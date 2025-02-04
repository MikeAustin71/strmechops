package strmech

import (
	"bufio"
	"errors"
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"strings"
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
//		by configuring both the internal bufio.Reader
//		and	the bufio.Writer simultaneously:
//
//		FileBufferReadWrite.NewIoReadWrite()
//		FileBufferReadWrite.NewFileMgrsReadWrite()
//		FileBufferReadWrite.NewPathFileNamesReadWrite()
//
//		Option-2
//
//		This second option calls for configuring the
//		internal bufio.Reader and bufio.Writer objects
//		individually. This approach allows the user to
//		choose different types of input parameters for
//		configuring the internal bufio.Reader and
//		bufio.Writer by calling 'Setter' methods.
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
//		FileBufferReadWrite.Read()
//		FileBufferReadWrite.ReadWriteBytes()
//		FileBufferReadWrite.Write()
//
//
//	Clean-Up
//
//	When all 'read' and 'write' operations have been
//	completed, the user is responsible for performing
//	Clean-Up operations by calling the following method:
//
//		FileBufferReadWrite.Close()
//
//	NOTE -	This 'Close' method will also 'flush' the
//			'write' buffer before closing the internal
//			bufio.Reader and bufio.Writer objects.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	Pointer receiver FileBufferReader methods
//		implement the following interfaces:
//
//			io.Closer
//			io.ReadCloser
//			io.Reader
//			io.ReadWriteCloser
//			io.ReadWriter
//			io.Seeker
//			io.Writer
//			io.WriteCloser
//
//	(2)	When all 'read' and 'write' operations have been
//		completed, the user is responsible for performing
//		Clean-Up operations by calling the following
//		method:
//
//			FileBufferReadWrite.Close()
type FileBufferReadWrite struct {
	writer             *FileBufferWriter
	reader             *FileBufferReader
	writerFilePathName string
	readerFilePathName string

	lock *sync.Mutex
}

// Close
//
// This method is provided in order to implement the
// io.Closer interface.
//
// FileBufferReadWrite.Close() effectively performs all
// required Clean-Up tasks. As such, this method should
// only be called after all 'read' and 'write' operations
// have been completed and the services of the current
// FileBufferReadWrite instance are no longer required.
//
// After calling this method, FileBufferReadWrite.Close(),
// the current instance of FileBufferReadWrite will be
// invalid and unavailable for further 'read' and/or
// 'write' operations.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	This method implements the io.Closer interface.
//
//	(2)	After completing all 'read' and 'write' operations,
//		calling this method will:
//
//		(a)	'Flush' the write buffer thereby ensuring all
//			data is written from the write buffer to the
//			underlying bufio.Writer object.
//
//		(b) Properly 'Close' the internal bufio.Writer
//			object.
//
//		(c) Properly 'Close' the internal bufio.Reader
//			object.
//
//		(d)	Release internal memory resources.
//
//			Releasing all internal memory resources will
//			synchronize internal flags and prevent
//			multiple calls to 'close' the underlying
//			bufio.Reader and bufio.Writer objects.
//
//			Calling 'close' on the same underlying
//			bufio.Reader or bufio.Writer object multiple
//			times can produce unexpected results.
//
//	(3)	Once this method completes all required Clean-Up
//		tasks, this current instance of
//		FileBufferReadWrite will become unavailable for
//		further 'read' and/or 'write' operations.
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
func (fBufReadWrite *FileBufferReadWrite) Close() error {

	if fBufReadWrite.lock == nil {
		fBufReadWrite.lock = new(sync.Mutex)
	}

	fBufReadWrite.lock.Lock()

	defer fBufReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"FileBufferReadWrite."+
			"Close()",
		"")

	if err != nil {
		return err
	}

	return new(fileBufferReadWriteMicrobot).
		flushCloseRelease(
			fBufReadWrite,
			"fBufReadWrite",
			true, // flushWriteBuffer
			true, // releaseReaderWriterMemResources
			true, // releaseFBuffReadWriteMemResources
			ePrefix.XCpy("Close-Reader&Writer"))
}

// CloseReader
//
// This method is designed to perform Clean-Up tasks
// after completion of all 'read' operations associated
// with the current instance of FileBufferReadWrite.
//
// After calling this method, the Clean-Up tasks
// performed will effectively render the internal
// bufio.Reader object, encapsulated by the current
// FileBufferReadWrite instance, invalid and unusable
// for any future 'read' operations.
//
// It is unlikely that the user will ever need to call
// this method. Typically, Clean-Up tasks are performed
// jointly on the internal bufio.Reader and bufio.Writer
// objects encapsulated in the current FileBufferReadWrite
// instance. These Clean-Up tasks should be performed
// after all 'read' and 'write' operations have been
// completed by calling the local method:
//
//	FileBufferReadWrite.Close()
//
// However, in the event of unforeseen use cases, this
// method is provided to exclusively close or Clean-Up
// the bufio.Reader object.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will:
//
//	(1) Properly 'Close' the 'read' file or internal
//		bufio.Reader object encapsulated by the current
//		instance of FileBufferReadWrite.
//
//	(2) Release the internal memory resources associated
//		with the internal bufio.Reader object
//		encapsulated by the current instance of
//		FileBufferReadWrite.
//
//		Releasing internal memory resources synchronizes
//		internal flags and prevents multiple calls to the
//		'close' method. Calling the 'close' method more
//		than once on the internal bufio.Reader object may
//		produce unexpected results.
//
//	(3) Effectively render the internal bufio.Reader
//		object, encapsulated by the current instance of
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
		readerCloseRelease(
			fBufReadWrite,
			"fBufReadWrite",
			true, // releaseReaderMemResources
			true, // releaseFBuffReaderLocalMemRes
			ePrefix)

	return err
}

// CloseWriter
//
// This method is designed to perform Clean-Up tasks
// after completion of all 'write' operations associated
// with the current instance of FileBufferReadWrite.
//
// After calling this method, the Clean-Up tasks
// performed will effectively render the internal
// bufio.Writer object, encapsulated by the current
// FileBufferReadWrite instance, invalid and unusable
// for any future 'write' operations.
//
// It is unlikely that the user will ever need to call
// this method. Typically, Clean-Up tasks are performed
// jointly on the internal io.Reader and io.Writer
// objects encapsulated in the current FileBufferReadWrite
// instance. These Clean-Up tasks should be performed
// after all 'read' and 'write' operations have been
// completed by calling the local method:
//
//	FileBufferReadWrite.Close()
//
// However, in the event of unforeseen use cases, this
// method is provided to exclusively close or Clean-Up
// the internal bufio.Writer object.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will:
//
//	(1)	'Flush' the write buffer thereby ensuring all
//		data is written from the write buffer to the
//		underlying io.Writer object. This 'Flush'
//		procedure is dependent on the value of input
//		parameter 'flushWriteBuffer'.
//
//	(2) Properly 'Close' the 'write' file or internal
//		bufio.Writer object encapsulated by the current
//		instance of FileBufferReadWrite.
//
//	(3) Release the internal memory resources associated
//		with the internal bufio.Writer object
//		encapsulated by the current instance of
//		FileBufferReadWrite.
//
//		Releasing internal memory resources synchronizes
//		internal flags and prevents multiple calls to the
//		'close' method. Calling the 'close' method more
//		than once on the internal bufio.Writer object may
//		produce unexpected results.
//
//	(4) Effectively render the internal io.Writer object,
//		encapsulated by the current instance of
//		FileBufferReadWrite, invalid and unusable for any
//		future 'write' operations.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	flushWriteBuffer			bool
//
//		If this parameter is set to 'true', this method
//		will flush the contents of the internal
//		bufio.Writer 'write' buffer before closing the
//		same bufio.Writer object. This means that 'write'
//		buffer contents are guaranteed to be written to
//		the internal bufio.Writer object encapsulated by
//		FileBufferReadWrite input parameter 'fBufReadWrite',
//		before the bufio.Writer object is closed.
//
//		If 'flushWriteBuffer', the internal bufio.Writer
//		object will be closed without flushing the 'write'
//		buffer. In this event, the contents of the 'write'
//		buffer will be lost.
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
	flushWriteBuffer bool,
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
		writerFlushCloseRelease(
			fBufReadWrite,
			"fBufReadWrite",
			flushWriteBuffer, // flushWriteBuffer
			true,             // releaseMemoryResources
			true,             // releaseFBuffWriterLocalMemRes
			ePrefix)

	return err
}

// Empty
//
// This method deletes all internal member variables and
// releases all the internal memory resources for the
// current instance of FileBufferReadWrite.
//
// Specifically the following internal member variables
// are set to nil or their initial zero values:
//
//	FileBufferReadWrite.writer = nil
//	FileBufferReadWrite.reader = nil
//	FileBufferReadWrite.writerFilePathName = ""
//	FileBufferReadWrite.readerFilePathName = ""
//
// After calling this method, the current instance of
// FileBufferReadWrite will become invalid and
// unavailable for future read/write operations.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	This method does NOT perform the 'flush' or 'close'
//	procedures. To perform the 'flush' and 'close'
//	procedures while simultaneously releasing all
//	internal memory resources, call local method:
//
//			FileBufferReadWrite.Close()
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
//	--- NONE ---
func (fBufReadWrite *FileBufferReadWrite) Empty() {

	if fBufReadWrite.lock == nil {
		fBufReadWrite.lock = new(sync.Mutex)
	}

	fBufReadWrite.lock.Lock()

	new(fileBufferReadWriteElectron).
		empty(fBufReadWrite)

	fBufReadWrite.lock.Unlock()

	fBufReadWrite.lock = nil

	return
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

	funcName := "FileBufferReadWrite." +
		"FlushWriteBuffer()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {
		return err
	}

	if fBufReadWrite.writer == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: The current instance of FileBufferReadWrite\n"+
			"is invalid! The internal io.Writer object was never\n"+
			"initialized. Call one of the 'New' methods or 'Setter'\n"+
			"methods to create a valid instance of FileBufferReadWrite.\n",
			ePrefix.String())

		return err
	}

	return new(fileBufferReadWriteMicrobot).
		flushWriteBuffer(
			fBufReadWrite,
			"fBufReadWrite",
			ePrefix)
}

// GetReadFile
//
// Returns a string containing the path and file name for
// the 'Read' file configured for the current instance of
// FileBufferReadWrite.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	NONE
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	string
//
//		If the current instance of FileBufferReadWrite has been
//		properly initialized, this returned string will contain
//		the path and file name of the 'Read' file.
func (fBufReadWrite *FileBufferReadWrite) GetReadFile() string {

	if fBufReadWrite.lock == nil {
		fBufReadWrite.lock = new(sync.Mutex)
	}

	fBufReadWrite.lock.Lock()

	defer fBufReadWrite.lock.Unlock()

	return fBufReadWrite.readerFilePathName
}

// GetWriteFile
//
// Returns a string containing the path and file name for
// the 'Write' file configured for the current instance of
// FileBufferReadWrite.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	NONE
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	string
//
//		If the current instance of FileBufferReadWrite has been
//		properly initialized, this returned string will contain
//		the path and file name of the 'Write' file.
func (fBufReadWrite *FileBufferReadWrite) GetWriteFile() string {

	if fBufReadWrite.lock == nil {
		fBufReadWrite.lock = new(sync.Mutex)
	}

	fBufReadWrite.lock.Lock()

	defer fBufReadWrite.lock.Unlock()

	return fBufReadWrite.writerFilePathName

}

// IsValidInstanceError
//
// Analyzes the current FileBufferReadWrite instance to
// determine if is invalid.
//
// If the current FileBufferReadWrite instance is found
// to be invalid, an error is returned.
//
// If the current FileBufferReadWrite instance is valid,
// this method returns 'nil'
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
//		If any of the internal member data variables
//		contained in the current instance of
//		FileBufferReadWrite are found to be invalid,
//		this method will return an error configured
//		with an appropriate message identifying the
//		invalid	member data variable.
//
//		If all internal member data variables evaluate
//		as valid, this returned error value will be set
//		to 'nil'.
//
//		If errors are encountered during processing or if
//		any FileBufferReadWrite internal member data
//	 	values are found to be invalid, the returned error
//	 	will encapsulate an appropriate error message.
//	 	This returned error message will incorporate the
//	 	method chain and text passed by input parameter,
//	 	'errorPrefix'. The 'errorPrefix' text will be
//	 	prefixed to the beginning of the error message.
func (fBufReadWrite *FileBufferReadWrite) IsValidInstanceError(
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
			"IsValidInstanceError()",
		"")

	if err != nil {

		return err
	}

	err = new(fileBufferReadWriteElectron).
		isFileBufferReadWriteValid(
			fBufReadWrite,
			"current",
			ePrefix.XCpy("fBufReadWrite"))

	return err
}

