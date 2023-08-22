package strmech

import (
	"bufio"
	"errors"
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"os"
	"strings"
	"sync"
)

// FileBufferWriter
//
// Type FileBufferWriter is a wrapper for bufio.Writer.
// It is designed to write data to a destination
// io.Writer object using a buffer.
//
// This structure and the associated methods facilitate
// data 'write' operations. The most common destination
// for these 'write' operations is assumed to be a data
// file residing on an attached storage drive. However,
// any object implementing the io.Writer interface may be
// used as a 'write' destination.
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
//	(1)	Use the methods 'New' and 'Setter' methods to
//		configure valid instances of FileBufferWriter.
//
//	(2)	FileBufferWriter implements the following
//		interfaces:
//
//			io.Writer
//			io.Closer
//
// ----------------------------------------------------------------
//
// # Best Practice
//
//	(1)	Create a new instance of FileBufferWriter using
//		either the New() method or the NewPathFileName()
//		method.
//
//		(a)	The New() method is used when an instance of
//			io.Writer is created externally by the user
//			and passed to the FileBufferWriter.New()
//			method.
//
//			Under this scenario, the user is independently
//			responsible for clean-up of the io.Writer
//			object after FileBufferWriter 'write'
//			operations have been completed.
//
//			Once all FileBufferWriter 'write' operations
//			have been completed, call methods Flush() and
//			Close() to perform local FileBufferWriter
//			clean-up tasks.
//
//		(b)	The NewPathFileName() method allows for the
//			creation of an internal file pointer to a
//			file passed as a path and file name by the
//			user. This file serves as the internal target
//			bufio.Writer object to which data will be
//			written.
//
//			Under this scenario, the user simply calls
//			methods Flush() and Close() in sequence to
//			perform all required clean-up tasks after
//			'write' operations have been completed.
//
//			Once method Close() is called, the current
//			FileBufferWriter instance becomes unusable
//			and should be discarded.
//
//	(2)	After creating an instance of FileBufferWriter,
//		the user calls the Write() method to write bytes
//		of data to the internal target bufio.Writer object.
//		This 'write' target may be a file or any other
//		object which implements the io.Writer interface.
//
//		The Write() method should be called repeatedly
//		until all data has been written to the underlying
//		bufio.Writer object.
//
//		Upon completion of the 'write' operation, call
//		methods Flush() and Close() in sequence to
//		perform required clean-up tasks.
//
//	(3)	After all data bytes have been written to the
//		internal bufio.Writer object, the user must call
//		methods Flush() and Close() to perform necessary
//		clean-up operations.
//
//		Once method Close() is called, the current
//		FileBufferWriter instance becomes unusable
//		and should be discarded.
type FileBufferWriter struct {
	fileWriter          *bufio.Writer
	filePtr             *os.File
	targetWriteFileName string

	lock *sync.Mutex
}

// Available
//
// This method returns the number of bytes that are
// unused in the write buffer.
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

	if fBufWriter.fileWriter == nil {

		return 0
	}

	return fBufWriter.fileWriter.Available()
}

// Close
//
// This method is used to close any open file pointers
// and perform necessary clean-up operations after all
// data has been written to the internal destination
// bufio.Writer object.
//
// These clean-up operations include flushing the
// write buffer to ensure that all data is written to
// the destination bufio.Writer object.
//
// After calling this method, the current instance of
// FileBufferWriter will be unusable and should be
// discarded.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	(1)	This method implements the io.Closer interface.
//
//	(2) This method is identical in function to method,
//		FileBufferWriter.FlushAndClose().
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	Call this method after completing all write
//		operations. Calling this method is essential to
//		performance of necessary clean-up tasks after
//		completion of all 'write' operations. Clean-up
//		tasks consist of:
//
//		(a)	Flushing the 'write' buffer to ensure that
//			all data is written from the 'write' buffer
//			to the underlying bufio.Writer object.
//
//		(b)	Properly closing the 'write' file or
//			bufio.Writer object.
//
//	(2)	Once this method completes the 'Close' operation,
//		this instance of FileBufferWriter becomes
//		invalid, unusable and unavailable for further
//		'write' operations.
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

	err = new(fileBufferWriterNanobot).
		flushAndClose(
			fBufWriter,
			"fBufWriter",
			ePrefix.XCpy("fBufWriter"))

	return err
}

