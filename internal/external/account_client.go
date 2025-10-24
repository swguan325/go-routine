package external

import (
	"context"
	"time"
)

type AccountClient interface {
	GetBalance(ctx context.Context, userID string) (float64, error)
}

type accountClientImpl struct{}

func NewAccountClient() AccountClient {
	return &accountClientImpl{}
}

func (a *accountClientImpl) GetBalance(ctx context.Context, userID string) (float64, error) {
	// simulate API latency
	time.Sleep(1200 * time.Millisecond)
	return 88888.00, nil
}
