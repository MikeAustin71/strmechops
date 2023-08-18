package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"os"
	"strings"
	"sync"
)

// FileIoReader
//
// This type serves as a wrapper for io.Reader. As such,
// it is designed to facilitate data 'read' operations.
// The most common data source for these read operations
// is typically a data file residing on an attached
// storage drive. However, any object implementing the
// io.Reader interface may be used as a data source.
//
// The methods associated with this type do NOT employ
// buffered read techniques. Instead, direct data reads
// are performed by means of the internal io.Reader
// object.
//
// For more information on 'buffered' data reads, see
// type 'FileBufferReader'
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://pkg.go.dev/io#Reader
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	When all read operations have been completed and
//	there is no further need for the FileIoReader
//	instance, the user is responsible for 'closing' and
//	releasing the associated memory resources.
//
//	For FileIoReader instances created with a Path and
//	File Name or a File Manager (FileMgr), the user must
//	call the local method FileIoReader.Close().
//
//	For FileIoReader instances created with an external
//	io.Reader object, the user need to apply any
//	required 'close' or clean-up operations externally.
type FileIoReader struct {
	ioReader           *io.Reader
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
// FileIoReader will be unusable and should be
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
func (fIoReader *FileIoReader) Close() error {

	if fIoReader.lock == nil {
		fIoReader.lock = new(sync.Mutex)
	}

	fIoReader.lock.Lock()

	defer fIoReader.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"FileIoReader."+
			"Close()",
		"")

	if err != nil {
		return err
	}

	err = new(fileIoReaderMolecule).close(
		fIoReader,
		"fIoReader",
		ePrefix.XCpy("fIoReader"))

	return err
}

// NewIoReader
//
// Receives input parameter, 'reader', which implements the
// io.Reader interface.
//
// The io.Reader object is used to configure and return a
// fully initialized instance of FileIoReader.
//
// The returned instance of FileIoReader may be used
// to read data from the data source identified by the
// io.Reader object, 'reader'.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	The returned instance of FileIoReader does NOT use
//	buffered read techniques. Instead, it implements a
//	direct read protocol using io.Reader.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://pkg.go.dev/io#Reader
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
//		in addition to calling FileIoReader.Close().
//
//		While the returned instance of FileIoReader
//		is primarily designed for reading data from disk
//		files, this 'reader' will in fact read data from
//		any object implementing the io.Reader interface.
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
//	FileIoReader
//
//		If this method completes successfully, a fully
//		configured instance of FileIoReader will
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
func (fIoReader *FileIoReader) NewIoReader(
	reader io.Reader,
	errorPrefix interface{}) (
	FileIoReader,
	error) {

	if fIoReader.lock == nil {
		fIoReader.lock = new(sync.Mutex)
	}

	fIoReader.lock.Lock()

	defer fIoReader.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error
	var newFileIoReader FileIoReader

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileIoReader."+
			"NewIoReader()",
		"")

	if err != nil {
		return newFileIoReader, err
	}

	err = new(fileIoReaderNanobot).
		setIoReader(
			&newFileIoReader,
			"newFileIoReader",
			reader,
			"reader",
			ePrefix.XCpy("newFileIoReader"))

	return newFileIoReader, err
}

