terraform {
  required_providers {
    haproxy = {
      version = "~> 1.0.0"
      source  = "matthisholleville/haproxy"
    }
  }
}
provider "haproxy" {
  server_addr = "CHANGE_ME"
  username    = "CHANGE_ME"
  password    = "CHANGE_ME"
}

locals {
  ips = ["92.91.233.55", "92.91.233.53"]
}

# resource "haproxy_maps" "test" {
#   for_each = { for ip in local.ips : ip => ip }
#   map      = "blacklist"
#   key      = each.key
# }

# resource "haproxy_maps" "test2" {
#   map = "blacklist"
#   key = "92.91.233.55"
# }
