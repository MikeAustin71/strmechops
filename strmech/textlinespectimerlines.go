package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"strings"
	"sync"
	"time"
)

// TextLineSpecTimerLines - Records, computes and formats an
// elapsed time. This format process requires user input specifying
// a start time and ending time.
//
// The final formatted output string is composed of four lines of
// text for output as text display, file output or printing.
//
// The first line of text shows the Starting Time. The second line
// shows the Ending Time. The third line displays the time duration
// or the difference between starting time and ending
// time. The fourth line displays the total elapsed time in
// nanoseconds.
//
// The third line contains Time duration , or elapsed time, and is
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
// The TextLineSpecTimerLines methods provide a lot of flexibility
// for modifying this output format.
//
//
// ----------------------------------------------------------------
//
// Member Variables
//
//  labelLeftMarginChars       []rune
//     - The character or characters which will comprise the left
//       margin of the text label.
//
//
//  startTimeLabel             []rune
//     - A string containing the text characters constituting the
//       starting time text label.
//
//
//  startTime                  time.Time
//     - A time value which will be used in conjunction with
//       'endTime' parameter to compute the time duration or
//       elapsed for the timer event.
//
//
//  endTimeLabel               []rune
//     - A string containing the text characters constituting the
//       ending time text label.
//
//
//  endTime                    time.Time
//     - A time value which will be used in conjunction with
//       'startTime' parameter to compute the time duration or
//       elapsed for the timer event.
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
//
//  textLabelFieldLen          int
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
//
//  textLabelJustification         TextJustify
//     - An enumeration which specifies the justification of the
//       three text labels 'startTimeLabel', 'endTimeLabel' and
//       'timeDurationLabel' within the field length specified by
//       'textLabelFieldLen'.
//
//       Label justification must be equal to one of these three
//       valid values:
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
//       Example:
//        Start Time[right margin chars]2010-01-02 15:04:05.000000000 -0700 MST
//
//       Often this parameter is set to a single white space
//       character (" ") or a colon plus white space character,
//       (": ").
//
//       If this parameter is submitted as a zero length array, or
//       if rune array contains invalid zero rune values,
//       'labelRightMarginChars' will be assigned a default value
//       of  ": ".
//       Example Output:
//        Start Time: 2010-01-02 15:04:05.000000000 -0700 MST
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//
//   ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
//              "YourType.YourMethodName()",
//               "")
//
//   var startTime, endTime time.Time
//   startTime = time.Now()
//
//   someReturnVal := someType.Method()
//
//   endTime = time.Now()
//
//   timerLines,
//   err := TextLineSpecTimerLines{}.NewDefaultFullTimerEvent(
//   startTime,
//   endTime,
//   ePrefix)
//
//   if err != nil {
//     return err
//   }
//
//
//   fmt.Println("Time Required to Run Your Method")
//   fmt.Println()
//   fmt.Println(timerLines.String())
//
//
// ------------------------------------------------------------------------
//
// Sample Output
//
//    Start Time: 2021-08-13 03:19:32.462108100 -0500 CDT
//      End Time: 2021-08-13 03:19:32.462163100 -0500 CDT
//  Elapsed Time: 55 Microseconds 0 Nanoseconds
//                Total Elapsed Nanoseconds: 55,000
//
type TextLineSpecTimerLines struct {
	labelLeftMarginChars   []rune
	startTimeLabel         []rune
	startTime              time.Time
	endTimeLabel           []rune
	endTime                time.Time
	timeFormat             string
	timeDurationLabel      []rune
	textLabelFieldLen      int
	textLabelJustification TextJustify
	labelRightMarginChars  []rune
	textLineReader         *strings.Reader
	lock                   *sync.Mutex
}

// CopyIn - Copies all the data fields from an incoming instance of
// TextLineSpecTimerLines ('incomingTimerLines') to the data fields
// of the current TextLineSpecTimerLines instance
// ('txtSpecTimerLines').
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All the data fields in current TextLineSpecTimerLines instance
// ('txtSpecTimerLines') will be modified and overwritten.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingTimerLines         *TextLineSpecTimerLines
//     - A pointer to an instance of TextLineSpecTimerLines. This
//       method will NOT change the member variable data values
//       contained in this instance.
//
//       All data values in this TextLineSpecTimerLines instance
//       will be copied to current TextLineSpecTimerLines
//       instance ('txtSpecTimerLines').
//
//       If 'incomingTimerLines' contains invalid member data
//       variables, this method will return an error.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings containing
//                      error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (txtSpecTimerLines *TextLineSpecTimerLines) CopyIn(
	incomingTimerLines *TextLineSpecTimerLines,
	errorPrefix interface{}) error {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecTimerLines.CopyIn()",
		"")

	if err != nil {
		return err
	}

	return textLineSpecTimerLinesNanobot{}.ptr().
		copyIn(
			txtSpecTimerLines,
			incomingTimerLines,
			ePrefix)
}

// CopyOut - Returns a deep copy of the current
// TextLineSpecTimerLines instance.
//
// If the current TextLineSpecTimerLines instance contains invalid
// member variable data values, this method will return an error.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings containing
//                      error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  TextLineSpecTimerLines
//     - If this method completes successfully and no errors are
//       encountered, this parameter will return a deep copy of the
//       current TextLineSpecTimerLines instance.
//
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (txtSpecTimerLines *TextLineSpecTimerLines) CopyOut(
	errorPrefix interface{}) (
	TextLineSpecTimerLines,
	error) {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecTimerLines.CopyOut()",
		"")

	if err != nil {
		return TextLineSpecTimerLines{}, err
	}

	return textLineSpecTimerLinesNanobot{}.ptr().copyOut(
		txtSpecTimerLines,
		ePrefix.XCpy("txtSpecTimerLines->"))
}

// CopyOutITextLine - Returns a deep copy of the current
// TextLineSpecTimerLines instance cast as a type
// ITextLineSpecification.
//
// This method fulfills requirements of ITextLineSpecification
// interface.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings containing
//                      error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  ITextLineSpecification
//     - If this method completes successfully and no errors are
//       encountered, this parameter will return a deep copy of the
//       current TextLineSpecTimerLines instance cast as an
//       ITextLineSpecification object.
//
//
//  error
//     - If the method completes successfully and no errors are
//       encountered, this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (txtSpecTimerLines *TextLineSpecTimerLines) CopyOutITextLine(
	errorPrefix interface{}) (
	ITextLineSpecification,
	error) {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecTimerLines.CopyOutITextLine()",
		"")

	if err != nil {
		return &TextLineSpecTimerLines{}, err
	}

	var newTxtLineSpecTimerLine TextLineSpecTimerLines

	newTxtLineSpecTimerLine,
		err = textLineSpecTimerLinesNanobot{}.ptr().
		copyOut(
			txtSpecTimerLines,
			ePrefix.XCpy(
				"txtSpecTimerLines->"))

	if err != nil {
		return &TextLineSpecTimerLines{}, err
	}

	return ITextLineSpecification(
		&newTxtLineSpecTimerLine), nil
}

// CopyOutPtr - Returns a pointer to a deep copy of the current
// TextLineSpecTimerLines instance.
//
// If the current TextLineSpecTimerLines instance contains invalid
// member variable data values, this method will return an error.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings containing
//                      error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  *TextLineSpecTimerLines
//     - If this method completes successfully and no errors are
//       encountered, this parameter will return a pointer to a
//       deep copy of the current TextLineSpecTimerLines
//       instance.
//
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (txtSpecTimerLines *TextLineSpecTimerLines) CopyOutPtr(
	errorPrefix interface{}) (
	*TextLineSpecTimerLines,
	error) {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecTimerLines.CopyOutPtr()",
		"")

	if err != nil {
		return &TextLineSpecTimerLines{}, err
	}

	var newTxtLineSpecTimerLine TextLineSpecTimerLines

	newTxtLineSpecTimerLine,
		err = textLineSpecTimerLinesNanobot{}.ptr().
		copyOut(
			txtSpecTimerLines,
			ePrefix.XCpy(
				"txtSpecTimerLines->"))

	return &newTxtLineSpecTimerLine, err
}

// Empty - Resets all internal member variables to their initial
// or zero states.
//
// This method fulfills requirements of the ITextLineSpecification
// interface.
//
func (txtSpecTimerLines *TextLineSpecTimerLines) Empty() {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	textLineSpecTimerLinesElectron{}.ptr().
		empty(
			txtSpecTimerLines)

	txtSpecTimerLines.lock.Unlock()

	txtSpecTimerLines.lock = nil
}

// Equal - Receives a pointer to another instance of
// TextLineSpecTimerLines and proceeds to compare the member
// variables to those of the current TextLineSpecTimerLines
// instance in order to determine if they are equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables of both instances are equal in
// all respects, this flag is set to 'true'. Otherwise, this method
// returns 'false'.
//
func (txtSpecTimerLines *TextLineSpecTimerLines) Equal(
	incomingTxtSpecTimerLines *TextLineSpecTimerLines) bool {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	return textLineSpecTimerLinesAtom{}.ptr().
		equal(
			txtSpecTimerLines,
			incomingTxtSpecTimerLines)
}

// EqualITextLine - Receives an object implementing the
// ITextLineSpecification interface and proceeds to compare
// the member variables to those of the current
// TextLineSpecTimerLines instance in order to determine if
// they are equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables from both instances are equal
// in all respects, this flag is set to 'true'. Otherwise, this
// method returns 'false'.
//
// This method is required by interface ITextLineSpecification.
//
func (txtSpecTimerLines *TextLineSpecTimerLines) EqualITextLine(
	iTextLine ITextLineSpecification) bool {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	incomingTxtSpecTimerLines, ok :=
		iTextLine.(*TextLineSpecTimerLines)

	if !ok {
		return false
	}

	return textLineSpecTimerLinesAtom{}.ptr().
		equal(
			txtSpecTimerLines,
			incomingTxtSpecTimerLines)
}

// GetEndTimeLabel - Returns the internal member variable
// 'endTimeLabel' as a string.
//
// 'endTimeLabel' is a text label inserted in the output
// string to describe the ending time presentation.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  NONE
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  string
//     - This method returns the current value of the End Time
//       Label as a string.
//
func (txtSpecTimerLines *TextLineSpecTimerLines) GetEndTimeLabel() string {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	if len(txtSpecTimerLines.endTimeLabel) == 0 {
		return ""
	}

	return string(txtSpecTimerLines.endTimeLabel)
}

// GetFormattedText - Returns the formatted text generated by this
// Text Line Specification for text display output and printing.
//
// The current TextLineSpecTimerLines instance will provide the
// necessary information for generating text output which describes
// a timer event. The output will always consist of three lines of
// text specifying the event start time, end time and time duration
// or elapsed time.
//
// If an error occurs, the error message will be included in the
// returned string.
//
// This method is similar to TextLineSpecTimerLines.String() with
// the sole difference being that this method returns an error.
//
// This method fulfills requirements of the ITextLineSpecification
// interface.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings
//                      containing error prefix and error context
//                      information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package,
//       "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  formattedText              string
//     - The formatted text line generated by the current instance
//       of TextLineSpecTimerLines.
//
//
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (txtSpecTimerLines *TextLineSpecTimerLines) GetFormattedText(
	errorPrefix interface{}) (
	formattedText string,
	err error) {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecTimerLines.GetFormattedText()",
		"")

	if err != nil {
		return formattedText, err
	}

	formattedText,
		_,
		_,
		err = textLineSpecTimerLinesMolecule{}.ptr().
		getFormattedText(
			txtSpecTimerLines,
			ePrefix)

	return formattedText, err
}

// GetEndTime - Returns the internal member variable 'endTime'
// as time.Time value.
//
// The ending time is the end of the timing event described by this
// instance of TextLineSpecTimerLines.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  NONE
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  time.Time
//     - This method returns the current value of the "End Time"
//       for this Time Duration Event.
//
func (txtSpecTimerLines *TextLineSpecTimerLines) GetEndTime() time.Time {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	return txtSpecTimerLines.endTime
}

