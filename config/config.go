package config

import (
	"log"
	"os"

	"github.com/labstack/echo"
	"gopkg.in/yaml.v3"
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

// LoadConfig load conf.yml
func LoadConfig() Config {
	f, err := os.Open("config/conf.yml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return cfg
}

func GetConfig(c echo.Context) Config {
	conf := c.Get("conf").(Config)
	return conf
}