// New
//
// This method returns a pointer to an empty or 'blank'
// instance of FileBufferReadWrite. All the member
// variables in this returned instance are initialized to
// their zero or initial values. This means the returned
// instance is invalid and unusable for standard 'read'
// and 'write' operations.
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
//	*FileBufferReadWrite
//
//		This method returns a pointer to an empty
//		instance of FileBufferReadWrite. After receiving
//		this instance, users must call 'Setter' methods
//		to complete the 'reader' and 'writer'
//		configuration process.
func (fBufReadWrite *FileBufferReadWrite) New() *FileBufferReadWrite {

	if fBufReadWrite.lock == nil {
		fBufReadWrite.lock = new(sync.Mutex)
	}

	fBufReadWrite.lock.Lock()

	defer fBufReadWrite.lock.Unlock()

	return new(FileBufferReadWrite)
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
//		performing any other required Clean-Up operations
//		in addition to calling local method:
//
//		FileBufferReadWrite.Close()
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
//		performing any other required Clean-Up operations
//		in addition to calling local method:
//
//		FileBufferReadWrite.Close()
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
//	*FileBufferReadWrite
//
//		If this method completes successfully, it will
//		return a pointer to a fully configured instance
//		of FileBufferReadWrite.
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
	*FileBufferReadWrite,
	error) {

	if fBufReadWrite.lock == nil {
		fBufReadWrite.lock = new(sync.Mutex)
	}

	fBufReadWrite.lock.Lock()

	defer fBufReadWrite.lock.Unlock()

	var newFBuffReadWrite = new(FileBufferReadWrite)

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
			newFBuffReadWrite,
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

// NewFileMgrsReadWrite
//
// Creates and returns a pointer to a new, fully
// configured instance of FileBufferReadWrite.
//
// The internal bufio.Reader and bufio.Writer member
// variables for this new instance of FileBufferReadWrite
// are generated from input parameters specifying
// 'reader' and 'writer' files extracted from input
// parameters of type File Manager (FileMgr).
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
//	newFBuffReadWrite			*FileBufferReadWrite
//
//		If this method completes successfully, it will
//		return a pointer to a fully configured instance
//		of FileBufferReadWrite.
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
func (fBufReadWrite *FileBufferReadWrite) NewFileMgrsReadWrite(
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
	newFBuffReadWrite *FileBufferReadWrite,
	err error) {

	if fBufReadWrite.lock == nil {
		fBufReadWrite.lock = new(sync.Mutex)
	}

	fBufReadWrite.lock.Lock()

	defer fBufReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "FileBufferReadWrite." +
		"NewFileMgrsReadWrite()"

	newFBuffReadWrite = new(FileBufferReadWrite)

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
			newFBuffReadWrite,
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

// NewPathFileNamesReadWrite
//
// Creates and returns a pointer to a new, fully
// configured instance of FileBufferReadWrite.
//
// The internal bufio.Reader and bufio.Writer member
// variables for this new instance of FileBufferReadWrite
// are generated from input parameters specifying
// 'reader' and 'writer' files extracted from path and
// file name strings.
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
//	newFBuffReadWrite			*FileBufferReadWrite
//
//		If this method completes successfully, it will
//		return a pointer to a fully configured instance
//		of FileBufferReadWrite.
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
func (fBufReadWrite *FileBufferReadWrite) NewPathFileNamesReadWrite(
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
	newFBuffReadWrite *FileBufferReadWrite,
	err error) {

	if fBufReadWrite.lock == nil {
		fBufReadWrite.lock = new(sync.Mutex)
	}

	fBufReadWrite.lock.Lock()

	defer fBufReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	newFBuffReadWrite = new(FileBufferReadWrite)

	funcName := "FileBufferReadWrite." +
		"NewPathFileNamesReadWrite()"

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
			newFBuffReadWrite,
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

// PeekReader
//
// This method returns the next n bytes without advancing
// the internal bufio.Reader. The bytes stop being valid
// at the next read call.
//
// If Peek returns fewer bytes than 'nextBytes' value, it
// also returns an error explaining why the read is
// short.
//
// If 'nextBytes' value is larger than the buffer size
// for the current read buffer, the error 'ErrBufferFull'
// will be returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	nextBytes					int
//
//		The 'Peek' operation performed by this method
//		will return the number of bytes specified by
//		'nextBytes' without advancing the reader.
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
func (fBufReadWrite *FileBufferReadWrite) PeekReader(
	nextBytes int,
	errorPrefix interface{}) (
	peekBytes []byte,
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
			"PeekReader()",
		"")

	if err != nil {

		return peekBytes, err
	}

	if fBufReadWrite.reader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: The current FileBufferReadWrite instance is INVALID!\n"+
			"The internal bufio.reader (FileBufferReader) is a nil pointer.\n"+
			"This FileBufferReadWrite instance has NOT been properly initialized.\n",
			ePrefix.String())

		return peekBytes, err
	}

	return fBufReadWrite.reader.Peek(
		nextBytes,
		ePrefix.XCpy("fBufReadWrite.reader"))
}

// Read
//
// Reads data from the internal bufio.Reader data source
// encapsulated in the current instance of FileBufferReadWrite.
//
// This method reads data into the input parameter byte
// array, 'bytesRead'. It returns the number of bytes
// read into the byte array as return parameter,
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
// See the bufio.Reader docs and 'Reference' section
// below.
//
// If the internal bufio.Reader object for the current
// instance of FileBufferReadWrite was improperly
// initialized, an error will be returned. To properly
// initialized an instance of FileBufferReadWrite, the
// user must call one or more of the 'New' and/or
// 'Setter' methods.
//
// Once all 'read' and 'write' operations have been
// completed, the user is responsible for performing
// Clean-Up operations by calling the local method:
//
//	FileBufferReadWrite.Close()
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	This method implements the io.Reader interface.
//
//	(2)	Keep calling this method until all the bytes have
//		been read from the io.Reader data source
//		configured for the current instance of
//		FileIoReader and the returned error is set to
//		'io.EOF'.
//
//	(3)	Callers should always process the returned number
//		of bytes read ('numBytesRead') before considering
//		the returned error parameter 'err'. Doing so
//		correctly handles I/O errors that occur after
//		reading some number of bytes. Also, this
//		technique allows both possible EOF behaviors to
//		be correctly processed (See the io.Reader docs
//		and 'Reference' section below).
//
//	(4)	When all 'read' operations have been completed,
//		perform the required Clean-Up operations by
//		calling local method:
//
//			FileIoReader.Close()
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://pkg.go.dev/bufio
//	https://pkg.go.dev/bufio#Reader
//	https://pkg.go.dev/io#Reader
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

// ReadAllTextLines
//
// Reads text lines from the internal bufio.Reader object
// encapsulated in the current instance of
// FileBufferReadWrite.
//
// Multiple custom end of line delimiters may be utilized
// to determine the end of each line of text read from
// the internal bufio.Reader object. End of line delimiters
// are specified by input parameter
// 'endOfLineDelimiters', an instance of StringArrayDto.
// 'endOfLineDelimiters' contains an array of strings any
// one of which may be used to identify, delimit and
// separate individual lines of text read from the internal
// bufio.Reader object configured for the current
// instance of FileBufferReadWrite.
//
// The extracted lines of text will be added to the
// StringArrayDto instance passed as input parameter
// 'outputLinesArray'.
//
// The returned individual lines of text will NOT
// include the end of line delimiters. End of line
// delimiters will therefore be stripped and deleted
// from the end of each configured text line.
//
// It naturally follows that this method will read the
// entire contents of the FileBufferReadWrite internal
// bufio.Reader object into memory when writing said
// contents to the StringArrayDto instance
// 'outputLinesArray'. Depending on the size of the
// target 'read' file, local memory constraints should
// be considered.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	This method is designed to read the entire
//		contents of the internal bufio.Reader object,
//		encapsulated by the current instance of
//		FileBufferReadWrite, into memory.
//
//		BE CAREFUL when reading large files!
//
//		Depending on the memory resources available to
//		your computer, you may run out of memory when
//		reading large files and writing their contents
//		to the output instance of StringArrayDto,
//		'outputLinesArray'.
//
//	(2)	This method will NOT automatically close the
//		internal bufio.Reader object upon completion.
//
//		The user is responsible for performing required
//		Clean-Up tasks by calling one of the following
//		local methods:
//
//			FileBufferReadWrite.Close()
//						OR
//			FileBufferReadWrite.CloseReader()
//
//	(3)	If the current instance of FileBufferReadWrite has
//		NOT been properly initialized, an error will be
//		returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	maxNumOfLines				int
//
//		Specifies the maximum number of text lines which
//		will be read from the internal bufio.Reader
//		encapsulated by the current instance of
//		FileBufferReadWrite.
//
//		If 'maxNumOfLines' is set to a value less than
//		zero (0) (Example: minus-one (-1) ),
//		'maxNumOfLines' will be automatically reset to
//		math.MaxInt(). This means all text lines existing
//		in the internal bufio.Reader object will be read
//		and processed. Reading all the text lines in a
//		file 'may' have memory implications depending on
//		the size of the file and the memory resources
//		available to your computer.
//
//		If 'maxNumOfLines' is set to a value of zero
//		('0'), no text lines will be read from
//		the internal bufio.Reader, and no error will be
//		returned.
//
//	endOfLineDelimiters				*StringArrayDto
//
//		A pointer to an instance of StringArrayDto.
//		'endOfLineDelimiters' encapsulates a string
//		array which contains the end-of-line delimiters
//		that will be used to identify and separate
//		individual lines of text.
//
//		Users have the flexibility to specify multiple
//		end-of-line delimiters for used in parsing text
//		lines extracted from file identified by
//		'pathFileName'.
//
//	outputLinesArray 			*StringArrayDto
//
//		A pointer to an instance of StringArrayDto.
//		Lines of text read from the internal bufio.Reader
//		object configured for the current instance of
//		FileBufferReadWrite will be stored as individual
//		strings in the string array encapsulated by
//		'outputLinesArray'.
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
//	numOfLinesRead				int
//
//		This integer value contains the number of text
//		lines read from the internal bufio.Reader object
//		encapsulated by the current instance of
//		FileBufferReadWrite. This number also specifies
//		the number of string array elements stored in
//		'outputLinesArray'.
//
//		When displayed in editors, the end-of-file
//		character is displayed on a separate line.
//		The returned 'numOfLinesRead' value does
//		not include this empty line containing an
//		end-of-file character. Therefore, the
//		returned 'numOfLinesRead' value will always
//		be one less than the number of lines shown
//		in a text editor.
//
//	numBytesRead				int64
//
//		If this method completes successfully, this
//		integer value will equal the number of bytes
//		read from the internal bufio.Reader object
//		encapsulated by the current instance of
//		FileBufferReadWrite.
//
//		Remember that the number of bytes read
//		includes the end-of-line delimiters which
//		were stripped off and deleted before the
//		text lines were stored in 'outputLinesArray'.
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
func (fBufReadWrite *FileBufferReadWrite) ReadAllTextLines(
	maxNumOfLines int,
	endOfLineDelimiters *StringArrayDto,
	outputLinesArray *StringArrayDto,
	errorPrefix interface{}) (
	numOfLinesRead int,
	numOfBytesRead int64,
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
			"ReadAllTextLines()",
		"")

	if err != nil {

		return numOfLinesRead, numOfBytesRead, err
	}

	if fBufReadWrite.reader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: The current FileBufferReadWrite instance is INVALID!\n"+
			"The internal bufio.reader (FileBufferReader) is a nil pointer.\n"+
			"This FileBufferReadWrite instance has NOT been properly initialized.\n",
			ePrefix.String())

		return numOfLinesRead, numOfBytesRead, err
	}

	return fBufReadWrite.reader.
		ReadAllTextLines(
			maxNumOfLines,
			endOfLineDelimiters,
			outputLinesArray,
			false,
			ePrefix.XCpy("fBufReadWrite.reader"))
}

