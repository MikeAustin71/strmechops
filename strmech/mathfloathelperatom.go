package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
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
func (mathFloatHelperAtom *mathFloatHelperAtom) floatNumToIntFracRunes(
	floatingPointNumber interface{},
	intDigits *RuneArrayDto,
	fracDigits *RuneArrayDto,
	numberValueType NumericValueType,
	numberSign *NumericSignValueType,
	isZeroValue *bool,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if mathFloatHelperAtom.lock == nil {
		mathFloatHelperAtom.lock = new(sync.Mutex)
	}

	mathFloatHelperAtom.lock.Lock()

	defer mathFloatHelperAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrKernelQuark."+
			"equalizeNStrFracDigitsLengths()",
		"")

	if err != nil {

		return err
	}

	if intDigits == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'intDigits' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if fracDigits == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'fracDigits' is a nil pointer!\n",
			ePrefix.String())

		return err
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

	switch floatingPointNumber.(type) {

	case float32:

		var float32Num float32

		float32Num, ok = floatingPointNumber.(float32)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: float32 cast to 'float32Num' failed!\n",
				ePrefix.String())

			return numberSign, err
		}

		float64Num = float64(float32Num)

	case float64:

		float64Num, ok = floatingPointNumber.(float64)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: float64 cast to 'float64Num' failed!\n",
				ePrefix.String())

			return numberSign, err
		}

	default:

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numericValue' is an invalid type!\n"+
			"'numericValue' is unsupported type '%T'\n",
			ePrefix.String(),
			floatingPointNumber)

		return numberSign, err

	}

	numberStr := strconv.FormatFloat(
		float64Num, 'f', -1, 64)

	numberRunes := []rune(numberStr)

	lenNumberRunes := len(numberRunes)

	foundMinusSign := false

	foundRadixPoint := false

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

			if !foundRadixPoint {

				intRunes = append(
					intRunes, numberRunes[i])
			} else {

				fracRunes = append(
					fracRunes, numberRunes[i])
			}
		}

	}

}
