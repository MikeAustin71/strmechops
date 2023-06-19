package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"os"
	"sync"
)

type dirMgrHelperPreon struct {
	lock *sync.Mutex
}

// getFilesInDir
//
// This method scans a designated directory and returns
// information on selected files as a File Manager
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
//		files, see the section on Definition of Terms,
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
func (dMgrHlprPreon *dirMgrHelperPreon) getFilesInDir(
	targetDMgr *DirMgr,
	getRegularFiles bool,
	getSymLinksFiles bool,
	getOtherNonRegularFiles bool,
	fileSelectCriteria FileSelectionCriteria,
	filesInDir *FileMgrCollection,
	targetDMgrLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	dirProfile DirectoryProfile,
	err error) {

	if dMgrHlprPreon.lock == nil {
		dMgrHlprPreon.lock = new(sync.Mutex)
	}

	dMgrHlprPreon.lock.Lock()

	defer dMgrHlprPreon.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "dirMgrHelperPreon." +
		"getSubdirectories()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		return dirProfile, err
	}

}

// getSubdirectories
//
// This method receives an instance of DirMgr ('dMgr')
// and proceeds to identify all the subdirectories
// located within the directory path specified by this
// DirMgr instance.
//
// Each subdirectory located in the 'dMgr' parent
// directory will be recorded as a separate DirMgr object
// and added to the Directory Manager collection passed
// as input parameter 'subDirectories'.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	(1)	This method will only identify, document and
//		return subdirectories located in the top level or
//		parent directory identified by input parameter
//		'dMgr'.
//
//	(2) The top level or parent directory specified by
//		input parameter 'dMgr' will NOT be included in
//		the Directory Manager collection returned by this
//		method ('subDirectories').
//
//	(3)	Directory entries for the current directory (".")
//		and the parent directory ("..") will be skipped.
//		These directory entries will not be added or
//		included in the subdirectories collection
//		('subDirectories'). Likewise, these two directory
//		entries will NOT be included in the subdirectory
//		profile and statistical information returned by
//		this method ('dirProfile').
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dMgr						*DirMgr
//
//		A pointer to an instance of DirMgr. This instance
//		specifies the absolute directory path which will be
//		searched to extract and return information on
//		subdirectories residing within that directory path.
//
//		If the directory specified by 'dMgr' does not
//		exist, an error will be returned.
//
//	subDirectories				*DirMgrCollection
//
//		A pointer to an instance of DirMgrCollection.
//		The DirMgrCollection contains an array of DirMgr
//		objects.
//
//		This method will scan the top level or parent
//		directory identified by input parameter 'dMgr'.
//		All subdirectories identified in this parent
//		directory will be documented by a new instance
//		of DirMgr which will be added to the
//		'subDirectories' collection.
//
//			type DirMgrCollection struct {
//				dirMgrs []DirMgr
//			}
//
//		Directory entries for the current directory (".")
//		and the parent directory ("..") will be skipped.
//		These directory entries will not be added or
//		included in the subdirectories collection
//		('subDirectories'). Likewise, these two directory
//		entries will NOT be included in the subdirectory
//		profile and statistical information returned by
//		this method ('dirProfile').
//
//	dMgrLabel					string
//
//		The name or label associated with input parameter
//		'dMgr', which will be used in error messages
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
//	dirProfile					DirectoryProfile
//
//		If this method completes successfully, this
//		returned instance of DirectoryProfile will be
//		populated with profile and statistical
//		information on the parent directory identified by
//		input parameter 'dMgr'.
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
func (dMgrHlprPreon *dirMgrHelperPreon) getSubdirectories(
	dMgr *DirMgr,
	subDirectories *DirMgrCollection,
	dMgrLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	dirProfile DirectoryProfile,
	err error) {

	if dMgrHlprPreon.lock == nil {
		dMgrHlprPreon.lock = new(sync.Mutex)
	}

	dMgrHlprPreon.lock.Lock()

	defer dMgrHlprPreon.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "dirMgrHelperPreon." +
		"getSubdirectories()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		return dirProfile, err
	}

	if subDirectories == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'subDirectories' is a 'nil' pointer!\n",
			ePrefix.String())

		return dirProfile, err
	}

	if len(dMgrLabel) == 0 {

		dMgrLabel = "dMgr"
	}

	_,
		_,
		err = new(dirMgrHelperPreon).
		validateDirMgr(
			dMgr,
			true, // Path MUST exist on disk
			dMgrLabel,
			ePrefix.XCpy(
				dMgrLabel))

	if err != nil {

		return dirProfile, err
	}

	dirProfile.DirAbsolutePath = dMgr.absolutePath
	dirProfile.DirExistsOnStorageDrive = true

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

		return dirProfile, err
	}

	if len(nameDirEntries) == 0 {

		return dirProfile, err
	}

	osPathSepStr := string(os.PathSeparator)

	var fInfo os.FileInfo

	for _, dirEntry := range nameDirEntries {

		fInfo,
			err2 = dirEntry.Info()

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Conversion of Direct Entry to os.FileInfo Failed!\n"+
				"Error returned by dirEntry.Info().\n"+
				"%v.absolutePath='%v'\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				dMgrLabel,
				dMgr.absolutePath,
				err2.Error())

			return dirProfile, err
		}

		if fInfo.IsDir() {

			if fInfo.Name() == "." ||
				fInfo.Name() == ".." {

				// Skip the current directory and
				// the parent directory entries.
				continue
			}

			dirProfile.DirSubDirectories++
			dirProfile.DirSubDirectoriesBytes +=
				uint64(fInfo.Size())

			err2 = subDirectories.
				AddDirMgrByKnownPathDirName(
					dMgr.absolutePath,
					fInfo.Name(),
					ePrefix.XCpy(
						"subDirectories<-dMgr"))

			if err2 != nil {

				err = fmt.Errorf("%v\n"+
					"Error occurred while adding subdirectory\n"+
					"to 'subDirectories' collection.\n"+
					"%v Absolute Path= '%v'\n"+
					"Subdirectory Name= '%v'\n"+
					"Subdirectory Path= '%v'\n"+
					"Error=\n%v\n",
					funcName,
					dMgrLabel,
					dMgr.absolutePath,
					fInfo.Name(),
					dMgr.absolutePath+
						osPathSepStr+
						fInfo.Name(),
					err2.Error())

				return dirProfile, err
			}

		} else {
			// This must be a file
			if fInfo.Mode().IsRegular() {

				dirProfile.DirTotalFiles++
				dirProfile.DirTotalFileBytes +=
					uint64(fInfo.Size())

				dirProfile.DirRegularFiles++
				dirProfile.DirRegularFileBytes +=
					uint64(fInfo.Size())

			} else if fInfo.Mode()&os.ModeSymlink != 0 {

				dirProfile.DirTotalFiles++
				dirProfile.DirTotalFileBytes +=
					uint64(fInfo.Size())

				dirProfile.DirSymLinkFiles++
				dirProfile.DirSymLinkFileBytes +=
					uint64(fInfo.Size())

			} else {

				dirProfile.DirTotalFiles++
				dirProfile.DirTotalFileBytes +=
					uint64(fInfo.Size())

				dirProfile.DirNonRegularFiles++
				dirProfile.DirNonRegularFileBytes +=
					uint64(fInfo.Size())

			}
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
			"dirProfile.DirSubDirectories +\n"+
			"dirProfile.DirRegularFiles +\n"+
			"dirProfile.DirSymLinkFiles +\n"+
			"dirProfile.DirNonRegularFiles = %v\n"+
			"The Total Number of files Processed = %v\n",
			ePrefix.String(),
			checkTotalFiles,
			dirProfile.DirTotalFiles)

		dirProfile.ComputeError = fmt.Errorf("%v",
			err.Error())
	}

	return dirProfile, err
}

