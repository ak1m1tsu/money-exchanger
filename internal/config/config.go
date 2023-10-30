package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	Logger   `yaml:"logger"`
	Server   `yaml:"server"`
	Database `yaml:"database"`
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
