package gateway

import (
	"context"
	"unchained/server/internal/entity/sms"
)

type Sms interface {
	SendVirifyCode(ctx context.Context, params sms.SendVirifyCodeParams) error
}
