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

	_,
		_,
		fatalErr = dMgrHlprPreon.
		validateDirMgr(
			targetDMgr,
			false,
			targetDMgrLabel,
			ePrefix)

	if fatalErr != nil {

		return dTreeCopyStats, nonfatalErrs, fatalErr
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

			return dTreeCopyStats, nonfatalErrs, fatalErr
		}

	}

	var subdirectories DirMgrCollection

	subdirectories,
		_,
		err = new(dirMgrHelperTachyon).
		getSubdirectories(
			sourceDMgr,
			sourceDMgrLabel,
			ePrefix)

	if err != nil {

		fatalErr = fmt.Errorf("%v\n"+
			"Failed to extract subdirectories from %v!\n"+
			"Error returned by dirMgrHelperTachyongetSubdirectories(%v).\n"+
			"%v = %v\n"+
			"Error= \n%v\n",
			funcName,
			sourceDMgrLabel,
			sourceDMgrLabel,
			sourceDMgrLabel,
			sourceDMgr.absolutePath,
			err.Error())

		return dTreeCopyStats, nonfatalErrs, fatalErr
	}

	err = sourceDirectories.
		AddDirMgrCollection(
			&subdirectories,
			ePrefix.XCpy("subdirectories"))

	if err != nil {

		fatalErr = fmt.Errorf("%v\n"+
			"Failed to add subdirectories to sourceDirectories Collection!\n"+
			"Error returned by sourceDirectories.AddDirMgrCollection(subdirectories).\n"+
			"%v = %v\n"+
			"Error= \n%v\n",
			funcName,
			sourceDMgrLabel,
			sourceDMgr.absolutePath,
			err.Error())

		return dTreeCopyStats, nonfatalErrs, fatalErr
	}

	var sourceDirMgr, targetDirMgr DirMgr

	sourceDirMgr, err = sourceDirectories.PopFirstDirMgr(
		ePrefix.XCpy("sourceDirectories"))

	if err != nil {

		fatalErr = fmt.Errorf("%v\n"+
			"Failed to extract first 'sourceDirMgr' from source directories collection!\n"+
			"Error returned by sourceDirectories.PopFirstDirMgr().\n"+
			"Error= \n%v\n",
			funcName,
			err.Error())

		return dTreeCopyStats, nonfatalErrs, fatalErr

	}

	var errs2 []error
	var err2 error
	var dirCopyStats DirectoryCopyStats
	dMgrHlprPlanck := new(dirMgrHelperPlanck)
	loopCount := 1

	for err == nil {

		targetDirMgr,
			err2 = new(DirMgr).New(
			targetDMgr.absolutePath+
				sourceDirMgr.absolutePath[baseSourceDirLen:],
			ePrefix.XCpy("targetDirMgr<-"))

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
				loopCount,
				sourceDMgrLabel,
				sourceDMgr.absolutePath,
				targetDMgrLabel,
				targetDMgr.absolutePath,
				sourceDirMgr.absolutePath,
				err2.Error())

			return dTreeCopyStats, nonfatalErrs, fatalErr
		}

		dirCopyStats,
			_,
			subdirectories,
			errs2,
			err2 = dMgrHlprPlanck.
			copyDirectoryFiles(
				&sourceDirMgr,
				&targetDirMgr,
				false, // returnCopiedFilesList
				copyEmptyTargetDirectory,
				copyRegularFiles,
				copySymLinkFiles,
				copyOtherNonRegularFiles,
				fileSelectCriteria,
				"sourceDirMgr",
				"targetDirMgr",
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
				loopCount,
				"sourceDirMgr",
				sourceDirMgr.absolutePath,
				"targetDirMgr",
				targetDirMgr.absolutePath,
				err2.Error())

			return dTreeCopyStats, nonfatalErrs, fatalErr
		}

		if len(errs2) > 0 {

			nonfatalErrs = append(
				nonfatalErrs,
				errs2...)

		}

		dTreeCopyStats.AddDirCopyStats(dirCopyStats)

		err2 = sourceDirectories.
			AddDirMgrCollection(
				&subdirectories,
				ePrefix.XCpy("subdirectories"))

		if err2 != nil {

			fatalErr = fmt.Errorf("%v\n"+
				"Failed to add subdirectories to sourceDirectories Collection!\n"+
				"Loop Count= %v\n"+
				"Error returned by sourceDirectories.AddDirMgrCollection(subdirectories).\n"+
				"%v = %v\n"+
				"Error= \n%v\n",
				funcName,
				loopCount,
				sourceDMgrLabel,
				sourceDMgr.absolutePath,
				err2.Error())

			return dTreeCopyStats, nonfatalErrs, fatalErr
		}

		sourceDirMgr, err = sourceDirectories.PopFirstDirMgr(
			ePrefix.XCpy("sourceDirectories"))

		if err != nil && err != io.EOF {

			fatalErr = fmt.Errorf("%v\n"+
				"Failed to extract 'sourceDirMgr' from source directories collection!\n"+
				"Loop Count= %v\n"+
				"Error returned by sourceDirectories.PopFirstDirMgr().\n"+
				"Error= \n%v\n",
				funcName,
				loopCount,
				err.Error())

			return dTreeCopyStats, nonfatalErrs, fatalErr
		}

		loopCount++
	}

	return dTreeCopyStats, nonfatalErrs, fatalErr
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
//		Characteristics requirements specified in paragraph
//		(1) above, files eligible for deletion must comply
//		with File Type specifications passed as input
//		parameters 'deleteRegularFiles',
//		'deleteSymLinkFiles' and
//		'deleteOtherNonRegularFiles'.
//
//	(3) If the target directory identified by input
//		parameter 'targetDMgr' contains NO Files
//		(0 Files), this method will exit and no error
//		will be returned.
//
//	(4) If the target directory identified by input
//		parameter 'targetDMgr' contains NO Files
//		matching the File Selection Criteria specified by
//		input parameter 'fileSelectCriteria', this method
//		will exit and no error will be returned.
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
//		file deletion operation.
//
//		If 'returnDeletedFilesList' is set to 'false',
//		the instance of FileMgrCollection returned by this
//		method will always be empty and unpopulated.
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
//	  This input parameter should be configured with the
//	  desired file selection criteria. Files matching
//	  this criteria will be deleted in the directory
//	  identified by input parameter, 'targetDMgr'.
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
				nonfatalErrs,
				fatalErr
		}
	}

	var subdirectories DirMgrCollection

	subdirectories,
		_,
		err2 = new(dirMgrHelperTachyon).
		getSubdirectories(
			targetDMgr,
			targetDMgrLabel,
			ePrefix.XCpy(targetDMgrLabel))

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
			nonfatalErrs,
			fatalErr
	}

	targetDirs.dirMgrs = append(
		targetDirs.dirMgrs,
		subdirectories.dirMgrs...)

	if len(targetDirs.dirMgrs) == 0 {

		return deletedDirTreeFileStats,
			deletedFiles,
			nonfatalErrs,
			fatalErr

	}

	var dMgrHlprPlanck = new(dirMgrHelperPlanck)
	var deletedDirFileStats DeleteDirFilesStats
	var deletedSubDirFiles FileMgrCollection
	var errs2 []error

	for idx, nextTargetDir := range targetDirs.dirMgrs {

		deletedDirFileStats,
			deletedSubDirFiles,
			errs2,
			err2 = dMgrHlprPlanck.
			deleteDirectoryFiles(
				&nextTargetDir,
				deleteRegularFiles,
				deleteSymLinkFiles,
				deleteOtherNonRegularFiles,
				returnDeletedFilesList,
				fileSelectCriteria,
				"nextTargetDir",
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
				"Error=\n%v\n",
				funcName,
				nextTargetDir.absolutePath,
				err2.Error())

			return deletedDirTreeFileStats,
				deletedFiles,
				nonfatalErrs,
				fatalErr
		}

		deletedDirTreeFileStats.
			AddStats(deletedDirFileStats)

		deletedDirTreeFileStats.TotalSubDirectories =
			uint64(idx + 1)

		if returnDeletedFilesList {

			deletedFiles.fileMgrs =
				append(
					deletedFiles.fileMgrs,
					deletedSubDirFiles.fileMgrs...)

		}
	}

	if !skipTopLevelDirectory {
		deletedDirTreeFileStats.TotalSubDirectories--
	}

	return deletedDirTreeFileStats,
		deletedFiles,
		nonfatalErrs,
		fatalErr

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
