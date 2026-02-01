# Terraform Simple Example

This example uses the Terraform provider to create a KubeStellar BindingPolicy.

## Prerequisites

- Terraform 1.x.
- KubeStellar WDS reachable (KUBECONFIG or in-cluster).
- Terraform provider built and available (see `integrations/terraform/` and [Installation](/integrations/terraform/installation)).

## Files

- `main.tf` — Provider configuration and one `kubestellar_binding_policy` resource (stub schema; full implementation may extend the resource schema).

## Run

```bash
terraform init
terraform plan
terraform apply
```

## Verify

```bash
kubectl get bindingpolicies.control.kubestellar.io
```
