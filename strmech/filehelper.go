package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"os"
	fp "path/filepath"
	"sync"
	"time"
)

// FileHelper
//
// The methods associated with this type provide
// generalized file creation, management and maintenance
// utilities.
//
// 'FileHelper' is a dependency for types 'DirMgr' and
// 'FileMgr'.
type FileHelper struct {
	Input  string
	Output string
	lock   *sync.Mutex
}

// AddPathSeparatorToEndOfPathStr
//
// Receives a path string as an input parameter. If the
// last character of the path string is not a path
// separator, this method will add a path separator to
// the end of that path string and return it to the
// calling method.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathStr						string
//
//		The path string which will be analyzed to
//		determine if the last character is a path
//		separator.
//
//		If the last character is NOT a path separator,
//		this method will add a path separator to the end
//		of that path string and return it to the calling
//		method.
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
//		The path string passed as input parameter,
//		'pathStr' will be analyzed to determine if the
//		last character is a path separator.
//
//		If the last character is NOT a path separator,
//		a path separator will be added to 'pathStr' and
//		returned through this parameter.
//
//		If the last character IS a path separator, no
//		action will be taken and an exact copy of
//		'pathStr' will be returned through this parameter.
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
func (fh *FileHelper) AddPathSeparatorToEndOfPathStr(
	pathStr string,
	errorPrefix interface{}) (string, error) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"AddPathSeparatorToEndOfPathStr()",
		"")

	if err != nil {

		return "", err
	}

	return new(fileHelperProton).
		addPathSeparatorToEndOfPathStr(
			pathStr,
			ePrefix.XCpy("pathStr"))
}

// AdjustPathSlash
//
// This method will standardize path separators according to
// operating system.
//
// If input parameter 'path' contains invalid file path
// separators for the current operating system, this
// method will apply standard, compatible file path
// separator characters.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	path				string
//
//		This 'path' string will be examined for
//		non-standard path separator characters in the
//		context of the current operating system. If
//		path separator characters are found to be
//		incompatible with the current operating system,
//		they will be replaced and a valid path string
//		will be returned to the calling function.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	string
//
//		This method will replace invalid path separators
//		found in input parameter 'path' and return a
//		string containing valid path separator characters
//		for the current operating system.
func (fh *FileHelper) AdjustPathSlash(path string) string {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	return new(fileHelperAtom).adjustPathSlash(
		path)
}

// AreSameFile
//
// Compares two paths or path/file names to determine if
// they are the same and equivalent.
//
// An error will be triggered if one or both of the input
// parameters, 'pathFile1' and 'pathFile2' are empty or
// zero length strings.
//
// If the path file input parameters identify the same
// file, this method returns 'true'.
//
// The two input parameters 'pathFile1' and 'pathFile2'
// will be converted to their absolute paths before
// comparisons are applied.
//
// This path/file comparison is not affected by the fact
// one file exists and the other does not.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathFile1					string
//
//		The first of two path file strings passed to this
//		method. 'pathFile1' will be compared to
//		determined if they are in fact pointing at the
//		same file.
//
//		Both 'pathFile1' and 'pathFile2' will be
//		converted to their absolute paths being subjected
//		to comparison.
//
//		If 'pathFile1' is submitted as an empty or zero
//		length string, an error will be returned.
//
//	pathFile2					string
//
//		The first of two path file strings passed to this
//		method. 'pathFile1' will be compared to determined
//		if they are in fact pointing at the same file.
//
//		Both 'pathFile1' and 'pathFile2' will be
//		converted to their absolute paths being subjected
//		to comparison.
//
//		If 'pathFile2' is submitted as an empty or zero
//		length string, an error will be returned.
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
//		If this return parameter is set to 'true' it
//		signals that input parameters 'pathFile1' and
//		'pathFile2' are pointing at the same file.
//		Effectively, the two path/file specifications
//		are the same or equivalent.
//
//		If this return parameter is set to 'false' it
//		signals that input parameters 'pathFile1' and
//		'pathFile2' are NOT pointing at the same file.
//		Effectively, the two path/file specifications
//		are NOT equivalent.
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
func (fh *FileHelper) AreSameFile(
	pathFile1,
	pathFile2 string,
	errorPrefix interface{}) (
	bool,
	error) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"AreSameFile()",
		"")

	if err != nil {
		return false, err
	}

	return new(fileHelperNanobot).areSameFile(
		pathFile1,
		pathFile2,
		ePrefix)
}

// ChangeFileMode
//
// This method changes the file mode of an existing file
// designated by input parameter 'pathFileName'.
//
// The new os.FileMode value to which the mode will be
// changed is extracted from input parameter
// 'filePermission'.
//
// If the target file does Not exist, an error will be
// returned.
//
// If the method succeeds and the file's mode is changed,
// an error value of 'nil' is returned.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	os.FileMode https://pkg.go.dev/os#FileMode
//	https://stackoverflow.com/questions/28969455/how-to-properly-instantiate-os-filemode
//	https://www.includehelp.com/golang/os-filemode-constants-with-examples.aspx
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathFileName				string
//
//		The file mode for the file identified by
//		'pathFileName' will be modified using the new
//		file permissions provided by input parameter
//		'filePermission'.
//
//	filePermission				FilePermissionConfig
//
//		This parameter contains the new permissions which
//		will be applied to the file identified by input
//		paramter 'pathFileName'.
//
//		The easiest way to configure permissions is
//		to call FilePermissionConfig.New() with
//		a mode string ('modeStr').
//
//		The first character of the 'modeStr' designates the
//		'Entry Type'. Currently, only two 'Entry Type'
//		characters are supported. Therefore, the first
//		character in the 10-character input parameter
//		'modeStr' MUST be either a "-" indicating a file, or
//		a "d" indicating a directory.
//
//		The remaining nine characters in the 'modeStr'
//		represent unix permission bits and consist of three
//		group fields each containing 3-characters. Each
//		character in the three group fields may consist of
//		'r' (Read-Permission), 'w' (Write-Permission), 'x'
//		(Execute-Permission) or '-' signaling no permission or
//		no access allowed. A typical 'modeStr' authorizing
//		permission for full access to a file would be styled
//		as:
//
//		Directory Example: "drwxrwxrwx"
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
//	    -----------------------------------------------------
//	           Directory Mode String Permission Codes
//	    -----------------------------------------------------
//	      Directory
//			10-Character
//			 'modeStr'
//			 Symbolic		  Directory Access
//			  Format	   Permission Descriptions
//			----------------------------------------------------
//
//			d---------		no permissions
//			drwx------		read, write, & execute only for owner
//			drwxrwx---		read, write, & execute for owner and group
//			drwxrwxrwx		read, write, & execute for owner, group and others
//			d--x--x--x		execute
//			d-w--w--w-		write
//			d-wx-wx-wx		write & execute
//			dr--r--r--		read
//			dr-xr-xr-x		read & execute
//			drw-rw-rw-		read & write
//			drwxr-----		Owner can read, write, & execute. Group can only read;
//			                others have no permissions
//
//			Note: drwxrwxrwx - identifies permissions for directory
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
func (fh *FileHelper) ChangeFileMode(
	pathFileName string,
	filePermission FilePermissionConfig,
	errorPrefix interface{}) error {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"ChangeFileMode()",
		"")

	if err != nil {
		return err
	}

	return new(fileHelperNanobot).changeFileMode(
		pathFileName,
		filePermission,
		ePrefix.XCpy("pathFileName"))
}

// ChangeFileTimes
//
// This method is a wrapper for "os.Chtimes()". New
// access and modification times will be set for a
// designated file.
//
// If the path and file name do not exist, this method
// will return an error.
//
// If successful, the access and modification times for
// the target file will be changed to those specified by
// input parameters, 'newAccessTime' and 'newModTime'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathFileName				string
//
//		This string specifies the path and file name of
//		the file whose access and modification times will
//		be modified.
//
//	newAccessTime				time.Time
//
//		The file identified by input parameter
//		'pathFileName' will have its access time reset to
//		the time specified by this parameter
//		('newAccessTime').
//
//	newModTime					time.Time
//
//		The file identified by input parameter
//		'pathFileName' will have its modification time
//		reset to the time specified by this parameter
//		('newModTime').
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
func (fh *FileHelper) ChangeFileTimes(
	pathFileName string,
	newAccessTime,
	newModTime time.Time,
	errorPrefix interface{}) error {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"ChangeFileTimes()",
		"")

	if err != nil {
		return err
	}

	return new(fileHelperNanobot).
		changeFileTimes(
			pathFileName,
			newAccessTime,
			newModTime,
			ePrefix)
}

// ChangeWorkingDir
//
// Changes the current working directory to the named
// directory passed by input parameter, 'dirPath'.
//
// If there is an error, it will be of type *PathError.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dirPath						string
//
//		The name of the new current working directory.
//		The existing current working directory will be
//		changed to this directory path specification.
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
func (fh *FileHelper) ChangeWorkingDir(
	dirPath string,
	errorPrefix interface{}) error {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"ChangeWorkingDir()",
		"")

	if err != nil {
		return err
	}

	return new(fileHelperAtom).changeWorkingDir(
		dirPath,
		ePrefix.XCpy(
			"<-dirPath"))
}

// CleanDirStr
//
// Cleans and formats a directory string. When a string
// contains both a directory and a file name, this method
// will extract the pure directory string.
//
// ----------------------------------------------------------------
//
// # Examples
//
//	dirName = '/someDir/xt_dirmgr_01_test.go' returns "./someDir"
//
//	dirName = '../dir1/dir2/fileName.ext'     returns "../dir1/dir2"
//
//	dirName = 'fileName.ext'                  returns "" isEmpty = true
//
//	dirName = 'xt_dirmgr_01_test.go'          returns "" isEmpty = true
//
//	dirName = '../dir1/dir2/'                 returns "../dir1/dir2"
//
//	dirName = '../../../dir1/'                returns '../../../dir1'
//
//	dirName = '../dir1/dir2/filename.ext'     returns "../dir1/dir2"
//
//	dirName = '../dir1/dir2/.git'             returns "../dir1/dir2"
//
//	dirName = 'somevalidcharacters'           returns "./somevalidchracters"
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
//	returnedDirName				string
//
//		Contains the formatted directory string extracted
//		from input paramter, 'dirNameStr'.
//
//		If 'dirNameStr' is empty or invalid,
//		'returnedDirName' will be set to zero length (an
//		empty string).
//
//	isEmpty						bool
//
//		This boolean value will be set to 'true' under
//		the following conditions:
//
//		(1)	Input parameter 'dirNameStr' is empty or
//			invalid.
//
//		(2) Input parameter 'dirNameStr' does NOT contain
//			a directory string. This occurs if
//			'dirNameStr' only contains a file name and/or
//			extension.
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
func (fh *FileHelper) CleanDirStr(
	dirNameStr string,
	errorPrefix interface{}) (
	returnedDirName string,
	isEmpty bool,
	err error) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	returnedDirName = ""

	isEmpty = true

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"CleanDirStr()",
		"")

	if err != nil {
		return returnedDirName, isEmpty, err
	}

	returnedDirName,
		isEmpty,
		err = new(fileHelperNanobot).cleanDirStr(
		dirNameStr,
		ePrefix)

	return returnedDirName, isEmpty, err
}

// CleanFileNameExtStr
//
// Cleans up a file name extension string and returns a
// string which contains only the file name and extension.
//
// ----------------------------------------------------------------
//
// # Usage
//
//	fileNameExt =	'../dir1/dir2/fileName.ext'
//	Method Returns:	"fileName.ext" and isEmpty=false
//
//	fileNameExt =	'fileName.ext"
//	Method Returns:	"fileName.ext" and isEmpty=false
//
//	fileNameExt =	"../filesfortest/newfilesfortest/" (actually exists on disk)
//	Method Returns:	"" and isEmpty=true and error = nil
//
//	fileNameExt =	'../dir1/dir2/'
//	Method Returns:	"" and isEmpty=true
//
//	fileNameExt =	'../filesfortest/newfilesfortest/newerFileForTest_01'
//	Method Returns:	"newerFileForTest_01" and isEmpty=false
//
//	fileNameExt =	'../filesfortest/newfilesfortest/.gitignore'
//	Method Returns:	".gitignore" and isEmpty=false
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fileNameExtStr				string
//
//		Typically, this string contains a directory/folder
//		string or a file path string. This method will
//		extract the file name and extension from this
//		string returning it through return parameter,
//		'returnedFileNameExt'.
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
//	returnedFileNameExt			string
//
//		This method receives a directory or file path
//		string and extracts the file name and extension.
//		Only the file name and extension is extracted and
//		returned through this parameter.
//
//	isEmpty						bool
//
//		If the returned file name and extension string
//		('returnedFileNameExt') is an empty, or zero
//		length string, this returned boolean value is
//		set to true.
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
func (fh *FileHelper) CleanFileNameExtStr(
	fileNameExtStr string,
	errorPrefix interface{}) (
	returnedFileNameExt string,
	isEmpty bool,
	err error) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	isEmpty = true

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"CleanFileNameExtStr()",
		"")

	if err != nil {
		return returnedFileNameExt, isEmpty, err
	}

	returnedFileNameExt,
		isEmpty,
		err = new(fileHelperNanobot).
		cleanFileNameExtStr(
			fileNameExtStr,
			ePrefix.XCpy(
				"returnedFileNameExt<-fileNameExtStr"))

	return returnedFileNameExt, isEmpty, err
}

// CleanPathStr
//
// Wrapper Function for filepath.Clean().
//
// See: https://golang.org/pkg/path/filepath/#Clean
//
// This method returns the shortest path name equivalent
// to path by purely lexical processing. It applies the
// following rules iteratively until no further
// processing can be done:
//
//  1. Replace multiple Separator elements with a single
//     one.
//
//  2. Eliminate each '.' path name element (the current
//     directory).
//
//  3. Eliminate each inner '..' path name element (the
//     parent directory) along with the 'non-..' element
//     that precedes it.
//
//  4. Eliminate '..' elements that begin a rooted path;
//     that is, replace '/..' with '/' at the beginning
//     of a path, assuming Separator is '/'.
//
// The returned path ends in a slash only if it
// represents a root directory, such as "/" on Unix or
// `C:\` on Windows.
//
// Finally, any occurrences of slash are replaced by
// Separator.
//
// If the result of this process is an empty string,
// Clean returns the string ".".
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathStr						string
//
//		The path or directory/folder string to be
//		cleaned. For a detail explanation of the
//		cleaning process, see above.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	string
//
//		This method returns a 'cleaned' version of the
//		file path or directory/folder string passed as
//		input parameter, 'pathStr'.
//
//		For a detail explanation of the path 'cleaning'
//		process, see above.
func (fh *FileHelper) CleanPathStr(pathStr string) string {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	return fp.Clean(pathStr)
}

// ConsolidateErrors
//
// Receives an array of type error and converts the
// individual error elements to a consolidated single
// instance of type 'error' which is returned to the
// caller.
//
// Multiple errors are separated by a new line character
// when returned as single consolidate 'error'.
//
// If the length of the error array is zero, this method
// returns nil.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errs						[]error
//
//		An array to type 'error'. The errors contained in
//		this array are consolidated and returned as a single
//		instance of type 'error'. Each 'error' element in this
//		array is automatically separated by a new line
//		character when returned as a single type of 'error'.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		A single instance of type 'error' is returned containing
//		all the consolidated individual errors contained in the
//		input parameter 'errs'. 'errs' is an array of type
//		'error'.
func (fh *FileHelper) ConsolidateErrors(errs []error) error {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	return new(strMechPreon).
		consolidateErrors(errs)
}

// ConvertDecimalToOctal
//
// Utility routine to convert a decimal (base 10)
// numeric value to an octal (base 8) numeric value. Useful in
// evaluating 'os.FileMode' values and associated constants.
//
//	Reference:
//	 https://www.cloudhadoop.com/2018/12/golang-example-convertcast-octal-to.html
//
// ----------------------------------------------------------------
//
// Usage:
//
//	 initialDecimalValue := 511
//	 expectedOctalValue := 777
//
//		var actualOctalValue int
//
//	 actualOctalValue =
//			new(numberConversionsMolecule).
//			convertDecimalToOctal(initialDecimalValue)
//
//	 'actualOctalValue' is now equal to integer value '777'.
//
// ----------------------------------------------------------------
//
// Warning:
//
// In the Go Programming Language, if you initialize an
// integer with a leading zero (e.g. x:= int(0777)), than
// number ('0777') is treated as an octal value and converted
// to a decimal value. Therefore, x:= int(0777) will mean
// that 'x' is set equal to 511. If you set x:= int(777), x
// will be set equal to '777'. '777' is the correct result.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	decimalNumber				int
//
//		The decimal or base 10 number which will be
//		converted to an octal value.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	octalNumber					int
//
//		The converted octal (Base 8) number. Input
//		parameter 'decimalNumber' (Base 10) will be
//		converted to an octal number and returned
//		through this parameter.
func (fh *FileHelper) ConvertDecimalToOctal(decimalNumber int) int {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	return new(numberConversionsMolecule).
		convertDecimalToOctal(decimalNumber)
}

// ConvertOctalToDecimal
//
// Utility routine to convert an octal (base 8) numeric
// value to a decimal (base 10) numeric value. Useful in
// evaluating 'os.FileMode' values and associated
// constants.
//
//	Reference:
//
//	https://www.cloudhadoop.com/2018/12/golang-example-convertcast-octal-to.html
//
// ----------------------------------------------------------------
//
// Usage:
//
//		expectedDecimalValue := 511
//		initialOctalValue := 777
//		actualDecimalValue :=
//			new(numberConversionsMolecule).
//				convertOctalToDecimal(initialOctalValue)
//
//	 actualDecimalValue is now equal to integer value, '511'.
//
// ----------------------------------------------------------------
//
// Warning:
//
// In the Go Programming Language, if you initialize an
// integer with a leading zero (e.g. x:= int(0777)), than
// number ('0777') is treated as an octal value and
// converted to a decimal value. Therefore, x:= int(0777)
// will mean that 'x' is set equal to 511. If you set
// x:= int(777), x will be set equal to '777'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	octalNumber					int
//
//		The octal or Base 8 number which will be
//		converted to a decimal numeric value (Base 10).
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	int
//
//		The converted decimal (Base 10) number. Input
//		parameter 'octalNumber' (Base 8) will be
//		converted to a decimal number and returned
//		through this parameter.
func (fh *FileHelper) ConvertOctalToDecimal(octalNumber int) int {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	return new(numberConversionsMolecule).
		convertOctalToDecimal(octalNumber)
}

// CopyFileByIoByLink
//
// Copies a file from source to destination using one of
// two techniques.
//
// First, this method will attempt to copy the designated
// file by means of creating a new destination file and
// using "io.Copy(out, in)" to copy the contents. This is
// accomplished by calling 'CopyFileByIo()'.
// If  the call to 'CopyFileByIo()' fails, this method
// will attempt a second copy method.
//
// The second attempt to copy the designated file will be
// accomplished by creating a 'hard link' to the source
// file. The second, 'hard link', attempt will call
// method, 'CopyFileByLink()'.
//
// If both attempted file copy operations fail, an error
// will be returned.
//
// See: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	src							string
//
//		This string holds the file name and path for the
//		source file which will be copied to the
//		destination file identified by input parameter,
//		'dst'.
//
//	dst							string
//
//		This string holds the file name and path for the
//		destination file. The source file identified by
//		input parameter 'src' will be copied to this
//		destination file ('dst').
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
func (fh *FileHelper) CopyFileByIoByLink(
	src string,
	dst string) (err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"FileHelper."+
			"CopyFileByIoByLink()",
		"")

	if err != nil {
		return err
	}

	return new(fileHelperDirector).copyFileByIoByLink(
		src,
		dst,
		ePrefix)
}

// CopyFileByLinkByIo
//
// Copies a file from source to destination using one of
// two techniques.
//
// First, this method will attempt to copy the designated
// file by means of creating a 'hard link' to the source file.
// The 'hard link' attempt will call 'CopyFileByLink()'.
//
// If that 'hard link' operation fails, this method will call
// 'CopyFileByIo()'.
//
// CopyFileByIo() will create a new destination file and attempt
// to write the contents of the source file to the new destination
// file using "io.Copy(out, in)".
//
// If both attempted file copy operations fail, an error will be
// returned.
//
// See: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	src							string
//
//		This string holds the file name and path for the
//		source file which will be copied to the
//		destination file identified by input parameter,
//		'dst'.
//
//	dst							string
//
//		This string holds the file name and path for the
//		destination file. The source file identified by
//		input parameter 'src' will be copied to this
//		destination file ('dst').
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
func (fh *FileHelper) CopyFileByLinkByIo(
	src,
	dst string,
	errorPrefix interface{}) error {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"CopyFileByLinkByIo()",
		"")

	if err != nil {
		return err
	}

	fHelperMech := new(fileHelperMechanics)

	err = fHelperMech.copyFileByLink(
		src,
		dst,
		ePrefix)

	if err == nil {
		return err
	}

	var err2 error

	// Copy by Link Failed. Try CopyFileByIo()
	err2 = fHelperMech.copyFileByIo(
		src,
		dst,
		ePrefix)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: After Copy By Link failed, an error was returned by fh.CopyFileByIo(src, dst).\n"+
			"src='%v'\n"+
			"dst='%v'\n"+
			"Error='%v'\n",
			ePrefix,
			src,
			dst,
			err2.Error())

		return err
	}

	return err
}

