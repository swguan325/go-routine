package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go-routine/internal/repo"
)

type AuthService interface {
	Login(ctx context.Context, userID, password string) (string, error)
}

type authServiceImpl struct {
	userRepo repo.UserRepo
}

func NewAuthService(userRepo repo.UserRepo) AuthService {
	return &authServiceImpl{
		userRepo: userRepo,
	}
}

func (s *authServiceImpl) Login(ctx context.Context, userID, password string) (string, error) {
	ok, err := s.userRepo.VerifyPassword(ctx, userID, password)
	if err != nil {
		return "", err
	}
	if !ok {
		return "", errors.New("invalid credentials")
	}

	token := fmt.Sprintf("token-%s-%d", userID, time.Now().UnixNano())
	return token, nil
}
