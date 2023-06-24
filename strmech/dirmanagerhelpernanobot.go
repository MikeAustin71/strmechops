package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
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
// To qualify as a selected file, the file entry must
// comply with two filters: File Type and File
// Characteristics.
//
// To be eligible for the copy operation, the file must
// first comply with the specified File Type criteria. In
// terms of File Type, files are classified as
// directories, regular files, SymLink files or other
// non-regular files. Since this method does NOT copy
// directories, the only allowed file types are Regular
// Files, SymLink Files and Other Non-Regular Files. For
// an explanation of Regular and Non-Regular files, see
// the Definition of Terms section below.
//
// Screening criteria for File Type is controlled by the
// following three input parameters:
//
//	copyRegularFiles bool
//	copySymLinkFiles bool
//	copyOtherNonRegularFiles bool
//
// File Types eligible for this copy operation include
// Regular Files such as text files, image files and
// executable files, SymLink files and other Non-Regular
// Files such as device files, named pipes and sockets.
//
// In addition to File Type, selected files must also
// comply with the File Characteristics criteria
// specified by input parameter, 'fileSelectCriteria'.
// The File Characteristics Selection criteria allows
// users to screen files for File Name, File Modification
// Date and File Mode.
// 'fileSelectCriteria'.
//
// The selected files are copied by a Copy IO operation.
// For information on the Copy IO procedure see
// FileHelper{}.CopyFileByIo() method and reference:
//
//	https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
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
//	(1)	This method copies ALL THE FILES residing in the
//		entire the directory tree identified by input
//		parameter 'sourceDMgr' to the directory tree
//		identified by input parameter 'targetDMgr'.
//
//	(2)	If a directory in the target directory tree does
//		not exist, this method will attempt to create it.
//
//	(3)	Files will only be copied if they meet the File
//		Type criteria and the File Characteristics
//		Criteria.
//
//		File Type criteria are specified by input
//		parameters:
//
//			copyRegularFiles bool
//			copySymLinkFiles bool
//			copyOtherNonRegularFiles bool
//
//		File Characteristics Selection criteria are
//		specified by input parameter
//		'fileSelectCriteria'.
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
//	returnCopiedFilesList			bool
//
//		If input parameter 'returnCopiedFilesList' is set
//		to 'true', this method will return a populated
//		File Manager Collection documenting all the files
//		actually included in the directory tree copy
//		operation.
//
//		If input parameter 'returnCopiedFilesList' is set
//		to 'false', this method will return an empty and
//		unpopulated instance of FileMgrCollection. This
//		means that the files actually copied by this
//		method will NOT be documented.
//
//	skipTopLevelDirectory		bool
//
//		If this parameter is set to 'true', the top level
//		source directory will be skipped, and it will not
//		be copied to the directory tree identified by
//		input parameter 'targetDMgr'.
//
//	copyEmptyTargetDirectory	bool
//
//		If set to 'true' the target directory will be
//		created regardless of whether any files are
//		copied to that directory. Remember that files are
//		only copied to the target directory if they meet
//		the File Type and File Characteristics selection
//		criteria.
//
//	copyRegularFiles			bool
//
//		If this parameter is set to 'true', Regular Files,
//		which also meet the File Selection Characteristics
//		criteria (fileSelectCriteria), will be included
//		in the copy operation.
//
//		Regular Files include text files, image files and
//		executable files.
//
//		For an explanation of Regular and Non-Regular
//		files, see the section on "Definition Of Terms",
//		above.
//
//		If input parameters 'copyRegularFiles',
//		'copySymLinkFiles' and 'copyOtherNonRegularFiles'
//		are all set to 'false', an error will be returned.
//
//	copySymLinkFiles				bool
//
//		If this parameter is set to 'true', SymLink Files
//		which also meet the File Selection Characteristics
//		criteria (fileSelectCriteria), will be included
//		in the copy operation.
//
//		If input parameters 'copyRegularFiles',
//		'copySymLinkFiles' and 'copyOtherNonRegularFiles'
//		are all set to 'false', an error will be returned.
//
//	copyOtherNonRegularFiles		bool
//
//		If this parameter is set to 'true', other
//		non-regular file types, besides SymLinks and
//		directories specified above, will be included
//		in the copy operation if they meet the file
//		selection criteria.
//
//		Examples of other non-regular file types
//		include device files, named pipes, and sockets.
//		See the Definition Of Terms section above.
//
//	fileSelectCriteria			FileSelectionCriteria
//
//		In addition to the File Type Selection Criteria,
//		selected files must conform to the File
//		Characteristics criteria specified by
//		'fileSelectCriteria'.
//
//		File Characteristics Selection criteria allows
//		users to screen files for File Name, File
//		Modification Date and File Mode.
//
//		Files matching these selection criteria, and the
//		File Type filter, will be included in the copy
//		operation performed by this method.
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
//			SelectByFileMode  	FilePermissionConfig
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
//			SelectCriterionMode	FileSelectCriterionMode
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
//		If this method completes successfully, this
//		return parameter will be populated with
//		information and statistics on the file copy
//		operation. This information includes the number
//		of files actually copied.
//
//		The data elements in this structure are used
//		to accumulate statistics and information
//		related to a file copy operation performed on
//		source and destination directory trees.
//
//		type DirTreeCopyStats struct {
//			TotalDirsScanned uint64
//				The total number of directories scanned
//				during the current directory tree copy
//				operation.
//
//			DirsCopied uint64
//				The number of directories copied.
//
//			DirsCreated uint64
//				The number of target directories created.
//
//			TotalFilesProcessed uint64
//				The total number of files processed during
//				the directory tree copy operation.
//
//			FilesCopied uint64
//				The total number of files copied to the
//				target directory tree during the directory
//				tree copy operation.
//
//			FileBytesCopied uint64
//				The total number of file bytes copied to the
//				target directory tree during the directory
//				tree copy operation.
//
//			FilesNotCopied uint64
//				The total number of files scanned and
//				processed, but NOT copied to the target
//				directory tree during the directory tree
//				copy operation.
//
//			FileBytesNotCopied uint64
//				The total number of bytes associated with
//				files scanned and processed, but NOT copied
//				to the target directory tree during the
//				directory tree copy operation.
//
//			SubDirs uint64
//				The total number of subdirectories identified
//				during the directory tree copy operation. This
//				does NOT include the parent directory.
//
//			SubDirsDocumented uint64
//				The number of subdirectories identified
//				and returned in a Directory Manager
//				Collection. Does NOT include the parent
//				directory. Subdirectories are only
//				documented if requested. This computation
//				value is therefore optional.
//
//			CopiedFilesDocumented uint64
//				The number of copied files documented
//				by adding a File Manager object to a
//				returned File Manager Collection.
//
//			Errors []error
//				An array of errors associated with the
//				calculation of these statistics.
//		}
//
//	copiedDirTreeFiles			FileMgrCollection
//
//		If input parameter 'returnCopiedFilesList' is set
//		to 'true', 'copiedDirTreeFiles' will return a
//		populated File Manager Collection including all
//		the files actually included in the directory tree
//		copy operation.
//
//		If input parameter 'returnCopiedFilesList' is set
//		to 'false', 'copiedDirTreeFiles' will return an
//		empty and unpopulated instance of
//		FileMgrCollection. This means that the files
//		actually copied by this method will NOT be
//		documented.
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
func (dMgrHlprNanobot *dirMgrHelperNanobot) copyDirectoryTree(
	sourceDMgr *DirMgr,
	targetDMgr *DirMgr,
	returnCopiedFilesList bool,
	skipTopLevelDirectory bool,
	copyEmptyTargetDirectory bool,
	copyRegularFiles bool,
	copySymLinkFiles bool,
	copyOtherNonRegularFiles bool,
	fileSelectCriteria FileSelectionCriteria,
	sourceDMgrLabel string,
	targetDMgrLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	dTreeCopyStats DirTreeCopyStats,
	copiedDirTreeFiles FileMgrCollection,
	nonfatalErrs []error,
	fatalErr error) {

	if dMgrHlprNanobot.lock == nil {
		dMgrHlprNanobot.lock = new(sync.Mutex)
	}

	dMgrHlprNanobot.lock.Lock()

	defer dMgrHlprNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "dirMgrHelperNanobot.copyDirectoryTree()"

	nonfatalErrs = make([]error, 0)

	ePrefix,
		fatalErr = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if fatalErr != nil {

		return dTreeCopyStats,
			copiedDirTreeFiles,
			nonfatalErrs,
			fatalErr
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

		return dTreeCopyStats,
			copiedDirTreeFiles,
			nonfatalErrs,
			fatalErr
	}

	if len(targetDMgrLabel) == 0 {

		targetDMgrLabel = "targetDMgr"
	}

	baseSourceDirLen := len(sourceDMgr.absolutePath)

	_,
		_,
		fatalErr = dMgrHlprPreon.
		validateDirMgr(
			targetDMgr,
			false,
			targetDMgrLabel,
			ePrefix)

	if fatalErr != nil {

		return dTreeCopyStats,
			copiedDirTreeFiles,
			nonfatalErrs,
			fatalErr
	}

	var sourceDirectories = new(DirMgrCollection).New()

	var err error

	if !skipTopLevelDirectory {

		err = sourceDirectories.
			AddDirMgr(
				*sourceDMgr,
				ePrefix.XCpy("sourceDMgr"))

		if err != nil {

			fatalErr = fmt.Errorf("%v\n"+
				"Fatal Error: Failed to add %v to subdirectory collection!\n"+
				"%v = %v\n"+
				"Error= \n%v\n",
				funcName,
				sourceDMgrLabel,
				sourceDMgrLabel,
				sourceDMgr.absolutePath,
				err.Error())

			return dTreeCopyStats,
				copiedDirTreeFiles,
				nonfatalErrs,
				fatalErr
		}

	}

	_,
		_,
		err = new(dirMgrHelperElectron).
		getSubDirsFilesInDirTree(
			sourceDMgr,
			true,                    // getSubdirectories
			false,                   // includeSubDirCurrenDirOneDot
			false,                   // includeSubDirParentDirTwoDots
			false,                   // getRegularFiles
			false,                   // getSymLinksFiles
			false,                   // getOtherNonRegularFiles
			FileSelectionCriteria{}, // subDirSelectCharacteristics
			FileSelectionCriteria{}, // fileSelectCriteria
			&sourceDirectories,      // subDirsInDir
			nil,                     // filesInDir
			sourceDMgrLabel,
			ePrefix)

	if err != nil {

		fatalErr = fmt.Errorf("%v\n"+
			"Failed to extract subdirectories from %v!\n"+
			"Error returned by dirMgrHelperElectron."+
			"getAllSubDirsInDirTree(%v).\n"+
			"%v = %v\n"+
			"Error= \n%v\n",
			funcName,
			sourceDMgrLabel,
			sourceDMgrLabel,
			sourceDMgrLabel,
			sourceDMgr.absolutePath,
			err.Error())

		return dTreeCopyStats,
			copiedDirTreeFiles,
			nonfatalErrs,
			fatalErr
	}

	if len(sourceDirectories.dirMgrs) == 0 {

		return dTreeCopyStats,
			copiedDirTreeFiles,
			nonfatalErrs,
			fatalErr
	}

	if skipTopLevelDirectory == true {

		dTreeCopyStats.TotalSubDirs =
			uint64(len(sourceDirectories.dirMgrs))

	} else {

		dTreeCopyStats.TotalSubDirs =
			uint64(len(sourceDirectories.dirMgrs)) - 1

	}

	var sourceDirMgr, targetDirMgr DirMgr

	var errStatus ArrayColErrorStatus

	var errs2 []error
	var err2 error
	var copiedDirFiles FileMgrCollection
	var dirCopyStats DirectoryCopyStats
	dMgrHlprPlanck := new(dirMgrHelperPlanck)
	cycleCount := 0

	for cycleCount > -1 {

		cycleCount++

		sourceDirMgr,
			errStatus = sourceDirectories.
			PopFirstDirMgr(
				ePrefix.XCpy(
					"sourceDirectories"))

		if errStatus.ProcessingError != nil {

			if errStatus.IsArrayCollectionEmpty {

				cycleCount = -1

				break

			}

			fatalErr = fmt.Errorf("%v\n"+
				"Failed to extract 'sourceDirMgr' from source directories collection!\n"+
				"Error returned by sourceDirectories.PopFirstDirMgr().\n"+
				"cycleCount= '%v'\n"+
				"Error= \n%v\n",
				funcName,
				cycleCount,
				errStatus.ProcessingError.Error())

			return dTreeCopyStats,
				copiedDirTreeFiles,
				nonfatalErrs,
				fatalErr
		}

		targetDirMgr,
			err2 = new(DirMgr).New(
			targetDMgr.absolutePath+
				sourceDirMgr.absolutePath[baseSourceDirLen:],
			ePrefix.XCpy("targetDirMgr<-sourceDirMgr"))

		if err2 != nil {

			fatalErr = fmt.Errorf("%v\n"+
				"Failed to create target Directory Manager for copy operation!\n"+
				"Error returned by DirMgr.New().\n"+
				"Loop Count= %v\n"+
				"%v = %v\n"+
				"%v = %v"+
				"New Source Directory= %v\n"+
				"Error= \n%v\n",
				funcName,
				cycleCount,
				sourceDMgrLabel,
				sourceDMgr.absolutePath,
				targetDMgrLabel,
				targetDMgr.absolutePath,
				sourceDirMgr.absolutePath,
				err2.Error())

			return dTreeCopyStats,
				copiedDirTreeFiles,
				nonfatalErrs,
				fatalErr
		}

		dirCopyStats,
			errs2,
			err2 = dMgrHlprPlanck.
			copyDirectoryFiles(
				&sourceDirMgr,
				&targetDirMgr,
				returnCopiedFilesList, // returnCopiedFilesList
				false,                 // returnSubDirsList
				copyEmptyTargetDirectory,
				copyRegularFiles,
				copySymLinkFiles,
				copyOtherNonRegularFiles,
				fileSelectCriteria,
				"sourceDirMgr",
				"targetDirMgr",
				&sourceDirectories,
				&copiedDirFiles,
				ePrefix)

		if err2 != nil {

			fatalErr = fmt.Errorf("%v\n"+
				"Fatal Error occurred while copying source directory to target directory!\n"+
				"Loop Count= %v\n"+
				"Error returned by dirMgrHelperPlanck.copyDirectoryFiles().\n"+
				"%v = %v\n"+
				"%v = %v"+
				"Error= \n%v\n",
				funcName,
				cycleCount,
				"sourceDirMgr",
				sourceDirMgr.absolutePath,
				"targetDirMgr",
				targetDirMgr.absolutePath,
				err2.Error())

			return dTreeCopyStats,
				copiedDirTreeFiles,
				nonfatalErrs,
				fatalErr
		}

		if len(errs2) > 0 {

			nonfatalErrs = append(
				nonfatalErrs,
				errs2...)

		}

		dTreeCopyStats.AddDirCopyStats(dirCopyStats)

		if returnCopiedFilesList {

			err2 = copiedDirTreeFiles.
				AddFileMgrCollection(
					&copiedDirFiles,
					ePrefix.XCpy("copiedDirFiles"))

			if err2 != nil {

				nonfatalErrs = append(
					nonfatalErrs, err2)
			}

		}
	}

	return dTreeCopyStats,
		copiedDirTreeFiles,
		nonfatalErrs,
		fatalErr
}

