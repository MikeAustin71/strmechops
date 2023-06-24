package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type DirHelper struct {
	lock *sync.Mutex
}

// GetDirectoryProfile
//
// This method returns an instance of DirectoryProfile
// which includes subdirectory and file statistics on the
// directory path specified by input parameter
// 'directoryPath'.
//
// Directory profile information will only be collected
// on the top-level or parent directory defined by input
// parameter 'directoryPath'. Information on the contents
// of subdirectories will NOT be included in the directory
// profile information returned by this method.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	directoryPath				string
//
//		This string contains the directory path which
//		will be analyzed for the following objectives:
//
//			(1)	Determine if the directory path exists
//				on an attached storage drive.
//
//			(2) If the path does exist, statistics on
//				the directory will be generated and
//				returned via an instance of
//				DirectoryProfile.
//
//		Directory profile information returned by this
//		method will only include the top-level or parent
//		directory defined by 'directoryPath'. Profile
//		Information on the contents of subdirectories
//		will NOT be included.
//
//		'directoryPath' may be formatted as a relative
//		path or an absolute path.
//
//		If 'directoryPath' is invalid, an error will be
//		returned.
//
//	includeSubDirCurrenDirOneDot		bool
//
//		All directories include an os.FileInfo entry for
//		the current directory. The current directory name
//		is always denoted as single dot ('.').
//
//		When this parameter, 'includeSubDirCurrenDirOneDot',
//		is set to 'true', the current directory, designated
//		as a single dot ('.'), will be included in the
//		directory profile information returned by this
//		method.
//
//	includeSubDirParentDirTwoDots 		bool
//
//		All directories include an os.FileInfo entry for
//		the parent directory. The parent directory name
//		is always denoted as two dots ('..').
//
//		When this parameter, 'includeSubDirParentDirTwoDots',
//		is set to 'true', the parent directory, designated
//		as two dots ('..'), will be included in the
//		directory profile information returned by this
//		method.
//
//	fileSelectCharacteristics	FileSelectionCriteria
//
//		Files selected for description in the directory
//		profile statistics must conform to the File
//		Characteristics Criteria specified by this input
//		parameter, 'fileSelectCharacteristics'.
//
//		Files matching these selection criteria will be
//		included in the directory profile information
//		returned by this method.
//
//		File Characteristics Criteria allow the user to
//		screen files based on File Name, File
//		Modification Date and File Mode. In addition,
//		users have the option to filter File Names using
//		pattern matches or regular expressions.
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
//		If all file selection criterion in the FileSelectionCriteria object
//		are 'Inactive' or 'Not Set' (set to their zero or default values),
//		then all the files in the directory path specified by input parameter
//		'dMgr' will be selected.
//
//			Example:
//			  fsc := FileSelectCriterionMode{}
//
//		 	In this example, 'fsc' is NOT initialized. Therefore,
//			all the selection criterion are 'Inactive'. Consequently,
//			all files encountered in the target directory
//			path during the search operation will meet the file characteristics
//			selection criteria and will therefore be classified as eligible for
//			selection.
//
//		------------------------------------------------------------------------
//
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	directoryPathDoesExist		bool
//
//		If this parameter returns a boolean value of
//		'true', it signals that the directory path
//		specified by input parameter 'directoryPath'
//		actually exists on an attached storage drive.
//
//	dirProfile					DirectoryProfile
//
//		If this method completes successfully, this
//		returned instance of DirectoryProfile will be
//		populated with profile and statistical
//		information on the top-level or parent directory
//		identified by input parameter 'directoryPath'.
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
func (dHlpr *DirHelper) GetDirectoryProfile(
	directoryPath string,
	includeSubDirCurrenDirOneDot bool,
	includeSubDirParentDirTwoDots bool,
	fileSelectCharacteristics FileSelectionCriteria,
	errorPrefix interface{}) (
	directoryPathDoesExist bool,
	dirProfile DirectoryProfile,
	err error) {

	if dHlpr.lock == nil {
		dHlpr.lock = new(sync.Mutex)
	}

	dHlpr.lock.Lock()

	defer dHlpr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "DirHelper." +
		"GetDirectoryProfile()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {
		return directoryPathDoesExist, dirProfile, err
	}

	var dMgr DirMgr
	var err2 error

	dMgr,
		err2 = new(DirMgr).New(
		directoryPath,
		ePrefix.XCpy("dMgr<-directoryPath"))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: directoryPath is NOT a valid directory path!\n"+
			"Error= \n%v\n",
			funcName,
			err2.Error())

		return directoryPathDoesExist, dirProfile, err
	}

	directoryPathDoesExist,
		dirProfile,
		err = new(dirMgrHelperTachyon).
		getDirectoryProfile(
			&dMgr,
			includeSubDirCurrenDirOneDot,
			includeSubDirParentDirTwoDots,
			fileSelectCharacteristics,
			"dMgr",
			ePrefix.XCpy("directoryPath->dMgr"))

	return directoryPathDoesExist, dirProfile, err
}

