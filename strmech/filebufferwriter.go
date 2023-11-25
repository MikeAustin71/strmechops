package strmech

import (
	"bufio"
	"errors"
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"math"
	"os"
	"sync"
)

// FileBufferWriter
//
// Type FileBufferWriter is a wrapper for 'bufio.Writer'.
// It is designed to write data to a destination
// io.Writer object using a buffer. As such,
// FileBufferWriter supports incremental or buffered
// 'write' operations to the target output destination.
//
// This structure and the associated methods facilitate
// data 'write' operations. The most common destination
// for these 'write' operations is assumed to be a data
// file residing on an attached storage drive. However,
// any object implementing the io.Writer interface may be
// used as a 'write' destination.
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
//	(1)	Use the methods 'New' and 'Setter' methods to
//		create and configure valid instances of
//		FileBufferWriter.
//
//	(2)	FileBufferWriter implements the following
//		interfaces:
//
//			io.Writer
//			io.Closer
//			io.ReadFrom
//			io.Seeker
//			io.WriteSeeker
//
// ----------------------------------------------------------------
//
// # Best Practice
//
//	(1)	Create a new, valid instance of FileBufferWriter
//		using one of the 'New' or 'Setter' methods. These
//		methods internally configure FileBufferWriter's
//		bufio.Writer object using a file or io.Writer
//		object.
//
//	(2)	After creating a valid instance of
//		FileBufferWriter, the user calls one of the
//		'Write' methods to write data to the internal
//		bufio.Writer object encapsulated by
//		FileBufferWriter. The 'Write' methods write data
//		to the target file or io.Writer object configured
//		by the 'New' or 'Setter' methods discussed above.
//
//	(3)	Upon completion of all 'write' operations, the
//		'Flush' and 'Close' tasks must be executed in
//		sequence to perform required clean-up tasks.
//
//		a.	The 'Flush' task can be performed by calling
//			the local method:
//
//				FileBufferWriter.Flush()
//
//		b.	The 'Close' and 'Flush' tasks can be
//			performed jointly by calling one local
//			method:
//
//				FileBufferWriter.Close()
//
//	(4)	Once method Close() is called, the current
//		FileBufferWriter instance becomes invalid
//		and unusable for future 'write' operations.
type FileBufferWriter struct {
	bufioWriter         *bufio.Writer
	ioWriter            *io.Writer
	filePtr             *os.File
	targetWriteFileName string

	lock *sync.Mutex
}

// Available
//
// This method returns the number of bytes that are
// unused in the write buffer.
//
// To acquire the 'size' of the buffer configured for the
// current instance of FileBufferWriter, call local
// method:
//
//	FileBufferWriter.Size()
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
//	int
//
//		This returned integer value specifies the number
//		of bytes that are unused in the 'write' buffer.
func (fBufWriter *FileBufferWriter) Available() int {

	if fBufWriter.lock == nil {
		fBufWriter.lock = new(sync.Mutex)
	}

	fBufWriter.lock.Lock()

	defer fBufWriter.lock.Unlock()

	if fBufWriter.bufioWriter == nil {

		return 0
	}

	return fBufWriter.bufioWriter.Available()
}

// AvailableBuffer
//
// This method returns an empty byte array buffer with
// FileBufferWriter.Available() capacity. This buffer is
// intended to be appended to and passed to an
// immediately succeeding Write call.
//
// The buffer is only valid until the next write
// operation on FileBufferWriter.
//
// If the current instance of FileBufferWriter is invalid
// or uninitialized, this method will return a zero
// length byte array.
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
//	[]byte
//
//		An empty byte array buffer with
//		FileBufferWriter.Available() capacity. This
//		buffer is intended to be appended to and
//		passed to an immediately succeeding Write
//		call.
//
//		This buffer is only valid until the next write
//		operation performed by FileBufferWriter.
//
//		If the current instance of FileBufferWriter is
//		invalid or uninitialized, this byte array will
//		be empty with a zero length.
func (fBufWriter *FileBufferWriter) AvailableBuffer() []byte {

	if fBufWriter.lock == nil {
		fBufWriter.lock = new(sync.Mutex)
	}

	fBufWriter.lock.Lock()

	defer fBufWriter.lock.Unlock()

	if fBufWriter.bufioWriter == nil {

		return make([]byte, 0)
	}

	return fBufWriter.bufioWriter.AvailableBuffer()

}

// Buffered
//
// This method returns the number of bytes that have been
// written into the 'write' buffer for the current
// instance of FileBufferWriter.
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
//	int
//
//		The number of bytes that have been written into
//		the 'write' buffer for the current instance of
//		FileBufferWriter.
func (fBufWriter *FileBufferWriter) Buffered() int {

	if fBufWriter.lock == nil {
		fBufWriter.lock = new(sync.Mutex)
	}

	fBufWriter.lock.Lock()

	defer fBufWriter.lock.Unlock()

	if fBufWriter.bufioWriter == nil {

		return 0
	}

	return fBufWriter.bufioWriter.Buffered()
}

// Close
//
// This method is provided in order to implement the
// io.Closer interface.
//
// FileBufferWriter.Close() effectively performs all
// required Clean-Up tasks. As such, this method should
// only be called after all 'write' operations have been
// completed and the services of the current
// FileBufferWriter instance are no longer required.
//
// After calling this method, FileBufferWriter.Close(),
// the current instance of FileBufferWriter will be
// invalid and unavailable for further 'write' operations.
//
// Calling this method will perform two elements of
// the clean-up operation required when all data
// has been written to the internal bufio.Writer and
// the services of the current FileBufferWriter
// instance are no longer required. The two clean-up
// procedures performed by this method are listed as
// follows:
//
//	(1)	'flush' procedure
//		This procedure flushes the 'write' buffer to
//		ensure that all data is written from the 'write'
//		buffer to the destination bufio.Writer object.
//
//	(2)	'close' procedure
//		This procedure properly closes  the underlying
//		bufio.Writer object configured for the current
//		instance of FileBufferWriter.
//
// The third procedure in the clean-up operation, namely
// the release of internal memory resources, will NOT be
// performed by this method.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	This method, FileBufferWriter.Close(), will
//		perform the following Clean-Up tasks:
//
//		(a)	'flush' procedure
//
//			This procedure flushes the 'write' buffer to
//			ensure that all data is written from the
//			'write' buffer to the destination bufio.Writer
//			object.
//
//		(b)	'close' procedure
//
//			This procedure properly closes the underlying
//			bufio.Writer object configured for the
//			current instance of FileBufferWriter.
//
//		(c)	Release internal memory resources
//
//			Releasing all internal memory resources will
//			synchronize internal flags and prevent
//			multiple calls to 'close' the underlying
//			io.Writer object. Calling 'close' on the same
//			underlying io.Writer object multiple times
//			can produce unexpected results.
//
//	(3)	Once this method completes all required Clean-Up
//		tasks, this current instance of FileBufferWriter
//		will become unavailable for further 'write'
//		operations.
//
//	(4)	If the user does not choose to flush the write
//		buffer before closing the FileBufferWriter
//		object, call the following local method to
//		close the FileBufferWriter instance and release
//		the internal memory resources:
//
//			FileBufferWriter.CloseAndRelease()
//
//	(5)	New instances of FileBufferWriter are created
//		using the 'New' methods. Existing instances of
//		FileBufferWriter may be reconfigured using the
//		'Setter' methods.
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
func (fBufWriter *FileBufferWriter) Close() error {

	if fBufWriter.lock == nil {
		fBufWriter.lock = new(sync.Mutex)
	}

	fBufWriter.lock.Lock()

	defer fBufWriter.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"FileBufferWriter."+
			"Close()",
		"")

	if err != nil {
		return err
	}

	return new(fileBufferWriterMolecule).
		flushCloseRelease(
			fBufWriter,
			"fBufWriter",
			true, // flushBuffer
			true, // releaseMemoryResources
			ePrefix)
}

