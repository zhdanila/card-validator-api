package cardvalidator

import (
	"context"
	"go.uber.org/zap"
)

type CardValidator interface {
	Validate(ctx context.Context, req *CardValidationRequest) (*CardValidationResponse, error)
}

type Service struct {
	//in the future, we can add repository there
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Validate(ctx context.Context, req *CardValidationRequest) (*CardValidationResponse, error) {
	// Normalize card number
	cardNumber := normalizeCardNumber(req.Number)
	maskedNumber := maskCardNumber(cardNumber) // Mask for logging

	// Validate card number
	if err := validateCardNumber(cardNumber); err != nil {
		zap.L().Sugar().Error("Card number validation failed",
			zap.String("card_number", maskedNumber),
			zap.Error(err),
		)

		return &CardValidationResponse{
			Valid:     false,
			HTTPError: errs.InvalidCardNumber,
		}, nil
	}

	// Validate expiration date (month/year)
	if valid, err := isValidExpirationDate(req.ExpMonth, req.ExpYear); !valid {
		zap.L().Sugar().Error(err.Error(),
			zap.String("card_number", maskedNumber),
			zap.Int("exp_month", req.ExpMonth),
			zap.Int("exp_year", req.ExpYear),
		)
		return &CardValidationResponse{
			Valid:     false,
			HTTPError: errs.ExpiredCard,
		}, nil
	}

	// Log success
	zap.L().Sugar().Info("Card validated successfully",
		zap.String("card_number", maskedNumber),
	)
	return &CardValidationResponse{
		Valid: true,
	}, nil
}
