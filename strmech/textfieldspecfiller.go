package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// TextFieldSpecFiller - The Filler Text Field Specification is a
// single character which is replicated for the entire length of
// the Filler Text Field.
//
// Text Field Specifications are designed to be configured within a
// line of text. Those lines of text can then be formatted for text
// displays, file output or printing. Type TextLineSpecStandardLine
// can be used to compose a line of text consisting of multiple Text
// Field Specifications like TextFieldSpecFiller. Text Field
// Specifications are therefore used as the components or building
// blocks for constructing a single lines of text.
//
// Typically, filler fields are used as margins containing multiple
// white space characters, or line separators containing multiple
// dashes, equal signs or underscore characters. Filler fields
// consist of filler characters ('fillerCharacters') and the filler
// characters repeat count ('fillerCharsRepeatCount'). A filler
// field is made up of one or more filler characters. These filler
// characters are repeated one or more times in order to construct
// the complete filler field as shown in the following examples:
//
//  Example 1:
//   Filler Characters = "-"
//   Filler Characters Repeat Count = 3
//   Formatted Text = "---"
//
//  Example 2:
//   Filler Characters = "-*"
//   Filler Characters Repeat Count = 3
//   Formatted Text = "-*-*-*"
//
// The 'fillerCharsRepeatCount' integer value is the number times
// that 'fillerCharacters' is repeated in order to construct the
// Filler Text Field. Be advised that Filler Text Fields requires a
// 'fillerCharsRepeatCount' value greater than zero.
// 'fillerCharsRepeatCount' values less than or equal to zero
// constitute an error condition.
//
// Member Variables
//
// ----------------------------------------------------------------
//
//  fillerCharacters           []rune
//     - A rune array containing the text characters which will be
//       included in the Text Filler Field. The final Text Filler
//       Field will be constructed from ths filler characters
//       repeated one or more times as specified by the
//       'fillerCharsRepeatCount' parameter.
//
//       The Text Field Filler final formatted text is equal to:
//          fillerCharacters X fillerCharsRepeatCount
//          Example: fillerCharacters = []rune{'-','*'}
//                   fillerRepeatCount = 3
//                   Final Text Filler Field = "-*-*-*"
//
//
//  fillerCharsRepeatCount     int
//     - Controls the number of times 'fillerCharacters' is
//       repeated when constructing the final Text Filler Field
//       returned by this method. The actual length of the string
//       which will populated the completed Text Filler Field is
//       equal to the length of 'fillerCharacters' times the value
//       of 'fillerCharsRepeatCount'.
//
//        Text Field Filler Length =
//          Length of fillerCharacters X fillerCharsRepeatCount
//
//          Example #1: fillerCharacters = []rune{'-','*'}
//                      fillerRepeatCount = 3
//                      Final Text Filler Field = "-*-*-*"
//
//          Example #2: fillerCharacters = []rune{'-'}
//                      fillerRepeatCount = 3
//                      Final Text Filler Field = "---"
//
type TextFieldSpecFiller struct {
	fillerCharacters []rune // The base characters which comprise the text filler
	//                                  //   field. See 'fillerCharsRepeatCount'.
	fillerCharsRepeatCount int // The number of times 'fillerCharacters'
	//                                  //  is repeated to create the complete filler string.
	lock *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming instance of
// TextFieldSpecFiller ('incomingTxtFieldFiller') to the data fields
// of the current TextFieldSpecFiller instance ('txtFillerField').
//
// IMPORTANT
// All the data fields in current TextFieldSpecFiller instance
// ('txtFillerField') will be modified and overwritten.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingTxtFieldFiller     *TextFieldSpecFiller
//     - A pointer to an instance of TextFieldSpecFiller. This
//       method will NOT change the values of internal member
//       variables contained in this instance.
//
//       All data values in this TextFieldSpecFiller instance
//       will be copied to current TextFieldSpecFiller
//       instance ('txtFillerField').
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
func (txtFillerField *TextFieldSpecFiller) CopyIn(
	incomingTxtFieldFiller *TextFieldSpecFiller,
	errorPrefix interface{}) error {

	if txtFillerField.lock == nil {
		txtFillerField.lock = new(sync.Mutex)
	}

	txtFillerField.lock.Lock()

	defer txtFillerField.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldSpecFiller.CopyIn()",
		"")

	if err != nil {
		return err
	}

	return textFieldSpecFillerMolecule{}.ptr().
		copyIn(
			txtFillerField,
			incomingTxtFieldFiller,
			ePrefix)
}

// CopyOut - Returns a deep copy of the current TextFieldSpecFiller
// instance.
//
// If the current TextFieldSpecFiller instance is invalid, an error
// is returned.
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
//  TextFieldSpecFiller
//     - If this method completes successfully and no errors are
//       encountered, this parameter will return a deep copy of the
//       current TextFieldSpecFiller instance.
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
func (txtFillerField *TextFieldSpecFiller) CopyOut(
	errorPrefix interface{}) (
	TextFieldSpecFiller,
	error) {

	if txtFillerField.lock == nil {
		txtFillerField.lock = new(sync.Mutex)
	}

	txtFillerField.lock.Lock()

	defer txtFillerField.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldSpecFiller.CopyOut()",
		"")

	if err != nil {
		return TextFieldSpecFiller{}, err
	}

	return textFieldSpecFillerMolecule{}.ptr().
		copyOut(
			txtFillerField,
			ePrefix)
}

// CopyOutITextField - Returns a deep copy of the current
// TextFieldSpecFiller instance cast as an ITextFieldSpecification
// object.
//
// If the current TextFieldSpecFiller instance is invalid, an error
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
//       the current TextFieldSpecFiller instance cast as an
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
func (txtFillerField *TextFieldSpecFiller) CopyOutITextField(
	errorPrefix interface{}) (
	ITextFieldSpecification,
	error) {

	if txtFillerField.lock == nil {
		txtFillerField.lock = new(sync.Mutex)
	}

	txtFillerField.lock.Lock()

	defer txtFillerField.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	var iTxtFieldSpec ITextFieldSpecification

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldSpecFiller.CopyOutPtr()",
		"")

	if err != nil {
		return iTxtFieldSpec, err
	}

	var newTextFillerField TextFieldSpecFiller

	newTextFillerField,
		err =
		textFieldSpecFillerMolecule{}.ptr().
			copyOut(
				txtFillerField,
				ePrefix)

	if err != nil {
		return iTxtFieldSpec, err
	}

	iTxtFieldSpec = ITextFieldSpecification(&newTextFillerField)

	return iTxtFieldSpec, nil
}

