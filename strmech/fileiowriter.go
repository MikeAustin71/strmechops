package strmech

import (
	"errors"
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"math"
	"os"
	"sync"
)

// FileIoWriter
//
// This type serves as a wrapper for io.writer. As such,
// it is designed to facilitate data 'write' operations.
// The most common data source for these 'write' operations
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
//	https://pkg.go.dev/io#Writer
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	Pointer receiver FileIoWriter methods implement
//		the following interfaces:
//
//			io.Writer
//			io.Closer
//			io.ReaderFrom
//			io.Seeker
//			io.WriteCloser
//			io.WriterAt
//			io.WriteSeeker
//
//	(2)	When all read operations have been completed and
//		there is no further need for the FileIoWriter
//		instance, the user is responsible for 'closing'
//		the underlying io.Writer object and releasing the
//		associated memory resources. This Clean-Up
//		operation is accomplished by calling local
//		method:
//
//			FileIoWriter.Close()
//
//		For FileIoWriter instances created with an
//		external io.Writer object, the user may need to
//		'close' that io.Writer object and perform
//		Clean-Up operations externally.
type FileIoWriter struct {
	ioWriter             *io.Writer
	filePtr              *os.File
	targetWriteFileName  string
	defaultByteArraySize int
	lock                 *sync.Mutex
}

// Close
//
// This method implements the io.Closer interface.
//
// FileIoWriter.Close() effectively performs all required
// Clean-Up tasks. As such, this method should only be
// called after all 'write' operations have been completed
// and the services of the current FileIoWriter instance
// are no longer required.
//
// After calling this method, FileIoWriter.Close(), the
// current instance of FileIoWriter will be invalid and
// unavailable for further 'write' operations.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	This method implements the io.Closer interface.
//
//	(2)	This method will perform all required Clean-Up
//		tasks after all 'write' operations have been
//		completed.
//
//	(3)	Clean-Up tasks performed by this method include
//		'closing' the underlying io.Writer object and
//		releasing the internal memory resources.
//
//		Releasing all internal memory resources will
//		synchronize internal flags and prevent multiple
//		calls to 'close' the underlying io.Writer object.
//		Calling 'close' on the same underlying io.Writer
//		object multiple times can produce unexpected
//		results.
//
//	(4)	After completing the call to this method,
//		FileIoWriter.Close(), the current FileIoWriter
//		instance will become invalid and unavailable
//		for future 'write' operations.
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
func (fIoWriter *FileIoWriter) Close() error {

	if fIoWriter.lock == nil {
		fIoWriter.lock = new(sync.Mutex)
	}

	fIoWriter.lock.Lock()

	defer fIoWriter.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"FileIoWriter."+
			"Close()",
		"")

	if err != nil {
		return err
	}

	err = new(fileIoWriterMolecule).
		closeAndRelease(
			fIoWriter,
			"fIoWriter",
			true, // releaseMemoryResources
			ePrefix.XCpy("fIoWriter"))

	return err
}

// Empty
//
// This method deletes all internal member variables and
// releases all the internal memory resources for the
// current instance of FileIoWriter.
//
// Specifically the following internal object pointers
// are set to nil or their initial zero values:
//
//	FileIoWriter.targetReadFileName = ""
//	FileIoWriter.filePtr = nil
//	FileIoWriter.ioReader = nil
//	FileIoWriter.defaultByteArraySize = 0
//
// After calling this method, the current instance of
// FileIoWriter will become invalid and unavailable
// for future 'write' operations.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	This method does NOT perform the 'close' protocol. To
//	perform both the 'close' protocol and release all
//	internal memory resources call local method:
//
//			FileIoWriter.Close()
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
func (fIoWriter *FileIoWriter) Empty() {

	if fIoWriter.lock == nil {
		fIoWriter.lock = new(sync.Mutex)
	}

	fIoWriter.lock.Lock()

	new(fileIoWriterAtom).empty(
		fIoWriter)

	fIoWriter.lock.Unlock()

	fIoWriter.lock = nil
}

// GetDefaultByteArraySize
//
// Returns the default byte array size for the current
// instance of FileIoWriter.
//
// Methods which utilize the Default Writer Byte Array
// Size include:
//
//	FileIoWriter.ReadFrom()
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
//		This integer value returns the default byte array
//		length for the current instance of FileIoWriter.
func (fIoWriter *FileIoWriter) GetDefaultByteArraySize() int {

	if fIoWriter.lock == nil {
		fIoWriter.lock = new(sync.Mutex)
	}

	fIoWriter.lock.Lock()

	defer fIoWriter.lock.Unlock()

	new(fileIoWriterMolecule).
		validateDefaultByteArraySize(fIoWriter)

	return fIoWriter.defaultByteArraySize
}

// GetIoWriter
//
// Returns the internal io.Writer object encapsulated by
// the current instance of FileIoWriter.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method returns the internal io.Writer object
//	configured for the current instance of FileIoWriter.
//
//	Be sure to release this io.Writer object when it
//	is no longer needed.
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
//	writer						io.Writer
//
//		The concrete instance of io.Writer encapsulated
//		by the current instance of FileIoWriter.
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
func (fIoWriter *FileIoWriter) GetIoWriter(
	errorPrefix interface{}) (
	writer io.Writer,
	err error) {

	if fIoWriter.lock == nil {
		fIoWriter.lock = new(sync.Mutex)
	}

	fIoWriter.lock.Lock()

	defer fIoWriter.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileIoWriter."+
			"GetIoWriter()",
		"")

	if err != nil {
		return writer, err
	}

	if fIoWriter.ioWriter == nil {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of 'FileIoWriter' (%v) is invalid!\n"+
			"The internal '%v' io.Writer object has NOT been initialized.\n"+
			"Call one of the 'New' or 'Setter' methods when creating\n"+
			"an instance of 'FileIoWriter'\n",
			ePrefix.String(),
			"fIoWriter",
			"fIoWriter")

		return writer, err
	}

	writer = *fIoWriter.ioWriter

	return writer, err
}

// IsValidInstanceError
//
// Analyzes the current FileIoWriter instance to
// determine if is invalid.
//
// If the current FileIoWriter instance is found
// to be invalid, an error is returned.
//
// If the current FileIoWriter instance is valid,
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
//		FileIoWriter are found to be invalid,
//		this method will return an error configured
//		with an appropriate message identifying the
//		invalid	member data variable.
//
//		If all internal member data variables evaluate
//		as valid, this returned error value will be set
//		to 'nil'.
//
//		If errors are encountered during processing or if
//		any FileIoWriter internal member data
//	 	values are found to be invalid, the returned error
//	 	will encapsulate an appropriate error message.
//	 	This returned error message will incorporate the
//	 	method chain and text passed by input parameter,
//	 	'errorPrefix'. The 'errorPrefix' text will be
//	 	prefixed to the beginning of the error message.
func (fIoWriter *FileIoWriter) IsValidInstanceError(
	errorPrefix interface{}) error {

	if fIoWriter.lock == nil {
		fIoWriter.lock = new(sync.Mutex)
	}

	fIoWriter.lock.Lock()

	defer fIoWriter.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileIoWriter."+
			"IsValidInstanceError()",
		"")

	if err != nil {
		return err
	}

	return new(fileIoWriterAtom).
		isFileIoWriterValid(
			fIoWriter,
			"fIoWriter",
			ePrefix)
}

// New
//
// This method returns a pointer to a new, empty instance
// of FileIoWriter.
//
// After creating an empty instance of FileIoWriter,
// the user may call the 'Setter' methods to configure
// this FileIoWriter instance with valid parameters.
//
// Step-1
//
//	Call this method FileIoWriter.New() to generate an
//	empty version of FileIoWriter.
//
// Step-2
//
//	Use this returned instance of FileIoWriter and call
//	the appropriate 'Setter' methods to individually
//	configure the internal 'io.writer' object.
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
//	*FileIoWriter
//
//		This method returns a pointer to an empty
//		instance of	FileIoWriter.
func (fIoWriter *FileIoWriter) New() *FileIoWriter {

	if fIoWriter.lock == nil {
		fIoWriter.lock = new(sync.Mutex)
	}

	fIoWriter.lock.Lock()

	defer fIoWriter.lock.Unlock()

	return new(FileIoWriter)
}