// CopyFileByLink
//
// Copies a file from source to destination by means of
// creating a 'hard link' to the source file,
// "os.Link(src, dst)".
//
// Note: This method of copying files does NOT create a
// new destination file and write the contents of the
// source file to destination file. (See CopyFileByIo
// Below).  Instead, this method performs the copy
// operation by creating a hard symbolic link to the
// source file.
//
// By creating a 'linked' file, changing the contents
// of one file will be reflected in the second. The
// two linked files are 'mirrors' of each other.
//
// Consider using CopyFileByIo() if this 'mirror' feature
// causes problems.
//
// "os.Link(src, dst)" is the only method employed to
// copy a designated file. If "os.Link(src, dst)" fails,
// an err is returned.
//
// See: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// REQUIREMENT: The destination Path must previously
// exist. The destination file need NOT exist as it will
// be created. If the destination file currently exists,
// it will first be deleted and a new linked file will be
// created.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	src							string
//
//		This string holds the file name and path for the
//		source file. which will be copied to the
//	 	destination file passed through input parameter
//	 	'dst'.
//
//	dst							string
//
//		This string holds the file name and path for the
//		destination file. The source file passed through
//	 	input parameter 'src' which will be copied to the
//	 	destination file identified by this parameter,
//	 	'dst'.
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
func (fh *FileHelper) CopyFileByLink(
	src string,
	dst string,
	errorPrefix interface{}) error {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"CopyFileByLink()",
		"")

	if err != nil {
		return err
	}

	return new(fileHelperMechanics).copyFileByLink(
		src,
		dst,
		ePrefix)
}

// CopyFileByIo
//
// Copies file from source path and file name to
// destination path and file name.
//
// Reference:
// https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// Note: Unlike the method CopyFileByLink above, this
// method does NOT rely on the creation of symbolic
// links. Instead, a new destination file is created and
// the contents of the source file are written to the new
// destination file using "io.Copy()".
//
// "io.Copy()" is the only method used to copy the
// designated source file. If this method fails, an error
// is returned.
//
// If source file is equivalent to the destination file,
// no action will be taken and no error will be returned.
//
// If the destination file does not exist, this method
// will create. However, it will NOT create the
// destination directory. If the destination directory
// does NOT exist, this method will abort the copy
// operation and return an error.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	sourceFile					string
//
//		This string holds the path and/or file name of
//	 	the source file. This source file will be copied
//		to the destination file.
//
//	destinationFile				string
//
//		This string holds the path and/or the file name
//		of the destination file. The source file taken
//		from input parameter 'sourceFile' will be copied
//		to this destination file.
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
func (fh *FileHelper) CopyFileByIo(
	sourceFile string,
	destinationFile string,
	errorPrefix interface{}) error {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"CopyFileByIo()",
		"")

	if err != nil {
		return err
	}

	return new(fileHelperMechanics).copyFileByIo(
		sourceFile,
		destinationFile,
		ePrefix)
}

// CreateFile
//
// A wrapper function for os.Create().
//
// This method will 'create' the file designated by input
// parameter 'pathFileName'. 'pathFileName' should
// consist of a valid path and file name. The file name
// may consist of a file name and file extension or simply
// a file name.
//
// If the path component of input parameter 'pathFileName'
// does not exist or invalid, the call to low level method
// os.Create() will a type *PathError. This *PathError
// will be returned through the return parameter
// 'lowLevelErr' for detailed review by the calling
// function.
//
// If successful, this method will return a valid pointer
// to the new file (type 'os.File') and an error value of
// 'nil'.
//
// os.Create creates or truncates the named file. If the
// file already exists, it is truncated. If the file does
// not exist, it is created with mode 0666 (before umask).
// If successful, methods on the returned File can be used
// for I/O; the associated file descriptor has mode O_RDWR.
// If there is an os.Create error, it will be of type
// *PathError.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	os.File:	https://pkg.go.dev/os#File
//	os.Create:	https://pkg.go.dev/os#Create
//	PathError:	https://pkg.go.dev/os#PathError
//
// ----------------------------------------------------------------
//
// # WARNING
//
//	If the file identified by 'pathFileName' previously
//	exists, it will be truncated.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathFileName				string
//
//		This string holds the path and file name for the
//		file which will be created by this method.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in the error message returned by
//		'msgError'.
//
//		Usually, 'errorPrefix' contains the name of the
//	 	calling method or methods listed as a method or
//	  	function chain of execution.
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
//	ptrOsFile					*os.File
//
//	If this method completes successfully, a pointer to
//	the newly created file will be returned by parameter
//	'ptrOsFile'.
//
//	msgError					error
//
//		'msgError' is a standard error containing
//		a brief high level message explaining the
//		error condition in narrative text.
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
//
//		If 'msgError' is 'nil' it signals that the
//		file creation operation succeeded.
//
//		NOTE: Not all returned errors have associated
//		low level errors ('lowLevelErr'). Always check
//		'msgError'. 'msgError' could be non-nil while
//		'lowLevelErr' is 'nil'.
//
//	lowLevelErr					error
//
//		If the call to os.Create fails it will return an
//		error of type *PathError. This error may be
//		unpacked to reveal additional technical details
//		regarding the failure to create a file.
//
//		If an error is returned by os.Create it will be
//		returned in its original form through parameter,
//		'lowLevelErr'.
//
//		If no *PathError occurs, 'lowLevelErr' will be set
//	 	to 'nil'.
func (fh *FileHelper) CreateFile(
	pathFileName string,
	errorPrefix interface{}) (
	ptrOsFile *os.File,
	msgError error,
	lowLevelErr error) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ptrOsFile = nil

	ePrefix,
		msgError = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"CreateFile()",
		"")

	if msgError != nil {
		return ptrOsFile, msgError, lowLevelErr
	}

	ptrOsFile,
		msgError,
		lowLevelErr = new(fileHelperAtom).
		createFile(
			pathFileName,
			ePrefix.XCpy(
				"pathFileName"))

	return ptrOsFile, msgError, lowLevelErr
}

// DeleteAllFilesInDirectory
//
// Deletes every file in a single directory.
//
// Only files in the target directory are deleted. Files
// is subdirectories or parent directories are NOT
// deleted.
//
// ----------------------------------------------------------------
//
// # WARNING!
//
//	Be careful. This method will delete all files in a
//	single directory.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	directoryPath				string
//
//		This string holds the target path or directory
//		name. All files residing in the directory will be
//		deleted.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in the error message returned by
//		'msgError'.
//
//		Usually, 'errorPrefix' contains the name of the
//	 	calling method or methods listed as a method or
//	  	function chain of execution.
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
//	msgError					error
//
//		'msgError' is a standard error containing
//		a brief high level message explaining the
//		error condition in narrative text.
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If 'msgError' is returned with a value of
//		'nil', it signals that the deletion operation
//		succeeded.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
//
//		NOTE: Not all returned errors have associated
//		low level errors ('lowLevelErr'). Always check
//		'msgError'. 'msgError' could be non-nil while
//		'lowLevelErr' is 'nil'.
//
//	lowLevelErr					error
//
//		If the call to os.Remove fails it will return an
//		error of type *PathError. This error may be
//		unpacked to reveal additional technical details
//		regarding the failure to create a file.
//
//		If an error is returned by os.Remove it will be
//		returned in its original form through parameter,
//		'lowLevelErr'.
//
//		If no *PathError occurs, 'lowLevelErr' will be set
//		to 'nil'.
func (fh *FileHelper) DeleteAllFilesInDirectory(
	directoryPath string,
	errorPrefix interface{}) (
	msgError error,
	lowLevelErr error) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		msgError = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"DeleteAllFilesInDirectory()",
		"")

	if msgError != nil {
		return msgError, lowLevelErr
	}

	msgError,
		lowLevelErr = new(fileHelperAtom).
		deleteAllFilesInDirectory(
			directoryPath,
			ePrefix.XCpy(
				"directoryPath"))

	return msgError, lowLevelErr
}

// DeleteDirFile
//
// Wrapper function for os.Remove.
//
// os.Remove removes the named file or an empty directory.
// If there is an error, it will be of type *PathError,
// and it will be returned in parameter 'lowLevelErr'.
//
// ----------------------------------------------------------------
//
// # WARNING!
//
//	Be careful. This method will delete files and
//	empty directories.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://pkg.go.dev/os#Remove
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathFile					string
//
//		This string holds the path and/or file name
//		identifying the file or directory to be deleted.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in the error message returned by
//		'msgError'.
//
//		Usually, 'errorPrefix' contains the name of the
//	 	calling method or methods listed as a method or
//	  	function chain of execution.
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
//	msgError					error
//
//		'msgError' is a standard error containing
//		a brief high level message explaining the
//		error condition in narrative text.
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If 'msgError' is returned with a value of
//		'nil', it signals that the deletion operation
//		succeeded.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
//
//		NOTE: Not all returned errors have associated
//		low level errors ('lowLevelErr'). Always check
//		'msgError'. 'msgError' could be non-nil while
//		'lowLevelErr' is 'nil'.
//
//	lowLevelErr					error
//
//		If the call to os.Remove fails it will return an
//		error of type *PathError. This error may be
//		unpacked to reveal additional technical details
//		regarding the failure to create a file.
//
//		If an error is returned by os.Remove it will be
//		returned in its original form through parameter,
//		'lowLevelErr'.
//
//		If no *PathError occurs, 'lowLevelErr' will be set
//		to 'nil'.
func (fh *FileHelper) DeleteDirFile(
	pathFile string,
	errorPrefix interface{}) (
	msgError error,
	lowLevelErr error) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		msgError = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"DeleteDirFile()",
		"")

	if msgError != nil {
		return msgError, lowLevelErr
	}

	msgError,
		lowLevelErr = new(fileHelperNanobot).
		deleteDirFile(
			pathFile,
			ePrefix.XCpy(
				"pathFile"))

	return msgError, lowLevelErr
}

// DeleteDirPathAll
//
// Wrapper function for os.RemoveAll.
//
// os.RemoveAll removes path and any children it
// contains. It removes everything it can but returns the
// first error it encounters. If the path does not exist,
// os.RemoveAll takes no action and returns 'nil'
// (no error). If os.RemoveAll does return an error, it
// will be of type *PathError.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://www.geeksforgeeks.org/how-to-remove-all-directories-and-files-in-golang/
//	https://thedeveloperblog.com/go/os-remove-go
//	https://pkg.go.dev/os#RemoveAll
//
// ----------------------------------------------------------------
//
// # Examples
//
//		(1)	D:\T08\x294_1\x394_1\x494_1  is a directory path
//			that currently exists and contains files.
//
//		(2) Call DeleteDirPathAll("D:\\T08\\x294_1")
//
//	 	(3) Upon return from method DeleteDirPathAll():
//			a.	Deletion Results:
//	    		Directory D:\T08\x294_1\x394_1\x494_1 and any
//	   			files in the 'x494_1' directory are deleted
//
//	    		Directory D:\T08\x294_1\x394_1\ and any files
//	   			in the 'x394_1' directory are deleted
//
//	    		Directory D:\T08\x294_1\ and any files in the
//	   			'x294_1' directory are deleted
//
//			b.	The Parent Path 'D:\T08' and any files in that
//				parent path 'D:\T08' directory are unaffected
//				and continue to exist.
//
// ----------------------------------------------------------------
//
// # WARNING
//
//	This method has the capability to delete entire
//	directory trees.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathDir						string
//
//		This strings holds the path and directory where
//		the deletion operation will occur.
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
//	msgError					error
//
//		'msgError' is a standard error containing
//		a brief high level message explaining the
//		error condition in narrative text.
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If 'msgError' is returned with a value of
//		'nil', it signals that the deletion operation
//		succeeded.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
//
//		NOTE: Not all returned errors have associated
//		low level errors ('lowLevelErr'). Always check
//		'msgError'. 'msgError' could be non-nil while
//		'lowLevelErr' is 'nil'.
//
//	lowLevelErr					error
//
//		If the call to os.RemoveAll fails it will return
//		an error of type *PathError. This error may be
//		unpacked to reveal additional technical details
//		regarding the failure to create a file.
//
//		If an error is returned by os.RemoveAll it will
//	 	be returned in its original form through parameter,
//		'lowLevelErr'.
//
//		If no *PathError occurs, 'lowLevelErr' will be set
//		to 'nil'.
func (fh *FileHelper) DeleteDirPathAll(
	pathDir string,
	errorPrefix interface{}) (
	msgError error,
	lowLevelErr error) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		msgError = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"DeleteDirPathAll()",
		"")

	if msgError != nil {
		return msgError, lowLevelErr
	}

	msgError,
		lowLevelErr = new(fileHelperNanobot).
		deleteDirPathAll(
			pathDir,
			ePrefix.XCpy(
				"pathDir"))

	return msgError, lowLevelErr
}

// DeleteFilesWalkDirectory
//
// This method 'walks' the directory tree searching for
// files which match the file selection criteria
// specified by input parameter 'fileSelectCriteria'.
// When a file matching said 'fileSelectCriteria' is
// found, that file is deleted.
//
// This method returns file information on files deleted.
//
// If a file matches the File Selection Criteria
// ('fileSelectCriteria') it is deleted and its file
// information is recorded in the returned instance of
// DirectoryDeleteFileInfo,
// 'DirectoryDeleteFileInfo.DeletedFiles'.
//
// By the way, if ALL the file selection criterion are
// set to zero values or 'Inactive', then ALL FILES in
// the directory are selected, deleted and returned in
// the field, 'DirectoryDeleteFileInfo.DeletedFiles'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	If all of the file selection criterion in the
//	FileSelectionCriteria object are 'Inactive' or
//	'Not Set' (set to their zero or default values), then
//	all the files processed in the directory tree WILL BE
//	DELETED!
//
//	Information on the deleted file will be returned in the
//	file manager collection,
//	'DirectoryDeleteFileInfo.DeletedFiles'.
//
//		Example:
//			FileNamePatterns  = ZERO Length Array
//			filesOlderThan    = time.Time{}
//			filesNewerThan    = time.Time{}
//
//	In this example, all the selection criterion are
//	'Inactive' and therefore all the files encountered
//	in the target directory will be selected and returned
//	as 'Found Files'.
//
//	This same effect can be achieved by simply creating an
//	empty file selection instance:
//
//		FileSelectionCriteria{}
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	startPath					string
//
//		A string consisting of the starting path or
//		directory from which the file search operation
//		will commence.
//
//	fileSelectCriteria			FileSelectionCriteria
//
//		This input parameter should be configured with
//		the desired file selection criteria. Files
//		matching this criteria will be returned as
//		'Found Files'.
//
//		If file 'fileSelectCriteria' is uninitialized
//		(FileSelectionCriteria{}), all directories the
//		'startPath' will be searched, and all files
//		within those directories WILL BE DELETED.
//
//		type FileSelectionCriteria struct {
//		  FileNamePatterns     []string    // An array of strings containing File Name Patterns
//		  FilesOlderThan       time.Time   // Match files with older modification date times
//		  FilesNewerThan       time.Time   // Match files with newer modification date times
//		  SelectByFileMode     FilePermissionConfig // Match file mode (os.FileMode).
//		  SelectCriterionMode  FileSelectCriterionMode // Specifies 'AND' or 'OR' selection mode
//		}
//
//		The FileSelectionCriteria type allows for configuration of single or multiple file
//		selection criterion. The 'SelectCriterionMode' can be used to specify whether the
//		file must match all, or any one, of the active file selection criterion.
//
//		Elements of the FileSelectionCriteria Type are described below:
//
//		FileNamePatterns []string
//
//			An array of strings which may define one or more search
//			patterns. If a file name matches any one of the search
//			pattern strings, it is deemed to be a 'match' for the
//			search pattern criterion.
//
//			Example Patterns:
//			 FileNamePatterns = []string{"*.log"}
//			 FileNamePatterns = []string{"current*.txt"}
//			 FileNamePatterns = []string{"*.txt", "*.log"}
//
//			If this string array has zero length or if
//			all the strings are empty strings, then this
//			file search criterion is considered 'Inactive'
//			or 'Not Set'.
//
//		FilesOlderThan  time.Time
//
//			This date time type is compared to file modification
//			date times in order to determine whether the file is
//			older than the 'FilesOlderThan' file selection
//			criterion. If the file is older than the
//			'FilesOlderThan' date time, that file is considered
//			a 'match'	for this file selection criterion.
//
//			If the value of 'FilesOlderThan' is set to time zero,
//			the default value for type time.Time{}, then this
//			file selection criterion is considered to be 'Inactive'
//			or 'Not Set'.
//
//		FilesNewerThan   time.Time
//
//			This date time type is compared to the file modification
//			date time in order to determine whether the file is newer
//			than the 'FilesNewerThan' file selection criterion. If
//			the file modification date time is newer than the
//			'FilesNewerThan' date time, that file is considered a
//			'match' for this file selection criterion.
//
//			If the value of 'FilesNewerThan' is set to time zero,
//			the default value for type time.Time{}, then this
//			file selection criterion is considered to be 'Inactive'
//			or 'Not Set'.
//
//		SelectByFileMode  FilePermissionConfig
//
//			Type FilePermissionConfig encapsulates an os.FileMode. The
//			file selection criterion allows for the selection of files
//			by File Mode.
//
//			File modes are compared to the value of 'SelectByFileMode'.
//			If the File Mode for a given file is equal to the value of
//	 		'SelectByFileMode', that file is considered to be a 'match'
//	 		for this file selection criterion. Examples for setting
//	 		SelectByFileMode are shown as follows:
//
//			fsc := FileSelectionCriteria{}
//
//			err = fsc.SelectByFileMode.SetByFileMode(os.FileMode(0666))
//
//			err = fsc.SelectByFileMode.SetFileModeByTextCode("-r--r--r--")
//
//		SelectCriterionMode FileSelectCriterionMode
//
//		This parameter selects the manner in which the file selection
//		criteria above are applied in determining a 'match' for file
//		selection purposes. 'SelectCriterionMode' may be set to one of
//		two constant values:
//
//		(1) FileSelectCriterionMode(0).ANDSelect()
//
//			File selected if all active selection criteria
//			are satisfied.
//
//			If this constant value is specified for the file selection mode,
//			then a given file will not be judged as 'selected' unless all
//			the active selection criterion are satisfied. In other words, if
//			three active search criterion are provided for 'FileNamePatterns',
//			'FilesOlderThan' and 'FilesNewerThan', then a file will NOT be
//			selected unless it has satisfied all three criterion in this example.
//
//		(2) FileSelectCriterionMode(0).ORSelect()
//
//			File selected if any active selection criterion is satisfied.
//
//			If this constant value is specified for the file selection mode,
//			then a given file will be selected if any one of the active file
//			selection criterion is satisfied. In other words, if three active
//			search criterion are provided for 'FileNamePatterns', 'FilesOlderThan'
//			and 'FilesNewerThan', then a file will be selected if it satisfies any
//			one of the three criterion in this example.
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
//	DirectoryDeleteFileInfo
//
//		If successful, files matching the file selection criteria input
//		parameter shown above will be deleted and returned in a
//		'DirectoryDeleteFileInfo' object. The file manager
//		'DirectoryDeleteFileInfo.DeletedFiles' contains information on all files
//		deleted during this operation.
//
//		Note:
//		It is a good idea to check the returned field 'DirectoryTreeInfo.ErrReturns'
//		to determine if any internal system errors were encountered while processing
//		the directory tree.
//
//		__________________________________________________________________________________________________
//
//		type DirectoryDeleteFileInfo struct {
//
//			StartPath             string
//
//		  		The starting path or directory for the file
//		  		search.
//
//			Directories           DirMgrCollection
//
//				Directory Manager instances found during the
//				directory tree search.
//
//			DeletedFiles          FileMgrCollection
//
//				Contains File Managers for Deleted Files matching
//				file selection criteria.
//
//			ErrReturns            []error
//
//				Internal System errors encountered during the search
//				and file deletion operations. This includes type
//				*PathError objects created by low level system
//				function calls.
//
//			FileSelectCriteria    FileSelectionCriteria
//
//				The File Selection Criteria submitted as an
//				input parameter to this method.
//		}
//
//		__________________________________________________________________________________________________
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
func (fh *FileHelper) DeleteFilesWalkDirectory(
	startPath string,
	fileSelectCriteria FileSelectionCriteria,
	errorPrefix interface{}) (
	DirectoryDeleteFileInfo,
	error) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	deleteFilesInfo := DirectoryDeleteFileInfo{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"DeleteFilesWalkDirectory()",
		"")

	if err != nil {
		return deleteFilesInfo, err
	}

	return new(fileHelperMechanics).
		deleteFilesWalkDirectory(
			startPath,
			fileSelectCriteria,
			ePrefix)
}