// GetDirectoryTreeProfile
//
// Returns an instance of DirectoryProfile which contains
// directory and file statistics on an entire directory
// tree. The target directory tree is defined by the
// directory absolute path passed by input parameter
// 'directoryPath'.
//
// Subject to input parameter 'skipTopLevelDirectory',
// directory profile information will be collected and
// returned on all selected files and directories in the
// target directory tree.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	directoryPath				string
//
//		This string contains the directory path which
//		defines the directory tree for which directory
//		profile information will be collected and
//		returned by this method.
//
//		Directory profile information returned by this
//		method will only include the entire directory
//		tree defined by 'directoryPath'. There returned
//		directory profile information will include the
//		parent directory 'directoryPath' plus the
//		contents of subdirectories in the directory tree.
//
//		'directoryPath' may be formatted as a relative
//		path or an absolute path.
//
//		If 'directoryPath' is invalid, an error will be
//		returned.
//
//	skipTopLevelDirectory		bool
//
//		If this parameter is set to 'true', the top level
//		directory, defined by input parameter 'dMgr',
//		will be skipped, and it will not be included in
//		the directory profile information accumulated for
//		the directory tree.
//
//	includeSubDirCurrenDirOneDot		bool
//
//		All directories include an os.FileInfo entry for
//		the current directory. The current directory name
//		is always denoted as single dot ('.').
//
//		When this parameter, 'includeSubDirCurrenDirOneDot',
//		is set to 'true', the current directory, designated
//		as a single dot ('.'), will be included in the
//		directory profile information returned by this
//		method.
//
//	includeSubDirParentDirTwoDots 		bool
//
//		All directories include an os.FileInfo entry for
//		the parent directory. The parent directory name
//		is always denoted as two dots ('..').
//
//		When this parameter, 'includeSubDirParentDirTwoDots',
//		is set to 'true', the parent directory, designated
//		as two dots ('..'), will be included in the
//		directory profile information returned by this
//		method.
//
//	fileSelectCharacteristics	FileSelectionCriteria
//
//		Files selected for inclusion in the directory
//		profile statistics must conform to the File
//		Characteristics Criteria specified by this input
//		parameter, 'fileSelectCharacteristics'.
//
//		Files matching these selection criteria will be
//		included in the directory profile information
//		returned by this method.
//
//		File Characteristics Criteria allow the user to
//		screen files based on File Name, File
//		Modification Date and File Mode. In addition,
//		users have the option to filter File Names using
//		pattern matches or regular expressions.
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
//		If all file selection criterion in the FileSelectionCriteria object
//		are 'Inactive' or 'Not Set' (set to their zero or default values),
//		then all the files in the directory path specified by input parameter
//		'dMgr' will be selected and included in the directory profile
//		returned by this method.
//
//			Example:
//			  fsc := FileSelectCriterionMode{}
//
//		 	In this example, 'fsc' is NOT initialized. Therefore,
//			all the selection criterion are 'Inactive'. Consequently,
//			all files encountered in the target directory
//			path during the search operation will meet the file characteristics
//			selection criteria and will therefore be included in the directory
//			profile returned by this method.
//
//		------------------------------------------------------------------------
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	directoryPathDoesExist		bool
//
//		If this parameter returns a boolean value of
//		'true', it signals that the directory path
//		specified by input parameter 'directoryPath'
//		actually exists on an attached storage drive.
//
//	dirTreeProfile				DirectoryProfile
//
//		If this method completes successfully, this
//		returned instance of DirectoryProfile will be
//		populated with profile and statistical
//		information on directories and files in the
//		entire directory tree defined by input parameter
//		'directoryPath'.
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
func (dHlpr *DirHelper) GetDirectoryTreeProfile(
	directoryPath string,
	skipTopLevelDirectory bool,
	includeSubDirCurrenDirOneDot bool,
	includeSubDirParentDirTwoDots bool,
	fileSelectCharacteristics FileSelectionCriteria,
	errorPrefix interface{}) (
	directoryPathDoesExist bool,
	dirTreeProfile DirectoryProfile,
	err error) {

	if dHlpr.lock == nil {
		dHlpr.lock = new(sync.Mutex)
	}

	dHlpr.lock.Lock()

	defer dHlpr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "DirHelper." +
		"GetDirectoryProfile()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {
		return directoryPathDoesExist, dirTreeProfile, err
	}

	var dMgr DirMgr
	var err2 error

	dMgr,
		err2 = new(DirMgr).New(
		directoryPath,
		ePrefix.XCpy("dMgr<-directoryPath"))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: directoryPath is NOT a valid directory path!\n"+
			"Error= \n%v\n",
			funcName,
			err2.Error())

		return directoryPathDoesExist, dirTreeProfile, err
	}

	directoryPathDoesExist,
		dirTreeProfile,
		err = new(dirMgrHelperPreon).
		getDirectoryTreeProfile(
			&dMgr,
			skipTopLevelDirectory,
			includeSubDirCurrenDirOneDot,
			includeSubDirParentDirTwoDots,
			fileSelectCharacteristics,
			"dMgr",
			ePrefix)

	return directoryPathDoesExist, dirTreeProfile, err
}

