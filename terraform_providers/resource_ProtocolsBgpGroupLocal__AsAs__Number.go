
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
type xmlProtocolsBgpGroupLocal__AsAs__Number struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_group  struct {
			XMLName xml.Name `xml:"group"`
			V_name  *string  `xml:"name,omitempty"`
			V_local__as  struct {
				XMLName xml.Name `xml:"local-as"`
				V_as__number  *string  `xml:"as-number,omitempty"`
			} `xml:"local-as"`
		} `xml:"protocols>bgp>group"`
	} `xml:"groups"`

}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosProtocolsBgpGroupLocal__AsAs__NumberCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_as__number := d.Get("as__number").(string)


	config := xmlProtocolsBgpGroupLocal__AsAs__Number{}
	config.Groups.Name = id
	config.Groups.V_group.V_name = &V_name
	config.Groups.V_group.V_local__as.V_as__number = &V_as__number

    err = client.SendTransaction("", config, false)
    check(ctx, err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosProtocolsBgpGroupLocal__AsAs__NumberRead(ctx,d,m)
}

func junosProtocolsBgpGroupLocal__AsAs__NumberRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlProtocolsBgpGroupLocal__AsAs__Number{}

	err = client.MarshalGroup(id, config)
	check(ctx, err)
 	if err :=d.Set("name", config.Groups.V_group.V_name);err != nil{
		return diag.FromErr(err)
	}
	if err :=d.Set("as__number", config.Groups.V_group.V_local__as.V_as__number);err != nil{
		return diag.FromErr(err)
	}

	return nil
}

func junosProtocolsBgpGroupLocal__AsAs__NumberUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_as__number := d.Get("as__number").(string)


	config := xmlProtocolsBgpGroupLocal__AsAs__Number{}
	config.Groups.Name = id
	config.Groups.V_group.V_name = &V_name
	config.Groups.V_group.V_local__as.V_as__number = &V_as__number

    err = client.SendTransaction(id, config, false)
    check(ctx, err)
    
	return junosProtocolsBgpGroupLocal__AsAs__NumberRead(ctx,d,m)
}

func junosProtocolsBgpGroupLocal__AsAs__NumberDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	_, err = client.DeleteConfig(id,false)
    check(ctx, err)

    d.SetId("")
    
	return nil
}

func junosProtocolsBgpGroupLocal__AsAs__Number() *schema.Resource {
	return &schema.Resource{
		CreateContext: junosProtocolsBgpGroupLocal__AsAs__NumberCreate,
		ReadContext: junosProtocolsBgpGroupLocal__AsAs__NumberRead,
		UpdateContext: junosProtocolsBgpGroupLocal__AsAs__NumberUpdate,
		DeleteContext: junosProtocolsBgpGroupLocal__AsAs__NumberDelete,

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
			"as__number": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_local__as.  Autonomous system number in plain number or 'higher 16bits'.'Lower 16 bits' (asdot notation) format",
			},
		},
	}
}