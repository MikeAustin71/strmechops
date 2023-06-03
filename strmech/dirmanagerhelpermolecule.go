package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

type dirMgrHelperMolecule struct {
	lock *sync.Mutex
}

// deleteAllSubDirectories
//
// The directory identified by the input parameter 'dMgr'
// is treated as the parent directory.
//
// This method will proceed to delete all directories and
// files which are subsidiary to the parent directory,
// or top level directory, identified by 'dMgr'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// All subdirectories and files which are subordinate to
// the parent or top level directory identified by 'dMgr'
// will be deleted.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dMgr						*DirMgr
//
//		A pointer to an instance of DirMgr. All
//		subdirectories and files subsidiary to the parent
//		or top level directory identified by 'dMgr' will
//		be deleted.
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
func (dMgrHlprMolecule *dirMgrHelperMolecule) deleteAllSubDirectories(
	dMgr *DirMgr,
	dMgrLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	errs []error) {

	if dMgrHlprMolecule.lock == nil {
		dMgrHlprMolecule.lock = new(sync.Mutex)
	}

	dMgrHlprMolecule.lock.Lock()

	defer dMgrHlprMolecule.lock.Unlock()

	funcName := "dirMgrHelper.doesDirectoryExist() "

	errs = make([]error, 0)

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		errs = append(errs, err)

		return errs
	}

	if len(dMgrLabel) == 0 {

		dMgrLabel = "dMgr"

	}

	dirPathDoesExist,
		_,
		err := new(dirMgrHelperAtom).doesDirectoryExist(
		dMgr,
		PreProcPathCode.None(),
		dMgrLabel,
		ePrefix)

	if err != nil {

		errs = append(errs, err)

		return errs
	}

	if !dirPathDoesExist {

		err = fmt.Errorf("%v\n"+
			"ERROR: %v Directory Path DOES NOT EXIST!\n"+
			"%v='%v'\n",
			ePrefix.String(),
			dMgrLabel,
			dMgrLabel,
			dMgr.absolutePath)

		errs = append(errs, err)

		return errs
	}

	var err2, err3 error

	dirMgrPtr, err := os.Open(dMgr.absolutePath)

	if err != nil {

		err2 = fmt.Errorf("%v\n"+
			"Error return by os.Open(dMgr.absolutePath)\n"+
			"dMgr.absolutePath='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			dMgr.absolutePath,
			err.Error())

		errs = append(errs, err2)

		return errs
	}

	var nameFileInfos []os.FileInfo

	err3 = nil

	osPathSeparatorStr := string(os.PathSeparator)

	for err3 != io.EOF {

		nameFileInfos, err3 = dirMgrPtr.Readdir(10000)

		if err3 != nil && err3 != io.EOF {

			_ = dirMgrPtr.Close()

			err2 = fmt.Errorf("%v\n"+
				"Error returned by dirMgrPtr.Readdirnames(10000).\n"+
				"dMgr.absolutePath='%v'\n"+
				"Error='%v'\n",
				ePrefix.String(),
				dMgr.absolutePath,
				err3.Error())

			errs = append(errs, err2)

			return errs
		}

		for _, nameFInfo := range nameFileInfos {

			if nameFInfo.IsDir() {

				err = os.RemoveAll(dMgr.absolutePath + osPathSeparatorStr + nameFInfo.Name())

				if err != nil {

					err2 = fmt.Errorf("%v\n"+
						"Error returned by os.RemoveAll(subDir)\n"+
						"subDir='%v'\n"+
						"Error= \n%v\n",
						ePrefix.String(),
						dMgr.absolutePath+osPathSeparatorStr+nameFInfo.Name(),
						err.Error())

					errs = append(errs, err2)

					continue
				}
			}
		}
	}

	if dirMgrPtr != nil {

		err = dirMgrPtr.Close()

		if err != nil {

			err2 = fmt.Errorf("%v\n"+
				"Error returned by %vPtr.Close().\n"+
				"%v='%v'\n"+
				"Error='%v'\n",
				ePrefix.String(),
				dMgrLabel,
				dMgrLabel,
				dMgr.absolutePath,
				err.Error())

			errs = append(errs, err2)
		}
	}

	return errs
}

