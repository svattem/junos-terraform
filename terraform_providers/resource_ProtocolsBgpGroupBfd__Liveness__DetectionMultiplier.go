
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
type xmlProtocolsBgpGroupBfd__Liveness__DetectionMultiplier struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_group  struct {
			XMLName xml.Name `xml:"group"`
			V_name  *string  `xml:"name,omitempty"`
			V_bfd__liveness__detection  struct {
				XMLName xml.Name `xml:"bfd-liveness-detection"`
				V_multiplier  *string  `xml:"multiplier,omitempty"`
			} `xml:"bfd-liveness-detection"`
		} `xml:"protocols>bgp>group"`
	} `xml:"groups"`

}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosProtocolsBgpGroupBfd__Liveness__DetectionMultiplierCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_multiplier := d.Get("multiplier").(string)


	config := xmlProtocolsBgpGroupBfd__Liveness__DetectionMultiplier{}
	config.Groups.Name = id
	config.Groups.V_group.V_name = &V_name
	config.Groups.V_group.V_bfd__liveness__detection.V_multiplier = &V_multiplier

    err = client.SendTransaction("", config, false)
    check(ctx, err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosProtocolsBgpGroupBfd__Liveness__DetectionMultiplierRead(ctx,d,m)
}

func junosProtocolsBgpGroupBfd__Liveness__DetectionMultiplierRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlProtocolsBgpGroupBfd__Liveness__DetectionMultiplier{}

	err = client.MarshalGroup(id, config)
	check(ctx, err)
 	if err :=d.Set("name", config.Groups.V_group.V_name);err != nil{
		return diag.FromErr(err)
	}
	if err :=d.Set("multiplier", config.Groups.V_group.V_bfd__liveness__detection.V_multiplier);err != nil{
		return diag.FromErr(err)
	}

	return nil
}

func junosProtocolsBgpGroupBfd__Liveness__DetectionMultiplierUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_multiplier := d.Get("multiplier").(string)


	config := xmlProtocolsBgpGroupBfd__Liveness__DetectionMultiplier{}
	config.Groups.Name = id
	config.Groups.V_group.V_name = &V_name
	config.Groups.V_group.V_bfd__liveness__detection.V_multiplier = &V_multiplier

    err = client.SendTransaction(id, config, false)
    check(ctx, err)
    
	return junosProtocolsBgpGroupBfd__Liveness__DetectionMultiplierRead(ctx,d,m)
}

func junosProtocolsBgpGroupBfd__Liveness__DetectionMultiplierDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	_, err = client.DeleteConfig(id,false)
    check(ctx, err)

    d.SetId("")
    
	return nil
}

func junosProtocolsBgpGroupBfd__Liveness__DetectionMultiplier() *schema.Resource {
	return &schema.Resource{
		CreateContext: junosProtocolsBgpGroupBfd__Liveness__DetectionMultiplierCreate,
		ReadContext: junosProtocolsBgpGroupBfd__Liveness__DetectionMultiplierRead,
		UpdateContext: junosProtocolsBgpGroupBfd__Liveness__DetectionMultiplierUpdate,
		DeleteContext: junosProtocolsBgpGroupBfd__Liveness__DetectionMultiplierDelete,

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
			"multiplier": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_bfd__liveness__detection. Detection time multiplier",
			},
		},
	}
}