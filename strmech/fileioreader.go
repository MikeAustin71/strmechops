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
// The most common data source for these 'read'
// operations is typically a data file residing on an
// attached storage drive. However, any object
// implementing the io.Reader interface may be used as a
// data source.
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
//	(1)	Pointer receiver FileIoReader methods implement
//		the following interfaces:
//
//			io.Reader
//			io.Closer
//			io.Seeker
//			io.ReaderAt
//			io.ReadCloser
//			io.ReadSeeker
//			io.ReadSeekCloser
//			io.WriterTo
//
//	(2)	When all read operations have been completed and
//		there is no further need for the FileIoReader
//		instance, the user is responsible for 'closing'
//		the underlying io.Reader object and releasing the
//		associated memory resources. This Clean-Up
//		operation is accomplished by calling local
//		method:
//
//			FileIoReader.Close()
//
//		For FileIoReader instances created with an
//		external io.Reader object, the user may need to
//		'close' that io.Reader object and perform
//		Clean-Up operations externally.
type FileIoReader struct {
	ioReader             *io.Reader
	filePtr              *os.File
	targetReadFileName   string
	defaultByteArraySize int
	lock                 *sync.Mutex
}

// Close
//
// This method is provided in order to implement the
// io.Closer interface.
//
// FileIoReader.Close() effectively performs all required
// Clean-Up tasks. As such, this method should only be
// called after all 'read' operations have been completed
// and the services of the current FileIoReader instance
// are no longer required.
//
// After calling this method, FileIoReader.Close(), the
// current instance of FileIoReader will be invalid and
// unavailable for further 'read' operations.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	This method implements the io.Closer interface.
//
//	(2)	This method will perform all required Clean-Up
//		tasks. FileIoReader.Close() should only be called
//		after all 'read' operations have been completed.
//
//	(3)	Clean-Up tasks performed by this method include
//		'closing' the underlying io.Reader object as well
//		as releasing all internal memory resources.
//
//		Releasing all internal memory resources will
//		synchronize internal flags and prevent multiple
//		calls to 'close' the underlying io.Reader object.
//		Calling 'close' on the same underlying io.Reader
//		object multiple times can produce unexpected
//		results.
//
//	(4)	After completing the call to this method,
//		FileIoReader.Close(), the current FileIoReader
//		instance will become invalid and unavailable
//		for future 'read' operations.
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

	err = new(fileIoReaderMolecule).closeAndRelease(
		fIoReader,
		"fIoReader",
		true, // releaseMemoryResources
		ePrefix.XCpy("fIoReader"))

	return err
}

// Empty
//
// This method deletes all internal member variables and
// releases all the internal memory resources for the
// current instance of FileIoReader.
//
// Specifically the following internal object pointers
// are set to nil or their initial zero values:
//
//	FileIoReader.targetReadFileName = ""
//	FileIoReader.filePtr = nil
//	FileIoReader.ioReader = nil
//	FileIoReader.defaultByteArraySize = 0
//
// After calling this method, the current instance of
// FileIoReader will become invalid and unavailable
// for future 'read' operations.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	This method does NOT perform the 'close' protocol. To
//	perform both the 'close' protocol and release all
//	internal memory resources call local method:
//
//			FileIoReader.Close()
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
func (fIoReader *FileIoReader) Empty() {

	if fIoReader.lock == nil {
		fIoReader.lock = new(sync.Mutex)
	}

	fIoReader.lock.Lock()

	new(fileIoReaderAtom).empty(
		fIoReader)

	fIoReader.lock.Unlock()

	fIoReader.lock = nil

	return
}

// GetDefaultByteArraySize
//
// Returns the default byte array size for the current
// instance of FileIoReader.
//
// The Default Reader Buffer Size controls the size of
// the byte array used by the following methods:
//
//	FileIoReader.ReadAllToStrBuilder()
//	FileIoReader.ReadAllToString()
//	FileIoReader.WriteTo()
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
//		length for the current instance of FileIoReader.
func (fIoReader *FileIoReader) GetDefaultByteArraySize() int {

	if fIoReader.lock == nil {
		fIoReader.lock = new(sync.Mutex)
	}

	fIoReader.lock.Lock()

	defer fIoReader.lock.Unlock()

	new(fileIoReaderMolecule).
		validateDefaultReaderArraySize(
			fIoReader)

	return fIoReader.defaultByteArraySize
}

// GetIoReader
//
// Returns the internal io.Reader object encapsulated by
// the current instance of FileIoReader.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method returns the internal io.Reader object
//	configured for the current instance of FileIoReader.
//
//	Be sure to release this io.Reader object when it
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
//	reader						io.Reader
//
//		The concrete instance of io.Reader encapsulated
//		by the current instance of FileIoReader.
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
func (fIoReader *FileIoReader) GetIoReader(
	errorPrefix interface{}) (
	reader io.Reader,
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
			"GetIoReader()",
		"")

	if err != nil {

		return reader, err
	}

	if fIoReader.ioReader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of 'FileIoReader' (%v) is invalid!\n"+
			"The internal '%v' io.Reader object has NOT been initialized.\n"+
			"Call one of the 'New' or 'Setter' methods when creating\n"+
			"an instance of 'FileIoReader'\n",
			ePrefix.String(),
			"fIoReader",
			"fIoReader")

		return reader, err
	}

	reader = *fIoReader.ioReader

	return reader, err
}

// IsValidInstanceError
//
// Analyzes the current FileIoReader instance to
// determine if is invalid.
//
// If the current FileIoReader instance is found
// to be invalid, an error is returned.
//
// If the current FileIoReader instance is valid,
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
//		FileIoReader are found to be invalid,
//		this method will return an error configured
//		with an appropriate message identifying the
//		invalid	member data variable.
//
//		If all internal member data variables evaluate
//		as valid, this returned error value will be set
//		to 'nil'.
//
//		If errors are encountered during processing or if
//		any FileIoReader internal member data
//	 	values are found to be invalid, the returned error
//	 	will encapsulate an appropriate error message.
//	 	This returned error message will incorporate the
//	 	method chain and text passed by input parameter,
//	 	'errorPrefix'. The 'errorPrefix' text will be
//	 	prefixed to the beginning of the error message.
func (fIoReader *FileIoReader) IsValidInstanceError(
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
			"IsValidInstanceError()",
		"")

	if err != nil {
		return err
	}

	return new(fileIoReaderAtom).
		isFileIoReaderValid(
			fIoReader,
			"fIoReader",
			ePrefix)
}

