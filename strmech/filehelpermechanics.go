package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"os"
	"sync"
)

type fileHelperMechanics struct {
	lock *sync.Mutex
}

// copyFileByIo
//
// Copies file from source path and file name to
// destination path and file name.
//
// Reference:
// https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// Note: Unlike the method CopyFileByLink above, this
// method does NOT rely on the creation of symbolic
// links. Instead, a new destination file is created and
// the contents of the source file are written to the new
// destination file using "io.Copy()".
//
// "io.Copy()" is the only method used to copy the
// designated source file. If this method fails, an error
// is returned.
//
// If source file is equivalent to the destination file,
// no action will be taken and no error will be returned.
//
// If the destination file does not exist, this method
// will create. However, it will NOT create the
// destination directory. If the destination directory
// does NOT exist, this method will abort the copy
// operation and return an error.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	sourceFile					string
//
//		This string holds the path and/or file name of
//	 	the source file. This source file will be copied
//		to the destination file.
//
//	destinationFile				string
//
//		This string holds the path and/or the file name
//		of the destination file. The source file taken
//		from input parameter 'sourceFile' will be copied
//		to this destination file.
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
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (fileHelpMech *fileHelperMechanics) copyFileByIo(
	sourceFile string,
	destinationFile string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fileHelpMech.lock == nil {
		fileHelpMech.lock = new(sync.Mutex)
	}

	fileHelpMech.lock.Lock()

	defer fileHelpMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileHelperMechanics."+
			"copyFileByIo()",
		"")

	if err != nil {
		return err
	}

	var err2, err3 error
	var srcFileDoesExist, dstFileDoesExist bool
	var srcFInfo, dstFileInfo FileInfoPlus

	fhMolecule := new(fileHelperMolecule)

	sourceFile,
		srcFileDoesExist,
		srcFInfo,
		err = fhMolecule.
		doesPathFileExist(
			sourceFile,
			PreProcPathCode.AbsolutePath(), // Convert to Absolute Path
			ePrefix,
			"sourceFile")

	if err != nil {
		return err
	}

	if !srcFileDoesExist {

		err = fmt.Errorf(
			"%v\n"+
				"Error: Source File DOES NOT EXIST!\n"+
				"sourceFile='%v'\n",
			ePrefix.String(),
			sourceFile)

		return err
	}

	destinationFile,
		dstFileDoesExist,
		dstFileInfo,
		err = fhMolecule.
		doesPathFileExist(
			destinationFile,
			PreProcPathCode.AbsolutePath(), // Convert to Absolute Path
			ePrefix,
			"destinationFile")

	if err != nil {
		return err
	}

	var areSameFile bool

	areSameFile,
		err2 = new(fileHelperNanobot).areSameFile(
		sourceFile,
		destinationFile,
		ePrefix.XCpy(
			"areSameFile<-"))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error occurred during path file name comparison.\n"+
			"Source File:'%v'\n"+
			"Destination File:'%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			sourceFile, destinationFile,
			err2.Error())

		return err
	}

	if areSameFile {

		err = fmt.Errorf("%v\n"+
			"Error: The source and destination file\n"+
			"are the same. They are equivalent.\n"+
			"Source File:'%v'\n"+
			"Destination File:'%v'\n",
			ePrefix.String(),
			sourceFile,
			destinationFile)

		return err
	}

	if srcFInfo.IsDir() {

		err = fmt.Errorf("%v\n"+
			"Error: Source File is 'Directory' and NOT a file!\n"+
			"Source File='%v'\n",
			ePrefix.String(),
			sourceFile)

		return err
	}

	if !srcFInfo.Mode().IsRegular() {
		// cannot copy non-regular files (e.g., directories,
		// symlinks, devices, etc.)

		err = fmt.Errorf("%v\n"+
			"Error non-regular source file ='%v'\n"+
			"source file Mode='%v'\n",
			ePrefix.String(),
			srcFInfo.Name(),
			srcFInfo.Mode().String())

		return err
	}

	if dstFileDoesExist && dstFileInfo.Mode().IsDir() {

		err = fmt.Errorf("%v\n"+
			"Error: 'destinationFile' is a Directory and NOT a File!\n"+
			"destinationFile='%v'",
			ePrefix.String(),
			destinationFile)

		return err
	}

	if dstFileDoesExist && !dstFileInfo.Mode().IsRegular() {
		err = fmt.Errorf("%v\n"+
			"Error: 'destinationFile' is NOT a 'Regular' File!\n"+
			"destinationFile='%v'\n",
			ePrefix.String(),
			destinationFile)

		return err
	}

	// If the destination file does NOT exist, this is not a problem
	// since it will be created later. If the destination 'Path' does
	// not exist, an error return will be triggered.

	// Create a new destination file and copy source
	// file contents to the destination file.

	// First, open the source file
	inSrcPtr, err2 := os.Open(sourceFile)

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"Error returned from os.Open(sourceFile) sourceFile='%v'\n"+
			"Error='%v'",
			ePrefix.String(),
			sourceFile,
			err2.Error())
		return err
	}

	// Next, 'Create' the destination file
	// If the destination file previously exists,
	// it will be truncated.
	outDestPtr, err2 := os.Create(destinationFile)

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"Error returned from os.Create(destinationFile)\n"+
			"destinationFile='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			destinationFile,
			err2.Error())

		_ = inSrcPtr.Close()

		return err
	}

	bytesCopied, err2 := io.Copy(outDestPtr, inSrcPtr)

	if err2 != nil {

		_ = inSrcPtr.Close()

		_ = outDestPtr.Close()

		err = fmt.Errorf("%v\n"+
			"Error returned from io.Copy(destination, source)\n"+
			"destination='%v'\n"+
			"source='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			destinationFile,
			sourceFile,
			err2.Error())

		return err
	}

	errs := make([]error, 0)

	// flush file buffers inSrcPtr memory
	err2 = outDestPtr.Sync()

	if err2 != nil {

		err3 = fmt.Errorf("%v\n"+
			"Error returned from outDestPtr.Sync()\n"+
			"outDestPtr=destination='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			destinationFile, err2.Error())

		errs = append(errs, err3)
	}

	err2 = inSrcPtr.Close()

	if err2 != nil {

		err3 = fmt.Errorf("%v\n"+
			"Error returned from inSrcPtr.Close()\n"+
			"inSrcPtr=source='%v'\nError='%v'\n",
			ePrefix.String(),
			sourceFile,
			err2.Error())

		errs = append(errs, err3)
	}

	inSrcPtr = nil

	err2 = outDestPtr.Close()

	if err2 != nil {

		err3 = fmt.Errorf("%v\n"+
			"Error returned from outDestPtr.Close()\n"+
			"outDestPtr=destination='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			destinationFile,
			err2.Error())

		errs = append(errs, err3)
	}

	outDestPtr = nil

	if len(errs) > 0 {

		return new(StrMech).ConsolidateErrors(errs)
	}

	_,
		dstFileDoesExist,
		dstFileInfo,
		err2 = fhMolecule.
		doesPathFileExist(
			destinationFile,
			PreProcPathCode.None(), // Do NOT alter path
			ePrefix,
			"destinationFile")

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: After Copy IO operation, destinationFile "+
			"generated non-path error!\n"+
			"destinationFile='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			destinationFile,
			err2.Error())

		return err
	}

	if !dstFileDoesExist {

		err = fmt.Errorf("%v\n"+
			"ERROR: After Copy IO operation, the destination file DOES NOT EXIST!\n"+
			"Destination File = 'destinationFile' = '%v'\n",
			ePrefix.String(),
			destinationFile)

		return err
	}

	srcFileSize := srcFInfo.Size()

	if bytesCopied != srcFileSize {

		err = fmt.Errorf("%v\n"+
			"Error: Bytes Copied does NOT equal bytes "+
			"in source file!\n"+
			"Source File Bytes='%v'\n"+
			"Bytes Coped='%v'\n"+
			"Source File=sourceFile='%v'\n"+
			"Destination File=destinationFile='%v'\n",
			ePrefix.String(),
			srcFileSize,
			bytesCopied,
			sourceFile,
			destinationFile)

		return err
	}

	if dstFileInfo.Size() != srcFileSize {

		err = fmt.Errorf("%v\n"+
			"Error: Bytes is source file do NOT equal bytes "+
			"in destination file!\n"+
			"Source File Bytes='%v'\n"+
			"Destination File Bytes='%v'\n"+
			"Source File=sourceFile='%v'\n"+
			"Destination File=destinationFile='%v'\n",
			ePrefix.String(),
			srcFileSize,
			dstFileInfo.Size(),
			sourceFile,
			destinationFile)

	}

	return err
}

