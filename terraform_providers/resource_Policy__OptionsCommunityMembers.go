
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
type xmlPolicy__OptionsCommunityMembers struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_community  struct {
			XMLName xml.Name `xml:"community"`
			V_name  *string  `xml:"name,omitempty"`
			V_members  *string  `xml:"members,omitempty"`
		} `xml:"policy-options>community"`
	} `xml:"groups"`

}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosPolicy__OptionsCommunityMembersCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_members := d.Get("members").(string)


	config := xmlPolicy__OptionsCommunityMembers{}
	config.Groups.Name = id
	config.Groups.V_community.V_name = &V_name
	config.Groups.V_community.V_members = &V_members

    err = client.SendTransaction("", config, false)
    check(ctx, err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosPolicy__OptionsCommunityMembersRead(ctx,d,m)
}

func junosPolicy__OptionsCommunityMembersRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlPolicy__OptionsCommunityMembers{}

	err = client.MarshalGroup(id, config)
	check(ctx, err)
 	if err :=d.Set("name", config.Groups.V_community.V_name);err != nil{
		return diag.FromErr(err)
	}
	if err :=d.Set("members", config.Groups.V_community.V_members);err != nil{
		return diag.FromErr(err)
	}

	return nil
}

func junosPolicy__OptionsCommunityMembersUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_members := d.Get("members").(string)


	config := xmlPolicy__OptionsCommunityMembers{}
	config.Groups.Name = id
	config.Groups.V_community.V_name = &V_name
	config.Groups.V_community.V_members = &V_members

    err = client.SendTransaction(id, config, false)
    check(ctx, err)
    
	return junosPolicy__OptionsCommunityMembersRead(ctx,d,m)
}

func junosPolicy__OptionsCommunityMembersDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	_, err = client.DeleteConfig(id,false)
    check(ctx, err)

    d.SetId("")
    
	return nil
}

func junosPolicy__OptionsCommunityMembers() *schema.Resource {
	return &schema.Resource{
		CreateContext: junosPolicy__OptionsCommunityMembersCreate,
		ReadContext: junosPolicy__OptionsCommunityMembersRead,
		UpdateContext: junosPolicy__OptionsCommunityMembersUpdate,
		DeleteContext: junosPolicy__OptionsCommunityMembersDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"name": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_community",
			},
			"members": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_community. Community members",
			},
		},
	}
}