// CloseAndRelease
//
// Performs the 'close' procedure and releases the
// internal memory resources for the current instance of
// FileBufferWriter. However, this method WILL NOT flush
// the write buffer
//
// Call this method when all 'write' operations have
// been completed and the services of the current
// FileBufferWriter instance are no longer required.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	This method will perform the following
//		two procedures:
//
//		(a)	'close' procedure
//
//			This procedure properly closes the underlying
//			bufio.Writer object configured for the current
//			instance of FileBufferWriter.
//
//		(b)	Release internal memory resources
//
//			This procedure releases all internal memory
//			resources and synchronizes internal flags
//			thereby preventing multiple calls to 'close'
//			the underlying io.writer object.
//
//	(2) This method, FileBufferWriter.CloseAndRelease(),
//		will NOT flush the internal 'write' buffer before
//		closing the FileBufferWriter. Any data remaining
//		in the internal 'write' buffer when this method
//		is called, will be lost. To 'flush', 'close' and
//		'release' memory resources call local method:
//
//			FileBufferWriter.Close()
//
//	(3)	Call this method only after completing all
//		'write' operations when the services of the
//		current FileBufferWriter instance are no longer
//		required.
//
//	(4)	Once this method executes the Clean-Up tasks,
//		this current instance of FileBufferWriter will
//		become invalid and unavailable for further
//		'write' operations.
//
//	(5)	New instances of FileBufferWriter are created
//		using the 'New' methods. Existing instances of
//		FileBufferWriter may be reconfigured using the
//		'Setter' methods.
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
//		appropriate error message.
func (fBufWriter *FileBufferWriter) CloseAndRelease(
	errorPrefix interface{}) error {

	if fBufWriter.lock == nil {
		fBufWriter.lock = new(sync.Mutex)
	}

	fBufWriter.lock.Lock()

	defer fBufWriter.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileBufferWriter."+
			"CloseAndRelease()",
		"")

	if err != nil {
		return err
	}

	return new(fileBufferWriterMolecule).
		flushCloseRelease(fBufWriter,
			"fBufWriter",
			false, // flushBuffer
			true,  // releaseMemoryResources
			ePrefix)
}

// Empty
//
// This method deletes all internal member variables and
// releases all the internal memory resources for the
// current instance of FileBufferWriter.
//
// Specifically the following internal object pointers
// are set to nil:
//
//	FileBufferWriter.bufioWriter = nil
//	FileBufferWriter.ioWriter = nil
//	FileBufferWriter.filePtr = nil
//
// In addition, the internal member variable
// 'targetWriteFileName' is set to an empty string.
//
//	FileBufferWriter.targetWriteFileName = ""
//
// After calling this method, the current instance of
// FileBufferReader will become invalid and unavailable
// for future 'read' operations.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	This method does NOT perform the 'close' protocol.
//	To perform both the 'close' protocol and release
//	all internal memory resources call one of the
//	following local methods:
//
//			FileBufferReader.Close()
//			FileBufferReader.CloseAndRelease()
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
func (fBufWriter *FileBufferWriter) Empty() {

	if fBufWriter.lock == nil {
		fBufWriter.lock = new(sync.Mutex)
	}

	fBufWriter.lock.Lock()

	defer fBufWriter.lock.Unlock()

	new(fileBufferWriterElectron).empty(
		fBufWriter)

}

// Flush
//
// Calling this method ensures that all remaining data in
// 'write' buffer will be written to the internal
// destination bufio.Writer object.
//
// This method does NOT close the internal destination
// bufio.Writer object. It only flushes the internal
// 'write' buffer.
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
func (fBufWriter *FileBufferWriter) Flush(
	errorPrefix interface{}) error {

	if fBufWriter.lock == nil {
		fBufWriter.lock = new(sync.Mutex)
	}

	fBufWriter.lock.Lock()

	defer fBufWriter.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileBufferWriter."+
			"Flush()",
		"")

	if err != nil {
		return err
	}

	err = new(fileBufferWriterAtom).
		flush(
			fBufWriter,
			"fBufWriter",
			ePrefix.XCpy("fBufWriter"))

	return err
}

// GetWriteBufferSize
//
// This method returns the size of the underlying 'write'
// buffer in bytes configured for the current instance of
// FileBufferWriter.
//
// To acquire the number of bytes unused in the buffer
// configured for the current instance of
// FileBufferWriter, call local method:
//
//	FileBufferWriter.Available()
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
//	int
//
//		This integer value returns the size of the
//		underlying 'write' buffer in bytes.
func (fBufWriter *FileBufferWriter) GetWriteBufferSize() int {

	if fBufWriter.lock == nil {
		fBufWriter.lock = new(sync.Mutex)
	}

	fBufWriter.lock.Lock()

	defer fBufWriter.lock.Unlock()

	if fBufWriter.bufioWriter == nil {

		return 0
	}

	return fBufWriter.bufioWriter.Size()
}

// IsClosed
//
// Returns a boolean value set to 'true' if the current
// instance of FileBufferWriter is 'closed'.
//
// The term 'closed' means that the FileBufferWriter
// member data values, including the internal
// bufio.Writer object, are invalid or uninitialized.
//
// If a FileBufferWriter object is closed, it is
// invalid and unavailable for future 'write' operations.
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
//	bool
//
//		If this boolean return value is set to 'true', it
//		signals that the current instance of
//		FileBufferWriter is 'closed' or uninitialized.
//
//		FileBufferWriter objects which are 'closed' are
//		invalid and unavailable for future 'write'
//		operations.
//
//		If this return parameter is set to 'false', it
//		signals that the current FileBufferWriter
//		instance is valid, properly configured and ready
//		to perform 'write' operations.
func (fBufWriter *FileBufferWriter) IsClosed() bool {

	if fBufWriter.lock == nil {
		fBufWriter.lock = new(sync.Mutex)
	}

	fBufWriter.lock.Lock()

	defer fBufWriter.lock.Unlock()

	if fBufWriter.bufioWriter != nil ||
		fBufWriter.ioWriter != nil {

		return false
	}

	return true
}

// NewIoWriter
//
// This method returns a fully initialized instance of
// FileBufferWriter.
//
// This returned instance of FileBufferWriter is created
// using an object implementing the io.Writer interface
// and passed as input parameter 'writer'.
//
// The size of the internal read buffer is controlled by
// input parameter 'bufSize'. The minimum buffer size is
// 16-bytes.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	Input parameter 'writer' will accept a pointer to an
//	instance of os.File because os.File implements the
//	io.Writer interface.
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
//	(1)	When all 'write' operations are completed and the
//		services of the returned new instance of
//		FileBufferWriter are no longer required, the user
//		MUST perform 'close' and Clean-Up tasks by
//		calling one of these local methods:
//
//			FileBufferWriter.Close()
//			FileBufferWriter.CloseAndRelease()
//
//	(2) If the input parameter io.writer base type is
//		NOT *os.File, the user will be required to
//		execute any 'close' or Clean-Up tasks
//		required by the external io.writer object in
//		addition to those 'close' and Clean-Up tasks
//		specified in paragraph (1), above.
//
//	(3)	Input parameter 'writer' will accept a pointer to
//		an instance of os.File because *os.File implements
//		the io.Writer interface.
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
//		tasks in addition to calling local method
//		FileBufferWriter.Close().
//
//		While the returned instance of FileBufferWriter
//		is primarily designed for writing data to disk
//		files, this 'writer' will in fact write data to
//		any object implementing the io.Writer interface.
//
//	bufSize						int
//
//		This integer value controls the size of the	write
//		buffer created for the returned instance of
//		FileBufferWriter.
//
//		'bufSize' should be configured to maximize
//		performance for 'write' operations subject to
//		prevailing memory limitations.
//
//		The minimum write buffer size is 1-byte. If
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
//	FileBufferWriter
//
//		If this method completes successfully, it will
//		return a fully configured instance of
//		FileBufferWriter.
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
func (fBufWriter *FileBufferWriter) NewIoWriter(
	writer io.Writer,
	bufSize int,
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

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileBufferWriter."+
			"NewIoWriter()",
		"")

	if err != nil {
		return newFileBufWriter, err
	}

	err = new(fileBufferWriterNanobot).
		setIoWriter(
			&newFileBufWriter,
			"newFileBufWriter",
			writer,
			"writer",
			bufSize,
			ePrefix.XCpy("newFileBufWriter"))

	return newFileBufWriter, err
}

