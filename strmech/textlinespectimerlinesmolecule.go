package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
	"time"
)

type textLineSpecTimerLinesMolecule struct {
	lock *sync.Mutex
}

// getFormattedText - Returns the formatted text generated by an
// instance of TextLineSpecTimerLines passed as an input parameter.
//
// Input parameter 'txtTimerLines' will provide the necessary
// information for generating text output which describes a timer
// event. The output will always consist of three lines of text
// specifying the event start time, end time and time duration or
// elapsed time.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  txtTimerLines              *TextLineSpecTimerLines
//     - A pointer to an instance of TextLineSpecTimerLines. The
//       member variables encapsulated by this instance will be
//       used to generate formatted text describing a time event
//       for text output display and printing.
//
//       If the member variable data values encapsulated by this
//       TextLineSpecTimerLines instance are found to be invalid,
//       this method will return an error.
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
//  string
//     - If this method completes successfully, a string of
//       formatted text will be generated from the data provided by
//       input parameter 'txtTimerLines'.
//
//
//  error
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (txtTimerLinesMolecule *textLineSpecTimerLinesMolecule) getFormattedText(
	txtTimerLines *TextLineSpecTimerLines,
	errPrefDto *ePref.ErrPrefixDto) (
	string,
	error) {

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
		"textLineSpecTimerLinesMolecule."+
			"getFormattedText()",
		"")

	if err != nil {
		return "", err
	}

	if txtTimerLines == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtTimerLines' is a nil pointer!\n",
			ePrefix.String())

		return "", err
	}

	txtTimerLinesAtom := textLineSpecTimerLinesAtom{}

	_,
		err = txtTimerLinesAtom.
		testValidityOfTxtSpecTimerLines(
			txtTimerLines,
			ePrefix.XCpy("txtTimerLines"))

	if err != nil {
		return "", err
	}

	// Text Formatting Setup
	var txtLabelLeftFiller *TextFieldSpecFiller

	var txtDescLabel *TextFieldSpecLabel

	var txtLabelRightFiller *TextFieldSpecFiller

	var txtOutputLabel *TextFieldSpecLabel

	stdLine := TextLineSpecStandardLine{}.New()

	// Begin First Line

	if len(txtTimerLines.labelLeftMarginChars) > 0 {

		txtLabelLeftFiller,
			err = TextFieldSpecFiller{}.NewPtrTextFillerRuneArray(
			txtTimerLines.labelLeftMarginChars,
			1,
			ePrefix.XCpy(
				"First Line: "+
					"txtTimerLines.labelLeftMarginChars"))

		if err != nil {
			return "", err
		}

		_,
			err = stdLine.AddTextField(
			txtLabelLeftFiller,
			ePrefix.XCpy(
				"First Line: txtLabelLeftFiller"))

		if err != nil {
			return "", err
		}

	}

	txtDescLabel,
		err = TextFieldSpecLabel{}.NewPtrTextLabelRunes(
		txtTimerLines.startTimeLabel,
		txtTimerLines.textLabelFieldLen,
		txtTimerLines.textLabelJustification,
		ePrefix.XCpy(
			"txtDescLabel: txtTimerLines.startTimeLabel"))

	if err != nil {
		return "", err
	}

	_,
		err = stdLine.AddTextField(
		txtDescLabel,
		ePrefix.XCpy(
			"First Line: txtDescLabel"+
				" - txtTimerLines.startTimeLabel"))

	if err != nil {
		return "", err
	}

	txtLabelRightFiller,
		err = TextFieldSpecFiller{}.NewPtrTextFillerRuneArray(
		txtTimerLines.labelRightMarginChars,
		1,
		ePrefix.XCpy(
			"First Line: txtTimerLines.labelRightMarginChars"))

	if err != nil {
		return "", err
	}

	_,
		err = stdLine.AddTextField(
		txtLabelRightFiller,
		ePrefix.XCpy(
			"First Line: txtLabelRightFiller"))

	if err != nil {
		return "", err
	}

	startTimeStr := txtTimerLines.startTime.Format(
		txtTimerLines.timeFormat)

	txtOutputLabel,
		err = TextFieldSpecLabel{}.NewPtrTextLabel(
		startTimeStr,
		-1,
		TxtJustify.Left(),
		ePrefix.XCpy(
			"startTimeStr: txtOutputLabel"))

	if err != nil {
		return "", err
	}

	_,
		err = stdLine.AddTextField(
		txtOutputLabel,
		ePrefix.XCpy(
			"First Line - startTimeStr: "+
				"txtOutputLabel"))

	if err != nil {
		return "", err
	}

	var err2 error

	sb := strings.Builder{}

	sb.Grow(512)

	_,
		err2 = sb.WriteString(stdLine.String())

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"Error returned by first standard line.\n"+
			"sb.WriteString(stdLine.String())\n"+
			"%v\n",
			ePrefix.ZCtxEmpty().String(),
			err2.Error())

		return "", err
	}

	stdLine.EmptyTextFields()

	// Begin Second Line

	if len(txtTimerLines.labelLeftMarginChars) > 0 {

		_,
			err = stdLine.AddTextField(
			txtLabelLeftFiller,
			ePrefix.XCpy(
				"Second Line - txtLabelLeftFiller"))

		if err != nil {
			return "", err
		}

	}

	err = txtDescLabel.SetTextRunes(
		txtTimerLines.endTimeLabel,
		ePrefix.XCpy(
			"Second Line - txtDescLabel: "+
				"txtTimerLines.endTimeLabel"))

	if err != nil {
		return "", err
	}

	_,
		err = stdLine.AddTextField(
		txtDescLabel,
		ePrefix.XCpy(
			"Second Line - txtDescLabel: "+
				"Add txtDescLabel"))

	if err != nil {
		return "", err
	}

	_,
		err =
		stdLine.AddTextField(
			txtLabelRightFiller,
			ePrefix.XCpy(
				"Second Line: txtLabelRightFiller"))

	if err != nil {
		return "", err
	}

	endTimeStr := txtTimerLines.endTime.Format(
		txtTimerLines.timeFormat)

	err = txtOutputLabel.SetText(
		endTimeStr,
		ePrefix.XCpy(
			"Second Line - txtOutputLabel: "+
				"endTimeStr"))

	if err != nil {
		return "", err
	}

	_,
		err = stdLine.AddTextField(
		txtOutputLabel,
		ePrefix.XCpy(
			"Second - endTimeStr: "+
				"txtOutputLabel"))

	if err != nil {
		return "", err
	}

	_,
		err2 = sb.WriteString(stdLine.String())

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"Error returned by second standard line.\n"+
			"sb.WriteString(stdLine.String())\n"+
			"%v\n",
			ePrefix.ZCtxEmpty().String(),
			err2.Error())

		return "", err
	}

	// Begin summary time duration lines

	totalLabelLen :=
		textLineSpecTimerLinesElectron{}.ptr().
			getTotalLabelLength(
				txtTimerLines.labelLeftMarginChars,
				txtTimerLines.startTimeLabel,
				txtTimerLines.endTimeLabel,
				txtTimerLines.timeDurationLabel,
				txtTimerLines.textLabelFieldLen,
				txtTimerLines.labelRightMarginChars)

	var txtFillerSumLeftMar *TextFieldSpecFiller

	txtFillerSumLeftMar,
		err = TextFieldSpecFiller{}.NewPtrTextFillerRuneArray(
		[]rune{' '},
		totalLabelLen,
		ePrefix.XCpy(
			"Summary Line Left Margin"))

	if err != nil {
		return "", err
	}

	var timeDurationStrs []string

	timeDurationStrs,
		err = textLineSpecTimerLinesElectron{}.ptr().
		computeTimeDuration(
			txtTimerLines.startTime,
			txtTimerLines.endTime,
			totalLabelLen,
			ePrefix.XCpy(
				"Timer Summary Line Calculation"))

	if err != nil {
		return "", err
	}

	for i := 0; i < len(timeDurationStrs); i++ {

		stdLine.EmptyTextFields()

		if i == 0 {

			if len(txtTimerLines.labelLeftMarginChars) > 0 {

				txtLabelLeftFiller,
					err = TextFieldSpecFiller{}.NewPtrTextFillerRuneArray(
					txtTimerLines.labelLeftMarginChars,
					1,
					ePrefix.XCpy(
						"Summary Third Line-First Segment: "+
							"txtTimerLines.labelLeftMarginChars"))

				if err != nil {
					return "", err
				}

				_,
					err = stdLine.AddTextField(
					txtLabelLeftFiller,
					ePrefix.XCpy(
						"Summary Third Line-First Segment: txtLabelLeftFiller"))

				if err != nil {
					return "", err
				}

			}

			txtDescLabel,
				err = TextFieldSpecLabel{}.NewPtrTextLabelRunes(
				txtTimerLines.timeDurationLabel,
				txtTimerLines.textLabelFieldLen,
				txtTimerLines.textLabelJustification,
				ePrefix.XCpy(
					"txtTimerLines.timeDurationLabel"))

			if err != nil {
				return "", err
			}

			_,
				err = stdLine.AddTextField(
				txtDescLabel,
				ePrefix.XCpy(
					fmt.Sprintf(
						"Time Summary Line #%v : "+
							"txtTimerLines.timeDurationLabel->"+
							"txtDescLabel", i)))

			if err != nil {
				return "", err
			}

			_,
				err =
				stdLine.AddTextField(
					txtLabelRightFiller,
					ePrefix.XCpy(
						fmt.Sprintf(
							"Time Summary Line #%v : "+
								"txtLabelRightFiller->"+
								"txtDescLabel", i+1)))

			if err != nil {
				return "", err
			}

			txtOutputLabel,
				err = TextFieldSpecLabel{}.NewPtrTextLabel(
				timeDurationStrs[i],
				-1,
				TxtJustify.Left(),
				ePrefix.XCpy(
					fmt.Sprintf(
						"Time Summary Line #%v : "+
							"txtLabelRightFiller->"+
							"txtDescLabel", i+1)))

			if err != nil {
				return "", err
			}

			_,
				err =
				stdLine.AddTextField(
					txtOutputLabel,
					ePrefix.XCpy(
						fmt.Sprintf(
							"Time Summary Line #%v : "+
								"txtOutputLabel->"+
								"txtDescLabel", i+1)))

			if err != nil {
				return "", err
			}

		} else {

			_,
				err =
				stdLine.AddTextField(
					txtFillerSumLeftMar,
					ePrefix.XCpy(
						fmt.Sprintf(
							"Time Summary Line #%v : "+
								"txtOutputLabel->"+
								"txtDescLabel", i+1)))

			if err != nil {
				return "", err
			}

			txtOutputLabel,
				err = TextFieldSpecLabel{}.NewPtrTextLabel(
				timeDurationStrs[i],
				-1,
				TxtJustify.Left(),
				ePrefix.XCpy(
					fmt.Sprintf(
						"Time Summary Line #%v : "+
							"timeDurationStrs[%v]->"+
							"txtOutputLabel", i+1, i)))

			if err != nil {
				return "", err
			}

			_,
				err = stdLine.AddTextField(
				txtOutputLabel,
				ePrefix.XCpy(
					fmt.Sprintf(
						"Time Summary Line #%v : "+
							"timeDurationStrs[%v]->"+
							"txtOutputLabel", i+1, i)))

			if err != nil {
				return "", err
			}

		} // End of 'else'

		_,
			err2 = sb.WriteString(stdLine.String())

		if err2 != nil {
			err = fmt.Errorf("%v\n"+
				"Error returned by summary line #%v .\n"+
				"sb.WriteString(stdLine.String())\n"+
				"%v\n",
				ePrefix.ZCtxEmpty().String(),
				i+1,
				err2.Error())

			return "", err
		}

	} // End of time duration strings for loop

	return sb.String(), nil
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
//     - A pointer to an instance of TextLineSpecTimerLines. All
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
//       If the string length of 'labelLeftMarginChars' plus
//       'labelRightMarginChars' plus the text label field length
//       ('textLabelFieldLen') exceeds the maximum length of
//       55-characters, this method will return an error.
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
//       If the string length of 'labelLeftMarginChars' plus
//       'labelRightMarginChars' plus the text label field length
//       ('textLabelFieldLen') exceeds the maximum length of
//       55-characters, this method will return an error.
//
//
//  endTime                    time.Time
//     - A time value which will be used in conjunction with
//       'startTime' parameter to compute the time duration or
//       elapsed for the timer event.
//
//       If this parameter is submitted as a zero time value,
//       'endTime' will be defaulted to value of 'startTime'.
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
//       If the string length of 'labelLeftMarginChars' plus
//       'labelRightMarginChars' plus the text label field length
//       ('textLabelFieldLen') exceeds the maximum length of
//       55-characters, this method will return an error.
//
//
//  textLabelFieldLen              int
//     - A user entered value which defines the length of the text
//       field used by all three text labels, 'startTimeLabel',
//       'endTimeLabel' and 'timeDurationLabel'.
//
//       The length of the text fields which will be used to
//       position and display the three text labels provided by
//       input parameters 'startTimeLabel', 'endTimeLabel' and
//       'timeDurationLabel'.
//
//       If 'textLabelFieldLen' is less than the length of the
//       longest text label it will be defaulted to the length
//       of the longest text label ('startTimeLabel',
//       'endTimeLabel' or 'timeDurationLabel').
//
//       If the string length of 'labelLeftMarginChars' plus
//       'labelRightMarginChars' plus the text label field length
//       ('textLabelFieldLen') exceeds the maximum length of
//       55-characters, this method will return an error.
//
//
//  textLabelJustification         TextJustify
//     - An enumeration which specifies the justification of the
//       three text labels 'startTimeLabel', 'endTimeLabel' and
//       'timeDurationLabel' within the field length specified by
//       'textLabelFieldLen'.
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
//  labelRightMarginChars      []rune
//     - This rune array contains the character or characters which
//       will be used to separate the text labels ('startTimeLabel',
//       'endTimeLabel' and 'timeDurationLabel') from the output or
//       data values displayed on the same line.
//
//       Example:
//        Start Time[sep chars]2010-01-02 15:04:05.000000000 -0700 MST
//
//       Often this parameter is set to a single white space
//       character (' ') or a colon plus white space character,
//       ([]rune{':', ' '}).
//
//
//       If this array is submitted as a zero length rune array,
//       'labelRightMarginChars' will be assigned a default
//       value of []rune{':', ' '}.
//       Example Output:
//         Start Time: 2010-01-02 15:04:05.000000000 -0700 MST
//
//       If the string length of 'labelLeftMarginChars' plus
//       'labelRightMarginChars' plus the text label field length
//       ('textLabelFieldLen') exceeds the maximum length of
//       55-characters, this method will return an error.
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
	labelLeftMarginChars []rune,
	startTimeLabel []rune,
	startTime time.Time,
	endTimeLabel []rune,
	endTime time.Time,
	timeFormat string,
	timeDurationLabel []rune,
	textLabelFieldLen int,
	labelJustification TextJustify,
	labelRightMarginChars []rune,
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
		"textLineSpecTimerLinesMolecule."+
			"setTxtLineSpecTimerLines()",
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

	if textLabelFieldLen > 1000000 {
		err = fmt.Errorf("%v\n"+
			"Error: 'textLabelFieldLen' is invalid!\n"+
			"'textLabelFieldLen' is greater than one million "+
			"(1,000,000)\n",
			ePrefix.String())

		return err
	}

	if textLabelFieldLen < -1 {
		err = fmt.Errorf("%v\n"+
			"Error: 'textLabelFieldLen' is invalid!\n"+
			"'textLabelFieldLen' is has a value "+
			"less than minus one (-1).\n",
			ePrefix.String())

		return err
	}

	if len(timeFormat) == 0 {
		timeFormat =
			textSpecificationMolecule{}.ptr().
				getDefaultTimeFormat()
	}

	txtTimerLinesElectron := textLineSpecTimerLinesElectron{}

	if startTime.IsZero() {
		startTime =
			txtTimerLinesElectron.
				getDefaultTime()
	}

	if endTime.IsZero() {
		endTime = startTime
	}

	if endTime.Before(startTime) {
		err = fmt.Errorf("%v\n"+
			"Error: 'endTime' occurs before 'startTime!\n",
			ePrefix.String())

		return err
	}

	labelLen := len(startTimeLabel)

	if len(startTimeLabel) == 0 {
		startTimeLabel =
			txtTimerLinesElectron.getDefaultStartTimeLabel()
	}

	labelLen = len(endTimeLabel)

	if labelLen == 0 {
		endTimeLabel =
			txtTimerLinesElectron.getDefaultEndTimeLabel()
	}

	labelLen = len(timeDurationLabel)

	if len(timeDurationLabel) == 0 {
		timeDurationLabel =
			txtTimerLinesElectron.getDefaultTimeDurationLabel()

	}

	labelLen = len(labelRightMarginChars)

	if len(labelRightMarginChars) == 0 {
		labelRightMarginChars =
			txtTimerLinesElectron.
				getDefaultLabelRightMarginChars()

	}

	lenLongestLabel := txtTimerLinesElectron.getLengthOfLongestLabel(
		txtTimerLines.startTimeLabel,
		txtTimerLines.endTimeLabel,
		txtTimerLines.timeDurationLabel)

	totalLabelLen := txtTimerLinesElectron.
		getTotalLabelLength(
			labelLeftMarginChars,
			startTimeLabel,
			endTimeLabel,
			timeDurationLabel,
			textLabelFieldLen,
			labelRightMarginChars)

	maxAllowableLabelLen := textLineSpecTimerLinesPreon{}.ptr().
		getMaximumTimerLabelLen()

	if totalLabelLen > maxAllowableLabelLen {
		err = fmt.Errorf("%v\n"+
			"Error: The total length of the text label field is invalid!\n"+
			"The maximum text label field length is %v-characters\n"+
			"The total length of 'labelLeftMarginChars' plus 'labelRightMarginChars'"+
			"plus the the text label field length is %v-characters."+
			"'text label field length' is computed by taking the longest"+
			"of the longest text label or the user entered text field length.\n"+
			"labelLeftMarginChars  = '%v'\n"+
			"startTimeLabel        = '%v'\n"+
			"endTimeLabel          = '%v'\n"+
			"timeDurationLabel     = '%v'\n"+
			"labelRightMarginChars = '%v'\n"+
			"textLabelFieldLen     = '%v'\n",
			ePrefix.String(),
			maxAllowableLabelLen,
			totalLabelLen,
			len(txtTimerLines.labelLeftMarginChars),
			len(txtTimerLines.startTimeLabel),
			len(txtTimerLines.endTimeLabel),
			len(txtTimerLines.timeDurationLabel),
			len(txtTimerLines.labelRightMarginChars),
			txtTimerLines.textLabelFieldLen)

		return err
	}

	if textLabelFieldLen < lenLongestLabel {
		textLabelFieldLen = lenLongestLabel
	}

	if !labelJustification.XIsValid() {
		err = fmt.Errorf("%v\n"+
			"Error: 'textLabelJustification' is invalid!\n"+
			"'textLabelJustification' should be 'Left',\n"+
			"'Right' or 'Center'.\n"+
			"'textLabelJustification' string value  = '%v'\n"+
			"'textLabelJustification' integer value = '%v'\n",
			ePrefix.String(),
			labelJustification.String(),
			labelJustification.XValueInt())

		return err
	}

	sMechPreon := strMechPreon{}

	err = sMechPreon.copyRuneArrays(
		&txtTimerLines.labelLeftMarginChars,
		&labelLeftMarginChars,
		true,
		ePrefix.XCpy("labelLeftMarginChars->"+
			"txtTimerLines.labelLeftMarginChars,"))

	if err != nil {
		return err
	}

	err = sMechPreon.copyRuneArrays(
		&txtTimerLines.startTimeLabel,
		&startTimeLabel,
		true,
		ePrefix.XCpy("startTimeLabel->"+
			"txtTimerLines.startTimeLabel"))

	if err != nil {
		return err
	}

	txtTimerLines.startTime = startTime

	err = sMechPreon.copyRuneArrays(
		&txtTimerLines.endTimeLabel,
		&endTimeLabel,
		true,
		ePrefix.XCpy("endTimeLabel"))

	if err != nil {
		return err
	}

	txtTimerLines.endTime = endTime

	txtTimerLines.timeFormat = timeFormat

	err = sMechPreon.copyRuneArrays(
		&txtTimerLines.timeDurationLabel,
		&timeDurationLabel,
		true,
		ePrefix.XCpy("timeDurationLabel"))

	if err != nil {
		return err
	}

	txtTimerLines.textLabelFieldLen =
		textLabelFieldLen

	txtTimerLines.textLabelJustification =
		labelJustification

	err = sMechPreon.copyRuneArrays(
		&txtTimerLines.labelRightMarginChars,
		&labelRightMarginChars,
		true,
		ePrefix.XCpy(
			"labelRightMarginChars"))

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
