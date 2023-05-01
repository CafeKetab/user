package http

import (
	"encoding/json"
	"fmt"

	"github.com/CafeKetab/user/internal/api/grpc"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Server struct {
	config *Config
	logger *zap.Logger
	auth   grpc.AuthClient
	app    *fiber.App
}

func New(cfg *Config, log *zap.Logger, auth grpc.AuthClient) *Server {
	server := &Server{config: cfg, logger: log, auth: auth}

	server.app = fiber.New(fiber.Config{JSONEncoder: json.Marshal, JSONDecoder: json.Unmarshal})

	v1 := server.app.Group("/v1")
	_ = v1
	// v1.Group("/users", server.redirect)
	// v1.Group("/books", server.redirect)
	// v1.Group("/books", server.authenticate, server.redirect)

	return server
}

func (server *Server) Serve() error {
	addr := fmt.Sprintf(":%d", server.config.ListenPort)
	if err := server.app.Listen(addr); err != nil {
		server.logger.Error("error resolving server", zap.Error(err))
		return err
	}
	return nil
}
