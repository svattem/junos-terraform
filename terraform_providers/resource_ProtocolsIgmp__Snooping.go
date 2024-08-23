
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
type xmlProtocolsIgmp__Snooping struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_igmp__snooping  struct {
			XMLName xml.Name `xml:"igmp-snooping"`
		} `xml:"protocols>igmp-snooping"`
	} `xml:"groups"`

}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosProtocolsIgmp__SnoopingCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     

	config := xmlProtocolsIgmp__Snooping{}
	config.Groups.Name = id

    err = client.SendTransaction("", config, false)
    check(ctx, err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosProtocolsIgmp__SnoopingRead(ctx,d,m)
}

func junosProtocolsIgmp__SnoopingRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlProtocolsIgmp__Snooping{}

	err = client.MarshalGroup(id, config)
	check(ctx, err)
 
	return nil
}

func junosProtocolsIgmp__SnoopingUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     

	config := xmlProtocolsIgmp__Snooping{}
	config.Groups.Name = id

    err = client.SendTransaction(id, config, false)
    check(ctx, err)
    
	return junosProtocolsIgmp__SnoopingRead(ctx,d,m)
}

func junosProtocolsIgmp__SnoopingDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	_, err = client.DeleteConfig(id,false)
    check(ctx, err)

    d.SetId("")
    
	return nil
}

func junosProtocolsIgmp__Snooping() *schema.Resource {
	return &schema.Resource{
		CreateContext: junosProtocolsIgmp__SnoopingCreate,
		ReadContext: junosProtocolsIgmp__SnoopingRead,
		UpdateContext: junosProtocolsIgmp__SnoopingUpdate,
		DeleteContext: junosProtocolsIgmp__SnoopingDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
		},
	}
}