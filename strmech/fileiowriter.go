package strmech

import (
	"errors"
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"math/big"
	"os"
	"strings"
	"sync"
)

type FileIoWriter struct {
	ioWriter            *io.Writer
	filePtr             *os.File
	targetWriteFileName string

	lock *sync.Mutex
}

// Close
//
// This method is used to close any open file pointers
// and perform required clean-up operations.
//
// Users MUST call this method after all 'write'
// operations have been completed.
//
// After calling this method, the current instance of
// FileIoWriter will be invalid and unavailable for
// future 'write' operations.
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

	err = new(fileIoWriterMolecule).close(
		fIoWriter,
		"fIoWriter",
		ePrefix.XCpy("fIoWriter"))

	return err
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
//	(2)	Input parameter 'writer' will accept a pointer to
//		an instance of os.File because os.File implements
//		the io.Writer interface.
//
//	(3) The returned instance of FileIoWriter will retain
//		a pointer reference to input parameter 'writer'.
//		Be sure to close and release this pointer when the
//		returned instance of FileIoWriter is no longer
//		needed. To perform the 'close' operation, call the
//		local method:
//
//			FileIoWriter.Close()
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
//		file and performing any other required clean-up
//		operations in addition to calling local method
//		FileIoWriter.Close().
//
//		While the returned instance of FileIoWriter
//		is primarily designed for writing data to disk
//		files, this 'writer' will in fact write data to
//		any object implementing the io.Writer interface.
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
	errorPrefix interface{}) (
	FileIoWriter,
	error) {

	if fIoWriter.lock == nil {
		fIoWriter.lock = new(sync.Mutex)
	}

	fIoWriter.lock.Lock()

	defer fIoWriter.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error
	var newFileIoWriter FileIoWriter

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
			&newFileIoWriter,
			"newFileIoWriter",
			writer,
			"writer",
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
//	newFileIoWriter			FileIoWriter
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
func (fIoWriter *FileIoWriter) NewFileMgr(
	fileMgr *FileMgr,
	openFileReadWrite bool,
	errorPrefix interface{}) (
	fInfoPlus FileInfoPlus,
	newFileIoWriter FileIoWriter,
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
			"NewFileMgr()",
		"")

	if err != nil {

		return fInfoPlus, newFileIoWriter, err
	}

	fInfoPlus,
		err = new(fileIoWriterMicrobot).
		setFileMgr(
			&newFileIoWriter,
			"newFileIoWriter",
			fileMgr,
			"fileMgr",
			openFileReadWrite,
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
//	newFileIoWriter			FileIoWriter
//
//		If this method completes successfully, a fully
//		configured instance of FileIoWriter will
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
func (fIoWriter *FileIoWriter) NewPathFileName(
	pathFileName string,
	openFileReadWrite bool,
	errorPrefix interface{}) (
	fInfoPlus FileInfoPlus,
	newFileIoWriter FileIoWriter,
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
			"NewPathFileName()",
		"")

	if err != nil {

		return fInfoPlus, newFileIoWriter, err
	}

	fInfoPlus,
		err = new(fileIoWriterNanobot).
		setPathFileName(
			&newFileIoWriter,
			"newFileIoWriter",
			pathFileName,
			"pathFileName",
			openFileReadWrite,
			ePrefix.XCpy(
				pathFileName))

	return fInfoPlus, newFileIoWriter, err
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
//		performing any other required clean-up operations
//		in addition to calling FileIoWriter.Close().
//
//		While the returned instance of FileIoWriter
//		is primarily designed for writing data from disk
//		files, this 'writer' will in fact write data from
//		any object implementing the io.Writer interface.
//
//		Remember that the user is responsible for
//		'closing' this io.Writer object. The FileIoWriter
//		method 'Close()' will NOT close this io.Writer
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
func (fIoWriter *FileIoWriter) SetIoWriter(
	writer io.Writer,
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
// reconfigure the internal bufio.Writer encapsulated in
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
			ePrefix.XCpy(
				"fileMgr"))

	return fInfoPlus, err
}

// SetPathFileName
//
// This method will completely re-initialize the current
// instance of FileIoWriter using the path and file
// name passed as input parameter 'pathFileName'.
//
// The path and file name specified by 'pathFileName'
// will be used to reconfigure the internal bufio.Writer
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
//	(2)	After all 'read' and 'write' operations have been
//		completed, the user MUST call the 'Close' method
//		to perform necessary clean-up operations:
//
//			FileIoWriter.Close()
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

	if fIoWriter.ioWriter == nil {

		err = fmt.Errorf("%v\n"+
			"-------------------------------------------------------\n"+
			"Error: This instance of 'FileIoWriter' is invalid!\n"+
			"The internal io.Writer has NOT been initialized.\n"+
			"Call one of the 'New' or 'Setter' methods when creating\n"+
			"a new valid instance of 'FileIoWriter'\n",
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

	var i64NumOfBytesWritten int64

	i64NumOfBytesWritten,
		err = new(fileIoWriterAtom).
		writeBytes(
			fIoWriter,
			"fIoWriter",
			bytesToWrite,
			"bytesToWrite",
			"", // endOfLineTerminator
			ePrefix)

	numBytesWritten = int(i64NumOfBytesWritten)

	return numBytesWritten, err
}

// WriteRunes
//
// Receives a rune array in the form of a RuneArrayDto
// instance and proceeds to write all the runes contained
// therein to the underlying io.Writer encapsulated by
// the current instance of FileIoWriter.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	runeArray					*RuneArrayDto
//
//		A pointer to an instance of RuneArrayDto. The
//		rune array contained in this RuneArrayDto
//		instance will be written to the io.Writer
//		encapsulated by the current instance of
//		FileIoWriter.
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
//	numOfBytesWritten			int
//
//		Returns the number of bytes extracted from
//		input parameter 'runeArray' and written to the
//		internal io.Writer object encapsulated by the
//		current instance of FileIoWriter.
//
//	autoCloseOnExit				bool
//
//		When this parameter is set to 'true', this
//		method will automatically perform the following
//		clean-up tasks upon exit, even if processing
//		errors are encountered:
//
//		(1)	The internal io.Writer object will be
//			properly closed and there will be no need
//			to make a separate call to local method,
//			FileIoWriter.Close().
//
//		(2) After performing this clean-up operation, the
//			current instance of FileIoWriter will invalid
//			and unusable for future 'write' operations.
//
//		If input parameter 'autoCloseOnExit' is
//		set to 'false', this method will NOT close
//		the internal io.Writer object. This means the user
//	 	is then responsible for performing the 'Close'
//		operation by calling the local method:
//
//			FileIoWriter.Close()
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
//		Returns the number of bytes extracted from
//		input parameter 'strBuilder' and written to the
//		internal io.Writer object encapsulated by the
//		current instance of FileIoWriter.
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
func (fIoWriter *FileIoWriter) WriteRunes(
	runeArray *RuneArrayDto,
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
			"WriteRunes()",
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

	if runeArray == nil {

		err = fmt.Errorf("%v\n"+
			"-------------------------------------------------------\n"+
			"Error: Input parameter 'runeArray' is invalid!\n"+
			"'runeArray' is a 'nil' pointer.\n",
			ePrefix.String())

		return numOfBytesWritten, err
	}

	if runeArray.GetRuneArrayLength() == 0 {

		err = fmt.Errorf("%v\n"+
			"-------------------------------------------------------\n"+
			"Error: Input parameter 'runeArray' is invalid!\n"+
			"'runeArray' contains an empty or zero length rune array.\n",
			ePrefix.String())

		return numOfBytesWritten, err
	}

	var err1, err2 error
	var localNumBytesWritten int
	var writer = *fIoWriter.ioWriter

	localNumBytesWritten,
		err2 = writer.Write(
		[]byte(runeArray.String()))

	if err2 != nil {

		err1 = fmt.Errorf("%v\n"+
			"Error returned by writer.Write((runeArray).\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			err2.Error())

	} else {

		numOfBytesWritten = int64(localNumBytesWritten)
	}

	err = new(fileIoWriterMicrobot).
		performAutoClose(
			fIoWriter,
			"fIoWriter",
			autoCloseOnExit,
			err1,
			ePrefix)

	return numOfBytesWritten, err
}

// WriteStrBuilder
//
// Receives an instance of strings.Builder and proceeds
// to write all the string data contained therein to the
// previously configured internal io.writer
// destination configured for the current instance of
// FileIoWriter.
//
// If input parameter 'autoCloseOnExit' is set to
// 'true', this method will automatically perform all
// required clean-up tasks upon completion. Clean-up
// tasks involve flushing the io.Writer object,
// closing the io.Writer object and then deleting
// the io.Writer structure values internal to
// the current FileIoWriter instance. When these
// clean-up tasks are completed, the current
// FileIoWriter instance will be invalid and
// unavailable for future 'write' operations.
//
// If input parameter 'autoCloseOnExit' is set to
// 'false', this method will automatically flush the
// 'write' buffer. This means that all data remaining in
// the 'write' buffer will be written to the underlying
// io.Writer output destination. However, most
// importantly, the user is then responsible for
// performing the 'Close' operation by calling the local
// method:
//
//	FileIoWriter.Close()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	strBuilder					*strings.Builder
//
//		A pointer to an instance of strings.Builder. The
//		entire string contents of 'strBuilder' will be
//		written to the internal io.Writer destination
//		encapsulated by the current instance of
//		FileIoWriter.
//
//		If the string contained in 'strBuilder' is empty,
//		or a zero length string, an error will be
//		returned.
//
//	autoCloseOnExit		bool
//
//		When this parameter is set to 'true', this
//		method will automatically perform the following
//		clean-up tasks upon exit, even if processing
//		errors are encountered:
//
//		(1)	The internal io.Writer object will be
//			properly closed and there will be no need
//			to make a separate call to local method,
//			FileIoWriter.Close().
//
//		(2) After performing this clean-up operation, the
//			current instance of FileIoWriter will invalid
//			and unusable for future 'write' operations.
//
//		If input parameter 'autoCloseOnExit' is
//		set to 'false', this method will NOT close
//		the internal io.Writer object. This means the user
//	 	is then responsible for performing the 'Close'
//		operation by calling the local method:
//
//			FileIoWriter.Close()
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
//		Returns the number of bytes extracted from
//		input parameter 'strBuilder' and written to the
//		internal io.Writer object encapsulated by the
//		current instance of FileIoWriter.
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
func (fIoWriter *FileIoWriter) WriteStrBuilder(
	strBuilder *strings.Builder,
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
			"WriteStrBuilder()",
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

	if strBuilder == nil {

		err = fmt.Errorf("%v\n"+
			"-------------------------------------------------------\n"+
			"Error: Input parameter 'strBuilder' is invalid!\n"+
			"'strBuilder' is a 'nil' pointer.\n",
			ePrefix.String())

		return numOfBytesWritten, err
	}

	if strBuilder.Len() == 0 {

		err = fmt.Errorf("%v\n"+
			"-------------------------------------------------------\n"+
			"Error: Input parameter 'strBuilder' is invalid!\n"+
			"'strBuilder' contains an empty or zero length string.\n",
			ePrefix.String())

		return numOfBytesWritten, err
	}

	var err1, err2 error
	var localNumBytesWritten int
	var writer = *fIoWriter.ioWriter

	localNumBytesWritten,
		err2 = writer.Write(
		[]byte(strBuilder.String()))

	if err2 != nil {

		err1 = fmt.Errorf("%v\n"+
			"Error returned by fIoWriter.writer.Write(strBuilder).\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			err2.Error())

	} else {

		numOfBytesWritten = int64(localNumBytesWritten)
	}

	err = new(fileIoWriterMicrobot).
		performAutoClose(
			fIoWriter,
			"fIoWriter",
			autoCloseOnExit,
			err1,
			ePrefix)

	return numOfBytesWritten, err
}

type fileIoWriterMechanics struct {
	lock *sync.Mutex
}

// writeCharacters
//
// This method will accept the following types:
//
//	string
//	*string
//	[]string
//	rune
//	[]rune
//	float32
//	*float32
//	float64
//	*float64
//	*BigFloatDto
//	BigFloatDto
//	*big.Float
//	big.Float
//	big.Rat
//	*big.Rat
//	int8
//	*int8
//	int16
//	*int16
//	int
//	*int
//	int32
//	*int32
//	int64
//	*int64
//	uint8
//	*uint8
//	uint16
//	*uint16
//	uint
//	*uint
//	uint32
//	*uint32
//	uint64,
//	*uint64
//	big.Int
//	*big.Int
//	TextFieldFormatDtoFloat64
//	*TextFieldFormatDtoFloat64
//	TextFieldFormatDtoBigFloat
//	*TextFieldFormatDtoBigFloat
//	NumberStrKernel
//	*NumberStrKernel
func (fIoWriterMech *fileIoWriterMechanics) writeCharacters(
	fIoWriter *FileIoWriter,
	fIoWriterLabel string,
	charsToWrite interface{},
	charsToWriteLabel string,
	endOfLineTerminator string,
	autoCloseOnExit bool,
	errPrefDto *ePref.ErrPrefixDto) (
	numOfBytesWritten int64,
	err error) {

	if fIoWriterMech.lock == nil {
		fIoWriterMech.lock = new(sync.Mutex)
	}

	fIoWriterMech.lock.Lock()

	defer fIoWriterMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileIoWriterMechanics." +
		"writeSomething()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		return numOfBytesWritten, err
	}

	if len(fIoWriterLabel) == 0 {

		fIoWriterLabel = "fIoWriter"
	}

	if len(charsToWriteLabel) == 0 {

		charsToWriteLabel = "charsToWrite"
	}

	if fIoWriter == nil {

		err = fmt.Errorf("%v\n"+
			"Error: The FileIoWriter instance passed\n"+
			"as input parameter '%v' is invalid!\n"+
			"'%v' is a 'nil' pointer.\n",
			ePrefix.String(),
			fIoWriterLabel,
			fIoWriterLabel)

		return numOfBytesWritten, err
	}

	if fIoWriter.ioWriter == nil {

		err = fmt.Errorf("%v\n"+
			"------------------------------------------------------------------------------\n"+
			"Error: Input parameter '%v' is invalid!\n"+
			"The internal io.Writer for this FileIoWriter instance has NOT been initialized.\n"+
			"Call one of the 'New' or 'Setter' methods when creating a new valid instance of\n"+
			"'FileIoWriter'\n",
			ePrefix.String(),
			fIoWriterLabel)

		return numOfBytesWritten, err
	}

	if charsToWrite == nil {

		err = fmt.Errorf("%v\n"+
			"-------------------------------------------------------\n"+
			"Error: Input parameter '%v' is invalid!\n"+
			"'%v' is a 'nil' value.\n",
			ePrefix.String(),
			charsToWriteLabel,
			charsToWriteLabel)

		return numOfBytesWritten, err
	}

	var stringToWrite string
	var ok bool
	var fIoWriterAtom = new(fileIoWriterAtom)

	switch charsToWrite.(type) {

	case string:
		// string

		stringToWrite, ok = charsToWrite.(string)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"Input parameter '%v' is ERROR!\n"+
				"'%v' was identified as a string.\n"+
				"string cast from '%v' to 'stringToWrite' Failed.\n",
				ePrefix.String(),
				charsToWriteLabel,
				charsToWriteLabel,
				charsToWriteLabel)

			return numOfBytesWritten, err
		}

		numOfBytesWritten,
			err = fIoWriterAtom.writeBytes(
			fIoWriter,
			"fIoWriter",
			[]byte(stringToWrite),
			"byteArray",
			endOfLineTerminator,
			ePrefix.XCpy("byteArray<-stringToWrite"))

	case *string:
		// string pointer

		var strPtr *string

		strPtr, ok = charsToWrite.(*string)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"Input parameter '%v' is ERROR!\n"+
				"'%v' was identified as a string pointer.\n"+
				"string cast from '%v' to string pointer Failed.\n",
				ePrefix.String(),
				charsToWriteLabel,
				charsToWriteLabel,
				charsToWriteLabel)

			return numOfBytesWritten, err
		}

		numOfBytesWritten,
			err = fIoWriterAtom.writeBytes(
			fIoWriter,
			"fIoWriter",
			[]byte(*strPtr),
			"byteArray",
			endOfLineTerminator,
			ePrefix.XCpy("byteArray<-*string"))

	case []string:
		// string array

		var strArray []string

		strArray, ok = charsToWrite.([]string)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"ERROR: Input parameter '%v' is invalid!\n"+
				"'%v' was identified as a string array.\n"+
				"string array cast from '%v' to string array Failed.\n",
				ePrefix.String(),
				charsToWriteLabel,
				charsToWriteLabel,
				charsToWriteLabel)

			return numOfBytesWritten, err
		}

		numOfBytesWritten,
			err = fIoWriterAtom.
			writeStringArray(
				fIoWriter,
				"fIoWriter",
				strArray,
				"strArray",
				endOfLineTerminator,
				ePrefix.XCpy("strArray<-[]string"))

	case StringArrayDto:

		var strArrayDto StringArrayDto

		strArrayDto, ok = charsToWrite.(StringArrayDto)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"Input parameter '%v' is ERROR!\n"+
				"'%v' was identified as a StringArrayDto.\n"+
				"The cast from '%v' to StringArrayDto Failed.\n",
				ePrefix.String(),
				charsToWriteLabel,
				charsToWriteLabel,
				charsToWriteLabel)

			return numOfBytesWritten, err
		}

		numOfBytesWritten,
			err = fIoWriterAtom.
			writeStringArray(
				fIoWriter,
				"fIoWriter",
				strArrayDto.StrArray,
				"strArrayDto.StrArray",
				endOfLineTerminator,
				ePrefix.XCpy("strArray<-StringArrayDto"))

	case *StringArrayDto:

		var strArrayDtoPtr *StringArrayDto

		strArrayDtoPtr, ok = charsToWrite.(*StringArrayDto)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"----------------------------------------------------------------\n"+
				"ERROR: Input parameter '%v' is invalid!\n"+
				"'%v' was identified as a StringArrayDto Pointer.\n"+
				"The cast from '%v' to *StringArrayDto Failed.\n",
				ePrefix.String(),
				charsToWriteLabel,
				charsToWriteLabel,
				charsToWriteLabel)

			return numOfBytesWritten, err
		}

		numOfBytesWritten,
			err = fIoWriterAtom.
			writeStringArray(
				fIoWriter,
				"fIoWriter",
				strArrayDtoPtr.StrArray,
				"strArrayDtoPtr.StrArray",
				endOfLineTerminator,
				ePrefix.XCpy("strArray<-*StringArrayDto"))

	case rune:

		var runeToWrite rune

		runeToWrite, ok = charsToWrite.(rune)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"Input parameter '%v' is ERROR!\n"+
				"'%v' was identified as a rune.\n"+
				"rune cast from '%v' to rune Failed.\n",
				ePrefix.String(),
				charsToWriteLabel,
				charsToWriteLabel,
				charsToWriteLabel)

			return numOfBytesWritten, err
		}

		numOfBytesWritten,
			err = fIoWriterAtom.writeBytes(
			fIoWriter,
			"fIoWriter",
			[]byte(string(runeToWrite)),
			"runeToWrite",
			endOfLineTerminator,
			ePrefix.XCpy("byteArray<-runeToWrite"))

	case []rune:

		var runeArray []rune

		runeArray, ok = charsToWrite.([]rune)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"Input parameter '%v' is ERROR!\n"+
				"'%v' was identified as a rune array ([]rune).\n"+
				"rune array cast from '%v' to 'runeArray' Failed.\n",
				ePrefix.String(),
				charsToWriteLabel,
				charsToWriteLabel,
				charsToWriteLabel)

			return numOfBytesWritten, err
		}

		numOfBytesWritten,
			err = fIoWriterAtom.writeBytes(
			fIoWriter,
			"fIoWriter",
			[]byte(string(runeArray)),
			"runeArray",
			endOfLineTerminator,
			ePrefix.XCpy("byteArray<-runeArray"))

	case float32, *float32, float64, *float64, *BigFloatDto,
		BigFloatDto, *big.Float, big.Float, big.Rat, *big.Rat,
		int8, *int8, int16, *int16, int, *int, //int32,
		*int32, int64, *int64, uint8, *uint8, uint16,
		*uint16, uint, *uint, uint32, *uint32, uint64,
		*uint64, big.Int, *big.Int, TextFieldFormatDtoFloat64,
		*TextFieldFormatDtoFloat64, TextFieldFormatDtoBigFloat,
		*TextFieldFormatDtoBigFloat, NumberStrKernel,
		*NumberStrKernel:

		stringToWrite,
			err = new(mathHelperNanobot).
			numericValueToNativeNumStr(
				charsToWrite,
				ePrefix.XCpy("<-charsToWrite"))

		if err != nil {

			return numOfBytesWritten,
				fmt.Errorf("%v\n"+
					"Error converting numeric value to a number string!\n"+
					"Error=\n%v\n",
					funcName,
					err.Error())

		}

		numOfBytesWritten,
			err = fIoWriterAtom.writeBytes(
			fIoWriter,
			"fIoWriter",
			[]byte(stringToWrite),
			"Number String",
			endOfLineTerminator,
			ePrefix.XCpy("byteArray<-number string"))

	default:

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'charsToWrite' is an invalid type!\n"+
			"'charsToWrite' is unsupported type '%T'\n",
			ePrefix.String(),
			charsToWrite)

		return numOfBytesWritten, err
	}

	if err != nil &&
		autoCloseOnExit == true {

		err = new(fileIoWriterMolecule).
			close(
				fIoWriter,
				"fIoWriter",
				ePrefix)

	}

	return numOfBytesWritten, err
}

