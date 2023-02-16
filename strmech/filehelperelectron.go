package strmech

import (
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