// copyFileByLink
//
// Copies a file from source to destination by means of
// creating a 'hard link' to the source file,
// "os.Link(src, dst)".
//
// Note: This method of copying files does NOT create a
// new destination file and write the contents of the
// source file to destination file. (See CopyFileByIo
// Below).  Instead, this method performs the copy
// operation by creating a hard symbolic link to the
// source file.
//
// By creating a 'linked' file, changing the contents
// of one file will be reflected in the second. The
// two linked files are 'mirrors' of each other.
//
// Consider using CopyFileByIo() if this 'mirror' feature
// causes problems.
//
// "os.Link(src, dst)" is the only method employed to
// copy a designated file. If "os.Link(src, dst)" fails,
// an err is returned.
//
// See: https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
//
// REQUIREMENT: The destination Path must previously
// exist. The destination file need NOT exist as it will
// be created. If the destination file currently exists,
// it will first be deleted and a new linked file will be
// created.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	src							string
//
//		This string holds the file name and path for the
//		source file. which will be copied to the
//	 	destination file passed through input parameter
//	 	'dst'.
//
//	dst							string
//
//		This string holds the file name and path for the
//		destination file. The source file passed through
//	 	input parameter 'src' which will be copied to the
//	 	destination file identified by this parameter,
//	 	'dst'.
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
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (fileHelpMech *fileHelperMechanics) copyFileByLink(
	src string,
	dst string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fileHelpMech.lock == nil {
		fileHelpMech.lock = new(sync.Mutex)
	}

	fileHelpMech.lock.Lock()

	defer fileHelpMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileHelperMechanics."+
			"copyFileByLink()",
		"")

	if err != nil {
		return err
	}

	var err2 error
	var srcFileDoesExist, dstFileDoesExist bool
	var srcFInfo, dstFInfo FileInfoPlus

	fHelpMolecule := fileHelperMolecule{}
	src,
		srcFileDoesExist,
		srcFInfo,
		err = fHelpMolecule.doesPathFileExist(
		src,
		PreProcPathCode.AbsolutePath(), // Covert to Absolute Path
		ePrefix,
		"src")

	if err != nil {
		return err
	}

	dst,
		dstFileDoesExist,
		dstFInfo,
		err = fHelpMolecule.doesPathFileExist(
		dst,
		PreProcPathCode.AbsolutePath(), // Convert to Absolute Path
		ePrefix,
		"dst")

	if err != nil {
		return err
	}

	areSameFile, err2 :=
		new(fileHelperNanobot).areSameFile(
			src,
			dst,
			ePrefix.XCpy("areSameFile<-"))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error occurred during path file name comparison.\n"+
			"Source File:'%v'\n"+
			"Destination File:'%v'\n"+
			"Error='%v'\n",
			ePrefix,
			src,
			dst,
			err2.Error())

		return err
	}

	if areSameFile {

		err = fmt.Errorf("%v\n"+
			"Error: The source and destination file"+
			" are the same - equivalent.\n"+
			"Source File:'%v'\n"+
			"Destination File:'%v'\n",
			ePrefix.String(),
			src,
			dst)

		return err
	}

	if !srcFileDoesExist {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'src' file DOES NOT EXIST!\n"+
			"src='%v'\n",
			ePrefix.String(),
			src)

		return err
	}

	if srcFInfo.IsDir() {

		err = fmt.Errorf("%v\n"+
			"ERROR: Source File (src) is a 'Directory' NOT A FILE!\n"+
			"Source File (src)='%v'\n",
			ePrefix.String(),
			src)

		return err
	}

	if !srcFInfo.Mode().IsRegular() {
		// cannot copy non-regular files (e.g., directories,
		// symlinks, devices, etc.)

		err = fmt.Errorf("%v\n"+
			"Error: Non-regular source file.\n"+
			"Source File Name='%v'\n"+
			"Source File Mode='%v'\n",
			ePrefix.String(),
			srcFInfo.Name(),
			srcFInfo.Mode().String())

		return err
	}

	// If the destination file does NOT exist - this is not a problem
	// because the destination file will be created later.

	if dstFileDoesExist {
		// The destination file exists. This IS a problem. Link will
		// fail when attempting to create a link to an existing file.

		if dstFInfo.IsDir() {

			err = fmt.Errorf("%v\n"+
				"Error: The destination file ('dst') is NOT A FILE.\n"+
				"It is a DIRECTORY!\n"+
				"Destination File ('dst') = '%v'\n",
				ePrefix.String(),
				dst)

			return err
		}

		if !dstFInfo.Mode().IsRegular() {

			err = fmt.Errorf("%v\n"+
				"Error: The destination file ('dst') is NOT A REGULAR FILE.\n"+
				"Destination File ('dst') = '%v'\n",
				ePrefix.String(),
				dst)

			return err
		}

		err2 = os.Remove(dst)

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Error: The target destination file exists and could NOT be deleted!\n"+
				"destination file='%v'\n"+
				"Error='%v'\n",
				ePrefix.String(),
				dst,
				err2.Error())

			return err
		}

		dst,
			dstFileDoesExist,
			_,
			err = fHelpMolecule.
			doesPathFileExist(
				dst,
				PreProcPathCode.None(), // Apply no pre-processing conversion to 'dst'
				ePrefix,
				"dst")

		if err != nil {
			return err
		}

		if dstFileDoesExist {

			err = fmt.Errorf("%v\n"+
				"Error: Deletion of preexisting "+
				"destination file failed!\n"+
				"The copy link operation cannot proceed!\n"+
				"destination file='%v' ",
				ePrefix.String(),
				dst)

			return err
		}
	}

	err2 = os.Link(src, dst)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"os.Link(src, dst) FAILED!\n"+
			"src='%v'\n"+
			"dst='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			src,
			dst,
			err2.Error())

		return err
	}

	dst,
		dstFileDoesExist,
		_,
		err2 = fHelpMolecule.
		doesPathFileExist(
			dst,
			PreProcPathCode.AbsolutePath(), // Convert to Absolute Path
			ePrefix,
			"dst")

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: After Copy By Link Operation, a non-path error was returned on 'dst'.\n"+
			"Error='%v'",
			ePrefix,
			err2.Error())

		return err
	}

	if !dstFileDoesExist {

		err = fmt.Errorf("%v\n"+
			"Error: After Copy By Link Operation, the destination file DOES NOT EXIST!\n"+
			"Destination File= dst = %v",
			ePrefix.String(),
			dst)

	}

	return err
}

