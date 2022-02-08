build:
	go build -o ~/go/bin/terraform-provider-haproxy

test:
	@docker rm tf_haproxy_acc_test -f || true
	@cd ./tools && docker build . -t tf_haproxy_acc_test:2.4
	@docker run --name tf_haproxy_acc_test --rm -d -p 8404:8404 -p 5555:5555 tf_haproxy_acc_test:2.4
	sleep 5

	TF_ACC=1 \
	HAPROXY_SERVER="localhost:5555" \
	HAPROXY_USERNAME="admin" \
	HAPROXY_PASSWORD="adminpwd" \
	HAPROXY_INSECURE="true" \
	go test -v -cover -count 1 ./internal/provider

	@docker rm tf_haproxy_acc_test -f

