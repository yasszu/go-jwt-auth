package config

import (
	"github.com/labstack/echo"
	"gopkg.in/yaml.v3"
	"os"
)

// Config is conf.yml
type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	Database struct {
		Username string `yaml:"user"`
		Password string `yaml:"password"`
		DB       string `yaml:"db"`
	} `yaml:"database"`
	JWT struct {
		Secret string `yaml:"secret"`
	} `yaml:"jwt"`
}

func NewConfig() Config {
	return Config{}
}

// LoadConfig load conf.yml
func (cfg Config) Load() (Config, error) {
	f, err := os.Open("config/conf.yml")
	if err != nil {
		return cfg, err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}

func GetConfig(c echo.Context) Config {
	conf := c.Get("conf").(Config)
	return conf
}
