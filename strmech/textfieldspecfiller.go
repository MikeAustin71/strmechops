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

// NewConstructor - Creates and returns a new, fully populated
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
//
//  fillerCharsRepeatCount    int
//     - Controls the number of times 'fillerCharacters' is
//       repeated when constructing the final Text Filler Field
//       returned by this method. The actual length of the string
//       which will populated the completed Text Filler Field is
//       equal to the length of 'fillerCharacters' times the value
//       of 'fillerCharsRepeatCount'.
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

	if len(fillerCharacters) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fillerCharacters' is a zero length string!\n",
			ePrefix.String())
		return TextFieldSpecFiller{}, err
	}

	if fillerCharsCount < 1 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fillerCharsRepeatCount' is less than one (1)!\n"+
			"'fillerCharsRepeatCount' controls the number of repetitions of 'fillerCharacters'\n"+
			"in the Filler Text Field.\n",
			ePrefix.String())
		return TextFieldSpecFiller{}, err
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

	return newTxtFillerField, nil
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
//       If this parameter is submitted as a zero length string,
//       an error will be returned.
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
