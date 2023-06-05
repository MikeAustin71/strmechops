package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type FileOperationsNanobot struct {
	lock *sync.Mutex
}

// copyIn
//
// This method receives two instances of FileOps, the
// Destination FileOps and the Source FileOps.
//
// All data contained in the Source FileOps instance
// will be copied to corresponding data elements in
// the Destination FileOps instance.
//
// The type of copy operation performed is a 'Deep'
// copy.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	destinationFOps				*FileOps
//
//		A pointer to an instance of FileOps. A deep copy
//		of the internal member data values contained in
//		'sourceFOps' will be copied to the corresponding
//		data elements contained this parameter,
//		'destinationFOps'.
//
//		If 'fOps' has not been properly initialized,
//		an error will be returned.
//
//	sourceFOps					*FileOps
//
//		A pointer to an instance of FileOps. A deep copy
//		of the internal member data values contained in
//		this instance will be copied to the corresponding
//		data elements contained parameter
//		'destinationFOps'.
//
//		If 'sourceFOps' evaluates as invalid, an error
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
func (fOpsNanobot *FileOperationsNanobot) copyIn(
	destinationFOps *FileOps,
	sourceFOps *FileOps,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fOpsNanobot.lock == nil {
		fOpsNanobot.lock = new(sync.Mutex)
	}

	fOpsNanobot.lock.Lock()

	defer fOpsNanobot.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName :=
		"FileOperationsNanobot." +
			"copyIn()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return err
	}

	if destinationFOps == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Target FileOps instance is invalid!\n"+
			"Input parameter 'destinationFOps' is a nil pointer.\n",
			ePrefix.String())

		return err
	}

	if sourceFOps == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Source FileOps instance is invalid!\n"+
			"Input parameter 'sourceFOps' is a nil pointer.\n",
			ePrefix.String())

		return err
	}

	if !sourceFOps.isInitialized {

		err = fmt.Errorf("%v\n"+
			"Error: Source FileOps instance is invalid!\n"+
			"Input parameter 'sourceFOps' has NOT been initialized.\n"+
			"sourceFOps.isInitialized = 'false'.",
			ePrefix.String())

		return err
	}

	fHelperAtom := fileMgrHelperAtom{}

	err = fHelperAtom.
		isFileMgrValid(
			&sourceFOps.source,
			ePrefix.XCpy("sourceFOps.source"))

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Error: The Source FileOps source File Manager\n"+
			"'sourceFOps.source' is Invalid!\n"+
			"Error= \n%v\n",
			funcName,
			err.Error())
	}

	err = fHelperAtom.
		isFileMgrValid(
			&sourceFOps.destination,
			ePrefix.XCpy("sourceFOps.source"))

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Error: The Source FileOps destination File Manager\n"+
			"'sourceFOps.destination' is Invalid!\n"+
			"Error= \n%v\n",
			funcName,
			err.Error())
	}

	err = sourceFOps.opToExecute.IsValid()

	if err != nil {

		sourceFOps.opToExecute = FileOpCode.None()

	}

	err = new(fileMgrHelperAtom).copyInFileMgr(
		&destinationFOps.source,
		&sourceFOps.source,
		ePrefix.XCpy(
			"destinationFOps.source<-"+
				"sourceFOps.source"))

	if err != nil {

		return fmt.Errorf("%v\n"+
			"The Source FileOps copy operation FAILED!\n"+
			"Error: destinationFOps.source <- sourceFOps.source\n"+
			"sourceFOps.source= '%v'\n"+
			"Error= \n%v\n",
			funcName,
			sourceFOps.source.absolutePathFileName,
			err.Error())
	}

	err = new(fileMgrHelperAtom).copyInFileMgr(
		&destinationFOps.destination,
		&sourceFOps.destination,
		ePrefix.XCpy(
			"destinationFOps.destination<-"+
				"sourceFOps.destination"))

	if err != nil {

		return fmt.Errorf("%v\n"+
			"The Source FileOps copy operation FAILED!\n"+
			"Error: destinationFOps.destination <- sourceFOps.destination\n"+
			"sourceFOps.destination= '%v'\n"+
			"Error= \n%v\n",
			funcName,
			sourceFOps.destination.absolutePathFileName,
			err.Error())
	}

	destinationFOps.opToExecute =
		sourceFOps.opToExecute

	return err
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

