package gateway

import "unchained/server/internal/entity/sms"

type Sms interface {
	SendVirifyCode(params sms.SendVirifyCodeParams) error
}
