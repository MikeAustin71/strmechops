package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"strconv"
	"sync"
)

type MathIntHelper struct {
	lock *sync.Mutex
}

//	IntegerToAbsValueNumStr
//
//	Receives one of several types of integer values
//	and converts that value to a pure number string
//	containing the absolute value of the original integer
//	input value.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	signedIntValue	interface{}
//
//		Integer numeric values passed by means of this
//		empty interface MUST BE convertible to one of the
//		following types:
//
//			int8
//			int16
//			int32
//			int	(equivalent to int32)
//			int64
//			*big.Int
//			uint8
//			uint16
//			uint32
//			uint (equivalent to uint32)
//			uint64
//
//		If parameter 'signedIntValue' is NOT convertible
//		to one of the types listed above, an error will
//		be returned.
//
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	absValueNumberStr			string
//
//		The integer input parameter, 'signedIntValue',
//		will be converted to an absolute value and
//		returned as	a string of numeric digits.
//
//	numberStrStats				NumberStrStatsDto
//
//		This returned instance of NumberStrStatsDto
//		contains a description of the numeric value
//		passed as input parameter 'signedIntValue'.
//		This returned information includes the number
//		sign, number of significant digits and a zero
//		value indicator.
//
//		type NumberStrStatsDto struct {
//
//		NumOfIntegerDigits					uint64
//
//			The total number of integer digits.
//
//		NumOfSignificantIntegerDigits		uint64
//
//			The number of nonzero integer digits.
//
//		NumOfFractionalDigits				uint64
//
//			This value is always zero. No fractional
//			digits are ever returned.
//
//		NumOfSignificantFractionalDigits	uint64
//
//			This value is always zero. No fractional
//			digits are ever returned.
//
//		NumberValueType 					NumericValueType
//
//			This enumeration value specifies whether the
//			subject numeric value is classified either as
//			an integer or a floating point number.
//
//			For purposes of this method, the NumberValueType
//			is always set to: NumValType.Integer()
//
//
//		NumberSign							NumericSignValueType
//
//			An enumeration specifying the number sign
//			associated with the input parameter
//			'signedIntValue'.
//
//			Possible values are listed as follows:
//				NumSignVal.Negative()	= -1
//				NumSignVal.Zero()		=  0
//				NumSignVal.Positive()	=  1
//
//		IsZeroValue							bool
//
//			If 'false', the subject numeric value is
//			greater than or less than zero ('0').
//
//			If 'true', the Numeric Value is equal to
//			zero.
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
func (mathIntHelper *MathIntHelper) IntegerToAbsValueNumStr(
	numericValue interface{},
	errorPrefix interface{}) (
	absValueNumberStr string,
	numberSign NumericSignValueType,
	err error) {

	if mathIntHelper.lock == nil {
		mathIntHelper.lock = new(sync.Mutex)
	}

	mathIntHelper.lock.Lock()

	defer mathIntHelper.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"MathIntHelper."+
			"IntegerToAbsValueNumStr()",
		"")

	if err != nil {
		return absValueNumberStr, numberSign, err
	}

	var ok bool

	switch numericValue.(type) {

	case int8:

		var int8Num int8

		int8Num, ok = numericValue.(int8)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: int8 cast to 'int8Num' failed!\n",
				ePrefix.String())

			return absValueNumberStr, numberSign, err
		}

		if int8Num < 0 {

			numberSign = NumSignVal.Negative()

			int8Num = int8Num * -1

		} else if int8Num == 0 {

			numberSign = NumSignVal.Zero()

		} else {
			// MUST BE -
			//	int8Num  > 0

			numberSign = NumSignVal.Positive()

		}

		absValueNumberStr = fmt.Sprintf("%v",
			int8Num)

	case int16:

		var int16Num int16

		int16Num, ok = numericValue.(int16)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: int16 cast to 'int16Num' failed!\n",
				ePrefix.String())

			return absValueNumberStr, numberSign, err
		}

		if int16Num < 0 {

			numberSign = NumSignVal.Negative()

			int16Num = int16Num * -1

		} else if int16Num == 0 {

			numberSign = NumSignVal.Zero()

		} else {
			// MUST BE -
			//	int16Num  > 0

			numberSign = NumSignVal.Positive()

		}

		absValueNumberStr = fmt.Sprintf("%v",
			int16Num)

	case int:

		var intNum int

		intNum, ok = numericValue.(int)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: int cast to 'intNum' failed!\n",
				ePrefix.String())

			return absValueNumberStr, numberSign, err
		}

		if intNum < 0 {

			numberSign = NumSignVal.Negative()

			intNum = intNum * -1

		} else if intNum == 0 {

			numberSign = NumSignVal.Zero()

		} else {
			// MUST BE -
			//	intNum  > 0

			numberSign = NumSignVal.Positive()

		}

		absValueNumberStr = fmt.Sprintf("%v",
			intNum)

	case int32:

		var int32Num int32

		int32Num, ok = numericValue.(int32)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: int32 cast to 'int32Num' failed!\n",
				ePrefix.String())

			return absValueNumberStr, numberSign, err
		}

		if int32Num < 0 {

			numberSign = NumSignVal.Negative()

			int32Num = int32Num * -1

		} else if int32Num == 0 {

			numberSign = NumSignVal.Zero()

		} else {
			// MUST BE -
			//	int32Num  > 0

			numberSign = NumSignVal.Positive()

		}

		absValueNumberStr = fmt.Sprintf("%v",
			int32Num)

	case int64:

		var int64Num int64

		int64Num, ok = numericValue.(int64)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: int64 cast to 'int64Num' failed!\n",
				ePrefix.String())

			return absValueNumberStr, numberSign, err
		}

		if int64Num < 0 {

			numberSign = NumSignVal.Negative()

			int64Num = int64Num * -1

		} else if int64Num == 0 {

			numberSign = NumSignVal.Zero()

		} else {
			// MUST BE -
			//	int64Num  > 0

			numberSign = NumSignVal.Positive()

		}

		absValueNumberStr = fmt.Sprintf("%v",
			int64Num)

	case *big.Int:

		var bigIntNum *big.Int

		bigIntNum, ok = numericValue.(*big.Int)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: *big.Int cast to 'bigIntNum' failed!\n",
				ePrefix.String())

			return absValueNumberStr, numberSign, err
		}

		var bigZero *big.Int

		bigZero = big.NewInt(0)

		comparison := bigIntNum.Cmp(bigZero)

		if comparison < 0 {

			numberSign = NumSignVal.Negative()

			bigIntNum.Neg(bigIntNum)

		} else if comparison == 0 {

			numberSign = NumSignVal.Zero()

		} else {
			// MUST BE -
			//	int64Num  > 0

			numberSign = NumSignVal.Positive()

		}

		absValueNumberStr = fmt.Sprintf("%v",
			bigIntNum.Text(10))

	case uint8:

		var uint8Num uint8

		uint8Num, ok = numericValue.(uint8)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: uint8 cast to 'uint8Num' failed!\n",
				ePrefix.String())

			return absValueNumberStr, numberSign, err
		}

		if uint8Num == 0 {

			numberSign = NumSignVal.Zero()

		} else {
			// MUST BE -
			//	uint8Num  > 0

			numberSign = NumSignVal.Positive()

		}

		absValueNumberStr = fmt.Sprintf("%v",
			uint8Num)

	case uint16:

		var uint16Num uint16

		uint16Num, ok = numericValue.(uint16)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: uint16 cast to 'uint16Num' failed!\n",
				ePrefix.String())

			return absValueNumberStr, numberSign, err
		}

		if uint16Num == 0 {

			numberSign = NumSignVal.Zero()

		} else {
			// MUST BE -
			//	uint16Num  > 0

			numberSign = NumSignVal.Positive()

		}

		absValueNumberStr = fmt.Sprintf("%v",
			uint16Num)

	case uint:

		var uintNum uint

		uintNum, ok = numericValue.(uint)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: uint cast to 'uintNum' failed!\n",
				ePrefix.String())

			return absValueNumberStr, numberSign, err
		}

		if uintNum == 0 {

			numberSign = NumSignVal.Zero()

		} else {
			// MUST BE -
			//	uintNum  > 0

			numberSign = NumSignVal.Positive()

		}

		absValueNumberStr = fmt.Sprintf("%v",
			uintNum)

	case uint32:

		var uint32Num uint32

		uint32Num, ok = numericValue.(uint32)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: uint32 cast to 'uint32Num' failed!\n",
				ePrefix.String())

			return absValueNumberStr, numberSign, err
		}

		if uint32Num == 0 {

			numberSign = NumSignVal.Zero()

		} else {
			// MUST BE -
			//	uint32Num  > 0

			numberSign = NumSignVal.Positive()

		}

		absValueNumberStr = fmt.Sprintf("%v",
			uint32Num)

	case uint64:

		var uint64Num uint64

		uint64Num, ok = numericValue.(uint64)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: uint64 cast to 'uint64Num' failed!\n",
				ePrefix.String())

			return absValueNumberStr, numberSign, err
		}

		if uint64Num == 0 {

			numberSign = NumSignVal.Zero()

		} else {
			// MUST BE -
			//	uint64Num  > 0

			numberSign = NumSignVal.Positive()

		}

		absValueNumberStr = fmt.Sprintf("%v",
			uint64Num)

	case float32:

		var float32Num float32

		float32Num, ok = numericValue.(float32)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: float32 cast to 'float32Num' failed!\n",
				ePrefix.String())

			return absValueNumberStr, numberSign, err
		}

		var float64Num float64

		float64Num = float64(float32Num)

		numberStr := strconv.FormatFloat(
			float64Num, 'f', -1, 32)

	default:

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numericValue' is an invalid type!\n"+
			"'numericValue' is unsupported type '%v'\n",
			ePrefix.String(),
			fmt.Sprintf("%T", numericValue))

		return absValueNumberStr, numberSign, err

	}

	return absValueNumberStr, numberSign, err
}
