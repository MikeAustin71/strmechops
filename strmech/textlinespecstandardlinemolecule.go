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
// ----------------------------------------------------------------
//
// IMPORTANT
//
// The pre-existing data fields for input parameter 'targetStdLine'
// will be overwritten and deleted.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  targetStdLine              *TextLineSpecStandardLine
//     - A pointer to an instance of TextLineSpecStandardLine. Data
//       extracted from input parameter 'incomingStdLine' will
//       be copied to this input parameter, 'targetStdLine'. If
//       this method completes successfully, all member data
//       variables encapsulated in 'targetStdLine' will be
//       identical to those contained in input parameter,
//       'incomingStdLine'.
//
//       Be advised that the pre-existing data fields in input
//       parameter 'targetStdLine' will be overwritten and
//       deleted.
//
//
//  incomingStdLine            *TextLineSpecStandardLine
//     - A pointer to an instance of TextLineSpecStandardLine.
//
//       All data values in this TextLineSpecStandardLine instance
//       will be copied to input parameter 'targetStdLine'.
//
//       The original member variable data values encapsulated in
//       'incomingStdLine' will remain unchanged and will NOT be
//       overwritten or deleted.
//
//       If 'incomingStdLine' contains invalid member data
//       variables, this method will return an error.
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
		"textLineSpecStandardLineMolecule."+
			"copyIn()",
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

	txtStdLineAtom :=
		textLineSpecStandardLineAtom{}

	_,
		err = txtStdLineAtom.
		testValidityOfTextLineSpecStdLine(
			incomingStdLine,
			ePrefix.XCtx(
				"incomingStdLine"))

	if err != nil {
		return err
	}

	if len(incomingStdLine.newLineChars) == 0 {
		incomingStdLine.newLineChars = []rune{'\n'}
	}

	sMechPreon := strMechPreon{}

	err = sMechPreon.copyRuneArrays(
		&targetStdLine.newLineChars,
		&incomingStdLine.newLineChars,
		true,
		ePrefix.XCtx(
			"incomingStdLine.newLineChars->"+
				"targetStdLine.newLineChars"))

	if err != nil {
		return err
	}

	targetStdLine.turnLineTerminatorOff =
		incomingStdLine.turnLineTerminatorOff

	targetStdLine.numOfStdLines =
		incomingStdLine.numOfStdLines

	return txtStdLineAtom.
		copyTextFields(
			&targetStdLine.textFields,
			&incomingStdLine.textFields,
			ePrefix.XCtx(
				"incomingStdLine.textFields->"+
					"targetStdLine.textFields"))
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
//       If the member variable data values encapsulated by this
//       'txtStdLine' are found to be invalid, this method will
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

	txtStdLineAtom :=
		textLineSpecStandardLineAtom{}

	_,
		err = txtStdLineAtom.
		testValidityOfTextLineSpecStdLine(
			txtStdLine,
			ePrefix.XCtx(
				"txtStdLine"))

	if err != nil {
		return TextLineSpecStandardLine{}, err
	}

	newStdLine := TextLineSpecStandardLine{}

	newStdLine.lock = new(sync.Mutex)

	lenInNewLineChars := len(txtStdLine.newLineChars)

	if lenInNewLineChars == 0 {
		txtStdLine.newLineChars = []rune{'\n'}
		lenInNewLineChars = 1
	}

	sMechPreon := strMechPreon{}

	err = sMechPreon.copyRuneArrays(
		&newStdLine.newLineChars,
		&txtStdLine.newLineChars,
		true,
		ePrefix.XCtx(
			"txtStdLine.newLineChars->"+
				"newStdLine.newLineChars"))

	if err != nil {
		return TextLineSpecStandardLine{}, err
	}

	newStdLine.turnLineTerminatorOff =
		txtStdLine.turnLineTerminatorOff

	newStdLine.numOfStdLines = txtStdLine.numOfStdLines

	err = txtStdLineAtom.
		copyTextFields(
			&newStdLine.textFields,
			&txtStdLine.textFields,
			ePrefix.XCtx(
				"txtStdLine.textFields->"+
					"newStdLine.textFields"))

	return newStdLine, err
}

// emptyStdLineTextFields - Receives a pointer to an instance of
// TextLineSpecStandardLine and proceeds to delete all the text
// fields contained in the internal text field collection.
//
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All the text fields stored in the text field collection
// maintained by input parameter 'txtStdLine' will be deleted.
//
func (txtStdLineMolecule *textLineSpecStandardLineMolecule) emptyStdLineTextFields(
	txtStdLine *TextLineSpecStandardLine) {

	if txtStdLineMolecule.lock == nil {
		txtStdLineMolecule.lock = new(sync.Mutex)
	}

	txtStdLineMolecule.lock.Lock()

	defer txtStdLineMolecule.lock.Unlock()

	textLineSpecStandardLineElectron{}.ptr().
		emptyTextFields(
			txtStdLine.textFields)
}

