
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
type xmlForwarding__OptionsStorm__Control__Profiles struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_storm__control__profiles  struct {
			XMLName xml.Name `xml:"storm-control-profiles"`
		} `xml:"forwarding-options>storm-control-profiles"`
	} `xml:"groups"`

}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosForwarding__OptionsStorm__Control__ProfilesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     

	config := xmlForwarding__OptionsStorm__Control__Profiles{}
	config.Groups.Name = id

    err = client.SendTransaction("", config, false)
    check(ctx, err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosForwarding__OptionsStorm__Control__ProfilesRead(ctx,d,m)
}

func junosForwarding__OptionsStorm__Control__ProfilesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlForwarding__OptionsStorm__Control__Profiles{}

	err = client.MarshalGroup(id, config)
	check(ctx, err)
 
	return nil
}

func junosForwarding__OptionsStorm__Control__ProfilesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     

	config := xmlForwarding__OptionsStorm__Control__Profiles{}
	config.Groups.Name = id

    err = client.SendTransaction(id, config, false)
    check(ctx, err)
    
	return junosForwarding__OptionsStorm__Control__ProfilesRead(ctx,d,m)
}

func junosForwarding__OptionsStorm__Control__ProfilesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	_, err = client.DeleteConfig(id,false)
    check(ctx, err)

    d.SetId("")
    
	return nil
}

func junosForwarding__OptionsStorm__Control__Profiles() *schema.Resource {
	return &schema.Resource{
		CreateContext: junosForwarding__OptionsStorm__Control__ProfilesCreate,
		ReadContext: junosForwarding__OptionsStorm__Control__ProfilesRead,
		UpdateContext: junosForwarding__OptionsStorm__Control__ProfilesUpdate,
		DeleteContext: junosForwarding__OptionsStorm__Control__ProfilesDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
		},
	}
}