// NewFileMgr
//
// Receives an instance of FileMgr as an input parameter
// 'fileMgr'.
//
// The target 'write' file identified by the 'fileMgr' is
// opened for either 'write-only' or 'read/write'
// operations depending on input parameter
// 'openFileReadWrite'.
//
// The target 'write' file identified by 'fileMgr' will
// be used to create a file pointer (*os.File) which in
// turn will be used to configure the internal
// bufio.Reader.
//
// The size of the internal 'write' buffer is controlled
// by input parameter 'bufSize'. If 'bufSize' is set to a
// value less than or equal to zero (0), it will be
// automatically reset to the default value of
// 4096-bytes.
//
// If the target path and file do not currently exist on
// an attached storage drive, this method will attempt to
// create them.
//
// Upon completion, this method returns a fully
// configured instance of FileBufferWriter.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	The returned type, 'FileBufferWriter', implements the
//	io.Writer and io.Closer interfaces.
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
//	(1)	As a precaution, the incoming 'fileMgr' object
//		will be closed before configuring the internal
//		bufio.writer object for the returned instance of
//		FileBufferWriter.
//
//	(2)	When all read operations have been completed and
//		there is no further need for the returned
//		instance of FileBufferWriter, the user is
//		responsible for 'closing' and releasing the
//		associated memory resources by calling one
//		of the following local methods:
//
//			FileBufferWriter.Close()
//			FileBufferWriter.CloseAndRelease()
//
//	(2)	After executing the 'close' operation described
//		in paragraph (2) above, the current instance of
//		FileBufferWriter will be rendered invalid and
//		unavailable for future 'write' operations.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fileMgr						*FileMgr
//
//		A pointer to an instance of FileMgr. The file
//		identified by 'fileMgr' will be used as an output
//		destination for 'write' operations performed by
//		method:
//
//			FileBufferWriter.Write()
//
//		If the path and file name encapsulated by
//		'fileMgr' do not currently exist on an attached
//		storage drive, this method will attempt to create
//		them.
//
//	openFileReadWrite			bool
//
//		If this parameter is set to 'true', the target
//		'write' file created from input parameter
//		'fileMgr' will be opened for 'read' and
//		'write' operations.
//
//		If 'openFileReadWrite' is set to 'false', the
//		target write file will be opened for 'write-only'
//		operations.
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
//		The minimum write buffer size is 1-byte. If
//		'bufSize' is set to a size less than or equal to
//		zero, it will be automatically reset to the
//		default buffer size of 4096-bytes.
//
//	truncateExistingFile			bool
//
//		If this parameter is set to 'true', the target
//		'write' file identified by 'fileMgr' will be
//		opened for 'write' operations. If the target file
//		previously existed, it will be truncated. This
//		means that the file's previous contents will be
//		deleted.
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
//	fileInfoPlus				FileInfoPlus
//
//		This returned instance of Type FileInfoPlus
//		contains data elements describing the file
//		identified by input parameter 'fileMgr'.
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
//	newFileBufWriter			FileBufferWriter
//
//		If this method completes successfully, a fully
//		configured instance of FileBufferWriter will
//		be returned.
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
func (fBufWriter *FileBufferWriter) NewFileMgr(
	fileMgr *FileMgr,
	openFileReadWrite bool,
	bufSize int,
	truncateExistingFile bool,
	errorPrefix interface{}) (
	fInfoPlus FileInfoPlus,
	newFileBufWriter FileBufferWriter,
	err error) {

	if fBufWriter.lock == nil {
		fBufWriter.lock = new(sync.Mutex)
	}

	fBufWriter.lock.Lock()

	defer fBufWriter.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileBufferWriter."+
			"NewFileMgr()",
		"")

	if err != nil {

		return fInfoPlus, newFileBufWriter, err
	}

	fInfoPlus,
		err = new(fileBufferWriterMicrobot).
		setFileMgr(
			&newFileBufWriter,
			"newFileBufWriter",
			fileMgr,
			"fileMgr",
			openFileReadWrite,
			bufSize,
			truncateExistingFile,
			ePrefix.XCpy("fileMgr"))

	return fInfoPlus, newFileBufWriter, err
}

// NewPathFileName
//
// Receives a path and file name as an input parameter.
// This target 'write' file is opened for either
// 'write-only' or 'read/write' operations depending on
// input parameter 'openFileReadWrite'.
//
// Upon completion, this method returns a fully
// configured instance of FileBufferWriter.
//
// If the target path and file do not currently exist on
// an attached storage drive, this method will attempt to
// create them.
//
// The size of the internal 'write' buffer is controlled
// by input parameter 'bufSize'. If 'bufSize' is set to a
// value less than or equal to zero (0), it will be
// automatically reset to the default value of
// 4096-bytes.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	The returned type, 'FileBufferWriter', implements the
//	io.Writer and io.Closer interfaces.
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
//	(1)	When all 'write' operations have been completed
//		and there is no further need for the returned
//		instance of FileBufferWriter, the user is
//		responsible for 'closing' and releasing the
//		associated memory resources by calling one
//		of the following local methods:
//
//			FileBufferWriter.Close()
//			FileBufferWriter.CloseAndRelease()
//
//	(2)	After executing the 'close' operation described
//		in paragraph (1) above, the current instance of
//		FileBufferWriter will be rendered invalid and
//		unavailable for future 'write' operations.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathFileName				string
//
//		This string contains the path and file name of
//		the target 'write' file which will be used as a
//		destination data file for 'write' operations
//		performed by method:
//
//			FileBufferWriter.Write()
//
//		If the target path and file name do not currently
//		exist on an attached storage drive, this method
//		will attempt to create them.
//
//	openFileReadWrite			bool
//
//		If this parameter is set to 'true', the target
//		'write' file created from input parameter
//		'pathFileName' will be opened for 'read' and
//		'write' operations.
//
//		If 'openFileReadWrite' is set to 'false', the
//		target write file will be opened for 'write-only'
//		operations.
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
//		The minimum write buffer size is 1-byte. If
//		'bufSize' is set to a size less than or equal to
//		zero, it will be automatically reset to the
//		default buffer size of 4096-bytes.
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
//	fileInfoPlus				FileInfoPlus
//
//		This returned instance of Type FileInfoPlus
//		contains data elements describing the file
//		identified by input parameter 'pathFileName'.
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
//	newFileBufWriter			FileBufferWriter
//
//		If this method completes successfully, a fully
//		configured instance of FileBufferWriter will
//		be returned.
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
func (fBufWriter *FileBufferWriter) NewPathFileName(
	pathFileName string,
	openFileReadWrite bool,
	bufSize int,
	truncateExistingFile bool,
	errorPrefix interface{}) (
	fInfoPlus FileInfoPlus,
	newFileBufWriter FileBufferWriter,
	err error) {

	if fBufWriter.lock == nil {
		fBufWriter.lock = new(sync.Mutex)
	}

	fBufWriter.lock.Lock()

	defer fBufWriter.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileBufferWriter."+
			"NewPathFileName()",
		"")

	if err != nil {

		return fInfoPlus, newFileBufWriter, err
	}

	fInfoPlus,
		err = new(fileBufferWriterNanobot).
		setPathFileName(
			&newFileBufWriter,
			"newFileBufWriter",
			pathFileName,
			"pathFileName",
			openFileReadWrite,
			bufSize,
			truncateExistingFile,
			ePrefix.XCpy("pathFileName"))

	return fInfoPlus, newFileBufWriter, err
}

// ReadFrom
//
// Implements the io.ReadFrom interface.
//
// This method will read data from the io.Reader object
// passed as input parameter 'reader' and write that data
// to bufio.Writer encapsulated by the current instance of
// FileBufferWriter.
//
// The data is read from 'reader' using an internal byte
// array equal in length to the buffer configured for the
// current instance of FileBufferWriter.
//
// The return parameter 'numOfBytesProcessed' records the
// number of bytes read from 'reader' and written to
// the FileBufferWriter bufio.Writer object. If the
// number of bytes read fails to match the number of
// bytes written, an error will be returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	reader						io.Reader
//
//		An object which implements the io.Reader interface.
//		This method will read the entire contents of this
//		io.Reader object and write the data to the
//		bufio.Writer object encapsulated by the current
//		instance of FileBufferWriter.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	numOfBytesProcessed			int64
//
//		This return parameter documents the number of
//		bytes read from 'reader' and written to the
//		FileBufferWriter bufio.Writer object. If the
//		number of bytes read fails to match the number
//		bytes written, an error will be returned.
//
//	err							error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message.
func (fBufWriter FileBufferWriter) ReadFrom(
	reader io.Reader) (
	numOfBytesProcessed int64,
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
			"ReadFrom()",
		"")

	if err != nil {

		return numOfBytesProcessed, err
	}

	if fBufWriter.bufioWriter == nil {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of 'FileBufferWriter' is invalid!\n"+
			"The internal bufio.Writer object has NOT been initialized.\n"+
			"Call one of the 'New' or 'Setter' methods to create a\n"+
			"valid instance of 'FileBufferWriter'\n",
			ePrefix.String())

		return numOfBytesProcessed, err
	}

	if reader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'reader' is invalid!\n"+
			"'reader' has a 'nil' value.\n",
			ePrefix.String())

		return numOfBytesProcessed, err
	}

	bufSize := fBufWriter.bufioWriter.Size()

	if bufSize <= 0 {

		err = fmt.Errorf("%v\n"+
			"Error: The Buffer Size for currnt instance of\n"+
			"FileBufferWriter is less than or equal to zero!\n"+
			"Therefore the current instance of FileBufferWriter\n"+
			"is invalid. Use one of the 'New' or 'Setter' methods\n"+
			"to create a valid instance of FileBufferWriter.\n"+
			"The current FileBufferWriter Size is '%v'\n",
			ePrefix.String(),
			bufSize)

		return numOfBytesProcessed, err
	}

	var bytesRead = make([]byte, bufSize)
	var numBytesRead, numBytesWritten int
	var err1, err2 error
	var maxCycle = math.MaxInt - 1
	var cycleCnt int

	for {

		cycleCnt++

		if cycleCnt >= maxCycle {

			err = fmt.Errorf("%v\n"+
				"Error: Infinite Loop!\n"+
				"The 'Read' operation failed to locate io.EOF\n"+
				"otherwise known as the end-of-file for this\n"+
				"underlying io.Reader object.\n"+
				"Read Cycle Count= %v\n",
				ePrefix.String(),
				cycleCnt)

			break
		}

		numBytesRead,
			err1 = reader.Read(bytesRead)

		if err1 != nil &&
			err1 != io.EOF {

			err = fmt.Errorf("%v\n"+
				"Error: reader.Read(bytesRead)\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				err1.Error())

			break

		}

		if numBytesRead > 0 {

			numBytesWritten,
				err2 = fBufWriter.bufioWriter.Write(
				bytesRead[:numBytesRead])

			if err2 != nil {

				err = fmt.Errorf("%v\n"+
					"Error returned by fBufWriter.bufioWriter.Write(bytesRead[:numBytesRead])\n"+
					"numBytesRead= '%v'\n"+
					"Error=\n%v\n",
					ePrefix.String(),
					numBytesRead,
					err2.Error())

				break
			}

			if numBytesWritten != numBytesRead {

				err = fmt.Errorf("%v\n"+
					"Error: Number of bytes read does NOT\n"+
					"match the number of bytes written.\n"+
					"Write Cycle Number: %v\n"+
					"   Number of Bytes Read: %v\n"+
					"Number of Bytes Written: %v\n",
					ePrefix.String(),
					cycleCnt,
					numBytesRead,
					numBytesWritten)

				break
			}

			numOfBytesProcessed += int64(numBytesWritten)
		}

		if err1 == io.EOF {

			break
		}

		clear(bytesRead)
	}

	return numOfBytesProcessed, err
}

