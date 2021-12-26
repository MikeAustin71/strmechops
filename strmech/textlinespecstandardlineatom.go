package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type textLineSpecStandardLineAtom struct {
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
//  targetTextFields           *[]ITextFieldSpecification
//     - A pointer to the target text fields array.
//
//       All the array elements within input parameter
//       'sourceTextFields' will be copied to this array,
//       'targetTextFields'. When the copy operation is completed
//       the elements and their data values contained in this array
//       will be identical to those in 'sourceTextFields'.
//
//       Be advised, all the elements in the target text fields
//       array will be deleted and overwritten.
//
//
//  sourceTextFields           *[]ITextFieldSpecification
//     - A pointer to the source text fields array.
//
//       All the data elements in this array will be copied to the
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
func (txtStdLineAtom *textLineSpecStandardLineAtom) copyTextFields(
	targetTextFields *[]ITextFieldSpecification,
	sourceTextFields *[]ITextFieldSpecification,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if txtStdLineAtom.lock == nil {
		txtStdLineAtom.lock = new(sync.Mutex)
	}

	txtStdLineAtom.lock.Lock()

	defer txtStdLineAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecStandardLineAtom."+
			"copyTextFields()",
		"")

	if err != nil {
		return err
	}

	if targetTextFields == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetTextFields' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if sourceTextFields == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'sourceTextFields' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if *sourceTextFields == nil {

		*targetTextFields = nil

		return err
	}

	lenSourceTxtFields := len(*sourceTextFields)

	if lenSourceTxtFields == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'sourceTextFields' is a zero length\n"+
			"empty array!\n",
			ePrefix.String())

		return err
	}

	*targetTextFields =
		make(
			[]ITextFieldSpecification,
			lenSourceTxtFields)

	itemsCopied :=
		copy(*targetTextFields, *sourceTextFields)

	if itemsCopied != lenSourceTxtFields {
		err = fmt.Errorf("%v\n"+
			"Error: Copy Operation Failed!\n"+
			"Number of elements copied from 'sourceTextFields' to 'targetTextFields'\n"+
			"DOES NOT MATCH the number of elements in 'sourceTextFields'\n"+
			"Number of elements copied to targetTextFields = '%v'\n"+
			"     Number of elements in 'sourceTextFields' = '%v'\n",
			ePrefix.String(),
			itemsCopied,
			lenSourceTxtFields)
	}

	return err
}

