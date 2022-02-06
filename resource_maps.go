package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceMaps() *schema.Resource {
	return &schema.Resource{
		Create: resourceMapsCreate,
		Read:   resourceMapsRead,
		Update: resourceMapsUpdate,
		Delete: resourceMapsDelete,

		Schema: map[string]*schema.Schema{
			"map": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"force_sync": &schema.Schema{
				Type:     schema.TypeBool,
				Required: false,
				Default:  true,
			},
		},
	}
}

func resourceMapsCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceMapsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceMapsUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceMapsRead(d, m)
}

func resourceMapsDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
