package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"strings"
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
// TextFieldSpecSpacer. Text Field Specifications are therefore
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
	fieldLen       int
	textLineReader *strings.Reader
	lock           *sync.Mutex
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

// CopyOutITextField - Returns a deep copy of the current
// TextFieldSpecSpacer instance cast as an ITextFieldSpecification
// object.
//
// If the current TextFieldSpecSpacer instance is invalid, an error
// is returned.
//
// This method is required by the ITextFieldSpecification
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
//  ITextFieldSpecification
//     - If this method completes successfully and no errors are
//       encountered, this parameter will return a deep copy of
//       the current TextFieldSpecSpacer instance cast as an
//       ITextFieldSpecification object.
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
func (txtFieldSpacer *TextFieldSpecSpacer) CopyOutITextField(
	errorPrefix interface{}) (
	ITextFieldSpecification,
	error) {

	if txtFieldSpacer.lock == nil {
		txtFieldSpacer.lock = new(sync.Mutex)
	}

	txtFieldSpacer.lock.Lock()

	defer txtFieldSpacer.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	var iTxtFieldSpec ITextFieldSpecification

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldSpecSpacer."+
			"CopyOutITextField()",
		"")

	if err != nil {
		return iTxtFieldSpec, err
	}

	var newTxtFieldSpacer TextFieldSpecSpacer

	newTxtFieldSpacer,
		err = textFieldSpecSpacerNanobot{}.ptr().
		copyOut(
			txtFieldSpacer,
			ePrefix)

	if err != nil {
		return iTxtFieldSpec, err
	}

	iTxtFieldSpec = ITextFieldSpecification(&newTxtFieldSpacer)

	return iTxtFieldSpec, nil
}

// CopyOutPtr - Returns a pointer to a deep copy of the current
// TextFieldSpecSpacer instance.
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
//  *TextFieldSpecSpacer
//     - If this method completes successfully and no errors are
//       encountered, this parameter will return a pointer to a
//       deep copy of the current TextFieldSpecSpacer instance.
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
func (txtFieldSpacer *TextFieldSpecSpacer) CopyOutPtr(
	errorPrefix interface{}) (
	*TextFieldSpecSpacer,
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
		"TextFieldSpecSpacer.CopyOutPtr()",
		"")

	if err != nil {
		return &TextFieldSpecSpacer{}, err
	}

	var newTxtFieldSpacer TextFieldSpecSpacer

	newTxtFieldSpacer,
		err = textFieldSpecSpacerNanobot{}.ptr().
		copyOut(
			txtFieldSpacer,
			ePrefix)

	return &newTxtFieldSpacer, err
}

// Empty - Resets all internal member variables for the current
// instance of TextFieldSpecSpacer to their initial or zero states.
//
func (txtFieldSpacer *TextFieldSpecSpacer) Empty() {

	if txtFieldSpacer.lock == nil {
		txtFieldSpacer.lock = new(sync.Mutex)
	}

	txtFieldSpacer.lock.Lock()

	textFieldSpecSpacerNanobot{}.ptr().empty(
		txtFieldSpacer)

	txtFieldSpacer.lock.Unlock()

	txtFieldSpacer.lock = nil

	return
}

// Equal - Receives a pointer to another instance of
// TextFieldSpecSpacer and proceeds to compare the member variables
// to those of the current TextFieldSpecSpacer instance in order to
// determine if they are equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables are equal in all respects,
// this flag is set to 'true'. Otherwise, this method returns
// 'false'.
//
func (txtFieldSpacer *TextFieldSpecSpacer) Equal(
	incomingTxtFieldSpacer *TextFieldSpecSpacer) bool {

	if txtFieldSpacer.lock == nil {
		txtFieldSpacer.lock = new(sync.Mutex)
	}

	txtFieldSpacer.lock.Lock()

	defer txtFieldSpacer.lock.Unlock()

	return textFieldSpecSpacerNanobot{}.ptr().
		equal(
			txtFieldSpacer,
			incomingTxtFieldSpacer)

}

