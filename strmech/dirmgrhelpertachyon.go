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

		dirProfile.Errors = append(
			dirProfile.Errors,
			fmt.Errorf("%v",
				err.Error()))
	}

	return directoryPathDoesExist, dirProfile, err
}
