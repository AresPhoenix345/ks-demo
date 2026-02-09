# Console Integration Cards

React components for KubeStellar Console dashboard. Copy these into your console fork:

```
web/src/components/cards/
├── PrometheusMetricsCard.tsx
├── GrafanaDashboardCard.tsx
└── GitHubWorkflowsCard.tsx
```

Then register them in your card registry (e.g. `lib/widgets/widgetRegistry.ts` or equivalent) and add the corresponding card types in `pkg/models/card.go` (GetCardTypes).

## Card Types to Add

- `prometheus_metrics` — PrometheusMetricsCard
- `grafana_dashboard` — GrafanaDashboardCard
- `github_workflows` — GitHubWorkflowsCard

## API Endpoints

The backend must expose:

- `GET /api/integrations/prometheus/metrics?cluster=&metric=&range=`
- `GET /api/integrations/grafana/embed?uid=&from=&to=`
- `GET /api/integrations/github/workflows?owner=&repo=`

These handlers use the packages in `../console-integrations/pkg/integrations/`.
