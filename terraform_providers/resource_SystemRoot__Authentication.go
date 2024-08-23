
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
type xmlSystemRoot__Authentication struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_root__authentication  struct {
			XMLName xml.Name `xml:"root-authentication"`
		} `xml:"system>root-authentication"`
	} `xml:"groups"`

}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSystemRoot__AuthenticationCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     

	config := xmlSystemRoot__Authentication{}
	config.Groups.Name = id

    err = client.SendTransaction("", config, false)
    check(ctx, err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSystemRoot__AuthenticationRead(ctx,d,m)
}

func junosSystemRoot__AuthenticationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSystemRoot__Authentication{}

	err = client.MarshalGroup(id, config)
	check(ctx, err)
 
	return nil
}

func junosSystemRoot__AuthenticationUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     

	config := xmlSystemRoot__Authentication{}
	config.Groups.Name = id

    err = client.SendTransaction(id, config, false)
    check(ctx, err)
    
	return junosSystemRoot__AuthenticationRead(ctx,d,m)
}

func junosSystemRoot__AuthenticationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	_, err = client.DeleteConfig(id,false)
    check(ctx, err)

    d.SetId("")
    
	return nil
}

func junosSystemRoot__Authentication() *schema.Resource {
	return &schema.Resource{
		CreateContext: junosSystemRoot__AuthenticationCreate,
		ReadContext: junosSystemRoot__AuthenticationRead,
		UpdateContext: junosSystemRoot__AuthenticationUpdate,
		DeleteContext: junosSystemRoot__AuthenticationDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
		},
	}
}