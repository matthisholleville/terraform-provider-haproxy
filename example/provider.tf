terraform {
  required_providers {
    haproxy = {
      version = "~> 1.0.0"
      source  = "matthisholleville/haproxy"
    }
  }
}
provider "haproxy" {
  server_addr = "10.100.0.130:5555"
  username    = "CHANGE_ME"
  password    = "CHANGE_ME"
  insecure    = true
}

resource "haproxy_maps" "test" {
  map   = "ratelimit"
  key   = "/metrics"
  value = "50"
}
