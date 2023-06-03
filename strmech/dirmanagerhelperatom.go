package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"os"
	pf "path/filepath"
	"strings"
	"sync"
)

type dirMgrHelperAtom struct {
	lock *sync.Mutex
}

// copyOut - Makes a duplicate copy of input parameter
// 'dMgr' values and returns them as a new DirMgr object.
func (dMgrHlprAtom *dirMgrHelperAtom) copyOut(
	dMgr *DirMgr,
	errPrefDto *ePref.ErrPrefixDto) (
	DirMgr,
	error) {

	if dMgrHlprAtom.lock == nil {
		dMgrHlprAtom.lock = new(sync.Mutex)
	}

	dMgrHlprAtom.lock.Lock()

	defer dMgrHlprAtom.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	dOut := DirMgr{}

	funcName := "dirMgrHelperAtom." +
		"copyOut()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return dOut, err
	}

	if dMgr == nil {

		err = fmt.Errorf("%v \n"+
			"ERROR: Input paramter 'dMgr' is a nil pointer!\n",
			ePrefix.String())

		return dOut, err
	}

	err = new(dirMgrHelper).isDirMgrValid(
		dMgr,
		ePrefix.XCpy("dMgr"))

	if err != nil {

		return dOut, fmt.Errorf("%v\n"+
			"Error: Input paramter 'dMgr' is INVALID!\n"+
			"Error= \n%v\n",
			funcName,
			err.Error())
	}

	dOut.isInitialized = dMgr.isInitialized
	dOut.originalPath = dMgr.originalPath
	dOut.path = dMgr.path
	dOut.isPathPopulated = dMgr.isPathPopulated
	dOut.doesPathExist = dMgr.doesPathExist
	dOut.parentPath = dMgr.parentPath
	dOut.isParentPathPopulated = dMgr.isParentPathPopulated
	dOut.absolutePath = dMgr.absolutePath
	dOut.isAbsolutePathPopulated = dMgr.isAbsolutePathPopulated
	dOut.doesAbsolutePathExist = dMgr.doesAbsolutePathExist
	dOut.isAbsolutePathDifferentFromPath = dMgr.isAbsolutePathDifferentFromPath
	dOut.directoryName = dMgr.directoryName
	dOut.volumeName = dMgr.volumeName
	dOut.isVolumePopulated = dMgr.isVolumePopulated
	dOut.actualDirFileInfo = dMgr.actualDirFileInfo.CopyOut()

	return dOut, err
}

