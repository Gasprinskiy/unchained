package rimport

import (
	"unchained/server/config"
	"unchained/server/internal/repository"

	"github.com/redis/go-redis/v9"
)

type Repository struct {
	AuthCache repository.AuthCache
}

func NewRepositoryImports(
	config *config.Config,
	rdb *redis.Client,
) *Repository {
	return &Repository{}
}