// validateDirMgr
//
// This method performs a comprehensive analysis to
// determine if an instance of DirMgr is valid.
//
// Users have the option to configure the validity test
// to require that the Directory Manager directory path
// actually exists on disk.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dMgr						*DirMgr
//
//		This instance of DirMgr will be analyzed to
//		determine if all data values are valid.
//
//	pathMustExist				bool
//
//		If this parameter is set to 'true', the directory
//		path contained in 'dMgr' must exist on disk as
//		requirement for validation.
//
//	dMgrLabel					string
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
//	finalDirMgrLabel			string
//
//		The final formatted version of input parameter
//		'dMgrLabel'.
//
//		The name or label associated with input parameter
//		'dMgr' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "dMgr" will be
//		automatically applied.
//
//	pathDoesExist               bool
//
//		If this return parameter is set to 'true' it
//		signals that the directory path contained in the
//		Directory Manager instance 'dMgr' actually exists
//		on disk.
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
func (dMgrHlprPreon *dirMgrHelperPreon) validateDirMgr(
	dMgr *DirMgr,
	pathMustExist bool,
	dMgrLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	finalDirMgrLabel string,
	pathDoesExist bool,
	err error) {

	if dMgrHlprPreon.lock == nil {
		dMgrHlprPreon.lock = new(sync.Mutex)
	}

	dMgrHlprPreon.lock.Lock()

	defer dMgrHlprPreon.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	pathDoesExist = false

	funcName := "dirMgrHelperPreon." +
		"validateDirMgr()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		return finalDirMgrLabel, pathDoesExist, err
	}

	if len(dMgrLabel) == 0 {

		finalDirMgrLabel = "dMgr"

	} else {

		finalDirMgrLabel = dMgrLabel
	}

	if dMgr == nil {

		err = fmt.Errorf("%v \n"+
			"ERROR: Input paramter '%v' is a nil pointer!\n",
			ePrefix.String(),
			finalDirMgrLabel)

		return finalDirMgrLabel, pathDoesExist, err
	}

	var err2 error

	err2 = new(dirMgrHelperPlanck).
		isDirMgrValid(
			dMgr,
			dMgrLabel,
			ePrefix.XCpy(finalDirMgrLabel))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input paramter '%v' is INVALID!\n"+
			"Error= \n%v\n",
			funcName,
			finalDirMgrLabel,
			err2.Error())

		return finalDirMgrLabel, pathDoesExist, err
	}

	pathDoesExist,
		_,
		err =
		new(dirMgrHelperAtom).
			doesDirectoryExist(
				dMgr,
				PreProcPathCode.AbsolutePath(),
				finalDirMgrLabel,
				ePrefix.XCpy(finalDirMgrLabel))

	if err != nil {

		return finalDirMgrLabel, pathDoesExist, err
	}

	if !pathMustExist {

		return finalDirMgrLabel, pathDoesExist, err
	}

	if !pathDoesExist {

		err = fmt.Errorf("%v\n"+
			"Error: The current DirMgr path DOES NOT EXIST!\n"+
			"%v.absolutePath='%v'\n",
			ePrefix.String(),
			finalDirMgrLabel,
			dMgr.absolutePath)

	}

	return finalDirMgrLabel, pathDoesExist, err
}
