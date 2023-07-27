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
// to facilitate data 'read' operations. The most common
// data source for these read operations is assumed to be
// a data file residing on an attached storage drive.
// However, any object implementing the io.Reader
// interface may be used as a data source.
//
// The FileBufferReader type is a wrapper for
// 'bufio.Reader'. As such, FileBufferReader supports
// incremental or buffered read operations from the target
// data source.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://pkg.go.dev/bufio
//	https://pkg.go.dev/bufio#Reader
//	https://pkg.go.dev/io#Reader
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	Use methods 'New' and 'NewPathFileName' to create
//		new instances of FileBufferReader.
//
//	(2)	FileBufferReader implements the io.Reader
//		interface.
//
// ----------------------------------------------------------------
//
// # Best Practice
//
//	(1)	Create a new instance of FileBufferReader using
//		either the New() method or the NewPathFileName()
//		method.
//
//		(a)	The New() method is used when an instance of
//			io.Reader is created externally by the user
//			and passed to the FileBufferReader.New()
//			method.
//
//			Under this scenario, the user is independently
//			responsible for clean-up of the io.Reader
//			object after FileBufferReader 'read'
//			operations have been completed.
//
//			Once all FileBufferReader 'read' operations
//			have been completed, call method Close() to
//			perform local FileBufferReader clean-up tasks.
//
//		(b)	The NewPathFileName() method allows for the
//			creation of an internal file pointer to a
//			file passed as a path and file name by the
//			user. This file serves at the target
//			io.Reader object from which data will be
//			read.
//
//			Under this scenario, the user simply calls
//			method Close() to perform all required
//			clean-up tasks after 'read' operations have
//			been completed.
//
//			Once method Close() is called, the current
//			FileBufferReader instance becomes unusable
//			and should be discarded.
//
//	(2)	After creating an instance of FileBufferReader,
//		the user calls the Read() method to read bytes
//		of data from the target io.Reader object. This
//		'read' target may be a file or any other object
//		which implements the io.Reader interface.
//
//		The Read() method should be called repeatedly
//		until all data has been read from the underlying
//		io.Reader object.
//
//		Upon completion of the 'read' operation, call
//		method Close() to perform required clean-up
//		tasks.
//
//	(3)	After all data bytes have been read from the
//		target io.Reader object, the user must call
//		method Close() to perform necessary clean-up
//		tasks.
//
//		Once method Close() is called, the current
//		FileBufferReader instance becomes unusable and
//		should be discarded.
type FileBufferReader struct {
	fileReader         *bufio.Reader
	filePtr            *os.File
	targetReadFileName string

	lock *sync.Mutex
}

// Close
//
// This method is used to close any open file pointers
// and perform required clean-up operations.
//
// Users MUST call this method after all 'read'
// operations have been completed.
//
// After calling this method, the current instance of
// FileBufferReader will be unusable and should be
// discarded.
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
func (fBufReader *FileBufferReader) Close(
	errorPrefix interface{}) error {

	if fBufReader.lock == nil {
		fBufReader.lock = new(sync.Mutex)
	}

	fBufReader.lock.Lock()

	defer fBufReader.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileBufferWriter."+
			"Close()",
		"")

	if err != nil {
		return err
	}

	err = new(fileBufferReaderMolecule).close(
		fBufReader,
		"fBufReader",
		ePrefix.XCpy("fBufReader"))

	return err
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
// # Reference:
//
//	https://pkg.go.dev/bufio
//	https://pkg.go.dev/bufio#Reader
//	https://pkg.go.dev/io#Reader
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	reader						io.Reader
//
//		An object which implements io.Reader interface.
//
//		This object may be a file pointer of type *os.File.
//		File pointers of this type implement the io.Reader
//		interface.
//
//		A file pointer (*os.File) will facilitate reading
//		data from files residing on an attached storage
//		drive. However, with this configuration, the user
//		is responsible for manually closing the file and
//		performing any other required clean-up operations
//		in addition to calling FileBufferReader.Close().
//
//		While the returned instance of FileBufferReader
//		is primarily designed for reading data from disk
//		files, this 'reader' will in fact read data from
//		any object implementing the io.Reader interface.
//
//	bufSize						int
//
//		This integer value controls the size of the
//		'read' buffer created for the returned instance
//		of FileBufferReader.
//
//		'bufSize' should be configured to maximize
//		performance for 'read' operations subject to
//		prevailing memory limitations.
//
//		The minimum reader buffer size is 16-bytes. If
//		'bufSize' is set to a size less than "16", it
//		will be automatically set to the default buffer
//		size of 4096-bytes.
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
			"New()",
		"")

	if err != nil {
		return newFileBufReader, err
	}

	err = new(fileBufferReaderNanobot).
		setIoReader(
			&newFileBufReader,
			"newFileBufReader",
			reader,
			"reader",
			bufSize,
			ePrefix.XCpy("newFileBufReader"))

	return newFileBufReader, err
}