// New
//
// This method returns a pointer to a new, empty instance
// of FileIoReader.
//
// After creating an empty instance of FileIoReader,
// the user may call the 'Setter' methods to configure
// this FileIoReader instance with valid parameters.
//
// Step-1
//
//	Call this method FileIoReader.New() to
//	generate an empty version of FileIoReader.
//
// Step-2
//
//	Use this returned instance of FileIoReader and
//	call the appropriate 'Setter' methods to individually
//	configure the internal 'io.reader' object.
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
//	*FileIoReader
//
//		This method returns a pointer to an empty
//		instance of	FileIoReader.
func (fIoReader *FileIoReader) New() *FileIoReader {

	if fIoReader.lock == nil {
		fIoReader.lock = new(sync.Mutex)
	}

	fIoReader.lock.Lock()

	defer fIoReader.lock.Unlock()

	return new(FileIoReader)
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
// # IMPORTANT
//
//	(1)	The returned instance of FileIoReader does NOT
//		use buffered read techniques. Instead, it
//		implements a direct read protocol using
//		io.Reader.
//
//	(2)	Input parameter 'reader' will accept a pointer to
//		an instance of os.File because os.File implements
//		the io.Reader interface.
//
//	(3) The returned instance of FileIoReader will retain
//		a pointer reference to input parameter 'reader'.
//		Be sure to close and release this pointer when the
//		returned instance of FileIoReader is no longer
//		needed. To perform the 'close' operation, call the
//		local method:
//
//			FileIoReader.Close()
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
//		performing any other required Clean-Up operations
//		in addition to calling FileIoReader.Close().
//
//		While the returned instance of FileIoReader
//		is primarily designed for reading data from disk
//		files, this 'reader' will in fact read data from
//		any object implementing the io.Reader interface.
//
//	defaultReaderByteArraySize		int
//
//		The size of the byte array which will be used to
//		read data from the internal io.Reader object
//		encapsulated by the returned FileIoReader
//		instance.
//
//		If the value of 'defaultReaderByteArraySize' is
//		less than '16', it will be reset to a size of
//		'4096'.
//
//		Although the FileIoReader type does not use the
//		'buffered' read protocol, the size of the byte
//		array used to store bytes read from the
//		underlying io.Reader object is variable.
//
//		Methods utilizing the Default Reader Buffer Size
//		include:
//
//			FileIoReader.ReadAllToStrBuilder()
//			FileIoReader.ReadAllToString()
//			FileIoReader.WriteTo()
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
//		If this method completes successfully, a pointer
//		to a fully configured instance of FileIoReader
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
func (fIoReader *FileIoReader) NewIoReader(
	reader io.Reader,
	defaultReaderByteArraySize int,
	errorPrefix interface{}) (
	*FileIoReader,
	error) {

	if fIoReader.lock == nil {
		fIoReader.lock = new(sync.Mutex)
	}

	fIoReader.lock.Lock()

	defer fIoReader.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error
	var newFileIoReader = new(FileIoReader)

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
			newFileIoReader,
			"newFileIoReader",
			reader,
			"reader",
			defaultReaderByteArraySize,
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
//	defaultReaderByteArraySize		int
//
//		The size of the byte array which will be used to
//		read data from the internal io.Reader object
//		encapsulated by the returned FileIoReader
//		instance.
//
//		If the value of 'defaultReaderByteArraySize' is less
//		than '16', it will be reset to a size of '4096'.
//
//		Although the FileIoReader type does not use the
//		'buffered' read protocol, the size of the byte
//		array used to read and store bytes read from the
//		underlying io.Reader object is variable.
//
//		Methods utilizing the Default Reader Buffer Size
//		include:
//
//			FileIoReader.ReadAllToStrBuilder()
//			FileIoReader.ReadAllToString()
//			FileIoReader.WriteTo()
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
//	*FileIoReader
//
//		If this method completes successfully, a pointer
//		to a fully configured instance of FileIoReader
//		will be returned.
//
//		This returned instance of FileIoReader will
//		configure the file identified by input parameter
//		'fileMgr' is a data source for file 'read'
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
func (fIoReader *FileIoReader) NewFileMgr(
	fileMgr *FileMgr,
	openFileReadWrite bool,
	defaultReaderByteArraySize int,
	errorPrefix interface{}) (
	FileInfoPlus,
	*FileIoReader,
	error) {

	if fIoReader.lock == nil {
		fIoReader.lock = new(sync.Mutex)
	}

	fIoReader.lock.Lock()

	defer fIoReader.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error
	var fInfoPlus FileInfoPlus
	var newFileIoReader = new(FileIoReader)

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
			newFileIoReader,
			"newFileIoReader",
			fileMgr,
			"fileMgr",
			openFileReadWrite,
			defaultReaderByteArraySize,
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
// io.Reader.
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
//	defaultReaderByteArraySize		int
//
//		The size of the byte array which will be used to
//		read data from the internal io.Reader object
//		encapsulated by the returned FileIoReader
//		instance.
//
//		If the value of 'defaultReaderByteArraySize' is
//		less than '16', it will be reset to a size of
//		'4096'.
//
//		Although the FileIoReader type does not use the
//		'buffered' read protocol, the size of the byte
//		array used to read and store bytes read from the
//		underlying io.Reader object is variable.
//
//		Methods utilizing the Default Reader Buffer Size
//		include:
//
//			FileIoReader.ReadAllToStrBuilder()
//			FileIoReader.ReadAllToString()
//			FileIoReader.WriteTo()
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
func (fIoReader *FileIoReader) NewPathFileName(
	pathFileName string,
	openFileReadWrite bool,
	defaultReaderByteArraySize int,
	errorPrefix interface{}) (
	FileInfoPlus,
	*FileIoReader,
	error) {

	if fIoReader.lock == nil {
		fIoReader.lock = new(sync.Mutex)
	}

	fIoReader.lock.Lock()

	defer fIoReader.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error
	var fInfoPlus FileInfoPlus
	var newFileIoReader = new(FileIoReader)

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
			newFileIoReader,
			"newFileIoReader",
			pathFileName,
			"pathFileName",
			openFileReadWrite,
			defaultReaderByteArraySize,
			ePrefix.XCpy(
				pathFileName))

	return fInfoPlus, newFileIoReader, err
}

// Read
//
// Reads a selection of data from the internal io.Reader
// data source encapsulated in the current instance of
// FileIoReader.
//
// This method reads data into the input parameter byte
// array, 'bytesRead', from the internal io.Reader object
// encapsulated by the current instance of FileIoReader.
// The number of bytes read into the byte array is
// returned as return parameter, 'numBytesRead'.
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
// is responsible for performing Clean-Up operations by
// calling the local method:
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
//	(5)	This method employs the direct 'read' technique
//		used by type io.Reader. It does NOT use the
//		buffered 'read' technique employed by the
//		bufio.Reader type.
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
//		Bytes will be read from the input data source and
//		stored in this byte array.
//
//		The input data source was previously configured
//		in the current instance of FileIoReader.
//
//		If the length of this byte array is less than
//		16-bytes, an error will be returned.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	numBytesRead				int
//
//		If this method completes successfully, the number
//		of bytes read from the data source
//		'FileIoReader.ioReader' and stored in the byte
//		array 'bytesRead', will be returned through this
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
	numBytesRead int,
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

		return numBytesRead, err
	}

	numBytesRead,
		err = new(fileIoReaderMicrobot).
		readBytes(
			fIoReader,
			"fIoReader",
			bytesRead,
			ePrefix)

	return numBytesRead, err
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
// Be advised that the StringArrayDto type includes
// methods for adding custom end of line delimiters.
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
//	(2)	This method may automatically close the io.Reader
//		object upon completion depending on the setting
//		for input parameter, 'autoCloseOnExit'.
//
//		If this method does NOT automatically close the
//		io.Reader object upon completion, the user will
//		be responsible for performing the Clean-Up tasks
//		by calling local method:
//
//			FileIoReader.Close()
//
//	(3)	If the current instance of FileIoReader has NOT
//		been properly initialized, an error will be
//		returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	readEndOfLineDelimiters		*StringArrayDto
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
//		Typical text line termination, or end-of-line
//		delimiters, which may be appropriate for use
//		with a given target 'read' file are listed as
//		follows:
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
//	outputLinesArray 			*StringArrayDto
//
//		A pointer to an instance of StringArrayDto.
//		Lines of text read from the file specified
//		by 'pathFileName' will be stored as
//		individual strings in the string array
//		encapsulated by 'outputLinesArray'.
//
//		-------------------------------------------------
//					IMPORTANT
//		-------------------------------------------------
//		The line termination or end-of-line delimiter
//		characters identified from 'endOfLineDelimiters'
//		will be stripped off and deleted from the end of
//		each line of text stored in the string array
//		encapsulated by 'outputLinesArray'. As such, the
//		text lines stored here are pure strings of text
//		without any line termination or end-of-line
//		delimiter characters append to the end of the
//		string.
//
//	maxNumOfTextLines			int
//
//		Specifies the maximum number of text lines which
//		will be read from the file identified by
//		'pathFileName'.
//
//		If 'maxNumOfLines' is set to a value less than
//		zero (0) (Example: minus-one (-1) ),
//		'maxNumOfLines' will be automatically reset to
//		math.MaxInt(). This means all text lines existing
//		in the file identified by 'pathFileName' will be
//		read and processed. Reading all the text lines in
//		a file 'may' have memory implications depending
//		on the size of the file and the memory resources
//		available to your computer.
//
//		If 'maxNumOfLines' is set to a value of zero
//		('0'), no text lines will be read from the file
//		identified by 'pathFileName', and no error will be
//		returned.
//
//	autoCloseOnExit				bool
//
//		When this parameter is set to 'true', this
//		method will automatically perform all required
//		Clean-Up tasks upon exit. These tasks include
//		closing the underlying FileIoReader io.Reader
//		object and releasing all internal memory
//		resources.
//
//		(1)	The FileIoReader internal io.Reader object
//			will be properly closed, the internal memory
//			resources will be released and there will be
//			no need to make a separate call to local
//			method,	FileIoReader.Close().
//
//		(2) After performing this Clean-Up operation, the
//			current instance of FileIoReader will invalid
//			and unavailable for future 'read' operations.
//
//		-------------------------------------------------
//						Be Advised
//		If processing errors are encountered during
//		method execution, the 'close' operation WILL NOT
//		be invoked or applied.
//		-------------------------------------------------
//
//		If input parameter 'autoCloseOnExit' is set to
//		'false', this method will NOT automatically
//		'close' the internal io.Reader object and release
//		the memory resources for the current instance of
//		FileBufferReader. Consequently, the user will then
//		be responsible for 'closing' the internal io.Reader
//		object by calling the local method:
//
//				FileIoReader.Close()
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
//	numLinesRead				int
//
//		This integer value contains the number of text
//		lines read from the file specified by input
//		parameter 'pathFileName'. This value also
//		specifies the number of array elements added to
//		the string array encapsulated by
//		'outputLinesArray'.
//
//		When displayed in editors, the end-of-file
//		character is displayed on a separate line.
//		The returned 'numLinesRead' value does
//		not include this empty line containing an
//		end-of-file character. Therefore, the
//		returned 'numLinesRead' value will always
//		be one less than the number of lines shown
//		in a text editor.
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
	readEndOfLineDelimiters *StringArrayDto,
	outputLinesArray *StringArrayDto,
	maxNumOfTextLines int,
	autoCloseOnExit bool,
	errorPrefix interface{}) (
	numLinesRead int,
	numBytesRead int64,
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

		return numLinesRead,
			numBytesRead,
			err
	}

	if fIoReader.ioReader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of 'FileIoReader' is invalid!\n"+
			"The internal io.Reader object has NOT been initialized.\n"+
			"Call one of the 'New' or 'Setter' methods when creating\n"+
			"an instance of 'FileIoReader'\n",
			ePrefix.String())

		return numLinesRead,
			numBytesRead,
			err
	}

	numLinesRead,
		numBytesRead,
		err = new(fileHelperMolecule).
		readerScanLines(
			*fIoReader.ioReader,
			"fIoReader.ioReader",
			readEndOfLineDelimiters,
			outputLinesArray,
			maxNumOfTextLines,
			ePrefix.XCpy("fIoReader.ioReader"))

	if autoCloseOnExit == true &&
		err == nil {

		// Do NOT Close if there is an error!
		err = new(fileIoReaderMolecule).closeAndRelease(
			fIoReader,
			"fIoReader",
			true, // releaseMemoryResources
			ePrefix.XCpy("fIoReader"))

	}

	return numLinesRead,
		numBytesRead,
		err
}

