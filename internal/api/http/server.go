package http

import (
	"encoding/json"
	"fmt"

	"github.com/CafeKetab/user/internal/api/grpc"
	"github.com/CafeKetab/user/internal/repository"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Server struct {
	config     *Config
	logger     *zap.Logger
	repository repository.Repository
	auth       grpc.AuthClient
	app        *fiber.App
}

func New(cfg *Config, log *zap.Logger, repo repository.Repository, auth grpc.AuthClient) *Server {
	server := &Server{config: cfg, logger: log, repository: repo, auth: auth}

	server.app = fiber.New(fiber.Config{JSONEncoder: json.Marshal, JSONDecoder: json.Unmarshal})

	v1 := server.app.Group("/v1")
	v1.Post("/register", server.register)
	v1.Post("/login", server.login)
	// v1.Get("/:id<uint64>", server.fetchUserId, server.user)
	// v1.Get("/me", server.fetchUserId, server.me)
	// v1.Post("/update-information", server.fetchUserId, server.updateInformation)
	// v1.Post("/update-password", server.fetchUserId, server.updatePassword)

	return server
}

func (server *Server) Serve() {
	addr := fmt.Sprintf(":%d", server.config.ListenPort)
	if err := server.app.Listen(addr); err != nil {
		server.logger.Fatal("error resolving server", zap.Error(err))
	}
}
