package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strconv"
	"sync"
	"time"
)

type textLineSpecTimerLinesElectron struct {
	lock *sync.Mutex
}

// computeTimeDuration - Computes the time duration from starting
// time and ending time. The computed time duration or elapsed time
// is returned in a string format.
//
// The returned time duration is formatted in days, hours, minutes,
// seconds, milliseconds, microseconds, nanoseconds and
// total nanoseconds. The output text display of these values begins
// with the first category that has the first non-zero value.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  startTime                  time.Time
//     - The starting time for the time duration event.
//
//  endTime                    time.Time
//     - The ending time for the time duration event.
//
//  summaryTextLineLeftMargin  int
//     - The left margin to be used in outputting the summary text
//       line strings for text display or printing. The length of
//       the returned time duration strings ('timeDurationStrs')
//       will be computed and adjusted using this left margin value
//       so that no time duration result string will exceed a
//       length of 78-characters.
//
//       If 'summaryTextLineLeftMargin' is less than zero (0) or
//       greater than fifty-five (55) characters, this method will
//       return an error.
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
//  timeDurationStrs           []string
//     - An array of text strings presenting the results of the time
//       duration computation.
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
func (txtTimerLinesElectron *textLineSpecTimerLinesElectron) computeTimeDuration(
	startTime time.Time,
	endTime time.Time,
	summaryTextLineLeftMargin int,
	errPrefDto *ePref.ErrPrefixDto) (
	timeDurationStrs []string,
	err error) {

	if txtTimerLinesElectron.lock == nil {
		txtTimerLinesElectron.lock = new(sync.Mutex)
	}

	txtTimerLinesElectron.lock.Lock()

	defer txtTimerLinesElectron.lock.Unlock()

	const maxSummaryTextLeftLineMargin = 55

	var ePrefix *ePref.ErrPrefixDto

	timeDurationStrs = nil

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecTimerLinesElectron."+
			"computeTimeDuration()",
		"")

	if err != nil {
		return timeDurationStrs, err
	}

	if startTime.IsZero() {
		err = fmt.Errorf("%v\n"+
			"Error: Input time parameter 'startTime' has a zero value!\n",
			ePrefix.String())

		return timeDurationStrs, err
	}

	if endTime.IsZero() {
		err = fmt.Errorf("%v\n"+
			"Error: Input time parameter 'endTime' has a zero value!\n",
			ePrefix.String())

		return timeDurationStrs, err
	}

	if endTime.Before(startTime) {

		timeFormat :=
			textSpecificationMolecule{}.ptr().
				getDefaultTimeFormat()

		err = fmt.Errorf("%v\n"+
			"Error: Input time parameters 'startTime' and "+
			"'endTime' are invalid!\n"+
			"'endTime' occurs before 'startTime'.\n"+
			"'startTime' = '%v'\n"+
			"  'endTime' = '%v'\n",
			ePrefix.String(),
			startTime.Format(timeFormat),
			endTime.Format(timeFormat))

		return timeDurationStrs, err
	}

	if summaryTextLineLeftMargin < 0 ||
		summaryTextLineLeftMargin > maxSummaryTextLeftLineMargin {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'summaryTextLineLeftMargin' is invalid!\n"+
			"The valid range for 'summaryTextLineLeftMargin' is 0-%v,"+
			"inclusive.\n"+
			"summaryTextLineLeftMargin= '%v'\n",
			ePrefix.String(),
			maxSummaryTextLeftLineMargin,
			summaryTextLineLeftMargin)

		return timeDurationStrs, err
	}

	// MicroSecondNanoseconds - Number of Nanoseconds in a Microsecond.
	// 	A MicroSecond is 1/1,000,000 or 1 one-millionth of a second
	MicroSecondNanoseconds := int64(time.Microsecond)

	// MilliSecondNanoseconds - Number of Nanoseconds in a MilliSecond.
	//	 A millisecond is 1/1,000 or 1 one-thousandth of a second
	MilliSecondNanoseconds := int64(time.Millisecond)

	// SecondNanoseconds - Number of Nanoseconds in a Second
	SecondNanoseconds := int64(time.Second)

	// MinuteNanoseconds - Number of Nanoseconds in a minute
	MinuteNanoseconds := int64(time.Minute)

	// HourNanoseconds - Number of Nanoseconds in an hour
	HourNanoseconds := int64(time.Hour)

	// DayNanoseconds - Number of Nanoseconds in a standard
	// 24-hour day.
	DayNanoseconds := int64(time.Hour) * int64(24)

	t2Dur := endTime.Sub(startTime)

	totalNanoseconds := t2Dur.Nanoseconds()
	summaryNanoseconds := totalNanoseconds

	numOfDays := int64(0)
	numOfHours := int64(0)
	numOfMinutes := int64(0)
	numOfSeconds := int64(0)
	numOfMilliseconds := int64(0)
	numOfMicroseconds := int64(0)
	numOfNanoseconds := int64(0)

	if totalNanoseconds >= DayNanoseconds {
		numOfDays = totalNanoseconds / DayNanoseconds
		totalNanoseconds =
			totalNanoseconds - (numOfDays * DayNanoseconds)
	}

	if totalNanoseconds >= HourNanoseconds {
		numOfHours = totalNanoseconds / HourNanoseconds
		totalNanoseconds = totalNanoseconds - (numOfHours * HourNanoseconds)
	}

	if totalNanoseconds >= MinuteNanoseconds {
		numOfMinutes = totalNanoseconds / MinuteNanoseconds
		totalNanoseconds = totalNanoseconds - (numOfMinutes * MinuteNanoseconds)
	}

	if totalNanoseconds >= SecondNanoseconds {
		numOfSeconds = totalNanoseconds / SecondNanoseconds
		totalNanoseconds = totalNanoseconds - (numOfSeconds * SecondNanoseconds)
	}

	if totalNanoseconds >= SecondNanoseconds {
		numOfSeconds = totalNanoseconds / SecondNanoseconds
		totalNanoseconds = totalNanoseconds - (numOfSeconds * SecondNanoseconds)
	}

	if totalNanoseconds >= MilliSecondNanoseconds {
		numOfMilliseconds = totalNanoseconds / MilliSecondNanoseconds
		totalNanoseconds = totalNanoseconds - (numOfMilliseconds * MilliSecondNanoseconds)
	}

	if totalNanoseconds >= MicroSecondNanoseconds {
		numOfMicroseconds = totalNanoseconds / MicroSecondNanoseconds
		totalNanoseconds = totalNanoseconds - (numOfMicroseconds * MicroSecondNanoseconds)
	}

	numOfNanoseconds = totalNanoseconds

	var nStrIntSeparator IntegerSeparatorDto

	nStrIntSeparator,
		err = IntegerSeparatorDto{}.
		NewUnitedStatesDefaults(
			ePrefix.XCtxEmpty())

	if err != nil {
		return timeDurationStrs, err
	}

	nStrIntSepMolecule := integerSeparatorDtoMolecule{}

	outputLine := ""
	outputValStr := ""
	var numStrWithIntSeps []rune
	var foundFirstValue = false
	maxSummaryLineLen := 78 - summaryTextLineLeftMargin
	newOutputLine := ""

	// numOfDays
	if numOfDays > 0 {

		outputValStr =
			strconv.FormatInt(numOfDays, 10)

		numStrWithIntSeps,
			err =
			nStrIntSepMolecule.applyIntSeparators(
				&nStrIntSeparator,
				[]rune(outputValStr),
				ePrefix.XCtx("numOfDays"))

		if err != nil {
			return timeDurationStrs, err
		}

		newOutputLine = fmt.Sprintf(
			"%v Days ",
			string(numStrWithIntSeps))

		outputLine += newOutputLine

		foundFirstValue = true
	}

	// numOfHours
	if numOfHours > 0 || foundFirstValue {

		outputValStr =
			strconv.FormatInt(numOfHours, 10)

		numStrWithIntSeps,
			err =
			nStrIntSepMolecule.applyIntSeparators(
				&nStrIntSeparator,
				[]rune(outputValStr),
				ePrefix.XCtx("numOfHours"))

		if err != nil {
			return timeDurationStrs, err
		}

		newOutputLine = fmt.Sprintf(
			"%v Hours ",
			string(numStrWithIntSeps))

		if len(outputLine)+
			len(newOutputLine) >= maxSummaryLineLen {
			timeDurationStrs = append(
				timeDurationStrs, outputLine)
			outputLine = ""
		}

		outputLine += newOutputLine

		foundFirstValue = true

	}

	// numOfMinutes
	if numOfMinutes > 0 || foundFirstValue {

		outputValStr =
			strconv.FormatInt(numOfMinutes, 10)

		numStrWithIntSeps,
			err =
			nStrIntSepMolecule.applyIntSeparators(
				&nStrIntSeparator,
				[]rune(outputValStr),
				ePrefix.XCtx("numOfMinutes"))

		if err != nil {
			return timeDurationStrs, err
		}

		newOutputLine =
			fmt.Sprintf("%v Minutes ",
				string(numStrWithIntSeps))

		if len(outputLine)+
			len(newOutputLine) >= maxSummaryLineLen {
			timeDurationStrs = append(
				timeDurationStrs, outputLine)
			outputLine = ""
		}

		outputLine += newOutputLine

		foundFirstValue = true

	}

	// numOfSeconds
	if numOfSeconds > 0 || foundFirstValue {

		outputValStr =
			strconv.FormatInt(numOfSeconds, 10)

		numStrWithIntSeps,
			err =
			nStrIntSepMolecule.applyIntSeparators(
				&nStrIntSeparator,
				[]rune(outputValStr),
				ePrefix.XCtx("numOfSeconds"))

		if err != nil {
			return timeDurationStrs, err
		}

		newOutputLine =
			fmt.Sprintf("%v Seconds ",
				string(numStrWithIntSeps))

		if len(outputLine)+
			len(newOutputLine) >= maxSummaryLineLen {
			timeDurationStrs = append(
				timeDurationStrs, outputLine)
			outputLine = ""
		}

		outputLine += newOutputLine

		foundFirstValue = true

	}

	// numOfMilliseconds
	if numOfMilliseconds > 0 || foundFirstValue {

		outputValStr =
			strconv.FormatInt(numOfMilliseconds, 10)

		numStrWithIntSeps,
			err =
			nStrIntSepMolecule.applyIntSeparators(
				&nStrIntSeparator,
				[]rune(outputValStr),
				ePrefix.XCtx("numOfMilliseconds"))

		if err != nil {
			return timeDurationStrs, err
		}

		newOutputLine =
			fmt.Sprintf("%v Milliseconds ",
				string(numStrWithIntSeps))

		if len(outputLine)+
			len(newOutputLine) >= maxSummaryLineLen {
			timeDurationStrs = append(
				timeDurationStrs, outputLine)
			outputLine = ""
		}

		outputLine += newOutputLine

		foundFirstValue = true

	}

	// numOfMicroseconds
	if numOfMicroseconds > 0 || foundFirstValue {

		outputValStr =
			strconv.FormatInt(numOfMicroseconds, 10)

		numStrWithIntSeps,
			err =
			nStrIntSepMolecule.applyIntSeparators(
				&nStrIntSeparator,
				[]rune(outputValStr),
				ePrefix.XCtx("numOfMicroseconds"))

		if err != nil {
			return timeDurationStrs, err
		}

		newOutputLine =
			fmt.Sprintf("%v Microseconds ",
				string(numStrWithIntSeps))

		if len(outputLine)+
			len(newOutputLine) >= maxSummaryLineLen {
			timeDurationStrs = append(
				timeDurationStrs, outputLine)
			outputLine = ""
		}

		outputLine += newOutputLine

		foundFirstValue = true

	}

	// numOfNanoseconds is always presented

	outputValStr =
		strconv.FormatInt(numOfNanoseconds, 10)

	numStrWithIntSeps,
		err =
		nStrIntSepMolecule.applyIntSeparators(
			&nStrIntSeparator,
			[]rune(outputValStr),
			ePrefix.XCtx("numOfNanoseconds"))

	if err != nil {
		return timeDurationStrs, err
	}

	newOutputLine =
		fmt.Sprintf("%v Nanoseconds",
			string(numStrWithIntSeps))

	if len(outputLine)+
		len(newOutputLine) >= maxSummaryLineLen {
		timeDurationStrs = append(
			timeDurationStrs, outputLine)
		outputLine = ""
	}

	outputLine += newOutputLine

	// summaryNanoseconds is always presented

	// Always Place Summary Nanoseconds on a
	// separate line.
	timeDurationStrs = append(
		timeDurationStrs, outputLine)

	outputLine = ""

	outputValStr =
		strconv.FormatInt(summaryNanoseconds, 10)

	numStrWithIntSeps,
		err =
		nStrIntSepMolecule.applyIntSeparators(
			&nStrIntSeparator,
			[]rune(outputValStr),
			ePrefix.XCtx("summaryNanoseconds"))

	if err != nil {
		return timeDurationStrs, err
	}

	outputLine =
		fmt.Sprintf("Total Elapsed Nanoseconds: %v",
			string(numStrWithIntSeps))

	timeDurationStrs = append(
		timeDurationStrs, outputLine)

	return timeDurationStrs, err
}

