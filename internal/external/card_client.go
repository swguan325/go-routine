package external

import (
	"context"
	"time"

	"go-routine/internal/model"
)

type CardClient interface {
	GetCreditCards(ctx context.Context, userID string) ([]model.Card, error)
}

type cardClientImpl struct{}

func NewCardClient() CardClient {
	return &cardClientImpl{}
}

func (c *cardClientImpl) GetCreditCards(ctx context.Context, userID string) ([]model.Card, error) {
	select {
	// simulate API latency
	case <-time.After(1500 * time.Millisecond):
		return []model.Card{
			{Number: "****-****-****-1234", Brand: "VISA"},
			{Number: "****-****-****-5678", Brand: "Mastercard"},
		}, nil
	case <-ctx.Done():
		return []model.Card{}, nil
	}
}