// ReadAllToStrBuilder
//
// Reads the entire contents of the internal bufio.Reader
// for the current instance of FileBufferReadWrite as
// a string. This string is then stored and returned
// through an instance of strings.Builder passed as input
// parameter 'strBuilder'.
//
// If a processing error is encountered, an appropriate
// error with an error message will be returned. When
// the end-of-file is encountered during the 'read'
// process, the returned error object will be set to
// 'nil' and no error will be returned.
//
// It naturally follows that this method will read the
// entire contents of the FileBufferReadWrite internal
// bufio.Reader object into memory when writing said
// contents to the strings.Builder instance 'strBuilder'.
// Depending on the size of the target 'read' data
// source, local memory constraints should be considered.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	This method is designed to read the entire
//		contents of the internal bufio.Reader object,
//		encapsulated by the current instance of
//		FileBufferReadWrite, into memory.
//
//		BE CAREFUL when reading large files!
//
//		Depending on the memory resources available to
//		your computer, you may run out of memory when
//		reading large files and writing their contents
//		to the strings.Builder input parameter,
//		'strBuilder'.
//
//	(2)	This method will NOT automatically close the
//		internal bufio.Reader object upon completion.
//
//		The user is responsible for performing required
//		Clean-Up tasks by calling one of the following
//		local methods:
//
//			FileBufferReadWrite.Close()
//						OR
//			FileBufferReadWrite.CloseReader()
//
//	(3)	If the current instance of FileBufferReadWrite has
//		NOT been properly initialized, an error will be
//		returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	strBuilder					*strings.Builder
//
//		A pointer to an instance of strings.Builder. The
//		entire contents of the internal bufio.Reader for
//		the current instance of FileBufferReadWrite will be
//		read and stored in 'strBuilder'.
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
//	numBytesRead				int64
//
//		If this method completes successfully, this
//		integer value will equal the number of bytes
//		read from the internal bufio.Reader object
//		encapsulated by the current instance of
//		FileBufferReadWrite.
//
//		This returned value will also be equal to the
//		number of bytes added to the strings.Builder
//		instance, 'strBuilder'.
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
//		An error will only be returned if a processing
//		or system error was encountered. When the
//		end-of-file is encountered during the 'read'
//		process, the returned error object will be set
//		to 'nil' and no error will be returned.
func (fBufReadWrite *FileBufferReadWrite) ReadAllToStrBuilder(
	strBuilder *strings.Builder,
	errorPrefix interface{}) (
	numOfBytesRead int64,
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
			"ReadAllTextLines()",
		"")

	if err != nil {

		return numOfBytesRead, err
	}

	if fBufReadWrite.reader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: The current FileBufferReadWrite instance is INVALID!\n"+
			"The internal bufio.reader (FileBufferReader) is a nil pointer.\n"+
			"This FileBufferReadWrite instance has NOT been properly initialized.\n",
			ePrefix.String())

		return numOfBytesRead, err
	}

	return fBufReadWrite.reader.
		ReadAllToStrBuilder(
			strBuilder,
			false,
			ePrefix.XCpy("fBufReadWrite.reader"))
}

// ReadAllToString
//
// Reads the entire contents of the internal bufio.Reader
// for the current instance of FileBufferReadWrite and
// returns these contents as a single string
// ('contentsStr').
//
// If a processing error is encountered, an appropriate
// error with an error message will be returned. When
// the end-of-file is encountered during the 'read'
// process, the returned error object will be set to
// 'nil' and no error will be returned.
//
// It naturally follows that this method will read the
// entire contents of the FileBufferReadWrite internal
// bufio.Reader object into memory when writing said
// contents to the returned string parameter
// 'contentsStr'. Depending on the size of the target
// 'read' data source, local memory constraints should be
// considered.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	This method is designed to read the entire
//		contents of the internal bufio.Reader object,
//		encapsulated by the current instance of
//		FileBufferReadWrite, into memory.
//
//		BE CAREFUL when reading large files!
//
//		Depending on the memory resources available to
//		your computer, you may run out of memory when
//		reading large files and returning their contents
//		as a single string ('contentsStr').
//
//	(2)	This method will NOT automatically close the
//		internal bufio.Reader object upon completion.
//
//		The user is responsible for performing required
//		Clean-Up tasks by calling one of the following
//		local methods:
//
//			FileBufferReadWrite.Close()
//						OR
//			FileBufferReadWrite.CloseReader()
//
//	(3)	If the current instance of FileBufferReadWrite has
//		NOT been properly initialized, an error will be
//		returned.
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
//	numBytesRead				int64
//
//		If this method completes successfully, this
//		integer value will equal the number of bytes
//		read from the internal bufio.Reader object
//		encapsulated by the current instance of
//		FileBufferReadWrite.
//
//		This integer value should also equal the
//		string length of the returned string,
//		'contentsStr'.
//
//	contentsStr					string
//
//		If this method completes successfully, the entire
//		contents if the internal bufio.Reader object
//		encapsulated by the current instance of
//		FileBufferReadWrite will be returned in this string.
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
//		An error will only be returned if a processing
//		or system error was encountered. When the
//		end-of-file is encountered during the 'read'
//		process, the returned error object will be set
//		to 'nil' and no error will be returned.
func (fBufReadWrite *FileBufferReadWrite) ReadAllToString(
	errorPrefix interface{}) (
	numOfBytesRead int64,
	contentsStr string,
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
			"ReadAllToString()",
		"")

	if err != nil {

		return numOfBytesRead, contentsStr, err
	}

	if fBufReadWrite.reader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: The current FileBufferReadWrite instance is INVALID!\n"+
			"The internal bufio.reader (FileBufferReader) is a nil pointer.\n"+
			"This FileBufferReadWrite instance has NOT been properly initialized.\n",
			ePrefix.String())

		return numOfBytesRead, contentsStr, err
	}

	return fBufReadWrite.reader.
		ReadAllToString(
			false,
			ePrefix.XCpy("fBufReadWrite.reader"))
}

// ReadWriteBytes
//
// This method reads all data from the 'reader' data
// source and writes all that data to the 'writer'
// output destination.
//
// If the total number of bytes read does NOT equal the
// total number of bytes written, an error will be
// returned.
//
// The 'read' and 'write' operations use the io.Reader
// and io.Writer objects created when the current
// instance of FileBufferReadWrite was first initialized.
//
// If input parameter 'autoFlushAndCloseOnExit' is set to
// 'true', this method will automatically perform all
// required Clean-Up tasks upon completion. Clean-Up
// tasks involve flushing the io.Writer object, closing
// the io.Reader and io.Writer objects and then deleting
// io.Reader and io.Writer structure values internal to
// the current FileBufferReadWrite instance. When these
// Clean-Up tasks are completed, the current
// FileBufferReadWrite instance will be invalid and
// unusable for future 'read' and/or 'write' operations.
//
// If input parameter 'autoFlushAndCloseOnExit' is set to
// 'false', this method will automatically flush the
// 'write' buffer. This means that all data remaining in
// the 'write' buffer will be written to the underlying
// io.Writer output destination. Most importantly, the
// user is then responsible for performing the 'Close'
// operation by calling the local method:
//
//	FileBufferReadWrite.Close()
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	If input parameter 'autoFlushAndCloseOnExit' is set
//	to 'false', the user is responsible for calling local
//	method FileBufferReadWrite.Close() in order to
//	perform the required Clean-Up operations on the
//	current instance of FileBufferReadWrite.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	autoFlushAndCloseOnExit		bool
//
//		When this parameter is set to 'true', this
//		method will automatically perform the following
//		Clean-Up tasks upon exit:
//
//		(1)	The write buffer will be flushed thereby
//			ensuring that all remaining data in the
//			'write' buffer will be written to the
//			underlying io.Writer object.
//
//		(2)	The io.Reader and io.Writer objects will be
//			properly closed.
//
//		(3) After performing these Clean-Up tasks, the
//			current instance of FileBufferReadWrite will
//			invalid and unusable for future 'read' and/or
//			'write' operations.
//
//		If input parameter 'autoFlushAndCloseOnExit' is
//		set to 'false', this method will automatically
//		flush the 'write' buffer. This means that all
//		data remaining in the 'write' buffer will be
//		written to the underlying io.Writer output
//		destination. Most importantly, the user is
//		then responsible for performing the 'Close'
//		operation by calling the local method:
//
//			FileBufferReadWrite.Close()
//
//	maxNumOfBytes				int64
//
//		This parameter can be used to limit the number of
//		bytes read from the source internal bufio.Reader
//		object and written to the destination
//		bufio.Writer object.
//
//		If this parameter is set to a value less than
//		one (1) (Examples: zero (0) or minus-one (-1) ),
//		'maxNumOfBytes' will be ignored, and all bytes
//		residing in the bufio.Reader object will be read
//		and then written to the bufio.Writer object.
//
//		When 'maxNumOfLines' is set to a value greater
//		than zero (0), it effectively limits the maximum
//		number of bytes which will be read from the
//		bufio.Reader object and written to the
//		bufio.Writer object.
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
func (fBufReadWrite *FileBufferReadWrite) ReadWriteBytes(
	autoFlushAndCloseOnExit bool,
	maxNumOfBytes int64,
	errorPrefix interface{}) (
	totalBytesRead int64,
	totalBytesWritten int64,
	err error) {

	if fBufReadWrite.lock == nil {
		fBufReadWrite.lock = new(sync.Mutex)
	}

	fBufReadWrite.lock.Lock()

	defer fBufReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "FileBufferReadWrite." +
		"ReadWriteBytes()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {

		return totalBytesRead, totalBytesWritten, err
	}

	err = new(fileBufferReadWriteElectron).
		isFileBufferReadWriteValid(
			fBufReadWrite,
			"current",
			ePrefix.XCpy("fBufReadWrite"))

	if err != nil {

		return totalBytesRead, totalBytesWritten, err
	}

	var numOfBytesRead, numOfBytesWritten,
		cycleCount, numEmptyReads int
	var readErr, err2 error
	var fBufReadWriteMicrobot = new(fileBufferReadWriteMicrobot)

	byteArray := make([]byte,
		fBufReadWrite.reader.bufioReader.Size())

	for {

		if numEmptyReads > 3 {

			break
		}

		if maxNumOfBytes > 0 &&
			totalBytesRead >= maxNumOfBytes {

		}
		cycleCount++

		numOfBytesRead,
			readErr =
			fBufReadWrite.reader.Read(byteArray)

		if readErr != nil &&
			readErr != io.EOF {

			err = fmt.Errorf("%v\n"+
				"Error Reading Target Read File!\n"+
				"Cycle Count= %v\n"+
				"Error= \n%v\n",
				funcName,
				cycleCount,
				readErr.Error())

			err2 = fBufReadWriteMicrobot.
				flushWriteBuffer(
					fBufReadWrite,
					"fBufReadWrite flush#1",
					ePrefix)

			if err2 != nil {

				err = errors.Join(err, err2)

			}

			return totalBytesRead, totalBytesWritten, err
		}

		if numOfBytesRead > 0 {

			numEmptyReads = 0

			totalBytesRead += int64(numOfBytesRead)

			numOfBytesWritten,
				err2 = fBufReadWrite.writer.Write(
				byteArray[0:numOfBytesRead])

			if err2 != nil {

				err = fmt.Errorf("%v\n"+
					"Error Writing Bytes To File!\n"+
					"Write Error=\n%v\n",
					funcName,
					err2.Error())

				err2 = fBufReadWriteMicrobot.
					flushWriteBuffer(
						fBufReadWrite,
						"fBufReadWrite flush#2",
						ePrefix)

				if err2 != nil {

					err = errors.Join(err, err2)
				}

				return totalBytesRead, totalBytesWritten, err
			}

			if numOfBytesRead != numOfBytesWritten {

				err = fmt.Errorf("%v\n"+
					"Error Writing Bytes To File!\n"+
					"numOfBytesRead != numOfBytesWritten\n"+
					"numOfBytesRead = %v\n"+
					"numOfBytesWritten = %v\n",
					funcName,
					numOfBytesRead,
					numOfBytesWritten)

				err2 = fBufReadWriteMicrobot.
					flushWriteBuffer(
						fBufReadWrite,
						"fBufReadWrite flush#3",
						ePrefix)

				if err2 != nil {

					err = errors.Join(err, err2)
				}

				return totalBytesRead, totalBytesWritten, err
			}

			totalBytesWritten += int64(numOfBytesWritten)

		} else {

			numEmptyReads++

		}

		if readErr == io.EOF {
			break
		}

		clear(byteArray)
	}

	if autoFlushAndCloseOnExit == true {

		err = fBufReadWriteMicrobot.
			flushCloseRelease(
				fBufReadWrite,
				"fBufReadWrite",
				true, // flushWriteBuffer
				true, // releaseReaderWriterMemResources
				true, // releaseFBuffReadWriteMemResources
				ePrefix.XCpy("FlushClose-Readers&Writers"))

	} else {

		err = fBufReadWriteMicrobot.
			flushWriteBuffer(
				fBufReadWrite,
				"fBufReadWrite final-flush",
				ePrefix)

	}

	return totalBytesRead, totalBytesWritten, err
}