// EqualITextField - Receives an object implementing the
// ITextFieldSpecification interface and proceeds to compare
// the member variables to those of the current TextFieldSpecSpacer
// instance in order to determine if they are equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables from both instances are equal
// in all respects, this flag is set to 'true'. Otherwise, this method returns
// 'false'.
//
func (txtFieldSpacer *TextFieldSpecSpacer) EqualITextField(
	iTextField ITextFieldSpecification) bool {

	if txtFieldSpacer.lock == nil {
		txtFieldSpacer.lock = new(sync.Mutex)
	}

	txtFieldSpacer.lock.Lock()

	defer txtFieldSpacer.lock.Unlock()

	if iTextField == nil {
		return false
	}

	txtSpacer, ok := iTextField.(*TextFieldSpecSpacer)

	if !ok {
		return false
	}

	return textFieldSpecSpacerNanobot{}.ptr().
		equal(
			txtFieldSpacer,
			txtSpacer)
}

// GetFieldLength - Returns an integer value specifying the length
// of the white space text characters produced by this the current
// instance of TextFieldSpecSpacer.
//
// TextFieldSpecSpacer is a Text Field Specification which produces
// a text string equal to one or more white space characters. The
// number of white space characters in the string is determined by
// the field length parameter. (See examples below)
//
//       Examples:
//          fieldLen = 1 produces text field " "
//          fieldLen = 2 produces text field "  "
//          fieldLen = 5 produces text field "     "
//
func (txtFieldSpacer *TextFieldSpecSpacer) GetFieldLength() int {

	if txtFieldSpacer.lock == nil {
		txtFieldSpacer.lock = new(sync.Mutex)
	}

	txtFieldSpacer.lock.Lock()

	defer txtFieldSpacer.lock.Unlock()

	return txtFieldSpacer.fieldLen
}

// GetFormattedStrLength - Returns the string length of the
// formatted text generated by the current instance of
// TextFieldSpecSpacer. Effectively, this is the length of the
// strings returned by methods:
//   TextFieldSpecSpacer.GetFormattedText()
//   TextFieldSpecSpacer.String()
//
// If an error is encountered, this method returns a value of minus
// one (-1).
//
func (txtFieldSpacer *TextFieldSpecSpacer) GetFormattedStrLength() int {

	if txtFieldSpacer.lock == nil {
		txtFieldSpacer.lock = new(sync.Mutex)
	}

	txtFieldSpacer.lock.Lock()

	defer txtFieldSpacer.lock.Unlock()

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TextFieldSpecSpacer.GetFormattedStrLength()",
		"")

	formattedTextStr,
		err := textFieldSpecSpacerNanobot{}.ptr().
		getFormattedText(
			txtFieldSpacer,
			ePrefix.XCpy(
				"txtFieldSpacer"))

	if err != nil {
		return -1
	}

	return len(formattedTextStr)

}

// GetFormattedText - Returns the formatted text generated by the
// current instance of TextFieldSpecSpacer.
//
// TextFieldSpecSpacer is a Text Field Specification which produces
// a text string equal to one or more white space characters. The
// number of white space characters in the string is determined by
// the field length parameter. (See examples below)
//
// This method is identical in function to
// TextFieldSpecSpacer.String()
//
// This method fulfills the requirements of the
// ITextFieldSpecification interface.
//
// Methods which return formatted text are listed as follows:
//  TextFieldSpecSpacer.String()
//  TextFieldSpecSpacer.GetFormattedText()
//  TextFieldSpecSpacer.TextBuilder()
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
//  string
//     - This method will return a string of white space characters
//       generated by the current instance of TextFieldSpecSpacer.
//       The number of white space characters in this returned
//       text string is controlled by the field length parameter.
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
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//  fieldLen = 1 produces formatted text string " "
//  fieldLen = 2 produces formatted text string "  "
//  fieldLen = 5 produces formatted text string "     "
//
func (txtFieldSpacer *TextFieldSpecSpacer) GetFormattedText(
	errorPrefix interface{}) (
	string,
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
		"TextFieldSpecSpacer.GetFormattedText()",
		"")

	if err != nil {
		return "", err
	}

	return textFieldSpecSpacerNanobot{}.ptr().
		getFormattedText(
			txtFieldSpacer,
			ePrefix.XCpy(
				"txtFieldSpacer"))
}

