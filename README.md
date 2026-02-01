# KubeStellar Integrations Demo

Production-ready integrations for KubeStellar multi-cluster orchestration.

**LFX Mentorship Project**: CNCF - KubeStellar: Integration and ecosystem development specialist (2026 Term 1)

## Project Overview

This project demonstrates two production-ready integrations for KubeStellar:

1. **ArgoCD Integration** — GitOps-based multi-cluster deployments using BindingPolicy and Argo CD Applications.
2. **Terraform Integration** — Manage KubeStellar BindingPolicies as Terraform resources (IaC).

Built as part of the CNCF LFX Mentorship Program 2026 Term 1.

## Quick Start

### Prerequisites

- Docker Desktop or similar
- kubectl 1.29+
- Helm 3.x
- Node.js 18+

### Install KubeStellar

```bash
bash <(curl -s https://raw.githubusercontent.com/kubestellar/kubestellar/main/scripts/install.sh)
```

### Install Integrations

```bash
git clone https://github.com/kubestellar/ks-demo.git
cd ks-demo
./scripts/install-integrations.sh
```

### Run Documentation Site Locally

```bash
cd docs-site
npm install
npm run dev
```

Visit http://localhost:3000

## Documentation

Full documentation is available in the docs-site (and when deployed at the Netlify URL):

- [Getting Started](docs-site/pages/getting-started/index.mdx)
- [ArgoCD Integration](docs-site/pages/integrations/argocd/index.mdx)
- [Terraform Integration](docs-site/pages/integrations/terraform/index.mdx)
- [API Reference](docs-site/pages/api-reference/index.mdx)

## Project Structure

```
ks-demo/
├── integrations/        # Integration implementations (ArgoCD, Terraform)
├── docs-site/           # Nextra documentation website (Netlify-deployable)
├── examples/            # Sample configurations
├── scripts/             # Automation scripts
└── tests/               # Test suites
```

## Testing

```bash
# Run integration tests (from repo root)
make -C integrations/argocd test
```

## Deployment

The documentation site is configured for Netlify:

- **Build**: `docs-site` base, `npm run build`, publish `.next`
- **Preview**: Deploy previews per PR

See [DEPLOYMENT.md](DEPLOYMENT.md) for details.

## Contributing

Contributions are welcome. See [CONTRIBUTING.md](CONTRIBUTING.md) if present, or open an issue/PR in the repository.

## License

Apache License 2.0

## Acknowledgments

- CNCF LFX Mentorship Program
- KubeStellar maintainers and community
- Mentors: Rishi Mondal, Andy Anderson, Shivam Kumar, Naman Jain, Onkar Shelke
