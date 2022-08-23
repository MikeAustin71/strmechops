package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
	"time"
)

// TextLineTimerStartStopDto - This Text Line Timer Start Stop Data
// Transfer Object is used to transmit data parameters and
// specifications required to construct a timer event string.
//
// Text Line Timer Start Stop type records, computes and formats an
// elapsed time. This format process requires user input specifying
// a start time and ending time.
//
// The final formatted output string is composed of four lines of
// text for output to screen display, file output or printing.
//
// The first line of text shows the Starting Time. The second line
// shows the Ending Time. The third line displays the time duration
// or the difference between starting time and ending time. The
// fourth line displays the total elapsed time in nanoseconds.
//
// The third line contains Time duration, or elapsed time, and is
// broken down by days, hours, minutes, seconds, microseconds,
// milliseconds and nanoseconds. The display has a variable line
// length and will begin with the first category containing valid
// time duration data.
//
// Sample Output
//
//	  Start Time: 2021-08-13 03:19:32.462108100 -0500 CDT
//	    End Time: 2021-08-13 03:19:32.462163100 -0500 CDT
//	Elapsed Time: 55 Microseconds 0 Nanoseconds
//	              Total Elapsed Nanoseconds: 55,000
type TextLineTimerStartStopDto struct {
	FormatType TextFieldType
	// Required. This enumeration value specifies the type of Text
	// Format Operation to be performed.
	//
	// For TextLineTimerStartStopDto Format Type, this value
	// should always be set to:
	//   TxtFieldType.TimerStartStop()

	LeftMarginStr string
	// The contents of the string will be used as the left margin
	// for the timer event lines.
	//
	// If no left margin is required, set 'LeftMarginStr' to a zero
	// length or empty string, and no left margin will be created.

	StartTimeLabel string
	// A string containing the text characters constituting the
	// starting time text label.
	//
	// If this parameter is submitted as a zero length or empty
	// rune array, or if the rune array contains invalid zero
	// rune values, 'StartTimeLabel' will be assigned a default
	// value of "Start Time".
	//
	// If the length in characters of 'LeftMarginStr' plus
	// 'LeftMarginStr' plus the text label field length
	// ('TextLabelFieldLength') exceeds the maximum length
	// of 55-characters, this method will return an error.

	StartTime time.Time
	// A time value which will be used in conjunction with
	// the 'EndTime' parameter to compute the time duration
	// or elapsed time for the timer event.
	//
	// If this parameter is submitted as a zero time value,
	// 'startTime' will be defaulted to value of July 4, 1776
	// 9:30AM UTC.

	EndTimeLabel string
	// A string containing the text characters constituting the
	// ending time text label.
	//
	// If this parameter is submitted as a zero length or empty
	// string, 'EndTimeLabel' will be assigned a default
	// value of "End Time".
	//
	// If the length in characters of 'LeftMarginStr' plus
	// 'RightMarginStr' plus the text label field length
	// ('TextLabelFieldLength') exceeds the maximum length
	// of 55-characters, this method will return an error.

	EndTime time.Time
	// A time value which will be used in conjunction with
	// the 'StartTime' parameter to compute the time duration
	// or elapsed time for the timer event.
	//
	// If this parameter is submitted as a zero time value,
	// 'EndTime' will be defaulted to value of 'StartTime'.

	TimeFormat string
	// This string holds the time format parameters used to
	// format starting time and ending time values for text
	// output.
	//
	// If this parameter is submitted as an empty string,
	// parameter 'TimeFormat' will be assigned a default
	// value of:
	//  "2006-01-02 15:04:05.000000000 -0700 MST"

	TimeDurationLabel string
	// The text label used to describe the time duration or
	// elapsed time computed from the 'startTime' and 'endTime'
	// parameters.
	//
	// If this string is submitted as a zero length or empty
	// string, 'TimeDurationLabel' will be assigned a default
	// value of "Elapsed Time".
	//
	// If the string length of 'LeftMarginStr' plus
	// 'RightMarginStr' plus the text label field length
	// ('TextLabelFieldLength') exceeds the maximum length
	// of 55-characters, this method will return an error.

	TextLabelFieldLength int
	// A user entered value which defines the length of the
	// text field used by all three text labels,
	// 'StartTimeLabel', 'EndTimeLabel' and
	// 'TimeDurationLabel'.
	//
	// This text length value will be used to position and
	// display the three text labels provided by
	// input parameters 'StartTimeLabel', 'EndTimeLabel' and
	// 'TimeDurationLabel'.
	//
	// If 'TextLabelFieldLength' is less than the length of
	// the longest text label it will be defaulted to the
	// length of the longest text label ('StartTimeLabel',
	// 'EndTimeLabel' or 'TimeDurationLabel').
	//
	// If the string length of 'LeftMarginStr' plus
	// 'RightMarginStr' plus the text label field length
	// ('TextLabelFieldLength') exceeds the maximum length
	// of 55-characters, this method will return an error.

	TextLabelJustification TextJustify
	// An enumeration which specifies the text justification of
	// the three text labels 'StartTimeLabel', 'EndTimeLabel'
	// and 'TimeDurationLabel' within the field length specified
	// by 'TextLabelFieldLength'.
	//
	// Label justification must be equal to one of these three
	// valid values:
	//   TextJustify(0).Left()
	//   TextJustify(0).Right()
	//   TextJustify(0).Center()
	//
	// The abbreviated text justification enumeration syntax can
	// also be used:
	//
	//   TxtJustify.Left()
	//   TxtJustify.Right()
	//   TxtJustify.Center()

	RightMarginStr string
	// This string contains the character or characters which
	// will be used to separate the text labels ('startTimeLabel',
	// 'endTimeLabel' and 'timeDurationLabel') from the output or
	// data values displayed on the same line.
	//  Example:
	//  Start Time[RightMarginStr]2010-01-02 15:04:05.000000000 -0700 MST
	//
	// Often this parameter is set to a single white space
	// character (" ") or a colon plus white space character,
	// (": ").
	//
	// If this string is submitted as a zero length or empty
	// string, 'RightMarginStr' will be assigned a default value
	// of  ": ".
	//  Example Output:
	//  Start Time: 2010-01-02 15:04:05.000000000 -0700 MST
	//
	// If the string length of 'LeftMarginStr' plus
	// 'RightMarginStr' plus the text label field length
	// ('TextLabelFieldLength') exceeds the maximum length
	// of 55-characters, this method will return an error.

	lock *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming instance of
// TextFormatterDto ('incomingTxtFmtDto') to the data fields
// of the current TextFormatterDto instance
// ('txtFmtDto').
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// All the data fields in current TextFormatterDto instance
// ('txtFmtDto') will be deleted and overwritten.
//
// NO DATA VALIDATION IS performed on input parameter,
// 'incomingTxtFmtDto'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	incomingTimerStartStopDto  TextLineTimerStartStopDto
//	   - An instance of TextLineTimerStartStopDto. This method will
//	     NOT change the data values of member variables contained
//	     in this instance.
//
//	     All data values in this TextLineTimerStartStopDto instance
//	     ('incomingTimerStartStopDto') will be copied to the
//	     current TextLineTimerStartStopDto instance
//	     ('lineTimerStartStopDto').
//
//	     No data validation is performed on input parameter,
//	     'incomingTimerStartStopDto'.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	NONE
func (lineTimerStartStopDto *TextLineTimerStartStopDto) CopyIn(
	incomingTimerStartStopDto TextLineTimerStartStopDto) {

	if lineTimerStartStopDto.lock == nil {
		lineTimerStartStopDto.lock = new(sync.Mutex)
	}

	lineTimerStartStopDto.lock.Lock()

	defer lineTimerStartStopDto.lock.Unlock()

	_ = new(textLineTimerStartStopDtoNanobot).
		copyData(
			lineTimerStartStopDto,
			&incomingTimerStartStopDto,
			nil)

	return
}

// CopyOut - Returns a deep copy of the current
// TextLineTimerStartStopDto instance.
//
// NO DATA VALIDATION is performed on the current instance of
// TextLineTimerStartStopDto.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	NONE
//
// ----------------------------------------------------------------
//
// Return Values
//
//	deepCopyTimerStartStopDto     TextLineTimerStartStopDto
//	   - This parameter will return a deep copy of the current
//	     TextLineTimerStartStopDto instance.
func (lineTimerStartStopDto *TextLineTimerStartStopDto) CopyOut() (
	deepCopyTimerStartStopDto TextLineTimerStartStopDto) {

	if lineTimerStartStopDto.lock == nil {
		lineTimerStartStopDto.lock = new(sync.Mutex)
	}

	lineTimerStartStopDto.lock.Lock()

	defer lineTimerStartStopDto.lock.Unlock()

	_ = new(textLineTimerStartStopDtoNanobot).
		copyData(
			&deepCopyTimerStartStopDto,
			lineTimerStartStopDto,
			nil)

	return deepCopyTimerStartStopDto
}

// Empty - Resets all internal member variables for the current
// instance of TextLineTimerStartStopDto to their zero or
// uninitialized states. This method will leave the current
// instance of TextLineTimerStartStopDto in an invalid state and
// unavailable for immediate reuse.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete all member variable data values in this
// current instance of TextLineTimerStartStopDto. All member
// variable data values will be reset to their zero or
// uninitialized states.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	NONE
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	NONE
func (lineTimerStartStopDto *TextLineTimerStartStopDto) Empty() {

	if lineTimerStartStopDto.lock == nil {
		lineTimerStartStopDto.lock = new(sync.Mutex)
	}

	lineTimerStartStopDto.lock.Lock()

	new(textLineTimerStartStopDtoAtom).
		empty(
			lineTimerStartStopDto)

	lineTimerStartStopDto.lock.Unlock()

	lineTimerStartStopDto.lock = nil

}

// Equal - Receives another instance of TextLineTimerStartStopDto
// and proceeds to compare the member variables to those of the
// current TextLineTimerStartStopDto instance in order to
// determine if they are equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables of both instances are equal in
// all respects, this flag is set to 'true'. Otherwise, this method
// returns 'false'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	incomingTimerStartStopDto     TextLineTimerStartStopDto
//	   - An incoming instance of TextLineTimerStartStopDto. This
//	     method will compare all member variable data values in
//	     this instance against those contained in the current
//	     instance of TextLineTimerStartStopDto. If the data values
//	     in both instances are found to be equal in all respects,
//	     this method will return a boolean value of 'true'.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	bool
//	   - If the member variable data values contained in input
//	     parameter 'incomingTimerStartStopDto' are equal in all
//	     respects to those contained in the current instance of
//	     TextLineTimerStartStopDto, this method will return a
//	     boolean value of 'true'. Otherwise, a value of 'false'
//	     will be returned to the calling function.
func (lineTimerStartStopDto *TextLineTimerStartStopDto) Equal(
	incomingTimerStartStopDto TextLineTimerStartStopDto) bool {

	if lineTimerStartStopDto.lock == nil {
		lineTimerStartStopDto.lock = new(sync.Mutex)
	}

	lineTimerStartStopDto.lock.Lock()

	defer lineTimerStartStopDto.lock.Unlock()

	return new(textLineTimerStartStopDtoAtom).
		equal(
			lineTimerStartStopDto,
			&incomingTimerStartStopDto)
}

// textLineTimerStartStopDtoNanobot - Provides helper methods for
// TextLineTimerStartStopDto
type textLineTimerStartStopDtoNanobot struct {
	lock *sync.Mutex
}

// copyDta - Copies all data from a source instance of
// TextLineTimerStartStopDto to a destination instance of
// TextLineTimerStartStopDto.
func (timerStartStopDtoNanobot *textLineTimerStartStopDtoNanobot) copyData(
	destinationTimerDto *TextLineTimerStartStopDto,
	sourceTimerDto *TextLineTimerStartStopDto,
	errPrefDto *ePref.ErrPrefixDto) error {

	if timerStartStopDtoNanobot.lock == nil {
		timerStartStopDtoNanobot.lock = new(sync.Mutex)
	}

	timerStartStopDtoNanobot.lock.Lock()

	defer timerStartStopDtoNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineTimerStartStopDtoNanobot."+
			"copyData()",
		"")

	if err != nil {

		return err

	}

	if sourceTimerDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'sourceTimerDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if destinationTimerDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'destinationTimerDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	new(textLineTimerStartStopDtoAtom).
		empty(destinationTimerDto)

	destinationTimerDto.FormatType =
		sourceTimerDto.FormatType

	destinationTimerDto.LeftMarginStr =
		sourceTimerDto.LeftMarginStr

	destinationTimerDto.StartTimeLabel =
		sourceTimerDto.StartTimeLabel

	destinationTimerDto.StartTime =
		sourceTimerDto.StartTime

	destinationTimerDto.EndTimeLabel =
		sourceTimerDto.EndTimeLabel

	destinationTimerDto.EndTime =
		sourceTimerDto.EndTime

	destinationTimerDto.TimeFormat =
		sourceTimerDto.TimeFormat

	destinationTimerDto.TimeDurationLabel =
		sourceTimerDto.TimeDurationLabel

	destinationTimerDto.TextLabelFieldLength =
		sourceTimerDto.TextLabelFieldLength

	destinationTimerDto.TextLabelJustification =
		sourceTimerDto.TextLabelJustification

	destinationTimerDto.RightMarginStr =
		sourceTimerDto.RightMarginStr

	return err
}

