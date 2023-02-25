package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"os"
	"sync"
)

// FileAccessControl encapsulates the codes required the open files and
// configure file permissions. As such this type encapsulates types
// FilePermissionConfig and FileOpenConfig.
type FileAccessControl struct {
	isInitialized bool
	permissions   FilePermissionConfig
	fileOpenCodes FileOpenConfig

	lock *sync.Mutex
}

// NewInitialized
//
// Returns a new FileAccessControl instances with all
// File Open Codes and File Permission Codes initialized
// to 'None'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	FileAccessControl
//
//		If this method completes successfully, a new
//		instance of FileAccessControl will be returned.
//
//		This new FileAccessControl instances will be
//		configured with all File Open Codes and File
//		Permission Codes initialized to 'None'.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fAccess *FileAccessControl) NewInitialized(
	errorPrefix interface{}) (
	FileAccessControl,
	error) {

	if fAccess.lock == nil {
		fAccess.lock = new(sync.Mutex)
	}

	fAccess.lock.Lock()

	defer fAccess.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "FileAccessControl.NewInitialized()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {
		return FileAccessControl{}, err
	}

	openCodes, err := new(FileOpenConfig).New(
		ePrefix,
		FOpenType.TypeNone(),
		FOpenMode.ModeNone())

	if err != nil {
		return FileAccessControl{},
			fmt.Errorf("%v\n"+
				"Error returned by FileOpenConfig{}.New("+
				"FOpenType.TypeNone(), FOpenMode.ModeNone())\n"+
				"Error= \n%v\n",
				funcName,
				err.Error())
	}

	entryType, err :=
		OsFilePermissionCode(0).
			GetNewFromFileMode(OsFilePermCode.ModeNone())

	if err != nil {
		return FileAccessControl{},
			fmt.Errorf("%v\n"+
				"Error returned by OsFilePermCode.GetNewFromFileMode("+
				"OsFilePermCode.ModeNone()).\n"+
				"Error= \n%v\n ",
				ePrefix.String(),
				err.Error())
	}

	permissions, err := new(FilePermissionConfig).
		NewByComponents(
			entryType,
			"---------",
			ePrefix)

	if err != nil {
		return FileAccessControl{},
			fmt.Errorf("%v\n"+
				"Error returned by FilePermissionConfig{}.NewByComponents("+
				"entryType, \"---------\")\n"+
				"entryType='OsFilePermCode.ModeNone()'\n"+
				"Error= \n%v\n",
				funcName,
				err.Error())
	}

	fA2 := FileAccessControl{}

	fA2.fileOpenCodes = openCodes.CopyOut()

	fA2.permissions = permissions.CopyOut()

	fA2.isInitialized = true

	return fA2, err
}

// New
//
// Creates and returns a new instance of type
// FileAccessControl.
func (fAccess *FileAccessControl) New(
	openCodes FileOpenConfig,
	permissions FilePermissionConfig) (
	FileAccessControl,
	error) {

	ePrefix := "FileAccessControl.New() "

	err := openCodes.IsValidInstanceError(ePrefix)

	if err != nil {
		return FileAccessControl{},
			fmt.Errorf(ePrefix+"Input parameter 'openCodes' is INVALID! "+
				"Error='%v' ", err.Error())
	}

	err = permissions.IsValid(ePrefix)

	if err != nil {
		return FileAccessControl{},
			fmt.Errorf(ePrefix+"Input parameter 'permissions' is INVALID! "+
				"Error='%v' ", err.Error())
	}

	fA2 := FileAccessControl{}

	fA2.fileOpenCodes = openCodes.CopyOut()

	fA2.permissions = permissions.CopyOut()

	fA2.isInitialized = true

	return fA2, nil
}

// NewReadWriteAccess
//
// Returns a FileAccessControl instance configured for
// Read/Write access.
func (fAccess FileAccessControl) NewReadWriteAccess() (
	FileAccessControl,
	error) {

	ePrefix := "FileAccessControl.NewReadWriteAccess() "

	fileOpenCfg, err :=
		new(FileOpenConfig).
			New(ePrefix,
				FOpenType.TypeReadWrite(),
				FOpenMode.ModeNone())

	if err != nil {
		return FileAccessControl{}, fmt.Errorf(ePrefix+"%v\n", err.Error())
	}

	filePermCfg, err := new(FilePermissionConfig).
		New("-rw-rw-rw-", ePrefix)

	if err != nil {
		return FileAccessControl{}, fmt.Errorf(ePrefix+"%v\n", err.Error())
	}

	fileAccessCfg, err := FileAccessControl{}.New(fileOpenCfg, filePermCfg)

	if err != nil {
		return FileAccessControl{}, fmt.Errorf(ePrefix+"%v\n", err.Error())
	}

	return fileAccessCfg, nil
}

