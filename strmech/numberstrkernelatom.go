package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// numberStrKernelAtom - Provides helper methods for type
// NumberStrKernel.
//
type numberStrKernelAtom struct {
	lock *sync.Mutex
}

// testValidityOfNumStrKernel - Receives a pointer to an instance
// of NumberStrKernel and performs a diagnostic analysis to
// determine if that instance is valid in all respects.
//
// If the input parameter 'numStrKernel' is determined to be
// invalid, this method will return a boolean flag ('isValid') of
// 'false'. In addition, an instance of type error ('err') will be
// returned configured with an appropriate error message.
//
// If the input parameter 'numStrKernel' is valid, this method will
// return a boolean flag ('isValid') of 'true' and the returned
// error type ('err') will be set to 'nil'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrKernel               *NumberStrKernel
//     - A pointer to an instance of NumberStrKernel. This object
//       will be subjected to diagnostic analysis in order to
//       determine if all the member variables contain valid
//       values.
//
//
//  errPrefDto                 *ePref.ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods listed
//       as a function chain.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       Type ErrPrefixDto is included in the 'errpref' software
//       package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  isValid                    bool
//     - If input parameter 'numStrKernel' is judged to be valid in
//       all respects, this return parameter will be set to 'true'.
//
//     - If input parameter 'numStrKernel' is found to be invalid,
//       this return parameter will be set to 'false'.
//
//
//  err                        error
//     - If input parameter 'numStrKernel' is judged to be valid in
//       all respects, this return parameter will be set to 'nil'.
//
//       If input parameter, 'numStrKernel' is found to be invalid,
//       this return parameter will be configured with an
//       appropriate error message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (numStrKernelAtom *numberStrKernelAtom) testValidityOfNumStrKernel(
	numStrKernel *NumberStrKernel,
	errPrefDto *ePref.ErrPrefixDto) (
	isValid bool,
	err error) {

	if numStrKernelAtom.lock == nil {
		numStrKernelAtom.lock = new(sync.Mutex)
	}

	numStrKernelAtom.lock.Lock()

	defer numStrKernelAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	isValid = false

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"negNumSearchSpecAtom."+
			"testValidityOfNegNumSearchSpec()",
		"")

	if err != nil {

		return isValid, err
	}

	if numStrKernel == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrKernel' is a nil pointer!\n",
			ePrefix.String())

		return isValid, err
	}

	lenIntDigits := numStrKernel.integerDigits.GetRuneArrayLength()

	lenFracDigits := numStrKernel.fractionalDigits.GetRuneArrayLength()

	if lenIntDigits == 0 &&
		lenFracDigits == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of NumberStrKernel is invalid!"+
			"Both Integer Digits and Fractional Digits are empty"+
			"and contain zero digits.\n",
			ePrefix.String())

		return isValid, err
	}

	if lenIntDigits == 0 &&
		lenFracDigits > 0 {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of NumberStrKernel is invalid!"+
			"The Fractional Digits rune array contains valid numeric digits."+
			"However, the Integer Digits rune array is empty and has a zero length.\n",
			ePrefix.String())

		return isValid, err

	}

	if !numStrKernel.numberSign.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of NumberStrKernel is invalid!"+
			"The Number Sign Value is invalid.\n"+
			"Valid Number Sign Values are:\n"+
			"   NumSignVal.Negative()\n"+
			"   NumSignVal.Zero()\n"+
			"   NumSignVal.Positive()\n"+
			"The current Number Sign Value is:\n"+
			"   Number Sign String Value = '%v'\n"+
			"  Number Sign Integer Value = '%v\n",
			ePrefix.String(),
			numStrKernel.numberSign.String(),
			numStrKernel.numberSign.XValueInt())

		return isValid, err
	}

	numValHasNonZeroVal := false

	for i := 0; i < lenIntDigits; i++ {

		if numStrKernel.integerDigits.CharsArray[i] < '0' &&
			numStrKernel.integerDigits.CharsArray[i] > '9' {

			err = fmt.Errorf("%v\n"+
				"Error: This instance of NumberStrKernel is invalid!"+
				"The Integer Digits rune array contains non-numeric characters.\n",
				ePrefix.String())

			return isValid, err
		}
	}

	for i := 0; i < lenIntDigits; i++ {

		if numStrKernel.integerDigits.CharsArray[i] >= '1' &&
			numStrKernel.integerDigits.CharsArray[i] <= '9' {
			numValHasNonZeroVal = true
			break
		}
	}

	for j := 0; j < lenFracDigits; j++ {

		if numStrKernel.fractionalDigits.CharsArray[j] < '0' &&
			numStrKernel.fractionalDigits.CharsArray[j] > '9' {

			err = fmt.Errorf("%v\n"+
				"Error: This instance of NumberStrKernel is invalid!"+
				"The Fractional Digits rune array contains non-numeric characters.\n",
				ePrefix.String())

			return isValid, err
		}

	}

	if !numValHasNonZeroVal {

		for j := 0; j < lenFracDigits; j++ {

			if numStrKernel.fractionalDigits.CharsArray[j] >= '1' &&
				numStrKernel.fractionalDigits.CharsArray[j] <= '9' {
				numValHasNonZeroVal = true
				break
			}

		}

	}

	if numValHasNonZeroVal &&
		numStrKernel.numberSign == NumSignVal.Zero() {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of NumberStrKernel is invalid!"+
			"The Number Sign Value is invalid.\n"+
			"NumberStrKernel has a non-zero numeric value.\n"+
			"However, Number Sign is equal to Zero.\n"+
			"Number Sign = NumSignVal.Zero()\n",
			ePrefix.String())

		return isValid, err
	}

	if numValHasNonZeroVal != numStrKernel.isNonZeroValue {

		if numValHasNonZeroVal == false &&
			numStrKernel.isNonZeroValue == true {

			err = fmt.Errorf("%v\n"+
				"Error: This instance of NumberStrKernel is invalid!"+
				"The Number Sign Value is invalid.\n"+
				"NumberStrKernel has a zero numeric value.\n"+
				"However, internal flag numStrKernel.isNonZeroValue\n"+
				"is set to 'true'.\n",
				ePrefix.String())

			return isValid, err

		}

		if numValHasNonZeroVal == true &&
			numStrKernel.isNonZeroValue == false {

			err = fmt.Errorf("%v\n"+
				"Error: This instance of NumberStrKernel is invalid!"+
				"The Number Sign Value is invalid.\n"+
				"NumberStrKernel has a non-zero numeric value.\n"+
				"However, internal flag numStrKernel.isNonZeroValue\n"+
				"is set to 'false'.\n",
				ePrefix.String())

			return isValid, err

		}
	}

	isValid = true

	return isValid, err
}

// ptr - Returns a pointer to a new instance of
// numberStrKernelAtom.
//
func (numStrKernelAtom numberStrKernelAtom) ptr() *numberStrKernelAtom {

	if numStrKernelAtom.lock == nil {
		numStrKernelAtom.lock = new(sync.Mutex)
	}

	numStrKernelAtom.lock.Lock()

	defer numStrKernelAtom.lock.Unlock()

	return &numberStrKernelAtom{
		lock: new(sync.Mutex),
	}
}
