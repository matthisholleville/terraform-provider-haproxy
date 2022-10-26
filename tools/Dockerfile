FROM haproxytech/haproxy-alpine:2.6.5

COPY ./haproxy.cfg /usr/local/etc/haproxy/haproxy.cfg

COPY ./dataplaneapi.hcl /usr/local/etc/haproxy/dataplaneapi.hcl

RUN mkdir -p /usr/local/etc/haproxy/maps/ && touch /usr/local/etc/haproxy/maps/test.map