// getValidPathStr
//
// Performs validation on a path string. If the string
// contains a filename and file extension, this method
// will declare an error.
//
// Once the path string is determined to be valid, this
// method returns a properly formatted string containing
// the valid path.
//
// If input parameter 'pathStr' is determined to be
// invalid, an error is returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathStr						string
//
//		This strings contains the path to be validated.
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
//	validPathDto				ValidPathStrDto
//
//		If this method completes successfully, an
//		instance of 'ValidPathStrDto' will be returned
//		containing a valid path string constructed from
//		input parameter 'pathStr' plus associated string
//		statistics.
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
func (dMgrHlprMolecule *dirMgrHelperMolecule) getValidPathStr(
	pathStr string,
	pathStrLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	validPathDto ValidPathStrDto,
	err error) {

	if dMgrHlprMolecule.lock == nil {
		dMgrHlprMolecule.lock = new(sync.Mutex)
	}

	dMgrHlprMolecule.lock.Lock()

	defer dMgrHlprMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	validPathDto = new(ValidPathStrDto).New()

	funcName := "dirMgrHelper.getValidPathStr()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return validPathDto, err
	}

	if len(pathStr) == 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input paramter 'pathStr' is invalid!\n"+
			"'pathStr' has a length of zero.\n",
			ePrefix.String())

		return validPathDto, err
	}

	fh := new(FileHelper)

	dMgrHlprElectron := dirMgrHelperElectron{}

	pathSepStr := string(os.PathSeparator)
	dotSeparator := "." + pathSepStr
	doubleDotSeparator := "." + dotSeparator
	doesPathExist := false
	fInfo := FileInfoPlus{}
	var volNameIndex, lSlashIdxs, lDotIdxs,
		strLen, firstCharIdx, lastCharIdx int
	var slashIdxs, dotIdxs []int
	var err2 error
	var volNameStr string

	pathStr,
		strLen,
		err =
		new(dirMgrHelperAtom).
			lowLevelScreenPathStrForInvalidChars(
				pathStr,
				pathStrLabel,
				ePrefix)

	if err != nil {
		goto errorExit
	}

	validPathDto.originalPathStr = pathStr

	if strLen > 2 &&
		// Remove trailing slash
		pathStr[strLen-1] == os.PathSeparator &&
		pathStr[strLen-2] != '.' &&
		pathStr[strLen-2] != os.PathSeparator {
		pathStr = pathStr[0 : strLen-1]
		strLen--
	}

	volNameIndex,
		_,
		volNameStr = fh.GetVolumeNameIndex(pathStr)

	slashIdxs, err2 = fh.GetPathSeparatorIndexesInPathStr(
		pathStr,
		ePrefix)

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"Error returned by fh.GetPathSeparatorIndexesInPathStr(%v).\n"+
			"%v='%v'\n"+
			"Error= \n%v\n",
			funcName,
			pathStrLabel,
			pathStrLabel,
			pathStr,
			err2.Error())

		goto errorExit
	}

	lSlashIdxs = len(slashIdxs)

	firstCharIdx,
		lastCharIdx,
		err2 =
		fh.GetFirstLastNonSeparatorCharIndexInPathStr(
			pathStr,
			ePrefix)

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"Error returned by fh.GetFirstLastNonSeparatorCharIndexInPathStr("+
			"%v).\n"+
			"%v='%v'\n"+
			"Error= \n%v\n",
			funcName,
			pathStrLabel,
			pathStrLabel,
			pathStr,
			err2.Error())

		goto errorExit
	}

	dotIdxs,
		err2 = fh.GetDotSeparatorIndexesInPathStr(
		pathStr,
		ePrefix)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by fh.GetDotSeparatorIndexesInPathStr(%v).\n"+
			"%v='%v'\n"+
			"Error= \n%v\n",
			funcName,
			pathStrLabel,
			pathStrLabel,
			pathStr,
			err2.Error())

		goto errorExit
	}

	lDotIdxs = len(dotIdxs)

	// identify obvious valid path strings

	if pathStr == "." {

		validPathDto.pathStr = dotSeparator
		validPathDto.pathIsValid = PathValidStatus.Valid()
		goto successExit

	} else if pathStr == ".." {

		validPathDto.pathStr = doubleDotSeparator
		validPathDto.pathIsValid = PathValidStatus.Valid()
		goto successExit

	} else if pathStr == dotSeparator {

		validPathDto.pathStr = dotSeparator
		validPathDto.pathIsValid = PathValidStatus.Valid()
		goto successExit

	} else if pathStr == doubleDotSeparator {

		validPathDto.pathStr = doubleDotSeparator
		validPathDto.pathIsValid = PathValidStatus.Valid()
		goto successExit

	}

	if volNameIndex == 0 &&
		strings.ToLower(volNameStr) == strings.ToLower(pathStr) {

		if strings.Contains(strings.ToLower(runtime.GOOS), "windows") {
			pathStr += pathSepStr
		}

		validPathDto.pathStr = pathStr
		validPathDto.pathIsValid = PathValidStatus.Valid()
		goto successExit
	}

	// Check conversion to absolute path
	validPathDto.absPathStr,
		err2 =
		fh.MakeAbsolutePath(pathStr, ePrefix)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by fh.MakeAbsolutePath("+
			"validPathDto.pathStr)\n"+
			"validPathDto.pathStr='%v'\n"+
			"Error= \n%v\n",
			funcName,
			validPathDto.pathStr,
			err2.Error())

		goto errorExit
	}

	if lastCharIdx == -1 &&
		lDotIdxs == 0 &&
		lSlashIdxs == 0 {
		// No characters, no dots and no slashes

		err = fmt.Errorf("%v\n"+
			"Error: %v is INVALID!\n"+
			"%v contains no valid characters in the string!\n"+
			"%v='%v'\n\n",
			ePrefix.String(),
			pathStrLabel,
			pathStrLabel,
			pathStrLabel,
			pathStr)

		goto errorExit

	} else if lastCharIdx == -1 &&
		lDotIdxs > 0 &&
		lSlashIdxs == 0 {
		// No characters, no slashes, but Has Dots
		// Note: good dots have already been processed

		err = fmt.Errorf("%v\n"+
			"Error: %v contains improperly formatted dot characters!\n"+
			"%v='%v'\n",
			ePrefix.String(),
			pathStrLabel,
			pathStrLabel,
			pathStr)

		goto errorExit

	} else if lastCharIdx == -1 &&
		lDotIdxs > 0 &&
		lSlashIdxs > 0 {
		// No characters but Has slashes and
		// has dots.
		validPathDto.pathStr = pathStr
		validPathDto.pathIsValid = PathValidStatus.Valid()

	} else if lastCharIdx == -1 &&
		lDotIdxs == 0 &&
		lSlashIdxs > 0 {

		// No characters, No dots, but Has slashes
		err = fmt.Errorf("%v\n"+
			"Error: '%v' contains improperly formatted path separator characters!\n"+
			"%v='%v'\n\n",
			ePrefix.String(),
			pathStrLabel,
			pathStrLabel,
			pathStr)

		goto errorExit

	} else if lastCharIdx > -1 &&
		lDotIdxs == 0 &&
		lSlashIdxs == 0 {
		// Has characters, but No dots and No slashes
		validPathDto.pathStr = pathStr
		validPathDto.pathIsValid = PathValidStatus.Valid()

	} else if lastCharIdx > -1 &&
		lDotIdxs > 0 &&
		lSlashIdxs == 0 {
		// Has characters, Has Dots, but No slashes
		// Example  someFileName.txt

		if lDotIdxs > 1 {

			// To many dots
			err = fmt.Errorf("%v\n"+
				"Error: %v contains improperly formatted path separator characters!\n"+
				"%v='%v'\n\n",
				ePrefix.String(),
				pathStrLabel,
				pathStrLabel,
				pathStr)

			goto errorExit

		} else {
			// lDotIdx must equal '1'
			if dotIdxs[0] < firstCharIdx {

				// Example .git = directory
				validPathDto.pathStr = pathStr
				validPathDto.pathIsValid = PathValidStatus.Valid()

			} else {

				err = fmt.Errorf("%v\n"+
					"Error: %v contains improperly formatted dot characters!\n"+
					"%v='%v'\n\n",
					ePrefix.String(),
					pathStrLabel,
					pathStrLabel,
					pathStr)

				goto errorExit
			}
		}

	} else if lastCharIdx > -1 &&
		lDotIdxs > 0 &&
		lSlashIdxs > 0 {
		// Has characters, Has slashes, Has dots

		if firstCharIdx < slashIdxs[0] {

			// Example somefile/
			err = fmt.Errorf("%v\n"+
				"Error: %v contains improperly "+
				"formatted characters and path separators!\n"+
				"%v='%v'\n\n",
				ePrefix.String(),
				pathStrLabel,
				pathStrLabel,
				pathStr)

			goto errorExit

		} else if dotIdxs[lDotIdxs-1] >
			lastCharIdx {

			// Example somedir.
			err = fmt.Errorf("%v\n"+
				"Error: %v contains improperly "+
				"formatted dot characters!\n"+
				"%v='%v'\n",
				ePrefix.String(),
				pathStrLabel,
				pathStrLabel,
				pathStr)

			goto errorExit

		} else if dotIdxs[lDotIdxs-1]-slashIdxs[lSlashIdxs-1] == 1 &&
			lastCharIdx > dotIdxs[lDotIdxs-1] {

			// ../dir1/dir2/.git
			validPathDto.pathStr = pathStr
			validPathDto.pathIsValid = PathValidStatus.Valid()

		} else if lastCharIdx > dotIdxs[lDotIdxs-1] &&
			dotIdxs[lDotIdxs-1] > slashIdxs[lSlashIdxs-1] &&
			dotIdxs[lDotIdxs-1]-slashIdxs[lSlashIdxs-1] != 1 {
			// ./dir1/dir2/fileName.ext

			// Trim off trailing file name
			validPathDto.pathStr = pathStr[0:slashIdxs[lSlashIdxs-1]]

			if len(validPathDto.pathStr) == 0 {
				err = fmt.Errorf("%v\n"+
					"Error: %v contains a "+
					"file name!\n"+
					"Attemp to trim trailing file name failed!\n"+
					"%v='%v'\n",
					ePrefix.String(),
					pathStrLabel,
					pathStrLabel,
					pathStr)

				goto errorExit

			} else {
				validPathDto.pathIsValid = PathValidStatus.Valid()
			}

		} else if lastCharIdx > slashIdxs[lSlashIdxs-1] &&
			slashIdxs[lSlashIdxs-1] > dotIdxs[lDotIdxs-1] {

			// ../dir1/dir2/git
			validPathDto.pathStr = pathStr
			validPathDto.pathIsValid = PathValidStatus.Valid()

		} else {

			// unknown error
			err = fmt.Errorf("%v\n"+
				"Error: '%v' contains a "+
				"file name!\n"+
				"%v='%v'\n\n",
				ePrefix.String(),
				pathStrLabel,
				pathStrLabel,
				pathStr)

			goto errorExit
		}

	} else if lastCharIdx > -1 &&
		lDotIdxs == 0 &&
		lSlashIdxs > 0 {
		// Has characters, No Dots, Has slashes

		if slashIdxs[lSlashIdxs-1] > lastCharIdx {

			err = fmt.Errorf("%v\n"+
				"Error: %v contains improperly "+
				"formatted path separators!\n"+
				"%v='%v'\n\n",
				ePrefix.String(),
				pathStrLabel,
				pathStrLabel,
				pathStr)

			goto errorExit

		} else {
			validPathDto.pathStr = pathStr
			validPathDto.pathIsValid = PathValidStatus.Valid()
		}

	} else {

		err = fmt.Errorf("%v\n"+
			"Error: %v is Invalid!\n"+
			"%v='%v'\n\n",
			ePrefix.String(),
			pathStrLabel,
			pathStrLabel,
			pathStr)

		goto errorExit
	}

