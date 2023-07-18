package strmech

import (
	"bufio"
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"os"
	"sync"
)

// FileBufferWriter
//
// This structure and the associated methods are designed
// to facilitate data 'write' operations. The most common
// destination for these 'write' operations is assumed to
// be a data file residing on an attached storage drive.
// However, any object implementing the io.Writer
// interface may be used as a 'write' destination.
//
// The FileBufferWriter type is a wrapper for
// 'bufio.Writer'. As such, FileBufferWriter supports
// incremental or buffered 'write' operations to the
// target output destination.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://pkg.go.dev/bufio
//	https://pkg.go.dev/bufio#Writer
//	https://pkg.go.dev/io#Writer
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	Use methods 'New' and 'NewPathFileName' to create
//		new instances of FileBufferWriter.
//
//	(2)	FileBufferWriter implements the io.Writer
//		interface.
type FileBufferWriter struct {
	fileWriter          *bufio.Writer
	filePtr             *os.File
	targetWriteFileName string

	lock *sync.Mutex
}

// Close
//
// This method is used to close any open file pointers
// and perform necessary clean-up operations after all
// data has been written to the destination io.Writer
// object.
//
// After calling this method, the current instance of
// FileBufferWriter will be unusable and should be
// discarded.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Call this method after completing all write operations.
//	Calling this method is essential to performance of
//	necessary clean-up tasks.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	--- NONE ---
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	--- NONE ---
func (fBufWriter *FileBufferWriter) Close() {

	if fBufWriter.lock == nil {
		fBufWriter.lock = new(sync.Mutex)
	}

	fBufWriter.lock.Lock()

	defer fBufWriter.lock.Unlock()

	if fBufWriter.filePtr != nil {

		_ = fBufWriter.filePtr.Close()

		fBufWriter.filePtr = nil
	}

	if fBufWriter.fileWriter != nil {

		fBufWriter.fileWriter = nil
	}

	return
}

// New
//
// This method returns a fully initialized instance of
// FileBufferWriter.
//
// The size of the internal read buffer is controlled by
// input parameter 'bufSize'. The minimum buffer size is
// 16-bytes.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	writer						io.Writer
//
//		An object which implements io.Writer interface.
//
//		This object may be a file pointer of type *os.File.
//		File pointers of this type implement the io.Writer
//		interface.
//
//		A file pointer (*os.File) will facilitate writing
//		file data to files residing on an attached
//		storage drive. However, with this configuration,
//		the user is responsible for manually closing the
//		file and performing any other required clean-up
//		operations.
//
//		While the returned instance of FileBufferWriter
//		is primarily designed for writing data to disk
//		files, this 'writer' will in fact write data to
//		any object implementing the io.Writer interface.
//
//	bufSize						int
//
//		This integer value controls the size of the
//		buffer created for the returned instance of
//		FileBufferWriter.
//
//		'bufSize' should be configured to maximize
//		performance for 'write' operations subject to
//		prevailing memory limitations.
//
//		If 'bufSize' is set to a value less than or equal
//		to zero (0), it will automatically be reset to
//		the default buffer size of 4096-bytes.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	FileBufferWriter
//
//		This method will return a fully configured
//		instance of FileBufferWriter.
func (fBufWriter *FileBufferWriter) New(
	writer io.Writer,
	bufSize int) FileBufferWriter {

	if fBufWriter.lock == nil {
		fBufWriter.lock = new(sync.Mutex)
	}

	fBufWriter.lock.Lock()

	defer fBufWriter.lock.Unlock()

	if bufSize <= 0 {

		bufSize = 4096
	}

	var newFileBufWriter FileBufferWriter

	newFileBufWriter.fileWriter = bufio.NewWriterSize(
		writer,
		bufSize)

	return newFileBufWriter
}

