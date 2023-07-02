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
	errPrefDto *ePref.ErrPrefixDto) (
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
		errPrefDto,
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
	errPrefDto *ePref.ErrPrefixDto) error {

	if fHelperNanobot.lock == nil {
		fHelperNanobot.lock = new(sync.Mutex)
	}

	fHelperNanobot.lock.Lock()

	defer fHelperNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errPrefDto,
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

	err = filePermission.IsValidInstanceError(
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
	errPrefDto *ePref.ErrPrefixDto) error {

	if fHelperNanobot.lock == nil {
		fHelperNanobot.lock = new(sync.Mutex)
	}

	fHelperNanobot.lock.Lock()

	defer fHelperNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errPrefDto,
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

// doesFileExist
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
func (fHelperNanobot *fileHelperNanobot) doesFileExist(
	pathFileName string) bool {

	if fHelperNanobot.lock == nil {
		fHelperNanobot.lock = new(sync.Mutex)
	}

	fHelperNanobot.lock.Lock()

	defer fHelperNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"fileHelperNanobot."+
			"doesFileExist()",
		"")

	if err != nil {
		return false
	}

	_,
		pathFileDoesExist,
		_,
		nonPathError :=
		new(fileHelperMolecule).
			doesPathFileExist(
				pathFileName,
				PreProcPathCode.AbsolutePath(), // Convert to Absolute Path
				ePrefix,
				"pathFileName")

	if !pathFileDoesExist || nonPathError != nil {
		return false
	}

	return true
}

// doesFileInfoExist
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
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (fHelperNanobot *fileHelperNanobot) doesFileInfoExist(
	pathFileName string,
	errPrefDto *ePref.ErrPrefixDto) (
	doesFInfoExist bool,
	fInfo os.FileInfo,
	err error) {

	if fHelperNanobot.lock == nil {
		fHelperNanobot.lock = new(sync.Mutex)
	}

	fHelperNanobot.lock.Lock()

	defer fHelperNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	doesFInfoExist = false

	fInfo = nil

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileHelperNanobot."+
			"doesFileInfoExist()",
		"")

	if err != nil {
		return doesFInfoExist, fInfo, err
	}

	fInfoPlus := FileInfoPlus{}

	pathFileName,
		doesFInfoExist,
		fInfoPlus,
		err = new(fileHelperMolecule).
		doesPathFileExist(
			pathFileName,
			PreProcPathCode.AbsolutePath(), // Convert to Absolute Path
			ePrefix,
			"pathFileName")

	if err != nil {

		doesFInfoExist = false

		return doesFInfoExist, fInfo, err
	}

	if !doesFInfoExist {

		err = nil

		fInfo = nil

		return doesFInfoExist, fInfo, err
	}

	fInfo = fInfoPlus.GetOriginalFileInfo()

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
//	 	text passed by input parameter, 'errPrefDto'.
//	 	The 'errPrefDto' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fHelperNanobot *fileHelperNanobot) doesFileInfoPlusExist(
	pathFileName string,
	errPrefDto *ePref.ErrPrefixDto) (
	doesFInfoExist bool,
	fInfoPlus FileInfoPlus,
	err error) {

	if fHelperNanobot.lock == nil {
		fHelperNanobot.lock = new(sync.Mutex)
	}

	fHelperNanobot.lock.Lock()

	defer fHelperNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	doesFInfoExist = false

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileHelperNanobot."+
			"doesFileInfoPlusExist()",
		"")

	if err != nil {
		return doesFInfoExist, fInfoPlus, err
	}

	pathFileName,
		doesFInfoExist,
		fInfoPlus,
		err = new(fileHelperMolecule).
		doesPathFileExist(
			pathFileName,
			PreProcPathCode.AbsolutePath(), // Convert to Absolute Path
			ePrefix,
			"pathFileName")

	if err != nil {

		doesFInfoExist = false

		return doesFInfoExist, fInfoPlus, err
	}

	return doesFInfoExist, fInfoPlus, err
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

// deleteDirPathAll
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
func (fHelperNanobot *fileHelperNanobot) deleteDirPathAll(
	pathDir string,
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
			"deleteDirPathAll()",
		"")

	if msgError != nil {
		return msgError, lowLevelErr
	}

	var pathFileDoesExist bool

	fHelpMolecule := fileHelperMolecule{}

	pathDir,
		pathFileDoesExist,
		_,
		msgError = fHelpMolecule.doesPathFileExist(
		pathDir,
		PreProcPathCode.AbsolutePath(), // Convert To Absolute Path
		ePrefix,
		"pathDir")

	if msgError != nil {
		return msgError, lowLevelErr
	}

	if !pathFileDoesExist {
		// Doesn't exist. Nothing to do.
		return msgError, lowLevelErr
	}

	// Make 3-attempts to delete the
	// directory path.
	for i := 0; i < 3; i++ {

		lowLevelErr = nil

		msgError = nil

		lowLevelErr = os.RemoveAll(pathDir)

		if lowLevelErr != nil {

			msgError = fmt.Errorf("%v\n"+
				"Error returned by os.RemoveAll(pathDir).\n"+
				"pathDir='%v'\nError=\n%v\n",
				ePrefix.String(),
				pathDir,
				lowLevelErr.Error())
		}

		if lowLevelErr == nil {
			break
		}

		time.Sleep(50 * time.Millisecond)
	}

	if lowLevelErr != nil {
		return msgError, lowLevelErr
	}

	pathDir,
		pathFileDoesExist,
		_,
		msgError = fHelpMolecule.doesPathFileExist(
		pathDir,
		PreProcPathCode.None(), // Apply No Pre-Processing. Take No Action.
		ePrefix,
		"pathDir")

	if msgError != nil {
		return msgError, lowLevelErr
	}

	if pathFileDoesExist {

		// Path still exists. Something is wrong.
		msgError = fmt.Errorf("%v\n"+
			"Delete Failed! 'pathDir' still exists!\n"+
			"pathDir='%v'\n",
			ePrefix.String(),
			pathDir)
	}

	return msgError, lowLevelErr
}