// IsValidInstance - Performs a diagnostic review of the data
// values encapsulated in the current TextFieldSpecSpacer instance
// to determine if they are valid.
//
// If all data element evaluate as valid, this method returns
// 'true'. If any data element is invalid, this method returns
// 'false'.
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
//  isValid             bool
//     - If all data elements encapsulated by the current instance
//       of TextFieldSpecSpacer are valid, this returned boolean
//       value is set to 'true'. If any data values are invalid,
//       this return parameter is set to 'false'.
//
func (txtFieldSpacer *TextFieldSpecSpacer) IsValidInstance() (
	isValid bool) {

	if txtFieldSpacer.lock == nil {
		txtFieldSpacer.lock = new(sync.Mutex)
	}

	txtFieldSpacer.lock.Lock()

	defer txtFieldSpacer.lock.Unlock()

	isValid,
		_ = textFieldSpecSpacerElectron{}.ptr().
		isFieldLenValidError(
			txtFieldSpacer.fieldLen,
			nil)

	return isValid
}

// IsValidInstanceError - Performs a diagnostic review of the data
// values encapsulated in the current TextFieldSpecSpacer instance
// to determine if they are valid.
//
// If any data element evaluates as invalid, this method will
// return an error.
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
//       the current instance of TextFieldSpecSpacer are found to be
//       invalid, this method will return an error.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' (error prefix) will be inserted or
//       prefixed at the beginning of the error message.
//
func (txtFieldSpacer *TextFieldSpecSpacer) IsValidInstanceError(
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
		"TextFieldSpecSpacer.IsValidInstanceError()",
		"")

	if err != nil {
		return err
	}

	_,
		err = textFieldSpecSpacerElectron{}.ptr().
		isFieldLenValidError(
			txtFieldSpacer.fieldLen,
			ePrefix.XCpy(
				"txtFieldSpacer.fieldLen invalid!"))

	return err
}

// New - Returns a new concrete instance of TextFieldSpecSpacer.
// This returned instance is empty and unpopulated. All the member
// variables contained in this new instance are set to their
// uninitialized or zero values.
//
// Be advised that setting member variables to their zero values
// means that the returned TextFieldSpecSpacer instance is
// invalid. Therefore, in order to use this TextFieldSpecSpacer
// instance, users must later call the setter methods on this type
// in order to configure valid and meaningful member variable data
// values.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  --- NONE ---
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  TextFieldSpecSpacer
//     - This parameter returns a new and empty concrete instance
//       of TextFieldSpecSpacer. Member variable data values are
//       set to their initial or zero values.
//
func (txtFieldSpacer TextFieldSpecSpacer) New() TextFieldSpecSpacer {

	if txtFieldSpacer.lock == nil {
		txtFieldSpacer.lock = new(sync.Mutex)
	}

	txtFieldSpacer.lock.Lock()

	defer txtFieldSpacer.lock.Unlock()

	newTxtFieldSpacer := TextFieldSpecSpacer{}

	newTxtFieldSpacer.textLineReader = nil

	return newTxtFieldSpacer
}

// NewPtr - Returns a pointer to a new unpopulated instance of
// TextFieldSpecSpacer. All the member variables contained in
// this new instance are set to their uninitialized or zero values.
//
// Be advised that setting member variables to their zero values
// means that the returned TextFieldSpecSpacer instance is invalid.
// Therefore, in order to use this TextFieldSpecSpacer instance,
// users must later call the setter methods on this type in order
// to configure valid and meaningful member variable data values.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  --- NONE ---
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  *TextFieldSpecSpacer
//     - This parameter returns a pointer to a new, empty instance
//       of TextFieldSpecSpacer. Member variable data values are
//       set to their initial or zero values. The returned Text
//       Field Spacer Specification is therefore invalid.
//
func (txtFieldSpacer TextFieldSpecSpacer) NewPtr() *TextFieldSpecSpacer {

	if txtFieldSpacer.lock == nil {
		txtFieldSpacer.lock = new(sync.Mutex)
	}

	txtFieldSpacer.lock.Lock()

	defer txtFieldSpacer.lock.Unlock()

	newTxtFieldSpacer := TextFieldSpecSpacer{}

	newTxtFieldSpacer.textLineReader = nil

	return &newTxtFieldSpacer
}

