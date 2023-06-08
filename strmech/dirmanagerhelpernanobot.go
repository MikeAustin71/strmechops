package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"os"
	"strings"
	"sync"
)

type dirMgrHelperNanobot struct {
	lock *sync.Mutex
}

// copyDirectoryTree
//
// Helper method for 'DirMgr'. This method is designed to
// copy entire directory trees.
//
// If the target directory does not exist, this method
// will attempt to create it.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	sourceDMgr					*DirMgr
//
//		A pointer to an instance of DirMgr which
//		identifies the source directory tree from which
//		files will be copied to the directory tree
//		identified by input parameter 'targetDMgr'.
//
//	targetDMgr					*DirMgr
//
//		A pointer to an instance of DirMgr which
//		identifies the destination directory tree to
//		which files will be copied from the directory
//		tree identified by input parameter 'sourceDMgr'.
//
//		If this target directory does not exist, this
//		method will attempt to create it.
//
//	copyEmptyDirectories		bool
//
//		If this parameter is set to 'true', directories
//		containing zero files will be created and no
//		errors will be returned.
//
//	skipTopLevelDirectory		bool
//
//		If this parameter is set to 'true', the top level
//		source directory will be skipped, and it will not
//		be copied to the directory tree identified by
//		input parameter 'targetDMgr'.
//
//	fileSelectCriteria			FileSelectionCriteria
//
//		This input parameter should be configured with the
//		desired file selection criteria. Files matching
//		this criteria will be copied  to the directory
//		identified by input parameter, 'targetDir'.
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
//	  The FileSelectionCriteria type allows for
//	  configuration of single or multiple file selection
//	  criterion. The 'SelectCriterionMode' can be used to
//	  specify whether the file must match all, or any one,
//	  of the active file selection criterion.
//
//		Elements of the FileSelectionCriteria are described
//		below:
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
//		If all of the file selection criterion in the FileSelectionCriteria object are
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
//	sourceDMgrLabel				string
//
//		The name or label associated with input parameter
//		'sourceDMgr' which will be used in error messages
//		returned by this method.
//
//	targetDMgrLabel				string
//
//		The name or label associated with input parameter
//		'targetDMgr' which will be used in error messages
//		returned by this method.
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
//	dTreeCopyStats				DirTreeCopyStats
//
//		If this method completes successfully, an
//		instance of DirTreeCopyStats will be returned
//		populated with information and statistics related
//		to the directory tree copy operation.
//
//			type DirTreeCopyStats struct {
//				TotalDirsScanned    uint64
//				DirsCopied          uint64
//				DirsCreated         uint64
//				TotalFilesProcessed uint64
//				FilesCopied         uint64
//				FileBytesCopied     uint64
//				FilesNotCopied      uint64
//				FileBytesNotCopied  uint64
//				ComputeError        error
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
func (dMgrHlprNanobot *dirMgrHelperNanobot) copyDirectoryTree(
	sourceDMgr *DirMgr,
	targetDMgr *DirMgr,
	copyEmptyDirectories bool,
	skipTopLevelDirectory bool,
	fileSelectCriteria FileSelectionCriteria,
	sourceDMgrLabel string,
	targetDMgrLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	dTreeCopyStats DirTreeCopyStats,
	errs []error) {

	if dMgrHlprNanobot.lock == nil {
		dMgrHlprNanobot.lock = new(sync.Mutex)
	}

	dMgrHlprNanobot.lock.Lock()

	defer dMgrHlprNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	funcName := "dirMgrHelperNanobot.copyDirectoryTree()"

	errs = make([]error, 0)

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		errs = append(errs, err)

		return dTreeCopyStats, errs
	}

	if len(sourceDMgrLabel) == 0 {

		sourceDMgrLabel = "sourceDMgr"
	}

	var dMgrHlprPreon = new(dirMgrHelperPreon)

	_,
		_,
		err = dMgrHlprPreon.
		validateDirMgr(
			sourceDMgr,
			true, // Path MUST exist on disk
			sourceDMgrLabel,
			ePrefix.XCpy(
				sourceDMgrLabel))

	if err != nil {

		errs = append(errs, err)

		return dTreeCopyStats, errs
	}

	if len(targetDMgrLabel) == 0 {

		targetDMgrLabel = "targetDMgr"
	}

	_,
		_,
		err = dMgrHlprPreon.
		validateDirMgr(
			targetDMgr,
			false, // Path is NOT required to exist on disk
			targetDMgrLabel,
			ePrefix.XCpy(
				targetDMgrLabel))

	if err != nil {

		errs = append(errs, err)

		return dTreeCopyStats, errs
	}

	dirs := DirMgrCollection{}

	err = dirs.AddDirMgr(
		*sourceDMgr,
		ePrefix.XCpy("sourceDMgr"))

	if err != nil {

		errs = append(errs, err)

		return dTreeCopyStats, errs
	}

	var err2 error

	baseDirLen := len(sourceDMgr.absolutePath)

	var nextTargetDMgr DirMgr

	osPathSepStr := string(os.PathSeparator)

	var nextDir DirMgr
	var srcFile, targetFile string
	fh := FileHelper{}
	dMgrHlprMolecule := dirMgrHelperMolecule{}
	dMgrHlprAtom := dirMgrHelperAtom{}
	dMgrHlprTachyon := new(dirMgrHelperTachyon)
	var fileInfos []FileInfoPlus
	var errs2 []error
	dirCreated := false
	mainLoopIsDone := false
	isMatch := false
	isTopLevelDir := true
	isNewDir := false
	isFirstLoop := true
	dMgrPathDoesExist := false
	var lenFileInfos int

	if !skipTopLevelDirectory {
		dTreeCopyStats.TotalDirsScanned++
	}

	for !mainLoopIsDone {

		if isFirstLoop {
			isTopLevelDir = true
			isFirstLoop = false
		} else {
			isTopLevelDir = false
		}

		nextDir, err = dirs.PopFirstDirMgr(
			ePrefix.XCpy("dirs"))

		if err != nil && err == io.EOF {

			mainLoopIsDone = true

			isTopLevelDir = false

			break

		} else if err != nil {

			err2 = fmt.Errorf("%v\n"+
				"Error returned by dirs.PopFirstDirMgr().\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				err.Error())

			errs = append(errs, err2)

			return dTreeCopyStats, errs
		}

		nextTargetDMgr, err = new(DirMgr).New(
			targetDMgr.absolutePath+
				nextDir.absolutePath[baseDirLen:],
			ePrefix)

		if err != nil {

			err2 = fmt.Errorf("%v\n"+
				"Error return by DirMgr{}.New(%v.absolutePath + "+
				"nextDir.absolutePath[baseDirLen:])\n"+
				"%v.absolutePath='%v'\n"+
				"nextDir.absolutePath[baseDirLen:]='%v'\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				targetDMgrLabel,
				targetDMgrLabel,
				targetDMgr.absolutePath,
				nextDir.absolutePath[baseDirLen:],
				err.Error())

			errs = append(errs, err2)

			continue
		}

		dirCreated = false

		if isTopLevelDir &&
			!skipTopLevelDirectory &&
			copyEmptyDirectories {

			dirCreated,
				err = dMgrHlprMolecule.
				lowLevelMakeDir(
					&nextTargetDMgr,
					"1-nextTargetDMgr",
					ePrefix.XCpy(
						"1-nextTargetDMgr"))

		} else if !isTopLevelDir && copyEmptyDirectories {

			dirCreated,
				err = dMgrHlprMolecule.
				lowLevelMakeDir(
					&nextTargetDMgr,
					"2-nextTargetDMgr",
					ePrefix.XCpy(
						"2-nextTargetDMgr"))

		} else {
			err = nil
		}

		if err != nil {

			err2 = fmt.Errorf("%v\n"+
				"Error creating target next directory!\n"+
				"Target Next Directory='%v'\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				nextTargetDMgr.absolutePath,
				err.Error())

			errs = append(errs, err2)

			isTopLevelDir = false

			continue

		} else if dirCreated {
			dTreeCopyStats.DirsCreated++
		}

		if !skipTopLevelDirectory && copyEmptyDirectories {

			dTreeCopyStats.DirsCopied++

		} else if skipTopLevelDirectory &&
			copyEmptyDirectories &&
			!isTopLevelDir {

			dTreeCopyStats.DirsCopied++
		}

		fileInfos,
			lenFileInfos,
			errs2 = dMgrHlprTachyon.
			getFileInfosFromDirectory(
				&nextDir,
				"nextDir",
				ePrefix.XCpy("nextDir"))

		if len(errs2) != 0 {

			errs = append(errs, errs2...)

			continue
		}

		if lenFileInfos == 0 {

			err = fmt.Errorf("%v\n"+
				"Error: dirMgrHelperElectron.getFileInfosFromDirectory()\n"+
				"returned a zero length array of File Info Objects from:\n"+
				"nextDir = %v\n",
				ePrefix.String(),
				nextDir.absolutePath)

			errs = append(errs, err)

			continue
		}

		isNewDir = true

		for _, fileInfo := range fileInfos {

			if fileInfo.IsDir() {
				// This is a directory

				err = dirs.AddDirMgrByPathNameStr(
					nextDir.absolutePath+
						osPathSepStr+
						fileInfo.Name(),
					ePrefix)

				if err != nil {
					err2 = fmt.Errorf("%v\n"+
						"Error returned by dirs.AddDirMgrByPathNameStr(newDir).\n"+
						"newDir='%v'\n"+
						"Error= \n%v\n",
						ePrefix.String(),
						nextDir.absolutePath+osPathSepStr+fileInfo.Name(),
						err.Error())

					errs = append(errs, err2)
					continue
				}

				// Count Directories Processed
				dTreeCopyStats.TotalDirsScanned++

				continue
			} // End of IsDir()

			// This is a file
			if isTopLevelDir && skipTopLevelDirectory {
				// Skip all files in the
				// parent directory.
				continue
			}

			// This is a file eligible for
			// matching with selection criteria
			dTreeCopyStats.TotalFilesProcessed++

			// Determine if it matches the find file criteria.
			isMatch,
				err,
				_ =
				fh.FilterFileName(
					fileInfo,
					fileSelectCriteria,
					ePrefix)

			if err != nil {

				err2 = fmt.Errorf("%v\n"+
					"Error returned by fh.FilterFileName(nameDirEntry, fileSelectCriteria).\n"+
					"Directory='%v'\n"+
					"File Name='%v'\n"+
					"Error= \n%v\n",
					ePrefix.String(),
					nextDir.absolutePath,
					fileInfo.Name(),
					err.Error())

				dTreeCopyStats.TotalFilesProcessed--

				errs = append(errs, err2)

				continue
			}

			if isMatch {
				// This file is a Match!
				dirCreated = false

				dMgrPathDoesExist,
					_,
					err =
					dMgrHlprAtom.doesDirectoryExist(
						&nextTargetDMgr,
						PreProcPathCode.None(),
						"nextTargetDMgr",
						ePrefix)

				if err != nil {
					errs = append(errs, err)
					break
				}

				// Create Directory if needed
				if !dMgrPathDoesExist {

					dirCreated,
						err = dMgrHlprMolecule.lowLevelMakeDir(
						&nextTargetDMgr,
						"3-nextTargetDMgr",
						ePrefix)

					if err != nil {
						err2 = fmt.Errorf("%v\n"+
							"Error creating targetFile directory!\n"+
							"Target Directory='%v'\n"+
							"Error= \n%v\n",
							funcName,
							nextTargetDMgr.absolutePath,
							err.Error())

						errs = append(errs, err2)
						break

					} else if dirCreated {
						dTreeCopyStats.DirsCreated++
						dTreeCopyStats.DirsCopied++
					}

				} else if isNewDir && !copyEmptyDirectories {

					dTreeCopyStats.DirsCopied++
				}

				isNewDir = false

				srcFile = nextDir.absolutePath +
					osPathSepStr + fileInfo.Name()

				targetFile = nextTargetDMgr.absolutePath +
					osPathSepStr + fileInfo.Name()

				err = dMgrHlprMolecule.
					lowLevelCopyFile(
						srcFile,
						fileInfo,
						targetFile,
						"srcFile",
						"targetFile",
						ePrefix)

				if err != nil {

					errs = append(errs, err)

					dTreeCopyStats.FilesNotCopied++
					dTreeCopyStats.FileBytesNotCopied += uint64(fileInfo.Size())

				} else {

					dTreeCopyStats.FilesCopied++
					dTreeCopyStats.FileBytesCopied += uint64(fileInfo.Size())
				}

			} else {
				// This file is NOT A Match
				// NOT Selected File
				dTreeCopyStats.FilesNotCopied++
				dTreeCopyStats.FileBytesNotCopied += uint64(fileInfo.Size())

			}
		} // End of range nameFileInfos

		isTopLevelDir = false
	} // End of main loop

	// Final verification of
	dMgrPathDoesExist,
		_,
		err = new(dirMgrHelperElectron).
		lowLevelDoesDirectoryExist(
			targetDMgr.absolutePath,
			"targetDMgr",
			ePrefix)

	if err != nil {

		err2 = fmt.Errorf("%v\n"+
			"After Copy Operation 'targetDMgr' path returned non-path error!\n"+
			"Error= \n%v\n",
			funcName,
			err.Error())

		errs = append(errs, err2)
	}

	if dTreeCopyStats.FilesCopied > 0 &&
		!dMgrPathDoesExist {

		err2 = fmt.Errorf("%v\n"+
			"\nERROR: The copy operation failed to create\n"+
			"the 'targetDMgr' path. 'targetDMgr' path DOES NOT EXIST!\n"+
			"targetDMgr Path='%v'\n",
			ePrefix.String(),
			targetDMgr.absolutePath)

		errs = append(errs, err2)
	}

	if dTreeCopyStats.TotalFilesProcessed !=
		dTreeCopyStats.FilesCopied+dTreeCopyStats.FilesNotCopied {

		err2 = fmt.Errorf("%v\n"+
			"File Counting Error: Number of Total Files Processed DOES\n"+
			"NOT EQUAL the Number of Files Copied Plus Number of Files NOT copied!\n"+
			"Total Number of Files Processed='%v'\n"+
			"         Number of Files Copied='%v'\n"+
			"     Number of Files NOT Copied='%v'\n\n",
			ePrefix.String(),
			dTreeCopyStats.TotalFilesProcessed,
			dTreeCopyStats.FilesCopied,
			dTreeCopyStats.FilesNotCopied)

		dTreeCopyStats.ComputeError = fmt.Errorf("%v", err2.Error())

		errs = append(errs, err2)
	}

	return dTreeCopyStats, errs
}

