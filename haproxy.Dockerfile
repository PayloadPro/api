FROM haproxy:1.7
COPY deployments/haproxy-dev.cfg /usr/local/etc/haproxy/haproxy.cfg
