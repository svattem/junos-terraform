
package main

import (
    "context"
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/diag"

)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlRouting__OptionsRouter__Id struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_  struct {
			XMLName xml.Name `xml:""`
			V_router__id  *string  `xml:"router-id,omitempty"`
		} `xml:"routing-options"`
	} `xml:"groups"`

}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosRouting__OptionsRouter__IdCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_router__id := d.Get("router__id").(string)


	config := xmlRouting__OptionsRouter__Id{}
	config.Groups.Name = id
	config.Groups.V_.V_router__id = &V_router__id

    err = client.SendTransaction("", config, false)
    check(ctx, err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosRouting__OptionsRouter__IdRead(ctx,d,m)
}

func junosRouting__OptionsRouter__IdRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlRouting__OptionsRouter__Id{}

	err = client.MarshalGroup(id, config)
	check(ctx, err)
 	if err :=d.Set("router__id", config.Groups.V_.V_router__id);err != nil{
		return diag.FromErr(err)
	}

	return nil
}

func junosRouting__OptionsRouter__IdUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_router__id := d.Get("router__id").(string)


	config := xmlRouting__OptionsRouter__Id{}
	config.Groups.Name = id
	config.Groups.V_.V_router__id = &V_router__id

    err = client.SendTransaction(id, config, false)
    check(ctx, err)
    
	return junosRouting__OptionsRouter__IdRead(ctx,d,m)
}

func junosRouting__OptionsRouter__IdDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	_, err = client.DeleteConfig(id,false)
    check(ctx, err)

    d.SetId("")
    
	return nil
}

func junosRouting__OptionsRouter__Id() *schema.Resource {
	return &schema.Resource{
		CreateContext: junosRouting__OptionsRouter__IdCreate,
		ReadContext: junosRouting__OptionsRouter__IdRead,
		UpdateContext: junosRouting__OptionsRouter__IdUpdate,
		DeleteContext: junosRouting__OptionsRouter__IdDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"router__id": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_. Router identifier",
			},
		},
	}
}