
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
type xmlSystemServicesRestEnable__Explorer struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_rest  struct {
			XMLName xml.Name `xml:"rest"`
			V_enable__explorer  *string  `xml:"enable-explorer,omitempty"`
		} `xml:"system>services>rest"`
	} `xml:"groups"`

}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSystemServicesRestEnable__ExplorerCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_enable__explorer := d.Get("enable__explorer").(string)


	config := xmlSystemServicesRestEnable__Explorer{}
	config.Groups.Name = id
	config.Groups.V_rest.V_enable__explorer = &V_enable__explorer

    err = client.SendTransaction("", config, false)
    check(ctx, err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSystemServicesRestEnable__ExplorerRead(ctx,d,m)
}

func junosSystemServicesRestEnable__ExplorerRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSystemServicesRestEnable__Explorer{}

	err = client.MarshalGroup(id, config)
	check(ctx, err)
 	if err :=d.Set("enable__explorer", config.Groups.V_rest.V_enable__explorer);err != nil{
		return diag.FromErr(err)
	}

	return nil
}

func junosSystemServicesRestEnable__ExplorerUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_enable__explorer := d.Get("enable__explorer").(string)


	config := xmlSystemServicesRestEnable__Explorer{}
	config.Groups.Name = id
	config.Groups.V_rest.V_enable__explorer = &V_enable__explorer

    err = client.SendTransaction(id, config, false)
    check(ctx, err)
    
	return junosSystemServicesRestEnable__ExplorerRead(ctx,d,m)
}

func junosSystemServicesRestEnable__ExplorerDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	_, err = client.DeleteConfig(id,false)
    check(ctx, err)

    d.SetId("")
    
	return nil
}

func junosSystemServicesRestEnable__Explorer() *schema.Resource {
	return &schema.Resource{
		CreateContext: junosSystemServicesRestEnable__ExplorerCreate,
		ReadContext: junosSystemServicesRestEnable__ExplorerRead,
		UpdateContext: junosSystemServicesRestEnable__ExplorerUpdate,
		DeleteContext: junosSystemServicesRestEnable__ExplorerDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"enable__explorer": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_rest. Enable REST API explorer tool",
			},
		},
	}
}