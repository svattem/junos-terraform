
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
type xmlInterfacesInterfaceUnitDescription struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_interface  struct {
			XMLName xml.Name `xml:"interface"`
			V_name  *string  `xml:"name,omitempty"`
			V_unit  struct {
				XMLName xml.Name `xml:"unit"`
				V_name__1  *string  `xml:"name,omitempty"`
				V_description  *string  `xml:"description,omitempty"`
			} `xml:"unit"`
		} `xml:"interfaces>interface"`
	} `xml:"groups"`

}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosInterfacesInterfaceUnitDescriptionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_name__1 := d.Get("name__1").(string)
	V_description := d.Get("description").(string)


	config := xmlInterfacesInterfaceUnitDescription{}
	config.Groups.Name = id
	config.Groups.V_interface.V_name = &V_name
	config.Groups.V_interface.V_unit.V_name__1 = &V_name__1
	config.Groups.V_interface.V_unit.V_description = &V_description

    err = client.SendTransaction("", config, false)
    check(ctx, err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosInterfacesInterfaceUnitDescriptionRead(ctx,d,m)
}

func junosInterfacesInterfaceUnitDescriptionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlInterfacesInterfaceUnitDescription{}

	err = client.MarshalGroup(id, config)
	check(ctx, err)
 	if err :=d.Set("name", config.Groups.V_interface.V_name);err != nil{
		return diag.FromErr(err)
	}
	if err :=d.Set("name__1", config.Groups.V_interface.V_unit.V_name__1);err != nil{
		return diag.FromErr(err)
	}
	if err :=d.Set("description", config.Groups.V_interface.V_unit.V_description);err != nil{
		return diag.FromErr(err)
	}

	return nil
}

func junosInterfacesInterfaceUnitDescriptionUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_name__1 := d.Get("name__1").(string)
	V_description := d.Get("description").(string)


	config := xmlInterfacesInterfaceUnitDescription{}
	config.Groups.Name = id
	config.Groups.V_interface.V_name = &V_name
	config.Groups.V_interface.V_unit.V_name__1 = &V_name__1
	config.Groups.V_interface.V_unit.V_description = &V_description

    err = client.SendTransaction(id, config, false)
    check(ctx, err)
    
	return junosInterfacesInterfaceUnitDescriptionRead(ctx,d,m)
}

func junosInterfacesInterfaceUnitDescriptionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	_, err = client.DeleteConfig(id,false)
    check(ctx, err)

    d.SetId("")
    
	return nil
}

func junosInterfacesInterfaceUnitDescription() *schema.Resource {
	return &schema.Resource{
		CreateContext: junosInterfacesInterfaceUnitDescriptionCreate,
		ReadContext: junosInterfacesInterfaceUnitDescriptionRead,
		UpdateContext: junosInterfacesInterfaceUnitDescriptionUpdate,
		DeleteContext: junosInterfacesInterfaceUnitDescriptionDelete,

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
			"name__1": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_interface.V_unit",
			},
			"description": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_interface.V_unit. Text description of interface",
			},
		},
	}
}