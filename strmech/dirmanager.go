package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"os"
	"strings"
	"sync"
)

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

// CopyDirectoryFiles
//
// Copies files from the directory identified by the
// current instance of DirMgr to a target directory.
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
//	(1)	This method ONLY selects and copies files from
//		the parent directory identified by 'sourceDMgr'
//		to the parent directory identified by
//		'targetDMgr'.
//
//	(2) No files in subdirectories of 'sourceDMgr 'will
//		be copied. Only files in the top level or parent
//		directory defined by input parameter 'sourceDMgr'
//		are eligible for the copy operation.
//
//	(3)	If the target directory does not exist, this
//		method will attempt to create it.
//
//	(4)	Files selected for the copy operation are
//		required to match two sets of selection criteria,
//		File Type Selection Criteria and File
//		Characteristics Selection Criteria.
//
//	(5) File Type Selection Criteria specifications are
//		passed as input parameters 'copyRegularFiles',
//		'copySymLinkFiles' and 'copyOtherNonRegularFiles'.
//		For an explanation of Regular and Non-Regular
//		files, see the section on Definition of Terms,
//		above.
//
//	(6) File Characteristics Selection Criteria are user
//		specified selection requirements passed as input
//		parameter 'fileSelectCriteria'. This file
//		selection criteria allows users to screen files
//		for File Name, File Modification Date and File
//		Mode.
//
//	(7) If the source directory identified by input
//		parameter 'sourceDMgr' contains NO Files meeting
//		(1) the File Type Selection Criteria and (2) the
//		File Characteristics Selection Criteria, this
//		method will exit, no files will be copied to the
//		target directory ('targetDMgr') and no error will
//		be returned.
//
//	(8) If the source directory identified by input
//		parameter 'sourceDMgr' contains NO Files
//		whatsoever (0 Files), this method will exit, no
//		files will be copied to the target directory
//		('targetDMgr') and no error will be returned.
//
//	(9) If input parameter 'returnCopiedFilesList' is set
//		to 'false', input parameter ('copiedFiles') may
//		safely be set to 'nil'.
//
//	(10) If input parameter 'returnSubDirsList' is set to
//		 'false', input parameter ('subDirectories') may
//		 safely be set to 'nil'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	targetDMgr   					DirMgr
//
//		An instance of 'DirMgr' initialized with the
//		directory path of the target directory to which
//		selected files will be copied. If the target
//		directory does not exist, this method will
//		attempt to create it.
//
//	returnCopiedFilesList			bool
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
//		If 'returnCopiedFilesList' is set to false, input
//		parameter 'copiedFiles' may safely be set to
//		'nil'.
//
//	returnSubDirsList			bool
//
//		If this parameter is set to 'true', this method
//		will create DirMgr objects for each subdirectory
//		in the parent directory ('sourceDMgr'), and add
//		them to the Directory Manager Collection passed
//		as input parameter 'subDirectories'. Only
//		subdirectories residing in the top level or
//		parent directory defined by 'targetDMgr' will be
//		included.
//
//		If 'returnSubDirsList' is set to 'false', no
//		subdirectories will be added to the Directory
//		Manager Collection (DirMgrCollection) passed as
//		input parameter 'subDirectories'.
//
//		If 'returnSubDirsList' is set to false,
//		input parameter 'subDirectories' may safely be set
//		to 'nil'.
//
//	copyEmptyTargetDirectory		bool
//
//		If set to 'true' the target directory will be
//		created regardless of whether any files are
//		copied to that directory. Remember that files are
//		only copied to the target directory if they meet
//		the File Type and File Characteristics selection
//		criteria.
//
//	copyRegularFiles				bool
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
//		For an explanation of Regular and Non-Regular
//		files, see the section on "Definition Of Terms",
//		above.
//
//		If input parameters 'copyRegularFiles',
//		'copySymLinkFiles' and 'copyOtherNonRegularFiles'
//		are all set to 'false', an error will be returned.
//
//	copyOtherNonRegularFiles		bool
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
//	fileSelectCriteria 				FileSelectionCriteria
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
//		object are 'Inactive' or 'Not Set' (set to their zero or default
//		values), then all the files meeting the File Type requirements in the
//		directory defined by the current DirMgr instance will be selected.
//
//			Example:
//			  fsc := FileSelectCriterionMode{}
//
//			  In this example, 'fsc' is NOT initialized. Therefore,
//			  all the selection criterion are 'Inactive'. Consequently,
//			  all the files meeting the File Type requirements in the
//			  directory defined by the current DirMgr instance will be
//			  selected.
//
//		------------------------------------------------------------------------
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
//		If input parameter 'returnSubDirsList' is set to
//		'false', no subdirectories will be added to this
//		Directory Manager Collection.
//
//		If input parameter 'returnSubDirsList' is set
//		to 'false', this parameter ('subDirectories') may
//		safely be set to 'nil'.
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
//	dirCopyStats				DirectoryCopyStats
//
//		If this method completes successfully, this
//		return parameter will be populated with
//		information and statistics on the copy
//		operation. This information includes the number
//		of files copied.
//
//		The data elements in this structure are used
//		to accumulate statistics and information
//		related to files copied from a single source
//		directory to a single destination or target
//		directory.
//
//		type DirectoryCopyStats struct {
//			DirsCreated uint64
//				The number of new directories created.
//
//			TotalFilesProcessed uint64
//				The total number of files processed.
//				Does NOT include directory entries.
//
//			FilesCopied uint64
//				The number of files copied. Does
//				NOT include directory entries.
//
//			FileBytesCopied uint64
//				The number of file bytes copied.
//				Does NOT include directory entries.
//
//			FilesNotCopied uint64
//				The number of files processed, but
//				NOT copied. Does NOT include directory
//				entries.
//
//			FileBytesNotCopied uint64
//				The number of bytes associated with
//				files processed but NOT copied. Does
//				NOT include directory entries.
//
//			SubDirs uint64
//				The total number of subdirectories identified
//				during the directory tree copy operation. This
//				does NOT include the parent directory.
//
//			SubDirsDocumented uint64
//				The number of subdirectories identified
//				and returned in a Directory Manager
//				Collection. Does NOT include the parent
//				directory. Subdirectories are only
//				documented if requested. This computation
//				value is therefore optional.
//
//			CopiedFilesDocumented uint64
//				The number of copied files documented
//				by adding a File Manager object to a
//				returned File Manager Collection.
//
//			Errors []error
//				An array of errors associated with the
//				calculation of these statistics.
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
//
//		Upon method completion, be sure to check both
//		Non-Fatal and Fatal errors.
func (dMgr *DirMgr) CopyDirectoryFiles(
	targetDMgr DirMgr,
	returnCopiedFilesList bool,
	returnSubDirsList bool,
	copyEmptyTargetDirectory bool,
	copyRegularFiles bool,
	copySymLinkFiles bool,
	copyOtherNonRegularFiles bool,
	fileSelectCriteria FileSelectionCriteria,
	subDirectories *DirMgrCollection,
	copiedFiles *FileMgrCollection,
	errorPrefix interface{}) (
	dirCopyStats DirectoryCopyStats,
	nonfatalErrs []error,
	fatalErr error) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		fatalErr = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgr.CopyDirectoryFiles()",
		"")

	if fatalErr != nil {

		return dirCopyStats, nonfatalErrs, fatalErr
	}

	return new(dirMgrHelperPlanck).copyDirectoryFiles(
		dMgr,
		&targetDMgr,
		returnCopiedFilesList,
		returnSubDirsList,
		copyEmptyTargetDirectory,
		copyRegularFiles,
		copySymLinkFiles,
		copyOtherNonRegularFiles,
		fileSelectCriteria,
		"dMgr",
		"targetDMgr",
		subDirectories,
		copiedFiles,
		ePrefix)
}