// CopyOutPtr - Returns a pointer to a deep copy of the current
// TextFieldSpecFiller instance.
//
// If the current TextFieldSpecFiller instance is invalid, an error
// is returned.
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
//  *TextFieldSpecFiller
//     - If this method completes successfully and no errors are
//       encountered, this parameter will return a pointer to a
//       deep copy of the current TextFieldSpecFiller instance.
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
func (txtFillerField *TextFieldSpecFiller) CopyOutPtr(
	errorPrefix interface{}) (
	*TextFieldSpecFiller,
	error) {

	if txtFillerField.lock == nil {
		txtFillerField.lock = new(sync.Mutex)
	}

	txtFillerField.lock.Lock()

	defer txtFillerField.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldSpecFiller.CopyOutPtr()",
		"")

	if err != nil {
		return &TextFieldSpecFiller{}, err
	}

	var newTextFillerField TextFieldSpecFiller

	newTextFillerField,
		err =
		textFieldSpecFillerMolecule{}.ptr().
			copyOut(
				txtFillerField,
				ePrefix)

	return &newTextFillerField, err
}

// Empty - Resets all internal member variables to their initial
// or zero states.
//
func (txtFillerField *TextFieldSpecFiller) Empty() {

	if txtFillerField.lock == nil {
		txtFillerField.lock = new(sync.Mutex)
	}

	txtFillerField.lock.Lock()

	textFieldSpecFillerMolecule{}.ptr().
		empty(txtFillerField)

	txtFillerField.lock.Unlock()

	txtFillerField.lock = nil
}

// Equal - Receives a pointer to another instance of
// TextFieldSpecFiller and proceeds to compare the member variables
// to those of the current TextFieldSpecFiller instance in order to
// determine if they are equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables are equal in all respects,
// this flag is set to 'true'. Otherwise, this method returns
// 'false'.
//
func (txtFillerField *TextFieldSpecFiller) Equal(
	incomingTxtFieldFiller *TextFieldSpecFiller) bool {

	if txtFillerField.lock == nil {
		txtFillerField.lock = new(sync.Mutex)
	}

	txtFillerField.lock.Lock()

	defer txtFillerField.lock.Unlock()

	return textFieldSpecFillerMolecule{}.ptr().
		equal(
			txtFillerField,
			incomingTxtFieldFiller)
}

// EqualITextField - Receives an object implementing the
// ITextFieldSpecification interface and proceeds to compare
// the member variables to those of the current TextFieldSpecFiller
// instance in order to determine if they are equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables from both instances are equal
// in all respects, this flag is set to 'true'. Otherwise, this method returns
// 'false'.
//
func (txtFillerField *TextFieldSpecFiller) EqualITextField(
	iTextField ITextFieldSpecification) bool {

	if txtFillerField.lock == nil {
		txtFillerField.lock = new(sync.Mutex)
	}

	txtFillerField.lock.Lock()

	defer txtFillerField.lock.Unlock()

	if iTextField == nil {
		return false
	}

	txtFiller, ok := iTextField.(*TextFieldSpecFiller)

	if !ok {
		return false
	}

	return textFieldSpecFillerMolecule{}.ptr().
		equal(
			txtFillerField,
			txtFiller)
}

// GetFillerChars - Returns the internal member variable
// TextFieldSpecFiller.fillerCharacters ([]rune) as a string.
//
// The filler characters are used to populate the Text Filler
// Field formatted text. The final formatted text is equal
// to:  Filler Characters  X  Filler Characters Repeat Count
//          Example: fillerCharacters = "-*"
//                   fillerRepeatCount = 3
//                   Final Text Filler Field = "-*-*-*"
//
// This method returns a string containing the Filler Characters.
//
func (txtFillerField *TextFieldSpecFiller) GetFillerChars() string {

	if txtFillerField.lock == nil {
		txtFillerField.lock = new(sync.Mutex)
	}

	txtFillerField.lock.Lock()

	defer txtFillerField.lock.Unlock()

	return string(txtFillerField.fillerCharacters)
}

// GetFillerCharsRepeatCount - This method returns the integer
// value of the internal member variable:
//     TextFieldSpecFiller.fillerCharsRepeatCount
//
// The Filler Characters Repeat Count is the number of times that
// the Filler Characters are repeated when constructing the Text
// Filler Field formatted text. The final formatted text is equal
// to:  Filler Characters  X  Filler Characters Repeat Count
//          Example: fillerCharacters = "-*"
//                   fillerRepeatCount = 3
//                   Final Text Filler Field = "-*-*-*"
//
// This method returns the Filler Characters Repeat Count.
//
func (txtFillerField *TextFieldSpecFiller) GetFillerCharsRepeatCount() int {

	if txtFillerField.lock == nil {
		txtFillerField.lock = new(sync.Mutex)
	}

	txtFillerField.lock.Lock()

	defer txtFillerField.lock.Unlock()

	return txtFillerField.fillerCharsRepeatCount
}

// GetFormattedText - Returns the formatted text generated by the
// current instance of TextFieldSpecFiller.
//
// If the length of the Filler Characters array (fillerCharacters
// []rune) is zero, this method returns an empty string.
//
// If the Filler Characters Repeat Count is less than one (+1),
// this method returns an empty string.
//
// If the Filler Characters Repeat Count is greater than
// one-million (+1,000,000), this method returns an empty string.
//
// The length of the final formatted text string is the product of:
//
//  Filler Characters Array Length  X
//               Filler Characters Repeat Count
//
// This method is identical in function to
// TextFieldSpecFiller.String()
//
// This method fulfills the requirements of the
// ITextFieldSpecification interface.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//  Example 1:
//   Filler Characters Array = []rune{'-'}
//   Filler Characters Repeat Count = 3
//   Formatted Text = "---"
//
//  Example 2:
//   Filler Characters Array = []rune{'-','*'}
//   Filler Characters Repeat Count = 3
//   Formatted Text = "-*-*-*"
//
func (txtFillerField *TextFieldSpecFiller) GetFormattedText() string {

	if txtFillerField.lock == nil {
		txtFillerField.lock = new(sync.Mutex)
	}

	txtFillerField.lock.Lock()

	defer txtFillerField.lock.Unlock()

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TextFieldSpecFiller.GetFormattedText()",
		"")

	result,
		err := textFieldSpecFillerMolecule{}.ptr().
		getFormattedText(
			txtFillerField,
			&ePrefix)

	if err != nil {
		result = fmt.Sprintf("%v\n",
			err.Error())
	}

	return result
}

