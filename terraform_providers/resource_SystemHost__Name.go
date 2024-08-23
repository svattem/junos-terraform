
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
type xmlSystemHost__Name struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_  struct {
			XMLName xml.Name `xml:""`
			V_host__name  *string  `xml:"host-name,omitempty"`
		} `xml:"system"`
	} `xml:"groups"`

}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSystemHost__NameCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_host__name := d.Get("host__name").(string)


	config := xmlSystemHost__Name{}
	config.Groups.Name = id
	config.Groups.V_.V_host__name = &V_host__name

    err = client.SendTransaction("", config, false)
    check(ctx, err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSystemHost__NameRead(ctx,d,m)
}

func junosSystemHost__NameRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSystemHost__Name{}

	err = client.MarshalGroup(id, config)
	check(ctx, err)
 	if err :=d.Set("host__name", config.Groups.V_.V_host__name);err != nil{
		return diag.FromErr(err)
	}

	return nil
}

func junosSystemHost__NameUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_host__name := d.Get("host__name").(string)


	config := xmlSystemHost__Name{}
	config.Groups.Name = id
	config.Groups.V_.V_host__name = &V_host__name

    err = client.SendTransaction(id, config, false)
    check(ctx, err)
    
	return junosSystemHost__NameRead(ctx,d,m)
}

func junosSystemHost__NameDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	_, err = client.DeleteConfig(id,false)
    check(ctx, err)

    d.SetId("")
    
	return nil
}

func junosSystemHost__Name() *schema.Resource {
	return &schema.Resource{
		CreateContext: junosSystemHost__NameCreate,
		ReadContext: junosSystemHost__NameRead,
		UpdateContext: junosSystemHost__NameUpdate,
		DeleteContext: junosSystemHost__NameDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"host__name": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_. Hostname for this router",
			},
		},
	}
}