package strmech

import (
	"errors"
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"os"
	"sync"
)

type fileHelperUtility struct {
	lock *sync.Mutex
}

// copyFileByIo
//
// Copies file from source path and file name to
// destination path and file name.
//
// Reference:
//
// https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang
// https://pkg.go.dev/io#Copy
//
// If the destination file does not exist, it will be
// created. If the directory path for the destination
// file does not exist, it too will be created
// depending on the setting for input parameter
// 'createDestDirPathIfNotExist'.
//
// Note: Unlike the method CopyFileByLink above, this
// method does NOT rely on the creation of symbolic
// links. Instead, a new destination file is created and
// the contents of the source file are written to the new
// destination file using:
//
//	io.Copy(dst Writer, src Reader)
//			(written int64, err error)
//
// "io.Copy()" is the only method used to copy the
// designated source file. If this method fails, an error
// is returned.
//
// If source file is equivalent to the destination file,
// an error will be returned.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	This method is a wrapper for io.Copy().
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
//		If the source file does NOT exist on an attached
//		storage device, an error will be returned.
//
//		If the source file is empty and contains zero (0)
//		bytes, an error will be returned.
//
//		If the source file is NOT classified as a
//		'regular' file, an error will be returned. This
//		means that Symlink files cannot be used as source
//		files.
//
//		If source file is equivalent to the destination
//		file, an error will be returned.
//
//	sourceFileLabel				string
//
//		The name or label associated with input parameter
//		'sourceFile' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "sourceFile" will be
//		automatically applied.
//
//	destinationFile				string
//
//		This string holds the path and/or the file name
//		of the destination file. The source file taken
//		from input parameter 'sourceFile' will be copied
//		to this destination file.
//
//		If the destination previously existed on an
//		attached storage device, it will be truncated.
//
//		If the destination file does not exist on an
//		attached storage device, that file will be
//		created automatically. If the directory path
//		for the destination file does not exist, it too
//		will be created depending on the setting for
//		input parameter 'createDestDirPathIfNotExist'.
//
//		If destination file is equivalent to the source
//		file, an error will be returned.
//
//	destinationFileLabel		string
//
//		The name or label associated with input parameter
//		'destinationFile' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "destinationFile" will
//		be automatically applied.
//
//	createDestDirPathIfNotExist	bool
//
//		If the directory path element of parameter
//		'destinationFile' does not exist on an attached
//		storage drive, and this parameter is set to
//		'true', this method will attempt to create
//		the directory path for 'destinationFile'.
//
//		If 'createDestDirPathIfNotExist' is set to
//		'false', and the directory path element of
//		parameter 'destinationFile' does not exist on
//		an attached storage drive, the copy operation
//		will fail and an error will be returned.
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
//	bytesCopied					int64
//
//		If this method completes successfully, this
//		return parameter will contain the number of
//		bytes copied from the source file to the
//		destination file. If no errors are present,
//		this value also represents the size of both
//		the source file and the destination file.
//
//	err							error
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
func (fHelperUtility *fileHelperUtility) copyFileByIo(
	sourceFile string,
	sourceFileLabel string,
	destinationFile string,
	destinationFileLabel string,
	createDestDirPathIfNotExist bool,
	errPrefDto *ePref.ErrPrefixDto) (
	bytesCopied int64,
	err error) {

	if fHelperUtility.lock == nil {
		fHelperUtility.lock = new(sync.Mutex)
	}

	fHelperUtility.lock.Lock()

	defer fHelperUtility.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileHelperUtility." +
		"copyFileByIo()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		return bytesCopied, err
	}

	if len(sourceFileLabel) == 0 {

		sourceFileLabel = "sourceFile"
	}

	if len(destinationFileLabel) == 0 {

		destinationFileLabel = "destinationFile"
	}

	var err2, err3 error
	var destFileDoesExist bool
	var srcFInfo, dstFileInfo FileInfoPlus

	var fHelperMicrobot = new(fileHelperMicrobot)

	sourceFile,
		srcFInfo,
		err2 = fHelperMicrobot.
		validateSourceFile(
			sourceFile,
			sourceFileLabel,
			true, // errorOnIrregularFile
			true, // errorOnEmptyFile
			ePrefix.XCpy(
				sourceFileLabel))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Source File Error!\n"+
			"Input parameter %v is invalid.\n"+
			"Error=\n%v\n",
			funcName,
			sourceFileLabel,
			err2.Error())

		return bytesCopied, err
	}

	destinationFile,
		destFileDoesExist,
		dstFileInfo,
		err2 = fHelperMicrobot.
		validateDestinationFile(
			destinationFile,
			destinationFileLabel,
			createDestDirPathIfNotExist,
			true, // errorOnIrregularFile
			ePrefix.XCpy(destinationFileLabel))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Destination File Error!\n"+
			"Input parameter %v is invalid.\n"+
			"Error=\n%v\n",
			funcName,
			destinationFileLabel,
			err2.Error())

		return bytesCopied, err
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

		return bytesCopied, err
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

		return bytesCopied, err
	}

	// Create a new destination file and copy source
	// file contents to the destination file.

	// First, open the source file
	var inSrcPtr *os.File

	inSrcPtr, err2 = os.Open(sourceFile)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned from os.Open(%v)\n"+
			"An error was returned while opening the source file!\n"+
			"%v= '%v'\n"+
			"Error='%v'",
			ePrefix.String(),
			sourceFileLabel,
			sourceFileLabel,
			sourceFile,
			err2.Error())

		return bytesCopied, err
	}

	// Next, 'Create' the destination file
	// If the destination file previously exists,
	// it will be truncated.

	var outDestPtr *os.File

	outDestPtr, err2 = os.Create(destinationFile)

	if err2 != nil {

		err3 = fmt.Errorf("%v\n"+
			"Error: os.Create(%v)\n"+
			"An error was returned while creating\n"+
			"the destination file, '%v'.\n"+
			"%v= '%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			destinationFileLabel,
			destinationFileLabel,
			destinationFileLabel,
			destinationFile,
			err2.Error())

		err = errors.Join(err, err3)

		err2 = inSrcPtr.Close()

		if err2 != nil {

			err3 = fmt.Errorf("%v\n"+
				"Error: inSrcPtr.Close()\n"+
				"Error returned while closing source\n"+
				"file, '%v'\n"+
				"%v= '%v'\n"+
				"Error=\n%v\n",
				ePrefix.String(),
				sourceFileLabel,
				sourceFileLabel,
				sourceFile,
				err2.Error())

			err = errors.Join(err, err3)

		}

		return bytesCopied, err
	}

	bytesCopied, err2 = io.Copy(outDestPtr, inSrcPtr)

	if err2 != nil {

		err3 = fmt.Errorf("%v\n"+
			"Error: io.Copy(%v, %v)\n"+
			"%v= '%v'\n"+
			"%v= '%v'\n"+
			"Error=\n'%v'\n",
			ePrefix.String(),
			destinationFileLabel,
			sourceFileLabel,
			destinationFileLabel,
			destinationFile,
			sourceFileLabel,
			sourceFile,
			err2.Error())

		err = errors.Join(err, err3)

		err2 = inSrcPtr.Close()

		if err2 != nil {

			err3 = fmt.Errorf("%v\n"+
				"Error: inSrcPtr.Close()\n"+
				"Error returned while closing source\n"+
				"file, '%v'\n"+
				"%v= '%v'\n"+
				"Error=\n%v\n",
				ePrefix.String(),
				sourceFileLabel,
				sourceFileLabel,
				sourceFile,
				err2.Error())

			err = errors.Join(err, err3)
		}

		err2 = outDestPtr.Close()

		if err2 != nil {

			err3 = fmt.Errorf("%v\n"+
				"Error: outDestPtr.Close()\n"+
				"Error returned while closing destination\n"+
				"file, '%v'\n"+
				"%v= '%v'\n"+
				"Error=\n%v\n",
				ePrefix.String(),
				destinationFileLabel,
				destinationFileLabel,
				destinationFile,
				err2.Error())

			err = errors.Join(err, err3)
		}

		inSrcPtr = nil
		outDestPtr = nil

		return bytesCopied, err
	}

	// flush file buffers inSrcPtr memory
	err2 = outDestPtr.Sync()

	if err2 != nil {

		err3 = fmt.Errorf("%v\n"+
			"Error: outDestPtr.Sync()\n"+
			"Error returned while flushing\n"+
			"destination file, '%v'\n"+
			"outDestPtr= %v ='%v'\n"+
			"Error=\n'%v'\n",
			ePrefix.String(),
			destinationFileLabel,
			destinationFileLabel,
			destinationFile,
			err2.Error())

		err = errors.Join(err, err3)
	}

	err2 = inSrcPtr.Close()

	if err2 != nil {

		err3 = fmt.Errorf("%v\n"+
			"Error: inSrcPtr.Close() after sync operation!\n"+
			"Error returned while closing source\n"+
			"file, '%v'.\n"+
			"%v= '%v'\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			sourceFileLabel,
			sourceFileLabel,
			sourceFile,
			err2.Error())

		err = errors.Join(err, err3)
	}

	err2 = outDestPtr.Close()

	if err2 != nil {

		err3 = fmt.Errorf("%v\n"+
			"Error: outDestPtr.Close() after sync operation!\n"+
			"Error returned while closing destination\n"+
			"file, '%v'\n"+
			"%v= '%v'\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			destinationFileLabel,
			destinationFileLabel,
			destinationFile,
			err2.Error())

		err = errors.Join(err, err3)
	}

	inSrcPtr = nil
	outDestPtr = nil

	if err != nil {

		return bytesCopied, err
	}

	_,
		destFileDoesExist,
		dstFileInfo,
		err2 = new(fileHelperMolecule).
		doesPathFileExist(
			destinationFile,
			PreProcPathCode.None(), // Do NOT alter path
			ePrefix,
			destinationFileLabel)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: Destination File Verification!\n"+
			"After the Io Buffer Copy operation,\n"+
			"Destination File, %v, generated a non-path\n"+
			"error.\n"+
			"%v='%v'\n"+
			"Error='%v'\n",
			funcName,
			destinationFileLabel,
			destinationFileLabel,
			destinationFile,
			err2.Error())

		return bytesCopied, err
	}

	if !destFileDoesExist {

		err = fmt.Errorf("%v\n"+
			"Error: Destination File Verification!\n"+
			"After Io Buffer Copy operation, the destination\n"+
			"file, %v, DOES NOT EXIST on disk!\n"+
			"%v= '%v'\n",
			ePrefix.String(),
			destinationFileLabel,
			destinationFileLabel,
			destinationFile)

		return bytesCopied, err
	}

	srcFileSize := srcFInfo.Size()

	if bytesCopied != srcFileSize {

		err = fmt.Errorf("%v\n"+
			"Error: Bytes Copied does NOT equal bytes "+
			"in source file, '%v'!\n"+
			"Source File, '%v', Size in Bytes='%v'\n"+
			"Bytes Copied='%v'\n"+
			"Destination File, '%v', Size in Bytes='%v'\n"+
			"Source File '%v'= '%v'\n"+
			"Destination File '%v'= '%v'\n",
			ePrefix.String(),
			sourceFileLabel,
			sourceFileLabel,
			srcFileSize,
			bytesCopied,
			destinationFileLabel,
			dstFileInfo.Size(),
			sourceFileLabel,
			sourceFile,
			destinationFileLabel,
			destinationFile)

		return bytesCopied, err
	}

	if dstFileInfo.Size() != srcFileSize {

		err = fmt.Errorf("%v\n"+
			"Error: Bytes is source file do NOT equal bytes "+
			"in destination file!\n"+
			"Source File, '%v', Size in Bytes='%v'\n"+
			"Destination File, '%v', Size in Bytes='%v'\n"+
			"%v= '%v'\n"+
			"%v= '%v'\n",
			ePrefix.String(),
			sourceFileLabel,
			srcFileSize,
			destinationFileLabel,
			dstFileInfo.Size(),
			sourceFileLabel,
			sourceFile,
			destinationFileLabel,
			destinationFile)

	}

	return bytesCopied, err
}

