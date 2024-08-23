
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
type xmlSnmpContact struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_  struct {
			XMLName xml.Name `xml:""`
			V_contact  *string  `xml:"contact,omitempty"`
		} `xml:"snmp"`
	} `xml:"groups"`

}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSnmpContactCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_contact := d.Get("contact").(string)


	config := xmlSnmpContact{}
	config.Groups.Name = id
	config.Groups.V_.V_contact = &V_contact

    err = client.SendTransaction("", config, false)
    check(ctx, err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSnmpContactRead(ctx,d,m)
}

func junosSnmpContactRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSnmpContact{}

	err = client.MarshalGroup(id, config)
	check(ctx, err)
 	if err :=d.Set("contact", config.Groups.V_.V_contact);err != nil{
		return diag.FromErr(err)
	}

	return nil
}

func junosSnmpContactUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_contact := d.Get("contact").(string)


	config := xmlSnmpContact{}
	config.Groups.Name = id
	config.Groups.V_.V_contact = &V_contact

    err = client.SendTransaction(id, config, false)
    check(ctx, err)
    
	return junosSnmpContactRead(ctx,d,m)
}

func junosSnmpContactDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	_, err = client.DeleteConfig(id,false)
    check(ctx, err)

    d.SetId("")
    
	return nil
}

func junosSnmpContact() *schema.Resource {
	return &schema.Resource{
		CreateContext: junosSnmpContactCreate,
		ReadContext: junosSnmpContactRead,
		UpdateContext: junosSnmpContactUpdate,
		DeleteContext: junosSnmpContactDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"contact": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_. Contact information for administrator",
			},
		},
	}
}