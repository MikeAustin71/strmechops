package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"os"
	"strings"
	"sync"
)

/*
  This source code file contains type 'DirMgr' .

  The Source Repository for this source code file is :
    https://github.com/MikeAustin71/pathfileopsgo.git

  Dependencies:
  -------------

  Types 'DirMgr' depends on types, 'FileHelper' and 'FileMgr'
  which are contained in source	code files, 'filehelper.go'
  and 'filemanager.go' located in this directory.

*/

// DirMgr - This type and its associated methods are used to
// manage directories, directory trees and directory permissions.
//
// Dependencies:
//
// Type 'DirMgr' depend on types, 'FileHelper' and 'FileMgr'
// which are contained in source code files, 'filehelper.go'
// and 'filemanager.go' located in this directory.
type DirMgr struct {
	isInitialized                   bool
	originalPath                    string
	path                            string // Stored with no trailing path separator
	isPathPopulated                 bool
	doesPathExist                   bool
	parentPath                      string // Stored with no trailing path separator
	isParentPathPopulated           bool
	absolutePath                    string
	isAbsolutePathPopulated         bool
	doesAbsolutePathExist           bool
	isAbsolutePathDifferentFromPath bool
	directoryName                   string // Name of directory without parent path.
	volumeName                      string
	isVolumePopulated               bool
	actualDirFileInfo               FileInfoPlus
	lock                            *sync.Mutex // Used internally to ensure thread safe operations
}

// CopyDirectory
//
// Copies files from the directory identified by DirMgr
// to a target directory. The files to be copied are
// selected according to file selection criteria
// specified by input parameter, 'fileSelectCriteria'.
//
// The selected files are copied by a Copy IO operation.
// For information on the Copy IO procedure see
// FileHelper{}.CopyFileByIo() method and reference:
//
//	https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// NOTE: This method ONLY copies files from the current
// directory identified by 'DirMgr'. It does NOT copy
// files from subdirectories.
//
// If the target directory does not exist and files are
// located matching the file selection criteria, this
// method will attempt to create the target directory.
// However, if no files meet the file selection criteria
// as defined by input parameter,'fileSelectCriteria',
// this method will NOT attempt to create the target
// directory.
//
// This method is optimized to support the copy of large
// numbers of files.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	targetDMgr   DirMgr
//
//		An instance of 'DirMgr' initialized with the
//		directory path of the target directory to which
//		selected files will be copied. If the target
//		directory does not exist, this method will
//		attempt to create it.
//
//
//	fileSelectCriteria FileSelectionCriteria
//
//		This input parameter should be configured with
//		the desired file selection criteria. Files
//		matching this criteria will be copied to the
//		directory identified by input parameter,
//		'targetDir'.
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
//	copyEmptyDirectory			bool
//
//		If set to 'true' the target directory will be
//		created regardless of whether any files are
//		copied to that directory. Remember that files are
//		only copied to the target directory if they meet
//		file selection criteria specified by input
//		parameter 'fileSelectCriteria'.
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
func (dMgr *DirMgr) CopyDirectory(
	targetDMgr DirMgr,
	fileSelectCriteria FileSelectionCriteria,
	copyEmptyDirectory bool,
	errorPrefix interface{}) (
	dirCopyStats DirectoryCopyStats,
	errs []error) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	dMgrHlpr := dirMgrHelper{}

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgr.CopyDirectory()",
		"")

	if err != nil {

		errs = append(errs, err)

		return dirCopyStats, errs
	}

	dirCopyStats,
		errs = dMgrHlpr.copyDirectory(
		dMgr,
		&targetDMgr,
		fileSelectCriteria,
		copyEmptyDirectory,
		"dMgr",
		"targetDMgr",
		ePrefix)

	return dirCopyStats, errs
}

// CopyDirectoryTree
//
// Copies all selected files in the directory tree to a
// specified target directory tree. If the target
// directory tree does not exist, this method will
// attempt to create it. See the details of target
// directory tree creation under input parameter
// 'copyEmptyDirectories'.
//
// If input parameter 'copyEmptyDirectories' is set to
// 'true', the entire directory tree will be created and
// may contain empty directories. If set to false, target
// directory tree elements will only be created if files
// meet the selection criteria and are subsequently
// copied to those target directory tree paths.
//
// Files eligible for copy to the target directory tree
// are selected on the basis of file selection criteria
// specified by input parameter, 'fileSelectCriteria'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	targetDMgr					DirMgr
//
//		An instance of 'DirMgr' initialized with the
//		directory path of the target directory to which
//		selected files will be copied. If the target
//		directory does not exist, this method will
//		attempt to create it.
//
//	copyEmptyDirectories		bool
//
//		If a target directory tree path does not
//		previously exist, the default behavior is to
//		create that directory ONLY if files matching the
//		file selection criteria are identified for that
//		directory. If no files match the file selection
//		criteria, the default is to NOT create the target
//		directory path.
//
//		If the parameter 'copyEmptyDirectories' is set to
//		'true' all target directory tree paths will be
//		created regardless of whether files are copied to
//		those directories.
//
//	fileSelectCriteria FileSelectionCriteria
//
//		This input parameter should be configured with
//		the desired file selection criteria. Files
//		matching this criteria will be copied to the
//		directory identified by input parameter,
//		'targetDir'.
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
// ----------------------------------------------------------------
//
// # Return Values
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
func (dMgr *DirMgr) CopyDirectoryTree(
	targetDMgr DirMgr,
	copyEmptyDirectories bool,
	fileSelectCriteria FileSelectionCriteria,
	errorPrefix interface{}) (
	dTreeCopyStats DirTreeCopyStats,
	errs []error) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgr.CopyDirectoryTree()",
		"")

	if err != nil {

		errs = append(errs, err)

		return dTreeCopyStats, errs
	}

	dTreeCopyStats,
		errs = new(dirMgrHelperNanobot).
		copyDirectoryTree(
			dMgr,
			&targetDMgr,
			copyEmptyDirectories,
			false, // skipTopLevelDirectory
			fileSelectCriteria,
			"dMgr",
			"targetDMgr",
			ePrefix)

	return dTreeCopyStats,
		errs
}

// CopyIn
//
// Receives a pointer to an incoming DirMgr object as an
// input parameter and copies the values from the
// incoming object to the current DirMgr object. When the
// copy operation is completed, the current DirMgr object
// is a duplicate of the incoming DirMgr object.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dMgrIn						*DirMgr
//
//		A pointer to an instance of DirMgr. The internal
//		member data values will be copied to the
//		corresponding data values contained in the
//		current DirMgr instance.
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
//	error
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
func (dMgr *DirMgr) CopyIn(
	dMgrIn *DirMgr,
	errorPrefix interface{}) error {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	dMgrHlpr := dirMgrHelper{}

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgr."+
			"CopyIn()",
		"")

	if err != nil {
		return err
	}

	return dMgrHlpr.copyIn(
		dMgr,
		dMgrIn,
		ePrefix.XCpy(
			"dMgr<-dMgrIn"))
}

// CopyOut
//
// Makes a duplicate copy of the current DirMgr values
// and returns them in a new DirMgr object.
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//	DirMgr
//
//		If this method completes successfully, a new,
//		fully populated instance of DirMgr will be
//		returned. This instance represents a deep copy of
//		the current DirMgr instance.
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
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (dMgr *DirMgr) CopyOut(
	errorPrefix interface{}) (
	DirMgr,
	error) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgr."+
			"CopyOut()",
		"")

	if err != nil {
		return DirMgr{}, err
	}

	return new(dirMgrHelperAtom).copyOut(
		dMgr,
		ePrefix.XCpy("dMgr->"))
}

// CopySubDirectoryTree
//
// For the purposes of this method, the directory
// identified by the current 'DirMgr' instance is treated
// as the parent directory.
//
// This method copies all the files matching specified
// file selection criteria from the parent directory
// subdirectories to the target directory specified by
// input parameter, 'targetDir'. Files residing in the
// parent directory are NOT eligible for the copy
// operation. Only files residing in the subdirectory
// tree are eligible for the file search and copy
// operation.
//
// Copied files must match the selection criteria
// specified by input parameter 'fileSelectCriteria'.
//
// If the target directory does not exist, and valid
// matching files are identified for that directory, this
// method will attempt to create the target directory.
//
// Conversely, if no files matching the file selection
// criteria are found, that corresponding target
// directory will NOT be created.
//
// If empty directories should be copied to the target
// directory, input parameter 'copyEmptyDirectories' must
// be set to 'true'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	targetDMgr					DirMgr
//
//		A concrete instance of DirMgr. This instance
//		specifies the target directory to which the
//		contents of the current DirMgr directory tree
//		will be copied.
//
//	copyEmptyDirectories		bool
//
//		If this boolean parameter is set to 'true', any
//		empty directories in the current DirMgr instance
//		subdirectory tree will be copied as empty
//		directories to the target directory tree
//	 	designated by input parameter 'targetDMgr'.
//
//	fileSelectCriteria			FileSelectionCriteria
//
//		This input parameter should be configured with
//		the desired file selection criteria. Files
//		matching this criteria will be copied from the
//		parent directory identified by the current
//		instance of DirMgr to the target directory tree
//		specified by input parameter 'targetDMgr'.
//
//		If file 'fileSelectCriteria' is uninitialized
//		(FileSelectionCriteria{}), all files residing
//		in the subdirectory tree defined by the current
//		DirMgr parent directory will be copied to the
//		target directory tree.
//
//			type FileSelectionCriteria struct {
//			 FileNamePatterns    []string
//				An array of strings containing File Name Patterns
//
//			 FilesOlderThan      time.Time
//			 	Match files with older modification date times
//
//			 FilesNewerThan      time.Time
//			 	Match files with newer modification date times
//
//			 SelectByFileMode    FilePermissionConfig
//			 	Match file mode (os.FileMode).
//
//			 SelectCriterionModeFileSelectCriterionMode
//			 	Specifies 'AND' or 'OR' selection mode
//			}
//
//		The FileSelectionCriteria type allows for configuration of single or multiple file
//		selection criterion. The 'SelectCriterionMode' can be used to specify whether the
//		file must match all, or any one, of the active file selection criterion.
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
//		If all of the file selection criterion in the FileSelectionCriteria object
//		are 'Inactive' or 'Not Set' (set to their zero or default values), then all
//		the files processed in the DirMgr subdirectory tree will be selected and
//		copied to the target directory tree.
//
//			Example:
//			  fsc := FileSelectCriterionMode{}
//
//			  In this example, 'fsc' is NOT initialized. Therefore,
//			  all the selection criterion are 'Inactive'. Consequently,
//			  all the files encountered in the DirMgr subdirectory tree
//			  during the search operation will be selected and copied
//			  to the target directory tree.
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
//	error
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
func (dMgr *DirMgr) CopySubDirectoryTree(
	targetDMgr DirMgr,
	copyEmptyDirectories bool,
	fileSelectCriteria FileSelectionCriteria,
	errorPrefix interface{}) (
	dTreeCopyStats DirTreeCopyStats,
	errs []error) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgr."+
			"CopySubDirectoryTree()",
		"")

	if err != nil {

		errs = append(errs, err)

		return dTreeCopyStats, errs
	}

	dTreeCopyStats,
		errs = new(dirMgrHelperNanobot).
		copyDirectoryTree(
			dMgr,
			&targetDMgr,
			copyEmptyDirectories,
			true, // skipTopLevelDirectory
			fileSelectCriteria,
			"dMgr",
			"targetDMgr",
			ePrefix)

	return dTreeCopyStats, errs
}

// DeleteAll
//
// ----------------------------------------------------------------
//
// # WARNING
//
//	This method will remove the directory, and all files
//	in that directory, identified by the current DirMgr
//	instance. It will also delete all child directories
//	and files in the directory tree.
//
// ----------------------------------------------------------------
//
// # Usage Example
//
//	Run DeleteAll on Directory:
//		"../pathfilego/003_filehelper/testdestdir/destdir"
//
//	All files and all subdirectories will be deleted.
//
//	Only the parent path will remain:
//		"../pathfilego/003_filehelper/testdestdir"
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//	error
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
func (dMgr *DirMgr) DeleteAll(
	errorPrefix interface{}) error {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgr.DeleteAll()",
		"")

	if err != nil {
		return err
	}

	return new(dirMgrHelper).
		deleteDirectoryAll(
			dMgr,
			"dMgr",
			ePrefix)
}

// DeleteAllFilesInDir
//
// ----------------------------------------------------------------
//
// # WARNING
//
// This method deletes all the files in the directory
// identified by the current instance of DirMgr.
//
// ONLY files in the top level directory identified
// by the current instance of DirMgr are deleted.
// Subdirectories are NOT deleted and files in
// subdirectories are NOT deleted.
//
// Reference:
// https://stackoverflow.com/questions/33450980/golang-remove-all-contents-of-a-directory
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//	deleteDirStats				DeleteDirFilesStats
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
//	errs						[]error
//
//		An array of errors is returned. If the method
//		completes successfully with no errors, a
//		ZERO-length array is returned.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errPrefDto'.
//	 	The 'errPrefDto' text will be prefixed or
//	 	attached to the	beginning of the error message.
//
//		Remember, this error array may contain multiple
//		errors.
func (dMgr *DirMgr) DeleteAllFilesInDir(
	errorPrefix interface{}) (
	deleteDirStats DeleteDirFilesStats,
	errs []error) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FilePermissionConfig."+
			"GetEntryTypeComponent()",
		"")

	if err != nil {

		errs = append(errs, err)

		return deleteDirStats, errs
	}

	deleteDirStats,
		errs = new(dirMgrHelper).
		deleteAllFilesInDirectory(
			dMgr,
			"dMgr",
			ePrefix)

	return deleteDirStats, errs
}

// DeleteAllSubDirectories
//
// ----------------------------------------------------------------
//
// # WARNING
//
// The directory identified by the current DirMgr
// instance is treated as the parent directory. This
// method will then proceed to delete all directories and
// files which are subsidiary to this parent directory.
//
// The parent directory and all files residing in the
// parent directory identified by the current instance of
// DirMgr, WILL NOT BE DELETED.
//
// Essentially, all subdirectories which are subordinate
// to the DirMgr directory will be deleted along with
// their constituent files.
//
// ----------------------------------------------------------------
//
// # Usage Example
//
//	 Parent Directory:
//	  DirMgr = d:\parentdirectory
//	  files    d:\parentdirectory\file1.txt
//	           d:\parentdirectory\file2.txt
//
//	 Sub-Directories:
//	           d:\parentdirectory\dir01
//	           d:\parentdirectory\dir02
//	           d:\parentdirectory\dir03
//
//	After Executing DirMgr.DeleteAllSubDirectories() all
//	subdirectories and any files they contain will be
//	deleted. The only directory which remains is the
//	parent directory and any files contained within the
//	parent directory.
//
//	  DirMgr = d:\parentdirectory
//	  files    d:\parentdirectory\file1.txt
//	           d:\parentdirectory\file2.txt
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//	errs						[]error
//
//		An array of errors is returned. If the method
//		completes successfully with no errors, a
//		ZERO-length array is returned.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errPrefDto'.
//	 	The 'errPrefDto' text will be prefixed or
//	 	attached to the	beginning of the error message.
//
//		Remember, this error array may contain multiple
//		errors.
func (dMgr *DirMgr) DeleteAllSubDirectories(
	errorPrefix interface{}) (
	errs []error) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgr.CopyDirectory()",
		"")

	if err != nil {

		errs = append(errs, err)

		return errs
	}

	errs = new(dirMgrHelperMolecule).
		deleteAllSubDirectories(
			dMgr,
			"dMgr",
			ePrefix)

	return errs
}

// DeleteDirectoryTreeFiles
//
// ----------------------------------------------------------------
//
// # Warning
//
//	This method deletes files in the directory tree. This
//	means that files in the parent directory and subsidiary
//	directories, identified by the current DirMgr instance,
//	may be deleted depending on specified file selection
//	criteria.
//
// ----------------------------------------------------------------
//
// The parent directory for this tree is the directory
// specified by the current 'DirMgr' instance.
//
// Files eligible for deletion must match the file
// selection criteria specified by input parameter
// 'deleteFileSelectionCriteria'. The file deletion
// operation will search the parent directory ('DirMgr')
// and all subdirectories screening for files which match
// the file selection criteria.
//
// The file deletion operation is conducted in three steps:
//
//  1. The criteria for selecting files to be deleted is
//     created using input parameter
//     'deleteFileSelectionCriteria'.
//
//  2. A file search is conducted which includes the
//     DirMgr parent directory and all subdirectories in
//     the tree.
//
//  3. Files processed during the directory tree search
//     are compared to the file selection criteria
//     specified by 'deleteFileSelectionCriteria'. Those
//     files which match the selection criteria are then
//     deleted.
//
// This method is similar to method
// 'DirMgr.DeleteWalkDirFiles()'. However, this method
// returns less tracking data and is designed to work
// with very large numbers of files and directories.
//
// Note: As a result of this operation, files within
// directory tree folders may be deleted, but the folders or
// directory elements will NEVER be deleted.
//
// ------------------------------------------------------------------------------
//
// Input Parameters:
//
//	deleteFileSelectionCriteria	FileSelectionCriteria
//
//		This input parameter should be configured with
//		the desired file selection criteria. Files
//		matching this criteria will be deleted.
//
//		If file 'fileSelectCriteria' is uninitialized
//		(FileSelectionCriteria{}), all files within
//		the current DirMgr parent directory plus all
//		subdirectories WILL BE DELETED.
//
//			type FileSelectionCriteria struct {
//			 FileNamePatterns    []string
//				An array of strings containing File Name Patterns
//
//			 FilesOlderThan      time.Time
//			 	Match files with older modification date times
//
//			 FilesNewerThan      time.Time
//			 	Match files with newer modification date times
//
//			 SelectByFileMode    FilePermissionConfig
//			 	Match file mode (os.FileMode).
//
//			 SelectCriterionModeFileSelectCriterionMode
//			 	Specifies 'AND' or 'OR' selection mode
//			}
//
//		The FileSelectionCriteria type allows for configuration of single or multiple file
//		selection criterion. The 'SelectCriterionMode' can be used to specify whether the
//		file must match all, or any one, of the active file selection criterion.
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
//		If all of the file selection criterion in the FileSelectionCriteria object
//		are 'Inactive' or 'Not Set' (set to their zero or default values), then all
//		the files processed in the target directory tree will be selected and
//		deleted.
//
//			Example:
//			  fsc := FileSelectCriterionMode{}
//
//			  In this example, 'fsc' is NOT initialized. Therefore,
//			  all the selection criterion are 'Inactive'. Consequently,
//			  all the files encountered in the DirMgr directory tree
//			  during the search operation will be selected and deleted.
//			  This includes files all files in the parent DirMgr
//			  directory plus all files in the subdirectory tree.
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
//	deleteDirStats				DeleteDirFilesStats
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
//	errs						[]error
//
//		An array of errors is returned. If the method
//		completes successfully with no errors, a
//		ZERO-length array is returned.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errPrefDto'.
//	 	The 'errPrefDto' text will be prefixed or
//	 	attached to the	beginning of the error message.
//
//		Remember, this error array may contain multiple
//		errors.
func (dMgr *DirMgr) DeleteDirectoryTreeFiles(
	deleteFileSelectionCriteria FileSelectionCriteria,
	errorPrefix interface{}) (
	deleteDirStats DeleteDirFilesStats,
	errs []error) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgr.DeleteDirectoryTreeFiles()",
		"")

	if err != nil {

		errs = append(errs, err)

		return deleteDirStats, errs
	}

	dMgrHlpr := dirMgrHelper{}

	deleteDirStats,
		errs = dMgrHlpr.deleteDirectoryTreeStats(
		dMgr,
		deleteFileSelectionCriteria,
		false, // skip top level (parent) directory
		true,  // scan sub-directories
		"dMgr",
		"deleteFileSelectionCriteria",
		ePrefix)

	return deleteDirStats, errs
}

