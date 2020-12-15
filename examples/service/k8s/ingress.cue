package k8s

resources: ingress: {
	apiVersion: "networking.k8s.io/v1beta1"
	kind:       "Ingress"
	metadata: {
		annotations: config.ingress.annotations
		name:        config.name
		namespace:   config.namespace
	}
	spec: rules: [{
		host: config.ingress.host
		http: paths: [{
			backend: {
				serviceName: config.ingress.serviceName
				servicePort: config.ingress.servicePort
			}
			path: "/"
		}, ...]
	}, ...]
}
