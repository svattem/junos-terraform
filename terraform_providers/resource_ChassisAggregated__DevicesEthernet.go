
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
type xmlChassisAggregated__DevicesEthernet struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_ethernet  struct {
			XMLName xml.Name `xml:"ethernet"`
		} `xml:"chassis>aggregated-devices>ethernet"`
	} `xml:"groups"`

}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosChassisAggregated__DevicesEthernetCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     

	config := xmlChassisAggregated__DevicesEthernet{}
	config.Groups.Name = id

    err = client.SendTransaction("", config, false)
    check(ctx, err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosChassisAggregated__DevicesEthernetRead(ctx,d,m)
}

func junosChassisAggregated__DevicesEthernetRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlChassisAggregated__DevicesEthernet{}

	err = client.MarshalGroup(id, config)
	check(ctx, err)
 
	return nil
}

func junosChassisAggregated__DevicesEthernetUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     

	config := xmlChassisAggregated__DevicesEthernet{}
	config.Groups.Name = id

    err = client.SendTransaction(id, config, false)
    check(ctx, err)
    
	return junosChassisAggregated__DevicesEthernetRead(ctx,d,m)
}

func junosChassisAggregated__DevicesEthernetDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	_, err = client.DeleteConfig(id,false)
    check(ctx, err)

    d.SetId("")
    
	return nil
}

func junosChassisAggregated__DevicesEthernet() *schema.Resource {
	return &schema.Resource{
		CreateContext: junosChassisAggregated__DevicesEthernetCreate,
		ReadContext: junosChassisAggregated__DevicesEthernetRead,
		UpdateContext: junosChassisAggregated__DevicesEthernetUpdate,
		DeleteContext: junosChassisAggregated__DevicesEthernetDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
		},
	}
}