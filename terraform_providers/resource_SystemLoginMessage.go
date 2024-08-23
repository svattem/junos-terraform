
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
type xmlSystemLoginMessage struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_login  struct {
			XMLName xml.Name `xml:"login"`
			V_message  *string  `xml:"message,omitempty"`
		} `xml:"system>login"`
	} `xml:"groups"`

}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSystemLoginMessageCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_message := d.Get("message").(string)


	config := xmlSystemLoginMessage{}
	config.Groups.Name = id
	config.Groups.V_login.V_message = &V_message

    err = client.SendTransaction("", config, false)
    check(ctx, err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSystemLoginMessageRead(ctx,d,m)
}

func junosSystemLoginMessageRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSystemLoginMessage{}

	err = client.MarshalGroup(id, config)
	check(ctx, err)
 	if err :=d.Set("message", config.Groups.V_login.V_message);err != nil{
		return diag.FromErr(err)
	}

	return nil
}

func junosSystemLoginMessageUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_message := d.Get("message").(string)


	config := xmlSystemLoginMessage{}
	config.Groups.Name = id
	config.Groups.V_login.V_message = &V_message

    err = client.SendTransaction(id, config, false)
    check(ctx, err)
    
	return junosSystemLoginMessageRead(ctx,d,m)
}

func junosSystemLoginMessageDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	_, err = client.DeleteConfig(id,false)
    check(ctx, err)

    d.SetId("")
    
	return nil
}

func junosSystemLoginMessage() *schema.Resource {
	return &schema.Resource{
		CreateContext: junosSystemLoginMessageCreate,
		ReadContext: junosSystemLoginMessageRead,
		UpdateContext: junosSystemLoginMessageUpdate,
		DeleteContext: junosSystemLoginMessageDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"message": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_login. System login message",
			},
		},
	}
}