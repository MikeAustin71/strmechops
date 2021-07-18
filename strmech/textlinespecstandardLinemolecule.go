package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type textLineSpecStandardLineMolecule struct {
	lock *sync.Mutex
}

// copyIn - Copies all data from input parameter 'incomingStdLine'
// to input parameter 'targetStdLine'.
//
// Be advised that the data fields in 'targetStdLine' will be
// overwritten.
//
func (txtStdLineMolecule *textLineSpecStandardLineMolecule) copyIn(
	targetStdLine *TextLineSpecStandardLine,
	incomingStdLine *TextLineSpecStandardLine,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtStdLineMolecule.lock == nil {
		txtStdLineMolecule.lock = new(sync.Mutex)
	}

	txtStdLineMolecule.lock.Lock()

	defer txtStdLineMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecStandardLineMolecule.copyIn()",
		"")

	if err != nil {
		return err
	}

	if targetStdLine == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetStdLine' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if incomingStdLine == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'incomingStdLine' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if incomingStdLine.numOfStdLines < 0 {
		incomingStdLine.numOfStdLines = 0
	}

	lenInNewLineChars := len(incomingStdLine.newLineChars)

	if lenInNewLineChars == 0 {
		incomingStdLine.newLineChars = []rune{'\n'}
		lenInNewLineChars = 1
	}

	targetStdLine.newLineChars =
		make([]rune, lenInNewLineChars)

	copy(targetStdLine.newLineChars,
		incomingStdLine.newLineChars)

	targetStdLine.turnLineTerminatorOff =
		incomingStdLine.turnLineTerminatorOff

	targetStdLine.numOfStdLines =
		incomingStdLine.numOfStdLines

	lenIncomingTxtFields := len(incomingStdLine.textFields)

	if lenIncomingTxtFields == 0 {
		targetStdLine.textFields = nil
		return nil
	}

	targetStdLine.textFields =
		make([]ITextFieldSpecification, lenIncomingTxtFields)

	var tempITextField ITextFieldSpecification

	for i := 0; i < lenIncomingTxtFields; i++ {

		tempITextField,
			err = incomingStdLine.textFields[i].CopyOutITextField(
			ePrefix.XCtx(
				fmt.Sprintf("i='%v'", i)))

		if err != nil {
			return err
		}

		targetStdLine.textFields[i] = tempITextField
	}

	return nil
}

// copyOut - Returns a deep copy of the input parameter
// 'txtStdLine'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  txtStdLine          *TextLineSpecStandardLine
//     - A pointer to an instance of TextLineSpecStandardLine. A
//       deep copy of the internal member variables will be created
//       and returned in a new instance of TextLineSpecStandardLine.
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
//  TextLineSpecStandardLine
//     - If this method completes successfully, a deep copy of
//       input parameter 'txtStdLine' will be created and returned
//       in a new instance of TextLineSpecStandardLine.
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
func (txtStdLineMolecule *textLineSpecStandardLineMolecule) copyOut(
	txtStdLine *TextLineSpecStandardLine,
	errPrefDto *ePref.ErrPrefixDto) (
	TextLineSpecStandardLine, error) {

	if txtStdLineMolecule.lock == nil {
		txtStdLineMolecule.lock = new(sync.Mutex)
	}

	txtStdLineMolecule.lock.Lock()

	defer txtStdLineMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecStandardLineMolecule.copyOut()",
		"")

	if err != nil {
		return TextLineSpecStandardLine{}, err
	}

	if txtStdLine == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtStdLine' is a nil pointer!\n",
			ePrefix.String())

		return TextLineSpecStandardLine{}, err
	}

	newStdLine := TextLineSpecStandardLine{}

	newStdLine.lock = new(sync.Mutex)

	lenInNewLineChars := len(txtStdLine.newLineChars)

	if lenInNewLineChars == 0 {
		txtStdLine.newLineChars = []rune{'\n'}
		lenInNewLineChars = 1
	}

	newStdLine.newLineChars =
		make([]rune, lenInNewLineChars)

	copy(newStdLine.newLineChars,
		txtStdLine.newLineChars)

	newStdLine.turnLineTerminatorOff =
		txtStdLine.turnLineTerminatorOff

	lenTxtFields := len(txtStdLine.textFields)

	if lenTxtFields > 0 {

		newStdLine.textFields = make([]ITextFieldSpecification,
			lenTxtFields)

		var tempITextField ITextFieldSpecification

		for i := 0; i < lenTxtFields; i++ {

			tempITextField,
				err = txtStdLine.textFields[i].CopyOutITextField(
				ePrefix.XCtx(
					fmt.Sprintf("i='%v'", i)))

			if err != nil {
				return TextLineSpecStandardLine{}, err
			}

			newStdLine.textFields[i] = tempITextField
		}
	}

	newStdLine.numOfStdLines = txtStdLine.numOfStdLines

	return newStdLine, nil
}