// findFilesInPath
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
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (fHelperNanobot *fileHelperNanobot) findFilesInPath(
	pathName string,
	fileSearchPattern string,
	errPrefDto *ePref.ErrPrefixDto) (
	[]string,
	error) {

	if fHelperNanobot.lock == nil {
		fHelperNanobot.lock = new(sync.Mutex)
	}

	fHelperNanobot.lock.Lock()

	defer fHelperNanobot.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileHelperNanobot."+
			"FindFilesInPath()",
		"")

	if err != nil {
		return []string{}, err
	}

	var fInfo FileInfoPlus
	var pathDoesExist bool
	var errCode int

	pathName,
		pathDoesExist,
		fInfo,
		err = new(fileHelperMolecule).doesPathFileExist(
		pathName,
		PreProcPathCode.AbsolutePath(), // Convert to Absolute Path
		ePrefix,
		"pathName")

	if err != nil {
		return []string{}, err
	}

	errCode, _, fileSearchPattern =
		new(fileHelperElectron).isStringEmptyOrBlank(fileSearchPattern)

	if errCode == -1 {

		return []string{},
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'fileSearchPattern' is "+
				"an empty string!\n",
				ePrefix.String())
	}

	if errCode == -2 {

		return []string{},
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'fileSearchPattern' consists "+
				"of blank spaces!\n",
				ePrefix.String())
	}

	if !pathDoesExist {
		return []string{},
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'pathName' DOES NOT EXIST!\n"+
				"pathName='%v'\n",
				ePrefix.String(),
				pathName)
	}

	if !fInfo.IsDir() {
		return []string{},
			fmt.Errorf("%v\n"+
				"Error: The path exists, but it NOT a directory!\n"+
				"pathName='%v'\n",
				ePrefix.String(),
				pathName)
	}

	// fInfo is a Directory.

	searchStr := new(fileHelperMolecule).
		joinPathsAdjustSeparators(pathName, fileSearchPattern)

	var results []string
	var err2 error

	results,
		err2 = fp.Glob(searchStr)

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"Error returned by fp.Glob(searchStr).\n"+
			"searchStr='%v'\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			searchStr,
			err2.Error())
	}

	return results, err
}

// getFileExtension
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
//	 	text passed by input parameter, 'errPrefDto'.
//	 	The 'errPrefDto' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fHelperNanobot *fileHelperNanobot) getFileExtension(
	pathFileNameExt string,
	errPrefDto *ePref.ErrPrefixDto) (
	ext string,
	isEmpty bool,
	err error) {

	if fHelperNanobot.lock == nil {
		fHelperNanobot.lock = new(sync.Mutex)
	}

	fHelperNanobot.lock.Lock()

	defer fHelperNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ext = ""

	isEmpty = true

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileHelperNanobot."+
			"getFileExtension()",
		"")

	if err != nil {
		return ext, isEmpty, err
	}

	fHelperElectron := new(fileHelperElectron)

	errCode := 0

	errCode,
		_,
		pathFileNameExt = fHelperElectron.
		isStringEmptyOrBlank(pathFileNameExt)

	if errCode == -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'pathFileNameExt' is an empty string!\n",
			ePrefix.String())

		return ext, isEmpty, err
	}

	if errCode == -2 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'pathFileNameExt' consists of blank spaces!\n",
			ePrefix.String())

		return ext, isEmpty, err
	}

	fHelperAtom := fileHelperAtom{}

	testPathFileNameExt := fHelperAtom.
		adjustPathSlash(pathFileNameExt)

	errCode,
		_,
		testPathFileNameExt =
		fHelperElectron.isStringEmptyOrBlank(testPathFileNameExt)

	if errCode < 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Cleaned version of 'pathFileNameExt', 'testPathFileNameExt'\n"+
			"is an empty string!\n",
			ePrefix.String())

		return ext, isEmpty, err
	}

	var dotIdxs []int
	var err2 error

	dotIdxs,
		err2 = fHelperAtom.
		getDotSeparatorIndexesInPathStr(
			testPathFileNameExt,
			ePrefix.XCpy(
				"dotIdxs<-testPathFileNameExt"))

	if err2 != nil {

		ext = ""

		isEmpty = true

		err = fmt.Errorf("%v\n"+
			"Error returned from GetDotSeparatorIndexesInPathStr(testPathFileNameExt).\n"+
			"testPathFileNameExt='%v'\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			testPathFileNameExt,
			err2.Error())

		return ext, isEmpty, err
	}

	lenDotIdxs := len(dotIdxs)

	// Deal with case where the pathFileNameExt contains
	// no dots.
	if lenDotIdxs == 0 {
		ext = ""
		isEmpty = true
		err = nil
		return ext, isEmpty, err

	}

	firstGoodCharIdx,
		lastGoodCharIdx,
		err2 := new(fileHelperMolecule).
		getFirstLastNonSeparatorCharIndexInPathStr(
			testPathFileNameExt,
			ePrefix)

	if err2 != nil {

		ext = ""

		isEmpty = true

		err = fmt.Errorf("%v\n"+
			"Error returned from GetFirstLastNonSeparatorCharIndexInPathStr(testPathFileNameExt).\n"+
			"testPathFileNameExt='%v'\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			testPathFileNameExt,
			err2.Error())

		return ext, isEmpty, err
	}

	// Deal with the case where pathFileNameExt contains no
	// valid alphanumeric characters
	if firstGoodCharIdx == -1 || lastGoodCharIdx == -1 {
		ext = ""
		isEmpty = true
		err = nil
		return ext, isEmpty, err
	}

	var slashIdxs []int

	slashIdxs,
		err2 = fHelperAtom.
		getPathSeparatorIndexesInPathStr(
			testPathFileNameExt,
			ePrefix)

	if err2 != nil {

		ext = ""

		isEmpty = true

		err = fmt.Errorf("%v\n"+
			"Error returned from GetPathSeparatorIndexesInPathStr(testPathFileNameExt).\n"+
			"testPathFileNameExt='%v'\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			testPathFileNameExt,
			err2.Error())

		return ext, isEmpty, err
	}

	lenSlashIdxs := len(slashIdxs)

	if lenSlashIdxs == 0 &&
		lenDotIdxs == 1 &&
		dotIdxs[lenDotIdxs-1] == 0 {
		// deal with the case .gitignore
		ext = ""
		isEmpty = true
		err = nil
		return ext, isEmpty, err
	}

	if lenSlashIdxs == 0 {
		ext = testPathFileNameExt[dotIdxs[lenDotIdxs-1]:]
		isEmpty = false
		err = nil
		return ext, isEmpty, err
	}

	// lenDotIdxs and lenSlasIdxs both greater than zero
	if dotIdxs[lenDotIdxs-1] > slashIdxs[lenSlashIdxs-1] &&
		dotIdxs[lenDotIdxs-1] < lastGoodCharIdx {

		ext = testPathFileNameExt[dotIdxs[lenDotIdxs-1]:]
		isEmpty = false
		err = nil
		return ext, isEmpty, err

	}

	ext = ""

	isEmpty = true

	err = nil

	return ext, isEmpty, err
}

