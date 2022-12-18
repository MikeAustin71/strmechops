package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
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

	RoundingType NumberRoundingType
	//	This enumeration parameter is used to specify the
	//	type of rounding algorithm that will be applied for
	//	the	rounding of fractional digits contained in the
	//	'Float64Num' value.
	//
	//	If in doubt as to a suitable rounding method,
	//	'HalfAwayFromZero' is recommended.
	//
	//	Possible values are listed as follows:
	//		NumRoundType.None() - Invalid Value
	//
	//		* Valid Values *
	//		NumRoundType.NoRounding()
	//		NumRoundType.HalfUpWithNegNums()
	//		NumRoundType.HalfDownWithNegNums()
	//		NumRoundType.HalfAwayFromZero()
	//		NumRoundType.HalfTowardsZero()
	//		NumRoundType.HalfToEven()
	//		NumRoundType.HalfToOdd()
	//		NumRoundType.Randomly()
	//		NumRoundType.Floor()
	//		NumRoundType.Ceiling()
	//		NumRoundType.Truncate()
	//
	//	Definitions:
	//
	//		NoRounding
	//
	//			Signals that no rounding operation will be
	//			performed on fractional digits. The
	//			fractional digits will therefore remain
	//			unchanged.
	//
	//		HalfUpWithNegNums
	//
	//			Half Round Up Including Negative Numbers.
	//			This method is intuitive but may produce
	//			unexpected results when applied to negative
	//			numbers.
	//
	//			'HalfUpWithNegNums' rounds .5 up.
	//
	//				Examples of 'HalfUpWithNegNums'
	//				7.6 rounds up to 8
	//				7.5 rounds up to 8
	//				7.4 rounds down to 7
	//				-7.4 rounds up to -7
	//				-7.5 rounds up to -7
	//				-7.6 rounds down to -8
	//
	//		HalfDownWithNegNums
	//
	//		Half Round Down Including Negative Numbers. This
	//		method is also considered intuitive but may
	//		produce unexpected results when applied to
	//		negative numbers.
	//
	//		'HalfDownWithNegNums' rounds .5 down.
	//
	//			Examples of HalfDownWithNegNums
	//
	//			7.6 rounds up to 8
	//			7.5 rounds down to 7
	//			7.4 rounds down to 7
	//			-7.4 rounds up to -7
	//			-7.5 rounds down to -8
	//			-7.6 rounds down to -8
	//
	//		HalfAwayFromZero
	//
	//			The 'HalfAwayFromZero' method rounds .5 further
	//			away from zero.	It provides clear and consistent
	//			behavior when dealing with negative numbers.
	//
	//				Examples of HalfAwayFromZero
	//
	//				7.6 rounds away to 8
	//				7.5 rounds away to 8
	//				7.4 rounds to 7
	//				-7.4 rounds to -7
	//				-7.5 rounds away to -8
	//				-7.6 rounds away to -8
	//
	//		HalfTowardsZero
	//
	//			Round Half Towards Zero. 'HalfTowardsZero' rounds
	//			0.5	closer to zero. It provides clear and
	//			consistent behavior	when dealing with negative
	//			numbers.
	//
	//				Examples of HalfTowardsZero
	//
	//				7.6 rounds away to 8
	//				7.5 rounds to 7
	//				7.4 rounds to 7
	//				-7.4 rounds to -7
	//				-7.5 rounds to -7
	//				-7.6 rounds away to -8
	//
	//		HalfToEven
	//
	//			Round Half To Even Numbers. 'HalfToEven' is
	//			also called	Banker's Rounding. This method
	//			rounds 0.5 to the nearest even digit.
	//
	//				Examples of HalfToEven
	//
	//				7.5 rounds up to 8 (because 8 is an even
	//				number)	but 6.5 rounds down to 6 (because
	//				6 is an even number)
	//
	//				HalfToEven only applies to 0.5. Other
	//				numbers (not ending	in 0.5) round to
	//				nearest as usual, so:
	//
	//				7.6 rounds up to 8
	//				7.5 rounds up to 8 (because 8 is an even number)
	//				7.4 rounds down to 7
	//				6.6 rounds up to 7
	//				6.5 rounds down to 6 (because 6 is an even number)
	//				6.4 rounds down to 6
	//
	//		HalfToOdd
	//
	//			Round Half to Odd Numbers. Similar to 'HalfToEven',
	//			but in this case 'HalfToOdd' rounds 0.5 towards odd
	//			numbers.
	//
	//				Examples of HalfToOdd
	//
	//				HalfToOdd only applies to 0.5. Other numbers
	//				(not ending	in 0.5) round to nearest as usual.
	//
	//				7.5 rounds down to 7 (because 7 is an odd number)
	//
	//				6.5 rounds up to 7 (because 7 is an odd number)
	//
	//				7.6 rounds up to 8
	//				7.5 rounds down to 7 (because 7 is an odd number)
	//				7.4 rounds down to 7
	//				6.6 rounds up to 7
	//				6.5 rounds up to 7 (because 7 is an odd number)
	//				6.4 rounds down to 6
	//
	//		Randomly
	//
	//			Round Half Randomly. Uses a Random Number Generator
	//			to choose between rounding 0.5 up or down.
	//
	//			All numbers other than 0.5 round to the nearest as
	//			usual.
	//
	//		Floor
	//
	//			Yields the nearest integer down. Floor does not apply
	//			any	special treatment to 0.5.
	//
	//			Floor Function: The greatest integer that is less than
	//			or equal to x
	//
	//			Source:
	//				https://www.mathsisfun.com/sets/function-floor-ceiling.html
	//
	//			In mathematics and computer science, the floor function
	//			is the function that takes as input a real number x,
	//			and gives as output the greatest integer less than or
	//			equal to x,	denoted floor(x) or ⌊x⌋.
	//
	//			Source:
	//				https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
	//
	//			Examples of Floor
	//
	//				Number     Floor
	//				 2           2
	//				 2.4         2
	//				 2.9         2
	//				-2.5        -3
	//				-2.7        -3
	//				-2          -2
	//
	//		Ceiling
	//
	//			Yields the nearest integer up. Ceiling does not
	//			apply any special treatment to 0.5.
	//
	//			Ceiling Function: The least integer that is
	//			greater than or	equal to x.
	//			Source:
	//				https://www.mathsisfun.com/sets/function-floor-ceiling.html
	//
	//			The ceiling function maps x to the least integer
	//			greater than or equal to x, denoted ceil(x) or
	//			⌈x⌉.[1]
	//
	//			Source:
	//				https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
	//
	//				Examples of Ceiling
	//
	//					Number    Ceiling
	//					 2           2
	//					 2.4         3
	//					 2.9         3
	//					-2.5        -2
	//					-2.7        -2
	//					-2          -2
	//
	//		Truncate
	//
	//			Apply NO Rounding whatsoever. The Round From Digit
	//			is dropped or deleted. The Round To Digit is NEVER
	//			changed.
	//
	//			Examples of Truncate
	//
	//				Example-1
	//				Number: 23.14567
	//				Objective: Round to two decimal places to
	//				the right of the decimal point.
	//				Rounding Method: Truncate
	//				Round To Digit:   4
	//				Round From Digit: 5
	//				Rounded Number:   23.14 - The Round From Digit
	//				is dropped.
	//
	//				Example-2
	//				Number: -23.14567
	//				Objective: Round to two decimal places to
	//				the right of the decimal point.
	//				Rounding Method: Truncate
	//				Round To Digit:   4
	//				Round From Digit: 5
	//				Rounded Number:  -23.14 - The Round From Digit
	//				is dropped.

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
func (txtFieldFmtDtoFloat64Atom *textFieldFormatDtoFloat64Atom) empty(
	txtFieldFmtDtoFloat64 *TextFieldFormatDtoFloat64) {

	if txtFieldFmtDtoFloat64Atom.lock == nil {
		txtFieldFmtDtoFloat64Atom.lock = new(sync.Mutex)
	}

	txtFieldFmtDtoFloat64Atom.lock.Lock()

	defer txtFieldFmtDtoFloat64Atom.lock.Unlock()

	if txtFieldFmtDtoFloat64 == nil {

		return
	}

	txtFieldFmtDtoFloat64.LeftMarginStr = ""

	txtFieldFmtDtoFloat64.Float64Num = 0.0

	txtFieldFmtDtoFloat64.RoundingType =
		NumRoundType.None()

	txtFieldFmtDtoFloat64.NumOfFractionalDigits = 0

	txtFieldFmtDtoFloat64.FieldLength = 0

	txtFieldFmtDtoFloat64.FieldJustify = TxtJustify.None()

	txtFieldFmtDtoFloat64.RightMarginStr = ""

	return
}