// DoesDirectoryExist
//
// This method tests for the existence of a directory
// path.
//
// If the directory path does exist, this method also
// returns an instance of FileInfoPlus containing
// detailed information on the directory.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dirPath						string
//
//		This strings holds the directory path which will
//		be examined to determine if it actually exists.
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
//	dirPathDoesExist			bool
//
//		If the directory path specified by input
//		parameter 'dirPath' actually exists, this
//		returned boolean value will be set to 'true'.
//
//	fInfoPlus					FileInfoPlus
//
//		If the directory path specified by input
//		parameter 'dirPath' actually exists, this
//		returned instance of FileInfoPlus will be
//		populated with detailed information on that
//		directory.
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
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fh *FileHelper) DoesDirectoryExist(
	dirPath string,
	errorPrefix interface{}) (
	dirPathDoesExist bool,
	fInfoPlus FileInfoPlus,
	err error) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"DoesDirectoryExist()",
		"")

	return new(fileHelperAtom).doesDirectoryExist(
		dirPath,
		ePrefix)
}

// DoesFileExist
//
// Returns a boolean value designating whether the passed
// file name exists.
//
// This method does not differentiate between Path Errors
// and Non-Path Errors returned by os.Stat(). The method
// only returns a boolean value.
//
// If a Non-Path Error is returned by os.Stat(), this
// method will classify the file as "Does NOT Exist" and
// return a value of 'false'.
//
// For a more granular test of whether a file exists, see
// method FileHelper.DoesThisFileExist().
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathFileName				string
//
//		This string holds the name of a path and file
//		name. This method will determine this path and
//		file name actually exists.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		If the path and file name passed by input
//		parameter 'pathFileName' is verified to actually
//		exist, this return parameter will be set to 'true'.
//
//		If the path and file name do not exist or if an
//		error occurs, this method returns 'false'.
func (fh *FileHelper) DoesFileExist(pathFileName string) bool {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	return new(fileHelperNanobot).
		doesFileExist(
			pathFileName)
}

// DoesFileInfoExist
//
// This method returns a boolean value indicating whether
// the path and file name passed to the function actually
// exists.
//
// If the file actually exists, the function will return
// the associated FileInfo structure.
//
// If 'pathFileName' does NOT exist, return parameters
// 'doesFInfoExist' will be set to 'false', 'fInfo' will
// be set to 'nil' and the returned error value will be
// 'nil'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathFileName				string
//
//		This path and file name will be evaluated to
//		determine if the file actually exist.
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
//	doesFInfoExist				bool
//
//		If the file identified in input parameter
//		'pathFileName' actually exists, this return
//		parameter will be set to 'true'.
//
//		If this file does NOT exist, this parameter
//		will return 'false'.
//
//	fInfo						os.FileInfo
//
//		If the file identified in input parameter
//		'pathFileName' actually exists, this return
//		parameter will be configured with a copy of
//		that file's	os.FileInfo object.
//
//		Type os.FileInfo contains data elements
//		describing a file.
//
//		type FileInfo interface {
//
//			Name() string
//				base name of the file
//
//			Size() int64
//				length in bytes for regular files;
//				system-dependent for others
//
//			Mode() FileMode
//				file mode bits
//
//			ModTime() time.Time
//				modification time
//
//			IsDir() bool
//				abbreviation for Mode().IsDir()
//
//			Sys() any
//				underlying data source (can return nil)
//		}
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
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fh *FileHelper) DoesFileInfoExist(
	pathFileName string,
	errorPrefix interface{}) (
	doesFInfoExist bool,
	fInfo os.FileInfo,
	err error) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	doesFInfoExist = false

	fInfo = nil

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"DoesFileInfoExist()",
		"")

	if err != nil {
		return doesFInfoExist, fInfo, err
	}

	doesFInfoExist,
		fInfo,
		err = new(fileHelperNanobot).
		doesFileInfoExist(
			pathFileName,
			ePrefix)

	return doesFInfoExist, fInfo, err
}

// DoesFileInfoPlusExist
//
// This method returns a boolean value indicating whether
// the path and file name passed to the function actually
// exists.
//
// If 'pathFileName' does NOT exist, return parameters
// 'doesFInfoExist' will be set to 'false' and the
// returned error value will be 'nil'.
//
// If the file actually exists, the function will return
// the associated FileInfoPlus structure.
//
// As its name implies, the returned type FileInfoPlus
// provides methods and information which go beyond the
// os.FileInfo interface in describing the file identified
// by 'pathFileName'. For detailed information on the
// features and capabilities offered by type FileInfoPlus,
// see the source code documentation detail located in
// source code file 'fileinfoplus.go'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathFileName				string
//
//		This path and file name will be evaluated to
//		determine if the file actually exist.
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
//	doesFInfoExist				bool
//
//		If the file identified in input parameter
//		'pathFileName' actually exists, this return
//		parameter will be set to 'true'.
//
//		If this file does NOT exist, this parameter
//		will return 'false'.
//
//	fInfoPlus					FileInfoPlus
//
//		If the file identified in input parameter
//		'pathFileName' actually exists, this return
//		parameter will be configured with an instance of
//		FileInfoPlus providing detailed information on
//		the file identified by input parameter
//		'pathFileName'.
//
//		Type FileInfoPlus conforms to the os.FileInfo
//		interface. This structure will store os.FileInfo
//	 	information plus additional information related
//	 	to a file or directory.
//
//		type os.FileInfo interface {
//
//				Name() string
//					base name of the file
//
//				Size() int64
//					length in bytes for regular files;
//					system-dependent for others
//
//				Mode() FileMode
//					file mode bits
//
//				ModTime() time.Time
//					modification time
//
//				IsDir() bool
//					abbreviation for Mode().IsDir()
//
//				Sys() any
//					underlying data source (can return nil)
//		}
//
//		See the detailed documentation for Type
//		FileInfoPlus in the source file,
//		'fileinfoplus.go'.
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
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fh *FileHelper) DoesFileInfoPlusExist(
	pathFileName string,
	errorPrefix interface{}) (
	doesFInfoExist bool,
	fInfoPlus FileInfoPlus,
	err error) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	doesFInfoExist = false

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"DoesFileInfoPlusExist()",
		"")

	if err != nil {
		return doesFInfoExist, fInfoPlus, err
	}

	doesFInfoExist,
		fInfoPlus,
		err = new(fileHelperNanobot).
		doesFileInfoPlusExist(
			pathFileName,
			ePrefix)

	return doesFInfoExist, fInfoPlus, err
}

// DoesPathFileExist
//
// A helper method which public methods use to determine
// whether a path and file does or does not exist.
//
// This method calls os.Stat(dirPath) which may return
// one of two types of errors:
//
//		(1)	A Non-Path Error
//
//			This is an error which is not path related. It
//			signals some other type of error which makes
//			it impossible to determine if the path actually
//			exists. These types of errors generally related
//			to as "access denied" situations, but there may
//			be other reasons behind non-path errors. If a
//			non-path error is returned, no valid existence
//			test can be performed on the file path.
//
//	    or
//
//		(2)	A Path Error
//
//			This type of error indicates that the path
//			definitely does not exist.
//
//		To deal with these types of errors, this method will
//		test path existence up to three times before returning
//		a non-path error.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	filePath					string
//
//		This string holds the path and file name. This
//		path and file name will be tested to determine if
//		they exist on disk.
//
//	preProcessCode			PreProcessPathCode
//
//		Type PreProcessPathCode is an enumeration. The
//		enumeration value passed as 'preProcessCode' is
//		used to determine the type of setup and
//		configuration that will be applied to 'filePath'.
//
//		This enumeration has three options:
//
//			PreProcessPathCode(0).None()
//			PreProcessPathCode(0).PathSeparator()
//			PreProcessPathCode(0).AbsolutePath()
//
//		Alternatively, the enumeration value could be
//		expressed using shorthand notation:
//
//			PreProcPathCode.None()
//			PreProcPathCode.PathSeparator()
//			PreProcPathCode.AbsolutePath()
//
//		PreProcPathCode.None()
//
//			This code specifies that no preprocessing
//			operation will be applied to the 'filePath'
//			string before determining whether the path
//			and file name identified by 'filePath'
//			actually exists.
//
//		PreProcPathCode.PathSeparator()
//
//			This code specifies that the file path
//			('filePath) will be screened for invalid path
//			separator characters. Invalid path separator
//			characters will be replaced with valid path
//			separator characters for the current
//			operating system. After completion of this
//			preprocessing operation, the revised path and
//			file name identified by 'filePath' will be
//			tested to determine if it actually exists on
//			disk.
//
//		PreProcPathCode.AbsolutePath()
//
//			This code specifies the file name and file
//			path ('filePath') will first be converted to
//			an absolute path and file name. Afterward,
//			the converted path and file name ('filePath')
//			will be tested to determine if it actually
//			exists on disk.
//
//	filePathTitle				string
//
//		This string constitutes a label which will be
//		used to describe parameter 'filePath' in any
//		error message returned to the user.
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
//	nonPathError				error
//
//		If this method completes successfully, the
//		returned 'nonPathError' is set equal to 'nil'.
//
//		This is an error which is not path related. It
//		signals some other type of error which makes
//		it impossible to determine if the path actually
//		exists. These types of errors generally related
//		to as "access denied" situations, but there may
//		be other reasons behind non-path errors. If a
//		non-path error is returned, no valid existence
//		test can be performed on the file path.
//
//		If errors are encountered during processing, the
//		returned 'nonPathError' will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be attached to the
//	 	beginning of the error message.
func (fh *FileHelper) DoesPathFileExist(
	filePath string,
	preProcessCode PreProcessPathCode,
	filePathTitle string,
	errorPrefix interface{}) (
	absFilePath string,
	filePathDoesExist bool,
	fInfo FileInfoPlus,
	nonPathError error) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		nonPathError = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"DoesPathFileExist()",
		"")

	if nonPathError != nil {

		return absFilePath, filePathDoesExist, fInfo, nonPathError
	}

	absFilePath,
		filePathDoesExist,
		fInfo,
		nonPathError = new(fileHelperMolecule).
		doesPathFileExist(
			filePath,
			preProcessCode,
			ePrefix,
			filePathTitle)

	return absFilePath, filePathDoesExist, fInfo, nonPathError
}

// DoesStringEndWithPathSeparator
//
// This method returns a boolean value of 'true' if the
// path string input paramter ('pathStr') ends with a
// valid Path Separator ('/' or '\' depending on the
// operating system).
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathStr						string
//
//		The file or directory path in this string will be
//		analyzed to determine if the string ends with a
//		path separator character.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		If input parameter 'pathStr' contains a string
//		which ends with a Path Separator character, this
//		return parameter will be set to 'true'.
//
//		If 'pathStr' does NOT end with a Path Separator
//		character, this parameter will return 'false'.
func (fh *FileHelper) DoesStringEndWithPathSeparator(
	pathStr string) bool {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	return new(fileHelperAtom).
		doesStringEndWithPathSeparator(pathStr)
}

// DoesThisFileExist
//
// Returns a boolean value signaling whether the path and
// file name represented by input parameter,
// 'pathFileName', does in fact exist. Unlike the similar
// method FileHelper.DoesFileExist(), this method returns
// an error in the case of Non-Path errors associated with
// 'pathFileName'.
//
// Non-Path errors may arise for a variety of reasons, but
// the most common is associated with 'access denied'
// situations.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathFileName				string
//
//		This string contains the path and file name which
//		will be analyzed to determine if in fact they
//		actually exist.
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
//	pathFileNameDoesExist		bool
//
//		If the path and file name identified by input
//		parameter 'pathFileName' actually exists, this
//		parameter will return a boolean value of 'true'.
//
//		If the path and file name do NOT exist, a value
//		of 'false' will be returned.
//
//	nonPathError				error
//
//		If this method completes successfully, the
//		returned 'nonPathError' is set equal to 'nil'.
//
//		This is an error which is not path related. It
//		signals some other type of error which makes
//		it impossible to determine if the path actually
//		exists. These types of errors generally related
//		to as "access denied" situations, but there may
//		be other reasons behind non-path errors. If a
//		non-path error is returned, no valid existence
//		test can be performed on the file path.
//
//		If errors are encountered during processing, the
//		returned 'nonPathError' will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be attached to the
//	 	beginning of the error message.
func (fh *FileHelper) DoesThisFileExist(
	pathFileName string,
	errorPrefix interface{}) (
	pathFileNameDoesExist bool,
	nonPathError error) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	pathFileNameDoesExist = false

	ePrefix,
		nonPathError = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"DoesThisFileExist()",
		"")

	if nonPathError != nil {

		return pathFileNameDoesExist, nonPathError
	}

	pathFileNameDoesExist,
		nonPathError = new(fileHelperAtom).
		doesThisFileExist(
			pathFileName,
			ePrefix)

	return pathFileNameDoesExist, nonPathError
}

// FilterFileName
//
// This method will determine whether a file described by
// an os.FileInfo object meets one or more of three file
// selection criteria:
//
//	(1)	A string or name pattern match.
//		('FileNamePatterns')
//
//	(2) A file modification time which is older than the
//		'FilesOlderThan' parameter.
//
//	(3)	A file modification time which is newer than the
//		'FilesNewerThan' parameter.
//
// The three file selection criterion are applied to the
// file name (fileInfo.Name()) passed as input parameter
// 'fileInfo'.
//
// If the three file search criteria are all set to their
// 'zero' or empty values, the no selection filter is
// applied and all files are deemed to be a match for the
// selection criteria ('isMatchedFile=true').
//
// If a given selection criterion is set to a zero value,
// then that criterion is defined as 'not set' and
// therefore is not used or applied in determining
// whether a file  'matches' the specified selection
// criteria.
//
// If a given criterion specifies a non-zero value, then
// that criterion is defined as 'set' and the file
// information must comply with that criterion in order
// to be judged as a match ('isMatchedFile=true').
//
// Again, if none of the three criterion are 'set', then
// all files are judged as matched ('isMatchedFile=true').
//
// When one or more of the selection criteria shown above
// are set, they are applied collectively or individually
// depending on input parameter
// 'fileSelectCriteria.SelectCriterionMode' which is of
// type 'FileSelectCriterionMode'.
//
// 'fileSelectCriteria.SelectCriterionMode' has two
// settings:
//
//	(1) FileSelectCriterionMode(0).ANDSelect()
//
//		A File is selected if all active selection
//		criteria are satisfied.
//
//		If this constant value is specified for the file
//		selection mode, then a given file will not be
//		judged as 'selected' unless all the active
//		selection criterion are satisfied. In other words,
//		if three active search criterion are provided for
//		'FileNamePatterns', 'FilesOlderThan' and
//		'FilesNewerThan', then a file will NOT be selected
//		unless it has satisfied all three criterion.
//
//	(2) FileSelectCriterionMode(0).ORSelect()
//
//		A File is selected if any active selection
//		criterion is satisfied.
//
//		If this constant value is specified for the file
//		selection mode, then a given file will be
//		selected if any one of the active file selection
//		criterion is satisfied. In other words, if three
//		active search criterion are provided for
//		'FileNamePatterns', 'FilesOlderThan' and
//		'FilesNewerThan', then a file will be selected if
//		it satisfies any one of these three active
//		criterion.
//
// The two settings for
// 'fileSelectCriteria.SelectCriterionMode' shown above,
// control the manner in which search criteria are
// applied by the file search algorithm.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fileInfo					os.FileInfo
//
//		Type os.FileInfo contains data elements
//		describing a file.
//
//		type FileInfo interface {
//
//			Name() string
//				base name of the file
//
//			Size() int64
//				length in bytes for regular files;
//				system-dependent for others
//
//			Mode() FileMode
//				file mode bits
//
//			ModTime() time.Time
//				modification time
//
//			IsDir() bool
//				abbreviation for Mode().IsDir()
//
//			Sys() any
//				underlying data source (can return nil)
//		}
//
//	fileSelectCriteria			FileSelectionCriteria
//
//		The FileSelectionCriteria type allows for
//		configuration of single or multiple file
//		selection criterion. The 'SelectCriterionMode'
//		can be used to specify whether the file must match all,
//		or any one, of the active file selection criterion.
//
//		type FileSelectionCriteria struct {
//		  FileNamePatterns     []string
//			An array of strings containing File Name Patterns
//
//		  FilesOlderThan       time.Time
//			Match files with older modification date times
//
//		  FilesNewerThan       time.Time
//			Match files with newer modification date times
//
//		  SelectByFileMode     FilePermissionConfig
//			Match file mode (os.FileMode).
//
//		  SelectCriterionMode  FileSelectCriterionMode
//			Specifies 'AND' or 'OR' selection mode
//		}
//
//		Elements of the FileSelectionCriteria Type are
//		described below:
//
//		FileNamePatterns []string
//
//			An array of strings which may define one or more search
//			patterns. If a file name matches any one of the search
//			pattern strings, it is deemed to be a 'match' for the
//			search pattern criterion.
//
//			Example Patterns:
//			 FileNamePatterns = []string{"*.log"}
//			 FileNamePatterns = []string{"current*.txt"}
//			 FileNamePatterns = []string{"*.txt", "*.log"}
//
//			If this string array has zero length or if
//			all the strings are empty strings, then this
//			file search criterion is considered 'Inactive'
//			or 'Not Set'.
//
//		FilesOlderThan  time.Time
//
//			This date time type is compared to file modification
//			date times in order to determine whether the file is
//			older than the 'FilesOlderThan' file selection
//			criterion. If the file is older than the
//			'FilesOlderThan' date time, that file is considered
//			a 'match'	for this file selection criterion.
//
//			If the value of 'FilesOlderThan' is set to time zero,
//			the default value for type time.Time{}, then this
//			file selection criterion is considered to be 'Inactive'
//			or 'Not Set'.
//
//		FilesNewerThan   time.Time
//
//			This date time type is compared to the file modification
//			date time in order to determine whether the file is newer
//			than the 'FilesNewerThan' file selection criterion. If
//			the file modification date time is newer than the
//			'FilesNewerThan' date time, that file is considered a
//			'match' for this file selection criterion.
//
//			If the value of 'FilesNewerThan' is set to time zero,
//			the default value for type time.Time{}, then this
//			file selection criterion is considered to be 'Inactive'
//			or 'Not Set'.
//
//		SelectByFileMode  FilePermissionConfig
//
//			Type FilePermissionConfig encapsulates an os.FileMode. The
//			file selection criterion allows for the selection of files
//			by File Mode.
//
//			File modes are compared to the value of 'SelectByFileMode'.
//			If the File Mode for a given file is equal to the value of
//	 		'SelectByFileMode', that file is considered to be a 'match'
//	 		for this file selection criterion. Examples for setting
//	 		SelectByFileMode are shown as follows:
//
//			fsc := FileSelectionCriteria{}
//
//			err = fsc.SelectByFileMode.SetByFileMode(os.FileMode(0666))
//
//			err = fsc.SelectByFileMode.SetFileModeByTextCode("-r--r--r--")
//
//		SelectCriterionMode FileSelectCriterionMode
//
//		This parameter selects the manner in which the file selection
//		criteria above are applied in determining a 'match' for file
//		selection purposes. 'SelectCriterionMode' may be set to one of
//		two constant values:
//
//		(1) FileSelectCriterionMode(0).ANDSelect()
//
//			File selected if all active selection criteria
//			are satisfied.
//
//			If this constant value is specified for the file selection mode,
//			then a given file will not be judged as 'selected' unless all
//			the active selection criterion are satisfied. In other words, if
//			three active search criterion are provided for 'FileNamePatterns',
//			'FilesOlderThan' and 'FilesNewerThan', then a file will NOT be
//			selected unless it has satisfied all three criterion in this example.
//
//		(2) FileSelectCriterionMode(0).ORSelect()
//
//			File selected if any active selection criterion is satisfied.
//
//			If this constant value is specified for the file selection mode,
//			then a given file will be selected if any one of the active file
//			selection criterion is satisfied. In other words, if three active
//			search criterion are provided for 'FileNamePatterns', 'FilesOlderThan'
//			and 'FilesNewerThan', then a file will be selected if it satisfies any
//			one of the three criterion in this example.
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
//	isMatchedFile				bool
//
//		If this return parameter is set to 'true' it
//		signals that the file identified by input
//		parameter 'fileInfo' has matched the specified
//		file selection criteria.
//
//		If 'isMatchedFile' is returned with a value of
//		'false', it means that the file identified by
//		input parameter 'fileInfo' does NOT match the
//		specified file selection criteria.
//
//	msgError					error
//
//		'msgError' is a standard error containing
//		a brief high level message explaining the
//		error condition in narrative text.
//
//		If this method completes successfully,
//		'msgError' will be set to 'nil'.
//
//		If errors are encountered during processing, the
//		returned 'msgError' will contain an appropriate
//		error message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
//
//		NOTE:
//
//		Not all returned errors have associated
//		low level system errors ('lowLevelErr'). Always
//		check 'msgError'. 'msgError' could be non-nil
//		while 'lowLevelErr' is 'nil'.
//
//	lowLevelErr					error
//
//		If calls to low level system methods fail, those
//		methods will often return a specialized error
//		containing a packet of subsidiary error messages.
//		These types of errors may be unpacked to reveal
//		additional technical details and information
//		regarding the causes of method failure.
//
//		If a low level system error is identified, it will
//		be returned in its original form through parameter
//		'lowLevelErr'.
//
//		If no low level system error is identified,
//		'lowLevelErr' will be set to 'nil'.
func (fh *FileHelper) FilterFileName(
	fileInfo os.FileInfo,
	fileSelectionCriteria FileSelectionCriteria,
	errorPrefix interface{}) (
	isMatchedFile bool,
	msgError error,
	lowLevelErr error) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	isMatchedFile = false

	ePrefix,
		msgError = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FilterFileName."+
			"FilterFileName()",
		"")

	if msgError != nil {

		return isMatchedFile, msgError, lowLevelErr
	}

	isMatchedFile,
		msgError,
		lowLevelErr = new(fileHelperAtom).
		filterFileName(
			fileInfo,
			fileSelectionCriteria,
			ePrefix)

	return isMatchedFile, msgError, lowLevelErr
}