// NewIoWriter
//
// This method returns a fully initialized instance of
// FileIoWriter.
//
// This returned instance of FileIoWriter is created
// using an object implementing the io.Writer interface
// and passed as input parameter 'writer'.
//
// The size of the internal read buffer is controlled by
// input parameter 'bufSize'. The minimum buffer size is
// 16-bytes.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	The returned instance of FileIoWriter does NOT
//		use buffered write techniques. Instead, it
//		implements a direct write protocol using
//		io.Writer.
//
//	(2)	When all 'write' operations are completed and the
//		services of the returned new instance of
//		FileIoWriter are no longer required, the user
//		MUST perform 'close' and Clean-Up operations by
//		calling this local methods:
//
//			FileIoWriter.Close()
//
//	(3)	After executing the 'close' operation described
//		in paragraph (2) above, the current instance of
//		FileIoWriter will be rendered invalid and
//		unavailable for future 'write' operations.
//
//	(4) If the input parameter 'writer' base type is
//		NOT *os.File, the user will be required to
//		execute any 'close' or Clean-Up operations
//		required by the external 'writer' object in
//		addition to those 'close' and Clean-Up operations
//		specified in paragraph (2), above.
//
//	(5)	Input parameter 'writer' will accept a pointer to
//		an instance of os.File because os.File implements
//		the io.Writer interface.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://pkg.go.dev/io
//	https://pkg.go.dev/io#Writer
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
//		file and performing any other required Clean-Up
//		operations in addition to calling local method
//		FileIoWriter.Close().
//
//		While the returned instance of FileIoWriter
//		is primarily designed for writing data to disk
//		files, this 'writer' will in fact write data to
//		any object implementing the io.Writer interface.
//
//	defaultWriterByteArraySize	int
//
//		The size of the byte array which will be used to
//		write data to the internal io.Writer object
//		encapsulated by the returned FileIoWriter
//		instance.
//
//		If the value of 'defaultWriterByteArraySize' is
//		less than one ('1'), it will be reset to a size
//		of '4096'.
//
//		Although the FileIoWriter type does not use the
//		'buffered' write protocol, the size of the byte
//		array used to write bytes to the underlying
//		io.Writer object is variable.
//
//		Methods utilizing the Default Byte Array Size
//		include:
//
//			FileIoWriter.ReadFrom()
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
//	FileIoWriter
//
//		If this method completes successfully, it will
//		return a fully configured instance of
//		FileIoWriter.
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
func (fIoWriter *FileIoWriter) NewIoWriter(
	writer io.Writer,
	defaultWriterByteArraySize int,
	errorPrefix interface{}) (
	*FileIoWriter,
	error) {

	if fIoWriter.lock == nil {
		fIoWriter.lock = new(sync.Mutex)
	}

	fIoWriter.lock.Lock()

	defer fIoWriter.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error
	var newFileIoWriter = new(FileIoWriter)

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileIoWriter."+
			"NewIoWriter()",
		"")

	if err != nil {
		return newFileIoWriter, err
	}

	err = new(fileIoWriterNanobot).
		setIoWriter(
			newFileIoWriter,
			"newFileIoWriter",
			writer,
			"writer",
			defaultWriterByteArraySize,
			ePrefix.XCpy("newFileIoWriter"))

	return newFileIoWriter, err
}

// NewFileMgr
//
// Receives an instance of FileMgr as input parameter
// 'fileMgr' and returns a new, fully configured instance
// of FileIoWriter.
//
// The FileIoWriter instance returned by this method will
// configure the file identified by 'fileMgr' as the data
// source for file 'write' operations.
//
// This target 'write' file identified by 'fileMgr' is
// opened for either 'write-only' or 'write/write'
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
//	https://pkg.go.dev/io#Writer
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	As a precaution, the incoming 'fileMgr' object
//		will be closed before configuring 'fIoWriter'
//		with a new internal io.Writer object.
//
//	(2)	The returned instance of FileIoWriter does NOT
//		use buffered write techniques. Instead, it
//		implements a direct write protocol using
//		io.Writer.
//
//	(3)	When all write operations have been completed and
//		there is no further need for the returned
//		instance of FileIoWriter, the user is responsible
//		for 'closing' and releasing the associated memory
//		resources by calling the local method
//		FileIoWriter.Close().
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fileMgr						*FileMgr
//
//		A pointer to an instance of FileMgr. The file
//		identified by 'fileMgr' will be used as a data
//		source for 'write' operations.
//
//		If the path and file name encapsulated by
//		'fileMgr' do not currently exist on an attached
//		storage drive, an error will be returned.
//
//		The instance of FileIoWriter returned by this
//		method will configure the file identified by
//		'fileMgr' as the data source for file 'write'
//		operations.
//
//		As a precaution, the incoming 'fileMgr' object
//		will be closed before configuring the returned
//		instance of 'FileIoWriter' with a new internal
//		io.Writer object.
//
//	openFileReadWrite			bool
//
//		If this parameter is set to 'true', the target
//		'write' file identified by input parameter
//		'fileMgr' will be opened for both 'write' and
//		'write' operations.
//
//		If 'openFileReadWrite' is set to 'false', the
//		target 'write' file will be opened for 'write-only'
//		operations.
//
//	defaultByteArraySize		int
//
//		The size of the byte array which will be used to
//		write data to the internal io.Writer object
//		encapsulated by the returned FileIoWriter
//		instance.
//
//		If the value of 'defaultByteArraySize' is
//		less than one ('1'), it will be reset to a size
//		of '4096'.
//
//		Although the FileIoWriter type does not use the
//		'buffered' write protocol, the size of the byte
//		array used to write bytes to the underlying
//		io.Writer object is variable.
//
//		Methods utilizing the Default Byte Array Size
//		include:
//
//			FileIoWriter.ReadFrom()
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
//	FileInfoPlus
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
//	FileIoWriter
//
//		If this method completes successfully, a fully
//		configured instance of FileIoWriter will
//		be returned.
//
//		This returned instance of FileIoWriter will
//		configure the file identified by input parameter
//		'fileMgr' is a data source for file 'write'
//		operations.
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
func (fIoWriter *FileIoWriter) NewFileMgr(
	fileMgr *FileMgr,
	openFileReadWrite bool,
	defaultByteArraySize int,
	truncateExistingFile bool,
	errorPrefix interface{}) (
	FileInfoPlus,
	*FileIoWriter,
	error) {

	if fIoWriter.lock == nil {
		fIoWriter.lock = new(sync.Mutex)
	}

	fIoWriter.lock.Lock()

	defer fIoWriter.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error
	var fInfoPlus FileInfoPlus
	var newFileIoWriter = new(FileIoWriter)

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileIoWriter."+
			"NewFileMgr()",
		"")

	if err != nil {

		return fInfoPlus, newFileIoWriter, err
	}

	fInfoPlus,
		err = new(fileIoWriterMicrobot).
		setFileMgr(
			newFileIoWriter,
			"newFileIoWriter",
			fileMgr,
			"fileMgr",
			openFileReadWrite,
			defaultByteArraySize,
			truncateExistingFile,
			ePrefix.XCpy(
				"fileMgr"))

	return fInfoPlus, newFileIoWriter, err
}

// NewPathFileName
//
// Receives a path and file name as an input parameter
// string, 'pathFileName' and returns a new, fully
// configured instance of FileIoWriter.
//
// The target 'write' file identified by 'pathFileName'
// is opened for either 'write-only' or 'write/write'
// operations depending on input parameter
// 'openFileReadWrite'.
//
// This target 'write' file identified by 'pathFileName'
// will be used to create a file pointer (*os.File) which
// in turn will be used to configure the internal
// io.Writer.
//
// If the target path and file identified by
// 'pathFileName' do not currently exist on an attached
// storage drive, an error will be returned.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://pkg.go.dev/io#Writer
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	The returned instance of FileIoWriter does NOT
//		use buffered write techniques. Instead, it
//		implements a direct write protocol using
//		io.Writer.
//
//	(2)	When all write operations have been completed and
//		there is no further need for the returned
//		instance of FileIoWriter, the user is responsible
//		for 'closing' and releasing the associated memory
//		resources by calling the local method
//		FileIoWriter.Close().
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathFileName				string
//
//		This string contains the path and file name of
//		the file which will be used a data source for
//		'write' operations.
//
//		If this file does not currently exist on an
//		attached storage drive, an error will be
//		returned.
//
//	openFileReadWrite			bool
//
//		If this parameter is set to 'true', the target
//		'write' file identified from input parameter
//		'pathFileName' will be opened for both 'write'
//		and 'write' operations.
//
//		If 'openFileReadWrite' is set to 'false', the
//		target 'write' file will be opened for 'write-only'
//		operations.
//
//	defaultByteArraySize		int
//
//		The size of the byte array which will be used to
//		write data to the internal io.Writer object
//		encapsulated by the returned FileIoWriter
//		instance.
//
//		If the value of 'defaultByteArraySize' is
//		less than one ('1'), it will be reset to a size
//		of '4096'.
//
//		Although the FileIoWriter type does not use the
//		'buffered' write protocol, the size of the byte
//		array used to write bytes to the underlying
//		io.Writer object is variable.
//
//		Methods utilizing the Default Byte Array Size
//		include:
//
//			FileIoWriter.ReadFrom()
//
//	truncateExistingFile			bool
//
//		If this parameter is set to 'true', the target
//		'write' file specified by 'pathFileName' will be
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
//	*FileIoWriter
//
//		If this method completes successfully, a pointer
//		to a fully configured instance of FileIoWriter
//		will be returned.
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
func (fIoWriter *FileIoWriter) NewPathFileName(
	pathFileName string,
	openFileReadWrite bool,
	defaultByteArraySize int,
	truncateExistingFile bool,
	errorPrefix interface{}) (
	FileInfoPlus,
	*FileIoWriter,
	error) {

	if fIoWriter.lock == nil {
		fIoWriter.lock = new(sync.Mutex)
	}

	fIoWriter.lock.Lock()

	defer fIoWriter.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error
	var newFileIoWriter = new(FileIoWriter)
	var fInfoPlus FileInfoPlus

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileIoWriter."+
			"NewPathFileName()",
		"")

	if err != nil {

		return fInfoPlus, newFileIoWriter, err
	}

	fInfoPlus,
		err = new(fileIoWriterNanobot).
		setPathFileName(
			newFileIoWriter,
			"newFileIoWriter",
			pathFileName,
			"pathFileName",
			openFileReadWrite,
			defaultByteArraySize,
			truncateExistingFile,
			ePrefix.XCpy(
				pathFileName))

	return fInfoPlus, newFileIoWriter, err
}

