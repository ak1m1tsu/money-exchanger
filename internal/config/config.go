package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	Logger      `yaml:"logger"`
	Server      `yaml:"server"`
	Database    `yaml:"database"`
	ExternalAPI `yaml:"external_api"`
}

type Logger struct {
	Level string `yaml:"level" env:"LOGGER_LEVEL"`
}

type Server struct {
	Port string `yaml:"port" env:"PORT"`
}

type Database struct {
	DSN string `yaml:"dsn" env:"DATABASE_DSN"`
}

type ExternalAPI struct {
	BaseURL string `yaml:"base_url" env:"EXTERNAL_API_BASE_URL"`
	Key     string `yaml:"key" env:"EXTERNAL_API_KEY"`
}

func New(file string) (*Config, error) {
	cfg := new(Config)

	err := cleanenv.ReadConfig(file, cfg)
	if err != nil {
		return nil, err
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