// getFileInfo
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
//	 	text passed by input parameter, 'errPrefDto'.
//	 	The 'errPrefDto' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fHelperNanobot *fileHelperNanobot) getFileInfo(
	pathFileName string,
	errPrefDto *ePref.ErrPrefixDto) (
	osFileInfo os.FileInfo,
	err error) {

	if fHelperNanobot.lock == nil {
		fHelperNanobot.lock = new(sync.Mutex)
	}

	fHelperNanobot.lock.Lock()

	defer fHelperNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileHelperNanobot."+
			"getFileInfo()",
		"")

	if err != nil {

		return osFileInfo, err
	}

	var pathDoesExist bool
	var fInfoPlus FileInfoPlus

	pathFileName,
		pathDoesExist,
		fInfoPlus,
		err = new(fileHelperMolecule).doesPathFileExist(
		pathFileName,
		PreProcPathCode.AbsolutePath(), // Convert to Absolute Path
		ePrefix,
		"pathFileName")

	if err != nil {
		return osFileInfo, err
	}

	if !pathDoesExist {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'pathFileName' does NOT exist!\n"+
			"pathFileName='%v'\n",
			ePrefix.String(),
			pathFileName)

		return osFileInfo, err
	}

	osFileInfo = fInfoPlus.GetOriginalFileInfo()

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
//	 	text passed by input parameter, 'errPrefDto'.
//	 	The 'errPrefDto' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fHelperNanobot *fileHelperNanobot) getFileInfoPlus(
	pathFileName string,
	errPrefDto *ePref.ErrPrefixDto) (
	fileInfoPlus FileInfoPlus,
	err error) {

	if fHelperNanobot.lock == nil {
		fHelperNanobot.lock = new(sync.Mutex)
	}

	fHelperNanobot.lock.Lock()

	defer fHelperNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileHelperNanobot."+
			"getFileInfoPlus()",
		"")

	if err != nil {

		return fileInfoPlus, err
	}

	var pathDoesExist bool

	pathFileName,
		pathDoesExist,
		fileInfoPlus,
		err = new(fileHelperMolecule).doesPathFileExist(
		pathFileName,
		PreProcPathCode.AbsolutePath(), // Convert to Absolute Path
		ePrefix,
		"pathFileName")

	if err != nil {
		return fileInfoPlus, err
	}

	if !pathDoesExist {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'pathFileName' does NOT exist!\n"+
			"pathFileName='%v'\n",
			ePrefix.String(),
			pathFileName)

	}

	return fileInfoPlus, err
}

// getFileLastModificationDate
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
//	 	text passed by input parameter, 'errPrefDto'.
//	 	The 'errPrefDto' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fHelperNanobot *fileHelperNanobot) getFileLastModificationDate(
	pathFileName string,
	customTimeFmt string,
	errPrefDto *ePref.ErrPrefixDto) (
	time.Time,
	string,
	error) {

	if fHelperNanobot.lock == nil {
		fHelperNanobot.lock = new(sync.Mutex)
	}

	fHelperNanobot.lock.Lock()

	defer fHelperNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileHelperNanobot."+
			"getFileLastModificationDate()",
		"")

	if err != nil {
		return time.Time{}, "", err
	}

	var pathFileNameDoesExist bool
	var fInfoPlus FileInfoPlus
	var errCode int

	pathFileName,
		pathFileNameDoesExist,
		fInfoPlus,
		err = new(fileHelperMolecule).doesPathFileExist(
		pathFileName,
		PreProcPathCode.AbsolutePath(), // Skip Convert to Absolute Path
		ePrefix,
		"pathFileName")

	if err != nil {
		return time.Time{}, "", err
	}

	if !pathFileNameDoesExist {

		return time.Time{}, "",
			fmt.Errorf("%v\n"+
				"ERROR: 'pathFileName' DOES NOT EXIST!\n"+
				"pathFileName='%v'\n",
				ePrefix.String(),
				pathFileName)
	}

	errCode,
		_,
		customTimeFmt =
		new(fileHelperElectron).isStringEmptyOrBlank(customTimeFmt)

	fmtStr := customTimeFmt

	if errCode < 0 {
		// Default Date Time Format
		fmtStr = "2006-01-02 15:04:05.000000000"
	}

	return fInfoPlus.ModTime(), fInfoPlus.ModTime().Format(fmtStr), nil
}

// getFileMode
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
//	 	text passed by input parameter, 'errPrefDto'.
//	 	The 'errPrefDto' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fHelperNanobot *fileHelperNanobot) getFileMode(
	pathFileName string,
	errPrefDto *ePref.ErrPrefixDto) (
	FilePermissionConfig,
	error) {

	if fHelperNanobot.lock == nil {
		fHelperNanobot.lock = new(sync.Mutex)
	}

	fHelperNanobot.lock.Lock()

	defer fHelperNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileHelperNanobot."+
			"getFileMode()",
		"")

	if err != nil {
		return FilePermissionConfig{}, err
	}

	var pathFileNameDoesExist bool

	var fInfoPlus FileInfoPlus

	pathFileName,
		pathFileNameDoesExist,
		fInfoPlus,
		err = new(fileHelperMolecule).doesPathFileExist(
		pathFileName,
		PreProcPathCode.AbsolutePath(), // Convert to Absolute Path
		ePrefix,
		"pathFileName")

	if err != nil {
		return FilePermissionConfig{}, err
	}

	if !pathFileNameDoesExist {

		return FilePermissionConfig{},
			fmt.Errorf("%v\n"+
				"ERROR: 'pathFileName' DOES NOT EXIST!\n"+
				"pathFileName='%v'\n",
				ePrefix.String(),
				pathFileName)
	}

	fPermCfg, err := new(FilePermissionConfig).
		NewByFileMode(fInfoPlus.Mode(), ePrefix)

	if err != nil {

		return FilePermissionConfig{},
			fmt.Errorf("%v\n"+
				"Error returned by FilePermissionConfig{}.NewByFileMode(fInfoPlus.Mode()).\n"+
				"fInfoPlus.Mode()='%v'\n"+
				"Error='%v'\n",
				ePrefix.String(),
				fInfoPlus.Mode(),
				err.Error())
	}

	return fPermCfg, nil
}