// CopyDirectoryTree
//
// Copies all selected files in the directory tree
// defined by the current instance of DirMgr to a
// specified target directory tree.
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
//		entire the source directory tree identified by
//		the current instance of DirMgr. The files in this
//		source directory tree will be copied to a
//		corresponding target directory tree identified by
//		input parameter 'targetDMgr'.
//
//	(2)	This copy operation INCLUDES the top level or
//		parent directory specified by the current instance
//		of DirMgr.
//
//	(3)	If a subdirectory in the target directory tree
//		does not exist, this method will attempt to
//		create it.
//
//	(4)	Files will only be copied if they meet the File
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
//		specified by input parameter 'fileSelectCriteria'.
//
//	(5)	To restrict the copy operation to selected files
//		in the subdirectory tree, thereby EXCLUDING the
//		source parent directory from copy operation, see
//		method:
//
//			DirMgr.CopySubDirectoryTree()
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//	returnCopiedFilesList			bool
//
//		If input parameter 'returnCopiedFilesList' is set
//		to 'true', this method will return a populated
//		File Manager Collection documenting all the files
//		actually included in the directory tree copy
//		operation.
//
//		If input parameter 'returnCopiedFilesList' is set
//		to 'false', this method will return an empty and
//		unpopulated instance of FileMgrCollection. This
//		means that the files actually copied by this
//		method will NOT be documented.
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
//		then all the files meeting the File Type requirements in the source
//		directory will be selected.
//
//			Example:
//			  fsc := FileSelectCriterionMode{}
//
//			  In this example, 'fsc' is NOT initialized. Therefore,
//			  all the selection criterion are 'Inactive'. Consequently,
//			  all the files meeting the File Type requirements in the
//			  source directory will be selected.
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
//				Errors        error
//			}
//
//	copiedDirTreeFiles			FileMgrCollection
//
//		If input parameter 'returnCopiedFilesList' is set
//		to 'true', 'copiedDirTreeFiles' will return a
//		populated File Manager Collection including all
//		the files actually included in the directory tree
//		copy operation.
//
//		If input parameter 'returnCopiedFilesList' is set
//		to 'false', 'copiedDirTreeFiles' will return an
//		empty and unpopulated instance of
//		FileMgrCollection. This means that the files
//		actually copied by this method will NOT be
//		documented.
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
func (dMgr *DirMgr) CopyDirectoryTree(
	targetDMgr DirMgr,
	returnCopiedFilesList bool,
	copyEmptyTargetDirectory bool,
	copyRegularFiles bool,
	copySymLinkFiles bool,
	copyOtherNonRegularFiles bool,
	fileSelectCriteria FileSelectionCriteria,
	errorPrefix interface{}) (
	dTreeCopyStats DirTreeCopyStats,
	copiedDirTreeFiles FileMgrCollection,
	nonfatalErrs []error,
	fatalErr error) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		fatalErr = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgr.CopyDirectoryTree()",
		"")

	if fatalErr != nil {

		nonfatalErrs = append(nonfatalErrs, err)

		return dTreeCopyStats,
			copiedDirTreeFiles,
			nonfatalErrs,
			fatalErr
	}

	dTreeCopyStats,
		copiedDirTreeFiles,
		nonfatalErrs,
		fatalErr = new(dirMgrHelperNanobot).
		copyDirectoryTree(
			dMgr,
			&targetDMgr,
			returnCopiedFilesList,
			false, // skipTopLevelDirectory
			copyEmptyTargetDirectory,
			copyRegularFiles,
			copySymLinkFiles,
			copyOtherNonRegularFiles,
			fileSelectCriteria,
			"dMgr",
			"targetDMgr",
			ePrefix)

	return dTreeCopyStats,
		copiedDirTreeFiles,
		nonfatalErrs,
		fatalErr
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

	return new(dirMgrHelperBoson).copyDirMgrs(
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

	newDirMgr := DirMgr{}

	err = new(dirMgrHelperBoson).copyDirMgrs(
		&newDirMgr,
		dMgr,
		ePrefix.XCpy("dMgr->"))

	return newDirMgr, err
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
// # BE ADVISED
//
//	(1)	This method will scan subdirectories for
//		selected files in the source directory tree
//		defined by the current instance of DirMgr.
//		The files in this source directory tree will be
//		copied to a corresponding target directory tree
//		identified by input parameter 'targetDMgr'.
//
//	(2)	This scan and copy operation EXCLUDES the top
//		level or parent directory specified by the
//		current instance of DirMgr.
//
//	(3)	If a subdirectory in the target directory tree
//		does not exist, this method will attempt to
//		create it.
//
//	(4)	Files will only be copied if they meet the File
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
//		specified by input parameter 'fileSelectCriteria'.
//
//	(5)	To scan and copy selected files in the entire
//		directory tree, thereby INCLUDING the parent
//		directory in the scan and copy operation, see
//		method:
//
//			DirMgr.CopyDirectoryTree()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	targetDMgr						DirMgr
//
//		A concrete instance of DirMgr. This instance
//		specifies the target directory to which the
//		contents of the current DirMgr directory tree
//		will be copied.
//
//	returnCopiedFilesList			bool
//
//		If input parameter 'returnCopiedFilesList' is set
//		to 'true', this method will return a populated
//		File Manager Collection documenting all the files
//		actually included in the directory tree copy
//		operation.
//
//		If input parameter 'returnCopiedFilesList' is set
//		to 'false', this method will return an empty and
//		unpopulated instance of FileMgrCollection. This
//		means that the files actually copied by this
//		method will NOT be documented.
//
//	copyEmptyTargetDirectory		bool
//
//		If set to 'true' the target directory will be
//		created regardless of whether any files are
//		copied to that directory. Remember that files are
//		only copied to the target directory if they meet
//		file selection criteria specified by input
//		parameter 'fileSelectCriteria'.
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
//	fileSelectCriteria				FileSelectionCriteria
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
//		then all the files meeting the File Type requirements in the source
//		directory will be selected.
//
//			Example:
//			  fsc := FileSelectCriterionMode{}
//
//			  In this example, 'fsc' is NOT initialized. Therefore,
//			  all the selection criterion are 'Inactive'. Consequently,
//			  all the files meeting the File Type requirements in the
//			  source directory will be selected.
//
//		------------------------------------------------------------------------
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
//	dTreeCopyStats					DirTreeCopyStats
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
//				Errors        error
//			}
//
//	copiedDirTreeFiles				FileMgrCollection
//
//		If input parameter 'returnCopiedFilesList' is set
//		to 'true', 'copiedDirTreeFiles' will return a
//		populated File Manager Collection including all
//		the files actually included in the directory tree
//		copy operation.
//
//		If input parameter 'returnCopiedFilesList' is set
//		to 'false', 'copiedDirTreeFiles' will return an
//		empty and unpopulated instance of
//		FileMgrCollection. This means that the files
//		actually copied by this method will NOT be
//		documented.
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
//		Non-fatal errors usually involve failure
//		to copy individual files.
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
func (dMgr *DirMgr) CopySubDirectoryTree(
	targetDMgr DirMgr,
	returnCopiedFilesList bool,
	copyEmptyTargetDirectory bool,
	copyRegularFiles bool,
	copySymLinkFiles bool,
	copyOtherNonRegularFiles bool,
	fileSelectCriteria FileSelectionCriteria,
	errorPrefix interface{}) (
	dTreeCopyStats DirTreeCopyStats,
	copiedDirTreeFiles FileMgrCollection,
	nonfatalErrs []error,
	fatalErr error) {

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

		nonfatalErrs = append(nonfatalErrs, err)

		return dTreeCopyStats,
			copiedDirTreeFiles,
			nonfatalErrs,
			fatalErr
	}

	dTreeCopyStats,
		copiedDirTreeFiles,
		nonfatalErrs,
		fatalErr = new(dirMgrHelperNanobot).
		copyDirectoryTree(
			dMgr,
			&targetDMgr,
			returnCopiedFilesList,
			true, // skipTopLevelDirectory
			copyEmptyTargetDirectory,
			copyRegularFiles,
			copySymLinkFiles,
			copyOtherNonRegularFiles,
			fileSelectCriteria,
			"dMgr",
			"targetDMgr",
			ePrefix)

	return dTreeCopyStats,
		copiedDirTreeFiles,
		nonfatalErrs,
		fatalErr
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
//	and files in the child directory trees.
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
// # IMPORTANT
//
// This method deletes all the files in the directory
// identified by the current instance of DirMgr.
//
// ONLY files in the top level directory identified
// by the current instance of DirMgr are deleted.
// Subdirectories are NOT deleted and files in
// subdirectories are NOT deleted.
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//		the instance of FileMgrCollection returned by
//		this method will always be empty and unpopulated.
//		This means that the files actually deleted by
//		this method will NOT be documented.
//
//		If 'returnDeletedFilesList' is set to false,
//		input parameter 'deletedFiles' may safely be set
//		to 'nil'.
//
//	deletedFiles				*FileMgrCollection
//
//		A pointer to an instance of FileMgrCollection
//		which encapsulates an array of File Manager
//		(FileMgr) objects.
//
//		If input parameter 'returnDeletedFilesList' is
//		set to 'true', all files actually deleted in the
//		target directory defined by input parameter
//		'targetDMgr' will be added as new FileMgr objects
//		to the 'deletedFiles' File Manager Collection.
//		Effectively, this provides a list documenting the
//		files actually deleted to the target directory.
//
//		If input parameter 'returnDeletedFilesList' is
//		set to 'false', no files will be added to this
//		File Manager collection ('deletedFiles'). This
//		means that no documentation of deleted files will
//		be provided.
//
//		If input parameter 'returnDeletedFilesList' is
//		set to 'false', this input parameter
//		('deletedFiles') may safely be set to 'nil'.
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
//		type DeleteDirFilesStats struct {
//			TotalFilesProcessed        uint64
//				The total number of files processed.
//				Does NOT include directory entries.
//
//			FilesDeleted               uint64
//				The number of files deleted. Does
//				NOT include directory entries.
//
//			FilesDeletedBytes          uint64
//				The number of file bytes deleted.
//				Does NOT include directory entries.
//
//			FilesRemaining             uint64
//				The number of files processed, but
//				NOT deleted. Does NOT include directory
//				entries.
//
//			FilesRemainingBytes        uint64
//				The number of bytes associated with
//				files processed but NOT copied. Does
//				NOT include directory entries.
//
//			TotalSubDirectories        uint64
//				Total SubDirectories processed
//
//			TotalDirsScanned           uint64
//				Total Directories Scanned.
//
//			NumOfDirsWhereFilesDeleted uint64
//				The number of parent directories and
//				subdirectories where files were deleted.
//
//			DirectoriesDeleted         uint64
//				The number of directories deleted.
//
//			SubDirsDocumented          uint64
//				The number of subdirectories identified
//				and returned in a Directory Manager
//				Collection. Does NOT include the parent
//				directory. Subdirectories are only
//				documented if requested. This computation
//				value is therefore optional.
//
//			DeletedFilesDocumented		uint64
//				The number of deleted files documented
//				through addition to a returned File
//				Manager Collection
//
//			Errors []error
//				An array of errors associated with the
//				calculation of these statistics.
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
func (dMgr *DirMgr) DeleteAllFilesInDir(
	returnDeletedFilesList bool,
	deletedFiles *FileMgrCollection,
	errorPrefix interface{}) (
	deleteDirStats DeleteDirFilesStats,
	nonfatalErrs []error,
	fatalErr error) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		fatalErr = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgr."+
			"DeleteAllFilesInDir()",
		"")

	if fatalErr != nil {

		return deleteDirStats, nonfatalErrs, fatalErr
	}

	deleteDirStats,
		nonfatalErrs,
		fatalErr = new(dirMgrHelperPlanck).
		deleteDirectoryFiles(
			dMgr,
			returnDeletedFilesList,
			false,                   // returnSubDirsList
			true,                    // deleteRegularFiles
			true,                    // deleteSymLinkFiles
			true,                    // deleteOtherNonRegularFiles
			FileSelectionCriteria{}, // Select All Files
			"dMgr",
			nil,          // subDirectories
			deletedFiles, // deletedFiles
			ePrefix)

	return deleteDirStats, nonfatalErrs, fatalErr
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
func (dMgr *DirMgr) DeleteAllSubDirectories(
	returnDeletedSubDirs bool,
	deletedSubDirs *DirMgrCollection,
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
		"DirMgr.DeleteAllSubDirectories()",
		"")

	if err != nil {

		return err
	}

	err = new(dirMgrHelperNanobot).
		deleteAllSubDirectories(
			dMgr,
			returnDeletedSubDirs,
			deletedSubDirs,
			"dMgr",
			ePrefix)

	return err
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

	var errs []error

	deleteFilesInfo,
		errs =
		new(dirMgrHelper).deleteDirectoryTreeInfo(
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

// DoesAbsolutePathExist
//
// Performs two operations.
//
// First this method determines whether the absolute
// directory path indicated by the DirMgr.absolutePath
// field actually does exist on disk and returns a 'true'
// or 'false' boolean value accordingly.
//
// Second, this method also updates the DirMgr field
// 'DirMgr.doesAbsolutePathExist'.
//
// ----------------------------------------------------------------
//
// # Definition of Terms
//
// An absolute or full path points to the same location
// in a file system, regardless of the current working
// directory. To do that, it must include the root
// directory.
//
//	https://en.wikipedia.org/wiki/Path_(computing)#Absolute_and_relative_paths
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	--- NONE ---
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		If the absolute directory path specified by the
//		current instance of DirMgr actually does exist on
//		disk, this return value is set to 'true'.
//
//		If the absolute directory path does NOT exist, a
//		value of 'false' is returned.
func (dMgr *DirMgr) DoesAbsolutePathExist() bool {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	var err error

	var dirPathDoesExist bool

	dirPathDoesExist,
		_,
		err =
		new(dirMgrHelperAtom).
			doesDirectoryExist(
				dMgr,
				PreProcPathCode.None(),
				"dMgr",
				nil)

	if err != nil {
		dirPathDoesExist = false
	}

	return dirPathDoesExist
}

// DoesDirectoryExist
//
// Returns a boolean value indicating whether the
// Directory path exists on an attached storage drive.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	--- NONE ---
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	doesPathExist				bool
//
//		If the directory path specified by the current
//		instance of DirMgr exists on an attached storage
//		drive, this return parameter will be set to
//		'true'.
func (dMgr *DirMgr) DoesDirectoryExist() (
	doesPathExist bool) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	var err error

	dirPathDoesExist := false

	dirPathDoesExist,
		_,
		err =
		new(dirMgrHelperAtom).
			doesDirectoryExist(
				dMgr,
				PreProcPathCode.None(),
				"dMgr",
				nil)

	if err != nil {
		dirPathDoesExist = false
	}

	return dirPathDoesExist
}

// DoesPathExist
//
// Performs two operations on the current instance of
// DirMgr.
//
// First the method determines whether the directory
// path indicated by the DirMgr.path field actually
// does exist on disk and returns a 'true' or 'false'
// boolean value accordingly.
//
// In addition, it also updates the DirMgr field
// DirMgr.doesPathExist field.
//
// The DirMgr directory may be configured either as
// an absolute path or a relative path.
//
// ----------------------------------------------------------------
//
// # Definition of Terms
//
//	An absolute or full path points to the same location
//	in a file system, regardless of the current working
//	directory. To do that, it must include the root
//	directory.
//
//	By contrast, a relative path starts from some given
//	working directory, avoiding the need to provide the
//	full absolute path. A filename can be considered as
//	a relative path based at the current working directory.
//
//	https://en.wikipedia.org/wiki/Path_(computing)#Absolute_and_relative_paths
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	--- NONE ---
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		If the directory path specified by the current
//		instance of DirMgr actually exists on disk, this
//		return parameter will be set to 'true'.
func (dMgr *DirMgr) DoesPathExist() bool {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	dirPathDoesExist := false

	var err error

	dirPathDoesExist,
		_,
		err =
		new(dirMgrHelperAtom).
			doesDirectoryExist(
				dMgr,
				PreProcPathCode.None(),
				"dMgr",
				nil)

	if err != nil {
		dirPathDoesExist = false
	}

	return dirPathDoesExist
}

// DoesThisDirectoryExist
//
// This method performs a test to determine if the
// directory specified by the current DirMgr instance
// actually exists on disk.
//
// This method returns a boolean value of 'true' if the
// directory identified by the current DirMgr instance
// does in fact exist. However, this method differs from
// similar methods in that it returns an error if any
// Non-Path errors are encountered during the path
// existence test.
//
// Non-Path errors are most commonly associated with
// 'access-denied' situations. However, there may be
// other reasons for triggering Non-Path errors.
//
// If a Non-Path error is encountered, an appropriate
// error message is returned along with a boolean value
// of 'false'.
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
//	directoryDoesExist			bool
//
//		If the directory path specified by the current
//		instance of DirMgr actually exists on disk, this
//		return parameter is set to 'true'.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If, during the process of verifying the existence
//		of the current directory, an error is encountered
//	 	it will be a non-path error.
//
//		Non-Path errors are most commonly associated with
//		'access-denied' situations. However, there may be
//		other reasons for triggering Non-Path errors.
//
//		If a Non-Path error is encountered, an appropriate
//		error message is returned and the return parameter
//		'directoryDoesExist' will be set to 'false'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
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

