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
// Typically, filler fields are used as margins containing multiple
// space characters, or line separators containing multiple dashes,
// equal signs or underscore characters.
//
// Which constructing a text line using type
// TextLineSpecStandardLine, it is common to include multiple Text
// Filler fields as required to separate labels or number strings.
//
// The 'fillerCharsRepeatCount' integer value is the number times
// that 'fillerCharacters' is repeated in order to construct the
// Filler Text Field. Be advised that Filler Text Fields requires a
// 'fillerCharsRepeatCount' value greater than zero.
// 'fillerCharsRepeatCount' values less than or equal to zero
// constitute an error condition.
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
// All of the data fields in current TextFieldSpecFiller instance
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
// If the length of the Filler Characters array
// (fillerCharacters []rune) is zero, this method returns an
// empty string.
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
func (txtFillerField *TextFieldSpecFiller) GetFormattedText() string {

	if txtFillerField.lock == nil {
		txtFillerField.lock = new(sync.Mutex)
	}

	txtFillerField.lock.Lock()

	defer txtFillerField.lock.Unlock()

	var isValid bool

	isValid,
		_ = textFieldSpecFillerAtom{}.ptr().
		isValidTextFieldSpecFiller(
			txtFillerField,
			nil)

	if !isValid {
		return ""
	}

	var result string

	for i := 0; i < txtFillerField.fillerCharsRepeatCount; i++ {
		result += string(txtFillerField.fillerCharacters)
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

// NewConstructor - Creates and returns a new, fully populated
// instance of TextFieldSpecFiller.
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
//          Example: fillerCharacters = "-*"
//                   fillerRepeatCount = 3
//                   Final Text Filler Field = "-*-*-*"
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
//       return a valid and fully populated Text Filler Field.
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
func (txtFillerField TextFieldSpecFiller) NewConstructor(
	fillerCharacters string,
	fillerCharsCount int,
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
		"TextFieldSpecFiller.NewConstructor()",
		"")

	if err != nil {
		return TextFieldSpecFiller{}, err
	}

	txtFillerElectron := textFieldSpecFillerElectron{}

	fillerCharsRunes := []rune(fillerCharacters)

	var lenFillerChars int

	lenFillerChars,
		err = txtFillerElectron.isFillerCharsValid(
		fillerCharsRunes,
		ePrefix.XCtx(
			"fillerCharacters"))

	if err != nil {
		return TextFieldSpecFiller{}, err
	}

	err = txtFillerElectron.isFillerCharsRepeatCountValid(
		fillerCharsCount,
		ePrefix.XCtx("fillerCharsCount"))

	if err != nil {
		return TextFieldSpecFiller{}, err
	}

	newTxtFillerField := textFieldSpecFillerMolecule{}.ptr().newEmpty()

	newTxtFillerField.fillerCharacters =
		make([]rune, lenFillerChars)

	copy(newTxtFillerField.fillerCharacters,
		fillerCharsRunes)

	newTxtFillerField.fillerCharsRepeatCount =
		fillerCharsCount

	return newTxtFillerField, nil
}

// NewConstructorRune - Creates and returns a pointer to a new,
// fully populated instance of TextFieldSpecFiller. This method is
// similar to method TextFieldSpecFiller.NewConstructor() with the
// distinction being that this method returns a pointer to a new
// instance of TextFieldSpecFiller and it accepts a single rune
// rune as a filler character.
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
func (txtFillerField TextFieldSpecFiller) NewConstructorRune(
	fillerCharacter rune,
	fillerCharsCount int,
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
		"TextFieldSpecFiller.NewConstructorRune()",
		"")

	if err != nil {
		return &TextFieldSpecFiller{}, err
	}

	txtFieldFillerElectron := textFieldSpecFillerElectron{}

	err = txtFieldFillerElectron.isFillerCharacterValid(
		fillerCharacter,
		ePrefix.XCtx("fillerCharacter"))

	if err != nil {
		return &TextFieldSpecFiller{}, err
	}

	err = txtFieldFillerElectron.isFillerCharsRepeatCountValid(
		fillerCharsCount,
		ePrefix.XCtx("fillerCharsCount"))

	if err != nil {
		return &TextFieldSpecFiller{}, err
	}

	newTxtFillerField := textFieldSpecFillerMolecule{}.ptr().
		newEmpty()

	newTxtFillerField.fillerCharacters =
		make([]rune, 1)

	newTxtFillerField.fillerCharacters[0] =
		fillerCharacter

	newTxtFillerField.fillerCharsRepeatCount =
		fillerCharsCount

	return &newTxtFillerField, nil
}