// NewPtrSpacer - Creates and returns a pointer to a new instance
// of TextFieldSpecSpacer.
//
// TextFieldSpecSpacer is a Text Field Specification which produces
// a text string equal to one or more white space characters. The
// number of white space characters in the string is determined by
// the field length parameter. (See examples below)
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
		"TextFieldSpecSpacer.NewPtrSpacer()",
		"")

	if err != nil {
		return &newTextSpacer, err
	}

	err = textFieldSpecSpacerNanobot{}.ptr().
		setTextFieldSpacer(
			&newTextSpacer,
			fieldLen,
			ePrefix.XCpy("->newTextSpacer"))

	if err != nil {

		return &newTextSpacer, err
	}

	return &newTextSpacer, err
}

// NewSpacer - Creates and returns a new, concrete instance of
// TextFieldSpecSpacer.
//
// TextFieldSpecSpacer is a Text Field Specification which produces
// a text string equal to one or more white space characters. The
// number of white space characters in the string is determined by
// the field length parameter. (See examples below)
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
		"TextFieldSpecSpacer.NewSpacer()",
		"")

	if err != nil {
		return newTextSpacer, err
	}

	err = textFieldSpecSpacerNanobot{}.ptr().
		setTextFieldSpacer(
			&newTextSpacer,
			fieldLen,
			ePrefix.XCpy("->newTextSpacer"))

	return newTextSpacer, err
}

// Read - Implements the io.Reader interface for type
// TextFieldSpecSpacer.
//
// The formatted text string generated by the current instance of
// TextFieldSpecSpacer will be written to the byte buffer 'p'. If
// the length of 'p' is less than the length of the formatted text
// string, multiple calls to this method will write the remaining
// unread characters to the byte buffer 'p'.
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
// This method fulfills the requirements of the
// ITextFieldSpecification interface.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  p                          []byte
//     - The byte buffer into which the formatted text string
//       generated by the current txtFillerField instance will be
//       written.
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
//    err = txtFieldSpacer01.Read(p)
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
//      "Error Returned From txtFieldSpacer01.Read(p)\n"+
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
//    err = txtFieldSpacer01.Read(p)
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
//      "Error Returned From txtFieldSpacer01.Read(p)\n"+
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
func (txtFieldSpacer *TextFieldSpecSpacer) Read(
	p []byte) (
	n int,
	err error) {

	if txtFieldSpacer.lock == nil {
		txtFieldSpacer.lock = new(sync.Mutex)
	}

	txtFieldSpacer.lock.Lock()

	defer txtFieldSpacer.lock.Unlock()

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TextFieldSpecSpacer.Read()",
		"")

	if txtFieldSpacer.textLineReader == nil {

		var formattedText string

		formattedText,
			err = textFieldSpecSpacerNanobot{}.ptr().
			getFormattedText(
				txtFieldSpacer,
				ePrefix.XCpy(
					"txtFieldSpacer"))

		if err != nil {
			return n, err
		}

		txtFieldSpacer.textLineReader =
			strings.NewReader(formattedText)

		if txtFieldSpacer.textLineReader == nil {
			err = fmt.Errorf("%v\n"+
				"Error: strings.NewReader(formattedText)\n"+
				"returned a nil pointer.\n"+
				"txtFillerField.textLineReader == nil\n",
				ePrefix.String())

			return n, err
		}
	}

	n,
		err = textSpecificationAtom{}.ptr().
		readBytes(
			txtFieldSpacer.textLineReader,
			p,
			ePrefix.XCpy(
				"p -> txtFieldSpacer.textLineReader"))

	if err == io.EOF {

		txtFieldSpacer.textLineReader = nil

	}

	return n, err
}

