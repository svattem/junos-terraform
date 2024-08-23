
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
type xmlSystemServicesRestHttpPort struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_http  struct {
			XMLName xml.Name `xml:"http"`
			V_port  *string  `xml:"port,omitempty"`
		} `xml:"system>services>rest>http"`
	} `xml:"groups"`

}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSystemServicesRestHttpPortCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_port := d.Get("port").(string)


	config := xmlSystemServicesRestHttpPort{}
	config.Groups.Name = id
	config.Groups.V_http.V_port = &V_port

    err = client.SendTransaction("", config, false)
    check(ctx, err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSystemServicesRestHttpPortRead(ctx,d,m)
}

func junosSystemServicesRestHttpPortRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSystemServicesRestHttpPort{}

	err = client.MarshalGroup(id, config)
	check(ctx, err)
 	if err :=d.Set("port", config.Groups.V_http.V_port);err != nil{
		return diag.FromErr(err)
	}

	return nil
}

func junosSystemServicesRestHttpPortUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_port := d.Get("port").(string)


	config := xmlSystemServicesRestHttpPort{}
	config.Groups.Name = id
	config.Groups.V_http.V_port = &V_port

    err = client.SendTransaction(id, config, false)
    check(ctx, err)
    
	return junosSystemServicesRestHttpPortRead(ctx,d,m)
}

func junosSystemServicesRestHttpPortDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	_, err = client.DeleteConfig(id,false)
    check(ctx, err)

    d.SetId("")
    
	return nil
}

func junosSystemServicesRestHttpPort() *schema.Resource {
	return &schema.Resource{
		CreateContext: junosSystemServicesRestHttpPortCreate,
		ReadContext: junosSystemServicesRestHttpPortRead,
		UpdateContext: junosSystemServicesRestHttpPortUpdate,
		DeleteContext: junosSystemServicesRestHttpPortDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"port": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_http. Port number to accept HTTP connections",
			},
		},
	}
}