// empty - Receives a pointer to an instance of
// TextLineSpecStandardLine and proceeds to set all of the internal
// member variables to their uninitialized or zero states.
//
// IMPORTANT
// ----------------------------------------------------------------
// The values of all member variables contained in input parameter
// 'txtStdLine' will be overwritten and deleted.
//
func (txtStdLineMolecule *textLineSpecStandardLineMolecule) empty(
	txtStdLine *TextLineSpecStandardLine) {

	if txtStdLineMolecule.lock == nil {
		txtStdLineMolecule.lock = new(sync.Mutex)
	}

	txtStdLineMolecule.lock.Lock()

	defer txtStdLineMolecule.lock.Unlock()

	if txtStdLine == nil {
		return
	}

	txtStdLine.numOfStdLines = 0

	txtStdLine.turnLineTerminatorOff = false

	txtStdLine.newLineChars = nil

	for i := 0; i < len(txtStdLine.textFields); i++ {
		txtStdLine.textFields[i].Empty()
		txtStdLine.textFields[i] = nil
	}

	txtStdLine.textFields = nil
}

// equal - Receives pointers to two TextLineSpecStandardLine
// instances and proceeds to compare the member data elements to
// determine whether they are equal.
//
// If the data elements of both input parameters 'blkLines' and
// 'incomingBlkLines' are equal in all respects, this method
// returns a boolean value of 'true'. Otherwise this method returns
// 'false'.
//
func (txtStdLineMolecule *textLineSpecStandardLineMolecule) equal(
	stdLineOne *TextLineSpecStandardLine,
	stdLineTwo *TextLineSpecStandardLine) bool {

	if txtStdLineMolecule.lock == nil {
		txtStdLineMolecule.lock = new(sync.Mutex)
	}

	txtStdLineMolecule.lock.Lock()

	defer txtStdLineMolecule.lock.Unlock()

	if stdLineOne == nil ||
		stdLineTwo == nil {

		return false
	}

	if stdLineOne.numOfStdLines !=
		stdLineTwo.numOfStdLines {

		return false
	}

	lenOneTextTermChars := len(stdLineOne.newLineChars)

	if lenOneTextTermChars != len(stdLineTwo.newLineChars) {
		return false
	}

	if lenOneTextTermChars > 0 {
		for i := 0; i < lenOneTextTermChars; i++ {
			if stdLineOne.newLineChars[i] !=
				stdLineTwo.newLineChars[i] {
				return false
			}
		}
	}

	if stdLineOne.turnLineTerminatorOff !=
		stdLineTwo.turnLineTerminatorOff {
		return false
	}

	lenOneTextFields := len(stdLineOne.textFields)

	if lenOneTextFields !=
		len(stdLineTwo.textFields) {
		return false
	}

	for i := 0; i < lenOneTextFields; i++ {

		if !stdLineOne.textFields[i].EqualITextField(
			stdLineTwo.textFields[i]) {
			return false
		}

	}

	return true
}

// ptr - Returns a pointer to a new instance of
// textLineSpecStandardLineMolecule.
//
func (txtStdLineMolecule textLineSpecStandardLineMolecule) ptr() *textLineSpecStandardLineMolecule {

	if txtStdLineMolecule.lock == nil {
		txtStdLineMolecule.lock = new(sync.Mutex)
	}

	txtStdLineMolecule.lock.Lock()

	defer txtStdLineMolecule.lock.Unlock()

	return &textLineSpecStandardLineMolecule{
		lock: new(sync.Mutex),
	}
}
