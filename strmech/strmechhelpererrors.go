package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type strMechHelperErrors struct {
	lock *sync.Mutex
}

// validateInputPointer
//
// This method is designed to screen for a nil pointer
// passed to a function as an input parameter. If the
// input parameter is identified as a nil pointer, a
// suitable error message is generated and returned in
// an instance of type error.
//
// This method will facilitate use of
// "If with assignment" if statements when testing for
// input errors.
//
//	Example "If with assignment"
//
//	if assignment-statement; condition {
//
//	}
//
//	Example Usage
//
//	if err = new(strMechHelperErrors).validateInputPointer(
//				inputPtr, "dirMgrPointer", ePrefix); err != nil {
//
//		return err
//
//	}
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	inputPtr					interface{}
//
//		A pointer to an object passed as an input
//		parameter for a function. If this parameter,
//		is nil, an error will be returned.
//
//	inputPtrLabel	string
//
//		The name or label associated with input parameter
//		'inputPtr' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "inputPtr" will be
//		automatically applied.
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string
//		which is included in all returned error
//		messages. Usually, it contains the name of the
//		calling method or methods listed as a function
//		chain.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref'
//		software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If input parameter 'inputPtr' non-nil, the
//		returned error Type is set equal to 'nil'.
//
//		If input parameter 'inputPtr' is nil, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errPrefDto'.
//	 	The 'errPrefDto' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (strMechHlprErrors *strMechHelperErrors) validateInputPointer(
	inputPtr interface{},
	inputPtrLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if strMechHlprErrors.lock == nil {
		strMechHlprErrors.lock = new(sync.Mutex)
	}

	strMechHlprErrors.lock.Lock()

	defer strMechHlprErrors.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"strMechHelperErrors."+
			"validateInputPointer()",
		"")

	if err != nil {

		return err
	}

	if len(inputPtrLabel) == 0 {

		inputPtrLabel = "inputPtr"
	}

	if inputPtr == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is invalid!\n"+
			"Input parameter '%v' a 'nil' pointer.\n",
			ePrefix.String(),
			inputPtrLabel,
			inputPtrLabel)

	}

	return err
}

// validateInputString
//
// This method is designed to validate input strings
// passed to a function. If the input string is
// determined to be an empty (zero length) string or
// consist entirely of blank (white) spaces (" "), it
// is classified as invalid.
//
// This method will return an error if the input
// string is found to be invalid.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	inputStr					string
//
//		This string will be analyzed to determine whether
//		it is valid. If 'inputStr' is empty (zero length)
//		or consists entirely of	blank or white spaces (" ")
//		it is classified as invalid.
//
//	inputStrLabel				string
//
//		The name or label associated with input parameter
//		'inputStr' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "inputStr" will be
//		automatically applied.
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string
//		which is included in all returned error
//		messages. Usually, it contains the name of the
//		calling method or methods listed as a function
//		chain.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref'
//		software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	err							error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errPrefDto'.
//	 	The 'errPrefDto' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (strMechHlprErrors *strMechHelperErrors) validateInputString(
	inputStr string,
	inputStrLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if strMechHlprErrors.lock == nil {
		strMechHlprErrors.lock = new(sync.Mutex)
	}

	strMechHlprErrors.lock.Lock()

	defer strMechHlprErrors.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"strMechHelperErrors."+
			"validateInputString()",
		"")

	if err != nil {

		return err
	}

	if len(inputStrLabel) == 0 {

		inputStrLabel = "inputStr"
	}

	errCode := 0

	errCode,
		_,
		_ =
		new(fileHelperElectron).
			isStringEmptyOrBlank(inputStr)

	if errCode == -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is invalid!\n"+
			"Input parameter '%v' is an empty string.\n",
			ePrefix.String(),
			inputStrLabel,
			inputStrLabel)

		return err
	}

	if errCode == -2 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is invalid!\n"+
			"Input parameter '%v' consists of blank spaces.\n",
			ePrefix.String(),
			inputStrLabel,
			inputStrLabel)
	}

	return err
}
