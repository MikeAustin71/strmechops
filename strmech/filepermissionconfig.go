package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"os"
	"strings"
	"sync"
)

// FilePermissionConfig - Provides methods to support the creation and management of
// File Permissions for use in controlling file access operations. The Go Programming
// Language uses os.FileMode (https://golang.org/pkg/os/#FileMode) and unix permission
// bits to configure file permissions.
//
//	Reference:
//	https://www.cyberciti.biz/faq/explain-the-nine-permissions-bits-on-files/
//	https://en.wikipedia.org/wiki/File_system_permissions
//
// The FilePermissionConfig methods will allow for configuration of valid file permissions
// which are subsequently stored as an os.FileMode in a private member variable,
// 'FilePermissionConfig.fileMode'.
//
// When evaluated as a string, file permission is defined by a 10-character string. The
// first character is an 'Entry Type' and the remaining 9-characters are unix permission
// bits.
//
//	Example: -rwxrwxrwx - Identifies permissions for a regular file
//	         drwxrwxrwx - Identifies permissions for directory
//	                      value = 020000000777
//
// ----------------------------------------------------------------
//
//	Symbolic and Numeric Notation
//
// Permission codes may be designated with Symbolic
// Notation or Numeric Octal Notation.
//
//				   Numeric
//		Symbolic	Octal
//		Notation   Notation
//		----------	0000	no permissions
//		-rwx------	0700	read, write, & execute only for owner
//		-rwxrwx---	0770	read, write, & execute for owner and group
//		-rwxrwxrwx	0777	read, write, & execute for owner, group and others
//		---x--x--x	0111	execute
//		--w--w--w-	0222	write
//		--wx-wx-wx	0333	write & execute
//		-r--r--r--	0444	read
//		-r-xr-xr-x	0555	read & execute
//		-rw-rw-rw-	0666	read & write
//		-rwxr-----	0740	owner can read, write, & execute; group can only read;
//	                       others have no permissions
//
// Internal private member variable stores the consolidated permission as a numerical
// value in 'FilePermissionConfig.fileMode'.
type FilePermissionConfig struct {
	isInitialized bool
	// If set to 'true', this boolean value signals that
	// the current instance of

	fileMode os.FileMode
	// Holds the consolidated file permission value which
	// consists of the two elements making up a permission
	// value:
	//	(1)	Entry Type
	// 	(2) Permission Bits
	//
	// Reference:
	//	https://pkg.go.dev/os#FileMode

	lock *sync.Mutex
}

// CopyIn
//
// Receives a FilePermissionConfig instance and copies
// all data fields to the current FilePermissionConfig
// instance. When complete, both the incoming and current
// FilePermissionConfig instances will be identical. The
// type of copy operation performed is a 'deep copy'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fPerm2						*FilePermissionConfig
//
//		A pointer to an incoming instance of
//		FilePermissionConfig. All the internal data
//		fields contained in this instance will be copied
//		to corresponding data fields in the current
//		FilePermissionConfig instance.
//
//		When the deep copy operation is completed, both
//		instances will contain identical data values.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	--- NONE ---
func (fPerm *FilePermissionConfig) CopyIn(
	fPerm2 *FilePermissionConfig) {

	if fPerm.lock == nil {
		fPerm.lock = new(sync.Mutex)
	}

	fPerm.lock.Lock()

	defer fPerm.lock.Unlock()

	fPerm.isInitialized = fPerm2.isInitialized

	fPerm.fileMode = fPerm2.fileMode

}

// CopyOut
//
// Returns a new instance of FilePermissionConfig which
// is in all respects an exact duplicate of the current
// FilePermissionConfig instance. The type of copy
// operation performed  is a 'deep copy'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	--- NONE ---
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	FilePermissionConfig
//
//		Returns a new instance of FilePermissionConfig
//		which is identical in all respects to the current
//		instance of FilePermissionConfig.
//
//		The type of copy operation performed  is a
//		'deep copy'.
func (fPerm *FilePermissionConfig) CopyOut() FilePermissionConfig {

	if fPerm.lock == nil {
		fPerm.lock = new(sync.Mutex)
	}

	fPerm.lock.Lock()

	defer fPerm.lock.Unlock()

	fPerm2 := FilePermissionConfig{}

	fPerm2.isInitialized = fPerm.isInitialized
	fPerm2.fileMode = fPerm.fileMode

	return fPerm2
}

// Empty
//
// ReInitializes the current FilePermissionConfig
// instance to empty or zero values.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the current instance of
//	FilePermissionConfig to their uninitialized or zero
//	values.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	--- NONE ---
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	--- NONE ---
func (fPerm *FilePermissionConfig) Empty() {

	if fPerm.lock == nil {
		fPerm.lock = new(sync.Mutex)
	}

	fPerm.lock.Lock()

	defer fPerm.lock.Unlock()

	fPerm.isInitialized = false

	fPerm.fileMode = os.FileMode(0)
}

// Equal
//
// Returns 'true' if the incoming FilePermissionConfig
// instance is equal in all respects to the current
// FilePermissionConfig instance.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fPerm2						*FilePermissionConfig
//
//		A pointer to an external instance of
//		FilePermissionConfig. The internal member
//		variable data values in this instance will be
//		compared to those in the current instance of
//		FilePermissionConfig. The results of this
//		comparison will be returned to the calling
//		function as a boolean value.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	boolean
//
//		If the internal member variable data values
//		contained in input parameter 'fPerm2' are
//		equivalent in all respects to those contained in
//		the current instance of FilePermissionConfig,
//		this return value will be set to 'true'.
//
//		Otherwise, this method will return 'false'.
func (fPerm *FilePermissionConfig) Equal(fPerm2 *FilePermissionConfig) bool {

	if fPerm.lock == nil {
		fPerm.lock = new(sync.Mutex)
	}

	fPerm.lock.Lock()

	defer fPerm.lock.Unlock()

	if fPerm.isInitialized != fPerm2.isInitialized {
		return false
	}

	if fPerm.fileMode != fPerm2.fileMode {
		return false
	}

	return true
}