successExit:

	// Check conversion to absolute path
	validPathDto.absPathStr,
		err2 =
		fh.MakeAbsolutePath(
			validPathDto.pathStr,
			ePrefix)

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"Error returned by fh.MakeAbsolutePath("+
			"validPathDto.pathStr)\n"+
			"validPathDto.pathStr='%v'\n"+
			"Error= \n%v\n",
			funcName,
			validPathDto.pathStr,
			err2.Error())

		goto errorExit
	}

	validPathDto.pathVolumeIndex,
		validPathDto.pathVolumeStrLength,
		validPathDto.pathVolumeName =
		fh.GetVolumeNameIndex(validPathDto.absPathStr)

	doesPathExist,
		fInfo,
		err = dMgrHlprElectron.
		lowLevelDoesDirectoryExist(
			validPathDto.pathStr,
			pathStrLabel,
			ePrefix)

	if err != nil {
		goto errorExit
	}

	if doesPathExist {
		validPathDto.pathDoesExist = PathExistsStatus.Exists()
		validPathDto.pathFInfoPlus = fInfo.CopyOut()

	} else {
		// doesPathExist = false
		validPathDto.pathDoesExist = PathExistsStatus.DoesNotExist()
	}

	doesPathExist,
		fInfo,
		err = dMgrHlprElectron.
		lowLevelDoesDirectoryExist(
			validPathDto.absPathStr,
			pathStrLabel+".absolutePath",
			ePrefix)

	if err != nil {
		goto errorExit
	}

	if doesPathExist {
		validPathDto.absPathDoesExist = PathExistsStatus.Exists()
		validPathDto.absPathFInfoPlus = fInfo.CopyOut()
	} else {
		// doesPathExist = false
		validPathDto.absPathDoesExist = PathExistsStatus.DoesNotExist()
	}

	if validPathDto.pathDoesExist != validPathDto.absPathDoesExist {

		err = fmt.Errorf("%v\n"+
			"ERROR: The path and absolute path show different values for "+
			"existence on disk.\n"+
			"validPathDto.pathDoesExist='%v'\n"+
			"validPathDto.absPathDoesExist='%v'\n",
			ePrefix.String(),
			validPathDto.pathDoesExist.String(),
			validPathDto.absPathDoesExist.String())

		goto errorExit
	}

	if validPathDto.pathDoesExist == PathExistsStatus.Exists() &&
		validPathDto.absPathDoesExist == PathExistsStatus.Exists() {

		if !validPathDto.absPathFInfoPlus.IsDir() {
			err = fmt.Errorf("%v\n"+
				"ERROR: The '%v' absolute path exists but it is classified "+
				"as a File, NOT a directory!\n"+
				"%v base path='%v'\n"+
				"%v absolute path='%v'\n",
				ePrefix.String(),
				pathStrLabel,
				pathStrLabel,
				validPathDto.pathStr,
				pathStrLabel,
				validPathDto.absPathStr)

			goto errorExit
		}

		if validPathDto.absPathFInfoPlus.Mode().IsRegular() {
			err = fmt.Errorf("%v\n"+
				"ERROR: The '%v' absolute path exists but it is classified\n"+
				"as a 'Regular' File, NOT a directory!\n"+
				"%v base path ='%v'\n"+
				"%v absolute path='%v'\n",
				ePrefix.String(),
				pathStrLabel,
				pathStrLabel,
				validPathDto.pathStr,
				pathStrLabel,
				validPathDto.absPathStr)

			goto errorExit

		}

		if !validPathDto.pathFInfoPlus.IsDir() {
			err = fmt.Errorf("%v\n"+
				"ERROR: The '%v' base path exists but it is classified\n"+
				"as a File, NOT a directory!\n"+
				"%v base path='%v'\n"+
				"%v absolute path='%v'\n",
				ePrefix.String(),
				pathStrLabel,
				pathStrLabel,
				validPathDto.pathStr,
				pathStrLabel,
				validPathDto.absPathStr)

			goto errorExit
		}

		if validPathDto.pathFInfoPlus.Mode().IsRegular() {
			err = fmt.Errorf("%v\n"+
				"ERROR: The '%v' base path exists but it is classified\n"+
				"as a 'Regular' File, NOT a directory!\n"+
				"%v base path='%v'\n"+
				"%v absolute path='%v'\n",
				ePrefix.String(),
				pathStrLabel,
				pathStrLabel,
				validPathDto.pathStr,
				pathStrLabel,
				validPathDto.absPathStr)

			goto errorExit

		}

		err = nil
	}

