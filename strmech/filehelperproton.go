package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"os"
	fp "path/filepath"
	"sync"
)

type fileHelperProton struct {
	lock *sync.Mutex
}

// addPathSeparatorToEndOfPathStr
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
func (fHelpProton *fileHelperProton) addPathSeparatorToEndOfPathStr(
	pathStr string,
	errorPrefix interface{}) (string, error) {

	if fHelpProton.lock == nil {
		fHelpProton.lock = new(sync.Mutex)
	}

	fHelpProton.lock.Lock()

	defer fHelpProton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"fileHelperProton."+
			"addPathSeparatorToEndOfPathStr()",
		"")

	if err != nil {

		return "", err
	}

	var errCode, lStr int

	errCode,
		lStr,
		pathStr = new(fileHelperElectron).
		isStringEmptyOrBlank(pathStr)

	if errCode == -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'pathStr' is an empty string!\n",
			ePrefix.String())

		return "", err

	}

	if errCode == -2 {

		err = fmt.Errorf(
			"%v\n"+
				"Error: Input parameter 'pathStr' consists of blank spaces!\n",
			ePrefix.String())

		return "", err
	}

	if pathStr[lStr-1] == os.PathSeparator {
		return pathStr, nil
	}

	var newPathStr string

	if pathStr[lStr-1] == '/' && '/' != os.PathSeparator {

		newPathStr = pathStr[0 : lStr-1]

		newPathStr += string(os.PathSeparator)

		return newPathStr, nil
	}

	if pathStr[lStr-1] == '\\' && '\\' != os.PathSeparator {

		newPathStr = pathStr[0 : lStr-1]

		newPathStr += string(os.PathSeparator)

		return newPathStr, nil
	}

	newPathStr = pathStr + string(os.PathSeparator)

	return newPathStr, nil
}

// makeAbsolutePath
//
// Supply a relative path or any path string and resolve
// that path to an Absolute path. This method calls
// filepath.Abs() to generate the absolute path.
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
// Note: Clean() is called on result by filepath.Abs().
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
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (fHelpProton *fileHelperProton) makeAbsolutePath(
	relPath string,
	errPrefDto *ePref.ErrPrefixDto) (
	string,
	error) {

	if fHelpProton.lock == nil {
		fHelpProton.lock = new(sync.Mutex)
	}

	fHelpProton.lock.Lock()

	defer fHelpProton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileHelperProton."+
			"makeAbsolutePath()",
		"")

	if err != nil {
		return "", err
	}

	errCode := 0

	fHelperElectron := fileHelperElectron{}

	errCode,
		_,
		relPath =
		fHelperElectron.
			isStringEmptyOrBlank(relPath)

	if errCode == -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'relPath' is an empty string!\n",
			ePrefix.String())

		return "", err
	}

	if errCode == -2 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'relPath' consists of blank spaces!\n",
			ePrefix.String())

		return "", err
	}

	testRelPath :=
		new(fileHelperAtom).adjustPathSlash(relPath)

	errCode, _, testRelPath =
		fHelperElectron.isStringEmptyOrBlank(testRelPath)

	if errCode < 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input Parameter 'relPath' adjusted for path Separators is an EMPTY string!\n",
			ePrefix.String())

		return "", err
	}

	var err2 error
	var p string

	p, err2 = fp.Abs(testRelPath)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned from fp.Abs(testRelPath).\n"+
			"testRelPath='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			testRelPath,
			err2.Error())

		return "", err
	}

	return p, err
}
