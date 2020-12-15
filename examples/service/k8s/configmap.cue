package k8s

for k, v in config.configMap {
	resources: "configMaps\(k)": {
		apiVersion: "v1"
		kind:       "ConfigMap"
		metadata: name: k
		data: v
	}
}
