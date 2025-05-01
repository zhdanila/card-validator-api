package cardvalidator

import (
	"fmt"
	"reflect"
	"testing"
)

// TestNormalizeCardNumber tests the normalizeCardNumber function
func TestNormalizeCardNumber(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Card number with spaces",
			input:    "4532 0151 1283 0366",
			expected: "4532015112830366",
		},
		{
			name:     "Card number with hyphens",
			input:    "4532-0151-1283-0366",
			expected: "4532015112830366",
		},
		{
			name:     "Card number with mixed spaces and hyphens",
			input:    "4532 0151-1283 0366",
			expected: "4532015112830366",
		},
		{
			name:     "Clean card number",
			input:    "4532015112830366",
			expected: "4532015112830366",
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := normalizeCardNumber(tt.input)
			if result != tt.expected {
				t.Errorf("normalizeCardNumber(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

// TestIsValidCardLength tests the isValidCardLength function
func TestIsValidCardLength(t *testing.T) {
	tests := []struct {
		name     string
		number   string
		length   int
		expected bool
	}{
		// Visa tests
		{
			name:     "Valid Visa 16 digits",
			number:   "4532015112830366",
			length:   16,
			expected: true,
		},
		{
			name:     "Valid Visa 13 digits",
			number:   "4532015112830",
			length:   13,
			expected: true,
		},
		{
			name:     "Invalid Visa length",
			number:   "453201511283036",
			length:   15,
			expected: false,
		},
		// Mastercard tests
		{
			name:     "Valid Mastercard 16 digits",
			number:   "5532015112830366",
			length:   16,
			expected: true,
		},
		{
			name:     "Invalid Mastercard length",
			number:   "553201511283036",
			length:   15,
			expected: false,
		},
		// Amex tests
		{
			name:     "Valid Amex 15 digits (prefix 34)",
			number:   "341234567890123",
			length:   15,
			expected: true,
		},
		{
			name:     "Valid Amex 15 digits (prefix 37)",
			number:   "371234567890123",
			length:   15,
			expected: true,
		},
		{
			name:     "Invalid Amex length",
			number:   "3412345678901234",
			length:   16,
			expected: false,
		},
		// JCB tests
		{
			name:     "Valid JCB 15 digits",
			number:   "351234567890123",
			length:   15,
			expected: true,
		},
		{
			name:     "Valid JCB 16 digits",
			number:   "3512345678901234",
			length:   16,
			expected: true,
		},
		{
			name:     "Invalid JCB length",
			number:   "35123456789012",
			length:   14,
			expected: false,
		},
		// Diners Club tests
		{
			name:     "Valid Diners 14 digits (prefix 30)",
			number:   "30123456789012",
			length:   14,
			expected: true,
		},
		{
			name:     "Valid Diners 14 digits (prefix 36)",
			number:   "36123456789012",
			length:   14,
			expected: true,
		},
		{
			name:     "Valid Diners 14 digits (prefix 38)",
			number:   "38123456789012",
			length:   14,
			expected: true,
		},
		{
			name:     "Invalid Diners length",
			number:   "301234567890123",
			length:   15,
			expected: false,
		},
		// Discover tests
		{
			name:     "Valid Discover 16 digits",
			number:   "6012345678901234",
			length:   16,
			expected: true,
		},
		{
			name:     "Invalid Discover length",
			number:   "601234567890123",
			length:   15,
			expected: false,
		},
		// Edge cases
		{
			name:     "Empty number",
			number:   "",
			length:   0,
			expected: false,
		},
		{
			name:     "Unknown card type",
			number:   "91234567890123",
			length:   14,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidCardLength(tt.number, tt.length)
			if result != tt.expected {
				t.Errorf("isValidCardLength(%q, %d) = %v; want %v", tt.number, tt.length, result, tt.expected)
			}
		})
	}
}

// TestIsValidLuhn tests the isValidLuhn function
func TestIsValidLuhn(t *testing.T) {
	tests := []struct {
		name     string
		number   string
		expected bool
	}{
		{
			name:     "Valid Luhn (Visa)",
			number:   "4532015112830366",
			expected: true,
		},
		{
			name:     "Valid Luhn (Mastercard)",
			number:   "5555555555554444",
			expected: true,
		},
		{
			name:     "Invalid Luhn",
			number:   "4532015112830367",
			expected: false,
		},
		{
			name:     "Non-numeric input",
			number:   "453201511283036a",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidLuhn(tt.number)
			if result != tt.expected {
				t.Errorf("isValidLuhn(%q) = %v; want %v", tt.number, result, tt.expected)
			}
		})
	}
}

// TestIsValidExpirationDate tests the isValidExpirationDate function
func TestIsValidExpirationDate(t *testing.T) {
	tests := []struct {
		name          string
		expMonth      int
		expYear       int
		expectedValid bool
		expectedErr   error
	}{
		{
			name:          "Valid future date",
			expMonth:      12,
			expYear:       2026,
			expectedValid: true,
			expectedErr:   nil,
		},
		{
			name:          "Valid current year, future month",
			expMonth:      6,
			expYear:       2025,
			expectedValid: true,
			expectedErr:   nil,
		},
		{
			name:          "Invalid month (0)",
			expMonth:      0,
			expYear:       2026,
			expectedValid: false,
			expectedErr:   fmt.Errorf("invalid expiration month"),
		},
		{
			name:          "Invalid month (13)",
			expMonth:      13,
			expYear:       2026,
			expectedValid: false,
			expectedErr:   fmt.Errorf("invalid expiration month"),
		},
		{
			name:          "Expired year",
			expMonth:      12,
			expYear:       2024,
			expectedValid: false,
			expectedErr:   fmt.Errorf("invalid expiration year"),
		},
		{
			name:          "Expired month in current year",
			expMonth:      3,
			expYear:       2025,
			expectedValid: false,
			expectedErr:   fmt.Errorf("card has expired"),
		},
		{
			name:          "Year too far in future",
			expMonth:      12,
			expYear:       2036,
			expectedValid: false,
			expectedErr:   fmt.Errorf("invalid expiration year"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, err := isValidExpirationDate(tt.expMonth, tt.expYear)
			if valid != tt.expectedValid {
				t.Errorf("isValidExpirationDate(%d, %d) valid = %v; want %v", tt.expMonth, tt.expYear, valid, tt.expectedValid)
			}
			if !reflect.DeepEqual(err, tt.expectedErr) {
				t.Errorf("isValidExpirationDate(%d, %d) error = %v; want %v", tt.expMonth, tt.expYear, err, tt.expectedErr)
			}
		})
	}
}

// TestValidateCardNumber tests the validateCardNumber function
func TestValidateCardNumber(t *testing.T) {
	tests := []struct {
		name        string
		number      string
		expectedErr error
	}{
		{
			name:        "Valid Visa number",
			number:      "4532015112830366",
			expectedErr: nil,
		},
		{
			name:        "Invalid non-numeric",
			number:      "453201511283036a",
			expectedErr: fmt.Errorf("invalid card number: must contain only digits"),
		},
		{
			name:        "Invalid length for Visa",
			number:      "453201511283036",
			expectedErr: fmt.Errorf("invalid card number length for detected card type"),
		},
		{
			name:        "Invalid Luhn",
			number:      "4532015112830367",
			expectedErr: fmt.Errorf("invalid card number: fails Luhn Algorithm check"),
		},
		{
			name:        "Empty number",
			number:      "",
			expectedErr: fmt.Errorf("invalid card number length for detected card type"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateCardNumber(tt.number)
			if !reflect.DeepEqual(err, tt.expectedErr) {
				t.Errorf("validateCardNumber(%q) error = %v; want %v", tt.number, err, tt.expectedErr)
			}
		})
	}
}

// TestMaskCardNumber tests the maskCardNumber function
func TestMaskCardNumber(t *testing.T) {
	tests := []struct {
		name     string
		number   string
		expected string
	}{
		{
			name:     "Normal card number",
			number:   "4532015112830366",
			expected: "****-****-****-0366",
		},
		{
			name:     "Short number (<4 digits)",
			number:   "123",
			expected: "****",
		},
		{
			name:     "Exactly 4 digits",
			number:   "1234",
			expected: "****-****-****-1234",
		},
		{
			name:     "Empty string",
			number:   "",
			expected: "****",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maskCardNumber(tt.number)
			if result != tt.expected {
				t.Errorf("maskCardNumber(%q) = %q; want %q", tt.number, result, tt.expected)
			}
		})
	}
}
