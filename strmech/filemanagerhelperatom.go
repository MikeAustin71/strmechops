package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// fileMgrHelperAtom
//
// Provides helper methods for Type
// fileMgrHelper.
type fileMgrHelperAtom struct {
	lock *sync.Mutex
}

// doesFileMgrPathFileExist
//
// Used by FileMgr type to test for the existence of a
// path and file name. In addition, this method performs
// validation on the 'FileMgr' instance.
func (fMgrHlprAtom *fileMgrHelperAtom) doesFileMgrPathFileExist(
	fileMgr *FileMgr,
	preProcessCode PreProcessPathCode,
	errPrefDto *ePref.ErrPrefixDto,
	filePathTitle string) (
	filePathDoesExist bool,
	nonPathError error) {

	if fMgrHlprAtom.lock == nil {
		fMgrHlprAtom.lock = new(sync.Mutex)
	}

	fMgrHlprAtom.lock.Lock()

	defer fMgrHlprAtom.lock.Unlock()

	filePathDoesExist = false

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileMgrHelper.doesFileMgrPathFileExist()"

	ePrefix,
		nonPathError = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if nonPathError != nil {
		return filePathDoesExist, nonPathError
	}

	if len(filePathTitle) == 0 {
		filePathTitle = "filePath"
	}

	if fileMgr == nil {

		nonPathError = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fMgr' is a nil pointer!\n",
			ePrefix.String())

		return filePathDoesExist, nonPathError
	}

	errCode := 0

	errCode,
		_, fileMgr.absolutePathFileName =
		new(fileHelperElectron).isStringEmptyOrBlank(fileMgr.absolutePathFileName)

	if errCode == -1 {

		fileMgr.isAbsolutePathFileNamePopulated = false

		nonPathError = fmt.Errorf("%v\n"+
			"Error: '%v' is an empty string!\n",
			ePrefix.String(),
			filePathTitle)

		return filePathDoesExist, nonPathError
	}

	if errCode == -2 {

		fileMgr.isAbsolutePathFileNamePopulated = false

		nonPathError = fmt.Errorf("%v\n"+
			"\nError: '%v' consists of blank spaces!\n",
			ePrefix.String(),
			filePathTitle)

		return filePathDoesExist, nonPathError
	}

	if !fileMgr.isInitialized {

		nonPathError = fmt.Errorf("%v\n"+
			"Error: This data structure is NOT initialized.\n"+
			"fileMgr.isInitialized='false'\n",
			ePrefix.String())

		return filePathDoesExist, nonPathError
	}

	var err2, err3 error

	err2 = fileMgr.dMgr.IsDirMgrValid(ePrefix.String())

	if err2 != nil {
		nonPathError = fmt.Errorf("\nFileMgr Directory Manager INVALID!\n"+
			"\nError='%v'", err2.Error())
		return filePathDoesExist, nonPathError
	}

	if preProcessCode == PreProcPathCode.PathSeparator() {

		fileMgr.absolutePathFileName =
			new(FileHelper).
				AdjustPathSlash(fileMgr.absolutePathFileName)

	} else if preProcessCode == PreProcPathCode.AbsolutePath() {

		fileMgr.absolutePathFileName, err2 =
			new(FileHelper).MakeAbsolutePath(
				fileMgr.absolutePathFileName,
				ePrefix)

		if err2 != nil {

			nonPathError = fmt.Errorf("%v\n"+
				"FileHelper{}.MakeAbsolutePath() FAILED!\n"+
				"%v='%v'\n"+
				"%v\n",
				ePrefix.String(),
				filePathTitle,
				fileMgr.absolutePathFileName,
				err2.Error())

			return filePathDoesExist, nonPathError
		}
	}

	var info FileInfoPlus

	filePathDoesExist,
		info,
		nonPathError =
		new(fileMgrHelperElectron).lowLevelDoesFileExist(
			fileMgr.absolutePathFileName,
			fileMgr.dMgr.absolutePath,
			ePrefix,
			"fileMgr.absolutePathFileName",
			"fileMgr.dMgr.absolutePath")

	if nonPathError != nil {
		fileMgr.doesAbsolutePathFileNameExist = false
		fileMgr.actualFileInfo = FileInfoPlus{}
		filePathDoesExist = false
		fileMgr.fileAccessStatus.Empty()
		return filePathDoesExist, nonPathError
	}

	if !filePathDoesExist {
		fileMgr.doesAbsolutePathFileNameExist = false
		fileMgr.actualFileInfo = FileInfoPlus{}
		fileMgr.fileAccessStatus.Empty()
		filePathDoesExist = false
		nonPathError = nil
		_ = fileMgr.dMgr.DoesPathExist()
		_ = fileMgr.dMgr.DoesAbsolutePathExist()
		return filePathDoesExist, nonPathError
	}

	// The path really does exist!
	errs := make([]error, 0, 10)
	filePathDoesExist = true
	nonPathError = nil
	fileMgr.doesAbsolutePathFileNameExist = true

	fileMgr.actualFileInfo,
		err2 =
		new(FileInfoPlus).
			NewFromPathFileInfo(
				fileMgr.dMgr.absolutePath,
				&info)

	if err2 != nil {

		err3 = fmt.Errorf("%v\n"+
			"Error returned by FileInfoPlus{}.NewFromPathFileInfo(fileMgr.dMgr.absolutePath, info)\n"+
			"fileMgr.dMgr.absolutePath='%v'\n"+
			"info.Name()='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			fileMgr.dMgr.absolutePath,
			info.Name(),
			err2.Error())

		errs = append(errs, err3)

	} else {

		permCode, err2 := new(FilePermissionConfig).
			NewByFileMode(
				fileMgr.actualFileInfo.Mode(),
				ePrefix.XCpy(
					"permCode<-fileMgr"))

		if err2 != nil {
			err3 = fmt.Errorf("%v\n"+
				"Error returned by FilePermissionConfig{}.NewByFileMode(fileMgr.actualFileInfo.Mode())\n"+
				"Error= \n%v\n",
				funcName,
				err2.Error())

			errs = append(errs, err3)

		} else {

			err2 = fileMgr.
				fileAccessStatus.
				SetFilePermissionCodes(
					permCode,
					ePrefix.XCpy(
						"fileMgr.fileAccessStatus->"))

			if err2 != nil {
				err3 = fmt.Errorf("%v\n"+
					"Error returned by fileMgr.fileAccessStatus.SetFilePermissionCodes(permCode)\n"+
					"Error='%v'\n",
					funcName,
					err2.Error())

				errs = append(errs, err3)
			}

		}
	}

	_ = fileMgr.dMgr.DoesPathExist()

	_ = fileMgr.dMgr.DoesAbsolutePathExist()

	if len(errs) > 0 {
		nonPathError = new(StrMech).ConsolidateErrors(errs)
	}

	return filePathDoesExist, nonPathError
}
