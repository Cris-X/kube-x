package initialize

import (
	"github.com/gin-gonic/gin"
	"kube-x/router"
)

// Routers initialize routers
func Routers() *gin.Engine {
	r := gin.Default()
	exampleGroup := router.RouterGroupApp.ExampleRouterGroup
	exampleGroup.InitExample(r)
	return r
}