// GetLabelJustification - Returns the internal member variable
// 'textLabelJustification'.
//
// The Label Justification specification is of type, TextJustify.
// TextJustify is an enumeration which specifies the position of a
// string of text within a text field.
//
// The text justification enumeration specification must be set to
// one of these three valid values:
//           TextJustify(0).Left()
//           TextJustify(0).Right()
//           TextJustify(0).Center()
//
// You can also use the abbreviated text justification enumeration
// syntax as follows:
//
//           TxtJustify.Left()
//           TxtJustify.Right()
//           TxtJustify.Center()
//
func (txtSpecTimerLines *TextLineSpecTimerLines) GetLabelJustification() TextJustify {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	return txtSpecTimerLines.textLabelJustification
}

// GetLabelOutputSeparationChars - Returns the internal member
// variable, 'labelRightMarginChars', as a string.
//
// The 'labelRightMarginChars' is one or more characters used
// to separate the text labels, 'startTime', 'endTime' and
// 'timeDuration' from their respective output values. Usually,
// 'labelRightMarginChars' is set to either a single white
// space character or the default value which is a colon and a
// white space character (": ").
//
func (txtSpecTimerLines *TextLineSpecTimerLines) GetLabelOutputSeparationChars() string {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	if len(txtSpecTimerLines.labelRightMarginChars) == 0 {
		return ""
	}

	return string(txtSpecTimerLines.labelRightMarginChars)
}

// GetLengthLongestTextLabel - Compares the string lengths of the
// three text labels, Start Time Label, End Time Label and Time
// Duration Label, and returns the longest character length as an
// integer value.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  NONE
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  int
//     - This method analyzes the string lengths of the three text
//       labels, Start Time Label, End Time Label and Time Duration
//       Label. It then returns the longest character length as an
//       integer value.
//
func (txtSpecTimerLines *TextLineSpecTimerLines) GetLengthLongestTextLabel() int {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	txtTimerLinesElectron := textLineSpecTimerLinesElectron{}

	return txtTimerLinesElectron.getLengthOfLongestLabel(
		txtSpecTimerLines.startTimeLabel,
		txtSpecTimerLines.endTimeLabel,
		txtSpecTimerLines.timeDurationLabel)
}

// GetLengthTotalLabel - Returns the total length of the text label
// field length plus the lengths of the left and right margin
// fields.
//
// The text label field length is computed by taking the greatest of
// the longest text label length or the current value of the text
// label field length.
//
// The longest text label length is determined by calculating the
// character lengths of the three text labels: 'startTimeLabel',
// 'endTimeLabel' and 'timeDurationLabel'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  NONE
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  int
//     - This method computes the total label length and returns
//       this value as an integer.
//
func (txtSpecTimerLines *TextLineSpecTimerLines) GetLengthTotalLabel() int {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	txtTimerLinesElectron := textLineSpecTimerLinesElectron{}

	return txtTimerLinesElectron.
		getTotalLabelLength(
			txtSpecTimerLines.labelLeftMarginChars,
			txtSpecTimerLines.startTimeLabel,
			txtSpecTimerLines.endTimeLabel,
			txtSpecTimerLines.timeDurationLabel,
			txtSpecTimerLines.textLabelFieldLen,
			txtSpecTimerLines.labelRightMarginChars)

}

// GetMaximumTextLabelLength - Returns the maximum allowable text
// label length.
//
// The total text label length is computed by adding the text
// field length plus the lengths of the left and right margin
// fields.
//
// The text label field length is computed by taking the greatest of
// the longest text label length or the current value of the text
// label field length.
//
// The longest text label length is determined by calculating the
// character lengths of the three text labels: 'startTimeLabel',
// 'endTimeLabel' and 'timeDurationLabel'.
//
// The total text label length cannot exceed the maximum value
// returned by this method.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  NONE
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  int
//     - Returns the maximum allowable text label length as an
//       integer value.
//
func (txtSpecTimerLines *TextLineSpecTimerLines) GetMaximumTextLabelLength() int {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	return textLineSpecTimerLinesPreon{}.ptr().
		getMaximumTimerLabelLen()
}

// GetSingleLineLength - Returns the maximum line length calculated
// by comparing individual text lines which comprise the formatted
// text produced by the current instance of TextLineSpecTimerLines.
//
// Type TextLineSpecTimerLines differs from other Text Line
// Specifications in that it always produces multiple lines of text
// of varying string lengths.
//
// This method therefore analyzes the multiple lines of text which
// comprise the formatted text produced by this instance of
// TextLineSpecTimerLines to return the length of the longest text
// line.
//
// To obtain the total length of all text lines produced by the
// current instance of TextLineSpecTimerLines, see method:
//   TextLineSpecTimerLines.GetTotalLinesLength()
//
// Remember that the returned line length always includes the
// trailing new line characters.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  --- NONE ---
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  maxSingleLineLen           int
//     - This method will return an integer value specifying
//       the length of the longest text line included in
//       the formatted text produced by the current instance of
//       TextLineSpecTimerLines.
//
//       In the event of an error, a zero value will be returned.
//
func (txtSpecTimerLines *TextLineSpecTimerLines) GetSingleLineLength() (
	maxSingleLineLen int) {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()
	var err error
	_,
		maxSingleLineLen,
		_,
		err = textLineSpecTimerLinesMolecule{}.ptr().
		getFormattedText(
			txtSpecTimerLines,
			nil)

	if err != nil {
		maxSingleLineLen = 0
	}

	return maxSingleLineLen
}

// GetStartTimeLabel - Returns the internal member variable
// 'startTimeLabel' as a string.
//
// 'startTimeLabel' is a text label inserted in the output
// string to describe the starting time presentation.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  NONE
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  string
//     - This method returns the current value of the Start Time
//       Label as a string.
//
func (txtSpecTimerLines *TextLineSpecTimerLines) GetStartTimeLabel() string {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	if len(txtSpecTimerLines.startTimeLabel) == 0 {
		return ""
	}

	return string(txtSpecTimerLines.startTimeLabel)
}

// GetStartTime - Returns the internal member variable 'startTime'
// as time.Time value.
//
// The starting time is the start or beginning of the timing event
// described by this instance of TextLineSpecTimerLines.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  NONE
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  time.Time
//     - This method returns the current value of the "Start Time"
//       for this Time Duration Event.
//
func (txtSpecTimerLines *TextLineSpecTimerLines) GetStartTime() time.Time {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	return txtSpecTimerLines.startTime
}

// GetTextLabelFieldLength - Returns the internal member variable
// 'textLabelFieldLen' as an integer value.
//
// The Text Label Field Length specifies the length of the text
// field in which all three labels, 'startTimeLabel',
// 'endTimeLabel' and 'timeDurationLabel', will be positioned and
// justified.
//
// If Text Label Field Length is longer than the length of the
// longest label string, the Text Justification specification
// 'textLabelJustification' will be applied to position all three
// labels in their text fields. Text Justification can be set to
// 'Right', 'Left' or 'Center'.
//
// ----------------------------------------------------------------
//
// Be Advised
//
// If the text field length is less than the length of the longest
// label, the text field length will be reset to the length of the
// longest label.
//
// The length of the longest label is determined by calculating the
// character lengths of the three text labels: 'startTimeLabel',
// 'endTimeLabel' and 'timeDurationLabel'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  NONE
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  int
//     - This method returns the Text Label Field Length as an
//       integer value.
//
func (txtSpecTimerLines *TextLineSpecTimerLines) GetTextLabelFieldLength() int {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	lenLongestLabel := textLineSpecTimerLinesElectron{}.ptr().
		getLengthOfLongestLabel(
			txtSpecTimerLines.endTimeLabel,
			txtSpecTimerLines.startTimeLabel,
			txtSpecTimerLines.timeDurationLabel)

	if txtSpecTimerLines.textLabelFieldLen < lenLongestLabel {
		txtSpecTimerLines.textLabelFieldLen = lenLongestLabel
	}

	return txtSpecTimerLines.textLabelFieldLen
}

// GetTimeDurationLabel - Returns the internal member variable
// 'timeDurationLabel' as a string.
//
// 'timeDurationLabel' is a text label inserted in the
// TextLineSpecTimerLines output string to describe the elapsed
// time presentation. Time duration, or elapsed time, for this
// timer event is computed by subtracting the starting time from
// the ending time.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  NONE
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  string
//     - This method returns the current value of the Time Duration
//       Label as a string.
//
func (txtSpecTimerLines *TextLineSpecTimerLines) GetTimeDurationLabel() string {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	if len(txtSpecTimerLines.timeDurationLabel) == 0 {
		return ""
	}

	return string(txtSpecTimerLines.timeDurationLabel)
}

// GetTimeFormat - Returns the internal member variable
// 'timeFormat'. This time format specification is used to format
// all time values in text strings.
//
// The Go Programming language uses pattern formatting to format
// time values. Reference:
//   https://pkg.go.dev/time#Time.Format
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  NONE
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  string
//     - This method returns the current value of the Time Format
//       as a string.
//
func (txtSpecTimerLines *TextLineSpecTimerLines) GetTimeFormat() string {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	return txtSpecTimerLines.timeFormat
}

// GetTotalLinesLength - Returns the total length of all the
// formatted lines of text produced by the current instance of
// TextLineSpecTimerLines.
//
// Type TextLineSpecTimerLines differs from other Text Line
// Specifications in that it always produces multiple lines of text
// of varying string lengths.
//
// This method totals the line lengths of each line of text and
// returns this total value.
//
// To obtain the length of the longest single line of text, see
// method:
//   TextLineSpecTimerLines.GetSingleLineLength()
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  --- NONE ---
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  totalLinesLength           int
//     - This method will return an integer value specifying the
//       total length of all the formatted lines produced by the
//       current instance of TextLineSpecStandardLine.
//
//       In the event of an error, a zero value will be returned.
//
func (txtSpecTimerLines *TextLineSpecTimerLines) GetTotalLinesLength() (
	totalLinesLength int) {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	var err error
	_,
		_,
		totalLinesLength,
		err = textLineSpecTimerLinesMolecule{}.ptr().
		getFormattedText(
			txtSpecTimerLines,
			nil)

	if err != nil {
		totalLinesLength = 0
	}

	return totalLinesLength
}

// IsValidInstance - Performs a diagnostic review of the data
// values encapsulated in the current TextLineSpecTimerLines
// instance to determine if they are valid.
//
// If any data element evaluates as invalid, this method will
// return a boolean value of 'false'.
//
// If all data elements are determined to be valid, this method
// returns a boolean value of 'true'.
//
// This method is functionally equivalent to
// TextLineSpecTimerLines.IsValidInstanceError() with the sole
// exception being that this method takes no input parameters and
// returns a boolean value.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  -- NONE --
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  bool
//     - If any of the internal member data variables contained in
//       the current instance of TextLineSpecTimerLines are found
//       to be invalid, this method will return a boolean value of
//       'false'.
//
//       If all internal member data variables contained in the
//       current instance of TextLineSpecTimerLines are found to be
//       valid, this method returns a boolean value of 'true'.
//
func (txtSpecTimerLines *TextLineSpecTimerLines) IsValidInstance() bool {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	isValid,
		_ := textLineSpecTimerLinesAtom{}.ptr().
		testValidityOfTxtSpecTimerLines(
			txtSpecTimerLines,
			nil)

	return isValid
}

// IsValidInstanceError - Performs a diagnostic review of the data
// values encapsulated in the current TextLineSpecTimerLines
// instance to determine if they are valid.
//
// If any data element evaluates as invalid, this method will
// return an error.
//
// If the number of standard lines for the current
// TextLineSpecTimerLines instance is set to a value less than one,
// this method will treat that instance as invalid and return an
// error.
//
// This method fulfills requirements of ITextLineSpecification
// interface.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  errorPrefix         interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings containing
//                      error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  error
//     - If any of the internal member data variables contained in
//       the current instance of TextLineSpecTimerLines are found
//       to be invalid, this method will return an error.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' (error prefix) will be inserted or
//       prefixed at the beginning of the error message.
//
func (txtSpecTimerLines *TextLineSpecTimerLines) IsValidInstanceError(
	errorPrefix interface{}) error {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecTimerLines.IsValidInstanceError()",
		"")

	if err != nil {
		return err
	}

	_,
		err = textLineSpecTimerLinesAtom{}.ptr().
		testValidityOfTxtSpecTimerLines(
			txtSpecTimerLines,
			ePrefix.XCpy("txtSpecTimerLines"))

	return err
}

