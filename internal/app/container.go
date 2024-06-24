package app

import (
	"database/sql"

	"user_registry/internal/adapter/pgsql/userrepo"
	"user_registry/internal/adapter/redis/storage"
	"user_registry/internal/service/incrementor"
	signservice "user_registry/internal/service/signer"
	"user_registry/internal/service/userregistry"
	"user_registry/internal/usecase"

	"github.com/redis/go-redis/v9"
)

type Container struct {
	redis *redis.Client
	pgsql *sql.DB
}

func NewContainer(redisClieint *redis.Client, pgsqlConnect *sql.DB) *Container {
	return &Container{
		redis: redisClieint,
		pgsql: pgsqlConnect,
	}
}

func (c *Container) GetUseCase() *usecase.UseCase {
	return usecase.New(
		c.getHMACsigner(),
		c.getUserRegistry(),
		c.getIncrementor(),
	)
}

func (c *Container) getHMACsigner() *signservice.Service {
	return signservice.New()
}

func (c *Container) getUserRegistry() *userregistry.Service {
	return userregistry.New(userrepo.New(c.pgsql))
}

func (c *Container) getIncrementor() *incrementor.Service {
	return incrementor.New(storage.New(c.redis))
}
