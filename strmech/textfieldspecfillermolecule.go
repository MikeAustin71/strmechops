package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type textFieldSpecFillerMolecule struct {
	lock *sync.Mutex
}

// copyIn - Copies all data from input parameter
// 'incomingTxtFiller' to input parameter 'incomingTxtFiller'.
//
// IMPORTANT
// ----------------------------------------------------------------
// Be advised that the data fields in 'targetTxtFiller' will be
// overwritten.
//
func (txtFieldFillerMolecule *textFieldSpecFillerMolecule) copyIn(
	targetTxtFiller *TextFieldSpecFiller,
	incomingTxtFiller *TextFieldSpecFiller,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtFieldFillerMolecule.lock == nil {
		txtFieldFillerMolecule.lock = new(sync.Mutex)
	}

	txtFieldFillerMolecule.lock.Lock()

	defer txtFieldFillerMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"txtFieldFillerMolecule.copyIn()",
		"")

	if err != nil {
		return err
	}

	if targetTxtFiller == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetTxtFiller' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if incomingTxtFiller == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'incomingTxtFieldLabel' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return err
	}

	_,
		err =
		textFieldSpecFillerAtom{}.ptr().
			isValidTextFieldSpecFiller(
				incomingTxtFiller,
				ePrefix.XCtx("incomingTxtFiller validation - "))

	if err != nil {
		return err
	}

	// Setting zero length array to nil.
	err = strMechPreon{}.ptr().copyRuneArrays(
		targetTxtFiller.fillerCharacters,
		incomingTxtFiller.fillerCharacters,
		true,
		ePrefix.XCtx("targetTxtFiller.fillerCharacters=Target "+
			"<-incomingTxtFiller.fillerCharacters=Source"))

	if err != nil {
		return err
	}

	targetTxtFiller.fillerCharsRepeatCount =
		incomingTxtFiller.fillerCharsRepeatCount

	return nil
}

// copyOut - Returns a deep copy of the input parameter
// 'txtFieldFiller'
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  txtFieldFiller      *TextFieldSpecFiller
//     - A pointer to an instance of TextFieldSpecFiller. A deep
//       copy of the internal member variables will be created
//       and returned in a new instance of TextFieldSpecFiller.
//
//
//  errPrefDto          *ePref.ErrPrefixDto
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
//  TextFieldSpecFiller
//     - If this method completes successfully, a deep copy of
//       input parameter 'txtFieldFiller' will be created and
//       returned in a new instance of TextFieldSpecFiller.
//
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
func (txtFieldFillerMolecule *textFieldSpecFillerMolecule) copyOut(
	txtFieldFiller *TextFieldSpecFiller,
	errPrefDto *ePref.ErrPrefixDto) (
	TextFieldSpecFiller, error) {

	if txtFieldFillerMolecule.lock == nil {
		txtFieldFillerMolecule.lock = new(sync.Mutex)
	}

	txtFieldFillerMolecule.lock.Lock()

	defer txtFieldFillerMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"txtFieldFillerMolecule.copyOut()",
		"")

	if err != nil {
		return TextFieldSpecFiller{}, err
	}

	if txtFieldFiller == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtFieldFiller' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return TextFieldSpecFiller{}, err
	}

	_,
		err =
		textFieldSpecFillerAtom{}.ptr().
			isValidTextFieldSpecFiller(
				txtFieldFiller,
				ePrefix.XCtx("txtFieldFiller validation - "))

	if err != nil {
		return TextFieldSpecFiller{}, err
	}

	newTxtFieldFiller := TextFieldSpecFiller{}

	// Set zero length array to nil = true
	err = strMechPreon{}.ptr().copyRuneArrays(
		newTxtFieldFiller.fillerCharacters,
		txtFieldFiller.fillerCharacters,
		true,
		ePrefix.XCtx(
			"newTxtFieldFiller.fillerCharacters=Target "+
				"<-txtFieldFiller.fillerCharacters=Source"))

	if err != nil {
		return TextFieldSpecFiller{}, err
	}

	newTxtFieldFiller.fillerCharsRepeatCount =
		txtFieldFiller.fillerCharsRepeatCount

	newTxtFieldFiller.lock = new(sync.Mutex)

	return newTxtFieldFiller, nil
}

