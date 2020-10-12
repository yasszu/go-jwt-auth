package conf

import (
	"gopkg.in/yaml.v3"
	"os"
)

// Conf is conf.yml
type Conf struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	Postgres struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"user"`
		Password string `yaml:"password"`
		DB       string `yaml:"db"`
	} `yaml:"postgres"`
	JWT struct {
		Secret string `yaml:"secret"`
	} `yaml:"jwt"`
}

func NewConf() (*Conf, error) {
	var cnf Conf
	f, err := os.Open("conf/conf.yml")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&cnf); err != nil {
		return nil, err
	}
	return &cnf, nil
}
