
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
type xmlSystemServicesExtension__ServiceNotification struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_notification  struct {
			XMLName xml.Name `xml:"notification"`
		} `xml:"system>services>extension-service>notification"`
	} `xml:"groups"`

}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSystemServicesExtension__ServiceNotificationCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     

	config := xmlSystemServicesExtension__ServiceNotification{}
	config.Groups.Name = id

    err = client.SendTransaction("", config, false)
    check(ctx, err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSystemServicesExtension__ServiceNotificationRead(ctx,d,m)
}

func junosSystemServicesExtension__ServiceNotificationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSystemServicesExtension__ServiceNotification{}

	err = client.MarshalGroup(id, config)
	check(ctx, err)
 
	return nil
}

func junosSystemServicesExtension__ServiceNotificationUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     

	config := xmlSystemServicesExtension__ServiceNotification{}
	config.Groups.Name = id

    err = client.SendTransaction(id, config, false)
    check(ctx, err)
    
	return junosSystemServicesExtension__ServiceNotificationRead(ctx,d,m)
}

func junosSystemServicesExtension__ServiceNotificationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	_, err = client.DeleteConfig(id,false)
    check(ctx, err)

    d.SetId("")
    
	return nil
}

func junosSystemServicesExtension__ServiceNotification() *schema.Resource {
	return &schema.Resource{
		CreateContext: junosSystemServicesExtension__ServiceNotificationCreate,
		ReadContext: junosSystemServicesExtension__ServiceNotificationRead,
		UpdateContext: junosSystemServicesExtension__ServiceNotificationUpdate,
		DeleteContext: junosSystemServicesExtension__ServiceNotificationDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
		},
	}
}