package auth

import "time"

type SmsSession struct {
	CreateParam      CreateVerificationCodeParam `json:"create_param"`
	VerificationCode string                      `json:"verification_code"`
	CreateTime       time.Time                   `json:"create_time"`
	AttemptCount     int                         `json:"attempt_count"`
}

func NewSmsSession(
	createParam CreateVerificationCodeParam,
	verificationCode string,
	createTime time.Time,
	attemptCount int,
) SmsSession {
	return SmsSession{
		CreateParam:      createParam,
		VerificationCode: verificationCode,
		CreateTime:       createTime,
		AttemptCount:     attemptCount,
	}
}