// GetIsDir
//
// Returns a boolean value signaling whether the
// FileMode contained in the current instance of
// FilePermissionConfig is a directory or not.
//
// A returned value of 'true' signals that the
// FileMode represents a directory.
//
// This method serves as a wrapper for method
// os.FileMode.IsDir()
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
//	bool
//
//		This returned boolean value signals whether the
//		FileMode contained in the current instance of
//		FilePermissionConfig is a directory or not.
//
//		A returned value of 'true' signals that the
//		FileMode represents a directory.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (fPerm *FilePermissionConfig) GetIsDir(
	errorPrefix interface{}) (
	bool,
	error) {

	if fPerm.lock == nil {
		fPerm.lock = new(sync.Mutex)
	}

	fPerm.lock.Lock()

	defer fPerm.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FilePermissionConfig.GetIsDir()",
		"")

	if err != nil {
		return false, err
	}

	_,
		err = new(filePermissionConfigElectron).
		testValidityOfFilePermissionConfig(
			fPerm,
			ePrefix)

	if err != nil {

		return false, err
	}

	return fPerm.fileMode.IsDir(), nil
}

// GetIsRegularFile
//
// Returns a boolean value signaling whether the
// FileMode contained in the current instance of
// FilePermissionConfig is a regular file or not.
//
// A returned value of 'true' signals that the
// FileMode represents a file.
//
// This method serves as a wrapper for method
// os.FileMode.IsRegular()
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
//	bool
//
//		This returned boolean value signals whether the
//		FileMode contained in the current instance of
//		FilePermissionConfig is a regular file or not.
//
//		A returned value of 'true' signals that the
//		FileMode represents a file.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (fPerm *FilePermissionConfig) GetIsRegularFile(
	errorPrefix interface{}) (
	bool,
	error) {

	if fPerm.lock == nil {
		fPerm.lock = new(sync.Mutex)
	}

	fPerm.lock.Lock()

	defer fPerm.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FilePermissionConfig."+
			"GetIsRegularFile()",
		"")

	if err != nil {
		return false, err
	}

	_,
		err = new(filePermissionConfigElectron).
		testValidityOfFilePermissionConfig(
			fPerm,
			ePrefix)

	if err != nil {

		return false, err
	}

	return fPerm.fileMode.IsRegular(), nil
}

// GetEntryTypeComponent
//
// Returns the 'Entry Type' component of the current
// os.FileMode permissions value.
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
//	OsFilePermissionCode
//
//		This returned OsFilePermissionCode type is set to
//		the value of the os.FileMode constant
//		representing the Entry Type associated with the
//		permission value encapsulated by the current
//		FilePermissionConfig instance.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (fPerm *FilePermissionConfig) GetEntryTypeComponent(
	errorPrefix interface{}) (
	OsFilePermissionCode,
	error) {

	if fPerm.lock == nil {
		fPerm.lock = new(sync.Mutex)
	}

	fPerm.lock.Lock()

	defer fPerm.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FilePermissionConfig."+
			"GetEntryTypeComponent()",
		"")

	if err != nil {
		return OsFilePermissionCode(0), err
	}

	return new(filePermissionConfigMolecule).
		getEntryTypeComponent(
			fPerm,
			ePrefix.XCpy("fPerm<-"))
}

// GetCompositePermissionMode
//
// Returns the os.FileMode from the internal data field,
// 'FilePermissionConfig.fileMode'.
//
// 'fileMode' represents the complete, consolidated
// permission code. It therefore contains the two
// elements which make up a consolidated permission code:
// Entry Type and Permission Bits.
//
// This method returns the complete permission code as a
// type 'os.FileMode'.
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
//	os.FileMode
//
//		(From Golang Docs)
//		A FileMode represents a file's mode and permission
//		bits. The bits have the same definition on all
//		systems, so that information about files can be
//		moved from one system to another portably. Not all
//		bits apply to all systems. The only required bit
//		is ModeDir for directories.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (fPerm *FilePermissionConfig) GetCompositePermissionMode(
	errorPrefix interface{}) (
	os.FileMode,
	error) {

	if fPerm.lock == nil {
		fPerm.lock = new(sync.Mutex)
	}

	fPerm.lock.Lock()

	defer fPerm.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FilePermissionConfig."+
			"GetCompositePermissionMode()",
		"")

	if err != nil {
		return os.FileMode(0), err
	}

	return new(filePermissionConfigMechanics).
		getCompositePermissionMode(
			fPerm,
			ePrefix.XCpy("<-fPerm"))
}

