
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
type xmlProtocolsBgpGroupNeighborDescription struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_group  struct {
			XMLName xml.Name `xml:"group"`
			V_name  *string  `xml:"name,omitempty"`
			V_neighbor  struct {
				XMLName xml.Name `xml:"neighbor"`
				V_name__1  *string  `xml:"name,omitempty"`
				V_description  *string  `xml:"description,omitempty"`
			} `xml:"neighbor"`
		} `xml:"protocols>bgp>group"`
	} `xml:"groups"`

}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosProtocolsBgpGroupNeighborDescriptionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_name__1 := d.Get("name__1").(string)
	V_description := d.Get("description").(string)


	config := xmlProtocolsBgpGroupNeighborDescription{}
	config.Groups.Name = id
	config.Groups.V_group.V_name = &V_name
	config.Groups.V_group.V_neighbor.V_name__1 = &V_name__1
	config.Groups.V_group.V_neighbor.V_description = &V_description

    err = client.SendTransaction("", config, false)
    check(ctx, err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosProtocolsBgpGroupNeighborDescriptionRead(ctx,d,m)
}

func junosProtocolsBgpGroupNeighborDescriptionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlProtocolsBgpGroupNeighborDescription{}

	err = client.MarshalGroup(id, config)
	check(ctx, err)
 	if err :=d.Set("name", config.Groups.V_group.V_name);err != nil{
		return diag.FromErr(err)
	}
	if err :=d.Set("name__1", config.Groups.V_group.V_neighbor.V_name__1);err != nil{
		return diag.FromErr(err)
	}
	if err :=d.Set("description", config.Groups.V_group.V_neighbor.V_description);err != nil{
		return diag.FromErr(err)
	}

	return nil
}

func junosProtocolsBgpGroupNeighborDescriptionUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_name__1 := d.Get("name__1").(string)
	V_description := d.Get("description").(string)


	config := xmlProtocolsBgpGroupNeighborDescription{}
	config.Groups.Name = id
	config.Groups.V_group.V_name = &V_name
	config.Groups.V_group.V_neighbor.V_name__1 = &V_name__1
	config.Groups.V_group.V_neighbor.V_description = &V_description

    err = client.SendTransaction(id, config, false)
    check(ctx, err)
    
	return junosProtocolsBgpGroupNeighborDescriptionRead(ctx,d,m)
}

func junosProtocolsBgpGroupNeighborDescriptionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	_, err = client.DeleteConfig(id,false)
    check(ctx, err)

    d.SetId("")
    
	return nil
}

func junosProtocolsBgpGroupNeighborDescription() *schema.Resource {
	return &schema.Resource{
		CreateContext: junosProtocolsBgpGroupNeighborDescriptionCreate,
		ReadContext: junosProtocolsBgpGroupNeighborDescriptionRead,
		UpdateContext: junosProtocolsBgpGroupNeighborDescriptionUpdate,
		DeleteContext: junosProtocolsBgpGroupNeighborDescriptionDelete,

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
			"name__1": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_neighbor",
			},
			"description": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_neighbor. Text description",
			},
		},
	}
}