// IsValidInstance - Performs a diagnostic review of the data
// values encapsulated in the current TextFieldSpecFiller instance
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
//       of TextFieldSpecFiller are valid, this returned boolean
//       value is set to 'true'. If any data values are invalid,
//       this return parameter is set to 'false'.
//
func (txtFillerField *TextFieldSpecFiller) IsValidInstance() (
	isValid bool) {

	if txtFillerField.lock == nil {
		txtFillerField.lock = new(sync.Mutex)
	}

	txtFillerField.lock.Lock()

	defer txtFillerField.lock.Unlock()

	isValid,
		_ = textFieldSpecFillerAtom{}.ptr().
		isValidTextFieldSpecFiller(
			txtFillerField,
			nil)

	return isValid
}

// IsValidInstanceError - Performs a diagnostic review of the data
// values encapsulated in the current TextFieldSpecFiller instance
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
//       the current instance of TextFieldSpecFiller are found to be
//       invalid, this method will return an error  along with an
//       appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' (error prefix) will be inserted or
//       prefixed at the beginning of the error message.
//
func (txtFillerField *TextFieldSpecFiller) IsValidInstanceError(
	errorPrefix interface{}) error {

	if txtFillerField.lock == nil {
		txtFillerField.lock = new(sync.Mutex)
	}

	txtFillerField.lock.Lock()

	defer txtFillerField.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldSpecFiller.IsValidInstanceError()",
		"")

	if err != nil {
		return err
	}

	_,
		err = textFieldSpecFillerAtom{}.ptr().
		isValidTextFieldSpecFiller(
			txtFillerField,
			ePrefix)

	return err
}

// New - Returns a new concrete instance of TextFieldSpecFiller.
// This returned instance is empty and unpopulated. All the member
// variables contained in this new instance are set to their
// uninitialized or zero values.
//
// Be advised that setting member variables to their zero values
// means that the returned TextFieldSpecFiller instance is invalid.
// Therefore, in order to use this TextFieldSpecFiller instance,
// users must later call the setter methods on this type in order
// to configure valid and meaningful meaningful member variable
// data values.
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
//  TextFieldSpecFiller
//     - This parameter returns a new and empty concrete instance
//       of TextFieldSpecFiller. Member variable data values are
//       set to their initial or zero values.
//
func (txtFillerField TextFieldSpecFiller) New() TextFieldSpecFiller {

	if txtFillerField.lock == nil {
		txtFillerField.lock = new(sync.Mutex)
	}

	txtFillerField.lock.Lock()

	defer txtFillerField.lock.Unlock()

	newFillerField := textFieldSpecFillerMolecule{}.ptr().
		newEmpty()

	return newFillerField
}

// NewPtr - Returns a pointer to a new unpopulated instance of
// TextFieldSpecFiller. All the member variables contained in
// this new instance are set to their uninitialized or zero values.
//
// Be advised that setting member variables to their zero values
// means that the returned TextFieldSpecFiller instance is invalid.
// Therefore, in order to use this TextFieldSpecFiller instance,
// users must later call the setter methods on this type in order
// to configure valid and meaningful meaningful member variable
// data values.
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
//  *TextFieldSpecFiller
//     - This parameter returns a pointer to a new, empty instance
//       of TextFieldSpecFiller. Member variable data values are
//       set to their initial or zero values.
//
func (txtFillerField TextFieldSpecFiller) NewPtr() *TextFieldSpecFiller {

	if txtFillerField.lock == nil {
		txtFillerField.lock = new(sync.Mutex)
	}

	txtFillerField.lock.Lock()

	defer txtFillerField.lock.Unlock()

	newFillerField := textFieldSpecFillerMolecule{}.ptr().
		newEmpty()

	return &newFillerField
}

// NewPtrTextFiller - Creates and returns a pointer to a new, fully
// populated instance of TextFieldSpecFiller.
//
// This method is identical to TextFieldSpecFiller.NewTextFiller()
// with the sole exception being that this method returns a pointer
// to an instance of TextFieldSpecFiller and
// TextFieldSpecFiller.NewTextFiller() returns a concrete instance
// of TextFieldSpecLabel.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  fillerCharacters           string
//     - A string containing the text characters which will be
//       included in the Text Filler Field. The final Text Filler
//       Field will be constructed from the filler characters
//       repeated one or more times as specified by the
//       'fillerCharsRepeatCount' parameter.
//
//       The Text Field Filler final formatted text is equal to:
//          fillerCharacter X fillerCharsRepeatCount
//          Example: fillerCharacters = "-*"
//                   fillerRepeatCount = 3
//                   Final Text Filler Field = "-*-*-*"
//
//       If 'fillerCharacters' is submitted as an empty or zero
//       length string, this method will return an error.
//
//
//  fillerCharsRepeatCount     int
//     - Controls the number of times 'fillerCharacters' is
//       repeated when constructing the final Text Filler Field
//       returned by this method. The actual length of the string
//       which will populated the completed Text Filler Field is
//       equal to the length of 'fillerCharacters' times the value
//       of 'fillerCharsRepeatCount'.
//
//        Text Field Filler Length =
//          Length of fillerCharacters X fillerCharsRepeatCount
//
//          Example #1: fillerCharacters = "-*"
//                      fillerRepeatCount = 3
//                      Final Text Filler Field = "-*-*-*"
//
//          Example #2: fillerCharacters = "-"
//                      fillerRepeatCount = 3
//                      Final Text Filler Field = "---"
//
//       If 'fillerCharsRepeatCount' has a value less than one (1) or
//       greater than one-million (1,000,000), an error will be
//       returned.
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
//  *TextFieldSpecFiller
//     - If this method completes successfully, this parameter will
//       return a pointer to a valid and fully populated Text Filler
//       Field object.
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
//   Filler Characters Array = "-"
//   Filler Characters Repeat Count = 3
//   Formatted Text = "---"
//
//  Example 2:
//   Filler Characters Array = "-*"
//   Filler Characters Repeat Count = 3
//   Formatted Text = "-*-*-*"
//
func (txtFillerField TextFieldSpecFiller) NewPtrTextFiller(
	fillerCharacters string,
	fillerCharsRepeatCount int,
	errorPrefix interface{}) (
	*TextFieldSpecFiller,
	error) {

	if txtFillerField.lock == nil {
		txtFillerField.lock = new(sync.Mutex)
	}

	txtFillerField.lock.Lock()

	defer txtFillerField.lock.Unlock()

	newTxtFillerField := TextFieldSpecFiller{}

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldSpecFiller.NewPtrTextFiller()",
		"")

	if err != nil {
		return &newTxtFillerField, err
	}

	fillerCharsRunes := []rune(fillerCharacters)

	err = textFieldSpecFillerNanobot{}.ptr().
		setTxtFieldSpecFiller(
			&newTxtFillerField,
			fillerCharsRunes,
			fillerCharsRepeatCount,
			ePrefix)

	return &newTxtFillerField, err
}