// ReadAllToStrBuilder
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
//	(2)	This method may automatically close the io.Reader
//		object upon completion depending on the setting
//		for input parameter, 'autoCloseOnExit'.
//
//		If this method does NOT automatically close the
//		io.Reader object upon completion, the user will
//		be responsible for performing the Clean-Up tasks
//		by calling local method:
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
//	autoCloseOnExit				bool
//
//		When this parameter is set to 'true', this
//		method will automatically perform all required
//		Clean-Up tasks upon exit. These tasks include
//		closing the underlying FileIoReader io.Reader
//		object and releasing all internal memory
//		resources.
//
//		(1)	The FileIoReader internal io.Reader object
//			will be properly closed, the internal memory
//			resources will be released and there will be
//			no need to make a separate call to local
//			method,	FileIoReader.Close().
//
//		(2) After performing this Clean-Up operation, the
//			current instance of FileIoReader will invalid
//			and unavailable for future 'read' operations.
//
//		-------------------------------------------------
//						Be Advised
//		If processing errors are encountered during
//		method execution, the 'close' operation WILL NOT
//		be invoked or applied.
//		-------------------------------------------------
//
//		If input parameter 'autoCloseOnExit' is set to
//		'false', this method will NOT automatically
//		'close' the internal io.Reader object and release
//		the memory resources for the current instance of
//		FileBufferReader. Consequently, the user will then
//		be responsible for 'closing' the internal io.Reader
//		object by calling the local method:
//
//				FileIoReader.Close()
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
//		operation, the returned error object will be set
//		to 'nil' and no error will be returned.
func (fIoReader *FileIoReader) ReadAllToStrBuilder(
	strBuilder *strings.Builder,
	autoCloseOnExit bool,
	errorPrefix interface{}) (
	numBytesRead int64,
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
			"ReadAllToStrBuilder()",
		"")

	if err != nil {

		return numBytesRead, err
	}

	numBytesRead,
		err = new(fileIoReaderMicrobot).
		readAllStrBuilder(
			fIoReader,
			"fIoReader",
			strBuilder,
			autoCloseOnExit,
			ePrefix)

	return numBytesRead, err
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
//	(2)	This method may automatically close the io.Reader
//		object upon completion depending on the setting
//		for input parameter, 'autoCloseOnExit'.
//
//		If this method does NOT automatically close the
//		io.Reader object upon completion, the user will
//		be responsible for performing the Clean-Up tasks
//		by calling local method:
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
//	autoCloseOnExit				bool
//
//		When this parameter is set to 'true', this
//		method will automatically perform all required
//		Clean-Up tasks upon exit. These tasks include
//		closing the underlying FileIoReader io.Reader
//		object and releasing all internal memory
//		resources.
//
//		(1)	The FileIoReader internal io.Reader object
//			will be properly closed, the internal memory
//			resources will be released and there will be
//			no need to make a separate call to local
//			method,	FileIoReader.Close().
//
//		(2) After performing this Clean-Up operation, the
//			current instance of FileIoReader will invalid
//			and unavailable for future 'read' operations.
//
//		-------------------------------------------------
//						Be Advised
//		If processing errors are encountered during
//		method execution, the 'close' operation WILL NOT
//		be invoked or applied.
//		-------------------------------------------------
//
//		If input parameter 'autoCloseOnExit' is set to
//		'false', this method will NOT automatically
//		'close' the internal io.Reader object and release
//		the memory resources for the current instance of
//		FileBufferReader. Consequently, the user will then
//		be responsible for 'closing' the internal io.Reader
//		object by calling the local method:
//
//				FileIoReader.Close()
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
//		operation, the returned error object will be set
//		to 'nil' and no error will be returned.
func (fIoReader *FileIoReader) ReadAllToString(
	autoCloseOnExit bool,
	errorPrefix interface{}) (
	numBytesRead int64,
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

		return numBytesRead, contentsStr, err
	}

	strBuilder := new(strings.Builder)

	numBytesRead,
		err = new(fileIoReaderMicrobot).
		readAllStrBuilder(
			fIoReader,
			"fIoReader",
			strBuilder,
			autoCloseOnExit,
			ePrefix)

	if err != nil {

		return numBytesRead, contentsStr, err
	}

	contentsStr = strBuilder.String()

	return numBytesRead, contentsStr, err
}

// ReadAt
//
// This method reads bytes beginning at the offset from
// the beginning of the input source as specified by
// input parameter 'offsetFromBeginning'.
//
// ----------------------------------------------------------------
//
// # Reference
//
// https://pkg.go.dev/io#ReaderAt
//
// ----------------------------------------------------------------
//
// # Be Advised
//
//	This method implements the io.ReaderAt interface.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	bytesRead					[]byte
//
//		Bytes will be read from the input data source and
//		stored in this byte array.
//
//		The input data source was previously configured
//		in the current instance of FileIoReader.
//
//		If the length of this byte array is less than
//		16-bytes, an error will be returned.
//
//	offsetFromBeginning			int64
//
//		The offset in bytes from the beginning of the
//		internal io.Reader input source at which the read
//		'operation' will commence.
//
//		If this value is less than zero, an error will be
//		returned.
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
//	numBytesRead				int
//
//		If this method completes successfully, this
//		returned integer will hold the number of bytes
//		read successfully from the internal io.Reader
//		object encapsulated by the current instance of
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
func (fIoReader *FileIoReader) ReadAt(
	bytesRead []byte,
	offsetFromBeginning int64) (
	numBytesRead int,
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
			"ReadAt()",
		"")

	if err != nil {

		return numBytesRead, err
	}

	if fIoReader.ioReader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of 'FileIoReader' is invalid!\n"+
			"The internal io.Reader object has NOT been initialized.\n"+
			"Call one of the 'New' or 'Setter' methods when creating\n"+
			"an instance of 'FileIoReader'\n",
			ePrefix.String())

		return numBytesRead, err
	}

	if len(bytesRead) < 16 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'bytesRead' is invalid!\n"+
			"The 'bytesRead' array has a length less than 16.\n"+
			"Length 'bytesRead'= '%v'\n",
			ePrefix.String(),
			len(bytesRead))

		return numBytesRead, err
	}

	if offsetFromBeginning < 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'offsetFromBeginning' is invalid!\n"+
			"The 'bytesRead' array has a length less than zero ('0).\n"+
			"offsetFromBeginning= '%v'\n",
			ePrefix.String(),
			offsetFromBeginning)

		return numBytesRead, err
	}

	var ok bool
	var readerAtObj io.ReaderAt
	var localReader io.Reader
	localReader = *fIoReader.ioReader

	readerAtObj, ok = localReader.(io.ReaderAt)

	if !ok {

		err = fmt.Errorf("%v\n"+
			"Error: This ReadAt method was invoked on a\n"+
			"'FileIoReader' internal io.Reader object\n"+
			"which does NOT support the io.ReaderAt\n"+
			"interface. This means:\n"+
			"(1) The 'ReadAt' method is unavailable.\n"+
			"              AND"+
			"(2) The 'FileIoReader' internal io.Reader\n"+
			"    was created from an object which does\n"+
			"    NOT implement the io.WriterAt interface.\n",
			ePrefix.String())

		return numBytesRead, err
	}

	numBytesRead,
		err = readerAtObj.ReadAt(
		bytesRead,
		offsetFromBeginning)

	return numBytesRead, err
}

