
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
type xmlRouting__OptionsForwarding__TableExport struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_forwarding__table  struct {
			XMLName xml.Name `xml:"forwarding-table"`
			V_export  *string  `xml:"export,omitempty"`
		} `xml:"routing-options>forwarding-table"`
	} `xml:"groups"`

}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosRouting__OptionsForwarding__TableExportCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_export := d.Get("export").(string)


	config := xmlRouting__OptionsForwarding__TableExport{}
	config.Groups.Name = id
	config.Groups.V_forwarding__table.V_export = &V_export

    err = client.SendTransaction("", config, false)
    check(ctx, err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosRouting__OptionsForwarding__TableExportRead(ctx,d,m)
}

func junosRouting__OptionsForwarding__TableExportRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlRouting__OptionsForwarding__TableExport{}

	err = client.MarshalGroup(id, config)
	check(ctx, err)
 	if err :=d.Set("export", config.Groups.V_forwarding__table.V_export);err != nil{
		return diag.FromErr(err)
	}

	return nil
}

func junosRouting__OptionsForwarding__TableExportUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_export := d.Get("export").(string)


	config := xmlRouting__OptionsForwarding__TableExport{}
	config.Groups.Name = id
	config.Groups.V_forwarding__table.V_export = &V_export

    err = client.SendTransaction(id, config, false)
    check(ctx, err)
    
	return junosRouting__OptionsForwarding__TableExportRead(ctx,d,m)
}

func junosRouting__OptionsForwarding__TableExportDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	_, err = client.DeleteConfig(id,false)
    check(ctx, err)

    d.SetId("")
    
	return nil
}

func junosRouting__OptionsForwarding__TableExport() *schema.Resource {
	return &schema.Resource{
		CreateContext: junosRouting__OptionsForwarding__TableExportCreate,
		ReadContext: junosRouting__OptionsForwarding__TableExportRead,
		UpdateContext: junosRouting__OptionsForwarding__TableExportUpdate,
		DeleteContext: junosRouting__OptionsForwarding__TableExportDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"export": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_forwarding__table. Export policy",
			},
		},
	}
}