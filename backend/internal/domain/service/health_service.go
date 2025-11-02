package service

import (
	"context"
	"database/sql"
	"time"

	"globepay/internal/infrastructure/cache"
)

// HealthService provides health check functionality
type HealthService struct {
	db    *sql.DB
	cache *cache.RedisClient
}

// NewHealthService creates a new health service
func NewHealthService(db *sql.DB, cache *cache.RedisClient) *HealthService {
	return &HealthService{
		db:    db,
		cache: cache,
	}
}

// CheckDatabase checks database connectivity
func (h *HealthService) CheckDatabase(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return h.db.PingContext(ctx)
}

// CheckCache checks cache connectivity
func (h *HealthService) CheckCache(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return h.cache.Ping(ctx)
}

// CheckAll checks all services
func (h *HealthService) CheckAll(ctx context.Context) map[string]string {
	results := make(map[string]string)

	// Check database
	if err := h.CheckDatabase(ctx); err != nil {
		results["database"] = "disconnected"
	} else {
		results["database"] = "connected"
	}

	// Check cache
	if err := h.CheckCache(ctx); err != nil {
		results["cache"] = "disconnected"
	} else {
		results["cache"] = "connected"
	}

	return results
}