// Seek
//
// This method sets the offset for the next 'write'
// operation within the 'write' file. This method only
// succeeds if the current FileBufferWriter instance
// was created as a file with a path and file name string
// or a File Manager object (FileMgr).
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
// Seek returns the new offset relative to the start of the
// file or an error, if any.
//
// Seek implements the 'io.Seeker' interface.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	The standard technique is to call method Seek()
//		and a 'write' method in tandem. In this case,
//		BE ADVISED, best practice specifies calling
//		Flush() to flush the write buffer BEFORE
//		calling this method, Seek().
//
//	(2)	Calling Seek() to move to an offset before the
//		beginning of the file is an error.
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
func (fBufWriter FileBufferWriter) Seek(
	targetOffset int64,
	whence int) (
	offsetFromFileStart int64,
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
		"FileBufferReader."+
			"WriteTo()",
		"")

	if err != nil {

		return offsetFromFileStart, err
	}

	if fBufWriter.bufioWriter == nil {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of 'FileBufferWriter' is invalid!\n"+
			"The internal bufio.Writer object has NOT been initialized.\n"+
			"Call one of the 'New' or 'Setter' methods to create a\n"+
			"valid instance of 'FileBufferWriter'\n",
			ePrefix.String())

		return offsetFromFileStart, err
	}

	if fBufWriter.ioWriter == nil {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of 'FileBufferWriter' is invalid!\n"+
			"The internal io.Writer object has NOT been initialized.\n"+
			"Call one of the 'New' or 'Setter' methods to create a\n"+
			"valid instance of 'FileBufferWriter'\n",
			ePrefix.String())

		return offsetFromFileStart, err
	}

	var ok bool
	var seekerObj io.Seeker
	var localWriter io.Writer
	localWriter = *fBufWriter.ioWriter

	seekerObj, ok = localWriter.(io.Seeker)

	if !ok {

		err = fmt.Errorf("%v\n"+
			"Error: This Seek method was invoked on a\n"+
			"'FileBufferWriter' internal io.Writer object\n"+
			"which does NOT support the io.Seeker interface.\n"+
			"This means:\n"+
			"(1) The 'Seek' method is unavailable.\n"+
			"\n"+
			"(2) The 'FileBufferWriter' internal io.Writer\n"+
			"      object was created from something\n"+
			"      other than a disk file (*os.File).\n",
			ePrefix.String())

		return offsetFromFileStart, err

	}

	var whenceCodeIsOk bool
	var whenceCodeStr string

	whenceCodeIsOk,
		whenceCodeStr = new(FileConstants).
		GetSeekerWhenceCodes(whence)

	if !whenceCodeIsOk {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'whence' is invalid!\n"+
			"'whence' MUST be equal to one of the following\n"+
			"constant values:\n"+
			"  io.SeekStart = 0\n"+
			"  io.SeekCurrent = 1\n"+
			"  io.SeekEnd = 2\n"+
			"Input 'whence' value = %v\n",
			ePrefix.String(),
			whence)

		return offsetFromFileStart, err
	}

	var err2 error

	offsetFromFileStart,
		err2 = seekerObj.Seek(
		targetOffset,
		whence)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: FileBufferWriter.Seek()\n"+
			"targetOffSet = %v\n"+
			"whence = %v\n"+
			"Error = \n%v\n",
			ePrefix.String(),
			targetOffset,
			whenceCodeStr,
			err2.Error())
	}

	return offsetFromFileStart, err
}

// SetFileMgr
//
// This method will completely re-initialize the current
// instance of FileBufferWriter using the path and file
// name passed as input parameter 'fileMgr'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the current instance of
//	FileBufferWriter.
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
// # Input Parameters
//
//	fileMgr						*FileMgr
//
//		A pointer to an instance of FileMgr. The file
//		identified by 'fileMgr' will be used as a
//		destination for 'write' operations performed by
//		method:
//
//			FileBufferWriter.Write()
//
//		If the path and file name encapsulated by
//		'fileMgr' do not currently exist on an attached
//		storage drive, this method will attempt to create
//		them.
//
//	openFileReadWrite			bool
//
//		If this parameter is set to 'true', the target
//		'write' file created from input parameter
//		'fileMgr' will be opened for 'read' and
//		'write' operations.
//
//		If 'openFileReadWrite' is set to 'false', the
//		target write file will be opened for 'write-only'
//		operations.
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
//		The minimum write buffer size is 1-byte. If
//		'bufSize' is set to a size less than or equal to
//		zero, it will be automatically reset to the
//		default buffer size of 4096-bytes.
//
//	truncateExistingFile			bool
//
//		If this parameter is set to 'true', the target
//		'write' file specified by 'fileMgr' will be
//		opened for write operations. If the target file
//		previously existed, it will be truncated. This
//		means that the file's previous contents will be
//		deleted.
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
//	fileInfoPlus				FileInfoPlus
//
//		This returned instance of Type FileInfoPlus
//		contains data elements describing the file
//		identified by input parameter 'fileMgr'.
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
func (fBufWriter *FileBufferWriter) SetFileMgr(
	fileMgr *FileMgr,
	openFileReadWrite bool,
	bufSize int,
	truncateExistingFile bool,
	errorPrefix interface{}) (
	fInfoPlus FileInfoPlus,
	err error) {

	if fBufWriter.lock == nil {
		fBufWriter.lock = new(sync.Mutex)
	}

	fBufWriter.lock.Lock()

	defer fBufWriter.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileBufferWriter."+
			"SetFileMgr()",
		"")

	if err != nil {

		return fInfoPlus, err
	}

	fInfoPlus,
		err = new(fileBufferWriterMicrobot).
		setFileMgr(
			fBufWriter,
			"fBufWriter",
			fileMgr,
			"fileMgr",
			openFileReadWrite,
			bufSize,
			truncateExistingFile,
			ePrefix.XCpy("fileMgr"))

	return fInfoPlus, err
}

