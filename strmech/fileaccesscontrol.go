package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"os"
	"sync"
)

// FileAccessControl
//
// Type FileAccessControl encapsulates the codes required
// to open files and configure file access permissions. As
// such this type encapsulates types FilePermissionConfig
// and FileOpenConfig.
//
// ----------------------------------------------------------------
//
// # Background
//
// The FileAccessControl type is used when opening files
// for read and write operations. The FileAccessControl
// type encapsulates File Permission Codes, File Open
// Type and File Open Modes.
//
// To open a file, three components may be required,
// File Permission Codes, File Open Type and File
// Open Modes:
//
//  1. File Permission Codes
//
//     The File Permission Codes is a 10-character
//     string containing the read, write and execute
//     permissions for the three groups or user
//     classes:
//
//     (1)	'Owner/User'
//
//     (2)	'Group'
//
//     (3)	'Other'
//
//     This 10-character string will be used to
//     configure the internal File Permission data
//     field for a configured instance of
//     FileAccessControl.
//
//     The File Permission string must conform to the
//     symbolic notation options shown below. Failure
//     to comply with this requirement will generate
//     an error. As indicated, the File Permission
//     string must consist of 10-characters.
//
//     The first character in the File Permission
//     string may be dash ('-') specifying a file or a
//     'd' specifying a directory.
//
//     The remaining nine characters in the File
//     Permission string represent unix permission
//     bits and consist of three group fields each
//     containing 3-characters. Each character in
//     the three group fields may consist of 'r'
//     (Read-Permission), 'w' (Write-Permission),
//     'x' (Execute-Permission) or dash ('-')
//     signaling no permission or no access allowed.
//     A typical File Permission string authorizing
//     permission for full access to a file would be
//     styled as:
//
//     Example: "-rwxrwxrwx"
//
//     Groups:	-	Owner/User, Group, Other
//
//     From left to right
//     First Characters is Entry Type
//     -----------------------------------------------------
//     First Char index 0	=	"-"   Designates a file
//
//     First Char index 0	=	"d"   Designates a directory
//     -----------------------------------------------------
//
//     Char indexes 1-3	=	Owner "rwx" Authorizing 'Read',
//     Write' & Execute Permissions
//     for 'Owner'
//
//     Char indexes 4-6	= 	Group "rwx" Authorizing 'Read', 'Write' & Execute
//     Permissions for 'Group'
//
//     Char indexes 7-9	=	Other "rwx" Authorizing 'Read', 'Write' & Execute
//     Permissions for 'Other'
//
//     The Symbolic notation provided by input parameter
//     'filePermissionStr' MUST conform to the options
//     presented below. The first character or 'Entry Type'
//     is listed as "-". However, in practice, the caller
//     may set the first character as either a "-",
//     specifying a file, or a "d", specifying a directory.
//     No other first character types are currently
//     supported.
//
//     Three SymbolicGroups:
//
//     The three group types are: User/Owners, Groups & Others.
//
//     Directory Permissions:
//
//     -----------------------------------------------------
//     Directory Mode String Permission Codes
//     -----------------------------------------------------
//     Directory
//     10-Character
//     File Permission
//     String
//     Symbolic		  	Directory Access
//     Format	   		Permission Descriptions
//     ----------------------------------------------------
//
//     d---------		no permissions
//     drwx------		read, write, & execute only for owner
//     drwxrwx---		read, write, & execute for owner and group
//     drwxrwxrwx		read, write, & execute for owner, group and others
//     d--x--x--x		execute
//     d-w--w--w-		write
//     d-wx-wx-wx		write & execute
//     dr--r--r--		read
//     dr-xr-xr-x		read & execute
//     drw-rw-rw-		read & write
//     drwxr-----		Owner can read, write, & execute. Group can only read;
//     others have no permissions
//
//     Note: drwxrwxrwx - identifies permissions for directory
//
//     File Permissions:
//
//     -----------------------------------------------------
//     File Mode String Permission Codes
//     -----------------------------------------------------
//
//     10-Character
//     File
//     Permission
//     String
//     Symbolic	Octal		Permission
//     Format		Notation	Descriptions
//     ------------------------------------------------------------
//
//     ----------	  0000		no permissions
//
//     -rwx------	  0700		read, write, & execute only for owner
//
//     -rwxrwx---	  0770		read, write, & execute for
//     owner and group
//
//     -rwxrwxrwx	  0777		read, write, & execute for owner,
//     group and others
//
//     ---x--x--x	  0111		execute
//
//     --w--w--w-	  0222		write
//
//     --wx-wx-wx	  0333		write & execute
//
//     -r--r--r--	  0444		read
//
//     -r-xr-xr-x	  0555		read & execute
//
//     -rw-rw-rw-	  0666		read & write
//
//     -rwxr-----	  0740		Owner can read, write, &
//     execute. Group can only
//     read; others have no
//     permissions
//
//  2. File Open Type
//
//     A file open type is an enumeration specifying
//     the manner in which  the file will be opened.
//     In order to open a file, exactly one of the
//     following File Open Codes MUST be specified:
//
//     FileOpenType(0).TypeReadOnly()
//     FileOpenType(0).TypeWriteOnly()
//     FileOpenType(0).TypeReadWrite()
//
//     -- AND --
//
//  3. File Open Mode
//
//     In addition to a File Open Type, a File Open
//     Mode may be specified. Zero or more of the
//     following File Open Modes may optionally be
//     specified to achieve granular control over
//     file open behavior.
//
//     FileOpenMode(0).ModeAppend()
//     FileOpenMode(0).ModeCreate()
//     FileOpenMode(0).ModeExclusive()
//     FileOpenMode(0).ModeSync()
//     FileOpenMode(0).ModeTruncate()
type FileAccessControl struct {
	isInitialized bool
	permissions   FilePermissionConfig
	fileOpenCodes FileOpenConfig

	lock *sync.Mutex
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
func (fAccess *FileAccessControl) GetCompositeFileOpenCode(
	errorPrefix interface{}) (int, error) {

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
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
func (fAccess *FileAccessControl) GetCompositePermissionMode(
	errorPrefix interface{}) (
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
		errorPrefix,
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
func (fAccess *FileAccessControl) GetFileOpenAndPermissionCodes(
	errorPrefix interface{}) (
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
		errorPrefix,
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
func (fAccess *FileAccessControl) GetFileOpenConfig(
	errorPrefix interface{}) (
	FileOpenConfig,
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
		errorPrefix,
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
func (fAccess *FileAccessControl) GetFileOpenType(
	errorPrefix interface{}) (
	FileOpenType,
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
		errorPrefix,
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
func (fAccess *FileAccessControl) GetFilePermissionConfig(
	errorPrefix interface{}) (FilePermissionConfig, error) {

	if fAccess.lock == nil {
		fAccess.lock = new(sync.Mutex)
	}

	fAccess.lock.Lock()

	defer fAccess.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
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
func (fAccess *FileAccessControl) GetFilePermissionTextCode(
	errorPrefix interface{}) (string,
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
		errorPrefix,
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
//		This new FileAccessControl instance will be
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

	newFAccessCtrl := FileAccessControl{}

	err = new(fileAccessControlMechanics).
		setInitializeNewFileAccessCtrl(
			&newFAccessCtrl,
			ePrefix.XCpy("fA2<-"))

	return newFAccessCtrl, err
}

// NewComponents
//
// Creates and returns a new instance of type
// FileAccessControl configured with the values provided
// by input parameters for file permissions, file open
// type and file open modes.
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
//	filePermissionStr			string
//
//		'filePermissionStr' is a 10-character string
//		containing the read, write and execute
//		permissions for the three groups or user
//		classes:
//
//			(1)	'Owner/User'
//
//			(2)	'Group'
//
//			(3)	'Other'
//
//		This 10-character string will be used to
//		configure the internal File Permission data field
//		for the new returned instance of FilePermissionConfig.
//
//		'filePermissionStr' must conform to the symbolic
//		notation options shown below. Failure to comply
//		with this requirement will generate an error. As
//		indicated, 'filePermissionStr' must consist of
//		10-characters.
//
//		The first character in 'filePermissionStr' may be
//		'-' specifying a fle or 'd' specifying a
//		directory.
//
//		The remaining nine characters in the
//		File Permission String represent unix permission
//		bits and consist of three group fields each
//		containing 3-characters. Each character in the
//		three group fields may consist of 'r'
//		(Read-Permission), 'w' (Write-Permission), 'x'
//		(Execute-Permission) or dash ('-') signaling no
//		permission or no access allowed. A typical
//		File Permission String authorizing permission
//		for full access to a file would be styled as:
//
//			Example: "-rwxrwxrwx"
//
//		Groups:	-	Owner/User, Group, Other
//
//		From left to right
//		First Characters is Entry Type
//		-----------------------------------------------------
//		First Char index 0	=	"-"   Designates a file
//
//		First Char index 0	=	"d"   Designates a directory
//		-----------------------------------------------------
//
//		Char indexes 1-3	=	Owner "rwx" Authorizing 'Read',
//								Write' & Execute Permissions
//								for 'Owner'
//
//		Char indexes 4-6	= 	Group "rwx" Authorizing 'Read', 'Write' & Execute
//								Permissions for 'Group'
//
//		Char indexes 7-9	=	Other "rwx" Authorizing 'Read', 'Write' & Execute
//								Permissions for 'Other'
//
//		The Symbolic notation provided by input parameter
//		'filePermissionStr' MUST conform to the options
//		presented below. The first character or 'Entry Type'
//		is listed as "-". However, in practice, the caller
//		may set the first character as either a "-",
//		specifying a file, or a "d", specifying a directory.
//		No other first character types are currently
//		supported.
//
//		Three SymbolicGroups:
//
//			The three group types are: User/Owners, Groups & Others.
//
//		Directory Permissions:
//
//			-----------------------------------------------------
//			        Directory Mode String Permission Codes
//			-----------------------------------------------------
//				Directory
//				10-Character
//				File Permission
//				String
//				Symbolic		  	Directory Access
//				Format	   		Permission Descriptions
//			----------------------------------------------------
//
//				d---------		no permissions
//				drwx------		read, write, & execute only for owner
//				drwxrwx---		read, write, & execute for owner and group
//				drwxrwxrwx		read, write, & execute for owner, group and others
//				d--x--x--x		execute
//				d-w--w--w-		write
//				d-wx-wx-wx		write & execute
//				dr--r--r--		read
//				dr-xr-xr-x		read & execute
//				drw-rw-rw-		read & write
//				drwxr-----		Owner can read, write, & execute. Group can only read;
//				                others have no permissions
//
//				Note: drwxrwxrwx - identifies permissions for directory
//
//		File Permissions:
//
//			-----------------------------------------------------
//			       File Mode String Permission Codes
//			-----------------------------------------------------
//
//			10-Character
//		       File
//			Permission
//			  String
//			 Symbolic	 Octal		File Access
//			  Format	Notation  Permission Descriptions
//			------------------------------------------------------------
//
//			----------	  0000		no permissions
//
//			-rwx------	  0700		read, write, & execute only for owner
//
//			-rwxrwx---	  0770		read, write, & execute for owner and
//						  				group
//
//			-rwxrwxrwx	  0777		read, write, & execute for owner,
//						  				group and others
//
//			---x--x--x	  0111		execute
//
//			--w--w--w-	  0222		write
//
//			--wx-wx-wx	  0333		write & execute
//
//			-r--r--r--	  0444		read
//
//			-r-xr-xr-x	  0555		read & execute
//
//			-rw-rw-rw-	  0666		read & write
//
//			-rwxr-----	  0740		Owner can read, write, & execute.
//									Group can only read; others
//									have no permissions
//
//	fOpenType					FileOpenType
//
//		The FileOpenType used to open a file.
//		FileOpenType is an enumeration. Valid options are
//		listed below using shorthand notation:
//
//			FOpenType.TypeReadOnly()
//			FOpenType.TypeWriteOnly()
//			FOpenType.TypeReadWrite()
//
//	fOpenModes					...FileOpenMode
//
//		As a golang variadic parameter, 'fOpenModes'
//		parameter accepts a variable number of arguments.
//
//		'fOpenModes' therefore transmits none, one or
//		more than one, File Open Modes used in the file
//		opening procedure.
//
//		Configure 'fOpenModes' with Zero or more
//		FileOpenMode concrete instances which will be
//		or'd with the input parameter 'fOpenType' in
//		order to generate the composite 'file open' code
//		used to open the target file.
//
//		If no File Open Modes are required, the caller
//		should pass nothing (blank/empty) for this
//		parameter.
//
//		FileOpenMode is an enumeration. Valid options are
//		listed below using the abbreviated notation:
//
//			FOpenMode.ModeAppend()
//			FOpenMode.ModeCreate()
//			FOpenMode.ModeExclusive()
//			FOpenMode.ModeSync()
//			FOpenMode.ModeTruncate()
//
//		Again, if no File Open Mode specifications are
//		required, leave this parameter blank/empty.
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
//		This new FileAccessControl instance will be
//		configured in accordance with the File
//		Permissions, Open Type and File Modes passed as
//		input parameters.
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
func (fAccess *FileAccessControl) NewComponents(
	errorPrefix interface{},
	filePermissionStr string,
	fOpenType FileOpenType,
	fOpenModes ...FileOpenMode) (
	FileAccessControl,
	error) {

	if fAccess.lock == nil {
		fAccess.lock = new(sync.Mutex)
	}

	fAccess.lock.Lock()

	defer fAccess.lock.Unlock()

	newFAccessCtrl := FileAccessControl{}

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileAccessControl."+
			"NewComponents()",
		"")

	if err != nil {

		return newFAccessCtrl, err
	}

	var filePermissionCfg FilePermissionConfig

	filePermissionCfg,
		err = new(FilePermissionConfig).New(
		filePermissionStr,
		ePrefix.XCpy("filePermissionCfg<-"))

	if err != nil {

		return newFAccessCtrl, err
	}

	var fileOpenCfg FileOpenConfig

	fileOpenCfg,
		err = new(FileOpenConfig).New(
		ePrefix.XCpy("fileOpenCfg<-"),
		fOpenType,
		fOpenModes...)

	if err != nil {

		return newFAccessCtrl, err
	}

	err = new(fileAccessControlMechanics).
		setFileAccessControl(
			&newFAccessCtrl,
			fileOpenCfg,
			filePermissionCfg,
			ePrefix.XCpy("newFAccessCtrl<-"))

	return newFAccessCtrl, err
}

// NewElements
//
// Creates and returns a new instance of type
// FileAccessControl configured with the values provided
// by input parameters 'openCodes' and 'permissions'.
//
// The FileAccessControl type is used when opening files
// for read and write operations.
//
// To open a file, two components are required:
//
//  1. A FileOpenType - Input parameter FileOpenConfig
//     In order to open a file, exactly one of the
//     following File Open Codes MUST be specified:
//
//     FileOpenType(0).TypeReadOnly()
//     FileOpenType(0).TypeWriteOnly()
//     FileOpenType(0).TypeReadWrite()
//
//     -- AND --
//
//  2. A FileOpenMode - Input parameter FilePermissionConfig
//
//     In addition to a 'FileOpenType', a File Open Mode
//     is also required. This code is also referred to as
//     'permissions'. Zero or more of the following File
//     Open Mode codes may optionally be specified to
//     better control file open behavior.
//
//     FileOpenMode(0).ModeAppend()
//     FileOpenMode(0).ModeCreate()
//     FileOpenMode(0).ModeExclusive()
//     FileOpenMode(0).ModeSync()
//     FileOpenMode(0).ModeTruncate()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	openCodes					FileOpenConfig
//
//		This parameter encapsulates the File Open
//		parameters which will be used to open subject
//		file. 'openCodes' are also referred to as the
//		File Open Type. For an explanation of File Open
//		parameters, see the source code documentation
//		for method FileOpenConfig.New().
//
//	permissions					FilePermissionConfig
//
//		This parameter encapsulates the File Permission
//		parameters which will be used to open the subject
//		file. 'permissions' is also referred to as the
//		File Open Mode. For an explanation of File
//		Permission parameters, see method FilePermissionConfig.New().
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
//		This new FileAccessControl instance will be
//		configured with File Open Codes and File
//		Permission Codes provided by input parameters
//		'openCodes' and 'permissions'.
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
func (fAccess *FileAccessControl) NewElements(
	openCodes FileOpenConfig,
	permissions FilePermissionConfig,
	errorPrefix interface{}) (
	FileAccessControl,
	error) {

	if fAccess.lock == nil {
		fAccess.lock = new(sync.Mutex)
	}

	fAccess.lock.Lock()

	defer fAccess.lock.Unlock()

	newFAccessCtrl := FileAccessControl{}

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileAccessControl."+
			"NewElements()",
		"")

	if err != nil {
		return newFAccessCtrl, err
	}

	err = new(fileAccessControlMechanics).
		setFileAccessControl(
			&newFAccessCtrl,
			openCodes,
			permissions,
			ePrefix.XCpy("newFAccessCtrl<-"))

	return newFAccessCtrl, err
}

// NewReadWriteAccess
//
// Returns a FileAccessControl instance configured for
// Read/Write access.
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
//		This new FileAccessControl instance will be
//		configured for Read-Write file access.
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
func (fAccess *FileAccessControl) NewReadWriteAccess(
	errorPrefix interface{}) (
	FileAccessControl,
	error) {

	if fAccess.lock == nil {
		fAccess.lock = new(sync.Mutex)
	}

	fAccess.lock.Lock()

	defer fAccess.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newFAccessCtrl := FileAccessControl{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileAccessControl."+
			"NewReadWriteAccess()",
		"")

	if err != nil {
		return newFAccessCtrl, err
	}

	var fileOpenCfg FileOpenConfig

	fileOpenCfg,
		err =
		new(FileOpenConfig).
			New(ePrefix.XCpy("fileOpenCfg<-"),
				FOpenType.TypeReadWrite(),
				FOpenMode.ModeNone())

	if err != nil {
		return newFAccessCtrl, err
	}

	var filePermCfg FilePermissionConfig

	filePermCfg,
		err = new(FilePermissionConfig).
		New("-rw-rw-rw-",
			ePrefix.XCpy(
				"filePermCfg<-'-rw-rw-rw-'"))

	if err != nil {
		return newFAccessCtrl, err
	}

	err = new(fileAccessControlMechanics).
		setFileAccessControl(
			&newFAccessCtrl,
			fileOpenCfg,
			filePermCfg,
			ePrefix.XCpy("newFAccessCtrl<-"))

	return newFAccessCtrl, nil
}

// NewReadWriteCreateAppendAccess
//
// Returns a FileAccessControl instance configured for
// Read, Write, Create and Append file access.
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
//		This new FileAccessControl instance will be
//		configured for Read, Write, Create and Truncate
//		file access.
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
func (fAccess *FileAccessControl) NewReadWriteCreateAppendAccess(
	errorPrefix interface{}) (
	FileAccessControl,
	error) {

	if fAccess.lock == nil {
		fAccess.lock = new(sync.Mutex)
	}

	fAccess.lock.Lock()

	defer fAccess.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newFileAccessCtrl := FileAccessControl{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileAccessControl."+
			"NewReadWriteCreateTruncateAccess()",
		"")

	if err != nil {
		return newFileAccessCtrl, err
	}

	var fOpenCfg FileOpenConfig

	//  OpenFile(name, O_RDWR|O_CREATE|O_APPEND, 0666)
	fOpenCfg,
		err = new(FileOpenConfig).New(
		ePrefix.XCpy("fOpenCfg<-"),
		FOpenType.TypeReadWrite(),
		FOpenMode.ModeCreate(),
		FOpenMode.ModeAppend())

	if err != nil {
		return newFileAccessCtrl, err
	}

	var fPermCfg FilePermissionConfig

	fPermCfg,
		err = new(FilePermissionConfig).
		New(
			"-rw-rw-rw-",
			ePrefix.XCpy(
				"<-'-rw-rw-rw-'"))

	if err != nil {
		return newFileAccessCtrl, err
	}

	err = new(fileAccessControlMechanics).
		setFileAccessControl(
			&newFileAccessCtrl,
			fOpenCfg,
			fPermCfg,
			ePrefix.XCpy("newFileAccessCtrl<-"))

	return newFileAccessCtrl, nil
}

// NewReadWriteCreateTruncateAccess
//
// Returns a FileAccessControl instance configured for
// Read, Write, Create and Truncate file access.
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
//		This new FileAccessControl instance will be
//		configured for Read, Write, Create and Truncate
//		file access.
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
func (fAccess *FileAccessControl) NewReadWriteCreateTruncateAccess(
	errorPrefix interface{}) (
	FileAccessControl,
	error) {

	if fAccess.lock == nil {
		fAccess.lock = new(sync.Mutex)
	}

	fAccess.lock.Lock()

	defer fAccess.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newFileAccessCtrl := FileAccessControl{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileAccessControl."+
			"NewReadWriteCreateTruncateAccess()",
		"")

	if err != nil {
		return newFileAccessCtrl, err
	}

	var fOpenCfg FileOpenConfig

	//  OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
	fOpenCfg,
		err = new(FileOpenConfig).New(
		ePrefix.XCpy("fOpenCfg<-"),
		FOpenType.TypeReadWrite(),
		FOpenMode.ModeCreate(),
		FOpenMode.ModeTruncate())

	if err != nil {
		return newFileAccessCtrl, err
	}

	var fPermCfg FilePermissionConfig

	fPermCfg,
		err = new(FilePermissionConfig).
		New(
			"-rw-rw-rw-",
			ePrefix.XCpy(
				"<-'-rw-rw-rw-'"))

	if err != nil {
		return newFileAccessCtrl, err
	}

	err = new(fileAccessControlMechanics).
		setFileAccessControl(
			&newFileAccessCtrl,
			fOpenCfg,
			fPermCfg,
			ePrefix.XCpy("newFileAccessCtrl<-"))

	return newFileAccessCtrl, nil
}

// NewReadOnlyAccess
//
// Returns a FileAccessControl instance configured for
// Read-Only access.
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
//		This new FileAccessControl instance will be
//		configured for Read-Only file access.
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
func (fAccess *FileAccessControl) NewReadOnlyAccess(
	errorPrefix interface{}) (
	FileAccessControl,
	error) {

	if fAccess.lock == nil {
		fAccess.lock = new(sync.Mutex)
	}

	fAccess.lock.Lock()

	defer fAccess.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	var newFileAccessCtrl FileAccessControl

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileAccessControl."+
			"NewReadOnlyAccess()",
		"")

	if err != nil {
		return newFileAccessCtrl, err
	}

	var fileOpenCfg FileOpenConfig

	fileOpenCfg,
		err = new(FileOpenConfig).
		New(ePrefix.XCpy("fileOpenCfg<-"),
			FOpenType.TypeReadOnly(),
			FOpenMode.ModeNone())

	if err != nil {
		return newFileAccessCtrl, err
	}

	var filePermCfg FilePermissionConfig

	filePermCfg,
		err = new(FilePermissionConfig).
		New("-r--r--r--",
			ePrefix.XCpy(
				"filePermCfg<-'-r--r--r--'"))

	if err != nil {
		return newFileAccessCtrl, err
	}

	err = new(fileAccessControlMechanics).
		setFileAccessControl(
			&newFileAccessCtrl,
			fileOpenCfg,
			filePermCfg,
			ePrefix.XCpy("newFileAccessCtrl<-"))

	return newFileAccessCtrl, err
}

// NewWriteOnlyAccess
//
// Returns a FileAccessControl instance configured for
// Write-Only file access.
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
//		This new FileAccessControl instance will be
//		configured for Write-Only file access.
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
func (fAccess *FileAccessControl) NewWriteOnlyAccess(
	errorPrefix interface{}) (
	FileAccessControl,
	error) {

	if fAccess.lock == nil {
		fAccess.lock = new(sync.Mutex)
	}

	fAccess.lock.Lock()

	defer fAccess.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	var newFileAccessCtrl FileAccessControl

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileAccessControl."+
			"NewWriteOnlyAccess()",
		"")

	if err != nil {
		return newFileAccessCtrl, err
	}

	var fileOpenCfg FileOpenConfig

	fileOpenCfg,
		err = new(FileOpenConfig).
		New(ePrefix.XCpy(
			"fileOpenCfg<-TypeWriteOnly"),
			FOpenType.TypeWriteOnly(),
			FOpenMode.ModeNone())

	if err != nil {
		return newFileAccessCtrl, err
	}

	var filePermCfg FilePermissionConfig

	filePermCfg,
		err = new(FilePermissionConfig).
		New("--w--w--w-",
			ePrefix.XCpy(
				"filePermCfg<- '--w--w--w-'"))

	if err != nil {
		return newFileAccessCtrl, err
	}

	err = new(fileAccessControlMechanics).
		setFileAccessControl(
			&newFileAccessCtrl,
			fileOpenCfg,
			filePermCfg,
			ePrefix.XCpy("newFileAccessCtrl<-"))

	return newFileAccessCtrl, err
}

// NewWriteOnlyCreateAppendAccess
//
// Returns a FileAccessControl instance configured for
// Write/Only - Create/Append file access.
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
//		This new FileAccessControl instance will be
//		configured for Write/Only - Create/Append file
//		access.
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
func (fAccess *FileAccessControl) NewWriteOnlyCreateAppendAccess(
	errorPrefix interface{}) (
	FileAccessControl,
	error) {

	if fAccess.lock == nil {
		fAccess.lock = new(sync.Mutex)
	}

	fAccess.lock.Lock()

	defer fAccess.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	var newFileAccessCtrl FileAccessControl

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileAccessControl."+
			"NewWriteOnlyCreateAppendAccess()",
		"")

	if err != nil {
		return newFileAccessCtrl, err
	}

	var fileOpenCfg FileOpenConfig

	fileOpenCfg,
		err = new(FileOpenConfig).
		New(ePrefix.XCpy(
			"fileOpenCfg<=TypeWriteOnly-ModeAppend"),
			FOpenType.TypeWriteOnly(),
			FOpenMode.ModeAppend(),
			FOpenMode.ModeCreate())

	if err != nil {
		return newFileAccessCtrl, err
	}

	var filePermCfg FilePermissionConfig

	filePermCfg,
		err = new(FilePermissionConfig).
		New("--w--w--w-",
			ePrefix.XCpy(
				"filePermCfg<- '--w--w--w-'"))

	if err != nil {
		return newFileAccessCtrl, err
	}

	err = new(fileAccessControlMechanics).
		setFileAccessControl(
			&newFileAccessCtrl,
			fileOpenCfg,
			filePermCfg,
			ePrefix.XCpy("newFileAccessCtrl<-"))

	return newFileAccessCtrl, err
}

// NewWriteOnlyAppendAccess
//
// Returns a FileAccessControl instance configured for
// Write/Only - Append file access.
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
//		This new FileAccessControl instance will be
//		configured for Write/Only - Append file access.
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
func (fAccess *FileAccessControl) NewWriteOnlyAppendAccess(
	errorPrefix interface{}) (
	FileAccessControl,
	error) {

	if fAccess.lock == nil {
		fAccess.lock = new(sync.Mutex)
	}

	fAccess.lock.Lock()

	defer fAccess.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	var newFileAccessCtrl FileAccessControl

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileAccessControl."+
			"NewWriteOnlyAccess()",
		"")

	if err != nil {
		return newFileAccessCtrl, err
	}

	var fileOpenCfg FileOpenConfig

	fileOpenCfg,
		err = new(FileOpenConfig).
		New(ePrefix.XCpy(
			"fileOpenCfg<=TypeWriteOnly-ModeAppend"),
			FOpenType.TypeWriteOnly(),
			FOpenMode.ModeAppend())

	if err != nil {
		return newFileAccessCtrl, err
	}

	var filePermCfg FilePermissionConfig

	filePermCfg,
		err = new(FilePermissionConfig).
		New("--w--w--w-",
			ePrefix.XCpy(
				"filePermCfg<- '--w--w--w-'"))

	if err != nil {
		return newFileAccessCtrl, err
	}

	err = new(fileAccessControlMechanics).
		setFileAccessControl(
			&newFileAccessCtrl,
			fileOpenCfg,
			filePermCfg,
			ePrefix.XCpy("newFileAccessCtrl<-"))

	return newFileAccessCtrl, err
}

// NewWriteOnlyTruncateAccess
//
// Returns a FileAccessControl instance configured for
// Write/Only - Truncate file access.
//
// If the file previously exists, it will be truncated
// before the writing operation commences.
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
//		This new FileAccessControl instance will be
//		configured for Write/Only - Truncate file access.
//
//		If the file previously exists, it will be
//		truncated before the writing operation commences.
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
func (fAccess *FileAccessControl) NewWriteOnlyTruncateAccess(
	errorPrefix interface{}) (
	FileAccessControl,
	error) {

	if fAccess.lock == nil {
		fAccess.lock = new(sync.Mutex)
	}

	fAccess.lock.Lock()

	defer fAccess.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	var newFileAccessCtrl FileAccessControl

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileAccessControl."+
			"NewWriteOnlyTruncateAccess()",
		"")

	if err != nil {
		return newFileAccessCtrl, err
	}

	var fileOpenCfg FileOpenConfig

	fileOpenCfg,
		err =
		new(FileOpenConfig).New(
			ePrefix.XCpy(
				"fileOpenCfg<=TypeWriteOnly-ModeTruncate"),
			FOpenType.TypeWriteOnly(),
			FOpenMode.ModeTruncate())

	if err != nil {
		return newFileAccessCtrl, err
	}

	var filePermCfg FilePermissionConfig

	filePermCfg,
		err = new(FilePermissionConfig).
		New("--w--w--w-",
			ePrefix.XCpy(
				"filePermCfg<= '--w--w--w-'"))

	if err != nil {
		return newFileAccessCtrl, err
	}

	err = new(fileAccessControlMechanics).
		setFileAccessControl(
			&newFileAccessCtrl,
			fileOpenCfg,
			filePermCfg,
			ePrefix.XCpy("newFileAccessCtrl<-"))

	return newFileAccessCtrl, err
}

