package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"sync"
)

// FileBufferReadWrite
//
// This type implements both the io.Reader and
// io.Writer interfaces. It is designed to provide a
// read and write capability for files and any objects
// supporting the io.Reader or io.Writer interfaces.
//
// Users can employ this type to perform read and
// write operations in a single method call.
//
// Read and Write operations are performed using
// private, internal FileBufferReader and
// FileBufferWriter objects.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://pkg.go.dev/bufio
//	https://pkg.go.dev/bufio#Reader
//	https://pkg.go.dev/io#Reader
//	https://pkg.go.dev/bufio
//	https://pkg.go.dev/bufio#Writer
//	https://pkg.go.dev/io#Writer
//
// ----------------------------------------------------------------
//
// # Best Practices
//
//	Initialization
//
//	Instances of FileBufferReadWrite are
//	created using one of three types of
//	input parameters:
//
//		(1)	io.Reader and io.Writer objects.
//
//		(2) File Manager objects (FileMgr).
//
//		(3) String parameters containing the
//			paths and file names for files which
//			will be converted to io.Reader and
//			io.Writer objects.
//
//	There are two options for creating and
//	initializing a new, valid instance of
//	FileBufferReadWrite.
//
//		Option-1
//
//		Use one of the following 'New' methods to
//		create a valid instance of FileBufferReadWrite
//		by configuring both the internal io.Reader and
//		the io.Writer simultaneously:
//
//		FileBufferReadWrite.NewIoReadWrite()
//		FileBufferReadWrite.NewFileMgrs()
//		FileBufferReadWrite.NewPathFileNames()
//
//		Option-2
//
//		This second option calls for configuring the
//		internal io.Reader and io.Writer objects
//		separately. This approach allows the user to
//		choose different types of input parameters for
//		configuring the internal io.Reader and io.Writer.
//		This approach relies on calling 'Setter' methods.
//
//		(a)	Call the 'New' method to generate a blank
//			or empty instance of FileBufferReadWrite.
//
//		(b) Next, call any appropriate combination of
//			'Setter' methods to finalize the configuration
//			of a valid FileBufferReadWrite instance:
//
//		FileBufferReadWrite.SetFileMgrReadWrite()
//		FileBufferReadWrite.SetFileMgrReader()
//		FileBufferReadWrite.SetFileMgrWriter()
//		FileBufferReadWrite.SetIoReadWrite()
//		FileBufferReadWrite.SetIoReader()
//		FileBufferReadWrite.SetIoWriter()
//		FileBufferReadWrite.SetPathFileNamesReadWrite()
//		FileBufferReadWrite.SetPathFileNameReader()
//		FileBufferReadWrite.SetFileMgrReadWrite()
//		FileBufferReadWrite.SetPathFileNameWriter()
//
//
//	Usage
//
//	There are three methods which are designed to
//	provide 'read' and 'write' services:
//
//		FileBufferReadWrite.ReadWriteAll()
//		FileBufferReadWrite.Read()
//		FileBufferReadWrite.Write()
//
//
//	Clean-Up
//
//	When all 'read' and 'write' operations have been
//	completed, the user is responsible for performing
//	Clean-Up operations by calling the following method:
//
//		FileBufferReadWrite.CloseFileBufferReadWrite()
//
//	NOTE -	This 'Close' method will also 'flush' the
//			internal 'write' buffer.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	When all 'read' and 'write' operations have been
//	completed, the user is responsible for performing
//	Clean-Up operations by calling the following method:
//
//		FileBufferReadWrite.CloseFileBufferReadWrite()
type FileBufferReadWrite struct {
	writer             *FileBufferWriter
	reader             *FileBufferReader
	writerFilePathName string
	readerFilePathName string

	lock *sync.Mutex
}

// CloseFileBufferReadWrite
//
// This method is designed to perform clean up tasks
// after completion of all 'read' and 'write' operations
// associated with the current instance of
// FileBufferReadWrite.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	After completing all 'read' and 'write' operations,
//	users MUST call this method in order to perform
//	required clean-up operations.
//
//	This method will:
//
//	(1)	'Flush' the write buffer thereby ensuring all
//		data is written from the write buffer to the
//		underlying io.Writer object.
//
//	(2) Properly 'Close' the 'write' file or io.Writer
//		object.
//
//	(3) Properly 'Close' the 'read' file or io.Reader
//		object.
//
//	(4) Effectively render the current instance of
//		FileBufferReadWrite invalid and unusable for
//		any future 'read' or 'write' operations.
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
func (fBufReadWrite *FileBufferReadWrite) CloseFileBufferReadWrite(
	errorPrefix interface{}) error {

	if fBufReadWrite.lock == nil {
		fBufReadWrite.lock = new(sync.Mutex)
	}

	fBufReadWrite.lock.Lock()

	defer fBufReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileBufferReadWrite."+
			"CloseFileBufferReadWrite()",
		"")

	if err != nil {
		return err
	}

	err = new(fileBufferReadWriteMicrobot).
		closeReaderWriter(
			fBufReadWrite,
			"fBufReadWrite",
			ePrefix.XCpy("Close-Readers&Writers"))

	return err
}

// CloseReader
//
// This method is designed to perform clean-up tasks
// after completion of all 'read' operations associated
// with the current instance of FileBufferReadWrite.
//
// After calling this method, the clean-up tasks
// performed will effectively render the internal
// io.Reader object, encapsulated by the current
// FileBufferReadWrite instance, invalid and unusable
// for any future 'read' operations.
//
// It is unlikely that the user will ever need to call
// this method. Typically, clean-up tasks are performed
// jointly on the internal io.Reader and io.Writer
// objects encapsulated in the current FileBufferReadWrite
// instance. These clean-up tasks should be performed
// after all 'read' and 'write' operations have been
// completed by calling the local method:
//
//	FileBufferReadWrite.CloseFileBufferReadWrite()
//
// However, in the event of unforeseen use cases, this
// method is provided to exclusively close or clean-up
// the io.Reader.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will:
//
//	(1) Properly 'Close' the 'read' file or internal
//		io.Reader object.
//
//	(2) Effectively render the internal io.Reader object,
//		encapsulated by the current instance of
//		FileBufferReadWrite, invalid and unusable for any
//		future 'read' operations.
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
func (fBufReadWrite *FileBufferReadWrite) CloseReader(
	errorPrefix interface{}) error {

	if fBufReadWrite.lock == nil {
		fBufReadWrite.lock = new(sync.Mutex)
	}

	fBufReadWrite.lock.Lock()

	defer fBufReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileBufferReadWrite."+
			"CloseReader()",
		"")

	if err != nil {
		return err
	}

	err = new(fileBufferReadWriteElectron).
		closeReader(
			fBufReadWrite,
			"fBufReadWrite",
			ePrefix)

	return err
}

// CloseWriter
//
// This method is designed to perform clean-up tasks
// after completion of all 'write' operations associated
// with the current instance of FileBufferReadWrite.
//
// These clean-up tasks include:
//
//	(a)	Flushing the write buffer to ensure all data
//		is written from the write buffer to the
//		underlying io.Writer object.
//
//	(b) Closing the 'write' file or io.Writer object.
//
// After calling this method, the clean-up tasks
// performed will effectively render the internal
// io.Writer object, encapsulated by the current
// FileBufferReadWrite instance, invalid and unusable
// for any future 'write' operations.
//
// It is unlikely that the user will ever need to call
// this method. Typically, clean-up tasks are performed
// jointly on the internal io.Reader and io.Writer
// objects encapsulated in the current FileBufferReadWrite
// instance. These clean-up tasks should be performed
// after all 'read' and 'write' operations have been
// completed by calling the local method:
//
//	FileBufferReadWrite.CloseFileBufferReadWrite()
//
// However, in the event of unforeseen use cases, this
// method is provided to exclusively close or clean-up
// the io.Writer.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will:
//
//	(1)	'Flush' the write buffer thereby ensuring all
//		data is written from the write buffer to the
//		underlying io.Writer object.
//
//	(2) Properly 'Close' the 'write' file or io.Writer
//		object.
//
//	(3) Effectively render the internal io.Writer object,
//		encapsulated by the current instance of
//		FileBufferReadWrite, invalid and unusable for any
//		future 'write' operations.
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
func (fBufReadWrite *FileBufferReadWrite) CloseWriter(
	errorPrefix interface{}) error {

	if fBufReadWrite.lock == nil {
		fBufReadWrite.lock = new(sync.Mutex)
	}

	fBufReadWrite.lock.Lock()

	defer fBufReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileBufferReadWrite."+
			"CloseWriter()",
		"")

	if err != nil {
		return err
	}

	err = new(fileBufferReadWriteElectron).
		flushAndCloseWriter(
			fBufReadWrite,
			"fBufReadWrite",
			ePrefix)

	return err
}

// FlushWriteBuffer
//
// This method will flush the write buffer to ensure that
// all data is written to the underlying output destination
// (a.k.a. destination file).
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
func (fBufReadWrite *FileBufferReadWrite) FlushWriteBuffer(
	errorPrefix interface{}) error {

	if fBufReadWrite.lock == nil {
		fBufReadWrite.lock = new(sync.Mutex)
	}

	fBufReadWrite.lock.Lock()

	defer fBufReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileBufferReadWrite."+
			"FlushWriteBuffer()",
		"")

	if fBufReadWrite.writer == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: The current instance of FileBufferReadWrite\n"+
			"is invalid! The internal io.Writer object was never\n"+
			"initialized. Call one of the 'New' methods or 'Setter'\n"+
			"methods to create a valid instance of FileBufferReadWrite.\n",
			ePrefix.String())

		return err
	}

	return fBufReadWrite.writer.
		Flush(ePrefix.XCpy("fBufReadWrite.writer"))
}

// New
//
// This method returns an empty or 'blank' instance of
// FileBufferReadWrite. All the member variables in this
// returned instance are initialized to their zero or
// initial values. This means the returned instance is
// invalid and unusable for standard 'read' and 'write'
// operations.
//
// This technique for creating a new working instance of
// FileBufferReadWrite requires two steps.
//
// Step-1
//
//	Call this method FileBufferReadWrite.New() to
//	generate an empty version of FileBufferReadWrite.
//
// Step-2
//
//	Use this returned instance of FileBufferReadWrite and
//	call the appropriate 'Setter' methods to individually
//	configure the internal 'reader' and 'writer' objects.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	-- NONE --
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	FileBufferReadWrite
//
//		This method returns an empty instance of
//		FileBufferReadWrite. After receiving this
//		instance, users must call 'Setter' methods
//		to complete the 'reader' and 'writer'
//		configuration process.
func (fBufReadWrite *FileBufferReadWrite) New() FileBufferReadWrite {

	if fBufReadWrite.lock == nil {
		fBufReadWrite.lock = new(sync.Mutex)
	}

	fBufReadWrite.lock.Lock()

	defer fBufReadWrite.lock.Unlock()

	return FileBufferReadWrite{}
}

// NewIoReadWrite
//
// Creates and returns a new, fully configured instance
// of FileBufferReadWrite using io.Reader and io.Writer
// input parameters.
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
//		in addition to calling local method:
//
//		FileBufferReadWrite.CloseFileBufferReadWrite()
//
//		While the 'read' services provided by
//		FileBufferReadWrite are primarily designed to
//		read data from disk files, this type of 'reader'
//		will in fact read data from any object
//		implementing the io.Reader interface.
//
//	readerBuffSize				int
//
//		This integer value controls the size of the
//		buffer created for the io.Reader object passed as
//		input parameter 'reader'.
//
//		'readerBuffSize' should be configured to maximize
//		performance for 'read' operations subject to
//		prevailing memory limitations.
//
//		The minimum reader buffer size is 16-bytes. If
//		'readerBuffSize' is set to a value less than
//		"16", it will be automatically reset to the
//		default buffer size of 4096-bytes.
//
//	writer						io.Writer
//
//		This parameter will accept any object
//		implementing the io.Writer interface.
//
//		This object may be a file pointer of type *os.File.
//		File pointers of this type implement the io.Writer
//		interface.
//
//		A file pointer (*os.File) will facilitate writing
//		data to files residing on an attached storage
//		drive. However, with this configuration, the user
//		is responsible for manually closing the file and
//		performing any other required clean-up operations
//		in addition to calling local method:
//
//		FileBufferReadWrite.CloseFileBufferReadWrite()
//
//		While the 'write' services provided by the
//		FileBufferReadWrite are primarily designed for
//		writing data to disk files, this type of 'writer'
//		will in fact write data to any object
//		implementing the io.Writer interface.
//
//	writerBuffSize				int
//
//		This integer value controls the size of the
//		buffer created for the io.Writer object passed as
//		input parameter 'writer'.
//
//		'writerBuffSize' should be configured to maximize
//		performance for 'write' operations subject to
//		prevailing memory limitations.
//
//		The minimum write buffer size is 1-byte. If
//		'writerBuffSize' is set to a size less than or
//		equal to zero, it will be automatically reset to
//		the default buffer size of 4096-bytes.
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
//	FileBufferReadWrite
//
//		If this method completes successfully, it will
//		return a fully configured instance of
//		FileBufferReadWrite.
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
func (fBufReadWrite *FileBufferReadWrite) NewIoReadWrite(
	reader io.Reader,
	readerBuffSize int,
	writer io.Writer,
	writerBuffSize int,
	errorPrefix interface{}) (
	FileBufferReadWrite,
	error) {

	if fBufReadWrite.lock == nil {
		fBufReadWrite.lock = new(sync.Mutex)
	}

	fBufReadWrite.lock.Lock()

	defer fBufReadWrite.lock.Unlock()

	var newFBuffReadWrite = FileBufferReadWrite{}

	var ePrefix *ePref.ErrPrefixDto
	var err error
	funcName := "FileBufferReadWrite." +
		"NewIoReadWrite()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {
		return newFBuffReadWrite, err
	}

	err = new(fileBufferReadWriteNanobot).
		setIoReaderWriter(
			&newFBuffReadWrite,
			"newFBuffReadWrite",
			reader,
			"reader",
			readerBuffSize,
			writer,
			"writer",
			writerBuffSize,
			ePrefix)

	return newFBuffReadWrite, err
}

// NewFileMgrs
//
// Creates and returns a new instance of
// FileBufferReadWrite.
//
// The internal io.Reader and io.Writer member variables
// for this new instance of FileBufferReadWrite are
// generated from input parameters specifying a 'reader'
// and a 'writer' extracted from File Manager instances
// (FileMgr).
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	readerFileMgr				*FileMgr
//
//		A pointer to an instance of FileMgr. The file
//		identified by 'fileMgr' will be used as a data
//		source for 'read' operations and will be
//		configured as an internal io.Reader for the
//		returned instance of FileBufferReadWrite.
//
//		If the path and file name encapsulated by
//		'readerFileMgr' do not currently exist on an
//		attached storage drive, an error will be
//		returned.
//
//	openReadFileReadWrite		bool
//
//		If this parameter is set to 'true', the target
//		'read' file identified by input parameter
//		'readerFileMgr' will be opened for both 'read'
//		and 'write' operations.
//
//		If 'openReadFileReadWrite' is set to 'false', the
//		target 'read' file will be opened for 'read-only'
//		operations.
//
//	readerBuffSize				int
//
//		This integer value controls the size of the
//		'read' buffer created for the internal io.Reader
//		object encapsulated in the returned instance of
//		FileBufferReadWrite.
//
//		'readerBuffSize' should be configured to maximize
//		performance for 'read' operations subject to
//		prevailing memory limitations.
//
//		The minimum reader buffer size is 16-bytes. If
//		'readerBuffSize' is set to a size less than "16",
//		it will be automatically reset to the default
//		buffer size of 4096-bytes.
//
//	writerFileMgr				*FileMgr
//
//		A pointer to an instance of FileMgr. The file
//		identified by 'fileMgr' will be used as an output
//		destination for 'write' operations and will be
//		configured as an internal io.Writer for the
//		returned instance of FileBufferReadWrite.
//
//		If the path and file name encapsulated by
//		'writerFileMgr' do not currently exist on an
//		attached storage drive, this method will attempt
//		to create them.
//
//	openWriteFileReadWrite		bool
//
//		If this parameter is set to 'true', the target
//		'write' file created from input parameter
//		'writerFileMgr' will be opened for 'read' and
//		'write' operations.
//
//		If 'openFileReadWrite' is set to 'false', the
//		target write file will be opened for 'write-only'
//		operations.
//
//	writerBuffSize				int
//
//		This integer value controls the size of the
//		'write' buffer created for the internal io.Writer
//		object encapsulated in the returned instance of
//		FileBufferReadWrite.
//
//		'writerBuffSize' should be configured to maximize
//		performance for 'write' operations subject to
//		prevailing memory limitations.
//
//		The minimum write buffer size is 1-byte. If
//		'writerBuffSize' is set to a size less than or
//		equal to zero, it will be automatically reset to
//		the default buffer size of 4096-bytes.
//
//	truncateExistingFile			bool
//
//		If this parameter is set to 'true', the target
//		'write' file identified by 'writerFileMgr' will
//		be opened for 'write' operations. If the target
//		file previously existed, it will be truncated.
//		This means that the file's previous contents will
//		be deleted.
//
//		If this parameter is set to 'false', the target
//		write file will be opened for 'write' operations.
//		If the target file previously existed, the new
//		text written to the file will be appended to the
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
//	FileBufferReadWrite
//
//		If this method completes successfully, it will
//		return a fully configured instance of
//		FileBufferReadWrite.
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
func (fBufReadWrite *FileBufferReadWrite) NewFileMgrs(
	readerFileMgr *FileMgr,
	openReadFileReadWrite bool,
	readerBuffSize int,
	writerFileMgr *FileMgr,
	openWriteFileReadWrite bool,
	writerBuffSize int,
	truncateExistingWriteFile bool,
	errorPrefix interface{}) (
	readerFileInfoPlus FileInfoPlus,
	writerFileInfoPlus FileInfoPlus,
	newFBuffReadWrite FileBufferReadWrite,
	err error) {

	if fBufReadWrite.lock == nil {
		fBufReadWrite.lock = new(sync.Mutex)
	}

	fBufReadWrite.lock.Lock()

	defer fBufReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "FileBufferReadWrite." +
		"NewFileMgrs()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {

		return readerFileInfoPlus,
			writerFileInfoPlus,
			newFBuffReadWrite,
			err
	}

	readerFileInfoPlus,
		writerFileInfoPlus,
		err = new(fileBufferReadWriteMicrobot).
		setFileMgrsReadWrite(
			&newFBuffReadWrite,
			"newFileBufReader",
			readerFileMgr,
			"readerFileMgr",
			openReadFileReadWrite,
			readerBuffSize,
			writerFileMgr,
			"writerFileMgr",
			openWriteFileReadWrite,
			writerBuffSize,
			truncateExistingWriteFile,
			ePrefix)

	return readerFileInfoPlus,
		writerFileInfoPlus,
		newFBuffReadWrite,
		err
}

