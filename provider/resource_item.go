package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"regexp"
	"strings"
	"terraform-provider-example/api/client"
	"terraform-provider-example/api/server"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func validateName(v interface{}, k string) (ws []string, es []error) {
	var errs []error
	var warns []string
	value, ok := v.(string)
	if !ok {
		errs = append(errs, fmt.Errorf("Expected name to be string"))
		return warns, errs
	}
	whiteSpace := regexp.MustCompile(`\s+`)
	if whiteSpace.Match([]byte(value)) {
		errs = append(errs, fmt.Errorf("name cannot contain whitespace. Got %s", value))
		return warns, errs
	}
	return warns, errs
}

func resourceItem() *schema.Resource {
	fmt.Print()
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "The name of the resource, also acts as it's unique ID",
				ForceNew:     true,
				ValidateFunc: validateName,
			},
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A description of an item",
			},
		},
		CreateContext: resourceCreateItem,
		ReadContext:   resourceReadItem,
		UpdateContext: resourceUpdateItem,
		DeleteContext: resourceDeleteItem,
	}
}

func resourceCreateItem(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)
	item := server.Item{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
	}

	err := apiClient.NewItem(&item)

	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(item.Name)
	return nil
}

func resourceReadItem(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)

	itemId := d.Id()
	item, err := apiClient.GetItem(itemId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			d.SetId("")
		} else {
			return diag.FromErr(err)
		}
	}

	d.SetId(item.Name)
	d.Set("name", item.Name)
	d.Set("description", item.Description)
	return nil
}

func resourceUpdateItem(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)

	item := server.Item{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
	}

	err := apiClient.UpdateItem(&item)
	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceDeleteItem(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*client.Client)

	itemId := d.Id()

	err := apiClient.DeleteItem(itemId)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return nil
}
