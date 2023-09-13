package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"noneland/backend/interview/config"
	"noneland/backend/interview/internal/pkg/logger"
)

func main() {
	// initial logger
	logger.InitZap()
	// initial global config
	config.Init()

	app := InitServer()
	server := app.Start()
	go func() {
		logger.Zap().Infof("app Server start at %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Zap().Fatalf("start follow server failed, err: %s", err)
		}
	}()

	// wait for signal
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	logger.Zap().Info("app Server is shutting down....")

	//cancel background task
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Zap().Errorf("app Server shutdown error: %s", err)
	}

	logger.Zap().Info("app Server shutdown complete")
}