// NewPathFileNames
//
// Creates and returns a new instance of
// FileBufferReadWrite.
//
// The internal io.Reader and io.Writer member variables
// for this new instance of FileBufferReadWrite are
// generated from input parameters specifying 'reader'
// and 'writer' files extracted from path and file name
// strings.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	readerPathFileName			string
//
//		This string contains the path and file name of
//		the file which will be used a data source for
//		'read' operations.
//
//		If this file does not currently exist on an
//		attached storage drive, an error will be
//		returned.
//
//	openReadFileReadWrite		bool
//
//		If this parameter is set to 'true', the target
//		'read' file identified by input parameter
//		'readerPathFileName' will be opened for both
//		'read' and 'write' operations.
//
//		If 'openReadFileReadWrite' is set to 'false', the
//		target 'read' file will be opened for 'read-only'
//		operations.
//
//	readerBuffSize					int
//
//		This integer value controls the size of the
//		'read' buffer created for the io.Reader object
//		generated from the file identified by
//		'readerPathFileName'. This io.Reader object is
//		encapsulated in the FileBufferReadWrite instance
//		returned by this method.
//
//		'readerBuffSize' should be configured to maximize
//		performance for 'read' operations subject to
//		prevailing memory limitations.
//
//		The minimum reader buffer size is 16-bytes. If
//		'readerBuffSize' is set to a size less than "16",
//		it will be automatically reset to the default
//		buffer size of 4096-bytes.
//
//	writerPathFileName			string
//
//		This string contains the path and file name of
//		the target 'write' file which will be used as
//		an output destination for 'write' operations.
//
//		If the target path and file do not currently
//		exist on an attached storage drive, this method
//		will attempt to create them.
//
//	openWriteFileReadWrite		bool
//
//		If this parameter is set to 'true', the target
//		'write' file identified by input parameter
//		'writerPathFileName' will be opened for 'read'
//		and 'write' operations.
//
//		If 'openWriteFileReadWrite' is set to 'false',
//		the target write file will be opened for
//		'write-only' operations.
//
//	writerBuffSize				int
//
//		This integer value controls the size of the
//		'write' buffer created for the io.Writer
//		object generated from the file identified by
//		input parameter 'writerPathFileName'. This
//		io.Writer object is encapsulated in the
//		FileBufferReadWrite instance returned by this
//		method.
//
//		'writerBuffSize' should be configured to maximize
//		performance for 'write' operations subject to
//		prevailing memory limitations.
//
//		If 'writerBuffSize' is set to a value less than
//		or equal to zero (0), it will be automatically
//		reset to the default value of 4096-bytes.
//
//	truncateExistingWriteFile	bool
//
//		If this parameter is set to 'true', the target
//		'write' file ('writerPathFileName') will be
//		opened for write operations. If the target write
//		file previously existed, it will be truncated.
//		This means that the file's previous contents will
//		be deleted.
//
//		If this parameter is set to 'false', the target
//		'write' file will be opened for write operations.
//		If the target 'write' file previously existed,
//		the new text written to this file will be appended
//		to the end of the previous file contents.
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
//	readerFileInfoPlus			FileInfoPlus
//
//		This returned instance of Type FileInfoPlus
//		contains data elements describing the file
//		identified by input parameter
//		'readerPathFileName'.
//
//		Type FileInfoPlus conforms to the os.FileInfo
//		interface. This structure will store os.FileInfo
//	 	information plus additional information related
//	 	to a file or directory.
//
//		type os.FileInfo interface {
//
//				Name() string
//					base name of the file
//
//				Size() int64
//					length in bytes for regular files;
//					system-dependent for others
//
//				Mode() FileMode
//					file mode bits
//
//				ModTime() time.Time
//					modification time
//
//				IsDir() bool
//					abbreviation for Mode().IsDir()
//
//				Sys() any
//					underlying data source (can return nil)
//		}
//
//		See the detailed documentation for Type
//		FileInfoPlus in the source file,
//		'fileinfoplus.go'.
//
//	writerFileInfoPlus			FileInfoPlus
//
//		This returned instance of Type FileInfoPlus
//		contains data elements describing the file
//		identified by input parameter
//		'writerPathFileName'.
//
//		Type FileInfoPlus conforms to the os.FileInfo
//		interface. This structure will store os.FileInfo
//	 	information plus additional information related
//	 	to a file or directory.
//
//		type os.FileInfo interface {
//
//				Name() string
//					base name of the file
//
//				Size() int64
//					length in bytes for regular files;
//					system-dependent for others
//
//				Mode() FileMode
//					file mode bits
//
//				ModTime() time.Time
//					modification time
//
//				IsDir() bool
//					abbreviation for Mode().IsDir()
//
//				Sys() any
//					underlying data source (can return nil)
//		}
//
//		See the detailed documentation for Type
//		FileInfoPlus in the source file,
//		'fileinfoplus.go'.
//
//	newFBuffReadWrite			FileBufferReadWrite
//
//		If this method completes successfully, it will
//		return a fully configured instance of
//		FileBufferReadWrite.
//
//	err							error
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
func (fBufReadWrite *FileBufferReadWrite) NewPathFileNames(
	readerPathFileName string,
	openReadFileReadWrite bool,
	readerBuffSize int,
	writerPathFileName string,
	openWriteFileReadWrite bool,
	writerBuffSize int,
	truncateExistingWriteFile bool,
	errorPrefix interface{}) (
	readerFileInfoPlus FileInfoPlus,
	writerFileInfoPlus FileInfoPlus,
	newFBuffReadWrite FileBufferReadWrite,
	err error) {

	if fBufReadWrite.lock == nil {
		fBufReadWrite.lock = new(sync.Mutex)
	}

	fBufReadWrite.lock.Lock()

	defer fBufReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "FileBufferReadWrite." +
		"NewPathFileNames()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {

		return readerFileInfoPlus,
			writerFileInfoPlus,
			newFBuffReadWrite,
			err
	}

	readerFileInfoPlus,
		writerFileInfoPlus,
		err = new(fileBufferReadWriteNanobot).
		setPathFileNamesReadWrite(
			&newFBuffReadWrite,
			"newFBuffReadWrite",
			readerPathFileName,
			"readerPathFileName",
			openReadFileReadWrite,
			readerBuffSize,
			writerPathFileName,
			"writerPathFileName",
			openWriteFileReadWrite,
			writerBuffSize,
			truncateExistingWriteFile,
			ePrefix)

	return readerFileInfoPlus,
		writerFileInfoPlus,
		newFBuffReadWrite,
		err
}

// ReadWriteAll
//
// This method reads all data from the 'reader' data
// source and writes all that data to the 'writer'
// destination.
//
// If the total number of bytes read does NOT equal the
// total number of bytes written, an error will be
// returned.
//
// The 'read' and 'write' operations use the io.Reader
// and io.Writer objects created when the current
// instance of FileBufferReadWrite was initialized.
//
// If input parameter 'autoFlushAndCloseOnExit' is set to
// 'true', this method will automatically perform all
// required clean-up tasks. Clean-up tasks involve
// flushing the io.Writer object, closing the io.Reader
// and io.Writer objects and then deleting io.Reader and
// io.Writer structures internal to the current
// FileBufferReadWrite instance. When these Clean-up tasks
// are completed, the current FileBufferReadWrite instance
// will be invalid and unusable for future 'read' and/or
// 'write' operations.
//
// If input parameter 'autoFlushAndCloseOnExit' is set to
// 'false', the user is responsible for performing
// clean-up tasks by calling the local method:
//
//	FileBufferReadWrite.CloseFileBufferReadWrite()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	autoFlushAndCloseOnExit		bool
//
//		When this parameter is set to 'true', this
//		method will automatically perform the following
//		clean-up tasks upon exit:
//
//		(1)	The write buffer will be flushed thereby
//			ensuring that all remaining data in the
//			'write' buffer will be written to the
//			underlying io.Writer object.
//
//		(2)	The io.Reader and io.Writer objects will be
//			properly closed.
//
//		(3) After performing these clean-up tasks, the
//			current instance of FileBufferReadWrite will
//			invalid and unusable for future 'read' and/or
//			'write' operations.
//
//		If input parameter 'autoFlushAndCloseOnExit' is
//		set to 'false', the user is responsible for
//		performing clean-up tasks by calling the local
//		method:
//
//			FileBufferReadWrite.CloseFileBufferReadWrite()
//
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
//	err							error
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
//
//		If the number of bytes written to the destination
//		is NOT equal to the number of bytes read from the
//		source, an error will be returned.
func (fBufReadWrite *FileBufferReadWrite) ReadWriteAll(
	autoFlushAndCloseOnExit bool,
	errorPrefix interface{}) (
	totalBytesRead int,
	totalBytesWritten int,
	err error) {

	if fBufReadWrite.lock == nil {
		fBufReadWrite.lock = new(sync.Mutex)
	}

	fBufReadWrite.lock.Lock()

	defer fBufReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "FileBufferReadWrite." +
		"ReadWriteAll()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {

		return totalBytesRead, totalBytesWritten, err
	}

	if fBufReadWrite.reader == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: The current instance of FileBufferReadWrite\n"+
			"is invalid! The internal io.Reader object was never\n"+
			"initialized. Call one of the 'New' methods or 'Setter'\n"+
			"methods to create a valid instance of FileBufferReadWrite.\n",
			ePrefix.String())

		return totalBytesRead, totalBytesWritten, err
	}

	if fBufReadWrite.writer == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: The current instance of FileBufferReadWrite\n"+
			"is invalid! The internal io.Writer object was never\n"+
			"initialized. Call one of the 'New' methods or 'Setter'\n"+
			"methods to create a valid instance of FileBufferReadWrite.\n",
			ePrefix.String())

		return totalBytesRead, totalBytesWritten, err
	}

	var readErr, writeErr error
	var numOfBytesRead, numOfBytesWritten, cycleCount int
	var errs []error
	var err2 error
	var fBufReadWriteMicrobot = new(fileBufferReadWriteMicrobot)

	byteArray := make([]byte,
		fBufReadWrite.reader.fileReader.Size())

	for {

		cycleCount++

		numOfBytesRead,
			readErr =
			fBufReadWrite.reader.Read(byteArray)

		if readErr != nil &&
			readErr != io.EOF {

			errs = append(
				errs,
				fmt.Errorf("%v\n"+
					"Error Reading Target Read File!\n"+
					"Cycle Count= %v\n"+
					"Error= \n%v\n",
					funcName,
					cycleCount,
					readErr.Error()))

			err2 = fBufReadWriteMicrobot.
				closeReaderWriter(
					fBufReadWrite,
					"fBufReadWrite",
					funcName)

			if err2 != nil {

				errs = append(
					errs,
					fmt.Errorf("%v\n"+
						"%v",
						funcName,
						err2.Error()))

			}

			if len(errs) > 0 {

				err = fmt.Errorf("%v\n"+
					"%v",
					ePrefix.String(),
					new(StrMech).ConsolidateErrors(errs).Error())
			}

			return totalBytesRead, totalBytesWritten, err
		}

		if numOfBytesRead > 0 {

			totalBytesRead += numOfBytesRead

			numOfBytesWritten,
				writeErr = fBufReadWrite.writer.Write(
				byteArray[0:numOfBytesRead])

			if writeErr != nil {

				errs = append(
					errs,
					fmt.Errorf("%v\n"+
						"Error Writing Bytes To File!\n"+
						"Write Error=\n%v\n",
						funcName,
						writeErr.Error()))

				err2 = fBufReadWriteMicrobot.
					closeReaderWriter(
						fBufReadWrite,
						"fBufReadWrite",
						funcName)

				if err2 != nil {

					errs = append(
						errs,
						fmt.Errorf("%v\n"+
							"%v",
							funcName,
							err2.Error()))

				}

				if len(errs) > 0 {

					err = fmt.Errorf("%v\n"+
						"%v",
						ePrefix.String(),
						new(StrMech).ConsolidateErrors(errs).Error())
				}

				return totalBytesRead, totalBytesWritten, err
			}

			if numOfBytesRead != numOfBytesWritten {

				errs = append(
					errs,
					fmt.Errorf("%v\n"+
						"Error Writing Bytes To File!\n"+
						"numOfBytesRead != numOfBytesWritten\n"+
						"numOfBytesRead = %v\n"+
						"numOfBytesWritten = %v\n",
						funcName,
						numOfBytesRead,
						numOfBytesWritten))

				err2 = fBufReadWriteMicrobot.
					closeReaderWriter(
						fBufReadWrite,
						"fBufReadWrite",
						funcName)

				if err2 != nil {

					errs = append(
						errs,
						fmt.Errorf("%v\n"+
							"%v",
							funcName,
							err2.Error()))

				}

				if len(errs) > 0 {

					err = fmt.Errorf("%v\n"+
						"%v",
						ePrefix.String(),
						new(StrMech).ConsolidateErrors(errs).Error())
				}

				return totalBytesRead, totalBytesWritten, err
			}

			totalBytesWritten += numOfBytesWritten
		}

		if readErr == io.EOF {
			break
		}

	}

	if autoFlushAndCloseOnExit == true {

		err = fBufReadWriteMicrobot.
			closeReaderWriter(
				fBufReadWrite,
				"fBufReadWrite",
				ePrefix.XCpy("Close-Readers&Writers"))

	}

	return totalBytesRead, totalBytesWritten, err
}