// OpenFile - wrapper for os.OpenFile. This method may be used to open or
// create files depending on the File Open and File Permission parameters.
//
// If successful, this method will return a pointer to the os.File object
// associated with the file designated for opening.
//
// The calling routine is responsible for calling "Close()" on this os.File
// pointer.
//
// ------------------------------------------------------------------------
//
// Input Parameters:
//
//	pathFileName                   string - A string containing the path and file name
//	                                        of the file which will be opened. If a parent
//	                                        path component does NOT exist, this method will
//	                                        trigger an error.
//
//	fileOpenCfg            FileOpenConfig - This parameter encapsulates the File Open parameters
//	                                        which will be used to open subject file. For an
//	                                        explanation of File Open parameters, see method
//	                                        FileOpenConfig.New().
//
// filePermissionCfg FilePermissionConfig - This parameter encapsulates the File Permission
//
//	parameters which will be used to open the subject
//	file. For an explanation of File Permission parameters,
//	see method FilePermissionConfig.New().
//
// ------------------------------------------------------------------------
//
// Return Values:
//
//	*os.File        - If successful, this method returns an os.File pointer
//	                  to the file designated by input parameter 'pathFileName'.
//	                  This file pointer can subsequently be used for reading
//	                  content from the subject file. It may NOT be used for
//	                  writing content to the subject file.
//
//	                  If this method fails, the *os.File return value is 'nil'.
//
//	                  Note: The caller is responsible for calling "Close()" on this
//	                  os.File pointer.
//
//
//	error           - If the method completes successfully, this return value
//	                  is 'nil'. If the method fails, the error type returned
//	                  is populated with an appropriate error message.
func (fileHelpMech *fileHelperMechanics) openFile(
	pathFileName string,
	fileOpenCfg FileOpenConfig,
	filePermissionCfg FilePermissionConfig,
	errorPrefix interface{}) (
	filePtr *os.File,
	err error) {

	if fileHelpMech.lock == nil {
		fileHelpMech.lock = new(sync.Mutex)
	}

	fileHelpMech.lock.Lock()

	defer fileHelpMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	filePtr = nil

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"fileHelperMechanics."+
			"openFile()",
		"")

	if err != nil {
		return filePtr, err
	}

	var pathFileNameDoesExist bool

	var fInfoPlus FileInfoPlus

	pathFileName,
		pathFileNameDoesExist,
		fInfoPlus,
		err = new(fileHelperMolecule).doesPathFileExist(
		pathFileName,
		PreProcPathCode.AbsolutePath(), // Convert to Absolute Path
		ePrefix,
		"pathFileName")

	if err != nil {
		return nil, err
	}

	if !pathFileNameDoesExist {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'pathFileName' DOES NOT EXIST!\n"+
			"pathFileName='%v'\n",
			ePrefix.String(),
			pathFileName)

		return filePtr, err
	}

	if fInfoPlus.IsDir() {
		err =
			fmt.Errorf("%v\n"+
				"ERROR: Input parameter 'pathFileName' is \n"+
				"a 'Directory' - NOT a file!\n"+
				"pathFileName='%v'\n",
				ePrefix.String(),
				pathFileName)

		return filePtr, err
	}

	err2 := fileOpenCfg.IsValid()

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Input Parameter 'fileOpenCfg' is INVALID!\n"+
			"Error='%v'\n",
			ePrefix.String(),
			err2.Error())

		return filePtr, err
	}

	fOpenCode, err2 := fileOpenCfg.GetCompositeFileOpenCode()

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"%v\n",
			ePrefix.String(),
			err2.Error())

		return filePtr, err
	}

	err2 = filePermissionCfg.IsValid(
		ePrefix)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Input Parameter 'filePermissionCfg' is INVALID!\n"+
			"Error='%v'\n",
			ePrefix.String(),
			err2.Error())

		return filePtr, err
	}

	fileMode, err2 := filePermissionCfg.
		GetCompositePermissionMode(ePrefix.XCpy(
			"fileMode<-"))

	if err2 != nil {
		err = fmt.Errorf(
			"%v\n"+"%v\n",
			ePrefix.String(),
			err2.Error())

		return filePtr, err
	}

	filePtr,
		err2 = os.OpenFile(
		pathFileName,
		fOpenCode,
		fileMode)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by os.OpenFile(pathFileName, fOpenCode, fileMode).\n"+
			"pathFileName='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			pathFileName,
			err2.Error())

		return filePtr, err
	}

	if filePtr == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: os.OpenFile() returned a 'nil' file pointer!\n",
			ePrefix.String())

		return filePtr, err
	}

	err = nil

	return filePtr, err
}