// FlushAndClose
//
// This method is used to perform all necessary clean-up
// operations after all data has been written to the
// internal destination bufio.Writer object.
//
// These clean-up operations consist of:
//
//	(1)	Flushing the write buffer to ensure that all
//		data is written to the internal destination
//		bufio.Writer object.
//
//				AND
//
//	(2) Closing the internal bufio.Writer object thereby
//		renders it invalid and unavailable for any future
//		'write' operations.
//
// After calling this method, the current instance of
// FileBufferWriter will be unusable and should be
// discarded.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	Call this method after completing all write
//		operations. Calling this method is essential to
//		performance of necessary clean-up tasks after
//		completion of all 'write' operations. Clean-up
//		tasks consist of:
//
//		(a)	Flushing the 'write' buffer to ensure that
//			all data is written from the 'write' buffer
//			to the underlying bufio.Writer object.
//
//		(b)	Properly closing the 'write' file or
//			internal bufio.Writer object.
//
//	(2)	Once this method completes the 'Close' operation,
//		this instance of FileBufferWriter becomes invalid
//		and unavailable for further	'write' operations.
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
func (fBufWriter *FileBufferWriter) FlushAndClose(
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
			"Close()",
		"")

	if err != nil {
		return err
	}

	err = new(fileBufferWriterNanobot).
		flushAndClose(
			fBufWriter,
			"fBufWriter",
			ePrefix.XCpy("fBufWriter"))

	return err
}

// Flush
//
// Calling this method ensures that all remaining data in
// 'write' buffer will be written to the internal
// destination bufio.Writer object.
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

	err = new(fileBufferWriterMolecule).
		flush(
			fBufWriter,
			"fBufWriter",
			ePrefix.XCpy("fBufWriter"))

	return err
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
//		FileBufferWriter.FlushAndClose().
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
//		FileBufferWriter.FlushAndClose().
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
// If for any reason, the returned number of bytes
// written ('numBytesWritten') to the internal
// destination bufio.Writer object is less than the length
// of the byte array passed as input parameter
// 'bytesToWrite', an error containing an explanation for
// this difference will be returned.
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
//			FileBufferWriter.FlushAndClose()
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
			"-------------------------------------------------------\n"+
			"Error: This instance of 'FileBufferWriter' is invalid!\n"+
			"The internal bufio.Writer has NOT been initialized.\n"+
			"Call one of the 'New' or 'Setter' methods when creating\n"+
			"a new valid instance of 'FileBufferWriter'\n",
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