// equal
//
// Compares two instances of TextFieldFormatDtoFloat64
// and returns a boolean value signaling whether the two
// instances are equivalent in all respects.
//
// If the two instances of TextFieldFormatDtoFloat64 are
// equal, this method returns 'true'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	txtFloat64FieldFmtDtoOne		*TextFieldFormatDtoFloat64
//
//		A pointer to an instance of
//		TextFieldFormatDtoFloat64.
//
//		The data values contained within this instance
//		will be compared to corresponding data values
//		contained within a second
//		TextFieldFormatDtoFloat64 instance
//		('txtFloat64FieldFmtDtoTwo') in order to
//		determine if they are equivalent.
//
//	txtFloat64FieldFmtDtoTwo		*TextFieldFormatDtoFloat64
//
//		A pointer to the second of two instances of
//		TextFieldFormatDtoFloat64. The data values
//		contained within this instance will be compared
//		to corresponding data values contained within the
//		first TextFieldFormatDtoFloat64 instance
//		('txtFloat64FieldFmtDtoOne') in order to
//		determine if they are equivalent.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		If all the data values within input parameters
//		'txtFloat64FieldFmtDtoOne' and
//		'txtFloat64FieldFmtDtoTwo' are found to be
//		equivalent in all respects, this return parameter
//		will be set to 'true'.
//
//		If the compared data values are NOT equivalent,
//		this method returns 'false'.
func (txtFieldFmtDtoFloat64Atom *textFieldFormatDtoFloat64Atom) equal(
	txtFloat64FieldFmtDtoOne *TextFieldFormatDtoFloat64,
	txtFloat64FieldFmtDtoTwo *TextFieldFormatDtoFloat64) bool {

	if txtFieldFmtDtoFloat64Atom.lock == nil {
		txtFieldFmtDtoFloat64Atom.lock = new(sync.Mutex)
	}

	txtFieldFmtDtoFloat64Atom.lock.Lock()

	defer txtFieldFmtDtoFloat64Atom.lock.Unlock()

	if txtFloat64FieldFmtDtoOne == nil ||
		txtFloat64FieldFmtDtoTwo == nil {

		return false
	}

	if txtFloat64FieldFmtDtoOne.LeftMarginStr !=
		txtFloat64FieldFmtDtoTwo.LeftMarginStr {

		return false
	}

	if txtFloat64FieldFmtDtoOne.RoundingType !=
		txtFloat64FieldFmtDtoTwo.RoundingType {

		return false
	}

	if txtFloat64FieldFmtDtoOne.NumOfFractionalDigits !=
		txtFloat64FieldFmtDtoTwo.NumOfFractionalDigits {

		return false
	}

	if txtFloat64FieldFmtDtoOne.FieldLength !=
		txtFloat64FieldFmtDtoTwo.FieldLength {

		return false
	}

	if txtFloat64FieldFmtDtoOne.FieldJustify !=
		txtFloat64FieldFmtDtoTwo.FieldJustify {

		return false
	}

	if txtFloat64FieldFmtDtoOne.RightMarginStr !=
		txtFloat64FieldFmtDtoTwo.RightMarginStr {

		return false
	}

	var float64NumStrOne, float64NumStrTwo string

	var float64NumberStrKernelOne,
		float64NumberStrKernelTwo NumberStrKernel

	var err error

	float64NumberStrKernelOne,
		err = new(NumberStrKernel).NewFromFloatValue(
		txtFloat64FieldFmtDtoOne.Float64Num,
		nil)

	if err != nil {
		return false
	}

	float64NumberStrKernelTwo,
		err = new(NumberStrKernel).NewFromFloatValue(
		txtFloat64FieldFmtDtoTwo.Float64Num,
		nil)

	if err != nil {
		return false
	}

	numberFieldSpec := NumStrNumberFieldSpec{
		fieldLength:        -1,
		fieldJustification: TxtJustify.Right(),
	}

	roundingSpecOne := NumStrRoundingSpec{
		roundingType:            txtFloat64FieldFmtDtoOne.RoundingType,
		roundToFractionalDigits: txtFloat64FieldFmtDtoOne.NumOfFractionalDigits,
	}

	float64NumStrOne,
		err = float64NumberStrKernelOne.FmtSignedNumStrUS(
		numberFieldSpec,
		roundingSpecOne,
		nil)

	if err != nil {
		return false
	}

	roundingSpecTwo := NumStrRoundingSpec{
		roundingType:            txtFloat64FieldFmtDtoTwo.RoundingType,
		roundToFractionalDigits: txtFloat64FieldFmtDtoTwo.NumOfFractionalDigits,
	}

	float64NumStrOne,
		err = float64NumberStrKernelTwo.FmtSignedNumStrUS(
		numberFieldSpec,
		roundingSpecTwo,
		nil)

	if err != nil {
		return false
	}

	if float64NumStrOne != float64NumStrTwo {
		return false
	}

	return true
}

