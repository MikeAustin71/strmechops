package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"os"
	fp "path/filepath"
	"strings"
	"sync"
	"time"
)

type fileHelperNanobot struct {
	lock *sync.Mutex
}

// areSameFile
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
func (fHelperNanobot *fileHelperNanobot) areSameFile(
	pathFile1,
	pathFile2 string,
	errorPrefix interface{}) (
	bool,
	error) {

	if fHelperNanobot.lock == nil {
		fHelperNanobot.lock = new(sync.Mutex)
	}

	fHelperNanobot.lock.Lock()

	defer fHelperNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"fileHelperNanobot."+
			"areSameFile()",
		"")

	if err != nil {
		return false, err
	}

	var pathFile1DoesExist, pathFile2DoesExist bool
	var fInfoPathFile1, fInfoPathFile2 FileInfoPlus

	pathFile1,
		pathFile1DoesExist,
		fInfoPathFile1,
		err = new(fileHelperMolecule).doesPathFileExist(
		pathFile1,
		PreProcPathCode.AbsolutePath(), // Convert To Absolute Path
		ePrefix,
		"pathFile1")

	if err != nil {
		return false, err
	}

	pathFile2,
		pathFile2DoesExist,
		fInfoPathFile2,
		err = new(fileHelperMolecule).doesPathFileExist(pathFile2,
		PreProcPathCode.AbsolutePath(), // Convert To Absolute Path
		ePrefix,
		"pathFile2")

	if err != nil {
		return false, err
	}

	pathFile1 = strings.ToLower(pathFile1)
	pathFile2 = strings.ToLower(pathFile2)

	if pathFile1DoesExist && pathFile2DoesExist {

		if os.SameFile(fInfoPathFile1.GetOriginalFileInfo(), fInfoPathFile2.GetOriginalFileInfo()) ||
			pathFile1 == pathFile2 {
			// pathFile1 and pathFile2 are the same
			// path and file name.

			return true, nil

		}

		return false, nil
	}

	if pathFile1 == pathFile2 {
		return true, nil
	}

	return false, nil
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
func (fHelperNanobot *fileHelperNanobot) changeFileMode(
	pathFileName string,
	filePermission FilePermissionConfig,
	errorPrefix interface{}) error {

	if fHelperNanobot.lock == nil {
		fHelperNanobot.lock = new(sync.Mutex)
	}

	fHelperNanobot.lock.Lock()

	defer fHelperNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"fileHelperNanobot."+
			"changeFileMode()",
		"")

	if err != nil {
		return err
	}

	var filePathDoesExist bool

	pathFileName,
		filePathDoesExist,
		_,
		err = new(fileHelperMolecule).doesPathFileExist(
		pathFileName,
		PreProcPathCode.AbsolutePath(), // Convert to Absolute Path
		ePrefix,
		"pathFileName")

	if err != nil {
		return err
	}

	if !filePathDoesExist {

		err = fmt.Errorf("%v\n"+
			"ERROR: 'pathFileName' DOES NOT EXIST!\n"+
			"pathFileName='%v'\n",
			ePrefix.String(),
			pathFileName)

		return err
	}

	err = filePermission.IsValid(
		ePrefix.XCpy("filePermission"))

	if err != nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'filePermission' is INVALID!\n"+
			"Error='%v'\n",
			ePrefix.String(),
			err.Error())

		return err
	}

	var newOsFileMode os.FileMode

	newOsFileMode, err = filePermission.GetFileMode(
		ePrefix.XCpy("newOsFileMode<-filePermission"))

	if err != nil {
		return err
	}

	var err2 error

	err2 = os.Chmod(pathFileName, newOsFileMode)

	if err2 != nil {

		var changeModeTxt, changeModeValue string

		changeModeTxt,
			_ = filePermission.GetPermissionTextCode(
			ePrefix.XCpy("changeModeTxt<-filePermission"))

		changeModeValue,
			_ =
			filePermission.GetPermissionFileModeValueText(
				ePrefix.XCpy("changeModeValue<-filePermission"))

		err = fmt.Errorf("%v\n"+
			"Error returned by os.Chmod(pathFileName, newOsFileMode).\n"+
			"pathFileName='%v'\n"+
			"newOsFileMode Text='%v'\n"+
			"newOsFileModeValue='%v'\n"+
			"Error='%v'",
			ePrefix.String(),
			pathFileName,
			changeModeTxt,
			changeModeValue,
			err.Error())
	}

	return err
}

