package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
	"time"
)

type textLineSpecTimerLinesAtom struct {
	lock *sync.Mutex
}

// getMaxTimerLabelLength - Returns the longest string length of
// the three text labels contained within an instance of
// TextLineSpecTimerLines. The maximum timer label length is
// therefore computed by comparing the string lengths of these
// three text label strings:
//  (1) TextLineSpecTimerLines.startTimeLabel
//  (2) TextLineSpecTimerLines.endTimeLabel
//  (3) TextLineSpecTimerLines.timeDurationLabel
//
func (txtTimerLinesAtom textLineSpecTimerLinesAtom) getMaxTimerLabelLength(
	txtTimerLines *TextLineSpecTimerLines) int {

	if txtTimerLinesAtom.lock == nil {
		txtTimerLinesAtom.lock = new(sync.Mutex)
	}

	txtTimerLinesAtom.lock.Lock()

	defer txtTimerLinesAtom.lock.Unlock()

	maxLabelLen := 0

	if txtTimerLines == nil {
		return maxLabelLen
	}

	if len(txtTimerLines.startTimeLabel) > maxLabelLen {
		maxLabelLen = len(txtTimerLines.startTimeLabel)
	}

	if len(txtTimerLines.endTimeLabel) > maxLabelLen {
		maxLabelLen = len(txtTimerLines.endTimeLabel)
	}

	if len(txtTimerLines.timeDurationLabel) > maxLabelLen {
		maxLabelLen = len(txtTimerLines.timeDurationLabel)
	}

	if txtTimerLines.labelFieldLen > maxLabelLen {
		maxLabelLen = txtTimerLines.labelFieldLen
	}

	return maxLabelLen
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

// equal - Receives pointers to two TextLineSpecTimerLines
// instances and proceeds to compare the member data elements to
// determine whether they are equal.
//
// If the data elements of both input parameters 'txtTimerLinesOne'
// and 'txtTimerLinesTwo' are equal in all respects, this method
// returns a boolean value of 'true'. Otherwise this method returns
// 'false'.
//
func (txtTimerLinesAtom *textLineSpecTimerLinesAtom) equal(
	txtTimerLinesOne *TextLineSpecTimerLines,
	txtTimerLinesTwo *TextLineSpecTimerLines) bool {

	if txtTimerLinesAtom.lock == nil {
		txtTimerLinesAtom.lock = new(sync.Mutex)
	}

	txtTimerLinesAtom.lock.Lock()

	defer txtTimerLinesAtom.lock.Unlock()

	if txtTimerLinesOne == nil {

		return false
	}

	if txtTimerLinesTwo == nil {

		return false
	}

	sMechPreon := strMechPreon{}

	if !sMechPreon.equalRuneArrays(
		txtTimerLinesOne.startTimeLabel,
		txtTimerLinesTwo.startTimeLabel) {
		return false
	}

	if txtTimerLinesOne.startTime !=
		txtTimerLinesTwo.startTime {
		return false
	}

	if !sMechPreon.equalRuneArrays(
		txtTimerLinesOne.endTimeLabel,
		txtTimerLinesTwo.endTimeLabel) {
		return false
	}

	if txtTimerLinesOne.endTime !=
		txtTimerLinesTwo.endTime {
		return false
	}

	if txtTimerLinesOne.timeFormat !=
		txtTimerLinesTwo.timeFormat {
		return false
	}

	if !sMechPreon.equalRuneArrays(
		txtTimerLinesOne.timeDurationLabel,
		txtTimerLinesTwo.timeDurationLabel) {
		return false
	}

	if txtTimerLinesOne.labelFieldLen !=
		txtTimerLinesTwo.labelFieldLen {
		return false
	}

	if txtTimerLinesOne.labelJustification !=
		txtTimerLinesTwo.labelJustification {
		return false
	}

	if !sMechPreon.equalRuneArrays(
		txtTimerLinesOne.labelOutputSeparationChars,
		txtTimerLinesTwo.labelOutputSeparationChars) {
		return false
	}

	return true
}

// testValidityOfEndTime - Tests the validity of the ending
// time parameter for an instance of TextLineSpecTimerLines.
//
// The purpose of a TextLineSpecTimerLines instance is to capture
// all the essential elements of a timer event and format that
// information for text display output or printing.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  endTime                    time.Time
//     - The ending time parameter must be a non-zero value.
//       If a zero value is submitted the time value will be
//       declared invalid.
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
//     - If input parameter 'endTime' is judged to be valid
//       in all respects, this return parameter will be set to
//       'true'.
//
//     - If input parameter 'endTime' is found to be invalid,
//       this return parameter will be set to 'false'.
//
//
//  err                        error
//     - If input parameter 'endTime' is judged to be valid
//       in all respects, this return parameter will be set to
//       'nil'.
//
//       If input parameter, 'endTime' is found to be invalid,
//       this return parameter will be configured with an
//       appropriate error message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (txtTimerLinesAtom *textLineSpecTimerLinesAtom) testValidityOfEndTime(
	endTime time.Time,
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
		"textLineSpecTimerLinesAtom.testValidityOfEndTime()",
		"")

	if err != nil {
		return isValid, err
	}

	if endTime.IsZero() {
		err = fmt.Errorf("%v\n"+
			"Error: 'endTime' is invalid!\n"+
			"The 'endTime' time value is zero.",
			ePrefix.String())
		return isValid, err
	}

	isValid = true

	return isValid, err
}

