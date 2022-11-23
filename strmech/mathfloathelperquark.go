package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"sync"
)

type mathFloatHelperQuark struct {
	lock *sync.Mutex
}

//	raiseToPositiveExponent
//
//	Receives a pointer to a big.Float floating point
//	number and raises that number to the power specified
//	by input parameter 'exponent'.
//
//		Example:	3.2 ^ 4 = 104.8576
//					base ^ exponent = raisedToExponent
//
//	This method will only process positive exponents.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	base						*big.Float
//
//		This floating point value will be raised to the
//		power of 'exponent' and returned to the calling
//		function.
//
//	exponent					int64
//
//		This value will be used to raise 'base' to the
//		power of 'exponent'.
//
//		Example:	3.2 ^ 4 = 104.8576
//					base ^ exponent = raisedToExponent
//
//	precisionBits				uint
//
//		The number of bits in the mantissa of the result
//		'raisedToExponent'. Effectively, this parameter
//		controls the precision and accuracy for the
//		calculation of 'base' raised to the power of
//		'exponent'.
//
//		If in doubt as to this number, identify the
//		total number of integer and fractional digits
//		required to store an accurate result and
//		multiply this number times four (+4).
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
//	raisedToExponent	*big.Float
//
//		If this method completes successfully, this will
//		return 'base' value raised to the power of the
//		'exponent' value.
//
//		Example:	3.2 ^ 4 = 104.8576
//					base ^ exponent = raisedToExponent
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
func (floatHelperQuark *mathFloatHelperQuark) raiseToPositiveExponent(
	base *big.Float,
	exponent int64,
	precisionBits uint,
	errPrefDto *ePref.ErrPrefixDto) (
	raisedToExponent *big.Float,
	err error) {

	if floatHelperQuark.lock == nil {
		floatHelperQuark.lock = new(sync.Mutex)
	}

	floatHelperQuark.lock.Lock()

	defer floatHelperQuark.lock.Unlock()

	raisedToExponent =
		new(big.Float).
			SetPrec(precisionBits).
			SetMode(big.AwayFromZero).
			SetInt64(0)

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"mathFloatHelperQuark."+
			"raiseToPositiveExponent()",
		"")

	if err != nil {
		return raisedToExponent, err
	}

	if exponent < 0 {

		err = fmt.Errorf("\n%v\n"+
			"Error: Input parameter 'exponent' is invalid!\n"+
			"'exponent' is less than zero and negative.\n"+
			"exponent = '%v'\n",
			ePrefix.String(),
			exponent)

		return raisedToExponent, err
	}

	if base == nil {

		err = fmt.Errorf("\n%v\n"+
			"Error: Input parameter 'base' is invalid!\n"+
			"'base' is a nil pointer.\n",
			ePrefix.String())

		return raisedToExponent, err
	}

	var ok bool
	_,
		ok = raisedToExponent.SetString("1.0")

	if !ok {

		err = fmt.Errorf("\n%v\n"+
			"Error: raisedToExponent.SetString(\"1.0\") Failed!\n",
			ePrefix.String())

		return raisedToExponent, err
	}

	if exponent == 0 {

		return raisedToExponent, err
	}

	for i := int64(0); i < exponent; i++ {

		raisedToExponent.Mul(raisedToExponent, base)
	}

	return raisedToExponent, err
}