// ReadBytesToString
//
// Reads a specified number of bytes from the io.Reader
// data input source configured for the current instance
// of FileIoReader (FileIoReader.ioReader). These bytes
// are then stored as a string and returned via return
// parameter 'contentsStr'.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	(1)	This method may automatically close the io.Reader
//		object upon completion depending on the setting
//		for input parameter, 'autoCloseOnExit'.
//
//		If this method does NOT automatically close the
//		io.Reader object upon completion, the user will
//		be responsible for performing the Clean-Up tasks
//		by calling local method:
//
//			FileIoReader.Close()
//
//	(2)	If the current instance of FileIoReader has NOT
//		been properly initialized, an error will be
//		returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numOfBytesToRead			int
//
//		This parameter specifies the number of bytes to
//		read from the data input source configured for
//		the current instance of FileIoReader.
//
//		If the value of 'numOfBytesToRead' is less than
//		one ('1'), this parameter will be automatically
//		set to the Default Reader Buffer Size previously
//		configured for this FileIoReader instance. For
//		more information on Default Reader Buffer Size,
//		reference local method:
//
//			FileIoReader.SetDefaultByteArraySize()
//
//		The actual number of bytes read from the data input
//		source may vary due to (1) unforeseen processing
//		errors or (2) an End-Of-File scenario. Be sure to
//		check the 'numBytesRead' and 'reachedEndOfFile'
//		parameters returned by this method.
//
//	autoCloseOnExit				bool
//
//		When this parameter is set to 'true', this
//		method will automatically perform all required
//		Clean-Up tasks upon exit. These tasks include
//		closing the underlying FileIoReader io.Reader
//		object and releasing all internal memory
//		resources.
//
//		(1)	The FileIoReader internal io.Reader object
//			will be properly closed, the internal memory
//			resources will be released and there will be
//			no need to make a separate call to local
//			method,	FileIoReader.Close().
//
//		(2) After performing this Clean-Up operation, the
//			current instance of FileIoReader will invalid
//			and unavailable for future 'read' operations.
//
//		-------------------------------------------------
//						Be Advised
//		If processing errors are encountered during
//		method execution, the 'close' operation WILL NOT
//		be invoked or applied.
//		-------------------------------------------------
//
//		If input parameter 'autoCloseOnExit' is set to
//		'false', this method will NOT automatically
//		'close' the internal io.Reader object and release
//		the memory resources for the current instance of
//		FileBufferReader. Consequently, the user will then
//		be responsible for 'closing' the internal io.Reader
//		object by calling the local method:
//
//				FileIoReader.Close()
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
//		integer value will equal the actual number of
//		bytes read from the input data source (io.Reader)
//		encapsulated by the current instance of
//		FileIoReader and stored in the strings.Builder
//		instance passed as input parameter 'strBuilder'.
//
//		This actual number of bytes read from the data
//		input source may vary from the 'numOfBytesToRead'
//		input parameter due to (1) unforeseen processing
//		errors or (2) an End-Of-File scenario. Be sure to
//		check the 'reachedEndOfFile' parameter returned
//		by this method.
//
//	contentsStr					string
//
//		If this method completes successfully, the bytes
//		read from the input data source configured for
//		the	current instance of FileIoReader will be
//		returned in this string.
//
//	reachedEndOfFile			bool
//
//		If during the 'read' operation, the End-Of-File
//		flag was encountered, this boolean parameter will
//		be set to 'true'. The End-Of-File flag signals that
//		the 'read' operation reached the end of the data
//		input source configured for the current
//		FileIoReader instance.
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
//		end-of-file flag is encountered during the 'read'
//		operation, 'reachedEndOfFile' will be set to
//		'true', the returned error object will be set
//		to 'nil', and no error will be returned.
func (fIoReader *FileIoReader) ReadBytesToString(
	numOfBytesToRead int,
	autoCloseOnExit bool,
	errorPrefix interface{}) (
	numBytesRead int64,
	contentsStr string,
	reachedEndOfFile bool,
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
			"ReadBytesToString()",
		"")

	if err != nil {

		return numBytesRead,
			contentsStr,
			reachedEndOfFile,
			err
	}

	if fIoReader.ioReader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of 'FileIoReader' is invalid!\n"+
			"The internal io.Reader object has NOT been initialized.\n"+
			"Call one of the 'New' or 'Setter' methods when creating\n"+
			"an instance of 'FileIoReader'\n",
			ePrefix.String())

		return numBytesRead,
			contentsStr,
			reachedEndOfFile,
			err
	}

	var strBuilder = new(strings.Builder)

	numBytesRead,
		reachedEndOfFile,
		err = new(fileIoReaderMicrobot).
		readBytesToStrBuilder(
			fIoReader,
			"fIoReader",
			numOfBytesToRead,
			strBuilder,
			autoCloseOnExit,
			ePrefix)

	if err != nil {
		contentsStr = strBuilder.String()
	}

	return numBytesRead,
		contentsStr,
		reachedEndOfFile,
		err
}

// ReadBytesToStringBuilder
//
// Reads a specified number of bytes from the data input
// source configured for the current instance of
// FileIoReader. These bytes are then stored in an
// instance of strings.Builder passed as input parameter
// 'strBuilder'.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	(1)	This method may automatically close the io.Reader
//		object upon completion depending on the setting
//		for input parameter, 'autoCloseOnExit'.
//
//		If this method does NOT automatically close the
//		io.Reader object upon completion, the user will
//		be responsible for performing the Clean-Up tasks
//		by calling local method:
//
//			FileIoReader.Close()
//
//	(2)	If the current instance of FileIoReader has NOT
//		been properly initialized, an error will be
//		returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numOfBytesToRead			int
//
//		This parameter specifies the number of bytes to
//		read from the data input source configured for
//		the current instance of FileIoReader.
//
//		If the value of 'numOfBytesToRead' is less than
//		one ('1'), this parameter will be automatically
//		set to the Default Reader Buffer Size previously
//		configured for this FileIoReader instance. For
//		more information on Default Reader Buffer Size,
//		reference local method:
//
//			FileIoReader.SetDefaultByteArraySize()
//
//		The actual number of bytes read from the data input
//		source may vary due to (1) unforeseen processing
//		errors or (2) an End-Of-File scenario. Be sure to
//		check the 'numBytesRead' and 'reachedEndOfFile'
//		parameters returned by this method.
//
//	strBuilder					*strings.Builder
//
//		A pointer to an instance of strings.Builder. This
//		method reads bytes from the input data source
//		configured for the current instance of
//		FileIoReader and adds those bytes to this
//		instance of strings.Builder.
//
//	autoCloseOnExit				bool
//
//		When this parameter is set to 'true', this
//		method will automatically perform all required
//		Clean-Up tasks upon exit. These tasks include
//		closing the underlying FileIoReader io.Reader
//		object and releasing all internal memory
//		resources.
//
//		(1)	The FileIoReader internal io.Reader object
//			will be properly closed, the internal memory
//			resources will be released and there will be
//			no need to make a separate call to local
//			method,	FileIoReader.Close().
//
//		(2) After performing this Clean-Up operation, the
//			current instance of FileIoReader will invalid
//			and unavailable for future 'read' operations.
//
//		-------------------------------------------------
//						Be Advised
//		If processing errors are encountered during
//		method execution, the 'close' operation WILL NOT
//		be invoked or applied.
//		-------------------------------------------------
//
//		If input parameter 'autoCloseOnExit' is set to
//		'false', this method will NOT automatically
//		'close' the internal io.Reader object and release
//		the memory resources for the current instance of
//		FileBufferReader. Consequently, the user will then
//		be responsible for 'closing' the internal io.Reader
//		object by calling the local method:
//
//				FileIoReader.Close()
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
//		integer value will equal the actual number of
//		bytes read from the input data source (io.Reader)
//		encapsulated by the current instance of
//		FileIoReader and stored in the strings.Builder
//		instance passed as input parameter 'strBuilder'.
//
//		This actual number of bytes read from the data
//		input source may vary from the 'numOfBytesToRead'
//		input parameter due to (1) unforeseen processing
//		errors or (2) an End-Of-File scenario. Be sure to
//		check the 'reachedEndOfFile' parameter returned
//		by this method.
//
//	reachedEndOfFile			bool
//
//		If during the 'read' operation, the End-Of-File
//		flag was encountered, this boolean parameter will
//		be set to 'true'. The End-Of-File flag signals that
//		the 'read' operation reached the end of the data
//		input source configured for the current
//		FileIoReader instance.
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
//		end-of-file flag is encountered during the 'read'
//		operation, 'reachedEndOfFile' will be set to
//		'true', the returned error object will be set
//		to 'nil', and no error will be returned.
func (fIoReader *FileIoReader) ReadBytesToStringBuilder(
	numOfBytesToRead int,
	strBuilder *strings.Builder,
	autoCloseOnExit bool,
	errorPrefix interface{}) (
	numBytesRead int64,
	reachedEndOfFile bool,
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
			"ReadBytesToStringBuilder()",
		"")

	if err != nil {

		return numBytesRead, reachedEndOfFile, err
	}

	if fIoReader.ioReader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of 'FileIoReader' is invalid!\n"+
			"The internal io.Reader object has NOT been initialized.\n"+
			"Call one of the 'New' or 'Setter' methods when creating\n"+
			"an instance of 'FileIoReader'\n",
			ePrefix.String())

		return numBytesRead, reachedEndOfFile, err
	}

	numBytesRead,
		reachedEndOfFile,
		err = new(fileIoReaderMicrobot).
		readBytesToStrBuilder(
			fIoReader,
			"fIoReader",
			numOfBytesToRead,
			strBuilder,
			autoCloseOnExit,
			ePrefix)

	return numBytesRead, reachedEndOfFile, err
}

