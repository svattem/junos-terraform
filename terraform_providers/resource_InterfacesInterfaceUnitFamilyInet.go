
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
type xmlInterfacesInterfaceUnitFamilyInet struct {
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
				V_inet  struct {
					XMLName xml.Name `xml:"inet"`
				} `xml:"family>inet"`
			} `xml:"unit"`
		} `xml:"interfaces>interface"`
	} `xml:"groups"`

}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosInterfacesInterfaceUnitFamilyInetCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_name__1 := d.Get("name__1").(string)


	config := xmlInterfacesInterfaceUnitFamilyInet{}
	config.Groups.Name = id
	config.Groups.V_interface.V_name = &V_name
	config.Groups.V_interface.V_unit.V_name__1 = &V_name__1

    err = client.SendTransaction("", config, false)
    check(ctx, err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosInterfacesInterfaceUnitFamilyInetRead(ctx,d,m)
}

func junosInterfacesInterfaceUnitFamilyInetRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlInterfacesInterfaceUnitFamilyInet{}

	err = client.MarshalGroup(id, config)
	check(ctx, err)
 	if err :=d.Set("name", config.Groups.V_interface.V_name);err != nil{
		return diag.FromErr(err)
	}
	if err :=d.Set("name__1", config.Groups.V_interface.V_unit.V_name__1);err != nil{
		return diag.FromErr(err)
	}

	return nil
}

func junosInterfacesInterfaceUnitFamilyInetUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_name__1 := d.Get("name__1").(string)


	config := xmlInterfacesInterfaceUnitFamilyInet{}
	config.Groups.Name = id
	config.Groups.V_interface.V_name = &V_name
	config.Groups.V_interface.V_unit.V_name__1 = &V_name__1

    err = client.SendTransaction(id, config, false)
    check(ctx, err)
    
	return junosInterfacesInterfaceUnitFamilyInetRead(ctx,d,m)
}

func junosInterfacesInterfaceUnitFamilyInetDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	_, err = client.DeleteConfig(id,false)
    check(ctx, err)

    d.SetId("")
    
	return nil
}

func junosInterfacesInterfaceUnitFamilyInet() *schema.Resource {
	return &schema.Resource{
		CreateContext: junosInterfacesInterfaceUnitFamilyInetCreate,
		ReadContext: junosInterfacesInterfaceUnitFamilyInetRead,
		UpdateContext: junosInterfacesInterfaceUnitFamilyInetUpdate,
		DeleteContext: junosInterfacesInterfaceUnitFamilyInetDelete,

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
		},
	}
}