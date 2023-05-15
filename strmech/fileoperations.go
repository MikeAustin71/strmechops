package strmech

import (
	"errors"
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// FileOps
//
// This type is used to manage and coordinate various
// operations performed on files.
//
// The FileOps type contains two member variables
// designed to identify source files and destination
// files.
//
// The source file is specified by an instance of File
// Manager configured as a private member data variable
// named 'FileOps.sources'
//
// The destination file is specified by a second instance
// of File Manager configured as a private member data
// variable named 'FileOps.destination'.
//
// The FileOps structure also includes an instance of
// FileOperationCode used to define File Operations
// performed on the source and destination files.
//
// # Creating New Instances Of FileOps
//
// In order to ensure that new instances of FileOps are
// properly initialized, users must call one of the
// following methods:
//
//	new(FileOps).NewByFileMgrs(...)
//	new(FileOps).NewByDirMgrFileName(...)
//	new(FileOps).NewByDirStrsAndFileNameExtStrs(...)
//	new(FileOps).NewByPathFileNameExtStrs(...)
//	new(FileOps).SetByFileMgrs(...)
//	new(FileOps).SetByDirMgrFileName(...)
type FileOps struct {
	isInitialized bool
	source        FileMgr
	destination   FileMgr
	opToExecute   FileOperationCode

	lock *sync.Mutex
}

// CopyOut
//
// Returns a deep copy of the current FileOps instance.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
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
//		current instance of FileOps.
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
func (fops *FileOps) CopyOut(
	errorPrefix interface{}) (
	FileOps,
	error) {

	if fops.lock == nil {
		fops.lock = new(sync.Mutex)
	}

	fops.lock.Lock()

	defer fops.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileOps."+
			"CopyOut()",
		"")

	if err != nil {
		return FileOps{}, err
	}

	return new(FileOperationsNanobot).copyOut(
		fops,
		ePrefix.XCpy("<-fops"))
}

// Equal - Returns 'true' if source, destination and opToExecute are
// equivalent. In other words both the current File Operations instance
// and the File Operations instance passed as an input parameter must
// have data fields which are equal in all respects.
//
// If any data field is found to be unequal, this method returns 'false'.
func (fops *FileOps) Equal(fops2 *FileOps) bool {

	if !fops.source.Equal(&fops2.source) {
		return false
	}

	if !fops.destination.Equal(&fops2.destination) {
		return false
	}

	if fops.opToExecute != fops2.opToExecute {
		return false
	}

	return true
}

// EqualPathFileNameExt - Compares two File Operations Types, 'FileOps'. The method
// returns 'true' if source and destination absolute paths, file
// names and file extensions are equivalent. Data Field 'opToExecute' is
// not included in the comparison.
//
// The absolute path, file name and file extension comparisons are
// case-insensitive. This means that all strings used in the comparisons are
// first converted to lower case before testing for equivalency.
//
// If the absolute paths, file names and file extensions are NOT equal,
// this method returns 'false'.
func (fops *FileOps) EqualPathFileNameExt(fops2 *FileOps) bool {

	if !fops.source.EqualPathFileNameExt(&fops2.source) {
		return false
	}

	if !fops.destination.EqualPathFileNameExt(&fops2.destination) {
		return false
	}

	return true
}

// IsInitialized - Returns a boolean value indicating whether
// this FileOps instance has been properly initialized.
func (fops *FileOps) IsInitialized() bool {
	return fops.isInitialized
}