// NewPathFileName
//
// Receives a path and file name as an input parameter.
// This file is opened for 'read' operations and is
// configured internally as a bufio.Reader contained in
// the returned instance of FileBufferReader.
//
// The path and file name will be used to create a file
// pointer (*os.File) which in turn will be used to
// configure the internal bufio.Reader.
//
// The size of the internal read buffer is controlled by
// input parameter 'bufSize'. The minimum buffer size is
// 16-bytes.
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
//	openFileReadWrite			bool
//
//		If this parameter is set to 'true', the target
//		'read' file identified from input parameter
//		'pathFileName' will be opened for both 'read'
//		and 'write' operations.
//
//		If 'openFileReadWrite' is set to 'false', the
//		target 'read' file will be opened for 'read-only'
//		operations.
//
//	bufSize						int
//
//		This integer value controls the size of the
//		'read' buffer created for the returned instance
//		of FileBufferReader.
//
//		'bufSize' should be configured to maximize
//		performance for 'read' operations subject to
//		prevailing memory limitations.
//
//		The minimum reader buffer size is 16-bytes. If
//		'bufSize' is set to a size less than "16", it
//		will be automatically set to the default buffer
//		size of 4096.
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
	openFileReadWrite bool,
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
			"NewPathFileName()",
		"")

	if err != nil {
		return newFileBufReader, err
	}

	err = new(fileBufferReaderNanobot).
		setPathFileName(
			&newFileBufReader,
			"newFileBufReader",
			pathFileName,
			"pathFileName",
			openFileReadWrite,
			bufSize,
			ePrefix.XCpy(
				pathFileName))

	return newFileBufReader, err
}

// Read
//
// Reads a selection data from the pre-configured data
// source associated with the current instance of
// FileBufferReader.
//
// This method is a wrapper for the bufio 'Reader.Read'
// method.
//
// Method 'Read' reads data into the byte array
// 'bytesRead'. It returns the number of bytes read
// into the byte array as return parameter,
// 'numOfBytesRead'.
//
// Under certain circumstances, the number of bytes read
// may be less than the length of the byte array
// (len(bytesRead)) due to the length of the underlying
// read buffer.
//
// To complete the read operation, repeat the call to
// this method until the returned error is set to
// 'io.EOF' signaling 'End of File'.
//
// See the io.Reader docs and 'Reference' section below.
//
// Once the 'read' operation has been completed, the user
// MUST call the 'Close' method to ensure clean-up
// operations are properly applied:
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
//		current instance of FileBufferReader and the
//		returned error is set to 'io.EOF'.
//
//	(3)	Callers should always process the
//		numOfBytesRead > 0 bytes returned before
//		considering the error err. Doing so correctly
//		handles I/O errors that happen after reading some
//		bytes and also both of the allowed EOF behaviors
//		(See the io.Reader docs and 'Reference' section
//		below).
//
//	(4)	When all 'read' operations have been completed,
//		call method:
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
//		FileBufferReader.
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

		} else {

			err = io.EOF

		}

	}

	return numOfBytesRead, err
}

