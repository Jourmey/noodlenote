package config

type Config struct {
	ServerConfig Server
}

func Default() Config {
	var config Config
	config.ServerConfig.DefaultServerConfig()
	return config
}
