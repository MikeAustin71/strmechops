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
//		input parameter, 'errPrefDto'. The 'errPrefDto'
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

// deleteAllFilesInDirectory
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
//		input parameter, 'errPrefDto'. The 'errPrefDto'
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
func (fHelperAtom *fileHelperAtom) deleteAllFilesInDirectory(
	directoryPath string,
	errPrefDto *ePref.ErrPrefixDto) (
	msgError error,
	lowLevelErr error) {

	if fHelperAtom.lock == nil {
		fHelperAtom.lock = new(sync.Mutex)
	}

	fHelperAtom.lock.Lock()

	defer fHelperAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		msgError = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileHelperAtom."+
			"deleteAllFilesInDirectory()",
		"")

	if msgError != nil {
		return msgError, lowLevelErr
	}

	var ptrDir *os.File

	ptrDir, lowLevelErr = os.Open(directoryPath)

	if lowLevelErr != nil {

		msgError = fmt.Errorf("%v\n"+
			"Error: os.Open(directoryPath) Failed!\n"+
			"directoryPath = '%v'\n"+
			"Error = \n%v\n",
			ePrefix.String(),
			directoryPath,
			lowLevelErr.Error())

		return msgError, lowLevelErr
	}

	var dirFiles []os.DirEntry

	dirFiles, lowLevelErr = ptrDir.ReadDir(0)

	if lowLevelErr != nil {

		msgError = fmt.Errorf("%v\n"+
			"Error: ptrDir.ReadDir(0) Failed!\n"+
			"ReadDir() did not return any files.\n"+
			"directoryPath = '%v'\n"+
			"Error = \n%v\n",
			ePrefix.String(),
			directoryPath,
			lowLevelErr.Error())

		return msgError, lowLevelErr

	}

	var dirPrefix string

	dirPrefix,
		msgError = new(fileHelperProton).
		addPathSeparatorToEndOfPathStr(
			directoryPath,
			ePrefix.XCpy("pathStr"))

	var fInfo os.FileInfo
	var fileName string

	for index, dirFile := range dirFiles {

		fInfo,
			lowLevelErr = dirFile.Info()

		if lowLevelErr != nil {

			msgError = fmt.Errorf("%v\n"+
				"Error: ptrDir.ReadDir(0) Failed!\n"+
				"ReadDir() did not return any files.\n"+
				"directoryPath = '%v'\n"+
				"Index = '%v'"+
				"Error = \n%v\n",
				ePrefix.String(),
				directoryPath,
				index,
				lowLevelErr.Error())

			lowLevelErr = nil

			return msgError, lowLevelErr

		}

		if fInfo.IsDir() {

			continue
		}

		fileName = fInfo.Name()

		fileName = dirPrefix + fileName

		lowLevelErr = os.Remove(fileName)

		if lowLevelErr != nil {

			msgError = fmt.Errorf("%v\n"+
				"Error: os.Remove(fileName) Failed!\n"+
				"directoryPath = '%v'\n"+
				"Index = '%v'"+
				"fileName = '%v'\n"+
				"Error = \n%v\n",
				ePrefix.String(),
				directoryPath,
				index,
				fileName,
				lowLevelErr.Error())

			return msgError, lowLevelErr

		}

	}

	return msgError, lowLevelErr
}

// doesStringEndWithPathSeparator
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
func (fHelperAtom *fileHelperAtom) doesStringEndWithPathSeparator(
	pathStr string) bool {

	if fHelperAtom.lock == nil {
		fHelperAtom.lock = new(sync.Mutex)
	}

	fHelperAtom.lock.Lock()

	defer fHelperAtom.lock.Unlock()

	errCode := 0
	lenStr := 0

	errCode, lenStr, pathStr =
		new(fileHelperElectron).isStringEmptyOrBlank(pathStr)

	if errCode < 0 {
		return false
	}

	if pathStr[lenStr-1] == '\\' || pathStr[lenStr-1] == '/' || pathStr[lenStr-1] == os.PathSeparator {
		return true
	}

	return false
}