// DeleteFilesByNamePattern
//
// Receives a string defining a pattern to use in
// searching file names for all files in the directory
// identified by the current DirMgr instance.
//
// ----------------------------------------------------------------
//
// # WARNING
//
//	This method will delete files matching the specified
//	file search pattern.
//
//	If a file name matches the pattern specified by input
//	parameter, 'fileSearchPattern', it will be deleted.
//
// ----------------------------------------------------------------
//
// Only files in the directory identified by the current
// DirMgr instance will be subject to deletion. Files
// residing in subdirectories of the parent directory
// identified by the current of DirMgr instance WILL NOT
// BE DELETED or altered in any way.
//
// If the 'fileSearchPattern' is improperly formatted, an
// error will be returned.
//
// If the directory path identified by the current DirMgr
// instance does NOT exist, an error will be returned.
//
// ----------------------------------------------------------------
//
// # Example 'fileSearchPattern'
//
//	*.*              will match  all files in directory
//	*.html           will match  anyfilename.html
//	a*               will match  appleJack.txt
//	j????row.txt     will match  j1x34row.txt
//	data[0-9]*       will match  data123.csv
//
//	Reference For Matching Details:
//	  https://golang.org/pkg/path/filepath/#Match
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fileSearchPattern			string
//
//		This string contains the file search pattern
//		which will be used to identify files for deletion
//		in the directory identified by the current DirMgr
//		instance.
//
//		See the 'fileSearchPattern' example shown above.
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
//	deleteDirStats				DeleteDirFilesStats
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
func (dMgr *DirMgr) DeleteFilesByNamePattern(
	fileSearchPattern string,
	errorPrefix interface{}) (
	deleteDirStats DeleteDirFilesStats,
	errs []error) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgr.DeleteFilesByNamePattern()",
		"")

	if err != nil {

		errs = append(errs, err)

		return deleteDirStats, errs
	}

	deleteDirStats,
		errs = new(dirMgrHelper).
		deleteFilesByNamePattern(
			dMgr,
			fileSearchPattern,
			"dMgr",
			"fileSearchPattern",
			ePrefix)

	return deleteDirStats, errs
}

// DeleteFilesBySelectionCriteria
//
// ----------------------------------------------------------------
//
// # WARNING
//
//	This method deletes selected files from the parent
//	directory identified by the current instance of
//	DirMgr. Files residing in subdirectories of this
//	parent directory WILL NEVER BE DELETED.
//
// ----------------------------------------------------------------
//
// The directory specified by the current DirMgr instance
// is treated as the parent directory.
//
// Files in subdirectories are NOT DELETED. ONLY files in
// the directory identified by the current DirMgr
// instance are deleted.
//
// The file deletion operation consists of three steps:
//
//  1. The criteria for selecting files to be deleted is
//     created using input parameter
//     'deleteFileSelectionCriteria'.
//
//  2. A file search is conducted which is limited ONLY
//     to the DirMgr parent directory. Files in this
//     parent directory may be deleted if they match the
//     search criteria. On the hand, Files in the
//     subdirectory tree ARE NEVER DELETED.
//
//  3. Files processed during the directory search are
//     compared to the file selection criteria specified
//     by input parameter 'deleteFileSelectionCriteria'.
//     Those files which match this selection criteria
//     in the DirMgr parent directory WILL BE DELETED.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	deleteFileSelectionCriteria FileSelectionCriteria
//
//		This input parameter should be configured with
//		the desired file selection criteria. Files
//		matching this criteria will be deleted in the
//		parent directory identified by the current
//		instance of DirMgr.
//
//		If file 'fileSelectCriteria' is uninitialized
//		(FileSelectionCriteria{}), all files within
//		the current DirMgr parent directory WILL BE
//		DELETED.
//
//			type FileSelectionCriteria struct {
//			 FileNamePatterns    []string
//				An array of strings containing File Name Patterns
//
//			 FilesOlderThan      time.Time
//			 	Match files with older modification date times
//
//			 FilesNewerThan      time.Time
//			 	Match files with newer modification date times
//
//			 SelectByFileMode    FilePermissionConfig
//			 	Match file mode (os.FileMode).
//
//			 SelectCriterionModeFileSelectCriterionMode
//			 	Specifies 'AND' or 'OR' selection mode
//			}
//
//		The FileSelectionCriteria type allows for configuration of single or multiple file
//		selection criterion. The 'SelectCriterionMode' can be used to specify whether the
//		file must match all, or any one, of the active file selection criterion.
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
//		If all of the file selection criterion in the FileSelectionCriteria object
//		are 'Inactive' or 'Not Set' (set to their zero or default values), then all
//		the files processed in the DirMgr parent directory will be selected and
//		deleted.
//
//			Example:
//			  fsc := FileSelectCriterionMode{}
//
//			  In this example, 'fsc' is NOT initialized. Therefore,
//			  all the selection criterion are 'Inactive'. Consequently,
//			  all the files encountered in the DirMgr parent directory
//			  during the search operation will be selected for deletion.
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
//	deleteDirStats				DeleteDirFilesStats
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
//	errs						[]error
//
//		An array of errors is returned. If the method
//		completes successfully with no errors, a
//		ZERO-length array is returned.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errPrefDto'.
//	 	The 'errPrefDto' text will be prefixed or
//	 	attached to the	beginning of the error message.
//
//		Remember, this error array may contain multiple
//		errors.
func (dMgr *DirMgr) DeleteFilesBySelectionCriteria(
	deleteFileSelectionCriteria FileSelectionCriteria,
	errorPrefix interface{}) (
	deleteDirStats DeleteDirFilesStats,
	errs []error) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgr.DeleteDirectoryTreeFiles() ",
		"")

	if err != nil {

		errs = append(errs, err)

		return deleteDirStats, errs
	}

	dMgrHlpr := dirMgrHelper{}

	deleteDirStats,
		errs = dMgrHlpr.deleteDirectoryTreeStats(
		dMgr,
		deleteFileSelectionCriteria,
		false, // skip top level (parent) directory
		false, //scan sub-directories
		"dMgr",
		"deleteFileSelectionCriteria",
		ePrefix)

	return deleteDirStats, errs
}

// DeleteSubDirectoryTreeFiles
//
// Deletes subdirectory files. For this operation, the
// current 'DirMgr' is classified as the top level or
// parent directory. Files in this parent directory will
// NEVER BE DELETED.
//
// ----------------------------------------------------------------
//
// # WARNING
//
//	This method deletes files in subdirectories.
//
// ----------------------------------------------------------------
//
// Files eligible for deletion must match the file
// selection criteria specified by input parameter
// 'deleteFileSelectionCriteria'. The file deletion
// operation will exclude the parent directory ('DirMgr')
// and confine the file search to the subdirectories
// underneath the parent directory. The file search will
// screen for files which match the file selection
// criteria in the subdirectory tree.
//
// The file deletion operation is conducted in three steps:
//
//  1. The criteria for selecting files to be deleted is
//     created using input parameter
//     'deleteFileSelectionCriteria'.
//
//  2. A file search is conducted which excludes the
//     DirMgr parent directory and focuses exclusively
//     on all subdirectories in the tree.
//
//  3. Files processed during the subdirectory tree
//     search are compared to the file selection
//     criteria specified by
//     'deleteFileSelectionCriteria'. Those files which
//     match the selection criteria are then deleted.
//
// Note: As a result of this operation, files within
// subdirectory tree folders may be deleted, but the
// folders or directory elements themselves will NEVER be
// deleted.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	deleteFileSelectionCriteria FileSelectionCriteria
//
//		This input parameter should be configured with
//		the desired file selection criteria. Files
//		matching this criteria will be deleted in the
//		subdirectory tree specified by the current
//		DirMgr parent directory.
//
//		If file 'fileSelectCriteria' is uninitialized
//		(FileSelectionCriteria{}), all directories in the
//		subdirectory tree will be searched, and all files
//		within those directories WILL BE DELETED.
//
//			type FileSelectionCriteria struct {
//			 FileNamePatterns    []string
//				An array of strings containing File Name Patterns
//
//			 FilesOlderThan      time.Time
//			 	Match files with older modification date times
//
//			 FilesNewerThan      time.Time
//			 	Match files with newer modification date times
//
//			 SelectByFileMode    FilePermissionConfig
//			 	Match file mode (os.FileMode).
//
//			 SelectCriterionModeFileSelectCriterionMode
//			 	Specifies 'AND' or 'OR' selection mode
//			}
//
//		The FileSelectionCriteria type allows for configuration of single or multiple file
//		selection criterion. The 'SelectCriterionMode' can be used to specify whether the
//		file must match all, or any one, of the active file selection criterion.
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
//		If all of the file selection criterion in the FileSelectionCriteria object
//		are 'Inactive' or 'Not Set' (set to their zero or default values), then all
//		the files processed in the target subdirectory tree will be selected and
//		deleted.
//
//			Example:
//			  fsc := FileSelectCriterionMode{}
//
//			  In this example, 'fsc' is NOT initialized. Therefore,
//			  all the selection criterion are 'Inactive'. Consequently,
//			  all the files encountered in the target directory tree
//			  during the search operation will be selected deleted.
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
//	deleteDirStats DeleteDirFilesStats - Statistics generated by the delete operation
//	                                     performed on the current directory identified
//	                                     by DirMgr.
//
//	errs						[]error
//
//		An array of errors is returned. If the method
//		completes successfully with no errors, a
//		ZERO-length array is returned.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errPrefDto'.
//	 	The 'errPrefDto' text will be prefixed or
//	 	attached to the	beginning of the error message.
//
//		Remember, this error array may contain multiple
//		errors.
func (dMgr *DirMgr) DeleteSubDirectoryTreeFiles(
	deleteFileSelectionCriteria FileSelectionCriteria,
	errorPrefix interface{}) (
	deleteDirStats DeleteDirFilesStats,
	errs []error) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgr.DeleteSubDirectoryTreeFiles()",
		"")

	if err != nil {

		errs = append(errs, err)

		return deleteDirStats, errs
	}

	dMgrHlpr := dirMgrHelper{}

	deleteDirStats,
		errs = dMgrHlpr.deleteDirectoryTreeStats(
		dMgr,
		deleteFileSelectionCriteria,
		true, // skip top level (parent) directory
		true, // scan sub-directories
		"dMgr",
		"deleteFileSelectionCriteria",
		ePrefix)

	return deleteDirStats, errs
}

// DeleteWalkDirFiles
//
// ----------------------------------------------------------------
//
// # WARNING
//
// This method deletes files in the directory tree
// identified by the current instance of DirMgr. The
// directory specified by DirMgr is treated as the
// top level or parent directory.
//
// Files are selected for deletion in the directory
// tree using file selection criteria defined by input
// parameter 'deleteFileSelectionCriteria'.
//
// The file selection and deletion operation starts in,
// and includes, the parent directory and every
// subdirectory in the parent directory tree.
//
// ----------------------------------------------------------------
//
// This method searches for files residing in the
// directory tree identified by the current DirMgr object
// which is treated as the parent directory.
//
// Starting with the parent directory, this method 'walks
// the directory tree' locating all files in the
// directory tree which match the file selection criteria
// submitted as method input parameter,
// 'deleteFileSelectionCriteria'.
//
// This method will delete files in the entire directory tree including
// the parent directory and its subdirectory tree.
//
// If a file matches the File Selection Criteria, it is DELETED. By the
// way, if ALL the file selection criterion are set to zero values or
// 'Inactive', then ALL FILES IN THE DIRECTORY ARE DELETED!!!
//
// A record of file deletions is included in the returned DirectoryDeleteFileInfo
// structure (DirectoryDeleteFileInfo.DeletedFiles).
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	deleteFileSelectionCriteria		FileSelectionCriteria
//
//	  This input parameter should be configured with the desired file
//	  selection criteria. Files matching this criteria will be deleted.
//
//		If file 'fileSelectCriteria' is uninitialized
//		(FileSelectionCriteria{}), all files within
//		the current DirMgr parent directory plus all
//		subdirectories WILL BE DELETED.
//
//			type FileSelectionCriteria struct {
//			 FileNamePatterns    []string
//				An array of strings containing File Name Patterns
//
//			 FilesOlderThan      time.Time
//			 	Match files with older modification date times
//
//			 FilesNewerThan      time.Time
//			 	Match files with newer modification date times
//
//			 SelectByFileMode    FilePermissionConfig
//			 	Match file mode (os.FileMode).
//
//			 SelectCriterionModeFileSelectCriterionMode
//			 	Specifies 'AND' or 'OR' selection mode
//			}
//
//		The FileSelectionCriteria type allows for configuration of single or multiple file
//		selection criterion. The 'SelectCriterionMode' can be used to specify whether the
//		file must match all, or any one, of the active file selection criterion.
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
//		If all of the file selection criterion in the FileSelectionCriteria object
//		are 'Inactive' or 'Not Set' (set to their zero or default values), then all
//		the files processed in the DirMgr directory tree will be selected and
//		deleted.
//
//			Example:
//			  fsc := FileSelectCriterionMode{}
//
//			  In this example, 'fsc' is NOT initialized. Therefore,
//			  all the selection criterion are 'Inactive'. Consequently,
//			  all the files encountered in the DirMgr directory tree
//			  during the search operation will be selected and deleted.
//			  This includes files all files in the parent DirMgr
//			  directory plus all files in the subdirectory tree.
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
//	DirectoryDeleteFileInfo -
//
//	        type DirectoryDeleteFileInfo struct {
//	          StartPath                string
//	          dirMgrs                  []DirMgr
//	          FoundFiles               []FileWalkInfo
//	          ErrReturns               []string
//	          DeleteFileSelectCriteria FileSelectionCriteria
//	          DeletedFiles             []FileWalkInfo
//	        }
//
//	        If successful, files matching the file selection criteria
//	        specified in input parameter 'deleteFileSelectionCriteria'
//	        will be DELETED and returned in a 'DirectoryDeleteFileInfo'
//	        structure field, 'DirectoryDeleteFileInfo.DeletedFiles.'
//
//	        Note: It is a good idea to check the returned field
//	              DirectoryDeleteFileInfo.ErrReturns to determine if any
//	              system errors were encountered during file processing.
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
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (dMgr *DirMgr) DeleteWalkDirFiles(
	deleteFileSelectionCriteria FileSelectionCriteria,
	errorPrefix interface{}) (
	DirectoryDeleteFileInfo,
	error) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	var err error

	var deleteFilesInfo DirectoryDeleteFileInfo

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgr.DeleteWalkDirFiles()",
		"")

	if err != nil {

		return deleteFilesInfo, err
	}

	dMgrHlpr := dirMgrHelper{}

	var errs []error

	deleteFilesInfo,
		errs =
		dMgrHlpr.deleteDirectoryTreeInfo(
			dMgr,
			deleteFileSelectionCriteria,
			false, // skip top level directory
			true,  // scan sub-directories
			"dMgr",
			"deleteFileSelectionCriteria",
			ePrefix)

	if len(errs) > 0 {
		err = new(StrMech).ConsolidateErrors(errs)
	}

	return deleteFilesInfo, err
}

// DoesAbsolutePathExist - Performs two operations.
// First the method determine whether the directory
// path indicated by the DirMgr.absolutePath field
// actually does exist on disk and returns a 'true'
// or 'false' boolean value accordingly. In addition,
// it also updates the DirMgr field
// 'DirMgr.doesAbsolutePathExist'.
func (dMgr *DirMgr) DoesAbsolutePathExist() bool {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"DirMgr."+
			"DoesAbsolutePathExist()",
		"")

	var dirPathDoesExist bool

	dirPathDoesExist,
		_,
		err =
		new(dirMgrHelperAtom).
			doesDirectoryExist(
				dMgr,
				PreProcPathCode.None(),
				"dMgr",
				ePrefix)

	if err != nil {
		dirPathDoesExist = false
	}

	return dirPathDoesExist
}

// DoesDirectoryExist
//
// Returns two boolean values indicating whether the
// Directory path exists and if the Directory absolute
// path exists.
func (dMgr *DirMgr) DoesDirectoryExist() (
	doesPathExist,
	doesAbsolutePathExist bool) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"DirMgr."+
			"DoesDirectoryExist()",
		"")

	dirPathDoesExist := false

	dirPathDoesExist,
		_,
		err =
		new(dirMgrHelperAtom).
			doesDirectoryExist(
				dMgr,
				PreProcPathCode.None(),
				"dMgr",
				ePrefix)

	if err != nil {
		dirPathDoesExist = false
	}

	return dirPathDoesExist, dirPathDoesExist
}

// DoesPathExist - Performs two operations.
// First the method determine whether the directory
// path indicated by the DirMgr.path field actually
// does exist on disk and returns a 'true' or 'false'
// boolean value accordingly. In addition, it also
// updates the DirMgr field DirMgr.doesPathExist field.
func (dMgr *DirMgr) DoesPathExist() bool {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	dirPathDoesExist := false

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"DirMgr."+
			"DoesPathExist()",
		"")

	dirPathDoesExist,
		_,
		err =
		new(dirMgrHelperAtom).
			doesDirectoryExist(
				dMgr,
				PreProcPathCode.None(),
				"dMgr",
				ePrefix)

	if err != nil {
		dirPathDoesExist = false
	}

	return dirPathDoesExist
}

// DoesThisDirectoryExist - Returns a boolean value of true if the directory identified
// by the current DirMgr instance does in fact exist.
//
// If, during the process of verifying the existence of the current directory, an error
// is encountered it will be a non-path error. Non-Path errors are most commonly associated
// with 'access-denied' situations. However, there may be other reasons for triggering Non-Path
// errors.
//
// If a Non-Path error is encountered, an appropriate error message is returned along with
// a boolean value of 'false'.
func (dMgr *DirMgr) DoesThisDirectoryExist(
	errorPrefix interface{}) (
	directoryDoesExist bool,
	nonPathError error) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	directoryDoesExist = false

	ePrefix,
		nonPathError = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgr."+
			"DoesThisDirectoryExist()",
		"")

	if nonPathError != nil {
		return directoryDoesExist, nonPathError
	}

	directoryDoesExist,
		_,
		nonPathError =
		new(dirMgrHelperAtom).
			doesDirectoryExist(
				dMgr,
				PreProcPathCode.None(),
				"dMgr",
				ePrefix)

	return directoryDoesExist, nonPathError
}