// GetIsRegular
//
// Returns a 'bool' indicating whether the encapsulated
// FileMode is a file or not.
//
// A returned value of 'true' signals that the FileMode
// represents a file.
//
// This method serves as a wrapper for:
//
//	os.FileMode.IsRegular()
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
//	bool
//
//		This returned boolean value signals whether the
//		FileMode encapsulated in the current instance of
//		FilePermissionConfig is a file or not.
//
//		A returned value of 'true' signals that the
//		FileMode represents a file.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (fPerm *FilePermissionConfig) GetIsRegular(
	errorPrefix interface{}) (
	bool,
	error) {

	if fPerm.lock == nil {
		fPerm.lock = new(sync.Mutex)
	}

	fPerm.lock.Lock()

	defer fPerm.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FilePermissionConfig."+
			"GetIsRegular()",
		"")

	if err != nil {
		return false, err
	}

	_,
		err = new(filePermissionConfigElectron).
		testValidityOfFilePermissionConfig(
			fPerm,
			ePrefix)

	if err != nil {

		return false, err
	}

	return fPerm.fileMode.IsRegular(), nil
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
//	os.FileMode
//
//		This parameter returns a FileMode containing only
//		the least significant 9-bits of the encapsulated
//		FileMode representing the unix permission bits.
//
//		The os.FileMode value contained in input
//		parameter 'fPerm' will be used to produce the
//		permission bits value returned as os.FileMode.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (fPerm *FilePermissionConfig) GetFileMode(
	errorPrefix interface{}) (
	os.FileMode,
	error) {

	if fPerm.lock == nil {
		fPerm.lock = new(sync.Mutex)
	}

	fPerm.lock.Lock()

	defer fPerm.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FilePermissionConfig.GetFileMode()",
		"")

	if err != nil {
		return os.FileMode(0), err
	}

	return new(filePermissionConfigMolecule).getFileMode(
		fPerm,
		ePrefix.XCpy("<-fPerm"))
}

// GetPermissionComponents
//
// Returns the two components of a permission
// configuration:
//
//	(1)	Entry Type
//
//		AND
//
//	(2)	Permission Bits
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
//	entryType					OsFilePermissionCode
//
//		The Entry Type or os mode value. Generally this
//		will either be OsFilePermissionCode(0).ModeNone()
//		for files or OsFilePermissionCode(0).ModeDir()
//		for directories.
//
//		For more information see method
//			FilePermissionConfig.GetEntryTypeComponent()
//
//	permissionBits				os.FileMode
//
//		The 9-least significant bits designate the unix
//		permission bits.
//
//		Be advised that if you call string on this result
//		(permissionBits.String()) you will receive a
//		10-character string the first character of which
//		is always a hyphen ("-"). Disregard this first
//		character, only the last 9-characters of the
//		string are valid permission descriptors.
//
//		For more information see method
//			FilePermissionConfig.GetFileMode()
//
//		To create a full and complete permission code,
//		permissionBits must be or'd with a valid Entry
//		Type os mode value.
//
//	err							error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (fPerm *FilePermissionConfig) GetPermissionComponents(
	errorPrefix interface{}) (
	osMode OsFilePermissionCode,
	permissionBits os.FileMode,
	err error) {

	if fPerm.lock == nil {
		fPerm.lock = new(sync.Mutex)
	}

	fPerm.lock.Lock()

	defer fPerm.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	osMode =
		OsFilePermissionCode(OsFilePermCode.ModeNone())

	permissionBits =
		os.FileMode(0)

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FilePermissionConfig."+
			"GetPermissionComponents()",
		"")

	if err != nil {
		return osMode, permissionBits, err
	}

	_,
		err = new(filePermissionConfigElectron).
		testValidityOfFilePermissionConfig(
			fPerm,
			ePrefix)

	if err != nil {

		return osMode, permissionBits, err
	}

	osMode,
		err = new(filePermissionConfigMolecule).
		getEntryTypeComponent(
			fPerm,
			ePrefix.XCpy("<-fPerm"))

	if err != nil {

		return osMode, permissionBits, err
	}

	permissionBits,
		err = new(filePermissionConfigMolecule).getFileMode(
		fPerm,
		ePrefix.XCpy("<-fPerm"))

	return osMode, permissionBits, err
}

// GetPermissionFileModeValueText
//
// Returns the Permission File Mode numeric value as
// text. The text presents the octal value of the File
// Mode.
//
//	Example:
//	      -rw-rw-rw- = returned value 0666
//	      drwxrwxrwx = returned value 020000000777
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
//	string
//
//		If this method completes successfully, the
//		Permission File Mode numeric value will be
//		returned as text. This text presents the
//		octal value of the File Mode.
//
//		Example:
//			-rw-rw-rw- = returned value 0666
//			drwxrwxrwx = returned value 020000000777
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (fPerm *FilePermissionConfig) GetPermissionFileModeValueText(
	errorPrefix interface{}) (
	string,
	error) {

	if fPerm.lock == nil {
		fPerm.lock = new(sync.Mutex)
	}

	fPerm.lock.Lock()

	defer fPerm.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FilePermissionConfig."+
			"GetPermissionFileModeValueText()",
		"")

	if err != nil {
		return "", err
	}

	sb := strings.Builder{}
	sb.Grow(300)

	_,
		err = new(filePermissionConfigElectron).
		testValidityOfFilePermissionConfig(
			fPerm,
			ePrefix)

	if err != nil {

		return "", err
	}

	fileMode, err := new(filePermissionConfigMechanics).
		getCompositePermissionMode(
			fPerm,
			ePrefix)

	if err != nil {

		return "", err

	} else {

		octalValStr := "0" + fmt.Sprintf("%d",
			new(NumberConversions).
				ConvertDecimalToOctal(int(fileMode)))

		octalValStr = strings.Trim(octalValStr, " ")

		sb.WriteString(octalValStr)

	}

	return sb.String(), err
}