// ReadFrom
//
// Implements the io.ReadFrom interface.
//
// This method will all the read data residing the
// io.Reader object passed as input parameter 'reader'
// and write that data to the io.Writer encapsulated by
// the current instance of FileIoWriter.
//
// The data is read from 'reader' using an internal byte
// array equal in length to the default array size
// previously configured for the current instance of
// FileIoWriter.
//
// The return parameter 'numOfBytesProcessed' records the
// number of bytes read from 'reader' and written to
// the io.Writer object. If the number of bytes read
// fails to match the number of bytes written, an error
// will be returned.
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
//		io.Writer object encapsulated by the current
//		instance of FileIoWriter.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	numOfBytesProcessed			int64
//
//		This return parameter documents the number of
//		bytes read from 'reader' and written to the
//		FileIoWriter io.Writer object. If the
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
func (fIoWriter *FileIoWriter) ReadFrom(
	reader io.Reader) (
	numOfBytesProcessed int64,
	err error) {

	if fIoWriter.lock == nil {
		fIoWriter.lock = new(sync.Mutex)
	}

	fIoWriter.lock.Lock()

	defer fIoWriter.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"FileIoWriter."+
			"ReadFrom()",
		"")

	if err != nil {

		return numOfBytesProcessed, err
	}

	if fIoWriter.ioWriter == nil {

		err = fmt.Errorf("%v\n"+
			"-------------------------------------------------------\n"+
			"Error: This instance of 'FileIoWriter' is invalid!\n"+
			"The internal io.Writer object has NOT been properly\n"+
			"initialized. Call one of the 'New' or 'Setter'\n"+
			"methods to create a valid instance of 'FileIoWriter'\n",
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

	new(fileIoWriterMolecule).
		validateDefaultByteArraySize(fIoWriter)

	var bytesRead = make([]byte,
		fIoWriter.defaultByteArraySize)

	var numBytesRead, numBytesWritten int
	var err1, err2 error
	var maxCycle = math.MaxInt - 1
	var cycleCnt int
	var writer io.Writer
	writer = *fIoWriter.ioWriter

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
			(errors.Is(err1, io.EOF) == false) {

			err = fmt.Errorf("%v\n"+
				"Error: reader.Read(bytesRead)\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				err1.Error())

			break

		}

		if numBytesRead > 0 {

			numBytesWritten,
				err2 = writer.Write(
				bytesRead[0:numBytesRead])

			if err2 != nil {

				err = fmt.Errorf("%v\n"+
					"Error returned by FileIoWriter.ioWriter.Write(bytesRead[:numBytesRead])\n"+
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
// This method sets the byte offset for the next 'write'
// operation within the internal io.Writer object
// encapsulated in the current FileIoWriter instance.
//
// The Seek() method only succeeds if the internal
// io.Writer object, encapsulated by the current
// FileIoWriter instance, implements the io.Seeker
// interface. Disk files with a base type of os.File and
// the FileMgr type are among those types which implement
// the io.Seeker interface.
//
// The new byte offset ('targetOffset') is interpreted
// according to input parameter 'whence'.
//
// 'whence' is an integer value designating whether the
// input parameter 'targetOffset' is interpreted to mean
// an offset from the start of the file, an offset from
// the current offset position or an offset from the end
// of the file. The 'whence' parameter must be passed as
// one of the following 'io' integer constant values:
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
// position specified by input parameter 'targetOffset'.
//
// Seek returns the new offset relative to the start of the
// io.Writer object or an error, if any.
//
// Method Seek() implements the 'io.Seeker' interface.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	If the current instance of FileIoWriter was
//		NOT initialized with an io.Writer object which
//		implements the io.Seeker interface, this method
//		will return an error.
//
//	(2)	Setting a byte offset which occurs before the
//		start of the internal io.Writer object
//		encapsulated by the current FileIoWriter instance
//		will result in an error.
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
//	(4) This method implements the 'io.Seeker' interface.
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
func (fIoWriter *FileIoWriter) Seek(
	targetOffset int64,
	whence int) (
	offsetFromFileStart int64,
	err error) {

	if fIoWriter.lock == nil {
		fIoWriter.lock = new(sync.Mutex)
	}

	fIoWriter.lock.Lock()

	defer fIoWriter.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"FileIoWriter."+
			"Seek()",
		"")

	if err != nil {

		return offsetFromFileStart, err
	}

	offsetFromFileStart,
		err = new(fileIoWriterMicrobot).
		seekByOffset(
			fIoWriter,
			"fIoWriter",
			targetOffset,
			whence,
			ePrefix)

	return offsetFromFileStart, err
}

// SetDefaultByteArraySize
//
// Sets the default size of the byte array used to write
// bytes to the output data destination object
// (io.Writer) encapsulated in the current instance of
// FileIoWriter.
//
// Although the FileIoWriter type does not use the
// 'buffered' write protocol, the size of the byte array
// used to write bytes to the underlying io.Writer object
// is variable is some cases.
//
// Methods which utilize the Default Writer Byte Array
// Size include:
//
//	FileIoWriter.ReadFrom()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	defaultWriterByteArraySize		int
//
//		The size of the byte array which will be used to
//		write data to the output data destination object
//		(io.Writer) encapsulated by the current instance
//		of FileIoWriter.
//
//		If the value of 'defaultWriterByteArraySize' is
//		less than  one ('1'), it will be automatically
//		reset to a size of '4096'.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	-- NONE --
func (fIoWriter *FileIoWriter) SetDefaultByteArraySize(
	defaultWriterByteArraySize int) {

	if fIoWriter.lock == nil {
		fIoWriter.lock = new(sync.Mutex)
	}

	fIoWriter.lock.Lock()

	defer fIoWriter.lock.Unlock()

	fIoWriter.defaultByteArraySize =
		defaultWriterByteArraySize

	new(fileIoWriterMolecule).
		validateDefaultByteArraySize(fIoWriter)

	return
}

// SetIoWriter
//
// This method will completely re-initialize the current
// instance of FileIoWriter using the io.Writer object
// passed as input parameter 'writer'.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://pkg.go.dev/io#Writer
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	This method will delete, overwrite and reset all
//		pre-existing data values in the current instance
//		of FileIoWriter.
//
//	(2)	The user is responsible for 'closing' the
//		instance of io.Writer passed as input parameter
//		'writer'. The FileIoWriter.Close() method will
//		NOT close the 'writer' object.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	writer						io.Writer
//
//		An object which implements io.Writer interface.
//		This object will be used as a data source for
//		'write' operations.
//
//		The io.Writer object may be a file pointer of
//		type *os.File because file pointers of this type
//		implement the io.Writer interface.
//
//		A file pointer (*os.File) will facilitate writing
//		data from files residing on an attached storage
//		drive. However, with this configuration, the user
//		is responsible for manually closing the file and
//		performing any other required Clean-Up operations
//		in addition to calling FileIoWriter.Close().
//
//		While the returned instance of FileIoWriter
//		is primarily designed for writing data from disk
//		files, this 'writer' will in fact write data from
//		any object implementing the io.Writer interface.
//
//	defaultByteArraySize		int
//
//		The size of the byte array which will be used to
//		write data to the internal io.Writer object
//		encapsulated by the current instance of
//		FileIoWriter.
//
//		If the value of 'defaultByteArraySize' is
//		less than one ('1'), it will be reset to a size
//		of '4096'.
//
//		Although the FileIoWriter type does not use the
//		'buffered' write protocol, the size of the byte
//		array used to write bytes to the underlying
//		io.Writer object is variable.
//
//		Methods utilizing the Default Byte Array Size
//		include:
//
//			FileIoWriter.ReadFrom()
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
func (fIoWriter *FileIoWriter) SetIoWriter(
	writer io.Writer,
	defaultByteArraySize int,
	errorPrefix interface{}) error {

	if fIoWriter.lock == nil {
		fIoWriter.lock = new(sync.Mutex)
	}

	fIoWriter.lock.Lock()

	defer fIoWriter.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileIoWriter."+
			"SetIoWriter()",
		"")

	if err != nil {
		return err
	}

	err = new(fileIoWriterNanobot).
		setIoWriter(
			fIoWriter,
			"fIoWriter",
			writer,
			"writer",
			defaultByteArraySize,
			ePrefix.XCpy("fIoWriter"))

	return err
}

// SetFileMgr
//
// This method will completely re-initialize the current
// instance of FileIoWriter using the path and file
// name identified by the FileMgr instance passed as
// input parameter 'fileMgr'.
//
// The file identified by 'fileMgr' will be used to
// reconfigure the internal io.Writer encapsulated in
// the current instance of FileIoWriter.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://pkg.go.dev/io#Writer
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	This method will delete, overwrite and reset all
//		pre-existing data values in the current instance
//		of FileIoWriter.
//
//	(2)	As a precaution, the incoming 'fileMgr' object
//		will be closed before configuring the current
//		FileIoWriter instance with a new, internal
//		io.Writer object.
//
//	(3)	When all write operations have been completed and
//		there is no further need for the current instance
//		of FileIoWriter, the user is responsible for
//		'closing' and releasing the associated memory
//		resources by calling the local method
//		FileIoWriter.Close().
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fileMgr						*FileMgr
//
//		A pointer to an instance of FileMgr. The file
//		identified by 'fileMgr' will be used as a data
//		source for 'write' operations.
//
//		If the path and file name encapsulated by
//		'fileMgr' do not currently exist on an attached
//		storage drive, an error will be returned.
//
//		The instance of FileIoWriter returned by this
//		method will configure the file identified by
//		'fileMgr' as the data source for file 'write'
//		operations.
//
//	openFileReadWrite			bool
//
//		If this parameter is set to 'true', the target
//		'write' file identified from input parameter
//		'pathFileName' will be opened for both 'write'
//		and 'write' operations.
//
//		If 'openFileReadWrite' is set to 'false', the
//		target 'write' file will be opened for 'write-only'
//		operations.
//
//	defaultByteArraySize		int
//
//		The size of the byte array which will be used to
//		write data to the internal io.Writer object
//		encapsulated by the current instance of
//		FileIoWriter.
//
//		If the value of 'defaultByteArraySize' is
//		less than one ('1'), it will be reset to a size
//		of '4096'.
//
//		Although the FileIoWriter type does not use the
//		'buffered' write protocol, the size of the byte
//		array used to write bytes to the underlying
//		io.Writer object is variable.
//
//		Methods utilizing the Default Byte Array Size
//		include:
//
//			FileIoWriter.ReadFrom()
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
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fIoWriter *FileIoWriter) SetFileMgr(
	fileMgr *FileMgr,
	openFileReadWrite bool,
	defaultByteArraySize int,
	truncateExistingFile bool,
	errorPrefix interface{}) (
	fInfoPlus FileInfoPlus,
	err error) {

	if fIoWriter.lock == nil {
		fIoWriter.lock = new(sync.Mutex)
	}

	fIoWriter.lock.Lock()

	defer fIoWriter.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileIoWriter."+
			"SetFileMgr()",
		"")

	if err != nil {

		return fInfoPlus, err
	}

	fInfoPlus,
		err = new(fileIoWriterMicrobot).
		setFileMgr(
			fIoWriter,
			"fIoWriter",
			fileMgr,
			"fileMgr",
			openFileReadWrite,
			defaultByteArraySize,
			truncateExistingFile,
			ePrefix.XCpy(
				"fileMgr"))

	return fInfoPlus, err
}

// SetPathFileName
//
// This method will completely re-initialize the current
// instance of FileIoWriter using the path and file name
// passed as input parameter 'pathFileName'.
//
// The path and file name specified by 'pathFileName'
// will be used to reconfigure the internal io.Writer
// encapsulated in the current instance of
// FileIoWriter.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://pkg.go.dev/io#Writer
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	This method will delete, overwrite and reset all
//		pre-existing data values in the current instance
//		of FileIoWriter.
//
//	(2)	When all write operations have been completed and
//		there is no further need for the current instance
//		of FileIoWriter, the user is responsible for
//		'closing' and releasing the associated memory
//		resources by calling the local method
//		FileIoWriter.Close().
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathFileName				string
//
//		This string contains the path and file name of
//		the file which will be configured as a new data
//		source for 'write' operations.
//
//		If this file does not currently exist on an
//		attached storage drive, an error will be
//		returned.
//
//	openFileReadWrite			bool
//
//		If this parameter is set to 'true', the target
//		'write' file identified from input parameter
//		'pathFileName' will be opened for both 'write'
//		and 'write' operations.
//
//		If 'openFileReadWrite' is set to 'false', the
//		target 'write' file will be opened for 'write-only'
//		operations.
//
//	defaultByteArraySize		int
//
//		The size of the byte array which will be used to
//		write data to the internal io.Writer object
//		encapsulated by the current instance of
//		FileIoWriter.
//
//		If the value of 'defaultByteArraySize' is
//		less than one ('1'), it will be reset to a size
//		of '4096'.
//
//		Although the FileIoWriter type does not use the
//		'buffered' write protocol, the size of the byte
//		array used to write bytes to the underlying
//		io.Writer object is variable.
//
//		Methods utilizing the Default Byte Array Size
//		include:
//
//			FileIoWriter.ReadFrom()
//
//	truncateExistingFile			bool
//
//		If this parameter is set to 'true', the target
//		'write' file specified by 'pathFileName' will be
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
func (fIoWriter *FileIoWriter) SetPathFileName(
	pathFileName string,
	openFileReadWrite bool,
	defaultByteArraySize int,
	truncateExistingFile bool,
	errorPrefix interface{}) (
	fInfoPlus FileInfoPlus,
	err error) {

	if fIoWriter.lock == nil {
		fIoWriter.lock = new(sync.Mutex)
	}

	fIoWriter.lock.Lock()

	defer fIoWriter.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileIoWriter."+
			"SetPathFileName()",
		"")

	if err != nil {

		return fInfoPlus, err
	}

	fInfoPlus,
		err = new(fileIoWriterNanobot).
		setPathFileName(
			fIoWriter,
			"fIoWriter",
			pathFileName,
			"pathFileName",
			openFileReadWrite,
			defaultByteArraySize,
			truncateExistingFile,
			ePrefix.XCpy(
				pathFileName))

	return fInfoPlus, err
}

// Write
//
// Writes the contents of the byte array input paramter
// ('bytesToWrite') to the internal destination
// io.Writer object previously configured for this
// instance of FileIoWriter.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://pkg.go.dev/io
//	https://pkg.go.dev/io#Writer
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1) This method implements the io.Writer interface.
//
//	(2)	After all 'write' operations have been completed,
//		the user MUST call the 'Close' method to perform
//		necessary Clean-Up operations:
//
//			FileIoWriter.Close()
//
//	(3)	If the planned number of bytes to be written as
//		specified by the length of 'bytesToWrite' does
//		NOT match the actual number of bytes written to
//		the internal io.Writer object, an error will be
//		returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	bytesToWrite				[]byte
//
//		The contents of this byte array will be written
//		to the internal destination io.Writer object
//		previously configured for the current instance of
//		FileIoWriter.
//
//		Typically, the internal destination io.Writer
//		object will be a data file existing on an attached
//		storage drive. However, the destination
//		io.Writer object may be any object implementing
//		the io.Writer interface.
//
//		This method WILL VERIFY that the number of bytes
//		written is equal to the length of the byte array
//		input parameter 'bytesToWrite'.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	numBytesWritten				int
//
//		This parameter returns the number of bytes
//		written to the internal destination io.Writer
//		object configured for the current instance of
//		FileIoWriter.
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
//		If the number of bytes written ('numBytesWritten')
//		is less than the number of bytes contained in
//		byte array input parameter 'bytesToWrite', an
//	 	error will be returned.
func (fIoWriter *FileIoWriter) Write(
	bytesToWrite []byte) (
	numBytesWritten int,
	err error) {

	if fIoWriter.lock == nil {
		fIoWriter.lock = new(sync.Mutex)
	}

	fIoWriter.lock.Lock()

	defer fIoWriter.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"FileIoWriter."+
			"Write()",
		"")

	if err != nil {

		return numBytesWritten, err
	}

	numBytesWritten,
		err = new(fileIoWriterMicrobot).
		writeBytes(
			fIoWriter,
			"fIoWriter",
			bytesToWrite,
			ePrefix.XCpy("fIoWriter"))

	return numBytesWritten, err
}

