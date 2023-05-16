package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type FileOperationsElectron struct {
	lock *sync.Mutex
}

// deleteDestinationFile
//
// Deletes the destination file identified in the FileOps
// instance passed as input parameter 'fOps'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete the destination file
//	contained in the FileOps instance passed as input
//	parameter 'fOps'. The 'fOps' internal member variable
//	specifying the destination file is 'fOps.destination'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fOps 						*FileOps
//
//		A pointer to an instance of FileOps. The file
//	 	identified by the internal member variable
//	 	'fOps.destination' will be deleted.
//
//		If 'fOps' has not been properly initialized,
//		an error will be returned.
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
func (fOpsElectron *FileOperationsElectron) deleteDestinationFile(
	fOps *FileOps,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fOpsElectron.lock == nil {
		fOpsElectron.lock = new(sync.Mutex)
	}

	fOpsElectron.lock.Lock()

	defer fOpsElectron.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "FileOperationsElectron." +
		"deleteDestinationFile()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return err
	}

	if fOps == nil {
		err = fmt.Errorf("%v\n"+
			"Error: FileOps instance is invalid!\n"+
			"Input parameter 'fOps' is a nil pointer.\n",
			ePrefix.String())

		return err
	}

	if !fOps.isInitialized {

		err = fmt.Errorf("%v\n"+
			"Error: FileOps instance is invalid!\n"+
			"Input parameter 'fOps' has NOT been properly initialized.\n"+
			"fOps.isInitialized = 'false'.",
			ePrefix.String())

		return err
	}

	var err2 error

	err2 = fOps.destination.
		DeleteThisFile(ePrefix.XCpy(
			"fOps.destination"))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"A an error occured while attempting to\n"+
			"delete the 'fOps.destination' file.\n"+
			"fOps.destination file = '%v'\n"+
			"Error= \n%v\n",
			funcName,
			fOps.destination.GetAbsolutePathFileName(),
			err2.Error())
	}

	return err
}

// deleteSourceFile
//
// Deletes the source file identified in the FileOps
// instance passed as input parameter 'fOps'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete the source file contained in
//	the FileOps instance passed as input parameter
//	'fOps'. The 'fOps' internal member variable
//	specifying the source file is 'fOps.source'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fOps 						*FileOps
//
//		A pointer to an instance of FileOps. The file
//	 	identified by the internal member variable
//	 	'fOps.source' will be deleted.
//
//		If 'fOps' has not been properly initialized,
//		an error will be returned.
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
func (fOpsElectron *FileOperationsElectron) deleteSourceFile(
	fOps *FileOps,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fOpsElectron.lock == nil {
		fOpsElectron.lock = new(sync.Mutex)
	}

	fOpsElectron.lock.Lock()

	defer fOpsElectron.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "FileOperationsElectron." +
		"deleteSourceFile()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return err
	}

	if fOps == nil {
		err = fmt.Errorf("%v\n"+
			"Error: FileOps instance is invalid!\n"+
			"Input parameter 'fOps' is a nil pointer.\n",
			ePrefix.String())

		return err
	}

	if !fOps.isInitialized {

		err = fmt.Errorf("%v\n"+
			"Error: FileOps instance is invalid!\n"+
			"Input parameter 'fOps' has NOT been properly initialized.\n"+
			"fOps.isInitialized = 'false'.",
			ePrefix.String())

		return err
	}

	var err2 error

	err2 = fOps.source.DeleteThisFile(ePrefix)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"A an error occured while attempting to\n"+
			"delete the 'fOps.source' file.\n"+
			"fOps.source file = '%v'\n"+
			"Error= \n%v\n",
			funcName,
			fOps.source.GetAbsolutePathFileName(),
			err2.Error())
	}

	return err
}

// empty
//
// This method receives a pointer to an instance of
// FileOps and proceeds to delete and reset all the
// contained member variable data values to their zero
// value or uninitialized state.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and reset all pre-existing
//	data values in the current instance of FileOps to
//	their zero value or uninitialized state.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fOps						*FileOps
//
//		A pointer to an instance of FileOps. All member
//		variable data values contained in this instance
//		will be deleted and reset to their zero value or
//		uninitialized state.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	-- NONE --
func (fOpsElectron *FileOperationsElectron) empty(
	fOps *FileOps) {

	if fOpsElectron.lock == nil {
		fOpsElectron.lock = new(sync.Mutex)
	}

	fOpsElectron.lock.Lock()

	defer fOpsElectron.lock.Unlock()

	if fOps == nil {

		return
	}

	fOps.source.Empty()

	fOps.destination.Empty()

	fOps.opToExecute = FileOpCode.None()

	return
}

