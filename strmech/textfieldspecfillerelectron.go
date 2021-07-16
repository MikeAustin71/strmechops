package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type textFieldSpecFillerElectron struct {
	lock *sync.Mutex
}

// isFillerCharsRepeatCountValid - Receives the
// 'fillerCharsRepeatCount' member variable of a
// TextFieldSpecFiller object and subjects it to a diagnostic
// review in order to determine if the filler characters are valid.
//
// If input parameter 'fillerCharsRepeatCount' is judged invalid,
// this method will return an error along with an appropriate error
// message.
//
// If input parameter 'fillerCharsRepeatCount' is valid this method
// will return a 'nil' value.
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
//     - If 'fillerCharsRepeatCount' is found to be invalid, this
//       method will return an error along with an appropriate
//       error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' (error prefix) will be inserted or
//       prefixed at the beginning of the error message.
//
//
func (txtFieldFillerElectron *textFieldSpecFillerElectron) isFillerCharsRepeatCountValid(
	fillerCharsRepeatCount int,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if txtFieldFillerElectron.lock == nil {
		txtFieldFillerElectron.lock = new(sync.Mutex)
	}

	txtFieldFillerElectron.lock.Lock()

	defer txtFieldFillerElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"txtFieldFillerElectron.isFillerCharsRepeatCountValid()",
		"")

	if err != nil {
		return err
	}

	if fillerCharsRepeatCount < 1 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fillerCharsRepeatCount' is less than one (1)!\n"+
			"'fillerCharsRepeatCount' controls the number of repetitions of 'fillerCharacters'\n"+
			"in the Filler Text Field.\n",
			ePrefix.String())
		return err
	}

	if fillerCharsRepeatCount > 1000000 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fillerCharsRepeatCount' is greater than one-million (+1,000,000)!\n"+
			"'fillerCharsRepeatCount' controls the number of repetitions of 'fillerCharacters'\n"+
			"in the Filler Text Field.\n",
			ePrefix.String())
		return err
	}

	return nil
}

// isFillerCharsValid - Receives the 'fillerCharacters' member
// variable of a TextFieldSpecFiller object and subjects it to a
// diagnostic review in order to determine if the filler characters
// are valid.
//
// If input parameter 'fillerChars' is judged invalid, this method
// will return an error along with an appropriate error message.
//
// If input parameter 'fillerChars' is valid this method
// will return a 'nil' value.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  fillerChars                []rune
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
//     - If 'fillerChars' is found to be invalid, this method will
//       return an error along with an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' (error prefix) will be inserted or
//       prefixed at the beginning of the error message.
//
func (txtFieldFillerElectron *textFieldSpecFillerElectron) isFillerCharsValid(
	fillerChars []rune,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if txtFieldFillerElectron.lock == nil {
		txtFieldFillerElectron.lock = new(sync.Mutex)
	}

	txtFieldFillerElectron.lock.Lock()

	defer txtFieldFillerElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"txtFieldFillerElectron.isFillerCharsValid()",
		"")

	if err != nil {
		return err
	}

	if fillerChars == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fillerChars' is 'nil' and "+
			"has a zero length!\n",
			ePrefix.String())

		return err
	}

	if len(fillerChars) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fillerChars' is has a "+
			"zero length!\n",
			ePrefix.String())

		return err
	}

	if len(fillerChars) == 1 &&
		fillerChars[0] == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fillerChars' has a\n"+
			"single character with a zero value!\n",
			ePrefix.String())

		return err
	}

	return nil
}

// ptr - Returns a pointer to a new instance of
// textFieldSpecFillerElectron.
//
func (txtFieldFillerElectron textFieldSpecFillerElectron) ptr() *textFieldSpecFillerElectron {

	if txtFieldFillerElectron.lock == nil {
		txtFieldFillerElectron.lock = new(sync.Mutex)
	}

	txtFieldFillerElectron.lock.Lock()

	defer txtFieldFillerElectron.lock.Unlock()

	return &textFieldSpecFillerElectron{
		lock: new(sync.Mutex),
	}
}
