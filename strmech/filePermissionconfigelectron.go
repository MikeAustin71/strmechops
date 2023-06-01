package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"os"
	"strings"
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
//	filePermConfig				*FilePermissionConfig
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
//		If an error occurs, the text value for input
//		parameter 'errPrefDto' (error prefix) will be
//		prefixed or attached at the beginning of the
//		error message.
//
//		If input parameter 'filePermConfig' is found to
//		be valid, and no errors are encountered during
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

	if !filePermConfig.isInitialized {

		err = fmt.Errorf("%v\n"+
			"Error: This FilePermissionConfig instance has\n"+
			"NOT been initialized and is INVALID!\n",
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
			"FilePermissionConfig File Mode Value  = %v\n"+
			"FilePermissionConfig File Mode String = %v\n",
			ePrefix.String(),
			filePermConfig.fileMode,
			filePermConfig.fileMode.String())
	}

	return isValid, err
}

// convertGroupToDecimal
//
// Receives a 3-character permission letter code for a
// user class specified as 'owner', 'group' or
// 'other'.
//
// The 3-character permission letter code must be
// formatted as one of the following character strings:
//
//	(1)	"rwx"	- Read, Write and Execute
//	(2)	"rw-"	- Read and Write
//	(3)	"r--"	- Read
//	(4)	"---"	- None
//	(5)	"--x"	- Execute
//	(6)	"-wx"	- Write and Execute
//	(7)	"-w-"	- Write
//	(8)	"r-x"	- Read and Execute
//
// If input parameter 'permissionLetterCode' does not
// match one of the 8-character strings shown above, an
// error will be returned.
//
// The user class must be specified as one of the
// following:
//
//	(1)	'owner'
//	(2)	'group'
//	(3)	'other'
//
// If input parameter 'userClass' does not match one of
// the three user classes listed above, an error will be
// returned.
//
// This method will return an integer value representing
// the octal digits comprising this permission letter
// code. For example, permissionLetterCode ="rwx" will
// return an integer value of '7' which can be treated as
// octal digit '7' for purposes of creating an os.FileMode.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://www.cyberciti.biz/faq/explain-the-nine-permissions-bits-on-files/
//	https://en.wikipedia.org/wiki/File_system_permissions
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	permissionLetterCode		string
//
//		This string must consist of 3-characters which
//		comprise the permission letter group code. The
//		permission letter group must be formatted as one
//		of the following:
//
//				"rwx"
//				"rw-"
//				"r--"
//				"---"
//				"--x"
//				"-wx"
//				"-w-"
//				"r-x"
//
//		This method will return an integer representing
//		the octal digits comprising this group code.
//
//		For example, groupStr="rwx" will return an
//		integer value of '7' which can be treated as
//		octal digit '7' for purposes of creating an
//		os.FileMode.
//
//		If 'permissionLetterGroup' does not match one
//		of the 3-character permission codes shown
//		above, an error will be returned.
//
//	userClass					string
//
//		This string must contain the user class name
//		associated with the permission letter code
//		submitted as input paramter,
//		'permissionLetterCode'.
//
//		'userClass' must contain one of the three
//		following user class names:
//
//			(1)	'owner'
//			(2)	'group'
//			(3)	'other'
//
//		If input parameter 'userClass' does not match one
//		of the three user classes listed above, an error
//		will be returned.
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
//	int
//
//		This method will return an integer value
//		representing the octal digits comprising the
//		permission letter code submitted as input
//		parameter, 'permissionLetterCode'.
//
//		For example, permissionLetterCode ="rwx" will
//		return an integer value of '7' which can be
//		treated as octal digit '7' for purposes of
//		creating an os.FileMode.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (filePermCfgElectron *filePermissionConfigElectron) convertGroupToDecimal(
	permissionLetterCode string,
	userClass string,
	errPrefDto *ePref.ErrPrefixDto) (
	int,
	error) {

	if filePermCfgElectron.lock == nil {
		filePermCfgElectron.lock = new(sync.Mutex)
	}

	filePermCfgElectron.lock.Lock()

	defer filePermCfgElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"filePermissionConfigElectron."+
			"testValidityOfFilePermissionConfig()",
		"")

	if err != nil {
		return 0, err
	}

	if len(permissionLetterCode) != 3 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'permissionLetterCode' must be exactly 3-characters in length.\n"+
			"This 'permissionLetterCode' is %v-characters in length.\n"+
			"permissionLetterCode='%v'\n"+
			"userClass='%v' ",
			ePrefix.String(),
			len(permissionLetterCode),
			permissionLetterCode,
			userClass)

		return -1, err
	}

	intVal := 0

	tstLetters := strings.ToLower(permissionLetterCode)

	switch tstLetters {
	case "rwx":
		intVal = 7
	case "rw-":
		intVal = 6
	case "r--":
		intVal = 4
	case "---":
		intVal = 0
	case "--x":
		intVal = 1
	case "-wx":
		intVal = 3
	case "-w-":
		intVal = 2
	case "r-x":
		intVal = 5
	default:
		err = fmt.Errorf(
			"%v\n"+
				"Error: Invalid 3-Letter 'permissionLetterCode' string!\n"+
				"permissionLetterCode = '%v'\n"+
				"userClass= '%v'\n"+
				"Converted 3-Letter Block='%v'",
			ePrefix.String(),
			permissionLetterCode,
			userClass,
			tstLetters)

		return -1, err
	}

	return intVal, err
}
