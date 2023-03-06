package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"os"
	"path"
	fp "path/filepath"
	"strings"
	"sync"
	"time"
)

type fileHelperMolecule struct {
	lock *sync.Mutex
}

// doesPathFileExist
//
// This method will determine whether a path and file
// does or does not exist.
//
// This method calls os.Stat(dirPath) which returns an
// error which is one of two types:
//
//  1. A Non-Path Error - An error which is not path
//     related. It signals some other type of error
//     which makes it impossible to determine if the
//     path actually exists. These types of errors
//     generally relate to "access denied" situations,
//     but there may be other reasons behind non-path
//     errors. If a non-path error is returned, no valid
//     existence test can be performed on the file path.
//
//     or
//
//  2. A Bona Fide Path Error - indicates that the path
//     definitely does not exist.
//
// To deal with these types of errors, this method will
// test path existence up to three times before returning
// a non-path error.
func (fHelpMolecule *fileHelperMolecule) doesPathFileExist(
	filePath string,
	preProcessCode PreProcessPathCode,
	errorPrefix interface{},
	filePathTitle string) (
	absFilePath string,
	filePathDoesExist bool,
	fInfo FileInfoPlus,
	nonPathError error) {

	if fHelpMolecule.lock == nil {
		fHelpMolecule.lock = new(sync.Mutex)
	}

	fHelpMolecule.lock.Lock()

	defer fHelpMolecule.lock.Unlock()
	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		nonPathError = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileHelper."+
			"AreSameFile()",
		"")

	if nonPathError != nil {
		return absFilePath, filePathDoesExist, fInfo, nonPathError
	}

	absFilePath = ""

	filePathDoesExist = false

	fInfo = FileInfoPlus{}

	nonPathError = nil

	if len(filePathTitle) == 0 {
		filePathTitle = "filePath"
	}

	errCode := 0

	errCode, _, filePath = new(fileHelperElectron).
		isStringEmptyOrBlank(filePath)

	if errCode == -1 {

		nonPathError = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is an empty string!\n",
			ePrefix.String(),
			filePathTitle)

		return absFilePath, filePathDoesExist, fInfo, nonPathError
	}

	if errCode == -2 {

		nonPathError = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' consists of blank spaces!\n",
			ePrefix.String(),
			filePathTitle)

		return absFilePath, filePathDoesExist, fInfo, nonPathError
	}

	if preProcessCode == PreProcPathCode.PathSeparator() {

		absFilePath = new(fileHelperAtom).adjustPathSlash(filePath)

	} else if preProcessCode == PreProcPathCode.AbsolutePath() {

		absFilePath,
			nonPathError = new(fileHelperProton).
			makeAbsolutePath(
				filePath,
				ePrefix.XCpy(
					"absFilePath<-filePath"))

		if nonPathError != nil {

			absFilePath = ""

			return absFilePath, filePathDoesExist, fInfo, nonPathError
		}

	} else {
		// For any other PreProcPathCode value, apply no pre-processing to
		absFilePath = filePath
	}

	var info os.FileInfo

	for i := 0; i < 3; i++ {

		filePathDoesExist = false
		fInfo = FileInfoPlus{}
		nonPathError = nil
		var err error

		info, err = os.Stat(absFilePath)

		if err != nil {

			if os.IsNotExist(err) {

				filePathDoesExist = false
				fInfo = FileInfoPlus{}
				nonPathError = nil
				return absFilePath, filePathDoesExist, fInfo, nonPathError
			}
			// err == nil and err != os.IsNotExist(err)
			// This is a non-path error. The non-path error will
			// be tested up to 3-times before it is returned.
			nonPathError = fmt.Errorf(
				"%v\n"+
					"Non-Path error returned by os.Stat(%v)\n"+
					"%v='%v'\n"+
					"Error='%v'\n",
				ePrefix.String(),
				filePathTitle,
				filePathTitle,
				filePath,
				err.Error())

			fInfo = FileInfoPlus{}
			filePathDoesExist = false

		} else {
			// err == nil
			// The path really does exist!
			filePathDoesExist = true
			nonPathError = nil
			fInfo = new(FileInfoPlus).NewFromFileInfo(info)
			return absFilePath, filePathDoesExist, fInfo, nonPathError
		}

		time.Sleep(30 * time.Millisecond)
	}

	return absFilePath, filePathDoesExist, fInfo, nonPathError
}

