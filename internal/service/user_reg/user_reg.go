package user_reg

import (
	"context"
	"user_registry/internal/entity"
	"user_registry/internal/usecase"
)

type Repository interface {
	CreateUser(ctx context.Context, user *entity.User) (int64, error)
}

type Service struct {
	repo Repository
}

var _ usecase.UserRegistry = &Service{}

func New(repo Repository) *Service {
	return &Service{
		repo: repo}
}

func (s *Service) CreateUser(ctx context.Context, user *entity.User) (int64, error) {
	id, err := s.repo.CreateUser(ctx, user)

	return id, err
}
