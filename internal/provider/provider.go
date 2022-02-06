package provider

import (
	"context"
	"log"

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
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("HAPROXY_USERNAME", nil),
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("HAPROXY_PASSWORD", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"haproxy_maps": resourceMaps(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	log := log.Default()

	log.Println(d.Get("server_addr").(string))
	server_addr := d.Get("server_addr").(string)
	username := d.Get("username").(string)
	password := d.Get("password").(string)

	apiClient := haproxy.NewClient(username, password, server_addr)

	err := apiClient.TestApiCall()
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return apiClient, nil

}