// ExecuteFileOperation - Executes specific operations on the source
// and/or destination files configured and identified in the current
// FileOps instance.
//
// ------------------------------------------------------------------------
//
// Input Parameters:
//
//	FileOperationCode
//
//	The FileOperationCode type consists of the following
//	constants.
//
//	FileOperationCode(0).MoveSourceFileToDestinationFile()
//		Moves the source file to the destination file and
//		then deletes the original source file
//
//	FileOperationCode(0).DeleteDestinationFile()
//		Deletes the Destination file if it exists
//
//	FileOperationCode(0).DeleteSourceFile()
//		Deletes the Source file if it exists
//
//	FileOperationCode(0).DeleteSourceAndDestinationFiles
//		Deletes both the Source and Destination files
//		if they exist.
//
//	FileOperationCode(0).CopySourceToDestinationByHardLinkByIo()
//		Copies the Source File to the Destination
//		using two copy attempts. The first copy is
//		by Hard Link. If the first copy attempt fails,
//		a second copy attempt is initiated/ by creating
//		a new file and copying the contents by 'io.Copy'.
//		An error is returned only if both copy attempts
//		fail. The source file is unaffected.
//
//		See: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
//
//	FileOperationCode(0).CopySourceToDestinationByIoByHardLink()
//		Copies the Source File to the Destination
//		using two copy attempts. The first copy is
//		by 'io.Copy' which creates a new file and copies
//		the contents to the new file. If the first attempt
//		fails, a second copy attempt is initiated using
//		'copy by hard link'. An error is returned only
//		if both copy attempts fail. The source file is
//		unaffected.
//
//		See: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
//
//	FileOperationCode(0).CopySourceToDestinationByHardLink()
//		Copies the Source File to the Destination
//		using one copy mode. The only copy attempt
//		utilizes 'Copy by Hard Link'. If this fails
//		an error is returned.  The source file is
//		unaffected.
//
//	FileOperationCode(0).CopySourceToDestinationByIo()
//		Copies the Source File to the Destination
//		using only one copy mode. The only copy
//		attempt is initiated using 'Copy by IO' or
//		'io.Copy'.  If this fails an error is returned.
//		The source file is unaffected.
//
//	FileOperationCode(0).CreateSourceDir()
//		Creates the Source Directory
//
//	FileOperationCode(0).CreateSourceDirAndFile()
//		Creates the Source Directory and File
//
//	FileOperationCode(0).CreateSourceFile()
//		Creates the Source File
//
//	FileOperationCode(0).CreateDestinationDir()
//		Creates the Destination Directory
//
//	FileOperationCode(0).CreateDestinationDirAndFile()
//		Creates the Destination Directory and File
//
//	FileOperationCode(0).CreateDestinationFile()
//		Creates the Destination File
func (fops *FileOps) ExecuteFileOperation(
	fileOp FileOperationCode,
	errorPrefix interface{}) error {

	if fops.lock == nil {
		fops.lock = new(sync.Mutex)
	}

	fops.lock.Lock()

	defer fops.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	funcName := "FileOps.ExecuteFileOperation()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {
		return err
	}

	fops.opToExecute = fileOp

	err = nil

	fOpsElectron := FileOperationsElectron{}
	fOpsNanobot := FileOperationsNanobot{}

	switch fops.opToExecute {

	case FileOpCode.None():

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fileOp' is 'NONE' or No Operation!\n",
			ePrefix.String())

	case FileOpCode.MoveSourceFileToDestinationDir():

		err = fOpsNanobot.
			moveSourceFileToDestinationDir(
				fops,
				ePrefix)

	case FileOpCode.MoveSourceFileToDestinationFile():

		err = fOpsNanobot.
			moveSourceFileToDestinationFile(
				fops,
				ePrefix)

	case FileOpCode.DeleteDestinationFile():

		err = fOpsElectron.
			deleteDestinationFile(
				fops,
				ePrefix)

	case FileOpCode.DeleteSourceFile():

		err = fOpsElectron.deleteSourceFile(
			fops,
			ePrefix)

	case FileOpCode.DeleteSourceAndDestinationFiles():

		err = new(FileOperationsAtom).
			deleteSourceAndDestinationFiles(
				fops,
				ePrefix)

	case FileOpCode.CopySourceToDestinationByHardLinkByIo():

		err = fOpsNanobot.
			copySrcToDestByHardLinkByIo(
				fops,
				ePrefix)

	case FileOpCode.CopySourceToDestinationByIoByHardLink():

		err = fOpsNanobot.
			copySrcToDestByIoByHardLink(
				fops,
				ePrefix)

	case FileOpCode.CopySourceToDestinationByHardLink():

		err = fOpsNanobot.
			copySrcToDestByHardLink(
				fops,
				ePrefix)

	case FileOpCode.CopySourceToDestinationByIo():

		err = fOpsNanobot.
			copySrcToDestByIo(
				fops,
				ePrefix)

	case FileOpCode.CreateSourceDir():

		err = fOpsNanobot.
			createSrcDirectory(
				fops,
				ePrefix)

	case FileOpCode.CreateSourceDirAndFile():

		err = fOpsNanobot.
			createSrcDirectoryAndFile(
				fops,
				ePrefix)

	case FileOpCode.CreateSourceFile():

		err = fOpsNanobot.
			createSrcFile(
				fops,
				ePrefix)

	case FileOpCode.CreateDestinationDir():

		err = fOpsNanobot.
			createDestDirectory(
				fops,
				ePrefix)

	case FileOpCode.CreateDestinationDirAndFile():

		err = fOpsNanobot.
			createDestDirectoryAndFile(
				fops,
				ePrefix)

	case FileOpCode.CreateDestinationFile():

		err = fOpsNanobot.
			createDestFile(
				fops,
				ePrefix)

	default:
		err = errors.New("Invalid 'FileOperationCode' Execution Command! ")
	}

	if err != nil {
		return fmt.Errorf("%v\n"+""+
			"Error= \n%v\n",
			funcName,
			err.Error())
	}

	return nil
}

