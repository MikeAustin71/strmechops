package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"os"
	pf "path/filepath"
	"strings"
	"sync"
)

// dirMgrHelper
//
// Provides helper methods for type DirMgr
type dirMgrHelper struct {
	dMgr DirMgr

	lock *sync.Mutex
}

// deleteDirectoryAll
//
// This method will remove the directory identified by
// the input parameter 'dMgr'. This means it will delete
// the parent directory and all files residing therein as
// well as deleting all child directories and files in
// the directory tree identified by 'dMgr'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	All files and directories, including the parent
//	directory, identified by input parameter 'dMgr' will
//	be deleted.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dMgr						*DirMgr
//
//		A pointer to an instance of DirMgr. The entire
//		directory tree identified by this parameter will
//		be deleted along with all the resident files.
//
//	dMgrLabel					string
//
//		The name or label associated with input parameter
//		'dMgr' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "dMgr" will be
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
func (dMgrHlpr *dirMgrHelper) deleteDirectoryAll(
	dMgr *DirMgr,
	dMgrLabel string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if dMgrHlpr.lock == nil {
		dMgrHlpr.lock = new(sync.Mutex)
	}

	dMgrHlpr.lock.Lock()

	defer dMgrHlpr.lock.Unlock()

	funcName := "dirMgrHelper.deleteDirectoryAll() "

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return err
	}

	if len(dMgrLabel) == 0 {

		dMgrLabel = "dMgr"

	}

	var dirPathDoesExist bool

	_,
		dirPathDoesExist,
		err = new(dirMgrHelperPreon).
		validateDirMgr(
			dMgr,
			false,
			dMgrLabel,
			ePrefix)

	if err != nil {

		return err
	}

	if !dirPathDoesExist {
		return nil
	}

	err = new(dirMgrHelperMolecule).
		lowLevelDeleteDirectoryAll(
			dMgr,
			dMgrLabel,
			ePrefix)

	if err != nil {
		return err
	}

	dirPathDoesExist,
		_,
		err =
		new(dirMgrHelperAtom).
			doesDirectoryExist(
				dMgr,
				PreProcPathCode.None(),
				dMgrLabel,
				ePrefix.XCpy(dMgrLabel))

	if err != nil {
		return err
	}

	if dirPathDoesExist {

		err = fmt.Errorf("%v\n"+
			"Error: FAILED TO DELETE DIRECTORY!!\n"+
			"Directory Path still exists!\n"+
			"Directory Path= '%v'\n",
			ePrefix.String(),
			dMgr.absolutePath)
	}

	return err
}

// deleteDirectoryTreeInfo
//
// This method will optionally delete files in the entire
// directory tree identified by input parameter 'dMgr'.
//
// This means that files in the top level directory, or
// exclusively in the subdirectory tree, will be deleted.
//
// The specific files to be deleted are identified by
// means of a selection criteria configured by the user
// and passed as input parameter
// 'deleteFileSelectionCriteria'.
//
// This Helper method similar to:
//
//	dirMgrHelper.deleteDirectoryTreeStats()
//
// This method differs from similar methods in that it
// returns a type DirectoryDeleteFileInfo containing
// information and statistics on the deleted files.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete files in the entire directory
//	tree identified by input parameter 'dMgr'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dMgr							*DirMgr
//
//		A pointer to an instance of DirMgr. This method
//		will delete files in the directory tree
//		identified by this parameter.
//
//	deleteFileSelectionCriteria		FileSelectionCriteria
//
//		This input parameter should be configured with the
//		desired file selection criteria. Files matching
//		this criteria will be deleted from the directory
//		tree identified by input parameter, 'dMgr'.
//
//		type FileSelectionCriteria struct {
//
//			FileNamePatterns    []string
//				An array of strings containing File Name Patterns
//
//			FilesOlderThan      time.Time
//				Match files with older modification date times
//
//			FilesNewerThan      time.Time
//				Match files with newer modification date times
//
//			RegularExp			*regexp.Regexp
//				Used to select file names with regular
//				expressions. If this parameter is NOT
//				equal to nil, file names will be
//				analyzed using MatchString().
//
//				Example:
//					RegularExp.MatchString("someFileName.txt")
//
//			SelectByFileMode    FilePermissionConfig
//				Match file mode (os.FileMode).
//
//			SelectCriterionModeFileSelectCriterionMode
//				Specifies 'AND' or 'OR' selection mode
//		}
//
//	  The FileSelectionCriteria type allows for
//	  configuration of single or multiple file selection
//	  criterion. The 'SelectCriterionMode' can be used to
//	  specify whether the file must match all, or any one,
//	  of the active file selection criterion.
//
//	  Elements of the FileSelectionCriteria are described
//	  below:
//
//	  FileNamePatterns		[]string
//	  	An array of strings which may define one or more
//	  	search patterns. If a file name matches any one
//	  	of the search pattern strings, it is deemed to be
//	  	a 'match' for the search pattern criterion.
//
//		Example Patterns:
//			FileNamePatterns = []string{"*.log"}
//			FileNamePatterns = []string{"current*.txt"}
//			FileNamePatterns = []string{"*.txt", "*.log"}
//
//		If this string array has zero length or if
//		all the strings are empty strings, then this
//		file search criterion is considered 'Inactive'
//		or 'Not Set'.
//
//
//		FilesOlderThan		time.Time
//			This date time type is compared to file
//			modification date times in order to determine
//			whether the file is older than the
//			'FilesOlderThan' file selection criterion. If
//			the file modification date time is older than
//			the 'FilesOlderThan' date time, that file is
//			considered a 'match' for this file selection
//			criterion.
//
//			If the value of 'FilesOlderThan' is set to
//			time zero, the default value for type
//			time.Time{}, then this file selection
//			criterion is considered to be 'Inactive' or
//			'Not Set'.
//
//		FilesNewerThan      time.Time
//			This date time type is compared to the file
//			modification date time in order to determine
//			whether the file is newer than the
//			'FilesNewerThan' file selection criterion. If
//			the file modification date time is newer than
//			the 'FilesNewerThan' date time, that file is
//			considered a 'match' for this file selection
//			criterion.
//
//			If the value of 'FilesNewerThan' is set to
//			time zero, the default value for type
//			time.Time{}, then this file selection
//			criterion is considered to be 'Inactive' or
//			'Not Set'.
//
//		RegularExp				*regexp.Regexp
//
//			Used to select file names with regular
//			expressions. If this parameter is NOT
//			equal to nil, file names will be
//			analyzed using MatchString().
//
//			Example:
//				RegularExp.MatchString("someFileName.txt")
//
//		SelectByFileMode		FilePermissionConfig
//			Type FilePermissionConfig encapsulates an
//			instance of os.FileMode. The file selection
//			criterion allows for the selection of files
//			by File Mode. File modes are compared to the
//			value of 'SelectByFileMode'. If the File Mode
//			for a given file is equal to the value of
//			'SelectByFileMode', that file is considered
//			to be a 'match' for this file selection
//			criterion. Examples for setting
//			SelectByFileMode are shown as follows:
//
//				fsc := FileSelectionCriteria{}
//
//				err =
//					fsc.SelectByFileMode.
//						SetByFileMode(
//							os.FileMode(0666))
//
//				err =
//					fsc.SelectByFileMode.
//						SetFileModeByTextCode(
//							"-r--r--r--")
//
//		SelectCriterionMode		FileSelectCriterionMode
//			This parameter selects the manner in which
//			the file selection criteria above are applied
//			in determining a 'match' for file selection
//			purposes. 'SelectCriterionMode' may be set to
//			one of two constant values:
//
//			FileSelectMode.ANDSelect()
//				File selected if all active selection
//				criteria are satisfied.
//
//				If this constant value is specified for
//				the file selection mode, then a given
//				file will not be judged as 'selected'
//				unless all the active selection criterion
//				are satisfied. In other words, if three
//				active search criterion are provided for
//				'FileNamePatterns', 'FilesOlderThan' and
//				'FilesNewerThan', then a file will NOT be
//				selected unless it has satisfied all three
//				criterion in this example.
//
//			FileSelectMode.ORSelect()
//				File selected if any active selection
//				criterion is satisfied.
//
//				If this constant value is specified for
//				the file selection mode, then a given
//				file will be selected if any one of the
//				active file selection criterion is
//				satisfied. In other words, if three
//				active search criterion are provided for
//				'FileNamePatterns', 'FilesOlderThan' and
//				'FilesNewerThan', then a file will be
//				selected if it satisfies any one of the
//				three criterion in this example.
//
//	skipTopLevelDirectory			bool
//
//		If this parameter is set to 'true', the parent or
//		top level directory identified by input parameter
//		'dMgr' will be skipped. This means no files will
//		be deleted from the parent or top level directory.
//
//		Be careful to ensure that parameters
//		'skipTopLevelDirectory' and 'scanSubDirectories'
//		are not in conflict. If 'skipTopLevelDirectory'
//		is set to 'true' and 'scanSubDirectories' is set
//		to 'false', an error will be returned.
//
//	scanSubDirectories				bool
//
//		If this parameter is set to 'true', it means that
//		child directories (a.k.a subdirectories) will be
//		searched and eligible files will be deleted from
//		subsidiary directories.
//
//		Conversely, if this parameter is set to 'false', no
//		files will be deleted from child directories (a.k.a
//		subdirectories).
//
//		Be careful to ensure that parameters
//		'skipTopLevelDirectory' and 'scanSubDirectories'
//		are not in conflict. If 'skipTopLevelDirectory'
//		is set to 'true' and 'scanSubDirectories' is set
//		to 'false', an error will be returned.
//
//	dMgrLabel						string
//
//		The name or label associated with input parameter
//		'dMgr' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "dMgr" will be
//		automatically applied.
//
//	deleteFileSelectLabel			string
//
//		The name or label used to describe the type of
//		files being deleted. This label will be used in
//		error messages returned by this method.
//
//		Example:
//			deleteFileSelectLabel = "Outdated files"
//
//		If this parameter is submitted as an empty
//		string, it will be automatically defaulted to a
//		value of "Target Files for Deletion".
//
//	errPrefDto						*ePref.ErrPrefixDto
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
//	deleteTreeInfo					DirectoryDeleteFileInfo
//
//		If this method completes successfully, this
//		return parameter will be populated with
//		information and statistics on the file deletion
//		operation. This information includes the number
//		of files deleted.
//
//		DirectoryDeleteFileInfo - A structure used
//		to delete files in a directory specified
//		by a starting directory.
//
//		The file selection criteria for target files
//		to be deleted is stored in member variable
//		'DeleteFileSelectCriteria'.
//
//		type DirectoryDeleteFileInfo struct {
//			StartPath                string
//			Directories              DirMgrCollection
//			ErrReturns               []error
//			DeleteFileSelectCriteria FileSelectionCriteria
//			DeletedFiles             FileMgrCollection
//		}
//
//	errs							[]error
//
//		An array of error objects.
//
//		If this method completes successfully, the
//		returned error array is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errPrefDto'.
//	 	The 'errPrefDto' text will be prefixed or
//	 	attached to the	beginning of the error message.
//
//		This error array may contain multiple errors.
//
//		An error array may be consolidated into a single
//		error using method StrMech.ConsolidateErrors()
func (dMgrHlpr *dirMgrHelper) deleteDirectoryTreeInfo(
	dMgr *DirMgr,
	deleteFileSelectionCriteria FileSelectionCriteria,
	skipTopLevelDirectory bool,
	scanSubDirectories bool,
	dMgrLabel string,
	deleteFileSelectLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	deleteTreeInfo DirectoryDeleteFileInfo,
	errs []error) {

	if dMgrHlpr.lock == nil {
		dMgrHlpr.lock = new(sync.Mutex)
	}

	dMgrHlpr.lock.Lock()

	defer dMgrHlpr.lock.Unlock()

	errs = make([]error, 0)

	funcName := "dirMgrHelper.deleteDirectoryTreeInfo() "

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		errs = append(errs, err)

		return deleteTreeInfo, errs
	}

	if len(dMgrLabel) == 0 {
		dMgrLabel = "dMgr"
	}

	_,
		_,
		err = new(dirMgrHelperPreon).
		validateDirMgr(
			dMgr,
			true,
			dMgrLabel,
			ePrefix)

	if err != nil {

		errs = append(errs, err)

		return deleteTreeInfo, errs
	}

	if len(deleteFileSelectLabel) == 0 {
		deleteFileSelectLabel =
			"Target Files for Deletion"
	}

	if skipTopLevelDirectory &&
		!scanSubDirectories {

		err = fmt.Errorf("%v\n"+
			"ERROR: Conflicted Input parameters!\n"+
			"skipTopLevelDirectory=true\n"+
			"scanSubDirectories=false.\n"+
			"Impossible combination!!\n",
			ePrefix.String())

		errs = append(errs, err)

		return deleteTreeInfo, errs
	}

	var err2 error
	var nextDir *DirMgr

	osPathSepStr := string(os.PathSeparator)

	fh := new(FileHelper)

	isMatch := false

	isTopLevelDir := true

	deleteTreeInfo.StartPath = dMgr.absolutePath

	deleteTreeInfo.DeleteFileSelectCriteria = deleteFileSelectionCriteria

	err = deleteTreeInfo.Directories.AddDirMgr(
		*dMgr,
		ePrefix.XCpy("deleteTreeInfo.Directories<-dMgr"))

	if err != nil {

		errs = append(errs, err)

		return deleteTreeInfo, errs
	}

	dTreeCnt := 1

	var nameFileInfo os.FileInfo

	var nameDirEntries []os.DirEntry

	var errStatus ArrayColErrorStatus

	for i := 0; i < dTreeCnt; i++ {

		if i == 0 {

			isTopLevelDir = true

		} else {

			isTopLevelDir = false
		}

		nextDir,
			errStatus = deleteTreeInfo.Directories.
			GetDirMgrAtIndex(
				i,
				ePrefix.XCpy(
					fmt.Sprintf(
						"deleteTreeInfo.Directories idx=%v",
						i)))

		if errStatus.ProcessingError != nil {

			errs = append(errs,
				fmt.Errorf("%v",
					errStatus.ProcessingError.Error()))

			continue
		}

		nameDirEntries,
			err2 = os.ReadDir(nextDir.absolutePath)

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Error returned by os.ReadDir(nextDir.absolutePath).\n"+
				"nextDir.absolutePath= '%v'\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				nextDir.absolutePath,
				err2.Error())

			errs = append(errs, err)

			continue
		}

		for _, nameDirEntry := range nameDirEntries {

			if nameDirEntry.IsDir() {

				if !scanSubDirectories {
					continue
				}

				err =
					deleteTreeInfo.Directories.
						AddDirMgrByKnownPathDirName(
							nextDir.absolutePath,
							nameDirEntry.Name(),
							ePrefix.XCpy(
								"nextDir.absolutePath"))

				if err != nil {
					err2 =
						fmt.Errorf("%v\n"+
							"Error returned by dirs.AddDirMgrByKnownPathDirName(newDirPathFileName)\n"+
							"newDirPathFileName='%v'\n"+
							"Error= \n%v\n",
							ePrefix.String(),
							nextDir.absolutePath+osPathSepStr+nameDirEntry.Name(),
							err.Error())

					errs = append(errs, err2)

					continue
				}

				dTreeCnt++

			} else {
				// This is a file which is eligible for processing

				if isTopLevelDir && skipTopLevelDirectory {
					continue
				}

				nameFileInfo,
					err2 = nameDirEntry.Info()

				if err2 != nil {

					err = fmt.Errorf("%v\n"+
						"Error: Error Returned by nameDirEntry.Info().\n"+
						"The conversion of DirEntry to os.FileInfo Failed.\n"+
						"nextDir= '%v'\n"+
						"FileName= '%v'\n"+
						"Error= \n%v\n",
						ePrefix.String(),
						nextDir.absolutePath,
						nextDir.absolutePath+osPathSepStr+nameFileInfo.Name(),
						err2.Error())

					errs = append(errs, err)

					continue
				}

				if !nameFileInfo.Mode().IsRegular() {

					err = fmt.Errorf("%v\n"+
						"Error: File Name is NOT classified as a 'Regular' File!\n"+
						"nextDir= '%v'\n"+
						"fileName='%v'\n",
						ePrefix.String(),
						nextDir.absolutePath,
						nextDir.absolutePath+osPathSepStr+nameFileInfo.Name())

					errs = append(errs, err)

					continue
				}

				// This is not a directory. It is a file.
				// Determine if it matches the find file criteria.
				isMatch,
					err,
					_ =
					fh.FilterFileName(nameFileInfo,
						deleteFileSelectionCriteria,
						ePrefix.XCpy("nameFileInfo"))

				if err != nil {

					err2 =
						fmt.Errorf("%v\n"+
							"Error returned by fh.FilterFileName(nameDirEntry, %v).\n"+
							"%v directory searched='%v'\n"+
							"fileName='%v'\n"+
							"Error= \n%v\n",
							funcName,
							deleteFileSelectLabel,
							dMgrLabel,
							dMgr.absolutePath,
							nameFileInfo.Name(),
							err.Error())

					errs = append(errs, err2)

					continue
				}

				if !isMatch {

					continue

				} else {

					// We have a match, save file to deleteTreeInfo
					fileToDelete :=
						nextDir.absolutePath + osPathSepStr + nameFileInfo.Name()

					err = os.Remove(fileToDelete)

					if err != nil {

						err2 = fmt.Errorf("%v\n"+
							"Error returned by os.Remove(fileToDelete)\n"+
							"fileToDelete='%v'\n"+
							"Error= \n%v\n",
							ePrefix.String(),
							fileToDelete,
							err.Error())

						errs = append(errs, err2)

						continue
					}

					var dirMgrCopy DirMgr

					dirMgrCopy,
						err = nextDir.CopyOut(ePrefix.XCpy(
						"nextDir->"))

					if err != nil {

						err2 = fmt.Errorf("%v\n"+
							"DirMgr Copy Errors"+
							"ERROR returned by nextDir.CopyOut()\n"+
							"nextDir='%v'\n"+
							"Error= \n%v\n",
							funcName,
							nextDir.absolutePath,
							err.Error())

						errs = append(errs, err2)

						return deleteTreeInfo, errs
					}

					err = deleteTreeInfo.DeletedFiles.
						AddFileMgrByDirFileNameExt(
							dirMgrCopy,
							nameFileInfo.Name(),
							ePrefix)

					if err != nil {
						err2 = fmt.Errorf("%v\n"+
							"ERROR returned by deleteTreeInfo.DeletedFiles.AddFileMgrByDirFileNameExt(nextDir, fileNameExt)\n"+
							"nextDir='%v'\n"+
							"fileNameExt='%v'"+
							"Error= \n%v\n",
							funcName,
							nextDir.absolutePath,
							nameFileInfo.Name(),
							err.Error())

						errs = append(errs, err2)

					}
				}
			}

		} // End of nameDirEntry := range nameDirEntries

	} // End of for i:=0; i < dTreeCnt; i ++

	if len(deleteTreeInfo.Directories.dirMgrs) > 0 && skipTopLevelDirectory {

		_,
			_ = deleteTreeInfo.Directories.PopFirstDirMgr(
			ePrefix.XCpy(
				"deleteTreeInfo.Directories"))
	}

	for i := 0; i < len(errs); i++ {

		err2 = fmt.Errorf("%v", errs[i].Error())

		deleteTreeInfo.ErrReturns =
			append(deleteTreeInfo.ErrReturns, err2)
	}

	return deleteTreeInfo, errs
}

