package bootstrap

import (
	"card-validator-api/internal/config"
	"card-validator-api/internal/service"
	"card-validator-api/pkg/logger"
	"card-validator-api/pkg/server"
	"context"
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
}