// getFileNameWithExt
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
//	 	text passed by input parameter, 'errPrefDto'.
//	 	The 'errPrefDto' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fHelperNanobot *fileHelperNanobot) getFileNameWithExt(
	pathFileNameExt string,
	errPrefDto *ePref.ErrPrefixDto) (
	fNameExt string,
	isEmpty bool,
	err error) {

	if fHelperNanobot.lock == nil {
		fHelperNanobot.lock = new(sync.Mutex)
	}

	fHelperNanobot.lock.Lock()

	defer fHelperNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	fNameExt = ""
	isEmpty = true
	errCode := 0

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileHelperNanobot."+
			"getFileNameWithExt()",
		"")

	if err != nil {
		return fNameExt, isEmpty, err
	}

	fHelperElectron := new(fileHelperElectron)

	errCode,
		_,
		pathFileNameExt =
		fHelperElectron.isStringEmptyOrBlank(
			pathFileNameExt)

	if errCode == -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'pathFileNameExt' is an empty string!\n",
			ePrefix.String())

		return fNameExt, isEmpty, err
	}

	if errCode == -2 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'pathFileNameExt' consists of blank spaces!",
			ePrefix.String())

		return fNameExt, isEmpty, err
	}

	fHelperAtom := new(fileHelperAtom)

	testPathFileNameExt := fHelperAtom.
		adjustPathSlash(pathFileNameExt)

	volName := fHelperAtom.
		getVolumeName(testPathFileNameExt)

	if volName != "" {
		testPathFileNameExt =
			strings.TrimPrefix(testPathFileNameExt, volName)
	}

	lTestPathFileNameExt := 0

	errCode,
		lTestPathFileNameExt,
		testPathFileNameExt =
		fHelperElectron.
			isStringEmptyOrBlank(testPathFileNameExt)

	if errCode < 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Cleaned version of 'pathFileNameExt', 'testPathFileNameExt' is an empty string!\n",
			ePrefix.String())

		return fNameExt, isEmpty, err
	}

	firstCharIdx,
		lastCharIdx,
		err2 := new(fileHelperMolecule).
		getFirstLastNonSeparatorCharIndexInPathStr(
			testPathFileNameExt,
			ePrefix)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by fh.GetFirstLastNonSeparatorCharIndexInPathStr(testPathFileNameExt).\n"+
			"testPathFileNameExt='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			testPathFileNameExt,
			err2.Error())

		return fNameExt, isEmpty, err
	}

	// There are no alphanumeric characters present.
	// Therefore, there is no file name and extension
	if firstCharIdx == -1 || lastCharIdx == -1 {
		isEmpty = true
		err = nil
		return fNameExt, isEmpty, err
	}

	slashIdxs,
		err2 := fHelperAtom.
		getPathSeparatorIndexesInPathStr(
			testPathFileNameExt,
			ePrefix)

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"Error returned by GetPathSeparatorIndexesInPathStr(testPathFileNameExt).\n"+
			"testPathFileNameExt='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			testPathFileNameExt,
			err2.Error())

		return fNameExt, isEmpty, err
	}

	var dotIdxs []int

	dotIdxs,
		err2 = fHelperAtom.
		getDotSeparatorIndexesInPathStr(
			testPathFileNameExt,
			ePrefix)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by GetDotSeparatorIndexesInPathStr(testPathFileNameExt).\n"+
			"testPathFileNameExt='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			testPathFileNameExt,
			err2.Error())

		return fNameExt, isEmpty, err
	}

	lSlashIdxs := len(slashIdxs)
	lDotIdxs := len(dotIdxs)

	if lSlashIdxs > 0 {
		// This string has path separators

		// Last char is a path separator. Therefore,
		// there is no file name and extension.
		if slashIdxs[lSlashIdxs-1] == lTestPathFileNameExt-1 {
			fNameExt = ""
		} else if lastCharIdx > slashIdxs[lSlashIdxs-1] {

			fNameExt = testPathFileNameExt[slashIdxs[lSlashIdxs-1]+1:]

		} else {
			fNameExt = ""
		}

		if len(fNameExt) == 0 {
			isEmpty = true
		} else {
			isEmpty = false
		}

		err = nil
		return fNameExt, isEmpty, err
	}

	// There are no path separators lSlashIdxs == 0

	if lDotIdxs > 0 {
		// This string has one or more dot separators ('.')

		fNameExt = ""

		if firstCharIdx > dotIdxs[lDotIdxs-1] {
			// Example '.txt' - Valid File name and extension
			// such as '.gitignore'
			fNameExt = testPathFileNameExt[dotIdxs[lDotIdxs-1]:]

		} else if firstCharIdx < dotIdxs[lDotIdxs-1] {
			fNameExt = testPathFileNameExt[firstCharIdx:]
		}

		if len(fNameExt) == 0 {
			isEmpty = true
		} else {
			isEmpty = false
		}

		err = nil
		return fNameExt, isEmpty, err
	}

	// Must be lSlashIdxs == 0 && lDotIdxs ==  0
	// There are no path Separators and there are
	// no dot separators ('.').

	fNameExt = testPathFileNameExt[firstCharIdx:]

	if len(fNameExt) == 0 {
		isEmpty = true
	} else {
		isEmpty = false
	}

	err = nil

	return fNameExt, isEmpty, err
}