// cleanFileNameExtStr
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
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (fHelperNanobot *fileHelperNanobot) cleanFileNameExtStr(
	fileNameExtStr string,
	errPrefDto *ePref.ErrPrefixDto) (
	returnedFileNameExt string,
	isEmpty bool,
	err error) {

	if fHelperNanobot.lock == nil {
		fHelperNanobot.lock = new(sync.Mutex)
	}

	fHelperNanobot.lock.Lock()

	defer fHelperNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	isEmpty = true

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileHelperNanobot."+
			"cleanFileNameExtStr()",
		"")

	if err != nil {
		return returnedFileNameExt, isEmpty, err
	}

	var pathDoesExist bool
	var fInfo FileInfoPlus
	var adjustedFileNameExt string

	fHelperMolecule := new(fileHelperMolecule)

	adjustedFileNameExt,
		pathDoesExist,
		fInfo,
		err = fHelperMolecule.
		doesPathFileExist(
			fileNameExtStr,
			PreProcPathCode.PathSeparator(), // Convert to os Path Separators
			ePrefix,
			"fileNameExtStr")

	if err != nil {

		returnedFileNameExt = ""

		isEmpty = true

		return returnedFileNameExt, isEmpty, err
	}

	osPathSepStr := string(os.PathSeparator)

	if strings.Contains(adjustedFileNameExt, osPathSepStr+osPathSepStr) {

		returnedFileNameExt = ""

		isEmpty = true

		err = fmt.Errorf("%v\n"+
			"Error: Invalid Directory string.\n"+
			"Directory string contains invalid Path Separators.\n"+
			"adjustedFileNameExt='%v'\n",
			ePrefix.String(),
			adjustedFileNameExt)

		return returnedFileNameExt, isEmpty, err
	}

	if strings.Contains(adjustedFileNameExt, "...") {

		err = fmt.Errorf("%v\n"+
			"Error: Invalid Directory string. Contains invalid dots.\n"+
			"adjustedFileNameExt='%v'\n",
			ePrefix.String(),
			adjustedFileNameExt)

		return returnedFileNameExt, isEmpty, err
	}

	// Find out if the file name extension path
	// actually exists.

	if pathDoesExist {
		// The path exists

		if fInfo.IsDir() {
			// The path exists and it is a directory.
			// There is no File Name present.
			returnedFileNameExt = ""
			isEmpty = true
			err = nil
			return

		} else {
			// The path exists, and it is a valid
			// file name.
			returnedFileNameExt = fInfo.Name()
			isEmpty = false

			err = nil
			return
		}
	} // End of if pathDoesExist

	firstCharIdx,
		lastCharIdx,
		err := fHelperMolecule.
		getFirstLastNonSeparatorCharIndexInPathStr(
			adjustedFileNameExt,
			ePrefix)

	if firstCharIdx == -1 || lastCharIdx == -1 {

		err = fmt.Errorf("%v\n"+
			"File Name Extension string contains no "+
			"valid file name characters!\n"+
			"adjustedFileNameExt='%v'\n",
			ePrefix.String(),
			adjustedFileNameExt)

		return returnedFileNameExt, isEmpty, err
	}

	// The file name extension path does not exist

	interiorDotPathIdx := strings.LastIndex(adjustedFileNameExt, "."+string(os.PathSeparator))

	if interiorDotPathIdx > firstCharIdx {

		err = fmt.Errorf("%v\n"+
			"Error: INVALID PATH.\n"+
			"Invalid interior relative path detected!\n"+
			"adjustedFileNameExt='%v'\n",
			ePrefix,
			adjustedFileNameExt)

		return returnedFileNameExt, isEmpty, err
	}

	var slashIdxs []int
	var err2 error

	fHelperAtom := new(fileHelperAtom)

	slashIdxs,
		err2 = fHelperAtom.
		getPathSeparatorIndexesInPathStr(
			adjustedFileNameExt,
			ePrefix)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned from fh.GetPathSeparatorIndexesInPathStr"+
			"(adjustedFileNameExt).\n"+
			"adustedFileNameExt='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			adjustedFileNameExt,
			err2.Error())

		return returnedFileNameExt, isEmpty, err
	}

	lSlashIdxs := len(slashIdxs)

	if lSlashIdxs == 0 {

		returnedFileNameExt = adjustedFileNameExt

		isEmpty = false

		err = nil

		return returnedFileNameExt, isEmpty, err
	}

	var dotIdxs []int

	dotIdxs,
		err2 = fHelperAtom.
		getDotSeparatorIndexesInPathStr(
			adjustedFileNameExt,
			ePrefix)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by GetDotSeparatorIndexesInPathStr(adjustedDirName).\n"+
			"adjustedFileNameExt='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			adjustedFileNameExt,
			err2.Error())

		returnedFileNameExt = ""

		isEmpty = true

		return returnedFileNameExt, isEmpty, err
	}

	lDotIdxs := len(dotIdxs)

	// Option # 1
	if lDotIdxs == 0 && lSlashIdxs == 0 && lastCharIdx == -1 {
		// There are no valid characters in the string
		returnedFileNameExt = ""

	} else if lDotIdxs == 0 && lSlashIdxs == 0 && lastCharIdx > -1 {
		// Option # 2
		// String consists only of eligible alphanumeric characters
		// "sometextstring"
		returnedFileNameExt = adjustedFileNameExt

	} else if lDotIdxs == 0 && lSlashIdxs > 0 && lastCharIdx == -1 {
		// Option # 3
		// There no dots no characters but string does contain
		// slashes
		returnedFileNameExt = ""

	} else if lDotIdxs == 0 && lSlashIdxs > 0 && lastCharIdx > -1 {
		// Option # 4
		// strings contains slashes and characters but no dots.

		if lastCharIdx < slashIdxs[lSlashIdxs-1] {
			// Example: ../dir1/dir2/
			returnedFileNameExt = ""
		} else {
			// Must be  lastCharIdx > slashIdxs[lSlashIdxs-1]
			// Example: ../dir1/dir2/filename
			returnedFileNameExt = adjustedFileNameExt[slashIdxs[lSlashIdxs-1]+1:]
		}

	} else if lDotIdxs > 0 && lSlashIdxs == 0 && lastCharIdx == -1 {
		// Option # 5
		// dots only. Must be "." or ".."
		// The is no file name
		returnedFileNameExt = ""

	} else if lDotIdxs > 0 && lSlashIdxs == 0 && lastCharIdx > -1 {
		// Option # 6
		// Dots and characters only. No slashes.
		// Maybe 'fileName.ext'
		returnedFileNameExt = adjustedFileNameExt

	} else if lDotIdxs > 0 && lSlashIdxs > 0 && lastCharIdx == -1 {
		// Option # 7
		// Dots and slashes, but no characters.
		returnedFileNameExt = ""

	} else {
		// Option # 8
		// MUST BE lDotIdxs > 0 && lSlashIdxs > 0 && lastCharIdx > -1
		// We have dots, slashes and characters

		if lastCharIdx > slashIdxs[lSlashIdxs-1] {
			// The last char comes after the last slash.
			// Take everything after the last slash
			returnedFileNameExt = adjustedFileNameExt[slashIdxs[lSlashIdxs-1]+1:]
		} else {
			// The last character comes before the last
			// slash. There is no file name here.
			returnedFileNameExt = ""
		}
	}

	if len(returnedFileNameExt) == 0 {

		isEmpty = true

	} else {

		isEmpty = false

	}

	err = nil

	return returnedFileNameExt, isEmpty, err
}

