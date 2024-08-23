
terraform {
	required_providers {
		junos-vsrx = {
			source = "{input source path}"
			version = "{input version here}"
		}
	}
}

resource "junos-vsrx_SystemSyslogUser" "vsrx_1" {
	resource_name = "vsrx_1"
}

resource "junos-vsrx_ProtocolsBgpGroupFamilyEvpnSignaling" "vsrx_2" {
	resource_name = "vsrx_2"
	name = "/protocols/bgp/group/name"
}

resource "junos-vsrx_Policy__OptionsPolicy__StatementTermThenCommunity" "vsrx_3" {
	resource_name = "vsrx_3"
	name = "/policy-options/policy-statement/name"
	name__1 = "/policy-options/policy-statement/term/name"
}

resource "junos-vsrx_ProtocolsBgpGroupName" "vsrx_4" {
	resource_name = "vsrx_4"
	name = "/protocols/bgp/group/name"
}

resource "junos-vsrx_ProtocolsBgpGroupBfd__Liveness__DetectionMultiplier" "vsrx_5" {
	resource_name = "vsrx_5"
	name = "/protocols/bgp/group/name"
	multiplier = "/protocols/bgp/group/bfd-liveness-detection/multiplier"
}

resource "junos-vsrx_Routing__OptionsStatic" "vsrx_6" {
	resource_name = "vsrx_6"
}

resource "junos-vsrx_ProtocolsBgpGroupMultipathMultiple__As" "vsrx_7" {
	resource_name = "vsrx_7"
	name = "/protocols/bgp/group/name"
	multiple__as = "/protocols/bgp/group/multipath/multiple-as"
}

resource "junos-vsrx_SystemServicesExtension__ServiceNotification" "vsrx_8" {
	resource_name = "vsrx_8"
}

resource "junos-vsrx_SystemServicesRestHttp" "vsrx_9" {
	resource_name = "vsrx_9"
}

resource "junos-vsrx_SnmpContact" "vsrx_10" {
	resource_name = "vsrx_10"
	contact = "/contact"
}

resource "junos-vsrx_Policy__OptionsCommunityMembers" "vsrx_11" {
	resource_name = "vsrx_11"
	name = "/policy-options/community/name"
	members = "/policy-options/community/members"
}

resource "junos-vsrx_SystemLoginUser" "vsrx_12" {
	resource_name = "vsrx_12"
}

resource "junos-vsrx_InterfacesInterfaceUnitName" "vsrx_13" {
	resource_name = "vsrx_13"
	name = "/interfaces/interface/name"
	name__1 = "/interfaces/interface/unit/name"
}

resource "junos-vsrx_ProtocolsBgpGroupFamilyEvpn" "vsrx_14" {
	resource_name = "vsrx_14"
	name = "/protocols/bgp/group/name"
}

resource "junos-vsrx_ProtocolsBgpGroupMtu__Discovery" "vsrx_15" {
	resource_name = "vsrx_15"
	name = "/protocols/bgp/group/name"
	mtu__discovery = "/protocols/bgp/group/mtu-discovery"
}

resource "junos-vsrx_SystemSyslogFileContents" "vsrx_16" {
	resource_name = "vsrx_16"
	name = "/system/syslog/file/name"
}

resource "junos-vsrx_ChassisAggregated__DevicesEthernet" "vsrx_17" {
	resource_name = "vsrx_17"
}

resource "junos-vsrx_ProtocolsBgpGroupType" "vsrx_18" {
	resource_name = "vsrx_18"
	name = "/protocols/bgp/group/name"
	type = "/protocols/bgp/group/type"
}

resource "junos-vsrx_ProtocolsLldp" "vsrx_19" {
	resource_name = "vsrx_19"
}

resource "junos-vsrx_SystemSyslogUserContentsName" "vsrx_20" {
	resource_name = "vsrx_20"
	name = "/system/syslog/user/name"
	name__1 = "/system/syslog/user/contents/name"
}

resource "junos-vsrx_Routing__OptionsStaticRoute" "vsrx_21" {
	resource_name = "vsrx_21"
}