// empty - Receives a pointer to an instance of
// TextLineSpecTimerLines and proceeds to set all the internal
// member variables to their uninitialized or zero states.
//
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// The data values of all member variables contained in input
// parameter 'txtTimerLines' will be overwritten and deleted.
//
func (txtTimerLinesElectron *textLineSpecTimerLinesElectron) empty(
	txtTimerLines *TextLineSpecTimerLines) {

	if txtTimerLinesElectron.lock == nil {
		txtTimerLinesElectron.lock = new(sync.Mutex)
	}

	txtTimerLinesElectron.lock.Lock()

	defer txtTimerLinesElectron.lock.Unlock()

	if txtTimerLines == nil {
		return
	}

	txtTimerLines.startTimeLabel = nil

	txtTimerLines.startTime = time.Time{}

	txtTimerLines.endTimeLabel = nil

	txtTimerLines.endTime = time.Time{}

	txtTimerLines.timeFormat = ""

	txtTimerLines.timeDurationLabel = nil

	txtTimerLines.textLabelFieldLen = 0

	txtTimerLines.textLabelJustification = TxtJustify.None()

	txtTimerLines.labelRightMarginChars = nil

	txtTimerLines.textLineReader = nil

	return
}

// getDefaultEndTimeLabel - Returns an array of runes containing
// the default End Time Label for instances of
// TextLineSpecTimerLines.
//
func (txtTimerLinesElectron *textLineSpecTimerLinesElectron) getDefaultEndTimeLabel() []rune {

	if txtTimerLinesElectron.lock == nil {
		txtTimerLinesElectron.lock = new(sync.Mutex)
	}

	txtTimerLinesElectron.lock.Lock()

	defer txtTimerLinesElectron.lock.Unlock()

	return []rune("End Time")
}

