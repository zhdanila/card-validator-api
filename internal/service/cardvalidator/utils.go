package cardvalidator

import (
	"card-validator-api/pkg/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	PrefixAmex1      = "34" // American Express card prefix (starts with "34")
	PrefixAmex2      = "37" // American Express card prefix (starts with "37")
	PrefixJCB        = "35" // JCB card prefix (starts with "35")
	PrefixDiners1    = "30" // Diners Club card prefix (starts with "30")
	PrefixDiners2    = "36" // Diners Club card prefix (starts with "36")
	PrefixDiners3    = "38" // Diners Club card prefix (starts with "38")
	PrefixVisa       = '4'  // Visa card prefix (starts with "4")
	PrefixMastercard = '5'  // Mastercard card prefix (starts with "5")
	PrefixDiscover   = '6'  // Discover card prefix (starts with "6")

	LengthAmex       = 15 // American Express card length (15 digits)
	LengthJCB1       = 15 // JCB card length variant 1 (15 digits)
	LengthJCB2       = 16 // JCB card length variant 2 (16 digits)
	LengthDiners     = 14 // Diners Club card length (14 digits)
	LengthVisa1      = 13 // Visa card length variant 1 (13 digits)
	LengthVisa2      = 16 // Visa card length variant 2 (16 digits)
	LengthMastercard = 16 // Mastercard card length (16 digits)
	LengthDiscover   = 16 // Discover card length (16 digits)
)

// normalizeCardNumber removes spaces and hyphens from the card number.
func normalizeCardNumber(cardNumber string) string {
	cardNumber = strings.ReplaceAll(cardNumber, " ", "")
	return strings.ReplaceAll(cardNumber, "-", "")
}

// isValidCardLength checks if the card number length matches the expected length for its type.
func isValidCardLength(number string, length int) bool {
	if len(number) == 0 {
		return false
	}

	switch number[0] {
	case '3':
		if len(number) >= 2 {
			prefix := number[:2]
			switch prefix {
			case PrefixAmex1, PrefixAmex2:
				return length == LengthAmex
			case PrefixJCB:
				return length == LengthJCB1 || length == LengthJCB2
			case PrefixDiners1, PrefixDiners2, PrefixDiners3:
				return length == LengthDiners
			}
		}
	case PrefixVisa:
		return length == LengthVisa1 || length == LengthVisa2
	case PrefixMastercard:
		return length == LengthMastercard
	case PrefixDiscover:
		return length == LengthDiscover
	}
	return false
}

// isValidLuhn implements the Luhn Algorithm to validate the card number.
func isValidLuhn(number string) bool {
	var sum int
	length := len(number)
	isEven := false

	for i := length - 1; i >= 0; i-- {
		digit, err := strconv.Atoi(string(number[i]))
		if err != nil {
			return false
		}

		if isEven {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		isEven = !isEven
	}

	return sum%10 == 0
}

// isValidExpirationDate checks if the expiration date is valid (month/year).
func isValidExpirationDate(expMonth, expYear int) (bool, error) {
	// Validate expiration month
	if expMonth < 1 || expMonth > 12 {
		return false, fmt.Errorf("invalid expiration month")
	}

	// Validate expiration year
	currentYear := time.Now().Year()
	if expYear < currentYear || expYear > currentYear+10 {
		return false, fmt.Errorf("invalid expiration year")
	}

	// Check if the card is expired (month-level precision)
	currentMonth := int(time.Now().Month())
	if expYear == currentYear && expMonth < currentMonth {
		return false, fmt.Errorf("card has expired")
	}

	return true, nil
}

// validateCardNumber checks if the card number is valid using Luhn Algorithm and length rules.
func validateCardNumber(number string) error {
	// Check if number contains only digits
	for _, r := range number {
		if !utils.IsDigit(r) {
			return fmt.Errorf("invalid card number: must contain only digits")
		}
	}

	// Check card length and type
	length := len(number)
	if !isValidCardLength(number, length) {
		return fmt.Errorf("invalid card number length for detected card type")
	}

	// Validate using Luhn Algorithm
	if !isValidLuhn(number) {
		return fmt.Errorf("invalid card number: fails Luhn Algorithm check")
	}

	return nil
}

// maskCardNumber masks all but the last 4 digits of the card number for logging.
func maskCardNumber(number string) string {
	if len(number) < 4 {
		return "****"
	}
	return "****-****-****-" + number[len(number)-4:]
}