resource "junos-vsrx_Routing__OptionsStaticRouteName" "vsrx_22" {
	resource_name = "vsrx_22"
	name = "/routing-options/static/route/name"
}

resource "junos-vsrx_SystemSyslogFile" "vsrx_23" {
	resource_name = "vsrx_23"
}

resource "junos-vsrx_SystemSyslogFileContentsNotice" "vsrx_24" {
	resource_name = "vsrx_24"
	name = "/system/syslog/file/name"
	name__1 = "/system/syslog/file/contents/name"
	notice = "/system/syslog/file/contents/notice"
}

resource "junos-vsrx_SystemSyslogFileContentsInfo" "vsrx_25" {
	resource_name = "vsrx_25"
	name = "/system/syslog/file/name"
	name__1 = "/system/syslog/file/contents/name"
	info = "/system/syslog/file/contents/info"
}

resource "junos-vsrx_SystemExtensionsProvidersLicense__TypeName" "vsrx_26" {
	resource_name = "vsrx_26"
	name = "/system/extensions/providers/name"
	name__1 = "/system/extensions/providers/license-type/name"
}

resource "junos-vsrx_ProtocolsBgpGroupLocal__AsAs__Number" "vsrx_27" {
	resource_name = "vsrx_27"
	name = "/protocols/bgp/group/name"
	as__number = "/protocols/bgp/group/local-as/as-number"
}

resource "junos-vsrx_ProtocolsIgmp__SnoopingVlanName" "vsrx_28" {
	resource_name = "vsrx_28"
	name = "/protocols/igmp-snooping/vlan/name"
}

resource "junos-vsrx_SystemLoginMessage" "vsrx_29" {
	resource_name = "vsrx_29"
	message = "/system/login/message"
}

resource "junos-vsrx_SystemExtensionsProvidersName" "vsrx_30" {
	resource_name = "vsrx_30"
	name = "/system/extensions/providers/name"
}

resource "junos-vsrx_InterfacesInterfaceUnit" "vsrx_31" {
	resource_name = "vsrx_31"
	name = "/interfaces/interface/name"
}

resource "junos-vsrx_InterfacesInterfaceUnitFamilyInetAddressName" "vsrx_32" {
	resource_name = "vsrx_32"
	name = "/interfaces/interface/name"
	name__1 = "/interfaces/interface/unit/name"
	name__2 = "/interfaces/interface/unit/family/inet/address/name"
}

resource "junos-vsrx_ProtocolsBgpGroupNeighborName" "vsrx_33" {
	resource_name = "vsrx_33"
	name = "/protocols/bgp/group/name"
	name__1 = "/protocols/bgp/group/neighbor/name"
}

resource "junos-vsrx_ProtocolsBgpGroupExport" "vsrx_34" {
	resource_name = "vsrx_34"
	name = "/protocols/bgp/group/name"
	export = "/protocols/bgp/group/export"
}

resource "junos-vsrx_SystemExtensionsProvidersLicense__TypeDeployment__Scope" "vsrx_35" {
	resource_name = "vsrx_35"
	name = "/system/extensions/providers/name"
	name__1 = "/system/extensions/providers/license-type/name"
	deployment__scope = "/system/extensions/providers/license-type/deployment-scope"
}

resource "junos-vsrx_SnmpCommunityName" "vsrx_36" {
	resource_name = "vsrx_36"
	name = "/snmp/community/name"
}

resource "junos-vsrx_ProtocolsBgpGroupMultipath" "vsrx_37" {
	resource_name = "vsrx_37"
	name = "/protocols/bgp/group/name"
}

resource "junos-vsrx_ProtocolsBgpGroupVpn__Apply__Export" "vsrx_38" {
	resource_name = "vsrx_38"
	name = "/protocols/bgp/group/name"
	vpn__apply__export = "/protocols/bgp/group/vpn-apply-export"
}

resource "junos-vsrx_SystemExtensionsProviders" "vsrx_39" {
	resource_name = "vsrx_39"
}

