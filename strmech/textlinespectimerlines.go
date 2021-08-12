package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
	"time"
)

// TextLineSpecTimerLines - Implements a time which produces three
// text lines of output for display or printing.
//
// The first line of text shows the Starting Time. The second line
// shows the Ending Time and the third line displays the time
// duration or the difference between starting time and ending
// time.
//
// The third line displaying time duration will format time
// duration in days, hours, minutes, seconds, microseconds,
// milliseconds and nanoseconds.
//
type TextLineSpecTimerLines struct {
	labelLeftMarginChars  []rune
	startTimeLabel        []rune
	startTime             time.Time
	endTimeLabel          []rune
	endTime               time.Time
	timeFormat            string
	timeDurationLabel     []rune
	labelFieldLen         int
	labelJustification    TextJustify
	labelRightMarginChars []rune
	lock                  *sync.Mutex
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
		ePrefix.XCtx("txtSpecTimerLines->"))
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
			ePrefix.XCtx(
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
			ePrefix.XCtx(
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
// This method is identical in function to
// TextLineSpecTimerLines.String()
//
// This method fulfills requirements of the ITextLineSpecification
// interface.
//
func (txtSpecTimerLines *TextLineSpecTimerLines) GetFormattedText() string {

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
		err = textLineSpecTimerLinesMolecule{}.ptr().
		getFormattedText(
			txtSpecTimerLines,
			&ePrefix)

	if err != nil {
		formattedText = fmt.Sprintf("%v\n",
			err.Error())
	}

	return formattedText
}

// GetEndTimeLabel - Returns the internal member variable
// 'endTimeLabel' as a rune array.
//
// 'endTimeLabel' is a text label inserted in the output
// string to describe the ending time presentation.
//
func (txtSpecTimerLines *TextLineSpecTimerLines) GetEndTimeLabel() []rune {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	return txtSpecTimerLines.endTimeLabel
}

// GetEndTime - Returns the internal member variable 'endTime'
// as time.Time value.
//
// The ending time is the end of the timing event described by this
// instance of TextLineSpecTimerLines.
//
func (txtSpecTimerLines *TextLineSpecTimerLines) GetEndTime() time.Time {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	return txtSpecTimerLines.endTime
}

// GetStartTimeLabel - Returns the internal member variable
// 'startTimeLabel' as a rune array.
//
// 'startTimeLabel' is a text label inserted in the output
// string to describe the starting time presentation.
//
func (txtSpecTimerLines *TextLineSpecTimerLines) GetStartTimeLabel() []rune {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	return txtSpecTimerLines.startTimeLabel
}

// GetStartTime - Returns the internal member variable 'startTime'
// as time.Time value.
//
// The starting time is the start or beginning of the timing event
// described by this instance of TextLineSpecTimerLines.
//
func (txtSpecTimerLines *TextLineSpecTimerLines) GetStartTime() time.Time {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	return txtSpecTimerLines.startTime
}

// GetLabelFieldLength - Returns the internal member variable
// 'labelFieldLen' as an integer value.
//
// The Label Field Length specifies the length of the text field in
// which all three labels, 'startTimeLabel', 'endTimeLabel' and
// 'timeDurationLabel', will be positioned and justified.
//
// If Label Field Length is longer than the length of the longest
// label string, the Text Justification specification
// 'labelJustification' will be applied to position all three
// labels in their text fields. Text Justification can be set to
// 'Right', 'Left' or 'Center'.
//
func (txtSpecTimerLines *TextLineSpecTimerLines) GetLabelFieldLength() int {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	return txtSpecTimerLines.labelFieldLen
}

// GetLabelJustification - Returns the internal member variable
// 'labelJustification'.
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

	return txtSpecTimerLines.labelJustification
}

// GetLabelOutputSeparationChars - Returns the internal member
// variable, 'labelRightMarginChars', as an array of runes.
//
// The 'labelRightMarginChars' is one or more characters used
// to separate the text labels, 'startTime', 'endTime' and
// 'timeDuration' from their respective output values. Usually,
// 'labelRightMarginChars' is set to either a single white
// space character or the default value which is a colon and a
// white space character (": ").
//
func (txtSpecTimerLines *TextLineSpecTimerLines) GetLabelOutputSeparationChars() []rune {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	return txtSpecTimerLines.labelRightMarginChars
}

// GetTimeDurationLabel - Returns the internal member variable
// 'timeDurationLabel' as an array of runes.
//
// 'timeDurationLabel' is a text label inserted in the
// TextLineSpecTimerLines output string to describe the elapsed
// time presentation. Time duration, or elapsed time, for this
// timer event is computed by subtracting the starting time from
// the ending time.
//
func (txtSpecTimerLines *TextLineSpecTimerLines) GetTimeDurationLabel() []rune {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	return txtSpecTimerLines.timeDurationLabel
}

// GetTimeFormat - Returns the internal member variable
// 'timeFormat'. This time format specification is used to format
// all time values in text strings.
//
// The Go Programming language uses pattern formatting to format
// time values. Reference:
//   https://pkg.go.dev/time#Time.Format
//
func (txtSpecTimerLines *TextLineSpecTimerLines) GetTimeFormat() string {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	return txtSpecTimerLines.timeFormat
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
			ePrefix.XCtx("txtSpecTimerLines"))

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
//  startTimeLabel      - Defaults to "Start Time"
//
//  endTimeLabel        - Defaults to "End Time".
//
//  timeFormat          - Defaults to "2006-01-02 15:04:05.000000000 -0700 MST"
//
//  timeDurationLabel   - Defaults to "Elapsed Time"
//
//  labelFieldLen       - Defaults to  '12'
//
//  labelJustification  - Defaults to TextJustify(0).Right()
//
//  labelRightMarginChars
//                      - Defaults to ": "
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
//  TextLineSpecTimerLines
//     - If this method completes successfully, it will create and
//       return a new instance of TextLineSpecTimerLines which is
//       fully configured with all the parameters necessary to
//       format a complete timer event for text display output or
//       printing.
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
	TextLineSpecTimerLines,
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
		return newTxtTimerLines, err
	}

	timeDurationLabel :=
		textLineSpecTimerLinesElectron{}.ptr().
			getDefaultTimeDurationLabel()

	err = textLineSpecTimerLinesMolecule{}.ptr().
		setTxtLineSpecTimerLines(
			&newTxtTimerLines,
			nil,
			startTime,
			nil,
			endTime,
			"",
			timeDurationLabel,
			len(timeDurationLabel),
			TxtJustify.Right(),
			nil,
			ePrefix.XCtx("newTxtTimerLines"))

	return newTxtTimerLines, err
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
// This method will automatically set the following default values:
//
// ------------------------------------------------------------------------
//
// Default Values
//
//  startTime           - Defaults to July 4, 1776 9:30AM UTC
//
//  startTimeLabel      - Defaults to "Start Time"
//
//  endTimeLabel        - Defaults to "End Time".
//
//  endTime             - Defaults to July 4, 1776 9:30AM UTC
//
//  timeFormat          - Defaults to "2006-01-02 15:04:05.000000000 -0700 MST"
//
//  timeDurationLabel   - Defaults to "Elapsed Time"
//
//  labelFieldLen       - Defaults to  '12'
//
//  labelJustification  - Defaults to TextJustify(0).Right()
//
//  labelRightMarginChars
//                      - Defaults to ": "
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
//  TextLineSpecTimerLines
//     - This method will create and return a new instance of
//       TextLineSpecTimerLines which is fully configured except
//       for the starting time and ending time.
//
func (txtSpecTimerLines TextLineSpecTimerLines) NewDefaultShellTimerEvent() TextLineSpecTimerLines {

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
			defaultTime,
			nil,
			defaultTime,
			"",
			timeDurationLabel,
			len(timeDurationLabel),
			TxtJustify.Right(),
			nil,
			nil)

	return newTxtTimerLines
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
// Use this method when both start time and ending time for the
// timer event are known quantities.
//
//
// ------------------------------------------------------------------------
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
//       If 'startTimeLabel' exceeds 50-characters in length, this
//       method will return an error.
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
//       If 'endTimeLabel' exceeds 50-characters in length, this
//       method will return an error.
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
//       If 'timeDurationLabel' exceeds 50-characters in length,
//       this method will return an error.
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
//       three text labels 'startTimeLabel', 'endTimeLabel' and
//       'timeDurationLabel' within the field length specified by
//       'labelFieldLen'.
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
//       If this string is submitted as a zero length or empty
//       string, 'labelRightMarginChars' will be assigned a
//       default value of  ": ". Example Output:
//        Start Time: 2010-01-02 15:04:05.000000000 -0700 MST
//
//       If the string length of 'labelRightMarginChars'
//       exceeds the maximum length of 5-characters, this method
//       will return an error.
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
//  TextLineSpecTimerLines
//     - If this method completes successfully, it will create and
//       return a new instance of TextLineSpecTimerLines which is
//       fully configured with all the parameters necessary to
//       format a complete timer event for text display output or
//       printing.
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
	startTimeLabel string,
	startTime time.Time,
	endTimeLabel string,
	endTime time.Time,
	timeFormat string,
	timeDurationLabel string,
	labelFieldLen int,
	labelJustification TextJustify,
	labelOutputSeparationChars string,
	errorPrefix interface{}) (
	TextLineSpecTimerLines,
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
		return newTxtTimerLines, err
	}

	err = textLineSpecTimerLinesMolecule{}.ptr().
		setTxtLineSpecTimerLines(
			&newTxtTimerLines,
			[]rune(startTimeLabel),
			startTime,
			[]rune(endTimeLabel),
			endTime,
			timeFormat,
			[]rune(timeDurationLabel),
			labelFieldLen,
			labelJustification,
			[]rune(labelOutputSeparationChars),
			ePrefix.XCtx("newTxtTimerLines"))

	return newTxtTimerLines, err
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
// This method will automatically set the following default values:
//
// ------------------------------------------------------------------------
//
// Default Values
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
//  startTimeLabel             string
//     - A string containing the text characters constituting the
//       starting time text label.
//
//       If this string is submitted as a zero length or empty
//       string, 'startTimeLabel' will be assigned a default value
//       of "Start Time".
//
//       If 'startTimeLabel' exceeds 50-characters in length, this
//       method will return an error.
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
//       If 'endTimeLabel' exceeds 50-characters in length, this
//       method will return an error.
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
//       If 'timeDurationLabel' exceeds 50-characters in length,
//       this method will return an error.
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
//       three text labels 'startTimeLabel', 'endTimeLabel' and
//       'timeDurationLabel' within the field length specified by
//       'labelFieldLen'.
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
//       If the string length of 'labelRightMarginChars'
//       exceeds the maximum length of 5-characters, this method
//       will return an error.
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
//  TextLineSpecTimerLines
//     - If this method completes successfully, it will create and
//       return a new instance of TextLineSpecTimerLines which is
//       fully configured except for the starting time and ending
//       time.
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
	startTimeLabel string,
	endTimeLabel string,
	timeFormat string,
	timeDurationLabel string,
	labelFieldLen int,
	labelJustification TextJustify,
	labelOutputSeparationChars string,
	errorPrefix interface{}) (
	TextLineSpecTimerLines,
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
		return newTxtTimerLines, err
	}

	defaultTime :=
		textLineSpecTimerLinesElectron{}.ptr().
			getDefaultTime()

	err = textLineSpecTimerLinesMolecule{}.ptr().
		setTxtLineSpecTimerLines(
			&newTxtTimerLines,
			[]rune(startTimeLabel),
			defaultTime,
			[]rune(endTimeLabel),
			defaultTime,
			timeFormat,
			[]rune(timeDurationLabel),
			labelFieldLen,
			labelJustification,
			[]rune(labelOutputSeparationChars),
			ePrefix.XCtx("newTxtTimerLines"))

	return newTxtTimerLines, err
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
// This method will automatically set the following default values:
//
// ------------------------------------------------------------------------
//
// Default Values
//
//  startTimeLabel      - Defaults to "Start Time"
//
//  endTimeLabel        - Defaults to "End Time".
//
//  timeFormat          - Defaults to "2006-01-02 15:04:05.000000000 -0700 MST"
//
//  timeDurationLabel   - Defaults to "Elapsed Time"
//
//  labelFieldLen       - Defaults to  '12'
//
//  labelJustification  - Defaults to TextJustify(0).Right()
//
//  labelRightMarginChars
//                      - Defaults to ": "
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
		"TextLineSpecTimerLines.SetDefaultFullTimerEvent()",
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
			startTime,
			nil,
			endTime,
			"",
			timeDurationLabel,
			len(timeDurationLabel),
			TxtJustify.Right(),
			nil,
			ePrefix.XCtx("txtSpecTimerLines"))

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
//
// This method will automatically set the following default values:
//
// ------------------------------------------------------------------------
//
// Default Values
//
//  startTime           - Defaults to July 4, 1776 9:30AM UTC
//
//  startTimeLabel      - Defaults to "Start Time"
//
//  endTime             - Defaults to July 4, 1776 9:30AM UTC
//
//  endTimeLabel        - Defaults to "End Time".
//
//  timeFormat          - Defaults to "2006-01-02 15:04:05.000000000 -0700 MST"
//
//  timeDurationLabel   - Defaults to "Elapsed Time"
//
//  labelFieldLen       - Defaults to  '12'
//
//  labelJustification  - Defaults to TextJustify(0).Right()
//
//  labelRightMarginChars
//                      - Defaults to ": "
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
// string, its value will be defaulted to "End Time".
//
// If the length of input parameter 'endTimeLabel' is greater than
// fifty (50) characters, this method will return an error.
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
		"TextLineSpecTimerLines.SetEndTimeLabel()",
		"")

	if err != nil {
		return err
	}

	lenLabel := len(endTimeLabel)

	maxLabelLen := textLineSpecTimerLinesElectron{}.ptr().
		getMaximumTimerLabelLen()

	if lenLabel > maxLabelLen {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'endTimeLabel'\n"+
			"exceeds the maximum label string length!\n"+
			"The maximum label string length is %v-characters\n"+
			"The length of 'endTimeLabel' is %v-characters\n",
			ePrefix.String(),
			maxLabelLen,
			lenLabel)

		return err
	}

	if len(endTimeLabel) == 0 {

		txtSpecTimerLines.endTimeLabel =
			textLineSpecTimerLinesElectron{}.ptr().
				getDefaultEndTimeLabel()

	} else {

		txtSpecTimerLines.endTimeLabel = []rune(endTimeLabel)
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
// information for text display output or printing.
//
// Use this method when both start time and ending time for the
// timer event are known quantities.
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
//  startTimeLabel             string
//     - A string containing the text characters constituting the
//       starting time text label.
//
//       If this string is submitted as a zero length or empty
//       string, 'startTimeLabel' will be assigned a default value
//       of "Start Time".
//
//       If 'startTimeLabel' exceeds 50-characters in length, this
//       method will return an error.
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
//       If 'endTimeLabel' exceeds 50-characters in length, this
//       method will return an error.
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
//       If 'timeDurationLabel' exceeds 50-characters in length,
//       this method will return an error.
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
//       three text labels 'startTimeLabel', 'endTimeLabel' and
//       'timeDurationLabel' within the field length specified by
//       'labelFieldLen'.
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
//       If this string is submitted as a zero length or empty
//       string, 'labelRightMarginChars' will be assigned a
//       default value of  ": ". Example Output:
//        Start Time: 2010-01-02 15:04:05.000000000 -0700 MST
//
//       If the string length of 'labelRightMarginChars'
//       exceeds the maximum length of 5-characters, this method
//       will return an error.
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
	startTimeLabel string,
	startTime time.Time,
	endTimeLabel string,
	endTime time.Time,
	timeFormat string,
	timeDurationLabel string,
	labelFieldLen int,
	labelJustification TextJustify,
	labelOutputSeparationChars string,
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
			[]rune(startTimeLabel),
			startTime,
			[]rune(endTimeLabel),
			endTime,
			timeFormat,
			[]rune(timeDurationLabel),
			labelFieldLen,
			labelJustification,
			[]rune(labelOutputSeparationChars),
			ePrefix.XCtx("txtSpecTimerLines"))

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
// This method will automatically set the following default values:
//
// ------------------------------------------------------------------------
//
// Default Values
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
//  startTimeLabel             string
//     - A string containing the text characters constituting the
//       starting time text label.
//
//       If this string is submitted as a zero length or empty
//       string, 'startTimeLabel' will be assigned a default value
//       of "Start Time".
//
//       If 'startTimeLabel' exceeds 50-characters in length, this
//       method will return an error.
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
//       If 'endTimeLabel' exceeds 50-characters in length, this
//       method will return an error.
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
//       If 'timeDurationLabel' exceeds 50-characters in length,
//       this method will return an error.
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
//       three text labels 'startTimeLabel', 'endTimeLabel' and
//       'timeDurationLabel' within the field length specified by
//       'labelFieldLen'.
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
//       If the string length of 'labelRightMarginChars'
//       exceeds the maximum length of 5-characters, this method
//       will return an error.
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
	startTimeLabel string,
	endTimeLabel string,
	timeFormat string,
	timeDurationLabel string,
	labelFieldLen int,
	labelJustification TextJustify,
	labelOutputSeparationChars string,
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
		"TextLineSpecTimerLines.SetShellTimerEvent()",
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
			[]rune(startTimeLabel),
			defaultTime,
			[]rune(endTimeLabel),
			defaultTime,
			timeFormat,
			[]rune(timeDurationLabel),
			labelFieldLen,
			labelJustification,
			[]rune(labelOutputSeparationChars),
			ePrefix.XCtx("newTxtTimerLines"))

	return err
}

