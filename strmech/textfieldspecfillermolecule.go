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

	if incomingTxtFiller.fillerCharsRepeatCount < 1 {
		err = fmt.Errorf("%v\n"+
			"Error: 'incomingTxtFiller.fillerCharsRepeatCount' is "+
			"less than one (+1)!\n"+
			"incomingTxtFiller.fillerCharsRepeatCount='%v'\n",
			ePrefix.String(),
			incomingTxtFiller.fillerCharsRepeatCount)

		return err
	}

	lenInTxtRunes := len(incomingTxtFiller.fillerCharacters)

	if lenInTxtRunes == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: 'incomingTxtFiller.fillerCharacters' is a zero "+
			"length array!\n",
			ePrefix.String())

		return err
	}

	targetTxtFiller.fillerCharacters =
		make([]rune, lenInTxtRunes)

	copy(targetTxtFiller.fillerCharacters,
		incomingTxtFiller.fillerCharacters)

	targetTxtFiller.fillerCharsRepeatCount =
		incomingTxtFiller.fillerCharsRepeatCount

	return nil
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
