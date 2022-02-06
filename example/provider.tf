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

resource "haproxy_maps" "test" {
  map   = "ratelimit"
  key   = "/metrics"
  value = "50"
}
