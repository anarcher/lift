package k8s

resources: [
	deployment,
	service,
]

output: [ for v in resources if len(v) > 0 {v}]
