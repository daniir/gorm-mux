package settings

import (
	_ "embed"

	"gopkg.in/yaml.v3"
)

//go:embed config.yaml
var configYamlFile []byte

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type ConfigSrv struct {
	Port int            `yaml:"port"`
	DB   DatabaseConfig `yaml:"database"`
}

func New() (*ConfigSrv, error) {
	var c ConfigSrv
	err := yaml.Unmarshal(configYamlFile, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