// SetStartTimeLabel - Sets the internal member variable
// 'startTimeLabel'.
//
// 'startTimeLabel' is a text label inserted in the output
// string to describe the starting time presentation.
//
// If input parameter 'startTimeLabel' is submitted as an empty
// string, its value will be defaulted to "Start Time".
//
// If the length of input parameter 'startTimeLabel' is greater
// than fifty (50) characters, this method will return an error.
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
		"TextLineSpecTimerLines.SetEndTimeLabel()",
		"")

	if err != nil {
		return err
	}

	lenLabel := len(startTimeLabel)

	maxLabelLen := textLineSpecTimerLinesElectron{}.ptr().
		getMaximumTimerLabelLen()

	if lenLabel > maxLabelLen {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'startTimeLabel'\n"+
			"exceeds the maximum label string length!\n"+
			"The maximum label string length is %v-characters\n"+
			"The length of 'startTimeLabel' is %v-characters\n",
			ePrefix.String(),
			maxLabelLen,
			lenLabel)

		return err
	}

	if lenLabel == 0 {
		txtSpecTimerLines.startTimeLabel =
			textLineSpecTimerLinesElectron{}.ptr().
				getDefaultStartTimeLabel()
	} else {
		txtSpecTimerLines.startTimeLabel =
			[]rune(startTimeLabel)
	}

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
		"TextLineSpecTimerLines.SetStartTime()",
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

