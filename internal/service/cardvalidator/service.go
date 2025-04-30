package cardvalidator

type CardValidator interface {
}

type Service struct {
	//in the future, we can add database there
}

func NewService() *Service {
	return &Service{}
}
