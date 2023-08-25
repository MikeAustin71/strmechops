package strmech

import (
	"bufio"
	"errors"
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"math"
	"os"
	"strings"
	"sync"
)

// FileBufferReader
//
// The FileBufferReader type is a wrapper for
// 'bufio.Reader'. As such, FileBufferReader supports
// incremental or buffered read operations from the target
// data source.
//
// This type and its associated methods are designed to
// facilitate data 'read' operations. The most common
// data source for these read operations is assumed to be
// a data file residing on an attached storage drive.
// However, any object implementing the io.Reader
// interface may be used as a data source.
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
//	(1)	Use the 'New' and 'Setter' methods to create
//		valid instances of FileBufferReader.
//
//	(2)	FileBufferReader implements the following
//		interfaces:
//
//			io.Reader
//			io.Closer
//			io.WriterTo
//
// ----------------------------------------------------------------
//
// # Best Practice
//
//	(1)	Create a new instance of FileBufferReader using
//		any of the 'New' methods:
//
//		FileBufferReader.New()
//		FileBufferReader.NewFileMgr()
//		FileBufferReader.NewPathFileName()
//
//		(a)	The New() method is used when an instance of
//			io.Reader is created externally by the user
//			and passed to the FileBufferReader.New()
//			method.
//
//			Under this scenario, the user is independently
//			responsible for clean-up of the internal
//			bufio.Reader object after all 'read' operations
//			have been completed for the current instance
//			of FileBufferReader.
//
//			Once all FileBufferReader 'read' operations
//			have been completed, call method Close() to
//			perform local FileBufferReader clean-up tasks.
//
//		(b)	The NewFileMgr() method receives an instance
//			of FileMgr which identifies a path and file
//			name which will be configured as data source
//			for subsequent file 'read' operations.
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
//		(c)	The NewPathFileName() method allows for the
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
//		of data from the internal bufio.Reader object.
//		This 'read' target may be a file or any other
//		object which implements the io.Reader interface.
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
//		internal bufio.Reader object, the user must call
//		method Close() to perform necessary clean-up
//		tasks.
//
//		Once method Close() is called, the current
//		FileBufferReader instance becomes unusable and
//		should be discarded.
type FileBufferReader struct {
	bufioReader        *bufio.Reader
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
// # BE ADVISED
//
//	This method implements the io.Closer interface.
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
func (fBufReader *FileBufferReader) Close() error {

	if fBufReader.lock == nil {
		fBufReader.lock = new(sync.Mutex)
	}

	fBufReader.lock.Lock()

	defer fBufReader.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"FileBufferReader."+
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

// Buffered
//
// This method returns the number of bytes that can be
// read from the current 'read' buffer.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://pkg.go.dev/bufio#Reader.Buffered
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
//		This return value contains the number of bytes
//		which can be read from the current 'read' buffer
//		encapsulated by the current instance of
//		FileBufferReader.
func (fBufReader *FileBufferReader) Buffered() int {

	if fBufReader.lock == nil {
		fBufReader.lock = new(sync.Mutex)
	}

	fBufReader.lock.Lock()

	defer fBufReader.lock.Unlock()

	if fBufReader.bufioReader == nil {

		return 0
	}

	return fBufReader.bufioReader.Buffered()
}

// Discard
//
// This method skips the next 'numBytesToDiscard' bytes,
// returning the number of bytes actually discarded.
//
// If Discard skips fewer than 'numBytesToDiscard' bytes,
// it also returns an error.
//
// If 0 <= discardedBytes <= FileBufferReader.Buffered(),
// Discard is guaranteed to succeed without reading from
// the underlying io.Reader.
//
// Effectively, this method provides a means of
// repositioning the 'reader' to beginning reading at the
// next byte after the discarded bytes.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://pkg.go.dev/bufio#Reader.Discard
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numBytesToDiscard			int
//
//		This value contains the number of bytes to
//		discard. Effectively, this repositions the reader
//		to begin the next 'read' operation at the next
//		byte beyond the discarded bytes.
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
//	discardedBytes				int
//
//		The number of bytes discarded. If this value is
//		not equal to input parameter 'numBytesToDiscard',
//		an error will also be returned.
//
//		Remember that the number of bytes 'discarded'
//		repositions the 'reader' to begin the next read
//		operation at the next byte beyond the bytes
//		'discarded'.
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
func (fBufReader *FileBufferReader) Discard(
	numBytesToDiscard int,
	errorPrefix interface{}) (
	discardedBytes int,
	err error) {

	if fBufReader.lock == nil {
		fBufReader.lock = new(sync.Mutex)
	}

	fBufReader.lock.Lock()

	defer fBufReader.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileBufferReader."+
			"Discard()",
		"")

	if err != nil {

		return discardedBytes, err
	}

	if fBufReader.bufioReader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of 'FileBufferReader' is invalid!\n"+
			"The internal bufio.Reader object has NOT been initialized.\n"+
			"Call one of the 'New' or 'Setter' methods when creating\n"+
			"an instance of 'FileBufferReader'\n",
			ePrefix.String())

		return discardedBytes, err
	}

	var err1 error

	discardedBytes,
		err1 = fBufReader.bufioReader.
		Discard(numBytesToDiscard)

	if err1 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by fBufReader.bufioReader.Discard()\n"+
			"Error= \n%v\n",
			ePrefix.String(),
			err1.Error())
	}

	return discardedBytes, err
}