// NewDefaultFullTimerEvent - Creates and returns a new instance of
// TextLineSpecTimerLines using minimal input parameters. All the
// formatting parameters are set to standard default values.
//
// The purpose of a TextLineSpecTimerLines instance is to capture
// all the essential elements of a timer event and format that
// information for text display output or printing.
//
// Use this method when both start time and ending time for the
// timer event are known quantities and default formatting
// parameters are acceptable.
//
// ------------------------------------------------------------------------
//
// Default Values
//
// This method will automatically set the following default values:
//
//  labelLeftMarginChars  - Defaults to "" (Empty or Zero left margin)
//
//  startTimeLabel        - Defaults to "Start Time"
//
//  endTimeLabel          - Defaults to "End Time".
//
//  timeFormat            - Defaults to "2006-01-02 15:04:05.000000000 -0700 MST"
//
//  timeDurationLabel     - Defaults to "Elapsed Time"
//
//  textLabelFieldLen     - Defaults to  '12'
//
//  textLabelJustification    - Defaults to TextJustify(0).Right()
//
//  labelRightMarginChars - Defaults to ": "
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
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
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings containing
//                      error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  *TextLineSpecTimerLines
//     - If this method completes successfully, it will create and
//       return a pointer to a new instance of
//       TextLineSpecTimerLines which is fully configured with all
//       the parameters necessary to format a complete timer event
//       for text display, file output or printing.
//
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (txtSpecTimerLines TextLineSpecTimerLines) NewDefaultFullTimerEvent(
	startTime time.Time,
	endTime time.Time,
	errorPrefix interface{}) (
	*TextLineSpecTimerLines,
	error) {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	newTxtTimerLines := TextLineSpecTimerLines{}

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecTimerLines.NewDefaultFullTimerEvent()",
		"")

	if err != nil {
		return &newTxtTimerLines, err
	}

	timeDurationLabel :=
		textLineSpecTimerLinesElectron{}.ptr().
			getDefaultTimeDurationLabel()

	err = textLineSpecTimerLinesMolecule{}.ptr().
		setTxtLineSpecTimerLines(
			&newTxtTimerLines,
			nil,
			nil,
			startTime,
			nil,
			endTime,
			"",
			timeDurationLabel,
			len(timeDurationLabel),
			TxtJustify.Right(),
			nil,
			ePrefix.XCpy("newTxtTimerLines"))

	return &newTxtTimerLines, err
}

// NewDefaultShellTimerEvent - Creates and returns a new instance
// of TextLineSpecTimerLines using minimal input parameters. All
// the formatting parameters are set to standard default values.
//
// The return TextLineSpecTimerLines instance is only configured
// with basic formatting parameters and is NOT configured with the
// 'startTime' and 'endTime' parameters necessary for a fully
// configured timer event.
//
// Use this method when the 'startTime' and 'endTime' parameters
// are unknown. Follow up later to complete the timer event by
// calling one or more of the following methods when both starting
// time and ending time are known:
//
//   TextLineSpecTimerLines.SetStartAndEndTime()
//   - When both 'startTime' and 'endTime' are known.
//
//   TextLineSpecTimerLines.SetStartTime()
//   - When the 'startTime' is known.
//
//   TextLineSpecTimerLines.SetEndTime()
//   - When the 'startTime' is known.
//
// The purpose of a TextLineSpecTimerLines instance is to capture
// all the essential elements of a timer event and format that
// information for text display output or printing.
//
// ------------------------------------------------------------------------
//
// Default Values
//
// This method will automatically set the following default values:
//
//  labelLeftMarginChars  - Defaults to "" (Empty or Zero left margin)
//
//  startTime             - Defaults to July 4, 1776 9:30AM UTC
//
//  startTimeLabel        - Defaults to "Start Time"
//
//  endTimeLabel          - Defaults to "End Time".
//
//  endTime               - Defaults to July 4, 1776 9:30AM UTC
//
//  timeFormat            - Defaults to "2006-01-02 15:04:05.000000000 -0700 MST"
//
//  timeDurationLabel     - Defaults to "Elapsed Time"
//
//  textLabelFieldLen     - Defaults to  '12'
//
//  textLabelJustification    - Defaults to TextJustify(0).Right()
//
//  labelRightMarginChars - Defaults to ": "
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  --- NONE ---
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  *TextLineSpecTimerLines
//     - This method will create and return a pointer to a new
//       instance of TextLineSpecTimerLines which is fully
//       configured except for the starting time and ending times.
//
func (txtSpecTimerLines TextLineSpecTimerLines) NewDefaultShellTimerEvent() *TextLineSpecTimerLines {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	newTxtTimerLines := TextLineSpecTimerLines{}

	timeDurationLabel :=
		textLineSpecTimerLinesElectron{}.ptr().
			getDefaultTimeDurationLabel()

	defaultTime :=
		textLineSpecTimerLinesElectron{}.ptr().
			getDefaultTime()

	_ = textLineSpecTimerLinesMolecule{}.ptr().
		setTxtLineSpecTimerLines(
			&newTxtTimerLines,
			nil,
			nil,
			defaultTime,
			nil,
			defaultTime,
			"",
			timeDurationLabel,
			len(timeDurationLabel),
			TxtJustify.Right(),
			nil,
			nil)

	return &newTxtTimerLines
}

// NewEmptyTimerEvent - Creates and returns a new instance of
// TextLineSpecTimerLines which is empty, invalid and
// uninitialized.
//
// The returned TextLineSpecTimerLines instance contains internal
// member variables set to their native zero values. In order to
// use this empty object for calculations, you must first call the
// following method in order to set the internal member variables
// to valid values:
//    TextLineSpecTimerLines.SetFullTimerEvent()
//
func (txtSpecTimerLines TextLineSpecTimerLines) NewEmptyTimerEvent() TextLineSpecTimerLines {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	return TextLineSpecTimerLines{}
}

// NewFullTimerEvent - Creates and returns a new instance of
// TextLineSpecTimerLines which is fully configured with all the
// parameters necessary to output a complete timer event.
//
// The purpose of a TextLineSpecTimerLines instance is to capture
// all the essential elements of a timer event and format that
// information for text display output or printing.
//
// This method is designed to be used when both start time and
// ending time for the timer event are known quantities. However,
// one also has the option of setting 'startTime' and 'endTime' to
// default values by passing time.Time{} (a.k.a. zero time) as
// input parameters. 'startTime' and 'endTime' can then be entered
// later using the appropriate setter methods.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  labelLeftMarginChars       string
//     - The character or characters which will comprise the left
//       margin of the text label.
//
//       If the string length of 'labelLeftMarginChars' plus
//       'labelRightMarginChars' plus the text label field length
//       ('textLabelFieldLen') exceeds the maximum length of
//       55-characters, this method will return an error.
//
//       If labelLeftMarginChars is submitted as a zero length
//       string, no error will be triggered. In this case, the left
//       margin will simply be omitted.
//
//
//  startTimeLabel             string
//     - A string containing the text characters constituting the
//       starting time text label.
//
//       If this string is submitted as a zero length or empty
//       string, 'startTimeLabel' will be assigned a default value
//       of "Start Time".
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
//  endTimeLabel               string
//     - A string containing the text characters constituting the
//       ending time text label.
//
//       If this string is submitted as a zero length or empty
//       string, 'endTimeLabel' will be assigned a default value of
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
//  timeDurationLabel          string
//     - The text label used to describe the time duration or
//       elapsed time computed from the 'startTime' and 'endTime'
//       parameters.
//
//       If this string is submitted as a zero length or empty
//       string, 'timeDurationLabel' will be assigned a default
//       value of "Elapsed Time".
//
//       If the string length of 'labelLeftMarginChars' plus
//       'labelRightMarginChars' plus the text label field length
//       ('textLabelFieldLen') exceeds the maximum length of
//       55-characters, this method will return an error.
//
//
//  textLabelFieldLen          int
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
//       Label justification must be equal to one of these three
//       valid values:
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
//  labelRightMarginChars      string
//     - This string contains the character or characters which
//       will be used to separate the text labels ('startTimeLabel',
//       'endTimeLabel' and 'timeDurationLabel') from the output or
//       data values displayed on the same line.
//       Example:
//        Start Time[right margin chars]2010-01-02 15:04:05.000000000 -0700 MST
//
//       Often this parameter is set to a single white space
//       character (" ") or a colon plus white space character,
//       (": ").
//
//       If this string is submitted as a zero length or empty
//       string, 'labelRightMarginChars' will be assigned a
//       default value of  ": ". Example Output:
//        Start Time: 2010-01-02 15:04:05.000000000 -0700 MST
//
//       If the string length of 'labelLeftMarginChars' plus
//       'labelRightMarginChars' plus the text label field length
//       ('textLabelFieldLen') exceeds the maximum length of
//       55-characters, this method will return an error.
//
//       If labelLeftMarginChars is submitted as a zero length
//       string, no error will be triggered. In this case, the left
//       margin will simply be omitted.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings containing
//                      error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  *TextLineSpecTimerLines
//     - If this method completes successfully, it will create and
//       return a pointer to a new instance of
//       TextLineSpecTimerLines which is fully configured with all
//       the parameters necessary to format a complete timer event
//       for text display, file output or printing.
//
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (txtSpecTimerLines TextLineSpecTimerLines) NewFullTimerEvent(
	labelLeftMarginChars string,
	startTimeLabel string,
	startTime time.Time,
	endTimeLabel string,
	endTime time.Time,
	timeFormat string,
	timeDurationLabel string,
	textLabelFieldLen int,
	labelJustification TextJustify,
	labelRightMarginChars string,
	errorPrefix interface{}) (
	*TextLineSpecTimerLines,
	error) {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	newTxtTimerLines := TextLineSpecTimerLines{}

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecTimerLines.NewFullTimerEvent()",
		"")

	if err != nil {
		return &newTxtTimerLines, err
	}

	err = textLineSpecTimerLinesMolecule{}.ptr().
		setTxtLineSpecTimerLines(
			&newTxtTimerLines,
			[]rune(labelLeftMarginChars),
			[]rune(startTimeLabel),
			startTime,
			[]rune(endTimeLabel),
			endTime,
			timeFormat,
			[]rune(timeDurationLabel),
			textLabelFieldLen,
			labelJustification,
			[]rune(labelRightMarginChars),
			ePrefix.XCpy("newTxtTimerLines"))

	return &newTxtTimerLines, err
}

// NewFullTimerEventRunes - Creates and returns a new instance of
// TextLineSpecTimerLines which is fully configured with all the
// parameters necessary to output a complete timer event.
//
// The purpose of a TextLineSpecTimerLines instance is to capture
// all the essential elements of a timer event and format that
// information for text display output or printing.
//
// This method is designed to be used when both start time and
// ending time for the timer event are known quantities. However,
// one also has the option of setting 'startTime' and 'endTime' to
// default values by passing time.Time{} (a.k.a. zero time) as
// input parameters. 'startTime' and 'endTime' can then be entered
// later using the appropriate setter methods.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  labelLeftMarginChars       []rune
//     - The character or characters which will comprise the left
//       margin of the text label.
//
//       If the length in characters of 'labelLeftMarginChars' plus
//       'labelRightMarginChars' plus the text label field length
//       ('textLabelFieldLen') exceeds the maximum length of
//       55-characters, this method will return an error.
//
//       If labelLeftMarginChars is submitted with a 'nil' value or
//       a zero length rune array, no error will be triggered. In
//       this case, the left margin will simply be omitted.
//
//
//  startTimeLabel             []rune
//     - A string containing the text characters constituting the
//       starting time text label.
//
//       If this parameter is submitted as a zero length or empty
//       rune array, or if the rune array contains invalid zero
//       rune values, 'startTimeLabel' will be assigned a default
//       value of "Start Time".
//
//       If the length in characters of 'labelLeftMarginChars' plus
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
//     - A string containing the text characters constituting the
//       ending time text label.
//
//       If this parameter is submitted as a zero length or empty
//       rune array, or if the rune array contains invalid zero
//       rune values, 'endTimeLabel' will be assigned a default
//       value of "End Time".
//
//       If the length in characters of 'labelLeftMarginChars' plus
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
//       If this parameter is submitted as a zero length or empty
//       rune array, or if the rune array contains invalid zero
//       rune values, 'timeDurationLabel' will be assigned a default
//       value of "Elapsed Time".
//
//       If the length in characters of 'labelLeftMarginChars' plus
//       'labelRightMarginChars' plus the text label field length
//       ('textLabelFieldLen') exceeds the maximum length of
//       55-characters, this method will return an error.
//
//
//  textLabelFieldLen          int
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
//       If the length in characters of 'labelLeftMarginChars' plus
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
//       Label justification must be equal to one of these three
//       valid values:
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
//       Example:
//        Start Time[right margin chars]2010-01-02 15:04:05.000000000 -0700 MST
//
//       Often this parameter is set to a single white space
//       character (" ") or a colon plus white space character,
//       (": ").
//
//       If this parameter is submitted as a zero length array, or
//       if rune array contains invalid zero rune values,
//       'labelRightMarginChars' will be assigned a default value
//       of  ": ".
//       Example Output:
//        Start Time: 2010-01-02 15:04:05.000000000 -0700 MST
//
//       If the length in characters of 'labelLeftMarginChars' plus
//       'labelRightMarginChars' plus the text label field length
//       ('textLabelFieldLen') exceeds the maximum length of
//       55-characters, this method will return an error.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings containing
//                      error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  *TextLineSpecTimerLines
//     - If this method completes successfully, it will create and
//       return a pointer to a new instance of
//       TextLineSpecTimerLines which is fully configured with all
//       the parameters necessary to format a complete timer event
//       for text display, file output or printing.
//
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (txtSpecTimerLines TextLineSpecTimerLines) NewFullTimerEventRunes(
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
	errorPrefix interface{}) (
	*TextLineSpecTimerLines,
	error) {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	newTxtTimerLines := TextLineSpecTimerLines{}

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecTimerLines.NewFullTimerEventRunes()",
		"")

	if err != nil {
		return &newTxtTimerLines, err
	}

	err = textLineSpecTimerLinesMolecule{}.ptr().
		setTxtLineSpecTimerLines(
			&newTxtTimerLines,
			labelLeftMarginChars,
			startTimeLabel,
			startTime,
			endTimeLabel,
			endTime,
			timeFormat,
			timeDurationLabel,
			textLabelFieldLen,
			labelJustification,
			labelRightMarginChars,
			ePrefix.XCpy("newTxtTimerLines"))

	return &newTxtTimerLines, err
}

