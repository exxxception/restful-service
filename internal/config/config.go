package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Server struct {
	Host string `yaml:"host" env-default:"localhost"`
	Port string `yaml:"port" env-default:"8080"`
}

type DB struct {
	DSN string `yaml:"DSN" env-requered:"true"`
}

type Config struct {
	Env    string `yaml:"env" env-default:"local"`
	Server `yaml:"server"`
	DB     `yaml:"db"`
	ApiKey string `yaml:"apiKey" env-requered:"true"`
}

func LoadConfig() *Config {
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatal("config file does not exist")
	}

	var cfg Config
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		log.Fatalf("connot read config: %s", err)
	}

	return &cfg
}