// Read
//
// Reads a selection data from the pre-configured
// io.Reader data source encapsulated in the current
// instance of FileBufferReadWrite.
//
// Method 'Read' reads data into the input parameter
// byte array, 'bytesRead'. It returns the number of
// bytes read into the byte array as return parameter,
// 'numOfBytesRead'.
//
// Under certain circumstances, the number of bytes read
// into the byte array may be less than the length of the
// byte array (len(bytesRead)) due to the length of the
// underlying read buffer.
//
// To complete the read operation, repeat the call to
// this method until the returned error is set to
// 'io.EOF' signaling 'End of File'.
//
// See the io.Reader docs and 'Reference' section below.
//
// If the internal io.Reader object for the current
// instance of FileBufferReadWrite was improperly
// initialized, an error will be returned. To properly
// initialized an instance of FileBufferReadWrite, the
// user must call one or more of the 'New' and/or
// 'Setter' methods.
//
// Once all 'read' and 'write' operations have been
// completed for the current instance of
// FileBufferReadWrite, the user MUST call the 'Close'
// method to ensure clean-up operations are properly
// applied:
//
//	FileBufferReadWrite.CloseFileBufferReadWrite()
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
//	(1)	This method implements the io.Reader interface.
//
//	(2)	Keep calling this method until all the bytes have
//		been read from the 'read' data source configured
//		for the current instance of FileBufferReadWrite
//		and the returned error is set to 'io.EOF'.
//
//	(3)	Callers should always process the
//		numOfBytesRead > 0 bytes returned before
//		considering the error err. Doing so correctly
//		handles I/O errors that happen after reading some
//		bytes and also both of the allowed EOF behaviors
//		(See the io.Reader docs and 'Reference' section
//		below).
//
//	(4)	When all 'read' and 'write' operations have been
//		completed, call method:
//
//			FileBufferReadWrite.CloseFileBufferReadWrite()
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
//		Bytes will be read from the 'read' data source
//		previously configured for the current instance of
//		FileBufferReadWrite.
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
//		section above for a discussion of 'io.EOF'. Disk
//		files will return an 'io.EOF'. However, some other
//		types of readers may not.
func (fBufReadWrite *FileBufferReadWrite) Read(
	bytesRead []byte) (
	numOfBytesRead int,
	err error) {

	if fBufReadWrite.lock == nil {
		fBufReadWrite.lock = new(sync.Mutex)
	}

	fBufReadWrite.lock.Lock()

	defer fBufReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"FileBufferReadWrite."+
			"Read()",
		"")

	if err != nil {

		return numOfBytesRead, err
	}

	if fBufReadWrite.reader == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: The current instance of FileBufferReadWrite\n"+
			"is invalid! The internal io.Reader object was never\n"+
			"initialized. Call one of the 'New' methods or 'Setter'\n"+
			"methods to create a valid instance of FileBufferReadWrite.\n",
			ePrefix.String())

		return numOfBytesRead, err
	}

	var err2 error

	numOfBytesRead,
		err2 = fBufReadWrite.reader.
		Read(bytesRead)

	if err2 != nil {

		if len(fBufReadWrite.readerFilePathName) > 0 {

			err = fmt.Errorf("%v\n"+
				"Error reading from 'read' data source.\n"+
				"Read File Path and File Name: %v\n"+
				"Error=\n%v\n",
				ePrefix.String(),
				fBufReadWrite.readerFilePathName,
				err2.Error())

		} else {

			err = fmt.Errorf("%v\n"+
				"Error reading from 'read' data source.\n"+
				"Error=\n%v\n",
				ePrefix.String(),
				err2.Error())

		}
	}

	return numOfBytesRead, err
}

// Write
//
// Writes the contents of the byte array input paramter
// ('bytesToWrite') to the output destination io.Writer
// object previously configured for the current instance
// of FileBufferReadWrite.
//
// If the internal io.Writer object for the current
// instance of FileBufferReadWrite was improperly
// initialized, an error will be returned. To properly
// initialized an instance of FileBufferReadWrite, the
// user must call one or more of the 'New' and 'Setter'
// methods.
//
// Once all 'read' and 'write' operations have been
// completed for the current instance of
// FileBufferReadWrite, the user MUST call the 'Close'
// method to ensure clean-up operations are properly
// applied:
//
//	FileBufferReadWrite.CloseFileBufferReadWrite()
//		Note: 	The Close operation performs both Flush
//				and Close tasks.
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
//	(1) This method implements the io.Writer interface.
//
//	(2)	After all 'read' and 'write' operations have been
//		completed, the user MUST call the 'Close' method
//		to perform necessary clean-up operations:
//
//		FileBufferReadWrite.CloseFileBufferReadWrite()
//
//	(3) This method WILL NOT VERIFY that the number of
//		bytes written is equal to the length of the
//		length of input parameter 'bytesToWrite'.
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
//		This method WILL NOT VERIFY that the number of
//		bytes written is equal to the length of the
//		length of input parameter 'bytesToWrite'.
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
func (fBufReadWrite *FileBufferReadWrite) Write(
	bytesToWrite []byte) (
	numBytesWritten int,
	err error) {

	if fBufReadWrite.lock == nil {
		fBufReadWrite.lock = new(sync.Mutex)
	}

	fBufReadWrite.lock.Lock()

	defer fBufReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"FileBufferReadWrite."+
			"Write()",
		"")

	if err != nil {

		return numBytesWritten, err
	}

	if fBufReadWrite.writer == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: The current instance of FileBufferReadWrite\n"+
			"is invalid! The internal io.Writer object was never\n"+
			"initialized. Call one of the 'New' methods or 'Setter'\n"+
			"methods to create a valid instance of FileBufferReadWrite.\n",
			ePrefix.String())

		return numBytesWritten, err
	}

	var err2 error

	numBytesWritten,
		err2 = fBufReadWrite.writer.
		Write(bytesToWrite)

	if err2 != nil {

		if len(fBufReadWrite.writerFilePathName) > 0 {

			err = fmt.Errorf("%v\n"+
				"Error writing data to the output data destination.\n"+
				"'Write' File Path and File Name: %v\n"+
				"Error=\n%v\n",
				ePrefix.String(),
				fBufReadWrite.writerFilePathName,
				err2.Error())

		} else {

			err = fmt.Errorf("%v\n"+
				"Error writing to the output data destination.\n"+
				"Error=\n%v\n",
				ePrefix.String(),
				err2.Error())

		}
	}

	return numBytesWritten, err
}

// SetFileMgrsReadWrite
//
// Receives two instances of FileMgr as input parameters
// identifying the io.Reader and io.Writer objects which
// will be configured for the current instance of
// FileBufferReadWrite.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the current instance of
//	FileBufferReadWrite.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	readerFileMgr				*FileMgr
//
//		A pointer to an instance of FileMgr. The file
//		identified by 'fileMgr' will be used as a data
//		source for 'read' operations and will be
//		configured as an internal io.Reader for the
//		returned instance of FileBufferReadWrite.
//
//		If the path and file name encapsulated by
//		'readerFileMgr' do not currently exist on an
//		attached storage drive, an error will be
//		returned.
//
//	openReadFileReadWrite		bool
//
//		If this parameter is set to 'true', the target
//		'read' file identified by input parameter
//		'readerFileMgr' will be opened for both 'read'
//		and 'write' operations.
//
//		If 'openReadFileReadWrite' is set to 'false', the
//		target 'read' file will be opened for 'read-only'
//		operations.
//
//	readerBuffSize				int
//
//		This integer value controls the size of the
//		'read' buffer created for the internal io.Reader
//		encapsulated in the current instance of
//		FileBufferReadWrite.
//
//		'readerBuffSize' should be configured to maximize
//		performance for 'read' operations subject to
//		prevailing memory limitations.
//
//		The minimum reader buffer size is 16-bytes. If
//		'readerBuffSize' is set to a size less than "16",
//		it will be automatically reset to the default
//		buffer size of 4096-bytes.
//
//	writerFileMgr				*FileMgr
//
//		A pointer to an instance of FileMgr. The file
//		identified by 'fileMgr' will be used as an output
//		destination for 'write' operations and will be
//		configured as an internal io.Writer for the
//		returned instance of FileBufferReadWrite.
//
//		If the path and file name encapsulated by
//		'writerFileMgr' do not currently exist on an
//		attached storage drive, this method will attempt
//		to create them.
//
//	openWriteFileReadWrite		bool
//
//		If this parameter is set to 'true', the target
//		'write' file created from input parameter
//		'writerFileMgr' will be opened for 'read' and
//		'write' operations.
//
//		If 'openFileReadWrite' is set to 'false', the
//		target write file will be opened for 'write-only'
//		operations.
//
//	writerBuffSize				int
//
//		This integer value controls the size of the
//		'write' buffer created for the internal io.Writer
//		object encapsulated in the current instance of
//		FileBufferReadWrite.
//
//		'writerBuffSize' should be configured to maximize
//		performance for 'write' operations subject to
//		prevailing memory limitations.
//
//		The minimum write buffer size is 1-byte. If
//		'writerBuffSize' is set to a size less than or
//		equal to zero, it will be automatically reset to
//		the default buffer size of 4096-bytes.
//
//	truncateExistingFile			bool
//
//		If this parameter is set to 'true', the target
//		'write' file identified by 'writerFileMgr' will
//		be opened for 'write' operations. If the target
//		file previously existed, it will be truncated.
//		This means that the file's previous contents will
//		be deleted.
//
//		If this parameter is set to 'false', the target
//		write file will be opened for 'write' operations.
//		If the target file previously existed, the new
//		text written to the file will be appended to the
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
//	readerFileInfoPlus			FileInfoPlus
//
//		This returned instance of Type FileInfoPlus
//		contains data elements describing the file
//		identified by input parameter 'readerFileMgr'.
//
//		Type FileInfoPlus conforms to the os.FileInfo
//		interface. This structure will store os.FileInfo
//	 	information plus additional information related
//	 	to a file or directory.
//
//		type os.FileInfo interface {
//
//				Name() string
//					base name of the file
//
//				Size() int64
//					length in bytes for regular files;
//					system-dependent for others
//
//				Mode() FileMode
//					file mode bits
//
//				ModTime() time.Time
//					modification time
//
//				IsDir() bool
//					abbreviation for Mode().IsDir()
//
//				Sys() any
//					underlying data source (can return nil)
//		}
//
//		See the detailed documentation for Type
//		FileInfoPlus in the source file,
//		'fileinfoplus.go'.
//
//	writerFileInfoPlus			FileInfoPlus
//
//		This returned instance of Type FileInfoPlus
//		contains data elements describing the file
//		identified by input parameter 'writerFileMgr'.
//
//		Type FileInfoPlus conforms to the os.FileInfo
//		interface. This structure will store os.FileInfo
//	 	information plus additional information related
//	 	to a file or directory.
//
//		type os.FileInfo interface {
//
//				Name() string
//					base name of the file
//
//				Size() int64
//					length in bytes for regular files;
//					system-dependent for others
//
//				Mode() FileMode
//					file mode bits
//
//				ModTime() time.Time
//					modification time
//
//				IsDir() bool
//					abbreviation for Mode().IsDir()
//
//				Sys() any
//					underlying data source (can return nil)
//		}
//
//		See the detailed documentation for Type
//		FileInfoPlus in the source file,
//		'fileinfoplus.go'.
//
//	err							error
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
func (fBufReadWrite *FileBufferReadWrite) SetFileMgrsReadWrite(
	readerFileMgr *FileMgr,
	openReadFileReadWrite bool,
	readerBuffSize int,
	writerFileMgr *FileMgr,
	openWriteFileReadWrite bool,
	writerBuffSize int,
	truncateExistingWriteFile bool,
	errorPrefix interface{}) (
	readerFileInfoPlus FileInfoPlus,
	writerFileInfoPlus FileInfoPlus,
	err error) {

	if fBufReadWrite.lock == nil {
		fBufReadWrite.lock = new(sync.Mutex)
	}

	fBufReadWrite.lock.Lock()

	defer fBufReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileBufferReadWrite."+
			"SetFileMgrsReadWrite()",
		"")

	if err != nil {

		return readerFileInfoPlus, writerFileInfoPlus, err
	}

	readerFileInfoPlus,
		writerFileInfoPlus,
		err = new(fileBufferReadWriteMicrobot).
		setFileMgrsReadWrite(
			fBufReadWrite,
			"fBufReadWrite",
			readerFileMgr,
			"readerFileMgr",
			openReadFileReadWrite,
			readerBuffSize,
			writerFileMgr,
			"writerFileMgr",
			openWriteFileReadWrite,
			writerBuffSize,
			truncateExistingWriteFile,
			ePrefix)

	return readerFileInfoPlus, writerFileInfoPlus, err
}

// SetFileMgrReader
//
// This method will close, delete and reconfigure the
// internal io.Reader object encapsulated in the current
// instance of FileBufferReadWrite. The internal
// io.Reader object is used to 'read' data from a data
// source such as a disk file.
//
// This internal io.Reader object will be reconfigured
// using the file identified by a File Manager instance
// passed as input parameter 'readerFileMgr'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reconfigure
//	the member variable io.Reader object encapsulated in
//	the current instance of FileBufferReadWrite.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	readerFileMgr				*FileMgr
//
//		A pointer to an instance of FileMgr. The file
//		identified by 'readerFileMgr' will be used as a
//		data source for 'read' operations and will be
//		configured as an internal io.Reader for the
//		FileBufferReadWrite instance passed as input
//		parameter 'fBufReadWrite'.
//
//		If the path and file name encapsulated by
//		'readerFileMgr' do not currently exist on an
//		attached storage drive, an error will be
//		returned.
//
//	openReadFileReadWrite		bool
//
//		If this parameter is set to 'true', the target
//		'read' file identified by input parameter
//		'readerFileMgr' will be opened for both 'read'
//		and 'write' operations.
//
//		If 'openReadFileReadWrite' is set to 'false', the
//		target 'read' file will be opened for 'read-only'
//		operations.
//
//	readerBuffSize				int
//
//		This integer value controls the size of the
//		'read' buffer created for the internal io.Reader
//		encapsulated in the current instance of
//		FileBufferReadWrite.
//
//		'readerBuffSize' should be configured to maximize
//		performance for 'read' operations subject to
//		prevailing memory limitations.
//
//		The minimum reader buffer size is 16-bytes. If
//		'readerBuffSize' is set to a size less than "16",
//		it will be automatically reset to the default
//		buffer size of 4096-bytes.
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
//	readerFileInfoPlus			FileInfoPlus
//
//		This returned instance of Type FileInfoPlus
//		contains data elements describing the file
//		identified by input parameter 'readerFileMgr'.
//
//		Type FileInfoPlus conforms to the os.FileInfo
//		interface. This structure will store os.FileInfo
//	 	information plus additional information related
//	 	to a file or directory.
//
//		type os.FileInfo interface {
//
//				Name() string
//					base name of the file
//
//				Size() int64
//					length in bytes for regular files;
//					system-dependent for others
//
//				Mode() FileMode
//					file mode bits
//
//				ModTime() time.Time
//					modification time
//
//				IsDir() bool
//					abbreviation for Mode().IsDir()
//
//				Sys() any
//					underlying data source (can return nil)
//		}
//
//		See the detailed documentation for Type
//		FileInfoPlus in the source file,
//		'fileinfoplus.go'.
//
//	err							error
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
func (fBufReadWrite *FileBufferReadWrite) SetFileMgrReader(
	readerFileMgr *FileMgr,
	openReadFileReadWrite bool,
	readerBuffSize int,
	errorPrefix interface{}) (
	readerFileInfoPlus FileInfoPlus,
	err error) {

	if fBufReadWrite.lock == nil {
		fBufReadWrite.lock = new(sync.Mutex)
	}

	fBufReadWrite.lock.Lock()

	defer fBufReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileBufferReadWrite."+
			"SetFileMgrReader()",
		"")

	if err != nil {

		return readerFileInfoPlus, err
	}

	readerFileInfoPlus,
		err = new(fileBufferReadWriteMolecule).
		setFileMgrReader(
			fBufReadWrite,
			"fBufReadWrite",
			readerFileMgr,
			"readerFileMgr",
			openReadFileReadWrite,
			readerBuffSize,
			ePrefix)

	return readerFileInfoPlus, err
}

