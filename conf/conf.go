package conf

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
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
	cnf.loadEnv()
	return &cnf, nil
}

func (c *Conf) loadEnv() {
	fmt.Println("load env:")
	if v, ok := os.LookupEnv("POSTGRES_HOST"); ok {
		fmt.Println("- POSTGRES_HOST:", v)
		c.Postgres.Host = v
	}
	if v, ok := os.LookupEnv("SERVER_HOST"); ok {
		fmt.Println("- SERVER_HOST:", v)
		c.Server.Host = v
	}
	if v, ok := os.LookupEnv("SERVER_PORT"); ok {
		fmt.Println("- SERVER_PORT:", v)
		c.Server.Port = v
	}
}
