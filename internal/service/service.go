package service

import "card-validator-api/internal/service/cardvalidator"

type Service struct {
	CardValidator cardvalidator.CardValidator
}

func NewService() *Service {
	return &Service{
		CardValidator: cardvalidator.NewService(),
	}
}
