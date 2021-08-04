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
	startTimeLabel             []rune
	startTime                  time.Time
	endTimeLabel               []rune
	endTime                    time.Time
	timeFormat                 string
	timeDurationLabel          []rune
	labelFieldLen              int
	labelJustification         TextJustify
	labelOutputSeparationChars []rune
	lock                       *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming instance of
// TextLineSpecTimerLines ('incomingTimerLines') to the data fields
// of the current TextLineSpecTimerLines instance
// ('txtSpecTimerLines').
//
// IMPORTANT
// All of the data fields in current TextLineSpecTimerLines
// instance ('txtSpecTimerLines') will be modified and overwritten.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingTimerLines         *TextLineSpecTimerLines
//     - A pointer to an instance of TextLineSpecTimerLines. This
//       method will NOT change the values of internal member
//       variables contained in this instance.
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

// GetFormattedText - Returns the formatted text generated by this
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
func (txtSpecTimerLines *TextLineSpecTimerLines) String() string {

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
