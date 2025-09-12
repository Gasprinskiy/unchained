package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"unchained/server/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	config := config.NewConfig()

	// подключение к redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPass,
	})
	defer rdb.Close()

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		log.Panic("ошибка при пинге redis: ", err)
	}

	// подключение к postgres
	pgdb, err := sqlx.Connect("pgx", config.PostgresURL)
	if err != nil {
		log.Fatalln("не удалось подключиться к базе postgres: ", err)
	}
	defer pgdb.Close()

	if err := pgdb.Ping(); err != nil {
		log.Fatal("ошибка при пинге postgres : ", err)
	}

	// Настройка HTTP-сервера
	ginConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Referer", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}

	router := gin.Default()
	router.Use(cors.New(ginConfig))

	srv := &http.Server{
		Addr:    config.ServerPort,
		Handler: router,
	}

	// инициализация логгера
	// hook := logger.NewPostgresHook(pgdb)
	// logger, err := logger.InitLogger(hook)
	// if err != nil {
	// log.Fatalln("Не удалось инициализировать логгер:", err)
	// }

	// инициализация session manager
	// sessionManager := transaction.NewSQLSessionManager(pgdb)

	// инициализация grpc соеденения
	grpcConn, err := grpc.NewClient(config.GRPCAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panic("ошибка при подключении к grpc серверу: ", err)
	}
	defer grpcConn.Close()

	// инициализация репо
	// ri := rimport.NewRepositoryImports(grpcConn, config, rdb)

	// инициализация usecase
	// ui := uimport.NewUsecaseImport(ri, logger)

	// payme middleware

	// v1Router := router.Group("/api/v1")

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Panic("Ошибка при запуске HTTP-сервера: ", err)
	}
}
