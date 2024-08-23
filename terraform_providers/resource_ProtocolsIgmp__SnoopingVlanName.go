
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
type xmlProtocolsIgmp__SnoopingVlanName struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_vlan  struct {
			XMLName xml.Name `xml:"vlan"`
			V_name  *string  `xml:"name,omitempty"`
		} `xml:"protocols>igmp-snooping>vlan"`
	} `xml:"groups"`

}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosProtocolsIgmp__SnoopingVlanNameCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)


	config := xmlProtocolsIgmp__SnoopingVlanName{}
	config.Groups.Name = id
	config.Groups.V_vlan.V_name = &V_name

    err = client.SendTransaction("", config, false)
    check(ctx, err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosProtocolsIgmp__SnoopingVlanNameRead(ctx,d,m)
}

func junosProtocolsIgmp__SnoopingVlanNameRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlProtocolsIgmp__SnoopingVlanName{}

	err = client.MarshalGroup(id, config)
	check(ctx, err)
 	if err :=d.Set("name", config.Groups.V_vlan.V_name);err != nil{
		return diag.FromErr(err)
	}

	return nil
}

func junosProtocolsIgmp__SnoopingVlanNameUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)


	config := xmlProtocolsIgmp__SnoopingVlanName{}
	config.Groups.Name = id
	config.Groups.V_vlan.V_name = &V_name

    err = client.SendTransaction(id, config, false)
    check(ctx, err)
    
	return junosProtocolsIgmp__SnoopingVlanNameRead(ctx,d,m)
}

func junosProtocolsIgmp__SnoopingVlanNameDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	_, err = client.DeleteConfig(id,false)
    check(ctx, err)

    d.SetId("")
    
	return nil
}

func junosProtocolsIgmp__SnoopingVlanName() *schema.Resource {
	return &schema.Resource{
		CreateContext: junosProtocolsIgmp__SnoopingVlanNameCreate,
		ReadContext: junosProtocolsIgmp__SnoopingVlanNameRead,
		UpdateContext: junosProtocolsIgmp__SnoopingVlanNameUpdate,
		DeleteContext: junosProtocolsIgmp__SnoopingVlanNameDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"name": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_vlan. VLAN name",
			},
		},
	}
}