// peekPopTextField - Performs either a 'Peek' or 'Pop' operation
// on an array element in the Text Fields Collection as specified
// by the input parameter, 'popTextField'.
//
// A 'Pop' operation returns a deep copy of the last Text Field
// in the Text Field Collection and then deletes that designated
// array element.
//
// On the other hand, a 'Peek' operation will return a deep copy of
// the designated Text Field in the Text Field Collection and WILL
// NOT delete that array element. The designated array element
// therefore remains in the collection after the 'Peek' operation
// is completed.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  txtStdLine                 *TextLineSpecStandardLine
//     - A pointer to an instance of TextLineSpecStandardLine. A
//       deep copy of the first Text Field in the Text Fields
//       Collection for this instance of TextLineSpecStandardLine
//       will be returned to the calling function.
//
//       Depending on the value of input parameter, 'popTextField',
//       either an 'Peek' or 'Pop' operation will be performed on
//       the  first Text Field in the Text Fields Collection.
//
//
//  indexId                    int
//     - The index number of the array element in the Text Fields
//       Collection on which the 'Pop' or 'Peek' operation will be
//       performed.
//
//
//  popTextField               bool
//     - If this parameter is set to 'true', it signals that a
//       'Pop' operation will be performed on the designated Text
//       Field in the Text Fields Collection encapsulated in
//       parameter 'txtStdLine'. A 'Pop' operation will delete the
//       designated Text Field from the Text Fields Collection.
//
//       If this parameter is set to 'false', it signals that a
//       'Peek' operation will be performed on the designated Text
//       Field in the Text Fields Collection encapsulated in
//       parameter 'txtStdLine'. A 'Peek' operation means that the
//       designated Text Field element in the Text Fields
//       Collection WILL NOT be deleted and will remain in the
//       collection.
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
//  iTextField                 ITextFieldSpecification
//     - If this method completes successfully, a deep copy of
//       if the designated member of the Text Fields Collection
//       will be returned to the calling function. The returned
//       object will implement the ITextFieldSpecification
//       interface.
//
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
func (txtStdLineAtom *textLineSpecStandardLineAtom) peekPopTextField(
	txtStdLine *TextLineSpecStandardLine,
	indexId int,
	popTextField bool,
	errPrefDto *ePref.ErrPrefixDto) (
	iTextField ITextFieldSpecification,
	err error) {

	if txtStdLineAtom.lock == nil {
		txtStdLineAtom.lock = new(sync.Mutex)
	}

	txtStdLineAtom.lock.Lock()

	defer txtStdLineAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	iTextField = nil

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecStandardLineAtom."+
			"peekPopFirstTextField()",
		"")

	if err != nil {
		return iTextField, err
	}

	if txtStdLine == nil {

		err = fmt.Errorf("%v - ERROR\n"+
			"Input parameter 'txtStdLine' is a nil pointer!\n",
			ePrefix.String())

		return iTextField, err
	}

	lenTextFieldCollection := len(txtStdLine.textFields)

	lastIdx := lenTextFieldCollection - 1

	if lenTextFieldCollection == 0 {

		err = fmt.Errorf("%v - ERROR\n"+
			"The Text Field Collection, 'txtStdLine.textFields' is EMPTY!\n",
			ePrefix.XCtxEmpty().String())

		return iTextField, err
	}

	if indexId < 0 {

		err = fmt.Errorf("%v - ERROR\n"+
			"Input parameter 'indexId' is invalid!\n"+
			"'indexId' is less than zero.\n"+
			"indexId = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			indexId)

		return iTextField, err
	}

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

		return iTextField, err
	}

	iTextField,
		err = txtStdLine.textFields[indexId].CopyOutITextField(
		ePrefix.XCtx(
			fmt.Sprintf(
				"txtStdLine.textFields[%v]",
				indexId)))

	if err != nil {
		return iTextField, err
	}

	if !popTextField {
		// popTextField == false
		// This means a 'peek' operation is
		// being performed and the designated
		// array element will NOT be deleted.
		return iTextField, err
	}

	err =
		textLineSpecStandardLineElectron{}.ptr().
			deleteTextField(
				txtStdLine,
				indexId,
				ePrefix.XCtx(
					fmt.Sprintf(
						"Deleting txtStdLine.textFields[%v]",
						indexId)))

	return iTextField, err
}

// ptr - Returns a pointer to a new instance of
// textLineSpecStandardLineAtom.
//
func (txtStdLineAtom textLineSpecStandardLineAtom) ptr() *textLineSpecStandardLineAtom {

	if txtStdLineAtom.lock == nil {
		txtStdLineAtom.lock = new(sync.Mutex)
	}

	txtStdLineAtom.lock.Lock()

	defer txtStdLineAtom.lock.Unlock()

	return &textLineSpecStandardLineAtom{
		lock: new(sync.Mutex),
	}
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
// If txtStdLine.newLineChars is a zero length array, this method
// will automatically set this value to the default new line
// character or characters.
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
func (txtStdLineAtom *textLineSpecStandardLineAtom) testValidityOfTextLineSpecStdLine(
	txtStdLine *TextLineSpecStandardLine,
	errPrefDto *ePref.ErrPrefixDto) (
	isValid bool,
	err error) {

	if txtStdLineAtom.lock == nil {
		txtStdLineAtom.lock = new(sync.Mutex)
	}

	txtStdLineAtom.lock.Lock()

	defer txtStdLineAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	isValid = false

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecStandardLineAtom.testValidityOfTextLineSpecStdLine()",
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

		err =
			textSpecificationMolecule{}.ptr().
				setDefaultNewLineChars(
					&txtStdLine.newLineChars,
					ePrefix)

		if err != nil {
			return isValid, err
		}

	}

	isValid = true

	return isValid, err
}