// GetPermissionNarrativeText
//
// Returns a string containing a narrative text
// description of the current permission codes contained
// in the current instance of FilePermissionConfig.
//
// Example Return Value
//
//	"Entry Type: ModeFile  -Permission Code: -rwxrwxrwx -File Mode Value: 0777"
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	--- NONE ---
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	string
//
//		Returns a narrative text description of the
//		permission codes contained in the current
//		instance of FilePermissionConfig.
//
//		Example Return Value
//		"Entry Type: ModeFile  -Permission Code: -rwxrwxrwx -File Mode Value: 0777"
func (fPerm *FilePermissionConfig) GetPermissionNarrativeText() string {

	if fPerm.lock == nil {
		fPerm.lock = new(sync.Mutex)
	}

	fPerm.lock.Lock()

	defer fPerm.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	sb := strings.Builder{}
	sb.Grow(300)

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"FilePermissionConfig."+
			"GetPermissionNarrativeText()",
		"")

	if err != nil {

		sb.WriteString(err.Error())

		return sb.String()
	}

	_,
		err = new(filePermissionConfigElectron).
		testValidityOfFilePermissionConfig(
			fPerm,
			ePrefix)

	if err != nil {
		sb.WriteString(err.Error())
		return sb.String()
	}

	osMode,
		err := new(filePermissionConfigMolecule).
		getEntryTypeComponent(
			fPerm,
			ePrefix)

	if err != nil {
		sb.WriteString("Entry Type: INVALID!\n" +
			err.Error() + "\n")

		return sb.String()
	}

	osModeStr := osMode.String()

	osModeStr =
		strings.Replace(
			osModeStr,
			"ModeNone",
			"ModeFile", 1)

	sb.WriteString(fmt.Sprintf("Entry Type: %s", osModeStr))

	var txtCode string

	txtCode,
		err = new(filePermissionConfigMechanics).getPermissionTextCode(
		fPerm,
		ePrefix.XCpy("txtCode<-"))

	if err != nil {
		sb.WriteString("  -Permission Code: INVALID!")
	} else {
		sb.WriteString("  -Permission Code: " + txtCode + " ")
	}

	var fileMode os.FileMode

	fileMode,
		err = new(filePermissionConfigMechanics).
		getCompositePermissionMode(
			fPerm,
			ePrefix.XCpy("fileMode<-"))

	if err != nil {

		sb.WriteString("  -File Mode Value: INVALID!")

	} else {

		octalValStr := "0" + fmt.Sprintf("%d",
			new(NumberConversions).
				ConvertDecimalToOctal(int(fileMode)))

		sb.WriteString(fmt.Sprintf("  -File Mode Value: %s",
			octalValStr))
	}

	sb.WriteString("\n")

	return sb.String()
}

// GetPermissionTextCode
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
func (fPerm *FilePermissionConfig) GetPermissionTextCode(
	errorPrefix interface{}) (
	string,
	error) {

	if fPerm.lock == nil {
		fPerm.lock = new(sync.Mutex)
	}

	fPerm.lock.Lock()

	defer fPerm.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FilePermissionConfig."+
			"GetPermissionTextCode()",
		"")

	if err != nil {
		return "", err
	}

	return new(filePermissionConfigMechanics).
		getPermissionTextCode(
			fPerm,
			ePrefix)
}

// IsValidInstanceError
//
// If the current FilePermissionConfig instance is judged
// to be 'Invalid', this method will return an error.
//
// Otherwise, if the current instance of
// FilePermissionConfig evaluates as 'Valid', this method
// will return 'nil'.
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
//		If errors are encountered during processing or
//		if the current instance of FilePermissionConfig
//		is found to be invalid, this returned error Type
//		will encapsulate an appropriate error message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
//
//		If the current instance of FilePermissionConfig
//		is found to be valid and no errors present during
//		processing, the returned error parameter is 'nil'.
func (fPerm *FilePermissionConfig) IsValidInstanceError(
	errorPrefix interface{}) error {

	if fPerm.lock == nil {
		fPerm.lock = new(sync.Mutex)
	}

	fPerm.lock.Lock()

	defer fPerm.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FilePermissionConfig."+
			"IsValidInstanceError()",
		"")

	if err != nil {
		return err
	}

	_,
		err = new(filePermissionConfigElectron).
		testValidityOfFilePermissionConfig(
			fPerm,
			ePrefix.XCpy("fPerm"))

	return err
}

