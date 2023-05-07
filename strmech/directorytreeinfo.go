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
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	baseDir						DirMgr
//
//		A concrete instance of DirMgr. The directory path
//		contained in this instance will be used as the source
//		for the directory tree which will be copied to the
//		directory path contained in 'newBaseDir'.
//
//	newBaseDir					DirMgr
//
//		A concrete instance of DirMgr. The directory path
//		contained in this instance will be used as the
//		destination path for the directory tree from the
//		'baseDir' directory path.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	DirectoryTreeInfo
//
//		If this method completes successfully, without
//		errors, a fully populated instance of
//		DirectoryTreeInfo is returned.
//
//		This structure contains detailed information on
//		files processed during copy operation.
//
//	          type DirectoryTreeInfo struct {
//
//	            StartPath             string
//					The starting path or directory for the file
//	                search.
//
//	            Directories           DirMgrCollection
//					Directory Managers found during directory tree
//					search.
//
//					This collection will ALWAYS return the parent
//					directory ('DirMgr') as the first entry in the
//					collection.
//
//	            FoundFiles            FileMgrCollection
//					Found Files matching file selection criteria
//
//	            ErrReturns            []error
//					Internal System errors encountered
//
//	            FileSelectCriteria    FileSelectionCriteria
//	            	The File Selection Criteria submitted as an
//					input parameter to this method.
//	           }
//
//	        If successful, files matching the file selection criteria input
//	        parameter shown above will be returned in a 'DirectoryTreeInfo'
//	        object. The field 'DirectoryTreeInfo.FoundFiles' contains information
//	        on all the files in the specified directory tree which match the file selection
//	        criteria.
//
//	        Note: It is a good idea to check the returned field 'DirectoryTreeInfo.ErrReturns'
//	              to determine if any internal system errors were encountered while processing
//	              the directory tree.
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
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
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
