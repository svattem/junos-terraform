
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
type xmlForwarding__OptionsStorm__Control__ProfilesName struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_storm__control__profiles  struct {
			XMLName xml.Name `xml:"storm-control-profiles"`
			V_name  *string  `xml:"name,omitempty"`
		} `xml:"forwarding-options>storm-control-profiles"`
	} `xml:"groups"`

}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosForwarding__OptionsStorm__Control__ProfilesNameCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)


	config := xmlForwarding__OptionsStorm__Control__ProfilesName{}
	config.Groups.Name = id
	config.Groups.V_storm__control__profiles.V_name = &V_name

    err = client.SendTransaction("", config, false)
    check(ctx, err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosForwarding__OptionsStorm__Control__ProfilesNameRead(ctx,d,m)
}

func junosForwarding__OptionsStorm__Control__ProfilesNameRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlForwarding__OptionsStorm__Control__ProfilesName{}

	err = client.MarshalGroup(id, config)
	check(ctx, err)
 	if err :=d.Set("name", config.Groups.V_storm__control__profiles.V_name);err != nil{
		return diag.FromErr(err)
	}

	return nil
}

func junosForwarding__OptionsStorm__Control__ProfilesNameUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)


	config := xmlForwarding__OptionsStorm__Control__ProfilesName{}
	config.Groups.Name = id
	config.Groups.V_storm__control__profiles.V_name = &V_name

    err = client.SendTransaction(id, config, false)
    check(ctx, err)
    
	return junosForwarding__OptionsStorm__Control__ProfilesNameRead(ctx,d,m)
}

func junosForwarding__OptionsStorm__Control__ProfilesNameDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	_, err = client.DeleteConfig(id,false)
    check(ctx, err)

    d.SetId("")
    
	return nil
}

func junosForwarding__OptionsStorm__Control__ProfilesName() *schema.Resource {
	return &schema.Resource{
		CreateContext: junosForwarding__OptionsStorm__Control__ProfilesNameCreate,
		ReadContext: junosForwarding__OptionsStorm__Control__ProfilesNameRead,
		UpdateContext: junosForwarding__OptionsStorm__Control__ProfilesNameUpdate,
		DeleteContext: junosForwarding__OptionsStorm__Control__ProfilesNameDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"name": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_storm__control__profiles. Storm control profile name",
			},
		},
	}
}