// NewReadWriteCreateTruncateAccess - Returns a FileAccessControl instance
// configured for Read, Write, Create and Truncate access.
func (fAccess FileAccessControl) NewReadWriteCreateTruncateAccess() (FileAccessControl, error) {

	ePrefix := "FileAccessControl.NewReadWriteCreateTruncateAccess() "

	//  OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
	fOpenCfg, err := new(FileOpenConfig).New(
		ePrefix,
		FOpenType.TypeReadWrite(),
		FOpenMode.ModeCreate(),
		FOpenMode.ModeTruncate())

	if err != nil {
		return FileAccessControl{}, fmt.Errorf(ePrefix+"%v\n", err.Error())
	}

	fPermCfg, err := new(FilePermissionConfig).
		New("-rw-rw-rw-", ePrefix)

	if err != nil {
		return FileAccessControl{}, fmt.Errorf(ePrefix+"%v\n", err.Error())
	}

	fileAccessCfg, err := FileAccessControl{}.New(fOpenCfg, fPermCfg)

	if err != nil {
		return FileAccessControl{}, fmt.Errorf(ePrefix+"%v\n", err.Error())
	}

	return fileAccessCfg, nil
}

// NewReadOnlyAccess - Returns a FileAccessControl instance configured for
// Read-Only access.
func (fAccess FileAccessControl) NewReadOnlyAccess() (FileAccessControl, error) {

	ePrefix := "FileAccessControl.NewReadOnlyAccess() "

	fileOpenCfg, err := new(FileOpenConfig).
		New(ePrefix,
			FOpenType.TypeReadOnly(),
			FOpenMode.ModeNone())

	if err != nil {
		return FileAccessControl{}, fmt.Errorf(ePrefix+"%v\n", err.Error())
	}

	filePermCfg, err := new(FilePermissionConfig).
		New("-r--r--r--", ePrefix)

	if err != nil {
		return FileAccessControl{}, fmt.Errorf(ePrefix+"%v\n", err.Error())
	}

	fileAccessCfg, err := FileAccessControl{}.New(fileOpenCfg, filePermCfg)

	if err != nil {
		return FileAccessControl{}, fmt.Errorf(ePrefix+"%v\n", err.Error())
	}

	return fileAccessCfg, nil
}

// NewWriteOnlyAccess - Returns a FileAccessControl instance configured for
// Write-Only access.
func (fAccess FileAccessControl) NewWriteOnlyAccess() (FileAccessControl, error) {

	ePrefix := "FileAccessControl.NewWriteOnlyAccess() "

	fileOpenCfg, err :=
		new(FileOpenConfig).
			New(ePrefix,
				FOpenType.TypeWriteOnly(),
				FOpenMode.ModeNone())

	if err != nil {
		return FileAccessControl{}, fmt.Errorf(ePrefix+"%v\n", err.Error())
	}

	filePermCfg, err := new(FilePermissionConfig).
		New("--w--w--w-", ePrefix)

	if err != nil {
		return FileAccessControl{}, fmt.Errorf(ePrefix+"%v\n", err.Error())
	}

	fileAccessCtrl, err := FileAccessControl{}.New(fileOpenCfg, filePermCfg)

	if err != nil {
		return FileAccessControl{}, fmt.Errorf(ePrefix+"%v\n", err.Error())
	}

	return fileAccessCtrl, nil
}

// NewWriteOnlyAppendAccess - Returns a FileAccessControl instance configured for
// Write/Only - Append access.
func (fAccess FileAccessControl) NewWriteOnlyAppendAccess() (FileAccessControl, error) {

	ePrefix := "FileAccessControl.NewWriteOnlyAccess() "

	fileOpenCfg, err := new(FileOpenConfig).
		New(ePrefix,
			FOpenType.TypeWriteOnly(),
			FOpenMode.ModeAppend())

	if err != nil {
		return FileAccessControl{}, fmt.Errorf(ePrefix+"%v\n", err.Error())
	}

	filePermCfg, err := new(FilePermissionConfig).
		New("--w--w--w-", ePrefix)

	if err != nil {
		return FileAccessControl{}, fmt.Errorf(ePrefix+"%v\n", err.Error())
	}

	fileAccessCfg, err := FileAccessControl{}.New(fileOpenCfg, filePermCfg)

	if err != nil {
		return FileAccessControl{}, fmt.Errorf(ePrefix+"%v\n", err.Error())
	}

	return fileAccessCfg, nil
}

