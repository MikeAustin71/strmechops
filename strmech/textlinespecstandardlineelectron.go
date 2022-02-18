package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type textLineSpecStandardLineElectron struct {
	lock *sync.Mutex
}

// deleteTextField - Deletes a member of the Text Fields
// Collection. The array element to be deleted is designated by
// input parameter 'indexId'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  txtStdLine                 *TextLineSpecStandardLine
//     - A pointer to an instance of TextLineSpecStandardLine which
//       encapsulates the Text Fields Collection. The member of
//       this collection designated by parameter, 'indexId' will be
//       deleted.
//
//
//  indexId                    int
//     - The index number of the array element in the Text Fields
//       Collection which will be deleted.
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
//  err                        error
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (txtStdLineElectron *textLineSpecStandardLineElectron) deleteTextField(
	txtStdLine *TextLineSpecStandardLine,
	indexId int,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtStdLineElectron.lock == nil {
		txtStdLineElectron.lock = new(sync.Mutex)
	}

	txtStdLineElectron.lock.Lock()

	defer txtStdLineElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecStandardLineElectron."+
			"testValidityOfTextFields()",
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

	if indexId < 0 {

		err = fmt.Errorf("%v - ERROR\n"+
			"Input parameter 'indexId' is invalid!\n"+
			"'indexId' is less than zero.\n"+
			"indexId = '%v'\n",
			ePrefix.String(),
			indexId)

		return err
	}

	lenTextFieldCollection := len(txtStdLine.textFields)

	if lenTextFieldCollection == 0 {

		err = fmt.Errorf("%v - ERROR\n"+
			"The Text Field Collection, 'txtStdLine.textFields' is EMPTY!\n",
			ePrefix.XCtxEmpty().String())

		return err
	}

	lastIdx := lenTextFieldCollection - 1

	if indexId > lastIdx {

		err = fmt.Errorf("%v - ERROR\n"+
			"Input parameter 'indexId' is invalid!\n"+
			"'indexId' is greater than the last index\n"+
			"in the Text Fields Collection.\n"+
			"Last index in collection = '%v'\n"+
			"indexId = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			lastIdx,
			indexId)

		return err

	}

	if indexId == 0 {

		txtStdLine.textFields = txtStdLine.textFields[1:]

	} else if indexId == lastIdx {

		txtStdLine.textFields = txtStdLine.textFields[0:lastIdx]

	} else {

		txtStdLine.textFields = append(
			txtStdLine.textFields[0:indexId],
			txtStdLine.textFields[indexId+1:]...)

	}

	return err
}

// emptyStandardLine - This method receives an instance of
// TextLineSpecStandardLine and proceeds to set all the internal
// member variables to their zero values.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// This method will delete all existing data values contained in
// input parameter 'txtStdLine'.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  txtStdLine                 *TextLineSpecStandardLine
//     - All the internal member variables contained in input
//       parameter 'txtStdLine' will be set to their initial or
//       zero values.
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
func (txtStdLineElectron textLineSpecStandardLineElectron) emptyStandardLine(
	txtStdLine *TextLineSpecStandardLine,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtStdLineElectron.lock == nil {
		txtStdLineElectron.lock = new(sync.Mutex)
	}

	txtStdLineElectron.lock.Lock()

	defer txtStdLineElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecStandardLineElectron."+
			"emptyStandardLine()",
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

	txtStdLine.numOfStdLines = 0
	txtStdLine.turnLineTerminatorOff = false
	txtStdLine.newLineChars = nil
	txtStdLine.textLineReader = nil

	lenTextFields := len(txtStdLine.textFields)

	if lenTextFields == 0 {
		txtStdLine.textFields = nil
		return nil
	}

	for i := 0; i < lenTextFields; i++ {

		if txtStdLine.textFields[i] == nil {
			continue
		}

		txtStdLine.textFields[i].Empty()

		txtStdLine.textFields[i] = nil

	}

	txtStdLine.textFields = nil

	return nil
}