// WriteAt
//
// This method implements the io.WriterAt interface.
//
// The bytes contained in the byte array passed as input
// parameter 'bytesToWrite' will be written to the
// internal io.Writer object encapsulated by the current
// FileIoWriter instance at the offset specified by
// 'targetOffset'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	bytesToWrite				[]byte
//
//		The contents of this byte array will be written
//		to the internal destination io.Writer object
//		previously configured for the current instance of
//		FileIoWriter.
//
//		Typically, the internal destination io.Writer
//		object will be a data file existing on an attached
//		storage drive. However, the destination
//		io.Writer object may be any object implementing
//		the io.Writer interface.
//
//		This method WILL NOT VERIFY that the number of
//		bytes written is equal to the length of the
//		length of input parameter 'bytesToWrite'.
//
//	targetOffset				int64
//
//		The number of bytes used to reset the file
//		offset to the beginning byte at which all
//		bytes contained in 'bytesToWrite' will be
//		written.
//
//		'targetOffset' represents the number of bytes
//		from the beginning of the io.Writer object.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	numBytesWritten				int
//
//		This parameter returns the number of bytes
//		written to the internal destination io.Writer
//		object configured for the current instance of
//		FileIoWriter.
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
//		If the number of bytes written ('numBytesWritten')
//		does not equal the number of bytes contained in
//		byte array input parameter 'bytesToWrite', an
//	 	error will be returned.
func (fIoWriter *FileIoWriter) WriteAt(
	bytesToWrite []byte,
	targetOffset int64) (
	numBytesWritten int,
	err error) {

	if fIoWriter.lock == nil {
		fIoWriter.lock = new(sync.Mutex)
	}

	fIoWriter.lock.Lock()

	defer fIoWriter.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "FileIoWriter." +
		"WriteAt()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		funcName,
		"")

	if err != nil {
		return numBytesWritten, err
	}

	lenBytesToWrite := len(bytesToWrite)

	if lenBytesToWrite == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'bytesToWrite' is invalid!\n"+
			"The 'bytesToWrite' byte array is empty. It has zero bytes.\n",
			ePrefix.String())

		return numBytesWritten, err
	}

	var ok bool
	var writerAtObj io.WriterAt
	var localWriter io.Writer
	localWriter = *fIoWriter.ioWriter

	writerAtObj, ok = localWriter.(io.WriterAt)

	if !ok {

		err = fmt.Errorf("%v\n"+
			"Error: This WriteAt method was invoked on a\n"+
			"'FileIoWriter' internal io.Writer object\n"+
			"which does NOT support the io.WriterAt\n"+
			"interface. This means:\n"+
			"(1) The 'WriteAt' method is unavailable.\n"+
			"              AND\n"+
			"(2) The 'FileIoWriter' internal io.Writer\n"+
			"    was created from an object which does\n"+
			"    NOT implement the io.WriterAt interface.\n",
			ePrefix.String())

		return numBytesWritten, err
	}

	numBytesWritten,
		err = writerAtObj.WriteAt(
		bytesToWrite,
		targetOffset)

	return numBytesWritten, err
}

