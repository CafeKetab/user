package cmd

import (
	"os"

	"github.com/CafeKetab/user/internal/config"
	"github.com/CafeKetab/user/pkg/logger"

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

	// crypto := crypto.New(cfg.Crypto)
	// token, err := token.New(cfg.Token)
	// if err != nil {
	// 	logger.Panic("Error creating token object", zap.Error(err))
	// }

	// go grpc.New(logger, crypto, token).Serve(9090)

	// Keep this at the bottom of the main function
	field := zap.String("signal trap", (<-trap).String())
	logger.Info("exiting by receiving a unix signal", field)
}
