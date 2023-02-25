package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type fileAccessControlElectron struct {
	lock *sync.Mutex
}

// testValidityOfFileAccessControl
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//	isValid						bool
//
//		If any of the internal member data variables
//		contained in the input parameter 'fAccessCtrl'
//		are found to be invalid, this return parameter
//		will be set to 'false'.
//
//		If all the internal member data variables
//		contained in the input parameter 'fAccessCtrl'
//		are found to be valid, this return parameter
//		will be set to 'true'.
//
//	err							error
//
//		If any of the internal member data variables
//		contained in the input parameter 'fAccessCtrl'
//		are found to be invalid, this method will return
//		an error configured with an appropriate message
//		identifying the invalid member data variable.
//
//		If all internal member data variables evaluate
//		as valid, this returned error value will be set
//		to 'nil'.
//
//		If errors are encountered during processing or if
//		any internal member data values are found to be
//		invalid, the returned error Type will encapsulate
//		an appropriate error message. This returned error
//		message will incorporate the method chain and text
//		passed by input parameter, 'errPrefDto'. The
//		'errPrefDto' text will be attached to the
//		beginning of the error message.
func (fAccessCtrlElectron *fileAccessControlElectron) testValidityOfFileAccessControl(
	fAccessCtrl *FileAccessControl,
	errPrefDto *ePref.ErrPrefixDto) (
	isValid bool,
	err error) {

	if fAccessCtrlElectron.lock == nil {
		fAccessCtrlElectron.lock = new(sync.Mutex)
	}

	fAccessCtrlElectron.lock.Lock()

	defer fAccessCtrlElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	isValid = false

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileAccessControlElectron."+
			"testValidityOfFileAccessControl()",
		"")

	if err != nil {

		return isValid, err
	}

	if fAccessCtrl == nil {
		err = fmt.Errorf("%v\n"+
			"Error: FileAccessControl object is invalid!\n"+
			"Input parameter 'fAccessCtrl' is a nil pointer!\n",
			ePrefix.String())

		return isValid, err
	}

	if !fAccessCtrl.isInitialized {
		return isValid,
			fmt.Errorf("%v\n"+
				"Error: FileAccessControl object is invalid!\n"+
				"The current FileAccessControl Instance has NOT been initialized!\n",
				ePrefix)
	}

	var err2 error

	err2 = fAccessCtrl.fileOpenCodes.IsValidInstanceError(ePrefix)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: FileAccessControl object is invalid!\n"+
			"File Open codes INVALID!\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			err2.Error())

		return isValid, err
	}

	err2 = fAccessCtrl.permissions.IsValid(ePrefix)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: FileAccessControl object is invalid!\n"+
			"File Permission codes INVALID!\n"+
			"Error = \n%v\n",
			ePrefix.String(),
			err2.Error())

		return isValid, err
	}

	isValid = true

	return isValid, err
}