// changeFileTimes
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
func (fHelperNanobot *fileHelperNanobot) changeFileTimes(
	pathFileName string,
	newAccessTime,
	newModTime time.Time,
	errorPrefix interface{}) error {

	if fHelperNanobot.lock == nil {
		fHelperNanobot.lock = new(sync.Mutex)
	}

	fHelperNanobot.lock.Lock()

	defer fHelperNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"fileHelperNanobot."+
			"changeFileTimes()",
		"")

	if err != nil {
		return err
	}

	var filePathDoesExist bool
	var fInfo FileInfoPlus

	pathFileName,
		filePathDoesExist,
		fInfo,
		err = new(fileHelperMolecule).doesPathFileExist(pathFileName,
		PreProcPathCode.AbsolutePath(), // Convert to Absolute Path
		ePrefix,
		"pathFileName")

	if err != nil {
		return err
	}

	if !filePathDoesExist {

		err = fmt.Errorf("%v\n"+
			"ERROR: 'pathFileName' DOES NOT EXIST!\n"+
			"pathFileName='%v'\n",
			ePrefix.String(),
			pathFileName)

		return err
	}

	if fInfo.IsDir() {

		err = fmt.Errorf("%v\n"+
			"ERROR: 'pathFileName' is a directory path and NOT a file!\n"+
			"pathFileName='%v'\n",
			ePrefix.String(),
			pathFileName)

		return err
	}

	tDefault := time.Time{}

	if newAccessTime == tDefault {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'newAccessTime' is INVALID!\n"+
			"newAccessTime='%v'\n",
			ePrefix.String(),
			newAccessTime)

		return err
	}

	if newModTime == tDefault {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'newModTime' is INVALID!\n"+
			"newModTime='%v'\n",
			ePrefix.String(),
			newModTime)

		return err
	}

	var err2 error

	err2 = os.Chtimes(pathFileName, newAccessTime, newModTime)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"ERROR returned by os.Chtimes(pathFileName,newAccessTime, newModTime)\n"+
			"newAccessTime='%v'\n"+
			"newModTime='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			newAccessTime.Format("2006-01-02 15:04:05.000000000 -0700 MST"),
			newModTime.Format("2006-01-02 15:04:05.000000000 -0700 MST"),
			err.Error())
	}

	return err
}

