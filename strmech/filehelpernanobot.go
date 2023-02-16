package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"os"
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