// NewFileMgr
//
// Receives an instance of FileMgr as input parameter
// 'fileMgr' and returns a new, fully configured instance
// of FileIoReader.
//
// The FileIoReader instance returned by this method will
// configure the file identified by 'fileMgr' as the data
// source for file 'read' operations.
//
// This target 'read' file identified by 'fileMgr' is
// opened for either 'read-only' or 'read/write'
// operations depending on input parameter
// 'openFileReadWrite'.
//
// If the target path and file identified by 'fileMgr' do
// not currently exist on an attached storage drive, an
// error will be returned.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://pkg.go.dev/io#Reader
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	As a precaution, the incoming 'fileMgr' object
//		will be closed before configuring 'fIoReader'
//		with a new internal io.Reader object.
//
//	(2)	The returned instance of FileIoReader does NOT
//		use buffered read techniques. Instead, it
//		implements a direct read protocol using
//		io.Reader.
//
//	(3)	When all read operations have been completed and
//		there is no further need for the returned
//		instance of FileIoReader, the user is responsible
//		for 'closing' and releasing the associated memory
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
//		The instance of FileIoReader returned by this
//		method will configure the file identified by
//		'fileMgr' as the data source for file 'read'
//		operations.
//
//		As a precaution, the incoming 'fileMgr' object
//		will be closed before configuring the returned
//		instance of 'FileIoReader' with a new internal
//		io.Reader object.
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
//	newFileIoReader			FileIoReader
//
//		If this method completes successfully, a fully
//		configured instance of FileIoReader will
//		be returned.
//
//		This returned instance of FileIoReader will
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
func (fIoReader *FileIoReader) NewFileMgr(
	fileMgr *FileMgr,
	openFileReadWrite bool,
	errorPrefix interface{}) (
	fInfoPlus FileInfoPlus,
	newFileIoReader FileIoReader,
	err error) {

	if fIoReader.lock == nil {
		fIoReader.lock = new(sync.Mutex)
	}

	fIoReader.lock.Lock()

	defer fIoReader.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileIoReader."+
			"NewFileMgr()",
		"")

	if err != nil {

		return fInfoPlus, newFileIoReader, err
	}

	fInfoPlus,
		err = new(fileIoReaderMicrobot).
		setFileMgr(
			&newFileIoReader,
			"newFileIoReader",
			fileMgr,
			"fileMgr",
			openFileReadWrite,
			ePrefix.XCpy(
				"fileMgr"))

	return fInfoPlus, newFileIoReader, err
}

// NewPathFileName
//
// Receives a path and file name as an input parameter
// string, 'pathFileName' and returns a new, fully
// configured instance of FileIoReader.
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
// If the target path and file identified by
// 'pathFileName' do not currently exist on an attached
// storage drive, an error will be returned.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://pkg.go.dev/io#Reader
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	The returned instance of FileIoReader does NOT
//		use buffered read techniques. Instead, it
//		implements a direct read protocol using
//		io.Reader.
//
//	(2)	When all read operations have been completed and
//		there is no further need for the returned
//		instance of FileIoReader, the user is responsible
//		for 'closing' and releasing the associated memory
//		resources by calling the local method
//		FileIoReader.Close().
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathFileName				string
//
//		This string contains the path and file name of
//		the file which will be used a data source for
//		'read' operations.
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
//	newFileIoReader			FileIoReader
//
//		If this method completes successfully, a fully
//		configured instance of FileIoReader will
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
func (fIoReader *FileIoReader) NewPathFileName(
	pathFileName string,
	openFileReadWrite bool,
	errorPrefix interface{}) (
	fInfoPlus FileInfoPlus,
	newFileIoReader FileIoReader,
	err error) {

	if fIoReader.lock == nil {
		fIoReader.lock = new(sync.Mutex)
	}

	fIoReader.lock.Lock()

	defer fIoReader.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileIoReader."+
			"NewPathFileName()",
		"")

	if err != nil {

		return fInfoPlus, newFileIoReader, err
	}

	fInfoPlus,
		err = new(fileIoReaderNanobot).
		setPathFileName(
			&newFileIoReader,
			"newFileIoReader",
			pathFileName,
			"pathFileName",
			openFileReadWrite,
			ePrefix.XCpy(
				pathFileName))

	return fInfoPlus, newFileIoReader, err
}

