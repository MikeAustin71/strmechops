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
//  errPrefDto                 *ErrPrefixDto
//     - This error prefix data transfer object encapsulates an
//       error prefix string which is included in all returned
//       error messages. Usually, it contains the names of the
//       calling functions or methods included in this chain of
//       code execution.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       Type ErrPrefixDto is included in the 'errpref' software
//       package, "github.com/MikeAustin71/errpref".
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  err                        error
//     - If 'fillerCharsRepeatCount' is found to be valid, this
//       return parameter will be set to 'nil'.
//
//     - If 'fillerCharsRepeatCount' is found to be invalid, this
//       method will return an error along with an appropriate
//       error message.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'errPrefDto'. The
//       'errPrefDto' text will be attached to the beginning of the
//       error message.
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
// If input parameter 'fillerChars' is valid, this method
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
//  errPrefDto                 *ErrPrefixDto
//     - This error prefix data transfer object encapsulates an
//       error prefix string which is included in all returned
//       error messages. Usually, it contains the names of the
//       calling functions or methods included in this chain of
//       code execution.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       Type ErrPrefixDto is included in the 'errpref' software
//       package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  lenFillerChars             int
//     - The length of input parameter 'fillerChars'.
//
//
//  err                        error
//     - If 'fillerChars' is found to be valid, this return
//       parameter will be set to 'nil'.
//
//     - If 'fillerChars' is found to be invalid, this method will
//       return an error along with an appropriate error message.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'errPrefDto'. The
//       'errPrefDto' text will be attached to the beginning of the
//       error message.
//
func (txtFieldFillerElectron *textFieldSpecFillerElectron) isFillerCharsValid(
	fillerChars []rune,
	errPrefDto *ePref.ErrPrefixDto) (
	lenFillerChars int,
	err error) {

	if txtFieldFillerElectron.lock == nil {
		txtFieldFillerElectron.lock = new(sync.Mutex)
	}

	txtFieldFillerElectron.lock.Lock()

	defer txtFieldFillerElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	lenFillerChars = 0

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"txtFieldFillerElectron.isFillerCharsValid()",
		"")

	if err != nil {
		return lenFillerChars, err
	}

	if fillerChars == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fillerChars' is 'nil' and "+
			"has a zero length!\n",
			ePrefix.String())

		return lenFillerChars, err
	}

	lenFillerChars = len(fillerChars)

	if lenFillerChars == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fillerChars' is has a "+
			"zero length!\n",
			ePrefix.String())

		return lenFillerChars, err
	}

	for i := 0; i < lenFillerChars; i++ {

		if fillerChars[i] == 0 {
			err = fmt.Errorf("%v\n"+
				"Error: At least one of the characters in Input parameter\n"+
				"'fillerChars' has a zero value!\n"+
				"fillerChars[%v] == 0\n",
				ePrefix.String(),
				i)

			lenFillerChars = -1

			return lenFillerChars, err
		}

	}

	return lenFillerChars, err
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
