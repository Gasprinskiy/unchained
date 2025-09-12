package auth

import (
	"strings"
	"time"
	"unchained/server/config"
	"unchained/server/internal/entity/global"
	"unchained/server/internal/entity/jwt_auth"
	"unchained/server/tools/logger"

	"github.com/golang-jwt/jwt/v5"
)

type Jwt struct {
	log    *logger.Logger
	config *config.Config
}

func NewJwt(log *logger.Logger, config *config.Config) *Jwt {
	return &Jwt{log, config}
}

func (j *Jwt) GenerateToken(userID int) (string, error) {
	expirationTime := time.Now().Add(j.config.JwtSecretTTL)

	claims := jwt_auth.Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := []byte(j.config.JwtSecret)

	result, err := token.SignedString(secretKey)
	if err != nil {
		err = global.ErrInternalError
	}

	return result, err
}

func (j *Jwt) ParseToken(tokenString string) (jwt_auth.Claims, error) {
	var zero jwt_auth.Claims

	token, err := jwt.ParseWithClaims(tokenString, &jwt_auth.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.config.JwtSecret), nil
	})

	if err != nil {
		errString := err.Error()
		if strings.Contains(errString, jwt.ErrTokenExpired.Error()) {
			return zero, global.ErrExpired
		}
		j.log.Db.Errorln("не спарсить токен пользователя:", err)
		return zero, global.ErrInternalError
	}

	claims, ok := token.Claims.(*jwt_auth.Claims)
	if !ok && !token.Valid {
		return zero, global.ErrInternalError
	}

	return *claims, nil
}
