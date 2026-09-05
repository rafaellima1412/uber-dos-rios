package config

import (
	"os"

	"github.com/rafaellima1412/uber-dos-rios/internal/common"
	"github.com/rafaellima1412/uber-dos-rios/internal/logger"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	Database struct {
		URL string `yaml:"url"`
	} `yaml:"database"`
}

func LoadConfig(path string) *Config {
	file, err := os.ReadFile(path)
	if err != nil {
		logger.Fatal(common.FailedToReadConfigFile, zap.Error(err))
	}

	var cfg Config
	if err := yaml.Unmarshal(file, &cfg); err != nil {
		logger.Fatal(common.FailedToUnmarshalConfigFile, zap.Error(err))
	}

	if envDB := os.Getenv("DATABASE_URL"); envDB != "" {
		cfg.Database.URL = envDB
	}

	return &cfg
}
