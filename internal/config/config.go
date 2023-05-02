package config

import (
	"github.com/CafeKetab/user/internal/api/grpc"
	"github.com/CafeKetab/user/internal/api/http"
	"github.com/CafeKetab/user/pkg/logger"
	"github.com/CafeKetab/user/pkg/rdbms"
)

type Config struct {
	Logger *logger.Config `koanf:"logger"`
	RDBMS  *rdbms.Config  `koanf:"rdbms"`
	HTTP   *http.Config   `koanf:"http"`
	GRPC   *grpc.Config   `koanf:"grpc"`
}