// Empty - Returns all DirMgr field values to their uninitialized
// or original zero values.
func (dMgr *DirMgr) Empty() {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	_ = new(dirMgrHelperElectron).
		empty(dMgr, "DirMgr.Empty() ", nil)

}

// Equal
//
// Compares two DirMgr objects to determine if they are
// equal.
//
// The Directory Path comparisons are NOT case-sensitive.
func (dMgr *DirMgr) Equal(dMgr2 *DirMgr) bool {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	return new(dirMgrHelper).
		equal(dMgr, dMgr2)
}

// EqualAbsPaths
//
// This method compares the absolute paths for the
// current directory manager and the input directory
// manager ('incomingDirMgr').
//
// If the two absolute paths are equal, this method
// returns 'true'.
//
// If the two absolute paths are NOT equal, this method
// returns 'false'.
//
// The comparison is NOT case-sensitive. In other words,
// both paths are converted to lower case before making
// the comparison.
//
// If either the current DirMgr ('dMgr') or the input
// parameter 'incomingDirMgr' are uninitialized, a value
// of 'false' is returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	incomingDirMgr				*DirMgr
//
//		A pointer to an incoming instance of DirMgr. This
//		instance will be compared to the current instance
//		of DirMgr to determine if the absolute paths are
//		equivalent.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		This method compares the absolute paths contained
//		in the current instance of DirMgr and a second
//		instance of DirMgr passed as input parameter
//		'incomingDirMgr'. If the absolute paths for both
//		instances are equivalent, a boolean value of
//		'true' will be returned.
func (dMgr *DirMgr) EqualAbsPaths(
	incomingDirMgr *DirMgr) bool {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	isEqual := false

	isEqual = new(dirMgrHelper).equalAbsolutePaths(
		dMgr,
		incomingDirMgr)

	return isEqual
}

// EqualPaths - Compares two DirMgr objects to determine
// if their paths are equal. Both Directory Path and
// absolute path must be equivalent.
//
// If the compared paths are equal, the method returns 'true'.
// If the paths are NOT equal, the method returns 'false'.
// The comparisons are NOT case-sensitive. In other words, all paths
// are converted to lower case before making the comparisons.
//
// If either the current DirMgr ('dMgr') or the input parameter
// 'dMgr2' are uninitialized, a value of 'false' is returned.
func (dMgr *DirMgr) EqualPaths(dMgr2 *DirMgr) bool {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	dMgrHlpr := dirMgrHelper{}
	isEqual := false

	isEqual = dMgrHlpr.equalPaths(
		dMgr,
		dMgr2)

	return isEqual
}

// ExecuteDirectoryFileOps - Performs a file operation on specified 'selected' files
// in the current directory ONLY. This function does NOT perform operations on the
// sub directories (a.k.a. the directory tree).
//
// To perform file operations on the entire Directory Tree, see Function 'ExecuteDirectoryTreeOps()',
// above.
//
// The types of File Operations performed are generally classified as 'file copy' and
// 'file deletion' operations. The precise file operation applied is defined by the type,
// 'FileOperationCode' which provides a series of constants used to identify the specific file
// operation applied.
//
// Input parameter, 'fileOps' is an array of type 'FileOperationCode' elements. Multiple file
// operations can be applied to a single file. For instance, a 'copy source to destination'
// operation can be followed by a 'delete source file' operation.
//
// The 'selected' files are identified by input parameter 'fileSelectCriteria' of type
// 'FileSelectionCriteria'. This file selection criteria is compared against all files
// in the directory (NOT the Directory Tree) identified by the current 'DirMgr' instance.
// When a match is found, that file is treated as a 'selected' source file and designated
// file operations are performed on that file.
//
// The results or final output from file operations utilizes the final input parameter,
// 'targetBaseDir' of type DirMgr. File operations are applied to selected source files
// and generated output is created in the 'targetBaseDir'.  For example 'copy' or 'move'
// file operations will transfer source files to 'targetBaseDir'.
// ------------------------------------------------------------------------
//
// IMPORTANT:
//
// This method performs File Operations ONLY on the directory
// identified by the current DirMgr instance.
//
// ------------------------------------------------------------------------
//
// Input Parameters:
//
//	fileSelectCriteria FileSelectionCriteria
//
//	  This input parameter should be configured with the desired file
//	  selection criteria. Files matching this criteria will be identified
//	  as 'Selected Files'. The specified File Operations (fileOps) will be
//	  performed on these selected files.
//
//	  type FileSelectionCriteria struct {
//	     FileNamePatterns    []string     // An array of strings containing File Name Patterns
//	     FilesOlderThan      time.Time    // Match files with older modification date times
//	     FilesNewerThan      time.Time    // Match files with newer modification date times
//	     SelectByFileMode   FilePermissionConfig  // Match by file mode (os.FileMode).
//	     SelectCriterionMode FileSelectCriterionMode // Specifies 'AND' or 'OR' selection mode
//	  }
//
//	  The FileSelectionCriteria type allows for configuration of single or multiple file
//	  selection criterion. The 'SelectCriterionMode' can be used to specify whether the
//	  file must match all, or any one, of the active file selection criterion.
//
//	  Elements of the FileSelectionCriteria are described below:
//
//
//	  type FileSelectionCriteria struct {
//	   FileNamePatterns    []string// An array of strings containing File Name Patterns
//	   FilesOlderThan      time.Time// Match files with older modification date times
//	   FilesNewerThan      time.Time// Match files with newer modification date times
//	   SelectByFileMode    FilePermissionConfig  // Match file mode (os.FileMode).
//	   SelectCriterionModeFileSelectCriterionMode // Specifies 'AND' or 'OR' selection mode
//	  }
//
//	  The FileSelectionCriteria type allows for configuration of single or multiple file
//	  selection criterion. The 'SelectCriterionMode' can be used to specify whether the
//	  file must match all, or any one, of the active file selection criterion.
//
//	  Elements of the FileSelectionCriteria are described below:
//
//	  FileNamePatterns    []string  - An array of strings which may define one or more
//	                                  search patterns. If a file name matches any one of the
//	                                  search pattern strings, it is deemed to be a 'match'
//	                                  for the search pattern criterion.
//
//	                                    Example Patterns:
//	                                     FileNamePatterns = []string{"*.log"}
//	                                     FileNamePatterns = []string{"current*.txt"}
//	                                     FileNamePatterns = []string{"*.txt", "*.log"}
//
//	                                  If this string array has zero length or if
//	                                  all the strings are empty strings, then this
//	                                  file search criterion is considered 'Inactive'
//	                                  or 'Not Set'.
//
//
//	  FilesOlderThan      time.Time - This date time type is compared to file
//	                                  modification date times in order to determine
//	                                  whether the file is older than the 'FilesOlderThan'
//	                                  file selection criterion. If the file modification
//	                                  date time is older than the 'FilesOlderThan' date time,
//	                                  that file is considered a 'match'	for this file selection
//	                                  criterion.
//
//	                                  If the value of 'FilesOlderThan' is set to time zero,
//	                                  the default value for type time.Time{}, then this
//	                                  file selection criterion is considered to be 'Inactive'
//	                                  or 'Not Set'.
//
//	  FilesNewerThan      time.Time - This date time type is compared to the file
//	                                  modification date time in order to determine
//	                                  whether the file is newer than the 'FilesNewerThan'
//	                                  file selection criterion. If the file modification date time
//	                                  is newer than the 'FilesNewerThan' date time, that file is
//	                                  considered a 'match' for this file selection criterion.
//
//	                                  If the value of 'FilesNewerThan' is set to time zero,
//	                                  the default value for type time.Time{}, then this
//	                                  file selection criterion is considered to be 'Inactive'
//	                                  or 'Not Set'.
//
//	  SelectByFileMode  FilePermissionConfig -
//	                                  Type FilePermissionConfig encapsulates an os.FileMode. The file
//	                                  selection criterion allows for the selection of files by File Mode.
//	                                  File modes are compared to the value	of 'SelectByFileMode'. If the
//	                                  File Mode for a given file is equal to the value of 'SelectByFileMode',
//	                                  that file is considered to be a 'match' for this file selection
//	                                  criterion. Examples for setting SelectByFileMode are shown as follows:
//
//	                                       fsc := FileSelectionCriteria{}
//	                                       err = fsc.SelectByFileMode.SetByFileMode(os.FileMode(0666))
//	                                       err = fsc.SelectByFileMode.SetFileModeByTextCode("-r--r--r--")
//
//	  SelectCriterionMode FileSelectCriterionMode -
//	                                  This parameter selects the manner in which the file selection
//	                                  criteria above are applied in determining a 'match' for file
//	                                  selection purposes. 'SelectCriterionMode' may be set to one of
//	                                  two constant values:
//
//	                                  FileSelectMode.ANDSelect() - File selected if all active selection
//	                                    criteria are satisfied.
//
//	                                    If this constant value is specified for the file selection mode,
//	                                    then a given file will not be judged as 'selected' unless all
//	                                    the active selection criterion are satisfied. In other words, if
//	                                    three active search criterion are provided for 'FileNamePatterns',
//	                                    'FilesOlderThan' and 'FilesNewerThan', then a file will NOT be
//	                                    selected unless it has satisfied all three criterion in this example.
//
//	                                  FileSelectMode.ORSelect() - File selected if any active selection
//	                                    criterion is satisfied.
//
//	                                    If this constant value is specified for the file selection mode,
//	                                    then a given file will be selected if any one of the active file
//	                                    selection criterion is satisfied. In other words, if three active
//	                                    search criterion are provided for 'FileNamePatterns', 'FilesOlderThan'
//	                                    and 'FilesNewerThan', then a file will be selected if it satisfies any
//	                                    one of the three criterion in this example.
//
// ------------------------------------------------------------------------
//
// IMPORTANT:
//
// If all of the file selection criterion in the FileSelectionCriteria object are
// 'Inactive' or 'Not Set' (set to their zero or default values), then all
// the files processed in the directory tree will be selected and returned as
// 'Found Files'.
//
//	     Example:
//
//	           fsc := FileSelectCriterionMode{}
//
//	           In this example, 'fsc' is NOT initialized. Therefore,
//	           all the selection criterion are 'Inactive'. Consequently,
//	           all the files encountered in the target directory during
//	           the search operation will be selected and returned as
//	           'Found Files'.
//
//	fileOps []FileOperationCode
//
//		An array of file operations to be performed on
//		each selected file. Selected files are identified
//		by matching the file selection criteria specified
//		by input parameter, 'fileSelectCriteria'. See above.
//
//		The FileOperationCode type consists of the following
//		constants.
//
//		FileOperationCode(0).MoveSourceFileToDestinationFile() FileOperationCode = iota
//		  Moves the source file to the destination file and
//		  then deletes the original source file
//
//		FileOperationCode(0).DeleteDestinationFile()
//		  Deletes the Destination file if it exists
//
//		FileOperationCode(0).DeleteSourceFile()
//		  Deletes the Source file if it exists
//
//		FileOperationCode(0).DeleteSourceAndDestinationFiles
//		  Deletes both the Source and Destination files
//		  if they exist.
//
//		FileOperationCode(0).CopySourceToDestinationByHardLinkByIo()
//		  Copies the Source File to the Destination
//		  using two copy attempts. The first copy is
//		  by Hard Link. If the first copy attempt fails,
//		  a second copy attempt is initiated/ by creating
//		  a new file and copying the contents by 'io.Copy'.
//		  An error is returned only if both copy attempts
//		  fail. The source file is unaffected.
//
//		  See: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
//		FileOperationCode(0).CopySourceToDestinationByIoByHardLink()
//		  Copies the Source File to the Destination
//		  using two copy attempts. The first copy is
//		  by 'io.Copy' which creates a new file and copies
//		  the contents to the new file. If the first attempt
//		  fails, a second copy attempt is initiated using
//		  'copy by hard link'. An error is returned only
//		  if both copy attempts fail. The source file is
//		  unaffected.
//
//		  See: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
//		FileOperationCode(0).CopySourceToDestinationByHardLink()
//		  Copies the Source File to the Destination
//		  using one copy mode. The only copy attempt
//		  utilizes 'Copy by Hard Link'. If this fails
//		  an error is returned.  The source file is
//		  unaffected.
//
//		FileOperationCode(0).CopySourceToDestinationByIo()
//		  Copies the Source File to the Destination
//		  using only one copy mode. The only copy
//		  attempt is initiated using 'Copy by IO' or
//		  'io.Copy'.  If this fails an error is returned.
//		  The source file is unaffected.
//
//		FileOperationCode(0).CreateSourceDir()
//		  Creates the Source Directory
//
//		FileOperationCode(0).CreateSourceDirAndFile()
//		  Creates the Source Directory and File
//
//		FileOperationCode(0).CreateSourceFile()
//		  Creates the Source File
//
//		FileOperationCode(0).CreateDestinationDir()
//		  Creates the Destination Directory
//
//		FileOperationCode(0).CreateDestinationDirAndFile()
//		  Creates the Destination Directory and File
//
//		FileOperationCode(0).CreateDestinationFile()
//		  Creates the Destination File
//
// ------------------------------------------------------------------------
//
// Input parameters (continued)
//
//	targetBaseDir - The file selection criteria, 'fileSelectCriteria', and
//	                the File Operations, 'fileOps' are applied to files in
//	                the target base directory. This input parameter is of
//	                type 'DirMgr'.
//
// ------------------------------------------------------------------------
//
// Return Value:
//
//	errs						[]error
//
//		An array of errors is returned. If the method
//		completes successfully with no errors, a
//		ZERO-length array is returned.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errPrefDto'.
//	 	The 'errPrefDto' text will be prefixed or
//	 	attached to the	beginning of the error message.
//
//		Remember, this error array may contain multiple
//		errors.
func (dMgr *DirMgr) ExecuteDirectoryFileOps(
	fileSelectCriteria FileSelectionCriteria,
	fileOps []FileOperationCode,
	targetBaseDir DirMgr) (errs []error) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"DirMgr.ExecuteDirectoryFileOps()",
		"")

	if err != nil {

		errs = append(errs, err)

		return errs
	}

	dMgrHlpr := dirMgrHelper{}

	errs = dMgrHlpr.executeDirectoryFileOps(
		dMgr,
		fileSelectCriteria,
		fileOps,
		&targetBaseDir,
		"dMgr",
		"targetBaseDir",
		"fileSelectCriteria",
		"fileOps",
		ePrefix)

	return errs
}