// Read
//
// Reads a selection data from the pre-configured
// io.Reader data source encapsulated in the current
// instance of FileIoReader.
//
// This method is a wrapper for the 'io.Reader.Read'
// method.
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
// See the io.Reader docs and 'Reference' section below.
//
// Once the 'read' operation has been completed, the user
// MUST call the 'Close' method to ensure clean-up
// operations are properly applied:
//
//	FileIoReader.Close()
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	This method implements the io.Reader interface.
//
//	(2)	Keep calling this method until all the bytes have
//		been read from the data source configured for the
//		current instance of FileIoReader and the
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
//			FileIoReader.Close()
//
//	(5)	This method employs the direct 'read' technique
//		used by io.Reader. It does NOT use the buffered
//		'read' technique used by bufio.Read.
//
// ----------------------------------------------------------------
//
// # Reference:
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
//		FileIoReader.
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
func (fIoReader *FileIoReader) Read(
	bytesRead []byte) (
	numOfBytesRead int,
	err error) {

	if fIoReader.lock == nil {
		fIoReader.lock = new(sync.Mutex)
	}

	fIoReader.lock.Lock()

	defer fIoReader.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"FileIoReader."+
			"Read()",
		"")

	if err != nil {

		return numOfBytesRead, err
	}

	if fIoReader.ioReader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of 'FileIoReader' is invalid!\n"+
			"The internal io.Reader object has NOT been initialized.\n"+
			"Call one of the 'New' or 'Setter' methods when creating\n"+
			"an instance of 'FileIoReader'\n",
			ePrefix.String())

		return numOfBytesRead, err
	}

	var err2 error
	var reader = *fIoReader.ioReader

	numOfBytesRead,
		err2 = reader.Read(bytesRead)

	if err2 != nil {

		if err2 != io.EOF {

			err = fmt.Errorf("%v\n"+
				"Error returned by fIoReader.ioReader.Read(bytesRead).\n"+
				"Error=\n%v\n",
				ePrefix.String(),
				err2.Error())

			return numOfBytesRead, err

		} else {

			err = io.EOF

		}

	}

	return numOfBytesRead, err
}

// ReadAllTextLines
//
// Reads text lines from the internal io.Reader object
// encapsulated in the current instance of FileIoReader.
// The entire contents of the io.Reader object are
// parsed and stored as individual lines of text in
// the instance of StringArrayDto passed as input
// parameter 'outputLinesArray'.
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
// StringArrayDto instance 'outputLinesArray'. Depending
// on the size of the target 'read' file, local memory
// constraints should be considered.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	This method is designed to read the entire
//		contents of the internal io.Reader object,
//		encapsulated by the current instance of
//		FileIoReader, into memory.
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
//			FileIoReader.Close()
//
//	(3)	If the current instance of FileIoReader has
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
//		will be read from the internal io.Reader
//		encapsulated by the current instance of
//		FileIoReader.
//
//		If 'maxNumOfLines' is set to a value less than
//		zero (0) (Example: minus-one (-1) ),
//		'maxNumOfLines' will be automatically reset to
//		math.MaxInt(). This means all text lines existing
//		in the internal io.Reader object will be read
//		and processed. Reading all the text lines in a
//		file 'may' have memory implications depending on
//		the size of the file and the memory resources
//		available to your computer.
//
//		If 'maxNumOfLines' is set to a value of zero
//		('0'), no text lines will be read from
//		the internal io.Reader, and no error will be
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
//		Lines of text read from the internal io.Reader
//		object configured for the current instance of
//		FileIoReader will be stored as individual
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
//		read from the internal io.Reader object
//		encapsulated by the current instance of
//		FileIoReader.
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
func (fIoReader *FileIoReader) ReadAllTextLines(
	maxNumOfLines int,
	endOfLineDelimiters *StringArrayDto,
	outputLinesArray *StringArrayDto,
	errorPrefix interface{}) (
	numOfLinesRead int,
	numOfBytesRead int64,
	err error) {

	if fIoReader.lock == nil {
		fIoReader.lock = new(sync.Mutex)
	}

	fIoReader.lock.Lock()

	defer fIoReader.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileIoReader."+
			"ReadAllTextLines()",
		"")

	if err != nil {

		return numOfLinesRead,
			numOfBytesRead,
			err
	}

	if fIoReader.ioReader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of 'FileIoReader' is invalid!\n"+
			"The internal io.Reader object has NOT been initialized.\n"+
			"Call one of the 'New' or 'Setter' methods when creating\n"+
			"an instance of 'FileIoReader'\n",
			ePrefix.String())

		return numOfLinesRead,
			numOfBytesRead,
			err
	}

	numOfLinesRead,
		numOfBytesRead,
		err = new(fileHelperMolecule).
		readerScanLines(
			*fIoReader.ioReader,
			"fIoReader.fileReader",
			maxNumOfLines,
			endOfLineDelimiters,
			outputLinesArray,
			ePrefix.XCpy("fIoReader.fileReader"))

	return numOfLinesRead,
		numOfBytesRead,
		err
}