// SetFileOpenCodes
//
// Assigns 'fileOpenCodes' to internal member variable,
// FileAccessControl.fileOpenCodes.
//
// This method only modify the File Open Codes. The
// Permission codes are NOT changed.
//
// ----------------------------------------------------------------
//
// # Background
//
// The FileAccessControl type is used when opening files
// for read and write operations.
//
// To open a file, two components are required:
//
//  1. A FileOpenType
//     Supported by this method. Input parameter
//     'fileOpenCodes' Type FileOpenConfig.
//
//     In order to open a file, exactly one of the
//     following File Open Codes MUST be specified:
//
//     FileOpenType(0).TypeReadOnly()
//     FileOpenType(0).TypeWriteOnly()
//     FileOpenType(0).TypeReadWrite()
//
//     -- AND --
//
//  2. A FileOpenMode
//     Not supported by this method. See
//     FileAccessControl.SetFilePermissionCodes()
//
//     In addition to a 'FileOpenType', a File Open Mode
//     is also required. This code is also referred to as
//     'permissions'. Zero or more of the following File
//     Open Mode codes may optionally be specified to
//     better control file open behavior.
//
//     FileOpenMode(0).ModeAppend()
//     FileOpenMode(0).ModeCreate()
//     FileOpenMode(0).ModeExclusive()
//     FileOpenMode(0).ModeSync()
//     FileOpenMode(0).ModeTruncate()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fileOpenCodes				FileOpenConfig
//
//		An instance of FileOpenConfig. If this instance
//		evaluates as invalid, an error will be returned.
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
func (fAccess *FileAccessControl) SetFileOpenCodes(
	fileOpenCodes FileOpenConfig,
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
			"SetFileOpenCodes()",
		"")

	if err != nil {
		return err
	}

	err = fileOpenCodes.IsValidInstanceError(
		ePrefix.XCpy("fileOpenCodes"))

	if err != nil {
		return err
	}

	fAccess.fileOpenCodes = fileOpenCodes.CopyOut()

	err = fAccess.permissions.IsValidInstanceError(ePrefix.XCpy(
		"fAccess.permissions"))

	if err == nil {

		fAccess.isInitialized = true

	} else {
		fAccess.isInitialized = false
	}

	return nil
}

