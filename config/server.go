package config

type Server struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}

func (s Server) GetPort() string {
	if s.Port != "" {
		return s.Port
	}

	return "8080"
}

func (s Server) GetHost() string {
	if s.Host != "" {
		return s.Host
	}

	return "127.0.0.1"
}
