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

	err = destinationFMgr.IsFileMgrValid(
		ePrefix.XCpy("destinationFMgr"))

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Input parameter 'destinationFMgr' is invalid!\n"+
			"Error= \n%v\n",
			funcName,
			err.Error())

	}

	var newDestinationFMgr FileMgr

	newDestinationFMgr,
		err = destinationFMgr.CopyOut(ePrefix.XCpy(
		"newDestinationFMgr"))

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Deep Copy of Destination File Manager Failed!\n"+
			"Error= \n%v\n",
			funcName,
			err.Error())
	}

	err = fOps.destination.CopyIn(
		&newDestinationFMgr,
		ePrefix.XCpy(
			"fOps.destination<-newDestinationFMgr"))

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Error Copying Destination File Manager to\n"+
			"fOps.destination."+
			"Error= \n%v\n",
			funcName,
			err.Error())
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

	err = sourceFMgr.IsFileMgrValid(
		ePrefix.XCpy("sourceFMgr"))

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Input parameter 'sourceFMgr' is invalid!\n"+
			"Error= \n%v\n",
			funcName,
			err.Error())

	}

	var newSourceFMgr FileMgr

	newSourceFMgr,
		err = sourceFMgr.CopyOut(ePrefix.XCpy(
		"newSourceFMgr"))

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Deep Copy of Source File Manager Failed!\n"+
			"Error= \n%v\n",
			funcName,
			err.Error())
	}

	err = fOps.source.CopyIn(
		&newSourceFMgr,
		ePrefix.XCpy(
			"fOps.source<-newSourceFMgr"))

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Error Copying Source File Manager to\n"+
			"fOps.source."+
			"Error= \n%v\n",
			funcName,
			err.Error())
	}

	return err
}
