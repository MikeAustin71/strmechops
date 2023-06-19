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

// getFilesInDir
//
// This method scans a designated directory and returns
// information on selected files through a File Manager
// Collection passed as input parameter 'filesInDir'.
//
// Files selected for addition to the File Manager
// Collection must satisfy two sets of criteria, File
// Type and File Characteristics.
//
// To qualify for selection, the file must first comply
// with the specified File Type criteria. In terms of
// File Type, files are classified as directories,
// regular files, SymLink files or other non-regular
// files.
//
// Since this method does NOT process directories, the
// only valid File Types eligible for selection are
// regular files, SymLink files or other non-regular
// files.
//
// For an explanation of Regular and Non-Regular files,
// see the Definition of Terms section below.
//
// Screening criteria for File Type is controlled by the
// following three input parameters:
//
//	getRegularFiles - bool
//	getSymLinksFiles - bool
//	getOtherNonRegularFiles - bool
//
// In addition to File Type, selected files must comply
// with the second set of file selection criteria, File
// Characteristics. File Characteristics Selection
// Criteria is specified by input parameter,
// 'fileSelectCriteria'. This file selection criteria
// allows users to screen files for File Name, File
// Modification Date and File Mode.
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
//		files in the directory specified by input
//		parameter 'targetDMgr'. No subdirectories will be
//		searched for eligible files. Only the top level
//		or parent directory identified by 'targetDMgr'
//		will be searched for eligible files.
//
//	(2)	The files to be selected are required to match
//		two sets of selection criteria, File Type
//		Selection Criteria and File Characteristics
//		Selection Criteria.
//
//	(3) File Type Selection Criteria specifications are
//		passed as input parameters 'getRegularFiles',
//		'getSymLinksFiles' and 'getOtherNonRegularFiles'.
//		For an explanation of Regular and Non-Regular
//		files, see the section on "Definition of Terms",
//		above.
//
//	(4) File Characteristics Selection Criteria are user
//		specified selection requirements passed as input
//		parameter 'fileSelectCriteria'. This file
//		selection criteria allows users to screen files
//		for File Name, File Modification Date and File
//		Mode.
//
//	(5) If the target directory identified by input
//		parameter 'targetDMgr' contains NO Files meeting
//		(1) the File Type Selection Criteria and (2) the
//		File Characteristics Selection Criteria, this
//		method will exit, no files will be added to the
//		'filesInDir' File Manager Collection and no error
//		will be returned.
//
//	(6) If the target directory identified by input
//		parameter 'targetDMgr' contains NO Files
//		whatsoever (0 Files), this method will exit, no
//		files will be added to the 'filesInDir' File
//		Manager Collection and no error will be returned.
//
//	(7)	This method will NOT return file information on
//		subdirectories.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	targetDMgr					*DirMgr
//
//		An instance of DirMgr which identifies the
//		target directory for a file search to return
//		information of selected files. Information on
//		files in this directory matching the specified
//		File Type and File Characteristics selection
//		criteria will be returned as File Manager objects
//		in a File Manager Collection passed as
//		'filesInDir'.
//
//	getRegularFiles				bool
//
//		If this parameter is set to 'true', Regular
//		Files, which also meet the File Characteristics
//		selection criteria ('fileSelectCriteria'), will
//		be included in the file information returned
//		through the File Manager Collection passed as
//		 input parameter 'filesInDir'.
//
//		Regular Files include text files, image files and
//		executable files.
//
//		For an explanation of Regular and Non-Regular
//		files, see the section on "Definition Of Terms",
//		above.
//
//		If input parameters 'getRegularFiles',
//		'getSymLinksFiles' and 'getOtherNonRegularFiles'
//		are all set to 'false', these parameters will be
//		classified as conflicted and an error will be
//		returned.
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
//		If input parameters 'getRegularFiles',
//		'getSymLinksFiles' and 'getOtherNonRegularFiles'
//		are all set to 'false', these parameters will be
//		classified as conflicted and an error will be
//		returned.
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
//		If input parameters 'getRegularFiles',
//		'getSymLinksFiles' and 'getOtherNonRegularFiles'
//		are all set to 'false', these parameters will be
//		classified as conflicted and an error will be
//		returned.
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
//		Modification Date and File Mode.
//
//		Files matching these selection criteria, and the
//		File Type filter, will be included in the file
//		information returned through the File Manager
//		Collection passed as input parameter
//		'filesInDir'.
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
//	numOfFilesLocated			uint64
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
func (dMgrHlprBoson *dirMgrHelperBoson) getFilesInDir(
	targetDMgr *DirMgr,
	getRegularFiles bool,
	getSymLinksFiles bool,
	getOtherNonRegularFiles bool,
	fileSelectCriteria FileSelectionCriteria,
	filesInDir *FileMgrCollection,
	targetDMgrLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	numOfFilesLocated uint64,
	err error) {

	if dMgrHlprBoson.lock == nil {
		dMgrHlprBoson.lock = new(sync.Mutex)
	}

	dMgrHlprBoson.lock.Lock()

	defer dMgrHlprBoson.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "dirMgrHelperBoson." +
		"getFilesInDir()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		return numOfFilesLocated, err
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

		return numOfFilesLocated, err
	}

	if getRegularFiles == false &&
		getSymLinksFiles == false &&
		getOtherNonRegularFiles == false {

		err = fmt.Errorf("%v\n"+
			"Fatal Error: File Type filters are conflicted!\n"+
			"All of the File Type filters are set to 'false'\n"+
			"This gurantees that NO files will be selected.\n"+
			"getRegularFiles == false\n"+
			"getSymLinksFiles == false\n"+
			"getOtherNonRegularFiles == false\n",
			ePrefix.String())

		return numOfFilesLocated, err
	}

	var isAllFileTypesSelected = false

	if getRegularFiles == true &&
		getSymLinksFiles == true &&
		getOtherNonRegularFiles == true {

		isAllFileTypesSelected = true
	}

	isFileSelectionCriteriaActive :=
		fileSelectCriteria.IsSelectionCriteriaActive()

	osPathSepStr := string(os.PathSeparator)

	var err2 error
	var nameDirEntries []os.DirEntry

	nameDirEntries,
		err2 = os.ReadDir(targetDMgr.absolutePath)

	var osFInfo os.FileInfo

	lenDirInfos := len(nameDirEntries)

	if lenDirInfos == 0 {

		return numOfFilesLocated, err
	}

	var isFileTypeFilterMatch bool
	var isFileCharacteristicsMatch bool

	for i := 0; i < lenDirInfos; i++ {

		isFileTypeFilterMatch = false
		isFileCharacteristicsMatch = false

		osFInfo,
			err2 = nameDirEntries[i].Info()

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Error: Error Returned by nameDirEntry.Info().\n"+
				"The conversion of DirEntry to os.FileInfo Failed."+
				"%v= '%v'\n"+
				"FileName= '%v'\n"+
				"Index= %v"+
				"Error= \n%v\n",
				ePrefix.String(),
				targetDMgrLabel,
				targetDMgr.absolutePath,
				targetDMgr.absolutePath+osPathSepStr+nameDirEntries[i].Name(),
				i,
				err2.Error())

		}

	}

	return numOfFilesLocated, err
}
