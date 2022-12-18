package strmech

import (
	"math/big"
	"sync"
)

type TextFieldFormatDtoFloat64 struct {
	LeftMarginStr string
	//	One or more characters used to create a left
	//	margin for the 'Float64Num' Text Field.
	//
	//	If this parameter is set to an empty string, no
	//	left margin will be configured for this
	//	'Float64Num' Text Field.

	Float64Num float64
	// The float64 floating point number to
	// be formatted for output as a text string.

	RoundingMode big.RoundingMode
	// The rounding mode used to round 'Float64Num'
	// to the number of fractional digits specified
	// by parameter, 'NumOfFractionalDigits'.
	//
	// Rounding Modes are defined in Golang as follows:
	//
	//	ToNearestEven RoundingMode == IEEE 754-2008 roundTiesToEven
	//	ToNearestAway == IEEE 754-2008 roundTiesToAway
	//	ToZero        == IEEE 754-2008 roundTowardZero
	//	AwayFromZero  == no IEEE 754-2008 equivalent
	//	ToNegativeInf == IEEE 754-2008 roundTowardNegative
	//	ToPositiveInf == IEEE 754-2008 roundTowardPositive

	NumOfFractionalDigits int
	// The number of digits to the right of the radix
	// point (a.k.a. decimal point) which will be
	// displayed in the formatted text string for the
	// Float64 floating point number, 'Float64Num'.
	//
	// If this value is set to minus one (-1), all
	// available fractional digits to the right of the
	// decimal point will be displayed

	FieldLength int
	//	The length of the text field in which the
	//	'Float64Num' string will be displayed. If
	//	'FieldLength' is less than the length of the
	//	'Float64Num' string, it will be automatically
	//	set equal to the 'Float64Num' string length.
	//
	//	To automatically set the value of 'FieldLength'
	//	to the length of the 'Float64Num' string, set
	//	this parameter to a value of minus one (-1).
	//
	//	If this parameter is submitted with a value less
	//	than minus one (-1) or greater than 1-million
	//	(1,000,000), an error will be returned.
	//
	//	Field Length Examples
	//
	//		Example-1
	//          Float64Num String = "5672.1234567"
	//			Float64Num String Length = 12
	//			FieldLength = 18
	//			FieldJustify = TxtJustify.Center()
	//			Text Field String =
	//				"   5672.1234567   "
	//
	//		Example-2
	//          Float64Num String = "5672.1234567"
	//			Float64Num String Length = 12
	//			FieldLength = 18
	//			FieldJustify = TxtJustify.Right()
	//			Text Field String =
	//				"      5672.1234567"
	//
	//		Example-3
	//          Float64Num String = "5672.1234567"
	//			Float64Num String Length = 12
	//			FieldLength = -1
	//			FieldJustify = TxtJustify.Center()
	//				// Text Justification Ignored. Field
	//				// Length Equals Title Line String Length
	//			Text Field String =
	//				"5672.1234567"
	//
	//		Example-4
	//          Float64Num String = "5672.1234567"
	//			Float64Num String Length = 12
	//			FieldLength = 2
	//			FieldJustify = TxtJustify.Center()
	//				// Justification Ignored because Field
	//				// Length Less Than Title Line String Length.
	//			Text Field String =
	//				"5672.1234567"

	FieldJustify TextJustify
	//	An enumeration which specifies the justification
	//	of the 'FieldDateTime' string within the text
	//	field length specified by 'FieldLength'.
	//
	//	Text justification can only be evaluated in the
	//	context of a text label ('FieldDateTime'), field
	//	length ('FieldLength') and a Text Justification
	//	object of type TextJustify. This is because text
	//	labels with a field length equal to or less than
	//	the length of the text label string will never
	//	use text justification. In these cases, text
	//	justification is completely ignored.
	//
	//	If the field length is greater than the length of
	//	the text label string, text justification must be
	//	equal to one of these three valid values:
	//
	//	    TextJustify(0).Left()
	//	    TextJustify(0).Right()
	//	    TextJustify(0).Center()
	//
	//	Users can also specify the abbreviated text
	//	justification enumeration syntax as follows:
	//
	//	    TxtJustify.Left()
	//	    TxtJustify.Right()
	//	    TxtJustify.Center()
	//
	//	Text Justification Examples
	//
	//		Example-1
	//          Float64Num String = "5672.1234567"
	//			Float64Num String Length = 12
	//			FieldLength = 18
	//			FieldJustify = TxtJustify.Center()
	//			Text Field String =
	//				"   5672.1234567   "
	//
	//		Example-2
	//          Float64Num String = "5672.1234567"
	//			Float64Num String Length = 12
	//			FieldLength = 18
	//			FieldJustify = TxtJustify.Right()
	//			Text Field String =
	//				"      5672.1234567"
	//
	//		Example-3
	//          Float64Num String = "5672.1234567"
	//			Float64Num String Length = 12
	//			FieldLength = -1
	//			FieldJustify = TxtJustify.Center()
	//				// Text Justification Ignored. Field
	//				// Length Equals Title Line String Length
	//			Text Field String =
	//				"5672.1234567"
	//
	//		Example-4
	//          Float64Num String = "5672.1234567"
	//			Float64Num String Length = 12
	//			FieldLength = 2
	//			FieldJustify = TxtJustify.Center()
	//				// Justification Ignored because Field
	//				// Length Less Than Title Line String Length.
	//			Text Field String =
	//				"5672.1234567"

	RightMarginStr string
	//	One or more characters used to create a right
	//	margin for this 'FieldDateTime' Text Field.
	//
	//	If this parameter is set to an empty string, no
	//	right margin will be configured for this
	//	'FieldDateTime' Text Field.

	lock *sync.Mutex
}

