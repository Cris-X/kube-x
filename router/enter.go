package router

import (
	"kube-x/router/example"
	"kube-x/router/k8s"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	ExampleRouterGroup example.ExampleRouter
	K8SRouterGroup     k8s.K8SRouter
}