// Empty
//
// Resets all internal member data field values to their
// uninitialized or original zero values for the current
// instance of DirMgr.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and reset all pre-existing
//	data values in the current instance of DirMgr to
//	their original zero or uninitialized values.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	--- NONE ---
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	--- NONE ---
func (dMgr *DirMgr) Empty() {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	_ = new(dirMgrHelperElectron).
		empty(dMgr, "DirMgr.Empty() ", nil)

	dMgr.lock.Unlock()

	dMgr.lock = nil
}

// Equal
//
// Compares two DirMgr objects to determine if they are
// equal in all respects.
//
// The Directory Path comparisons are NOT case-sensitive.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dMgr2						*DirMgr
//
//		A pointer to an external instance of DirMgr. The
//		internal member data field values contained in
//		this instance will be compared to the
//	 	corresponding data values in the current instance
//		of DirMgr. If the compared data values are equal
//		in all respects, a boolean value of 'true' will
//		be returned.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		This method compares the internal data values
//		contained in the current instance of DirMgr to
//		the corresponding data values in an external
//		DirMgr instance passed as 'dMgr2'.
//
//		If all the data values contained in the current
//		instance of DirMgr and input parameter 'dMgr2'
//		are equal in all respects, this return parameter
//		will be set to 'true'.
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
// current directory manager and the external directory
// manager passed as input paramter 'incomingDirMgr'.
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
// # Definition of Terms
//
// An absolute or full path points to the same location
// in a file system, regardless of the current working
// directory. To do that, it must include the root
// directory.
//
//	https://en.wikipedia.org/wiki/Path_(computing)#Absolute_and_relative_paths
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

// EqualPaths
//
// Compares two DirMgr objects to determine if their
// paths are equal. Both Directory Path and absolute path
// must be equivalent.
//
// Be advised that the Directory Path for a DirMgr object
// may be an absolute path or a relative path.
//
// ----------------------------------------------------------------
//
// # Definition of Terms
//
//	An absolute or full path points to the same location
//	in a file system, regardless of the current working
//	directory. To do that, it must include the root
//	directory.
//
//	By contrast, a relative path starts from some given
//	working directory, avoiding the need to provide the
//	full absolute path. A filename can be considered as
//	a relative path based at the current working directory.
//
//	https://en.wikipedia.org/wiki/Path_(computing)#Absolute_and_relative_paths
//
// ----------------------------------------------------------------
//
// If the compared paths are equal, the method returns
// 'true'. If the paths are NOT equal, the method returns
// 'false'.
//
// The comparisons are NOT case-sensitive. In other
// words, all paths are converted to lower case before
// making the comparisons.
//
// If either the current DirMgr ('dMgr') or the input
// parameter 'dMgr2' are uninitialized, a value of
// 'false' is returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dMgr2						*DirMgr
//
//		A pointer to an instance of DirMgr. If both the
//		Directory Path and absolute path contained in
//		this instance are equivalent to those contained
//		in the current DirMgr instance, this method will
//		return a boolean value of 'true'.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		If both the Directory Path and absolute path
//		contained in input parameter 'dMgr2' are
//		equivalent to those contained in the current
//		DirMgr instance, this method will return a
//		boolean value of 'true'.
func (dMgr *DirMgr) EqualPaths(dMgr2 *DirMgr) bool {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	isEqual := false

	isEqual = new(dirMgrHelper).equalPaths(
		dMgr,
		dMgr2)

	return isEqual
}

// ExecuteDirectoryFileOps
//
// Performs a file operation on specified 'selected'
// files in the current directory ONLY. This function
// does NOT perform operations on the subdirectories
// (a.k.a. the directory tree).
//
// To perform file operations on the entire Directory
// Tree, see Function 'ExecuteDirectoryTreeOps()', below.
//
// The types of File Operations performed are generally
// classified as 'file copy' and 'file deletion
// ' operations. The precise file operation applied is
// defined by the type, 'FileOperationCode' which
// provides a series of constants used to identify the
// specific file operation applied.
//
// Input parameter, 'fileOps' is an array of type
// 'FileOperationCode' elements. Multiple file operations
// can be applied to a single file. For instance, a
// 'copy source to destination' operation can be followed
// by a 'delete source file' operation.
//
// The 'selected' files are identified by input parameter
// 'fileSelectCriteria' of type 'FileSelectionCriteria'.
// This file selection criteria is compared against all
// files in the directory (NOT the Directory Tree)
// identified by the current 'DirMgr' instance. When a
// match is found, that file is treated as a 'selected'
// source file and designated file operations are
// performed on that file.
//
// The results or final output from file operations
// utilizes the final input parameter, 'targetBaseDir' of
// type DirMgr. File operations are applied to selected
// source files and generated output is created in the
// 'targetBaseDir'.  For example 'copy' or 'move' file
// operations will transfer source files to 'targetBaseDir'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method performs File Operations ONLY on the
//	parent directory identified by the current DirMgr
//	instance.
//
//	No file operations are performed on the
//	subdirectories underneath this parent directory.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fileSelectCriteria			FileSelectionCriteria
//
//		This input parameter should be configured with
//		the desired file selection criteria. Files
//		matching this criteria will be subject to the
//		specified File Operations (fileOps).
//
//		If file 'fileSelectCriteria' is uninitialized
//		(FileSelectionCriteria{}), all files within
//		the current DirMgr parent directory will be
//		subject to the specified File Operations
//		(fileOps).
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
//		subjected to the specified file operations (fileOps).
//
//			Example:
//			  fsc := FileSelectCriterionMode{}
//
//			  In this example, 'fsc' is NOT initialized. Therefore,
//			  all the selection criterion are 'Inactive'. Consequently,
//			  all the files encountered in the DirMgr parent directory
//			  during the search operation will be subjected to the
//			  specified file operations.
//
//		------------------------------------------------------------------------
//
//	fileOps						[]FileOperationCode
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
//	targetBaseDir				DirMgr
//
//		The file selection criteria, 'fileSelectCriteria',
//		and the File Operations, 'fileOps' are applied to
//		files in the target base directory.
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
func (dMgr *DirMgr) ExecuteDirectoryFileOps(
	fileSelectCriteria FileSelectionCriteria,
	fileOps []FileOperationCode,
	targetBaseDir DirMgr,
	errorPrefix interface{}) (
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

// ExecuteDirectoryTreeOps
//
// Performs File Operations on specified 'selected' files
// in the directory tree identified by the current
// 'DirMgr' instance. The 'DirMgr' path therefore serves
// as the parent directory for file operations performed
// on the directory tree.
//
// If you wish to perform File Operations ONLY on the
// current directory and NOT THE ENTIRE DIRECTORY TREE,
// see Function "ExecuteDirectoryFileOps()", above.
//
// The types of File Operations performed are generally
// classified as 'file copy' and 'file deletion'
// operations. The precise file operation applied is
// defined by the type, 'FileOperationCode' which
// provides a series of constants, or enumerations, used
// to identify the specific file operation applied. Input
// parameter, 'fileOps' is an array of type
// 'FileOperationCode' elements. Multiple file operations
// can be applied to a single file. For instance, a 'copy
// source to destination' operation can be followed by a
// 'delete source file' operation.
//
// The 'selected' files are identified by input parameter
// 'fileSelectCriteria' of type 'FileSelectionCriteria'.
// This file selection criteria is compared against all
// files in the current directory tree identified by the
// current 'DirMgr' instance. When a match is found, that
// file is treated as a 'selected' file and designated
// file operations are performed on that file.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method performs File Operations on THE ENTIRE
// DIRECTORY TREE identified by the current instance
// of DirMgr.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fileSelectCriteria FileSelectionCriteria
//
//		This input parameter should be configured with
//		the desired file selection criteria. Files
//		matching this criteria will be identified as
//		'Selected Files'. The specified File Operations
//		(fileOps) will be performed on these selected
//		files.
//
//		If file 'fileSelectCriteria' is uninitialized
//		(FileSelectionCriteria{}), all files within
//		the current DirMgr directory tree will be
//		identified as 'Selected Files'. The specified
//		File Operations (fileOps) will be performed on
//		all these selected files.
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
//		The FileSelectionCriteria type allows for configuration of single
//		or multiple file selection criterion. The 'SelectCriterionMode' can
//		be used to specify whether the file must match all, or any one, of
//		the active file selection criterion.
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
//		subjected to the specified file operations.
//
//			Example:
//			  fsc := FileSelectCriterionMode{}
//
//			  In this example, 'fsc' is NOT initialized. Therefore,
//			  all the selection criterion are 'Inactive'. Consequently,
//			  all the files encountered in the DirMgr directory tree
//			  search operation will be subjected to the specified file
//			  operations.
//
//		------------------------------------------------------------------------
//
//	fileOps []FileOperationCode
//
//		An array of file operations to be performed on
//		each selected file. Selected files are identified
//		by matching the file selection criteria specified
//		by input parameter, 'fileSelectCriteria'.
//		See above.
//
//		The FileOperationCode type consists of the
//		following constants:
//
//		FileOperationCode(0).None()
//		  No Action
//
//		FileOperationCode(0).MoveSourceFileToDestinationFile()
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
//	targetBaseDir				DirMgr
//
//		The file selection criteria, 'fileSelectCriteria',
//		and the File Operations, 'fileOps' are applied to
//		files in the target base directory. This input
//		parameter is of type 'DirMgr'.
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
	targetBaseDir DirMgr,
	errorPrefix interface{}) (
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
// is treated as the parent directory for purposes of
// the search operation.
//
// The files selected must match selection criteria
// provided by input parameter 'fileSelectCriteria'.
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
//	fileSelectCriteria			FileSelectionCriteria
//
//	  This input parameter should be configured with the
//	  desired file selection criteria. Files matching
//	  this criteria will be returned as 'Found Files'.
//
//			type FileSelectionCriteria struct {
//
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
//		are 'Inactive' or 'Not Set' (set to their zero or default values), then
//		all the files processed in the directory tree will be selected and returned
//		as 'Found Files'.
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
//
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

// FindFilesByNamePattern
//
// This method searches the directory identified by the
// current DirMgr instance using a name pattern file
// search criteria.
//
// The directory specified by the current instance of
// DirMgr is treated as the target directory for the
// purposes of this file search.
//
// The file search operation will always be restricted
// to the target directory identified by the current
// DirMgr instance. No subdirectories will be included
// in the file search operation.
//
// If the 'fileSearchPattern' is an empty string or
// improperly formatted, an error will be returned.
//
// ------------------------------------------------------------------------
//
// # Input parameter
//
//	fileSearchPattern			string
//
//		The fileSearchPattern is string containing
//		parameters used to select target files in the
//		directory identified by the current 'DirMgr'
//		instance.
//
//		Example 'fileSearchPattern' strings
//
//		*.*             will match all files in directory.
//		*.html          will match  anyfilename.html
//		a*              will match  appleJack.txt
//		j????row.txt    will match  j1x34row.txt
//		data[0-9]*      will match 	data123.csv
//
//		Reference For File Pattern Matching Details:
//		  https://golang.org/pkg/path/filepath/#Match
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
//	FileMgrCollection
//
//		If this method completes successfully without
//		error, this returned FileMgrCollection instance
//		will contain an array of FileMgr types
//		identifying each of the files matched by input
//		parameter, 'fileSearchPattern'.
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
func (dMgr *DirMgr) FindFilesByNamePattern(
	fileSearchPattern string,
	errorPrefix interface{}) (
	FileMgrCollection,
	error) {

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
		errorPrefix,
		"DirMgr."+
			"FindFilesByNamePattern()",
		"")

	if err != nil {
		return fileMgrCol, err
	}

	return new(dirMgrHelper).findFilesByNamePattern(
		dMgr,
		fileSearchPattern,
		"dMgr",
		"fileSearchPattern",
		ePrefix)
}