// deleteDirectoryTreeFiles
//
// Deletes files in a directory tree based on the File
// Type and File Characteristics criteria passed as input
// parameters. No parent directories or subdirectories
// will be deleted.
//
// The files to be deleted are selected according to file
// to two sets of criteria, File Type and File
// Characteristics.
//
// First, the file must comply with the specified File
// Type criteria. In terms of File Type, files are
// classified as directories, regular files, SymLink
// files or other non-regular files.
//
// This method does NOT delete directories.
//
// For an explanation of Regular and Non-Regular files,
// see the Definition of Terms section below.
//
// Screening criteria for File Type is controlled by the
// following three input parameters:
//
//	deleteRegularFiles - bool
//	deleteSymLinkFiles - bool
//	deleteOtherNonRegularFiles - bool
//
// This method deletes regular files plus certain
// non-regular files depending on input parameter values
// supplied by the user.
//
// In addition to File Type, selected files must comply
// with the File Characteristics criteria specified by
// input parameter, 'fileSelectCriteria'. The File
// Characteristics Selection criteria allows users to
// screen files for File Name, File Modification Date and
// File Mode.
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
//	https://www.computerhope.com/jargon/r/regular-file.htm
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
// # IMPORTANT
//
//	(1)	This method will delete files in the entire
//		directory tree specified by input parameter
//		'targetDMgr'. The files to be deleted must match
//		the File Selection Characteristics Criteria passed
//		as input parameter 'fileSelectCriteria'.
//
//	(2) In addition to meeting the File Selection
//		Characteristics Criteria specified in paragraph
//		(1) above, files eligible for deletion must comply
//		with File Type Criteria passed as input parameters
//		'deleteRegularFiles', 'deleteSymLinkFiles' and
//		'deleteOtherNonRegularFiles'.
//
//	(3) If the target directory tree identified by input
//		parameter 'targetDMgr' contains NO Files
//		meeting the (1) File Selection Characteristics
//		Criteria and the (2) File Type Selection
//		Criteria, this method will exit, and no error
//		will be returned.
//
//	(4) If the target directory tree identified by input
//		parameter 'targetDMgr' contains NO Files
//		(0 Files), this method will exit and no error
//		will be returned.
//
//	(5)	This method will NOT delete directories.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	targetDMgr					*DirMgr
//
//		An instance of DirMgr which identifies the
//		target directory from which files matching the
//		File Selection Criteria ('fileSelectCriteria')
//		will be deleted.
//
//		If the target directory identified by
//		'targetDMgr' contains NO Files (0 Files), this
//		method will exit and no error will be returned.
//
//		If the target directory identified by
//		'targetDMgr' contains NO Files matching the File
//		Selection Criteria specified by input parameter
//		'fileSelectCriteria', this method will exit and
//		no error will be returned.
//
//	returnDeletedFilesList		bool
//
//		If this parameter is set to 'true', the return
//		parameter 'deletedFiles' will be returned as a
//		populated instance of File Manager Collection
//		(FileMgrCollection). This collection will contain
//		an array of File Manager (FileMgr) objects
//		identifying all the files deleted in the current
//		file deletion operation for the designated target
//		directory tree.
//
//		If 'returnDeletedFilesList' is set to 'false',
//		the instance of FileMgrCollection returned by
//		this method will always be empty and unpopulated.
//		This means that the files actually deleted by
//		this method will NOT be documented.
//
//	skipTopLevelDirectory		bool
//
//		If this parameter is set to 'true' it means that
//		no files in the top level or parent directory
//		specified by input parameter 'targetDMgr' will
//		be deleted.
//
//	deleteRegularFiles			bool
//
//		If this parameter is set to 'true', regular files
//		will be eligible for deletion if they meet the
//		File Selection criteria specified by input
//		paramter 'fileSelectCriteria'.
//
//		Regular Files include text files, image files and
//		executable files.
//
//		For an explanation of Regular and Non-Regular
//		files, see the section on "Definition Of Terms",
//		above.
//
//		If 'deleteRegularFiles', 'deleteSymLinkFiles' and
//		'deleteOtherNonRegularFiles' are all set to
//		'false', an error will be returned.
//
//	deleteSymLinkFiles			bool
//
//		If this parameter is set to 'true', SymLink files
//		which meet the file selection criteria, will be
//		deleted.
//
//		If 'deleteRegularFiles', 'deleteSymLinkFiles' and
//		'deleteOtherNonRegularFiles' are all set to
//		'false', an error will be returned.
//
//	deleteOtherNonRegularFiles bool
//
//		If this parameter is set to 'true', other
//		non-regular file types, besides SymLinks and
//		directories, will be deleted if they meet the file
//		selection criteria.
//
//		Examples of other non-regular file types
//		include device files, named pipes, and sockets.
//		See the Definition Of Terms section above.
//
//		Note:	Although directories are non-regular
//				files; however, this method WILL NOT
//				DELETE directories.
//
//		If 'deleteRegularFiles', 'deleteSymLinkFiles' and
//		'deleteOtherNonRegularFiles' are all set to
//		'false', an error will be returned.
//
//	fileSelectCriteria			FileSelectionCriteria
//
//		In addition to the File Type Selection Criteria,
//		selected files must conform to the File
//		Characteristics Criteria specified by this
//		parameter, 'fileSelectCriteria'.
//
//		Failure to comply with File Characteristics
//		Selection Criteria ('fileSelectCriteria') means
//		that the subject file will NOT be deleted.
//
//		File Characteristics Selection criteria allows
//		users to screen files for File Name, File
//		Modification Date and File Mode.
//
//		Files matching these File Characteristics
//		Selection Criteria, and the File Type filter,
//		will be included in the file deletion operation
//		performed by this method.
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
//		If all of the file selection criterion in the FileSelectionCriteria object
//		are 'Inactive' or 'Not Set' (set to their zero or default values), then all
//		the files processed in the target directory will be selected and deleted.
//
//			Example:
//			  fsc := FileSelectCriterionMode{}
//
//			  In this example, 'fsc' is NOT initialized. Therefore,
//			  all the selection criterion are 'Inactive'. Consequently,
//			  all the files encountered in the target directory during
//			  the search operation will be selected and deleted.
//
//		------------------------------------------------------------------------
//
//	targetDMgrLabel string
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
//	deletedDirTreeFileStats		DeleteDirFilesStats
//
//		If this method completes successfully, this
//		return parameter will be populated with
//		information and statistics on the file deletion
//		operation. This information includes the number
//		of files deleted.
//
//		The data elements in the DeleteDirFilesStats
//		structure are used to accumulate statistics and
//		information related to the deletion of files from
//		a target directory tree.
//
//		type DeleteDirFilesStats struct {
//
//			TotalFilesProcessed uint64
//				The total number of files processed.
//				Does NOT include directory entries.
//
//			FilesDeleted uint64
//				The number of files deleted. Does
//				NOT include directory entries.
//
//			FilesDeletedBytes uint64
//				The number of file bytes deleted.
//				Does NOT include directory entries.
//
//			FilesRemaining uint64
//				The number of files processed, but
//				NOT deleted. Does NOT include directory
//				entries.
//
//			FilesRemainingBytes uint64
//				The number of bytes associated with
//				files processed but NOT copied. Does
//				NOT include directory entries.
//
//			TotalSubDirectories uint64
//				Total SubDirectories processed
//
//			TotalDirsScanned uint64
//				Total Directories Scanned.
//
//			NumOfDirsWhereFilesDeleted uint64
//				The number of parent directories and
//				subdirectories where files were deleted.
//
//			DirectoriesDeleted uint64
//				The number of directories deleted.
//
//			SubDirsDocumented uint64
//				The number of subdirectories identified
//				and returned in a Directory Manager
//				Collection. Does NOT include the parent
//				directory. Subdirectories are only
//				documented if requested. This computation
//				value is therefore optional.
//
//			DeletedFilesDocumented uint64
//				The number of deleted files documented
//				by adding a File Manager object to a
//				returned File Manager Collection.
//
//			Errors []error
//				An array of errors associated with the
//				calculation of these statistics.
//		}
//
//	deletedFiles				FileMgrCollection
//
//		If this method completes successfully and input
//		paramter 'returnDeletedFilesList' is set to
//		'true', 'deletedFiles' will return a collection
//		of File Manager objects identifying all the
//		files actually deleted. Again, this return
//		parameter will ONLY be populated when input
//		parameter 'returnDeletedFilesList' is set to
//		'true'.
//
//	remainingTargetDirTreeStats	DirectoryProfile
//
//		If this method completes successfully, without
//		errors, this parameter will return an instance
//		of DirectoryProfile containing directory profile
//		information on the target directory tree
//		('targetDMgr') after target files have been
//		deleted.
//
//		If 'skipTopLevelDirectory' is set to false,
//		this returned Directory Profile will include
//		the parent directory defined by 'targetDMgr'.
//		Conversely, if 'skipTopLevelDirectory', the
//		statistics will only include the subdirectories
//		below 'targetDMgr'.
//
//		type DirectoryProfile struct {
//
//			ParentDirAbsolutePath 			string
//				The absolute directory path for the
//				directory described by this profile
//				information.
//
//			ParentDirManager					DirMgr
//				An instance of DirMgr encapsulating the
//				Directory Path and associated parameters
//				for the directory described by this profile
//				information.
//
//			DirExistsOnStorageDrive 	bool
//				If 'true', this paramter signals
//				that the directory actually exists on
//				a storage drive.
//
//			ParentDirIsIncludedInStats bool
//				If this parameter is set to 'true', it
//				signals that the directory statistics and
//				information provided by this instance of
//				DirectoryProfile includes metrics from
//				the parent directory.
//
//			DirTotalFiles				uint64
//				The number of total files, of all types,
//				residing in the subject directory. This
//				includes directory entry files, Regular
//				Files, SymLink Files and Non-Regular
//				Files.
//
//			DirTotalFileBytes			uint64
//				The size of all files, of all types,
//				residing in the subject directory
//				expressed in bytes. This includes
//				directory entry files, Regular Files,
//				SymLink Files and Non-Regular Files.
//
//			DirSubDirectories			uint64
//				The number of subdirectories residing
//				within the subject directory. This
//
//			DirSubDirectoriesBytes		uint64
//				The total size of all Subdirectory entries
//				residing in the subject directory expressed
//				in bytes.
//
//			SubDirsIncludeCurrentDirOneDot bool
//				All directories include an os.FileInfo entry for
//				the current directory. The current directory name
//				is always denoted as single dot ('.').
//
//				When data element, 'SubDirsIncludeCurrentDirOneDot',
//				is set to 'true', the one dot current directory ('.')
//				will be included in the directory profile information
//				and counted as a separate subdirectory.
//
//			SubDirsIncludeParentDirTwoDot bool
//				All directories include an os.FileInfo entry for
//				the parent directory. The parent directory name
//				is always denoted as two dots ('..').
//
//				When data element, 'SubDirsIncludeParentDirTwoDot',
//				is set to 'true', the two dot ('..') parent directory,
//				will be included in the directory profile information
//				and counted as a separate subdirectory.
//
//			DirRegularFiles				uint64
//				The number of 'Regular' Files residing
//				within the subject Directory. Regular
//				files include text files, image files
//				and executable files. Reference:
//				https://www.computerhope.com/jargon/r/regular-file.htm
//
//			DirRegularFileBytes			uint64
//				The total size of all 'Regular' files
//				residing in the subject directory expressed
//				in bytes.
//
//			DirSymLinkFiles				uint64
//				The number of SymLink files residing in the
//				subject directory.
//
//			DirSymLinkFileBytes			uint64
//				The total size of all SymLink files
//				residing in the subject directory
//				expressed in bytes.
//
//			DirNonRegularFiles			uint64
//				The total number of Non-Regular files residing
//				in the subject directory.
//
//				Non-Regular files include directories, device
//				files, named pipes, sockets, and symbolic links.
//
//			DirNonRegularFileBytes		uint64
//				The total size of all Non-Regular files residing
//				in the subject directory expressed in bytes.
//
//			Errors						error
//				Computational or processing errors will be
//				recorded through this parameter.
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
//		If a file which meets the selection criteria
//		fails to delete, the error is classified as
//		a "Fatal" error.
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
//
//		If a file which meets the selection criteria
//		fails to delete, the error is classified as
//		a "Fatal" error.
func (dMgrHlprNanobot *dirMgrHelperNanobot) deleteDirectoryTreeFiles(
	targetDMgr *DirMgr,
	returnDeletedFilesList bool,
	skipTopLevelDirectory bool,
	deleteRegularFiles bool,
	deleteSymLinkFiles bool,
	deleteOtherNonRegularFiles bool,
	fileSelectCriteria FileSelectionCriteria,
	targetDMgrLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	deletedDirTreeFileStats DeleteDirFilesStats,
	deletedFiles FileMgrCollection,
	remainingTargetDirTreeStats DirectoryProfile,
	nonfatalErrs []error,
	fatalErr error) {

	if dMgrHlprNanobot.lock == nil {
		dMgrHlprNanobot.lock = new(sync.Mutex)
	}

	dMgrHlprNanobot.lock.Lock()

	defer dMgrHlprNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "dirMgrHelperNanobot.deleteDirectoryTreeFiles()"

	nonfatalErrs = make([]error, 0)

	ePrefix,
		fatalErr = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if fatalErr != nil {

		return deletedDirTreeFileStats,
			deletedFiles,
			remainingTargetDirTreeStats,
			nonfatalErrs,
			fatalErr
	}

	if len(targetDMgrLabel) == 0 {

		targetDMgrLabel = "targetDMgr"
	}

	_,
		_,
		fatalErr = new(dirMgrHelperPreon).
		validateDirMgr(
			targetDMgr,
			false,
			targetDMgrLabel,
			ePrefix.XCpy(targetDMgrLabel))

	if fatalErr != nil {

		return deletedDirTreeFileStats,
			deletedFiles,
			remainingTargetDirTreeStats,
			nonfatalErrs,
			fatalErr
	}

	var targetDirs DirMgrCollection
	var err2 error

	if !skipTopLevelDirectory {

		err2 = targetDirs.
			AddDirMgr(
				*targetDMgr,
				ePrefix.XCpy(targetDMgrLabel))

		if err2 != nil {

			fatalErr = fmt.Errorf("%v\n"+
				"Error adding %v to Directory Manager Collection.\n"+
				"targetDirs.AddDirMgr() returned an error."+
				"%v Path = %v\n"+
				"Error=\n%v\n",
				funcName,
				targetDMgrLabel,
				targetDMgrLabel,
				targetDMgr.absolutePath,
				err2.Error())

			return deletedDirTreeFileStats,
				deletedFiles,
				remainingTargetDirTreeStats,
				nonfatalErrs,
				fatalErr
		}
	}

	_,
		_,
		err2 = new(dirMgrHelperElectron).
		getSubDirsFilesInDirTree(
			targetDMgr,
			true,                    // getSubdirectories
			false,                   // includeSubDirCurrenDirOneDot
			false,                   // includeSubDirParentDirTwoDots
			false,                   // getRegularFiles
			false,                   // getSymLinksFiles
			false,                   // getOtherNonRegularFiles
			FileSelectionCriteria{}, // subDirSelectCharacteristics
			FileSelectionCriteria{}, // fileSelectCriteria
			&targetDirs,             // subDirsInDir
			nil,                     // filesInDir
			targetDMgrLabel,
			ePrefix)

	if err2 != nil {

		fatalErr = fmt.Errorf("%v\n"+
			"Error extracting subdirectories from %v parent directory.\n"+
			"dirMgrHelperTachyon.getSubdirectories() returned an error.\n"+
			"%v Path = %v\n"+
			"Error=\n%v\n",
			funcName,
			targetDMgrLabel,
			targetDMgrLabel,
			targetDMgr.absolutePath,
			err2.Error())

		return deletedDirTreeFileStats,
			deletedFiles,
			remainingTargetDirTreeStats,
			nonfatalErrs,
			fatalErr
	}

	if len(targetDirs.dirMgrs) == 0 {

		// There are no subdirectories
		_,
			remainingTargetDirTreeStats,
			fatalErr = new(dirMgrHelperPreon).
			getDirectoryTreeProfile(
				targetDMgr,
				skipTopLevelDirectory,
				false, // includeSubDirCurrenDirOneDot
				false, // includeSubDirParentDirTwoDots
				FileSelectionCriteria{},
				"dMgr",
				ePrefix)

		return deletedDirTreeFileStats,
			deletedFiles,
			remainingTargetDirTreeStats,
			nonfatalErrs,
			fatalErr

	}

	if skipTopLevelDirectory == true {

		deletedDirTreeFileStats.TotalSubDirectories =
			uint64(len(targetDirs.dirMgrs))

	} else {

		deletedDirTreeFileStats.TotalSubDirectories =
			uint64(len(targetDirs.dirMgrs)) - 1

	}

	var dMgrHlprPlanck = new(dirMgrHelperPlanck)
	var deletedDirFileStats DeleteDirFilesStats
	var deletedSubDirFiles FileMgrCollection
	var errs2 []error
	var nextTargetDir DirMgr

	var idx = -1

	var errStatus ArrayColErrorStatus
	var dMgrColHelper = new(dirMgrCollectionHelper)

	for idx > -2 {

		idx++

		nextTargetDir,
			errStatus = dMgrColHelper.
			peekOrPopAtIndex(
				&targetDirs,
				idx,
				false, // deleteIndex
				ePrefix.XCpy("targetDirs"))

		if errStatus.ProcessingError != nil {

			if errStatus.IsArrayCollectionEmpty {

				idx = -2

				break

			}

			fatalErr = fmt.Errorf("%v\n"+
				"Failed to extract first 'nextTargetDir' from target\n"+
				"directories collection (targetDirs)!\n"+
				"idx= '%v'\n"+
				"Error returned by sourceDirectories.PopFirstDirMgr().\n"+
				"Error= \n%v\n",
				funcName,
				idx,
				errStatus.ProcessingError.Error())

			return deletedDirTreeFileStats,
				deletedFiles,
				remainingTargetDirTreeStats,
				nonfatalErrs,
				fatalErr
		}

		deletedDirFileStats,
			_,
			errs2,
			err2 = dMgrHlprPlanck.
			deleteDirectoryFiles(
				&nextTargetDir,
				returnDeletedFilesList,
				false, // returnSubDirsList
				deleteRegularFiles,
				deleteSymLinkFiles,
				deleteOtherNonRegularFiles,
				fileSelectCriteria,
				"nextTargetDir",
				&targetDirs,
				&deletedFiles,
				ePrefix.XCpy("nextTargetDir"))

		if len(errs2) > 0 {

			nonfatalErrs = append(
				nonfatalErrs,
				err2)

		}

		if err2 != nil {

			fatalErr = fmt.Errorf("%v\n"+
				"Fatal Error occurred while deleting files.\n"+
				"dMgrHlprPlanck.deleteDirectoryFiles() returned an error.\n"+
				"nextTargetDir= %v\n"+
				"idx= '%v'\n"+
				"Error=\n%v\n",
				funcName,
				nextTargetDir.absolutePath,
				idx,
				err2.Error())

			return deletedDirTreeFileStats,
				deletedFiles,
				remainingTargetDirTreeStats,
				nonfatalErrs,
				fatalErr
		}

		deletedDirTreeFileStats.
			AddStats(deletedDirFileStats)

		if returnDeletedFilesList {

			deletedFiles.fileMgrs =
				append(
					deletedFiles.fileMgrs,
					deletedSubDirFiles.fileMgrs...)

		}

	}

	_,
		remainingTargetDirTreeStats,
		fatalErr = new(dirMgrHelperPreon).
		getDirectoryTreeProfile(
			targetDMgr,
			skipTopLevelDirectory,
			false, // includeSubDirCurrenDirOneDot
			false, // includeSubDirParentDirTwoDots
			FileSelectionCriteria{},
			"dMgr",
			ePrefix)

	targetDirs.Empty()

	return deletedDirTreeFileStats,
		deletedFiles,
		remainingTargetDirTreeStats,
		nonfatalErrs,
		fatalErr
}

