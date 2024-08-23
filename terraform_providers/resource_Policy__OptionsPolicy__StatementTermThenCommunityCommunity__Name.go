
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
type xmlPolicy__OptionsPolicy__StatementTermThenCommunityCommunity__Name struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_policy__statement  struct {
			XMLName xml.Name `xml:"policy-statement"`
			V_name  *string  `xml:"name,omitempty"`
			V_term  struct {
				XMLName xml.Name `xml:"term"`
				V_name__1  *string  `xml:"name,omitempty"`
				V_community  struct {
					XMLName xml.Name `xml:"community"`
					V_community__name  *string  `xml:"community-name,omitempty"`
				} `xml:"then>community"`
			} `xml:"term"`
		} `xml:"policy-options>policy-statement"`
	} `xml:"groups"`

}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosPolicy__OptionsPolicy__StatementTermThenCommunityCommunity__NameCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_name__1 := d.Get("name__1").(string)
	V_community__name := d.Get("community__name").(string)


	config := xmlPolicy__OptionsPolicy__StatementTermThenCommunityCommunity__Name{}
	config.Groups.Name = id
	config.Groups.V_policy__statement.V_name = &V_name
	config.Groups.V_policy__statement.V_term.V_name__1 = &V_name__1
	config.Groups.V_policy__statement.V_term.V_community.V_community__name = &V_community__name

    err = client.SendTransaction("", config, false)
    check(ctx, err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosPolicy__OptionsPolicy__StatementTermThenCommunityCommunity__NameRead(ctx,d,m)
}

func junosPolicy__OptionsPolicy__StatementTermThenCommunityCommunity__NameRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlPolicy__OptionsPolicy__StatementTermThenCommunityCommunity__Name{}

	err = client.MarshalGroup(id, config)
	check(ctx, err)
 	if err :=d.Set("name", config.Groups.V_policy__statement.V_name);err != nil{
		return diag.FromErr(err)
	}
	if err :=d.Set("name__1", config.Groups.V_policy__statement.V_term.V_name__1);err != nil{
		return diag.FromErr(err)
	}
	if err :=d.Set("community__name", config.Groups.V_policy__statement.V_term.V_community.V_community__name);err != nil{
		return diag.FromErr(err)
	}

	return nil
}

func junosPolicy__OptionsPolicy__StatementTermThenCommunityCommunity__NameUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_name__1 := d.Get("name__1").(string)
	V_community__name := d.Get("community__name").(string)


	config := xmlPolicy__OptionsPolicy__StatementTermThenCommunityCommunity__Name{}
	config.Groups.Name = id
	config.Groups.V_policy__statement.V_name = &V_name
	config.Groups.V_policy__statement.V_term.V_name__1 = &V_name__1
	config.Groups.V_policy__statement.V_term.V_community.V_community__name = &V_community__name

    err = client.SendTransaction(id, config, false)
    check(ctx, err)
    
	return junosPolicy__OptionsPolicy__StatementTermThenCommunityCommunity__NameRead(ctx,d,m)
}

func junosPolicy__OptionsPolicy__StatementTermThenCommunityCommunity__NameDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	_, err = client.DeleteConfig(id,false)
    check(ctx, err)

    d.SetId("")
    
	return nil
}

func junosPolicy__OptionsPolicy__StatementTermThenCommunityCommunity__Name() *schema.Resource {
	return &schema.Resource{
		CreateContext: junosPolicy__OptionsPolicy__StatementTermThenCommunityCommunity__NameCreate,
		ReadContext: junosPolicy__OptionsPolicy__StatementTermThenCommunityCommunity__NameRead,
		UpdateContext: junosPolicy__OptionsPolicy__StatementTermThenCommunityCommunity__NameUpdate,
		DeleteContext: junosPolicy__OptionsPolicy__StatementTermThenCommunityCommunity__NameDelete,

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
			"name__1": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_policy__statement.V_term",
			},
			"community__name": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_policy__statement.V_term.V_community. Name to identify a BGP community",
			},
		},
	}
}