// testValidityOfTxtFieldFmtDtoFloat64
//
// Receives a pointer to an instance of
// TextFieldFormatDtoFloat64 and performs a diagnostic
// analysis to determine if the data values contained in
// that instance are valid in all respects.
//
// If the input parameter 'txtFieldFmtDtoFloat64' is
// determined to be invalid, this method will return a
// boolean flag ('isValid') of 'false'. In addition, an
// instance of type error ('err') will be returned
// configured with an appropriate error message.
//
// If the input parameter 'txtFieldFmtDtoFloat64' is
// valid, this method will return a boolean flag
// ('isValid') of 'true' and the returned error type
// ('err') will be set to 'nil'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	txtFieldFmtDtoFloat64		*TextFieldFormatDtoFloat64
//
//		A pointer to an instance of
//		TextFieldFormatDtoFloat64.
//
//		The data values contained in this instance will
//		be reviewed and analyzed to determine if they
//		are valid in all respects.
//
//		None of the data values in this instance will be
//		changed or modified.
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string
//		which is included in all returned error
//		messages. Usually, it contains the name of the
//		calling method or methods listed as a function
//		chain.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref'
//		software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	isValid						bool
//
//		If all data elements contained within input
//		parameter 'txtFieldFmtDtoFloat64' are judged to
//		be valid, this returned boolean value will be set
//		to 'true'. If any data values are invalid, this
//		return parameter will be set to 'false'.
//
//	error
//
//		If this method completes successfully and all the
//		data values contained in input parameter
//		'txtFieldFmtDtoFloat64' are judged to be valid,
//		the returned error Type will be set equal to
//		'nil'.
//
//		If the data values contained in input parameter
//		'txtFieldFmtDtoFloat64' are invalid, the
//		returned 'error' will be non-nil and configured
//		with an appropriate error message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (txtFieldFmtDtoFloat64Atom *textFieldFormatDtoFloat64Atom) testValidityOfTxtFieldFmtDtoFloat64(
	txtFieldFmtDtoFloat64 *TextFieldFormatDtoFloat64,
	errPrefDto *ePref.ErrPrefixDto) (
	isValid bool,
	err error) {

	if txtFieldFmtDtoFloat64Atom.lock == nil {
		txtFieldFmtDtoFloat64Atom.lock = new(sync.Mutex)
	}

	txtFieldFmtDtoFloat64Atom.lock.Lock()

	defer txtFieldFmtDtoFloat64Atom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	isValid = false

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textFieldFormatDtoFloat64Atom."+
			"testValidityOfTxtFieldFmtDtoFloat64()",
		"")

	if err != nil {

		return isValid, err

	}

	if txtFieldFmtDtoFloat64 == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'txtFieldFmtDtoFloat64' is a nil pointer!\n",
			ePrefix.String())

		return isValid, err
	}

	if txtFieldFmtDtoFloat64.FieldLength < -1 {

		err = fmt.Errorf("%v\n"+
			"ERROR: TextFieldFormatDtoFloat64 parameter 'FieldLength' is INVALID!\n"+
			"txtFieldFmtDtoFloat64.FieldLength has a value less than minus one (-1)\n"+
			"txtFieldFmtDtoFloat64.FieldLength = %v\n",
			ePrefix.String(),
			txtFieldFmtDtoFloat64.FieldLength)

		return isValid, err
	}

	if txtFieldFmtDtoFloat64.FieldLength > 1000000 {

		err = fmt.Errorf("%v\n"+
			"ERROR: TextFieldFormatDtoFloat64 parameter 'FieldLength' is INVALID!\n"+
			"txtFieldFmtDtoFloat64.FieldLength has a value greater than one-million (1,000,000)\n"+
			"txtFieldFmtDtoFloat64.FieldLength = %v\n",
			ePrefix.String(),
			txtFieldFmtDtoFloat64.FieldLength)

		return isValid, err
	}

	if txtFieldFmtDtoFloat64.NumOfFractionalDigits < -1 {

		err = fmt.Errorf("%v\n"+
			"ERROR: TextFieldFormatDtoFloat64 parameter 'NumOfFractionalDigits' is INVALID!\n"+
			"txtFieldFmtDtoFloat64.NumOfFractionalDigits has a value less than minus one (-1)\n"+
			"txtFieldFmtDtoFloat64.NumOfFractionalDigits = %v\n",
			ePrefix.String(),
			txtFieldFmtDtoFloat64.NumOfFractionalDigits)

		return isValid, err
	}

	if !txtFieldFmtDtoFloat64.RoundingType.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"ERROR: TextFieldFormatDtoFloat64 parameter 'RoundingType' is INVALID!\n"+
			"txtFieldFmtDtoFloat64.RoundingType must be set to a valid value as follows:\n"+
			" NumRoundType.NoRounding()\n"+
			" NumRoundType.HalfUpWithNegNums()\n"+
			" NumRoundType.HalfDownWithNegNums()\n"+
			" NumRoundType.HalfAwayFromZero()\n"+
			" NumRoundType.HalfTowardsZero()\n"+
			" NumRoundType.HalfToEven()\n"+
			" NumRoundType.HalfToOdd()\n"+
			" NumRoundType.Randomly()\n"+
			" NumRoundType.Floor()\n"+
			" NumRoundType.Ceiling()\n"+
			" NumRoundType.Truncate()\n"+
			" txtFieldFmtDtoFloat64.RoundingType String Value = %v\n"+
			"txtFieldFmtDtoFloat64.RoundingType Integer Value = %v\n",
			ePrefix.String(),
			txtFieldFmtDtoFloat64.RoundingType.String(),
			txtFieldFmtDtoFloat64.RoundingType.XValueInt())

		return isValid, err
	}

	isValid = true

	return isValid, err
}

