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
	dMgr.doesAbsolutePathExist = true
	dMgr.doesPathExist = true
	fInfo = dMgr.actualDirFileInfo.CopyOut()
	dirPathDoesExist = true
	err = nil

	return dirPathDoesExist, fInfo, err
}
