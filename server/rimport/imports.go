package rimport

import (
	"unchained/server/config"

	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

type RepositoryImports struct {
	Repository
}

func NewRepositoryImports(
	grpcConn *grpc.ClientConn,
	config *config.Config,
	rdb *redis.Client,
) *RepositoryImports {
	return &RepositoryImports{}
}