// FindFilesBySelectCriteria
//
// This method searches the directory identified by the
// current DirMgr instance using file selection criteria
// as defined by input parameter 'fileSelectCriteria'.
//
// The directory specified by the current instance of
// DirMgr is treated as the target directory for the
// purposes of this file search.
//
// The file search operation will always be restricted
// to the target directory identified by the current
// DirMgr instance. No subdirectories will be included
// in the file search operation.
//
// Files matching the "FileSectionCriteria" instance
// passed as an input parameter will be used to screen
// available files. Any files matching the file selection
// criteria will be returned in a File Manager Collection
// (Type: FileMgrCollection).
//
// Only matched files will be returned.
//
// The use of a 'FileSelectionCriteria' structure allows
// for very flexible and granular file searches.
//
// ------------------------------------------------------------------------
//
// # Input parameter
//
//	fileSelectCriteria			FileSelectionCriteria
//
//	  This input parameter should be configured with the
//	  desired file selection criteria. Files matching
//	  this criteria will be returned as 'Found Files'.
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
//		If all of the file selection criterion in the FileSelectionCriteria
//		object are 'Inactive' or 'Not Set' (set to their zero or default values),
//		then all the files processed in the directory tree will be selected and
//		returned as 'Found Files'.
//
//			Example:
//				fsc := FileSelectCriterionMode{}
//
//				In this example, 'fsc' is NOT initialized. Therefore,
//				all the selection criterion are 'Inactive'. Consequently,
//				all the files encountered in the target directory during
//				the search operation will be selected and returned as
//				'Found Files'.
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
//	FileMgrCollection
//
//		If this method completes successfully without
//		error, the returned FileMgrCollection type will
//		contain an array of FileMgr types identifying
//		each of the files matched by input parameter,
//		'fileSelectCriteria'.
//
//		In summary, this returned File Manager Collection
//		will contain information on all files matching
//		the file selection criteria in the directory
//		specified by the current instance of DirMgr.
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
func (dMgr *DirMgr) FindFilesBySelectCriteria(
	fileSelectCriteria FileSelectionCriteria,
	errorPrefix interface{}) (
	FileMgrCollection,
	error) {

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
		errorPrefix,
		"DirMgr."+
			"FindFilesBySelectCriteria()",
		"")

	if err != nil {
		return dTreeInfo.FoundFiles, err
	}

	dMgrHlpr := new(dirMgrHelper)

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

// FindWalkDirFiles
//
// This method returns file information on files residing
// in a specific directory tree identified by the current
// DirMgr instance. The directory identified by 'DirMgr'
// is treated as the parent directory for the search.
//
// In addition to file information, this method also
// returns data on the directory tree being searched
// including the parent directory and all subdirectories
// in the tree.
//
// This method 'walks the directory tree' locating all
// files in the directory tree which match the file
// selection criteria submitted as input parameter,
// 'fileSelectCriteria'.
//
// All directories in the tree are searched including the
// top level parent directory identified by the current
// DirMgr instance. This differs from method
// 'DirMgr.FindWalkSubDirFiles()' which only searches the
// subdirectory tree.
//
// If a file matches the File Selection Criteria passed
// as input parameter 'fileSelectCriteria', that file is
// included in the returned DirectoryTreeInfo structure
// (DirectoryTreeInfo.FoundFiles). If ALL the file
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
//	fileSelectCriteria			FileSelectionCriteria
//
//		This input parameter should be configured with
//		the desired file selection criteria. Files
//		matching this criteria will be returned as
//		'Found Files'.
//
//			type FileSelectionCriteria struct {
//
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
//		If all of the file selection criterion in the FileSelectionCriteria
//		object are 'Inactive' or 'Not Set' (set to their zero or default values),
//		then all the files processed in the target directory tree will be
//		selected and returned as 'Found Files'.
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
//	DirectoryTreeInfo
//
//		If this method completes successfully, files
//		matching the file selection criteria input
//		parameter 'fileSelectCriteria' will be returned
//		in a 'DirectoryTreeInfo' object. The field
//		'DirectoryTreeInfo.FoundFiles' contains
//		information on all the files in the specified
//		directory tree which match the file selection
//		criteria.
//
//		Note: It is a good idea to check the returned field 'DirectoryTreeInfo.ErrReturns'
//			  to determine if any internal system errors were encountered while processing
//			  the directory tree.
//
//		This structure contains detailed information on
//		files found during the search operation.
//
//	          type DirectoryTreeInfo struct {
//
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
//		In addition, internal system errors are
//	 	documented in 'DirectoryTreeInfo.ErrReturns' as
//	 	shown above.
func (dMgr *DirMgr) FindWalkDirFiles(
	fileSelectCriteria FileSelectionCriteria,
	errorPrefix interface{}) (
	DirectoryTreeInfo,
	error) {

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
		errorPrefix,
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

// FindWalkSubDirFiles
//
// This method returns file information on files residing
// in a subdirectory tree identified by the current
// DirMgr instance. As such, this method will NOT search
// the top level directory, parent directory identified
// by the current DirMgr instance. However, all
// directories subsidiary to the parent directory
// ('DirMgr') will be searched.
//
// This method 'walks the directory tree' locating all
// files in the subdirectory tree which match the file
// selection criteria submitted as input parameter,
// 'fileSelectCriteria'.
//
// If a file matches the File Selection Criteria, it is
// included in the returned field,
// 'DirectoryTreeInfo.FoundFiles'. If ALL the file
// selection criterion are set to zero values or
// 'Inactive', then ALL FILES in the suc-directory tree
// are selected and returned in the field,
// 'DirectoryTreeInfo.FoundFiles'.
//
// All directories searched will be included in the
// returned collection 'DirectoryTreeInfo.Directories'.
// If the parent directory has NO subdirectories, this
// returned collection will be empty.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	The directory identified by the current instance of
//	DirMgr is treated as the parent directory.
//
//	The file search conducted by this method will NOT
//	include the parent directory. ONLY subdirectories
//	of the parent directory will be included in the file
//	search operation.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fileSelectCriteria			FileSelectionCriteria
//
//		This input parameter should be configured with
//		the desired file selection criteria. Files
//		matching this criteria will be returned as
//		'Found Files'.
//
//			type FileSelectionCriteria struct {
//
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
//		If all of the file selection criterion in the FileSelectionCriteria
//		object are 'Inactive' or 'Not Set' (set to their zero or default values),
//		then all the files processed in the target directory tree will be
//		selected and returned as 'Found Files'.
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
//	DirectoryTreeInfo
//
//		If this method completes successfully, files
//		matching the file selection criteria input
//		parameter 'fileSelectCriteria' will be returned
//		in a 'DirectoryTreeInfo' object. The field
//		'DirectoryTreeInfo.FoundFiles' contains
//		information on all the files in the specified
//		directory tree which match the file selection
//		criteria.
//
//		Note: It is a good idea to check the returned field 'DirectoryTreeInfo.ErrReturns'
//			  to determine if any internal system errors were encountered while processing
//			  the directory tree.
//
//		This structure contains detailed information on
//		files found during the search operation.
//
//	          type DirectoryTreeInfo struct {
//
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
//		In addition, internal system errors are
//	 	documented in 'DirectoryTreeInfo.ErrReturns' as
//	 	shown above.
func (dMgr *DirMgr) FindWalkSubDirFiles(
	fileSelectCriteria FileSelectionCriteria,
	errorPrefix interface{}) (
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
		errorPrefix,
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

	return dTreeInfo, err
}

// GetAbsolutePath
//
// Returns a string containing the absolute path for the
// current Directory Manager instance (DirMgr). This
// string returned by this method WILL NEVER contain a
// trailing path separator (Linux='/' or Windows='\').
//
// The returned absolute path string may consist of upper
// and lower case characters if the current DirMgr
// path was initialized with upper and lower case
// characters.
//
// See companion method GetAbsolutePathLc() to
// acquire a lower case version of absolute path.
//
// ----------------------------------------------------------------
//
// # Definition of Terms
//
// An absolute or full path points to the same location
// in a file system, regardless of the current working
// directory. To do that, it must include the root
// directory.
//
//	https://en.wikipedia.org/wiki/Path_(computing)#Absolute_and_relative_paths
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	(1)	The absolute path string returned by this method
//		WILL NEVER contain a trailing path separator
//		(Linux='/' or Windows='\').
//
//	(2)	The absolute path string returned by this method
//		may consist of upper and lower case characters if
//		the current DirMgr path was initialized with
//		upper and lower case characters.
//
//	(3) If this method returns an empty string, it
//		signals that some type of error was encountered.
//		To examine detailed error messages, call method:
//			DirMgr.DoesThisDirectoryExist()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	--- NONE ---
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	string
//
//		This method returns a string containing the
//		absolute path specified by the current instance
//		of DirMgr. Remember that the returned absolute
//		path string WILL NEVER contain a trailing path
//		separator (Linux='/' or Windows='\').
//
//		The absolute path string returned by this
//		parameter may consist of upper and lower case
//		characters if the current DirMgr path was
//		initialized with upper and lower case characters.
//
//		If this method returns an empty string it signals
//		that some type of error was encountered. To
//		examine detailed error messages, call method:
//			DirMgr.DoesThisDirectoryExist()
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

	if err == nil {

		absolutePath = dMgr.absolutePath

	}

	return absolutePath
}

// GetAbsolutePathLc
//
// Returns a string containing the lower case version of
// the absolute path specified by the current Directory
// Manager instance (DirMgr).
//
// This string returned by this method will NOT have a
// trailing path separator (Linux='/' or Windows='\').
// And, it will consist of all lower case characters.
//
// See the companion method GetAbsolutePath() to return
// an absolute path string with upper and lower case
// characters.
//
// ----------------------------------------------------------------
//
// # Definition of Terms
//
// An absolute or full path points to the same location
// in a file system, regardless of the current working
// directory. To do that, it must include the root
// directory.
//
//	https://en.wikipedia.org/wiki/Path_(computing)#Absolute_and_relative_paths
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	(1)	The absolute path string returned by this method
//		will NOT contain a trailing path separator
//		(Linux='/' or Windows='\').
//
//	(2)	All characters returned by this method in the
//		absolute path string are lower case.
//
//	(3) If this method returns an empty string, it
//		signals that some type of error was encountered.
//		To examine detailed error messages, call method:
//			DirMgr.DoesThisDirectoryExist()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	--- NONE ---
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	string
//
//		This method returns a string containing the
//		absolute path specified by the current instance
//		of DirMgr. Remember that the returned absolute
//		path string will NOT contain a trailing path
//		separator (Linux='/' or Windows='\').
//
//		All characters returned in this string are lower
//		case.
//
//		If this method returns an empty string it signals
//		that some type of error was encountered. To
//		examine detailed error messages, call method:
//			DirMgr.DoesThisDirectoryExist()
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

	if err == nil {

		absolutePath = strings.ToLower(dMgr.absolutePath)
	}

	return absolutePath
}

// GetAbsolutePathElements
//
// Returns all the drive and directory elements from
// the absolute path specified by the current instance of
// DirMgr. Each component of the absolute path is
// isolated and returned as a single element in an array
// of strings.
//
// ----------------------------------------------------------------
//
// # Definition of Terms
//
// An absolute or full path points to the same location
// in a file system, regardless of the current working
// directory. To do that, it must include the root
// directory.
//
//	https://en.wikipedia.org/wiki/Path_(computing)#Absolute_and_relative_paths
//
// ----------------------------------------------------------------
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
//	pathElements				[]string
//
//		Returns the absolute path specified by the
//		current instance of DirMgr as individual drive
//		and directory elements in an array of strings.
//
//		See the example path elements array shown above.
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
func (dMgr *DirMgr) GetAbsolutePathElements(
	errorPrefix interface{}) (
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
		errorPrefix,
		"DirMgr."+
			"GetAbsolutePathElements()",
		"")

	if err != nil {

		return pathElements, err
	}

	pathElements,
		err = new(dirMgrHelper).
		getAbsolutePathElements(
			dMgr,
			"",
			ePrefix.XCpy("dMgr"))

	return pathElements, err
}