// SetFilePermissionCodes
//
// Assigns 'filePermissions' to internal member variable
// FileAccessControl.permissions.
//
// ----------------------------------------------------------------
//
// # Background
//
// The FileAccessControl type is used when opening files
// for read and write operations.
//
// To open a file, two components are required:
//
//  1. A FileOpenType
//     Not Supported by this method. See method
//     FileAccessControl.SetFileOpenCodes().
//
//     In order to open a file, exactly one of the
//     following File Open Codes MUST be specified:
//
//     FileOpenType(0).TypeReadOnly()
//     FileOpenType(0).TypeWriteOnly()
//     FileOpenType(0).TypeReadWrite()
//
//     -- AND --
//
//  2. A FileOpenMode
//     Supported by this method. See input parameter
//     'filePermissions'.
//
//     In addition to a 'FileOpenType', a File Open Mode
//     is also required. This code is also referred to as
//     the File Permissions code. One or more of the
//     following File Open Mode codes may optionally be
//     specified to better control file open behavior.
//
//     FileOpenMode(0).None()
//     FileOpenMode(0).ModeAppend()
//     FileOpenMode(0).ModeCreate()
//     FileOpenMode(0).ModeExclusive()
//     FileOpenMode(0).ModeSync()
//     FileOpenMode(0).ModeTruncate()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	filePermissions				FilePermissionConfig
//
//		An instance of FilePermissionConfig. If this
//		instance evaluates as invalid, an error will be
//		returned.
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
func (fAccess *FileAccessControl) SetFilePermissionCodes(
	filePermissions FilePermissionConfig,
	errorPrefix interface{}) error {

	if fAccess.lock == nil {
		fAccess.lock = new(sync.Mutex)
	}

	fAccess.lock.Lock()

	defer fAccess.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	funcName := "FileAccessControl.SetFilePermissionCodes()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {
		return err
	}

	err = filePermissions.IsValidInstanceError(
		ePrefix.XCpy(
			"filePermissions"))

	if err != nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'filePermissions' INVALID!\n"+
			"Error = \n%v\n",
			funcName,
			err.Error())
	}

	fAccess.permissions = filePermissions.CopyOut()

	err = fAccess.fileOpenCodes.IsValidInstanceError(ePrefix)

	if err == nil {
		fAccess.isInitialized = true
	} else {
		fAccess.isInitialized = false
	}

	return nil
}
