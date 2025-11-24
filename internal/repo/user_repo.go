package repo

import (
	"context"
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
	select {
	// simulate DB lookup
	case <-time.After(1000 * time.Millisecond):
		return "Bruce", nil
	case <-ctx.Done():
		return "", nil
	}
}