type fileIoWriterMicrobot struct {
	lock *sync.Mutex
}

// performAutoClose
//
// Receives an instance of FileIoWriter and performs
// clean-up operations as specified by input parameter
// 'autoCloseOnExit'.
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
//	autoCloseOnExit				bool
//
//		When this parameter is set to 'true', this
//		method will automatically perform the following
//		clean-up tasks upon exit, even if processing
//		errors are encountered:
//
//		(1)	The internal io.Writer object will be
//			properly closed and there will be no need
//			to make a separate call to local method,
//			FileIoWriter.Close().
//
//		(2) After performing this clean-up operation, the
//			current instance of FileIoWriter will invalid
//			and unusable for future 'write' operations.
//
//		If input parameter 'autoCloseOnExit' is
//		set to 'false', this method will NOT close
//		the internal io.Writer object. This means the user
//	 	is then responsible for performing the 'Close'
//		operation by calling the local method:
//
//			FileIoWriter.Close()
//
//	accumulatedErrors			error
//
//		This parameter includes any errors accumulated
//		to this point by the calling method. Any errors
//		encountered while performing the clean-up
//		operation on 'fIoWriter' will be added or
//		'joined' to 'accumulatedErrors' and returned as
//		a single error by this method.
//
//		It is assumed that this method will be called by
//		high level methods which may have already
//		encountered and recorded processing errors. If
//		this is the case, this parameter will contain one
//		or more processing errors. If this method produces
//		any additional errors, they will be joined to
//		'accumulatedErrors' and all errors will be
//		returned to the caller.
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
//	err							error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//		Remember that any errors encountered by this
//		method will be joined to existing errors passed
//		as input parameter 'accumulatedErrors'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errPrefDto'.
//	 	The 'errPrefDto' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fIoWriterMicrobot *fileIoWriterMicrobot) performAutoClose(
	fIoWriter *FileIoWriter,
	fIoWriterLabel string,
	autoCloseOnExit bool,
	accumulatedErrors error,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if fIoWriterMicrobot.lock == nil {
		fIoWriterMicrobot.lock = new(sync.Mutex)
	}

	fIoWriterMicrobot.lock.Lock()

	defer fIoWriterMicrobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err2 error

	funcName := "fileIoWriterMicrobot." +
		"performAutoClose()"

	err = errors.Join(err, accumulatedErrors)

	ePrefix,
		err2 = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err2 != nil {

		err = errors.Join(err, err2)

		return err
	}

	if len(fIoWriterLabel) == 0 {

		fIoWriterLabel = "fIoWriter"
	}

	if fIoWriter == nil {

		err2 = fmt.Errorf("%v\n"+
			"Error: The FileIoWriter instance passed\n"+
			"as input parameter '%v' is invalid!\n"+
			"'%v' is a 'nil' pointer.\n",
			ePrefix.String(),
			fIoWriterLabel,
			fIoWriterLabel)

		err = errors.Join(err, err2)

		return err
	}

	if fIoWriter.ioWriter == nil {

		err2 = fmt.Errorf("%v\n"+
			"-------------------------------------------------------\n"+
			"Error: This instance of 'FileIoWriter' is invalid!\n"+
			"The internal io.Writer has NOT been initialized.\n"+
			"Call one of the 'New' or 'Setter' methods when creating\n"+
			"a new valid instance of 'FileIoWriter'\n",
			ePrefix.String())

		err = errors.Join(err, err2)

		return err
	}

	var cleanUpStatus string

	if accumulatedErrors != nil {

		cleanUpStatus = "Processing Error"

	} else {

		cleanUpStatus = "Successful Completion"

	}

	if autoCloseOnExit == true {

		err2 = new(fileIoWriterMolecule).close(
			fIoWriter,
			fIoWriterLabel,
			ePrefix.XCpy(fmt.Sprintf(
				"%v Close-Writer",
				cleanUpStatus)))

		if err2 != nil {

			err = errors.Join(err, err2)
		}

	}

	return err
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
			ePrefix.XCpy(fIoWriterLabel+"<-"+fileMgrLabel))

	return fInfoPlus, err
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
//		file and performing any other required clean-up
//		operations in addition to calling local method
//		FileIoWriter.FlushAndClose().
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
	_ = fIoWriterMolecule.close(
		fIoWriter,
		"",
		nil)

	fIoWriter.ioWriter = &writer

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

	err = new(fileIoWriterMolecule).close(
		fIoWriter,
		fIoWriterLabel,
		ePrefix.XCpy(fIoWriterLabel))

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

		return fInfoPlus, err
	}

	fIoWriter.targetWriteFileName = pathFileName

	var writer io.Writer

	writer = fIoWriter.filePtr

	fIoWriter.ioWriter = &writer

	return fInfoPlus, err
}