// ReadWriteTextLines
//
// This method reads all available data from the internal
// bufio.Reader object previously configured for the
// current instance of FileBufferReadWrite. It then parses
// this data into lines of text based on the end-of-line
// delimiter characters passed as input parameter
// 'endOfLineDelimiters'. These end-of-line delimiters
// are stripped off the ends of all text lines processed.
// New line termination or end-of-line characters
// specified by user input parameters will then be
// appended to the text lines before they are written to
// the output destination bufio.Writer object configured
// for the current instance of FileBufferReadWrite.
//
// When writing final text lines to the internal
// bufio.Writer object, the line termination or
// end-of-line characters appended to each text line
// written will be specified by input parameter
// 'writeEndOfLineChars'.
//
// If input parameter 'autoFlushAndCloseOnExit' is set to
// 'true', this method will automatically perform all
// required Clean-Up tasks upon completion. Clean-Up
// tasks involve flushing the bufio.Writer object, closing
// the bufio.Reader and bufio.Writer objects and then
// deleting bufio.Reader and bufio.Writer structure
// values internal to the current FileBufferReadWrite
// instance. When these Clean-Up tasks are completed, the
// current FileBufferReadWrite instance will be invalid
// and unavailable for future 'read' and/or 'write'
// operations.
//
// If input parameter 'autoFlushAndCloseOnExit' is set to
// 'false', this method will automatically flush the
// 'write' buffer. This means that all data remaining in
// the 'write' buffer will be written to the underlying
// bufio.Writer output destination. However, most
// importantly, the user is then responsible for
// performing the 'Close' operation by calling the local
// method:
//
//	FileBufferReadWrite.Close()
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	  Platform Conventions For Text End-Of-Line Characters
//
//	The line termination or end-of-line character, or
//	characters, written to the bufio.Writer output
//	destination are specified by input parameter
//	'writeEndOfLineChars'.
//
//	The following are the various line termination
//	conventions:
//
//	On Windows, line-endings are terminated with a
//	combination of a carriage return (ASCII 0x0d or \r)
//	and a newline(\n), also referred to as CR/LF (\r\n).
//
//	On UNIX, text file line-endings are terminated with a
//	newline character (ASCII 0x0a, represented by the \n
//	escape sequence in most languages), also referred to
//	as a linefeed (LF).
//
//	On the Mac Classic (Mac systems using any system prior
//	to Mac OS X), line-endings are terminated with a single
//	carriage return (\r or CR). (Mac OS X uses the UNIX
//	convention.)
//
//	Reference
//		portal.perforce.com/s/article/3096
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	'Read' End-of-line characters specified by input
//		parameter 'readEndOfLineDelimiters' are used to
//		parse raw data read from the bufio.Reader object and
//		extract individual lines of text.
//
//		The end-of-line delimiters specified by
//		'readEndOfLineDelimiters' are NOT written to the
//	 	output destination bufio.Writer object. They are
//	 	stripped off before being written to the
//	 	io.Writer object.
//
//	(2)	The text lines actually written
//	 	to the bufio.Writer object are terminated with the
//	 	end-of-line characters specified by user input
//	 	parameter 'writeEndOfLineChars'.
//
//	(3)	If input parameter 'autoFlushAndCloseOnExit' is
//		set to 'false', the user is responsible for
//		calling local method FileBufferReadWrite.Close()
//		in order to perform the required Clean-Up
//		tasks on the current instance of
//		FileBufferReadWrite after all read and write
//		operations have been completed. Clean-Up tasks
//		can be independently performed by calling the
//		local method:
//
//			FileBufferReadWrite.Close()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	readEndOfLineDelimiters		*StringArrayDto
//
//		A pointer to an instance of StringArrayDto.
//		'readEndOfLineDelimiters' encapsulates a string
//		array which contains one or more end-of-line
//		delimiters used to identify, parse and separate
//		individual lines of text.
//
//		Users have the flexibility to specify multiple
//		end-of-line delimiters for use in parsing text
//		lines extracted from file identified by
//		internal bufio.Reader object encapsulated by
//		the current instance of FileBufferReadWrite.
//
//		Typical text line termination, or end-of-line
//		delimiters, which may be appropriate for use
//		with a given target bufio.Reader object are
//		listed as follows:
//
//		Windows
//			Line-endings are terminated with a
//			combination of a carriage return (ASCII 0x0d
//			or \r) and a newline(\n), also referred to as
//			carriage return/line feed or CR/LF (\r\n).
//
//		UNIX/Linux
//			Text file line-endings are terminated with a
//			newline character (ASCII 0x0a, represented
//			by the \n escape sequence in most languages),
//			also referred to as a linefeed (LF).
//
//		Mac Classic Prior to Mac OS X
//			Text Line-endings are terminated with a single
//			carriage return (\r or CR).
//
//		Mac OS X or Later
//			Line termination uses the UNIX convention.
//			Text file line-endings are terminated with a
//			newline character (ASCII 0x0a, represented
//			by the \n escape sequence in most languages),
//			also referred to as a linefeed (LF).
//
//		Again, the 'read' end-of-line delimiters
//		specified by this parameter are NOT written to
//		the output destination bufio.Writer object. They
//		are stripped off before being written to the
//		internal bufio.Writer object.
//
//		The text line termination, or end-of-line
//		characters, actually written to the output
//		destination bufio.Writer object are controlled by
//		the 'write' end-of-line characters specified by
//		input parameter 'writeEndOfLineChars'.
//
//	writeEndOfLineChars			string
//
//		This string contains the end-of-line characters
//		which will be configured for each line of text
//		written to the output destination specified by
//		the internal bufio.Writer object.
//
//		On Windows, line-endings are terminated with a
//		combination of a carriage return (ASCII 0x0d or
//		\r) and a newline(\n), also referred to as CR/LF
//		(\r\n).
//
//		On UNIX, text file line-endings are terminated
//		with a newline character (ASCII 0x0a, represented
//		by the \n escape sequence in most languages),
//		also referred to as a linefeed (LF).
//
//		On the Mac Classic (Mac systems using any system
//		prior to Mac OS X), line-endings are terminated
//		with a single carriage return (\r or CR). (Mac OS
//		X uses the UNIX convention.)
//
//		If 'writeEndOfLineChars' is submitted as an empty
//		or zero length string, no end-of-line characters
//		will be written to the bufio.Writer output
//		destination and no error will be returned.
//
//	maxNumOfTextLines			int64
//
//		Specifies the maximum number of text lines which
//		will be read and processed.
//
//		If this parameter is set to a value less than
//		one (+1) (Examples: zero (0), minus-one (-1) ),
//		'maxNumOfLines' will be automatically reset to
//		the maximum positive int64 value of
//		9,223,372,036,854,775,807 (+9-Quintillion) text
//		lines. This effectively means that all text lines
//		existing in the internal bufio.Reader will be
//		read, parsed and processed.
//
//		When 'maxNumOfLines' is set to a value greater
//		than zero (0), it effectively limits the
//		maximum number of text lines which will be
//		parsed and written to the internal bufio.Writer
//		object.
//
//	initialBufferSizeBytes		int
//
//		Sets the initial size of the buffer which will
//		be used by the internal scanner to parse and
//		read individual lines of text.
//
//		The internal scanner used by this method is
//		an instance of bufio.Scanner and therefore uses
//		a buffered bufio.procedure to read data from the
//		internal bufio.Reader. The initial size of this
//		buffer is controlled by this input parameter,
//		'initialBufferSizeBytes'.
//
//		Assuming the number of bytes in a line of text
//		contained in the internal bufio.Reader is
//		reasonably standardized and constant, execution
//		time and memory resources can be minimized with
//		the careful configuration of this parameter.
//
//		If the initial buffer size proves to be too
//		small to accommodate a line of text, it will
//		be automatically increased in stages up to a
//		maximum of 65,536-bytes.
//
//		If the value of 'initialBufferSizeBytes' is
//		less than two (2-bytes), it will be automatically
//		be reset to the default initial buffer size of
//		4,096-bytes.
//
//		If the value of 'initialBufferSizeBytes' is
//		greater than 65,536 (65,536-bytes), it will be
//		automatically reset to 65,536-bytes.
//
//	autoFlushAndCloseOnExit		bool
//
//		When this parameter is set to 'true', this
//		method will automatically perform the following
//		Clean-Up tasks upon exit:
//
//		(1)	The write buffer for the internal bufio.Writer
//			object will be flushed thereby ensuring that
//			all remaining data in the 'write' buffer will
//			be written to the underlying bufio.Writer object.
//
//		(2)	The bufio.Reader and bufio.Writer objects will be
//			properly closed.
//
//		(3) After performing these Clean-Up tasks, the
//			current instance of FileBufferReadWrite will
//			become invalid and unusable for future 'read'
//			and/or 'write' operations.
//
//		If input parameter 'autoFlushAndCloseOnExit' is
//		set to 'false', this method will automatically
//		flush the 'write' buffer. This means that all
//		data remaining in the 'write' buffer will be
//		written to the underlying bufio.Writer object. Most
//		importantly, the user is then responsible for
//		performing the 'Close' operation by calling the
//		local method:
//
//			FileBufferReadWrite.Close()
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
//	numOfLinesProcessed			int64
//
//		This int64 value contains the number of text
//		lines read and parsed from the internal bufio.Reader
//		object and written to the bufio.Writer object.
//
//	numTextLineBytes			int64
//
//		If this method completes successfully, this
//		parameter will return the number of bytes read
//		as discrete lines of text. Remember that this
//		number excludes the end-of-line delimiters
//		specified by input parameter
//		'readEndOfLineDelimiters' which are stripped off
//		and deleted.
//
//		This means that in many if not most cases,
//		the number of text line bytes
//		('numTextLineBytes') will NOT match the number
//		of bytes written ('numBytesWritten').
//
//		See description of return parameter
//		'numBytesWritten' below.
//
//	numBytesWritten				int64
//
//		If this method completes successfully, this
//		integer value will equal the number of bytes
//		written to the internal bufio.Writer object. Text
//		lines written to the bufio.Writer object have new
//		line termination or end-of-line characters
//		('writeEndOfLineChars') automatically added
//		to the end of the text line string.
//		'numBytesWritten' therefore includes the
//		length of these end-of-line characters
//		('writeEndOfLineChars').
//
//		In many if not most cases, the number of text
//		line bytes shown in return parameter
//		'numTextLineBytes' will NOT match the number
//		of bytes written as shown in this return
//		parameter, 'numBytesWritten'. This discrepancy
//		occurs because 'numTextLineBytes' reflects the
//		size of text lines read from bufio.Reader after
//		their end-of-line delimiter characters have
//		been deleted or stripped off. 'numBytesWritten'
//		includes the length of these end-of-line
//		characters ('writeEndOfLineChars').
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
func (fBufReadWrite *FileBufferReadWrite) ReadWriteTextLines(
	readEndOfLineDelimiters *StringArrayDto,
	writeEndOfLineChars string,
	maxNumOfTextLines int64,
	initialBufferSizeBytes int,
	autoFlushAndCloseOnExit bool,
	errorPrefix interface{}) (
	numOfLinesProcessed int64,
	numTextLineBytes int64,
	numBytesWritten int64,
	err error) {

	if fBufReadWrite.lock == nil {
		fBufReadWrite.lock = new(sync.Mutex)
	}

	fBufReadWrite.lock.Lock()

	defer fBufReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "FileBufferReadWrite." +
		"ReadWriteTextLines()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {

		return numOfLinesProcessed,
			numTextLineBytes,
			numBytesWritten,
			err
	}

	if fBufReadWrite.reader == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: The current instance of FileBufferReadWrite\n"+
			"is invalid! The internal io.Reader object was never\n"+
			"initialized. Call one of the 'New' methods or 'Setter'\n"+
			"methods to create a valid instance of FileBufferReadWrite.\n",
			ePrefix.String())

		return numOfLinesProcessed,
			numTextLineBytes,
			numBytesWritten,
			err
	}

	if fBufReadWrite.writer == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: The current instance of FileBufferReadWrite\n"+
			"is invalid! The internal io.Writer object was never\n"+
			"initialized. Call one of the 'New' methods or 'Setter'\n"+
			"methods to create a valid instance of FileBufferReadWrite.\n",
			ePrefix.String())

		return numOfLinesProcessed,
			numTextLineBytes,
			numBytesWritten,
			err
	}

	var localNumBytesWritten, numOfEmptyReads int
	var readerLabel, writerLabel string

	var fBufReadWriteMicrobot = new(fileBufferReadWriteMicrobot)

	if len(fBufReadWrite.readerFilePathName) == 0 {

		readerLabel = "FileBufferReadWrite.IoReader"

	} else {

		readerLabel = fBufReadWrite.readerFilePathName
	}

	if len(fBufReadWrite.writerFilePathName) == 0 {

		writerLabel = "FileBufferReadWrite.IoWriter"

	} else {

		writerLabel = fBufReadWrite.writerFilePathName
	}

	if initialBufferSizeBytes < 2 {

		initialBufferSizeBytes = 4096

	} else if initialBufferSizeBytes > 65536 {

		initialBufferSizeBytes = 65536
	}

	var textLineScanner *bufio.Scanner

	textLineScanner,
		err = new(fileHelperAtom).
		getStdTextLineScanner(
			fBufReadWrite.reader,
			readerLabel,
			readEndOfLineDelimiters,
			ePrefix.XCpy("textLineScanner<-"))

	if err != nil {

		return numOfLinesProcessed,
			numTextLineBytes,
			numBytesWritten,
			err

	}

	buf := make([]byte, initialBufferSizeBytes)

	textLineScanner.Buffer(buf, bufio.MaxScanTokenSize)

	var textLine string
	var lenTextLine int
	var ok bool
	var err2, err3 error

	for {

		if numOfEmptyReads > 3 {

			break
		}

		if maxNumOfTextLines > 0 &&
			numOfLinesProcessed >= maxNumOfTextLines {

			break
		}

		err2 = nil

		ok = textLineScanner.Scan()

		if !ok {

			err2 = textLineScanner.Err()

			if err2 != nil &&
				err2 != io.EOF {

				err = fmt.Errorf("%v\n"+
					"System Errror returned by textLineScanner.Scan()\n"+
					"Error=\n%v\n",
					ePrefix.String(),
					err2)

				break
			}
		}

		textLine = textLineScanner.Text()

		lenTextLine = len(textLine)

		numTextLineBytes += int64(lenTextLine)

		if lenTextLine == 0 &&
			(errors.Is(err2, io.EOF) == true ||
				ok == false) {

			break
		}

		textLine += writeEndOfLineChars

		lenTextLine = len(textLine)

		if lenTextLine > 0 {

			numOfEmptyReads = 0

			localNumBytesWritten,
				err3 = fBufReadWrite.writer.Write(
				[]byte(textLine))

			if err3 != nil {

				err = fmt.Errorf("%v\n"+
					"Error Writing Bytes To File!\n"+
					"io.Writer= %v\n"+
					"Write Error=\n%v\n",
					funcName,
					writerLabel,
					err3.Error())

				break
			}

			numOfLinesProcessed++

			numBytesWritten += int64(localNumBytesWritten)

			// if len(textLine) > 0
		} else {

			numOfEmptyReads++
		}

		if errors.Is(err2, io.EOF) == true ||
			ok == false {

			break
		}

	} // for numOfLinesProcessed < maxNumOfTextLines

	var cleanUpStatus string

	if err != nil {

		cleanUpStatus = "After Processing Error"

	} else {

		cleanUpStatus = "Success"
	}

	if autoFlushAndCloseOnExit == true {

		err2 = fBufReadWriteMicrobot.
			flushCloseRelease(
				fBufReadWrite,
				"fBufReadWrite",
				true, // flushWriteBuffer
				true, // releaseReaderWriterMemResources
				true, // releaseFBuffReadWriteMemResources
				ePrefix.XCpy(fmt.Sprintf(
					"%v Flush/Close-Readers & Writers",
					cleanUpStatus)))

		if err2 != nil {

			err = errors.Join(err, err2)
		}

	} else {

		err2 = fBufReadWriteMicrobot.
			flushWriteBuffer(
				fBufReadWrite,
				"fBufReadWrite",
				ePrefix.XCpy(fmt.Sprintf(
					"%v Flush fBufReadWrite.writer",
					cleanUpStatus)))

		if err2 != nil {

			err = errors.Join(err, err2)
		}

	}

	return numOfLinesProcessed,
		numTextLineBytes,
		numBytesWritten,
		err
}

