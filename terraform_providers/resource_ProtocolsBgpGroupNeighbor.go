
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
type xmlProtocolsBgpGroupNeighbor struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_group  struct {
			XMLName xml.Name `xml:"group"`
			V_name  *string  `xml:"name,omitempty"`
			V_neighbor  struct {
				XMLName xml.Name `xml:"neighbor"`
			} `xml:"neighbor"`
		} `xml:"protocols>bgp>group"`
	} `xml:"groups"`

}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosProtocolsBgpGroupNeighborCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)


	config := xmlProtocolsBgpGroupNeighbor{}
	config.Groups.Name = id
	config.Groups.V_group.V_name = &V_name

    err = client.SendTransaction("", config, false)
    check(ctx, err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosProtocolsBgpGroupNeighborRead(ctx,d,m)
}

func junosProtocolsBgpGroupNeighborRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlProtocolsBgpGroupNeighbor{}

	err = client.MarshalGroup(id, config)
	check(ctx, err)
 	if err :=d.Set("name", config.Groups.V_group.V_name);err != nil{
		return diag.FromErr(err)
	}

	return nil
}

func junosProtocolsBgpGroupNeighborUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)


	config := xmlProtocolsBgpGroupNeighbor{}
	config.Groups.Name = id
	config.Groups.V_group.V_name = &V_name

    err = client.SendTransaction(id, config, false)
    check(ctx, err)
    
	return junosProtocolsBgpGroupNeighborRead(ctx,d,m)
}

func junosProtocolsBgpGroupNeighborDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	_, err = client.DeleteConfig(id,false)
    check(ctx, err)

    d.SetId("")
    
	return nil
}

func junosProtocolsBgpGroupNeighbor() *schema.Resource {
	return &schema.Resource{
		CreateContext: junosProtocolsBgpGroupNeighborCreate,
		ReadContext: junosProtocolsBgpGroupNeighborRead,
		UpdateContext: junosProtocolsBgpGroupNeighborUpdate,
		DeleteContext: junosProtocolsBgpGroupNeighborDelete,

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
		},
	}
}