// createDestDirectory
//
// Creates the destination directory using information
// from the FileOps destination file manager. The FileOps
// object is passed as input parameter 'fOps'.
//
// The destination directory is specified in the internal
// member File Manager object, 'fOps.destination'.
//
// If the source directory already exists, no action will
// be taken and this method will exit without error.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fOps 						*FileOps
//
//		A pointer to an instance of FileOps. This
//		structure contains a destination File Manager.
//		The destination file directory path identified by
//		the destination File Manager will be used to
//		create the destination directory.
//
//		The destination File Manager is represented by
//		internal member variable 'fOps.destination'.
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
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fOpsNanobot *FileOperationsNanobot) createDestDirectory(
	fOps *FileOps,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fOpsNanobot.lock == nil {
		fOpsNanobot.lock = new(sync.Mutex)
	}

	fOpsNanobot.lock.Lock()

	defer fOpsNanobot.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName :=
		"FileOperationsNanobot.createDestDirectory()"

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
			"Input parameter 'fOps' has NOT been initialized.\n"+
			"fOps.isInitialized = 'false'.",
			ePrefix.String())

		return err
	}

	var err2 error

	err2 = fOps.destination.CreateDir(
		ePrefix.XCpy(
			"fOps.destination"))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"An error occurred while creating the\n"+
			"destination file directory.\n"+
			"fOps.destination File Directory = '%v'\n"+
			"Error= \n%v\n",
			funcName,
			fOps.destination.GetAbsolutePath(),
			err2.Error())
	}

	return err
}

// createDestDirectoryAndFile
//
// Creates the destination directory and file using
// information from a FileOps object passed as input
// parameter 'fOps'.
//
// The FileOps structure contains a destination File
// Manager (fOps.destination  FileMgr) identifying the
// destination directory and destination file which will
// be created by this method.
//
// The newly created destination file will be empty and
// contain zero (0) bytes.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fOps 						*FileOps
//
//		A pointer to an instance of FileOps. This
//		structure contains a source File Manager. The
//		source directory and file identified by this
//		source File Manager will be created. The source
//		file will be created as an empty file with zero
//		(0) bytes.
//
//		The source File Manager is represented by
//		internal member variable 'fOps.source'.
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
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fOpsNanobot *FileOperationsNanobot) createDestDirectoryAndFile(
	fOps *FileOps,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fOpsNanobot.lock == nil {
		fOpsNanobot.lock = new(sync.Mutex)
	}

	fOpsNanobot.lock.Lock()

	defer fOpsNanobot.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName :=
		"FileOperationsNanobot.createDestDirectoryAndFile()"

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
			"Input parameter 'fOps' has NOT been initialized.\n"+
			"fOps.isInitialized = 'false'.",
			ePrefix.String())

		return err
	}

	var err2 error

	err2 = fOps.destination.CreateDirAndFile(
		ePrefix.XCpy(
			"fOps.destination"))

	if err2 != nil {

		return fmt.Errorf("%v\n"+
			"An error occurred while creating the\n"+
			"Destination Directory and Destination File.\n"+
			"fOps.destination Directory= '%v\n"+
			"fOps.destination File= '%v'\n"+
			"Error= \n%v\n",
			funcName,
			fOps.destination.GetAbsolutePath(),
			fOps.destination.GetAbsolutePathFileName(),
			err2.Error())
	}

	err2 = fOps.destination.CloseThisFile(ePrefix.XCpy(
		"fOps.destination"))

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"An error occurred while closing the\n"+
			"Destination File.\n"+
			"fOps.destination File= '%v'\n"+
			"Error= \n%v\n",
			funcName,
			fOps.destination.GetAbsolutePathFileName(),
			err2.Error())
	}

	return err
}

// createDestFile
//
// Creates the destination file using information from a
// FileOps object passed as input parameter 'fOps'.
//
// The FileOps structure contains a destination File
// Manager (fOps.destination  FileMgr) identifying the
// destination file which will be created by this method.
//
// The newly created destination file will be empty and
// contain zero (0) bytes.
//
// If the destination directory does NOT exist, an error
// will be returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fOps 						*FileOps
//
//		A pointer to an instance of FileOps. This
//		structure contains a destination File Manager.
//		The destination directory and file identified by
//	 	this destination File Manager will be created.
//	 	The destination file will be created as an empty
//	 	file with zero (0) bytes.
//
//		The destination File Manager is represented by
//		internal member variable 'fOps.destination'.
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
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fOpsNanobot *FileOperationsNanobot) createDestFile(
	fOps *FileOps,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fOpsNanobot.lock == nil {
		fOpsNanobot.lock = new(sync.Mutex)
	}

	fOpsNanobot.lock.Lock()

	defer fOpsNanobot.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName :=
		"FileOperationsNanobot.createDestFile()"

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
			"Input parameter 'fOps' has NOT been initialized.\n"+
			"fOps.isInitialized = 'false'.",
			ePrefix.String())

		return err
	}

	var err2 error

	err2 = fOps.destination.CreateThisFile(
		ePrefix.XCpy(
			"fOps.destination"))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"An error occurred while creating the\n"+
			"Destination file.\n"+
			"fOps.destination File = '%v'\n"+
			"Error= \n%v\n",
			funcName,
			fOps.destination.GetAbsolutePathFileName(),
			err2.Error())

		return err
	}

	err2 = fOps.destination.CloseThisFile(ePrefix.XCpy(
		"fOps.destination"))

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"An error occurred while closing the\n"+
			"Destination File.\n"+
			"fOps.destination File= '%v'\n"+
			"Error= \n%v\n",
			funcName,
			fOps.destination.GetAbsolutePathFileName(),
			err2.Error())
	}

	return err
}

