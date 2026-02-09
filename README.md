# KubeStellar Console Integrations

Production-ready integrations for the **KubeStellar Console** — an AI-powered multi-cluster Kubernetes dashboard.

**LFX Mentorship Project**: CNCF - KubeStellar: Integration and ecosystem development specialist (2026 Term 1)

## Overview

This project adds two major integrations to KubeStellar Console:

1. **Prometheus/Grafana** — Historical metrics, alerts, embedded dashboards
2. **GitHub Actions** — CI/CD pipeline visibility, workflow triggers

## What is KubeStellar Console?

KubeStellar Console (kc) is a proactive dashboard that:

- Adapts to your workflow using AI
- Shows multi-cluster Kubernetes data in real-time
- Auto-swaps dashboard cards based on your focus

[Learn more →](https://github.com/kubestellar/console)

## Quick Start

### Prerequisites

- KubeStellar Console installed
- Prometheus + Grafana (for metrics integration)
- GitHub Personal Access Token (for GitHub integration)

### Installation

```bash
# 1. Clone Console (if not already)
git clone https://github.com/kubestellar/console.git
cd console
./scripts/prod.sh

# 2. Clone kss-demo and install integrations
git clone https://github.com/kubestellar/kss-demo.git
cd kss-demo
./scripts/install-console-integrations.sh
```

Set `CONSOLE_DIR` if your Console is elsewhere:

```bash
CONSOLE_DIR=/path/to/console ./scripts/install-console-integrations.sh
```

### Configure Integrations

Add to Console `.env`:

```bash
PROMETHEUS_URL=http://localhost:9090
GRAFANA_URL=http://localhost:3000
GRAFANA_API_KEY=your_grafana_api_key
GITHUB_TOKEN=ghp_xxxxxxxx
```

### Using the Integrations

1. Open Console at http://localhost:5174
2. Sign in with GitHub
3. Add Card → Prometheus Metrics, Grafana Dashboard, or GitHub Workflows
4. Configure each card (metric, dashboard UID, owner/repo)

## Documentation

Full documentation: [https://kss-demo.netlify.app](https://kss-demo.netlify.app) (when deployed)

- [Console Installation](docs-site/pages/getting-started/console-installation.mdx)
- [Prometheus/Grafana Integration](docs-site/pages/integrations/prometheus-grafana/index.mdx)
- [GitHub Actions Integration](docs-site/pages/integrations/github-actions/index.mdx)

## Project Structure

```
kss-demo/
├── console-integrations/   # Go packages (Prometheus, Grafana, GitHub)
├── web-cards/             # React dashboard cards
├── docs-site/             # Nextra documentation (Netlify)
├── examples/              # Example configurations
├── integrations/          # Legacy: ArgoCD, Terraform (standalone)
└── scripts/               # Installation scripts
```

## Contributing

Contributions welcome. Open an issue or PR in the repository.

## License

Apache License 2.0

## Acknowledgments

- CNCF LFX Mentorship Program
- KubeStellar maintainers and Console contributors
- Mentors: Rishi Mondal, Andy Anderson, Shivam Kumar, Naman Jain, Onkar Shelke