type fileIoWriterMolecule struct {
	lock *sync.Mutex
}

// close
//
// This method is designed to perform clean up
// operations after completion of all 'write'
// operations.
//
// All internal member variable data values for the
// instance of FileIoWriter passed as input parameter
// 'fIoWriter' will be deleted and reset to their zero
// states.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	This method will delete all pre-existing data
//		values in the instance of FileIoWriter passed
//		as input parameter 'fIoWriter'.
//
//		After completion of this method this
//		FileIoWriter instance will be unusable,
//		invalid and unavailable for future 'write'
//		operations.
//
//	(2)	This 'close' method will NOT flush the 'write'
//		buffer. To flush the 'write' buffer call:
//			fileIoWriterMolecule.flush()
//
//		Be sure to call the 'flush()' method before you
//		call this method ('close()').
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
//		this instance will be deleted.
//
//		If a file pointer (*os.File) was previously
//		configured for 'fIoWriter', it will be closed
//		and set to 'nil' by this method.
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
func (fIoWriterMolecule *fileIoWriterMolecule) close(
	fIoWriter *FileIoWriter,
	fIoWriterLabel string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fIoWriterMolecule.lock == nil {
		fIoWriterMolecule.lock = new(sync.Mutex)
	}

	fIoWriterMolecule.lock.Lock()

	defer fIoWriterMolecule.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileIoWriterMolecule." +
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

	if fIoWriter.filePtr != nil {

		var err2 error

		err2 = fIoWriter.filePtr.Close()

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Error returned while closing the target 'target' file!\n"+
				"%v.filePtr.Close()\n"+
				"Target Read File = '%v'\n"+
				"Error = \n%v\n",
				ePrefix.String(),
				fIoWriterLabel,
				fIoWriter.targetWriteFileName,
				err2.Error())

		}
	}

	fIoWriter.targetWriteFileName = ""

	fIoWriter.filePtr = nil

	fIoWriter.ioWriter = nil

	return err
}

