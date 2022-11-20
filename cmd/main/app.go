package main

import (
	"context"
	"errors"
	"github.com/joho/godotenv"
	"github.com/onemgvv/fakapi/internal/config"
	deliveryHttp "github.com/onemgvv/fakapi/internal/delivery/http"
	"github.com/onemgvv/fakapi/internal/logger"
	"github.com/onemgvv/fakapi/internal/server"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	configDir = "configs"
	timeout   = 5 * time.Second
)

func main() {
	if err := godotenv.Load(); err != nil {
		logger.ErrorLogger.Fatalf("[Config INIT ERROR]: %v", err)
	}

	cfg, err := config.Init(configDir, config.Mode(os.Args[1:][0]))
	if err != nil {
		logger.ErrorLogger.Fatalf("[Config INIT ERROR]: %v", err)
	}

	handler := deliveryHttp.NewHandler()
	app := server.NewServer(cfg, handler.InitRoutes())

	go func() {
		if err := app.Run(); errors.Is(err, http.ErrServerClosed) {
			logger.ErrorLogger.Fatalf("[Server Starting Error]: %s", err.Error())
		}
	}()

	logger.InfoLogger.Infof("[SUCCESSFUL START]: Application started on port: %s", cfg.HTTP.Port)

	/**
	 *	Graceful Shutdown
	 */

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err = app.Stop(ctx); err != nil {
		logger.ErrorLogger.Fatalf("[SHUTDOWN APP ERROR]: %v", err)
	}

	logger.InfoLogger.Infof("[SHUTTING DOWN]: Application successfully stoped")
}
