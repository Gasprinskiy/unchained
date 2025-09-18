package repository

import (
	"context"
	"unchained/server/internal/entity/auth"
)

type AuthCache interface {
	SetSmsSession(ctx context.Context, key string, param auth.SmsSession) error
	GetSmsSession(ctx context.Context, key string) (auth.SmsSession, error)
}

type UserProfiles interface {
}
