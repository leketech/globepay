package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

// TracingMiddleware adds tracing to HTTP requests
func TracingMiddleware(tracer opentracing.Tracer) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract tracing context from headers
		spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))

		// Start a new span
		span := tracer.StartSpan(c.FullPath(), ext.RPCServerOption(spanCtx))
		defer span.Finish()

		// Add HTTP-specific tags
		ext.HTTPMethod.Set(span, c.Request.Method)
		ext.HTTPUrl.Set(span, c.Request.URL.String())
		ext.Component.Set(span, "gin")

		// Store span in context
		c.Set("tracing-span", span)

		// Process request
		c.Next()

		// Set status code tag
		statusCode := c.Writer.Status()
		if statusCode >= 0 && statusCode <= 65535 {
			ext.HTTPStatusCode.Set(span, uint16(statusCode))
		}

		// Mark as error if status is 5xx
		if statusCode >= 500 {
			ext.Error.Set(span, true)
		}
	}
}