// testValidityOfStartTime - Tests the validity of the starting
// time parameter for an instance of TextLineSpecTimerLines.
//
// The purpose of a TextLineSpecTimerLines instance is to capture
// all the essential elements of a timer event and format that
// information for text display output or printing.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  startTime                  time.Time
//     - The starting time parameter must be a non-zero
//       value. If a zero value is submitted the time value
//       will be declared invalid.
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
//     - If input parameter 'startTime' is judged to be valid
//       in all respects, this return parameter will be set to
//       'true'.
//
//     - If input parameter 'startTime' is found to be invalid,
//       this return parameter will be set to 'false'.
//
//
//  err                        error
//     - If input parameter 'startTime' is judged to be valid
//       in all respects, this return parameter will be set to
//       'nil'.
//
//       If input parameter, 'startTime' is found to be invalid,
//       this return parameter will be configured with an
//       appropriate error message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (txtTimerLinesAtom *textLineSpecTimerLinesAtom) testValidityOfStartTime(
	startTime time.Time,
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
		"textLineSpecTimerLinesAtom.testValidityOfStartTime()",
		"")

	if err != nil {
		return isValid, err
	}

	if startTime.IsZero() {
		err = fmt.Errorf("%v\n"+
			"Error: 'startTime' is invalid!\n"+
			"The 'startTime' time value is zero.",
			ePrefix.String())
		return isValid, err
	}

	isValid = true

	return isValid, err
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
//     - If input parameter 'txtTimerLines' is judged to be valid
//       in all respects, this return parameter will be set to
//       'true'.
//
//     - If input parameter 'txtTimerLines' is found to be invalid,
//       this return parameter will be set to 'false'.
//
//
//  err                        error
//     - If input parameter 'txtTimerLines' is judged to be valid
//       in all respects, this return parameter will be set to
//       'nil'.
//
//       If input parameter, 'txtTimerLines' is found to be
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

	if txtTimerLines.labelFieldLen > 1000000 {
		err = fmt.Errorf("%v\n"+
			"Error: 'txtTimerLines.labelFieldLen' is invalid!\n"+
			"'txtTimerLines.labelFieldLen' is greater than one million "+
			"(1,000,000)'\n",
			ePrefix.String())

		return isValid, err
	}

	if len(txtTimerLines.timeFormat) == 0 {
		txtTimerLines.timeFormat =
			textSpecificationMolecule{}.ptr().
				getDefaultTimeFormat()
	}

	txtTimerLinesElectron := textLineSpecTimerLinesElectron{}

	if len(txtTimerLines.startTimeLabel) == 0 {
		txtTimerLines.startTimeLabel =
			txtTimerLinesElectron.getDefaultStartTimeLabel()
	}

	if len(txtTimerLines.endTimeLabel) == 0 {
		txtTimerLines.endTimeLabel =
			txtTimerLinesElectron.getDefaultEndTimeLabel()
	}

	if len(txtTimerLines.timeDurationLabel) == 0 {
		txtTimerLines.timeDurationLabel =
			txtTimerLinesElectron.getDefaultTimeDurationLabel()
	}

	if len(txtTimerLines.labelOutputSeparationChars) == 0 {
		txtTimerLines.labelOutputSeparationChars =
			txtTimerLinesElectron.getDefaultLabelOutputSeparationCharsLabel()
	}

	if txtTimerLines.labelFieldLen < -1 {
		txtTimerLines.labelFieldLen = -1
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

	isValid = true

	return isValid, err
}
