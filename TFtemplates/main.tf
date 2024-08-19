
# replace {text} with your own test setup

terraform {
	required_providers {
		junos-vsrx = {
			source = "{input source path}"
			version = "{input version here}"
		}
	}
}

provider "junos-vsrx" {
	host = "localhost"
	port = 8300
	username = "root"
	password = "juniper123"
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
	