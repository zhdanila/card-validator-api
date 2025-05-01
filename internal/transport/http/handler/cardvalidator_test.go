package handler

import (
	"bytes"
	"card-validator-api/internal/service"
	"card-validator-api/internal/service/cardvalidator"
	"card-validator-api/internal/service/cardvalidator/mock"
	"card-validator-api/pkg/validator"
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupCardHandlerTest(t *testing.T) (*echo.Echo, *mocks.MockCardValidator, *CardHandler) {
	var err error

	ctrl := gomock.NewController(t)
	mockCardValidator := mocks.NewMockCardValidator(ctrl)
	service := &service.Service{CardValidator: mockCardValidator}
	handler := NewCardHandler(service)
	e := echo.New()

	if e.Validator, err = validator.CustomValidator(); err != nil {
		zap.L().Fatal("Error setting up custom validator", zap.Error(err))
	}

	return e, mockCardValidator, handler
}

func TestCardHandler_Validate(t *testing.T) {
	e, mockCardValidator, handler := setupCardHandlerTest(t)
	defer gomock.NewController(t).Finish()

	tests := []struct {
		name              string
		input             cardvalidator.CardValidationRequest
		mockResponse      *cardvalidator.CardValidationResponse
		mockError         error
		expectServiceCall bool
		expectedStatus    int
	}{
		{
			name: "successful card validation",
			input: cardvalidator.CardValidationRequest{
				Number:   "4532015112830366",
				ExpMonth: 12,
				ExpYear:  2025,
			},
			mockResponse: &cardvalidator.CardValidationResponse{
				Valid: true,
			},
			mockError:         nil,
			expectServiceCall: true,
			expectedStatus:    http.StatusCreated,
		},
		{
			name: "invalid card number",
			input: cardvalidator.CardValidationRequest{
				Number:   "1234567890123456",
				ExpMonth: 12,
				ExpYear:  2025,
			},
			mockResponse:      nil,
			mockError:         errors.New("invalid card number"),
			expectServiceCall: true,
			expectedStatus:    http.StatusBadRequest,
		},
		{
			name: "invalid expiration date",
			input: cardvalidator.CardValidationRequest{
				Number:   "4532015112830366",
				ExpMonth: 13,
				ExpYear:  2020,
			},
			mockResponse:      nil,
			mockError:         nil,
			expectServiceCall: false,
			expectedStatus:    http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.input)
			req := httptest.NewRequest(http.MethodPost, "/card/validate", bytes.NewReader(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if tt.expectServiceCall {
				mockCardValidator.EXPECT().
					Validate(gomock.Any(), gomock.Any()).
					Return(tt.mockResponse, tt.mockError)
			}

			err := handler.Validate(c)

			if tt.expectedStatus == http.StatusBadRequest {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, rec.Code)

			if tt.mockResponse != nil {
				var response cardvalidator.CardValidationResponse
				err = json.Unmarshal(rec.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, tt.mockResponse.Valid, response.Valid)
			}
		})
	}
}
