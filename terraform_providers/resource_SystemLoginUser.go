
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
type xmlSystemLoginUser struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_user  struct {
			XMLName xml.Name `xml:"user"`
		} `xml:"system>login>user"`
	} `xml:"groups"`

}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSystemLoginUserCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     

	config := xmlSystemLoginUser{}
	config.Groups.Name = id

    err = client.SendTransaction("", config, false)
    check(ctx, err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSystemLoginUserRead(ctx,d,m)
}

func junosSystemLoginUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSystemLoginUser{}

	err = client.MarshalGroup(id, config)
	check(ctx, err)
 
	return nil
}

func junosSystemLoginUserUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     

	config := xmlSystemLoginUser{}
	config.Groups.Name = id

    err = client.SendTransaction(id, config, false)
    check(ctx, err)
    
	return junosSystemLoginUserRead(ctx,d,m)
}

func junosSystemLoginUserDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	_, err = client.DeleteConfig(id,false)
    check(ctx, err)

    d.SetId("")
    
	return nil
}

func junosSystemLoginUser() *schema.Resource {
	return &schema.Resource{
		CreateContext: junosSystemLoginUserCreate,
		ReadContext: junosSystemLoginUserRead,
		UpdateContext: junosSystemLoginUserUpdate,
		DeleteContext: junosSystemLoginUserDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
		},
	}
}