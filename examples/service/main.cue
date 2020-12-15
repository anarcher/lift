package main

import "example.com/k8s"

svc: k8s
svc: config: {
	name:  "helloworld"
	image: "helloworld:v0.0.1"

	livenessProbe:  probe
	readinessProbe: probe

	envSpec: "MONITOR_TRACER_DATADOG_HOST": {
		valueFrom: fieldRef: fieldPath: "status.hostIP"
	}
	env: "TEST": "true"

	irsa: "IRSA"

	configMap: "helloworld": {
		ENV: "alpha"
		RGN: "kr"
	}
	envFrom: configMsp: ["helloworld"]
	ingress: host: "example.com"
}

probe: httpGet: path: "/health-check"

svc: resources: deployment: metadata: annotations: {
	env: "test"
}
