package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type textLineSpecTimerLinesAtom struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// textLineSpecTimerLinesAtom.
//
func (txtTimerLinesAtom textLineSpecTimerLinesAtom) ptr() *textLineSpecTimerLinesAtom {

	if txtTimerLinesAtom.lock == nil {
		txtTimerLinesAtom.lock = new(sync.Mutex)
	}

	txtTimerLinesAtom.lock.Lock()

	defer txtTimerLinesAtom.lock.Unlock()

	return &textLineSpecTimerLinesAtom{
		lock: new(sync.Mutex),
	}
}

// testValidityOfTxtSpecTimerLines - Receives a pointer to an
// instance of TextLineSpecTimerLines and performs a diagnostic
// analysis to determine if that instance is valid in all respects.
//
// If the input parameter 'txtTimerLines' is determined to be
// invalid, this method will return a boolean flag ('isValid') of
// 'false'. In addition, an instance of type error ('err') will be
// returned configured with an appropriate error message.
//
// If the input parameter 'txtTimerLines' is valid, this method
// will return a boolean flag ('isValid') of 'true' and the
// returned error type ('err') will be set to 'nil'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  txtTimerLines              *TextLineSpecTimerLines
//     - A pointer to an instance of TextLineSpecTimerLines. This
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
func (txtTimerLinesAtom *textLineSpecTimerLinesAtom) testValidityOfTxtSpecTimerLines(
	txtTimerLines *TextLineSpecTimerLines,
	errPrefDto *ePref.ErrPrefixDto) (
	isValid bool,
	err error) {

	if txtTimerLinesAtom.lock == nil {
		txtTimerLinesAtom.lock = new(sync.Mutex)
	}

	txtTimerLinesAtom.lock.Lock()

	defer txtTimerLinesAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	isValid = false

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecTimerLinesAtom.testValidityOfTxtSpecTimerLines()",
		"")

	if err != nil {
		return isValid, err
	}

	if txtTimerLines == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtTimerLines' is a nil pointer!\n",
			ePrefix.String())

		return isValid, err
	}

	if txtTimerLines.startTime.IsZero() {
		err = fmt.Errorf("%v\n"+
			"Error: 'txtTimerLines.startTime' is invalid!\n"+
			"'txtTimerLines.startTime' has a zero value.",
			ePrefix.String())

		return isValid, err
	}

	if txtTimerLines.endTime.IsZero() {
		err = fmt.Errorf("%v\n"+
			"Error: 'txtTimerLines.endTime' is invalid!\n"+
			"'txtTimerLines.endTime' has a zero value.",
			ePrefix.String())

		return isValid, err
	}

	if txtTimerLines.endTime.Before(txtTimerLines.startTime) {
		err = fmt.Errorf("%v\n"+
			"Error: 'txtTimerLines.endTime' is invalid!\n"+
			"'txtTimerLines.endTime' occurs before 'txtTimerLines.startTime'\n",
			ePrefix.String())

		return isValid, err
	}

	if len(txtTimerLines.startTimeLabel) == 0 {
		txtTimerLines.startTimeLabel = []rune("Start Time")
	}

	if len(txtTimerLines.endTimeLabel) == 0 {
		txtTimerLines.endTimeLabel = []rune("End Time")
	}

	if len(txtTimerLines.timeDurationLabel) == 0 {
		txtTimerLines.timeDurationLabel = []rune("Elapsed Time")
	}

	if len(txtTimerLines.labelOutputSeparationChars) == 0 {
		txtTimerLines.labelOutputSeparationChars = []rune{' '}
	}

	if !txtTimerLines.labelJustification.XIsValid() {
		err = fmt.Errorf("%v\n"+
			"Error: 'txtTimerLines.labelJustification' is invalid!\n"+
			"'txtTimerLines.labelJustification' should be 'Left',\n"+
			"'Right' or 'Center'.\n"+
			"'txtTimerLines.labelJustification' string value  = '%v'\n"+
			"'txtTimerLines.labelJustification' integer value = '%v'\n",
			ePrefix.String(),
			txtTimerLines.labelJustification.String(),
			txtTimerLines.labelJustification.XValueInt())

		return isValid, err
	}

	if txtTimerLines.labelFieldLen < -1 {
		err = fmt.Errorf("%v\n"+
			"Error: 'txtTimerLines.labelFieldLen' is invalid!\n"+
			"'txtTimerLines.labelFieldLen' is less than minus one (-1).'\n",
			ePrefix.String())

		return isValid, err
	}

	if len(txtTimerLines.timeFormat) == 0 {
		txtTimerLines.timeFormat =
			textSpecificationMolecule{}.ptr().
				getDefaultTimeFormat()
	}

	if txtTimerLines.labelFieldLen > 1000000 {
		err = fmt.Errorf("%v\n"+
			"Error: 'txtTimerLines.labelFieldLen' is invalid!\n"+
			"'txtTimerLines.labelFieldLen' is greater than one million "+
			"(1,000,000)'\n",
			ePrefix.String())

		return isValid, err
	}

	isValid = true

	return isValid, err
}