// deleteDirectoryTreeStats
//
// Helper method designed to delete files using file
// selection criteria. Scope of scans and file deletions
// is controlled by input parameter 'scanSubDirectories'.
//
// If 'scanSubDirectories' is set to 'true', files may be
// deleted in the entire directory tree. If set to 'false'
// the file deletions are limited solely to the parent or
// top level directory identified by the current instance
// of DirMgr.
//
// This method differs from
// dMgrHlpr.deleteDirectoryTreeInfo() in that this method
// returns a type DeleteDirFilesStats.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete files in the directory
//	tree specified by the current instance of DirMgr.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dMgr							*DirMgr
//
//		A pointer to an instance of DirMgr. This method
//		will delete files in the directory tree
//		identified by this parameter.
//
//
//	deleteFileSelectionCriteria		FileSelectionCriteria
//
//	  This input parameter should be configured with the
//	  desired file selection criteria. Files matching
//	  this criteria will be deleted from the directory
//	  tree identified by input parameter, 'dMgr'.
//
//		type FileSelectionCriteria struct {
//
//			FileNamePatterns    []string
//				An array of strings containing File Name Patterns
//
//			FilesOlderThan      time.Time
//				Match files with older modification date times
//
//			FilesNewerThan      time.Time
//				Match files with newer modification date times
//
//			RegularExp			*regexp.Regexp
//				Used to select file names with regular
//				expressions. If this parameter is NOT
//				equal to nil, file names will be
//				analyzed using MatchString().
//
//				Example:
//					RegularExp.MatchString("someFileName.txt")
//
//			SelectByFileMode    FilePermissionConfig
//				Match file mode (os.FileMode).
//
//			SelectCriterionModeFileSelectCriterionMode
//				Specifies 'AND' or 'OR' selection mode
//		}
//
//	  The FileSelectionCriteria type allows for
//	  configuration of single or multiple file selection
//	  criterion. The 'SelectCriterionMode' can be used to
//	  specify whether the file must match all, or any one,
//	  of the active file selection criterion.
//
//	  Elements of the FileSelectionCriteria are described
//	  below:
//
//	  FileNamePatterns		[]string
//	  	An array of strings which may define one or more
//	  	search patterns. If a file name matches any one
//	  	of the search pattern strings, it is deemed to be
//	  	a 'match' for the search pattern criterion.
//
//		Example Patterns:
//			FileNamePatterns = []string{"*.log"}
//			FileNamePatterns = []string{"current*.txt"}
//			FileNamePatterns = []string{"*.txt", "*.log"}
//
//		If this string array has zero length or if
//		all the strings are empty strings, then this
//		file search criterion is considered 'Inactive'
//		or 'Not Set'.
//
//
//		FilesOlderThan		time.Time
//			This date time type is compared to file
//			modification date times in order to determine
//			whether the file is older than the
//			'FilesOlderThan' file selection criterion. If
//			the file modification date time is older than
//			the 'FilesOlderThan' date time, that file is
//			considered a 'match' for this file selection
//			criterion.
//
//			If the value of 'FilesOlderThan' is set to
//			time zero, the default value for type
//			time.Time{}, then this file selection
//			criterion is considered to be 'Inactive' or
//			'Not Set'.
//
//		FilesNewerThan      time.Time
//			This date time type is compared to the file
//			modification date time in order to determine
//			whether the file is newer than the
//			'FilesNewerThan' file selection criterion. If
//			the file modification date time is newer than
//			the 'FilesNewerThan' date time, that file is
//			considered a 'match' for this file selection
//			criterion.
//
//			If the value of 'FilesNewerThan' is set to
//			time zero, the default value for type
//			time.Time{}, then this file selection
//			criterion is considered to be 'Inactive' or
//			'Not Set'.
//
//		RegularExp			*regexp.Regexp
//
//			Used to select file names with regular
//			expressions. If this parameter is NOT
//			equal to nil, file names will be
//			analyzed using MatchString().
//
//			Example:
//				RegularExp.MatchString("someFileName.txt")
//
//		SelectByFileMode	FilePermissionConfig
//			Type FilePermissionConfig encapsulates an
//			instance of os.FileMode. The file selection
//			criterion allows for the selection of files
//			by File Mode. File modes are compared to the
//			value of 'SelectByFileMode'. If the File Mode
//			for a given file is equal to the value of
//			'SelectByFileMode', that file is considered
//			to be a 'match' for this file selection
//			criterion. Examples for setting
//			SelectByFileMode are shown as follows:
//
//				fsc := FileSelectionCriteria{}
//
//				err =
//					fsc.SelectByFileMode.
//						SetByFileMode(
//							os.FileMode(0666))
//
//				err =
//					fsc.SelectByFileMode.
//						SetFileModeByTextCode(
//							"-r--r--r--")
//
//		SelectCriterionMode	FileSelectCriterionMode
//			This parameter selects the manner in which
//			the file selection criteria above are applied
//			in determining a 'match' for file selection
//			purposes. 'SelectCriterionMode' may be set to
//			one of two constant values:
//
//			FileSelectMode.ANDSelect()
//				File selected if all active selection
//				criteria are satisfied.
//
//				If this constant value is specified for
//				the file selection mode, then a given
//				file will not be judged as 'selected'
//				unless all the active selection criterion
//				are satisfied. In other words, if three
//				active search criterion are provided for
//				'FileNamePatterns', 'FilesOlderThan' and
//				'FilesNewerThan', then a file will NOT be
//				selected unless it has satisfied all three
//				criterion in this example.
//
//			FileSelectMode.ORSelect()
//				File selected if any active selection
//				criterion is satisfied.
//
//				If this constant value is specified for
//				the file selection mode, then a given
//				file will be selected if any one of the
//				active file selection criterion is
//				satisfied. In other words, if three
//				active search criterion are provided for
//				'FileNamePatterns', 'FilesOlderThan' and
//				'FilesNewerThan', then a file will be
//				selected if it satisfies any one of the
//				three criterion in this example.
//
//	skipTopLevelDirectory			bool
//
//		If this parameter is set to 'true', the parent or
//		top level directory identified by input parameter
//		'dMgr' will be skipped. This means no files will
//		be deleted from the parent or top level directory.
//
//		Be careful to ensure that parameters
//		'skipTopLevelDirectory' and 'scanSubDirectories'
//		are not in conflict. If 'skipTopLevelDirectory'
//		is set to 'true' and 'scanSubDirectories' is set
//		to 'false', an error will be returned.
//
//	scanSubDirectories				bool
//
//		If this parameter is set to 'true', it means that
//		child directories (a.k.a subdirectories) will be
//		searched and eligible files will be deleted from
//		subsidiary directories.
//
//		Conversely, if this parameter is set to 'false', no
//		files will be deleted from child directories (a.k.a
//		subdirectories).
//
//		Be careful to ensure that parameters
//		'skipTopLevelDirectory' and 'scanSubDirectories'
//		are not in conflict. If 'skipTopLevelDirectory'
//		is set to 'true' and 'scanSubDirectories' is set
//		to 'false', an error will be returned.
//
//	dMgrLabel						string
//
//		The name or label associated with input parameter
//		'dMgr' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "dMgr" will be
//		automatically applied.
//
//	deleteFileSelectLabel			string
//
//		The name or label used to describe the type of
//		files being deleted. This label will be used in
//		error messages returned by this method.
//
//		Example:
//			deleteFileSelectLabel = "Outdated files"
//
//		If this parameter is submitted as an empty
//		string, it will be automatically defaulted to a
//		value of "Target Files for Deletion".
//
//	errPrefDto						*ePref.ErrPrefixDto
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
//	deleteDirStats					DeleteDirFilesStats
//
//		If this method completes successfully, this
//		return parameter will be populated with
//		information and statistics on the file deletion
//		operation. This information includes the number
//		of files deleted.
//
//			type DeleteDirFilesStats struct {
//				TotalFilesProcessed        uint64
//				FilesDeleted               uint64
//				FilesDeletedBytes          uint64
//				FilesRemaining             uint64
//				FilesRemainingBytes        uint64
//				TotalSubDirectories        uint64
//				TotalDirsScanned           uint64
//				NumOfDirsWhereFilesDeleted uint64
//				DirectoriesDeleted         uint64
//			}
//
//	errs							[]error
//
//		An array of error objects.
//
//		If this method completes successfully, the
//		returned error array is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errPrefDto'.
//	 	The 'errPrefDto' text will be prefixed or
//	 	attached to the	beginning of the error message.
//
//		This error array may contain multiple errors.
//
//		An error array may be consolidated into a single
//		error using method StrMech.ConsolidateErrors()
func (dMgrHlpr *dirMgrHelper) deleteDirectoryTreeStats(
	dMgr *DirMgr,
	deleteFileSelectionCriteria FileSelectionCriteria,
	skipTopLevelDirectory bool,
	scanSubDirectories bool,
	dMgrLabel string,
	deleteSelectionLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	deleteDirStats DeleteDirFilesStats,
	errs []error) {

	if dMgrHlpr.lock == nil {
		dMgrHlpr.lock = new(sync.Mutex)
	}

	dMgrHlpr.lock.Lock()

	defer dMgrHlpr.lock.Unlock()

	funcName := "dirMgrHelper.deleteDirectoryTreeStats()"

	errs = make([]error, 0)

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		errs = append(errs, err)

		return deleteDirStats, errs
	}

	if len(dMgrLabel) == 0 {
		dMgrLabel = "dMgr"
	}

	_,
		_,
		err = new(dirMgrHelperPreon).
		validateDirMgr(
			dMgr,
			true,
			dMgrLabel,
			ePrefix)

	if err != nil {

		errs = append(errs, err)

		return deleteDirStats, errs
	}

	if skipTopLevelDirectory &&
		!scanSubDirectories {

		err := fmt.Errorf("%v\n"+
			"ERROR: Conflicted Input parameters!\n"+
			"skipTopLevelDirectory=true and scanSubDirectories=false.\n"+
			"Impossible combination!!\n",
			ePrefix.String())

		errs = append(errs, err)

		return deleteDirStats, errs
	}

	var err2 error
	osPathSepStr := string(os.PathSeparator)
	dirs := DirMgrCollection{}
	fh := FileHelper{}
	var nextDir DirMgr
	mainLoopIsDone := false
	isMatch := false
	isNewDir := false
	isTopLevelDir := true
	isFirstLoop := true

	err = dirs.AddDirMgr(
		*dMgr,
		ePrefix.XCpy("dirs<-dMgr"))

	if err != nil {

		err2 = fmt.Errorf("%v\n"+
			"Error returned when adding 'dirMgrCopy' to 'dirs' collection.\n"+
			"Error= \n%v\n",
			funcName,
			err.Error())

		errs = append(errs, err2)

		return deleteDirStats, errs
	}

	var nameDirEntries []os.DirEntry
	var nameFileInfo os.FileInfo

	var errStatus ArrayColErrorStatus

	for !mainLoopIsDone {

		if isFirstLoop {
			isTopLevelDir = true
			isFirstLoop = false
		} else {
			isTopLevelDir = false
		}

		nextDir,
			errStatus = dirs.
			PopFirstDirMgr(ePrefix.XCpy(
				"dirs"))

		if errStatus.ProcessingError != nil &&
			errStatus.IsIndexOutOfBounds == true {

			mainLoopIsDone = true
			break

		} else if errStatus.ProcessingError != nil {

			err2 = fmt.Errorf("%v\n"+
				"Error returned by dirs.PopFirstDirMgr().\n"+
				"Error='%v'\n",
				ePrefix.String(),
				errStatus.ProcessingError.Error())

			errs = append(errs, err2)

			return deleteDirStats, errs
		}

		deleteDirStats.TotalDirsScanned++

		isNewDir = true

		nameDirEntries,
			err2 = os.ReadDir(nextDir.absolutePath)

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Error returned by os.ReadDir(nextDir.absolutePath).\n"+
				"%v.absolutePath='%v'\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				"nextDir",
				nextDir.absolutePath,
				err2.Error())

			errs = append(errs, err)

			continue
		}

		for _, nameDirEntryInfo := range nameDirEntries {

			if nameDirEntryInfo.IsDir() {

				deleteDirStats.TotalSubDirectories++

				if !scanSubDirectories {
					continue
				}

				err = dirs.
					AddDirMgrByPathNameStr(
						nextDir.absolutePath+
							osPathSepStr+
							nameDirEntryInfo.Name(),
						ePrefix)

				if err != nil {

					err2 =
						fmt.Errorf("%v\n"+
							"Error returned by dirs.AddDirMgrByPathNameStr(newDirPathFileName).\n"+
							"newDirPathFileName='%v'\n"+
							"Error= \n%v\n",
							ePrefix.String(),
							nextDir.absolutePath+osPathSepStr+nameDirEntryInfo.Name(),
							err.Error())

					errs = append(errs, err2)
					continue
				}

			} else {
				// This is a file which is eligible for processing

				if isTopLevelDir && skipTopLevelDirectory {
					continue
				}

				nameFileInfo,
					err2 = nameDirEntryInfo.Info()

				if err2 != nil {

					err = fmt.Errorf("%v\n"+
						"Error: Error Returned by nameDirEntryInfo.Info().\n"+
						"The conversion of DirEntry to os.FileInfo Failed."+
						"Error= \n%v\n",
						ePrefix.String(),
						err2.Error())

					errs = append(errs, err)

					continue
				}

				deleteDirStats.TotalFilesProcessed++

				// This is not a directory. It is a file.
				// Determine if it matches the find file criteria.
				isMatch,
					err,
					_ =
					fh.FilterFileName(
						nameFileInfo,
						deleteFileSelectionCriteria,
						ePrefix.XCpy("nameDirEntryInfo"))

				if err != nil {

					err2 =
						fmt.Errorf("%v\n"+
							"Error returned by fh.FilterFileName(nameDirEntryInfo, %v).\n"+
							"%v directory searched='%v'\n"+
							"fileName='%v'\n"+
							"Error= \n%v\n",
							funcName,
							deleteSelectionLabel,
							dMgrLabel,
							dMgr.absolutePath,
							nameFileInfo.Name(),
							err.Error())

					errs = append(errs, err2)

					continue
				}

				if !isMatch {

					deleteDirStats.FilesRemaining++
					deleteDirStats.FilesRemainingBytes +=
						uint64(nameFileInfo.Size())

					continue

				} else {

					// We have a match, delete the file

					err = os.Remove(
						nextDir.absolutePath + osPathSepStr + nameFileInfo.Name())

					if err != nil {
						err2 = fmt.Errorf("%v\n"+
							"ERROR returned by os.Remove(pathFileName)\n"+
							"pathFileName='%v'\n"+
							"Error= \n%v\n",
							ePrefix.String(),
							nextDir.absolutePath+osPathSepStr+nameFileInfo.Name(),
							err.Error())

						deleteDirStats.FilesRemaining++
						deleteDirStats.FilesRemainingBytes +=
							uint64(nameFileInfo.Size())

						errs = append(errs, err2)

					} else {
						deleteDirStats.FilesDeleted++
						deleteDirStats.FilesDeletedBytes +=
							uint64(nameFileInfo.Size())

						if isNewDir {
							deleteDirStats.NumOfDirsWhereFilesDeleted++
						}

						isNewDir = false

					}
				}
			}

		} // End of nameDirEntryInfo := range nameFileInfos

	} // End of for !mainLoopIsDone

	nameDirEntries = make([]os.DirEntry, 0)

	return deleteDirStats, errs
}

