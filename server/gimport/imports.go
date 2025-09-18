package gimport

import (
	"unchained/server/config"
	"unchained/server/internal/gateway"

	"google.golang.org/grpc"
)

type Getaway struct {
	Sms gateway.Sms
}

func NewGetawayImports(
	grpcConn *grpc.ClientConn,
	config *config.Config,
) *Getaway {
	return &Getaway{}
}
