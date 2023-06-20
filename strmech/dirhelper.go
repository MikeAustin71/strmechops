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
// This method returns an instance of DirectoryProfile which
// includes file breakdowns and statistics on the directory
// path specified by input parameter 'directoryPath'.
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
//		'directoryPath' may be formatted as a relative
//		path or an absolute path.
//
//		If 'directoryPath' is invalid, an error will be
//		returned.
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
//		parameter 'directoryPath' actually exists on an
//		attached storage drive.
//
//	dirProfile					DirectoryProfile
//
//		If this method completes successfully, this
//		returned instance of DirectoryProfile will be
//		populated with profile and statistical
//		information on the directory identified by input
//		parameter 'directoryPath'.
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
func (dHlpr *DirHelper) GetDirectoryProfile(
	directoryPath string,
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
			"dMgr",
			ePrefix.XCpy("directoryPath->dMgr"))

	return directoryPathDoesExist, dirProfile, err
}

// GetSubdirectoriesDirTree
//
// This method scans and identifies all the
// subdirectories in the entire directory tree defined by
// input parameter 'directoryPath'. Subdirectories
// located in this directory tree are returned to the
// user by means of a Directory Manager Collection
// (DirMgrCollection) passed as input parameter
// 'subDirectories'.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	(1)	All subdirectories residing in all levels of the
//		directory tree defined by input parameter
//		'directoryPath' will be added and returned in the
//		Directory Manager Collection passed as input
//		parameter 'subDirectories'.
//
//	(2)	While scanning for subdirectories, Directory
//		entries for the current directory (".") and the
//		parent directory ("..") will be skipped and will
//		NOT be added to the 'subDirectories' Directory
//		Manager Collection.
//
//	(3)	For a collection of subdirectories residing in
//		the top level or parent directory specified by
//		a directory path, see method:
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
//		input parameter 'subDirectories'.
//
//		If this directory path does not exist on an
//		attached storage drive, an error will be
//		returned.
//
//	subDirectories				*DirMgrCollection
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
//	numOfSubdirectories			uint64
//
//		If this method completes successfully without
//		error, this parameter will return the number
//		of subdirectories located in the directory tree
//		defined by input parameter 'directoryPath'. This
//		uint64 value also represents the number of
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
func (dHlpr *DirHelper) GetSubdirectoriesDirTree(
	directoryPath string,
	subDirectories *DirMgrCollection,
	errorPrefix interface{}) (
	numOfSubdirectories uint64,
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

	if subDirectories == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'subDirectories' is invalid!\n"+
			"'subDirectories' is a nil pointer.\n",
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

	var dirProfile DirectoryProfile

	dirProfile,
		err = new(dirMgrHelperElectron).
		getAllSubDirsInDirTree(
			&dMgr,
			subDirectories,
			"dMgr",
			ePrefix)

	if err == nil {
		numOfSubdirectories =
			dirProfile.DirSubDirectories
	}

	return numOfSubdirectories, err
}

// GetSubdirectoriesParentDir
//
// This method scans and identifies all the
// subdirectories in the top level or parent directory
// specified by input parameter 'directoryPath'. These
// subdirectories are returned to the user by means of
// a Directory Manager Collection (DirMgrCollection)
// passed as input parameter 'subDirectories'.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	(1)	Only the subdirectories residing in the single
//		parent directory defined by input parameter
//		'directoryPath' will be added and returned in the
//		Directory Manager Collection passed as input
//		parameter 'subDirectories'.
//
//		The search for subdirectories will NOT extend to
//		the directory tree below the 'directoryPath'
//		parent directory.
//
//	(2)	While scanning for subdirectories, Directory
//		entries for the current directory (".") and the
//		parent directory ("..") will be skipped and will
//		NOT be added to the 'subDirectories' Directory
//		Manager Collection.
//
//	(3)	For a collection of all subdirectories in the
//		directory tree specified by a directory path, see
//		method:
//
//			DirHelper.GetSubdirectoriesDirTree
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	directoryPath				string
//
//		This string defines a directory path. This path
//		will be treated as a top level or parent
//		directory. Subdirectories residing in this single
//		parent directory will be added and returned to
//		the user by means of the Directory Manager
//		Collection passed as input parameter
//		'subDirectories'.
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
//	subDirectories				*DirMgrCollection
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
//		Directory Manager Collection ('subDirectories').
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
//	numOfSubdirectories			uint64
//
//		If this method completes successfully, without
//		error, this parameter will return the number
//		of subdirectories located in the parent directory
//		defined by the input parameter 'directoryPath'.
//		This uint64 value also represents the number of
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
func (dHlpr *DirHelper) GetSubdirectoriesParentDir(
	directoryPath string,
	subDirectories *DirMgrCollection,
	errorPrefix interface{}) (
	numOfSubdirectories uint64,
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

	if subDirectories == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'subDirectories' is invalid!\n"+
			"'subDirectories' is a nil pointer.\n",
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

	var dirProfile DirectoryProfile

	dirProfile,
		err = new(dirMgrHelperPreon).
		getSubdirectories(
			&dMgr,
			subDirectories,
			"dMgr",
			ePrefix)

	if err == nil {
		numOfSubdirectories =
			dirProfile.DirSubDirectories
	}

	return numOfSubdirectories, err
}
