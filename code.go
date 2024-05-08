// Package sex provides a simple implementation for representing and working with gender codes as defined in ISO/IEC 5218.
package sex

import (
	"errors"
)

// GenderCode represents a gender code as per ISO/IEC 5218.
type GenderCode int

const (
	// NotKnown represents a gender code of 'Not known'.
	NotKnown GenderCode = 0
	// Male represents a gender code of 'Male'.
	Male GenderCode = 1
	// Female represents a gender code of 'Female'.
	Female GenderCode = 2
	// NotApplicable represents a gender code of 'Not applicable'.
	NotApplicable GenderCode = 9
)

var (
	// ErrInvalidGenderCode represents the error returned when an invalid code or string is provided.
	ErrInvalidGenderCode = errors.New("invalid gender code")
)

// String returns the string representation of the GenderCode.
func (gc GenderCode) String() string {
	switch gc {
	case NotKnown:
		return "Not Known"
	case Male:
		return "Male"
	case Female:
		return "Female"
	case NotApplicable:
		return "Not Applicable"
	default:
		return "Invalid"
	}
}

// FromInt converts an integer to a GenderCode.
func FromInt(code int) (GenderCode, error) {
	switch code {
	case 0:
		return NotKnown, nil
	case 1:
		return Male, nil
	case 2:
		return Female, nil
	case 9:
		return NotApplicable, nil
	default:
		return -1, ErrInvalidGenderCode
	}
}

// FromString converts a string to a GenderCode.
func FromString(code string) (GenderCode, error) {
	switch code {
	case "Not Known":
		return NotKnown, nil
	case "Male":
		return Male, nil
	case "Female":
		return Female, nil
	case "Not Applicable":
		return NotApplicable, nil
	default:
		return -1, ErrInvalidGenderCode
	}
}