// ReadAllStrBuilder
//
// Reads the entire contents of the internal io.Reader
// for the current instance of FileIoReader as
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
// entire contents of the target io.Reader object into
// memory when writing said contents to the
// strings.Builder instance 'strBuilder'. Depending
// on the size of the target 'read' file, local memory
// constraints should be considered.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	This method is designed to read the entire
//		contents of the internal io.Reader object,
//		encapsulated by the current instance of
//		FileIoReader, into memory.
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
//			FileIoReader.Close()
//
//	(3)	If the current instance of FileIoReader has
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
//		entire contents of the internal io.Reader for the
//		current instance of FileIoReader and stores the
//		resulting string in 'strBuilder'.
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
//		read from the internal io.Reader object
//		encapsulated by the current instance of
//		FileIoReader.
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
func (fIoReader *FileIoReader) ReadAllStrBuilder(
	strBuilder *strings.Builder,
	errorPrefix interface{}) (
	numOfBytesRead int64,
	err error) {

	if fIoReader.lock == nil {
		fIoReader.lock = new(sync.Mutex)
	}

	fIoReader.lock.Lock()

	defer fIoReader.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileIoReader."+
			"ReadAllStrBuilder()",
		"")

	if err != nil {

		return numOfBytesRead, err
	}

	numOfBytesRead,
		err = new(fileIoReaderMicrobot).
		readAllStrBuilder(
			fIoReader,
			"fIoReader",
			strBuilder,
			ePrefix)

	return numOfBytesRead, err
}

// ReadAllToString
//
// Reads the entire contents of the internal io.Reader
// for the current instance of FileIoReader and returns
// these contents as a single string ('contentsStr').
//
// If a processing error is encountered, an appropriate
// error with an error message will be returned. When
// the end-of-file is encountered during the 'read'
// process, the returned error object will be set to
// 'nil' and no error will be returned.
//
// It naturally follows that this method will read the
// entire contents of the target io.Reader object into
// memory when writing said contents to the returned
// string parameter 'contentsStr'. Depending on the size
// of the target 'read' file, local memory constraints
// should be considered.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	This method is designed to read the entire
//		contents of the internal io.Reader object,
//		encapsulated by the current instance of
//		FileIoReader, into memory.
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
//			FileIoReader.Close()
//
//	(3)	If the current instance of FileIoReader has
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
//		read from the internal io.Reader object
//		encapsulated by the current instance of
//		FileIoReader.
//
//		This integer value should also equal the
//		string length of the returned string,
//		'contentsStr'.
//
//	contentsStr					string
//
//		If this method completes successfully, the entire
//		contents if the internal io.Reader object
//		encapsulated by the current instance of
//		FileIoReader will be returned in this string.
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
func (fIoReader *FileIoReader) ReadAllToString(
	errorPrefix interface{}) (
	numOfBytesRead int64,
	contentsStr string,
	err error) {

	if fIoReader.lock == nil {
		fIoReader.lock = new(sync.Mutex)
	}

	fIoReader.lock.Lock()

	defer fIoReader.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileIoReader."+
			"ReadAllToString()",
		"")

	if err != nil {

		return numOfBytesRead, contentsStr, err
	}

	strBuilder := new(strings.Builder)

	numOfBytesRead,
		err = new(fileIoReaderMicrobot).
		readAllStrBuilder(
			fIoReader,
			"fIoReader",
			strBuilder,
			ePrefix)

	if err == nil {

		contentsStr = strBuilder.String()
	}

	return numOfBytesRead, contentsStr, err
}