// getFirstLastNonSeparatorCharIndexInPathStr
//
// Basically, this method returns the two path
// string indexes:
//
//		(1) The first alphanumeric character in a path string
//	     (reading from left to right).
//
//		(2) The last alphanumeric character in a path string
//	     (reading from left to right).
//
// Specifically, the character identified by the index
// must not be a path Separator ('\', '/') and it must
// not be a dot ('.').
//
// If the first Non-Separator char is found, this method
// will return an integer index which is greater than or
// equal to zero plus an error value of nil.
//
// The first character found will never be part of the
// volume name.
//
// Example On Windows:
//
//	"D:\fDir1\fDir2" - first character index will
//						be 3 denoting character 'f'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathStr						string
//
//		A string specifying a file or directory path.
//		Basically, this method returns the indexes of
//		the first and last alphanumeric characters in
//		this path string. See the example above.
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
//		firstIdx					int
//
//			The string index in input parameter 'pathStr'
//			identifying the first alphanumeric character.
//
//			The character identified by this index will not
//			be a path Separator ('\', '/') and or a dot ('.').
//
//			If this returned value is less than zero, it signals
//			that no valid first character was found.
//
//			The first character found will never be part of the
//			volume name.
//
//			Example On Windows:
//
//			"D:\fDir1\fDir2" - first character index will
//	                         be 3 denoting character 'f'.
//
//		lastIdx						int
//
//			The string index in input parameter 'pathStr'
//			identifying the last alphanumeric character.
//
//			The character identified by this index will not
//			be a path Separator ('\', '/') and or a dot ('.').
//
//			The last character found will never be part of the
//			volume name.
//
//		error
//
//			If this method completes successfully, the
//			returned error Type is set equal to 'nil'. If
//			errors are encountered during processing, the
//			returned error Type will encapsulate an error
//			message.
//
//			If an error message is returned, the text value
//			for input parameter 'errPrefDto' (error prefix)
//			will be prefixed or attached at the beginning of
//			the error message.
func (fHelpMolecule *fileHelperMolecule) getFirstLastNonSeparatorCharIndexInPathStr(
	pathStr string,
	errPrefDto *ePref.ErrPrefixDto) (
	firstIdx,
	lastIdx int,
	err error) {

	if fHelpMolecule.lock == nil {
		fHelpMolecule.lock = new(sync.Mutex)
	}

	fHelpMolecule.lock.Lock()

	defer fHelpMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	firstIdx = -1
	lastIdx = -1
	errCode := 0

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errPrefDto,
		"fileHelperMolecule."+
			"getFirstLastNonSeparatorCharIndexInPathStr()",
		"")

	if err != nil {
		return firstIdx, lastIdx, err
	}

	errCode, _, pathStr =
		new(fileHelperElectron).isStringEmptyOrBlank(pathStr)

	if errCode == -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'pathStr' is an empty string!\n",
			ePrefix.String())

		return firstIdx, lastIdx, err
	}

	if errCode == -2 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'pathStr' consists of blank spaces!\n",
			ePrefix.String())

		return firstIdx, lastIdx, err
	}

	pathStr = new(fileHelperAtom).adjustPathSlash(pathStr)

	lPathStr := 0

	errCode, lPathStr, pathStr = new(fileHelperElectron).isStringEmptyOrBlank(pathStr)

	if errCode < 0 {

		err = fmt.Errorf("%v\n"+
			"Error: After path Separator adjustment, 'pathStr' is an empty string!\n",
			ePrefix.String())

		return firstIdx, lastIdx, err
	}

	// skip the volume name. Don't count
	// first characters in the volume name
	volName := fp.VolumeName(pathStr)
	lVolName := len(volName)

	startIdx := 0

	if lVolName > 0 {
		startIdx = lVolName
	}

	var rChar rune

	forbiddenTextChars := []rune{os.PathSeparator,
		'\\',
		'/',
		'|',
		'.',
		'&',
		'!',
		'%',
		'$',
		'#',
		'@',
		'^',
		'*',
		'(',
		')',
		'-',
		'_',
		'+',
		'=',
		'[',
		'{',
		']',
		'}',
		'|',
		'<',
		'>',
		',',
		'~',
		'`',
		':',
		';',
		'"',
		'\'',
		'\n',
		'\t',
		'\r'}

	lForbiddenTextChars := len(forbiddenTextChars)

	for i := startIdx; i < lPathStr; i++ {
		rChar = rune(pathStr[i])
		isForbidden := false

		for j := 0; j < lForbiddenTextChars; j++ {
			if rChar == forbiddenTextChars[j] {
				isForbidden = true
			}

		}

		if isForbidden == false {

			if firstIdx == -1 {
				firstIdx = i
			}

			lastIdx = i
		}

	}

	err = nil

	return firstIdx, lastIdx, err
}