// WriteTextLines
//
// Writes all the string elements of a string array to
// the internal bufio.Writer object encapsulated in the
// current instance of FileBufferWriter.
//
// When writing final text lines to the internal
// bufio.Writer object, the line termination or
// end-of-line characters appended to each text line will
// be specified by input parameter 'writeEndOfLineChars'.
//
// If input parameter 'autoFlushAndCloseOnExit' is set to
// 'true', this method will automatically perform all
// required clean-up tasks upon completion. Clean-up
// tasks involve flushing the bufio.Writer object,
// closing the bufio.Writer object and then deleting
// the bufio.Writer structure values internal to
// the current FileBufferWriter instance. When these
// clean-up tasks are completed, the current
// FileBufferWriter instance will be invalid and
// unavailable for future 'write' operations.
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
//	FileBufferWriter.Close()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	outputLinesArray			*StringArrayDto
//
//		This instance of StringArrayDto encapsulates a
//		string array. All the string elements of this
//		array will be written to the underlying
//		bufio.Writer object contained in the current
//		instance of FileBufferWriter.
//
//		If the encapsulated string array is empty and
//		contains zero string array elements, an error is
//		returned.
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
//	autoFlushAndCloseOnExit		bool
//
//		When this parameter is set to 'true', this
//		method will automatically perform the following
//		clean-up tasks upon exit, even if processing
//		errors are encountered:
//
//		(1)	The write buffer will be flushed thereby
//			ensuring that all remaining data in the
//			'write' buffer will be written to the
//			underlying bufio.Writer object.
//
//		(2)	The internal bufio.Writer object will be
//			properly closed and there will be need
//			to make a separate call to local method,
//			FileBufferReadWrite.Close().
//
//		(3) After performing these clean-up tasks, the
//			current instance of FileBufferReadWrite will
//			invalid and unusable for future 'write'
//			operations.
//
//		If input parameter 'autoFlushAndCloseOnExit' is
//		set to 'false', this method will automatically
//		flush the 'write' buffer, but it will NOT close
//		internal bufio.Writer object. This means that all
//		data remaining in the 'write' buffer will be
//		written to the underlying bufio.Writer output
//		destination. However, most importantly, the user
//	 	is then responsible for performing the 'Close'
//		operation by calling the local method:
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
//	numBytesWritten				int
//
//		Returns the number of bytes written to the
//		bufio.Writer object encapsulated by the current
//		instance of FileBufferWriter.
//
//	numOfLinesWritten			int
//
//		Returns the number to text lines extracted from
//		the 'outputLinesArray' string array and written
//		to the bufio.Writer object encapsulated by the
//		current instance of FileBufferWriter.
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
func (fBufWriter *FileBufferWriter) WriteTextLines(
	outputLinesArray *StringArrayDto,
	writeEndOfLineChars string,
	autoFlushAndCloseOnExit bool,
	errorPrefix interface{}) (
	numBytesWritten int,
	numOfLinesWritten int,
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
			"WriteTextLines()",
		"")

	if err != nil {

		return numBytesWritten, numOfLinesWritten, err
	}

	if outputLinesArray == nil {

		err = fmt.Errorf("%v\n"+
			"-----------------------------------------------------\n"+
			"Error: Input parameter 'outputLinesArray' is invalid!\n"+
			"'outputLinesArray' is a 'nil' pointer.\n",
			ePrefix.String())

		return numBytesWritten, numOfLinesWritten, err
	}

	if fBufWriter.fileWriter == nil {

		err = fmt.Errorf("%v\n"+
			"-------------------------------------------------------\n"+
			"Error: This instance of 'FileBufferWriter' is invalid!\n"+
			"The internal bufio.Writer has NOT been initialized.\n"+
			"Call one of the 'New' or 'Setter' methods when creating\n"+
			"a new valid instance of 'FileBufferWriter'\n",
			ePrefix.String())

		return numBytesWritten, numOfLinesWritten, err
	}

	lenStrArray := len(outputLinesArray.StrArray)

	if lenStrArray == 0 {

		err = fmt.Errorf("%v\n"+
			"---------------------------------------------------------\n"+
			"Error: Input parameter 'outputLinesArray' is invalid!\n"+
			"The 'outputLinesArray' string array is empty!\n"+
			"'outputLinesArray.StrArray' is a zero length string array.\n",
			ePrefix.String())

		return numBytesWritten, numOfLinesWritten, err
	}

	var str string
	var localNumBytesWritten int
	var err2 error

	for i := 0; i < lenStrArray; i++ {

		str = outputLinesArray.StrArray[i] +
			writeEndOfLineChars

		if len(str) == 0 {
			continue
		}

		localNumBytesWritten,
			err2 = fBufWriter.fileWriter.Write([]byte(str))

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Error returned by fBufWriter.fileWriter.Write(bytesToWrite).\n"+
				"String Array Index= %v\n"+
				"Error=\n%v\n",
				ePrefix.String(),
				i,
				err2.Error())

			break
		}

		numBytesWritten += localNumBytesWritten
		numOfLinesWritten++
	}

	var cleanUpStatus string

	if err != nil {

		cleanUpStatus = "After Processing Error"

	} else {

		cleanUpStatus = "Success"
	}

	if autoFlushAndCloseOnExit == true {

		err2 = new(fileBufferWriterNanobot).flushAndClose(
			fBufWriter,
			"fBufWriter",
			ePrefix.XCpy(fmt.Sprintf(
				"%v Flush/Close-Readers & Writers",
				cleanUpStatus)))

		if err2 != nil {

			err = errors.Join(err, err2)
		}

	} else {

		err2 = new(fileBufferWriterMolecule).
			flush(
				fBufWriter,
				"fBufWriter",
				ePrefix.XCpy(fmt.Sprintf(
					"%v Flush/Close-Readers & Writers",
					cleanUpStatus)))

		if err2 != nil {

			err = errors.Join(err, err2)
		}

	}

	return numBytesWritten, numOfLinesWritten, err
}