// NewShellTimerEvent - Creates and returns a new instance of
// TextLineSpecTimerLines. This instance is only configured with
// basic parameters and is NOT configured with the 'startTime'
// and 'endTime' parameters necessary for a fully configured timer event.
//
// Use this method when the 'startTime' and 'endTime' parameters
// are unknown. Follow up later to complete the timer event by
// calling one or more of the following methods when both starting
// time and ending time are known:
//
//   TextLineSpecTimerLines.SetStartAndEndTime()
//   - When both 'startTime' and 'endTime' are known.
//
//   TextLineSpecTimerLines.SetStartTime()
//   - When the 'startTime' is known.
//
//   TextLineSpecTimerLines.SetEndTime()
//   - When the 'startTime' is known.
//
// The purpose of a TextLineSpecTimerLines instance is to capture
// all the essential elements of a timer event and format that
// information for text display output or printing.
//
//
// ------------------------------------------------------------------------
//
// Default Values
//
// This method will automatically set the following default values:
//
//  startTime           - Defaults to July 4, 1776 9:30AM UTC
//
//  endTime             - Defaults to July 4, 1776 9:30AM UTC
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  labelLeftMarginChars       string
//     - The character or characters which will comprise the left
//       margin of the text label.
//
//       If the string length of 'labelLeftMarginChars' plus
//       'labelRightMarginChars' plus the text label field length
//       ('textLabelFieldLen') exceeds the maximum length of
//       55-characters, this method will return an error.
//
//
//  startTimeLabel             string
//     - A string containing the text characters constituting the
//       starting time text label.
//
//       If this string is submitted as a zero length or empty
//       string, 'startTimeLabel' will be assigned a default value
//       of "Start Time".
//
//
//       If the string length of 'labelLeftMarginChars' plus
//       'labelRightMarginChars' plus the text label field length
//       ('textLabelFieldLen') exceeds the maximum length of
//       55-characters, this method will return an error.
//
//
//  endTimeLabel               string.
//     - A string containing the text characters constituting the
//       ending time text label.
//
//       If this string is submitted as a zero length or empty
//       string, 'endTimeLabel' will be assigned a default value of
//       "End Time".
//
//       If the string length of 'labelLeftMarginChars' plus
//       'labelRightMarginChars' plus the text label field length
//       ('textLabelFieldLen') exceeds the maximum length of
//       55-characters, this method will return an error.
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
//  timeDurationLabel          string
//     - The text label used to describe the time duration or
//       elapsed time computed from the 'startTime' and 'endTime'
//       parameters.
//
//       If this string is submitted as a zero length or empty
//       string, 'timeDurationLabel' will be assigned a default
//       value of "Elapsed Time".
//
//       If the string length of 'labelLeftMarginChars' plus
//       'labelRightMarginChars' plus the text label field length
//       ('textLabelFieldLen') exceeds the maximum length of
//       55-characters, this method will return an error.
//
//
//  textLabelFieldLen          int
//     - The length of the text fields which will be used to
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
//       Label justification must be equal to one of these three
//       valid values:
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
//  labelRightMarginChars string
//     - This string contains the character or characters which
//       will be used to separate the text labels ('startTimeLabel',
//       'endTimeLabel' and 'timeDurationLabel') from the output or
//       data values displayed on the same line.
//       Example:
//        Start Time[sep chars]2010-01-02 15:04:05.000000000 -0700 MST
//
//       Often this parameter is set to a single white space
//       character (" ") or a colon plus white space character,
//       (": ").
//
//
//       If this string is submitted as a zero length or empty
//       string, 'labelRightMarginChars' will be assigned a
//       default value of  ": ". Example Output:
//        Start Time: 2010-01-02 15:04:05.000000000 -0700 MST
//
//       If the string length of 'labelLeftMarginChars' plus
//       'labelRightMarginChars' plus the text label field length
//       ('textLabelFieldLen') exceeds the maximum length of
//       55-characters, this method will return an error.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings containing
//                      error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  *TextLineSpecTimerLines
//     - If this method completes successfully, it will create and
//       return a pointer to a new instance of
//       TextLineSpecTimerLines which is fully configured except
//       for the starting time and ending time.
//
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (txtSpecTimerLines TextLineSpecTimerLines) NewShellTimerEvent(
	labelLeftMarginChars string,
	startTimeLabel string,
	endTimeLabel string,
	timeFormat string,
	timeDurationLabel string,
	textLabelFieldLen int,
	labelJustification TextJustify,
	labelRightMarginChars string,
	errorPrefix interface{}) (
	*TextLineSpecTimerLines,
	error) {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	newTxtTimerLines := TextLineSpecTimerLines{}

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecTimerLines.NewShellTimerEvent()",
		"")

	if err != nil {
		return &newTxtTimerLines, err
	}

	defaultTime :=
		textLineSpecTimerLinesElectron{}.ptr().
			getDefaultTime()

	err = textLineSpecTimerLinesMolecule{}.ptr().
		setTxtLineSpecTimerLines(
			&newTxtTimerLines,
			[]rune(labelLeftMarginChars),
			[]rune(startTimeLabel),
			defaultTime,
			[]rune(endTimeLabel),
			defaultTime,
			timeFormat,
			[]rune(timeDurationLabel),
			textLabelFieldLen,
			labelJustification,
			[]rune(labelRightMarginChars),
			ePrefix.XCpy("newTxtTimerLines"))

	return &newTxtTimerLines, err
}

// Read - Implements the io.Reader interface for type
// TextLineSpecTimerLines.
//
// The formatted text line string generated by the current
// instance of TextLineSpecTimerLines will be written to the
// byte buffer 'p'. If the length of 'p' is less than the length
// of the formatted text line string, multiple calls to this method
// will write the remaining unread characters to the byte buffer
// 'p'.
//
// Read() supports buffered 'read' operations.
//
// This method reads up to len(p) bytes into p. It returns the
// number of bytes read (0 <= n <= len(p)) and any error
// encountered. Even if read returns n < len(p), it may use all
// of p as scratch space during the call.
//
// If some data is available but not len(p) bytes, readBytes()
// conventionally returns what is available instead of waiting
// for more.
//
// When this method encounters an error or end-of-file condition
// after successfully reading n > 0 bytes, it returns the number
// of bytes read. It may return the (non-nil) error from the same
// call or return the error (and n == 0) from a subsequent call.
// An instance of this general case is that a Reader returning
// a non-zero number of bytes at the end of the input stream may
// return either err == EOF or err == nil. The next read operation
// should return 0, EOF.
//
// Callers should always process the n > 0 bytes returned before
// considering the error err. Doing so correctly handles I/O errors
// that happen after reading some bytes and also both of the
// allowed EOF behaviors.
//
// The last read operation performed on the formatted text string
// will always return n==0 and err==io.EOF.
//
// This method fulfills requirements of the ITextLineSpecification
// interface.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  p                          []byte
//     - The byte buffer into which the formatted text line string
//       generated by the current TextLineSpecTimerLines instance
//       will be written.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  n                          int
//     - The number of bytes written to byte buffer 'p'.
//
//       Read() reads up to len(p) bytes into p. It returns
//       the number of bytes read (0 <= n <= len(p)) and any error
//       encountered. Even if Read() returns n < len(p), it may use
//       all of 'p' as scratch space during the call. If some
//       data is available but not len(p) bytes, Read()
//       conventionally returns what is available instead of
//       waiting for more.
//
//
//  err                        error
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered
//       during processing, the returned error Type will
//       encapsulate an error message.
//
//       When Read() encounters an error or end-of-file condition
//       after successfully reading n > 0 bytes, it returns the
//       number of bytes read. It may return the (non-nil) error
//       from the same call or return the error (and n == 0) from
//       a subsequent call. An instance of this general case is
//       that a Reader returning a non-zero number of bytes at the
//       end of the input stream may return either err == EOF or
//       err == nil. The next read operation should return 0, EOF.
//
//
// ------------------------------------------------------------------------
//
// Usage Examples:
//
//  Example # 1
//
//  p := make([]byte, 50)
//
//  var n, readBytesCnt int
//  sb := strings.Builder{}
//
//  for {
//
//    n,
//    err = txtSpecTimerLines01.Read(p)
//
//    if n == 0 {
//      break
//    }
//
//    sb.Write(p[:n])
//    readBytesCnt += n
//  }
//
//  if err != nil &&
//    err != io.EOF {
//     return fmt.Errorf(
//      "Error Returned From txtSpecTimerLines01.Read(p)\n"+
//      "Error = \n%v\n",
//       err.Error())
//  }
//
//  fmt.Printf("Text Line String: %s\n",
//                sb.String())
//
//  fmt.Printf("Number of bytes Read: %v\n",
//                readBytesCnt)
//
//  Example # 2
//
//  p := make([]byte, 50)
//
//  var n, readBytesCnt int
//  var actualStr string
//
//  for {
//
//    n,
//    err = txtSpecTimerLines01.Read(p)
//
//    if n == 0 {
//      break
//    }
//
//    actualStr += string(p[:n])
//    readBytesCnt += n
//  }
//
//  if err != nil &&
//    err != io.EOF {
//     return fmt.Errorf(
//      "Error Returned From txtSpecTimerLines01.Read(p)\n"+
//      "Error = \n%v\n",
//       err.Error())
//  }
//
//  fmt.Printf("Text Line String: %v\n",
//                actualStr)
//
//  fmt.Printf("Number of bytes Read: %v\n",
//                readBytesCnt)
//
func (txtSpecTimerLines *TextLineSpecTimerLines) Read(
	p []byte) (
	n int,
	err error) {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TextLineSpecTimerLines.Read()",
		"")

	if txtSpecTimerLines.textLineReader == nil {

		var formattedText string

		formattedText,
			_,
			_,
			err = textLineSpecTimerLinesMolecule{}.ptr().
			getFormattedText(
				txtSpecTimerLines,
				ePrefix.XCpy("txtSpecTimerLines"))

		if err != nil {
			return n, err
		}

		txtSpecTimerLines.textLineReader =
			strings.NewReader(formattedText)

		if txtSpecTimerLines.textLineReader == nil {
			err = fmt.Errorf("%v\n"+
				"Error: strings.NewReader(formattedText)\n"+
				"returned a nil pointer.\n"+
				"txtSpecTimerLines.textLineReader == nil\n",
				ePrefix.String())

			return n, err
		}
	}

	n,
		err = textSpecificationAtom{}.ptr().
		readBytes(
			txtSpecTimerLines.textLineReader,
			p,
			ePrefix.XCpy(
				"p -> txtSpecTimerLines.textLineReader"))

	if err == io.EOF {

		txtSpecTimerLines.textLineReader = nil

	}

	return n, err
}

