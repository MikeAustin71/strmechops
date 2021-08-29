package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// TextFieldSpecSpacer - The Text Field Specification for one or
// more white space characters.
//
// Text Field Specifications are designed to be configured within a
// line of text. Those lines of text can then be formatted for text
// displays, file output or printing. The type
// TextLineSpecStandardLine can be used to compose a line of text
// consisting of multiple Text Field Specifications like
// TextFieldSpecLabel. Text Field Specifications are therefore
// used as the components or building blocks for single lines of
// text.
//
// Member Variables
//
// ----------------------------------------------------------------
//
//  fieldLen                   int
//     - An integer value greater than zero and less than 1,000,001
//       which is used to specify the number of white space
//       characters in the 'spacer' text field.
//
//       Examples:
//          fieldLen = 1 produces text field " "
//          fieldLen = 2 produces text field "  "
//          fieldLen = 5 produces text field "     "
//
type TextFieldSpecSpacer struct {
	fieldLen int
	lock     *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming instance of
// TextFieldSpecSpacer ('incomingTxtFieldSpacer') to the data fields
// of the current TextFieldSpecSpacer instance ('txtFieldSpacer').
//
// IMPORTANT
// All the data fields in current TextFieldSpecSpacer instance
// ('txtFieldSpacer') will be overwritten and modified.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingTxtFieldSpacer     *TextFieldSpecSpacer
//     - A pointer to an instance of TextFieldSpecSpacer. This
//       method will NOT change the values of internal member
//       variables contained in this instance.
//
//       All data values in this TextFieldSpecSpacer instance
//       will be copied to current TextFieldSpecSpacer
//       instance ('txtFieldSpacer').
//
//       If parameter 'incomingTxtFieldSpacer' is determined to be
//       invalid, an error will be returned.
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
//     - If this method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (txtFieldSpacer *TextFieldSpecSpacer) CopyIn(
	incomingTxtFieldSpacer *TextFieldSpecSpacer,
	errorPrefix interface{}) error {

	if txtFieldSpacer.lock == nil {
		txtFieldSpacer.lock = new(sync.Mutex)
	}

	txtFieldSpacer.lock.Lock()

	defer txtFieldSpacer.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldSpecSpacer.CopyIn()",
		"")

	if err != nil {
		return err
	}

	err = textFieldSpecSpacerNanobot{}.ptr().
		copyIn(
			txtFieldSpacer,
			incomingTxtFieldSpacer,
			ePrefix)

	return err
}

// CopyOut - Returns a deep copy of the current TextFieldSpecSpacer
// instance.
//
// If the current TextFieldSpecSpacer instance is invalid, an error
// will be returned.
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
//  TextFieldSpecSpacer
//     - If this method completes successfully and no errors are
//       encountered, this parameter will return a deep copy of the
//       current TextFieldSpecSpacer instance.
//
//
//  error
//     - If this method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (txtFieldSpacer *TextFieldSpecSpacer) CopyOut(
	errorPrefix interface{}) (
	TextFieldSpecSpacer,
	error) {

	if txtFieldSpacer.lock == nil {
		txtFieldSpacer.lock = new(sync.Mutex)
	}

	txtFieldSpacer.lock.Lock()

	defer txtFieldSpacer.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldSpecSpacer.CopyOut()",
		"")

	if err != nil {
		return TextFieldSpecSpacer{}, err
	}

	return textFieldSpecSpacerNanobot{}.ptr().
		copyOut(
			txtFieldSpacer,
			ePrefix)
}

// NewPtrSpacer - Creates and returns a pointer to a new instance
// of TextFieldSpecSpacer.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  fieldLen                   int
//     - An integer value which specifies the number of white space
//       characters to be included in the spacer text field.
//
//       Examples:
//          fieldLen = 1 produces text field " "
//          fieldLen = 2 produces text field "  "
//          fieldLen = 5 produces text field "     "
//
//       If 'fieldLen' is less than one (+1), an error will be
//       returned.
//
//       If 'fieldLen' is greater than one-million (+1,000,000), an
//       error will be returned.
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
//  *TextFieldSpecSpacer
//     - This method will return a pointer to a new instance of
//       TextFieldSpecSpacer constructed from information provided
//       by the input parameter, 'fieldLen'.
//
//
//  error
//     - If this method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (txtFieldSpacer TextFieldSpecSpacer) NewPtrSpacer(
	fieldLen int,
	errorPrefix interface{}) (*TextFieldSpecSpacer, error) {

	if txtFieldSpacer.lock == nil {
		txtFieldSpacer.lock = new(sync.Mutex)
	}

	txtFieldSpacer.lock.Lock()

	defer txtFieldSpacer.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newTextSpacer := TextFieldSpecSpacer{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldSpecLabel.NewPtrSpacer()",
		"")

	if err != nil {
		return &newTextSpacer, err
	}

	err = textFieldSpecSpacerNanobot{}.ptr().
		setTextFieldSpacer(
			&newTextSpacer,
			fieldLen,
			ePrefix.XCtx("fieldLen invalid"))

	if err != nil {

		return &newTextSpacer, err
	}

	return &newTextSpacer, err
}