// NewIoReader
//
// Receives input parameter, 'reader', which implements the
// io.Reader interface.
//
// The io.Reader object is used to configure and return a
// fully initialized instance of FileBufferReader.
//
// The returned instance of FileBufferReader may be used
// to read data from the data source identified by the
// io.Reader object, 'reader'.
//
// The size of the internal 'read' buffer is controlled by
// input parameter 'bufSize'. The minimum buffer size is
// 16-bytes. If 'bufSize' is set to a value less than
// "16", it will be automatically reset to the default
// value of 4096-bytes.
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
//	The user is responsible for 'closing' the instance of
//	io.Reader passed as input parameter 'reader'. The
//	FileBufferReader.Close() method	will NOT close the
//	'reader' object.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	reader						io.Reader
//
//		An object which implements io.Reader interface.
//		This object will be used as a data source for
//		'read' operations.
//
//		The io.Reader object may be a file pointer of
//		type *os.File because file pointers of this type
//		implement the io.Reader interface.
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
//		Remember that the user is responsible for
//		'closing' this io.Reader object. The
//		FileBufferReader.Close() method will NOT close
//		this io.Reader object.
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
//		'bufSize' is set to a value less than "16", it
//		will be automatically reset to the default buffer
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
func (fBufReader *FileBufferReader) NewIoReader(
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
			"NewIoReader()",
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

// NewFileMgr
//
// Receives an instance of FileMgr as input parameter
// 'fileMgr'.
//
// The instance of FileBufferReader returned by this
// method will configure the file identified by 'fileMgr'
// as the data source for file 'read' operations.
//
// This target 'read' file identified by 'fileMgr' is
// opened for either 'read-only' or 'read/write'
// operations depending on input parameter
// 'openFileReadWrite'.
//
// The size of the internal 'read' buffer is controlled by
// input parameter 'bufSize'. The minimum buffer size is
// 16-bytes. If 'bufSize' is set to a value less than
// "16", it will be automatically reset to the default
// value of 4096-bytes.
//
// If the target path and file identified by 'fileMgr' do
// not currently exist on an attached storage drive, an
// error will be returned.
//
// Upon completion, this method returns a fully
// configured instance of FileBufferReader.
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
//	(1)	As a precaution, the incoming 'fileMgr' object
//		will be closed before configuring the internal
//		io.Reader object for the returned instance of
//		FileBufferReader.
//
//	(2)	When all read operations have been completed and
//		there is no further need for the returned
//		instance of FileBufferReader, the user is
//		responsible for 'closing' and releasing the
//		associated memory resources by calling the method
//		FileIoReader.Close().
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fileMgr						*FileMgr
//
//		A pointer to an instance of FileMgr. The file
//		identified by 'fileMgr' will be used as a data
//		source for 'read' operations.
//
//		If the path and file name encapsulated by
//		'fileMgr' do not currently exist on an attached
//		storage drive, an error will be returned.
//
//		The instance of FileBufferReader returned by this
//		method will configure the file identified by
//		'fileMgr' as the data source for file 'read'
//		operations.
//
//		As a precaution, the incoming 'fileMgr' object
//		will be closed before configuring the returned
//		FileBufferReader with a new internal bufio.Reader object.
//
//		If the path and file name specified by 'fileMgr'
//		does NOT exist on an attached storage drive, an
//		error will be returned.
//
//	openFileReadWrite			bool
//
//		If this parameter is set to 'true', the target
//		'read' file identified by input parameter
//		'fileMgr' will be opened for both 'read' and
//		'write' operations.
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
//	newFileBufReader			FileBufferReader
//
//		If this method completes successfully, a fully
//		configured instance of FileBufferReader will
//		be returned.
//
//		This returned instance of FileBufferReader will
//		configure the file identified by input parameter
//		'fileMgr' is a data source for file 'read'
//		operations.
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
func (fBufReader *FileBufferReader) NewFileMgr(
	fileMgr *FileMgr,
	openFileReadWrite bool,
	bufSize int,
	errorPrefix interface{}) (
	fInfoPlus FileInfoPlus,
	newFileBufReader FileBufferReader,
	err error) {

	if fBufReader.lock == nil {
		fBufReader.lock = new(sync.Mutex)
	}

	fBufReader.lock.Lock()

	defer fBufReader.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileBufferReader."+
			"NewFileMgr()",
		"")

	if err != nil {

		return fInfoPlus, newFileBufReader, err
	}

	fInfoPlus,
		err = new(fileBufferReaderMicrobot).
		setFileMgr(
			&newFileBufReader,
			"newFileBufReader",
			fileMgr,
			"fileMgr",
			openFileReadWrite,
			bufSize,
			ePrefix.XCpy(
				"fileMgr"))

	return fInfoPlus, newFileBufReader, err
}

