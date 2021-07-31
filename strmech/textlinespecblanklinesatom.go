package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type textLineSpecBlankLinesAtom struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// textLineSpecBlankLinesAtom.
//
func (txtBlankLinesAtom textLineSpecBlankLinesAtom) ptr() *textLineSpecBlankLinesAtom {

	if txtBlankLinesAtom.lock == nil {
		txtBlankLinesAtom.lock = new(sync.Mutex)
	}

	txtBlankLinesAtom.lock.Lock()

	defer txtBlankLinesAtom.lock.Unlock()

	return &textLineSpecBlankLinesAtom{
		lock: new(sync.Mutex),
	}
}

// testValidityOfTextLineSpecBlankLines - Receives a pointer to an
// instance of TextLineSpecBlankLines and performs a diagnostic
// analysis to determine if that instance is valid in all respects.
//
// If the input parameter 'txtBlankLines' is determined to be
// invalid, this method will return a boolean flag ('isValid') of
// 'false'. In addition, an instance of type error ('err') will be
// returned configured with an appropriate error message.
//
// If the input parameter 'txtBlankLines' is valid, this method
// will return a boolean flag ('isValid') of 'true' and the
// returned error type ('err') will be set to 'nil'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  txtBlankLines              *TextLineSpecBlankLines
//     - A pointer to an instance of TextLineSpecBlankLines. This
//       object will be subjected to diagnostic analysis in order
//       to determine if all the member variables contain valid
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
//     - If input parameter 'txtBlankLines' is judged to be valid
//       in all respects, this return parameter will be set to
//       'true'.
//
//     - If input parameter 'txtBlankLines' is found to be invalid,
//       this return parameter will be set to 'false'.
//
//
//  err                        error
//     - If input parameter 'txtBlankLines' is judged to be valid
//       in all respects, this return parameter will be set to
//       'nil'.
//
//       If input parameter, 'txtBlankLines' is found to be
//       invalid, this return parameter will be configured with an
//       appropriate error message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (txtBlankLinesAtom *textLineSpecBlankLinesAtom) testValidityOfTextLineSpecBlankLines(
	txtBlankLines *TextLineSpecBlankLines,
	errPrefDto *ePref.ErrPrefixDto) (
	isValid bool,
	err error) {

	if txtBlankLinesAtom.lock == nil {
		txtBlankLinesAtom.lock = new(sync.Mutex)
	}

	txtBlankLinesAtom.lock.Lock()

	defer txtBlankLinesAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	isValid = false

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecBlankLinesAtom.testValidityOfTextLineSpecBlankLines()",
		"")

	if err != nil {
		return isValid, err
	}

	if txtBlankLines == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtBlankLines' is a nil pointer!\n",
			ePrefix.String())

		return isValid, err
	}

	if len(txtBlankLines.newLineChars) == 0 {
		txtBlankLines.newLineChars = []rune{'\n'}
	}

	if txtBlankLines.numBlankLines == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: The specified number of blank lines is zero!\n"+
			"'TextLineSpecBlankLines.numBlankLines' must be greater than zero.\n",
			ePrefix.String())

		return isValid, err
	}

	if txtBlankLines.numBlankLines > 1000000 {

		err = fmt.Errorf("%v\n"+
			"Error:  The number of specified blank lines is\n"+
			"greater than 1,000,000!\n",
			ePrefix.String())

		return isValid, err
	}

	if txtBlankLines.numBlankLines < 0 {

		err = fmt.Errorf("%v\n"+
			"Error: The specified number of blank lines is less than zero!\n"+
			"'TextLineSpecBlankLines.numBlankLines' must be greater than zero.\n",
			ePrefix.String())

		return isValid, err
	}

	isValid = true

	return isValid, err
}
