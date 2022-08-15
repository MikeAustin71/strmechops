package strmech

// SignedNumberFormatSpec - The Signed Number Format
// Specification is used to format signed numeric values
// as number strings. This type contains all the
// specifications and parameters required to format a
// numeric value for text displays.
type SignedNumberFormatSpec struct {
	DecSeparator DecimalSeparatorSpec
	// Contains the decimal separator character or
	// characters which will separate integer and
	// fractional digits in a floating point number

	IntGroupingType NumStrIntegerGroupingSpec
	// This parameter specifies the type of integer
	// grouping and integer separator characters
	// which will be applied to the number string
	// formatting operation.

	RoundingSpec NumStrRoundingSpec
	// Controls the rounding algorithm applied to
	// floating point numbers.

	PositiveNumberSign NumStrPositiveNumberSignSpec
	// Positive number signs are commonly implied and
	// not specified. However, the user as the option
	// to specify a positive number sign character or
	// characters for positive numeric values using a
	// Number String Positive Number Sign
	// Specification.

	NegativeNumberSign NumStrNegativeNumberSignSpec
	// The Number String Negative Number Sign
	// Specification is used to configure negative
	// number sign symbols for negative numeric values
	// formatted and displayed in number stings.

	SignedNumberFieldLength int
	// The length of the text field in which the Signed Number will
	// be displayed. If 'SignedNumberFieldLength' is less than the
	// length of the Signed Number string, it will be automatically
	// set equal to the Signed Number string length.
	//
	// To automatically set the value of SignedNumberFieldLength to
	// the length of the Signed Number string, set this parameter to
	// a value of minus one (-1).
	//
	// If this parameter is submitted with a value less than minus
	// one (-1) or greater than 1-million (1,000,000), an error will
	// be returned.

	SignedNumberJustification TextJustify
	// An enumeration which specifies the justification of the
	// Signed Number string within the field specified by
	// 'SignedNumberFieldLength'.
	//
	// Text justification can only be evaluated in the context of
	// a Signed Number string, field length and a
	// 'textJustification' object of type TextJustify. This is
	// because Signed Number strings with a field length equal to
	// or less than the length of the Signed Number string never
	// use text justification. In these cases, text justification
	// is completely ignored.
	//
	// If the field length ('SignedNumberFieldLength') is greater
	// than the length of the Signed Number string, text
	// justification must be equal to one of these
	// three valid values:
	//           TextJustify(0).Left()
	//           TextJustify(0).Right()
	//           TextJustify(0).Center()
	//
	// You can also use the abbreviated text justification
	// enumeration syntax as follows:
	//
	//           TxtJustify.Left()
	//           TxtJustify.Right()
	//           TxtJustify.Center()

}