// New
//
// Creates and returns a new FilePermissionConfig
// instance initialized with an os.FileMode value
// generated from the input parameter 'modeStr'.
//
// 'modeStr' is a 10-character string containing the
// read, write and execute permissions for the three
// groups, 'Owner', 'Group' and 'Other'.
//
// The text codes used in the 'modeStr' mimic the Unix
// permission codes.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://www.cyberciti.biz/faq/explain-the-nine-permissions-bits-on-files/.
//	https://en.wikipedia.org/wiki/File-system_permissions
//	https://linuxconfig.org/how-to-use-special-permissions-the-setuid-setgid-and-sticky-bits
//
// ----------------------------------------------------------------
//
// The first character of the 'modeStr' designates the
// 'Entry Type'. Currently, only two 'Entry Type'
// characters are supported. Therefore, the first
// character in the 10-character input parameter
// 'modeStr' MUST be either a "-" indicating a file, or
// a "d" indicating a directory.
//
// The remaining nine characters in the 'modeStr'
// represent unix permission bits and consist of three
// group fields each containing 3-characters. Each
// character in the three group fields may consist of
// 'r' (Read-Permission), 'w' (Write-Permission), 'x'
// (Execute-Permission) or '-' signaling no permission or
// no access allowed. A typical 'modeStr' authorizing
// permission for full access to a file would be styled
// as:
//
//		Example: "-rwxrwxrwx"
//
//		Groups: - Owner/User, Group, Other
//		From left to right
//		First Characters is Entry Type index 0 ("-")
//
//		First Char index 0 =     "-"   Designates a file
//
//		First Char index 0 =     "d"   Designates a directory
//
//		Char indexes 1-3 = Owner "rwx" Authorizing 'Read',
//	                                  Write' & Execute Permissions for 'Owner'
//
//		Char indexes 4-6 = Group "rwx" Authorizing 'Read', 'Write' & Execute
//	                                  Permissions for 'Group'
//
//		Char indexes 7-9 = Other "rwx" Authorizing 'Read', 'Write' & Execute
//	                                  Permissions for 'Other'
//
// The Symbolic notation provided by input parameter 'modeStr' MUST conform to
// the options presented below. The first character or 'Entry Type' is listed as
// "-". However, in practice, the caller may set the first character as either a
// "-", specifying a file, or a "d", specifying a directory. No other first character
// types are currently supported.
//
// Three SymbolicGroups:
//
//	The three group types are: User/Owners, Groups & Others.
//
// Directory Permissions:
//
//	        -----------------------------------------------------
//	               Directory Mode String Permission Codes
//	        -----------------------------------------------------
//	          Directory
//				10-Character
//				 'modeStr'
//				 Symbolic		  Directory Access
//				  Format	   Permission Descriptions
//				----------------------------------------------------
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
// File Permissions:
//
//	        -----------------------------------------------------
//	               File Mode String Permission Codes
//	        -----------------------------------------------------
//
//	               File
//				10-Character
//				 'modeStr'
//				 Symbolic	Octal		File Access
//				  Format	Notation  Permission Descriptions
//				------------------------------------------------------------
//
//				----------	0000	no permissions
//				-rwx------	0700	read, write, & execute only for owner
//				-rwxrwx---	0770	read, write, & execute for owner and group
//				-rwxrwxrwx	0777	read, write, & execute for owner, group and others
//				---x--x--x	0111	execute
//				--w--w--w-	0222	write
//				--wx-wx-wx	0333	write & execute
//				-r--r--r--	0444	read
//				-r-xr-xr-x	0555	read & execute
//				-rw-rw-rw-	0666	read & write
//				-rwxr-----	0740	Owner can read, write, & execute. Group can only read;
//				                             others have no permissions
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//	FilePermissionConfig
//
//		If this method completes successfully, a new,
//		fully populated instance of FilePermissionConfig
//		will be returned configured with the permission
//		codes contained in input parameter 'modeStr'.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (fPerm *FilePermissionConfig) New(
	filePermissionStr string,
	errorPrefix interface{}) (
	FilePermissionConfig,
	error) {

	if fPerm.lock == nil {
		fPerm.lock = new(sync.Mutex)
	}

	fPerm.lock.Lock()

	defer fPerm.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	fPerm2 := FilePermissionConfig{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FilePermissionConfig."+
			"New()",
		"")

	if err != nil {
		return fPerm2, err
	}

	err = new(filePermissionConfigNanobot).
		setFileModeByTextCode(
			&fPerm2,
			filePermissionStr,
			ePrefix.XCpy(
				"fPerm2<-filePermissionStr"))

	return fPerm2, nil
}

// NewByComponents
//
// Creates and returns a new instance of
// FilePermissionConfig using permission value supplied
// by two input parameters, 'entryType' and
// 'unixPermissionTextStr'.
//
// ------------------------------------------------------------------------
//
// # Warning
//
// Incorrect or invalid File Permissions can cause extensive damage. If you
// don't know what you are doing, you would be well advised to use one of
// the other methods in this type which provide additional safeguards.
//
// If you decide to proceed, be guided by the wisdom of Davy Crockett:
//
//	"Be always sure you are right - then go ahead."
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://golang.org/pkg/os/#FileMode
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//		entryType					OsFilePermissionCode
//
//			The code which makes up the first character in a
//			10-digit unix permission character string.
//
//			This a wrapper for os.FileMode constants.
//			Reference:
//				https://golang.org/pkg/os/#FileMode
//
//			Select this value with caution.
//			See the warning above.
//
//		unixPermissionStr			string
//
//			A 9-character string containing the unix
//			permission bits expressed as three groups of
//		 	3-characters each.
//
//			The 9-characters are constituents of the three
//			Symbolic Groups or User Classes:
//				(1) Owners/Users
//				(2) Groups
//				(3) Others.
//			Each group has three characters which may be 'r',
//			'w', 'x'. If a permission is not set, that
//			character position contains a '-'.
//
//		 	Unix Permission String Options:
//
//		      9-Character          File Access
//		      Notation             Permission Descriptions
//		-----------------------------------------------------------
//		      ---------            File - no permissions
//		      rwx------            File - read, write, & execute only for owner
//		      rwxrwx---            File - read, write, & execute for owner and group
//		      rwxrwxrwx            File - read, write, & execute for owner, group and others
//		      --x--x--x            File - execute
//		      -w--w--w-            File - write
//		      -wx-wx-wx            File - write & execute
//		      r--r--r--            File - read
//		      r-xr-xr-x            File - read & execute
//		      rw-rw-rw-            File - read & write
//		      rwxr-----            File - Owner can read, write, & execute. Group can only read;
//
//
//			Note: drwxrwxrwx - identifies permissions for
//	                        directory
//
//		errorPrefix					interface{}
//
//			This object encapsulates error prefix text which
//			is included in all returned error messages.
//			Usually, it contains the name of the calling
//			method or methods listed as a method or function
//			chain of execution.
//
//			If no error prefix information is needed, set
//			this parameter to 'nil'.
//
//			This empty interface must be convertible to one
//			of the following types:
//
//			1.	nil
//					A nil value is valid and generates an
//					empty collection of error prefix and
//					error context information.
//
//			2.	string
//					A string containing error prefix
//					information.
//
//			3.	[]string
//					A one-dimensional slice of strings
//					containing error prefix information.
//
//			4.	[][2]string
//					A two-dimensional slice of strings
//			   		containing error prefix and error
//			   		context information.
//
//			5.	ErrPrefixDto
//					An instance of ErrPrefixDto.
//					Information from this object will
//					be copied for use in error and
//					informational messages.
//
//			6.	*ErrPrefixDto
//					A pointer to an instance of
//					ErrPrefixDto. Information from
//					this object will be copied for use
//					in error and informational messages.
//
//			7.	IBasicErrorPrefix
//					An interface to a method
//					generating a two-dimensional slice
//					of strings containing error prefix
//					and error context information.
//
//			If parameter 'errorPrefix' is NOT convertible
//			to one of the valid types listed above, it will
//			be considered invalid and trigger the return of
//			an error.
//
//			Types ErrPrefixDto and IBasicErrorPrefix are
//			included in the 'errpref' software package:
//				"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	FilePermissionConfig
//
//		If this method completes successfully, a new,
//		fully populated instance of FilePermissionConfig
//		will be returned configured with the permission
//		codes contained in input parameter 'modeStr'.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (fPerm *FilePermissionConfig) NewByComponents(
	entryType OsFilePermissionCode,
	unixPermissionTextStr string,
	errorPrefix interface{}) (FilePermissionConfig, error) {

	if fPerm.lock == nil {
		fPerm.lock = new(sync.Mutex)
	}

	fPerm.lock.Lock()

	defer fPerm.lock.Unlock()

	fPerm2 := FilePermissionConfig{}

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FilePermissionConfig."+
			"NewByComponents()",
		"")

	if err != nil {
		return fPerm2, err
	}

	err = new(filePermissionConfigNanobot).
		setFileModeByComponents(
			&fPerm2,
			entryType,
			unixPermissionTextStr,
			ePrefix.XCpy(
				"fPerm2<-"))

	return fPerm2, err
}

