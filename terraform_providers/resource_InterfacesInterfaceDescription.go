
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
type xmlInterfacesInterfaceDescription struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_interface  struct {
			XMLName xml.Name `xml:"interface"`
			V_name  *string  `xml:"name,omitempty"`
			V_description  *string  `xml:"description,omitempty"`
		} `xml:"interfaces>interface"`
	} `xml:"groups"`

}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosInterfacesInterfaceDescriptionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_description := d.Get("description").(string)


	config := xmlInterfacesInterfaceDescription{}
	config.Groups.Name = id
	config.Groups.V_interface.V_name = &V_name
	config.Groups.V_interface.V_description = &V_description

    err = client.SendTransaction("", config, false)
    check(ctx, err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosInterfacesInterfaceDescriptionRead(ctx,d,m)
}

func junosInterfacesInterfaceDescriptionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlInterfacesInterfaceDescription{}

	err = client.MarshalGroup(id, config)
	check(ctx, err)
 	if err :=d.Set("name", config.Groups.V_interface.V_name);err != nil{
		return diag.FromErr(err)
	}
	if err :=d.Set("description", config.Groups.V_interface.V_description);err != nil{
		return diag.FromErr(err)
	}

	return nil
}

func junosInterfacesInterfaceDescriptionUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_description := d.Get("description").(string)


	config := xmlInterfacesInterfaceDescription{}
	config.Groups.Name = id
	config.Groups.V_interface.V_name = &V_name
	config.Groups.V_interface.V_description = &V_description

    err = client.SendTransaction(id, config, false)
    check(ctx, err)
    
	return junosInterfacesInterfaceDescriptionRead(ctx,d,m)
}

func junosInterfacesInterfaceDescriptionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	_, err = client.DeleteConfig(id,false)
    check(ctx, err)

    d.SetId("")
    
	return nil
}

func junosInterfacesInterfaceDescription() *schema.Resource {
	return &schema.Resource{
		CreateContext: junosInterfacesInterfaceDescriptionCreate,
		ReadContext: junosInterfacesInterfaceDescriptionRead,
		UpdateContext: junosInterfacesInterfaceDescriptionUpdate,
		DeleteContext: junosInterfacesInterfaceDescriptionDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"name": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_interface",
			},
			"description": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_interface. Text description of interface",
			},
		},
	}
}