// ReaderInitialize - This method will reset the internal member
// variable 'TextLineSpecTimerLines.textLineReader' to its initial
// zero state of 'nil'.
//
// This method is rarely used. It provides a means of
// reinitializing the internal strings.Reader in case an error
// occurs during a read operation initiated by method
// TextLineSpecTimerLines.Read().
//
// Calling this method cleans up the residue from an aborted read
// operation and allows the calling function to start a new read
// operation.
//
// If any errors are returned by method
// TextLineSpecTimerLines.Read() which are NOT equal to io.EOF, call
// this method, TextLineSpecTimerLines.ReaderInitialize(), to reset
// the internal reader for future read operations.
//
// This method fulfills requirements of the ITextLineSpecification
// interface.
//
func (txtSpecTimerLines *TextLineSpecTimerLines) ReaderInitialize() {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	txtSpecTimerLines.textLineReader = nil

	return
}

// SetDefaultFullTimerEvent - This method configures the current
// instance of TextLineSpecTimerLines using minimal input
// parameters. All the formatting parameters are set to standard
// default values.
//
// The purpose of a TextLineSpecTimerLines instance is to capture
// all the essential elements of a timer event and format that
// information for text display output or printing.
//
// Use this method when both start time and ending time for the
// timer event are known quantities and default formatting
// parameters are acceptable.
//
// ------------------------------------------------------------------------
//
// IMPORTANT
//
// This method will overwrite and reset all the data values
// encapsulated by the current TextLineSpecTimerLines instance.
//
//
// ------------------------------------------------------------------------
//
// Default Values
//
// This method will automatically set the following default values:
//
//  labelLeftMarginChars  - Defaults to "" (Empty or Zero left margin)
//
//  startTimeLabel        - Defaults to "Start Time"
//
//  endTimeLabel          - Defaults to "End Time".
//
//  timeFormat            - Defaults to "2006-01-02 15:04:05.000000000 -0700 MST"
//
//  timeDurationLabel     - Defaults to "Elapsed Time"
//
//  textLabelFieldLen     - Defaults to  '12'
//
//  textLabelJustification    - Defaults to TextJustify(0).Right()
//
//  labelRightMarginChars - Defaults to ": "
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
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
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings containing
//                      error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (txtSpecTimerLines *TextLineSpecTimerLines) SetDefaultFullTimerEvent(
	startTime time.Time,
	endTime time.Time,
	errorPrefix interface{}) (
	err error) {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecTimerLines."+
			"SetDefaultFullTimerEvent()",
		"")

	if err != nil {
		return err
	}

	timeDurationLabel :=
		textLineSpecTimerLinesElectron{}.ptr().
			getDefaultTimeDurationLabel()

	err = textLineSpecTimerLinesMolecule{}.ptr().
		setTxtLineSpecTimerLines(
			txtSpecTimerLines,
			nil,
			nil,
			startTime,
			nil,
			endTime,
			"",
			timeDurationLabel,
			len(timeDurationLabel),
			TxtJustify.Right(),
			nil,
			ePrefix.XCpy("txtSpecTimerLines"))

	return err
}

// SetDefaultShellTimerEvent - Reconfigures the current instance of
// TextLineSpecTimerLines using minimal input parameters. All the
// formatting parameters are set to standard default values.
//
// Upon completion of this method, the current
// TextLineSpecTimerLines instance is only configured with basic
// default formatting parameters and is NOT configured with the
// 'startTime' and 'endTime' parameters necessary for a fully
// configured timer event.
//
// The intent is to create a shell or incomplete timer event with
// default formatting parameters. The 'startTime' and 'endTime'
// parameters can then be added later to create a fully configured
// timer event.
//
// Use this method when the 'startTime' and 'endTime' parameters
// are unknown and default formatting parameters are acceptable.
// Follow up later to complete the timer event by calling one or
// more of the following methods when both starting time and ending
// time are known:
//
//   TextLineSpecTimerLines.SetStartAndEndTime()
//   - When both 'startTime' and 'endTime' are known.
//
//   TextLineSpecTimerLines.SetStartTime()
//   - When the 'startTime' is known.
//
//   TextLineSpecTimerLines.SetEndTime()
//   - When the 'startTime' is known.
//
// The purpose of a TextLineSpecTimerLines instance is to capture
// all the essential elements of a timer event and format that
// information for text display output or printing.
//
// ------------------------------------------------------------------------
//
// IMPORTANT
//
// This method will overwrite and reset all the data values
// encapsulated by the current TextLineSpecTimerLines instance.
//
// ------------------------------------------------------------------------
//
// Default Values
//
// This method will automatically set the following default values:
//
//  labelLeftMarginChars  - Defaults to "" (Empty or Zero left margin)
//
//  startTime             - Defaults to July 4, 1776 9:30AM UTC
//
//  startTimeLabel        - Defaults to "Start Time"
//
//  endTime               - Defaults to July 4, 1776 9:30AM UTC
//
//  endTimeLabel          - Defaults to "End Time".
//
//  timeFormat            - Defaults to "2006-01-02 15:04:05.000000000 -0700 MST"
//
//  timeDurationLabel     - Defaults to "Elapsed Time"
//
//  textLabelFieldLen     - Defaults to  '12'
//
//  textLabelJustification    - Defaults to TextJustify(0).Right()
//
//  labelRightMarginChars - Defaults to ": "
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  --- NONE ---
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  --- NONE ---
//     - This method will automatically set the internal member
//       variables encapsulated by the current
//       TextLineSpecTimerLines to default values.
//
func (txtSpecTimerLines *TextLineSpecTimerLines) SetDefaultShellTimerEvent() {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	timeDurationLabel :=
		textLineSpecTimerLinesElectron{}.ptr().
			getDefaultTimeDurationLabel()

	defaultTime :=
		textLineSpecTimerLinesElectron{}.ptr().
			getDefaultTime()

	_ = textLineSpecTimerLinesMolecule{}.ptr().
		setTxtLineSpecTimerLines(
			txtSpecTimerLines,
			nil,
			nil,
			defaultTime,
			nil,
			defaultTime,
			"",
			timeDurationLabel,
			len(timeDurationLabel),
			TxtJustify.Right(),
			nil,
			nil)

}

// SetEndTimeLabel - Sets the internal member variable
// 'endTimeLabel'.
//
// 'endTimeLabel' is a text label inserted in the output
// string to describe the ending time presentation.
//
// If the input parameter 'endTimeLabel' is submitted as an empty
// string, its value will be defaulted to the string "End Time".
//
// If the total character length of the new End Time Label, plus
// the lengths of the Start Time Label, Time Duration label, the
// Left Margin Characters and the Right Margin Characters, exceeds
// the maximum available label length (55-characters), an error
// will be returned.
//
// ----------------------------------------------------------------
//
// Be Advised
//
// If this method completes successfully and the text field length
// is less than the length of the longest label, the text field
// length will be reset to the length of the longest label.
//
// The length of the longest text label is determined by
// calculating the character lengths of the three text labels:
// 'startTimeLabel', 'endTimeLabel' and 'timeDurationLabel'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  endTimeLabel               string
//     - A string containing the text characters constituting the
//       ending time text label.
//
//       If this string is submitted as a zero length or empty
//       string, 'endTimeLabel' will be assigned a default value of
//       "End Time".
//
//       If 'endTimeLabel' exceeds 50-characters in length,
//       this method will return an error.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings containing
//                      error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (txtSpecTimerLines *TextLineSpecTimerLines) SetEndTimeLabel(
	endTimeLabel string,
	errorPrefix interface{}) error {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error = nil

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecTimerLines."+
			"SetEndTimeLabel()",
		"")

	if err != nil {
		return err
	}

	lengthOfNewEndTimeLabel := len(endTimeLabel)

	maxAllowableLabelLen := textLineSpecTimerLinesPreon{}.ptr().
		getMaximumTimerLabelLen()

	if lengthOfNewEndTimeLabel > maxAllowableLabelLen {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'endTimeLabel'\n"+
			"exceeds the maximum label string length!\n"+
			"The maximum text label string length is: %v-characters\n"+
			"        The length of 'endTimeLabel' is: %v-characters\n",
			ePrefix.String(),
			maxAllowableLabelLen,
			lengthOfNewEndTimeLabel)

		return err
	}

	if lengthOfNewEndTimeLabel == 0 {

		txtSpecTimerLines.endTimeLabel =
			textLineSpecTimerLinesElectron{}.ptr().
				getDefaultEndTimeLabel()

	}

	endTimeLabelRunes := []rune(endTimeLabel)

	txtTimerLinesElectron := textLineSpecTimerLinesElectron{}

	lenLongestLabel := txtTimerLinesElectron.getLengthOfLongestLabel(
		txtSpecTimerLines.startTimeLabel,
		endTimeLabelRunes,
		txtSpecTimerLines.timeDurationLabel)

	totalLabelLen := txtTimerLinesElectron.
		getTotalLabelLength(
			txtSpecTimerLines.labelLeftMarginChars,
			txtSpecTimerLines.startTimeLabel,
			endTimeLabelRunes,
			txtSpecTimerLines.timeDurationLabel,
			txtSpecTimerLines.textLabelFieldLen,
			txtSpecTimerLines.labelRightMarginChars)

	if totalLabelLen > maxAllowableLabelLen {
		err = fmt.Errorf("%v\n"+
			"Error: The total length of the text label field is invalid!\n"+
			"The New 'End Time Label' will cause the entire text label\n"+
			"to exceed the maximum available text label field length.\n"+
			"The maximum text label field length is %v-characters\n"+
			"The total length of 'labelLeftMarginChars' plus 'labelRightMarginChars'"+
			"plus the the text label field length is %v-characters."+
			"'text label field length' is computed by taking the longest"+
			"text label length from the text labels, 'startTimeLabel',\n"+
			"'endTimeLabel' and 'timeDurationLabel'.\n"+
			"labelLeftMarginChars  = '%v'\n"+
			"startTimeLabel        = '%v'\n"+
			"New endTimeLabel      = '%v'\n"+
			"timeDurationLabel     = '%v'\n"+
			"labelRightMarginChars = '%v'\n"+
			"textLabelFieldLen     = '%v'\n",
			ePrefix.String(),
			maxAllowableLabelLen,
			totalLabelLen,
			len(txtSpecTimerLines.labelLeftMarginChars),
			len(endTimeLabelRunes),
			len(txtSpecTimerLines.endTimeLabel),
			len(txtSpecTimerLines.timeDurationLabel),
			len(txtSpecTimerLines.labelRightMarginChars),
			txtSpecTimerLines.textLabelFieldLen)

		return err
	}

	txtSpecTimerLines.endTimeLabel = endTimeLabelRunes

	if txtSpecTimerLines.textLabelFieldLen < lenLongestLabel {
		txtSpecTimerLines.textLabelFieldLen = lenLongestLabel
	}

	return err
}

// SetEndTime - Sets the internal member variable 'endTime'
// as time.Time value.
//
// The ending time is the end of the timing event described by this
// instance of TextLineSpecTimerLines.
//
// If the input parameter 'endTime' is submitted as a zero time
// value, this method will return an error.
//
// Of course, the ending time should occur after the starting
// time for this instance of TextLineSpecTimerLines.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  endTime                    time.Time
//     - The time at which the time event ends. The timer event
//       refers to the subject timer event of the current
//       TextLineSpecTimerLines instance.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings containing
//                      error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (txtSpecTimerLines *TextLineSpecTimerLines) SetEndTime(
	endTime time.Time,
	errorPrefix interface{}) error {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error = nil

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecTimerLines.SetEndTime()",
		"")

	if err != nil {
		return err
	}

	if endTime.IsZero() {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'endTime' is set to a ZERO value!\n",
			ePrefix.String())

		return err
	}

	txtSpecTimerLines.endTime = endTime

	return err
}

