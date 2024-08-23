
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
type xmlProtocolsBgpGroupVpn__Apply__Export struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_group  struct {
			XMLName xml.Name `xml:"group"`
			V_name  *string  `xml:"name,omitempty"`
			V_vpn__apply__export  *string  `xml:"vpn-apply-export,omitempty"`
		} `xml:"protocols>bgp>group"`
	} `xml:"groups"`

}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosProtocolsBgpGroupVpn__Apply__ExportCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_vpn__apply__export := d.Get("vpn__apply__export").(string)


	config := xmlProtocolsBgpGroupVpn__Apply__Export{}
	config.Groups.Name = id
	config.Groups.V_group.V_name = &V_name
	config.Groups.V_group.V_vpn__apply__export = &V_vpn__apply__export

    err = client.SendTransaction("", config, false)
    check(ctx, err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosProtocolsBgpGroupVpn__Apply__ExportRead(ctx,d,m)
}

func junosProtocolsBgpGroupVpn__Apply__ExportRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlProtocolsBgpGroupVpn__Apply__Export{}

	err = client.MarshalGroup(id, config)
	check(ctx, err)
 	if err :=d.Set("name", config.Groups.V_group.V_name);err != nil{
		return diag.FromErr(err)
	}
	if err :=d.Set("vpn__apply__export", config.Groups.V_group.V_vpn__apply__export);err != nil{
		return diag.FromErr(err)
	}

	return nil
}

func junosProtocolsBgpGroupVpn__Apply__ExportUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_vpn__apply__export := d.Get("vpn__apply__export").(string)


	config := xmlProtocolsBgpGroupVpn__Apply__Export{}
	config.Groups.Name = id
	config.Groups.V_group.V_name = &V_name
	config.Groups.V_group.V_vpn__apply__export = &V_vpn__apply__export

    err = client.SendTransaction(id, config, false)
    check(ctx, err)
    
	return junosProtocolsBgpGroupVpn__Apply__ExportRead(ctx,d,m)
}

func junosProtocolsBgpGroupVpn__Apply__ExportDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	_, err = client.DeleteConfig(id,false)
    check(ctx, err)

    d.SetId("")
    
	return nil
}

func junosProtocolsBgpGroupVpn__Apply__Export() *schema.Resource {
	return &schema.Resource{
		CreateContext: junosProtocolsBgpGroupVpn__Apply__ExportCreate,
		ReadContext: junosProtocolsBgpGroupVpn__Apply__ExportRead,
		UpdateContext: junosProtocolsBgpGroupVpn__Apply__ExportUpdate,
		DeleteContext: junosProtocolsBgpGroupVpn__Apply__ExportDelete,

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
			"vpn__apply__export": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group. Apply BGP export policy when exporting VPN routes",
			},
		},
	}
}