type fileIoWriterAtom struct {
	lock *sync.Mutex
}

// writeBytes
//
// Writes a byte array to the io.Writer object
// contained in the FileIoWriter instance passed as input
// parameter 'fIoWriter'
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	If the byte array passed as input parameter
//	'byteArray' is empty or contains zero array elements,
//	this method will take no action, no error will be
//	returned and the returned number of bytes written
//	('numOfBytesWritten') will be set to zero.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fIoWriter					*FileIoWriter
//
//		A pointer to an instance of FileIoWriter.
//
//		The contents of the byte array passed as input
//		parameter '' will be written to the internal
//		io.Writer object encapsulated by this
//		FileIoWriter instance.
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
//	byteArray					[]string
//
//		An array of strings which will be written to
//		the internal io.Writer object encapsulated
//		within the FileIoWriter instance passed as input
//		parameter 'fIoWriter'.
//
//		If parameter 'endOfLineTerminator' has a length
//		greater than zero, 'endOfLineTerminator' will be
//		appended to each string written to the io.Writer
//		object.
//
//		If 'byteArray' is empty or passed as a zero
//		length byte array, the method will take no
//		action, no error will be returned and the
//		returned number of bytes written
//		('numOfBytesWritten') will be set to zero.
//
//	byteArrayLabel				string
//
//		The name or label associated with input parameter
//		'byteArray' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "byteArray" will be
//		automatically applied.
//
//	endOfLineTerminator string
//
//		If this parameter has a string length greater
//		than zero, this string will be appended to
//		each byte array element ('byteArray') written
//		to the io.Writer object contained in 'fIoWriter'.
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
//		The number of bytes written to the io.Writer
//		object encapsulated in the FileIoWriter instance
//		passed as input parameter 'fIoWriter'.
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
func (fIoWriterAtom *fileIoWriterAtom) writeBytes(
	fIoWriter *FileIoWriter,
	fIoWriterLabel string,
	byteArray []byte,
	byteArrayLabel string,
	endOfLineTerminator string,
	errPrefDto *ePref.ErrPrefixDto) (
	numOfBytesWritten int64,
	err error) {

	if fIoWriterAtom.lock == nil {
		fIoWriterAtom.lock = new(sync.Mutex)
	}

	fIoWriterAtom.lock.Lock()

	defer fIoWriterAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileIoWriterAtom." +
		"writeBytes()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		return numOfBytesWritten, err
	}

	if len(fIoWriterLabel) == 0 {

		fIoWriterLabel = "fIoWriter"
	}

	if len(byteArrayLabel) == 0 {

		byteArrayLabel = "byteArray"
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

		return numOfBytesWritten, err
	}

	if fIoWriter.ioWriter == nil {

		err = fmt.Errorf("%v\n"+
			"-------------------------------------------------------\n"+
			"Error: This instance of 'FileIoWriter' passed"+
			"as input parameter '%v' is invalid!\n"+
			"The internal io.Writer has NOT been initialized.\n"+
			"Call one of the 'New' or 'Setter' methods when creating\n"+
			"a new valid instance of 'FileIoWriter'\n",
			fIoWriterLabel,
			ePrefix.String())

		return numOfBytesWritten, err
	}

	lenByteArray := len(byteArray)

	if lenByteArray == 0 {

		return numOfBytesWritten, err
	}

	var writer = *fIoWriter.ioWriter
	var err2 error
	var localNumBytesWritten int

	localNumBytesWritten,
		err2 = writer.Write(
		byteArray)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by writer.Write(byteArray)"+
			"while writing original byte array.\n"+
			"byteArray= '%v'\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			string(byteArray),
			err2.Error())

		return numOfBytesWritten, err

	} else {

		numOfBytesWritten += int64(localNumBytesWritten)
	}

	if len(endOfLineTerminator) > 0 {

		localNumBytesWritten,
			err2 = writer.Write(
			[]byte(endOfLineTerminator))

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Error returned by writer.Write(byteArray)"+
				"while writing original byte array.\n"+
				"byteArray= '%v'\n"+
				"Error=\n%v\n",
				ePrefix.String(),
				string(byteArray),
				err2.Error())

		} else {

			numOfBytesWritten += int64(localNumBytesWritten)
		}
	}

	return numOfBytesWritten, err
}

