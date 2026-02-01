#!/usr/bin/env bash
# setup-dev-env.sh — Create a minimal dev environment for ks-demo (Kind + KubeStellar placeholder)
set -euo pipefail

echo "KubeStellar Integrations Demo — Dev environment setup"
echo "This script is a placeholder. For full KubeStellar demo env, use:"
echo "  bash <(curl -s https://raw.githubusercontent.com/kubestellar/kubestellar/main/scripts/create-kubestellar-demo-env.sh) --platform kind"
echo ""
echo "To install only the ArgoCD integration CRD and controller (after KubeStellar is installed):"
echo "  ./scripts/install-integrations.sh"
echo ""
echo "To run the docs site locally:"
echo "  cd docs-site && npm install && npm run dev"
