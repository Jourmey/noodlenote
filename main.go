package main

import (
	"noodlenote/config"
	"noodlenote/router"
)

func main() {
	conf := config.Default()
	r := router.NewRouter()

	r.Run(conf.ServerConfig.Host + ":" + conf.ServerConfig.Port)
}