// FindFilesInPath
//
// This method will apply a search pattern to files and
// directories in the path designated by input parameter,
// 'pathName'. If the files and or directory names match
// the input parameter, 'fileSearchPattern' they will be
// returned in an array of strings.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
// The names returned in the string array may consist of
// both files and directory names, depending on the
// specified file search pattern, 'fileSearchPattern'.
//
// ----------------------------------------------------------------
//
// This method uses the "path/filepath" function, 'Glob'.
//
// Glob Reference:
//
//	https://golang.org/pkg/path/filepath/#Glob
//	https://pkg.go.dev/path/filepath#Match
//
// The File matching patterns depend on the 'go'
// "path/filepath" function, 'Match'.
//
// Match Reference:
//
//	https://golang.org/pkg/path/filepath/#Match
//
// Note:
// This method will NOT search subdirectories. It will
// return the names of subdirectories existing in the
// designated, 'pathName', depending on the file search
// pattern, 'fileSearchPattern', passed as an input
// parameter.
//
// If input parameters 'pathName' or 'fileSearchPattern'
// are empty strings or consist of all space characters,
// this method will return an error.
//
//	Example 'fileSearchPattern' values:
//	      "*"     = Returns all files and directories (everything)
//	      "*.*"   = Returns files which have a file extension
//	      "*.txt" = Returns only files with a "txt" file extension
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathName					string
//
//		This string holds the path which will be searched
//		in an attempt to find the files specified by the
//		file search pattern.
//
//	fileSearchPattern			string
//
//		This string holds the file search pattern which
//		will be applied to filter, select and return the
//		target files residing in the path specified by
//		input parameter 'pathName'.
//
//		For more search pattern examples, see the
//		documentation at:
//
//		https://golang.org/pkg/path/filepath/#Glob
//		https://pkg.go.dev/path/filepath#Match
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
//	[]string
//
//		An array of strings containing the file names
//		matching the file search criteria specified
//		by input parameter 'fileSearchPattern'.
//
//		The names returned in this string array may
//		consist of both files and directory names,
//		depending on the specified file search
//		pattern, 'fileSearchPattern'.
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
func (fh *FileHelper) FindFilesInPath(
	pathName,
	fileSearchPattern string,
	errorPrefix interface{}) ([]string, error) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"FindFilesInPath()",
		"")

	if err != nil {
		return []string{}, err
	}

	return new(fileHelperNanobot).
		findFilesInPath(
			pathName,
			fileSearchPattern,
			ePrefix.XCpy("pathName"))
}

// FindFilesWalkDirectory
//
// This method returns file information on files residing
// in a specified directory tree identified by the input
// parameter, 'startPath'.
//
// This method 'walks the directory tree' locating all
// files in the directory tree which match the file
// selection criteria submitted as input parameter,
// 'fileSelectCriteria'.
//
// If a file matches the File Selection Criteria, it is
// included in the returned field:
//
//	'DirectoryTreeInfo.FoundFiles'
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	If all of the file selection criterion in the
//	FileSelectionCriteria object are 'Inactive' or
//	'Not Set' (set to their zero or default values), then
//	all the files processed in the directory tree will be
//	selected and returned as 'Found Files'.
//
//	  Example:
//	     FileNamePatterns  = ZERO Length Array
//	     filesOlderThan    = time.Time{}
//	     filesNewerThan    = time.Time{}
//
//	  In this example, all the selection criterion are
//	  'Inactive' and therefore all the files encountered
//	  in the target directory will be selected and returned
//	  as 'Found Files' stored in the following member
//	  variable:
//
//		'DirectoryTreeInfo.FoundFiles'.
//
//	  This same effect can be achieved by simply creating
//	  an empty file selection instance:
//
//	          FileSelectionCriteria{}
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	startPath					string
//
//		A string consisting of the starting path or
//		directory from which the find files search
//		operation will commence.
//
//	fileSelectCriteria			FileSelectionCriteria
//
//		This input parameter should be configured with
//		the desired file selection criteria. Files
//		matching this criteria will be returned as
//		'Found Files'.
//
//		If file 'fileSelectCriteria' is uninitialized
//		(FileSelectionCriteria{}), all directories and
//		files will be returned from the 'startPath'
//
//		type FileSelectionCriteria struct {
//		  FileNamePatterns     []string
//			An array of strings containing File Name Patterns
//
//		  FilesOlderThan       time.Time
//			Match files with older modification date times
//
//		  FilesNewerThan       time.Time
//			Match files with newer modification date times
//
//		  SelectByFileMode     FilePermissionConfig
//			Match file mode (os.FileMode).
//
//		  SelectCriterionMode  FileSelectCriterionMode
//			Specifies 'AND' or 'OR' selection mode
//		}
//
//		Elements of the FileSelectionCriteria Type are
//		described below:
//
//		FileNamePatterns []string
//
//			An array of strings which may define one or more search
//			patterns. If a file name matches any one of the search
//			pattern strings, it is deemed to be a 'match' for the
//			search pattern criterion.
//
//			Example Patterns:
//			 FileNamePatterns = []string{"*.log"}
//			 FileNamePatterns = []string{"current*.txt"}
//			 FileNamePatterns = []string{"*.txt", "*.log"}
//
//			If this string array has zero length or if
//			all the strings are empty strings, then this
//			file search criterion is considered 'Inactive'
//			or 'Not Set'.
//
//		FilesOlderThan  time.Time
//
//			This date time type is compared to file modification
//			date times in order to determine whether the file is
//			older than the 'FilesOlderThan' file selection
//			criterion. If the file is older than the
//			'FilesOlderThan' date time, that file is considered
//			a 'match'	for this file selection criterion.
//
//			If the value of 'FilesOlderThan' is set to time zero,
//			the default value for type time.Time{}, then this
//			file selection criterion is considered to be 'Inactive'
//			or 'Not Set'.
//
//		FilesNewerThan   time.Time
//
//			This date time type is compared to the file modification
//			date time in order to determine whether the file is newer
//			than the 'FilesNewerThan' file selection criterion. If
//			the file modification date time is newer than the
//			'FilesNewerThan' date time, that file is considered a
//			'match' for this file selection criterion.
//
//			If the value of 'FilesNewerThan' is set to time zero,
//			the default value for type time.Time{}, then this
//			file selection criterion is considered to be 'Inactive'
//			or 'Not Set'.
//
//		SelectByFileMode  	FilePermissionConfig
//
//			Type FilePermissionConfig encapsulates an os.FileMode. The
//			file selection criterion allows for the selection of files
//			by File Mode.
//
//			File modes are compared to the value of 'SelectByFileMode'.
//			If the File Mode for a given file is equal to the value of
//	 		'SelectByFileMode', that file is considered to be a 'match'
//	 		for this file selection criterion. Examples for setting
//	 		SelectByFileMode are shown as follows:
//
//			fsc := FileSelectionCriteria{}
//
//			err = fsc.SelectByFileMode.SetByFileMode(os.FileMode(0666))
//
//			err = fsc.SelectByFileMode.SetFileModeByTextCode("-r--r--r--")
//
//			SelectCriterionMode 	FileSelectCriterionMode
//
//			This parameter selects the manner in which the file selection
//			criteria above are applied in determining a 'match' for file
//			selection purposes. 'SelectCriterionMode' may be set to one of
//			two constant values:
//
//			(1) FileSelectCriterionMode(0).ANDSelect()
//
//				File selected if all active selection criteria
//				are satisfied.
//
//				If this constant value is specified for the file selection mode,
//				then a given file will not be judged as 'selected' unless all
//				the active selection criterion are satisfied. In other words, if
//				three active search criterion are provided for 'FileNamePatterns',
//				'FilesOlderThan' and 'FilesNewerThan', then a file will NOT be
//				selected unless it has satisfied all three criterion in this example.
//
//			(2) FileSelectCriterionMode(0).ORSelect()
//
//				File selected if any active selection criterion is satisfied.
//
//				If this constant value is specified for the file selection mode,
//				then a given file will be selected if any one of the active file
//				selection criterion is satisfied. In other words, if three active
//				search criterion are provided for 'FileNamePatterns', 'FilesOlderThan'
//				and 'FilesNewerThan', then a file will be selected if it satisfies any
//				one of the three criterion in this example.
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
// ------------------------------------------------------------------------
//
// Return Values:
//
//	DirectoryTreeInfo
//
//		If successful, files matching the file selection
//		criteria input parameter shown above will be
//		returned in a 'DirectoryTreeInfo' object. The
//		field 'DirectoryTreeInfo.FoundFiles' contains
//		information on all the files in the specified
//		path or directory tree which match the file
//		selection criteria.
//
//		Note:
//
//		It's a good idea to check the returned field
//		'DirectoryTreeInfo.ErrReturns' to determine if
//		any internal system errors were encountered while
//		processing the directory tree.
//
//	        ________________________________________________
//
//	        type DirectoryTreeInfo struct {
//
//	          StartPath             string
//	          	The starting path or directory for the
//	          	file search.
//
//	          dirMgrs               []DirMgr
//	          	Directories found during directory tree
//	          	search are stored here.
//
//	          FoundFiles            []FileWalkInfo
//				Files matching the file search selection
//				criteria are stored here.
//
//	          ErrReturns            []string
//				Internal System errors encountered during
//				the file search operation are stored here
//				as text messages.
//
//	          FileSelectCriteria    FileSelectionCriteria
//	          	The File Selection Criteria submitted as an
//				input parameter to this method.
//	        }
//
//	        ________________________________________________
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
//
//		BE ADVISED
//
//		It's a good idea to check the returned field
//		'DirectoryTreeInfo.ErrReturns' to determine if
//		any internal system errors were encountered while
//		processing the directory tree.
func (fh *FileHelper) FindFilesWalkDirectory(
	startPath string,
	fileSelectCriteria FileSelectionCriteria,
	errorPrefix interface{}) (
	DirectoryTreeInfo,
	error) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	findFilesInfo := DirectoryTreeInfo{}

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"FindFilesWalkDirectory()",
		"")

	if err != nil {
		return findFilesInfo, err
	}

	return new(fileHelperMechanics).
		findFilesWalkDirectory(
			startPath,
			fileSelectCriteria,
			ePrefix)
}

// GetAbsCurrDir
//
// Returns the absolute path of the current working
// directory.
//
// The current work directory is determined by a call to
// 'os.Getwd()'. 'os.Getwd()' returns a rooted path name
// corresponding to the current directory.
//
// If the current directory can be reached via multiple
// paths (due to symbolic links), 'os.Getwd()' may return
// any one of them.
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
//		Returns the absolute path of the current working
//		directory.
//
//		The current work directory is determined by a
//		call to 'os.Getwd()'. 'os.Getwd()' returns a
//		rooted path name corresponding to the current
//		directory.
//
//		If the current working directory can be reached
//		via multiple paths (due to symbolic links),
//		'os.Getwd()' may return any one of them.
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
func (fh *FileHelper) GetAbsCurrDir(
	errorPrefix interface{}) (
	string,
	error) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"GetAbsCurrDir()",
		"")

	if err != nil {
		return "", err
	}

	return new(fileHelperAtom).
		getAbsCurrDir(
			ePrefix)
}

// GetAbsPathFromFilePath
//
// Receives a string containing the path, file name
// and extension.
//
// This method will then return the absolute value of
// that path, file name and file extension.
//
// "An absolute or full path points to the same location
// in a file system, regardless of the current working
// directory. To do that, it must include the root
// directory."
//
//	Wikipedia
//
// This method therefore converts path element contained
// in input parameter 'filePath' to an absolute path.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://en.wikipedia.org/wiki/Path_(computing)#Absolute_and_relative_paths
//	https://en.wikipedia.org/wiki/Path_(computing)
//	https://pkg.go.dev/path/filepath@go1.20.1#Abs
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	filePath					string
//
//		This strings contains the path, file name and
//		file extension. This method will convert the path
//		element to an absolute path.
//
//		"An absolute or full path points to the same
//	 	location in a file system, regardless of the
//	 	current working directory. To do that, it must
//	 	include the root directory."
//
//			Wikipedia
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
//		This string returns the absolute file path.
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
func (fh *FileHelper) GetAbsPathFromFilePath(
	filePath string,
	errorPrefix interface{}) (
	string,
	error) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"GetAbsPathFromFilePath()",
		"")

	if err != nil {
		return "", err
	}

	return new(fileHelperMolecule).
		getAbsPathFromFilePath(
			filePath,
			ePrefix.XCpy(
				"<-filePath"))
}

// GetCurrentDir
//
// This wrapper function calls os.Getwd().
//
// Getwd returns a rooted path name corresponding to the
// current directory. If the current directory can be
// reached via multiple paths (due to symbolic links),
// Getwd may return any one of them.
//
// In this context the returned current directory is
// the current working directory.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://pkg.go.dev/os#Getwd
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
//		This method returns a string containing the
//		current working directory path.
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
func (fh *FileHelper) GetCurrentDir(
	errorPrefix interface{}) (
	string,
	error) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"GetCurrentDir()",
		"")

	if err != nil {
		return "", err
	}

	return new(fileHelperElectron).
		getCurrentDir(ePrefix)
}

// GetDotSeparatorIndexesInPathStr
//
// Returns an array of integers representing the indexes
// of dot ('.') characters located in input parameter
// 'pathStr'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathStr						string
//
//		Typically, this is a file path or directory path
//		string. This method will analyze this string and
//		report all the identified dot ('.') characters
//		in an integer array of string indexes returned by
//		this method.
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
//	[]int
//
//		An array of integers specifying the string
//		indexes of the dot ('.') characters located in
//		input parameter string 'pathStr'.
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
func (fh *FileHelper) GetDotSeparatorIndexesInPathStr(
	pathStr string,
	errorPrefix interface{}) (
	[]int,
	error) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	var dotIdxs []int

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"GetDotSeparatorIndexesInPathStr()",
		"")

	if err != nil {
		return dotIdxs, err
	}

	dotIdxs,
		err = new(fileHelperAtom).
		getDotSeparatorIndexesInPathStr(
			pathStr,
			ePrefix.XCpy("pathStr"))

	return dotIdxs, nil
}

// GetExecutablePathFileName
//
// Returns the path and file name of the executable that
// started the current process.
//
// This executable path and file name is generated by a
// call to os.Executable().
//
// os.Executable() returns the path (and file) name for
// the executable that started the current process. There
// is no guarantee that the path is still pointing to the
// correct executable. If a symlink was used to start the
// process, depending on the operating system, the result
// might be the symlink or the path it pointed to. If a
// stable result is needed, path/filepath.EvalSymlinks
// might help.
//
// Executable returns an absolute path unless an error
// occurs.
//
// The main use case for this method is finding resources
// located relative to an executable.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://pkg.go.dev/os#Executable
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
//		If this method completes successfully, this
//		string will return the path and file name of the
//		executable that started the current process.
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
func (fh *FileHelper) GetExecutablePathFileName(
	errorPrefix interface{}) (
	string,
	error) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"GetExecutablePathFileName()",
		"")

	if err != nil {
		return "", err
	}

	return new(fileHelperElectron).
		getExecutablePathFileName(
			ePrefix)
}

// GetFileExtension
//
// Returns the File Extension with the dot. If there is
// no File Extension an empty string is returned (NO dot
// included). If the returned File Extension is an empty
// string, the returned parameter 'isEmpty' is set equal
// to 'true'.
//
// When an extension is returned in the 'ext' variable,
// this extension includes a leading dot.
//
//	Example: '.txt'
//
// ----------------------------------------------------------------
//
// # Usage
//
//		Actual File Name Plus Extension: "newerFileForTest_01.txt"
//		        Returned File Extension: ".txt"
//
//		Actual File Name Plus Extension: "newerFileForTest_01"
//		        Returned File Extension: ""
//
//		Actual File Name Plus Extension: ".gitignore"
//		        Returned File Extension: ""
//
//		-------------------------------------
//
//		fh := new(FileHelper)
//
//		commonDir :=
//			fh.AdjustPathSlash(".\\xt_dirmgr_01_test.go")
//
//		result, isEmpty, err := fh.GetFileExtension(commonDir)
//
//	 result is now equal to ".go"
//		isEmpty is now equal to 'false'
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathFileNameExt				string
//
//		This string contains the path, file name and file
//		extension. This method will extract and return
//		the file extension.
//
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
//	ext							string
//
//		If this method completes successfully, this
//		parameter will return the file extension
//		extracted from the path, file name and file
//		extension passed as input parameter
//		'pathFileNameExt'.
//
//	isEmpty						bool
//
//		If the returned file extension ('ext') is an
//		empty string, 'isEmpty' is set to 'true'.
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
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fh *FileHelper) GetFileExtension(
	pathFileNameExt string,
	errorPrefix interface{}) (
	ext string,
	isEmpty bool,
	err error) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ext = ""

	isEmpty = true

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"GetFileExtension()",
		"")

	if err != nil {
		return ext, isEmpty, err
	}

	ext,
		isEmpty,
		err = new(fileHelperNanobot).
		getFileExtension(
			pathFileNameExt,
			ePrefix)

	return ext, isEmpty, err
}

// GetFileInfo
//
// Wrapper function for os.Stat(). This method returns an
// instance of os.FileInfo associated with the file
// identified by input parameter 'pathFileName'.
//
// If the file identified by 'pathFileName' does NOT
// exist, an error will be returned.
//
// The returned os.FileInfo provides useful information
// about the file identified by 'pathFileName'. See the
// documentation detail for return parameter 'osFileInfo'
// shown below.
//
// This method is similar to:
//
//	FileHelper.DoesFileInfoExist()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathFileName				string
//
//		This string contains the path and file name of
//		the file which is described by the os.FileInfo
//		object returned by this method.
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
//	osFileInfo					os.FileInfo
//
//		This returned instance of Type os.FileInfo
//		contains data elements describing the file
//		identified by input parameter 'pathFileName'.
//
//		type os.FileInfo interface {
//
//				Name() string
//					base name of the file
//
//				Size() int64
//					length in bytes for regular files;
//					system-dependent for others
//
//				Mode() FileMode
//					file mode bits
//
//				ModTime() time.Time
//					modification time
//
//				IsDir() bool
//					abbreviation for Mode().IsDir()
//
//				Sys() any
//					underlying data source (can return nil)
//		}
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
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fh *FileHelper) GetFileInfo(
	pathFileName string,
	errorPrefix interface{}) (
	osFileInfo os.FileInfo,
	err error) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"GetFileInfo()",
		"")

	if err != nil {

		return osFileInfo, err
	}

	osFileInfo,
		err = new(fileHelperNanobot).
		getFileInfo(
			pathFileName,
			ePrefix)

	return osFileInfo, err
}

// GetFileInfoPlus
//
// Wrapper function for os.Stat(). This method returns an
// instance of FileInfoPlus associated with the file
// identified by input parameter 'pathFileName'.
//
// If the file identified by 'pathFileName' does NOT
// exist, an error will be returned.
//
// As its name implies, the returned type FileInfoPlus
// provides methods and information which go beyond the
// os.FileInfo interface in describing the file identified
// by 'pathFileName'. For detailed information on the
// features and capabilities offered by type FileInfoPlus,
// see the source code documentation detail located in
// source code file 'fileinfoplus.go'.
//
// This method is similar to:
//
//	FileHelper.DoesFileInfoPlusExist()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathFileName				string
//
//		This string contains the path and file name of
//		the file which is described by the os.FileInfo
//		object returned by this method.
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
//	fileInfoPlus				FileInfoPlus
//
//		This returned instance of Type FileInfoPlus
//		contains data elements describing the file
//		identified by input parameter 'pathFileName'.
//
//		Type FileInfoPlus conforms to the os.FileInfo
//		interface. This structure will store os.FileInfo
//	 	information plus additional information related
//	 	to a file or directory.
//
//		type os.FileInfo interface {
//
//				Name() string
//					base name of the file
//
//				Size() int64
//					length in bytes for regular files;
//					system-dependent for others
//
//				Mode() FileMode
//					file mode bits
//
//				ModTime() time.Time
//					modification time
//
//				IsDir() bool
//					abbreviation for Mode().IsDir()
//
//				Sys() any
//					underlying data source (can return nil)
//		}
//
//		See the detailed documentation for Type
//		FileInfoPlus in the source file,
//		'fileinfoplus.go'.
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
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fh *FileHelper) GetFileInfoPlus(
	pathFileName string,
	errorPrefix interface{}) (
	fileInfoPlus FileInfoPlus,
	err error) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"GetFileInfoPlus()",
		"")

	if err != nil {

		return fileInfoPlus, err
	}

	fileInfoPlus,
		err = new(fileHelperNanobot).
		getFileInfoPlus(
			pathFileName,
			ePrefix)

	return fileInfoPlus, err
}

