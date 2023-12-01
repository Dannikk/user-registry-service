package app

import (
	"database/sql"
	"user_registry/internal/usecase"

	"github.com/redis/go-redis/v9"

	"user_registry/internal/adapter/pgsql/userrepo"
	"user_registry/internal/adapter/redis/storage"
	"user_registry/internal/service/incrementor"
	"user_registry/internal/service/signer"
	"user_registry/internal/service/user_reg"
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

func (c *Container) getHMACsigner() *sign_service.Service {
	return sign_service.New()
}

func (c *Container) getUserRegistry() *user_reg.Service {
	return user_reg.New(userrepo.New(c.pgsql))
}

func (c *Container) getIncrementor() *incrementor.Service {
	return incrementor.New(storage.New(c.redis))
}
