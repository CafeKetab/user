package config

import (
	"github.com/CafeKetab/user-go/pkg/logger"
)

type Config struct {
	Logger *logger.Config `koanf:"logger"`
}