// isPathFileString
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
//	 	text passed by input parameter, 'errPrefDto'.
//	 	The 'errPrefDto' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fHelperNanobot *fileHelperNanobot) isPathFileString(
	pathFileStr string,
	errPrefDto *ePref.ErrPrefixDto) (
	pathFileType PathFileTypeCode,
	absolutePathFile string,
	err error) {

	if fHelperNanobot.lock == nil {
		fHelperNanobot.lock = new(sync.Mutex)
	}

	fHelperNanobot.lock.Lock()

	defer fHelperNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	absolutePathFile = ""

	pathFileType = PathFileType.None()

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileHelperNanobot."+
			"isPathFileString()",
		"")

	if err != nil {
		return pathFileType, absolutePathFile, err
	}

	var pathFileDoesExist bool
	var fInfo FileInfoPlus
	var correctedPathFileStr string

	fHelperMolecule := new(fileHelperMolecule)

	correctedPathFileStr,
		pathFileDoesExist,
		fInfo,
		err = fHelperMolecule.doesPathFileExist(
		pathFileStr,
		PreProcPathCode.PathSeparator(), // Convert Path Separators to os default Path Separators
		ePrefix,
		"pathFileStr")

	if err != nil {
		return pathFileType, absolutePathFile, err
	}

	if strings.Contains(correctedPathFileStr, "...") {

		pathFileType = PathFileType.None()

		absolutePathFile = ""

		err = fmt.Errorf("%v\n"+
			"Error: INVALID PATH STRING!\n"+
			"pathFileStr='%v'\n",
			ePrefix.String(),
			correctedPathFileStr)

		return pathFileType, absolutePathFile, err
	}

	osPathSepStr := string(os.PathSeparator)

	if strings.Contains(correctedPathFileStr,
		osPathSepStr+osPathSepStr) {

		pathFileType = PathFileType.None()

		absolutePathFile = ""

		err = fmt.Errorf("%v\n"+
			"Error: Invalid Path File string.\n"+
			"Path File string contains invalid Path Separators.\n"+
			"pathFileStr='%v'\n",
			ePrefix.String(),
			correctedPathFileStr)

		return pathFileType, absolutePathFile, err
	}

	testAbsPathFileStr, err2 := new(fileHelperProton).
		makeAbsolutePath(
			correctedPathFileStr,
			ePrefix)

	if err2 != nil {
		err = fmt.Errorf("Error converting pathFileStr to absolute path.\n"+
			"pathFileStr='%v'\nError='%v'\n",
			correctedPathFileStr, err2.Error())

		return pathFileType, absolutePathFile, err
	}

	volName := fp.VolumeName(testAbsPathFileStr)

	if strings.ToLower(volName) == strings.ToLower(testAbsPathFileStr) ||
		strings.ToLower(volName) == strings.ToLower(pathFileStr) {
		// This is a volume name not a file Name!
		pathFileType = PathFileType.Volume()
		absolutePathFile = volName
		err = nil
		return pathFileType, absolutePathFile, err
	}

	// See if path actually exists on disk and
	// then examine the File Info object returned.
	if pathFileDoesExist {

		if fInfo.IsDir() {

			pathFileType = PathFileType.Path()

			absolutePathFile = testAbsPathFileStr

			err = nil

		} else {

			pathFileType = PathFileType.PathFile()

			absolutePathFile = testAbsPathFileStr

			err = nil

		}

		return pathFileType, absolutePathFile, err
	} // End of if pathFileDoesExist

	// Ok - We know the testPathFileStr does NOT exist on disk

	firstCharIdx,
		lastCharIdx,
		err2 :=
		fHelperMolecule.
			getFirstLastNonSeparatorCharIndexInPathStr(
				correctedPathFileStr,
				ePrefix)

	if err2 != nil {

		pathFileType = PathFileType.None()

		absolutePathFile = ""

		err = fmt.Errorf("%v\n"+
			"Error returned from getFirstLastNonSeparatorCharIndexInPathStr"+
			"(correctedPathFileStr)\n"+
			"correctedPathFileStr='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			correctedPathFileStr,
			err2.Error())

		return pathFileType, absolutePathFile, err
	}

	interiorDotPathIdx :=
		strings.LastIndex(correctedPathFileStr,
			"."+string(os.PathSeparator))

	if interiorDotPathIdx > firstCharIdx {

		pathFileType = PathFileType.None()

		absolutePathFile = ""

		err = fmt.Errorf("%v\n"+
			"Error: INVALID PATH. Invalid interior relative path detected!\n"+
			"correctedPathFileStr='%v'\n",
			ePrefix.String(),
			correctedPathFileStr)

		return pathFileType, absolutePathFile, err
	}

	fHelperAtom := new(fileHelperAtom)

	var slashIdxs []int

	slashIdxs,
		err2 = fHelperAtom.
		getPathSeparatorIndexesInPathStr(
			correctedPathFileStr,
			ePrefix)

	if err2 != nil {

		pathFileType = PathFileType.None()

		absolutePathFile = ""

		err = fmt.Errorf("%v\n"+
			"GetPathSeparatorIndexesInPathStr(correctedPathFileStr) returned error.\n"+
			"testAbsPathFileStr='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			correctedPathFileStr,
			err2.Error())

		return pathFileType, absolutePathFile, err
	}

	dotIdxs, err2 := fHelperAtom.
		getDotSeparatorIndexesInPathStr(
			correctedPathFileStr,
			ePrefix)

	if err2 != nil {

		pathFileType = PathFileType.None()

		absolutePathFile = ""

		err = fmt.Errorf("%v\n"+
			"fh.GetDotSeparatorIndexesInPathStr(correctedPathFileStr) retured error.\n"+
			"correctedPathFileStr='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			correctedPathFileStr,
			err2.Error())

		return pathFileType, absolutePathFile, err
	}

	lDotIdxs := len(dotIdxs)

	lSlashIdxs := len(slashIdxs)

	if lDotIdxs == 0 && lSlashIdxs == 0 && lastCharIdx == -1 {
		// Option # 1
		// There are no valid characters in the string
		pathFileType = PathFileType.None()

		absolutePathFile = ""

		err = fmt.Errorf("%v\n"+
			"No valid characters in parameter 'pathFileStr'\n"+
			"pathFileStr='%v'\n",
			ePrefix.String(),
			correctedPathFileStr)

	} else if lDotIdxs == 0 && lSlashIdxs == 0 && lastCharIdx > -1 {
		// Option # 2
		// String consists only of eligible alphanumeric characters
		// "sometextstring"

		// the string has no dots and no path separators.
		absolutePathFile = testAbsPathFileStr

		// Call it a file!
		// Example: "somefilename"
		pathFileType = PathFileType.File()
		err = nil

	} else if lDotIdxs == 0 && lSlashIdxs > 0 && lastCharIdx == -1 {
		// Option # 3
		// There no dots no characters but string does contain
		// slashes

		if lSlashIdxs > 1 {

			pathFileType = PathFileType.None()

			absolutePathFile = ""

			err = fmt.Errorf("%v\n"+
				"Invalid parameter 'pathFileStr'!\n"+
				"'pathFileStr' consists entirely of path separators!\n"+
				"pathFileStr='%v'\n",
				ePrefix.String(),
				correctedPathFileStr)

		} else {
			// lSlashIdxs must be '1'
			absolutePathFile = testAbsPathFileStr

			// Call it a Directory!
			// Example: "/"
			pathFileType = PathFileType.Path()
			err = nil
		}

	} else if lDotIdxs == 0 && lSlashIdxs > 0 && lastCharIdx > -1 {
		// Option # 4
		// strings contains slashes and characters but no dots.

		absolutePathFile = testAbsPathFileStr

		if lastCharIdx > slashIdxs[lSlashIdxs-1] {
			// Example D:\dir1\dir2\xray
			// It could be a file or a path. We simply
			// can't tell.
			pathFileType = PathFileType.Indeterminate()
			err = nil

		} else {

			// Call it a Directory!
			// Example: "D:\dir1\dir2\"
			pathFileType = PathFileType.Path()
			err = nil

		}

	} else if lDotIdxs > 0 && lSlashIdxs == 0 && lastCharIdx == -1 {
		// Option # 5
		// dots only. Must be "." or ".."

		absolutePathFile = testAbsPathFileStr

		// Call it a Directory!
		// Example: "D:\dir1\dir2\"
		pathFileType = PathFileType.Path()
		err = nil

	} else if lDotIdxs > 0 && lSlashIdxs == 0 && lastCharIdx > -1 {
		// Option # 6
		// Dots and characters only. No slashes.
		// Maybe 'fileName.ext'

		absolutePathFile = testAbsPathFileStr

		// Call it a file!
		// Example: "somefilename"
		pathFileType = PathFileType.File()
		err = nil

	} else if lDotIdxs > 0 && lSlashIdxs > 0 && lastCharIdx == -1 {
		// Option # 7
		// Dots and slashes, but no characters.

		absolutePathFile = testAbsPathFileStr

		// Call it a Directory!
		// Example: ".\"
		pathFileType = PathFileType.Path()
		err = nil

	} else {
		// Option # 8
		// Must be else if lDotIdxs > 0 && lSlashIdxs > 0 && lastCharIdx > -1
		// Contains dots slashes and characters

		absolutePathFile = testAbsPathFileStr

		// If there is a dot after the last path separator
		// and there is a character following the dot.
		// Therefore, this is a filename extension, NOT a directory
		if dotIdxs[lDotIdxs-1] > slashIdxs[lSlashIdxs-1] &&
			dotIdxs[lDotIdxs-1] < lastCharIdx {

			// Call it a path/file!
			// Example: "D:\dir1\dir2\somefilename.txt"
			pathFileType = PathFileType.PathFile()
			err = nil

		} else if lastCharIdx > slashIdxs[lSlashIdxs-1] &&
			dotIdxs[lDotIdxs-1] < slashIdxs[lSlashIdxs-1] {
			// Example: "..\dir1\dir2\sometext"
			// Could be a file could be a directory
			// Cannot determine.
			pathFileType = PathFileType.Indeterminate()
			err = nil

		} else {

			// Call it a Directory!
			// Example: "D:\dir1\dir2\"
			pathFileType = PathFileType.Path()
			err = nil
		}

	}

	return pathFileType, absolutePathFile, err
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
//		dirPath						string
//
//			This string contains the directory path which
//			will be created by this method.
//
//		permission					FilePermissionConfig
//
//			An instance of FilePermissionConfig containing
//			the permission specifications for the new
//			directory to be created from input paramter,
//			'dirPath'.
//
//			The easiest way to configure permissions is
//			to call FilePermissionConfig.New() with
//			a mode string ('modeStr').
//
//			The first character of the 'modeStr' designates the
//			'Entry Type'. Currently, only two 'Entry Type'
//			characters are supported. Therefore, the first
//			character in the 10-character input parameter
//			'modeStr' MUST be either a "-" indicating a file, or
//			a "d" indicating a directory.
//
//			The remaining nine characters in the 'modeStr'
//			represent unix permission bits and consist of three
//			group fields each containing 3-characters. Each
//			character in the three group fields may consist of
//			'r' (Read-Permission), 'w' (Write-Permission), 'x'
//			(Execute-Permission) or '-' signaling no permission or
//			no access allowed. A typical 'modeStr' authorizing
//			permission for full access to a file would be styled
//			as:
//
//			Directory Example: "drwxrwxrwx"
//
//			Groups: - Owner/User, Group, Other
//			From left to right
//			First Characters is Entry Type index 0 ("-")
//
//			First Char index 0 =     "-"   Designates a file
//
//			First Char index 0 =     "d"   Designates a directory
//
//			Char indexes 1-3 = Owner "rwx" Authorizing 'Read',
//		                                  Write' & Execute Permissions for 'Owner'
//
//			Char indexes 4-6 = Group "rwx" Authorizing 'Read', 'Write' & Execute
//		                                  Permissions for 'Group'
//
//			Char indexes 7-9 = Other "rwx" Authorizing 'Read', 'Write' & Execute
//		                                  Permissions for 'Other'
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
//
//	dirPathLabel				string
//
//		The name or label associated with input parameter
//		'dirPath' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "dirPath" will be
//		automatically applied.
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
func (fHelperNanobot *fileHelperNanobot) makeDirAllPerm(
	dirPath string,
	permission FilePermissionConfig,
	dirPathLabel string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fHelperNanobot.lock == nil {
		fHelperNanobot.lock = new(sync.Mutex)
	}

	fHelperNanobot.lock.Lock()

	defer fHelperNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error
	funcName := "fileHelperNanobot." +
		"makeDirAllPerm()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return err
	}

	if len(dirPathLabel) == 0 {
		dirPathLabel = "dirPath"
	}

	errCode := 0

	errCode,
		_,
		dirPath = new(fileHelperElectron).
		isStringEmptyOrBlank(dirPath)

	if errCode == -1 {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is an empty string!\n",
			ePrefix.String(),
			dirPathLabel)
	}

	if errCode == -2 {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' consists of blank spaces!\n",
			ePrefix.String(),
			dirPathLabel)
	}

	var err2 error

	err2 = permission.IsValidInstanceError(
		ePrefix.XCpy("permission"))

	if err2 != nil {

		return fmt.Errorf("%v\n"+
			"Input parameter 'permission' is INVALID!\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			err2.Error())
	}

	var dirPermCode os.FileMode

	dirPermCode,
		err2 = permission.
		GetCompositePermissionMode(
			ePrefix.XCpy(
				"permission"))

	if err2 != nil {
		return fmt.Errorf("%v\n"+
			"ERROR: INVALID Permission Code!\n"+
			"Error returned by permission.GetCompositePermissionMode()\n"+
			"Error=\n%v\n",
			funcName,
			err2.Error())
	}

	dirPath,
		err2 = new(fileHelperProton).
		makeAbsolutePath(
			dirPath,
			ePrefix.XCpy("dirPath"))

	if err2 != nil {
		return fmt.Errorf("%v\n"+
			"Error returned by fh.MakeAbsolutePath(dirPath).\n"+
			"dirPath='%v'\n"+
			"Error=\n%v\n",
			funcName,
			dirPath,
			err2.Error())
	}

	err2 = os.MkdirAll(dirPath, dirPermCode)

	if err2 != nil {
		return fmt.Errorf("%v\n"+
			"Error return from os.MkdirAll(%v, permission).\n"+
			"%v= '%v'\n"+
			"Permission Code ('dirPermCode') = '%v'\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			dirPathLabel,
			dirPathLabel,
			dirPath,
			dirPermCode,
			err2.Error())
	}

	var pathDoesExist bool

	_,
		pathDoesExist,
		_,
		err2 = new(fileHelperMolecule).doesPathFileExist(
		dirPath,
		PreProcPathCode.None(), // Take no Pre-Processing Action
		ePrefix,
		dirPathLabel)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"After creating the %v directory, an error\n"+
			"was returned by fileHelperMolecule.doesPathFileExist()\n"+
			"Error=\n%v\n",
			funcName,
			dirPathLabel,
			err2.Error())

		return err
	}

	if !pathDoesExist {

		err = fmt.Errorf("%v\n"+
			"Error: Directory creation FAILED!. New Directory Path DOES NOT EXIST!\n"+
			"%v='%v'\n",
			ePrefix.String(),
			dirPathLabel,
			dirPath)
	}

	return err
}

