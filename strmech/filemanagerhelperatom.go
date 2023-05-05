package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"os"
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

// setFileMgrDirMgrFileName
//
// This configures a FileMgr instance based on a
// Directory Manager instance plus a string containing
// the file name and file extension.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fMgr						*FileMgr
//
//		A pointer to an instance of FileMgr. This method
//		will configure this FileMgr instance using two
//		components. First, the directory component will
//		be taken from the DirMgr instance passed as input
//		parameter 'dMgr'. Second, the file name and file
//		extension will be taken from input parameter
//	 	'fileNameExt'.
//
//	dMgr						*DirMgr
//
//		A pointer to an instance of Directory Manager
//		('DirMgr'). The directory name contained in this
//		DirMgr instance will be extracted and used to
//		construct the directory name configured in
//		'fMgr'.
//
//	fileNameExt					string
//
//		This string holds the file name and file
//		extension used to configure the FileMgr instance
//		passed as input parameter 'fMgr'.
//
//		If 'fileNameExt' is submitted as an empty string
//		or a string containing all white space, an error
//		will be returned.
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
//	isEmpty						bool
//
//		If input parameter 'fileNameExt' is submitted as
//		an empty string or a string containing all white
//		space, this boolean parameter will return a value
//		of 'true'.
//
//	err							error
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
func (fMgrHlprAtom *fileMgrHelperAtom) setFileMgrDirMgrFileName(
	fMgr *FileMgr,
	dMgr *DirMgr,
	fileNameExt string,
	errPrefDto *ePref.ErrPrefixDto) (
	isEmpty bool,
	err error) {

	if fMgrHlprAtom.lock == nil {
		fMgrHlprAtom.lock = new(sync.Mutex)
	}

	fMgrHlprAtom.lock.Lock()

	defer fMgrHlprAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileMgrHelperAtom."+
			"setFileMgrDirMgrFileName()",
		"")

	if err != nil {
		return false, err
	}

	isEmpty = true
	err = nil

	if fMgr == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fMgr' is a nil pointer!\n",
			ePrefix.String())

		return isEmpty, err
	}

	err2 := dMgr.IsDirMgrValid(ePrefix.String())

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'dMgr' is INVALID!\n"+
			"dMgr.absolutePath='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			dMgr.absolutePath,
			err2.Error())

		return isEmpty, err
	}

	errCode,
		_,
		fileNameExt :=
		new(fileHelperElectron).
			isStringEmptyOrBlank(fileNameExt)

	if errCode == -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fileNameExt' is a Zero length string!\n",
			ePrefix.String())

		return isEmpty, err
	}

	if errCode == -2 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fileNameExt' consists entirely of blank spaces!\n",
			ePrefix.String())

		return isEmpty, err
	}

	fh := new(FileHelper)

	adjustedFileNameExt,
		isFileNameEmpty,
		err2 :=
		fh.CleanFileNameExtStr(
			fileNameExt,
			ePrefix.XCpy(
				"adjustedFileNameExt<-fileNameExt"))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned from fh.CleanFileNameExtStr(fileNameExt).\n"+
			"fileNameExt='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			fileNameExt,
			err2.Error())

		return isEmpty, err
	}

	if isFileNameEmpty {

		err = fmt.Errorf("%v\n"+
			"Error: fileName returned from fh.CleanFileNameExtStr(fileNameExt)\n"+
			"is a ZERO length string!\n"+
			"fileNameExt='%v'\n",
			ePrefix.String(),
			fileNameExt)

		return isEmpty, err
	}

	fMgrHelperBoson := new(fileMgrHelperBoson)

	err = fMgrHelperBoson.
		emptyFileMgr(fMgr, ePrefix)

	if err != nil {
		return isEmpty, err
	}

	fMgr.dMgr,
		err = dMgr.CopyOut(ePrefix.XCpy("dMgr"))

	if err != nil {
		return isEmpty, err
	}

	s, fNameIsEmpty, err2 := fh.GetFileNameWithoutExt(
		adjustedFileNameExt,
		ePrefix.XCpy(
			"s<-adjustedFileNameExt"))

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"Error returned from fh.GetFileNameWithoutExt(adjustedFileNameExt).\n"+
			"adjustedFileNameExt='%v'\n"+
			"Error='%v'\n ",
			ePrefix.String(),
			adjustedFileNameExt,
			err2.Error())

		isEmpty = true

		_ = fMgrHelperBoson.
			emptyFileMgr(fMgr, ePrefix)

		return isEmpty, err
	}

	if fNameIsEmpty {

		err = fmt.Errorf("%v\n"+
			"Error: fileName returned from fh.GetFileNameWithoutExt(adjustedFileNameExt)\n"+
			"is Zero length string!\n"+
			"adjustedFileNameExt='%v'\n",
			ePrefix.String(),
			adjustedFileNameExt)

		_ = fMgrHelperBoson.
			emptyFileMgr(fMgr, ePrefix)

		isEmpty = true

		return isEmpty, err
	}

	fMgr.isFileNamePopulated = true
	fMgr.fileName = s

	s,
		extIsEmpty,
		err2 := fh.GetFileExtension(
		adjustedFileNameExt,
		ePrefix)

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"Error returned from fh.GetFileExt(fileNameAndExt).\n"+
			"fileNameAndExt='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			adjustedFileNameExt,
			err2.Error())

		isEmpty = true

		_ = fMgrHelperBoson.emptyFileMgr(
			fMgr,
			ePrefix)

		return isEmpty, err
	}

	if !extIsEmpty {
		fMgr.isFileExtPopulated = true
		fMgr.fileExt = s
	}

	if fMgr.isFileNamePopulated {
		fMgr.isFileNameExtPopulated = true
		fMgr.fileNameExt = fMgr.fileName + fMgr.fileExt
	}

	lPath := len(fMgr.dMgr.absolutePath)
	if lPath == 0 {
		fMgr.absolutePathFileName = fMgr.fileNameExt

	} else if fMgr.dMgr.absolutePath[lPath-1] == os.PathSeparator {
		fMgr.absolutePathFileName = fMgr.dMgr.absolutePath + fMgr.fileNameExt

	} else {
		fMgr.absolutePathFileName =
			fMgr.dMgr.absolutePath + string(os.PathSeparator) + fMgr.fileNameExt

	}

	lPath = len(fMgr.dMgr.path)

	if lPath == 0 {
		fMgr.originalPathFileName = fMgr.fileNameExt

	} else if fMgr.dMgr.path[lPath-1] == os.PathSeparator {
		fMgr.originalPathFileName = fMgr.dMgr.path + fMgr.fileNameExt

	} else {
		fMgr.originalPathFileName = fMgr.dMgr.path + string(os.PathSeparator) + fMgr.fileNameExt
	}

	fMgr.isAbsolutePathFileNamePopulated = true

	_,
		filePathDoesExist,
		fInfoPlus,
		nonPathError :=
		new(fileHelperMolecule).doesPathFileExist(
			fMgr.absolutePathFileName,
			PreProcPathCode.None(), // Do NOT perform pre-processing on path
			ePrefix,
			"fMgr.absolutePathFileName")

	if filePathDoesExist && nonPathError == nil {
		fMgr.doesAbsolutePathFileNameExist = true
		fMgr.actualFileInfo = fInfoPlus.CopyOut()

		err2 = fMgr.actualFileInfo.SetDirectoryPath(
			dMgr.absolutePath,
			ePrefix.XCpy(
				"dMgr.absolutePath"))

		if err2 != nil {
			isEmpty = true
			err = fmt.Errorf("%v\n"+
				"Error returned by fMgr.actualFileInfo.SetDirectoryPath(dMgr.absolutePath)\n"+
				"dMgr.absolutePath='%v'\n"+
				"%v",
				ePrefix.String(),
				dMgr.absolutePath,
				err2.Error())

			_ = fMgrHelperBoson.emptyFileMgr(
				fMgr,
				ePrefix)

			return isEmpty, err
		}

	} else {
		fMgr.doesAbsolutePathFileNameExist = false
		fMgr.actualFileInfo = FileInfoPlus{}
	}

	fMgr.isInitialized = true

	err = nil
	isEmpty = false

	return isEmpty, err
}
