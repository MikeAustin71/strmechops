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
// The 'fillerCharsCount' integer value the number times that
// 'fillerCharacters' is repeated in order to construct the Filler
// Text Field. Be advised that Filler Text Fields requires a
// 'fillerCharsCount' value greater than zero. 'fillerCharsCount' values less
// than or equal to zero constitute an error condition.
//
type TextFieldSpecFiller struct {
	fillerCharacters []rune
	fillerCharsCount int
	lock             *sync.Mutex
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
//       'fillerCharsCount' parameter.
//
//
//  fillerCharsCount    int
//     - Controls the number of times 'fillerCharacters' is
//       repeated when constructing the final Text Filler Field
//       returned by this method. The actual length of the string
//       which will populated the completed Text Filler Field is
//       equal to the length of 'fillerCharacters' times the value
//       of 'fillerCharsCount'.
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
		"TextFieldSpecLabel.CopyIn()",
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
			"Error: Input parameter 'fillerCharsCount' is less than one (1)!\n"+
			"'fillerCharsCount' controls the number of repetitions of 'fillerCharacters'\n"+
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

	newTxtFillerField.fillerCharsCount =
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
