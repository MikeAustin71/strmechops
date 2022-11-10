package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"strconv"
	"sync"
)

// mathFloatHelperAtom
//
// Provides helper methods for type MathFloatHelper
type mathFloatHelperAtom struct {
	lock *sync.Mutex
}

//	floatNumToIntFracRunes
//
//	Receives one of several types of floating point
//	values and converts that value to an integer digit
//	rune array and a fractional digit rune array.
//
//	The integer and fractional digit rune arrays
//	represent and absolute value of the original floating
//	point number.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	floatingPointNumber interface{}
//
//		Numeric values passed by means of this empty
//		interface MUST BE convertible to one of the
//		following types:
//
//			float32
//			float64
//			*big.Float
//
//	intDigits					*RuneArrayDto
//
//		A pointer to an instance of RuneArrayDto. The
//		integer numeric digits extracted from
//		'floatingPointNumber' will be stored as text
//		characters in the rune array encapsulated by
//		this RuneArrayDto object.
//
//		The positive or negative number sign for the
//		extracted integer digits, can be determined by
//		examining the statistics returned by parameter
//		'numberStats'.
//
//	fracDigits					*RuneArrayDto
//
//		A pointer to an instance of RuneArrayDto. The
//		fractional numeric digits extracted from
//		'floatingPointNumber' will be stored as text
//		characters in the rune array encapsulated by
//		this RuneArrayDto object.
//
//		The positive or negative number sign for the
//		extracted integer digits, can be determined by
//		examining the statistics returned by parameter
//		'numberStats'.
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
//	numberStats					NumberStrStatsDto
//
//		This data transfer object will return critical
//		statistics on the numeric value represented
//		by the integer and fractional digits extracted
//		from 'floatingPointNumber' and stored in the
//		'intDigits' and 'fracDigits' RuneArrayDto
//		objects.
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
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (mathFloatHelperAtom *mathFloatHelperAtom) floatNumToIntFracRunes(
	floatingPointNumber interface{},
	intDigits *RuneArrayDto,
	fracDigits *RuneArrayDto,
	errPrefDto *ePref.ErrPrefixDto) (
	numberStats NumberStrStatsDto,
	err error) {

	if mathFloatHelperAtom.lock == nil {
		mathFloatHelperAtom.lock = new(sync.Mutex)
	}

	mathFloatHelperAtom.lock.Lock()

	defer mathFloatHelperAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	numberStats.NumberSign = NumSignVal.Zero()

	numberStats.IsZeroValue = true

	numberStats.NumberValueType = NumValType.Integer()

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrKernelQuark."+
			"equalizeNStrFracDigitsLengths()",
		"")

	if err != nil {

		return numberStats, err
	}

	if intDigits == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'intDigits' is a nil pointer!\n",
			ePrefix.String())

		return numberStats, err
	}

	if fracDigits == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'fracDigits' is a nil pointer!\n",
			ePrefix.String())

		return numberStats, err
	}

	new(runeArrayDtoAtom).empty(
		intDigits)

	intDigits.charSearchType =
		CharSearchType.LinearTargetStartingIndex()

	new(runeArrayDtoAtom).empty(
		fracDigits)

	fracDigits.charSearchType =
		CharSearchType.LinearTargetStartingIndex()

	var ok bool
	var float64Num float64
	var numberStr string

	switch floatingPointNumber.(type) {

	case float32:

		var float32Num float32

		float32Num, ok = floatingPointNumber.(float32)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: float32 cast to 'float32Num' failed!\n",
				ePrefix.String())

			return numberStats, err
		}

		float64Num = float64(float32Num)

	case float64:

		float64Num, ok = floatingPointNumber.(float64)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: float64 cast to 'float64Num' failed!\n",
				ePrefix.String())

			return numberStats, err
		}

	case *big.Float:

		var bigFloatNum *big.Float

		bigFloatNum, ok = floatingPointNumber.(*big.Float)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: *big.Float cast to 'bigFloatNum' failed!\n",
				ePrefix.String())

			return numberStats, err
		}

		numberStr = fmt.Sprintf("%v",
			bigFloatNum.Text('f', -1))

		goto standardPrep

	default:

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numericValue' is an invalid type!\n"+
			"'numericValue' is unsupported type '%T'\n",
			ePrefix.String(),
			floatingPointNumber)

		return numberStats, err

	}

	numberStr = strconv.FormatFloat(
		float64Num, 'f', -1, 64)

standardPrep:

	numberRunes := []rune(numberStr)

	lenNumberRunes := len(numberRunes)

	foundMinusSign := false

	foundRadixPoint := false

	isZero := true

	for i := 0; i < lenNumberRunes; i++ {

		if numberRunes[i] == '-' {

			foundMinusSign = true

			continue
		}

		if numberRunes[i] == '.' {

			foundRadixPoint = true

			continue
		}

		// intDigits *RuneArrayDto,
		// fracDigits *RuneArrayDto,

		if numberRunes[i] >= '0' &&
			numberRunes[i] <= '9' {

			if numberRunes[i] > '0' {
				isZero = false
			}

			if !foundRadixPoint {

				intDigits.CharsArray = append(
					intDigits.CharsArray, numberRunes[i])

			} else {

				fracDigits.CharsArray = append(
					fracDigits.CharsArray, numberRunes[i])
			}
		}

	}

	numberStats.NumOfIntegerDigits =
		uint64(len(intDigits.CharsArray))

	numberStats.NumOfSignificantIntegerDigits =
		numberStats.NumOfIntegerDigits -
			intDigits.GetCountLeadingZeros()

	numberStats.NumOfFractionalDigits =
		uint64(len(fracDigits.CharsArray))

	numberStats.NumOfSignificantFractionalDigits =
		numberStats.NumOfFractionalDigits -
			fracDigits.GetCountTrailingZeros()

	if numberStats.NumOfFractionalDigits > 0 {

		numberStats.NumberValueType =
			NumValType.FloatingPoint()

	} else if numberStats.NumOfIntegerDigits > 0 {

		numberStats.NumberValueType =
			NumValType.Integer()
	} else {

		numberStats.NumberValueType =
			NumValType.None()
	}

	numberStats.IsZeroValue = isZero

	if !numberStats.IsZeroValue {

		if foundMinusSign {

			numberStats.NumberSign =
				NumSignVal.Negative()

		} else {

			numberStats.NumberSign =
				NumSignVal.Positive()

		}

	} else {
		// MUST BE -
		// numberStats.IsZeroValue == true

		numberStats.NumberSign =
			NumSignVal.Zero()

	}

	return numberStats, err
}
