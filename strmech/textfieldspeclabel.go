package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"strings"
	"sync"
)

// TextFieldSpecLabel - This Text Field Specification is used to
// generate a text label string. The text label is positioned
// inside a text field with a given field length. Text
// Justification within this text field is controlled by the
// Text Justification specification value which may be set to
// 'Left', 'Right' or 'Center'. A text label contains a string of
// text characters.
//
// Text Label Examples:
//   'Hello World  ' - Left Justified, Field Length= 13
//   '  Hello World' - Right Justified, Field Length= 13
//   ' Hello World ' - Centered, Field Length= 13
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
//  textLabel                  []rune
//     - An array of runes or text characters which is used to
//       generate string content for display as a text label.
//
//
//  fieldLen                   int
//     - The length of the text field in which the 'textLabel'
//       characters will be displayed. If 'fieldLen' is less than the length
//       of the 'textLabel' array, 'fieldLen' will be automatically
//
//
//  textJustification          TextJustify
//     - An enumeration which specifies the justification of the
//       'textLabelChars' within the field specified by 'fieldLen'.
//
//       Text justification can only be evaluated in the context of
//       a text label, field length and a 'textJustification'
//       object of type TextJustify. This is because text labels
//       with a field length equal to or less than the length of
//       the text label never use text justification. In these
//       cases, text justification is completely ignored.
//
//       If the field length is greater than the length of the text
//       label, text justification must be equal to one of these
//       three valid values:
//           TextJustify(0).Left()
//           TextJustify(0).Right()
//           TextJustify(0).Center()
//
//       You can also use the abbreviated text justification
//       enumeration syntax as follows:
//
//           TxtJustify.Left()
//           TxtJustify.Right()
//           TxtJustify.Center()
//
type TextFieldSpecLabel struct {
	textLabel []rune // The text content of the label.
	fieldLen  int    // The length of the text field in which
	//               //  the text label will be positioned.
	textJustification TextJustify // The specification which controls
	//                            //  how the text label will be positioned
	//                            //  within the text field: 'Left', 'Right'
	//                            //  or 'Center'.
	textLineReader *strings.Reader
	lock           *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming instance of
// TextFieldSpecLabel ('incomingTxtFieldLabel') to the data fields
// of the current TextFieldSpecLabel instance ('txtFieldLabel').
//
// IMPORTANT
// All the data fields in current TextFieldSpecLabel instance
// ('txtFieldLabel') will be modified and overwritten.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingTxtFieldLabel     *TextFieldSpecLabel
//     - A pointer to an instance of TextFieldSpecLabel. This
//       method will NOT change the values of internal member
//       variables contained in this instance.
//
//       All data values in this TextFieldSpecLabel instance
//       will be copied to current TextFieldSpecLabel
//       instance ('txtFieldLabel').
//
//       If parameter 'incomingTxtFieldLabel' is determined to be
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
func (txtFieldLabel *TextFieldSpecLabel) CopyIn(
	incomingTxtFieldLabel *TextFieldSpecLabel,
	errorPrefix interface{}) error {

	if txtFieldLabel.lock == nil {
		txtFieldLabel.lock = new(sync.Mutex)
	}

	txtFieldLabel.lock.Lock()

	defer txtFieldLabel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldSpecLabel.CopyIn()",
		"")

	if err != nil {
		return err
	}

	return textFieldSpecLabelMolecule{}.ptr().
		copyIn(
			txtFieldLabel,
			incomingTxtFieldLabel,
			ePrefix)
}

// CopyOut - Returns a deep copy of the current TextFieldSpecLabel
// instance.
//
// If the current TextFieldSpecLabel instance is invalid, an error
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
//  TextFieldSpecLabel
//     - If this method completes successfully and no errors are
//       encountered, this parameter will return a deep copy of the
//       current TextFieldSpecLabel instance.
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
func (txtFieldLabel *TextFieldSpecLabel) CopyOut(
	errorPrefix interface{}) (
	TextFieldSpecLabel,
	error) {

	if txtFieldLabel.lock == nil {
		txtFieldLabel.lock = new(sync.Mutex)
	}

	txtFieldLabel.lock.Lock()

	defer txtFieldLabel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldSpecLabel.CopyOut()",
		"")

	if err != nil {
		return TextFieldSpecLabel{}, err
	}

	return textFieldSpecLabelMolecule{}.ptr().
		copyOut(txtFieldLabel,
			ePrefix)
}

// CopyOutITextField - Returns a deep copy of the current
// TextFieldSpecLabel instance cast as an ITextFieldSpecification
// object.
//
// If the current TextFieldSpecLabel instance is invalid, an error
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
//       the current TextFieldSpecLabel instance cast as an
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
func (txtFieldLabel *TextFieldSpecLabel) CopyOutITextField(
	errorPrefix interface{}) (
	ITextFieldSpecification,
	error) {

	if txtFieldLabel.lock == nil {
		txtFieldLabel.lock = new(sync.Mutex)
	}

	txtFieldLabel.lock.Lock()

	defer txtFieldLabel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	var iTxtFieldSpec ITextFieldSpecification

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldSpecLabel.CopyOutITextField()",
		"")

	if err != nil {
		return iTxtFieldSpec, err
	}

	var newTxtFieldLabel TextFieldSpecLabel

	newTxtFieldLabel,
		err = textFieldSpecLabelMolecule{}.ptr().
		copyOut(txtFieldLabel,
			ePrefix)

	if err != nil {
		return iTxtFieldSpec, err
	}

	iTxtFieldSpec = ITextFieldSpecification(&newTxtFieldLabel)

	return iTxtFieldSpec, nil
}

// CopyOutPtr - Returns a pointer to a deep copy of the current
// TextFieldSpecLabel instance.
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
//  *TextFieldSpecLabel
//     - If this method completes successfully and no errors are
//       encountered, this parameter will return a pointer to a
//       deep copy of the current TextFieldSpecLabel instance.
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
func (txtFieldLabel *TextFieldSpecLabel) CopyOutPtr(
	errorPrefix interface{}) (
	*TextFieldSpecLabel,
	error) {

	if txtFieldLabel.lock == nil {
		txtFieldLabel.lock = new(sync.Mutex)
	}

	txtFieldLabel.lock.Lock()

	defer txtFieldLabel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldSpecLabel.CopyOutPtr()",
		"")

	if err != nil {
		return &TextFieldSpecLabel{}, err
	}

	var newTxtFieldLabel TextFieldSpecLabel

	newTxtFieldLabel,
		err = textFieldSpecLabelMolecule{}.ptr().
		copyOut(txtFieldLabel,
			ePrefix)

	return &newTxtFieldLabel, err
}

// Empty - Resets all internal member variables for the current
// instance of TextFieldSpecLabel to their initial or zero states.
//
func (txtFieldLabel *TextFieldSpecLabel) Empty() {

	if txtFieldLabel.lock == nil {
		txtFieldLabel.lock = new(sync.Mutex)
	}

	txtFieldLabel.lock.Lock()

	textFieldSpecLabelMolecule{}.ptr().
		empty(txtFieldLabel)

	txtFieldLabel.lock.Unlock()

	txtFieldLabel.lock = nil

}

