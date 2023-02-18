package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"os"
	fp "path/filepath"
	"strings"
	"sync"
)

type fileHelperElectron struct {
	lock *sync.Mutex
}

// cleanPathStr
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
func (fHelpElectron *fileHelperElectron) cleanPathStr(pathStr string) string {

	if fHelpElectron.lock == nil {
		fHelpElectron.lock = new(sync.Mutex)
	}

	fHelpElectron.lock.Lock()

	defer fHelpElectron.lock.Unlock()

	return fp.Clean(pathStr)
}

// isStringEmptyOrBlank
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
// /
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
func (fHelpElectron *fileHelperElectron) isStringEmptyOrBlank(
	testStr string) (
	errCode int,
	strLen int,
	newStr string) {

	if fHelpElectron.lock == nil {
		fHelpElectron.lock = new(sync.Mutex)
	}

	fHelpElectron.lock.Lock()

	defer fHelpElectron.lock.Unlock()

	errCode = 0
	strLen = 0
	newStr = ""

	if len(testStr) == 0 {
		errCode = -1
		return errCode, strLen, newStr
	}

	newStr = strings.TrimLeft(testStr, " ")

	newStr = strings.TrimRight(newStr, " ")

	strLen = len(newStr)

	if strLen == 0 {
		errCode = -2
		newStr = ""
		return errCode, strLen, newStr
	}

	errCode = 0

	return errCode, strLen, newStr
}

// SearchFilePatternMatch
//
// This method is used to determine whether a file
// described by the 'info' parameter meets the specified
// File Selection Criteria and is judged to be a match
// for the fileSelectCriteria FileNamePattern.
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
// Input parameter 'info' is of type os.FileInfo.  You
// can substitute a type 'FileInfoPlus' object for the
// 'info' parameter because 'FileInfoPlus' implements
// the 'os.FileInfo' interface.
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
func (fHelpElectron *fileHelperElectron) SearchFilePatternMatch(
	info os.FileInfo,
	fileSelectCriteria FileSelectionCriteria,
	errorPrefix interface{}) (
	isPatternSet bool,
	isPatternMatch bool,
	msgError error,
	lowLevelErr error) {

	if fHelpElectron.lock == nil {
		fHelpElectron.lock = new(sync.Mutex)
	}

	fHelpElectron.lock.Lock()

	defer fHelpElectron.lock.Unlock()

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

	isPatternSet = fileSelectCriteria.ArePatternsActive()

	if !isPatternSet {
		// bona fide result. No errors here!
		isPatternSet = false
		isPatternMatch = false
		msgError = nil
		lowLevelErr = nil

		return isPatternSet, isPatternMatch, msgError, lowLevelErr
	}

	lPats := len(fileSelectCriteria.FileNamePatterns)

	var matched bool

	for i := 0; i < lPats; i++ {

		matched,
			lowLevelErr = fp.Match(
			fileSelectCriteria.FileNamePatterns[i],
			info.Name())

		if lowLevelErr != nil {

			isPatternSet = true

			msgError = fmt.Errorf("%v\n"+
				"Error returned from filepath.Match(fileSelectCriteria.FileNamePatterns[i], "+
				"info.Name())\n"+
				"fileSelectCriteria.FileNamePatterns[i]='%v'\n"+
				"info.Name()='%v'\nError='%v'\n",
				ePrefix.String(),
				fileSelectCriteria.FileNamePatterns[i],
				info.Name(),
				lowLevelErr.Error())

			isPatternMatch = false

			return isPatternSet, isPatternMatch, msgError, lowLevelErr
		}

		if matched {

			isPatternSet = true
			isPatternMatch = true

			return isPatternSet, isPatternMatch, msgError, lowLevelErr
		}
	}

	isPatternSet = true
	isPatternMatch = false

	return isPatternSet, isPatternMatch, msgError, lowLevelErr
}
