package strmech

import (
	"bufio"
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"os"
	"sync"
)

// FileBufferReader
//
// This structure and the associated methods are designed
// to facilitate data read operations. The most common
// data source for these read operations is assumed to be
// a data file residing on an attached storage drive.
// However, any object implementing the io.Reader
// interface may be used as a data source.
//
// The FileBufferReader type is a wrapper for
// 'bufio.Reader'. As such, FileBufferReader supports
// incremental read operations from the data source.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://pkg.go.dev/bufio
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	Use the 'New' method to create new instances of
//		FileBufferReader.
//
//	(2)	FileBufferReader implements the io.Reader
//		interface.
type FileBufferReader struct {
	fileReader *bufio.Reader
	filePtr    *os.File

	lock *sync.Mutex
}

// Close
//
// This method is used to close any open file pointers.
//
// File pointers are closed automatically by
// FileBufferReader. However, in the case of processing
// errors or if there is any doubt about whether the
// file pointer is actually closed, simply call this
// method to ensure proper closure of an open file
// pointer.
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
func (fBufReader *FileBufferReader) Close() {

	if fBufReader.lock == nil {
		fBufReader.lock = new(sync.Mutex)
	}

	fBufReader.lock.Lock()

	defer fBufReader.lock.Unlock()

	if fBufReader.filePtr != nil {

		_ = fBufReader.filePtr.Close()
	}

	return
}

// New
//
// This method returns a fully initialized instance of
// FileBufferReader.
//
// The size of the internal read buffer is controlled by
// input parameter 'bufSize'. The minimum buffer size is
// 16-bytes.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	reader						io.Reader
//
//		An object which implements io.Reader interface.
//		Typically, this is a file pointer of type
//		*os.File.
//
//		A file pointer (*os.File) will facilitate reading
//		file data from files residing on an attached
//		storage drive.
//
//		While the returned instance of FileBufferReader
//		is primarily designed for reading file data, the
//		reader will in fact read data from any object
//		implementing the io.Reader interface.
//
//	bufSize						int
//
//		This integer value controls the size of the
//		buffer created for the returned instance of
//		FileBufferReader.
//
//		'bufSize' should be configured to maximize
//		performance for 'read' operations subject to
//		prevailing memory limitations.
//
//		The minimum reader buffer size is 16-bytes. If
//		'bufSize' is set to a size less than '16', an
//		error will be returned.
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
//	FileBufferReader
//
//		If this method completes successfully, a fully
//		configured instance of FileBufferReader will
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
func (fBufReader *FileBufferReader) New(
	reader io.Reader,
	bufSize int,
	errorPrefix interface{}) (
	FileBufferReader,
	error) {

	if fBufReader.lock == nil {
		fBufReader.lock = new(sync.Mutex)
	}

	fBufReader.lock.Lock()

	defer fBufReader.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error
	var newFileBufReader FileBufferReader

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileBufferReader."+
			"NewInitializeReader()",
		"")

	if err != nil {
		return newFileBufReader, err
	}

	if bufSize < 16 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'bufSize' is invalid!\n"+
			"The minimum buffer size is '16'.\n"+
			"Input parameter 'bufSize' = '%v'\n",
			ePrefix.String(),
			bufSize)

		return newFileBufReader, err
	}

	newFileBufReader.fileReader = bufio.NewReaderSize(
		reader,
		bufSize)

	return newFileBufReader, err
}

