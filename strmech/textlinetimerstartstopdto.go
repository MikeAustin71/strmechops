package strmech

import (
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
//    Start Time: 2021-08-13 03:19:32.462108100 -0500 CDT
//      End Time: 2021-08-13 03:19:32.462163100 -0500 CDT
//  Elapsed Time: 55 Microseconds 0 Nanoseconds
//                Total Elapsed Nanoseconds: 55,000
//
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
	//  Start Time[right margin chars]2010-01-02 15:04:05.000000000 -0700 MST
	//
	// Often this parameter is set to a single white space
	// character (" ") or a colon plus white space character,
	// (": ").
	//
	// If this string is submitted as a zero length or empty
	// string, 'LabelRightMarginStr' will be assigned a
	// default value of  ": ".
	//  Example Output:
	//  Start Time: 2010-01-02 15:04:05.000000000 -0700 MST
	//
	// If the string length of 'LeftMarginStr' plus
	// 'RightMarginStr' plus the text label field length
	// ('TextLabelFieldLength') exceeds the maximum length
	// of 55-characters, this method will return an error.

	lock *sync.Mutex
}
