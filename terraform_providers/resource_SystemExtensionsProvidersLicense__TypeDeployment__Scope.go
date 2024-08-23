
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
type xmlSystemExtensionsProvidersLicense__TypeDeployment__Scope struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_providers  struct {
			XMLName xml.Name `xml:"providers"`
			V_name  *string  `xml:"name,omitempty"`
			V_license__type  struct {
				XMLName xml.Name `xml:"license-type"`
				V_name__1  *string  `xml:"name,omitempty"`
				V_deployment__scope  *string  `xml:"deployment-scope,omitempty"`
			} `xml:"license-type"`
		} `xml:"system>extensions>providers"`
	} `xml:"groups"`

}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSystemExtensionsProvidersLicense__TypeDeployment__ScopeCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_name__1 := d.Get("name__1").(string)
	V_deployment__scope := d.Get("deployment__scope").(string)


	config := xmlSystemExtensionsProvidersLicense__TypeDeployment__Scope{}
	config.Groups.Name = id
	config.Groups.V_providers.V_name = &V_name
	config.Groups.V_providers.V_license__type.V_name__1 = &V_name__1
	config.Groups.V_providers.V_license__type.V_deployment__scope = &V_deployment__scope

    err = client.SendTransaction("", config, false)
    check(ctx, err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSystemExtensionsProvidersLicense__TypeDeployment__ScopeRead(ctx,d,m)
}

func junosSystemExtensionsProvidersLicense__TypeDeployment__ScopeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSystemExtensionsProvidersLicense__TypeDeployment__Scope{}

	err = client.MarshalGroup(id, config)
	check(ctx, err)
 	if err :=d.Set("name", config.Groups.V_providers.V_name);err != nil{
		return diag.FromErr(err)
	}
	if err :=d.Set("name__1", config.Groups.V_providers.V_license__type.V_name__1);err != nil{
		return diag.FromErr(err)
	}
	if err :=d.Set("deployment__scope", config.Groups.V_providers.V_license__type.V_deployment__scope);err != nil{
		return diag.FromErr(err)
	}

	return nil
}

func junosSystemExtensionsProvidersLicense__TypeDeployment__ScopeUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_name__1 := d.Get("name__1").(string)
	V_deployment__scope := d.Get("deployment__scope").(string)


	config := xmlSystemExtensionsProvidersLicense__TypeDeployment__Scope{}
	config.Groups.Name = id
	config.Groups.V_providers.V_name = &V_name
	config.Groups.V_providers.V_license__type.V_name__1 = &V_name__1
	config.Groups.V_providers.V_license__type.V_deployment__scope = &V_deployment__scope

    err = client.SendTransaction(id, config, false)
    check(ctx, err)
    
	return junosSystemExtensionsProvidersLicense__TypeDeployment__ScopeRead(ctx,d,m)
}

func junosSystemExtensionsProvidersLicense__TypeDeployment__ScopeDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	_, err = client.DeleteConfig(id,false)
    check(ctx, err)

    d.SetId("")
    
	return nil
}

func junosSystemExtensionsProvidersLicense__TypeDeployment__Scope() *schema.Resource {
	return &schema.Resource{
		CreateContext: junosSystemExtensionsProvidersLicense__TypeDeployment__ScopeCreate,
		ReadContext: junosSystemExtensionsProvidersLicense__TypeDeployment__ScopeRead,
		UpdateContext: junosSystemExtensionsProvidersLicense__TypeDeployment__ScopeUpdate,
		DeleteContext: junosSystemExtensionsProvidersLicense__TypeDeployment__ScopeDelete,

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
			"name__1": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_providers.V_license__type",
			},
			"deployment__scope": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_providers.V_license__type. ",
			},
		},
	}
}