// SeekReader
//
// This method sets the offset for the next 'read'
// operation within the bufio.Reader object.
//
// This method only succeeds if the internal bufio.Reader
// for the current FileBufferReadWrite instance was
// created by means of a path and file name string, a
// File Manager object (FileMgr) or an io.Reader object
// with a base type of file pointer (*os.File).
//
// This target offset is interpreted according to input
// parameter 'whence'.
//
// 'whence' is an integer value designating whether the
// input parameter 'targetOffset' is interpreted to mean
// an offset from the start of the file, an offset from
// the current offset position or an offset from the end
// of the file. The 'whence' parameter must be passed as
// one of the following 'io' constant values:
//
//	io.SeekStart = 0
//		Means relative to the start of the file.
//
//	io.SeekCurrent = 1
//		Means relative to the current file offset.
//
//	io.SeekEnd = 2
//		Means relative to the end (for example,
//		offset = -2 specifies the penultimate byte of
//		the file).
//
// If the Seek method completes successfully, the next
// 'read' operation will occur at the new offset
// position.
//
// Seek returns the new offset relative to the start of the
// file or an error, if any.
//
// Seek implements the 'io.Seeker' interface.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	If the current instance of FileBufferReadWrite was
//		NOT initialized with a path and file name, a File
//		Manager (FileMgr) object or an io.Reader object
//		with a base type of *os.File, this method will
//		return an error.
//
//		Said another way, if the current instance of
//		FileBufferReadWrite was initialized with some object
//		other than a disk file, an error will be returned.
//
//	(2)	Seeking to an offset before the start of the file
//		is an error.
//
//	(3) If input parameter 'whence' is not set to one of
//		these three constant integer values, an error
//		will be returned.
//
//		io.SeekStart = 0
//			Means relative to the start of the file.
//
//		io.SeekCurrent = 1
//			Means relative to the current file offset.
//
//		io.SeekEnd = 2
//			Means relative to the end (for example,
//			offset = -2 specifies the penultimate byte of
//			the file).
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	targetOffset				int64
//
//		The number of bytes used to reset the file
//		offset for the next 'read' operation.
//
//		This offset value is interpreted according to
//		input parameter 'whence'.
//
//	whence						int
//
//		'whence' is an integer value designating whether
//		the input parameter 'targetOffset' is interpreted
//		to mean an offset from the start of the file, an
//		offset from the current offset position or an
//		offset from the end of the file. The 'whence'
//		parameter must be passed as one of the following
//		'io' constant values:
//
//		io.SeekStart = 0
//			Means relative to the start of the file.
//
//		io.SeekCurrent = 1
//			Means relative to the current file offset.
//
//		io.SeekEnd = 2
//			Means relative to the end (for example,
//			offset = -2 specifies the penultimate byte of
//			the file).
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
//	offsetFromFileStart			int64
//
//		If this method completes successfully, this
//		parameter will return the new file offset
//		in bytes from the beginning of the io.Reader
//		object
//
//	err							error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message.
func (fBufReadWrite *FileBufferReadWrite) SeekReader(
	targetOffset int64,
	whence int,
	errorPrefix interface{}) (
	offsetFromFileStart int64,
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
			"SeekReader()",
		"")

	if err != nil {

		return offsetFromFileStart, err

	}

	if fBufReadWrite.reader == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: The current instance of FileBufferReadWrite\n"+
			"is invalid! The internal bufio.Reader object was never\n"+
			"initialized. Call one of the 'New' methods or 'Setter'\n"+
			"methods to create a valid instance of FileBufferReadWrite.\n",
			ePrefix.String())

		return offsetFromFileStart, err
	}

	var err2 error

	offsetFromFileStart,
		err2 = fBufReadWrite.reader.
		Seek(
			targetOffset,
			whence)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by fBufReadWrite.reader.Seek()\n"+
			"Error:\n%v\n",
			ePrefix.String(),
			err2.Error())
	}

	return offsetFromFileStart, err
}

// SeekWriter
//
// This method sets the offset for the next 'write'
// operation within the internal bufio.Writer object.
// This method only succeeds if the current
// FileBufferReadWrite instance was created as a disk
// file with a path and file name string, a File Manager
// object (FileMgr) or an io.Writer object with a base
// type of file pointer (*os.File).
//
// This target offset is interpreted according to input
// parameter 'whence'.
//
// 'whence' is an integer value designating whether the
// input parameter 'targetOffset' is interpreted to mean
// an offset from the start of the file, an offset from
// the current offset position or an offset from the end
// of the file. The 'whence' parameter must be passed as
// one of the following 'io' constant values:
//
//	io.SeekStart = 0
//		Means relative to the start of the file.
//
//	io.SeekCurrent = 1
//		Means relative to the current file offset.
//
//	io.SeekEnd = 2
//		Means relative to the end (for example,
//		offset = -2 specifies the penultimate byte of
//		the file).
//
// If the Seek method completes successfully, the next
// 'write' operation will occur at the new offset
// position.
//
// SeekWriter returns the new offset relative to the
// start of the file or an error, if any.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	The standard technique is to call method
//		SeekWriter() and a 'write' method in tandem. In
//		this case, BE ADVISED, best practice specifies
//		calling FlushWriteBuffer() to flush the write
//		buffer BEFORE calling this method, SeekWriter().
//
//	(2)	Calling SeekWriter() to move to an offset before
//		the beginning of the 'write' file is an error.
//
//	(3) If input parameter 'whence' is not set to one of
//		these three constant integer values, an error
//		will be returned.
//
//		io.SeekStart = 0
//			Means relative to the start of the file.
//
//		io.SeekCurrent = 1
//			Means relative to the current file offset.
//
//		io.SeekEnd = 2
//			Means relative to the end (for example,
//			offset = -2 specifies the penultimate byte of
//			the file).
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	targetOffset				int64
//
//		The number of bytes used to reset the file
//		offset for the next 'write' operation.
//
//		This offset value is interpreted according to
//		input parameter 'whence'.
//
//	whence						int
//
//		'whence' is an integer value designating whether
//		the input parameter 'targetOffset' is interpreted
//		to mean an offset from the start of the file, an
//		offset from the current offset position or an
//		offset from the end of the file. The 'whence'
//		parameter must be passed as one of the following
//		'io' constant values:
//
//		io.SeekStart = 0
//			Means relative to the start of the file.
//
//		io.SeekCurrent = 1
//			Means relative to the current file offset.
//
//		io.SeekEnd = 2
//			Means relative to the end (for example,
//			offset = -2 specifies the penultimate byte of
//			the file).
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	offsetFromFileStart			int64
//
//		If this method completes successfully, this
//		parameter will return the new file offset
//		in bytes from the beginning of the file.
//
//	err							error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message.
func (fBufReadWrite *FileBufferReadWrite) SeekWriter(
	targetOffset int64,
	whence int) (
	offsetFromFileStart int64,
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
			"SeekWriter()",
		"")

	if err != nil {

		return offsetFromFileStart, err

	}

	if fBufReadWrite.writer == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: The current instance of FileBufferReadWrite\n"+
			"is invalid! The internal bufio.Writer object was never\n"+
			"initialized. Call one of the 'New' methods or 'Setter'\n"+
			"methods to create a valid instance of FileBufferReadWrite.\n",
			ePrefix.String())

		return offsetFromFileStart, err
	}

	var err2 error

	offsetFromFileStart,
		err2 = fBufReadWrite.writer.
		Seek(
			targetOffset,
			whence)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by fBufReadWrite.writer.Seek()\n"+
			"Error:\n%v\n",
			ePrefix.String(),
			err2.Error())
	}

	return offsetFromFileStart, err
}

