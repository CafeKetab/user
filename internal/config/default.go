package config

import (
	"github.com/CafeKetab/user-go/pkg/logger"
)

func Default() *Config {
	return &Config{
		Logger: &logger.Config{
			Development: true,
			Level:       "debug",
		},
	}
}
