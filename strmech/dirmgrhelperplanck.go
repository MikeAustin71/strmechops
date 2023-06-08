package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"os"
	"sync"
)

type dirMgrHelperPlanck struct {
	lock *sync.Mutex
}

// copyDirectoryFiles
//
// Helper method used by DirMgr. This method copies
// files from the directory identified by DirMgr object
// (sourceDMgr) to a target directory (targetDMgr).
//
// The files to be copied are selected according to
// file selection criteria specified by input parameter,
// 'fileSelectCriteria'.
//
// The selected files are copied by a Copy IO operation.
// For information on the Copy IO procedure see
// FileHelper{}.CopyFileByIo() method and reference:
//
//	https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// This method copies "regular" files plus certain
// non-regular files depending on input parameter values
// supplied by the user.
//
// In Go programming language, a regular file is a file
// that contains data in any format that can be read by
// a user or an application. It is not a directory or a
// device file.
//
// Examples of "regular" files include text files, image
// files and executable files.
//
// Examples of non-regular files include directories,
// device files, named pipes, sockets, and symbolic
// links.
//
// Input parameters 'copyEmptyDirectories',
// 'copySymLinkFiles', and 'copyOtherNonRegularFiles'
// allow the users to specify that these non-regular
// files should be included in the copy operation.
//
// If the target directory does not exist and files are
// located matching the file selection criteria, this
// method will attempt to create the target directory.
// However, if no files meet the file selection criteria
// as defined by input parameter,'fileSelectCriteria',
// this method will NOT attempt to create the target
// directory.
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
//	different from a standard shortcut that, say, a program
//	installer has placed on your Windows desktop to make the
//	program easier to run.
//
//	Sure, clicking on either type of shortcut opens the
//	linked object, but what goes on beneath the hood is
//	different in both cases as we'll see next.
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
// This method ONLY copies files from the directory
// identified by 'sourceDMgr' to the directory identified
// by 'targetDMgr'. It does NOT copy files from
// subdirectories of 'sourceDMgr'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	sourceDMgr					*DirMgr
//
//		An instance of DirMgr which identifies the source
//		from which files will be copied to the directory
//		identified by input parameter 'targetDMgr'.
//
//	targetDMgr					*DirMgr
//
//		An instance of DirMgr which identifies the
//		destination directory to which files from
//		'sourceDMgr' will be copied.
//
//	fileSelectCriteria			FileSelectionCriteria
//
//	  This input parameter should be configured with the
//	  desired file selection criteria. Files matching
//	  this criteria will be copied to the directory
//	  identified by input parameter, 'targetDir'.
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
//	copyEmptyDirectories			bool
//
//		If set to 'true' the target directory will be
//		created regardless of whether any files are
//		copied to that directory. Remember that files are
//		only copied to the target directory if they meet
//		file selection criteria specified by input
//		parameter 'fileSelectCriteria'.
//
//	copySymLinkFiles				bool
//
//		If this parameter is set to 'true', SymLink files
//		which meet the file selection criteria, will be
//		included in the copy operation.
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
//	dirCopyStats				DirectoryCopyStats
//
//		If this method completes successfully, this
//		return parameter will be populated with
//		information and statistics on the copy
//		operation. This information includes the number
//		of files copied.
//
//		type DirectoryCopyStats struct {
//			DirCreated          uint64
//			TotalFilesProcessed uint64
//			FilesCopied         uint64
//			FileBytesCopied     uint64
//			FilesNotCopied      uint64
//			FileBytesNotCopied  uint64
//			ComputeError        error
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
func (dMgrHlprPlanck *dirMgrHelperPlanck) copyDirectoryFiles(
	sourceDMgr *DirMgr,
	targetDMgr *DirMgr,
	fileSelectCriteria FileSelectionCriteria,
	copyEmptyDirectories bool,
	copySymLinkFiles bool,
	copyOtherNonRegularFiles bool,
	sourceDMgrLabel string,
	targetDMgrLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	dirCopyStats DirectoryCopyStats,
	errs []error) {

	if dMgrHlprPlanck.lock == nil {
		dMgrHlprPlanck.lock = new(sync.Mutex)
	}

	dMgrHlprPlanck.lock.Lock()

	defer dMgrHlprPlanck.lock.Unlock()

	var err, err2 error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "dirMgrHelper." +
		"copyDirectoryFiles()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		errs = append(errs, err)

		return dirCopyStats, errs
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
			ePrefix)

	if err != nil {

		errs = append(errs, err)

		return dirCopyStats, errs
	}

	if len(targetDMgrLabel) == 0 {

		targetDMgrLabel = "targetDMgr"
	}

	var targetPathDoesExist bool

	_,
		targetPathDoesExist,
		err = dMgrHlprPreon.
		validateDirMgr(
			targetDMgr,
			false,
			targetDMgrLabel,
			ePrefix)

	if err != nil {

		errs = append(errs, err)

		return dirCopyStats, errs
	}

	var dirCreated bool

	dMgrHlprMolecule := new(dirMgrHelperMolecule)

	if !targetPathDoesExist && copyEmptyDirectories {

		dirCreated,
			err = dMgrHlprMolecule.lowLevelMakeDir(
			targetDMgr,
			targetDMgrLabel,
			ePrefix.XCpy(targetDMgrLabel))

		if err != nil {
			errs = append(errs, err)
			return dirCopyStats, errs
		}

		if dirCreated {
			dirCopyStats.DirCreated++
		}

		targetPathDoesExist = true
	}

	osPathSeparatorStr := string(os.PathSeparator)

	var src, target string
	var isMatch bool
	var fileInfos []FileInfoPlus
	var lenFileInfos int
	var errs2 []error

	fileInfos,
		lenFileInfos,
		errs2 = new(dirMgrHelperElectron).
		getFileInfosFromDirectory(
			sourceDMgr,
			sourceDMgrLabel,
			ePrefix.XCpy(sourceDMgrLabel))

	if len(errs2) != 0 {

		errs = append(errs, errs2...)

	}

	if lenFileInfos == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: dirMgrHelperElectron.getFileInfosFromDirectory()\n"+
			"returned a zero length array of File Info Objects from:\n"+
			"%v = %v\n",
			ePrefix.String(),
			sourceDMgrLabel,
			sourceDMgr.absolutePath)

		errs = append(errs, err)

		return dirCopyStats, errs
	}

	fh := new(FileHelper)
	var fileMode os.FileMode
	var doCopy bool

	for _, nameFileInfo := range fileInfos {

		if nameFileInfo.IsDir() {
			// We don't care about sub-directories
			continue

		}

		// This is a File. Proceed...
		dirCopyStats.TotalFilesProcessed++

		// This is not a directory. It is a file.
		// Determine if it matches the find file criteria.

		isMatch,
			err2,
			_ =
			fh.FilterFileName(
				nameFileInfo,
				fileSelectCriteria,
				ePrefix.XCpy("nameFileInfo"))

		if err2 != nil {

			err =
				fmt.Errorf("%v\n"+
					"Error returned by fh.FilterFileName(nameFileInfo, fileSelectCriteria).\n"+
					"%v directorySearched='%v'\n"+
					"fileName='%v'\n"+
					"Error= \n%v\n",
					ePrefix.String(),
					sourceDMgrLabel,
					sourceDMgr.absolutePath,
					nameFileInfo.Name(),
					err2.Error())

			errs = append(errs, err)

			continue
		}

		if !isMatch {
			// File DOES NOT Match Selection Criteria
			dirCopyStats.FilesNotCopied++

			dirCopyStats.FileBytesNotCopied += uint64(nameFileInfo.Size())

			continue

		} else {

			// We have a match

			// Create Directory if needed
			if !targetPathDoesExist {

				dirCreated,
					err = dMgrHlprMolecule.lowLevelMakeDir(
					targetDMgr,
					"targetDMgr",
					ePrefix)

				if err != nil {
					err2 = fmt.Errorf("%v\n"+
						"Error creating target directory!\n"+
						"%v Directory='%v'\n"+
						"Error= \n%v\n",
						ePrefix.String(),
						targetDMgrLabel,
						targetDMgr.absolutePath,
						err.Error())

					errs = append(errs, err2)

					continue
				}

				targetPathDoesExist = true

				if dirCreated {
					dirCopyStats.DirCreated++
				}
			}

			doCopy = false

			fileMode = nameFileInfo.Mode()

			if nameFileInfo.Mode().IsRegular() {

				doCopy = true

			} else if fileMode&os.ModeSymlink != 0 &&
				copySymLinkFiles {

				doCopy = true

			} else {
				// MUST BE Other Non-Regular File
				if copyOtherNonRegularFiles {
					doCopy = true
				}

			}

			if doCopy {

				src = sourceDMgr.absolutePath +
					osPathSeparatorStr + nameFileInfo.Name()

				target = targetDMgr.absolutePath +
					osPathSeparatorStr + nameFileInfo.Name()

				err = dMgrHlprMolecule.lowLevelCopyFile(
					src,
					nameFileInfo,
					target,
					"srcFile",
					"destinationFile",
					ePrefix)

				if err != nil {

					errs = append(errs, err)

					dirCopyStats.FilesNotCopied++

					dirCopyStats.FileBytesNotCopied += uint64(nameFileInfo.Size())

				} else {

					dirCopyStats.FilesCopied++

					dirCopyStats.FileBytesCopied += uint64(nameFileInfo.Size())
				}
			}
		}
	}

	return dirCopyStats, errs
}

