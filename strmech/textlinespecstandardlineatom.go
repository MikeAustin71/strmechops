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
			textLineSpecStandardLineProton{}.ptr().
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