// SetFileMgrWriter
//
// This method will close, delete and reconfigure the
// internal io.Writer object encapsulated in the current
// instance of FileBufferReadWrite. The internal
// io.Writer object is used to 'write' to an output data
// destination such as a disk file.
//
// This internal io.Writer object will be reconfigured
// using the file identified by a File Manager instance
// passed as input parameter 'writerFileMgr'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reconfigure
//	the member variable io.Writer object encapsulated in
//	the current instance of FileBufferReadWrite.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	writerFileMgr				*FileMgr
//
//		A pointer to an instance of FileMgr. The file
//		identified by 'writerFileMgr' will be used as an
//		output data destination for 'write' operations
//		and will be configured as an internal io.Writer
//		for the current instance of FileBufferReadWrite.
//
//		If the path and file name encapsulated by
//		'writerFileMgr' do not currently exist on an
//		attached storage drive, this method will attempt
//		to create them.
//
//	openWriteFileReadWrite		bool
//
//		If this parameter is set to 'true', the target
//		'write' file identified by input parameter
//		'writerFileMgr' will be opened for 'read'
//		and 'write' operations.
//
//		If 'openWriteFileReadWrite' is set to 'false',
//		the target write file will be opened for
//		'write-only' operations.
//
//	writerBuffSize				int
//
//		This integer value controls the size of the
//		'write' buffer created for the io.Writer object
//		generated from the file identified by
//		'writerFileMgr'. This io.Writer object will in
//		turn be configured and encapsulated in the
//		current instance of FileBufferWriter.
//
//		'writerBuffSize' should be configured to maximize
//		performance for 'write' operations subject to
//		prevailing memory limitations.
//
//		If 'writerBuffSize' is set to a value less than
//		or equal to zero (0), it will be automatically
//		reset to the default value of 4096-bytes.
//
//	truncateExistingWriteFile	bool
//
//		If this parameter is set to 'true', the target
//		'write' file ('writerFileMgr') will be opened for
//		'write' operations. If the target 'write' file
//		previously existed, it will be truncated. This
//		means that the file's previous contents will be
//		deleted.
//
//		If this parameter is set to 'false', the target
//		'write' file will be opened for write operations.
//		If the target 'write' file previously existed,
//		the new text written to this file will be appended
//		to the end of the previous file contents.
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
//	writerFileInfoPlus			FileInfoPlus
//
//		This returned instance of Type FileInfoPlus
//		contains data elements describing the file
//		identified by input parameter 'writerFileMgr'.
//
//		Type FileInfoPlus conforms to the os.FileInfo
//		interface. This structure will store os.FileInfo
//	 	information plus additional information related
//	 	to a file or directory.
//
//		type os.FileInfo interface {
//
//				Name() string
//					base name of the file
//
//				Size() int64
//					length in bytes for regular files;
//					system-dependent for others
//
//				Mode() FileMode
//					file mode bits
//
//				ModTime() time.Time
//					modification time
//
//				IsDir() bool
//					abbreviation for Mode().IsDir()
//
//				Sys() any
//					underlying data source (can return nil)
//		}
//
//		See the detailed documentation for Type
//		FileInfoPlus in the source file,
//		'fileinfoplus.go'.
//
//	err							error
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
func (fBufReadWrite *FileBufferReadWrite) SetFileMgrWriter(
	writerFileMgr *FileMgr,
	openWriteFileReadWrite bool,
	writerBuffSize int,
	truncateExistingWriteFile bool,
	errorPrefix interface{}) (
	writerFileInfoPlus FileInfoPlus,
	err error) {

	if fBufReadWrite.lock == nil {
		fBufReadWrite.lock = new(sync.Mutex)
	}

	fBufReadWrite.lock.Lock()

	defer fBufReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileBufferReadWrite."+
			"SetFileMgrWriter()",
		"")

	if err != nil {

		return writerFileInfoPlus, err
	}

	writerFileInfoPlus,
		err = new(fileBufferReadWriteMolecule).
		setFileMgrWriter(
			fBufReadWrite,
			"fBufReadWrite",
			writerFileMgr,
			"writerFileMgr",
			openWriteFileReadWrite,
			writerBuffSize,
			truncateExistingWriteFile,
			ePrefix)

	return writerFileInfoPlus, err
}

// SetIoReadWrite
//
// This method will close, delete and reconfigure the
// internal io.Reader and io.Writer objects encapsulated
// in the current instance of FileBufferReadWrite. The
// internal io.Reader object is used to 'read' data from
// a data source such as a disk file. In contrast, the
// io.Writer object is used to 'write' data to an output
// destination such as a disk file.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the current instance
//	of FileBufferReadWrite.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	reader						io.Reader
//
//		Any object which implements io.Reader interface.
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
//		in addition to calling local method:
//
//		FileBufferReadWrite.CloseFileBufferReadWrite()
//			This method will flush the 'write' buffer
//			in addition to closing and performing
//			clean-up tasks for the io.Reader and
//			io.Writer objects.
//
//		While the 'read' services provided by
//		FileBufferReadWrite are primarily designed to
//		read data from disk files, this type of 'reader'
//		will in fact read data from any object
//		implementing the io.Reader interface.
//
//	readerBuffSize				int
//
//		This integer value controls the size of the
//		buffer created for the io.Reader object passed as
//		input parameter 'reader'.
//
//		'readerBuffSize' should be configured to maximize
//		performance for 'read' operations subject to
//		prevailing memory limitations.
//
//		The minimum reader buffer size is 16-bytes. If
//		'readerBuffSize' is set to a value less than
//		"16", it will be automatically reset to the
//		default buffer size of 4096-bytes.
//
//	writer						io.Writer
//
//		This parameter will accept any object
//		implementing the io.Writer interface.
//
//		This object may be a file pointer of type
//		*os.File. File pointers of this type implement
//		the io.Writer interface.
//
//		A file pointer (*os.File) will facilitate writing
//		output data to destination files residing on an
//		attached storage drive. However, with this
//		configuration, the user is responsible for
//		manually closing the file and performing any
//		other required clean-up operations in addition to
//		calling local method:
//
//		FileBufferReadWrite.CloseFileBufferReadWrite()
//			This method will flush the 'write' buffer
//			in addition to closing and performing
//			clean-up tasks for the io.Reader and
//			io.Writer objects.
//
//		While the 'write' services provided by the
//		FileBufferReadWrite are primarily designed for
//		writing data to disk files, this type of 'writer'
//		will in fact write data to any object
//		implementing the io.Writer interface.
//
//	writerBuffSize				int
//
//		This integer value controls the size of the
//		'write' buffer created for the io.Writer
//		object passed as input parameter 'writer'.
//		This io.Writer object will in turn be configured
//		and encapsulated in the current instance of
//		FileBufferWriter.
//
//		'writerBuffSize' should be configured to maximize
//		performance for 'write' operations subject to
//		prevailing memory limitations.
//
//		The minimum write buffer size is 1-byte. If
//		'writerBuffSize' is set to a size less than or
//		equal to zero, it will be automatically reset to
//		the default buffer size of 4096-bytes.
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
func (fBufReadWrite *FileBufferReadWrite) SetIoReadWrite(
	reader io.Reader,
	readerBuffSize int,
	writer io.Writer,
	writerBuffSize int,
	errorPrefix interface{}) error {

	if fBufReadWrite.lock == nil {
		fBufReadWrite.lock = new(sync.Mutex)
	}

	fBufReadWrite.lock.Lock()

	defer fBufReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileBufferReadWrite."+
			"SetIoReadWrite()",
		"")

	if err != nil {
		return err
	}

	return new(fileBufferReadWriteNanobot).
		setIoReaderWriter(
			fBufReadWrite,
			"fBufReadWrite",
			reader,
			"reader",
			readerBuffSize,
			writer,
			"writer",
			writerBuffSize,
			ePrefix)
}

// SetIoReader
//
// This method will close, delete and reconfigure the
// internal io.Reader object encapsulated in the current
// instance of FileBufferReadWrite. The internal
// io.Reader object is used to 'read' data from a data
// source such as a disk file.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reconfigure
//	the member variable io.Reader object encapsulated in
//	the current instance of FileBufferReadWrite.
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
//		in addition to calling local method:
//
//		FileBufferReadWrite.CloseFileBufferReadWrite()
//
//		While the 'read' services provided by
//		FileBufferReadWrite are primarily designed to
//		read data from disk files, this type of 'reader'
//		will in fact read data from any object
//		implementing the io.Reader interface.
//
//	readerBuffSize				int
//
//		This integer value controls the size of the
//		buffer created for the io.Reader object passed as
//		input parameter 'reader'.
//
//		'readerBuffSize' should be configured to maximize
//		performance for 'read' operations subject to
//		prevailing memory limitations.
//
//		The minimum reader buffer size is 16-bytes. If
//		'readerBuffSize' is set to a value less than
//		"16", it will be automatically reset to the
//		default buffer size of 4096-bytes.
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
func (fBufReadWrite *FileBufferReadWrite) SetIoReader(
	reader io.Reader,
	readerBuffSize int,
	errorPrefix interface{}) error {

	if fBufReadWrite.lock == nil {
		fBufReadWrite.lock = new(sync.Mutex)
	}

	fBufReadWrite.lock.Lock()

	defer fBufReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileBufferReadWrite."+
			"SetIoReader()",
		"")

	if err != nil {
		return err
	}

	return new(fileBufferReadWriteAtom).
		setIoReader(
			fBufReadWrite,
			"fBufReadWrite",
			reader,
			"reader",
			readerBuffSize,
			ePrefix)
}

// SetIoWriter
//
// This method will close, delete and reconfigure the
// internal io.Writer object encapsulated in the current
// instance of FileBufferReadWrite. The internal
// io.Writer object is used to 'write' data to an output
// destination such as a disk file.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reconfigure
//	the member variable io.Writer object encapsulated in
//	the current instance of FileBufferReadWrite.
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
//		data to files residing on an attached storage
//		drive. However, with this configuration, the user
//		is responsible for manually closing the file and
//		performing any other required clean-up operations
//		in addition to calling local method:
//
//		FileBufferReadWrite.CloseFileBufferReadWrite()
//			This method will flush the 'write' buffer
//			in addition to closing and performing
//			clean-up tasks for the io.Reader and
//			io.Writer objects.
//
//		While the 'write' services provided by
//		FileBufferReadWrite are primarily designed to
//		write data to disk files, this type of 'writer'
//		will in fact write data to any object
//		implementing the io.Writer interface.
//
//	writerBuffSize				int
//
//		This integer value controls the size of the
//		'write' buffer created for the io.Writer
//		object passed as input parameter 'writer'.
//		This io.Writer object will in turn be configured
//		and encapsulated in the current instance of
//		FileBufferWriter.
//
//		'writerBuffSize' should be configured to maximize
//		performance for 'write' operations subject to
//		prevailing memory limitations.
//
//		The minimum write buffer size is 1-byte. If
//		'writerBuffSize' is set to a size less than or
//		equal to zero, it will be automatically reset to
//		the default buffer size of 4096-bytes.
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
func (fBufReadWrite *FileBufferReadWrite) SetIoWriter(
	writer io.Writer,
	writerBuffSize int,
	errorPrefix interface{}) error {

	if fBufReadWrite.lock == nil {
		fBufReadWrite.lock = new(sync.Mutex)
	}

	fBufReadWrite.lock.Lock()

	defer fBufReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileBufferReadWrite."+
			"SetIoWriter()",
		"")

	if err != nil {
		return err
	}

	return new(fileBufferReadWriteAtom).
		setIoWriter(
			fBufReadWrite,
			"fBufReadWrite",
			writer,
			"writer",
			writerBuffSize,
			ePrefix)
}

// SetPathFileNamesReadWrite
//
// Receives two strings as input parameters. These
// strings contain the path and file names of files
// used to construct the io.Reader and io.Writer
// objects which will be configured and encapsulated in
// the current instance of FileBufferReadWrite.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the current instance of
//	FileBufferReadWrite.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	readerPathFileName			string
//
//		This string contains the path and file name of
//		the file which will be used a data source for
//		'read' operations.
//
//		If this file does not currently exist on an
//		attached storage drive, an error will be
//		returned.
//
//	openReadFileReadWrite		bool
//
//		If this parameter is set to 'true', the target
//		'read' file identified by input parameter
//		'readerPathFileName' will be opened for both
//		'read' and 'write' operations.
//
//		If 'openReadFileReadWrite' is set to 'false', the
//		target 'read' file will be opened for 'read-only'
//		operations.
//
//	readerBuffSize					int
//
//		This integer value controls the size of the
//		'read' buffer created for the io.Reader object
//		generated from the file identified by
//		'readerPathFileName'. This io.Reader object is
//		encapsulated in the current instance of
//		FileBufferReadWrite.
//
//		'readerBuffSize' should be configured to maximize
//		performance for 'read' operations subject to
//		prevailing memory limitations.
//
//		The minimum reader buffer size is 16-bytes. If
//		'readerBuffSize' is set to a size less than "16",
//		it will be automatically reset to the default
//		buffer size of 4096-bytes.
//
//	writerPathFileName			string
//
//		This string contains the path and file name of
//		the target 'write' file which will be used as
//		an output data destination for 'write'
//		operations.
//
//		If the target path and file do not currently
//		exist on an attached storage drive, this method
//		will attempt to create them.
//
//	openWriteFileReadWrite		bool
//
//		If this parameter is set to 'true', the target
//		'write' file identified by input parameter
//		'writerPathFileName' will be opened for 'read'
//		and 'write' operations.
//
//		If 'openWriteFileReadWrite' is set to 'false',
//		the target write file will be opened for
//		'write-only' operations.
//
//	writerBuffSize				int
//
//		This integer value controls the size of the
//		'write' buffer created for the io.Writer
//		object generated from the file identified by
//		input parameter 'writerPathFileName'. This
//		io.Writer object is encapsulated in the
//		current instance of FileBufferReadWrite.
//
//		'writerBuffSize' should be configured to maximize
//		performance for 'write' operations subject to
//		prevailing memory limitations.
//
//		If 'writerBuffSize' is set to a value less than
//		or equal to zero (0), it will be automatically
//		reset to the default value of 4096-bytes.
//
//	truncateExistingWriteFile	bool
//
//		If this parameter is set to 'true', the target
//		'write' file ('writerPathFileName') will be
//		opened for write operations. If the target write
//		file previously existed, it will be truncated.
//		This means that the file's previous contents will
//		be deleted.
//
//		If this parameter is set to 'false', the target
//		'write' file will be opened for write operations.
//		If the target 'write' file previously existed,
//		the new text written to this file will be appended
//		to the end of the previous file contents.
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
//	readerFileInfoPlus			FileInfoPlus
//
//		This returned instance of Type FileInfoPlus
//		contains data elements describing the file
//		identified by input parameter
//		'readerPathFileName'.
//
//		Type FileInfoPlus conforms to the os.FileInfo
//		interface. This structure will store os.FileInfo
//	 	information plus additional information related
//	 	to a file or directory.
//
//		type os.FileInfo interface {
//
//				Name() string
//					base name of the file
//
//				Size() int64
//					length in bytes for regular files;
//					system-dependent for others
//
//				Mode() FileMode
//					file mode bits
//
//				ModTime() time.Time
//					modification time
//
//				IsDir() bool
//					abbreviation for Mode().IsDir()
//
//				Sys() any
//					underlying data source (can return nil)
//		}
//
//		See the detailed documentation for Type
//		FileInfoPlus in the source file,
//		'fileinfoplus.go'.
//
//	writerFileInfoPlus			FileInfoPlus
//
//		This returned instance of Type FileInfoPlus
//		contains data elements describing the file
//		identified by input parameter
//		'writerPathFileName'.
//
//		Type FileInfoPlus conforms to the os.FileInfo
//		interface. This structure will store os.FileInfo
//	 	information plus additional information related
//	 	to a file or directory.
//
//		type os.FileInfo interface {
//
//				Name() string
//					base name of the file
//
//				Size() int64
//					length in bytes for regular files;
//					system-dependent for others
//
//				Mode() FileMode
//					file mode bits
//
//				ModTime() time.Time
//					modification time
//
//				IsDir() bool
//					abbreviation for Mode().IsDir()
//
//				Sys() any
//					underlying data source (can return nil)
//		}
//
//		See the detailed documentation for Type
//		FileInfoPlus in the source file,
//		'fileinfoplus.go'.
//
//	err							error
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
func (fBufReadWrite *FileBufferReadWrite) SetPathFileNamesReadWrite(
	readerPathFileName string,
	openReadFileReadWrite bool,
	readerBuffSize int,
	writerPathFileName string,
	openWriteFileReadWrite bool,
	writerBuffSize int,
	truncateExistingWriteFile bool,
	errorPrefix interface{}) (
	readerFileInfoPlus FileInfoPlus,
	writerFileInfoPlus FileInfoPlus,
	err error) {

	if fBufReadWrite.lock == nil {
		fBufReadWrite.lock = new(sync.Mutex)
	}

	fBufReadWrite.lock.Lock()

	defer fBufReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileBufferReadWrite."+
			"SetPathFileNamesReadWrite()",
		"")

	if err != nil {

		return readerFileInfoPlus, writerFileInfoPlus, err
	}

	readerFileInfoPlus,
		writerFileInfoPlus,
		err = new(fileBufferReadWriteNanobot).
		setPathFileNamesReadWrite(
			fBufReadWrite,
			"fBufReadWrite",
			readerPathFileName,
			"readerPathFileName",
			openReadFileReadWrite,
			readerBuffSize,
			writerPathFileName,
			"writerPathFileName",
			openWriteFileReadWrite,
			writerBuffSize,
			truncateExistingWriteFile,
			ePrefix)

	return readerFileInfoPlus, writerFileInfoPlus, err
}