// GetSource
//
// Returns a deep copy of the source FileMgr instance
// encapsulated by the current instance of FileOps.
//
// The internal member variable which is copied and
// returned is identified as:
//
//	FileOps.source
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	FileMgr
//
//		If this method completes successfully, this
//		parameter will return a deep copy of the 'source'
//		File Manager encapsulated in the current instance
//		of FileOps.
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
func (fops *FileOps) GetSource(
	errorPrefix interface{}) (
	FileMgr,
	error) {

	if fops.lock == nil {
		fops.lock = new(sync.Mutex)
	}

	fops.lock.Lock()

	defer fops.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileOps."+
			"GetSource()",
		"")

	if err != nil {
		return FileMgr{}, err
	}

	return fops.source.CopyOut(
		ePrefix.XCpy(
			"fops.source"))
}

// GetDestination
//
// Returns a deep copy of the destination FileMgr
// instance.
//
// The internal member variable which is copied and
// returned is identified as:
//
//	FileOps.destination
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	FileMgr
//
//		If this method completes successfully, this
//		parameter will return a deep copy of the
//	 	'destination' File Manager encapsulated in the
//	 	current instance of FileOps.
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
func (fops *FileOps) GetDestination(
	errorPrefix interface{}) (
	FileMgr,
	error) {

	if fops.lock == nil {
		fops.lock = new(sync.Mutex)
	}

	fops.lock.Lock()

	defer fops.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileOps."+
			"GetDestination()",
		"")

	if err != nil {
		return FileMgr{}, err
	}

	return fops.destination.CopyOut(
		ePrefix.XCpy("fops.destination"))
}

// NewByFileMgrs
//
// Creates and returns a new FileOps instance based on
// input parameters 'source' and 'destination' File
// Managers.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	sourceFMgr					FileMgr
//
//		A concrete instance of FileMgr which is used
//		to configure the source File Manager contained
//		in the returned instance of FileOps.
//
//	destinationFMgr				FileMgr
//
//		A concrete instance of FileMgr which is used
//		to configure the destination File Manager
//		contained in the returned instance of FileOps.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	FileOps
//
//		If this method completes successfully, this
//		parameter will return a new, and properly
//		initialized instance, of FileOps.
//
//		Data values for this new instance will be
//		constructed from the source and destination File
//		Managers passed as input parameters 'sourceFMgr'
//		and 'destinationFMgr'.
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
func (fops *FileOps) NewByFileMgrs(
	sourceFMgr FileMgr,
	destinationFMgr FileMgr,
	errorPrefix interface{}) (
	FileOps,
	error) {

	if fops.lock == nil {
		fops.lock = new(sync.Mutex)
	}

	fops.lock.Lock()

	defer fops.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileOps."+
			"NewByFileMgrs()",
		"")

	if err != nil {
		return FileOps{}, err
	}

	fOpsNew := FileOps{}

	err = new(FileOperationsAtom).
		setFileOps(
			&fOpsNew,
			sourceFMgr,
			destinationFMgr,
			ePrefix.XCpy("fOpsNew"))

	return fOpsNew, nil
}