// NewPathFileName
//
// Receives a path and file name as an input parameter
// string, 'pathFileName' and returns a new, fully
// configured instance of FileBufferReader.
//
// The target 'read' file identified by 'pathFileName'
// is opened for either 'read-only' or 'read/write'
// operations depending on input parameter
// 'openFileReadWrite'.
//
// This target 'read' file identified by 'pathFileName'
// will be used to create a file pointer (*os.File) which
// in turn will be used to configure the internal
// bufio.Reader.
//
// The size of the internal 'read' buffer is controlled by
// input parameter 'bufSize'. The minimum buffer size is
// 16-bytes. If 'bufSize' is set to a value less than
// "16", it will be automatically reset to the default
// value of 4096-bytes.
//
// If the target path and file identified by
// 'pathFileName' do not currently exist on an attached
// storage drive, an error will be returned.
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
//		the file which will be configured as a new data
//		source for 'read' operations.
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
//		will be automatically reset to the default buffer
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
//	newFileBufReader			FileBufferReader
//
//		If this method completes successfully, a fully
//		configured instance of FileBufferReader will
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
func (fBufReader *FileBufferReader) NewPathFileName(
	pathFileName string,
	openFileReadWrite bool,
	bufSize int,
	errorPrefix interface{}) (
	fInfoPlus FileInfoPlus,
	newFileBufReader FileBufferReader,
	err error) {

	if fBufReader.lock == nil {
		fBufReader.lock = new(sync.Mutex)
	}

	fBufReader.lock.Lock()

	defer fBufReader.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileBufferReader."+
			"NewPathFileName()",
		"")

	if err != nil {

		return fInfoPlus, newFileBufReader, err
	}

	fInfoPlus,
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

	return fInfoPlus, newFileBufReader, err
}

// Peek
//
// This method returns the next n bytes without advancing
// the reader. The bytes stop being valid at the next
// read call.
//
// If Peek returns fewer bytes than 'nextBytes' value, it
// also returns an error explaining why the read is
// short.
//
// If 'nextBytes' value is larger than the buffer size
// for the current FileBufferReader, the error
// 'ErrBufferFull' will be returned.
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
func (fBufReader *FileBufferReader) Peek(
	nextBytes int,
	errorPrefix interface{}) (
	peekBytes []byte,
	err error) {

	if fBufReader.lock == nil {
		fBufReader.lock = new(sync.Mutex)
	}

	fBufReader.lock.Lock()

	defer fBufReader.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileBufferReader."+
			"Peek()",
		"")

	if err != nil {

		return peekBytes, err
	}

	if fBufReader.bufioReader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of 'FileBufferReader' is invalid!\n"+
			"The internal bufio.Reader object has NOT been initialized.\n"+
			"Call one of the 'New' or 'Setter' methods when creating\n"+
			"an instance of 'FileBufferReader'\n",
			ePrefix.String())

		return peekBytes, err
	}

	var err2 error

	peekBytes,
		err2 = fBufReader.bufioReader.Peek(nextBytes)

	if err2 != nil {

		if errors.Is(err2, bufio.ErrBufferFull) {

			err = fmt.Errorf("%v\n"+
				"Error: %v\n"+
				"The input parameter 'nextBytes' is greater\n"+
				"than the length of the FileBufferReader buffer.\n"+
				" Input parameter 'nextBytes'= %v-bytes\n"+
				"FileBufferReader Buffer Size= %v-bytes\n",
				ePrefix.String(),
				bufio.ErrBufferFull.Error(),
				nextBytes,
				fBufReader.bufioReader.Size())

		} else {

			err = fmt.Errorf("%v\n"+
				"Error returned by fBufReader.bufioReader.Peek()\n"+
				"Input parameter 'nextBytes'= %v-bytes\n"+
				"Error=\n%v\n",
				ePrefix.String(),
				nextBytes,
				err2.Error())

		}

	}

	return peekBytes, err
}

// Read
//
// Reads a selection data from the pre-configured
// internal bufio.Reader data source encapsulated in the
// current instance of FileBufferReader.
//
// This method is a wrapper for the bufio 'Reader.Read'
// method.
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
//		appropriate error message.
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

	if fBufReader.bufioReader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of 'FileBufferReader' is invalid!\n"+
			"The internal bufio.Reader object has NOT been initialized.\n"+
			"Call one of the 'New' or 'Setter' methods when creating\n"+
			"an instance of 'FileBufferReader'\n",
			ePrefix.String())

		return numOfBytesRead, err
	}

	var err2 error

	numOfBytesRead,
		err2 = fBufReader.bufioReader.Read(bytesRead)

	if err2 != nil {

		if err2 != io.EOF {

			err = fmt.Errorf("%v\n"+
				"Error returned by fBufReader.bufioReader.Read(bytesRead).\n"+
				"Error=\n%v\n",
				ePrefix.String(),
				err2.Error())

		} else {

			err = io.EOF

		}

	}

	return numOfBytesRead, err
}