// GetFileLastModificationDate
//
// Returns the last modification' date/time on a specific
// file. If input parameter 'customTimeFmt' string is
// empty, a default time format will be used to format
// the returned time string.
//
// The default date time format is:
//
//	"2006-01-02 15:04:05.000000000"
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathFileName				string
//
//		This string contains the path and file name of
//		the file which the last file modification time
//		will be returned.
//
//	customTimeFmt				string
//
//		This string holds the date/time format which will
//		be used to format the last file modification time.
//		Reference:
//			https://pkg.go.dev/time#pkg-constants
//			https://yourbasic.org/golang/format-parse-string-time-date-example
//
//		If this parameter is configured as an empty or
//		zero length string, a default date/time format
//		will be applied. The default date time format is:
//
//			"2006-01-02 15:04:05.000000000"
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
//	time.Time
//
//		If this method completes successfully, the last
//		file modification date/time associated with the
//		file identified by input parameter 'pathFileName'
//		will be returned by this parameter.
//
//	string
//
//		If this method completes successfully, the last
//		file modification date/time associated with the
//		file identified by input parameter 'pathFileName'
//		will be formatted as a string and returned by this
//		parameter.
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
func (fh *FileHelper) GetFileLastModificationDate(
	pathFileName string,
	customTimeFmt string,
	errorPrefix interface{}) (
	time.Time,
	string,
	error) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"GetFileLastModificationDate()",
		"")

	if err != nil {
		return time.Time{}, "", err
	}

	return new(fileHelperNanobot).
		getFileLastModificationDate(
			pathFileName,
			customTimeFmt,
			ePrefix)
}

// GetFileMode
//
// Returns the mode (os.FileMode) for the file designated
// by input parameter 'pathFileName'. If the file does
// not exist, an error will be returned.
//
// The 'os.FileMode' is returned via type
// 'FilePermissionCfg' which includes methods necessary to
// interpret the 'os.FileMode'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathFileName				string
//
//		The file mode for the file identified by
//		'pathFileName' will be returned as an
//		instance of FilePermissionConfig. Type
//		FilePermissionConfig contains file
//		permission information.
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
//		This parameter contains the file permissions
//		currently applied to the file identified by
//		input paramter 'pathFileName'.
//
//		Type FilePermissionConfig contains methods
//		providing information and management options
//		for the encapsulated file permissions
//		associated with 'pathFileName'.
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
func (fh *FileHelper) GetFileMode(
	pathFileName string,
	errorPrefix interface{}) (
	FilePermissionConfig,
	error) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"GetFileMode()",
		"")

	if err != nil {
		return FilePermissionConfig{}, err
	}

	return new(fileHelperNanobot).getFileMode(
		pathFileName,
		ePrefix)
}

// GetFileNameWithExt
//
// This method expects to receive a valid directory path
// and file name or file name plus extension. It then
// extracts the File Name and Extension from the file
// path and returns it as a string.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathFileNameExt string
//
//		This input parameter is expected to contain a
//		properly formatted directory path and File Name.
//		The File Name may or may not include a File
//		Extension.
//
//		The directory path must include the correct
//		delimiters such as path Separators ('/' or'\'),
//		dots ('.') and in the case of Windows, a volume
//		designation (Example: 'F:').
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
//	fNameExt					string
//
//		If successful, this method will the return value
//		'fNameExt' equal to the File Name and File
//		Extension extracted from the input file path,
//		'pathFileNameExt'.
//
//		Example-1:	'fNameExt' return value:
//					'someFilename.txt'
//
//			If the File Extension is not present, only
//			the File Name will be returned.
//
//		Example-2:	Return value with no file extension:
//					'someFilename'.
//
//	isEmpty						bool
//
//		If this method CAN NOT parse out a valid File
//		Name and Extension from input parameter string,
//		'pathFileNameExt', return value 'fNameExt' will
//		be set to an empty string and return value
//		'isEmpty' will be set to 'true'.
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
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fh *FileHelper) GetFileNameWithExt(
	pathFileNameExt string,
	errorPrefix interface{}) (
	fNameExt string,
	isEmpty bool,
	err error) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	fNameExt = ""
	isEmpty = true

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"GetFileNameWithExt()",
		"")

	if err != nil {
		return fNameExt, isEmpty, err
	}

	fNameExt,
		isEmpty,
		err = new(fileHelperNanobot).
		getFileNameWithExt(
			pathFileNameExt,
			ePrefix)

	return fNameExt, isEmpty, err
}

// GetFileNameWithoutExt
//
// Returns the file name without the path or extension.
//
// If the returned File Name is an empty string, return
// parameter 'isEmpty' is set to 'true'.
//
// ----------------------------------------------------------------
//
// # Usage Examples
//
//	     Actual Path Plus File Name: = "./pathfilego/003_filehelper/common/xt_dirmgr_01_test.go"
//	             Returned File Name: = "dirmgr_01_test"
//
//	Actual File Name Plus Extension: "newerFileForTest_01.txt"
//	             Returned File Name: "newerFileForTest_01"
//
//	Actual File Name Plus Extension: "newerFileForTest_01"
//	             Returned File Name: "newerFileForTest_01"
//
//	Actual File Name Plus Extension: ".gitignore"
//	             Returned File Name: ".gitignore"
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathFileNameExt				string
//
//		This string holds the path, file name and file
//		extension. This method will extract the file name
//		from this string and return it to the calling
//		function.
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
//	fName						string
//
//			This return parameter contains the file name
//			extracted from input parameter
//			'pathFileNameExt'.
//
//	isEmpty						bool
//
//		If the returned File Name is an empty string,
//		'isEmpty' is set to 'true'.
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
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fh *FileHelper) GetFileNameWithoutExt(
	pathFileNameExt string,
	errorPrefix interface{}) (
	fName string,
	isEmpty bool,
	err error) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	fName = ""
	isEmpty = true

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"GetFileNameWithoutExt()",
		"")

	if err != nil {
		return fName, isEmpty, err
	}

	fName,
		isEmpty,
		err = new(fileHelperMechanics).
		getFileNameWithoutExt(
			pathFileNameExt,
			ePrefix)

	return fName, isEmpty, err
}

// GetFirstLastNonSeparatorCharIndexInPathStr
//
// Basically, this method returns the first index of the
// first alphanumeric character in a path string
// (reading from left to right).
//
// Specifically, the character must not be a path
// Separator ('\', '/') and it must not be a dot ('.').
//
// If the first Non-Separator char is found, this method
// will return an integer index which is greater than or
// equal to zero plus an error value of nil.
//
// The first character found will never be part of the
// volume name.
//
// Example On Windows:
//
//	"D:\fDir1\fDir2" - first character index will
//						be 3 denoting character 'f'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathStr						string
//
//		A string specifying a file or directory path.
//		Basically, this method returns the first index
//		of the first alphanumeric character in this path
//		string. See the example above.
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
//		firstIdx					int
//
//			The string index in input parameter 'pathStr'
//			identifying the first alphanumeric character.
//
//			The character identified by this index will not
//			be a path Separator ('\', '/') and or a dot ('.').
//
//			If this returned value is less than zero, it signals
//			that no valid first character was found.
//
//			The first character found will never be part of the
//			volume name.
//
//			Example On Windows:
//
//			"D:\fDir1\fDir2" - first character index will
//	                         be 3 denoting character 'f'.
//
//		lastIdx						int
//
//			The string index in input parameter 'pathStr'
//			identifying the last alphanumeric character.
//
//			The character identified by this index will not
//			be a path Separator ('\', '/') and or a dot ('.').
//
//			The last character found will never be part of the
//			volume name.
//
//		error
//
//			If this method completes successfully, the
//			returned error Type is set equal to 'nil'.
//
//			If errors are encountered during processing, the
//			returned error Type will encapsulate an error
//			message. This returned error message will
//			incorporate the method chain and text passed by
//			input parameter, 'errorPrefix'. The 'errorPrefix'
//			text will be attached to the beginning of the
//			error message.
func (fh *FileHelper) GetFirstLastNonSeparatorCharIndexInPathStr(
	pathStr string,
	errorPrefix interface{}) (
	firstIdx,
	lastIdx int,
	err error) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	firstIdx = -1

	lastIdx = -1

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"GetFirstLastNonSeparatorCharIndexInPathStr()",
		"")

	if err != nil {
		return firstIdx, lastIdx, err
	}

	firstIdx,
		lastIdx,
		err = new(fileHelperMolecule).
		getFirstLastNonSeparatorCharIndexInPathStr(
			pathStr,
			ePrefix)

	return firstIdx, lastIdx, err
}

// GetLastPathElement
//
// Analyzes a 'pathName' string and returns the last
// element in the path. If 'pathName' ends in a path
// separator ('/'), this method returns an empty
// string.
//
// ----------------------------------------------------------------
//
// # Usage
//
//	pathName = '../dir1/dir2/fileName.ext' will return "fileName.ext"
//	pathName = '../dir1/dir2/' will return ""
//	pathName = 'fileName.ext' will return "fileName.ext"
//	pathName = '../dir1/dir2/dir3' will return "dir3"
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
//		If this method completes successfully, this
//		string will return the last path element found in
//		input parameter 'pathName'.
//
//		If 'pathName' ends in a path separator ('/'),
//		this string parameter will be returned as an
//		empty or zero length string.
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
func (fh *FileHelper) GetLastPathElement(
	pathName string,
	errorPrefix interface{}) (
	string,
	error) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"GetLastPathElement()",
		"")

	if err != nil {
		return "", err
	}

	return new(fileHelperMolecule).
		getLastPathElement(
			pathName,
			ePrefix)
}

// GetPathAndFileNameExt
//
// Breaks out path and fileName+Ext elements from a path
// string. If both path and fileName are empty strings,
// this method returns an error.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathFileNameExt				string
//
//		This string holds the file path, file name and
//		file extension. The file path will be returned
//		as in the first parameter. The file name and file
//		extension will be returned in a second paramter.
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
//	pathDir						string
//
//		This returned string will contain the directory
//		path extracted from input parameter 'pathFileNameExt'.
//
//	fileNameExt					string
//
//		This returned string will contain the file name
//		and file extension extracted from input parameter
//		'pathFileNameExt'.
//
//	bothAreEmpty				bool
//
//		If both 'pathDir' and 'fileNameExt' are returned
//		as empty strings, this return parameter will be
//		set to 'true'.
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
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fh *FileHelper) GetPathAndFileNameExt(
	pathFileNameExt string,
	errorPrefix interface{}) (
	pathDir string,
	fileNameExt string,
	bothAreEmpty bool,
	err error) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	pathDir = ""

	fileNameExt = ""

	bothAreEmpty = true

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"GetPathAndFileNameExt()",
		"")

	if err != nil {
		return pathDir, fileNameExt, bothAreEmpty, err
	}

	pathDir,
		fileNameExt,
		bothAreEmpty,
		err = new(fileHelperDirector).
		getPathAndFileNameExt(
			pathFileNameExt,
			ePrefix)

	return pathDir, fileNameExt, bothAreEmpty, err
}

// GetPathFromPathFileName
//
// Returns the path from a path and file name string. If
// the returned path is an empty string, return parameter
// 'isEmpty' is set to 'true'.
//
// ----------------------------------------------------------------
//
// # Usage Examples
//
//	pathFileNameExt = ""                  returns isEmpty==true  err==nil
//	pathFileNameExt = "D:\"               returns "D:\"
//	pathFileNameExt = "."                 returns ".\"
//	pathFileNameExt = "..\"               returns "..\"
//	pathFileNameExt = "...\"              returns ERROR
//
//	pathFileNameExt = ".\pathfile\003_filehelper\wt_HowToRunTests.md"
//	                                      returns ".\pathfile\003_filehelper"
//
//	pathFileNameExt = "someFile.go"       returns ""
//	pathFileNameExt = "..\dir1\dir2\.git" returns "..\dir1\dir2"
//	                                       '.git' is assumed to be a file.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathFileNameExt  string
//
//		This is an input parameter. The method expects to
//		receive a single, properly formatted path and
//		file name string delimited by dots ('.') and path
//		Separators ('/' or '\').
//
//		On Windows, if the 'pathFileNameExt' string
//		contains valid volume designations
//		(Example: "D:"), these are returned as part of
//		the path.
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
//	path						string
//
//		This is the directory path extracted from the
//		input parameter 'pathFileNameExt'. If successful,
//		the 'path' string that is returned by this method
//		WILL NOT include a trailing path separator
//		('/' or '\' depending on the os).
//
//		Example Return 'path': "./pathfile/003_filehelper"
//
//	isEmpty						bool
//
//		If the method determines that it cannot extract a
//		valid directory path from input parameter
//		'pathFileNameExt', this boolean value will be set
//		to 'true'. Failure to extract a valid directory
//		path will occur if the input parameter
//		'pathFileNameExt' is not properly formatted as a
//		valid path and file name.
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
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fh *FileHelper) GetPathFromPathFileName(
	pathFileNameExt string,
	errorPrefix interface{}) (
	dirPath string,
	isEmpty bool,
	err error) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	dirPath = ""
	isEmpty = true

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"GetPathFromPathFileName()",
		"")

	if err != nil {
		return dirPath, isEmpty, err
	}

	dirPath,
		isEmpty,
		err = new(fileHelperMechanics).
		getPathFromPathFileName(
			pathFileNameExt,
			ePrefix)

	return dirPath, isEmpty, err
}

// GetPathSeparatorIndexesInPathStr
//
// Returns an array containing the indexes of path Separators
// (Forward slashes or backward slashes depending on operating
// system).
//
// The returned integer array identifies the indexes of the
// forward or backward slashes within input paramter string,
// 'pathStr'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathStr						string
//
//		'pathStr' is a string specifying a directory or
//		file path. Depending on the operating system, the
//		directories will be delimited by forward or
//		backward slashes.
//
//		'pathStr' will be examined and the indexes of the
//		forward or backward slashes will be recorded and
//		returned in an integer array.
//
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
//	[]int
//
//		This array contains the string indexes of the
//		forward or backward slashes contained in input
//		paramter string, 'pathStr'.
//
//		'pathStr' is a string specifying a directory or
//		file path. Depending on the operating system the
//		directories will be delimited by forward or
//		backward slashes.
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
func (fh *FileHelper) GetPathSeparatorIndexesInPathStr(
	pathStr string,
	errorPrefix interface{}) (
	[]int,
	error) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"GetPathSeparatorIndexesInPathStr()",
		"")

	if err != nil {
		return []int{}, err
	}

	return new(fileHelperAtom).
		getPathSeparatorIndexesInPathStr(
			pathStr,
			ePrefix.XCpy("<-pathStr"))
}

// GetVolumeName
//
// Returns the volume name of associated with a given
// directory path. The method calls the function
// 'path/filepath.VolumeName().
//
// VolumeName() returns the leading volume name if it
// exists in input parameter 'pathStr'.
//
// ----------------------------------------------------------------
//
// # Usage
//
// Given "C:\foo\bar" it returns "C:" on Windows.
// Given "c:\foo\bar" it returns "c:" on Windows.
// Given "\\host\share\foo" it returns "\\host\share" on linux
// On other platforms, it returns "".
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://pkg.go.dev/path/filepath#VolumeName
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathStr						string
//
//		This string holds the directory path. The Volume
//		Name will be extracted from the directory path
//		and returned to the calling function.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	string
//
//		This returned string holds the Volume Name
//		extracted from input parameter 'pathStr'.
func (fh *FileHelper) GetVolumeName(
	pathStr string) string {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	return new(fileHelperAtom).
		getVolumeName(
			pathStr)
}

// GetVolumeNameIndex
//
// Analyzes input parameter 'pathStr' to determine if it
// contains a volume name.
//
// This method calls the function
// path/filepath.VolumeName().
//
// VolumeName() returns the leading volume name if it
// exists in input parameter 'pathStr'.
//
// ----------------------------------------------------------------
//
// # Usage Examples
//
// Given "C:\foo\bar" it returns "C:" on Windows.
// Given "c:\foo\bar" it returns "c:" on Windows.
// Given "\\host\share\foo" it returns "\\host\share" on linux
// On other platforms, it returns "".
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathStr						string
//
//		This strings contains the file path which will be
//		searched for the existence of a volume name.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	volNameIndex				int
//
//		The zero based string index of the first
//		character of the Volume Name located in input
//		parameter 'pathStr'.
//
//	volNameLength				int
//
//		This integer value contains the string length of
//		the Volume Name located in input parameter
//		'pathStr'.
//
//	volNameStr				string
//
//		This string contains the Volume Name extracted
//		from input parameter 'pathStr'.
func (fh *FileHelper) GetVolumeNameIndex(
	pathStr string) (
	volNameIndex int,
	volNameLength int,
	volNameStr string) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	volNameIndex,
		volNameLength,
		volNameStr = new(fileHelperAtom).
		getVolumeNameIndex(
			pathStr)

	return volNameIndex, volNameLength, volNameStr
}

// IsAbsolutePath
//
// Compares the input parameter 'pathStr' to the absolute
// path representation for 'pathStr' to determine whether
// 'pathStr' represents an absolute path.
//
// This method differs from isAbsolutePathByCompare() in
// that this method calls low level method
// filePath.IsAbs() to determine if a path is an absolute
// path.
//
// ----------------------------------------------------------------
//
// Absolute Path Definition (Wikipedia):
//
//	An absolute or full path points to the same location
//	in a file system, regardless of the current working
//	directory. To do that, it must include the root
//	directory.
//
//	By contrast, a relative path starts from some given
//	working directory, avoiding the need to provide the
//	full absolute path. A filename can be considered as
//	a relative path based at the current working
//	directory. If the working directory is not the file's
//	parent directory, a file not found error will result
//	if the file is addressed by its name.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathStr						string
//
//		This string holds the file path which will be
//		analyzed to determine if it is an absolute path.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		This method will analyze input parameter
//		'pathStr' to determine if it is an absolute path.
//
//		This determination is made by calling low level
//		function filePath.IsAbs().
//
//		If input parameter 'pathStr' is determined to be
//		an absolute path, this returned boolean value
//		will be set to true.
func (fh *FileHelper) IsAbsolutePath(
	pathStr string) bool {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	return new(fileHelperMolecule).isAbsolutePath(
		pathStr)
}

// IsAbsolutePathByCompare
//
// Compares the input parameter 'pathStr' to the absolute
// path representation for 'pathStr' to determine whether
// 'pathStr' represents an absolute path.
//
// This method differs from isAbsolutePath() in that
// this method does NOT call filePath.IsAbs(). Instead,
// this method constructs an absolute path from 'pathStr'
// and compares the two paths.
//
// ----------------------------------------------------------------
//
// Absolute Path Definition (Wikipedia):
//
//	An absolute or full path points to the same location
//	in a file system, regardless of the current working
//	directory. To do that, it must include the root
//	directory.
//
//	By contrast, a relative path starts from some given
//	working directory, avoiding the need to provide the
//	full absolute path. A filename can be considered as
//	a relative path based at the current working
//	directory. If the working directory is not the file's
//	parent directory, a file not found error will result
//	if the file is addressed by its name.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathStr						string
//
//		This string holds the file path which will be
//		analyzed to determine if it is an absolute path.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		This method will analyze input parameter
//		'pathStr' to determine if it is an absolute path.
//
//		This determination is made by comparing 'pathStr'
//		to the absolute path constructed with 'pathStr'.
//
//		If input parameter '' is determined to be an
//		absolute path, this returned boolean value will be
//		set to true.
func (fh *FileHelper) IsAbsolutePathByCompare(
	pathStr string) bool {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	return new(fileHelperMolecule).
		isAbsolutePathByCompare(pathStr)
}

// IsPathFileString
//
// Returns 'true' if it is determined that input
// parameter, 'pathFileStr', represents a directory path
// plus a file name or file name and file extension.
//
// If 'pathFileStr' is judged to be a directory path and
// file name, by definition it cannot be solely a
// directory path.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathFileStr					string
//
//		This string will be analyzed to determine if it
//		contains both a directory path and a file name or
//		file name and file extension.
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
//	pathFileType PathFileTypeCode
//
//		Path File Type Code indicating whether the input
//		parameter 'pathFileStr' is a Path, a Path and
//		File, a File or "Indeterminate". "Indeterminate"
//		signals that the nature of 'pathFileStr' cannot
//		be classified as either a Path or a Path and File
//		or a File.
//
//		--------------------------------------------------------
//		PathFileTypeCodes
//		   0 = None
//		   1 = Path
//		   2 = PathFile
//		   3 = File (with no path)
//		   4 = Volume
//		   5 = Indeterminate
//		   		Cannot determine whether string is a
//		   		Path, Path & File or File
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
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fh *FileHelper) IsPathFileString(
	pathFileStr string,
	errorPrefix interface{}) (
	pathFileType PathFileTypeCode,
	absolutePathFile string,
	err error) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	absolutePathFile = ""

	pathFileType = PathFileType.None()

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"IsPathFileString()",
		"")

	if err != nil {
		return pathFileType, absolutePathFile, err
	}

	pathFileType,
		absolutePathFile,
		err = new(fileHelperNanobot).isPathFileString(
		pathFileStr,
		ePrefix)

	return pathFileType, absolutePathFile, err
}

