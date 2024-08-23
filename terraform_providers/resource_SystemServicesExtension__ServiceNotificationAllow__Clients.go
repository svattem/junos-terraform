
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
type xmlSystemServicesExtension__ServiceNotificationAllow__Clients struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_allow__clients  struct {
			XMLName xml.Name `xml:"allow-clients"`
		} `xml:"system>services>extension-service>notification>allow-clients"`
	} `xml:"groups"`

}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSystemServicesExtension__ServiceNotificationAllow__ClientsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     

	config := xmlSystemServicesExtension__ServiceNotificationAllow__Clients{}
	config.Groups.Name = id

    err = client.SendTransaction("", config, false)
    check(ctx, err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSystemServicesExtension__ServiceNotificationAllow__ClientsRead(ctx,d,m)
}

func junosSystemServicesExtension__ServiceNotificationAllow__ClientsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSystemServicesExtension__ServiceNotificationAllow__Clients{}

	err = client.MarshalGroup(id, config)
	check(ctx, err)
 
	return nil
}

func junosSystemServicesExtension__ServiceNotificationAllow__ClientsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     

	config := xmlSystemServicesExtension__ServiceNotificationAllow__Clients{}
	config.Groups.Name = id

    err = client.SendTransaction(id, config, false)
    check(ctx, err)
    
	return junosSystemServicesExtension__ServiceNotificationAllow__ClientsRead(ctx,d,m)
}

func junosSystemServicesExtension__ServiceNotificationAllow__ClientsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	_, err = client.DeleteConfig(id,false)
    check(ctx, err)

    d.SetId("")
    
	return nil
}

func junosSystemServicesExtension__ServiceNotificationAllow__Clients() *schema.Resource {
	return &schema.Resource{
		CreateContext: junosSystemServicesExtension__ServiceNotificationAllow__ClientsCreate,
		ReadContext: junosSystemServicesExtension__ServiceNotificationAllow__ClientsRead,
		UpdateContext: junosSystemServicesExtension__ServiceNotificationAllow__ClientsUpdate,
		DeleteContext: junosSystemServicesExtension__ServiceNotificationAllow__ClientsDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
		},
	}
}