// SetReader
//
// This method will completely re-initialize the current
// instance of FileBufferReader using the io.Reader object
// passed as input parameter 'reader'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the current instance of
//	FileBufferReader.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://pkg.go.dev/bufio
//	https://pkg.go.dev/bufio#Reader
//	https://pkg.go.dev/io#Reader
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	reader						io.Reader
//
//		An object which implements io.Reader interface.
//
//		This object may be a file pointer of type *os.File.
//		File pointers of this type implement the io.Reader
//		interface.
//
//		A file pointer (*os.File) will facilitate reading
//		data from files residing on an attached storage
//		drive. However, with this configuration, the user
//		is responsible for manually closing the file and
//		performing any other required clean-up operations
//		in addition to calling FileBufferReader.Close().
//
//		While the returned instance of FileBufferReader
//		is primarily designed for reading data from disk
//		files, this 'reader' will in fact read data from
//		any object implementing the io.Reader interface.
//
//	bufSize						int
//
//		This integer value controls the size of the
//		'read' buffer created for the returned instance
//		of FileBufferReader.
//
//		'bufSize' should be configured to maximize
//		performance for 'read' operations subject to
//		prevailing memory limitations.
//
//		The minimum reader buffer size is 16-bytes. If
//		'bufSize' is set to a size less than "16", it
//		will be automatically set to the default buffer
//		size of 4096-bytes.
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
func (fBufReader *FileBufferReader) SetReader(
	reader io.Reader,
	bufSize int,
	errorPrefix interface{}) error {

	if fBufReader.lock == nil {
		fBufReader.lock = new(sync.Mutex)
	}

	fBufReader.lock.Lock()

	defer fBufReader.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileBufferReader."+
			"SetReader()",
		"")

	if err != nil {
		return err
	}

	err = new(fileBufferReaderNanobot).
		setIoReader(
			fBufReader,
			"fBufReader",
			reader,
			"reader",
			bufSize,
			ePrefix.XCpy("fBufReader"))

	return err
}

// SetPathFileName
//
// This method will completely re-initialize the current
// instance of FileBufferReader using the path and file
// name passed as input parameter 'pathFileName'.
//
// The path and file name will be used to create a file
// pointer (*os.File) which in turn will be used to
// configure the internal bufio.Reader for the current
// instance of FileBufferReader.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the current instance of
//	FileBufferReader.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://pkg.go.dev/bufio
//	https://pkg.go.dev/bufio#Reader
//	https://pkg.go.dev/io#Reader
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
//	openFileReadWrite			bool
//
//		If this parameter is set to 'true', the target
//		'read' file identified from input parameter
//		'pathFileName' will be opened for both 'read'
//		and 'write' operations.
//
//		If 'openFileReadWrite' is set to 'false', the
//		target 'read' file will be opened for 'read-only'
//		operations.
//
//	bufSize						int
//
//		This integer value controls the size of the
//		'read' buffer created for the returned instance
//		of FileBufferReader.
//
//		'bufSize' should be configured to maximize
//		performance for 'read' operations subject to
//		prevailing memory limitations.
//
//		The minimum reader buffer size is 16-bytes. If
//		'bufSize' is set to a size less than "16", it
//		will be automatically set to the default buffer
//		size of 4096.
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
func (fBufReader *FileBufferReader) SetPathFileName(
	pathFileName string,
	openFileReadWrite bool,
	bufSize int,
	errorPrefix interface{}) error {

	if fBufReader.lock == nil {
		fBufReader.lock = new(sync.Mutex)
	}

	fBufReader.lock.Lock()

	defer fBufReader.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileBufferReader."+
			"NewPathFileName()",
		"")

	if err != nil {
		return err
	}

	err = new(fileBufferReaderNanobot).
		setPathFileName(
			fBufReader,
			"fBufReader",
			pathFileName,
			"pathFileName",
			openFileReadWrite,
			bufSize,
			ePrefix.XCpy(
				pathFileName))

	return err
}

type fileBufferReaderNanobot struct {
	lock *sync.Mutex
}

