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

// searchFileOlderThan
//
// This method is called to determine whether the file
// described by the input parameter 'info' is a 'match'
// for the File Selection Criteria,
// 'fileSelectCriteria.FilesOlderThan'.
//
// If the file modification date time occurs before the
// 'fileSelectCriteria.FilesOlderThan' date time, the
// return value 'isFileOlderThanMatch' is set to 'true'.
//
// If 'fileSelectCriteria.FilesOlderThan' is set to
// time.Time zero ( the default or zero value for this
// type), the return value 'isFileOlderThanSet' is set to
// 'false' signaling that this search criterion is NOT
// active.
//
// Note:
//
// Input parameter 'info' is of type os.FileInfo.  This
// means the user may substitute a type 'FileInfoPlus'
// object for the 'info' parameter because 'FileInfoPlus'
// implements the 'os.FileInfo' interface.
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
// ----------------------------------------------------------------
//
// # Return Values
//
//	isFileOlderThanSet			bool
//
//		This boolean return value signals whether the
//		FileOlderThan file search criteria is set,
//		engaged and active.
//
//		If the 'fileSelectCriteria.FileOlderThan'
//	 	data element is set to time.Time == zero (0), it
//	 	means that the 'FileOlderThan' search criteria is
//	 	not set or engaged. In this case, the
//	 	'isFileOlderThanSet' parameter will return a
//	 	value of 'false' meaning that the search criteria
//	 	is NOT active or engaged.
//
//	isFileOlderThanMatch		bool
//
//		'fileSelectCriteria.FileOlderThan' consists of a
//		time.Time date element. This represents a file
//		search pattern designed to identify files for
//	 	deletion which are older than the
//	 	'fileSelectCriteria.FileOlderThan' date. If
//		the file specified by input parameter 'info',
//		is older than this date, the file search
//		operation is classified as a 'match', and
//		the return value of this parameter,
//		'isFileOlderThanMatch', is set to 'true'.
func (fHelpElectron *fileHelperElectron) searchFileOlderThan(
	info os.FileInfo,
	fileSelectCriteria FileSelectionCriteria) (
	isFileOlderThanSet bool,
	isFileOlderThanMatch bool) {

	if fHelpElectron.lock == nil {
		fHelpElectron.lock = new(sync.Mutex)
	}

	fHelpElectron.lock.Lock()

	defer fHelpElectron.lock.Unlock()

	if fileSelectCriteria.FilesOlderThan.IsZero() {

		isFileOlderThanSet = false

		isFileOlderThanMatch = false

		return isFileOlderThanSet, isFileOlderThanMatch
	}

	if fileSelectCriteria.FilesOlderThan.After(info.ModTime()) {

		isFileOlderThanSet = true

		isFileOlderThanMatch = true

		return isFileOlderThanSet, isFileOlderThanMatch
	}

	isFileOlderThanSet = true
	isFileOlderThanMatch = false

	return isFileOlderThanSet, isFileOlderThanMatch
}

// searchFilePatternMatch
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
// Input parameter 'info' is of type os.FileInfo.  This
// means the user may substitute a type 'FileInfoPlus'
// object for the 'info' parameter because 'FileInfoPlus'
// implements the 'os.FileInfo' interface.
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
//	isPatternSet				bool
//
//		This boolean return value signals whether one or
//		more file search patterns are set, engaged and
//		active.
//
//		If the 'fileSelectCriteria.FileNamePatterns'
//	 	array is empty or if it contains only empty
//	  	strings, this return value is set to 'false'
//	  	signaling that the pattern file search
//	   	selection criterion is NOT active.
//
//		When set to 'true', this parameter signals that
//		one or more file search patterns are set, engaged
//		and active.
//
//	isPatternMatch				bool
//
//		'fileSelectCriteria.FileNamePatterns' consist of
//		a string array. If the pattern signified by any
//		element in the string array is a 'match', the
//		return value 'isPatternMatch' is set to true.
//
//		The criteria for a 'match' can be set through
//		input paramter,
//		'fileSelectCriteria.SelectCriterionMode'.
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
//		If the call to path.Match() fails it will return
//		an error of type ErrBadPattern. This error may be
//		unpacked to reveal additional technical details
//		regarding the failure to create a file.
//
//		If an error is returned by path.Match() it will
//		be returned in its original form through
//		parameter, 'lowLevelErr'.
//
//		If no ErrBadPattern occurs, 'lowLevelErr' will be
//		set to 'nil'.
func (fHelpElectron *fileHelperElectron) searchFilePatternMatch(
	info os.FileInfo,
	fileSelectCriteria FileSelectionCriteria,
	errPrefDto *ePref.ErrPrefixDto) (
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
		msgError = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileHelperElectron."+
			"searchFilePatternMatch()",
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
