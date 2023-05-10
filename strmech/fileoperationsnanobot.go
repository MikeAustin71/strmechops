package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type FileOperationsNanobot struct {
	lock *sync.Mutex
}

// copyOut
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

// copySrcToDestByIoByHardLink
//
// This method receives an instance of FileOps. This
// structure contains both a source File Manager
// (fOps.source  FileMgr) and a destination File
// Manager (fOps.destination  FileMgr).
//
// This method proceeds to copy the source file to
// the destination location represented by the
// destination File Manager.
//
// Note that if the destination directory does not
// exist, this method will attempt to create it.
//
// The copy operation will be carried out in two
// attempts. The first attempt will try to copy the
// source file to the destination by creating a new
// destination file and copying the source file contents
// to that new destination file using a technique known
// as 'io.Copy'.
//
// If this first file copy operation fails, a second
// attempt will be made using a technique known as a
// 'Hard Link'. This technique will utilize a hard
// symbolic link to the existing source file in order to
// create the destination file.
//
// If both attempted copy operations fail, and error will
// be returned.
//
// ----------------------------------------------------------------
//
// # Reference:
//
// https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fOps 						*FileOps
//
//		A pointer to an instance of FileOps. This
//		structure contains a source File Manager and a
//		destination File Manager. The source file
//		identified by the source File Manager will be
//		copied to the location and file name identified
//		by the destination File Manager.
//
//		The source File Manager is represented by
//		internal member variable 'fOps.source'.
//
//		The destination File Manager is represented by
//		internal member variable 'fOps.destination'.
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
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fOpsNanobot *FileOperationsNanobot) copySrcToDestByIoByHardLink(
	fOps *FileOps,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fOpsNanobot.lock == nil {
		fOpsNanobot.lock = new(sync.Mutex)
	}

	fOpsNanobot.lock.Lock()

	defer fOpsNanobot.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"FileOperationsNanobot."+
			"copySrcToDestByIoByHardLink()",
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
			"Input parameter 'fOps' has NOT been initialized.\n"+
			"fOps.isInitialized = 'false'.",
			ePrefix.String())

		return err
	}

	return fOps.source.
		CopyFileMgrByIoByLink(
			&fOps.destination,
			ePrefix.XCpy(
				"fOps.source<-fOps.destination"))
}
