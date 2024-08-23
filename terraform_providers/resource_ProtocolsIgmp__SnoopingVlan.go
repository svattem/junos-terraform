
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
type xmlProtocolsIgmp__SnoopingVlan struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_vlan  struct {
			XMLName xml.Name `xml:"vlan"`
		} `xml:"protocols>igmp-snooping>vlan"`
	} `xml:"groups"`

}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosProtocolsIgmp__SnoopingVlanCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     

	config := xmlProtocolsIgmp__SnoopingVlan{}
	config.Groups.Name = id

    err = client.SendTransaction("", config, false)
    check(ctx, err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosProtocolsIgmp__SnoopingVlanRead(ctx,d,m)
}

func junosProtocolsIgmp__SnoopingVlanRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlProtocolsIgmp__SnoopingVlan{}

	err = client.MarshalGroup(id, config)
	check(ctx, err)
 
	return nil
}

func junosProtocolsIgmp__SnoopingVlanUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     

	config := xmlProtocolsIgmp__SnoopingVlan{}
	config.Groups.Name = id

    err = client.SendTransaction(id, config, false)
    check(ctx, err)
    
	return junosProtocolsIgmp__SnoopingVlanRead(ctx,d,m)
}

func junosProtocolsIgmp__SnoopingVlanDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	_, err = client.DeleteConfig(id,false)
    check(ctx, err)

    d.SetId("")
    
	return nil
}

func junosProtocolsIgmp__SnoopingVlan() *schema.Resource {
	return &schema.Resource{
		CreateContext: junosProtocolsIgmp__SnoopingVlanCreate,
		ReadContext: junosProtocolsIgmp__SnoopingVlanRead,
		UpdateContext: junosProtocolsIgmp__SnoopingVlanUpdate,
		DeleteContext: junosProtocolsIgmp__SnoopingVlanDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
		},
	}
}