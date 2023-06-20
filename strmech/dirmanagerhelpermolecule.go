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

// lowLevelGetFileInfosFromDir
//
// Receives an instance of DirMgr ('dMgr') and proceeds
// to extract os.FileInfo data describing selected
// subdirectories and files contained in that DirMgr's
// absolute directory path. The results are returned to
// the user as an array of os.FileInfo objects.
//
// This is a low level method. It is therefore assumed
// that input parameter 'dMgr' (DirMgr) has been
// validated by the calling function. Calling functions
// are therefore responsible for validating 'dMgr' and
// ensuring that the directory path currently exists on
// an attached storage drive.
//
// This method returns os.FileInfo objects describing
// selected subdirectories and files. The os.FileInfo
// interface is defined as follows:
//
//	 	type FileInfo interface {
//			 Name() string
//				Base name of the file
//
//			 Size() int64
//			 	Length in bytes for regular files;
//			 	system-dependent for others
//
//			 Mode() FileMode
//			 	File mode bits
//
//			 ModTime() time.Time
//			 	Modification time
//
//			 IsDir() bool
//			 	Abbreviation for Mode().IsDir()
//
//			 Sys() interface{}
//			 	Underlying data source (can return nil)
//	 	}
//
// Upon completion, this method returns an array of
// FileInfoPlus objects containing os.FileInfo
// information on selected subdirectories and files
// residing in the directory path specified by input
// parameter 'dMgr'. Type FileInfoPlus implements the
// os.FileInfo interface, but provides additional
// information on the subject subdirectory or file.
//
// To qualify for selection and inclusion in the returned
// array of os.FileInfo objects, items residing in the
// 'dMgr' target directory are divided into two classes,
// subdirectories and files. Subdirectories are standard
// directory entries. Files are defined as all artifacts
// residing in the target directory which are not
// subdirectories.
//
// To qualify as a selected subdirectory, the
// subdirectory must satisfy two filters. First, input
// parameter 'getSubdirectoryFileInfos' must be set to
// 'true'. Second, the subdirectory must satisfy the
// Directory Characteristics Selection Criteria specified
// by input parameter, 'subdirectorySelectCharacteristics'.
// If both of these filter requirements are satisfied, the
// subdirectory will be added to, and returned by, the
// os.FileInfo array, 'fileInfos'. Be advised that users
// can control behavior for current directories (".") and
// parent directories ("..") with input parameters
// 'includeSubDirCurrenDirOneDot' and
// 'includeSubDirParentDirTwoDots'.
//
// To qualify as a selected file, the file entry must
// also comply with two filters: File Type and File
// Characteristics. Remember that files are defined as
// all artifacts residing in the target directory which
// are not subdirectories.
//
// To be eligible for file selection, the file entry must
// first comply with the specified File Type criteria. In
// terms of File Type, files are classified as
// regular files, SymLink files or other non-regular files.
//
// Screening criteria for File Type is controlled by the
// following three input parameters:
//
//	getRegularFileInfos - bool
//	getSymLinksFileInfos - bool
//	getOtherNonRegularFiles - bool
//
// File Types eligible for selection include Regular
// Files such as text files, image files and executable
// files, SymLink files and other Non-Regular Files such
// as device files, named pipes and sockets.
//
// In addition to File Type, selected files,
// must also comply with the File Characteristics
// Selection Criteria specified by input parameter,
// 'fileSelectCharacteristics'.
//
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
//	(1)	This is a low level method. It is therefore
//		assumed that input parameter 'dMgr' (DirMgr) was
//		previously validated by the calling function.
//		Calling	functions are responsible for validating
//		'dMgr' and ensuring that this absolute directory
//		path currently exists on an attached storage
//		drive.
//
//	(2)	To qualify as a selected subdirectory, the
//		subdirectory must satisfy two filters. First,
//		input parameter 'getSubdirectoryFileInfos' must
//		be set to 'true'. Second, the subdirectory must
//		satisfy the Directory Characteristics Selection
//		Criteria specified by input parameter,
//		'subdirectorySelectCharacteristics'. If both of
//		these filter requirements are satisfied, the
//		subdirectory will be added to, and returned by,
//		the os.FileInfo array, 'fileInfos'. Be advised
//		that users can control behavior for current
//		directories (".") and parent directories ("..")
//		with input parameters 'includeSubDirCurrenDirOneDot'
//		and 'includeSubDirParentDirTwoDots'.
//
//	(3) All artifacts located in the target directory
//		defined by input parameter 'dMgr', which are not
//		subdirectories, are treated as files. To qualify
//		as a selected file, the file entry must comply
//		with two filters: File Type and File
//		Characteristics.
//
//		File Type Selection Criteria is controlled and
//		specified by three input parameters:
//
//			getRegularFileInfos - bool
//			getSymLinksFileInfos - bool
//			getOtherNonRegularFiles - bool
//
//		File Characteristics Selection Criteria is
//		specified by input parameter
//		'fileSelectCharacteristics'.
//
//		If both of these filter requirements are
//		satisfied, the subject file will be added to, and
//		returned by, the os.FileInfo array, 'fileInfos'.
//
//	(4) If the directory identified by input parameter
//		'dMgr' contains NO Subdirectories or Files
//		meeting (1) the File Type Selection Criteria and
//		(2) the File Characteristics Selection Criteria,
//		this method will exit, no files will be included
//		in the returned os.FileInfo array, and no error
//		will be returned.
//
//	(5) If the directory identified by input parameter
//		'dMgr' contains NO Files whatsoever (0 Files),
//		this method will exit, no files will be included
//		in the returned os.FileInfo array, and no error
//		will be returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	dMgr								*DirMgr
//
//		A pointer to an instance of DirMgr. This instance
//		specifies the absolute directory path which will
//		be searched to extract and return os.FileInfo
//		information on all files and directories
//		contained therein.
//
//		Calling functions are responsible for validating
//		'dMgr' and ensuring that the absolute directory
//		path currently exists on an attached storage
//		drive. No validation is performed by this method.
//
//		If the directory specified by 'dMgr' does not
//		exist on an attached storage drive, errors
//		will be returned.
//
//	getSubdirectoryFileInfos			bool
//
//		If this parameter is set to 'true', directory
//		entries which also meet the File Selection
//		Characteristics criteria (fileSelectCharacteristics),
//		will be included in the os.FileInfo information
//		('fileInfos') returned by this method.
//
//		If input parameters 'getSubdirectoryFileInfos',
//		'getRegularFileInfos', 'getSymLinksFileInfos' and
//		'getOtherNonRegularFileInfos' are all set to
//		'false', they are classified as conflicted and an
//		error will be returned.
//
//	includeSubDirCurrenDirOneDot		bool
//
//		This parameter is only used, if input parameter
//		'getSubdirectoryFileInfos' is set to 'true'.
//
//		When this parameter, 'includeSubDirCurrenDirOneDot',
//		is set to 'true' and input parameter
//		getSubdirectoryFileInfos' is set to 'true', the current
//		directory, designated by a single dot ('.'), will be
//		included in the returned array of os.FileInfo objects.
//
//	includeSubDirParentDirTwoDots 		bool
//
//		This parameter is only used, if input parameter
//		'getSubdirectoryFileInfos' is set to 'true'.
//
//		When this parameter, 'includeSubDirParentDirTwoDots',
//		is set to 'true' and input parameter
//		'getSubdirectoryFileInfos' is set to 'true', the parent
//		directory, designated by a two dots ('..'), will be
//		included in the returned array of os.FileInfo objects.
//
//	getRegularFileInfos					bool
//
//		If this parameter is set to 'true', regular files,
//		which also meet the File Selection
//		Characteristics criteria (fileSelectCharacteristics),
//		will be included in the os.FileInfo information
//		('fileInfos') returned by this method.
//
//		Regular Files include text files, image files and
//		executable files.
//
//		For an explanation of Regular and Non-Regular
//		files, see the section on "Definition Of Terms",
//		above.
//
//		If input parameters 'getSubdirectoryFileInfos',
//		'getRegularFileInfos', 'getSymLinksFileInfos' and
//		'getOtherNonRegularFileInfos' are all set to
//		'false', they are classified as conflicted and an
//		error will be returned.
//
//	getSymLinksFileInfos				bool
//
//		If this parameter is set to 'true', SymLink files,
//		which also meet the File Selection
//		Characteristics criteria (fileSelectCharacteristics),
//		will be included in the os.FileInfo information
//		('fileInfos') returned by this method.
//
//		For an explanation of Regular and Non-Regular
//		files, see the section on "Definition Of Terms",
//		above.
//
//		If input parameters 'getSubdirectoryFileInfos',
//		'getRegularFileInfos', 'getSymLinksFileInfos' and
//		'getOtherNonRegularFileInfos' are all set to
//		'false', they are classified as conflicted and an
//		error will be returned.
//
//	getOtherNonRegularFileInfos			bool
//
//		If this parameter is set to 'true', Other
//		Non-Regular files, which also meet the File
//		Selection Characteristics criteria
//		(fileSelectCharacteristics), will be included in the
//		os.FileInfo information ('fileInfos') returned by
//		this method.
//
//		Other Non-regular files include	device files,
//		named pipes, and sockets. For an explanation of
//		Regular and Non-Regular files, see the section on
//		"Definition Of Terms", above.
//
//		If input parameters 'getSubdirectoryFileInfos',
//		'getRegularFileInfos', 'getSymLinksFileInfos' and
//		'getOtherNonRegularFileInfos' are all set to
//		'false', they are classified as conflicted and an
//		error will be returned.
//
//	subdirectorySelectCharacteristics	FileSelectionCriteria
//
//		In addition to the File Type Selection Criteria,
//		selected subdirectories must conform to the File
//		Characteristics Selection Criteria specified by
//		directorySelectCriteria.
//
//		Directory Files matching this selection criteria,
//		and the	File Type filter specified by input
//		parameter 'getSubdirectoryFileInfos', will be
//		included in the array of os.FileInfo objects
//		returned by	this method.
//
//		Remember that setting 'directorySelectCriteria'
//		to an empty instance of FileSelectionCriteria will
//		ensure that all directories are selected.
//
//			Example:
//			subdirectorySelectCharacteristics =
//				FileSelectionCriteria{}
//
//			This ensures that all subdirectories will satisfy
//			the Directory Characteristics Selection Criteria.
//
//		For a detailed explanation of the File Characteristics
//		Criteria specifications offered by Type
//		FileSelectionCriteria, see the documentation for
//		'fileSelectCharacteristics', below.
//
//	fileSelectCharacteristics			FileSelectionCriteria
//
//		In addition to the File Type Selection Criteria,
//		selected files must conform to the File
//		Characteristics Criteria specified by
//		'fileSelectCharacteristics'.
//
//		Files matching these selection criteria, and the
//		File Type filter, will be included in the array of
//		FileInfoPlus objects containing os.FileInfo
//		information returned by this method.
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
//	  The FileSelectionCriteria type allows for
//	  configuration of single or multiple file selection
//	  criterion. The 'SelectCriterionMode' can be used to
//	  specify whether the file must match all, or any one,
//	  of the active file selection criterion.
//
//	  Elements of the FileSelectionCriteria are described
//	  below:
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
//		If all file selection criterion in the FileSelectionCriteria object
//		are 'Inactive' or 'Not Set' (set to their zero or default values),
//		then all the files and subdirectories in the directory path specified
//		by input parameter 'dMgr' will be selected.
//
//			Example:
//			  fsc := FileSelectCriterionMode{}
//
//		 	In this example, 'fsc' is NOT initialized. Therefore,
//			all the selection criterion are 'Inactive'. Consequently,
//			all subdirectories and files encountered in the target directory
//			path during the search operation will meet the file characteristics
//			selection criteria and will therefore be classified as eligible for
//			selection.
//
//		------------------------------------------------------------------------
//
//	dMgrLabel							string
//
//		The name or label associated with input parameter
//		'dMgr', which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "dMgr" will be
//		automatically applied.
//
//	errPrefDto							*ePref.ErrPrefixDto
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
//	fileInfos							[]FileInfoPlus
//
//		If this method completes successfully, this
//		method will return an array of FileInfoPlus
//		objects containing os.FileInfo data on the
//		subdirectories and files contained in the
//		directory path specified by input parameter
//		'dMgr' which meet the File Type and File
//		Characteristics Selection Criteria.
//
//		The returned Type, FileInfoPlus, implements the
//		os.FileInfo interface, but provides additional
//		file information over and above that provided by
//		the standard os.FileInfo interface.
//
//		The os.FileInfo interface is defined as follows:
//
//	 	type FileInfo interface {
//			 Name() string
//				Base name of the file
//
//			 Size() int64
//			 	Length in bytes for regular files;
//			 	system-dependent for others
//
//			 Mode() FileMode
//			 	File mode bits
//
//			 ModTime() time.Time
//			 	Modification time
//
//			 IsDir() bool
//			 	Abbreviation for Mode().IsDir()
//
//			 Sys() interface{}
//			 	Underlying data source (can return nil)
//	 	}
//
//	lenFileInfos						int
//
//		This returned integer value specifies the length
//		of the array of os.FileInfo objects, returned as
//		'fileInfos'.
//
//	nonfatalErrs						[]error
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
//		Non-fatal errors usually involve failures
//		associated with reading individual files.
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
//	fatalErr							error
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
func (dMgrHlprMolecule *dirMgrHelperMolecule) lowLevelGetFileInfosFromDir(
	dMgr *DirMgr,
	getSubdirectoryFileInfos bool,
	includeSubDirCurrenDirOneDot bool,
	includeSubDirParentDirTwoDots bool,
	getRegularFileInfos bool,
	getSymLinksFileInfos bool,
	getOtherNonRegularFileInfos bool,
	subdirectorySelectCharacteristics FileSelectionCriteria,
	fileSelectCharacteristics FileSelectionCriteria,
	dMgrLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	fileInfos []FileInfoPlus,
	lenFileInfos int,
	nonfatalErrs []error,
	fatalErr error) {

	if dMgrHlprMolecule.lock == nil {
		dMgrHlprMolecule.lock = new(sync.Mutex)
	}

	dMgrHlprMolecule.lock.Lock()

	defer dMgrHlprMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	funcName := "dirMgrHelperMolecule." +
		"lowLevelGetFileInfosFromDir()"

	ePrefix,
		fatalErr = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		return fileInfos, lenFileInfos, nonfatalErrs, fatalErr
	}

	if len(dMgrLabel) == 0 {

		dMgrLabel = "dMgr"
	}

	if dMgr == nil {

		fatalErr = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is invalid!\n"+
			"'%v' is a 'nil' pointer.\n",
			ePrefix.String(),
			dMgrLabel,
			dMgrLabel)

		return fileInfos, lenFileInfos, nonfatalErrs, fatalErr
	}

	if getSubdirectoryFileInfos == false &&
		getRegularFileInfos == false &&
		getSymLinksFileInfos == false &&
		getOtherNonRegularFileInfos == false {

		fatalErr = fmt.Errorf("%v\n"+
			"Fatal Error: File Type filters are conflicted!\n"+
			"All of the File Type filters are set to 'false'\n"+
			"This gurantees that NO files will be selected.\n"+
			"getSubdirectoryFileInfos == false\n"+
			"getRegularFileInfos == false\n"+
			"getSymLinksFileInfos == false\n"+
			"getOtherNonRegularFileInfos == false\n",
			ePrefix.String())

		return fileInfos, lenFileInfos, nonfatalErrs, fatalErr
	}

	var isSelectAllFileTypes = false

	if getRegularFileInfos == true &&
		getSymLinksFileInfos == true &&
		getOtherNonRegularFileInfos == true {

		isSelectAllFileTypes = true

	}

	isSubdirectorySelectionCriteriaActive :=
		subdirectorySelectCharacteristics.IsSelectionCriteriaActive()

	isFileSelectionCriteriaActive :=
		fileSelectCharacteristics.IsSelectionCriteriaActive()

	var err2 error
	var nameDirEntries []os.DirEntry

	nameDirEntries,
		err2 = os.ReadDir(dMgr.absolutePath)

	if err2 != nil {

		fatalErr = fmt.Errorf("%v\n"+
			"Error returned by os.ReadDir(%v.absolutePath).\n"+
			"%v.absolutePath='%v'\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			dMgrLabel,
			dMgrLabel,
			dMgr.absolutePath,
			err2.Error())

		return fileInfos, lenFileInfos, nonfatalErrs, fatalErr
	}

	lenFileInfos = len(nameDirEntries)

	if lenFileInfos == 0 {
		// There are no files!
		return fileInfos, lenFileInfos, nonfatalErrs, fatalErr
	}

	var osFInfo os.FileInfo

	fip := FileInfoPlus{}

	osPathSepStr := string(os.PathSeparator)

	var isFileTypeMatch bool

	var isFileCharacteristicsMatch bool

	var fh = new(FileHelper)

	for i := 0; i < lenFileInfos; i++ {

		osFInfo,
			err2 = nameDirEntries[i].Info()

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Error: Error Returned by nameDirEntry.Info().\n"+
				"The conversion of DirEntry to os.FileInfo Failed."+
				"%v= '%v'\n"+
				"FileName= '%v'\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				dMgrLabel,
				dMgr.absolutePath,
				dMgr.absolutePath+osPathSepStr+nameDirEntries[i].Name(),
				err2.Error())

			nonfatalErrs = append(nonfatalErrs, err)

			continue
		}

		isFileTypeMatch = false
		isFileCharacteristicsMatch = false

		if osFInfo.IsDir() {

			if !getSubdirectoryFileInfos {
				// The user did NOT request
				// subdirectories!
				continue
			}

			// Process Current Directory Entry
			if osFInfo.Name() == "." {

				if includeSubDirCurrenDirOneDot == true {

					fileInfos = append(
						fileInfos, fip.NewFromFileInfo(osFInfo))

				}

				continue
			}

			// Process Parent Directory Entry
			if osFInfo.Name() == ".." {

				if includeSubDirParentDirTwoDots == true {

					fileInfos = append(
						fileInfos, fip.NewFromFileInfo(osFInfo))

				}

				continue
			}

			// This is a subdirectory. Subdirectories
			// have been requested by the user.
			// Does it meet File Characteristics
			// Criteria?

			if isSubdirectorySelectionCriteriaActive == false {

				isFileCharacteristicsMatch = true

			} else {
				// Directory File Characteristics Filter
				// is ACTIVE!

				isFileCharacteristicsMatch,
					err2,
					_ =
					fh.FilterFileName(
						osFInfo,
						subdirectorySelectCharacteristics,
						ePrefix.XCpy("osFInfo-Directory"))

				if err2 != nil {

					err =
						fmt.Errorf("%v\n"+
							"Error returned by fh.FilterFileName(osFInfo, fileSelectCharacteristics).\n"+
							"%v directorySearched='%v'\n"+
							"Directory Name='%v'\n"+
							"Directory Path='%v'\n"+
							"Error= \n%v\n",
							funcName,
							dMgrLabel,
							dMgr.absolutePath,
							osFInfo.Name(),
							dMgr.absolutePath+
								osPathSepStr+
								osFInfo.Name(),
							err2.Error())

					nonfatalErrs = append(nonfatalErrs, err)

					continue
				}

			}

			if isFileCharacteristicsMatch {

				fileInfos = append(
					fileInfos, fip.NewFromFileInfo(osFInfo))

			}

			continue
		}

		// This is a file. NOT a subdirectory!

		if isSelectAllFileTypes == true {

			isFileTypeMatch = true

		} else if osFInfo.Mode()&os.ModeSymlink != 0 &&
			getSymLinksFileInfos {

			isFileTypeMatch = true

		} else if !osFInfo.Mode().IsRegular() &&
			getOtherNonRegularFileInfos {

			isFileTypeMatch = true

		} else if osFInfo.Mode().IsRegular() &&
			getRegularFileInfos {

			isFileTypeMatch = true

		} else {

			if getOtherNonRegularFileInfos == true {

				isFileTypeMatch = true

			} else {

				isFileTypeMatch = false
			}

		}

		if !isFileTypeMatch {
			// This file fails the File Type
			// filter test. Skip it.
			continue
		}

		// MUST BE: This file passes the File
		// Type filter test. Continue Processing.
		// Test for File Characteristics Filter.

		if isFileSelectionCriteriaActive == true {

			isFileCharacteristicsMatch,
				err2,
				_ =
				fh.FilterFileName(
					osFInfo,
					fileSelectCharacteristics,
					ePrefix.XCpy("nameFileInfo"))

			if err2 != nil {

				err =
					fmt.Errorf("%v\n"+
						"Error returned by fh.FilterFileName(nameFileInfo, fileSelectCharacteristics).\n"+
						"%v directorySearched='%v'\n"+
						"fileName='%v'\n"+
						"File Path='%v'\n"+
						"Error= \n%v\n",
						funcName,
						dMgrLabel,
						dMgr.absolutePath,
						osFInfo.Name(),
						dMgr.absolutePath+
							osPathSepStr+
							osFInfo.Name(),
						err2.Error())

				nonfatalErrs = append(nonfatalErrs, err)

				continue
			}

		} else {

			// isFileSelectionCriteriaActive == false
			isFileCharacteristicsMatch = true
		}

		if isFileCharacteristicsMatch == true {

			fileInfos = append(
				fileInfos, fip.NewFromFileInfo(osFInfo))
		}
	}

	lenFileInfos = len(fileInfos)

	return fileInfos, lenFileInfos, nonfatalErrs, fatalErr
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

	if len(dMgrLabel) == 0 {
		dMgrLabel = "dMgr"
	}

	var dMgrHlprPreon = new(dirMgrHelperPreon)

	var targetDMgrPathDoesExist bool

	_,
		targetDMgrPathDoesExist,
		err = dMgrHlprPreon.
		validateDirMgr(
			dMgr,
			false, // Path is NOT required to exist on disk
			dMgrLabel,
			ePrefix.XCpy(
				dMgrLabel))

	if !targetDMgrPathDoesExist {

		return err
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

	var dirPathDoesExist bool

	dirPathDoesExist,
		_,
		err2 =
		new(dirMgrHelperAtom).
			doesDirectoryExist(
				dMgr,
				PreProcPathCode.None(),
				dMgrLabel,
				ePrefix.XCpy(dMgrLabel))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned while testing for the\n"+
			"existance of the directory path.\n"+
			"dirMgrHelperAtom.doesDirectoryExist()\n"+
			"%v Directory Path= '%v'\n"+
			"Error= \n%v\n",
			funcName,
			dMgrLabel,
			dMgr.absolutePath,
			err2.Error())

		return err
	}

	if dirPathDoesExist {

		err = fmt.Errorf("%v\n"+
			"Error: FAILED TO DELETE DIRECTORY!!\n"+
			"Directory Path still exists!\n"+
			"Directory Path= '%v'\n",
			ePrefix.String(),
			dMgr.absolutePath)
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
