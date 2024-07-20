package main

import (
	"kube-x/global"
	"kube-x/initialize"
)

// program entry point
func main() {
	// 1.initialize routers
	r := initialize.Routers()

	// 2.initialize viper
	initialize.Viper()

	// 3.initialize k8s
	initialize.K8S()

	panic(r.Run(global.CONF.System.Addr))
}
