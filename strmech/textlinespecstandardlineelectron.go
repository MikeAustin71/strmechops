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
			ePrefix.String())

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
			ePrefix.String(),
			lastIdx,
			indexId)

		return err

	}

	if indexId == 0 {

		if txtStdLine.textFields[0] != nil {

			txtStdLine.textFields[0].Empty()

			txtStdLine.textFields[0] = nil

		}

		txtStdLine.textFields = txtStdLine.textFields[1:]

	} else if indexId == lastIdx {

		if txtStdLine.textFields[lastIdx] != nil {

			txtStdLine.textFields[lastIdx].Empty()

			txtStdLine.textFields[lastIdx] = nil

		}

		txtStdLine.textFields = txtStdLine.textFields[0:lastIdx]

	} else {

		if txtStdLine.textFields[indexId] != nil {

			txtStdLine.textFields[indexId].Empty()

			txtStdLine.textFields[indexId] = nil

		}

		txtStdLine.textFields = append(
			txtStdLine.textFields[0:indexId],
			txtStdLine.textFields[indexId+1:]...)

	}

	return err
}

// emptyTextFields - Receives a pointer to an array of Text Fields
// and proceeds to empty or delete the elements contained in that
// array.
//
// Text Fields is an array of the ITextFieldSpecification objects.
//
// The ITextFieldSpecification interface defines a text field used
// in conjunction with the type, TextLineSpecStandardLine. This
// type contains an array of text field or ITextFieldSpecification
// objects. Text fields are the building blocks of lines of text
// which are formatted by TextLineSpecStandardLine for text
// displays, file output or printing.
//
// This method will call method 'Empty()' on each of the elements
// contained in the Text Field array passed as input parameter,
// 'textFields'. Upon completion the concrete 'textFields' array
// will be set to 'nil'.
//
// IMPORTANT
//
// ----------------------------------------------------------------
//
// This method deletes and overwrites all the member elements, and
// their data values, contained in input parameter,
// 'textFields'. Upon completion, the concrete instance of the
// 'textFields' array is also set to 'nil'.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  textFields                 *[]ITextFieldSpecification
//     - A pointer to an array of text fields.
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
func (txtStdLineElectron *textLineSpecStandardLineElectron) emptyTextFields(
	textFields *[]ITextFieldSpecification,
	errPrefDto *ePref.ErrPrefixDto) (err error) {

	if txtStdLineElectron.lock == nil {
		txtStdLineElectron.lock = new(sync.Mutex)
	}

	txtStdLineElectron.lock.Lock()

	defer txtStdLineElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecStandardLineElectron."+
			"emptyTextFields()",
		"")

	if err != nil {
		return err
	}

	if textFields == nil {
		err = fmt.Errorf("%v - ERROR\n"+
			"Input parameter textFields is a 'nil' pointer!\n",
			ePrefix.String())

		return err
	}

	concreteTxtFields := *textFields

	lenTextFields := len(concreteTxtFields)

	if lenTextFields == 0 {
		return err
	}

	for i := 0; i < lenTextFields; i++ {

		if concreteTxtFields[i] == nil {
			continue
		}

		concreteTxtFields[i].Empty()

		concreteTxtFields[i] = nil
	}

	*textFields = nil

	endingLenTextFields := len(*textFields)

	if endingLenTextFields != 0 {

		err = fmt.Errorf("%v - ERROR\n"+
			"Empty Text Fields Array Operation Failed!\n"+
			"'textFields' has an array length greater than zero.\n"+
			"Begining Length 'textFields' = '%v'\n"+
			"  Ending Length 'textFields' = '%v'\n",
			ePrefix.String(),
			lenTextFields,
			endingLenTextFields)

	}

	return err
}

