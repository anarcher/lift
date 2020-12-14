package main

import k8s "example.com/k8s"

http_server: k8s
http_server: config: {
	namespace: config.namespace
	name:      "http_server"
	image:     "http_server:v0.0.1"
}

grpc_server: k8s & {
	config: {
		namespace: config.namespace
		name:      "grpc_server"
		image:     "grpc_server:v0.0.1"
	}
}

config: {
	namespace: "server"
}
