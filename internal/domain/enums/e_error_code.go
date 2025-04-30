package enums

type Enum string

const (
	ErrInvalidCardNumber Enum = "001"
	ErrCardExpired       Enum = "002"
)

func (r Enum) IsValid() bool {
	switch r {
	case ErrCardExpired, ErrInvalidCardNumber:
		return true
	default:
		return false
	}
}

func (r Enum) String() string {
	return string(r)
}