// doesDirectoryExist
//
// Helper method used by DirMgr to test for existence of
// directory path. In addition, this method performs
// validation on the 'DirMgr' instance.
func (dMgrHlprAtom *dirMgrHelperAtom) doesDirectoryExist(
	dMgr *DirMgr,
	preProcessCode PreProcessPathCode,
	dMgrLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	dirPathDoesExist bool,
	fInfo FileInfoPlus,
	err error) {

	if dMgrHlprAtom.lock == nil {
		dMgrHlprAtom.lock = new(sync.Mutex)
	}

	dMgrHlprAtom.lock.Lock()

	defer dMgrHlprAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"dirMgrHelperAtom."+
			"doesDirectoryExist()",
		"")

	if err != nil {
		return dirPathDoesExist, fInfo, err
	}

	dirPathDoesExist = false
	fInfo = FileInfoPlus{}
	err = nil

	if dMgr == nil {

		err = fmt.Errorf("%v\n"+
			"Error: %v pointer is 'nil'!\n",
			ePrefix.String(),
			dMgrLabel)

		return dirPathDoesExist, fInfo, err
	}

	if !dMgr.isInitialized {

		err = fmt.Errorf(
			"%v\n"+
				"Error: DirMgr is NOT Initialized.\n",
			ePrefix.String())

		return dirPathDoesExist, fInfo, err
	}

	fh := new(FileHelper)

	errCode := 0

	errCode, _, dMgr.absolutePath =
		fh.IsStringEmptyOrBlank(dMgr.absolutePath)

	if errCode == -1 {

		dMgr.isInitialized = false
		dMgr.absolutePath = ""
		dMgr.path = ""
		dMgr.doesAbsolutePathExist = false
		dMgr.doesPathExist = false
		dMgr.actualDirFileInfo = FileInfoPlus{}
		dirPathDoesExist = false

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v'.absolutePath is an empty string!\n",
			ePrefix.String(),
			dMgrLabel)

		return dirPathDoesExist, fInfo, err
	}

	if errCode == -2 {
		dMgr.isInitialized = false
		dMgr.absolutePath = ""
		dMgr.path = ""
		dMgr.doesAbsolutePathExist = false
		dMgr.doesPathExist = false
		dMgr.actualDirFileInfo = FileInfoPlus{}
		dirPathDoesExist = false

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' consists of blank spaces!\n",
			ePrefix.String(),
			dMgrLabel)

		return dirPathDoesExist, fInfo, err
	}

	var err2 error

	if preProcessCode == PreProcPathCode.PathSeparator() {

		dMgr.absolutePath = fh.AdjustPathSlash(dMgr.absolutePath)

	} else if preProcessCode == PreProcPathCode.AbsolutePath() {

		dMgr.absolutePath,
			err2 =
			fh.MakeAbsolutePath(
				dMgr.absolutePath,
				ePrefix.XCpy("dMgr.absolutePath<-"))

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Error: fh.MakeAbsolutePath(%v.absolutePath) FAILED!\n"+
				"%v.absolutePath='%v'\n"+
				"Error='%v'\n\n",
				ePrefix.String(),
				dMgrLabel,
				dMgrLabel,
				dMgr.absolutePath,
				err2.Error())

			return dirPathDoesExist, fInfo, err
		}
	}

	errCode, _, dMgr.path =
		fh.IsStringEmptyOrBlank(dMgr.path)

	if errCode < 0 {
		dMgr.path = dMgr.absolutePath
	}

	dMgr.isPathPopulated = true

	strAry := strings.Split(dMgr.absolutePath, string(os.PathSeparator))
	lStr := len(strAry)
	idxStr := strAry[lStr-1]

	idx := strings.Index(dMgr.absolutePath, idxStr)

	dMgr.parentPath = fh.RemovePathSeparatorFromEndOfPathString(dMgr.absolutePath[0:idx])

	dMgr.isParentPathPopulated = true

	if dMgr.parentPath == "" {
		dMgr.isParentPathPopulated = false
	}

	if idxStr != "" {
		dMgr.directoryName = idxStr
	} else {
		dMgr.directoryName = dMgr.absolutePath
	}

	errCode, _, dMgr.path =
		fh.IsStringEmptyOrBlank(dMgr.path)

	if dMgr.path != dMgr.absolutePath {
		dMgr.isAbsolutePathDifferentFromPath = true
	}

	var vn string
	if dMgr.isAbsolutePathPopulated {
		vn = pf.VolumeName(dMgr.absolutePath)
	} else if dMgr.isPathPopulated {
		vn = pf.VolumeName(dMgr.path)
	}

	dMgr.isVolumePopulated = false

	if vn != "" {
		dMgr.isVolumePopulated = true
		dMgr.volumeName = vn
	}

	var absFInfo, pathFInfo FileInfoPlus

	dMgr.doesAbsolutePathExist,
		absFInfo,
		err = new(dirMgrHelperElectron).
		lowLevelDoesDirectoryExist(
			dMgr.absolutePath,
			dMgrLabel+".absolutePath",
			ePrefix)

	if err != nil {
		dMgr.doesAbsolutePathExist = false
		dMgr.doesPathExist = false
		dMgr.actualDirFileInfo = FileInfoPlus{}
		dirPathDoesExist = false
		return dirPathDoesExist, fInfo, err
	}

	if !dMgr.doesAbsolutePathExist {
		dMgr.doesAbsolutePathExist = false
		dMgr.doesPathExist = false
		dMgr.actualDirFileInfo = FileInfoPlus{}
		dirPathDoesExist = false
		err = nil
		return dirPathDoesExist, fInfo, err
	}

	if !absFInfo.Mode().IsDir() {
		dMgr.doesAbsolutePathExist = false
		dMgr.doesPathExist = false
		dMgr.actualDirFileInfo = FileInfoPlus{}

		err = fmt.Errorf("%v\n"+
			"Error: Directory absolute path exists, but "+
			"it is a file - NOT A DIRECTORY!\n"+
			"%v='%v'\n",
			ePrefix.String(),
			dMgrLabel,
			dMgr.absolutePath)

		dirPathDoesExist = false

		return dirPathDoesExist, fInfo, err
	}

	if absFInfo.Mode().IsRegular() {

		dMgr.doesAbsolutePathExist = false
		dMgr.doesPathExist = false
		dMgr.actualDirFileInfo = FileInfoPlus{}

		err = fmt.Errorf("%v\n"+
			"Error: Directory absolute path exists,\n"+
			"but it is classified as as a Regular File!\n"+
			"%v='%v'\n",
			ePrefix.String(),
			dMgrLabel,
			dMgr.absolutePath)

		dirPathDoesExist = false

		return dirPathDoesExist, fInfo, err
	}

	dMgr.doesPathExist,
		pathFInfo,
		err = new(dirMgrHelperElectron).
		lowLevelDoesDirectoryExist(
			dMgr.path,
			dMgrLabel+".path",
			ePrefix)

	if err != nil {

		dMgr.doesAbsolutePathExist = false

		dMgr.doesPathExist = false

		dMgr.actualDirFileInfo = FileInfoPlus{}

		dirPathDoesExist = false

		return dirPathDoesExist, fInfo, err
	}

	if !dMgr.doesPathExist {

		err = fmt.Errorf("%v\n"+
			"Error: Directory absolute path exists,\n"+
			"but original directory 'path' DOES NOT EXIST!\n"+
			"%v.absolutePath='%v'\n"+
			"%v.path='%v'\n",
			ePrefix.String(),
			dMgrLabel,
			dMgr.absolutePath,
			dMgrLabel,
			dMgr.path)

		dMgr.doesAbsolutePathExist = false
		dMgr.doesPathExist = false
		dMgr.actualDirFileInfo = FileInfoPlus{}
		dirPathDoesExist = false

		return dirPathDoesExist, fInfo, err
	}

	if !pathFInfo.Mode().IsDir() {

		err = fmt.Errorf("%v\n"+
			"Error: Directory path absolute path exists,\n"+
			"but original directory 'path' is NOT A DIRECTORY!!\n"+
			"%v.absolutePath= '%v'\n"+
			"%v.path='%v'\n",
			ePrefix.String(),
			dMgrLabel,
			dMgr.absolutePath,
			dMgrLabel,
			dMgr.path)

		dMgr.doesAbsolutePathExist = false
		dMgr.doesPathExist = false
		dMgr.actualDirFileInfo = FileInfoPlus{}
		dirPathDoesExist = false
		return dirPathDoesExist, fInfo, err
	}

	if pathFInfo.Mode().IsRegular() {
		err = fmt.Errorf("%v\n"+
			"Error: Directory path exists,\n"+
			"but original directory 'path' is\n"+
			"classified  as a Regular File!\n"+
			"%v.absolutePath='%v'\n"+
			"%v.path='%v'\n",
			ePrefix.String(),
			dMgrLabel,
			dMgr.absolutePath,
			dMgrLabel,
			dMgr.path)

		dMgr.doesAbsolutePathExist = false
		dMgr.doesPathExist = false
		dMgr.actualDirFileInfo = FileInfoPlus{}
		dirPathDoesExist = false
		return dirPathDoesExist, fInfo, err
	}

	// both dMgr.path and dMgr.doesAbsolutePathExist
	// exist. And, there are no errors

	dMgr.actualDirFileInfo = absFInfo.CopyOut()

	if err != nil {
		return dirPathDoesExist, fInfo, err
	}

	dMgr.doesAbsolutePathExist = true
	dMgr.doesPathExist = true

	fInfo = dMgr.actualDirFileInfo.CopyOut()

	if err != nil {
		return dirPathDoesExist, fInfo, err
	}

	dirPathDoesExist = true

	return dirPathDoesExist, fInfo, err
}

