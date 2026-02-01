/*
Copyright 2025 The KubeStellar Authors.

Terraform provider for KubeStellar BindingPolicy (stub).
*/

package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/kubestellar/kss-demo/integrations/terraform/internal/provider"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.New,
	})
}
