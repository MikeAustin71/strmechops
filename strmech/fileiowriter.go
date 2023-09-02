package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"os"
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

type fileIoWriterMicrobot struct {
	lock *sync.Mutex
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
