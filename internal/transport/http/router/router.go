package router

import (
	"card-validator-api/internal/service"
	"card-validator-api/internal/transport/http/handler"
	"card-validator-api/pkg/server"
	"github.com/labstack/echo/v4"
	"net/http"
)

func RegisterRoutes(server *server.Server, services *service.Service) {
	// Register handlers
	handler.NewCardHandler(services).Register(server)

	server.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
}
