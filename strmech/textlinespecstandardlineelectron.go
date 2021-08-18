package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type textLineSpecStandardLineElectron struct {
	lock *sync.Mutex
}

// copyTextFields - Copies an array of ITextFieldSpecification
// objects from a source array to a target array.
//
// If any of the ITextFieldSpecification objects in the source
// array input parameter 'sourceTextFields' are found to be
// invalid, an error will be returned.
//
// The ITextFieldSpecification interface defines a text field used
// in conjunction with the type, TextLineSpecStandardLine. This
// type contains an array of text field or ITextFieldSpecification
// objects. Text fields are the building blocks of lines of text
// which are formatted by TextLineSpecStandardLine for text
// displays, file output or printing.
//
// Often, the need arises to copy text fields between
// TextLineSpecStandardLine objects. This method is designed to
// facilitate those copy operations.
//
// IMPORTANT
//
// ----------------------------------------------------------------
//
// This method deletes and overwrites all the member elements, and
// their data values, contained in input parameter,
// 'targetTextFields'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  targetTextFields           []ITextFieldSpecification
//     - All the array elements within input parameter
//       'sourceTextFields' will be copied to this array,
//       'targetTextFields'. When the copy operation is completed
//       the elements and their data values contained in this array
//       will be identical to those in 'sourceTextFields'.
//
//
//  sourceTextFields           []ITextFieldSpecification
//     - All the data elements in this array will be copied to the
//       input parameter 'targetTextFields'. When the copy
//       operation is completed all the array elements and their
//       data values in 'targetTextFields' will be identical to
//       those found in this array, 'sourceTextFields'.
//
//       If 'sourceTextFields' contains an empty or zero length
//       array, an error will be returned.
//
//       If any of the ITextFieldSpecification objects in this
//       array are found to be invalid, an error will be returned.
//
//
//  errPrefDto          *ePref.ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods listed
//       as a function chain.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
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
func (txtStdLineElectron *textLineSpecStandardLineElectron) copyTextFields(
	targetTextFields []ITextFieldSpecification,
	sourceTextFields []ITextFieldSpecification,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if txtStdLineElectron.lock == nil {
		txtStdLineElectron.lock = new(sync.Mutex)
	}

	txtStdLineElectron.lock.Lock()

	defer txtStdLineElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecStandardLineElectron.copyTextFields()",
		"")

	if err != nil {
		return err
	}

	lenSourceTxtFields := len(sourceTextFields)

	if lenSourceTxtFields == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'sourceTextFields' is a zero length\n"+
			"empty array!\n",
			ePrefix.String())

		return err
	}

	lenTargetTxtFields := len(targetTextFields)

	if lenTargetTxtFields > 0 {

		for i := 0; i < lenTargetTxtFields; i++ {

			if targetTextFields[i] == nil {
				continue
			}

			targetTextFields[i].Empty()

			targetTextFields[i] = nil
		}
	}

	targetTextFields = nil

	targetTextFields =
		make(
			[]ITextFieldSpecification,
			lenSourceTxtFields)

	var newITextField ITextFieldSpecification

	for i := 0; i < lenSourceTxtFields; i++ {

		if sourceTextFields[i] == nil {

			err = fmt.Errorf("%v\n"+
				"Error: Incoming Text Field is invalid!\n"+
				"sourceTextFields[%v] has a 'nil' value.\n",
				ePrefix.XCtx(
					fmt.Sprintf(
						"sourceTextFields[%v] == nil",
						i)),
				i)

			return err
		}

		newITextField,
			err = sourceTextFields[i].CopyOutITextField(
			ePrefix.XCtx(
				fmt.Sprintf("sourceTextFields[%v]",
					i)))

		if err != nil {
			return err
		}

		targetTextFields[i] = newITextField
	}

	return err
}

