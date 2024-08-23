
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
type xmlSystemExtensionsProvidersLicense__Type struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_providers  struct {
			XMLName xml.Name `xml:"providers"`
			V_name  *string  `xml:"name,omitempty"`
			V_license__type  struct {
				XMLName xml.Name `xml:"license-type"`
			} `xml:"license-type"`
		} `xml:"system>extensions>providers"`
	} `xml:"groups"`

}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSystemExtensionsProvidersLicense__TypeCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)


	config := xmlSystemExtensionsProvidersLicense__Type{}
	config.Groups.Name = id
	config.Groups.V_providers.V_name = &V_name

    err = client.SendTransaction("", config, false)
    check(ctx, err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSystemExtensionsProvidersLicense__TypeRead(ctx,d,m)
}

func junosSystemExtensionsProvidersLicense__TypeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSystemExtensionsProvidersLicense__Type{}

	err = client.MarshalGroup(id, config)
	check(ctx, err)
 	if err :=d.Set("name", config.Groups.V_providers.V_name);err != nil{
		return diag.FromErr(err)
	}

	return nil
}

func junosSystemExtensionsProvidersLicense__TypeUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)


	config := xmlSystemExtensionsProvidersLicense__Type{}
	config.Groups.Name = id
	config.Groups.V_providers.V_name = &V_name

    err = client.SendTransaction(id, config, false)
    check(ctx, err)
    
	return junosSystemExtensionsProvidersLicense__TypeRead(ctx,d,m)
}

func junosSystemExtensionsProvidersLicense__TypeDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	_, err = client.DeleteConfig(id,false)
    check(ctx, err)

    d.SetId("")
    
	return nil
}

func junosSystemExtensionsProvidersLicense__Type() *schema.Resource {
	return &schema.Resource{
		CreateContext: junosSystemExtensionsProvidersLicense__TypeCreate,
		ReadContext: junosSystemExtensionsProvidersLicense__TypeRead,
		UpdateContext: junosSystemExtensionsProvidersLicense__TypeUpdate,
		DeleteContext: junosSystemExtensionsProvidersLicense__TypeDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"name": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_providers",
			},
		},
	}
}