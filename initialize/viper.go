package initialize

import (
	"github.com/spf13/viper"
	"kube-x/global"
)

func Viper() {
	v := viper.New()

	// set config file path
	v.SetConfigFile("config.yaml")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}

	// map config to global variable
	err = v.Unmarshal(&global.CONF)
	if err != nil {
		panic(err.Error())
	}
}