// GetAbsolutePathWithSeparator
//
// Returns the absolute path specified by the current
// DirMgr instance. The returned path WILL ALWAYS contain
// a trailing os.PathSeparator character (Linux='/' or
// Windows='\').
//
// The returned absolute path string may consist of upper
// and lower case characters if the current DirMgr
// path was initialized with upper and lower case
// characters.
//
// See the companion method
// DirMgr.GetAbsolutePathWithSeparatorLc() which returns a
// path string consisting of all lower case characters.
//
// ----------------------------------------------------------------
//
// # Definition of Terms
//
// An absolute or full path points to the same location
// in a file system, regardless of the current working
// directory. To do that, it must include the root
// directory.
//
//	https://en.wikipedia.org/wiki/Path_(computing)#Absolute_and_relative_paths
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	(1)	The absolute path string returned by this method
//		WILL ALWAYS contain a trailing os.PathSeparator
//		character (Linux='/' or Windows='\').
//
//	(2)	The absolute path string returned by this method
//		may consist of upper and lower case characters if
//		the current DirMgr path was initialized with
//		upper and lower case characters.
//
//	(3) If this method returns an empty string, it
//		signals that some type of error was encountered.
//		To examine detailed error messages, call method:
//			DirMgr.DoesThisDirectoryExist()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	--- NONE ---
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	string
//
//		This method returns a string containing the
//		absolute path specified by the current instance
//		of DirMgr. Remember that the returned absolute
//		path string WILL ALWAYS contain a trailing
//		os.PathSeparator character (Linux='/' or
//		Windows='\').
//
//		The absolute path string returned by this
//		parameter may consist of upper and lower case
//		characters if the current DirMgr path was
//		initialized with upper and lower case characters.
//
//		If this method returns an empty string it signals
//		that some type of error was encountered. To
//		examine detailed error messages, call method:
//			DirMgr.DoesThisDirectoryExist()
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

	if err == nil {
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

// GetAbsolutePathWithSeparatorLc
//
// Returns the absolute path specified by the current
// DirMgr instance. The returned path WILL ALWAYS contain
// a trailing os.PathSeparator character (Linux='/' or
// Windows='\').
//
// The returned absolute path string WILL ALWAYS consist
// of lower case characters.
//
// See the companion method
// DirMgr.GetAbsolutePathWithSeparator() which returns a
// path string containing upper and lower case characters.
//
// ----------------------------------------------------------------
//
// # Definition of Terms
//
// An absolute or full path points to the same location
// in a file system, regardless of the current working
// directory. To do that, it must include the root
// directory.
//
//	https://en.wikipedia.org/wiki/Path_(computing)#Absolute_and_relative_paths
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	(1)	The absolute path string returned by this method
//		WILL ALWAYS contain a trailing os.PathSeparator
//		character (Linux='/' or Windows='\').
//
//	(2)	The absolute path string returned by this method
//		WILL ALWAYS consist of lower case characters.
//
//	(3) If this method returns an empty string, it
//		signals that some type of error was encountered.
//		To examine detailed error messages, call method:
//			DirMgr.DoesThisDirectoryExist()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	--- NONE ---
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	string
//
//		This method returns a string containing the
//		absolute path specified by the current instance
//		of DirMgr. Remember that the returned absolute
//		path string WILL ALWAYS contain a trailing
//		os.PathSeparator character (Linux='/' or
//		Windows='\').
//
//		The absolute path string returned by this
//		parameter WILL ALWAYS consist of lower case
//		characters.
//
//		If this method returns an empty string it signals
//		that some type of error was encountered. To
//		examine detailed error messages, call method:
//			DirMgr.DoesThisDirectoryExist()
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

	if err == nil {
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

// GetDirectoryProfile
//
// This method returns an instance of DirectoryProfile
// which includes file classifications and statistics on
// the directory specified by the current instance of DirMgr.
//
// The directory analysis required to generate this Directory
// Profile is designed to accomplish the following objectives.
//
//	(1)	Determine if the directory path exists on an
//		attached storage drive.
//
//	(2)	If the path does exist, statistics on the
//		directory will be generated and returned via an
//		instance of	DirectoryProfile.
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//	dirProfile					DirectoryProfile
//
//		If this method completes successfully, this
//		returned instance of DirectoryProfile will be
//		populated with profile and statistical
//		information on the directory identified by the
//		current instance of DirMgr.
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
//			Errors error
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
func (dMgr *DirMgr) GetDirectoryProfile(
	errorPrefix interface{}) (
	directoryPathDoesExist bool,
	dirProfile DirectoryProfile,
	err error) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "DirMgr." +
		"GetDirectoryProfile()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {
		return directoryPathDoesExist, dirProfile, err
	}

	directoryPathDoesExist,
		dirProfile,
		err = new(dirMgrHelperTachyon).
		getDirectoryProfile(
			dMgr,
			"dMgr",
			ePrefix.XCpy("dMgr"))

	return directoryPathDoesExist, dirProfile, err
}

// GetDirectoryStats
//
// Returns statistics and information on the directory
// structure identified by the current instance of
// DirMgr.
//
// For the purposes of this method, the directory
// specified by the current DirMgr instance is treated
// as the top-level or 'parent' directory.
//
// This method returns an instance of DirectoryStatsDto
// which contains information and statistics on the
// specified directory structure.
//
// The statistics and information returned will cover the
// parent directory, the parent directory and
// subdirectories or subdirectories exclusively,
// depending on the input parameters
// 'includeParentDirectory' and 'includeSubDirectories'.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	(1)	Be careful to understand the settings for input
//		parameters 'includeParentDirectory' and
//		'includeSubDirectories' in the context of the
//		returned directory statistics. For example,
//		if 'includeParentDirectory' is set to 'false' and
//		'includeSubDirectories' is set to 'true', the
//		returned directory statistics will ONLY include
//		the subdirectory tree and files in the
//		subdirectory tree.
//
//	(1)	If 'includeSubDirectories' is set to 'false' and
//		'includeSubDirectories' is set to 'false' an
//		error will be returned. These settings specify
//		that NO directories will be included in the
//	 	accumulated directory statistics. This represents
//	 	a NULL set and an error condition.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	includeParentDirectory		bool
//
//		If this parameter is set to 'true' it signals
//		that the top-level or parent directory
//		specified by the current DirMgr instance will
//		be included in the returned directory statistics.
//
//		If both parameters 'includeParentDirectory' and
//		'includeSubDirectories' are set to 'false', it
//		defines a NULL set and an error will be returned.
//
//	includeSubDirectories bool
//
//		If this parameter is set to 'true' it signals
//		that the subdirectory tree relative to the
//		parent directory specified by the current DirMgr
//		instance will be included in the returned
//		directory statistics.
//
//		If both parameters 'includeParentDirectory' and
//		'includeSubDirectories' are set to 'false', it
//		defines a NULL set and an error will be returned.
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
//	dTreeStats					DirectoryStatsDto
//
//		If this method completes successfully, this
//		returned instance of DirectoryStatsDto will
//		contain information on files and directories
//		contained in the directory tree specified by
//		the current instance of DirMgr and the input
//		parameters 'includeParentDirectory' and
//		'includeSubDirectories'.
//
//		The DirectoryStatsDto structure is used to
//		accumulate and disseminate statistical
//		information relating to a specific directory
//		tree.
//
//			type DirectoryStatsDto struct {
//				numOfFiles    uint64
//				numOfSubDirs  uint64
//				numOfBytes    uint64
//				isInitialized bool // 'false' = error condition
//			}
//
//		Type DirectoryStatsDto contains public methods
//		for retrieving the specified directory statistics
//		and information.
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
func (dMgr *DirMgr) GetDirectoryStats(
	includeParentDirectory bool,
	includeSubDirectories bool,
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

	if includeParentDirectory == false &&
		includeSubDirectories == false {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameters 'includeParentDirectory'\n"+
			"and 'includeSubDirectories' are conflicted.\n"+
			"includeParentDirectory= 'false'\n"+
			"includeSubDirectories= 'false'\n"+
			"This represents a NULL set and an error condition.\n"+
			"In means that neither the parent directory nor the\n"+
			"the subdirectory tree will be included in the returned\n"+
			"directory statistics.\n",
			ePrefix.String())

		errs = append(errs, err)

		return dirStats, errs
	}

	dirStats,
		errs = new(dirMgrHelper).findDirectoryTreeStats(
		dMgr,
		!includeParentDirectory,
		includeSubDirectories,
		"dMgr",
		ePrefix)

	return dirStats, errs
}

// GetDirectoryTree
//
// Returns a Directory Manager Collection
// (DirMgrCollection) containing all the subdirectories
// in the path of the parent directory identified by the
// current DirMgr instance.
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
//		Type DirMgrCollection contains public methods
//		retrieving DirMgr objects from the Directory
//		Manager Collection
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

// GetDirectoryName
//
// Returns a string containing the name of the directory
// without the parent path.
//
// ----------------------------------------------------------------
//
// # Usage
//
//	--------------------
//	Example-1 (Windows)
//	--------------------
//
//	var dMgr DirMgr
//	var err error
//	dirPathString := "C:\t01\MyDirectory"
//
//	dMgr,
//	err = new(DirMgr).New(
//			dirPathString,
//			"CallingMethod()")
//
//	if err == nil {
//		fmt.Println(err.Error())
//	}
//
//	var directoryName string
//
//	directoryName = dMgr.GetDirectoryName()
//
//	'directoryName' is now equal to "MyDirectory"
//
//	--------------------
//	Example-2 (Linux)
//	--------------------
//
//	var dMgr DirMgr
//	var err error
//	dirPathString := "/c/t01/mydirectory"
//
//	dMgr,
//	err = new(DirMgr).New(
//			dirPathString,
//			"CallingMethod()")
//
//	if err == nil {
//		fmt.Println(err.Error())
//	}
//
//	var directoryName string
//
//	directoryName = dMgr.GetDirectoryName()
//
//	'directoryName' is now equal to "mydirectory"
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	--- NONE ---
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	string
//
//		This method returns the last directory name
//		element extracted from the directory path
//		specified by the current instance of DirMgr.
//
//		See examples shown above.
func (dMgr *DirMgr) GetDirectoryName() string {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	if !dMgr.isInitialized {

		return ""

	}

	_,
		_,
		_ = new(dirMgrHelperAtom).doesDirectoryExist(
		dMgr,
		PreProcPathCode.None(),
		"dMgr",
		nil)

	return dMgr.directoryName
}

// GetFileInfoPlus
//
// Returns a FileInfoPlus instance detailing file system
// information on the directory identified by the current
// Directory Manager instance (DirMgr).
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
//	fileInfoPlus				FileInfoPlus
//
//		This method returns an instance of FileInfoPlus
//		which details information on the directory
//		identified by the current instance of DirMgr.
//
//		This returned FileInfoPlus instance implements
//		the	os.FileInfo interface. In addition,
//		FileInfoPlus provides additional data on the
//		current Directory Manager instance beyond that
//		provided by the standard os.FileInfo interface.
//
//		The os.FileInfo interface is defined as follows:
//
//		type FileInfo interface {
//			Name() string       // base name of the file
//			Size() int64        // length in bytes for regular files; system-dependent for others
//			Mode() FileMode     // file mode bits
//			ModTime() time.Time // modification time
//			IsDir() bool        // abbreviation for Mode().IsDir()
//			Sys() any           // underlying data source (can return nil)
//		}
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

// GetDirPermissionCodes
//
// If the current directory exists on disk, this method
// will return the Directory Permission Codes
// encapsulated in an instance of type
// 'FilePermissionConfig'.
//
// If the current Directory does NOT exist, this method
// will return an error.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://golang.org/pkg/os/#FileMode
//	https://www.cyberciti.biz/faq/explain-the-nine-permissions-bits-on-files/
//	https://en.wikipedia.org/wiki/File_system_permissions
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
//	FilePermissionConfig
//
//		An instance of FilePermissionConfig containing
//		the permission specifications associated with
//		the directory identified by the current instance
//		of DirMgr.
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
//		Example: -rwxrwxrwx - Identifies permissions for a regular file
//		         drwxrwxrwx - Identifies permissions for directory
//		                      value = 020000000777
//
//		----------------------------------------------------------------
//
//			Symbolic and Numeric Notation
//
//		Permission codes may be designated with Symbolic
//		Notation or Numeric Octal Notation.
//
//						   Numeric
//				Symbolic	Octal
//				Notation   Notation
//				----------	0000	no permissions
//				-rwx------	0700	read, write, & execute only for owner
//				-rwxrwx---	0770	read, write, & execute for owner and group
//				-rwxrwxrwx	0777	read, write, & execute for owner, group and others
//				---x--x--x	0111	execute
//				--w--w--w-	0222	write
//				--wx-wx-wx	0333	write & execute
//				-r--r--r--	0444	read
//				-r-xr-xr-x	0555	read & execute
//				-rw-rw-rw-	0666	read & write
//				-rwxr-----	0740	owner can read, write, & execute; group can only read;
//			                       others have no permissions
//
//		See the source code documentation for method
//		'FilePermissionConfig.New()' for a detailed
//		explanation of permission codes.
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

// GetNumberOfAbsPathElements
//
// Returns the number of elements or path components in
// the absolute path of the current Directory Manager
// instance (DirMgr).
//
// ----------------------------------------------------------------
//
// # Definition of Terms
//
// An absolute or full path points to the same location
// in a file system, regardless of the current working
// directory. To do that, it must include the root
// directory.
//
//	https://en.wikipedia.org/wiki/Path_(computing)#Absolute_and_relative_paths
//
// ----------------------------------------------------------------
//
// # Example
//
//					Number of Path Elements
//	                1   2    3    4    5
//		DirMgr Path = "D:\ADir\BDir\CDir\EDir"
//
//		The number of path elements returned by this method
//		given the absolute path specified above is '5'.
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
//	int
//
//		This method returns the number of path elements
//		in the absolute path specified by the current
//		instance of DirMgr.
//
//		See the example shown above.
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

