package main

import (
	"noodlenote/config"
	"noodlenote/model"
	"noodlenote/router"
)

func main() {
	config.SetUp()
	model.SetUp()

	router.NewRouter().Run(config.Cfg.ServerConfig.Host + ":" + config.Cfg.ServerConfig.Port)
}