// textLineTimerStartStopDtoAtom - Provides helper methods for
// TextLineTimerStartStopDto
type textLineTimerStartStopDtoAtom struct {
	lock *sync.Mutex
}

// empty - Receives a pointer to an instance of
// TextLineTimerStartStopDto and proceeds to set all the
// internal member variables to their zero or uninitialized
// states.
//
// This method will therefore delete all data currently held
// by this instance of TextLineTimerStartStopDto.
func (timerStartStopDtoAtom *textLineTimerStartStopDtoAtom) empty(
	txtTimerStartStopDto *TextLineTimerStartStopDto) {

	if timerStartStopDtoAtom.lock == nil {
		timerStartStopDtoAtom.lock = new(sync.Mutex)
	}

	timerStartStopDtoAtom.lock.Lock()

	defer timerStartStopDtoAtom.lock.Unlock()

	txtTimerStartStopDto.FormatType = TxtFieldType.None()

	txtTimerStartStopDto.LeftMarginStr = ""

	txtTimerStartStopDto.StartTimeLabel = ""

	txtTimerStartStopDto.StartTime = time.Time{}

	txtTimerStartStopDto.EndTimeLabel = ""

	txtTimerStartStopDto.EndTime = time.Time{}

	txtTimerStartStopDto.TimeFormat = ""

	txtTimerStartStopDto.TimeDurationLabel = ""

	txtTimerStartStopDto.TextLabelFieldLength = -99

	txtTimerStartStopDto.TextLabelJustification = TxtJustify.None()

	txtTimerStartStopDto.RightMarginStr = ""

	return
}