// SetFullTimerEvent - This method configures the current
// TextLineSpecTimerLines instance with new timer event parameters.
// All internal member variable data values for the current
// TextLineSpecTimerLines instance will be reset and overwritten.
//
// The purpose of a TextLineSpecTimerLines instance is to capture
// all the essential elements of a timer event and format that
// information for text display, file output or printing.
//
// This method is designed to be used when both start time and
// ending time for the timer event are known quantities. However,
// one also has the option of setting 'startTime' and 'endTime' to
// default values by passing time.Time{} (a.k.a. zero time) as
// input parameters. 'startTime' and 'endTime' can then be entered
// later using the appropriate setter methods.
//
// This method is very similar to
// TextLineSpecTimerLines.SetFullTimerEventRunes() and differs only
// in the respect that many input parameters use strings instead of
// rune arrays.
//
// ------------------------------------------------------------------------
//
// IMPORTANT
//
// This method will overwrite and reset all the data values
// encapsulated by the current TextLineSpecTimerLines instance.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  labelLeftMarginChars       string
//     - The character or characters which will comprise the left
//       margin of the text label.
//
//       If the string length of 'labelLeftMarginChars' plus
//       'labelRightMarginChars' plus the text label field length
//       ('textLabelFieldLen') exceeds the maximum length of
//       55-characters, this method will return an error.
//
//
//  startTimeLabel             string
//     - A string containing the text characters constituting the
//       starting time text label.
//
//       If this string is submitted as a zero length or empty
//       string, 'startTimeLabel' will be assigned a default value
//       of "Start Time".
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
//  endTimeLabel               string.
//     - A string containing the text characters constituting the
//       ending time text label.
//
//       If this string is submitted as a zero length or empty
//       string, 'endTimeLabel' will be assigned a default value of
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
//  timeDurationLabel          string
//     - The text label used to describe the time duration or
//       elapsed time computed from the 'startTime' and 'endTime'
//       parameters.
//
//       If this string is submitted as a zero length or empty
//       string, 'timeDurationLabel' will be assigned a default
//       value of "Elapsed Time".
//
//       If the string length of 'labelLeftMarginChars' plus
//       'labelRightMarginChars' plus the text label field length
//       ('textLabelFieldLen') exceeds the maximum length of
//       55-characters, this method will return an error.
//
//
//  textLabelFieldLen          int
//     - The length of the text fields which will be used to
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
//       Label justification must be equal to one of these three
//       valid values:
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
//  labelRightMarginChars      string
//     - This string contains the character or characters which
//       will be used to separate the text labels ('startTimeLabel',
//       'endTimeLabel' and 'timeDurationLabel') from the output or
//       data values displayed on the same line.
//       Example:
//        Start Time[right margin chars]2010-01-02 15:04:05.000000000 -0700 MST
//
//       Often this parameter is set to a single white space
//       character (" ") or a colon plus white space character,
//       (": ").
//
//       If this string is submitted as a zero length or empty
//       string, 'labelRightMarginChars' will be assigned a
//       default value of  ": ". Example Output:
//        Start Time: 2010-01-02 15:04:05.000000000 -0700 MST
//
//       If the string length of 'labelLeftMarginChars' plus
//       'labelRightMarginChars' plus the text label field length
//       ('textLabelFieldLen') exceeds the maximum length of
//       55-characters, this method will return an error.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings containing
//                      error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (txtSpecTimerLines *TextLineSpecTimerLines) SetFullTimerEvent(
	labelLeftMarginChars string,
	startTimeLabel string,
	startTime time.Time,
	endTimeLabel string,
	endTime time.Time,
	timeFormat string,
	timeDurationLabel string,
	textLabelFieldLen int,
	labelJustification TextJustify,
	labelRightMarginChars string,
	errorPrefix interface{}) (
	err error) {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecTimerLines.SetFullTimerEvent()",
		"")

	if err != nil {
		return err
	}

	err = textLineSpecTimerLinesMolecule{}.ptr().
		setTxtLineSpecTimerLines(
			txtSpecTimerLines,
			[]rune(labelLeftMarginChars),
			[]rune(startTimeLabel),
			startTime,
			[]rune(endTimeLabel),
			endTime,
			timeFormat,
			[]rune(timeDurationLabel),
			textLabelFieldLen,
			labelJustification,
			[]rune(labelRightMarginChars),
			ePrefix.XCpy("txtSpecTimerLines"))

	return err
}

// SetFullTimerEventRunes - This method configures the current
// TextLineSpecTimerLines instance with new timer event parameters.
// All internal member variable data values for the current
// TextLineSpecTimerLines instance will be reset and overwritten.
//
// The purpose of a TextLineSpecTimerLines instance is to capture
// all the essential elements of a timer event and format that
// information for text display, file output or printing.
//
// This method is designed to be used when both start time and
// ending time for the timer event are known quantities. However,
// one also has the option of setting 'startTime' and 'endTime' to
// default values by passing time.Time{} (a.k.a. zero time) as
// input parameters. 'startTime' and 'endTime' can then be entered
// later using the appropriate setter methods.
//
// This method is very similar to
// TextLineSpecTimerLines.SetFullTimerEvent() and differs only in
// the respect that many input parameters use rune arrays instead of
// strings.
//
// ------------------------------------------------------------------------
//
// IMPORTANT
//
// This method will overwrite and reset all the data values
// encapsulated by the current TextLineSpecTimerLines instance.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  labelLeftMarginChars       []rune
//     - The character or characters which will comprise the left
//       margin of the text label.
//
//       If the length in characters of 'labelLeftMarginChars' plus
//       'labelRightMarginChars' plus the text label field length
//       ('textLabelFieldLen') exceeds the maximum length of
//       55-characters, this method will return an error.
//
//       If labelLeftMarginChars is submitted with a 'nil' value or
//       a zero length rune array, no error will be triggered. In
//       this case, the left margin will simply be omitted.
//
//
//  startTimeLabel             []rune
//     - A string containing the text characters constituting the
//       starting time text label.
//
//       If this parameter is submitted as a zero length or empty
//       rune array, or if the rune array contains invalid zero
//       rune values, 'startTimeLabel' will be assigned a default
//       value of "Start Time".
//
//       If the length in characters of 'labelLeftMarginChars' plus
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
//     - A string containing the text characters constituting the
//       ending time text label.
//
//       If this parameter is submitted as a zero length or empty
//       rune array, or if the rune array contains invalid zero
//       rune values, 'endTimeLabel' will be assigned a default
//       value of "End Time".
//
//       If the length in characters of 'labelLeftMarginChars' plus
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
//       If this parameter is submitted as a zero length or empty
//       rune array, or if the rune array contains invalid zero
//       rune values, 'timeDurationLabel' will be assigned a default
//       value of "Elapsed Time".
//
//       If the length in characters of 'labelLeftMarginChars' plus
//       'labelRightMarginChars' plus the text label field length
//       ('textLabelFieldLen') exceeds the maximum length of
//       55-characters, this method will return an error.
//
//
//  textLabelFieldLen          int
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
//       If the length in characters of 'labelLeftMarginChars' plus
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
//       Label justification must be equal to one of these three
//       valid values:
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
//       Example:
//        Start Time[right margin chars]2010-01-02 15:04:05.000000000 -0700 MST
//
//       Often this parameter is set to a single white space
//       character (" ") or a colon plus white space character,
//       (": ").
//
//       If this parameter is submitted as a zero length array, or
//       if rune array contains invalid zero rune values,
//       'labelRightMarginChars' will be assigned a default value
//       of  ": ".
//       Example Output:
//        Start Time: 2010-01-02 15:04:05.000000000 -0700 MST
//
//       If the length in characters of 'labelLeftMarginChars' plus
//       'labelRightMarginChars' plus the text label field length
//       ('textLabelFieldLen') exceeds the maximum length of
//       55-characters, this method will return an error.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings containing
//                      error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (txtSpecTimerLines *TextLineSpecTimerLines) SetFullTimerEventRunes(
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
	errorPrefix interface{}) (
	err error) {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecTimerLines."+
			"SetFullTimerEventRunes()",
		"")

	if err != nil {
		return err
	}

	err = textLineSpecTimerLinesMolecule{}.ptr().
		setTxtLineSpecTimerLines(
			txtSpecTimerLines,
			labelLeftMarginChars,
			startTimeLabel,
			startTime,
			endTimeLabel,
			endTime,
			timeFormat,
			timeDurationLabel,
			textLabelFieldLen,
			labelJustification,
			labelRightMarginChars,
			ePrefix.XCpy("txtSpecTimerLines"))

	return err
}

// SetLabelJustification - Sets the text justification for the text
// labels used to describe various aspects of the timer event.
// Valid text justifications are 'Left', 'Right' or 'Center'.
//
// Each instance of TextLineSpecTimerLines describes a timer event
// which uses three labels to describe the key elements of that
// event. These labels are:
//   (1) TextLineSpecTimerLines.startTimeLabel
//       - Describes the starting time for this timer event.
//
//   (2) TextLineSpecTimerLines.endTimeLabel
//       - Describes the ending time for this timer event.
//
//   (3) TextLineSpecTimerLines.timeDuration
//       - Describes the time duration or elapsed time for this
//         timer event.
//
// The length of the text field used to display these three labels
// is defined either by the user specified parameter 'Label Field
// Length' or the string length of the longest label string.
//
// Within this text field labels are positioned according to the
// label justification specified by the user.
//
// The Label Justification specification is of type, TextJustify.
// TextJustify is an enumeration which specifies the position of a
// string of text within a text field.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  textLabelJustification         TextJustify
//     - An enumeration which specifies the justification of the
//       three text labels 'startTimeLabel', 'endTimeLabel' and
//       'timeDurationLabel' within the field length specified by
//       'textLabelFieldLen'.
//
//       Label justification must be equal to one of these three
//       valid values:
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
//       If the input parameter for label justification
//       ('textLabelJustification') is set to a value other than 'Left',
//       'Right' or 'Center', this method will return an error.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings containing
//                      error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (txtSpecTimerLines *TextLineSpecTimerLines) SetLabelJustification(
	labelJustification TextJustify,
	errorPrefix interface{}) error {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error = nil

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecTimerLines.SetLabelJustification()",
		"")

	if err != nil {
		return err
	}

	if !labelJustification.XIsValid() {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'textLabelJustification' is invalid!\n"+
			"'textLabelJustification' must be set to 'Left', 'Right' or 'Center'\n"+
			"textLabelJustification string = '%v'\n"+
			"textLabelJustification int = '%v'\n",
			ePrefix.String(),
			labelJustification.String(),
			labelJustification.XValueInt())

		return err
	}

	txtSpecTimerLines.textLabelJustification =
		labelJustification

	return err
}

// SetLabelRightMarginChars - Sets the internal member
// variable, 'labelRightMarginChars'.
//
// The 'labelRightMarginChars' is one or more characters used
// to separate the text labels, 'startTime', 'endTime' and
// 'timeDuration' from their respective output values. Usually,
// 'labelRightMarginChars' is set to either a single white
// space character or the default value which is a colon and a
// white space character (": ").
//       Examples:
//        Start Time[sep chars]2010-01-02 15:04:05.000000000 -0700 MST
//        Start Time: 2010-01-02 15:04:05.000000000 -0700 MST
//
// If input parameter 'labelRightMarginChars' is set to an
// empty string, its value will be defaulted to a colon and a white
// space character  (": ").
//
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  labelRightMarginChars      string
//     - This string contains the character or characters which
//       will be used to separate the text labels ('startTimeLabel',
//       'endTimeLabel' and 'timeDurationLabel') from the output or
//       data values displayed on the same line.
//       Example:
//        Start Time[right margin chars]2010-01-02 15:04:05.000000000 -0700 MST
//
//       Often this parameter is set to a single white space
//       character (" ") or a colon plus white space character,
//       (": ").
//
//       If this string is submitted as a zero length or empty
//       string, 'labelRightMarginChars' will be assigned a
//       default value of  ": ". Example Output:
//        Start Time: 2010-01-02 15:04:05.000000000 -0700 MST
//
//       If the string length of 'labelLeftMarginChars' plus
//       'labelRightMarginChars' plus the text label field length
//       ('textLabelFieldLen') exceeds the maximum length of
//       55-characters, this method will return an error.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings containing
//                      error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (txtSpecTimerLines *TextLineSpecTimerLines) SetLabelRightMarginChars(
	labelRightMarginChars string,
	errorPrefix interface{}) error {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error = nil

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecTimerLines."+
			"SetLabelRightMarginChars()",
		"")

	if err != nil {
		return err
	}

	var labelRightMarginCharsRunes []rune

	if len(labelRightMarginChars) == 0 {

		labelRightMarginCharsRunes =
			textLineSpecTimerLinesElectron{}.ptr().
				getDefaultLabelRightMarginChars()

	} else {

		labelRightMarginCharsRunes =
			[]rune(labelRightMarginChars)
	}

	sMechPreon := strMechPreon{}

	_,
		err = sMechPreon.testValidityOfRuneCharArray(
		labelRightMarginCharsRunes,
		ePrefix.XCpy(
			"labelRightMarginChars->"+
				"labelRightMarginCharsRunes"))

	if err != nil {
		return err
	}

	err = sMechPreon.copyRuneArrays(
		&txtSpecTimerLines.labelRightMarginChars,
		&labelRightMarginCharsRunes,
		true,
		ePrefix.XCpy(
			"labelRightMarginChars-> "+
				"labelRightMarginCharsRunes-> "+
				"txtSpecTimerLines.labelRightMarginChars"))

	return err
}