// empty - Receives a pointer to an instance of TextFieldSpecLabel
// and proceeds to set all of the internal member variables to
// their uninitialized or zero states.
//
// IMPORTANT
// ----------------------------------------------------------------
// The values of all member variables contained in input parameter
// 'txtFieldLabel' will be overwritten and deleted.
//
func (txtFieldFillerMolecule *textFieldSpecFillerMolecule) empty(
	txtFieldFiller *TextFieldSpecFiller) {

	if txtFieldFillerMolecule.lock == nil {
		txtFieldFillerMolecule.lock = new(sync.Mutex)
	}

	txtFieldFillerMolecule.lock.Lock()

	defer txtFieldFillerMolecule.lock.Unlock()

	if txtFieldFiller == nil {
		return
	}

	txtFieldFiller.fillerCharacters = nil

	txtFieldFiller.fillerCharsRepeatCount = 0

	return
}

// equal - Receives a pointer to two instances of
// TextFieldSpecFiller and proceeds to compare their member
// variables in order to determine if they are equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables for both instances are equal
// in all respects, this flag is set to 'true'. Otherwise, this
// method returns 'false'.
//
func (txtFieldFillerMolecule *textFieldSpecFillerMolecule) equal(
	txtFieldFiller *TextFieldSpecFiller,
	incomingTxtFieldFiller *TextFieldSpecFiller) bool {

	if txtFieldFillerMolecule.lock == nil {
		txtFieldFillerMolecule.lock = new(sync.Mutex)
	}

	txtFieldFillerMolecule.lock.Lock()

	defer txtFieldFillerMolecule.lock.Unlock()

	if txtFieldFiller == nil ||
		incomingTxtFieldFiller == nil {
		return false
	}

	if txtFieldFiller.fillerCharsRepeatCount !=
		incomingTxtFieldFiller.fillerCharsRepeatCount {
		return false
	}

	lenInTxtFiller := len(incomingTxtFieldFiller.fillerCharacters)

	if lenInTxtFiller != len(txtFieldFiller.fillerCharacters) {
		return false
	}

	if lenInTxtFiller > 0 {
		for i := 0; i < lenInTxtFiller; i++ {
			if incomingTxtFieldFiller.fillerCharacters[i] !=
				txtFieldFiller.fillerCharacters[i] {
				return false
			}
		}
	}

	return true
}

// newEmpty - Returns a new unpopulated instance of
// TextFieldSpecFiller. All of the member variables contained in
// this new instance are set to their uninitialized or zero values.
//
// Be advised that setting member variables to their zero values
// means that the TextFieldSpecFiller is invalid.
//
func (txtFieldFillerMolecule *textFieldSpecFillerMolecule) newEmpty() TextFieldSpecFiller {

	if txtFieldFillerMolecule.lock == nil {
		txtFieldFillerMolecule.lock = new(sync.Mutex)
	}

	txtFieldFillerMolecule.lock.Lock()

	defer txtFieldFillerMolecule.lock.Unlock()

	newFillerField := TextFieldSpecFiller{}

	newFillerField.fillerCharacters = nil

	newFillerField.fillerCharsRepeatCount = 0

	newFillerField.lock = new(sync.Mutex)

	return newFillerField
}

// ptr - Returns a pointer to a new instance of
// textFieldSpecFillerMolecule.
//
func (txtFieldFillerMolecule textFieldSpecFillerMolecule) ptr() *textFieldSpecFillerMolecule {

	if txtFieldFillerMolecule.lock == nil {
		txtFieldFillerMolecule.lock = new(sync.Mutex)
	}

	txtFieldFillerMolecule.lock.Lock()

	defer txtFieldFillerMolecule.lock.Unlock()

	return &textFieldSpecFillerMolecule{
		lock: new(sync.Mutex),
	}
}
