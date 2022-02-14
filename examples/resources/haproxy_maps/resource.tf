resource "haproxy_maps" "my-key" {
  map   = "ratelimit"
  key   = "/metrics"
  value = "50"
}
