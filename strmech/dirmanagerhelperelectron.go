package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"os"
	"strings"
	"sync"
	"time"
)

type dirMgrHelperElectron struct {
	lock *sync.Mutex
}

// empty
//
// Resets all internal member variables for the current
// instance of DirMgr to their initial or zero
// values.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete all pre-existing internal member
// variable data values in the current instance of DirMgr.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dMgr							*DirMgr
//
//		A pointer to an instance of DirMgr. This method
//		will delete and reset all member variable data
//		values to their initial or zero states.
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
func (dMgrHlprElectron *dirMgrHelperElectron) empty(
	dMgr *DirMgr,
	dMgrLabel string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if dMgrHlprElectron.lock == nil {
		dMgrHlprElectron.lock = new(sync.Mutex)
	}

	dMgrHlprElectron.lock.Lock()

	defer dMgrHlprElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"dirMgrHelperElectron.empty()",
		"")

	if err != nil {
		return err
	}

	if dMgr == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' pointer is 'nil'!\n",
			ePrefix.String(),
			dMgrLabel)
	}

	if len(dMgrLabel) == 0 {
		dMgrLabel = "dMgr"
	}

	dMgr.isInitialized = false
	dMgr.originalPath = ""
	dMgr.path = ""
	dMgr.isPathPopulated = false
	dMgr.doesPathExist = false
	dMgr.parentPath = ""
	dMgr.isParentPathPopulated = false
	dMgr.absolutePath = ""
	dMgr.isAbsolutePathPopulated = false
	dMgr.doesAbsolutePathExist = false
	dMgr.isAbsolutePathDifferentFromPath = false
	dMgr.directoryName = ""
	dMgr.volumeName = ""
	dMgr.isVolumePopulated = false
	dMgr.actualDirFileInfo = FileInfoPlus{}

	return err
}

// isPathStringEmptyOrBlank
//
// Determines whether a path string is blank.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathStr						string
//
//		Contains the path string to be validated.
//
//	trimTrailingPathSeparator	bool
//
//		If this parameter is set to 'true', the returned
//		path string will be trimmed. This means leading
//		and trailing spaces will be deleted from the
//		returned path string.
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
//	pathFileNameExt				string
//
//		This parameter returns the path, file name and
//		file extension (if present). If input parameter
//		'trimTrailingPathSeparator' is set to 'true',
//		leading and trailing spaces will be deleted.
//
//	strLen						int
//
//		Returns the length of the path, file name and
//		file extension string, 'pathFileNameExt'.
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
func (dMgrHlprElectron *dirMgrHelperElectron) isPathStringEmptyOrBlank(
	pathStr string,
	trimTrailingPathSeparator bool,
	pathStrLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	pathFileNameExt string,
	strLen int,
	err error) {

	if dMgrHlprElectron.lock == nil {
		dMgrHlprElectron.lock = new(sync.Mutex)
	}

	dMgrHlprElectron.lock.Lock()

	defer dMgrHlprElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"dirMgrHelperElectron.isPathStringEmptyOrBlank()",
		"")

	if err != nil {
		return pathFileNameExt, strLen, err
	}

	strLen = len(pathStr)

	if strLen == 0 {
		err = fmt.Errorf("%v\n"+
			"ERROR: %v is an empty string!\n",
			ePrefix.String(),
			pathStrLabel)

		return pathFileNameExt, strLen, err
	}

	if len(pathStrLabel) == 0 {
		pathStrLabel = "pathStr"
	}

	pathFileNameExt = strings.TrimLeft(pathStr, " ")

	pathFileNameExt = strings.TrimRight(pathFileNameExt, " ")

	strLen = len(pathFileNameExt)

	if strLen == 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: %v consists entirely of blank spaces!\n",
			ePrefix.String(),
			pathStrLabel)

		return pathFileNameExt, strLen, err
	}

	pathFileNameExt =
		new(FileHelper).AdjustPathSlash(pathFileNameExt)

	dotPathSeparator := "." + string(os.PathSeparator)

	if strings.HasSuffix(pathFileNameExt, dotPathSeparator) {
		trimTrailingPathSeparator = false
	}

	strLen = len(pathFileNameExt)

	if trimTrailingPathSeparator &&
		pathFileNameExt[strLen-1] == os.PathSeparator {

		pathFileNameExt = pathFileNameExt[0 : strLen-1]
		strLen = len(pathFileNameExt)

	}

	return pathFileNameExt, strLen, err
}