// SetFileMgrsReadWrite
//
// Receives two input parameters of type FileMgr. These
// File Manager instances will be used to reconfigure the
// internal bufio.Reader and bufio.Writer objects
// encapsulated by the current instance of
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
// internal bufio.Reader object encapsulated in the
// current instance of FileBufferReadWrite. This internal
// bufio.Reader object will be reconfigured using the
// file identified by a File Manager instance passed as
// input parameter 'readerFileMgr'.
//
// The internal bufio.Reader object is used to 'read'
// data from a data source such as a disk file.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reconfigure
//	the member variable bufio.Reader object encapsulated
//	in the current instance of FileBufferReadWrite.
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
// internal bufio.Writer object encapsulated in the
// current instance of FileBufferReadWrite. This internal
// bufio.Writer object will be reconfigured using the
// file identified by a File Manager instance passed as
// input parameter 'writerFileMgr'.
//
// The internal bufio.Writer object is used to 'write'
// data to an output destination such as a disk file.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reconfigure
//	the member variable bufio.Writer object encapsulated
//	in the current instance of FileBufferReadWrite.
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
// internal bufio.Reader and bufio.Writer objects
// encapsulated in the current instance of
// FileBufferReadWrite.
//
// The internal bufio.Reader object is used to 'read'
// data from a data source such as a disk file. This
// internal bufio.Reader object will be reconfigured
// using the io.Reader object passed as input parameter
// 'reader'.
//
// In contrast, the internal bufio.Writer object is used
// to 'write' data to an output destination such as a
// disk file. This internal bufio.Writer object will be
// reconfigured using the io.Writer object passed as
// input parameter 'writer'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing member data values in the current
//	instance of FileBufferReadWrite.
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
//		performing any other required Clean-Up operations
//		in addition to calling local method:
//
//		FileBufferReadWrite.Close()
//			This method will flush the 'write' buffer
//			in addition to closing and performing
//			Clean-Up tasks for the io.Reader and
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
//		other required Clean-Up operations in addition to
//		calling local method:
//
//		FileBufferReadWrite.Close()
//			This method will flush the 'write' buffer
//			in addition to closing and performing
//			Clean-Up tasks for the io.Reader and
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
// internal bufio.Reader object encapsulated in the
// current instance of FileBufferReadWrite. The internal
// bufio.Reader object will be reconfigured using the
// io.Reader object passed as input parameter 'reader'.
//
// The internal bufio.Reader object, encapsulated in the
// current instance of FileBufferReadWrite, is used to
// 'read' data from a data source such as a disk file.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reconfigure
//	the internal bufio.Reader object encapsulated in
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
//		performing any other required Clean-Up operations
//		in addition to calling local method:
//
//		FileBufferReadWrite.Close()
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
// internal bufio.Writer object encapsulated in the
// current instance of FileBufferReadWrite. The internal
// bufio.Writer object will be reconfigured using the
// io.Writer object passed as input parameter 'writer'.
//
// The internal bufio.Writer object, encapsulated in the
// current instance of FileBufferReadWrite, is used to
// 'write' data to an output destination such as a disk
// file.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reconfigure
//	the internal bufio.Writer object encapsulated in the
//	current instance of FileBufferReadWrite.
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
//		performing any other required Clean-Up operations
//		in addition to calling local method:
//
//		FileBufferReadWrite.Close()
//			This method will flush the 'write' buffer
//			in addition to closing and performing
//			Clean-Up tasks for the io.Reader and
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
// configured as a data source for 'read' operations
// performed by the current instance of
// FileBufferReadWrite.
//
// This 'read' file will be configured as an internal
// io.Reader object encapsulated in the current instance
// of FileBufferReadWrite.
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
// internal bufio.Writer object for the current instance
// of FileBufferReadWrite.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reconfigure
//	the member variable bufio.Writer object encapsulated
//	in	the current instance of FileBufferReadWrite.
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
// method to ensure Clean-Up operations are properly
// applied:
//
//	FileBufferReadWrite.Close()
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
//	(2) This method DOES NOT flush the write buffer.
//
//	(3)	After all 'read' and 'write' operations have been
//		completed, the user MUST call the 'Close' method
//		to perform necessary Clean-Up operations:
//
//		FileBufferReadWrite.Close()
//
//	(4) This method WILL NOT VERIFY that the number of
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

// WriteTextOrNumbers
//
// This method will accept many different text or numeric
// data types which are then converted to a byte or
// string array and written to the internal bufio.Writer
// object encapsulated in the current instance of
// FileBufferReadWrite.
//
// The internal bufio.Writer is usually configured by
// calling one of the 'New' methods for the current
// instance of FileBufferReadWrite.
//
// The text or numeric data type passed as input
// parameter 'charsToWrite' must match one of over sixty
// eligible data types.
//
// If 'charsToWrite' is set to an ineligible data type,
// an error will be returned.
//
// Don't forget to 'close' the current FileBufferReadWrite
// instance when you are finished with it.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	charsToWrite				interface{}
//
//		This empty interface is used to transmit an
//		eligible text or numeric data type which will be
//		to a string or byte array and written to the
//		io.Writer object passed as input parameter
//		'ioWriter'.
//
//		If the type transmitted through this parameter
//		does not one of the following data types, an
//		error will be returned.
//
//			Eligible Data Types
//
//			 1.	[]byte
//			 2.	*[]byte
//			 3.	string
//			 4.	*string
//			 5.	[]string
//			 6.	*[]string
//			 7.	Stringer (fmt.Stringer) Interface
//			 8.	strings.Builder
//			 9.	*strings.Builder
//			10.	StringArrayDto
//			11.	*StringArrayDto
//			12.	[]rune
//			13.	*[]rune
//			14.	RuneArrayDto
//			15.	*RuneArrayDto
//			16.	RuneArrayCollection
//			17.	*RuneArrayCollection
//			18.	ITextFieldFormatDto
//			19.	ITextFieldSpecification
//			20.	ITextLineSpecification
//			21.	TextLineSpecLinesCollection
//			22.	bool
//			23.	TextLineTitleMarqueeDto
//			24.	time.Time
//			25.	TextInputParamFieldDateTimeDto
//			26.	float32
//			27.	*float32
//			28.	float64
//			29.	*float64
//			30.	BigFloatDto
//			31.	*BigFloatDto
//			32.	big.Float
//			33.	*big.Float
//			34.	big.Rat
//			35.	*big.Rat
//			36.	int8
//			37.	*int8
//			38.	int16
//			39.	*int16
//			40.	int
//			41.	*int
//			42.	int32
//			43.	*int32
//			44.	int64
//			45.	*int64
//			46.	uint8
//			47.	*uint8
//			48.	uint16
//			49.	*uint16
//			50.	uint
//			51.	*uint
//			52.	uint32
//			53.	*uint32
//			54.	uint64,
//			55.	*uint64
//			56.	big.Int
//			57.	*big.Int
//			58.	TextFieldFormatDtoFloat64
//			59.	*TextFieldFormatDtoFloat64
//			60.	TextFieldFormatDtoBigFloat
//			61.	*TextFieldFormatDtoBigFloat
//			62.	NumberStrKernel
//			63.	*NumberStrKernel
//			64.	[]NumberStrKernel
//			65.	*[]NumberStrKernel
//
//	writeEndOfLineChars string
//
//		This character string is appended to each line of
//		text written to the bufio.Writer object. This
//		capability is more useful when processing string
//		arrays when each element of the array is written
//		separately to the io.Writer object.
//
//		Remember that on Windows, line-endings are
//		terminated with a combination of a carriage
//		return (ASCII 0x0d or \r) and a newline(\n), also
//		referred to as CR/LF (\r\n).
//
//		On UNIX or Linux, text file line-endings are
//		terminated with a newline character (ASCII 0x0a,
//		represented by the \n escape sequence in most
//		languages), also referred to as a linefeed (LF).
//
//		If 'writeEndOfLineChars' is set to an empty
//		string, it will be ignored and no additional
//		characters will be appended to each line written
//		to the bufio.Writer object.
//
//	writeEndOfTextChars			string
//
//		A character string which will be written to the
//		internal bufio.Writer object after all other text
//		from 'charsToWrite' and 'writeEndOfLineChars'
//		has been processed and written.
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
//	numOfBytesWritten			int64
//
//
//		Returns the number of bytes extracted from the
//		string array contained in input parameter
//		'strArrayDto' and written to the internal
//		bufio.Writer object encapsulated by the current
//		instance of FileBufferReadWrite.
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
func (fBufReadWrite *FileBufferReadWrite) WriteTextOrNumbers(
	charsToWrite interface{},
	writeEndOfLineChars string,
	writeEndOfTextChars string,
	errorPrefix interface{}) (
	numOfBytesWritten int64,
	err error) {

	if fBufReadWrite.lock == nil {
		fBufReadWrite.lock = new(sync.Mutex)
	}

	fBufReadWrite.lock.Lock()

	defer fBufReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "FileBufferReadWrite." +
		"WriteTextOrNumbers()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {

		return numOfBytesWritten, err
	}

	if fBufReadWrite.writer == nil {

		err = fmt.Errorf("%v\n"+
			"---------------------------------------------------------\n"+
			"Error: This instance of 'FileBufferReadWrite' is invalid!\n"+
			"The internal bufio.Writer has NOT been initialized.\n"+
			"Call one of the 'New' or 'Setter' methods when creating\n"+
			"a new valid instance of 'FileBufferReadWrite'.\n",
			ePrefix.String())

		return numOfBytesWritten, err
	}

	var writeBytesFunc = fBufReadWrite.lowLevelWriteBytes

	numOfBytesWritten,
		err = new(fileWriterHelperMicrobot).
		writeCharacters(
			writeBytesFunc,
			charsToWrite,
			"charsToWrite",
			writeEndOfLineChars,
			writeEndOfTextChars,
			ePrefix.XCpy("fBufReadWrite.writer<-charsToWrite"))

	if err != nil {

		return numOfBytesWritten, err
	}

	var fBufWriterAtom = new(fileBufferWriterAtom)
	var err2 error

	err2 = fBufWriterAtom.flush(
		fBufReadWrite.writer,
		"fBufReadWrite.writer",
		ePrefix.XCpy("Flush fBufReadWrite.writer"))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by fBufWriterMolecule.flush(fBufReadWrite.writer)\n"+
			"Error= \n%v\n",
			funcName,
			err2.Error())
	}

	return numOfBytesWritten, err
}

// lowLevelWriteBytes
//
// This method is designed to be passed as a text handler
// and text writer for method:
//
//	fileWriterHelperMicrobot.writeCharacters()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	bytesToWrite				[]byte
//
//		A byte array containing the characters which will
//		be written to the bufio.Writer encapsulated in
//		the current instance of FileBufferReadWrite.
//
//	writeEndOfTextChars			string
//
//		If the length of string 'writeEndOfTextChars' is
//		greater than zero (0), the characters contained
//		therein will be the last characters written to
//		the current instance of FileBufferReadWrite.
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
//	numOfBytesWritten			int64
//
//		If this method completes successfully, this int64
//		value will contain the number of characters
//		written to the current instance of
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
//	 	text passed by input parameter, 'errPrefDto'.
//	 	The 'errPrefDto' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fBufReadWrite *FileBufferReadWrite) lowLevelWriteBytes(
	bytesToWrite []byte,
	writeEndOfTextChars string,
	errPrefDto *ePref.ErrPrefixDto) (
	numOfBytesWritten int64,
	err error) {
	//
	//if fBufReadWrite.lock == nil {
	//	fBufReadWrite.lock = new(sync.Mutex)
	//}
	//
	//fBufReadWrite.lock.Lock()
	//
	//defer fBufReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "FileBufferReadWrite." +
		"lowLevelWriteBytes()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		return numOfBytesWritten, err
	}

	if fBufReadWrite.writer == nil {

		err = fmt.Errorf("%v\n"+
			"---------------------------------------------------------\n"+
			"Error: This instance of 'FileBufferReadWrite' is invalid!\n"+
			"The internal bufio.Writer has NOT been initialized.\n"+
			"Call one of the 'New' or 'Setter' methods when creating\n"+
			"a new valid instance of 'FileBufferReadWrite'\n",
			ePrefix.String())

		return numOfBytesWritten, err
	}

	lenBytesToWrite := len(bytesToWrite)

	if lenBytesToWrite == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'bytesToWrite' is invalid!\n"+
			"The 'bytesToWrite' byte array is empty. It has zero bytes.\n",
			ePrefix.String())

		return numOfBytesWritten, err
	}

	var err2 error
	var localNumOfBytesWritten int
	var expectedNumOfBytesToWrite int64

	expectedNumOfBytesToWrite = int64(lenBytesToWrite)

	localNumOfBytesWritten,
		err2 = fBufReadWrite.writer.Write(bytesToWrite)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by fBufReadWrite.writer.Write(bytesToWrite).\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			err2.Error())

		return numOfBytesWritten, err
	}

	if expectedNumOfBytesToWrite != int64(localNumOfBytesWritten) {

		err = fmt.Errorf("%v\n"+
			"Error returned by fBufReadWrite.writer.Write(bytesToWrite)\n"+
			"while attempting to write a line or collection of text characters.\n"+
			"ERROR: Expected number of bytes to write NOT equal to actual number\n"+
			"of bytes written!\n"+
			"Expected number of bytes to write= %v\n"+
			"Actual Number of Bytes Written= %v\n"+
			"Attempted to write byte array (string):\n"+
			"%v",
			ePrefix.String(),
			expectedNumOfBytesToWrite,
			localNumOfBytesWritten,
			string(bytesToWrite))

		return numOfBytesWritten, err
	}

	numOfBytesWritten +=
		int64(localNumOfBytesWritten)

	lenWriteEndOfTextChars := len(writeEndOfTextChars)

	if lenWriteEndOfTextChars > 0 {

		expectedNumOfBytesToWrite += int64(lenWriteEndOfTextChars)

		localNumOfBytesWritten,
			err2 = fBufReadWrite.writer.Write(
			[]byte(writeEndOfTextChars))

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Error returned by fBufReadWrite.writer.Write(writeEndOfTextChars).\n"+
				"writeEndOfTextChars= '%v'\n"+
				"Error=\n%v\n",
				ePrefix.String(),
				writeEndOfTextChars,
				err2.Error())

			return numOfBytesWritten, err
		}

		if expectedNumOfBytesToWrite != int64(localNumOfBytesWritten) {

			err = fmt.Errorf("%v\n"+
				"Error returned by fBufReadWrite.writer.Write(writeEndOfTextChars)\n"+
				"while attempting to write the end of line or end of text characters.\n"+
				"ERROR: Expected number of bytes to write NOT equal to actual number\n"+
				"of bytes written!\n"+
				"Expected number of bytes to write= %v\n"+
				"Actual Number of Bytes Written= %v\n"+
				"Attempted to write byte array (string):\n"+
				"%v",
				ePrefix.String(),
				expectedNumOfBytesToWrite,
				localNumOfBytesWritten,
				writeEndOfTextChars)

			return numOfBytesWritten, err
		}

		numOfBytesWritten +=
			int64(localNumOfBytesWritten)
	}

	return numOfBytesWritten, err
}