errorExit:
	if err != nil {

		validPathDto.pathStr = pathStr
		validPathDto.absPathStr = ""
		validPathDto.pathStrLength =
			len(pathStr)
		validPathDto.pathIsValid = PathValidStatus.Invalid()
		validPathDto.isInitialized = true
		validPathDto.pathIsValid = PathValidStatus.Invalid()
		validPathDto.err = fmt.Errorf("%v", err.Error())
		return validPathDto, err
	}

	validPathDto.pathStrLength =
		len(validPathDto.pathStr)

	validPathDto.absPathStrLength =
		len(validPathDto.absPathStr)

	validPathDto.pathType = PathFileType.Path()
	validPathDto.pathIsValid = PathValidStatus.Valid()
	validPathDto.isInitialized = true

	err = validPathDto.IsDtoValid(
		ePrefix)

	return validPathDto, err
}

// lowLevelCopyFile
//
// This low level helper method is designed
// to copy files from a source file to a destination file.
//
// No validation or error checking is performed on the input
// parameters.
func (dMgrHlprMolecule *dirMgrHelperMolecule) lowLevelCopyFile(
	srcFile string,
	srcFInfo os.FileInfo,
	dstFile string,
	srcLabel string,
	dstLabel string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if dMgrHlprMolecule.lock == nil {
		dMgrHlprMolecule.lock = new(sync.Mutex)
	}

	dMgrHlprMolecule.lock.Lock()

	defer dMgrHlprMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errPrefDto,
		"dirMgrHelperMolecule."+
			"lowLevelCopyFile()",
		"")

	if err != nil {
		return err
	}

	if len(srcLabel) == 0 {
		srcLabel = "srcFile"
	}

	if len(srcFile) == 0 {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter %v is an empty string!\n",
			ePrefix.String(),
			srcLabel)
	}

	errCode := 0

	errCode,
		_,
		srcFile =
		new(fileHelperElectron).isStringEmptyOrBlank(srcFile)

	if errCode == -1 {

		err = fmt.Errorf("%v\n"+
			"Error: '%v' is an empty string!\n",
			ePrefix.String(),
			srcLabel)

		return err
	}

	if errCode == -2 {

		err = fmt.Errorf("%v\n"+
			"\nError: '%v' consists of blank spaces!\n",
			ePrefix.String(),
			srcLabel)

		return err
	}

	var doesFileExist bool

	srcFile,
		doesFileExist,
		_,
		err = new(fileHelperMolecule).
		doesPathFileExist(
			srcFile,
			PreProcPathCode.AbsolutePath(),
			ePrefix.XCpy("srcFile"),
			srcLabel)

	if err != nil {
		return err
	}

	if !doesFileExist {

		err = fmt.Errorf("%v\n"+
			"Error: %v does NOT exist on disk!\n",
			ePrefix.String(),
			srcLabel)

		return err
	}

	if len(dstLabel) == 0 {
		dstLabel = "dstFile"
	}

	if len(dstFile) == 0 {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter %v is an empty string!\n",
			ePrefix.String(),
			dstLabel)
	}

	if !srcFInfo.Mode().IsRegular() {

		return fmt.Errorf("%v\n"+
			"Error: %v is a Non-Regular File and cannot be copied!\n"+
			"%v='%v'\n",
			ePrefix.String(),
			srcLabel,
			srcLabel,
			srcFile)
	}

	// First, open the source file
	inSrcPtr, err := os.Open(srcFile)

	if err != nil {
		return fmt.Errorf("%v\n"+
			"Error returned from os.Open(srcFile)\n"+
			"%v='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			srcLabel,
			srcFile,
			err.Error())
	}

	// Next, 'Create' the destination file
	// If the destination file previously exists,
	// it will be truncated.
	outDestPtr, err := os.Create(dstFile)

	if err != nil {

		_ = inSrcPtr.Close()

		return fmt.Errorf("%v\n"+
			"Error returned from os.Create(destinationFile)\n"+
			"%v = '%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			dstLabel,
			dstFile,
			err.Error())
	}

	bytesCopied, err2 := io.Copy(outDestPtr, inSrcPtr)

	if err2 != nil {

		_ = inSrcPtr.Close()
		_ = outDestPtr.Close()

		err = fmt.Errorf("%v\n"+
			"Error returned from io.Copy(%v, %v)\n"+
			"%v='%v'\n"+
			"%v='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			dstLabel,
			srcLabel,
			dstLabel,
			dstFile,
			srcLabel,
			srcFile,
			err2.Error())

		return err
	}

	errs := make([]error, 0)

	// flush file buffers inSrcPtr memory
	err = outDestPtr.Sync()

	if err != nil {

		err2 = fmt.Errorf("%v\n"+
			"Error returned from outDestPtr.Sync()\n"+
			"%v='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			dstLabel,
			dstFile,
			err.Error())

		errs = append(errs, err2)
	}

	err = inSrcPtr.Close()

	if err != nil {

		err2 = fmt.Errorf("%v\n"+
			"Error returned from inSrcPtr.Close()\n"+
			"inSrcPtr=source='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			srcFile,
			err.Error())

		errs = append(errs, err2)
	}

	inSrcPtr = nil

	err = outDestPtr.Close()

	if err != nil {

		err2 = fmt.Errorf("%v\n"+
			"Error returned from outDestPtr.Close()\n"+
			"outDestPtr=destination='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			dstFile,
			err.Error())

		errs = append(errs, err2)
	}

	outDestPtr = nil

	if len(errs) > 0 {

		return new(StrMech).ConsolidateErrors(errs)
	}

	var dstFileDoesExist bool
	var dstFileInfo FileInfoPlus

	dstFileDoesExist,
		dstFileInfo,
		err = new(dirMgrHelperElectron).
		lowLevelDoesDirectoryExist(
			dstFile,
			dstLabel,
			ePrefix)

	if err != nil {
		return fmt.Errorf("%v\n"+
			"Error: After Copy IO operation, %v\n"+
			"generated non-path error!\n"+
			"%v='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			dstLabel,
			dstLabel,
			dstFile,
			err.Error())
	}

	if !dstFileDoesExist {

		err = fmt.Errorf("%v\n"+
			"ERROR: After Copy IO operation, the destination file DOES NOT EXIST!\n"+
			"Destination File = '%v' = '%v'\n",
			ePrefix.String(),
			dstLabel,
			dstFile)

		return err
	}

	srcFileSize := srcFInfo.Size()

	if bytesCopied != srcFileSize {

		err = fmt.Errorf("%v\n"+
			"Error: Bytes Copied does NOT equal bytes in source file!\n"+
			"Source File Bytes='%v'   Bytes Coped='%v'\n"+
			"Source File=%v='%v'\n"+
			"Destination File=%v='%v'\n",
			ePrefix.String(),
			srcFileSize,
			bytesCopied,
			srcLabel,
			srcFile,
			dstLabel,
			dstFile)

		return err
	}

	err = nil

	if dstFileInfo.Size() != srcFileSize {

		err = fmt.Errorf("%v\n"+
			"Error: Bytes is source file do NOT equal bytes\n"+
			"in the destination file!\n"+
			"Source File Bytes='%v'   Destination File Bytes='%v'\n"+
			"Source File=%v='%v'\n"+
			"Destination File=%v='%v'\n",
			ePrefix.String(),
			srcFileSize,
			dstFileInfo.Size(),
			srcLabel,
			srcFile,
			dstLabel,
			dstFile)
	}

	return err
}