// getAbsPathFromFilePath
//
// Receives a string containing the path, file name
// and extension.
//
// This method will then return the absolute value of
// that path, file name and file extension.
//
// "An absolute or full path points to the same location
// in a file system, regardless of the current working
// directory. To do that, it must include the root
// directory."
//
//	Wikipedia
//
// This method therefore converts path element contained
// in input parameter 'filePath' to an absolute path.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://en.wikipedia.org/wiki/Path_(computing)#Absolute_and_relative_paths
//	https://en.wikipedia.org/wiki/Path_(computing)
//	https://pkg.go.dev/path/filepath@go1.20.1#Abs
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	filePath					string
//
//		This strings contains the path, file name and
//		file extension. This method will convert the path
//		element to an absolute path.
//
//		"An absolute or full path points to the same
//	 	location in a file system, regardless of the
//	 	current working directory. To do that, it must
//	 	include the root directory."
//
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
//		This string returns the absolute file path.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (fHelpMolecule *fileHelperMolecule) getAbsPathFromFilePath(
	filePath string,
	errorPrefix interface{}) (
	string,
	error) {

	if fHelpMolecule.lock == nil {
		fHelpMolecule.lock = new(sync.Mutex)
	}

	fHelpMolecule.lock.Lock()

	defer fHelpMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"fileHelperMolecule."+
			"getAbsPathFromFilePath()",
		"")

	if err != nil {
		return "", err
	}

	errCode := 0

	fHelperElectron := new(fileHelperElectron)

	errCode,
		_,
		filePath = fHelperElectron.isStringEmptyOrBlank(filePath)

	if errCode == -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'filePath' is an empty string!\n",
			ePrefix.String())

		return "", err

	}

	if errCode == -2 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'filePath' consists of blank spaces!\n",
			ePrefix.String())

		return "", err

	}

	testFilePath := new(fileHelperAtom).adjustPathSlash(filePath)

	errCode,
		_,
		testFilePath =
		fHelperElectron.
			isStringEmptyOrBlank(
				testFilePath)

	if errCode < 0 {

		err = fmt.Errorf("%v\n"+
			"Error: After adjusting path Separators,\n"+
			"'filePath' resolves to an empty string!\n",
			ePrefix.String())

		return "", err
	}

	var absPath string
	var err2 error

	absPath,
		err2 = new(fileHelperProton).
		makeAbsolutePath(
			testFilePath,
			ePrefix.XCpy("absPath<-"))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned from fh.MakeAbsolutePath(testFilePath).\n"+
			"testFilePath='%v'\nError='%v'\n",
			ePrefix.String(),
			testFilePath,
			err.Error())

		return "", err
	}

	return absPath, nil
}