// createSrcDirectory
//
// This method receives an instance of FileOps and then
// proceeds to create the source directory using the path
// information taken the FileOps source file manager
// (fOps.source  FileMgr).
//
// If the source directory already exists, no action will
// be taken and this method will exit without error.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fOps 						*FileOps
//
//		A pointer to an instance of FileOps. This
//		structure contains a source File Manager. The
//		source file directory path identified by the
//		source File Manager will be used to create
//		the source directory.
//
//		The source File Manager is represented by
//		internal member variable 'fOps.source'.
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
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fOpsNanobot *FileOperationsNanobot) createSrcDirectory(
	fOps *FileOps,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fOpsNanobot.lock == nil {
		fOpsNanobot.lock = new(sync.Mutex)
	}

	fOpsNanobot.lock.Lock()

	defer fOpsNanobot.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "FileOperationsNanobot.createSrcDirectory()"

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
			"Input parameter 'fOps' has NOT been initialized.\n"+
			"fOps.isInitialized = 'false'.",
			ePrefix.String())

		return err
	}

	var err2 error

	err2 = fOps.source.CreateDir(
		ePrefix.XCpy(
			"fOps.source"))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"An error occurred while creating the\n"+
			"source file directory.\n"+
			"fOps.source File Directory = '%v'\n"+
			"Error= \n%v\n",
			funcName,
			fOps.source.GetAbsolutePath(),
			err2.Error())
	}

	return err
}

// createSrcDirectoryAndFile
//
// Creates the source directory and file using
// information from a FileOps object passed as input
// parameter 'fOps'.
//
// The FileOps structure contains a source File Manager
// (fOps.source  FileMgr) identifying the source
// directory and source file which will be created by
// this method.
//
// The newly created source file will be empty and
// contain zero (0) bytes.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fOps 						*FileOps
//
//		A pointer to an instance of FileOps. This
//		structure contains a source File Manager. The
//		source directory and file identified by this
//		source File Manager will be created. The source
//		file will be created as an empty file with zero
//		(0) bytes.
//
//		The source File Manager is represented by
//		internal member variable 'fOps.source'.
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
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fOpsNanobot *FileOperationsNanobot) createSrcDirectoryAndFile(
	fOps *FileOps,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fOpsNanobot.lock == nil {
		fOpsNanobot.lock = new(sync.Mutex)
	}

	fOpsNanobot.lock.Lock()

	defer fOpsNanobot.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName :=
		"FileOperationsNanobot.createSrcDirectoryAndFile()"

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
			"Input parameter 'fOps' has NOT been initialized.\n"+
			"fOps.isInitialized = 'false'.",
			ePrefix.String())

		return err
	}

	var err2 error

	err2 = fOps.source.CreateDirAndFile(
		ePrefix.XCpy(
			"fOps.source"))

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"An error occurred while creating the\n"+
			"Source Directory and Source File.\n"+
			"fOps.source Directory= '%v\n"+
			"fOps.source File= '%v'\n"+
			"Error= \n%v\n",
			funcName,
			fOps.source.GetAbsolutePath(),
			fOps.source.GetAbsolutePathFileName(),
			err2.Error())

		return err
	}

	err2 = fOps.source.CloseThisFile(ePrefix.XCpy(
		"fOps.source"))

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"An error occurred while closing the\n"+
			"Source File.\n"+
			"fOps.source File= '%v'\n"+
			"Error= \n%v\n",
			funcName,
			fOps.source.GetAbsolutePathFileName(),
			err2.Error())
	}

	return err
}

