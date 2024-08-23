
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
type xmlSystemLoginUserClass struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_user  struct {
			XMLName xml.Name `xml:"user"`
			V_name  *string  `xml:"name,omitempty"`
			V_class  *string  `xml:"class,omitempty"`
		} `xml:"system>login>user"`
	} `xml:"groups"`

}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSystemLoginUserClassCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_class := d.Get("class").(string)


	config := xmlSystemLoginUserClass{}
	config.Groups.Name = id
	config.Groups.V_user.V_name = &V_name
	config.Groups.V_user.V_class = &V_class

    err = client.SendTransaction("", config, false)
    check(ctx, err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSystemLoginUserClassRead(ctx,d,m)
}

func junosSystemLoginUserClassRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSystemLoginUserClass{}

	err = client.MarshalGroup(id, config)
	check(ctx, err)
 	if err :=d.Set("name", config.Groups.V_user.V_name);err != nil{
		return diag.FromErr(err)
	}
	if err :=d.Set("class", config.Groups.V_user.V_class);err != nil{
		return diag.FromErr(err)
	}

	return nil
}

func junosSystemLoginUserClassUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_class := d.Get("class").(string)


	config := xmlSystemLoginUserClass{}
	config.Groups.Name = id
	config.Groups.V_user.V_name = &V_name
	config.Groups.V_user.V_class = &V_class

    err = client.SendTransaction(id, config, false)
    check(ctx, err)
    
	return junosSystemLoginUserClassRead(ctx,d,m)
}

func junosSystemLoginUserClassDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	_, err = client.DeleteConfig(id,false)
    check(ctx, err)

    d.SetId("")
    
	return nil
}

func junosSystemLoginUserClass() *schema.Resource {
	return &schema.Resource{
		CreateContext: junosSystemLoginUserClassCreate,
		ReadContext: junosSystemLoginUserClassRead,
		UpdateContext: junosSystemLoginUserClassUpdate,
		DeleteContext: junosSystemLoginUserClassDelete,

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
			"class": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_user. Login class",
			},
		},
	}
}