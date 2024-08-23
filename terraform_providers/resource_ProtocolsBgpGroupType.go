
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
type xmlProtocolsBgpGroupType struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_group  struct {
			XMLName xml.Name `xml:"group"`
			V_name  *string  `xml:"name,omitempty"`
			V_type  *string  `xml:"type,omitempty"`
		} `xml:"protocols>bgp>group"`
	} `xml:"groups"`

}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosProtocolsBgpGroupTypeCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_type := d.Get("type").(string)


	config := xmlProtocolsBgpGroupType{}
	config.Groups.Name = id
	config.Groups.V_group.V_name = &V_name
	config.Groups.V_group.V_type = &V_type

    err = client.SendTransaction("", config, false)
    check(ctx, err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosProtocolsBgpGroupTypeRead(ctx,d,m)
}

func junosProtocolsBgpGroupTypeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlProtocolsBgpGroupType{}

	err = client.MarshalGroup(id, config)
	check(ctx, err)
 	if err :=d.Set("name", config.Groups.V_group.V_name);err != nil{
		return diag.FromErr(err)
	}
	if err :=d.Set("type", config.Groups.V_group.V_type);err != nil{
		return diag.FromErr(err)
	}

	return nil
}

func junosProtocolsBgpGroupTypeUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_type := d.Get("type").(string)


	config := xmlProtocolsBgpGroupType{}
	config.Groups.Name = id
	config.Groups.V_group.V_name = &V_name
	config.Groups.V_group.V_type = &V_type

    err = client.SendTransaction(id, config, false)
    check(ctx, err)
    
	return junosProtocolsBgpGroupTypeRead(ctx,d,m)
}

func junosProtocolsBgpGroupTypeDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	_, err = client.DeleteConfig(id,false)
    check(ctx, err)

    d.SetId("")
    
	return nil
}

func junosProtocolsBgpGroupType() *schema.Resource {
	return &schema.Resource{
		CreateContext: junosProtocolsBgpGroupTypeCreate,
		ReadContext: junosProtocolsBgpGroupTypeRead,
		UpdateContext: junosProtocolsBgpGroupTypeUpdate,
		DeleteContext: junosProtocolsBgpGroupTypeDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"name": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group",
			},
			"type": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group. Type of peer group",
			},
		},
	}
}