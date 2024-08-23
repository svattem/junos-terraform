
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
type xmlSystemRoot__AuthenticationEncrypted__Password struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_root__authentication  struct {
			XMLName xml.Name `xml:"root-authentication"`
			V_encrypted__password  *string  `xml:"encrypted-password,omitempty"`
		} `xml:"system>root-authentication"`
	} `xml:"groups"`

}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSystemRoot__AuthenticationEncrypted__PasswordCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_encrypted__password := d.Get("encrypted__password").(string)


	config := xmlSystemRoot__AuthenticationEncrypted__Password{}
	config.Groups.Name = id
	config.Groups.V_root__authentication.V_encrypted__password = &V_encrypted__password

    err = client.SendTransaction("", config, false)
    check(ctx, err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSystemRoot__AuthenticationEncrypted__PasswordRead(ctx,d,m)
}

func junosSystemRoot__AuthenticationEncrypted__PasswordRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSystemRoot__AuthenticationEncrypted__Password{}

	err = client.MarshalGroup(id, config)
	check(ctx, err)
 	if err :=d.Set("encrypted__password", config.Groups.V_root__authentication.V_encrypted__password);err != nil{
		return diag.FromErr(err)
	}

	return nil
}

func junosSystemRoot__AuthenticationEncrypted__PasswordUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_encrypted__password := d.Get("encrypted__password").(string)


	config := xmlSystemRoot__AuthenticationEncrypted__Password{}
	config.Groups.Name = id
	config.Groups.V_root__authentication.V_encrypted__password = &V_encrypted__password

    err = client.SendTransaction(id, config, false)
    check(ctx, err)
    
	return junosSystemRoot__AuthenticationEncrypted__PasswordRead(ctx,d,m)
}

func junosSystemRoot__AuthenticationEncrypted__PasswordDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	_, err = client.DeleteConfig(id,false)
    check(ctx, err)

    d.SetId("")
    
	return nil
}

func junosSystemRoot__AuthenticationEncrypted__Password() *schema.Resource {
	return &schema.Resource{
		CreateContext: junosSystemRoot__AuthenticationEncrypted__PasswordCreate,
		ReadContext: junosSystemRoot__AuthenticationEncrypted__PasswordRead,
		UpdateContext: junosSystemRoot__AuthenticationEncrypted__PasswordUpdate,
		DeleteContext: junosSystemRoot__AuthenticationEncrypted__PasswordDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"encrypted__password": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_root__authentication. Encrypted password string",
			},
		},
	}
}