// WriteTextOrNumbers
//
// This method will accept many different text or numeric
// data types which are then converted to a byte or
// string array and written to the io.Writer object
// configured for the current instance of FileIoWriter.
//
// The text or numeric data type passed as input
// parameter 'charsToWrite' must match one of over sixty
// eligible data types. Eligible data types include
// strings.Builder, strings, string arrays, byte arrays
// and numeric values such as big.Int, big.Rat, float64
// and number strings (NumberStrKernel).
//
// If 'charsToWrite' is set to an ineligible data type,
// an error will be returned.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://pkg.go.dev/io
//	https://pkg.go.dev/io#Writer
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1) This method implements the io.Writer interface.
//
//	(2)	After all 'write' operations have been completed,
//		the user MUST call the 'Close' method to perform
//		necessary Clean-Up operations:
//
//			FileIoWriter.Close()
//
//	(3)	If the planned number of bytes to be written as
//		specified by the length of 'bytesToWrite' does
//		NOT match the actual number of bytes written to
//		the internal io.Writer object, an error will be
//		returned.
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
//	writeEndOfLineChars			string
//
//		This character string is appended to each line of
//		text written to the io.Writer object. This
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
//		to the io.Writer object.
//
//	writeEndOfTextChars			string
//
//		A character string which will be written to the
//		io.Writer object after all other text from
//		'charsToWrite' and 'writeEndOfLineChars' have
//		been processed and written.
//
//	autoCloseOnExit				bool
//
//		When this parameter is set to 'true' and no
//		processing errors are encountered, this method
//		will automatically perform the following Clean-Up
//		tasks upon exit:
//
//		(1)	The FileIoWriter internal io.Writer object
//			will be properly closed and there will be no
//			need to make a separate call to local method,
//			FileIoWriter.Close().
//
//		(2) After performing this Clean-Up operation, the
//			current instance of FileIoWriter will invalid
//			and unusable for future 'write' operations.
//
//		If input parameter 'autoCloseOnExit' is
//		set to 'false', this method will NOT close
//		the internal io.Writer object. This means the user
//	 	is then responsible for performing the Clean-Up
//		tasks when all 'write' operations have been
//		completed by calling the local method:
//
//			FileIoWriter.Close()
//
//		If a processing error is encountered during
//		method execution, 'autoCloseOnExit' will be
//		ignored and no action will be taken to close
//		the current instance of FileIoWriter.
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
//		The number of bytes written to the io.Writer
//		object configured for the current instance of
//		FileIoWriter.
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

func (fIoWriter *FileIoWriter) WriteTextOrNumbers(
	charsToWrite interface{},
	writeEndOfLineChars string,
	writeEndOfTextChars string,
	autoCloseOnExit bool,
	errorPrefix interface{}) (
	numOfBytesWritten int64,
	err error) {

	if fIoWriter.lock == nil {
		fIoWriter.lock = new(sync.Mutex)
	}

	fIoWriter.lock.Lock()

	defer fIoWriter.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileIoWriter."+
			"WriteTextOrNumbers()",
		"")

	if err != nil {

		return numOfBytesWritten, err
	}

	if fIoWriter.ioWriter == nil {

		err = fmt.Errorf("%v\n"+
			"-------------------------------------------------------\n"+
			"Error: This instance of 'FileIoWriter' is invalid!\n"+
			"The internal io.Writer has NOT been initialized.\n"+
			"Call one of the 'New' or 'Setter' methods when creating\n"+
			"a new valid instance of 'FileIoWriter'\n",
			ePrefix.String())

		return numOfBytesWritten, err
	}

	if fIoWriter.ioWriter == nil {

		err = fmt.Errorf("%v\n"+
			"-------------------------------------------------------\n"+
			"Error: This instance of 'FileIoWriter' is invalid!\n"+
			"The internal io.Writer object has NOT been properly\n"+
			"initialized. Call one of the 'New' or 'Setter'\n"+
			"methods to create a valid instance of 'FileIoWriter'.\n",
			ePrefix.String())

		return numOfBytesWritten, err
	}

	var writeBytesFunc = fIoWriter.lowLevelWriteToBytes

	numOfBytesWritten,
		err = new(fileWriterHelperMicrobot).
		writeCharacters(
			writeBytesFunc,
			charsToWrite,
			"charsToWrite",
			writeEndOfLineChars,
			writeEndOfTextChars,
			ePrefix.XCpy("fIoWriter.ioWriter<-charsToWrite"))

	if err == nil &&
		autoCloseOnExit == true {

		err = new(fileIoWriterMolecule).
			closeAndRelease(
				fIoWriter,
				"fIoWriter",
				true, // releaseMemoryResources
				ePrefix)
	}

	return numOfBytesWritten, err
}

// lowLevelWriteToBytes
//
// This method is designed to be passed as a text handler
// and text writer for method:
//
//	fileWriterHelperMicrobot.writeCharacters()
func (fIoWriter *FileIoWriter) lowLevelWriteToBytes(
	bytesToWrite []byte,
	writeLastTextChars string,
	xErrPref *ePref.ErrPrefixDto) (
	numOfBytesWritten int64,
	err error) {

	funcName := "FileIoWriter." +
		"writeToBytes()"

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		xErrPref,
		funcName,
		"")

	if fIoWriter.ioWriter == nil {

		err = fmt.Errorf("%v\n"+
			"-------------------------------------------------------\n"+
			"Error: This instance of 'FileIoWriter' is invalid!\n"+
			"The internal io.Writer has NOT been initialized.\n"+
			"Call one of the 'New' or 'Setter' methods when creating\n"+
			"a new valid instance of 'FileIoWriter'\n",
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

	var writer io.Writer
	var localNumOfBytesWritten int
	var expectedNumOfBytesToWrite int64
	var err2 error
	writer = *fIoWriter.ioWriter

	expectedNumOfBytesToWrite = int64(lenBytesToWrite)

	localNumOfBytesWritten,
		err2 = writer.Write(bytesToWrite)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: writer.Write(bytesToWrite)\n"+
			"bytesToWrite = '%v'\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			string(bytesToWrite),
			err2.Error())

		return numOfBytesWritten, err
	}

	numOfBytesWritten +=
		int64(localNumOfBytesWritten)

	lenLastTextChars := len(writeLastTextChars)

	if lenLastTextChars > 0 {

		expectedNumOfBytesToWrite += int64(lenLastTextChars)

		localNumOfBytesWritten,
			err2 = writer.Write([]byte(writeLastTextChars))

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Error: writer.Write([]byte(writeLastTextChars))\n"+
				"writeLastTextChars= '%v'\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				writeLastTextChars,
				err2.Error())

			return numOfBytesWritten, err

		}

		numOfBytesWritten +=
			int64(localNumOfBytesWritten)
	}

	if expectedNumOfBytesToWrite != numOfBytesWritten {

		err = fmt.Errorf("%v\n"+
			"Error: Number of bytes written does NOT\n"+
			"match the expected number of bytes to write!"+
			"Expected number of bytes to write: %v\n"+
			"   Actual number of bytes written: %v\n",
			ePrefix.String(),
			expectedNumOfBytesToWrite,
			numOfBytesWritten)

	}

	return numOfBytesWritten, err
}

type fileIoWriterMicrobot struct {
	lock *sync.Mutex
}

