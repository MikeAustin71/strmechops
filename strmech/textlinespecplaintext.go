package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// TextLineSpecPlainText - Used to create a line of text from a
// simple text string.
//
// The TextLineSpecPlainText type provides formatting
// specifications for a simple line of text characters for text
// display, file output or printing.
//
// The plain text line consists of a left margin, the text string,
// the right margin and a line termination character or characters.
// The line termination character is usually a new line character
// ('\n').
//
// Left and right margins consist of zero or more characters
// customized and provided by the calling function.
//
// IMPORTANT
//
// ----------------------------------------------------------------
//
// By default, the TextLineSpecPlainText type will generate a
// line of text with a "new line" ('\n') line termination
// character. The application of this "new line" ('\n') character
// is automatic and by default.
//
// However, users may customize and override the application of
// line termination characters through the following methods:
//
//  TextLineSpecPlainText.SetLineTerminationChars()
//  TextLineSpecPlainText.TurnAutoLineTerminationOff()
//  TextLineSpecPlainText.TurnAutoLineTerminationOn()
//
//
// The TextLineSpecPlainText type implements the
// ITextLineSpecification interface.
//
//
// Member Variables
//
// ----------------------------------------------------------------
//
//  leftMarginChars            []rune
//     - An array of runes containing the text characters which
//       will be used as the left margin of the formatted text
//       string. The number of left margin characters is restricted
//       to a range from zero (0) to one-million (1,000,000),
//       inclusive.
//
//
//  rightMarginChars           []rune
//     - An array of runes containing the text characters which
//       will be used as the right margin of the formatted text
//       string. The number of right margin characters is
//       restricted to a range from zero (0) to one-million
//       (1,000,000), inclusive.
//
//
//  textString                 string
//     - A string of text characters supplied by the user which
//       will be positioned between the left and right margins when
//       constructing the formatted line of text generated for text
//       display, file output or printing.
//
//
//  newLineChars               []rune
//     - By default, each line of text generated by
//       TextLineSpecPlainText will be terminated with a new line
//       character ('\n'). However, users have the option to
//       override and modify this behavior by supplying an
//       alternative character or characters to be used as a line
//       termination sequence for each line of text produced by the
//       current TextLineSpecPlainText instance.
//
//       To override, change or control the behavior of
//       'newLineChars', see the following methods:
//         TextLineSpecPlainText.GetLineTerminationRunes()
//         TextLineSpecPlainText.SetLineTerminationChars()
//         TextLineSpecPlainText.SetLineTerminationRunes()
//         TextLineSpecPlainText.TurnAutoLineTerminationOff()
//
//
type TextLineSpecPlainText struct {
	leftMarginChars       []rune
	rightMarginChars      []rune
	textString            string
	turnLineTerminatorOff bool
	newLineChars          []rune
	lock                  *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming instance of
// TextLineSpecPlainText ('incomingPlainTxtLine') to the data
// fields of the current TextLineSpecPlainText instance
// ('plainTextLine').
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All the data fields in current TextLineSpecPlainText instance
// ('plainTextLine') will be deleted and overwritten.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingPlainTxtLine       *TextLineSpecPlainText
//     - A pointer to an instance of TextLineSpecPlainText. This
//       method will NOT change the data values of member variables
//       contained in this instance.
//
//       All data values in this TextLineSpecPlainText instance
//       will be copied to the current TextLineSpecPlainText
//       instance ('plainTextLine').
//
//       If 'incomingPlainTextLine' contains invalid member
//       variable data values, this method will return an error.
//
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
func (plainTextLine *TextLineSpecPlainText) CopyIn(
	incomingPlainTxtLine *TextLineSpecPlainText,
	errorPrefix interface{}) error {

	if plainTextLine.lock == nil {
		plainTextLine.lock = new(sync.Mutex)
	}

	plainTextLine.lock.Lock()

	defer plainTextLine.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecPlainText.CopyIn()",
		"")

	if err != nil {
		return err
	}

	err = textLineSpecPlainTextNanobot{}.ptr().
		copyIn(
			plainTextLine,
			incomingPlainTxtLine,
			ePrefix)

	return err
}