// Seek
//
// This method sets the byte offset for the next 'read'
// operation within the internal io.Reader object
// encapsulated in the current FileIoReader instance.
//
// The Seek() method only succeeds if the internal
// io.Reader object, encapsulated by the current
// FileIoReader instance, implements the io.Seeker
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
// 'read' operation will occur at the new offset
// position specified by input parameter 'targetOffset'.
//
// Method Seek() returns the new offset relative to the
// io.Reader object or an error, if any.
//
// Method Seek() implements the 'io.Seeker' interface.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	If the current instance of FileIoReader was
//		NOT initialized with an io.Reader object which
//		implements the io.Seeker interface, this method
//		will return an error.
//
//	(2)	Setting a byte offset which occurs before the
//		start of the internal io.Reader object
//		encapsulated by the current FileIoReader instance
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
func (fIoReader *FileIoReader) Seek(
	targetOffset int64,
	whence int) (
	offsetFromFileStart int64,
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
			"Seek()",
		"")

	if err != nil {

		return offsetFromFileStart, err
	}

	offsetFromFileStart,
		err = new(fileIoReaderMicrobot).
		seekByOffset(
			fIoReader,
			"fIoReader",
			targetOffset,
			whence,
			ePrefix)

	return offsetFromFileStart, err
}

// SetDefaultByteArraySize
//
// Sets the default size of the array used to read bytes
// from the internal io.Reader encapsulated in the
// current instance of FileIoReader.
//
// Although the FileIoReader type does not use the
// 'buffered' read protocol, the size of the byte array
// used to read and store bytes read from the underlying
// io.Reader object is variable in some cases.
//
// The Default Reader Buffer Size controls the size of
// the byte array used by the following methods:
//
//	FileIoReader.ReadAllToStrBuilder()
//	FileIoReader.ReadAllToString()
//	FileIoReader.WriteTo()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	defaultReaderByteArraySize		int
//
//		The size of the byte array which will be used to
//		read data from the internal io.Reader object
//		encapsulated by the current instance of
//		FileIoReader.
//
//		If the value of 'defaultReaderByteArraySize' is less
//		than '16', it will be automatically reset to a
//		size of '4096'.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	-- NONE --
func (fIoReader *FileIoReader) SetDefaultByteArraySize(
	defaultReaderByteArraySize int) {

	if fIoReader.lock == nil {
		fIoReader.lock = new(sync.Mutex)
	}

	fIoReader.lock.Lock()

	defer fIoReader.lock.Unlock()

	fIoReader.defaultByteArraySize =
		defaultReaderByteArraySize

	new(fileIoReaderMolecule).
		validateDefaultReaderArraySize(
			fIoReader)

	return
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
//		performing any other required Clean-Up operations
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
//	defaultReaderByteArraySize		int
//
//		The size of the byte array which will be used to
//		read data from the internal io.Reader object
//		encapsulated by the current FileIoReader
//		instance.
//
//		If the value of 'defaultReaderByteArraySize' is
//		less than '16', it will be reset to a size of
//		'4096'.
//
//		Although the FileIoReader type does not use the
//		'buffered' read protocol, the size of the byte
//		array used to read and store bytes read from the
//		underlying io.Reader object is variable.
//
//		Methods utilizing the Default Reader Buffer Size
//		include:
//
//			FileIoReader.ReadAllToStrBuilder()
//			FileIoReader.ReadAllToString()
//			FileIoReader.WriteTo()
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
	defaultReaderByteArraySize int,
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
			defaultReaderByteArraySize,
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
// reconfigure the internal io.Reader encapsulated in
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
//	defaultReaderByteArraySize		int
//
//		The size of the byte array which will be used to
//		read data from the internal io.Reader object
//		encapsulated by the current FileIoReader
//		instance.
//
//		If the value of 'defaultReaderByteArraySize' is
//		less than '16', it will be reset to a size of
//		'4096'.
//
//		Although the FileIoReader type does not use the
//		'buffered' read protocol, the size of the byte
//		array used to read and store bytes read from the
//		underlying io.Reader object is variable.
//
//		Methods utilizing the Default Reader Buffer Size
//		include:
//
//			FileIoReader.ReadAllToStrBuilder()
//			FileIoReader.ReadAllToString()
//			FileIoReader.WriteTo()
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
	defaultReaderByteArraySize int,
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
			defaultReaderByteArraySize,
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
// will be used to reconfigure the internal io.Reader
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
//	defaultReaderByteArraySize		int
//
//		The size of the byte array which will be used to
//		read data from the internal io.Reader object
//		encapsulated by the current FileIoReader
//		instance.
//
//		If the value of 'defaultReaderByteArraySize' is
//		less than '16', it will be reset to a size of
//		'4096'.
//
//		Although the FileIoReader type does not use the
//		'buffered' read protocol, the size of the byte
//		array used to read and store bytes read from the
//		underlying io.Reader object is variable.
//
//		Methods utilizing the Default Reader Buffer Size
//		include:
//
//			FileIoReader.ReadAllToStrBuilder()
//			FileIoReader.ReadAllToString()
//			FileIoReader.WriteTo()
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
	defaultReaderByteArraySize int,
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
			defaultReaderByteArraySize,
			ePrefix.XCpy(
				pathFileName))

	return fInfoPlus, err
}

// WriteTo
//
// This method implements the io.ReaderTo interface.
//
// Input parameter 'writer' passes an io.Writer object.
// This method will then proceed to read the entire
// contents of the io.Reader object encapsulated by
// the current instance of FileIoReader and write
// this data to the io.Writer object passed as input
// parameter 'writer'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	This method implements the io.WriterTo interface.
//
//	(2) When the current instance of FileIoReader is no
//		longer needed, the user is responsible for
//		performing 'close' and Clean-Up operations by
//		calling the local method:
//
//			FileIoReader.Close()
//
//	(3) If the number of bytes read from the FileIoReader
//		internal io.Reader object does NOT match the bytes
//		written to the io.Writer object passed as input
//		parameter 'writer', an error will be returned.
//
//	(4) The size of the internal byte array used to read
//		store and write bytes is controlled by the
//		FileIoReader internal member variable:
//
//			FileIoReader.defaultByteArraySize
//
//		Reference local method:
//
//			FileIoReader.SetDefaultReaderByteArraySize()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	writer						io.Writer
//
//		This instance of io.Writer will be used as the
//		'write' destination for all data read from the
//		io.Reader object encapsulated by the current
//		instance of FileIoReader.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	numOfBytesProcessed			int64
//
//		The number of bytes read from the internal
//		FileIoReader io.Reader object and written to the
//		destination io.Writer object passed as input
//		parameter 'writer'.
//
//	err							error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message.
func (fIoReader *FileIoReader) WriteTo(
	writer io.Writer) (
	numOfBytesProcessed int64,
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
		"FileBufferReader."+
			"WriteTo()",
		"")

	if err != nil {

		return numOfBytesProcessed, err
	}

	if fIoReader.ioReader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of 'FileIoReader' is invalid!\n"+
			"The internal io.Reader object has NOT been initialized.\n"+
			"Call one of the 'New' or 'Setter' methods to create a\n"+
			"valid instance of 'FileIoReader'\n",
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

	new(fileIoReaderMolecule).
		validateDefaultReaderArraySize(
			fIoReader)

	var bytesRead = make([]byte,
		fIoReader.defaultByteArraySize)

	var localBytesRead, localBytesWritten int
	var err2, err3 error
	var localReader = *fIoReader.ioReader

	for {

		localBytesRead,
			err2 = localReader.Read(bytesRead)

		if err2 != nil &&
			err2 != io.EOF {

			err = fmt.Errorf("%v\n"+
				"Error: localReader.Read(bytesRead).\n"+
				"The Read operation failed with errors.\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				err2.Error())

			return numOfBytesProcessed, err
		}

		if localBytesRead > 0 {

			localBytesWritten,
				err3 = writer.Write(bytesRead[:localBytesRead])

			if err3 != nil {

				err = fmt.Errorf("\n%v\n"+
					"Error: writer.Write(bytesRead[:localBytesRead])\n"+
					"The Write operation failed with errors.\n"+
					"Error=\n%v\n",
					ePrefix.String(),
					err3.Error())

				return numOfBytesProcessed, err
			}

			if localBytesWritten != localBytesRead {

				err = fmt.Errorf("%v\n"+
					"Error: The number of bytes written\n"+
					"DOES NOT equal the number of bytes read!\n"+
					"   Bytes Read: %v\n"+
					"Bytes Written: %v\n",
					ePrefix.String(),
					localBytesRead,
					localBytesWritten)

				return numOfBytesProcessed, err

			} else {

				numOfBytesProcessed += int64(localBytesRead)
			}

		}

		if err2 == io.EOF {

			break
		}

	}

	return numOfBytesProcessed, err
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
//	autoCloseOnExit				bool
//
//		When this parameter is set to 'true', this
//		method will automatically perform all required
//		FileIoReader Clean-Up tasks upon exit for
//		'fIoReader'.
//
//		(1)	The FileIoReader internal io.Reader object
//			will be properly closed and there will be no
//			need to make a separate call to local method,
//			FileIoReader.Close().
//
//		(2) After performing this Clean-Up operation, the
//			current instance of FileIoReader will invalid
//			and unavailable for future 'read' operations.
//
//		-------------------------------------------------
//						Be Advised
//		If processing errors are encountered during method
//		execution, the 'close' operation WILL NOT be
//		invoked or applied.
//		-------------------------------------------------
//
//		If input parameter 'autoCloseOnExit' is set to
//		'false', this method will NOT automatically
//		'close' the internal io.Reader object for the
//		'fIoReader' FileBufferReader. Consequently, the
//		user will then be responsible for performing
//		required Clean-Up tasks on the FileIoReader
//		instance by calling the local method:
//
//				FileIoReader.Close()
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
//	reachedEndOfFile			bool
//
//		If during the 'read' operation, the End-Of-File
//		flag was encountered, this boolean parameter will
//		be set to 'true'. The End-Of-File flag signals that
//		the 'read' operation reached the end of the data
//		input source configured for the 'fIoReader'
//		FileIoReader instance.
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
	autoCloseOnExit bool,
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

	if fIoReader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is invalid!\n"+
			"'%v' is a 'nil' pointer.\n",
			ePrefix.String(),
			fIoReaderLabel,
			fIoReaderLabel)

		return numOfBytesRead, err
	}

	if strBuilder == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'strBuilder' is invalid!\n"+
			"'strBuilder' is a 'nil' pointer.\n",
			ePrefix.String())

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

	var fIoReaderMolecule = new(fileIoReaderMolecule)

	fIoReaderMolecule.
		validateDefaultReaderArraySize(
			fIoReader)

	var reader = *fIoReader.ioReader
	var bytesRead = make([]byte, fIoReader.defaultByteArraySize)
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

			return numOfBytesRead, err
		}

		clear(bytesRead)

	}

	if autoCloseOnExit == true {

		// Do NOT Close if there is an error!
		err = fIoReaderMolecule.closeAndRelease(
			fIoReader,
			"fIoReader",
			true, // releaseMemoryResources
			ePrefix.XCpy("fIoReader"))

	}

	return numOfBytesRead, err
}

