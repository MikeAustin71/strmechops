package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// DirectoryTreeInfo - structure used
// to 'Find' files in a directory specified
// by 'StartPath'. The file search will be
// filtered by a 'FileSelectCriteria' object.
//
// 'FileSelectCriteria' is a FileSelectionCriteria type
// which contains FileNamePatterns strings and
// 'FilesOlderThan' or 'FilesNewerThan' date time
// parameters which can be used as a selection
// criteria.
type DirectoryTreeInfo struct {
	StartPath          string
	Directories        DirMgrCollection
	FoundFiles         FileMgrCollection
	ErrReturns         []error
	FileSelectCriteria FileSelectionCriteria
	lock               *sync.Mutex
}

// CopyToDirectoryTree
//
// Copies an entire directory tree to an alternate location.
//
// The copy operation includes all files and all directories in the designated directory
// tree.
func (dirTree *DirectoryTreeInfo) CopyToDirectoryTree(
	baseDir DirMgr,
	newBaseDir DirMgr,
	errorPrefix interface{}) (
	DirectoryTreeInfo,
	error) {

	if dirTree.lock == nil {
		dirTree.lock = new(sync.Mutex)
	}

	dirTree.lock.Lock()

	defer dirTree.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newDirTree := DirectoryTreeInfo{}

	funcName := "DirectoryTreeInfo.CopyToDirectoryTree()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {
		return newDirTree, err
	}

	err = baseDir.IsDirMgrValid(
		ePrefix.XCpy("baseDir"))

	if err != nil {
		return newDirTree, fmt.Errorf("%v\n"+
			"Error: Input Parameter 'baseDir' is INVALID!\n"+
			"Error= \n%v\n",
			funcName,
			err.Error())
	}

	err = newBaseDir.IsDirMgrValid(ePrefix.XCpy(
		"newBaseDir"))

	if err != nil {
		return newDirTree, fmt.Errorf("%v\n"+
			"Error: Input Parameter 'newBaseDir' is INVALID!\n"+
			"Error= \n%v\n",
			funcName,
			err.Error())
	}

	err = newBaseDir.MakeDir(ePrefix.XCpy(
		"newBaseDir"))

	if err != nil {
		return newDirTree,
			fmt.Errorf("%v\n"+
				"Error returned from newBaseDir.MakeDir().\n"+
				"newBaseDir.absolutePath='%v'\n"+
				"Error= \n%v\n",
				funcName,
				newBaseDir.absolutePath,
				err.Error())
	}

	if dirTree.Directories.dirMgrs == nil {

		dirTree.Directories.dirMgrs = make([]DirMgr, 10)

	}

	lAry := len(dirTree.Directories.dirMgrs)

	var newDMgr DirMgr
	var fileDMgr DirMgr
	var newFileMgr FileMgr

	// Make the new Directory Tree
	for i := 0; i < lAry; i++ {

		newDMgr,
			err = dirTree.Directories.dirMgrs[i].
			SubstituteBaseDir(
				baseDir,
				newBaseDir,
				ePrefix.XCpy(fmt.Sprintf(
					"dirTree.Directories.dirMgrs[%v]", i)))

		if err != nil {
			return DirectoryTreeInfo{},
				fmt.Errorf("%v\n"+
					"Error returned from SubstituteBaseDir(baseDir, newBaseDir).\n"+
					"i='%v'\n"+
					"Error= \n%v\n",
					funcName,
					i,
					err.Error())
		}

		err = newDMgr.MakeDir(ePrefix.XCpy(
			"newDMgr"))

		if err != nil {
			return DirectoryTreeInfo{},
				fmt.Errorf("%v\n"+
					"Error returned from newDMgr.MakeDir()\n"+
					"Error= \n%v\n",
					funcName,
					err.Error())

		}

		err = newDirTree.Directories.
			AddDirMgr(newDMgr, ePrefix.XCpy(
				"newDirTree<-newDMgr"))

		if err != nil {
			return DirectoryTreeInfo{},
				fmt.Errorf("%v\n"+
					"Error returned from newDirTree.Directories.AddDirMgr()\n"+
					"Error= \n%v\n",
					funcName,
					err.Error())

		}
	}

	if dirTree.FoundFiles.fileMgrs == nil {
		dirTree.FoundFiles.fileMgrs = make([]FileMgr, 0, 10)
	}

	lAry = len(dirTree.FoundFiles.fileMgrs)

	for j := 0; j < lAry; j++ {

		fileDMgr,
			err = dirTree.FoundFiles.fileMgrs[j].dMgr.
			SubstituteBaseDir(
				baseDir,
				newBaseDir,
				ePrefix.XCpy(
					fmt.Sprintf("dirTree.FoundFiles.fileMgrs[%v].dMgr",
						j)))

		if err != nil {
			return DirectoryTreeInfo{},
				fmt.Errorf("%v\n"+
					"Error returned by dirTree.FoundFiles.fileMgrs[j].dMgr.SubstituteBaseDir(baseDir, newBaseDir).\n"+
					"Error= \n%v\n",
					funcName,
					err.Error())
		}

		newFileMgr,
			err =
			new(FileMgr).NewFromDirMgrFileNameExt(
				fileDMgr,
				dirTree.FoundFiles.fileMgrs[j].fileNameExt,
				ePrefix.XCpy(
					fmt.Sprintf("dirTree.FoundFiles.fileMgrs[%v].fileNameExt",
						j)))

		if err != nil {
			return DirectoryTreeInfo{},
				fmt.Errorf("%v\n"+
					"Error returned by FileMgr{}.NewFromDirMgrFileNameExt("+
					"dMgr, dirTree.FoundFiles.fileMgrs[j].fileNameExt)\n"+
					"dirTree.FoundFiles.fileMgrs[j].fileNameExt='%v'\n"+
					"j='%v'\n"+
					"Error= \n%v\n",
					funcName,
					dirTree.FoundFiles.fileMgrs[j].fileNameExt,
					j,
					err.Error())
		}

		err = dirTree.FoundFiles.fileMgrs[j].
			CopyFileMgrByIoByLink(
				&newFileMgr,
				ePrefix.XCpy(
					fmt.Sprintf("newFileMgr<-"+
						"dirTree.FoundFiles.fileMgrs[%v]",
						j)))

		if err != nil {
			return DirectoryTreeInfo{},
				fmt.Errorf("%v\n"+
					"Error returned by fileMgrs[j].CopyFileMgrByIoByLink(&newFileMgr)\n"+
					"SrcFileName:'%v'\n"+
					"DestFileName:'%v'\n"+
					"Error= \n%v\n",
					funcName,
					dirTree.FoundFiles.fileMgrs[j].fileNameExt,
					newFileMgr.fileNameExt,
					err.Error())

		}

		err = newDirTree.FoundFiles.
			AddFileMgr(
				newFileMgr,
				ePrefix.XCpy(
					"newDirTree.FoundFiles<-newFileMgr"))

		if err != nil {

			return DirectoryTreeInfo{},
				fmt.Errorf("%v\n"+
					"Error returned by newDirTree.FoundFiles.AddFileMgr()\n"+
					"SrcFileName:'%v'\n"+
					"DestFileName:'%v'\n"+
					"'j' Loop No :'%v'"+
					"Error= \n%v\n",
					funcName,
					dirTree.FoundFiles.fileMgrs[j].fileNameExt,
					newFileMgr.fileNameExt,
					j,
					err.Error())

		}
	}

	return newDirTree, nil
}
