package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"sync"
)

// mathBigRatHelperAtom
//
// Provides helper methods for type MathBigRatHelper.
type mathBigRatHelperAtom struct {
	lock *sync.Mutex
}

//	ratToBigFloat
//
//	Converts a big.Rat number to a big.Float number.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	bigRatNum					*big.Rat
//
//		The rational number which will be converted to a
//		big.Float numeric value.
//
//	roundToFractionalDigits 	int
//
//		Controls the number of fractional digits returned
//		as a big float number ('bigFloatNum').
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
//	bigFloatNum					*big.Float
//
//		If this method completes successfully, this
//		parameter will return a big.Float number
//		representing the numeric value of the rational
//		number passed through input parameter,
//		'bigRatNum'.
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
func (mathBigRatHelpAtom *mathBigRatHelperAtom) ratToBigFloat(
	bigRatNum *big.Rat,
	roundToFractionalDigits int,
	errPrefDto *ePref.ErrPrefixDto) (
	bigFloatNum *big.Float,
	err error) {

	if mathBigRatHelpAtom.lock == nil {
		mathBigRatHelpAtom.lock = new(sync.Mutex)
	}

	mathBigRatHelpAtom.lock.Lock()

	defer mathBigRatHelpAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	bigFloatNum = big.NewFloat(0)

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"mathBigRatHelperAtom."+
			"ratToBigFloat()",
		"")

	if err != nil {

		return bigFloatNum, err
	}

	if bigRatNum == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'bigRatNum' is a nil pointer!\n",
			ePrefix.String())

		return bigFloatNum, err
	}

	ratFloatStr := bigRatNum.FloatString(roundToFractionalDigits)

	var ok bool

	bigFloatNum,
		ok = bigFloatNum.SetString(ratFloatStr)

	if !ok {
		err = fmt.Errorf("%v\n"+
			"Error: bigFloatNum.SetString(ratFloatStr) Failed!\n"+
			"ratFloatStr = '%v'\n",
			ePrefix.String(),
			ratFloatStr)

		return bigFloatNum, err

	}

	return bigFloatNum, err
}