// isDirMgrValid
//
// This method examines the current DirMgr object to
// determine whether it has been properly configured. If
// the current DirMgr object is valid, the method returns
// 'nil' for no errors.
//
// Otherwise, if the DirMgr object is INVALID, an error
// is returned.
//
// In order to qualify as a valid DirMgr instance the
// DirMgr.path and DirMgr.absolutePath member string
// variables must be populated with alphanumeric
// characters.
func (dMgrHlprPlanck *dirMgrHelperPlanck) isDirMgrValid(
	dMgr *DirMgr,
	dMgrLabel string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if dMgrHlprPlanck.lock == nil {
		dMgrHlprPlanck.lock = new(sync.Mutex)
	}

	dMgrHlprPlanck.lock.Lock()

	defer dMgrHlprPlanck.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"dirMgrHelperPlanck.isDirMgrValid()",
		"")

	if err != nil {
		return err
	}

	if len(dMgrLabel) == 0 {
		dMgrLabel = "dMgr"
	}

	if dMgr == nil {

		return fmt.Errorf("%v\n"+
			"Error: This Directory Manager instance is invalid!\n"+
			"Input parameter '%v' a 'nil' pointer.\n",
			ePrefix.String(),
			dMgrLabel)
	}

	if !dMgr.isInitialized {

		dMgr.absolutePath = ""
		dMgr.path = ""
		dMgr.doesAbsolutePathExist = false
		dMgr.doesPathExist = false
		dMgr.actualDirFileInfo = FileInfoPlus{}

		return fmt.Errorf("%v\n"+
			"Error: This Directory Manager instance is invalid!\n"+
			"Input parameter '%v' has NOT been properly initialized.\n",
			ePrefix.String(),
			dMgrLabel)

	}

	fh := new(FileHelper)

	errCode := 0

	errCode,
		_,
		dMgr.path =
		fh.IsStringEmptyOrBlank(dMgr.path)

	if errCode == -1 {

		dMgr.isInitialized = false
		dMgr.absolutePath = ""
		dMgr.path = ""
		dMgr.doesAbsolutePathExist = false
		dMgr.doesPathExist = false
		dMgr.actualDirFileInfo = FileInfoPlus{}

		return fmt.Errorf("%v\n"+
			"Error: This Directory Manager instance is invalid!\n"+
			"Input parameter '%v'.path is an empty string!\n",
			ePrefix.String(),
			dMgrLabel)
	}

	if errCode == -2 {

		dMgr.isInitialized = false
		dMgr.absolutePath = ""
		dMgr.path = ""
		dMgr.doesAbsolutePathExist = false
		dMgr.doesPathExist = false
		dMgr.actualDirFileInfo = FileInfoPlus{}

		return fmt.Errorf("%v\n"+
			"Error: This Directory Manager instance is invalid!\n"+
			"Input parameter '%v'.path consists of blank spaces.\n",
			ePrefix.String(),
			dMgrLabel)
	}

	errCode,
		_,
		dMgr.absolutePath =
		fh.IsStringEmptyOrBlank(dMgr.absolutePath)

	if errCode == -1 {

		dMgr.isInitialized = false
		dMgr.absolutePath = ""
		dMgr.absolutePath = ""
		dMgr.doesAbsolutePathExist = false
		dMgr.doesPathExist = false
		dMgr.actualDirFileInfo = FileInfoPlus{}

		return fmt.Errorf("%v\n"+
			"Error: This Directory Manager instance is invalid!\n"+
			"Input parameter '%v'.absolutePath is an empty string.\n",
			ePrefix.String(),
			dMgrLabel)
	}

	if errCode == -2 {

		dMgr.isInitialized = false
		dMgr.absolutePath = ""
		dMgr.absolutePath = ""
		dMgr.doesAbsolutePathExist = false
		dMgr.doesPathExist = false
		dMgr.actualDirFileInfo = FileInfoPlus{}

		return fmt.Errorf("%v\n"+
			"Error: This Directory Manager instance is invalid!\n"+
			"Input parameter '%v'.absolutePath consists of blank spaces.\n",
			ePrefix.String(),
			dMgrLabel)
	}

	return err
}