// GetSubdirectoriesDirTree
//
// This method scans and identifies selected
// subdirectories residing in the directory tree defined
// by input parameter 'directoryPath'. Subdirectories
// meeting the selection criteria and located in this
// directory tree are returned to the user by means of a
// Directory Manager Collection (DirMgrCollection) passed
// as input parameter 'subDirsInDir'.
//
//	Subdirectory Screening and Selection
//
// To qualify as a selected subdirectory, the
// subdirectory must satisfy the Directory
// Characteristics Selection Criteria specified by input
// parameter, 'subDirSelectCharacteristics'. This
// parameter is of Type FileSelectionCriteria and allows
// users to screen and select subdirectories by Name,
// Directory Modification Date and Mode. Directory Name
// selections can be based on pattern matches or regular
// expression matches.
//
// If this filter requirement is satisfied, the
// subdirectory will be added to, and returned by,
// the Directory Manager Collection passed as input
// parameter 'subDirsInDir'.
//
// To select all subdirectories in the directory tree,
// turn off the Directory Characteristics Selection
// Criteria by setting this parameter to an empty
// instance of FileSelectionCriteria.
//
//	Example: subDirSelectCharacteristics =
//				FileSelectionCriteria{}
//
// Be advised that users control the behavior for current
// directories (".") and parent directories ("..") with
// input parameters 'includeSubDirCurrenDirOneDot' and
// 'includeSubDirParentDirTwoDots'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	This method will search for subdirectories
//		located in the directory tree defined by input
//		parameter 'directoryPath'. Those subdirectories
//		which satisfy the Directory Characteristics
//		Selection Criteria specified by input parameter
//		'subDirSelectCharacteristics', will be added and
//		returned in the Directory Manager Collection
//		passed as input parameter 'subDirsInDir'.
//
//	(2)	To select all subdirectories in the directory
//		tree, turn off the Directory Characteristics
//		Selection Criteria by setting this parameter
//		('subDirSelectCharacteristics') to an empty
//		instance of FileSelectionCriteria.
//
//			Example:
//				subDirSelectCharacteristics =
//					FileSelectionCriteria{}
//
//	(3)	All directories include an os.FileInfo entry for
//		the current directory. The current directory name
//		is always denoted as single dot ('.'). Users can
//		include or exclude the selection of the current
//		directory entry by configuring the boolean input
//		parameter 'includeSubDirCurrenDirOneDot'.
//
//	(4)	All directories include an os.FileInfo entry for
//		the parent directory. The parent directory name
//		is always denoted as two dots ('..'). Users can
//		include or exclude the selection of the parent
//		directory entry by configuring the boolean input
//		parameter 'includeSubDirParentDirTwoDots'.
//
//	(5)	For a collection of subdirectories residing
//		exclusively in the top level or parent directory
//		specified by a directory path, see method:
//
//			DirHelper.GetSubdirectoriesParentDir
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	directoryPath				string
//
//		This string defines a directory path. This path
//		will be treated as a directory tree. All
//		subdirectories residing in this directory tree
//		will be added and returned to the user by means
//		of the Directory Manager Collection passed as
//		input parameter 'subDirsInDir'.
//
//		If this directory path does not exist on an
//		attached storage drive, an error will be
//		returned.
//
//	includeSubDirCurrenDirOneDot		bool
//
//		This parameter is only used, if input parameter
//		'getSubdirectories' is set to 'true'.
//
//		All directories include an os.FileInfo entry for
//		the current directory. The current directory name
//		is always denoted as single dot ('.').
//
//		When this parameter, 'includeSubDirCurrenDirOneDot',
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
//	subDirSelectCharacteristics	FileSelectionCriteria
//
//		This subdirectory selection criteria allows users
//		to screen subdirectories for Name, Modification
//		Date and File Mode. Subdirectory Name selections
//		can be configured for pattern matches or regular
//		expression matches.
//
//		Directory os.FileIno entries matching this
//		selection criteria will be included in the
//		Directory Manager Collection returned by input
//		parameter 'subDirsInDir'.
//
//		Remember that setting 'subDirSelectCharacteristics'
//		to an empty instance of FileSelectionCriteria will
//		ensure that all subdirectories are selected.
//
//			Example:
//			subDirSelectCharacteristics =
//				FileSelectionCriteria{}
//
//			This ensures that all subdirectories will satisfy
//			the Directory Characteristics Selection Criteria.
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
//		If all of the file selection criterion in the FileSelectionCriteria
//		object are 'Inactive' or 'Not Set' (set to their zero or default values),
//		then all the subdirectories meeting the File Type requirements in the
//		directory tree defined by 'targetDMgr' will be selected.
//
//			Example:
//			  fsc := FileSelectCriterionMode{}
//
//			  In this example, 'fsc' is NOT initialized. Therefore,
//			  all the selection criterion are 'Inactive'. Consequently,
//			  all the subdirectories in the directory tree defined by
//			  'targetDMgr' will be selected.
//
//		------------------------------------------------------------------------
//
//	subDirsInDir				*DirMgrCollection
//
//		A pointer to an instance of DirMgrCollection
//		which encapsulates an array of Directory Manager
//		(DirMgr) objects.
//
//		This method will scan the entire directory tree
//		defined by input parameter 'directoryPath'.	All
//		subdirectories found in this directory tree will
//		be configured as Directory Manager (DirMgr)
//		objects and added to this Directory Manager
//		Collection ('subDirsInDir').
//
//		Directory entries for the current directory (".")
//		and the parent directory ("..") will be skipped
//		and will NOT be added to the 'subDirsInDir'
//		Directory Manager Collection.
//
//			type DirMgrCollection struct {
//				dirMgrs []DirMgr
//			}
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	numOfSubdirectories			int
//
//		If this method completes successfully without
//		error, this parameter will return the number
//		of subdirectories located in the directory tree
//		defined by input parameter 'directoryPath'.
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
func (dHlpr *DirHelper) GetSubdirectoriesDirTree(
	directoryPath string,
	includeSubDirCurrenDirOneDot bool,
	includeSubDirParentDirTwoDots bool,
	subDirSelectCharacteristics FileSelectionCriteria,
	subDirsInDir *DirMgrCollection,
	errorPrefix interface{}) (
	numOfSubdirectories int,
	err error) {

	if dHlpr.lock == nil {
		dHlpr.lock = new(sync.Mutex)
	}

	dHlpr.lock.Lock()

	defer dHlpr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "DirHelper." +
		"GetSubdirectoriesDirTree()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {
		return numOfSubdirectories, err
	}

	var dMgr DirMgr
	var err2 error

	dMgr,
		err2 = new(DirMgr).New(
		directoryPath,
		ePrefix.XCpy("dMgr<-directoryPath"))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input paramter 'directoryPath' is invalid!\n"+
			"'directoryPath' is NOT a valid directory path!\n"+
			"Error= \n%v\n",
			funcName,
			err2.Error())

		return numOfSubdirectories, err
	}

	if subDirsInDir == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'subDirsInDir' is invalid!\n"+
			"'subDirsInDir' is a nil pointer.\n",
			ePrefix.String())

		return numOfSubdirectories, err
	}

	if !dMgr.DoesDirectoryExist() {

		err = fmt.Errorf("%v\n"+
			"Error: Input paramter 'directoryPath' is invalid!\n"+
			"'directoryPath' does NOT exist on an attached\n"+
			"storage drive.\n",
			ePrefix.String())

		return numOfSubdirectories, err
	}

	numOfSubdirectories,
		_,
		err = new(dirMgrHelperElectron).
		getSubDirsFilesInDirTree(
			&dMgr,
			true, // getSubdirectories
			includeSubDirCurrenDirOneDot,
			includeSubDirParentDirTwoDots,
			false, // getRegularFiles
			false, // getSymLinksFiles
			false, // getOtherNonRegularFiles
			subDirSelectCharacteristics,
			FileSelectionCriteria{}, // fileSelectCriteria
			subDirsInDir,
			nil,
			"dMgr",
			ePrefix)

	return numOfSubdirectories, err
}

