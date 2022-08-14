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

	IntGroupingType IntegerGroupingType
	// This enumeration specifies the type of inter
	// grouping which will be applied to the integer
	// digits displayed in the number string.
	//
	// Possible Valid Values are:
	//	IntGroupingType.None()
	//	IntGroupingType.Thousands()
	//	IntGroupingType.IndiaNumbering()
	//	IntGroupingType.ChineseNumbering()

	IntGroupingChar string

	RoundingSpec NumStrRoundingSpec
	// Controls the rounding algorithm applied to
	// floating point numbers.

	LeadingPositiveNumberSign string
	// Positive number signs are commonly implied and
	// not specified. However, the user as the option
	// to specify a leading positive number sign
	// character or characters for positive numeric
	// values.

	TrailingPositiveNumberSign string
	// Positive number signs are commonly implied and
	// not specified. However, the user as the option
	// to specify a trailing positive number sign
	// character or characters for positive numeric
	// values.

	LeadingNegativeNumberSign string
	// The leading negative number sign character or
	// characters applied to negative numeric values.
	// Users have the option to configure both leading
	// and trailing negative number signs for negative
	// numeric values.

	TrailingNegativeNumberSign string
	// The trailing negative number sign character or
	// characters applied to negative numeric values.
	// Users have the option to configure both leading
	// and trailing negative number signs for negative
	// numeric values.

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
