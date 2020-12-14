package k8s

service: {
	apiVersion: "v1"
	kind:       "Service"
	metadata: {
		name:      *config.service.name | string
		namespace: config.namespace
	}
	spec: {
		ports: [
			{
				name:       config.name
				port:       80
				targetPort: "http"
			},
		]
		selector: {
			app: config.name
		}
	}
}
