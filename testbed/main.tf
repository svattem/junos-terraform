terraform {
  required_providers {
    junos-vsrx = {
      source = "juniper/providers/junos-vsrx"
      version = "19.41.101"
    }
  }
}

provider "junos-vsrx" {
    host = "100.123.12.0"
    port = 22
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