// NewByDirMgrFileName
//
// Creates and returns a new FileOps instance based on
// input parameters consisting of a source Directory
// Manger, a source file name and extension string, a
// destination Directory Manager and a destination file
// name and extension string.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	sourceDir					DirMgr
//
//		This instance of Directory Manager specifies the
//		source file directory which will be combined with
//		the source file name parameter to configure the
//		source file member variable 'fOps.source'
//		contained in the returned instance of FileOps.
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
//		contained in the returned instance of FileOps.
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
//		'fOps.destination' contained in the returned
//		instance of FileOps.
//
//		If this parameter is evaluated as invalid, an
//		error will be returned.
//
//	destinationFileNameExt		string
//
//		This string specifies the destination file name
//		and file extension which will be combined with
//		the destination directory parameter to configure
//		the destination file member variable
//		'fOps.destination' contained in the returned
//		instance of FileOps.
//
//		If this parameter is submitted as an empty string
//		with a zero (0) string length, an error will be
//		returned.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	FileOps
//
//		If this method completes successfully, this
//		parameter will return a new, and properly
//		initialized instance, of FileOps.
//
//		Data values for this new instance will be
//		constructed from the source and destination
//		Directory Managers file name strings passed as
//		input parameters.
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
func (fops *FileOps) NewByDirMgrFileName(
	sourceDir DirMgr,
	sourceFileNameExt string,
	destinationDir DirMgr,
	destinationFileNameExt string,
	errorPrefix interface{}) (
	FileOps,
	error) {

	if fops.lock == nil {
		fops.lock = new(sync.Mutex)
	}

	fops.lock.Lock()

	defer fops.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileOps."+
			"NewByDirMgrFileName()",
		"")

	if err != nil {
		return FileOps{}, err
	}

	fOpsNew := FileOps{}

	err = new(FileOperationsNanobot).
		setFileOpsByDirMgrFileName(
			&fOpsNew,
			sourceDir,
			sourceFileNameExt,
			destinationDir,
			destinationFileNameExt,
			ePrefix.XCpy("fOpsNew<-"))

	return fOpsNew, err
}

// NewByDirStrsAndFileNameExtStrs
//
// Creates and returns a new FileOps instance based on
// source and destination input parameters which consist
// of two pairs of directory strings and file name and
// extension strings.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	sourceDirStr				string
//
//		This string specifies the source file directory
//		which will be combined with the source file name
//		parameter to configure the source file member
//		variable 'fOps.source' contained in the returned
//		instance of FileOps.
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
//		contained in the returned instance of FileOps.
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
//		contained in the returned instance of FileOps.
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
//		'fOps.destination' contained in the returned
//		instance of FileOps.
//
//		If this parameter is submitted as an empty string
//		with a zero (0) string length, an error will be
//		returned.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	FileOps
//
//		If this method completes successfully, this
//		parameter will return a new, and properly
//		initialized instance, of FileOps.
//
//		Data values for this new instance will be
//		constructed from the source and destination
//		values passed as input parameters.
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
func (fops *FileOps) NewByDirStrsAndFileNameExtStrs(
	sourceDirStr string,
	sourceFileNameExtStr string,
	destinationDirStr string,
	destinationFileNameExtStr string,
	errorPrefix interface{}) (
	FileOps,
	error) {

	if fops.lock == nil {
		fops.lock = new(sync.Mutex)
	}

	fops.lock.Lock()

	defer fops.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileOps."+
			"NewByDirStrsAndFileNameExtStrs()",
		"")

	if err != nil {
		return FileOps{}, err
	}

	fOpsNew := FileOps{}

	err = new(FileOperationsNanobot).
		setFileOpsByDirAndFileNameStr(
			&fOpsNew,
			sourceDirStr,
			sourceFileNameExtStr,
			destinationDirStr,
			destinationFileNameExtStr,
			ePrefix.XCpy(
				"fOpsNew<-"))

	return fOpsNew, err
}