// NewByFileMode
//
// Creates and returns a new instance of
// FilePermissionConfig. This instance is initialized
// using the input parameter 'fMode' of type
// 'os.FileMode'.  'fMode' is assumed to contain all
// the codes necessary for the configuration of unix file
// permission bits.
//
// Unix file permission bits are used by the Go
// Programming language to configure file permissions on
// all supported operating systems.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://pkg.go.dev/os#FileMode
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fMode						os.FileMode
//
//		An instance of os.FileMode containing file or
//		directory permission codes. These permission
//		codes will be used to reset the internal
//		FileMode data field in the FilePermissionConfig
//		instance returned by this method.
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
//	FilePermissionConfig
//
//		If this method completes successfully, a new,
//		fully populated instance of FilePermissionConfig
//		will be returned configured with the permission
//		codes contained in input parameter 'fMode'.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (fPerm *FilePermissionConfig) NewByFileMode(
	fMode os.FileMode,
	errorPrefix interface{}) (
	FilePermissionConfig,
	error) {

	if fPerm.lock == nil {
		fPerm.lock = new(sync.Mutex)
	}

	fPerm.lock.Lock()

	defer fPerm.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	fPerm2 := FilePermissionConfig{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FilePermissionConfig."+
			"NewByFileMode()",
		"")

	if err != nil {
		return fPerm2, err
	}

	err = new(filePermissionConfigNanobot).
		setByFileMode(
			&fPerm2,
			fMode,
			ePrefix.XCpy(
				"fPerm<-fMode"))

	return fPerm2, err
}

// NewByOctalDigits
//
// Creates and returns a new FilePermissionConfig
// instance by initializing the internal FileMode data
// field (FilePermissionConfig.fileMode) to the value
// represented by input parameter, 'octalFileModeCode'.
//
// ------------------------------------------------------------------------
//
// # Warning
//
// In the Go Programming Language, if you initialize an
// integer with a leading zero (e.g. x:= int(0777)), than
// number ('0777') is treated as an octal value and
// converted to a decimal value. Therefore, x:= int(0777)
// will mean that 'x' is set equal to 511. If you set
// x:= int(777), x will be set equal to '777'. For purposes
// of this method enter the octal code as x:= int(777).
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	octalFileModeCode			int
//
//		This parameter contains the integer value of the
//		of the permission code which will be used to
//		initialize the current FilePermissionConfig
//		instance (FilePermissionConfig.fileMode). The
//		integer digits in 'octalFileModeCode' represent
//		the octal value for the file permissions.
//
//		If the input parameter 'octalFileModeCode'
//		contains an invalid Entry Type, an error will be
//		returned.
//
//		A partial list of valid file permission value
//	 	examples are shown as follows:
//
//	 ____________________________________________________________________________
//
//	          Input Parameter
//	              integer            Equivalent
//	 Octal    'octalFileModeCode'    Symbolic      File Access
//	 Digits        value             Notation      Permission Descriptions
//	 0000 	         0               ----------    File - no permissions
//	 0700 	       700               -rwx------    File - read, write, & execute only for owner
//	 0770 	       770               -rwxrwx---    File - read, write, & execute for owner and group
//	 0777 	       777               -rwxrwxrwx    File - read, write, & execute for owner, group and others
//	 0111 	       111               ---x--x--x    File - execute
//	 0222 	       222               --w--w--w-    File - write
//	 0333 	       333               --wx-wx-wx    File - write & execute
//	 0444 	       444               -r--r--r--    File - read
//	 0555 	       555               -r-xr-xr-x    File - read & execute
//	 0666 	       666               -rw-rw-rw-    File - read & write
//	 0740 	       740               -rwxr-----    File - Owner can read, write, & execute. Group can only read;
//	                                                      others have no permissions
//
//	 drwxrwxrwx    Directory - read, write, & execute for owner, group and others
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
//	FilePermissionConfig
//
//		If this method completes successfully, a new,
//		fully populated instance of FilePermissionConfig
//		will be returned configured with the permission
//		codes contained in input parameter
//		'octalFileModeCode'.
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
//
//		If the input parameter 'octalFileModeCode'
//		contains an invalid Entry Type, an error will be
//		returned.
func (fPerm *FilePermissionConfig) NewByOctalDigits(
	octalFileModeCode int,
	errorPrefix interface{}) (
	FilePermissionConfig,
	error) {

	if fPerm.lock == nil {
		fPerm.lock = new(sync.Mutex)
	}

	fPerm.lock.Lock()

	defer fPerm.lock.Unlock()

	fPerm2 := FilePermissionConfig{}

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FilePermissionConfig."+
			"NewByOctalDigits()",
		"")

	if err != nil {
		return fPerm2, err
	}

	err = new(filePermissionConfigNanobot).
		setFileModeByOctalDigits(
			&fPerm2,
			octalFileModeCode,
			ePrefix.XCpy(
				"fPerm<-octalFileModeCode"))

	return fPerm2, nil
}