// NewSpacer - Creates and returns a new, concrete instance of
// TextFieldSpecSpacer.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  fieldLen                   int
//     - An integer value which specifies the number of white space
//       characters to be included in the spacer text field.
//
//       Examples:
//          fieldLen = 1 produces text field " "
//          fieldLen = 2 produces text field "  "
//          fieldLen = 5 produces text field "     "
//
//       If 'fieldLen' is less than one (+1), an error will be
//       returned.
//
//       If 'fieldLen' is greater than one-million (+1,000,000), an
//       error will be returned.
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
//  TextFieldSpecSpacer
//     - This method will return a new, concrete instance of
//       TextFieldSpecSpacer constructed from information provided
//       by the input parameter, 'fieldLen'.
//
//
//  error
//     - If this method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (txtFieldSpacer TextFieldSpecSpacer) NewSpacer(
	fieldLen int,
	errorPrefix interface{}) (
	TextFieldSpecSpacer, error) {

	if txtFieldSpacer.lock == nil {
		txtFieldSpacer.lock = new(sync.Mutex)
	}

	txtFieldSpacer.lock.Lock()

	defer txtFieldSpacer.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newTextSpacer := TextFieldSpecSpacer{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldSpecLabel.NewSpacer()",
		"")

	if err != nil {
		return newTextSpacer, err
	}

	err = textFieldSpecSpacerNanobot{}.ptr().
		setTextFieldSpacer(
			&newTextSpacer,
			fieldLen,
			ePrefix.XCtx("fieldLen invalid"))

	return newTextSpacer, err
}

// SetFieldLen - Sets the field length for the current instance of
// TextFieldSpecSpacer.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  fieldLen                   int
//     - An integer value which specifies the number of white space
//       characters to be included in the spacer text field.
//
//       Examples:
//          fieldLen = 1 produces text field " "
//          fieldLen = 2 produces text field "  "
//          fieldLen = 5 produces text field "     "
//
//       If 'fieldLen' is less than one (+1), an error will be
//       returned.
//
//       If 'fieldLen' is greater than one-million (+1,000,000), an
//       error will be returned.
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
//     - If this method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (txtFieldSpacer *TextFieldSpecSpacer) SetFieldLen(
	fieldLen int,
	errorPrefix interface{}) (
	err error) {

	if txtFieldSpacer.lock == nil {
		txtFieldSpacer.lock = new(sync.Mutex)
	}

	txtFieldSpacer.lock.Lock()

	defer txtFieldSpacer.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldSpecLabel.SetFieldLen()",
		"")

	if err != nil {
		return err
	}

	err = textFieldSpecSpacerNanobot{}.ptr().
		setTextFieldSpacer(
			txtFieldSpacer,
			fieldLen,
			ePrefix)

	return err
}

// TextFieldName - returns a string specifying the name of the Text
// Field specification.
//
// This method fulfills the requirements of the
// ITextFieldSpecification interface.
//
func (txtFieldSpacer *TextFieldSpecSpacer) TextFieldName() string {

	if txtFieldSpacer.lock == nil {
		txtFieldSpacer.lock = new(sync.Mutex)
	}

	txtFieldSpacer.lock.Lock()

	defer txtFieldSpacer.lock.Unlock()

	return "TextFieldSpecSpacer"
}

// TextTypeName - returns a string specifying the type of Text
// Field specification.
//
// This method fulfills the requirements of the ITextSpecification
// interface.
//
func (txtFieldSpacer *TextFieldSpecSpacer) TextTypeName() string {

	if txtFieldSpacer.lock == nil {
		txtFieldSpacer.lock = new(sync.Mutex)
	}

	txtFieldSpacer.lock.Lock()

	defer txtFieldSpacer.lock.Unlock()

	return "TextFieldSpecSpacer"
}