// ExecuteDirectoryTreeOps - Performs File Operations on specified 'selected'
// files in the directory tree identified by the current 'DirMgr' instance.
// The 'DirMgr' path therefore serves as the parent directory for file operations
// performed on the directory tree.
//
// If you wish to perform File Operations ONLY on the current directory and
// NOT THE ENTIRE DIRECTORY TREE, see Function "ExecuteDirectoryFileOps()", above.
//
// The types of File Operations performed are generally classified as 'file copy'
// and 'file deletion' operations. The precise file operation applied is defined
// by the type, 'FileOperationCode' which provides a series of constants, or
// enumerations, used to identify the specific file operation applied. Input
// parameter, 'fileOps' is an array of type 'FileOperationCode' elements. Multiple
// file operations can be applied to a single file. For instance, a 'copy source to
// destination' operation can be followed by a 'delete source file' operation.
//
// The 'selected' files are identified by input parameter 'fileSelectCriteria' of
// type 'FileSelectionCriteria'. This file selection criteria is compared against
// all files in the current directory tree identified by the current 'DirMgr'
// instance. When a match is found, that file is treated as a 'selected' file and
// designated file operations are performed on that file.
//
// ------------------------------------------------------------------------
//
// IMPORTANT:
//
// This method performs File Operations on THE ENTIRE DIRECTORY
// TREE identified by this DirMgr instance.
//
// ------------------------------------------------------------------------
//
// Input Parameters:
//
//	fileSelectCriteria FileSelectionCriteria
//
//	  This input parameter should be configured with the desired file
//	  selection criteria. Files matching this criteria will be identified
//	  as 'Selected Files'. The specified File Operations (fileOps) will be
//	  performed on these selected files.
//
//	  type FileSelectionCriteria struct {
//	    FileNamePatterns     []string	    // An array of strings containing File Name Patterns
//	    FilesOlderThan       time.Time    // Match files with older modification date times
//	    FilesNewerThan       time.Time    // Match files with newer modification date times
//	    SelectByFileMode     FilePermissionConfig  // Match file mode (os.FileMode).
//	                                      //   is set to 'false'.
//	    SelectCriterionMode  FileSelectCriterionMode // Specifies 'AND' or 'OR' selection mode
//	  }
//
//	  The FileSelectionCriteria type allows for configuration of single or multiple file
//	  selection criterion. The 'SelectCriterionMode' can be used to specify whether the
//	  file must match all, or any one, of the active file selection criterion.
//
//	  Elements of the FileSelectionCriteria are described below:
//
//
//	  type FileSelectionCriteria struct {
//	   FileNamePatterns    []string// An array of strings containing File Name Patterns
//	   FilesOlderThan      time.Time// Match files with older modification date times
//	   FilesNewerThan      time.Time// Match files with newer modification date times
//	   SelectByFileMode    FilePermissionConfig  // Match file mode (os.FileMode).
//	   SelectCriterionModeFileSelectCriterionMode // Specifies 'AND' or 'OR' selection mode
//	  }
//
//	  The FileSelectionCriteria type allows for configuration of single or multiple file
//	  selection criterion. The 'SelectCriterionMode' can be used to specify whether the
//	  file must match all, or any one, of the active file selection criterion.
//
//	  Elements of the FileSelectionCriteria are described below:
//
//	  FileNamePatterns    []string  - An array of strings which may define one or more
//	                                  search patterns. If a file name matches any one of the
//	                                  search pattern strings, it is deemed to be a 'match'
//	                                  for the search pattern criterion.
//
//	                                    Example Patterns:
//	                                     FileNamePatterns = []string{"*.log"}
//	                                     FileNamePatterns = []string{"current*.txt"}
//	                                     FileNamePatterns = []string{"*.txt", "*.log"}
//
//	                                  If this string array has zero length or if
//	                                  all the strings are empty strings, then this
//	                                  file search criterion is considered 'Inactive'
//	                                  or 'Not Set'.
//
//
//	  FilesOlderThan      time.Time - This date time type is compared to file
//	                                  modification date times in order to determine
//	                                  whether the file is older than the 'FilesOlderThan'
//	                                  file selection criterion. If the file modification
//	                                  date time is older than the 'FilesOlderThan' date time,
//	                                  that file is considered a 'match'	for this file selection
//	                                  criterion.
//
//	                                  If the value of 'FilesOlderThan' is set to time zero,
//	                                  the default value for type time.Time{}, then this
//	                                  file selection criterion is considered to be 'Inactive'
//	                                  or 'Not Set'.
//
//	  FilesNewerThan      time.Time - This date time type is compared to the file
//	                                  modification date time in order to determine
//	                                  whether the file is newer than the 'FilesNewerThan'
//	                                  file selection criterion. If the file modification date time
//	                                  is newer than the 'FilesNewerThan' date time, that file is
//	                                  considered a 'match' for this file selection criterion.
//
//	                                  If the value of 'FilesNewerThan' is set to time zero,
//	                                  the default value for type time.Time{}, then this
//	                                  file selection criterion is considered to be 'Inactive'
//	                                  or 'Not Set'.
//
//	  SelectByFileMode  FilePermissionConfig -
//	                                  Type FilePermissionConfig encapsulates an os.FileMode. The file
//	                                  selection criterion allows for the selection of files by File Mode.
//	                                  File modes are compared to the value	of 'SelectByFileMode'. If the
//	                                  File Mode for a given file is equal to the value of 'SelectByFileMode',
//	                                  that file is considered to be a 'match' for this file selection
//	                                  criterion. Examples for setting SelectByFileMode are shown as follows:
//
//	                                       fsc := FileSelectionCriteria{}
//	                                       err = fsc.SelectByFileMode.SetByFileMode(os.FileMode(0666))
//	                                       err = fsc.SelectByFileMode.SetFileModeByTextCode("-r--r--r--")
//
//	  SelectCriterionMode FileSelectCriterionMode -
//	                                  This parameter selects the manner in which the file selection
//	                                  criteria above are applied in determining a 'match' for file
//	                                  selection purposes. 'SelectCriterionMode' may be set to one of
//	                                  two constant values:
//
//	                                  FileSelectMode.ANDSelect() - File selected if all active selection
//	                                    criteria are satisfied.
//
//	                                    If this constant value is specified for the file selection mode,
//	                                    then a given file will not be judged as 'selected' unless all
//	                                    the active selection criterion are satisfied. In other words, if
//	                                    three active search criterion are provided for 'FileNamePatterns',
//	                                    'FilesOlderThan' and 'FilesNewerThan', then a file will NOT be
//	                                    selected unless it has satisfied all three criterion in this example.
//
//	                                  FileSelectMode.ORSelect() - File selected if any active selection
//	                                    criterion is satisfied.
//
//	                                    If this constant value is specified for the file selection mode,
//	                                    then a given file will be selected if any one of the active file
//	                                    selection criterion is satisfied. In other words, if three active
//	                                    search criterion are provided for 'FileNamePatterns', 'FilesOlderThan'
//	                                    and 'FilesNewerThan', then a file will be selected if it satisfies any
//	                                    one of the three criterion in this example.
//
// ------------------------------------------------------------------------
//
// IMPORTANT:
//
// If all of the file selection criterion in the FileSelectionCriteria object are
// 'Inactive' or 'Not Set' (set to their zero or default values), then all
// the files processed in the directory tree will be selected and returned as
// 'Found Files'.
//
//	Example:
//	     fsc := FileSelectCriterionMode{}
//
//	     In this example, 'fsc' is NOT initialized. Therefore,
//	     all the selection criterion are 'Inactive'. Consequently,
//	     all the files encountered in the target directory during
//	     the search operation will be selected and returned as
//	     'Found Files'.
//
// ---------------------------------------------------------------------------
//
//	fileOps []FileOperationCode - An array of file operations to be performed
//	                              on each selected file. Selected files are
//	                              identified by matching the file selection
//	                              criteria specified by input parameter,
//	                              'fileSelectCriteria'. See above.
//
//	  The FileOperationCode type consists of the following
//	  constants.
//
//	  FileOperationCode(0).None()
//	    No Action
//
//	  FileOperationCode(0).MoveSourceFileToDestinationFile()
//	    Moves the source file to the destination file and
//	    then deletes the original source file
//
//	  FileOperationCode(0).DeleteDestinationFile()
//	    Deletes the Destination file if it exists
//
//	  FileOperationCode(0).DeleteSourceFile()
//	    Deletes the Source file if it exists
//
//	  FileOperationCode(0).DeleteSourceAndDestinationFiles
//	    Deletes both the Source and Destination files
//	    if they exist.
//
//	  FileOperationCode(0).CopySourceToDestinationByHardLinkByIo()
//	    Copies the Source File to the Destination
//	    using two copy attempts. The first copy is
//	    by Hard Link. If the first copy attempt fails,
//	    a second copy attempt is initiated/ by creating
//	    a new file and copying the contents by 'io.Copy'.
//	    An error is returned only if both copy attempts
//	    fail. The source file is unaffected.
//
//	    See: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
//	  FileOperationCode(0).CopySourceToDestinationByIoByHardLink()
//	    Copies the Source File to the Destination
//	    using two copy attempts. The first copy is
//	    by 'io.Copy' which creates a new file and copies
//	    the contents to the new file. If the first attempt
//	    fails, a second copy attempt is initiated using
//	    'copy by hard link'. An error is returned only
//	    if both copy attempts fail. The source file is
//	    unaffected.
//
//	    See: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
//	  FileOperationCode(0).CopySourceToDestinationByHardLink()
//	    Copies the Source File to the Destination
//	    using one copy mode. The only copy attempt
//	    utilizes 'Copy by Hard Link'. If this fails
//	    an error is returned.  The source file is
//	    unaffected.
//
//	  FileOperationCode(0).CopySourceToDestinationByIo()
//	    Copies the Source File to the Destination
//	    using only one copy mode. The only copy
//	    attempt is initiated using 'Copy by IO' or
//	    'io.Copy'.  If this fails an error is returned.
//	    The source file is unaffected.
//
//	  FileOperationCode(0).CreateSourceDir()
//	    Creates the Source Directory
//
//	  FileOperationCode(0).CreateSourceDirAndFile()
//	    Creates the Source Directory and File
//
//	  FileOperationCode(0).CreateSourceFile()
//	    Creates the Source File
//
//	  FileOperationCode(0).CreateDestinationDir()
//	    Creates the Destination Directory
//
//	  FileOperationCode(0).CreateDestinationDirAndFile()
//	    Creates the Destination Directory and File
//
//	  FileOperationCode(0).CreateDestinationFile()
//	    Creates the Destination File
//
// ------------------------------------------------------------------------
//
// Input parameters (continued)
//
// targetBaseDir -  The file selection criteria, 'fileSelectCriteria', and
//
//	the File Operations, 'fileOps' are applied to files in
//	the target base directory. This input parameter is of
//	type 'DirMgr'.
//
// ---------------------------------------------------------------------------
//
// Return Values:
//
//	errs						[]error
//
//		An array of errors is returned. If the method
//		completes successfully with no errors, a
//		ZERO-length array is returned.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errPrefDto'.
//	 	The 'errPrefDto' text will be prefixed or
//	 	attached to the	beginning of the error message.
//
//		Remember, this error array may contain multiple
//		errors.
func (dMgr *DirMgr) ExecuteDirectoryTreeOps(
	fileSelectCriteria FileSelectionCriteria,
	fileOps []FileOperationCode,
	targetBaseDir DirMgr) (errs []error) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"DirMgr.ExecuteDirectoryTreeOps()",
		"")

	if err != nil {

		errs = append(errs, err)

		return errs
	}

	dMgrHlpr := dirMgrHelper{}

	errs = dMgrHlpr.executeDirectoryTreeOps(
		dMgr,
		fileSelectCriteria,
		fileOps,
		&targetBaseDir,
		"dMgr",
		"targetBaseDir",
		"fileOps",
		ePrefix)

	return errs
}

// FindDirectoryTreeFiles
//
// This method returns file information on files residing
// in a specific directory tree identified by the current
// DirMgr instance. The directory identified by 'DirMgr'
// is treated as the parent directory for the search.
//
// In addition to file information, this method also
// returns data on the directory tree being searched,
// including the parent directory, and all subdirectories
// in the tree.
//
// This method 'walks the directory tree' locating all
// files in the directory tree which match the file
// selection criteria submitted as input parameter,
// 'fileSelectCriteria'.
//
// All directories including the top level parent
// directory ('DirMgr') are searched. This differs from
// method 'DirMgr.FindWalkSubDirFiles()' which only
// searches the subdirectory tree.
//
// If a file matches the File Selection Criteria, it is
// included in the returned field,
// 'DirectoryTreeInfo.FoundFiles'. If ALL the file
// selection criterion are set to zero values or
// 'Inactive', then ALL FILES in the directory are
// selected and returned in the field,
// 'DirectoryTreeInfo.FoundFiles'.
//
// All directories searched will be included in the
// returned collection 'DirectoryTreeInfo.Directories'.
// This returned 'DirectoryTreeInfo.Directories'
// collection will always include the top level parent
// directory identified by 'DirMgr'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fileSelectCriteria FileSelectionCriteria
//
//	  This input parameter should be configured with the
//	  desired file selection criteria. Files matching
//	  this criteria will be returned as 'Found Files'.
//
//		type FileSelectionCriteria struct {
//		  FileNamePatterns     []string     // An array of strings containing File Name Patterns
//		  FilesOlderThan       time.Time    // Match files with older modification date times
//		  FilesNewerThan       time.Time    // Match files with newer modification date times
//		  SelectByFileMode     FilePermissionConfig  // Match file mode (os.FileMode).
//		}
//
//		The FileSelectionCriteria type allows for configuration of single or multiple file
//		selection criterion. The 'SelectCriterionMode' can be used to specify whether the
//		file must match all, or any one, of the active file selection criterion.
//
//		Elements of the FileSelectionCriteria are described below:
//
//
//		type FileSelectionCriteria struct {
//		 FileNamePatterns    []string// An array of strings containing File Name Patterns
//		 FilesOlderThan      time.Time// Match files with older modification date times
//		 FilesNewerThan      time.Time// Match files with newer modification date times
//		 SelectByFileMode    FilePermissionConfig  // Match file mode (os.FileMode).
//		 SelectCriterionModeFileSelectCriterionMode // Specifies 'AND' or 'OR' selection mode
//		}
//
//		The FileSelectionCriteria type allows for configuration of single or multiple file
//		selection criterion. The 'SelectCriterionMode' can be used to specify whether the
//		file must match all, or any one, of the active file selection criterion.
//
//		Elements of the FileSelectionCriteria are described below:
//
//		FileNamePatterns    []string  - An array of strings which may define one or more
//		                                search patterns. If a file name matches any one of the
//		                                search pattern strings, it is deemed to be a 'match'
//		                                for the search pattern criterion.
//
//		                                  Example Patterns:
//		                                   FileNamePatterns = []string{"*.log"}
//		                                   FileNamePatterns = []string{"current*.txt"}
//		                                   FileNamePatterns = []string{"*.txt", "*.log"}
//
//		                                If this string array has zero length or if
//		                                all the strings are empty strings, then this
//		                                file search criterion is considered 'Inactive'
//		                                or 'Not Set'.
//
//
//		FilesOlderThan      time.Time - This date time type is compared to file
//		                                modification date times in order to determine
//		                                whether the file is older than the 'FilesOlderThan'
//		                                file selection criterion. If the file modification
//		                                date time is older than the 'FilesOlderThan' date time,
//		                                that file is considered a 'match'	for this file selection
//		                                criterion.
//
//		                                If the value of 'FilesOlderThan' is set to time zero,
//		                                the default value for type time.Time{}, then this
//		                                file selection criterion is considered to be 'Inactive'
//		                                or 'Not Set'.
//
//		FilesNewerThan      time.Time - This date time type is compared to the file
//		                                modification date time in order to determine
//		                                whether the file is newer than the 'FilesNewerThan'
//		                                file selection criterion. If the file modification date time
//		                                is newer than the 'FilesNewerThan' date time, that file is
//		                                considered a 'match' for this file selection criterion.
//
//		                                If the value of 'FilesNewerThan' is set to time zero,
//		                                the default value for type time.Time{}, then this
//		                                file selection criterion is considered to be 'Inactive'
//		                                or 'Not Set'.
//
//		SelectByFileMode  FilePermissionConfig -
//		                                Type FilePermissionConfig encapsulates an os.FileMode. The file
//		                                selection criterion allows for the selection of files by File Mode.
//		                                File modes are compared to the value	of 'SelectByFileMode'. If the
//		                                File Mode for a given file is equal to the value of 'SelectByFileMode',
//		                                that file is considered to be a 'match' for this file selection
//		                                criterion. Examples for setting SelectByFileMode are shown as follows:
//
//		                                     fsc := FileSelectionCriteria{}
//		                                     err = fsc.SelectByFileMode.SetByFileMode(os.FileMode(0666))
//		                                     err = fsc.SelectByFileMode.SetFileModeByTextCode("-r--r--r--")
//
//		SelectCriterionMode FileSelectCriterionMode -
//	                                  This parameter selects the manner in which the file selection
//	                                  criteria above are applied in determining a 'match' for file
//	                                  selection purposes. 'SelectCriterionMode' may be set to one of
//	                                  two constant values:
//
//	                                  FileSelectMode.ANDSelect() - File selected if all active selection
//	                                    criteria are satisfied.
//
//	                                    If this constant value is specified for the file selection mode,
//	                                    then a given file will not be judged as 'selected' unless all
//	                                    the active selection criterion are satisfied. In other words, if
//	                                    three active search criterion are provided for 'FileNamePatterns',
//	                                    'FilesOlderThan' and 'FilesNewerThan', then a file will NOT be
//	                                    selected unless it has satisfied all three criterion in this example.
//
//	                                  FileSelectMode.ORSelect() - File selected if any active selection
//	                                    criterion is satisfied.
//
//	                                    If this constant value is specified for the file selection mode,
//	                                    then a given file will be selected if any one of the active file
//	                                    selection criterion is satisfied. In other words, if three active
//	                                    search criterion are provided for 'FileNamePatterns', 'FilesOlderThan'
//	                                    and 'FilesNewerThan', then a file will be selected if it satisfies any
//	                                    one of the three criterion in this example.
//
//		------------------------------------------------------------------------
//
//		IMPORTANT:
//
//		If all of the file selection criterion in the FileSelectionCriteria object are
//		'Inactive' or 'Not Set' (set to their zero or default values), then all
//		the files processed in the directory tree will be selected and returned as
//		'Found Files'.
//
//			Example:
//			     fsc := FileSelectionCriteria{} // fsc is NOT initialized
//
//			     In this example, all the selection criterion are
//			     'Inactive' and therefore all the files encountered
//			     in the target directory will be selected and returned
//			     as 'Found Files'.
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
//	dTreeInfo					DirectoryTreeInfo
//
//		If this method completes successfully, without
//		errors, a fully populated instance of
//		DirectoryTreeInfo is returned.
//
//		This structure contains detailed information on
//		files found during the search operation.
//
//	          type DirectoryTreeInfo struct {
//	            StartPath             string
//					The starting path or directory for the file
//	                search.
//
//	            Directories           DirMgrCollection
//					Directory Managers found during directory tree
//					search.
//
//					This collection will ALWAYS return the parent
//					directory ('DirMgr') as the first entry in the
//					collection.
//
//	            FoundFiles            FileMgrCollection
//					Found Files matching file selection criteria
//
//	            ErrReturns            []error
//					Internal System errors encountered
//
//	            FileSelectCriteria    FileSelectionCriteria
//	            	The File Selection Criteria submitted as an
//					input parameter to this method.
//	           }
//
//	        If successful, files matching the file selection criteria input
//	        parameter shown above will be returned in a 'DirectoryTreeInfo'
//	        object. The field 'DirectoryTreeInfo.FoundFiles' contains information
//	        on all the files in the specified directory tree which match the file selection
//	        criteria.
//
//	        Note: It is a good idea to check the returned field 'DirectoryTreeInfo.ErrReturns'
//	              to determine if any internal system errors were encountered while processing
//	              the directory tree.
//
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
func (dMgr *DirMgr) FindDirectoryTreeFiles(
	fileSelectionCriteria FileSelectionCriteria,
	errorPrefix interface{}) (
	dTreeInfo DirectoryTreeInfo,
	errs []error) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgr."+
			"FindDirectoryTreeFiles()",
		"")

	if err != nil {

		errs = append(errs, err)

		return dTreeInfo, errs
	}

	dTreeInfo,
		errs = new(dirMgrHelper).findDirectoryTreeFiles(
		dMgr,
		fileSelectionCriteria,
		false,
		true,
		"dMgr",
		"fileSelectionCriteria",
		ePrefix)

	return dTreeInfo, errs
}

// FindFilesByNamePattern - searches the current directory using a name pattern file
// search criteria.
//
// Regardless of the search pattern used, this method will never return subdirectories
// of the target search directory.
//
// Again, the file search will always be limited to the directory identified by the
// current DirMgr instance. No subdirectories will be searched.
//
// If the 'fileSearchPattern' is an empty string or improperly formatted, an error
// will be returned.
//
// ------------------------------------------------------------------------
//
// # Input parameter
//
//	fileSearchPattern  string -  The fileSearchPattern is string containing
//	                             parameters used to select target files in
//	                             directory identified by the 'DirMgr' type.
//
//	                             Example 'fileSearchPattern' strings
//
//	                             *.*             will match all files in directory.
//	                             *.html          will match  anyfilename.html
//	                             a*              will match  appleJack.txt
//	                             j????row.txt    will match  j1x34row.txt
//	                             data[0-9]*      will match 	data123.csv
//
//	                             Reference For File Pattern Matching Details:
//	                               https://golang.org/pkg/path/filepath/#Match
//
// ---------------------------------------------------------------------------
//
// Return Values:
//
//	FileMgrCollection - If this method completes successfully without error, the
//	                    returned FileMgrCollection type will contain an array of
//	                    FileMgr types identifying each of the files matched by
//	                    input parameter, 'fileSearchPattern'.
//
//	error             - If this method completes successfully, this return value
//	                    will be set to 'nil'. Otherwise, a valid error message will
//	                    be encapsulated in the returned type 'error'.
func (dMgr *DirMgr) FindFilesByNamePattern(
	fileSearchPattern string) (
	FileMgrCollection, error) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	fileMgrCol := FileMgrCollection{}.New()

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"DirMgr."+
			"FindFilesByNamePattern()",
		"")

	if err != nil {
		return fileMgrCol, err
	}

	dMgrHlpr := dirMgrHelper{}

	fileMgrCol,
		err = dMgrHlpr.findFilesByNamePattern(
		dMgr,
		fileSearchPattern,
		"dMgr",
		"fileSearchPattern",
		ePrefix)

	return fileMgrCol, err
}