// GetOriginalPath
//
// Returns the original path used to initialize this
// Directory Manager instance. This returned path may
// be an absolute path or a relative path depending on
// how the DirMgr instance was initialized.
//
// ----------------------------------------------------------------
//
// # Definition of Terms
//
//	An absolute or full path points to the same location
//	in a file system, regardless of the current working
//	directory. To do that, it must include the root
//	directory.
//
//	By contrast, a relative path starts from some given
//	working directory, avoiding the need to provide the
//	full absolute path. A filename can be considered as
//	a relative path based at the current working directory.
//
//	https://en.wikipedia.org/wiki/Path_(computing)#Absolute_and_relative_paths
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	--- NONE ---
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	string
//
//		This string returns the original path used to
//		initialize the current Directory Manager
//		instance.
//
//		The returned path may be an absolute path or a
//		relative path depending on how the current DirMgr
//		instance was initialized.
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

// GetOriginalAbsolutePath
//
// Returns the original path used to initialize this
// Directory Manager instance as an absolute path.
//
// This method differs from DirMgr.GetOriginalPath in
// that this method will ALWAYS return the absolute
// path used to initialize the current instance of
// DirMgr.
//
// If the current instance of DirMgr has not been
// correctly initialized, an error will be returned.
//
// ----------------------------------------------------------------
//
// # Definition of Terms
//
// An absolute or full path points to the same location
// in a file system, regardless of the current working
// directory. To do that, it must include the root
// directory.
//
//	https://en.wikipedia.org/wiki/Path_(computing)#Absolute_and_relative_paths
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
//		This string returns the original path used to
//		initialize the current Directory Manager
//		instance.
//
//		The returned path may be an absolute path or a
//		relative path depending on how the current DirMgr
//		instance was initialized.
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
func (dMgr *DirMgr) GetOriginalAbsolutePath(
	errorPrefix interface{}) (
	string,
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
		"DirMgr.GetOriginalAbsolutePath()",
		"")

	if err != nil {
		return "", err
	}

	if !dMgr.isInitialized {

		err = fmt.Errorf("%v\n"+
			"Error: The current instance DirMgr\n"+
			"was NOT correctly initialized.\n",
			ePrefix.String())

		return "", err

	}

	validPathDto := new(ValidPathStrDto).New()

	validPathDto,
		err =
		new(dirMgrHelperMolecule).
			getValidPathStr(
				dMgr.originalPath,
				"dMgr.originalPath",
				ePrefix)

	if err != nil {
		return "", err
	}

	return validPathDto.absPathStr, err
}

// GetParentDirMgr
//
// Returns a new Directory Manager instance which
// represents the parent path for the current Directory
// Manager (DirMgr). The current Directory Manager
// absolute path is used in extracting the parent
// Directory Manager.
//
// ----------------------------------------------------------------
//
// # Definition of Terms
//
// An absolute or full path points to the same location
// in a file system, regardless of the current working
// directory. To do that, it must include the root
// directory.
//
//	https://en.wikipedia.org/wiki/Path_(computing)#Absolute_and_relative_paths
//
// ----------------------------------------------------------------
//
// # Usage
//
//	--------------------
//	Example-1 (Windows)
//	--------------------
//
//		Current DirMgr Absolute Path =
//			"D:\ADir\BDir\CDir\EDir"
//
//		The returned Parent DirMgr will be configured
//		with the absolute path:
//
//			"D:\ADir\BDir\CDir\EDir"
//
//	--------------------
//	Example-2 (Linux)
//	--------------------
//
//		Current DirMgr Absolute Path =
//			"/d/adir/bdir/cdir/edir"
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
//	dirMgrParent				DirMgr
//
//		If this method completes successfully, this
//		returned instance of DirMgr will contain the
//		absolute parent path extracted from the current
//		instance of DirMgr.
//
//	hasParent					bool
//
//		If the current DirMgr instance has a parent path,
//		this returned boolean value will be set to 'true'.
//
//		Otherwise, this value is set to 'false' signaling
//		that the current DirMgr instance does not have a
//		parent path and the absolute path specified by
//		the current DirMgr instance is a parent path.
//
//			Example: "D:\MyDir"
//
//				The original absolute path is "D:\MyDir"
//				and the parent absolute path is also
//				"D:\MyDir".
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
// current Directory Manager instance (DirMgr).
//
// The returned Parent Path string will NOT contain a
// trailing os.PathSeparator character (Linux='/' or
// Windows='\').
//
// The returned path will be an absolute path.
//
// ----------------------------------------------------------------
//
// # Definition of Terms
//
//	Absolute Path
//
//		An absolute or full path points to the same location
//		in a file system, regardless of the current working
//		directory. To do that, it must include the root
//		directory.
//
//		https://en.wikipedia.org/wiki/Path_(computing)#Absolute_and_relative_paths
//
//	os.PathSeparator
//
//		The type of path separator applied depends on the
//		operating system. Linux='/' or Windows='\'
//
//		https://pkg.go.dev/os#pkg-constants
//
// ----------------------------------------------------------------
//
// # Usage
//
//	--------------------
//	Example-1 (Windows)
//	--------------------
//
//		Current DirMgr Absolute Path =
//			"D:\ADir\BDir\CDir\EDir"
//
//		The returned Parent Absolute Path:
//
//			"D:\ADir\BDir\CDir"
//
//	--------------------
//	Example-2 (Linux)
//	--------------------
//
//		Current DirMgr Absolute Path =
//			"/d/adir/bdir/cdir/edir"
//
//		The returned Parent Absolute Path:
//
//			"/d/adir/bdir/cdir"
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
//	 	absolute parent path for the directory specified
//	 	by the current instance of Directory Manager
//	 	(DirMgr).
//
//		Be advised that the returned path will NOT
//		contain a trailing path separator
//		(os.PathSeparator).
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

// GetPath
//
// Returns a string containing the directory path used to
// configure the current Directory Manager Instance
// (DirMgr).
//
// The returned directory path string will NOT contain a
// trailing os.PathSeparator character (Linux='/' or
// Windows='\').
//
// In addition, the returned directory path WILL ALWAYS
// be an absolute path.
//
// ----------------------------------------------------------------
//
// # Definition of Terms
//
//	Absolute Path
//
//		An absolute or full path points to the same
//		location in a file system, regardless of the
//		current working directory. To do that, it must
//		include the root directory.
//
//		https://en.wikipedia.org/wiki/Path_(computing)#Absolute_and_relative_paths
//
//	os.PathSeparator
//
//		The type of path separator applied depends on the
//		operating system. Linux='/' or Windows='\'
//
//		https://pkg.go.dev/os#pkg-constants
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	--- NONE ---
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	string
//
//		A string containing the directory path used to
//		configure the current Directory Manager Instance
//		(DirMgr).
//
//		The returned directory path string will NOT
//		contain a trailing os.PathSeparator character
//		(Linux='/' or Windows='\').
//
//		The returned path will be an absolute path.
func (dMgr *DirMgr) GetPath() string {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

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

		return ""

	}

	return dMgr.path
}

// GetPathWithSeparator
//
// Returns the absolute directory path representing the
// directory specified by the current instance of DirMgr.
//
// The returned path WILL ALWAYS contain a trailing
// os.PathSeparator character (Linux='/' or Windows='\').
//
// ----------------------------------------------------------------
//
// # Definition of Terms
//
//	Absolute Path
//
//		An absolute or full path points to the same location
//		in a file system, regardless of the current working
//		directory. To do that, it must include the root
//		directory.
//
//		https://en.wikipedia.org/wiki/Path_(computing)#Absolute_and_relative_paths
//
//	os.PathSeparator
//
//		The type of path separator applied depends on the
//		operating system. Linux='/' or Windows='\'
//
//		https://pkg.go.dev/os#pkg-constants
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	--- NONE ---
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	string
//
//		This string parameter returns the absolute
//		directory path specified by the current instance
//		of DirMgr.
//
//		The returned absolute directory path WILL ALWAYS
//		contain a trailing os.PathSeparator character
//		(Linux='/' or Windows='\').
func (dMgr *DirMgr) GetPathWithSeparator() string {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()
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

		return ""

	}

	dPath := dMgr.path

	lPath := len(dPath)

	if lPath == 0 {
		return ""
	}

	if dPath[lPath-1] != os.PathSeparator {
		return dPath + string(os.PathSeparator)
	}

	return dPath
}

// GetSubdirectoriesDirTree
//
// This method scans and identifies all the
// subdirectories in the entire directory tree
// defined by the current instance of DirMgr.
// Subdirectories located in this directory tree are
// returned to the user by means of a Directory Manager
// Collection (DirMgrCollection) passed as input
// parameter 'subDirectories'.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	(1)	All subdirectories residing in all levels of the
//		directory tree defined by the current instance of
//		DirMgr will be added and returned in the
//		Directory Manager Collection passed as input
//		parameter 'subDirectories'.
//
//	(2)	While scanning for subdirectories, Directory
//		entries for the current directory (".") and the
//		parent directory ("..") will be skipped and will
//		NOT be added to the 'subDirectories' Directory
//		Manager Collection.
//
//	(3)	For a collection of subdirectories residing in the
//		top level or parent directory specified by the
//		current DirMgr instance, see method:
//
//			DirMgr.GetSubdirectoriesParentDir
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	subDirectories				*DirMgrCollection
//
//		A pointer to an instance of DirMgrCollection
//		which encapsulates an array of Directory Manager
//		(DirMgr) objects.
//
//		This method will scan the entire directory tree
//		defined by the current DirMgr instance.	All
//		subdirectories found in this directory tree will
//		be configured as Directory Manager (DirMgr)
//		objects and added to this Directory Manager
//		Collection ('subDirectories').
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
//		defined by the current instance of DirMgr.
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
func (dMgr *DirMgr) GetSubdirectoriesDirTree(
	subDirectories *DirMgrCollection,
	errorPrefix interface{}) (
	numOfSubdirectories int,
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
			"GetSubdirectoriesDirTree()",
		"")

	if err != nil {
		return numOfSubdirectories, err
	}

	if subDirectories == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'subDirectories' is invalid!\n"+
			"'subDirectories' is a nil pointer.\n",
			ePrefix.String())

		return numOfSubdirectories, err
	}

	numOfSubdirectories,
		err = new(dirMgrHelperElectron).
		getAllSubDirsInDirTree(
			dMgr,
			subDirectories,
			"dMgr",
			ePrefix)

	return numOfSubdirectories, err
}

// GetSubdirectoriesParentDir
//
// This method scans and identifies all the
// subdirectories in the top level or parent directory
// specified by the current instance of DirMgr. These
// subdirectories are returned to the user by means of
// a Directory Manager Collection (DirMgrCollection)
// passed as input parameter 'subDirectories'.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	(1)	Only the subdirectories residing in the single
//		parent directory defined by the current instance
//		of DirMgr will be added and returned in the
//		Directory Manager Collection passed as input
//		parameter 'subDirectories'.
//
//	(2)	While scanning for subdirectories, Directory entries
//		for the current directory (".") and the parent directory
//		("..") will be skipped and will NOT be added to the
//		'subDirectories' Directory Manager Collection.
//
//	(3)	For a collection of all subdirectories in the
//		directory tree specified by the current instance
//		of DirMgr, see method:
//
//			DirMgr.GetSubdirectoriesDirTree
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	subDirectories				*DirMgrCollection
//
//		A pointer to an instance of DirMgrCollection
//		which encapsulates an array of Directory Manager
//		(DirMgr) objects.
//
//		This method will scan the top level or parent
//		directory defined by the current DirMgr instance.
//		Any subdirectories found in this parent directory
//		will be configured as Directory Manager (DirMgr)
//		objects and added to this Directory Manager
//		Collection ('subDirectories').
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
//		If this method completes successfully, this
//		integer value represents the number of
//		subdirectories added to the Directory Manager
//		Collection passed as input parameter
//		'subDirectories'.
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
func (dMgr *DirMgr) GetSubdirectoriesParentDir(
	subDirectories *DirMgrCollection,
	errorPrefix interface{}) (
	numOfSubdirectories int,
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
			"GetSubdirectoriesParentDir()",
		"")

	if err != nil {
		return numOfSubdirectories, err
	}

	if subDirectories == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'subDirectories' is invalid!\n"+
			"'subDirectories' is a nil pointer.\n",
			ePrefix.String())

		return numOfSubdirectories, err
	}

	numOfSubdirectories,
		err = new(dirMgrHelperPreon).
		getSubdirectories(
			dMgr,
			subDirectories,
			"dMgr",
			ePrefix)

	return numOfSubdirectories, err
}