// NewPtrTextFillerRune - Creates and returns a pointer to a new,
// fully populated instance of TextFieldSpecFiller.
//
// This method is identical to
// TextFieldSpecFiller.NewTextFillerRune() with the sole exception
// being that this method returns a pointer to an instance of
// TextFieldSpecFiller and TextFieldSpecFiller.NewTextFillerRune()
// returns a concrete instance of TextFieldSpecFiller.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  fillerCharacter            rune
//     - A rune containing the text character which will be
//       included in the Text Filler Field. The final Text Filler
//       Field will be constructed from ths filler character
//       repeated one or more times as specified by the
//       'fillerCharsRepeatCount' parameter.
//
//       The Text Field Filler final formatted text is equal to:
//           fillerCharacter X fillerCharsRepeatCount
//           Example: fillerCharacter = '-'
//                    fillerRepeatCount = 3
//                    Final Text Filler Field = "---"
//
//       If 'fillerCharacter' is submitted with a zero value,
//       this method will return an error.
//
//
//  fillerCharsRepeatCount     int
//     - Controls the number of times 'fillerCharacter' is
//       repeated when constructing the final Text Filler Field
//       returned by this method. The actual length of the string
//       which will populated the completed Text Filler Field is
//       equal to the length of 'fillerCharacter' (1) times the
//       value of 'fillerCharsRepeatCount'.
//
//         Text Field Filler Length =
//           Length of fillerCharacter (1) X fillerCharsRepeatCount
//           Example: fillerCharacter = '-'
//                    fillerRepeatCount = 3
//                    Final Text Filler Field = "---"
//
//       If 'fillerCharsRepeatCount' has a value less than one (1) or
//       greater than one-million (1,000,000), an error will be
//       returned.
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
//  *TextFieldSpecFiller
//     - If this method completes successfully, this parameter will
//       return a pointer to a new, valid and fully populated Text
//       Filler Field.
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
//   Filler Characters Array = '-'
//   Filler Characters Repeat Count = 3
//   Formatted Text = "---"
//
//  Example 2:
//   Filler Characters Array = '*'
//   Filler Characters Repeat Count = 3
//   Formatted Text = "***"
//
func (txtFillerField TextFieldSpecFiller) NewPtrTextFillerRune(
	fillerCharacter rune,
	fillerCharsRepeatCount int,
	errorPrefix interface{}) (
	*TextFieldSpecFiller,
	error) {

	if txtFillerField.lock == nil {
		txtFillerField.lock = new(sync.Mutex)
	}

	txtFillerField.lock.Lock()

	defer txtFillerField.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newTxtFillerField := TextFieldSpecFiller{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldSpecFiller.NewPtrTextFillerRune()",
		"")

	if err != nil {
		return &newTxtFillerField, err
	}

	fillerCharsRunes := []rune{fillerCharacter}

	err = textFieldSpecFillerNanobot{}.ptr().
		setTxtFieldSpecFiller(
			&newTxtFillerField,
			fillerCharsRunes,
			fillerCharsRepeatCount,
			ePrefix)

	return &newTxtFillerField, err
}

// NewPtrTextFillerRuneArray - Creates and returns a pointer to a
// new, fully populated instance of TextFieldSpecFiller.
//
// This method is identical to
// TextFieldSpecFiller.NewTextFillerRuneArray() with the sole
// exception being that this method returns a pointer to an
// instance of TextFieldSpecFiller and
// TextFieldSpecFiller.NewTextFillerRuneArray() returns a concrete
// instance of TextFieldSpecFiller.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  fillerCharacters           []rune
//     - A rune array containing the text characters which will be
//       included in the Text Filler Field. The final Text Filler
//       Field will be constructed from ths filler characters
//       repeated one or more times as specified by the
//       'fillerCharsRepeatCount' parameter.
//
//       The Text Field Filler final formatted text is equal to:
//          fillerCharacters X fillerCharsRepeatCount
//          Example: fillerCharacters = []rune{'-','*'}
//                   fillerRepeatCount = 3
//                   Final Text Filler Field = "-*-*-*"
//
//       If 'fillerCharacters' is submitted with a zero length rune
//       array, this method will return an error.
//
//
//  fillerCharsRepeatCount     int
//     - Controls the number of times 'fillerCharacters' is
//       repeated when constructing the final Text Filler Field
//       returned by this method. The actual length of the string
//       which will populated the completed Text Filler Field is
//       equal to the length of 'fillerCharacters' times the value
//       of 'fillerCharsRepeatCount'.
//
//        Text Field Filler Length =
//          Length of fillerCharacters X fillerCharsRepeatCount
//
//          Example #1: fillerCharacters = []rune{'-','*'}
//                      fillerRepeatCount = 3
//                      Final Text Filler Field = "-*-*-*"
//
//          Example #2: fillerCharacters = []rune{'-'}
//                      fillerRepeatCount = 3
//                      Final Text Filler Field = "---"
//
//       If 'fillerCharsRepeatCount' has a value less than one (1) or
//       greater than one-million (1,000,000), an error will be
//       returned.
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
//  *TextFieldSpecFiller
//     - If this method completes successfully, this parameter will
//       return a pointer to a new, valid and fully populated Text
//       Filler Field.
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
//   Filler Characters Array = []rune{'-'}
//   Filler Characters Repeat Count = 3
//   Formatted Text = "---"
//
//  Example 2:
//   Filler Characters Array = []rune{'-','*'}
//   Filler Characters Repeat Count = 3
//   Formatted Text = "-*-*-*"
//
func (txtFillerField TextFieldSpecFiller) NewPtrTextFillerRuneArray(
	fillerCharacters []rune,
	fillerCharsRepeatCount int,
	errorPrefix interface{}) (
	*TextFieldSpecFiller,
	error) {

	if txtFillerField.lock == nil {
		txtFillerField.lock = new(sync.Mutex)
	}

	txtFillerField.lock.Lock()

	defer txtFillerField.lock.Unlock()

	newTxtFillerField := TextFieldSpecFiller{}

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldSpecFiller."+
			"NewPtrTextFillerRuneArray()",
		"")

	if err != nil {
		return &newTxtFillerField, err
	}

	err = textFieldSpecFillerNanobot{}.ptr().
		setTxtFieldSpecFiller(
			&newTxtFillerField,
			fillerCharacters,
			fillerCharsRepeatCount,
			ePrefix)

	return &newTxtFillerField, err
}