// GetSubdirectoriesParentDir
//
// This method scans and identifies selected
// subdirectories residing in the top-level or parent
// directory defined by input parameter 'directoryPath'.
// Subdirectories meeting the selection criteria and
// located in this single directory are returned to the
// user by means of a Directory Manager Collection
// (DirMgrCollection) passed as input parameter
// 'subDirsInDir'.
//
//	Subdirectory Screening and Selection
//
// To qualify as a selected subdirectory, the
// subdirectory must satisfy the Directory
// Characteristics Selection Criteria specified by input
// parameter, 'subDirSelectCharacteristics'. This
// parameter is of Type FileSelectionCriteria and allows
// users to screen and select subdirectories by Name,
// Directory Modification Date and Mode. Directory Name
// selections can be based on pattern matches or regular
// expression matches.
//
// If this filter requirement is satisfied, the
// subdirectory will be added to, and returned by,
// the Directory Manager Collection passed as input
// parameter 'subDirsInDir'.
//
// To select all subdirectories in the parent directory,
// turn off the Directory Characteristics Selection
// Criteria by setting this parameter to an empty
// instance of FileSelectionCriteria.
//
//	Example: subDirSelectCharacteristics =
//				FileSelectionCriteria{}
//
// Be advised that users control the behavior for current
// directories (".") and parent directories ("..") with
// input parameters 'includeSubDirCurrenDirOneDot' and
// 'includeSubDirParentDirTwoDots'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	This method will search for subdirectories
//		located in the parent directory defined by input
//		parameter 'directoryPath'. Those subdirectories
//		which satisfy the Directory Characteristics
//		Selection Criteria specified by input parameter
//		'subDirSelectCharacteristics', will be added and
//		returned in the Directory Manager Collection
//		passed as input parameter 'subDirsInDir'.
//
//	(2)	To select all subdirectories in the parent
//		directory, turn off the Directory Characteristics
//		Selection Criteria by setting this parameter
//		('subDirSelectCharacteristics') to an empty
//		instance of FileSelectionCriteria.
//
//			Example:
//				subDirSelectCharacteristics =
//					FileSelectionCriteria{}
//
//	(3)	All directories include an os.FileInfo entry for
//		the current directory. The current directory name
//		is always denoted as single dot ('.'). Users can
//		include or exclude the selection of the current
//		directory entry by configuring the boolean input
//		parameter 'includeSubDirCurrenDirOneDot'.
//
//	(4)	All directories include an os.FileInfo entry for
//		the parent directory. The parent directory name
//		is always denoted as two dots ('..'). Users can
//		include or exclude the selection of the parent
//		directory entry by configuring the boolean input
//		parameter 'includeSubDirParentDirTwoDots'.
//
//	(5)	For a collection of subdirectories in the
//		directory tree specified by a directory path, see
//		method:
//
//			DirHelper.GetSubdirectoriesDirTree
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	directoryPath					string
//
//		This string defines a directory path. This path
//		will be treated as a top level or parent
//		directory. Subdirectories residing in this single
//		parent directory will be added and returned to
//		the user by means of the Directory Manager
//		Collection passed as input parameter
//		'subDirsInDir'.
//
//		The search for subdirectories is limited
//		exclusively to this	parent directory and
//		does NOT extend to lower levels of the
//		directory tree.
//
//		If this directory path does not exist on an
//		attached storage drive, an error will be
//		returned.
//
//	includeSubDirCurrenDirOneDot	bool
//
//		All directories automatically include an
//		os.FileInfo entry for the current directory. The
//		current directory name is always denoted as
//		single dot ('.').
//
//		When this parameter, 'includeSubDirCurrenDirOneDot',
//		is set to 'true', the current directory, designated
//		as a single dot ('.'), will be added to the Directory
//		Manager Collection passed as input parameter
//		'subDirsInDir'.
//
//	includeSubDirParentDirTwoDots	bool
//
//		All directories include an os.FileInfo entry for
//		the parent directory. The parent directory name
//		is always denoted as two dots ('..').
//
//		When this parameter, 'includeSubDirParentDirTwoDots',
//		is set to 'true', the parent directory, designated
//		as two dots ('..'), will be added to the Directory
//		Manager Collection passed as input parameter
//		'subDirsInDir'.
//
//	subDirSelectCharacteristics	FileSelectionCriteria
//
//		This subdirectory selection criteria allows users
//		to screen subdirectories for Name, Modification
//		Date and File Mode. Subdirectory Name selections
//		can be configured for pattern matches or regular
//		expression matches.
//
//		Directory os.FileIno entries matching this
//		selection criteria will be included in the
//		Directory Manager Collection returned by input
//		parameter 'subDirsInDir'.
//
//		Remember that setting 'subDirSelectCharacteristics'
//		to an empty instance of FileSelectionCriteria will
//		ensure that all subdirectories are selected.
//
//			Example:
//			subDirSelectCharacteristics =
//				FileSelectionCriteria{}
//
//			This ensures that all subdirectories will satisfy
//			the Directory Characteristics Selection Criteria.
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
//		If all of the file selection criterion in the FileSelectionCriteria
//		object are 'Inactive' or 'Not Set' (set to their zero or default values),
//		then all the subdirectories meeting the File Type requirements in the
//		directory tree defined by 'targetDMgr' will be selected.
//
//			Example:
//			  fsc := FileSelectCriterionMode{}
//
//			  In this example, 'fsc' is NOT initialized. Therefore,
//			  all the selection criterion are 'Inactive'. Consequently,
//			  all the subdirectories in the directory tree defined by
//			  'targetDMgr' will be selected.
//
//		------------------------------------------------------------------------
//
//	subDirsInDir					*DirMgrCollection
//
//		A pointer to an instance of DirMgrCollection
//		which encapsulates an array of Directory Manager
//		(DirMgr) objects.
//
//		This method will scan the top level or parent
//		directory defined by the input parameter
//		'directoryPath'. Any subdirectories found in this
//		parent directory will be configured as Directory
//		Manager (DirMgr) objects and added to this
//		Directory Manager Collection ('subDirsInDir').
//
//		Directory entries for the current directory (".")
//		and the parent directory ("..") will be skipped
//		and will NOT be added to the 'subDirsInDir'
//		Directory Manager Collection.
//
//			type DirMgrCollection struct {
//				dirMgrs []DirMgr
//			}
//
//	errorPrefix						interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	numOfSubdirectories				int
//
//		If this method completes successfully, this
//		integer value represents the number of
//		subdirectories added to the Directory Manager
//		Collection passed as input parameter
//		'subDirsInDir'.
//
//	err								error
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
func (dHlpr *DirHelper) GetSubdirectoriesParentDir(
	directoryPath string,
	includeSubDirCurrenDirOneDot bool,
	includeSubDirParentDirTwoDots bool,
	subDirSelectCharacteristics FileSelectionCriteria,
	subDirsInDir *DirMgrCollection,
	errorPrefix interface{}) (
	numOfSubdirectories int,
	err error) {

	if dHlpr.lock == nil {
		dHlpr.lock = new(sync.Mutex)
	}

	dHlpr.lock.Lock()

	defer dHlpr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "DirHelper." +
		"GetSubdirectoriesParentDir()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {
		return numOfSubdirectories, err
	}

	var dMgr DirMgr
	var err2 error

	dMgr,
		err2 = new(DirMgr).New(
		directoryPath,
		ePrefix.XCpy("dMgr<-directoryPath"))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input paramter 'directoryPath' is invalid!\n"+
			"'directoryPath' is NOT a valid directory path!\n"+
			"Error= \n%v\n",
			funcName,
			err2.Error())

		return numOfSubdirectories, err
	}

	if subDirsInDir == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'subDirsInDir' is invalid!\n"+
			"'subDirsInDir' is a nil pointer.\n",
			ePrefix.String())

		return numOfSubdirectories, err
	}

	if !dMgr.DoesDirectoryExist() {

		err = fmt.Errorf("%v\n"+
			"Error: Input paramter 'directoryPath' is invalid!\n"+
			"'directoryPath' does NOT exist on an attached\n"+
			"storage drive.\n",
			ePrefix.String())

		return numOfSubdirectories, err

	}

	numOfSubdirectories,
		_,
		err = new(dirMgrHelperBoson).
		getSubDirsFilesInDir(
			&dMgr,
			true, // getSubdirectories
			includeSubDirCurrenDirOneDot,
			includeSubDirParentDirTwoDots,
			false,                       // getRegularFiles
			false,                       // getSymLinksFiles
			false,                       // getOtherNonRegularFiles
			subDirSelectCharacteristics, // subDirSelectCharacteristics
			FileSelectionCriteria{},     // fileSelectCriteria
			subDirsInDir,                // subDirsInDir
			nil,                         // filesInDir
			"dMgr",                      // targetDMgrLabel
			ePrefix.XCpy("dMgr"))

	return numOfSubdirectories, err
}

