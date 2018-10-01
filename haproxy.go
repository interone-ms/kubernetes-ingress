package main

// fixed paths to haproxy items
const (
	HAProxyCFG    = "/etc/haproxy/haproxy.cfg"
	HAProxySocket = "/var/run/haproxy-runtime-api.sock"
)

func getGlobal() string {
	return `
global
    daemon
    stats socket /var/run/haproxy-runtime-api.sock level admin expose-fd listeners
    maxconn 2000
    tune.ssl.default-dh-param 1024
    ssl-default-bind-ciphers ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-AES256-GCM-SHA384:DHE-RSA-AES128-GCM-SHA256:DHE-DSS-AES128-GCM-SHA256:kEDH+AESGCM:ECDHE-RSA-AES128-SHA256:ECDHE-ECDSA-AES128-SHA256:ECDHE-RSA-AES128-SHA:ECDHE-ECDSA-AES128-SHA:ECDHE-RSA-AES256-SHA384:ECDHE-ECDSA-AES256-SHA384:ECDHE-RSA-AES256-SHA:ECDHE-ECDSA-AES256-SHA:DHE-RSA-AES128-SHA256:DHE-RSA-AES128-SHA:DHE-DSS-AES128-SHA256:DHE-RSA-AES256-SHA256:DHE-DSS-AES256-SHA:DHE-RSA-AES256-SHA:!aNULL:!eNULL:!EXPORT:!DES:!RC4:!3DES:!MD5:!PSK
	ssl-default-bind-options no-sslv3 no-tls-tickets
`
}

func getDefault() string {
	return `
defaults
    log global
    maxconn 2000
    option redispatch
    option dontlognull
    option http-server-close
    option http-keep-alive
    timeout http-request    5s
    timeout connect         5s
    timeout client          50s
    timeout client-fin      50s
    timeout queue           5s
    timeout server          50s
    timeout server-fin      50s
    timeout tunnel          1h
    timeout http-keep-alive 1m
`
}