// NewTextFiller - Creates and returns a new, concrete instance of
// TextFieldSpecFiller.
//
// This method is identical to
// TextFieldSpecFiller.NewPtrTextFiller() with the sole exception
// being that this method returns a concrete instance of
// TextFieldSpecFiller and TextFieldSpecFiller.NewPtrTextFiller()
// returns a pointer to a TextFieldSpecLabel instance.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  fillerCharacters           string
//     - A string containing the text characters which will be
//       included in the Text Filler Field. The final Text Filler
//       Field will be constructed from the filler characters
//       repeated one or more times as specified by the
//       'fillerCharsRepeatCount' parameter.
//
//        Text Field Filler Length =
//          Length of fillerCharacters X fillerCharsRepeatCount
//
//          Example #1: fillerCharacters = "-*"
//                      fillerRepeatCount = 3
//                      Final Text Filler Field = "-*-*-*"
//
//          Example #2: fillerCharacters = "-"
//                      fillerRepeatCount = 3
//                      Final Text Filler Field = "---"
//
//       If 'fillerCharacters' is submitted as an empty or zero
//       length string, this method will return an error.
//
//
//  fillerCharsRepeatCount     int
//     - Controls the number of times 'fillerCharacters' is
//       repeated when constructing the final Text Filler Field
//       returned by this method. The actual length of the string
//       which will populated the completed Text Filler Field is
//       equal to the length of 'fillerCharacters' times the value
//       of 'fillerCharsRepeatCount'.
//
//        Text Field Filler Length =
//          Length of fillerCharacters X fillerCharsRepeatCount
//
//          Example #1: fillerCharacters = "-*"
//                      fillerRepeatCount = 3
//                      Final Text Filler Field = "-*-*-*"
//
//          Example #2: fillerCharacters = "-"
//                      fillerRepeatCount = 3
//                      Final Text Filler Field = "---"
//
//       If 'fillerCharsRepeatCount' has a value less than one (1) or
//       greater than one-million (1,000,000), an error will be
//       returned.
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
//  TextFieldSpecFiller
//     - If this method completes successfully, this parameter will
//       return a new concrete Text Filler Field object.
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
//   Filler Characters Array = "-"
//   Filler Characters Repeat Count = 3
//   Formatted Text = "---"
//
//  Example 2:
//   Filler Characters Array = "-*"
//   Filler Characters Repeat Count = 3
//   Formatted Text = "-*-*-*"
//
func (txtFillerField TextFieldSpecFiller) NewTextFiller(
	fillerCharacters string,
	fillerCharsRepeatCount int,
	errorPrefix interface{}) (
	TextFieldSpecFiller,
	error) {

	if txtFillerField.lock == nil {
		txtFillerField.lock = new(sync.Mutex)
	}

	txtFillerField.lock.Lock()

	defer txtFillerField.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newTxtFillerField := TextFieldSpecFiller{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldSpecFiller.NewTextFiller()",
		"")

	if err != nil {
		return newTxtFillerField, err
	}

	fillerCharsRunes := []rune(fillerCharacters)

	err = textFieldSpecFillerNanobot{}.ptr().
		setTxtFieldSpecFiller(
			&newTxtFillerField,
			fillerCharsRunes,
			fillerCharsRepeatCount,
			ePrefix)

	return newTxtFillerField, err
}

