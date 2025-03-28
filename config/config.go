package config

import (
	"fmt"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"gopkg.in/yaml.v3"
)

type Endpoint struct {
	Name        string
	Url         string
	AuthHeaders map[string]string `yaml:"auth_headers"`
}

type Config struct {
	Endpoints []Endpoint               `yaml:"endpoints"`
	Polling   map[string]time.Duration `yaml:"polling"`
	Filters   map[string]string        `yaml:"filters"`
}

func Load() (*Config, error) {
	userConfig := &Config{}
	cfgData, err := os.ReadFile("config.yaml")

	if err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	err = yaml.Unmarshal(cfgData, &userConfig)

	if err != nil {
		throwErr := fmt.Errorf("Error parsing config file: %v", err)
		panic(throwErr)
	}

	return userConfig, nil
}