// SetIoReader
//
// This method will completely re-initialize the current
// instance of FileIoReader using the io.Reader object
// passed as input parameter 'reader'.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://pkg.go.dev/io#Reader
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	This method will delete, overwrite and reset all
//		pre-existing data values in the current instance
//		of FileIoReader.
//
//	(2)	The user is responsible for 'closing' the
//		instance of io.Reader passed as input parameter
//		'reader'. The FileIoReader.Close() method will
//		NOT close the 'reader' object.
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
//		in addition to calling FileIoReader.Close().
//
//		While the returned instance of FileIoReader
//		is primarily designed for reading data from disk
//		files, this 'reader' will in fact read data from
//		any object implementing the io.Reader interface.
//
//		Remember that the user is responsible for
//		'closing' this io.Reader object. The FileIoReader
//		method 'Close()' will NOT close this io.Reader
//		object.
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
func (fIoReader *FileIoReader) SetIoReader(
	reader io.Reader,
	errorPrefix interface{}) error {

	if fIoReader.lock == nil {
		fIoReader.lock = new(sync.Mutex)
	}

	fIoReader.lock.Lock()

	defer fIoReader.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileIoReader."+
			"SetIoReader()",
		"")

	if err != nil {
		return err
	}

	err = new(fileIoReaderNanobot).
		setIoReader(
			fIoReader,
			"fIoReader",
			reader,
			"reader",
			ePrefix.XCpy("fIoReader"))

	return err
}

// SetFileMgr
//
// This method will completely re-initialize the current
// instance of FileIoReader using the path and file
// name identified by the FileMgr instance passed as
// input parameter 'fileMgr'.
//
// The file identified by 'fileMgr' will be used to
// reconfigure the internal bufio.Reader encapsulated in
// the current instance of FileIoReader.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://pkg.go.dev/io#Reader
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	This method will delete, overwrite and reset all
//		pre-existing data values in the current instance
//		of FileIoReader.
//
//	(2)	As a precaution, the incoming 'fileMgr' object
//		will be closed before configuring the current
//		FileIoReader instance with a new, internal
//		io.Reader object.
//
//	(3)	When all read operations have been completed and
//		there is no further need for the current instance
//		of FileIoReader, the user is responsible for
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
//		The instance of FileIoReader returned by this
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
func (fIoReader *FileIoReader) SetFileMgr(
	fileMgr *FileMgr,
	openFileReadWrite bool,
	errorPrefix interface{}) (
	fInfoPlus FileInfoPlus,
	err error) {

	if fIoReader.lock == nil {
		fIoReader.lock = new(sync.Mutex)
	}

	fIoReader.lock.Lock()

	defer fIoReader.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileIoReader."+
			"SetFileMgr()",
		"")

	if err != nil {

		return fInfoPlus, err
	}

	fInfoPlus,
		err = new(fileIoReaderMicrobot).
		setFileMgr(
			fIoReader,
			"fIoReader",
			fileMgr,
			"fileMgr",
			openFileReadWrite,
			ePrefix.XCpy(
				"fileMgr"))

	return fInfoPlus, err
}

// SetPathFileName
//
// This method will completely re-initialize the current
// instance of FileIoReader using the path and file
// name passed as input parameter 'pathFileName'.
//
// The path and file name specified by 'pathFileName'
// will be used to reconfigure the internal bufio.Reader
// encapsulated in the current instance of
// FileIoReader.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://pkg.go.dev/io#Reader
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	This method will delete, overwrite and reset all
//		pre-existing data values in the current instance
//		of FileIoReader.
//
//	(2)	When all read operations have been completed and
//		there is no further need for the current instance
//		of FileIoReader, the user is responsible for
//		'closing' and releasing the associated memory
//		resources by calling the local method
//		FileIoReader.Close().
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
func (fIoReader *FileIoReader) SetPathFileName(
	pathFileName string,
	openFileReadWrite bool,
	errorPrefix interface{}) (
	fInfoPlus FileInfoPlus,
	err error) {

	if fIoReader.lock == nil {
		fIoReader.lock = new(sync.Mutex)
	}

	fIoReader.lock.Lock()

	defer fIoReader.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileIoReader."+
			"SetPathFileName()",
		"")

	if err != nil {

		return fInfoPlus, err
	}

	fInfoPlus,
		err = new(fileIoReaderNanobot).
		setPathFileName(
			fIoReader,
			"fIoReader",
			pathFileName,
			"pathFileName",
			openFileReadWrite,
			ePrefix.XCpy(
				pathFileName))

	return fInfoPlus, err
}

type fileIoReaderMicrobot struct {
	lock *sync.Mutex
}

