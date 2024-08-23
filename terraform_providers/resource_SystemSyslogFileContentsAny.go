
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
type xmlSystemSyslogFileContentsAny struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_file  struct {
			XMLName xml.Name `xml:"file"`
			V_name  *string  `xml:"name,omitempty"`
			V_contents  struct {
				XMLName xml.Name `xml:"contents"`
				V_name__1  *string  `xml:"name,omitempty"`
				V_any  *string  `xml:"any,omitempty"`
			} `xml:"contents"`
		} `xml:"system>syslog>file"`
	} `xml:"groups"`

}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSystemSyslogFileContentsAnyCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_name__1 := d.Get("name__1").(string)
	V_any := d.Get("any").(string)


	config := xmlSystemSyslogFileContentsAny{}
	config.Groups.Name = id
	config.Groups.V_file.V_name = &V_name
	config.Groups.V_file.V_contents.V_name__1 = &V_name__1
	config.Groups.V_file.V_contents.V_any = &V_any

    err = client.SendTransaction("", config, false)
    check(ctx, err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSystemSyslogFileContentsAnyRead(ctx,d,m)
}

func junosSystemSyslogFileContentsAnyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSystemSyslogFileContentsAny{}

	err = client.MarshalGroup(id, config)
	check(ctx, err)
 	if err :=d.Set("name", config.Groups.V_file.V_name);err != nil{
		return diag.FromErr(err)
	}
	if err :=d.Set("name__1", config.Groups.V_file.V_contents.V_name__1);err != nil{
		return diag.FromErr(err)
	}
	if err :=d.Set("any", config.Groups.V_file.V_contents.V_any);err != nil{
		return diag.FromErr(err)
	}

	return nil
}

func junosSystemSyslogFileContentsAnyUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_name__1 := d.Get("name__1").(string)
	V_any := d.Get("any").(string)


	config := xmlSystemSyslogFileContentsAny{}
	config.Groups.Name = id
	config.Groups.V_file.V_name = &V_name
	config.Groups.V_file.V_contents.V_name__1 = &V_name__1
	config.Groups.V_file.V_contents.V_any = &V_any

    err = client.SendTransaction(id, config, false)
    check(ctx, err)
    
	return junosSystemSyslogFileContentsAnyRead(ctx,d,m)
}

func junosSystemSyslogFileContentsAnyDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	_, err = client.DeleteConfig(id,false)
    check(ctx, err)

    d.SetId("")
    
	return nil
}

func junosSystemSyslogFileContentsAny() *schema.Resource {
	return &schema.Resource{
		CreateContext: junosSystemSyslogFileContentsAnyCreate,
		ReadContext: junosSystemSyslogFileContentsAnyRead,
		UpdateContext: junosSystemSyslogFileContentsAnyUpdate,
		DeleteContext: junosSystemSyslogFileContentsAnyDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"name": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_file",
			},
			"name__1": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_file.V_contents",
			},
			"any": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_file.V_contents. All levels",
			},
		},
	}
}