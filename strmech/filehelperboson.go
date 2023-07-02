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
//	pathFileName				string
//
//		A string containing the path and file name of the
//		file which will be opened. If a parent path
//		component does NOT exist, this method will
//		trigger an error.
//
//	fileOpenCfg					FileOpenConfig
//
//		This parameter encapsulates the File Open
//		parameters which will be used to open subject
//		file. For an explanation of File Open parameters,
//		see method:
//			FileOpenConfig.New().
//
//	filePermissionCfg			FilePermissionConfig
//
//		This parameter encapsulates the File Permission
//		parameters which will be used to open the subject
//		file. For an explanation of File Permission
//		parameters, see method:
//			FilePermissionConfig.New().
//
//	pathFileNameLabel			string
//
//		The name or label associated with input parameter
//		'pathFileName' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "pathFileName" will be
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

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errPrefDto,
		"fileHelperBoson."+
			"openFile()",
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

	err2 := fileOpenCfg.IsValidInstanceError(ePrefix)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Input Parameter 'fileOpenCfg' is INVALID!\n"+
			"Error='%v'\n",
			ePrefix.String(),
			err2.Error())

		return filePtr, err
	}

	fOpenCode, err2 := fileOpenCfg.
		GetCompositeFileOpenCode(ePrefix)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"%v\n",
			ePrefix.String(),
			err2.Error())

		return filePtr, err
	}

	err2 = filePermissionCfg.IsValidInstanceError(
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