// FindFilesBySelectCriteria - Conducts a file search in the directory
// identified by the current DirMgr instance. The file search is limited
// to that directory ONLY. No subdirectories will be searched.
//
// Files matching the "FileSectionCriteria" instance passed as an input
// parameter will be used to screen available files. Any files matching
// the file selection criteria will be returned in a 'FileMgrCollection'.
//
// Only matched files will be returned. No subdirectory names will ever
// be included.
//
// The use of a 'FileSelectionCriteria' structure allows for very flexible
// and granular file searches.
//
// ------------------------------------------------------------------------
//
// Input Parameter:
//
//	fileSelectCriteria FileSelectionCriteria
//	  This input parameter should be configured with the desired file
//	  selection criteria. Files matching this criteria will be returned as
//	  'Found Files'.
//
//	  type FileSelectionCriteria struct {
//	   FileNamePatterns    []string// An array of strings containing File Name Patterns
//	   FilesOlderThan      time.Time// Match files with older modification date times
//	   FilesNewerThan      time.Time// Match files with newer modification date times
//	   SelectByFileMode    FilePermissionConfig  // Match file mode (os.FileMode).
//	   SelectCriterionModeFileSelectCriterionMode // Specifies 'AND' or 'OR' selection mode
//	  }
//
//	  The FileSelectionCriteria type allows for configuration of single or multiple file
//	  selection criterion. The 'SelectCriterionMode' can be used to specify whether the
//	  file must match all, or any one, of the active file selection criterion.
//
//	  Elements of the FileSelectionCriteria are described below:
//
//	  FileNamePatterns    []string  - An array of strings which may define one or more
//	                                  search patterns. If a file name matches any one of the
//	                                  search pattern strings, it is deemed to be a 'match'
//	                                  for the search pattern criterion.
//
//	                                    Example Patterns:
//	                                     FileNamePatterns = []string{"*.log"}
//	                                     FileNamePatterns = []string{"current*.txt"}
//	                                     FileNamePatterns = []string{"*.txt", "*.log"}
//
//	                                  If this string array has zero length or if
//	                                  all the strings are empty strings, then this
//	                                  file search criterion is considered 'Inactive'
//	                                  or 'Not Set'.
//
//
//	  FilesOlderThan      time.Time - This date time type is compared to file
//	                                  modification date times in order to determine
//	                                  whether the file is older than the 'FilesOlderThan'
//	                                  file selection criterion. If the file modification
//	                                  date time is older than the 'FilesOlderThan' date time,
//	                                  that file is considered a 'match'	for this file selection
//	                                  criterion.
//
//	                                  If the value of 'FilesOlderThan' is set to time zero,
//	                                  the default value for type time.Time{}, then this
//	                                  file selection criterion is considered to be 'Inactive'
//	                                  or 'Not Set'.
//
//	  FilesNewerThan      time.Time - This date time type is compared to the file
//	                                  modification date time in order to determine
//	                                  whether the file is newer than the 'FilesNewerThan'
//	                                  file selection criterion. If the file modification date time
//	                                  is newer than the 'FilesNewerThan' date time, that file is
//	                                  considered a 'match' for this file selection criterion.
//
//	                                  If the value of 'FilesNewerThan' is set to time zero,
//	                                  the default value for type time.Time{}, then this
//	                                  file selection criterion is considered to be 'Inactive'
//	                                  or 'Not Set'.
//
//	  SelectByFileMode  FilePermissionConfig -
//	                                  Type FilePermissionConfig encapsulates an os.FileMode. The file
//	                                  selection criterion allows for the selection of files by File Mode.
//	                                  File modes are compared to the value	of 'SelectByFileMode'. If the
//	                                  File Mode for a given file is equal to the value of 'SelectByFileMode',
//	                                  that file is considered to be a 'match' for this file selection
//	                                  criterion. Examples for setting SelectByFileMode are shown as follows:
//
//	                                       fsc := FileSelectionCriteria{}
//	                                       err = fsc.SelectByFileMode.SetByFileMode(os.FileMode(0666))
//	                                       err = fsc.SelectByFileMode.SetFileModeByTextCode("-r--r--r--")
//
//	  SelectCriterionMode FileSelectCriterionMode -
//	                                  This parameter selects the manner in which the file selection
//	                                  criteria above are applied in determining a 'match' for file
//	                                  selection purposes. 'SelectCriterionMode' may be set to one of
//	                                  two constant values:
//
//	                                  FileSelectMode.ANDSelect() - File selected if all active selection
//	                                    criteria are satisfied.
//
//	                                    If this constant value is specified for the file selection mode,
//	                                    then a given file will not be judged as 'selected' unless all
//	                                    the active selection criterion are satisfied. In other words, if
//	                                    three active search criterion are provided for 'FileNamePatterns',
//	                                    'FilesOlderThan' and 'FilesNewerThan', then a file will NOT be
//	                                    selected unless it has satisfied all three criterion in this example.
//
//	                                  FileSelectMode.ORSelect() - File selected if any active selection
//	                                    criterion is satisfied.
//
//	                                    If this constant value is specified for the file selection mode,
//	                                    then a given file will be selected if any one of the active file
//	                                    selection criterion is satisfied. In other words, if three active
//	                                    search criterion are provided for 'FileNamePatterns', 'FilesOlderThan'
//	                                    and 'FilesNewerThan', then a file will be selected if it satisfies any
//	                                    one of the three criterion in this example.
//
// ------------------------------------------------------------------------
//
// IMPORTANT:
//
// If all of the file selection criterion in the FileSelectionCriteria object are
// 'Inactive' or 'Not Set' (set to their zero or default values), then all
// the files processed in the directory tree will be selected and returned as
// 'Found Files'.
//
//	Example:
//	  fsc := FileSelectCriterionMode{}
//
//	  In this example, 'fsc' is NOT initialized. Therefore,
//	  all the selection criterion are 'Inactive'. Consequently,
//	  all the files encountered in the target directory during
//	  the search operation will be selected and returned as
//	  'Found Files'.
//
// ---------------------------------------------------------------------------
//
// Return Values:
//
//	FileMgrCollection - If this method completes successfully without error, the
//	                    returned FileMgrCollection type will contain an array of
//	                    FileMgr types identifying each of the files matched by
//	                    input parameter, 'fileSelectCriteria'.
//
//	error             - If this method completes successfully, this return value
//	                    will be set to 'nil'. Otherwise, a valid error message will
//	                    be encapsulated in the returned type 'error'.
func (dMgr *DirMgr) FindFilesBySelectCriteria(
	fileSelectCriteria FileSelectionCriteria) (FileMgrCollection, error) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	dTreeInfo := DirectoryTreeInfo{}

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"DirMgr."+
			"FindFilesBySelectCriteria()",
		"")

	if err != nil {
		return dTreeInfo.FoundFiles, err
	}

	dMgrHlpr := dirMgrHelper{}

	dTreeInfo,
		errs := dMgrHlpr.findDirectoryTreeFiles(
		dMgr,
		fileSelectCriteria,
		false, // skip top level directory
		false, // scan sub-directories
		"dMgr",
		"fileSelectCriteria",
		ePrefix)

	if len(errs) > 0 {
		err = new(StrMech).ConsolidateErrors(errs)
	}

	if err != nil {
		return FileMgrCollection{}, err
	}

	return dTreeInfo.FoundFiles, err
}

// FindWalkDirFiles - This method returns file information on files residing in a
// specific directory tree identified by the current DirMgr instance. The directory
// identified by 'DirMgr' is treated as the parent directory for the search.
//
// In addition to file information, this method also returns data on the directory tree
// being searched including the parent directory and all subdirectories in the tree.
//
// This method 'walks the directory tree' locating all files in the directory tree which
// match the file selection criteria submitted as input parameter, 'fileSelectCriteria'.
//
// All directories including the top level parent directory ('DirMgr') are searched. This
// differs from method 'DirMgr.FindWalkSubDirFiles()' which only searches the subdirectory
// tree.
//
// If a file matches the File Selection Criteria, it is included in the returned field,
// 'DirectoryTreeInfo.FoundFiles'. If ALL the file selection criterion are set to zero
// values or 'Inactive', then ALL FILES in the directory are selected and returned in
// the field, 'DirectoryTreeInfo.FoundFiles'.
//
// All directories searched will be included in the returned collection
// 'DirectoryTreeInfo.Directories'. This returned 'DirectoryTreeInfo.Directories'
// collection will always include the top level parent directory identified by 'DirMgr'.
//
// ------------------------------------------------------------------------
//
// Input Parameters:
//
//	fileSelectCriteria FileSelectionCriteria
//	  This input parameter should be configured with the desired file
//	  selection criteria. Files matching this criteria will be returned as
//	  'Found Files'.
//
//	  type FileSelectionCriteria struct {
//	    FileNamePatterns     []string     // An array of strings containing File Name Patterns
//	    FilesOlderThan       time.Time    // Match files with older modification date times
//	    FilesNewerThan       time.Time    // Match files with newer modification date times
//	    SelectByFileMode     FilePermissionConfig  // Match file mode (os.FileMode).
//	  }
//
//	  The FileSelectionCriteria type allows for configuration of single or multiple file
//	  selection criterion. The 'SelectCriterionMode' can be used to specify whether the
//	  file must match all, or any one, of the active file selection criterion.
//
//	  Elements of the FileSelectionCriteria are described below:
//
//
//	  type FileSelectionCriteria struct {
//	   FileNamePatterns    []string// An array of strings containing File Name Patterns
//	   FilesOlderThan      time.Time// Match files with older modification date times
//	   FilesNewerThan      time.Time// Match files with newer modification date times
//	   SelectByFileMode    FilePermissionConfig  // Match file mode (os.FileMode).
//	   SelectCriterionModeFileSelectCriterionMode // Specifies 'AND' or 'OR' selection mode
//	  }
//
//	  The FileSelectionCriteria type allows for configuration of single or multiple file
//	  selection criterion. The 'SelectCriterionMode' can be used to specify whether the
//	  file must match all, or any one, of the active file selection criterion.
//
//	  Elements of the FileSelectionCriteria are described below:
//
//	  FileNamePatterns    []string  - An array of strings which may define one or more
//	                                  search patterns. If a file name matches any one of the
//	                                  search pattern strings, it is deemed to be a 'match'
//	                                  for the search pattern criterion.
//
//	                                    Example Patterns:
//	                                     FileNamePatterns = []string{"*.log"}
//	                                     FileNamePatterns = []string{"current*.txt"}
//	                                     FileNamePatterns = []string{"*.txt", "*.log"}
//
//	                                  If this string array has zero length or if
//	                                  all the strings are empty strings, then this
//	                                  file search criterion is considered 'Inactive'
//	                                  or 'Not Set'.
//
//
//	  FilesOlderThan      time.Time - This date time type is compared to file
//	                                  modification date times in order to determine
//	                                  whether the file is older than the 'FilesOlderThan'
//	                                  file selection criterion. If the file modification
//	                                  date time is older than the 'FilesOlderThan' date time,
//	                                  that file is considered a 'match'	for this file selection
//	                                  criterion.
//
//	                                  If the value of 'FilesOlderThan' is set to time zero,
//	                                  the default value for type time.Time{}, then this
//	                                  file selection criterion is considered to be 'Inactive'
//	                                  or 'Not Set'.
//
//	  FilesNewerThan      time.Time - This date time type is compared to the file
//	                                  modification date time in order to determine
//	                                  whether the file is newer than the 'FilesNewerThan'
//	                                  file selection criterion. If the file modification date time
//	                                  is newer than the 'FilesNewerThan' date time, that file is
//	                                  considered a 'match' for this file selection criterion.
//
//	                                  If the value of 'FilesNewerThan' is set to time zero,
//	                                  the default value for type time.Time{}, then this
//	                                  file selection criterion is considered to be 'Inactive'
//	                                  or 'Not Set'.
//
//	  SelectByFileMode  FilePermissionConfig -
//	                                  Type FilePermissionConfig encapsulates an os.FileMode. The file
//	                                  selection criterion allows for the selection of files by File Mode.
//	                                  File modes are compared to the value	of 'SelectByFileMode'. If the
//	                                  File Mode for a given file is equal to the value of 'SelectByFileMode',
//	                                  that file is considered to be a 'match' for this file selection
//	                                  criterion. Examples for setting SelectByFileMode are shown as follows:
//
//	                                       fsc := FileSelectionCriteria{}
//	                                       err = fsc.SelectByFileMode.SetByFileMode(os.FileMode(0666))
//	                                       err = fsc.SelectByFileMode.SetFileModeByTextCode("-r--r--r--")
//
//	  SelectCriterionMode FileSelectCriterionMode -
//	                                  This parameter selects the manner in which the file selection
//	                                  criteria above are applied in determining a 'match' for file
//	                                  selection purposes. 'SelectCriterionMode' may be set to one of
//	                                  two constant values:
//
//	                                  FileSelectMode.ANDSelect() - File selected if all active selection
//	                                    criteria are satisfied.
//
//	                                    If this constant value is specified for the file selection mode,
//	                                    then a given file will not be judged as 'selected' unless all
//	                                    the active selection criterion are satisfied. In other words, if
//	                                    three active search criterion are provided for 'FileNamePatterns',
//	                                    'FilesOlderThan' and 'FilesNewerThan', then a file will NOT be
//	                                    selected unless it has satisfied all three criterion in this example.
//
//	                                  FileSelectMode.ORSelect() - File selected if any active selection
//	                                    criterion is satisfied.
//
//	                                    If this constant value is specified for the file selection mode,
//	                                    then a given file will be selected if any one of the active file
//	                                    selection criterion is satisfied. In other words, if three active
//	                                    search criterion are provided for 'FileNamePatterns', 'FilesOlderThan'
//	                                    and 'FilesNewerThan', then a file will be selected if it satisfies any
//	                                    one of the three criterion in this example.
//
// ------------------------------------------------------------------------
//
// IMPORTANT:
//
// If all of the file selection criterion in the FileSelectionCriteria object are
// 'Inactive' or 'Not Set' (set to their zero or default values), then all
// the files processed in the directory tree will be selected and returned as
// 'Found Files'.
//
//	Example:
//	     fsc := FileSelectionCriteria{} // fsc is NOT initialized
//
//	     In this example, all the selection criterion are
//	     'Inactive' and therefore all the files encountered
//	     in the target directory will be selected and returned
//	     as 'Found Files'.
//
// ------------------------------------------------------------------------
//
// Return Values:
//
//	DirectoryTreeInfo structure	-
//	        type DirectoryTreeInfo struct {
//	          StartPath             string                // The starting path or directory for the file search
//	          Directories           DirMgrCollection      // Directory Managers found during directory tree search.
//	                                                      // This collection will ALWAYS return the parent directory
//	                                                      // ('DirMgr') as the first entry in the collection.
//	          FoundFiles            FileMgrCollection     // Found Files matching file selection criteria
//	          ErrReturns            []error               // Internal System errors encountered
//	          FileSelectCriteria    FileSelectionCriteria // The File Selection Criteria submitted as an
//	                                                      // input parameter to this method.
//	        }
//
//	        If successful, files matching the file selection criteria input
//	        parameter shown above will be returned in a 'DirectoryTreeInfo'
//	        object. The field 'DirectoryTreeInfo.FoundFiles' contains information
//	        on all the files in the specified directory tree which match the file selection
//	        criteria.
//
//	        Note: It is a good idea to check the returned field 'DirectoryTreeInfo.ErrReturns'
//	              to determine if any internal system errors were encountered while processing
//	              the directory tree.
//
//	error - If a program execution error is encountered during processing, it will
//	        be returned as an 'error' type. Also, see the comment on
//	        'DirectoryTreeInfo.ErrReturns', above.
func (dMgr *DirMgr) FindWalkDirFiles(
	fileSelectCriteria FileSelectionCriteria) (DirectoryTreeInfo, error) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	dTreeInfo := DirectoryTreeInfo{}
	var errs []error

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"DirMgr."+
			"FindWalkDirFiles()",
		"")

	if err != nil {
		return dTreeInfo, err
	}

	dMgrHlpr := dirMgrHelper{}

	dTreeInfo,
		errs = dMgrHlpr.findDirectoryTreeFiles(
		dMgr,
		fileSelectCriteria,
		false, // skip top level directory
		true,  // scan sub-directories
		"dMgr",
		"fileSelectCriteria",
		ePrefix)

	if len(errs) > 0 {
		err = new(StrMech).ConsolidateErrors(errs)
	}

	return dTreeInfo, err
}