// setIoReader
//
// This 'setter' method is used to initialize new values
// for internal member variables contained in the
// instance of FileBufferReader passed as input parameter
// 'fBufReader'. The new configuration will be based on
// an io.Reader object passed as input parameter
// 'reader'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the instance of
//	FileBufferReader passed as input parameter
//	'fBufReader'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fBufReader					*FileBufferReader
//
//		A pointer to an instance of FileBufferReader.
//
//		All internal member variable data values in
//		this instance will be deleted and initialized
//		to new values based on the following input
//		parameters.
//
//	fBufReaderLabel				string
//
//		The name or label associated with input parameter
//		'fBufReader' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fBufReader" will be
//		automatically applied.
//
//	reader						io.Reader
//
//		This parameter will accept any object
//		implementing the io.Reader interface.
//
//	readerLabel					string
//
//		The name or label associated with input parameter
//		'reader' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "reader" will be
//		automatically applied.
//
//	bufSize						int
//
//		This integer value controls the size of the
//		'read' buffer created for the returned instance
//		of FileBufferReader.
//
//		'bufSize' should be configured to maximize
//		performance for 'read' operations subject to
//		prevailing memory limitations.
//
//		The minimum reader buffer size is 16-bytes. If
//		'bufSize' is set to a size less than "16", it
//		will be automatically set to the default buffer
//		size of 4096-bytes.
func (fBufReaderNanobot *fileBufferReaderNanobot) setIoReader(
	fBufReader *FileBufferReader,
	fBufReaderLabel string,
	reader io.Reader,
	readerLabel string,
	bufSize int,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fBufReaderNanobot.lock == nil {
		fBufReaderNanobot.lock = new(sync.Mutex)
	}

	fBufReaderNanobot.lock.Lock()

	defer fBufReaderNanobot.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileBufferReaderNanobot." +
		"setIoReader()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if len(fBufReaderLabel) == 0 {

		fBufReaderLabel = "fBufReader"
	}

	if len(readerLabel) == 0 {

		readerLabel = "reader"
	}

	if fBufReader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: The FileBufferReader instance passed\n"+
			"as input parameter '%v' is invalid!\n"+
			"'%v' is a 'nil' pointer.\n",
			ePrefix.String(),
			fBufReaderLabel,
			fBufReaderLabel)

		return err
	}

	err = new(fileBufferReaderMolecule).close(
		fBufReader,
		fBufReaderLabel,
		ePrefix.XCpy(fBufReaderLabel))

	if err != nil {
		return err
	}

	if bufSize < 16 {

		bufSize = 4096
	}

	fBufReader.fileReader = bufio.NewReaderSize(
		reader,
		bufSize)

	return err
}

// setPathFileName
//
// This 'setter' method is used to initialize new values
// for internal member variables contained in the
// instance of FileBufferReader passed as input parameter
// 'fBufReader'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the instance of
//	FileBufferReader passed as input parameter
//	'fBufReader'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fBufReader					*FileBufferReader
//
//		A pointer to an instance of FileBufferReader.
//
//		All internal member variable data values in
//		this instance will be deleted and initialized
//		to new values based on the following input
//		parameters.
//
//	fBufReaderLabel				string
//
//		The name or label associated with input parameter
//		'fBufReader' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fBufReader" will be
//		automatically applied.
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
//	pathFileNameLabel			string
//
//		The name or label associated with input parameter
//		'pathFileName' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "pathFileName" will be
//		automatically applied.
//
//	openFileReadWrite			bool
//
//		If this parameter is set to 'true', the target
//		'read' file identified from input parameter
//		'pathFileName' will be opened for both 'read'
//		and 'write' operations.
//
//		If 'openFileReadWrite' is set to 'false', the
//		target 'read' file will be opened for 'read-only'
//		operations.
//
//	bufSize						int
//
//		This integer value controls the size of the
//		'read' buffer created for the returned instance
//		of FileBufferReader.
//
//		'bufSize' should be configured to maximize
//		performance for 'read' operations subject to
//		prevailing memory limitations.
//
//		The minimum reader buffer size is 16-bytes. If
//		'bufSize' is set to a size less than "16", it
//		will be automatically reset to the default buffer
//		size of 4096.
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
func (fBufReaderNanobot *fileBufferReaderNanobot) setPathFileName(
	fBufReader *FileBufferReader,
	fBufReaderLabel string,
	pathFileName string,
	pathFileNameLabel string,
	openFileReadWrite bool,
	bufSize int,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fBufReaderNanobot.lock == nil {
		fBufReaderNanobot.lock = new(sync.Mutex)
	}

	fBufReaderNanobot.lock.Lock()

	defer fBufReaderNanobot.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileBufferReaderNanobot." +
		"setPathFileName()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return err
	}

	if len(fBufReaderLabel) == 0 {

		fBufReaderLabel = "fBufReader"
	}

	if fBufReader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: The FileBufferReader instance passed\n"+
			"as input parameter '%v' is invalid!\n"+
			"'%v' is a 'nil' pointer.\n",
			ePrefix.String(),
			fBufReaderLabel,
			fBufReaderLabel)

		return err
	}

	if len(pathFileNameLabel) == 0 {

		pathFileNameLabel = "pathFileName"
	}

	if len(pathFileName) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is invalid!\n"+
			"'%v' is an empty string with a length of zero (0).\n",
			ePrefix.String(),
			pathFileNameLabel,
			pathFileNameLabel)

		return err
	}

	err = new(fileBufferReaderMolecule).close(
		fBufReader,
		fBufReaderLabel,
		ePrefix.XCpy(fBufReaderLabel))

	if err != nil {
		return err
	}

	if bufSize < 16 {

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
				pathFileNameLabel)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"An error occurred while testing for the existance\n"+
			"of '%v' on an attached storage drive.\n"+
			"%v = '%v'\n"+
			"Error= \n%v\n",
			funcName,
			pathFileNameLabel,
			pathFileNameLabel,
			pathFileName,
			err2.Error())

		return err
	}

	if !pathFileDoesExist {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is invalid!\n"+
			"The path and file name do NOT exist on an attached\n"+
			"storage drive.\n"+
			"%v= '%v'\n",
			ePrefix.String(),
			pathFileNameLabel,
			pathFileNameLabel,
			pathFileName)

		return err
	}

	if fInfoPlus.IsDir() {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'pathFileName' is invalid!\n"+
			"'pathFileName' is directory and NOT a file name.\n"+
			"pathFileName= '%v'\n",
			ePrefix.String(),
			pathFileName)

		return err
	}

	var filePermissionCfg FilePermissionConfig
	var fileOpenPermissions = "-r--r--r--"

	if openFileReadWrite == true {
		fileOpenPermissions = "-rw-rw-rw-"
	}

	filePermissionCfg,
		err = new(FilePermissionConfig).New(
		fileOpenPermissions,
		ePrefix.XCpy("filePermissionCfg<-"))

	if err != nil {

		return err
	}

	var fileOpenCfg FileOpenConfig

	fileOpenCfg,
		err = new(FileOpenConfig).New(
		ePrefix.XCpy("fileOpenCfg<-"),
		FOpenType.TypeReadOnly())

	if err != nil {

		return err
	}

	fBufReader.filePtr,
		err = new(fileHelperBoson).
		openFile(
			pathFileName,
			false,
			fileOpenCfg,
			filePermissionCfg,
			pathFileNameLabel,
			ePrefix)

	if err != nil {

		return err
	}

	fBufReader.targetReadFileName = pathFileName

	fBufReader.fileReader = bufio.NewReaderSize(
		fBufReader.filePtr,
		bufSize)

	return err
}

