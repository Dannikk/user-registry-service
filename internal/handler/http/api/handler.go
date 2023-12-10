package api

import (
	"context"
	"user_registry/internal/entity"
)

//go:generate mockgen -source=handler.go -destination=mocks/ucmock.go

type UseCase interface {
	Sign(ctx context.Context, tk *entity.TextKey) (string, error)
	CreateUser(ctx context.Context, user *entity.User) (int64, error)
	Increment(ctx context.Context, kv *entity.KeyValue) (int64, error)
}

type Handler struct {
	uc UseCase
}

func NewHandler(uc UseCase) *Handler {
	return &Handler{uc: uc}
}
