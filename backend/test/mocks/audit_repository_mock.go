package mocks

import (
	"context"
	"globepay/internal/domain/model"
	"github.com/stretchr/testify/mock"
)

type AuditRepositoryMock struct {
	mock.Mock
}

func (m *AuditRepositoryMock) Create(ctx context.Context, auditLog *model.AuditLog) error {
	args := m.Called(ctx, auditLog)
	return args.Error(0)
}

func (m *AuditRepositoryMock) GetByUser(ctx context.Context, userID string, limit, offset int) ([]*model.AuditLog, error) {
	args := m.Called(ctx, userID, limit, offset)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]*model.AuditLog), args.Error(1)
}

func (m *AuditRepositoryMock) GetByAction(ctx context.Context, action string, limit, offset int) ([]*model.AuditLog, error) {
	args := m.Called(ctx, action, limit, offset)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]*model.AuditLog), args.Error(1)
}

func (m *AuditRepositoryMock) GetByTable(ctx context.Context, tableName string, limit, offset int) ([]*model.AuditLog, error) {
	args := m.Called(ctx, tableName, limit, offset)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.([]*model.AuditLog), args.Error(1)
}