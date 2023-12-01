package incrementor

import (
	"context"
	"log"
	"user_registry/internal/entity"
	"user_registry/internal/usecase"
)


type KeyValueRepo interface {
	UpdateKeyValue(ctx context.Context, key string, value int64) (int64, error)
}

type Service struct {
	repo KeyValueRepo
}


var _ usecase.Incrementor = &Service{}


func New(repo KeyValueRepo) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Increment(ctx context.Context, kv *entity.KeyValue) (int64, error) {
	if val, err := s.repo.UpdateKeyValue(ctx, kv.Key, kv.Value); err != nil {
		log.Printf("Failed to increment: %v", err)
		return 0, err
	} else {
		return val, nil
	}
}
	