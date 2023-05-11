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
// operations performed on files. Hence, the name File
// Operations.
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

	switch fops.opToExecute {

	case FileOpCode.None():

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fileOp' is 'NONE' or No Operation!\n",
			ePrefix.String())

	case FileOpCode.MoveSourceFileToDestinationDir():
		err = fops.moveSourceFileToDestinationDir()

	case FileOpCode.MoveSourceFileToDestinationFile():
		err = fops.moveSourceFileToDestinationFile()

	case FileOpCode.DeleteDestinationFile():

		err = fOpsElectron.
			deleteDestinationFile(
				fops,
				ePrefix)

	case FileOpCode.DeleteSourceFile():
		err = fops.deleteSourceFile()

	case FileOpCode.DeleteSourceAndDestinationFiles():
		err = fops.deleteSourceAndDestinationFiles()

	case FileOpCode.CopySourceToDestinationByHardLinkByIo():
		err = fops.copySrcToDestByHardLinkByIo()

	case FileOpCode.CopySourceToDestinationByIoByHardLink():

		err = new(FileOperationsNanobot).
			copySrcToDestByIoByHardLink(
				fops,
				ePrefix)

	case FileOpCode.CopySourceToDestinationByHardLink():
		err = fops.copySrcToDestByHardLink()

	case FileOpCode.CopySourceToDestinationByIo():
		err = fops.copySrcToDestByIo()

	case FileOpCode.CreateSourceDir():
		err = fops.createSrcDirectory()

	case FileOpCode.CreateSourceDirAndFile():
		err = fops.createSrcDirectoryAndFile()

	case FileOpCode.CreateSourceFile():
		err = fops.createSrcFile()

	case FileOpCode.CreateDestinationDir():
		err = fops.createDestDirectory()

	case FileOpCode.CreateDestinationDirAndFile():
		err = fops.createDestDirectoryAndFile()

	case FileOpCode.CreateDestinationFile():
		err = fops.createDestFile()

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

// GetSource - Returns a deep copy of the
// source FileMgr instance.
func (fops *FileOps) GetSource() FileMgr {
	return fops.source.CopyOut()
}

// GetDestination - Returns a deep copy of the
// destination FileMgr instance.
func (fops *FileOps) GetDestination() FileMgr {
	return fops.destination.CopyOut()
}

// NewByFileMgrs - Creates and returns a new FileOps
// instance based on input parameters 'source' and
// 'destination' File Managers.
func (fops FileOps) NewByFileMgrs(
	source,
	destination FileMgr) (FileOps, error) {

	ePrefix := "FileOps.NewByFileMgrs() "

	err := source.IsFileMgrValid(ePrefix + "sourceFileMgr Error: ")

	if err != nil {
		return FileOps{},
			fmt.Errorf("Source File Manager INVALID!\n%v", err.Error())
	}

	err = destination.IsFileMgrValid(ePrefix + "destinationFileMgr Error: ")

	if err != nil {
		return FileOps{},
			fmt.Errorf("Destination File Manager INVALID!\n%v", err.Error())
	}

	fOpsNew := FileOps{}

	fOpsNew.source = source.CopyOut()
	fOpsNew.destination = destination.CopyOut()
	fOpsNew.isInitialized = true

	return fOpsNew, nil
}

// NewByDirMgrFileName - Creates and returns a new FileOps instance
// based on input parameters, source Directory Manger, source file name
// and extension string, destination Directory Manager and destination
// file name and extension string.
//
// If the destinationFileNameExt string is an empty string, it will default
// to the 'sourceFileNameExt'
func (fops FileOps) NewByDirMgrFileName(
	sourceDir DirMgr,
	sourceFileNameExt string,
	destinationDir DirMgr,
	destinationFileNameExt string) (FileOps, error) {

	ePrefix := "FileOps.NewByDirMgrFileName() "

	var err error

	fOpsNew := FileOps{}

	err = sourceDir.IsDirMgrValid(ePrefix + "sourceDir Error: ")

	if err != nil {
		return FileOps{}, err
	}

	fOpsNew.source, err = new(FileMgr).
		NewFromDirMgrFileNameExt(
			sourceDir,
			sourceFileNameExt,
			ePrefix)

	if err != nil {
		return FileOps{},
			fmt.Errorf(ePrefix+"Source File Error:\n%v\n", err.Error())
	}

	if len(destinationFileNameExt) == 0 {
		destinationFileNameExt = sourceFileNameExt
	}

	err = destinationDir.IsDirMgrValid(ePrefix + "destinationDir Error: ")

	if err != nil {
		return FileOps{}, err
	}

	fOpsNew.destination, err = new(FileMgr).
		NewFromDirMgrFileNameExt(
			destinationDir,
			destinationFileNameExt,
			ePrefix)

	if err != nil {
		return FileOps{},
			fmt.Errorf(ePrefix+"Destination File Error: %v", err.Error())
	}

	fOpsNew.isInitialized = true

	return fOpsNew, nil
}