// NewPathFileName
//
// Receives a path and file name as an input parameter.
// This file is opened for 'write-only' operations and
// is configured in the returned instance
// FileBufferWriter.
//
// The size of the internal 'write' buffer is controlled
// by input parameter 'bufSize'. If 'bufSize' is set to a
// value less than or equal to zero (0), it will be
// automatically reset to the default value of
// 4096-bytes.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathFileName				string
//
//		This string contains the path and file name of
//		the file which will be used a data source for
//		'read' operations performed by method:
//			FileBufferReader.Read()
//
//		If this file does not currently exist on an
//		attached storage drive, an error will be
//		returned.
//
//	bufSize						int
//
//		This integer value controls the size of the
//		'write' buffer created for the returned instance
//		of FileBufferWriter.
//
//		'bufSize' should be configured to maximize
//		performance for 'write' operations subject to
//		prevailing memory limitations.
//
//		If 'bufSize' is set to a value less than or equal
//		to zero (0), it will be automatically reset to
//		the default value of 4096-bytes.
//
//	truncateExistingFile			bool
//
//		If this parameter is set to 'true', the target
//		'write' file will be opened for write operations.
//		If the target file previously existed, it will be
//		truncated. This means that the file's previous
//		contents will be deleted.
//
//		If this parameter is set to 'false', the target
//		file will be opened for write operations. If the
//		target file previously existed, the new text
//		written to the file will be appended to the
//		end of the previous file contents.
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
//	FileBufferWriter
//
//		If this method completes successfully, a fully
//		configured instance of FileBufferWriter will
//		be returned.
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
func (fBufWriter *FileBufferWriter) NewPathFileName(
	pathFileName string,
	bufSize int,
	truncateExistingFile bool,
	errorPrefix interface{}) (
	FileBufferWriter,
	error) {

	if fBufWriter.lock == nil {
		fBufWriter.lock = new(sync.Mutex)
	}

	fBufWriter.lock.Lock()

	defer fBufWriter.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error
	var newFileBufWriter FileBufferWriter

	funcName := "FileBufferWriter." +
		"NewPathFileName()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {
		return newFileBufWriter, err
	}

	if len(pathFileName) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'pathFileName' is invalid!\n"+
			"'pathFileName' is an empty string with a length of zero (0).\n",
			ePrefix.String())

		return newFileBufWriter, err
	}

	if bufSize <= 0 {

		bufSize = 4096

	}

	var fInfoPlus FileInfoPlus
	var pathFileDoesExist bool
	var err2 error

	pathFileName,
		pathFileDoesExist,
		fInfoPlus,
		err2 =
		new(fileHelperMolecule).
			doesPathFileExist(
				pathFileName,
				PreProcPathCode.AbsolutePath(), // Convert to Absolute Path
				ePrefix,
				"pathFileName")

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"An error occurred while testing for the existance\n"+
			"of 'pathFileName' on an attached storage drive.\n"+
			"pathFileName = '%v'\n"+
			"Error= \n%v\n",
			funcName,
			pathFileName,
			err2.Error())

		return newFileBufWriter, err
	}

	if !pathFileDoesExist {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'pathFileName' is invalid!\n"+
			"The path and file name do NOT exist on an attached\n"+
			"storage drive.\n"+
			"pathFileName= '%v'\n",
			ePrefix.String(),
			pathFileName)

		return newFileBufWriter, err
	}

	if fInfoPlus.IsDir() {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'pathFileName' is invalid!\n"+
			"'pathFileName' is directory and NOT a file name.\n"+
			"pathFileName= '%v'\n",
			ePrefix.String(),
			pathFileName)

		return newFileBufWriter, err
	}

	var filePermissionCfg FilePermissionConfig

	filePermissionCfg,
		err = new(FilePermissionConfig).New(
		"--w--w--w-",
		ePrefix.XCpy("filePermissionCfg<-"))

	if err != nil {

		return newFileBufWriter, err
	}

	var fileOpenCfg FileOpenConfig

	if truncateExistingFile {

		fileOpenCfg,
			err = new(FileOpenConfig).New(
			ePrefix.XCpy("fileOpenCfg<-"),
			FOpenType.TypeWriteOnly(),
			FOpenMode.ModeCreate(),
			FOpenMode.ModeTruncate())

		if err != nil {

			return newFileBufWriter, err
		}

	} else {
		// truncateExistingFile = 'false'
		// This signals Append to existing file.

		fileOpenCfg,
			err = new(FileOpenConfig).New(
			ePrefix.XCpy("fileOpenCfg<-"),
			FOpenType.TypeWriteOnly(),
			FOpenMode.ModeCreate(),
			FOpenMode.ModeAppend())

		if err != nil {

			return newFileBufWriter, err
		}

	}

	newFileBufWriter.filePtr,
		err = new(fileHelperBoson).
		openFile(
			pathFileName,
			false,
			fileOpenCfg,
			filePermissionCfg,
			"pathFileName",
			ePrefix)

	if err != nil {

		if newFileBufWriter.filePtr != nil {
			_ = newFileBufWriter.filePtr.Close()
		}

		return newFileBufWriter, err
	}

	newFileBufWriter.targetWriteFileName = pathFileName

	newFileBufWriter.fileWriter = bufio.NewWriterSize(
		newFileBufWriter.filePtr,
		bufSize)

	return newFileBufWriter, err
}

