package provider

import (
	"context"
	"errors"

	"github.com/avast/retry-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/matthisholleville/terraform-provider-haproxy/internal/haproxy"
	"github.com/matthisholleville/terraform-provider-haproxy/internal/haproxy/models"
)

func resourceAcl() *schema.Resource {
	return &schema.Resource{
		Description:   "`haproxy_acl` manage acls.",
		CreateContext: resourceAclCreate,
		ReadContext:   resourceAclRead,
		UpdateContext: resourceAclUpdate,
		DeleteContext: resourceAclDelete,
		Schema: map[string]*schema.Schema{
			"parent_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Parent name.",
			},
			"parent_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Parent type. Possible value `frontend` or `backend`.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of the ACL to describe it as much as possible. It must have upper and lower case letters, digits, - (dash), _ (underscore) , . (dot) and : (colon). It is case sensitive; hence my_acl and My_Acl are two different ACLs.",
			},
			"criterion": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Based on sample fetches, it describes the portion of the request or response where this ACL applies.",
			},
			"index": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     nil,
				Description: "Acl line index in file.",
			},
			"value": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Data provided by <criterion> is compared to a<pattern> list.",
			},
		},
	}
}

func resourceAclRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceAclUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceAclDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceAclCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*haproxy.Client)
	parentName := d.Get("parent_name").(string)
	parentType := d.Get("parent_type").(string)
	newAcl := &models.ACL{
		ACLName:   d.Get("name").(string),
		Criterion: d.Get("criterion").(string),
		Index:     d.Get("index").(int),
		Value:     d.Get("value").(string),
	}
	err := retry.Do(
		func() error {
			configuration, err := client.GetConfiguration()
			if err != nil {
				return err
			}
			transaction, err := client.CreateTransaction(configuration.Version)
			if err != nil {
				return err
			}

			acls, err := client.GetAcls(transaction.Id, parentName, parentType)
			if err != nil {
				return err
			}

			for _, v := range *acls {
				if v.ACLName == newAcl.ACLName {
					return errors.New("Cannot insert " + newAcl.ACLName + " which already exist in " + parentName + " " + parentType + ".")
				}
			}

			_, err = client.CreateAcl(transaction.Id, newAcl, parentName, parentType)
			if err != nil {
				return err
			}

			_, err = client.CommitTransaction(transaction.Id)
			if err != nil {
				return err
			}
			return nil

		},
	)

	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(newAcl.ACLName)
	return nil

}
