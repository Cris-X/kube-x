package api

import "kube-x/api/example"

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	ExampleApiGroup example.ApiGroup
}
