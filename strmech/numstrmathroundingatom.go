package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type numStrMathRoundingAtom struct {
	lock *sync.Mutex
}

// roundHalfAwayFromZero - Performs a rounding operation on the
// integer and fractional arrays contained in an instance of
// NumberStrKernel. This method performs applies the 'Round Half
// Away From Zero' algorithm.
//
// Examples of HalfAwayFromZero Rounding Algorithm
//
//	7.6 rounds away to 8
//	7.5 rounds away to 8
//	7.4 rounds to 7
//	-7.4 rounds to -7
//	-7.5 rounds away to -8
//	-7.6 rounds away to -8
//
// Reference:
//
//	https://www.mathsisfun.com/numbers/rounding-methods.html
//	https://en.wikipedia.org/wiki/Rounding
//	https://www.mathsisfun.com/rounding-numbers.html
//	https://www.vedantu.com/maths/rounding-methods
//	https://rounding.to/the-most-common-rounding-methods/
//	https://www.wikihow.com/Round-Numbers
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// No data validation is performed on 'numStrKernel' or
// 'roundToFractionalDigits'. Both are assumed to have been
// validated by upstream methods.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//		numStrKernel                    *NumberStrKernel
//		    - An instance of NumberStrKernel which contains the
//		      rune array of fractional digits which will be
//		      subjected to the 'Round Half Away From Zero' algorithm.
//
//		 roundToFractionalDigits        int
//		    - If the number of digits to the right of the decimal
//		      separator (a.k.a. decimal point) which will remain
//		      after this rounding algorithm is applied.
//
//	       Example:
//	        Floating Point Number  0.12345
//	        roundToFractionalDigits 3
//	        Floating Point Number after rounding: 0.123
//
//
//			errPrefDto          *ePref.ErrPrefixDto
//			   - This object encapsulates an error prefix string which is
//			     included in all returned error messages. Usually, it
//			     contains the name of the calling method or methods listed
//			     as a function chain.
//
//			     If no error prefix information is needed, set this parameter
//			     to 'nil'.
//
//			     Type ErrPrefixDto is included in the 'errpref' software
//			     package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If this method completes successfully, this returned error
//	     Type is set equal to 'nil'. If errors are encountered during
//	     processing, the returned error Type will encapsulate an error
//	     message.
//
//	     If an error message is returned, the text value for input
//	     parameter 'errPrefDto' (error prefix) will be prefixed or
//	     attached at the beginning of the error message.
func (nStrMathRoundAtom *numStrMathRoundingAtom) roundHalfAwayFromZero(
	numStrKernel *NumberStrKernel,
	roundToFractionalDigits int,
	errPrefDto *ePref.ErrPrefixDto) error {

	if nStrMathRoundAtom.lock == nil {
		nStrMathRoundAtom.lock = new(sync.Mutex)
	}

	nStrMathRoundAtom.lock.Lock()

	defer nStrMathRoundAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrMathRoundingAtom."+
			"roundHalfAwayFromZero()",
		"")

	if err != nil {
		return err
	}

	if numStrKernel == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrKernel' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return err
	}

	// roundToFractionalDigits MUST BE
	//  less than existingNumOfFracDigits

	lastIdx := roundToFractionalDigits - 1
	roundIdx := roundToFractionalDigits
	var carry = '0'

	numStrKernel.fractionalDigits.CharsArray[roundIdx] += 5

	if numStrKernel.fractionalDigits.CharsArray[roundIdx] > '9' {
		numStrKernel.fractionalDigits.CharsArray[roundIdx] = '0'
		carry = '1'
	}

	if carry == '1' && lastIdx > -1 {

		for i := lastIdx; i > -1; i-- {

			numStrKernel.fractionalDigits.CharsArray[i] += 1

			carry = '0'

			if numStrKernel.fractionalDigits.CharsArray[i] > '9' {
				carry = '1'
			}

			if carry == '0' {
				break
			}
		}

	}

	// Add carry digit to integer digits if necessary
	existingNumOfIntDigits := numStrKernel.GetNumberOfIntegerDigits()

	if carry == '1' && existingNumOfIntDigits > 0 {
		lastIdx = existingNumOfIntDigits - 1

		for i := lastIdx; i > -1; i-- {

			numStrKernel.integerDigits.CharsArray[i] += 1

			carry = '0'

			if numStrKernel.integerDigits.CharsArray[i] > '9' {
				carry = '1'
			}

			if carry == '0' {
				break
			}
		}
	}

	if carry == '1' {

		numStrKernel.integerDigits.CharsArray =
			append(
				[]rune{carry},
				numStrKernel.integerDigits.CharsArray...)
	}

	return err
}
