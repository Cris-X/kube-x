package api

import (
	"kube-x/api/example"
	"kube-x/api/k8s"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	ExampleApiGroup example.ApiGroup // ExampleApiGroup is a struct that contains all the example api methods
	K8SApiGroup     k8s.ApiGroup     // K8SApiGroup is a struct that contains all the k8s api methods
}