// openFileReadOnly
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
//	 	text passed by input parameter, 'errPrefDto'.
//	 	The 'errPrefDto' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fHelperNanobot *fileHelperNanobot) openFileReadOnly(
	pathFileName string,
	errPrefDto *ePref.ErrPrefixDto) (
	filePtr *os.File,
	err error) {

	if fHelperNanobot.lock == nil {
		fHelperNanobot.lock = new(sync.Mutex)
	}

	fHelperNanobot.lock.Lock()

	defer fHelperNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	filePtr = nil

	funcName := "fileHelperNanobot.openFileReadOnly()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return filePtr, err
	}

	var err2 error
	var pathFileNameDoesExist bool
	var fInfoPlus FileInfoPlus

	pathFileName,
		pathFileNameDoesExist,
		fInfoPlus,
		err = new(fileHelperMolecule).doesPathFileExist(
		pathFileName,
		PreProcPathCode.AbsolutePath(), // Convert to Absolute Path
		ePrefix,
		"pathFileName")

	if err != nil {
		return filePtr, err
	}

	if !pathFileNameDoesExist {

		err = fmt.Errorf("%v\n"+
			"ERROR: The input parameter 'pathFileName' DOES NOT EXIST!\n"+
			"pathFileName='%v'\n",
			ePrefix.String(),
			pathFileName)

		return filePtr, err
	}

	if fInfoPlus.IsDir() {

		err = fmt.Errorf("%v\n"+
			"ERROR: The input parameter 'pathFileName' is a 'Directory'\n"+
			"and NOT a path file name.\n"+
			"'pathFileName' is therefore INVALID!\n"+
			"pathFileName='%v'\n",
			ePrefix.String(),
			pathFileName)

		return filePtr, err
	}

	fileOpenCfg, err2 := new(FileOpenConfig).
		New(ePrefix,
			FOpenType.TypeReadOnly(),
			FOpenMode.ModeNone())

	if err2 != nil {
		err =
			fmt.Errorf("%v\n"+
				"Error returned by FileOpenConfig.New("+
				"FOpenType.TypeReadOnly(),"+
				"FOpenMode.ModeNone()).\n"+
				"Error=\n%v\n",
				ePrefix.String(),
				err2.Error())

		return filePtr, err
	}

	fOpenCode, err2 := fileOpenCfg.
		GetCompositeFileOpenCode(ePrefix)

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"Error Creating File Open Code.\n"+
			"fileOpenCfg.GetCompositeFileOpenCode()\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			err2.Error())

		return filePtr, err
	}

	fPermCfg, err2 := new(FilePermissionConfig).
		New(
			"-r--r--r--",
			ePrefix)

	if err2 != nil {

		err =
			fmt.Errorf("%v\n"+
				"Error returned by FilePermissionConfig."+
				"New(\"-r--r--r--\")\n"+
				"Error=\n%v\n",
				funcName,
				err2.Error())

		return filePtr, err
	}

	var fileMode os.FileMode

	fileMode, err2 = fPermCfg.GetCompositePermissionMode(
		ePrefix.XCpy("fPermCfg"))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error Creating File Mode Code.\n"+
			"fPermCfg.GetCompositePermissionMode()\nn"+
			"Error=\n%v\n",
			funcName,
			err2.Error())

		return filePtr, err
	}

	filePtr, err2 = os.OpenFile(
		pathFileName,
		fOpenCode,
		fileMode)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"File Open Error: os.OpenFile()\n"+
			"pathFileName= '%v'"+
			"fOpenCode= '%v'\n"+
			"fileMode= '%v'\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			pathFileName,
			fOpenCode,
			fileMode,
			err2.Error())

		filePtr = nil

		return filePtr, err
	}

	if filePtr == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: os.OpenFile() returned a 'nil' file pointer!\n",
			ePrefix.String())
	}

	return filePtr, err
}