// writeStringArray
//
// Writes a string array to the io.Writer object
// contained in the FileIoWriter instance passed as input
// parameter 'fIoWriter'
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	If the string array passed as input parameter
//	'strArray' is empty or contains zero array elements,
//	this method will take no action, no error will be
//	returned and the returned number of bytes written
//	'numOfBytesWritten' will be set to zero.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fIoWriter					*FileIoWriter
//
//		A pointer to an instance of FileIoWriter.
//
//		The contents of the string array passed as input
//		parameter '' will be written to the internal
//		io.Writer object encapsulated by this
//		FileIoWriter instance.
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
//	strArray					[]string
//
//		An array of strings which will be written to
//		the internal io.Writer object encapsulated
//		within the FileIoWriter instance passed as input
//		parameter 'fIoWriter'.
//
//		If parameter 'endOfLineTerminator' has a length
//		greater than zero, 'endOfLineTerminator' will be
//		appended to each string written to the io.Writer
//		object.
//
//		If 'strArray' is empty or passed as a zero length
//		byte array, the method will take no action, no
//		error will be returned and the returned number of
//		bytes written ('numOfBytesWritten') will be set
//		to zero.
//
//	strArrayLabel				string
//
//		The name or label associated with input parameter
//		'strArray' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "strArray" will be
//		automatically applied.
//
//	endOfLineTerminator string
//
//		If this parameter has a string length greater
//		than zero, this string will be appended to
//		each string array element ('strArray') written
//		to the io.Writer object contained in 'fIoWriter'.
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
//		The number of bytes written to the io.Writer
//		object encapsulated in the FileIoWriter instance
//		passed as input parameter 'fIoWriter'.
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
func (fIoWriterAtom *fileIoWriterAtom) writeStringArray(
	fIoWriter *FileIoWriter,
	fIoWriterLabel string,
	strArray []string,
	strArrayLabel string,
	endOfLineTerminator string,
	errPrefDto *ePref.ErrPrefixDto) (
	numOfBytesWritten int64,
	err error) {

	if fIoWriterAtom.lock == nil {
		fIoWriterAtom.lock = new(sync.Mutex)
	}

	fIoWriterAtom.lock.Lock()

	defer fIoWriterAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileIoWriterAtom." +
		"writeStringArray()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		return numOfBytesWritten, err
	}

	if len(fIoWriterLabel) == 0 {

		fIoWriterLabel = "fIoWriter"
	}

	if len(strArrayLabel) == 0 {

		strArrayLabel = "strArray"
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

		return numOfBytesWritten, err
	}

	if fIoWriter.ioWriter == nil {

		err = fmt.Errorf("%v\n"+
			"-------------------------------------------------------\n"+
			"Error: This instance of 'FileIoWriter' passed"+
			"as input parameter '%v' is invalid!\n"+
			"The internal io.Writer has NOT been initialized.\n"+
			"Call one of the 'New' or 'Setter' methods when creating\n"+
			"a new valid instance of 'FileIoWriter'\n",
			fIoWriterLabel,
			ePrefix.String())

		return numOfBytesWritten, err
	}

	lenStrArray := len(strArray)
	lenEOLTerminator := len(endOfLineTerminator)

	if lenStrArray == 0 {

		return numOfBytesWritten, err
	}

	var err2 error
	var localNumBytesWritten int
	var writer = *fIoWriter.ioWriter
	var strToWrite string

	for i := 0; i <= lenStrArray; i++ {

		strToWrite = strArray[i]

		if lenEOLTerminator > 0 {
			strToWrite += endOfLineTerminator
		}

		localNumBytesWritten,
			err2 = writer.Write(
			[]byte(strToWrite))

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Error returned by writer.Write([]byte(strToWrite)).\n"+
				"strToWrite= '%v'\n"+
				"Error=\n%v\n",
				ePrefix.String(),
				strToWrite,
				err2.Error())

			return numOfBytesWritten, err

		} else {

			numOfBytesWritten += int64(localNumBytesWritten)
		}

		strToWrite = ""
	}

	return numOfBytesWritten, err
}