// textFieldFormatDtoBigFloatElectron - Provides helper
// methods for TextFieldFormatDtoBigFloat.
type textFieldFormatDtoFloat64Electron struct {
	lock *sync.Mutex
}

// getFloat64PureNumberStr
//
// Receives a pointer to an instance of
// TextFieldFormatDtoFloat64 and extracts the
// specifications necessary to format and return a
// floating point, pure number string.
//
// The floating point pure number string returned by
// this method will:
//
//  1. Consist entirely of numeric digit characters.
//
//  2. Separate integer and fractional digits with a
//     decimal point ('.').
//
//  3. Designate negative values with a leading minus
//     sign ('-').
//
//  4. NOT include integer separators such as commas
//     (',') to separate integer digits by thousands.
//
//     NOT THIS: 1,000,000
//     Pure Number String: 1000000
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	Pure number strings Do NOT include integer separators
//	(i.e. commas ',') to separate integer number strings
//	into thousands.
//
//					  NOT THIS: 1,000,000
//			Pure Number String: 1000000
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	txtFieldFmtDtoFloat64		*TextFieldFormatDtoFloat64
//
//		A pointer to an instance of
//		TextFieldFormatDtoFloat64.
//
//		This instance of TextFieldFormatDtoFloat64 will
//		be converted, formatted and returned as a
//		floating point pure number string.
//
//		If this instance of TextFieldFormatDtoFloat64
//		contains invalid data elements, an error will
//		be returned.
//
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string
//		which is included in all returned error
//		messages. Usually, it contains the name of the
//		calling method or methods listed as a function
//		chain.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref'
//		software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	string
//
//		If this method completes successfully, this
//		string parameter will return a floating point
//		pure number string representation of the
//		float64 value passed by input paramter,
//		'txtFieldFmtDtoFloat64'.
//
//		The returned floating point pure number string
//		will:
//
//		1.	Consist entirely of numeric digit characters.
//
//		2.	Separate integer and fractional digits with a
//			decimal point ('.').
//
//		3.	Designate negative values with a leading minus
//			sign ('-').
//
//		4.	NOT include integer separators such as commas
//			(',') to separate integer digits by thousands.
//
//						  NOT THIS: 1,000,000
//				Pure Number String: 1000000
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (txtFieldFmtDtoFloat64Electron *textFieldFormatDtoFloat64Electron) getFloat64PureNumberStr(
	txtFieldFmtDtoFloat64 *TextFieldFormatDtoFloat64,
	errPrefDto *ePref.ErrPrefixDto) (
	string,
	error) {

	if txtFieldFmtDtoFloat64Electron.lock == nil {
		txtFieldFmtDtoFloat64Electron.lock = new(sync.Mutex)
	}

	txtFieldFmtDtoFloat64Electron.lock.Lock()

	defer txtFieldFmtDtoFloat64Electron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textFieldFormatDtoFloat64Electron."+
			"getFloat64PureNumberStr()",
		"")

	if err != nil {

		return "", err

	}

	if txtFieldFmtDtoFloat64 == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'txtFieldFmtDtoFloat64' is a nil pointer!\n",
			ePrefix.String())

		return "", err
	}

	_,
		err = new(textFieldFormatDtoFloat64Atom).
		testValidityOfTxtFieldFmtDtoFloat64(
			txtFieldFmtDtoFloat64,
			ePrefix.XCpy(
				"txtFieldFmtDtoFloat64 Invalid"))

	if err != nil {

		return "", err

	}

	var float64NumStr string

	var float64NumberStrKernel NumberStrKernel

	float64NumberStrKernel,
		err = new(NumberStrKernel).NewFromFloatValue(
		txtFieldFmtDtoFloat64.Float64Num,
		ePrefix.XCpy(
			"txtFieldFmtDtoFloat64.Float64Num"))

	if err != nil {

		return "", err

	}

	roundingSpecOne := NumStrRoundingSpec{
		roundingType:            txtFieldFmtDtoFloat64.RoundingType,
		roundToFractionalDigits: txtFieldFmtDtoFloat64.NumOfFractionalDigits,
	}

	float64NumStr,
		err = float64NumberStrKernel.FmtSimpleSignedNumber(
		".",
		",",
		true,
		-1,
		TxtJustify.Right(),
		roundingSpecOne,
		ePrefix.XCpy(
			"float64NumStr<-float64NumberStrKernel"))

	return float64NumStr, err
}