// deleteFilesByNamePattern
//
// Receives a string defining a pattern to use in
// searching file names for all files in the directory
// identified by the input parameter 'dMgr'.
//
// During this search, files are deleted if the file name
// matches the pattern specified by input parameter,
// 'fileSearchPattern'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method deletes files in the directory specified
//	by the current instance of DirMgr. Only files in the
//	parent or top level directory identified by DirMgr
//	are eligible for deletion. No files in the DirMgr
//	subdirectory tree will be deleted.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dMgr							*DirMgr
//
//		A pointer to an instance of DirMgr. This method
//		will delete files in the directory
//		identified by this parameter.
//
//	fileSearchPattern				string
//
//		This string holds the pattern used to identify
//		files for deletion in the directory specified by
//		input parameter 'dMgr'.
//
//		Example Patterns
//			"*.*"
//			"*.txt"
//			"*My*.txt"
//
//	errPrefDto						*ePref.ErrPrefixDto
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
//	deleteDirStats					DeleteDirFilesStats
//
//		If this method completes successfully, this
//		return parameter will be populated with
//		information and statistics on the file deletion
//		operation. This information includes the number
//		of files deleted.
//
//			type DeleteDirFilesStats struct {
//				TotalFilesProcessed        uint64
//				FilesDeleted               uint64
//				FilesDeletedBytes          uint64
//				FilesRemaining             uint64
//				FilesRemainingBytes        uint64
//				TotalSubDirectories        uint64
//				TotalDirsScanned           uint64
//				NumOfDirsWhereFilesDeleted uint64
//				DirectoriesDeleted         uint64
//			}
//
//	errs							[]error
//
//		An array of error objects.
//
//		If this method completes successfully, the
//		returned error array is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errPrefDto'.
//	 	The 'errPrefDto' text will be prefixed or
//	 	attached to the	beginning of the error message.
//
//		This error array may contain multiple errors.
//
//		An error array may be consolidated into a single
//		error using method StrMech.ConsolidateErrors()
func (dMgrHlpr *dirMgrHelper) deleteFilesByNamePattern(
	dMgr *DirMgr,
	fileSearchPattern string,
	dMgrLabel string,
	fileSearchLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	deleteDirStats DeleteDirFilesStats,
	errs []error) {

	if dMgrHlpr.lock == nil {
		dMgrHlpr.lock = new(sync.Mutex)
	}

	dMgrHlpr.lock.Lock()

	defer dMgrHlpr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	funcName := "dirMgrHelper.deleteFilesByNamePattern()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		errs = append(errs, err)

		return deleteDirStats, errs
	}

	if len(dMgrLabel) == 0 {

		dMgrLabel = "dMgr"
	}

	dMgrHlprPreon := new(dirMgrHelperPreon)

	_,
		_,
		err = dMgrHlprPreon.
		validateDirMgr(
			dMgr,
			true,
			dMgrLabel,
			ePrefix)

	if err != nil {

		errs = append(errs, err)

		return deleteDirStats, errs
	}

	if err = new(strMechHelperErrors).
		validateInputString(
			fileSearchPattern,
			"fileSearchPattern",
			ePrefix); err != nil {

		errs = append(errs, err)

		return deleteDirStats, errs
	}

	deleteDirStats.TotalDirsScanned++

	var nameDirEntries []os.DirEntry
	osPathSepStr := string(os.PathSeparator)
	var isMatch bool
	var err2 error

	nameDirEntries,
		err2 = os.ReadDir(dMgr.absolutePath)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by os.ReadDir(dMgr.absolutePath).\n"+
			"%v.absolutePath='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			dMgrLabel,
			dMgr.absolutePath,
			err2.Error())

		errs = append(errs, err)
		return deleteDirStats, errs
	}

	var nameFileInfo os.FileInfo

	for _, nameDirEntry := range nameDirEntries {

		if nameDirEntry.IsDir() {

			// This is a Directory.
			// Skip it and proceed.
			deleteDirStats.TotalSubDirectories++

			continue

		} else {
			// This is a file.
			// Look for a match with selection criteria.

			nameFileInfo,
				err2 = nameDirEntry.Info()

			if err2 != nil {

				err = fmt.Errorf("%v\n"+
					"Error: Error Returned by nameDirEntry.Info().\n"+
					"The conversion of DirEntry to os.FileInfo Failed."+
					"%v= '%v'\n"+
					"FileName= '%v'\n"+
					"Error= \n%v\n",
					ePrefix.String(),
					dMgrLabel,
					dMgr.absolutePath,
					dMgr.absolutePath+osPathSepStr+nameFileInfo.Name(),
					err2.Error())

				errs = append(errs, err)

				continue
			}

			if !nameFileInfo.Mode().IsRegular() {

				err = fmt.Errorf("%v\n"+
					"Error: File Name is NOT classified as a 'Regular' File!\n"+
					"%v= '%v'\n"+
					"fileName= '%v'\n",
					ePrefix.String(),
					dMgrLabel,
					dMgr.absolutePath,
					dMgr.absolutePath+osPathSepStr+nameFileInfo.Name())

				errs = append(errs, err)

				deleteDirStats.FilesRemaining++

				deleteDirStats.FilesRemainingBytes +=
					uint64(nameFileInfo.Size())

				continue
			}

			// This is a Regular File

			isMatch,
				err = pf.Match(
				fileSearchPattern,
				nameFileInfo.Name())

			if err != nil {

				err2 = fmt.Errorf("%v\n"+
					"Error returned by (path/filepath) pf.Match(%v, fileName).\n"+
					"%v Directory Searched='%v'\n"+
					"%v='%v'\n"+
					"fileName='%v'\n"+
					"Error='%v'\n\n",
					ePrefix.String(),
					fileSearchLabel,
					dMgrLabel,
					dMgr.absolutePath,
					fileSearchLabel,
					fileSearchPattern,
					nameFileInfo.Name(),
					err.Error())

				errs = append(errs, err2)

				continue
			}

			deleteDirStats.TotalFilesProcessed++

			if !isMatch {
				// This file failed to match the
				// file selection criteria.
				deleteDirStats.FilesRemaining++
				deleteDirStats.FilesRemainingBytes += uint64(nameFileInfo.Size())
				continue

			} else {

				err = os.Remove(dMgr.absolutePath + osPathSepStr + nameFileInfo.Name())

				if err != nil {
					err2 = fmt.Errorf("%v\n"+
						"Error returned by os.Remove(pathFileName).\n"+
						"pathFileName='%v'\n"+
						"Error= \n%v\n",
						ePrefix.String(),
						dMgr.absolutePath+osPathSepStr+nameFileInfo.Name(),
						err.Error())

					deleteDirStats.FilesRemaining++
					deleteDirStats.FilesRemainingBytes -= uint64(nameFileInfo.Size())

					errs = append(errs, err2)
					continue
				}

				deleteDirStats.FilesDeleted++
				deleteDirStats.FilesDeletedBytes += uint64(nameFileInfo.Size())
			}
		}
	}

	nameDirEntries = make([]os.DirEntry, 0)

	return deleteDirStats, errs
}

// equal
//
// Compares two DirMgr objects to determine if they are
// equal is all respects.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dMgr1						*DirMgr
//
//		A pointer to an instance of DirMgr. If the
//		internal member data values are equal to the
//		corresponding internal member data values in
//		input parameter 'dMgr2', this method will
//		return a boolean value of 'true'.
//
//	dMgr2						*DirMgr
//
//		A pointer to an instance of DirMgr. If the
//		internal member data values are equal to the
//		corresponding internal member data values in
//		input parameter 'dMgr1', this method will
//		return a boolean value of 'true'.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		If the internal member data values in input
//		parameters 'dMgr1' and 'dMgr2' are equal in all
//		respects, this parameter returns a value of
//		'true'.
func (dMgrHlpr *dirMgrHelper) equal(
	dMgr1 *DirMgr,
	dMgr2 *DirMgr) bool {

	if dMgrHlpr.lock == nil {
		dMgrHlpr.lock = new(sync.Mutex)
	}

	dMgrHlpr.lock.Lock()

	defer dMgrHlpr.lock.Unlock()

	if dMgr1 == nil || dMgr2 == nil {
		return false
	}

	if dMgr1.isInitialized != dMgr2.isInitialized ||
		dMgr1.originalPath != dMgr2.originalPath ||
		dMgr1.path != dMgr2.path ||
		dMgr1.isPathPopulated != dMgr2.isPathPopulated ||
		dMgr1.doesPathExist != dMgr2.doesPathExist ||
		dMgr1.parentPath != dMgr2.parentPath ||
		dMgr1.isParentPathPopulated != dMgr2.isParentPathPopulated ||
		dMgr1.absolutePath != dMgr2.absolutePath ||
		dMgr1.isAbsolutePathPopulated != dMgr2.isAbsolutePathPopulated ||
		dMgr1.doesAbsolutePathExist != dMgr2.doesAbsolutePathExist ||
		dMgr1.isAbsolutePathDifferentFromPath != dMgr2.isAbsolutePathDifferentFromPath ||
		dMgr1.directoryName != dMgr2.directoryName ||
		dMgr1.volumeName != dMgr2.volumeName ||
		dMgr1.isVolumePopulated != dMgr2.isVolumePopulated {

		return false
	}

	if !dMgr1.actualDirFileInfo.Equal(&dMgr2.actualDirFileInfo) {
		return false
	}

	return true
}

// equalAbsolutePaths
//
// Compares the absolute paths for the input parameters
// 'dMgr1' and 'dMgr2'.
//
// If the two absolute paths are equal, the method
// returns 'true'.
//
// If the two absolute paths are NOT equal, this method
// returns 'false'.
//
// The comparison is NOT case-sensitive. In other words,
// both paths are first converted to lower case before
// making the comparison.
//
// If either the input parameter ('dMgr1') or the input
// parameter 'dMgr2' are uninitialized, a value of
// 'false' is returned.
//
// An absolute path is defined as follows:
//
//	"An absolute or full path points to the same location
//	in a file system, regardless of the current working
//	directory. To do that, it must include the root
//	directory."
//
//			Wikipedia
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dMgr1						*DirMgr
//
//		A pointer to an instance of DirMgr. The absolute
//		path contained in this instance is compared to
//		the absolute path encapsulated in the DirMgr
//		instance passed as 'dMgr2'. If the two absolute
//		paths are equal, this method returns 'true'.
//
//	dMgr2						*DirMgr
//
//		A pointer to an instance of DirMgr. The absolute
//		path contained in this instance is compared to
//		the absolute path encapsulated in the DirMgr
//		instance passed as 'dMgr1'. If the two absolute
//		paths are equal, this method returns 'true'.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		If the absolute paths contained in input
//		parameters 'dMgr1' and 'dMgr2' are equal, this
//		return parameter is set to 'true'.
//
//		If the absolute paths are not equal or, if one of
//		the DirMgr instances is uninitialized, a boolean
//		value of 'false' is returned.
func (dMgrHlpr *dirMgrHelper) equalAbsolutePaths(
	dMgr1 *DirMgr,
	dMgr2 *DirMgr) bool {

	if dMgrHlpr.lock == nil {
		dMgrHlpr.lock = new(sync.Mutex)
	}

	dMgrHlpr.lock.Lock()

	defer dMgrHlpr.lock.Unlock()

	dMgrHlprAtom := dirMgrHelperAtom{}

	if dMgr1 == nil || dMgr2 == nil {

		return false
	}

	if !dMgr1.isInitialized || !dMgr2.isInitialized {
		return false
	}

	_,
		_,
		_ = dMgrHlprAtom.doesDirectoryExist(
		dMgr1,
		PreProcPathCode.None(),
		"dMgr1",
		nil)

	lcDMgr1Path := strings.ToLower(dMgr1.absolutePath)

	_,
		_,
		_ = dMgrHlprAtom.doesDirectoryExist(
		dMgr2,
		PreProcPathCode.None(),
		"dMgr2",
		nil)

	lcDMgr2Path := strings.ToLower(dMgr2.absolutePath)

	if lcDMgr1Path != lcDMgr2Path {
		return false
	}

	return true
}

// equalPaths
//
// Compares the current instance of DirMgr with another
// DirMgr instance passed as an input parameter to
// determine if their directory and absolute paths are
// equal. Both Directory Path and absolute path must be
// equivalent for a comparison match.
//
// If the compared paths are equal, this method returns
// 'true'. If the paths are NOT equal, the method returns
// 'false'.
//
// The comparisons are NOT case-sensitive. In other
// words, all paths are converted to lower case before
// making the comparisons.
//
// If either the current DirMgr ('dMgr') or the input
// parameter 'dMgr2' are uninitialized, a value of
// 'false' is returned.
//
// An absolute path is defined as follows:
//
//	"An absolute or full path points to the same location
//	in a file system, regardless of the current working
//	directory. To do that, it must include the root
//	directory."
//
//			Wikipedia
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dMgr2						*DirMgr
//
//		A pointer to an instance of DirMgr. Both the
//		directory path and the absolute path contained in
//		this instance are compared to the directory path
//		and the absolute path encapsulated in the DirMgr
//		instance passed as 'dMgr2'. If the two paths are
//		equal, this method returns 'true'.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		If the directory path and the absolute path
//		contained in the current DirMgr instance are
//		equal to the directory path and absolute path
//		encapsulated in 'dMgr2' are equal, this return
//		parameter is set to 'true'.
//
//		If the directory and absolute paths are not equal
//		or, if one of the DirMgr instances is
//		uninitialized, a boolean value of 'false' is
//		returned.
func (dMgrHlpr *dirMgrHelper) equalPaths(
	dMgr *DirMgr,
	dMgr2 *DirMgr) bool {

	if dMgrHlpr.lock == nil {
		dMgrHlpr.lock = new(sync.Mutex)
	}

	dMgrHlpr.lock.Lock()

	defer dMgrHlpr.lock.Unlock()

	dMgrHlprAtom := dirMgrHelperAtom{}

	if !dMgr.isInitialized || !dMgr2.isInitialized {
		return false
	}

	_,
		_,
		_ = dMgrHlprAtom.doesDirectoryExist(
		dMgr,
		PreProcPathCode.None(),
		"",
		nil)

	lcDMgrPath := strings.ToLower(dMgr.absolutePath)

	_,
		_,
		_ = dMgrHlprAtom.doesDirectoryExist(
		dMgr2,
		PreProcPathCode.None(),
		"",
		nil)

	lcDMgr2Path := strings.ToLower(dMgr2.absolutePath)

	if lcDMgrPath != lcDMgr2Path {
		return false
	}

	lcDMgrPath = strings.ToLower(dMgr.path)
	lcDMgr2Path = strings.ToLower(dMgr2.path)

	if lcDMgrPath != lcDMgr2Path {
		return false
	}

	return true
}

