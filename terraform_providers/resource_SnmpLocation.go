
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
type xmlSnmpLocation struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_  struct {
			XMLName xml.Name `xml:""`
			V_location  *string  `xml:"location,omitempty"`
		} `xml:"snmp"`
	} `xml:"groups"`

}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSnmpLocationCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_location := d.Get("location").(string)


	config := xmlSnmpLocation{}
	config.Groups.Name = id
	config.Groups.V_.V_location = &V_location

    err = client.SendTransaction("", config, false)
    check(ctx, err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSnmpLocationRead(ctx,d,m)
}

func junosSnmpLocationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSnmpLocation{}

	err = client.MarshalGroup(id, config)
	check(ctx, err)
 	if err :=d.Set("location", config.Groups.V_.V_location);err != nil{
		return diag.FromErr(err)
	}

	return nil
}

func junosSnmpLocationUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_location := d.Get("location").(string)


	config := xmlSnmpLocation{}
	config.Groups.Name = id
	config.Groups.V_.V_location = &V_location

    err = client.SendTransaction(id, config, false)
    check(ctx, err)
    
	return junosSnmpLocationRead(ctx,d,m)
}

func junosSnmpLocationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	_, err = client.DeleteConfig(id,false)
    check(ctx, err)

    d.SetId("")
    
	return nil
}

func junosSnmpLocation() *schema.Resource {
	return &schema.Resource{
		CreateContext: junosSnmpLocationCreate,
		ReadContext: junosSnmpLocationRead,
		UpdateContext: junosSnmpLocationUpdate,
		DeleteContext: junosSnmpLocationDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"location": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_. Physical location of system",
			},
		},
	}
}