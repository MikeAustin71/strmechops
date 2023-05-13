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
		return fmt.Errorf("%v\n"+
			"An error occurred while creating the\n"+
			"Source Directory and Source File.\n"+
			"fOps.source Directory= '%v\n"+
			"fOps.source File= '%v'\n"+
			"Error= \n%v\n",
			funcName,
			fOps.source.GetAbsolutePath(),
			fOps.source.GetAbsolutePathFileName(),
			err2.Error())
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
// This method receives an instance of FileOps. This
// structure contains a source File Manager
// (fOps.source  FileMgr) identifying the source file
// which be created by this method.
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

	err2 = fOps.source.CreateThisFile(ePrefix)

	if err != nil {

		err = fmt.Errorf("%v\n"+
			"An error occurred while creating the\n"+
			"source file.\n"+
			"fOps.source File = '%v'\n"+
			"Error= \n%v\n",
			funcName,
			fOps.source.GetAbsolutePathFileName(),
			err2.Error())
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
