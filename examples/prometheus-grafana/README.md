# Prometheus/Grafana Integration Examples

## Files

- `prometheus-config.yaml` — Example Prometheus scrape config

## Console Card Config

Add to your dashboard:

```json
{
  "card_type": "prometheus_metrics",
  "config": {
    "metric": "container_cpu_usage_seconds_total",
    "cluster": "",
    "timeRange": "1h"
  }
}
```

## Grafana Embed

```json
{
  "card_type": "grafana_dashboard",
  "config": {
    "dashboardUid": "your-dashboard-uid",
    "from": "now-1h",
    "to": "now"
  }
}
```