// NewByDirStrsAndFileNameExtStrs - Creates and returns a new FileOps instance
// based on source and destination input parameters which consist of two pairs
// of directory or path strings and file name and extension strings.
//
// If the input parameter, 'destinationFileNameExtStr' is an empty string, it
// will be defaulted to a string value equal to 'sourceFileNameExtStr'.
func (fops *FileOps) NewByDirStrsAndFileNameExtStrs(
	sourceDirStr string,
	sourceFileNameExtStr string,
	destinationDirStr string,
	destinationFileNameExtStr string) (FileOps, error) {

	ePrefix := "FileOps.NewByPathFileNameExtStrs() "

	fOpsNew := FileOps{}

	var err error

	if len(sourceDirStr) == 0 {
		return FileOps{},
			fmt.Errorf("%v\n"+
				"Error: 'sourceDirStr' is an EMPTY STRING!\n",
				ePrefix)
	}

	if len(sourceFileNameExtStr) == 0 {
		return FileOps{},
			fmt.Errorf("%v\n"+
				"Error: 'sourceFileNameExtStr' is an EMPTY STRING!\n",
				ePrefix)
	}

	if len(destinationFileNameExtStr) == 0 {
		destinationFileNameExtStr = sourceFileNameExtStr
	}

	fOpsNew.source, err = new(FileMgr).
		NewFromDirStrFileNameStr(
			sourceDirStr,
			sourceFileNameExtStr,
			ePrefix)

	if err != nil {
		return FileOps{},
			fmt.Errorf(ePrefix+"Source File Error: %v", err.Error())
	}

	fOpsNew.destination, err = new(FileMgr).
		NewFromDirStrFileNameStr(
			destinationDirStr,
			destinationFileNameExtStr,
			ePrefix)

	if err != nil {
		return FileOps{},
			fmt.Errorf(ePrefix+"Destination File Error: %v", err.Error())
	}

	fOpsNew.isInitialized = true

	return fOpsNew, nil

}

// NewByPathFileNameExtStrs - Creates and returns a new FileOps instance
// based on two string input parameters. The first represents the path name,
// file name and extension of the source file. The second represents the
// path name, file name and extension of the destination file.
func (fops FileOps) NewByPathFileNameExtStrs(
	sourcePathFileNameExt string,
	destinationPathFileNameExt string) (FileOps, error) {

	ePrefix := "FileOps.NewByPathFileNameExtStrs() "

	fOpsNew := FileOps{}

	var err error

	fOpsNew.source, err = new(FileMgr).
		NewFromPathFileNameExtStr(
			sourcePathFileNameExt,
			ePrefix)

	if err != nil {
		return FileOps{},
			fmt.Errorf(ePrefix+"Source File Error: %v", err.Error())
	}

	fOpsNew.destination, err = new(FileMgr).
		NewFromPathFileNameExtStr(
			destinationPathFileNameExt,
			ePrefix)

	if err != nil {
		return FileOps{},
			fmt.Errorf(ePrefix+"Destination File Error: %v", err.Error())
	}

	fOpsNew.isInitialized = true

	return fOpsNew, nil
}

// SetFileOpsCode - Sets the File Operations code for the current FileOps
// instance.
func (fops *FileOps) SetFileOpsCode(fOpCode FileOperationCode) error {

	err := fOpCode.IsValid()

	if err != nil {
		return fmt.Errorf("FileOps.SetFileOpsCode()\n"+
			"Error returned by fOpCode.IsValidInstanceError()\nError='%v'", err.Error())
	}

	fops.opToExecute = fOpCode

	return nil
}

// deleteSourceFile - Deletes the source file in the current
// FileOps instance.
func (fops *FileOps) deleteSourceFile() error {

	ePrefix := "FileOps.FileOperationCode(0).DeleteSourceFile()() Source Deletion Failed: "

	err := fops.source.DeleteThisFile(ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+"%v", err.Error())
	}

	return nil
}

