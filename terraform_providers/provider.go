
// Copyright (c) 2017-2022, Juniper Networks Inc. All rights reserved.
//
// License: Apache 2.0
//
// THIS SOFTWARE IS PROVIDED BY Juniper Networks, Inc. ''AS IS'' AND ANY
// EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL Juniper Networks, Inc. BE LIABLE FOR ANY
// DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//

package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"os"
	"sort"
	s "strings"
	"sync"
	"terraform-provider-junos-vsrx/netconf"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const groupStrXML = `<load-configuration action="merge" format="xml">
%s
</load-configuration>`

const deleteStr = `<edit-config>
	<target>
		<candidate/>
	</target>
	<default-operation>none</default-operation>
	<config>
		<configuration>
			<groups operation="delete">
				<name>%s</name>
			</groups>
			<apply-groups operation="delete">%s</apply-groups>
		</configuration>
	</config>
</edit-config>`

const commitStr = `<commit/>`

const getGroupXMLStr = `<get-configuration>
	<configuration>
	<groups><name>%s</name></groups>
	</configuration>
</get-configuration>`

const ApplyGroupXML = `<load-configuration action="merge" format="xml">
%s
</load-configuration>`

type configuration struct {
	ApplyGroup []string `xml:"apply-groups"`
}

var mockMapMutex sync.Mutex

// ProviderConfig is to hold client information
type ProviderConfig struct {
	netconf.Client
	Host string
}

func init() {
	schema.DescriptionKind = schema.StringMarkdown
}

