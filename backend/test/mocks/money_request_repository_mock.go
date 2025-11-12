package mocks

import (
	"context"
	"time"

	"github.com/stretchr/testify/mock"

	"globepay/internal/domain/model"
)

type MoneyRequestRepositoryMock struct {
	mock.Mock
}

func (m *MoneyRequestRepositoryMock) Create(ctx context.Context, request *model.MoneyRequest) error {
	args := m.Called(ctx, request)
	return args.Error(0)
}

func (m *MoneyRequestRepositoryMock) GetByID(ctx context.Context, id string) (*model.MoneyRequest, error) {
	args := m.Called(ctx, id)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*model.MoneyRequest), args.Error(1)
}

func (m *MoneyRequestRepositoryMock) GetByRequester(ctx context.Context, requesterID string) ([]*model.MoneyRequest, error) {
	args := m.Called(ctx, requesterID)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]*model.MoneyRequest), args.Error(1)
}

func (m *MoneyRequestRepositoryMock) GetByRecipient(ctx context.Context, recipientID string) ([]*model.MoneyRequest, error) {
	args := m.Called(ctx, recipientID)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]*model.MoneyRequest), args.Error(1)
}

func (m *MoneyRequestRepositoryMock) UpdateStatus(ctx context.Context, id, status string, paidAt *time.Time) error {
	args := m.Called(ctx, id, status, paidAt)
	return args.Error(0)
}

func (m *MoneyRequestRepositoryMock) UpdatePaymentLink(ctx context.Context, id, paymentLink string) error {
	args := m.Called(ctx, id, paymentLink)
	return args.Error(0)
}
