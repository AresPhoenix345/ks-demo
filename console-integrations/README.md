# KubeStellar Console Integrations

Integration packages for the KubeStellar Console: Prometheus/Grafana observability and GitHub Actions CI/CD.

## Structure

```
pkg/integrations/
├── prometheus/   # Prometheus API client, metrics, alerts
├── grafana/      # Grafana API client, dashboards, embed URLs
└── github/       # GitHub Actions API client, workflows, dispatch
```

## Usage

These packages are designed to be **merged into** a fork of [kubestellar/console](https://github.com/kubestellar/console):

1. Copy `pkg/integrations/` into your console repo under `pkg/`.
2. Add API handlers that use these clients (e.g. `/api/integrations/prometheus/metrics`).
3. Add the corresponding React card components in `web/src/components/cards/`.
4. Register new card types in `pkg/models/card.go` (GetCardTypes) and the frontend card registry.

## Environment Variables

- `PROMETHEUS_URL` — Prometheus server URL (e.g. http://localhost:9090)
- `GRAFANA_URL` — Grafana server URL
- `GRAFANA_API_KEY` — Grafana API key for listing dashboards
- `GITHUB_TOKEN` — GitHub PAT for Actions API (workflows, dispatch)

## See Also

- [Prometheus/Grafana Integration Docs](../docs-site/pages/integrations/prometheus-grafana/index.mdx)
- [GitHub Actions Integration Docs](../docs-site/pages/integrations/github-actions/index.mdx)