// equal - Receives pointers to two instances of
// TextLineTimerStartStopDto and proceeds to compare all the member
// data variables for both instances.
//
// If the two instances of TextLineTimerStartStopDto are found to
// be equal in all respects, this method will return a boolean
// value of 'true'.
func (timerStartStopDtoAtom *textLineTimerStartStopDtoAtom) equal(
	timerStartStopDto1 *TextLineTimerStartStopDto,
	timerStartStopDto2 *TextLineTimerStartStopDto) bool {

	if timerStartStopDtoAtom.lock == nil {
		timerStartStopDtoAtom.lock = new(sync.Mutex)
	}

	timerStartStopDtoAtom.lock.Lock()

	defer timerStartStopDtoAtom.lock.Unlock()

	if timerStartStopDto1 == nil ||
		timerStartStopDto2 == nil {

		return false
	}

	if timerStartStopDto1.FormatType !=
		timerStartStopDto2.FormatType {

		return false
	}

	if timerStartStopDto1.LeftMarginStr !=
		timerStartStopDto2.LeftMarginStr {

		return false
	}

	if timerStartStopDto1.StartTimeLabel !=
		timerStartStopDto2.StartTimeLabel {

		return false
	}

	if timerStartStopDto1.StartTime !=
		timerStartStopDto2.StartTime {

		return false
	}

	if timerStartStopDto1.EndTimeLabel !=
		timerStartStopDto2.EndTimeLabel {

		return false
	}

	if timerStartStopDto1.EndTime !=
		timerStartStopDto2.EndTime {

		return false
	}

	if timerStartStopDto1.TimeFormat !=
		timerStartStopDto2.TimeFormat {

		return false
	}

	if timerStartStopDto1.TimeDurationLabel !=
		timerStartStopDto2.TimeDurationLabel {

		return false
	}

	if timerStartStopDto1.TextLabelFieldLength !=
		timerStartStopDto2.TextLabelFieldLength {

		return false
	}

	if timerStartStopDto1.TextLabelJustification !=
		timerStartStopDto2.TextLabelJustification {

		return false
	}

	if timerStartStopDto1.RightMarginStr !=
		timerStartStopDto2.RightMarginStr {

		return false
	}

	return true
}
