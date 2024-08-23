
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
type xmlSystemServicesExtension__ServiceRequest__ResponseGrpcMax__Connections struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_grpc  struct {
			XMLName xml.Name `xml:"grpc"`
			V_max__connections  *string  `xml:"max-connections,omitempty"`
		} `xml:"system>services>extension-service>request-response>grpc"`
	} `xml:"groups"`

}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSystemServicesExtension__ServiceRequest__ResponseGrpcMax__ConnectionsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_max__connections := d.Get("max__connections").(string)


	config := xmlSystemServicesExtension__ServiceRequest__ResponseGrpcMax__Connections{}
	config.Groups.Name = id
	config.Groups.V_grpc.V_max__connections = &V_max__connections

    err = client.SendTransaction("", config, false)
    check(ctx, err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSystemServicesExtension__ServiceRequest__ResponseGrpcMax__ConnectionsRead(ctx,d,m)
}

func junosSystemServicesExtension__ServiceRequest__ResponseGrpcMax__ConnectionsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSystemServicesExtension__ServiceRequest__ResponseGrpcMax__Connections{}

	err = client.MarshalGroup(id, config)
	check(ctx, err)
 	if err :=d.Set("max__connections", config.Groups.V_grpc.V_max__connections);err != nil{
		return diag.FromErr(err)
	}

	return nil
}

func junosSystemServicesExtension__ServiceRequest__ResponseGrpcMax__ConnectionsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_max__connections := d.Get("max__connections").(string)


	config := xmlSystemServicesExtension__ServiceRequest__ResponseGrpcMax__Connections{}
	config.Groups.Name = id
	config.Groups.V_grpc.V_max__connections = &V_max__connections

    err = client.SendTransaction(id, config, false)
    check(ctx, err)
    
	return junosSystemServicesExtension__ServiceRequest__ResponseGrpcMax__ConnectionsRead(ctx,d,m)
}

func junosSystemServicesExtension__ServiceRequest__ResponseGrpcMax__ConnectionsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	_, err = client.DeleteConfig(id,false)
    check(ctx, err)

    d.SetId("")
    
	return nil
}

func junosSystemServicesExtension__ServiceRequest__ResponseGrpcMax__Connections() *schema.Resource {
	return &schema.Resource{
		CreateContext: junosSystemServicesExtension__ServiceRequest__ResponseGrpcMax__ConnectionsCreate,
		ReadContext: junosSystemServicesExtension__ServiceRequest__ResponseGrpcMax__ConnectionsRead,
		UpdateContext: junosSystemServicesExtension__ServiceRequest__ResponseGrpcMax__ConnectionsUpdate,
		DeleteContext: junosSystemServicesExtension__ServiceRequest__ResponseGrpcMax__ConnectionsDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"max__connections": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_grpc. Maximum number of connections",
			},
		},
	}
}