// executeDirectoryFileOps
//
// Performs a one or more file operations on selected
// files in the parent or top level directory specified
// by the DirMgr input parameter 'sourceDMgr'.
//
// This method does NOT perform operations on
// subdirectories (a.k.a. the directory tree or child
// directories) of the 'sourceDMgr' instance.
//
// The files selected for this operation are determined
// by the file selection criteria configured by the
// user.
//
// The type of file operation performed is likewise
// configured by the user through input parameter,
// 'fileOps'.
//
// This method is designed to perform file operations on
// two directories; an input or source directory supplied
// by input parameter 'sourceDMgr', and an output or
// target directory specified by input parameter
// 'targetBaseDir'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	Depending on the type of File Operation code
//		submitted as an input parameter, this method may
//		be used to move, delete or copy files contained
//		in the top level or parent directory specified by
//		the DirMgr input parameter 'sourceDMgr'.
//
//	(2)	If the parent or top level directory specified by
//		'sourceDMgr' does not exist, an error will be
//		returned.
//
//	(3)	If the parent or top level directory specified by
//		'targetBaseDir' does not exist, an error will be
//		returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	sourceDMgr					*DirMgr
//
//		A pointer to an instance of DirMgr. This instance
//		is used as the source for file operations
//		performed on the 'sourceDMgr' top level or parent
//		directory path.
//
//		If the parent or top level directory specified by
//		'sourceDMgr' does not exist, an error will be
//		returned.
//
//	fileSelectCriteria			FileSelectionCriteria
//
//	  This input parameter should be configured with the
//	  desired file selection criteria. Files matching
//	  this criteria will be subject to the file
//	  operations specified by input parameter, 'fileOps'.
//
//		type FileSelectionCriteria struct {
//
//			FileNamePatterns    []string
//				An array of strings containing File Name Patterns
//
//			FilesOlderThan      time.Time
//				Match files with older modification date times
//
//			FilesNewerThan      time.Time
//				Match files with newer modification date times
//
//			RegularExp			*regexp.Regexp
//				Used to select file names with regular
//				expressions. If this parameter is NOT
//				equal to nil, file names will be
//				analyzed using MatchString().
//
//				Example:
//					RegularExp.MatchString("someFileName.txt")
//
//			SelectByFileMode    FilePermissionConfig
//				Match file mode (os.FileMode).
//
//			SelectCriterionModeFileSelectCriterionMode
//				Specifies 'AND' or 'OR' selection mode
//		}
//
//	  The FileSelectionCriteria type allows for
//	  configuration of single or multiple file selection
//	  criterion. The 'SelectCriterionMode' can be used to
//	  specify whether the file must match all, or any one,
//	  of the active file selection criterion.
//
//	  Elements of the FileSelectionCriteria are described
//	  below:
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
//			RegularExp			*regexp.Regexp
//
//				Used to select file names with regular
//				expressions. If this parameter is NOT
//				equal to nil, file names will be
//				analyzed using MatchString().
//
//				Example:
//					RegularExp.MatchString("someFileName.txt")
//
//			SelectByFileMode	FilePermissionConfig
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
//		If all the file selection criterion in the FileSelectionCriteria object are
//		'Inactive' or 'Not Set' (set to their zero or default values), then all the
//		files processed in the directory tree will be selected.
//
//			Example:
//			  fsc := FileSelectCriterionMode{}
//
//			  In this example, 'fsc' is NOT initialized. Therefore,
//			  all the selection criterion are 'Inactive'. Consequently,
//			  all the files encountered in the target directory during
//			  the search operation will be selected.
//
//		------------------------------------------------------------------------
//
//	fileOps						[]FileOperationCode
//
//		An array of file operations to be performed on
//		each selected file. Selected files are identified
//		by matching the file selection criteria specified
//		by input parameter, 'fileSelectCriteria'. See above.
//
//		The FileOperationCode type consists of the following
//		constants.
//
//		FileOperationCode(0).MoveSourceFileToDestinationFile() FileOperationCode = iota
//		  Moves the source file to the destination file and
//		  then deletes the original source file
//
//		FileOperationCode(0).DeleteDestinationFile()
//		  Deletes the Destination file if it exists
//
//		FileOperationCode(0).DeleteSourceFile()
//		  Deletes the Source file if it exists
//
//		FileOperationCode(0).DeleteSourceAndDestinationFiles
//		  Deletes both the Source and Destination files
//		  if they exist.
//
//		FileOperationCode(0).CopySourceToDestinationByHardLinkByIo()
//		  Copies the Source File to the Destination
//		  using two copy attempts. The first copy is
//		  by Hard Link. If the first copy attempt fails,
//		  a second copy attempt is initiated/ by creating
//		  a new file and copying the contents by 'io.Copy'.
//		  An error is returned only if both copy attempts
//		  fail. The source file is unaffected.
//
//		  See: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
//		FileOperationCode(0).CopySourceToDestinationByIoByHardLink()
//		  Copies the Source File to the Destination
//		  using two copy attempts. The first copy is
//		  by 'io.Copy' which creates a new file and copies
//		  the contents to the new file. If the first attempt
//		  fails, a second copy attempt is initiated using
//		  'copy by hard link'. An error is returned only
//		  if both copy attempts fail. The source file is
//		  unaffected.
//
//		  See: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
//		FileOperationCode(0).CopySourceToDestinationByHardLink()
//		  Copies the Source File to the Destination
//		  using one copy mode. The only copy attempt
//		  utilizes 'Copy by Hard Link'. If this fails
//		  an error is returned.  The source file is
//		  unaffected.
//
//		FileOperationCode(0).CopySourceToDestinationByIo()
//		  Copies the Source File to the Destination
//		  using only one copy mode. The only copy
//		  attempt is initiated using 'Copy by IO' or
//		  'io.Copy'.  If this fails an error is returned.
//		  The source file is unaffected.
//
//		FileOperationCode(0).CreateSourceDir()
//		  Creates the Source Directory
//
//		FileOperationCode(0).CreateSourceDirAndFile()
//		  Creates the Source Directory and File
//
//		FileOperationCode(0).CreateSourceFile()
//		  Creates the Source File
//
//		FileOperationCode(0).CreateDestinationDir()
//		  Creates the Destination Directory
//
//		FileOperationCode(0).CreateDestinationDirAndFile()
//		  Creates the Destination Directory and File
//
//		FileOperationCode(0).CreateDestinationFile()
//		  Creates the Destination File
//
//	targetBaseDir				*DirMgr
//
//		A pointer to an instance of DirMgr. The top level
//		or parent directory of this DirMgr instance serves
//		as the destination or target for file operations
//		performed on the parent directory of input
//		parameter 'sourceDMgr'.
//
//		If the parent or top level directory specified by
//		'targetBaseDir' does not exist, an error will be
//		returned.
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
//		'targetBaseDir' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "targetBaseDir" will
//		be automatically applied.
//
//	fileSelectLabel				string
//
//		The name or label used to describe the type of
//		files selected for file operations. This label
//		will be used in error messages returned by this
//		method.
//
//		Example:
//			fileSelectLabel = "Files For Deletion"
//
//		If this parameter is submitted as an empty
//		string, it will be automatically defaulted to a
//		value of "Target Files".
//
//	fileOpsLabel				string
//
//		The name or label used to describe the file
//		operations being performed. This label will be
//		used in error messages returned by this method.
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
//	errs						[]error
//
//		An array of error objects.
//
//		If this method completes successfully, the
//		returned error array is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errPrefDto'.
//	 	The 'errPrefDto' text will be prefixed or
//	 	attached to the	beginning of the error message.
//
//		This error array may contain multiple errors.
//
//		An error array may be consolidated into a single
//		error using method StrMech.ConsolidateErrors()
func (dMgrHlpr *dirMgrHelper) executeDirectoryFileOps(
	sourceDMgr *DirMgr,
	fileSelectCriteria FileSelectionCriteria,
	fileOps []FileOperationCode,
	targetBaseDir *DirMgr,
	sourceDMgrLabel string,
	targetDMgrLabel string,
	fileSelectLabel string,
	fileOpsLabel string,
	errPrefDto *ePref.ErrPrefixDto) (errs []error) {

	if dMgrHlpr.lock == nil {
		dMgrHlpr.lock = new(sync.Mutex)
	}

	dMgrHlpr.lock.Lock()

	defer dMgrHlpr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	funcName := "dirMgrHelper.executeDirectoryFileOps()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		errs = append(errs, err)

		return errs
	}

	if len(sourceDMgrLabel) == 0 {

		sourceDMgrLabel = "sourceDMgr"
	}

	dMgrHlprPreon := new(dirMgrHelperPreon)

	_,
		_,
		err = dMgrHlprPreon.
		validateDirMgr(
			sourceDMgr,
			true,
			sourceDMgrLabel,
			ePrefix.XCpy(
				sourceDMgrLabel))

	if err != nil {

		errs = append(errs, err)

		return errs
	}

	if len(targetDMgrLabel) == 0 {

		targetDMgrLabel = "targetBaseDir"
	}

	_,
		_,
		err = dMgrHlprPreon.
		validateDirMgr(
			targetBaseDir,
			true,
			targetDMgrLabel,
			ePrefix.XCpy(
				sourceDMgrLabel))

	if err != nil {

		errs = append(errs, err)

		return errs
	}

	if len(fileOpsLabel) == 0 {

		fileOpsLabel = "File Operations"
	}

	if len(fileSelectLabel) == 0 {

		fileSelectLabel = "Files For Deletion"
	}

	if len(fileOps) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: The input parameter '%v' is a ZERO LENGTH ARRAY!\n",
			ePrefix.String(),
			fileOpsLabel)

		errs = append(errs, err)
		return errs
	}

	var nameDirEntries []os.DirEntry

	var err2 error

	nameDirEntries,
		err2 = os.ReadDir(sourceDMgr.absolutePath)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by os.ReadDir(sourceDMgr.absolutePath).\n"+
			"%v.absolutePath='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			sourceDMgrLabel,
			sourceDMgr.absolutePath,
			err2.Error())

		errs = append(errs, err)

		return errs
	}

	var nameFileInfo os.FileInfo
	fh := new(FileHelper)
	var isMatch bool
	var fileOp FileOps
	srcFileNameExt := ""

	for _, nameDirEntryInfo := range nameDirEntries {

		if nameDirEntryInfo.IsDir() {
			continue
		}

		// Must be a file - process it!
		// Determine if it matches the find file criteria.

		nameFileInfo,
			err2 = nameDirEntryInfo.Info()

		isMatch,
			err2,
			_ = fh.FilterFileName(
			nameFileInfo,
			fileSelectCriteria,
			ePrefix)

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Error returned by FileHelper{}.FilterFileName(nameFInfo, %v).\n"+
				"%v Directory Searched='%v'\n"+
				"fileName='%v'\n"+
				"Error='%v'\n",
				ePrefix.String(),
				fileSelectLabel,
				sourceDMgrLabel,
				sourceDMgr.absolutePath,
				nameFileInfo.Name(),
				err2.Error())

			errs = append(errs, err)

			return errs
		}

		if !isMatch {

			continue

		}

		// Must be a match - this is a 'selected' file!
		srcFileNameExt = nameFileInfo.Name()

		err2 = new(FileOperationsNanobot).
			setFileOpsByDirAndFileNameStr(
				&fileOp,
				sourceDMgr.absolutePath,
				srcFileNameExt,
				targetBaseDir.absolutePath,
				srcFileNameExt,
				ePrefix.XCpy("fileOp<-"))

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Error returned by FileOps.NewByDirStrsAndFileNameExtStrs()\n"+
				"%v Source Path='%v'\n"+
				"srcFileNameExt='%v'\n"+
				"%v Destination Directory='%v'\n"+
				"Destination File='%v'\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				sourceDMgrLabel,
				sourceDMgr.absolutePath,
				srcFileNameExt,
				targetDMgrLabel,
				targetBaseDir.absolutePath,
				srcFileNameExt,
				err2.Error())

			errs = append(errs, err2)

			return errs
		}

		for i := 0; i < len(fileOps); i++ {

			err2 = fileOp.ExecuteFileOperation(
				fileOps[i],
				ePrefix)

			if err2 != nil {

				err = fmt.Errorf("%v\n"+
					"Error returned by fileOp.ExecuteFileOperation(fileOps[%v])\n"+
					"FileOps='%v'\n"+
					"Error= \n%v\n",
					ePrefix.String(),
					i,
					fileOps[i].String(),
					err2.Error())

				// Store the error and continue processing
				// file operations.
				errs = append(errs, err)
			}

		} // END OF: for i := 0; i < len(fileOps); i++

	} // END OF: for _, nameDirEntryInfo := range nameDirEntries

	nameDirEntries = make([]os.DirEntry, 0)

	return errs
}

// executeDirectoryTreeOps
//
// Performs File Operations on 'selected' files in the
// directory tree identified by the input parameter,
// 'sourceDMgr'.
//
// The 'sourceDMgr' path therefore serves as the source
// parent directory for file operations performed on the
// directory tree. Designated file operations will
// therefore be performed on all files in the parent
// directory as well as all files in all subdirectories.
//
// The destination for these file operations is specified
// by input parameter 'targetBaseDir'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	If the parent or top level directory specified by
//		'sourceDMgr' does not exist, an error will be
//		returned.
//
//	(2)	If the parent or top level directory specified by
//		'targetBaseDir' does not exist, an error will be
//		returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	sourceDMgr					*DirMgr
//
//		A pointer to an instance of DirMgr. This instance
//		defines the parent directory for file operations
//		performed on the parent directory and all
//		subdirectories in the directory tree specified by
//		'sourceDMgr'.
//
//		If the parent or top level directory specified by
//		'sourceDMgr' does not exist on disk, an error
//		will be returned.
//
//	fileSelectCriteria			FileSelectionCriteria
//
//	  This input parameter should be configured with the
//	  desired file selection criteria. Files matching
//	  this criteria will be subject to the file
//	  operations specified by input parameter, 'fileOps'.
//
//		type FileSelectionCriteria struct {
//
//			FileNamePatterns    []string
//				An array of strings containing File Name Patterns
//
//			FilesOlderThan      time.Time
//				Match files with older modification date times
//
//			FilesNewerThan      time.Time
//				Match files with newer modification date times
//
//			RegularExp			*regexp.Regexp
//				Used to select file names with regular
//				expressions. If this parameter is NOT
//				equal to nil, file names will be
//				analyzed using MatchString().
//
//				Example:
//					RegularExp.MatchString("someFileName.txt")
//
//			SelectByFileMode    FilePermissionConfig
//				Match file mode (os.FileMode).
//
//			SelectCriterionModeFileSelectCriterionMode
//				Specifies 'AND' or 'OR' selection mode
//		}
//
//	  The FileSelectionCriteria type allows for
//	  configuration of single or multiple file selection
//	  criterion. The 'SelectCriterionMode' can be used to
//	  specify whether the file must match all, or any one,
//	  of the active file selection criterion.
//
//	  Elements of the FileSelectionCriteria are described
//	  below:
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
//			RegularExp			*regexp.Regexp
//
//				Used to select file names with regular
//				expressions. If this parameter is NOT
//				equal to nil, file names will be
//				analyzed using MatchString().
//
//				Example:
//					RegularExp.MatchString("someFileName.txt")
//
//			SelectByFileMode	FilePermissionConfig
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
//		If all the file selection criterion in the FileSelectionCriteria object are
//		'Inactive' or 'Not Set' (set to their zero or default values), then all the
//		files processed in the directory tree will be selected.
//
//			Example:
//			  fsc := FileSelectCriterionMode{}
//
//			  In this example, 'fsc' is NOT initialized. Therefore,
//			  all the selection criterion are 'Inactive'. Consequently,
//			  all the files encountered in the target directory during
//			  the search operation will be selected.
//
//		------------------------------------------------------------------------
//
//	fileOps						[]FileOperationCode
//
//		An array of file operations to be performed on
//		each selected file. Selected files are identified
//		by matching the file selection criteria specified
//		by input parameter, 'fileSelectCriteria'. See above.
//
//		The FileOperationCode type consists of the following
//		constants.
//
//		FileOperationCode(0).MoveSourceFileToDestinationFile() FileOperationCode = iota
//		  Moves the source file to the destination file and
//		  then deletes the original source file
//
//		FileOperationCode(0).DeleteDestinationFile()
//		  Deletes the Destination file if it exists
//
//		FileOperationCode(0).DeleteSourceFile()
//		  Deletes the Source file if it exists
//
//		FileOperationCode(0).DeleteSourceAndDestinationFiles
//		  Deletes both the Source and Destination files
//		  if they exist.
//
//		FileOperationCode(0).CopySourceToDestinationByHardLinkByIo()
//		  Copies the Source File to the Destination
//		  using two copy attempts. The first copy is
//		  by Hard Link. If the first copy attempt fails,
//		  a second copy attempt is initiated/ by creating
//		  a new file and copying the contents by 'io.Copy'.
//		  An error is returned only if both copy attempts
//		  fail. The source file is unaffected.
//
//		  See: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
//		FileOperationCode(0).CopySourceToDestinationByIoByHardLink()
//		  Copies the Source File to the Destination
//		  using two copy attempts. The first copy is
//		  by 'io.Copy' which creates a new file and copies
//		  the contents to the new file. If the first attempt
//		  fails, a second copy attempt is initiated using
//		  'copy by hard link'. An error is returned only
//		  if both copy attempts fail. The source file is
//		  unaffected.
//
//		  See: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
//		FileOperationCode(0).CopySourceToDestinationByHardLink()
//		  Copies the Source File to the Destination
//		  using one copy mode. The only copy attempt
//		  utilizes 'Copy by Hard Link'. If this fails
//		  an error is returned.  The source file is
//		  unaffected.
//
//		FileOperationCode(0).CopySourceToDestinationByIo()
//		  Copies the Source File to the Destination
//		  using only one copy mode. The only copy
//		  attempt is initiated using 'Copy by IO' or
//		  'io.Copy'.  If this fails an error is returned.
//		  The source file is unaffected.
//
//		FileOperationCode(0).CreateSourceDir()
//		  Creates the Source Directory
//
//		FileOperationCode(0).CreateSourceDirAndFile()
//		  Creates the Source Directory and File
//
//		FileOperationCode(0).CreateSourceFile()
//		  Creates the Source File
//
//		FileOperationCode(0).CreateDestinationDir()
//		  Creates the Destination Directory
//
//		FileOperationCode(0).CreateDestinationDirAndFile()
//		  Creates the Destination Directory and File
//
//		FileOperationCode(0).CreateDestinationFile()
//		  Creates the Destination File
//
//	targetBaseDir				*DirMgr
//
//		A pointer to an instance of DirMgr. The top level
//		or parent directory of this DirMgr instance serves
//		as the destination or target for file operations
//		performed on the parent directory of input
//		parameter 'sourceDMgr'.
//
//		If the parent or top level directory specified by
//		'targetBaseDir' does not exist on disk, an error
//		will be returned.
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
//	targetBaseDirLabel			string
//
//		The name or label associated with input parameter
//		'targetBaseDir' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "targetBaseDir" will
//		be automatically applied.
//
//	fileOpsLabel				string
//
//		The name or label used to describe the file
//		operations being performed. This label will be
//		used in error messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "File Operations"
//		will be automatically applied.
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
//	errs						[]error
//
//		An array of error objects.
//
//		If this method completes successfully, the
//		returned error array is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errPrefDto'.
//	 	The 'errPrefDto' text will be prefixed or
//	 	attached to the	beginning of the error message.
//
//		This error array may contain multiple errors.
//
//		An error array may be consolidated into a single
//		error using method StrMech.ConsolidateErrors()
func (dMgrHlpr *dirMgrHelper) executeDirectoryTreeOps(
	sourceDMgr *DirMgr,
	fileSelectCriteria FileSelectionCriteria,
	fileOps []FileOperationCode,
	targetBaseDir *DirMgr,
	sourceDMgrLabel string,
	targetBaseDirLabel string,
	fileOpsLabel string,
	errPrefDto *ePref.ErrPrefixDto) (errs []error) {

	if dMgrHlpr.lock == nil {
		dMgrHlpr.lock = new(sync.Mutex)
	}

	dMgrHlpr.lock.Lock()

	defer dMgrHlpr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	funcName := "dirMgrHelper.executeDirectoryTreeOps()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		errs = append(errs, err)

		return errs
	}

	if len(sourceDMgrLabel) == 0 {

		sourceDMgrLabel = "sourceDMgrLabel"
	}

	dMgrHlprPreon := new(dirMgrHelperPreon)

	_,
		_,
		err = dMgrHlprPreon.
		validateDirMgr(
			sourceDMgr,
			true,
			sourceDMgrLabel,
			ePrefix.XCpy(
				sourceDMgrLabel))

	if err != nil {

		errs = append(errs, err)

		return errs
	}

	if len(targetBaseDirLabel) == 0 {

		targetBaseDirLabel = "targetBaseDir"
	}

	_,
		_,
		err = dMgrHlprPreon.
		validateDirMgr(
			targetBaseDir,
			true,
			targetBaseDirLabel,
			ePrefix.XCpy(
				sourceDMgrLabel))

	if err != nil {

		errs = append(errs, err)

		return errs
	}

	if len(fileOpsLabel) == 0 {

		fileOpsLabel = "File Operations"
	}

	var err2 error

	if len(fileOps) == 0 {

		err2 = fmt.Errorf("%v\n"+
			"Error: The input parameter '%v' is a ZERO LENGTH ARRAY!\n",
			ePrefix.String(),
			fileOpsLabel)

		errs = append(errs, err2)
		return errs
	}

	dirOp := new(DirTreeOp).New()
	dirOp.CallingFunc = funcName + "\n"
	dirOp.FileOps = append(dirOp.FileOps, fileOps...)

	dirOp.TargetBaseDir, err = new(DirMgr).
		New(targetBaseDir.absolutePath, ePrefix)

	if err != nil {

		err2 = fmt.Errorf("%v\n"+
			"Error returned by dirOp.TargetBaseDir = DirMgr{}.New(%v.absolutePath)\n"+
			"%v.absolutePath='%v'\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			targetBaseDirLabel,
			targetBaseDirLabel,
			targetBaseDir.absolutePath,
			err.Error())

		errs = append(errs, err2)
		return errs
	}

	dirOp.SourceBaseDir,
		err = new(DirMgr).
		New(
			sourceDMgr.absolutePath,
			ePrefix)

	if err != nil {

		err2 = fmt.Errorf("%v\n"+
			"Error returned by dirOp.SourceBaseDir = DirMgr{}.New(%v.absolutePath)\n"+
			"%v.absolutePath='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			sourceDMgrLabel,
			sourceDMgrLabel,
			sourceDMgr.absolutePath,
			err.Error())

		errs = append(errs, err2)
		return errs
	}

	dirOp.FileSelectCriteria = fileSelectCriteria

	err = pf.Walk(
		sourceDMgr.absolutePath,
		new(dirMgrHelperAtom).executeFileOpsOnFoundFiles(
			&dirOp,
			ePrefix.XCpy("dirOp")))

	if err != nil {
		err2 = fmt.Errorf("%v\n"+
			"Error returned by (path/filepath) pf.Walk("+
			"%v.absolutePath, dMgrHlpr.executeFileOpsOnFoundFiles(dirOp)).\n"+
			"%v.absolutePath='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			sourceDMgrLabel,
			sourceDMgrLabel,
			sourceDMgr.absolutePath,
			err.Error())

		errs = append(errs, dirOp.ErrReturns...)
		errs = append(errs, err2)
		return errs
	}

	return dirOp.ErrReturns
}

