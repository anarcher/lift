package k8s

resources: serviceAccount: {
	if config.serviceAccount.enable == true {
		apiVersion: "v1"
		kind:       "ServiceAccount"
		metadata: {
			name:      config.serviceAccount.name
			namespace: config.namespace
			if config.irsa != "" {
				annotations: "eks.amazonaws.com/role-arn": config.irsa
			}

		}
	}
}
