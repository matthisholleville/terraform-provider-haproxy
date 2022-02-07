<a href="https://www.haproxy.com">
    <img src="https://upload.wikimedia.org/wikipedia/commons/a/ab/Haproxy-logo.png" alt="Pritunl logo" title="Pritunl" align="right" height="100" />
</a>
<a href="https://terraform.io">
    <img src="https://dashboard.snapcraft.io/site_media/appmedia/2019/11/terraform.png" alt="Terraform logo" title="Terraform" align="right" height="100" />
</a>

# Terraform Provider to configure HAProxy

- Website: https://www.terraform.io
- HAProxy: https://haproxy.com
- HAProxy Dataplaneapi: https://github.com/haproxytech/dataplaneapi

## Requirements
-	[Terraform](https://www.terraform.io/downloads.html) >=1.1.x
-	[Go](https://golang.org/doc/install) 1.17.x (to build the provider plugin)

## Building The Provider

```sh
$ git clone git@github.com:matthisholleville/terraform-provider-haproxy.git
$ go build -o terraform-provider-haproxy
```

## Example usage

```hcl
# Set the required provider and versions
terraform {
  required_providers {
    haproxy = {
      version = "~> 1.0.0"
      source  = "matthisholleville/haproxy"
    }
  }
}

# Configure the haproxy provider
provider "haproxy" {
  server_addr = "10.100.0.130:5555"
  username    = "CHANGE_ME"
  password    = "CHANGE_ME"
  insecure    = true
}

# Create a new entrie in the ratelimit HAProxy Maps file
resource "haproxy_maps" "test" {
  map   = "ratelimit"
  key   = "/metrics"
  value = "50"
}
```


### Ressources implemented

- [x] maps

### Ressources in the roadmap

- [ ] backend
- [ ] frontend
- [ ] server
- [ ] acl
- [ ] httpRequestRule
- [ ] httpResponseRule

## License

The Terraform HAProxy Provider is available to everyone under the terms of the Mozilla Public License Version 2.0. [Take a look the LICENSE file](LICENSE).