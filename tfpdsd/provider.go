package tfpdsd

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"dsd_api": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("DSD_API", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"reserved_name": resourceReservedName(),
		},
		DataSourcesMap:       map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, data *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	apiServer := data.Get("dsd_api").(string)
	if apiServer != "" {
		return apiServer, diags
	} else {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "You need to specify an API Server",
			Detail:   "Specify the dsd_api configuration or DSD_API",
		})
		return nil, diags
	}
}
