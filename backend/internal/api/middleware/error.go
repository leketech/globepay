package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Error represents an API error
type Error struct {
	Error     string    `json:"error"`
	Code      string    `json:"code"`
	Timestamp time.Time `json:"timestamp"`
}

// ErrorHandler handles errors and returns standardized responses
func ErrorHandler(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set request time for error responses
		c.Set("request_time", time.Now())

		// Process request
		c.Next()

		// Handle any errors that occurred
		if len(c.Errors) > 0 {
			// Log the error
			for _, err := range c.Errors {
				logger.WithFields(logrus.Fields{
					"error": err.Error(),
					"meta":  err.Meta,
				}).Error("Request error")
			}

			// Return the last error
			lastError := c.Errors.Last()
			c.JSON(http.StatusInternalServerError, Error{
				Error:     lastError.Error(),
				Code:      "INTERNAL_ERROR",
				Timestamp: time.Now(),
			})
			return
		}
	}
}