type fileBufferReadWriteMicrobot struct {
	lock *sync.Mutex
}

// flushCloseRelease
//
// This method will perform all required Clean-Up
// operations on an instance of FileBufferReadWrite
// passed as input parameter 'fBufReadWrite'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	After calling this method the FileBufferReadWrite
//	instance passed as 'fBufReadWrite' will be invalid
//	and unavailable for any future read/write operations.
//
//	The specific Clean-Up procedures performed by this
//	method are listed as follows:
//
//	(1)	Flushing the internal write buffer.
//		This guarantees that any data remaining in the
//		'write' will be written to the underlying
//		bufio.writer object.
//
//	(2)	Closing the internal bufio.reader object.
//
//	(3)	Closing the internal bufio.writer object.
//
//	(4) Releasing all internal memory resources.
//		This action will synchronize internal flags and
//		prevent multiple calls to 'close' methods.
//		Performing a 'close' operation multiple times
//		on a single bufio.reader or bufio.writer object
//		can produce unexpected results.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fBufReadWrite						*FileBufferReadWrite
//
//		A pointer to an instance of FileBufferReadWrite.
//
//		This method will perform Clean-Up operations on
//		the internal bufio.reader and bufio.writer objects
//		encapsulated in	this FileBufferReadWrite instance.
//
//		This method will effectively render the
//		FileBufferReadWrite instance 'fBufReadWrite'
//		invalid and unusable for any future 'read' and/or
//		'write' operations.
//
//	fBufReadWriteLabel					string
//
//		The name or label associated with input parameter
//		'fBufReadWrite' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fBufReadWrite" will
//		be automatically applied.
//
//	flushWriteBuffer					bool
//
//		If this parameter is set to 'true', this method
//		will flush the contents of the internal
//		bufio.Writer 'write' buffer. This means that
//		'write' buffer contents are guaranteed to be
//		written to the internal bufio.Writer object
//		encapsulated by	FileBufferReadWrite input
//		parameter 'fBufReadWrite'.
//
//		If 'flushWriteBuffer' is set to 'false', the
//		contents of the 'write' buffer will be lost.
//
//	releaseReaderWriterMemResources		bool
//
//		If this parameter is set to 'true', this method
//		will release the memory resources for the
//		internal bufio.reader and bufio.writer objects
//		encapsulated by 'fBufReadWrite'.
//
//	releaseFBuffReadWriteMemResources	bool
//
//		If this parameter is set to 'true', this method
//		will release the internal memory resources for
//		the FileBufferReadWrite instance passed as
//		'fBufReadWrite'.
//
//	errPrefDto							*ePref.ErrPrefixDto
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
func (fBufReadWriteMicrobot *fileBufferReadWriteMicrobot) flushCloseRelease(
	fBufReadWrite *FileBufferReadWrite,
	fBufReadWriteLabel string,
	flushWriteBuffer bool,
	releaseReaderWriterMemResources bool,
	releaseFBuffReadWriteMemResources bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fBufReadWriteMicrobot.lock == nil {
		fBufReadWriteMicrobot.lock = new(sync.Mutex)
	}

	fBufReadWriteMicrobot.lock.Lock()

	defer fBufReadWriteMicrobot.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileBufferReadWriteMicrobot." +
		"flushCloseRelease()"

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

	var err2, err3 error
	var fBuffReadWriteElectron = new(fileBufferReadWriteElectron)

	err2 = fBuffReadWriteElectron.
		writerFlushCloseRelease(
			fBufReadWrite,
			fBufReadWriteLabel,
			flushWriteBuffer,
			releaseReaderWriterMemResources,
			releaseFBuffReadWriteMemResources,
			ePrefix)

	if err2 != nil {

		err3 = fmt.Errorf("%v\n"+
			"Error occurred while closing the %v.\n"+
			"Error:\n%v\n",
			funcName,
			fBufReadWriteLabel+".writer",
			err2.Error())

		err = errors.Join(err3)
	}

	err2 = fBuffReadWriteElectron.
		readerCloseRelease(
			fBufReadWrite,
			fBufReadWriteLabel,
			releaseReaderWriterMemResources,
			releaseFBuffReadWriteMemResources,
			ePrefix)

	if err2 != nil {

		err3 = fmt.Errorf("%v\n"+
			"Error occurred while closing the %v.\n"+
			"Error:\n%v\n",
			funcName,
			fBufReadWriteLabel+".reader",
			err2.Error())

		err = errors.Join(err3)
	}

	return err
}

// flushWriteBuffer
//
// This method will flush the write buffer to ensure that
// all data is written to the underlying bufio.Writer which
// encapsulates the output destination.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fBufReadWrite						*FileBufferReadWrite
//
//		A pointer to an instance of FileBufferWriter.
//
//		This method will flush the contents of the internal
//		bufio.Writer 'write' buffer encapsulated by
//		'fBufReadWrite'. This means that 'write' buffer
//		contents are guaranteed to be written to the
//		internal bufio.Writer object containing the data
//		output destination.
//
//	fBufReadWriteLabel					string
//
//		The name or label associated with input parameter
//		'fBufReadWrite' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fBufReadWrite" will
//		be automatically applied.
//
//	errPrefDto							*ePref.ErrPrefixDto
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
func (fBufReadWriteMicrobot *fileBufferReadWriteMicrobot) flushWriteBuffer(
	fBufReadWrite *FileBufferReadWrite,
	fBufReadWriteLabel string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fBufReadWriteMicrobot.lock == nil {
		fBufReadWriteMicrobot.lock = new(sync.Mutex)
	}

	fBufReadWriteMicrobot.lock.Lock()

	defer fBufReadWriteMicrobot.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileBufferReadWriteMicrobot." +
		"flushWriteBuffer()"

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

	if fBufReadWrite.writer == nil {

		err = fmt.Errorf("%v\n"+
			"Error: The internal '%v.writer' object is a nil pointer!\n"+
			"%v is invalid and was NOT properly initialized.\n",
			ePrefix.String(),
			fBufReadWriteLabel,
			fBufReadWriteLabel)

		return err
	}

	var err2 error

	err2 = fBufReadWrite.writer.
		Flush(ePrefix.XCpy(fBufReadWriteLabel + ".writer"))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"An error occurred while flushing the write buffer\n"+
			"for the internal bufio.Writer object, %v.writer.\n"+
			"Error:\n%v\n",
			funcName,
			fBufReadWriteLabel,
			err2.Error())

	}

	return err
}

// setFileMgrsReadWrite
//
// Receives two instances of FileMgr as input parameters
// identifying the internal bufio.Reader and bufio.Writer
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
//		A pointer to an instance of FileBufferReadWrite.
//
//		The internal bufio.Reader and bufio.Writer objects
//		encapsulated in this instance will be deleted and
//		reinitialized using the 'read' and 'write' files
//		passed as input parameters 'readerFileMgr' and
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
//		'readerFileMgr' will be opened for 'read' and
//		'write' operations.
//
//		If 'openFileReadWrite' is set to 'false', the
//		target read file will be opened for 'read-only'
//		operations.
//
//	readerBuffSize				int
//
//		This integer value controls the size of the
//		'read' buffer created for the bufio.Reader object
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
//		output destination for 'write' operations
//		performed by the instance of FileBufferReadWrite
//		passed as input parameter 'fBufReadWrite'.
//
//		If the path and file name encapsulated by
//		'writerFileMgr' do not currently exist on an
//		attached storage drive, this method will attempt
//		to create them.
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
//		performing any other required Clean-Up operations
//		in addition to calling the local method:
//
//		FileBufferReadWrite.Close()
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
//		other required Clean-Up operations in addition to
//		calling local method:
//
//		FileBufferReadWrite.Close()
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

	var fBufReadWriteAtom = new(fileBufferReadWriteAtom)

	err = fBufReadWriteAtom.
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

	err = fBufReadWriteAtom.
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
//		A pointer to an instance of FileBufferReadWrite.
//
//		The internal FileBufferReader and
//		FileBufferWriter objects encapsulated in this
//		instance will be deleted and reinitialized using
//		the path and file names passed as input
//		parameters 'readerPathFileName' and
//		'writerPathFileName'.
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
//		the file which will be configured as the internal
//		bufio.Reader instance and used a data source for
//		all 'read' operations.
//
//		If this file does not currently exist on an
//		attached storage drive, an error will be
//		returned.
//
//	readerPathFileNameLabel			string
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
//	writerPathFileNameLabel			string
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
// will be used to configure a bufio.Reader object
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
//	parameter 'fBufReadWrite':
//
//		fBufReadWrite.reader
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fBufReadWrite				*FileBufferReadWrite
//
//		A pointer to an instance of FileBufferReadWrite.
//
//		The internal bufio.Reader object encapsulated in
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
//		configured as an internal bufio.Reader for the
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
			"Reader FileMgr %v failed the validity test.\n"+
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
//	parameter 'fBufReadWrite':
//
//		fBufReadWrite.reader
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fBufReadWrite				*FileBufferReadWrite
//
//		A pointer to an instance of FileBufferReadWrite.
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
//		performing any other required Clean-Up operations
//		in addition to calling the local method:
//
//		FileBufferReadWrite.Close()
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

	if reader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is a nil pointer!\n"+
			"%v is invalid.\n",
			ePrefix.String(),
			readerLabel,
			readerLabel)

		return err
	}

	err = new(fileBufferReadWriteElectron).readerCloseRelease(
		fBufReadWrite,
		fBufReadWriteLabel,
		true, // releaseMemoryResources
		true, // releaseFBuffReaderLocalMemRes
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

	fBufReadWrite.readerFilePathName =
		newBuffReader.targetReadFileName

	return err
}