// NewTextFillerRune - Creates and returns a new concrete instance
// of TextFieldSpecFiller.
//
// This method is identical to
// TextFieldSpecFiller.NewPtrTextFillerRune() with the sole
// exception being that this method returns a concrete instance of
// TextFieldSpecFiller and TextFieldSpecFiller.NewPtrTextFillerRune()
// returns a pointer to a TextFieldSpecFiller instance.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  fillerCharacter            rune
//     - A rune containing the text character which will be
//       included in the Text Filler Field. The final Text Filler
//       Field will be constructed from ths filler character
//       repeated one or more times as specified by the
//       'fillerCharsRepeatCount' parameter.
//
//       The Text Field Filler final formatted text is equal to:
//           fillerCharacter X fillerCharsRepeatCount
//           Example: fillerCharacter = '-'
//                    fillerRepeatCount = 3
//                    Final Text Filler Field = "---"
//
//       If 'fillerCharacter' is submitted with a zero value,
//       this method will return an error.
//
//
//  fillerCharsRepeatCount     int
//     - Controls the number of times 'fillerCharacter' is
//       repeated when constructing the final Text Filler Field
//       returned by this method. The actual length of the string
//       which will populated the completed Text Filler Field is
//       equal to the length of 'fillerCharacter' (1) times the
//       value of 'fillerCharsRepeatCount'.
//
//         Text Field Filler Length =
//           Length of fillerCharacter (1) X fillerCharsRepeatCount
//           Example: fillerCharacter = '-'
//                    fillerRepeatCount = 3
//                    Final Text Filler Field = "---"
//
//       If 'fillerCharsRepeatCount' has a value less than one (1) or
//       greater than one-million (1,000,000), an error will be
//       returned.
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
//  TextFieldSpecFiller
//     - If this method completes successfully, this parameter will
//       return a new, concrete instance of TextFieldSpecFiller.
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
//   Filler Characters Array = '-'
//   Filler Characters Repeat Count = 3
//   Formatted Text = "---"
//
//  Example 2:
//   Filler Characters Array = '*'
//   Filler Characters Repeat Count = 3
//   Formatted Text = "***"
//
func (txtFillerField TextFieldSpecFiller) NewTextFillerRune(
	fillerCharacter rune,
	fillerCharsRepeatCount int,
	errorPrefix interface{}) (
	TextFieldSpecFiller,
	error) {

	if txtFillerField.lock == nil {
		txtFillerField.lock = new(sync.Mutex)
	}

	txtFillerField.lock.Lock()

	defer txtFillerField.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newTxtFillerField := TextFieldSpecFiller{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldSpecFiller.NewTextFillerRune()",
		"")

	if err != nil {
		return newTxtFillerField, err
	}

	fillerCharsRunes := []rune{fillerCharacter}

	err = textFieldSpecFillerNanobot{}.ptr().
		setTxtFieldSpecFiller(
			&newTxtFillerField,
			fillerCharsRunes,
			fillerCharsRepeatCount,
			ePrefix)

	return newTxtFillerField, err
}

// NewTextFillerRuneArray - Creates and returns a new concrete
// instance of TextFieldSpecFiller.
//
// This method is identical to
// TextFieldSpecFiller.NewPtrTextFillerRuneArray() with the sole
// exception being that this method returns a concrete instance of
// TextFieldSpecFiller while
// TextFieldSpecFiller.NewPtrTextFillerRuneArray() returns a pointer
// to an instance of TextFieldSpecFiller.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  fillerCharacters           []rune
//     - A rune array containing the text characters which will be
//       included in the Text Filler Field. The final Text Filler
//       Field will be constructed from ths filler characters
//       repeated one or more times as specified by the
//       'fillerCharsRepeatCount' parameter.
//
//       The Text Field Filler final formatted text is equal to:
//          fillerCharacters X fillerCharsRepeatCount
//          Example: fillerCharacters = []rune{'-','*'}
//                   fillerRepeatCount = 3
//                   Final Text Filler Field = "-*-*-*"
//
//       If 'fillerCharacters' is submitted with a zero length rune
//       array, this method will return an error.
//
//
//  fillerCharsRepeatCount     int
//     - Controls the number of times 'fillerCharacters' is
//       repeated when constructing the final Text Filler Field
//       returned by this method. The actual length of the string
//       which will populated the completed Text Filler Field is
//       equal to the length of 'fillerCharacters' times the value
//       of 'fillerCharsRepeatCount'.
//
//        Text Field Filler Length =
//          Length of fillerCharacters X fillerCharsRepeatCount
//
//          Example #1: fillerCharacters = []rune{'-','*'}
//                      fillerRepeatCount = 3
//                      Final Text Filler Field = "-*-*-*"
//
//          Example #2: fillerCharacters = []rune{'-'}
//                      fillerRepeatCount = 3
//                      Final Text Filler Field = "---"
//
//       If 'fillerCharsRepeatCount' has a value less than one (1) or
//       greater than one-million (1,000,000), an error will be
//       returned.
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
//  TextFieldSpecFiller
//     - If this method completes successfully, this parameter will
//       return a new concrete instance of TextFieldSpecFiller.
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
//   Filler Characters Array = []rune{'-'}
//   Filler Characters Repeat Count = 3
//   Formatted Text = "---"
//
//  Example 2:
//   Filler Characters Array = []rune{'-','*'}
//   Filler Characters Repeat Count = 3
//   Formatted Text = "-*-*-*"
//
func (txtFillerField TextFieldSpecFiller) NewTextFillerRuneArray(
	fillerCharacters []rune,
	fillerCharsRepeatCount int,
	errorPrefix interface{}) (
	TextFieldSpecFiller,
	error) {

	if txtFillerField.lock == nil {
		txtFillerField.lock = new(sync.Mutex)
	}

	txtFillerField.lock.Lock()

	defer txtFillerField.lock.Unlock()

	newTxtFillerField := TextFieldSpecFiller{}

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldSpecFiller."+
			"NewTextFillerRuneArray()",
		"")

	if err != nil {
		return newTxtFillerField, err
	}

	err = textFieldSpecFillerNanobot{}.ptr().
		setTxtFieldSpecFiller(
			&newTxtFillerField,
			fillerCharacters,
			fillerCharsRepeatCount,
			ePrefix)

	return newTxtFillerField, err
}

