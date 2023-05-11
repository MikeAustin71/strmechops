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
