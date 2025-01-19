package config

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var SecretKey = []byte("your-secret-key")

type ServerConfig struct {
	BindAddr    string `yaml:"bindAddr"`
	LogLevel    string `yaml:"logLevel"`
	DatabaseURL string `yaml:"databaseURL"`
}

func New(configFilePath string) *ServerConfig {
	cfg := &ServerConfig{}
	fileData, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	err = yaml.Unmarshal(fileData, cfg)
	if err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}

	cfg.BindAddr = getEnv("BIND_ADDR", cfg.BindAddr)
	cfg.LogLevel = getEnv("LOG_LEVEL", cfg.LogLevel)
	cfg.DatabaseURL = getEnv("DATABASE_URL", cfg.DatabaseURL)

	return cfg
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
