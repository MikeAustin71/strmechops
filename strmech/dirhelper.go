package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type dirHelper struct {
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
func (dHlpr *dirHelper) GetDirectoryProfile(
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

	funcName := "dirHelper." +
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
