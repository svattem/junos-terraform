
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
type xmlSystemServicesExtension__ServiceRequest__ResponseGrpc struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_grpc  struct {
			XMLName xml.Name `xml:"grpc"`
		} `xml:"system>services>extension-service>request-response>grpc"`
	} `xml:"groups"`

}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSystemServicesExtension__ServiceRequest__ResponseGrpcCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     

	config := xmlSystemServicesExtension__ServiceRequest__ResponseGrpc{}
	config.Groups.Name = id

    err = client.SendTransaction("", config, false)
    check(ctx, err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosSystemServicesExtension__ServiceRequest__ResponseGrpcRead(ctx,d,m)
}

func junosSystemServicesExtension__ServiceRequest__ResponseGrpcRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlSystemServicesExtension__ServiceRequest__ResponseGrpc{}

	err = client.MarshalGroup(id, config)
	check(ctx, err)
 
	return nil
}

func junosSystemServicesExtension__ServiceRequest__ResponseGrpcUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     

	config := xmlSystemServicesExtension__ServiceRequest__ResponseGrpc{}
	config.Groups.Name = id

    err = client.SendTransaction(id, config, false)
    check(ctx, err)
    
	return junosSystemServicesExtension__ServiceRequest__ResponseGrpcRead(ctx,d,m)
}

func junosSystemServicesExtension__ServiceRequest__ResponseGrpcDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var err error
	client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	_, err = client.DeleteConfig(id,false)
    check(ctx, err)

    d.SetId("")
    
	return nil
}

func junosSystemServicesExtension__ServiceRequest__ResponseGrpc() *schema.Resource {
	return &schema.Resource{
		CreateContext: junosSystemServicesExtension__ServiceRequest__ResponseGrpcCreate,
		ReadContext: junosSystemServicesExtension__ServiceRequest__ResponseGrpcRead,
		UpdateContext: junosSystemServicesExtension__ServiceRequest__ResponseGrpcUpdate,
		DeleteContext: junosSystemServicesExtension__ServiceRequest__ResponseGrpcDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
		},
	}
}