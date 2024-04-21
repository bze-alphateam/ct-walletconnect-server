package config

type App struct {
	Server   *Server   `yaml:"server"`
	Security *Security `yaml:"security"`
	Redis    *Redis    `yaml:"redis"`
}