// openFileReadWrite
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
//	 	text passed by input parameter, 'errPrefDto'.
//	 	The 'errPrefDto' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fHelperNanobot *fileHelperNanobot) openFileReadWrite(
	pathFileName string,
	truncateFile bool,
	errPrefDto *ePref.ErrPrefixDto) (
	*os.File,
	error) {

	if fHelperNanobot.lock == nil {
		fHelperNanobot.lock = new(sync.Mutex)
	}

	fHelperNanobot.lock.Lock()

	defer fHelperNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var fPtr *os.File

	var err error

	funcName := "FileHelper.openFileReadWrite()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return fPtr, err
	}

	var pathFileNameDoesExist bool

	var fInfoPlus FileInfoPlus

	pathFileName,
		pathFileNameDoesExist,
		fInfoPlus,
		err = new(fileHelperMolecule).doesPathFileExist(
		pathFileName,
		PreProcPathCode.AbsolutePath(), // Convert to Absolute Path
		ePrefix,
		"pathFileName")

	if err != nil {
		return nil, err
	}

	var fileOpenCfg FileOpenConfig

	if !pathFileNameDoesExist {
		// pathFileName does NOT exist

		fileOpenCfg, err = new(FileOpenConfig).New(
			ePrefix.XCpy("fileOpenCfg<-"),
			FOpenType.TypeReadWrite(),
			FOpenMode.ModeCreate(),
			FOpenMode.ModeAppend())

		if err != nil {
			return nil,
				fmt.Errorf("%v\n"+
					"Error returned by FileOpenConfig{}.New(FOpenType.TypeWriteOnly(),\n"+
					"FOpenMode.ModeCreate(), FOpenMode.ModeAppend()).\n"+
					"Error= \n%v\n",
					funcName,
					err.Error())
		}

	} else {
		// pathFileName does exist

		if fInfoPlus.IsDir() {
			return nil,
				fmt.Errorf("%v\n"+
					"ERROR: Input parameter 'pathFileName' is\n"+
					"a 'Directory' and NOT a file.\n"+
					"pathFileName='%v'\n",
					ePrefix.String(),
					pathFileName)
		}

		if truncateFile {
			// truncateFile == true
			fileOpenCfg, err = new(FileOpenConfig).New(
				ePrefix.XCpy("fileOpenCfg<-truncateFile=true"),
				FOpenType.TypeReadWrite(),
				FOpenMode.ModeTruncate())

			if err != nil {
				return nil,
					fmt.Errorf("%v\n"+
						"Error returned by FileOpenConfig{}.New(FOpenType.TypeReadWrite(),\n"+
						"FOpenMode.ModeTruncate()).\n"+
						"Error= \n%v\n",
						funcName,
						err.Error())
			}

		} else {
			// truncateFile == false
			fileOpenCfg, err = new(FileOpenConfig).New(
				ePrefix.XCpy("fileOpenCfg<-truncate=false"),
				FOpenType.TypeReadWrite(),
				FOpenMode.ModeAppend())

			if err != nil {
				return nil,
					fmt.Errorf("%v\n"+
						"Error returned by FileOpenConfig{}.New(FOpenType.TypeReadWrite(),\n"+
						"FOpenMode.ModeAppend()).\n"+
						"Error= \n%v\n",
						funcName,
						err.Error())
			}
		}
	}

	fOpenCode, err := fileOpenCfg.
		GetCompositeFileOpenCode(ePrefix)

	if err != nil {
		return nil,
			fmt.Errorf("%v\n"+
				"Error Return from fileOpenCfg.GetCompositeFileOpenCode()\n"+
				"Error= \n%v\n,",
				ePrefix.String(),
				err.Error())
	}

	var fPermCfg FilePermissionConfig

	fPermCfg, err = new(FilePermissionConfig).New(
		"-rwxrwxrwx",
		ePrefix.XCpy("fPermCfg<-"))

	if err != nil {
		return nil,
			fmt.Errorf("%v\n"+
				"Error returned by FilePermissionConfig{}.New(\"-rwxrwxrwx\")\n"+
				"Error= \n%v\n",
				funcName,
				err.Error())
	}

	var fileMode os.FileMode

	fileMode,
		err = fPermCfg.GetCompositePermissionMode(
		ePrefix.XCpy("fileMode<-fPermCfg"))

	if err != nil {
		return nil,
			fmt.Errorf("%v\n"+
				"Error return from fPermCfg.GetCompositePermissionMode()\n"+
				"Error= \n%v\n",
				funcName,
				err.Error())
	}

	fPtr, err = os.OpenFile(pathFileName, fOpenCode, fileMode)

	if err != nil {

		return nil, fmt.Errorf("%v\n"+
			"File Open Error returned from\n"+
			"os.OpenFile(pathFileName, fOpenCode, fileMode)\n"+
			"pathFileName= '%v'\n"+
			"fOpenCode= '%v'\n"+
			"fileMode= '%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			pathFileName,
			fOpenCode,
			fileMode,
			err.Error())
	}

	if fPtr == nil {

		return nil, fmt.Errorf("%v\n"+
			"ERROR: os.OpenFile() returned a 'nil' file pointer!\n"+
			"pathFileName='%v'\n",
			ePrefix.String(),
			pathFileName)
	}

	return fPtr, err
}