// getDefaultLabelRightMarginChars - Returns an array of
// runes containing the default Label Output Separation Characters
// Label for instances of TextLineSpecTimerLines.
//
func (txtTimerLinesElectron *textLineSpecTimerLinesElectron) getDefaultLabelRightMarginChars() []rune {

	if txtTimerLinesElectron.lock == nil {
		txtTimerLinesElectron.lock = new(sync.Mutex)
	}

	txtTimerLinesElectron.lock.Lock()

	defer txtTimerLinesElectron.lock.Unlock()

	return []rune{':', ' '}
}

// getDefaultStartTimeLabel - Returns an array of runes containing
// the default Start Time Label for instances of
// TextLineSpecTimerLines.
//
func (txtTimerLinesElectron *textLineSpecTimerLinesElectron) getDefaultStartTimeLabel() []rune {

	if txtTimerLinesElectron.lock == nil {
		txtTimerLinesElectron.lock = new(sync.Mutex)
	}

	txtTimerLinesElectron.lock.Lock()

	defer txtTimerLinesElectron.lock.Unlock()

	return []rune("Start Time")
}

// getDefaultTimeDurationLabel - Returns an array of runes containing
// the default Time Duration or Elapsed Time Label for instances of
// TextLineSpecTimerLines.
//
func (txtTimerLinesElectron *textLineSpecTimerLinesElectron) getDefaultTimeDurationLabel() []rune {

	if txtTimerLinesElectron.lock == nil {
		txtTimerLinesElectron.lock = new(sync.Mutex)
	}

	txtTimerLinesElectron.lock.Lock()

	defer txtTimerLinesElectron.lock.Unlock()

	return []rune("Elapsed Time")
}