// readAllStrBuilder
//
// Reads the entire contents of the internal io.Reader
// for instance of FileIoReader passed as input parameter
// 'fIoReader'. These contents are converted to a string
// which is then stored and returned through an instance
// of strings.Builder passed as input parameter
// 'strBuilder'.
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
//	fIoReader					*FileIoReader
//
//		A pointer to an instance of FileIoReader.
//
//		The entire contents of the io.Reader object
//		encapsulated in this FileIoReader instance
//		will be extracted and returned as a string
//		through input parameter 'strBuilder'.
//
//	fIoReaderLabel				string
//
//		The name or label associated with input parameter
//		'fIoReader' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fIoReader" will be
//		automatically applied.
//
//	strBuilder					*strings.Builder
//
//		A pointer to an instance of strings.Builder. The
//		entire contents of the internal io.Reader for the
//		FileIoReader instance passed as 'fIoReader' will
//		be extracted and stored as a string in
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
//		read from the internal io.Reader object
//		encapsulated by the FileIoReader instance passed
//		as input parameter 'fIoReader'.
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
func (fIoReaderMicrobot *fileIoReaderMicrobot) readAllStrBuilder(
	fIoReader *FileIoReader,
	fIoReaderLabel string,
	strBuilder *strings.Builder,
	errPrefDto *ePref.ErrPrefixDto) (
	numOfBytesRead int64,
	err error) {

	if fIoReaderMicrobot.lock == nil {
		fIoReaderMicrobot.lock = new(sync.Mutex)
	}

	fIoReaderMicrobot.lock.Lock()

	defer fIoReaderMicrobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileIoReaderMicrobot."+
			"readAllStrBuilder()",
		"")

	if err != nil {

		return numOfBytesRead, err
	}

	if len(fIoReaderLabel) == 0 {

		fIoReaderLabel = "fIoReader"
	}

	if strBuilder == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'strBuilder' is invalid!\n"+
			"'strBuilder' is a 'nil' pointer.\n",
			ePrefix)

		return numOfBytesRead, err
	}

	if fIoReader.ioReader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of 'FileIoReader' (%v) is invalid!\n"+
			"The internal '%v' io.Reader object has NOT been initialized.\n"+
			"Call one of the 'New' or 'Setter' methods when creating\n"+
			"an instance of 'FileIoReader'\n",
			ePrefix.String(),
			fIoReaderLabel,
			fIoReaderLabel)

		return numOfBytesRead, err
	}

	var reader = *fIoReader.ioReader
	var bytesRead = make([]byte, 4096)
	var localBytesRead int
	var err2 error

	for {

		localBytesRead,
			err2 = reader.Read(bytesRead)

		if localBytesRead > 0 {

			strBuilder.Write(bytesRead[:localBytesRead])

			numOfBytesRead += int64(localBytesRead)

		}

		if err2 == io.EOF {

			break

		}

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Error returned by fIoReader.ioReader.Read(bytesRead).\n"+
				"Error=\n%v\n",
				ePrefix.String(),
				err2.Error())

			break
		}

	}

	return numOfBytesRead, err
}

