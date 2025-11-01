# Monitoring Documentation

This document provides comprehensive information about the monitoring setup for the Globepay application.

## Overview

Globepay uses a full observability stack consisting of:

- **Prometheus** for metrics collection
- **Grafana** for visualization and dashboards
- **Loki** for log aggregation
- **Promtail** for log shipping
- **Jaeger** for distributed tracing
- **Alertmanager** for alerting

## Architecture

```
┌─────────────────┐    ┌──────────────┐    ┌──────────────┐
│   Application   │    │   Metrics    │    │    Logs      │
│     (Go/React)  │───▶│  (Prometheus)│───▶│   (Loki)     │
└─────────────────┘    └──────────────┘    └──────────────┘
                              │                   │
                              ▼                   ▼
                       ┌──────────────┐    ┌──────────────┐
                       │ Visualization│    │ Visualization│
                       │   (Grafana)  │    │   (Grafana)  │
                       └──────────────┘    └──────────────┘
                              │
                              ▼
                       ┌──────────────┐
                       │   Alerting   │
                       │(Alertmanager)│
                       └──────────────┘
```

## Metrics Collection

### Prometheus Configuration

Prometheus is configured to scrape metrics from:

1. **Kubernetes Services** - API server, nodes, pods
2. **Application Endpoints** - `/metrics` endpoints
3. **Database** - PostgreSQL exporter
4. **Cache** - Redis exporter
5. **Infrastructure** - Node exporter

**Configuration File:**
```yaml
# monitoring/prometheus/prometheus-config.yaml
global:
  scrape_interval: 15s
  evaluation_interval: 15s

scrape_configs:
  - job_name: 'kubernetes-apiservers'
    kubernetes_sd_configs:
    - role: endpoints
    scheme: https
    tls_config:
      ca_file: /var/run/secrets/kubernetes.io/serviceaccount/ca.crt
    bearer_token_file: /var/run/secrets/kubernetes.io/serviceaccount/token

  - job_name: 'globepay-backend'
    static_configs:
    - targets: ['backend.globepay-prod.svc.cluster.local:9090']

  - job_name: 'postgres-exporter'
    static_configs:
    - targets: ['postgres-exporter.monitoring.svc.cluster.local:9187']

  - job_name: 'redis-exporter'
    static_configs:
    - targets: ['redis-exporter.monitoring.svc.cluster.local:9121']
```

### Application Metrics

The backend exposes the following metrics:

#### HTTP Metrics
- `http_requests_total` - Total HTTP requests
- `http_request_duration_seconds` - HTTP request duration
- `http_requests_in_flight` - Current HTTP requests in flight

#### Business Metrics
- `transfers_total` - Total money transfers
- `transfer_amount_total` - Total transfer amount
- `users_registered_total` - Total registered users
- `transactions_processed_total` - Total processed transactions

#### System Metrics
- `go_goroutines` - Number of goroutines
- `go_memstats_alloc_bytes` - Memory allocation
- `process_cpu_seconds_total` - CPU usage

**Example Metric Endpoint:**
```go
// backend/internal/api/middleware/metrics.go
func MetricsMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        c.Next()
        duration := time.Since(start)
        
        // Record metrics
        httpRequestsTotal.WithLabelValues(
            c.Request.Method,
            c.FullPath(),
            strconv.Itoa(c.Writer.Status()),
        ).Inc()
        
        httpRequestDuration.WithLabelValues(
            c.Request.Method,
            c.FullPath(),
        ).Observe(duration.Seconds())
    }
}
```

## Dashboards

### Grafana Dashboards

Grafana provides the following pre-configured dashboards:

#### 1. API Dashboard

Shows API performance metrics:
- Request rate and error rate
- Latency percentiles (P50, P95, P99)
- Top endpoints by request volume
- HTTP status code distribution

#### 2. Infrastructure Dashboard

Displays infrastructure health:
- Node CPU and memory usage
- Pod resource utilization
- Network I/O
- Disk I/O

#### 3. Database Dashboard

Monitors database performance:
- Connection pool usage
- Query performance
- Table sizes
- Index usage

#### 4. Business Metrics Dashboard

Tracks business KPIs:
- Daily transfer volume
- User growth trends
- Revenue metrics
- Geographic distribution

### Dashboard Configuration

Dashboards are provisioned through ConfigMaps:

```yaml
# monitoring/grafana/dashboards/api-dashboard.json
{
  "dashboard": {
    "id": null,
    "title": "API Metrics",
    "panels": [
      {
        "title": "Request Rate",
        "type": "graph",
        "targets": [
          {
            "expr": "rate(http_requests_total[5m])",
            "legendFormat": "{{method}} {{path}} {{status}}"
          }
        ]
      }
    ]
  }
}
```

## Logging

### Loki Configuration

Loki is configured to collect logs from:

1. **Application Pods** - Backend and frontend logs
2. **Infrastructure Pods** - Kubernetes system components
3. **Node Logs** - System logs from worker nodes

**Promtail Configuration:**
```yaml
# monitoring/promtail/promtail-config.yaml
server:
  http_listen_port: 9080
  grpc_listen_port: 0

positions:
  filename: /tmp/positions.yaml

clients:
  - url: http://loki:3100/loki/api/v1/push

scrape_configs:
  - job_name: kubernetes-pods
    kubernetes_sd_configs:
      - role: pod
    relabel_configs:
      - source_labels: [__meta_kubernetes_pod_annotation_promtail_io_scrape]
        action: keep
        regex: true
      - source_labels: [__meta_kubernetes_pod_name]
        target_label: pod
```

### Log Structure

Application logs follow a structured format:

```json
{
  "timestamp": "2023-01-01T00:00:00Z",
  "level": "info",
  "service": "backend",
  "component": "transfer-service",
  "message": "Transfer created successfully",
  "transfer_id": "12345",
  "user_id": "67890",
  "amount": 100.50
}
```

### Log Queries

Common log queries in Grafana:

1. **Error Logs**: `{level="error"} |= "transfer"`
2. **User Activity**: `{user_id="12345"}`
3. **Performance Issues**: `{duration_ms > 1000}`
4. **Authentication**: `{component="auth-service"} |= "login"`

## Distributed Tracing

### Jaeger Configuration

Jaeger collects traces from:

1. **API Requests** - HTTP request tracing
2. **Database Queries** - SQL query tracing
3. **External Calls** - Third-party service calls
4. **Background Jobs** - Worker process tracing

**Application Tracing Setup:**
```go
// backend/internal/api/middleware/tracing.go
import (
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/exporters/jaeger"
    "go.opentelemetry.io/otel/sdk/resource"
    sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func InitTracing() (*sdktrace.TracerProvider, error) {
    exp, err := jaeger.New(jaeger.WithCollectorEndpoint(
        jaeger.WithEndpoint("http://jaeger-collector:14268/api/traces"),
    ))
    if err != nil {
        return nil, err
    }

    tp := sdktrace.NewTracerProvider(
        sdktrace.WithBatcher(exp),
        sdktrace.WithResource(resource.NewWithAttributes(
            semconv.SchemaURL,
            semconv.ServiceNameKey.String("globepay-backend"),
        )),
    )
    
    otel.SetTracerProvider(tp)
    return tp, nil
}
```

### Trace Structure

Traces include the following spans:

1. **HTTP Request Span** - Overall request processing
2. **Authentication Span** - User authentication
3. **Database Span** - Database operations
4. **External Service Span** - Third-party API calls

## Alerting

### Alertmanager Configuration

Alertmanager handles alert routing and notification:

```yaml
# monitoring/alertmanager/alertmanager-config.yaml
route:
  group_by: ['alertname']
  group_wait: 10s
  group_interval: 10s
  repeat_interval: 1h
  receiver: 'slack-notifications'

receivers:
  - name: 'slack-notifications'
    slack_configs:
      - api_url: 'https://hooks.slack.com/services/YOUR/SLACK/WEBHOOK'
        channel: '#alerts'
        send_resolved: true
```

### Alert Rules

#### Critical Alerts

1. **High Error Rate**
   ```yaml
   - alert: HighErrorRate
     expr: rate(http_requests_total{status=~"5.."}[5m]) / rate(http_requests_total[5m]) > 0.05
     for: 5m
     labels:
       severity: critical
     annotations:
       summary: "High error rate detected"
   ```

2. **High Latency**
   ```yaml
   - alert: HighLatency
     expr: histogram_quantile(0.95, rate(http_request_duration_seconds_bucket[5m])) > 1
     for: 5m
     labels:
       severity: critical
     annotations:
       summary: "High API latency detected"
   ```

