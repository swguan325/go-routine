package repo

import (
	"context"
	"errors"
	"time"
)

type UserRepo interface {
	VerifyPassword(ctx context.Context, userID, plainPwd string) (bool, error)
	GetUsername(ctx context.Context, userID string) (string, error)
}

type userRepoImpl struct{}

func NewUserRepo() UserRepo {
	return &userRepoImpl{}
}

func (r *userRepoImpl) VerifyPassword(ctx context.Context, userID, plainPwd string) (bool, error) {
	// simulate DB lookup
	time.Sleep(50 * time.Millisecond)

	// demo rule
	if userID == "user123" && plainPwd == "pw123" {
		return true, nil
	}
	return false, nil
}

func (r *userRepoImpl) GetUsername(ctx context.Context, userID string) (string, error) {
	// simulate DB lookup
	time.Sleep(1000 * time.Millisecond)

	if userID == "user123" {
		return "Bruce", nil
	}
	return "", errors.New("user not found")
}