// readBytes
//
// Reads a selection of data from the internal io.Reader
// data source encapsulated in the FileIoReader instance
// passed as input parameter 'fIoReader'.
//
// This method is a wrapper for the 'io.Reader.Read'
// method.
//
// This method reads data into the input parameter byte
// array, 'bytesRead', from the internal io.Reader object
// encapsulated by 'fIoReader'. The number of bytes read
// into the byte array is returned as return parameter,
// 'numBytesRead'.
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
// is responsible for performing Clean-Up operations by
// calling the local method:
//
//	FileIoReader.Close()
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	Keep calling this method until all the bytes have
//		been read from the io.Reader data source
//		configured for the FileIoReader instance
//		'fIoReader' and the returned error is set to
//		'io.EOF'.
//
//	(2)	Callers should always process the returned number
//		of bytes read ('numBytesRead') before considering
//		the returned error parameter 'err'. Doing so
//		correctly handles I/O errors that occur after
//		reading some number of bytes. Also, this
//		technique allows both possible EOF behaviors to
//		be correctly processed (See the io.Reader docs
//		and 'Reference' section below).
//
//	(3)	When all 'read' operations have been completed,
//		perform the required Clean-Up operations by
//		calling local method:
//
//			FileIoReader.Close()
//
//	(4)	This method employs the direct 'read' technique
//		used by type io.Reader. It does NOT use the
//		buffered 'read' technique employed by the
//		bufio.Reader type.
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
//	fIoReader					*FileIoReader
//
//		A pointer to an instance of FileIoReader.
//
//		A number of bytes equal to, or less than, the
//		length of byte array 'bytesRead' will be read
//		from the internal io.Reader object encapsulated
//		in this FileIoReader instance and stored in
//		byte array 'bytesRead'.
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
//	bytesRead					[]byte
//
//		Bytes will be read from the io.Reader data source
//		encapsulated in 'fIoReader' and stored in this
//		byte array.
//
//		If the length of this byte array is less than
//		16-bytes, an error will be returned.
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
//	numBytesRead				int
//
//		If this method completes successfully, the number
//		of bytes read from the data source
//		'fIoReader.ioReader' and stored in the byte array
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
func (fIoReaderMicrobot *fileIoReaderMicrobot) readBytes(
	fIoReader *FileIoReader,
	fIoReaderLabel string,
	bytesRead []byte,
	errPrefDto *ePref.ErrPrefixDto) (
	numBytesRead int,
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
			"readBytes()",
		"")

	if err != nil {

		return numBytesRead, err
	}

	if len(fIoReaderLabel) == 0 {

		fIoReaderLabel = "fIoReader"
	}

	if fIoReader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is invalid!\n"+
			"'%v' is a 'nil' pointer.\n",
			ePrefix.String(),
			fIoReaderLabel,
			fIoReaderLabel)

		return numBytesRead, err
	}

	if fIoReader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: The FileIoReader instance passed\n"+
			"as input parameter '%v' is invalid!\n"+
			"'%v' is a 'nil' pointer.\n",
			ePrefix.String(),
			fIoReaderLabel,
			fIoReaderLabel)

		return numBytesRead, err
	}

	if fIoReader.ioReader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of FileIoReader ('%v') is invalid!\n"+
			"The internal io.Reader object has NOT been properly initialized.\n"+
			"%v.writer is equal to 'nil'.\n"+
			"Call one of the 'New' or 'Setter' methods to create a valid instance of 'FileIoReader'.\n",
			ePrefix.String(),
			fIoReaderLabel,
			fIoReaderLabel)

		return numBytesRead, err
	}

	lenBytesRead := len(bytesRead)

	if lenBytesRead < 16 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'bytesRead' is invalid!\n"+
			"The length of the 'bytesRead' byte array is less than 16-bytes.\n"+
			"Length 'bytesRead' = %v\n",
			ePrefix.String(),
			lenBytesRead)

		return numBytesRead, err
	}

	var err2 error

	// Produces concrete io.Reader reference
	var reader = *fIoReader.ioReader

	numBytesRead,
		err2 = reader.Read(bytesRead)

	if err2 != nil {

		if err2 != io.EOF {

			err = fmt.Errorf("%v\n"+
				"Error returned by %v.ioReader.Read(bytesRead).\n"+
				"Error=\n%v\n",
				ePrefix.String(),
				fIoReaderLabel,
				err2.Error())

			return numBytesRead, err

		} else {

			err = io.EOF

		}

	}

	return numBytesRead, err
}

