resource "haproxy_maps" "test" {
  map   = "ratelimit"
  key   = "/metrics"
  value = "50"
}
