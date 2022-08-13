package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type numStrMathAtom struct {
	lock *sync.Mutex
}

// extendFractionalDigits - This method will add a specified
// number of zero digits ('0') to the end of the fractional
// digits array within an instance of NumberStrKernel.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//		numStrKernel                    *NumberStrKernel
//	    - An instance of NumberStrKernel which contains the
//	      rune array of fractional digits which will be
//	      extended to comply with the requirement for a
//	      specific number of fractional digits.
//
//	 requestedNumOfFractionalDigits  int
//	    - If the number of digits is the fractional digits array
//	      contained in parameter 'numStrKernel' is less than
//	      this integer value, that fractional digits array will
//	      be extended with zero digits ('0') to meet the
//	      required number of fractional digits. Zero digits
//	      ('0') will be appended to the fractional digits array
//	      until the total length of the fractional digits array
//	      equals the integer value of
//	      'requestedNumOfFractionalDigits'.
//
//	      If the actual number of fractional digits is greater
//	      than or equal to 'requestedNumOfFractionalDigits',
//	      no action will be taken and the method will return
//	      without error.
//
//
//		errPrefDto          *ePref.ErrPrefixDto
//		   - This object encapsulates an error prefix string which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods listed
//		     as a function chain.
//
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
//
//		     Type ErrPrefixDto is included in the 'errpref' software
//		     package, "github.com/MikeAustin71/errpref".
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
func (nStrMathAtom *numStrMathAtom) extendFractionalDigits(
	numStrKernel *NumberStrKernel,
	requestedNumOfFractionalDigits int,
	errPrefDto *ePref.ErrPrefixDto) error {

	if nStrMathAtom.lock == nil {
		nStrMathAtom.lock = new(sync.Mutex)
	}

	nStrMathAtom.lock.Lock()

	defer nStrMathAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrMathAtom."+
			"extendFractionalDigits()",
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

	if requestedNumOfFractionalDigits < 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'requestedNumOfFractionalDigits' is invalid!\n"+
			"'requestedNumOfFractionalDigits' has a value which is less than zero (0).\n"+
			"requestedNumOfFractionalDigits = '%v'\n",
			ePrefix.String(),
			requestedNumOfFractionalDigits)

		return err

	}

	existingNumOfFracDigits := numStrKernel.GetNumberOfFractionalDigits()

	if requestedNumOfFractionalDigits <=
		existingNumOfFracDigits {

		// Nothing to do.
		return err
	}

	requiredNumOfFracDigits :=
		requestedNumOfFractionalDigits -
			existingNumOfFracDigits

	for i := 0; i < requiredNumOfFracDigits; i++ {

		err = numStrKernel.AddFractionalDigit(
			'0',
			ePrefix.XCpy(
				fmt.Sprintf("Add 0-digit #%v",
					i)))

		if err != nil {
			return err
		}
	}

	return err
}