// findDirectoryTreeFiles
//
// A multifunctional helper method which can be used to
// scan a parent directory or an entire directory tree to
// locate files which match the file selection criteria.
//
// Files matching the selection criteria defined by input
// parameter 'fileSelectCriteria', will be identified in
// the returned instance of type DirectoryTreeInfo
// ('dTreeInfo').
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	If the parent or top level directory specified by
//	input parameter 'targetBaseDir' does not exist on
//	disk, an error will be returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	targetBaseDir				*DirMgr
//
//		A pointer to an instance of DirMgr. This DirMgr
//		instance identifies the parent directory and
//		directory tree where the search for files matching
//		the file selection criteria will be conducted.
//
//		If the parent or top level directory specified by
//		'targetBaseDir' does not exist on disk, an error
//		will be returned.
//
//	fileSelectCriteria			FileSelectionCriteria
//
//		This input parameter should be configured with the
//		desired file selection criteria. Files matching
//		this criteria will be subject to the file
//		operations specified by input parameter, 'fileOps'.
//
//		type FileSelectionCriteria struct {
//
//			FileNamePatterns    []string
//				An array of strings containing File Name Patterns
//
//			FilesOlderThan      time.Time
//				Match files with older modification date times
//
//			FilesNewerThan      time.Time
//				Match files with newer modification date times
//
//			RegularExp			*regexp.Regexp
//				Used to select file names with regular
//				expressions. If this parameter is NOT
//				equal to nil, file names will be
//				analyzed using MatchString().
//
//				Example:
//					RegularExp.MatchString("someFileName.txt")
//
//			SelectByFileMode    FilePermissionConfig
//				Match file mode (os.FileMode).
//
//			SelectCriterionModeFileSelectCriterionMode
//				Specifies 'AND' or 'OR' selection mode
//		}
//
//		The FileSelectionCriteria type allows for
//		configuration of single or multiple file selection
//		criterion. The 'SelectCriterionMode' can be used to
//		specify whether the file must match all, or any one,
//		of the active file selection criterion.
//
//	  Elements of the FileSelectionCriteria are described
//	  below:
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
//			RegularExp			*regexp.Regexp
//
//				Used to select file names with regular
//				expressions. If this parameter is NOT
//				equal to nil, file names will be
//				analyzed using MatchString().
//
//				Example:
//					RegularExp.MatchString("someFileName.txt")
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
//		If all the file selection criterion in the FileSelectionCriteria object are
//		'Inactive' or 'Not Set' (set to their zero or default values), then all the
//		files processed in the directory tree will be selected.
//
//			Example:
//			  fsc := FileSelectCriterionMode{}
//
//			  In this example, 'fsc' is NOT initialized. Therefore,
//			  all the selection criterion are 'Inactive'. Consequently,
//			  all the files encountered in the target directory during
//			  the search operation will be selected.
//
//		------------------------------------------------------------------------
//
//	skipTopLevelDirectory		bool
//
//		When this parameter is set to 'true', the search
//		for matching files will NOT be conducted in the
//		top level or parent directory.
//
//		If this parameter is set to 'true' and parameter
//		'scanSubDirectories' is also set to 'false', the
//		parameters are in conflict and an error will be
//		returned.
//
//	scanSubDirectories			bool
//
//		When this parameter is to 'true', the search for
//		matching files will include the subdirectories
//		below the parent directory in the specified
//		directory tree.
//
//		If this parameter is set to 'false' and parameter
//		'skipTopLevelDirectory' is also set to 'true',
//		the parameters are in conflict and an error will
//		be returned.
//
//	targetBaseDirLabel			string
//
//		The name or label associated with input parameter
//		'targetBaseDir' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "targetBaseDir" will
//		be automatically applied.
//
//	fileSelectLabel				string
//
//		The name or label used to describe the file
//		operations being performed. This label will be
//		used in error messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "File Selection Criteria"
//		will be automatically applied.
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
//	dTreeInfo					DirectoryTreeInfo
//
//		If this method completes successfully, this
//		return parameter will be populated with
//		information and statistics on the file search
//		operation.
//
//		DirectoryTreeInfo - A structure used
//		to 'Find' files in a directory specified
//		by 'StartPath'. The file search will be
//		filtered by a 'FileSelectCriteria' object.
//
//		'FileSelectCriteria' is a FileSelectionCriteria
//		type which contains FileNamePatterns strings and
//		'FilesOlderThan' or 'FilesNewerThan' date time
//		parameters which can be used as file selection
//		criteria.
//
//		type DirectoryTreeInfo struct {
//			StartPath          string
//			Directories        DirMgrCollection
//			FoundFiles         FileMgrCollection
//			ErrReturns         []error
//			FileSelectCriteria FileSelectionCriteria
//		}
//
//	errs						[]error
//
//		An array of error objects.
//
//		If this method completes successfully, the
//		returned error array is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errPrefDto'.
//	 	The 'errPrefDto' text will be prefixed or
//	 	attached to the	beginning of the error message.
//
//		This error array may contain multiple errors.
//
//		An error array may be consolidated into a single
//		error using method StrMech.ConsolidateErrors()
func (dMgrHlpr *dirMgrHelper) findDirectoryTreeFiles(
	targetBaseDir *DirMgr,
	fileSelectCriteria FileSelectionCriteria,
	skipTopLevelDirectory bool,
	scanSubDirectories bool,
	targetBaseDirLabel string,
	fileSelectLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	dTreeInfo DirectoryTreeInfo,
	errs []error) {

	if dMgrHlpr.lock == nil {
		dMgrHlpr.lock = new(sync.Mutex)
	}

	dMgrHlpr.lock.Lock()

	defer dMgrHlpr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	funcName := "dirMgrHelper.findDirectoryTreeFiles()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		errs = append(errs, err)

		return dTreeInfo, errs
	}

	if len(targetBaseDirLabel) == 0 {

		targetBaseDirLabel = "targetBaseDir"
	}

	_,
		_,
		err = new(dirMgrHelperPreon).
		validateDirMgr(
			targetBaseDir,
			true, // Path MUST exist on disk
			targetBaseDirLabel,
			ePrefix.XCpy(
				targetBaseDirLabel))

	if err != nil {

		errs = append(errs, err)

		return dTreeInfo, errs
	}

	if len(fileSelectLabel) == 0 {

		fileSelectLabel = "File Selection Criteria"
	}

	if skipTopLevelDirectory &&
		!scanSubDirectories {

		err := fmt.Errorf("%v\n"+
			"ERROR: Conflicted Input parameters!\n"+
			"skipTopLevelDirectory=true and scanSubDirectories=false.\n"+
			"Impossible combination!!\n",
			ePrefix.String())

		errs = append(errs, err)
		return dTreeInfo, errs
	}

	var err2 error
	osPathSepStr := string(os.PathSeparator)
	fh := FileHelper{}
	var nextDir DirMgr
	isMatch := false
	isTopLevelDir := true

	// Adds a copy of targetBaseDir
	err = dTreeInfo.Directories.AddDirMgr(
		*targetBaseDir,
		ePrefix.XCpy("targetBaseDirCopy"))

	if err != nil {

		errs = append(errs, err)

		return dTreeInfo, errs
	}

	dTreeCnt := 1

	var nameDirEntries []os.DirEntry
	var nameFileInfo os.FileInfo

	var errStatus = ArrayColErrorStatus{}

	for i := 0; i < dTreeCnt; i++ {

		if i == 0 {
			isTopLevelDir = true
		} else {
			isTopLevelDir = false
		}

		nextDir,
			errStatus = new(dirMgrCollectionHelper).
			peekOrPopAtIndex(
				&dTreeInfo.Directories,
				i,
				false, // Do NOT Delete 'idx'
				ePrefix.XCpy(
					fmt.Sprintf(
						"dMgrs[%v]",
						i)))

		if errStatus.ProcessingError != nil {

			errs = append(errs,
				fmt.Errorf("%v",
					errStatus.ProcessingError.Error()))

			continue
		}

		nameDirEntries,
			err2 = os.ReadDir(nextDir.absolutePath)

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Error returned by os.ReadDir(nextDir.absolutePath).\n"+
				"%v.absolutePath='%v'\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				"nextDir",
				nextDir.absolutePath,
				err2.Error())

			errs = append(errs, err)

			continue
		}

		for _, nameDirEntry := range nameDirEntries {

			if nameDirEntry.IsDir() {

				if !scanSubDirectories {
					continue
				}

				err = dTreeInfo.Directories.
					AddDirMgrByKnownPathDirName(
						nextDir.absolutePath,
						nameDirEntry.Name(),
						ePrefix.XCpy(
							"nextDir.absolutePath"))

				if err != nil {

					err2 =
						fmt.Errorf("%v\n"+
							"Error returned by dirs.AddDirMgrByKnownPathDirName(newDirPathFileName).\n"+
							"newDirPathFileName='%v'\n"+
							"Error=\n%v\n",
							ePrefix.String(),
							nextDir.absolutePath+osPathSepStr+nameDirEntry.Name(),
							err.Error())

					errs = append(errs, err2)
					continue
				}

				dTreeCnt++

			} else {
				// This is a file which is eligible for processing

				if isTopLevelDir && skipTopLevelDirectory {
					continue
				}

				nameFileInfo,
					err2 = nameDirEntry.Info()

				if err2 != nil {

					err = fmt.Errorf("%v\n"+
						"Error: Error Returned by nameDirEntry.Info().\n"+
						"The conversion of DirEntry to os.FileInfo Failed.\n"+
						"nextDir= '%v'\n"+
						"fileName='%v'\n"+
						"Error= \n%v\n",
						ePrefix.String(),
						nextDir.absolutePath,
						nextDir.absolutePath+osPathSepStr+nameFileInfo.Name(),
						err2.Error())

					errs = append(errs, err)

					continue
				}

				if !nameFileInfo.Mode().IsRegular() {

					err = fmt.Errorf("%v\n"+
						"Error: File Name is NOT classified as a 'Regular' File!\n"+
						"nextDir= '%v'\n"+
						"fileName='%v'\n",
						ePrefix.String(),
						nextDir.absolutePath,
						nextDir.absolutePath+osPathSepStr+nameFileInfo.Name())

					errs = append(errs, err)

					continue
				}

				// This is not a directory. It is a file.
				// Determine if it matches the find file criteria.
				isMatch,
					err,
					_ =
					fh.FilterFileName(
						nameFileInfo,
						fileSelectCriteria,
						ePrefix.XCpy("nameFileInfo"))

				if err != nil {

					err2 =
						fmt.Errorf("%v\n"+
							"Error returned by fh.FilterFileName(nameDirEntry, %v).\n"+
							"%v directory searched='%v'\n"+
							"fileName='%v'\n"+
							"Error= \n%v\n",
							funcName,
							fileSelectLabel,
							targetBaseDirLabel,
							targetBaseDir.absolutePath,
							nameFileInfo.Name(),
							err.Error())

					errs = append(errs, err2)

					continue
				}

				if !isMatch {

					continue

				} else {

					// We have a match, save file to dTreeInfo

					err = dTreeInfo.FoundFiles.AddFileMgrByDirFileNameExt(
						nextDir,
						nameFileInfo.Name(),
						ePrefix.XCpy("nameDirEntry"))

					if err != nil {

						err2 = fmt.Errorf("%v\n"+
							"ERROR returned by dTreeInfo.FoundFiles."+
							"AddFileMgrByDirFileNameExt(nextDir, fileNameExt)\n"+
							"nextDir= '%v'\n"+
							"fileNameExt= '%v'\n"+
							"Error= \n%v\n",
							funcName,
							nextDir.absolutePath,
							nameFileInfo.Name(),
							err.Error())

						errs = append(errs, err2)

					}
				}
			}

		} // End of nameDirEntry := range nameFileInfos

	} // End of for i := 0; i < dTreeCnt; i++ {

	nameDirEntries = make([]os.DirEntry, 0)

	if len(dTreeInfo.Directories.dirMgrs) > 0 &&
		skipTopLevelDirectory {

		_, _ = dTreeInfo.Directories.PopFirstDirMgr(
			ePrefix.XCpy(
				"dTreeInfo.Directories"))

	}

	return dTreeInfo, errs
}

// findDirectoryTreeStats
//
// Scans the parent directory or the entire directory
// tree to calculate and return information and
// statistics pertaining to files and subdirectories.
//
// Unlike similar methods, this method will scan the
// directory specified by input parameter 'targetBaseDir'
// and return statistical information on ALL files and
// subdirectories residing in that directory tree.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	If the parent or top level directory specified by
//	input parameter 'targetBaseDir' does not exist on
//	disk, an error will be returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	targetBaseDir				*DirMgr
//
//		A pointer to an instance of DirMgr. This DirMgr
//		instance identifies the parent directory and
//		directory tree where statistical information
//		on ALL files and subdirectories residing in that
//		directory tree will be accumulated.
//
//		If the parent or top level directory specified by
//		'targetBaseDir' does not exist on disk, an error
//		will be returned.
//
//	skipTopLevelDirectory		bool
//
//		When this parameter is set to 'true', the search
//		for files will NOT be conducted in the top level
//		or parent directory.
//
//		If this parameter is set to 'false' and parameter
//		'scanSubDirectories' is also set to 'false', an
//		error will be returned.
//
//	scanSubDirectories			bool
//
//		When this parameter is to 'true', the search for
//		files will include the subdirectories below the
//		parent directory in the specified directory tree.
//
//		If this parameter is set to 'false' and parameter
//		'skipTopLevelDirectory' is also set to 'false',
//		an error will be returned.
//
//	targetBaseDirLabel			string
//
//		The name or label associated with input parameter
//		'targetBaseDir' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "targetBaseDir" will
//		be automatically applied.
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
//	dTreeStats					DirectoryStatsDto
//
//		The DirectoryStatsDto structure is used to
//		accumulate and disseminate statistical
//		information relating to a specific directory
//		tree.
//
//		type DirectoryStatsDto struct {
//
//			dMgr DirMgr
//				Identifies the parent directory associated with
//				this directory information.
//
//			numOfFiles uint64
//				The number of files (all types) residing
//				within this directory ('dMgr').
//
//			numOfSubDirs uint64
//				The number of subdirectories residing
//				within this directory
//
//			numOfBytes uint64
//				The total number of bytes for all files
//				contained in this directory.
//				isInitialized bool
//		}
//
//		Type DirectoryStatsDto contains public methods
//		for retrieving the specified directory statistics
//		and information.
//
//		If this method completes successfully, this
//		returned instance of DirectoryStatsDto will
//		contain information on files and directories
//		contained in the directory tree specified by
//		input parameter 'targetBaseDir'.
//
//	errs						[]error
//
//		An array of error objects.
//
//		If this method completes successfully, the
//		returned error array is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errPrefDto'.
//	 	The 'errPrefDto' text will be prefixed or
//	 	attached to the	beginning of the error message.
//
//		This error array may contain multiple errors.
//
//		An error array may be consolidated into a single
//		error using method StrMech.ConsolidateErrors()
func (dMgrHlpr *dirMgrHelper) findDirectoryTreeStats(
	targetBaseDir *DirMgr,
	skipTopLevelDirectory bool,
	scanSubDirectories bool,
	targetBaseDirLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	dTreeStats DirectoryStatsDto,
	errs []error) {

	if dMgrHlpr.lock == nil {
		dMgrHlpr.lock = new(sync.Mutex)
	}

	dMgrHlpr.lock.Lock()

	defer dMgrHlpr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	funcName := "dirMgrHelper.findDirectoryTreeStats()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		errs = append(errs, err)

		return dTreeStats, errs
	}

	if len(targetBaseDirLabel) == 0 {

		targetBaseDirLabel = "targetBaseDir"
	}

	_,
		_,
		err = new(dirMgrHelperPreon).
		validateDirMgr(
			targetBaseDir,
			true, // Path MUST exist on disk
			targetBaseDirLabel,
			ePrefix.XCpy(
				targetBaseDirLabel))

	if err != nil {

		errs = append(errs, err)

		return dTreeStats, errs
	}

	var err2 error
	dirs := DirMgrCollection{}
	var nameDirEntries []os.DirEntry
	var nameFileInfo os.FileInfo
	var nextDir DirMgr
	mainLoopIsDone := false
	isFirstLoop := true
	isTopLevelDir := true
	osPathSepStr := string(os.PathSeparator)

	err = dirs.AddDirMgr(
		*targetBaseDir,
		ePrefix.XCpy("targetBaseDir"))

	if err != nil {

		errs = append(errs, err)

		return dTreeStats, errs
	}

	var errStatus = ArrayColErrorStatus{}

	for !mainLoopIsDone {

		if isFirstLoop {
			isTopLevelDir = true
			isFirstLoop = false
		} else {
			isTopLevelDir = false
		}

		nextDir,
			errStatus = dirs.PopFirstDirMgr(ePrefix.XCpy(
			"dirs"))

		if errStatus.ProcessingError != nil &&
			errStatus.IsIndexOutOfBounds == true {

			mainLoopIsDone = true
			break

		} else if errStatus.ProcessingError != nil {

			err2 = fmt.Errorf("%v\n"+
				"Error returned by dirs.PopFirstDirMgr().\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				errStatus.ProcessingError.Error())

			errs = append(errs, err2)

			return dTreeStats, errs
		}

		nameDirEntries,
			err2 = os.ReadDir(nextDir.absolutePath)

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Error returned by os.ReadDir(nextDir.absolutePath).\n"+
				"%v.absolutePath='%v'\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				"nextDir",
				nextDir.absolutePath,
				err2.Error())

			errs = append(errs, err)

			continue
		}

		for _, nameDirEntry := range nameDirEntries {

			if nameDirEntry.IsDir() {
				// This is a directory
				err = dirs.
					AddDirMgrByKnownPathDirName(
						nextDir.absolutePath,
						nameDirEntry.Name(),
						ePrefix.XCpy(
							"nextDir.absolutePath"))

				if err != nil {
					errs = append(errs, err2)
					continue
				}

				dTreeStats.numOfSubDirs++

			} else {

				if isTopLevelDir && skipTopLevelDirectory {
					continue
				}

				// This is a file

				nameFileInfo,
					err2 = nameDirEntry.Info()

				if err2 != nil {

					err = fmt.Errorf("%v\n"+
						"Error: Error Returned by nameDirEntry.Info().\n"+
						"The conversion of DirEntry to os.FileInfo Failed.\n"+
						"nextDir= '%v'\n"+
						"FileName= '%v'\n"+
						"Error= \n%v\n",
						ePrefix.String(),
						nextDir.absolutePath,
						nextDir.absolutePath+osPathSepStr+nameFileInfo.Name(),
						err2.Error())

					errs = append(errs, err)

					continue
				}

				if !nameFileInfo.Mode().IsRegular() {

					err = fmt.Errorf("%v\n"+
						"Error: File Name is NOT classified as a 'Regular' File!\n"+
						"nextDir= '%v'\n"+
						"fileName='%v'\n",
						ePrefix.String(),
						nextDir.absolutePath,
						nextDir.absolutePath+osPathSepStr+nameFileInfo.Name())

					errs = append(errs, err)

					continue
				}

				dTreeStats.numOfFiles++
				dTreeStats.numOfBytes += uint64(nameFileInfo.Size())
			}
		} // for _, nameDirEntry := range nameFileInfos

		if isTopLevelDir && !scanSubDirectories {
			mainLoopIsDone = true
			break
		}

	} // for !mainLoopIsDone

	nameDirEntries = make([]os.DirEntry, 0)

	return dTreeStats, errs
}