// getDefaultTime - Returns default start and end time for
// TextLineSpecTimerLines objects.
//
// The default time is July 4, 1776 09:30AM UTC
//
func (txtTimerLinesElectron *textLineSpecTimerLinesElectron) getDefaultTime() time.Time {

	if txtTimerLinesElectron.lock == nil {
		txtTimerLinesElectron.lock = new(sync.Mutex)
	}

	txtTimerLinesElectron.lock.Lock()

	defer txtTimerLinesElectron.lock.Unlock()

	defaultTime := time.Date(
		1776,
		7,
		4,
		9,
		30,
		0,
		0,
		time.UTC)

	return defaultTime
}

// getLengthOfLongestLabel - Returns the length of the longest text
// label currently configured for a TextLineSpecTimerLines timer
// event.
//
// The length of the longest text label is determined by comparing
// the individual string lengths of the three text labels
// associated with every TextLineSpecTimerLines timer event.
// Namely, these are 'startTimeLabel', 'endTimeLabel', and
// 'timeDurationLabel'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  startTimeLabel             []rune
//     - An array of runes containing the text characters
//       constituting the starting time text label.
//
//
//  endTimeLabel               []rune
//     - An array of runes containing the text characters
//       constituting the ending time text label.
//
//
//  timeDurationLabel          []rune
//     - The text label used to describe the time duration or
//       elapsed time computed from the 'startTime' and 'endTime'
//       parameters.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  int
//     - An integer value identifying the length in runes of the
//       longest text label for an instance of
//       TextLineSpecTimerLines. All instances of
//       TextLineSpecTimerLines encapsulate three text labels,
//       'startTimeLabel', 'endTimeLabel', and 'timeDurationLabel'.
//       This return parameter holds the length of the longest
//       of these three text labels.
//
func (txtTimerLinesElectron *textLineSpecTimerLinesElectron) getLengthOfLongestLabel(
	startTimeLabel []rune,
	endTimeLabel []rune,
	timeDurationLabel []rune) int {

	if txtTimerLinesElectron.lock == nil {
		txtTimerLinesElectron.lock = new(sync.Mutex)
	}

	txtTimerLinesElectron.lock.Lock()

	defer txtTimerLinesElectron.lock.Unlock()

	maxLabelLen := len(startTimeLabel)

	itemLen := len(endTimeLabel)

	if itemLen > maxLabelLen {
		maxLabelLen = itemLen
	}

	itemLen = len(timeDurationLabel)

	if itemLen > maxLabelLen {
		maxLabelLen = itemLen
	}

	return maxLabelLen
}

