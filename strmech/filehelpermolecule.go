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
	var err error

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

	if len(filePathTitle) == 0 {

		filePathTitle = "filePath"
	}

	errCode := 0

	errCode, _, filePath = new(fileHelperElectron).
		isStringEmptyOrBlank(filePath)

	if errCode == -1 {

		nonPathError =
			fmt.Errorf("%v\n"+
				"Error: Input parameter '%v' is an empty string!\n",
				ePrefix.String(),
				filePathTitle)

		return absFilePath, filePathDoesExist, fInfo, nonPathError
	}

	if errCode == -2 {

		nonPathError =
			fmt.Errorf("%v\n"+
				"Error: Input parameter '%v' consists of blank spaces!",
				ePrefix.String(),
				filePathTitle)

		return absFilePath, filePathDoesExist, fInfo, nonPathError
	}

	if preProcessCode == PreProcPathCode.PathSeparator() {

		absFilePath = new(fileHelperAtom).adjustPathSlash(filePath)

	} else if preProcessCode == PreProcPathCode.AbsolutePath() {

		absFilePath, err = new(fileHelperProton).
			makeAbsolutePath(
				filePath,
				ePrefix.XCpy(
					"filePath"))

		if err != nil {

			absFilePath = ""
			nonPathError =
				fmt.Errorf("%v\n"+
					"fh.MakeAbsolutePath() FAILED!\n"+
					"%v",
					ePrefix.String(),
					err.Error())

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

		info, err = os.Stat(absFilePath)

		if err != nil {

			if os.IsNotExist(err) {

				filePathDoesExist = false
				fInfo = FileInfoPlus{}
				nonPathError = nil
				return absFilePath, filePathDoesExist, fInfo, nonPathError
			}

			// err == nil and err != os.IsNotExist(err)
			// This is a non-path error. The non-path error will be
			// tested up to 3-times before it is returned.
			nonPathError =
				fmt.Errorf("%v\n"+
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

			fInfo = FileInfoPlus{}.NewFromFileInfo(info)

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

			subDir, err = DirMgr{}.New(pathFile)

			if err != nil {

				ex := fmt.Errorf("%v\n"+
					"Error returned from DirMgr{}.New(pathFile).\n"+
					"pathFile:='%v'\nError='%v'\n",
					ePrefix.String(),
					pathFile,
					err.Error())

				dInfo.ErrReturns = append(dInfo.ErrReturns, ex)

				return nil
			}

			subDir.actualDirFileInfo, err = FileInfoPlus{}.NewFromPathFileInfo(pathFile, info)

			if err != nil {

				ex := fmt.Errorf("%v\n"+
					"Error returned by FileInfoPlus{}.NewFromPathFileInfo(pathFile, info)\n"+
					"pathFile='%v'\n"+
					"info.Name='%v'\n"+
					"Error='%v'\n",
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
				"Error='%v'\n",
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
					"pathFile='%v'\nError='%v'\n",
					ePrefix.String(),
					pathFile,
					err.Error())

				dInfo.ErrReturns = append(dInfo.ErrReturns, ex)

				return nil
			}

			err = dInfo.DeletedFiles.AddFileMgrByFileInfo(pathFile, info)

			if err != nil {

				ex := fmt.Errorf("%v\n"+
					"Error returned from dInfo.DeletedFiles.AddFileMgrByFileInfo( pathFile,  info).\n"+
					"pathFile='%v'\nError='%v'\n",
					ePrefix.String(),
					pathFile,
					err.Error())

				dInfo.ErrReturns = append(dInfo.ErrReturns, ex)

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