// openFileWriteOnly
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
//	 	text passed by input parameter, 'errPrefDto'.
//	 	The 'errPrefDto' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fHelperNanobot *fileHelperNanobot) openFileWriteOnly(
	pathFileName string,
	truncateFile bool,
	errPrefDto *ePref.ErrPrefixDto) (
	*os.File,
	error) {

	if fHelperNanobot.lock == nil {
		fHelperNanobot.lock = new(sync.Mutex)
	}

	fHelperNanobot.lock.Lock()

	defer fHelperNanobot.lock.Unlock()

	funcName := "FileHelper.OpenFileWriteOnly() "

	var ePrefix *ePref.ErrPrefixDto

	var fPtr *os.File

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return fPtr, err
	}

	var pathFileNameDoesExist bool
	var fInfoPlus FileInfoPlus

	pathFileName,
		pathFileNameDoesExist,
		fInfoPlus,
		err = new(fileHelperMolecule).doesPathFileExist(
		pathFileName,
		PreProcPathCode.AbsolutePath(), // Convert to Absolute Path
		ePrefix,
		"pathFileName")

	if err != nil {
		return nil, err
	}

	var fileOpenCfg FileOpenConfig

	if !pathFileNameDoesExist {
		// The pathFileName DOES NOT EXIST!

		fileOpenCfg, err = new(FileOpenConfig).New(
			ePrefix.XCpy("fileOpenCfg<-!pathFileNameDoesExist"),
			FOpenType.TypeWriteOnly(),
			FOpenMode.ModeCreate(),
			FOpenMode.ModeAppend())

		if err != nil {
			return nil,
				fmt.Errorf("%v\n"+
					"Error returned by FileOpenConfig{}.New(FOpenType.TypeWriteOnly(),\n"+
					"FOpenMode.ModeCreate(), FOpenMode.ModeAppend()).\n"+
					"Error= \n%v\n",
					funcName,
					err.Error())
		}

	} else {
		// The pathFileName DOES EXIST!

		if fInfoPlus.IsDir() {

			return nil,
				fmt.Errorf("%v\n"+
					"ERROR: Input parameter 'pathFileName' is\n"+
					"a 'Directory' and NOT a file.\n"+
					"pathFileName='%v'\n",
					ePrefix.String(),
					pathFileName)
		}

		if truncateFile {

			// truncateFile == true; Set Mode 'Truncate'
			fileOpenCfg, err = new(FileOpenConfig).New(
				ePrefix.XCpy("fileOpenCfg<-truncateFile=true"),
				FOpenType.TypeWriteOnly(),
				FOpenMode.ModeTruncate())

			if err != nil {
				return nil,
					fmt.Errorf("%v\n"+
						"Error returned by FileOpenConfig{}.New(FOpenType.TypeWriteOnly(),\n"+
						"FOpenMode.ModeTruncate()).\n"+
						"Error= \n%v\n",
						funcName,
						err.Error())
			}

		} else {

			// truncateFile == false; Set Mode 'Append'
			fileOpenCfg, err = new(FileOpenConfig).New(
				ePrefix.XCpy("fileOpenCfg<-ModeAppend"),
				FOpenType.TypeWriteOnly(),
				FOpenMode.ModeAppend())

			if err != nil {
				return nil,
					fmt.Errorf("%v\n"+
						"Error returned by FileOpenConfig{}.New(FOpenType.TypeWriteOnly(),\n"+
						"FOpenMode.ModeAppend()).\n"+
						"Error= \n%v\n",
						funcName,
						err.Error())
			}
		}
	}

	fOpenCode, err := fileOpenCfg.
		GetCompositeFileOpenCode(ePrefix)

	if err != nil {

		return nil,
			fmt.Errorf("%v\n"+
				"Error creating File Open Code.\n"+
				"Error return from fileOpenCfg.GetCompositeFileOpenCode()\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				err.Error())
	}

	fPermCfg, err := new(FilePermissionConfig).New(
		"--wx-wx-wx",
		ePrefix.XCpy(
			"fPermCfg<---wx-wx-wx"))

	if err != nil {

		return nil,
			fmt.Errorf("%v\n"+
				"Error returned by FilePermissionConfig{}.New(\"-rwxrwxrwx\")\n"+
				"Error= \n%v\n",
				funcName,
				err.Error())
	}

	var fileMode os.FileMode

	fileMode,
		err = fPermCfg.GetCompositePermissionMode(
		ePrefix.XCpy(
			"fileMode<-fPermCfg"))

	if err != nil {

		return nil, fmt.Errorf("%v\n"+
			"Error creating file mode code.\n"+
			"Error= \n%v\n",
			funcName,
			err.Error())
	}

	fPtr, err = os.OpenFile(pathFileName, fOpenCode, fileMode)

	if err != nil {

		return nil, fmt.Errorf("%v\n"+
			"Error returned from os.OpenFile().\n"+
			"pathFileName='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			pathFileName,
			err.Error())
	}

	if fPtr == nil {
		return nil, fmt.Errorf("%v\n"+
			"ERROR: File pointer returned from os.OpenFile() is 'nil'!\n",
			ePrefix.String())
	}

	return fPtr, err
}
