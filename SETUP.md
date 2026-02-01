# Development Setup Guide

## System Requirements

- **OS**: Linux, macOS, or Windows with WSL2
- **RAM**: 8GB minimum, 16GB recommended
- **Disk**: 20GB free space
- **CPU**: 4 cores recommended

## Software Prerequisites

### Required Tools

1. **Docker Desktop** (or Docker Engine + Kind)
2. **kubectl** — [Install guide](https://kubernetes.io/docs/tasks/tools/)
3. **Helm 3.x** — `curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash`
4. **Go 1.21+** — [Download](https://go.dev/dl/)
5. **Node.js 18+** — [Download](https://nodejs.org/)
6. **Kind** (optional) — `go install sigs.k8s.io/kind@latest`

## Repository Setup

### Clone the Repository

```bash
git clone https://github.com/kubestellar/kss-demo.git
cd kss-demo
```

### Set Up Development Environment

```bash
./scripts/setup-dev-env.sh
```

This script can create Kind clusters, install KubeStellar, and prepare the environment (adjust script to your needs).

## Integration Development

### ArgoCD Integration

```bash
cd integrations/argocd
go mod tidy     # generate go.sum if missing
go mod download
make generate   # if codegen is configured
make build
make test
```

### Terraform Integration

```bash
cd integrations/terraform
go mod tidy     # generate go.sum if missing
go mod download
make build
make test
```

## Documentation Site Development

```bash
cd docs-site
npm install
npm run dev
```

Visit http://localhost:3000

**Production build:**

```bash
npm run build
npm start
```

## Troubleshooting

### Kind Cluster Issues

```bash
kind delete cluster --name kubeflex
# Re-run setup-dev-env.sh
```

### Documentation Build Issues

```bash
cd docs-site
rm -rf .next node_modules
npm install
npm run build
```

### Integration Controller Issues

```bash
kubectl logs -n kubestellar-system -l app=argocd-integration-controller --tail=100
```

## Next Steps

- Review [IMPLEMENTATION_PLAN.md](IMPLEMENTATION_PLAN.md)
- Read integration docs under `docs-site/pages/integrations/`
- Try examples in `examples/`
