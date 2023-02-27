package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type App struct {
	Port int `yaml:"port"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	SSLMode  string `yaml:"sslmode"`
}

type Configuration struct {
	App App      `yaml:"App"`
	DB  Database `yaml:"Database"`
}

const defaultPath string = "config.yml"

func Load(filename string) (*Configuration, error) {
	config := &Configuration{}

	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(content, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func LoadConf() (*Configuration, error) {
	return Load(defaultPath)
}