// lowLevelDeleteDirectoryAll
//
// Helper method designed for use by DirMgr.
//
// This method will delete the designated directory and
// constituent file specified by input parameter 'dMgr',
// as well as all subsidiary directories and files. This
// means that the entire directory tree designated by
// 'dMgr', along with all the contained files, will be
// deleted.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	This low-level method will not perform validation
//		services. It assumes that 'dMgr' specifies a
//		directory path which actually exists on disk.
//
//	(2) This method will delete the directory and
//		constituent files identified by input parameter
//		'dMgr'. In addition, all the child directories
//		and files subordinate to the directory designated
//		by 'dMgr' will likewise be deleted.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dMgr						*DirMgr
//
//		A pointer to an instance of DirMgr. The entire
//		directory tree identified by this parameter will
//		be deleted along with all the resident files.
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
func (dMgrHlprMolecule *dirMgrHelperMolecule) lowLevelDeleteDirectoryAll(
	dMgr *DirMgr,
	dMgrLabel string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if dMgrHlprMolecule.lock == nil {
		dMgrHlprMolecule.lock = new(sync.Mutex)
	}

	dMgrHlprMolecule.lock.Lock()

	defer dMgrHlprMolecule.lock.Unlock()

	funcName := "dirMgrHelper.lowLevelDeleteDirectoryAll() "

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return err
	}

	var err2 error

	if dMgr == nil {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter %v pointer is 'nil' !\n",
			ePrefix.String(),
			dMgrLabel)
	}

	if len(dMgrLabel) == 0 {
		dMgrLabel = "dMgr"
	}

	for i := 0; i < 3; i++ {

		err2 = os.RemoveAll(dMgr.absolutePath)

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Error returned by os.RemoveAll(%v.absolutePath)\n"+
				"%v.absolutePath='%v'\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				dMgrLabel,
				dMgrLabel,
				dMgr.absolutePath,
				err2.Error())

		} else {
			// err2 == nil
			// Deletion was successful
			dMgr.doesAbsolutePathExist = false
			dMgr.doesPathExist = false
			dMgr.actualDirFileInfo = FileInfoPlus{}
			return nil
		}

		time.Sleep(50 * time.Millisecond)
	}

	return err
}

