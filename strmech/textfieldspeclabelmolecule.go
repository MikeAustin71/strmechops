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
// 'sourceTxtFieldLabel' to input parameter
// 'destinationTxtFieldLabel'. Both instances are of type
// TextFieldSpecLabel.
//
// IMPORTANT
// -----------------------------------------------------------------
// Be advised that the data fields in 'destinationTxtFieldLabel'
// will be deleted and overwritten.
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//	destinationTxtFieldLabel   *TextFieldSpecLabel
//	   - A pointer to a TextFieldSpecLabel instance. All the
//	     member variable data fields in this object will be
//	     replaced by data values extracted from input parameter
//	     'sourceTxtFieldLabel'.
//
//	     'destinationTxtFieldLabel' is the target of this copy
//	     operation.
//
//
//	sourceTxtFieldLabel        *TextFieldSpecLabel
//	   - A pointer to another TextFieldSpecLabel instance. All
//	     the member variable data values from this object will
//	     be copied to corresponding member variables in
//	     'destinationTxtFieldLabel'.
//
//	     'sourceTxtFieldLabel' is the source for this copy
//	     operation.
//
//	     If 'sourceTxtFieldLabel' is determined to be invalid,
//	     an error will be returned.
//
//
//	errPrefDto          *ePref.ErrPrefixDto
//	   - This object encapsulates an error prefix string which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods listed
//	     as a function chain.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     Type ErrPrefixDto is included in the 'errpref' software
//	     package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If this method completes successfully, this returned error
//	     Type is set equal to 'nil'. If errors are encountered during
//	     processing, the returned error Type will encapsulate an error
//	     message.
//
//	     If an error message is returned, the text value for input
//	     parameter 'errPrefDto' (error prefix) will be prefixed or
//	     attached at the beginning of the error message.
func (txtFieldLabelMolecule *textFieldSpecLabelMolecule) copyTextFieldLabel(
	destinationTxtFieldLabel *TextFieldSpecLabel,
	sourceTxtFieldLabel *TextFieldSpecLabel,
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
		"textFieldSpecLabelMolecule."+
			"copyTextFieldLabel()",
		"")

	if err != nil {
		return err
	}

	if destinationTxtFieldLabel == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'destinationTxtFieldLabel' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if sourceTxtFieldLabel == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'sourceTxtFieldLabel' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return err
	}

	_,
		err = textFieldSpecLabelAtom{}.ptr().
		isValidTextFieldLabel(
			sourceTxtFieldLabel,
			ePrefix.XCpy("sourceTxtFieldLabel"))

	if err != nil {
		return err
	}

	destinationTxtFieldLabel.textLineReader = nil

	// Set zero length arrays to nil == true
	err = strMechPreon{}.ptr().copyRuneArrays(
		&destinationTxtFieldLabel.textLabel,
		&sourceTxtFieldLabel.textLabel,
		true,
		ePrefix.XCpy("destinationTxtFieldLabel.textLabel=Target "+
			"<-sourceTxtFieldLabel.textLabel=Source"))

	if err != nil {
		return err
	}

	destinationTxtFieldLabel.fieldLen =
		sourceTxtFieldLabel.fieldLen

	destinationTxtFieldLabel.textJustification =
		sourceTxtFieldLabel.textJustification

	return nil
}

// empty - Receives a pointer to an instance of TextFieldSpecLabel
// and proceeds to set all the internal member variables to
// their uninitialized or zero states.
//
// IMPORTANT
// ----------------------------------------------------------------
// The values of all member variables contained in input parameter
// 'txtFieldLabel' will be overwritten and replaced.
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

	txtFieldLabel.textLineReader = nil

	return
}

// equal - Receives pointers to two instances of
// TextFieldSpecLabel and proceeds to compare their member
// variables in order to determine if they are equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables for both instances are equal
// in all respects, this flag is set to 'true'. Otherwise, this
// method returns 'false'.
func (txtFieldLabelMolecule *textFieldSpecLabelMolecule) equal(
	txtLabelOne *TextFieldSpecLabel,
	txtLabelTwo *TextFieldSpecLabel) bool {

	if txtFieldLabelMolecule.lock == nil {
		txtFieldLabelMolecule.lock = new(sync.Mutex)
	}

	txtFieldLabelMolecule.lock.Lock()

	defer txtFieldLabelMolecule.lock.Unlock()

	if txtLabelOne == nil ||
		txtLabelTwo == nil {
		return false
	}

	sMechPreon := strMechPreon{}

	if !sMechPreon.equalRuneArrays(
		txtLabelOne.textLabel,
		txtLabelTwo.textLabel) {
		return false
	}

	if txtLabelOne.fieldLen !=
		txtLabelTwo.fieldLen {
		return false
	}

	if txtLabelOne.textJustification !=
		txtLabelTwo.textJustification {
		return false
	}

	return true
}
