package example

import (
	"github.com/gin-gonic/gin"
	"kube-x/api"
)

type ExampleRouter struct {
}

func (*ExampleRouter) InitExample(r *gin.Engine) {
	routerGroup := r.Group("/example")
	apiGroup := api.ApiGroupApp.ExampleApiGroup
	routerGroup.GET("/ping", apiGroup.ExampleTest)
}