// SetPathFileName
//
// This method will completely re-initialize the current
// instance of FileBufferWriter using the path and file
// name passed as input parameter 'pathFileName'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the current instance of
//	FileBufferWriter.
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
// # Input Parameters
//
//	pathFileName				string
//
//		This string contains the path and file name of
//		the target 'write' file which will be used as a
//		destination data file for 'write' operations
//		performed by method:
//
//			FileBufferWriter.Write()
//
//		If the target path and file name do not currently
//		exist on an attached storage drive, this method
//		will attempt to create them.
//
//	openFileReadWrite			bool
//
//		If this parameter is set to 'true', the target
//		'write' file created from input parameter
//		'pathFileName' will be opened for 'read' and
//		'write' operations.
//
//		If 'openFileReadWrite' is set to 'false', the
//		target write file will be opened for 'write-only'
//		operations.
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
//		The minimum write buffer size is 1-byte. If
//		'bufSize' is set to a size less than or equal to
//		zero, it will be automatically reset to the
//		default buffer size of 4096-bytes.
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
//	fileInfoPlus				FileInfoPlus
//
//		This returned instance of Type FileInfoPlus
//		contains data elements describing the file
//		identified by input parameter 'pathFileName'.
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
func (fBufWriter *FileBufferWriter) SetPathFileName(
	pathFileName string,
	openFileReadWrite bool,
	bufSize int,
	truncateExistingFile bool,
	errorPrefix interface{}) (
	fInfoPlus FileInfoPlus,
	err error) {

	if fBufWriter.lock == nil {
		fBufWriter.lock = new(sync.Mutex)
	}

	fBufWriter.lock.Lock()

	defer fBufWriter.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileBufferWriter."+
			"SetPathFileName()",
		"")

	if err != nil {

		return fInfoPlus, err
	}

	fInfoPlus,
		err = new(fileBufferWriterNanobot).
		setPathFileName(
			fBufWriter,
			"fBufWriter",
			pathFileName,
			"pathFileName",
			openFileReadWrite,
			bufSize,
			truncateExistingFile,
			ePrefix.XCpy("pathFileName"))

	return fInfoPlus, err
}

// SetIoWriter
//
// This method will completely re-initialize the current
// instance of FileBufferWriter using the io.Writer object
// passed as input parameter 'writer'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the current instance of
//	FileBufferWriter.
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
//		operations in addition to calling local method
//		FileBufferWriter.Close().
//
//		While the returned instance of FileBufferWriter
//		is primarily designed for writing data to disk
//		files, this 'writer' will in fact write data to
//		any object implementing the io.Writer interface.
//
//	bufSize						int
//
//		This integer value controls the size of the	write
//		buffer created for the returned instance of
//		FileBufferWriter.
//
//		'bufSize' should be configured to maximize
//		performance for 'write' operations subject to
//		prevailing memory limitations.
//
//		The minimum write buffer size is 1-byte. If
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
func (fBufWriter *FileBufferWriter) SetIoWriter(
	writer io.Writer,
	bufSize int,
	errorPrefix interface{}) error {

	if fBufWriter.lock == nil {
		fBufWriter.lock = new(sync.Mutex)
	}

	fBufWriter.lock.Lock()

	defer fBufWriter.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileBufferWriter."+
			"SetIoWriter()",
		"")

	if err != nil {
		return err
	}

	err = new(fileBufferWriterNanobot).
		setIoWriter(
			fBufWriter,
			"fBufWriter",
			writer,
			"writer",
			bufSize,
			ePrefix.XCpy("fBufWriter"))

	return err
}

// Write
//
// Writes the contents of the byte array input paramter
// ('bytesToWrite') to the internal destination
// bufio.Writer object previously configured for this
// instance of FileBufferWriter.
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
//			FileBufferWriter.Close()
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
//		to the internal destination bufio.Writer object
//		previously configured for the current instance of
//		FileBufferWriter.
//
//		Typically, the internal destination bufio.Writer
//		object will be a data file existing on an attached
//		storage drive. However, the destination
//		bufio.Writer object may be any object implementing
//		the io.Writer interface.
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
//		written to the internal destination bufio.Writer
//		object configured for the current instance of
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
func (fBufWriter FileBufferWriter) Write(
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

	if fBufWriter.bufioWriter == nil {

		err = fmt.Errorf("%v\n"+
			"-------------------------------------------------------\n"+
			"Error: This instance of 'FileBufferWriter' is invalid!\n"+
			"The internal bufio.Writer has NOT been initialized.\n"+
			"Call one of the 'New' or 'Setter' methods when creating\n"+
			"a new valid instance of 'FileBufferWriter'\n",
			ePrefix.String())

		return numBytesWritten, err
	}

	if len(bytesToWrite) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'bytesToWrite' is invalid!\n"+
			"The 'bytesToWrite' byte array is empty. It has zero bytes.\n",
			ePrefix.String())

		return numBytesWritten, err
	}

	var err2 error

	numBytesWritten,
		err2 = fBufWriter.bufioWriter.Write(bytesToWrite)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by fBufWriter.bufioWriter.Write(bytesToWrite).\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			err2.Error())
	}

	return numBytesWritten, err
}

// WriteTextOrNumbers
//
// This method will accept many different text or numeric
// data types which are then converted to a byte or
// string array and written to the internal bufio.Writer
// object encapsulated in the current instance of
// FileBufferWriter.
//
// The text or numeric data type passed as input
// parameter 'charsToWrite' must match one of over sixty
// eligible data types.
//
// If 'charsToWrite' is set to an ineligible data type,
// an error will be returned.
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
//	autoFlushAndCloseOnExit		bool
//
//		When this parameter is set to 'true' and no
//		processing errors are encountered during method
//		execution, this method will automatically perform
//		the following clean-up tasks upon exit:
//
//		(1)	The write buffer will be flushed thereby
//			ensuring that all remaining data in the
//			'write' buffer will be written to the
//			underlying bufio.Writer object.
//
//		(2)	The internal bufio.Writer object will be
//			properly closed and there will be no need
//			to make a separate call to local method,
//			FileBufferWriter.Close().
//
//		(3) After performing these clean-up tasks, the
//			current instance of FileBufferWriter will
//			invalid and unusable for future 'write'
//			operations.
//
//		If input parameter 'autoFlushAndCloseOnExit' is
//		set to 'false', this method will still
//		automatically flush the 'write' buffer. However,
//		it will NOT close the internal bufio.Writer
//		object. This means that all data remaining in the
//		'write' buffer will be written to the underlying
//		bufio.Writer output destination. But, most
//		importantly, when 'autoFlushAndCloseOnExit' is
//		set to false, the user is then responsible for
//		performing the 'Close' operation by calling the
//		local method:
//
//			FileBufferWriter.Close()
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
//		instance of FileBufferWriter.
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
func (fBufWriter *FileBufferWriter) WriteTextOrNumbers(
	charsToWrite interface{},
	writeEndOfLineChars string,
	writeEndOfTextChars string,
	autoFlushAndCloseOnExit bool,
	errorPrefix interface{}) (
	numOfBytesWritten int64,
	err error) {

	if fBufWriter.lock == nil {
		fBufWriter.lock = new(sync.Mutex)
	}

	fBufWriter.lock.Lock()

	defer fBufWriter.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "FileBufferWriter." +
		"WriteTextOrNumbers()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {

		return numOfBytesWritten, err
	}

	if fBufWriter.bufioWriter == nil {

		err = fmt.Errorf("%v\n"+
			"-------------------------------------------------------\n"+
			"Error: This instance of 'FileBufferWriter' is invalid!\n"+
			"The internal bufio.Writer has NOT been initialized.\n"+
			"Call one of the 'New' or 'Setter' methods when creating\n"+
			"a new valid instance of 'FileBufferWriter'\n",
			ePrefix.String())

		return numOfBytesWritten, err
	}

	var writeBytesFunc = fBufWriter.lowLevelWriteBytes

	numOfBytesWritten,
		err = new(fileWriterHelperMicrobot).
		writeCharacters(
			writeBytesFunc,
			charsToWrite,
			"charsToWrite",
			writeEndOfLineChars,
			writeEndOfTextChars,
			ePrefix.XCpy("fIoWriter.ioWriter<-charsToWrite"))

	if err != nil {

		return numOfBytesWritten, err
	}

	var fBufWriterAtom = new(fileBufferWriterAtom)
	var err2, err3 error

	err2 = fBufWriterAtom.flush(
		fBufWriter,
		"fBufWriter",
		ePrefix.XCpy("Flush fBufWriter"))

	if err2 != nil {

		err3 = fmt.Errorf("%v\n"+
			"Error returned by fBufWriterMolecule.flush(fBufWriter)\n"+
			"Error= \n%v\n",
			funcName,
			err2.Error())

		err = errors.Join(err, err3)
	}

	if autoFlushAndCloseOnExit == true {

		err2 = fBufWriterAtom.close(
			fBufWriter,
			"fBufWriter",
			ePrefix.XCpy("Close fBufWriter"))

		if err2 != nil {

			err3 = fmt.Errorf("%v\n"+
				"Error returned by fBufWriterMolecule.close(fBufWriter)\n"+
				"Error= \n%v\n",
				funcName,
				err2.Error())

			err = errors.Join(err, err3)
		}

	}

	return numOfBytesWritten,
		err
}