resource "junos-vsrx_InterfacesInterfaceUnitFamily" "vsrx_40" {
	resource_name = "vsrx_40"
	name = "/interfaces/interface/name"
	name__1 = "/interfaces/interface/unit/name"
}

resource "junos-vsrx_Forwarding__OptionsStorm__Control__ProfilesAll" "vsrx_41" {
	resource_name = "vsrx_41"
	name = "/forwarding-options/storm-control-profiles/name"
}

resource "junos-vsrx_Routing__OptionsRouter__Id" "vsrx_42" {
	resource_name = "vsrx_42"
	router__id = "/router-id"
}

resource "junos-vsrx_InterfacesInterfaceUnitFamilyInetAddress" "vsrx_43" {
	resource_name = "vsrx_43"
	name = "/interfaces/interface/name"
	name__1 = "/interfaces/interface/unit/name"
}

resource "junos-vsrx_Policy__OptionsPolicy__Statement" "vsrx_44" {
	resource_name = "vsrx_44"
}

resource "junos-vsrx_SystemServices" "vsrx_45" {
	resource_name = "vsrx_45"
}

resource "junos-vsrx_SystemServicesNetconfSsh" "vsrx_46" {
	resource_name = "vsrx_46"
}

resource "junos-vsrx_InterfacesInterfaceName" "vsrx_47" {
	resource_name = "vsrx_47"
	name = "/interfaces/interface/name"
}

resource "junos-vsrx_ProtocolsBgpGroupNeighborDescription" "vsrx_48" {
	resource_name = "vsrx_48"
	name = "/protocols/bgp/group/name"
	name__1 = "/protocols/bgp/group/neighbor/name"
	description = "/protocols/bgp/group/neighbor/description"
}

resource "junos-vsrx_SystemLoginUserAuthenticationEncrypted__Password" "vsrx_49" {
	resource_name = "vsrx_49"
	name = "/system/login/user/name"
	encrypted__password = "/system/login/user/authentication/encrypted-password"
}

resource "junos-vsrx_SystemServicesExtension__ServiceNotificationAllow__ClientsAddress" "vsrx_50" {
	resource_name = "vsrx_50"
	address = "/system/services/extension-service/notification/allow-clients/address"
}

resource "junos-vsrx_SystemServicesRest" "vsrx_51" {
	resource_name = "vsrx_51"
}

resource "junos-vsrx_ChassisAggregated__Devices" "vsrx_52" {
	resource_name = "vsrx_52"
}

resource "junos-vsrx_SnmpLocation" "vsrx_53" {
	resource_name = "vsrx_53"
	location = "/location"
}

resource "junos-vsrx_Policy__OptionsPolicy__StatementTermThenCommunityCommunity__Name" "vsrx_54" {
	resource_name = "vsrx_54"
	name = "/policy-options/policy-statement/name"
	name__1 = "/policy-options/policy-statement/term/name"
	community__name = "/policy-options/policy-statement/term/then/community/community-name"
}

resource "junos-vsrx_SystemLoginUserName" "vsrx_55" {
	resource_name = "vsrx_55"
	name = "/system/login/user/name"
}

resource "junos-vsrx_SystemLoginUserAuthentication" "vsrx_56" {
	resource_name = "vsrx_56"
	name = "/system/login/user/name"
}

resource "junos-vsrx_SystemSyslogFileName" "vsrx_57" {
	resource_name = "vsrx_57"
	name = "/system/syslog/file/name"
}

resource "junos-vsrx_Forwarding__OptionsStorm__Control__ProfilesName" "vsrx_58" {
	resource_name = "vsrx_58"
	name = "/forwarding-options/storm-control-profiles/name"
}

resource "junos-vsrx_Policy__OptionsPolicy__StatementThenLoad__Balance" "vsrx_59" {
	resource_name = "vsrx_59"
	name = "/policy-options/policy-statement/name"
}

resource "junos-vsrx_SystemServicesExtension__ServiceNotificationAllow__Clients" "vsrx_60" {
	resource_name = "vsrx_60"
}

resource "junos-vsrx_InterfacesInterface" "vsrx_61" {
	resource_name = "vsrx_61"
}

