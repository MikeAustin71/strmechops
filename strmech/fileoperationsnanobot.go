package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type FileOperationsNanobot struct {
	lock *sync.Mutex
}

// CopyOut
//
// Returns a deep copy of the current FileOps instance
// passed as input parameter 'fOps'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fOps 						*FileOps
//
//		A pointer to an instance of FileOps. A deep copy
//		of the internal member data values contained in
//		this instance will be returned to the calling
//		function.
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
//	FileOps
//
//		If this method completes successfully without
//		error, this parameter will return a new instance
//		of 'FileOps' containing a deep copy of all the
//		internal member data values encapsulated in the
//		input parameter 'fOps'.
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
func (fOpsNanobot *FileOperationsNanobot) copyOut(
	fOps *FileOps,
	errPrefDto *ePref.ErrPrefixDto) (
	FileOps,
	error) {

	if fOpsNanobot.lock == nil {
		fOpsNanobot.lock = new(sync.Mutex)
	}

	fOpsNanobot.lock.Lock()

	defer fOpsNanobot.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	var newFOps FileOps

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"FileOperationsNanobot."+
			"copyOut()",
		"")

	if err != nil {
		return newFOps, err
	}

	if fOps == nil {
		err = fmt.Errorf("%v\n"+
			"Error: FileOps instance is invalid!\n"+
			"Input parameter 'fOps' is a nil pointer.\n",
			ePrefix.String())

		return newFOps, err
	}

	if !fOps.isInitialized {

		err = fmt.Errorf("%v\n"+
			"Error: FileOps instance is invalid!\n"+
			"Input parameter 'fOps' has NOT been initialized.\n"+
			"fOps.isInitialized = 'false'.",
			ePrefix.String())

		return FileOps{}, err
	}

	newFOps.source,
		err = fOps.source.CopyOut(
		ePrefix.XCpy("newFOps.source<-fOps.source"))

	if err != nil {
		return newFOps, err
	}

	newFOps.destination,
		err = fOps.destination.CopyOut(
		ePrefix.XCpy("newFOps.destination" +
			"<-fOps.destination"))

	if err != nil {
		return newFOps, err
	}

	newFOps.opToExecute = fOps.opToExecute

	newFOps.isInitialized = true

	return newFOps, err
}
