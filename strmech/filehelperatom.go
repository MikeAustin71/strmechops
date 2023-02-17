package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"os"
	fp "path/filepath"
	"strings"
	"sync"
)

type fileHelperAtom struct {
	lock *sync.Mutex
}

// adjustPathSlash
//
// This method will standardize path separators according
// to the current operating system.
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
func (fHelperAtom *fileHelperAtom) adjustPathSlash(
	path string) string {

	if fHelperAtom.lock == nil {
		fHelperAtom.lock = new(sync.Mutex)
	}

	fHelperAtom.lock.Lock()

	defer fHelperAtom.lock.Unlock()

	errCode := 0

	errCode, _, path = new(fileHelperElectron).
		isStringEmptyOrBlank(path)

	if errCode == -1 {
		return ""
	}

	if errCode == -2 {
		return ""
	}

	if os.PathSeparator != '\\' {
		return strings.ReplaceAll(path, "\\", string(os.PathSeparator))
	}

	if os.PathSeparator != '/' {
		return strings.ReplaceAll(path, "/", string(os.PathSeparator))
	}

	return fp.FromSlash(path)
}

// createFile
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
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates error prefix text which
//		is included in the error message returned by
//		'msgError'.
//
//		Usually, 'errPrefDto' contains the name of the
//	 	calling method or methods listed as a method or
//	  	function chain of execution.
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
//		to 'nil'.
func (fHelperAtom *fileHelperAtom) createFile(
	pathFileName string,
	errPrefDto *ePref.ErrPrefixDto) (
	ptrOsFile *os.File,
	msgError error,
	lowLevelErr error) {

	if fHelperAtom.lock == nil {
		fHelperAtom.lock = new(sync.Mutex)
	}

	fHelperAtom.lock.Lock()

	defer fHelperAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ptrOsFile = nil

	ePrefix,
		msgError = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileHelperAtom."+
			"createFile()",
		"")

	if msgError != nil {
		return ptrOsFile, msgError, lowLevelErr
	}

	var errCode int

	errCode,
		_,
		pathFileName = new(fileHelperElectron).
		isStringEmptyOrBlank(pathFileName)

	if errCode == -1 {

		msgError = fmt.Errorf("%v\n"+
			"Error: Input parameter 'pathFileName' is an empty string!\n",
			ePrefix)

		return ptrOsFile, msgError, lowLevelErr
	}

	if errCode == -2 {

		msgError = fmt.Errorf("%v\n"+
			"Error: Input parameter 'pathFileName' consists of blank spaces!\n",
			ePrefix)

		return ptrOsFile, msgError, lowLevelErr
	}

	pathFileName,
		msgError = new(fileHelperProton).
		makeAbsolutePath(
			pathFileName,
			ePrefix)

	if msgError != nil {

		return ptrOsFile, msgError, lowLevelErr

	}

	ptrOsFile,
		lowLevelErr = os.Create(pathFileName)

	if lowLevelErr != nil {

		msgError = fmt.Errorf("%v\n"+
			"Error returned from os.Create(pathFileName)\n"+
			"pathFileName='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			pathFileName,
			lowLevelErr.Error())
	}

	return ptrOsFile, msgError, lowLevelErr
}

// changeWorkingDir
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
func (fHelperAtom *fileHelperAtom) changeWorkingDir(
	dirPath string,
	errorPrefix interface{}) error {

	if fHelperAtom.lock == nil {
		fHelperAtom.lock = new(sync.Mutex)
	}

	fHelperAtom.lock.Lock()

	defer fHelperAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"fileHelperAtom."+
			"changeWorkingDir()",
		"")

	if err != nil {
		return err
	}

	errCode := 0

	errCode, _, dirPath =
		new(fileHelperElectron).isStringEmptyOrBlank(dirPath)

	if errCode == -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'dirPath' is an empty string!\n",
			ePrefix.String())

		return err
	}

	if errCode == -2 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'dirPath' consists of blank spaces!\n",
			ePrefix.String())

		return err
	}

	var err2 error

	err2 = os.Chdir(dirPath)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by os.Chdir(dirPath).\n"+
			"dirPath='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			dirPath,
			err2.Error())
	}

	return err
}

