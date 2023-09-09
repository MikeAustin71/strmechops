package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"strings"
	"sync"
)

// TextFieldSpecFiller - The Filler Text Field Specification is a
// single character or character sequence which is replicated
// multiple times to create the entire length of the Filler Text
// Field.
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
//	Example 1:
//	 Filler Characters = "-"
//	 Filler Characters Repeat Count = 3
//	 Formatted Text = "---"
//
//	Example 2:
//	 Filler Characters = "-*"
//	 Filler Characters Repeat Count = 3
//	 Formatted Text = "-*-*-*"
//
// The 'fillerCharsRepeatCount' integer value is the number times
// that 'fillerCharacters' is repeated in order to construct the
// Filler Text Field. Be advised that Filler Text Fields requires a
// 'fillerCharsRepeatCount' value greater than zero.
// 'fillerCharsRepeatCount' values less than or equal to zero
// constitute an error condition.
//
// # Member Variables
//
// ----------------------------------------------------------------
//
//	fillerCharacters           []rune
//	   - A rune array containing the text characters which will be
//	     included in the Text Filler Field. The final Text Filler
//	     Field will be constructed from ths filler characters
//	     repeated one or more times as specified by the
//	     'fillerCharsRepeatCount' parameter.
//
//	     The Text Field Filler final formatted text is equal to:
//	        fillerCharacters X fillerCharsRepeatCount
//	        Example: fillerCharacters = []rune{'-','*'}
//	                 fillerRepeatCount = 3
//	                 Final Text Filler Field = "-*-*-*"
//
//
//	fillerCharsRepeatCount     int
//	   - Controls the number of times 'fillerCharacters' is
//	     repeated when constructing the final Text Filler Field
//	     returned by this method. The actual length of the string
//	     which will populate the completed Text Filler Field is
//	     equal to the length of 'fillerCharacters' times the value
//	     of 'fillerCharsRepeatCount'.
//
//	      Text Field Filler Length =
//	        Length of fillerCharacters X fillerCharsRepeatCount
//
//	        Example #1: fillerCharacters = []rune{'-','*'}
//	                    fillerRepeatCount = 3
//	                    Final Text Filler Field = "-*-*-*"
//
//	        Example #2: fillerCharacters = []rune{'-'}
//	                    fillerRepeatCount = 3
//	                    Final Text Filler Field = "---"
type TextFieldSpecFiller struct {
	fillerCharacters []rune // The base characters which comprise the text filler
	//                                  //   field. See 'fillerCharsRepeatCount'.
	fillerCharsRepeatCount int // The number of times 'fillerCharacters'
	//                                  //  is repeated to create the complete filler string.
	textLineReader *strings.Reader
	lock           *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming instance of
// TextFieldSpecFiller ('incomingTxtFieldFiller') to the data fields
// of the current TextFieldSpecFiller instance ('txtFillerField').
//
// IMPORTANT
// All the data fields in current TextFieldSpecFiller instance
// ('txtFillerField') will be modified and overwritten.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	 incomingTxtFieldFiller     *TextFieldSpecFiller
//	    - A pointer to an instance of TextFieldSpecFiller. This
//	      method will NOT change the values of internal member
//	      variables contained in this instance.
//
//	      All data values in this TextFieldSpecFiller instance
//	      will be copied to current TextFieldSpecFiller
//	      instance ('txtFillerField').
//
//	      If parameter 'incomingTxtFieldFiller' is determined to be
//	      invalid, an error will be returned.
//
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
//
//		     This empty interface must be convertible to one of the
//		     following types:
//
//		     1. nil - A nil value is valid and generates an empty
//		        collection of error prefix and error context
//		        information.
//
//		     2. string - A string containing error prefix information.
//
//		     3. []string A one-dimensional slice of strings containing
//		        error prefix information
//
//		     4. [][2]string A two-dimensional slice of strings
//		        containing error prefix and error context information.
//
//		     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		        from this object will be copied for use in error and
//		        informational messages.
//
//		     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		        Information from this object will be copied for use in
//		        error and informational messages.
//
//		     7. IBasicErrorPrefix - An interface to a method generating
//		        a two-dimensional slice of strings containing error
//		        prefix and error context information.
//
//		     If parameter 'errorPrefix' is NOT convertible to one of
//		     the valid types listed above, it will be considered
//		     invalid and trigger the return of an error.
//
//		     Types ErrPrefixDto and IBasicErrorPrefix are included in
//		     the 'errpref' software package,
//		     "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If this method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
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
// ----------------------------------------------------------------
//
// Input Parameters
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
//
//		     This empty interface must be convertible to one of the
//		     following types:
//
//		     1. nil - A nil value is valid and generates an empty
//		        collection of error prefix and error context
//		        information.
//
//		     2. string - A string containing error prefix information.
//
//		     3. []string A one-dimensional slice of strings containing
//		        error prefix information
//
//		     4. [][2]string A two-dimensional slice of strings
//		        containing error prefix and error context information.
//
//		     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		        from this object will be copied for use in error and
//		        informational messages.
//
//		     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		        Information from this object will be copied for use in
//		        error and informational messages.
//
//		     7. IBasicErrorPrefix - An interface to a method generating
//		        a two-dimensional slice of strings containing error
//		        prefix and error context information.
//
//		     If parameter 'errorPrefix' is NOT convertible to one of
//		     the valid types listed above, it will be considered
//		     invalid and trigger the return of an error.
//
//		     Types ErrPrefixDto and IBasicErrorPrefix are included in
//		     the 'errpref' software package,
//		     "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	TextFieldSpecFiller
//	   - If this method completes successfully and no errors are
//	     encountered, this parameter will return a deep copy of the
//	     current TextFieldSpecFiller instance.
//
//
//	error
//	   - If this method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error occurs, the text value of input parameter
//	     'errorPrefix' will be inserted or prefixed at the
//	     beginning of the error message.
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
// ----------------------------------------------------------------
//
// Input Parameters
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
//
//		     This empty interface must be convertible to one of the
//		     following types:
//
//		     1. nil - A nil value is valid and generates an empty
//		        collection of error prefix and error context
//		        information.
//
//		     2. string - A string containing error prefix information.
//
//		     3. []string A one-dimensional slice of strings containing
//		        error prefix information
//
//		     4. [][2]string A two-dimensional slice of strings
//		        containing error prefix and error context information.
//
//		     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		        from this object will be copied for use in error and
//		        informational messages.
//
//		     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		        Information from this object will be copied for use in
//		        error and informational messages.
//
//		     7. IBasicErrorPrefix - An interface to a method generating
//		        a two-dimensional slice of strings containing error
//		        prefix and error context information.
//
//		     If parameter 'errorPrefix' is NOT convertible to one of
//		     the valid types listed above, it will be considered
//		     invalid and trigger the return of an error.
//
//		     Types ErrPrefixDto and IBasicErrorPrefix are included in
//		     the 'errpref' software package,
//		     "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	ITextFieldSpecification
//	   - If this method completes successfully and no errors are
//	     encountered, this parameter will return a deep copy of
//	     the current TextFieldSpecFiller instance cast as an
//	     ITextFieldSpecification object.
//
//
//	error
//	   - If this method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error occurs, the text value of input parameter
//	     'errorPrefix' will be inserted or prefixed at the
//	     beginning of the error message.
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
// ----------------------------------------------------------------
//
// Input Parameters
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
//
//		     This empty interface must be convertible to one of the
//		     following types:
//
//		     1. nil - A nil value is valid and generates an empty
//		        collection of error prefix and error context
//		        information.
//
//		     2. string - A string containing error prefix information.
//
//		     3. []string A one-dimensional slice of strings containing
//		        error prefix information
//
//		     4. [][2]string A two-dimensional slice of strings
//		        containing error prefix and error context information.
//
//		     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		        from this object will be copied for use in error and
//		        informational messages.
//
//		     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		        Information from this object will be copied for use in
//		        error and informational messages.
//
//		     7. IBasicErrorPrefix - An interface to a method generating
//		        a two-dimensional slice of strings containing error
//		        prefix and error context information.
//
//		     If parameter 'errorPrefix' is NOT convertible to one of
//		     the valid types listed above, it will be considered
//		     invalid and trigger the return of an error.
//
//		     Types ErrPrefixDto and IBasicErrorPrefix are included in
//		     the 'errpref' software package,
//		     "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	*TextFieldSpecFiller
//	   - If this method completes successfully and no errors are
//	     encountered, this parameter will return a pointer to a
//	     deep copy of the current TextFieldSpecFiller instance.
//
//
//	error
//	   - If this method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error occurs, the text value of input parameter
//	     'errorPrefix' will be inserted or prefixed at the
//	     beginning of the error message.
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
// This method is identical to method
// TextFieldSpecFiller.GetFillerRunes() with the sole exception
// being that this method returns a string while method
// TextFieldSpecFiller.GetFillerRunes() returns an array of runes.
//
// The filler characters are used to populate the Text Filler
// Field formatted text. The final Text Filler Field will be
// constructed from ths filler characters repeated one or more
// times as specified by the 'fillerCharsRepeatCount' parameter.
//
// The Text Field Filler final formatted text is equal to:
//
//	        fillerCharacters X fillerCharsRepeatCount
//
//	Example 1:
//	 Filler Characters Array = "-"
//	 Filler Characters Repeat Count = 3
//	 Formatted Text = "---"
//
//	Example 2:
//	 Filler Characters Array = "-*"
//	 Filler Characters Repeat Count = 3
//	 Formatted Text = "-*-*-*"
//
// This method returns a string containing the Filler Characters.
func (txtFillerField *TextFieldSpecFiller) GetFillerChars() string {

	if txtFillerField.lock == nil {
		txtFillerField.lock = new(sync.Mutex)
	}

	txtFillerField.lock.Lock()

	defer txtFillerField.lock.Unlock()

	if len(txtFillerField.fillerCharacters) == 0 {
		return ""
	}

	return string(txtFillerField.fillerCharacters)
}

// GetFillerRunes - Returns the internal member variable
// TextFieldSpecFiller.fillerCharacters as an array of runes.
//
// This method is identical to method
// TextFieldSpecFiller.GetFillerChars() with the sole exception
// being that this method returns an array of runes while method
// TextFieldSpecFiller.GetFillerChars() returns a string.
//
// The filler characters are used to populate the Text Filler
// Field formatted text. The final Text Filler Field will be
// constructed from ths filler characters repeated one or more
// times as specified by the 'fillerCharsRepeatCount' parameter.
//
// The Text Field Filler final formatted text is equal to:
//
//	        fillerCharacters X fillerCharsRepeatCount
//
//	Example #1: fillerCharacters = []rune{'-','*'}
//	            fillerRepeatCount = 3
//	            Final Text Filler Field = "-*-*-*"
//
//	Example #2: fillerCharacters = []rune{'-'}
//	            fillerRepeatCount = 3
//	            Final Text Filler Field = "---"
//
// This method returns an array of runes containing the Filler
// Characters.
func (txtFillerField *TextFieldSpecFiller) GetFillerRunes() []rune {

	if txtFillerField.lock == nil {
		txtFillerField.lock = new(sync.Mutex)
	}

	txtFillerField.lock.Lock()

	defer txtFillerField.lock.Unlock()

	newRuneArray := make([]rune, 0)

	sMechPreon := strMechPreon{}

	err := sMechPreon.copyRuneArrays(
		&newRuneArray,
		&txtFillerField.fillerCharacters,
		true,
		nil)

	if err != nil {
		newRuneArray = nil
	}

	return newRuneArray
}

// GetFillerCharsRepeatCount - This method returns the integer
// value of the internal member variable:
//
//	TextFieldSpecFiller.fillerCharsRepeatCount
//
// The Filler Characters Repeat Count is the number of times that
// the Filler Characters are repeated when constructing the Text
// Filler Field formatted text. The final formatted text is equal
// to:  Filler Characters  X  Filler Characters Repeat Count
//
//	Example: fillerCharacters = "-*"
//	         fillerRepeatCount = 3
//	         Final Text Filler Field = "-*-*-*"
//
// This method returns the Filler Characters Repeat Count.
func (txtFillerField *TextFieldSpecFiller) GetFillerCharsRepeatCount() int {

	if txtFillerField.lock == nil {
		txtFillerField.lock = new(sync.Mutex)
	}

	txtFillerField.lock.Lock()

	defer txtFillerField.lock.Unlock()

	return txtFillerField.fillerCharsRepeatCount
}

// GetFormattedStrLength - Returns the string length of the
// formatted text generated by the current instance of
// TextFieldSpecFiller. Effectively, this is the length of the
// strings returned by methods:
//
//	TextFieldSpecFiller.GetFormattedText()
//	TextFieldSpecFiller.String()
//
// If an error is encountered, this method returns a value of minus
// one (-1).
func (txtFillerField *TextFieldSpecFiller) GetFormattedStrLength() int {

	if txtFillerField.lock == nil {
		txtFillerField.lock = new(sync.Mutex)
	}

	txtFillerField.lock.Lock()

	defer txtFillerField.lock.Unlock()

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TextFieldSpecFiller.GetFormattedStrLength()",
		"")

	formattedTextStr,
		err := textFieldSpecFillerMolecule{}.ptr().
		getFormattedText(
			txtFillerField,
			ePrefix.XCpy(
				"txtFillerField"))

	if err != nil {
		return -1
	}

	return len(formattedTextStr)
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
//	Filler Characters Array Length  X
//	             Filler Characters Repeat Count
//
// This method is similar to method TextFieldSpecFiller.String()
// with the exception that this method returns an error.
//
// This method fulfills the requirements of the
// ITextFieldSpecification interface.
//
// Methods which return formatted text are listed as follows:
//
//	TextFieldSpecFiller.String()
//	TextFieldSpecFiller.GetFormattedText()
//	TextFieldSpecFiller.TextBuilder()
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
//
//		     This empty interface must be convertible to one of the
//		     following types:
//
//		     1. nil - A nil value is valid and generates an empty
//		        collection of error prefix and error context
//		        information.
//
//		     2. string - A string containing error prefix information.
//
//		     3. []string A one-dimensional slice of strings containing
//		        error prefix information
//
//		     4. [][2]string A two-dimensional slice of strings
//		        containing error prefix and error context information.
//
//		     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		        from this object will be copied for use in error and
//		        informational messages.
//
//		     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		        Information from this object will be copied for use in
//		        error and informational messages.
//
//		     7. IBasicErrorPrefix - An interface to a method generating
//		        a two-dimensional slice of strings containing error
//		        prefix and error context information.
//
//		     If parameter 'errorPrefix' is NOT convertible to one of
//		     the valid types listed above, it will be considered
//		     invalid and trigger the return of an error.
//
//		     Types ErrPrefixDto and IBasicErrorPrefix are included in
//		     the 'errpref' software package,
//		     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	string
//	   - The formatted text string generated by the current
//	     instance of TextFieldSpecFiller.
//
//
//	error
//	   - If the method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
//
// ----------------------------------------------------------------
//
// Example Usage
//
//	Example 1:
//	 Filler Characters Array = []rune{'-'}
//	 Filler Characters Repeat Count = 3
//	 Formatted Text = "---"
//
//	Example 2:
//	 Filler Characters Array = []rune{'-','*'}
//	 Filler Characters Repeat Count = 3
//	 Formatted Text = "-*-*-*"
func (txtFillerField *TextFieldSpecFiller) GetFormattedText(
	errorPrefix interface{}) (
	string,
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
		"TextFieldSpecFiller.GetFormattedText()",
		"")

	if err != nil {
		return "", err
	}

	return textFieldSpecFillerMolecule{}.ptr().
		getFormattedText(
			txtFillerField,
			ePrefix.XCpy(
				"txtFillerField"))
}