type fileBufferReaderMolecule struct {
	lock *sync.Mutex
}

// close
//
// This method will effectively close and invalidate the
// instance of FileBufferReader passed as input parameter
// 'fBufReader'.
//
// If a file pointer (*os.File) was previously configured
// for 'fBufReader', it will be closed and set to 'nil'.
//
// After completion of this method this FileBufferReader
// instance will be unusable, invalid and unavailable for
// future 'read' operations.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete all pre-existing data values
//	in the instance of FileBufferReader passed as input
//	parameter 'fBufReader'.
//
//	After completion of this method this FileBufferReader
//	instance will be unusable, invalid and unavailable
//	for future 'read' operations.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fBufReader					*FileBufferReader
//
//		A pointer to an instance of FileBufferReader.
//
//		All internal member variable data values in
//		this instance will be deleted.
//
//		If a file pointer (*os.File) was previously
//		configured for 'fBufReader', it will be closed
//		and set to 'nil' by this method.
//
//	fBufReaderLabel				string
//
//		The name or label associated with input parameter
//		'fBufReader' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fBufReader" will be
//		automatically applied.
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
func (fBuffReaderMolecule *fileBufferReaderMolecule) close(
	fBufReader *FileBufferReader,
	fBufReaderLabel string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fBuffReaderMolecule.lock == nil {
		fBuffReaderMolecule.lock = new(sync.Mutex)
	}

	fBuffReaderMolecule.lock.Lock()

	defer fBuffReaderMolecule.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileBufferReaderMolecule." +
		"close()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return err
	}

	if len(fBufReaderLabel) == 0 {

		fBufReaderLabel = "fBufReader"
	}

	if fBufReader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: The FileBufferReader instance passed\n"+
			"as input parameter '%v' is invalid!\n"+
			"'%v' is a 'nil' pointer.\n",
			ePrefix.String(),
			fBufReaderLabel,
			fBufReaderLabel)

		return err
	}

	if fBufReader.filePtr != nil {

		var err2 error

		err2 = fBufReader.filePtr.Close()

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Error returned while closing the target 'target' file!\n"+
				"fBufWriter.filePtr.Close()\n"+
				"Target Read File = '%v'\n"+
				"Error = \n%v\n",
				ePrefix.String(),
				fBufReader.targetReadFileName,
				err2.Error())

		}
	}

	fBufReader.targetReadFileName = ""

	fBufReader.filePtr = nil

	fBufReader.fileReader = nil

	return err
}