// findFilesByNamePattern
//
// Searches for matching files in the top level or parent
// directory specified by input parameter 'targetBaseDir'.
//
// Only the parent directory specified by 'targetBaseDir'
// will be searched for matching files - NOT the
// subdirectories in the directory tree.
//
// Files matching the search pattern specified by input
// parameter 'fileSearchPattern' will be selected, stored
// and returned as a type FileMgrCollection.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	If the parent or top level directory specified by
//	input parameter 'targetBaseDir' does not exist on
//	disk, an error will be returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	targetBaseDir				*DirMgr
//
//		A pointer to an instance of DirMgr. The top level
//		or parent directory of this DirMgr instance will
//		be searched for files matching the search criteria
//		specified by input parameter 'fileSearchPattern'.
//
//		If the parent or top level directory specified by
//		'targetBaseDir' does not exist, an error will be
//		returned.
//
//	fileSearchPattern				string
//
//		This string holds the pattern used to identify
//		files for deletion in the directory specified by
//		input parameter 'dMgr'.
//
//		Example 'fileSearchPattern' strings
//
//		*.*             will match all files in directory.
//		*.html          will match  anyfilename.html
//		a*              will match  appleJack.txt
//		j????row.txt    will match  j1x34row.txt
//		data[0-9]*      will match 	data123.csv
//
//	targetBaseDirLabel			string
//
//		The name or label associated with input parameter
//		'targetBaseDir' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "targetBaseDir" will
//		be automatically applied.
//
//	fileSearchPatternLabel		string
//
//		The name or label used to describe the file
//		search criteria used to select matching files.
//		This label will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "File Search Pattern"
//		will be automatically applied.
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
//	FileMgrCollection
//
//		If this method completes successfully, the
//		returned error Type will be populated with the
//		files matching the search criteria specified by
//		input parameter 'fileSearchPattern'.
//
//		FileMgrCollection - Manages a collection of
//		FileMgr instances. These file manager instances
//		will be populated with information on the
//		matching files identified by the file search
//		operation. Information on files contained in the
//		FileMgrCollection can be retrieved through
//		methods on the FileMgrCollection type.
//
//			type FileMgrCollection struct {
//				fileMgrs []FileMgr
//			}
//
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
func (dMgrHlpr *dirMgrHelper) findFilesByNamePattern(
	targetBaseDir *DirMgr,
	fileSearchPattern string,
	targetBaseDirLabel string,
	fileSearchPatternLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	FileMgrCollection,
	error) {

	if dMgrHlpr.lock == nil {
		dMgrHlpr.lock = new(sync.Mutex)
	}

	dMgrHlpr.lock.Lock()

	defer dMgrHlpr.lock.Unlock()

	fileMgrCol := FileMgrCollection{}.New()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "dirMgrHelper.findFilesByNamePattern() "

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		return fileMgrCol, err
	}

	if len(targetBaseDirLabel) == 0 {

		targetBaseDirLabel = "targetBaseDir"
	}

	_,
		_,
		err = new(dirMgrHelperPreon).
		validateDirMgr(
			targetBaseDir,
			true, // Path MUST exist on disk
			targetBaseDirLabel,
			ePrefix.XCpy(
				targetBaseDirLabel))

	if err != nil {

		return fileMgrCol, err
	}

	if len(fileSearchPatternLabel) == 0 {

		fileSearchPatternLabel = "File Search Pattern"
	}

	if err = new(strMechHelperErrors).
		validateInputString(
			fileSearchPattern,
			fileSearchPatternLabel,
			ePrefix.XCpy(fileSearchPatternLabel)); err != nil {

		return fileMgrCol, err

	}

	var isMatch bool
	var nameDirEntries []os.DirEntry
	var nameFileInfo os.FileInfo
	var err2 error
	errs := make([]error, 0, 10)
	osPathSepStr := string(os.PathSeparator)

	nameDirEntries,
		err2 = os.ReadDir(targetBaseDir.absolutePath)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by os.ReadDir(targetBaseDir.absolutePath).\n"+
			"%v.absolutePath='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			targetBaseDirLabel,
			targetBaseDir.absolutePath,
			err2.Error())

		return fileMgrCol, err
	}

	for _, nameDirEntry := range nameDirEntries {

		if nameDirEntry.IsDir() {
			continue

		} else {
			// This is a file.

			nameFileInfo,
				err2 = nameDirEntry.Info()

			if err2 != nil {

				err = fmt.Errorf("%v\n"+
					"Error: Error Returned by nameDirEntry.Info().\n"+
					"The conversion of DirEntry to os.FileInfo Failed."+
					"Error= \n%v\n",
					ePrefix.String(),
					err2.Error())

				errs = append(errs, err)

				continue
			}

			if !nameFileInfo.Mode().IsRegular() {

				err = fmt.Errorf("%v\n"+
					"Error: File Name is NOT classified as a 'Regular' File!\n"+
					"%v= '%v'\n"+
					"fileName= '%v'\n",
					ePrefix.String(),
					targetBaseDirLabel,
					targetBaseDir.absolutePath,
					targetBaseDir.absolutePath+osPathSepStr+nameFileInfo.Name())

				errs = append(errs, err)

				continue
			}

			// Check for pattern match.
			isMatch, err = pf.Match(fileSearchPattern, nameDirEntry.Name())

			if err != nil {

				err2 = fmt.Errorf("%v\n"+
					"Error returned by fp.Match(%v, fileName).\n"+
					"directorySearched='%v'\n"+
					"%v='%v'\n"+
					"fileName='%v'\n"+
					"Error= \n%v\n ",
					ePrefix.String(),
					fileSearchPatternLabel,
					targetBaseDir.absolutePath,
					fileSearchPatternLabel,
					fileSearchPattern,
					nameDirEntry.Name(),
					err.Error())

				errs = append(errs, err2)
				continue
			}

			if !isMatch {

				continue

			} else {
				// This file is a match. Process it.
				err = fileMgrCol.AddFileMgrByFileInfo(
					targetBaseDir.absolutePath,
					nameFileInfo,
					ePrefix)

				if err != nil {

					err2 = fmt.Errorf("%v\n"+
						"Error returned by fileMgrCol.AddFileMgrByFileInfo(%v.absolutePath, nameDirEntry).\n"+
						"Directory='%v'\n"+
						"FileName='%v'\n"+
						"Error= \n%v\n",
						funcName,
						targetBaseDirLabel,
						targetBaseDir.absolutePath,
						nameDirEntry.Name(),
						err.Error())

					errs = append(errs, err2)

					continue
				}
			}
		}
	}

	nameDirEntries = make([]os.DirEntry, 0)

	return fileMgrCol, new(StrMech).ConsolidateErrors(errs)
}

// getAbsolutePathElements
//
// Receives an instance of DirMgr and returns the drive
// and directory specifications as an array of
// strings.
//
// Example:
//
//	DirMgr Path = "D:\ADir\BDir\CDir\EDir"
//
//	Returned pathElements string array:
//
//		pathElements[0] = "D:"
//		pathElements[1] = "ADir"
//		pathElements[2] = "BDir"
//		pathElements[3] = "CDir"
//		pathElements[4] = "DDir"
//		pathElements[4] = "EDir"
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dMgr						*DirMgr
//
//		A pointer to an instance of DirMgr.
//
//		The drive and directory specifications in the
//		directory path contained in this instance will
//		be returned as individual elements in an array
//		of strings.
//
//		The directory path for this DirMgr instance is
//		NOT required to exist on disk.
//
//	dMgrLabel					string
//
//		The name or label associated with input parameter
//		'dMgr', which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "dMgr" will be
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
//	pathElements				[]string
//
//		If this method completes successfully, this
//		return parameter will be populated with an array
//		of strings containing the elements of the
//		directory specified by input parameter 'dMgr'.
//
//		Example:
//
//			dMgr Path = "D:\ADir\BDir\CDir\EDir"
//
//			Returned pathElements string array:
//
//				pathElements[0] = "D:"
//				pathElements[1] = "ADir"
//				pathElements[2] = "BDir"
//				pathElements[3] = "CDir"
//				pathElements[4] = "DDir"
//				pathElements[4] = "EDir"
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
func (dMgrHlpr *dirMgrHelper) getAbsolutePathElements(
	dMgr *DirMgr,
	dMgrLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	pathElements []string,
	err error) {

	if dMgrHlpr.lock == nil {
		dMgrHlpr.lock = new(sync.Mutex)
	}

	dMgrHlpr.lock.Lock()

	defer dMgrHlpr.lock.Unlock()

	pathElements = make([]string, 0, 50)

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"dirMgrHelper.getAbsolutePathElements()",
		"")

	if err != nil {

		return pathElements, err
	}

	if len(dMgrLabel) == 0 {

		dMgrLabel = "dMgr"
	}

	_,
		_,
		err = new(dirMgrHelperPreon).
		validateDirMgr(
			dMgr,
			false, // Path is NOT required to exist on disk
			dMgrLabel,
			ePrefix.XCpy(
				dMgrLabel))

	if err != nil {

		return pathElements, err
	}

	var absolutePath string

	absolutePath = dMgr.absolutePath

	absolutePath = strings.Replace(absolutePath, "\\", "/", -1)

	pathElements = strings.Split(absolutePath, "/")

	return pathElements, err
}

// getDirectoryTree
//
// Returns a DirMgrCollection containing all the
// subdirectories in the directory tree identified by the
// input parameter 'dMgr'.
//
// The returned DirMgrCollection will always contain the
// parent directory at the top of the array (index=0).
// Therefore, if no errors are encountered, the returned
// DirMgrCollection will always consist of at least one
// directory.
//
// If subdirectories are found, then the returned
// DirMgrCollection will contain more than one directory.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	If the parent or top level directory specified by
//	input parameter 'dMgr' does not exist on disk, an
//	error will be returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dMgr						*DirMgr
//
//		A pointer to an instance of DirMgr specifying a
//		directory. This method will identify the 'dMgr'
//		parent directory and all existing subdirectories
//		through a returned instance of DirMgrCollection.
//
//		If the parent or top level directory specified by
//		'dMgr' does not exist on disk, an error	will be
//		returned.
//
//	dMgrLabel					string
//
//		The name or label associated with input parameter
//		'dMgr', which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "dMgr" will be
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
//	dirMgrs 					DirMgrCollection
//
//		If this method completes successfully, this
//		method will return an instance of DirMgrCollection
//		populated with an array of 'DirMgr' objects
//		identifying the parent directory and all
//		subdirectories specified by input parameter
//		'dMgr'.
//
//			type DirMgrCollection struct {
//				dirMgrs []DirMgr
//			}
//
//	errs						[]error
//
//		An array of error objects.
//
//		If this method completes successfully, the
//		returned error array is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errPrefDto'.
//	 	The 'errPrefDto' text will be prefixed or
//	 	attached to the	beginning of the error message.
//
//		This error array may contain multiple errors.
//
//		An error array may be consolidated into a single
//		error using method StrMech.ConsolidateErrors()
func (dMgrHlpr *dirMgrHelper) getDirectoryTree(
	dMgr *DirMgr,
	dMgrLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	dirMgrs DirMgrCollection,
	errs []error) {

	if dMgrHlpr.lock == nil {
		dMgrHlpr.lock = new(sync.Mutex)
	}

	dMgrHlpr.lock.Lock()

	defer dMgrHlpr.lock.Unlock()

	dirMgrs = new(DirMgrCollection).New()

	var err, err2 error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"dirMgrHelper.getDirectoryTree()",
		"")

	if err != nil {

		errs = append(errs, err)

		return dirMgrs, errs
	}

	if len(dMgrLabel) == 0 {

		dMgrLabel = "dMgr"
	}

	_,
		_,
		err = new(dirMgrHelperPreon).
		validateDirMgr(
			dMgr,
			true, // Path MUST exist on disk
			dMgrLabel,
			ePrefix.XCpy(
				dMgrLabel))

	if err != nil {

		errs = append(errs, err)

		return dirMgrs, errs
	}

	err = dirMgrs.AddDirMgr(
		*dMgr,
		ePrefix.XCpy("dMgr"))

	if err != nil {

		errs = append(errs, err)

		return dirMgrs, errs
	}

	var fh = new(FileHelper)
	var nameDirEntries []os.DirEntry

	maxLen := dirMgrs.GetNumOfDirs()

	for i := 0; i < maxLen; i++ {

		nameDirEntries,
			err2 = os.ReadDir(dirMgrs.dirMgrs[i].absolutePath)

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Error returned by os.ReadDir(dirMgrs.dirMgrs[%v].absolutePath).\n"+
				"%v.absolutePath='%v'\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				i,
				"dirMgrs.dirMgrs[i]",
				dirMgrs.dirMgrs[i].absolutePath,
				err2.Error())

			errs = append(errs, err)

			continue
		}

		for _, nameDirEntry := range nameDirEntries {

			if nameDirEntry.IsDir() {

				newDirPathFileName :=
					fh.JoinPathsAdjustSeparators(
						dirMgrs.dirMgrs[i].absolutePath, nameDirEntry.Name())

				err = dirMgrs.AddDirMgrByPathNameStr(
					newDirPathFileName,
					ePrefix)

				if err != nil {

					err2 =
						fmt.Errorf("%v\n"+
							"Error returned by dirMgrs.AddDirMgrByPathNameStr(newDirPathFileName).\n"+
							"dirPtr='%v'\n"+
							"Error='%v'\n",
							ePrefix.String(),
							newDirPathFileName,
							err.Error())

					errs = append(errs, err2)
					continue
				}

				maxLen = dirMgrs.GetNumOfDirs()
			}
		} // END OF: for _, nameDirEntry := range nameDirEntries

	} // END OF: for i := 0; i < maxLen; i++

	return dirMgrs, errs
}

// getParentDirMgr
//
// Returns a new Directory Manager instance which
// represents the parent path for the input Directory
// Manager, 'dMgr'. The 'dMgr' absolute path is used
// in extracting the parent path in the form of a
// new Directory Manager instance.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dMgr						*DirMgr
//
//		A pointer to an instance of DirMgr. The absolute
//		path specified by this instance will be analyzed
//		to extract and return the parent path.
//
//	dMgrLabel					string
//
//		The name or label associated with input parameter
//		'dMgr', which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "dMgr" will be
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
//	dirMgrParent				DirMgr
//
//		If this method completes successfully, this
//		returned instance of DirMgr will contain the
//		absolute parent path extracted from input
//		parameter 'dMgr'.
//
//	hasParent					bool
//
//		If input parameter 'dMgr' has a parent path, this
//		returned boolean value will be set to 'true'.
//
//		Otherwise, this value is set to 'false' signaling
//		that 'dMgr' does not have a parent path and the
//		absolute path specified by 'dMgr' is a parent path.
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
func (dMgrHlpr *dirMgrHelper) getParentDirMgr(
	dMgr *DirMgr,
	dMgrLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	dirMgrParent DirMgr,
	hasParent bool,
	err error) {

	if dMgrHlpr.lock == nil {
		dMgrHlpr.lock = new(sync.Mutex)
	}

	dMgrHlpr.lock.Lock()

	defer dMgrHlpr.lock.Unlock()

	hasParent = false

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"dirMgrHelper.getParentDirMgr()",
		"")

	if err != nil {

		return dirMgrParent, hasParent, err
	}

	if len(dMgrLabel) == 0 {

		dMgrLabel = "dMgr"
	}

	_,
		_,
		err = new(dirMgrHelperPreon).
		validateDirMgr(
			dMgr,
			false, // Path is NOT required to exist on disk
			dMgrLabel,
			ePrefix.XCpy(
				dMgrLabel))

	if err != nil {

		return dirMgrParent, hasParent, err
	}

	if len(dMgr.parentPath) == 0 {

		err = new(dirMgrHelperBoson).
			copyDirMgrs(
				&dirMgrParent,
				dMgr,
				ePrefix.XCpy("dMgr Copy Errors"))

		hasParent = false

		return dirMgrParent, hasParent, err

	} else {
		hasParent = true
	}

	var err2 error

	dirMgrParent,
		err2 = new(DirMgr).New(
		dMgr.parentPath,
		ePrefix.XCpy(
			"dMgr.parentPath"))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by DirMgr{}.New(%v.parentPath).\n"+
			"%v.parentPath=%v\n"+
			"Error='%v'\n",
			ePrefix.String(),
			dMgrLabel,
			dMgrLabel,
			dMgr.parentPath,
			err2.Error())

		hasParent = true

		dirMgrParent = DirMgr{}

		return dirMgrParent, hasParent, err
	}

	return dirMgrParent, hasParent, err
}

