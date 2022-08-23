package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type NumStrSignedNum struct {
	lock *sync.Mutex
}

// GetNumStr - Returns a formatted number string for the current
// instance of NumStrSignedNum.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	numStrKernel				NumberStrKernel
//
//		An instance of NumberStrKernel containing the numeric
//		digits which will be used to create and format the
//		returned number string.
//
//
//	signedNumFormatSpec			SignedNumberFormatSpec
//
//		This instance of SignedNumberFormatSpec contains all
//		the parameters required to format a signed number for
//		text display.
//
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix
//			information.
//
//		3. []string A one-dimensional slice of strings
//			containing error prefix information.
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context
//		   information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto.
//			Information from this object will be copied for use
//			in error and informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of
//			ErrPrefixDto. Information from this object will be
//			copied for use in error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method
//			generating a two-dimensional slice of strings
//			containing error prefix and error context
//			information.
//
//		If parameter 'errorPrefix' is NOT convertible to one
//		of the valid types listed above, it will be
//		considered invalid and trigger the return of an
//		error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included
//		in the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// # Return Values
//
//	numberStr					string
//
//		If this method completes successfully, this parameter
//		will return a signed numeric value formatted as a
//		number string.
//
//	err							error
//
//		If this method completes successfully and no errors
//		are encountered this return value is set to 'nil'.
//		Otherwise, if errors are encountered, this return
//		value will contain an appropriate error message.
//
//		If an error message is returned, the text value of
//		input parameter 'errorPrefix' will be inserted or
//		prefixed at	the beginning of the error message.
func (nStrSignedNum *NumStrSignedNum) GetNumStr(
	numStrKernel NumberStrKernel,
	signedNumFormatSpec SignedNumberFormatSpec,
	errorPrefix interface{}) (
	numberStr string,
	err error) {

	if nStrSignedNum.lock == nil {
		nStrSignedNum.lock = new(sync.Mutex)
	}

	nStrSignedNum.lock.Lock()

	defer nStrSignedNum.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldSpecLabel.CopyIn()",
		"")

	if err != nil {
		return numberStr, err
	}

	if numStrKernel.GetNumberOfIntegerDigits() == 0 &&
		numStrKernel.GetNumberOfFractionalDigits() == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input Parameter 'numStrKernel' is invalid!\n"+
			"'numStrKernel' is empty and contains zero integer and\n"+
			"zero fractional numeric digits.\n",
			ePrefix.String())

		return numberStr, err
	}

	return new(numStrSignedNumNanobot).
		formatSignedNumStr(
			numStrKernel,
			signedNumFormatSpec,
			ePrefix.XCpy("numStrKernel"))

}

// numStrSignedNumNanobot - Provides helper methods for
// type NumStrSignedNum.
type numStrSignedNumNanobot struct {
	lock *sync.Mutex
}

func (nStrSignedNumNanobot numStrSignedNumNanobot) formatSignedNumStr(
	numStrKernel NumberStrKernel,
	signedNumFormatSpec SignedNumberFormatSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	numStr string,
	err error) {

	if nStrSignedNumNanobot.lock == nil {
		nStrSignedNumNanobot.lock = new(sync.Mutex)
	}

	nStrSignedNumNanobot.lock.Lock()

	defer nStrSignedNumNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrSignedNumNanobot."+
			"formatSignedNumStr()",
		"")

	if err != nil {
		return numStr, err
	}

	if numStrKernel.GetNumberOfIntegerDigits() == 0 &&
		numStrKernel.GetNumberOfFractionalDigits() == 0 {
		numStr = "0"

		return numStr, err
	}

	var newNumStrKernel NumberStrKernel

	newNumStrKernel,
		err = numStrKernel.CopyOut(
		ePrefix.XCpy(
			"newNumStrKernel<-numStrKernel"))

	if err != nil {
		return numStr, err
	}

	var roundingSpec NumStrRoundingSpec

	roundingSpec,
		err = signedNumFormatSpec.GetRoundingSpec(
		ePrefix.XCpy(
			"roundingSpec<-signedNumFormatSpec"))

	if err != nil {
		return numStr, err
	}

	// Performing fractional digit rounding
	err = new(numStrMathNanobot).roundNumStrKernel(
		&newNumStrKernel,
		roundingSpec,
		ePrefix.XCpy(
			"newNumStrKernel Rounding"))

	if err != nil {
		return numStr, err
	}

	var decSeparator DecimalSeparatorSpec

	decSeparator,
		err = signedNumFormatSpec.GetDecSeparatorSpec(
		ePrefix.XCpy(
			"decSeparator<-signedNumFormatSpec"))

	if err != nil {
		return numStr, err
	}

	if newNumStrKernel.GetNumberOfFractionalDigits() > 0 &&
		decSeparator.GetNumberOfSeparatorChars() == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: This is a floating point number and the number\n"+
			"of decimal separator characters specified is zero.\n"+
			"Input parameter 'signedNumFormatSpec.DecSeparator'\n"+
			"is invalid!\n",
			ePrefix.String())

		return numStr, err
	}

	return numStr, err
}