// NewWriteOnlyTruncateAccess - Returns a FileAccessControl instance configured for
// Write/Only - Truncate access.
//
// If the file previously exists, it will be truncated before the writing operation
// commences.
func (fAccess FileAccessControl) NewWriteOnlyTruncateAccess() (FileAccessControl, error) {

	ePrefix := "FileAccessControl.NewWriteOnlyTruncateAccess() "

	fileOpenCfg, err :=
		new(FileOpenConfig).New(
			ePrefix,
			FOpenType.TypeWriteOnly(),
			FOpenMode.ModeTruncate())

	if err != nil {
		return FileAccessControl{}, fmt.Errorf(ePrefix+"%v\n", err.Error())
	}

	filePermCfg, err := new(FilePermissionConfig).
		New("--w--w--w-", ePrefix)

	if err != nil {
		return FileAccessControl{}, fmt.Errorf(ePrefix+"%v\n", err.Error())
	}

	fileAccessCfg, err := FileAccessControl{}.New(fileOpenCfg, filePermCfg)

	if err != nil {
		return FileAccessControl{}, fmt.Errorf(ePrefix+"%v\n", err.Error())
	}

	return fileAccessCfg, nil
}

// CopyIn - Receives a FileAccessControl instance and copies all the data
// fields to the current FileAccessControl instance. When complete, both
// the incoming and current FileAccessControl instances will be identical.
//
// The type of copy operation performed is a 'deep copy'.
func (fAccess *FileAccessControl) CopyIn(fA2 *FileAccessControl) {

	fAccess.isInitialized = fA2.isInitialized

	fAccess.fileOpenCodes.CopyIn(&fA2.fileOpenCodes)

	fAccess.permissions.CopyIn(&fA2.permissions)

}

// CopyOut - Creates and returns a deep copy of the current
// FileAccessControl instance.
func (fAccess *FileAccessControl) CopyOut() FileAccessControl {

	fA2 := FileAccessControl{}

	fA2.isInitialized = fAccess.isInitialized
	fA2.fileOpenCodes = fAccess.fileOpenCodes.CopyOut()
	fA2.permissions = fAccess.permissions.CopyOut()

	return fA2
}

// Empty - ReInitializes the current FileAccessControl instance to
// empty or zero values.
func (fAccess *FileAccessControl) Empty() {
	fAccess.fileOpenCodes.Empty()
	fAccess.permissions.Empty()
	fAccess.isInitialized = false
}

// Equal - Returns 'true' if the incoming FileAccessControl instance
// is equal in all respects to the current FileAccessControl instance.
func (fAccess *FileAccessControl) Equal(fA2 *FileAccessControl) bool {

	if fAccess.isInitialized != fA2.isInitialized {
		return false
	}

	if !fAccess.fileOpenCodes.Equal(&fA2.fileOpenCodes) {
		return false
	}

	if !fAccess.permissions.Equal(&fA2.permissions) {
		return false
	}

	return true
}

// GetCompositeFileOpenCode - Returns the composite 'file open' code. This code
// is generated by combining the single FileOpenType value and zero
// or more FileOpenMode values.
func (fAccess *FileAccessControl) GetCompositeFileOpenCode() (int, error) {

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"FileAccessControl."+
			"GetCompositeFileOpenCode()",
		"")

	_,
		err = new(fileAccessControlElectron).
		testValidityOfFileAccessControl(
			fAccess,
			ePrefix)

	if err != nil {
		return 0, err
	}

	fileOpenCodes, err :=
		fAccess.fileOpenCodes.
			GetCompositeFileOpenCode(ePrefix)

	if err != nil {
		return 0, err
	}

	return fileOpenCodes, nil
}

