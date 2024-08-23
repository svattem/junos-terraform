
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
type xmlProtocolsBgpGroupMultipathMultiple__As struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_group  struct {
			XMLName xml.Name `xml:"group"`
			V_name  *string  `xml:"name,omitempty"`
			V_multipath  struct {
				XMLName xml.Name `xml:"multipath"`
				V_multiple__as  *string  `xml:"multiple-as,omitempty"`
			} `xml:"multipath"`
		} `xml:"protocols>bgp>group"`
	} `xml:"groups"`

}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosProtocolsBgpGroupMultipathMultiple__AsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_multiple__as := d.Get("multiple__as").(string)


	config := xmlProtocolsBgpGroupMultipathMultiple__As{}
	config.Groups.Name = id
	config.Groups.V_group.V_name = &V_name
	config.Groups.V_group.V_multipath.V_multiple__as = &V_multiple__as

    err = client.SendTransaction("", config, false)
    check(ctx, err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosProtocolsBgpGroupMultipathMultiple__AsRead(ctx,d,m)
}

func junosProtocolsBgpGroupMultipathMultiple__AsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlProtocolsBgpGroupMultipathMultiple__As{}

	err = client.MarshalGroup(id, config)
	check(ctx, err)
 	if err :=d.Set("name", config.Groups.V_group.V_name);err != nil{
		return diag.FromErr(err)
	}
	if err :=d.Set("multiple__as", config.Groups.V_group.V_multipath.V_multiple__as);err != nil{
		return diag.FromErr(err)
	}

	return nil
}

func junosProtocolsBgpGroupMultipathMultiple__AsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_multiple__as := d.Get("multiple__as").(string)


	config := xmlProtocolsBgpGroupMultipathMultiple__As{}
	config.Groups.Name = id
	config.Groups.V_group.V_name = &V_name
	config.Groups.V_group.V_multipath.V_multiple__as = &V_multiple__as

    err = client.SendTransaction(id, config, false)
    check(ctx, err)
    
	return junosProtocolsBgpGroupMultipathMultiple__AsRead(ctx,d,m)
}

func junosProtocolsBgpGroupMultipathMultiple__AsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	_, err = client.DeleteConfig(id,false)
    check(ctx, err)

    d.SetId("")
    
	return nil
}

func junosProtocolsBgpGroupMultipathMultiple__As() *schema.Resource {
	return &schema.Resource{
		CreateContext: junosProtocolsBgpGroupMultipathMultiple__AsCreate,
		ReadContext: junosProtocolsBgpGroupMultipathMultiple__AsRead,
		UpdateContext: junosProtocolsBgpGroupMultipathMultiple__AsUpdate,
		DeleteContext: junosProtocolsBgpGroupMultipathMultiple__AsDelete,

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
			"multiple__as": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_multipath. Use paths received from different ASs",
			},
		},
	}
}