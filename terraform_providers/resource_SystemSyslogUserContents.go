
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
type xmlSystemSyslogUserContents struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_user  struct {
			XMLName xml.Name `xml:"user"`
			V_name  *string  `xml:"name,omitempty"`
			V_contents  struct {
				XMLName xml.Name `xml:"contents"`
			} `xml:"contents"`
		} `xml:"system>syslog>user"`
	} `xml:"groups"`

}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSystemSyslogUserContentsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)


	config := xmlSystemSyslogUserContents{}
	config.Groups.Name = id
	config.Groups.V_user.V_name = &V_name

    err = client.SendTransaction("", config, false)
    check(ctx, err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSystemSyslogUserContentsRead(ctx,d,m)
}

func junosSystemSyslogUserContentsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSystemSyslogUserContents{}

	err = client.MarshalGroup(id, config)
	check(ctx, err)
 	if err :=d.Set("name", config.Groups.V_user.V_name);err != nil{
		return diag.FromErr(err)
	}

	return nil
}

func junosSystemSyslogUserContentsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)


	config := xmlSystemSyslogUserContents{}
	config.Groups.Name = id
	config.Groups.V_user.V_name = &V_name

    err = client.SendTransaction(id, config, false)
    check(ctx, err)
    
	return junosSystemSyslogUserContentsRead(ctx,d,m)
}

func junosSystemSyslogUserContentsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	_, err = client.DeleteConfig(id,false)
    check(ctx, err)

    d.SetId("")
    
	return nil
}

func junosSystemSyslogUserContents() *schema.Resource {
	return &schema.Resource{
		CreateContext: junosSystemSyslogUserContentsCreate,
		ReadContext: junosSystemSyslogUserContentsRead,
		UpdateContext: junosSystemSyslogUserContentsUpdate,
		DeleteContext: junosSystemSyslogUserContentsDelete,

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
		},
	}
}