// seekByOffset
//
// This method sets the byte offset for the next 'write'
// operation within the internal io.Writer object
// encapsulated in the FileIoReader instance passed as
// input parameter 'fIoReader'.
//
// This method only succeeds if the internal io.Writer
// object implements the io.Seeker interface. Disk
// files with a base type of os.File and FileMgr
// instances are among those types which implement the
// io.Seeker interface.
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
// This method implements the 'io.Seeker' interface.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	Seeking to an offset before the start of the file
//		is an error.
//
//	(2) If input parameter 'whence' is not set to one of
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
//	fIoWriter					*FileIoWriter
//
//		A pointer to an instance of FileIoWriter.
//
//		The io.Writer object encapsulated in this
//		FileIoWriter instance will use the target offset
//		specified by 'targetOffset' to reset the file
//		offset for the next 'write' operation.
//
//	fIoWriterLabel				string
//
//		The name or label associated with input parameter
//		'fIoWriter' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fIoWriter" will be
//		automatically applied.
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
func (fIoWriterMicrobot *fileIoWriterMicrobot) seekByOffset(
	fIoWriter *FileIoWriter,
	fIoWriterLabel string,
	targetOffset int64,
	whence int,
	errPrefDto *ePref.ErrPrefixDto) (
	offsetFromFileStart int64,
	err error) {

	if fIoWriterMicrobot.lock == nil {
		fIoWriterMicrobot.lock = new(sync.Mutex)
	}

	fIoWriterMicrobot.lock.Lock()

	defer fIoWriterMicrobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileIoWriterMicrobot." +
		"seekByOffset()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return offsetFromFileStart, err
	}

	if len(fIoWriterLabel) == 0 {

		fIoWriterLabel = "fIoWriter"
	}

	if fIoWriter == nil {

		err = fmt.Errorf("%v\n"+
			"Error: The FileIoWriter instance passed\n"+
			"as input parameter '%v' is invalid!\n"+
			"'%v' is a 'nil' pointer.\n",
			ePrefix.String(),
			fIoWriterLabel,
			fIoWriterLabel)

		return offsetFromFileStart, err
	}

	if fIoWriter.ioWriter == nil {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of FileIoWriter ('%v') is invalid!\n"+
			"The internal io.Writer object has NOT been initialized.\n"+
			"%v.ioWriter is equal to 'nil'.\n"+
			"Call one of the 'New' or 'Setter' methods to create a\n"+
			"valid instance of 'FileIoWriter'\n",
			ePrefix.String(),
			fIoWriterLabel,
			fIoWriterLabel)

		return offsetFromFileStart, err
	}

	if targetOffset < 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetOffset' is invalid!\n"+
			"'targetOffset' has a value less than zero.\n"+
			"targetOffset= %v\n",
			ePrefix.String(),
			targetOffset)

		return offsetFromFileStart, err
	}

	var ok bool
	var seekerObj io.Seeker
	var localWriter io.Writer
	localWriter = *fIoWriter.ioWriter

	seekerObj, ok = localWriter.(io.Seeker)

	if !ok {

		err = fmt.Errorf("%v\n"+
			"Error: This Seek method was invoked on an\n"+
			"'FileIoWriter' internal io.Writer object\n"+
			"which does NOT support the io.Seeker\n"+
			"interface. This means:\n"+
			"(1) The 'Seek' method is unavailable.\n"+
			"           AND\n"+
			"(2) The 'FileIoWriter' internal io.Writer\n"+
			"      object was created from something\n"+
			"      other than a disk file (*os.File)\n"+
			"      or FileMgr type.",
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
			"Error returned by low-level 'seek' operation\n"+
			"seekerObj.Seek(). The operation to set a byte\n"+
			"offset ('targetOffset') Failed.\n"+
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

// setFileMgr
//
// This 'setter' method is used to initialize new values
// for internal member variables contained in the
// instance of FileIoWriter passed as input parameter
// 'fIoWriter'.
//
// The new io.Writer object assigned to 'fIoWriter' is
// generated from the File Manager (FileMgr) instance
// passed as input parameter 'fileMgr'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1) This method will delete, overwrite and reset all
//		pre-existing data values in the instance of
//		FileIoWriter passed as input parameter
//		'fIoWriter'.
//
//	(2) As a precaution, the incoming 'fileMgr' object
//		will be closed before configuring 'fIoWriter'
//		with a new internal io.Writer object.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fIoWriter					*FileIoWriter
//
//		A pointer to an instance of FileIoWriter.
//
//		All internal member variable data values in
//		this instance will be deleted and initialized
//		to new values based on the following input
//		parameters.
//
//	fIoWriterLabel				string
//
//		The name or label associated with input parameter
//		'fIoWriter' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fIoWriter" will be
//		automatically applied.
//
//	fileMgr						*FileMgr
//
//		A pointer to an instance of FileMgr. The file
//		identified by 'fileMgr' will be used as a
//		data source for 'write' operations.
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
//		'write' file identified by input parameter
//		'fileMgr' will be opened for 'write' and 'write'
//		operations.
//
//		If 'openFileReadWrite' is set to 'false', the
//		target write file will be opened for 'write-only'
//		operations.
//
//	defaultByteArraySize		int
//
//		The size of the byte array which will be used to
//		write data to the internal io.Writer object
//		encapsulated by FileIoWriter input parameter
//		'fIoWriter'.
//
//		If the value of 'defaultByteArraySize' is
//		less than one ('1'), it will be reset to a size
//		of '4096'.
//
//		Although the FileIoWriter type does not use the
//		'buffered' write protocol, the size of the byte
//		array used to write bytes to the underlying
//		io.Writer object is variable.
//
//		Methods utilizing the Default Byte Array Size
//		include:
//
//			FileIoWriter.ReadFrom()
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
func (fIoWriterMicrobot *fileIoWriterMicrobot) setFileMgr(
	fIoWriter *FileIoWriter,
	fIoWriterLabel string,
	fileMgr *FileMgr,
	fileMgrLabel string,
	openFileReadWrite bool,
	defaultByteArraySize int,
	truncateExistingFile bool,
	errPrefDto *ePref.ErrPrefixDto) (
	fInfoPlus FileInfoPlus,
	err error) {

	if fIoWriterMicrobot.lock == nil {
		fIoWriterMicrobot.lock = new(sync.Mutex)
	}

	fIoWriterMicrobot.lock.Lock()

	defer fIoWriterMicrobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileIoWriterMicrobot." +
		"setFileMgr()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return fInfoPlus, err
	}

	if len(fIoWriterLabel) == 0 {

		fIoWriterLabel = "fIoWriter"
	}

	if len(fileMgrLabel) == 0 {

		fileMgrLabel = "fileMgr"
	}

	if fIoWriter == nil {

		err = fmt.Errorf("%v\n"+
			"Error: The FileIoWriter instance passed\n"+
			"as input parameter '%v' is invalid!\n"+
			"'%v' is a 'nil' pointer.\n",
			ePrefix.String(),
			fIoWriterLabel,
			fIoWriterLabel)

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
		err = new(fileIoWriterNanobot).
		setPathFileName(
			fIoWriter,
			fIoWriterLabel,
			fileMgr.absolutePathFileName,
			fileMgrLabel,
			openFileReadWrite,
			defaultByteArraySize,
			truncateExistingFile,
			ePrefix.XCpy(fIoWriterLabel+"<-"+fileMgrLabel))

	return fInfoPlus, err
}

// writeBytes
//
// Writes the contents of the byte array input paramter
// ('bytesToWrite') to the internal destination
// io.Writer object previously configured for the
// FileIoWriter instance passed as 'fIoWriter'.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://pkg.go.dev/io
//	https://pkg.go.dev/io#Writer
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fIoWriter					*FileIoWriter
//
//		A pointer to an instance of FileIoWriter.
//
//		Data contained in the byte array 'bytesToWrite'
//		will be written to the destination io.Writer
//		object encapsulated in this FileIoWriter
//		instance.
//
//	fIoWriterLabel				string
//
//		The name or label associated with input parameter
//		'fIoWriter' which will be used in error messages
//		returned by this method.
//
//	bytesToWrite				[]byte
//
//		The contents of this byte array will be written
//		to the internal destination io.Writer object
//		previously configured for the current instance of
//		FileIoWriter.
//
//		Typically, the internal destination io.Writer
//		object will be a data file existing on an attached
//		storage drive. However, the destination
//		io.Writer object may be any object implementing
//		the io.Writer interface.
//
//		This method WILL VERIFY that the number of bytes
//		written is equal to the length of the byte array
//		input parameter 'bytesToWrite'.
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
func (fIoWriterMicrobot *fileIoWriterMicrobot) writeBytes(
	fIoWriter *FileIoWriter,
	fIoWriterLabel string,
	bytesToWrite []byte,
	errPrefDto *ePref.ErrPrefixDto) (
	numBytesWritten int,
	err error) {

	if fIoWriterMicrobot.lock == nil {
		fIoWriterMicrobot.lock = new(sync.Mutex)
	}

	fIoWriterMicrobot.lock.Lock()

	defer fIoWriterMicrobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileIoWriterMicrobot." +
		"setFileMgr()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return numBytesWritten, err
	}

	if len(fIoWriterLabel) == 0 {

		fIoWriterLabel = "fIoWriter"
	}

	if fIoWriter == nil {

		err = fmt.Errorf("%v\n"+
			"Error: The FileIoWriter instance passed\n"+
			"as input parameter '%v' is invalid!\n"+
			"'%v' is a 'nil' pointer.\n",
			ePrefix.String(),
			fIoWriterLabel,
			fIoWriterLabel)

		return numBytesWritten, err
	}

	if fIoWriter.ioWriter == nil {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of FileIoWriter ('%v') is invalid!\n"+
			"The internal io.Writer object has NOT been properly initialized.\n"+
			"%v.writer is equal to 'nil'.\n"+
			"Call one of the 'New' or 'Setter' methods to create a valid instance of 'FileIoWriter'.\n",
			ePrefix.String(),
			fIoWriterLabel,
			fIoWriterLabel)

		return numBytesWritten, err
	}

	lenBytesToWrite := len(bytesToWrite)

	if lenBytesToWrite == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'bytesToWrite' is invalid!\n"+
			"The 'bytesToWrite' byte array is empty. It contains zero bytes.\n",
			ePrefix.String())

		return numBytesWritten, err
	}

	var writer io.Writer
	var err2 error
	writer = *fIoWriter.ioWriter

	numBytesWritten,
		err2 = writer.Write(bytesToWrite)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: writer.Write(bytesToWrite)\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			err2.Error())

		return numBytesWritten, err
	}

	if lenBytesToWrite != numBytesWritten {

		err = fmt.Errorf("%v\n"+
			"Error: The number of bytes provided by\n"+
			"input parameter 'bytesToWrite' is NOT equal\n"+
			"to the number of bytes actually written.\n"+
			"         Length 'bytesToWrite' = %v\n"+
			"Actual Number of Bytes Written = %v\n",
			ePrefix.String(),
			lenBytesToWrite,
			numBytesWritten)

	}

	return numBytesWritten, err
}

type fileIoWriterNanobot struct {
	lock *sync.Mutex
}