// FindWalkSubDirFiles - This method returns file information on files residing in a
// subdirectory tree identified by the current DirMgr instance. As such, this method
// will NOT search the top level directory, parent directory identified by the current
// DirMgr instance. However, all directories subsidiary to the parent directory ('DirMgr')
// will be searched.
//
// This method 'walks the directory tree' locating all files in the subdirectory tree which
// match the file selection criteria submitted as input parameter, 'fileSelectCriteria'.
//
// If a file matches the File Selection Criteria, it is included in the returned field,
// 'DirectoryTreeInfo.FoundFiles'. If ALL the file selection criterion are set to zero
// values or 'Inactive', then ALL FILES in the suc-directory tree are selected and returned
// in the field, 'DirectoryTreeInfo.FoundFiles'.
//
// All directories searched will be included in the returned collection 'DirectoryTreeInfo.Directories'.
// If the parent directory has NO subdirectories, this returned collection will be empty.
//
// ------------------------------------------------------------------------
//
// Input Parameters:
//
//	fileSelectCriteria FileSelectionCriteria
//	  This input parameter should be configured with the desired file
//	  selection criteria. Files matching this criteria will be returned as
//	  'Found Files'.
//
//	  type FileSelectionCriteria struct {
//	    FileNamePatterns     []string     // An array of strings containing File Name Patterns
//	    FilesOlderThan       time.Time    // Match files with older modification date times
//	    FilesNewerThan       time.Time    // Match files with newer modification date times
//	    SelectByFileMode     FilePermissionConfig  // Match file mode (os.FileMode).
//	  }
//
//	  The FileSelectionCriteria type allows for configuration of single or multiple file
//	  selection criterion. The 'SelectCriterionMode' can be used to specify whether the
//	  file must match all, or any one, of the active file selection criterion.
//
//	  Elements of the FileSelectionCriteria are described below:
//
//
//	  type FileSelectionCriteria struct {
//	   FileNamePatterns    []string// An array of strings containing File Name Patterns
//	   FilesOlderThan      time.Time// Match files with older modification date times
//	   FilesNewerThan      time.Time// Match files with newer modification date times
//	   SelectByFileMode    FilePermissionConfig  // Match file mode (os.FileMode).
//	   SelectCriterionModeFileSelectCriterionMode // Specifies 'AND' or 'OR' selection mode
//	  }
//
//	  The FileSelectionCriteria type allows for configuration of single or multiple file
//	  selection criterion. The 'SelectCriterionMode' can be used to specify whether the
//	  file must match all, or any one, of the active file selection criterion.
//
//	  Elements of the FileSelectionCriteria are described below:
//
//	  FileNamePatterns    []string  - An array of strings which may define one or more
//	                                  search patterns. If a file name matches any one of the
//	                                  search pattern strings, it is deemed to be a 'match'
//	                                  for the search pattern criterion.
//
//	                                    Example Patterns:
//	                                     FileNamePatterns = []string{"*.log"}
//	                                     FileNamePatterns = []string{"current*.txt"}
//	                                     FileNamePatterns = []string{"*.txt", "*.log"}
//
//	                                  If this string array has zero length or if
//	                                  all the strings are empty strings, then this
//	                                  file search criterion is considered 'Inactive'
//	                                  or 'Not Set'.
//
//
//	  FilesOlderThan      time.Time - This date time type is compared to file
//	                                  modification date times in order to determine
//	                                  whether the file is older than the 'FilesOlderThan'
//	                                  file selection criterion. If the file modification
//	                                  date time is older than the 'FilesOlderThan' date time,
//	                                  that file is considered a 'match'	for this file selection
//	                                  criterion.
//
//	                                  If the value of 'FilesOlderThan' is set to time zero,
//	                                  the default value for type time.Time{}, then this
//	                                  file selection criterion is considered to be 'Inactive'
//	                                  or 'Not Set'.
//
//	  FilesNewerThan      time.Time - This date time type is compared to the file
//	                                  modification date time in order to determine
//	                                  whether the file is newer than the 'FilesNewerThan'
//	                                  file selection criterion. If the file modification date time
//	                                  is newer than the 'FilesNewerThan' date time, that file is
//	                                  considered a 'match' for this file selection criterion.
//
//	                                  If the value of 'FilesNewerThan' is set to time zero,
//	                                  the default value for type time.Time{}, then this
//	                                  file selection criterion is considered to be 'Inactive'
//	                                  or 'Not Set'.
//
//	  SelectByFileMode  FilePermissionConfig -
//	                                  Type FilePermissionConfig encapsulates an os.FileMode. The file
//	                                  selection criterion allows for the selection of files by File Mode.
//	                                  File modes are compared to the value	of 'SelectByFileMode'. If the
//	                                  File Mode for a given file is equal to the value of 'SelectByFileMode',
//	                                  that file is considered to be a 'match' for this file selection
//	                                  criterion. Examples for setting SelectByFileMode are shown as follows:
//
//	                                       fsc := FileSelectionCriteria{}
//	                                       err = fsc.SelectByFileMode.SetByFileMode(os.FileMode(0666))
//	                                       err = fsc.SelectByFileMode.SetFileModeByTextCode("-r--r--r--")
//
//	  SelectCriterionMode FileSelectCriterionMode -
//	                                  This parameter selects the manner in which the file selection
//	                                  criteria above are applied in determining a 'match' for file
//	                                  selection purposes. 'SelectCriterionMode' may be set to one of
//	                                  two constant values:
//
//	                                  FileSelectMode.ANDSelect() - File selected if all active selection
//	                                    criteria are satisfied.
//
//	                                    If this constant value is specified for the file selection mode,
//	                                    then a given file will not be judged as 'selected' unless all
//	                                    the active selection criterion are satisfied. In other words, if
//	                                    three active search criterion are provided for 'FileNamePatterns',
//	                                    'FilesOlderThan' and 'FilesNewerThan', then a file will NOT be
//	                                    selected unless it has satisfied all three criterion in this example.
//
//	                                  FileSelectMode.ORSelect() - File selected if any active selection
//	                                    criterion is satisfied.
//
//	                                    If this constant value is specified for the file selection mode,
//	                                    then a given file will be selected if any one of the active file
//	                                    selection criterion is satisfied. In other words, if three active
//	                                    search criterion are provided for 'FileNamePatterns', 'FilesOlderThan'
//	                                    and 'FilesNewerThan', then a file will be selected if it satisfies any
//	                                    one of the three criterion in this example.
//
// ------------------------------------------------------------------------
//
// IMPORTANT:
//
// If all of the file selection criterion in the FileSelectionCriteria object are
// 'Inactive' or 'Not Set' (set to their zero or default values), then all
// the files processed in the directory tree will be selected and returned as
// 'Found Files'.
//
//	Example:
//	     fsc := FileSelectionCriteria{} // fsc is NOT initialized
//
//	     In this example, all the selection criterion are
//	     'Inactive' and therefore all the files encountered
//	     in the target directory will be selected and returned
//	     as 'Found Files'.
//
// ------------------------------------------------------------------------
//
// Return Values:
//
//	DirectoryTreeInfo structure	-
//	        type DirectoryTreeInfo struct {
//	          StartPath             string                // The starting path or directory for the file search
//	          Directories           DirMgrCollection      // Directory Managers found during directory tree
//	                                                      // search. Note: The top level parent directory will
//	                                                      // NEVER be included.
//	          FoundFiles            FileMgrCollection     // Found Files matching file selection criteria
//	          ErrReturns            []error               // Internal System errors encountered
//	          FileSelectCriteria    FileSelectionCriteria // The File Selection Criteria submitted as an
//	                                                      // input parameter to this method.
//	        }
//
//	        If successful, files matching the file selection criteria input
//	        parameter shown above will be returned in a 'DirectoryTreeInfo'
//	        object. The field 'DirectoryTreeInfo.FoundFiles' contains information
//	        on all the files in the specified subdirectory tree which match the file
//	        selection criteria.
//
//	        Note: It is a good idea to check the returned field 'DirectoryTreeInfo.ErrReturns'
//	              to determine if any internal system errors were encountered while processing
//	              the directory tree.
//
//	error - If a program execution error is encountered during processing, it will
//	        be returned as an 'error' type. Also, see the comment on
//	        'DirectoryTreeInfo.ErrReturns', above.
func (dMgr *DirMgr) FindWalkSubDirFiles(
	fileSelectCriteria FileSelectionCriteria) (
	dTreeInfo DirectoryTreeInfo,
	err error) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"DirMgr."+
			"FindWalkSubDirFiles()",
		"")

	if err != nil {
		return dTreeInfo, err
	}

	dMgrHlpr := dirMgrHelper{}
	var errs []error

	dTreeInfo,
		errs = dMgrHlpr.findDirectoryTreeFiles(
		dMgr,
		fileSelectCriteria,
		true, // skip top level directory
		true, // scan sub-directories
		"dMgr",
		"fileSelectCriteria",
		ePrefix)

	if len(errs) > 0 {
		err = new(StrMech).ConsolidateErrors(errs)
	}

	/*
	   dTreeInfo,
	     err = dMgrHlpr.findFilesWalkDirectoryTree(
	     dMgr,
	     fileSelectCriteria,
	     true, // skipTopLevelDirectory
	     ePrefix,
	     "dMgr")
	*/
	return dTreeInfo, err
}

// GetAbsolutePath - Returns a string containing the
// absolute path for the current Directory Manager
// instance. This string returned by this method
// will NOT have a trailing path separator.
//
// See companion method GetAbsolutePathLc() to
// acquire a lower case version of absolute path.
func (dMgr *DirMgr) GetAbsolutePath() string {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	absolutePath := ""

	_,
		_,
		err := new(dirMgrHelperAtom).
		doesDirectoryExist(
			dMgr,
			PreProcPathCode.None(),
			"",
			nil)

	if err != nil {
		absolutePath = ""
	} else {
		absolutePath = dMgr.absolutePath
	}

	return absolutePath
}

// GetAbsolutePathLc - Returns a string containing the
// low case version of the absolute path for the current
// Directory Manager instance.
//
// This string returned by this method will NOT have a
// trailing path separator. It will consist of all lower
// case characters.
//
// See the companion method GetAbsolutePath() to return
// an absolute path string with upper and lower case
// characters.
func (dMgr *DirMgr) GetAbsolutePathLc() string {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	absolutePath := ""

	_,
		_,
		err := new(dirMgrHelperAtom).
		doesDirectoryExist(
			dMgr,
			PreProcPathCode.None(),
			"",
			nil)

	if err != nil {
		absolutePath = ""
	} else {
		absolutePath = strings.ToLower(dMgr.absolutePath)
	}

	return absolutePath
}

// GetAbsolutePathElements - Returns all the directories and drive
// specifications as an array of strings.
//
// # Example
//
// Path = "D:\ADir\BDir\CDir\EDir"
//
// Returned pathElements string array:
//
//	pathElements[0] = "D:"
//	pathElements[1] = "ADir"
//	pathElements[2] = "BDir"
//	pathElements[3] = "CDir"
//	pathElements[4] = "DDir"
//	pathElements[4] = "EDir"
func (dMgr *DirMgr) GetAbsolutePathElements() (
	pathElements []string,
	err error) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"DirMgr."+
			"GetAbsolutePathElements()",
		"")

	if err != nil {

		return pathElements, err
	}

	pathElements = make([]string, 0, 50)
	dMgrHlpr := dirMgrHelper{}

	pathElements,
		err = dMgrHlpr.
		getAbsolutePathElements(
			dMgr,
			"",
			ePrefix.XCpy("dMgr"))

	return pathElements, err
}

// GetAbsolutePathWithSeparator - Returns the current
// DirMgr.absolutePath with a trailing os.PathSeparator
// character. The path string may consist of upper and
// lower case characters.
//
// See the companion method GetAbsolutePathWithSeparatorLc()
// which returns a path string consisting of all lower case
// characters.
func (dMgr *DirMgr) GetAbsolutePathWithSeparator() string {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	absolutePath := ""

	_,
		_,
		err := new(dirMgrHelperAtom).
		doesDirectoryExist(
			dMgr,
			PreProcPathCode.None(),
			"",
			nil)

	if err != nil {
		absolutePath = ""
	} else {
		absolutePath = dMgr.absolutePath
	}

	lPath := len(absolutePath)

	if lPath == 0 {
		return ""
	}

	if absolutePath[lPath-1] != os.PathSeparator {
		return absolutePath + string(os.PathSeparator)
	}

	return absolutePath
}

// GetAbsolutePathWithSeparatorLc - Returns the current
// DirMgr.absolutePath with a trailing os.PathSeparator
// character. The path string will consist of all lower
// case characters.
//
// See the companion method GetAbsolutePathWithSeparator()
// which returns a path string consisting of upper and lower
// case characters.
func (dMgr *DirMgr) GetAbsolutePathWithSeparatorLc() string {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	absolutePath := ""

	_,
		_,
		err := new(dirMgrHelperAtom).
		doesDirectoryExist(
			dMgr,
			PreProcPathCode.None(),
			"",
			nil)

	if err != nil {
		absolutePath = ""
	} else {
		absolutePath = strings.ToLower(dMgr.absolutePath)
	}

	lPath := len(absolutePath)

	if lPath == 0 {
		return ""
	}

	if absolutePath[lPath-1] != os.PathSeparator {
		return absolutePath + string(os.PathSeparator)
	}

	return absolutePath
}

// GetDirectoryStats - Returns the number of bytes in the current directory identified
// by the 'DirMgr' instance. This method only returns bytes in the current directory
func (dMgr *DirMgr) GetDirectoryStats(
	errorPrefix interface{}) (
	dirStats DirectoryStatsDto,
	errs []error) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgr."+
			"GetDirectoryStats()",
		"")

	if err != nil {

		errs = append(errs, err)

		return dirStats, errs
	}

	dirStats,
		errs = new(dirMgrHelper).findDirectoryTreeStats(
		dMgr,
		false,
		true,
		"dMgr",
		ePrefix)

	return dirStats, errs
}

// GetDirectoryTreeStats - Returns all the bytes in a directory tree.
// The parent directory for the search is identified by the current
// DirMgr instance.
func (dMgr *DirMgr) GetDirectoryTreeStats() (
	dirStats DirectoryStatsDto,
	errs []error) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"DirMgr."+
			"GetDirectoryTreeStats()",
		"")

	if err != nil {

		errs = append(errs, err)

		return dirStats, errs
	}

	dirStats,
		errs = new(dirMgrHelper).findDirectoryTreeStats(
		dMgr,
		false,
		true,
		"dMgr",
		ePrefix)

	return dirStats, errs
}

// GetDirectoryTree
//
// Returns a DirMgrCollection containing all the
// subdirectories in the path of the parent directory
// identified by the current DirMgr instance.
//
// The returned DirMgrCollection will always contain the
// parent directory at the top of the array (index=0).
// Therefore, if no errors are encountered, the returned
// DirMgrCollection will always consist of at least one
// directory.
//
// If subdirectories are found, then the returned
// DirMgrCollection will contain more than one directory.
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//	dirMgrs DirMgrCollection
//
//		If this method completes successfully, this
//		method will return an instance of DirMgrCollection
//		populated with an array of 'DirMgr' objects
//		identifying the parent directory and all
//		subdirectories specified by current instance of
//		'DirMgr'.
//
//			type DirMgrCollection struct {
//				dirMgrs []DirMgr
//			}
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
func (dMgr *DirMgr) GetDirectoryTree(
	errorPrefix interface{}) (
	dirMgrs DirMgrCollection,
	errs []error) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgr.GetDirectoryTree()",
		"")

	if err != nil {

		errs = append(errs, err)

		return dirMgrs, errs
	}

	dMgrHlpr := dirMgrHelper{}

	dirMgrs, errs =
		dMgrHlpr.getDirectoryTree(
			dMgr,
			"dMgr",
			ePrefix)

	return dirMgrs, errs
}

// GetDirectoryName - Returns a string containing the name
// of the directory without the parent path.
func (dMgr *DirMgr) GetDirectoryName() string {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	directoryName := ""

	_,
		_,
		_ = new(dirMgrHelperAtom).doesDirectoryExist(
		dMgr,
		PreProcPathCode.None(),
		"dMgr",
		nil)

	if !dMgr.isInitialized {

		directoryName = ""

	} else {

		directoryName = dMgr.directoryName

	}

	return directoryName
}

// GetFileInfoPlus - Returns a FileInfoPlus instance detailing file
// system information on the directory identified by the current
// Directory Manager instance.
func (dMgr *DirMgr) GetFileInfoPlus(
	errorPrefix interface{}) (
	FileInfoPlus,
	error) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	fileInfoPlus := FileInfoPlus{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgr.GetFileInfoPlus()",
		"")

	if err != nil {
		return fileInfoPlus, err
	}

	var dirDoesExist bool

	dirDoesExist,
		fileInfoPlus,
		err = new(dirMgrHelperAtom).
		doesDirectoryExist(
			dMgr,
			PreProcPathCode.None(),
			"dMgr",
			ePrefix)

	if err == nil && !dirDoesExist {

		fileInfoPlus = FileInfoPlus{}

		err = fmt.Errorf("%v\n"+
			"DirMgr Path DOES NOT EXIST!\n"+
			"DirMgr Path= '%v'\n",
			ePrefix.String(),
			dMgr.absolutePath)
	}

	return fileInfoPlus, err
}

// GetDirPermissionCodes - If the current directory exists on disk,
// this method will return the Directory Permission Codes encapsulated
// in a type 'FilePermissionConfig'.
//
// If the current Directory does NOT exist, this method will return an
// error.
func (dMgr *DirMgr) GetDirPermissionCodes(
	errorPrefix interface{}) (
	FilePermissionConfig,
	error) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	fileInfoPlus := FileInfoPlus{}
	var err error
	var dirDoesExist bool

	fPermCfg := FilePermissionConfig{}

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgr."+
			"GetDirPermissionCodes()",
		"")

	if err != nil {
		return fPermCfg, err
	}

	dirDoesExist,
		fileInfoPlus,
		err = new(dirMgrHelperAtom).
		doesDirectoryExist(
			dMgr,
			PreProcPathCode.None(),
			"dMgr",
			ePrefix)

	if err == nil && !dirDoesExist {

		err = fmt.Errorf("%v\n"+
			"DirMgr Path DOES NOT EXIST!\n"+
			"DirMgr Path='%v'\n",
			ePrefix.String(),
			dMgr.absolutePath)

	} else if err == nil && dirDoesExist {

		fPermCfg,
			err = new(FilePermissionConfig).
			NewByFileMode(
				fileInfoPlus.Mode(),
				ePrefix)
	}

	return fPermCfg, err
}

// GetNumberOfAbsPathElements - Returns the number of elements
// or path components in the absolute path of the current
// Directory Manager instance.
func (dMgr *DirMgr) GetNumberOfAbsPathElements(
	errorPrefix interface{}) (
	int,
	error) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgr."+
			"GetNumberOfAbsPathElements()",
		"")

	if err != nil {
		return -1, err
	}

	var pathElements []string

	pathElements,
		err = new(dirMgrHelper).
		getAbsolutePathElements(
			dMgr,
			"",
			ePrefix.XCpy(
				"dMgr"))

	return len(pathElements), err
}

// GetOriginalPath - Returns the original path used to initialize
// this Directory Manager instance.
func (dMgr *DirMgr) GetOriginalPath() string {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	originalPath := ""

	if !dMgr.isInitialized {

		originalPath = ""

	} else {

		originalPath = dMgr.originalPath

	}

	return originalPath
}

// GetParentDirMgr - Returns a new Directory Manager instance
// which represents the parent path for the current
// Directory Manager. The current Directory Manager absolute
// path is used in extracting the parent Directory Manager.
//
// Return Values:
//
//	dirMgrOut  DirMgr - If successful, this method returns a Directory Manager
//	                  which is a parent of the current Directory Manager.
//
//	hasParent  bool - If 'true', it signals that the current Directory Manager
//	                  has a valid parent directory. If 'false', it signals that
//	                  the current Directory Manager represents a top level directory
//	                  which has no parent directory. In that case a copy of the
//	                  current Directory will be returned.
//
//	err       error - If an error is encountered this error type will be populated
//	                  with an appropriate error message. Otherwise, a value of 'nil'
//	                  will be returned.
//
//	                  If 'hasParent' is 'false', no error will be returned.
func (dMgr *DirMgr) GetParentDirMgr(
	errorPrefix interface{}) (
	dirMgrParent DirMgr,
	hasParent bool,
	err error) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgr.GetParentDirMgr()",
		"")

	if err != nil {
		return dirMgrParent, hasParent, err
	}

	dMgrHlpr := dirMgrHelper{}

	dirMgrParent,
		hasParent,
		err = dMgrHlpr.getParentDirMgr(
		dMgr,
		"dMgr",
		ePrefix)

	return dirMgrParent, hasParent, err
}

