package config

type Config struct {
	ServerConfig Server
	AppConfig    App
}

var Cfg Config

func SetUp() {
	var config Config
	config.ServerConfig.DefaultServerConfig()
	config.AppConfig.DefaultAppConfig()

	Cfg = config
}
