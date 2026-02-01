#!/usr/bin/env bash
# install-integrations.sh — Install ArgoCD (and optionally Terraform) integration CRDs and controllers
set -euo pipefail

REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$REPO_ROOT"

echo "Installing KubeStellar integrations from $REPO_ROOT"

# Check KubeStellar CRD exists
if ! kubectl get crd bindingpolicies.control.kubestellar.io &>/dev/null; then
  echo "KubeStellar not detected (bindingpolicies.control.kubestellar.io CRD missing)."
  echo "Install KubeStellar first: bash <(curl -s https://raw.githubusercontent.com/kubestellar/kubestellar/main/scripts/install.sh)"
  exit 1
fi

# ArgoCD integration: apply CRD and RBAC if present
if [[ -d "$REPO_ROOT/integrations/argocd/config/crd/bases" ]]; then
  echo "Applying ArgoCD integration CRDs..."
  kubectl apply -f "$REPO_ROOT/integrations/argocd/config/crd/bases/" || true
fi
if [[ -d "$REPO_ROOT/integrations/argocd/config/rbac" ]]; then
  echo "Applying ArgoCD integration RBAC..."
  kubectl apply -f "$REPO_ROOT/integrations/argocd/config/rbac/" || true
fi
if [[ -d "$REPO_ROOT/integrations/argocd/deploy/kustomize" ]]; then
  echo "Deploying ArgoCD integration controller..."
  kubectl apply -k "$REPO_ROOT/integrations/argocd/deploy/kustomize/" || true
fi

echo "Done. Verify with: kubectl get crd argocdbindings.integrations.kubestellar.io"
echo "Controller (when deployed): kubectl get pods -n kubestellar-system -l app=argocd-integration-controller"
