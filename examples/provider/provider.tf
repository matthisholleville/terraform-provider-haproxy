provider "haproxy" {
  server_addr = "localhost:5555" # optionally use HAPROXY_SERVER env var
  username    = "admin"          # optionally use HAPROXY_USERNAME env var
  password    = "adminpwd"       # optionally use HAPROXY_PASSWORD env var

  # you may need to allow insecure TLS communications unless you have configured
  # certificates for your server
  insecure = true # optionally use HAPROXY_INSECURE env var
}