// Equal - Receives a pointer to another instance of
// TextFieldSpecLabel and proceeds to compare the member variables
// to those of the current TextFieldSpecLabel instance in order to
// determine if they are equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables are equal in all respects,
// this flag is set to 'true'. Otherwise, this method returns
// 'false'.
//
func (txtFieldLabel *TextFieldSpecLabel) Equal(
	incomingTextFieldLabel *TextFieldSpecLabel) bool {

	if txtFieldLabel.lock == nil {
		txtFieldLabel.lock = new(sync.Mutex)
	}

	txtFieldLabel.lock.Lock()

	defer txtFieldLabel.lock.Unlock()

	return textFieldSpecLabelMolecule{}.ptr().
		equal(
			txtFieldLabel,
			incomingTextFieldLabel)
}

// EqualITextField - Receives an object implementing the
// ITextFieldSpecification interface and proceeds to compare
// the member variables to those of the current TextFieldSpecLabel
// instance in order to determine if they are equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables from both instances are equal
// in all respects, this flag is set to 'true'. Otherwise, this
// method returns 'false'.
//
func (txtFieldLabel *TextFieldSpecLabel) EqualITextField(
	iTextField ITextFieldSpecification) bool {

	if txtFieldLabel.lock == nil {
		txtFieldLabel.lock = new(sync.Mutex)
	}

	txtFieldLabel.lock.Lock()

	defer txtFieldLabel.lock.Unlock()

	if iTextField == nil {
		return false
	}

	txtLabel, ok := iTextField.(*TextFieldSpecLabel)

	if !ok {
		return false
	}

	return textFieldSpecLabelMolecule{}.ptr().
		equal(
			txtFieldLabel,
			txtLabel)
}

// GetFieldLength - Returns the length of the text field in which
// the text label string will be positioned.
//
// ------------------------------------------------------------------------
//
// Background
//
// Field Length specifies the length of the text field in which the
// formatted text label will be displayed.
//
// If Field Length is less than the length of the formatted text
// label string, it will be automatically set equal to the length
// of the formatted text label string.
//
// If Field Length is greater than the length of the length of the
// formatted text label string, the text label text will be
// positioned within a text field with a length equal to Field
// Length. In this case, the position of the text label string
// within the text field will be controlled by the text
// justification specification.
//
func (txtFieldLabel *TextFieldSpecLabel) GetFieldLength() int {

	if txtFieldLabel.lock == nil {
		txtFieldLabel.lock = new(sync.Mutex)
	}

	txtFieldLabel.lock.Lock()

	defer txtFieldLabel.lock.Unlock()

	return txtFieldLabel.fieldLen
}

// GetFormattedStrLength - Returns the string length of the
// formatted text generated by the current instance of
// TextFieldSpecDateTime. Effectively, this is the length of the
// strings returned by methods:
//   TextFieldSpecLabel.GetFormattedText()
//   TextFieldSpecLabel.String()
//
// If an error is encountered, this method returns a value of minus
// one (-1).
//
func (txtFieldLabel *TextFieldSpecLabel) GetFormattedStrLength() int {

	if txtFieldLabel.lock == nil {
		txtFieldLabel.lock = new(sync.Mutex)
	}

	txtFieldLabel.lock.Lock()

	defer txtFieldLabel.lock.Unlock()

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TextFieldSpecLabel.GetFormattedStrLength()",
		"")

	formattedTextStr,
		err := textSpecificationMolecule{}.ptr().
		getFormattedText(
			txtFieldLabel.textLabel,
			txtFieldLabel.fieldLen,
			txtFieldLabel.textJustification,
			ePrefix.XCpy(
				"txtFieldLabel"))

	if err != nil {
		return -1
	}

	return len(formattedTextStr)
}

// GetFormattedText - Returns the formatted text generated by the
// current instance of TextFieldSpecLabel.
//
// If the length of the text label string is zero and the field
// length is zero this method returns an empty string.
//
// If the length of the text label string is zero and the field
// length is greater than zero, this method returns a string with
// a length equal to field length and content equal to white space
// (the space character " " x field length).
//
// This method is similar to method TextFieldSpecLabel.String()
// with the exception that this method returns an error.
//
// This method fulfills the requirements of the
// ITextFieldSpecification interface.
//
// Methods which return formatted text are listed as follows:
//  TextFieldSpecLabel.String()
//  TextFieldSpecLabel.GetFormattedText()
//  TextFieldSpecLabel.TextBuilder()
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
//     - The formatted text string generated by the current
//       instance of TextFieldSpecLabel.
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
//  Example 1:
//   textLabel = "Hi There" (Length = 8)
//    fieldLen = 12
//    textJustification = TextJustify(0).Center()
//      result = "  Hi There  "
//
//  Example 2:
//   textLabel = "Hi There" (Length = 8)
//    fieldLen = 12
//    textJustification = TextJustify(0).Left()
//      result = "Hi There    "
//
//  Example 3:
//   textLabel = "Hi There" (Length = 8)
//    fieldLen = 12
//    textJustification = TextJustify(0).Right()
//      result = "    Hi There"
//
//  Example 4:
//   textLabel = "Hi There" (Length = 8)
//    fieldLen = -1
//    textJustification = TextJustify(0).Right()
//      result = "Hi There"
//
func (txtFieldLabel *TextFieldSpecLabel) GetFormattedText(
	errorPrefix interface{}) (
	string,
	error) {

	if txtFieldLabel.lock == nil {
		txtFieldLabel.lock = new(sync.Mutex)
	}

	txtFieldLabel.lock.Lock()

	defer txtFieldLabel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldSpecLabel.GetFormattedText()",
		"")

	if err != nil {
		return "", err
	}

	return textSpecificationMolecule{}.ptr().
		getFormattedText(
			txtFieldLabel.textLabel,
			txtFieldLabel.fieldLen,
			txtFieldLabel.textJustification,
			ePrefix.XCpy(
				"txtFieldLabel"))
}

// GetTextJustification - Returns the value of the text
// justification specification which will be used to position the
// text label string with a text field.
//
// ------------------------------------------------------------------------
//
// Background
//
// The text justification enumeration specification should be set
// to one of three valid values:
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
// Text justification is only applied when the text field length is
// greater than the length of the formatted text label string.
//
func (txtFieldLabel *TextFieldSpecLabel) GetTextJustification() TextJustify {

	if txtFieldLabel.lock == nil {
		txtFieldLabel.lock = new(sync.Mutex)
	}

	txtFieldLabel.lock.Lock()

	defer txtFieldLabel.lock.Unlock()

	return txtFieldLabel.textJustification
}

// GetTextLabel - Returns the un-formatted text label string
// associated with the current instance of TextFieldSpecLabel.
//
// ------------------------------------------------------------------------
//
// Background
//
// The Text Field Specification is used to generate a text label
// string. The text label is positioned inside a text field with a
// given field length. Text Justification within this text field
// is controlled by the Text Justification specification value
// which may be set to 'Left', 'Right' or 'Center'. A text label
// contains a string of text characters.
//
// Text Label Examples:
//   'Hello World  ' - Left Justified, Field Length= 13
//   '  Hello World' - Right Justified, Field Length= 13
//   ' Hello World ' - Centered, Field Length= 13
//
func (txtFieldLabel *TextFieldSpecLabel) GetTextLabel() string {

	if txtFieldLabel.lock == nil {
		txtFieldLabel.lock = new(sync.Mutex)
	}

	txtFieldLabel.lock.Lock()

	defer txtFieldLabel.lock.Unlock()

	return string(txtFieldLabel.textLabel)
}

