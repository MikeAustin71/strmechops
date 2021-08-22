package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type textFieldSpecFillerNanobot struct {
	lock *sync.Mutex
}

// setTxtFieldSpecFiller - Receives a pointer to an instance of
// TextFieldSpecFiller and proceeds to reset the data values
// based on the input parameters.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// The pre-existing data fields for input parameter
// 'textFieldLabel' will be deleted and overwritten.
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
//  errPrefDto                 *ePref.ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods listed
//       as a function chain.
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
//  error
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (txtFieldFillerNanobot *textFieldSpecFillerNanobot) setTxtFieldSpecFiller(
	txtFieldFiller *TextFieldSpecFiller,
	fillerCharacters []rune,
	fillerCharsRepeatCount int,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtFieldFillerNanobot.lock == nil {
		txtFieldFillerNanobot.lock = new(sync.Mutex)
	}

	txtFieldFillerNanobot.lock.Lock()

	defer txtFieldFillerNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"txtFieldFillerMolecule."+
			"setTxtFieldSpecFiller()",
		"")

	if err != nil {
		return err
	}

	if txtFieldFiller == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtFieldFiller' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return err
	}

	txtFillerElectron := textFieldSpecFillerElectron{}

	var lenFillerChars int

	lenFillerChars,
		err = txtFillerElectron.isFillerCharsValid(
		fillerCharacters,
		ePrefix.XCtx(
			"Input parameter 'fillerCharacters' invalid!"))

	if err != nil {
		return err
	}

	err = txtFillerElectron.isFillerCharsRepeatCountValid(
		fillerCharsRepeatCount,
		ePrefix.XCtx(
			"Input parameter 'fillerCharsRepeatCount' invalid!"))

	if err != nil {
		return err
	}

	txtFieldFiller.fillerCharsRepeatCount =
		fillerCharsRepeatCount

	txtFieldFiller.fillerCharacters = nil

	txtFieldFiller.fillerCharacters =
		make([]rune, lenFillerChars)

	itemsCopied :=
		copy(
			txtFieldFiller.fillerCharacters,
			fillerCharacters)

	if itemsCopied != lenFillerChars {

		err = fmt.Errorf("%v\n"+
			"Error: 'fillerCharacters' copy operation failed!\n"+
			"Should have copied %v characters to txtFieldFiller.fillerCharacters.\n"+
			"Instead, %v characters were copied to txtFieldFiller.fillerCharacters.\n",
			ePrefix.XCtx(
				"fillerCharacters->txtFieldFiller.fillerCharacters"),
			lenFillerChars,
			itemsCopied)

	}

	return err
}

// ptr - Returns a pointer to a new instance of
// textFieldSpecFillerNanobot.
//
func (txtFieldFillerNanobot textFieldSpecFillerNanobot) ptr() *textFieldSpecFillerNanobot {

	if txtFieldFillerNanobot.lock == nil {
		txtFieldFillerNanobot.lock = new(sync.Mutex)
	}

	txtFieldFillerNanobot.lock.Lock()

	defer txtFieldFillerNanobot.lock.Unlock()

	return &textFieldSpecFillerNanobot{
		lock: new(sync.Mutex),
	}
}