3. **Database Connection Pool Exhausted**
   ```yaml
   - alert: DatabaseConnectionPoolExhausted
     expr: db_connections_in_use / db_connections_max > 0.9
     for: 5m
     labels:
       severity: critical
     annotations:
       summary: "Database connection pool nearly exhausted"
   ```

#### Warning Alerts

1. **High Memory Usage**
   ```yaml
   - alert: HighMemoryUsage
     expr: container_memory_usage_bytes / container_spec_memory_limit_bytes > 0.8
     for: 10m
     labels:
       severity: warning
     annotations:
       summary: "High memory usage detected"
   ```

2. **High CPU Usage**
   ```yaml
   - alert: HighCPUUsage
     expr: rate(container_cpu_usage_seconds_total[5m]) > 0.8
     for: 10m
     labels:
       severity: warning
     annotations:
       summary: "High CPU usage detected"
   ```

3. **Low Cache Hit Rate**
   ```yaml
   - alert: LowCacheHitRate
     expr: redis_keyspace_hits_total / (redis_keyspace_hits_total + redis_keyspace_misses_total) < 0.8
     for: 10m
     labels:
       severity: warning
     annotations:
       summary: "Low Redis cache hit rate"
   ```

### Notification Channels

Alerts are sent to:

1. **Slack** - #alerts channel
2. **Email** - operations@globepay.com
3. **PagerDuty** - For critical alerts
4. **Webhook** - For integration with other systems

## Health Checks

### Application Health Endpoints

Applications expose health check endpoints:

1. **Liveness Probe** - `/health`
   ```json
   {
     "status": "healthy",
     "timestamp": "2023-01-01T00:00:00Z"
   }
   ```

2. **Readiness Probe** - `/ready`
   ```json
   {
     "status": "ready",
     "database": "connected",
     "cache": "connected",
     "timestamp": "2023-01-01T00:00:00Z"
   }
   ```

### Kubernetes Probes

Kubernetes uses these probes for pod management:

```yaml
# k8s/base/backend/deployment.yaml
livenessProbe:
  httpGet:
    path: /health
    port: http
  initialDelaySeconds: 30
  periodSeconds: 10

readinessProbe:
  httpGet:
    path: /ready
    port: http
  initialDelaySeconds: 10
  periodSeconds: 5
```

## Monitoring Commands

### Accessing Monitoring Tools

```bash
# Port-forward Grafana
kubectl port-forward -n monitoring svc/prometheus-grafana 3000:80

# Port-forward Prometheus
kubectl port-forward -n monitoring svc/prometheus-kube-prometheus-prometheus 9090:9090

# Port-forward Loki
kubectl port-forward -n monitoring svc/loki 3100:3100

# Port-forward Jaeger
kubectl port-forward -n monitoring svc/jaeger-query 16686:16686
```

### Useful Queries

1. **API Error Rate:**
   ```
   rate(http_requests_total{status=~"5.."}[5m]) / rate(http_requests_total[5m])
   ```

2. **95th Percentile Latency:**
   ```
   histogram_quantile(0.95, rate(http_request_duration_seconds_bucket[5m]))
   ```

3. **Database Connections:**
   ```
   db_connections_in_use / db_connections_max
   ```

4. **Memory Usage:**
   ```
   container_memory_usage_bytes / container_spec_memory_limit_bytes
   ```

## Troubleshooting

### Common Issues

1. **Metrics Not Showing**
   - Check if the application exposes `/metrics` endpoint
   - Verify Prometheus target configuration
   - Check network connectivity

2. **Logs Not Appearing**
   - Verify Promtail configuration
   - Check if pods have correct annotations
   - Ensure Loki is running

3. **Alerts Not Firing**
   - Check Alertmanager configuration
   - Verify alert rules syntax
   - Test alert expressions in Prometheus

### Debugging Commands

```bash
# Check Prometheus targets
kubectl get targets -n monitoring

# Check Grafana dashboards
kubectl get configmaps -n monitoring | grep dashboard

# Check Loki readiness
kubectl get pods -n monitoring -l app=loki

# Check Alertmanager alerts
kubectl get alerts -n monitoring
```

This monitoring documentation provides a comprehensive guide to the observability setup for the Globepay application, ensuring proper monitoring, alerting, and troubleshooting capabilities.