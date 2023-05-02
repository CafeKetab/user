package cmd

import (
	"os"

	"github.com/CafeKetab/user/internal/api/grpc"
	"github.com/CafeKetab/user/internal/api/http"
	"github.com/CafeKetab/user/internal/config"
	"github.com/CafeKetab/user/internal/repository"
	"github.com/CafeKetab/user/pkg/logger"
	"github.com/CafeKetab/user/pkg/rdbms"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

type Server struct{}

func (cmd Server) Command(trap chan os.Signal) *cobra.Command {
	run := func(_ *cobra.Command, _ []string) {
		cmd.main(config.Load(true), trap)
	}

	return &cobra.Command{
		Use:   "server",
		Short: "run user server",
		Run:   run,
	}
}

func (cmd *Server) main(cfg *config.Config, trap chan os.Signal) {
	logger := logger.NewZap(cfg.Logger)

	rdbms, err := rdbms.NewPostgres(cfg.RDBMS)
	if err != nil {
		logger.Panic("Error creating rdbms database", zap.Error(err))
	}

	repo := repository.New(logger, rdbms)
	authGrpcClient := grpc.NewAuthClient(cfg.GRPC, logger)

	server := http.New(cfg.HTTP, logger, repo, authGrpcClient)
	go server.Serve()

	// Keep this at the bottom of the main function
	field := zap.String("signal trap", (<-trap).String())
	logger.Info("exiting by receiving a unix signal", field)
}
