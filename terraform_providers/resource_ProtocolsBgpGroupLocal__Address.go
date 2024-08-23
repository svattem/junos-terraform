
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
type xmlProtocolsBgpGroupLocal__Address struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_group  struct {
			XMLName xml.Name `xml:"group"`
			V_name  *string  `xml:"name,omitempty"`
			V_local__address  *string  `xml:"local-address,omitempty"`
		} `xml:"protocols>bgp>group"`
	} `xml:"groups"`

}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosProtocolsBgpGroupLocal__AddressCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_local__address := d.Get("local__address").(string)


	config := xmlProtocolsBgpGroupLocal__Address{}
	config.Groups.Name = id
	config.Groups.V_group.V_name = &V_name
	config.Groups.V_group.V_local__address = &V_local__address

    err = client.SendTransaction("", config, false)
    check(ctx, err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosProtocolsBgpGroupLocal__AddressRead(ctx,d,m)
}

func junosProtocolsBgpGroupLocal__AddressRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlProtocolsBgpGroupLocal__Address{}

	err = client.MarshalGroup(id, config)
	check(ctx, err)
 	if err :=d.Set("name", config.Groups.V_group.V_name);err != nil{
		return diag.FromErr(err)
	}
	if err :=d.Set("local__address", config.Groups.V_group.V_local__address);err != nil{
		return diag.FromErr(err)
	}

	return nil
}

func junosProtocolsBgpGroupLocal__AddressUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_local__address := d.Get("local__address").(string)


	config := xmlProtocolsBgpGroupLocal__Address{}
	config.Groups.Name = id
	config.Groups.V_group.V_name = &V_name
	config.Groups.V_group.V_local__address = &V_local__address

    err = client.SendTransaction(id, config, false)
    check(ctx, err)
    
	return junosProtocolsBgpGroupLocal__AddressRead(ctx,d,m)
}

func junosProtocolsBgpGroupLocal__AddressDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	_, err = client.DeleteConfig(id,false)
    check(ctx, err)

    d.SetId("")
    
	return nil
}

func junosProtocolsBgpGroupLocal__Address() *schema.Resource {
	return &schema.Resource{
		CreateContext: junosProtocolsBgpGroupLocal__AddressCreate,
		ReadContext: junosProtocolsBgpGroupLocal__AddressRead,
		UpdateContext: junosProtocolsBgpGroupLocal__AddressUpdate,
		DeleteContext: junosProtocolsBgpGroupLocal__AddressDelete,

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
			"local__address": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group. Address of local end of BGP session",
			},
		},
	}
}