// NewPathFileName
//
// Receives a path and file name as an input parameter.
// This file is opened for 'read-only' operations and
// is configured in the returned instance
// FileBufferReader.
//
// The size of the internal read buffer is controlled by
// input parameter 'bufSize'. The minimum buffer size is
// 16-bytes.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	reader						io.Reader
//
//		An object which implements io.Reader interface.
//		Typically, this is a file pointer of type
//		*os.File.
//
//		A file pointer (*os.File) will facilitate reading
//		file data from files residing on an attached
//		storage drive.
//
//		While the returned instance of FileBufferReader
//		is primarily designed for reading file data, the
//		reader will in fact read data from any object
//		implementing the io.Reader interface.
//
//	bufSize						int
//
//		This integer value controls the size of the
//		buffer created for the returned instance of
//		FileBufferReader.
//
//		'bufSize' should be configured to maximize
//		performance for 'read' operations subject to
//		prevailing memory limitations.
//
//		The minimum reader buffer size is 16-bytes. If
//		'bufSize' is set to a size less than '16', an
//		error will be returned.
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
//	FileBufferReader
//
//		If this method completes successfully, a fully
//		configured instance of FileBufferReader will
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
func (fBufReader *FileBufferReader) NewPathFileName(
	pathFileName string,
	bufSize int,
	errorPrefix interface{}) (
	FileBufferReader,
	error) {

	if fBufReader.lock == nil {
		fBufReader.lock = new(sync.Mutex)
	}

	fBufReader.lock.Lock()

	defer fBufReader.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error
	var newFileBufReader FileBufferReader
	funcName := "FileBufferReader." +
		"NewPathFileName()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {
		return newFileBufReader, err
	}

	if bufSize < 16 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'bufSize' is invalid!\n"+
			"The minimum buffer size is '16'.\n"+
			"Input parameter 'bufSize' = '%v'\n",
			ePrefix.String(),
			bufSize)

		return newFileBufReader, err
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

		return newFileBufReader, err
	}

	if !pathFileDoesExist {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'pathFileName' is invalid!\n"+
			"The path and file name do NOT exist on an attached\n"+
			"storage drive.\n"+
			"pathFileName= '%v'\n",
			ePrefix.String(),
			pathFileName)

		return newFileBufReader, err
	}

	if fInfoPlus.IsDir() {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'pathFileName' is invalid!\n"+
			"'pathFileName' is directory and NOT a file name.\n"+
			"pathFileName= '%v'\n",
			ePrefix.String(),
			pathFileName)

		return newFileBufReader, err
	}

	var filePermissionCfg FilePermissionConfig

	filePermissionCfg,
		err = new(FilePermissionConfig).New(
		"-r--r--r--",
		ePrefix.XCpy("filePermissionCfg<-"))

	if err != nil {

		return newFileBufReader, err
	}

	var fileOpenCfg FileOpenConfig

	fileOpenCfg,
		err = new(FileOpenConfig).New(
		ePrefix.XCpy("fileOpenCfg<-"),
		FOpenType.TypeReadOnly())

	if err != nil {

		return newFileBufReader, err
	}

	newFileBufReader.filePtr,
		err = new(fileHelperBoson).
		openFile(
			pathFileName,
			false,
			fileOpenCfg,
			filePermissionCfg,
			"pathFileName",
			ePrefix)

	if err != nil {

		if newFileBufReader.filePtr != nil {
			_ = newFileBufReader.filePtr.Close()
		}

		return newFileBufReader, err
	}

	newFileBufReader.fileReader = bufio.NewReaderSize(
		newFileBufReader.filePtr,
		bufSize)

	return newFileBufReader, err
}

