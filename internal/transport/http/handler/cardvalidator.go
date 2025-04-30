package handler

import (
	"card-validator-api/internal/service"
	"card-validator-api/internal/service/cardvalidator"
	"card-validator-api/pkg/bind"
	"card-validator-api/pkg/server"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
)

type CardHandler struct {
	service *service.Service
}

func NewCardHandler(s *service.Service) *CardHandler {
	return &CardHandler{
		service: s,
	}
}

func (s *CardHandler) Register(server *server.Server) {
	group := server.Group("/card")

	group.POST("/validate", s.Validate)
}

// @Summary Validate a card
// @Description Validates a card's number, expiration month, and year
// @Tags Card
// @ID validate-card
// @Accept json
// @Produce json
// @Param card body cardvalidator.CardValidationRequest true "Card Details"
// @Success 200 {object} cardvalidator.CardValidationResponse
// @Failure 400 {object} echo.HTTPError "Invalid request"
// @Failure 500 {object} echo.HTTPError "Internal server error"
// @Router /card/validate [post]
func (s *CardHandler) Validate(c echo.Context) error {
	var (
		err error
		obj cardvalidator.CardValidationRequest
	)

	if err = bind.Validate(c, &obj); err != nil {
		zap.L().Error("error binding and validating request", zap.Error(err))
		return err
	}

	res, err := s.service.CardValidator.Validate(c.Request().Context(), &obj)
	if err != nil {
		zap.L().Error("error validating card", zap.Error(err))
		return err
	}

	return c.JSON(http.StatusCreated, res)
}