// IsPathString
//
// Attempts to determine whether a string is a path
// string designating a directory (and not a path file
// name file extension string).
//
// If the path exists on disk, this method will examine
// the associated file information and determine whether
// the path string represents a directory.
//
// ------------------------------------------------------------------------
//
// Input Parameter:
//
//	pathStr						string
//
//		The path string to be analyzed. This will method
//		will determine whether 'pathStr' is a directory
//		path.
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
// ------------------------------------------------------------------------
//
// Return Values:
//
//	isPathStr       			bool
//
//		If the input parameter, 'pathStr' is determined
//		to be a directory path, this return value is set
//		to 'true'. Here, a 'directory path' is defined as
//		a true directory and the path does NOT contain a
//		file name.
//
//	cannotDetermine bool
//
//		If the method cannot determine whether the input
//		parameter 'pathStr' is a valid directory path,
//		this return value will be set to 'true'.
//
//		The 'cannotDetermine=true' condition occurs with
//		path names like 'D:\DirA\common'. The method
//		cannot determine whether 'common' is a file name
//		or a directory name.
//
//
//	testPathStr					string
//
//		Input parameter 'pathStr' is subjected to
//		cleaning routines designed to exclude extraneous
//		characters from the analysis. testPathFileStr'
//		is the actual string on which the analysis was
//		performed.
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
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fh *FileHelper) IsPathString(
	pathStr string,
	errorPrefix interface{}) (
	isPathStr bool,
	cannotDetermine bool,
	testPathStr string,
	err error) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"IsPathString()",
		"")

	if err != nil {
		return isPathStr, cannotDetermine, testPathStr, err
	}

	isPathStr,
		cannotDetermine,
		testPathStr,
		err = new(fileHelperMechanics).
		isPathString(
			pathStr,
			ePrefix)

	return isPathStr, cannotDetermine, testPathStr, err
}

// IsStringEmptyOrBlank
//
// Analyzes a string to determine if the string is
// 'empty' or if the string consists of all blanks
// (spaces).
//
// In addition, if the string contains characters,
// all leading and trailing spaces are 'trimmed' or
// deleted.
//
// If the string consists entirely of blank spaces or
// white space (" "), it will be reset to an empty
// string, ("").
//
// ----------------------------------------------------------------
//
// # Usage
//
//	testStr string
//
//	The original input parameter string ('testStr') from
//	which the leading and trailing spaces will be deleted.
//
//	         Examples:
//
//	            1. 'testStr'     = "  a string   "
//	               return string = "a string"
//
//	            2. 'testStr'     = "    "
//	               return string = ""
//
//	            3. 'testStr'     = ""
//	               return string = ""
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	testStr						string
//
//		This string will be analyzed to determine if
//		is empty (zero length) or consists entirely of
//		blank or white spaces (" ").
//
//		This string will be returned as return parameter
//		'newStr'.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	errCode			int
//
//		This method will analyze input parameter
//		'testStr' and return one of the following status
//		codes
//
//		Integer Codes Returned:
//			-1 = string is empty
//			-2 = string consists entirely of spaces
//			 0 = string contains non-white space
//			 	 characters.
//
//	strLen			int
//
//		The length of 'newStr' returned by this method.
//
//	newStr			string
//
//		If 'testStr' consists of non-white space
//		characters, this method will delete leading
//		and trailing white spaces.
//
//		If 'testStr' consists entirely of white space
//		characters, 'newStr' will be returned as an empty
//		string ("").
func (fh *FileHelper) IsStringEmptyOrBlank(
	testStr string) (
	errCode int,
	strLen int,
	newStr string) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	return new(fileHelperElectron).
		isStringEmptyOrBlank(
			testStr)
}

// JoinPathsAdjustSeparators
//
// Joins two path strings and standardizes the path
// separators according to the current operating system.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	p1							string
//
//		This string holds one of two path strings which
//		will be joined together in tandem (p1+p2). 'p1'
//		will be located at the beginning of the
//		composite, joined path string returned to the
//		calling function.
//
//	p2							string
//
//		This string holds the second of two path strings
//		which will be joined together in tandem (p1+p2).
//		'p2' will be located at the end of the composite,
//		joined path string returned to the calling
//		function.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	string
//
//		This returned path string contains the two joined
//		path strings provided by input parameter 'p1' and
//		'p2' (p1+p2).
func (fh *FileHelper) JoinPathsAdjustSeparators(
	p1 string, p2 string) string {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	return new(fileHelperMolecule).
		joinPathsAdjustSeparators(p1, p2)
}

// MakeAbsolutePath
//
// Receives a relative path or any path string and
// resolves that path to an Absolute path. This method
// calls the low lever method filepath.Abs() to
// generate the absolute path.
//
// "An absolute or full path points to the same location
// in a file system, regardless of the current working
// directory. To do that, it must include the root
// directory.
//
// By contrast, a relative path starts from some given
// working directory, avoiding the need to provide the
// full absolute path. A filename can be considered as a
// relative path based at the current working directory.
// If the working directory is not the file's parent
// directory, a file not found error will result if the
// file is addressed by its name."
//
//	Wikipedia
//
// Note:
//
// Clean() is called on the result generated by
// filepath.Abs().
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://en.wikipedia.org/wiki/Path_(computing)#Absolute_and_relative_paths
//	https://pkg.go.dev/path/filepath@go1.20.1#Abs
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	relPath						string
//
//		This string holds a relative path. This method
//		will convert this relative path to an absolute
//		path.
//
//		A relative path is defined as follows:
//
//		"A relative path starts from some given working
//		directory, avoiding the need to provide the full
//		absolute path. A filename can be considered as a
//		relative path based at the current working
//		directory. If the working directory is not the
//		file's parent directory, a file not found error
//		will result if the file is addressed by its name."
//			Wikipedia
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
//		If this method completes successfully, this
//		method will convert the relative path received
//		from input parameter '', to an absolute path.
//
//		"An absolute or full path points to the same
//		location in a file system, regardless of the
//		current working directory. To do that, it must
//		include the root directory."	Wikipedia
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
func (fh *FileHelper) MakeAbsolutePath(
	relPath string,
	errorPrefix interface{}) (
	string,
	error) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"MakeAbsolutePath()",
		"")

	if err != nil {
		return "", err
	}

	return new(fileHelperProton).makeAbsolutePath(
		relPath,
		ePrefix.XCpy("<-relPath"))
}

// MakeDirAll
//
// Creates a directory named path, along with any
// necessary parent directories. In other words, all
// directories in the path are created.
//
// The permission bits 'drwxrwxrwx' are used for all
// directories that the method creates.
//
// If path is a directory which already exists, this
// method does nothing and returns and error value of
// 'nil'.
//
// Note:
//
// This method calls MakeDirAllPerm()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dirPath						string
//
//		This string contains the name of the directory
//		path which will be created by this method.
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
func (fh *FileHelper) MakeDirAll(
	dirPath string,
	errorPrefix interface{}) error {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"MakeDirAll()",
		"")

	if err != nil {
		return err
	}

	return new(fileHelperMechanics).makeDirAll(
		dirPath,
		ePrefix)
}

// MakeDir
//
// Creates a single directory. The method returns an
// error type. If the operation succeeds, the error value
// is 'nil'. If the operation fails the error value is
// populated with an appropriate error message.
//
// This method will fail if the parent directory does not
// exist.
//
// The permission bits 'drwxrwxrwx' are used for
// directory creation. If path is already a directory,
// this method does nothing and returns an error value of
// 'nil'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dirPath						string
//
//		This string holds the new directory path which
//		will be created by this method.
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
//		A returned value of 'nil' signals that the new
//		directory path was successfully created.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fh *FileHelper) MakeDir(
	dirPath string,
	errorPrefix interface{}) error {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"MakeDir()",
		"")

	if err != nil {
		return err
	}

	var permission FilePermissionConfig

	permission,
		err = new(FilePermissionConfig).
		New("drwxrwxrwx",
			ePrefix)

	if err != nil {

		return err
	}

	return new(fileHelperAtom).makeDirPerm(
		dirPath,
		permission,
		ePrefix.XCpy(
			fmt.Sprintf("dirPath= %v",
				dirPath)))
}

// MakeDirAllPerm
//
// Creates a directory path along with any necessary
// parent paths and configures permissions for the
// new directory path.
//
// If the target directory path already exists, this
// method does nothing and returns.
//
// The input parameter 'permission' is of type
// 'FilePermissionConfig'. See method the documentation
// for method 'FilePermissionConfig.New()' for an
// explanation of permission codes.
//
// If you wish to grant total access to a directory,
// consider setting permission code as follows:
//
//	FilePermissionConfig{}.New("drwxrwxrwx")
//
// If the parent directories in parameter 'dirPath' do
// not yet exist, this method will create them.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dirPath						string
//
//		This string contains the directory path which
//		will be created by this method.
//
//	permission					FilePermissionConfig
//
//		An instance of FilePermissionConfig containing
//		the permission specifications for the new
//		directory to be created from input paramter,
//		'dirPath'.
//
//		The easiest way to configure permissions is
//		to call FilePermissionConfig.New() with
//		a mode string ('modeStr').
//
//		The first character of the 'modeStr' designates the
//		'Entry Type'. Currently, only two 'Entry Type'
//		characters are supported. Therefore, the first
//		character in the 10-character input parameter
//		'modeStr' MUST be either a "-" indicating a file, or
//		a "d" indicating a directory.
//
//		The remaining nine characters in the 'modeStr'
//		represent unix permission bits and consist of three
//		group fields each containing 3-characters. Each
//		character in the three group fields may consist of
//		'r' (Read-Permission), 'w' (Write-Permission), 'x'
//		(Execute-Permission) or '-' signaling no permission or
//		no access allowed. A typical 'modeStr' authorizing
//		permission for full access to a file would be styled
//		as:
//
//		Directory Example: "drwxrwxrwx"
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
//	    -----------------------------------------------------
//	           Directory Mode String Permission Codes
//	    -----------------------------------------------------
//	      Directory
//			10-Character
//			 'modeStr'
//			 Symbolic		  Directory Access
//			  Format	   Permission Descriptions
//			----------------------------------------------------
//
//			d---------		no permissions
//			drwx------		read, write, & execute only for owner
//			drwxrwx---		read, write, & execute for owner and group
//			drwxrwxrwx		read, write, & execute for owner, group and others
//			d--x--x--x		execute
//			d-w--w--w-		write
//			d-wx-wx-wx		write & execute
//			dr--r--r--		read
//			dr-xr-xr-x		read & execute
//			drw-rw-rw-		read & write
//			drwxr-----		Owner can read, write, & execute. Group can only read;
//			                others have no permissions
//
//			Note: drwxrwxrwx - identifies permissions for directory
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
func (fh *FileHelper) MakeDirAllPerm(
	dirPath string,
	permission FilePermissionConfig,
	errorPrefix interface{}) error {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"MakeDirAllPerm()",
		"")

	if err != nil {
		return err
	}

	return new(fileHelperNanobot).
		makeDirAllPerm(
			dirPath,
			permission,
			ePrefix)
}

// MakeDirPerm
//
// Creates a single directory using the permission codes
// passed by input parameter 'permission'.
//
// This method will fail if the parent directory does not
// exist. To create all parent directories in the path use
// method 'FileHelper.MakeDirAllPerm()'.
//
// The input parameter 'permission' is of type
// 'FilePermissionConfig'. See the source code
// documentation for method 'FilePermissionConfig.New()'
// for an explanation of permission codes.
//
// If you wish to grant total access to a directory,
// consider setting permission code as follows:
//
//	FilePermissionConfig{}.New("drwxrwxrwx")
//
// An error will be triggered if the 'dirPath' input
// parameter represents an invalid path or if parent
// directories in the path do not exist.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dirPath						string
//
//		The directory path to be created,
//
//	permission					FilePermissionConfig
//
//		The instance of FilePermissionConfig contains the
//		permissions which will be applied to the newly
//		created directory ('dirPath').
//
//		For an explanation of permission bit settings,
//		see the source code documentation for:
//
//			FilePermissionConfig{}.New()
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
func (fh *FileHelper) MakeDirPerm(
	dirPath string,
	permission FilePermissionConfig,
	errorPrefix interface{}) error {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"MakeDirPerm()",
		"")

	if err != nil {
		return err
	}

	return new(fileHelperAtom).
		makeDirPerm(
			dirPath,
			permission,
			ePrefix)
}

// MakeFileHelperWalkDirDeleteFilesFunc
//
// Used in conjunction with DirMgr.DeleteWalDirFiles to
// select and delete files residing the directory tree
// identified by the current DirMgr object.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dInfo				*DirectoryDeleteFileInfo
//
//		A pointer to an instance of DirectoryDeleteFileInfo.
//
//		DirectoryDeleteFileInfo
//		This structure is used to delete files in a
//		directory specified	by 'StartPath'. Deleted files
//		will be selected based on 'DeleteFileSelectCriteria'
//		value.
//
//		'DeleteFileSelectCriteria' is a 'FileSelectionCriteria'
//		type which contains FileNamePatterns strings and the
//		FilesOlderThan or FilesNewerThan date time parameters
//		which can be used as file selection criteria.
//
//		type DirectoryDeleteFileInfo struct {
//			StartPath                string
//			Directories              DirMgrCollection
//			ErrReturns               []error
//			DeleteFileSelectCriteria FileSelectionCriteria
//			DeletedFiles             FileMgrCollection
//		}
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	This method returns a function with the following
//	signature:
//
//		func(string, os.FileInfo, error) error
func (fh *FileHelper) MakeFileHelperWalkDirDeleteFilesFunc(dInfo *DirectoryDeleteFileInfo) func(string, os.FileInfo, error) error {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	return new(fileHelperMolecule).
		makeFileHelperWalkDirDeleteFilesFunc(dInfo)
}

// MakeFileHelperWalkDirFindFilesFunc
//
// This function is designed to work in conjunction with
// a walk directory function like FindWalkDirFiles. It
// will process files extracted from a 'Directory Walk'
// operation initiated by the 'filepath.Walk' method.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dInfo						*DirectoryTreeInfo
//
//	A pointer to a DirectoryTreeInfo instance. This
//	structure is used to 'Find' files in a directory
//	specified by member variable, 'StartPath'. The file
//	search will be filtered by using member variable
//	'FileSelectCriteria' selection criteria
//	specifications.
//
//	'FileSelectCriteria' is a FileSelectionCriteria type
//	which contains FileNamePatterns strings and
//	'FilesOlderThan' or 'FilesNewerThan' date time
//	parameters which can be used as a selection
//	criteria.
//
//		type DirectoryTreeInfo struct {
//			StartPath          string
//			Directories        DirMgrCollection
//			FoundFiles         FileMgrCollection
//			ErrReturns         []error
//			FileSelectCriteria FileSelectionCriteria
//		}
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	This method returns a function with the following
//	signature:
//
//		func(string, os.FileInfo, error) error
func (fh *FileHelper) MakeFileHelperWalkDirFindFilesFunc(
	dInfo *DirectoryTreeInfo) func(string, os.FileInfo, error) error {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	return new(fileHelperMolecule).
		makeFileHelperWalkDirFindFilesFunc(
			dInfo)
}

// MoveFile
//
// Copies a file from source to destination and, if
// successful, then deletes the original source file.
//
// The copy procedure will be carried out using the
// 'Copy By Io' technique. See FileHelper.CopyFileByIo().
//
// The 'move' operation will create the destination file,
// but it will NOT create the destination directory. If
// the destination directory does NOT exist, an error will
// be returned.
//
// If this copy operation fails, the method will return an
// error, and the source file will NOT be deleted.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete the source file after the
//	copy from source to destination file is completed.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	src							string
//
//		This string holds the path and file name of the
//		source file. This source file will be copied to
//		the destination path and file name specified by
//		input parameter 'dst'.
//
//		After the source file ('src') is copied to the
//		destination file ('dst'), the source file ('src')
//		will be deleted.
//
//	dst							string
//
//		This string holds the path and file name of the
//		destination file. The source file ('src') will be
//		copied to this destination file ('dst') and the
//		source file will be deleted.
//
//		If the directory path for the destination file
//		('dst') does not previously exist, an error will
//		be returned, the source file ('src') will not be
//		copied to the destination file ('dst') and the
//		source file will NOT be deleted.
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
func (fh *FileHelper) MoveFile(
	src string,
	dst string,
	errorPrefix interface{}) error {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"MoveFile()",
		"")

	if err != nil {
		return err
	}

	return new(fileHelperDirector).moveFile(
		src,
		dst,
		ePrefix)
}

// OpenDirectory
//
// Opens a directory and returns the associated 'os.File'
// pointer. This method will open a directory designated
// by input parameter, 'directoryPath'.
//
// The input parameter 'createDir' determines the action
// taken if 'directoryPath' does not exist. If
// 'createDir' is set to 'true' and 'directoryPath' does
// not currently exist, this method will attempt to
// create 'directoryPath'. Directories created in this
// manner are configured with Open Type of 'Read-Write'
// and a Permission code of 'drwxrwxrwx'.
//
// Alternatively, if 'createDir' is set to 'false' and
// 'directoryPath' does NOT exist, an error will be
// returned.
//
// Regardless of whether the target directory path
// already exists or is created by this method, the
// returned os.File pointer is opened with the
// 'Read-Only' attribute (O_RDONLY) and a
// permission code of zero ("----------").
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// The caller is responsible for calling "Close()" on the
// returned os.File pointer.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	directoryPath				string
//
//		A string containing the path name of the
//		directory which will be opened.
//
//	createDir					bool
//
//		Determines what action will be taken if
//		'directoryPath' does NOT exist. If 'createDir' is
//		set to 'true' and 'directoryPath' does NOT exist,
//		this method will attempt to create
//		'directoryPath'. Alternatively, if 'createDir' is
//		set to false and 'directoryPath' does NOT exist,
//		this method will terminate and an error will be
//		returned.
//
//		Directories created in this manner will have an
//		Open Type of 'Read-Write' and a Permission code
//		of 'drwxrwxrwx'. This differs from the Open Type
//		and permission mode represented by the returned
//		os.File pointer.
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
//	*os.File
//
//		If successful, this method returns an os.File
//		pointer to the directory designated by input
//		parameter 'directoryPath'.
//
//		If successful, the returned os.File pointer is
//		opened with the 'Read-Only' attribute (O_RDONLY)
//		and a permission code of zero ("----------").
//
//		If this method fails, the *os.File return value
//		is 'nil'.
//
//		Note:
//			The caller is responsible for calling
//			"Close()" on this os.File pointer.
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
func (fh *FileHelper) OpenDirectory(
	directoryPath string,
	createDir bool,
	errorPrefix interface{}) (*os.File, error) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"OpenDirectory()",
		"")

	if err != nil {
		return nil, err
	}

	return new(fileHelperDirector).
		openDirectory(
			directoryPath,
			createDir,
			ePrefix)
}

// OpenFile
//
// This method is a wrapper for os.OpenFile. This method
// may be used to open or create files depending on the
// File Open and File Permission parameters.
//
// If successful, this method will return a pointer to
// the os.File object associated with the file designated
// for opening.
//
// The calling routine is responsible for calling
// "Close()" on this os.File pointer.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	The calling method is responsible for calling
//	"Close()" on the os.File pointer returned by this
//	method.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathFileName				string
//
//		A string containing the path and file name of the
//		file which will be opened. If a parent path
//		component does NOT exist, this method will
//		trigger an error.
//
//	fileOpenCfg					FileOpenConfig
//
//		This parameter encapsulates the File Open
//		parameters which will be used to open subject
//		file. For an explanation of File Open parameters,
//		see the source code documentation for method
//		FileOpenConfig.New().
//
//	filePermissionCfg			FilePermissionConfig
//
//		This parameter encapsulates the File Permission
//		parameters which will be used to open the subject
//		file. For an explanation of File Permission
//	 	parameters, see method FilePermissionConfig.New().
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
// Return Values:
//
//	*os.File
//
//		If successful, this method returns an os.File
//		pointer to the file designated by input parameter
//		'pathFileName'. This file pointer can
//		subsequently be used for reading content from the
//		subject file. It may NOT be used for writing
//		content to the subject file.
//
//		If this method fails, the *os.File return value
//		is 'nil'.
//
//		Note:
//		The caller is responsible for calling "Close()"
//		on this os.File pointer.
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
func (fh *FileHelper) OpenFile(
	pathFileName string,
	fileOpenCfg FileOpenConfig,
	filePermissionCfg FilePermissionConfig,
	errorPrefix interface{}) (
	filePtr *os.File,
	err error) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"OpenFile()",
		"")

	if err != nil {

		return filePtr, err
	}

	return new(fileHelperMechanics).openFile(
		pathFileName,
		fileOpenCfg,
		filePermissionCfg,
		ePrefix)
}