func (fBufWriter *FileBufferWriter) lowLevelWriteBytes(
	bytesToWrite []byte,
	writeEndOfTextChars string,
	errPrefDto *ePref.ErrPrefixDto) (
	numOfBytesWritten int64,
	err error) {

	var ePrefix *ePref.ErrPrefixDto

	funcName := "FileBufferWriter." +
		"lowLevelWriteBytes()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		return numOfBytesWritten, err
	}

	if fBufWriter.bufioWriter == nil {

		err = fmt.Errorf("%v\n"+
			"-------------------------------------------------------\n"+
			"Error: This instance of 'FileBufferWriter' is invalid!\n"+
			"The internal bufio.Writer has NOT been initialized.\n"+
			"Call one of the 'New' or 'Setter' methods when creating\n"+
			"a new valid instance of 'FileBufferWriter'\n",
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
		err2 = fBufWriter.bufioWriter.Write(bytesToWrite)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by fBufWriter.bufioWriter.Write(bytesToWrite).\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			err2.Error())

		return numOfBytesWritten, err
	}

	numOfBytesWritten +=
		int64(localNumOfBytesWritten)

	lenWriteEndOfTextChars := len(writeEndOfTextChars)

	if lenWriteEndOfTextChars > 0 {

		expectedNumOfBytesToWrite += int64(lenWriteEndOfTextChars)

		localNumOfBytesWritten,
			err2 = fBufWriter.bufioWriter.Write(
			[]byte(writeEndOfTextChars))

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Error returned by fBufWriter.bufioWriter.Write(writeEndOfTextChars).\n"+
				"writeEndOfTextChars= '%v'\n"+
				"Error=\n%v\n",
				ePrefix.String(),
				writeEndOfTextChars,
				err2.Error())

			return numOfBytesWritten, err
		}

	}

	return numOfBytesWritten, err
}

type fileBufferWriterMicrobot struct {
	lock *sync.Mutex
}

// setFileMgr
//
// This 'setter' method is used to initialize new values
// for internal member variables contained in the
// instance of FileBufferWriter passed as input parameter
// 'fBufWriter'.
//
// The new bufio.Writer object assigned to 'fBufWriter' is
// generated from the File Manager (FileMgr) instance
// passed as input parameter 'fileMgr'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the instance of
//	FileBufferWriter passed as input parameter
//	'fBufWriter'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fBufWriter					*FileBufferWriter
//
//		A pointer to an instance of FileBufferWriter.
//
//		All internal member variable data values in
//		this instance will be deleted and initialized
//		to new values based on the following input
//		parameters.
//
//	fBufWriterLabel				string
//
//		The name or label associated with input parameter
//		'fBufWriter' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fBufWriter" will be
//		automatically applied.
//
//	fileMgr						*FileMgr
//
//		A pointer to an instance of FileMgr. The file
//		identified by 'fileMgr' will be used as an output
//		destination for 'write' operations performed by
//		method:
//
//			FileBufferWriter.Write()
//
//		If the path and file name encapsulated by
//		'fileMgr' do not currently exist on an attached
//		storage drive, this method will attempt to create
//		them.
//
//	fileMgrLabel				string
//
//		The name or label associated with input parameter
//		'fileMgr' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fileMgr" will be
//		automatically applied.
//
//	openFileReadWrite			bool
//
//		If this parameter is set to 'true', the target
//		'write' file created from input parameter
//		'fileMgr' will be opened for 'read' and 'write'
//		operations.
//
//		If 'openFileReadWrite' is set to 'false', the
//		target write file will be opened for 'write-only'
//		operations.
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
//		The minimum write buffer size is 1-byte. If
//		'bufSize' is set to a size less than or equal to
//		zero, it will be automatically reset to the
//		default buffer size of 4096-bytes.
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
//		identified by input parameter 'fileMgr'.
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
func (fBufWriterMicrobot *fileBufferWriterMicrobot) setFileMgr(
	fBufWriter *FileBufferWriter,
	fBufWriterLabel string,
	fileMgr *FileMgr,
	fileMgrLabel string,
	openFileReadWrite bool,
	bufSize int,
	truncateExistingFile bool,
	errPrefDto *ePref.ErrPrefixDto) (
	fInfoPlus FileInfoPlus,
	err error) {

	if fBufWriterMicrobot.lock == nil {
		fBufWriterMicrobot.lock = new(sync.Mutex)
	}

	fBufWriterMicrobot.lock.Lock()

	defer fBufWriterMicrobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileBufferWriterMicrobot." +
		"setFileMgr()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		return fInfoPlus, err
	}

	if len(fBufWriterLabel) == 0 {

		fBufWriterLabel = "fBufWriter"
	}

	if len(fileMgrLabel) == 0 {

		fileMgrLabel = "fileMgr"
	}

	if fBufWriter == nil {

		err = fmt.Errorf("%v\n"+
			"Error: The FileBufferWriter instance passed\n"+
			"as input parameter '%v' is invalid!\n"+
			"'%v' is a 'nil' pointer.\n",
			ePrefix.String(),
			fBufWriterLabel,
			fBufWriterLabel)

		return fInfoPlus, err
	}

	if fileMgr == nil {

		err = fmt.Errorf("%v\n"+
			"Error: The FileMgr instance passed\n"+
			"as input parameter '%v' is invalid!\n"+
			"'%v' is a 'nil' pointer.\n",
			ePrefix.String(),
			fileMgrLabel,
			fileMgrLabel)

		return fInfoPlus, err
	}

	var err2 error

	err = new(fileMgrHelperAtom).isFileMgrValid(
		fileMgr,
		ePrefix.XCpy(fileMgrLabel))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter %v is invalid.\n"+
			"%v failed the validity test.\n"+
			"Error=\n%v\n",
			funcName,
			fileMgrLabel,
			fileMgrLabel,
			err2.Error())

		return fInfoPlus, err
	}

	err2 = new(fileMgrHelper).closeFile(
		fileMgr,
		ePrefix.XCpy(fileMgrLabel))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"An error occurred while closing the file pointer.\n"+
			"for FileMgr input parameter '%v'.\n"+
			"Error=\n%v\n",
			funcName,
			fileMgrLabel,
			err2.Error())

		return fInfoPlus, err
	}

	fileMgrLabel += ".absolutePathFileName"

	fInfoPlus,
		err = new(fileBufferWriterNanobot).
		setPathFileName(
			fBufWriter,
			fBufWriterLabel,
			fileMgr.absolutePathFileName,
			fileMgrLabel,
			openFileReadWrite,
			bufSize,
			truncateExistingFile,
			ePrefix.XCpy(fBufWriterLabel+"<-"+fileMgrLabel))

	return fInfoPlus, err
}

type fileBufferWriterNanobot struct {
	lock *sync.Mutex
}

// setIoWriter
//
// This 'setter' method is used to initialize new values
// for internal member variables contained in the
// instance of FileBufferWriter passed as input parameter
// 'fBufWriter'. The new configuration will be based on
// an io.Writer object passed as input parameter
// 'writer'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the instance of
//	FileBufferWriter passed as input parameter
//	'fBufWriter'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fBufWriter					*FileBufferWriter
//
//		A pointer to an instance of FileBufferWriter.
//
//		All internal member variable data values in
//		this instance will be deleted and initialized
//		to new values based on the following input
//		parameters.
//
//	fBufWriterLabel				string
//
//		The name or label associated with input parameter
//		'fBufWriter' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fBufWriter" will be
//		automatically applied.
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
//		file data to files residing on an attached
//		storage drive. However, with this configuration,
//		the user is responsible for manually closing the
//		file and performing any other required clean-up
//		operations in addition to calling local method
//		FileBufferWriter.Close().
//
//		While the configured instance of FileBufferWriter
//		(fBufWriter) is primarily designed for writing
//		data to disk files, this 'writer' will in fact
//		write data to any object implementing the
//		io.Writer interface.
//
//		This instance of io.Writer will be used to
//		configure the internal bufio.Writer contained in
//		'fBufWriter' and used to conduct 'write'
//		operations.
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
//	bufSize						int
//
//		This integer value controls the size of the
//		'write' buffer configured for the
//		FileBufferWriter instance passed as input
//		parameter 'fBufWriter'.
//
//		'bufSize' should be configured to maximize
//		performance for 'write' operations subject to
//		prevailing memory limitations.
//
//		The minimum write buffer size is 1-byte. If
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
func (fBufWriterNanobot *fileBufferWriterNanobot) setIoWriter(
	fBufWriter *FileBufferWriter,
	fBufWriterLabel string,
	writer io.Writer,
	writerLabel string,
	bufSize int,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fBufWriterNanobot.lock == nil {
		fBufWriterNanobot.lock = new(sync.Mutex)
	}

	fBufWriterNanobot.lock.Lock()

	defer fBufWriterNanobot.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileBufferWriterNanobot." +
		"setIoWriter()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return err
	}

	if len(fBufWriterLabel) == 0 {

		fBufWriterLabel = "fBufWriter"
	}

	if len(writerLabel) == 0 {

		writerLabel = "writer"
	}

	if fBufWriter == nil {

		err = fmt.Errorf("%v\n"+
			"Error: The FileBufferWriter instance passed\n"+
			"as input parameter '%v' is invalid!\n"+
			"'%v' is a 'nil' pointer.\n",
			ePrefix.String(),
			fBufWriterLabel,
			fBufWriterLabel)

		return err
	}

	if writer == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is invalid!\n"+
			"'%v' is a 'nil' value.\n",
			ePrefix.String(),
			writerLabel,
			writerLabel)

		return err

	}

	var fBufWriterMolecule = new(fileBufferWriterMolecule)

	// Flush and close the old fBufWriter
	var err2 error

	err2 = fBufWriterMolecule.
		flushCloseRelease(fBufWriter,
			"fBufWriter",
			true, // flushBuffer
			true, // releaseMemoryResources
			ePrefix)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"A error occurred while closing the current\n"+
			"FileBufferWriter instance in preparation for\n"+
			"setting a new io.Writer object.\n"+
			"Error: \n%v\n",
			funcName,
			err2.Error())

		return err
	}

	if bufSize <= 0 {

		bufSize = 4096
	}

	var ok bool

	fBufWriter.filePtr, ok = writer.(*os.File)

	if ok == true {

		fBufWriter.ioWriter = &writer

		fBufWriter.bufioWriter = bufio.NewWriterSize(
			fBufWriter.filePtr,
			bufSize)

		fBufWriter.targetWriteFileName =
			fBufWriter.filePtr.Name()

	} else {
		// ok == false - this is NOT a file

		fBufWriter.filePtr = nil

		fBufWriter.ioWriter = &writer

		fBufWriter.bufioWriter = bufio.NewWriterSize(
			writer,
			bufSize)

	}

	return err
}

