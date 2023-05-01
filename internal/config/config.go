package config

import (
	"github.com/CafeKetab/user/internal/repository"
	"github.com/CafeKetab/user/pkg/logger"
)

type Config struct {
	Logger     *logger.Config     `koanf:"logger"`
	Repository *repository.Config `koanf:"repository"`
}
