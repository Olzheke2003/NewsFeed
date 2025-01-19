package config

import (
	"io/ioutil"
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type ServerConfig struct {
	BindAddr       string        `yaml:"bindAddr"`
	LogLevel       string        `yaml:"logLevel"`
	DatabaseURL    string        `yaml:"databaseURL"`
	JwtSecretKey   string        `yaml:"JwtSecretKey"`
	TokenExpiry    time.Duration `yaml:"-"`           // Сохраняем как Duration
	TokenExpiryRaw string        `yaml:"TokenExpiry"` // Храним исходное значение как строку
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
	cfg.JwtSecretKey = getEnv("JWT_SECRET_KEY", cfg.JwtSecretKey)

	// Конвертируем строку в time.Duration
	if expiry, err := time.ParseDuration(cfg.TokenExpiryRaw); err == nil {
		cfg.TokenExpiry = expiry
	} else {
		log.Fatalf("Invalid TokenExpiry format: %v", err)
	}

	return cfg
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