// setFileMgr
//
// This 'setter' method is used to initialize new values
// for internal member variables contained in the
// instance of FileIoReader passed as input parameter
// 'fIoReader'.
//
// The new io.Reader object assigned to 'fIoReader' is
// generated from the File Manager (FileMgr) instance
// passed as input parameter 'fileMgr'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1) This method will delete, overwrite and reset all
//		pre-existing data values in the instance of
//		FileIoReader passed as input parameter
//		'fIoReader'.
//
//	(2) As a precaution, the incoming 'fileMgr' object
//		will be closed before configuring 'fIoReader'
//		with a new internal io.Reader object.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fIoReader					*FileIoReader
//
//		A pointer to an instance of FileIoReader.
//
//		All internal member variable data values in
//		this instance will be deleted and initialized
//		to new values based on the following input
//		parameters.
//
//	fIoReaderLabel				string
//
//		The name or label associated with input parameter
//		'fIoReader' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fIoReader" will be
//		automatically applied.
//
//	fileMgr						*FileMgr
//
//		A pointer to an instance of FileMgr. The file
//		identified by 'fileMgr' will be used as a
//		data source for 'read' operations.
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
func (fIoReaderMicrobot *fileIoReaderMicrobot) setFileMgr(
	fIoReader *FileIoReader,
	fIoReaderLabel string,
	fileMgr *FileMgr,
	fileMgrLabel string,
	openFileReadWrite bool,
	errPrefDto *ePref.ErrPrefixDto) (
	fInfoPlus FileInfoPlus,
	err error) {

	if fIoReaderMicrobot.lock == nil {
		fIoReaderMicrobot.lock = new(sync.Mutex)
	}

	fIoReaderMicrobot.lock.Lock()

	defer fIoReaderMicrobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileIoReaderMicrobot." +
		"setFileMgr()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return fInfoPlus, err
	}

	if len(fIoReaderLabel) == 0 {

		fIoReaderLabel = "fIoReader"
	}

	if len(fileMgrLabel) == 0 {

		fileMgrLabel = "fileMgr"
	}

	if fIoReader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: The FileIoReader instance passed\n"+
			"as input parameter '%v' is invalid!\n"+
			"'%v' is a 'nil' pointer.\n",
			ePrefix.String(),
			fIoReaderLabel,
			fIoReaderLabel)

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
		err = new(fileIoReaderNanobot).
		setPathFileName(
			fIoReader,
			fIoReaderLabel,
			fileMgr.absolutePathFileName,
			fileMgrLabel,
			openFileReadWrite,
			ePrefix.XCpy(fIoReaderLabel+"<-"+fileMgrLabel))

	return fInfoPlus, err
}

type fileIoReaderNanobot struct {
	lock *sync.Mutex
}

// setIoReader
//
// This 'setter' method is used to initialize new values
// for internal member variables contained in the
// instance of FileIoReader passed as input parameter
// 'fIoReader'. The new configuration will be based on
// an io.Reader object passed as input parameter
// 'reader'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the instance of
//	FileIoReader passed as input parameter
//	'fIoReader'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fIoReader					*FileIoReader
//
//		A pointer to an instance of FileIoReader.
//
//		All internal member variable data values in
//		this instance will be deleted and initialized
//		to new values based on the following input
//		parameters.
//
//	fIoReaderLabel				string
//
//		The name or label associated with input parameter
//		'fIoReader' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fIoReader" will be
//		automatically applied.
//
//	reader						io.Reader
//
//		This parameter will accept any object
//		implementing the io.Reader interface.
//
//		This instance of io.Reader will be used to
//		configure the internal bufio.Reader contained in
//		'fIoReader' and used to conduct 'read'
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
func (fIoReaderNanobot *fileIoReaderNanobot) setIoReader(
	fIoReader *FileIoReader,
	fIoReaderLabel string,
	reader io.Reader,
	readerLabel string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fIoReaderNanobot.lock == nil {
		fIoReaderNanobot.lock = new(sync.Mutex)
	}

	fIoReaderNanobot.lock.Lock()

	defer fIoReaderNanobot.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileIoReaderNanobot." +
		"setIoReader()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if len(fIoReaderLabel) == 0 {

		fIoReaderLabel = "fIoReader"
	}

	if len(readerLabel) == 0 {

		readerLabel = "reader"
	}

	if fIoReader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: The FileIoReader instance passed\n"+
			"as input parameter '%v' is invalid!\n"+
			"'%v' is a 'nil' pointer.\n",
			ePrefix.String(),
			fIoReaderLabel,
			fIoReaderLabel)

		return err
	}

	err = new(fileIoReaderMolecule).close(
		fIoReader,
		fIoReaderLabel,
		ePrefix.XCpy(fIoReaderLabel))

	if err != nil {
		return err
	}

	fIoReader.ioReader = &reader

	return err
}

