package entities

type Config struct {
	Env string `yaml:"env"`

	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
}