// getAbsCurrDir
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
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (fHelperAtom *fileHelperAtom) getAbsCurrDir(
	errPrefDto *ePref.ErrPrefixDto) (
	string,
	error) {

	if fHelperAtom.lock == nil {
		fHelperAtom.lock = new(sync.Mutex)
	}

	fHelperAtom.lock.Lock()

	defer fHelperAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileHelperAtom."+
			"getAbsCurrDir()",
		"")

	if err != nil {
		return "", err
	}

	var dir string
	var err2 error

	dir, err2 = os.Getwd()

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned from os.Getwd().\n"+
			"Error='%v'\n",
			ePrefix.String(),
			err2.Error())

		return "", err
	}

	var absDir string

	absDir,
		err = new(fileHelperProton).
		makeAbsolutePath(
			dir,
			ePrefix.XCpy(
				"absDir<-dir"))

	if err != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by fh.MakeAbsolutePath(dir).\n"+
			"Error='%v'\n",
			ePrefix.String(),
			err.Error())

	}

	return absDir, err
}

// getDotSeparatorIndexesInPathStr
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
//		An array of integers specifying the string
//		indexes of the dot ('.') characters located in
//		input parameter string 'pathStr'.
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
func (fHelperAtom *fileHelperAtom) getDotSeparatorIndexesInPathStr(
	pathStr string,
	errPrefDto *ePref.ErrPrefixDto) (
	[]int,
	error) {

	if fHelperAtom.lock == nil {
		fHelperAtom.lock = new(sync.Mutex)
	}

	fHelperAtom.lock.Lock()

	defer fHelperAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	var dotIdxs []int

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileHelperAtom."+
			"getDotSeparatorIndexesInPathStr()",
		"")

	if err != nil {
		return dotIdxs, err
	}

	errCode := 0

	lPathStr := 0

	errCode,
		lPathStr,
		pathStr = new(fileHelperElectron).
		isStringEmptyOrBlank(pathStr)

	if errCode == -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'pathStr' is an empty string!\n",
			ePrefix.String())

		return dotIdxs, err
	}

	if errCode == -2 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'pathStr' consists of blank spaces!\n",
			ePrefix.String())

		return dotIdxs, err
	}

	for i := 0; i < lPathStr; i++ {

		rChar := pathStr[i]

		if rChar == '.' {

			dotIdxs = append(dotIdxs, i)
		}

	}

	return dotIdxs, nil
}

// getPathSeparatorIndexesInPathStr
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
func (fHelperAtom *fileHelperAtom) getPathSeparatorIndexesInPathStr(
	pathStr string,
	errPrefDto *ePref.ErrPrefixDto) (
	[]int,
	error) {

	if fHelperAtom.lock == nil {
		fHelperAtom.lock = new(sync.Mutex)
	}

	fHelperAtom.lock.Lock()

	defer fHelperAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileHelperAtom."+
			"getPathSeparatorIndexesInPathStr()",
		"")

	if err != nil {
		return []int{}, err
	}
	errCode := 0
	lPathStr := 0

	errCode, lPathStr, pathStr =
		new(fileHelperElectron).isStringEmptyOrBlank(pathStr)

	if errCode == -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'pathStr' consists of blank spaces!\n",
			ePrefix.String())

		return []int{}, err
	}

	if errCode == -2 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'pathStr' consists of blank spaces!\n",
			ePrefix.String())

		return []int{}, err
	}

	var slashIdxs []int

	for i := 0; i < lPathStr; i++ {

		rChar := pathStr[i]

		if rChar == os.PathSeparator ||
			rChar == '\\' ||
			rChar == '/' {

			slashIdxs = append(slashIdxs, i)
		}

	}

	return slashIdxs, err
}