// ReadAllTextLines
//
// Reads text lines from the internal bufio.Reader object
// encapsulated in the current instance of
// FileBufferReader.
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
// instance of FileBufferReader.
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
// entire contents of the FileBufferReader internal
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
//		FileBufferReader, into memory.
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
//		io.Reader object upon completion.
//
//		The user is responsible for performing required
//		clean-up tasks by calling the local method:
//
//			FileBufferReader.Close()
//
//	(3)	If the current instance of FileBufferReader has
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
//		FileBufferReader.
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
//		FileBufferReader will be stored as individual
//		strings in the string array encapsulated by
//		'outputLinesArray'.
//
//	autoCloseOnExit				bool
//
//		When this parameter is set to 'true', this
//		method will automatically perform all required
//		clean-up tasks for the current instance of
//		FileBufferReader upon completion. Specifically,
//		the internal bufio.Reader object will be properly
//		'closed' and there will be no need to make a
//		separate call to the local method:
//
//			FileBufferReader.Close()
//
//		After completing the 'close' operation, the
//		current instance of FileBufferReader will be
//		invalid and unavailable for all future 'read'
//		operations.
//
//		If input parameter 'autoCloseOnExit' is set to
//		'false', this method will NOT automatically
//		'close' the internal bufio.Reader object for the
//		current instance of FileBufferReader.
//		Consequently, the user will then be responsible
//		for 'closing' the internal bufio.Reader object by
//		calling the local method:
//
//			FileBufferReader.Close()
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
//		FileBufferReader. This number also specifies
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
//		FileBufferReader.
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
func (fBufReader *FileBufferReader) ReadAllTextLines(
	maxNumOfLines int,
	endOfLineDelimiters *StringArrayDto,
	outputLinesArray *StringArrayDto,
	autoCloseOnExit bool,
	errorPrefix interface{}) (
	numOfLinesRead int,
	numOfBytesRead int64,
	err error) {

	if fBufReader.lock == nil {

		fBufReader.lock = new(sync.Mutex)

	}

	fBufReader.lock.Lock()

	defer fBufReader.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileBufferReader."+
			"ReadAllTextLines()",
		"")

	if err != nil {

		return numOfLinesRead,
			numOfBytesRead,
			err
	}

	if fBufReader.bufioReader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of 'FileBufferReader' is invalid!\n"+
			"The internal bufio.Reader object has NOT been initialized.\n"+
			"Call one of the 'New' or 'Setter' methods when creating\n"+
			"an instance of 'FileBufferReader'\n",
			ePrefix.String())

		return numOfLinesRead,
			numOfBytesRead,
			err
	}

	var err1 error

	numOfLinesRead,
		numOfBytesRead,
		err1 = new(fileHelperMolecule).
		readerScanLines(
			fBufReader.bufioReader,
			"fBufReader.bufioReader",
			endOfLineDelimiters,
			outputLinesArray,
			maxNumOfLines,
			ePrefix.XCpy("fBufReader.bufioReader"))

	var err2 error

	if autoCloseOnExit == true {

		err2 = new(fileBufferReaderMolecule).close(
			fBufReader,
			"fBufReader",
			ePrefix.XCpy("fBufReader"))

	}

	err = errors.Join(err1, err2)

	return numOfLinesRead,
		numOfBytesRead,
		err
}

// ReadAllStrBuilder
//
// Reads the entire contents of the internal bufio.Reader
// for the current instance of FileBufferReader as
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
// entire contents of the FileBufferReader internal
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
//		FileBufferReader, into memory.
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
//		io.Reader object upon completion.
//
//		The user is responsible for performing required
//		clean-up tasks by calling the local method:
//
//			FileBufferReader.Close()
//
//	(3)	If the current instance of FileBufferReader has
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
//		the current instance of FileBufferReader will be
//		read and stored in 'strBuilder'.
//
//	autoCloseOnExit				bool
//
//		When this parameter is set to 'true', this
//		method will automatically perform all required
//		clean-up tasks for the current instance of
//		FileBufferReader upon completion. Specifically,
//		the internal bufio.Reader object will be properly
//		'closed' and there will be no need to make a
//		separate call to the local method:
//
//			FileBufferReader.Close()
//
//		After completing the 'close' operation, the
//		current instance of FileBufferReader will be
//		invalid and unavailable for all future 'read'
//		operations.
//
//		If input parameter 'autoCloseOnExit' is set to
//		'false', this method will NOT automatically
//		'close' the internal bufio.Reader object for the
//		current instance of FileBufferReader.
//		Consequently, the user will then be responsible
//		for 'closing' the internal bufio.Reader object by
//		calling the local method:
//
//			FileBufferReader.Close()
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
//		FileBufferReader.
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
func (fBufReader *FileBufferReader) ReadAllStrBuilder(
	strBuilder *strings.Builder,
	autoCloseOnExit bool,
	errorPrefix interface{}) (
	numOfBytesRead int64,
	err error) {

	if fBufReader.lock == nil {
		fBufReader.lock = new(sync.Mutex)
	}

	fBufReader.lock.Lock()

	defer fBufReader.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileBufferReader."+
			"ReadAllStrBuilder()",
		"")

	if err != nil {

		return numOfBytesRead, err
	}

	if fBufReader.bufioReader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of 'FileBufferReader' is invalid!\n"+
			"The internal bufio.Reader object has NOT been initialized.\n"+
			"Call one of the 'New' or 'Setter' methods when creating\n"+
			"an instance of 'FileBufferReader'\n",
			ePrefix.String())

		return numOfBytesRead, err
	}

	var err1 error

	numOfBytesRead,
		err1 = new(fileBufferReaderMicrobot).
		readAllStrBuilder(
			fBufReader,
			"fBufReader",
			strBuilder,
			ePrefix)

	var err2 error

	if autoCloseOnExit == true {

		err2 = new(fileBufferReaderMolecule).close(
			fBufReader,
			"fBufReader",
			ePrefix.XCpy("fBufReader"))

	}

	err = errors.Join(err1, err2)

	return numOfBytesRead, err
}