// setIoWriter
//
// This 'setter' method is used to initialize new values
// for internal member variables contained in the
// instance of FileIoWriter passed as input parameter
// 'fIoWriter'. The new configuration will be based on
// an io.Writer object passed as input parameter
// 'writer'.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://pkg.go.dev/io
//	https://pkg.go.dev/io#Writer
//	https://pkg.go.dev/io#Writer
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	This method will delete, overwrite and reset all
//		pre-existing data values in the instance of
//		FileIoWriter passed as input parameter
//		'fIoWriter'.
//
//	(2) The instance of FileIoWriter passed as 'fIoWriter'
//		will retain a pointer reference to input parameter
//		'writer'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fIoWriter					*FileIoWriter
//
//		A pointer to an instance of FileIoWriter.
//
//		All internal member variable data values in
//		this instance will be deleted and initialized
//		to new values based on the following input
//		parameters.
//
//	fIoWriterLabel				string
//
//		The name or label associated with input parameter
//		'fIoWriter' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fIoWriter" will be
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
//		file and performing any other required Clean-Up
//		operations in addition to calling local method
//		FileIoWriter.Close().
//
//		While the configured instance of FileIoWriter
//		(fIoWriter) is primarily designed for writing
//		data to disk files, this 'writer' will in fact
//		write data to any object implementing the
//		io.Writer interface.
//
//		This instance of io.Writer will be used to
//		configure the internal io.Writer contained in
//		'fIoWriter' and used to conduct 'write'
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
//	defaultByteArraySize		int
//
//		The size of the byte array which will be used to
//		write data to the internal io.Writer object
//		encapsulated by FileIoWriter input parameter
//		'fIoWriter'.
//
//		If the value of 'defaultByteArraySize' is
//		less than one ('1'), it will be reset to a size
//		of '4096'.
//
//		Although the FileIoWriter type does not use the
//		'buffered' write protocol, the size of the byte
//		array used to write bytes to the underlying
//		io.Writer object is variable.
//
//		Methods utilizing the Default Byte Array Size
//		include:
//
//			FileIoWriter.ReadFrom()
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
func (fIoWriterNanobot *fileIoWriterNanobot) setIoWriter(
	fIoWriter *FileIoWriter,
	fIoWriterLabel string,
	writer io.Writer,
	writerLabel string,
	defaultByteArraySize int,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fIoWriterNanobot.lock == nil {
		fIoWriterNanobot.lock = new(sync.Mutex)
	}

	fIoWriterNanobot.lock.Lock()

	defer fIoWriterNanobot.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileIoWriterNanobot." +
		"setIoWriter()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return err
	}

	if len(fIoWriterLabel) == 0 {

		fIoWriterLabel = "fIoWriter"
	}

	if len(writerLabel) == 0 {

		writerLabel = "writer"
	}

	if fIoWriter == nil {

		err = fmt.Errorf("%v\n"+
			"Error: The FileIoWriter instance passed\n"+
			"as input parameter '%v' is invalid!\n"+
			"'%v' is a 'nil' pointer.\n",
			ePrefix.String(),
			fIoWriterLabel,
			fIoWriterLabel)

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

	var fIoWriterMolecule = new(fileIoWriterMolecule)

	// Close the old fIoWriter
	err = fIoWriterMolecule.closeAndRelease(
		fIoWriter,
		writerLabel,
		true,
		ePrefix)

	if err != nil {

		return err
	}

	// Set Default Byte Array Size

	fIoWriter.defaultByteArraySize =
		defaultByteArraySize

	fIoWriterMolecule.
		validateDefaultByteArraySize(fIoWriter)

	// Set io.Writer
	fIoWriter.ioWriter = &writer

	// Set File Data
	var ok bool

	var fileMgr *FileMgr

	fileMgr, ok = writer.(*FileMgr)

	if ok == true {

		fIoWriter.targetWriteFileName =
			fileMgr.GetAbsolutePathFileName()

		fIoWriter.filePtr = fileMgr.filePtr

	} else {

		fIoWriter.filePtr, ok = writer.(*os.File)

		if ok == true {

			fIoWriter.targetWriteFileName =
				fIoWriter.filePtr.Name()

		}

	}

	return err
}

// setPathFileName
//
// This 'setter' method is used to initialize new values
// for internal member variables contained in the
// instance of FileIoWriter passed as input parameter
// 'fIoWriter'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the instance of
//	FileIoWriter passed as input parameter
//	'fIoWriter'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fIoWriter					*FileIoWriter
//
//		A pointer to an instance of FileIoWriter.
//
//		All internal member variable data values in
//		this instance will be deleted and initialized
//		to new values based on the following input
//		parameters.
//
//	fIoWriterLabel				string
//
//		The name or label associated with input parameter
//		'fIoWriter' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fIoWriter" will be
//		automatically applied.
//
//	pathFileName				string
//
//		This string contains the path and file name of
//		the file which will be used a data source for
//		'write' operations.
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
//		'write' file identified from input parameter
//		'pathFileName' will be opened for both 'write'
//		and 'write' operations.
//
//		If 'openFileReadWrite' is set to 'false', the
//		target 'write' file will be opened for 'write-only'
//		operations.
//
//	defaultByteArraySize		int
//
//		The size of the byte array which will be used to
//		write data to the internal io.Writer object
//		encapsulated by FileIoWriter input parameter
//		'fIoWriter'.
//
//		If the value of 'defaultByteArraySize' is
//		less than one ('1'), it will be reset to a size
//		of '4096'.
//
//		Although the FileIoWriter type does not use the
//		'buffered' write protocol, the size of the byte
//		array used to write bytes to the underlying
//		io.Writer object is variable.
//
//		Methods utilizing the Default Byte Array Size
//		include:
//
//			FileIoWriter.ReadFrom()
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
func (fIoWriterNanobot *fileIoWriterNanobot) setPathFileName(
	fIoWriter *FileIoWriter,
	fIoWriterLabel string,
	pathFileName string,
	pathFileNameLabel string,
	openFileReadWrite bool,
	defaultByteArraySize int,
	truncateExistingFile bool,
	errPrefDto *ePref.ErrPrefixDto) (
	fInfoPlus FileInfoPlus,
	err error) {

	if fIoWriterNanobot.lock == nil {
		fIoWriterNanobot.lock = new(sync.Mutex)
	}

	fIoWriterNanobot.lock.Lock()

	defer fIoWriterNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileIoWriterNanobot." +
		"setPathFileName()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		return fInfoPlus, err
	}

	if len(fIoWriterLabel) == 0 {

		fIoWriterLabel = "fIoWriter"
	}

	if fIoWriter == nil {

		err = fmt.Errorf("%v\n"+
			"Error: The FileIoWriter instance passed\n"+
			"as input parameter '%v' is invalid!\n"+
			"'%v' is a 'nil' pointer.\n",
			ePrefix.String(),
			fIoWriterLabel,
			fIoWriterLabel)

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

	var fIoWriterMolecule = new(fileIoWriterMolecule)

	// Close the old fIoWriter
	_ = fIoWriterMolecule.closeAndRelease(
		fIoWriter,
		fIoWriterLabel,
		true,
		ePrefix)

	var err2 error

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
			"An error occurred while testing for the existence\n"+
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

	var filePermissionCfg FilePermissionConfig
	var fileOpenPermissions = "--w--w--w-"

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

	fIoWriter.filePtr,
		err = new(fileHelperBoson).
		openFile(
			pathFileName,
			false,
			fileOpenCfg,
			filePermissionCfg,
			pathFileNameLabel,
			ePrefix)

	if err != nil {

		if fIoWriter.filePtr != nil {
			_ = fIoWriter.filePtr.Close()
		}

		return fInfoPlus, err
	}

	fIoWriter.targetWriteFileName = pathFileName

	var writer io.Writer

	writer = fIoWriter.filePtr

	fIoWriter.ioWriter = &writer

	fIoWriter.defaultByteArraySize =
		defaultByteArraySize

	fIoWriterMolecule.
		validateDefaultByteArraySize(fIoWriter)

	return fInfoPlus, err
}

type fileIoWriterMolecule struct {
	lock *sync.Mutex
}

