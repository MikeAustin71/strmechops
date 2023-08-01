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
// read/write capability for files and any objects
// supporting the io.Reader/io.Writer interfaces.
//
// User can utilize this type to perform read and
// write operations in a single method call.
//
// Read and Write operations are performed using
// private, internal FileBufferReader and
// FileBufferWriter objects.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	To create a valid instance of FileBufferReadWrite,
//	users MUST call one of the 'New' methods.
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
//	This method will effectively render the current
//	instance of FileBufferReadWrite invalid and unusable
//	for any future 'read' or 'write' operations.
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
			"NewPathFileNames()",
		"")

	if err != nil {
		return err
	}

	err = new(fileBufferReadWriteNanobot).
		closeFileBufferReadWrite(
			fBufReadWrite,
			"fBufReadWrite",
			ePrefix.XCpy("Close-Readers&Writers"))

	return err
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
		"New()"

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
//		encapsulated in the returned instance of
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
//		encapsulated in the returned instance of
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
	FileBufferReadWrite,
	error) {

	if fBufReadWrite.lock == nil {
		fBufReadWrite.lock = new(sync.Mutex)
	}

	fBufReadWrite.lock.Lock()

	defer fBufReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	var newFBuffReadWrite = FileBufferReadWrite{}

	funcName := "FileBufferReadWrite." +
		"NewFileMgrs()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {

		return newFBuffReadWrite, err
	}

	err = new(fileBufferReadWriteMicrobot).
		setFileMgrs(
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

	return newFBuffReadWrite, err
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
//		The minimum reader buffer size is 16-bytes. If
//		'readerBuffSize' is set to a size less than "16",
//		it will be automatically reset to the default
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
//		'read' buffer created for the io.Writer
//		associated with the file identified by
//		'writerPathFileName'.
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
func (fBufReadWrite *FileBufferReadWrite) NewPathFileNames(
	readerPathFileName string,
	openReadFileReadWrite bool,
	readerBuffSize int,
	writerPathFileName string,
	openWriteFileReadWrite bool,
	writerBuffSize int,
	truncateExistingWriteFile bool,
	errorPrefix interface{}) (
	FileBufferReadWrite,
	error) {

	if fBufReadWrite.lock == nil {
		fBufReadWrite.lock = new(sync.Mutex)
	}

	fBufReadWrite.lock.Lock()

	defer fBufReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	var newFBuffReadWrite = FileBufferReadWrite{}

	funcName := "FileBufferReadWrite." +
		"NewPathFileNames()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {
		return newFBuffReadWrite, err
	}

	err = new(fileBufferReadWriteNanobot).
		setPathFileNames(
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

	return newFBuffReadWrite, err
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
// instance of FileBufferReadWrite was initialized with
// one of the 'New' methods.
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
func (fBufReadWrite *FileBufferReadWrite) ReadWriteAll(
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

	if fBufReadWrite.reader == nil ||
		fBufReadWrite.writer == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: The current instance of\n"+
			"FileBufferReadWrite is invalid!\n"+
			"Call one of the 'New' methods to\n"+
			"create a valid instance of FileBufferReadWrite.\n",
			ePrefix.String())

		return totalBytesRead, totalBytesWritten, err
	}

	var readErr, writeErr error
	var numOfBytesRead, numOfBytesWritten, cycleCount int

	byteArray := make([]byte,
		fBufReadWrite.reader.fileReader.Size())

	var errs []error

	for {

		cycleCount++

		numOfBytesRead,
			readErr =
			fBufReadWrite.reader.Read(byteArray)

		if readErr != nil &&
			readErr != io.EOF {

			var err2 error

			err2 = fmt.Errorf("%v\n"+
				"Error Reading Target Read File!\n"+
				"Cycle Count= %v\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				cycleCount,
				readErr.Error())

			errs = append(
				errs, err2)

			var err3 error

			err3 = fBufReadWrite.reader.Close(
				ePrefix.XCpy("reader-Close"))

			if err3 != nil {

				errs = append(
					errs, err3)

			}

			var err4 error

			err4 = fBufReadWrite.writer.Close(
				ePrefix.XCpy("writer-Close"))

			if err4 != nil {

				errs = append(
					errs, err4)

			}

			err = new(StrMech).ConsolidateErrors(errs)

			return totalBytesRead, totalBytesWritten, err
		}

		if numOfBytesRead > 0 {

			totalBytesRead += numOfBytesRead

			numOfBytesWritten,
				writeErr = fBufReadWrite.writer.Write(
				byteArray[0:numOfBytesRead])

			if writeErr != nil {

				var err2 error

				err2 = fmt.Errorf("%v\n"+
					"Error Writing Bytes To File!\n"+
					"Write Error=\n%v\n",
					ePrefix.String(),
					writeErr.Error())

				errs = append(
					errs, err2)

				var err3 error

				err3 = fBufReadWrite.reader.Close(
					ePrefix.XCpy("reader-Close"))

				if err3 != nil {

					errs = append(
						errs, err3)

				}

				var err4 error

				err4 = fBufReadWrite.writer.Close(
					ePrefix.XCpy("writer-Close"))

				if err4 != nil {

					errs = append(
						errs, err4)

				}

				err = new(StrMech).ConsolidateErrors(errs)

				return totalBytesRead, totalBytesWritten, err
			}

			totalBytesWritten += numOfBytesWritten
		}

		if numOfBytesRead != numOfBytesWritten {

			var err2 error

			err2 = fmt.Errorf("%v\n"+
				"Error Writing Bytes To File!\n"+
				"numOfBytesRead != numOfBytesWritten\n"+
				"numOfBytesRead = %v\n"+
				"numOfBytesWritten = %v\n",
				ePrefix.String(),
				numOfBytesRead,
				numOfBytesWritten)

			errs = append(
				errs, err2)

			var err3 error

			err3 = fBufReadWrite.reader.Close(
				ePrefix.XCpy("reader-Close"))

			if err3 != nil {

				errs = append(
					errs, err3)

			}

			var err4 error

			err4 = fBufReadWrite.writer.Close(
				ePrefix.XCpy("writer-Close"))

			if err4 != nil {

				errs = append(
					errs, err4)

			}

			err = new(StrMech).ConsolidateErrors(errs)

			return totalBytesRead, totalBytesWritten, err
		}

		if readErr == io.EOF {

			break
		}

	}

	return totalBytesRead, totalBytesWritten, err
}