// GetTextLabelRunes - Returns an array of rune characters representing
// the un-formatted text label associated with the current instance of TextFieldSpecLabel.
//
func (txtFieldLabel *TextFieldSpecLabel) GetTextLabelRunes() []rune {

	if txtFieldLabel.lock == nil {
		txtFieldLabel.lock = new(sync.Mutex)
	}

	txtFieldLabel.lock.Lock()

	defer txtFieldLabel.lock.Unlock()

	var newTextLabelRunes []rune

	err :=
		strMechPreon{}.ptr().copyRuneArrays(
			&newTextLabelRunes,
			&txtFieldLabel.textLabel,
			true,
			nil)

	if err != nil {
		return nil
	}

	return newTextLabelRunes
}

// IsValidInstance - Performs a diagnostic review of the data
// values encapsulated in the current TextFieldSpecLabel instance
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
//       of TextFieldSpecLabel are valid, this returned boolean
//       value is set to 'true'. If any data values are invalid,
//       this return parameter is set to 'false'.
//
func (txtFieldLabel *TextFieldSpecLabel) IsValidInstance() (
	isValid bool) {

	if txtFieldLabel.lock == nil {
		txtFieldLabel.lock = new(sync.Mutex)
	}

	txtFieldLabel.lock.Lock()

	defer txtFieldLabel.lock.Unlock()

	isValid,
		_ = textFieldSpecLabelAtom{}.ptr().
		isValidTextFieldLabel(
			txtFieldLabel,
			nil)

	return isValid
}

// IsValidInstanceError - Performs a diagnostic review of the data
// values encapsulated in the current TextFieldSpecLabel instance
// to determine if they are all valid.
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
//       the current instance of TextFieldSpecLabel are found to be
//       invalid, this method will return an error. If the member
//       data variables are determined to be valid, this error
//       return parameter will be set to 'nil'.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' (error prefix) will be inserted or
//       prefixed at the beginning of the error message.
//
func (txtFieldLabel *TextFieldSpecLabel) IsValidInstanceError(
	errorPrefix interface{}) error {

	if txtFieldLabel.lock == nil {
		txtFieldLabel.lock = new(sync.Mutex)
	}

	txtFieldLabel.lock.Lock()

	defer txtFieldLabel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldSpecLabel.IsValidInstanceError()",
		"")

	if err != nil {
		return err
	}

	_,
		err = textFieldSpecLabelAtom{}.ptr().
		isValidTextFieldLabel(
			txtFieldLabel,
			ePrefix)

	return err
}

// New - Returns a new concrete instance of TextFieldSpecLabel.
// This returned instance is empty and unpopulated. All the member
// variables contained in this new instance are set to their
// uninitialized or zero values.
//
// Be advised that setting member variables to their zero values
// means that the returned TextFieldSpecLabel instance is invalid.
// Therefore, in order to use this TextFieldSpecLabel instance,
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
//  TextFieldSpecLabel
//     - This parameter returns a new and empty concrete instance
//       of TextFieldSpecLabel. Member variable data values are
//       set to their initial or zero values.
//
func (txtFieldLabel TextFieldSpecLabel) New() TextFieldSpecLabel {

	if txtFieldLabel.lock == nil {
		txtFieldLabel.lock = new(sync.Mutex)
	}

	txtFieldLabel.lock.Lock()

	defer txtFieldLabel.lock.Unlock()

	newTxtFieldLabel := TextFieldSpecLabel{}

	newTxtFieldLabel.textJustification = TextJustify(0).None()

	newTxtFieldLabel.textLineReader = nil

	newTxtFieldLabel.lock = new(sync.Mutex)

	return newTxtFieldLabel
}

// NewPtr - Returns a pointer to a new unpopulated instance of
// TextFieldSpecLabel. All the member variables contained in
// this new instance are set to their uninitialized or zero values.
//
// Be advised that setting member variables to their zero values
// means that the returned TextFieldSpecLabel instance is invalid.
// Therefore, in order to use this TextFieldSpecLabel instance,
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
//  *TextFieldSpecLabel
//     - This parameter returns a pointer to a new, empty instance
//       of TextFieldSpecLabel. Member variable data values are
//       set to their initial or zero values.
//
func (txtFieldLabel TextFieldSpecLabel) NewPtr() *TextFieldSpecLabel {

	if txtFieldLabel.lock == nil {
		txtFieldLabel.lock = new(sync.Mutex)
	}

	txtFieldLabel.lock.Lock()

	defer txtFieldLabel.lock.Unlock()

	newLabelField := TextFieldSpecLabel{}

	newLabelField.textJustification = TextJustify(0).None()

	newLabelField.textLineReader = nil

	newLabelField.lock = new(sync.Mutex)

	return &newLabelField
}

// NewPtrTextLabel - Creates and returns a pointer to a new, fully
// populated instance of TextFieldSpecLabel. This type encapsulates
// a string which is formatted as a text label.
//
// The new returned instance of TextFieldSpecLabel is constructed
// from input parameters, 'textLabel', 'fieldLen' and
// 'textJustification'.
//
// This method is identical to TextFieldSpecLabel.NewTextLabel()
// with the sole exception being that this method returns a pointer
// to an instance of TextFieldSpecLabel and
// TextFieldSpecLabel.NewTextLabel() returns a concrete instance of
// TextFieldSpecLabel.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  textLabel                  string
//     - String content to be displayed within the text label.
//
//       If this parameter is submitted as a zero length string,
//       an error will be returned.
//
//
//  fieldLen                   int
//     - The length of the text field in which the 'textLabel' will
//       be displayed. If 'fieldLen' is less than the length of the
//       'textLabel' string, it will be automatically set equal to
//       the 'textLabel' string length.
//
//       To automatically set the value of 'fieldLen' to the length
//       of 'textLabel', set this parameter to a value of minus one
//       (-1).
//
//       If this parameter is submitted with a value less than
//       minus one (-1) or greater than 1-million (1,000,000), an
//       error will be returned.
//
//
//  textJustification          TextJustify
//     - An enumeration which specifies the justification of the
//       'textLabel' within the field specified by 'fieldLen'.
//
//       Text justification can only be evaluated in the context of
//       a text label, field length and a 'textJustification'
//       object of type TextJustify. This is because text labels
//       with a field length equal to or less than the length of
//       the text label never use text justification. In these
//       cases, text justification is completely ignored.
//
//       If the field length is greater than the length of the text
//       label, text justification must be equal to one of these
//       three valid values:
//           TextJustify(0).Left()
//           TextJustify(0).Right()
//           TextJustify(0).Center()
//
//       You can also use the abbreviated text justification
//       enumeration syntax as follows:
//
//           TxtJustify.Left()
//           TxtJustify.Right()
//           TxtJustify.Center()
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
//  *TextFieldSpecLabel
//     - This method will return a pointer to a new instance of
//       TextFieldSpecLabel constructed from information provided
//       by the input parameters.
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
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//  Example 1:
//   textLabel = "Hi There" (Length = 8)
//    fieldLen = 12
//    textJustification = TextJustify(0).Center()
//      result = "  Hi There  "
//
//  Example 2:
//   textLabel = "Hi There" (Length = 8)
//    fieldLen = 12
//    textJustification = TextJustify(0).Left()
//      result = "Hi There    "
//
//  Example 3:
//   textLabel = "Hi There" (Length = 8)
//    fieldLen = 12
//    textJustification = TextJustify(0).Right()
//      result = "    Hi There"
//
//  Example 4:
//   textLabel = "Hi There" (Length = 8)
//    fieldLen = -1
//    textJustification = TextJustify(0).Right()
//      result = "Hi There"
//
func (txtFieldLabel TextFieldSpecLabel) NewPtrTextLabel(
	textLabel string,
	fieldLen int,
	textJustification TextJustify,
	errorPrefix interface{}) (
	*TextFieldSpecLabel,
	error) {

	if txtFieldLabel.lock == nil {
		txtFieldLabel.lock = new(sync.Mutex)
	}

	txtFieldLabel.lock.Lock()

	defer txtFieldLabel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newTextLabel := TextFieldSpecLabel{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldSpecLabel.NewPtrTextLabel()",
		"")

	if err != nil {
		return &newTextLabel, err
	}

	err = textFieldSpecLabelNanobot{}.ptr().
		setTextFieldLabel(
			&newTextLabel,
			[]rune(textLabel),
			fieldLen,
			textJustification,
			ePrefix)

	return &newTextLabel, err
}

