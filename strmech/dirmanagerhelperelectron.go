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

// getAllSubDirsInDirTree
//
// Identifies and returns all the subdirectories in a
// directory tree.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	(1)	This method will identify, document and return
//		subdirectories located in the entire directory
//		tree identified by input parameter 'dMgr'.
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
//		this method ('dirTreeProfile').
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dMgr						*DirMgr
//
//		A pointer to an instance of DirMgr. This instance
//		specifies the absolute directory path which will
//		be searched to extract and return subdirectories
//		for the entire directory tree.
//
//		If the directory specified by 'dMgr' does not
//		exist on an attached storage drive, an error will
//		be returned.
//
//	subDirectories				*DirMgrCollection
//
//		A pointer to an instance of DirMgrCollection.
//		The DirMgrCollection contains an array of DirMgr
//		objects.
//
//		This method will scan the entire directory tree
//		identified by input parameter 'dMgr'.
//
//		All subdirectories identified in this directory
//		tree will be documented by a new instance of
//		DirMgr which will be added to the
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
//		this method ('dirTreeProfile').
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
//	dirTreeProfile				DirectoryProfile
//
//		If this method completes successfully, this
//		returned instance of DirectoryProfile will be
//		populated with cumulative profile and statistical
//		information on the entire directory tree defined
//		by input parameter 'dMgr'.
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
//		If this method completes successfully, this
//		returned error Type is set equal to 'nil'.
//
//		If an error is encountered during processing,
//		this returned error Type will encapsulate an
//		appropriate error message. This returned error
//		message will incorporate the method chain and
//		text passed by input parameter, 'errPrefDto'. The
//		'errPrefDto' text will be prefixed or attached to
//		the beginning of the error message.
func (dMgrHlprElectron *dirMgrHelperElectron) getAllSubDirsInDirTree(
	dMgr *DirMgr,
	subDirectories *DirMgrCollection,
	dMgrLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	dirTreeProfile DirectoryProfile,
	err error) {

	if dMgrHlprElectron.lock == nil {
		dMgrHlprElectron.lock = new(sync.Mutex)
	}

	dMgrHlprElectron.lock.Lock()

	defer dMgrHlprElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "dirMgrHelperElectron." +
		"getAllSubDirsInDirTree()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return dirTreeProfile, err
	}

	if subDirectories == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'subDirectories' is a 'nil' pointer!\n",
			ePrefix.String())

		return dirTreeProfile, err
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

		return dirTreeProfile, err
	}

	dirTreeProfile.DirAbsolutePath =
		dMgr.absolutePath

	dirTreeProfile.DirExistsOnStorageDrive =
		true

	originalLenOfSubDirsCol := len(subDirectories.dirMgrs)

	var dirProfile DirectoryProfile
	var dMgrHlprPreon = new(dirMgrHelperPreon)

	dirProfile,
		err = dMgrHlprPreon.
		getSubdirectories(
			dMgr,
			subDirectories,
			dMgrLabel,
			ePrefix.XCpy(dMgrLabel))

	if err != nil {

		return dirTreeProfile, err
	}

	if len(subDirectories.dirMgrs) <= originalLenOfSubDirsCol {
		// There are no subdirectories

		return dirTreeProfile, err
	}

	dirTreeProfile.
		AddDirProfileStats(dirProfile)

	var idx = 0

	if originalLenOfSubDirsCol <= 0 {

		idx = 0

	} else {

		idx = originalLenOfSubDirsCol
	}

	var subDirDMgr DirMgr
	var dMgrColHelper = new(dirMgrCollectionHelper)
	var errStatus ArrayColErrorStatus

	for idx >= 0 {

		// This is a peek operation
		subDirDMgr,
			errStatus = dMgrColHelper.
			peekOrPopAtIndex(
				subDirectories,
				idx,
				false, // deleteIndex
				ePrefix)

		if errStatus.ProcessingError != nil {

			if errStatus.IsIndexOutOfBounds == true {

				idx = -1

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

			return dirTreeProfile, err
		}

		dirProfile,
			err = dMgrHlprPreon.
			getSubdirectories(
				&subDirDMgr,
				subDirectories,
				"subDirDMgr",
				ePrefix.XCpy("subDirDMgr"))

		if err != nil {

			return dirTreeProfile, err
		}

		dirTreeProfile.
			AddDirProfileStats(dirProfile)

		idx++
	}

	return dirTreeProfile, errStatus.ProcessingError
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