func (dMgrHlprNanobot *dirMgrHelperNanobot) newCopyDirectoryTree(
	sourceDMgr *DirMgr,
	targetDMgr *DirMgr,
	skipTopLevelDirectory bool,
	copyEmptyDirectories bool,
	copySymLinkFiles bool,
	copyOtherNonRegularFiles bool,
	fileSelectCriteria FileSelectionCriteria,
	sourceDMgrLabel string,
	targetDMgrLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	dTreeCopyStats DirTreeCopyStats,
	nonfatalErrs []error,
	fatalErr error) {

	if dMgrHlprNanobot.lock == nil {
		dMgrHlprNanobot.lock = new(sync.Mutex)
	}

	dMgrHlprNanobot.lock.Lock()

	defer dMgrHlprNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "dirMgrHelperNanobot.newCopyDirectoryTree()"

	nonfatalErrs = make([]error, 0)

	ePrefix,
		fatalErr = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if fatalErr != nil {

		return dTreeCopyStats, nonfatalErrs, fatalErr
	}

	if len(sourceDMgrLabel) == 0 {

		sourceDMgrLabel = "sourceDMgr"
	}

	if len(targetDMgrLabel) == 0 {

		targetDMgrLabel = "targetDMgr"
	}

	dMgrHlprPreon := new(dirMgrHelperPreon)

	_,
		_,
		fatalErr = dMgrHlprPreon.
		validateDirMgr(
			sourceDMgr,
			true, // Path MUST exist on disk
			sourceDMgrLabel,
			ePrefix.XCpy(
				sourceDMgrLabel))

	if fatalErr != nil {

		return dTreeCopyStats, nonfatalErrs, fatalErr
	}

	if len(targetDMgrLabel) == 0 {

		targetDMgrLabel = "targetDMgr"
	}

	baseSourceDirLen := len(sourceDMgr.absolutePath)

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

		return dTreeCopyStats, nonfatalErrs, fatalErr
	}

	baseTargetDirLen := len(targetDMgr.absolutePath)

	dMgrHlprMolecule := dirMgrHelperMolecule{}
	var dirCreated bool
	var err2 error

	if !targetPathDoesExist && copyEmptyDirectories {

		dirCreated,
			err2 = dMgrHlprMolecule.
			lowLevelMakeDir(
				targetDMgr,
				"targetDMgr",
				ePrefix.XCpy(
					"targetDMgr"))

		if err2 != nil {

			fatalErr = fmt.Errorf("%v\n"+
				"Error occurred while creating 'targetDMgr' Directory!\n"+
				"targetDMgr= '%v'\n"+
				"Error= \n%v\n",
				funcName,
				targetDMgr.absolutePath,
				err2.Error())

			return dTreeCopyStats, nonfatalErrs, fatalErr
		}

		if !dirCreated {

			fatalErr = fmt.Errorf("%v\n"+
				"Error: Attempted creation of 'targetDMgr' Directory FAILED!\n"+
				"targetDMgr= '%v'\n"+
				ePrefix.String(),
				targetDMgr.absolutePath)

			return dTreeCopyStats, nonfatalErrs, fatalErr

		} else {

			dTreeCopyStats.DirsCreated++
		}
	}

	dMgrHlprPlanck := new(dirMgrHelperPlanck)
	var subDirectories DirMgrCollection
	var subDirCopyStats DirectoryCopyStats
	var errs2 []error

	if !skipTopLevelDirectory {
		dTreeCopyStats.TotalDirsScanned++

		subDirCopyStats,
			subDirectories,
			errs2,
			err2 = dMgrHlprPlanck.
			copyDirectoryFiles(
				sourceDMgr,
				targetDMgr,
				fileSelectCriteria,
				copyEmptyDirectories,
				copySymLinkFiles,
				copyOtherNonRegularFiles,
				sourceDMgrLabel,
				targetDMgrLabel,
				ePrefix)

		if len(errs2) > 0 {

			nonfatalErrs = append(nonfatalErrs, errs2...)

		}

		if err2 != nil {

			fatalErr = fmt.Errorf("%v\n"+
				"Fatal Error occurred while copying the\n"+
				"source directory to the target directory.\n"+
				"Source Directory %v= %v\n,"+
				"Target Directory %v= %v\n"+
				"Error=\n%v\n",
				funcName,
				sourceDMgrLabel,
				sourceDMgr.absolutePath,
				targetDMgrLabel,
				targetDMgr.absolutePath,
				err2.Error())

			return dTreeCopyStats, nonfatalErrs, fatalErr
		}

		dTreeCopyStats.AddDirCopyStats(
			subDirCopyStats)

	}

}

