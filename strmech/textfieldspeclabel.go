package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type TextFieldSpecLabel struct {
	textLabel         []rune
	fieldLen          int
	textJustification TextJustify
	lock              *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming instance of
// TextFieldSpecLabel ('incomingTxtFieldLabel') to the data fields
// of the current TextFieldSpecLabel instance ('txtFieldLabel').
//
// IMPORTANT
// All of the data fields in current TextFieldSpecLabel instance
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
func (txtFieldLabel *TextFieldSpecLabel) CopyOut() TextFieldSpecLabel {

	if txtFieldLabel.lock == nil {
		txtFieldLabel.lock = new(sync.Mutex)
	}

	txtFieldLabel.lock.Lock()

	defer txtFieldLabel.lock.Unlock()

	newTxtFieldLabel,
		_ := textFieldSpecLabelMolecule{}.ptr().
		copyOut(txtFieldLabel,
			nil)

	return newTxtFieldLabel
}

// CopyOutITextField - Returns a deep copy of the current
// TextFieldSpecLabel instance cast as an ITextFieldSpecification
// object.
//
func (txtFieldLabel *TextFieldSpecLabel) CopyOutITextField() ITextFieldSpecification {

	if txtFieldLabel.lock == nil {
		txtFieldLabel.lock = new(sync.Mutex)
	}

	txtFieldLabel.lock.Lock()

	defer txtFieldLabel.lock.Unlock()

	newTxtFieldLabel,
		_ := textFieldSpecLabelMolecule{}.ptr().
		copyOut(txtFieldLabel,
			nil)

	iTxtFieldSpec := ITextFieldSpecification(&newTxtFieldLabel)

	return iTxtFieldSpec
}

// CopyOutPtr - Returns a pointer to a deep copy of the current
// TextFieldSpecLabel instance.
//
func (txtFieldLabel *TextFieldSpecLabel) CopyOutPtr() *TextFieldSpecLabel {

	if txtFieldLabel.lock == nil {
		txtFieldLabel.lock = new(sync.Mutex)
	}

	txtFieldLabel.lock.Lock()

	defer txtFieldLabel.lock.Unlock()

	newTxtFieldLabel,
		_ := textFieldSpecLabelMolecule{}.ptr().
		copyOut(txtFieldLabel,
			nil)

	return &newTxtFieldLabel
}

// Empty - Resets all internal member variables to their initial
// or zero states.
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
// in all respects, this flag is set to 'true'. Otherwise, this method returns
// 'false'.
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
func (txtFieldLabel *TextFieldSpecLabel) GetFieldLength() int {

	if txtFieldLabel.lock == nil {
		txtFieldLabel.lock = new(sync.Mutex)
	}

	txtFieldLabel.lock.Lock()

	defer txtFieldLabel.lock.Unlock()

	return txtFieldLabel.fieldLen
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
func (txtFieldLabel *TextFieldSpecLabel) GetFormattedText() string {

	if txtFieldLabel.lock == nil {
		txtFieldLabel.lock = new(sync.Mutex)
	}

	txtFieldLabel.lock.Lock()

	defer txtFieldLabel.lock.Unlock()

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TextFieldSpecLabel.GetFormattedText()",
		"")

	result := textSpecificationMolecule{}.ptr().
		getFormattedText(
			txtFieldLabel.textLabel,
			txtFieldLabel.fieldLen,
			txtFieldLabel.textJustification,
			&ePrefix)

	return result
}

// GetTextJustification - Returns the value of the text
// justification enumeration specification which will be used to
// position the text label string with a text field.
//
// The text justification enumeration specification should be set
// to one of three values:
//           TextJustify(0).Left()
//           TextJustify(0).Right()
//           TextJustify(0).Center()
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
func (txtFieldLabel *TextFieldSpecLabel) GetTextLabel() string {

	if txtFieldLabel.lock == nil {
		txtFieldLabel.lock = new(sync.Mutex)
	}

	txtFieldLabel.lock.Lock()

	defer txtFieldLabel.lock.Unlock()

	return string(txtFieldLabel.textLabel)
}

// IsValidInstanceError - Performs a diagnostic review of the data
// values encapsulated in the current TextFieldSpecLabel instance
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
//       the current instance of TextFieldSpecLabel are found to be
//       invalid, this method will return an error.
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

	lenTextLabel := len(txtFieldLabel.textLabel)

	if lenTextLabel == 0 && txtFieldLabel.fieldLen == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Text Label is a zero length string AND\n"+
			"Field Length is also zero!\n",
			ePrefix.String())

		return err
	}

	txtJustificationIsValid := txtFieldLabel.textJustification.XIsValid()

	if txtFieldLabel.fieldLen > lenTextLabel &&
		!txtJustificationIsValid {
		err = fmt.Errorf("%v\n"+
			"Error: Text Justification is INVALID!\n"+
			"Text Lable Length = '%v'\n"+
			"Field Length = '%v'\n"+
			"Text Justification Integer Value = '%v'\n",
			ePrefix.String(),
			lenTextLabel,
			txtFieldLabel.fieldLen,
			txtFieldLabel.textJustification.XValueInt())

		return err
	}

	return nil
}