// createSrcFile
//
// Creates the source file using information from a
// FileOps object passed as input parameter 'fOps'.
//
// The FileOps structure contains a source File
// Manager (fOps.source  FileMgr) identifying the
// source file which will be created by this method.
//
// The newly created source file will be empty and
// contain zero (0) bytes.
//
// If the source directory does NOT exist, an error
// will be returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fOps 						*FileOps
//
//		A pointer to an instance of FileOps. This
//		structure contains a source File Manager. The
//		source directory path and source file identified
//		by the source File Manager will be used to create
//		the source file.
//
//		The source File Manager is represented by
//		internal member variable 'fOps.source'.
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
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fOpsNanobot *FileOperationsNanobot) createSrcFile(
	fOps *FileOps,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fOpsNanobot.lock == nil {
		fOpsNanobot.lock = new(sync.Mutex)
	}

	fOpsNanobot.lock.Lock()

	defer fOpsNanobot.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName :=
		"FileOperationsNanobot.createSrcFile()"

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
			"Input parameter 'fOps' has NOT been initialized.\n"+
			"fOps.isInitialized = 'false'.",
			ePrefix.String())

		return err
	}

	var err2 error

	err2 = fOps.source.CreateThisFile(
		ePrefix.XCpy(
			"fOps.source"))

	if err != nil {

		err = fmt.Errorf("%v\n"+
			"An error occurred while creating the\n"+
			"source file.\n"+
			"fOps.source File = '%v'\n"+
			"Error= \n%v\n",
			funcName,
			fOps.source.GetAbsolutePathFileName(),
			err2.Error())

		return err
	}

	err2 = fOps.source.CloseThisFile(ePrefix.XCpy(
		"fOps.source"))

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"An error occurred while closing the\n"+
			"Source File.\n"+
			"fOps.source File= '%v'\n"+
			"Error= \n%v\n",
			funcName,
			fOps.source.GetAbsolutePathFileName(),
			err2.Error())
	}

	return err
}

// copySrcToDestByHardLink
//
// This method receives an instance of FileOps. This
// structure contains both a source File Manager
// (fOps.source  FileMgr) and a destination File
// Manager (fOps.destination  FileMgr) identifying
// both source and destination files.
//
// This method proceeds to copy the source file to
// the destination file location represented by the
// destination File Manager. The copy operation will be
// made using a technique known as a 'Hard Link'. This
// technique will utilize a hard symbolic link to the
// existing source file in order to create the
// destination file.
//
// If the copy operation fails, an error will be
// returned.
//
// Note that if the destination directory does not
// exist, this method will attempt to create it.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//	https://golang.org/pkg/os/#Link
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
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fOpsNanobot *FileOperationsNanobot) copySrcToDestByHardLink(
	fOps *FileOps,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fOpsNanobot.lock == nil {
		fOpsNanobot.lock = new(sync.Mutex)
	}

	fOpsNanobot.lock.Lock()

	defer fOpsNanobot.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName :=
		"FileOperationsNanobot.copySrcToDestByHardLink()"

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
			"Input parameter 'fOps' has NOT been initialized.\n"+
			"fOps.isInitialized = 'false'.",
			ePrefix.String())

		return err
	}

	var err2 error

	err2 = fOps.source.CopyFileMgrByLink(
		&fOps.destination,
		ePrefix)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"An error occurred while copying the\n"+
			"source file to the destination file.\n"+
			"fOps.source File = '%v'\n"+
			"fOps.destination File = '%v'\n"+
			"Error= \n%v\n",
			funcName,
			fOps.source.GetAbsolutePathFileName(),
			fOps.destination.GetAbsolutePathFileName(),
			err2.Error())
	}

	return err
}

// copySrcToDestByHardLinkByIo
//
// This method receives an instance of FileOps. This
// structure contains both a source File Manager
// (fOps.source  FileMgr) and a destination File
// Manager (fOps.destination  FileMgr) identifying
// both source and destination files.
//
// This method proceeds to copy the source file to
// the destination file location represented by the
// destination File Manager.
//
// Note that if the destination directory does not
// exist, this method will attempt to create it.
//
// The copy operation will be carried out in two
// attempts.
//
// If this first attempted file copy operation will be
// made using a technique known as a 'Hard Link'. This
// technique will utilize a hard symbolic link to the
// existing source file in order to create the
// destination file.
//
// If the first attempted copy operation fails, a second
// attempt will try to copy the source file to the
// destination by creating a new destination file and
// copying the source file contents to that new
// destination file using a technique known as 'io.Copy'.
// This technique involves copying the source file to the
// destination by creating a new destination file before
// copying the source file contents to that new
// destination file.
//
// If both attempted copy operations fail, and error will
// be returned.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//	https://golang.org/pkg/os/#Link
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
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fOpsNanobot *FileOperationsNanobot) copySrcToDestByHardLinkByIo(
	fOps *FileOps,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fOpsNanobot.lock == nil {
		fOpsNanobot.lock = new(sync.Mutex)
	}

	fOpsNanobot.lock.Lock()

	defer fOpsNanobot.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "FileOperationsNanobot." +
		"copySrcToDestByHardLinkByIo()"

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
			"Input parameter 'fOps' has NOT been initialized.\n"+
			"fOps.isInitialized = 'false'.",
			ePrefix.String())

		return err
	}

	var err2 error

	err2 = fOps.source.CopyFileMgrByLinkByIo(
		&fOps.destination,
		ePrefix)

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"An error occurred while copying the\n"+
			"source file to the destination file.\n"+
			"fOps.source File = '%v'\n"+
			"fOps.destination File = '%v'\n"+
			"Error= \n%v\n",
			funcName,
			fOps.source.GetAbsolutePathFileName(),
			fOps.destination.GetAbsolutePathFileName(),
			err2.Error())
	}

	return err
}