// equal - Receives pointers to two TextLineSpecStandardLine
// instances and proceeds to compare the member data elements to
// determine whether they are equal.
//
// If the data elements of both input parameters 'stdLineOne' and
// 'stdLineTwo' are equal in all respects, this method returns a
// boolean value of 'true'. Otherwise, this method returns 'false'.
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

	sMechPreon := strMechPreon{}

	if !sMechPreon.equalRuneArrays(
		stdLineOne.newLineChars,
		stdLineTwo.newLineChars) {
		return false
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

// getFormattedText - Returns the formatted text generated by this
// Text Line Specification, 'txtStdLine', for output and printing.
//
// The standard line may be replicated multiple times if the
// value of internal member variable' stdLine.numOfStdLines' is
// greater than one ('1').
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  txtStdLine                 *TextLineSpecStandardLine
//     - A pointer to an instance of TextLineSpecStandardLine. The
//       member variables encapsulated by this instance will be
//       used to generate formatted text for output display and
//       printing.
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
//  formattedText              string
//     - If this method completes successfully, a string of
//       formatted text will be generated from the data provided by
//       input parameter 'txtStdLine'.
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
func (txtStdLineMolecule *textLineSpecStandardLineMolecule) getFormattedText(
	txtStdLine *TextLineSpecStandardLine,
	errPrefDto *ePref.ErrPrefixDto) (
	formattedText string,
	err error) {

	if txtStdLineMolecule.lock == nil {
		txtStdLineMolecule.lock = new(sync.Mutex)
	}

	txtStdLineMolecule.lock.Lock()

	defer txtStdLineMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	formattedText = ""

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecStandardLineMolecule.getFormattedText()",
		"")

	if err != nil {
		return formattedText, err
	}

	if txtStdLine == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtStdLine' is a nil pointer!\n",
			ePrefix.String())

		return formattedText, err
	}

	_,
		err = textLineSpecStandardLineAtom{}.ptr().
		testValidityOfTextLineSpecStdLine(
			txtStdLine,
			ePrefix.XCtx("txtStdLine"))

	if err != nil {
		return formattedText, err
	}

	if len(txtStdLine.newLineChars) == 0 {
		txtStdLine.newLineChars = []rune{'\n'}
	}

	lenTextFields := len(txtStdLine.textFields)

	if lenTextFields == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: txtStdLine.textFields are invalid!\n"+
			"Length of txtStdLine.textFields is zero.",
			ePrefix.String())

		return formattedText, err
	}

	var lineStr string

	for i := 0; i < lenTextFields; i++ {

		if txtStdLine.textFields[i] == nil {
			err = fmt.Errorf("%v\n"+
				"Error: 'txtStdLine.textFields' is an invalid array!\n"+
				"txtStdLine.textFields[%v] is a 'nil'value.\n",
				ePrefix.XCtx(
					"txtStdLine.textFields[i] == nil"),
				i)

			return formattedText, err
		}

		lineStr += txtStdLine.textFields[i].GetFormattedText()
	}

	for j := 0; j < txtStdLine.numOfStdLines; j++ {
		formattedText += lineStr
	}

	if txtStdLine.turnLineTerminatorOff == true {

		return formattedText, err
	}

	formattedText += string(txtStdLine.newLineChars)

	return formattedText, err
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

// setTxtSpecStandardLine - Reconfigures all the data values for
// input parameter 'txtStdLine', a pointer to an instance of
// TextLineSpecStandardLine.
func (txtStdLineMolecule *textLineSpecStandardLineMolecule) setTxtSpecStandardLine(
	txtStdLine *TextLineSpecStandardLine,
	numOfStdLines int,
	textFields []ITextFieldSpecification,
	newLineChars []rune,
	turnLineTerminatorOff bool,
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
		"textLineSpecStandardLineMolecule."+
			"setTxtSpecStandardLine()",
		"")

	if err != nil {
		return err
	}

	if txtStdLine == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtStdLine' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if len(textFields) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'textFields' is invalid!\n"+
			"'textFields' is an empty or zero length array.\n",
			ePrefix.String())

		return err
	}

	if numOfStdLines < 1 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numOfStdLines' is invalid!\n"+
			"'numOfStdLines' has a value less than one (1).\n"+
			"numOfStdLines = '%v'\n",
			ePrefix.String(),
			numOfStdLines)

		return err
	}

	if numOfStdLines > 1000000 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numOfStdLines' is invalid!\n"+
			"'numOfStdLines' has a value greater than one-million (1,000,000).\n"+
			"numOfStdLines = '%v'\n",
			ePrefix.String(),
			numOfStdLines)

		return err
	}

	if len(newLineChars) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'newLineChars' is invalid!\n"+
			"'newLineChars' is an empty or zero length array.\n",
			ePrefix.String())

		return err
	}

	sMechPreon := strMechPreon{}

	_,
		err = sMechPreon.testValidityOfRuneCharArray(
		newLineChars,
		ePrefix.XCtx(
			"newLineChars"))

	if err != nil {
		return err
	}

	err = textLineSpecStandardLineAtom{}.ptr().
		copyTextFields(
			&txtStdLine.textFields,
			&textFields,
			ePrefix.XCtx(
				"textFields->"+
					"txtStdLine.textFields"))

	if err != nil {
		return err
	}

	txtStdLine.numOfStdLines =
		numOfStdLines

	txtStdLine.turnLineTerminatorOff =
		turnLineTerminatorOff

	err = sMechPreon.copyRuneArrays(
		&txtStdLine.newLineChars,
		&newLineChars,
		true,
		ePrefix.XCtx(
			"newLineChars->txtStdLine.newLineChars"))

	return err
}