// NewPtrTextLabelRunes - Creates and returns a pointer to a new,
// fully populated instance of TextFieldSpecLabel. This type
// encapsulates a string which is formatted as a text label.
//
// The new returned instance of TextFieldSpecLabel is constructed
// from input parameters, 'textLabelChars', 'fieldLen' and
// 'textJustification'.
//
// This method is identical to
// TextFieldSpecLabel.NewTextLabelRunes() with the sole exception
// being that this method returns a pointer to an instance of
// TextFieldSpecLabel while TextFieldSpecLabel.NewTextLabelRunes()
// returns a concrete instance of TextFieldSpecLabel.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  textLabelChars             []rune
//     - An array of runes or text characters which is used to
//       generate string content for display as a text label.
//
//       If this parameter is submitted as a zero length array,
//       an error will be returned.
//
//       If this parameter is submitted with invalid zero character
//       values, an error will be returned.
//
//
//  fieldLen                   int
//     - The length of the text field in which the 'textLabelChars'
//       characters will be displayed. If 'fieldLen' is less than
//       the length of the 'textLabelChars' array, it will be
//       automatically set equal to the 'textLabelChars' array
//       length.
//
//       To automatically set the value of 'fieldLen' to the length
//       of 'textLabelChars', set this parameter to a value of
//       minus one (-1).
//
//       If this parameter is submitted with a value less than
//       minus one (-1) or greater than 1-million (1,000,000), an
//       error will be returned.
//
//
//  textJustification          TextJustify
//     - An enumeration which specifies the justification of the
//       'textLabelChars' within the text field specified by
//       'fieldLen'.
//
//       Text justification can only be evaluated in the context of
//       a text label, field length and a 'textJustification'
//       object of type TextJustify. This is because text labels
//       with a field length equal to or less than the length of
//       the text label never use text justification. In these
//       cases, text justification is completely ignored.
//
//       If the field length is greater than the length of the text
//       label, text justification must be equal to one of these
//       three valid values:
//           TextJustify(0).Left()
//           TextJustify(0).Right()
//           TextJustify(0).Center()
//
//       You can also use the abbreviated text justification
//       enumeration syntax as follows:
//
//           TxtJustify.Left()
//           TxtJustify.Right()
//           TxtJustify.Center()
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
//  *TextFieldSpecLabel
//     - This method will return a pointer to a new instance of
//       TextFieldSpecLabel constructed from information provided
//       by the input parameters.
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
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//  Example 1:
//   textLabel = "Hi There" (Length = 8)
//    fieldLen = 12
//    textJustification = TextJustify(0).Center()
//      result = "  Hi There  "
//
//  Example 2:
//   textLabel = "Hi There" (Length = 8)
//    fieldLen = 12
//    textJustification = TextJustify(0).Left()
//      result = "Hi There    "
//
//  Example 3:
//   textLabel = "Hi There" (Length = 8)
//    fieldLen = 12
//    textJustification = TextJustify(0).Right()
//      result = "    Hi There"
//
//  Example 4:
//   textLabel = "Hi There" (Length = 8)
//    fieldLen = -1
//    textJustification = TextJustify(0).Right()
//      result = "Hi There"
//
func (txtFieldLabel TextFieldSpecLabel) NewPtrTextLabelRunes(
	textLabelChars []rune,
	fieldLen int,
	textJustification TextJustify,
	errorPrefix interface{}) (
	*TextFieldSpecLabel,
	error) {

	if txtFieldLabel.lock == nil {
		txtFieldLabel.lock = new(sync.Mutex)
	}

	txtFieldLabel.lock.Lock()

	defer txtFieldLabel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newTextLabel := TextFieldSpecLabel{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldSpecLabel.NewPtrTextLabelRunes()",
		"")

	if err != nil {
		return &newTextLabel, err
	}

	err = textFieldSpecLabelNanobot{}.ptr().
		setTextFieldLabel(
			&newTextLabel,
			textLabelChars,
			fieldLen,
			textJustification,
			ePrefix)

	return &newTextLabel, err
}