// ReadAllToString
//
// Reads the entire contents of the internal bufio.Reader
// for the current instance of FileBufferReader and returns
// these contents as a single string ('contentsStr').
//
// If a processing error is encountered, an appropriate
// error with an error message will be returned. When
// the end-of-file is encountered during the 'read'
// process, the returned error object will be set to
// 'nil' and no error will be returned.
//
// It naturally follows that this method will read the
// entire contents of the FileBufferReader internal
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
//		FileBufferReader, into memory.
//
//		BE CAREFUL when reading large files!
//
//		Depending on the memory resources available to
//		your computer, you may run out of memory when
//		reading large files and returning their contents
//		as a single string ('contentsStr').
//
//	(2)	This method will NOT automatically close the
//		io.Reader object upon completion.
//
//		The user is responsible for performing required
//		clean-up tasks by calling the local method:
//
//			FileBufferReader.Close()
//
//	(3)	If the current instance of FileBufferReader has
//		NOT been properly initialized, an error will be
//		returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	autoCloseOnExit				bool
//
//		When this parameter is set to 'true', this
//		method will automatically perform all required
//		clean-up tasks for the current instance of
//		FileBufferReader upon completion. Specifically,
//		the internal bufio.Reader object will be properly
//		'closed' and there will be no need to make a
//		separate call to the local method:
//
//			FileBufferReader.Close()
//
//		After completing the 'close' operation, the
//		current instance of FileBufferReader will be
//		invalid and unavailable for all future 'read'
//		operations.
//
//		If input parameter 'autoCloseOnExit' is set to
//		'false', this method will NOT automatically
//		'close' the internal bufio.Reader object for the
//		current instance of FileBufferReader.
//		Consequently, the user will then be responsible
//		for 'closing' the internal bufio.Reader object by
//		calling the local method:
//
//			FileBufferReader.Close()
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
//		FileBufferReader.
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
//		FileBufferReader will be returned in this string.
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
func (fBufReader *FileBufferReader) ReadAllToString(
	autoCloseOnExit bool,
	errorPrefix interface{}) (
	numOfBytesRead int64,
	contentsStr string,
	err error) {

	if fBufReader.lock == nil {
		fBufReader.lock = new(sync.Mutex)
	}

	fBufReader.lock.Lock()

	defer fBufReader.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileBufferReader."+
			"ReadAllToString()",
		"")

	if err != nil {

		return numOfBytesRead, contentsStr, err
	}

	if fBufReader.bufioReader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of 'FileBufferReader' is invalid!\n"+
			"The internal bufio.Reader object has NOT been initialized.\n"+
			"Call one of the 'New' or 'Setter' methods when creating\n"+
			"an instance of 'FileBufferReader'\n",
			ePrefix.String())

		return numOfBytesRead, contentsStr, err
	}

	strBuilder := new(strings.Builder)
	var err1 error

	numOfBytesRead,
		err1 = new(fileBufferReaderMicrobot).
		readAllStrBuilder(
			fBufReader,
			"fBufReader",
			strBuilder,
			ePrefix)

	if err1 == nil {

		contentsStr = strBuilder.String()
	}

	var err2 error

	if autoCloseOnExit == true {

		err2 = new(fileBufferReaderMolecule).close(
			fBufReader,
			"fBufReader",
			ePrefix.XCpy("fBufReader"))

	}

	err = errors.Join(err1, err2)

	return numOfBytesRead, contentsStr, err
}