// executeFileOpsOnFoundFiles
//
// This function is designed to work in conjunction with
// a walk directory function like FindWalkDirFiles. It
// will process files extracted from a 'Directory Walk'
// operation initiated by the 'filepath.Walk' method.
//
// Thereafter, file operations will be performed on files
// in the directory tree as specified by the 'dirOp'
// parameter.
func (dMgrHlprAtom *dirMgrHelperAtom) executeFileOpsOnFoundFiles(
	dirOp *DirTreeOp,
	errPrefDto *ePref.ErrPrefixDto) func(string, os.FileInfo, error) error {

	if dMgrHlprAtom.lock == nil {
		dMgrHlprAtom.lock = new(sync.Mutex)
	}

	dMgrHlprAtom.lock.Lock()

	defer dMgrHlprAtom.lock.Unlock()

	return func(pathFile string, info os.FileInfo, erIn error) error {

		var ePrefix *ePref.ErrPrefixDto

		var err error

		ePrefix,
			err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
			errPrefDto,
			"dirMgrHelper.executeFileOpsOnFoundFiles()",
			"")

		if err != nil {
			return err
		}

		if dirOp == nil {

			err = fmt.Errorf("%v \n"+
				"ERROR: Input paramter 'dirOp' is a nil pointer!\n",
				ePrefix.String())

			return err
		}

		var err2 error

		if erIn != nil {

			err2 = fmt.Errorf("%v\n"+
				"Error returned from directory walk function.\n"+
				"pathFile='%v'\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				pathFile,
				erIn.Error())

			dirOp.ErrReturns = append(dirOp.ErrReturns, err2)

			return nil
		}

		if info.IsDir() {
			return nil
		}

		fh := FileHelper{}

		// This is not a directory. It is a file.
		// Determine if it matches the find file criteria.
		isFoundFile,
			err,
			_ := fh.FilterFileName(
			info,
			dirOp.FileSelectCriteria,
			ePrefix)

		if err != nil {

			err2 = fmt.Errorf("%v\n"+
				"Error returned from dMgr.FilterFileName(info, dInfo.FileSelectCriteria)\n"+
				"pathFile='%v'\n"+
				"info.Name()='%v'\n"+
				"Error='%v'\n",
				ePrefix.String(),
				pathFile, info.Name(),
				err.Error())

			dirOp.ErrReturns = append(dirOp.ErrReturns, err2)
			return nil
		}

		if !isFoundFile {
			return nil
		}

		srcFileNameExt := info.Name()

		destDir, err := fh.SwapBasePath(
			dirOp.SourceBaseDir.absolutePath,
			dirOp.TargetBaseDir.absolutePath,
			pathFile,
			ePrefix)

		if err != nil {

			err2 = fmt.Errorf("%v\n"+
				"Error returned by fh.SwapBasePath(dirOp.SourceBaseDir, "+
				"dirOp.TargetBaseDir, pathFile).\n"+
				"dirOp.SourceBaseDir='%v'\n"+
				"dirOp.TargetBaseDir='%v'\n"+
				"pathFile='%v'\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				dirOp.SourceBaseDir.absolutePath,
				dirOp.TargetBaseDir.absolutePath,
				pathFile,
				err.Error())

			dirOp.ErrReturns = append(dirOp.ErrReturns, err2)
			return nil
		}

		fileOp,
			err := new(FileOps).
			NewByDirStrsAndFileNameExtStrs(
				pathFile,
				srcFileNameExt,
				destDir,
				srcFileNameExt,
				ePrefix.XCpy("fileOp<-"))

		if err != nil {
			err2 = fmt.Errorf("%v\n"+
				"Error returned by FileOps{}.NewByDirStrsAndFileNameExtStrs(pathFile, "+
				"srcFileNameExt, destDir, srcFileNameExt)\n"+
				"pathFile='%v'\n"+
				"srcFileNameExt='%v'\n"+
				"destDir='%v'\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				pathFile,
				srcFileNameExt,
				destDir,
				err.Error())

			dirOp.ErrReturns = append(dirOp.ErrReturns, err2)
			return nil
		}

		for i := 0; i < len(dirOp.FileOps); i++ {

			err = fileOp.ExecuteFileOperation(
				dirOp.FileOps[i],
				ePrefix)

			if err != nil {
				err2 = fmt.Errorf("%v\n"+
					"Error returned by fileOp.ExecuteFileOperation(dirOp.FileOps[i]).\n"+
					"i='%v'\n"+
					"FileOps='%v'\n"+
					"Error= \n%v\n",
					ePrefix.String(),
					i, dirOp.FileOps[i].String(),
					err.Error())

				dirOp.ErrReturns = append(dirOp.ErrReturns, err2)

			}
		}

		return nil
	}
}

