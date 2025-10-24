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
	// simulate API latency
	time.Sleep(1500 * time.Millisecond)

	return []model.Card{
		{Number: "****-****-****-1234", Brand: "VISA"},
		{Number: "****-****-****-5678", Brand: "Mastercard"},
	}, nil
}