// SetIoReader
//
// This method will completely re-initialize the current
// instance of FileBufferReader using the io.Reader object
// passed as input parameter 'reader'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	This method will delete, overwrite and reset all
//		pre-existing data values in the current instance
//		of FileBufferReader.
//
//	(2)	The user is responsible for 'closing' the
//		instance of io.Reader passed as input parameter
//		'reader'. The FileBufferReader.Close() method
//		will NOT close the 'reader' object.
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
//
//		An object which implements io.Reader interface.
//		This object will be used as a data source for
//		'read' operations.
//
//		The io.Reader object may be a file pointer of
//		type *os.File because file pointers of this type
//		implement the io.Reader interface.
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
//		Remember that the user is responsible for
//		'closing' this io.Reader object. The FileIoReader
//		method 'Close()' will NOT close this io.Reader
//		object.
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
func (fBufReader *FileBufferReader) SetIoReader(
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
			"SetIoReader()",
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

// SetFileMgr
//
// This method will completely re-initialize the current
// instance of FileBufferReader using the path and file
// name identified by the FileMgr instance passed as
// input parameter 'fileMgr'.
//
// The file identified by 'fileMgr' will be used to
// reconfigure the internal bufio.Reader encapsulated in
// the current instance of FileBufferReader.
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
//	(1)	This method will delete, overwrite and reset all
//		pre-existing data values in the current instance
//		of FileBufferReader.
//
//	(2)	As a precaution, the incoming 'fileMgr' object
//		will be closed before configuring the current
//		FileBufferReader instance with a new, internal
//		io.Reader object.
//
//	(3)	When all read operations have been completed and
//		there is no further need for the current instance
//		of FileBufferReader, the user is responsible for
//		'closing' and releasing the associated memory
//		resources by calling the local method
//		FileIoReader.Close().
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fileMgr						*FileMgr
//
//		A pointer to an instance of FileMgr. The file
//		identified by 'fileMgr' will be used as a data
//		source for 'read' operations.
//
//		If the path and file name encapsulated by
//		'fileMgr' do not currently exist on an attached
//		storage drive, an error will be returned.
//
//		The instance of FileBufferReader returned by this
//		method will configure the file identified by
//		'fileMgr' as the data source for file 'read'
//		operations.
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
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fBufReader *FileBufferReader) SetFileMgr(
	fileMgr *FileMgr,
	openFileReadWrite bool,
	bufSize int,
	errorPrefix interface{}) (
	fInfoPlus FileInfoPlus,
	err error) {

	if fBufReader.lock == nil {
		fBufReader.lock = new(sync.Mutex)
	}

	fBufReader.lock.Lock()

	defer fBufReader.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileBufferReader."+
			"SetFileMgr()",
		"")

	if err != nil {

		return fInfoPlus, err
	}

	fInfoPlus,
		err = new(fileBufferReaderMicrobot).
		setFileMgr(
			fBufReader,
			"fBufReader",
			fileMgr,
			"fileMgr",
			openFileReadWrite,
			bufSize,
			ePrefix.XCpy(
				"fileMgr"))

	return fInfoPlus, err
}

// SetPathFileName
//
// This method will completely re-initialize the current
// instance of FileBufferReader using the path and file
// name passed as input parameter 'pathFileName'.
//
// The path and file name specified by 'pathFileName'
// will be used to reconfigure the internal bufio.Reader
// encapsulated in the current instance of
// FileBufferReader.
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
//	(1)	This method will delete, overwrite and reset all
//		pre-existing data values in the current instance
//		of FileBufferReader.
//
//	(2)	When all read operations have been completed and
//		there is no further need for the current instance
//		of FileBufferReader, the user is responsible for
//		'closing' and releasing the associated memory
//		resources by calling the method
//		FileBufferReader.Close().
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathFileName				string
//
//		This string contains the path and file name of
//		the file which will be configured as a new data
//		source for 'read' operations.
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
//		will be automatically reset to the default buffer
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
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fBufReader *FileBufferReader) SetPathFileName(
	pathFileName string,
	openFileReadWrite bool,
	bufSize int,
	errorPrefix interface{}) (
	fInfoPlus FileInfoPlus,
	err error) {

	if fBufReader.lock == nil {
		fBufReader.lock = new(sync.Mutex)
	}

	fBufReader.lock.Lock()

	defer fBufReader.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileBufferReader."+
			"SetPathFileName()",
		"")

	if err != nil {

		return fInfoPlus, err
	}

	fInfoPlus,
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

	return fInfoPlus, err
}

// Size
//
// This method returns the size of the underlying 'read'
// buffer in bytes.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://pkg.go.dev/bufio#Reader.Size
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
//		This return value contains the size of the
//		underlying 'read' buffer in bytes for the current
//		instance of FileBufferReader.
func (fBufReader *FileBufferReader) Size() int {

	if fBufReader.lock == nil {
		fBufReader.lock = new(sync.Mutex)
	}

	fBufReader.lock.Lock()

	defer fBufReader.lock.Unlock()

	if fBufReader.bufioReader == nil {

		return 0
	}

	return fBufReader.bufioReader.Size()
}