// SetShellTimerEvent - Reconfigures the current instance of
// TextLineSpecTimerLines.
//
// Upon completion of this method, the current
// TextLineSpecTimerLines instance is only configured with basic
// formatting parameters and is NOT configured with the 'startTime'
// and 'endTime' parameters necessary for a fully configured timer
// event.
//
// Use this method when the 'startTime' and 'endTime' parameters
// are unknown. Follow up later to complete the timer event by
// calling one or more of the following methods when both starting
// time and ending time are known:
//
//   TextLineSpecTimerLines.SetStartAndEndTime()
//   - When both 'startTime' and 'endTime' are known.
//
//   TextLineSpecTimerLines.SetStartTime()
//   - When the 'startTime' is known.
//
//   TextLineSpecTimerLines.SetEndTime()
//   - When the 'startTime' is known.
//
// The purpose of a TextLineSpecTimerLines instance is to capture
// all the essential elements of a timer event and format that
// information for text display output or printing.
//
//
// ------------------------------------------------------------------------
//
// Default Values
//
// This method will automatically set the following default values:
//
//  startTime           - Defaults to July 4, 1776 9:30AM UTC
//
//  endTime             - Defaults to July 4, 1776 9:30AM UTC
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  labelLeftMarginChars       string
//     - The character or characters which will comprise the left
//       margin of the text label.
//
//       If the string length of 'labelLeftMarginChars' plus
//       'labelRightMarginChars' plus the text label field length
//       ('textLabelFieldLen') exceeds the maximum length of
//       55-characters, this method will return an error.
//
//
//  startTimeLabel             string
//     - A string containing the text characters constituting the
//       starting time text label.
//
//       If this string is submitted as a zero length or empty
//       string, 'startTimeLabel' will be assigned a default value
//       of "Start Time".
//
//
//       If the string length of 'labelLeftMarginChars' plus
//       'labelRightMarginChars' plus the text label field length
//       ('textLabelFieldLen') exceeds the maximum length of
//       55-characters, this method will return an error.
//
//
//  endTimeLabel               string.
//     - A string containing the text characters constituting the
//       ending time text label.
//
//       If this string is submitted as a zero length or empty
//       string, 'endTimeLabel' will be assigned a default value of
//       "End Time".
//
//       If the string length of 'labelLeftMarginChars' plus
//       'labelRightMarginChars' plus the text label field length
//       ('textLabelFieldLen') exceeds the maximum length of
//       55-characters, this method will return an error.
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
//  timeDurationLabel          string
//     - The text label used to describe the time duration or
//       elapsed time computed from the 'startTime' and 'endTime'
//       parameters.
//
//       If this string is submitted as a zero length or empty
//       string, 'timeDurationLabel' will be assigned a default
//       value of "Elapsed Time".
//
//       If the string length of 'labelLeftMarginChars' plus
//       'labelRightMarginChars' plus the text label field length
//       ('textLabelFieldLen') exceeds the maximum length of
//       55-characters, this method will return an error.
//
//
//  textLabelFieldLen          int
//     - The length of the text fields which will be used to
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
//       Label justification must be equal to one of these three
//       valid values:
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
//  labelRightMarginChars string
//     - This string contains the character or characters which
//       will be used to separate the text labels ('startTimeLabel',
//       'endTimeLabel' and 'timeDurationLabel') from the output or
//       data values displayed on the same line.
//       Example:
//        Start Time[sep chars]2010-01-02 15:04:05.000000000 -0700 MST
//
//       Often this parameter is set to a single white space
//       character (" ") or a colon plus white space character,
//       (": ").
//
//
//       If this string is submitted as a zero length or empty
//       string, 'labelRightMarginChars' will be assigned a
//       default value of  ": ". Example Output:
//        Start Time: 2010-01-02 15:04:05.000000000 -0700 MST
//
//       If the string length of 'labelLeftMarginChars' plus
//       'labelRightMarginChars' plus the text label field length
//       ('textLabelFieldLen') exceeds the maximum length of
//       55-characters, this method will return an error.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings containing
//                      error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (txtSpecTimerLines *TextLineSpecTimerLines) SetShellTimerEvent(
	labelLeftMarginChars string,
	startTimeLabel string,
	endTimeLabel string,
	timeFormat string,
	timeDurationLabel string,
	textLabelFieldLen int,
	labelJustification TextJustify,
	labelRightMarginChars string,
	errorPrefix interface{}) (
	err error) {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecTimerLines."+
			"SetShellTimerEvent()",
		"")

	if err != nil {
		return err
	}

	defaultTime :=
		textLineSpecTimerLinesElectron{}.ptr().
			getDefaultTime()

	err = textLineSpecTimerLinesMolecule{}.ptr().
		setTxtLineSpecTimerLines(
			txtSpecTimerLines,
			[]rune(labelLeftMarginChars),
			[]rune(startTimeLabel),
			defaultTime,
			[]rune(endTimeLabel),
			defaultTime,
			timeFormat,
			[]rune(timeDurationLabel),
			textLabelFieldLen,
			labelJustification,
			[]rune(labelRightMarginChars),
			ePrefix.XCpy("newTxtTimerLines"))

	return err
}

// SetStartTime - Sets the internal member variable 'startTime'
// as time.Time value.
//
// The starting time is the start or beginning of the timing event
// described by this instance of TextLineSpecTimerLines.
//
// If the input parameter 'startTime' is submitted as a zero time
// value, this method will return an error.
//
// Of course, the starting time should occur before the ending
// time for this instance of TextLineSpecTimerLines.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  startTime                  time.Time
//     - The time at which the time event starts. The timer event
//       refers to the subject timer event of the current
//       TextLineSpecTimerLines instance.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings containing
//                      error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (txtSpecTimerLines *TextLineSpecTimerLines) SetStartTime(
	startTime time.Time,
	errorPrefix interface{}) error {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error = nil

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecTimerLines."+
			"SetStartTime()",
		"")

	if err != nil {
		return err
	}

	if startTime.IsZero() {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'startTime' is set to a ZERO value!\n",
			ePrefix.String())

		return err
	}

	txtSpecTimerLines.startTime = startTime

	return err
}

// SetStartTimeLabel - Sets the internal member variable
// 'startTimeLabel'.
//
// 'startTimeLabel' is a text label inserted in the output
// string to describe the starting time presentation.
//
// If input parameter 'startTimeLabel' is submitted as an empty
// string, its value will be defaulted to the string "Start Time".
//
// If the total character length of the new Start Time Label, plus
// the lengths of the End Time Label, Time Duration label, the Left
// Margin Characters and the Right Margin Characters, exceeds the
// maximum available label length (55-characters), an error will be
// returned.
//
// ----------------------------------------------------------------
//
// Be Advised
//
// If this method completes successfully and the text field length
// is less than the length of the longest label, the text field
// length will be reset to the length of the longest label.
//
// The length of the longest label is determined by calculating the
// character lengths of the three text labels: 'startTimeLabel',
// 'endTimeLabel' and 'timeDurationLabel'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  startTimeLabel             string
//     - A string containing the text characters constituting the
//       starting time text label.
//
//       If this string is submitted as a zero length or empty
//       string, 'startTimeLabel' will be assigned a default value
//       of "Start Time".
//
//       If 'startTimeLabel' exceeds 50-characters in length,
//       this method will return an error.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings containing
//                      error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (txtSpecTimerLines *TextLineSpecTimerLines) SetStartTimeLabel(
	startTimeLabel string,
	errorPrefix interface{}) error {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error = nil

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecTimerLines."+
			"SetEndTimeLabel()",
		"")

	if err != nil {
		return err
	}

	lengthOfNewStartTimeLabel := len(startTimeLabel)

	maxAllowableLabelLen := textLineSpecTimerLinesPreon{}.ptr().
		getMaximumTimerLabelLen()

	if lengthOfNewStartTimeLabel > maxAllowableLabelLen {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'startTimeLabel'\n"+
			"exceeds the maximum label string length!\n"+
			"The maximum text label string length is: %v-characters\n"+
			"      The length of 'startTimeLabel' is: %v-characters\n",
			ePrefix.String(),
			maxAllowableLabelLen,
			lengthOfNewStartTimeLabel)

		return err
	}

	if lengthOfNewStartTimeLabel == 0 {
		txtSpecTimerLines.startTimeLabel =
			textLineSpecTimerLinesElectron{}.ptr().
				getDefaultStartTimeLabel()
	}

	startTimeLabelRunes := []rune(startTimeLabel)

	txtTimerLinesElectron := textLineSpecTimerLinesElectron{}

	lenLongestLabel := txtTimerLinesElectron.getLengthOfLongestLabel(
		startTimeLabelRunes,
		txtSpecTimerLines.endTimeLabel,
		txtSpecTimerLines.timeDurationLabel)

	totalLabelLen := txtTimerLinesElectron.
		getTotalLabelLength(
			txtSpecTimerLines.labelLeftMarginChars,
			startTimeLabelRunes,
			txtSpecTimerLines.endTimeLabel,
			txtSpecTimerLines.timeDurationLabel,
			txtSpecTimerLines.textLabelFieldLen,
			txtSpecTimerLines.labelRightMarginChars)

	if totalLabelLen > maxAllowableLabelLen {
		err = fmt.Errorf("%v\n"+
			"Error: The total length of the text label field is invalid!\n"+
			"The New 'Start Time Label' will cause the entire text label\n"+
			"to exceed the maximum available text label field length.\n"+
			"The maximum text label field length is %v-characters\n"+
			"The total length of 'labelLeftMarginChars' plus 'labelRightMarginChars'"+
			"plus the the text label field length is %v-characters."+
			"'text label field length' is computed by taking the longest"+
			"text label length.\n"+
			"labelLeftMarginChars  = '%v'\n"+
			"New startTimeLabel    = '%v'\n"+
			"endTimeLabel          = '%v'\n"+
			"timeDurationLabel     = '%v'\n"+
			"labelRightMarginChars = '%v'\n"+
			"textLabelFieldLen     = '%v'\n",
			ePrefix.String(),
			maxAllowableLabelLen,
			totalLabelLen,
			len(txtSpecTimerLines.labelLeftMarginChars),
			len(startTimeLabelRunes),
			len(txtSpecTimerLines.endTimeLabel),
			len(txtSpecTimerLines.timeDurationLabel),
			len(txtSpecTimerLines.labelRightMarginChars),
			txtSpecTimerLines.textLabelFieldLen)

		return err
	}

	txtSpecTimerLines.startTimeLabel = startTimeLabelRunes

	if txtSpecTimerLines.textLabelFieldLen < lenLongestLabel {
		txtSpecTimerLines.textLabelFieldLen = lenLongestLabel
	}

	return err
}

// SetStartAndEndTime - Sets both the starting and ending times for
// the timing event described by this instance of
// TextLineSpecTimerLines.
//
// If either of the input parameters 'startTime' or 'endTime' is
// submitted as a zero time value, this method will return an
// error.
//
// If 'endTime' occurs BEFORE 'startTime', this method will return
// an error.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  startTime                  time.Time
//     - The time at which the time event starts. The timer event
//       refers to the subject timer event of the current
//       TextLineSpecTimerLines instance.
//
//
//  endTime                    time.Time
//     - The time at which the time event ends. The timer event
//       refers to the subject timer event of the current
//       TextLineSpecTimerLines instance.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings containing
//                      error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (txtSpecTimerLines *TextLineSpecTimerLines) SetStartAndEndTime(
	startTime time.Time,
	endTime time.Time,
	errorPrefix interface{}) error {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error = nil

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecTimerLines.SetStartAndEndTime()",
		"")

	if err != nil {
		return err
	}

	if startTime.IsZero() {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'startTime' is set to a ZERO value!\n",
			ePrefix.String())

		return err
	}

	if endTime.IsZero() {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'endTime' is set to a ZERO value!\n",
			ePrefix.String())

		return err
	}

	if endTime.Before(startTime) {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'endTime' occurs before\n"+
			"input parameter 'startTime'\n",
			ePrefix.String())

		return err
	}

	txtSpecTimerLines.startTime = startTime

	txtSpecTimerLines.endTime = endTime

	return err
}

