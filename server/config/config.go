package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

// Config структура для хранения переменных окружения
type Config struct {
	PostgresURL         string
	GRPCAddr            string
	ServerPort          string
	RedisAddr           string
	RedisPass           string
	RedisTTL            time.Duration
	JwtSecret           string
	JwtSecretTTL        time.Duration
	SmsSessionTTL       time.Duration
	SmsAtemptLimit      int
	VerificationCodeLen int
}

// NewConfig загружает переменные из .env и возвращает структуру Config
func NewConfig() *Config {
	redisTtl, err := strconv.Atoi(os.Getenv("REDIS_TTL"))
	if err != nil {
		log.Panic("не удалось получить время жизни кеша: ", err)
	}

	jwtSecretTtl, err := strconv.Atoi(os.Getenv("JWT_SECRET_TTL"))
	if err != nil {
		log.Panic("не удалось получить время жизни jwt токена: ", err)
	}

	smsSessionTTL, err := strconv.Atoi(os.Getenv("SMS_SESSION_TTL"))
	if err != nil {
		log.Panic("не удалось получить время жизни сессии sms верификации: ", err)
	}

	smsAtemptLimit, err := strconv.Atoi(os.Getenv("SMS_ATEMPT_LIMIT"))
	if err != nil {
		log.Panic("не удалось получить лимит попыток запроса смс: ", err)
	}

	verificationCodeLen, err := strconv.Atoi(os.Getenv("VERIFICATION_CODE_LEN"))
	if err != nil {
		log.Panic("не удалось получить лимит попыток запроса смс: ", err)
	}

	return &Config{
		PostgresURL:         os.Getenv("POSTGRES_URL"),
		GRPCAddr:            os.Getenv("GRPC_ADDR"),
		ServerPort:          fmt.Sprintf(":%s", os.Getenv("HTTP_SERVER_PORT")),
		RedisPass:           os.Getenv("REDIS_PASSWORD"),
		RedisAddr:           fmt.Sprintf("redis:%s", os.Getenv("REDIS_PORT")),
		RedisTTL:            time.Minute * time.Duration(redisTtl),
		JwtSecretTTL:        time.Hour * time.Duration(jwtSecretTtl),
		SmsSessionTTL:       time.Minute * time.Duration(smsSessionTTL),
		SmsAtemptLimit:      smsAtemptLimit,
		VerificationCodeLen: verificationCodeLen,
	}
}