// equalTextFieldArrays - Compares two text fields arrays and
// returns a boolean value of 'true' if the two arrays are equal
// in all respects.
//
// To qualify as equivalent arrays of Text Fields, both arrays must
// be equal in length and contain Text Fields of equal values.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  textFields01               *[]ITextFieldSpecification
//     - If parameter 'textFields01' is equivalent, in all
//       respects, to parameter 'textFields02', this method will
//       return a boolean value of 'true'.
//
//       If this pointer to an array of ITextFieldSpecification's
//       is nil, a value of 'false' will be returned.
//
//
//  textFields02               *[]ITextFieldSpecification
//     - If parameter 'textFields02' is equivalent, in all
//       respects, to parameter 'textFields01', this method will
//       return a boolean value of 'true'.
//
//       If this pointer to an array of ITextFieldSpecification's
//       is nil, a value of 'false' will be returned.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  areEqual                   bool
//     - If Text Field Arrays 'textFields01' and 'textFields02' are
//       equal in all respects, a boolean value of 'true' is
//       returned. Otherwise, the return value is 'false'.
//
func (txtStdLineElectron *textLineSpecStandardLineElectron) equalTextFieldArrays(
	textFields01 *[]ITextFieldSpecification,
	textFields02 *[]ITextFieldSpecification) (
	areEqual bool) {

	if txtStdLineElectron.lock == nil {
		txtStdLineElectron.lock = new(sync.Mutex)
	}

	txtStdLineElectron.lock.Lock()

	defer txtStdLineElectron.lock.Unlock()

	if textFields01 == nil ||
		textFields02 == nil {
		return false
	}

	txtF01 := *textFields01

	txtF02 := *textFields02

	if txtF01 == nil &&
		txtF02 == nil {

		return true
	}

	lenTxtFields01 := len(txtF01)

	if lenTxtFields01 != len(txtF02) {
		return false
	}

	if lenTxtFields01 == 0 {
		return true
	}

	for i := 0; i < lenTxtFields01; i++ {

		if txtF01[i] == nil &&
			txtF02[i] != nil {

			return false

		} else if txtF02[i] == nil &&
			txtF01[i] != nil {

			return false

		} else if txtF02[i] == nil &&
			txtF01[i] == nil {

			continue

		} else {
			// txtF01 and txtF02 are both
			// != nil
			if !txtF01[i].EqualITextField(
				txtF02[i]) {

				return false

			}

		}
	}

	return true
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
// ------------------------------------------------------------------
//
// Input Parameters
//
//  textFields                 *[]ITextFieldSpecification
//     - 'textFields' is a pointer to a collection of objects
//       implementing the ITextLineSpecification interface. These
//       text fields are assembled by the TextLineSpecStandardLine
//       type and formatted as a single line of text. Text fields
//       are the building blocks for a standard line of text
//       characters produced by type TextLineSpecStandardLine.
//
//       If this pointer is nil, an error will be returned.
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
//  allowZeroLengthTextFieldsArray  bool
//     - When set to 'true', no error will be generated if the
//       input parameter 'textFields' contains a zero length array.
//
//       Conversely, if 'allowZeroLengthTextFieldsArray' is set to
//       'false', an error WILL be returned if 'textFields' is a
//       zero length array.
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
// -----------------------------------------------------------------
//
// Return Values
//
//  lengthTextFields           int
//     - If input parameter 'textFields' is judged to be valid in
//       all respects, this return parameter will be set to the
//       array length of 'textFields'.
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
	textFields *[]ITextFieldSpecification,
	allowZeroLengthTextFieldsArray bool,
	errPrefDto *ePref.ErrPrefixDto) (
	lengthTextFields int,
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
		"textLineSpecStandardLineElectron."+
			"testValidityOfTextFields()",
		"")

	if err != nil {
		return lengthTextFields, err
	}

	if textFields == nil {
		err = fmt.Errorf("%v - ERROR\n"+
			"Input parameter 'textFields' is a 'nil' pointer!\n",
			ePrefix.String())

		return lengthTextFields, err

	}

	lengthTextFields = 0

	var err2 error

	for idx, val := range *textFields {

		if val == nil {
			err = fmt.Errorf("%v - ERROR\n"+
				"textFields[%v] is 'nil'!\n",
				idx,
				ePrefix.String())

			return lengthTextFields, err
		}

		err2 = val.IsValidInstanceError(
			nil)

		if err2 != nil {
			err = fmt.Errorf("%v\n"+
				"Error: 'textFields' element is invalid!\n"+
				"textFields[%v] failed validity test.\n"+
				"Validity Error:\n%v\n",
				ePrefix.String(),
				idx,
				err2.Error())

			return lengthTextFields, err
		}

		lengthTextFields++
	}

	if lengthTextFields == 0 &&
		!allowZeroLengthTextFieldsArray {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'textFields' is empty!\n"+
			"'textFields' is a zero length array.\n",
			ePrefix.String())

		return lengthTextFields, err
	}

	return lengthTextFields, err
}