// cleanDirStr
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
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (fHelperNanobot *fileHelperNanobot) cleanDirStr(
	dirNameStr string,
	errPrefDto *ePref.ErrPrefixDto) (
	returnedDirName string,
	isEmpty bool,
	err error) {

	if fHelperNanobot.lock == nil {
		fHelperNanobot.lock = new(sync.Mutex)
	}

	fHelperNanobot.lock.Lock()

	defer fHelperNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	returnedDirName = ""

	isEmpty = true

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileHelperNanobot."+
			"cleanDirStr()",
		"")

	if err != nil {
		return returnedDirName, isEmpty, err
	}

	var pathDoesExist bool
	var fInfo FileInfoPlus
	var adjustedDirName string

	fHelperMolecule := new(fileHelperMolecule)

	adjustedDirName,
		pathDoesExist,
		fInfo,
		err = fHelperMolecule.
		doesPathFileExist(dirNameStr,
			PreProcPathCode.PathSeparator(), // Convert to os Path Separators
			ePrefix,
			"dirNameStr")

	if err != nil {
		returnedDirName = ""
		isEmpty = true
		return returnedDirName, isEmpty, err
	}

	lAdjustedDirName := len(adjustedDirName)

	osPathSepStr := string(os.PathSeparator)

	if strings.Contains(adjustedDirName, osPathSepStr+osPathSepStr) {

		returnedDirName = ""

		isEmpty = true

		err = fmt.Errorf("%v\n"+
			"Error: Invalid Directory string.\n"+
			"Directory string contains invalid Path Separators.\n"+
			"adjustedDirName='%v'\n",
			ePrefix.String(),
			adjustedDirName)

		return returnedDirName, isEmpty, err
	}

	if strings.Contains(adjustedDirName, "...") {

		returnedDirName = ""

		isEmpty = true

		err = fmt.Errorf("%v\n"+
			"Error: Invalid Directory string.\n"+
			"Directory string contains invalid dots.\n"+
			"adjustedDirName='%v'\n",
			ePrefix.String(),
			adjustedDirName)

		return returnedDirName, isEmpty, err
	}

	absPath, err := new(fileHelperProton).
		makeAbsolutePath(adjustedDirName,
			ePrefix)

	if err != nil {

		err = fmt.Errorf("%v\n"+
			"Error occurred while converting path\n"+
			"to absolute path!\n"+
			"dirPath='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			adjustedDirName,
			err.Error())

		returnedDirName = ""

		isEmpty = true

		return returnedDirName, isEmpty, err
	}

	volName := fp.VolumeName(absPath)

	volumeIdx := strings.Index(adjustedDirName, volName)

	if strings.ToLower(volName) == strings.ToLower(adjustedDirName) {

		returnedDirName = adjustedDirName

		if len(returnedDirName) == 0 {
			isEmpty = true
		} else {
			isEmpty = false
		}

		err = nil
		return returnedDirName, isEmpty, err
	}

	isAbsPath := false

	if strings.ToLower(absPath) == strings.ToLower(adjustedDirName) {
		isAbsPath = true
	}

	dotPlusOsPathSepStr := "." + osPathSepStr

	if pathDoesExist {
		// The path exists

		if fInfo.IsDir() {

			// The path exists and it is a directory
			returnedDirName = adjustedDirName

		} else {
			// The path exists, but it is
			// a File Name and NOT a directory name.
			adjustedDirName = adjustedDirName[0 : lAdjustedDirName-len(fInfo.Name())]
			lAdjustedDirName = len(adjustedDirName)

			if lAdjustedDirName < 1 {
				returnedDirName = ""
			} else {
				returnedDirName = adjustedDirName
			}

		}

		lAdjustedDirName = len(returnedDirName)

		if lAdjustedDirName == 0 {
			isEmpty = true
			err = nil
			return returnedDirName, isEmpty, err
		}

		if returnedDirName[lAdjustedDirName-1] == os.PathSeparator {
			// Last character in path is a Path Separator

			if lAdjustedDirName >= 2 &&
				returnedDirName[lAdjustedDirName-2] == '.' {
				// The char immediately preceding the Path Separator
				// in the last char position is a dot.
				// Keep the path separator
				returnedDirName = adjustedDirName
			} else {
				// The char immediately preceding the Path Separator in
				// the last char position is NOT a dot. Delete the
				// trailing path separator
				returnedDirName = returnedDirName[0 : lAdjustedDirName-1]
			}
		}

		lAdjustedDirName = len(returnedDirName)

		if returnedDirName[0] != '.' &&
			returnedDirName[0] != os.PathSeparator &&
			volumeIdx == -1 &&
			!isAbsPath {

			// First letter in path is a character
			// Add a leading dot
			returnedDirName = dotPlusOsPathSepStr + returnedDirName

		} else if returnedDirName[0] == os.PathSeparator &&
			volumeIdx == -1 &&
			!isAbsPath {

			returnedDirName = "." + returnedDirName
		}

		if len(returnedDirName) == 0 {
			isEmpty = true
		} else {
			isEmpty = false
		}

		err = nil
		return returnedDirName, isEmpty, err

	} // End of pathDoesExist == true

	// The Path DOES NOT EXIST ON DISK
	var firstCharIdx, lastCharIdx int
	var err2 error

	firstCharIdx,
		lastCharIdx,
		err2 = fHelperMolecule.
		getFirstLastNonSeparatorCharIndexInPathStr(
			adjustedDirName,
			ePrefix)

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"Error returned by "+
			"GetFirstLastNonSeparatorCharIndexInPathStr(adjustedDirName).\n"+
			"adjustedDirName='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			adjustedDirName,
			err2.Error())

		returnedDirName = ""

		isEmpty = true

		return returnedDirName, isEmpty, err
	}

	interiorDotPathIdx := strings.LastIndex(adjustedDirName, "."+string(os.PathSeparator))

	if interiorDotPathIdx > firstCharIdx {

		err = fmt.Errorf("%v\n"+
			"Error: INVALID PATH. Invalid interior relative path detected!\n"+
			"adjustedDirName='%v'\n",
			ePrefix.String(),
			adjustedDirName)

		returnedDirName = ""

		isEmpty = true

		return returnedDirName, isEmpty, err
	}

	var slashIdxs []int

	fHelperAtom := new(fileHelperAtom)

	slashIdxs, err2 = fHelperAtom.
		getPathSeparatorIndexesInPathStr(
			adjustedDirName,
			ePrefix)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by fh."+
			"GetPathSeparatorIndexesInPathStr(adjustedDirName).\n"+
			"adjusteDirName='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			adjustedDirName,
			err2.Error())

		returnedDirName = ""

		isEmpty = true

		return returnedDirName, isEmpty, err
	}

	lSlashIdxs := len(slashIdxs)

	var dotIdxs []int

	dotIdxs,
		err2 = fHelperAtom.getDotSeparatorIndexesInPathStr(
		adjustedDirName,
		ePrefix)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by fh."+
			"GetDotSeparatorIndexesInPathStr(adjustedDirName).\n"+
			"adjustedDirName='%v'\nError='%v'\n",
			ePrefix.String(),
			adjustedDirName,
			err2.Error())

		returnedDirName = ""

		isEmpty = true

		return returnedDirName, isEmpty, err
	}

	lDotIdxs := len(dotIdxs)

	if lDotIdxs == 0 && lSlashIdxs == 0 && lastCharIdx == -1 {
		// Option # 1
		// There are no valid characters in the string
		returnedDirName = ""

	} else if lDotIdxs == 0 && lSlashIdxs == 0 && lastCharIdx > -1 {
		// Option # 2
		// String consists only of eligible alphanumeric characters
		// "sometextstring"
		returnedDirName = dotPlusOsPathSepStr + adjustedDirName

	} else if lDotIdxs == 0 && lSlashIdxs > 0 && lastCharIdx == -1 {
		// Option # 3
		// There no dots no characters but string does contain
		// slashes

		if lSlashIdxs > 1 {
			returnedDirName = ""
		} else {
			// lSlashIdxs must be '1'
			returnedDirName = dotPlusOsPathSepStr
		}

	} else if lDotIdxs == 0 && lSlashIdxs > 0 && lastCharIdx > -1 {
		// Option # 4
		// strings contains slashes and characters but no dots.

		returnedDirName = adjustedDirName

	} else if lDotIdxs > 0 && lSlashIdxs == 0 && lastCharIdx == -1 {
		// Option # 5
		// dots only. Must be "." or ".."
		// Add trailing path separator
		returnedDirName = adjustedDirName + osPathSepStr

	} else if lDotIdxs > 0 && lSlashIdxs == 0 && lastCharIdx > -1 {
		// Option # 6
		// Dots and characters only. No slashes.
		// Maybe 'fileName.ext'
		returnedDirName = ""

	} else if lDotIdxs > 0 && lSlashIdxs > 0 && lastCharIdx == -1 {
		// Option # 7
		// Dots and slashes, but no characters.
		returnedDirName = adjustedDirName

	} else {
		// Option # 8
		// MUST BE lDotIdxs > 0 && lSlashIdxs > 0 && lastCharIdx > -1
		// Has dots, slashes and characters

		returnedDirName = adjustedDirName

		// If there is a dot after the last path separator
		// and there is a character following the dot.
		// Therefore, this is a filename extension, NOT a directory
		if dotIdxs[lDotIdxs-1] > slashIdxs[lSlashIdxs-1] &&
			dotIdxs[lDotIdxs-1] < lastCharIdx {

			returnedDirName = adjustedDirName[0:slashIdxs[lSlashIdxs-1]]
		}
	}

	lAdjustedDirName = len(returnedDirName)

	if lAdjustedDirName == 0 {
		isEmpty = true
		err = nil
		return returnedDirName, isEmpty, err
	}

	if returnedDirName[lAdjustedDirName-1] == os.PathSeparator {
		// Last character in path is a Path Separator
		if lAdjustedDirName >= 2 &&
			returnedDirName[lAdjustedDirName-2] == '.' {
			// The char immediately preceding the Path Separator
			// is a dot. Keep the path separator
			// do nothing
			lAdjustedDirName = len(returnedDirName)
		} else {
			returnedDirName = returnedDirName[0 : lAdjustedDirName-1]
		}
	}

	lAdjustedDirName = len(returnedDirName)

	if returnedDirName[0] != '.' &&
		returnedDirName[0] != os.PathSeparator &&
		volumeIdx == -1 &&
		!isAbsPath {

		// First letter in path is a character
		// Add a leading dot
		returnedDirName = dotPlusOsPathSepStr + returnedDirName

	} else if returnedDirName[0] == os.PathSeparator &&
		volumeIdx == -1 &&
		!isAbsPath {

		returnedDirName = "." + returnedDirName
	}

	if len(returnedDirName) == 0 {
		isEmpty = true
	} else {
		isEmpty = false
	}

	err = nil

	return returnedDirName, isEmpty, err
}

