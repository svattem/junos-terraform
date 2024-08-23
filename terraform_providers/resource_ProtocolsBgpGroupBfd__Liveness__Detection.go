
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
type xmlProtocolsBgpGroupBfd__Liveness__Detection struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_group  struct {
			XMLName xml.Name `xml:"group"`
			V_name  *string  `xml:"name,omitempty"`
			V_bfd__liveness__detection  struct {
				XMLName xml.Name `xml:"bfd-liveness-detection"`
			} `xml:"bfd-liveness-detection"`
		} `xml:"protocols>bgp>group"`
	} `xml:"groups"`

}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosProtocolsBgpGroupBfd__Liveness__DetectionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)


	config := xmlProtocolsBgpGroupBfd__Liveness__Detection{}
	config.Groups.Name = id
	config.Groups.V_group.V_name = &V_name

    err = client.SendTransaction("", config, false)
    check(ctx, err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosProtocolsBgpGroupBfd__Liveness__DetectionRead(ctx,d,m)
}

func junosProtocolsBgpGroupBfd__Liveness__DetectionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlProtocolsBgpGroupBfd__Liveness__Detection{}

	err = client.MarshalGroup(id, config)
	check(ctx, err)
 	if err :=d.Set("name", config.Groups.V_group.V_name);err != nil{
		return diag.FromErr(err)
	}

	return nil
}

func junosProtocolsBgpGroupBfd__Liveness__DetectionUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)


	config := xmlProtocolsBgpGroupBfd__Liveness__Detection{}
	config.Groups.Name = id
	config.Groups.V_group.V_name = &V_name

    err = client.SendTransaction(id, config, false)
    check(ctx, err)
    
	return junosProtocolsBgpGroupBfd__Liveness__DetectionRead(ctx,d,m)
}

func junosProtocolsBgpGroupBfd__Liveness__DetectionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	_, err = client.DeleteConfig(id,false)
    check(ctx, err)

    d.SetId("")
    
	return nil
}

func junosProtocolsBgpGroupBfd__Liveness__Detection() *schema.Resource {
	return &schema.Resource{
		CreateContext: junosProtocolsBgpGroupBfd__Liveness__DetectionCreate,
		ReadContext: junosProtocolsBgpGroupBfd__Liveness__DetectionRead,
		UpdateContext: junosProtocolsBgpGroupBfd__Liveness__DetectionUpdate,
		DeleteContext: junosProtocolsBgpGroupBfd__Liveness__DetectionDelete,

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