// OpenFileReadOnly
//
// Opens the designated path file name for reading only.
//
// If successful, this method returns a pointer of type
// *os.File which can only be used for reading content
// from the subject file.
//
// This returned file pointer is configured for
// 'Read-Only' operations. You may not write to the
// subject file using this pointer.
//
// If the designated file ('pathFileName') does NOT
// exist, an error will be returned.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	The calling method is responsible for calling
//	"Close()" on the os.File pointer returned by this
//	method.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathFileName        string
//
//		A string containing the path and file name of the
//		file which will be opened in the 'Read-Only'
//		mode. If the path or file does NOT exist, this
//		method will return an error.
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
// ------------------------------------------------------------------------
//
// Return Values:
//
//	*os.File
//
//		If successful, this method  opens the file
//		designated by input parameter 'pathFileName' and
//		returns an os.File pointer to that file. This
//		file pointer can subsequently be used for reading
//		content from the subject file. It may NOT be used
//		for writing content to the subject file.
//
//		If this method fails, the *os.File return value
//		is 'nil'.
//
//		Note:
//
//		The caller is responsible for calling "Close()"
//		on this os.File pointer.
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
func (fh *FileHelper) OpenFileReadOnly(
	pathFileName string,
	errorPrefix interface{}) (
	filePtr *os.File,
	err error) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	filePtr = nil

	funcName := "FileHelper.OpenFileReadOnly()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {
		return filePtr, err
	}

	filePtr,
		err = new(fileHelperNanobot).openFileReadOnly(
		pathFileName,
		ePrefix)

	return filePtr, err
}

// OpenFileReadWrite
//
// Opens the file designated by input parameter
// 'fileName' for 'Writing'. The actual permission code
// used to open the file is 'Read/Write'.
//
// If the method is successful, a pointer to the opened
// file is returned along with an error value of 'nil'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	If the file designated by input parameter
//		'pathFileName' does NOT exist, this method will
//		attempt to create it. However, if the path
//		component of 'pathFileName' does not exist, an
//		error will be returned.
//
//	(2)	If this method completes successfully, the caller
//		is responsible for calling "Close()" on the
//		returned os.File pointer.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathFileName				string
//
//		A string containing the path and file name of the
//		file which will be opened in the 'Read/Write'
//		mode.
//
//		If this file does NOT exist, this method will
//		attempt to create it. However, if the path
//		component of 'pathFileName' does not exist, an
//		error will be returned.
//
//	truncateFile				bool
//
//		If set to 'true' and the target file will be
//		truncated to zero bytes in length before it is
//		opened.
//
//		If set to 'false', the target file will be opened
//		in the 'Append' mode and any bytes written to the
//		file will be appended to the end of the file.
//		Under this scenario, the original file contents
//		are preserved and newly written bytes are added
//		to the end of the file.
//
//		If the file designated by input parameter
//		'pathFileName' does not exist, this parameter
//		('truncateFile') is ignored and the new created
//		file is initialized containing zero bytes.
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
//	*os.File
//
//		If successful, this method returns an os.File
//		pointer to the file opened for 'Read/Write'
//		operations. This file pointer can be used
//		subsequently for writing content to, or reading
//		content from, the subject file.
//
//		If this method fails, this pointer to os.File is
//		set to 'nil'.
//
//		Note:
//
//		The caller is responsible for calling "Close()"
//		on this os.File pointer.
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
func (fh *FileHelper) OpenFileReadWrite(
	pathFileName string,
	truncateFile bool,
	errorPrefix interface{}) (
	*os.File,
	error) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var fPtr *os.File

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper.OpenFileReadWrite()",
		"")

	if err != nil {
		return fPtr, err
	}

	return new(fileHelperNanobot).openFileReadWrite(
		pathFileName,
		truncateFile,
		ePrefix)
}

// OpenFileWriteOnly
//
// Opens a file for 'Write-Only' operations. Input
// parameter 'pathFileName' specifies the path and file
// name of the file which will be opened.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	If the file designated by input parameter
//		'pathFileName' does not exist, this method will
//		attempt to create that file.
//
//	(2)	If the path or directory component of
//		'pathFileName' does not exist, an error will be
//		returned.
//
//	(3)	If the method completes successfully, the caller
//		is responsible for calling 'Close()' on the
//		returned os.File pointer.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathFileName				string
//
//		A string containing the path and file name of the
//		file which will be opened in the 'Write Only'
//		mode.
//
//		If the file does NOT exist, this method will
//		attempt to create it. However, if the path
//		or directory component of 'pathFileName' does not
//		exist, an error will be returned.
//
//	truncateFile				bool
//
//		If set to 'true' and the target file will be
//		truncated to zero bytes in length before it is
//		opened.
//
//		If set to 'false', the target file will be opened
//		in the 'Append' mode and any bytes written to the
//		file will be appended to the end of the file.
//		Under this scenario, the original file contents
//		are preserved and newly written bytes are added
//		to the end of the file.
//
//		If the file designated by input parameter
//		'pathFileName' does not exist, this parameter
//		('truncateFile') is ignored and the newly created
//		file is initialized with zero bytes of content.
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
//	*os.File
//
//		If successful, this method returns an os.File
//		pointer to the file opened for 'Write Only'
//		operations. This file pointer can be used for
//		writing content to the subject file.
//
//		If this method fails, this return parameter will
//		be set to 'nil'.
//
//		Note:
//
//		The caller is responsible for calling "Close()"
//		on this os.File pointer.
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
func (fh *FileHelper) OpenFileWriteOnly(
	pathFileName string,
	truncateFile bool,
	errorPrefix interface{}) (
	*os.File,
	error) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	funcName := "FileHelper.OpenFileWriteOnly() "

	var ePrefix *ePref.ErrPrefixDto

	var fPtr *os.File

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {
		return fPtr, err
	}

	fPtr,
		err = new(fileHelperNanobot).openFileWriteOnly(
		pathFileName,
		truncateFile,
		ePrefix.XCpy(
			"fPtr<-"))

	return fPtr, err
}

// RemovePathSeparatorFromEndOfPathString
//
// This method will remove or delete the Trailing path
// separator from a path string.
//
// If the path string does not have a Trailing path
// separator, no action is taken and the original path
// string is returned to the calling function.
//
// This path separator character will vary depending on
// the operating system ('/' or '\').
//
// ----------------------------------------------------------------
//
// # Usage
//
//	Example-1
//
//	pathStr = "../testdestdir/destdir/"
//	Returned string = "../testdestdir/destdir/"
//
//	Example-2
//
//	pathStr = "D:\logTest\testoverwrite\"
//	Returned string = "D:\logTest\testoverwrite"
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathStr						string
//
//		This string holds the file or directory path from
//		which any trailing path separator will be removed.
//
//		If 'pathStr' consists of all white space, it will
//		be returned by this method as an empty string
//		("").
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	string
//
//		If input parameter 'pathStr' contains a trailing
//		path separator, that trailing path separator will
//		be removed and the remaining characters will be
//		returned by this string parameter.
//
//		If 'pathStr' contains no trailing path separator,
//		the original value of input parameter 'pathStr'
//		will be	returned.
func (fh *FileHelper) RemovePathSeparatorFromEndOfPathString(
	pathStr string) string {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	return new(fileHelperAtom).
		removePathSeparatorFromEndOfPathString(pathStr)
}

// SearchFileModeMatch
//
// This method determines whether the file mode of the
// file described by input parameter, 'fileInfo', is a
// match for the File Selection Criteria
// 'fileSelectCriteria.SelectByFileMode'.
//
// In this context, the term 'file mode' refers to
// os.FileMode. An os.FileMode represents a file's mode
// and permission bits. The bits have the same definition
// on all systems, so that information about files can be
// moved from one system to another portably. For more
// information on 'File Mode', see the Reference Section
// below.
//
// If the file's FileMode matches the
// 'fileSelectCriteria.SelectByFileMode' value, the
// return value, 'isFileModeMatch' is set to 'true'.
//
// If 'fileSelectCriteria.SelectByFileMode' is NOT
// initialized to a valid File Mode value, the return
// value 'isFileModeSet' is set to 'false' signaling that
// the File Mode File Selection Criterion is NOT active
// or engaged.
//
// Note:
//
// Input parameter 'fileInfo' is of type os.FileInfo.  This
// means the user may substitute a type 'FileInfoPlus'
// object for the 'fileInfo' parameter because 'FileInfoPlus'
// implements the 'os.FileInfo' interface.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	os.FileMode https://pkg.go.dev/os#FileMode
//	https://stackoverflow.com/questions/28969455/how-to-properly-instantiate-os-filemode
//	https://www.includehelp.com/golang/os-filemode-constants-with-examples.aspx
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fileInfo					os.FileInfo
//
//		The type os.FileInfo contains data elements
//		describing a file.
//
//		type FileInfo interface {
//
//			Name() string
//				base name of the file
//
//			Size() int64
//				length in bytes for regular files;
//				system-dependent for others
//
//			Mode() FileMode
//				file mode bits
//
//			ModTime() time.Time
//				modification time
//
//			IsDir() bool
//				abbreviation for Mode().IsDir()
//
//			Sys() any
//				underlying data source (can return nil)
//		}
//
//	fileSelectCriteria			FileSelectionCriteria
//
//		The FileSelectionCriteria type allows for
//		configuration of single or multiple file
//		selection criterion. The 'SelectCriterionMode'
//		can be used to specify whether the file must
//		match all, or any one, of the active file
//		selection criterion.
//
//		type FileSelectionCriteria struct {
//		  FileNamePatterns     []string
//			An array of strings containing File Name Patterns
//
//		  FilesOlderThan       time.Time
//			Match files with older modification date times
//
//		  FilesNewerThan       time.Time
//			Match files with newer modification date times
//
//		  SelectByFileMode     FilePermissionConfig
//			Match file mode (os.FileMode).
//
//		  SelectCriterionMode  FileSelectCriterionMode
//			Specifies 'AND' or 'OR' selection mode
//		}
//
//		Elements of the FileSelectionCriteria Type are
//		described below:
//
//		FileNamePatterns []string
//
//			An array of strings which may define one or
//			more search patterns. If a file name matches
//			any one of the search pattern strings, it is
//			deemed to be a 'match' for the search pattern
//			criterion.
//
//			Example Patterns:
//			 FileNamePatterns = []string{"*.log"}
//			 FileNamePatterns = []string{"current*.txt"}
//			 FileNamePatterns = []string{"*.txt", "*.log"}
//
//			If this string array has zero length or if
//			all the strings are empty strings, then this
//			file search criterion is considered 'Inactive'
//			or 'Not Set'.
//
//		FilesOlderThan  time.Time
//
//			This date time type is compared to file
//			modification date times in order to determine
//			whether the file is older than the
//			'FilesOlderThan' file selection criterion. If
//			the file is older than the 'FilesOlderThan'
//			date time, that file is considered a 'match'
//			for this file selection criterion.
//
//			If the value of 'FilesOlderThan' is set to
//			time zero, the default value for type
//	 		time.Time{}, then this
//	 		file selection criterion is considered to be
//	 		'Inactive' or 'Not Set'.
//
//		FilesNewerThan   time.Time
//
//			This date time type is compared to the file
//			modification date time in order to determine
//			whether the file is newer than the
//			'FilesNewerThan' file selection criterion. If
//			the file modification date time is newer than
//			the 'FilesNewerThan' date time, that file is
//			considered a 'match' for this file selection
//			criterion.
//
//			If the value of 'FilesNewerThan' is set to
//			time zero, the default value for type
//			time.Time{}, then this file selection
//			criterion is considered to be 'Inactive' or
//			'Not Set'.
//
//		SelectByFileMode  FilePermissionConfig
//
//			Type FilePermissionConfig encapsulates an
//			os.FileMode. The file selection criterion
//			allows for the selection of files by File Mode.
//
//			File modes are compared to the value of
//			'SelectByFileMode'. If the File Mode for a
//			given file is equal to the value of
//			'SelectByFileMode', that file is considered
//			to be a 'match' for this file selection
//			criterion. Examples for setting
//			SelectByFileMode are shown as follows:
//
//			fsc := FileSelectionCriteria{}
//
//			err = fsc.SelectByFileMode.SetByFileMode(os.FileMode(0666))
//
//			err = fsc.SelectByFileMode.SetFileModeByTextCode("-r--r--r--")
//
//		SelectCriterionMode FileSelectCriterionMode
//
//		This parameter selects the manner in which the
//		file selection criteria above are applied in
//		determining a 'match' for file selection
//		purposes. 'SelectCriterionMode' may be set to
//		one of two constant values:
//
//		(1) FileSelectCriterionMode(0).ANDSelect()
//
//			File selected if all active selection criteria
//			are satisfied.
//
//			If this constant value is specified for the
//			file selection mode, then a given file will
//			not be judged as 'selected' unless all the
//			active selection criterion are satisfied. In
//			other words, if three active search criterion
//			are provided for 'FileNamePatterns',
//			'FilesOlderThan' and 'FilesNewerThan', then a
//			file will NOT be selected unless it has
//			satisfied all three criterion in this example.
//
//		(2) FileSelectCriterionMode(0).ORSelect()
//
//			File selected if any active selection
//			criterion is satisfied.
//
//			If this constant value is specified for the
//			file selection mode, then a given file will
//			be selected if any one of the active file
//			selection criterion is satisfied. In other
//			words, if three active search criterion are
//			provided for 'FileNamePatterns',
//			'FilesOlderThan' and 'FilesNewerThan', then a
//			file will be selected if it satisfies any
//			one of the three criterion in this example.
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
//	isFileModeSet				bool
//
//		This boolean return value signals whether the
//		'SelectByFileMode' file search criteria is set,
//		engaged and active.
//
//		If the 'fileSelectCriteria.SelectByFileMode'
//	 	data element is empty or set to an invalid
//	 	FileMode, it means that the 'SelectByFileMode'
//	 	search criteria is NOT set or engaged. In this
//	 	case, the 'isFileModeSet' parameter will return a
//	 	value of 'false' meaning that the search criteria
//	 	is NOT active or engaged.
//
//		If this parameter is returned with a value of
//		'true', it means that the 'SelectByFileMode'
//		search criterion is set, engaged and active.
//
//	isFileModeMatch				bool
//
//		'fileSelectCriteria.SelectByFileMode' represents
//		a file search pattern designed to identify files
//		which match the specified FileMode. If the file
//	 	specified by input parameter, 'fileInfo',
//		has a FileMode which matches the
//		'fileSelectCriteria.SelectByFileMode'
//		specification, the file search operation is
//		classified as a 'match', and the return value of
//		this parameter, 'isFileNewerThanMatch', is set to
//		'true'.
//
//		If this parameter, 'isFileModeMatch', is
//		returned with a value of 'false', it signals that
//		the file has a FileMode which is NOT equal to
//		FileMode contained in the
//		'fileSelectCriteria.SelectByFileMode'
//		specification.
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
func (fh *FileHelper) SearchFileModeMatch(
	fileInfo os.FileInfo,
	fileSelectCriteria FileSelectionCriteria,
	errorPrefix interface{}) (
	isFileModeSet bool,
	isFileModeMatch bool,
	err error) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"SearchFileModeMatch()",
		"")

	if err != nil {
		return isFileModeSet, isFileModeMatch, err
	}

	isFileModeSet,
		isFileModeMatch,
		err = new(fileHelperElectron).searchFileModeMatch(
		fileInfo,
		fileSelectCriteria,
		ePrefix)

	return isFileModeSet, isFileModeMatch, err
}

// SearchFileNewerThan
//
// This method is called to determine whether the file
// described by the input parameter 'fileInfo' is a
// 'match' for the File Selection Criteria,
// 'fileSelectCriteria.FilesNewerThan'.
//
// If the file modification date time occurs after the
// 'fileSelectCriteria.FilesNewerThan' date time, the
// return value 'isFileNewerThanMatch' is set to 'true'.
//
// If 'fileSelectCriteria.FilesNewerThan' is set to
// time.Time zero ( the default or zero value for this type),
// the return value 'isFileNewerThanSet' is set to 'false'
// signaling that this search criterion is NOT active.
//
// Note:
//
// Input parameter 'fileInfo' is of type os.FileInfo.
// This means the user may substitute a type
// 'FileInfoPlus' object for the 'fileInfo' parameter
// because 'FileInfoPlus' implements the 'os.FileInfo'
// interface.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fileInfo					os.FileInfo
//
//		The type os.FileInfo contains data elements
//		describing a file.
//
//		type FileInfo interface {
//
//			Name() string
//				base name of the file
//
//			Size() int64
//				length in bytes for regular files;
//				system-dependent for others
//
//			Mode() FileMode
//				file mode bits
//
//			ModTime() time.Time
//				modification time
//
//			IsDir() bool
//				abbreviation for Mode().IsDir()
//
//			Sys() any
//				underlying data source (can return nil)
//		}
//
//	fileSelectCriteria			FileSelectionCriteria
//
//		The FileSelectionCriteria type allows for
//		configuration of single or multiple file
//		selection criterion. The 'SelectCriterionMode'
//		can be used to specify whether the file must
//		match all, or any one, of the active file
//		selection criterion.
//
//		type FileSelectionCriteria struct {
//		  FileNamePatterns     []string
//			An array of strings containing File Name Patterns
//
//		  FilesOlderThan       time.Time
//			Match files with older modification date times
//
//		  FilesNewerThan       time.Time
//			Match files with newer modification date times
//
//		  SelectByFileMode     FilePermissionConfig
//			Match file mode (os.FileMode).
//
//		  SelectCriterionMode  FileSelectCriterionMode
//			Specifies 'AND' or 'OR' selection mode
//		}
//
//		Elements of the FileSelectionCriteria Type are
//		described below:
//
//		FileNamePatterns []string
//
//			An array of strings which may define one or
//			more search patterns. If a file name matches
//			any one of the search pattern strings, it is
//			deemed to be a 'match' for the search pattern
//			criterion.
//
//			Example Patterns:
//			 FileNamePatterns = []string{"*.log"}
//			 FileNamePatterns = []string{"current*.txt"}
//			 FileNamePatterns = []string{"*.txt", "*.log"}
//
//			If this string array has zero length or if
//			all the strings are empty strings, then this
//			file search criterion is considered 'Inactive'
//			or 'Not Set'.
//
//		FilesOlderThan  time.Time
//
//			This date time type is compared to file
//			modification date times in order to determine
//			whether the file is older than the
//			'FilesOlderThan' file selection criterion. If
//			the file is older than the 'FilesOlderThan'
//			date time, that file is considered a 'match'
//			for this file selection criterion.
//
//			If the value of 'FilesOlderThan' is set to
//			time zero, the default value for type
//	 		time.Time{}, then this
//	 		file selection criterion is considered to be
//	 		'Inactive' or 'Not Set'.
//
//		FilesNewerThan   time.Time
//
//			This date time type is compared to the file
//			modification date time in order to determine
//			whether the file is newer than the
//			'FilesNewerThan' file selection criterion. If
//			the file modification date time is newer than
//			the 'FilesNewerThan' date time, that file is
//			considered a 'match' for this file selection
//			criterion.
//
//			If the value of 'FilesNewerThan' is set to
//			time zero, the default value for type
//			time.Time{}, then this file selection
//			criterion is considered to be 'Inactive' or
//			'Not Set'.
//
//		SelectByFileMode  FilePermissionConfig
//
//			Type FilePermissionConfig encapsulates an
//			os.FileMode. The file selection criterion
//			allows for the selection of files by File Mode.
//
//			File modes are compared to the value of
//			'SelectByFileMode'. If the File Mode for a
//			given file is equal to the value of
//			'SelectByFileMode', that file is considered
//			to be a 'match' for this file selection
//			criterion. Examples for setting
//			SelectByFileMode are shown as follows:
//
//			fsc := FileSelectionCriteria{}
//
//			err = fsc.SelectByFileMode.SetByFileMode(os.FileMode(0666))
//
//			err = fsc.SelectByFileMode.SetFileModeByTextCode("-r--r--r--")
//
//		SelectCriterionMode FileSelectCriterionMode
//
//		This parameter selects the manner in which the
//		file selection criteria above are applied in
//		determining a 'match' for file selection
//		purposes. 'SelectCriterionMode' may be set to
//		one of two constant values:
//
//		(1) FileSelectCriterionMode(0).ANDSelect()
//
//			File selected if all active selection criteria
//			are satisfied.
//
//			If this constant value is specified for the
//			file selection mode, then a given file will
//			not be judged as 'selected' unless all the
//			active selection criterion are satisfied. In
//			other words, if three active search criterion
//			are provided for 'FileNamePatterns',
//			'FilesOlderThan' and 'FilesNewerThan', then a
//			file will NOT be selected unless it has
//			satisfied all three criterion in this example.
//
//		(2) FileSelectCriterionMode(0).ORSelect()
//
//			File selected if any active selection
//			criterion is satisfied.
//
//			If this constant value is specified for the
//			file selection mode, then a given file will
//			be selected if any one of the active file
//			selection criterion is satisfied. In other
//			words, if three active search criterion are
//			provided for 'FileNamePatterns',
//			'FilesOlderThan' and 'FilesNewerThan', then a
//			file will be selected if it satisfies any
//			one of the three criterion in this example.
//
// /
// ----------------------------------------------------------------
//
// # Return Values
//
//	isFileNewerThanSet			bool
//
//		This boolean return value signals whether the
//		'FilesNewerThan' file search criteria is set,
//		engaged and active.
//
//		If the 'fileSelectCriteria.FilesNewerThan'
//	 	data element is set to time.Time == zero (0), it
//	 	means that the 'FilesNewerThan' search criteria
//	 	is NOT set or engaged. In this case, the
//	 	'isFileNewerThanSet' parameter will return a
//	 	value of 'false' meaning that the search criteria
//	 	is NOT active or engaged.
//
//		If this parameter is returned with a value of
//		'true', it means that the 'FilesNewerThan' search
//		criterion is set, engaged and active.
//
//	isFileNewerThanMatch		bool
//
//		'fileSelectCriteria.FilesNewerThan' consists of a
//		time.Time date element. This represents a file
//		search pattern designed to identify files with
//		modification dates which are newer than the
//	 	'fileSelectCriteria.FilesNewerThan' date. If
//		the file specified by input parameter 'fileInfo',
//		is newer than this date, the file search
//		operation is classified as a 'match', and
//		the return value of this parameter,
//		'isFileNewerThanMatch', is set to 'true'.
//
//		If this parameter, 'isFileNewerThanMatch', is
//		returned with a value of 'false', it signals that
//		the file has a modification date which is less
//		than or equal to the
//		'fileSelectCriteria.FilesNewerThan' date.
func (fh *FileHelper) SearchFileNewerThan(
	fileInfo os.FileInfo,
	fileSelectCriteria FileSelectionCriteria) (
	isFileNewerThanSet bool,
	isFileNewerThanMatch bool) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	isFileNewerThanSet,
		isFileNewerThanMatch = new(fileHelperElectron).
		searchFileNewerThan(
			fileInfo,
			fileSelectCriteria)

	return isFileNewerThanSet, isFileNewerThanMatch
}