// doesThisFileExist
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
//		If errors are encountered during processing, the
//		returned 'nonPathError' will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errPrefDto'.
//	 	The 'errPrefDto' text will be attached to the
//	 	beginning of the error message.
func (fHelperAtom *fileHelperAtom) doesThisFileExist(
	pathFileName string,
	errPrefDto *ePref.ErrPrefixDto) (
	pathFileNameDoesExist bool,
	nonPathError error) {

	if fHelperAtom.lock == nil {
		fHelperAtom.lock = new(sync.Mutex)
	}

	fHelperAtom.lock.Lock()

	defer fHelperAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	pathFileNameDoesExist = false

	ePrefix,
		nonPathError = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileHelperAtom."+
			"doesThisFileExist()",
		"")

	if nonPathError != nil {

		return pathFileNameDoesExist, nonPathError
	}

	_,
		pathFileNameDoesExist,
		_,
		nonPathError =
		new(fileHelperMolecule).doesPathFileExist(
			pathFileName,
			PreProcPathCode.AbsolutePath(), // Skip Absolute Path Conversion
			ePrefix,
			"pathFileName")

	if nonPathError != nil {

		pathFileNameDoesExist = false

		return pathFileNameDoesExist, nonPathError
	}

	return pathFileNameDoesExist, nonPathError
}

