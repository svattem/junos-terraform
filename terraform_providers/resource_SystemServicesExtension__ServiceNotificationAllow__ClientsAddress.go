
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
type xmlSystemServicesExtension__ServiceNotificationAllow__ClientsAddress struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_allow__clients  struct {
			XMLName xml.Name `xml:"allow-clients"`
			V_address  *string  `xml:"address,omitempty"`
		} `xml:"system>services>extension-service>notification>allow-clients"`
	} `xml:"groups"`

}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSystemServicesExtension__ServiceNotificationAllow__ClientsAddressCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_address := d.Get("address").(string)


	config := xmlSystemServicesExtension__ServiceNotificationAllow__ClientsAddress{}
	config.Groups.Name = id
	config.Groups.V_allow__clients.V_address = &V_address

    err = client.SendTransaction("", config, false)
    check(ctx, err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSystemServicesExtension__ServiceNotificationAllow__ClientsAddressRead(ctx,d,m)
}

func junosSystemServicesExtension__ServiceNotificationAllow__ClientsAddressRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSystemServicesExtension__ServiceNotificationAllow__ClientsAddress{}

	err = client.MarshalGroup(id, config)
	check(ctx, err)
 	if err :=d.Set("address", config.Groups.V_allow__clients.V_address);err != nil{
		return diag.FromErr(err)
	}

	return nil
}

func junosSystemServicesExtension__ServiceNotificationAllow__ClientsAddressUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_address := d.Get("address").(string)


	config := xmlSystemServicesExtension__ServiceNotificationAllow__ClientsAddress{}
	config.Groups.Name = id
	config.Groups.V_allow__clients.V_address = &V_address

    err = client.SendTransaction(id, config, false)
    check(ctx, err)
    
	return junosSystemServicesExtension__ServiceNotificationAllow__ClientsAddressRead(ctx,d,m)
}

func junosSystemServicesExtension__ServiceNotificationAllow__ClientsAddressDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	_, err = client.DeleteConfig(id,false)
    check(ctx, err)

    d.SetId("")
    
	return nil
}

func junosSystemServicesExtension__ServiceNotificationAllow__ClientsAddress() *schema.Resource {
	return &schema.Resource{
		CreateContext: junosSystemServicesExtension__ServiceNotificationAllow__ClientsAddressCreate,
		ReadContext: junosSystemServicesExtension__ServiceNotificationAllow__ClientsAddressRead,
		UpdateContext: junosSystemServicesExtension__ServiceNotificationAllow__ClientsAddressUpdate,
		DeleteContext: junosSystemServicesExtension__ServiceNotificationAllow__ClientsAddressDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"address": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_allow__clients. IPv4/IPv6 addresses, prefix length optional, or hostnames",
			},
		},
	}
}