# ArgoCD Integration Kustomize

Place Kustomize overlay here to deploy the ArgoCD integration controller:

- `kustomization.yaml` (resources: deployment, service, etc.)
- `deployment.yaml` (or reference from config/manager)
- `service.yaml` (optional)

Then `kubectl apply -k deploy/kustomize/` to deploy (after CRD and RBAC are installed).
