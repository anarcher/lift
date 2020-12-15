package k8s

output: [ for k, v in resources if len(v) > 0 {v}]