// IsValidInstance - Performs a diagnostic review of the data
// values encapsulated in the current TextFieldSpecFiller instance
// to determine if they are valid.
//
// If all data element evaluate as valid, this method returns
// 'true'. If any data element is invalid, this method returns
// 'false'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	--- NONE ---
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	isValid             bool
//	   - If all data elements encapsulated by the current instance
//	     of TextFieldSpecFiller are valid, this returned boolean
//	     value is set to 'true'. If any data values are invalid,
//	     this return parameter is set to 'false'.
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
// ----------------------------------------------------------------
//
// Input Parameters
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
//
//		     This empty interface must be convertible to one of the
//		     following types:
//
//		     1. nil - A nil value is valid and generates an empty
//		        collection of error prefix and error context
//		        information.
//
//		     2. string - A string containing error prefix information.
//
//		     3. []string A one-dimensional slice of strings containing
//		        error prefix information
//
//		     4. [][2]string A two-dimensional slice of strings
//		        containing error prefix and error context information.
//
//		     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		        from this object will be copied for use in error and
//		        informational messages.
//
//		     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		        Information from this object will be copied for use in
//		        error and informational messages.
//
//		     7. IBasicErrorPrefix - An interface to a method generating
//		        a two-dimensional slice of strings containing error
//		        prefix and error context information.
//
//		     If parameter 'errorPrefix' is NOT convertible to one of
//		     the valid types listed above, it will be considered
//		     invalid and trigger the return of an error.
//
//		     Types ErrPrefixDto and IBasicErrorPrefix are included in
//		     the 'errpref' software package,
//		     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If any of the internal member data variables contained in
//	     the current instance of TextFieldSpecFiller are found to be
//	     invalid, this method will return an error  along with an
//	     appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' (error prefix) will be inserted or
//	     prefixed at the beginning of the error message.
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
// The Filler Text Field consists of a single character or multiple
// character sequence which is replicated some number of times to
// create the entire length of the Filler Text Field.
//
// The Filler Text Field
// Be advised that setting member variables to their zero values
// means that the returned TextFieldSpecFiller instance is invalid.
// Therefore, in order to use this TextFieldSpecFiller instance,
// users must later call the setter methods on this type in order
// to configure valid and meaningful member variable data values.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	--- NONE ---
//
// -----------------------------------------------------------------
//
// Return Values
//
//	TextFieldSpecFiller
//	   - This parameter returns a new and empty concrete instance
//	     of TextFieldSpecFiller. Member variable data values are
//	     set to their initial or zero values.
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
// The Filler Text Field consists of a single character or multiple
// character sequence which is replicated some number of times to
// create the entire length of the Filler Text Field.
//
// Be advised that setting member variables to their zero values
// means that the returned TextFieldSpecFiller instance is invalid.
// Therefore, in order to use this TextFieldSpecFiller instance,
// users must later call the setter methods on this type in order
// to configure valid and meaningful member variable data values.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	--- NONE ---
//
// -----------------------------------------------------------------
//
// Return Values
//
//	*TextFieldSpecFiller
//	   - This parameter returns a pointer to a new, empty instance
//	     of TextFieldSpecFiller. Member variable data values are
//	     set to their initial or zero values.
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
// The Filler Text Field consists of a single character or multiple
// character sequence which is replicated some number of times to
// create the entire length of the Filler Text Field.
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
//	 fillerCharacters           string
//	    - A string containing the text characters which will be
//	      included in the Text Filler Field. The final Text Filler
//	      Field will be constructed from the filler characters
//	      repeated one or more times as specified by the
//	      'fillerCharsRepeatCount' parameter.
//
//	      The Text Field Filler final formatted text is equal to:
//	         fillerCharacter X fillerCharsRepeatCount
//	         Example: fillerCharacters = "-*"
//	                  fillerRepeatCount = 3
//	                  Final Text Filler Field = "-*-*-*"
//
//	      If 'fillerCharacters' is submitted as an empty or zero
//	      length string, this method will return an error.
//
//
//	 fillerCharsRepeatCount     int
//	    - Controls the number of times 'fillerCharacters' is
//	      repeated when constructing the final Text Filler Field
//	      returned by this method. The actual length of the string
//	      which will populated the completed Text Filler Field is
//	      equal to the length of 'fillerCharacters' times the value
//	      of 'fillerCharsRepeatCount'.
//
//	       Text Field Filler Length =
//	         Length of fillerCharacters X fillerCharsRepeatCount
//
//	         Example #1: fillerCharacters = "-*"
//	                     fillerRepeatCount = 3
//	                     Final Text Filler Field = "-*-*-*"
//
//	         Example #2: fillerCharacters = "-"
//	                     fillerRepeatCount = 3
//	                     Final Text Filler Field = "---"
//
//	      If 'fillerCharsRepeatCount' has a value less than one (1) or
//	      greater than one-million (1,000,000), an error will be
//	      returned.
//
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
//
//		     This empty interface must be convertible to one of the
//		     following types:
//
//		     1. nil - A nil value is valid and generates an empty
//		        collection of error prefix and error context
//		        information.
//
//		     2. string - A string containing error prefix information.
//
//		     3. []string A one-dimensional slice of strings containing
//		        error prefix information
//
//		     4. [][2]string A two-dimensional slice of strings
//		        containing error prefix and error context information.
//
//		     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		        from this object will be copied for use in error and
//		        informational messages.
//
//		     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		        Information from this object will be copied for use in
//		        error and informational messages.
//
//		     7. IBasicErrorPrefix - An interface to a method generating
//		        a two-dimensional slice of strings containing error
//		        prefix and error context information.
//
//		     If parameter 'errorPrefix' is NOT convertible to one of
//		     the valid types listed above, it will be considered
//		     invalid and trigger the return of an error.
//
//		     Types ErrPrefixDto and IBasicErrorPrefix are included in
//		     the 'errpref' software package,
//		     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	*TextFieldSpecFiller
//	   - If this method completes successfully, this parameter will
//	     return a pointer to a valid and fully populated Text Filler
//	     Field object.
//
//
//	error
//	   - If this method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//	Example 1:
//	 Filler Characters = "-"
//	 Filler Characters Repeat Count = 3
//	 Formatted Text = "---"
//
//	Example 2:
//	 Filler Characters = "-*"
//	 Filler Characters Repeat Count = 3
//	 Formatted Text = "-*-*-*"
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
// The Filler Text Field consists of a single character or multiple
// character sequence which is replicated some number of times to
// create the entire length of the Filler Text Field.
//
// This method is identical to
// TextFieldSpecFiller.NewTextFillerRune() with the sole exception
// being that this method returns a pointer to an instance of
// TextFieldSpecFiller and TextFieldSpecFiller.NewTextFillerRune()
// returns a concrete instance of TextFieldSpecFiller.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	 fillerCharacter            rune
//	    - A rune containing the text character which will be
//	      included in the Text Filler Field. The final Text Filler
//	      Field will be constructed from ths filler character
//	      repeated one or more times as specified by the
//	      'fillerCharsRepeatCount' parameter.
//
//	      The Text Field Filler final formatted text is equal to:
//	          fillerCharacter X fillerCharsRepeatCount
//	          Example: fillerCharacter = '-'
//	                   fillerRepeatCount = 3
//	                   Final Text Filler Field = "---"
//
//	      If 'fillerCharacter' is submitted with a zero value,
//	      this method will return an error.
//
//
//	 fillerCharsRepeatCount     int
//	    - Controls the number of times 'fillerCharacter' is
//	      repeated when constructing the final Text Filler Field
//	      returned by this method. The actual length of the string
//	      which will populated the completed Text Filler Field is
//	      equal to the length of 'fillerCharacter' (1) times the
//	      value of 'fillerCharsRepeatCount'.
//
//	        Text Field Filler Length =
//	          Length of fillerCharacter (1) X fillerCharsRepeatCount
//	          Example: fillerCharacter = '-'
//	                   fillerRepeatCount = 3
//	                   Final Text Filler Field = "---"
//
//	      If 'fillerCharsRepeatCount' has a value less than one (1) or
//	      greater than one-million (1,000,000), an error will be
//	      returned.
//
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
//
//		     This empty interface must be convertible to one of the
//		     following types:
//
//		     1. nil - A nil value is valid and generates an empty
//		        collection of error prefix and error context
//		        information.
//
//		     2. string - A string containing error prefix information.
//
//		     3. []string A one-dimensional slice of strings containing
//		        error prefix information
//
//		     4. [][2]string A two-dimensional slice of strings
//		        containing error prefix and error context information.
//
//		     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		        from this object will be copied for use in error and
//		        informational messages.
//
//		     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		        Information from this object will be copied for use in
//		        error and informational messages.
//
//		     7. IBasicErrorPrefix - An interface to a method generating
//		        a two-dimensional slice of strings containing error
//		        prefix and error context information.
//
//		     If parameter 'errorPrefix' is NOT convertible to one of
//		     the valid types listed above, it will be considered
//		     invalid and trigger the return of an error.
//
//		     Types ErrPrefixDto and IBasicErrorPrefix are included in
//		     the 'errpref' software package,
//		     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	*TextFieldSpecFiller
//	   - If this method completes successfully, this parameter will
//	     return a pointer to a new, valid and fully populated Text
//	     Filler Field.
//
//
//	error
//	   - If this method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//	Example 1:
//	 Filler Character = '-'
//	 Filler Characters Repeat Count = 3
//	 Formatted Text = "---"
//
//	Example 2:
//	 Filler Character = '*'
//	 Filler Characters Repeat Count = 3
//	 Formatted Text = "***"
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
// The Filler Text Field consists of a single character or multiple
// character sequence which is replicated some number of times to
// create the entire length of the Filler Text Field.
//
// This method is identical to
// TextFieldSpecFiller.NewTextFillerRuneArray() with the sole
// exception being that this method returns a pointer to an
// instance of TextFieldSpecFiller and
// TextFieldSpecFiller.NewTextFillerRuneArray() returns a concrete
// instance of TextFieldSpecFiller.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	 fillerCharacters           []rune
//	    - A rune array containing the text characters which will be
//	      included in the Text Filler Field. The final Text Filler
//	      Field will be constructed from ths filler characters
//	      repeated one or more times as specified by the
//	      'fillerCharsRepeatCount' parameter.
//
//	      The Text Field Filler final formatted text is equal to:
//	         fillerCharacters X fillerCharsRepeatCount
//	         Example: fillerCharacters = []rune{'-','*'}
//	                  fillerRepeatCount = 3
//	                  Final Text Filler Field = "-*-*-*"
//
//	      If 'fillerCharacters' is submitted with a zero length rune
//	      array, this method will return an error.
//
//
//	 fillerCharsRepeatCount     int
//	    - Controls the number of times 'fillerCharacters' is
//	      repeated when constructing the final Text Filler Field
//	      returned by this method. The actual length of the string
//	      which will populated the completed Text Filler Field is
//	      equal to the length of 'fillerCharacters' times the value
//	      of 'fillerCharsRepeatCount'.
//
//	       Text Field Filler Length =
//	         Length of fillerCharacters X fillerCharsRepeatCount
//
//	         Example #1: fillerCharacters = []rune{'-','*'}
//	                     fillerRepeatCount = 3
//	                     Final Text Filler Field = "-*-*-*"
//
//	         Example #2: fillerCharacters = []rune{'-'}
//	                     fillerRepeatCount = 3
//	                     Final Text Filler Field = "---"
//
//	      If 'fillerCharsRepeatCount' has a value less than one (1) or
//	      greater than one-million (1,000,000), an error will be
//	      returned.
//
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
//
//		     This empty interface must be convertible to one of the
//		     following types:
//
//		     1. nil - A nil value is valid and generates an empty
//		        collection of error prefix and error context
//		        information.
//
//		     2. string - A string containing error prefix information.
//
//		     3. []string A one-dimensional slice of strings containing
//		        error prefix information
//
//		     4. [][2]string A two-dimensional slice of strings
//		        containing error prefix and error context information.
//
//		     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		        from this object will be copied for use in error and
//		        informational messages.
//
//		     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		        Information from this object will be copied for use in
//		        error and informational messages.
//
//		     7. IBasicErrorPrefix - An interface to a method generating
//		        a two-dimensional slice of strings containing error
//		        prefix and error context information.
//
//		     If parameter 'errorPrefix' is NOT convertible to one of
//		     the valid types listed above, it will be considered
//		     invalid and trigger the return of an error.
//
//		     Types ErrPrefixDto and IBasicErrorPrefix are included in
//		     the 'errpref' software package,
//		     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	*TextFieldSpecFiller
//	   - If this method completes successfully, this parameter will
//	     return a pointer to a new, valid and fully populated Text
//	     Filler Field.
//
//
//	error
//	   - If this method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//	Example 1:
//	 Filler Characters Array = []rune{'-'}
//	 Filler Characters Repeat Count = 3
//	 Formatted Text = "---"
//
//	Example 2:
//	 Filler Characters Array = []rune{'-','*'}
//	 Filler Characters Repeat Count = 3
//	 Formatted Text = "-*-*-*"
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
// The Filler Text Field consists of a single character or multiple
// character sequence which is replicated some number of times to
// create the entire length of the Filler Text Field.
//
// This method is identical to
// TextFieldSpecFiller.NewPtrTextFiller() with the sole exception
// being that this method returns a concrete instance of
// TextFieldSpecFiller and TextFieldSpecFiller.NewPtrTextFiller()
// returns a pointer to a TextFieldSpecLabel instance.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	 fillerCharacters           string
//	    - A string containing the text characters which will be
//	      included in the Text Filler Field. The final Text Filler
//	      Field will be constructed from the filler characters
//	      repeated one or more times as specified by the
//	      'fillerCharsRepeatCount' parameter.
//
//	       Text Field Filler Length =
//	         Length of fillerCharacters X fillerCharsRepeatCount
//
//	         Example #1: fillerCharacters = "-*"
//	                     fillerRepeatCount = 3
//	                     Final Text Filler Field = "-*-*-*"
//
//	         Example #2: fillerCharacters = "-"
//	                     fillerRepeatCount = 3
//	                     Final Text Filler Field = "---"
//
//	      If 'fillerCharacters' is submitted as an empty or zero
//	      length string, this method will return an error.
//
//
//	 fillerCharsRepeatCount     int
//	    - Controls the number of times 'fillerCharacters' is
//	      repeated when constructing the final Text Filler Field
//	      returned by this method. The actual length of the string
//	      which will populated the completed Text Filler Field is
//	      equal to the length of 'fillerCharacters' times the value
//	      of 'fillerCharsRepeatCount'.
//
//	       Text Field Filler Length =
//	         Length of fillerCharacters X fillerCharsRepeatCount
//
//	         Example #1: fillerCharacters = "-*"
//	                     fillerRepeatCount = 3
//	                     Final Text Filler Field = "-*-*-*"
//
//	         Example #2: fillerCharacters = "-"
//	                     fillerRepeatCount = 3
//	                     Final Text Filler Field = "---"
//
//	      If 'fillerCharsRepeatCount' has a value less than one (1) or
//	      greater than one-million (1,000,000), an error will be
//	      returned.
//
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
//
//		     This empty interface must be convertible to one of the
//		     following types:
//
//		     1. nil - A nil value is valid and generates an empty
//		        collection of error prefix and error context
//		        information.
//
//		     2. string - A string containing error prefix information.
//
//		     3. []string A one-dimensional slice of strings containing
//		        error prefix information
//
//		     4. [][2]string A two-dimensional slice of strings
//		        containing error prefix and error context information.
//
//		     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		        from this object will be copied for use in error and
//		        informational messages.
//
//		     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		        Information from this object will be copied for use in
//		        error and informational messages.
//
//		     7. IBasicErrorPrefix - An interface to a method generating
//		        a two-dimensional slice of strings containing error
//		        prefix and error context information.
//
//		     If parameter 'errorPrefix' is NOT convertible to one of
//		     the valid types listed above, it will be considered
//		     invalid and trigger the return of an error.
//
//		     Types ErrPrefixDto and IBasicErrorPrefix are included in
//		     the 'errpref' software package,
//		     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	TextFieldSpecFiller
//	   - If this method completes successfully, this parameter will
//	     return a new concrete Text Filler Field object.
//
//
//	error
//	   - If this method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//	Example 1:
//	 Filler Characters = "-"
//	 Filler Characters Repeat Count = 3
//	 Formatted Text = "---"
//
//	Example 2:
//	 Filler Characters = "-*"
//	 Filler Characters Repeat Count = 3
//	 Formatted Text = "-*-*-*"
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
// The Filler Text Field consists of a single character or multiple
// character sequence which is replicated some number of times to
// create the entire length of the Filler Text Field.
//
// This method is identical to
// TextFieldSpecFiller.NewPtrTextFillerRune() with the sole
// exception being that this method returns a concrete instance of
// TextFieldSpecFiller and TextFieldSpecFiller.NewPtrTextFillerRune()
// returns a pointer to a TextFieldSpecFiller instance.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	 fillerCharacter            rune
//	    - A rune containing the text character which will be
//	      included in the Text Filler Field. The final Text Filler
//	      Field will be constructed from ths filler character
//	      repeated one or more times as specified by the
//	      'fillerCharsRepeatCount' parameter.
//
//	      The Text Field Filler final formatted text is equal to:
//	          fillerCharacter X fillerCharsRepeatCount
//	          Example: fillerCharacter = '-'
//	                   fillerRepeatCount = 3
//	                   Final Text Filler Field = "---"
//
//	      If 'fillerCharacter' is submitted with a zero value,
//	      this method will return an error.
//
//
//	 fillerCharsRepeatCount     int
//	    - Controls the number of times 'fillerCharacter' is
//	      repeated when constructing the final Text Filler Field
//	      returned by this method. The actual length of the string
//	      which will populated the completed Text Filler Field is
//	      equal to the length of 'fillerCharacter' (1) times the
//	      value of 'fillerCharsRepeatCount'.
//
//	        Text Field Filler Length =
//	          Length of fillerCharacter (1) X fillerCharsRepeatCount
//	          Example: fillerCharacter = '-'
//	                   fillerRepeatCount = 3
//	                   Final Text Filler Field = "---"
//
//	      If 'fillerCharsRepeatCount' has a value less than one (1)
//	      or greater than one-million (1,000,000), an error will be
//	      returned.
//
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
//
//		     This empty interface must be convertible to one of the
//		     following types:
//
//		     1. nil - A nil value is valid and generates an empty
//		        collection of error prefix and error context
//		        information.
//
//		     2. string - A string containing error prefix information.
//
//		     3. []string A one-dimensional slice of strings containing
//		        error prefix information
//
//		     4. [][2]string A two-dimensional slice of strings
//		        containing error prefix and error context information.
//
//		     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		        from this object will be copied for use in error and
//		        informational messages.
//
//		     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		        Information from this object will be copied for use in
//		        error and informational messages.
//
//		     7. IBasicErrorPrefix - An interface to a method generating
//		        a two-dimensional slice of strings containing error
//		        prefix and error context information.
//
//		     If parameter 'errorPrefix' is NOT convertible to one of
//		     the valid types listed above, it will be considered
//		     invalid and trigger the return of an error.
//
//		     Types ErrPrefixDto and IBasicErrorPrefix are included in
//		     the 'errpref' software package,
//		     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	TextFieldSpecFiller
//	   - If this method completes successfully, this parameter will
//	     return a new, concrete instance of TextFieldSpecFiller.
//
//
//	error
//	   - If this method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//	Example 1:
//	 Filler Characters Array = '-'
//	 Filler Characters Repeat Count = 3
//	 Formatted Text = "---"
//
//	Example 2:
//	 Filler Characters Array = '*'
//	 Filler Characters Repeat Count = 3
//	 Formatted Text = "***"
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
// The Filler Text Field consists of a single character or multiple
// character sequence which is replicated some number of times to
// create the entire length of the Filler Text Field.
//
// This method is identical to
// TextFieldSpecFiller.NewPtrTextFillerRuneArray() with the sole
// exception being that this method returns a concrete instance of
// TextFieldSpecFiller while
// TextFieldSpecFiller.NewPtrTextFillerRuneArray() returns a pointer
// to an instance of TextFieldSpecFiller.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	 fillerCharacters           []rune
//	    - A rune array containing the text characters which will be
//	      included in the Text Filler Field. The final Text Filler
//	      Field will be constructed from these filler characters
//	      repeated one or more times as specified by the
//	      'fillerCharsRepeatCount' parameter.
//
//	      The Text Field Filler final formatted text is equal to:
//	         fillerCharacters X fillerCharsRepeatCount
//	         Example: fillerCharacters = []rune{'-','*'}
//	                  fillerRepeatCount = 3
//	                  Final Text Filler Field = "-*-*-*"
//
//	      If 'fillerCharacters' is submitted with a zero length rune
//	      array, this method will return an error.
//
//
//	 fillerCharsRepeatCount     int
//	    - Controls the number of times 'fillerCharacters' is
//	      repeated when constructing the final Text Filler Field
//	      returned by this method. The actual length of the string
//	      which will populated the completed Text Filler Field is
//	      equal to the length of 'fillerCharacters' times the value
//	      of 'fillerCharsRepeatCount'.
//
//	       Text Field Filler Length =
//	         Length of fillerCharacters X fillerCharsRepeatCount
//
//	         Example #1: fillerCharacters = []rune{'-','*'}
//	                     fillerRepeatCount = 3
//	                     Final Text Filler Field = "-*-*-*"
//
//	         Example #2: fillerCharacters = []rune{'-'}
//	                     fillerRepeatCount = 3
//	                     Final Text Filler Field = "---"
//
//	      If 'fillerCharsRepeatCount' has a value less than one (1) or
//	      greater than one-million (1,000,000), an error will be
//	      returned.
//
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
//
//		     This empty interface must be convertible to one of the
//		     following types:
//
//		     1. nil - A nil value is valid and generates an empty
//		        collection of error prefix and error context
//		        information.
//
//		     2. string - A string containing error prefix information.
//
//		     3. []string A one-dimensional slice of strings containing
//		        error prefix information
//
//		     4. [][2]string A two-dimensional slice of strings
//		        containing error prefix and error context information.
//
//		     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		        from this object will be copied for use in error and
//		        informational messages.
//
//		     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		        Information from this object will be copied for use in
//		        error and informational messages.
//
//		     7. IBasicErrorPrefix - An interface to a method generating
//		        a two-dimensional slice of strings containing error
//		        prefix and error context information.
//
//		     If parameter 'errorPrefix' is NOT convertible to one of
//		     the valid types listed above, it will be considered
//		     invalid and trigger the return of an error.
//
//		     Types ErrPrefixDto and IBasicErrorPrefix are included in
//		     the 'errpref' software package,
//		     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	TextFieldSpecFiller
//	   - If this method completes successfully, this parameter will
//	     return a new concrete instance of TextFieldSpecFiller.
//
//
//	error
//	   - If this method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//	Example 1:
//	 Filler Characters Array = []rune{'-'}
//	 Filler Characters Repeat Count = 3
//	 Formatted Text = "---"
//
//	Example 2:
//	 Filler Characters Array = []rune{'-','*'}
//	 Filler Characters Repeat Count = 3
//	 Formatted Text = "-*-*-*"
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