// SearchFileOlderThan
//
// This method is called to determine whether the file
// described by the input parameter 'fileInfo' is a 'match'
// for the File Selection Criteria,
// 'fileSelectCriteria.FilesOlderThan'.
//
// If the file modification date time occurs before the
// 'fileSelectCriteria.FilesOlderThan' date time, the
// return value 'isFileOlderThanMatch' is set to 'true'.
//
// If 'fileSelectCriteria.FilesOlderThan' is set to
// time.Time zero ( the default or zero value for this
// type), the return value 'isFileOlderThanSet' is set to
// 'false' signaling that this search criterion is NOT
// active.
//
// Note:
//
// Input parameter 'fileInfo' is of type os.FileInfo.  This
// means the user may substitute a type 'FileInfoPlus'
// object for the 'fileInfo' parameter because 'FileInfoPlus'
// implements the 'os.FileInfo' interface.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fileInfo					os.FileInfo
//
//		The type os.FileInfo contains data elements
//		describing a file.
//
//		type FileInfo interface {
//			Name() string
//				base name of the file
//
//			Size() int64
//				length in bytes for regular files;
//				system-dependent for others
//
//			Mode() FileMode
//				file mode bits
//
//			ModTime() time.Time
//				modification time
//
//			IsDir() bool
//				abbreviation for Mode().IsDir()
//
//			Sys() any
//				underlying data source (can return nil)
//		}
//
//	fileSelectCriteria			FileSelectionCriteria
//
//		The FileSelectionCriteria type allows for
//		configuration of single or multiple file
//		selection criterion. The 'SelectCriterionMode'
//		can be used to specify whether the file must
//		match all, or any one, of the active file
//		selection criterion.
//
//		type FileSelectionCriteria struct {
//		  FileNamePatterns     []string
//			An array of strings containing File Name Patterns
//
//		  FilesOlderThan       time.Time
//			Match files with older modification date times
//
//		  FilesNewerThan       time.Time
//			Match files with newer modification date times
//
//		  SelectByFileMode     FilePermissionConfig
//			Match file mode (os.FileMode).
//
//		  SelectCriterionMode  FileSelectCriterionMode
//			Specifies 'AND' or 'OR' selection mode
//		}
//
//		Elements of the FileSelectionCriteria Type are
//		described below:
//
//		FileNamePatterns []string
//
//			An array of strings which may define one or
//			more search patterns. If a file name matches
//			any one of the search pattern strings, it is
//			deemed to be a 'match' for the search pattern
//			criterion.
//
//			Example Patterns:
//			 FileNamePatterns = []string{"*.log"}
//			 FileNamePatterns = []string{"current*.txt"}
//			 FileNamePatterns = []string{"*.txt", "*.log"}
//
//			If this string array has zero length or if
//			all the strings are empty strings, then this
//			file search criterion is considered 'Inactive'
//			or 'Not Set'.
//
//		FilesOlderThan  time.Time
//
//			This date time type is compared to file
//			modification date times in order to determine
//			whether the file is older than the
//			'FilesOlderThan' file selection criterion. If
//			the file is older than the 'FilesOlderThan'
//			date time, that file is considered a 'match'
//			for this file selection criterion.
//
//			If the value of 'FilesOlderThan' is set to
//			time zero, the default value for type
//	 		time.Time{}, then this
//	 		file selection criterion is considered to be
//	 		'Inactive' or 'Not Set'.
//
//		FilesNewerThan   time.Time
//
//			This date time type is compared to the file
//			modification date time in order to determine
//			whether the file is newer than the
//			'FilesNewerThan' file selection criterion. If
//			the file modification date time is newer than
//			the 'FilesNewerThan' date time, that file is
//			considered a 'match' for this file selection
//			criterion.
//
//			If the value of 'FilesNewerThan' is set to
//			time zero, the default value for type
//			time.Time{}, then this file selection
//			criterion is considered to be 'Inactive' or
//			'Not Set'.
//
//		SelectByFileMode  FilePermissionConfig
//
//			Type FilePermissionConfig encapsulates an
//			os.FileMode. The file selection criterion
//			allows for the selection of files by File Mode.
//
//			File modes are compared to the value of
//			'SelectByFileMode'. If the File Mode for a
//			given file is equal to the value of
//			'SelectByFileMode', that file is considered
//			to be a 'match' for this file selection
//			criterion. Examples for setting
//			SelectByFileMode are shown as follows:
//
//			fsc := FileSelectionCriteria{}
//
//			err = fsc.SelectByFileMode.SetByFileMode(os.FileMode(0666))
//
//			err = fsc.SelectByFileMode.SetFileModeByTextCode("-r--r--r--")
//
//		SelectCriterionMode FileSelectCriterionMode
//
//		This parameter selects the manner in which the
//		file selection criteria above are applied in
//		determining a 'match' for file selection
//		purposes. 'SelectCriterionMode' may be set to
//		one of two constant values:
//
//		(1) FileSelectCriterionMode(0).ANDSelect()
//
//			File selected if all active selection criteria
//			are satisfied.
//
//			If this constant value is specified for the
//			file selection mode, then a given file will
//			not be judged as 'selected' unless all the
//			active selection criterion are satisfied. In
//			other words, if three active search criterion
//			are provided for 'FileNamePatterns',
//			'FilesOlderThan' and 'FilesNewerThan', then a
//			file will NOT be selected unless it has
//			satisfied all three criterion in this example.
//
//		(2) FileSelectCriterionMode(0).ORSelect()
//
//			File selected if any active selection
//			criterion is satisfied.
//
//			If this constant value is specified for the
//			file selection mode, then a given file will
//			be selected if any one of the active file
//			selection criterion is satisfied. In other
//			words, if three active search criterion are
//			provided for 'FileNamePatterns',
//			'FilesOlderThan' and 'FilesNewerThan', then a
//			file will be selected if it satisfies any
//			one of the three criterion in this example.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	isFileOlderThanSet			bool
//
//		This boolean return value signals whether the
//		FileOlderThan file search criteria is set,
//		engaged and active.
//
//		If the 'fileSelectCriteria.FilesOlderThan'
//	 	data element is set to time.Time == zero (0), it
//	 	means that the 'FilesOlderThan' search criteria is
//	 	NOT set or engaged. In this case, the
//	 	'isFileOlderThanSet' parameter will return a
//	 	value of 'false' meaning that the search criteria
//	 	is NOT active or engaged.
//
//		If this parameter is returned with a value of
//		'true', it means that the 'FilesOlderThan' search
//		criterion is set, engaged and active.
//
//	isFileOlderThanMatch		bool
//
//		'fileSelectCriteria.FilesOlderThan' consists of a
//		time.Time date element. This represents a file
//		search pattern designed to identify files with
//		modification dates which are older than the
//	 	'fileSelectCriteria.FilesOlderThan' date. If
//		the file specified by input parameter 'fileInfo',
//		is older than this date, the file search
//		operation is classified as a 'match', and
//		the return value of this parameter,
//		'isFileOlderThanMatch', is set to 'true'.
//
//		If this parameter, 'isFileOlderThanMatch', is
//		returned with a value of 'false', it signals that
//		the file has a modification date which is greater
//		than or equal to the
//		'fileSelectCriteria.FilesOlderThan' date.
func (fh *FileHelper) SearchFileOlderThan(
	fileInfo os.FileInfo,
	fileSelectCriteria FileSelectionCriteria) (
	isFileOlderThanSet bool,
	isFileOlderThanMatch bool) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	return new(fileHelperElectron).
		searchFileOlderThan(
			fileInfo,
			fileSelectCriteria)
}

// SearchFilePatternMatch
//
// This method is used to determine whether a file
// described by the 'fileInfo' parameter meets the
// specified File Selection Criteria and is judged to be
// a match for the fileSelectCriteria FileNamePattern.
//
// 'fileSelectCriteria.FileNamePatterns' consist of a
// string array. If the pattern signified by any element
// in the string array is a 'match', the return value
// 'isPatternMatch' is set to true.
//
// If the 'fileSelectCriteria.FileNamePatterns' array is
// empty or if it contains only empty strings, the return
// value isPatternSet is set to 'false' signaling that
// the pattern file search selection criterion is NOT
// active.
//
// Note:
//
// Input parameter 'fileInfo' is of type os.FileInfo. You
// can substitute a type 'FileInfoPlus' object for the
// 'fileInfo' parameter because 'FileInfoPlus' implements
// the 'os.FileInfo' interface.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fileInfo					os.FileInfo
//
//		The type os.FileInfo contains data elements
//		describing a file.
//
//		type FileInfo interface {
//			Name() string
//				base name of the file
//
//			Size() int64
//				length in bytes for regular files;
//				system-dependent for others
//
//			Mode() FileMode
//				file mode bits
//
//			ModTime() time.Time
//				modification time
//
//			IsDir() bool
//				abbreviation for Mode().IsDir()
//
//			Sys() any
//				underlying data source (can return nil)
//		}
//
//	fileSelectCriteria			FileSelectionCriteria
//
//		The FileSelectionCriteria type allows for
//		configuration of single or multiple file
//		selection criterion. The 'SelectCriterionMode'
//		can be used to specify whether the file must
//		match all, or any one, of the active file
//		selection criterion.
//
//		type FileSelectionCriteria struct {
//		  FileNamePatterns     []string
//			An array of strings containing File Name Patterns
//
//		  FilesOlderThan       time.Time
//			Match files with older modification date times
//
//		  FilesNewerThan       time.Time
//			Match files with newer modification date times
//
//		  SelectByFileMode     FilePermissionConfig
//			Match file mode (os.FileMode).
//
//		  SelectCriterionMode  FileSelectCriterionMode
//			Specifies 'AND' or 'OR' selection mode
//		}
//
//		Elements of the FileSelectionCriteria Type are
//		described below:
//
//		FileNamePatterns []string
//
//			An array of strings which may define one or
//			more search patterns. If a file name matches
//			any one of the search pattern strings, it is
//			deemed to be a 'match' for the search pattern
//			criterion.
//
//			Example Patterns:
//			 FileNamePatterns = []string{"*.log"}
//			 FileNamePatterns = []string{"current*.txt"}
//			 FileNamePatterns = []string{"*.txt", "*.log"}
//
//			If this string array has zero length or if
//			all the strings are empty strings, then this
//			file search criterion is considered 'Inactive'
//			or 'Not Set'.
//
//		FilesOlderThan  time.Time
//
//			This date time type is compared to file
//			modification date times in order to determine
//			whether the file is older than the
//			'FilesOlderThan' file selection criterion. If
//			the file is older than the 'FilesOlderThan'
//			date time, that file is considered a 'match'
//			for this file selection criterion.
//
//			If the value of 'FilesOlderThan' is set to
//			time zero, the default value for type
//	 		time.Time{}, then this
//	 		file selection criterion is considered to be
//	 		'Inactive' or 'Not Set'.
//
//		FilesNewerThan   time.Time
//
//			This date time type is compared to the file
//			modification date time in order to determine
//			whether the file is newer than the
//			'FilesNewerThan' file selection criterion. If
//			the file modification date time is newer than
//			the 'FilesNewerThan' date time, that file is
//			considered a 'match' for this file selection
//			criterion.
//
//			If the value of 'FilesNewerThan' is set to
//			time zero, the default value for type
//			time.Time{}, then this file selection
//			criterion is considered to be 'Inactive' or
//			'Not Set'.
//
//		SelectByFileMode  FilePermissionConfig
//
//			Type FilePermissionConfig encapsulates an
//			os.FileMode. The file selection criterion
//			allows for the selection of files by File Mode.
//
//			File modes are compared to the value of
//			'SelectByFileMode'. If the File Mode for a
//			given file is equal to the value of
//			'SelectByFileMode', that file is considered
//			to be a 'match' for this file selection
//			criterion. Examples for setting
//			SelectByFileMode are shown as follows:
//
//			fsc := FileSelectionCriteria{}
//
//			err = fsc.SelectByFileMode.SetByFileMode(os.FileMode(0666))
//
//			err = fsc.SelectByFileMode.SetFileModeByTextCode("-r--r--r--")
//
//		SelectCriterionMode FileSelectCriterionMode
//
//		This parameter selects the manner in which the
//		file selection criteria above are applied in
//		determining a 'match' for file selection
//		purposes. 'SelectCriterionMode' may be set to
//		one of two constant values:
//
//		(1) FileSelectCriterionMode(0).ANDSelect()
//
//			File selected if all active selection criteria
//			are satisfied.
//
//			If this constant value is specified for the
//			file selection mode, then a given file will
//			not be judged as 'selected' unless all the
//			active selection criterion are satisfied. In
//			other words, if three active search criterion
//			are provided for 'FileNamePatterns',
//			'FilesOlderThan' and 'FilesNewerThan', then a
//			file will NOT be selected unless it has
//			satisfied all three criterion in this example.
//
//		(2) FileSelectCriterionMode(0).ORSelect()
//
//			File selected if any active selection
//			criterion is satisfied.
//
//			If this constant value is specified for the
//			file selection mode, then a given file will
//			be selected if any one of the active file
//			selection criterion is satisfied. In other
//			words, if three active search criterion are
//			provided for 'FileNamePatterns',
//			'FilesOlderThan' and 'FilesNewerThan', then a
//			file will be selected if it satisfies any
//			one of the three criterion in this example.
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
//	isPatternSet				bool
//
//		This boolean return value signals whether one or
//		more file search patterns are set, engaged and
//		active.
//
//		If the 'fileSelectCriteria.FileNamePatterns'
//	 	array is empty or if it contains only empty
//	  	strings, this return value is set to 'false'
//	  	signaling that the pattern file search
//	   	selection criterion is NOT active.
//
//		When set to 'true', this parameter signals that
//		one or more file search patterns are set, engaged
//		and active.
//
//	isPatternMatch				bool
//
//		'fileSelectCriteria.FileNamePatterns' consist of
//		a string array. If the pattern signified by any
//		element in the string array is a 'match', the
//		return value 'isPatternMatch' is set to true.
//
//		The criteria for a 'match' can be set through
//		input paramter,
//		'fileSelectCriteria.SelectCriterionMode'.
//
//	msgError					error
//
//		'msgError' is a standard error containing
//		a brief high level message explaining the
//		error condition in narrative text.
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If 'msgError' is returned with a value of
//		'nil', it signals that the deletion operation
//		succeeded.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
//
//		NOTE: Not all returned errors have associated
//		low level errors ('lowLevelErr'). Always check
//		'msgError'. 'msgError' could be non-nil while
//		'lowLevelErr' is 'nil'.
//
//	lowLevelErr					error
//
//		If the call to path.Match() fails it will return
//		an error of type ErrBadPattern. This error may be
//		unpacked to reveal additional technical details
//		regarding the failure to create a file.
//
//		If an error is returned by path.Match() it will
//		be returned in its original form through
//		parameter, 'lowLevelErr'.
//
//		If no ErrBadPattern occurs, 'lowLevelErr' will be
//		set to 'nil'.
func (fh *FileHelper) SearchFilePatternMatch(
	fileInfo os.FileInfo,
	fileSelectCriteria FileSelectionCriteria,
	errorPrefix interface{}) (
	isPatternSet bool,
	isPatternMatch bool,
	msgError error,
	lowLevelErr error) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	isPatternMatch = false

	isPatternSet = false

	ePrefix,
		msgError = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"SearchFilePatternMatch()",
		"")

	if msgError != nil {

		return isPatternSet, isPatternMatch, msgError, lowLevelErr
	}

	isPatternSet,
		isPatternMatch,
		msgError,
		lowLevelErr = new(fileHelperElectron).
		searchFilePatternMatch(
			fileInfo,
			fileSelectCriteria,
			ePrefix)

	return isPatternSet, isPatternMatch, msgError, lowLevelErr
}

// StripLeadingDotSeparatorChars
//
// Strips or deletes the following characters from the
// front of path or directory names.
//
// Leading Characters To Be Removed:
//
//	(1)	" " (Space)
//
//	(2)	PathSeparator
//
//	(3)	"."
//
//	(4)	".."
//
//	(5)	"." + PathSeparator
//
//	(6)	".." + PathSeparator
//
// Removal of these characters will convert the path or
// directory name to a valid set of text characters
// suitable as input for file or directory processing
// functions.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathName					string
//
//		The path or directory name to be processed. This
//		method will strip or delete selected characters
//		from the front or right side of this string.
//		Removal of these characters will convert the path
//		or directory name to a valid set of text
//		characters suitable as input for file or
//	 	directory processing functions.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	string
//
//		The converted path or directory name form which
//		selected invalid characters have been removed.
//
//	int
//
//		The string length of the returned 'string'
//		parameter.
func (fh *FileHelper) StripLeadingDotSeparatorChars(pathName string) (string, int) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	return new(fileHelperMolecule).
		stripLeadingDotSeparatorChars(
			pathName)
}

// SwapBasePath
//
// Searches the 'targetPath' string for the existence of
// 'oldBasePath'. If 'oldBasePath' is found, it is
// replaced with 'newBasePath' in 'targetPath'.
//
// If 'oldBasePath' is not found in 'targetPath' an error
// is returned.
//
// Likewise, if 'oldBasePath' is not located at the
// beginning of 'targetPath', an error will be returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	oldBasePath					string
//
//		The Old Base Path. If this path is found at the
//		beginning of 'targetPath' it will be replaced by
//		'newBasePath' and returned to the caller.
//
//		If 'oldBasePath' is NOT found at the beginning of
//		'targetPath', an error will be returned.
//
//	newBasePath					string
//
//		The New Base Path. If 'oldBasePath' is found at
//		the beginning of 'targetPath', 'oldBasePath' will
//		be replaced with 'newBasePath' in 'targetPath'.
//
//		If 'oldBasePath' is NOT found at the beginning of
//		'targetPath', an error will be returned.
//
//	targetPath					string
//
//		This string contains the entire path.
//
//		If 'oldBasePath' is not found at the beginning of
//		'targetPath', an error will be returned.
//
//		If 'oldBasePath' is found at the beginning of
//		'targetPath', it will be replaced by 'newBasePath'
//		and returned to the caller.
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
//		This string parameter returns a new version of
//		input parameter 'targetPath' in which
//		'oldBasePath' has been replaced with
//		'newBasePath'.
//
//		If 'oldBasePath' is not found at the beginning of
//		'targetPath', an error will be returned.
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
func (fh *FileHelper) SwapBasePath(
	oldBasePath string,
	newBasePath string,
	targetPath string,
	errorPrefix interface{}) (string, error) {

	if fh.lock == nil {
		fh.lock = new(sync.Mutex)
	}

	fh.lock.Lock()

	defer fh.lock.Unlock()

	funcName := "FileHelper.SwapBasePath() "

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {
		return "", err
	}

	return new(fileHelperAtom).
		swapBasePath(
			oldBasePath,
			newBasePath,
			targetPath,
			ePrefix)
}
