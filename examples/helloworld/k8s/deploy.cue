package k8s

deployment: {
	apiVersion: "apps/v1"
	kind:       "Deployment"
	metadata: {
		namespace: config.namespace
		name:      config.name
	}
	spec: {
		replicas: config.replicas
		selector: matchLabels: app: config.name
		template: {
			metadata: labels: app: config.name
			spec: containers: [{
				name:  config.name
				image: config.image
				ports: [{
					name:          "http"
					containerPort: 8080
				}]
			}]
		}
	}
}
