
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
type xmlSnmpCommunityAuthorization struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_community  struct {
			XMLName xml.Name `xml:"community"`
			V_name  *string  `xml:"name,omitempty"`
			V_authorization  *string  `xml:"authorization,omitempty"`
		} `xml:"snmp>community"`
	} `xml:"groups"`

}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSnmpCommunityAuthorizationCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_authorization := d.Get("authorization").(string)


	config := xmlSnmpCommunityAuthorization{}
	config.Groups.Name = id
	config.Groups.V_community.V_name = &V_name
	config.Groups.V_community.V_authorization = &V_authorization

    err = client.SendTransaction("", config, false)
    check(ctx, err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSnmpCommunityAuthorizationRead(ctx,d,m)
}

func junosSnmpCommunityAuthorizationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSnmpCommunityAuthorization{}

	err = client.MarshalGroup(id, config)
	check(ctx, err)
 	if err :=d.Set("name", config.Groups.V_community.V_name);err != nil{
		return diag.FromErr(err)
	}
	if err :=d.Set("authorization", config.Groups.V_community.V_authorization);err != nil{
		return diag.FromErr(err)
	}

	return nil
}

func junosSnmpCommunityAuthorizationUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_authorization := d.Get("authorization").(string)


	config := xmlSnmpCommunityAuthorization{}
	config.Groups.Name = id
	config.Groups.V_community.V_name = &V_name
	config.Groups.V_community.V_authorization = &V_authorization

    err = client.SendTransaction(id, config, false)
    check(ctx, err)
    
	return junosSnmpCommunityAuthorizationRead(ctx,d,m)
}

func junosSnmpCommunityAuthorizationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	_, err = client.DeleteConfig(id,false)
    check(ctx, err)

    d.SetId("")
    
	return nil
}

func junosSnmpCommunityAuthorization() *schema.Resource {
	return &schema.Resource{
		CreateContext: junosSnmpCommunityAuthorizationCreate,
		ReadContext: junosSnmpCommunityAuthorizationRead,
		UpdateContext: junosSnmpCommunityAuthorizationUpdate,
		DeleteContext: junosSnmpCommunityAuthorizationDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"name": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_community",
			},
			"authorization": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_community. Authorization type",
			},
		},
	}
}