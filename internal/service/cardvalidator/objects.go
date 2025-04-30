package cardvalidator

import "github.com/labstack/echo/v4"

type CardValidationRequest struct {
	Number   string `json:"number" validate:"required"`
	ExpMonth int    `json:"exp_month" validate:"required"`
	ExpYear  int    `json:"exp_year" validate:"required"`
}

type CardValidationResponse struct {
	Valid bool `json:"valid"`
	*echo.HTTPError
}