// copySrcToDestByIo
//
// This method receives an instance of FileOps. This
// structure contains both a source File Manager
// (fOps.source  FileMgr) and a destination File
// Manager (fOps.destination  FileMgr) identifying
// both source and destination files.
//
// This method proceeds to copy the source file to
// the destination location represented by the
// destination File Manager.
//
// The copy operation will be accomplished using a
// technique known as 'io.Copy'. This technique
// involves copying the source file to the destination by
// creating a new destination file before copying the
// source file contents to that new destination file.
//
// Note that if the destination directory does not
// exist, this method will attempt to create it.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//	https://golang.org/pkg/os/#Link
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
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fOpsNanobot *FileOperationsNanobot) copySrcToDestByIo(
	fOps *FileOps,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fOpsNanobot.lock == nil {
		fOpsNanobot.lock = new(sync.Mutex)
	}

	fOpsNanobot.lock.Lock()

	defer fOpsNanobot.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "FileOperationsNanobot." +
		"copySrcToDestByIo()"

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
			"Input parameter 'fOps' has NOT been initialized.\n"+
			"fOps.isInitialized = 'false'.",
			ePrefix.String())

		return err
	}

	var err2 error

	err2 = fOps.source.CopyFileMgrByIo(
		&fOps.destination,
		ePrefix)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"An error occurred while copying the\n"+
			"source file to the destination file.\n"+
			"fOps.source File = '%v'\n"+
			"fOps.destination File = '%v'\n"+
			"Error= \n%v\n",
			funcName,
			fOps.source.GetAbsolutePathFileName(),
			fOps.destination.GetAbsolutePathFileName(),
			err2.Error())
	}

	return err
}

// copySrcToDestByIoByHardLink
//
// This method receives an instance of FileOps. This
// structure contains both a source File Manager
// (fOps.source  FileMgr) and a destination File
// Manager (fOps.destination  FileMgr) identifying
// both source and destination files.
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
//	https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//	https://golang.org/pkg/os/#Link
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

	funcName := "FileOperationsNanobot." +
		"copySrcToDestByIoByHardLink()"

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
			"Input parameter 'fOps' has NOT been initialized.\n"+
			"fOps.isInitialized = 'false'.",
			ePrefix.String())

		return err
	}

	var err2 error

	err2 = fOps.source.
		CopyFileMgrByIoByLink(
			&fOps.destination,
			ePrefix)

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"An error occurred while copying the\n"+
			"source file to the destination file.\n"+
			"fOps.source File = '%v'\n"+
			"fOps.destination File = '%v'\n"+
			"Error= \n%v\n",
			funcName,
			fOps.source.GetAbsolutePathFileName(),
			fOps.destination.GetAbsolutePathFileName(),
			err2.Error())
	}

	return err
}

// moveSourceFileFileToDestinationDir
//
// Moves a source file to the destination directory.
// The source file and destination file names are
// specified by the FileOps object passed as input
// parameter 'fOps'.
//
// The FileOps structure contains a destination File
// Manager (fOps.destination FileMgr) identifying the
// destination file which will be created by this method.
//
// This same FileOps structure also contains a source
// File Manager which specifies the source directory and
// file name (fOps.source FileMgr).
//
// If the destination directory does NOT exist, this
// method will attempt to create that directory.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This is a 'move' file operation. Therefore, by
//	definition, the original source file in the original
//	source directory will be deleted after it is moved to
//	the destination directory.
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
//		moved to the location and file name identified
//		by the destination File Manager. After the 'move'
//		operation is completed, the source file will be
//		deleted.
//
//		The source File Manager is represented by
//		internal member variable 'fOps.source'.
//
//		The destination File Manager is represented by
//		internal member variable 'fOps.destination'.
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
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fOpsNanobot *FileOperationsNanobot) moveSourceFileToDestinationDir(
	fOps *FileOps,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fOpsNanobot.lock == nil {
		fOpsNanobot.lock = new(sync.Mutex)
	}

	fOpsNanobot.lock.Lock()

	defer fOpsNanobot.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName :=
		"FileOperationsNanobot." +
			"moveSourceFileToDestinationDir()"

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
			"Input parameter 'fOps' has NOT been initialized.\n"+
			"fOps.isInitialized = 'false'.",
			ePrefix.String())

		return err
	}

	var deepCopyDirMgr DirMgr

	var err2 error

	deepCopyDirMgr,
		err2 = fOps.destination.GetDirMgr(
		ePrefix)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"An error occurred while retrieving the\n"+
			"Destination Directory Manager.\n"+
			"fOps.destination Directory = '%v'\n"+
			"Error= \n%v\n",
			funcName,
			fOps.destination.GetAbsolutePath(),
			err2.Error())

		return err
	}

	_,
		err2 = fOps.source.
		MoveFileToNewDirMgr(
			deepCopyDirMgr,
			ePrefix.XCpy(
				"fOps.source"))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"An error occurred while moving the\n"+
			"Source file to the Destination Directory.\n"+
			"fOps.source File Name = '%v'\n"+
			"fOps.destination Directory = '%v'\n"+
			"Error= \n%v\n",
			funcName,
			fOps.source.GetAbsolutePathFileName(),
			deepCopyDirMgr.GetAbsolutePath(),
			err2.Error())
	}

	return err
}

