package usecase

import (
	"context"
	"user_registry/internal/entity"
	handler "user_registry/internal/handler/http/api"
)

//go:generate mockgen -source=usecase.go -destination=mocks/mock_uc.go

type Signer interface {
	GetSign(ctx context.Context, tk *entity.TextKey) (string, error)
}

type UserRegistry interface {
	CreateUser(ctx context.Context, user *entity.User) (int64, error)
}

type Incrementor interface {
	Increment(ctx context.Context, kv *entity.KeyValue) (int64, error)
}

type UseCase struct {
	signer      Signer
	userReg     UserRegistry
	incrementor Incrementor
}

var _ handler.UseCase = &UseCase{}

func New(signer Signer, userReg UserRegistry, incrementor Incrementor) *UseCase {
	return &UseCase{
		signer,
		userReg,
		incrementor,
	}
}