type fileBufferReadWriteMicrobot struct {
	lock *sync.Mutex
}

// setFileMgrs
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
//		'read' buffer created for the io.Writer
//		associated with the file identified by
//		'writerPathFileName'.
//
//		'writerFileMgr' should be configured to maximize
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
//		write operations. If the target write file
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
func (fBufReadWriteMicrobot *fileBufferReadWriteMicrobot) setFileMgrs(
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
	errPrefDto *ePref.ErrPrefixDto) error {

	if fBufReadWriteMicrobot.lock == nil {
		fBufReadWriteMicrobot.lock = new(sync.Mutex)
	}

	fBufReadWriteMicrobot.lock.Lock()

	defer fBufReadWriteMicrobot.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileBufferReadWriteMicrobot." +
		"setFileMgrs()"

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

		return err
	}

	var fBuffReadWriteMolecule = new(fileBufferReadWriteMolecule)

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

		return err
	}

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

	return err
}

type fileBufferReadWriteNanobot struct {
	lock *sync.Mutex
}

// closeFileBufferReadWrite
//
// This method is designed to perform clean up tasks
// after completion of all 'read' and 'write' operations
// associated with an instance of FileBufferReadWrite
// passed as input parameter 'fBufReadWrite'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will effectively render the instance of
//	FileBufferReadWrite, passed as input parameter
//	'fBufReadWrite', invalid and unusable for any
//	future 'read' or 'write' operations.
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
//		instance of FileBufferReadWrite will be deleted
//		as part of this 'close' operation.
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
func (fBufReadWriteNanobot *fileBufferReadWriteNanobot) closeFileBufferReadWrite(
	fBufReadWrite *FileBufferReadWrite,
	fBufReadWriteLabel string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fBufReadWriteNanobot.lock == nil {
		fBufReadWriteNanobot.lock = new(sync.Mutex)
	}

	fBufReadWriteNanobot.lock.Lock()

	defer fBufReadWriteNanobot.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileBufferReadWriteNanobot." +
		"closeFileBufferReadWrite()"

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

	fBuffReadWriteAtom := new(fileBufferReadWriteElectron)

	err = fBuffReadWriteAtom.closeReader(
		fBufReadWrite,
		fBufReadWriteLabel,
		ePrefix.XCpy("Close-Reader"))

	if err != nil {

		return err
	}

	err = fBuffReadWriteAtom.closeWriter(
		fBufReadWrite,
		fBufReadWriteLabel,
		ePrefix.XCpy("Close-Writer"))

	return err
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
//		in addition to calling FileBufferReadWrite.Close().
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
//		FileBufferReadWrite.Close().
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

// setPathFileNames
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
//		'read' buffer created for the io.Writer
//		associated with the file identified by
//		'writerPathFileName'.
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
func (fBufReadWriteNanobot *fileBufferReadWriteNanobot) setPathFileNames(
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
	errPrefDto *ePref.ErrPrefixDto) error {

	if fBufReadWriteNanobot.lock == nil {
		fBufReadWriteNanobot.lock = new(sync.Mutex)
	}

	fBufReadWriteNanobot.lock.Lock()

	defer fBufReadWriteNanobot.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileBufferReadWriteNanobot." +
		"setPathFileNames()"

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
		setPathFileNameReader(
			fBufReadWrite,
			fBufReadWriteLabel,
			readerPathFileName,
			readerPathFileNameLabel,
			openReadFileReadWrite,
			readerBuffSize,
			ePrefix)

	if err != nil {

		return err
	}

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

	return err
}

type fileBufferReadWriteMolecule struct {
	lock *sync.Mutex
}

func (fBuffReadWriteMolecule *fileBufferReadWriteMolecule) setFileMgrReader(
	fBufReadWrite *FileBufferReadWrite,
	fBufReadWriteLabel string,
	readerFileMgr *FileMgr,
	readerFileMgrLabel string,
	openReadFileReadWrite bool,
	readerBuffSize int,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fBuffReadWriteMolecule.lock == nil {
		fBuffReadWriteMolecule.lock = new(sync.Mutex)
	}

	fBuffReadWriteMolecule.lock.Lock()

	defer fBuffReadWriteMolecule.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileBufferReadWriteMolecule." +
		"setFileMgrReader()"

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

		return err
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

		return err
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

		return err
	}

	readerFileMgrLabel += ".absolutePathFileName"

	err = new(fileBufferReadWriteAtom).
		setPathFileNameReader(
			fBufReadWrite,
			fBufReadWriteLabel,
			readerFileMgr.absolutePathFileName,
			readerFileMgrLabel,
			openReadFileReadWrite,
			readerBuffSize,
			ePrefix)

	return err
}