// WriteTo
//
// This method implements the io.WriterTo interface.
//
// Input parameter 'writer' passes an io.Writer object.
// This method will then proceed to read the entire
// contents of the bufio.Reader object encapsulated by
// the current instance of FileBufferReader and write
// this data to the io.Writer object.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	This method implements the io.WriterTo interface.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	writer						io.Writer
//
//		This instance of io.Writer will be used as the
//		'write' destination for all data read from the
//		bufio.Reader object encapsulated by the current
//		instance of FileBufferReader.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	numOfBytesProcessed			int64
//
//		The number of bytes read from the internal
//		bufio.Reader and written to the destination
//		io.Writer object passed as input parameter
//		'writer'.
//
//	err							error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message.
func (fBufReader *FileBufferReader) WriteTo(
	writer io.Writer) (
	numOfBytesProcessed int64,
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
			"WriteTo()",
		"")

	if err != nil {

		return numOfBytesProcessed, err
	}

	if fBufReader.bufioReader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of 'FileBufferReader' is invalid!\n"+
			"The internal bufio.Reader object has NOT been initialized.\n"+
			"Call one of the 'New' or 'Setter' methods to create a\n"+
			"valid instance of 'FileBufferReader'\n",
			ePrefix.String())

		return numOfBytesProcessed, err
	}

	if writer == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'writer' is invalid!\n"+
			"'writer' has a 'nil' value.\n",
			ePrefix.String())

		return numOfBytesProcessed, err
	}

	bufSize := fBufReader.bufioReader.Size()

	if bufSize < 16 {
		// Reset to default buffer size of 4096
		fBufReader.bufioReader.Reset(nil)

		bufSize = fBufReader.bufioReader.Size()

	}

	if bufSize < 16 {

		err = fmt.Errorf("%v\n"+
			"Error: The attempt to reset the bufio.Reader buffer\n"+
			"size to 4096-bytes Failed!\n",
			ePrefix.String())

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
				"The 'Read' operation failed locate io.EOF\n"+
				"otherwise known as the end-of-file for this\n"+
				"underlying io.Reader object.\n"+
				"Read Cycle Count= %v\n",
				ePrefix.String(),
				cycleCnt)

			break
		}

		numBytesRead,
			err1 = fBufReader.bufioReader.Read(bytesRead)

		if err1 != nil &&
			err1 != io.EOF {

			err = fmt.Errorf("%v\n"+
				"Error: fBufReader.bufioReader.Read(bytesRead)\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				err1.Error())

			break

		}

		if numBytesRead > 0 {

			numBytesWritten,
				err2 = writer.Write(bytesRead[:numBytesRead])

			if err2 != nil {

				err = fmt.Errorf("%v\n"+
					"Error returned by writer.Write(bytesRead[:numBytesRead])\n"+
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

	} // for

	return numOfBytesProcessed, err
}

type fileBufferReaderMicrobot struct {
	lock *sync.Mutex
}

// readAllStrBuilder
//
// Reads the entire contents of the internal bufio.Reader
// for instance of FileBufferReader passed as input
// parameter 'fBufReader'. These contents are converted
// to a string which is then stored and returned through
// an instance of strings.Builder passed as input
// parameter 'strBuilder'.
//
// If a processing error is encountered, an appropriate
// error with an error message will be returned. When
// the end-of-file is encountered during the 'read'
// process, the returned error object will be set to
// 'nil' and no error will be returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fBufReader					*FileBufferReader
//
//		A pointer to an instance of FileBufferReader.
//
//		The entire contents of the internal bufio.Reader
//		object encapsulated in this FileBufferReader
//		instance will be extracted and returned as a
//		string through input parameter 'strBuilder'.
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
//	strBuilder					*strings.Builder
//
//		A pointer to an instance of strings.Builder. The
//		entire contents of the internal bufio.Reader for the
//		FileBufferReader instance passed as 'fBufReader'
//		will be extracted and stored as a string in
//		'strBuilder'.
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
//	numBytesRead				int64
//
//		If this method completes successfully, this
//		integer value will equal the number of bytes
//		read from the internal bufio.Reader object
//		encapsulated by the FileBufferReader instance
//		passed as input parameter 'fBufReader'.
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
func (fBufReaderMicrobot *fileBufferReaderMicrobot) readAllStrBuilder(
	fBufReader *FileBufferReader,
	fBufReaderLabel string,
	strBuilder *strings.Builder,
	errPrefDto *ePref.ErrPrefixDto) (
	numOfBytesRead int64,
	err error) {

	if fBufReaderMicrobot.lock == nil {
		fBufReaderMicrobot.lock = new(sync.Mutex)
	}

	fBufReaderMicrobot.lock.Lock()

	defer fBufReaderMicrobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileBufferReaderMicrobot."+
			"readAllStrBuilder()",
		"")

	if err != nil {

		return numOfBytesRead, err
	}

	if len(fBufReaderLabel) == 0 {

		fBufReaderLabel = "fBufReader"
	}

	if fBufReader.bufioReader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of 'FileBufferReader' (%v) is invalid!\n"+
			"The internal '%v' bufio.Reader object has NOT been initialized.\n"+
			"Call one of the 'New' or 'Setter' methods when creating\n"+
			"an instance of 'FileBufferReader'\n",
			ePrefix.String(),
			fBufReaderLabel,
			fBufReaderLabel)

		return numOfBytesRead, err
	}

	var bytesRead = make([]byte, 4096)
	var localBytesRead int
	var err2 error

	for {

		localBytesRead,
			err2 = fBufReader.bufioReader.Read(bytesRead)

		if localBytesRead > 0 {

			strBuilder.Write(bytesRead[:localBytesRead])

			numOfBytesRead += int64(localBytesRead)

		}

		if err2 == io.EOF {

			break

		}

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Error returned by fBufReader.bufioReader.Read(bytesRead).\n"+
				"Error=\n%v\n",
				ePrefix.String(),
				err2.Error())

			break
		}

		clear(bytesRead)

	}

	return numOfBytesRead, err
}

