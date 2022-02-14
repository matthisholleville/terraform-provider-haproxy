package provider

import (
	"context"
	"errors"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/matthisholleville/terraform-provider-haproxy/internal/haproxy"
	"github.com/matthisholleville/terraform-provider-haproxy/internal/haproxy/models"
)

func resourceMaps() *schema.Resource {
	return &schema.Resource{
		Description:   "`haproxy_maps` manage maps.",
		CreateContext: resourceMapsCreate,
		ReadContext:   resourceMapsRead,
		UpdateContext: resourceMapsUpdate,
		DeleteContext: resourceMapsDelete,
		Importer: &schema.ResourceImporter{
			State: resourceMapEntrieImport,
		},

		Schema: map[string]*schema.Schema{
			"map": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The HAProxy map name. More informations : https://www.haproxy.com/fr/blog/introduction-to-haproxy-maps/",
				ValidateFunc: func(i interface{}, s string) ([]string, []error) {
					return validation.StringIsNotWhiteSpace(i, s)
				},
			},
			"key": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Key name",
				ValidateFunc: func(i interface{}, s string) ([]string, []error) {
					return validation.StringIsNotWhiteSpace(i, s)
				},
			},
			"value": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "defaultValue",
				Description: "Value name. Default value 'defaultValue'",
				ValidateFunc: func(i interface{}, s string) ([]string, []error) {
					return validation.StringIsNotWhiteSpace(i, s)
				},
			},
			"force_sync": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "If true, immediately syncs changes to disk",
			},
		},
	}
}

func resourceMapEntrieImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	client := meta.(*haproxy.Client)

	idMatchFormat, _ := regexp.MatchString("map/(.*?)/entrie/(.*?)", d.Id())
	if !idMatchFormat {
		return nil, fmt.Errorf("invalid format: expected map/<mapName>/entrie/<entrieName>, e.g. map/test/entrie/my-key, actual id is %s", d.Id())
	}

	mapName := haproxy.ExtractStringWithRegex(d.Id(), "map/(.*?)/")
	mapEntrie := haproxy.ExtractStringWithRegex(d.Id(), "entrie/(.*?)$")

	d.SetId(mapEntrie)
	d.Set("map", mapName)

	_, err := client.GetMapEntrie(mapEntrie, mapName)
	if err != nil {
		return nil, fmt.Errorf("error on getting map entrie during import: %s", err)
	}

	return []*schema.ResourceData{d}, nil
}

func resourceMapsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*haproxy.Client)
	mapName := d.Get("map").(string)
	newEntrie := &models.MapEntrie{
		Key:   d.Get("key").(string),
		Value: d.Get("value").(string),
	}
	forceSync := d.Get("force_sync").(bool)
	_, err := client.CreateMapEntrie(newEntrie, mapName, forceSync)
	if err != nil {
		return diag.FromErr(err)
	}

	_, err = client.GetMapEntrie(newEntrie.Key, mapName)
	if err != nil {
		errMessage := errors.New("Cannot insert " + newEntrie.Key + ". Space is not allowed.")
		return diag.FromErr(errMessage)
	}

	d.SetId(newEntrie.Key)

	return resourceMapsRead(ctx, d, meta)
}

func resourceMapsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*haproxy.Client)

	mapEntrie, err := client.GetMapEntrie(d.Id(), d.Get("map").(string))
	if err != nil {
		diag.FromErr(err)
	}
	d.Set("key", mapEntrie.Key)
	d.Set("value", mapEntrie.Value)
	d.Set("map", d.Get("map").(string))
	d.Set("force_sync", d.Get("force_sync").(bool))
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