// GetParentPath
//
// Returns a string containing the parent path for the
// current Directory Manager instance. The Parent Path
// string will NOT contain a trailing path separator.
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//	string
//
//		If this method completes successfully, this
//		method will return a string containing the
//	 	absolute value of the parent path for the
//	 	directory specified by the current instance of
//		Directory Manager (DirMgr).
//
//		Be advised that the returned path will NOT
//		contain a trailing path separator.
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
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (dMgr *DirMgr) GetParentPath(
	errorPrefix interface{}) (
	string,
	error) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	dMgrHlpr := dirMgrHelper{}
	dirMgrOut := DirMgr{}
	parentPath := ""

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgr.GetParentPath()",
		"")

	if err != nil {
		return parentPath, err
	}

	if !dMgr.isInitialized {

		err = fmt.Errorf("%v\n"+
			"Error: The current 'dMgr' instance is NOT initialized!\n",
			ePrefix)

		return parentPath, err
	}

	dirMgrOut,
		_,
		err = dMgrHlpr.getParentDirMgr(
		dMgr,
		"dMgr",
		ePrefix)

	if err == nil {

		parentPath = dirMgrOut.absolutePath
	}

	return parentPath, err
}

// GetPath - Returns the path used to configure this
// Directory Manager Instance. It will NOT contain a
// trailing path separator. It may or may not be an
// absolute path.
func (dMgr *DirMgr) GetPath() string {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	dPath := ""

	_,
		_,
		_ = new(dirMgrHelperAtom).
		doesDirectoryExist(
			dMgr,
			PreProcPathCode.None(),
			"dMgr",
			nil)

	if len(dMgr.path) == 0 ||
		!dMgr.isInitialized {
		dPath = ""
	} else {
		dPath = dMgr.path
	}

	return dPath
}

// GetPathWithSeparator - Returns the current
// DirMgr.absolutePath with a trailing os.PathSeparator
// character.
func (dMgr *DirMgr) GetPathWithSeparator() string {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	dPath := ""

	_,
		_,
		_ = new(dirMgrHelperAtom).
		doesDirectoryExist(
			dMgr,
			PreProcPathCode.None(),
			"dMgr",
			nil)

	if len(dMgr.path) == 0 ||
		!dMgr.isInitialized {

		dPath = ""

	} else {

		dPath = dMgr.path
	}

	lPath := len(dPath)

	if lPath == 0 {
		return ""
	}

	if dPath[lPath-1] != os.PathSeparator {
		return dPath + string(os.PathSeparator)
	}

	return dPath
}

// GetVolumeName - Returns a string containing the volume name
// of the directory identified by the current Directory Manager
// instance.
func (dMgr *DirMgr) GetVolumeName() string {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	volumeName := ""

	_,
		_,
		_ = new(dirMgrHelperAtom).
		doesDirectoryExist(
			dMgr,
			PreProcPathCode.None(),
			"dMgr",
			nil)

	if len(dMgr.volumeName) == 0 ||
		!dMgr.isInitialized {

		volumeName = ""

	} else {

		volumeName = dMgr.volumeName

	}

	return volumeName
}

// IsDirMgrValid - This method examines the current DirMgr object
// to determine whether it has been properly configured.
// If the current DirMgr object is valid, the method returns
// 'nil' for no errors.
//
// Otherwise, if the DirMgr object is INVALID, an error is
// returned.
func (dMgr *DirMgr) IsDirMgrValid(
	errorPrefix interface{}) error {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgr.IsDirMgrValid()",
		"")

	if err != nil {
		return err
	}

	return new(dirMgrHelper).isDirMgrValid(
		dMgr,
		ePrefix.XCpy("dMgr"))
}

// IsInitialized - Returns a boolean value indicating
// whether the Directory Manager instance is initialized.
func (dMgr *DirMgr) IsInitialized() bool {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	isInitialized := false

	_,
		_,
		_ = new(dirMgrHelperAtom).
		doesDirectoryExist(
			dMgr,
			PreProcPathCode.None(),
			"dMgr",
			nil)

	isInitialized = dMgr.isInitialized

	return isInitialized
}

// IsParentPathPopulated
//
// Returns a boolean value indicating whether the parent
// path for the current Directory Manager instance is
// populated.
func (dMgr *DirMgr) IsParentPathPopulated() bool {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	dMgrHlpr := dirMgrHelper{}
	isParentPathPopulated := false

	_,
		_,
		_ = dMgrHlpr.getParentDirMgr(
		dMgr,
		"dMgr",
		nil)

	if len(dMgr.parentPath) == 0 ||
		!dMgr.isInitialized {
		isParentPathPopulated = false
	} else {
		isParentPathPopulated = true
	}

	return isParentPathPopulated
}

// IsPathPopulated
//
// Returns a boolean value indicating whether the current
// Directory Manager path string is populated.
func (dMgr *DirMgr) IsPathPopulated() bool {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	dMgrHlpr := dirMgrHelper{}
	isDMgrPathPopulated := false
	var err error
	_,
		_,
		err = dMgrHlpr.getParentDirMgr(
		dMgr,
		"",
		nil)

	if len(dMgr.path) == 0 ||
		!dMgr.isInitialized ||
		err != nil {

		isDMgrPathPopulated = false

	} else {

		isDMgrPathPopulated = true
	}

	return isDMgrPathPopulated
}

// ParseValidPathStr
//
// Receives a valid path string and parses that string
// into is basic elements. Those elements are returned
// in a type ValidPathStrDto
func (dMgr *DirMgr) ParseValidPathStr(
	pathStr string,
	errorPrefix interface{}) (
	ValidPathStrDto,
	error) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()
	var ePrefix *ePref.ErrPrefixDto
	var err error
	validPathDto := ValidPathStrDto{}.New()

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FilePermissionConfig."+
			"GetEntryTypeComponent()",
		"")

	if err != nil {
		return validPathDto, err
	}

	validPathDto,
		err = new(dirMgrHelperMolecule).
		getValidPathStr(
			pathStr,
			"pathStr",
			ePrefix)

	return validPathDto, err
}

// IsVolumeNamePopulated - Returns a boolean value indicating
// whether the Volume Name for the current Directory Manager
// instance is populated.
func (dMgr *DirMgr) IsVolumeNamePopulated() bool {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	isDMgrVolumePopulated := false

	_,
		_,
		_ = new(dirMgrHelperAtom).
		doesDirectoryExist(
			dMgr,
			PreProcPathCode.None(),
			"dMgr",
			nil)

	if len(dMgr.volumeName) == 0 ||
		!dMgr.isInitialized {

		isDMgrVolumePopulated = false

	} else {

		isDMgrVolumePopulated = true

	}

	return isDMgrVolumePopulated
}

// MakeDirWithPermission
//
// If the directory path identified by the current DirMgr
// object does not exist, this method will create that
// directory path.
//
// The path will be created using permission
// specifications passed through input parameter
// 'fPermCfg'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fPermCfg					FilePermissionConfig
//
//		An instance of FilePermissionConfig containing
//		the permission specifications for the new
//		directory to be created from the directory path
//		specified by the current instance of DirMgr.
//
//		The easiest way to configure permissions is
//		to call FilePermissionConfig.New() with
//		a mode string ('modeStr').
//
//		The first character of the 'modeStr' designates the
//		'Entry Type'. Currently, only two 'Entry Type'
//		characters are supported. Therefore, the first
//		character in the 10-character input parameter
//		'modeStr' MUST be either a "-" indicating a file, or
//		a "d" indicating a directory.
//
//		The remaining nine characters in the 'modeStr'
//		represent unix permission bits and consist of three
//		group fields each containing 3-characters. Each
//		character in the three group fields may consist of
//		'r' (Read-Permission), 'w' (Write-Permission), 'x'
//		(Execute-Permission) or '-' signaling no permission or
//		no access allowed. A typical 'modeStr' authorizing
//		permission for full access to a file would be styled
//		as:
//
//		Directory Example: "drwxrwxrwx"
//
//		Groups: - Owner/User, Group, Other
//		From left to right
//		First Characters is Entry Type index 0 ("-")
//
//		First Char index 0 =     "-"   Designates a file
//
//		First Char index 0 =     "d"   Designates a directory
//
//		Char indexes 1-3 = Owner "rwx" Authorizing 'Read',
//	                                  Write' & Execute Permissions for 'Owner'
//
//		Char indexes 4-6 = Group "rwx" Authorizing 'Read', 'Write' & Execute
//	                                  Permissions for 'Group'
//
//		Char indexes 7-9 = Other "rwx" Authorizing 'Read', 'Write' & Execute
//	                                  Permissions for 'Other'
//
//	    -----------------------------------------------------
//	           Directory Mode String Permission Codes
//	    -----------------------------------------------------
//	      Directory
//			10-Character
//			 'modeStr'
//			 Symbolic		  Directory Access
//			  Format	   Permission Descriptions
//			----------------------------------------------------
//
//			d---------		no permissions
//			drwx------		read, write, & execute only for owner
//			drwxrwx---		read, write, & execute for owner and group
//			drwxrwxrwx		read, write, & execute for owner, group and others
//			d--x--x--x		execute
//			d-w--w--w-		write
//			d-wx-wx-wx		write & execute
//			dr--r--r--		read
//			dr-xr-xr-x		read & execute
//			drw-rw-rw-		read & write
//			drwxr-----		Owner can read, write, & execute. Group can only read;
//			                others have no permissions
//
//			Note: drwxrwxrwx - identifies permissions for directory
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
//	error
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
func (dMgr *DirMgr) MakeDirWithPermission(
	fPermCfg FilePermissionConfig,
	errorPrefix interface{}) error {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgr."+
			"MakeDirWithPermission()",
		"")

	if err != nil {
		return err
	}

	_,
		err = new(dirMgrHelperMolecule).
		lowLevelMakeDirWithPermission(
			dMgr,
			fPermCfg,
			"dMgr",
			ePrefix)

	return err
}

// MakeDir - If the directory path identified by the current DirMgr
// object does not exist, this method will create that directory path.
// The permission specification used to create the directory is
// 'drwxrwxrwx' which is equivalent to octal value, '020000000777'
//
// MakeDir creates a directory named path, along with any necessary
// parent directories.
//
// If the directory creation fails, an error is returned.
func (dMgr *DirMgr) MakeDir(
	errorPrefix interface{}) error {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgr.MakeDir()",
		"")

	if err != nil {
		return err
	}

	_,
		err = new(dirMgrHelperMolecule).
		lowLevelMakeDir(
			dMgr,
			"dMgr",
			ePrefix)

	// No errors = directory created.
	return err
}

// MoveDirectory
//
// Moves files from the source directory identified by
// DirMgr to a target directory. The 'move' operation is
// accomplished in three steps. First, the files to be
// copied are selected according to file selection
// criteria specified by input parameter,
// 'fileSelectCriteria'.
//
// Second, the selected files are copied to target
// directory identified by the input parameter,
// 'targetDir'. Finally, after verifying the copy, the
// files are deleted from the source directory (DirMgr).
//
// If, at the conclusion of the 'move' operation, there
// are no files or subdirectories remaining in the source
// directory (DirMgr), the source directory will be
// deleted.
//
// The selected files are copied using Copy IO operation.
// For information on the Copy IO procedure reference:
//
//	FileHelper{}.CopyFileByIo() method
//	https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// If the target directory does not previously exist,
// this method will attempt to create the target
// directory, provided, that files are selected for
// movement to that directory. If no files match the file
// selection criteria, the target directory will NOT be
// created.
//
// NOTE: This method ONLY moves files from the current
// directory identified by 'DirMgr'. It does NOT move
// files from subdirectories.
//
// This method is optimized to support the movement of
// large numbers of files.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete files in the current DirMgr
// path.  If all files have been moved out of the
// directory and there are no sub-Directories remaining,
// the DirMgr directory will likewise be deleted.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	targetDMgr					DirMgr
//
//		An instance of 'DirMgr' initialized with the
//		directory path of the target directory to which
//		selected files will be moved. If the target
//		directory does not exist, this method will
//		attempt to create it.
//
//	fileSelectCriteria			FileSelectionCriteria
//
//		This input parameter should be configured with
//		the desired file selection criteria. Files
//		matching this criteria will be moved from the
//		current DirMgr path to the directory identified
//		by input parameter, 'targetDir'.
//
//			type FileSelectionCriteria struct {
//			 FileNamePatterns    []string
//				An array of strings containing File Name Patterns
//
//			 FilesOlderThan      time.Time
//			 	Match files with older modification date times
//
//			 FilesNewerThan      time.Time
//			 	Match files with newer modification date times
//
//			 SelectByFileMode    FilePermissionConfig
//			 	Match file mode (os.FileMode).
//
//			 SelectCriterionModeFileSelectCriterionMode
//			 	Specifies 'AND' or 'OR' selection mode
//			}
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
//			FileNamePatterns	[]string
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
//		If all of the file selection criterion in the FileSelectionCriteria object are
//		'Inactive' or 'Not Set' (set to their zero or default values), then all the
//		files processed in the directory tree will be selected and moved
//		to the target directory.
//
//			Example:
//			  fsc := FileSelectCriterionMode{}
//
//			  In this example, 'fsc' is NOT initialized. Therefore,
//			  all the selection criterion are 'Inactive'. Consequently,
//			  all the files encountered in the target directory during
//			  the search operation will be selected and moved
//			  to the target directory.
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
//	dirMoveStats				DirectoryMoveStats
//
//		If this method completes successfully, this
//		structure will contain information and statistics
//		describing the outcome of the file 'move'
//		operation.
//
//		type DirectoryMoveStats struct {
//			TotalSrcFilesProcessed   uint64
//			SourceFilesMoved         uint64
//			SourceFileBytesMoved     uint64
//			SourceFilesRemaining     uint64
//			SourceFileBytesRemaining uint64
//			TotalDirsProcessed       uint64
//			DirsCreated              uint64
//			NumOfSubDirectories      uint64
//			SourceDirWasDeleted      bool
//			ComputeError             error
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
func (dMgr *DirMgr) MoveDirectory(
	targetDMgr DirMgr,
	fileSelectCriteria FileSelectionCriteria,
	errorPrefix interface{}) (
	dirMoveStats DirectoryMoveStats,
	errs []error) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgr.MoveDirectory()",
		"")

	if err != nil {

		errs = append(errs, err)

		return dirMoveStats, errs
	}

	dMgrHlpr := dirMgrHelper{}

	dirMoveStats,
		errs = dMgrHlpr.moveDirectory(
		dMgr,
		&targetDMgr,
		fileSelectCriteria,
		"dMgr",
		"targetDMgr",
		"fileSelectCriteria",
		ePrefix)

	return dirMoveStats, errs
}

// MoveDirectoryTree
//
// Moves all subdirectories and files plus files in the
// parent DirMgr directory to a target directory tree
// specified by input parameter 'targetDMgr'. The
// directory path specified by the current instance of
// DirMgr is therefore treated as the source directory
// tree.
//
// If successful, the parent directory DirMgr will be
// deleted along with the entire subdirectory tree. This
// directory tree will be copied or "moved" to the
// directory tree specified by input parameter
// 'targetDMgr'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete the entire directory tree
// identified by the current instance of DirMgr. This
// means that all files in this directory tree will also
// be deleted.
//
// --------------------------------------------------------------------
//
// Input Parameters:
//
//	targetDMgr					DirMgr
//
//		An instance of 'DirMgr' initialized with the
//		directory path of the target directory to which
//		all source files will be moved. If the target
//		directory does not exist, this method will
//		attempt to create it.
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
//	dirMoveStats				DirectoryMoveStats
//
//		If this method completes successfully, this
//		structure will contain information and statistics
//		describing the outcome of the file 'move'
//		operation.
//
//		type DirectoryMoveStats struct {
//			TotalSrcFilesProcessed   uint64
//			SourceFilesMoved         uint64
//			SourceFileBytesMoved     uint64
//			SourceFilesRemaining     uint64
//			SourceFileBytesRemaining uint64
//			TotalDirsProcessed       uint64
//			DirsCreated              uint64
//			NumOfSubDirectories      uint64
//			SourceDirWasDeleted      bool
//			ComputeError             error
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
func (dMgr *DirMgr) MoveDirectoryTree(
	targetDMgr DirMgr,
	errorPrefix interface{}) (
	dirMoveStats DirectoryMoveStats,
	errs []error) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgr.MoveDirectoryTree()",
		"")

	if err != nil {

		errs = append(errs, err)

		return dirMoveStats, errs
	}

	dMgrHlpr := dirMgrHelper{}

	dirMoveStats,
		errs = dMgrHlpr.moveDirectoryTree(
		dMgr,
		&targetDMgr,
		"dMgr",
		"targetDMgr",
		ePrefix)

	return dirMoveStats, errs
}

// MoveSubDirectoryTree - Moves all subdirectories and their constituent
// files from the source or parent directory 'DirMgr' to a target directory
// tree specified by input parameter 'targetDMgr'. If successful, all
// subdirectories and files in the source directory tree will be deleted.
// The source or parent directory identified by 'DirMgr' and the files
// within 'DirMgr' will NOT be deleted.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete the entire subdirectory tree in
// the parent directory identified by the current instance
// of DirMgr. However, the source or parent directory for
// the current instance of DirMgr, and its constituent
// files, will NOT be deleted.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	targetDMgr   DirMgr
//
//		An instance of 'DirMgr' initialized with the
//		directory path of the target directory to which
//		all source files will be moved.
//
//		If the target directory does not exist, this
//		method will attempt to create it.
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
//	dirMoveStats				DirectoryMoveStats
//
//		If this method completes successfully, this
//		structure will contain information and statistics
//		describing the outcome of the 'move' operation.
//
//		type DirectoryMoveStats struct {
//			TotalSrcFilesProcessed   uint64
//			SourceFilesMoved         uint64
//			SourceFileBytesMoved     uint64
//			SourceFilesRemaining     uint64
//			SourceFileBytesRemaining uint64
//			TotalDirsProcessed       uint64
//			DirsCreated              uint64
//			NumOfSubDirectories      uint64
//			SourceDirWasDeleted      bool
//			ComputeError             error
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
func (dMgr *DirMgr) MoveSubDirectoryTree(
	targetDMgr DirMgr,
	errorPrefix interface{}) (
	dirMoveStats DirectoryMoveStats,
	errs []error) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgr."+
			"MoveSubDirectoryTree()",
		"")

	if err != nil {

		errs = append(errs, err)

		return dirMoveStats, errs
	}

	dMgrHlpr := dirMgrHelper{}

	dirMoveStats,
		errs =
		dMgrHlpr.moveSubDirectoryTree(
			dMgr,
			&targetDMgr,
			"dMgr",
			"destinationDMgr",
			ePrefix)

	return dirMoveStats, errs
}