// lowLevelMakeDir
//
// Helper Method used by 'DirMgr'. This method will
// create the directory path including parent directories
// for the path specified by 'dMgr'.
//
// The permissions used to create the new directory are
// automatic, fixed and specified as "drwxrwxrwx".
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dMgr						*DirMgr
//
//		A pointer to an instance of DirMgr. This method
//		will create the directory path including parent
//		directories path specified by this parameter.
//
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
//	dirCreated					bool
//
//		If this returned boolean value is set to 'true',
//		it signals that the directory as been created
//		with the default permissions.
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
func (dMgrHlprMolecule *dirMgrHelperMolecule) lowLevelMakeDir(
	dMgr *DirMgr,
	dMgrLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	dirCreated bool,
	err error) {

	if dMgrHlprMolecule.lock == nil {
		dMgrHlprMolecule.lock = new(sync.Mutex)
	}

	dMgrHlprMolecule.lock.Lock()

	defer dMgrHlprMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	dirCreated = false

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errPrefDto,
		"dirMgrHelperMolecule."+
			"lowLevelMakeDir()",
		"")

	if err != nil {
		return dirCreated, err
	}

	if dMgr == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'dMgr' is a 'nil' pointer!\n",
			ePrefix)

		return dirCreated, err
	}

	if len(dMgrLabel) == 0 {
		dMgrLabel = "dMgr"
	}

	dMgrHlprAtom := dirMgrHelperAtom{}

	dMgrPathDoesExist,
		_,
		err :=
		dMgrHlprAtom.doesDirectoryExist(
			dMgr,
			PreProcPathCode.None(),
			dMgrLabel,
			ePrefix)

	if err != nil {
		return dirCreated, err
	}

	if dMgrPathDoesExist {
		// The directory exists
		// Nothing to do.
		return dirCreated, err
	}

	var fPermCfg FilePermissionConfig

	fPermCfg, err =
		new(FilePermissionConfig).New(
			"drwxrwxrwx",
			ePrefix)

	if err != nil {

		return dirCreated, err
	}

	var modePerm os.FileMode

	modePerm,
		err = fPermCfg.GetCompositePermissionMode(
		ePrefix.XCpy(
			"modePerm<-fPermCfg"))

	if err != nil {

		return dirCreated, err
	}

	var err2 error

	err2 = os.MkdirAll(dMgr.absolutePath, modePerm)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by os.MkdirAll(dMgr.absolutePath, modePerm).\n"+
			"dMgr.absolutePath='%v'\n"+
			"modePerm=\"drwxrwxrwx\"\n"+
			"Error='%v'\n",
			ePrefix.String(),
			dMgr.absolutePath,
			err2.Error())

		return dirCreated, err
	}

	dMgrPathDoesExist,
		_,
		err2 =
		dMgrHlprAtom.doesDirectoryExist(
			dMgr,
			PreProcPathCode.None(),
			dMgrLabel,
			ePrefix.XCpy(
				"dMgr"))

	if err2 != nil {
		err = fmt.Errorf("Error: After attempted directory creation, "+
			"a non-path error was generated!\n"+
			"%v.absolutePath='%v'\n"+
			"Error='%v'\n",
			dMgrLabel,
			dMgr.absolutePath,
			err2.Error())
		return dirCreated, err
	}

	if !dMgrPathDoesExist {
		err = fmt.Errorf("Error: After attempted directory creation,\n"+
			"the directory DOES NOT EXIST!\n"+
			"%v=%v\n", dMgrLabel, dMgr.absolutePath)

		return dirCreated, err
	}

	dirCreated = true
	err = nil

	return dirCreated, err
}

