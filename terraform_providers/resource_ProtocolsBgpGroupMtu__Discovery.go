
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
type xmlProtocolsBgpGroupMtu__Discovery struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_group  struct {
			XMLName xml.Name `xml:"group"`
			V_name  *string  `xml:"name,omitempty"`
			V_mtu__discovery  *string  `xml:"mtu-discovery,omitempty"`
		} `xml:"protocols>bgp>group"`
	} `xml:"groups"`

}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosProtocolsBgpGroupMtu__DiscoveryCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_mtu__discovery := d.Get("mtu__discovery").(string)


	config := xmlProtocolsBgpGroupMtu__Discovery{}
	config.Groups.Name = id
	config.Groups.V_group.V_name = &V_name
	config.Groups.V_group.V_mtu__discovery = &V_mtu__discovery

    err = client.SendTransaction("", config, false)
    check(ctx, err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosProtocolsBgpGroupMtu__DiscoveryRead(ctx,d,m)
}

func junosProtocolsBgpGroupMtu__DiscoveryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlProtocolsBgpGroupMtu__Discovery{}

	err = client.MarshalGroup(id, config)
	check(ctx, err)
 	if err :=d.Set("name", config.Groups.V_group.V_name);err != nil{
		return diag.FromErr(err)
	}
	if err :=d.Set("mtu__discovery", config.Groups.V_group.V_mtu__discovery);err != nil{
		return diag.FromErr(err)
	}

	return nil
}

func junosProtocolsBgpGroupMtu__DiscoveryUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_mtu__discovery := d.Get("mtu__discovery").(string)


	config := xmlProtocolsBgpGroupMtu__Discovery{}
	config.Groups.Name = id
	config.Groups.V_group.V_name = &V_name
	config.Groups.V_group.V_mtu__discovery = &V_mtu__discovery

    err = client.SendTransaction(id, config, false)
    check(ctx, err)
    
	return junosProtocolsBgpGroupMtu__DiscoveryRead(ctx,d,m)
}

func junosProtocolsBgpGroupMtu__DiscoveryDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	_, err = client.DeleteConfig(id,false)
    check(ctx, err)

    d.SetId("")
    
	return nil
}

func junosProtocolsBgpGroupMtu__Discovery() *schema.Resource {
	return &schema.Resource{
		CreateContext: junosProtocolsBgpGroupMtu__DiscoveryCreate,
		ReadContext: junosProtocolsBgpGroupMtu__DiscoveryRead,
		UpdateContext: junosProtocolsBgpGroupMtu__DiscoveryUpdate,
		DeleteContext: junosProtocolsBgpGroupMtu__DiscoveryDelete,

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
			"mtu__discovery": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group. Enable TCP path MTU discovery",
			},
		},
	}
}