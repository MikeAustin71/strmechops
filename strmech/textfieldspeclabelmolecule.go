package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type textFieldSpecLabelMolecule struct {
	lock *sync.Mutex
}

// copyOut - Returns a deep copy of the input parameter
// 'txtFieldLabel'
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  txtFieldLabel       *TextFieldSpecLabel
//     - A pointer to an instance of TextFieldSpecLabel. A deep
//       copy of the internal member variables will be created
//       and returned in a new instance of TextFieldSpecLabel.
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
//  TextFieldSpecLabel
//     - If this method completes successfully, a deep copy of
//       input parameter 'txtFieldLabel' will be created and
//       returned in a new instance of TextFieldSpecLabel.
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
func (txtFieldLabelMolecule *textFieldSpecLabelMolecule) copyOut(
	txtFieldLabel *TextFieldSpecLabel,
	errPrefDto *ePref.ErrPrefixDto) (
	TextFieldSpecLabel, error) {

	if txtFieldLabelMolecule.lock == nil {
		txtFieldLabelMolecule.lock = new(sync.Mutex)
	}

	txtFieldLabelMolecule.lock.Lock()

	defer txtFieldLabelMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textFieldSpecLabelMolecule.copyOut()",
		"")

	if err != nil {
		return TextFieldSpecLabel{}, err
	}

	if txtFieldLabel == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtFieldLabel' is a nil pointer!\n",
			ePrefix.String())

		return TextFieldSpecLabel{}, err
	}

	newTxtFieldLabel := TextFieldSpecLabel{}

	newTxtFieldLabel.textLabel =
		txtFieldLabel.textLabel

	newTxtFieldLabel.fieldLen =
		txtFieldLabel.fieldLen

	newTxtFieldLabel.textJustification =
		txtFieldLabel.textJustification

	newTxtFieldLabel.lock = new(sync.Mutex)

	return newTxtFieldLabel, nil
}

// equal - Receives a pointer to two instances of
// TextFieldSpecLabel and proceeds to compare their member
// variables in order to determine if they are equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables for both instances are equal
// in all respects, this flag is set to 'true'. Otherwise, this
// method returns 'false'.
//
func (txtFieldLabelMolecule *textFieldSpecLabelMolecule) equal(
	txtLabel *TextFieldSpecLabel,
	incomingTxtLabel *TextFieldSpecLabel) bool {

	if txtFieldLabelMolecule.lock == nil {
		txtFieldLabelMolecule.lock = new(sync.Mutex)
	}

	txtFieldLabelMolecule.lock.Lock()

	defer txtFieldLabelMolecule.lock.Unlock()

	if txtLabel == nil {
		return false
	}

	if incomingTxtLabel == nil {
		return false
	}

	if txtLabel.textLabel !=
		incomingTxtLabel.textLabel {
		return false
	}

	if txtLabel.fieldLen !=
		incomingTxtLabel.fieldLen {
		return false
	}

	if txtLabel.textJustification !=
		incomingTxtLabel.textJustification {
		return false
	}

	return true
}

// ptr - Returns a pointer to a new instance of
// textFieldSpecLabelMolecule.
//
func (txtFieldLabelMolecule textFieldSpecLabelMolecule) ptr() *textFieldSpecLabelMolecule {

	if txtFieldLabelMolecule.lock == nil {
		txtFieldLabelMolecule.lock = new(sync.Mutex)
	}

	txtFieldLabelMolecule.lock.Lock()

	defer txtFieldLabelMolecule.lock.Unlock()

	return &textFieldSpecLabelMolecule{
		lock: new(sync.Mutex),
	}
}
