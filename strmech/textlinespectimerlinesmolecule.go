package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
	"time"
)

type textLineSpecTimerLinesMolecule struct {
	lock *sync.Mutex
}

// setTxtLineSpecTimerLines - Receives a pointer to an instance of
// TextLineSpecTimerLines and proceeds to overwrite and set all the
// member variable data values using the input parameters provided.
//
// Be advised that the pre-existing data fields in input parameter
// 'txtTimerLines' will be overwritten and deleted.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  txtTimerLines              *TextLineSpecTimerLines
//     - A pointer to an instance of TextLineSpecTimerLines. All of
//       the member variable data values for this instance will be
//       overwritten and reset to new values extracted from the
//       following input parameters.
//
//       Be advised that the pre-existing data fields in input
//       parameter 'txtTimerLines' will be overwritten and deleted.
//
//
//  startTimeLabel             []rune
//     - An array of runes containing the text characters
//       constituting the starting time text label.
//
//       If this array is submitted as a zero length rune array,
//       'startTimeLabel' will be assigned a default value of
//       "Start Time".
//
//
//  startTime                  time.Time
//     - A time value which will be used in conjunction with
//       'endTime' parameter to compute the time duration or
//       elapsed for the timer event.
//
//       If this parameter is submitted as a zero time value,
//       'startTime' will be defaulted to value of July 4, 1776
//       9:30AM UTC.
//
//
//  endTimeLabel               []rune
//     - An array of runes containing the text characters
//       constituting the ending time text label.
//
//       If this array is submitted as a zero length rune array,
//       'endTimeLabel' will be assigned a default value of
//       "End Time".
//
//
//  endTime                    time.Time
//     - A time value which will be used in conjunction with
//       'startTime' parameter to compute the time duration or
//       elapsed for the timer event.
//
//       If this parameter is submitted as a zero time value,
//       'endTime' will be defaulted to value of July 4, 1776
//       9:30AM UTC.
//
//
//  timeFormat                 string
//     - This string holds the time format parameters used to
//       format starting time and ending time values for text
//       display and output.
//
//       If this parameter is submitted as an empty string,
//       parameter 'timeFormat' will be assigned a default value
//       of "2006-01-02 15:04:05.000000000 -0700 MST"
//
//
//  timeDurationLabel          []rune
//     - The text label used to describe the time duration or
//       elapsed time computed from the 'startTime' and 'endTime'
//       parameters.
//
//       If this array is submitted as a zero length rune array,
//       'timeDurationLabel' will be assigned a default value of
//       "Elapsed Time".
//
//
//  labelFieldLen              int
//     - The length of the text fields which will be used to
//       position and display the three text labels provided by
//       input parameters 'startTimeLabel', 'endTimeLabel' and
//       'timeDurationLabel'.
//
//       If labelFieldLen is less than the length of the longest
//       text label it will be defaulted to the length of the
//       longest text label.
//
//
//  labelJustification         TextJustify
//     - An enumeration which specifies the justification of the
//       the three text labels 'startTimeLabel', 'endTimeLabel' and
//       'timeDurationLabel' within the field length specified by
//       'labelFieldLen'.
//
//       If the field length is greater than the length of the
//       longest of the three text labels, label justification must
//       be equal to one of these three valid values:
//           TextJustify(0).Left()
//           TextJustify(0).Right()
//           TextJustify(0).Center()
//
//       The abbreviated text justification enumeration syntax can
//       also be used:
//
//           TxtJustify.Left()
//           TxtJustify.Right()
//           TxtJustify.Center()
//
//
//  labelOutputSeparationChars []rune
//     - This rune array contains the character or characters which
//       will be used to separate the text labels ('startTimeLabel',
//       'endTimeLabel' and 'timeDurationLabel') from the output or
//       data values displayed on the same line.
//       Example:
//        Start Time[sep chars]2010-01-02 15:04:05.000000000 -0700 MST
//
//       Often this parameter is set to a single white space
//       character (' ') or a colon plus white space character,
//       ([]rune{':', ' "}).
//
//
//       If this array is submitted as a zero length rune array,
//       'labelOutputSeparationChars' will be assigned a default value of
//       []rune{':', ' "}. Example Output:
//        Start Time: 2010-01-02 15:04:05.000000000 -0700 MST
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
//  error
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
func (txtTimerLinesMolecule *textLineSpecTimerLinesMolecule) setTxtLineSpecTimerLines(
	txtTimerLines *TextLineSpecTimerLines,
	startTimeLabel []rune,
	startTime time.Time,
	endTimeLabel []rune,
	endTime time.Time,
	timeFormat string,
	timeDurationLabel []rune,
	labelFieldLen int,
	labelJustification TextJustify,
	labelOutputSeparationChars []rune,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtTimerLinesMolecule.lock == nil {
		txtTimerLinesMolecule.lock = new(sync.Mutex)
	}

	txtTimerLinesMolecule.lock.Lock()

	defer txtTimerLinesMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecTimerLinesMolecule.setTxtLineSpecTimerLines()",
		"")

	if err != nil {
		return err
	}

	if txtTimerLines == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtTimerLines' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if len(startTimeLabel) == 0 {
		startTimeLabel = []rune("Start Time")
	}

	if len(endTimeLabel) == 0 {
		endTimeLabel = []rune("End Time")
	}

	if len(timeDurationLabel) == 0 {
		timeDurationLabel = []rune("Elapsed Time")
	}

	if len(labelOutputSeparationChars) == 0 {
		labelOutputSeparationChars = []rune{':', ' '}
	}

	if !labelJustification.XIsValid() {
		err = fmt.Errorf("%v\n"+
			"Error: 'labelJustification' is invalid!\n"+
			"'labelJustification' should be 'Left',\n"+
			"'Right' or 'Center'.\n"+
			"'labelJustification' string value  = '%v'\n"+
			"'labelJustification' integer value = '%v'\n",
			ePrefix.String(),
			labelJustification.String(),
			labelJustification.XValueInt())

		return err
	}

	if labelFieldLen < -1 {
		labelFieldLen = -1
	}

	if labelFieldLen > 1000000 {
		err = fmt.Errorf("%v\n"+
			"Error: 'labelFieldLen' is invalid!\n"+
			"'labelFieldLen' is greater than one million "+
			"(1,000,000)'\n",
			ePrefix.String())

		return err
	}

	if len(timeFormat) == 0 {
		timeFormat =
			textSpecificationMolecule{}.ptr().
				getDefaultTimeFormat()
	}

	if startTime.IsZero() {
		startTime =
			textLineSpecTimerLinesElectron{}.ptr().
				getDefaultTime()
	}

	if endTime.IsZero() {
		endTime =
			textLineSpecTimerLinesElectron{}.ptr().
				getDefaultTime()
	}

	sMechPreon := strMechPreon{}

	err = sMechPreon.copyRuneArrays(
		&txtTimerLines.startTimeLabel,
		&startTimeLabel,
		true,
		ePrefix.XCtx("startTimeLabel"))

	if err != nil {
		return err
	}

	txtTimerLines.startTime = startTime

	err = sMechPreon.copyRuneArrays(
		&txtTimerLines.endTimeLabel,
		&endTimeLabel,
		true,
		ePrefix.XCtx("endTimeLabel"))

	if err != nil {
		return err
	}

	txtTimerLines.endTime = endTime

	txtTimerLines.timeFormat = timeFormat

	err = sMechPreon.copyRuneArrays(
		&txtTimerLines.timeDurationLabel,
		&timeDurationLabel,
		true,
		ePrefix.XCtx("timeDurationLabel"))

	if err != nil {
		return err
	}

	txtTimerLines.labelFieldLen =
		labelFieldLen

	txtTimerLines.labelJustification =
		labelJustification

	err = sMechPreon.copyRuneArrays(
		&txtTimerLines.labelOutputSeparationChars,
		&labelOutputSeparationChars,
		true,
		ePrefix.XCtx(
			"labelOutputSeparationChars"))

	return err
}

// ptr - Returns a pointer to a new instance of
// textLineSpecTimerLinesMolecule.
//
func (txtTimerLinesMolecule textLineSpecTimerLinesMolecule) ptr() *textLineSpecTimerLinesMolecule {

	if txtTimerLinesMolecule.lock == nil {
		txtTimerLinesMolecule.lock = new(sync.Mutex)
	}

	txtTimerLinesMolecule.lock.Lock()

	defer txtTimerLinesMolecule.lock.Unlock()

	return &textLineSpecTimerLinesMolecule{
		lock: new(sync.Mutex),
	}
}
