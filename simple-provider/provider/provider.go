package simpleprovider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	entityclient "simple-provider/client"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"simpleprovider_item": resourceItem(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	// username := d.Get("username").(string)
	// password := d.Get("password").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	// diags = append(diags, diag.Diagnostic{
	// 	Severity: diag.Warning,
	// 	Summary:  "Warning Message Summary",
	// 	Detail:   "This is the detailed warning message from providerConfigure",
	// })
	/*
		if (username != "") && (password != "") {
			c, err := hashicups.NewClient(nil, &username, &password)
			if err != nil {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Unable to create HashiCups client",
					Detail:   "Unable to auth user for authenticated HashiCups client",
				})
				return nil, diags
			}

			return c, diags
		}
	*/
	c, err := entityclient.NewClient(nil, nil, nil)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create HashiCups client",
			Detail:   "Unable to auth user for authenticated HashiCups client",
		})
		return nil, diags
	}

	return c, diags
}
