package strmech

import (
	"strings"
	"sync"
)

type fileHelperElectron struct {
	lock *sync.Mutex
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