// NewByPathFileNameExtStrs
//
// Creates and returns a new FileOps instance based on
// two string input parameters.
//
// The first parameter represents the path name, file
// name and extension of the source file.
//
// The second represents the path name, file name and
// extension of the destination file.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	sourcePathFileNameExt		string
//
//		This string contains the source path and file
//		name used to configure the source File Manager
//		encapsulated in the returned instance of FileOps
//		(FileOps.source).
//
//		If this parameter is submitted as an empty string
//		with a zero (0) string length, an error will be
//		returned.
//
//	destinationPathFileNameExt	string
//
//		This string contains the destination path and
//		file name used to configure the destination File
//		Manager encapsulated in  the returned instance of
//		FileOps (FileOps.destination).
//
//		If this parameter is submitted as an empty string
//		with a zero (0) string length, an error will be
//		returned.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	FileOps
//
//		If this method completes successfully, this
//		parameter will return a new, and properly
//		initialized instance, of FileOps.
//
//		Data values for this new instance will be
//		constructed from the source and destination
//		values passed as input parameters.
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
func (fops *FileOps) NewByPathFileNameExtStrs(
	sourcePathFileNameExt string,
	destinationPathFileNameExt string,
	errorPrefix interface{}) (
	FileOps,
	error) {

	if fops.lock == nil {
		fops.lock = new(sync.Mutex)
	}

	fops.lock.Lock()

	defer fops.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileOps."+
			"NewByPathFileNameExtStrs()",
		"")

	if err != nil {
		return FileOps{}, err
	}

	fOpsNew := FileOps{}

	err = new(FileOperationsNanobot).
		setByPathFileNameExtStrs(
			&fOpsNew,
			sourcePathFileNameExt,
			destinationPathFileNameExt,
			ePrefix.XCpy(
				"fOpsNew<-"))

	return fOpsNew, err
}

// SetByFileMgrs
//
// Reconfigures the current instance of FileOps using new
// values for source and destination files passed as input
// parameters.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and reconfigure all
//	pre-existing source and destination data values in
//	the current instance of FileOps.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	sourceFMgr					FileMgr
//
//		A concrete instance of FileMgr which is used
//		to configure the source File Manager contained
//		in the current instance of FileOps.
//
//	destinationFMgr				FileMgr
//
//		A concrete instance of FileMgr which is used
//		to configure the destination File Manager
//		contained in the current instance of FileOps.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
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
func (fops *FileOps) SetByFileMgrs(
	sourceFMgr FileMgr,
	destinationFMgr FileMgr,
	errorPrefix interface{}) error {

	if fops.lock == nil {
		fops.lock = new(sync.Mutex)
	}

	fops.lock.Lock()

	defer fops.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileOps."+
			"SetByFileMgrs()",
		"")

	if err != nil {
		return err
	}

	return new(FileOperationsAtom).
		setFileOps(
			fops,
			sourceFMgr,
			destinationFMgr,
			ePrefix.XCpy("fops"))
}