// moveSourceFileToDestinationFile
//
// Moves a source file to the destination file by first
// copying the source file to the destination and then
// deleting the source file. The source file and
// destination file names are specified by the FileOps
// object passed as input parameter 'fOps'.
//
// The final file name consist of the destination file
// name in the destination directory.
//
// The FileOps structure contains a source File Manager
// which specifies the source directory and source file
// name (fOps.source FileMgr).
//
// The FileOps structure also contains a destination File
// Manager (fOps.destination FileMgr) identifying the
// destination file which will be created by this method.
//
// The contents of the source file will be copied to the
// destination file name. After the 'move' operation is
// completed, the source file will be deleted.
//
// If the destination directory does NOT exist, an error
// will be returned.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This is a 'move' file operation. Therefore, by
//	definition, the original source file in the original
//	source directory will be deleted after it is moved to
//	the destination directory and destination file name.
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
//		moved to the location and file name identified
//		by the destination File Manager. After the 'move'
//		operation is completed, the source file will be
//		deleted.
//
//		The source File Manager is represented by
//		internal member variable 'fOps.source'.
//
//		The destination File Manager is represented by
//		internal member variable 'fOps.destination'.
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
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fOpsNanobot *FileOperationsNanobot) moveSourceFileToDestinationFile(
	fOps *FileOps,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fOpsNanobot.lock == nil {
		fOpsNanobot.lock = new(sync.Mutex)
	}

	fOpsNanobot.lock.Lock()

	defer fOpsNanobot.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName :=
		"FileOperationsNanobot." +
			"moveSourceFileToDestinationFile()"

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
			"Input parameter 'fOps' has NOT been initialized.\n"+
			"fOps.isInitialized = 'false'.",
			ePrefix.String())

		return err
	}

	var err2 error

	err2 = fOps.source.MoveFileToFileMgr(
		fOps.destination,
		ePrefix.XCpy(
			"fOps.destination<-fOps.source"))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"An error occurred while moving the\n"+
			"Source File to the Destination Directory.\n"+
			"fOps.destination Directory = '%v'\n"+
			"fOps.source File = '%v'\n"+
			"Error= \n%v\n",
			funcName,
			fOps.destination.GetAbsolutePath(),
			fOps.source.GetAbsolutePathFileName(),
			err2.Error())
	}

	return err
}

