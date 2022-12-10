package strmech

import (
	"math/big"
	"sync"
)

// TextBigFloatFieldFormatDto
//
// The TextBigFloatFieldFormatDto type encapsulates input
// specifications for a text field populated with a
// big.Float floating point value formatted as a number
// string.
//
// Field Format Data Transfer Objects (Dto) are used
// to facilitate easy data entry which creating and
// configuring text lines strings for screen display,
// file output or printing.
type TextBigFloatFieldFormatDto struct {
	LeftMarginStr string
	//	One or more characters used to create a left
	//	margin for the 'BigFloatNum' Text Field.
	//
	//	If this parameter is set to an empty string, no
	//	left margin will be configured for this
	//	'BigFloatNum' Text Field.

	BigFloatNum big.Float
	// The big.Float floating point number to
	// be formatted for output as a text string.

	RoundingMode big.RoundingMode
	// The rounding mode used to round 'BigFloatNum'
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
	// big.Float floating point number, 'BigFloatNum'.
	//
	// If this value is set to minus one (-1), all
	// available fractional digits to the right of the
	// decimal point will be displayed

	FieldLength int
	//	The length of the text field in which the
	//	'BigFloatNum' string will be displayed. If
	//	'FieldLength' is less than the length of the
	//	'BigFloatNum' string, it will be automatically
	//	set equal to the 'BigFloatNum' string length.
	//
	//	To automatically set the value of 'FieldLength'
	//	to the length of the 'BigFloatNum' string, set
	//	this parameter to a value of minus one (-1).
	//
	//	If this parameter is submitted with a value less
	//	than minus one (-1) or greater than 1-million
	//	(1,000,000), an error will be returned.
	//
	//	Field Length Examples
	//
	//		Example-1
	//          BigFloatNum String = "5672.12345678901234"
	//			BigFloatNum String Length = 19
	//			FieldLength = 25
	//			FieldJustify = TxtJustify.Center()
	//			Text Field String =
	//				"   5672.12345678901234   "
	//
	//		Example-2
	//          BigFloatNum String = "5672.12345678901234"
	//			BigFloatNum String Length = 19
	//			FieldLength = 25
	//			FieldJustify = TxtJustify.Right()
	//			Text Field String =
	//				"      5672.12345678901234"
	//
	//		Example-3
	//          BigFloatNum String = "5672.12345678901234"
	//			BigFloatNum String Length = 19
	//			FieldLength = -1
	//			FieldJustify = TxtJustify.Center() // Ignored
	//			Text Field String =
	//				"5672.12345678901234"

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
	//          BigFloatNum String = "5672.12345678901234"
	//			BigFloatNum String Length = 19
	//			FieldLength = 25
	//			FieldJustify = TxtJustify.Center()
	//			Text Field String =
	//				"   5672.12345678901234   "
	//
	//		Example-2
	//          BigFloatNum String = "5672.12345678901234"
	//			BigFloatNum String Length = 19
	//			FieldLength = 25
	//			FieldJustify = TxtJustify.Right()
	//			Text Field String =
	//				"      5672.12345678901234"
	//
	//		Example-3
	//          BigFloatNum String = "5672.12345678901234"
	//			BigFloatNum String Length = 19
	//			FieldLength = -1
	//			FieldJustify = TxtJustify.Center() // Ignored
	//			Text Field String =
	//				"5672.12345678901234"

	RightMarginStr string
	//	One or more characters used to create a right
	//	margin for this 'FieldDateTime' Text Field.
	//
	//	If this parameter is set to an empty string, no
	//	right margin will be configured for this
	//	'FieldDateTime' Text Field.

	lock *sync.Mutex
}

// GetFormattedText
//
// Returns a string containing the formatted text
// representation of the big.Float floating point number
// encapsulated in the current instance of
// TextBigFloatFieldFormatDto.
func (bFloatFmtDto *TextBigFloatFieldFormatDto) GetFormattedText() string {

	bFloatFmtDto.BigFloatNum.SetMode(
		bFloatFmtDto.RoundingMode)

	return bFloatFmtDto.BigFloatNum.Text(
		'f', bFloatFmtDto.NumOfFractionalDigits)
}

// textBigFloatFieldFormatDtoAtom - Provides helper methods for
// TextBigFloatFieldFormatDto.
type textBigFloatFieldFormatDtoAtom struct {
	lock *sync.Mutex
}

