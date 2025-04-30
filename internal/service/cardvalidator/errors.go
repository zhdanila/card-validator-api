package cardvalidator

import (
	"card-validator-api/internal/domain/enums"
	"github.com/labstack/echo/v4"
	"net/http"
)

var errs = struct {
	InvalidCardNumber *echo.HTTPError
	ExpiredCard       *echo.HTTPError
}{
	InvalidCardNumber: echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
		"code":    enums.ErrInvalidCardNumber.String(),
		"message": "Invalid card number: must be 13 to 19 digits",
	}),
	ExpiredCard: echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
		"code":    enums.ErrCardExpired.String(),
		"message": "Card has expired",
	}),
}