// closeAndRelease
//
// This method will perform the 'close' procedure on the
// internal io.Writer object encapsulated by the instance
// of FileIoWriter passed as input parameter 'fIoWriter'.
//
// In addition, this method will also release the
// internal memory resources encapsulated by 'fIoWriter'.
//
// This method, fileIoReaderMolecule.closeAndRelease(),
// is the recommended method for 'closing' the internal
// io.Writer object.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	(1)	This method will close the underlying io.Writer
//		object encapsulated by 'fIoWriter'.
//
//	(2)	In addition, this method will release the
//		internal memory	resources for the 'fIoWriter'
//		object.
//
//	(3)	After completion of this method, the FileIoWriter
//		instance passed by 'fIoWriter' will be unusable,
//		invalid and unavailable	for future 'write'
//		operations.
//
//	(4)	If an error is returned by the 'close' operation,
//		this method will still proceed to release all
//		internal memory resources.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fIoWriter					*FileIoWriter
//
//		A pointer to an instance of FileIoWriter.
//
//		This method will close the underlying io.Writer
//		object contained in 'fIoWriter' and release
//		all internal memory resources.
//
//		Upon completion, 'fIoWriter' will be rendered
//		invalid, and unusable for all future 'write'
//		operations.
//
//	fIoWriterLabel				string
//
//		The name or label associated with input parameter
//		'fIoWriter' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fIoWriter" will be
//		automatically applied.
//
//	releaseMemoryResources		bool
//
//		If this parameter is set to 'true', this method
//		will release all internal memory resources for
//		the passed instance of FileIoWriter ('fIoWriter').
//		Releasing internal memory resources synchronizes
//		internal flags and prevents multiple calls to the
//		'close' method. Calling the 'close' method more
//		than once may produce unexpected results.
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
//	 	text passed by input parameter, 'errorPrefix'.
//	 	The 'errorPrefix' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fIoWriterMolecule *fileIoWriterMolecule) closeAndRelease(
	fIoWriter *FileIoWriter,
	fIoWriterLabel string,
	releaseMemoryResources bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fIoWriterMolecule.lock == nil {
		fIoWriterMolecule.lock = new(sync.Mutex)
	}

	fIoWriterMolecule.lock.Lock()

	defer fIoWriterMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	funcName := "fileIoWriterMolecule." +
		"closeAndRelease()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		return err
	}

	if len(fIoWriterLabel) == 0 {

		fIoWriterLabel = "fIoWriter"
	}

	if fIoWriter == nil {

		err = fmt.Errorf("%v\n"+
			"Error: The FileIoWriter instance passed\n"+
			"as input parameter '%v' is invalid!\n"+
			"'%v' is a 'nil' pointer.\n",
			ePrefix.String(),
			fIoWriterLabel,
			fIoWriterLabel)

		return err
	}

	var fIoWriterAtom = new(fileIoWriterAtom)

	err = fIoWriterAtom.close(
		fIoWriter,
		fIoWriterLabel,
		ePrefix)

	if err != nil {

		return err
	}

	if releaseMemoryResources == true {
		fIoWriterAtom.empty(fIoWriter)
	}

	return err
}

// validateDefaultByteArraySize
//
// Validates the Default Byte Array Size internal
// member variable for the FileIoWriter instance passed
// as input parameter 'fIoWriter'.
//
//	FileIoWriter.defaultByteArraySize
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fIoWriter					*FileIoWriter
//
//		A pointer to an instance of FileIoWriter.
//
//		The default buffer size for this FileIoWriter
//		instance will be validated. If an invalid value
//		is detected that value will be automatically
//		reset to a value of 4096-bytes.
//
//		This internal member variable is styled as:
//
//			FileIoWriter.defaultByteArraySize
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	-- NONE --
func (fIoWriterMolecule *fileIoWriterMolecule) validateDefaultByteArraySize(
	fIoWriter *FileIoWriter) {

	if fIoWriterMolecule.lock == nil {
		fIoWriterMolecule.lock = new(sync.Mutex)
	}

	fIoWriterMolecule.lock.Lock()

	defer fIoWriterMolecule.lock.Unlock()

	if fIoWriter == nil {

		return
	}

	if fIoWriter.defaultByteArraySize < 1 {

		fIoWriter.defaultByteArraySize = 4096
	}

	return
}

type fileIoWriterAtom struct {
	lock *sync.Mutex
}

// close
//
// This method will effectively close the underlying
// io.Writer object contained in the FileIoWriter
// instance passed as input parameter 'fIoWriter'.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	This method will NOT release the internal memory
//	resources for the 'fIoWriter' instance. To release
//	the internal memory resources, call this method:
//
//			fileIoWriterAtom.empty()
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	After completion of this method th3 FileIoWriter
//	instance passed as 'fIoWriter' will be unusable,
//	invalid and unavailable for future 'write'
//	operations.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fIoWriter					*FileIoWriter
//
//		A pointer to an instance of FileIoWriter.
//
//		The underlying io.Writer object	configured for
//		'fIoWriter' will be closed.
//
//	fIoWriterLabel				string
//
//		The name or label associated with input parameter
//		'fIoWriter' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fIoWriter" will be
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
func (fIoWriterAtom *fileIoWriterAtom) close(
	fIoWriter *FileIoWriter,
	fIoWriterLabel string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fIoWriterAtom.lock == nil {
		fIoWriterAtom.lock = new(sync.Mutex)
	}

	fIoWriterAtom.lock.Lock()

	defer fIoWriterAtom.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileIoWriterAtom." +
		"close()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return err
	}

	if len(fIoWriterLabel) == 0 {

		fIoWriterLabel = "fIoWriter"
	}

	if fIoWriter == nil {

		err = fmt.Errorf("%v\n"+
			"-------------------------------------------\n"+
			"Error: The FileIoWriter instance passed\n"+
			"as input parameter '%v' is invalid!\n"+
			"'%v' is a 'nil' pointer.\n",
			ePrefix.String(),
			fIoWriterLabel,
			fIoWriterLabel)

		return err
	}

	if fIoWriter.ioWriter == nil {

		return err
	}

	var ok bool
	var closerObj io.Closer
	var localWriter io.Writer

	localWriter = *fIoWriter.ioWriter

	closerObj, ok = localWriter.(io.Closer)

	if ok {

		var err2 error

		err2 = closerObj.Close()

		if err2 != nil {

			errText := fmt.Sprintf(
				"%v\n"+
					"Error returned while closing the 'FileIoWriter'\n"+
					"internal io.Writer object.\n",
				ePrefix.String())

			if len(fIoWriter.targetWriteFileName) > 0 {

				errText += fmt.Sprintf(
					"Target Write File Name: %v\n",
					fIoWriter.targetWriteFileName)

			}

			err = fmt.Errorf("%v"+
				"closerObj.Close() Error=\n%v\n",
				errText,
				err2.Error())

		}

	}

	return err
}

// empty
//
// This method releases all internal member resources
// contained in the FileIoWriter instance passed as input
// parameter 'fIoWriter'.
//
// Specifically, the following internal member variables
// are set to 'nil' or their initial zero values:
//
//	fIoWriter.targetReadFileName = ""
//	fIoWriter.filePtr = nil
//	fIoWriter.ioReader = nil
//	fIoWriter.defaultByteArraySize = 0
//
// After calling this method, the current instance of
// FileIoWriter will become invalid and unavailable
// for future 'write' operations.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	This method will NOT 'close' the underlying io.Writer
//	object encapsulated within 'fIoWriter'. Best
//	practices call for 'closing' the underlying io.Writer
//	object first and then releasing the internal memory
//	resources with a call to this method,
//	fileIoReaderAtom.empty().
//
//	To 'close' the underlying io.Writer object, first
//	call method, fileIoWriterAtom.close().
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fIoWriter					*FileIoWriter
//
//		A pointer to an instance of FileIoWriter.
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
func (fIoWriterAtom *fileIoWriterAtom) empty(
	fIoWriter *FileIoWriter) {

	if fIoWriterAtom.lock == nil {
		fIoWriterAtom.lock = new(sync.Mutex)
	}

	fIoWriterAtom.lock.Lock()

	defer fIoWriterAtom.lock.Unlock()

	if fIoWriter == nil {

		return
	}

	fIoWriter.ioWriter = nil

	fIoWriter.filePtr = nil

	fIoWriter.targetWriteFileName = ""

	fIoWriter.defaultByteArraySize = 0
}

// isFileIoWriterValid
//
// This method receives a pointer to an instance of
// FileIoWriter ('fIoWriter') which will be analyzed to
// determine if all the member variables contain valid
// data values.
//
// If input parameter 'fIoWriter' is determined to be
// invalid, this method returns an error.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fIoWriter				*FileIoWriter
//
//		A pointer to an instance of FileIoWriter.
//
//		If any of the internal member variable data
//		values encapsulated in 'fIoWriter' are
//		determined to be invalid, this method will return
//		an error.
//
//	fIoWriterLabel			string
//
//		The name or label associated with input parameter
//		'fIoWriter' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fIoWriter" will
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
//		contained in the instance of FileIoWriter
//		passed as 'fIoWriter' are found to be
//		invalid, this method will return an error
//		configured with an appropriate message
//		identifying the invalid	member data variable.
//
//		If all internal member data variables evaluate
//		as valid, this returned error value will be set
//		to 'nil'.
//
//		If errors are encountered during processing or if
//		any 'fIoWriter' internal member data values
//		are found to be invalid, the returned error Type
//		will encapsulate an appropriate error message.
//	 	This returned error message will incorporate the
//	 	method chain and text passed by input parameter,
//	 	'errorPrefix'. The 'errorPrefix' text will be
//	 	prefixed to the beginning of the error message.
func (fIoWriterAtom *fileIoWriterAtom) isFileIoWriterValid(
	fIoWriter *FileIoWriter,
	fIoWriterLabel string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fIoWriterAtom.lock == nil {
		fIoWriterAtom.lock = new(sync.Mutex)
	}

	fIoWriterAtom.lock.Lock()

	defer fIoWriterAtom.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileIoWriterAtom."+
			"isFileIoWriterValid()",
		"")

	if err != nil {
		return err
	}

	if len(fIoWriterLabel) == 0 {

		fIoWriterLabel = "fIoWriter"
	}

	if fIoWriter == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is a nil pointer!\n"+
			"%v is invalid.\n",
			ePrefix.String(),
			fIoWriterLabel,
			fIoWriterLabel)

		return err
	}

	if fIoWriter.ioWriter == nil {

		err = fmt.Errorf("%v\n"+
			" -------------------------------------------------------------------\n"+
			" ERROR: The %v instance of FileIoWriter is invalid!\n"+
			" The internal io.Writer object was never initialized.\n"+
			" Call one of the 'New' methods or 'Setter' methods to create\n"+
			" a valid instance of FileIoWriter.\n",
			ePrefix.String(),
			fIoWriterLabel)
	}

	return err
}
