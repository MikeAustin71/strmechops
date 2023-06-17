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
// selected files from the directory identified by DirMgr
// object (sourceDMgr) to a target directory
// (targetDMgr).
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
//	(1)	This method ONLY copies files from the directory
//		identified by 'sourceDMgr' to the directory identified
//		by 'targetDMgr'. It does NOT copy files from
//		subdirectories of 'sourceDMgr'.
//
//	(2)	If the target directory does not exist, this method
//		will attempt to create it.
//
//	(3)	Files will only be copied if they meet the File Type
//		criteria and the File Characteristics Criteria.
//
//		File Type criteria are specified by input parameters:
//
//			copyRegularFiles bool
//			copySymLinkFiles bool
//			copyOtherNonRegularFiles bool
//
//		File Characteristics Selection criteria is specified by
//		input parameter 'fileSelectCriteria'.
//
//	(4) If input parameter 'returnCopiedFilesList' is set
//		to 'false', input parameter ('copiedFiles') can be
//		set to nil.
//
//	(5)	If input parameter 'returnSubDirsList' is set to
//		'false', input parameter ('subDirectories') can
//		be set to nil.
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
//	returnCopiedFilesList		bool
//
//		If input parameter 'returnCopiedFilesList' is set
//		to 'true', this method will return a populated
//		File Manager Collection documenting all the files
//		actually included in the copy operation.
//
//		If input parameter 'returnCopiedFilesList' is set
//		to 'false', this method will return an empty and
//		unpopulated instance of FileMgrCollection. This
//		means that the files actually copied by this
//		method will NOT be documented.
//
//	returnSubDirsList			bool
//
//		If input parameter 'returnSubDirsList' is set
//		to 'true', this method will create DirMgr objects
//		for each subdirectory in the parent directory,
//		and add them to Subdirectory collection passed as
//		input parameter 'subDirectories'.
//
//		If input parameter 'returnCopiedFilesList' is set
//		to 'false', no subdirectories will be added to
//		Directory Manager Collection (DirMgrCollection)
//		passed as input parameter 'subDirectories'.
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
//	copySymLinkFiles			bool
//
//		If this parameter is set to 'true', SymLink Files
//		which also meet the File Selection Characteristics
//		criteria (fileSelectCriteria), will be included
//		in the copy operation.
//
//		For an explanation of Regular and Non-Regular
//		files, see the section on "Definition Of Terms",
//		above.
//
//		If input parameters 'copyRegularFiles',
//		'copySymLinkFiles' and 'copyOtherNonRegularFiles'
//		are all set to 'false', an error will be returned.
//
//	copyOtherNonRegularFiles	bool
//
//		If this parameter is set to 'true', Other
//		Non-Regular Files, which also meet the File
//		Selection Characteristics criteria
//		(fileSelectCriteria), will be included in the
//		copy operation.
//
//		Examples of other non-regular file types
//		include device files, named pipes, and sockets.
//		See the Definition Of Terms section above.
//
//		If input parameters 'copyRegularFiles',
//		'copySymLinkFiles' and 'copyOtherNonRegularFiles'
//		are all set to 'false', an error will be returned.
//
//	fileSelectCriteria			FileSelectionCriteria
//
//		In addition to the File Type Selection Criteria,
//		selected files must conform to the File
//		Characteristics criteria specified by
//		'fileSelectCriteria'.
//
//		File Characteristics Selection criteria allow
//		users to screen files for File Name, File
//		Modification Date and File Mode.
//
//		Files matching these selection criteria, and the
//		File Type filter, will be included in the copy
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
//	subDirectories				*DirMgrCollection
//
//		A pointer to an instance of DirMgrCollection
//		which encapsulates an array of Directory Manager
//		(DirMgr) objects.
//
//		If input parameter 'returnSubDirsList' is set to
//		'true', all subdirectories residing the parent
//		directory defined by input parameter 'sourceDMgr'
//		will be added as new DirMgr objects to the
//		'subDirectories' Directory Manager Collection.
//
//		Directory entries for the current directory (".")
//		and the parent directory ("..") will be skipped
//		and will NOT be added to the 'subDirectories'
//		Directory Manager Collection.
//
//			type DirMgrCollection struct {
//				dirMgrs []DirMgr
//			}
//
//		Directory entries for the current directory (".")
//		and the parent directory ("..") will be skipped and
//		will NOT be added to the 'subDirectories' Directory
//		Manager Collection.
//
//		If input parameter 'returnSubDirsList' is set to
//		'false', no subdirectories will be added to this
//		Directory Manager Collection.
//
//		If input parameter 'returnSubDirsList' is set
//		to 'false', this parameter ('subDirectories') can be
//		set to nil.
//
//	copiedFiles					*FileMgrCollection
//
//		A pointer to an instance of FileMgrCollection
//		which encapsulates an array of File Manager
//		(FileMgr) objects.
//
//		If input parameter 'returnCopiedFilesList' is
//		set to 'true', all files actually copied to the
//		target directory defined by input parameter
//		'targetDMgr' will be added as new FileMgr objects
//		to the 'copiedFiles' File Manager Collection.
//		Effectively, this provides a list documenting the
//		files actually copied to the target directory.
//
//		If input parameter 'returnCopiedFilesList' is set
//		to 'false', no files will be added to this File
//		Manager collection.
//
//		If input parameter 'returnCopiedFilesList' is set
//		to 'false', this parameter ('copiedFiles') can be
//		set to nil.
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
//			DirsCreated          uint64
//			TotalFilesProcessed uint64
//			FilesCopied         uint64
//			FileBytesCopied     uint64
//			FilesNotCopied      uint64
//			FileBytesNotCopied  uint64
//			ComputeError        error
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
//		Non-fatal errors usually involves failures
//		associated with individual files.
//
//		The default behavior for Non-Fatal errors
//		accumulates these errors and returns them in an
//		array of errors. However, under the default
//		behavior, processing continues until a Fatal
//		Error is encountered or the method completes
//		processing and exits normally.
//
//		Any returned error messages will incorporate
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
//		Upon method completion, be sure to check both
//		Non-Fatal and Fatal errors.
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
func (dMgrHlprPlanck *dirMgrHelperPlanck) copyDirectoryFiles(
	sourceDMgr *DirMgr,
	targetDMgr *DirMgr,
	returnCopiedFilesList bool,
	returnSubDirsList bool,
	copyEmptyTargetDirectory bool,
	copyRegularFiles bool,
	copySymLinkFiles bool,
	copyOtherNonRegularFiles bool,
	fileSelectCriteria FileSelectionCriteria,
	sourceDMgrLabel string,
	targetDMgrLabel string,
	subDirectories *DirMgrCollection,
	copiedFiles *FileMgrCollection,
	errPrefDto *ePref.ErrPrefixDto) (
	dirCopyStats DirectoryCopyStats,
	nonfatalErrs []error,
	fatalErr error) {

	if dMgrHlprPlanck.lock == nil {
		dMgrHlprPlanck.lock = new(sync.Mutex)
	}

	dMgrHlprPlanck.lock.Lock()

	defer dMgrHlprPlanck.lock.Unlock()

	var err, err2 error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "dirMgrHelperPlanck." +
		"copyDirectoryFiles()"

	ePrefix,
		fatalErr = ePref.ErrPrefixDto{}.NewIEmpty(
		errPrefDto,
		funcName,
		"")

	if fatalErr != nil {

		return dirCopyStats, nonfatalErrs, fatalErr
	}

	if returnCopiedFilesList == true &&
		copiedFiles == nil {

		fatalErr = fmt.Errorf("%v\n"+
			"Error: Input parameters 'returnCopiedFilesList'\n"+
			"and 'copiedFiles' are conflicted.\n"+
			"'returnCopiedFilesList' is set to 'true'; however\n"+
			"the 'copiedFiles' pointer is 'nil'.\n"+
			"Provide a valid pointer to a 'copiedFiles'\n"+
			"File Manager Collection (FileMgrCollection)!\n",
			ePrefix.String())

		return dirCopyStats, nonfatalErrs, fatalErr
	}

	if returnSubDirsList == true &&
		subDirectories == nil {

		fatalErr = fmt.Errorf("%v\n"+
			"Error: Input parameters 'returnSubDirsList'\n"+
			"and 'subDirectories' are conflicted.\n"+
			"'returnSubDirsList' is set to 'true'; however\n"+
			"the 'subDirectories' pointer is 'nil'.\n"+
			"Provide a valid pointer to a 'subDirectories'\n"+
			"Directory Manager Collection (DirMgrCollection)!\n",
			ePrefix.String())

		return dirCopyStats, nonfatalErrs, fatalErr
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

		return dirCopyStats, nonfatalErrs, fatalErr
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
			false, //pathMustExist
			targetDMgrLabel,
			ePrefix)

	if fatalErr != nil {

		return dirCopyStats, nonfatalErrs, fatalErr
	}

	if copyRegularFiles == false &&
		copySymLinkFiles == false &&
		copyOtherNonRegularFiles == false {

		fatalErr = fmt.Errorf("%v\n"+
			"Fatal Error: File Type filters are conflicted!\n"+
			"All of the File Type filters are set to 'false'\n"+
			"This gurantees that NO files will be selected.\n"+
			"copyRegularFiles == false\n"+
			"copySymLinkFiles == false\n"+
			"copyOtherNonRegularFiles == false\n",
			ePrefix.String())

		return dirCopyStats, nonfatalErrs, fatalErr
	}

	isFileSelectionCriteriaActive :=
		fileSelectCriteria.IsSelectionCriteriaActive()

	dMgrHlprMolecule := new(dirMgrHelperMolecule)

	if !targetPathDoesExist && copyEmptyTargetDirectory {

		targetPathDoesExist,
			fatalErr = dMgrHlprMolecule.lowLevelMakeDir(
			targetDMgr,
			targetDMgrLabel,
			ePrefix.XCpy(targetDMgrLabel))

		if fatalErr != nil {

			return dirCopyStats, nonfatalErrs, fatalErr
		}

		if targetPathDoesExist {
			dirCopyStats.DirsCreated++
		}

		targetPathDoesExist = true
	}

	osPathSeparatorStr := string(os.PathSeparator)

	var src, target string
	var isMatch bool
	var fileInfos []FileInfoPlus
	var lenFileInfos int
	var errs2 []error

	// If returnSubDirsList is false
	// no subdirectory entries will be
	// returned.

	fileInfos,
		lenFileInfos,
		errs2,
		fatalErr = new(dirMgrHelperTachyon).
		getFileInfosFromDirectory(
			sourceDMgr,
			returnSubDirsList,        // getDirectoryFileInfos
			copyRegularFiles,         // getRegularFileInfos
			copySymLinkFiles,         // copySymLinkFiles,
			copyOtherNonRegularFiles, // copyOtherNonRegularFiles
			FileSelectionCriteria{},
			sourceDMgrLabel,
			ePrefix.XCpy(sourceDMgrLabel))

	if len(errs2) != 0 {

		nonfatalErrs = append(nonfatalErrs, errs2...)

	}

	if fatalErr != nil {

		return dirCopyStats, nonfatalErrs, fatalErr
	}

	if lenFileInfos == 0 {

		fatalErr = fmt.Errorf("%v\n"+
			"Error: The %v directory is EMPTY!\n"+
			"The copy files operation cannot proceed.\n"+
			"Method dirMgrHelperElectron.getFileInfosFromDirectory()\n"+
			"returned a zero length array of File Info Objects from:\n"+
			"%v = %v\n",
			ePrefix.String(),
			sourceDMgrLabel,
			sourceDMgrLabel,
			sourceDMgr.absolutePath)

		return dirCopyStats, nonfatalErrs, fatalErr
	}

	var fh = new(FileHelper)
	var fileMode os.FileMode
	var doCopy bool

	for _, nameFileInfo := range fileInfos {

		if returnSubDirsList &&
			nameFileInfo.IsDir() {
			// This is a Subdirectory!

			if nameFileInfo.Name() == "." ||
				nameFileInfo.Name() == ".." {

				continue
			}

			err2 = subDirectories.
				AddDirMgrByKnownPathDirName(
					sourceDMgr.absolutePath,
					nameFileInfo.Name(),
					ePrefix.XCpy("sourceDMgr+nameFileInfo"))

			if err2 != nil {

				fatalErr = fmt.Errorf("%v\n"+
					"Error returned adding subdirectory DirMgrCollection!\n"+
					"Parent Directory = '%v'\n"+
					"Subdirectory Name= '%v'\n"+
					"Full Subdirectory Path= '%v'\n"+
					"Error= \n%v\n",
					funcName,
					sourceDMgr.absolutePath,
					nameFileInfo.Name(),
					sourceDMgr.absolutePath+
						osPathSeparatorStr+
						nameFileInfo.Name(),
					err2.Error())

				return dirCopyStats, nonfatalErrs, fatalErr
			}

			dirCopyStats.SubDirsDocumented++

			continue
		}

		// This is a File. Proceed to Copy Operation!
		dirCopyStats.TotalFilesProcessed++

		// This is not a directory. It is a file.
		// Determine if it matches the find file criteria.

		if isFileSelectionCriteriaActive == true {

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
						"%v directory= '%v'\n"+
						"fileName= '%v'\n"+
						"Error= \n%v\n",
						funcName,
						sourceDMgrLabel,
						sourceDMgr.absolutePath,
						nameFileInfo.Name(),
						err2.Error())

				nonfatalErrs = append(nonfatalErrs, err)

				continue
			}

		} else {

			isMatch = true
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

				targetPathDoesExist,
					err = dMgrHlprMolecule.lowLevelMakeDir(
					targetDMgr,
					"targetDMgr",
					ePrefix)

				if err != nil {
					fatalErr = fmt.Errorf("%v\n"+
						"Error creating target directory!\n"+
						"%v Directory='%v'\n"+
						"Error= \n%v\n",
						funcName,
						targetDMgrLabel,
						targetDMgr.absolutePath,
						err.Error())

					return dirCopyStats, nonfatalErrs, fatalErr
				}

				if targetPathDoesExist == true {
					dirCopyStats.DirsCreated++
				}
			}

			doCopy = false

			fileMode = nameFileInfo.Mode()

			if fileMode.IsRegular() &&
				copyRegularFiles {

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

				err2 = dMgrHlprMolecule.lowLevelCopyFile(
					src,
					nameFileInfo,
					target,
					"srcFile",
					"destinationFile",
					ePrefix)

				if err2 != nil {

					fatalErr = fmt.Errorf("%v\n"+
						"Error: Attempted file copy FAILED!\n"+
						"Source File: %v\n"+
						"Target File: %v\n"+
						"Error= \n%v\n",
						funcName,
						src,
						target,
						err2.Error())

					dirCopyStats.FilesNotCopied++

					dirCopyStats.FileBytesNotCopied +=
						uint64(nameFileInfo.Size())

					return dirCopyStats, nonfatalErrs, fatalErr

				} else {

					// The file was successfully copied

					if returnCopiedFilesList {

						err2 = copiedFiles.AddFileMgrByPathFileNameExt(
							target,
							ePrefix.XCpy("target"))

						if err2 != nil {

							nonfatalErrs = append(
								nonfatalErrs,
								fmt.Errorf(
									"%v\n"+
										"Non-Fatal Error\n"+
										"Error from copiedFiles.AddFileMgrByPathFileNameExt()\n"+
										"%v Directory Path= %v\n"+
										"File Name = %v\n"+
										"File Mode = %v\n"+
										"File Path = %v\n"+
										"Error= \n%v\n",
									funcName,
									targetDMgrLabel,
									targetDMgr.absolutePath,
									nameFileInfo.Name(),
									nameFileInfo.Mode(),
									target,
									err2.Error()))
						}

					}

					dirCopyStats.FilesCopied++

					dirCopyStats.FileBytesCopied += uint64(nameFileInfo.Size())
				}

			} else {

				// The file was NOT copied

				dirCopyStats.FilesNotCopied++

				dirCopyStats.FileBytesNotCopied += uint64(nameFileInfo.Size())
			}
		}
	}

	return dirCopyStats, nonfatalErrs, fatalErr
}

