package config

type Server struct {
	Host    string `yaml:"Host"`
	Port    string `yaml:"Port"`
	RunMode string `yaml:"RunMode"`
}

func (server *Server) DefaultServerConfig() {
	server.Host = "localhost"
	server.Port = "9000"
	server.RunMode = "release"
}
