
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
type xmlSystemLoginUserAuthenticationEncrypted__Password struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_user  struct {
			XMLName xml.Name `xml:"user"`
			V_name  *string  `xml:"name,omitempty"`
			V_authentication  struct {
				XMLName xml.Name `xml:"authentication"`
				V_encrypted__password  *string  `xml:"encrypted-password,omitempty"`
			} `xml:"authentication"`
		} `xml:"system>login>user"`
	} `xml:"groups"`

}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSystemLoginUserAuthenticationEncrypted__PasswordCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_encrypted__password := d.Get("encrypted__password").(string)


	config := xmlSystemLoginUserAuthenticationEncrypted__Password{}
	config.Groups.Name = id
	config.Groups.V_user.V_name = &V_name
	config.Groups.V_user.V_authentication.V_encrypted__password = &V_encrypted__password

    err = client.SendTransaction("", config, false)
    check(ctx, err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSystemLoginUserAuthenticationEncrypted__PasswordRead(ctx,d,m)
}

func junosSystemLoginUserAuthenticationEncrypted__PasswordRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSystemLoginUserAuthenticationEncrypted__Password{}

	err = client.MarshalGroup(id, config)
	check(ctx, err)
 	if err :=d.Set("name", config.Groups.V_user.V_name);err != nil{
		return diag.FromErr(err)
	}
	if err :=d.Set("encrypted__password", config.Groups.V_user.V_authentication.V_encrypted__password);err != nil{
		return diag.FromErr(err)
	}

	return nil
}

func junosSystemLoginUserAuthenticationEncrypted__PasswordUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_encrypted__password := d.Get("encrypted__password").(string)


	config := xmlSystemLoginUserAuthenticationEncrypted__Password{}
	config.Groups.Name = id
	config.Groups.V_user.V_name = &V_name
	config.Groups.V_user.V_authentication.V_encrypted__password = &V_encrypted__password

    err = client.SendTransaction(id, config, false)
    check(ctx, err)
    
	return junosSystemLoginUserAuthenticationEncrypted__PasswordRead(ctx,d,m)
}

func junosSystemLoginUserAuthenticationEncrypted__PasswordDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	_, err = client.DeleteConfig(id,false)
    check(ctx, err)

    d.SetId("")
    
	return nil
}

func junosSystemLoginUserAuthenticationEncrypted__Password() *schema.Resource {
	return &schema.Resource{
		CreateContext: junosSystemLoginUserAuthenticationEncrypted__PasswordCreate,
		ReadContext: junosSystemLoginUserAuthenticationEncrypted__PasswordRead,
		UpdateContext: junosSystemLoginUserAuthenticationEncrypted__PasswordUpdate,
		DeleteContext: junosSystemLoginUserAuthenticationEncrypted__PasswordDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"name": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_user",
			},
			"encrypted__password": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_user.V_authentication. Encrypted password string",
			},
		},
	}
}