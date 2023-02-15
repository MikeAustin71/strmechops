package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"os"
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
	filePathTitle string,
	errPrefDto *ePref.ErrPrefixDto) (
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
		nonPathError = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileHelperAtom."+
			"doesPathFileExist()",
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
			// This is a non-path error. The non-path error will be test
			// up to 3-times before it is returned.
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