// deleteAllSubDirectories
//
// The directory identified by the input parameter 'dMgr'
// is treated as the parent directory.
//
// This method will proceed to delete all directories and
// files which are subsidiary to the parent directory,
// or top level directory, identified by 'dMgr'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// All subdirectories and files which are subordinate to
// the parent or top level directory identified by 'dMgr'
// will be deleted.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dMgr						*DirMgr
//
//		A pointer to an instance of DirMgr. All
//		subdirectories and files subsidiary to the parent
//		or top level directory identified by 'dMgr' will
//		be deleted.
//
//		If the directory identified by 'dMgr' does NOT
//		exist on persistent (drive) storage, an error will
//		be returned.
//
//	returnDeletedSubDirs		bool
//
//	deletedSubDirs				*DirMgrCollection
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
//	err							error
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
func (dMgrHlprNanobot *dirMgrHelperNanobot) deleteAllSubDirectories(
	dMgr *DirMgr,
	returnDeletedSubDirs bool,
	deletedSubDirs *DirMgrCollection,
	dMgrLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if dMgrHlprNanobot.lock == nil {
		dMgrHlprNanobot.lock = new(sync.Mutex)
	}

	dMgrHlprNanobot.lock.Lock()

	defer dMgrHlprNanobot.lock.Unlock()

	funcName := "dirMgrHelperNanobot.deleteAllSubDirectories() "

	var ePrefix *ePref.ErrPrefixDto

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

	if returnDeletedSubDirs == true &&
		deletedSubDirs == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameters 'returnDeletedSubDirs'\n"+
			"and 'deletedSubDirs' are conflicted!\n"+
			"'returnDeletedSubDirs' is set to 'true', but\n"+
			"'deletedSubDirs' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	_,
		_,
		err = new(dirMgrHelperPreon).
		validateDirMgr(
			dMgr,
			true, // Path Must Exist
			dMgrLabel,
			ePrefix.XCpy(
				dMgrLabel))

	if err != nil {

		return err
	}

	var targetSubDirs DirMgrCollection
	_,
		_,
		err = new(dirMgrHelperBoson).
		getSubDirsFilesInDir(
			dMgr,
			true,                    // getSubdirectories
			false,                   // includeSubDirCurrenDirOneDot
			false,                   // includeSubDirParentDirTwoDots
			false,                   // getRegularFiles
			false,                   // getSymLinksFiles
			false,                   // getOtherNonRegularFiles
			FileSelectionCriteria{}, // subDirSelectCharacteristics
			FileSelectionCriteria{}, // fileSelectCriteria
			&targetSubDirs,          // subDirsInDir
			nil,                     // filesInDir
			dMgrLabel,               // targetDMgrLabel
			ePrefix.XCpy(dMgrLabel))

	if err != nil {

		return err
	}

	lenSubDirs := len(targetSubDirs.dirMgrs)

	if lenSubDirs == 0 {

		return err
	}

	var err2 error
	var dMgrHlprMolecule = new(dirMgrHelperMolecule)

	for i := 0; i < lenSubDirs; i++ {

		err2 = dMgrHlprMolecule.
			lowLevelDeleteDirectoryAll(
				&targetSubDirs.dirMgrs[i],
				fmt.Sprintf("targetSubDirs.dirMgrs[%v]",
					i),
				ePrefix)

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Error occurred while deleting target subdirectory.\n"+
				"dirMgrHelperMolecule.lowLevelDeleteDirectoryAll(targetSubDirs.dirMgrs[%v])\n"+
				"targetSubDirs.dirMgrs[%v]= '%v'\n"+
				"Error=\n%v\n",
				funcName,
				i,
				i,
				targetSubDirs.dirMgrs[i].absolutePath,
				err2.Error())

			return err
		}

	}

	if returnDeletedSubDirs == true {

		err = deletedSubDirs.AddDirMgrCollection(
			&targetSubDirs,
			ePrefix.XCpy("targetSubDirs"))

	}

	return err
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