// SetFillerCharsRepeatCount - Sets the Filler Characters Repeat
// Count for the current instance of TextFieldSpecFiller.
//
// Text Field Specifications like TextFieldSpecFiller are used as
// building blocks for constructing single lines of text which can
// be configured for text display, file output or printing.
//
// Typically, filler fields (TextFieldSpecFiller) are used as
// margins containing multiple white space characters, or line
// separators containing multiple dashes, equal signs or underscore
// characters. Filler fields consist of filler characters
// ('fillerCharacters') and the filler characters repeat count
// ('fillerCharsRepeatCount'). A filler field is made up of one
// or more filler characters. These filler characters are
// repeated one or more times in order to construct the complete
// filler field as shown in the following examples:
//
//  Example 1:
//   Filler Characters = "-"
//   Filler Characters Repeat Count = 3
//   Formatted Text = "---"
//
//  Example 2:
//   Filler Characters = "-*"
//   Filler Characters Repeat Count = 3
//   Formatted Text = "-*-*-*"
//
// The 'fillerCharsRepeatCount' integer value is the number times
// that 'fillerCharacters' is repeated in order to construct the
// Filler Text Field. Be advised that Filler Text Fields requires a
// 'fillerCharsRepeatCount' value greater than zero.
// 'fillerCharsRepeatCount' values less than or equal to zero
// will trigger an error.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  fillerCharsRepeatCount     int
//     - Controls the number of times 'fillerCharacters' is
//       repeated when constructing the final Text Filler Field
//       returned by this method. The actual length of the string
//       which will populated the completed Text Filler Field is
//       equal to the length of 'fillerCharacters' times the value
//       of 'fillerCharsRepeatCount'.
//
//        Text Field Filler Length =
//          Length of fillerCharacters X fillerCharsRepeatCount
//
//          Example #1: fillerCharacters = "-*"
//                      fillerRepeatCount = 3
//                      Final Text Filler Field = "-*-*-*"
//
//          Example #2: fillerCharacters = "-"
//                      fillerRepeatCount = 3
//                      Final Text Filler Field = "---"
//
//       If 'fillerCharsRepeatCount' has a value less than one (1) or
//       greater than one-million (1,000,000), an error will be
//       returned.
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
func (txtFillerField *TextFieldSpecFiller) SetFillerCharsRepeatCount(
	fillerCharsRepeatCount int,
	errorPrefix interface{}) error {

	if txtFillerField.lock == nil {
		txtFillerField.lock = new(sync.Mutex)
	}

	txtFillerField.lock.Lock()

	defer txtFillerField.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldSpecFiller.SetFillerCharsRepeatCount()",
		"")

	if err != nil {
		return err
	}

	err = textFieldSpecFillerElectron{}.ptr().
		isFillerCharsRepeatCountValid(
			fillerCharsRepeatCount,
			ePrefix.XCtx(
				"Input parameter 'fillerCharsRepeatCount' invalid!"))

	if err != nil {
		return err
	}

	txtFillerField.fillerCharsRepeatCount =
		fillerCharsRepeatCount

	return err
}

// SetTextFiller - Overwrites the internal member variables for the
// current instance of TextFieldSpecFiller. The new data values are
// generated from the input parameters.
//
// IMPORTANT
//
// This method will overwrite and delete the existing data values
// for the current TextFieldSpecFiller instance (txtFillerField).
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  fillerCharacters           string
//     - A string containing the text characters which will be
//       included in the Text Filler Field. The final Text Filler
//       Field will be constructed from the filler characters
//       repeated one or more times as specified by the
//       'fillerCharsRepeatCount' parameter.
//
//       The Text Field Filler final formatted text is equal to:
//          fillerCharacter X fillerCharsRepeatCount
//          Example: fillerCharacters = "-*"
//                   fillerRepeatCount = 3
//                   Final Text Filler Field = "-*-*-*"
//
//       If 'fillerCharacters' is submitted as an empty or zero
//       length string, this method will return an error.
//
//
//  fillerCharsRepeatCount     int
//     - Controls the number of times 'fillerCharacters' is
//       repeated when constructing the final Text Filler Field
//       returned by this method. The actual length of the string
//       which will populated the completed Text Filler Field is
//       equal to the length of 'fillerCharacters' times the value
//       of 'fillerCharsRepeatCount'.
//
//        Text Field Filler Length =
//          Length of fillerCharacters X fillerCharsRepeatCount
//
//          Example #1: fillerCharacters = "-*"
//                      fillerRepeatCount = 3
//                      Final Text Filler Field = "-*-*-*"
//
//          Example #2: fillerCharacters = "-"
//                      fillerRepeatCount = 3
//                      Final Text Filler Field = "---"
//
//       If 'fillerCharsRepeatCount' has a value less than one (1) or
//       greater than one-million (1,000,000), an error will be
//       returned.
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
//   Filler Characters Array = "-"
//   Filler Characters Repeat Count = 3
//   Formatted Text = "---"
//
//  Example 2:
//   Filler Characters Array = "-*"
//   Filler Characters Repeat Count = 3
//   Formatted Text = "-*-*-*"
//
func (txtFillerField *TextFieldSpecFiller) SetTextFiller(
	fillerCharacters string,
	fillerCharsRepeatCount int,
	errorPrefix interface{}) error {

	if txtFillerField.lock == nil {
		txtFillerField.lock = new(sync.Mutex)
	}

	txtFillerField.lock.Lock()

	defer txtFillerField.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldSpecFiller.SetTextFiller()",
		"")

	if err != nil {
		return err
	}

	fillerCharsRunes := []rune(fillerCharacters)

	err = textFieldSpecFillerNanobot{}.ptr().
		setTxtFieldSpecFiller(
			txtFillerField,
			fillerCharsRunes,
			fillerCharsRepeatCount,
			ePrefix)

	return err
}

// SetTextFillerRune - Overwrites the internal member variables for
// the current instance of TextFieldSpecFiller. The new data values
// are generated from the input parameters.
//
// This method differs from TextFieldSpecFiller.SetTextFiller() in
// that this method accepts a single rune as the Filler Character.
//
//
// IMPORTANT
//
// This method will overwrite and delete the existing data values
// for the current TextFieldSpecFiller instance (txtFillerField).
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  fillerCharacter            rune
//     - A rune containing the text character which will be
//       included in the Text Filler Field. The final Text Filler
//       Field will be constructed from ths filler character
//       repeated one or more times as specified by the
//       'fillerCharsRepeatCount' parameter.
//
//       The Text Field Filler final formatted text is equal to:
//           fillerCharacter X fillerCharsRepeatCount
//           Example: fillerCharacter = '-'
//                    fillerRepeatCount = 3
//                    Final Text Filler Field = "---"
//
//       If 'fillerCharacter' is submitted with a zero value,
//       this method will return an error.
//
//
//  fillerCharsRepeatCount     int
//     - Controls the number of times 'fillerCharacter' is
//       repeated when constructing the final Text Filler Field
//       returned by this method. The actual length of the string
//       which will populated the completed Text Filler Field is
//       equal to the length of 'fillerCharacter' (1) times the
//       value of 'fillerCharsRepeatCount'.
//
//         Text Field Filler Length =
//           Length of fillerCharacter (1) X fillerCharsRepeatCount
//           Example: fillerCharacter = '-'
//                    fillerRepeatCount = 3
//                    Final Text Filler Field = "---"
//
//       If 'fillerCharsRepeatCount' has a value less than one (1) or
//       greater than one-million (1,000,000), an error will be
//       returned.
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
//   Filler Characters Array = '-'
//   Filler Characters Repeat Count = 3
//   Formatted Text = "---"
//
//  Example 2:
//   Filler Characters Array = '*'
//   Filler Characters Repeat Count = 3
//   Formatted Text = "***"
//
func (txtFillerField *TextFieldSpecFiller) SetTextFillerRune(
	fillerCharacter rune,
	fillerCharsRepeatCount int,
	errorPrefix interface{}) error {

	if txtFillerField.lock == nil {
		txtFillerField.lock = new(sync.Mutex)
	}

	txtFillerField.lock.Lock()

	defer txtFillerField.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldSpecFiller.SetTextFillerRune()",
		"")

	if err != nil {
		return err
	}

	fillerCharsRunes := []rune{fillerCharacter}

	err = textFieldSpecFillerNanobot{}.ptr().
		setTxtFieldSpecFiller(
			txtFillerField,
			fillerCharsRunes,
			fillerCharsRepeatCount,
			ePrefix)

	return err
}