resource "junos-vsrx_InterfacesInterfaceUnitDescription" "vsrx_62" {
	resource_name = "vsrx_62"
	name = "/interfaces/interface/name"
	name__1 = "/interfaces/interface/unit/name"
	description = "/interfaces/interface/unit/description"
}

resource "junos-vsrx_ProtocolsBgpGroupBfd__Liveness__Detection" "vsrx_63" {
	resource_name = "vsrx_63"
	name = "/protocols/bgp/group/name"
}

resource "junos-vsrx_Policy__OptionsPolicy__StatementName" "vsrx_64" {
	resource_name = "vsrx_64"
	name = "/policy-options/policy-statement/name"
}

resource "junos-vsrx_Policy__OptionsPolicy__StatementTermFrom" "vsrx_65" {
	resource_name = "vsrx_65"
	name = "/policy-options/policy-statement/name"
	name__1 = "/policy-options/policy-statement/term/name"
}

resource "junos-vsrx_SystemLoginUserClass" "vsrx_66" {
	resource_name = "vsrx_66"
	name = "/system/login/user/name"
	class = "/system/login/user/class"
}

resource "junos-vsrx_InterfacesInterfaceUnitFamilyInet" "vsrx_67" {
	resource_name = "vsrx_67"
	name = "/interfaces/interface/name"
	name__1 = "/interfaces/interface/unit/name"
}

resource "junos-vsrx_Routing__OptionsForwarding__Table" "vsrx_68" {
	resource_name = "vsrx_68"
}

resource "junos-vsrx_ProtocolsBgpGroupLocal__As" "vsrx_69" {
	resource_name = "vsrx_69"
	name = "/protocols/bgp/group/name"
}

resource "junos-vsrx_Policy__OptionsPolicy__StatementThen" "vsrx_70" {
	resource_name = "vsrx_70"
	name = "/policy-options/policy-statement/name"
}

resource "junos-vsrx_SystemServicesSshRoot__Login" "vsrx_71" {
	resource_name = "vsrx_71"
	root__login = "/system/services/ssh/root-login"
}

resource "junos-vsrx_SystemServicesExtension__ServiceRequest__ResponseGrpcMax__Connections" "vsrx_72" {
	resource_name = "vsrx_72"
	max__connections = "/system/services/extension-service/request-response/grpc/max-connections"
}

resource "junos-vsrx_SystemServicesRestEnable__Explorer" "vsrx_73" {
	resource_name = "vsrx_73"
	enable__explorer = "/system/services/rest/enable-explorer"
}

resource "junos-vsrx_SystemSyslogFileContentsAny" "vsrx_74" {
	resource_name = "vsrx_74"
	name = "/system/syslog/file/name"
	name__1 = "/system/syslog/file/contents/name"
	any = "/system/syslog/file/contents/any"
}

resource "junos-vsrx_ChassisAggregated__DevicesEthernetDevice__Count" "vsrx_75" {
	resource_name = "vsrx_75"
	device__count = "/chassis/aggregated-devices/ethernet/device-count"
}

resource "junos-vsrx_Policy__OptionsPolicy__StatementTermFromProtocol" "vsrx_76" {
	resource_name = "vsrx_76"
	name = "/policy-options/policy-statement/name"
	name__1 = "/policy-options/policy-statement/term/name"
	protocol = "/policy-options/policy-statement/term/from/protocol"
}

resource "junos-vsrx_SystemHost__Name" "vsrx_77" {
	resource_name = "vsrx_77"
	host__name = "/host-name"
}

resource "junos-vsrx_SystemServicesExtension__Service" "vsrx_78" {
	resource_name = "vsrx_78"
}

resource "junos-vsrx_ProtocolsBgpGroupNeighborPeer__As" "vsrx_79" {
	resource_name = "vsrx_79"
	name = "/protocols/bgp/group/name"
	name__1 = "/protocols/bgp/group/neighbor/name"
	peer__as = "/protocols/bgp/group/neighbor/peer-as"
}