// setFileOpsByDirAndFileNameStr
//
// Reconfigures an instance of FileOps using source and
// destination file names passed as directory and file
// name component strings.
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
//		A pointer to an instance of FileOps. This
//		instance will be reconfigured using information
//		provided by the source and destination parameters
//		listed below.
//
//	sourceDirStr				string
//
//		This string specifies the source file directory
//		which will be combined with the source file name
//		parameter to configure the source file member
//		variable 'fOps.source' contained in the 'fOps'
//		FileOps	object.
//
//		If this parameter is submitted as an empty string
//		with a zero (0) string length, an error will be
//		returned.
//
//	sourceFileNameExtStr		string
//
//		This string specifies the source file name and
//		file extension which will be combined with the
//		source directory parameter to configure the
//		source file member variable 'fOps.source'
//		contained in the 'fOps' FileOps object.
//
//		If this parameter is submitted as an empty string
//		with a zero (0) string length, an error will be
//		returned.
//
//	destinationDirStr			string
//
//		This string specifies the destination file
//		directory which will be combined with the
//		destination file name parameter to configure the
//		destination file member	variable 'fOps.destination'
//		contained in the 'fOps' FileOps object.
//
//		If this parameter is submitted as an empty string
//		with a zero (0) string length, an error will be
//		returned.
//
//	destinationFileNameExtStr	string
//
//		This string specifies the destination file name
//		and file extension which will be combined with
//		the destination directory parameter to configure
//		the destination file member variable
//		'fOps.destination' contained in the 'fOps'
//		FileOps object.
//
//		If this parameter is submitted as an empty string
//		with a zero (0) string length, an error will be
//		returned.
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
func (fOpsNanobot *FileOperationsNanobot) setFileOpsByDirAndFileNameStr(
	fOps *FileOps,
	sourceDirStr string,
	sourceFileNameExtStr string,
	destinationDirStr string,
	destinationFileNameExtStr string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fOpsNanobot.lock == nil {
		fOpsNanobot.lock = new(sync.Mutex)
	}

	fOpsNanobot.lock.Lock()

	defer fOpsNanobot.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName :=
		"FileOperationsNanobot." +
			"setFileOpsByDirAndFileNameStr()"

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

	if len(sourceDirStr) == 0 {

		return fmt.Errorf("%v\n"+
			"Error: 'sourceDirStr' is an EMPTY STRING!\n",
			ePrefix.String())
	}

	if len(sourceFileNameExtStr) == 0 {

		return fmt.Errorf("%v\n"+
			"Error: 'sourceFileNameExtStr' is an EMPTY STRING!\n",
			ePrefix.String())
	}

	if len(destinationDirStr) == 0 {

		return fmt.Errorf("%v\n"+
			"Error: 'sourceFileNameExtStr' is an EMPTY STRING!\n",
			ePrefix.String())
	}

	if len(destinationFileNameExtStr) == 0 {

		return fmt.Errorf("%v\n"+
			"Error: 'sourceFileNameExtStr' is an EMPTY STRING!\n",
			ePrefix.String())
	}

	var sourceFMgr, destinationFMgr FileMgr

	fMgrObj := new(FileMgr)

	sourceFMgr,
		err = fMgrObj.
		NewFromDirStrFileNameStr(
			sourceDirStr,
			sourceFileNameExtStr,
			ePrefix.XCpy("sourceFMgr<-"))

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Creation of intermediate source File Manager Failed!\n"+
			"sourceDirStr = '%v'\n"+
			"sourceFileNameExtStr = '%v'\n"+
			"Error= \n%v\n",
			funcName,
			sourceDirStr,
			sourceFileNameExtStr,
			err.Error())
	}

	destinationFMgr,
		err = fMgrObj.
		NewFromDirStrFileNameStr(
			destinationDirStr,
			destinationFileNameExtStr,
			ePrefix.XCpy("destinationFMgr<-"))

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Creation of intermediate destination File Manager Failed!\n"+
			"destinationDirStr = '%v'\n"+
			"destinationFileNameExtStr = '%v'\n"+
			"Error= \n%v\n",
			funcName,
			destinationDirStr,
			destinationFileNameExtStr,
			err.Error())
	}

	return new(FileOperationsAtom).setFileOps(
		fOps,
		sourceFMgr,
		destinationFMgr,
		ePrefix)
}

// setFileOpsByDirMgrFileName
//
// Reconfigures an instance of FileOps using source and
// destination file names passed as directory managers
// and file name string components.
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
//		A pointer to an instance of FileOps. This
//		instance will be reconfigured using information
//		provided by the source and destination parameters
//		listed below.
//
//	sourceDir					DirMgr
//
//		This instance of Directory Manager specifies the
//		source file directory which will be combined with
//		the source file name parameter to configure the
//		source file member variable 'fOps.source'
//		contained in the 'fOps' FileOps object.
//
//		If this parameter is evaluated as invalid, an
//		error will be returned.
//
//	sourceFileNameExt			string
//
//		This string specifies the source file name and
//		file extension which will be combined with the
//		source directory parameter to configure the
//		source file member variable 'fOps.source'
//		contained in the 'fOps' FileOps object.
//
//		If this parameter is submitted as an empty string
//		with a zero (0) string length, an error will be
//		returned.
//
//	destinationDir				DirMgr
//
//		This instance of Directory Manager specifies the
//		destination file directory which will be combined
//		with the destination file name parameter to
//		configure the destination file member variable
//		'fOps.destination' contained in the 'fOps'
//		FileOps object.
//
//		If this parameter is evaluated as invalid, an
//		error will be returned.
//
//	destinationFileNameExtStr	string
//
//		This string specifies the destination file name
//		and file extension which will be combined with
//		the destination directory parameter to configure
//		the destination file member variable
//		'fOps.destination' contained in the 'fOps'
//		FileOps object.
//
//		If this parameter is submitted as an empty string
//		with a zero (0) string length, an error will be
//		returned.
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
func (fOpsNanobot *FileOperationsNanobot) setFileOpsByDirMgrFileName(
	fOps *FileOps,
	sourceDir DirMgr,
	sourceFileNameExt string,
	destinationDir DirMgr,
	destinationFileNameExt string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fOpsNanobot.lock == nil {
		fOpsNanobot.lock = new(sync.Mutex)
	}

	fOpsNanobot.lock.Lock()

	defer fOpsNanobot.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName :=
		"FileOperationsNanobot." +
			"setFileOpsByDirMgrFileName()"

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

	dMgrHelperPreon := new(dirMgrHelperPreon)

	_,
		_,
		err = dMgrHelperPreon.validateDirMgr(
		&sourceDir,
		false, // sourceDir NOT required to exist on disk
		"sourceDir",
		ePrefix.XCpy("sourceDir"))

	if err != nil {
		return err
	}

	if len(sourceFileNameExt) == 0 {

		return fmt.Errorf("%v\n"+
			"Error: 'sourceFileNameExtStr' is an EMPTY STRING!\n",
			ePrefix.String())
	}

	_,
		_,
		err = dMgrHelperPreon.validateDirMgr(
		&destinationDir,
		false, // sourceDir NOT required to exist on disk
		"destinationDir",
		ePrefix.XCpy("destinationDir"))

	if err != nil {
		return err
	}

	if len(destinationFileNameExt) == 0 {

		return fmt.Errorf("%v\n"+
			"Error: 'destinationFileNameExtStr' is an EMPTY STRING!\n",
			ePrefix.String())
	}

	var sourceFMgr, destinationFMgr FileMgr

	sourceFMgr,
		err = new(FileMgr).
		NewFromDirMgrFileNameExt(
			sourceDir,
			sourceFileNameExt,
			ePrefix)

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Creation of intermediate source File Manager Failed!\n"+
			"sourceDir= '%v'\n"+
			"sourceFileNameExt= '%v'\n"+
			"Error= \n%v\n",
			funcName,
			sourceDir.GetAbsolutePath(),
			sourceFileNameExt,
			err.Error())
	}

	destinationFMgr,
		err = new(FileMgr).
		NewFromDirMgrFileNameExt(
			destinationDir,
			destinationFileNameExt,
			ePrefix)

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Creation of intermediate destination File Manager Failed!\n"+
			"destinationDir= '%v'\n"+
			"destinationFileNameExt= '%v'\n"+
			"Error= \n%v\n",
			funcName,
			destinationDir.GetAbsolutePath(),
			destinationFileNameExt,
			err.Error())
	}

	return new(FileOperationsAtom).setFileOps(
		fOps,
		sourceFMgr,
		destinationFMgr,
		ePrefix)
}