// SetPathFileNameReader
//
// Receives an input parameter string specifying the path
// and file name identifying the file which will be
// configured as a data source for 'read' operations.
// This file will be configured as an internal io.Reader
// object for the current instance of
// FileBufferReadWrite.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reconfigure
//	the member variable io.Reader object encapsulated in
//	the current instance of FileBufferReadWrite.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	readerPathFileName			string
//
//		This string contains the path and file name of
//		the file which will be configured as an io.Reader
//		object encapsulated in the current instance of
//		FileBufferReadWrite. As such, the file identified
//		by 'readerPathFileName' will be used a data source
//		for 'read' operations.
//
//		If this file does not currently exist on an
//		attached storage drive, an error will be
//		returned.
//
//	openReadFileReadWrite		bool
//
//		If this parameter is set to 'true', the target
//		'read' file identified from input parameter
//		'readerPathFileName' will be opened for both
//		'read' and 'write' operations.
//
//		If 'openReadFileReadWrite' is set to 'false', the
//		target 'read' file will be opened for 'read-only'
//		operations.
//
//	readerBuffSize					int
//
//		This integer value controls the size of the
//		'read' buffer created for the io.Reader object
//		generated from the file identified by
//		'readerPathFileName'. This io.Reader object is
//		encapsulated in the current instance of
//		FileBufferReadWrite.
//
//		'readerBuffSize' should be configured to maximize
//		performance for 'read' operations subject to
//		prevailing memory limitations.
//
//		The minimum read buffer size is 1-byte. If
//		'bufSize' is set to a size less than or equal to
//		zero, it will be automatically reset to the
//		default buffer size of 4096-bytes.
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
//	readerFileInfoPlus			FileInfoPlus
//
//		This returned instance of Type FileInfoPlus
//		contains data elements describing the file
//		identified by input parameter
//		'readerPathFileName'.
//
//		Type FileInfoPlus conforms to the os.FileInfo
//		interface. This structure will store os.FileInfo
//	 	information plus additional information related
//	 	to a file or directory.
//
//		type os.FileInfo interface {
//
//				Name() string
//					base name of the file
//
//				Size() int64
//					length in bytes for regular files;
//					system-dependent for others
//
//				Mode() FileMode
//					file mode bits
//
//				ModTime() time.Time
//					modification time
//
//				IsDir() bool
//					abbreviation for Mode().IsDir()
//
//				Sys() any
//					underlying data source (can return nil)
//		}
//
//		See the detailed documentation for Type
//		FileInfoPlus in the source file,
//		'fileinfoplus.go'.
//
//	err							error
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
func (fBufReadWrite *FileBufferReadWrite) SetPathFileNameReader(
	readerPathFileName string,
	openReadFileReadWrite bool,
	readerBuffSize int,
	errorPrefix interface{}) (
	readerFileInfoPlus FileInfoPlus,
	err error) {

	if fBufReadWrite.lock == nil {
		fBufReadWrite.lock = new(sync.Mutex)
	}

	fBufReadWrite.lock.Lock()

	defer fBufReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileBufferReadWrite."+
			"SetPathFileNameReader()",
		"")

	if err != nil {

		return readerFileInfoPlus, err
	}

	readerFileInfoPlus,
		err = new(fileBufferReadWriteAtom).
		setPathFileNameReader(
			fBufReadWrite,
			"fBufReadWrite",
			readerPathFileName,
			"readerPathFileName",
			openReadFileReadWrite,
			readerBuffSize,
			ePrefix)

	return readerFileInfoPlus, err
}

// SetPathFileNameWriter
//
// Receives an input parameter string specifying the path
// and file name identifying the file which will be
// configured as an output destination for 'write'
// operations. This file will be configured as an
// internal io.Writer object for the FileBufferReadWrite
// instance passed as input parameter 'fBufReadWrite'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reconfigure
//	the member variable io.Writer object encapsulated in
//	the instance of FileBufferReadWrite passed as input
//	parameter 'fBufReadWrite'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	writerPathFileName			string
//
//		This string contains the path and file name of
//		the target 'write' file which will be used as
//		a data destination for 'write' operations.
//
//		The target 'write' file will be configured
//		as an io.Writer object encapsulated in the
//		current FileBufferReadWrite instance.
//
//		If the target path and file do not currently
//		exist on an attached storage drive, this method
//		will attempt to create them.
//
//	openWriteFileReadWrite		bool
//
//		If this parameter is set to 'true', the target
//		'write' file identified by input parameter
//		'writerPathFileName' will be opened for 'read'
//		and 'write' operations.
//
//		If 'openWriteFileReadWrite' is set to 'false',
//		the target write file will be opened for
//		'write-only' operations.
//
//	writerBuffSize				int
//
//		This integer value controls the size of the
//		'write' buffer created for the io.Writer
//		object generated from the file identified by
//		input parameter 'writerPathFileName'. This
//		io.Writer object is encapsulated in the
//		current instance of FileBufferReadWrite.
//
//		'writerBuffSize' should be configured to maximize
//		performance for 'write' operations subject to
//		prevailing memory limitations.
//
//		If 'writerBuffSize' is set to a value less than
//		or equal to zero (0), it will be automatically
//		reset to the default value of 4096-bytes.
//
//	truncateExistingWriteFile	bool
//
//		If this parameter is set to 'true', the target
//		'write' file ('writerPathFileName') will be
//		opened for write operations. If the target write
//		file previously existed, it will be truncated.
//		This means that the file's previous contents will
//		be deleted.
//
//		If this parameter is set to 'false', the target
//		'write' file will be opened for write operations.
//		If the target 'write' file previously existed,
//		the new text written to this file will be appended
//		to the end of the previous file contents.
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
//	writerFileInfoPlus			FileInfoPlus
//
//		This returned instance of Type FileInfoPlus
//		contains data elements describing the file
//		identified by input parameter
//		'writerPathFileName'.
//
//		Type FileInfoPlus conforms to the os.FileInfo
//		interface. This structure will store os.FileInfo
//	 	information plus additional information related
//	 	to a file or directory.
//
//		type os.FileInfo interface {
//
//				Name() string
//					base name of the file
//
//				Size() int64
//					length in bytes for regular files;
//					system-dependent for others
//
//				Mode() FileMode
//					file mode bits
//
//				ModTime() time.Time
//					modification time
//
//				IsDir() bool
//					abbreviation for Mode().IsDir()
//
//				Sys() any
//					underlying data source (can return nil)
//		}
//
//		See the detailed documentation for Type
//		FileInfoPlus in the source file,
//		'fileinfoplus.go'.
//
//	err							error
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
func (fBufReadWrite *FileBufferReadWrite) SetPathFileNameWriter(
	writerPathFileName string,
	openWriteFileReadWrite bool,
	writerBuffSize int,
	truncateExistingWriteFile bool,
	errorPrefix interface{}) (
	writerFileInfoPlus FileInfoPlus,
	err error) {

	if fBufReadWrite.lock == nil {
		fBufReadWrite.lock = new(sync.Mutex)
	}

	fBufReadWrite.lock.Lock()

	defer fBufReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileBufferReadWrite."+
			"SetPathFileNameWriter()",
		"")

	if err != nil {

		return writerFileInfoPlus, err
	}

	writerFileInfoPlus,
		err = new(fileBufferReadWriteAtom).
		setPathFileNameWriter(
			fBufReadWrite,
			"fBufReadWrite",
			writerPathFileName,
			"writerPathFileName",
			openWriteFileReadWrite,
			writerBuffSize,
			truncateExistingWriteFile,
			ePrefix)

	return writerFileInfoPlus, err
}

type fileBufferReadWriteMicrobot struct {
	lock *sync.Mutex
}

// closeReaderWriter
//
// This method is designed to perform clean up tasks
// after completion of all 'read' and 'write' operations
// associated with the instance of FileBufferReadWrite
// passed as input parameter 'fBufReadWrite'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	This method will effectively render the
//		FileBufferReadWrite instance passed as
//		'fBufReadWrite' invalid and unusable for any
//		future 'read' and/or 'write' operations.
//
//	(2) After completing all 'read' and 'write'
//		operations, users MUST perform these required
//		clean-up tasks.
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
//	 	text passed by input parameter, 'errPrefDto'.
//	 	The 'errPrefDto' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fBufReadWriteMicrobot *fileBufferReadWriteMicrobot) closeReaderWriter(
	fBufReadWrite *FileBufferReadWrite,
	fBufReadWriteLabel string,
	errorPrefix interface{}) error {

	if fBufReadWriteMicrobot.lock == nil {
		fBufReadWriteMicrobot.lock = new(sync.Mutex)
	}

	fBufReadWriteMicrobot.lock.Lock()

	defer fBufReadWriteMicrobot.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileBufferReadWriteMicrobot." +
		"closeReaderWriter()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {

		return err
	}

	if len(fBufReadWriteLabel) == 0 {

		fBufReadWriteLabel = "fBufReadWrite"
	}

	if fBufReadWrite == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is a nil pointer!\n"+
			"%v is invalid.\n",
			ePrefix.String(),
			fBufReadWriteLabel,
			fBufReadWriteLabel)

		return err
	}

	var errs []error
	var err2 error
	var fBuffReadWriteElectron = new(fileBufferReadWriteElectron)

	err2 = fBuffReadWriteElectron.
		closeReader(
			fBufReadWrite,
			fBufReadWriteLabel,
			ePrefix)

	if err2 != nil {
		errs = append(
			errs,
			fmt.Errorf("%v", err2.Error()))
	}

	err2 = fBuffReadWriteElectron.
		flushAndCloseWriter(
			fBufReadWrite,
			fBufReadWriteLabel,
			ePrefix)

	if err2 != nil {
		errs = append(
			errs,
			fmt.Errorf("%v", err2.Error()))
	}

	if len(errs) > 0 {

		err = new(StrMech).ConsolidateErrors(errs)

	}

	return err
}

// setFileMgrsReadWrite
//
// Receives two instances of FileMgr as input parameters
// identifying the io.Reader and io.Writer objects which
// will be configured for the FileBufferReadWrite
// instance passed as input parameter 'fBufReadWrite'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the instance of
//	FileBufferReadWrite passed as input parameter
//	'fBufReadWrite'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fBufReadWrite				*FileBufferReadWrite
//
//		A pointer to an instance of FileBufferWriter.
//
//		The internal FileBufferReader and
//		FileBufferWriter objects encapsulated in this
//		instance be deleted and reinitialized using the
//		io.Reader and io.Writer objects passed as input
//		parameters 'reader' and 'writer'.
//
//	fBufReadWriteLabel			string
//
//		The name or label associated with input parameter
//		'fBufReadWrite' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fBufReadWrite" will
//		be automatically applied.
//
//	readerFileMgr				*FileMgr
//
//		A pointer to an instance of FileMgr. The file
//		identified by 'readerFileMgr' will be used as a
//		data source for 'read' operations performed by
//		the instance of FileBufferReadWrite passed as
//		input parameter 'fBufReadWrite'.
//
//		If the path and file name encapsulated by
//		'readerFileMgr' do not currently exist on an
//		attached storage drive, an error will be
//		returned.
//
//	readerFileMgrLabel			string
//
//		The name or label associated with input parameter
//		'readerFileMgr' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "readerFileMgr" will
//		be automatically applied.
//
//	openReadFileReadWrite		bool
//
//		If this parameter is set to 'true', the target
//		'read' file identified by input parameter
//		'fileMgr' will be opened for 'read' and 'write'
//		operations.
//
//		If 'openFileReadWrite' is set to 'false', the
//		target read file will be opened for 'read-only'
//		operations.
//
//	readerBuffSize					int
//
//		This integer value controls the size of the
//		'read' buffer created for the io.Reader
//		associated with the file identified by
//		'readerFileMgr'.
//
//		'readerBuffSize' should be configured to maximize
//		performance for 'read' operations subject to
//		prevailing memory limitations.
//
//		The minimum reader buffer size is 16-bytes. If
//		'bufSize' is set to a size less than "16", it
//		will be automatically reset to the default buffer
//		size of 4096.
//
//	writerFileMgr				*FileMgr
//
//		A pointer to an instance of FileMgr. The file
//		identified by 'writerFileMgr' will be used as an
//		output destination for 'write' operations performed by
//		the instance of FileBufferReadWrite passed as
//		input parameter 'fBufReadWrite'.
//
//		If the path and file name encapsulated by
//		'fileMgr' do not currently exist on an attached
//		storage drive, this method will attempt to create
//		them.
//
//	writerFileMgrLabel			string
//
//		The name or label associated with input parameter
//		'writerFileMgr' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "writerFileMgr" will
//		be automatically applied.
//
//	openWriteFileReadWrite		bool
//
//		If this parameter is set to 'true', the target
//		'write' file identified by input parameter
//		'writerFileMgr' will be opened for 'read'
//		and 'write' operations.
//
//		If 'openWriteFileReadWrite' is set to 'false',
//		the target write file will be opened for
//		'write-only' operations.
//
//	writerBuffSize				int
//
//		This integer value controls the size of the
//		'write' buffer created for the io.Writer object
//		generated from the file identified by
//		'writerFileMgr'. This io.Writer object will in
//		turn be configured and encapsulated in the
//		FileBufferWriter instance passed as input
//		parameter 'fBufReadWrite'.
//
//		'writerBuffSize' should be configured to maximize
//		performance for 'write' operations subject to
//		prevailing memory limitations.
//
//		If 'writerBuffSize' is set to a value less than
//		or equal to zero (0), it will be automatically
//		reset to the default value of 4096-bytes.
//
//	truncateExistingWriteFile	bool
//
//		If this parameter is set to 'true', the target
//		'write' file ('writerFileMgr') will be opened for
//		'write' operations. If the target 'write' file
//		previously existed, it will be truncated. This
//		means that the file's previous contents will be
//		deleted.
//
//		If this parameter is set to 'false', the target
//		'write' file will be opened for write operations.
//		If the target 'write' file previously existed,
//		the new text written to this file will be appended
//		to the end of the previous file contents.
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
//	readerFileInfoPlus			FileInfoPlus
//
//		This returned instance of Type FileInfoPlus
//		contains data elements describing the file
//		identified by input parameter 'readerFileMgr'.
//
//		Type FileInfoPlus conforms to the os.FileInfo
//		interface. This structure will store os.FileInfo
//	 	information plus additional information related
//	 	to a file or directory.
//
//		type os.FileInfo interface {
//
//				Name() string
//					base name of the file
//
//				Size() int64
//					length in bytes for regular files;
//					system-dependent for others
//
//				Mode() FileMode
//					file mode bits
//
//				ModTime() time.Time
//					modification time
//
//				IsDir() bool
//					abbreviation for Mode().IsDir()
//
//				Sys() any
//					underlying data source (can return nil)
//		}
//
//		See the detailed documentation for Type
//		FileInfoPlus in the source file,
//		'fileinfoplus.go'.
//
//	writerFileInfoPlus			FileInfoPlus
//
//		This returned instance of Type FileInfoPlus
//		contains data elements describing the file
//		identified by input parameter 'writerFileMgr'.
//
//		Type FileInfoPlus conforms to the os.FileInfo
//		interface. This structure will store os.FileInfo
//	 	information plus additional information related
//	 	to a file or directory.
//
//		type os.FileInfo interface {
//
//				Name() string
//					base name of the file
//
//				Size() int64
//					length in bytes for regular files;
//					system-dependent for others
//
//				Mode() FileMode
//					file mode bits
//
//				ModTime() time.Time
//					modification time
//
//				IsDir() bool
//					abbreviation for Mode().IsDir()
//
//				Sys() any
//					underlying data source (can return nil)
//		}
//
//		See the detailed documentation for Type
//		FileInfoPlus in the source file,
//		'fileinfoplus.go'.
//
//	err							error
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
func (fBufReadWriteMicrobot *fileBufferReadWriteMicrobot) setFileMgrsReadWrite(
	fBufReadWrite *FileBufferReadWrite,
	fBufReadWriteLabel string,
	readerFileMgr *FileMgr,
	readerFileMgrLabel string,
	openReadFileReadWrite bool,
	readerBuffSize int,
	writerFileMgr *FileMgr,
	writerFileMgrLabel string,
	openWriteFileReadWrite bool,
	writerBuffSize int,
	truncateExistingWriteFile bool,
	errPrefDto *ePref.ErrPrefixDto) (
	readerFileInfoPlus FileInfoPlus,
	writerFileInfoPlus FileInfoPlus,
	err error) {

	if fBufReadWriteMicrobot.lock == nil {
		fBufReadWriteMicrobot.lock = new(sync.Mutex)
	}

	fBufReadWriteMicrobot.lock.Lock()

	defer fBufReadWriteMicrobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileBufferReadWriteMicrobot." +
		"setFileMgrsReadWrite()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		return readerFileInfoPlus, writerFileInfoPlus, err
	}

	if len(fBufReadWriteLabel) == 0 {

		fBufReadWriteLabel = "fBufReadWrite"
	}

	if len(readerFileMgrLabel) == 0 {

		readerFileMgrLabel = "readerFileMgr"
	}

	if len(writerFileMgrLabel) == 0 {

		writerFileMgrLabel = "writerFileMgr"
	}

	if fBufReadWrite == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is a nil pointer!\n"+
			"%v is invalid.\n",
			ePrefix.String(),
			fBufReadWriteLabel,
			fBufReadWriteLabel)

		return readerFileInfoPlus, writerFileInfoPlus, err
	}

	var fBuffReadWriteMolecule = new(fileBufferReadWriteMolecule)

	readerFileInfoPlus,
		err = fBuffReadWriteMolecule.
		setFileMgrReader(
			fBufReadWrite,
			fBufReadWriteLabel,
			readerFileMgr,
			readerFileMgrLabel,
			openReadFileReadWrite,
			readerBuffSize,
			ePrefix)

	if err != nil {

		return readerFileInfoPlus, writerFileInfoPlus, err
	}

	writerFileInfoPlus,
		err = fBuffReadWriteMolecule.
		setFileMgrWriter(
			fBufReadWrite,
			fBufReadWriteLabel,
			writerFileMgr,
			writerFileMgrLabel,
			openWriteFileReadWrite,
			writerBuffSize,
			truncateExistingWriteFile,
			ePrefix)

	return readerFileInfoPlus, writerFileInfoPlus, err
}

