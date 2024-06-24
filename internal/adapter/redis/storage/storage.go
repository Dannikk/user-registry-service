package storage

import (
	"context"
	"errors"
	"user_registry/internal/service/incrementor"

	"github.com/redis/go-redis/v9"
)

type Storage struct {
	redis *redis.Client
}

var _ incrementor.KeyValueRepo = &Storage{}

func New(redis *redis.Client) *Storage {
	return &Storage{
		redis: redis,
	}
}

func (repo *Storage) UpdateKeyValue(ctx context.Context, key string, newVal int64) (int64, error) {
	value, err := repo.redis.Do(ctx, "get", key).Int64()
	if err != nil && !errors.Is(err, redis.Nil) {
		return 0, err
	}

	if err := repo.redis.Set(ctx, key, value+newVal, 0).Err(); err != nil {
		return 0, err
	}

	return value + newVal, nil
}