// filterFileName
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
// file name (info.Name()) passed as input parameter
// 'info'.
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
// If a given criterion is set to a non-zero value, then
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
//	info						os.FileInfo
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
//	isMatchedFile				bool
//
//		If this return parameter is set to 'true' it
//		signals that the file identified by input
//		parameter 'info' has matched the specified
//		file selection criteria.
//
//		If 'isMatchedFile' is returned with a value of
//		'false', it means that the file identified by
//		input parameter 'info' does NOT match the
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
//		input parameter, 'errPrefDto'. The 'errPrefDto'
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
func (fHelperAtom *fileHelperAtom) filterFileName(
	fileInfo os.FileInfo,
	fileSelectionCriteria FileSelectionCriteria,
	errPrefDto *ePref.ErrPrefixDto) (
	isMatchedFile bool,
	msgError error,
	lowLevelErr error) {

	if fHelperAtom.lock == nil {
		fHelperAtom.lock = new(sync.Mutex)
	}

	fHelperAtom.lock.Lock()

	defer fHelperAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	isMatchedFile = false

	ePrefix,
		msgError = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileHelperAtom."+
			"filterFileName()",
		"")

	if msgError != nil {

		return isMatchedFile, msgError, lowLevelErr
	}

	if fileInfo == nil {

		msgError = fmt.Errorf("%v\n"+
			"Input parameter 'fileInfo' is 'nil' and INVALID!",
			ePrefix.String())

		return isMatchedFile, msgError, lowLevelErr
	}

	var isPatternSet, isPatternMatch bool
	var err2 error

	fHelperElectron := new(fileHelperElectron)

	isPatternSet,
		isPatternMatch,
		err2,
		lowLevelErr = fHelperElectron.
		searchFilePatternMatch(
			fileInfo,
			fileSelectionCriteria,
			ePrefix)

	if err2 != nil {

		msgError = fmt.Errorf(
			"Error returned from SearchFilePatternMatch(fileInfo, fileSelectionCriteria).\n"+
				"fileInfo.Name()='%v'\n"+
				"Message Error=\n%v'\n\n"+
				"Low Level Error=\n%v\n",
			fileInfo.Name(),
			err2.Error(),
			lowLevelErr.Error())

		isMatchedFile = false

		return isMatchedFile, msgError, lowLevelErr
	}

	isFileOlderThanSet,
		isFileOlderThanMatch := fHelperElectron.
		searchFileOlderThan(
			fileInfo,
			fileSelectionCriteria)

	var isFileNewerThanSet, isFileNewerThanMatch bool

	isFileNewerThanSet,
		isFileNewerThanMatch =
		fHelperElectron.
			searchFileNewerThan(
				fileInfo,
				fileSelectionCriteria)

	var isFileModeSearchSet, isFileModeSearchMatch bool

	isFileModeSearchSet,
		isFileModeSearchMatch,
		err2 = fHelperElectron.
		searchFileModeMatch(
			fileInfo,
			fileSelectionCriteria,
			ePrefix)

	if err2 != nil {

		var err3 error
		var errs []error
		var fileModeTxt string

		fileModeTxt,
			err3 = fileSelectionCriteria.SelectByFileMode.
			GetPermissionFileModeValueText(
				ePrefix)

		if err3 != nil {

			errs = append(
				errs, err2)

			errs = append(
				errs, err3)

			err2 = new(StrMech).ConsolidateErrors(errs)

		}

		msgError = fmt.Errorf("%v\n"+
			"Error returned from searchFileModeMatch(fileInfo, fileSelectionCriteria).\n"+
			"fileSelectionCriteria.SelectByFileMode='%v'\n"+
			"fileInfo.Name()='%v' Error=\n%v\n",
			ePrefix.String(),
			fileModeTxt,
			fileInfo.Name(), err2.Error())

		isMatchedFile = false

		return isMatchedFile, msgError, lowLevelErr
	}

	// If no file selection criterion are set, then always select the file
	if !isPatternSet &&
		!isFileOlderThanSet &&
		!isFileNewerThanSet &&
		!isFileModeSearchSet {

		isMatchedFile = true

		return isMatchedFile, msgError, lowLevelErr
	}

	// If using the AND File Select Criterion Mode, then for criteria that
	// are set and active, they must all be 'matched'.
	if fileSelectionCriteria.SelectCriterionMode == FileSelectMode.ANDSelect() {

		if isPatternSet && !isPatternMatch {

			isMatchedFile = false

			return isMatchedFile, msgError, lowLevelErr
		}

		if isFileOlderThanSet && !isFileOlderThanMatch {

			isMatchedFile = false

			return isMatchedFile, msgError, lowLevelErr
		}

		if isFileNewerThanSet && !isFileNewerThanMatch {

			isMatchedFile = false

			return isMatchedFile, msgError, lowLevelErr
		}

		if isFileModeSearchSet && !isFileModeSearchMatch {

			isMatchedFile = false

			return isMatchedFile, msgError, lowLevelErr
		}

		isMatchedFile = true

		return isMatchedFile, msgError, lowLevelErr

	} // End of fileSelectMode.ANDSelect()

	// Must be fileSelectMode.ORSelect() Mode
	// If ANY of the section criterion are active and 'matched', then
	// classify the file as matched.

	if isPatternSet && isPatternMatch {

		isMatchedFile = true

		return isMatchedFile, msgError, lowLevelErr
	}

	if isFileOlderThanSet && isFileOlderThanMatch {

		isMatchedFile = true

		return isMatchedFile, msgError, lowLevelErr
	}

	if isFileNewerThanSet && isFileNewerThanMatch {

		isMatchedFile = true

		return isMatchedFile, msgError, lowLevelErr
	}

	if isFileModeSearchSet && isFileModeSearchMatch {

		isMatchedFile = true

		return isMatchedFile, msgError, lowLevelErr
	}

	isMatchedFile = false

	return isMatchedFile, msgError, lowLevelErr
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

// removePathSeparatorFromEndOfPathString
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
func (fHelperAtom *fileHelperAtom) removePathSeparatorFromEndOfPathString(
	pathStr string) string {

	if fHelperAtom.lock == nil {
		fHelperAtom.lock = new(sync.Mutex)
	}

	fHelperAtom.lock.Lock()

	defer fHelperAtom.lock.Unlock()

	errCode := 0
	lPathStr := 0

	errCode,
		lPathStr,
		pathStr = new(fileHelperElectron).
		isStringEmptyOrBlank(pathStr)

	if errCode < 0 {
		return ""

	}

	lastChar := rune(pathStr[lPathStr-1])

	if lastChar == os.PathSeparator ||
		lastChar == '\\' ||
		lastChar == '/' {

		return pathStr[0 : lPathStr-1]
	}

	return pathStr
}