// NewConstructor - Creates and returns a new, fully populated
// instance of TextFieldSpecLabel. This type encapsulates a string
// which is formatted as a text label.
//
// The new returned instance of TextFieldSpecLabel is constructed
// from input parameters, 'textLabel', 'fieldLen' and
// 'textJustification'.
//
// This method is identical to TextFieldSpecLabel.NewPtr()
// with the sole distinction being that this method returns a
// new concrete instance of TextFieldSpecLabel.
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
//     - The length of the field in which the 'textLabel' will be
//       displayed. If 'fieldLen' is less than the length of the
//       'textLabel' string, it will be automatically set equal to
//       the 'textLabel' string length.
//
//       If this parameter is submitted with a value less than
//       minus one (-1) or greater than 1-million (1,000,000), an
//       error will be returned.
//
//
//  textJustification          TextJustify
//     - An enumeration which specifies the justification of the
//       'textLabel' within the field specified by 'fieldLen'.
//       Options for 'textJustification' are:
//           TextJustify(0).Left()
//           TextJustify(0).Right()
//           TextJustify(0).Center()
//
//       If input parameter 'textJustification' is not equal to one
//       of the three values listed above, an error will be returned.
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
//     - This method will return a new instance of
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
func (txtFieldLabel TextFieldSpecLabel) NewConstructor(
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

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldSpecFiller.NewConstructor()",
		"")

	if err != nil {
		return TextFieldSpecLabel{}, err
	}

	textRunes := []rune(textLabel)

	lenTxtRunes := len(textRunes)

	if lenTxtRunes == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'textLabel' is a zero length string!\n",
			ePrefix.String())
		return TextFieldSpecLabel{}, err
	}

	if fieldLen < -1 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fieldLen' is less than minus one (-1)!\n"+
			"'fieldLen' controls the length of the Text Label Field.\n",
			ePrefix.String())

		return TextFieldSpecLabel{}, err
	}

	if fieldLen > 1000000 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fieldLen' is greater than one-million (1,000,000)!\n"+
			"'fieldLen' controls the length of the Text Label Field.\n",
			ePrefix.String())
		return TextFieldSpecLabel{}, err
	}

	if !textJustification.XIsValid() {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'textJustification' is INVALID!\n"+
			"'textJustification' must be equal to 'Left', 'Center' or 'Right'.\n"+
			"'textJustification'='%v'\n"+
			"'textJustification' integer value ='%v'\n",
			ePrefix.String(),
			textJustification.String(),
			textJustification.XValueInt())

		return TextFieldSpecLabel{}, err
	}

	newTextLabel := TextFieldSpecLabel{}

	newTextLabel.textLabel = make([]rune, lenTxtRunes)

	copy(newTextLabel.textLabel,
		textRunes)

	if fieldLen < len(textLabel) {
		fieldLen = len(textLabel)
	}

	newTextLabel.fieldLen = fieldLen

	newTextLabel.textJustification = textJustification

	newTextLabel.lock = new(sync.Mutex)

	return newTextLabel, nil
}