resource "junos-vsrx_SystemSyslogUserName" "vsrx_80" {
	resource_name = "vsrx_80"
	name = "/system/syslog/user/name"
}

resource "junos-vsrx_Routing__OptionsForwarding__TableExport" "vsrx_81" {
	resource_name = "vsrx_81"
	export = "/routing-options/forwarding-table/export"
}

resource "junos-vsrx_ProtocolsBgp" "vsrx_82" {
	resource_name = "vsrx_82"
}

resource "junos-vsrx_ProtocolsBgpGroupCluster" "vsrx_83" {
	resource_name = "vsrx_83"
	name = "/protocols/bgp/group/name"
	cluster = "/protocols/bgp/group/cluster"
}

resource "junos-vsrx_Policy__OptionsPolicy__StatementTermThen" "vsrx_84" {
	resource_name = "vsrx_84"
	name = "/policy-options/policy-statement/name"
	name__1 = "/policy-options/policy-statement/term/name"
}

resource "junos-vsrx_Policy__OptionsPolicy__StatementTermThenAccept" "vsrx_85" {
	resource_name = "vsrx_85"
	name = "/policy-options/policy-statement/name"
	name__1 = "/policy-options/policy-statement/term/name"
	accept = "/policy-options/policy-statement/term/then/accept"
}

resource "junos-vsrx_SystemServicesExtension__ServiceRequest__Response" "vsrx_86" {
	resource_name = "vsrx_86"
}

resource "junos-vsrx_Policy__OptionsPolicy__StatementTerm" "vsrx_87" {
	resource_name = "vsrx_87"
	name = "/policy-options/policy-statement/name"
}

resource "junos-vsrx_Policy__OptionsCommunity" "vsrx_88" {
	resource_name = "vsrx_88"
}

resource "junos-vsrx_SnmpCommunity" "vsrx_89" {
	resource_name = "vsrx_89"
}

resource "junos-vsrx_ProtocolsBgpGroupImport" "vsrx_90" {
	resource_name = "vsrx_90"
	name = "/protocols/bgp/group/name"
	import = "/protocols/bgp/group/import"
}

resource "junos-vsrx_ProtocolsIgmp__Snooping" "vsrx_91" {
	resource_name = "vsrx_91"
}

resource "junos-vsrx_SystemServicesSsh" "vsrx_92" {
	resource_name = "vsrx_92"
}

resource "junos-vsrx_SystemServicesExtension__ServiceRequest__ResponseGrpc" "vsrx_93" {
	resource_name = "vsrx_93"
}

resource "junos-vsrx_SystemSyslogUserContents" "vsrx_94" {
	resource_name = "vsrx_94"
	name = "/system/syslog/user/name"
}

resource "junos-vsrx_SnmpCommunityAuthorization" "vsrx_95" {
	resource_name = "vsrx_95"
	name = "/snmp/community/name"
	authorization = "/snmp/community/authorization"
}

resource "junos-vsrx_ProtocolsBgpGroupLocal__Address" "vsrx_96" {
	resource_name = "vsrx_96"
	name = "/protocols/bgp/group/name"
	local__address = "/protocols/bgp/group/local-address"
}

resource "junos-vsrx_ProtocolsBgpGroupFamily" "vsrx_97" {
	resource_name = "vsrx_97"
	name = "/protocols/bgp/group/name"
}

resource "junos-vsrx_SystemLogin" "vsrx_98" {
	resource_name = "vsrx_98"
}

resource "junos-vsrx_SystemExtensionsProvidersLicense__Type" "vsrx_99" {
	resource_name = "vsrx_99"
	name = "/system/extensions/providers/name"
}

resource "junos-vsrx_ProtocolsLldpInterfaceName" "vsrx_100" {
	resource_name = "vsrx_100"
	name = "/protocols/lldp/interface/name"
}

resource "junos-vsrx_Policy__OptionsPolicy__StatementThenLoad__BalancePer__Packet" "vsrx_101" {
	resource_name = "vsrx_101"
	name = "/policy-options/policy-statement/name"
	per__packet = "/policy-options/policy-statement/then/load-balance/per-packet"
}