func (fBuffReadWriteMolecule *fileBufferReadWriteMolecule) setFileMgrWriter(
	fBufReadWrite *FileBufferReadWrite,
	fBufReadWriteLabel string,
	writerFileMgr *FileMgr,
	writerFileMgrLabel string,
	openWriteFileReadWrite bool,
	writerBuffSize int,
	truncateExistingWriteFile bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fBuffReadWriteMolecule.lock == nil {
		fBuffReadWriteMolecule.lock = new(sync.Mutex)
	}

	fBuffReadWriteMolecule.lock.Lock()

	defer fBuffReadWriteMolecule.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileBufferReadWriteMolecule." +
		"setFileMgrWriter()"

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

		return err
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

		return err
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

		return err
	}

	writerFileMgrLabel += ".absolutePathFileName"

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

	return err
}

type fileBufferReadWriteAtom struct {
	lock *sync.Mutex
}

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

	err = new(fileBufferReadWriteElectron).closeWriter(
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

func (fBuffReadWriteAtom *fileBufferReadWriteAtom) setPathFileNameReader(
	fBufReadWrite *FileBufferReadWrite,
	fBufReadWriteLabel string,
	readerPathFileName string,
	readerPathFileNameLabel string,
	openReadFileReadWrite bool,
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
		"setPathFileNameReader()"

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

		return err
	}

	if len(readerPathFileName) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is invalid!\n"+
			"'%v' is an empty string with a length of zero (0).\n",
			ePrefix.String(),
			readerPathFileNameLabel,
			readerPathFileNameLabel)

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

		return err
	}

	fBufReadWrite.reader = &newBuffReader
	fBufReadWrite.readerFilePathName = readerPathFileName

	return err
}

func (fBuffReadWriteAtom *fileBufferReadWriteAtom) setPathFileNameWriter(
	fBufReadWrite *FileBufferReadWrite,
	fBufReadWriteLabel string,
	writerPathFileName string,
	writerPathFileNameLabel string,
	openWriteFileReadWrite bool,
	writerBuffSize int,
	truncateExistingWriteFile bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fBuffReadWriteAtom.lock == nil {
		fBuffReadWriteAtom.lock = new(sync.Mutex)
	}

	fBuffReadWriteAtom.lock.Lock()

	defer fBuffReadWriteAtom.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileBufferReadWriteAtom." +
		"setPathFileNameWriter()"

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

		return err
	}

	if len(writerPathFileName) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is invalid!\n"+
			"'%v' is an empty string with a length of zero (0).\n",
			ePrefix.String(),
			writerPathFileNameLabel,
			writerPathFileNameLabel)

		return err
	}

	err = new(fileBufferReadWriteElectron).closeWriter(
		fBufReadWrite,
		fBufReadWriteLabel,
		ePrefix.XCpy("Close-Writer"))

	if err != nil {

		return err
	}

	var newBuffWriter FileBufferWriter
	var err2 error

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

		return err
	}

	fBufReadWrite.writer = &newBuffWriter
	fBufReadWrite.writerFilePathName = writerPathFileName

	return err
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
//		will be invalid and unusable with respect to
//		'read' operations.
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

// closeWriter
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
//	future 'write' operations.
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
//		will be invalid and unusable with respect to
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
func (fBuffReadWriteElectron *fileBufferReadWriteElectron) closeWriter(
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
		"closeWriter()"

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

		err2 = new(fileBufferWriterMolecule).close(
			fBufReadWrite.writer,
			fBufReadWriteLabel+".writer",
			ePrefix.XCpy(
				fBufReadWriteLabel+".writer"))

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"An error occurred while closing %v.writer.\n"+
				"Error=\n%v\n",
				funcName,
				fBufReadWriteLabel,
				err2.Error())

			return err
		}

	}

	fBufReadWrite.writer = nil
	fBufReadWrite.writerFilePathName = ""

	return err
}
