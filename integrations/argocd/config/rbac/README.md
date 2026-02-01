# ArgoCD Integration RBAC

Place RBAC manifests here for the ArgoCD integration controller:

- `service_account.yaml`
- `role.yaml`
- `role_binding.yaml`

Then `kubectl apply -f config/rbac/` to install (after creating the namespace if needed).