// NewEmpty - Returns a pointer to a new, empty instance of
// TextFieldSpecLabel. The member variables of this returned
// instance will all be set to their native zero values.
//
func (txtFieldLabel TextFieldSpecLabel) NewEmpty() *TextFieldSpecLabel {

	if txtFieldLabel.lock == nil {
		txtFieldLabel.lock = new(sync.Mutex)
	}

	txtFieldLabel.lock.Lock()

	defer txtFieldLabel.lock.Unlock()

	newTxtFieldLabel := TextFieldSpecLabel{}

	newTxtFieldLabel.textJustification = TextJustify(0).None()

	newTxtFieldLabel.lock = new(sync.Mutex)

	return &newTxtFieldLabel
}

// NewPtr - Returns a pointer to a new, populated instance
// of TextFieldSpecLabel. This type encapsulates a string which
// is formatted as a text label.
//
// The new returned instance of TextFieldSpecLabel is constructed
// from input parameters, 'textLabel', 'fieldLen' and
// 'textJustification'.
//
// This method is identical to TextFieldSpecLabel.NewConstructor()
// with the sole distinction being that this method returns a
// pointer to the new instance of TextFieldSpecLabel.
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
//     - The length of the field in which the 'textLabel' will be
//       displayed. If 'fieldLen' is less than the length of the
//       'textLabel' string, it will be automatically set equal to
//       the 'textLabel' string length.
//
//       If this parameter is submitted with a value less than
//       minus one (-1) or greater than 1-million (1,000,000), an
//       error will be returned.
//
//
//  textJustification          TextJustify
//     - An enumeration which specifies the justification of the
//       'textLabel' within the field specified by 'fieldLen'.
//       Options for 'textJustification' are:
//           TextJustify(0).Left()
//           TextJustify(0).Right()
//           TextJustify(0).Center()
//
//       If input parameter 'textJustification' is not equal to one
//       of the three values listed above, an error will be returned.
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
//       TextFieldSpecLabel constructed from the information
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
func (txtFieldLabel TextFieldSpecLabel) NewPtr(
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

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldSpecFiller.NewPtr()",
		"")

	if err != nil {
		return &TextFieldSpecLabel{}, err
	}

	textRunes := []rune(textLabel)

	lenTxtRunes := len(textRunes)

	if lenTxtRunes == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'textLabel' is a zero length string!\n",
			ePrefix.String())
		return &TextFieldSpecLabel{}, err
	}

	if fieldLen < -1 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fieldLen' is less than minus one (-1)!\n"+
			"'fieldLen' controls the length of the Text Label Field.\n",
			ePrefix.String())

		return &TextFieldSpecLabel{}, err
	}

	if fieldLen > 1000000 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fieldLen' is greater than one-million (1,000,000)!\n"+
			"'fieldLen' controls the length of the Text Label Field.\n",
			ePrefix.String())
		return &TextFieldSpecLabel{}, err
	}

	if !textJustification.XIsValid() {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'textJustification' is INVALID!\n"+
			"'textJustification' must be equal to 'Left', 'Center' or 'Right'.\n"+
			"'textJustification'='%v'\n"+
			"'textJustification' integer value ='%v'\n",
			ePrefix.String(),
			textJustification.String(),
			textJustification.XValueInt())

		return &TextFieldSpecLabel{}, err
	}

	newTextLabel := TextFieldSpecLabel{}

	newTextLabel.textLabel = make([]rune, lenTxtRunes)

	copy(newTextLabel.textLabel, textRunes)

	if fieldLen < len(textLabel) {
		fieldLen = len(textLabel)
	}

	newTextLabel.fieldLen = fieldLen

	newTextLabel.textJustification = textJustification

	newTextLabel.lock = new(sync.Mutex)

	return &newTextLabel, nil
}

