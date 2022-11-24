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

//	floatNumToSignedPureNumStr
//
//	Receives one of several types of floating point
//	values and converts that value to a signed pure
//	number string containing the signed numeric value
//	extracted from the original floating point input
//	parameter, 'floatingPointNumber'.
//
//	A signed pure number string consists entirely of
//	numeric digit text characters. In the case of
//	negative numeric values, the negative numeric value
//	is prefixed with a leading minus sign ('-').
//
//	Integer and fractional numeric digits contained in
//	the signed pure number string will be separated by a
//	period character ('.') known as the decimal point.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	floatingPointNumber 		interface{}
//
//		Numeric values passed by means of this empty
//		interface MUST BE convertible to one of the
//		following types:
//
//			float32
//			float64
//			*big.Float
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
//	signedPureNumStr			string
//
//		If this method completes successfully, this
//		parameter will return a signed pure number string
//		representing the floating point numeric value
//		extracted from input parameter,
//		'floatingPointNumber'.
//
//		A signed pure number string consists entirely of
//		numeric digit text characters. In the case of
//		negative numeric values, the negative numeric
//		value is prefixed with a leading minus sign
//		('-').
//
//		Integer and fractional numeric digits contained
//		in this signed pure number string will be
//		separated by radix point, or decimal separator,
//		specified by input paramter, 'decSeparatorChars'.
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
func (mathFloatHelpAtom *mathFloatHelperAtom) floatNumToSignedPureNumStr(
	floatingPointNumber interface{},
	errPrefDto *ePref.ErrPrefixDto) (
	signedPureNumStr string,
	err error) {

	if mathFloatHelpAtom.lock == nil {
		mathFloatHelpAtom.lock = new(sync.Mutex)
	}

	mathFloatHelpAtom.lock.Lock()

	defer mathFloatHelpAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"mathFloatHelperAtom."+
			"floatNumToSignedPureNumStr()",
		"")

	if err != nil {

		return signedPureNumStr, err
	}

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

			return signedPureNumStr, err
		}

		float64Num = float64(float32Num)

	case float64:

		float64Num, ok = floatingPointNumber.(float64)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: float64 cast to 'float64Num' failed!\n",
				ePrefix.String())

			return signedPureNumStr, err
		}

	case *big.Float:

		var bigFloatNum *big.Float

		bigFloatNum, ok = floatingPointNumber.(*big.Float)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: *big.Float cast to 'bigFloatNum' failed!\n",
				ePrefix.String())

			return signedPureNumStr, err
		}

		// bigFloatNum.SetPrec(bigFloatNum.MinPrec())

		signedPureNumStr = fmt.Sprintf("%v",
			bigFloatNum.Text('f', -1))

		goto skipNumStrConversion

	default:

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'floatingPointNumber' is an invalid type!\n"+
			"'floatingPointNumber' is unsupported type '%T'\n",
			ePrefix.String(),
			floatingPointNumber)

		return signedPureNumStr, err
	}

	signedPureNumStr = strconv.FormatFloat(
		float64Num, 'f', -1, 64)

skipNumStrConversion:

	return signedPureNumStr, err
}

//	precisionBitsFromRequiredDigits
//
//	Generates the number of precision bits in the
//	mantissa of a big.Float number based on the
//	number of numerical digits required to produce
//	an accurate calculation result.
//
//	Be advised that the number of mantissa precision bits
//	required to store a process an accurate numeric value
//	includes both integer and fractional numeric digits.
//
//	The number of precision bits generated is an estimate
//	with a margin of error of plus or minus sixteen
//	(+ or - 16) bits.
//
//	The value of 'precisionBits' returned by this method
//	will always be a multiple of eight (+8).
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//
//	requiredIntegerDigits		int64
//
//		The number of integer digits required for the
//		pending calculation.
//
//		If this parameter has a value less than zero,
//		an error will be returned.
//
//		If the sum of parameters 'requiredIntegerDigits'
//		and 'requiredFractionalDigits' is equal to zero,
//		an error will be returned.
//
//	requiredFractionalDigits	int64
//
//		The number of fractional digits required to
//		ensure accuracy for the pending calculation.
//
//		If this parameter has a value less than zero,
//		an error will be returned.
//
//		If the sum of parameters 'requiredIntegerDigits'
//		and 'requiredFractionalDigits' is equal to zero,
//		an error will be returned.
//
//	requestedBufferDigits		int64
//
//		The number of extra numerical digits required to
//		ensure accuracy for the pending calculation. It
//		is generally a good idea to add space for extra
//		numerical digits to accommodate rounding and/or
//		complex numerical calculations.
//
//		If this parameter has a value less than zero,
//		an error will be returned.
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
//	precisionBits				uint
//
//		If this method completes successfully, this
//		parameter will return the number of precision
//		bits required to store and accurately process
//		the number of numerical digits identified by
//		input parameters, 'requiredIntegerDigits',
//		'requiredFractionalDigits' and
//		'requestedBufferDigits'.
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
func (mathFloatHelpAtom *mathFloatHelperAtom) precisionBitsFromRequiredDigits(
	requiredIntegerDigits,
	requiredFractionalDigits,
	requestedBufferDigits int64,
	errPrefDto *ePref.ErrPrefixDto) (
	precisionBits uint,
	err error) {

	if mathFloatHelpAtom.lock == nil {
		mathFloatHelpAtom.lock = new(sync.Mutex)
	}

	mathFloatHelpAtom.lock.Lock()

	defer mathFloatHelpAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"mathFloatHelperAtom."+
			"precisionBitsFromRequiredDigits()",
		"")

	if err != nil {

		return precisionBits, err
	}

	if requiredIntegerDigits < 0 {

		err = fmt.Errorf("\n%v\n"+
			"Error: Input parameter 'requiredIntegerDigits' is invalid!\n"+
			"'requiredIntegerDigits' is less than zero and negative.\n"+
			"requiredIntegerDigits = '%v'\n",
			ePrefix.String(),
			requiredIntegerDigits)

		return precisionBits, err
	}

	if requiredFractionalDigits < 0 {

		err = fmt.Errorf("\n%v\n"+
			"Error: Input parameter 'requiredFractionalDigits' is invalid!\n"+
			"'requiredFractionalDigits' is less than zero and negative.\n"+
			"requiredFractionalDigits = '%v'\n",
			ePrefix.String(),
			requiredFractionalDigits)

		return precisionBits, err
	}

	if requestedBufferDigits < 0 {

		err = fmt.Errorf("\n%v\n"+
			"Error: Input parameter 'requestedBufferDigits' is invalid!\n"+
			"'requestedBufferDigits' is less than zero and negative.\n"+
			"requestedBufferDigits = '%v'\n",
			ePrefix.String(),
			requestedBufferDigits)

		return precisionBits, err
	}

	if requiredIntegerDigits == 0 &&
		requiredFractionalDigits == 0 {

		err = fmt.Errorf("\n%v\n"+
			"Error: The combination of 'requiredIntegerDigits' and 'requiredFractionalDigits' is invalid!\n"+
			"Both 'requiredIntegerDigits' and 'requiredFractionalDigits' are equal to zero.\n",
			ePrefix.String())

		return precisionBits, err
	}

	totalRequiredNumericalDigits :=
		requiredIntegerDigits +
			requiredFractionalDigits +
			requestedBufferDigits

	precisionBits,
		err = new(mathFloatHelperPreon).
		estimateDigitsToPrecision(
			totalRequiredNumericalDigits,
			ePrefix.XCpy(
				fmt.Sprintf("totalRequiredNumericalDigits= %v",
					totalRequiredNumericalDigits)))

	return precisionBits, err
}
