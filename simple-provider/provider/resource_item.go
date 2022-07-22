package simpleprovider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	entityclient "simple-provider/client"
)

type Entity struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func resourceItem() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
		CreateContext: resourceCreateItem,
		ReadContext:   resourceReadItem,
		UpdateContext: resourceUpdateItem,
		DeleteContext: resourceDeleteItem,
	}
}

func resourceCreateItem(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*entityclient.Client)

	var diags diag.Diagnostics

	item := entityclient.Entity{
		ID:          d.Get("id").(string),
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
	}

	token := "xxx"
	_, err := apiClient.CreateEntity(item, &token)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(item.ID)
	return diags
}

func resourceReadItem(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*entityclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	itemId := d.Id()

	token := "xxx"
	item, err := apiClient.GetEntity(itemId, &token)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(item.ID)
	err = d.Set("name", item.Name)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("description", item.Description)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceUpdateItem(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*entityclient.Client)

	var diags diag.Diagnostics

	item := entityclient.Entity{
		ID:          d.Get("id").(string),
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
	}
	token := "xxx"

	_, err := apiClient.UpdateEntity(item.ID, item, &token)
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceDeleteItem(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	apiClient := m.(*entityclient.Client)

	itemId := d.Id()
	token := "xxx"

	err := apiClient.DeleteEntity(itemId, &token)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return nil
}
