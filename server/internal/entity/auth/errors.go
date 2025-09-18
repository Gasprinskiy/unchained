package auth

import "errors"

var (
	ErrSmsSessionAlreadyCreated = errors.New("SmsSessionAlreadyCreated")
	ErrOutOfSmsAttemptLimit     = errors.New("OutOfSmsAttemptLimit")
	ErrCountryCodeNotAllowed    = errors.New("CountryCodeNotAllowed")
)
