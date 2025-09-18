package usecase

import (
	"context"
	"time"
	"unchained/server/config"
	"unchained/server/gimport"
	"unchained/server/internal/entity/auth"
	"unchained/server/internal/entity/global"
	"unchained/server/internal/entity/sms"
	"unchained/server/rimport"
	"unchained/server/tools/arbit"
	"unchained/server/tools/logger"
)

type Auth struct {
	ri     *rimport.Repository
	ge     *gimport.Getaway
	log    *logger.Logger
	config *config.Config
}

func NewAuth(
	ri *rimport.Repository,
	ge *gimport.Getaway,
	log *logger.Logger,
	config *config.Config,
) *Auth {
	return &Auth{ri, ge, log, config}
}

func (u *Auth) CreateVerificationCode(ctx context.Context, params auth.CreateVerificationCodeParam) error {
	if _, exists := auth.AllowedCountryCodesMap[params.CountryCode]; !exists {
		return auth.ErrCountryCodeNotAllowed
	}

	fullNumber := params.FullNumber()

	existSession, err := u.ri.AuthCache.GetSmsSession(ctx, fullNumber)
	if err != nil && err != global.ErrNoData {
		u.log.Db.Errorln("не удалось достать смс сессию из кеша:", err)
		return global.ErrInternalError
	}

	var (
		attemptCount = 1
	)

	if !existSession.CreateTime.IsZero() {
		existSessionExpireTime := existSession.CreateTime.Add(u.config.SmsSessionTTL)
		if existSession.CreateTime.Before(existSessionExpireTime) {
			return auth.ErrSmsSessionAlreadyCreated
		}

		if existSession.AttemptCount >= u.config.SmsAtemptLimit {
			return auth.ErrOutOfSmsAttemptLimit
		}

		attemptCount = existSession.AttemptCount + 1
	}

	verificationCode := arbit.GenerateRandDigits(u.config.VerificationCodeLen)

	verificationCodeSmsParams := sms.NewSendVirifyCodeParams(params.PhoneNumber, verificationCode)
	if err := u.ge.Sms.SendVirifyCode(ctx, verificationCodeSmsParams); err != nil {
		u.log.Db.Errorln("не удалось отправить смс с кодом верификации:", err)
		return global.ErrInternalError
	}

	smsSession := auth.NewSmsSession(
		params,
		verificationCode,
		time.Now(),
		attemptCount,
	)
	if err := u.ri.AuthCache.SetSmsSession(ctx, fullNumber, smsSession); err != nil {
		u.log.Db.Errorln("не удалось отправить смс с кодом верификации:", err)
		return global.ErrInternalError
	}

	return nil
}

// func (u *Auth)
