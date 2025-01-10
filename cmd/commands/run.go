package commands

import (
	"errors"
	"os"
	"os/signal"
	"syscall"

	"github.com/dezh-tech/geb/cmd/daemon"
	"github.com/dezh-tech/geb/config"
	"github.com/dezh-tech/geb/pkg/logger"
)

func HandleRun(args []string) {
	if len(args) < 3 {
		ExitOnError(errors.New("at least 1 arguments expected\nuse help command for more information"))
	}

	cfg, err := config.Load(args[2])
	if err != nil {
		ExitOnError(err)
	}

	logger.InitGlobalLogger(&cfg.Logger)

	d, err := daemon.New(cfg)
	if err != nil {
		ExitOnError(err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	errCh := d.Start()

	select {
	case sig := <-sigChan:
		logger.Info("Initiating graceful shutdown", "signal", sig.String())
		if err := d.Stop(); err != nil {
			ExitOnError(err)
		}

	case err := <-errCh:
		logger.Error("Initiating shutdown due to the error", "err", err)
		if err := d.Stop(); err != nil {
			ExitOnError(err)
		}
	}
}
