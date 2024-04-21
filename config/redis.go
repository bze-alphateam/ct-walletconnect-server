package config

type Redis struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	Database int    `yaml:"database"`
}