// lowLevelDirMgrFieldConfig
//
// Receives an instance of DirMgr and proceeds to
// configure the internal data elements for that instance
// using data from input parameter 'validPathDto'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dMgr						*DirMgr
//
//		A pointer to an instance of DirMgr. This instance
//		will be reconfigured using the data provided by
//		input parameter 'validPathDto'.
//
//	validPathDto				ValidPathStrDto
//
//		This structure is used to transfer a file/path
//		string plus associated attributes and errors.
//
//		type ValidPathStrDto struct {
//			isInitialized bool
//				Signals whether the current ValidPathStrDto instance
//				has been properly initialized.
//
//			originalPathStr string
//				The original, unformatted path string
//
//			pathStr string
//				The path string which may or may not be
//				the absolute path
//
//			pathFInfoPlus FileInfoPlus
//				Only populated if absValidPath exists on disk.
//
//			pathDoesExist PathExistsStatusCode
//				-1 = don't know, file/path existence has not been tested
//				0 - No, tests show the file/path doesn't exist on disk.
//				1 - Yes, tests show the file/path does exist on disk.
//
//			pathStrLength int
//				Length of the path string
//
//			absPathStr string
//				The absolute path version of 'path'
//
//			absPathFInfoPlus FileInfoPlus
//				Only populated if absValidPath exists on disk.
//
//			absPathDoesExist PathExistsStatusCode
//				-1 - don't know, has not been tested
//				 0 - No, tests shown path doesn't exist
//				 1 - Yes, tests show path does exist
//
//			absPathStrLength int
//				Length of the absolute path string
//
//			pathType PathFileTypeCode
//				The path type. Path File, Path Directory
//
//			pathIsValid PathValidityStatusCode
//				-1 - don't know
//				 0 - No path is NOT valid
//				 1 - Yes, path is valid
//
//			pathVolumeName string
//				Volume name associated with current path
//
//			pathVolumeIndex int
//				Index of the starting character of Volume Name
//				in the path string.
//
//			pathVolumeStrLength int
//				Length of the Volume name in the path string.
//
//			err error
//				If no error is encountered this value is nil
//		}
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
func (dMgrHlprNanobot *dirMgrHelperNanobot) lowLevelDirMgrFieldConfig(
	dMgr *DirMgr,
	validPathDto ValidPathStrDto,
	dMgrLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	isEmpty bool,
	err error) {

	if dMgrHlprNanobot.lock == nil {
		dMgrHlprNanobot.lock = new(sync.Mutex)
	}

	dMgrHlprNanobot.lock.Lock()

	defer dMgrHlprNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"dirMgrHelper."+
			"lowLevelDirMgrFieldConfig()",
		"")

	if err != nil {
		return false, err
	}

	isEmpty = false
	err = nil

	if dMgr == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter %v pointer is 'nil'!\n",
			ePrefix.String(),
			dMgrLabel)

		return isEmpty, err
	}

	err = validPathDto.IsDtoValid(ePrefix)

	if err != nil {
		return isEmpty, err
	}

	if len(dMgrLabel) == 0 {

		dMgrLabel = "dMgr"
	}

	dMgr.originalPath = validPathDto.originalPathStr

	dMgr.path = validPathDto.pathStr
	dMgr.isPathPopulated = true

	dMgr.absolutePath = validPathDto.absPathStr
	dMgr.isAbsolutePathPopulated = true

	fh := FileHelper{}

	fHelpMolecule := fileHelperMolecule{}

	dMgrHlprElectron := dirMgrHelperElectron{}

	var dirPathDoesExist, dirAbsPathDoesExist bool
	var pathFInfoPlus, absPathFInfoPlus FileInfoPlus

	if validPathDto.pathDoesExist != PathExistsStatus.Unknown() &&
		validPathDto.absPathDoesExist != PathExistsStatus.Unknown() {

		if validPathDto.pathDoesExist ==
			PathExistsStatus.DoesNotExist() {
			dirPathDoesExist = false
		} else {
			dirPathDoesExist = true
			pathFInfoPlus = validPathDto.pathFInfoPlus.CopyOut()
		}

		if validPathDto.absPathDoesExist ==
			PathExistsStatus.DoesNotExist() {
			dirAbsPathDoesExist = false
		} else {
			dirAbsPathDoesExist = true
			absPathFInfoPlus = validPathDto.absPathFInfoPlus.CopyOut()
		}

	} else {
		_,
			dirPathDoesExist,
			pathFInfoPlus,
			err =
			fHelpMolecule.doesPathFileExist(
				dMgr.path,
				PreProcPathCode.None(),
				ePrefix,
				dMgrLabel+".path")

		if err != nil {

			_ = dMgrHlprElectron.empty(
				dMgr,
				dMgrLabel,
				ePrefix)

			isEmpty = true
			return isEmpty, err
		}

		_,
			dirAbsPathDoesExist,
			absPathFInfoPlus,
			err =
			fHelpMolecule.doesPathFileExist(
				dMgr.absolutePath,
				PreProcPathCode.None(),
				ePrefix.XCpy(
					"dirAbsPathDoesExist<-dMgr.absolutePath"),
				dMgrLabel+".absolutePath")

		if err != nil {

			_ = dMgrHlprElectron.empty(
				dMgr,
				dMgrLabel,
				ePrefix)

			isEmpty = true
			return isEmpty, err
		}
	}

	if !dirPathDoesExist {
		dMgr.doesPathExist = false

	} else {

		if !pathFInfoPlus.IsDir() {
			_ = dMgrHlprElectron.empty(
				dMgr,
				dMgrLabel,
				ePrefix)

			err = fmt.Errorf("%v\n"+
				"ERROR: Directory path exists, but it is a File - NOT a directory!\n"+
				"%v='%v'\n",
				ePrefix.String(),
				dMgrLabel,
				dMgr.path)

			isEmpty = true
			return isEmpty, err
		}

		if pathFInfoPlus.Mode().IsRegular() {

			_ = dMgrHlprElectron.empty(
				dMgr,
				dMgrLabel,
				ePrefix)

			err = fmt.Errorf("%v\n"+
				"Error: Directory path exists, but\n"+
				"it is classified as as a Regular File!\n"+
				"%v='%v'\n",
				ePrefix.String(),
				dMgrLabel,
				dMgr.path)

			isEmpty = true
			return isEmpty, err
		}

		dMgr.doesPathExist = true
	}

	if dirAbsPathDoesExist {

		if !absPathFInfoPlus.IsDir() {
			_ = dMgrHlprElectron.empty(
				dMgr,
				dMgrLabel,
				ePrefix)

			err = fmt.Errorf("%v\n"+
				"The Directory Manager absolute path exists and IS NOT A DIRECTORY!.\n"+
				"%v Path='%v'\n",
				ePrefix.String(),
				dMgrLabel,
				dMgr.absolutePath)

			isEmpty = true
			return isEmpty, err
		}

		if absPathFInfoPlus.Mode().IsRegular() {

			_ = dMgrHlprElectron.empty(
				dMgr,
				dMgrLabel,
				ePrefix)

			err = fmt.Errorf("%v\n"+
				"Error: Directory absolute path exists, but\n"+
				"it is classified as as a Regular File!\n"+
				"%v='%v'\n",
				ePrefix.String(),
				dMgrLabel,
				dMgr.absolutePath)

			isEmpty = true
			return isEmpty, err
		}

		dMgr.doesAbsolutePathExist = true
		dMgr.actualDirFileInfo = absPathFInfoPlus.CopyOut()

	} else {
		dMgr.doesAbsolutePathExist = false
		dMgr.actualDirFileInfo = FileInfoPlus{}
	}

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

	if dMgr.path != dMgr.absolutePath {
		dMgr.isAbsolutePathDifferentFromPath = true
	}

	if validPathDto.pathVolumeName != "" {
		dMgr.isVolumePopulated = true
		dMgr.volumeName = validPathDto.pathVolumeName
	}

	if dMgr.isAbsolutePathPopulated && dMgr.isPathPopulated {
		dMgr.isInitialized = true
		isEmpty = false
	} else {
		isEmpty = true
	}

	err = nil

	return isEmpty, err
}

