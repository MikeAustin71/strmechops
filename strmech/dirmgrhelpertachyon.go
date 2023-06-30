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
func (dMgrHlprTachyon *dirMgrHelperTachyon) getDirectoryProfile(
	dMgr *DirMgr,
	includeSubDirCurrenDirOneDot bool,
	includeSubDirParentDirTwoDots bool,
	fileSelectCharacteristics FileSelectionCriteria,
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

	err = dirProfile.SetDirMgr(
		dMgr,
		dMgrLabel,
		ePrefix)

	if err != nil {

		return directoryPathDoesExist, dirProfile, err
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

	dirProfile.SubDirsIncludeCurrentDirOneDot =
		includeSubDirCurrenDirOneDot

	dirProfile.SubDirsIncludeParentDirTwoDot =
		includeSubDirParentDirTwoDots

	dirProfile.DirExistsOnStorageDrive =
		directoryPathDoesExist

	if !directoryPathDoesExist {

		return directoryPathDoesExist, dirProfile, err
	}

	var fileInfos []FileInfoPlus

	exitFInfo := func() {
		fileInfos = nil
	}

	defer exitFInfo()

	var lenFileInfos int
	var errs2 []error

	fileInfos,
		lenFileInfos,
		errs2,
		err = new(dirMgrHelperMolecule).
		lowLevelGetFileInfosFromDir(
			dMgr,
			true, // getDirectoryFileInfos
			includeSubDirCurrenDirOneDot,
			includeSubDirParentDirTwoDots,
			true,                      // getRegularFileInfos
			true,                      // getSymLinkFiles,
			true,                      // getOtherNonRegularFiles
			FileSelectionCriteria{},   // subdirectorySelectCharacteristics
			fileSelectCharacteristics, // fileSelectCharacteristics
			dMgrLabel,
			ePrefix.XCpy(dMgrLabel))

	if len(errs2) != 0 {

		if err != nil {
			errs2 = append(errs2, err)
		}

		err = new(StrMech).ConsolidateErrors(errs2)

		return directoryPathDoesExist, dirProfile, err
	}

	if err != nil {

		return directoryPathDoesExist, dirProfile, err
	}

	if lenFileInfos == 0 {

		return directoryPathDoesExist, dirProfile, err
	}

	for i := 0; i < lenFileInfos; i++ {

		if fileInfos[i].IsDir() {

			dirProfile.DirSubDirectories++
			dirProfile.DirSubDirectoriesBytes +=
				uint64(fileInfos[i].Size())

		} else if fileInfos[i].Mode().IsRegular() {

			dirProfile.DirTotalFiles++
			dirProfile.DirTotalFileBytes +=
				uint64(fileInfos[i].Size())

			dirProfile.DirRegularFiles++
			dirProfile.DirRegularFileBytes +=
				uint64(fileInfos[i].Size())

		} else if fileInfos[i].Mode()&os.ModeSymlink != 0 {

			dirProfile.DirTotalFiles++
			dirProfile.DirTotalFileBytes +=
				uint64(fileInfos[i].Size())

			dirProfile.DirSymLinkFiles++
			dirProfile.DirSymLinkFileBytes +=
				uint64(fileInfos[i].Size())

		} else {

			dirProfile.DirTotalFiles++
			dirProfile.DirTotalFileBytes +=
				uint64(fileInfos[i].Size())

			dirProfile.DirNonRegularFiles++
			dirProfile.DirNonRegularFileBytes +=
				uint64(fileInfos[i].Size())

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
			"The Total Number of files Processed = %v"+
			"%v Directory Path = '%v'\n",
			ePrefix.String(),
			checkTotalFiles,
			dirProfile.DirTotalFiles,
			dMgrLabel,
			dMgr.absolutePath)

		dirProfile.Errors = append(
			dirProfile.Errors,
			fmt.Errorf("%v",
				err.Error()))
	}

	return directoryPathDoesExist, dirProfile, err
}
