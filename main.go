package main

import (
	"ginDemo/config"
	"ginDemo/router"
)

func main() {
	config.InitConfig()
	r := router.SetupRouter()
	_ = r.Run(config.AppConfig.App.Port)
}