type fileBufferReadWriteNanobot struct {
	lock *sync.Mutex
}

// setIoReaderWriter
//
// This 'setter' method is used to initialize new values
// for internal member variables contained in the
// instance of FileBufferReadWrite passed as input
// parameter 'fBufReadWrite'. The new configuration will
// be based on the io.Reader and io.Writer object passed
// as input parameters.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the instance of
//	FileBufferReadWrite passed as input parameter
//	'fBufReadWrite'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fBufReadWrite				*FileBufferReadWrite
//
//		A pointer to an instance of FileBufferWriter.
//
//		The internal FileBufferReader and
//		FileBufferWriter objects encapsulated in this
//		instance be deleted and reinitialized using the
//		io.Reader and io.Writer objects passed as input
//		parameters 'reader' and 'writer'.
//
//	fBufReadWriteLabel			string
//
//		The name or label associated with input parameter
//		'fBufReadWrite' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fBufReadWrite" will
//		be automatically applied.
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
//		in addition to calling the local method:
//
//		FileBufferReadWrite.CloseFileBufferReadWrite()
//
//		While the 'read' services provided by
//		FileBufferReadWrite are primarily designed to
//		read data from disk files, this type of 'reader'
//		will in fact read data from any object
//		implementing the io.Reader interface.
//
//	readerLabel					string
//
//		The name or label associated with input parameter
//		'reader' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "reader" will
//		be automatically applied.
//
//	readerBuffSize				int
//
//		This integer value controls the size of the
//		buffer created for the io.Reader object passed as
//		input parameter 'reader'.
//
//		'readerBuffSize' should be configured to maximize
//		performance for 'read' operations subject to
//		prevailing memory limitations.
//
//		The minimum reader buffer size is 16-bytes. If
//		'readerBuffSize' is set to a value less than
//		"16", it will be automatically reset to the
//		default buffer size of 4096-bytes.
//
//	writer						io.Writer
//
//		This parameter will accept any object
//		implementing the io.Writer interface.
//
//		This object may be a file pointer of type
//		*os.File. File pointers of this type implement
//		the io.Writer interface.
//
//		A file pointer (*os.File) will facilitate writing
//		output data to destination files residing on an
//		attached storage drive. However, with this
//		configuration, the user is responsible for
//		manually closing the file and performing any
//		other required clean-up operations in addition to
//		calling local method:
//
//		FileBufferReadWrite.CloseFileBufferReadWrite()
//
//		While the 'write' services provided by the
//		FileBufferReadWrite are primarily designed for
//		writing data to disk files, this type of 'writer'
//		will in fact write data to any object
//		implementing the io.Writer interface.
//
//	writerLabel					string
//
//		The name or label associated with input parameter
//		'writer' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "writer" will be
//		automatically applied.
//
//	writerBuffSize				int
//
//		This integer value controls the size of the
//		'write' buffer created for the io.Writer
//		object passed as input parameter 'writer'.
//		This io.Writer object will in turn be configured
//		and encapsulated in the FileBufferWriter
//		instance passed as input parameter 'fBufReadWrite'.
//
//		'writerBuffSize' should be configured to maximize
//		performance for 'write' operations subject to
//		prevailing memory limitations.
//
//		The minimum write buffer size is 1-byte. If
//		'writerBuffSize' is set to a size less than or
//		equal to zero, it will be automatically reset to
//		the default buffer size of 4096-bytes.
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
func (fBufReadWriteNanobot *fileBufferReadWriteNanobot) setIoReaderWriter(
	fBufReadWrite *FileBufferReadWrite,
	fBufReadWriteLabel string,
	reader io.Reader,
	readerLabel string,
	readerBuffSize int,
	writer io.Writer,
	writerLabel string,
	writerBuffSize int,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fBufReadWriteNanobot.lock == nil {
		fBufReadWriteNanobot.lock = new(sync.Mutex)
	}

	fBufReadWriteNanobot.lock.Lock()

	defer fBufReadWriteNanobot.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileBufferReadWriteNanobot." +
		"setIoReaderWriter()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return err
	}

	var fBufReadWriteMolecule = new(fileBufferReadWriteAtom)

	err = fBufReadWriteMolecule.
		setIoReader(
			fBufReadWrite,
			fBufReadWriteLabel,
			reader,
			readerLabel,
			readerBuffSize,
			ePrefix)

	if err != nil {
		return err
	}

	err = fBufReadWriteMolecule.
		setIoWriter(
			fBufReadWrite,
			fBufReadWriteLabel,
			writer,
			writerLabel,
			writerBuffSize,
			ePrefix)

	return err
}

// setPathFileNamesReadWrite
//
// Receives two strings as input parameters for the path
// and file names identifying the io.Reader and io.Writer
// objects which will be configured for the
// FileBufferReadWrite instance passed as input parameter
// 'fBufReadWrite'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the instance of
//	FileBufferReadWrite passed as input parameter
//	'fBufReadWrite'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fBufReadWrite				*FileBufferReadWrite
//
//		A pointer to an instance of FileBufferWriter.
//
//		The internal FileBufferReader and
//		FileBufferWriter objects encapsulated in this
//		instance be deleted and reinitialized using the
//		io.Reader and io.Writer objects passed as input
//		parameters 'reader' and 'writer'.
//
//	fBufReadWriteLabel			string
//
//		The name or label associated with input parameter
//		'fBufReadWrite' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fBufReadWrite" will
//		be automatically applied.
//
//	readerPathFileName			string
//
//		This string contains the path and file name of
//		the file which will be configured as an io.Reader
//		instance and used a data source for 'read'
//		operations.
//
//		If this file does not currently exist on an
//		attached storage drive, an error will be
//		returned.
//
//	openReadFileReadWrite		bool
//
//		If this parameter is set to 'true', the target
//		'read' file identified from input parameter
//		'readerPathFileName' will be opened for both
//		'read' and 'write' operations.
//
//		If 'openReadFileReadWrite' is set to 'false', the
//		target 'read' file will be opened for 'read-only'
//		operations.
//
//	readerBuffSize					int
//
//		This integer value controls the size of the
//		'read' buffer created for the io.Reader
//		associated with the file identified by
//		'readerPathFileName'.
//
//		'readerBuffSize' should be configured to maximize
//		performance for 'read' operations subject to
//		prevailing memory limitations.
//
//		The minimum read buffer size is 1-byte. If
//		'bufSize' is set to a size less than or equal to
//		zero, it will be automatically set to the default
//		buffer size of 4096-bytes.
//
//	writerPathFileName			string
//
//		This string contains the path and file name of
//		the target 'write' file which will be used as
//		a data destination for 'write' operations.
//
//		If the target path and file do not currently
//		exist on an attached storage drive, this method
//		will attempt to create them.
//
//	openWriteFileReadWrite		bool
//
//		If this parameter is set to 'true', the target
//		'write' file identified by input parameter
//		'writerPathFileName' will be opened for 'read'
//		and 'write' operations.
//
//		If 'openWriteFileReadWrite' is set to 'false',
//		the target write file will be opened for
//		'write-only' operations.
//
//	writerBuffSize				int
//
//		This integer value controls the size of the
//		'write' buffer created for the io.Writer
//		object generated from the file identified by
//		input parameter 'writerPathFileName'. This
//		io.Writer object is encapsulated in the
//		FileBufferReadWrite instance passed as input
//		parameter 'fBufReadWrite'.
//
//		'writerBuffSize' should be configured to maximize
//		performance for 'write' operations subject to
//		prevailing memory limitations.
//
//		If 'writerBuffSize' is set to a value less than
//		or equal to zero (0), it will be automatically
//		reset to the default value of 4096-bytes.
//
//	truncateExistingWriteFile	bool
//
//		If this parameter is set to 'true', the target
//		'write' file ('writerPathFileName') will be
//		opened for write operations. If the target write
//		file previously existed, it will be truncated.
//		This means that the file's previous contents will
//		be deleted.
//
//		If this parameter is set to 'false', the target
//		'write' file will be opened for write operations.
//		If the target 'write' file previously existed,
//		the new text written to this file will be appended
//		to the end of the previous file contents.
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
//	readerFileInfoPlus			FileInfoPlus
//
//		This returned instance of Type FileInfoPlus
//		contains data elements describing the file
//		identified by input parameter
//		'readerPathFileName'.
//
//		Type FileInfoPlus conforms to the os.FileInfo
//		interface. This structure will store os.FileInfo
//	 	information plus additional information related
//	 	to a file or directory.
//
//		type os.FileInfo interface {
//
//				Name() string
//					base name of the file
//
//				Size() int64
//					length in bytes for regular files;
//					system-dependent for others
//
//				Mode() FileMode
//					file mode bits
//
//				ModTime() time.Time
//					modification time
//
//				IsDir() bool
//					abbreviation for Mode().IsDir()
//
//				Sys() any
//					underlying data source (can return nil)
//		}
//
//		See the detailed documentation for Type
//		FileInfoPlus in the source file,
//		'fileinfoplus.go'.
//
//	writerFileInfoPlus			FileInfoPlus
//
//		This returned instance of Type FileInfoPlus
//		contains data elements describing the file
//		identified by input parameter
//		'writerPathFileName'.
//
//		Type FileInfoPlus conforms to the os.FileInfo
//		interface. This structure will store os.FileInfo
//	 	information plus additional information related
//	 	to a file or directory.
//
//		type os.FileInfo interface {
//
//				Name() string
//					base name of the file
//
//				Size() int64
//					length in bytes for regular files;
//					system-dependent for others
//
//				Mode() FileMode
//					file mode bits
//
//				ModTime() time.Time
//					modification time
//
//				IsDir() bool
//					abbreviation for Mode().IsDir()
//
//				Sys() any
//					underlying data source (can return nil)
//		}
//
//		See the detailed documentation for Type
//		FileInfoPlus in the source file,
//		'fileinfoplus.go'.
//
//	err							error
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
func (fBufReadWriteNanobot *fileBufferReadWriteNanobot) setPathFileNamesReadWrite(
	fBufReadWrite *FileBufferReadWrite,
	fBufReadWriteLabel string,
	readerPathFileName string,
	readerPathFileNameLabel string,
	openReadFileReadWrite bool,
	readerBuffSize int,
	writerPathFileName string,
	writerPathFileNameLabel string,
	openWriteFileReadWrite bool,
	writerBuffSize int,
	truncateExistingWriteFile bool,
	errPrefDto *ePref.ErrPrefixDto) (
	readerFileInfoPlus FileInfoPlus,
	writerFileInfoPlus FileInfoPlus,
	err error) {

	if fBufReadWriteNanobot.lock == nil {
		fBufReadWriteNanobot.lock = new(sync.Mutex)
	}

	fBufReadWriteNanobot.lock.Lock()

	defer fBufReadWriteNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileBufferReadWriteNanobot." +
		"setPathFileNamesReadWrite()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		return readerFileInfoPlus, writerFileInfoPlus, err
	}

	var fBufReadWriteMolecule = new(fileBufferReadWriteAtom)

	readerFileInfoPlus,
		err = fBufReadWriteMolecule.
		setPathFileNameReader(
			fBufReadWrite,
			fBufReadWriteLabel,
			readerPathFileName,
			readerPathFileNameLabel,
			openReadFileReadWrite,
			readerBuffSize,
			ePrefix)

	if err != nil {

		return readerFileInfoPlus, writerFileInfoPlus, err
	}

	writerFileInfoPlus,
		err = fBufReadWriteMolecule.
		setPathFileNameWriter(
			fBufReadWrite,
			fBufReadWriteLabel,
			writerPathFileName,
			writerPathFileNameLabel,
			openWriteFileReadWrite,
			writerBuffSize,
			truncateExistingWriteFile,
			ePrefix)

	return readerFileInfoPlus, writerFileInfoPlus, err
}

type fileBufferReadWriteMolecule struct {
	lock *sync.Mutex
}

// setFileMgrReader
//
// Receives an instance of File Manager (FileMgr) which
// will be used to configure an io.Reader object
// encapsulated by the FileBufferReadWrite instance
// passed as input parameter 'fBufReadWrite'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reconfigure
//	the member variable io.Reader object encapsulated in
//	the instance of FileBufferReadWrite passed as input
//	parameter 'fBufReadWrite'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fBufReadWrite				*FileBufferReadWrite
//
//		A pointer to an instance of FileBufferWriter.
//
//		The internal io.Reader object encapsulated in
//		this instance of FileBufferReadWrite will be
//		deleted and reconfigured using the FileMgr
//		instance passed as input parameter
//		'readerFileMgr'.
//
//	fBufReadWriteLabel			string
//
//		The name or label associated with input parameter
//		'fBufReadWrite' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fBufReadWrite" will
//		be automatically applied.
//
//	readerFileMgr				*FileMgr
//
//		A pointer to an instance of FileMgr. The file
//		identified by 'readerFileMgr' will be used as a
//		data source for 'read' operations and will be
//		configured as an internal io.Reader for the
//		FileBufferReadWrite instance passed as input
//		parameter 'fBufReadWrite'.
//
//		If the path and file name encapsulated by
//		'readerFileMgr' do not currently exist on an
//		attached storage drive, an error will be
//		returned.
//
//	readerLabel					string
//
//		The name or label associated with input parameter
//		'reader' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "reader" will
//		be automatically applied.
//
//	openReadFileReadWrite		bool
//
//		If this parameter is set to 'true', the target
//		'read' file identified by input parameter
//		'readerFileMgr' will be opened for both 'read'
//		and 'write' operations.
//
//		If 'openReadFileReadWrite' is set to 'false', the
//		target 'read' file will be opened for 'read-only'
//		operations.
//
//	readerBuffSize				int
//
//		This integer value controls the size of the
//		'read' buffer created for the internal io.Reader
//		encapsulated in the instance of
//		FileBufferReadWrite passed as input parameter
//		'fBufReadWrite'.
//
//		'readerBuffSize' should be configured to maximize
//		performance for 'read' operations subject to
//		prevailing memory limitations.
//
//		The minimum reader buffer size is 16-bytes. If
//		'readerBuffSize' is set to a size less than "16",
//		it will be automatically reset to the default
//		buffer size of 4096-bytes.
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
//	fileInfoPlus				FileInfoPlus
//
//		This returned instance of Type FileInfoPlus
//		contains data elements describing the file
//		identified by input parameter 'readerFileMgr'.
//
//		Type FileInfoPlus conforms to the os.FileInfo
//		interface. This structure will store os.FileInfo
//	 	information plus additional information related
//	 	to a file or directory.
//
//		type os.FileInfo interface {
//
//				Name() string
//					base name of the file
//
//				Size() int64
//					length in bytes for regular files;
//					system-dependent for others
//
//				Mode() FileMode
//					file mode bits
//
//				ModTime() time.Time
//					modification time
//
//				IsDir() bool
//					abbreviation for Mode().IsDir()
//
//				Sys() any
//					underlying data source (can return nil)
//		}
//
//		See the detailed documentation for Type
//		FileInfoPlus in the source file,
//		'fileinfoplus.go'.
//
//	err							error
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
func (fBuffReadWriteMolecule *fileBufferReadWriteMolecule) setFileMgrReader(
	fBufReadWrite *FileBufferReadWrite,
	fBufReadWriteLabel string,
	readerFileMgr *FileMgr,
	readerFileMgrLabel string,
	openReadFileReadWrite bool,
	readerBuffSize int,
	errPrefDto *ePref.ErrPrefixDto) (
	fInfoPlus FileInfoPlus,
	err error) {

	if fBuffReadWriteMolecule.lock == nil {
		fBuffReadWriteMolecule.lock = new(sync.Mutex)
	}

	fBuffReadWriteMolecule.lock.Lock()

	defer fBuffReadWriteMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileBufferReadWriteMolecule." +
		"setFileMgrReader()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		return fInfoPlus, err
	}

	if len(fBufReadWriteLabel) == 0 {

		fBufReadWriteLabel = "fBufReadWrite"
	}

	if len(readerFileMgrLabel) == 0 {

		readerFileMgrLabel = "readerFileMgr"
	}

	if fBufReadWrite == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is a nil pointer!\n"+
			"%v is invalid.\n",
			ePrefix.String(),
			fBufReadWriteLabel,
			fBufReadWriteLabel)

		return fInfoPlus, err
	}

	var err2 error

	err2 = new(fileMgrHelperAtom).isFileMgrValid(
		readerFileMgr,
		ePrefix.XCpy(readerFileMgrLabel))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter %v is invalid.\n"+
			"%v failed the validity test.\n"+
			"Error=\n%v\n",
			funcName,
			readerFileMgrLabel,
			readerFileMgrLabel,
			err2.Error())

		return fInfoPlus, err
	}

	err2 = new(fileMgrHelper).closeFile(
		readerFileMgr,
		ePrefix.XCpy(readerFileMgrLabel))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"An error occurred while closing the file pointer.\n"+
			"for FileMgr input parameter '%v'.\n"+
			"Error=\n%v\n",
			funcName,
			readerFileMgrLabel,
			err2.Error())

		return fInfoPlus, err
	}

	readerFileMgrLabel += ".absolutePathFileName"

	fInfoPlus,
		err = new(fileBufferReadWriteAtom).
		setPathFileNameReader(
			fBufReadWrite,
			fBufReadWriteLabel,
			readerFileMgr.absolutePathFileName,
			readerFileMgrLabel,
			openReadFileReadWrite,
			readerBuffSize,
			ePrefix)

	return fInfoPlus, err
}