// empty
//
// Receives a pointer to an instance of
// TextBigFloatFieldFormatDto and proceeds to set all the
// member variables to their zero or uninitialized
// states.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and reset all pre-existing
//	data values contained within the
//	TextBigFloatFieldFormatDto instance passed as input
//	parameter 'txtBigFloatFieldFmtDto' to their zero or
//	uninitialized states.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	txtBigFloatFieldFmtDto		*TextBigFloatFieldFormatDto
//
//		A pointer to an instance of TextBigFloatFieldFormatDto.
//		All data values contained within this instance
//		will be deleted and reset to their zero or
//		uninitialized states.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	NONE
func (txtBigFloatFieldFmtDtoAtom *textBigFloatFieldFormatDtoAtom) empty(
	txtBigFloatFieldFmtDto *TextBigFloatFieldFormatDto) {

	if txtBigFloatFieldFmtDtoAtom.lock == nil {
		txtBigFloatFieldFmtDtoAtom.lock = new(sync.Mutex)
	}

	txtBigFloatFieldFmtDtoAtom.lock.Lock()

	defer txtBigFloatFieldFmtDtoAtom.lock.Unlock()

	if txtBigFloatFieldFmtDto == nil {

		return
	}

	txtBigFloatFieldFmtDto.LeftMarginStr = ""

	txtBigFloatFieldFmtDto.BigFloatNum.SetInt64(0)

	txtBigFloatFieldFmtDto.RoundingMode =
		big.ToNearestEven

	txtBigFloatFieldFmtDto.NumOfFractionalDigits = 0

	txtBigFloatFieldFmtDto.FieldLength = 0

	txtBigFloatFieldFmtDto.FieldJustify = 0

	txtBigFloatFieldFmtDto.RightMarginStr = ""

	return
}

// equal
//
// Compares two instances of TextBigFloatFieldFormatDto
// and returns a boolean value signaling whether the two
// instances are equivalent in all respects.
//
// If the two instances of TextBigFloatFieldFormatDto are
// equal, this method returns 'true'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	txtBigFloatFieldFmtDtoOne		*TextBigFloatFieldFormatDto
//
//		A pointer to an instance of
//		TextBigFloatFieldFormatDto.
//
//		The data values contained within this instance
//		will be compared to corresponding data values
//		contained within a second
//		TextBigFloatFieldFormatDto instance
//		('txtBigFloatFieldFmtDtoTwo') in order to
//		determine if they are equivalent.
//
//	txtBigFloatFieldFmtDtoTwo		*TextBigFloatFieldFormatDto
//
//		A pointer to the second of two instances of
//		TextBigFloatFieldFormatDto. The data values
//		contained within this instance will be compared
//		to corresponding data values contained within the
//		first TextBigFloatFieldFormatDto instance
//		('txtBigFloatFieldFmtDtoOne') in order to
//		determine if they are equivalent.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		If all the data values within input parameters
//		'txtBigFloatFieldFmtDtoOne' and
//		'txtBigFloatFieldFmtDtoTwo' are found to be
//		equivalent in all respects, this return parameter
//		will be set to 'true'.
//
//		If the compared data values are NOT equivalent,
//		this method returns 'false'.
func (txtBigFloatFieldFmtDtoAtom *textBigFloatFieldFormatDtoAtom) equal(
	txtBigFloatFieldFmtDtoOne *TextBigFloatFieldFormatDto,
	txtBigFloatFieldFmtDtoTwo *TextBigFloatFieldFormatDto) bool {

	if txtBigFloatFieldFmtDtoAtom.lock == nil {
		txtBigFloatFieldFmtDtoAtom.lock = new(sync.Mutex)
	}

	txtBigFloatFieldFmtDtoAtom.lock.Lock()

	defer txtBigFloatFieldFmtDtoAtom.lock.Unlock()

	if txtBigFloatFieldFmtDtoOne == nil ||
		txtBigFloatFieldFmtDtoTwo == nil {

		return false
	}

	if txtBigFloatFieldFmtDtoOne.LeftMarginStr !=
		txtBigFloatFieldFmtDtoOne.LeftMarginStr {

		return false
	}

	if txtBigFloatFieldFmtDtoOne.RoundingMode !=
		txtBigFloatFieldFmtDtoOne.RoundingMode {

		return false
	}

	if txtBigFloatFieldFmtDtoOne.NumOfFractionalDigits !=
		txtBigFloatFieldFmtDtoOne.NumOfFractionalDigits {

		return false
	}

	if txtBigFloatFieldFmtDtoOne.FieldLength !=
		txtBigFloatFieldFmtDtoOne.FieldLength {

		return false
	}

	if txtBigFloatFieldFmtDtoOne.FieldJustify !=
		txtBigFloatFieldFmtDtoOne.FieldJustify {

		return false
	}

	if txtBigFloatFieldFmtDtoOne.RightMarginStr !=
		txtBigFloatFieldFmtDtoOne.RightMarginStr {

		return false
	}

	var bFloatNumStrOne, bFloatNumStrTwo string
	var tempBFloatOne, tempBFloatTwo big.Float

	tempBFloatOne.Copy(
		&txtBigFloatFieldFmtDtoOne.BigFloatNum)

	tempBFloatOne.SetMode(
		txtBigFloatFieldFmtDtoOne.RoundingMode)

	bFloatNumStrOne = tempBFloatOne.Text(
		'f',
		txtBigFloatFieldFmtDtoOne.NumOfFractionalDigits)

	tempBFloatTwo.Copy(
		&txtBigFloatFieldFmtDtoTwo.BigFloatNum)

	tempBFloatTwo.SetMode(
		txtBigFloatFieldFmtDtoTwo.RoundingMode)

	bFloatNumStrTwo = tempBFloatTwo.Text(
		'f',
		txtBigFloatFieldFmtDtoTwo.NumOfFractionalDigits)

	if bFloatNumStrOne != bFloatNumStrTwo {
		return false
	}

	return true
}
