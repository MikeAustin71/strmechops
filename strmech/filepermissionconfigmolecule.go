package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"os"
	"sync"
)

type filePermissionConfigMolecule struct {
	lock *sync.Mutex
}

// getEntryTypeComponent
//
// Returns the 'Entry Type' component of the os.FileMode
// permissions value contained in the
// FilePermissionConfig instance passed as input
// paramter 'fPerm'.
//
// The 'Entry Type' is the first character in a
// 10-character permissions text string. For the majority
// of applications, the leading character in a
// 10-character permissions text string is either a
// hyphen ('-') indicating the subject is a file - or -
// a 'd' indicating the subject is a directory.
//
// For a file, the File Mode Entry Type value is zero
// ('0').  For a directory, the File Mode Entry Type
// value is equal to 'os.ModDir'.
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
//	OsFilePermissionCode
//
//		The OsFilePermissionCode type is set to the value
//		of the os.FileMode constant representing the
//		Entry Type associated with the permission value
//		encapsulated by the FilePermissionConfig instance
//		passed as input parameter 'fPerm'.
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
func (fPermConfigMolecule *filePermissionConfigMolecule) getEntryTypeComponent(
	fPerm *FilePermissionConfig,
	errPrefDto *ePref.ErrPrefixDto) (
	OsFilePermissionCode,
	error) {

	if fPermConfigMolecule.lock == nil {
		fPermConfigMolecule.lock = new(sync.Mutex)
	}

	fPermConfigMolecule.lock.Lock()

	defer fPermConfigMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"filePermissionConfigMolecule."+
			"getEntryTypeComponent()",
		"")

	if err != nil {
		return OsFilePermissionCode(0), err
	}

	if fPerm == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fPerm' is a nil pointer!\n",
			ePrefix.String())

		return OsFilePermissionCode(0), err
	}

	_,
		err = new(filePermissionConfigElectron).
		testValidityOfFilePermissionConfig(
			fPerm,
			ePrefix)

	if err != nil {

		return OsFilePermissionCode(0), err
	}

	fMode := fPerm.fileMode &^ os.FileMode(0777)

	for idx := range mOsPermissionCodeToString {

		if fMode == idx {

			return OsFilePermissionCode(idx), nil

		}

	}

	err = fmt.Errorf("%v\n"+
		"The Entry Type for this FilePermissionConfig instance is INVALID!\n"+
		"No matching Os File Permission Code could be located.\n"+
		"FilePermissionConfig File Mode = %v\n",
		ePrefix.String(),
		fPerm.fileMode)

	return OsFilePermissionCode(0), err
}

// GetFileMode
//
// Returns a FileMode containing only the least
// significant 9-bits of the encapsulated FileMode
// representing the unix permission bits.
//
// If this value is converted to a permissions string,
// the actual string returned will contain 10-characters,
// have the first character (index=0) will always be a
// hyphen ("-"). The hyphen ("-") indicates a file,
// however it should be ignored in this case.
//
// The only valid a reliable unix permission bits are in
// the last 9-characters (string indexes 1-8). When
// evaluating permission bits returned by this method as
// permission strings always ignore the first character
// which will always be a hyphen ("-").
//
// To acquire the full and valid 10-digit permission
// string use method:
//
//	FilePermissionConfig.GetPermissionTextCode()
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
//	os.FileMode
//
//		This parameter returns a FileMode containing only
//		the least significant 9-bits of the encapsulated
//		FileMode representing the unix permission bits.
//
//		The os.FileMode value contained in input
//		parameter 'fPerm' will be used to produce this
//		permission bits value return as os.FileMode.
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
func (fPermConfigMolecule *filePermissionConfigMolecule) getFileMode(
	fPerm *FilePermissionConfig,
	errPrefDto *ePref.ErrPrefixDto) (
	os.FileMode,
	error) {

	if fPermConfigMolecule.lock == nil {
		fPermConfigMolecule.lock = new(sync.Mutex)
	}

	fPermConfigMolecule.lock.Lock()

	defer fPermConfigMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"filePermissionConfigMolecule."+
			"getFileMode()",
		"")

	if err != nil {
		return os.FileMode(0), err
	}

	if fPerm == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fPerm' is a nil pointer!\n",
			ePrefix.String())

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

	return fPerm.fileMode.Perm(), nil
}
