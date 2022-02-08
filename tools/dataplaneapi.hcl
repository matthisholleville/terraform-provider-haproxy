config_version = 2

name = "famous_condor"

mode = "single"

dataplaneapi {
  host = "0.0.0.0"
  port = 8080

  user "admin" {
    insecure = true
    password = "adminpwd"
  }

  resources {
    maps_dir      = "/etc/haproxy/maps"
    ssl_certs_dir = "/etc/haproxy/ssl"
    spoe_dir      = "/etc/haproxy/spoe"
  }

  advertised {}
}

haproxy {
  config_file = "/etc/haproxy/haproxy.cfg"
  haproxy_bin = "/usr/local/sbin/haproxy"

  reload {
    reload_cmd  = "kill -SIGUSR2 1"
    restart_cmd = "kill -SIGUSR2 1"
  }
}
