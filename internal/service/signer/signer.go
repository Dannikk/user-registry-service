package sign_service

import (
	"context"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"user_registry/internal/entity"
	"user_registry/internal/usecase"
)

type Service struct{}

var _ usecase.Signer = &Service{}

func New() *Service {
	return &Service{}
}

func (s *Service) GetSign(ctx context.Context, tk *entity.TextKey) (string, error) {
	hasher := hmac.New(sha512.New, []byte(tk.Key))

	if _, err := hasher.Write([]byte(tk.Text)); err != nil {
		return "", err
	}
	summa := hasher.Sum(nil)
	hexstring := hex.EncodeToString(summa)

	return hexstring, nil
}
