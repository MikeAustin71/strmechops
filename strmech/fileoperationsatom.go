package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type FileOperationsAtom struct {
	lock *sync.Mutex
}

// deleteSourceAndDestinationFiles
//
// Deletes the source and destination files identified in
// the FileOps instance passed as input parameter 'fOps'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete both the destination file and
//	the source file contained in the FileOps instance
//	passed as input parameter 'fOps'.
//
//	The 'fOps' internal member variable specifying the
//	source file is 'fOps.source'.
//
//	The 'fOps' internal member variable specifying the
//	destination file is 'fOps.destination'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fOps 						*FileOps
//
//		A pointer to an instance of FileOps. The two
//		files identified by the internal member variables
//	 	'fOps.source' and 'fOps.destination' will be
//	 	deleted.
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
func (fOpsAtom *FileOperationsAtom) deleteSourceAndDestinationFiles(
	fOps *FileOps,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fOpsAtom.lock == nil {
		fOpsAtom.lock = new(sync.Mutex)
	}

	fOpsAtom.lock.Lock()

	defer fOpsAtom.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "FileOperationsAtom." +
		"deleteSourceAndDestinationFiles()"

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

	fOpsElectron := FileOperationsElectron{}

	err = fOpsElectron.deleteSourceFile(
		fOps,
		ePrefix)

	if err != nil {
		return err
	}

	err = fOpsElectron.deleteDestinationFile(
		fOps,
		ePrefix)

	return err
}

// setFileOps
//
// Reconfigures an instance of FileOps passed as input
// parameter 'fOps'.
//
// 'fOps' is reconfigured using a source File Manager and
// a destination File Manager passed as input parameters.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and reconfigure all
//	pre-existing data values in the FileOps instance
//	passed as input parameter 'fOps'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fOps						FileOps
//
//		An instance of FileOps. This instance will
//		be reconfigured using information provided by
//		input parameters 'sourceFMgr' and
//		'destinationFMgr'.
//
//	sourceFMgr					FileMgr
//
//		An instance of FileMgr. A deep copy of this
//		object will be used to configure the source File
//		Manager contained in 'fOps'.
//
//			fOps.source = sourceFMgr
//
//	destinationFMgr				FileMgr
//
//		An instance of FileMgr. A deep copy of this
//		object will be used to configure the destination
//		File Manager contained in 'fOps'.
//
//			fOps.destination = destinationFMgr
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
func (fOpsAtom *FileOperationsAtom) setFileOps(
	fOps *FileOps,
	sourceFMgr FileMgr,
	destinationFMgr FileMgr,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fOpsAtom.lock == nil {
		fOpsAtom.lock = new(sync.Mutex)
	}

	fOpsAtom.lock.Lock()

	defer fOpsAtom.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "FileOperationsAtom." +
		"setFileOps()"

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
			"Source File Manager is INVALID!\n"+
			"Error= \n%v\n",
			funcName,
			err.Error())
	}

	err = destinationFMgr.IsFileMgrValid(
		ePrefix.XCpy("destinationFMgr"))

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Destination File Manager is INVALID!\n"+
			"Error= \n%v\n",
			funcName,
			err.Error())
	}

	var newSourceFMgr, newDestinationFMgr FileMgr

	newSourceFMgr,
		err = sourceFMgr.CopyOut(ePrefix.XCpy(
		"sourceFMgr"))

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Deep Copy of Source File Manager Failed!\n"+
			"Error= \n%v\n",
			funcName,
			err.Error())
	}

	newDestinationFMgr,
		err = destinationFMgr.CopyOut(ePrefix.XCpy(
		"destinationFMgr"))

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Deep Copy of Destination File Manager Failed!\n"+
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

	fOps.isInitialized = true

	return err
}
