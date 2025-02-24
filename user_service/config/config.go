package config

import (
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
)

type Postgresql struct {
	Name     string `yaml:"name"`
	Password string `yaml:"password"`
	User     string `yaml:"user"`
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
}

type App struct {
	Port string `yaml:"port"`
}

type ParseConfig struct {
	Postgres Postgresql `yaml:"postgresql"`
	App      App        `yaml:"app"`
}

type Config struct {
	DSN  string
	ADDR string
}

func MustLoad() (*Config, error) {
	configPath := "config.yaml"
	configFile, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("MustLoad: failed read config file: %w", err)
	}

	pointerConfigParse := &ParseConfig{}
	if err := yaml.Unmarshal(configFile, pointerConfigParse); err != nil {
		return nil, fmt.Errorf("MustLoad: failed parse config on path %s: %w", configPath, err)
	}

	postgres := pointerConfigParse.Postgres
	addr := fmt.Sprintf("localhost:%s", pointerConfigParse.App.Port)
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", postgres.User, postgres.Password, postgres.Host, postgres.Port, postgres.Name)

	return &Config{
		ADDR: addr,
		DSN:  dsn,
	}, nil
}
