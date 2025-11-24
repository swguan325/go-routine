package service

import (
	"context"
	"fmt"
	"sync"
	"time"

	"go-routine/internal/external"
	"go-routine/internal/model"
	"go-routine/internal/repo"
)

type DashboardService interface {
	GetDashboard(ctx context.Context, userID string) (*model.Dashboard, error)
}

type dashboardServiceImpl struct {
	cardClient    external.CardClient
	accountClient external.AccountClient
	userRepo      repo.UserRepo
	timeout       time.Duration
}

func NewDashboardService(
	cardClient external.CardClient,
	accountClient external.AccountClient,
	userRepo repo.UserRepo,
) DashboardService {
	// 指向 dashboardServiceImpl 的指標
	return &dashboardServiceImpl{
		cardClient:    cardClient,
		accountClient: accountClient,
		userRepo:      userRepo,
		timeout:       1250 * time.Millisecond,
	}
}

func (s *dashboardServiceImpl) GetDashboard(ctx context.Context, userID string) (*model.Dashboard, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(3)

	var (
		cardsResult []model.Card
		cardsErr    error

		balanceResult float64
		balanceErr    error

		usernameResult string
		usernameErr    error
	)

	// goroutine 1: credit cards
	go func() {
		defer wg.Done()
		cardsResult, cardsErr = s.cardClient.GetCreditCards(ctx, userID)
	}()

	// goroutine 2: balance
	go func() {
		defer wg.Done()
		balanceResult, balanceErr = s.accountClient.GetBalance(ctx, userID)
	}()

	// goroutine 3: balance
	go func() {
		defer wg.Done()
		usernameResult, usernameErr = s.userRepo.GetUsername(ctx, userID)
	}()

	wg.Wait()

	// unifying error checks
	if cardsErr != nil {
		return nil, fmt.Errorf("GetCreditCards failed: %w", cardsErr)
	}
	if balanceErr != nil {
		return nil, fmt.Errorf("GetBalance failed: %w", balanceErr)
	}
	if usernameErr != nil {
		return nil, fmt.Errorf("GetUsername failed: %w", usernameErr)
	}

	return &model.Dashboard{
		Username: usernameResult,
		Balance:  balanceResult,
		Cards:    cardsResult,
	}, nil
}