func check(ctx context.Context, err error) {
	if err != nil {
		// Some of these errors will be "normal".
		tflog.Debug(ctx, err.Error())
		f, _ := os.OpenFile("jtaf_logging.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		f.WriteString(err.Error() + "\n")
		f.Close()
		return
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	config := Config{
		Host:     d.Get("host").(string),
		Port:     d.Get("port").(int),
		Username: d.Get("username").(string),
		Password: d.Get("password").(string),
		SSHKey:   d.Get("sshkey").(string),
	}

	configFilePath, ok := os.LookupEnv("MOCK_FILE")
	var client netconf.Client
	var err error

	if ok {
		filePtr, err := os.OpenFile(configFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return nil, diag.FromErr(err)
		}
		client = FileClient{filePtr: filePtr}

	} else {
		client, err = config.Client()
		if err != nil {
			return nil, diag.FromErr(err)
		}
	}

	return &ProviderConfig{client, config.Host}, nil
}

var _ netconf.Client = &FileClient{}

// FileClient represents a fake client for testing purposes.
type FileClient struct {
	// You can add fields for testing purposes here.
	filePtr *os.File
}

// Close is a functional thing to close the FileClient (no-op in this case).
func (bc FileClient) Close() error {
	return nil
}

// updateRawConfig simulates updating the configuration on a network device.
func (bc FileClient) updateRawConfig(applyGroup string, netconfCall string, commit bool) (string, error) {
	// Simulate the update operation (you can customize this part).
	// Extract the string between <name> tags
	nameStart := s.Index(netconfCall, "<name>")
	nameEnd := s.Index(netconfCall, "</name>")
	if nameStart == -1 || nameEnd == -1 {
		return "", fmt.Errorf("Failed to extract the group name from the netconfCall")
	}
	groupName := netconfCall[nameStart+6 : nameEnd]

	// Add the groupName to the applyGroupsList
	addToApplyGroupsList(groupName)

	var groupString string
	groupString = fmt.Sprintf(groupStrXML, netconfCall)
	_, err := bc.filePtr.WriteString(groupString)
	if err != nil {
		return "", err
	}
	bc.filePtr.WriteString("\n\n")
	if commit {
		bc.filePtr.WriteString("\nCommiting from Update\n")
		_, err := bc.filePtr.WriteString(commitStr)
		if err != nil {
			return "", err
		}
	}

	return fmt.Sprintf("Updated config for group: %s", applyGroup), nil
}

// DeleteConfig simulates deleting a configuration on a network device.
func (bc FileClient) DeleteConfig(applyGroup string, commit bool) (string, error) {
	// Simulate the delete operation (you can customize this part).
	return fmt.Sprintf("Deleted config for group: %s", applyGroup), nil
}

// SendCommit simulates sending a commit to a network device.
func (bc FileClient) SendCommit() error {
	// Simulate the commit operation (you can customize this part).
	bc.sortApplyGroupsList()
	if err := bc.SendApplyGroups(); err != nil {
		return err
	}
	bc.filePtr.WriteString("\nCommiting from the SendCommit function\n")
	return nil
}

// MarshalGroup simulates retrieving and marshaling configuration data for a group.
func (bc FileClient) MarshalGroup(id string, obj interface{}) error {
	// Simulate the retrieval and marshaling of configuration data (you can customize this part).
	// For testing purposes, let's just marshal an example object and save it to a file.
	return nil
}

// SendTransaction simulates sending a transaction to a network device.
func (bc FileClient) SendTransaction(id string, obj interface{}, commit bool) error {
	// Simulate sending a transaction (you can customize this part).
	// For testing purposes, let's just write the transaction data to a file.
	cfg, err := xml.Marshal(obj) // Indent with four spaces
	if err != nil {
		return err
	}
	mockMapMutex.Lock()
	defer mockMapMutex.Unlock()

	// updateRawConfig deletes old group by, re-creates it then commits.
	// As far as Junos cares, it's an edit.
	if id != "" {
		bc.filePtr.WriteString("Sending groups to device via Update Function:\n")
		if _, err = bc.updateRawConfig(id, string(cfg), commit); err != nil {
			return err
		}
		return nil
	}
	bc.filePtr.WriteString("Sending groups to device via Send Raw Function:\n")
	if _, err = bc.sendRawConfig(string(cfg), commit); err != nil {
		return err
	}
	return nil
}

// sendRawConfig is a wrapper for driver.SendRaw()
func (bc FileClient) sendRawConfig(netconfCall string, commit bool) (string, error) {

	// Extract the string between <name> tags
	nameStart := s.Index(netconfCall, "<name>")
	nameEnd := s.Index(netconfCall, "</name>")
	if nameStart == -1 || nameEnd == -1 {
		return "", fmt.Errorf("Failed to extract the group name from the netconfCall")
	}
	groupName := netconfCall[nameStart+6 : nameEnd]

	// Add the groupName to the applyGroupsList
	addToApplyGroupsList(groupName)

	groupString := fmt.Sprintf(groupStrXML, netconfCall)

	_, err := bc.filePtr.WriteString(groupString)
	if err != nil {
		return "", err
	}
	bc.filePtr.WriteString("\n\n")
	if commit {
		bc.filePtr.WriteString("\nCommiting from Sending\n")
		_, err := bc.filePtr.WriteString(commitStr)
		if err != nil {
			return "", err
		}
	}

	return "", nil
}

// Helper function to add an id to the global list.
func addToApplyGroupsList(id string) {
	applyGroupsMutex.Lock()
	defer applyGroupsMutex.Unlock()
	applyGroupsList = append(applyGroupsList, id)
}

// Helper function to sort the global list.
func (bc FileClient) sortApplyGroupsList() {
	applyGroupsMutex.Lock()
	defer applyGroupsMutex.Unlock()

	// Create a map to track unique items
	uniqueGroups := make(map[string]bool)

	// Filter out empty s and remove duplicates
	filteredGroups := make([]string, 0)
	for _, group := range applyGroupsList {
		if group != "" && !uniqueGroups[group] {
			uniqueGroups[group] = true
			filteredGroups = append(filteredGroups, group)
		}
	}

	// Sort the filtered list
	sort.Strings(filteredGroups)

	// Update the global applyGroupsList with the sorted and filtered list
	applyGroupsList = filteredGroups
}

var applyGroupsList []string
var applyGroupsMutex sync.Mutex

func (bc FileClient) SendApplyGroups() error {

	// Concatenate the s in applyGroupsList.
	applyGroupsMutex.Lock()
	defer applyGroupsMutex.Unlock()

	var applyG configuration
	applyG.ApplyGroup = make([]string, len(applyGroupsList))
	for i, item := range applyGroupsList {
		applyG.ApplyGroup[i] = item
	}

	cfg, err := xml.Marshal(applyG)
	if err != nil {
		return err
	}

	_, err = bc.filePtr.WriteString("\n")
	if err != nil {
		return err
	}

	_, err = bc.filePtr.WriteString("Sending Apply-Groups to device\n")
	if err != nil {
		return err
	}

	applyGroupString := fmt.Sprintf(ApplyGroupXML, string(cfg))

	_, err = bc.filePtr.WriteString(applyGroupString)
	if err != nil {
		fmt.Printf("Error writing to XML file: %v\n", err)
		return err
	}

	return nil
}

// Provider returns a Terraform Provider.
func Provider() *schema.Provider {
	return &schema.Provider{

		Schema: map[string]*schema.Schema{
			"host": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},

			"username": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"password": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"sshkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},

		ResourcesMap: map[string]*schema.Resource{
				"junos-vsrx_SystemSyslogUser": junosSystemSyslogUser(),
			"junos-vsrx_ProtocolsBgpGroupFamilyEvpnSignaling": junosProtocolsBgpGroupFamilyEvpnSignaling(),
			"junos-vsrx_Policy__OptionsPolicy__StatementTermThenCommunity": junosPolicy__OptionsPolicy__StatementTermThenCommunity(),
			"junos-vsrx_ProtocolsBgpGroupName": junosProtocolsBgpGroupName(),
			"junos-vsrx_ProtocolsBgpGroupBfd__Liveness__DetectionMultiplier": junosProtocolsBgpGroupBfd__Liveness__DetectionMultiplier(),
			"junos-vsrx_Routing__OptionsStatic": junosRouting__OptionsStatic(),
			"junos-vsrx_ProtocolsBgpGroupMultipathMultiple__As": junosProtocolsBgpGroupMultipathMultiple__As(),
			"junos-vsrx_SystemServicesExtension__ServiceNotification": junosSystemServicesExtension__ServiceNotification(),
			"junos-vsrx_SystemServicesRestHttp": junosSystemServicesRestHttp(),
			"junos-vsrx_SnmpContact": junosSnmpContact(),
			"junos-vsrx_Policy__OptionsCommunityMembers": junosPolicy__OptionsCommunityMembers(),
			"junos-vsrx_SystemLoginUser": junosSystemLoginUser(),
			"junos-vsrx_InterfacesInterfaceUnitName": junosInterfacesInterfaceUnitName(),
			"junos-vsrx_ProtocolsBgpGroupFamilyEvpn": junosProtocolsBgpGroupFamilyEvpn(),
			"junos-vsrx_ProtocolsBgpGroupMtu__Discovery": junosProtocolsBgpGroupMtu__Discovery(),
			"junos-vsrx_SystemSyslogFileContents": junosSystemSyslogFileContents(),
			"junos-vsrx_ChassisAggregated__DevicesEthernet": junosChassisAggregated__DevicesEthernet(),
			"junos-vsrx_ProtocolsBgpGroupType": junosProtocolsBgpGroupType(),
			"junos-vsrx_ProtocolsLldp": junosProtocolsLldp(),
			"junos-vsrx_SystemSyslogUserContentsName": junosSystemSyslogUserContentsName(),
			"junos-vsrx_Routing__OptionsStaticRoute": junosRouting__OptionsStaticRoute(),
			"junos-vsrx_Routing__OptionsStaticRouteName": junosRouting__OptionsStaticRouteName(),
			"junos-vsrx_SystemSyslogFile": junosSystemSyslogFile(),
			"junos-vsrx_SystemSyslogFileContentsNotice": junosSystemSyslogFileContentsNotice(),
			"junos-vsrx_SystemSyslogFileContentsInfo": junosSystemSyslogFileContentsInfo(),
			"junos-vsrx_SystemExtensionsProvidersLicense__TypeName": junosSystemExtensionsProvidersLicense__TypeName(),
			"junos-vsrx_ProtocolsBgpGroupLocal__AsAs__Number": junosProtocolsBgpGroupLocal__AsAs__Number(),
			"junos-vsrx_ProtocolsIgmp__SnoopingVlanName": junosProtocolsIgmp__SnoopingVlanName(),
			"junos-vsrx_SystemLoginMessage": junosSystemLoginMessage(),
			"junos-vsrx_SystemExtensionsProvidersName": junosSystemExtensionsProvidersName(),
			"junos-vsrx_InterfacesInterfaceUnit": junosInterfacesInterfaceUnit(),
			"junos-vsrx_InterfacesInterfaceUnitFamilyInetAddressName": junosInterfacesInterfaceUnitFamilyInetAddressName(),
			"junos-vsrx_ProtocolsBgpGroupNeighborName": junosProtocolsBgpGroupNeighborName(),
			"junos-vsrx_ProtocolsBgpGroupExport": junosProtocolsBgpGroupExport(),
			"junos-vsrx_SystemExtensionsProvidersLicense__TypeDeployment__Scope": junosSystemExtensionsProvidersLicense__TypeDeployment__Scope(),
			"junos-vsrx_SnmpCommunityName": junosSnmpCommunityName(),
			"junos-vsrx_ProtocolsBgpGroupMultipath": junosProtocolsBgpGroupMultipath(),
			"junos-vsrx_ProtocolsBgpGroupVpn__Apply__Export": junosProtocolsBgpGroupVpn__Apply__Export(),
			"junos-vsrx_SystemExtensionsProviders": junosSystemExtensionsProviders(),
			"junos-vsrx_InterfacesInterfaceUnitFamily": junosInterfacesInterfaceUnitFamily(),
			"junos-vsrx_Forwarding__OptionsStorm__Control__ProfilesAll": junosForwarding__OptionsStorm__Control__ProfilesAll(),
			"junos-vsrx_Routing__OptionsRouter__Id": junosRouting__OptionsRouter__Id(),
			"junos-vsrx_InterfacesInterfaceUnitFamilyInetAddress": junosInterfacesInterfaceUnitFamilyInetAddress(),
			"junos-vsrx_Policy__OptionsPolicy__Statement": junosPolicy__OptionsPolicy__Statement(),
			"junos-vsrx_SystemServices": junosSystemServices(),
			"junos-vsrx_SystemServicesNetconfSsh": junosSystemServicesNetconfSsh(),
			"junos-vsrx_InterfacesInterfaceName": junosInterfacesInterfaceName(),
			"junos-vsrx_ProtocolsBgpGroupNeighborDescription": junosProtocolsBgpGroupNeighborDescription(),
			"junos-vsrx_SystemLoginUserAuthenticationEncrypted__Password": junosSystemLoginUserAuthenticationEncrypted__Password(),
			"junos-vsrx_SystemServicesExtension__ServiceNotificationAllow__ClientsAddress": junosSystemServicesExtension__ServiceNotificationAllow__ClientsAddress(),
			"junos-vsrx_SystemServicesRest": junosSystemServicesRest(),
			"junos-vsrx_ChassisAggregated__Devices": junosChassisAggregated__Devices(),
			"junos-vsrx_SnmpLocation": junosSnmpLocation(),
			"junos-vsrx_Policy__OptionsPolicy__StatementTermThenCommunityCommunity__Name": junosPolicy__OptionsPolicy__StatementTermThenCommunityCommunity__Name(),
			"junos-vsrx_SystemLoginUserName": junosSystemLoginUserName(),
			"junos-vsrx_SystemLoginUserAuthentication": junosSystemLoginUserAuthentication(),
			"junos-vsrx_SystemSyslogFileName": junosSystemSyslogFileName(),
			"junos-vsrx_Forwarding__OptionsStorm__Control__ProfilesName": junosForwarding__OptionsStorm__Control__ProfilesName(),
			"junos-vsrx_Policy__OptionsPolicy__StatementThenLoad__Balance": junosPolicy__OptionsPolicy__StatementThenLoad__Balance(),
			"junos-vsrx_SystemServicesExtension__ServiceNotificationAllow__Clients": junosSystemServicesExtension__ServiceNotificationAllow__Clients(),
			"junos-vsrx_InterfacesInterface": junosInterfacesInterface(),
			"junos-vsrx_InterfacesInterfaceUnitDescription": junosInterfacesInterfaceUnitDescription(),
			"junos-vsrx_ProtocolsBgpGroupBfd__Liveness__Detection": junosProtocolsBgpGroupBfd__Liveness__Detection(),
			"junos-vsrx_Policy__OptionsPolicy__StatementName": junosPolicy__OptionsPolicy__StatementName(),
			"junos-vsrx_Policy__OptionsPolicy__StatementTermFrom": junosPolicy__OptionsPolicy__StatementTermFrom(),
			"junos-vsrx_SystemLoginUserClass": junosSystemLoginUserClass(),
			"junos-vsrx_InterfacesInterfaceUnitFamilyInet": junosInterfacesInterfaceUnitFamilyInet(),
			"junos-vsrx_Routing__OptionsForwarding__Table": junosRouting__OptionsForwarding__Table(),
			"junos-vsrx_ProtocolsBgpGroupLocal__As": junosProtocolsBgpGroupLocal__As(),
			"junos-vsrx_Policy__OptionsPolicy__StatementThen": junosPolicy__OptionsPolicy__StatementThen(),
			"junos-vsrx_SystemServicesSshRoot__Login": junosSystemServicesSshRoot__Login(),
			"junos-vsrx_SystemServicesExtension__ServiceRequest__ResponseGrpcMax__Connections": junosSystemServicesExtension__ServiceRequest__ResponseGrpcMax__Connections(),
			"junos-vsrx_SystemServicesRestEnable__Explorer": junosSystemServicesRestEnable__Explorer(),
			"junos-vsrx_SystemSyslogFileContentsAny": junosSystemSyslogFileContentsAny(),
			"junos-vsrx_ChassisAggregated__DevicesEthernetDevice__Count": junosChassisAggregated__DevicesEthernetDevice__Count(),
			"junos-vsrx_Policy__OptionsPolicy__StatementTermFromProtocol": junosPolicy__OptionsPolicy__StatementTermFromProtocol(),
			"junos-vsrx_SystemHost__Name": junosSystemHost__Name(),
			"junos-vsrx_SystemServicesExtension__Service": junosSystemServicesExtension__Service(),
			"junos-vsrx_ProtocolsBgpGroupNeighborPeer__As": junosProtocolsBgpGroupNeighborPeer__As(),
			"junos-vsrx_SystemSyslogUserName": junosSystemSyslogUserName(),
			"junos-vsrx_Routing__OptionsForwarding__TableExport": junosRouting__OptionsForwarding__TableExport(),
			"junos-vsrx_ProtocolsBgp": junosProtocolsBgp(),
			"junos-vsrx_ProtocolsBgpGroupCluster": junosProtocolsBgpGroupCluster(),
			"junos-vsrx_Policy__OptionsPolicy__StatementTermThen": junosPolicy__OptionsPolicy__StatementTermThen(),
			"junos-vsrx_Policy__OptionsPolicy__StatementTermThenAccept": junosPolicy__OptionsPolicy__StatementTermThenAccept(),
			"junos-vsrx_SystemServicesExtension__ServiceRequest__Response": junosSystemServicesExtension__ServiceRequest__Response(),
			"junos-vsrx_Policy__OptionsPolicy__StatementTerm": junosPolicy__OptionsPolicy__StatementTerm(),
			"junos-vsrx_Policy__OptionsCommunity": junosPolicy__OptionsCommunity(),
			"junos-vsrx_SnmpCommunity": junosSnmpCommunity(),
			"junos-vsrx_ProtocolsBgpGroupImport": junosProtocolsBgpGroupImport(),
			"junos-vsrx_ProtocolsIgmp__Snooping": junosProtocolsIgmp__Snooping(),
			"junos-vsrx_SystemServicesSsh": junosSystemServicesSsh(),
			"junos-vsrx_SystemServicesExtension__ServiceRequest__ResponseGrpc": junosSystemServicesExtension__ServiceRequest__ResponseGrpc(),
			"junos-vsrx_SystemSyslogUserContents": junosSystemSyslogUserContents(),
			"junos-vsrx_SnmpCommunityAuthorization": junosSnmpCommunityAuthorization(),
			"junos-vsrx_ProtocolsBgpGroupLocal__Address": junosProtocolsBgpGroupLocal__Address(),
			"junos-vsrx_ProtocolsBgpGroupFamily": junosProtocolsBgpGroupFamily(),
			"junos-vsrx_SystemLogin": junosSystemLogin(),
			"junos-vsrx_SystemExtensionsProvidersLicense__Type": junosSystemExtensionsProvidersLicense__Type(),
			"junos-vsrx_ProtocolsLldpInterfaceName": junosProtocolsLldpInterfaceName(),
			"junos-vsrx_Policy__OptionsPolicy__StatementThenLoad__BalancePer__Packet": junosPolicy__OptionsPolicy__StatementThenLoad__BalancePer__Packet(),
			"junos-vsrx_SystemSyslogUserContentsEmergency": junosSystemSyslogUserContentsEmergency(),
			"junos-vsrx_Policy__OptionsCommunityName": junosPolicy__OptionsCommunityName(),
			"junos-vsrx_SystemSyslog": junosSystemSyslog(),
			"junos-vsrx_InterfacesInterfaceDescription": junosInterfacesInterfaceDescription(),
			"junos-vsrx_Forwarding__OptionsStorm__Control__Profiles": junosForwarding__OptionsStorm__Control__Profiles(),
			"junos-vsrx_Routing__OptionsStaticRouteNext__Hop": junosRouting__OptionsStaticRouteNext__Hop(),
			"junos-vsrx_ProtocolsBgpGroupNeighbor": junosProtocolsBgpGroupNeighbor(),
			"junos-vsrx_ProtocolsLldpInterface": junosProtocolsLldpInterface(),
			"junos-vsrx_SystemRoot__Authentication": junosSystemRoot__Authentication(),
			"junos-vsrx_SystemRoot__AuthenticationEncrypted__Password": junosSystemRoot__AuthenticationEncrypted__Password(),
			"junos-vsrx_SystemExtensions": junosSystemExtensions(),
			"junos-vsrx_SystemServicesNetconf": junosSystemServicesNetconf(),
			"junos-vsrx_SystemSyslogFileContentsName": junosSystemSyslogFileContentsName(),
			"junos-vsrx_ProtocolsBgpGroup": junosProtocolsBgpGroup(),
			"junos-vsrx_ProtocolsIgmp__SnoopingVlan": junosProtocolsIgmp__SnoopingVlan(),
			"junos-vsrx_Policy__OptionsPolicy__StatementTermThenReject": junosPolicy__OptionsPolicy__StatementTermThenReject(),
			"junos-vsrx_SystemServicesRestHttpPort": junosSystemServicesRestHttpPort(),
			"junos-vsrx_ProtocolsBgpGroupBfd__Liveness__DetectionMinimum__Interval": junosProtocolsBgpGroupBfd__Liveness__DetectionMinimum__Interval(),
			"junos-vsrx_SystemLoginUserUid": junosSystemLoginUserUid(),
			"junos-vsrx_ProtocolsBgpGroupAllow": junosProtocolsBgpGroupAllow(),
			"junos-vsrx_Policy__OptionsPolicy__StatementTermName": junosPolicy__OptionsPolicy__StatementTermName(),
			"junos-vsrx_commit": junosCommit(),
	        "junos-vsrx_destroycommit": junosDestroyCommit(),
			},
		    ConfigureContextFunc: providerConfigure,
	    } 
    }