// NewTextLabel - Returns a new, populated concrete instance of
// TextFieldSpecLabel. This type encapsulates a string which
// is formatted as a text label.
//
// The new returned instance of TextFieldSpecLabel is constructed
// from input parameters, 'textLabel', 'fieldLen' and
// 'textJustification'.
//
// This method is identical to TextFieldSpecLabel.NewPtrTextLabel()
// with the sole exception being that this method returns a
// concrete instance of TextFieldSpecLabel and
// TextFieldSpecLabel.NewPtrTextLabel() returns a pointer to a
// TextFieldSpecLabel instance.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  textLabel                  string
//     - The string content to be displayed within the label.
//
//       If this parameter is submitted as a zero length string,
//       an error will be returned.
//
//
//  fieldLen                   int
//     - The length of the text field in which the 'textLabel' will
//       be displayed. If 'fieldLen' is less than the length of the
//       'textLabel' string, it will be automatically set equal to
//       the 'textLabel' string length.
//
//       To automatically set the value of 'fieldLen' to the length
//       of 'textLabel', set this parameter to a value of minus one
//       (-1).
//
//       If this parameter is submitted with a value less than
//       minus one (-1) or greater than 1-million (1,000,000), an
//       error will be returned.
//
//
//  textJustification          TextJustify
//     - An enumeration which specifies the justification of the
//       'textLabel' string within the text field specified by
//       'fieldLen'.
//
//       Text justification can only be evaluated in the context of
//       a text label, field length and a Text Justification object
//       of type TextJustify. This is because text labels with a
//       field length equal to or less than the length of the text
//       label never use text justification. In these cases, text
//       justification is completely ignored.
//
//       If the field length is greater than the length of the text
//       label, text justification must be equal to one of these
//       three valid values:
//           TextJustify(0).Left()
//           TextJustify(0).Right()
//           TextJustify(0).Center()
//
//       You can also use the abbreviated text justification
//       enumeration syntax as follows:
//
//           TxtJustify.Left()
//           TxtJustify.Right()
//           TxtJustify.Center()
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
//  TextFieldSpecLabel
//     - This method will return a new, populated concrete instance
//       of TextFieldSpecLabel constructed from the information
//       provided by the input parameters.
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
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//  Example 1:
//   textLabel = "Hi There" (Length = 8)
//    fieldLen = 12
//    textJustification = TextJustify(0).Center()
//      result = "  Hi There  "
//
//  Example 2:
//   textLabel = "Hi There" (Length = 8)
//    fieldLen = 12
//    textJustification = TextJustify(0).Left()
//      result = "Hi There    "
//
//  Example 3:
//   textLabel = "Hi There" (Length = 8)
//    fieldLen = 12
//    textJustification = TextJustify(0).Right()
//      result = "    Hi There"
//
//  Example 4:
//   textLabel = "Hi There" (Length = 8)
//    fieldLen = -1
//    textJustification = TextJustify(0).Right()
//      result = "Hi There"
//
func (txtFieldLabel TextFieldSpecLabel) NewTextLabel(
	textLabel string,
	fieldLen int,
	textJustification TextJustify,
	errorPrefix interface{}) (
	TextFieldSpecLabel,
	error) {

	if txtFieldLabel.lock == nil {
		txtFieldLabel.lock = new(sync.Mutex)
	}

	txtFieldLabel.lock.Lock()

	defer txtFieldLabel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newTextLabel := TextFieldSpecLabel{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldSpecLabel.NewTextLabel()",
		"")

	if err != nil {
		return newTextLabel, err
	}

	err = textFieldSpecLabelNanobot{}.ptr().
		setTextFieldLabel(
			&newTextLabel,
			[]rune(textLabel),
			fieldLen,
			textJustification,
			ePrefix)

	return newTextLabel, err
}

// NewTextLabelRunes - Creates and returns a new, concrete instance
// of TextFieldSpecLabel. This type encapsulates a string which is
// formatted as a text label.
//
// The new returned instance of TextFieldSpecLabel is constructed
// from input parameters, 'textLabelChars', 'fieldLen' and
// 'textJustification'.
//
// This method is identical to
// TextFieldSpecLabel.NewPtrTextLabelRunes() with the sole
// exception being that this method returns a concrete instance of
// TextFieldSpecLabel and TextFieldSpecLabel.NewPtrTextLabelRunes()
// returns a pointer to a TextFieldSpecLabel instance.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  textLabelChars             []rune
//     - An array of runes or text characters which is used to
//       generate string content for display as a text label.
//
//       If this parameter is submitted as a zero length array,
//       an error will be returned.
//
//       If this parameter is submitted with invalid zero character
//       values, an error will be returned.
//
//
//  fieldLen                   int
//     - The length of the text field in which the 'textLabelChars'
//       characters will be displayed. If 'fieldLen' is less than
//       the length of the 'textLabelChars' array, it will be
//       automatically set equal to the 'textLabelChars' array
//       length.
//
//       To automatically set the value of 'fieldLen' to the length
//       of 'textLabelChars', set this parameter to a value of
//       minus one (-1).
//
//       If this parameter is submitted with a value less than
//       minus one (-1) or greater than 1-million (1,000,000), an
//       error will be returned.
//
//
//  textJustification          TextJustify
//     - An enumeration which specifies the justification of the
//       'textLabelChars' within the text field specified by
//       'fieldLen'.
//
//       Text justification can only be evaluated in the context of
//       a text label, field length and a 'textJustification'
//       object of type TextJustify. This is because text labels
//       with a field length equal to or less than the length of
//       the text label never use text justification. In these
//       cases, text justification is completely ignored.
//
//       If the field length is greater than the length of the text
//       label, text justification must be equal to one of these
//       three valid values:
//           TextJustify(0).Left()
//           TextJustify(0).Right()
//           TextJustify(0).Center()
//
//       You can also use the abbreviated text justification
//       enumeration syntax as follows:
//
//           TxtJustify.Left()
//           TxtJustify.Right()
//           TxtJustify.Center()
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
//  TextFieldSpecLabel
//     - This method will return a new concrete instance of
//       TextFieldSpecLabel constructed from information provided
//       by the input parameters.
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
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//  Example 1:
//   textLabel = "Hi There" (Length = 8)
//    fieldLen = 12
//    textJustification = TextJustify(0).Center()
//      result = "  Hi There  "
//
//  Example 2:
//   textLabel = "Hi There" (Length = 8)
//    fieldLen = 12
//    textJustification = TextJustify(0).Left()
//      result = "Hi There    "
//
//  Example 3:
//   textLabel = "Hi There" (Length = 8)
//    fieldLen = 12
//    textJustification = TextJustify(0).Right()
//      result = "    Hi There"
//
//  Example 4:
//   textLabel = "Hi There" (Length = 8)
//    fieldLen = -1
//    textJustification = TextJustify(0).Right()
//      result = "Hi There"
//
func (txtFieldLabel TextFieldSpecLabel) NewTextLabelRunes(
	textLabelChars []rune,
	fieldLen int,
	textJustification TextJustify,
	errorPrefix interface{}) (
	TextFieldSpecLabel,
	error) {

	if txtFieldLabel.lock == nil {
		txtFieldLabel.lock = new(sync.Mutex)
	}

	txtFieldLabel.lock.Lock()

	defer txtFieldLabel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newTextLabel := TextFieldSpecLabel{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldSpecLabel.NewTextLabelRunes()",
		"")

	if err != nil {
		return newTextLabel, err
	}

	err = textFieldSpecLabelNanobot{}.ptr().
		setTextFieldLabel(
			&newTextLabel,
			textLabelChars,
			fieldLen,
			textJustification,
			ePrefix)

	return newTextLabel, err
}