// NewConstructorRuneArray - Creates and returns a pointer to a
// new, fully populated instance of TextFieldSpecFiller. This
// method is similar to method TextFieldSpecFiller.NewConstructor()
// with the distinction being that this method returns a pointer
// to a new instance of TextFieldSpecFiller and it accepts an array
// runes filler characters.
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
//         Text Field Filler Length =
//           Length of fillerCharacters X fillerCharsRepeatCount
//           Example: fillerCharacters = []rune{'-','*'}
//                    fillerRepeatCount = 3
//                    Final Text Filler Field = "-*-*-*"
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
func (txtFillerField TextFieldSpecFiller) NewConstructorRuneArray(
	fillerCharacters []rune,
	fillerCharsCount int,
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
		"TextFieldSpecFiller.NewConstructorRune()",
		"")

	if err != nil {
		return &TextFieldSpecFiller{}, err
	}

	lenRuneArray := len(fillerCharacters)

	if lenRuneArray == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fillerCharacters' is a zero length rune array!\n",
			ePrefix.String())

		return &TextFieldSpecFiller{}, err
	}

	if fillerCharsCount < 1 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fillerCharsRepeatCount' is less than one (1)!\n"+
			"'fillerCharsRepeatCount' controls the number of repetitions of 'fillerCharacters'\n"+
			"in the Filler Text Field.\n",
			ePrefix.String())
		return &TextFieldSpecFiller{}, err
	}

	if fillerCharsCount > 1000000 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fillerCharsRepeatCount' is greater than one-million (+1,000,000)!\n"+
			"'fillerCharsRepeatCount' controls the number of repetitions of 'fillerCharacters'\n"+
			"in the Filler Text Field.\n",
			ePrefix.String())
		return &TextFieldSpecFiller{}, err
	}

	newTxtFillerField := textFieldSpecFillerMolecule{}.ptr().
		newEmpty()

	newTxtFillerField.fillerCharacters =
		make([]rune, lenRuneArray)

	copy(
		newTxtFillerField.fillerCharacters,
		fillerCharacters)

	newTxtFillerField.fillerCharsRepeatCount =
		fillerCharsCount

	return &newTxtFillerField, nil
}

// NewEmpty - Returns a new unpopulated instance of
// TextFieldSpecFiller. All of the member variables contained in
// this new instance are set to their uninitialized or zero values.
//
// Be advised that setting member variables to their zero values
// means that the TextFieldSpecFiller is invalid.
//
func (txtFillerField TextFieldSpecFiller) NewEmpty() TextFieldSpecFiller {

	if txtFillerField.lock == nil {
		txtFillerField.lock = new(sync.Mutex)
	}

	txtFillerField.lock.Lock()

	defer txtFillerField.lock.Unlock()

	return textFieldSpecFillerMolecule{}.ptr().
		newEmpty()
}

