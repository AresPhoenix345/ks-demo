# ArgoCD Integration CRD bases

Place generated CRD YAML here (e.g. from `controller-gen` or `kubebuilder`).

Example:

```bash
# From integrations/argocd, after installing controller-gen:
controller-gen crd paths="./api/..." output:crd:dir=./config/crd/bases
```

Then `kubectl apply -f config/crd/bases/` to install the ArgoCDBinding CRD.