// getLastPathElement
//
// Analyzes a 'pathName' string and returns the last
// element in the path. If 'pathName' ends in a path
// separator ('/'), this method returns an empty
// string.
//
// ----------------------------------------------------------------
//
// # Usage
//
//	pathName = '../dir1/dir2/fileName.ext' will return "fileName.ext"
//	pathName = '../dir1/dir2/' will return ""
//	pathName = 'fileName.ext' will return "fileName.ext"
//	pathName = '../dir1/dir2/dir3' will return "dir3"
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
//		If this method completes successfully, this
//		string will return the last path element found in
//		input parameter 'pathName'.
//
//		If 'pathName' ends in a path separator ('/'),
//		this string parameter will be returned as an
//		empty or zero length string.
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
func (fHelpMolecule *fileHelperMolecule) getLastPathElement(
	pathName string,
	errPrefDto *ePref.ErrPrefixDto) (
	string,
	error) {

	if fHelpMolecule.lock == nil {
		fHelpMolecule.lock = new(sync.Mutex)
	}

	fHelpMolecule.lock.Lock()

	defer fHelpMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileHelperMolecule."+
			"getLastPathElement()",
		"")

	if err != nil {
		return "", err
	}

	errCode := 0

	errCode, _, pathName = new(fileHelperElectron).
		isStringEmptyOrBlank(pathName)

	if errCode == -1 {
		return "",
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'pathName' is an empty string!\n",
				ePrefix.String())
	}

	if errCode == -2 {
		return "",
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'pathName' consists of blank spaces!\n",
				ePrefix.String())
	}

	adjustedPath := new(fileHelperAtom).
		adjustPathSlash(pathName)

	resultAry := strings.
		Split(
			adjustedPath,
			string(os.PathSeparator))

	lResultAry := len(resultAry)

	if lResultAry == 0 {
		return adjustedPath, nil
	}

	return resultAry[lResultAry-1], nil
}

// isAbsolutePath
//
// Compares the input parameter 'pathStr' to the absolute
// path representation for 'pathStr' to determine whether
// 'pathStr' represents an absolute path.
//
// This method differs from isAbsolutePathByCompare() in
// that this method calls low level method
// filePath.IsAbsolute() to determine if a path is an
// absolute path.
//
// ----------------------------------------------------------------
//
// Absolute Path Definition (Wikipedia):
//
//	An absolute or full path points to the same location
//	in a file system, regardless of the current working
//	directory. To do that, it must include the root
//	directory.
//
//	By contrast, a relative path starts from some given
//	working directory, avoiding the need to provide the
//	full absolute path. A filename can be considered as
//	a relative path based at the current working
//	directory. If the working directory is not the file's
//	parent directory, a file not found error will result
//	if the file is addressed by its name.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathStr						string
//
//		This string holds the file path which will be
//		analyzed to determine if it is an absolute path.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		This method will analyze input parameter
//		'pathStr' to determine if it is an absolute path.
//
//		This determination is made by comparing 'pathStr'
//		to the absolute path constructed with 'pathStr'.
//
//		If input parameter 'pathStr' is determined to be
//		an absolute path, this returned boolean value
//		will be set to true.
func (fHelpMolecule *fileHelperMolecule) isAbsolutePath(
	pathStr string) bool {

	if fHelpMolecule.lock == nil {
		fHelpMolecule.lock = new(sync.Mutex)
	}

	fHelpMolecule.lock.Lock()

	defer fHelpMolecule.lock.Unlock()

	errCode := 0

	errCode, _, pathStr =
		new(fileHelperElectron).
			isStringEmptyOrBlank(pathStr)

	if errCode < 0 {
		return false
	}

	return fp.IsAbs(pathStr)
}

