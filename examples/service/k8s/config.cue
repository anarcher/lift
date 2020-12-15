package k8s

config: {
	name:      string
	namespace: *name | string

	image: string

	labels: [string]: string

	replicas:             *1 | int
	minReadySeconds:      *5 | int
	revisionHistoryLimit: *3 | int

	selector: [string]: string
	if len(selector) <= 0 {
		selector: app: config.name
	}

	podAnnotations: [string]: string

	livenessProbe?: {
		httpGet?: {
			path: *"/health-check" | string
			port: *"http" | string | int
		}
		initialDelaySeconds: *5 | int
		timeoutSeconds:      *1 | int
	}
	readinessProbe?: {
		httpGet?: {
			path: *"/health-check" | string
			port: *"http" | string | int
		}
		initialDelaySeconds: *5 | int
		timeoutSeconds:      *1 | int
	}

	expose: ports: [string]: int
	ports: [string]: int

	if len(expose.ports) <= 0 {
		expose: ports: http: 8080
	}

	arg: [string]: string
	args: [ for k, v in arg {"-\(k)=\(v)"}] | [...string]

	env: [string]: string
	envSpec: [string]: {}
	envSpec: {
		for k, v in env {
			"\(k)": value: v
		}
	}

	envFrom: configMap: [...string]
	envFrom: secret: [...string]

	resources: {
		limits: {
			cpu:    *"250m" | string
			memory: *"64Mi" | string
		}
		requests: {
			cpu:    *"250m" | string
			memory: *"64Mi" | string
		}
	}

	serviceAccount: {
		name:   *config.name | string
		enable: *true | bool
	}

	serviceAccountName:            *serviceAccount.name | string
	terminationGracePeriodSeconds: *60 | int

	irsa: *"" | string

	configMap: [string]: {}
	secret: [string]: {}

	ingress: {
		host:        string
		serviceName: *config.name | string
		servicePort: *"http" | string
		annotations: [string]: string
	}
}