// setFileMgrWriter
//
// Receives an instance of File Manager (FileMgr) which
// will be used to configure an io.Writer object
// encapsulated by the FileBufferReadWrite instance
// passed as input parameter 'fBufReadWrite'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reconfigure
//	the member variable io.Writer object encapsulated in
//	the instance of FileBufferReadWrite passed as input
//	parameter 'fBufReadWrite'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fBufReadWrite				*FileBufferReadWrite
//
//		A pointer to an instance of FileBufferWriter.
//
//		The internal io.Writer object encapsulated in
//		this instance of FileBufferReadWrite will be
//		deleted and reconfigured using the FileMgr
//		instance passed as input parameter
//		'writerFileMgr'.
//
//	fBufReadWriteLabel			string
//
//		The name or label associated with input parameter
//		'fBufReadWrite' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fBufReadWrite" will
//		be automatically applied.
//
//	writerFileMgr				*FileMgr
//
//		A pointer to an instance of FileMgr. The file
//		identified by 'writerFileMgr' will be used as an
//		output destination for 'write' operations
//		performed by the instance of FileBufferReadWrite
//		passed as input parameter 'fBufReadWrite'.
//
//		If the path and file name encapsulated by
//		'fileMgr' do not currently exist on an attached
//		storage drive, this method will attempt to create
//		them.
//
//	writerFileMgrLabel			string
//
//		The name or label associated with input parameter
//		'writerFileMgr' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "writerFileMgr" will
//		be automatically applied.
//
//	openWriteFileReadWrite		bool
//
//		If this parameter is set to 'true', the target
//		'write' file identified by input parameter
//		'writerFileMgr' will be opened for 'read'
//		and 'write' operations.
//
//		If 'openWriteFileReadWrite' is set to 'false',
//		the target write file will be opened for
//		'write-only' operations.
//
//	writerBuffSize				int
//
//		This integer value controls the size of the
//		'write' buffer created for the io.Writer object
//		generated from the file identified by
//		'writerFileMgr'. This io.Writer object will in
//		turn be configured and encapsulated in the
//		FileBufferWriter instance passed as input
//		parameter 'fBufReadWrite'.
//
//		'writerBuffSize' should be configured to maximize
//		performance for 'write' operations subject to
//		prevailing memory limitations.
//
//		If 'writerBuffSize' is set to a value less than
//		or equal to zero (0), it will be automatically
//		reset to the default value of 4096-bytes.
//
//	truncateExistingWriteFile	bool
//
//		If this parameter is set to 'true', the target
//		'write' file ('writerFileMgr') will be opened for
//		'write' operations. If the target 'write' file
//		previously existed, it will be truncated. This
//		means that the file's previous contents will be
//		deleted.
//
//		If this parameter is set to 'false', the target
//		'write' file will be opened for write operations.
//		If the target 'write' file previously existed,
//		the new text written to this file will be appended
//		to the end of the previous file contents.
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
//	fileInfoPlus				FileInfoPlus
//
//		This returned instance of Type FileInfoPlus
//		contains data elements describing the file
//		identified by input parameter 'writerFileMgr'.
//
//		Type FileInfoPlus conforms to the os.FileInfo
//		interface. This structure will store os.FileInfo
//	 	information plus additional information related
//	 	to a file or directory.
//
//		type os.FileInfo interface {
//
//				Name() string
//					base name of the file
//
//				Size() int64
//					length in bytes for regular files;
//					system-dependent for others
//
//				Mode() FileMode
//					file mode bits
//
//				ModTime() time.Time
//					modification time
//
//				IsDir() bool
//					abbreviation for Mode().IsDir()
//
//				Sys() any
//					underlying data source (can return nil)
//		}
//
//		See the detailed documentation for Type
//		FileInfoPlus in the source file,
//		'fileinfoplus.go'.
//
//	err							error
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
func (fBuffReadWriteMolecule *fileBufferReadWriteMolecule) setFileMgrWriter(
	fBufReadWrite *FileBufferReadWrite,
	fBufReadWriteLabel string,
	writerFileMgr *FileMgr,
	writerFileMgrLabel string,
	openWriteFileReadWrite bool,
	writerBuffSize int,
	truncateExistingWriteFile bool,
	errPrefDto *ePref.ErrPrefixDto) (
	fInfoPlus FileInfoPlus,
	err error) {

	if fBuffReadWriteMolecule.lock == nil {
		fBuffReadWriteMolecule.lock = new(sync.Mutex)
	}

	fBuffReadWriteMolecule.lock.Lock()

	defer fBuffReadWriteMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileBufferReadWriteMolecule." +
		"setFileMgrWriter()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		return fInfoPlus, err
	}

	if len(fBufReadWriteLabel) == 0 {

		fBufReadWriteLabel = "fBufReadWrite"
	}

	if len(writerFileMgrLabel) == 0 {

		writerFileMgrLabel = "writerFileMgr"
	}

	if fBufReadWrite == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is a nil pointer!\n"+
			"%v is invalid.\n",
			ePrefix.String(),
			fBufReadWriteLabel,
			fBufReadWriteLabel)

		return fInfoPlus, err
	}

	var err2 error

	err2 = new(fileMgrHelperAtom).isFileMgrValid(
		writerFileMgr,
		ePrefix.XCpy(writerFileMgrLabel))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter %v is invalid.\n"+
			"%v failed the validity test.\n"+
			"Error=\n%v\n",
			funcName,
			writerFileMgrLabel,
			writerFileMgrLabel,
			err2.Error())

		return fInfoPlus, err
	}

	err2 = new(fileMgrHelper).closeFile(
		writerFileMgr,
		ePrefix.XCpy(writerFileMgrLabel))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"An error occurred while closing the file pointer.\n"+
			"for FileMgr input parameter '%v'.\n"+
			"Error=\n%v\n",
			funcName,
			writerFileMgrLabel,
			err2.Error())

		return fInfoPlus, err
	}

	writerFileMgrLabel += ".absolutePathFileName"

	fInfoPlus,
		err = new(fileBufferReadWriteAtom).
		setPathFileNameWriter(
			fBufReadWrite,
			fBufReadWriteLabel,
			writerFileMgr.absolutePathFileName,
			writerFileMgrLabel,
			openWriteFileReadWrite,
			writerBuffSize,
			truncateExistingWriteFile,
			ePrefix)

	return fInfoPlus, err
}

type fileBufferReadWriteAtom struct {
	lock *sync.Mutex
}

// setIoReader
//
// Receives an object which implements io.Reader
// interface. This object is then used to configure
// the internal io.Reader member variable encapsulated in
// the FileBufferReadWrite instance passed as input
// parameter 'fBufReadWrite'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reconfigure
//	the member variable io.Reader object encapsulated in
//	the instance of FileBufferReadWrite passed as input
//	parameter 'fBufReadWrite'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fBufReadWrite				*FileBufferReadWrite
//
//		A pointer to an instance of FileBufferWriter.
//
//		The internal io.Reader object encapsulated in
//		this instance of FileBufferReadWrite will be
//		deleted and configured using the io.Reader
//		instance passed as input parameter 'reader'.
//
//	fBufReadWriteLabel			string
//
//		The name or label associated with input parameter
//		'fBufReadWrite' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fBufReadWrite" will
//		be automatically applied.
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
//		in addition to calling the local method:
//
//		FileBufferReadWrite.CloseFileBufferReadWrite()
//
//		While the 'read' services provided by
//		FileBufferReadWrite are primarily designed to
//		read data from disk files, this type of 'reader'
//		will in fact read data from any object
//		implementing the io.Reader interface.
//
//	readerLabel					string
//
//		The name or label associated with input parameter
//		'reader' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "reader" will
//		be automatically applied.
//
//	readerBuffSize				int
//
//		This integer value controls the size of the
//		'read' buffer created for the internal io.Reader
//		encapsulated in the instance of
//		FileBufferReadWrite passed as input parameter
//		'fBufReadWrite'.
//
//		'readerBuffSize' should be configured to maximize
//		performance for 'read' operations subject to
//		prevailing memory limitations.
//
//		The minimum reader buffer size is 16-bytes. If
//		'readerBuffSize' is set to a value less than
//		"16", it will be automatically reset to the
//		default buffer size of 4096-bytes.
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
func (fBuffReadWriteAtom *fileBufferReadWriteAtom) setIoReader(
	fBufReadWrite *FileBufferReadWrite,
	fBufReadWriteLabel string,
	reader io.Reader,
	readerLabel string,
	readerBuffSize int,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fBuffReadWriteAtom.lock == nil {
		fBuffReadWriteAtom.lock = new(sync.Mutex)
	}

	fBuffReadWriteAtom.lock.Lock()

	defer fBuffReadWriteAtom.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileBufferReadWriteAtom." +
		"setIoReader()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return err
	}

	if len(fBufReadWriteLabel) == 0 {

		fBufReadWriteLabel = "fBufReadWrite"
	}

	if len(readerLabel) == 0 {

		readerLabel = "reader"
	}

	if fBufReadWrite == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is a nil pointer!\n"+
			"%v is invalid.\n",
			ePrefix.String(),
			fBufReadWriteLabel,
			fBufReadWriteLabel)

		return err
	}

	err = new(fileBufferReadWriteElectron).closeReader(
		fBufReadWrite,
		fBufReadWriteLabel,
		ePrefix.XCpy("Close-Reader"))

	if err != nil {

		return err
	}

	var newBuffReader FileBufferReader
	var err2 error

	err2 =
		new(fileBufferReaderNanobot).setIoReader(
			&newBuffReader,
			"newBuffReader",
			reader,
			readerLabel,
			readerBuffSize,
			ePrefix.XCpy("newBuffReader<-reader"))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"An error occurred while creating the new %v.reader.\n"+
			"Error=\n%v\n",
			funcName,
			fBufReadWriteLabel,
			err2.Error())

		return err
	}

	fBufReadWrite.reader = &newBuffReader

	return err
}

// setIoWriter
//
// Receives an instance of File Manager (FileMgr) which
// will be used to configure an io.Writer object
// encapsulated by the FileBufferReadWrite instance
// passed as input parameter 'fBufReadWrite'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reconfigure
//	the member variable io.Writer object encapsulated in
//	the instance of FileBufferReadWrite passed as input
//	parameter 'fBufReadWrite'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fBufReadWrite				*FileBufferReadWrite
//
//		A pointer to an instance of FileBufferWriter.
//
//		The internal io.Writer object encapsulated in
//		this instance of FileBufferReadWrite will be
//		deleted and reconfigured using the io.Writer
//		object passed as input parameter 'writer'.
//
//	fBufReadWriteLabel			string
//
//		The name or label associated with input parameter
//		'fBufReadWrite' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fBufReadWrite" will
//		be automatically applied.
//
//	writer						io.Writer
//
//		This parameter will accept any object
//		implementing the io.Writer interface.
//
//		This object may be a file pointer of type
//		*os.File. File pointers of this type implement
//		the io.Writer interface.
//
//		A file pointer (*os.File) will facilitate writing
//		output data to destination files residing on an
//		attached storage drive. However, with this
//		configuration, the user is responsible for
//		manually closing the file and performing any
//		other required clean-up operations in addition to
//		calling local method:
//
//		FileBufferReadWrite.CloseFileBufferReadWrite()
//
//		While the 'write' services provided by the
//		FileBufferReadWrite are primarily designed for
//		writing data to disk files, this type of 'writer'
//		will in fact write data to any object
//		implementing the io.Writer interface.
//
//	writerLabel					string
//
//		The name or label associated with input parameter
//		'writer' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "writer" will be
//		automatically applied.
//
//	writerBuffSize				int
//
//		This integer value controls the size of the
//		'write' buffer created for the io.Writer
//		object passed as input parameter 'writer'.
//		This io.Writer object will in turn be configured
//		and encapsulated in the FileBufferWriter
//		instance passed as input parameter 'fBufReadWrite'.
//
//		'writerBuffSize' should be configured to maximize
//		performance for 'write' operations subject to
//		prevailing memory limitations.
//
//		The minimum write buffer size is 1-byte. If
//		'writerBuffSize' is set to a size less than or
//		equal to zero, it will be automatically reset to
//		the default buffer size of 4096-bytes.
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
func (fBuffReadWriteAtom *fileBufferReadWriteAtom) setIoWriter(
	fBufReadWrite *FileBufferReadWrite,
	fBufReadWriteLabel string,
	writer io.Writer,
	writerLabel string,
	writerBuffSize int,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fBuffReadWriteAtom.lock == nil {
		fBuffReadWriteAtom.lock = new(sync.Mutex)
	}

	fBuffReadWriteAtom.lock.Lock()

	defer fBuffReadWriteAtom.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileBufferReadWriteAtom." +
		"setIoWriter()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return err
	}

	if len(fBufReadWriteLabel) == 0 {

		fBufReadWriteLabel = "fBufReadWrite"
	}

	if len(fBufReadWriteLabel) == 0 {

		fBufReadWriteLabel = "fBufReadWrite"
	}

	if len(writerLabel) == 0 {

		writerLabel = "writer"
	}

	if fBufReadWrite == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is a nil pointer!\n"+
			"%v is invalid.\n",
			ePrefix.String(),
			fBufReadWriteLabel,
			fBufReadWriteLabel)

		return err
	}

	err = new(fileBufferReadWriteElectron).flushAndCloseWriter(
		fBufReadWrite,
		fBufReadWriteLabel,
		ePrefix.XCpy("Close-Writer"))

	if err != nil {

		return err
	}

	var newBuffWriter FileBufferWriter
	var err2 error

	err2 =
		new(fileBufferWriterNanobot).setIoWriter(
			&newBuffWriter,
			"newBuffWriter",
			writer,
			writerLabel,
			writerBuffSize,
			ePrefix.XCpy("newBuffWriter<-writer"))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"An error occurred while creating the new %v.writer.\n"+
			"Error=\n%v\n",
			funcName,
			fBufReadWriteLabel,
			err2.Error())

		return err

	}

	fBufReadWrite.writer = &newBuffWriter

	return err
}