// SetFileModeByComponents
//
// Sets the value of the current FilePermissionConfig
// instance by initializing the internal FileMode data
// field (FilePermissionConfig.fileMode).
//
// The final FileMode value is computed by integrating
// the 'entryType' FileMode with the unix permission
// symbolic values provided by the input parameter,
// 'unixPermissionStr'. This approach allows the caller
// to created custom File Permissions.
//
// ------------------------------------------------------------------------
//
// # Warning
//
// Incorrect or invalid File Permissions can cause
// extensive damage. If you don't know what you are
// doing, you would be well advised to use one of the
// other methods in this type which provide additional
// safeguards.
//
// If you decide to proceed, be guided by the wisdom of
// Davy Crockett:
//
//	"Be always sure you are right - then go ahead."
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//		entryType					OsFilePermissionCode
//
//			The code which makes up the first character in a
//			10-digit unix permission character string.
//
//			This a wrapper for os.FileMode constants.
//			Reference:
//				https://golang.org/pkg/os/#FileMode
//
//			Select this value with caution.
//			See the warning above.
//
//		unixPermissionStr			string
//
//			A 9-character string containing the unix
//			permission bits expressed as three groups of
//		 	3-characters each.
//
//			The 9-characters are constituents of the three
//			Symbolic Groups or User Classes:
//				(1) Owners/Users
//				(2) Groups
//				(3) Others.
//			Each group has three characters which may be 'r',
//			'w', 'x'. If a permission is not set, that
//			character position contains a '-'.
//
//		 	Unix Permission String Options:
//
//		      9-Character          File Access
//		      Notation             Permission Descriptions
//		-----------------------------------------------------------
//		      ---------            File - no permissions
//		      rwx------            File - read, write, & execute only for owner
//		      rwxrwx---            File - read, write, & execute for owner and group
//		      rwxrwxrwx            File - read, write, & execute for owner, group and others
//		      --x--x--x            File - execute
//		      -w--w--w-            File - write
//		      -wx-wx-wx            File - write & execute
//		      r--r--r--            File - read
//		      r-xr-xr-x            File - read & execute
//		      rw-rw-rw-            File - read & write
//		      rwxr-----            File - Owner can read, write, & execute. Group can only read;
//
//
//			Note: drwxrwxrwx - identifies permissions for
//	                        directory
//
//		errorPrefix					interface{}
//
//			This object encapsulates error prefix text which
//			is included in all returned error messages.
//			Usually, it contains the name of the calling
//			method or methods listed as a method or function
//			chain of execution.
//
//			If no error prefix information is needed, set
//			this parameter to 'nil'.
//
//			This empty interface must be convertible to one
//			of the following types:
//
//			1.	nil
//					A nil value is valid and generates an
//					empty collection of error prefix and
//					error context information.
//
//			2.	string
//					A string containing error prefix
//					information.
//
//			3.	[]string
//					A one-dimensional slice of strings
//					containing error prefix information.
//
//			4.	[][2]string
//					A two-dimensional slice of strings
//			   		containing error prefix and error
//			   		context information.
//
//			5.	ErrPrefixDto
//					An instance of ErrPrefixDto.
//					Information from this object will
//					be copied for use in error and
//					informational messages.
//
//			6.	*ErrPrefixDto
//					A pointer to an instance of
//					ErrPrefixDto. Information from
//					this object will be copied for use
//					in error and informational messages.
//
//			7.	IBasicErrorPrefix
//					An interface to a method
//					generating a two-dimensional slice
//					of strings containing error prefix
//					and error context information.
//
//			If parameter 'errorPrefix' is NOT convertible
//			to one of the valid types listed above, it will
//			be considered invalid and trigger the return of
//			an error.
//
//			Types ErrPrefixDto and IBasicErrorPrefix are
//			included in the 'errpref' software package:
//				"github.com/MikeAustin71/errpref".
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
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (fPerm *FilePermissionConfig) SetFileModeByComponents(
	entryType OsFilePermissionCode,
	unixPermissionTextStr string,
	errorPrefix interface{}) error {

	if fPerm.lock == nil {
		fPerm.lock = new(sync.Mutex)
	}

	fPerm.lock.Lock()

	defer fPerm.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FilePermissionConfig."+
			"SetFileModeByComponents()",
		"")

	if err != nil {
		return err
	}

	return new(filePermissionConfigNanobot).
		setFileModeByComponents(
			fPerm,
			entryType,
			unixPermissionTextStr,
			ePrefix.XCpy(
				"fPerm<-"))
}