// isAbsolutePathByCompare
//
// Compares the input parameter 'pathStr' to the absolute
// path representation for 'pathStr' to determine whether
// 'pathStr' represents an absolute path.
//
// This method differs from isAbsolutePath() in that
// this method does NOT call filePath.IsAbs(). Instead,
// this method constructs an absolute path from 'pathStr'
// and compares the two paths.
//
// ----------------------------------------------------------------
//
// Absolute Path Definition (Wikipedia):
//
//	An absolute or full path points to the same location
//	in a file system, regardless of the current working
//	directory. To do that, it must include the root
//	directory.
//
//	By contrast, a relative path starts from some given
//	working directory, avoiding the need to provide the
//	full absolute path. A filename can be considered as
//	a relative path based at the current working
//	directory. If the working directory is not the file's
//	parent directory, a file not found error will result
//	if the file is addressed by its name.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathStr						string
//
//		This string holds the file path which will be
//		analyzed to determine if it is an absolute path.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		This method will analyze input parameter
//		'pathStr' to determine if it is an absolute path.
//
//		This determination is made by comparing 'pathStr'
//		to the absolute path constructed with 'pathStr'.
//
//		If input parameter '' is determined to be an
//		absolute path, this returned boolean value will be
//		set to true.
func (fHelpMolecule *fileHelperMolecule) isAbsolutePathByCompare(
	pathStr string) bool {

	if fHelpMolecule.lock == nil {
		fHelpMolecule.lock = new(sync.Mutex)
	}

	fHelpMolecule.lock.Lock()

	defer fHelpMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"fileHelperMolecule."+
			"isAbsolutePathByCompare()",
		"")

	if err != nil {
		return false
	}

	errCode := 0

	errCode, _, pathStr =
		new(fileHelperElectron).
			isStringEmptyOrBlank(pathStr)

	if errCode < 0 {
		return false
	}

	// Adjust the path separators for the current operating
	// system.
	correctDelimPathStr := strings.ToLower(
		new(fileHelperAtom).adjustPathSlash(pathStr))

	correctDelimPathStr =
		strings.TrimLeft(correctDelimPathStr, " ")

	correctDelimPathStr =
		strings.TrimRight(correctDelimPathStr, " ")

	absPath, err := new(fileHelperProton).makeAbsolutePath(
		pathStr,
		ePrefix)

	if err != nil {
		return false
	}

	absPath = strings.ToLower(absPath)

	absPath =
		strings.TrimLeft(absPath, " ")

	absPath =
		strings.TrimRight(absPath, " ")

	if absPath == correctDelimPathStr {
		return true
	}

	return false
}

// joinPathsAdjustSeparators
//
// Joins two path strings and standardizes the path
// separators according to the current operating system.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	p1							string
//
//		This string holds one of two path strings which
//		will be joined together in tandem (p1+p2). 'p1'
//		will be located at the beginning of the
//		composite, joined path string returned to the
//		calling function.
//
//	p2							string
//
//		This string holds the second of two path strings
//		which will be joined together in tandem (p1+p2).
//		'p2' will be located at the end of the composite,
//		joined path string returned to the calling
//		function.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	string
//
//		This returned path string contains the two joined
//		path strings provided by input parameter 'p1' and
//		'p2' (p1+p2).
func (fHelpMolecule *fileHelperMolecule) joinPathsAdjustSeparators(
	p1 string, p2 string) string {

	if fHelpMolecule.lock == nil {
		fHelpMolecule.lock = new(sync.Mutex)
	}

	fHelpMolecule.lock.Lock()

	defer fHelpMolecule.lock.Unlock()

	errCode := 0

	fHelpElectron := new(fileHelperElectron)

	errCode, _, p1 = fHelpElectron.isStringEmptyOrBlank(p1)

	if errCode < 0 {
		p1 = ""
	}

	errCode, _, p2 = fHelpElectron.isStringEmptyOrBlank(p2)

	if errCode < 0 {
		p2 = ""
	}

	if p1 == "" &&
		p2 == "" {

		return ""
	}

	fHelperAtom := new(fileHelperAtom)

	ps1 := fHelperAtom.adjustPathSlash(fp.Clean(p1))

	ps2 := fHelperAtom.adjustPathSlash(fp.Clean(p2))

	joinedPath := fHelperAtom.
		adjustPathSlash(
			path.Join(ps1, ps2))

	return fp.Clean(joinedPath)
}