// Read - Implements the io.Reader interface for type
// TextFieldSpecFiller.
//
// The formatted text string generated by the current instance of
// TextFieldSpecFiller will be written to the byte buffer 'p'. If
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
// ----------------------------------------------------------------
//
// Input Parameters
//
//	p                          []byte
//	   - The byte buffer into which the formatted text string
//	     generated by the current txtFillerField instance will be
//	     written.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	n                          int
//	   - The number of bytes written to byte buffer 'p'.
//
//	     Read() reads up to len(p) bytes into p. It returns
//	     the number of bytes read (0 <= n <= len(p)) and any error
//	     encountered. Even if Read() returns n < len(p), it may use
//	     all of 'p' as scratch space during the call. If some
//	     data is available but not len(p) bytes, Read()
//	     conventionally returns what is available instead of
//	     waiting for more.
//
//
//	err                        error
//	   - If this method completes successfully, this returned error
//	     Type is set equal to 'nil'. If errors are encountered
//	     during processing, the returned error Type will
//	     encapsulate an error message.
//
//	     When Read() encounters an error or end-of-file condition
//	     after successfully reading n > 0 bytes, it returns the
//	     number of bytes read. It may return the (non-nil) error
//	     from the same call or return the error (and n == 0) from
//	     a subsequent call. An instance of this general case is
//	     that a Reader returning a non-zero number of bytes at the
//	     end of the input stream may return either err == EOF or
//	     err == nil. The next read operation should return 0, EOF.
//
// ------------------------------------------------------------------------
//
// Usage Examples:
//
//	Example # 1
//
//	p := make([]byte, 50)
//
//	var n, readBytesCnt int
//	sb := strings.Builder{}
//
//	for {
//
//	  n,
//	  err = txtFillerField01.Read(p)
//
//	  if n == 0 {
//	    break
//	  }
//
//	  sb.Write(p[:n])
//	  readBytesCnt += n
//	}
//
//	if err != nil &&
//	  err != io.EOF {
//	   return fmt.Errorf(
//	    "Error Returned From txtFillerField01.Read(p)\n"+
//	    "Error = \n%v\n",
//	     err.Error())
//	}
//
//	fmt.Printf("Text Line String: %s\n",
//	              sb.String())
//
//	fmt.Printf("Number of bytes Read: %v\n",
//	              readBytesCnt)
//
//	Example # 2
//
//	p := make([]byte, 50)
//
//	var n, readBytesCnt int
//	var actualStr string
//
//	for {
//
//	  n,
//	  err = txtFillerField01.Read(p)
//
//	  if n == 0 {
//	    break
//	  }
//
//	  actualStr += string(p[:n])
//	  readBytesCnt += n
//	}
//
//	if err != nil &&
//	  err != io.EOF {
//	   return fmt.Errorf(
//	    "Error Returned From txtFillerField01.Read(p)\n"+
//	    "Error = \n%v\n",
//	     err.Error())
//	}
//
//	fmt.Printf("Text Line String: %v\n",
//	              actualStr)
//
//	fmt.Printf("Number of bytes Read: %v\n",
//	              readBytesCnt)
func (txtFillerField *TextFieldSpecFiller) Read(
	p []byte) (
	n int,
	err error) {

	if txtFillerField.lock == nil {
		txtFillerField.lock = new(sync.Mutex)
	}

	txtFillerField.lock.Lock()

	defer txtFillerField.lock.Unlock()

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TextFieldSpecFiller.Read()",
		"")

	if txtFillerField.textLineReader == nil {

		var formattedText string

		formattedText,
			err = textFieldSpecFillerMolecule{}.ptr().
			getFormattedText(
				txtFillerField,
				ePrefix.XCpy("txtFillerField"))

		if err != nil {
			return n, err
		}

		txtFillerField.textLineReader =
			strings.NewReader(formattedText)

		if txtFillerField.textLineReader == nil {
			err = fmt.Errorf("%v\n"+
				"Error: strings.NewReader(formattedText)\n"+
				"returned a nil pointer.\n"+
				"txtFillerField.textLineReader == nil\n",
				ePrefix.String())

			return n, err
		}
	}

	n,
		err = new(textSpecificationAtom).
		readBytes(
			txtFillerField.textLineReader,
			p,
			ePrefix.XCpy(
				"p -> txtFillerField.textLineReader"))

	if err == io.EOF {

		txtFillerField.textLineReader = nil

	}

	return n, err
}

