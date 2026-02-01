# Terraform simple example — KubeStellar BindingPolicy via provider (stub).
# Configure the provider with kubeconfig for your KubeStellar WDS.

terraform {
  required_version = ">= 1.0"
  # When provider is published or using dev overrides:
  # required_providers {
  #   kubestellar = {
  #     source  = "kubestellar/kubestellar"
  #     version = "~> 0.1"
  #   }
  # }
}

provider "kubestellar" {
  # kubeconfig = "/path/to/kubeconfig"  # or set KUBECONFIG env
}

resource "kubestellar_binding_policy" "example" {
  name = "terraform-example-policy"
  # cluster_selectors and downsync (schema TBD in full implementation)
}