// setPathFileName
//
// This 'setter' method is used to initialize new values
// for internal member variables contained in the
// instance of FileBufferWriter passed as input parameter
// 'fBufWriter'.
//
// The new bufio.Writer object assigned to 'fBufWriter' is
// generated from the file name passed as input parameter
// 'pathFileName'
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the instance of
//	FileBufferWriter passed as input parameter
//	'fBufWriter'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fBufWriter					*FileBufferWriter
//
//		A pointer to an instance of FileBufferWriter.
//
//		All internal member variable data values in
//		this instance will be deleted and initialized
//		to new values based on the following input
//		parameters.
//
//	fBufWriterLabel				string
//
//		The name or label associated with input parameter
//		'fBufWriter' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fBufWriter" will be
//		automatically applied.
//
//	pathFileName				string
//
//		This string contains the path and file name of
//		the target 'write' file which will be used as a
//		destination data source for 'write' operations
//		performed by method:
//
//			FileBufferWriter.Write()
//
//		If the target path and file name do not currently
//		exist on an attached storage drive, this method
//		will attempt to create them.
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
//		'write' file created from input parameter
//		'pathFileName' will be opened for 'read' and
//		'write' operations.
//
//		If 'openFileReadWrite' is set to 'false', the
//		target write file will be opened for 'write-only'
//		operations.
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
//		The minimum write buffer size is 1-byte. If
//		'bufSize' is set to a size less than or equal to
//		zero, it will be automatically set to the default
//		buffer size of 4096-bytes.
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
//		identified by input parameter 'pathFileName'.
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
func (fBufWriterNanobot *fileBufferWriterNanobot) setPathFileName(
	fBufWriter *FileBufferWriter,
	fBufWriterLabel string,
	pathFileName string,
	pathFileNameLabel string,
	openFileReadWrite bool,
	bufSize int,
	truncateExistingFile bool,
	errPrefDto *ePref.ErrPrefixDto) (
	fInfoPlus FileInfoPlus,
	err error) {

	if fBufWriterNanobot.lock == nil {
		fBufWriterNanobot.lock = new(sync.Mutex)
	}

	fBufWriterNanobot.lock.Lock()

	defer fBufWriterNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileBufferWriterNanobot." +
		"setPathFileName()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		return fInfoPlus, err
	}

	if len(fBufWriterLabel) == 0 {

		fBufWriterLabel = "fBufWriter"
	}

	if len(pathFileNameLabel) == 0 {

		pathFileNameLabel = "pathFileName"
	}

	if fBufWriter == nil {

		err = fmt.Errorf("%v\n"+
			"Error: The FileBufferWriter instance passed\n"+
			"as input parameter '%v' is invalid!\n"+
			"'%v' is a 'nil' pointer.\n",
			ePrefix.String(),
			fBufWriterLabel,
			fBufWriterLabel)

		return fInfoPlus, err
	}

	if len(pathFileName) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is invalid!\n"+
			"'%v' is an empty string with a length of zero (0).\n",
			ePrefix.String(),
			pathFileNameLabel,
			pathFileNameLabel)

		return fInfoPlus, err

	}

	var fBufWriterMolecule = new(fileBufferWriterMolecule)

	// Flush and close the old fBufWriter
	var err2 error

	err2 = fBufWriterMolecule.
		flushCloseRelease(fBufWriter,
			"fBufWriter",
			true, // flushBuffer
			true, // releaseMemoryResources
			ePrefix)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"A error occurred while closing the current\n"+
			"FileBufferWriter instance in preparation for\n"+
			"setting a new Path and File Name object.\n"+
			"Error: \n%v\n",
			funcName,
			err2.Error())

		return fInfoPlus, err
	}

	if bufSize <= 0 {

		bufSize = 4096
	}

	// var pathFileDoesExist bool

	pathFileName,
		_,
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
			"of 'pathFileName' on an attached storage drive.\n"+
			"%v = '%v'\n"+
			"Error= \n%v\n",
			funcName,
			pathFileNameLabel,
			pathFileName,
			err2.Error())

		return fInfoPlus, err
	}

	var filePermissionCfg FilePermissionConfig

	var filePermissionStr = "--w--w--w-"

	if openFileReadWrite == true {

		filePermissionStr = "-rw-rw-rw-"
	}

	filePermissionCfg,
		err = new(FilePermissionConfig).New(
		filePermissionStr,
		ePrefix.XCpy("filePermissionCfg<-"))

	if err != nil {

		return fInfoPlus, err
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

			return fInfoPlus, err
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

			return fInfoPlus, err
		}

	}

	fBufWriter.filePtr,
		err = new(fileHelperBoson).
		openFile(
			pathFileName,
			false,
			fileOpenCfg,
			filePermissionCfg,
			"pathFileName",
			ePrefix)

	if err != nil {

		if fBufWriter.filePtr != nil {
			_ = fBufWriter.filePtr.Close()
		}

		return fInfoPlus, err
	}

	var writer io.Writer

	writer = fBufWriter.filePtr

	fBufWriter.ioWriter = &writer

	fBufWriter.targetWriteFileName = pathFileName

	fBufWriter.bufioWriter = bufio.NewWriterSize(
		fBufWriter.filePtr,
		bufSize)

	return fInfoPlus, err
}

type fileBufferWriterMolecule struct {
	lock *sync.Mutex
}

