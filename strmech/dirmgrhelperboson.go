package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"os"
	"sync"
)

type dirMgrHelperBoson struct {
	lock *sync.Mutex
}

// copyDirMgrs
//
// Receives a pointer to an incoming DirMgr object
// ('sourceDMgrIn') as an input parameter and copies the
// data values from the incoming object to the input
// parameter, 'destinationDMgr'.
//
// When the copy operation is completed, the
// 'destinationDMgr' object is configured as a duplicate
// of the incoming DirMgr object ('sourceDMgrIn').
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	If input parameter 'sourceDMgrIn' is invalid, this
//	method will return an error.
//
//	If input parameter 'sourceDMgrIn' does NOT exist on
//	disk, no error will be returned and the copy operation
//	will proceed to completion.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	destinationDMgr				*DirMgr
//
//		A pointer to an instance of DirMgr. All the data
//		values contained in input parameter
//		'sourceDMgrIn' will be copied to the
//		corresponding member data values contained in
//		'destinationDMgr'.
//
//		When the copy operation is completed, all data
//		values in 'destinationDMgr' will duplicate
//		corresponding data values contained in
//		'sourceDMgrIn'.
//
//	sourceDMgrIn				*DirMgr
//
//		A pointer to an instance of DirMgr. Data values
//		contained in this instance will be copied to
//		corresponding member data values encapsulated
//		by input parameter 'destinationDMgr'.
//
//		When the copy operation is completed, all data
//		values in 'destinationDMgr' will duplicate
//		corresponding data values contained in
//		'sourceDMgrIn'.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	--- NONE ---
func (dMgrHlprBoson *dirMgrHelperBoson) copyDirMgrs(
	destinationDMgr *DirMgr,
	sourceDMgrIn *DirMgr,
	errPrefDto *ePref.ErrPrefixDto) error {

	if dMgrHlprBoson.lock == nil {
		dMgrHlprBoson.lock = new(sync.Mutex)
	}

	dMgrHlprBoson.lock.Lock()

	defer dMgrHlprBoson.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "dirMgrHelperBoson." +
		"copyDirMgrs()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return err
	}

	if destinationDMgr == nil {

		err = fmt.Errorf("%v \n"+
			"ERROR: Input paramter 'destinationDMgr' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if sourceDMgrIn == nil {

		err = fmt.Errorf("%v \n"+
			"ERROR: Input paramter 'sourceDMgrIn' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	err = new(dirMgrHelperPlanck).isDirMgrValid(
		sourceDMgrIn,
		"sourceDMgrIn",
		ePrefix.XCpy("sourceDMgrIn"))

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Error: Input paramter 'sourceDMgrIn' is INVALID!\n"+
			"Error= \n%v\n",
			funcName,
			err.Error())
	}

	destinationDMgr.isInitialized = sourceDMgrIn.isInitialized
	destinationDMgr.originalPath = sourceDMgrIn.originalPath
	destinationDMgr.path = sourceDMgrIn.path
	destinationDMgr.isPathPopulated = sourceDMgrIn.isPathPopulated
	destinationDMgr.doesPathExist = sourceDMgrIn.doesPathExist
	destinationDMgr.parentPath = sourceDMgrIn.parentPath
	destinationDMgr.isParentPathPopulated = sourceDMgrIn.isParentPathPopulated
	destinationDMgr.absolutePath = sourceDMgrIn.absolutePath
	destinationDMgr.isAbsolutePathPopulated = sourceDMgrIn.isAbsolutePathPopulated
	destinationDMgr.doesAbsolutePathExist = sourceDMgrIn.doesAbsolutePathExist
	destinationDMgr.isAbsolutePathDifferentFromPath = sourceDMgrIn.isAbsolutePathDifferentFromPath
	destinationDMgr.directoryName = sourceDMgrIn.directoryName
	destinationDMgr.volumeName = sourceDMgrIn.volumeName
	destinationDMgr.isVolumePopulated = sourceDMgrIn.isVolumePopulated
	destinationDMgr.actualDirFileInfo =
		sourceDMgrIn.actualDirFileInfo.CopyOut()

	return err
}