// SetByDirMgrFileName
//
// Reconfigures the current FileOps instance based on
// input parameters consisting of a source Directory
// Manger, a source file name and extension string, a
// destination Directory Manager and a destination file
// name and extension string.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and reconfigure all
//	pre-existing source and destination data values in
//	the current instance of FileOps.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	sourceDirStr				DirMgr
//
//		This instance of Directory Manager specifies the
//		source file directory which will be combined with
//		the source file name parameter to configure the
//		source file member variable 'FileOps.source'
//		contained in the current instance of FileOps.
//
//		If this parameter is evaluated as invalid, an
//		error will be returned.
//
//	sourceFileNameExtStr		string
//
//		This string specifies the source file name and
//		file extension which will be combined with the
//		source directory parameter to configure the
//		source file member variable 'FileOps.source'
//		contained in the current instance of FileOps.
//
//		If this parameter is submitted as an empty string
//		with a zero (0) string length, an error will be
//		returned.
//
//	destinationDirStr			DirMgr
//
//		This instance of Directory Manager specifies the
//		destination file directory which will be combined
//		with the destination file name parameter to
//		configure the destination file member variable
//		'FileOps.destination' contained in the current
//		instance of FileOps.
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
//		'FileOps.destination' contained in the current
//		instance of FileOps.
//
//		If this parameter is submitted as an empty string
//		with a zero (0) string length, an error will be
//		returned.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
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
func (fops *FileOps) SetByDirMgrFileName(
	sourceDir DirMgr,
	sourceFileNameExt string,
	destinationDir DirMgr,
	destinationFileNameExt string,
	errorPrefix interface{}) error {

	if fops.lock == nil {
		fops.lock = new(sync.Mutex)
	}

	fops.lock.Lock()

	defer fops.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileOps."+
			"SetByFileMgrs()",
		"")

	if err != nil {
		return err
	}

	return new(FileOperationsNanobot).
		setFileOpsByDirMgrFileName(
			fops,
			sourceDir,
			sourceFileNameExt,
			destinationDir,
			destinationFileNameExt,
			ePrefix.XCpy("fops<-"))
}

// SetByDirStrsAndFileNameExtStrs
//
// Reconfigures the current FileOps instance based on
// input parameters consisting of a source Directory
// Manger, a source file name and extension string, a
// destination Directory Manager and a destination file
// name and extension string.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and reconfigure all
//	pre-existing source and destination data values in
//	the current instance of FileOps.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	sourceDirStr				string
//
//		This string specifies the source file directory
//		which will be combined with the source file name
//		parameter to configure the source file member
//		variable 'FileOps.source' contained in the
//		current instance of FileOps.
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
//		source file member variable 'FileOps.source'
//		contained in the current instance of FileOps.
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
//		contained in the current instance of FileOps.
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
//		'fOps.destination' contained in the current
//		instance of FileOps.
//
//		If this parameter is submitted as an empty string
//		with a zero (0) string length, an error will be
//		returned.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
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
func (fops *FileOps) SetByDirStrsAndFileNameExtStrs(
	sourceDirStr string,
	sourceFileNameExtStr string,
	destinationDirStr string,
	destinationFileNameExtStr string,
	errorPrefix interface{}) error {

	if fops.lock == nil {
		fops.lock = new(sync.Mutex)
	}

	fops.lock.Lock()

	defer fops.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileOps."+
			"SetByDirStrsAndFileNameExtStrs()",
		"")

	if err != nil {
		return err
	}

	return new(FileOperationsNanobot).
		setFileOpsByDirAndFileNameStr(
			fops,
			sourceDirStr,
			sourceFileNameExtStr,
			destinationDirStr,
			destinationFileNameExtStr,
			ePrefix.XCpy(
				"fops"))
}

