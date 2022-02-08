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