// GetVolumeName
//
// Returns a string containing the volume name of the
// directory identified by the current Directory Manager
// instance (DirMgr).
//
// If the current instance of DirMgr was improperly
// initialized, an empty string will be returned.
//
// ----------------------------------------------------------------
//
// # Definition Of Terms:
//
//	Volume Name
//		VolumeName returns leading volume name.
//		Given "C:\foo\bar" it returns "C:" on Windows.
//		Given "\\host\share\foo" it returns "\\host\share".
//		On other platforms it returns "".
//
//	https://pkg.go.dev/path/filepath#VolumeName
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	--- NONE ---
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	string
//
//		This method returns a string containing the
//		volume name associated with the directory
//		specified by the current instance of DirMgr.
//		Reference the definition of Volume Name above.
//
//		If the current instance of DirMgr was improperly
//		initialized, an empty string will be returned.
func (dMgr *DirMgr) GetVolumeName() string {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

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

		return ""

	}

	return dMgr.volumeName
}

// IsValidInstance
//
// Performs a diagnostic review of the data values
// encapsulated in the current DirMgr instance to
// determine if they are valid.
//
// If any data element evaluates as invalid, this
// method will return a boolean value of 'false'.
//
// If all data elements are determined to be valid in all
// respects, this method returns a boolean value of
// 'true'.
//
// This method is functionally equivalent to
// DirMgr.IsValidInstanceError() with the sole exceptions
// being that this method takes no input parameters and
// returns a boolean value.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	-- NONE --
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		If any of the internal member data values
//		contained in the current instance of DirMgr are
//		found to be invalid, this method will return a
//		boolean value of 'false'.
//
//		If all internal member data variables contained
//		in the current instance of DirMgr are found to be
//		valid in all respects, this method returns a
//		boolean value of 'true'.
func (dMgr *DirMgr) IsValidInstance() bool {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	var err error

	_,
		_,
		err = new(dirMgrHelperPreon).validateDirMgr(
		dMgr,
		false, // Path is NOT required to exit on disk
		"dMgr",
		nil)

	if err != nil {

		return false
	}

	return true
}

// IsValidInstanceError
//
// Performs a diagnostic review of the data values
// encapsulated in the current DirMgr instance to
// determine if they are valid in all respects.
//
// If any data element evaluates as invalid, this
// method will return an error containing an appropriate
// message describing the nature of the invalid error.
//
// If all member data elements are determined to be
// valid in all respects, this method returns a value
// of 'nil'.
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
func (dMgr *DirMgr) IsValidInstanceError(
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
		"DirMgr.IsValidInstanceError()",
		"")

	if err != nil {
		return err
	}

	_,
		_,
		err = new(dirMgrHelperPreon).validateDirMgr(
		dMgr,
		false, // Path is NOT required to exit on disk
		"dMgr",
		ePrefix)

	return err
}

// IsInitialized
//
// Returns a boolean value indicating whether the current
// Directory Manager instance (DirMgr) has been properly
// initialized.
//
// If this method returns a boolean value of 'false' it
// signals that the current DirMgr instance is malformed
// and may produce incorrect operational results.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	--- NONE ---
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		This method analyzes the current instance of
//		DirMgr to determine if it has been properly
//		initialized.
//
//		If the current DirMgr instance has been properly
//		initialize and is operationally ready, this
//		return parameter is set to 'true'.
//
//		If the analysis determines that the current
//		DirMgr not been properly initialized, this
//		return parameter is set to 'false'. A return
//		value of 'false' signals that the current DirMgr
//		is malformed and may produce incorrect
//		operational results.
func (dMgr *DirMgr) IsInitialized() bool {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	_,
		_,
		_ = new(dirMgrHelperAtom).
		doesDirectoryExist(
			dMgr,
			PreProcPathCode.None(),
			"dMgr",
			nil)

	return dMgr.isInitialized
}

// IsParentPathPopulated
//
// Returns a boolean value indicating whether the parent
// path for the current Directory Manager instance is
// populated.
//
// ----------------------------------------------------------------
//
// # Example - Parent Path
//
//	--------------------
//	Example-1 (Windows)
//	--------------------
//
//		Current DirMgr Absolute Path =
//			"D:\ADir\BDir\CDir\EDir"
//
//		The Parent Absolute Path:
//
//			"D:\ADir\BDir\CDir"
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	--- NONE ---
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		If the Parent Path for the current DirMgr
//		instance exists and properly configured,
//		this return parameter is set to 'true'.
//
//		Otherwise, a boolean value of 'false' is
//		returned.
func (dMgr *DirMgr) IsParentPathPopulated() bool {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	_,
		_,
		_ = new(dirMgrHelper).getParentDirMgr(
		dMgr,
		"dMgr",
		nil)

	if len(dMgr.parentPath) == 0 ||
		!dMgr.isInitialized {

		return false
	}

	return true
}

// IsPathPopulated
//
// Returns a boolean value indicating whether the current
// Directory Manager (DirMgr) path string is populated.
//
// The Director Manager is stored in private internal
// member variables. If the path has not been populated,
// it indicates that the current DirMgr instance is not
// serviceable and will produce invalid operational
// results.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	--- NONE ---
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		This return parameter is set to 'true' if the
//		directory path for the current DirMgr instance
//		has been correctly configured and populated.
//
//		Otherwise, a boolean value of 'false' is
//		returned.
func (dMgr *DirMgr) IsPathPopulated() bool {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()
	var err error
	_,
		_,
		err = new(dirMgrHelper).getParentDirMgr(
		dMgr,
		"",
		nil)

	if len(dMgr.path) == 0 ||
		!dMgr.isInitialized ||
		err != nil {

		return false

	}

	return true
}

// ParseValidPathStr
//
// Receives a valid path string and parses that string
// into is basic elements. Those elements are returned
// in an instance of type ValidPathStrDto.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathStr						string
//
//		This string contains a directory path which
//		will parsed into its basic elements and
//		returned as an instance of ValidPathStrDto.
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
//	ValidPathStrDto
//
//		The directory path passed as input parameter
//		'pathStr' will be broken down to its basic
//		path elements and returned as an instance
//		of ValidPathStrDto.
//
//		Type ValidPathStrDto encapsulates a series of
//		private member variables describing the basic
//		elements of a directory path. All of these
//		private member variables can be accessed through
//		public methods provided by Type ValidPathStrDto.
//
//		type ValidPathStrDto struct {
//			isInitialized bool
//			//	Signals whether the current ValidPathStrDto instance
//			//	has been properly initialized.
//
//			originalPathStr string
//			//	The original, unformatted path string
//
//			pathStr string
//			//	The path string which may or may not be
//			//	the absolute path.
//
//			pathFInfoPlus FileInfoPlus
//			// Only populated if absValidPath exists on disk.
//
//			pathDoesExist PathExistsStatusCode
//			//	-1 = don't know, file/path existence has not been tested
//			//	 0 - No, tests show the file/path doesn't exist on disk.
//			//	 1 - Yes, tests show the file/path does exist on disk.
//
//			pathStrLength int
//			// Length of the path string
//
//			absPathStr string
//			// The absolute path version of 'path'
//
//			absPathFInfoPlus FileInfoPlus
//			// Only populated if absPathStr
//			// exists on disk.
//
//			absPathDoesExist PathExistsStatusCode
//			//	-1 = don't know, has not been tested
//			//	 0 - No, tests shown path doesn't exist
//			//	 1 - Yes, tests show path does exist
//
//			absPathStrLength int
//			//	Length of the absolute path string
//
//			pathType PathFileTypeCode
//			//	The path type. Path File, Path Directory
//
//			pathIsValid PathValidityStatusCode
//			//	-1 - don't know
//			//	 0 - No path is NOT valid
//			//	 1 - Yes, path is valid
//
//			pathVolumeName string
//			//	Volume name associated with current path
//
//			pathVolumeIndex int
//			// Index of the starting character of Volume Name
//			// in the path string.
//
//			pathVolumeStrLength int
//			// Length of the Volume name in the path string.
//
//			err error
//			// If no error is encountered
//			// this value is nil
//		}
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
	validPathDto := new(ValidPathStrDto).New()

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgr."+
			"ParseValidPathStr()",
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

// IsVolumeNamePopulated
//
// Returns a boolean value indicating whether the
// Volume Name for the current Directory Manager
// instance is populated.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	--- NONE ---
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		This method returns a boolean value of 'true' if
//		the Volume Name for the current Directory Manager
//		instance is populated.
func (dMgr *DirMgr) IsVolumeNamePopulated() bool {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

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

		return false

	}

	return true
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

// MakeDir
//
// If the directory path identified by the current DirMgr
// object does not exist, this method will create that
// directory path.
//
// The permission specification used to create the
// directory is 'drwxrwxrwx' which is equivalent to octal
// value, '020000000777'
//
// MakeDir creates a directory named path, along with any
// necessary parent directories.
//
// If the directory creation fails, an error is returned.
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

