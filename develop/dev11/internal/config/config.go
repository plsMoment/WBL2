package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type Config struct {
	HTTPServer `yaml:"http_server"`
	DB         `yaml:"db"`
}

type HTTPServer struct {
	Host string `yaml:"host" env-default:"localhost"`
	Port string `yaml:"port" env-default:"8080"`
}

type DB struct {
	DSN string `yaml:"dsn"`
}

func MustLoad() *Config {
	configPath := "./config.yml"
	// checking the existence of the file
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file %s doesn't exist", configPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Cannot read config: %s", err)
	}

	return &cfg
}