// setPathFileNameReader
//
// Receives an input parameter string specifying the path
// and file name identifying the file which will be
// configured as a data source for 'read' operations.
// This file will be configured as an internal io.Reader
// object for the FileBufferReadWrite instance passed as
// input parameter 'fBufReadWrite'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reconfigure
//	the member variable io.Reader object encapsulated in
//	the instance of FileBufferReadWrite passed as input
//	parameter 'fBufReadWrite'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	readerPathFileName			string
//
//		This string contains the path and file name of
//		the file which will be configured as an io.Reader
//		object encapsulated in the FileBufferReadWrite
//		instance passed as input parameter
//		'fBufReadWrite'. As such, the file identified by
//		'readerPathFileName' will be used a data source
//		for 'read' operations.
//
//		If this file does not currently exist on an
//		attached storage drive, an error will be
//		returned.
//
//	readerPathFileNameLabel		string
//
//		The name or label associated with input parameter
//		'readerPathFileName' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "readerPathFileName"
//		will be automatically applied.
//
//	openReadFileReadWrite		bool
//
//		If this parameter is set to 'true', the target
//		'read' file identified from input parameter
//		'readerPathFileName' will be opened for both
//		'read' and 'write' operations.
//
//		If 'openReadFileReadWrite' is set to 'false', the
//		target 'read' file will be opened for 'read-only'
//		operations.
//
//	readerBuffSize					int
//
//		This integer value controls the size of the
//		'read' buffer created for the io.Reader object
//		generated from the file identified by
//		'readerPathFileName'. This io.Reader object is
//		encapsulated in the FileBufferReadWrite instance
//		passed as input	parameter 'fBufReadWrite'.
//
//		'readerBuffSize' should be configured to maximize
//		performance for 'read' operations subject to
//		prevailing memory limitations.
//
//		The minimum read buffer size is 1-byte. If
//		'bufSize' is set to a size less than or equal to
//		zero, it will be automatically reset to the
//		default buffer size of 4096-bytes.
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
//	fileInfoPlus				FileInfoPlus
//
//		This returned instance of Type FileInfoPlus
//		contains data elements describing the file
//		identified by input parameter
//		'readerPathFileName'.
//
//		Type FileInfoPlus conforms to the os.FileInfo
//		interface. This structure will store os.FileInfo
//	 	information plus additional information related
//	 	to a file or directory.
//
//		type os.FileInfo interface {
//
//				Name() string
//					base name of the file
//
//				Size() int64
//					length in bytes for regular files;
//					system-dependent for others
//
//				Mode() FileMode
//					file mode bits
//
//				ModTime() time.Time
//					modification time
//
//				IsDir() bool
//					abbreviation for Mode().IsDir()
//
//				Sys() any
//					underlying data source (can return nil)
//		}
//
//		See the detailed documentation for Type
//		FileInfoPlus in the source file,
//		'fileinfoplus.go'.
//
//	err							error
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
func (fBuffReadWriteAtom *fileBufferReadWriteAtom) setPathFileNameReader(
	fBufReadWrite *FileBufferReadWrite,
	fBufReadWriteLabel string,
	readerPathFileName string,
	readerPathFileNameLabel string,
	openReadFileReadWrite bool,
	readerBuffSize int,
	errPrefDto *ePref.ErrPrefixDto) (
	fInfoPlus FileInfoPlus,
	err error) {

	if fBuffReadWriteAtom.lock == nil {
		fBuffReadWriteAtom.lock = new(sync.Mutex)
	}

	fBuffReadWriteAtom.lock.Lock()

	defer fBuffReadWriteAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileBufferReadWriteAtom." +
		"setPathFileNameReader()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		return fInfoPlus, err
	}

	if len(fBufReadWriteLabel) == 0 {

		fBufReadWriteLabel = "fBufReadWrite"
	}

	if len(readerPathFileNameLabel) == 0 {

		readerPathFileNameLabel = "readerPathFileName"
	}

	if fBufReadWrite == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is a nil pointer!\n"+
			"%v is invalid.\n",
			ePrefix.String(),
			fBufReadWriteLabel,
			fBufReadWriteLabel)

		return fInfoPlus, err
	}

	if len(readerPathFileName) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is invalid!\n"+
			"'%v' is an empty string with a length of zero (0).\n",
			ePrefix.String(),
			readerPathFileNameLabel,
			readerPathFileNameLabel)

		return fInfoPlus, err
	}

	err = new(fileBufferReadWriteElectron).closeReader(
		fBufReadWrite,
		fBufReadWriteLabel,
		ePrefix.XCpy("Close-Reader"))

	if err != nil {

		return fInfoPlus, err
	}

	var newBuffReader FileBufferReader
	var err2 error

	fInfoPlus,
		err2 = new(fileBufferReaderNanobot).
		setPathFileName(
			&newBuffReader,
			"newBuffReader",
			readerPathFileName,
			readerPathFileNameLabel,
			openReadFileReadWrite, // openFileReadWrite
			readerBuffSize,
			ePrefix.XCpy("newBuffReader<-readerPathFileName"))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"An error occurred while creating the new %v.reader.\n"+
			"Error=\n%v\n",
			funcName,
			fBufReadWriteLabel,
			err2.Error())

		return fInfoPlus, err
	}

	fBufReadWrite.reader = &newBuffReader
	fBufReadWrite.readerFilePathName = readerPathFileName

	return fInfoPlus, err
}

// setPathFileNameWriter
//
// Receives an input parameter string specifying the path
// and file name identifying the file which will be
// configured as an output destination for 'write'
// operations. This file will be configured as an
// internal io.Writer object for the FileBufferReadWrite
// instance passed as input parameter 'fBufReadWrite'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reconfigure
//	the member variable io.Writer object encapsulated in
//	the instance of FileBufferReadWrite passed as input
//	parameter 'fBufReadWrite'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fBufReadWrite				*FileBufferReadWrite
//
//		A pointer to an instance of FileBufferWriter.
//
//		The internal io.Writer object encapsulated in
//		this instance of FileBufferReadWrite will be
//		deleted and configured using the file identified
//		by input parameter 'writerPathFileName'.
//
//	fBufReadWriteLabel			string
//
//		The name or label associated with input parameter
//		'fBufReadWrite' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fBufReadWrite" will
//		be automatically applied.
//
//	writerPathFileName			string
//
//		This string contains the path and file name of
//		the target 'write' file which will be used as
//		a data destination for 'write' operations.
//
//		If the target path and file do not currently
//		exist on an attached storage drive, this method
//		will attempt to create them.
//
//	writerPathFileNameLabel		string
//
//		The name or label associated with input parameter
//		'writerPathFileName' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "writerPathFileName"
//		will be automatically applied.
//
//	openWriteFileReadWrite		bool
//
//		If this parameter is set to 'true', the target
//		'write' file identified by input parameter
//		'writerPathFileName' will be opened for 'read'
//		and 'write' operations.
//
//		If 'openWriteFileReadWrite' is set to 'false',
//		the target write file will be opened for
//		'write-only' operations.
//
//	writerBuffSize				int
//
//		This integer value controls the size of the
//		'write' buffer created for the io.Writer
//		object generated from the file identified by
//		input parameter 'writerPathFileName'. This
//		io.Writer object is encapsulated in the
//		FileBufferReadWrite instance passed as input
//		parameter 'fBufReadWrite'.
//
//		'writerBuffSize' should be configured to maximize
//		performance for 'write' operations subject to
//		prevailing memory limitations.
//
//		If 'writerBuffSize' is set to a value less than
//		or equal to zero (0), it will be automatically
//		reset to the default value of 4096-bytes.
//
//	truncateExistingWriteFile	bool
//
//		If this parameter is set to 'true', the target
//		'write' file ('writerPathFileName') will be
//		opened for write operations. If the target write
//		file previously existed, it will be truncated.
//		This means that the file's previous contents will
//		be deleted.
//
//		If this parameter is set to 'false', the target
//		'write' file will be opened for write operations.
//		If the target 'write' file previously existed,
//		the new text written to this file will be appended
//		to the end of the previous file contents.
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
//	fileInfoPlus				FileInfoPlus
//
//		This returned instance of Type FileInfoPlus
//		contains data elements describing the file
//		identified by input parameter
//		'writerPathFileName'.
//
//		Type FileInfoPlus conforms to the os.FileInfo
//		interface. This structure will store os.FileInfo
//	 	information plus additional information related
//	 	to a file or directory.
//
//		type os.FileInfo interface {
//
//				Name() string
//					base name of the file
//
//				Size() int64
//					length in bytes for regular files;
//					system-dependent for others
//
//				Mode() FileMode
//					file mode bits
//
//				ModTime() time.Time
//					modification time
//
//				IsDir() bool
//					abbreviation for Mode().IsDir()
//
//				Sys() any
//					underlying data source (can return nil)
//		}
//
//		See the detailed documentation for Type
//		FileInfoPlus in the source file,
//		'fileinfoplus.go'.
//
//	err							error
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
func (fBuffReadWriteAtom *fileBufferReadWriteAtom) setPathFileNameWriter(
	fBufReadWrite *FileBufferReadWrite,
	fBufReadWriteLabel string,
	writerPathFileName string,
	writerPathFileNameLabel string,
	openWriteFileReadWrite bool,
	writerBuffSize int,
	truncateExistingWriteFile bool,
	errPrefDto *ePref.ErrPrefixDto) (
	fInfoPlus FileInfoPlus,
	err error) {

	if fBuffReadWriteAtom.lock == nil {
		fBuffReadWriteAtom.lock = new(sync.Mutex)
	}

	fBuffReadWriteAtom.lock.Lock()

	defer fBuffReadWriteAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileBufferReadWriteAtom." +
		"setPathFileNameWriter()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		return fInfoPlus, err
	}

	if len(fBufReadWriteLabel) == 0 {

		fBufReadWriteLabel = "fBufReadWrite"
	}

	if len(writerPathFileNameLabel) == 0 {

		writerPathFileNameLabel = "writerPathFileName"
	}

	if fBufReadWrite == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is a nil pointer!\n"+
			"%v is invalid.\n",
			ePrefix.String(),
			fBufReadWriteLabel,
			fBufReadWriteLabel)

		return fInfoPlus, err
	}

	if len(writerPathFileName) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is invalid!\n"+
			"'%v' is an empty string with a length of zero (0).\n",
			ePrefix.String(),
			writerPathFileNameLabel,
			writerPathFileNameLabel)

		return fInfoPlus, err
	}

	err = new(fileBufferReadWriteElectron).flushAndCloseWriter(
		fBufReadWrite,
		fBufReadWriteLabel,
		ePrefix.XCpy("Close-Writer"))

	if err != nil {

		return fInfoPlus, err
	}

	var newBuffWriter FileBufferWriter
	var err2 error

	fInfoPlus,
		err2 = new(fileBufferWriterNanobot).
		setPathFileName(
			&newBuffWriter,
			"newBuffWriter",
			writerPathFileName,
			writerPathFileNameLabel,
			openWriteFileReadWrite, // openFileReadWrite
			writerBuffSize,
			truncateExistingWriteFile,
			ePrefix.XCpy("newBuffWriter<-writerPathFileName"))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"An error occurred while creating the new %v.writer.\n"+
			"Error=\n%v\n",
			funcName,
			fBufReadWriteLabel,
			err2.Error())

		return fInfoPlus, err
	}

	fBufReadWrite.writer = &newBuffWriter
	fBufReadWrite.writerFilePathName = writerPathFileName

	return fInfoPlus, err
}

type fileBufferReadWriteElectron struct {
	lock *sync.Mutex
}

// closeReader
//
// This method is designed to perform clean up tasks
// associated with the io.Reader configured for the
// instance of FileBufferReadWrite passed as input
// parameter 'fBufReadWrite'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will effectively render the instance of
//	FileBufferReadWrite, passed as input parameter
//	'fBufReadWrite', invalid and unusable for any
//	future 'read' operations.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fBufReadWrite				*FileBufferReadWrite
//
//		A pointer to an instance of FileBufferWriter.
//
//		The internal io.Reader object encapsulated
//		in this instance of FileBufferReadWrite will be
//		deleted as part of this 'close' operation.
//
//		Upon completion of this method, 'fBufReadWrite'
//		will be invalid and unusable for future 'read'
//		operations.
//
//	fBufReadWriteLabel			string
//
//		The name or label associated with input parameter
//		'fBufReadWrite' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fBufReadWrite" will
//		be automatically applied.
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
func (fBuffReadWriteElectron *fileBufferReadWriteElectron) closeReader(
	fBufReadWrite *FileBufferReadWrite,
	fBufReadWriteLabel string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fBuffReadWriteElectron.lock == nil {
		fBuffReadWriteElectron.lock = new(sync.Mutex)
	}

	fBuffReadWriteElectron.lock.Lock()

	defer fBuffReadWriteElectron.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileBufferReadWriteElectron." +
		"closeReader()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return err
	}

	if len(fBufReadWriteLabel) == 0 {

		fBufReadWriteLabel = "fBufReadWrite"
	}

	if fBufReadWrite == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is a nil pointer!\n"+
			"%v is invalid.\n",
			ePrefix.String(),
			fBufReadWriteLabel,
			fBufReadWriteLabel)

		return err
	}

	var err2 error

	if fBufReadWrite.reader != nil {

		err2 = new(fileBufferReaderMolecule).close(
			fBufReadWrite.reader,
			fBufReadWriteLabel+".reader",
			ePrefix.XCpy(
				fBufReadWriteLabel+".reader"))

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"An error occurred while closing %v.reader.\n"+
				"Error=\n%v\n",
				funcName,
				fBufReadWriteLabel,
				err2.Error())

			return err
		}

	}

	fBufReadWrite.reader = nil
	fBufReadWrite.readerFilePathName = ""

	return err
}

// flushAndCloseWriter
//
// This method is designed to perform clean up tasks
// associated with the io.Writer configured for the
// instance of FileBufferReadWrite passed as input
// parameter 'fBufReadWrite'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will effectively render the instance of
//	FileBufferReadWrite, passed as input parameter
//	'fBufReadWrite', invalid and unusable for any
//	future 'write' operations. The io.Writer object
//	configured for 'fBufReadWrite' will be deleted.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fBufReadWrite				*FileBufferReadWrite
//
//		A pointer to an instance of FileBufferWriter.
//
//		The internal io.Writer object encapsulated
//		in this instance of FileBufferReadWrite will be
//		deleted as part of this 'close' operation.
//
//		Upon completion of this method, 'fBufReadWrite'
//		will be invalid and unusable for any future
//		'write' operations.
//
//	fBufReadWriteLabel			string
//
//		The name or label associated with input parameter
//		'fBufReadWrite' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fBufReadWrite" will
//		be automatically applied.
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
func (fBuffReadWriteElectron *fileBufferReadWriteElectron) flushAndCloseWriter(
	fBufReadWrite *FileBufferReadWrite,
	fBufReadWriteLabel string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fBuffReadWriteElectron.lock == nil {
		fBuffReadWriteElectron.lock = new(sync.Mutex)
	}

	fBuffReadWriteElectron.lock.Lock()

	defer fBuffReadWriteElectron.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileBufferReadWriteElectron." +
		"flushAndCloseWriter()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return err
	}

	if len(fBufReadWriteLabel) == 0 {

		fBufReadWriteLabel = "fBufReadWrite"
	}

	if fBufReadWrite == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is a nil pointer!\n"+
			"%v is invalid.\n",
			ePrefix.String(),
			fBufReadWriteLabel,
			fBufReadWriteLabel)

		return err
	}

	var err2 error
	var fBufWriterMolecule = new(fileBufferWriterMolecule)
	var errs []error

	if fBufReadWrite.writer != nil {

		err2 = fBufWriterMolecule.
			flush(
				fBufReadWrite.writer,
				fBufReadWriteLabel+".writer",
				ePrefix.XCpy(
					fBufReadWriteLabel+".writer"))

		if err2 != nil {

			errs = append(
				errs,
				fmt.Errorf("%v",
					err2.Error()))
		}

		err2 = fBufWriterMolecule.close(
			fBufReadWrite.writer,
			fBufReadWriteLabel+".writer",
			ePrefix.XCpy(
				fBufReadWriteLabel+".writer"))

		if err2 != nil {

			errs = append(
				errs,
				fmt.Errorf("%v\n"+
					"An error occurred while closing %v.writer.\n"+
					"Error=\n%v\n",
					funcName,
					fBufReadWriteLabel,
					err2.Error()))
		}

		if len(errs) > 0 {

			err = new(StrMech).ConsolidateErrors(errs)
		}
	}

	fBufReadWrite.writer = nil
	fBufReadWrite.writerFilePathName = ""

	return err
}
