terraform {
  required_providers {
    haproxy = {
      source  = "matthisholleville/haproxy"
      version = "0.1.0"
    }
  }
}
provider "haproxy" {
  server_addr = "localhost:5555"
  username    = "admin"
  password    = "adminpwd"
  insecure    = true
}

# resource "haproxy_maps" "test" {
#   map   = "ratelimit"
#   key   = "/metrics"
#   value = "50"
# }

resource "haproxy_frontend" "test" {
  name = "%[1]s"
}