// GetSubDirsFilesInDirTree
//
// Returns selected subdirectories and/or files in an
// entire directory tree defined by the parent directory
// input parameter 'directoryPath'. This means that the
// parent directory and all subdirectories in the
// directory tree will be scanned to identify and return
// subdirectories and/or files matching the specified
// selection criteria. Selected subdirectories will be
// returned in a Directory Manager Collection while
// selected files will be returned in a File Manager
// Collection.
//
// The parent directory 'directoryPath' will NOT be
// returned with selected subdirectories unless input
// paramter 'includeParentDirectory' is set to 'true'.
//
// To qualify for selection and inclusion in the returned
// Directory and File Manager Collections, items residing
// in the 'directoryPath' target directory tree are
// divided into two classes, subdirectories and files.
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
// Directory Modification Date and Directory Mode.
// Directory Name selections can be based on pattern
// matches or regular expression matches.
//
// If both of these filter requirements are satisfied,
// the subdirectory will be added to, and returned by,
// the Directory Manager Collection passed as input
// parameter 'directoriesLocated'.
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
// target directory tree ('directoryPath') which are not
// classified as subdirectories.
//
// To be eligible for file selection, the os.FileInfo
// entry must first comply with the specified File Type
// criteria. In terms of File Type, files are classified
// as Regular Files, SymLink Files or Other Non-Regular
// Files.
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
// as device files, named pipes and sockets. Other
// Non-Regular Files is a catch-all category including
// all files which are NOT classified as Regular Files or
// SymLink Files.
//
// In addition to File Type, selected files must also
// comply with the File Characteristics Selection
// Criteria specified by input parameter,
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
//		'directoryPath'. The entire target directory
//		tree will be searched for eligible subdirectories
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
//	 	Directory Mode. Subdirectory Name selections can
//	 	be configured for pattern matches or regular
//	 	expression matches.
//
//	(3)	The parent directory ('directoryPath') will NOT
//		be returned with selected subdirectories unless
//		input parameters 'includeParentDirectory' and
//		'getSubdirectories' are set to 'true'.
//
//	(4)	Selected subdirectories will be returned to the
//		user in a Directory Manager Collection
//		('directoriesLocated').
//
//	(5)	Selected files are required to match two sets
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
//	(6)	Selected files will be returned to the user in a
//		File Manager Collection ('filesLocated').
//
//	(7) If the target directory identified by input
//		parameter 'directoryPath' contains NO subdirectories
//		or files matching the Type and Characteristics
//		selection criteria, this method will exit, no
//		subdirectories and/or files will be returned in
//		the Directory Manager or File Manager
//		Collections, and no error will be returned.
//
//	(8) If the target directory identified by input
//		parameter 'directoryPath' contains NO Files
//		whatsoever (0 Files), this method will exit, no
//		subdirectories or files will be added to the
//		Directory Manager or File Manager Collections,
//		and no error will be returned.
//
//	(9) If the target directory identified by input
//		parameter 'directoryPath' does NOT exist on an
//		attached storage drive, an error will be
//		returned.
//
//	(10) If input parameters 'getSubdirectories'
//		'getRegularFiles', 'getSymLinksFiles' and
//		'getOtherNonRegularFiles' are all set to
//		'false', an error will be returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	directoryPath				string
//
//		This string defines a directory path. This path
//		will be treated as the target directory tree for
//		the search operation designed to return
//		subdirectories and/or files. All subdirectories
//		and/or files which reside in this directory tree
//		and meet the specified selection criteria will
//		be returned to the user by means of the Directory
//		Manager ('directoriesLocated') and File Manager
//		('filesLocated') Collections.
//
//		If this directory path does not exist on an
//		attached storage drive, an error will be
//		returned.
//
//	getSubdirectories			bool
//
//		If this parameter is set to 'true', directory
//		entries which also meet the Directory Selection
//		Characteristics criteria (subDirSelectCharacteristics),
//		will be stored and returned in the Directory
//		Manager Collection passed as input parameter
//		'directoriesLocated'.
//
//		If input parameters 'getSubdirectories',
//		'getRegularFiles', 'getSymLinksFiles' and
//		'getOtherNonRegularFiles' are all set to 'false',
//		these parameters will be classified as conflicted
//		and an error will be returned.
//
//	includeParentDirectory				bool
//
//		If this parameter is set to 'true' and input parameter
//		'getSubdirectories' is also set to 'true', the parent
//		directory defined by input parameter 'directoryPath'
//		will be added to the returned Directory Manager
//		Collection 'directoriesLocated'.
//
//		If input parameter 'getSubdirectories' is set to
//		'false', 'includeParentDirectory' is ignored and
//		has no impact on method operations.
//
//		If this parameter is set to 'false' and input parameter
//		'getSubdirectories' is also set to 'true', the parent
//		directory defined by input parameter 'directoryPath'
//		WILL NOT be added to the returned Directory Manager
//		Collection 'directoriesLocated'.
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
//		input parameter 'directoriesLocated'.
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
//		input parameter 'directoriesLocated'.
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
//		When 'getSubdirectories' is set to 'true',
//		Directory os.FileIno entries matching this
//		selection criteria will be included in the
//		Directory Manager Collection returned by
//		parameter 'directoriesLocated'.
//
//		Remember that setting
//		'subDirSelectCharacteristics' to an empty
//		instance of FileSelectionCriteria will ensure
//		that all subdirectories are selected.
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
//		Characteristics Criteria specifications offered by
//		Type FileSelectionCriteria, see the documentation
//		for 'fileSelectCriteria', below.
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
//			If all of the file selection criterion in the FileSelectionCriteria
//			object are 'Inactive' or 'Not Set' (set to their zero or default values),
//			then all the files meeting the File Type requirements in the directory
//			defined by 'directoriesLocated' will be selected.
//
//				Example:
//				  fsc := FileSelectCriterionMode{}
//
//				  In this example, 'fsc' is NOT initialized. Therefore,
//				  all the selection criterion are 'Inactive'. Consequently,
//				  all the files meeting the File Type requirements in the
//				  directory defined by 'directoriesLocated' will be selected.
//
//			------------------------------------------------------------------------
//
//	directoriesLocated				*DirMgrCollection
//
//		A pointer to an instance of DirMgrCollection
//		which encapsulates an array of Directory
//		Manager (DirMgr) objects.
//
//		Information on subdirectories in the directory
//		defined by input parameter 'directoriesLocated' which
//		match the specified Type and Characteristics
//		Selection Criteria will be converted to Directory
//		Manager (DirMgr) objects and added to this
//		Directory Manager Collection (DirMgrCollection).
//
//		If input parameter 'getSubdirectories' is set to
//		'true' and this input parameter, 'directoriesLocated',
//		is set to 'nil', an error will be returned.
//
//		If input parameter 'getSubdirectories' is set to
//		'false', input parameter 'directoriesLocated' may be
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
//		input parameter 'directoriesLocated' which match the
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
//	directoriesLocatedLabel				string
//
//		The name or label associated with input parameter
//		'directoriesLocated' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "directoriesLocated" will be
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
//	numOfDirectoriesLocated		int
//
//		If this method completes successfully, this
//		parameter will return number of subdirectories
//		selected and returned in the Directory Manager
//		Collection 'directoriesLocated'.
//
// isParentDirectoryIncluded	bool
//
//		When set to 'true' this parameter signals that
//		the parent directory is included in
//		'numOfDirectoriesLocated' and the returned
//		Directory Manager Collection 'directoriesLocated'.
//
//	numOfFilesLocated			int
//
//		If this method completes successfully, this
//		parameter will return number of files selected
//		and returned in the File Manager Collection
//		'filesLocated'.
//
//	directoriesLocated			*DirMgrCollection
//
//		A pointer to an instance of DirMgrCollection
//		which encapsulates an array of Directory
//		Manager (DirMgr) objects.
//
//		Information on subdirectories in the directory
//		tree defined by input parameter 'directoryPath'
//		which match the specified Subdirectory
//		Characteristics Selection Criteria will be
//		returned by this Directory Manager Collection.
//
//	filesLocated				*FileMgrCollection
//
//		A pointer to an instance of FileMgrCollection
//		which encapsulates an array of File Manager
//		(FileMgr) objects.
//
//		Information on files in the directory tree
//		defined by input parameter 'directoryPath'
//		which match the specified File Type and File
//		Characteristics Selection Criteria will be
//		returned by this File Manager Collection.
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
func (dHlpr *DirHelper) GetSubDirsFilesInDirTree(
	directoryPath string,
	getSubdirectories bool,
	includeParentDirectory bool,
	includeSubDirCurrentDirOneDot bool,
	includeSubDirParentDirTwoDots bool,
	getRegularFiles bool,
	getSymLinksFiles bool,
	getOtherNonRegularFiles bool,
	subDirSelectCharacteristics FileSelectionCriteria,
	fileSelectCriteria FileSelectionCriteria,
	targetDMgrLabel string,
	errorPrefix interface{}) (
	numOfDirectoriesLocated int,
	isParentDirectoryIncluded bool,
	directoriesLocated *DirMgrCollection,
	numOfFilesLocated int,
	filesLocated *FileMgrCollection,
	err error) {

}
