package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// LoggingMiddleware logs incoming requests
func LoggingMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		
		// Process request
		c.Next()
		
		// Log request details
		logger.WithFields(logrus.Fields{
			"method":     c.Request.Method,
			"path":       c.Request.URL.Path,
			"status":     c.Writer.Status(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
			"latency":    time.Since(start),
		}).Info("HTTP request")
	}
}

// RequestIDMiddleware adds a request ID to each request for tracking
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Generate or get request ID
		requestID := generateRequestID()

		// Set request ID in context
		c.Set("request_id", requestID)

		// Set request ID in response headers
		c.Header("X-Request-ID", requestID)

		// Continue with the next middleware or handler
		c.Next()
	}
}

// generateRequestID generates a unique request ID
func generateRequestID() string {
	// In a real implementation, you would generate a unique ID
	// For now, we'll use a simple timestamp-based ID
	return time.Now().Format("20060102150405.000")
}

// SecurityHeadersMiddleware adds security headers to responses
func SecurityHeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Add security headers
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Frame-Options", "DENY")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		c.Header("Content-Security-Policy", "default-src 'self'")

		// Continue with the next middleware or handler
		c.Next()
	}
}