// emptyTextFields - Receives a pointer to an instance of
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
func (txtStdLineElectron *textLineSpecStandardLineElectron) emptyTextFields(
	txtStdLine *TextLineSpecStandardLine) {

	if txtStdLineElectron.lock == nil {
		txtStdLineElectron.lock = new(sync.Mutex)
	}

	txtStdLineElectron.lock.Lock()

	defer txtStdLineElectron.lock.Unlock()

	if txtStdLine == nil {
		return
	}

	for i := 0; i < len(txtStdLine.textFields); i++ {

		if txtStdLine.textFields[i] == nil {
			continue
		}

		txtStdLine.textFields[i].Empty()

		txtStdLine.textFields[i] = nil
	}

	txtStdLine.textFields = nil

	return
}

// testValidityOfTextLineSpecStdLine - Receives a pointer to an
// instance of TextLineSpecStandardLine and performs a diagnostic
// analysis to determine if that instance is valid in all respects.
//
// If the input parameter 'txtStdLine' is determined to be invalid,
// this method will return a boolean flag ('isValid') of 'false'.
// In addition, an instance of type error ('err') will be returned
// configured with an appropriate error message.
//
// If the input parameter 'txtStdLine' is valid, this method will
// return a boolean flag ('isValid') of 'true' and the returned
// error type ('err') will be set to 'nil'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  txtStdLine                 *TextLineSpecStandardLine
//     - A pointer to an instance of TextLineSpecStandardLine. This
//       object will be subjected to diagnostic analysis in order
//       to determine if all the member variables contain valid
//       values.
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
//  isValid                    bool
//     - If input parameter 'txtStdLine' is judged to be valid in
//       all respects, this return parameter will be set to 'true'.
//
//     - If input parameter 'txtStdLine' is found to be invalid,
//       this return parameter will be set to 'false'.
//
//
//  err                        error
//     - If input parameter 'txtStdLine' is judged to be valid in
//       all respects, this return parameter will be set to 'nil'.
//
//       If input parameter, 'txtStdLine' is found to be invalid,
//       this return parameter will be configured with an appropriate
//       error message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (txtStdLineElectron *textLineSpecStandardLineElectron) testValidityOfTextLineSpecStdLine(
	txtStdLine *TextLineSpecStandardLine,
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
		"textLineSpecStandardLineElectron.testValidityOfTextLineSpecStdLine()",
		"")

	if err != nil {
		return isValid, err
	}

	if txtStdLine == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtStdLine' is a nil pointer!\n",
			ePrefix.String())

		return isValid, err
	}

	if txtStdLine.numOfStdLines < 1 {
		err = fmt.Errorf("%v\n"+
			"Error: The number of standard lines is less than one ('1')!\n"+
			"This means that no lines will be generated by this specification.\n",
			ePrefix.String())

		return isValid, err
	}

	lenTextFields := len(txtStdLine.textFields)

	if lenTextFields == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: No Text Fields have been configured for\n"+
			"this standard line specification!\n",
			ePrefix.String())

		return isValid, err
	}

	var err2 error

	for i := 0; i < lenTextFields; i++ {

		if txtStdLine.textFields[i] == nil {
			err = fmt.Errorf("%v\n"+
				"Collection Element Error: A Starndard Line Text Field is invalid!\n"+
				" txtStdLine.textFields[%v] has a 'nil' value.\n",
				ePrefix.XCtx("txtStdLine"),
				i)

			return isValid, err
		}

		err2 = txtStdLine.textFields[i].IsValidInstanceError(ePrefix)

		if err2 != nil {
			err = fmt.Errorf("%v\n"+
				"Error: Text Field Element [%v] is invalid!\n"+
				"Text Field Element Error = \n%v\n",
				ePrefix.String(),
				i,
				err2.Error())

			return isValid, err
		}
	}

	if len(txtStdLine.newLineChars) == 0 {
		txtStdLine.newLineChars = []rune{'\n'}
	}

	isValid = true

	return isValid, err
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
