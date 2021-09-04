package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type textLineSpecStandardLineElectron struct {
	lock *sync.Mutex
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
