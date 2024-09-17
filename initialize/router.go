package initialize

import (
	"github.com/gin-gonic/gin"
	"kube-x/middleware"
	"kube-x/router"
)

// Routers initialize routers
func Routers() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors)

	// get router group
	exampleGroup := router.RouterGroupApp.ExampleRouterGroup
	k8sGroup := router.RouterGroupApp.K8SRouterGroup

	// initialize routers
	exampleGroup.InitExample(r)
	k8sGroup.InitK8SRouter(r)
	return r
}