// moveDirectoryTree
//
// Moves all subdirectories and files plus files in the
// parent 'dMgr' directory to a target directory tree
// specified by input parameter 'targetDMgr'. If
// successful, the parent directory, 'dMgr, will be
// deleted along with the entire sub-directory tree.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	This method will delete the entire directory tree
//		identified by 'dMgr' along with ALL the files in
//		that directory tree. This operation includes the
//		top level or parent directory identified by input
//		parameter 'dMgr'.
//
//	(2)	To move the subdirectory tree, thereby excluding
//		the parent directory identified by 'dMgr' see
//		method:
//
//			dirMgrHelper.moveSubDirectoryTree()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dMgr						*DirMgr
//
//		A pointer to an instance of DirMgr. This instance
//		identifies the source directory tree which will
//		be moved to a target directory identified by input
//		parameter 'targetDMgr'. After source files are
//		copied to this target directory, the entire
//		directory tree identified by 'dMgr' along with
//		ALL the files in that directory tree will be
//		deleted.
//
//		ALL the files in this directory tree will be
//		moved to the target directory specified by input
//		parameter 'targetDMgr'.
//
//		If the directory specified by 'dMgr' does not
//		exist, an error will be returned.
//
//	targetDMgr					*DirMgr
//
//		A pointer to an instance of DirMgr. Source files
//		selected in the source directory ('dMgr') will be
//		copied to a corresponding directory in the target
//		directory tree specified by this input parameter
//		('targetDMgr').
//
//		If this target directory does not exist, this
//		method will attempt to create it.
//
//	dMgrLabel					string
//
//		The name or label associated with input parameter
//		'dMgr', which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "dMgr" will be
//		automatically applied.
//
//	targetDMgrLabel				string
//
//		The name or label associated with input parameter
//		'targetDMgr' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "targetDMgr" will
//		be automatically applied.
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
//		structure will contain information and statistics
//		describing the outcome of the
//		'move directory tree' operation.
//
//		type DirectoryMoveStats struct {
//			TotalSrcFilesProcessed   uint64
//			SourceFilesMoved         uint64
//			SourceFileBytesMoved     uint64
//			SourceFilesRemaining     uint64
//			SourceFileBytesRemaining uint64
//			TotalDirsProcessed       uint64
//			DirsCreated              uint64
//			SourceOriginalSubDirs      uint64
//			SourceDirWasDeleted      bool
//			Errors             error
//		}
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
func (dMgrHlpr *dirMgrHelper) moveDirectoryTree(
	dMgr *DirMgr,
	targetDMgr *DirMgr,
	dMgrLabel string,
	targetDMgrLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	dirMoveStats DirectoryMoveStats,
	nonfatalErrs []error,
	fatalErr error) {

	dMgrHlpr.lock.Lock()

	defer dMgrHlpr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "dirMgrHelper.moveDirectoryTree()"

	ePrefix,
		fatalErr = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if fatalErr != nil {

		return dirMoveStats, nonfatalErrs, fatalErr
	}

	if len(dMgrLabel) == 0 {
		dMgrLabel = "dMgr"
	}

	var dMgrHlprPreon = new(dirMgrHelperPreon)

	_,
		_,
		fatalErr = dMgrHlprPreon.
		validateDirMgr(
			dMgr,
			true, // Path MUST exist on disk
			dMgrLabel,
			ePrefix.XCpy(
				dMgrLabel))

	if fatalErr != nil {

		return dirMoveStats, nonfatalErrs, fatalErr
	}

	if len(targetDMgrLabel) == 0 {

		targetDMgrLabel = "targetDMgr"
	}

	_,
		_,
		fatalErr = dMgrHlprPreon.
		validateDirMgr(
			targetDMgr,
			false, // Path is NOT required to exist on disk
			targetDMgrLabel,
			ePrefix.XCpy(
				targetDMgrLabel))

	if fatalErr != nil {

		return dirMoveStats, nonfatalErrs, fatalErr
	}

	var err2 error

	fileSelectCriteria := FileSelectionCriteria{}
	var errs2 []error
	var dTreeCopyStats DirTreeCopyStats

	dTreeCopyStats,
		_,
		errs2,
		err2 =
		new(dirMgrHelperNanobot).
			copyDirectoryTree(
				dMgr,
				targetDMgr,
				false, // returnCopiedFilesList
				false, // skipTopLevelDirectory
				false, // copyEmptyTargetDirectory
				true,  // copyRegularFiles
				true,  // copySymLinkFiles
				true,  // copyOtherNonRegularFiles
				fileSelectCriteria,
				dMgrLabel,
				targetDMgrLabel,
				ePrefix)

	if len(errs2) > 0 {

		nonfatalErrs = append(nonfatalErrs, errs2...)

	}

	if err2 != nil {

		fatalErr = fmt.Errorf("%v\n"+
			"Errors occurred while copying directory tree to target directory.\n"+
			"The source directory WAS NOT DELETED!\n"+
			"%v Source Directory='%v'\n"+
			"%v Target Directory='%v'\n"+
			"Fatal Error = \n%v\n",
			funcName,
			dMgrLabel,
			dMgr.absolutePath,
			targetDMgrLabel,
			targetDMgr.absolutePath,
			err2.Error())

		return dirMoveStats, nonfatalErrs, fatalErr
	}

	dirMoveStats.TotalDirsProcessed =
		dTreeCopyStats.TotalDirsScanned

	dirMoveStats.TargetDirsCreated =
		dTreeCopyStats.DirsCreated

	dirMoveStats.SourceOriginalSubDirs =
		dTreeCopyStats.TotalDirsScanned - 1

	dirMoveStats.SourceFilesRemaining =
		dTreeCopyStats.FilesNotCopied

	dirMoveStats.SourceFileBytesRemaining =
		dTreeCopyStats.FileBytesNotCopied

	if dirMoveStats.SourceFilesRemaining > 0 {

		fatalErr = fmt.Errorf("%v\n"+
			"Error: Some of the files designated to be moved\n"+
			"to the target directory, were NOT copied!\n"+
			"Therefore the source directory WILL NOT BE DELETED!\n"+
			"Number of Files NOT Copied='%v'\n"+
			"%v Source Directory='%v'\n"+
			"%v Target Directory= '%v'\n",
			ePrefix.String(),
			dTreeCopyStats.FilesNotCopied,
			dMgrLabel,
			dMgr.absolutePath,
			targetDMgrLabel,
			targetDMgr.absolutePath)

		return dirMoveStats, nonfatalErrs, fatalErr
	}

	dirMoveStats.TotalSrcFilesProcessed =
		dTreeCopyStats.TotalFilesProcessed

	err2 = new(dirMgrHelperMolecule).
		lowLevelDeleteDirectoryAll(
			dMgr,
			dMgrLabel,
			ePrefix)

	if err2 != nil {

		fatalErr = fmt.Errorf("%v\n"+
			"Files were copied successfuly to target directory.\n"+
			"However, errors occurred while deleting the source directory tree.\n"+
			"%v.absolutePath='%v'\n"+
			"Error= \n%v\n",
			funcName,
			dMgrLabel,
			dMgr.absolutePath,
			err2.Error())

		return dirMoveStats, nonfatalErrs, fatalErr
	}

	dirMoveStats.SourceDirWasDeleted = true

	dirMoveStats.SourceFilesMoved =
		dTreeCopyStats.FilesCopied

	dirMoveStats.SourceFileBytesMoved =
		dTreeCopyStats.FileBytesCopied

	return dirMoveStats, nonfatalErrs, fatalErr
}

// moveSubDirectoryTree
//
// Moves all subdirectories in the 'dMgr' tree to the
// 'targetDMgr' subdirectory tree.
//
// Moves all subdirectories and their constituent files
// from the source or parent directory 'DirMgr' to a
// target directory tree specified by input parameter
// 'targetDMgr'.
//
// If this method completes successfully, all
// subdirectories and files in the source directory tree
// will be deleted.
//
// The source or parent directory identified by 'DirMgr'
// and the files within 'DirMgr' will NOT be deleted.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete the entire subdirectory tree in
// the parent directory identified by the current instance
// of DirMgr. However, the source or parent directory for
// the current instance of DirMgr, and its constituent
// files, will NOT be deleted.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dMgr						*DirMgr
//
//		A pointer to an instance of DirMgr. This instance
//		identifies the source parent directory from which
//		all subdirectories and files will be moved to the
//		target directory tree identified by input
//		parameter 'targetDMgr'. After source
//		subdirectories and files are copied to this
//		target directory tree, the entire source
//		subdirectory tree identified by the 'dMgr' parent
//		directory along with ALL the files in that
//		subdirectory tree will be deleted.
//
//		Be advised that the parent directory and
//		constituent files in that parent directory will
//		NOT be deleted.
//
//		ALL the files in the subdirectory tree specified
//		by the 'dMgr' parent directory will be moved to
//		the target directory tree specified by input
//		parameter 'targetDMgr'.
//
//		If the directory specified by 'dMgr' does not
//		exist, an error will be returned.
//
//	targetDMgr					*DirMgr
//
//		A pointer to an instance of DirMgr. Source files
//		selected in the source subdirectory tree identified
//		by the 'dMgr' parent directory will be copied to a
//		corresponding directory in the target subdirectory
//		tree specified by this input parameter
//		('targetDMgr').
//
//		If a target directory does not exist, this
//		method will attempt to create it.
//
//	dMgrLabel					string
//
//		The name or label associated with input parameter
//		'dMgr', which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "dMgr" will be
//		automatically applied.
//
//	targetDMgrLabel				string
//
//		The name or label associated with input parameter
//		'targetDMgr' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "targetDMgr" will
//		be automatically applied.
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
//		structure will contain information and statistics
//		describing the outcome of the
//		'move directory tree' operation.
//
//		type DirectoryMoveStats struct {
//			TotalSrcFilesProcessed   uint64
//			SourceFilesMoved         uint64
//			SourceFileBytesMoved     uint64
//			SourceFilesRemaining     uint64
//			SourceFileBytesRemaining uint64
//			TotalDirsProcessed       uint64
//			DirsCreated              uint64
//			SourceOriginalSubDirs      uint64
//			SourceDirWasDeleted      bool
//			Errors             error
//		}
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
func (dMgrHlpr *dirMgrHelper) moveSubDirectoryTree(
	dMgr *DirMgr,
	targetDMgr *DirMgr,
	dMgrLabel string,
	targetDMgrLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	dirMoveStats DirectoryMoveStats,
	nonfatalErrs []error,
	fatalErr error) {

	dMgrHlpr.lock.Lock()

	defer dMgrHlpr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "dirMgrHelper.moveSubDirectoryTree()"

	ePrefix,
		fatalErr = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if fatalErr != nil {

		return dirMoveStats, nonfatalErrs, fatalErr
	}

	if len(dMgrLabel) == 0 {
		dMgrLabel = "dMgr"
	}

	if len(targetDMgrLabel) == 0 {

		targetDMgrLabel = "targetDMgr"
	}

	fileSelectCriteria := FileSelectionCriteria{}

	var dTreeCopyStats DirTreeCopyStats
	var errs2 []error
	var err2 error

	dTreeCopyStats,
		_,
		errs2,
		err2 =
		new(dirMgrHelperNanobot).
			copyDirectoryTree(
				dMgr,
				targetDMgr,
				false, // returnCopiedFilesList
				true,  // skipTopLevelDirectory
				true,  // copyEmptyDirectories
				true,  // copyRegularFiles
				true,  // copySymLinkFiles,
				true,  // copyOtherNonRegularFiles
				fileSelectCriteria,
				"dMgr",
				"targetDMgr",
				ePrefix)

	if len(errs2) > 0 {

		nonfatalErrs = append(nonfatalErrs, errs2...)

	}

	if err2 != nil {

		fatalErr = fmt.Errorf("%v\n"+
			"Errors occurred while copying the source directory tree\n"+
			"to the target directory tree.\n"+
			"The source directory WAS NOT DELETED!\n"+
			"%v Source Directory='%v'\n"+
			"%v Target Directory='%v'\n"+
			"Fatal Error = \n%v\n",
			funcName,
			dMgrLabel,
			dMgr.absolutePath,
			targetDMgrLabel,
			targetDMgr.absolutePath,
			err2.Error())

		return dirMoveStats, nonfatalErrs, fatalErr
	}

	dirMoveStats.TotalDirsProcessed =
		dTreeCopyStats.TotalDirsScanned

	dirMoveStats.TargetDirsCreated =
		dTreeCopyStats.DirsCreated

	dirMoveStats.SourceOriginalSubDirs =
		dTreeCopyStats.TotalDirsScanned

	dirMoveStats.SourceFilesRemaining =
		dTreeCopyStats.FilesNotCopied

	dirMoveStats.SourceFileBytesRemaining =
		dTreeCopyStats.FileBytesNotCopied

	if dirMoveStats.SourceFilesRemaining > 0 {

		fatalErr = fmt.Errorf("%v\n"+
			"Error: Some of the files designated to be moved\n"+
			"to the target directory, were NOT copied!\n"+
			"Therefore the source directory WILL NOT BE DELETED!\n"+
			"Number of Files NOT Copied='%v'\n"+
			"%v Source Directory='%v'\n"+
			"%v Target Directory= '%v'\n",
			ePrefix.String(),
			dTreeCopyStats.FilesNotCopied,
			dMgrLabel,
			dMgr.absolutePath,
			targetDMgrLabel,
			targetDMgr.absolutePath)

		return dirMoveStats, nonfatalErrs, fatalErr
	}

	dirMoveStats.TotalSrcFilesProcessed =
		dTreeCopyStats.TotalFilesProcessed

	err2 = new(dirMgrHelperNanobot).
		deleteAllSubDirectories(
			dMgr,
			false, // returnDeletedSubDirs
			nil,   // deletedSubDirs
			"dMgr",
			ePrefix)

	if err2 != nil {

		fatalErr = fmt.Errorf("%v\n"+
			"Fatal Errror returned by dirMgrHelperMolecule.deleteAllSubDirectories()\n"+
			"Error= \n%v\n",
			funcName,
			err2.Error())

		return dirMoveStats, nonfatalErrs, fatalErr
	}

	dirMoveStats.SourceDirWasDeleted = true

	dirMoveStats.SourceFilesMoved =
		dTreeCopyStats.FilesCopied

	dirMoveStats.SourceFileBytesMoved =
		dTreeCopyStats.FileBytesCopied

	return dirMoveStats, nonfatalErrs, fatalErr
}

// setDirMgrFromKnownPathDirName
//
// Configures the internal field values for the 'dMgr'
// instance using a parent path name and a directory
// name. The parent path and directory name are combined
// to form the full path for the 'dMgr' instance.
//
// This method will replace all previous field values
// with new values based on input parameters
// 'parentPathName' and 'directoryName'.
//
// This method differs from other "Set" methods in that
// it assumes the input parameters are known values and
// do not require the usual analysis and validation
// screening applied by similar methods.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dMgr						*DirMgr
//
//		A pointer to an instance of DirMgr. This instance
//		will be reconfigured using the path supplied by
//		input parameter 'pathStr' and the directory name
//		specified by input parameter 'dirName'.
//
//	parentPath					string
//
//		This string contains the parent path which will
//		be combined with input parameter 'dirName' to
//		create the final path used to configure the DirMgr
//		instance supplied by input parameter 'dMgr'.
//
//	dirName						string
//
//		This string contains the directory which will be
//		combined with the parent directory supplied by
//		input parameter 'parentPath' to create the final
//		directory path which will be used to configure
//		the instance of DirMgr identified by input
//		parameter 'dMgr'.
//
//	dMgrLabel					string
//
//		The name or label associated with input parameter
//		'dMgr', which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "dMgr" will be
//		automatically applied.
//
//	parentPathLabel				string
//
//		The name or label associated with input parameter
//		'parentPath', which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "parentPath" will be
//		automatically applied.
//
//	dirNameLabel				string
//
//		The name or label associated with input parameter
//		'dirName', which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "dirName" will be
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
//	isEmpty						bool
//
//		If the directory path is empty or blank, this
//		returned boolean value is set to 'true'.
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
func (dMgrHlpr *dirMgrHelper) setDirMgrFromKnownPathDirName(
	dMgr *DirMgr,
	parentPath string,
	dirName string,
	dMgrLabel string,
	parentPathLabel string,
	dirNameLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	isEmpty bool,
	err error) {

	if dMgrHlpr.lock == nil {
		dMgrHlpr.lock = new(sync.Mutex)
	}

	dMgrHlpr.lock.Lock()

	defer dMgrHlpr.lock.Unlock()

	isEmpty = true

	var ePrefix *ePref.ErrPrefixDto

	funcName := "dirMgrHelper." +
		"setDirMgrFromKnownPathDirName()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return false, err
	}

	if len(dMgrLabel) == 0 {
		dMgrLabel = "dMgr"
	}

	if dMgr == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter %v pointer is 'nil' !\n",
			ePrefix.String(),
			dMgrLabel)

		return isEmpty, err
	}

	if len(parentPathLabel) == 0 {
		parentPathLabel = "parentPath"
	}

	if len(dirNameLabel) == 0 {
		dirNameLabel = "dirName"
	}

	strLen := 0

	dMgrHlprAtom := dirMgrHelperAtom{}

	parentPath,
		strLen,
		err = dMgrHlprAtom.
		lowLevelScreenPathStrForInvalidChars(
			parentPath,
			parentPathLabel,
			ePrefix)

	if err != nil {

		isEmpty = true
		return isEmpty, err
	}

	lDirName := 0

	dirName,
		lDirName,
		err = dMgrHlprAtom.
		lowLevelScreenPathStrForInvalidChars(
			dirName,
			dirNameLabel,
			ePrefix)

	if err != nil {

		isEmpty = true
		return isEmpty, err
	}

	fh := FileHelper{}

	dotSeparator := "." + string(os.PathSeparator)
	doubleDotSeparator := ".." + string(os.PathSeparator)

	if lDirName > 2 &&
		strings.HasPrefix(dirName, doubleDotSeparator) {

		dirName = dirName[3:]

	} else if lDirName > 1 &&
		(strings.HasPrefix(dirName, dotSeparator) ||
			strings.HasPrefix(dirName, "..")) {

		dirName = dirName[2:]

	} else if lDirName > 0 &&
		(dirName[0] == os.PathSeparator ||
			dirName[0] == '.') {

		dirName = dirName[1:]

	}

	finalPathStr := ""

	if parentPath[strLen-1] != os.PathSeparator {
		finalPathStr = parentPath + string(os.PathSeparator) + dirName
	} else {
		finalPathStr = parentPath + dirName
	}

	var err2 error

	validPathDto := new(ValidPathStrDto).New()

	validPathDto.originalPathStr = finalPathStr
	validPathDto.pathStr = finalPathStr

	validPathDto.absPathStr,
		err2 = new(fileHelperProton).
		makeAbsolutePath(
			validPathDto.pathStr,
			ePrefix.XCpy("validPathDto.absPathStr<-"))

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"Error returned by fh.MakeAbsolutePath(parentPath).\n"+
			"Directory Path='%v'\n"+
			"Error='%v'\n",
			funcName,
			validPathDto.pathStr,
			err2.Error())

		isEmpty = true
		return isEmpty, err
	}

	validPathDto.pathVolumeIndex,
		validPathDto.pathVolumeStrLength,
		validPathDto.pathVolumeName =
		fh.GetVolumeNameIndex(validPathDto.absPathStr)

	validPathDto.pathStrLength = len(parentPath)
	validPathDto.absPathStrLength = len(validPathDto.absPathStr)
	validPathDto.pathDoesExist = PathExistsStatus.Unknown()
	validPathDto.absPathDoesExist = PathExistsStatus.Unknown()
	validPathDto.isInitialized = true
	validPathDto.pathIsValid = PathValidStatus.Valid()

	err = validPathDto.IsDtoValid(ePrefix)

	if err != nil {

		isEmpty = true
		return isEmpty, err
	}

	err = new(dirMgrHelperElectron).
		empty(
			dMgr,
			dMgrLabel,
			ePrefix)

	if err != nil {

		return isEmpty, err
	}

	isEmpty,
		err = new(dirMgrHelperMolecule).
		lowLevelDirMgrFieldConfig(
			dMgr,
			validPathDto,
			dMgrLabel,
			ePrefix)

	return isEmpty, err
}