// doesFileExist - Returns a boolean value designating whether the passed
// file name exists.
//
// This method does not differentiate between Path Errors and Non-Path
// Errors returned by os.Stat(). The method only returns a boolean
// value.
//
// If a Non-Path Error is returned by os.Stat(), this method will
// classify the file as "Does NOT Exist" and return a value of
// false.
//
// For a more granular test of whether a file exists, see method
// FileHelper.DoesThisFileExist().
func (fHelperNanobot *fileHelperNanobot) doesFileExist(
	pathFileName string) bool {

	_,
		pathFileDoesExist,
		_,
		nonPathError :=
		new(fileHelperMolecule).
			doesPathFileExist(
				pathFileName,
				PreProcPathCode.AbsolutePath(), // Convert to Absolute Path
				"fileHelperNanobot.doesFileExist()",
				"pathFileName")

	if !pathFileDoesExist || nonPathError != nil {
		return false
	}

	return true
}

// deleteDirFile
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
func (fHelperNanobot *fileHelperNanobot) deleteDirFile(
	pathFile string,
	errPrefDto *ePref.ErrPrefixDto) (
	msgError error,
	lowLevelErr error) {

	if fHelperNanobot.lock == nil {
		fHelperNanobot.lock = new(sync.Mutex)
	}

	fHelperNanobot.lock.Lock()

	defer fHelperNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		msgError = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileHelperNanobot."+
			"deleteDirFile()",
		"")

	if msgError != nil {
		return msgError, lowLevelErr
	}

	var fileDoesExist bool

	fHelpMolecule := fileHelperMolecule{}

	pathFile,
		fileDoesExist,
		_,
		msgError = fHelpMolecule.doesPathFileExist(
		pathFile,
		PreProcPathCode.AbsolutePath(), // Convert to Absolute File Path
		ePrefix,
		"pathFile")

	if msgError != nil {
		return msgError, lowLevelErr
	}

	if !fileDoesExist {
		// Doesn't exist. Nothing to do.
		return msgError, lowLevelErr
	}

	lowLevelErr = os.Remove(pathFile)

	if lowLevelErr != nil {
		msgError = fmt.Errorf("%v\n"+
			"Error returned from os.Remove(pathFile).\n"+
			"pathFile='%v'\nError= \n%v\n",
			ePrefix.String(),
			pathFile,
			lowLevelErr.Error())

		return msgError, lowLevelErr
	}

	var err2 error
	_,
		fileDoesExist,
		_,
		err2 = fHelpMolecule.doesPathFileExist(
		pathFile,
		PreProcPathCode.None(), // Apply No Pre-Processing. Take no action
		ePrefix,
		"pathFile")

	if err2 != nil {

		msgError = fmt.Errorf(
			"After attempted deletion, file error occurred!\n"+
				"pathFile='%v'\n"+
				"Error = \n%v\n",
			pathFile, err2.Error())

		return msgError, lowLevelErr
	}

	if fileDoesExist {

		// File STILL Exists! ERROR!
		msgError = fmt.Errorf("%v\n"+
			"ERROR: After attempted deletion, file still exists!\n"+
			"pathFile='%v'\n",
			ePrefix.String(),
			pathFile)
	}

	return msgError, lowLevelErr
}