// setByPathFileNameExtStrs
//
// Reconfigures an instance of FileOps using source and destination
// file names passed as strings via input parameters
// 'sourcePathFileName' and 'destinationPathFileName'.
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
//		input parameters 'sourcePathFileName' and
//		'destinationPathFileName'.
//
//	sourcePathFileNameExt		string
//
//		This string contains the source path and file
//		name used to configure the source File Manager
//		encapsulated in 'fOps' (fOps.source).
//
//		If this parameter is submitted as an empty string
//		with a zero (0) string length, an error will be
//		returned.
//
//	destinationPathFileNameExt	string
//
//		This string contains the destination path and
//		file name used to configure the destination File
//		Manager encapsulated in 'fOps'
//		(fOps.destination).
//
//		If this parameter is submitted as an empty string
//		with a zero (0) string length, an error will be
//		returned.
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
func (fOpsNanobot *FileOperationsNanobot) setByPathFileNameExtStrs(
	fOps *FileOps,
	sourcePathFileNameExt string,
	destinationPathFileNameExt string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fOpsNanobot.lock == nil {
		fOpsNanobot.lock = new(sync.Mutex)
	}

	fOpsNanobot.lock.Lock()

	defer fOpsNanobot.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName :=
		"FileOperationsNanobot." +
			"setByPathFileNameExtStrs()"

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

	if len(sourcePathFileNameExt) == 0 {

		return fmt.Errorf("%v\n"+
			"Error: 'sourceFileNameExtStr' is an EMPTY STRING!\n",
			ePrefix.String())
	}

	if len(destinationPathFileNameExt) == 0 {

		return fmt.Errorf("%v\n"+
			"Error: 'destinationPathFileNameExt' is an EMPTY STRING!\n",
			ePrefix.String())
	}

	var sourceFMgr, destinationFMgr FileMgr

	sourceFMgr,
		err = new(FileMgr).New(
		sourcePathFileNameExt,
		ePrefix.XCpy(
			"sourceFMgr<-sourcePathFileNameExt"))

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Creation of intermediate source File Manager Failed!\n"+
			"sourcePathFileNameExt= '%v'\n"+
			"Error= \n%v\n",
			funcName,
			sourcePathFileNameExt,
			err.Error())
	}

	destinationFMgr,
		err = new(FileMgr).New(
		destinationPathFileNameExt,
		ePrefix.XCpy(
			"destinationFMgr<-destinationPathFileNameExt"))

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Creation of intermediate destination File Manager Failed!\n"+
			"destinationPathFileNameExt= '%v'\n"+
			"Error= \n%v\n",
			funcName,
			destinationPathFileNameExt,
			err.Error())
	}

	return new(FileOperationsAtom).setFileOps(
		fOps,
		sourceFMgr,
		destinationFMgr,
		ePrefix)
}
