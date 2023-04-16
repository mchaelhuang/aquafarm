package main

import (
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"github.com/mchaelhuang/aquafarm/internal/provider"
)

func main() {
	logger := provider.LoggerProvider()
	defer logger.Sync()

	app := provider.ProvideRESTApp()

	go func() {
		if err := app.Run(); err != nil {
			logger.Fatal("run is failed", zap.Error(err))
		}
	}()

	// Wait for interrupt
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	logger.Info("shutting down app")

	// Gracefully shutdown
	if err := app.Stop(); err != nil {
		logger.Fatal("error when shutting down app", zap.Error(err))
	}

	logger.Info("bye.")
}
