package strmech

import "sync"

// SignedNumberFormatSpec - The Signed Number Format
// Specification is used to format signed numeric values
// as number strings. This type contains all the
// specifications and parameters required to format a
// numeric value for text displays.
type SignedNumberFormatSpec struct {
	decSeparator DecimalSeparatorSpec
	// Contains the decimal separator character or
	// characters which will separate integer and
	// fractional digits in a floating point number

	intGroupingType NumStrIntegerGroupingSpec
	// This parameter specifies the type of integer
	// grouping and integer separator characters
	// which will be applied to the number string
	// formatting operation.

	roundingSpec NumStrRoundingSpec
	// Controls the rounding algorithm applied to
	// floating point numbers.

	positiveNumberSign NumStrPositiveNumberSignSpec
	// Positive number signs are commonly implied and
	// not specified. However, the user as the option
	// to specify a positive number sign character or
	// characters for positive numeric values using a
	// Number String Positive Number Sign
	// Specification.

	negativeNumberSign NumStrNegativeNumberSignSpec
	// The Number String Negative Number Sign
	// Specification is used to configure negative
	// number sign symbols for negative numeric values
	// formatted and displayed in number stings.

	numberField NumStrNumberFieldSpec
	// This Number String Number Field Specification
	// contains the field length and text justification
	// parameter necessary to display a numeric value
	// within a number field for display as a number
	// string.

	lock *sync.Mutex
}
