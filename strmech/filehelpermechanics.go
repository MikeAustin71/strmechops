package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"os"
	"sync"
)

type fileHelperMechanics struct {
	lock *sync.Mutex
}

// OpenFile - wrapper for os.OpenFile. This method may be used to open or
// create files depending on the File Open and File Permission parameters.
//
// If successful, this method will return a pointer to the os.File object
// associated with the file designated for opening.
//
// The calling routine is responsible for calling "Close()" on this os.File
// pointer.
//
// ------------------------------------------------------------------------
//
// Input Parameters:
//
//	pathFileName                   string - A string containing the path and file name
//	                                        of the file which will be opened. If a parent
//	                                        path component does NOT exist, this method will
//	                                        trigger an error.
//
//	fileOpenCfg            FileOpenConfig - This parameter encapsulates the File Open parameters
//	                                        which will be used to open subject file. For an
//	                                        explanation of File Open parameters, see method
//	                                        FileOpenConfig.New().
//
// filePermissionCfg FilePermissionConfig - This parameter encapsulates the File Permission
//
//	parameters which will be used to open the subject
//	file. For an explanation of File Permission parameters,
//	see method FilePermissionConfig.New().
//
// ------------------------------------------------------------------------
//
// Return Values:
//
//	*os.File        - If successful, this method returns an os.File pointer
//	                  to the file designated by input parameter 'pathFileName'.
//	                  This file pointer can subsequently be used for reading
//	                  content from the subject file. It may NOT be used for
//	                  writing content to the subject file.
//
//	                  If this method fails, the *os.File return value is 'nil'.
//
//	                  Note: The caller is responsible for calling "Close()" on this
//	                  os.File pointer.
//
//
//	error           - If the method completes successfully, this return value
//	                  is 'nil'. If the method fails, the error type returned
//	                  is populated with an appropriate error message.
func (fileHelpMech *fileHelperMechanics) openFile(
	pathFileName string,
	fileOpenCfg FileOpenConfig,
	filePermissionCfg FilePermissionConfig,
	errorPrefix interface{}) (
	filePtr *os.File,
	err error) {

	if fileHelpMech.lock == nil {
		fileHelpMech.lock = new(sync.Mutex)
	}

	fileHelpMech.lock.Lock()

	defer fileHelpMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	filePtr = nil

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"fileHelperMechanics."+
			"openFile()",
		"")

	if err != nil {
		return filePtr, err
	}

	var pathFileNameDoesExist bool

	var fInfoPlus FileInfoPlus

	pathFileName,
		pathFileNameDoesExist,
		fInfoPlus,
		err = new(fileHelperMolecule).doesPathFileExist(
		pathFileName,
		PreProcPathCode.AbsolutePath(), // Convert to Absolute Path
		ePrefix,
		"pathFileName")

	if err != nil {
		return nil, err
	}

	if !pathFileNameDoesExist {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'pathFileName' DOES NOT EXIST!\n"+
			"pathFileName='%v'\n",
			ePrefix.String(),
			pathFileName)

		return filePtr, err
	}

	if fInfoPlus.IsDir() {
		err =
			fmt.Errorf("%v\n"+
				"ERROR: Input parameter 'pathFileName' is \n"+
				"a 'Directory' - NOT a file!\n"+
				"pathFileName='%v'\n",
				ePrefix.String(),
				pathFileName)

		return filePtr, err
	}

	err2 := fileOpenCfg.IsValid()

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Input Parameter 'fileOpenCfg' is INVALID!\n"+
			"Error='%v'\n",
			ePrefix.String(),
			err2.Error())

		return filePtr, err
	}

	fOpenCode, err2 := fileOpenCfg.GetCompositeFileOpenCode()

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"%v\n",
			ePrefix.String(),
			err2.Error())

		return filePtr, err
	}

	err2 = filePermissionCfg.IsValid(
		ePrefix)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Input Parameter 'filePermissionCfg' is INVALID!\n"+
			"Error='%v'\n",
			ePrefix.String(),
			err2.Error())

		return filePtr, err
	}

	fileMode, err2 := filePermissionCfg.
		GetCompositePermissionMode(ePrefix.XCpy(
			"fileMode<-"))

	if err2 != nil {
		err = fmt.Errorf(
			"%v\n"+"%v\n",
			ePrefix.String(),
			err2.Error())

		return filePtr, err
	}

	filePtr,
		err2 = os.OpenFile(
		pathFileName,
		fOpenCode,
		fileMode)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by os.OpenFile(pathFileName, fOpenCode, fileMode).\n"+
			"pathFileName='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			pathFileName,
			err2.Error())

		return filePtr, err
	}

	if filePtr == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: os.OpenFile() returned a 'nil' file pointer!\n",
			ePrefix.String())

		return filePtr, err
	}

	err = nil

	return filePtr, err
}
