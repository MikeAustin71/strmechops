package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"os"
	"strings"
	"sync"
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

// DoesFileExist - Returns a boolean value designating whether the passed
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
