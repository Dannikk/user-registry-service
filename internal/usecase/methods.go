package usecase

import (
	"context"
	"user_registry/internal/entity"
)

func (uc *UseCase) Sign(ctx context.Context, tk *entity.TextKey) (string, error) {
	sign, err := uc.signer.GetSign(ctx, tk)

	return sign, err
}

func (uc *UseCase) CreateUser(ctx context.Context, user *entity.User) (int64, error) {
	res, err := uc.userReg.CreateUser(ctx, user)

	return res, err
}

func (uc *UseCase) Increment(ctx context.Context, kv *entity.KeyValue) (int64, error) {
	res, err := uc.incrementor.Increment(ctx, kv)

	return res, err
}
