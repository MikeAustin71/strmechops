package strmech

import "sync"

type textFieldSpecFillerMolecule struct {
	lock *sync.Mutex
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

	newFillerField.fillerCharacter = 0

	newFillerField.fieldLen = 0

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
