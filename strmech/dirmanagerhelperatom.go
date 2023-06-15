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

	funcName := "dirMgrHelperAtom." +
		"doesDirectoryExist()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return dirPathDoesExist, fInfo, err
	}

	dirPathDoesExist = false
	fInfo = FileInfoPlus{}
	err = nil

	if len(dMgrLabel) == 0 {
		dMgrLabel = "dMgr"
	}

	if dMgr == nil {

		err = fmt.Errorf("%v\n"+
			"Error: %v pointer is 'nil'!\n",
			ePrefix.String(),
			dMgrLabel)

		return dirPathDoesExist, fInfo, err
	}

	fh := new(FileHelper)

	errCode := 0

	errCode,
		_,
		dMgr.absolutePath =
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
			"Error: This Directory Manager instance is invalid!\n"+
			"Input parameter '%v'.absolutePath is an empty string.\n",
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
			"Error: This Directory Manager instance is invalid!\n"+
			"Input parameter '%v' consists of blank spaces.\n",
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
				"Error: This Directory Manager instance is invalid!\n"+
				"fh.MakeAbsolutePath(%v.absolutePath) FAILED.\n"+
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

	errCode,
		_,
		dMgr.path =
		fh.IsStringEmptyOrBlank(dMgr.path)

	if errCode < 0 {
		dMgr.path = dMgr.absolutePath
	}

	dMgr.isPathPopulated = true

	strAry := strings.Split(dMgr.absolutePath, string(os.PathSeparator))
	lStr := len(strAry)
	idxStr := strAry[lStr-1]

	idx := strings.Index(dMgr.absolutePath, idxStr)

	dMgr.parentPath =
		fh.RemovePathSeparatorFromEndOfPathString(
			dMgr.absolutePath[0:idx])

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
			"This Directory Manager instance is invalid!\n"+
			"Directory absolute path exists, but "+
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
			"Error: This Directory Manager instance is invalid!\n"+
			"The Directory absolute path exists, but it is\n"+
			"classified as as a Regular File.\n"+
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

		if info == nil {

			err2 = fmt.Errorf("%v\n"+
				"Error: Input parameter 'info' is nil and Invalid!\n",
				ePrefix.String())

			dirOp.ErrReturns = append(dirOp.ErrReturns, err2)

			return nil
		}

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

		fh := new(FileHelper)

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