// Read - Implements the io.Reader interface for type
// TextFieldSpecLabel.
//
// The formatted text string generated by the current instance of
// TextFieldSpecLabel will be written to the byte buffer 'p'. If
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
//    err = txtFieldLabel01.Read(p)
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
//      "Error Returned From txtFieldLabel01.Read(p)\n"+
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
//    err = txtFieldLabel01.Read(p)
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
//      "Error Returned From txtFieldLabel01.Read(p)\n"+
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
func (txtFieldLabel *TextFieldSpecLabel) Read(
	p []byte) (
	n int,
	err error) {

	if txtFieldLabel.lock == nil {
		txtFieldLabel.lock = new(sync.Mutex)
	}

	txtFieldLabel.lock.Lock()

	defer txtFieldLabel.lock.Unlock()

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TextFieldSpecLabel.Read()",
		"")

	if txtFieldLabel.textLineReader == nil {

		var formattedText string

		formattedText,
			err = textSpecificationMolecule{}.ptr().
			getFormattedText(
				txtFieldLabel.textLabel,
				txtFieldLabel.fieldLen,
				txtFieldLabel.textJustification,
				ePrefix.XCpy(
					"txtFieldLabel"))

		if err != nil {
			return n, err
		}

		txtFieldLabel.textLineReader =
			strings.NewReader(formattedText)

		if txtFieldLabel.textLineReader == nil {
			err = fmt.Errorf("%v\n"+
				"Error: strings.NewReader(formattedText)\n"+
				"returned a nil pointer.\n"+
				"txtFieldLabel.textLineReader == nil\n",
				ePrefix.String())

			return n, err
		}
	}

	n,
		err = textSpecificationAtom{}.ptr().
		readBytes(
			txtFieldLabel.textLineReader,
			p,
			ePrefix.XCpy(
				"p -> txtFieldLabel.textLineReader"))

	if err == io.EOF {

		txtFieldLabel.textLineReader = nil

	}

	return n, err
}

// ReaderInitialize - This method will reset the internal member
// variable 'TextFieldSpecLabel.textLineReader' to its initial
// zero state of 'nil'. Effectively, this resets the internal
// strings.Reader object for use in future read operations.
//
// This method is rarely used or needed. It provides a means of
// reinitializing the internal strings.Reader object in case an
// error occurs during a read operation initiated by method
// TextFieldSpecLabel.Read().
//
// Calling this method cleans up the residue from an aborted read
// operation and prepares the strings.Reader object for future read
// operations.
//
// If any errors are returned by method
// TextFieldSpecLabel.Read() which are NOT equal to io.EOF, call
// this method, TextFieldSpecLabel.ReaderInitialize(), to reset
// and prepare the internal reader for future read operations.
//
// This method fulfills the requirements of the
// ITextFieldSpecification interface.
//
func (txtFieldLabel *TextFieldSpecLabel) ReaderInitialize() {

	if txtFieldLabel.lock == nil {
		txtFieldLabel.lock = new(sync.Mutex)
	}

	txtFieldLabel.lock.Lock()

	defer txtFieldLabel.lock.Unlock()

	txtFieldLabel.textLineReader = nil

	return
}

// SetFieldLength - Sets The length of the text field in which the
// text label string will be positioned for text display, file
// output or printing.
//
// If field length is greater than the length of text label, the
// text label string will be positioned within the text field using
// text justification specifications as shown in the following
// examples.
//
//  Example 1:
//   textLabel = "Hi There" (Length = 8)
//   fieldLen = 12
//   textJustification = TextJustify(0).Center()
//   result = "  Hi There  "
//
//  Example 2:
//   textLabel = "Hi There" (Length = 8)
//   fieldLen = 12
//   textJustification = TextJustify(0).Left()
//   result = "Hi There    "
//
//  Example 3:
//   textLabel = "Hi There" (Length = 8)
//   fieldLen = 12
//   textJustification = TextJustify(0).Right()
//   result = "    Hi There"
//
//  Example 4:
//   textLabel = "Hi There" (Length = 8)
//   fieldLen = -1
//   textJustification = TextJustify(0).Right()
//   result = "Hi There"
//
// For more information on Text Justification, see method:
//   TextFieldSpecLabel.SetTextJustification()
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  fieldLen                   int
//     - The length of the text field in which the 'textLabelChars'
//       will be displayed. If 'fieldLen' is less than the length
//       of the 'textLabelChars' array, it will be automatically
//       set equal to the 'textLabelChars' array length.
//
//       To automatically set the value of 'fieldLen' to the length
//       of 'textLabelChars', set this parameter to a value of
//       minus one (-1).
//
//       If this parameter is submitted with a value less than
//       minus one (-1), an error will be returned.
//
//       If 'fieldLen' is submitted with a value greater than
//       1-million (1,000,000), an error will be returned.
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
func (txtFieldLabel *TextFieldSpecLabel) SetFieldLength(
	fieldLen int,
	errorPrefix interface{}) error {

	if txtFieldLabel.lock == nil {
		txtFieldLabel.lock = new(sync.Mutex)
	}

	txtFieldLabel.lock.Lock()

	defer txtFieldLabel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldSpecLabel.SetFieldLength()",
		"")

	if err != nil {
		return err
	}

	err =
		textFieldSpecLabelElectron{}.ptr().
			isFieldLengthValid(
				fieldLen,
				ePrefix.XCpy(
					"'fieldLen' Invalid!"))

	if err != nil {
		return err
	}

	txtFieldLabel.fieldLen = fieldLen

	return err
}

// SetText - Sets the text string which will be used as the text
// label for this instance of TextFieldSpecLabel.
//
// When the text label is formatted for output, the existing field
// length, and text justification parameters will be applied.
//
// If input parameter 'textLabel' is submitted as an empty string,
// an error will be returned.
//
// This method is identical to TextFieldSpecLabel.SetTextRunes()
// with the sole difference being that this method receives a
// string as an input parameter.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  textLabel                  string
//     - String content to be displayed within the text label.
//
//       If this parameter is submitted as a zero length string,
//       an error will be returned.
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
//     - If this method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (txtFieldLabel *TextFieldSpecLabel) SetText(
	textLabel string,
	errorPrefix interface{}) (
	err error) {

	if txtFieldLabel.lock == nil {
		txtFieldLabel.lock = new(sync.Mutex)
	}

	txtFieldLabel.lock.Lock()

	defer txtFieldLabel.lock.Unlock()

	txtFieldLabel.textLabel = []rune(textLabel)

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldSpecLabel.SetText()",
		"")

	if err != nil {
		return err
	}

	if len(textLabel) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'textLabel' is invalid!\n"+
			"'textLabel' is a zero length string.\n",
			ePrefix.String())

		return err
	}

	textLabelRunes := []rune(textLabel)

	sMechPreon := strMechPreon{}
	_,
		err = sMechPreon.testValidityOfRuneCharArray(
		textLabelRunes,
		ePrefix.XCpy("textLabel->"+
			"textLabelRunes"))

	if err != nil {
		return err
	}

	err = sMechPreon.copyRuneArrays(
		&txtFieldLabel.textLabel,
		&textLabelRunes,
		true,
		ePrefix.XCpy(
			"textLabel->textLabelRunes->"+
				"txtFieldLabel.textLabel"))

	return err
}