// deleteSourceAndDestinationFiles - Deletes both the source
// and destination files configured in the current FileOps
// instance.
func (fops *FileOps) deleteSourceAndDestinationFiles() error {

	cumErrMsg := "FileOps.FileOperationCode(0).DeleteSourceAndDestinationFiles()- "

	cumErrLen := len(cumErrMsg)

	err := fops.destination.DeleteThisFile(cumErrMsg)

	if err != nil {
		cumErrMsg += "Destination Deletion Failed: " + err.Error() + "\n"
	}

	err = fops.source.DeleteThisFile(cumErrMsg)

	if err != nil {
		cumErrMsg += "Source Deletion Failed: " + err.Error() + "\n"
	}

	if len(cumErrMsg) != cumErrLen {
		return errors.New(cumErrMsg)
	}

	return nil
}

func (fops *FileOps) copySrcToDestByHardLinkByIo() error {

	ePrefix := "FileOps.copySrcToDestByHardLinkByIo() "

	err := fops.source.CopyFileMgrByLinkByIo(
		&fops.destination,
		ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+"%v", err.Error())
	}

	return nil
}

// copySrcToDestByIo - Copies source file to destination
// using IO.
func (fops *FileOps) copySrcToDestByIo() error {

	ePrefix := "FileOps.copySrcToDestByIo() "

	err := fops.source.CopyFileMgrByIo(
		&fops.destination,
		ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+"%v", err.Error())
	}

	return nil
}

func (fops *FileOps) copySrcToDestByHardLink() error {

	ePrefix := "FileOps.copySrcToDestByHardLink() "

	err := fops.source.CopyFileMgrByLink(
		&fops.destination,
		ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+"%v", err.Error())
	}

	return nil
}

// createSrcDirectory - Creates the source directory using
// information from the FileOps source file manager.
func (fops *FileOps) createSrcDirectory() error {

	ePrefix := "FileOps.createSrcDirectory() "

	err := fops.source.CreateDir(ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+"%v", err.Error())
	}

	return nil
}

// createSrcDirectoryAndFile - Creates the source file
// and directory using information from the FileOps
// source file manager.
func (fops *FileOps) createSrcDirectoryAndFile() error {

	ePrefix := "FileOps.createSrcDirectoryAndFile() "

	err := fops.source.CreateDirAndFile(ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+"%v\n", err.Error())
	}

	err = fops.source.CloseThisFile(ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+"%v\n", err.Error())
	}

	return nil
}

// createSrcFile - Creates the source file using
// information from the FileOps source file manager.
func (fops *FileOps) createSrcFile() error {

	ePrefix := "FileOps.createSrcFile() "

	err := fops.source.CreateThisFile(ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+"%v\n", err.Error())
	}

	err = fops.source.CloseThisFile(ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+"%v\n", err.Error())
	}

	return nil
}

// createDestDirectory - Creates the destination
// directory using information from the FileOps
// destination file manager.
func (fops *FileOps) createDestDirectory() error {

	ePrefix := "FileOps.createSrcDirectory() "

	err := fops.destination.CreateDir(ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+"%v", err.Error())
	}

	return nil
}

// createDestDirectoryAndFile - Creates the destination
// directory and file using information from the FileOps
// destination file manager.
func (fops *FileOps) createDestDirectoryAndFile() error {

	ePrefix := "FileOps.createDestDirectoryAndFile() "

	err := fops.destination.CreateDirAndFile(ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+"%v\n", err.Error())
	}

	err = fops.destination.CloseThisFile(ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+"%v\n", err.Error())
	}

	return nil
}

// createDestFile - Creates the destination file using
// information from the FileOps destination file manager.
func (fops *FileOps) createDestFile() error {

	ePrefix := "FileOps.createDestFile() "

	err := fops.destination.CreateThisFile(ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+"%v", err.Error())
	}

	err = fops.destination.CloseThisFile(ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+"%v\n", err.Error())
	}

	return nil
}

// moveSourceFileFileToDestinationDir - Moves the source file
// to the destination directory. The file name will be designated
// by the destination file name.
func (fops *FileOps) moveSourceFileToDestinationDir() error {

	ePrefix := "FileOps.moveSourceFileToDestinationDir() "

	_, err := fops.source.
		MoveFileToNewDirMgr(
			fops.destination.GetDirMgr(),
			ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+"%v\n",
			err.Error())
	}

	return nil
}

// moveSourceFileToDestinationFile - Moves the source file
// to the destination by fist copying the source file
// to the destination and then deleting the source file.
//
// The final name will be the destination file in the
// destination directory.
func (fops *FileOps) moveSourceFileToDestinationFile() error {

	ePrefix := "FileOps.moveSourceFileToDestinationFile() "

	err := fops.source.MoveFileToFileMgr(
		fops.destination,
		ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+"%v\n", err.Error())
	}

	return nil
}
