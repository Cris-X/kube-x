package initialize

import (
	"kube-x/global"
	"kube-x/logger"
)

func InitLogger() {
	// init logger
	global.Logger = logger.InitLogger()
	return
}