// SetTextJustification - Sets the text justification specification
// for the current instance of TextFieldSpecLabel.
//
// TextJustify is An enumeration which specifies the justification
// of the text field label string within a text field. The calling
// function sets the text justification specification by passing an
// appropriate TextJustify value.
//
// For more information on the text label, see methods:
//  TextFieldSpecLabel.SetText()
//  TextFieldSpecLabel.SetTextRunes()
//
// For more information on text field length, see method:
//  TextFieldSpecLabel.SetFieldLength()
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  textJustification          TextJustify
//     - An enumeration which specifies the justification of the
//       Text label string within the text field specified by
//       field length.
//
//       Text justification can only be evaluated in the context of
//       a Text label, field length and a 'textJustification'
//       object of type TextJustify. This is because text labels
//       with a field length equal to or less than the length of
//       the text label never use text justification. In these
//       cases, text justification is completely ignored.
//
//       This method requires that text justification be set to one
//       of the following three values:
//           TextJustify(0).Left()
//           TextJustify(0).Right()
//           TextJustify(0).Center()
//
//       You can also use the abbreviated text justification
//       enumeration syntax as follows:
//
//           TxtJustify.Left()
//           TxtJustify.Right()
//           TxtJustify.Center()
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
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//  Example 1:
//   textLabel = "Hi There" (Length = 8)
//    fieldLen = 12
//    textJustification = TextJustify(0).Center()
//      result = "  Hi There  "
//
//  Example 2:
//   textLabel = "Hi There" (Length = 8)
//    fieldLen = 12
//    textJustification = TextJustify(0).Left()
//      result = "Hi There    "
//
//  Example 3:
//   textLabel = "Hi There" (Length = 8)
//    fieldLen = 12
//    textJustification = TextJustify(0).Right()
//      result = "    Hi There"
//
//  Example 4:
//   textLabel = "Hi There" (Length = 8)
//    fieldLen = -1
//    textJustification = TextJustify(0).Right()
//      result = "Hi There"
//
func (txtFieldLabel *TextFieldSpecLabel) SetTextJustification(
	textJustification TextJustify,
	errorPrefix interface{}) error {

	if txtFieldLabel.lock == nil {
		txtFieldLabel.lock = new(sync.Mutex)
	}

	txtFieldLabel.lock.Lock()

	defer txtFieldLabel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldSpecLabel.SetFieldLength()",
		"")

	if err != nil {
		return err
	}

	if !textJustification.XIsValid() {
		err = fmt.Errorf("%v\n"+
			"Error: Text Justification is INVALID!\n"+
			"Text Justification MUST be set to\n"+
			"Left, Right or Center.\n"+
			"Text Justification  String Value= '%v'\n"+
			"Text Justification Integer Value= '%v'\n",
			ePrefix.String(),
			textJustification.String(),
			textJustification.XValueInt())

		return err
	}

	txtFieldLabel.textJustification =
		textJustification

	return nil
}

// SetTextLabel - Sets the text label component values for the
// current instance of TextFieldSpecLabel.
//
// IMPORTANT
//
// This method will overwrite and delete the existing data values
// for the current TextFieldSpecLabel instance (txtFieldLabel).
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  textLabel                  string
//     - String content to be displayed within the text label.
//
//       If this parameter is submitted as a zero length string,
//       an error will be returned.
//
//
//  fieldLen                   int
//     - The length of the text field in which the 'textLabel' will
//       be displayed. If 'fieldLen' is less than the length of the
//       'textLabel' string, it will be automatically set equal to
//       the 'textLabel' string length.
//
//       To automatically set the value of 'fieldLen' to the length
//       of 'textLabel', set this parameter to a value of minus one
//       (-1).
//
//       If this parameter is submitted with a value less than
//       minus one (-1) or greater than 1-million (1,000,000), an
//       error will be returned.
//
//
//  textJustification          TextJustify
//     - An enumeration which specifies the justification of the
//       'textLabel' within the field specified by 'fieldLen'.
//
//       Text justification can only be evaluated in the context of
//       a text label, field length and a 'textJustification'
//       object of type TextJustify. This is because text labels
//       with a field length equal to or less than the length of
//       the text label never use text justification. In these
//       cases, text justification is completely ignored.
//
//       If the field length is greater than the length of the text
//       label, text justification must be equal to one of these
//       three valid values:
//           TextJustify(0).Left()
//           TextJustify(0).Right()
//           TextJustify(0).Center()
//
//       You can also use the abbreviated text justification
//       enumeration syntax as follows:
//
//           TxtJustify.Left()
//           TxtJustify.Right()
//           TxtJustify.Center()
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
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//  Example 1:
//   textLabel = "Hi There" (Length = 8)
//    fieldLen = 12
//    textJustification = TextJustify(0).Center()
//      result = "  Hi There  "
//
//  Example 2:
//   textLabel = "Hi There" (Length = 8)
//    fieldLen = 12
//    textJustification = TextJustify(0).Left()
//      result = "Hi There    "
//
//  Example 3:
//   textLabel = "Hi There" (Length = 8)
//    fieldLen = 12
//    textJustification = TextJustify(0).Right()
//      result = "    Hi There"
//
//  Example 4:
//   textLabel = "Hi There" (Length = 8)
//    fieldLen = -1
//    textJustification = TextJustify(0).Right()
//      result = "Hi There"
//
func (txtFieldLabel *TextFieldSpecLabel) SetTextLabel(
	textLabel string,
	fieldLen int,
	textJustification TextJustify,
	errorPrefix interface{}) error {

	if txtFieldLabel.lock == nil {
		txtFieldLabel.lock = new(sync.Mutex)
	}

	txtFieldLabel.lock.Lock()

	defer txtFieldLabel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldSpecLabel.SetTextLabel()",
		"")

	if err != nil {
		return err
	}

	return textFieldSpecLabelNanobot{}.ptr().
		setTextFieldLabel(
			txtFieldLabel,
			[]rune(textLabel),
			fieldLen,
			textJustification,
			ePrefix)
}

// SetTextLabelRunes - Sets the text label component values for
// the current instance of TextFieldSpecLabel. The input parameter
// required to set the text label characters is submitted as an
// array of runes.
//
// IMPORTANT
//
// This method will overwrite and delete the existing data values
// for the current TextFieldSpecLabel instance (txtFieldLabel).
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  textLabelChars             []rune
//     - An array of runes or text characters which is used to
//       generate string content for display as a text label.
//
//       If this parameter is submitted as a zero length array,
//       an error will be returned.
//
//
//  fieldLen                   int
//     - The length of the text field in which the 'textLabelChars'
//       will be displayed. If 'fieldLen' is less than the length
//       of the 'textLabelChars' array, it will be automatically
//       set equal to the 'textLabelChars' array length.
//
//       To automatically set the value of 'fieldLen' to the length
//       of 'textLabelChars', set this parameter to a value of
//       minus one (-1).
//
//       If this parameter is submitted with a value less than
//       minus one (-1) or greater than 1-million (1,000,000), an
//       error will be returned.
//
//
//  textJustification          TextJustify
//     - An enumeration which specifies the justification of the
//       'textLabelChars' within the field specified by 'fieldLen'.
//
//       Text justification can only be evaluated in the context of
//       a text label, field length and a 'textJustification'
//       object of type TextJustify. This is because text labels
//       with a field length equal to or less than the length of
//       the text label never use text justification. In these
//       cases, text justification is completely ignored.
//
//       If the field length is greater than the length of the text
//       label, text justification must be equal to one of these
//       three valid values:
//           TextJustify(0).Left()
//           TextJustify(0).Right()
//           TextJustify(0).Center()
//
//       You can also use the abbreviated text justification
//       enumeration syntax as follows:
//
//           TxtJustify.Left()
//           TxtJustify.Right()
//           TxtJustify.Center()
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
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//  Example 1:
//   textLabel = "Hi There" (Length = 8)
//    fieldLen = 12
//    textJustification = TextJustify(0).Center()
//      result = "  Hi There  "
//
//  Example 2:
//   textLabel = "Hi There" (Length = 8)
//    fieldLen = 12
//    textJustification = TextJustify(0).Left()
//      result = "Hi There    "
//
//  Example 3:
//   textLabel = "Hi There" (Length = 8)
//    fieldLen = 12
//    textJustification = TextJustify(0).Right()
//      result = "    Hi There"
//
//  Example 4:
//   textLabel = "Hi There" (Length = 8)
//    fieldLen = -1
//    textJustification = TextJustify(0).Right()
//      result = "Hi There"
//
func (txtFieldLabel *TextFieldSpecLabel) SetTextLabelRunes(
	textLabelChars []rune,
	fieldLen int,
	textJustification TextJustify,
	errorPrefix interface{}) error {

	if txtFieldLabel.lock == nil {
		txtFieldLabel.lock = new(sync.Mutex)
	}

	txtFieldLabel.lock.Lock()

	defer txtFieldLabel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldSpecLabel.SetTextLabelRunes()",
		"")

	if err != nil {
		return err
	}

	return textFieldSpecLabelNanobot{}.ptr().
		setTextFieldLabel(
			txtFieldLabel,
			textLabelChars,
			fieldLen,
			textJustification,
			ePrefix)
}