// getTotalLabelLength - Returns the true length of a series of
// text labels plus the label's left and right margins.
//
// 'textLabelFieldLen' is a user entered value defining the length
// of all text label strings. If this value is greater than the
// length of the longest text label, it is used to determine the
// length of text label strings. Otherwise, the length of the
// longest text label is used.
//
func (txtTimerLinesElectron *textLineSpecTimerLinesElectron) getTotalLabelLength(
	labelLeftMarginChars []rune,
	startTimeLabel []rune,
	endTimeLabel []rune,
	timeDurationLabel []rune,
	textLabelFieldLen int,
	labelRightMarginChars []rune) int {

	if txtTimerLinesElectron.lock == nil {
		txtTimerLinesElectron.lock = new(sync.Mutex)
	}

	txtTimerLinesElectron.lock.Lock()

	defer txtTimerLinesElectron.lock.Unlock()

	itemLen := len(startTimeLabel)
	longestLabelLen := itemLen

	itemLen = len(endTimeLabel)

	if itemLen > longestLabelLen {
		longestLabelLen = itemLen
	}

	itemLen = len(timeDurationLabel)

	if itemLen > longestLabelLen {
		longestLabelLen = itemLen
	}

	if textLabelFieldLen > longestLabelLen {
		longestLabelLen = textLabelFieldLen
	}

	return textLabelFieldLen +
		len(labelLeftMarginChars) +
		len(labelRightMarginChars)
}

// ptr - Returns a pointer to a new instance of
// textLineSpecTimerLinesElectron.
//
func (txtTimerLinesElectron textLineSpecTimerLinesElectron) ptr() *textLineSpecTimerLinesElectron {

	if txtTimerLinesElectron.lock == nil {
		txtTimerLinesElectron.lock = new(sync.Mutex)
	}

	txtTimerLinesElectron.lock.Lock()

	defer txtTimerLinesElectron.lock.Unlock()

	return &textLineSpecTimerLinesElectron{
		lock: new(sync.Mutex),
	}
}