// GetCompositePermissionMode - Returns the complete permission code as a type
// os.FileMode.
func (fAccess *FileAccessControl) GetCompositePermissionMode() (os.FileMode, error) {

	if fAccess.lock == nil {
		fAccess.lock = new(sync.Mutex)
	}

	fAccess.lock.Lock()

	defer fAccess.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"FileAccessControl."+
			"GetCompositePermissionMode()",
		"")

	_,
		err = new(fileAccessControlElectron).
		testValidityOfFileAccessControl(
			fAccess,
			ePrefix)

	if err != nil {
		return os.FileMode(9999), err
	}

	var permissionCode os.FileMode

	permissionCode,
		err = fAccess.permissions.
		GetCompositePermissionMode(ePrefix)

	if err != nil {
		return os.FileMode(9999), err
	}

	return permissionCode, err
}

// GetCompositePermissionModeText - Returns the composite permission file mode
// numerical value expressed as text.
//
//	Example:
//
//	      -rw-rw-rw- = returned value "0666"
//	      drwxrwxrwx = returned value "020000000777"
func (fAccess *FileAccessControl) GetCompositePermissionModeText() string {

	if fAccess.lock == nil {
		fAccess.lock = new(sync.Mutex)
	}

	fAccess.lock.Lock()

	defer fAccess.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"FileAccessControl."+
			"GetCompositePermissionModeText()",
		"")

	_,
		err = new(fileAccessControlElectron).
		testValidityOfFileAccessControl(
			fAccess,
			ePrefix)

	if err != nil {
		return err.Error()
	}

	var permissionModeText string

	permissionModeText,
		err = fAccess.permissions.
		GetPermissionFileModeValueText(
			ePrefix.XCpy(
				"fAccess.permissions->"))

	if err != nil {
		permissionModeText = err.Error()
	}

	return permissionModeText
}

// GetFileOpenAndPermissionCodes - Returns both the complete File Open Code
// and complete Permission code.
func (fAccess *FileAccessControl) GetFileOpenAndPermissionCodes() (
	int,
	os.FileMode,
	error) {

	if fAccess.lock == nil {
		fAccess.lock = new(sync.Mutex)
	}

	fAccess.lock.Lock()

	defer fAccess.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"FileAccessControl."+
			"GetFileOpenAndPermissionCodes()",
		"")

	if err != nil {
		return -1, os.FileMode(9999), err
	}

	_,
		err = new(fileAccessControlElectron).
		testValidityOfFileAccessControl(
			fAccess,
			ePrefix)

	if err != nil {
		return -1, os.FileMode(9999), err
	}

	fileOpenCode, err :=
		fAccess.fileOpenCodes.
			GetCompositeFileOpenCode(
				ePrefix.XCpy(
					"fAccess.fileOpenCodes->"))

	if err != nil {
		return -1, os.FileMode(9999), err
	}

	permissionCode, err := fAccess.permissions.
		GetCompositePermissionMode(
			ePrefix.XCpy(
				"fAccess.permissions->"))

	if err != nil {
		return -1, os.FileMode(9999), err
	}

	return fileOpenCode, permissionCode, err
}

// GetFileOpenConfig - Returns a deep copy of the FileOpenConfig type
// encapsulated by the current FileAccessControl instance.
func (fAccess *FileAccessControl) GetFileOpenConfig() (FileOpenConfig, error) {

	if fAccess.lock == nil {
		fAccess.lock = new(sync.Mutex)
	}

	fAccess.lock.Lock()

	defer fAccess.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"FileAccessControl."+
			"GetFileOpenConfig()",
		"")

	if err != nil {
		return FileOpenConfig{}, err
	}

	_,
		err = new(fileAccessControlElectron).
		testValidityOfFileAccessControl(
			fAccess,
			ePrefix)

	if err != nil {
		return FileOpenConfig{}, err
	}

	return fAccess.fileOpenCodes.CopyOut(), err
}

// GetFileOpenType - Returns the File Open Type associated with the
// FileOpenConfig type stored as 'FileAccessControl.fileOpenCodes'.
func (fAccess *FileAccessControl) GetFileOpenType() (FileOpenType, error) {

	if fAccess.lock == nil {
		fAccess.lock = new(sync.Mutex)
	}

	fAccess.lock.Lock()

	defer fAccess.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"FileAccessControl."+
			"GetFileOpenConfig()",
		"")

	if err != nil {
		return FileOpenType(99999), err
	}

	_,
		err = new(fileAccessControlElectron).
		testValidityOfFileAccessControl(
			fAccess,
			ePrefix)

	if err != nil {
		return FileOpenType(99999), err
	}

	return fAccess.fileOpenCodes.GetFileOpenType(), err
}