// readBytesToStrBuilder
//
// Reads a specified number of bytes from the data input
// source configured for the instance of FileIoReader
// passed as input parameter 'fIoReader'. These bytes are
// then stored in an instance of strings.Builder passed
// as input parameter 'strBuilder'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fIoReader					*FileIoReader
//
//		A pointer to an instance of FileIoReader.
//
//		A number of bytes specified by input parameter
//		'numOfBytesToRead' will be read from the data
//		input source (io.Reader) configured for this
//		instance of FileIoReader.
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
//	numOfBytesToRead			int
//
//		This parameter specifies the number of bytes to
//		read from the data input source configured for
//		the current instance of FileIoReader.
//
//		If the value of 'numOfBytesToRead' is less than
//		one ('1'), this parameter will be automatically
//		set to the Default Reader Buffer Size previously
//		configured for the 'fIoReader' FileIoReader
//		instance. For more information on Default Reader
//		Buffer Size, reference local method:
//
//			FileIoReader.SetDefaultByteArraySize()
//
//		The actual number of bytes read from the data input
//		source may vary due to (1) unforeseen processing
//		errors or (2) an End-Of-File scenario. Be sure to
//		check the 'numOfBytesRead' and 'reachedEndOfFile'
//		parameters returned by this method.
//
//	strBuilder					*strings.Builder
//
//		A pointer to an instance of strings.Builder.
//		Bytes read from the input data source configured
//		for the 'fIoReader'	FileIoReader instance will be
//		extracted and stored as a string in 'strBuilder'.
//
//	autoCloseOnExit				bool
//
//		When this parameter is set to 'true', this
//		method will automatically perform all required
//		FileIoReader Clean-Up tasks upon exit for
//		'fIoReader'.
//
//		(1)	The FileIoReader internal io.Reader object
//			will be properly closed and there will be no
//			need to make a separate call to local method,
//			FileIoReader.Close().
//
//		(2) After performing this Clean-Up operation, the
//			current instance of FileIoReader will invalid
//			and unavailable for future 'read' operations.
//
//		-------------------------------------------------
//						Be Advised
//		If processing errors are encountered during method
//		execution, the 'close' operation WILL NOT be
//		invoked or applied.
//		-------------------------------------------------
//
//		If input parameter 'autoCloseOnExit' is set to
//		'false', this method will NOT automatically
//		'close' the internal io.Reader object for the
//		'fIoReader' FileBufferReader. Consequently, the
//		user will then be responsible for performing
//		required Clean-Up tasks on the FileIoReader
//		instance by calling the local method:
//
//				FileIoReader.Close()
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
//	numOfBytesRead				int64
//
//		If this method completes successfully, this
//		integer value will equal the actual number of
//		bytes read from the input data source (io.Reader)
//		encapsulated by the 'fIoReader' and stored in the
//		strings.Builder instance passed as input
//		parameter 'strBuilder'.
//
//		This actual number of bytes read from the data
//		input source may vary from the 'numOfBytesToRead'
//		input parameter due to (1) unforeseen processing
//		errors or (2) an End-Of-File scenario. Be sure to
//		check the 'reachedEndOfFile' parameter returned
//		by this method.
//
//	reachedEndOfFile			bool
//
//		If during the 'read' operation, the End-Of-File
//		flag was encountered, this boolean parameter will
//		be set to 'true'. The End-Of-File flag signals that
//		the 'read' operation reached the end of the data
//		input source configured for the 'fIoReader'
//		FileIoReader instance.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
//
//		An error will only be returned if a processing
//		or system error was encountered. When the
//		end-of-file is encountered during the 'read'
//		process, the returned error object will be set
//		to 'nil', 'reachedEndOfFile' will be set to
//		'true' and no error will be returned.
func (fIoReaderMicrobot *fileIoReaderMicrobot) readBytesToStrBuilder(
	fIoReader *FileIoReader,
	fIoReaderLabel string,
	numOfBytesToRead int,
	strBuilder *strings.Builder,
	autoCloseOnExit bool,
	errPrefDto *ePref.ErrPrefixDto) (
	numOfBytesRead int64,
	reachedEndOfFile bool,
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

		return numOfBytesRead, reachedEndOfFile, err
	}

	if len(fIoReaderLabel) == 0 {

		fIoReaderLabel = "fIoReader"
	}

	if fIoReader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is invalid!\n"+
			"The FileIoReader instance '%v' is a 'nil' pointer.\n",
			ePrefix.String(),
			fIoReaderLabel,
			fIoReaderLabel)

		return numOfBytesRead, reachedEndOfFile, err
	}

	if strBuilder == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'strBuilder' is invalid!\n"+
			"'strBuilder' is a 'nil' pointer.\n",
			ePrefix)

		return numOfBytesRead, reachedEndOfFile, err
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

		return numOfBytesRead, reachedEndOfFile, err
	}

	var fIoReaderMolecule = new(fileIoReaderMolecule)

	fIoReaderMolecule.
		validateDefaultReaderArraySize(
			fIoReader)

	if numOfBytesToRead < 1 {

		numOfBytesToRead = fIoReader.defaultByteArraySize

	}

	bytesToRead := make([]byte, numOfBytesToRead)

	var err2 error
	var reader = *fIoReader.ioReader
	var localNumBytesRead int

	localNumBytesRead,
		err2 = reader.Read(bytesToRead)

	if err2 != nil {

		if err2 != io.EOF {

			err = fmt.Errorf("%v\n"+
				"Error returned by fIoReader.ioReader.Read(bytesRead).\n"+
				"Error=\n%v\n",
				ePrefix.String(),
				err2.Error())

			return numOfBytesRead, reachedEndOfFile, err

		} else {

			// Must be io.EOF
			reachedEndOfFile = true
		}

	}

	strBuilder.Write(bytesToRead[:localNumBytesRead])

	numOfBytesRead = int64(localNumBytesRead)

	if autoCloseOnExit == true {

		// Do NOT Close if there is an error!
		err = fIoReaderMolecule.closeAndRelease(
			fIoReader,
			"fIoReader",
			true, // releaseMemoryResources
			ePrefix.XCpy("fIoReader"))

	}

	return numOfBytesRead, reachedEndOfFile, err
}

// seekByOffset
//
// This method sets the byte offset for the next 'read'
// operation within the internal io.Reader object
// encapsulated in the FileIoReader instance passed as
// input parameter 'fIoReader'.
//
// This method only succeeds if the internal io.Reader
// object implements the io.Seeker interface. Disk
// files with a base type of os.File and FileMgr
// instances are among those types which implement the
// io.Seeker interface.
//
// The new byte offset (targetOffset) is interpreted
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
//	(1)	Setting a byte offset at a point before the start
//		of the file is an error.
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
func (fIoReaderMicrobot *fileIoReaderMicrobot) seekByOffset(
	fIoReader *FileIoReader,
	fIoReaderLabel string,
	targetOffset int64,
	whence int,
	errPrefDto *ePref.ErrPrefixDto) (
	offsetFromFileStart int64,
	err error) {

	if fIoReaderMicrobot.lock == nil {
		fIoReaderMicrobot.lock = new(sync.Mutex)
	}

	fIoReaderMicrobot.lock.Lock()

	defer fIoReaderMicrobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileIoReaderMicrobot." +
		"seekOffset()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		return offsetFromFileStart, err
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

		return offsetFromFileStart, err
	}

	if fIoReader.ioReader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of FileIoReader ('%v') is invalid!\n"+
			"The internal io.Reader object has NOT been initialized.\n"+
			"%v.ioReader is equal to 'nil'.\n"+
			"Call one of the 'New' or 'Setter' methods to create a\n"+
			"valid instance of 'FileIoReader'\n",
			ePrefix.String(),
			fIoReaderLabel,
			fIoReaderLabel)

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
	var localReader io.Reader
	localReader = *fIoReader.ioReader

	seekerObj, ok = localReader.(io.Seeker)

	if !ok {

		err = fmt.Errorf("%v\n"+
			"Error: This Seek method was invoked on a\n"+
			"'FileIoReader' internal io.Reader object\n"+
			"which does NOT support the io.Seeker\n"+
			"interface. This means:\n"+
			"(1) The 'Seek' method is unavailable.\n"+
			"           AND\n"+
			"(2) The 'FileIoReader' internal io.Reader\n"+
			"      object was created from something\n"+
			"      other than a disk file (*os.File)\n"+
			"      or FileMgr type.\n",
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
//	defaultReaderByteArraySize		int
//
//		The size of the byte array which will be used to
//		read data from the internal io.Reader object
//		encapsulated by the FileIoReader instance passed
//		as input parameter 'fIoReader'.
//
//		If the value of 'defaultReaderByteArraySize' is
//		less than '16', it will be reset to a size of
//		'4096'.
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
	defaultReaderByteArraySize int,
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
			defaultReaderByteArraySize,
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
//	(1)	This method will delete, overwrite and reset all
//		pre-existing data values in the instance of
//		FileIoReader passed as input parameter
//		'fIoReader'.
//
//	(2) The instance of FileIoReader passed as input
//		parameter 'fIoReader' will retain a pointer
//		reference to input parameter 'reader'. Be sure to
//		close and release this pointer when 'fIoReader'
//		is no longer needed. To perform the 'close'
//		operation, call the local method:
//
//			FileIoReader.Close()
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
//		configure the internal io.Reader contained in
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
//	defaultReaderByteArraySize		int
//
//		The size of the byte array which will be used to
//		read data from the internal io.Reader object
//		encapsulated by the FileIoReader instance passed
//		as input parameter 'fIoReader'.
//
//		If the value of 'defaultReaderByteArraySize' is
//		less than '16', it will be reset to a size of
//		'4096'.
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
	defaultReaderByteArraySize int,
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

	if reader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: The io.Reader instance passed\n"+
			"as input parameter '%v' is invalid!\n"+
			"'%v' is a 'nil' pointer.\n",
			ePrefix.String(),
			readerLabel,
			readerLabel)

		return err
	}

	// Close Old FileIoReader
	var fIoReaderMolecule = new(fileIoReaderMolecule)

	err = fIoReaderMolecule.closeAndRelease(
		fIoReader,
		fIoReaderLabel,
		true, // releaseMemoryResources
		ePrefix.XCpy(fIoReaderLabel))

	if err != nil {
		return err
	}

	// Set Default Byte Array Size
	fIoReader.defaultByteArraySize =
		defaultReaderByteArraySize

	fIoReaderMolecule.
		validateDefaultReaderArraySize(fIoReader)

	// Set io.Reader
	fIoReader.ioReader = &reader

	// Set File Data

	var ok bool

	var fileMgr *FileMgr

	fileMgr, ok = reader.(*FileMgr)

	if ok == true {

		fIoReader.targetReadFileName =
			fileMgr.GetAbsolutePathFileName()

	}

	fIoReader.filePtr, ok = reader.(*os.File)

	if ok == true &&
		len(fIoReader.targetReadFileName) == 0 {

		fIoReader.targetReadFileName =
			fIoReader.filePtr.Name()

	}

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
//	defaultReaderByteArraySize		int
//
//		The size of the byte array which will be used to
//		read data from the internal io.Reader object
//		encapsulated by the FileIoReader instance passed
//		as input paramter 'fIoReader'.
//
//		If the value of 'defaultReaderByteArraySize' is
//		less than '16', it will be reset to a size of
//		'4096'.
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
	defaultReaderByteArraySize int,
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

	var fIoReaderMolecule = new(fileIoReaderMolecule)

	// Close old FileIoReader

	err = fIoReaderMolecule.closeAndRelease(
		fIoReader,
		fIoReaderLabel,
		true, // releaseMemoryResources
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

	fIoReader.defaultByteArraySize =
		defaultReaderByteArraySize

	fIoReaderMolecule.
		validateDefaultReaderArraySize(
			fIoReader)

	return fInfoPlus, err
}