// textFieldFormatDtoFloat64Atom
//
// Provides helper methods for TextFieldFormatDtoFloat64.
type textFieldFormatDtoFloat64Atom struct {
	lock *sync.Mutex
}

// empty
//
// Receives a pointer to an instance of
// TextFieldFormatDtoFloat64 and proceeds to set all the
// member variables to their zero or uninitialized
// states.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and reset all pre-existing
//	data values contained within the
//	TextFieldFormatDtoFloat64 instance passed as input
//	parameter 'txtBigFloatFieldFmtDto' to their zero or
//	uninitialized states.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	txtFieldFmtDtoFloat64		*TextFieldFormatDtoFloat64
//
//		A pointer to an instance of
//		TextFieldFormatDtoFloat64. All data values
//		contained within this instance will be deleted
//		and reset to their zero or uninitialized states.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	NONE
func (txtFieldFmtFloat64DtoAtom *textFieldFormatDtoFloat64Atom) empty(
	txtFieldFmtDtoFloat64 *TextFieldFormatDtoFloat64) {

	if txtFieldFmtFloat64DtoAtom.lock == nil {
		txtFieldFmtFloat64DtoAtom.lock = new(sync.Mutex)
	}

	txtFieldFmtFloat64DtoAtom.lock.Lock()

	defer txtFieldFmtFloat64DtoAtom.lock.Unlock()

	if txtFieldFmtDtoFloat64 == nil {

		return
	}

	txtFieldFmtDtoFloat64.LeftMarginStr = ""

	txtFieldFmtDtoFloat64.Float64Num = 0.0

	txtFieldFmtDtoFloat64.RoundingMode =
		big.ToNearestEven

	txtFieldFmtDtoFloat64.NumOfFractionalDigits = 0

	txtFieldFmtDtoFloat64.FieldLength = 0

	txtFieldFmtDtoFloat64.FieldJustify = TxtJustify.None()

	txtFieldFmtDtoFloat64.RightMarginStr = ""

	return
}