// getSubDirsFilesInDirTree
//
// Returns selected subdirectories and/or files in an
// entire directory tree defined by the parent directory
// input parameter 'targetDMgr'. This means that the
// parent directory and all subdirectories in the
// directory tree will be scanned to identify and return
// subdirectories and/or files matching the specified
// selection criteria. Selected subdirectories will be
// returned in a Directory Manager Collection while
// selected files will be returned in a File Manager
// Collection.
//
// The parent directory 'targetDMgr' will NOT be
// returned with selected subdirectories.
//
// To qualify for selection and inclusion in the returned
// Directory and File Manager Collections, items residing
// in the 'targetDMgr' target directory tree are divided
// into two classes, subdirectories and files.
// Subdirectories are standard os.FileInfo directory
// entries. Files are defined as all artifacts residing
// in the target directory tree which are not classified
// as subdirectories.
//
//	Subdirectory Screening and Selection
//
// To qualify as a selected subdirectory, the
// subdirectory must satisfy two filters. First, input
// parameter 'getSubdirectories' must be set to
// 'true'.
//
// Second, the subdirectory must satisfy the Directory
// Characteristics Selection Criteria specified by input
// parameter, 'subDirSelectCharacteristics'. This
// parameter is of Type FileSelectionCriteria and allows
// users to screen and select subdirectories by Name,
// Directory Modification Date and Mode. Directory Name
// selections can be based on pattern matches or regular
// expression matches.
//
// If both of these filter requirements are satisfied,
// the subdirectory will be added to, and returned by,
// the Directory Manager Collection passed as input
// parameter 'subDirsInDir'.
//
// Be advised that users control the behavior for current
// directories (".") and parent directories ("..") with
// input parameters 'includeSubDirCurrentDirOneDot' and
// 'includeSubDirParentDirTwoDots'.
//
//	File Screening and Selection
//
// To qualify as a selected file, the os.FileInfo entry
// must comply with two filters: File Type and File
// Characteristics Selection Criteria. Remember that
// files are defined as all artifacts residing in the
// target directory tree which are not classified as
// subdirectories.
//
// To be eligible for file selection, the os.FileInfo
// entry must first comply with the specified File Type
// criteria. In terms of File Type, files are classified
// as Regular Files, SymLink Files or Other Non-Regular Files.
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
// files, SymLink files and Other Non-Regular Files such
// as device files, named pipes and sockets. The Other
// Non-Regular Files is a catch-all category including all
// files which are NOT classified as Regular Files or
// SymLink Files.
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
//		subdirectories and/or files in the target
//		directory tree specified by input parameter
//		'targetDMgr'. The entire target directory tree
//		will be searched for eligible subdirectories
//		and/or files.
//
//	(2) Selected subdirectories must fulfill two
//		requirements.
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
//		subdirectories and/or files will be added to the
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
//			------------------------------------------------------------------------
//
//			IMPORTANT:
//
//			If all the file selection criterion in the FileSelectionCriteria
//			object are 'Inactive' or 'Not Set' (set to their zero or default values),
//			then all the files meeting the File Type requirements in the directory
//			defined by 'targetDMgr' will be selected.
//
//				Example:
//				  fsc := FileSelectCriterionMode{}
//
//				  In this example, 'fsc' is NOT initialized. Therefore,
//				  all the selection criterion are 'Inactive'. Consequently,
//				  all the files meeting the File Type requirements in the
//				  directory defined by 'targetDMgr' will be selected.
//
//			------------------------------------------------------------------------
//
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
func (dMgrHlprElectron *dirMgrHelperElectron) getSubDirsFilesInDirTree(
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

	if dMgrHlprElectron.lock == nil {
		dMgrHlprElectron.lock = new(sync.Mutex)
	}

	dMgrHlprElectron.lock.Lock()

	defer dMgrHlprElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "dirMgrHelperElectron." +
		"getSubDirsFilesInDirTree()"

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

	if err != nil {

		return numOfSubDirsLocated, numOfFilesLocated, err
	}

	var idx = 0

	if subDirsInDir != nil {

		idx = len(subDirsInDir.dirMgrs)
	}

	var dMgrHlprBoson = new(dirMgrHelperBoson)

	numOfSubDirsLocated,
		numOfFilesLocated,
		err = dMgrHlprBoson.
		getSubDirsFilesInDir(
			targetDMgr,
			true,
			includeSubDirCurrentDirOneDot,
			includeSubDirParentDirTwoDots,
			getRegularFiles,
			getSymLinksFiles,
			getOtherNonRegularFiles,
			subDirSelectCharacteristics,
			fileSelectCriteria, // fileSelectCriteria
			subDirsInDir,
			filesInDir,
			targetDMgrLabel,
			ePrefix.XCpy(targetDMgrLabel))

	if err != nil {

		return numOfSubDirsLocated, numOfFilesLocated, err
	}

	if numOfSubDirsLocated == 0 {
		// No subdirectories matched the search
		// criteria in the top-level target
		// directory. The search operation is
		// there finished.

		return numOfSubDirsLocated, numOfFilesLocated, err
	}

	// Subdirectories were located.

	var subDirDMgr DirMgr
	var newNumOfSubDirsLocated, newNumOfFilesLocated int
	var dMgrColHelper = new(dirMgrCollectionHelper)
	var errStatus ArrayColErrorStatus

	idx--

	for idx > -2 {

		idx++

		// This is a peek operation
		subDirDMgr,
			errStatus = dMgrColHelper.
			peekOrPopAtIndex(
				subDirsInDir,
				idx,
				false, // deleteIndex
				ePrefix)

		if errStatus.ProcessingError != nil {

			if errStatus.IsIndexOutOfBounds == true {

				idx = -2

				break
			}

			err = fmt.Errorf("%v\n"+
				"Error occurred while extracting DirMgr from 'subDirectories'.\n"+
				"dirMgrCollectionHelper.peekOrPopAtIndex(subDirectories,index)\n"+
				"index= %v\n"+
				"Error= \n%v\n",
				funcName,
				idx,
				errStatus.ProcessingError.Error())

			return numOfSubDirsLocated, numOfFilesLocated, err
		}

		newNumOfSubDirsLocated,
			newNumOfFilesLocated,
			err = dMgrHlprBoson.
			getSubDirsFilesInDir(
				&subDirDMgr,
				getSubdirectories,
				includeSubDirCurrentDirOneDot,
				includeSubDirParentDirTwoDots,
				getRegularFiles,
				getSymLinksFiles,
				getOtherNonRegularFiles,
				subDirSelectCharacteristics,
				fileSelectCriteria, // fileSelectCriteria
				subDirsInDir,
				filesInDir,
				"subDirDMgr",
				ePrefix.XCpy(fmt.Sprintf("subDirDMgr idx=[%v]", idx)))

		if err != nil {

			return numOfSubDirsLocated, numOfFilesLocated, err
		}

		numOfSubDirsLocated += newNumOfSubDirsLocated

		numOfFilesLocated += newNumOfFilesLocated

	}

	return numOfSubDirsLocated, numOfFilesLocated, err
}