// makeFileHelperWalkDirDeleteFilesFunc
//
// Used in conjunction with DirMgr.DeleteWalDirFiles to
// select and delete files residing the directory tree
// identified by the current DirMgr object.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dInfo				*DirectoryDeleteFileInfo
//
//		A pointer to an instance of DirectoryDeleteFileInfo.
//
//		DirectoryDeleteFileInfo
//		This structure is used to delete files in a
//		directory specified	by 'StartPath'. Deleted files
//		will be selected based on 'DeleteFileSelectCriteria'
//		value.
//
//		'DeleteFileSelectCriteria' is a 'FileSelectionCriteria'
//		type which contains FileNamePatterns strings and the
//		FilesOlderThan or FilesNewerThan date time parameters
//		which can be used as file selection criteria.
//
//		type DirectoryDeleteFileInfo struct {
//			StartPath                string
//			Directories              DirMgrCollection
//			ErrReturns               []error
//			DeleteFileSelectCriteria FileSelectionCriteria
//			DeletedFiles             FileMgrCollection
//		}
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	This method returns a function with the following
//	signature:
//
//		func(string, os.FileInfo, error) error
func (fHelpMolecule *fileHelperMolecule) makeFileHelperWalkDirDeleteFilesFunc(
	dInfo *DirectoryDeleteFileInfo) func(string, os.FileInfo, error) error {

	if fHelpMolecule.lock == nil {
		fHelpMolecule.lock = new(sync.Mutex)
	}

	fHelpMolecule.lock.Lock()

	defer fHelpMolecule.lock.Unlock()

	return func(pathFile string, info os.FileInfo, erIn error) error {

		if erIn != nil {
			dInfo.ErrReturns = append(dInfo.ErrReturns, erIn)
			return nil
		}

		var err error
		var ePrefix *ePref.ErrPrefixDto

		ePrefix,
			err = ePref.ErrPrefixDto{}.NewIEmpty(
			nil,
			"fileHelperMolecule."+
				"makeFileHelperWalkDirDeleteFilesFunc()",
			"")

		if erIn != nil {
			dInfo.ErrReturns = append(dInfo.ErrReturns, err)
			return nil
		}

		if info.IsDir() {

			var subDir DirMgr

			subDir, err = new(DirMgr).New(pathFile)

			if err != nil {

				ex := fmt.Errorf("%v\n"+
					"Error returned from DirMgr{}.New(pathFile).\n"+
					"pathFile:='%v'\n"+
					"Error= \n%v\n",
					ePrefix.String(),
					pathFile,
					err.Error())

				dInfo.ErrReturns = append(dInfo.ErrReturns, ex)

				return nil
			}

			subDir.actualDirFileInfo, err =
				new(FileInfoPlus).NewFromPathFileInfo(pathFile, info)

			if err != nil {

				ex := fmt.Errorf("%v\n"+
					"Error returned by FileInfoPlus{}.NewFromPathFileInfo(pathFile, info)\n"+
					"pathFile='%v'\n"+
					"info.Name='%v'\n"+
					"Error= \n%v\n",
					ePrefix.String(),
					pathFile,
					info.Name(),
					err.Error())

				dInfo.ErrReturns = append(dInfo.ErrReturns, ex)

			} else {

				dInfo.Directories.AddDirMgr(subDir)
			}

			return nil
		}

		var isFoundFile bool

		isFoundFile,
			err,
			_ = new(fileHelperAtom).
			filterFileName(
				info,
				dInfo.DeleteFileSelectCriteria,
				ePrefix)

		if err != nil {

			ex := fmt.Errorf("%v\n"+
				"Error returned from fh.FilterFileName(info, dInfo.DeleteFileSelectCriteria)\n"+
				"pathFile='%v'\n"+
				"info.Name()='%v'\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				pathFile,
				info.Name(),
				err.Error())

			dInfo.ErrReturns = append(dInfo.ErrReturns, ex)

			return nil
		}

		if isFoundFile {

			err = os.Remove(pathFile)

			if err != nil {

				ex := fmt.Errorf("%v\n"+
					"Error returned from os.Remove(pathFile).\n"+
					"pathFile='%v'\n"+
					"Error= \n%v\n",
					ePrefix.String(),
					pathFile,
					err.Error())

				dInfo.ErrReturns = append(dInfo.ErrReturns, ex)

				return nil
			}

			err = dInfo.DeletedFiles.AddFileMgrByFileInfo(
				pathFile,
				info,
				ePrefix.XCpy(
					"dInfo.DeletedFiles<-pathFile"))

			if err != nil {

				ex := fmt.Errorf("fileHelperMolecule.makeFileHelperWalkDirFindFilesFunc()\n"+
					"Error returned from dInfo.DeletedFiles.AddFileMgrByFileInfo( pathFile,  info).\n"+
					"pathFile='%v'\n"+
					"Error=\n%v\n",
					pathFile,
					err.Error())

				dInfo.ErrReturns = append(dInfo.ErrReturns, ex)

				return nil
			}

		}

		return nil
	}
}