// setDirMgrWithPathDirectoryName
//
// Configures a Directory Manager instance based on
// 'path' and 'directory name' parameters.
//
// 'path' is treated as the parent directory. The
// directory name ('directoryName') will be added to the
// parent directory to construct the new directory path.
//
// The newly constructed directory path will be used to
// reconfigure the instance of DirMgr passed as input
// parameter 'dMgr'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dMgr						*DirMgr
//
//		A pointer to an instance of DirMgr. This instance
//		will be reconfigured using a new directory path
//		constructed from input parameters 'pathStr' and
//		'directoryName'.
//
//	parentDirectoryPath			string
//
//		The directory specified by input parameter
//		'directoryName' will be added to this parent
//		directory to create the new directory path used
//		to reconfigure 'dMgr'.
//
//	directoryName				string
//
//		A directory name which will be added to the
//		parent directory to construct a new directory
//		path used to reconfigure 'dMgr'.
//
//	dMgrLabel					string
//
//		The name or label associated with input parameter
//		'dMgr', which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "dMgr" will be
//		automatically applied.
//
//	parentDirectoryLabel		string
//
//		The name or label associated with input parameter
//		'parentDirectoryPath' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "parentDirectoryPath"
//		will be automatically applied.
//
//	directoryNameLabel			string
//
//		The name or label associated with input parameter
//		'directoryName' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "directoryName" will
//		be automatically applied.
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
//		If the directory path, constructed from
//		'parentDirectoryPath' and 'directoryName' is
//		empty or blank, this returned boolean value is
//		set to 'true'.
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
func (dMgrHlpr *dirMgrHelper) setDirMgrWithPathDirectoryName(
	dMgr *DirMgr,
	parentDirectoryPath string,
	directoryName string,
	dMgrLabel string,
	parentDirectoryLabel string,
	directoryNameLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	isEmpty bool,
	err error) {

	if dMgrHlpr.lock == nil {
		dMgrHlpr.lock = new(sync.Mutex)
	}

	dMgrHlpr.lock.Lock()

	defer dMgrHlpr.lock.Unlock()

	funcName := "dirMgrHelper.setDirMgrWithPathDirectoryName()"

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		return isEmpty, err
	}

	if len(dMgrLabel) == 0 {
		dMgrLabel = "dMgr"
	}

	if dMgr == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter %v pointer is 'nil' !\n",
			ePrefix.String(),
			dMgrLabel)

		return isEmpty, err
	}

	if len(parentDirectoryLabel) == 0 {
		parentDirectoryLabel = "parentDirectoryPath"
	}

	if len(directoryNameLabel) == 0 {
		directoryNameLabel = "directoryName"
	}

	var strLen int

	dMgrHlprElectron := dirMgrHelperElectron{}

	parentDirectoryPath,
		strLen,
		err = dMgrHlprElectron.
		isPathStringEmptyOrBlank(
			parentDirectoryPath,
			true,
			parentDirectoryLabel,
			ePrefix)

	if err != nil {
		isEmpty = true
		return isEmpty, err
	}

	directoryName,
		_,
		err = dMgrHlprElectron.
		isPathStringEmptyOrBlank(
			directoryName,
			true,
			directoryNameLabel,
			ePrefix)

	if err != nil {
		isEmpty = true
		return isEmpty, err
	}

	if directoryName[0] == os.PathSeparator {
		directoryName = directoryName[1:]
	}

	finalPathStr := ""

	if parentDirectoryPath[strLen-1] != os.PathSeparator {
		finalPathStr =
			parentDirectoryPath + string(os.PathSeparator) + directoryName

	} else {
		finalPathStr = parentDirectoryPath + directoryName
	}

	validPathDto := new(ValidPathStrDto).New()

	validPathDto,
		err =
		new(dirMgrHelperMolecule).
			getValidPathStr(
				finalPathStr,
				"parentDirectoryPath",
				ePrefix.XCpy(
					"finalPathStr"))

	if err != nil {
		isEmpty = true
		return isEmpty, err
	}

	err = new(dirMgrHelperElectron).
		empty(
			dMgr,
			dMgrLabel,
			ePrefix)

	if err != nil {

		return isEmpty, err
	}

	isEmpty,
		err = new(dirMgrHelperMolecule).
		lowLevelDirMgrFieldConfig(
			dMgr,
			validPathDto,
			dMgrLabel,
			ePrefix)

	return isEmpty, err
}

// setPermissions
//
// Sets the read/write and execute permissions for the
// directory identified by input parameter 'dMgr'.
//
// Remember that the treatment of 'execute' permissions
// may vary according to the host operating system.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dMgr						*DirMgr
//
//		A pointer to an instance of DirMgr. This method
//		will set the read/write and execute permissions
//		for the directory identified by this DirMgr
//		instance.
//
//	permissionConfig			FilePermissionConfig
//
//		Use FilePermissionConfig 'New' methods to create
//		directory permissions for the directory
//		identified by input parameter 'dMgr'.
//
//		Type FilePermissionConfig provides methods to
//		support the creation and management of File
//		Permissions for use in controlling file access
//		operations. The Go Programming Language uses
//		os.FileMode and unix permission bits to configure
//		file permissions.
//
//		Reference:
//		https://golang.org/pkg/os/#FileMode
//		https://www.cyberciti.biz/faq/explain-the-nine-permissions-bits-on-files/
//		https://en.wikipedia.org/wiki/File_system_permissions
//
//		The FilePermissionConfig methods will allow for
//		configuration of valid file permissions which are
//		subsequently stored as an os.FileMode in a
//		private member variable,
//		'FilePermissionConfig.fileMode'.
//
//		When evaluated as a string, file permission is
//		defined by a 10-character string. The first
//		character is an 'Entry Type' and the remaining
//		9-characters are unix permission bits.
//
//		Symbolic and Numeric Notation
//
//		Permission codes may be designated with Symbolic
//		Notation or Numeric Octal Notation.
//
//				   Numeric
//		Symbolic	Octal
//		Notation   Notation
//		----------	0000	no permissions
//		-rwx------	0700	read, write, & execute only for owner
//		-rwxrwx---	0770	read, write, & execute for owner and group
//		-rwxrwxrwx	0777	read, write, & execute for owner, group and others
//		---x--x--x	0111	execute
//		--w--w--w-	0222	write
//		--wx-wx-wx	0333	write & execute
//		-r--r--r--	0444	read
//		-r-xr-xr-x	0555	read & execute
//		-rw-rw-rw-	0666	read & write
//		-rwxr-----	0740	owner can read, write, & execute; group can only read;
//		                   others have no permissions
//
//		The permissions used by this method are designed
//		to be used for directories, not files.
//
//		Example:
//
//		drwxrwxrwx - Identifies permissions for directory
//						value = 020000000777
//
//	dMgrLabel					string
//
//		The name or label associated with input parameter
//		'dMgr', which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "dMgr" will be
//		automatically applied.
//
//	permissionConfigLabel		string
//
//		The name or label associated with input parameter
//		'permissionConfig', which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "permissionConfig"
//		will be automatically applied.
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
func (dMgrHlpr *dirMgrHelper) setPermissions(
	dMgr *DirMgr,
	permissionConfig FilePermissionConfig,
	dMgrLabel string,
	permissionConfigLabel string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if dMgrHlpr.lock == nil {
		dMgrHlpr.lock = new(sync.Mutex)
	}

	dMgrHlpr.lock.Lock()

	defer dMgrHlpr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	funcName := "dirMgrHelper.setPermissions()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return err
	}

	if len(dMgrLabel) == 0 {
		dMgrLabel = "dMgr"
	}

	if dMgr == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter %v pointer is 'nil' !\n",
			ePrefix.String(),
			dMgrLabel)

		return err
	}

	if len(permissionConfigLabel) == 0 {
		dMgrLabel = "permissionConfig"
	}

	err = permissionConfig.IsValidInstanceError(ePrefix.XCpy(
		permissionConfigLabel))

	if err != nil {
		return fmt.Errorf("%v\n"+
			"Input parameter '%v' is INVALID!\n"+
			"Error returned by %v.IsValidInstanceError()\n"+
			"%v='%v'\n"+
			"Error='%v'\n\n",
			ePrefix.String(),
			permissionConfigLabel,
			permissionConfigLabel,
			permissionConfigLabel,
			permissionConfig.GetPermissionNarrativeText(),
			err.Error())
	}

	dirPathDoesExist,
		_,
		err := new(dirMgrHelperAtom).
		doesDirectoryExist(
			dMgr,
			PreProcPathCode.None(),
			dMgrLabel,
			ePrefix)

	if err != nil {
		return err
	}

	if !dirPathDoesExist {
		err = fmt.Errorf("%v\n"+
			"ERROR: %v Directory Path DOES NOT EXIST!\n"+
			"%v='%v'\n",
			ePrefix.String(),
			dMgrLabel,
			dMgrLabel,
			dMgr.absolutePath)

		return err
	}

	err = new(FileHelper).ChmodFilePermConfig(
		dMgr.absolutePath,
		permissionConfig,
		ePrefix.XCpy(
			fmt.Sprintf("%v.absolutePath",
				dMgrLabel)))

	if err != nil {
		return fmt.Errorf("%v\n"+
			"Error retrned by FileHelper{}.ChmodFilePermConfig("+
			"%v.absolutePath, %v)\n"+
			"%v.absolutePath=%v\n"+
			"%v='%v'"+
			"Error= \n%v\n",
			funcName,
			dMgrLabel,
			permissionConfigLabel,
			dMgrLabel,
			dMgr.absolutePath,
			permissionConfigLabel,
			permissionConfig.GetPermissionNarrativeText(),
			err.Error())
	}

	return err
}

// substituteBaseDir
//
// Substitute 'baseDir' segment of the current DirMgr
// with a new parent directory identified by input
// parameter 'substituteBaseDir'. This is useful in
// copying files to new directory trees.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dMgr						*DirMgr
//
//		A pointer to an instance of DirMgr. This method
//		will substitute 'baseDir' segment of 'dMgr' with
//		a new parent directory identified by input
//		parameter 'substituteBaseDir'. This operation is
//		useful when copying files to new directory trees.
//
//	baseDir						*DirMgr
//
//		A pointer to an instance of DirMgr. This method
//		will substitute the 'baseDir' segment of 'dMgr'
//		with a new parent directory identified by input
//		parameter 'substituteBaseDir'. This operation is
//		useful when copying files to new directory trees.
//
//	substituteBaseDir			*DirMgr
//
//		A pointer to an instance of DirMgr. This method
//		will substitute the 'baseDir' segment of 'dMgr'
//		with a new parent directory identified by input
//		parameter 'substituteBaseDir'. This operation is
//		useful when copying files to new directory trees.
//
//	dMgrLabel					string
//
//		The name or label associated with input parameter
//		'dMgr', which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "dMgr" will be
//		automatically applied.
//
//	baseDirLabel				string
//
//		The name or label associated with input parameter
//		'baseDir', which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "baseDir" will be
//		automatically applied.
//
//	substituteBaseDirLabel		string
//
//		The name or label associated with input parameter
//		'substituteBaseDir', which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "substituteBaseDir"
//		will be automatically applied.
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
//	newDMgr						DirMgr
//
//		If this method completes successfully, an
//		instance of DirMgr populated with a directory
//		path constructed from input parameters 'dMgr'
//		and 'substituteBaseDir'.
//
//		This method will substitute the 'baseDir' segment
//		of 'dMgr' with a new parent directory identified
//		by input parameter 'substituteBaseDir'. The newly
//		configured directory path will be returned by
//		this parameter, 'newDMgr'.
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
func (dMgrHlpr *dirMgrHelper) substituteBaseDir(
	dMgr *DirMgr,
	baseDir *DirMgr,
	substituteBaseDir *DirMgr,
	dMgrLabel string,
	baseDirLabel string,
	substituteBaseDirLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	newDMgr DirMgr,
	err error) {

	if dMgrHlpr.lock == nil {
		dMgrHlpr.lock = new(sync.Mutex)
	}

	dMgrHlpr.lock.Lock()

	defer dMgrHlpr.lock.Unlock()

	funcName := "dirMgrHelper.substituteBaseDir()"

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return newDMgr, err
	}

	if len(dMgrLabel) == 0 {
		dMgrLabel = "dMgr"
	}

	if dMgr == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter %v pointer is 'nil' !\n",
			ePrefix.String(),
			dMgrLabel)

		return newDMgr, err
	}

	if len(baseDirLabel) == 0 {
		baseDirLabel = "baseDir"
	}

	if baseDir == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter %v pointer is 'nil' !\n",
			ePrefix.String(),
			baseDirLabel)

		return newDMgr, err
	}

	if len(substituteBaseDirLabel) == 0 {
		substituteBaseDirLabel = "substituteBaseDir"
	}

	if substituteBaseDir == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter %v pointer is 'nil' !\n",
			ePrefix.String(),
			substituteBaseDirLabel)

		return newDMgr, err
	}

	dMgrHlprAtom := dirMgrHelperAtom{}

	dMgrHlprElectron := dirMgrHelperElectron{}

	_,
		_,
		err = dMgrHlprAtom.
		doesDirectoryExist(
			dMgr,
			PreProcPathCode.None(),
			dMgrLabel,
			ePrefix)

	if err != nil {
		return newDMgr, err
	}

	_,
		_,
		err = dMgrHlprAtom.
		doesDirectoryExist(
			baseDir,
			PreProcPathCode.None(),
			baseDirLabel,
			ePrefix)

	if err != nil {
		return newDMgr, err
	}

	_,
		_,
		err = dMgrHlprAtom.
		doesDirectoryExist(
			substituteBaseDir,
			PreProcPathCode.None(),
			substituteBaseDirLabel,
			ePrefix)

	if err != nil {
		return newDMgr, err
	}

	thisDirAbsPath := strings.ToLower(dMgr.absolutePath)

	oldBaseAbsPath := strings.ToLower(baseDir.absolutePath)

	newBaseAbsPath := strings.ToLower(substituteBaseDir.absolutePath)

	idx := strings.Index(thisDirAbsPath, oldBaseAbsPath)

	if idx < 0 {

		err = fmt.Errorf("%v\n"+
			"The base directory was NOT found in the current %v path!\n"+
			"%v Path='%v'\n"+
			"%v Path='%v'\n",
			ePrefix.String(),
			dMgrLabel,
			dMgrLabel,
			thisDirAbsPath,
			baseDirLabel,
			oldBaseAbsPath)

		return newDMgr, err
	}

	if idx != 0 {
		err = fmt.Errorf("%v\n"+
			"The %v directory was NOT found at the beginning of the %v path!\n"+
			"%v Path='%v'\n"+
			"%v Path='%v'\n",
			ePrefix.String(),
			baseDirLabel,
			dMgrLabel,
			dMgrLabel,
			thisDirAbsPath,
			baseDirLabel,
			oldBaseAbsPath)

		return newDMgr, err
	}

	oldBaseLen := len(oldBaseAbsPath)

	newAbsPath := newBaseAbsPath + thisDirAbsPath[oldBaseLen:]

	isEmpty := false

	isEmpty,
		err = new(dirMgrHelperNanobot).
		setDirMgr(
			&newDMgr,
			newAbsPath,
			dMgrLabel,
			"newAbsPath",
			ePrefix)

	if err != nil {

		_ = dMgrHlprElectron.
			empty(
				&newDMgr,
				dMgrLabel,
				ePrefix.XCpy(
					"newDMgr"))

		return newDMgr, err
	}

	if isEmpty {

		_ = dMgrHlprElectron.
			empty(
				&newDMgr,
				dMgrLabel,
				ePrefix)

		err = fmt.Errorf("%v\n"+
			"ERROR: New generated Directory Path Is Invalid!\n"+
			"isEmpty='true'\n"+
			"newAbsPath='%v'\n",
			ePrefix,
			newAbsPath)
	}

	return newDMgr, err
}
