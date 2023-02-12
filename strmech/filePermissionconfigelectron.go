package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"os"
	"sync"
)

type filePermissionConfigElectron struct {
	lock *sync.Mutex
}

// testValidityOfFilePermissionConfig
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fPerm						*FilePermissionConfig
//
//		A pointer to an instance of FilePermissionConfig.
//		No data elements in this instance will be modified.
//
//		The internal member data elements contained in this
//		instance will be analyzed to determine if they are
//		valid in all respects.
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
//		If input parameter 'filePermConfig' is judged to
//		be valid in all respects, this return parameter
//		will be set to 'true'.
//
//		If input parameter 'filePermConfig' is found to
//		be invalid, this return parameter will be set to
//		'false'.
//
//	err							error
//
//		If errors are encountered during processing or
//		if input parameter 'filePermConfig' is found to
//		be invalid, this returned error Type will
//		encapsulate an appropriate error message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
//
//		If input parameter 'filePermConfig' is found to
//		be valid and no errors are encountered during
//		processing, this returned error parameter is set
//		to 'nil'.
func (filePermCfgElectron *filePermissionConfigElectron) testValidityOfFilePermissionConfig(
	filePermConfig *FilePermissionConfig,
	errPrefDto *ePref.ErrPrefixDto) (
	isValid bool,
	err error) {

	if filePermCfgElectron.lock == nil {
		filePermCfgElectron.lock = new(sync.Mutex)
	}

	filePermCfgElectron.lock.Lock()

	defer filePermCfgElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	isValid = false

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"filePermissionConfigElectron."+
			"testValidityOfFilePermissionConfig()",
		"")

	if err != nil {
		return isValid, err
	}

	if filePermConfig == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'filePermConfig' is a nil pointer!\n",
			ePrefix.String())

		return isValid, err
	}

	fMode := filePermConfig.fileMode &^ os.FileMode(0777)

	isValid = false

	for idx := range mOsPermissionCodeToString {

		if fMode == idx {

			isValid = true

			break
		}

	}

	if !isValid {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of FilePermissionConfig is invalid!\n"+
			"The Entry Type File Mode value is invalid!\n"+
			"FilePermissionConfig File Mode = %v\n",
			ePrefix.String(),
			filePermConfig.fileMode)
	}

	return isValid, err
}
