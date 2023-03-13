package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// fileAccessControlMechanics
//
// Provides helper methods for Type
// FileAccessControl.
type fileAccessControlMechanics struct {
	lock *sync.Mutex
}

// setFileAccessControl
//
// Receives an instance of type FileAccessControl and
// configures that instance with the values provided
// by input parameters 'openCodes' and 'permissions'.
//
// The FileAccessControl type is used when opening files
// for read and write operations.
//
// To open a file, two components are required:
//
//  1. A FileOpenType - Input parameter FileOpenConfig
//     In order to open a file, exactly one of the
//     following File Open Codes MUST be specified:
//
//     FileOpenType(0).TypeReadOnly()
//     FileOpenType(0).TypeWriteOnly()
//     FileOpenType(0).TypeReadWrite()
//
//     -- AND --
//
//  2. A FileOpenMode - Input parameter FilePermissionConfig
//
//     In addition to a 'FileOpenType', a File Open Mode
//     is also required. This code is also referred to as
//     'permissions'. Zero or more of the following File
//     Open Mode codes may optionally be specified to
//     better control file open behavior.
//
//     FileOpenMode(0).ModeAppend()
//     FileOpenMode(0).ModeCreate()
//     FileOpenMode(0).ModeExclusive()
//     FileOpenMode(0).ModeSync()
//     FileOpenMode(0).ModeTruncate()
//
// Type FileAccessControl encapsulates these two
// components required for file access operations.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fAccessCtrl					*FileAccessControl
//
//		A pointer to an instance of FileAccessControl.
//
//		This FileAccessControl object will be
//		reconfigured with the File Open Type and File
//		Open mode values provided by input parameters
//		'openCodes' and 'permissions'.
//
//	openCodes					FileOpenConfig
//
//		This parameter encapsulates the File Open
//		parameters which will be used to open subject
//		file. 'openCodes' are also referred to as the
//		File Open Type. For an explanation of File Open
//		parameters, see the source code documentation
//		for method FileOpenConfig.New().
//
//	permissions					FilePermissionConfig
//
//		This parameter encapsulates the File Permission
//		parameters which will be used to open the subject
//		file. 'permissions' is also referred to as the
//		File Open Mode. For an explanation of File
//		Permission parameters, see method FilePermissionConfig.New().
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
func (fAccessMech *fileAccessControlMechanics) setFileAccessControl(
	fAccessCtrl *FileAccessControl,
	openCodes FileOpenConfig,
	permissions FilePermissionConfig,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fAccessMech.lock == nil {
		fAccessMech.lock = new(sync.Mutex)
	}

	fAccessMech.lock.Lock()

	defer fAccessMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	funcName := "fileAccessControlMechanics." +
		"setFileAccessControl()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return err
	}

	if fAccessCtrl == nil {

		err = fmt.Errorf("%v\n"+
			"Error: FileAccessControl instance is invalid!\n"+
			"Input parameter 'fAccessCtrl' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	err = openCodes.IsValidInstanceError(ePrefix)

	if err != nil {

		return fmt.Errorf(
			"%v\n"+
				"Input parameter 'openCodes' is INVALID!\n"+
				"Error= \n%v\n",
			funcName,
			err.Error())
	}

	err = permissions.IsValidInstanceError(ePrefix)

	if err != nil {
		return fmt.Errorf("%v\n"+
			"Input parameter 'permissions' is INVALID!\n"+
			"Error= \n%v\n",
			funcName,
			err.Error())
	}

	fAccessCtrl.fileOpenCodes = openCodes.CopyOut()

	fAccessCtrl.permissions = permissions.CopyOut()

	fAccessCtrl.isInitialized = true

	return err
}

// setInitializeNewFileAccessCtrl
//
// Configures a FileAccessControl instances with all
// File Open Codes and File Permission Codes initialized
// to 'None'. The FileAccessControl instance is
// essentially reconfigured with all internal data
// members set to their 'zero' or initial states.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fAccessCtrl					*FileAccessControl
//
//		A pointer to an instance of FileAccessControl.
//
//		This FileAccessControl object will be
//		reconfigured setting all internal data members
//		to their 'zero' or initial states.
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
func (fAccessMech *fileAccessControlMechanics) setInitializeNewFileAccessCtrl(
	fAccessCtrl *FileAccessControl,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fAccessMech.lock == nil {
		fAccessMech.lock = new(sync.Mutex)
	}

	fAccessMech.lock.Lock()

	defer fAccessMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	funcName := "fileAccessControlMechanics." +
		"setInitializeNewFileAccessCtrl()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return err
	}

	if fAccessCtrl == nil {
		err = fmt.Errorf("%v\n"+
			"Error: FileAccessControl instance is invalid!\n"+
			"Input parameter 'fAccessCtrl' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	openCodes,
		err := new(FileOpenConfig).New(
		ePrefix.XCpy("openCodes<-"),
		FOpenType.TypeNone(),
		FOpenMode.ModeNone())

	if err != nil {
		return fmt.Errorf("%v\n"+
			"Error returned by FileOpenConfig{}.New("+
			"FOpenType.TypeNone(), FOpenMode.ModeNone())\n"+
			"Error= \n%v\n",
			funcName,
			err.Error())
	}

	var entryType OsFilePermissionCode

	entryType,
		err = OsFilePermissionCode(0).
		GetNewFromFileMode(OsFilePermCode.ModeNone())

	if err != nil {
		return fmt.Errorf("%v\n"+
			"Error returned by OsFilePermCode.GetNewFromFileMode("+
			"OsFilePermCode.ModeNone()).\n"+
			"Error= \n%v\n ",
			ePrefix.String(),
			err.Error())
	}

	var permissions FilePermissionConfig

	permissions,
		err = new(FilePermissionConfig).
		NewByComponents(
			entryType,
			"---------",
			ePrefix.XCpy("permissions<-"))

	if err != nil {
		return fmt.Errorf("%v\n"+
			"Error returned by FilePermissionConfig{}.NewByComponents("+
			"entryType, \"---------\")\n"+
			"entryType='OsFilePermCode.ModeNone()'\n"+
			"Error= \n%v\n",
			funcName,
			err.Error())
	}

	fAccessCtrl.fileOpenCodes = openCodes.CopyOut()

	fAccessCtrl.permissions = permissions.CopyOut()

	fAccessCtrl.isInitialized = true

	return err
}
