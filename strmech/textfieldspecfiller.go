package strmech

import "sync"

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
// Be advised that Filler Text Fields require a field length
// greater than zero. Field Lengths less than or equal to zero
// represent an error condition.
//
type TextFieldSpecFiller struct {
	fillerCharacter rune
	fieldLen        int
	lock            *sync.Mutex
}

// NewEmpty - Returns a new unpopulated instance of
// TextFieldSpecFiller. All of the member variables contained in
// this new instance are set to their uninitialized or zero values.
//
// Be advised that setting member variables to their zero values
// means that the TextFieldSpecFiller is invalid.
//
func (txtFillerField *TextFieldSpecFiller) NewEmpty() TextFieldSpecFiller {

	if txtFillerField.lock == nil {
		txtFillerField.lock = new(sync.Mutex)
	}

	txtFillerField.lock.Lock()

	defer txtFillerField.lock.Unlock()

	return textFieldSpecFillerMolecule{}.ptr().
		newEmpty()
}