// flushCloseRelease
//
// This method is used to perform all necessary clean-up
// operations after final data has been written to the
// internal destination bufio.Writer object.
//
// These clean-up operations consist of the following:
//
//	(1)	Flushing the write buffer to ensure that all
//		data is written to the internal destination
//		bufio.Writer object. This procedure is optional
//		but highly recommended. If the write buffer is
//		NOT flushed, any data contained in the write
//		buffer will be lost. Reference input parameter,
//		'flushBuffer'.
//
//	(2) Closing the internal bufio.Writer object thereby
//		rendering it invalid and unavailable for any future
//		'write' operations.
//
//	(3)	Releasing all internal memory resources for the
//		passed instance of FileBufferWriter
//		('fBufWriter'). Releasing internal memory
//		resources synchronizes internal flags and
//		prevents multiple calls to the 'close' method.
//
// Individual Clean-up procedures may be performed
// depending on the settings for input parameters
// 'flushWriteBuffer' and 'releaseMemoryResources'.
//
// After calling this method, the passed instance of
// FileBufferWriter ('fBufWriter') will be unusable and
// should be discarded.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fBufWriter					*FileBufferWriter
//
//		A pointer to an instance of FileBufferWriter.
//
//		All internal member variable data values in
//		this instance will be deleted and initialized
//		to new values based on the following input
//		parameters.
//
//	fBufWriterLabel				string
//
//		The name or label associated with input parameter
//		'fBufWriter' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fBufWriter" will be
//		automatically applied.
//
//	flushWriteBuffer			bool
//
//		If this parameter is set to 'true', this method
//		will flush the contents of the write buffer. This
//		means that write buffer contents are guaranteed
//		to be written to the internal bufio.Writer object
//		encapsulated by FileBufferWriter input parameter
//		'fBufWriter'.
//
//		If 'flushWriteBuffer' is set to 'false', the
//		contents of the write buffer will be lost.
//
//	releaseMemoryResources		bool
//
//		If this parameter is set to 'true', this method
//		will release all internal memory resources for
//		the passed instance of FileBufferWriter
//		('fBufWriter'). Releasing internal memory
//		resources synchronizes internal flags and
//		prevents multiple calls to the 'close' method.
//		Calling the 'close' method more	than once may
//		produce unexpected results.
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
func (fBufWriterMolecule *fileBufferWriterMolecule) flushCloseRelease(
	fBufWriter *FileBufferWriter,
	fBufWriterLabel string,
	flushWriteBuffer bool,
	releaseMemoryResources bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fBufWriterMolecule.lock == nil {
		fBufWriterMolecule.lock = new(sync.Mutex)
	}

	fBufWriterMolecule.lock.Lock()

	defer fBufWriterMolecule.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileBufferWriterMolecule." +
		"flushCloseRelease()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return err
	}

	if len(fBufWriterLabel) == 0 {

		fBufWriterLabel = "fBufWriter"
	}

	if fBufWriter == nil {

		err = fmt.Errorf("%v\n"+
			"Error: The FileBufferWriter instance passed\n"+
			"as input parameter '%v' is invalid!\n"+
			"'%v' is a 'nil' pointer.\n",
			ePrefix.String(),
			fBufWriterLabel,
			fBufWriterLabel)

		return err
	}

	var fBuffWriterAtom = new(fileBufferWriterAtom)

	if flushWriteBuffer == true {

		err = fBuffWriterAtom.
			flush(
				fBufWriter,
				fBufWriterLabel,
				ePrefix)

		if err != nil {
			return err
		}

	}

	err = fBuffWriterAtom.close(
		fBufWriter,
		fBufWriterLabel,
		ePrefix)

	if err != nil {
		return err
	}

	if releaseMemoryResources == true {

		new(fileBufferWriterElectron).
			empty(
				fBufWriter)

	}

	return err
}

type fileBufferWriterAtom struct {
	lock *sync.Mutex
}

// close
//
// This method is designed to perform clean up
// operations after completion of all 'write'
// operations.
//
// All internal member variable data values for the
// instance of FileBufferWriter passed as input parameter
// 'fBufWriter' will be deleted and reset to their zero
// states.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	This method will delete all pre-existing data
//		values in the instance of FileBufferWriter passed
//		as input parameter 'fBufWriter'.
//
//		After completion of this method this
//		FileBufferWriter instance will be unusable,
//		invalid and unavailable for future 'write'
//		operations.
//
//	(2)	This 'close' method will NOT flush the 'write'
//		buffer. To flush the 'write' buffer call:
//			fileBufferWriterAtom.flush()
//
//		Be sure to call the 'flush()' method before you
//		call this method ('close()').
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fBufWriter					*FileBufferWriter
//
//		A pointer to an instance of FileBufferWriter.
//
//		All internal member variable data values in
//		this instance will be deleted.
//
//		If a file pointer (*os.File) was previously
//		configured for 'fBufWriter', it will be closed
//		and set to 'nil' by this method.
//
//	fBufWriterLabel				string
//
//		The name or label associated with input parameter
//		'fBufWriter' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fBufWriter" will be
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
func (fBufWriterAtom *fileBufferWriterAtom) close(
	fBufWriter *FileBufferWriter,
	fBufWriterLabel string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fBufWriterAtom.lock == nil {
		fBufWriterAtom.lock = new(sync.Mutex)
	}

	fBufWriterAtom.lock.Lock()

	defer fBufWriterAtom.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileBufferWriterAtom." +
		"close()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return err
	}

	if len(fBufWriterLabel) == 0 {

		fBufWriterLabel = "fBufWriter"
	}

	if fBufWriter == nil {

		err = fmt.Errorf("%v\n"+
			"-------------------------------------------\n"+
			"Error: The FileBufferWriter instance passed\n"+
			"as input parameter '%v' is invalid!\n"+
			"'%v' is a 'nil' pointer.\n",
			ePrefix.String(),
			fBufWriterLabel,
			fBufWriterLabel)

		return err
	}

	if fBufWriter.ioWriter != nil {

		var ok bool
		var closerObj io.Closer
		var localWriter io.Writer

		localWriter = *fBufWriter.ioWriter

		closerObj, ok = localWriter.(io.Closer)

		if ok {

			var err2 error

			err2 = closerObj.Close()

			if err2 != nil {

				errText := fmt.Sprintf(
					"%v\n"+
						"Error returned while closing the 'fBufWriter'\n"+
						"internal io.Writer object.\n",
					ePrefix.String())

				if len(fBufWriter.targetWriteFileName) > 0 {

					errText += fmt.Sprintf(
						"Target Write File Name: %v\n",
						fBufWriter.targetWriteFileName)

				}

				err = fmt.Errorf("%v"+
					"closerObj.Close() Error=\n%v\n",
					errText,
					err2.Error())

			}

		}

	}

	return err
}

// flush
//
// This method performs one function. Namely, it flushes
// all data from the write file effectively ensuring that
// all data in the buffer is written to the file or
// underlying device defined by the internal bufio.Writer
// encapsulated in the FileBufferWriter instance passed
// as input parameter 'fBufWriter'.
//
// Specifically, this method does NOT 'close' the
// 'fBufWriter' FileBufferWriter instance. As such, this
// method does not delete member variable data contained
// in 'fBufWriter'. To fully close the 'fBufWriter'
// instance, make a separate call to local method:
//
//	fileBufferWriterAtom.close()'.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	This method will NOT modify the internal member
//	variable data values for the instance of
//	FileBufferWriter passed as input parameter
//	'fBufWriter'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fBufWriter					*FileBufferWriter
//
//		A pointer to an instance of FileBufferWriter.
//		Any data remaining in the 'write' buffer will
//		be written to the underlying data file by the
//		flush 'operation' performed by this method.
//
//	fBufWriterLabel				string
//
//		The name or label associated with input parameter
//		'fBufWriter' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fBufWriter" will be
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
func (fBufWriterAtom *fileBufferWriterAtom) flush(
	fBufWriter *FileBufferWriter,
	fBufWriterLabel string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fBufWriterAtom.lock == nil {
		fBufWriterAtom.lock = new(sync.Mutex)
	}

	fBufWriterAtom.lock.Lock()

	defer fBufWriterAtom.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileBufferWriterAtom." +
		"flush()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return err
	}

	if len(fBufWriterLabel) == 0 {

		fBufWriterLabel = "fBufWriter"
	}

	if fBufWriter == nil {

		err = fmt.Errorf("%v\n"+
			"Error: The FileBufferWriter instance passed\n"+
			"as input parameter '%v' is invalid!\n"+
			"'%v' is a 'nil' pointer.\n",
			ePrefix.String(),
			fBufWriterLabel,
			fBufWriterLabel)

		return err
	}

	if fBufWriter.bufioWriter == nil {

		// Nothing to do. Nothing to close.
		return err

	}

	var err2 error

	err2 = fBufWriter.bufioWriter.Flush()

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned while flushing the 'write' buffer!\n"+
			"%v.bufioWriter.Flush()\n"+
			"Error = \n%v\n",
			ePrefix.String(),
			fBufWriterLabel,
			err2.Error())

	}

	return err
}

type fileBufferWriterElectron struct {
	lock *sync.Mutex
}

// empty
//
// This method deletes all internal member variables and
// releases all the internal memory resources for an
// instance of FileBufferWriter passed as input parameter
// 'fBufWriter'.
//
// Specifically the following internal member variables
// are set to 'nil' or their initial zero values:
//
//	FileBufferWriter.targetWriteFileName = ""
//	FileBufferWriter.filePtr = nil
//	FileBufferWriter.ioWriter = nil
//	FileBufferWriter.bufioWriter = nil
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fBufWriter					*FileBufferWriter
//
//		A pointer to an instance of FileBufferWriter.
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
func (fBuffWriterElectron *fileBufferWriterElectron) empty(
	fBufWriter *FileBufferWriter) {

	if fBuffWriterElectron.lock == nil {
		fBuffWriterElectron.lock = new(sync.Mutex)
	}

	fBuffWriterElectron.lock.Lock()

	defer fBuffWriterElectron.lock.Unlock()

	if fBufWriter == nil {

		return
	}

	fBufWriter.targetWriteFileName = ""

	fBufWriter.filePtr = nil

	fBufWriter.ioWriter = nil

	fBufWriter.bufioWriter = nil

	return
}
