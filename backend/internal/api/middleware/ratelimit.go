package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// RateLimiter handles rate limiting for API requests
type RateLimiter struct {
	limiters map[string]*rate.Limiter
	mu       sync.RWMutex
	rate     rate.Limit
	burst    int
}

// NewRateLimiter creates a new RateLimiter
func NewRateLimiter(r rate.Limit, burst int) *RateLimiter {
	return &RateLimiter{
		limiters: make(map[string]*rate.Limiter),
		rate:     r,
		burst:    burst,
	}
}

// getLimiter returns the rate limiter for the given key
func (rl *RateLimiter) getLimiter(key string) *rate.Limiter {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	limiter, exists := rl.limiters[key]
	if !exists {
		limiter = rate.NewLimiter(rl.rate, rl.burst)
		rl.limiters[key] = limiter
	}

	return limiter
}

// RateLimitMiddleware returns a Gin middleware function for rate limiting
func (rl *RateLimiter) RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Use IP address as the key for rate limiting
		key := c.ClientIP()

		// Get the rate limiter for this key
		limiter := rl.getLimiter(key)

		// Check if the request is allowed
		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Rate limit exceeded",
			})
			c.Abort()
			return
		}

		// Continue with the next middleware or handler
		c.Next()
	}
}

// RateLimitByUserMiddleware returns a Gin middleware function for rate limiting by user ID
func (rl *RateLimiter) RateLimitByUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get user ID from context (set by AuthMiddleware)
		userID, exists := c.Get("user_id")
		if !exists {
			// If user is not authenticated, use IP address
			c.Next()
			return
		}

		// Use user ID as the key for rate limiting
		key := userID.(string)

		// Get the rate limiter for this key
		limiter := rl.getLimiter(key)

		// Check if the request is allowed
		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Rate limit exceeded",
			})
			c.Abort()
			return
		}

		// Continue with the next middleware or handler
		c.Next()
	}
}

// SlidingWindowRateLimiter implements a sliding window rate limiter
type SlidingWindowRateLimiter struct {
	mu       sync.Mutex
	requests map[string][]time.Time
	window   time.Duration
	limit    int
}

// NewSlidingWindowRateLimiter creates a new sliding window rate limiter
func NewSlidingWindowRateLimiter(window time.Duration, limit int) *SlidingWindowRateLimiter {
	return &SlidingWindowRateLimiter{
		requests: make(map[string][]time.Time),
		window:   window,
		limit:    limit,
	}
}

// Allow checks if a request is allowed based on the sliding window algorithm
func (swrl *SlidingWindowRateLimiter) Allow(key string) bool {
	swrl.mu.Lock()
	defer swrl.mu.Unlock()

	now := time.Now()
	windowStart := now.Add(-swrl.window)

	// Remove old requests outside the window
	requests := swrl.requests[key]
	filtered := make([]time.Time, 0)
	for _, reqTime := range requests {
		if reqTime.After(windowStart) {
			filtered = append(filtered, reqTime)
		}
	}

	// Check if we're within the limit
	if len(filtered) >= swrl.limit {
		return false
	}

	// Add the current request
	swrl.requests[key] = append(filtered, now)
	return true
}

// SlidingWindowRateLimitMiddleware returns a Gin middleware function for sliding window rate limiting
func (swrl *SlidingWindowRateLimiter) SlidingWindowRateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Use IP address as the key for rate limiting
		key := c.ClientIP()

		// Check if the request is allowed
		if !swrl.Allow(key) {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Rate limit exceeded",
			})
			c.Abort()
			return
		}

		// Continue with the next middleware or handler
		c.Next()
	}
}