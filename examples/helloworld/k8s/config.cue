package k8s

#config: {
	name:      string
	namespace: *name | string
	replicas:  *1 | int
	image:     string

	service: name: *config.name | string
}

config: #config
