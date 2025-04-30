package cardvalidator

import "context"

type CardValidator interface {
	Validate(ctx context.Context, req *CardValidationRequest) (*CardValidationResponse, error)
}

type Service struct {
	//in the future, we can add database there
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Validate(ctx context.Context, req *CardValidationRequest) (*CardValidationResponse, error) {
	return nil, nil
}
