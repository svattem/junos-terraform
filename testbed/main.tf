
# replace {text} with your own test setup

terraform {
	required_providers {
		junos-vsrx = {
            source = "juniper/providers/junos-vsrx"
            version = "19.41.101"
		}
	}
}

provider "junos-vsrx" {
	host = "66.129.234.201"
	port = 38015
	username = "jcluser"
	password = "Juniper!1"
	sshkey = ""
}

module "vsrx_1" {
	source = "./vsrx_1"

	providers = {junos-vsrx = junos-vsrx}

	depends_on = [junos-vsrx_destroycommit.commit-main]
}


resource "junos-vsrx_commit" "commit-main" {
	resource_name = "commit"
	depends_on = [module.vsrx_1]
}

resource "junos-vsrx_destroycommit" "commit-main" {
	resource_name = "destroycommit"
}
	