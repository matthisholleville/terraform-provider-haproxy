package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/matthisholleville/terraform-provider-haproxy/internal/haproxy"
)

func resourceMaps() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceMapsCreate,
		ReadContext:   resourceMapsRead,
		UpdateContext: resourceMapsUpdate,
		DeleteContext: resourceMapsDelete,

		Schema: map[string]*schema.Schema{
			"map": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"value": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "defaultValue",
			},
			"force_sync": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}

func resourceMapsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*haproxy.Client)
	mapName := d.Get("map").(string)
	newEntrie := &haproxy.MapEntrie{
		Key:   d.Get("key").(string),
		Value: d.Get("value").(string),
	}
	forceSync := d.Get("force_sync").(bool)
	_, err := client.CreateMapEntrie(newEntrie, mapName, forceSync)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(newEntrie.Key)

	return nil
}

func resourceMapsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceMapsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*haproxy.Client)

	if d.HasChange("value") {
		mapName := d.Get("map").(string)
		entrie, err := client.GetMapEntrie(d.Get("key").(string), mapName)
		if err != nil {
			diag.FromErr(err)
		}

		entrie.Value = d.Get("value").(string)

		_, err = client.UpdateMapEntrie(entrie, mapName, d.Get("force_sync").(bool))
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceMapsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*haproxy.Client)
	mapName := d.Get("map").(string)
	forceSync := d.Get("force_sync").(bool)
	err := client.DeleteMapEntrie(d.Id(), mapName, forceSync)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return nil
}
