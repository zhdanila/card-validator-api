package cardvalidator

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestService_Validate(t *testing.T) {
	// Setup mock logger
	zap.ReplaceGlobals(zap.NewNop())

	// Create service
	svc := NewService()

	// Define test cases
	tests := []struct {
		name           string
		req            *CardValidationRequest
		expectedResp   *CardValidationResponse
		expectedErr    error
		currentTime    time.Time
		modifyTimeFunc func(*testing.T, time.Time)
	}{
		{
			name: "Valid card",
			req: &CardValidationRequest{
				Number:   "4532015112830366", // Valid Visa card number
				ExpMonth: 12,
				ExpYear:  time.Now().Year() + 1,
			},
			expectedResp: &CardValidationResponse{
				Valid:     true,
				HTTPError: nil,
			},
			expectedErr: nil,
			currentTime: time.Now(),
		},
		{
			name: "Invalid card number - too short",
			req: &CardValidationRequest{
				Number:   "453201511283", // Too short
				ExpMonth: 12,
				ExpYear:  time.Now().Year() + 1,
			},
			expectedResp: &CardValidationResponse{
				Valid:     false,
				HTTPError: errs.InvalidCardNumber,
			},
			expectedErr: nil,
			currentTime: time.Now(),
		},
		{
			name: "Invalid card number - non-numeric",
			req: &CardValidationRequest{
				Number:   "453201511283036a", // Contains letter
				ExpMonth: 12,
				ExpYear:  time.Now().Year() + 1,
			},
			expectedResp: &CardValidationResponse{
				Valid:     false,
				HTTPError: errs.InvalidCardNumber,
			},
			expectedErr: nil,
			currentTime: time.Now(),
		},
		{
			name: "Expired card - past month",
			req: &CardValidationRequest{
				Number:   "4532015112830366",
				ExpMonth: 1,
				ExpYear:  time.Now().Year(),
			},
			modifyTimeFunc: func(t *testing.T, currentTime time.Time) {},
			expectedResp: &CardValidationResponse{
				Valid:     false,
				HTTPError: errs.ExpiredCard,
			},
			expectedErr: nil,
			currentTime: time.Now(),
		},
		{
			name: "Invalid expiration month - zero",
			req: &CardValidationRequest{
				Number:   "4532015112830366",
				ExpMonth: 0,
				ExpYear:  time.Now().Year() + 1,
			},
			expectedResp: &CardValidationResponse{
				Valid:     false,
				HTTPError: errs.ExpiredCard,
			},
			expectedErr: nil,
			currentTime: time.Now(),
		},
		{
			name: "Invalid expiration month - too high",
			req: &CardValidationRequest{
				Number:   "4532015112830366",
				ExpMonth: 13,
				ExpYear:  time.Now().Year() + 1,
			},
			expectedResp: &CardValidationResponse{
				Valid:     false,
				HTTPError: errs.ExpiredCard,
			},
			expectedErr: nil,
			currentTime: time.Now(),
		},
		{
			name: "Card number with spaces",
			req: &CardValidationRequest{
				Number:   "4532 0151 1283 0366",
				ExpMonth: 12,
				ExpYear:  time.Now().Year() + 1,
			},
			expectedResp: &CardValidationResponse{
				Valid:     true,
				HTTPError: nil,
			},
			expectedErr: nil,
			currentTime: time.Now(),
		},
	}

	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			// Apply time modification if needed
			if tt.modifyTimeFunc != nil {
				tt.modifyTimeFunc(t, tt.currentTime)
			}

			// Execute validation
			resp, err := svc.Validate(ctx, tt.req)

			// Check error
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.expectedErr)
			}

			// Check response
			if resp.Valid != tt.expectedResp.Valid {
				t.Errorf("Validate() Valid = %v, want %v", resp.Valid, tt.expectedResp.Valid)
			}
			if !errors.Is(resp.HTTPError, tt.expectedResp.HTTPError) {
				t.Errorf("Validate() HTTPError = %v, want %v", resp.HTTPError, tt.expectedResp.HTTPError)
			}
		})
	}
}
