package sms

type SendVirifyCodeParams struct {
	PhoneNumber string
	Code        string
}

func NewSendVirifyCodeParams(phoneNumber, code string) SendVirifyCodeParams {
	return SendVirifyCodeParams{
		PhoneNumber: phoneNumber,
		Code:        code,
	}
}
