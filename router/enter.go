package router

import "kube-x/router/example"

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	ExampleRouterGroup example.ExampleRouter
}