// SetFieldLength - Sets and replaces the current value of text
// field length in the current instance of TextFieldSpecLabel.
//
// The field length specifies the length of the text field in which
// the text label string will be positioned. If the field length is
// greater than the Text Label string length, the field will be
// padded with spaces according to the Text Justification setting.
//
// If the field length is less than the length of the text label
// string, it will automatically reset to the length of the text
// label string.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  fieldLen                   int
//     - The length of the field in which the 'textLabel' will be
//       displayed.
//
//       If this parameter is submitted with a value less than
//       minus one (-1) or greater than 1-million (1,000,000), an
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
		"TextFieldSpecFiller.SetFieldLength()",
		"")

	if err != nil {
		return err
	}

	if fieldLen < -1 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fieldLen' is less than minus one (-1)!\n"+
			"'fieldLen' controls the length of the Text Label Field.\n",
			ePrefix.String())

		return err
	}

	if fieldLen > 1000000 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fieldLen' is greater than one-million (1,000,000)!\n"+
			"'fieldLen' controls the length of the Text Label Field.\n",
			ePrefix.String())
		return err
	}

	txtFieldLabel.fieldLen = fieldLen

	return nil
}

// SetTextJustification - Sets and replaces the current value of
// the Text Justification enumeration specification. This
// specification will be used to position the text label string
// within a text field.
//
// The field length specifies the length of the text field in which
// the text label string will be positioned. If the field length is
// greater than the Text Label string length, the field will be
// padded with spaces according to the Text Justification setting.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  textJustification          TextJustify
//     - An enumeration which specifies the justification of the
//       'textLabel' within the field specified by 'fieldLen'.
//       Options for 'textJustification' are:
//           TextJustify(0).Left()
//           TextJustify(0).Right()
//           TextJustify(0).Center()
//
//       If input parameter 'textJustification' is not equal to one
//       of the three values listed above, an error will be returned.
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
		"TextFieldSpecFiller.SetTextJustification()",
		"")

	if err != nil {
		return err
	}

	if !textJustification.XIsValid() {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'textJustification' is INVALID!\n"+
			"'textJustification' must be equal to 'Left', 'Center' or 'Right'.\n"+
			"'textJustification'='%v'\n"+
			"'textJustification' integer value ='%v'\n",
			ePrefix.String(),
			textJustification.String(),
			textJustification.XValueInt())

		return err
	}

	txtFieldLabel.textJustification = textJustification

	return nil
}

// SetTextLabel - Sets the text label string associated with the
// current instance of TextFieldSpecLabel.
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
func (txtFieldLabel *TextFieldSpecLabel) SetTextLabel(
	textLabel string,
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
		"TextFieldSpecFiller.SetTextLabel()",
		"")

	if err != nil {
		return err
	}

	textRunes := []rune(textLabel)

	lenTxtRunes := len(textRunes)

	if lenTxtRunes == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'textLabel' is a zero length string!\n",
			ePrefix.String())
		return err
	}

	txtFieldLabel.textLabel = make([]rune, lenTxtRunes)

	copy(txtFieldLabel.textLabel, textRunes)

	return nil
}

// SetTextLabelRunes - Sets the text label string associated with
// the current instance of TextFieldSpecLabel. The input parameter
// required to set this text label is submitted as an array of
// runes.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  textLabelRunes             []rune
//     - The string content to be displayed within the label will
//       be created from this rune array.
//
//       If this parameter is submitted as a zero length array of,
//       runes, an error will be returned.
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
func (txtFieldLabel *TextFieldSpecLabel) SetTextLabelRunes(
	textLabelRunes []rune,
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
		"TextFieldSpecFiller.SetTextLabelRunes()",
		"")

	if err != nil {
		return err
	}

	lenTxtRunes := len(textLabelRunes)

	if lenTxtRunes == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'textLabelRunes' is a zero length array!\n",
			ePrefix.String())
		return err
	}

	txtFieldLabel.textLabel = make([]rune, lenTxtRunes)

	copy(txtFieldLabel.textLabel, textLabelRunes)

	return nil
}

// TextFieldName - returns a string specifying the name
// of the Text Field specification. This method fulfills
// the requirements of the ITextFieldSpecification interface.
//
func (txtFieldLabel TextFieldSpecLabel) TextFieldName() string {

	return "TextFieldSpecLabel"
}

// TextTypeName - returns a string specifying the type
// of Text Field specification. This method fulfills
// the requirements of the ITextSpecification interface.
//
func (txtFieldLabel TextFieldSpecLabel) TextTypeName() string {

	return "TextFieldSpecLabel"
}
