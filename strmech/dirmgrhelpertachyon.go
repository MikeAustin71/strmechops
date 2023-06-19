package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"os"
	"sync"
)

type dirMgrHelperTachyon struct {
	lock *sync.Mutex
}

// getDirectoryProfile
//
// This method returns an instance of DirectoryProfile which
// includes file breakdowns and statistics on the directory
// path specified by input parameter 'dMgr'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dMgr				*DirMgr
//
//		A pointer to an instance of DirMgr. The
//		directory path identified by this instance will
//		be analyzed for the following objectives:
//
//			(1)	Determine if the directory path exists
//				on a local storage drive.
//
//			(2) If the path does exist, statistics on
//				the directory will be generated and
//				returned via an instance of
//				DirectoryProfile.
//
//	dMgrLabel string
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
//	directoryPathDoesExist		bool
//
//		If this parameter returns a boolean value of 'true',
//		it signals that the directory path specified by input
//		parameter 'dMgr' actually exists on a storage drive.
//
//	dirProfile					DirectoryProfile
//
//		If this method completes successfully, this
//		returned instance of DirectoryProfile will be
//		populated with profile and statistical
//		information on the directory identified by input
//		parameter 'dMgr'.
//
//		type DirectoryProfile struct {
//
//			DirAbsolutePath string
//				The absolute directory path for the
//				directory described by this profile
//				information.
//
//			DirExistsOnStorageDrive bool
//				If 'true', this paramter signals
//				that the directory actually exists on
//				a storage drive.
//
//			DirTotalFiles uint64
//				The number of total files, of all types,
//				residing in the subject directory. This
//				includes directory entry files, Regular
//				Files, SymLink Files and Non-Regular
//				Files.
//
//			DirTotalFileBytes uint64
//				The size of all files, of all types,
//				residing in the subject directory
//				expressed in bytes. This includes
//				directory entry files, Regular Files,
//				SymLink Files and Non-Regular Files.
//
//			DirSubDirectories uint64
//				The number of subdirectories residing
//				within the subject directory. This
//
//			DirSubDirectoriesBytes uint64
//				The total size of all Subdirectory entries
//				residing in the subject directory expressed
//				in bytes.
//
//			DirRegularFiles uint64
//				The number of 'Regular' Files residing
//				within the subject Directory. Regular
//				files include text files, image files
//				and executable files. Reference:
//				https://www.computerhope.com/jargon/r/regular-file.htm
//
//			DirRegularFileBytes uint64
//				The total size of all 'Regular' files
//				residing in the subject directory expressed
//				in bytes.
//
//			DirSymLinkFiles uint64
//				The number of SymLink files residing in the
//				subject directory.
//
//			DirSymLinkFileBytes uint64
//				The total size of all SymLink files
//				residing in the subject directory
//				expressed in bytes.
//
//			DirNonRegularFiles uint64
//				The total number of Non-Regular files residing
//				in the subject directory.
//
//				Non-Regular files include directories, device
//				files, named pipes, sockets, and symbolic links.
//
//			DirNonRegularFileBytes uint64
//				The total size of all Non-Regular files residing
//				in the subject directory expressed in bytes.
//
//			ComputeError error
//				Computational or processing errors will be
//				recorded through this parameter.
//		}
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
func (dMgrHlprTachyon *dirMgrHelperTachyon) getDirectoryProfile(
	dMgr *DirMgr,
	dMgrLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	directoryPathDoesExist bool,
	dirProfile DirectoryProfile,
	err error) {

	if dMgrHlprTachyon.lock == nil {
		dMgrHlprTachyon.lock = new(sync.Mutex)
	}

	dMgrHlprTachyon.lock.Lock()

	defer dMgrHlprTachyon.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "dirMgrHelperTachyon." +
		"getDirectoryProfile()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		return directoryPathDoesExist, dirProfile, err
	}

	if len(dMgrLabel) == 0 {

		dMgrLabel = "dMgr"
	}

	_,
		directoryPathDoesExist,
		err = new(dirMgrHelperPreon).
		validateDirMgr(
			dMgr,
			false, // pathMustExist
			dMgrLabel,
			ePrefix)

	if err != nil {

		return directoryPathDoesExist, dirProfile, err
	}

	dirProfile.DirAbsolutePath =
		dMgr.absolutePath

	dirProfile.DirExistsOnStorageDrive =
		directoryPathDoesExist

	if !directoryPathDoesExist {

		return directoryPathDoesExist, dirProfile, err
	}

	var err2 error
	var nameDirEntries []os.DirEntry

	nameDirEntries,
		err2 = os.ReadDir(dMgr.absolutePath)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by os.ReadDir(%v.absolutePath).\n"+
			"%v.absolutePath='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			dMgrLabel,
			dMgrLabel,
			dMgr.absolutePath,
			err2.Error())

		return directoryPathDoesExist, dirProfile, err
	}

	var osFInfo os.FileInfo

	osPathSepStr := string(os.PathSeparator)

	for _, directoryEntry := range nameDirEntries {

		osFInfo,
			err2 = directoryEntry.Info()

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Error: Error Returned by directoryEntry.Info().\n"+
				"The conversion of DirEntry to os.FileInfo Failed."+
				"%v= '%v'\n"+
				"FileName= '%v'\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				dMgrLabel,
				dMgr.absolutePath,
				dMgr.absolutePath+osPathSepStr+osFInfo.Name(),
				err2.Error())

			return directoryPathDoesExist, dirProfile, err
		}

		if osFInfo.IsDir() {

			dirProfile.DirSubDirectories++
			dirProfile.DirSubDirectoriesBytes +=
				uint64(osFInfo.Size())

		} else if osFInfo.Mode().IsRegular() {

			dirProfile.DirTotalFiles++
			dirProfile.DirTotalFileBytes +=
				uint64(osFInfo.Size())

			dirProfile.DirRegularFiles++
			dirProfile.DirRegularFileBytes +=
				uint64(osFInfo.Size())

		} else if osFInfo.Mode()&os.ModeSymlink != 0 {

			dirProfile.DirTotalFiles++
			dirProfile.DirTotalFileBytes +=
				uint64(osFInfo.Size())

			dirProfile.DirSymLinkFiles++
			dirProfile.DirSymLinkFileBytes +=
				uint64(osFInfo.Size())

		} else {

			dirProfile.DirTotalFiles++
			dirProfile.DirTotalFileBytes +=
				uint64(osFInfo.Size())

			dirProfile.DirNonRegularFiles++
			dirProfile.DirNonRegularFileBytes +=
				uint64(osFInfo.Size())

		}

	}

	var checkTotalFiles uint64

	checkTotalFiles =
		dirProfile.DirRegularFiles +
			dirProfile.DirSymLinkFiles +
			dirProfile.DirNonRegularFiles

	if dirProfile.DirTotalFiles !=
		checkTotalFiles {

		err = fmt.Errorf("%v\n"+
			"Error: The Total Number of Files Processed"+
			"does NOT equal the sum of file type categories.\n"+
			"dirProfile.DirRegularFiles +\n"+
			"dirProfile.DirSymLinkFiles +\n"+
			"dirProfile.DirNonRegularFiles = %v\n"+
			"The Total Number of files Processed = %v\n",
			ePrefix.String(),
			checkTotalFiles,
			dirProfile.DirTotalFiles)

		dirProfile.ComputeError = fmt.Errorf("%v",
			err.Error())
	}

	return directoryPathDoesExist, dirProfile, err
}

