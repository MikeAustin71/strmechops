package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"strconv"
	"sync"
)

// MathHelper
//
// Provides math utility methods
type MathHelper struct {
	lock *sync.Mutex
}

// NumericValueToPureNumStr
//
// Receives a numeric value as an empty interface and
// converts that value to a signed pure number string.
// This signed pure number string facilitates further
// conversion to integer and floating point number
// types.
//
// Packages and functions in the Go programming
// language, as well as other programming languages,
// typically use this type number string as an input
// parameter.
//
// A signed pure number string is defined as follows:
//
//  1. A number string consisting entirely of numeric
//     digit characters with the following exceptions.
//
//  2. A number string which separates integer and
//     fractional digits with a decimal point ('.').
//
//  3. A number string which designates negative values
//     with a leading minus sign ('-'). Positive numeric
//     values will have no leading number sign.
//
//  4. A number string which DOES NOT include integer
//     separators such as commas (',') to separate
//     integer digits by thousands.
//
//     NOT THIS: 1,000,000
//     Signed Pure Number String: 1000000
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//		numericValue				interface{}
//
//			An empty interface containing the numeric value
//			which will be converted and returned as signed
//			pure number string.
//
//			An error will be returned if the concrete type
//			passed through this parameter does not match one
//			of the supported types below.
//
//			Supported Numeric Value Types:
//	     		float32, float64, big.Float
//				*float32, *float64, *big.Float
//				*BigFloatDto, BigFloatDto
//				int8, int16, int, int32, int64, big.Int
//				*int8, *int16, *int, *int32, *int64, *big.Int
//				uint8, uint16, uint, uint32, uint64
//				*uint8, *uint16, *uint, *uint32, *uint64
//				*TextFieldFormatDtoFloat64, TextFieldFormatDtoFloat64
//				*TextFieldFormatDtoBigFloat, TextFieldFormatDtoBigFloat
//				*NumberStrKernel, NumberStrKernel
//
//		errorPrefix					interface{}
//
//			This object encapsulates error prefix text which
//			is included in all returned error messages.
//			Usually, it contains the name of the calling
//			method or methods listed as a method or function
//			chain of execution.
//
//			If no error prefix information is needed, set this
//			parameter to 'nil'.
//
//			This empty interface must be convertible to one of
//			the following types:
//
//			1.	nil
//					A nil value is valid and generates an
//					empty collection of error prefix and
//					error context information.
//
//			2.	string
//					A string containing error prefix
//					information.
//
//			3.	[]string
//					A one-dimensional slice of strings
//					containing error prefix information.
//
//			4.	[][2]string
//					A two-dimensional slice of strings
//			   		containing error prefix and error
//			   		context information.
//
//			5.	ErrPrefixDto
//					An instance of ErrPrefixDto.
//					Information from this object will
//					be copied for use in error and
//					informational messages.
//
//			6.	*ErrPrefixDto
//					A pointer to an instance of
//					ErrPrefixDto. Information from
//					this object will be copied for use
//					in error and informational messages.
//
//			7.	IBasicErrorPrefix
//					An interface to a method
//					generating a two-dimensional slice
//					of strings containing error prefix
//					and error context information.
//
//			If parameter 'errorPrefix' is NOT convertible
//			to one of the valid types listed above, it will
//			be considered invalid and trigger the return of
//			an error.
//
//			Types ErrPrefixDto and IBasicErrorPrefix are
//			included in the 'errpref' software package:
//				"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//		signedPureNumStr			string
//
//			If this method completes successfully, a signed
//			pure number string representing the numeric value
//			passed as input	parameter 'numericValue' will be
//			returned.
//
//			A signed pure number string is defined as
//			follows:
//
//	 	1.	A number string consisting entirely of
//			numeric digit characters with the following
//			exceptions.
//
//	 	2.	A number string which separates integer and
//	 	  	fractional digits with a decimal point ('.').
//
//	 	3.	A number string which designates negative
//			values with a leading minus sign ('-').
//			Positive numeric values will have no leading
//			number sign.
//
//	 	4.	A number string which DOES NOT include
//			integer separators such as commas (',') to
//			separate integer digits by thousands.
//
//	   						 NOT THIS: 1,000,000
//	   		Signed Pure Number String: 1000000
//
//		err							error
//
//			If this method completes successfully, the
//			returned error Type is set equal to 'nil'.
//
//			If errors are encountered during processing, the
//			returned error Type will encapsulate an error
//			message. This returned error message will
//			incorporate the method chain and text passed by
//			input parameter, 'errorPrefix'. The 'errorPrefix'
//			text will be attached to the beginning of the
//			error message.
func (mathHelper *MathHelper) NumericValueToPureNumStr(
	numericValue interface{},
	errorPrefix interface{}) (
	signedPureNumStr string,
	err error) {

	if mathHelper.lock == nil {
		mathHelper.lock = new(sync.Mutex)
	}

	mathHelper.lock.Lock()

	defer mathHelper.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"MathHelper."+
			"NumericValueToPureNumStr()",
		"")

	if err != nil {

		return signedPureNumStr, err
	}

	var ok bool

	if numericValue == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numericValue' is a nil pointer!\n",
			ePrefix.String())

		return signedPureNumStr, err
	}

	var int64Num int64

	var uint64Num uint64

	switch numericValue.(type) {

	case float32:

		var float32Num float32

		float32Num, ok = numericValue.(float32)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: float32 cast to 'float32Num' failed!\n",
				ePrefix.String())

			return signedPureNumStr, err
		}

		signedPureNumStr = strconv.FormatFloat(
			float64(float32Num),
			'f',
			-1,
			32)

		return signedPureNumStr, err

	case *float32:

		var ptrFloat32 *float32

		ptrFloat32, ok = numericValue.(*float32)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: *float32 cast to 'ptrFloat32' failed!\n",
				ePrefix.String())

			return signedPureNumStr, err
		}

		if ptrFloat32 == nil {

			err = fmt.Errorf("%v\n"+
				"ERROR: *float32 cast to 'ptrFloat32' failed!\n"+
				"ptrFloat32 is a nil pointer.",
				ePrefix.String())

			return signedPureNumStr, err

		}

		signedPureNumStr = strconv.FormatFloat(
			float64(*ptrFloat32),
			'f',
			-1,
			32)

		return signedPureNumStr, err

	case float64:

		var float64Num float64

		float64Num, ok = numericValue.(float64)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: float64 cast to 'float64Num' failed!\n",
				ePrefix.String())

			return signedPureNumStr, err
		}

		signedPureNumStr = strconv.FormatFloat(
			float64Num,
			'f',
			-1,
			64)

		return signedPureNumStr, err

	case *float64:

		var ptrFloat64 *float64

		ptrFloat64, ok = numericValue.(*float64)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: *float64 cast to 'ptrFloat64' failed!\n",
				ePrefix.String())

			return signedPureNumStr, err
		}

		if ptrFloat64 == nil {

			err = fmt.Errorf("%v\n"+
				"ERROR:  *float64 cast to 'ptrFloat64' failed!\n"+
				"ptrFloat64 is a nil pointer.",
				ePrefix.String())

			return signedPureNumStr, err

		}

		signedPureNumStr = strconv.FormatFloat(
			*ptrFloat64,
			'f',
			-1,
			64)

		return signedPureNumStr, err

	case *BigFloatDto:

		var ptrBigFloatDto *BigFloatDto

		ptrBigFloatDto, ok = numericValue.(*BigFloatDto)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: *BigFloatDto cast to 'ptrBigFloatDto' failed!\n",
				ePrefix.String())

			return signedPureNumStr, err
		}

		if ptrBigFloatDto == nil {

			err = fmt.Errorf("%v\n"+
				"ERROR:  *BigFloatDto cast to 'ptrBigFloatDto' failed!\n"+
				"ptrBigFloatDto is a nil pointer.",
				ePrefix.String())

			return signedPureNumStr, err

		}

		signedPureNumStr =
			ptrBigFloatDto.Value.Text('f', -1)

		return signedPureNumStr, err

	case BigFloatDto:

		var bigFloatDto BigFloatDto

		bigFloatDto, ok = numericValue.(BigFloatDto)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: BigFloatDto cast to 'bigFloatDto' failed!\n",
				ePrefix.String())

			return signedPureNumStr, err
		}

		signedPureNumStr =
			bigFloatDto.Value.Text('f', -1)

		return signedPureNumStr, err

	case *big.Float:

		var ptrBigFloatNum *big.Float

		ptrBigFloatNum, ok = numericValue.(*big.Float)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: *big.Float cast to 'ptrBigFloatNum' failed!\n",
				ePrefix.String())

			return signedPureNumStr, err
		}

		if ptrBigFloatNum == nil {

			err = fmt.Errorf("%v\n"+
				"ERROR:  *big.Float cast to 'ptrBigFloatNum' failed!\n"+
				"ptrBigFloatNum is a nil pointer.",
				ePrefix.String())

			return signedPureNumStr, err

		}

		signedPureNumStr =
			ptrBigFloatNum.Text('f', -1)

		return signedPureNumStr, err

	case big.Float:

		var bigFloatNum big.Float

		bigFloatNum, ok = numericValue.(big.Float)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: big.Float cast to 'bigFloatNum' failed!\n",
				ePrefix.String())

			return signedPureNumStr, err
		}

		signedPureNumStr =
			bigFloatNum.Text('f', -1)

		return signedPureNumStr, err

	case int8:

		var int8Num int8

		int8Num, ok = numericValue.(int8)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: int8 cast to 'int8Num' failed!\n",
				ePrefix.String())

			return signedPureNumStr, err
		}

		int64Num = int64(int8Num)

		goto conversionInteger

	case *int8:

		var ptrInt8Num *int8

		ptrInt8Num, ok = numericValue.(*int8)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: *int8 cast to 'ptrInt8Num' failed!\n",
				ePrefix.String())

			return signedPureNumStr, err
		}

		if ptrInt8Num == nil {

			err = fmt.Errorf("%v\n"+
				"ERROR:  *int8 cast to 'ptrInt8Num' failed!\n"+
				"ptrInt8Num is a nil pointer.",
				ePrefix.String())

			return signedPureNumStr, err

		}

		int64Num = int64(*ptrInt8Num)

		goto conversionInteger

	case int16:

		var int16Num int16

		int16Num, ok = numericValue.(int16)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: int16 cast to 'int16Num' failed!\n",
				ePrefix.String())

			return signedPureNumStr, err
		}

		int64Num = int64(int16Num)

		goto conversionInteger

	case *int16:

		var ptrInt16Num *int16

		ptrInt16Num, ok = numericValue.(*int16)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: *int16 cast to 'ptrInt16Num' failed!\n",
				ePrefix.String())

			return signedPureNumStr, err
		}

		if ptrInt16Num == nil {

			err = fmt.Errorf("%v\n"+
				"ERROR:  *int16 cast to 'ptrInt16Num' failed!\n"+
				"ptrInt16Num is a nil pointer.",
				ePrefix.String())

			return signedPureNumStr, err

		}

		int64Num = int64(*ptrInt16Num)

		goto conversionInteger

	case int:

		var intNum int

		intNum, ok = numericValue.(int)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: int cast to 'intNum' failed!\n",
				ePrefix.String())

			return signedPureNumStr, err
		}

		int64Num = int64(intNum)

		goto conversionInteger

	case *int:

		var ptrIntNum *int

		ptrIntNum, ok = numericValue.(*int)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: *int cast to 'ptrIntNum' failed!\n",
				ePrefix.String())

			return signedPureNumStr, err
		}

		if ptrIntNum == nil {

			err = fmt.Errorf("%v\n"+
				"ERROR:  *int cast to 'ptrIntNum' failed!\n"+
				"ptrIntNum is a nil pointer.",
				ePrefix.String())

			return signedPureNumStr, err

		}

		int64Num = int64(*ptrIntNum)

		goto conversionInteger

	case int32:

		var int32Num int32

		int32Num, ok = numericValue.(int32)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: int32 cast to 'int32Num' failed!\n",
				ePrefix.String())

			return signedPureNumStr, err
		}

		int64Num = int64(int32Num)

		goto conversionInteger

	case *int32:

		var ptrInt32Num *int32

		ptrInt32Num, ok = numericValue.(*int32)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: *int32 cast to 'ptrInt32Num' failed!\n",
				ePrefix.String())

			return signedPureNumStr, err
		}

		if ptrInt32Num == nil {

			err = fmt.Errorf("%v\n"+
				"ERROR:  *int32 cast to 'ptrInt32Num' failed!\n"+
				"ptrInt32Num is a nil pointer.",
				ePrefix.String())

			return signedPureNumStr, err

		}

		int64Num = int64(*ptrInt32Num)

		goto conversionInteger

	case int64:

		int64Num, ok = numericValue.(int64)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: int64 cast to 'int64Num' failed!\n",
				ePrefix.String())

			return signedPureNumStr, err
		}

		goto conversionInteger

	case *int64:

		var ptrInt64Num *int64

		ptrInt64Num, ok = numericValue.(*int64)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: *int64 cast to 'ptrInt64Num' failed!\n",
				ePrefix.String())

			return signedPureNumStr, err
		}

		if ptrInt64Num == nil {

			err = fmt.Errorf("%v\n"+
				"ERROR:  *int64 cast to 'ptrInt64Num' failed!\n"+
				"ptrInt64Num is a nil pointer.",
				ePrefix.String())

			return signedPureNumStr, err

		}

		int64Num = *ptrInt64Num

		goto conversionInteger

	case uint8:

		var uint8Num uint8

		uint8Num, ok = numericValue.(uint8)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: uint8 cast to 'uint8Num' failed!\n",
				ePrefix.String())

			return signedPureNumStr, err
		}

		uint64Num = uint64(uint8Num)

		goto conversionUnsignedInteger

	case *uint8:

		var ptrUint8Num *uint8

		ptrUint8Num, ok = numericValue.(*uint8)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: *uint8 cast to 'ptrUint8Num' failed!\n",
				ePrefix.String())

			return signedPureNumStr, err
		}

		if ptrUint8Num == nil {

			err = fmt.Errorf("%v\n"+
				"ERROR:  *uint8 cast to 'ptrUint8Num' failed!\n"+
				"ptrUint8Num is a nil pointer.",
				ePrefix.String())

			return signedPureNumStr, err

		}

		uint64Num = uint64(*ptrUint8Num)

		goto conversionUnsignedInteger

	case uint16:

		var uint16Num uint16

		uint16Num, ok = numericValue.(uint16)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: uint16 cast to 'uint16Num' failed!\n",
				ePrefix.String())

			return signedPureNumStr, err
		}

		uint64Num = uint64(uint16Num)

		goto conversionUnsignedInteger

	case *uint16:

		var ptrUint16Num *uint16

		ptrUint16Num, ok = numericValue.(*uint16)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: *uint16 cast to 'ptrUint16Num' failed!\n",
				ePrefix.String())

			return signedPureNumStr, err
		}

		if ptrUint16Num == nil {

			err = fmt.Errorf("%v\n"+
				"ERROR:  *uint16 cast to 'ptrUint16Num' failed!\n"+
				"ptrUint16Num is a nil pointer.",
				ePrefix.String())

			return signedPureNumStr, err

		}

		uint64Num = uint64(*ptrUint16Num)

		goto conversionUnsignedInteger

	case uint:

		var uintNum uint

		uintNum, ok = numericValue.(uint)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: uint cast to 'uintNum' failed!\n",
				ePrefix.String())

			return signedPureNumStr, err
		}

		uint64Num = uint64(uintNum)

		goto conversionUnsignedInteger

	case *uint:

		var ptrUintNum *uint

		ptrUintNum, ok = numericValue.(*uint)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: *uint cast to 'ptrUintNum' failed!\n",
				ePrefix.String())

			return signedPureNumStr, err
		}

		if ptrUintNum == nil {

			err = fmt.Errorf("%v\n"+
				"ERROR:  *uint cast to 'ptrUintNum' failed!\n"+
				"ptrUintNum is a nil pointer.",
				ePrefix.String())

			return signedPureNumStr, err

		}

		uint64Num = uint64(*ptrUintNum)

		goto conversionUnsignedInteger

	case uint32:

		var uint32Num uint32

		uint32Num, ok = numericValue.(uint32)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: uint32 cast to 'uint32Num' failed!\n",
				ePrefix.String())

			return signedPureNumStr, err
		}

		uint64Num = uint64(uint32Num)

		goto conversionUnsignedInteger

	case *uint32:

		var ptrUint32Num *uint32

		ptrUint32Num, ok = numericValue.(*uint32)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: *uint32 cast to 'ptrUint32Num' failed!\n",
				ePrefix.String())

			return signedPureNumStr, err
		}

		if ptrUint32Num == nil {

			err = fmt.Errorf("%v\n"+
				"ERROR:  *uint32 cast to 'ptrUint32Num' failed!\n"+
				"ptrUint32Num is a nil pointer.",
				ePrefix.String())

			return signedPureNumStr, err

		}

		uint64Num = uint64(*ptrUint32Num)

		goto conversionUnsignedInteger

	case uint64:

		uint64Num, ok = numericValue.(uint64)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: uint64 cast to 'uint64Num' failed!\n",
				ePrefix.String())

			return signedPureNumStr, err
		}

		goto conversionUnsignedInteger

	case *uint64:

		var ptrUint32Num *uint64

		ptrUint32Num, ok = numericValue.(*uint64)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: *uint64 cast to 'ptrUint32Num' failed!\n",
				ePrefix.String())

			return signedPureNumStr, err
		}

		if ptrUint32Num == nil {

			err = fmt.Errorf("%v\n"+
				"ERROR:  *uint64 cast to 'ptrUint32Num' failed!\n"+
				"ptrUint32Num is a nil pointer.",
				ePrefix.String())

			return signedPureNumStr, err

		}

		uint64Num = *ptrUint32Num

		goto conversionUnsignedInteger

	case *big.Int:

		var ptrBigIntNum *big.Int

		ptrBigIntNum, ok = numericValue.(*big.Int)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: *big.Int cast to 'ptrBigIntNum' failed!\n",
				ePrefix.String())

			return signedPureNumStr, err
		}

		signedPureNumStr = ptrBigIntNum.Text(10)

		return signedPureNumStr, err

	case big.Int:

		var bigIntNum big.Int

		bigIntNum, ok = numericValue.(big.Int)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: big.Int cast to 'bigIntNum' failed!\n",
				ePrefix.String())

			return signedPureNumStr, err
		}

		signedPureNumStr = bigIntNum.Text(10)

		return signedPureNumStr, err

	case *TextFieldFormatDtoFloat64:

		var ptrTxtFieldFmtDtoF64 *TextFieldFormatDtoFloat64

		ptrTxtFieldFmtDtoF64, ok =
			numericValue.(*TextFieldFormatDtoFloat64)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: *TextFieldFormatDtoFloat64 cast to 'ptrTxtFieldFmtDtoF64' failed!\n",
				ePrefix.String())

			return signedPureNumStr, err
		}

		signedPureNumStr,
			err = ptrTxtFieldFmtDtoF64.FmtNumStrNative(
			ePrefix.XCpy(
				"signedPureNumStr<-" +
					"ptrTxtFieldFmtDtoF64"))

		return signedPureNumStr, err

	case TextFieldFormatDtoFloat64:

		var txtFieldFmtDtoF64 TextFieldFormatDtoFloat64

		txtFieldFmtDtoF64, ok =
			numericValue.(TextFieldFormatDtoFloat64)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: TextFieldFormatDtoFloat64 cast to 'txtFieldFmtDtoF64' failed!\n",
				ePrefix.String())

			return signedPureNumStr, err
		}

		signedPureNumStr,
			err = txtFieldFmtDtoF64.FmtNumStrNative(
			ePrefix.XCpy(
				"signedPureNumStr<-" +
					"txtFieldFmtDtoF64"))

		return signedPureNumStr, err

	case *TextFieldFormatDtoBigFloat:

		var ptrTxtFieldFmtDtoBigFloat *TextFieldFormatDtoBigFloat

		ptrTxtFieldFmtDtoBigFloat, ok =
			numericValue.(*TextFieldFormatDtoBigFloat)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: *TextFieldFormatDtoBigFloat cast to 'ptrTxtFieldFmtDtoBigFloat' failed!\n",
				ePrefix.String())

			return signedPureNumStr, err
		}

		signedPureNumStr,
			err = ptrTxtFieldFmtDtoBigFloat.FmtNumStrNative(
			ePrefix.XCpy(
				"signedPureNumStr<-" +
					"ptrTxtFieldFmtDtoBigFloat"))

		return signedPureNumStr, err

	case TextFieldFormatDtoBigFloat:

		var txtFieldFmtDtoBigFloat TextFieldFormatDtoBigFloat

		txtFieldFmtDtoBigFloat, ok =
			numericValue.(TextFieldFormatDtoBigFloat)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: TextFieldFormatDtoBigFloat cast to 'txtFieldFmtDtoBigFloat' failed!\n",
				ePrefix.String())

			return signedPureNumStr, err
		}

		signedPureNumStr,
			err = txtFieldFmtDtoBigFloat.FmtNumStrNative(
			ePrefix.XCpy(
				"signedPureNumStr<-" +
					"txtFieldFmtDtoBigFloat"))

		return signedPureNumStr, err

	case *NumberStrKernel:

		var ptrNumStrKernel *NumberStrKernel

		ptrNumStrKernel, ok =
			numericValue.(*NumberStrKernel)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: *NumberStrKernel cast to 'ptrNumStrKernel' failed!\n",
				ePrefix.String())

			return signedPureNumStr, err
		}

		signedPureNumStr,
			err = ptrNumStrKernel.FmtNumStrNative(
			ePrefix.XCpy(
				"signedPureNumStr<-" +
					"ptrNumStrKernel"))

		return signedPureNumStr, err

	case NumberStrKernel:

		var numStrKernel NumberStrKernel

		numStrKernel, ok =
			numericValue.(NumberStrKernel)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: NumberStrKernel cast to 'numStrKernel' failed!\n",
				ePrefix.String())

			return signedPureNumStr, err
		}

		signedPureNumStr,
			err = numStrKernel.FmtNumStrNative(
			ePrefix.XCpy(
				"signedPureNumStr<-" +
					"numStrKernel"))

		return signedPureNumStr, err

	default:

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numericValue' is an invalid type!\n"+
			"'numericValue' is unsupported type '%v'\n",
			ePrefix.String(),
			fmt.Sprintf("%T", numericValue))

		return signedPureNumStr, err

	}

conversionUnsignedInteger:

	signedPureNumStr = strconv.FormatUint(uint64Num, 10)

	return signedPureNumStr, err

conversionInteger:

	signedPureNumStr = strconv.FormatInt(int64Num, 10)

	return signedPureNumStr, err

}