// MoveDirectoryFiles
//
// This method will 'move' selected files from a source
// directory to a target directory.
//
// This move operation is accomplished in two steps by
// first copying selected source files to the target
// directory and then deleting the original source
// files. The copy step is executed using a Copy IO
// operation. For information on the Copy IO procedure
// see FileHelper{}.CopyFileByIo() method and reference:
//
//	https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// No subdirectories will be included in this move
// operation. Only source files in a single directory
// will be moved to the target directory.
//
// To qualify as a selected source file eligible for the
// move operation, a file must comply with two filters:
// File Type and File Characteristics.
//
// To be eligible for the move operation, a source file
// must first comply with the specified File Type
// criteria. The File Type Filter classifies artifacts
// residing in a parent directory as subdirectories,
// regular files, SymLink files or other non-regular
// files. Since this method does NOT move subdirectories,
// the only valid file types are Regular Files, SymLink
// Files and Other Non-Regular Files. For an explanation
// of Regular and Non-Regular files, see the Definition
// of Terms section below.
//
// Screening criteria for File Type is controlled by the
// following three input parameters:
//
//	moveRegularFiles bool
//	moveSymLinkFiles bool
//	moveOtherNonRegularFiles bool
//
// File Types eligible for this move operation therefore
// include Regular Files such as text files, image files
// and executable files, SymLink files and other Non-Regular
// Files such as device files, named pipes and sockets.
//
// In addition to File Type, selected files must also
// comply with the File Characteristics criteria
// specified by input parameter, 'fileSelectCriteria'.
// The File Characteristics Selection criteria allows
// users to screen files for File Name, File Modification
// Date and File Mode.
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
//	(1)	This method ONLY moves files from the single
//		parent source directory identified by the
//		current	instance of DirMgr. These selected source
//		files are moved to the target directory
//		identified by input parameter 'targetDMgr'.
//
//	(2)	This method does NOT move subdirectories residing
//		in parent source directory identified by the
//		current DirMgr instance. Therefore, it NEVER
//		moves files residing in subdirectories of the
//		parent directory identified by the current DirMgr
//		instance.
//
//	(3)	If the target directory does not exist, this method
//		will attempt to create it.
//
//	(4)	Files will only be moved if they meet the File Type
//		criteria and the File Characteristics Criteria.
//
//		File Type criteria are specified by input parameters:
//
//			moveRegularFiles bool
//			moveSymLinkFiles bool
//			moveOtherNonRegularFiles bool
//
//		File Characteristics criteria are specified by
//		input parameter 'fileSelectCriteria'.
//
//		Files will NOT be selected for the move operation
//		unless they satisfy both the File Type and File
//		Characteristics filters.
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
//	returnMovedFilesList		bool
//
//		If input parameter 'returnMovedFilesList' is set
//		to 'true', this method will return a populated
//		File Manager Collection documenting all the files
//		actually included in the move operation.
//
//		If input parameter 'returnMovedFilesList' is set
//		to 'false', this method will return an empty and
//		unpopulated instance of FileMgrCollection. This
//		means that the files actually moved from the
//		source directory to the target directory by this
//		method, will NOT be documented.
//
//	copyEmptyTargetDirectory	bool
//
//		If set to 'true' the target directory will be
//		created regardless of whether any files are
//		selected to be moved to that directory. Remember
//		that files are only moved to the target directory
//		if they meet the File Type and File Characteristics
//		criteria for file selection.
//
//	deleteEmptySourceDirectory	bool
//
//		This parameter controls whether empty source
//		directories will be deleted after completion of
//		the move operation.
//
//		If 'deleteEmptySourceDirectory' is set to 'true'
//		and there are no files or subdirectories
//		remaining in the source directory (sourceDMgr)
//		after completion of the move operation, the
//		source directory will be deleted.
//
//		If 'deleteEmptySourceDirectory' is set to
//		'false', the source directory (sourceDMgr) will
//		NOT be deleted after completion of the move
//		operation.
//
//	moveRegularFiles			bool
//
//		If this parameter is set to 'true', Regular Files,
//		which also meet the File Selection Characteristics
//		criteria (fileSelectCriteria), will be included
//		in the move operation.
//
//		Regular Files include text files, image files and
//		executable files.
//
//		For an explanation of Regular and Non-Regular
//		files, see the section on "Definition Of Terms",
//		above.
//
//		If input parameters 'moveRegularFiles',
//		'moveSymLinkFiles' and 'moveOtherNonRegularFiles'
//		are all set to 'false', an error will be returned.
//
//	moveSymLinkFiles			bool
//
//		If this parameter is set to 'true', SymLink Files
//		which also meet the File Selection Characteristics
//		criteria (fileSelectCriteria), will be included
//		in the move operation.
//
//		For an explanation of Regular and Non-Regular
//		files, see the section on "Definition Of Terms",
//		above.
//
//		If input parameters 'moveRegularFiles',
//		'moveSymLinkFiles' and 'moveOtherNonRegularFiles'
//		are all set to 'false', an error will be returned.
//
//	moveOtherNonRegularFiles	bool
//
//		If this parameter is set to 'true', Other
//		Non-Regular Files, which also meet the File
//		Selection Characteristics criteria
//		(fileSelectCriteria), will be included in the
//		move operation.
//
//		Examples of other non-regular file types
//		include device files, named pipes, and sockets.
//		See the Definition Of Terms section above.
//
//		If input parameters 'moveRegularFiles',
//		'moveSymLinkFiles' and 'moveOtherNonRegularFiles'
//		are all set to 'false', an error will be returned.
//
//	fileSelectCriteria			FileSelectionCriteria
//
//		In addition to the File Type Selection Criteria,
//		selected files must conform to the File
//		Characteristics criteria specified by this
//		parameter, 'fileSelectCriteria'.
//
//		File Characteristics Selection criteria allow
//		users to screen files for File Name, File
//		Modification Date and File Mode.
//
//		Files matching these selection criteria, and the
//		File Type filter, will be included in the move
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
//			SourceOriginalSubDirs      uint64
//			SourceDirWasDeleted      bool
//			Errors             error
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
//
//		Upon method completion, be sure to check both
//		Non-Fatal and Fatal errors.
func (dMgr *DirMgr) MoveDirectoryFiles(
	targetDMgr DirMgr,
	returnMovedFilesList bool,
	copyEmptyTargetDirectory bool,
	deleteEmptySourceDirectory bool,
	moveRegularFiles bool,
	moveSymLinkFiles bool,
	moveOtherNonRegularFiles bool,
	fileSelectCriteria FileSelectionCriteria,
	errorPrefix interface{}) (
	dirMoveStats DirectoryMoveStats,
	movedFiles FileMgrCollection,
	nonfatalErrs []error,
	fatalErr error) {

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
		"DirMgr.MoveDirectoryFiles()",
		"")

	if err != nil {

		nonfatalErrs = append(nonfatalErrs, err)

		return dirMoveStats, movedFiles, nonfatalErrs, fatalErr
	}

	dirMoveStats,
		movedFiles,
		nonfatalErrs,
		fatalErr = new(dirMgrHelperAtom).
		moveDirectoryFiles(
			dMgr,
			&targetDMgr,
			returnMovedFilesList,
			copyEmptyTargetDirectory,
			deleteEmptySourceDirectory,
			moveRegularFiles,
			moveSymLinkFiles,
			moveOtherNonRegularFiles,
			fileSelectCriteria,
			"dMgr",
			"targetDMgr",
			ePrefix)

	return dirMoveStats, movedFiles, nonfatalErrs, fatalErr
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
// ----------------------------------------------------------------
//
// # Input Parameters
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
//			SourceOriginalSubDirs      uint64
//			SourceDirWasDeleted      bool
//			Errors             error
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
func (dMgr *DirMgr) MoveDirectoryTree(
	targetDMgr DirMgr,
	errorPrefix interface{}) (
	dirMoveStats DirectoryMoveStats,
	nonfatalErrs []error,
	fatalErr error) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		fatalErr = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgr.MoveDirectoryTree()",
		"")

	if fatalErr != nil {

		return dirMoveStats, nonfatalErrs, fatalErr
	}

	dirMoveStats,
		nonfatalErrs,
		fatalErr = new(dirMgrHelper).moveDirectoryTree(
		dMgr,
		&targetDMgr,
		"dMgr",
		"targetDMgr",
		ePrefix)

	return dirMoveStats, nonfatalErrs, fatalErr
}

// MoveSubDirectoryTree
//
// Moves all subdirectories and their constituent files
// from the source or parent directory 'DirMgr' to a
// target directory tree specified by input parameter
// 'targetDMgr'. If successful, all subdirectories and
// files in the source subdirectory tree will be deleted.
//
// The source or parent directory identified by 'DirMgr'
// and the files within 'DirMgr' will NOT be deleted.
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
// # Return Values
//
//	targetDMgr					DirMgr
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
//			SourceOriginalSubDirs      uint64
//			SourceDirWasDeleted      bool
//			Errors             error
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
func (dMgr *DirMgr) MoveSubDirectoryTree(
	targetDMgr DirMgr,
	errorPrefix interface{}) (
	dirMoveStats DirectoryMoveStats,
	nonfatalErrs []error,
	fatalErr error) {

	if dMgr.lock == nil {
		dMgr.lock = new(sync.Mutex)
	}

	dMgr.lock.Lock()

	defer dMgr.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		fatalErr = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DirMgr."+
			"MoveSubDirectoryTree()",
		"")

	if fatalErr != nil {

		return dirMoveStats, nonfatalErrs, fatalErr
	}

	dirMoveStats,
		nonfatalErrs,
		fatalErr =
		new(dirMgrHelper).moveSubDirectoryTree(
			dMgr,
			&targetDMgr,
			"dMgr",
			"destinationDMgr",
			ePrefix)

	return dirMoveStats, nonfatalErrs, fatalErr
}

// New
//
// Returns a new DirMgr object and populates the the data
// fields.
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
//		(os.PathSeparator Linux='/' or Windows='\').
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

// NewFromFileInfo
//
// Returns a new DirMgr object based on two input
// parameters:
//
//	(1)	A parent directory path string
//
//	(2)	An os.FileInfo object containing the directory
//		name.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	parentDirectoryPath			string
//
//		This string contains the parent directory which
//		will be added to the directory name supplied by
//		input parameter 'fInfo'. These two elements are
//		combined to create the new directory path.
//
//	fInfo						os.FileInfo
//
//		An os.FileInfo object containing the directory
//		name which will be added to the parent directory,
//		specified by input parameter 'parentDirectoryPath',
//		to create the new directory path.
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
//	DirMgr
//
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of Directory Manager (DirMgr).
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
func (dMgr *DirMgr) NewFromFileInfo(
	parentDirectoryPath string,
	fInfo os.FileInfo,
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

	if fInfo == nil {
		return newDirMgr,
			fmt.Errorf("%v\n"+
				"ERROR: Input parameter 'fInfo' is 'nil' and INVALID!\n",
				ePrefix.String())
	}

	var isEmpty bool

	isEmpty,
		err = new(dirMgrHelper).
		setDirMgrWithPathDirectoryName(
			&newDirMgr,
			parentDirectoryPath,
			fInfo.Name(),
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
				fInfo.Name())
	}

	return newDirMgr, err
}

// NewFromDirMgrFileInfo
//
// Configures and returns a new 'DirMgr' instance based
// on two input parameters:
//
//	(1)	A parent directory path extracted from an
//		instance of DirMgr passed as input parameter
//		'parentDirectory'.
//
//	(2)	An os.FileInfo object containing the directory
//		name passed as input parameter
//		'directoryFileInfo'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	parentDirectory				DirMgr
//
//		This DirMgr instance contains the parent
//		directory which will be added to the directory
//		name supplied by input parameter
//		'directoryFileInfo'. These two elements are
//		combined to create the new directory path.
//
//	directoryFileInfo			os.FileInfo
//
//		An os.FileInfo structure containing the directory
//		name which will be added to the parent directory,
//		specified by input parameter 'parentDirectory',
//		to create the new directory path.
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
//	DirMgr
//
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of Directory Manager (DirMgr).
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

	err = new(dirMgrHelperPlanck).isDirMgrValid(
		&parentDirectory,
		"parentDirectory",
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
		err = new(dirMgrHelper).setDirMgrFromKnownPathDirName(
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
//		construct and a return a new, fully populated
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
//	DirMgr
//
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of Directory Manager (DirMgr).
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

	newDirMgr := DirMgr{}

	err = new(dirMgrHelperBoson).copyDirMgrs(
		&newDirMgr,
		&fMgr.dMgr,
		ePrefix.XCpy("fMgr.dMgr->"))

	return newDirMgr, err
}

// NewFromKnownPathDirectoryName
//
// Configures and returns a new 'DirMgr' instance using
// a parent path name and directory name. The parent path
// and directory name are combined to form the full path
// for the new 'DirMgr' instance.
//
// This method will populate all internal field values
// with new values based on input parameters
// 'parentPathName' and 'directoryName'.
//
// This method differs from similar methods in that it
// assumes the input parameters are known values and do
// not require the usual analysis and validation
// screening applied by other methods.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	parentPathName				string
//
//		This string contains the parent directory which
//		will be added to the directory name supplied by
//		input parameter 'directoryName'. These two
//		elements are combined to create the new directory
//		path.
//
//	directoryName				string
//
//		This string contains the directory name which
//		will be added to the parent directory, specified
//		by input parameter 'parentPathName', to create
//		the new directory path.
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
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of Directory Manager (DirMgr).
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

	var isEmpty bool

	isEmpty,
		err = new(dirMgrHelper).setDirMgrFromKnownPathDirName(
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
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the current instance of
//	DirMgr.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathStr						string
//
//		A path string designating a path or directory.
//		To reduce errors, the 'pathStr' should be
//		terminated with an appropriate os.PathSeparator
//		character (Linux='/' or Windows='\').
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

// SetDirMgrFromKnownPathDirName
//
// Configures the internal data field values for the
// current DirMgr instance using a parent path name and
// a directory name. The parent path and directory name
// are combined to form the full path for the current
// 'DirMgr' instance.
//
// This method differs from other "Set" methods in that
// it assumes the input parameters are known values and
// do not require the usual analysis and validation
// screening applied by similar methods.
//
// If more rigorous input parameter validation is required,
// consider using method, DirMgr.SetDirMgr().
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the current instance of
//	DirMgr.
//
//	All previous data field values will be replaced with
//	new values based on input parameters 'parentPathName'
//	and 'directoryName'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	parentPathName				string
//
//		This string contains the parent directory which
//		will be added to the directory name supplied by
//		input parameter 'directoryName'. These two
//		elements are combined to create the new directory
//		path.
//
//	directoryName string
//
//		This string contains the directory name which
//		will be added to the parent directory, specified
//		by input parameter 'parentPathName', to create
//		the new directory path.
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
//		If this returned boolean value is set to 'true',
//		it signals that the 'parentPathName' and
//		'directoryName' input parameters resulted in an
//		empty or zero length directory string.
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
func (dMgr *DirMgr) SetDirMgrFromKnownPathDirName(
	parentPathName string,
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
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the current instance of
//	DirMgr.
//
//	All previous data field values will be replaced with
//	new values based on input parameters 'parentPathName'
//	and 'directoryName'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	parentDirectoryPath			string
//
//		This string contains the parent directory which
//		will be added to the directory name supplied by
//		input parameter 'fInfo'. These two elements are
//		combined to create the new directory path.
//
//	fInfo						os.FileInfo
//
//		An os.FileInfo object containing the directory
//		name which will be added to the parent directory,
//		specified by input parameter 'parentDirectoryPath',
//		to create the new directory path.
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
	fInfo os.FileInfo,
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

	if fInfo == nil {

		return fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'fInfo' is 'nil' and INVALID!\n",
			ePrefix.String())
	}

	isEmpty := true

	isEmpty,
		err = new(dirMgrHelper).
		setDirMgrWithPathDirectoryName(
			dMgr,
			parentDirectoryPath,
			fInfo.Name(),
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
			fInfo.Name())
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
// # BE ADVISED
//
//	This method will NOT modify the current instance of
//	DirMgr.
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