// lowLevelScreenPathStrForInvalidChars
//
// Examines input parameter 'pathStr' to determine if it
// contains invalid characters.
//
// If 'pathStr' evaluates as 'valid', lead and trailing
// spaces are deleted and the valid path string is
// returned through parameter 'validPathStr'.
//
// If 'pathStr' is determined to contain invalid
// characters, an error is returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathStr						string
//
//		This string contains the path to be validated.
//
//	pathStrLabel				string
//
//		The name or label associated with input parameter
//		'pathStr', which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "pathStr" will be
//		automatically applied.
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
//	validPathStr				string
//
//		If input parameter 'pathStr' is determined to be
//		a valid path, leading and trailing spaces in
//		'pathStr' will be deleted and returned through
//		'validPathStr'.
//
//	validPathStrLength			int
//
//		This returned integer value specifies the length of
//		'validPathStr'
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
func (dMgrHlprAtom *dirMgrHelperAtom) lowLevelScreenPathStrForInvalidChars(
	pathStr string,
	pathStrLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	validPathStr string,
	validPathStrLength int,
	err error) {

	if dMgrHlprAtom.lock == nil {
		dMgrHlprAtom.lock = new(sync.Mutex)
	}

	dMgrHlprAtom.lock.Lock()

	defer dMgrHlprAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"dirMgrHelperAtom."+
			"lowLevelScreenPathStrForInvalidChars()",
		"")

	if err != nil {

		return validPathStr, validPathStrLength, err
	}

	strLen := 0

	if len(pathStrLabel) == 0 {
		pathStrLabel = "pathStr"
	}

	pathStr,
		strLen,
		err = new(dirMgrHelperElectron).
		isPathStringEmptyOrBlank(
			pathStr,
			true, // trim trailing path separator
			pathStrLabel,
			ePrefix)

	if err != nil {

		return validPathStr, validPathStrLength, err
	}

	tripleDotSeparator := "..."
	doublePathSeparator := string(os.PathSeparator) + string(os.PathSeparator)

	if strings.Contains(pathStr, tripleDotSeparator) {

		err = fmt.Errorf("%v"+
			"ERROR: Input parameter '%v' contains invalid dot characters!\n"+
			"%v = %v\n",
			ePrefix.String(),
			pathStrLabel,
			pathStrLabel,
			pathStr)

		return validPathStr, validPathStrLength, err

	}

	if strings.Contains(pathStr, doublePathSeparator) {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter '%v' contains invalid path separator characters!\n"+
			"%v = %v\n",
			ePrefix.String(),
			pathStrLabel,
			pathStrLabel,
			pathStr)

		return validPathStr, validPathStrLength, err
	}

	validPathStr = pathStr
	validPathStrLength = strLen
	err = nil

	return validPathStr, validPathStrLength, err
}