// CopyOut - Returns a deep copy of the current
// TextLineSpecPlainText instance.
//
// If the current TextLineSpecPlainText instance contains invalid
// member variables, this method will return an error.
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
//  TextLineSpecPlainText
//     - If this method completes successfully and no errors are
//       encountered, this parameter will return a deep copy of the
//       current TextLineSpecPlainText instance.
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
func (plainTextLine *TextLineSpecPlainText) CopyOut(
	errorPrefix interface{}) (
	TextLineSpecPlainText,
	error) {

	if plainTextLine.lock == nil {
		plainTextLine.lock = new(sync.Mutex)
	}

	plainTextLine.lock.Lock()

	defer plainTextLine.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecPlainText.CopyOut()",
		"")

	if err != nil {
		return TextLineSpecPlainText{}, err
	}

	var newPlainTxtLine TextLineSpecPlainText

	newPlainTxtLine,
		err = textLineSpecPlainTextNanobot{}.ptr().
		copyOut(
			plainTextLine,
			ePrefix.XCtx(
				"plainTextLine"))

	return newPlainTxtLine, err
}

// CopyOutITextLine - Returns a deep copy of the current
// TextLineSpecPlainText instance cast as a type
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
//       current TextLineSpecPlainText instance cast as an
//       ITextLineSpecification object.
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
func (plainTextLine *TextLineSpecPlainText) CopyOutITextLine(
	errorPrefix interface{}) (
	ITextLineSpecification,
	error) {

	if plainTextLine.lock == nil {
		plainTextLine.lock = new(sync.Mutex)
	}

	plainTextLine.lock.Lock()

	defer plainTextLine.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecPlainText.CopyOutITextLine()",
		"")

	if err != nil {
		return ITextLineSpecification(&TextLineSpecPlainText{}), err
	}

	var newPlainTxtLine TextLineSpecPlainText

	newPlainTxtLine,
		err = textLineSpecPlainTextNanobot{}.ptr().
		copyOut(
			plainTextLine,
			ePrefix.XCtx(
				"plainTextLine"))

	return ITextLineSpecification(&newPlainTxtLine), err
}

// CopyOutPtr - Returns a pointer to a deep copy of the current
// TextLineSpecPlainText instance.
//
// If the current TextLineSpecPlainText instance contains invalid
// member variables, this method will return an error.
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
//  *TextLineSpecPlainText
//     - If this method completes successfully and no errors are
//       encountered, this parameter will return a pointer to a
//       deep copy of the current TextLineSpecPlainText
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
func (plainTextLine *TextLineSpecPlainText) CopyOutPtr(
	errorPrefix interface{}) (
	*TextLineSpecPlainText,
	error) {

	if plainTextLine.lock == nil {
		plainTextLine.lock = new(sync.Mutex)
	}

	plainTextLine.lock.Lock()

	defer plainTextLine.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecPlainText.CopyOutPtr()",
		"")

	if err != nil {
		return &TextLineSpecPlainText{}, err
	}

	var newPlainTxtLine TextLineSpecPlainText

	newPlainTxtLine,
		err = textLineSpecPlainTextNanobot{}.ptr().
		copyOut(
			plainTextLine,
			ePrefix.XCtx(
				"plainTextLine"))

	return &newPlainTxtLine, err
}

// Empty - Deletes and resets all the member variable data values
// contained in the current TextLineSpecPlainText instance to their
// initial or zero values.
//
// After calling 'Empty', the caller CAN NOT reuse this instance of
// TextLineSpecStandardLine. All the internal member variables
// will have to be re-initialized.
//
// This method fulfills requirements of the ITextLineSpecification
// interface.
//
func (plainTextLine *TextLineSpecPlainText) Empty() {

	if plainTextLine.lock == nil {
		plainTextLine.lock = new(sync.Mutex)
	}

	plainTextLine.lock.Lock()

	textLineSpecPlainTextElectron{}.ptr().
		empty(plainTextLine)

	plainTextLine.lock.Unlock()

	plainTextLine.lock = nil
}

// Equal - Receives a pointer to another instance of
// TextLineSpecPlainText and proceeds to compare the member
// variables to those of the current TextLineSpecPlainText
// instance in order to determine if they are equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables of both instances are equal in
// all respects, this flag is set to 'true'. Otherwise, this method
// returns 'false'.
//
func (plainTextLine *TextLineSpecPlainText) Equal(
	incomingPlainTxtLine *TextLineSpecPlainText) bool {

	if plainTextLine.lock == nil {
		plainTextLine.lock = new(sync.Mutex)
	}

	plainTextLine.lock.Lock()

	defer plainTextLine.lock.Unlock()

	return textLineSpecPlainTextElectron{}.ptr().
		equal(
			plainTextLine,
			incomingPlainTxtLine)
}

