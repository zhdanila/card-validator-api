package cardvalidator

type CardValidationRequest struct {
	Number   string `json:"number" validate:"required"`
	ExpMonth int    `json:"exp_month" validate:"required"`
	ExpYear  int    `json:"exp_year" validate:"required"`
}

type CardValidationResponse struct {
	Valid bool                 `json:"valid"`
	Error *CardValidationError `json:"error,omitempty"`
}

type CardValidationError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