// makeFileHelperWalkDirFindFilesFunc
//
// This function is designed to work in conjunction with
// a walk directory function like FindWalkDirFiles. It
// will process files extracted from a 'Directory Walk'
// operation initiated by the 'filepath.Walk' method.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dInfo						*DirectoryTreeInfo
//
//	A pointer to a DirectoryTreeInfo instance. This
//	structure is used to 'Find' files in a directory
//	specified by member variable, 'StartPath'. The file
//	search will be filtered by using member variable
//	'FileSelectCriteria' selection criteria
//	specifications.
//
//	'FileSelectCriteria' is a FileSelectionCriteria type
//	which contains FileNamePatterns strings and
//	'FilesOlderThan' or 'FilesNewerThan' date time
//	parameters which can be used as a selection
//	criteria.
//
//		type DirectoryTreeInfo struct {
//			StartPath          string
//			Directories        DirMgrCollection
//			FoundFiles         FileMgrCollection
//			ErrReturns         []error
//			FileSelectCriteria FileSelectionCriteria
//		}
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	This method returns a function with the following
//	signature:
//
//		func(string, os.FileInfo, error) error
func (fHelpMolecule *fileHelperMolecule) makeFileHelperWalkDirFindFilesFunc(
	dInfo *DirectoryTreeInfo) func(string, os.FileInfo, error) error {

	if fHelpMolecule.lock == nil {
		fHelpMolecule.lock = new(sync.Mutex)
	}

	fHelpMolecule.lock.Lock()

	defer fHelpMolecule.lock.Unlock()

	return func(pathFile string, info os.FileInfo, erIn error) error {

		var err, er2, ex2 error

		var ePrefix *ePref.ErrPrefixDto

		ePrefix,
			err = ePref.ErrPrefixDto{}.NewIEmpty(
			nil,
			"DirMgr."+
				"MakeFileHelperWalkDirFindFilesFunc()",
			"")

		if err != nil {
			dInfo.ErrReturns = append(dInfo.ErrReturns, err)
			return nil
		}

		if erIn != nil {

			ex2 = fmt.Errorf("%v\n"+
				"Error returned from directory walk function.\n"+
				"pathFile= '%v'\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				pathFile,
				erIn.Error())

			dInfo.ErrReturns = append(dInfo.ErrReturns, ex2)

			return nil
		}

		if info.IsDir() {

			var subDir DirMgr

			subDir, err = new(DirMgr).NewFromFileInfo(pathFile, info)

			if err != nil {

				er2 = fmt.Errorf("%v\n"+
					"Error returned by DirMgr{}.New(pathFile).\n"+
					"pathFile='%v'\n"+
					"Error= \n%v\n",
					ePrefix.String(),
					pathFile,
					err.Error())

				dInfo.ErrReturns = append(dInfo.ErrReturns, er2)

				return nil
			}

			dInfo.Directories.AddDirMgr(subDir)

			return nil
		}

		// This is not a directory. It is a file.
		// Determine if it matches the find file criteria.
		var isFoundFile bool

		isFoundFile,
			err,
			ex2 = new(fileHelperAtom).filterFileName(
			info,
			dInfo.FileSelectCriteria,
			ePrefix)

		if err != nil {

			er2 = fmt.Errorf("%v\n"+
				"Error returned from dMgr.FilterFileName(info, "+
				"dInfo.FileSelectCriteria)\n"+
				"pathFile='%v'\ninfo.Name()='%v'\n"+
				"Message Error= \n%v\n"+
				"Low Level Error= \n%v\n",
				ePrefix.String(),
				pathFile,
				info.Name(),
				err.Error(),
				ex2.Error())

			dInfo.ErrReturns = append(dInfo.ErrReturns, er2)

			return nil
		}

		if isFoundFile {

			var fMgr FileMgr

			var isEmpty bool

			isEmpty,
				er2 = new(fileMgrHelper).setFileMgrPathFileName(
				&fMgr,
				pathFile,
				ePrefix.XCpy(
					"fMgr<-pathFile"))

			if er2 != nil {

				err = fmt.Errorf("fileHelperMolecule.makeFileHelperWalkDirFindFilesFunc()\n"+
					"Error returned by fileMgrHelper.setFileMgrPathFileName(pathFile)\n"+
					"pathFile='%v'\n"+
					"Error= \n%v\n ",
					pathFile,
					er2.Error())

				dInfo.ErrReturns = append(dInfo.ErrReturns, err)

				return nil
			}

			if isEmpty {

				err = fmt.Errorf("%v\n"+
					"Error: fileMgrHelper.setFileMgrPathFileName(pathFile)\n"+
					"returned an Empty Result!\n"+
					"pathFile='%v'\n",
					ePrefix.String(),
					pathFile)

				dInfo.ErrReturns = append(dInfo.ErrReturns, err)

				return nil
			}

			fMgr, er2 = new(FileMgr).NewFromPathFileNameExtStr(pathFile)

			if er2 != nil {

				err = fmt.Errorf("%v\n"+
					"Error returned by FileMgr{}.NewFromPathFileNameExtStr(pathFile)\n"+
					"pathFile='%v'\n"+
					"Error= \n%v\n ",
					ePrefix.String(),
					pathFile,
					er2.Error())

				dInfo.ErrReturns = append(dInfo.ErrReturns, err)

				return nil
			}

			err = dInfo.FoundFiles.AddFileMgrByFileInfo(
				fMgr.dMgr.GetAbsolutePath(),
				info,
				ePrefix)

			if err != nil {

				er2 = fmt.Errorf("%v\n"+
					"Error returned from dInfo.FoundFiles.AddFileMgrByFileInfo(pathFile, info)\n"+
					"pathFile='%v'\ninfo.Name()='%v'\n"+
					"Error= \n%v\n",
					ePrefix.String(),
					pathFile,
					info.Name(),
					err.Error())

				dInfo.ErrReturns = append(dInfo.ErrReturns, er2)

				return nil
			}
		}

		return nil
	}
}

