
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
type xmlPolicy__OptionsPolicy__StatementThenLoad__BalancePer__Packet struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_policy__statement  struct {
			XMLName xml.Name `xml:"policy-statement"`
			V_name  *string  `xml:"name,omitempty"`
			V_load__balance  struct {
				XMLName xml.Name `xml:"load-balance"`
				V_per__packet  *string  `xml:"per-packet,omitempty"`
			} `xml:"then>load-balance"`
		} `xml:"policy-options>policy-statement"`
	} `xml:"groups"`

}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosPolicy__OptionsPolicy__StatementThenLoad__BalancePer__PacketCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_per__packet := d.Get("per__packet").(string)


	config := xmlPolicy__OptionsPolicy__StatementThenLoad__BalancePer__Packet{}
	config.Groups.Name = id
	config.Groups.V_policy__statement.V_name = &V_name
	config.Groups.V_policy__statement.V_load__balance.V_per__packet = &V_per__packet

    err = client.SendTransaction("", config, false)
    check(ctx, err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosPolicy__OptionsPolicy__StatementThenLoad__BalancePer__PacketRead(ctx,d,m)
}

func junosPolicy__OptionsPolicy__StatementThenLoad__BalancePer__PacketRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlPolicy__OptionsPolicy__StatementThenLoad__BalancePer__Packet{}

	err = client.MarshalGroup(id, config)
	check(ctx, err)
 	if err :=d.Set("name", config.Groups.V_policy__statement.V_name);err != nil{
		return diag.FromErr(err)
	}
	if err :=d.Set("per__packet", config.Groups.V_policy__statement.V_load__balance.V_per__packet);err != nil{
		return diag.FromErr(err)
	}

	return nil
}

func junosPolicy__OptionsPolicy__StatementThenLoad__BalancePer__PacketUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_per__packet := d.Get("per__packet").(string)


	config := xmlPolicy__OptionsPolicy__StatementThenLoad__BalancePer__Packet{}
	config.Groups.Name = id
	config.Groups.V_policy__statement.V_name = &V_name
	config.Groups.V_policy__statement.V_load__balance.V_per__packet = &V_per__packet

    err = client.SendTransaction(id, config, false)
    check(ctx, err)
    
	return junosPolicy__OptionsPolicy__StatementThenLoad__BalancePer__PacketRead(ctx,d,m)
}

func junosPolicy__OptionsPolicy__StatementThenLoad__BalancePer__PacketDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	_, err = client.DeleteConfig(id,false)
    check(ctx, err)

    d.SetId("")
    
	return nil
}

func junosPolicy__OptionsPolicy__StatementThenLoad__BalancePer__Packet() *schema.Resource {
	return &schema.Resource{
		CreateContext: junosPolicy__OptionsPolicy__StatementThenLoad__BalancePer__PacketCreate,
		ReadContext: junosPolicy__OptionsPolicy__StatementThenLoad__BalancePer__PacketRead,
		UpdateContext: junosPolicy__OptionsPolicy__StatementThenLoad__BalancePer__PacketUpdate,
		DeleteContext: junosPolicy__OptionsPolicy__StatementThenLoad__BalancePer__PacketDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"name": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_policy__statement",
			},
			"per__packet": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_policy__statement.V_load__balance. Load balance on a per-packet basis",
			},
		},
	}
}