// WriteStrBuilder
//
// Receives an instance of strings.Builder and proceeds
// to write all the string data contained therein to the
// previously configured internal bufio.writer
// destination configured for the current instance of
// FileBufferWriter.
//
// If input parameter 'autoFlushAndCloseOnExit' is set to
// 'true', this method will automatically perform all
// required clean-up tasks upon completion. Clean-up
// tasks involve flushing the bufio.Writer object,
// closing the bufio.Writer object and then deleting
// the bufio.Writer structure values internal to
// the current FileBufferWriter instance. When these
// clean-up tasks are completed, the current
// FileBufferWriter instance will be invalid and
// unavailable for future 'write' operations.
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
//	FileBufferWriter.Close()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	strBuilder					*strings.Builder
//
//		A pointer to an instance of strings.Builder. The
//		entire string contents of 'strBuilder' will be
//		written to the internal bufio.Writer destination
//		encapsulated by the current instance of
//		FileBufferWriter.
//
//		If the string contained in 'strBuilder' is empty,
//		or a zero length string, an error will be
//		returned.
//
//	autoFlushAndCloseOnExit		bool
//
//		When this parameter is set to 'true', this
//		method will automatically perform the following
//		clean-up tasks upon exit, even if processing
//		errors are encountered:
//
//		(1)	The write buffer will be flushed thereby
//			ensuring that all remaining data in the
//			'write' buffer will be written to the
//			underlying bufio.Writer object.
//
//		(2)	The internal bufio.Writer object will be
//			properly closed and there will be need
//			to make a separate call to local method,
//			FileBufferReadWrite.Close().
//
//		(3) After performing these clean-up tasks, the
//			current instance of FileBufferReadWrite will
//			invalid and unusable for future 'write'
//			operations.
//
//		If input parameter 'autoFlushAndCloseOnExit' is
//		set to 'false', this method will automatically
//		flush the 'write' buffer, but it will NOT close
//		internal bufio.Writer object. This means that all
//		data remaining in the 'write' buffer will be
//		written to the underlying bufio.Writer output
//		destination. However, most importantly, the user
//		is then responsible for performing the 'Close'
//		operation by calling the local method:
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
//	numBytesWritten				int
//
//		Returns the number of bytes extracted from
//		input parameter 'strBuilder' and written to the
//		internal bufio.Writer object encapsulated by the
//		current instance of FileBufferWriter.
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
func (fBufWriter *FileBufferWriter) WriteStrBuilder(
	strBuilder *strings.Builder,
	autoFlushAndCloseOnExit bool,
	errorPrefix interface{}) (
	numBytesWritten int64,
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
			"WriteStrBuilder()",
		"")

	if err != nil {

		return numBytesWritten, err
	}

	if fBufWriter.fileWriter == nil {

		err = fmt.Errorf("%v\n"+
			"-------------------------------------------------------\n"+
			"Error: This instance of 'FileBufferWriter' is invalid!\n"+
			"The internal bufio.Writer has NOT been initialized.\n"+
			"Call one of the 'New' or 'Setter' methods when creating\n"+
			"a new valid instance of 'FileBufferWriter'\n",
			ePrefix.String())

		return numBytesWritten, err
	}

	if strBuilder == nil {

		err = fmt.Errorf("%v\n"+
			"-------------------------------------------------------\n"+
			"Error: Input parameter 'strBuilder' is invalid!\n"+
			"'strBuilder' is a 'nil' pointer.\n",
			ePrefix.String())

		return numBytesWritten, err
	}

	if strBuilder.Len() == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'strBuilder' is invalid!\n"+
			"'strBuilder' contains an empty or zero length string.\n",
			ePrefix.String())

		return numBytesWritten, err
	}

	var err2 error
	var localNumBytesWritten int

	localNumBytesWritten,
		err2 = fBufWriter.fileWriter.Write(
		[]byte(strBuilder.String()))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by fBufWriter.fileWriter.Write(strBuilder).\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			err2.Error())

	} else {

		numBytesWritten = int64(localNumBytesWritten)
	}

	var cleanUpStatus string

	if err != nil {

		cleanUpStatus = "After Processing Error"

	} else {

		cleanUpStatus = "Success"
	}

	if autoFlushAndCloseOnExit == true {

		err2 = new(fileBufferWriterNanobot).flushAndClose(
			fBufWriter,
			"fBufWriter",
			ePrefix.XCpy(fmt.Sprintf(
				"%v Flush/Close-Readers & Writers",
				cleanUpStatus)))

		if err2 != nil {

			err = errors.Join(err, err2)
		}

	} else {

		err2 = new(fileBufferWriterMolecule).
			flush(
				fBufWriter,
				"fBufWriter",
				ePrefix.XCpy(fmt.Sprintf(
					"%v Flush/Close-Readers & Writers",
					cleanUpStatus)))

		if err2 != nil {

			err = errors.Join(err, err2)
		}

	}

	return numBytesWritten, err
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

