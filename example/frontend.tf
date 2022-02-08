locals {
  test = [
    "frontend1",
    "frontend2",
    "frontend3",
    "frontend4",
    "frontend5",
    "frontend6",
    "frontend7",
    "frontend8",
    "frontend9",
    "frontend10",
    "frontend11",
    "frontend12",
    "frontend13",
  ]
}
resource "haproxy_frontend" "test" {
  for_each = { for frontend in local.test : frontend => frontend }
  name     = each.key
}