// ptr - Returns a pointer to a new instance of
// textLineSpecStandardLineElectron.
//
func (txtStdLineElectron textLineSpecStandardLineElectron) ptr() *textLineSpecStandardLineElectron {

	if txtStdLineElectron.lock == nil {
		txtStdLineElectron.lock = new(sync.Mutex)
	}

	txtStdLineElectron.lock.Lock()

	defer txtStdLineElectron.lock.Unlock()

	return &textLineSpecStandardLineElectron{
		lock: new(sync.Mutex),
	}
}

// testValidityOfTextFields - This method receives an array of
// ITextFieldSpecification objects and subjects that array to
// an analysis which determines whether the array, and all of its
// member elements, is valid in all respects.
//
// Type TextLineSpecStandardLine is a text specification for a
// standard line of text. A standard line of text comprises a
// single line of text which may be repeated one or more times.
//
// TextLineSpecStandardLine encapsulates an array of
// ITextFieldSpecification objects which are used to format text
// fields within a single line of text. Essentially, a standard
// text line is a collection of text fields which implement the
// ITextFieldSpecification interface. Text fields can be thought of
// as the building blocks for a standard line of text.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  textFields                 []ITextFieldSpecification
//     - 'textFields' is a collection of objects implementing the
//       ITextLineSpecification interface. These text fields are
//       assembled by the TextLineSpecStandardLine type and
//       formatted as a single line of text. Text fields are the
//       building blocks for a standard line of text characters
//       produced by type TextLineSpecStandardLine.
//
//       If this array is 'nil' or has a zero length, an error will
//       be returned.
//
//       If any of the ITextFieldSpecification objects within this
//       array are found to be invalid, an error will be returned.
//
//       If any of the individual elements within this array as a
//       'nil' value, an error will be returned.
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
//  isValid                    bool
//     - If input parameter 'textFields' is judged to be valid in
//       all respects, this return parameter will be set to 'true'.
//
//     - If input parameter 'textFields' is found to be invalid,
//       this return parameter will be set to 'false'.
//
//
//  err                        error
//     - If input parameter 'textFields' is judged to be valid in
//       all respects, this return parameter will be set to 'nil'.
//
//       If input parameter, 'textFields' is found to be invalid,
//       this return parameter will be configured with an appropriate
//       error message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (txtStdLineElectron *textLineSpecStandardLineElectron) testValidityOfTextFields(
	textFields []ITextFieldSpecification,
	errPrefDto *ePref.ErrPrefixDto) (
	isValid bool,
	err error) {

	if txtStdLineElectron.lock == nil {
		txtStdLineElectron.lock = new(sync.Mutex)
	}

	txtStdLineElectron.lock.Lock()

	defer txtStdLineElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	isValid = false

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecStandardLineElectron."+
			"testValidityOfTextFields()",
		"")

	if err != nil {
		return isValid, err
	}

	if textFields == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'textFields' is invalid!\n"+
			"The value of 'textFields' is 'nil'.\n",
			ePrefix.String())

		return isValid, err
	}

	lenTxtFields := len(textFields)

	if lenTxtFields == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'textFields' is invalid!\n"+
			"'textFields' is a zero length array.\n",
			ePrefix.String())

		return isValid, err
	}

	var err2 error

	for i := 0; i < lenTxtFields; i++ {

		if textFields[i] == nil {
			err = fmt.Errorf("%v\n"+
				"Error: 'textFields' element is invalid!\n"+
				"textFields[%v] is 'nil'.\n",
				ePrefix.String(),
				i)

			return isValid, err
		}

		err2 = textFields[i].
			IsValidInstanceError(
				ePrefix)

		if err2 != nil {
			err = fmt.Errorf("%v\n"+
				"Error: 'textFields' element is invalid!\n"+
				"textFields[%v] failed validity test.\n"+
				"Validity Error:\n%v\n",
				ePrefix.String(),
				i,
				err2.Error())

			return isValid, err
		}
	}

	isValid = true

	return isValid, err
}
