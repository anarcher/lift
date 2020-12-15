package k8s

resources: deployment: {
	apiVersion: "apps/v1"
	kind:       "Deployment"
	metadata: {
		name:      config.name
		namespace: config.namespace
	}
	spec: {
		minReadySeconds:      config.minReadySeconds
		replicas:             config.replicas
		revisionHistoryLimit: config.revisionHistoryLimit
		selector: matchLabels: config.selector
		template: {
			metadata: labels: config.selector
			if len(config.podAnnotations) > 0 {
				metadata: annotations: config.podAnnotations
			}
			spec: {
				containers: [{
					if len(config.envSpec) > 0 {
						env: [ for k, v in config.envSpec {v, name: k}]
					}
					if len(config.envFrom) > 0 {
						envFrom: [ for k, v in config.envFrom for vv in v {"\(k)Ref": name: vv}]
					}

					image:           config.image
					imagePullPolicy: "IfNotPresent"
					livenessProbe:   config.livenessProbe
					name:            config.name

					ports: [ for k, p in config.expose.ports & config.ports {
						name:          k
						containerPort: p
					}]

					readinessProbe: config.readinessProbe
					resources:      config.resources
				}, ...]
				securityContext: fsGroup: 65534
				serviceAccountName:            config.serviceAccountName
				terminationGracePeriodSeconds: config.terminationGracePeriodSeconds
			}
		}
	}
}