// setFileMgr
//
// This 'setter' method is used to initialize new values
// for internal member variables contained in the
// instance of FileBufferReader passed as input parameter
// 'fBufReader'.
//
// The new internal bufio.Reader object assigned to
// 'fBufReader' is generated from the File Manager
// (FileMgr) instance passed as input parameter
// 'fileMgr'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1) This method will delete, overwrite and reset all
//		pre-existing data values in the instance of
//		FileBufferReader passed as input parameter
//		'fBufReader'.
//
//	(2) As a precaution, the incoming 'fileMgr' object
//		will be closed before configuring 'fBufReader'
//		with a new internal bufio.Reader object.
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
//	fileMgr						*FileMgr
//
//		A pointer to an instance of FileMgr. The file
//		identified by 'fileMgr' will be used as a
//		data source for 'read' operations performed by
//		method:
//
//			FileBufferReader.Read()
//
//		If the path and file name encapsulated by
//		'fileMgr' do not currently exist on an attached
//		storage drive, an error will be returned.
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
//		'read' file identified by input parameter
//		'fileMgr' will be opened for 'read' and 'write'
//		operations.
//
//		If 'openFileReadWrite' is set to 'false', the
//		target read file will be opened for 'read-only'
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
func (fBufReaderMicrobot *fileBufferReaderMicrobot) setFileMgr(
	fBufReader *FileBufferReader,
	fBufReaderLabel string,
	fileMgr *FileMgr,
	fileMgrLabel string,
	openFileReadWrite bool,
	bufSize int,
	errPrefDto *ePref.ErrPrefixDto) (
	fInfoPlus FileInfoPlus,
	err error) {

	if fBufReaderMicrobot.lock == nil {
		fBufReaderMicrobot.lock = new(sync.Mutex)
	}

	fBufReaderMicrobot.lock.Lock()

	defer fBufReaderMicrobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileBufferReaderMicrobot." +
		"setFileMgr()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return fInfoPlus, err
	}

	if len(fBufReaderLabel) == 0 {

		fBufReaderLabel = "fBufReader"
	}

	if len(fileMgrLabel) == 0 {

		fileMgrLabel = "fileMgr"
	}

	if fBufReader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: The FileBufferReader instance passed\n"+
			"as input parameter '%v' is invalid!\n"+
			"'%v' is a 'nil' pointer.\n",
			ePrefix.String(),
			fBufReaderLabel,
			fBufReaderLabel)

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

	err2 = new(fileMgrHelperAtom).isFileMgrValid(
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
		err = new(fileBufferReaderNanobot).
		setPathFileName(
			fBufReader,
			fBufReaderLabel,
			fileMgr.absolutePathFileName,
			fileMgrLabel,
			openFileReadWrite,
			bufSize,
			ePrefix.XCpy(fBufReaderLabel+"<-"+fileMgrLabel))

	return fInfoPlus, err
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
//		This instance of io.Reader will be used to
//		configure the internal bufio.Reader contained in
//		'fBufReader' and used to conduct 'read'
//		operations.
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
//		'read' buffer configured for the
//		FileBufferReader instance passed as input
//		parameter 'fBufReader'.
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

	fBufReader.bufioReader = bufio.NewReaderSize(
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
func (fBufReaderNanobot *fileBufferReaderNanobot) setPathFileName(
	fBufReader *FileBufferReader,
	fBufReaderLabel string,
	pathFileName string,
	pathFileNameLabel string,
	openFileReadWrite bool,
	bufSize int,
	errPrefDto *ePref.ErrPrefixDto) (
	fInfoPlus FileInfoPlus,
	err error) {

	if fBufReaderNanobot.lock == nil {
		fBufReaderNanobot.lock = new(sync.Mutex)
	}

	fBufReaderNanobot.lock.Lock()

	defer fBufReaderNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileBufferReaderNanobot." +
		"setPathFileName()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		return fInfoPlus, err
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

		return fInfoPlus, err
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

		return fInfoPlus, err
	}

	err = new(fileBufferReaderMolecule).close(
		fBufReader,
		fBufReaderLabel,
		ePrefix.XCpy(fBufReaderLabel))

	if err != nil {

		return fInfoPlus, err
	}

	if bufSize < 16 {

		bufSize = 4096
	}

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

		return fInfoPlus, err
	}

	if !pathFileDoesExist {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is invalid!\n"+
			"The 'reader' path and file name do NOT exist\n"+
			"on an attached storage drive.\n"+
			"%v= '%v'\n",
			ePrefix.String(),
			pathFileNameLabel,
			pathFileNameLabel,
			pathFileName)

		return fInfoPlus, err
	}

	if fInfoPlus.IsDir() {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'pathFileName' is invalid!\n"+
			"'pathFileName' is directory and NOT a file name.\n"+
			"pathFileName= '%v'\n",
			ePrefix.String(),
			pathFileName)

		return fInfoPlus, err
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

		return fInfoPlus, err
	}

	var fileOpenCfg FileOpenConfig

	fileOpenCfg,
		err = new(FileOpenConfig).New(
		ePrefix.XCpy("fileOpenCfg<-"),
		FOpenType.TypeReadOnly())

	if err != nil {

		return fInfoPlus, err
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

		return fInfoPlus, err
	}

	fBufReader.targetReadFileName = pathFileName

	fBufReader.bufioReader = bufio.NewReaderSize(
		fBufReader.filePtr,
		bufSize)

	return fInfoPlus, err
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

	fBufReader.bufioReader = nil

	return err
}
