package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// mathIntHelperMechanics
//
// Provides helper methods for type MathIntHelper
type mathIntHelperMechanics struct {
	lock *sync.Mutex
}

// intToAbsoluteValueStr
//
// Receives one of several types of integer values
// and converts that value to a pure number string
// containing the absolute value of the original integer
// input value ('intNumericValue').
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	intNumericValue				interface{}
//
//		Integer numeric values passed by means of this
//		empty interface MUST BE convertible to one of the
//		following types:
//
//			int8
//			int16
//			int32
//			int	(currently equivalent to int32)
//			int64
//			uint8
//			uint16
//			uint32
//			uint (currently equivalent to uint32)
//			uint64
//			*big.Int
//
//		If parameter 'intNumericValue' is NOT convertible
//		to one of the types listed above, an error will
//		be returned.
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
//	absValueNumberStr			string
//
//		A pure number string consisting entirely of
//		numeric digit text characters. These numeric
//		digits represent the absolute value of the
//		integer number passed through input paramter
//		intNumericValue.
//
//		The absolute value of a number will always be
//		a positive numeric value ('+').
//
//		The positive or negative number sign for the
//		returned 'absValueStr' parameter can be
//		determined by examining the statistics returned
//		by parameter 'numberStrStats'
//		(numberStrStats.NumberSign).
//
//	numberStrStats				NumberStrStatsDto
//
//		This data transfer object will return critical
//		statistics on the numeric value represented
//		by the returned absolute value number string,
//		'absValueNumberStr'.
//
//		Since this is an integer value, statistics on
//		fractional digits are not relevant.
//
//		type NumberStrStatsDto struct {
//
//		NumOfIntegerDigits					uint64
//
//			The total number of integer digits to the
//			left of the radix point or, decimal point, in
//			the subject numeric value.
//
//		NumOfSignificantIntegerDigits		uint64
//
//			The number of nonzero integer digits to the
//			left of the radix point or, decimal point, in
//			the subject numeric value.
//
//		NumOfFractionalDigits				uint64
//
//			The total number of fractional digits to the
//			right of the radix point or, decimal point,
//			in the subject numeric value.
//
//		NumOfSignificantFractionalDigits	uint64
//
//			The number of nonzero fractional digits to
//			the right of the radix point or, decimal
//			point, in the subject numeric value.
//
//		NumberValueType 					NumericValueType
//
//			This enumeration value specifies whether the
//			subject numeric value is classified either as
//			an integer or a floating point number.
//
//			Possible enumeration values are listed as
//			follows:
//				NumValType.None()
//				NumValType.FloatingPoint()
//				NumValType.Integer()
//
//		NumberSign							NumericSignValueType
//
//			An enumeration specifying the number sign
//			associated with the numeric value. Possible
//			values are listed as follows:
//				NumSignVal.None()		= Invalid Value
//				NumSignVal.Negative()	= -1
//				NumSignVal.Zero()		=  0
//				NumSignVal.Positive()	=  1
//
//		IsZeroValue							bool
//
//			If 'true', the subject numeric value is equal
//			to zero ('0').
//
//			If 'false', the subject numeric value is
//			greater than or less than zero ('0').
//		}
//
//	err							error
//
//		If this method completes successfully, this
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (mathIntHelpMech *mathIntHelperMechanics) intToAbsoluteValueStr(
	intNumericValue interface{},
	errPrefDto *ePref.ErrPrefixDto) (
	absValueNumberStr string,
	numberStrStats NumberStrStatsDto,
	err error) {

	if mathIntHelpMech.lock == nil {
		mathIntHelpMech.lock = new(sync.Mutex)
	}

	mathIntHelpMech.lock.Lock()

	defer mathIntHelpMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"mathIntHelperMechanics."+
			"intToAbsoluteValueStr()",
		"")

	if err != nil {

		return absValueNumberStr,
			numberStrStats,
			err

	}

	var signedPureNumStr string

	signedPureNumStr,
		err = new(mathIntHelperAtom).
		intNumToSignedPureNumStr(
			intNumericValue,
			ePrefix.XCpy(
				"intNumericValue"))

	if err != nil {

		return absValueNumberStr,
			numberStrStats,
			err

	}

	var intDigits, fracDigits, decSeparatorChars RuneArrayDto

	decSeparatorChars,
		err = new(RuneArrayDto).NewString(
		".",
		CharSearchType.LinearTargetStartingIndex(),
		ePrefix)

	numberStrStats,
		err = new(numStrMathQuark).
		pureNumStrToRunes(
			signedPureNumStr,
			&intDigits,
			&fracDigits,
			&decSeparatorChars,
			true,
			ePrefix)

	if err != nil {

		return absValueNumberStr,
			numberStrStats,
			err

	}

	if len(fracDigits.CharsArray) > 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Integer Value returned fractional digits!\n"+
			"fracDigits.CharsArray = '%v'\n",
			ePrefix.String(),
			string(fracDigits.CharsArray))

		return absValueNumberStr,
			numberStrStats,
			err
	}

	absValueNumberStr =
		string(intDigits.CharsArray)

	return absValueNumberStr,
		numberStrStats,
		err
}