// SetTextRunes - Sets the text characters which will be used as
// the text label for this instance of TextFieldSpecLabel.
//
// When the text label is formatted for output, the existing field
// length, and text justification parameters will be applied.
//
// This method is identical to TextFieldSpecLabel.SetText()
// with the sole difference being that this method receives an
// array of runes as an input parameter.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  textLabelChars             []rune
//     - An array of runes or text characters which is used to
//       generate string content for display as a text label.
//
//       If this parameter is submitted as a zero length array,
//       an error will be returned.
//
//       If this parameter is submitted with invalid zero character
//       values, an error will be returned.
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
//     - If this method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (txtFieldLabel *TextFieldSpecLabel) SetTextRunes(
	textLabelChars []rune,
	errorPrefix interface{}) (
	err error) {

	if txtFieldLabel.lock == nil {
		txtFieldLabel.lock = new(sync.Mutex)
	}

	txtFieldLabel.lock.Lock()

	defer txtFieldLabel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldSpecLabel.SetTextRunes()",
		"")

	if err != nil {
		return err
	}

	sMechPreon := strMechPreon{}
	_,
		err = sMechPreon.testValidityOfRuneCharArray(
		textLabelChars,
		ePrefix.XCpy("textLabelChars"))

	if err != nil {
		return err
	}

	err = sMechPreon.copyRuneArrays(
		&txtFieldLabel.textLabel,
		&textLabelChars,
		true,
		ePrefix.XCpy(
			"textLabelChars->txtFieldLabel.textLabel"))

	return err
}

// String - Returns the formatted text generated by the
// current instance of TextFieldSpecLabel.
//
// If the length of the text label string is zero and the field
// length is zero this method returns an empty string.
//
// If the length of the text label string is zero and the field
// length is greater than zero, this method returns a string with
// a length equal to field length and content equal to white space
// (the space character " " x field length).
//
// This method is identical in function to
// TextFieldSpecLabel.GetFormattedText()
//
// This method fulfills the requirements of the
// ITextFieldSpecification interface.
//
// This method also fulfills the requirements of the 'Stringer'
// interface defined in the Golang package 'fmt'. Reference:
//   https://pkg.go.dev/fmt#Stringer
//
// Methods which return formatted text are listed as follows:
//  TextFieldSpecLabel.String()
//  TextFieldSpecLabel.GetFormattedText()
//  TextFieldSpecLabel.TextBuilder()
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//  Example 1:
//   textLabel = "Hi There" (Length = 8)
//    fieldLen = 12
//    textJustification = TextJustify(0).Center()
//      result = "  Hi There  "
//
//  Example 2:
//   textLabel = "Hi There" (Length = 8)
//    fieldLen = 12
//    textJustification = TextJustify(0).Left()
//      result = "Hi There    "
//
//  Example 3:
//   textLabel = "Hi There" (Length = 8)
//    fieldLen = 12
//    textJustification = TextJustify(0).Right()
//      result = "    Hi There"
//
//  Example 4:
//   textLabel = "Hi There" (Length = 8)
//    fieldLen = -1
//    textJustification = TextJustify(0).Right()
//      result = "Hi There"
//
func (txtFieldLabel TextFieldSpecLabel) String() string {

	if txtFieldLabel.lock == nil {
		txtFieldLabel.lock = new(sync.Mutex)
	}

	txtFieldLabel.lock.Lock()

	defer txtFieldLabel.lock.Unlock()

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TextFieldSpecLabel.String()",
		"")

	result,
		err := textSpecificationMolecule{}.ptr().
		getFormattedText(
			txtFieldLabel.textLabel,
			txtFieldLabel.fieldLen,
			txtFieldLabel.textJustification,
			&ePrefix)

	if err != nil {
		result = fmt.Sprintf("%v",
			err.Error())
	}

	return result
}

// TextBuilder - Configures the line of text produced by this
// instance of TextFieldSpecLabel, and writes it to an instance
// of strings.Builder.
//
// This method fulfills the requirements of the
// ITextFieldSpecification interface.
//
// Methods which return formatted text are listed as follows:
//  TextFieldSpecLabel.String()
//  TextFieldSpecLabel.GetFormattedText()
//  TextFieldSpecLabel.TextBuilder()
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  sBuilder                   *strings.Builder
//    - A pointer to an instance of strings.Builder. The line of
//      text produced by the current instance of
//      TextFieldSpecLabel and writes that text to 'sBuilder'.
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
func (txtFieldLabel *TextFieldSpecLabel) TextBuilder(
	sBuilder *strings.Builder,
	errorPrefix interface{}) error {

	if txtFieldLabel.lock == nil {
		txtFieldLabel.lock = new(sync.Mutex)
	}

	txtFieldLabel.lock.Lock()

	defer txtFieldLabel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldSpecLabel.TextBuilder()",
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
		err = textSpecificationMolecule{}.ptr().
		getFormattedText(
			txtFieldLabel.textLabel,
			txtFieldLabel.fieldLen,
			txtFieldLabel.textJustification,
			ePrefix.XCpy(
				"txtFieldLabel"))

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
func (txtFieldLabel *TextFieldSpecLabel) TextFieldName() string {

	if txtFieldLabel.lock == nil {
		txtFieldLabel.lock = new(sync.Mutex)
	}

	txtFieldLabel.lock.Lock()

	defer txtFieldLabel.lock.Unlock()

	return "Label"
}

// TextTypeName - returns a string specifying the type of Text
// Field specification.
//
// This method fulfills the requirements of the ITextSpecification
// interface.
//
func (txtFieldLabel *TextFieldSpecLabel) TextTypeName() string {

	if txtFieldLabel.lock == nil {
		txtFieldLabel.lock = new(sync.Mutex)
	}

	txtFieldLabel.lock.Lock()

	defer txtFieldLabel.lock.Unlock()

	return "TextFieldSpecLabel"
}
