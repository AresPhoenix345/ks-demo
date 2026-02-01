# ArgoCD Basic Example

This example creates a KubeStellar BindingPolicy and an ArgoCDBinding so that the ArgoCD integration controller creates one Argo CD Application per destination cluster.

## Prerequisites

- KubeStellar installed with at least one BindingPolicy and resolved Binding (or equivalent cluster inventory).
- Argo CD installed.
- ArgoCD integration controller installed (see repo root `scripts/install-integrations.sh`).

## Files

- `bindingpolicy.yaml` — Sample BindingPolicy (adjust clusterSelectors and downsync to match your environment).
- `argocd-binding.yaml` — ArgoCDBinding that references the BindingPolicy and defines the Application template.

## Apply

```bash
kubectl apply -f bindingpolicy.yaml
kubectl apply -f argocd-binding.yaml
```

## Verify

- List Argo CD Applications: you should see one Application per destination from the BindingPolicy/Binding.
- Check ArgoCDBinding status: `kubectl get argocdbinding -o yaml`