// setDirMgr
//
// Sets internal values for a DirMgr instance based on a
// path or path/file name string passed as an input
// parameter.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dMgr						*DirMgr
//
//		A pointer to an instance of DirMgr. This instance
//		will be reconfigured using the path supplied by
//		input parameter 'pathStr'.
//
//	pathStr						string
//
//		This string contains the path or path/filename
//		used to configure the DirMgr instance supplied by
//		input parameter 'dMgr'.
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
func (dMgrHlprNanobot *dirMgrHelperNanobot) setDirMgr(
	dMgr *DirMgr,
	pathStr string,
	dMgrLabel string,
	pathStrLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	isEmpty bool,
	err error) {

	if dMgrHlprNanobot.lock == nil {
		dMgrHlprNanobot.lock = new(sync.Mutex)
	}

	dMgrHlprNanobot.lock.Lock()

	defer dMgrHlprNanobot.lock.Unlock()

	isEmpty = true

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"dirMgrHelperNanobot."+
			"setDirMgr()",
		"")

	if err != nil {
		return isEmpty, err
	}

	if len(dMgrLabel) == 0 {
		dMgrLabel = "dMgr"
	}

	if dMgr == nil {

		err = fmt.Errorf("%v\n"+
			"Input parameter %v pointer is 'nil'!\n",
			ePrefix.String(),
			dMgrLabel)

		return isEmpty, err
	}

	if len(pathStrLabel) == 0 {
		pathStrLabel = "pathStr"
	}

	if len(pathStr) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter %v pointer is an\n"+
			"empty or zero length string\n",
			ePrefix.String(),
			pathStrLabel)

		return isEmpty, err

	}

	validPathDto := new(ValidPathStrDto).New()

	validPathDto,
		err =
		new(dirMgrHelperMolecule).
			getValidPathStr(
				pathStr,
				pathStrLabel,
				ePrefix)

	if err != nil {
		isEmpty = true
		return isEmpty, err
	}

	err = new(dirMgrHelperElectron).empty(
		dMgr,
		dMgrLabel,
		ePrefix)

	if err != nil {

		return isEmpty, err
	}

	return new(dirMgrHelperNanobot).
		lowLevelDirMgrFieldConfig(
			dMgr,
			validPathDto,
			dMgrLabel,
			ePrefix)
}