// Write
//
// Writes the contents of the byte array input paramter
// ('bytesToWrite') to the destination io.Writer object
// previously configured for this instance of
// FileBufferWriter.
//
// If for any reason, the returned number of bytes
// written ('numBytesWritten') to the destination
// io.Writer object is less than the length of the byte
// array passed as input parameter 'bytesToWrite', an
// error containing an explanation for this event will
// be returned.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	After all write operations have been completed, the
//	user MUST call FileBufferWriter.Close() to perform
//	necessary clean-up operations.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	bytesToWrite				[]byte
//
//		The contents of this byte array will be written
//		to the destination io.Writer object previously
//		configured for the current instance of
//		FileBufferWriter.
//
//		Typically, the destination io.Writer object will
//		be a data file existing on an attached storage
//		drive. However, the destination io.Writer object
//		may be any object implementing the io.Writer
//		interface.
//
//		If for any reason, the returned number of bytes
//		written ('numBytesWritten') to the destination
//		io.Writer object is less than the length of this
//		byte array ('bytesToWrite'), an error containing
//		an explanation for this event will be returned.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	numBytesWritten				int
//
//		This parameter returns the number of bytes
//		written to the destination io.Writer object
//		configured for the current instance of
//		FileBufferWriter.
//
//		If 'numBytesWritten' is less than the length
//		of the byte array input parameter 'bytesToWrite',
//		an error will also be returned.
//
//	err							error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If processing errors are encountered, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
//
//		If an end of file is encountered (after reading
//		all data source contents), this returned error
//		will be set to 'io.EOF'. See the 'Reference'
//		section for a discussion of 'io.EOF'. Disk files
//		will return an 'io.EOF'. However, some other
//		types of readers may not.
func (fBufWriter *FileBufferWriter) Write(
	bytesToWrite []byte) (
	numBytesWritten int,
	err error) {

	if fBufWriter.lock == nil {
		fBufWriter.lock = new(sync.Mutex)
	}

	fBufWriter.lock.Lock()

	defer fBufWriter.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"FileBufferWriter."+
			"Write()",
		"")

	if err != nil {

		return numBytesWritten, err
	}

	if fBufWriter.fileWriter == nil {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of 'FileBufferWriter' is invalid!\n"+
			"The internal bufio.Writer has NOT been initialized.\n"+
			"Call one of the 'New' methods when creating an instance\n"+
			"of 'FileBufferWriter'\n",
			ePrefix.String())

		return numBytesWritten, err
	}

	if len(bytesToWrite) <= 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'bytesToWrite' is invalid!\n"+
			"The 'bytesToWrite' byte array is empty. It has zero bytes.\n",
			ePrefix.String())

		return numBytesWritten, err
	}

	var err2 error

	numBytesWritten,
		err2 = fBufWriter.fileWriter.Write(bytesToWrite)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by fBufWriter.fileWriter.Write(bytesToWrite).\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			err2.Error())
	}

	return numBytesWritten, err
}
