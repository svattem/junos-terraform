
terraform {
	required_providers {
		junos-vsrx = {
			source = "{input source path}"
			version = "{input version here}"
		}
	}
}

resource "junos-vsrx_InterfacesInterfaceDescription" "vsrx_1" {
	resource_name = "vsrx_1"
	name = "/interfaces/interface/name"
	description = "/interfaces/interface/description"
}

resource "junos-vsrx_InterfacesInterfaceUnitFamilyInetAddressName" "vsrx_2" {
	resource_name = "vsrx_2"
	name = "/interfaces/interface/name"
	name__1 = "/interfaces/interface/unit/name"
	name__2 = "/interfaces/interface/unit/family/inet/address/name"
}