resource "junos-vsrx_SystemSyslogUserContentsEmergency" "vsrx_102" {
	resource_name = "vsrx_102"
	name = "/system/syslog/user/name"
	name__1 = "/system/syslog/user/contents/name"
	emergency = "/system/syslog/user/contents/emergency"
}

resource "junos-vsrx_Policy__OptionsCommunityName" "vsrx_103" {
	resource_name = "vsrx_103"
	name = "/policy-options/community/name"
}

resource "junos-vsrx_SystemSyslog" "vsrx_104" {
	resource_name = "vsrx_104"
}

resource "junos-vsrx_InterfacesInterfaceDescription" "vsrx_105" {
	resource_name = "vsrx_105"
	name = "/interfaces/interface/name"
	description = "/interfaces/interface/description"
}

resource "junos-vsrx_Forwarding__OptionsStorm__Control__Profiles" "vsrx_106" {
	resource_name = "vsrx_106"
}

resource "junos-vsrx_Routing__OptionsStaticRouteNext__Hop" "vsrx_107" {
	resource_name = "vsrx_107"
	name = "/routing-options/static/route/name"
	next__hop = "/routing-options/static/route/next-hop"
}

resource "junos-vsrx_ProtocolsBgpGroupNeighbor" "vsrx_108" {
	resource_name = "vsrx_108"
	name = "/protocols/bgp/group/name"
}

resource "junos-vsrx_ProtocolsLldpInterface" "vsrx_109" {
	resource_name = "vsrx_109"
}

resource "junos-vsrx_SystemRoot__Authentication" "vsrx_110" {
	resource_name = "vsrx_110"
}

resource "junos-vsrx_SystemRoot__AuthenticationEncrypted__Password" "vsrx_111" {
	resource_name = "vsrx_111"
	encrypted__password = "/system/root-authentication/encrypted-password"
}

resource "junos-vsrx_SystemExtensions" "vsrx_112" {
	resource_name = "vsrx_112"
}

resource "junos-vsrx_SystemServicesNetconf" "vsrx_113" {
	resource_name = "vsrx_113"
}

resource "junos-vsrx_SystemSyslogFileContentsName" "vsrx_114" {
	resource_name = "vsrx_114"
	name = "/system/syslog/file/name"
	name__1 = "/system/syslog/file/contents/name"
}

resource "junos-vsrx_ProtocolsBgpGroup" "vsrx_115" {
	resource_name = "vsrx_115"
}

resource "junos-vsrx_ProtocolsIgmp__SnoopingVlan" "vsrx_116" {
	resource_name = "vsrx_116"
}

resource "junos-vsrx_Policy__OptionsPolicy__StatementTermThenReject" "vsrx_117" {
	resource_name = "vsrx_117"
	name = "/policy-options/policy-statement/name"
	name__1 = "/policy-options/policy-statement/term/name"
	reject = "/policy-options/policy-statement/term/then/reject"
}

resource "junos-vsrx_SystemServicesRestHttpPort" "vsrx_118" {
	resource_name = "vsrx_118"
	port = "/system/services/rest/http/port"
}

resource "junos-vsrx_ProtocolsBgpGroupBfd__Liveness__DetectionMinimum__Interval" "vsrx_119" {
	resource_name = "vsrx_119"
	name = "/protocols/bgp/group/name"
	minimum__interval = "/protocols/bgp/group/bfd-liveness-detection/minimum-interval"
}

resource "junos-vsrx_SystemLoginUserUid" "vsrx_120" {
	resource_name = "vsrx_120"
	name = "/system/login/user/name"
	uid = "/system/login/user/uid"
}

resource "junos-vsrx_ProtocolsBgpGroupAllow" "vsrx_121" {
	resource_name = "vsrx_121"
	name = "/protocols/bgp/group/name"
	allow = "/protocols/bgp/group/allow"
}

resource "junos-vsrx_Policy__OptionsPolicy__StatementTermName" "vsrx_122" {
	resource_name = "vsrx_122"
	name = "/policy-options/policy-statement/name"
	name__1 = "/policy-options/policy-statement/term/name"
}