// SetByPathFileNameExtStrs
//
// Reconfigures the current instance of FileOps using
// source and destination file names passed as input
// string parameters 'sourcePathFileName' and
// 'destinationPathFileName'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and reconfigure all
//	pre-existing source and destination data values in
//	the current instance of FileOps.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	sourcePathFileNameExt		string
//
//		This string contains the source path and file
//		name used to configure the source File Manager
//		encapsulated in the current instance of FileOps
//		(FileOps.source).
//
//		If this parameter is submitted as an empty string
//		with a zero (0) string length, an error will be
//		returned.
//
//	destinationPathFileNameExt	string
//
//		This string contains the destination path and
//		file name used to configure the destination File
//		Manager encapsulated in the current instance of
//		FileOps	(FileOps.destination).
//
//		If this parameter is submitted as an empty string
//		with a zero (0) string length, an error will be
//		returned.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
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
func (fops *FileOps) SetByPathFileNameExtStrs(
	sourcePathFileNameExt string,
	destinationPathFileNameExt string,
	errorPrefix interface{}) error {

	if fops.lock == nil {
		fops.lock = new(sync.Mutex)
	}

	fops.lock.Lock()

	defer fops.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileOps."+
			"SetByPathFileNameExtStrs()",
		"")

	if err != nil {
		return err
	}

	return new(FileOperationsNanobot).
		setByPathFileNameExtStrs(
			fops,
			sourcePathFileNameExt,
			destinationPathFileNameExt,
			ePrefix.XCpy(
				"fops"))
}

// SetFileOpsCode
//
// Sets the File Operations code for the current FileOps
// instance.
//
// The File Operations Code is an integer enumeration. It
// signals the type of operation to be performed on a file.
//
// This method stores the user specified File Operations
// Code internally for later use in the performance of file
// operations.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fOpCode						FileOperationCode
//
//		An enumeration value which will be stored
//		internal to the current FileOps instance for
//		later use in specifying the performance of
//		file operations.
//
//		Valid File Operation Codes are listed as
//		follows:
//
//			FileOperationCode(0).None()
//			FileOperationCode(0).MoveSourceFileToDestinationFile()
//			FileOperationCode(0).MoveSourceFileToDestinationDir()
//			FileOperationCode(0).DeleteDestinationFile()
//			FileOperationCode(0).DeleteSourceFile()
//			FileOperationCode(0).DeleteSourceAndDestinationFiles()
//			FileOperationCode(0).CopySourceToDestinationByHardLinkByIo()
//			FileOperationCode(0).CopySourceToDestinationByIoByHardLink()
//			FileOperationCode(0).CopySourceToDestinationByHardLink()
//			FileOperationCode(0).CopySourceToDestinationByIo()
//			FileOperationCode(0).CreateSourceDir()
//			FileOperationCode(0).CreateSourceDirAndFile()
//			FileOperationCode(0).CreateSourceFile()
//			FileOperationCode(0).CreateDestinationDir()
//			FileOperationCode(0).CreateDestinationDirAndFile()
//			FileOperationCode(0).CreateDestinationFile()
//
//		Users may find it easier to use the shorthand
//		notation for designating valid File Operation
//		Codes.
//
//			FileOpCode.None()
//			FileOpCode.MoveSourceFileToDestinationFile()
//			FileOpCode.MoveSourceFileToDestinationDir()
//			FileOpCode.DeleteDestinationFile()
//			FileOpCode.DeleteSourceFile()
//			FileOpCode.DeleteSourceAndDestinationFiles()
//			FileOpCode.CopySourceToDestinationByHardLinkByIo()
//			FileOpCode.CopySourceToDestinationByIoByHardLink()
//			FileOpCode.CopySourceToDestinationByHardLink()
//			FileOpCode.CopySourceToDestinationByIo()
//			FileOpCode.CreateSourceDir()
//			FileOpCode.CreateSourceDirAndFile()
//			FileOpCode.CreateSourceFile()
//			FileOpCode.CreateDestinationDir()
//			FileOpCode.CreateDestinationDirAndFile()
//			FileOpCode.CreateDestinationFile()
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
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
func (fops *FileOps) SetFileOpsCode(
	fOpCode FileOperationCode,
	errorPrefix interface{}) error {

	if fops.lock == nil {
		fops.lock = new(sync.Mutex)
	}

	fops.lock.Lock()

	defer fops.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	funcName := "FileOps.SetFileOpsCode()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {
		return err
	}

	err = fOpCode.IsValid()

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Error returned by fOpCode.IsValidInstanceError()\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			err.Error())
	}

	fops.opToExecute = fOpCode

	return err
}
