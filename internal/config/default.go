package config

import (
	"github.com/CafeKetab/user/internal/api/grpc"
	"github.com/CafeKetab/user/internal/api/http"
	"github.com/CafeKetab/user/pkg/logger"
	"github.com/CafeKetab/user/pkg/rdbms"
)

func Default() *Config {
	return &Config{
		Logger: &logger.Config{
			Development: true,
			Level:       "debug",
			Encoding:    "console",
		},
		RDBMS: &rdbms.Config{
			Host:     "localhost",
			Port:     5432,
			Username: "TEST_USER",
			Password: "TEST_PASSWORD",
			Database: "USER_DB",
		},
		HTTP: &http.Config{
			ListenPort: 8081,
		},
		GRPC: &grpc.Config{
			AuthGrpcClientAddress: "localhost:9090",
		},
	}
}
