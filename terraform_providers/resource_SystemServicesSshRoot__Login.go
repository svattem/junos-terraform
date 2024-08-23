
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
type xmlSystemServicesSshRoot__Login struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_ssh  struct {
			XMLName xml.Name `xml:"ssh"`
			V_root__login  *string  `xml:"root-login,omitempty"`
		} `xml:"system>services>ssh"`
	} `xml:"groups"`

}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSystemServicesSshRoot__LoginCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_root__login := d.Get("root__login").(string)


	config := xmlSystemServicesSshRoot__Login{}
	config.Groups.Name = id
	config.Groups.V_ssh.V_root__login = &V_root__login

    err = client.SendTransaction("", config, false)
    check(ctx, err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSystemServicesSshRoot__LoginRead(ctx,d,m)
}

func junosSystemServicesSshRoot__LoginRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSystemServicesSshRoot__Login{}

	err = client.MarshalGroup(id, config)
	check(ctx, err)
 	if err :=d.Set("root__login", config.Groups.V_ssh.V_root__login);err != nil{
		return diag.FromErr(err)
	}

	return nil
}

func junosSystemServicesSshRoot__LoginUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_root__login := d.Get("root__login").(string)


	config := xmlSystemServicesSshRoot__Login{}
	config.Groups.Name = id
	config.Groups.V_ssh.V_root__login = &V_root__login

    err = client.SendTransaction(id, config, false)
    check(ctx, err)
    
	return junosSystemServicesSshRoot__LoginRead(ctx,d,m)
}

func junosSystemServicesSshRoot__LoginDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	_, err = client.DeleteConfig(id,false)
    check(ctx, err)

    d.SetId("")
    
	return nil
}

func junosSystemServicesSshRoot__Login() *schema.Resource {
	return &schema.Resource{
		CreateContext: junosSystemServicesSshRoot__LoginCreate,
		ReadContext: junosSystemServicesSshRoot__LoginRead,
		UpdateContext: junosSystemServicesSshRoot__LoginUpdate,
		DeleteContext: junosSystemServicesSshRoot__LoginDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"root__login": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_ssh. Configure root access via ssh",
			},
		},
	}
}