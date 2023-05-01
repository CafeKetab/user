package config

import (
	"github.com/CafeKetab/user/internal/repository"
	"github.com/CafeKetab/user/pkg/logger"
)

func Default() *Config {
	return &Config{
		Logger: &logger.Config{
			Development: true,
			Level:       "debug",
			Encoding:    "console",
		},
		Repository: &repository.Config{
			Host:     "localhost",
			Port:     5432,
			Username: "TEST_USER",
			Password: "TEST_PASSWORD",
			Database: "TEST_DB",
		},
	}
}