// getSubDirsFilesInDir
//
// This method scans a single target directory path
// designated by input parameter 'targetDMgr' and returns
// information on selected subdirectories and selected
// files. The selected subdirectories are returned
// through a Directory Manager Collection passed as input
// parameter 'subDirsInDir'. The selected files are
// returned through a File Manager Collection passed as
// input parameter 'filesInDir'.
//
// To qualify for selection and inclusion in the returned
// Directory and File Manager Collections, items residing
// in the 'targetDMgr' target directory are divided into
// two classes, subdirectories and files. Subdirectories
// are standard directory entries. Files are defined as
// all artifacts residing in the target directory which
// are not subdirectories.
//
// To qualify as a selected subdirectory, the
// subdirectory must satisfy two filters. First, input
// parameter 'getSubdirectories' must be set to
// 'true'. Second, the subdirectory must satisfy the
// Directory Characteristics Selection Criteria specified
// by input parameter, 'subDirSelectCharacteristics'. If
// both of these filter requirements are satisfied, the
// subdirectory will be added to, and returned by, the
// 'subDirsInDir' Directory Manager Collection. Be
// advised that users control the behavior for current
// directories (".") and parent directories ("..") with
// input parameters 'includeSubDirCurrentDirOneDot' and
// 'includeSubDirParentDirTwoDots'.
//
// To qualify as a selected file, the file entry must
// also comply with two filters: File Type and File
// Characteristics Selection Criteria. Remember that
// files are defined as all artifacts residing in the
// target directory which are not subdirectories.
//
// To be eligible for file selection, the file entry must
// first comply with the specified File Type criteria. In
// terms of File Type, files are classified as
// regular files, SymLink files or other non-regular files.
//
// Screening criteria for File Type is controlled by the
// following three input parameters:
//
//	getRegularFiles - bool
//	getSymLinksFiles - bool
//	getOtherNonRegularFiles - bool
//
// File Types eligible for selection include Regular
// Files such as text files, image files and executable
// files, SymLink files and other Non-Regular Files such
// as device files, named pipes and sockets.
//
// In addition to File Type, selected files,
// must also comply with the File Characteristics
// Selection Criteria specified by input parameter,
// 'fileSelectCharacteristics'.
//
// The File Characteristics Selection criteria allows
// users to screen files for File Name, File Modification
// Date and File Mode. File Name selections can be based
// on pattern matches or regular expression matches.
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
//	(1)	This method will select and return information on
//		subdirectories and files in the target directory
//	 	specified by input parameter 'targetDMgr'. No
//	 	subdirectories will be searched for eligible
//	 	files. Only the top level or parent directory
//	 	identified by 'targetDMgr' will be searched for
//	 	eligible subdirectories and files.
//
//	(2) Selected subdirectories are required to fulfill
//		two requirements.
//
//		First, to select subdirectories, input parameter
//		'getSubdirectories' must be set to 'true'.
//
//		Second, subdirectories must conform to the
//		Directory Characteristics Selection Criteria
//		specified by input parameter
//		'subDirSelectCharacteristics'. This subdirectory
//		selection criteria allows users to screen
//	 	subdirectories for Name, Modification Date and
//	 	File Mode. Subdirectory Name selections can be
//	 	configured for pattern matches or regular
//	 	expression matches.
//
//	(3)	Selected files are required to match two sets
//		of selection criteria, File Type and File
//		Characteristics	Selection Criteria.
//
//		File Type Selection Criteria specifications are
//		passed as input parameters 'getRegularFiles',
//		'getSymLinksFiles' and 'getOtherNonRegularFiles'.
//		For an explanation of Regular and Non-Regular
//		files, see the section on "Definition of Terms",
//		above.
//
//		File Characteristics Selection Criteria are user
//		specified selection requirements passed as input
//		parameter 'fileSelectCriteria'. This file
//		selection criteria allows users to screen files
//		for File Name, File Modification Date and File
//		Mode. File Name selections can be based on
//		pattern matches or regular expression matches.
//
//	(4) If the target directory identified by input
//		parameter 'targetDMgr' contains NO subdirectories
//		or files matching the Type and Characteristics
//		selection criteria, this method will exit, no
//		subdirectories or files will be added to the
//		Directory Manager or File Manager Collections,
//		and no error will be returned.
//
//	(5) If the target directory identified by input
//		parameter 'targetDMgr' contains NO Files
//		whatsoever (0 Files), this method will exit, no
//		subdirectories or files will be added to the
//		Directory Manager or File Manager Collections,
//		and no error will be returned.
//
//	(6) If the target directory identified by input
//		parameter 'targetDMgr' does NOT exist on an
//		attached storage drive, an error will be
//		returned.
//
//	(7)	If input parameter 'getSubdirectories' is set to
//		'true' and input parameter 'subDirsInDir' is set
//		to 'nil', an error will be returned.
//
//		If input parameter 'getSubdirectories' is set to
//		'false', input parameter 'subDirsInDir' may be
//		safely configured as 'nil', and no error will
//		be returned.
//
//	(8)	If any file type input parameters (
//		'getRegularFiles', 'getSymLinksFiles' or
//		'getOtherNonRegularFiles') are set to
//		'true', and input parameter 'filesInDir' is set
//		to 'nil', an error will be returned.
//
//		If all file type input parameters (
//		'getRegularFiles', 'getSymLinksFiles' or
//		'getOtherNonRegularFiles') are set to
//		'false', input parameter 'filesInDir' may be
//		safely configured as 'nil', and no error will be
//		returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	targetDMgr					*DirMgr
//
//		An instance of DirMgr which identifies the
//		target directory where the search for
//		subdirectories and/or files will be conducted.
//		Information on subdirectories and files in this
//		directory matching the required selection
//		criteria will be returned in Directory Manager
//		and File Manager Collections.
//
//		If this target directory path does not currently
//		exist on an attached storage drive, an error will
//		be returned.
//
//	getSubdirectories			bool
//
//		If this parameter is set to 'true', directory
//		entries which also meet the Directory Selection
//		Characteristics criteria (subDirSelectCharacteristics),
//		will be stored and returned in the Directory Manager
//		Collection passed as input parameter 'subDirsInDir'.
//
//		If input parameters 'getSubdirectories',
//		'getRegularFiles', 'getSymLinksFiles' and
//		'getOtherNonRegularFiles' are all set to 'false',
//		these parameters will be classified as conflicted
//		and an error will be returned.
//
//	includeSubDirCurrentDirOneDot		bool
//
//		This parameter is only used, if input parameter
//		'getSubdirectories' is set to 'true'.
//
//		All directories include an os.FileInfo entry for
//		the current directory. The current directory name
//		is always denoted as single dot ('.').
//
//		When this parameter, 'includeSubDirCurrentDirOneDot',
//		is set to 'true' and input parameter
//		getSubdirectories' is set to 'true', the current
//		directory, designated as a single dot ('.'), will be
//		added to the Directory Manager Collection passed as
//		input parameter 'subDirsInDir'.
//
//	includeSubDirParentDirTwoDots 		bool
//
//		This parameter is only used, if input parameter
//		'getSubdirectories' is set to 'true'.
//
//		All directories include an os.FileInfo entry for
//		the parent directory. The parent directory name
//		is always denoted as two dots ('..').
//
//		When this parameter, 'includeSubDirParentDirTwoDots',
//		is set to 'true' and input parameter
//		getSubdirectories' is set to 'true', the parent
//		directory, designated as two dots ('..'), will be
//		added to the Directory Manager Collection passed as
//		input parameter 'subDirsInDir'.
//
//	getRegularFiles				bool
//
//		If this parameter is set to 'true', Regular
//		Files, which also meet the File Characteristics
//		selection criteria ('fileSelectCriteria'), will
//		be included in the file information returned
//		through the File Manager Collection passed as
//		input parameter 'filesInDir'.
//
//		Regular Files include text files, image files and
//		executable files.
//
//		For an explanation of Regular and Non-Regular
//		files, see the section on "Definition Of Terms",
//		above.
//
//		If input parameters 'getSubdirectories',
//		'getRegularFiles', 'getSymLinksFiles' and
//		'getOtherNonRegularFiles' are all set to 'false',
//		these parameters will be classified as conflicted
//		and an error will be returned.
//
//	getSymLinksFiles			bool
//
//		If this parameter is set to 'true', SymLink Files
//		which also meet the File Characteristics selection
//		criteria ('fileSelectCriteria'), will be included
//		in the file information returned through the File
//		Manager Collection passed as input parameter
//		'filesInDir'.
//
//		For an explanation of Regular and Non-Regular
//		files, see the section on "Definition Of Terms",
//		above.
//
//		If input parameters 'getSubdirectories',
//		'getRegularFiles', 'getSymLinksFiles' and
//		'getOtherNonRegularFiles' are all set to 'false',
//		these parameters will be classified as conflicted
//		and an error will be returned.
//
//	getOtherNonRegularFiles		bool
//
//		If this parameter is set to 'true', Other
//		Non-Regular Files, which also meet the File
//		Characteristics selection criteria
//		('fileSelectCriteria'), will be included in the
//		file information returned through the File
//		Manager Collection passed as input parameter
//		'filesInDir'.
//
//		Examples of other non-regular file types
//		include device files, named pipes, and sockets.
//		See the Definition Of Terms section above.
//
//		If input parameters 'getSubdirectories',
//		'getRegularFiles', 'getSymLinksFiles' and
//		'getOtherNonRegularFiles' are all set to 'false',
//		these parameters will be classified as conflicted
//		and an error will be returned.
//
//	subDirSelectCharacteristics FileSelectionCriteria
//
//		In addition to input parameter
//		'getSubdirectories' being set to 'true', selected
//		subdirectories must conform to the Subdirectory
//		Characteristics Selection Criteria specified by
//		this parameter, 'subDirSelectCharacteristics'.
//
//		This subdirectory selection criteria allows users
//		to screen subdirectories for Name, Modification
//		Date and File Mode. Subdirectory Name selections
//		can be configured for pattern matches or regular
//		expression matches.
//
//		Directory os.FileIno entries matching this
//		selection criteria, and the filter specified by
//		input parameter 'getSubdirectories', will be
//		included in the Directory Manager Collection
//		returned by input parameter 'subDirsInDir'.
//
//		Remember that setting 'subDirSelectCharacteristics'
//		to an empty instance of FileSelectionCriteria will
//		ensure that all subdirectories are selected.
//
//			Example:
//			subDirSelectCharacteristics =
//				FileSelectionCriteria{}
//
//			This ensures that all subdirectories will
//			satisfy the Subdirectory Characteristics
//			Selection Criteria.
//
//		For a detailed explanation of the Subdirectory
//		Characteristics Criteria specifications offered
//		by Type FileSelectionCriteria, see the
//		documentation for 'fileSelectCriteria', below.
//
//	fileSelectCriteria			FileSelectionCriteria
//
//		In addition to the File Type Selection Criteria,
//		selected files must conform to the File
//		Characteristics Selection Criteria specified by
//		this parameter, 'fileSelectCriteria'.
//
//		File Characteristics Selection Criteria allow
//		users to screen files for File Name, File
//		Modification Date and File Mode. File Name
//		selections can be configured for pattern matches
//		or regular expression matches.
//
//		Files matching these selection criteria, and the
//		File Type filter, will be included in the file
//		information returned through the File Manager
//		Collection passed as input parameter
//		'filesInDir'.
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
//		The FileSelectionCriteria Type allows for
//		configuration of single or multiple file selection
//		criterion. The 'SelectCriterionMode' can be used to
//		specify whether the file must match all, or any one,
//		of the active file selection criterion.
//
//		Elements of the File Characteristics Selection
//		Criteria are described below:
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
//		If all the file selection criterion in the FileSelectionCriteria
//		object are 'Inactive' or 'Not Set' (set to their zero or default values),
//		then all the files meeting the File Type requirements in the directory
//		defined by 'targetDMgr' will be selected.
//
//			Example:
//			  fsc := FileSelectCriterionMode{}
//
//			  In this example, 'fsc' is NOT initialized. Therefore,
//			  all the selection criterion are 'Inactive'. Consequently,
//			  all the files meeting the File Type requirements in the
//			  directory defined by 'targetDMgr' will be selected.
//
//		------------------------------------------------------------------------
//
//	subDirsInDir				*DirMgrCollection
//
//		A pointer to an instance of DirMgrCollection
//		which encapsulates an array of Directory
//		Manager (DirMgr) objects.
//
//		Information on subdirectories in the directory
//		defined by input parameter 'targetDMgr' which
//		match the specified Type and Characteristics
//		Selection Criteria will be converted to Directory
//		Manager (DirMgr) objects and added to this
//		Directory Manager Collection (DirMgrCollection).
//
//		If input parameter 'getSubdirectories' is set to
//		'true' and this input parameter, 'subDirsInDir',
//		is set to 'nil', an error will be returned.
//
//		If input parameter 'getSubdirectories' is set to
//		'false', input parameter 'subDirsInDir' may be
//		safely configured as 'nil' and no error will
//		be returned.
//
//	filesInDir					*FileMgrCollection
//
//		A pointer to an instance of FileMgrCollection
//		which encapsulates an array of File Manager
//		(FileMgr) objects.
//
//		Information on files in the directory defined by
//		input parameter 'targetDMgr' which match the
//		specified File Type and File Characteristics
//		Selection Criteria will be converted to File
//		Manager (FileMgr) objects and added to this File
//		Manager	Collection (FileMgrCollection).
//
//		If any file type input parameters (
//		'getRegularFiles', 'getSymLinksFiles' or
//		'getOtherNonRegularFiles') are set to
//		'true', and this input parameter, 'filesInDir',
//		is set to 'nil', an error will be returned.
//
//		If all file type input parameters (
//		'getRegularFiles', 'getSymLinksFiles' or
//		'getOtherNonRegularFiles') are set to
//		'false', input parameter 'filesInDir' may be
//		safely configured as 'nil', and no error will be
//		returned.
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
//	numOfSubDirsLocated			int
//
//		If this method completes successfully, this
//		parameter will return number of subdirectories
//		selected for addition to the Directory Manager
//		Collection passed as input parameter 'subDirsInDir'.
//
//	numOfFilesLocated			int
//
//		If this method completes successfully, this
//		parameter will return number of files selected
//		for addition to the File Manager Collection
//		passed as input parameter 'filesInDir'.
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
func (dMgrHlprBoson *dirMgrHelperBoson) getSubDirsFilesInDir(
	targetDMgr *DirMgr,
	getSubdirectories bool,
	includeSubDirCurrentDirOneDot bool,
	includeSubDirParentDirTwoDots bool,
	getRegularFiles bool,
	getSymLinksFiles bool,
	getOtherNonRegularFiles bool,
	subDirSelectCharacteristics FileSelectionCriteria,
	fileSelectCriteria FileSelectionCriteria,
	subDirsInDir *DirMgrCollection,
	filesInDir *FileMgrCollection,
	targetDMgrLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	numOfSubDirsLocated int,
	numOfFilesLocated int,
	err error) {

	if dMgrHlprBoson.lock == nil {
		dMgrHlprBoson.lock = new(sync.Mutex)
	}

	dMgrHlprBoson.lock.Lock()

	defer dMgrHlprBoson.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "dirMgrHelperBoson." +
		"getSubDirsFilesInDir()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		return numOfSubDirsLocated, numOfFilesLocated, err
	}

	if len(targetDMgrLabel) == 0 {

		targetDMgrLabel = "targetDMgr"
	}

	_,
		_,
		err = new(dirMgrHelperPreon).
		validateDirMgr(
			targetDMgr,
			true, //pathMustExist
			targetDMgrLabel,
			ePrefix)

	if err != nil {

		return numOfSubDirsLocated, numOfFilesLocated, err
	}

	if getSubdirectories == false &&
		getRegularFiles == false &&
		getSymLinksFiles == false &&
		getOtherNonRegularFiles == false {

		err = fmt.Errorf("%v\n"+
			"Fatal Error: File Type filters are conflicted!\n"+
			"All of the File Type filters are set to 'false'\n"+
			"This gurantees that NO subdirectories or files\n"+
			"will be selected.\n"+
			"getSubdirectories == false"+
			"getRegularFiles == false\n"+
			"getSymLinksFiles == false\n"+
			"getOtherNonRegularFiles == false\n",
			ePrefix.String())

		return numOfSubDirsLocated, numOfFilesLocated, err
	}

	if subDirsInDir == nil &&
		getSubdirectories == true {

		err = fmt.Errorf("%v\n"+
			"Error: Subdirectories were requested, but\n"+
			"Directory Manager Collection 'subDirsInDir'\n"+
			"is a 'nil' pointer. 'subDirsInDir' is invalid!\n",
			ePrefix.String())

		return numOfSubDirsLocated, numOfFilesLocated, err
	}

	if filesInDir == nil {

		if getRegularFiles == true ||
			getSymLinksFiles == true ||
			getOtherNonRegularFiles == true {

			err = fmt.Errorf("%v\n"+
				"Error: Files were requested, but File Manager\n"+
				"Collection 'filesInDir is a 'nil' pointer.\n"+
				"'filesInDir' is invalid!\n"+
				"getRegularFiles= '%v'\n"+
				"getSymLinksFiles= '%v'\n"+
				"getOtherNonRegularFiles= '%v'\n",
				getRegularFiles,
				getSymLinksFiles,
				getOtherNonRegularFiles,
				ePrefix.String())

			return numOfSubDirsLocated, numOfFilesLocated, err

		}
	}

	osPathSepStr := string(os.PathSeparator)
	var fileInfos []FileInfoPlus
	var lenFileInfos int
	var errs2 []error
	var fatalErr error

	fileInfos,
		lenFileInfos,
		errs2,
		fatalErr = new(dirMgrHelperMolecule).
		lowLevelGetFileInfosFromDir(
			targetDMgr,
			getSubdirectories,             // getDirectoryFileInfos
			includeSubDirCurrentDirOneDot, // includeSubDirCurrentDirOneDot
			includeSubDirParentDirTwoDots, // includeSubDirParentDirTwoDots
			getRegularFiles,               // getRegularFileInfos
			getSymLinksFiles,              // getSymLinksFileInfos
			getOtherNonRegularFiles,       // getOtherNonRegularFileInfos
			subDirSelectCharacteristics,   // subdirectorySelectCharacteristics
			fileSelectCriteria,            // fileSelectCharacteristics
			targetDMgrLabel,
			ePrefix.XCpy(targetDMgrLabel))

	if len(errs2) > 0 {

		if fatalErr != nil {

			errs2 = append(
				errs2, fatalErr)
		}

		err = new(StrMech).ConsolidateErrors(errs2)

		return numOfSubDirsLocated, numOfFilesLocated, err
	}

	if fatalErr != nil {

		err = fmt.Errorf("%v"+
			"Error returned from lowLevelGetFileInfosFromDir()\n"+
			"Error= \n%v\n",
			funcName,
			fatalErr.Error())

		return numOfSubDirsLocated, numOfFilesLocated, err
	}

	if lenFileInfos == 0 {
		// No files matching select criteria were
		// found in the directory.
		return numOfSubDirsLocated, numOfFilesLocated, err
	}

	for i := 0; i < lenFileInfos; i++ {

		if fileInfos[i].IsDir() {

			fatalErr = subDirsInDir.AddFileInfo(
				targetDMgr.absolutePath,
				fileInfos[i],
				ePrefix.XCpy("Dir targetDMgr:fileInfos[i]"))

			if fatalErr != nil {

				err = fmt.Errorf("%v\n"+
					"Error occurred while adding Directory\n"+
					"to Directory Manager Collection!\n"+
					"subDirsInDir.AddFileInfo()\n"+
					"targetDMgr Path= '%v'\n"+
					"SubDirectory Name= '%v'\n"+
					"SubDirectory Path= '%v'\n"+
					"Index= fileInfos[%v]\n"+
					"Error=\n%v\n",
					funcName,
					targetDMgr.absolutePath,
					fileInfos[i].Name(),
					targetDMgr.absolutePath+
						osPathSepStr+
						fileInfos[i].Name(),
					i,
					fatalErr.Error())

				return numOfSubDirsLocated, numOfFilesLocated, err
			}

			numOfSubDirsLocated++

			continue
		}

		// Must be a file
		fatalErr = filesInDir.AddFileMgrByFileInfo(
			targetDMgr.absolutePath,
			fileInfos[i],
			ePrefix.XCpy("File targetDMgr: fileInfos[i]"))

		if fatalErr != nil {

			err = fmt.Errorf("%v\n"+
				"Error occurred while adding os.FileInfo object\n"+
				"to the File Manager Collection!\n"+
				"filesInDir.AddFileMgrByFileInfo()\n"+
				"Target Directory '%v' = '%v'"+
				"fileInfos[%v].Name = '%v'\n"+
				"File Name = '%v'\n"+
				"Error = \n%v\n",
				funcName,
				targetDMgrLabel,
				targetDMgr.absolutePath,
				i,
				fileInfos[i].Name(),
				targetDMgr.absolutePath+
					osPathSepStr+
					fileInfos[i].Name(),
				fatalErr.Error())

			return numOfSubDirsLocated, numOfFilesLocated, err

		}

		numOfFilesLocated++
	}

	return numOfSubDirsLocated, numOfFilesLocated, err
}
