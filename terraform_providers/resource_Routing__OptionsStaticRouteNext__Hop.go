
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
type xmlRouting__OptionsStaticRouteNext__Hop struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_route  struct {
			XMLName xml.Name `xml:"route"`
			V_name  *string  `xml:"name,omitempty"`
			V_next__hop  *string  `xml:"next-hop,omitempty"`
		} `xml:"routing-options>static>route"`
	} `xml:"groups"`

}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosRouting__OptionsStaticRouteNext__HopCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_next__hop := d.Get("next__hop").(string)


	config := xmlRouting__OptionsStaticRouteNext__Hop{}
	config.Groups.Name = id
	config.Groups.V_route.V_name = &V_name
	config.Groups.V_route.V_next__hop = &V_next__hop

    err = client.SendTransaction("", config, false)
    check(ctx, err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosRouting__OptionsStaticRouteNext__HopRead(ctx,d,m)
}

func junosRouting__OptionsStaticRouteNext__HopRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlRouting__OptionsStaticRouteNext__Hop{}

	err = client.MarshalGroup(id, config)
	check(ctx, err)
 	if err :=d.Set("name", config.Groups.V_route.V_name);err != nil{
		return diag.FromErr(err)
	}
	if err :=d.Set("next__hop", config.Groups.V_route.V_next__hop);err != nil{
		return diag.FromErr(err)
	}

	return nil
}

func junosRouting__OptionsStaticRouteNext__HopUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_next__hop := d.Get("next__hop").(string)


	config := xmlRouting__OptionsStaticRouteNext__Hop{}
	config.Groups.Name = id
	config.Groups.V_route.V_name = &V_name
	config.Groups.V_route.V_next__hop = &V_next__hop

    err = client.SendTransaction(id, config, false)
    check(ctx, err)
    
	return junosRouting__OptionsStaticRouteNext__HopRead(ctx,d,m)
}

func junosRouting__OptionsStaticRouteNext__HopDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	_, err = client.DeleteConfig(id,false)
    check(ctx, err)

    d.SetId("")
    
	return nil
}

func junosRouting__OptionsStaticRouteNext__Hop() *schema.Resource {
	return &schema.Resource{
		CreateContext: junosRouting__OptionsStaticRouteNext__HopCreate,
		ReadContext: junosRouting__OptionsStaticRouteNext__HopRead,
		UpdateContext: junosRouting__OptionsStaticRouteNext__HopUpdate,
		DeleteContext: junosRouting__OptionsStaticRouteNext__HopDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"name": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_route",
			},
			"next__hop": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_route. Next hop to destination",
			},
		},
	}
}