// Read
//
// Reads a selection data from the pre-configured data
// source associated with the current instance of
// FileBufferReader.
//
// This method is a wrapper for the bufio
// 'Reader.Read' method.
//
// Method 'Read' reads data into the byte array
// 'bytesRead'. It returns the number of bytes read
// into 'bytesRead' as return parameter 'numOfBytesRead'.
// The bytes are taken from at most one Read on the
// underlying Reader, hence 'numOfBytesRead' may be less
// than len(bytesRead).
//
// If the underlying Reader can return a non-zero count
// with io.EOF, then this Read method can do so as well
// (See the io.Reader docs and 'Reference' section
// below).
//
// The underlying data source will be automatically
// closed when an error occurs or when all bytes have
// been read from the data source.
//
// If there are any doubts about whether the data source
// has been 'closed', call method:
//
//	FileBufferReader.Close()
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	This method implements the io.Reader interface.
//
//	(2)	Keep calling this method until all the bytes have
//		been read from the data source configured for the
//		current instance of FileBufferReader.
//
//	(3)	Callers should always process the
//		numOfBytesRead > 0 bytes returned before
//		considering the error err. Doing so correctly
//		handles I/O errors that happen after reading some
//		bytes and also both of the allowed EOF behaviors.
//
//	(4)	When a processing error occurs (err != io.EOF)
//		during the 'read' operation, a file data source
//		will be automatically closed and the current
//		FileBufferReader instance will be no longer
//		operational.
//
//	(5)	When all bytes have been successfully read from
//		the data source, it will be automatically closed.
//
//	(6)	If there are any doubts about whether a file data
//		source has been 'closed', call method:
//
//			FileBufferReader.Close()
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	Read(p []byte) (n int, err error)
//	https://pkg.go.dev/bufio#Reader.Read
//
//	io.Reader
//	https://pkg.go.dev/io#Reader
//
//	Reader is the interface that wraps the basic Read
//	method.
//
//	Read reads up to len(p) bytes into p. It returns the
//	number of bytes read (0 <= n <= len(p)) and any error
//	encountered. Even if Read returns n < len(p), it may
//	use all of p as scratch space during the call. If some
//	data is available but not len(p) bytes, Read
//	conventionally returns what is available instead of
//	waiting for more.
//
//	When Read encounters an error or end-of-file
//	condition after successfully reading n > 0 bytes, it
//	returns the number of bytes read. It may return the
//	(non-nil) error from the same call or return the
//	error (and n == 0) from a subsequent call. An
//	instance of this general case is that a Reader
//	returning a non-zero number of bytes at the end of
//	the input stream may return either err == EOF or
//	err == nil. The next Read should return 0, EOF.
//
//	Callers should always process the n > 0 bytes returned
//	before considering the error err. Doing so correctly
//	handles I/O errors that happen after reading some bytes
//	and also both of the allowed EOF behaviors.
//
//	Implementations of Read are discouraged from returning a
//	zero byte count with a nil error, except when
//	len(p) == 0. Callers should treat a return of 0 and nil
//	as indicating that nothing happened; in particular it
//	does not indicate EOF.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	bytesRead					[]byte
//
//		If the length of this byte array is less than
//		16-bytes, an error will be returned.
//
//		Bytes will be read from the data source
//		configured for the current instance of
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	numOfBytesRead				int
//
//		If this method completes successfully, the number
//		of bytes read from the data source, and stored in
//		the byte array passed as input parameter
//		'bytesRead', will be returned through this
//		parameter.
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
func (fBufReader *FileBufferReader) Read(
	bytesRead []byte) (
	numOfBytesRead int,
	err error) {

	if fBufReader.lock == nil {
		fBufReader.lock = new(sync.Mutex)
	}

	fBufReader.lock.Lock()

	defer fBufReader.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"FileBufferReader."+
			"Read()",
		"")

	if err != nil {

		return numOfBytesRead, err
	}

	if len(bytesRead) < 16 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'bytesRead' is invalid!\n"+
			"The minimum byte array size is '16'.\n"+
			"The length of input parameter 'bytesRead' = '%v'\n",
			ePrefix.String(),
			bytesRead)

		return numOfBytesRead, err
	}

	if fBufReader.fileReader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of 'FileBufferReader' is invalid!\n"+
			"The internal bufio.Reader has NOT been initialized.\n"+
			"Call one of the 'New' methods when creating an instance\n"+
			"of 'FileBufferReader'\n",
			ePrefix.String())

		return numOfBytesRead, err
	}

	var err2 error

	numOfBytesRead,
		err2 = fBufReader.fileReader.Read(bytesRead)

	if err2 != nil {

		if err2 != io.EOF {

			err = fmt.Errorf("%v\n"+
				"Error returned by fBufReader.fileReader.Read(bytesRead).\n"+
				"Error=\n%v\n",
				ePrefix.String(),
				err2.Error())

		}

	}

	if err != nil {

		if fBufReader.filePtr != nil {

			_ = fBufReader.filePtr.Close()
		}
	}

	return numOfBytesRead, err
}