// NewPtr - Creates and returns a pointer to a new, fully populated
// instance of TextFieldSpecFiller. This method is identical to
// method TextFieldSpecFiller.NewConstructor() with the sole
// distinction being that this method returns a pointer to a new
// instance of TextFieldSpecFiller.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  fillerCharacters    string
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
//  fillerCharsRepeatCount    int
//     - Controls the number of times 'fillerCharacters' is
//       repeated when constructing the final Text Filler Field
//       returned by this method. The actual length of the string
//       which will populated the completed Text Filler Field is
//       equal to the length of 'fillerCharacters' times the value
//       of 'fillerCharsRepeatCount'.
//
//       Text Field Filler Length =
//          Length of fillerCharacters X fillerCharsRepeatCount
//          Example: fillerCharacters = "-*"
//                   fillerRepeatCount = 3
//                   Final Text Filler Field = "-*-*-*"
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
func (txtFillerField TextFieldSpecFiller) NewPtr(
	fillerCharacters string,
	fillerCharsCount int,
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
		"TextFieldSpecFiller.NewPtr()",
		"")

	if err != nil {
		return &TextFieldSpecFiller{}, err
	}

	if len(fillerCharacters) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fillerCharacters' is a zero length string!\n",
			ePrefix.String())
		return &TextFieldSpecFiller{}, err
	}

	if fillerCharsCount < 1 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fillerCharsRepeatCount' is less than one (1)!\n"+
			"'fillerCharsRepeatCount' controls the number of repetitions of 'fillerCharacters'\n"+
			"in the Filler Text Field.\n",
			ePrefix.String())
		return &TextFieldSpecFiller{}, err
	}

	if fillerCharsCount > 1000000 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fillerCharsRepeatCount' is greater than one million (1,000,000)!\n"+
			"'fillerCharsRepeatCount' controls the number of repetitions of 'fillerCharacters'\n"+
			"in the Filler Text Field.\n",
			ePrefix.String())
		return &TextFieldSpecFiller{}, err
	}

	fillerCharsRunes := []rune(fillerCharacters)

	lenFillerChars := len(fillerCharsRunes)

	newTxtFillerField := textFieldSpecFillerMolecule{}.ptr().newEmpty()

	newTxtFillerField.fillerCharacters =
		make([]rune, lenFillerChars)

	copy(newTxtFillerField.fillerCharacters,
		fillerCharsRunes)

	newTxtFillerField.fillerCharsRepeatCount =
		fillerCharsCount

	return &newTxtFillerField, nil
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
//          Example: fillerCharacters = "-*"
//                   fillerRepeatCount = 3
//                   Final Text Filler Field = "-*-*-*"
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
func (txtFillerField *TextFieldSpecFiller) SetTextFiller(
	fillerCharacters string,
	fillerCharsCount int,
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

	lenFillerChars := len(fillerCharsRunes)

	if lenFillerChars == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fillerCharacters' is a zero length string!\n",
			ePrefix.String())
		return err
	}

	if fillerCharsCount < 1 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fillerCharsRepeatCount' is less than one (1)!\n"+
			"'fillerCharsRepeatCount' controls the number of repetitions of 'fillerCharacters'\n"+
			"in the Filler Text Field.\n",
			ePrefix.String())
		return err
	}

	if fillerCharsCount > 1000000 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fillerCharsRepeatCount' is greater than one-million (+1,000,000)!\n"+
			"'fillerCharsRepeatCount' controls the number of repetitions of 'fillerCharacters'\n"+
			"in the Filler Text Field.\n",
			ePrefix.String())
		return err
	}

	txtFillerField.fillerCharacters =
		make([]rune, lenFillerChars)

	copy(txtFillerField.fillerCharacters,
		fillerCharsRunes)

	txtFillerField.fillerCharsRepeatCount =
		fillerCharsCount

	return nil
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
func (txtFillerField *TextFieldSpecFiller) SetTextFillerRune(
	fillerCharacter rune,
	fillerCharsCount int,
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

	if fillerCharacter == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fillerCharacter' is a rune with a zero value!\n",
			ePrefix.String())

		return err
	}

	if fillerCharsCount < 1 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fillerCharsRepeatCount' is less than one (1)!\n"+
			"'fillerCharsRepeatCount' controls the number of repetitions of 'fillerCharacters'\n"+
			"in the Filler Text Field.\n",
			ePrefix.String())
		return err
	}

	if fillerCharsCount > 1000000 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fillerCharsRepeatCount' is greater than one-million (+1,000,000)!\n"+
			"'fillerCharsRepeatCount' controls the number of repetitions of 'fillerCharacters'\n"+
			"in the Filler Text Field.\n",
			ePrefix.String())
		return err
	}

	newTxtFillerField := textFieldSpecFillerMolecule{}.ptr().
		newEmpty()

	newTxtFillerField.fillerCharacters =
		make([]rune, 1)

	newTxtFillerField.fillerCharacters[0] =
		fillerCharacter

	newTxtFillerField.fillerCharsRepeatCount =
		fillerCharsCount

	return nil
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
//         Text Field Filler Length =
//           Length of fillerCharacters X fillerCharsRepeatCount
//           Example: fillerCharacters = []rune{'-','*'}
//                    fillerRepeatCount = 3
//                    Final Text Filler Field = "-*-*-*"
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
func (txtFillerField *TextFieldSpecFiller) SetTextFillerRuneArray(
	fillerCharacters []rune,
	fillerCharsCount int,
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

	lenRuneArray := len(fillerCharacters)

	if lenRuneArray == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fillerCharacters' is a zero length rune array!\n",
			ePrefix.String())

		return err
	}

	if fillerCharsCount < 1 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fillerCharsRepeatCount' is less than one (1)!\n"+
			"'fillerCharsRepeatCount' controls the number of repetitions of 'fillerCharacters'\n"+
			"in the Filler Text Field.\n",
			ePrefix.String())
		return err
	}

	if fillerCharsCount > 1000000 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fillerCharsRepeatCount' is greater than one-million (+1,000,000)!\n"+
			"'fillerCharsRepeatCount' controls the number of repetitions of 'fillerCharacters'\n"+
			"in the Filler Text Field.\n",
			ePrefix.String())
		return err
	}

	newTxtFillerField := textFieldSpecFillerMolecule{}.ptr().
		newEmpty()

	newTxtFillerField.fillerCharacters =
		make([]rune, lenRuneArray)

	copy(
		newTxtFillerField.fillerCharacters,
		fillerCharacters)

	newTxtFillerField.fillerCharsRepeatCount =
		fillerCharsCount

	return nil
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
