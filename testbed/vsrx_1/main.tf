terraform {
    required_providers {
        junos-vsrx = {
            source = "juniper/providers/junos-vsrx"
            version = "19.41.101"
        }
    }
}
 
resource "junos-vsrx_InterfacesInterfaceUnitFamilyInetAddressName" "vsrx_1" {
    resource_name = "vsrx_1"
    name = "et-0/0/0"
    name__1 = "0"
    name__2 = "10.0.0.1/30"
}
 
resource "junos-vsrx_InterfacesInterfaceUnitFamilyInetAddressName" "vsrx_2" {
    resource_name = "vsrx_2"
    name = "lo0"
    name__1 = "0"
    name__2 = "1.1.1.1/32"
}
 

 
resource "junos-vsrx_InterfacesInterfaceDescription" "vsrx_4" {
    resource_name = "vsrx_4"
    name = "lo0"
    description = "Loopback 1"
}
 
resource "junos-vsrx_Routing__OptionsRouter__Id" "vsrx_5" {
    resource_name = "vsrx_5"
    router__id = "1.1.1.1"
}
 

 
