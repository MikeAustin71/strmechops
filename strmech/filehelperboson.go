package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"os"
	"sync"
)

type fileHelperBoson struct {
	lock *sync.Mutex
}

// openFile
//
// This method is a wrapper for os.OpenFile. This method
// may be used to open or create files depending on the
// File Open and File Permission parameters.
//
// If successful, this method will return a pointer to
// the os.File object associated with the file designated
// for opening.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	The calling routine is responsible for calling
//	"Close()" on this os.File pointer.
//
//	Calling "Close()" multiple times on an os.File
//	pointer will generate a 'panic' error.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathFileName					string
//
//		A string containing the path and file name of the
//		file which will be opened. If a parent path
//		component does NOT exist, this method will
//		trigger an error.
//
//	createDirectoryPathIfNotExist	bool
//
//		If the directory path element of parameter
//		'pathFileName' does not exist on an attached
//		storage drive, and this parameter is set to
//		'true', this method will attempt to create
//		the directory path.
//
//		If 'createDirectoryPathIfNotExist' is set to
//		'false', and the directory path element of
//		parameter 'pathFileName' does not exist on an
//		attached storage drive, an error will be returned.
//
//	fileOpenCfg						FileOpenConfig
//
//		This parameter encapsulates the File Open
//		parameters which will be used to open subject
//		file. For an explanation of File Open parameters,
//		see method:
//			FileOpenConfig.New().
//
//	filePermissionCfg				FilePermissionConfig
//
//		This parameter encapsulates the File Permission
//		parameters which will be used to open the subject
//		file. For an explanation of File Permission
//		parameters, see method:
//			FilePermissionConfig.New().
//
//	pathFileNameLabel				string
//
//		The name or label associated with input parameter
//		'pathFileName' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "pathFileName" will be
//		automatically applied.
//
//	errPrefDto						*ePref.ErrPrefixDto
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
//	*os.File
//
//		If successful, this method returns an os.File
//		pointer to the file designated by input parameter
//		'pathFileName'. This file pointer can
//		subsequently be used for reading content from the
//		subject file. It may NOT be used for writing
//		content to the subject file.
//
//		If this method fails, the *os.File return value
//		is 'nil'.
//
//		Note:
//		The caller is responsible for calling "Close()"
//		on this os.File pointer.
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
func (fileHelpBoson *fileHelperBoson) openFile(
	pathFileName string,
	createDirectoryPathIfNotExist bool,
	fileOpenCfg FileOpenConfig,
	filePermissionCfg FilePermissionConfig,
	pathFileNameLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	filePtr *os.File,
	err error) {

	if fileHelpBoson.lock == nil {
		fileHelpBoson.lock = new(sync.Mutex)
	}

	fileHelpBoson.lock.Lock()

	defer fileHelpBoson.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	filePtr = nil

	funcName := "fileHelperBoson.openFile()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return filePtr, err
	}

	if len(pathFileNameLabel) == 0 {
		pathFileNameLabel = "pathFileName"
	}

	if len(pathFileName) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is invalid!\n"+
			"'%v' is empty and has a zero string length.\n",
			ePrefix.String(),
			pathFileNameLabel,
			pathFileNameLabel)

		return filePtr, err
	}

	var err2 error

	pathFileName,
		err2 = new(fileHelperMolecule).
		getAbsPathFromFilePath(
			pathFileName,
			ePrefix.XCpy(
				"<-pathFileName"))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"'%v' conversion to an absolute path failed!\n"+
			"Error returned by fileHelperMolecule.getAbsPathFromFilePath()\n"+
			"%v= '%v'\n"+
			"Error= \n%v\n",
			funcName,
			pathFileNameLabel,
			pathFileNameLabel,
			pathFileName,
			err2.Error())

		return filePtr, err
	}

	var directoryPath, fileNameExt string
	var bothAreEmpty bool

	directoryPath,
		fileNameExt,
		bothAreEmpty,
		err2 = new(fileHelperDirector).
		getPathAndFileNameExt(
			pathFileName,
			pathFileNameLabel,
			ePrefix.XCpy("<-pathFileName"))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is invalid!\n"+
			"An error occurred while breaking '%v'\n"+
			"into directory path, file name and file extension\n"+
			"components.\n"+
			"Error = \n%v\n",
			funcName,
			pathFileNameLabel,
			pathFileNameLabel,
			err2.Error())

		return filePtr, err
	}

	if len(fileNameExt) == 0 || bothAreEmpty {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is invalid!\n"+
			"No valid file name could be extracted from\n"+
			"'%v'.\n"+
			"%v= '%v'\n"+
			"Directory Path Element= '%v'\n"+
			"File Name Element= '%v'\n",
			ePrefix.String(),
			pathFileNameLabel,
			pathFileNameLabel,
			pathFileNameLabel,
			pathFileName,
			directoryPath,
			fileNameExt)

		return filePtr, err
	}

	var pathFileNameDoesExist bool

	var fInfoPlus FileInfoPlus

	pathFileNameDoesExist,
		fInfoPlus,
		err2 = new(fileHelperAtom).doesDirectoryExist(
		directoryPath,
		pathFileNameLabel+" Dir Path",
		ePrefix.XCpy(pathFileNameLabel+" Dir Path"))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by fileHelperAtom.doesDirectoryExist()\n"+
			"%v= '%v'\n"+
			"Error=\n%v\n",
			funcName,
			pathFileNameLabel,
			pathFileName,
			err2.Error())

		return filePtr, err
	}

	if !fInfoPlus.IsDir() {

		err = fmt.Errorf("%v\n"+
			"Error: The directory path extracted from '%v'\n"+
			"is NOT a valid directory. '%v' is therefore invalid!\n"+
			"%v= '%v'\n",
			ePrefix.String(),
			pathFileNameLabel,
			pathFileNameLabel,
			pathFileNameLabel,
			pathFileName)

		return filePtr, err
	}

	if pathFileNameDoesExist {
		// The 'pathFileName' directory path does NOT
		// exist on an attached storage volume.

		if createDirectoryPathIfNotExist {

			err2 = new(fileHelperMechanics).makeDirAll(
				directoryPath,
				pathFileNameLabel+" directoryPath",
				ePrefix)

			if err2 != nil {

				err = fmt.Errorf("%v\n"+
					"Attempted creation of directory path failed!\n"+
					"%v= '%v'\n"+
					"%v Directory Path= '%v'\n"+
					"Error returned by fileHelperMechanics.makeDirAll()\n"+
					"Error=\n%v\n",
					funcName,
					pathFileNameLabel,
					pathFileName,
					pathFileNameLabel,
					directoryPath,
					err2.Error())

				return filePtr, err
			}

		} else {
			// The Path File Name Directory DOES NOT EXIST
			// on an attached storage drive and
			// createDirectoryPathIfNotExist = 'false'.

			err = fmt.Errorf("%v\n"+
				"Error: The %v Directory\n"+
				"does NOT exist on an attached storage drive and\n"+
				"Input Parameter 'createDirectoryPathIfNotExist'\n"+
				"was set to 'false'. Therefore the file cannot be\n"+
				"opened.\n"+
				"%v= '%v\n"+
				"%v Directory = '%v'\n",
				ePrefix.String(),
				pathFileNameLabel,
				pathFileNameLabel,
				pathFileName,
				pathFileNameLabel,
				directoryPath)

			return filePtr, err
		}

	}

	// At this point the 'pathFileName' directory exists
	// on an attached storage drive. Time to open the file.

	err2 = fileOpenCfg.IsValidInstanceError(
		ePrefix.XCpy("fileOpenCfg"))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Input Parameter 'fileOpenCfg' is INVALID!\n"+
			"Error='%v'\n",
			ePrefix.String(),
			err2.Error())

		return filePtr, err
	}

	var fOpenCode int

	fOpenCode,
		err2 = fileOpenCfg.
		GetCompositeFileOpenCode(
			ePrefix.XCpy("fileOpenCfg"))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Input parameter 'fileOpenCfg' is invalid!\n"+
			"fileOpenCfg caused an error return from \n"+
			"fileOpenCfg.GetGetCompositeFileOpenCode()\n"+
			"Error = \n%v\n",
			funcName,
			err2.Error())

		return filePtr, err
	}

	err2 = filePermissionCfg.IsValidInstanceError(
		ePrefix.XCpy("filePermissionCfg"))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Input Parameter 'filePermissionCfg' is INVALID!\n"+
			"filePermissionCfg.IsValidInstanceError() returned\n"+
			"an error.\n"+
			"Error= \n%v'\n",
			funcName,
			err2.Error())

		return filePtr, err
	}

	var fileMode os.FileMode

	fileMode,
		err2 = filePermissionCfg.
		GetCompositePermissionMode(ePrefix.XCpy(
			"fileMode<-"))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Input parameter 'filePermissionCfg' is invalid!\n"+
			"filePermissionCfg.GetCompositePermissionMode()\n"+
			"returned an error.\n"+
			"Error= \n%v\n",
			funcName,
			err2.Error())

		return filePtr, err
	}

	var err3 error

	filePtr,
		err2 = os.OpenFile(
		pathFileName,
		fOpenCode,
		fileMode)

	if err2 != nil {

		if filePtr != nil {

			err3 = filePtr.Close()

		}

		if err3 != nil {

			err = fmt.Errorf("%v\n"+
				"There were two errors!\n"+
				"Error #1\n"+
				"Error #1 returned by os.OpenFile(%v, fOpenCode, fileMode).\n"+
				"%v= '%v'\n"+
				"Error #1=\n%v\n\n"+
				"Error #2\n"+
				"Error #2 returned by filePtr.Close() while attempting to\n"+
				"close the file pointer returned by os.OpenFile().\n"+
				"Error #2=\n%v\n",
				ePrefix.String(),
				pathFileNameLabel,
				pathFileNameLabel,
				pathFileName,
				err2.Error(),
				err3.Error())

			return filePtr, err

		} else {

			err = fmt.Errorf("%v\n"+
				"Error returned by os.OpenFile(%v, fOpenCode, fileMode).\n"+
				"%v= '%v'\n"+
				"Error=\n%v\n\n",
				ePrefix.String(),
				pathFileNameLabel,
				pathFileNameLabel,
				pathFileName,
				err2.Error())

		}

		return filePtr, err
	}

	if filePtr == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: os.OpenFile() returned a 'nil' file pointer!\n",
			ePrefix.String())
	}

	return filePtr, err
}