// setIoWriter
//
// Receives an object which implements io.Writer
// interface. This object is then used to configure
// the internal io.Writer member variable encapsulated in
// the FileBufferReadWrite instance passed as input
// parameter 'fBufReadWrite'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reconfigure
//	the member variable io.Writer object encapsulated in
//	the instance of FileBufferReadWrite passed as input
//	parameter 'fBufReadWrite':
//
//		fBufReadWrite.writer
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fBufReadWrite				*FileBufferReadWrite
//
//		A pointer to an instance of FileBufferReadWrite.
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
//		other required Clean-Up operations in addition to
//		calling local method:
//
//		FileBufferReadWrite.Close()
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

	if writer == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is a nil pointer!\n"+
			"%v is invalid.\n",
			ePrefix.String(),
			writerLabel,
			writerLabel)

		return err
	}

	err = new(fileBufferReadWriteElectron).
		writerFlushCloseRelease(
			fBufReadWrite,
			fBufReadWriteLabel,
			true, // flushWriteBuffer
			true, // releaseMemoryResources
			true, // releaseFBuffWriterLocalMemRes
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

	fBufReadWrite.writerFilePathName =
		newBuffWriter.targetWriteFileName

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
//			fBufReadWrite.reader
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
//		deleted and configured using the file identified
//		by input parameter 'readerPathFileName'.
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

	err = new(fileBufferReadWriteElectron).readerCloseRelease(
		fBufReadWrite,
		fBufReadWriteLabel,
		true, // releaseMemoryResources
		true, // releaseFBuffReaderLocalMemRes
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
			fBufReadWriteLabel+".newBuffReader",
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
//			fBufReadWrite.writer
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

	err = new(fileBufferReadWriteElectron).
		writerFlushCloseRelease(
			fBufReadWrite,
			fBufReadWriteLabel,
			true, // flushWriteBuffer
			true, // releaseWriterMemResources
			true, // releaseFBuffWriterLocalMemRes
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

// empty
//
// This method deletes all internal member variables and
// releases all the internal memory resources for an
// instance of FileBufferReadWrite passed as input
// parameter 'fBufReadWrite'.
//
// Specifically the following internal member variables
// are set to 'nil' or their initial zero values:
//
//	FileBufferReadWrite.reader = nil
//	FileBufferReadWrite.writer = nil
//	FileBufferReadWrite.readerFilePathName = ""
//	FileBufferReadWrite.writerFilePathName = ""
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fBufReadWrite				*FileBufferReadWrite
//
//		A pointer to an instance of FileBufferReadWrite.
//
//		All internal member variable data values in
//		this instance will be deleted and reset to
//		their initial zero values.
//
//		All member variable object pointers will be set
//		to 'nil'.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	--- NONE ---
func (fBuffReadWriteElectron *fileBufferReadWriteElectron) empty(
	fBufReadWrite *FileBufferReadWrite) {

	if fBuffReadWriteElectron.lock == nil {
		fBuffReadWriteElectron.lock = new(sync.Mutex)
	}

	fBuffReadWriteElectron.lock.Lock()

	defer fBuffReadWriteElectron.lock.Unlock()

	if fBufReadWrite == nil {

		return
	}

	fBufReadWrite.reader = nil

	fBufReadWrite.writer = nil

	fBufReadWrite.readerFilePathName = ""

	fBufReadWrite.writerFilePathName = ""

	return
}

// isFileBufferReadWriteValid
//
// This method receives a pointer to an instance of
// FileBufferReadWrite ('fBufReadWrite') which will be
// analyzed to determine if all the member variables
// contain valid values.
//
// If input parameter 'fBufReadWrite' is determined to be
// invalid, this method returns an error.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fBufReadWrite				*FileBufferReadWrite
//
//		A pointer to an instance of FileBufferReadWrite.
//
//		If any of the internal member variable data
//		values encapsulated in 'fBufReadWrite' are
//		determined to be invalid, this method will return
//		an error.
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
//		If any of the internal member data variables
//		contained in the instance of FileBufferReadWrite
//		passed as 'fBufReadWrite' are found to be
//		invalid, this method will return an error
//		configured with an appropriate message
//		identifying the invalid	member data variable.
//
//		If all internal member data variables evaluate
//		as valid, this returned error value will be set
//		to 'nil'.
//
//		If errors are encountered during processing or if
//		any 'fBufReadWrite' internal member data values
//		are found to be invalid, the returned error Type
//		will encapsulate an appropriate error message.
//	 	This returned error message will incorporate the
//	 	method chain and text passed by input parameter,
//	 	'errorPrefix'. The 'errorPrefix' text will be
//	 	prefixed to the beginning of the error message.
func (fBuffReadWriteElectron *fileBufferReadWriteElectron) isFileBufferReadWriteValid(
	fBufReadWrite *FileBufferReadWrite,
	fBufReadWriteLabel string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fBuffReadWriteElectron.lock == nil {
		fBuffReadWriteElectron.lock = new(sync.Mutex)
	}

	fBuffReadWriteElectron.lock.Lock()

	defer fBuffReadWriteElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	funcName := "fileBufferReadWriteElectron." +
		"isFileBufferReadWriteValid()"

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

	if fBufReadWrite.reader == nil &&
		fBufReadWrite.writer == nil {

		err = fmt.Errorf("%v\n"+
			" -----------------------------------------------------------\n"+
			" ERROR: The %v instance of FileBufferReadWrite\n"+
			" is invalid! The internal bufio.Reader and bufio.Writer objects\n"+
			" were never initialized. Call one of the 'New' methods or\n"+
			" 'Setter' methods to create a valid instance of\n"+
			" FileBufferReadWrite.\n",
			ePrefix.String(),
			fBufReadWriteLabel)

		return err
	}

	if fBufReadWrite.reader == nil {

		err = fmt.Errorf("%v\n"+
			" -----------------------------------------------------------\n"+
			" ERROR: The %v instance of FileBufferReadWrite\n"+
			" is invalid! The internal io.Reader object was never\n"+
			" initialized. Call one of the 'New' methods or 'Setter'\n"+
			" methods to create a valid instance of FileBufferReadWrite.\n",
			ePrefix.String(),
			fBufReadWriteLabel)

		return err
	}

	if fBufReadWrite.writer == nil {

		err = fmt.Errorf("%v\n"+
			" -----------------------------------------------------------\n"+
			" ERROR: The %v instance of FileBufferReadWrite is invalid!\n"+
			" The internal io.Writer object was never initialized.\n"+
			" Call one of the 'New' methods or 'Setter' methods to\n"+
			" create a valid instance of FileBufferReadWrite.\n",
			ePrefix.String(),
			fBufReadWriteLabel)

		return err
	}

	var err2 error

	err2 = new(fileBufferReaderAtom).isFileBufferReaderValid(
		fBufReadWrite.reader,
		fBufReadWriteLabel+".reader",
		ePrefix)

	if err2 != nil {

		err = fmt.Errorf(" %v\n"+
			" -----------------------------------------------------------------------\n"+
			" ERROR: The %v instance of FileBufferReadWrite is invalid!\n"+
			" The internal bufio.reader returned a validation error.\n"+
			" Error:\n"+
			"%v\n",
			funcName,
			fBufReadWriteLabel,
			err2.Error())

		return err
	}

	err2 = new(fileBufferWriterElectron).isFileBufferWriterValid(
		fBufReadWrite.writer,
		fBufReadWriteLabel+".writer",
		ePrefix)

	if err2 != nil {

		err = fmt.Errorf(" %v\n"+
			" -----------------------------------------------------------------------\n"+
			" ERROR: The %v instance of FileBufferReadWrite is invalid!\n"+
			" The internal bufio.Writer returned a validation error.\n"+
			" Error:\n"+
			"%v\n",
			funcName,
			fBufReadWriteLabel,
			err2.Error())

	}

	return err
}

// readerCloseRelease
//
// This method will perform Clean-Up operations on the
// internal bufio.Reader object encapsulated in the
// FileBufferReadWrite instance passed as input parameter
// 'fBufReadWrite':
//
//	fBufReadWrite.reader
//
// Upon completion, method will effectively render the
// 'fBufReadWrite' instance	invalid and unusable for any
// future 'read' operations.
//
// This operation is accomplished by closing the internal
// bufio.Reader object and release the memory resources
// associated with that object:
//
//	fBufReadWrite.reader
//
// The 'release memory resources' actions are implemented
// independently based on the values passed for input
// parameters 'releaseReaderMemResources', and
// 'releaseFBuffReaderLocalMemRes'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fBufReadWrite						*FileBufferReadWrite
//
//		A pointer to an instance of FileBufferWriter.
//
//		This method will perform Clean-Up operations on
//		the internal bufio.Reader object encapsulated in
//		this FileBufferReadWrite instance:
//
//			fBufReadWrite.reader
//
//		Upon completion, this method will effectively
//		render the FileBufferReadWrite instance,
//		'fBufReadWrite', invalid and unusable for any
//		future 'read' operations.
//
//	fBufReadWriteLabel					string
//
//		The name or label associated with input parameter
//		'fBufReadWrite' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fBufReadWrite" will
//		be automatically applied.
//
//	releaseReaderMemResources			bool
//
//		If this parameter is set to 'true', this method
//		will release the internal memory resources for
//		the	internal bufio.Reader object encapsulated by
//		'fBufReadWrite':
//
//			fBufReadWrite.reader
//
//		Releasing internal memory resources synchronizes
//		internal flags and prevents multiple calls to the
//		'close' method. Calling the 'close' method more
//		than once may produce unexpected results.
//
//	releaseFBuffReaderLocalMemRes		bool
//
//		If 'releaseFBuffReaderLocalMemRes' is set to
//		'true', this method will release the local memory
//		resources for the FileBufferReadWrite reader object
//		(fBufReadWrite):
//
//			fBufReadWrite.reader = nil
//			fBufReadWrite.readerFilePathName = ""
//
//		Releasing internal memory resources synchronizes
//		internal flags and prevents multiple calls to the
//		'close' method. Calling the 'close' method more
//		than once may produce unexpected results.
//
//	errPrefDto							*ePref.ErrPrefixDto
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
func (fBuffReadWriteElectron *fileBufferReadWriteElectron) readerCloseRelease(
	fBufReadWrite *FileBufferReadWrite,
	fBufReadWriteLabel string,
	releaseReaderMemResources bool,
	releaseFBuffReaderLocalMemRes bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fBuffReadWriteElectron.lock == nil {
		fBuffReadWriteElectron.lock = new(sync.Mutex)
	}

	fBuffReadWriteElectron.lock.Lock()

	defer fBuffReadWriteElectron.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileBufferReadWriteElectron." +
		"readerCloseRelease()"

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

		err2 = new(fileBufferReaderMolecule).closeAndRelease(
			fBufReadWrite.reader,
			fBufReadWriteLabel+".reader",
			releaseReaderMemResources,
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

	if releaseFBuffReaderLocalMemRes == true {

		fBufReadWrite.reader = nil
		fBufReadWrite.readerFilePathName = ""

	}

	return err
}

// writerFlushCloseRelease
//
// This method will perform Clean-Up operations on the
// internal bufio.Writer object encapsulated in the
// FileBufferReadWrite instance passed as input parameter
// 'fBufReadWrite':
//
//	fBufReadWrite.writer
//
// Upon completion, method will effectively render the
// 'fBufReadWrite' instance	invalid and unusable for any
// future 'write' operations.
//
// This operation is accomplished by flushing and closing
// the internal bufio.Writer object before finally
// releasing the memory resources associated with that
// object:
//
//	fBufReadWrite.writer
//
// The 'flush' and 'release memory resources' actions are
// implemented independently based on the values passed
// for input parameters 'flushWriteBuffer',
// 'releaseWriterMemResources', and
// 'releaseFBuffWriterLocalMemRes'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fBufReadWrite						*FileBufferReadWrite
//
//		A pointer to an instance of FileBufferWriter.
//
//		This method will perform Clean-Up operations on
//		the internal bufio.Writer object encapsulated in
//		this FileBufferReadWrite instance:
//
//			fBufReadWrite.writer
//
//		Upon completion, this method will effectively
//		render the FileBufferReadWrite instance,
//		'fBufReadWrite', invalid and unusable for any
//		future 'write' operations.
//
//	fBufReadWriteLabel					string
//
//		The name or label associated with input parameter
//		'fBufReadWrite' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fBufReadWrite" will
//		be automatically applied.
//
//	flushWriteBuffer					bool
//
//		If 'flushWriteBuffer' is set to 'true', this
//		method will flush the contents of the write
//		buffer. This means that write buffer contents are
//		guaranteed to be written to the internal
//		bufio.Writer object encapsulated by the
//		FileBufferReadWrite input parameter
//		'fBufReadWrite'.
//
//		If 'flushWriteBuffer' is set to 'false', the
//		contents of the write buffer will be lost.
//
//	releaseWriterMemResources			bool
//
//		If 'releaseWriterMemResources' is set to 'true',
//		this method	will release the internal memory
//		resources for the internal bufio.writer object
//		encapsulated by	'fBufReadWrite':
//
//			fBufReadWrite.writer
//
//		Releasing internal memory resources synchronizes
//		internal flags and prevents multiple calls to the
//		'close' method. Calling the 'close' method more
//		than once may produce unexpected results.
//
//	releaseFBuffReaderLocalMemRes		bool
//
//		If 'releaseFBuffReaderLocalMemRes' is set to
//		'true', this method will release the local memory
//		resources for the FileBufferReadWrite writer object
//		(fBufReadWrite):
//
//			fBufReadWrite.writer = nil
//			fBufReadWrite.writerFilePathName = ""
//
//		Releasing internal memory resources synchronizes
//		internal flags and prevents multiple calls to the
//		'close' method. Calling the 'close' method more
//		than once may produce unexpected results.
//
//	errPrefDto							*ePref.ErrPrefixDto
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
func (fBuffReadWriteElectron *fileBufferReadWriteElectron) writerFlushCloseRelease(
	fBufReadWrite *FileBufferReadWrite,
	fBufReadWriteLabel string,
	flushWriteBuffer bool,
	releaseWriterMemResources bool,
	releaseFBuffWriterLocalMemRes bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fBuffReadWriteElectron.lock == nil {
		fBuffReadWriteElectron.lock = new(sync.Mutex)
	}

	fBuffReadWriteElectron.lock.Lock()

	defer fBuffReadWriteElectron.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileBufferReadWriteElectron." +
		"writerFlushCloseRelease()"

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

	if fBufReadWrite.writer != nil {

		err2 = new(fileBufferWriterMolecule).
			flushCloseRelease(
				fBufReadWrite.writer,
				fBufReadWriteLabel+".writer",
				flushWriteBuffer,
				releaseWriterMemResources,
				ePrefix)

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Error occurred while closing the %v.\n"+
				"Error:\n%v\n",
				funcName,
				fBufReadWriteLabel+".writer",
				err2.Error())

			return err
		}

	}

	if releaseFBuffWriterLocalMemRes == true {

		fBufReadWrite.writer = nil
		fBufReadWrite.writerFilePathName = ""

	}

	return err
}