// lowLevelDoesDirectoryExist
//
// This method tests for the existence of directory path.
func (dMgrHlprElectron *dirMgrHelperElectron) lowLevelDoesDirectoryExist(
	dirPath,
	dirPathLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	dirPathDoesExist bool,
	fInfoPlus FileInfoPlus,
	err error) {

	if dMgrHlprElectron.lock == nil {
		dMgrHlprElectron.lock = new(sync.Mutex)
	}

	dMgrHlprElectron.lock.Lock()

	defer dMgrHlprElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	dirPathDoesExist = false

	funcName := "dirMgrHelperElectron." +
		"lowLevelDoesDirectoryExist()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return dirPathDoesExist, fInfoPlus, err
	}

	if len(dirPathLabel) == 0 {
		dirPathLabel = "DirMgr"
	}

	errCode := 0

	errCode,
		_,
		dirPath = new(FileHelper).
		IsStringEmptyOrBlank(dirPath)

	if errCode < 0 {
		err = fmt.Errorf("%v\n"+
			"Error: This Directory Manager instance is invalid!\n"+
			"Input paramter %v.path is an empty string.\n",
			ePrefix.String(),
			dirPathLabel)

		return dirPathDoesExist, fInfoPlus, err
	}

	var err2 error
	var info os.FileInfo

	for i := 0; i < 3; i++ {

		dirPathDoesExist = false
		fInfoPlus = FileInfoPlus{}
		err = nil

		info,
			err2 = os.Stat(dirPath)

		if err2 != nil {

			if os.IsNotExist(err2) {

				dirPathDoesExist = false
				fInfoPlus = FileInfoPlus{}
				err = nil

				return dirPathDoesExist, fInfoPlus, err
			}

			// err == nil and err != os.IsNotExist(err)
			// This is a non-path error. The non-path error will be
			// tested up to 3-times before it is returned.
			err = fmt.Errorf("%v\n"+
				"Error: This Directory Manager instance is invalid!\n"+
				"Non-Path error returned by os.Stat(%v) while attempty\n"+
				"to acquire os.FileInfo data on this directory path.\n"+
				"%v= %v\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				dirPathLabel,
				dirPathLabel,
				dirPath,
				err2.Error())

			fInfoPlus = FileInfoPlus{}

			dirPathDoesExist = false

		} else {
			// err == nil
			// The path really does exist!
			dirPathDoesExist = true
			err = nil

			fInfoPlus,
				err2 = new(FileInfoPlus).
				NewFromPathFileInfo(
					dirPath,
					info,
					ePrefix.XCpy("dirPath"))

			if err2 != nil {

				err = fmt.Errorf("%v\n"+
					"Error: An internal error occurred while processing\n"+
					"the os.FileInfo data for %v directory path.\n"+
					"Error returned by FileInfoPlus{}.NewFromPathFileInfo(dirPath, info)\n"+
					"dirPath= '%v'\n"+
					"Error= \n%v\n",
					funcName,
					dirPathLabel,
					dirPath,
					err2.Error())

				fInfoPlus = FileInfoPlus{}
			}

			return dirPathDoesExist, fInfoPlus, err
		}

		time.Sleep(30 * time.Millisecond)
	}

	return dirPathDoesExist, fInfoPlus, err
}