// stripLeadingDotSeparatorChars
//
// Strips or deletes the following characters from the
// front of path or directory names.
//
// Leading Characters To Be Removed:
//
//	(1)	" " (Space)
//
//	(2)	PathSeparator
//
//	(3)	"."
//
//	(4)	".."
//
//	(5)	"." + PathSeparator
//
//	(6)	".." + PathSeparator
//
// Removal of these characters will convert the path or
// directory name to a valid set of text characters
// suitable as input for file or directory processing
// functions.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathName					string
//
//		The path or directory name to be processed. This
//		method will strip or delete selected characters
//		from the front or right side of this string.
//		Removal of these characters will convert the path
//		or directory name to a valid set of text
//		characters suitable as input for file or
//	 	directory processing functions.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	string
//
//		The converted path or directory name form which
//		selected invalid characters have been removed.
//
//	int
//
//		The string length of the returned 'string'
//		parameter.
func (fHelpMolecule *fileHelperMolecule) stripLeadingDotSeparatorChars(
	pathName string) (
	string,
	int) {

	if fHelpMolecule.lock == nil {
		fHelpMolecule.lock = new(sync.Mutex)
	}

	fHelpMolecule.lock.Lock()

	defer fHelpMolecule.lock.Unlock()

	pathName = new(fileHelperAtom).adjustPathSlash(pathName)

	pathSeparatorStr := string(os.PathSeparator)

	space := " "
	dot := "."
	doubleDot := ".."
	dotSeparator := dot + pathSeparatorStr
	doubleDotSeparator := doubleDot + pathSeparatorStr
	strLen := len(pathName)

	if strLen == 0 {
		return pathName, 0
	}

	badChars := []string{
		doubleDotSeparator,
		dotSeparator,
		doubleDot,
		dot,
		pathSeparatorStr,
		space}

	for i := 0; i < len(badChars); i++ {

		for j := 0; j < 100; j++ {

			if !strings.HasPrefix(pathName, badChars[i]) {
				break
			}

			pathName = pathName[len(badChars[i]):]
		}

		strLen = len(pathName)

		if len(pathName) == 0 {
			break
		}
	}

	return pathName, strLen
}