// SetByFileMode
//
// Sets the permission codes for the current instance of
// FilePermissionConfig.
//
// Using input parameter 'fMode' of type 'os.FileMode'.
// If the value does not include a valid os mode
// constant, an error will be returned.
//
// If successful, this method will assign the os.FileMode
// input value to the internal data field,
// 'fPerm.fileMode'.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://pkg.go.dev/os#FileMode
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fMode						os.FileMode
//
//		An instance of os.FileMode containing file or
//		directory permission codes. These permission
//		codes will be used to reset the internal
//		FileMode data field in the current instance of
//		FilePermissionConfig.
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
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (fPerm *FilePermissionConfig) SetByFileMode(
	fMode os.FileMode,
	errorPrefix interface{}) error {

	if fPerm.lock == nil {
		fPerm.lock = new(sync.Mutex)
	}

	fPerm.lock.Lock()

	defer fPerm.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FilePermissionConfig."+
			"SetByFileMode()",
		"")

	if err != nil {
		return err
	}

	return new(filePermissionConfigNanobot).
		setByFileMode(
			fPerm,
			fMode,
			ePrefix.XCpy(
				"fPerm<-fMode"))
}

// SetFileModeByOctalDigits
//
// Sets the permissions value for the current instance of
// FilePermissionConfig. The internal FileMode data field
// (FilePermissionConfig.fileMode) is reset to the value
// represented by input parameter, 'octalFileModeCode'.
// Any previous internal FileMode value is overwritten.
//
// ------------------------------------------------------------------------
//
// # Warning
//
// In the Go Programming Language, if you initialize an
// integer with a leading zero (e.g. x:= int(0777)), than
// number ('0777') is treated as an octal value and
// converted to a decimal value. Therefore, x:= int(0777)
// will mean that 'x' is set equal to 511. If you set
// x:= int(777), x will be set equal to '777'. For purposes
// of this method enter the octal code as x:= int(777).
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	octalFileModeCode			int
//
//		This parameter contains the integer value of the
//		of the permission code which will be used to
//		initialize the current FilePermissionConfig
//		instance (FilePermissionConfig.fileMode). The
//		integer digits in 'octalFileModeCode' represent
//		the octal value for the file permissions.
//
//		If the input parameter 'octalFileModeCode'
//		contains an invalid Entry Type, an error will be
//		returned.
//
//		A partial list of valid file permission value
//	 	examples are shown as follows:
//
//	 ____________________________________________________________________________
//
//	          Input Parameter
//	              integer            Equivalent
//	 Octal    'octalFileModeCode'    Symbolic      File Access
//	 Digits        value             Notation      Permission Descriptions
//	 0000 	         0               ----------    File - no permissions
//	 0700 	       700               -rwx------    File - read, write, & execute only for owner
//	 0770 	       770               -rwxrwx---    File - read, write, & execute for owner and group
//	 0777 	       777               -rwxrwxrwx    File - read, write, & execute for owner, group and others
//	 0111 	       111               ---x--x--x    File - execute
//	 0222 	       222               --w--w--w-    File - write
//	 0333 	       333               --wx-wx-wx    File - write & execute
//	 0444 	       444               -r--r--r--    File - read
//	 0555 	       555               -r-xr-xr-x    File - read & execute
//	 0666 	       666               -rw-rw-rw-    File - read & write
//	 0740 	       740               -rwxr-----    File - Owner can read, write, & execute. Group can only read;
//	                                                      others have no permissions
//
//	 drwxrwxrwx    Directory - read, write, & execute for owner, group and others
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
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
//
//		If the input parameter 'octalFileModeCode'
//		contains an invalid Entry Type, an error will be
//		returned.
func (fPerm *FilePermissionConfig) SetFileModeByOctalDigits(
	octalFileModeCode int,
	errorPrefix interface{}) error {

	if fPerm.lock == nil {
		fPerm.lock = new(sync.Mutex)
	}

	fPerm.lock.Lock()

	defer fPerm.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FilePermissionConfig."+
			"SetFileModeByOctalDigits()",
		"")

	if err != nil {
		return err
	}

	return new(filePermissionConfigNanobot).
		setFileModeByOctalDigits(
			fPerm,
			octalFileModeCode,
			ePrefix.XCpy(
				"fPerm<-octalFileModeCode"))
}

// SetFileModeByTextCode
//
// Sets the internal FileMode data field for the current
// instance of FilePermissionConfig using input parameter
// 'filePermissionStr'. 'filePermissionStr' is a
// 10-character string containing the read, write and
// execute permissions for the three groups or user
// classes:
//
//	(1)	'Owner/User'
//
//	(2)	'Group'
//
//	(3)	'Other'
//
// The text codes used in the 'filePermissionStr' mimic
// the Unix permission codes.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://www.cyberciti.biz/faq/explain-the-nine-permissions-bits-on-files/.
//	https://en.wikipedia.org/wiki/File_system_permissions
//	https://linuxconfig.org/how-to-use-special-permissions-the-setuid-setgid-and-sticky-bits
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (fPerm *FilePermissionConfig) SetFileModeByTextCode(
	filePermissionStr string,
	errorPrefix interface{}) error {

	if fPerm.lock == nil {
		fPerm.lock = new(sync.Mutex)
	}

	fPerm.lock.Lock()

	defer fPerm.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FilePermissionConfig."+
			"SetFileModeByTextCode()",
		"")

	if err != nil {
		return err
	}

	return new(filePermissionConfigNanobot).
		setFileModeByTextCode(
			fPerm,
			filePermissionStr,
			ePrefix.XCpy(
				"fPerm<-filePermissionStr"))
}