// moveDirectoryFiles
//
// This method will 'move' files from a source directory
// to a target directory.
//
// This move operation is accomplished in two steps by
// first copying selected source files to the target
// directory and then deleting the original source
// files. The copy step is executed using a Copy IO
// operation. For information on the Copy IO procedure
// see FileHelper{}.CopyFileByIo() method and reference:
//
//	https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// No subdirectories will be included in this move
// operation. Only source files in a single directory
// will be moved to the target directory.
//
// To qualify as a selected source file eligible for the
// move operation, a file must comply with two filters:
// File Type and File Characteristics.
//
// To be eligible for the move operation, a source file
// must first comply with the specified File Type
// criteria. The File Type Filter classifies artifacts
// residing in a parent directory as subdirectories,
// regular files, SymLink files or other non-regular
// files. Since this method does NOT move subdirectories,
// the only valid file types are Regular Files, SymLink
// Files and Other Non-Regular Files. For an explanation
// of Regular and Non-Regular files, see the Definition
// of Terms section below.
//
// Screening criteria for File Type is controlled by the
// following three input parameters:
//
//	moveRegularFiles bool
//	moveSymLinkFiles bool
//	moveOtherNonRegularFiles bool
//
// File Types eligible for this move operation therefore
// include Regular Files such as text files, image files
// and executable files, SymLink files and other Non-Regular
// Files such as device files, named pipes and sockets.
//
// In addition to File Type, selected files must also
// comply with the File Characteristics criteria
// specified by input parameter, 'fileSelectCriteria'.
// The File Characteristics Selection criteria allows
// users to screen files for File Name, File Modification
// Date and File Mode.
//
// ----------------------------------------------------------------
//
// # Definition Of Terms
//
//	Regular & Non-Regular Files
//
//	In Go programming language, a regular file is a file
//	that contains data in any format that can be read by
//	a user or an application. It is not a directory or a
//	device file.
//
//	Regular files include text files, image files and
//	executable files.
//
//	Non-regular files include directories, device files,
//	named pipes, sockets, and symbolic links.
//
//	https://docs.studygolang.com/src/io/fs/fs.go
//	https://go.dev/src/os/types.go
//	https://go.dev/src/os/types.go?s=1237:1275#L31
//	https://pkg.go.dev/gopkg.in/src-d/go-git.v4/plumbing/filemode
//	https://www.linode.com/docs/guides/creating-reading-and-writing-files-in-go-a-tutorial/
//
//	SymLink Files
//
//	In computing, a symbolic link (also symlink or soft
//	link) is a file whose purpose is to point to a file
//	or directory (called the "target") by specifying a
//	path thereto.
//
//		https://en.wikipedia.org/wiki/Symbolic_link
//
//	It's true that a symlink is a shortcut file. But it's
//	different from a standard shortcut that a program
//	installer might place on your Windows desktop to make
//	the program easier to run.
//
//	Clicking on either type of shortcut opens the linked
//	object. However, what goes on beneath the hood is
//	different in both cases.
//
//	While a standard shortcut points to a certain object,
//	a symlink makes it appear as if the linked object is
//	actually there. Your computer and the apps on it will
//	read the symlink as the target object itself.
//
//		https://www.thewindowsclub.com/create-symlinks-in-windows-10
//		https://www.makeuseof.com/tag/what-is-a-symbolic-link-what-are-its-uses-makeuseof-explains/
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	(1)	This method ONLY moves files from the single
//		parent directory identified by source directory
//		'sourceDMgr' to the target directory identified
//		by 'targetDMgr'.
//
//	(2)	This method does NOT move subdirectories residing
//		in 'sourceDMgr' and therefore, it does NOT move
//		files residing in subdirectories of 'sourceDMgr'.
//
//	(3)	If the target directory does not exist, this method
//		will attempt to create it.
//
//	(4)	Files will only be moved if they meet the File Type
//		criteria and the File Characteristics Criteria.
//
//		File Type criteria are specified by input parameters:
//
//			moveRegularFiles bool
//			moveSymLinkFiles bool
//			moveOtherNonRegularFiles bool
//
//		File Characteristics Selection criteria is specified by
//		input parameter 'fileSelectCriteria'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	sourceDMgr					*DirMgr
//
//		An instance of DirMgr which identifies the source
//		directory from which files will be moved to a target
//		directory identified by input parameter 'targetDMgr'.
//
//		The move operation is accomplished in two steps.
//		First, a file is copied from 'sourceDMgr' to
//		'targetDMgr'. Then, the original file in
//		'sourceDMgr' is deleted.
//
//	targetDMgr					*DirMgr
//
//		An instance of DirMgr which identifies the
//		destination directory to which files from the
//		directory specified 'sourceDMgr' will be moved.
//
//		The move operation is accomplished in two steps.
//		First, a file is copied from 'sourceDMgr' to
//		'targetDMgr'. Then, the original file in
//		'sourceDMgr' is deleted.
//
//	returnMovedFilesList		bool
//
//		If input parameter 'returnMovedFilesList' is set
//		to 'true', this method will return a populated
//		File Manager Collection documenting all the files
//		actually included in the move operation.
//
//		If input parameter 'returnMovedFilesList' is set
//		to 'false', this method will return an empty and
//		unpopulated instance of FileMgrCollection. This
//		means that the files actually moved from the
//		source directory to the target directory by this
//		method, will NOT be documented.
//
//	copyEmptyTargetDirectory	bool
//
//		If set to 'true' the target directory will be
//		created regardless of whether any files are
//		selected to be moved to that directory. Remember
//		that files are only moved to the target directory
//		if they meet the File Type and File Characteristics
//		criteria for file selection.
//
//	deleteEmptySourceDirectory	bool
//
//		This parameter controls whether empty source
//		directories will be deleted after completion of
//		the move operation.
//
//		If 'deleteEmptySourceDirectory' is set to 'true'
//		and there are no files or subdirectories
//		remaining in the source directory (sourceDMgr)
//		after completion of the move operation, the
//		source directory will be deleted.
//
//		If 'deleteEmptySourceDirectory' is set to
//		'false', the source directory (sourceDMgr) will
//		NOT be deleted after completion of the move
//		operation.
//
//	moveRegularFiles			bool
//
//		If this parameter is set to 'true', Regular Files,
//		which also meet the File Selection Characteristics
//		criteria (fileSelectCriteria), will be included
//		in the move operation.
//
//		Regular Files include text files, image files and
//		executable files.
//
//		For an explanation of Regular and Non-Regular
//		files, see the section on "Definition Of Terms",
//		above.
//
//		If input parameters 'moveRegularFiles',
//		'moveSymLinkFiles' and 'moveOtherNonRegularFiles'
//		are all set to 'false', an error will be returned.
//
//	moveSymLinkFiles			bool
//
//		If this parameter is set to 'true', SymLink Files
//		which also meet the File Selection Characteristics
//		criteria (fileSelectCriteria), will be included
//		in the move operation.
//
//		For an explanation of Regular and Non-Regular
//		files, see the section on "Definition Of Terms",
//		above.
//
//		If input parameters 'moveRegularFiles',
//		'moveSymLinkFiles' and 'moveOtherNonRegularFiles'
//		are all set to 'false', an error will be returned.
//
//	moveOtherNonRegularFiles	bool
//
//		If this parameter is set to 'true', Other
//		Non-Regular Files, which also meet the File
//		Selection Characteristics criteria
//		(fileSelectCriteria), will be included in the
//		move operation.
//
//		Examples of other non-regular file types
//		include device files, named pipes, and sockets.
//		See the Definition Of Terms section above.
//
//		If input parameters 'moveRegularFiles',
//		'moveSymLinkFiles' and 'moveOtherNonRegularFiles'
//		are all set to 'false', an error will be returned.
//
//	fileSelectCriteria			FileSelectionCriteria
//
//		In addition to the File Type Selection Criteria,
//		selected files must conform to the File
//		Characteristics criteria specified by this
//		parameter, 'fileSelectCriteria'.
//
//		File Characteristics Selection criteria allow
//		users to screen files for File Name, File
//		Modification Date and File Mode.
//
//		Files matching these selection criteria, and the
//		File Type filter, will be included in the move
//		operation performed by this method.
//
//		type FileSelectionCriteria struct {
//		 FileNamePatterns    []string
//			An array of strings containing File Name Patterns
//
//		 FilesOlderThan      time.Time
//		 	Match files with older modification date times
//
//		 FilesNewerThan      time.Time
//		 	Match files with newer modification date times
//
//		 SelectByFileMode    FilePermissionConfig
//		 	Match file mode (os.FileMode).
//
//		 SelectCriterionModeFileSelectCriterionMode
//		 	Specifies 'AND' or 'OR' selection mode
//		}
//
//	  The FileSelectionCriteria Type allows for
//	  configuration of single or multiple file selection
//	  criterion. The 'SelectCriterionMode' can be used to
//	  specify whether the file must match all, or any one,
//	  of the active file selection criterion.
//
//	  Elements of the File Characteristics Selection
//	  Criteria are described below:
//
//			FileNamePatterns		[]string
//
//				An array of strings which may define one or more
//				search patterns. If a file name matches any one
//				of the search pattern strings, it is deemed to be
//				a 'match' for the search pattern criterion.
//
//				Example Patterns:
//					FileNamePatterns = []string{"*.log"}
//					FileNamePatterns = []string{"current*.txt"}
//					FileNamePatterns = []string{"*.txt", "*.log"}
//
//				If this string array has zero length or if
//				all the strings are empty strings, then this
//				file search criterion is considered 'Inactive'
//				or 'Not Set'.
//
//
//			FilesOlderThan		time.Time
//
//				This date time type is compared to file
//				modification date times in order to determine
//				whether the file is older than the
//				'FilesOlderThan' file selection criterion. If
//				the file modification date time is older than
//				the 'FilesOlderThan' date time, that file is
//				considered a 'match' for this file selection
//				criterion.
//
//				If the value of 'FilesOlderThan' is set to
//				time zero, the default value for type
//				time.Time{}, then this file selection
//				criterion is considered to be 'Inactive' or
//				'Not Set'.
//
//			FilesNewerThan      time.Time
//
//				This date time type is compared to the file
//				modification date time in order to determine
//				whether the file is newer than the
//				'FilesNewerThan' file selection criterion. If
//				the file modification date time is newer than
//				the 'FilesNewerThan' date time, that file is
//				considered a 'match' for this file selection
//				criterion.
//
//				If the value of 'FilesNewerThan' is set to
//				time zero, the default value for type
//				time.Time{}, then this file selection
//				criterion is considered to be 'Inactive' or
//				'Not Set'.
//
//			SelectByFileMode  FilePermissionConfig
//
//				Type FilePermissionConfig encapsulates an os.FileMode. The
//				file selection criterion allows for the selection of files
//				by File Mode.
//
//				File modes are compared to the value of 'SelectByFileMode'.
//				If the File Mode for a given file is equal to the value of
//				'SelectByFileMode', that file is considered to be a 'match'
//				for this file selection criterion. Examples for setting
//				SelectByFileMode are shown as follows:
//
//				fsc := FileSelectionCriteria{}
//
//				err = fsc.SelectByFileMode.SetByFileMode(os.FileMode(0666))
//
//				err = fsc.SelectByFileMode.SetFileModeByTextCode("-r--r--r--")
//
//			SelectCriterionMode FileSelectCriterionMode
//
//			This parameter selects the manner in which the file selection
//			criteria above are applied in determining a 'match' for file
//			selection purposes. 'SelectCriterionMode' may be set to one of
//			two constant values:
//
//			(1) FileSelectCriterionMode(0).ANDSelect()
//
//				File selected if all active selection criteria
//				are satisfied.
//
//				If this constant value is specified for the file selection mode,
//				then a given file will not be judged as 'selected' unless all
//				the active selection criterion are satisfied. In other words, if
//				three active search criterion are provided for 'FileNamePatterns',
//				'FilesOlderThan' and 'FilesNewerThan', then a file will NOT be
//				selected unless it has satisfied all three criterion in this example.
//
//			(2) FileSelectCriterionMode(0).ORSelect()
//
//				File selected if any active selection criterion is satisfied.
//
//				If this constant value is specified for the file selection mode,
//				then a given file will be selected if any one of the active file
//				selection criterion is satisfied. In other words, if three active
//				search criterion are provided for 'FileNamePatterns', 'FilesOlderThan'
//				and 'FilesNewerThan', then a file will be selected if it satisfies any
//				one of the three criterion in this example.
//
//		------------------------------------------------------------------------
//
//		IMPORTANT:
//
//		If all of the file selection criterion in the FileSelectionCriteria
//		object are 'Inactive' or 'Not Set' (set to their zero or default values),
//		then all the files meeting the File Type requirements in the directory
//		defined by 'sourceDMgr' will be selected.
//
//			Example:
//			  fsc := FileSelectCriterionMode{}
//
//			  In this example, 'fsc' is NOT initialized. Therefore,
//			  all the selection criterion are 'Inactive'. Consequently,
//			  all the files meeting the File Type requirements in the
//			  directory defined by 'sourceDMgr' will be selected.
//
//		------------------------------------------------------------------------
//
//	sourceDMgrLabel				string
//
//		The name or label associated with input parameter
//		'sourceDMgr' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "sourceDMgr" will be
//		automatically applied.
//
//	targetDMgrLabel				string
//
//		The name or label associated with input parameter
//		'targetDMgr' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "targetDMgr" will be
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
//	dirMoveStats				DirectoryMoveStats
//
//		If this method completes successfully, this
//		returned instance of DirectoryMoveStats will be
//		populated with statistics and information
//		describing details of the move operation
//		performed by this method.
//
//		type DirectoryMoveStats struct {
//			TotalSrcFilesProcessed   uint64
//			SourceFilesMoved         uint64
//			SourceFileBytesMoved     uint64
//			SourceFilesRemaining     uint64
//			SourceFileBytesRemaining uint64
//			SourceSubDirsMoved       uint64
//			SourceSubDirsRemaining   uint64
//			TotalDirsProcessed       uint64
//			TargetDirsCreated        uint64
//			NumOfSubDirectories      uint64
//			SourceDirWasDeleted      bool
//			ComputeError             error
//		}
//
//	movedFiles					FileMgrCollection
//
//		If input parameter 'returnMovedFilesList' is set
//		to 'true', 'movedFiles' will return a populated
//		File Manager Collection documenting all the files
//		actually included in this move operation.
//
//		If input parameter 'returnMovedFilesList' is set
//		to 'false', 'movedFiles' will return an empty and
//		unpopulated instance of FileMgrCollection.
//
//	nonfatalErrs				[]error
//
//		An array of error objects.
//
//		If this method completes successfully, the
//		returned error array is set equal to 'nil'.
//
//		If non-fatal errors are encountered during
//		processing, the returned error Type will
//		encapsulate appropriate error messages.
//
//		Non-fatal errors usually involve processing
//		failures associated with individual files.
//
//		The returned error messages will incorporate
//		the method chain and text passed by input
//		parameter, 'errPrefDto'. The 'errPrefDto' text
//		will be prefixed or attached to the beginning of
//		the error message.
//
//		This error array may contain multiple errors.
//
//		An error array may be consolidated into a single
//		error using method StrMech.ConsolidateErrors()
//
//	fatalErr					error
//
//		If this method completes successfully, this
//		returned error Type is set equal to 'nil'.
//
//		If a fatal error is encountered during
//		processing, this returned error Type will
//		encapsulate an appropriate error message. This
//		returned error message will incorporate the
//		method chain and text passed by input parameter,
//		'errPrefDto'. The 'errPrefDto' text will be
//		prefixed or attached to the	beginning of the error
//		message.
//
//		Fatal errors are returned when the nature of the
//		processing failure is such that it is no longer
//		reasonable to continue code execution.
func (dMgrHlprAtom *dirMgrHelperAtom) moveDirectoryFiles(
	sourceDMgr *DirMgr,
	targetDMgr *DirMgr,
	returnMovedFilesList bool,
	copyEmptyTargetDirectory bool,
	deleteEmptySourceDirectory bool,
	moveRegularFiles bool,
	moveSymLinkFiles bool,
	moveOtherNonRegularFiles bool,
	fileSelectCriteria FileSelectionCriteria,
	sourceDMgrLabel string,
	targetDMgrLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	dirMoveStats DirectoryMoveStats,
	movedFiles FileMgrCollection,
	nonfatalErrs []error,
	fatalErr error) {

	if dMgrHlprAtom.lock == nil {
		dMgrHlprAtom.lock = new(sync.Mutex)
	}

	dMgrHlprAtom.lock.Lock()

	defer dMgrHlprAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "dirMgrHelperAtom." +
		"moveDirectoryFiles()"

	ePrefix,
		fatalErr = ePref.ErrPrefixDto{}.NewIEmpty(
		errPrefDto,
		funcName,
		"")

	if fatalErr != nil {

		return dirMoveStats,
			movedFiles,
			nonfatalErrs,
			fatalErr
	}

	if len(sourceDMgrLabel) == 0 {

		sourceDMgrLabel = "sourceDMgr"
	}

	dMgrHlprPreon := new(dirMgrHelperPreon)

	_,
		_,
		fatalErr = dMgrHlprPreon.
		validateDirMgr(
			sourceDMgr,
			true,
			sourceDMgrLabel,
			ePrefix)

	if fatalErr != nil {

		return dirMoveStats,
			movedFiles,
			nonfatalErrs,
			fatalErr
	}

	if len(targetDMgrLabel) == 0 {

		targetDMgrLabel = "targetDMgr"
	}

	var targetPathDoesExist bool

	_,
		targetPathDoesExist,
		fatalErr = dMgrHlprPreon.
		validateDirMgr(
			targetDMgr,
			false,
			targetDMgrLabel,
			ePrefix)

	if fatalErr != nil {

		return dirMoveStats,
			movedFiles,
			nonfatalErrs,
			fatalErr
	}

	if moveRegularFiles == false &&
		moveSymLinkFiles == false &&
		moveOtherNonRegularFiles == false {

		fatalErr = fmt.Errorf("%v\n"+
			"Fatal Error: File Type filters are conflicted!\n"+
			"All of the File Type filters are set to 'false'\n"+
			"This gurantees that NO files will be selected.\n"+
			"moveRegularFiles == false\n"+
			"moveSymLinkFiles == false\n"+
			"moveOtherNonRegularFiles == false\n",
			ePrefix.String())

		return dirMoveStats,
			movedFiles,
			nonfatalErrs,
			fatalErr
	}

	var err2 error
	dMgrHlprMolecule := dirMgrHelperMolecule{}

	if !targetPathDoesExist &&
		copyEmptyTargetDirectory {

		targetPathDoesExist,
			err2 = dMgrHlprMolecule.
			lowLevelMakeDir(
				targetDMgr,
				targetDMgrLabel,
				ePrefix)

		if err2 != nil {

			fatalErr = fmt.Errorf("%v\n"+
				"Error creating target directory!\n"+
				"copyEmptyTargetDirectory == true\n"+
				"%v Directory='%v'\n"+
				"Error= \n%v\n",
				funcName,
				targetDMgrLabel,
				targetDMgr.absolutePath,
				err2.Error())

			return dirMoveStats,
				movedFiles,
				nonfatalErrs,
				fatalErr
		}

	}

	isFileSelectionCriteriaActive :=
		fileSelectCriteria.IsSelectionCriteriaActive()

	var fileInfos []FileInfoPlus
	var lenFileInfos int
	var dMgrHlprTachyon = new(dirMgrHelperTachyon)
	var errs2 []error

	fileInfos,
		lenFileInfos,
		errs2,
		fatalErr = dMgrHlprTachyon.
		getFileInfosFromDirectory(
			sourceDMgr,
			false,                    // getDirectoryFileInfos
			moveRegularFiles,         // getRegularFileInfos
			moveSymLinkFiles,         // copySymLinkFiles,
			moveOtherNonRegularFiles, // copyOtherNonRegularFiles
			FileSelectionCriteria{},
			sourceDMgrLabel,
			ePrefix.XCpy(sourceDMgrLabel))

	if len(errs2) != 0 {

		nonfatalErrs = append(nonfatalErrs, errs2...)

	}

	if fatalErr != nil {

		return dirMoveStats,
			movedFiles,
			nonfatalErrs,
			fatalErr
	}

	if lenFileInfos == 0 {

		fatalErr = fmt.Errorf("%v\n"+
			"Error: The %v source directory is EMPTY!\n"+
			"The move files operation cannot proceed.\n"+
			"Method dirMgrHelperElectron.getFileInfosFromDirectory()\n"+
			"returned a zero length array of File Info Objects from:\n"+
			"%v = %v\n",
			ePrefix.String(),
			sourceDMgrLabel,
			sourceDMgrLabel,
			sourceDMgr.absolutePath)

		return dirMoveStats,
			movedFiles,
			nonfatalErrs,
			fatalErr
	}

	var fh = new(FileHelper)
	var isMatch bool
	var srcFile, targetFile string

	osPathSepStr := string(os.PathSeparator)

	for _, nameFileInfo := range fileInfos {

		dirMoveStats.TotalSrcFilesProcessed++

		// nameFileInfo is an os.FileInfo

		if isFileSelectionCriteriaActive == true {

			isMatch,
				err2,
				_ =
				fh.FilterFileName(
					nameFileInfo,
					fileSelectCriteria,
					ePrefix.XCpy("nameFileInfo"))

			if err2 != nil {

				fatalErr =
					fmt.Errorf("%v\n"+
						"Error returned by fh.FilterFileName(nameFileInfo, fileSelectCriteria).\n"+
						"%v directory= '%v'\n"+
						"fileName= '%v'\n"+
						"Error= \n%v\n",
						funcName,
						sourceDMgrLabel,
						sourceDMgr.absolutePath,
						nameFileInfo.Name(),
						err2.Error())

				return dirMoveStats,
					movedFiles,
					nonfatalErrs,
					fatalErr
			}

		} else {

			isMatch = true
		}

		if !isMatch {

			dirMoveStats.SourceFilesRemaining++
			continue

		} else {
			// We have a Match!

			if !targetPathDoesExist {

				targetPathDoesExist,
					err2 = dMgrHlprMolecule.
					lowLevelMakeDir(
						targetDMgr,
						targetDMgrLabel,
						ePrefix)

				if err2 != nil {

					fatalErr = fmt.Errorf("%v\n"+
						"Error creating target directory!\n"+
						"%v Directory='%v'\n"+
						"Error= \n%v\n",
						funcName,
						targetDMgrLabel,
						targetDMgr.absolutePath,
						err2.Error())

					return dirMoveStats,
						movedFiles,
						nonfatalErrs,
						fatalErr
				}

			}

			srcFile = sourceDMgr.absolutePath +
				osPathSepStr + nameFileInfo.Name()

			targetFile = targetDMgr.absolutePath +
				osPathSepStr + nameFileInfo.Name()

			err2 = dMgrHlprMolecule.
				lowLevelCopyFile(
					srcFile,
					nameFileInfo,
					targetFile,
					"sourceFile",
					"destinationFile",
					ePrefix)

			if err2 != nil {

				fatalErr = fmt.Errorf("%v\n"+
					"Error copying source file to target directory!\n"+
					"%v Directory='%v'\n"+
					"%v Directory='%v'\n"+
					"Source File='%v\n"+
					"Target File='%v'\n"+
					"Error= \n%v\n",
					funcName,
					sourceDMgrLabel,
					sourceDMgr.absolutePath,
					targetDMgrLabel,
					targetDMgr.absolutePath,
					srcFile,
					targetFile,
					err2.Error())

				return dirMoveStats,
					movedFiles,
					nonfatalErrs,
					fatalErr
			}

			err2 = os.Remove(srcFile)

			if err2 != nil {

				fatalErr = fmt.Errorf("%v\n"+
					"Error occurred after file copy completed during the\n"+
					"delete source file operation!\n"+
					"Error returned by os.Remove(sourceFile)\n"+
					"sourceFile='%v'\n"+
					"Error= \n%v\n",
					ePrefix.String(),
					srcFile,
					err2.Error())

				return dirMoveStats,
					movedFiles,
					nonfatalErrs,
					fatalErr
			}

			dirMoveStats.SourceFilesMoved++

			if returnMovedFilesList {

				err2 = movedFiles.AddFileMgrByDirFileNameExt(
					*sourceDMgr,
					nameFileInfo.Name(),
					ePrefix)

				if err2 != nil {
					nonfatalErrs =
						append(nonfatalErrs, err2)
				}
			}
		}

	}

	fileInfos = nil

	if deleteEmptySourceDirectory {

		var dirProfile DirectoryProfile

		_,
			dirProfile,
			err2 = dMgrHlprTachyon.
			getDirectoryProfile(
				sourceDMgr,
				sourceDMgrLabel,
				ePrefix)

		if err2 != nil {

			fatalErr = fmt.Errorf("%v\n"+
				"Error occurred reading the %v Source Directory\n"+
				"Profile after the move operation was completed.\n"+
				"Error returned by dirMgrHelperTachyon.getDirectoryProfile()\n"+
				"%v Directory = '%v'\n"+
				"Error= \n%v\n",
				funcName,
				sourceDMgrLabel,
				sourceDMgrLabel,
				sourceDMgr.absolutePath,
				err2.Error())

			return dirMoveStats,
				movedFiles,
				nonfatalErrs,
				fatalErr

		}

		if dirProfile.DirExistsOnStorageDrive &&
			dirProfile.DirTotalFiles == 0 &&
			dirProfile.DirSubDirectories == 0 &&
			dirProfile.DirRegularFiles == 0 &&
			dirProfile.DirSymLinkFiles == 0 &&
			dirProfile.DirNonRegularFiles == 0 {
			// The source directory is empty

			err2 = dMgrHlprMolecule.
				lowLevelDeleteDirectoryAll(
					sourceDMgr,
					sourceDMgrLabel,
					ePrefix)

			fatalErr = fmt.Errorf("%v\n"+
				"Error occurred deleting the %v Source Directory\n"+
				"The %v Directory Profile showed 'Empty' with zero files remaining.\n"+
				"Error returned by dMgrHlprMolecule.lowLevelDeleteDirectoryAll()\n"+
				"%v Directory = '%v'\n"+
				"Error= \n%v\n",
				funcName,
				sourceDMgrLabel,
				sourceDMgrLabel,
				sourceDMgrLabel,
				sourceDMgr.absolutePath,
				err2.Error())

		}

	}

	return dirMoveStats,
		movedFiles,
		nonfatalErrs,
		fatalErr
}