// GetFilePermissionConfig - Returns a deep copy of the FilePermissionConfig type
// encapsulated by the current FileAccessControl instance.
func (fAccess *FileAccessControl) GetFilePermissionConfig() (FilePermissionConfig, error) {

	if fAccess.lock == nil {
		fAccess.lock = new(sync.Mutex)
	}

	fAccess.lock.Lock()

	defer fAccess.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"FileAccessControl."+
			"GetFilePermissionConfig()",
		"")

	if err != nil {
		return FilePermissionConfig{}, err
	}

	_,
		err = new(fileAccessControlElectron).
		testValidityOfFileAccessControl(
			fAccess,
			ePrefix)

	if err != nil {
		return FilePermissionConfig{}, err
	}

	return fAccess.permissions.CopyOut(), nil
}

// GetFilePermissionTextCode - Returns the file mode permissions expressed as
// a text string. The returned string includes the full and complete
// 10-character permission code.
//
//	Example Return Values:
//	      -rwxrwxrwx
//	      -rw-rw-rw-
//	      drwxrwxrwx
func (fAccess *FileAccessControl) GetFilePermissionTextCode() (string,
	error) {

	if fAccess.lock == nil {
		fAccess.lock = new(sync.Mutex)
	}

	fAccess.lock.Lock()

	defer fAccess.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"FileAccessControl."+
			"GetFilePermissionTextCode()",
		"")

	if err != nil {
		return "", err
	}

	_,
		err = new(fileAccessControlElectron).
		testValidityOfFileAccessControl(
			fAccess,
			ePrefix)

	if err != nil {
		return "", err
	}

	permTxtCode, err := fAccess.permissions.
		GetPermissionTextCode(
			ePrefix.XCpy(
				"fAccess.permissions->"))

	if err != nil {
		return "", err
	}

	return permTxtCode, err
}

// IsValidInstanceError
//
// If the current FileAccessControl instance is valid and
// properly initialized, this method returns an error
// value of 'nil'.
//
// If the current FileAccessControl instance is invalid,
// this method returns an error encapsulating an
// appropriate error message describing the cause of the
// error.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If any of the internal member data variables
//		contained in the current instance of
//		FileAccessControl are found to be invalid, this
//		method will return an error configured with an
//		appropriate message identifying the invalid
//		member data variable.
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
//		passed by input parameter, 'errorPrefix'. The
//		'errorPrefix' text will be attached to the
//		beginning of the error message.
func (fAccess *FileAccessControl) IsValidInstanceError(
	errorPrefix interface{}) error {

	if fAccess.lock == nil {
		fAccess.lock = new(sync.Mutex)
	}

	fAccess.lock.Lock()

	defer fAccess.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileAccessControl."+
			"IsValidInstanceError()",
		"")

	if err != nil {
		return err
	}

	_,
		err = new(fileAccessControlElectron).
		testValidityOfFileAccessControl(
			fAccess,
			ePrefix)

	return err
}

// SetFileOpenCodes - Assigns 'fileOpenCodes' to internal member variable,
// FileAccessControl.fileOpenCodes
func (fAccess *FileAccessControl) SetFileOpenCodes(fileOpenCodes FileOpenConfig) error {

	ePrefix := "FileAccessControl.SetFileOpenCodes() "

	err := fileOpenCodes.IsValidInstanceError(ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+"INVALID 'fileOpenCodes'! - %v", err.Error())
	}

	fAccess.fileOpenCodes = fileOpenCodes.CopyOut()

	err = fAccess.permissions.IsValid(ePrefix)

	if err == nil {

		fAccess.isInitialized = true

	}

	return nil
}

// SetFilePermissionCodes - Assigns 'filePermissions' to internal
// member variable FileAccessControl.permissions.
func (fAccess *FileAccessControl) SetFilePermissionCodes(
	filePermissions FilePermissionConfig) error {

	ePrefix := "FileAccessControl.SetFilePermissionCodes() "

	err := filePermissions.IsValid(
		ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+"Error: 'filePermissions' INVALID! - %v",
			err.Error())
	}

	fAccess.permissions = filePermissions.CopyOut()

	err = fAccess.fileOpenCodes.IsValidInstanceError(ePrefix)

	if err == nil {
		fAccess.isInitialized = true
	}

	return nil
}