// deleteDirectoryFiles
//
// This method deletes selected files in a single
// directory. No subdirectories will be deleted.
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
//	(1)	This method will delete files in the directory
//		specified by input parameter 'targetDMgr'. The
//		files to be deleted must match the File Selection
//		Characteristics Criteria passed as input parameter
//		'fileSelectCriteria'.
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
//		Characteristics criteria specified by
//		'fileSelectCriteria'.
//
//		File Characteristics Selection criteria allows
//		users to screen files for File Name, File
//		Modification Date and File Mode.
//
//		Files matching these selection criteria, and the
//		File Type filter, will be included in the file
//		deletion operation performed by this method.
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
//	deletedDirFileStats			DeleteDirFilesStats
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
func (dMgrHlprPlanck *dirMgrHelperPlanck) deleteDirectoryFiles(
	targetDMgr *DirMgr,
	returnDeletedFilesList bool,
	deleteRegularFiles bool,
	deleteSymLinkFiles bool,
	deleteOtherNonRegularFiles bool,
	fileSelectCriteria FileSelectionCriteria,
	targetDMgrLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	deletedDirFileStats DeleteDirFilesStats,
	deletedFiles FileMgrCollection,
	nonfatalErrs []error,
	fatalErr error) {

	if dMgrHlprPlanck.lock == nil {
		dMgrHlprPlanck.lock = new(sync.Mutex)
	}

	dMgrHlprPlanck.lock.Lock()

	defer dMgrHlprPlanck.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "dirMgrHelperPlanck." +
		"deleteDirectoryFiles()"

	ePrefix,
		fatalErr = ePref.ErrPrefixDto{}.NewIEmpty(
		errPrefDto,
		funcName,
		"")

	if fatalErr != nil {

		return deletedDirFileStats,
			deletedFiles,
			nonfatalErrs,
			fatalErr
	}

	if len(targetDMgrLabel) == 0 {

		targetDMgrLabel = "targetDMgr"
	}

	var fileInfos []FileInfoPlus
	var lenFileInfos int
	var errs2 []error

	if deleteRegularFiles == false &&
		deleteSymLinkFiles == false &&
		deleteOtherNonRegularFiles == false {

		fatalErr = fmt.Errorf("%v\n"+
			"Fatal Error: File Type filters are conflicted!\n"+
			"All of the File Type filters are set to 'false'\n"+
			"This gurantees that NO files will be selected.\n"+
			"deleteRegularFiles == false\n"+
			"deleteSymLinkFiles == false\n"+
			"deleteOtherNonRegularFiles == false\n",
			ePrefix.String())

		return deletedDirFileStats,
			deletedFiles,
			nonfatalErrs,
			fatalErr
	}

	fileInfos,
		lenFileInfos,
		errs2,
		fatalErr = new(dirMgrHelperTachyon).
		getFileInfosFromDirectory(
			targetDMgr,
			false,                      // getDirectoryFileInfos
			deleteRegularFiles,         // getRegularFileInfos
			deleteSymLinkFiles,         // getSymLinksFileInfos
			deleteOtherNonRegularFiles, // getOtherNonRegularFileInfos
			fileSelectCriteria,
			targetDMgrLabel,
			ePrefix.XCpy(targetDMgrLabel))

	if len(errs2) != 0 {

		nonfatalErrs = append(nonfatalErrs, errs2...)

	}

	if fatalErr != nil {

		return deletedDirFileStats,
			deletedFiles,
			nonfatalErrs,
			fatalErr
	}

	if lenFileInfos == 0 {

		return deletedDirFileStats,
			deletedFiles,
			nonfatalErrs,
			fatalErr
	}

	var err2 error
	osPathSepStr := string(os.PathSeparator)

	deletedDirFileStats.TotalDirsScanned = 1

	for _, nameFileInfo := range fileInfos {

		deletedDirFileStats.TotalFilesProcessed++

		// This is not a directory. It is a file.
		// The file matches the File Characteristics
		// Criteria and the File Type Criteria

		err2 = os.Remove(targetDMgr.absolutePath + osPathSepStr + nameFileInfo.Name())

		if err2 != nil {

			fatalErr = fmt.Errorf("%v\n"+
				"Error returned by os.Remove(pathFileName).\n"+
				"pathFileName='%v'\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				targetDMgr.absolutePath+osPathSepStr+nameFileInfo.Name(),
				err2.Error())

			return deletedDirFileStats,
				deletedFiles,
				nonfatalErrs,
				fatalErr
		}

		// File successfully deleted

		if returnDeletedFilesList {

			err2 = deletedFiles.AddFileMgrByFileInfo(
				targetDMgr.absolutePath,
				nameFileInfo,
				ePrefix.XCpy("targetDMgr+nameFileInfo"))

			if err2 != nil {

				nonfatalErrs = append(
					nonfatalErrs,
					fmt.Errorf(
						"%v\n"+
							"Non-Fatal Error\n"+
							"Error from deletedFiles.AddFileMgrByFileInfo()\n"+
							"%v Directory Path= %v\n"+
							"File Name = %v\n"+
							"File Mode = %v\n"+
							"File Path = %v\n"+
							"Error= \n%v\n",
						funcName,
						targetDMgrLabel,
						targetDMgr.absolutePath,
						nameFileInfo.Name(),
						nameFileInfo.Mode(),
						targetDMgr.absolutePath+osPathSepStr+nameFileInfo.Name(),
						err2.Error()))
			}
		}

		deletedDirFileStats.FilesDeleted++
		deletedDirFileStats.FilesDeletedBytes +=
			uint64(nameFileInfo.Size())

		deletedDirFileStats.NumOfDirsWhereFilesDeleted = 1
	}

	return deletedDirFileStats,
		deletedFiles,
		nonfatalErrs,
		fatalErr
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
