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

	if len(pathStrLabel) == 0 {
		pathStrLabel = "pathStr"
	}

	strLen = len(pathStr)

	if strLen == 0 {
		err = fmt.Errorf("%v\n"+
			"ERROR: %v is an empty string!\n",
			ePrefix.String(),
			pathStrLabel)

		return pathFileNameExt, strLen, err
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

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"dirMgrHelperElectron."+
			"lowLevelDoesDirectoryExist()",
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
			"Error: Input paramter %v is an empty string!\n",
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
				"Non-Path error returned by os.Stat(%v)\n"+
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
				NewFromPathFileInfo(dirPath, info)

			if err2 != nil {

				err = fmt.Errorf("%v\n"+
					"Error returned by FileInfoPlus{}.NewFromPathFileInfo(dirPath, info)\n"+
					"Error= \n%v\n",
					ePrefix.String(),
					err2.Error())

				fInfoPlus = FileInfoPlus{}
			}

			return dirPathDoesExist, fInfoPlus, err
		}

		time.Sleep(30 * time.Millisecond)
	}

	return dirPathDoesExist, fInfoPlus, err
}
