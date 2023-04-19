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

	if targetDMgr == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter %v pointer is 'nil'!\n",
			ePrefix.String(),
			targetDMgrLabel)

		errs = append(errs, err)

		return dTreeCopyStats, errs

	}

	dMgrHlprAtom := dirMgrHelperAtom{}

	dMgrPathDoesExist,
		_,
		err :=
		dMgrHlprAtom.doesDirectoryExist(
			sourceDMgr,
			PreProcPathCode.None(),
			sourceDMgrLabel,
			ePrefix.XCpy(
				"sourceDMgr"))

	if err != nil {

		errs = append(errs, err)

		return dTreeCopyStats, errs
	}

	if !dMgrPathDoesExist {

		err = fmt.Errorf("%v\n"+
			"Error: %v directory path DOES NOT EXIST!\n"+
			"%v='%v'\n",
			ePrefix,
			sourceDMgrLabel,
			sourceDMgrLabel,
			sourceDMgr.absolutePath)

		errs = append(errs, err)

		return dTreeCopyStats, errs
	}

	var err2 error
	_,
		_,
		err2 =
		dMgrHlprAtom.doesDirectoryExist(
			targetDMgr,
			PreProcPathCode.None(),
			targetDMgrLabel,
			ePrefix.XCpy(
				"targetDMgr"))

	if err2 != nil {

		errs = append(errs, err2)
		return dTreeCopyStats, errs
	}

	baseDirLen := len(sourceDMgr.absolutePath)

	var nextTargetDMgr DirMgr
	var nameFileInfos []os.FileInfo

	osPathSepStr := string(os.PathSeparator)

	dirs := DirMgrCollection{}

	var nextDir DirMgr
	var dirPtr *os.File
	dirPtr = nil
	var srcFile, targetFile string
	fh := FileHelper{}

	dirs.AddDirMgr(dMgrHlprAtom.copyOut(sourceDMgr))

	if !skipTopLevelDirectory {
		dTreeCopyStats.TotalDirsScanned++
	}

	dirCreated := false
	mainLoopIsDone := false
	file2LoopIsDone := false
	isMatch := false
	isTopLevelDir := true
	isNewDir := false
	isFirstLoop := true

	dMgrHlprMolecule := dirMgrHelperMolecule{}

	for !mainLoopIsDone {

		if isFirstLoop {
			isTopLevelDir = true
			isFirstLoop = false
		} else {
			isTopLevelDir = false
		}

		if dirPtr != nil {

			err = dirPtr.Close()

			if err != nil {

				err2 = fmt.Errorf("%v"+
					"Error returned by dirPtr.Close()\n"+
					"Error= \n%v\n",
					ePrefix.String(),
					err.Error())

				errs = append(errs, err2)
			}

			dirPtr = nil
		}

		nextDir, err = dirs.PopFirstDirMgr()

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
			targetDMgr.absolutePath +
				nextDir.absolutePath[baseDirLen:])

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

		isNewDir = true

		dirPtr, err = os.Open(nextDir.absolutePath)

		if err != nil {

			err2 = fmt.Errorf("%v\n"+
				"Error return by os.Open(nextDir.absolutePath).\n"+
				"nextDir.absolutePath='%v'\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				nextDir.absolutePath,
				err.Error())

			errs = append(errs, err2)

			continue
		}

		file2LoopIsDone = false

		for !file2LoopIsDone {

			nameFileInfos, err = dirPtr.Readdir(0)

			if err != nil && err == io.EOF {

				file2LoopIsDone = true

				if len(nameFileInfos) == 0 {

					break
				}

			} else if err != nil {

				err2 = fmt.Errorf("%v\n"+
					"Error returned by dirPtr.Readdir(0).\n"+
					"Error= \n%v\n",
					ePrefix.String(),
					err.Error())

				errs = append(errs, err2)

				file2LoopIsDone = true

				break
			}

			for _, nameFInfo := range nameFileInfos {

				if nameFInfo.IsDir() {
					// This is a directory

					err = dirs.AddDirMgrByPathNameStr(
						nextDir.absolutePath +
							osPathSepStr +
							nameFInfo.Name())

					if err != nil {
						err2 = fmt.Errorf("%v\n"+
							"Error returned by dirs.AddDirMgrByPathNameStr(newDir).\n"+
							"newDir='%v'\n"+
							"Error= \n%v\n",
							ePrefix.String(),
							nextDir.absolutePath+osPathSepStr+nameFInfo.Name(),
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
					fh.FilterFileName(nameFInfo,
						fileSelectCriteria,
						ePrefix)

				if err != nil {

					err2 = fmt.Errorf("%v\n"+
						"Error returned by fh.FilterFileName(nameFInfo, fileSelectCriteria).\n"+
						"Directory='%v'\n"+
						"File Name='%v'\n"+
						"Error= \n%v\n",
						ePrefix.String(),
						nextDir.absolutePath,
						nameFInfo.Name(),
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
						file2LoopIsDone = true
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
							file2LoopIsDone = true
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
						osPathSepStr + nameFInfo.Name()

					targetFile = nextTargetDMgr.absolutePath +
						osPathSepStr + nameFInfo.Name()

					err = dMgrHlprMolecule.
						lowLevelCopyFile(
							srcFile,
							nameFInfo,
							targetFile,
							"srcFile",
							"destinationFile",
							ePrefix)

					if err != nil {

						errs = append(errs, err)

						dTreeCopyStats.FilesNotCopied++
						dTreeCopyStats.FileBytesNotCopied += uint64(nameFInfo.Size())

					} else {

						dTreeCopyStats.FilesCopied++
						dTreeCopyStats.FileBytesCopied += uint64(nameFInfo.Size())
					}

				} else {
					// This file is NOT A Match
					// NOT Selected File
					dTreeCopyStats.FilesNotCopied++
					dTreeCopyStats.FileBytesNotCopied += uint64(nameFInfo.Size())

				}
			} // End of range nameFileInfos
		} // End of file 2 Loop

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
