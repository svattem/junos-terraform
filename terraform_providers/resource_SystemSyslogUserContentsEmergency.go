
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
type xmlSystemSyslogUserContentsEmergency struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_user  struct {
			XMLName xml.Name `xml:"user"`
			V_name  *string  `xml:"name,omitempty"`
			V_contents  struct {
				XMLName xml.Name `xml:"contents"`
				V_name__1  *string  `xml:"name,omitempty"`
				V_emergency  *string  `xml:"emergency,omitempty"`
			} `xml:"contents"`
		} `xml:"system>syslog>user"`
	} `xml:"groups"`

}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSystemSyslogUserContentsEmergencyCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_name__1 := d.Get("name__1").(string)
	V_emergency := d.Get("emergency").(string)


	config := xmlSystemSyslogUserContentsEmergency{}
	config.Groups.Name = id
	config.Groups.V_user.V_name = &V_name
	config.Groups.V_user.V_contents.V_name__1 = &V_name__1
	config.Groups.V_user.V_contents.V_emergency = &V_emergency

    err = client.SendTransaction("", config, false)
    check(ctx, err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSystemSyslogUserContentsEmergencyRead(ctx,d,m)
}

func junosSystemSyslogUserContentsEmergencyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSystemSyslogUserContentsEmergency{}

	err = client.MarshalGroup(id, config)
	check(ctx, err)
 	if err :=d.Set("name", config.Groups.V_user.V_name);err != nil{
		return diag.FromErr(err)
	}
	if err :=d.Set("name__1", config.Groups.V_user.V_contents.V_name__1);err != nil{
		return diag.FromErr(err)
	}
	if err :=d.Set("emergency", config.Groups.V_user.V_contents.V_emergency);err != nil{
		return diag.FromErr(err)
	}

	return nil
}

func junosSystemSyslogUserContentsEmergencyUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_name__1 := d.Get("name__1").(string)
	V_emergency := d.Get("emergency").(string)


	config := xmlSystemSyslogUserContentsEmergency{}
	config.Groups.Name = id
	config.Groups.V_user.V_name = &V_name
	config.Groups.V_user.V_contents.V_name__1 = &V_name__1
	config.Groups.V_user.V_contents.V_emergency = &V_emergency

    err = client.SendTransaction(id, config, false)
    check(ctx, err)
    
	return junosSystemSyslogUserContentsEmergencyRead(ctx,d,m)
}

func junosSystemSyslogUserContentsEmergencyDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	_, err = client.DeleteConfig(id,false)
    check(ctx, err)

    d.SetId("")
    
	return nil
}

func junosSystemSyslogUserContentsEmergency() *schema.Resource {
	return &schema.Resource{
		CreateContext: junosSystemSyslogUserContentsEmergencyCreate,
		ReadContext: junosSystemSyslogUserContentsEmergencyRead,
		UpdateContext: junosSystemSyslogUserContentsEmergencyUpdate,
		DeleteContext: junosSystemSyslogUserContentsEmergencyDelete,

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
			"name__1": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_user.V_contents",
			},
			"emergency": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_user.V_contents. Panic conditions",
			},
		},
	}
}