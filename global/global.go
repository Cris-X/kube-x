package global

import (
	"k8s.io/client-go/kubernetes"
	"kube-x/config"
)

var (
	CONF          config.Server         // global config: manage project's configuration
	KubeConfigSet *kubernetes.Clientset // global k8s clientset: manage k8s clientset
)