// SetTextFillerRuneArray - Overwrites the internal member
// variables for the current instance of TextFieldSpecFiller. The
// new data values are generated from the input parameters.
//
// This method differs from TextFieldSpecFiller.SetTextFiller() in
// that this method accepts an array of runes as the Filler
// Characters.
//
//
// IMPORTANT
//
// This method will overwrite and delete the existing data values
// for the current TextFieldSpecFiller instance (txtFillerField).
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  fillerCharacters           []rune
//     - A rune array containing the text characters which will be
//       included in the Text Filler Field. The final Text Filler
//       Field will be constructed from ths filler characters
//       repeated one or more times as specified by the
//       'fillerCharsRepeatCount' parameter.
//
//       The Text Field Filler final formatted text is equal to:
//          fillerCharacters X fillerCharsRepeatCount
//          Example: fillerCharacters = []rune{'-','*'}
//                   fillerRepeatCount = 3
//                   Final Text Filler Field = "-*-*-*"
//
//       If 'fillerCharacters' is submitted with a zero length rune
//       array, this method will return an error.
//
//
//  fillerCharsRepeatCount     int
//     - Controls the number of times 'fillerCharacters' is
//       repeated when constructing the final Text Filler Field
//       returned by this method. The actual length of the string
//       which will populated the completed Text Filler Field is
//       equal to the length of 'fillerCharacters' times the value
//       of 'fillerCharsRepeatCount'.
//
//        Text Field Filler Length =
//          Length of fillerCharacters X fillerCharsRepeatCount
//
//          Example #1: fillerCharacters = []rune{'-','*'}
//                      fillerRepeatCount = 3
//                      Final Text Filler Field = "-*-*-*"
//
//          Example #2: fillerCharacters = []rune{'-'}
//                      fillerRepeatCount = 3
//                      Final Text Filler Field = "---"
//
//       If 'fillerCharsRepeatCount' has a value less than one (1) or
//       greater than one-million (1,000,000), an error will be
//       returned.
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
//   Filler Characters Array = []rune{'-'}
//   Filler Characters Repeat Count = 3
//   Formatted Text = "---"
//
//  Example 2:
//   Filler Characters Array = []rune{'-','*'}
//   Filler Characters Repeat Count = 3
//   Formatted Text = "-*-*-*"
//
func (txtFillerField *TextFieldSpecFiller) SetTextFillerRuneArray(
	fillerCharacters []rune,
	fillerCharsRepeatCount int,
	errorPrefix interface{}) error {

	if txtFillerField.lock == nil {
		txtFillerField.lock = new(sync.Mutex)
	}

	txtFillerField.lock.Lock()

	defer txtFillerField.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldSpecFiller.SetTextFillerRuneArray()",
		"")

	if err != nil {
		return err
	}

	err = textFieldSpecFillerNanobot{}.ptr().
		setTxtFieldSpecFiller(
			txtFillerField,
			fillerCharacters,
			fillerCharsRepeatCount,
			ePrefix)

	return err
}

// String - Returns the formatted text generated by the
// current instance of TextFieldSpecFiller.
//
// If the length of the Filler Characters array (fillerCharacters
// []rune) is zero, this method returns an empty string.
//
// If the Filler Characters Repeat Count is less than one (+1),
// this method returns an empty string.
//
// If the Filler Characters Repeat Count is greater than
// one-million (+1,000,000), this method returns an empty string.
//
// The length of the final formatted text string is the product of:
//
//  Filler Characters Array Length  X
//               Filler Characters Repeat Count
//
// This method is identical in function to
// TextFieldSpecFiller.GetFormattedText()
//
// This method fulfills the requirements of the
// ITextFieldSpecification interface.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//  Example 1:
//   Filler Characters Array = []rune{'-'}
//   Filler Characters Repeat Count = 3
//   Formatted Text = "---"
//
//  Example 2:
//   Filler Characters Array = []rune{'-','*'}
//   Filler Characters Repeat Count = 3
//   Formatted Text = "-*-*-*"
//
func (txtFillerField TextFieldSpecFiller) String() string {

	if txtFillerField.lock == nil {
		txtFillerField.lock = new(sync.Mutex)
	}

	txtFillerField.lock.Lock()

	defer txtFillerField.lock.Unlock()

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TextFieldSpecFiller.GetFormattedText()",
		"")

	result,
		err := textFieldSpecFillerMolecule{}.ptr().
		getFormattedText(
			&txtFillerField,
			&ePrefix)

	if err != nil {
		result = fmt.Sprintf("%v\n",
			err.Error())
	}

	return result
}

// TextFieldName - returns a string specifying the name
// of the Text Field specification. This method fulfills
// the requirements of the ITextFieldSpecification interface.
//
func (txtFillerField TextFieldSpecFiller) TextFieldName() string {

	return "TextFieldSpecFiller"
}

// TextTypeName - returns a string specifying the type
// of Text Field specification. This method fulfills
// the requirements of the ITextSpecification interface.
//
func (txtFillerField TextFieldSpecFiller) TextTypeName() string {

	return "TextFieldSpecFiller"
}
