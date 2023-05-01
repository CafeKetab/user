package config

import (
	"github.com/CafeKetab/user/pkg/logger"
	"github.com/CafeKetab/user/pkg/rdbms"
)

type Config struct {
	Logger *logger.Config `koanf:"logger"`
	RDBMS  *rdbms.Config  `koanf:"rdbms"`
}
