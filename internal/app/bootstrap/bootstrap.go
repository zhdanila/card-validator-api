package bootstrap

import (
	"card-validator-api/internal/config"
	"card-validator-api/internal/service"
	"card-validator-api/internal/transport/http/router"
	"card-validator-api/pkg/logger"
	"card-validator-api/pkg/server"
	"context"
	"errors"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Website() {
	logger.InitLogger()

	ctx := context.Background()
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	services := service.NewService()
	srv := server.NewServer(cfg.HTTPPort)
	router.RegisterRoutes(srv, services)

	zap.L().Sugar().Infof("Finly backend started on port %s", cfg.HTTPPort)
	go func() {
		zap.L().Sugar().Info("Starting server...")
		if err = srv.Start(); err != nil && !errors.Is(http.ErrServerClosed, err) {
			zap.L().Sugar().Fatalf("error with starting server: %s", err.Error())
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	zap.L().Sugar().Info("Finly backend shutting down")

	if err = srv.Shutdown(ctx); err != nil {
		zap.L().Sugar().Fatalf("error with shutting down server: %s", err.Error())
	}
}
