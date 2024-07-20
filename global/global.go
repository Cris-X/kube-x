package global

import (
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
	"kube-x/config"
)

var (
	CONF          config.Server         // global config: manage project's configuration
	Logger        *zap.Logger           // global logger: manage project's log
	KubeConfigSet *kubernetes.Clientset // global k8s clientset: manage k8s clientset
)