// copyFileByIoBuffer
//
// Copies file from source path and file name to
// destination path and file name.
//
// Reference:
//
//	https://pkg.go.dev/io#CopyBuffer
//
// If the destination file does not exist, it will be
// created.
//
// The contents of the source file are written to the
// destination file using:
//
//	io.CopyBuffer(dst Writer, src Reader, buf []byte)
//			 (written int64, err error)
//
// If source file is equivalent to the destination file,
// an error will be returned.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	This method is a wrapper for io.CopyBuffer.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	sourceFile						string
//
//		This string holds the path and/or file name of
//	 	the source file. This source file will be copied
//		to the destination file.
//
//		If the source file does NOT exist on an attached
//		storage device, an error will be returned.
//
//		If the source file is empty and contains zero (0)
//		bytes, an error will be returned.
//
//		If the source file is NOT classified as a
//		'regular' file, an error will be returned. This
//		means that Symlink files cannot be used as source
//		files.
//
//		If source file is equivalent to the destination
//		file, an error will be returned.
//
//	sourceFileLabel					string
//
//		The name or label associated with input parameter
//		'sourceFile' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "sourceFile" will be
//		automatically applied.
//
//	destinationFile					string
//
//		This string holds the path and/or the file name
//		of the destination file. The source file taken
//		from input parameter 'sourceFile' will be copied
//		to this destination file.
//
//		If the destination previously existed on an
//		attached storage device, it will be truncated.
//
//		If the destination file does not exist on an
//		attached storage device, that file will be
//		created automatically. If the directory path
//		for the destination file does not exist, it too
//		will be created depending on the setting for
//		input parameter 'createDestDirPathIfNotExist'.
//
//		If destination file is equivalent to the source
//		file, an error will be returned.
//
//	destinationFileLabel			string
//
//		The name or label associated with input parameter
//		'destinationFile' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "destinationFile" will
//		be automatically applied.
//
//	buffer							[]byte
//
//		This byte array will be used as an internal
//		buffer for the copy operation performed by this
//		method.
//
//		If 'buffer' is set to 'nil', a default internal
//		buffer will be used.
//
//		If 'buffer' has an array length of zero, a
//		default internal buffer will be used.
//
//	createDestDirPathIfNotExist	bool
//
//		If the directory path element of parameter
//		'destinationFile' does not exist on an attached
//		storage drive, and this parameter is set to
//		'true', this method will attempt to create
//		the directory path for 'destinationFile'.
//
//		If 'createDestDirPathIfNotExist' is set to
//		'false', and the directory path element of
//		parameter 'destinationFile' does not exist on
//		an attached storage drive, the copy operation
//		will fail and an error will be returned.
//
//	errPrefDto						*ePref.ErrPrefixDto
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
//	bytesCopied					int64
//
//		If this method completes successfully, this
//		return parameter will contain the number of
//		bytes copied from the source file to the
//		destination file. If no errors are present,
//		this value also represents the size of both
//		the source file and the destination file.
//
//	err							error
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
func (fHelperUtility *fileHelperUtility) copyFileByIoBuffer(
	sourceFile string,
	sourceFileLabel string,
	destinationFile string,
	destinationFileLabel string,
	buffer []byte,
	createDestDirPathIfNotExist bool,
	errPrefDto *ePref.ErrPrefixDto) (
	bytesCopied int64,
	err error) {

	if fHelperUtility.lock == nil {
		fHelperUtility.lock = new(sync.Mutex)
	}

	fHelperUtility.lock.Lock()

	defer fHelperUtility.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileHelperUtility." +
		"copyFileByIoBuffer()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		return bytesCopied, err
	}

	if len(sourceFileLabel) == 0 {

		sourceFileLabel = "sourceFile"
	}

	if len(destinationFileLabel) == 0 {

		destinationFileLabel = "destinationFile"
	}

	var err2, err3 error
	var destFileDoesExist bool
	var srcFInfo, dstFileInfo FileInfoPlus

	var fHelperMicrobot = new(fileHelperMicrobot)

	sourceFile,
		srcFInfo,
		err2 = fHelperMicrobot.
		validateSourceFile(
			sourceFile,
			sourceFileLabel,
			true, // errorOnIrregularFile
			true, // errorOnEmptyFile
			ePrefix.XCpy(
				sourceFileLabel))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Source File Error!\n"+
			"Input parameter %v is invalid.\n"+
			"Error=\n%v\n",
			funcName,
			sourceFileLabel,
			err2.Error())

		return bytesCopied, err
	}

	destinationFile,
		destFileDoesExist,
		dstFileInfo,
		err2 = fHelperMicrobot.
		validateDestinationFile(
			destinationFile,
			destinationFileLabel,
			createDestDirPathIfNotExist,
			true, // errorOnIrregularFile
			ePrefix.XCpy(destinationFileLabel))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Destination File Error!\n"+
			"Input parameter %v is invalid.\n"+
			"Error=\n%v\n",
			funcName,
			destinationFileLabel,
			err2.Error())

		return bytesCopied, err
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
			"%v: '%v'\n"+
			"%v: '%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			sourceFileLabel,
			sourceFile,
			destinationFileLabel,
			destinationFile,
			err2.Error())

		return bytesCopied, err
	}

	if areSameFile {

		err = fmt.Errorf("%v\n"+
			"Error: The source and destination file\n"+
			"are the same. They are equivalent.\n"+
			"%v: '%v'\n"+
			"%v: '%v'\n",
			ePrefix.String(),
			sourceFileLabel,
			sourceFile,
			destinationFileLabel,
			destinationFile)

		return bytesCopied, err
	}

	// Create a new destination file and copy source
	// file contents to the destination file.

	// First, open the source file
	var inSrcPtr *os.File

	inSrcPtr, err2 = os.Open(sourceFile)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned from os.Open(%v)\n"+
			"An error was returned while opening the source file!\n"+
			"%v= '%v'\n"+
			"Error='%v'",
			ePrefix.String(),
			sourceFileLabel,
			sourceFileLabel,
			sourceFile,
			err2.Error())

		return bytesCopied, err
	}

	// Next, 'Create' the destination file
	// If the destination file previously exists,
	// it will be truncated.

	var outDestPtr *os.File

	outDestPtr,
		err2 = os.Create(destinationFile)

	if err2 != nil {

		err3 = fmt.Errorf("%v\n"+
			"Error: os.Create(%v)\n"+
			"An error was returned while creating\n"+
			"the destination file, '%v'.\n"+
			"%v= '%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			destinationFileLabel,
			destinationFileLabel,
			destinationFileLabel,
			destinationFile,
			err2.Error())

		err = errors.Join(err, err3)

		err2 = inSrcPtr.Close()

		if err2 != nil {

			err3 = fmt.Errorf("%v\n"+
				"Error: inSrcPtr.Close()\n"+
				"Error returned while closing source\n"+
				"file, '%v'\n"+
				"%v= '%v'\n"+
				"Error=\n%v\n",
				ePrefix.String(),
				sourceFileLabel,
				sourceFileLabel,
				sourceFile,
				err2.Error())

			err = errors.Join(err, err3)

		}

		return bytesCopied, err
	}

	bytesCopied,
		err2 = io.CopyBuffer(
		outDestPtr,
		inSrcPtr,
		buffer)

	if err2 != nil {

		err3 = fmt.Errorf("%v\n"+
			"Error: io.CopyBuffer(%v, %v)\n"+
			"%v= '%v'\n"+
			"%v= '%v'\n"+
			"Error=\n'%v'\n",
			ePrefix.String(),
			destinationFileLabel,
			sourceFileLabel,
			destinationFileLabel,
			destinationFile,
			sourceFileLabel,
			sourceFile,
			err2.Error())

		err = errors.Join(err, err3)

		err2 = inSrcPtr.Close()

		if err2 != nil {

			err3 = fmt.Errorf("%v\n"+
				"Error: inSrcPtr.Close()\n"+
				"Error returned while closing source\n"+
				"file, '%v'\n"+
				"%v= '%v'\n"+
				"Error=\n%v\n",
				ePrefix.String(),
				sourceFileLabel,
				sourceFileLabel,
				sourceFile,
				err2.Error())

			err = errors.Join(err, err3)
		}

		err2 = outDestPtr.Close()

		if err2 != nil {

			err3 = fmt.Errorf("%v\n"+
				"Error: outDestPtr.Close()\n"+
				"Error returned while closing destination\n"+
				"file, '%v'\n"+
				"%v= '%v'\n"+
				"Error=\n%v\n",
				ePrefix.String(),
				destinationFileLabel,
				destinationFileLabel,
				destinationFile,
				err2.Error())

			err = errors.Join(err, err3)
		}

		inSrcPtr = nil
		outDestPtr = nil

		return bytesCopied, err
	}

	// flush file buffers outDestPtr memory
	err2 = outDestPtr.Sync()

	if err2 != nil {

		err3 = fmt.Errorf("%v\n"+
			"Error: outDestPtr.Sync()\n"+
			"Error returned while flushing\n"+
			"destination file, '%v'\n"+
			"outDestPtr= %v ='%v'\n"+
			"Error=\n'%v'\n",
			ePrefix.String(),
			destinationFileLabel,
			destinationFileLabel,
			destinationFile,
			err2.Error())

		err = errors.Join(err, err3)

	}

	err2 = inSrcPtr.Close()

	if err2 != nil {

		err3 = fmt.Errorf("%v\n"+
			"Error: inSrcPtr.Close() after sync operation!\n"+
			"Error returned while closing source\n"+
			"file, '%v'.\n"+
			"%v= '%v'\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			sourceFileLabel,
			sourceFileLabel,
			sourceFile,
			err2.Error())

		err = errors.Join(err, err3)
	}

	err2 = outDestPtr.Close()

	if err2 != nil {

		err3 = fmt.Errorf("%v\n"+
			"Error: outDestPtr.Close() after sync operation!\n"+
			"Error returned while closing destination\n"+
			"file, '%v'\n"+
			"%v= '%v'\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			destinationFileLabel,
			destinationFileLabel,
			destinationFile,
			err2.Error())

		err = errors.Join(err, err3)
	}

	inSrcPtr = nil
	outDestPtr = nil

	if err != nil {

		return bytesCopied, err
	}

	_,
		destFileDoesExist,
		dstFileInfo,
		err2 = new(fileHelperMolecule).
		doesPathFileExist(
			destinationFile,
			PreProcPathCode.None(), // Do NOT alter path
			ePrefix,
			destinationFileLabel)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: Destination File Verification!\n"+
			"After the Io Buffer Copy operation,\n"+
			"Destination File, %v, generated a non-path\n"+
			"error.\n"+
			"%v='%v'\n"+
			"Error='%v'\n",
			funcName,
			destinationFileLabel,
			destinationFileLabel,
			destinationFile,
			err2.Error())

		return bytesCopied, err
	}

	if !destFileDoesExist {

		err = fmt.Errorf("%v\n"+
			"Error: Destination File Verification!\n"+
			"After Io Buffer Copy operation, the destination\n"+
			"file, %v, DOES NOT EXIST on disk!\n"+
			"%v= '%v'\n",
			ePrefix.String(),
			destinationFileLabel,
			destinationFileLabel,
			destinationFile)

		return bytesCopied, err
	}

	srcFileSize := srcFInfo.Size()

	if bytesCopied != srcFileSize {

		err = fmt.Errorf("%v\n"+
			"Error: Bytes Copied does NOT equal bytes "+
			"in source file, '%v'!\n"+
			"Source File, '%v', Size in Bytes='%v'\n"+
			"Bytes Copied='%v'\n"+
			"Destination File, '%v', Size in Bytes='%v'\n"+
			"Source File '%v'= '%v'\n"+
			"Destination File '%v'= '%v'\n",
			ePrefix.String(),
			sourceFileLabel,
			sourceFileLabel,
			srcFileSize,
			bytesCopied,
			destinationFileLabel,
			dstFileInfo.Size(),
			sourceFileLabel,
			sourceFile,
			destinationFileLabel,
			destinationFile)

		return bytesCopied, err
	}

	if dstFileInfo.Size() != srcFileSize {

		err = fmt.Errorf("%v\n"+
			"Error: Bytes is source file do NOT equal bytes "+
			"in destination file!\n"+
			"Source File, '%v', Size in Bytes='%v'\n"+
			"Destination File, '%v', Size in Bytes='%v'\n"+
			"%v= '%v'\n"+
			"%v= '%v'\n",
			ePrefix.String(),
			sourceFileLabel,
			srcFileSize,
			destinationFileLabel,
			dstFileInfo.Size(),
			sourceFileLabel,
			sourceFile,
			destinationFileLabel,
			destinationFile)

	}

	return bytesCopied, err
}
