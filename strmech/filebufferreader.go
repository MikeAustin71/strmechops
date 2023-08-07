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
// This structure and its associated methods are designed
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
//			responsible for clean-up of the io.Reader
//			object after FileBufferReader 'read'
//			operations have been completed.
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
// # Input Parameters
//
//	fileMgr						*FileMgr
//
//		A pointer to an instance of FileMgr. The file
//		identified by 'fileMgr' will be used as a data
//		source for 'read' operations performed by
//		method:
//
//			FileBufferReader.Read()
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
// string, 'pathFileName'.
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
// This target 'read' file identified by 'fileMgr' will
// be used to create a file pointer (*os.File) which in
// turn will be used to configure the internal
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

// Read
//
// Reads a selection data from the pre-configured
// io.Reader data source encapsulated in the current
// instance of FileBufferReader.
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
			"Call one of the 'New' or 'Setter' methods when creating\n"+
			"an instance of 'FileBufferReader'\n",
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

// ReadAllTextLines
//
// Reads text lines from the internal io.Reader object
// encapsulated in the current instance of
// FileBufferReader.
//
// Multiple custom end of line delimiters may be utilized
// to determine the end of each line of text read from
// the internal io.Reader object. End of line delimiters
// are specified by input parameter
// 'endOfLineDelimiters', an instance of StringArrayDto.
// 'endOfLineDelimiters' contains an array of strings any
// one of which may be used to identify, delimit and
// separate individual lines of text read from the target
// io.Reader object.
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
// entire contents of the target io.Reader object into
// memory when writing said contents to the
// StringArrayDto instance 'outputLinesArray'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	This method is designed to read the entire
//		contents of the internal io.Reader object,
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
//		Lines of text read from the internal io.Reader
//		object configured for the current instance of
//		FileBufferReader will be stored as individual
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
//		lines read from the file specified by input
//		parameter 'pathFileName'. This value also
//		specifies the number of array elements added to
//		the string array encapsulated by
//		'outputLinesArray'.
//
//	numBytesRead				int64
//
//		If this method completes successfully, this
//		integer value will equal the number of bytes
//		read from the target input file 'pathFileName'
//		and added to the string array encapsulated by
//		'outputLinesArray'.
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
	endOfLineDelimiters *StringArrayDto,
	outputLinesArray *StringArrayDto,
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

	if fBufReader.fileReader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of 'FileBufferReader' is invalid!\n"+
			"The internal bufio.Reader has NOT been initialized.\n"+
			"Call one of the 'New' or 'Setter' methods when creating\n"+
			"an instance of 'FileBufferReader'\n",
			ePrefix.String())

		return numOfLinesRead,
			numOfBytesRead,
			err
	}

	numOfLinesRead,
		numOfBytesRead,
		err = new(fileHelperMolecule).
		readerScanLines(
			fBufReader.fileReader,
			"fBufReader.fileReader",
			endOfLineDelimiters,
			outputLinesArray,
			ePrefix.XCpy("fBufReader.fileReader"))

	return numOfLinesRead,
		numOfBytesRead,
		err
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
//	fileMgr						*FileMgr
//
//		A pointer to an instance of FileMgr. The file
//		identified by 'fileMgr' will be used as a data
//		source for 'read' operations performed by
//		method:
//
//			FileBufferReader.Read()
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
//		the file which will be configured as a new data
//		source for 'read' operations performed by method:
//
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

type fileBufferReaderMicrobot struct {
	lock *sync.Mutex
}

// setFileMgr
//
// This 'setter' method is used to initialize new values
// for internal member variables contained in the
// instance of FileBufferReader passed as input parameter
// 'fBufReader'.
//
// The new io.Reader object assigned to 'fBufReader' is
// generated from the File Manager (FileMgr) instance
// passed as input parameter 'fileMgr'.
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
			"The path and file name do NOT exist on an attached\n"+
			"storage drive.\n"+
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

	fBufReader.fileReader = bufio.NewReaderSize(
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

	fBufReader.fileReader = nil

	return err
}
