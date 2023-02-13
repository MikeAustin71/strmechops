package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"os"
	"sync"
)

type filePermissionConfigMechanics struct {
	lock *sync.Mutex
}

// getCompositePermissionMode
//
// Returns the os.FileMode from the internal data field,
// 'fPerm.fileMode', contained in the instance of
// FilePermissionConfig passed as input parameter
// 'fPerm'.
//
// 'fileMode' represents the complete, consolidated
// permission code. It therefore contains the two
// elements which make up a consolidated permission code:
//
//	(1)	Entry Type
//	(2) Permission Bits
//
// This method returns the complete permission code as a
// type 'os.FileMode'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fPerm						*FilePermissionConfig
//
//		A pointer to an instance of FilePermissionConfig.
//		The 'Entry Type' component of the os.FileMode
//		permissions value contained in this
//		FilePermissionConfig instance will be extracted
//		and returned to the calling function.
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
//		os.FileMode
//
//			Holds the consolidated file permission value
//			which consists of the two elements making up a
//			permission value:
//	     		(1)	Entry Type
//	     		(2) Permission Bits
//
//			The os.FileMode contained in input parameter
//			'fPerm' will be extracted and returned through
//			this parameter.
//
//		error
//
//			If this method completes successfully, the
//			returned error Type is set equal to 'nil'. If
//			errors are encountered during processing, the
//			returned error Type will encapsulate an error
//			message.
//
//			If an error message is returned, the text value
//			for input parameter 'errPrefDto' (error prefix)
//			will be prefixed or attached at the beginning of
//			the error message.
func (fPermConfigMech *filePermissionConfigMechanics) getCompositePermissionMode(
	fPerm *FilePermissionConfig,
	errPrefDto *ePref.ErrPrefixDto) (
	os.FileMode,
	error) {

	if fPermConfigMech.lock == nil {
		fPermConfigMech.lock = new(sync.Mutex)
	}

	fPermConfigMech.lock.Lock()

	defer fPermConfigMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"filePermissionConfigMechanics."+
			"getCompositePermissionMode()",
		"")

	if err != nil {
		return os.FileMode(0), err
	}

	_,
		err = new(filePermissionConfigElectron).
		testValidityOfFilePermissionConfig(
			fPerm,
			ePrefix)

	if err != nil {
		return os.FileMode(0), err
	}

	return fPerm.fileMode, nil
}

// getPermissionTextCode
//
// Returns the file mode permissions expressed as a text
// string. The returned string includes the full and
// complete 10-character permission code.
//
//	Example Return Values:
//	      -rwxrwxrwx
//	      -rw-rw-rw-
//	      drwxrwxrwx
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fPerm						*FilePermissionConfig
//
//		A pointer to an instance of FilePermissionConfig.
//		The file mode permissions expressed as a text
//		will be extracted from this instance and returned
//		to the calling function.
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
//	string
//
//		If this method completes successfully, this
//		parameter returns the file mode permissions
//		contained in 'fPerm' expressed as a text
//		string. The returned string includes the full
//		and complete 10-character permission code.
//
//			Example Return Values:
//	      		-rwxrwxrwx
//				-rw-rw-rw-
//				drwxrwxrwx
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
func (fPermConfigMech *filePermissionConfigMechanics) getPermissionTextCode(
	fPerm *FilePermissionConfig,
	errPrefDto *ePref.ErrPrefixDto) (
	string,
	error) {

	if fPermConfigMech.lock == nil {
		fPermConfigMech.lock = new(sync.Mutex)
	}

	fPermConfigMech.lock.Lock()

	defer fPermConfigMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"filePermissionConfigMechanics."+
			"getPermissionTextCode()",
		"")

	if err != nil {
		return "", err
	}

	_,
		err = new(filePermissionConfigElectron).
		testValidityOfFilePermissionConfig(
			fPerm,
			ePrefix)

	if err != nil {
		return "", err
	}

	return fPerm.fileMode.String(), nil
}
