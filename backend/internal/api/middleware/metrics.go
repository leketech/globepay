package middleware

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"globepay/internal/infrastructure/metrics"
)

// MetricsMiddleware collects metrics for HTTP requests
func MetricsMiddleware(m *metrics.Metrics) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Increment in-flight requests
		m.HTTPRequestsInFlight.Inc()
		defer m.HTTPRequestsInFlight.Dec()

		start := time.Now()
		path := c.FullPath()

		// Process request
		c.Next()

		// Record metrics
		status := strconv.Itoa(c.Writer.Status())
		method := c.Request.Method

		// Record total requests
		m.HTTPRequestsTotal.WithLabelValues(method, path, status).Inc()

		// Record request duration
		duration := time.Since(start).Seconds()
		m.HTTPRequestDuration.WithLabelValues(method, path).Observe(duration)
	}
}