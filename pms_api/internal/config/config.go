package config

import "github.com/ilyakaznacheev/cleanenv"

const PATH_TO_CONFIG = "./pms_api/config/config.yaml"

type Config struct {
	Http     Http
	Database Database
	LogLevel string `yaml:"log_level"`
}

type Http struct {
	Port     string `yaml:"port"`
	BasePath string `yaml:"base_path"`
}

type Database struct {
	ConnectionString string `yaml:"connection_string"`
}

func NewConfig() (*Config, error) {
	var config Config
	if err := cleanenv.ReadConfig(PATH_TO_CONFIG, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