// flushAndClose
//
// This method is used to perform all necessary clean-up
// operations after all data has been written to the
// internal destination bufio.Writer object.
//
// These clean-up operations consist of:
//
//	(1)	Flushing the write buffer to ensure that all
//		data is written to the internal destination
//		bufio.Writer object.
//
//				AND
//
//	(2) Closing the internal bufio.Writer object thereby
//		it invalid and unavailable for any future
//		'write' operations.
//
// After calling this method, the current instance of
// FileBufferWriter will be unusable and should be
// discarded.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	Call this method after completing all write
//		operations. Calling this method is essential to
//		performance of necessary clean-up tasks. Clean-up
//		tasks consist of:
//
//		(a)	Flushing the 'write' buffer to ensure that
//			all data is written from the 'write' buffer
//			to the underlying bufio.Writer object.
//
//		(b)	Properly closing the 'write' file or
//			bufio.Writer object.
//
//	(2)	Once this method completes the 'Close' operation,
//		this instance of FileBufferWriter becomes
//		invalid, unusable and unavailable for further
//		'write' operations.
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
func (fBufWriterNanobot *fileBufferWriterNanobot) flushAndClose(
	fBufWriter *FileBufferWriter,
	fBufWriterLabel string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fBufWriterNanobot.lock == nil {
		fBufWriterNanobot.lock = new(sync.Mutex)
	}

	fBufWriterNanobot.lock.Lock()

	defer fBufWriterNanobot.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileBufferWriterNanobot." +
		"flushAndClose()"

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
			"Error: The FileBufferReader instance passed\n"+
			"as input parameter '%v' is invalid!\n"+
			"'%v' is a 'nil' pointer.\n",
			ePrefix.String(),
			fBufWriterLabel,
			fBufWriterLabel)

		return err
	}

	var fBufWriterMolecule = new(fileBufferWriterMolecule)

	err = fBufWriterMolecule.
		flush(
			fBufWriter,
			fBufWriterLabel,
			ePrefix.XCpy("fBufWriter"))

	var err2 error

	err2 = fBufWriterMolecule.
		close(
			fBufWriter,
			fBufWriterLabel,
			ePrefix.XCpy("fBufWriter"))

	if err2 != nil {

		err = errors.Join(err, err2)

	}

	return err
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
//		FileBufferWriter.FlushAndClose().
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

	if bufSize <= 0 {

		bufSize = 4096
	}

	fBufWriter.fileWriter = bufio.NewWriterSize(
		writer,
		bufSize)

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

	if bufSize <= 0 {

		bufSize = 4096
	}

	var err2 error
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
	/*
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

			return fInfoPlus, err
		}
	*/
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

	fBufWriter.targetWriteFileName = pathFileName

	fBufWriter.fileWriter = bufio.NewWriterSize(
		fBufWriter.filePtr,
		bufSize)

	return fInfoPlus, err
}

