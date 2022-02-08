build:
	go build -o ~/go/bin/terraform-provider-haproxy

test:
	@docker rm tf_haproxy_acc_test -f || true
	@docker run --name tf_haproxy_acc_test --rm -d -v $(shell pwd)/tools:/usr/local/etc/haproxy:rw -p 8404:8404 -p 5555:5555 haproxytech/haproxy-alpine:2.4

	@docker exec tf_haproxy_acc_test mkdir -p /usr/local/etc/haproxy/maps/test.map

	sleep 10