// ReaderInitialize - This method will reset the internal member
// variable 'TextFieldSpecSpacer.textLineReader' to its initial
// zero state of 'nil'. Effectively, this resets the internal
// strings.Reader object for use in future read operations.
//
// This method is rarely used or needed. It provides a means of
// reinitializing the internal strings.Reader object in case an
// error occurs during a read operation initiated by method
// TextFieldSpecSpacer.Read().
//
// Calling this method cleans up the residue from an aborted read
// operation and prepares the strings.Reader object for future read
// operations.
//
// If any errors are returned by method
// TextFieldSpecSpacer.Read() which are NOT equal to io.EOF, call
// this method, TextFieldSpecSpacer.ReaderInitialize(), to reset
// and prepare the internal reader for future read operations.
//
// This method fulfills the requirements of the
// ITextFieldSpecification interface.
//
func (txtFieldSpacer *TextFieldSpecSpacer) ReaderInitialize() {

	if txtFieldSpacer.lock == nil {
		txtFieldSpacer.lock = new(sync.Mutex)
	}

	txtFieldSpacer.lock.Lock()

	defer txtFieldSpacer.lock.Unlock()

	txtFieldSpacer.textLineReader = nil

	return
}

// SetFieldLen - Sets the field length for the current instance of
// TextFieldSpecSpacer.
//
// TextFieldSpecSpacer is a Text Field Specification which produces
// a text string equal to one or more white space characters. The
// number of white space characters in the string is determined by
// the field length parameter. (See examples below)
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
		"TextFieldSpecSpacer.SetFieldLen()",
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

// String - Returns the formatted text generated by the
// current instance of TextFieldSpecSpacer.
//
// TextFieldSpecSpacer is a Text Field Specification which produces
// a text string equal to one or more white space characters. The
// number of white space characters in the string is determined by
// the field length parameter. (See examples below)
//
// This method is identical in function to
// TextFieldSpecSpacer.GetFormattedText()
//
// This method fulfills the requirements of the
// ITextFieldSpecification interface.
//
// This method also fulfills the requirements of the 'Stringer'
// interface defined in the Golang package 'fmt'. Reference:
//   https://pkg.go.dev/fmt#Stringer
//
// Methods which return formatted text are listed as follows:
//  TextFieldSpecSpacer.String()
//  TextFieldSpecSpacer.GetFormattedText()
//  TextFieldSpecSpacer.TextBuilder()
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//  fieldLen = 1 produces text field " "
//  fieldLen = 2 produces text field "  "
//  fieldLen = 5 produces text field "     "
//
//
// ----------------------------------------------------------------
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
//  string
//     - This method will return a string of white space characters
//       generated by the current instance of TextFieldSpecSpacer.
//       The number of white space characters in this returned
//       text string is controlled by the field length parameter.
//
func (txtFieldSpacer *TextFieldSpecSpacer) String() string {

	if txtFieldSpacer.lock == nil {
		txtFieldSpacer.lock = new(sync.Mutex)
	}

	txtFieldSpacer.lock.Lock()

	defer txtFieldSpacer.lock.Unlock()

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TextFieldSpecSpacer.String()",
		"")

	formattedText,
		err := textFieldSpecSpacerNanobot{}.ptr().
		getFormattedText(
			txtFieldSpacer,
			&ePrefix)

	if err != nil {
		formattedText = fmt.Sprintf(
			"%v", err.Error())
	}

	return formattedText
}

// TextBuilder - Configures the line of text produced by this
// instance of TextFieldSpecSpacer, and writes it to an instance
// of strings.Builder.
//
// This method fulfills the requirements of the
// ITextFieldSpecification interface.
//
// Methods which return formatted text are listed as follows:
//  TextFieldSpecSpacer.String()
//  TextFieldSpecSpacer.GetFormattedText()
//  TextFieldSpecSpacer.TextBuilder()
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  sBuilder                   *strings.Builder
//    - A pointer to an instance of strings.Builder. The line of
//      text produced by the current instance of
//      TextFieldSpecSpacer and writes that text to 'sBuilder'.
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
func (txtFieldSpacer *TextFieldSpecSpacer) TextBuilder(
	sBuilder *strings.Builder,
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
		"TextFieldSpecSpacer.TextBuilder()",
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
		err = textFieldSpecSpacerNanobot{}.ptr().
		getFormattedText(
			txtFieldSpacer,
			ePrefix.XCpy(
				"txtFieldSpacer"))

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