// New - Returns a new DirMgr object and populates the
// the data fields.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathStr						string
//
//		A path string designating a path or directory.
//
//		To reduce errors, the 'pathStr' should be
//		terminated with an appropriate path separator
//		('/' or '\').
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
//	DirMgr
//
//		If this method completes successfully, a fully
//		populated instance of DirMgr will be returned
//		configured using the path or path/filename
//		specified by input parameter 'pathStr'.
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
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
//
// ----------------------------------------------------------------
//
// # Usage
//
//		Example 'pathStr': "C:\dirA\dirB\dirC\"
//
//		Example Output After DirMgr Configuration:
//
//	    ----------------------------
//	    	DirMgr Fields
//	    ----------------------------
//
//	                     isInitialized:  true
//	                      originalPath:  D:\go\work\src\MikeAustin71\pathfilego\003_filehelper\logTest\testoverwrite
//	                              path:  D:\go\work\src\MikeAustin71\pathfilego\003_filehelper\logTest\testoverwrite
//	                   IsPathPopulated:  true
//	                     doesPathExist:  true
//	                        parentPath:  D:\go\work\src\MikeAustin71\pathfilego\003_filehelper\logTest
//	             isParentPathPopulated:  true
//	                      relativePath:  testoverwrite
//	           isRelativePathPopulated:  true
//	                      absolutePath:  D:\go\work\src\MikeAustin71\pathfilego\003_filehelper\logTest\testoverwrite
//	           isAbsolutePathPopulated:  true
//	   isAbsolutePathDifferentFromPath:  false
//	             doesAbsolutePathExist:  true
//	                     directoryName:  testoverwrite
//	                        volumeName:  D:
//	                 isVolumePopulated:  true
//	                 actualDirFileInfo:
//	                        ========== File Info Data ==========
//	                          File Info IsDir():  true
//	                           File Info Name():  testoverwrite
//	                           File Info Size():  0
//	                        File Info ModTime():  2018-01-06 Sat 00:06:56.421236800 -0600 CST
//	                           File Info Mode():  drwxrwxrwx
//	                            File Info Sys():  &{16 {617269082 30594119} {2388100752 30639796} {2388100752 30639796} 0 0}
//	                                   Dir path:  D:\go\work\src\MikeAustin71\pathfilego\003_filehelper\logTest\testoverwrite
func (dMgr *DirMgr) New(
	pathStr string,
	errorPrefix interface{}) (
	DirMgr,
	error) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	var newDirMgr DirMgr

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgr.New()",
		"")

	if err != nil {
		return newDirMgr, err
	}

	isEmpty,
		err := new(dirMgrHelperNanobot).
		setDirMgr(
			&newDirMgr,
			pathStr,
			"newDirMgr",
			"pathStr",
			ePrefix)

	if err != nil {
		return DirMgr{}, err
	}

	if isEmpty {

		return DirMgr{}, fmt.Errorf("%v"+
			"ERROR: dMgrHlpr.SetDirMgr(pathStr) returned an EMPTY DirMgr\n"+
			"pathStr='%v'\n",
			ePrefix.String(),
			pathStr)
	}

	return newDirMgr, err
}

// NewFromFileInfo - Returns a new DirMgr object based on two input parameters:
//   - A parent directory path string
//   - An os.FileInfo object containing the directory name.
func (dMgr *DirMgr) NewFromFileInfo(
	parentDirectoryPath string,
	info os.FileInfo,
	errorPrefix interface{}) (
	DirMgr,
	error) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	newDirMgr := DirMgr{}

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgr.NewFromFileInfo()",
		"")

	if err != nil {
		return newDirMgr, err
	}

	if info == nil {
		return newDirMgr,
			fmt.Errorf("%v\n"+
				"ERROR: Input parameter 'info' is 'nil' and INVALID!\n",
				ePrefix.String())
	}

	var isEmpty bool

	isEmpty,
		err = new(dirMgrHelper).
		setDirMgrWithPathDirectoryName(
			&newDirMgr,
			parentDirectoryPath,
			info.Name(),
			"newDirMgr",
			"parentDirectoryPath",
			"FileInfo.Name()",
			ePrefix)

	if err != nil {
		return DirMgr{}, err
	}

	if isEmpty {
		return DirMgr{},
			fmt.Errorf("%v\n"+
				"Newly generated 'DirMgr' is Empty!\n"+
				"dMgrHlpr.setDirMgrFromKnownPathDirName() returned an empty 'DirMgr'\n"+
				"parentDirectoryPath= '%v'\n"+
				"FileInfo.Name()= '%v'\n",
				ePrefix.String(),
				parentDirectoryPath,
				info.Name())
	}

	return newDirMgr, err
}

// NewFromDirMgrFileInfo
//
// Configures and returns a new 'DirMgr' instance based
// on two input parameters, 'directory' and 'fileNameExt'.
//
// Input parameter 'directory' is of type 'DirMgr' and is
// treated as the parent directory. The final directory
// name is provided by the input parameter 'fileInfo' of
// type 'os.FileInfo'.
func (dMgr *DirMgr) NewFromDirMgrFileInfo(
	parentDirectory DirMgr,
	directoryFileInfo os.FileInfo,
	errorPrefix interface{}) (
	DirMgr,
	error) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	newDirMgr := DirMgr{}

	var err error

	funcName := "DirMgr.NewFromDirMgrFileInfo()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {
		return newDirMgr, err
	}

	if directoryFileInfo == nil {
		return newDirMgr,
			fmt.Errorf("%v\n"+
				"ERROR: Input parameter 'directoryFileInfo' is 'nil' and therefore invalid!\n",
				ePrefix.String())
	}

	dMgrHlpr := dirMgrHelper{}

	err = dMgrHlpr.isDirMgrValid(
		&parentDirectory,
		ePrefix.XCpy(
			"parentDirectory"))

	if err != nil {
		return DirMgr{},
			fmt.Errorf("%v\n"+
				"Input parameter 'parentDirectory' is invalid!\n"+
				"Error= \n%v\n",
				funcName,
				err.Error())
	}

	isEmpty := false

	isEmpty,
		err = dMgrHlpr.setDirMgrFromKnownPathDirName(
		&newDirMgr,
		parentDirectory.absolutePath,
		directoryFileInfo.Name(),
		"newDirMgr",
		"parentDirectory",
		"directoryFileInfo.Name()",
		ePrefix)

	if err == nil && isEmpty {

		err = fmt.Errorf("%v\n"+
			"ERROR: The DirMgr instance generated is empty and contains no data!\n"+
			"parentDirectory= '%v'\n"+
			"directory= '%v'\n",
			ePrefix.String(),
			parentDirectory.absolutePath,
			directoryFileInfo.Name())
	}

	if err != nil {
		return DirMgr{}, err
	}

	return newDirMgr, nil
}

// NewFromFileMgr
//
// Configures and returns a new 'DirMgr' instance based
// on input parameter 'fileMgr' which is of type
// 'FileMgr'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fMgr						FileMgr
//
//		A valid, concrete instance of FileMgr. The data
//		elements in this instance will be used to
//		construct and a return a new, full populated
//		instance of DirMgr.
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
//	error
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
func (dMgr *DirMgr) NewFromFileMgr(
	fMgr FileMgr,
	errorPrefix interface{}) (
	DirMgr,
	error) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgr.NewFromFileMgr()",
		"")

	if err != nil {
		return DirMgr{}, err
	}

	err = fMgr.IsValidInstanceError(
		ePrefix.XCpy("fMgr"))

	if err != nil {
		return DirMgr{}, err
	}

	return new(dirMgrHelperAtom).copyOut(
		&fMgr.dMgr,
		ePrefix.XCpy("fMgr.dMgr->"))
}

// NewFromKnownPathDirectoryName - Configures and returns
// a new 'DirMgr' instance using a parent path name and
// directory name. The parent path and directory name are
// combined to form the full path for the new 'DirMgr'
// instance.
//
// This method will populate all internal field values
// with new values based on input parameters 'parentPathName'
// and 'directoryName'.
//
// This method differs from similar methods in that it assumes
// the input parameters are known values and do not require
// the usual analysis and validation screening applied by
// other methods.
func (dMgr *DirMgr) NewFromKnownPathDirectoryName(
	parentPathName string,
	directoryName string,
	errorPrefix interface{}) (
	DirMgr,
	error) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	newDirMgr := DirMgr{}

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgr.NewFromKnownPathDirectoryName()",
		"")

	if err != nil {
		return newDirMgr, err
	}

	dMgrHlpr := dirMgrHelper{}

	var isEmpty bool

	isEmpty,
		err = dMgrHlpr.setDirMgrFromKnownPathDirName(
		&newDirMgr,
		parentPathName,
		directoryName,
		"newDirMgr",
		"parentPathName",
		"directoryName",
		ePrefix)

	if err != nil {
		return DirMgr{}, err
	}

	if isEmpty {

		return DirMgr{},
			fmt.Errorf("%v\n"+
				"Newly generated 'DirMgr' is Empty!\n"+
				"dMgrHlpr.setDirMgrFromKnownPathDirName() returned an empty 'DirMgr'\n"+
				"parentPathName='%v'\n"+
				"directoryName='%v'\n",
				ePrefix.String(),
				parentPathName,
				directoryName)
	}

	return newDirMgr, nil
}

// SetDirMgr
//
// Sets the DirMgr fields and path strings for the
// current DirMgr instance.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathStr string
//
//		A path string designating a path or directory.
//		To reduce errors, the 'pathStr' should be
//		terminated with an appropriate path separator
//		('/' or '\').
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
//	isEmpty						bool
//
//		If the outcome of setting new values for DirMgr
//		is an 'empty' DirMgr instance, this value will be
//		set to 'true'.
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
//
// ----------------------------------------------------------------
//
// # Usage
//
//	Example 'pathStr': "C:\dirA\dirB\dirC\"
//
//	Example Output After DirMgr Configuration:
//
//	 ----------------------------
//	 	DirMgr Fields
//	 ----------------------------
//
//	                  isInitialized:  true
//	                   originalPath:  D:\go\work\src\MikeAustin71\pathfilego\003_filehelper\logTest\testoverwrite
//	                           path:  D:\go\work\src\MikeAustin71\pathfilego\003_filehelper\logTest\testoverwrite
//	                IsPathPopulated:  true
//	                  doesPathExist:  true
//	                     parentPath:  D:\go\work\src\MikeAustin71\pathfilego\003_filehelper\logTest
//	          isParentPathPopulated:  true
//	                   relativePath:  testoverwrite
//	        isRelativePathPopulated:  true
//	                   absolutePath:  D:\go\work\src\MikeAustin71\pathfilego\003_filehelper\logTest\testoverwrite
//	        isAbsolutePathPopulated:  true
//	isAbsolutePathDifferentFromPath:  false
//	          doesAbsolutePathExist:  true
//	                  directoryName:  testoverwrite
//	                     volumeName:  D:
//	              isVolumePopulated:  true
//	              actualDirFileInfo:
//	                     ========== File Info Data ==========
//	                       File Info IsDir():  true
//	                        File Info Name():  testoverwrite
//	                        File Info Size():  0
//	                     File Info ModTime():  2018-01-06 Sat 00:06:56.421236800 -0600 CST
//	                        File Info Mode():  drwxrwxrwx
//	                         File Info Sys():  &{16 {617269082 30594119} {2388100752 30639796} {2388100752 30639796} 0 0}
//	                                Dir path:  D:\go\work\src\MikeAustin71\pathfilego\003_filehelper\logTest\testoverwrite
func (dMgr *DirMgr) SetDirMgr(
	pathStr string,
	errorPrefix interface{}) (
	isEmpty bool,
	err error) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	isEmpty = true

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgr.SetDirMgr()",
		"")

	if err != nil {
		return isEmpty, err
	}

	isEmpty,
		err = new(dirMgrHelperNanobot).
		setDirMgr(
			dMgr,
			pathStr,
			"dMgr",
			"pathStr",
			ePrefix)

	return isEmpty, err
}

// SetDirMgrFromKnownPathDirName - Configures the internal
// field values for the current DirMgr instance using a parent
// path name and a directory name. The parent path and directory
// name are combined to form the full path for the current 'DirMgr'
// instance.
//
// This method will replace all previous field values with new
// values based on input parameters 'parentPathName' and
// 'directoryName'.
//
// This method differs from other "Set" methods in that it
// assumes the input parameters are known values and do not
// require the usual analysis and validation screening applied
// by similar methods.
//
// If more rigours input parameter validation is required,
// consider using method, DirMgr.SetDirMgr().
func (dMgr *DirMgr) SetDirMgrFromKnownPathDirName(
	parentPathName,
	directoryName string,
	errorPrefix interface{}) (
	isEmpty bool,
	err error) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	isEmpty = true

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgr."+
			"setDirMgrFromKnownPathDirName()",
		"")

	if err != nil {
		return isEmpty, err
	}

	isEmpty,
		err = new(dirMgrHelper).
		setDirMgrFromKnownPathDirName(
			dMgr,
			parentPathName,
			directoryName,
			"dMgr",
			"parentPathName",
			"directoryName",
			ePrefix)

	return isEmpty, err
}

// SetDirMgrWithFileInfo
//
// Sets the DirMgr fields and path strings for the
// current DirMgr object based on an input
// 'parentDirectoryPath' parameter and an os.FileInfo
// input parameter, 'info'.
//
// The os.FileInfo directory information will be added to
// the parent directory when constructing the new
// directory path which will be configured in the current
// instance of DirMgr.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	parentDirectoryPath			string
//
//		The parent directory path.
//
//	info						os.FileInfo
//
//		An os.FileInfo structure containing the directory
//		name which will be added to the parent directory
//		in order to create the new directory path.
//
//		Note:
//
//			An instance of FileInfoPlus may be submitted
//			for this parameter because FileInfoPlus
//			implements the os.FileInfo interface.
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
func (dMgr *DirMgr) SetDirMgrWithFileInfo(
	parentDirectoryPath string,
	info os.FileInfo,
	errorPrefix interface{}) (
	err error) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgr."+
			"SetDirMgrWithFileInfo()",
		"")

	if err != nil {
		return err
	}

	if info == nil {

		return fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'info' is 'nil' and INVALID!\n",
			ePrefix.String())
	}

	isEmpty := true

	isEmpty,
		err = new(dirMgrHelper).
		setDirMgrWithPathDirectoryName(
			dMgr,
			parentDirectoryPath,
			info.Name(),
			"dMgr",
			"parentDirectoryPath",
			"FileInfo.Name()",
			ePrefix)

	if err == nil && isEmpty {

		err = fmt.Errorf("%v\n"+
			"Newly generated 'DirMgr' is Empty!\n"+
			"dMgrHlpr.setDirMgrFromKnownPathDirName() returned an empty 'DirMgr'\n"+
			"parentDirectoryPath= '%v'\n"+
			"FileInfo.Name()= '%v'\n",
			ePrefix.String(),
			parentDirectoryPath,
			info.Name())
	}

	return err
}

// SetPermissions
//
// Sets the read/write and execute permissions for the
// directory identified by the current DirMgr instance.
// Note the treatment of execute permissions may vary by
// operating system.
//
// The permissions are configured based on input
// parameter 'permissionConfig' which is of type,
// 'FilePermissionConfig'. For an explanation of
// permission codes, see method
// 'FilePermissionConfig.New()'.
//
// If the directory identified by the current DirMgr
// instance does not exist, an error will be returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	permissionConfig			FilePermissionConfig
//
//		Use FilePermissionConfig 'New' methods to create
//		directory permissions for the directory
//		identified by the current instance of DirMgr.
//
//		Type FilePermissionConfig provides methods to
//		support the creation and management of File
//		Permissions for use in controlling file access
//		operations. The Go Programming Language uses
//		os.FileMode and unix permission bits to configure
//		file permissions.
//
//		Reference:
//		https://golang.org/pkg/os/#FileMode
//		https://www.cyberciti.biz/faq/explain-the-nine-permissions-bits-on-files/
//		https://en.wikipedia.org/wiki/File_system_permissions
//
//		The FilePermissionConfig methods will allow for
//		configuration of valid file permissions which are
//		subsequently stored as an os.FileMode in a
//		private member variable,
//		'FilePermissionConfig.fileMode'.
//
//		When evaluated as a string, file permission is
//		defined by a 10-character string. The first
//		character is an 'Entry Type' and the remaining
//		9-characters are unix permission bits.
//
//		Symbolic and Numeric Notation
//
//		Permission codes may be designated with Symbolic
//		Notation or Numeric Octal Notation.
//
//				   Numeric
//		Symbolic	Octal
//		Notation   Notation
//		----------	0000	no permissions
//		-rwx------	0700	read, write, & execute only for owner
//		-rwxrwx---	0770	read, write, & execute for owner and group
//		-rwxrwxrwx	0777	read, write, & execute for owner, group and others
//		---x--x--x	0111	execute
//		--w--w--w-	0222	write
//		--wx-wx-wx	0333	write & execute
//		-r--r--r--	0444	read
//		-r-xr-xr-x	0555	read & execute
//		-rw-rw-rw-	0666	read & write
//		-rwxr-----	0740	owner can read, write, & execute; group can only read;
//		                   others have no permissions
//
//		The permissions used by this method are designed
//		to be used for directories, not files.
//
//		Example:
//
//		drwxrwxrwx - Identifies permissions for directory
//						value = 020000000777
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
func (dMgr *DirMgr) SetPermissions(
	permissionConfig FilePermissionConfig,
	errorPrefix interface{}) (
	err error) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgr.SetPermissions()",
		"")

	if err != nil {
		return err
	}

	dMgrHlpr := dirMgrHelper{}

	err = dMgrHlpr.setPermissions(
		dMgr,
		permissionConfig,
		"dMgr",
		"permissionConfig",
		ePrefix)

	return err
}

// SubstituteBaseDir
//
// Substitute 'baseDir' segment of the current DirMgr
// with a new parent directory identified by input
// parameter 'substituteBaseDir'. This is useful in
// copying files to new directory trees.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	baseDir						DirMgr
//
//		A pointer to an instance of DirMgr. Input
//		parameter 'baseDir' must be equivalent to a
//		segment of the directory identified by the
//		current instance of DirMgr.
//
//		This segment will be replaced by the directory
//		string specified by input parameter
//		'substituteBaseDir'.
//
//	substituteBaseDir			DirMgr
//
//		A pointer to an instance of DirMgr. This directory
//		string will be substituted for the directory
//		segment in the current DirMgr instance identified
//		by input parameter 'baseDir'.
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
//	newDMgr						DirMgr
//
//		If this method completes successfully, an
//		instance of DirMgr populated with a directory
//		path constructed from the current instance of
//		DirMgr and input parameter 'substituteBaseDir'.
//
//		This method will substitute the 'baseDir' segment
//		of directory identified by the current instance
//		of DirMgr with the new directory string specified
//		by input parameter 'substituteBaseDir'. The newly
//		configured directory path will be returned by this
//		parameter, 'newDMgr'.
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
func (dMgr *DirMgr) SubstituteBaseDir(
	baseDir DirMgr,
	substituteBaseDir DirMgr,
	errorPrefix interface{}) (
	newDMgr DirMgr,
	err error) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgr."+
			"SubstituteBaseDir()",
		"")

	if err != nil {
		return newDMgr, err
	}

	newDMgr,
		err = new(dirMgrHelper).
		substituteBaseDir(
			dMgr,
			&baseDir,
			&substituteBaseDir,
			"DirMgr",
			"baseDir",
			"substituteBaseDir",
			ePrefix)

	return newDMgr, err
}
