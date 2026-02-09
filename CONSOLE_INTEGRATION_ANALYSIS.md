# Console Integration Analysis

Analysis of KubeStellar Console (console-main) for integration development.

## Key Integration Points

### 1. Dashboard Cards

- **Model**: `pkg/models/card.go` — CardType, Card, CardPosition, GetCardTypes()
- **Registry**: `web/src/lib/widgets/widgetRegistry.ts` — WIDGET_CARDS, WidgetCardDefinition
- **Criteria**: `component-criteria.md` — Card patterns (Data List, Metric, Chart, Single Select, Specialized)
- **Props**: Cards receive `config` (JSON) and optional cluster/context

### 2. Backend API

- **Server**: `pkg/api/server.go` — Fiber app, Config, routes
- **Handlers**: `pkg/api/handlers/` — auth, cards, dashboard, events, mcp, etc.
- **Pattern**: Handlers use `*fiber.Ctx`, return JSON or error
- **MCP Bridge**: `pkg/mcp/bridge.go` — Wraps MCP servers as HTTP/WS APIs

### 3. Frontend Components

- **Path**: `web/src/components/` — 439+ files
- **Unified system**: `web/src/lib/unified/` — UnifiedCard, UnifiedDashboard, UnifiedStats
- **Hooks**: `useCardData()`, `useChartFilters()`, `useSingleSelectCluster()`
- **API client**: Fetch from `/api/...` (proxied to backend)

### 4. Database Schema

- **Store**: `pkg/store/store.go`, `sqlite.go`
- **Tables**: Users, dashboards, cards, feedback, etc.
- **Integration settings**: Can extend schema for per-user Prometheus/Grafana URLs

### 5. Environment Variables

- **.env.example**: GITHUB_CLIENT_ID, GITHUB_CLIENT_SECRET, DEV_MODE, FRONTEND_URL, etc.
- **Integration vars**: Add PROMETHEUS_URL, GRAFANA_URL, GRAFANA_API_KEY, GITHUB_TOKEN

## Recommended Patterns

1. **New card type**: Add to GetCardTypes() in card.go; add React component; register in widgetRegistry
2. **New API route**: Add handler in pkg/api/handlers/; register route in server.go
3. **Integration package**: Create pkg/integrations/<name>/ with client, handlers
4. **Card config**: Use json.RawMessage in Card.Config; parse in frontend

## Console Tech Stack

- **Backend**: Go, Fiber v2, SQLite, gorilla/websocket
- **Frontend**: React 18, Vite, TypeScript, Tailwind
- **Ports**: Backend 8080, Frontend 5174, kc-agent 8585

## Files to Extend

| File | Purpose |
|------|---------|
| pkg/models/card.go | Add CardType constants and GetCardTypes entries |
| pkg/api/server.go | Add routes for /api/integrations/* |
| pkg/api/handlers/ | New handlers: prometheus.go, grafana.go, github.go |
| web/src/lib/widgets/widgetRegistry.ts | Add WIDGET_CARDS entries |
| web/src/components/cards/ | New card components |
