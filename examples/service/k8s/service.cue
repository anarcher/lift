package k8s

resources: service: {
	apiVersion: "v1"
	kind:       "Service"
	metadata: {
		labels:    config.labels
		name:      config.name
		namespace: config.namespace
	}
	spec: {
		ports: [ for n, p in config.expose.ports {
			name:       n
			port:       p
			targetPort: p
		}]
		selector: config.selector
	}
}
