package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Metrics holds all application metrics
type Metrics struct {
	HTTPRequestsTotal    *prometheus.CounterVec
	HTTPRequestDuration  *prometheus.HistogramVec
	HTTPRequestsInFlight prometheus.Gauge
	UsersRegisteredTotal prometheus.Counter
	LoginAttemptsTotal   prometheus.Counter
	TransfersTotal       prometheus.Counter
	TransferAmountTotal  prometheus.Counter
}

// NewMetrics creates a new metrics instance
func NewMetrics() *Metrics {
	return &Metrics{
		HTTPRequestsTotal: promauto.NewCounterVec(prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		}, []string{"method", "path", "status"}),

		HTTPRequestDuration: promauto.NewHistogramVec(prometheus.HistogramOpts{
			Name: "http_request_duration_seconds",
			Help: "HTTP request duration in seconds",
		}, []string{"method", "path"}),

		HTTPRequestsInFlight: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "http_requests_in_flight",
			Help: "Number of HTTP requests currently in flight",
		}),

		UsersRegisteredTotal: promauto.NewCounter(prometheus.CounterOpts{
			Name: "users_registered_total",
			Help: "Total number of registered users",
		}),

		LoginAttemptsTotal: promauto.NewCounter(prometheus.CounterOpts{
			Name: "login_attempts_total",
			Help: "Total number of login attempts",
		}),

		TransfersTotal: promauto.NewCounter(prometheus.CounterOpts{
			Name: "transfers_total",
			Help: "Total number of money transfers",
		}),

		TransferAmountTotal: promauto.NewCounter(prometheus.CounterOpts{
			Name: "transfer_amount_total",
			Help: "Total amount of money transferred",
		}),
	}
}