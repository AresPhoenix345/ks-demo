/*
Copyright 2025 The KubeStellar Authors.

Terraform provider for KubeStellar (stub).
*/

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func New() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"kubestellar_binding_policy": resourceBindingPolicy(),
		},
		DataSourcesMap: map[string]*schema.Resource{},
		Schema: map[string]*schema.Schema{
			"kubeconfig": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("KUBECONFIG", nil),
				Description: "Path to kubeconfig for KubeStellar WDS.",
			},
		},
		ConfigureContextFunc: configureProvider,
	}
}

func configureProvider(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	// TODO: build KubeStellar client from kubeconfig
	return struct{}{}, nil
}

func resourceBindingPolicy() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceBindingPolicyCreate,
		ReadContext:   resourceBindingPolicyRead,
		UpdateContext: resourceBindingPolicyUpdate,
		DeleteContext: resourceBindingPolicyDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cluster_selectors": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeMap},
			},
		},
	}
}

func resourceBindingPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// TODO: create BindingPolicy via KubeStellar client
	d.SetId(d.Get("name").(string))
	return nil
}

func resourceBindingPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceBindingPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return resourceBindingPolicyRead(ctx, d, meta)
}

func resourceBindingPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// TODO: delete BindingPolicy
	return nil
}