type fileIoReaderMolecule struct {
	lock *sync.Mutex
}

// validateDefaultReaderArraySize
//
// Validates the default buffer size for the instance of
// FileIoReader passed as input parameter 'fIoReader'.
//
// If the buffer size is less than 16-bytes, it will be
// reset to 4096 bytes.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fIoReader					*FileIoReader
//
//		A pointer to an instance of FileIoReader.
//
//		The default buffer size for this FileIoReader
//		instance will be validated. If an invalid value
//		is detected that value will be automatically
//		reset to a value of 4096-bytes.
//
//		This internal member variable is styled as:
//
//			FileIoReader.defaultByteArraySize
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	-- NONE --
func (fIoReaderMolecule *fileIoReaderMolecule) validateDefaultReaderArraySize(
	fIoReader *FileIoReader) {

	if fIoReaderMolecule.lock == nil {
		fIoReaderMolecule.lock = new(sync.Mutex)
	}

	fIoReaderMolecule.lock.Lock()

	defer fIoReaderMolecule.lock.Unlock()

	if fIoReader == nil {
		return
	}

	if fIoReader.defaultByteArraySize < 16 {

		fIoReader.defaultByteArraySize = 4096
	}

	return
}

// closeAndRelease
//
// This method will perform the 'close' procedure on the
// internal io.Reader object encapsulated by the instance
// of FileIoReader passed as input parameter 'fIoReader'.
//
// In addition, this method will also release the
// internal memory resources encapsulated by 'fIoReader'.
//
// This method, fileIoReaderMolecule.closeAndRelease(),
// is the recommended method for 'closing' the internal
// io.Reader object.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	(1)	This method will close the underlying io.Reader
//		object encapsulated by 'fIoReader'.
//
//	(2)	In addition, this method will release the
//		internal memory	resources for the 'fIoReader'
//		object.
//
//	(3)	After completion of this method, the FileIoReader
//		instance passed by 'fIoReader' will be unusable,
//		invalid and unavailable	for future 'read'
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
//	fIoReader					*FileIoReader
//
//		A pointer to an instance of FileIoReader.
//
//		This method will close the underlying io.Reader
//		object contained in 'fIoReader' and release
//		all internal memory resources.
//
//		Upon completion, 'fIoReader' will be rendered
//		invalid, and unusable for all future 'read'
//		operations.
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
//	releaseMemoryResources		bool
//
//		If this parameter is set to 'true', this method
//		will release all internal memory resources for
//		the passed instance of FileIoReader ('fIoReader').
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
func (fIoReaderMolecule *fileIoReaderMolecule) closeAndRelease(
	fIoReader *FileIoReader,
	fIoReaderLabel string,
	releaseMemoryResources bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fIoReaderMolecule.lock == nil {
		fIoReaderMolecule.lock = new(sync.Mutex)
	}

	fIoReaderMolecule.lock.Lock()

	defer fIoReaderMolecule.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileIoReaderAtom." +
		"closeAndRelease()"

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

	var fIoReaderAtom = new(fileIoReaderAtom)

	err = fIoReaderAtom.close(
		fIoReader,
		fIoReaderLabel,
		ePrefix)

	if err != nil {

		return err
	}

	if releaseMemoryResources == true {

		fIoReaderAtom.empty(fIoReader)

	}

	return err
}

type fileIoReaderAtom struct {
	lock *sync.Mutex
}

// close
//
// This method will effectively close the underlying
// io.Reader object contained in the FileIoReader
// instance passed as input parameter 'fIoReader'.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	This method will NOT release the internal memory
//	resources for the 'fIoReader' instance. To release
//	the internal memory resources, call this method:
//
//			fileIoReaderAtom.empty()
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	After completion of this method th3 FileIoReader
//	instance passed as 'fIoReader' will be unusable,
//	invalid and unavailable for future 'read' operations.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fIoReader					*FileIoReader
//
//		A pointer to an instance of FileIoReader.
//
//		The underlying io.Reader object configured for
//		'fIoReader' will be closed.
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
func (fIoReaderAtom *fileIoReaderAtom) close(
	fIoReader *FileIoReader,
	fIoReaderLabel string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fIoReaderAtom.lock == nil {
		fIoReaderAtom.lock = new(sync.Mutex)
	}

	fIoReaderAtom.lock.Lock()

	defer fIoReaderAtom.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileIoReaderAtom." +
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

	if fIoReader.ioReader == nil {

		return err
	}

	var ok bool
	var closerObj io.Closer
	var localReader io.Reader

	localReader = *fIoReader.ioReader

	closerObj, ok = localReader.(io.Closer)

	if ok {

		var err2 error

		err2 = closerObj.Close()

		if err2 != nil {

			errText := fmt.Sprintf(
				"%v\n"+
					"Error returned while closing the 'FileIoReader'\n"+
					"internal io.Writer object.\n",
				ePrefix.String())

			if len(fIoReader.targetReadFileName) > 0 {

				errText += fmt.Sprintf(
					"Target Read File Name: %v\n",
					fIoReader.targetReadFileName)

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
// contained in the FileIoReader instance passed as input
// parameter 'fIoReader'.
//
// Specifically, the following internal member variables
// are set to 'nil' or their initial zero values:
//
//	fIoReader.targetReadFileName = ""
//	fIoReader.filePtr = nil
//	fIoReader.ioReader = nil
//	fIoReader.defaultByteArraySize = 0
//
// After calling this method, the current instance of
// FileIoReader will become invalid and unavailable
// for future 'read' operations.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	This method will NOT 'close' the underlying io.Reader
//	object encapsulated within 'fIoReader'. Best
//	practices call for 'closing' the underlying io.Reader
//	object first and then releasing the internal memory
//	resources with a call to this method,
//	fileIoReaderAtom.empty().
//
//	To 'close' the underlying io.Reader object, first
//	call method, fileIoReaderAtom.close().
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
func (fIoReaderAtom *fileIoReaderAtom) empty(
	fIoReader *FileIoReader) {

	if fIoReaderAtom.lock == nil {
		fIoReaderAtom.lock = new(sync.Mutex)
	}

	fIoReaderAtom.lock.Lock()

	defer fIoReaderAtom.lock.Unlock()

	if fIoReader == nil {

		return
	}

	fIoReader.targetReadFileName = ""

	fIoReader.filePtr = nil

	fIoReader.ioReader = nil

	fIoReader.defaultByteArraySize = 0

	return
}

// isFileIoReaderValid
//
// This method receives a pointer to an instance of
// FileIoReader ('fIoReader') which will be analyzed to
// determine if all the member variables contain valid
// data values.
//
// If input parameter 'fIoReader' is determined to be
// invalid, this method returns an error.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fIoReader				*FileIoReader
//
//		A pointer to an instance of FileIoReader.
//
//		If any of the internal member variable data
//		values encapsulated in 'fIoReader' are
//		determined to be invalid, this method will return
//		an error.
//
//	fIoReaderLabel			string
//
//		The name or label associated with input parameter
//		'fIoReader' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fIoReader" will
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
//		contained in the instance of FileIoReader
//		passed as 'fIoReader' are found to be
//		invalid, this method will return an error
//		configured with an appropriate message
//		identifying the invalid	member data variable.
//
//		If all internal member data variables evaluate
//		as valid, this returned error value will be set
//		to 'nil'.
//
//		If errors are encountered during processing or if
//		any 'fIoReader' internal member data values
//		are found to be invalid, the returned error Type
//		will encapsulate an appropriate error message.
//	 	This returned error message will incorporate the
//	 	method chain and text passed by input parameter,
//	 	'errorPrefix'. The 'errorPrefix' text will be
//	 	prefixed to the beginning of the error message.
func (fIoReaderAtom *fileIoReaderAtom) isFileIoReaderValid(
	fIoReader *FileIoReader,
	fIoReaderLabel string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fIoReaderAtom.lock == nil {
		fIoReaderAtom.lock = new(sync.Mutex)
	}

	fIoReaderAtom.lock.Lock()

	defer fIoReaderAtom.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fileIoReaderAtom."+
			"isFileIoReaderValid()",
		"")

	if err != nil {
		return err
	}

	if len(fIoReaderLabel) == 0 {

		fIoReaderLabel = "fIoReader"
	}

	if fIoReader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is a nil pointer!\n"+
			"%v is invalid.\n",
			ePrefix.String(),
			fIoReaderLabel,
			fIoReaderLabel)

		return err
	}

	if fIoReader.ioReader == nil {

		err = fmt.Errorf("%v\n"+
			" -------------------------------------------------------------------\n"+
			" ERROR: The %v instance of FileIoReader is invalid!\n"+
			" The internal io.Reader object was never initialized.\n"+
			" Call one of the 'New' methods or 'Setter' methods to create\n"+
			" a valid instance of FileIoReader.\n",
			ePrefix.String(),
			fIoReaderLabel)
	}

	return err
}