// ReaderInitialize - This method will reset the internal member
// variable 'TextFieldSpecFiller.textLineReader' to its initial
// zero state of 'nil'. Effectively, this resets the internal
// strings.Reader object for use in future read operations.
//
// This method is rarely used or needed. It provides a means of
// reinitializing the internal strings.Reader object in case an
// error occurs during a read operation initiated by method
// TextFieldSpecFiller.Read().
//
// Calling this method cleans up the residue from an aborted read
// operation and prepares the strings.Reader object for future read
// operations.
//
// If any errors are returned by method
// TextFieldSpecFiller.Read() which are NOT equal to io.EOF, call
// this method, TextFieldSpecFiller.ReaderInitialize(), to reset
// and prepare the internal reader for future read operations.
//
// This method fulfills the requirements of the
// ITextFieldSpecification interface.
func (txtFillerField *TextFieldSpecFiller) ReaderInitialize() {

	if txtFillerField.lock == nil {
		txtFillerField.lock = new(sync.Mutex)
	}

	txtFillerField.lock.Lock()

	defer txtFillerField.lock.Unlock()

	txtFillerField.textLineReader = nil

	return
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
//	Example 1:
//	 Filler Characters = "-"
//	 Filler Characters Repeat Count = 3
//	 Formatted Text = "---"
//
//	Example 2:
//	 Filler Characters = "-*"
//	 Filler Characters Repeat Count = 3
//	 Formatted Text = "-*-*-*"
//
// The 'fillerCharsRepeatCount' integer value is the number times
// that 'fillerCharacters' is repeated in order to construct the
// Filler Text Field. Be advised that Filler Text Fields requires a
// 'fillerCharsRepeatCount' value greater than zero.
// 'fillerCharsRepeatCount' values less than or equal to zero
// will trigger an error.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	 fillerCharsRepeatCount     int
//	    - Controls the number of times 'fillerCharacters' is
//	      repeated when constructing the final Text Filler Field
//	      returned by this method. The actual length of the string
//	      which will populated the completed Text Filler Field is
//	      equal to the length of 'fillerCharacters' times the value
//	      of 'fillerCharsRepeatCount'.
//
//	       Text Field Filler Length =
//	         Length of fillerCharacters X fillerCharsRepeatCount
//
//	         Example #1: fillerCharacters = "-*"
//	                     fillerRepeatCount = 3
//	                     Final Text Filler Field = "-*-*-*"
//
//	         Example #2: fillerCharacters = "-"
//	                     fillerRepeatCount = 3
//	                     Final Text Filler Field = "---"
//
//	      If 'fillerCharsRepeatCount' has a value less than one (1) or
//	      greater than one-million (1,000,000), an error will be
//	      returned.
//
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
//
//		     This empty interface must be convertible to one of the
//		     following types:
//
//		     1. nil - A nil value is valid and generates an empty
//		        collection of error prefix and error context
//		        information.
//
//		     2. string - A string containing error prefix information.
//
//		     3. []string A one-dimensional slice of strings containing
//		        error prefix information
//
//		     4. [][2]string A two-dimensional slice of strings
//		        containing error prefix and error context information.
//
//		     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		        from this object will be copied for use in error and
//		        informational messages.
//
//		     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		        Information from this object will be copied for use in
//		        error and informational messages.
//
//		     7. IBasicErrorPrefix - An interface to a method generating
//		        a two-dimensional slice of strings containing error
//		        prefix and error context information.
//
//		     If parameter 'errorPrefix' is NOT convertible to one of
//		     the valid types listed above, it will be considered
//		     invalid and trigger the return of an error.
//
//		     Types ErrPrefixDto and IBasicErrorPrefix are included in
//		     the 'errpref' software package,
//		     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If this method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
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
			ePrefix.XCpy(
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
// # IMPORTANT
//
// This method will overwrite and delete the existing data values
// for the current TextFieldSpecFiller instance (txtFillerField).
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	 fillerCharacters           string
//	    - A string containing the text characters which will be
//	      included in the Text Filler Field. The final Text Filler
//	      Field will be constructed from the filler characters
//	      repeated one or more times as specified by the
//	      'fillerCharsRepeatCount' parameter.
//
//	      The Text Field Filler final formatted text is equal to:
//	         fillerCharacter X fillerCharsRepeatCount
//	         Example: fillerCharacters = "-*"
//	                  fillerRepeatCount = 3
//	                  Final Text Filler Field = "-*-*-*"
//
//	      If 'fillerCharacters' is submitted as an empty or zero
//	      length string, this method will return an error.
//
//
//	 fillerCharsRepeatCount     int
//	    - Controls the number of times 'fillerCharacters' is
//	      repeated when constructing the final Text Filler Field
//	      returned by this method. The actual length of the string
//	      which will populated the completed Text Filler Field is
//	      equal to the length of 'fillerCharacters' times the value
//	      of 'fillerCharsRepeatCount'.
//
//	       Text Field Filler Length =
//	         Length of fillerCharacters X fillerCharsRepeatCount
//
//	         Example #1: fillerCharacters = "-*"
//	                     fillerRepeatCount = 3
//	                     Final Text Filler Field = "-*-*-*"
//
//	         Example #2: fillerCharacters = "-"
//	                     fillerRepeatCount = 3
//	                     Final Text Filler Field = "---"
//
//	      If 'fillerCharsRepeatCount' has a value less than one (1)
//	      or greater than one-million (1,000,000), an error will be
//	      returned.
//
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
//
//		     This empty interface must be convertible to one of the
//		     following types:
//
//		     1. nil - A nil value is valid and generates an empty
//		        collection of error prefix and error context
//		        information.
//
//		     2. string - A string containing error prefix information.
//
//		     3. []string A one-dimensional slice of strings containing
//		        error prefix information
//
//		     4. [][2]string A two-dimensional slice of strings
//		        containing error prefix and error context information.
//
//		     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		        from this object will be copied for use in error and
//		        informational messages.
//
//		     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		        Information from this object will be copied for use in
//		        error and informational messages.
//
//		     7. IBasicErrorPrefix - An interface to a method generating
//		        a two-dimensional slice of strings containing error
//		        prefix and error context information.
//
//		     If parameter 'errorPrefix' is NOT convertible to one of
//		     the valid types listed above, it will be considered
//		     invalid and trigger the return of an error.
//
//		     Types ErrPrefixDto and IBasicErrorPrefix are included in
//		     the 'errpref' software package,
//		     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If this method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
//
// ----------------------------------------------------------------
//
// Example Usage
//
//	Example 1:
//	 Filler Characters Array = "-"
//	 Filler Characters Repeat Count = 3
//	 Formatted Text = "---"
//
//	Example 2:
//	 Filler Characters Array = "-*"
//	 Filler Characters Repeat Count = 3
//	 Formatted Text = "-*-*-*"
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
// # IMPORTANT
//
// This method will overwrite and delete the existing data values
// for the current TextFieldSpecFiller instance (txtFillerField).
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	 fillerCharacter            rune
//	    - A rune containing the text character which will be
//	      included in the Text Filler Field. The final Text Filler
//	      Field will be constructed from ths filler character
//	      repeated one or more times as specified by the
//	      'fillerCharsRepeatCount' parameter.
//
//	      The Text Field Filler final formatted text is equal to:
//	          fillerCharacter X fillerCharsRepeatCount
//	          Example: fillerCharacter = '-'
//	                   fillerRepeatCount = 3
//	                   Final Text Filler Field = "---"
//
//	      If 'fillerCharacter' is submitted with a zero value,
//	      this method will return an error.
//
//
//	 fillerCharsRepeatCount     int
//	    - Controls the number of times 'fillerCharacter' is
//	      repeated when constructing the final Text Filler Field
//	      returned by this method. The actual length of the string
//	      which will populated the completed Text Filler Field is
//	      equal to the length of 'fillerCharacter' (1) times the
//	      value of 'fillerCharsRepeatCount'.
//
//	        Text Field Filler Length =
//	          Length of fillerCharacter (1) X fillerCharsRepeatCount
//	          Example: fillerCharacter = '-'
//	                   fillerRepeatCount = 3
//	                   Final Text Filler Field = "---"
//
//	      If 'fillerCharsRepeatCount' has a value less than one (1) or
//	      greater than one-million (1,000,000), an error will be
//	      returned.
//
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
//
//		     This empty interface must be convertible to one of the
//		     following types:
//
//		     1. nil - A nil value is valid and generates an empty
//		        collection of error prefix and error context
//		        information.
//
//		     2. string - A string containing error prefix information.
//
//		     3. []string A one-dimensional slice of strings containing
//		        error prefix information
//
//		     4. [][2]string A two-dimensional slice of strings
//		        containing error prefix and error context information.
//
//		     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		        from this object will be copied for use in error and
//		        informational messages.
//
//		     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		        Information from this object will be copied for use in
//		        error and informational messages.
//
//		     7. IBasicErrorPrefix - An interface to a method generating
//		        a two-dimensional slice of strings containing error
//		        prefix and error context information.
//
//		     If parameter 'errorPrefix' is NOT convertible to one of
//		     the valid types listed above, it will be considered
//		     invalid and trigger the return of an error.
//
//		     Types ErrPrefixDto and IBasicErrorPrefix are included in
//		     the 'errpref' software package,
//		     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If this method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
//
// ----------------------------------------------------------------
//
// Example Usage
//
//	Example 1:
//	 Filler Characters Array = '-'
//	 Filler Characters Repeat Count = 3
//	 Formatted Text = "---"
//
//	Example 2:
//	 Filler Characters Array = '*'
//	 Filler Characters Repeat Count = 3
//	 Formatted Text = "***"
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
// # IMPORTANT
//
// This method will overwrite and delete the existing data values
// for the current TextFieldSpecFiller instance (txtFillerField).
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	 fillerCharacters           []rune
//	    - A rune array containing the text characters which will be
//	      included in the Text Filler Field. The final Text Filler
//	      Field will be constructed from ths filler characters
//	      repeated one or more times as specified by the
//	      'fillerCharsRepeatCount' parameter.
//
//	      The Text Field Filler final formatted text is equal to:
//	         fillerCharacters X fillerCharsRepeatCount
//	         Example: fillerCharacters = []rune{'-','*'}
//	                  fillerRepeatCount = 3
//	                  Final Text Filler Field = "-*-*-*"
//
//	      If 'fillerCharacters' is submitted with a zero length rune
//	      array, this method will return an error.
//
//
//	 fillerCharsRepeatCount     int
//	    - Controls the number of times 'fillerCharacters' is
//	      repeated when constructing the final Text Filler Field
//	      returned by this method. The actual length of the string
//	      which will populated the completed Text Filler Field is
//	      equal to the length of 'fillerCharacters' times the value
//	      of 'fillerCharsRepeatCount'.
//
//	       Text Field Filler Length =
//	         Length of fillerCharacters X fillerCharsRepeatCount
//
//	         Example #1: fillerCharacters = []rune{'-','*'}
//	                     fillerRepeatCount = 3
//	                     Final Text Filler Field = "-*-*-*"
//
//	         Example #2: fillerCharacters = []rune{'-'}
//	                     fillerRepeatCount = 3
//	                     Final Text Filler Field = "---"
//
//	      If 'fillerCharsRepeatCount' has a value less than one (1) or
//	      greater than one-million (1,000,000), an error will be
//	      returned.
//
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
//
//		     This empty interface must be convertible to one of the
//		     following types:
//
//		     1. nil - A nil value is valid and generates an empty
//		        collection of error prefix and error context
//		        information.
//
//		     2. string - A string containing error prefix information.
//
//		     3. []string A one-dimensional slice of strings containing
//		        error prefix information
//
//		     4. [][2]string A two-dimensional slice of strings
//		        containing error prefix and error context information.
//
//		     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		        from this object will be copied for use in error and
//		        informational messages.
//
//		     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		        Information from this object will be copied for use in
//		        error and informational messages.
//
//		     7. IBasicErrorPrefix - An interface to a method generating
//		        a two-dimensional slice of strings containing error
//		        prefix and error context information.
//
//		     If parameter 'errorPrefix' is NOT convertible to one of
//		     the valid types listed above, it will be considered
//		     invalid and trigger the return of an error.
//
//		     Types ErrPrefixDto and IBasicErrorPrefix are included in
//		     the 'errpref' software package,
//		     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If this method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
//
// ----------------------------------------------------------------
//
// Example Usage
//
//	Example 1:
//	 Filler Characters Array = []rune{'-'}
//	 Filler Characters Repeat Count = 3
//	 Formatted Text = "---"
//
//	Example 2:
//	 Filler Characters Array = []rune{'-','*'}
//	 Filler Characters Repeat Count = 3
//	 Formatted Text = "-*-*-*"
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
//	Filler Characters Array Length  X
//	             Filler Characters Repeat Count
//
// This method is similar to method
// TextFieldSpecFiller.GetFormattedText() with the exception that
// this method does NOT return an error.
//
// This method fulfills the requirements of the
// ITextFieldSpecification interface.
//
// This method also fulfills the requirements of the 'Stringer'
// interface defined in the Golang package 'fmt'. Reference:
//
//	https://pkg.go.dev/fmt#Stringer
//
// Methods which return formatted text are listed as follows:
//
//	TextFieldSpecFiller.String()
//	TextFieldSpecFiller.GetFormattedText()
//	TextFieldSpecFiller.TextBuilder()
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//	Example 1:
//	 Filler Characters Array = []rune{'-'}
//	 Filler Characters Repeat Count = 3
//	 Formatted Text = "---"
//
//	Example 2:
//	 Filler Characters Array = []rune{'-','*'}
//	 Filler Characters Repeat Count = 3
//	 Formatted Text = "-*-*-*"
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

// TextBuilder - Configures the line of text produced by this
// instance of TextFieldSpecFiller, and writes it to an instance
// of strings.Builder.
//
// This method fulfills the requirements of the
// ITextFieldSpecification interface.
//
// Methods which return formatted text are listed as follows:
//
//	TextFieldSpecFiller.String()
//	TextFieldSpecFiller.GetFormattedText()
//	TextFieldSpecFiller.TextBuilder()
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	 strBuilder                 *strings.Builder
//	    - A pointer to an instance of *strings.Builder. The
//	      formatted text characters produced by this method will be
//	      written to this instance of strings.Builder.
//
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
//
//		     This empty interface must be convertible to one of the
//		     following types:
//
//		     1. nil - A nil value is valid and generates an empty
//		        collection of error prefix and error context
//		        information.
//
//		     2. string - A string containing error prefix information.
//
//		     3. []string A one-dimensional slice of strings containing
//		        error prefix information
//
//		     4. [][2]string A two-dimensional slice of strings
//		        containing error prefix and error context information.
//
//		     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		        from this object will be copied for use in error and
//		        informational messages.
//
//		     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		        Information from this object will be copied for use in
//		        error and informational messages.
//
//		     7. IBasicErrorPrefix - An interface to a method generating
//		        a two-dimensional slice of strings containing error
//		        prefix and error context information.
//
//		     If parameter 'errorPrefix' is NOT convertible to one of
//		     the valid types listed above, it will be considered
//		     invalid and trigger the return of an error.
//
//		     Types ErrPrefixDto and IBasicErrorPrefix are included in
//		     the 'errpref' software package,
//		     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If the method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (txtFillerField *TextFieldSpecFiller) TextBuilder(
	strBuilder *strings.Builder,
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
		"TextFieldSpecFiller.TextBuilder()",
		"")

	if err != nil {
		return err
	}

	if strBuilder == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'strBuilder' is invalid!\n"+
			"'strBuilder' is a nil pointer.\n",
			ePrefix.String())

		return err
	}

	var formattedTxtStr string

	formattedTxtStr,
		err = textFieldSpecFillerMolecule{}.ptr().
		getFormattedText(
			txtFillerField,
			ePrefix.XCpy(
				"txtFillerField->formattedTxtStr"))

	if err != nil {
		return err
	}

	lenFormattedText := len(formattedTxtStr)

	netCapacityStrBuilder :=
		strBuilder.Cap() -
			strBuilder.Len()

	requiredCapacity :=
		lenFormattedText - netCapacityStrBuilder

	if requiredCapacity > 0 {

		strBuilder.Grow(requiredCapacity + 16)
	}

	var err2 error

	_,
		err2 = strBuilder.WriteString(formattedTxtStr)

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"Error returned by sBuilder.WriteString(formattedTxtStr)\n"+
			"%v\n",
			ePrefix.String(),
			err2.Error())
	}

	return err
}

// TextFieldName - returns a string specifying the name
// of the Text Field specification. This method fulfills
// the requirements of the ITextFieldSpecification interface.
func (txtFillerField *TextFieldSpecFiller) TextFieldName() string {

	if txtFillerField.lock == nil {
		txtFillerField.lock = new(sync.Mutex)
	}

	txtFillerField.lock.Lock()

	defer txtFillerField.lock.Unlock()

	return "Filler"
}

// TextTypeName - returns a string specifying the type
// of Text Field specification. This method fulfills
// the requirements of the ITextSpecification interface.
func (txtFillerField *TextFieldSpecFiller) TextTypeName() string {

	if txtFillerField.lock == nil {
		txtFillerField.lock = new(sync.Mutex)
	}

	txtFillerField.lock.Lock()

	defer txtFillerField.lock.Unlock()

	return "TextFieldSpecFiller"
}
