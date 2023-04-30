package config

import (
	"github.com/CafeKetab/auth-go/pkg/crypto"
	"github.com/CafeKetab/auth-go/pkg/logger"
	"github.com/CafeKetab/auth-go/pkg/token"
)

type Config struct {
	Logger *logger.Config `koanf:"log"`
	Token  *token.Config  `koanf:"token"`
	Crypto *crypto.Config `koanf:"crypto"`
}
