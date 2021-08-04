package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type textFieldSpecLabelMolecule struct {
	lock *sync.Mutex
}

// copyIn - Copies all data from input parameter
// 'incomingTxtFieldLabel' to input parameter
// 'targetTxtFieldLabel'.
//
// IMPORTANT
// -----------------------------------------------------------------
// Be advised that the data fields in 'targetTxtFieldLabel' will be
// overwritten.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  targetTxtFieldLabel        *TextFieldSpecLabel
//     - A pointer to a TextFieldSpecLabel instance. All of the
//       member variable data fields in this object will be
//       replaced by data values extracted from input parameter
//       'incomingTxtFieldLabel'.
//
//       'targetTxtFieldLabel' is the target of this copy
//       operation.
//
//
//  incomingTxtFieldLabel      *TextFieldSpecLabel
//     - A pointer to a another TextFieldSpecLabel instance. All
//       of the member variable data values from this object will
//       be copied to corresponding member variables in
//       'targetTxtFieldLabel'.
//
//       'incomingTxtFieldLabel' is the source for this copy
//       operation.
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
func (txtFieldLabelMolecule *textFieldSpecLabelMolecule) copyIn(
	targetTxtFieldLabel *TextFieldSpecLabel,
	incomingTxtFieldLabel *TextFieldSpecLabel,
	errPrefDto *ePref.ErrPrefixDto) error {

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
		"textFieldSpecLabelMolecule.copyIn()",
		"")

	if err != nil {
		return err
	}

	if targetTxtFieldLabel == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetTxtFieldLabel' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if incomingTxtFieldLabel == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'incomingTxtFieldLabel' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return err
	}

	_,
		err = textFieldSpecLabelAtom{}.ptr().
		isValidTextFieldLabel(
			incomingTxtFieldLabel,
			ePrefix.XCtx("incomingTxtFieldLabel"))

	if err != nil {
		return err
	}

	// Set zero length arrays to nil == true
	err = strMechPreon{}.ptr().copyRuneArrays(
		&targetTxtFieldLabel.textLabel,
		&incomingTxtFieldLabel.textLabel,
		true,
		ePrefix.XCtx("targetTxtFieldLabel.textLabel=Target "+
			"<-incomingTxtFieldLabel.textLabel=Source"))

	if err != nil {
		return err
	}

	targetTxtFieldLabel.fieldLen =
		incomingTxtFieldLabel.fieldLen

	targetTxtFieldLabel.textJustification =
		incomingTxtFieldLabel.textJustification

	return nil
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
//       If the member variable data values encapsulated by this
//       'txtFieldLabel' are found to be invalid, this method will
//       return an error
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

	_,
		err = textFieldSpecLabelAtom{}.ptr().
		isValidTextFieldLabel(
			txtFieldLabel,
			ePrefix.XCtx("txtFieldLabel"))

	if err != nil {
		return TextFieldSpecLabel{}, err
	}

	newTxtFieldLabel := TextFieldSpecLabel{}

	lenTxtRunes := len(txtFieldLabel.textLabel)

	if lenTxtRunes > 0 {

		newTxtFieldLabel.textLabel =
			make([]rune, lenTxtRunes)

		copy(newTxtFieldLabel.textLabel,
			txtFieldLabel.textLabel)
	}

	newTxtFieldLabel.fieldLen =
		txtFieldLabel.fieldLen

	newTxtFieldLabel.textJustification =
		txtFieldLabel.textJustification

	newTxtFieldLabel.lock = new(sync.Mutex)

	return newTxtFieldLabel, nil
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
func (txtFieldLabelMolecule *textFieldSpecLabelMolecule) empty(
	txtFieldLabel *TextFieldSpecLabel) {

	if txtFieldLabelMolecule.lock == nil {
		txtFieldLabelMolecule.lock = new(sync.Mutex)
	}

	txtFieldLabelMolecule.lock.Lock()

	defer txtFieldLabelMolecule.lock.Unlock()

	if txtFieldLabel == nil {
		return
	}

	txtFieldLabel.textLabel = nil

	txtFieldLabel.fieldLen = 0

	txtFieldLabel.textJustification = TextJustify(0).None()

	return
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

	sMechPreon := strMechPreon{}

	if !sMechPreon.equalRuneArrays(
		txtLabel.textLabel,
		incomingTxtLabel.textLabel) {
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
