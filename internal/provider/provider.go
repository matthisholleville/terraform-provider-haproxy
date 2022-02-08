package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/matthisholleville/terraform-provider-haproxy/internal/haproxy"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"server_addr": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("HAPROXY_SERVER", nil),
				Description: "HAProxy Dataplaneapi server address.",
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("HAPROXY_USERNAME", nil),
				Description: "Username use for authentification",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("HAPROXY_PASSWORD", nil),
				Description: "Password use for authentification",
			},
			"insecure": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Scheme for request. If not set, https will be use.",
				DefaultFunc: schema.EnvDefaultFunc("HAPROXY_INSECURE", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"haproxy_maps":     resourceMaps(),
			"haproxy_frontend": resourceFrontend(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	server_addr := d.Get("server_addr").(string)
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	insecure := d.Get("insecure").(bool)

	apiClient := haproxy.NewClient(username, password, server_addr, insecure)

	err := apiClient.TestApiCall()
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return apiClient, nil

}