// SetTextLabelFieldLength - Sets the Text Label Field Length. This
// value is stored in the internal member variable
// 'textLabelFieldLen' as an integer value.
//
// The Text Label Field Length specifies the length of the text
// field in which all three time event description labels, will
// be positioned and justified.
//
// The three description labels associated with the time event
// specified by this instance of TextLineSpecTimerLines are listed
// as follows:
//   (1) TextLineSpecTimerLines.startTimeLabel
//       - Describes the starting time for this timer event.
//
//   (2) TextLineSpecTimerLines.endTimeLabel
//       - Describes the ending time for this timer event.
//
//   (3) TextLineSpecTimerLines.timeDuration
//       - Describes the time duration or elapsed time for this
//         timer event.
//
// If Text Label Field Length ('fieldLength') is longer than the
// length of the longest of the three label strings, the Text
// Justification specification 'textLabelJustification' will be
// applied to position all three labels within a standard text
// field length. Text Justification can be set to 'Right', 'Left'
// or 'Center'.
//
// If Text Label Field Length ('fieldLength') is less than the
// length of the longest of the three label strings, it will be
// automatically set equal to the longest label length. Label
// strings with a length less than the maximum length will be
// justified according to the Text Justification Specification.
// Reference:
//    TextLineSpecTimerLines.SetLabelJustification()
//
// If input parameter 'fieldLength' is set to a value less than
// minus one (-1), an error will be generated.
//
// If input parameter 'fieldLength' is set to a value greater than
// one-million (1,000,000), an error will be generated.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  fieldLength               int
//     - The new value for the Text Label Field Length. If the
//       value of 'fieldLength' is less than minus one (-1) or
//       greater than one-million (1,000,000) an error will be
//       returned.
//
//       If this input value is less than the length of the
//       currently configured longest text label, Text Label Field
//       Length will be automatically set to the length of that
//       longest text label.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings containing
//                      error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (txtSpecTimerLines *TextLineSpecTimerLines) SetTextLabelFieldLength(
	fieldLength int,
	errorPrefix interface{}) (err error) {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecTimerLines."+
			"SetTextLabelFieldLength()",
		"")

	if err != nil {
		return err
	}

	lenLongestLabel := textLineSpecTimerLinesElectron{}.ptr().
		getLengthOfLongestLabel(
			txtSpecTimerLines.endTimeLabel,
			txtSpecTimerLines.startTimeLabel,
			txtSpecTimerLines.timeDurationLabel)

	if fieldLength < -1 {

		err = fmt.Errorf("%v\n"+
			"Input parameter 'fieldLength' is invalid.\n"+
			"The value of 'fieldLength' is less than minus one (-1)\n"+
			"fieldLength = '%v'\n",
			ePrefix.String(),
			fieldLength)

		return err
	}

	if fieldLength > 1000000 {

		err = fmt.Errorf("%v\n"+
			"Input parameter 'fieldLength' is invalid.\n"+
			"The value of 'fieldLength' is greater than one-million (1,000,000)\n"+
			"fieldLength = '%v'\n",
			ePrefix.String(),
			fieldLength)

		return err
	}

	if fieldLength < lenLongestLabel {

		txtSpecTimerLines.textLabelFieldLen = lenLongestLabel

	} else {

		txtSpecTimerLines.textLabelFieldLen = fieldLength

	}

	return err
}

// SetTimeDurationLabel - Sets the internal member variable
// 'timeDurationLabel'.
//
// 'timeDurationLabel' is a text label inserted in the
// TextLineSpecTimerLines output string to describe the elapsed
// time presentation. Time duration, or elapsed time, for this
// timer event is computed by subtracting the starting time from
// the ending time.
//
// If the input parameter 'timeDurationLabel' is submitted as an
// empty string, its value will be defaulted to the string
// "Elapsed Time".
//
// If the total character length of the new Time Duration Label,
// plus the lengths of the Start Time Label, End Time label, the
// Left Margin Characters and the Right Margin Characters, exceeds
// the maximum available label length (55-characters), an error
// will be returned.
//
// ----------------------------------------------------------------
//
// Be Advised
//
// If this method completes successfully and the text field length
// is less than the length of the longest label, the text field
// length will be reset to the length of the longest label.
//
// The length of the longest label is determined by calculating the
// character lengths of the three text labels: 'startTimeLabel',
// 'endTimeLabel' and 'timeDurationLabel'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  timeDurationLabel          string
//     - The text label used to describe the time duration or
//       elapsed time computed from the 'startTime' and 'endTime'
//       parameters.
//
//       If this string is submitted as a zero length or empty
//       string, 'timeDurationLabel' will be assigned a default
//       value of "Elapsed Time".
//
//       If 'timeDurationLabel' exceeds 50-characters in length,
//       this method will return an error.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings containing
//                      error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (txtSpecTimerLines *TextLineSpecTimerLines) SetTimeDurationLabel(
	timeDurationLabel string,
	errorPrefix interface{}) error {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error = nil

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecTimerLines."+
			"SetTimeDurationLabel()",
		"")

	if err != nil {
		return err
	}

	lengthOfNewTimeDurationLabel := len(timeDurationLabel)

	maxAllowableLabelLen := textLineSpecTimerLinesPreon{}.ptr().
		getMaximumTimerLabelLen()

	if lengthOfNewTimeDurationLabel > maxAllowableLabelLen {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'timeDurationLabel'\n"+
			"exceeds the maximum label string length!\n"+
			"The maximum label string length is %v-characters\n"+
			"The length of 'timeDurationLabel' is %v-characters\n",
			ePrefix.String(),
			maxAllowableLabelLen,
			lengthOfNewTimeDurationLabel)

		return err
	}

	if lengthOfNewTimeDurationLabel == 0 {
		txtSpecTimerLines.timeDurationLabel =
			textLineSpecTimerLinesElectron{}.ptr().
				getDefaultTimeDurationLabel()
	}

	timeDurationLabelRunes := []rune(timeDurationLabel)

	txtTimerLinesElectron := textLineSpecTimerLinesElectron{}

	lenLongestLabel := txtTimerLinesElectron.getLengthOfLongestLabel(
		txtSpecTimerLines.startTimeLabel,
		txtSpecTimerLines.endTimeLabel,
		timeDurationLabelRunes)

	totalLabelLen := txtTimerLinesElectron.
		getTotalLabelLength(
			txtSpecTimerLines.labelLeftMarginChars,
			txtSpecTimerLines.startTimeLabel,
			txtSpecTimerLines.endTimeLabel,
			timeDurationLabelRunes,
			txtSpecTimerLines.textLabelFieldLen,
			txtSpecTimerLines.labelRightMarginChars)

	if totalLabelLen > maxAllowableLabelLen {
		err = fmt.Errorf("%v\n"+
			"Error: The total length of the text label field is invalid!\n"+
			"The New 'Time Duration Label' will cause the entire text label\n"+
			"to exceed the maximum available text label field length.\n"+
			"The maximum text label field length is %v-characters\n"+
			"The total length of 'labelLeftMarginChars' plus 'labelRightMarginChars'"+
			"plus the the text label field length is %v-characters."+
			"'text label field length' is computed by taking the longest"+
			"text label length.\n"+
			"labelLeftMarginChars  = '%v'\n"+
			"startTimeLabel        = '%v'\n"+
			"endTimeLabel          = '%v'\n"+
			"New timeDurationLabel = '%v'\n"+
			"labelRightMarginChars = '%v'\n"+
			"textLabelFieldLen     = '%v'\n",
			ePrefix.String(),
			maxAllowableLabelLen,
			totalLabelLen,
			len(txtSpecTimerLines.labelLeftMarginChars),
			len(txtSpecTimerLines.startTimeLabel),
			len(txtSpecTimerLines.endTimeLabel),
			len(timeDurationLabelRunes),
			len(txtSpecTimerLines.labelRightMarginChars),
			txtSpecTimerLines.textLabelFieldLen)

		return err
	}

	txtSpecTimerLines.timeDurationLabel =
		timeDurationLabelRunes

	if txtSpecTimerLines.textLabelFieldLen < lenLongestLabel {
		txtSpecTimerLines.textLabelFieldLen = lenLongestLabel
	}

	return err
}

// SetTimeFormat - Sets the internal member variable
// 'timeFormat'. This time format specification is used to format
// all time values in text strings.
//
// The Go Programming language uses pattern formatting to format
// time values. Reference:
//   https://pkg.go.dev/time#Time.Format
//
//
// ----------------------------------------------------------------
//
// Input Parameters
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
// ------------------------------------------------------------------------
//
// Return Values
//
//  --- None ---
//
func (txtSpecTimerLines *TextLineSpecTimerLines) SetTimeFormat(
	timeFormat string) {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	if len(timeFormat) == 0 {

		txtSpecTimerLines.timeFormat =
			textSpecificationMolecule{}.ptr().
				getDefaultDateTimeFormat()

	} else {

		txtSpecTimerLines.timeFormat =
			timeFormat
	}

	return
}

// String - Returns the formatted text generated by this
// Text Line Specification for text display output and printing.
//
// The current 'TextLineSpecTimerLines' instance will provide the
// necessary information for generating text output which describes
// a timer event. The output will always consist of three lines of
// text specifying the event start time, end time and time duration
// or elapsed time.
//
// If an error occurs, the error message will be included in the
// returned string.
//
// This method is similar to
// TextLineSpecTimerLines.GetFormattedText() with the sole
// difference being that this method does not return an error.
//
// This method fulfills requirements of the ITextLineSpecification
// interface.
//
func (txtSpecTimerLines TextLineSpecTimerLines) String() string {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TextLineSpecTimerLines.GetFormattedText()",
		"")

	var formattedText string
	var err error

	formattedText,
		_,
		_,
		err = textLineSpecTimerLinesMolecule{}.ptr().
		getFormattedText(
			&txtSpecTimerLines,
			&ePrefix)

	if err != nil {
		formattedText = fmt.Sprintf("%v\n",
			err.Error())
	}

	return formattedText
}

// TextBuilder - Configures the lines of text produced by this
// instance of TextLineSpecTimerLines, and writes it to an instance
// of strings.Builder.
//
// This method fulfills requirements of the ITextLineSpecification
// interface.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  sBuilder                   *strings.Builder
//    - An instance of strings.Builder. The line of text produced
//      by the current instance of TextLineSpecTimerLines and writes
//      that text to 'sBuilder'.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings containing
//                      error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (txtSpecTimerLines *TextLineSpecTimerLines) TextBuilder(
	sBuilder *strings.Builder,
	errorPrefix interface{}) error {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecTimerLines.TextBuilder()",
		"")

	if err != nil {
		return err
	}

	if sBuilder == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'sBuilder' (strings.Builder)\n"+
			"is invalid! 'sBuilder' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	var formattedTxtStr string

	formattedTxtStr,
		_,
		_,
		err = textLineSpecTimerLinesMolecule{}.ptr().
		getFormattedText(
			txtSpecTimerLines,
			ePrefix)

	if err != nil {
		return err
	}

	var err2 error

	_,
		err2 = sBuilder.WriteString(formattedTxtStr)

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"Error returned by sBuilder.WriteString(formattedTxtStr)\n"+
			"%v\n",
			ePrefix.String(),
			err2.Error())
	}

	return err
}

// TextLineSpecName - returns a string specifying the name
// of this Text Line Specification.
//
// This method fulfills requirements of ITextSpecification
// interface.
//
func (txtSpecTimerLines *TextLineSpecTimerLines) TextLineSpecName() string {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	return "TimerLines"
}

// TextTypeName - returns a string specifying the type
// of Text Line specification.
//
// This method fulfills requirements of ITextSpecification interface.
//
func (txtSpecTimerLines *TextLineSpecTimerLines) TextTypeName() string {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	return "TextLineSpecTimerLines"

}