type fileBufferWriterMolecule struct {
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
//			fileBufferWriterMolecule.flush()
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
func (fBufWriterMolecule *fileBufferWriterMolecule) close(
	fBufWriter *FileBufferWriter,
	fBufWriterLabel string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fBufWriterMolecule.lock == nil {
		fBufWriterMolecule.lock = new(sync.Mutex)
	}

	fBufWriterMolecule.lock.Lock()

	defer fBufWriterMolecule.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileBufferWriterMolecule." +
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
			"Error: The FileBufferReader instance passed\n"+
			"as input parameter '%v' is invalid!\n"+
			"'%v' is a 'nil' pointer.\n",
			ePrefix.String(),
			fBufWriterLabel,
			fBufWriterLabel)

		return err
	}

	if fBufWriter.fileWriter == nil {

		err = fmt.Errorf("%v\n"+
			"-------------------------------------------\n"+
			"Error: The FileBufferReader instance passed\n"+
			"as input parameter '%v' is invalid!\n"+
			"The internal bufio.Writer object is a 'nil' pointer.\n",
			ePrefix.String(),
			fBufWriterLabel)

		return err
	}

	if fBufWriter.filePtr != nil {

		var err2 error

		err2 = fBufWriter.filePtr.Close()

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Error returned while closing the target 'target' file!\n"+
				"fBufWriter.filePtr.Close()\n"+
				"Target Read File = '%v'\n"+
				"Error = \n%v\n",
				ePrefix.String(),
				fBufWriter.targetWriteFileName,
				err2.Error())

		}
	}

	fBufWriter.targetWriteFileName = ""

	fBufWriter.filePtr = nil

	fBufWriter.fileWriter = nil

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
//	fileBufferWriterMolecule.close()'.
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
func (fBufWriterMolecule *fileBufferWriterMolecule) flush(
	fBufWriter *FileBufferWriter,
	fBufWriterLabel string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fBufWriterMolecule.lock == nil {
		fBufWriterMolecule.lock = new(sync.Mutex)
	}

	fBufWriterMolecule.lock.Lock()

	defer fBufWriterMolecule.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileBufferWriterMolecule." +
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
			"Error: The FileBufferReader instance passed\n"+
			"as input parameter '%v' is invalid!\n"+
			"'%v' is a 'nil' pointer.\n",
			ePrefix.String(),
			fBufWriterLabel,
			fBufWriterLabel)

		return err
	}

	if fBufWriter.fileWriter == nil {

		err = fmt.Errorf("%v\n"+
			"-------------------------------------------\n"+
			"Error: The FileBufferReader instance passed\n"+
			"as input parameter '%v' is invalid!\n"+
			"The internal bufio.Writer object is a 'nil' pointer.\n",
			ePrefix.String(),
			fBufWriterLabel)

		return err
	}

	var err2 error

	err2 = fBufWriter.fileWriter.Flush()

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned while flushing the 'write' buffer!\n"+
			"%v.fileWriter.Flush()\n"+
			"Error = \n%v\n",
			ePrefix.String(),
			fBufWriterLabel,
			err2.Error())

	}

	return err
}