// EqualITextLine - Receives an object implementing the
// ITextLineSpecification interface and proceeds to compare
// the member variables to those of the current
// TextLineSpecPlainText instance in order to determine if
// they are equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables from both instances are equal
// in all respects, this flag is set to 'true'. Otherwise, this
// method returns 'false'.
//
// This method is required by interface ITextLineSpecification.
//
func (plainTextLine *TextLineSpecPlainText) EqualITextLine(
	iTextLine ITextLineSpecification) bool {

	if plainTextLine.lock == nil {
		plainTextLine.lock = new(sync.Mutex)
	}

	plainTextLine.lock.Lock()

	defer plainTextLine.lock.Unlock()

	plainTxtLineTwo, ok :=
		iTextLine.(*TextLineSpecPlainText)

	if !ok {
		return false
	}

	return textLineSpecPlainTextElectron{}.ptr().
		equal(
			plainTextLine,
			plainTxtLineTwo)
}

// GetFormattedText - Returns the formatted text generated by this
// Text Line Specification for text display, file output and
// printing.
//
// If an error occurs, the error message will be included in the
// returned string.
//
// This method is identical in function to
// TextLineSpecPlainText.String()
//
// This method fulfills requirements of the ITextLineSpecification
// interface.
//
func (plainTextLine *TextLineSpecPlainText) GetFormattedText() string {

	if plainTextLine.lock == nil {
		plainTextLine.lock = new(sync.Mutex)
	}

	plainTextLine.lock.Lock()

	defer plainTextLine.lock.Unlock()

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TextLineSpecPlainText.GetFormattedText()",
		"")

	formattedTxtStr,
		err := textLineSpecPlainTextNanobot{}.ptr().
		getFormattedText(
			plainTextLine,
			ePrefix.XCtx("plainTextLine"))

	if err != nil {
		formattedTxtStr = fmt.Sprintf("%v\n",
			err.Error())
	}

	return formattedTxtStr
}

// IsValidInstanceError - Performs a diagnostic review of the data
// values encapsulated in the current TextLineSpecPlainText
// instance to determine if they are valid.
//
// If any data element evaluates as invalid, this method will
// return an error.
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
//       the current instance of TextLineSpecPlainText are found
//       to be invalid, this method will return an error.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' (error prefix) will be inserted or
//       prefixed at the beginning of the error message.
//
func (plainTextLine *TextLineSpecPlainText) IsValidInstanceError(
	errorPrefix interface{}) error {

	if plainTextLine.lock == nil {
		plainTextLine.lock = new(sync.Mutex)
	}

	plainTextLine.lock.Lock()

	defer plainTextLine.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecPlainText."+
			"IsValidInstanceError()",
		"")

	if err != nil {
		return err
	}

	_,
		err = textLineSpecPlainTextAtom{}.ptr().
		testValidityOfTextLineSpecPlainText(
			plainTextLine,
			ePrefix.XCtx(
				"plainTextLine"))

	return err
}

// String - Returns the formatted text generated by this Text Line
// Specification for output display and printing.
//
// If an error occurs, the error message will be included in the
// returned string.
//
// This method is identical in function to
// TextLineSpecPlainText.GetFormattedText()
//
// This method fulfills requirements of the ITextLineSpecification
// interface.
//
func (plainTextLine TextLineSpecPlainText) String() string {

	if plainTextLine.lock == nil {
		plainTextLine.lock = new(sync.Mutex)
	}

	plainTextLine.lock.Lock()

	defer plainTextLine.lock.Unlock()

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TextLineSpecPlainText.GetFormattedText()",
		"")

	formattedTxtStr,
		err := textLineSpecPlainTextNanobot{}.ptr().
		getFormattedText(
			&plainTextLine,
			ePrefix.XCtx("plainTextLine"))

	if err != nil {
		formattedTxtStr = fmt.Sprintf("%v\n",
			err.Error())
	}

	return formattedTxtStr
}

// TextTypeName - returns a string specifying the type
// of Text Line Specification.
//
// This method fulfills requirements of ITextSpecification
// interface.
//
func (plainTextLine *TextLineSpecPlainText) TextTypeName() string {

	if plainTextLine.lock == nil {
		plainTextLine.lock = new(sync.Mutex)
	}

	plainTextLine.lock.Lock()

	defer plainTextLine.lock.Unlock()

	return "TextLineSpecPlainText"
}

// TextLineSpecName - returns a string specifying the name
// of this Text Line Specification.
//
// This method fulfills requirements of ITextLineSpecification
// interface.
//
func (plainTextLine *TextLineSpecPlainText) TextLineSpecName() string {

	if plainTextLine.lock == nil {
		plainTextLine.lock = new(sync.Mutex)
	}

	plainTextLine.lock.Lock()

	defer plainTextLine.lock.Unlock()

	return "TextLineSpecPlainText"
}