// setFileOpsDestination
//
// Reconfigures the 'destination' File Manager
// encapsulated in the FileOps instance passed as input
// parameter 'fOps'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and overwrite the pre-existing
//	destination File Manager configured in the FileOps instance
//	passed as input parameter 'fOps'.
//
//	The 'fOps' internal member variable which will be reset
//	to a new value is identified as 'fOps.destination'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fOps 						*FileOps
//
//		A pointer to an instance of FileOps. The
//		destination File Manager object contained in this
//		instance will be reconfigured with destination
//		file information taken from input parameter
//		'destinationFMgr'.
//
//	destinationFMgr				FileMgr
//
//		A deep copy of this File Manager instance will be
//		used to reconfigure the 'fOps' destination member
//		variable 'fOps.destination'.
//
//		If 'destinationFMgr' evaluates as invalid, an error
//		will be returned.
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
func (fOpsElectron *FileOperationsElectron) setFileOpsDestination(
	fOps *FileOps,
	destinationFMgr FileMgr,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fOpsElectron.lock == nil {
		fOpsElectron.lock = new(sync.Mutex)
	}

	fOpsElectron.lock.Lock()

	defer fOpsElectron.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "FileOperationsElectron." +
		"setFileOpsDestination()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return err
	}

	if fOps == nil {
		err = fmt.Errorf("%v\n"+
			"Error: FileOps instance is invalid!\n"+
			"Input parameter 'fOps' is a nil pointer.\n",
			ePrefix.String())

		return err
	}

	fMgrHelperAtom := new(fileMgrHelperAtom)

	err = fMgrHelperAtom.isFileMgrValid(
		&destinationFMgr,
		ePrefix.XCpy("destinationFMgr"))

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Input parameter 'destinationFMgr' is invalid!\n"+
			"Error= \n%v\n",
			funcName,
			err.Error())

	}

	err = fOps.destination.CopyIn(
		&destinationFMgr,
		ePrefix.XCpy(
			"fOps.destination<-destinationFMgr"))

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Error Copying Destination File Manager to\n"+
			"fOps.destination."+
			"Error= \n%v\n",
			funcName,
			err.Error())
	}

	err = fMgrHelperAtom.isFileMgrValid(
		&fOps.source,
		ePrefix.XCpy("fOps.source"))

	if err == nil {

		fOps.isInitialized = true

	} else {

		err = nil
	}

	return err
}

// setFileOpsSource
//
// Reconfigures the 'source' File Manager encapsulated in
// the FileOps instance passed as input parameter 'fOps'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and overwrite the pre-existing
//	source File Manager configured in the FileOps instance
//	passed as input parameter 'fOps'.
//
//	The 'fOps' internal member variable which will be reset
//	to a new value is identified as 'fOps.source'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fOps 						*FileOps
//
//		A pointer to an instance of FileOps. The source
//		File Manager object contained in this instance
//		will be reconfigured with source file information
//		taken from input parameter 'sourceFMgr'.
//
//	sourceFMgr					FileMgr
//
//		A deep copy of this File Manager instance will be
//		used to reconfigure the 'fOps' source member
//		variable 'fOps.source'.
//
//		If 'sourceFMgr' evaluates as invalid, an error
//		will be returned.
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
func (fOpsElectron *FileOperationsElectron) setFileOpsSource(
	fOps *FileOps,
	sourceFMgr FileMgr,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fOpsElectron.lock == nil {
		fOpsElectron.lock = new(sync.Mutex)
	}

	fOpsElectron.lock.Lock()

	defer fOpsElectron.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "FileOperationsElectron." +
		"setFileOpsSource()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return err
	}

	if fOps == nil {
		err = fmt.Errorf("%v\n"+
			"Error: FileOps instance is invalid!\n"+
			"Input parameter 'fOps' is a nil pointer.\n",
			ePrefix.String())

		return err
	}

	fMgrHelperAtom := new(fileMgrHelperAtom)

	err = fMgrHelperAtom.isFileMgrValid(
		&sourceFMgr,
		ePrefix.XCpy("sourceFMgr"))

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Input parameter 'sourceFMgr' is invalid!\n"+
			"Error= \n%v\n",
			funcName,
			err.Error())

	}

	err = fOps.source.CopyIn(
		&sourceFMgr,
		ePrefix.XCpy(
			"fOps.source<-sourceFMgr"))

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Error Copying Source File Manager to\n"+
			"fOps.source."+
			"Error= \n%v\n",
			funcName,
			err.Error())
	}

	err = fMgrHelperAtom.isFileMgrValid(
		&fOps.destination,
		ePrefix.XCpy("fOps.destination"))

	if err == nil {

		fOps.isInitialized = true

	} else {

		err = nil
	}

	return err
}