// setPathFileName
//
// This 'setter' method is used to initialize new values
// for internal member variables contained in the
// instance of FileIoReader passed as input parameter
// 'fIoReader'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the instance of
//	FileIoReader passed as input parameter
//	'fIoReader'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fIoReader					*FileIoReader
//
//		A pointer to an instance of FileIoReader.
//
//		All internal member variable data values in
//		this instance will be deleted and initialized
//		to new values based on the following input
//		parameters.
//
//	fIoReaderLabel				string
//
//		The name or label associated with input parameter
//		'fIoReader' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fIoReader" will be
//		automatically applied.
//
//	pathFileName				string
//
//		This string contains the path and file name of
//		the file which will be used a data source for
//		'read' operations.
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
func (fIoReaderNanobot *fileIoReaderNanobot) setPathFileName(
	fIoReader *FileIoReader,
	fIoReaderLabel string,
	pathFileName string,
	pathFileNameLabel string,
	openFileReadWrite bool,
	errPrefDto *ePref.ErrPrefixDto) (
	fInfoPlus FileInfoPlus,
	err error) {

	if fIoReaderNanobot.lock == nil {
		fIoReaderNanobot.lock = new(sync.Mutex)
	}

	fIoReaderNanobot.lock.Lock()

	defer fIoReaderNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileIoReaderNanobot." +
		"setPathFileName()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		return fInfoPlus, err
	}

	if len(fIoReaderLabel) == 0 {

		fIoReaderLabel = "fIoReader"
	}

	if fIoReader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: The FileIoReader instance passed\n"+
			"as input parameter '%v' is invalid!\n"+
			"'%v' is a 'nil' pointer.\n",
			ePrefix.String(),
			fIoReaderLabel,
			fIoReaderLabel)

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

	err = new(fileIoReaderMolecule).close(
		fIoReader,
		fIoReaderLabel,
		ePrefix.XCpy(fIoReaderLabel))

	if err != nil {

		return fInfoPlus, err
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

	fIoReader.filePtr,
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

	fIoReader.targetReadFileName = pathFileName

	var reader io.Reader

	reader = fIoReader.filePtr

	fIoReader.ioReader = &reader

	return fInfoPlus, err
}

type fileIoReaderMolecule struct {
	lock *sync.Mutex
}

// close
//
// This method will effectively close and invalidate the
// instance of FileIoReader passed as input parameter
// 'fIoReader'.
//
// If a file pointer (*os.File) was previously configured
// for 'fIoReader', it will be closed and set to 'nil'.
//
// After completion of this method this FileIoReader
// instance will be unusable, invalid and unavailable for
// future 'read' operations.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete all pre-existing data values
//	in the instance of FileIoReader passed as input
//	parameter 'fIoReader'.
//
//	After completion of this method this FileIoReader
//	instance will be unusable, invalid and unavailable
//	for future 'read' operations.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fIoReader					*FileIoReader
//
//		A pointer to an instance of FileIoReader.
//
//		All internal member variable data values in
//		this instance will be deleted.
//
//		If a file pointer (*os.File) was previously
//		configured for 'fIoReader', it will be closed
//		and set to 'nil' by this method.
//
//	fIoReaderLabel				string
//
//		The name or label associated with input parameter
//		'fIoReader' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fIoReader" will be
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
func (fIoReaderMolecule *fileIoReaderMolecule) close(
	fIoReader *FileIoReader,
	fIoReaderLabel string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fIoReaderMolecule.lock == nil {
		fIoReaderMolecule.lock = new(sync.Mutex)
	}

	fIoReaderMolecule.lock.Lock()

	defer fIoReaderMolecule.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileIoReaderMolecule." +
		"close()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return err
	}

	if len(fIoReaderLabel) == 0 {

		fIoReaderLabel = "fIoReader"
	}

	if fIoReader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: The FileIoReader instance passed as\n"+
			"input parameter '%v' is invalid!\n"+
			"'%v' is a 'nil' pointer.\n",
			ePrefix.String(),
			fIoReaderLabel,
			fIoReaderLabel)

		return err
	}

	if fIoReader.filePtr != nil {

		var err2 error

		err2 = fIoReader.filePtr.Close()

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Error returned while closing the target 'target' file!\n"+
				"fBufWriter.filePtr.Close()\n"+
				"Target Read File = '%v'\n"+
				"Error = \n%v\n",
				ePrefix.String(),
				fIoReader.targetReadFileName,
				err2.Error())

		}
	}

	fIoReader.targetReadFileName = ""

	fIoReader.filePtr = nil

	fIoReader.ioReader = nil

	return err
}
