
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
type xmlChassisAggregated__Devices struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_aggregated__devices  struct {
			XMLName xml.Name `xml:"aggregated-devices"`
		} `xml:"chassis>aggregated-devices"`
	} `xml:"groups"`

}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosChassisAggregated__DevicesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     

	config := xmlChassisAggregated__Devices{}
	config.Groups.Name = id

    err = client.SendTransaction("", config, false)
    check(ctx, err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosChassisAggregated__DevicesRead(ctx,d,m)
}

func junosChassisAggregated__DevicesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlChassisAggregated__Devices{}

	err = client.MarshalGroup(id, config)
	check(ctx, err)
 
	return nil
}

func junosChassisAggregated__DevicesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     

	config := xmlChassisAggregated__Devices{}
	config.Groups.Name = id

    err = client.SendTransaction(id, config, false)
    check(ctx, err)
    
	return junosChassisAggregated__DevicesRead(ctx,d,m)
}

func junosChassisAggregated__DevicesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	_, err = client.DeleteConfig(id,false)
    check(ctx, err)

    d.SetId("")
    
	return nil
}

func junosChassisAggregated__Devices() *schema.Resource {
	return &schema.Resource{
		CreateContext: junosChassisAggregated__DevicesCreate,
		ReadContext: junosChassisAggregated__DevicesRead,
		UpdateContext: junosChassisAggregated__DevicesUpdate,
		DeleteContext: junosChassisAggregated__DevicesDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
		},
	}
}