// getFileInfosFromDirectory
//
// Receives an instance of DirMgr and proceeds to extract
// os.FileInfo data describing the files and directories
// contained in that DirMgr's absolute directory path.
//
// The os.FileInfo interface is defined as follows:
//
//	 	type FileInfo interface {
//			 Name() string
//				Base name of the file
//
//			 Size() int64
//			 	Length in bytes for regular files;
//			 	system-dependent for others
//
//			 Mode() FileMode
//			 	File mode bits
//
//			 ModTime() time.Time
//			 	Modification time
//
//			 IsDir() bool
//			 	Abbreviation for Mode().IsDir()
//
//			 Sys() interface{}
//			 	Underlying data source (can return nil)
//	 	}
//
// Upon completion, this method returns an array of
// FileInfoPlus objects containing os.FileInfo
// information on files residing in the directory path
// specified by input parameter 'dMgr'. Type FileInfoPlus
// implements the os.FileInfo interface.
//
// To qualify as a selected file, the file entry must
// comply with two filters: File Type and File
// Characteristics.
//
// To be eligible for selection, the file must first
// conform to the specified File Type criteria. In
// terms of File Type, files are classified as
// directories, regular files, SymLink files or other
// non-regular files. For an explanation of Regular and
// Non-Regular files, see the Definition of Terms section
// below.
//
// Screening criteria for File Type is controlled by the
// following four input parameters:
//
//	getDirectoryFileInfos - bool
//	getRegularFileInfos - bool
//	getSymLinksFileInfos - bool
//	getOtherNonRegularFiles - bool
//
// File Types eligible for this selection operation
// include Directories, Regular Files such as text files,
// image files and/ executable files, SymLink files and
// other Non-Regular Files such as device files, named
// pipes and sockets.
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
// # Input Parameters
//
//	dMgr							*DirMgr
//
//		A pointer to an instance of DirMgr. This instance
//		specifies the absolute directory path which will be
//		searched to extract and return os.FileInfo
//		information on all files and directories contained
//		therein.
//
//		If the directory specified by 'dMgr' does not
//		exist, an error will be returned.
//
//	getDirectoryFileInfos			bool
//
//		If this parameter is set to 'true', directory
//		entries which also meet the File Selection
//		Characteristics criteria (fileSelectCriteria),
//		will be included in the os.FileInfo information
//		('fileInfos') returned by this method.
//
//		If input parameters 'getDirectoryFileInfos',
//		'getRegularFileInfos', 'getSymLinksFileInfos' and
//		'getOtherNonRegularFileInfos' are all set to
//		'false', an error will be returned.
//
//	getRegularFileInfos				bool
//
//		If this parameter is set to 'true', regular files,
//		which also meet the File Selection
//		Characteristics criteria (fileSelectCriteria),
//		will be included in the os.FileInfo information
//		('fileInfos') returned by this method.
//
//		Regular Files include text files, image files and
//		executable files.
//
//		For an explanation of Regular and Non-Regular
//		files, see the section on "Definition Of Terms",
//		above.
//
//		If input parameters 'getDirectoryFileInfos',
//		'getRegularFileInfos', 'getSymLinksFileInfos' and
//		'getOtherNonRegularFileInfos' are all set to
//		'false', an error will be returned.
//
//	getSymLinksFileInfos			bool
//
//		If this parameter is set to 'true', SymLink files,
//		which also meet the File Selection
//		Characteristics criteria (fileSelectCriteria),
//		will be included in the os.FileInfo information
//		('fileInfos') returned by this method.
//
//		For an explanation of Regular and Non-Regular
//		files, see the section on "Definition Of Terms",
//		above.
//
//		If input parameters 'getDirectoryFileInfos',
//		'getRegularFileInfos', 'getSymLinksFileInfos' and
//		'getOtherNonRegularFileInfos' are all set to
//		'false', an error will be returned.
//
//	getOtherNonRegularFileInfos		bool
//
//		If this parameter is set to 'true', Other
//		Non-Regular files, which also meet the File
//		Selection Characteristics criteria
//		(fileSelectCriteria), will be included in the
//		os.FileInfo information ('fileInfos') returned by
//		this method.
//
//		Other Non-regular files include	device files,
//		named pipes, and sockets. For an explanation of
//		Regular and Non-Regular files, see the section on
//		"Definition Of Terms", above.
//
//		If input parameters 'getDirectoryFileInfos',
//		'getRegularFileInfos', 'getSymLinksFileInfos' and
//		'getOtherNonRegularFileInfos' are all set to
//		'false', an error will be returned.
//
//	fileSelectCriteria				FileSelectionCriteria
//
//		In addition to the File Type Selection Criteria,
//		selected files must conform to the File
//		Characteristics specified by 'fileSelectCriteria'.
//
//		Files matching these selection criteria, and the
//		File Type filter, will be included in the array of
//		FileInfoPlus objects containing os.FileInfo
//		information returned by this method.
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
//		If all of the file selection criterion in the FileSelectionCriteria
//		object are 'Inactive' or 'Not Set' (set to their zero or default values),
//		then all the files processed in the directory tree will be selected.
//
//			Example:
//			  fsc := FileSelectCriterionMode{}
//
//		 	In this example, 'fsc' is NOT initialized. Therefore,
//			all the selection criterion are 'Inactive'. Consequently,
//			all the files encountered in the target directory during
//			the search operation will meet the file characteristics
//			selection criteria.
//
//		------------------------------------------------------------------------
//
//	dMgrLabel						string
//
//		The name or label associated with input parameter
//		'dMgr', which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "dMgr" will be
//		automatically applied.
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
//	fileInfos						[]FileInfoPlus
//
//		If this method completes successfully, this
//		method will return an array of FileInfoPlus
//		objects containing os.FileInfo data on the files
//		contained in the directory path specified by
//		input parameter 'dMgr'.
//
//		The types of file and directory entries included
//		will be controlled by the following input
//		parameters:
//
//			excludeDirectoryFileInfos
//			excludeSymLinks
//			excludeOtherNonRegularFiles
//
//
//		Type FileInfoPlus implements the os.FileInfo
//		interface, but provides additional file information
//		over and above that provided by the standard
//		os.FileInfo interface.
//
//		The os.FileInfo interface is defined as follows:
//
//	 	type FileInfo interface {
//			 Name() string
//				Base name of the file
//
//			 Size() int64
//			 	Length in bytes for regular files;
//			 	system-dependent for others
//
//			 Mode() FileMode
//			 	File mode bits
//
//			 ModTime() time.Time
//			 	Modification time
//
//			 IsDir() bool
//			 	Abbreviation for Mode().IsDir()
//
//			 Sys() interface{}
//			 	Underlying data source (can return nil)
//	 	}
//
//	nonfatalErrs					[]error
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
//		Non-fatal errors usually involve failures
//		associated with reading individual files.
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
//	fatalErr						error
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
func (dMgrHlprTachyon *dirMgrHelperTachyon) getFileInfosFromDirectory(
	dMgr *DirMgr,
	getDirectoryFileInfos bool,
	getRegularFileInfos bool,
	getSymLinksFileInfos bool,
	getOtherNonRegularFileInfos bool,
	fileSelectCriteria FileSelectionCriteria,
	dMgrLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	fileInfos []FileInfoPlus,
	lenFileInfos int,
	nonfatalErrs []error,
	fatalErr error) {

	if dMgrHlprTachyon.lock == nil {
		dMgrHlprTachyon.lock = new(sync.Mutex)
	}

	dMgrHlprTachyon.lock.Lock()

	defer dMgrHlprTachyon.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	funcName := "dirMgrHelperTachyon." +
		"getFileInfosFromDirectory()"

	ePrefix,
		fatalErr = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		return fileInfos, lenFileInfos, nonfatalErrs, fatalErr
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

		return fileInfos, lenFileInfos, nonfatalErrs, fatalErr
	}

	if getDirectoryFileInfos == false &&
		getRegularFileInfos == false &&
		getSymLinksFileInfos == false &&
		getOtherNonRegularFileInfos == false {

		fatalErr = fmt.Errorf("%v\n"+
			"Fatal Error: File Type filters are conflicted!\n"+
			"All of the File Type filters are set to 'false'\n"+
			"This gurantees that NO files will be selected.\n"+
			"getDirectoryFileInfos == false\n"+
			"getRegularFileInfos == false\n"+
			"getSymLinksFileInfos == false\n"+
			"getOtherNonRegularFileInfos == false\n",
			ePrefix.String())

		return fileInfos, lenFileInfos, nonfatalErrs, fatalErr
	}

	isFileSelectionCriteriaActive :=
		fileSelectCriteria.IsSelectionCriteriaActive()

	var err2 error
	var nameDirEntries []os.DirEntry

	nameDirEntries,
		err2 = os.ReadDir(dMgr.absolutePath)

	if err2 != nil {

		fatalErr = fmt.Errorf("%v\n"+
			"Error returned by os.ReadDir(%v.absolutePath).\n"+
			"%v.absolutePath='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			dMgrLabel,
			dMgrLabel,
			dMgr.absolutePath,
			err2.Error())

		return fileInfos, lenFileInfos, nonfatalErrs, fatalErr
	}

	var nameFileInfo os.FileInfo

	osPathSepStr := string(os.PathSeparator)

	fip := FileInfoPlus{}

	lenFileInfos = len(nameDirEntries)

	fileInfos = make([]FileInfoPlus, lenFileInfos)

	var osFInfo os.FileInfo

	var isFileTypeFilterMatch bool

	var isMatch bool

	var fh = new(FileHelper)

	for i := 0; i < lenFileInfos; i++ {

		isFileTypeFilterMatch = false

		osFInfo,
			err2 = nameDirEntries[i].Info()

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

			nonfatalErrs = append(nonfatalErrs, err)

			continue
		}

		if osFInfo.IsDir() && getDirectoryFileInfos {

			isFileTypeFilterMatch = true

		} else if osFInfo.Mode()&os.ModeSymlink != 0 &&
			getSymLinksFileInfos {

			isFileTypeFilterMatch = true

		} else if !osFInfo.Mode().IsRegular() &&
			getOtherNonRegularFileInfos {

			isFileTypeFilterMatch = true

		} else if osFInfo.Mode().IsRegular() &&
			getRegularFileInfos {

			isFileTypeFilterMatch = true

		} else {

			if getOtherNonRegularFileInfos == true {

				isFileTypeFilterMatch = true

			}

		}

		if !isFileTypeFilterMatch {
			// This file fails the File Type
			// filter test. Skip it.
			continue
		}

		// MUST BE: This file passes the File
		// Type filter test. Process it.

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
						"%v directorySearched='%v'\n"+
						"fileName='%v'\n"+
						"File Path='%v'\n"+
						"Error= \n%v\n",
						funcName,
						dMgrLabel,
						dMgr.absolutePath,
						nameFileInfo.Name(),
						dMgr.absolutePath+
							osPathSepStr+
							nameFileInfo.Name(),
						err2.Error())

				nonfatalErrs = append(nonfatalErrs, err)

				continue
			}

			if !isMatch {
				continue
			}
		}

		// MUST BE:
		// isFileSelectionCriteriaActive == FALSE
		// File Type Filter has already been satisfied.
		// This file is a hit. Save it to fileInfos.

		fileInfos = append(
			fileInfos, fip.NewFromFileInfo(osFInfo))
	}

	lenFileInfos = len(fileInfos)

	return fileInfos, lenFileInfos, nonfatalErrs, fatalErr
}