// lowLevelMakeDirWithPermission
//
// Helper Method used by 'DirMgr'. This method will
// create the directory path including parent directories
// for the path specified by 'dMgr'. The permissions used
// to create the directory path are specified by input
// parameter 'fPermCfg'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dMgr						*DirMgr
//
//		A pointer to an instance of DirMgr. This method
//		will create the directory path including parent
//		directories path specified by this parameter.
//
//	fPermCfg					FilePermissionConfig
//
//		An instance of FilePermissionConfig containing
//		the permission specifications for the new
//		directory to be created from input paramter,
//		'dMgr'.
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
//	dirCreated					bool
//
//		If this returned boolean value is set to 'true',
//		it signals that the directory as been created
//		with the default permissions.
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
func (dMgrHlprMolecule *dirMgrHelperMolecule) lowLevelMakeDirWithPermission(
	dMgr *DirMgr,
	fPermCfg FilePermissionConfig,
	dMgrLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	dirCreated bool,
	err error) {

	if dMgrHlprMolecule.lock == nil {
		dMgrHlprMolecule.lock = new(sync.Mutex)
	}

	dMgrHlprMolecule.lock.Lock()

	defer dMgrHlprMolecule.lock.Unlock()

	dirCreated = false

	funcName := "dirMgrHelper.lowLevelMakeDirWithPermission() "

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return false, err
	}

	dMgrHlprAtom := dirMgrHelperAtom{}

	dMgrPathDoesExist,
		_,
		err :=
		dMgrHlprAtom.doesDirectoryExist(
			dMgr,
			PreProcPathCode.None(),
			dMgrLabel,
			ePrefix)

	if err != nil {
		return dirCreated, err
	}

	if dMgrPathDoesExist {
		// The directory exists
		// Nothing to do.
		return dirCreated, err
	}

	err2 := fPermCfg.IsValidInstanceError(
		ePrefix.XCpy(
			"fPermCfg"))

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"Input Parameter 'fPermCfg' is INVALID!\n"+
			"Error returned by fPermCfg.IsValidInstanceError().\n"+
			"Error= \n%v\n",
			funcName,
			err2.Error())

		return dirCreated, err
	}

	modePerm, err2 := fPermCfg.GetCompositePermissionMode(
		ePrefix.XCpy(
			"fPermCfg"))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by fPermCfg.GetCompositePermissionMode().\n"+
			"Error= \n%v\n",
			funcName,
			err2.Error())

		return dirCreated, err
	}

	err2 = os.MkdirAll(dMgr.absolutePath, modePerm)

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"Error returned by os.MkdirAll(%v.absolutePath, modePerm).\n"+
			"%v.absolutePath='%v'\n"+
			"modePerm=\"drwxrwxrwx\"\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			dMgrLabel,
			dMgrLabel,
			dMgr.absolutePath,
			err2.Error())

		return dirCreated, err
	}

	dMgrPathDoesExist,
		_,
		err2 =
		dMgrHlprAtom.doesDirectoryExist(
			dMgr,
			PreProcPathCode.None(),
			dMgrLabel,
			ePrefix)

	if err2 != nil {
		err = fmt.Errorf("%v"+
			"Error: After attempted directory creation, "+
			"a non-path error was generated!\n"+
			"%v.absolutePath='%v'\n"+
			"Error= \n%v\n",
			funcName,
			dMgrLabel,
			dMgr.absolutePath,
			err2.Error())

		return dirCreated, err
	}

	if !dMgrPathDoesExist {
		err = fmt.Errorf("Error: After attempted directory creation,\n"+
			"the directory DOES NOT EXIST!\n"+
			"%v=%v\n", dMgrLabel, dMgr.absolutePath)
		return dirCreated, err
	}

	dirCreated = true

	return dirCreated, err
}
