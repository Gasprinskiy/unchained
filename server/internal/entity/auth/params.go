package auth

import "fmt"

type CreateVerificationCodeParam struct {
	CountryCode string `json:"country_code"`
	PhoneNumber string `json:"phone_number"`
}

func (p CreateVerificationCodeParam) FullNumber() string {
	return fmt.Sprintf("%s-%s", p.CountryCode, p.PhoneNumber)
}

type VerifyCodeParam struct {
	CreateParam      CreateVerificationCodeParam `json:"create_param"`
	VerificationCode string                      `json:"verification_code"`
}

func (p VerifyCodeParam) IsEqual(verificationCode string) bool {
	return p.VerificationCode == verificationCode
}
