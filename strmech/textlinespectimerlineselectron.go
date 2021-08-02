package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
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
func (txtTimerLinesElectron *textLineSpecTimerLinesElectron) computeTimeDuration(
	startTime time.Time,
	endTime time.Time,
	errPrefDto *ePref.ErrPrefixDto) (
	timeDurationStr string,
	err error) {

	if txtTimerLinesElectron.lock == nil {
		txtTimerLinesElectron.lock = new(sync.Mutex)
	}

	txtTimerLinesElectron.lock.Lock()

	defer txtTimerLinesElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	timeDurationStr = ""

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecTimerLinesElectron."+
			"computeTimeDuration()",
		"")

	if err != nil {
		return timeDurationStr, err
	}

	if startTime.IsZero() {
		err = fmt.Errorf("%v\n"+
			"Error: Input time parameter 'startTime' has a zero value!\n",
			ePrefix.String())

		return timeDurationStr, err
	}

	if endTime.IsZero() {
		err = fmt.Errorf("%v\n"+
			"Error: Input time parameter 'endTime' has a zero value!\n",
			ePrefix.String())

		return timeDurationStr, err
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

		return timeDurationStr, err
	}

	// MicroSecondNanoseconds - Number of Nanoseconds in a Microsecond
	// 	A MicroSecond is 1/1,000,000 or 1 one-millionth of a second
	MicroSecondNanoseconds := int64(time.Microsecond)

	// MilliSecondNanoseconds - Number of Nanoseconds in a MilliSecond
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

	if numOfDays > 0 {
		timeDurationStr += fmt.Sprintf(
			"%v-Days ", numOfDays)
	}

	if numOfHours > 0 || timeDurationStr != "" {

		timeDurationStr += fmt.Sprintf("%v-Hours ", numOfHours)

	}

	if numOfMinutes > 0 || timeDurationStr != "" {

		timeDurationStr += fmt.Sprintf("%v-Minutes ", numOfMinutes)

	}

	if numOfSeconds > 0 || timeDurationStr != "" {

		timeDurationStr += fmt.Sprintf("%v-Seconds ", numOfSeconds)

	}

	if numOfMilliseconds > 0 || timeDurationStr != "" {

		timeDurationStr += fmt.Sprintf("%v-Milliseconds ", numOfMilliseconds)

	}

	if numOfMicroseconds > 0 || timeDurationStr != "" {

		timeDurationStr += fmt.Sprintf("%v-Microseconds ", numOfMicroseconds)

	}

	timeDurationStr += fmt.Sprintf("%v-Nanoseconds", numOfNanoseconds)

	timeDurationStr += fmt.Sprintf(" - Total Elapsed Nanoseconds: %v", numOfNanoseconds)

	return timeDurationStr, err

}

// empty - Receives a pointer to an instance of
// TextLineSpecTimerLines and proceeds to set all of the internal
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

	txtTimerLines.labelFieldLen = 0

	txtTimerLines.labelJustification = TxtJustify.None()

	txtTimerLines.labelOutputSeparationChars = nil

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

// getDefaultLabelOutputSeparationCharsLabel - Returns an array of
// runes containing the default Label Output Separation Characters
// Label for instances of TextLineSpecTimerLines.
//
func (txtTimerLinesElectron *textLineSpecTimerLinesElectron) getDefaultLabelOutputSeparationCharsLabel() []rune {

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