// SetLabelFieldLength - Sets the internal member variable
// 'labelFieldLen' as an integer value.
//
// The Label Field Length specifies the length of the text field in
// which all three time event description labels, will be
// positioned and justified.
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
// If Label Field Length ('labelFieldLength') is longer than the
// length of the longest of the three label strings, the Text
// Justification specification 'labelJustification' will be applied
// to position all three labels in their text fields. Text
// Justification can be set to 'Right', 'Left' or 'Center'.
//
// If Label Field Length ('labelFieldLength') is less than the
// length of the longest of the three label strings, it will be set
// equal to that maximum length. Label strings with a length less
// than the maximum length will be justified according to the
// Text Justification Specification. Reference:
//    TextLineSpecTimerLines.SetLabelJustification()
//
// If input parameter 'labelFieldLength' is set to a value less
// than minus one (-1), it will be defaulted to a value of minus
// one (-1).
//
// If input parameter 'labelFieldLength' is set to a value greater
// than one-million (1,000,000), it will be defaulted to a value of
// minus one (-1).
//
func (txtSpecTimerLines *TextLineSpecTimerLines) SetLabelFieldLength(
	labelFieldLength int) {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	if labelFieldLength < -1 {
		labelFieldLength = -1
	}

	if labelFieldLength > 1000000 {
		labelFieldLength = -1
	}

	txtSpecTimerLines.labelFieldLen = labelFieldLength

	return
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
//  labelJustification         TextJustify
//     - An enumeration which specifies the justification of the
//       three text labels 'startTimeLabel', 'endTimeLabel' and
//       'timeDurationLabel' within the field length specified by
//       'labelFieldLen'.
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
//       ('labelJustification') is set to a value other than 'Left',
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
			"Error: Input parameter 'labelJustification' is invalid!\n"+
			"'labelJustification' must be set to 'Left', 'Right' or 'Center'\n"+
			"labelJustification string = '%v'\n"+
			"labelJustification int = '%v'\n",
			ePrefix.String(),
			labelJustification.String(),
			labelJustification.XValueInt())

		return err
	}

	txtSpecTimerLines.labelJustification =
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
// If the string length of 'labelRightMarginChars' exceeds the
// maximum length of 5-characters, this method will return an
// error.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
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
//       If the string length of 'labelRightMarginChars'
//       exceeds the maximum length of 5-characters, this method
//       will return an error.
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
	labelOutputSeparationChars string,
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

	lenLabel := len(labelOutputSeparationChars)

	maxLabelLen := textLineSpecTimerLinesElectron{}.ptr().
		getMaximumLabelOutputSeparationCharsLen()

	if lenLabel > maxLabelLen {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'labelRightMarginChars'\n"+
			"exceeds the maximum label string length!\n"+
			"The maximum label string length is %v-characters\n"+
			"The length of 'labelRightMarginChars' is %v-characters\n",
			ePrefix.String(),
			maxLabelLen,
			lenLabel)

		return err
	}

	if len(labelOutputSeparationChars) == 0 {

		txtSpecTimerLines.labelRightMarginChars =
			textLineSpecTimerLinesElectron{}.ptr().
				getDefaultLabelRightMarginChars()
	} else {

		txtSpecTimerLines.labelRightMarginChars =
			[]rune(labelOutputSeparationChars)
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
// empty string, its value will be defaulted to "Elapsed Time".
//
// If the length of input parameter 'timeDurationLabel' is greater
// than fifty (50) characters, this method will return an error.
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
		"TextLineSpecTimerLines.SetTimeDurationLabel()",
		"")

	if err != nil {
		return err
	}

	lenLabel := len(timeDurationLabel)

	maxLabelLen := textLineSpecTimerLinesElectron{}.ptr().
		getMaximumTimerLabelLen()

	if lenLabel > maxLabelLen {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'timeDurationLabel'\n"+
			"exceeds the maximum label string length!\n"+
			"The maximum label string length is %v-characters\n"+
			"The length of 'timeDurationLabel' is %v-characters\n",
			ePrefix.String(),
			maxLabelLen,
			lenLabel)

		return err
	}

	if lenLabel == 0 {
		txtSpecTimerLines.timeDurationLabel =
			textLineSpecTimerLinesElectron{}.ptr().
				getDefaultTimeDurationLabel()
	} else {
		txtSpecTimerLines.timeDurationLabel =
			[]rune(timeDurationLabel)
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
				getDefaultTimeFormat()

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
// This method is identical in function to
// TextLineSpecTimerLines.GetFormattedText()
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

// TextLineSpecName - returns a string specifying the name
// of this Text Line Specification.
//
// This method fulfills requirements of ITextSpecification
// interface.
//
func (txtSpecTimerLines TextLineSpecTimerLines) TextLineSpecName() string {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	return "TextLineSpecTimerLines"
}

// TextTypeName - returns a string specifying the type
// of Text Line specification.
//
// This method fulfills requirements of ITextSpecification interface.
//
func (txtSpecTimerLines TextLineSpecTimerLines) TextTypeName() string {

	if txtSpecTimerLines.lock == nil {
		txtSpecTimerLines.lock = new(sync.Mutex)
	}

	txtSpecTimerLines.lock.Lock()

	defer txtSpecTimerLines.lock.Unlock()

	return "TextLineSpecTimerLines"

}
