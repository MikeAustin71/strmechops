package strmech

import (
	"errors"
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"strings"
	"sync"
)

type FileIoReadWrite struct {
	writer             *FileIoWriter
	reader             *FileIoReader
	writerFilePathName string
	readerFilePathName string

	lock *sync.Mutex
}

// Close
//
// This method is provided in order to implement the
// io.Closer interface.
//
// FileIoReadWrite.Close() effectively performs all
// required Clean-Up tasks. As such, this method should
// only be called after all 'read' and 'write' operations
// have been completed and the services of the current
// FileIoReadWrite instance are no longer required.
//
// After calling this method, FileIoReadWrite.Close(),
// the current instance of FileIoReadWrite will be
// invalid and unavailable for further 'read' and/or
// 'write' operations.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	This method implements the io.Closer interface.
//
//	(2)	After completing all 'read' and 'write' operations,
//		calling this method will:
//
//		(a) Properly 'Close' the internal io.Writer
//			object.
//
//		(b) Properly 'Close' the internal io.Reader
//			object.
//
//		(c)	Release internal memory resources.
//
//			Releasing all internal memory resources will
//			synchronize internal flags and prevent
//			multiple calls to 'close' the underlying
//			io.Reader and io.Writer objects.
//
//			Calling 'close' on the same underlying
//			io.Reader or io.Writer object multiple
//			times can produce unexpected results.
//
//	(3)	Once this method completes all required Clean-Up
//		tasks, this current instance of FileIoReadWrite
//		will become unavailable for further 'read' and/or
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
func (fIoReadWrite *FileIoReadWrite) Close() error {

	if fIoReadWrite.lock == nil {
		fIoReadWrite.lock = new(sync.Mutex)
	}

	fIoReadWrite.lock.Lock()

	defer fIoReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"FileIoReadWrite."+
			"Close()",
		"")

	if err != nil {
		return err
	}

	return new(fileIoReadWriteMicrobot).
		readerWriterCloseRelease(
			fIoReadWrite,
			"fIoReadWrite",
			true, // releaseReaderWriterMemResources
			true, // releaseFIoReadWriteMemResources
			ePrefix.XCpy(
				"Close-Reader&Writer"))
}

// CloseReader
//
// This method is designed to perform Clean-Up tasks
// after completion of all 'read' operations associated
// with the current instance of FileIoReadWrite.
//
// After calling this method, the Clean-Up tasks
// performed will effectively render the internal
// io.Reader object, encapsulated by the current
// FileIoReadWrite instance, invalid and unusable
// for any future 'read' operations.
//
// It is unlikely that the user will ever need to call
// this method. Typically, Clean-Up tasks are performed
// jointly on the internal io.Reader and io.Writer
// objects encapsulated in the current FileIoReadWrite
// instance. These Clean-Up tasks should be performed
// after all 'read' and 'write' operations have been
// completed by calling the local method:
//
//	FileIoReadWrite.Close()
//
// However, in the event of unforeseen use cases, this
// method is provided to exclusively close or Clean-Up
// the io.Reader.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will:
//
//	(1) Properly 'Close' the 'read' file or internal
//		io.Reader object.
//
//	(2) Release the internal memory resources associated
//		with the io.Reader object.
//
//		Releasing internal memory resources synchronizes
//		internal flags and prevents multiple calls to the
//		'close' method. Calling the 'close' method more
//		than once on the internal io.Reader object may
//		produce unexpected results.
//
//	(3) Effectively render the internal io.Reader object,
//		encapsulated by the current instance of
//		FileIoReadWrite, invalid and unusable for any
//		future 'read' operations.
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
func (fIoReadWrite *FileIoReadWrite) CloseReader(
	errorPrefix interface{}) error {

	if fIoReadWrite.lock == nil {
		fIoReadWrite.lock = new(sync.Mutex)
	}

	fIoReadWrite.lock.Lock()

	defer fIoReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileIoReadWrite."+
			"CloseReader()",
		"")

	if err != nil {
		return err
	}

	return new(fileIoReadWriteElectron).
		readerCloseRelease(
			fIoReadWrite,
			"fIoReadWrite",
			true, // releaseReaderMemResources
			true, // releaseFIoReaderLocalMemRes
			ePrefix)
}

// CloseWriter
//
// This method is designed to perform Clean-Up tasks
// after completion of all 'write' operations associated
// with the current instance of FileIoReadWrite.
//
// After calling this method, the Clean-Up tasks
// performed will effectively render the internal
// io.Writer object, encapsulated by the current
// FileIoReadWrite instance, invalid and unusable
// for any future 'write' operations.
//
// It is unlikely that the user will ever need to call
// this method. Typically, Clean-Up tasks are performed
// jointly on the internal io.Reader and io.Writer
// objects encapsulated in the current FileIoReadWrite
// instance. These Clean-Up tasks should be performed
// after all 'read' and 'write' operations have been
// completed by calling the local method:
//
//	FileIoReadWrite.Close()
//
// However, in the event of unforeseen use cases, this
// method is provided to exclusively close or Clean-Up
// the internal io.Writer object.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will:
//
//	(1) Properly 'Close' the 'write' file or internal
//		io.Writer object encapsulated by the current
//		instance of FileIoReadWrite.
//
//	(2) Release the internal memory resources associated
//		with the internal io.Writer object encapsulated
//		by the current instance of FileBufferReadWrite.
//
//		Releasing internal memory resources synchronizes
//		internal flags and prevents multiple calls to the
//		'close' method. Calling the 'close' method more
//		than once on the internal io.Writer object may
//		produce unexpected results.
//
//	(3) Effectively render the internal io.Writer object,
//		encapsulated by the current instance of
//		FileIoReadWrite, invalid and unusable for any
//		future 'write' operations.
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
func (fIoReadWrite *FileIoReadWrite) CloseWriter(
	errorPrefix interface{}) error {

	if fIoReadWrite.lock == nil {
		fIoReadWrite.lock = new(sync.Mutex)
	}

	fIoReadWrite.lock.Lock()

	defer fIoReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileIoReadWrite."+
			"CloseReader()",
		"")

	if err != nil {
		return err
	}

	return new(fileIoReadWriteElectron).
		writerCloseRelease(
			fIoReadWrite,
			"fIoReadWrite",
			true, // releaseWriterMemResources
			true, // releaseFBuffWriterLocalMemRes
			ePrefix)

}

// Empty
//
// This method deletes all internal member variables and
// release all the internal memory resources for the
// current instance of FileIoReadWrite.
//
// Specifically the following internal member variables
// are set to 'nil' or their initial zero values:
//
//	FileIoReadWrite.reader = nil
//	FileIoReadWrite.writer = nil
//	FileIoReadWrite.readerFilePathName = ""
//	FileIoReadWrite.writerFilePathName = ""
//
// After calling this method, the current FileIoReadWrite
// instance will become invalid and unavailable for
// future read/write operations.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	This method does NOT perform the 'close' procedure.
//	To perform the 'close' procedures while
//	simultaneously releasing all internal memory
//	resources, call local method:
//
//			FileBufferReadWrite.Close()
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
func (fIoReadWrite *FileIoReadWrite) Empty() {

	if fIoReadWrite.lock == nil {
		fIoReadWrite.lock = new(sync.Mutex)
	}

	fIoReadWrite.lock.Lock()

	new(fileIoReadWriteElectron).empty(
		fIoReadWrite)

	fIoReadWrite.lock.Unlock()

	fIoReadWrite.lock = nil
}

// IsValidInstanceError
//
// Analyzes the current FileIoReadWrite instance to
// determine if is invalid.
//
// If the current FileIoReadWrite instance is found
// to be invalid, an error is returned.
//
// If the current FileIoReadWrite instance is valid,
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
//		FileIoReadWrite are found to be invalid,
//		this method will return an error configured
//		with an appropriate message identifying the
//		invalid	member data variable.
//
//		If all internal member data variables evaluate
//		as valid, this returned error value will be set
//		to 'nil'.
//
//		If errors are encountered during processing or if
//		any FileIoReadWrite internal member data
//	 	values are found to be invalid, the returned error
//	 	will encapsulate an appropriate error message.
//	 	This returned error message will incorporate the
//	 	method chain and text passed by input parameter,
//	 	'errorPrefix'. The 'errorPrefix' text will be
//	 	prefixed to the beginning of the error message.
func (fIoReadWrite *FileIoReadWrite) IsValidInstanceError(
	errorPrefix interface{}) error {

	if fIoReadWrite.lock == nil {
		fIoReadWrite.lock = new(sync.Mutex)
	}

	fIoReadWrite.lock.Lock()

	defer fIoReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileIoReadWrite."+
			"IsValidInstanceError()",
		"")

	if err != nil {
		return err
	}

	return new(fileIoReadWriteElectron).
		isFileIoReadWriteValid(
			fIoReadWrite,
			"fIoReadWrite",
			ePrefix)
}

// New
//
// This method returns a pointer to an empty or 'blank'
// instance of FileIoReadWrite. All the member variables
// in this returned instance are initialized to their
// zero or initial values. This means the returned
// instance is invalid and unusable for standard 'read'
// and 'write' operations.
//
// This technique for creating a new working instance of
// FileIoReadWrite requires two steps.
//
// Step-1
//
//	Call this method FileIoReadWrite.New() to
//	generate an empty version of FileIoReadWrite.
//
// Step-2
//
//	Use this returned instance of FileIoReadWrite and
//	call the appropriate 'Setter' methods to individually
//	configure the internal 'io.reader' and 'io.writer'
//	objects.
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
//	*FileIoReadWrite
//
//		This method returns a pointer to an empty
//		instance of FileIoReadWrite. After receiving
//		this instance, users must call 'Setter' methods
//		to complete the 'reader' and 'writer'
//		configuration process.
func (fIoReadWrite *FileIoReadWrite) New() *FileIoReadWrite {

	if fIoReadWrite.lock == nil {
		fIoReadWrite.lock = new(sync.Mutex)
	}

	fIoReadWrite.lock.Lock()

	defer fIoReadWrite.lock.Unlock()

	return new(FileIoReadWrite)
}

// NewFileMgrsReadWrite
//
// Creates and returns a pointer to a new, fully
// configured instance of FileIoReadWrite.
//
// The internal io.Reader and io.Writer member variables
// for this new FileIoReadWrite instance are generated
// from input parameters specifying 'reader' and 'writer'
// extracted files from extracted from input parameters
// of type File Manager (FileMgr).
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	readerFileMgr					*FileMgr
//
//		A pointer to an instance of FileMgr. The file
//		identified by 'readerFileMgr' will be used as a
//		data source for 'read' operations performed by
//		the instance of FileIoReadWrite returned by
//		this method
//
//		If the path and file name encapsulated by
//		'readerFileMgr' do not currently exist on an
//		attached storage drive, an error will be
//		returned.
//
//	openReadFileReadWrite			bool
//
//		If this parameter is set to 'true', the target
//		'read' file identified by input parameter
//		'readerFileMgr' will be opened for 'read' and
//		'write' operations.
//
//		If 'openFileReadWrite' is set to 'false', the
//		target read file will be opened for 'read-only'
//		operations.
//
//	defaultReaderByteArraySize		int
//
//		The size of the default byte array which will be
//		used to read data from the internal io.Reader
//		object encapsulated in the instance of
//		FileIoReadWrite returned by this method.
//
//		If the value of 'defaultReaderByteArraySize' is
//		less than '16', it will be reset to a size of
//		'4096'.
//
//		Although the FileIoReadWrite type does not use
//		the 'buffered' read protocol, the size of the
//		byte array used to store bytes read from the
//		underlying io.Reader object is variable.
//
//	writerFileMgr					*FileMgr
//
//		A pointer to an instance of FileMgr. The file
//		identified by 'writerFileMgr' will be used as an
//		output destination for 'write' operations
//		performed by the instance of FileIoReadWrite
//		returned by this method.
//
//		If the path and file name encapsulated by
//		'writerFileMgr' do not currently exist on an
//		attached storage drive, this method will attempt
//		to create them.
//
//	openWriteFileReadWrite			bool
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
//	defaultWriterByteArraySize		int
//
//		The size of the default byte array which will
//		be used to write data to the internal io.Writer
//		object encapsulated by the FileIoReadWrite
//		instance returned by this method.
//
//		If the value of 'defaultWriterByteArraySize' is
//		less than one ('1'), it will be reset to a size
//		of '4096'.
//
//		Although the FileIoReadWrite type does not use
//		the 'buffered' write protocol, the size of the
//		byte array used to store bytes written to the
//		underlying io.Writer object is variable.
//
//	truncateExistingWriteFile		bool
//
//		If this parameter is set to 'true', the target
//		'write' file ('writerFileMgr') will be opened for
//		'write' operations. If the target 'write' file
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
//	errorPrefix						interface{}
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
//	readerFileInfoPlus				FileInfoPlus
//
//		This returned instance of Type FileInfoPlus
//		contains data elements describing the file
//		identified by input parameter 'readerFileMgr'.
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
//	writerFileInfoPlus				FileInfoPlus
//
//		This returned instance of Type FileInfoPlus
//		contains data elements describing the file
//		identified by input parameter 'writerFileMgr'.
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
//	newFileIoReadWrite				*FileIoReadWrite
//
//		If this method completes successfully, it will
//		return a pointer to a fully configured instance
//		of FileIoReadWrite.
//
//	err								error
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
func (fIoReadWrite *FileIoReadWrite) NewFileMgrsReadWrite(
	readerFileMgr *FileMgr,
	openReadFileReadWrite bool,
	defaultReaderByteArraySize int,
	writerFileMgr *FileMgr,
	openWriteFileReadWrite bool,
	defaultWriterByteArraySize int,
	truncateExistingWriteFile bool,
	errorPrefix interface{}) (
	readerFileInfoPlus FileInfoPlus,
	writerFileInfoPlus FileInfoPlus,
	newFileIoReadWrite *FileIoReadWrite,
	err error) {

	if fIoReadWrite.lock == nil {
		fIoReadWrite.lock = new(sync.Mutex)
	}

	fIoReadWrite.lock.Lock()

	defer fIoReadWrite.lock.Unlock()

	newFileIoReadWrite = new(FileIoReadWrite)

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileIoReadWrite."+
			"NewFileMgrsReadWrite()",
		"")

	if err != nil {

		return readerFileInfoPlus,
			writerFileInfoPlus,
			newFileIoReadWrite,
			err
	}

	readerFileInfoPlus,
		writerFileInfoPlus,
		err = new(fileIoReadWriteMicrobot).
		setFileMgrsReadWrite(
			newFileIoReadWrite,
			"newFBuffReadWrite",
			readerFileMgr,
			"readerFileMgr",
			openReadFileReadWrite,
			defaultReaderByteArraySize,
			writerFileMgr,
			"writerFileMgr",
			openWriteFileReadWrite,
			defaultWriterByteArraySize,
			truncateExistingWriteFile,
			ePrefix)

	return readerFileInfoPlus,
		writerFileInfoPlus,
		newFileIoReadWrite,
		err
}

// NewIoReadWrite
//
// Creates and returns a pointer to a new, fully
// configured instance of FileIoReadWrite using
// io.Reader and io.Writer input parameters.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	reader							io.Reader
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
//		FileIoReadWrite.Close()
//
//		While the 'read' services provided by
//		FileIoReadWrite are primarily designed to
//		read data from disk files, this type of 'reader'
//		will in fact read data from any object
//		implementing the io.Reader interface.
//
//	defaultReaderByteArraySize		int
//
//		The size of the byte array which will be used to
//		read data from the internal io.Reader object
//		encapsulated by the returned FileIoReadWrite
//		instance.
//
//		If the value of 'defaultReaderByteArraySize' is
//		less than '16', it will be reset to a size of
//		'4096'.
//
//		Although the FileIoReadWrite type does not use
//		the 'buffered' read protocol, the size of the
//		byte array used to store bytes read from the
//		underlying io.Reader object is variable.
//
//	writer							io.Writer
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
//			FileIoReadWrite.Close()
//
//		While the 'write' services provided by the
//		FileIoReadWrite are primarily designed for
//		writing data to disk files, this type of 'writer'
//		will in fact write data to any object
//		implementing the io.Writer interface.
//
//	defaultWriterByteArraySize		int
//
//		The size of the byte array which will be used to
//		write data to the internal io.Writer object
//		encapsulated by the returned FileIoReadWrite
//		instance.
//
//		If the value of 'defaultWriterByteArraySize' is
//		less than one ('1'), it will be reset to a size
//		of '4096'.
//
//		Although the FileIoReadWrite type does not use
//		the 'buffered' write protocol, the size of the
//		byte array used to store bytes written to the
//		underlying io.Writer object is variable.
//
//	errorPrefix						interface{}
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
//	newFileIoReadWrite			*FileIoReadWrite
//
//		If this method completes successfully, it will
//		return a pointer to a fully configured instance
//		of FileIoReadWrite.
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
func (fIoReadWrite *FileIoReadWrite) NewIoReadWrite(
	reader io.Reader,
	defaultReaderByteArraySize int,
	writer io.Writer,
	defaultWriterByteArraySize int,
	errorPrefix interface{}) (
	newFileIoReadWrite *FileIoReadWrite,
	err error) {

	if fIoReadWrite.lock == nil {
		fIoReadWrite.lock = new(sync.Mutex)
	}

	fIoReadWrite.lock.Lock()

	defer fIoReadWrite.lock.Unlock()

	newFileIoReadWrite = new(FileIoReadWrite)

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileIoReadWrite."+
			"NewIoReadWrite()",
		"")

	if err != nil {
		return newFileIoReadWrite, err
	}

	err = new(fileIoReadWriteMolecule).
		setIoReaderIoWriter(
			newFileIoReadWrite,
			"newFileIoReadWrite",
			reader,
			"reader",
			defaultReaderByteArraySize,
			writer,
			"writer",
			defaultWriterByteArraySize,
			ePrefix)

	return newFileIoReadWrite, err
}

// NewPathFileNamesReadWrite
//
// Creates and returns a pointer to a new, fully
// configured instance of FileIoReadWrite.
//
// The internal io.Reader and io.Writer member variables
// for this new instance of FileIoReadWrite are generated
// from input parameters specifying 'reader' and 'writer'
// files extracted from path and file name strings.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	readerPathFileName				string
//
//		This string contains the path and file name of
//		the file which will be configured as an io.Reader
//		object encapsulated in the FileIoReadWrite
//		instance returned by this method. As such, the
//		file identified by 'readerPathFileName' will be
//		used a data source for 'read' operations.
//
//		If this file does not currently exist on an
//		attached storage drive, an error will be
//		returned.
//
//	openReadFileReadWrite			bool
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
//	defaultReaderByteArraySize		int
//
//		The default size of the byte array which will be
//	 	used to read data from the internal io.Reader
//	 	object encapsulated by the FileIoReadWrite
//	 	instance returned by this method.
//
//		If the value of 'defaultReaderByteArraySize' is
//		less than '16', it will be reset to a size of
//		'4096'.
//
//		Although the FileIoReadWrite type does not use
//		the 'buffered' read protocol, the size of the
//		byte array used to store bytes read from the
//		underlying io.Reader object is variable.
//
//	writerPathFileName				string
//
//		This string contains the path and file name of
//		the target 'write' file which will be used as
//		a data destination for 'write' operations.
//
//		If the target path and file do not currently
//		exist on an attached storage drive, this method
//		will attempt to create them.
//
//	openWriteFileReadWrite			bool
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
//	defaultWriterByteArraySize		int
//
//		The size of the byte array which will be used to
//		write data to the internal io.Writer object
//		encapsulated by the FileIoReadWrite instance
//		returned by this method.
//
//		If the value of 'defaultByteArraySize' is
//		less than one ('1'), it will be reset to a size
//		of '4096'.
//
//		Although the FileIoReadWrite type does not use
//		the 'buffered' write protocol, the size of the
//		byte array used to store bytes written to the
//		underlying io.Writer object is variable.
//
//	truncateExistingWriteFile		bool
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
//	errorPrefix						interface{}
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
//	readerFileInfoPlus				FileInfoPlus
//
//		This returned instance of Type FileInfoPlus
//		contains data elements describing the file
//		identified by input parameter
//		'readerPathFileName'.
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
//	writerFileInfoPlus				FileInfoPlus
//
//		This returned instance of Type FileInfoPlus
//		contains data elements describing the file
//		identified by input parameter
//		'writerPathFileName'.
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
//	newFileIoReadWrite				*FileIoReadWrite
//
//		If this method completes successfully, it will
//		return a pointer to a fully configured instance
//		of FileIoReadWrite.
//
//	err								error
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
func (fIoReadWrite *FileIoReadWrite) NewPathFileNamesReadWrite(
	readerPathFileName string,
	openReadFileReadWrite bool,
	defaultReaderByteArraySize int,
	writerPathFileName string,
	openWriteFileReadWrite bool,
	defaultWriterByteArraySize int,
	truncateExistingWriteFile bool,
	errorPrefix interface{}) (
	readerFileInfoPlus FileInfoPlus,
	writerFileInfoPlus FileInfoPlus,
	newFileIoReadWrite *FileIoReadWrite,
	err error) {

	if fIoReadWrite.lock == nil {
		fIoReadWrite.lock = new(sync.Mutex)
	}

	fIoReadWrite.lock.Lock()

	defer fIoReadWrite.lock.Unlock()

	newFileIoReadWrite = new(FileIoReadWrite)

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileIoReadWrite."+
			"NewPathFileNamesReadWrite()",
		"")

	if err != nil {

		return readerFileInfoPlus,
			writerFileInfoPlus,
			newFileIoReadWrite,
			err
	}

	readerFileInfoPlus,
		writerFileInfoPlus,
		err = new(fileIoReadWriteNanobot).
		setPathFileNamesReadWrite(
			newFileIoReadWrite,
			"newFileIoReadWrite",
			readerPathFileName,
			"readerPathFileName",
			openReadFileReadWrite,
			defaultReaderByteArraySize,
			writerPathFileName,
			"writerPathFileName",
			openWriteFileReadWrite,
			defaultWriterByteArraySize,
			truncateExistingWriteFile,
			ePrefix)

	return readerFileInfoPlus,
		writerFileInfoPlus,
		newFileIoReadWrite,
		err
}

// Read
//
// Reads a selection of data from the internal io.Reader
// data source encapsulated in the current instance of
// FileIoReadWrite.
//
// This method reads data into the input parameter byte
// array, 'bytesRead', from the internal io.Reader object
// encapsulated by the current instance of FileIoReadWrite.
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
//	FileIoReadWrite.Close()
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
//		FileIoReadWrite and the returned error is set to
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
//			FileIoReadWrite.Close()
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
//		in the current instance of FileIoReadWrite.
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
//		'FileIoReadWrite.ioReader' and stored in the byte
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
func (fIoReadWrite *FileIoReadWrite) Read(
	bytesRead []byte) (
	numBytesRead int,
	err error) {

	if fIoReadWrite.lock == nil {
		fIoReadWrite.lock = new(sync.Mutex)
	}

	fIoReadWrite.lock.Lock()

	defer fIoReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"FileIoReadWrite."+
			"Read()",
		"")

	if err != nil {

		return numBytesRead, err
	}

	if fIoReadWrite.reader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of FileIoReadWrite is invalid.\n"+
			"The internal io.Reader object has not been proplery\n"+
			"initialized. FileIoReadWrite.reader == 'nil'\n"+
			"To properly initialize an instance of FileIoReadWrite,\n"+
			"call one or more of the 'New' or 'Setter' methods.\n",
			ePrefix.String())

		return numBytesRead, err
	}

	numBytesRead,
		err = new(fileIoReaderMicrobot).
		readBytes(
			fIoReadWrite.reader,
			"fIoReadWrite.reader",
			bytesRead,
			ePrefix)

	return numBytesRead, err
}

// ReadAllTextLines
//
// Reads text lines from the internal io.Reader object
// encapsulated in the current instance of
// FileIoReadWrite. The entire contents of the io.Reader
// object are parsed and stored as individual lines of
// text in the instance of StringArrayDto passed as input
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
//		FileIoReadWrite, into memory.
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
//			FileIoReadWrite.Close()
//
//	(3)	If the current instance of FileIoReadWrite has NOT
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
//		closing the underlying FileIoReadWrite io.Reader
//		object and releasing all internal memory
//		resources.
//
//		(1)	The FileIoReadWrite internal io.Reader object
//			will be properly closed, the internal memory
//			resources will be released and there will be
//			no need to make a separate call to local
//			method,	FileIoReadWrite.Close().
//
//		(2) After performing this Clean-Up operation, the
//			current instance of FileIoReadWrite will invalid
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
//				FileIoReadWrite.Close()
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
//		FileIoReadWrite.
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
func (fIoReadWrite *FileIoReadWrite) ReadAllTextLines(
	readEndOfLineDelimiters *StringArrayDto,
	outputLinesArray *StringArrayDto,
	maxNumOfTextLines int,
	errorPrefix interface{}) (
	numLinesRead int,
	numBytesRead int64,
	err error) {

	if fIoReadWrite.lock == nil {
		fIoReadWrite.lock = new(sync.Mutex)
	}

	fIoReadWrite.lock.Lock()

	defer fIoReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileIoReadWrite."+
			"SetPathFileNameWriter()",
		"")

	if err != nil {

		return numLinesRead,
			numBytesRead,
			err
	}

	if fIoReadWrite.reader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of FileIoReadWrite is invalid.\n"+
			"The internal io.Reader object has not been proplery\n"+
			"initialized. FileIoReadWrite.reader == 'nil'\n"+
			"To properly initialize an instance of FileIoReadWrite,\n"+
			"call one or more of the 'New' or 'Setter' methods.\n",
			ePrefix.String())

		return numLinesRead,
			numBytesRead,
			err
	}

	numLinesRead,
		numBytesRead,
		err = fIoReadWrite.reader.
		ReadAllTextLines(
			readEndOfLineDelimiters,
			outputLinesArray,
			maxNumOfTextLines,
			false, // autoCloseOnExit
			ePrefix)

	return numLinesRead,
		numBytesRead,
		err
}

// ReadAllToStrBuilder
//
// Reads the entire contents of the internal io.Reader,
// encapsulated in the current instance of
// FileIoReadWrite, as a string. This string is then
// stored and returned through an instance of
// strings.Builder passed as input parameter
// 'strBuilder'.
//
// If a processing error is encountered, an appropriate
// error with an error message will be returned. When
// the end-of-file is encountered during the 'read'
// process, the returned error object will be set to
// 'nil' and no error will be returned.
//
// It naturally follows that this method will read the
// entire contents of the target io.Reader object
// (FileIoReadWrite.reader) into memory when writing
// said contents to the strings.Builder instance
// 'strBuilder'. Depending on the size of the target
// 'read' file, local memory constraints should be
// considered.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	This method is designed to read the entire
//		contents of the internal io.Reader object,
//		encapsulated by the current instance of
//		FileIoReadWrite, into memory.
//
//		BE CAREFUL when reading large files!
//
//		Depending on the memory resources available to
//		your computer, you may run out of memory when
//		reading large files and writing their contents
//		to the strings.Builder input parameter,
//		'strBuilder'.
//
//	(2)	When all read and write operations have been
//		completed for this instance of FileIoReadWrite,
//		the user is responsible for performing required
//		Clean-Up tasks by calling local method:
//
//			FileIoReadWrite.Close()
//
//	(3)	If the current instance of FileIoReadWrite has
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
//		entire contents of the FileIoReadWrite internal
//		io.Reader object and stores the resulting string
//		in this instance of strings.Builder
//		('strBuilder').
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
//		FileIoReadWrite. This value will also equate
//		to the length of the string returned through
//		input parameter, 'strBuilder'.
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
func (fIoReadWrite *FileIoReadWrite) ReadAllToStrBuilder(
	strBuilder *strings.Builder,
	errorPrefix interface{}) (
	numBytesRead int64,
	err error) {

	if fIoReadWrite.lock == nil {
		fIoReadWrite.lock = new(sync.Mutex)
	}

	fIoReadWrite.lock.Lock()

	defer fIoReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileIoReadWrite."+
			"ReadAllToStrBuilder()",
		"")

	if err != nil {

		return numBytesRead, err
	}

	if fIoReadWrite.reader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of FileIoReadWrite is invalid.\n"+
			"The internal io.Reader object has not been proplery\n"+
			"initialized. FileIoReadWrite.reader == 'nil'\n"+
			"To properly initialize an instance of FileIoReadWrite,\n"+
			"call one or more of the 'New' or 'Setter' methods.\n",
			ePrefix.String())

		return numBytesRead, err
	}

	numBytesRead,
		err = new(fileIoReaderMicrobot).
		readAllStrBuilder(
			fIoReadWrite.reader,
			"fIoReadWrite.reader",
			strBuilder,
			false, //autoCloseOnExit
			ePrefix)

	return numBytesRead, err
}

// ReadAllToString
//
// Reads the entire contents of the internal io.Reader
// for the current instance of FileIoReadWrite and returns
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
//		FileIoReadWrite, into memory.
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
//			FileIoReadWrite.Close()
//
//	(3)	If the current instance of FileIoReadWrite has
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
//		FileIoReadWrite.
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
//		FileIoReadWrite will be returned in this string.
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
func (fIoReadWrite *FileIoReadWrite) ReadAllToString(
	errorPrefix interface{}) (
	numBytesRead int64,
	contentsStr string,
	err error) {

	if fIoReadWrite.lock == nil {
		fIoReadWrite.lock = new(sync.Mutex)
	}

	fIoReadWrite.lock.Lock()

	defer fIoReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileIoReadWrite."+
			"ReadAllToString()",
		"")

	if err != nil {

		return numBytesRead,
			contentsStr, err
	}

	if fIoReadWrite.reader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of FileIoReadWrite is invalid.\n"+
			"The internal io.Reader object has not been proplery\n"+
			"initialized. FileIoReadWrite.reader == 'nil'\n"+
			"To properly initialize an instance of FileIoReadWrite,\n"+
			"call one or more of the 'New' or 'Setter' methods.\n",
			ePrefix.String())

		return numBytesRead,
			contentsStr, err
	}

	strBuilder := new(strings.Builder)

	numBytesRead,
		err = new(fileIoReaderMicrobot).
		readAllStrBuilder(
			fIoReadWrite.reader,
			"fIoReader",
			strBuilder,
			false, //autoCloseOnExit
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
// input parameter 'offsetFromBeginning'. The data source
// for this read operation is provided by the internal
// io.Reader object encapsulated by the current instance
// of FileIoReadWrite.
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
//		in the current instance of FileIoReadWrite.
//
//		If the length of this byte array is less than
//		16-bytes, an error will be returned.
//
//	offsetFromBeginning			int64
//
//		The offset in bytes from the beginning of the
//		internal io.Reader input source at which the
//		read 'operation' will commence.
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
//		FileIoReadWrite.
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
func (fIoReadWrite *FileIoReadWrite) ReadAt(
	bytesRead []byte,
	offsetFromBeginning int64) (
	numBytesRead int,
	err error) {

	if fIoReadWrite.lock == nil {
		fIoReadWrite.lock = new(sync.Mutex)
	}

	fIoReadWrite.lock.Lock()

	defer fIoReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"FileIoReadWrite."+
			"ReadAt()",
		"")

	if err != nil {

		return numBytesRead, err
	}

	if fIoReadWrite.reader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of FileIoReadWrite is invalid.\n"+
			"The internal io.Reader object has not been proplery\n"+
			"initialized. FileIoReadWrite.reader == 'nil'\n"+
			"To properly initialize an instance of FileIoReadWrite,\n"+
			"call one or more of the 'New' or 'Setter' methods.\n",
			ePrefix.String())

		return numBytesRead, err
	}

	var err2 error

	numBytesRead,
		err2 = fIoReadWrite.reader.
		ReadAt(
			bytesRead,
			offsetFromBeginning)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by fIoReadWrite.reader.ReadAt()\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			err2.Error())

	}

	return numBytesRead, err
}

// ReadBytesToString
//
// Reads a specified number of bytes from the io.Reader
// data input source configured for the current instance
// of FileIoReadWrite (FileIoReadWrite.reader). These
// bytes are then stored as a string and returned via
// return parameter 'contentsStr'.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	If the current instance of FileIoReadWrite has NOT
//	been properly initialized, an error will be returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numOfBytesToRead			int
//
//		This parameter specifies the number of bytes to
//		read from the io.Reader data input source
//		configured for the current instance of
//		FileIoReadWrite.
//
//		If the value of 'numOfBytesToRead' is less than
//		one ('1'), this parameter will be automatically
//		set to the Default Reader Buffer Size previously
//		configured for this FileIoReadWrite io.Reader
//		instance. For more information on Default Reader
//		Buffer Size, reference local method:
//
//			FileIoReadWrite.SetDefaultReaderByteArraySize()
//
//		The actual number of bytes read from the data input
//		source may vary due to (1) unforeseen processing
//		errors or (2) an End-Of-File scenario. Be sure to
//		check the 'numBytesRead' and 'reachedEndOfFile'
//		parameters returned by this method.
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
//		FileIoReadWrite and stored in the strings.Builder
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
//		the	current instance of FileIoReadWrite will be
//		returned in this string.
//
//	reachedEndOfFile			bool
//
//		If during the 'read' operation, the End-Of-File
//		flag was encountered, this boolean parameter will
//		be set to 'true'. The End-Of-File flag signals that
//		the 'read' operation reached the end of the data
//		input source configured for the current
//		FileIoReadWrite instance.
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
func (fIoReadWrite *FileIoReadWrite) ReadBytesToString(
	numOfBytesToRead int,
	errorPrefix interface{}) (
	numBytesRead int64,
	contentsStr string,
	reachedEndOfFile bool,
	err error) {

	if fIoReadWrite.lock == nil {
		fIoReadWrite.lock = new(sync.Mutex)
	}

	fIoReadWrite.lock.Lock()

	defer fIoReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileIoReadWrite."+
			"ReadAt()",
		"")

	if err != nil {

		return numBytesRead,
			contentsStr, reachedEndOfFile, err
	}

	if fIoReadWrite.reader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of FileIoReadWrite is invalid.\n"+
			"The internal io.Reader object has not been proplery\n"+
			"initialized. FileIoReadWrite.reader == 'nil'\n"+
			"To properly initialize an instance of FileIoReadWrite,\n"+
			"call one or more of the 'New' or 'Setter' methods.\n",
			ePrefix.String())

		return numBytesRead,
			contentsStr, reachedEndOfFile, err
	}

	var strBuilder = new(strings.Builder)

	numBytesRead,
		reachedEndOfFile,
		err = new(fileIoReaderMicrobot).
		readBytesToStrBuilder(
			fIoReadWrite.reader,
			"fIoReadWrite.reader",
			numOfBytesToRead,
			strBuilder,
			false, //autoCloseOnExit,
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
// FileIoReadWrite. These bytes are then stored in an
// instance of strings.Builder passed as input parameter
// 'strBuilder'.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	If the current instance of FileIoReadWrite has NOT been
//	properly initialized, an error will be returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numOfBytesToRead			int
//
//		This parameter specifies the number of bytes to
//		read from the data input source configured for
//		the current instance of FileIoReadWrite.
//
//		If the value of 'numOfBytesToRead' is less than
//		one ('1'), this parameter will be automatically
//		set to the Default Reader Buffer Size previously
//		configured for this FileIoReadWrite instance. For
//		more information on Default Reader Buffer Size,
//		reference local method:
//
//			FileIoReadWrite.SetDefaultReaderByteArraySize()
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
//		FileIoReadWrite and adds those bytes to this
//		instance of strings.Builder.
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
//		FileIoReadWrite and stored in the strings.Builder
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
//		FileIoReadWrite instance.
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
func (fIoReadWrite *FileIoReadWrite) ReadBytesToStringBuilder(
	numOfBytesToRead int,
	strBuilder *strings.Builder,
	errorPrefix interface{}) (
	numBytesRead int64,
	reachedEndOfFile bool,
	err error) {

	if fIoReadWrite.lock == nil {
		fIoReadWrite.lock = new(sync.Mutex)
	}

	fIoReadWrite.lock.Lock()

	defer fIoReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileIoReadWrite."+
			"SetPathFileNameWriter()",
		"")

	if err != nil {

		return numBytesRead,
			reachedEndOfFile,
			err
	}

	if fIoReadWrite.reader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of FileIoReadWrite is invalid.\n"+
			"The internal io.Reader object has not been proplery\n"+
			"initialized. FileIoReadWrite.reader == 'nil'\n"+
			"To properly initialize an instance of FileIoReadWrite,\n"+
			"call one or more of the 'New' or 'Setter' methods.\n",
			ePrefix.String())

		return numBytesRead,
			reachedEndOfFile,
			err
	}

	numBytesRead,
		reachedEndOfFile,
		err = new(fileIoReaderMicrobot).
		readBytesToStrBuilder(
			fIoReadWrite.reader,
			"fIoReadWrite.reader",
			numOfBytesToRead,
			strBuilder,
			false, // autoCloseOnExit,
			ePrefix)

	return numBytesRead, reachedEndOfFile, err
}

// ReadWriteAll
//
// This method will read all data residing in the
// internal io.Reader object and write that data to the
// internal io.Writer object. Both the io.Reader and
// io.Writer objects are encapsulated in the current
// instance of FileIoReadWrite.
//
// The data is read from 'reader' using an internal byte
// array equal in length to the default array size
// previously configured for the io.Reader.
//
// The return parameter 'numOfBytesProcessed' records the
// number of bytes read from 'reader' and written to
// the io.Writer object. If the number of bytes read fails
// to match the number of bytes written, an error will be
// returned.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	If input parameter 'autoCloseOnExit' is set to
//	'false', the user is responsible for performing
//	required Clean-Up operations on the current instance
//	of FileIoReadWrite by calling local method:
//
//		FileIoReadWrite.Close()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	autoCloseOnExit				bool
//
//		If this parameter is set to 'true', the current
//		instance of FileIoReadWrite will be automatically
//		'closed' upon successful completion of this
//		method. This 'close' operation will perform all
//		required Clean-Up tasks on the current
//		FileIoReadWrite instance and users will therefore
//		NOT be required to make a separate call to local
//		method:
//
//				FileIoReadWrite.Close()
//
//		Conversely, if this parameter is set to 'false'
//		the current instance of FileIoReadWrite will NOT
//		be 'closed' and users WILL BE required to make
//		a separate call to local method:
//
//				FileIoReadWrite.Close()
//
//		Once the current instance of FileIoReadWrite has
//		been 'closed' and all Clean-Up operations have
//		been completed, no further 'read' or 'write'
//		operations may be performed using the current
//		FileIoReadWrite instance.
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
//	numOfBytesProcessed			int64
//
//		This return parameter documents the number of
//		bytes read from 'reader' and written to the
//		FileIoReadWrite io.Writer object. If the
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
func (fIoReadWrite *FileIoReadWrite) ReadWriteAll(
	autoCloseOnExit bool,
	errorPrefix interface{}) (
	numOfBytesProcessed int64,
	err error) {

	if fIoReadWrite.lock == nil {
		fIoReadWrite.lock = new(sync.Mutex)
	}

	fIoReadWrite.lock.Lock()

	defer fIoReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileIoReadWrite."+
			"SetPathFileNameWriter()",
		"")

	if err != nil {

		return numOfBytesProcessed, err
	}

	if fIoReadWrite.reader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of FileIoReadWrite is invalid.\n"+
			"The internal io.Reader object has not been proplery\n"+
			"initialized. FileIoReadWrite.reader == 'nil'\n"+
			"To properly initialize an instance of FileIoReadWrite,\n"+
			"call one or more of the 'New' or 'Setter' methods.\n",
			ePrefix.String())

		return numOfBytesProcessed, err
	}

	if fIoReadWrite.writer == nil {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of FileIoReadWrite is invalid.\n"+
			"The internal io.Writer object has not been proplery\n"+
			"initialized. FileIoReadWrite.writer == 'nil'\n"+
			"To properly initialize an instance of FileIoReadWrite,\n"+
			"call one or more of the 'New' or 'Setter' methods.\n",
			ePrefix.String())

		return numOfBytesProcessed, err
	}

	var err2 error

	numOfBytesProcessed,
		err2 = fIoReadWrite.writer.ReadFrom(
		*fIoReadWrite.reader.ioReader)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by fIoReadWrite.writer.ReadFrom().\n"+
			"The error occurred while reading and writing the\n"+
			"entire contents of the internal io.Reader object\n"+
			"encapsulated by the current instance of FileIoReadWrite.\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			err2.Error())

		return numOfBytesProcessed, err
	}

	if autoCloseOnExit == true {

		err = new(fileIoReadWriteMicrobot).
			readerWriterCloseRelease(
				fIoReadWrite,
				"fIoReadWrite",
				true, // releaseReaderWriterMemResources
				true, // releaseFIoReadWriteMemResources
				ePrefix.XCpy(
					"Close-Reader&Writer"))

	}

	return numOfBytesProcessed, err
}

// SeekReader
//
// This method sets the byte offset for the next 'read'
// operation within the internal io.Reader object
// encapsulated in the current FileIoReadWrite instance.
//
// The SeekReader() method only succeeds if the internal
// io.Reader object, encapsulated by the current
// FileIoReadWrite instance, implements the io.Seeker
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
// If the SeekReader() method completes successfully, the
// next 'read' operation will occur at the new offset
// position.
//
// Method SeekReader() returns the new offset relative to
// the start of the io.Reader object or an error, if any.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	If the current FileIoReadWrite instance was NOT
//		initialized with an io.Reader object which
//		implements the io.Seeker interface, this method
//		will return an error.
//
//	(2)	Setting a byte offset which occurs before the
//		start of the internal io.Reader object
//		encapsulated by the current FileIoReadWrite
//		instance will result in an error.
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
//	offsetFromFileStart			int64
//
//		If this method completes successfully, this
//		parameter will return the new file offset
//		in bytes from the beginning of the io.Reader
//		object.
//
//	err							error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message.
func (fIoReadWrite *FileIoReadWrite) SeekReader(
	targetOffset int64,
	whence int,
	errorPrefix interface{}) (
	offsetFromFileStart int64,
	err error) {

	if fIoReadWrite.lock == nil {
		fIoReadWrite.lock = new(sync.Mutex)
	}

	fIoReadWrite.lock.Lock()

	defer fIoReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileIoReadWrite."+
			"SetPathFileNameWriter()",
		"")

	if err != nil {

		return offsetFromFileStart, err
	}

	offsetFromFileStart,
		err = new(fileIoReaderMicrobot).
		seekByOffset(
			fIoReadWrite.reader,
			"fIoReadWrite.reader",
			targetOffset,
			whence,
			ePrefix)

	return offsetFromFileStart, err
}

// SeekWriter
//
// This method sets the byte offset for the next 'write'
// operation within the internal io.Writer object
// encapsulated in the current FileIoReadWrite instance.
//
// The SeekWriter() method only succeeds if the internal
// io.Writer object, encapsulated by the current
// FileIoReadWrite instance, implements the io.Seeker
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
// If the SeekWriter() method completes successfully, the
// next 'write' operation will occur at the new offset
// position.
//
// Method SeekWriter() returns the new offset relative to
// the start of the io.Writer object or an error, if any.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	If the current FileIoReadWrite instance was NOT
//		initialized with an io.Writer object which
//		implements the io.Seeker interface, this method
//		will return an error.
//
//	(2)	Setting a byte offset which occurs before the
//		start of the internal io.Writer object
//		encapsulated by the current FileIoReadWrite
//		instance will result in an error.
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
//	offsetFromFileStart			int64
//
//		If this method completes successfully, this
//		parameter will return the new file offset
//		in bytes from the beginning of the io.Writer
//		object.
//
//	err							error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message.
func (fIoReadWrite *FileIoReadWrite) SeekWriter(
	targetOffset int64,
	whence int,
	errorPrefix interface{}) (
	offsetFromFileStart int64,
	err error) {

	if fIoReadWrite.lock == nil {
		fIoReadWrite.lock = new(sync.Mutex)
	}

	fIoReadWrite.lock.Lock()

	defer fIoReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileIoReadWrite."+
			"SetPathFileNameWriter()",
		"")

	if err != nil {

		return offsetFromFileStart, err
	}

	offsetFromFileStart,
		err = new(fileIoWriterMicrobot).
		seekByOffset(
			fIoReadWrite.writer,
			"fIoReadWrite.writer",
			targetOffset,
			whence,
			ePrefix)

	return offsetFromFileStart, err
}

// SetDefaultReaderByteArraySize
//
// Sets the default size of the byte array used to read
// bytes from the internal io.Reader object encapsulated
// in the current instance of FileIoReadWrite.
//
// Although the FileIoReader type does not use the
// 'buffered' read protocol, the size of the byte array
// used to read and store bytes read from the underlying
// io.Reader object is variable in some cases.
//
// The Default Reader Byte Array Size controls the size
// of the byte array used by the following methods:
//
//	FileIoReadWrite.ReadAllToStrBuilder()
//	FileIoReadWrite.ReadAllToString()
//	FileIoReadWrite.ReadWriteAll()
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
//	errorPrefix						interface{}
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
//	err								error
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
func (fIoReadWrite *FileIoReadWrite) SetDefaultReaderByteArraySize(
	defaultReaderByteArraySize int,
	errorPrefix interface{}) (
	err error) {

	if fIoReadWrite.lock == nil {
		fIoReadWrite.lock = new(sync.Mutex)
	}

	fIoReadWrite.lock.Lock()

	defer fIoReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileIoReadWrite."+
			"SetDefaultReaderByteArraySize()",
		"")

	if err != nil {

		return err
	}

	if fIoReadWrite.reader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of FileIoReadWrite is invalid.\n"+
			"The internal io.Reader object has not been proplery\n"+
			"initialized. FileIoReadWrite.reader == 'nil'\n"+
			"To properly initialize an instance of FileIoReadWrite,\n"+
			"call one or more of the 'New' or 'Setter' methods.\n",
			ePrefix.String())

		return err
	}

	fIoReadWrite.reader.defaultByteArraySize =
		defaultReaderByteArraySize

	new(fileIoReaderMolecule).
		validateDefaultReaderArraySize(
			fIoReadWrite.reader)

	return err
}

// SetDefaultWriterByteArraySize
//
// Sets the default size of the array used to read bytes
// from the internal io.Writer encapsulated in the
// current instance of FileIoReadWrite.
//
// Although the FileIoWriter type does not use the
// 'buffered' write protocol, the size of the byte array
// used to write bytes to the underlying io.Writer object
// is variable is some cases.
//
// The Default Writer Byte Array Size controls the size
// of the byte array used by the following methods:
//
//	FileIoReadWrite.ReadWriteAll()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	defaultWriterByteArraySize		int
//
//		The default size of the byte array which will be
//	 	used to write data to the internal io.Writer
//	 	object encapsulated by the current instance of
//		FileIoReadWrite.
//
//		If the value of 'defaultWriterByteArraySize' is
//		less than  one ('1'), it will be automatically
//		reset to a size of '4096'.
//
//	errorPrefix						interface{}
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
//	err								error
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
func (fIoReadWrite *FileIoReadWrite) SetDefaultWriterByteArraySize(
	defaultWriterByteArraySize int,
	errorPrefix interface{}) (
	err error) {

	if fIoReadWrite.lock == nil {
		fIoReadWrite.lock = new(sync.Mutex)
	}

	fIoReadWrite.lock.Lock()

	defer fIoReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileIoReadWrite."+
			"SetDefaultWriterByteArraySize()",
		"")

	if err != nil {

		return err
	}

	if fIoReadWrite.reader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of FileIoReadWrite is invalid.\n"+
			"The internal io.Writer object has not been proplery\n"+
			"initialized. FileIoReadWrite.reader == 'nil'\n"+
			"To properly initialize an instance of FileIoReadWrite,\n"+
			"call one or more of the 'New' or 'Setter' methods.\n",
			ePrefix.String())

		return err
	}

	fIoReadWrite.reader.defaultByteArraySize =
		defaultWriterByteArraySize

	new(fileIoWriterMolecule).
		validateDefaultByteArraySize(
			fIoReadWrite.writer)

	return err
}

// SetFileMgrsReadWrite
//
// Receives two input parameters of type FileMgr. These
// File Manager instances will be used to reconfigure the
// internal io.Reader and io.Writer objects encapsulated
// by the current instance of FileIoReadWrite.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the current instance of
//	FileIoReadWrite.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	readerFileMgr					*FileMgr
//
//		A pointer to an instance of FileMgr. The file
//		identified by 'readerFileMgr' will be used as a
//		data source for all future 'read' operations
//	 	performed by the current instance of
//	 	FileIoReadWrite.
//
//		If the path and file name encapsulated by
//		'readerFileMgr' do not currently exist on an
//		attached storage drive, an error will be
//		returned.
//
//	openReadFileReadWrite			bool
//
//		If this parameter is set to 'true', the target
//		'read' file identified by input parameter
//		'readerFileMgr' will be opened for 'read' and
//		'write' operations.
//
//		If 'openFileReadWrite' is set to 'false', the
//		target read file will be opened for 'read-only'
//		operations.
//
//	defaultReaderByteArraySize		int
//
//		The size of the default byte array which will be
//		used to read data from the internal io.Reader
//		object encapsulated in the current instance of
//		FileIoReadWrite.
//
//		If the value of 'defaultReaderByteArraySize' is
//		less than '16', it will be reset to a size of
//		'4096'.
//
//		Although the FileIoReadWrite type does not use
//		the 'buffered' read protocol, the size of the
//		byte array used to store bytes read from the
//		underlying io.Writer object is variable.
//
//	writerFileMgr					*FileMgr
//
//		A pointer to an instance of FileMgr. The file
//		identified by 'writerFileMgr' will be used as an
//		output destination for all future 'write'
//		operations performed by the current instance of
//		FileIoReadWrite.
//
//		If the path and file name encapsulated by
//		'writerFileMgr' do not currently exist on an
//		attached storage drive, this method will attempt
//		to create them.
//
//	openWriteFileReadWrite			bool
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
//	defaultWriterByteArraySize		int
//
//		The size of the default byte array which will
//		be used to write data to the internal io.Writer
//		object encapsulated by the current instance of
//		FileIoReadWrite.
//
//		If the value of 'defaultWriterByteArraySize' is
//		less than one ('1'), it will be reset to a size
//		of '4096'.
//
//		Although the FileIoReadWrite type does not use
//		the 'buffered' write protocol, the size of the
//		byte array used to store bytes written to the
//		underlying io.Writer object is variable.
//
//	truncateExistingWriteFile		bool
//
//		If this parameter is set to 'true', the target
//		'write' file ('writerFileMgr') will be opened for
//		'write' operations. If the target 'write' file
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
//	errorPrefix						interface{}
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
//	readerFileInfoPlus				FileInfoPlus
//
//		This returned instance of Type FileInfoPlus
//		contains data elements describing the file
//		identified by input parameter 'readerFileMgr'.
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
//	writerFileInfoPlus				FileInfoPlus
//
//		This returned instance of Type FileInfoPlus
//		contains data elements describing the file
//		identified by input parameter 'writerFileMgr'.
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
//	err								error
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
func (fIoReadWrite *FileIoReadWrite) SetFileMgrsReadWrite(
	readerFileMgr *FileMgr,
	openReadFileReadWrite bool,
	defaultReaderByteArraySize int,
	writerFileMgr *FileMgr,
	openWriteFileReadWrite bool,
	defaultWriterByteArraySize int,
	truncateExistingWriteFile bool,
	errorPrefix interface{}) (
	readerFileInfoPlus FileInfoPlus,
	writerFileInfoPlus FileInfoPlus,
	err error) {

	if fIoReadWrite.lock == nil {
		fIoReadWrite.lock = new(sync.Mutex)
	}

	fIoReadWrite.lock.Lock()

	defer fIoReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileIoReadWrite."+
			"SetFileMgrsReadWrite()",
		"")

	if err != nil {

		return readerFileInfoPlus,
			writerFileInfoPlus,
			err
	}

	readerFileInfoPlus,
		writerFileInfoPlus,
		err = new(fileIoReadWriteMicrobot).
		setFileMgrsReadWrite(
			fIoReadWrite,
			"fIoReadWrite",
			readerFileMgr,
			"readerFileMgr",
			openReadFileReadWrite,
			defaultReaderByteArraySize,
			writerFileMgr,
			"writerFileMgr",
			openWriteFileReadWrite,
			defaultWriterByteArraySize,
			truncateExistingWriteFile,
			ePrefix)

	return readerFileInfoPlus,
		writerFileInfoPlus,
		err
}

// SetFileMgrReader
//
// This method will close, delete and reconfigure the
// internal io.Reader object encapsulated in the current
// instance of FileIoReadWrite. This internal io.Reader
// object will be reconfigured using the file identified
// by a File Manager instance passed as input parameter
// 'readerFileMgr'.
//
// The internal io.Reader object is used to 'read' data
// from a data source such as a disk file.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reconfigure
//	the member variable io.Reader object encapsulated in
//	the current instance of FileIoReadWrite.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	readerFileMgr					*FileMgr
//
//		A pointer to an instance of FileMgr. The file
//		identified by 'readerFileMgr' will be used as a
//		data source for all future 'read' operations
//	 	performed by the current instance of
//	 	FileIoReadWrite.
//
//		If the path and file name encapsulated by
//		'readerFileMgr' do not currently exist on an
//		attached storage drive, an error will be
//		returned.
//
//	openReadFileReadWrite			bool
//
//		If this parameter is set to 'true', the target
//		'read' file identified by input parameter
//		'readerFileMgr' will be opened for 'read' and
//		'write' operations.
//
//		If 'openFileReadWrite' is set to 'false', the
//		target read file will be opened for 'read-only'
//		operations.
//
//	defaultReaderByteArraySize		int
//
//		The size of the default byte array which will be
//		used to read data from the internal io.Reader
//		object encapsulated in the current instance of
//		FileIoReadWrite.
//
//		If the value of 'defaultReaderByteArraySize' is
//		less than '16', it will be reset to a size of
//		'4096'.
//
//		Although the FileIoReadWrite type does not use
//		the 'buffered' read protocol, the size of the
//		byte array used to store bytes read from the
//		underlying io.Writer object is variable.
//
//	errorPrefix						interface{}
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
//	readerFileInfoPlus				FileInfoPlus
//
//		This returned instance of Type FileInfoPlus
//		contains data elements describing the file
//		identified by input parameter 'readerFileMgr'.
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
//	err								error
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
func (fIoReadWrite *FileIoReadWrite) SetFileMgrReader(
	readerFileMgr *FileMgr,
	openReadFileReadWrite bool,
	defaultReaderByteArraySize int,
	errorPrefix interface{}) (
	readerFileInfoPlus FileInfoPlus,
	err error) {

	if fIoReadWrite.lock == nil {
		fIoReadWrite.lock = new(sync.Mutex)
	}

	fIoReadWrite.lock.Lock()

	defer fIoReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileIoReadWrite."+
			"SetFileMgrReader()",
		"")

	if err != nil {

		return readerFileInfoPlus, err
	}

	readerFileInfoPlus,
		err = new(fileIoReadWriteNanobot).
		setFileMgrReader(
			fIoReadWrite,
			"fIoReadWrite",
			readerFileMgr,
			"readerFileMgr",
			openReadFileReadWrite,
			defaultReaderByteArraySize,
			ePrefix)

	return readerFileInfoPlus, err
}

// SetFileMgrWriter
//
// This method will close, delete and reconfigure the
// internal io.Writer object encapsulated in the current
// instance of FileIoReadWrite. This internal io.Writer
// object will be reconfigured using the file identified
// by a File Manager instance passed as input parameter
// 'writerFileMgr'.
//
// The internal io.Writer object is used to 'write' data
// to an output destination such as a disk file.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reconfigure
//	the member variable io.Writer object encapsulated
//	in the current instance of FileIoReadWrite.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	writerFileMgr					*FileMgr
//
//		A pointer to an instance of FileMgr. The file
//		identified by 'writerFileMgr' will be used as an
//		output destination for all future 'write'
//		operations performed by the current instance of
//		FileIoReadWrite.
//
//		If the path and file name encapsulated by
//		'writerFileMgr' do not currently exist on an
//		attached storage drive, this method will attempt
//		to create them.
//
//	openWriteFileReadWrite			bool
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
//	defaultWriterByteArraySize		int
//
//		The size of the default byte array which will
//		be used to write data to the internal io.Writer
//		object encapsulated by the current instance of
//		FileIoReadWrite.
//
//		If the value of 'defaultWriterByteArraySize' is
//		less than one ('1'), it will be reset to a size
//		of '4096'.
//
//		Although the FileIoReadWrite type does not use
//		the 'buffered' write protocol, the size of the
//		byte array used to store bytes written to the
//		underlying io.Writer object is variable.
//
//	truncateExistingWriteFile		bool
//
//		If this parameter is set to 'true', the target
//		'write' file ('writerFileMgr') will be opened for
//		'write' operations. If the target 'write' file
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
//	errorPrefix						interface{}
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
//	writerFileInfoPlus				FileInfoPlus
//
//		This returned instance of Type FileInfoPlus
//		contains data elements describing the file
//		identified by input parameter 'writerFileMgr'.
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
//	err								error
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
func (fIoReadWrite *FileIoReadWrite) SetFileMgrWriter(
	writerFileMgr *FileMgr,
	openWriteFileReadWrite bool,
	defaultWriterByteArraySize int,
	truncateExistingWriteFile bool,
	errorPrefix interface{}) (
	writerFileInfoPlus FileInfoPlus,
	err error) {

	if fIoReadWrite.lock == nil {
		fIoReadWrite.lock = new(sync.Mutex)
	}

	fIoReadWrite.lock.Lock()

	defer fIoReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileIoReadWrite."+
			"SetFileMgrWriter()",
		"")

	if err != nil {

		return writerFileInfoPlus, err
	}

	writerFileInfoPlus,
		err = new(fileIoReadWriteNanobot).
		setFileMgrWriter(
			fIoReadWrite,
			"fIoReadWrite",
			writerFileMgr,
			"writerFileMgr",
			openWriteFileReadWrite,
			defaultWriterByteArraySize,
			truncateExistingWriteFile,
			ePrefix)

	return writerFileInfoPlus, err
}

// SetIoReadWrite
//
// This method will close, delete and reconfigure the
// internal io.Reader and io.Writer objects encapsulated
// in the current instance of FileIoReadWrite.
//
// The internal io.Reader object is used to 'read' data
// from a data source such as a disk file. This internal
// io.Reader object will be reconfigured using the
// io.Reader object passed as input parameter 'reader'.
//
// In contrast, the internal io.Writer object is used
// to 'write' data to an output destination such as a
// disk file. This internal io.Writer object will be
// reconfigured using the io.Writer object passed as
// input parameter 'writer'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing member data values in the current
//	instance of FileIoReadWrite.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	reader							io.Reader
//
//		An object which implements io.Reader interface.
//		This object will be used to reconfigure the
//		internal io.Reader object encapsulated by the
//		current FileIoReadWrite instance.
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
//			FileIoReadWrite.Close()
//
//		While the 'read' services provided by
//		FileIoReadWrite are primarily designed to
//		read data from disk files, this type of 'reader'
//		will in fact read data from any object
//		implementing the io.Reader interface.
//
//	defaultReaderByteArraySize		int
//
//		The size of the byte array which will be used to
//		read data from the internal io.Reader object
//		encapsulated by the current instance of
//		FileIoReadWrite.
//
//		If the value of 'defaultReaderByteArraySize' is
//		less than '16', it will be reset to a size of
//		'4096'.
//
//		Although the FileIoReadWrite type does not use
//		the 'buffered' read protocol, the size of the
//		byte array used to store bytes read from the
//		underlying io.Reader object is variable.
//
//	writer							io.Writer
//
//		This parameter will accept any object
//		implementing the io.Writer interface.
//		This 'writer' object will be used to reconfigure
//		the internal io.Writer object encapsulated by the
//		current FileIoReadWrite instance.
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
//		FileIoReadWrite.Close()
//
//		While the 'write' services provided by the
//		FileIoReadWrite are primarily designed for
//		writing data to disk files, this type of 'writer'
//		will in fact write data to any object
//		implementing the io.Writer interface.
//
//	defaultWriterByteArraySize		int
//
//		The size of the byte array which will be used to
//		write data to the internal io.Writer object
//		encapsulated by the current FileIoReadWrite
//		instance.
//
//		If the value of 'defaultWriterByteArraySize' is
//		less than one ('1'), it will be reset to a size
//		of '4096'.
//
//		Although the FileIoReadWrite type does not use
//		the 'buffered' write protocol, the size of the
//		byte array used to write bytes to the underlying
//		io.Writer object is variable.
//
//	errorPrefix						interface{}
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
func (fIoReadWrite *FileIoReadWrite) SetIoReadWrite(
	reader io.Reader,
	defaultReaderByteArraySize int,
	writer io.Writer,
	defaultWriterByteArraySize int,
	errorPrefix interface{}) (
	err error) {

	if fIoReadWrite.lock == nil {
		fIoReadWrite.lock = new(sync.Mutex)
	}

	fIoReadWrite.lock.Lock()

	defer fIoReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileIoReadWrite."+
			"SetIoReadWrite()",
		"")

	if err != nil {

		return err

	}

	err = new(fileIoReadWriteMolecule).
		setIoReaderIoWriter(
			fIoReadWrite,
			"fIoReadWrite",
			reader,
			"reader",
			defaultReaderByteArraySize,
			writer,
			"writer",
			defaultWriterByteArraySize,
			ePrefix)

	return err
}

// SetIoReader
//
// This method will close, delete and reconfigure the
// internal io.Reader object encapsulated in the current
// instance of FileIoReadWrite. The internal io.Reader
// object will be reconfigured using the io.Reader object
// passed as input parameter 'reader'.
//
// The internal io.Reader object, encapsulated in the
// current instance of FileIoReadWrite, is used to 'read'
// data from a data source such as a disk file.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reconfigure
//	the internal io.Reader object encapsulated in the
//	current instance of FileIoReadWrite.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	reader							io.Reader
//
//		An object which implements io.Reader interface.
//		This object will be used to reconfigure the
//		internal io.Reader object encapsulated by the
//		current FileIoReadWrite instance.
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
//			FileIoReadWrite.Close()
//
//		While the 'read' services provided by
//		FileIoReadWrite are primarily designed to
//		read data from disk files, this type of 'reader'
//		will in fact read data from any object
//		implementing the io.Reader interface.
//
//	defaultReaderByteArraySize		int
//
//		The size of the byte array which will be used to
//		read data from the internal io.Reader object
//		encapsulated by the current instance of
//		FileIoReadWrite.
//
//		If the value of 'defaultReaderByteArraySize' is
//		less than '16', it will be reset to a size of
//		'4096'.
//
//		Although the FileIoReadWrite type does not use
//		the 'buffered' read protocol, the size of the
//		byte array used to store bytes read from the
//		underlying io.Reader object is variable.
//
//	errorPrefix						interface{}
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
func (fIoReadWrite *FileIoReadWrite) SetIoReader(
	reader io.Reader,
	defaultReaderByteArraySize int,
	errorPrefix interface{}) (
	err error) {

	if fIoReadWrite.lock == nil {
		fIoReadWrite.lock = new(sync.Mutex)
	}

	fIoReadWrite.lock.Lock()

	defer fIoReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileIoReadWrite."+
			"SetIoReader()",
		"")

	if err != nil {

		return err

	}

	return new(fileIoReadWriteAtom).
		setIoReader(
			fIoReadWrite,
			"fIoReadWrite",
			reader,
			"reader",
			defaultReaderByteArraySize,
			ePrefix)
}

// SetIoWriter
//
// This method will close, delete and reconfigure the
// internal io.Writer object encapsulated in the current
// instance of FileIoReadWrite. The internal io.Writer
// object will be reconfigured using the io.Writer object
// passed as input parameter 'writer'.
//
// The internal io.Writer object, encapsulated in the
// current instance of FileIoReadWrite, is used to
// 'write' data to an output destination such as a disk
// file.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reconfigure
//	the internal io.Writer object encapsulated in the
//	current instance of FileIoReadWrite.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	writer							io.Writer
//
//		This parameter will accept any object
//		implementing the io.Writer interface. This
//		'writer' object will be used to reconfigure the
//		internal io.Writer object encapsulated by the
//		current FileIoReadWrite instance.
//
//		This object may be a file pointer of type
//		*os.File. File pointers of this type implement
//		the io.Writer interface.
//
//		A file pointer (*os.File) will facilitate writing
//		data to files residing on an attached storage
//		drive. However, with this configuration, the user
//		is responsible for manually closing the file and
//		performing any other required clean-up operations
//		in addition to calling local method:
//
//		FileIoReadWrite.Close()
//
//		While the 'write' services provided by the
//		FileIoReadWrite are primarily designed for
//		writing data to disk files, this type of 'writer'
//		will in fact write data to any object
//		implementing the io.Writer interface.
//
//	defaultWriterByteArraySize		int
//
//		The size of the byte array which will be used to
//		write data to the internal io.Writer object
//		encapsulated by the current FileIoReadWrite
//		instance.
//
//		If the value of 'defaultWriterByteArraySize' is
//		less than one ('1'), it will be reset to a size
//		of '4096'.
//
//		Although the FileIoReadWrite type does not use
//		the 'buffered' write protocol, the size of the
//		byte array used to write bytes to the underlying
//		io.Writer object is variable.
//
//	errorPrefix						interface{}
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
func (fIoReadWrite *FileIoReadWrite) SetIoWriter(
	writer io.Writer,
	defaultWriterByteArraySize int,
	errorPrefix interface{}) (
	err error) {

	if fIoReadWrite.lock == nil {
		fIoReadWrite.lock = new(sync.Mutex)
	}

	fIoReadWrite.lock.Lock()

	defer fIoReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileIoReadWrite."+
			"SetIoWriter()",
		"")

	if err != nil {

		return err

	}

	return new(fileIoReadWriteAtom).
		setIoWriter(
			fIoReadWrite,
			"fIoReadWrite",
			writer,
			"writer",
			defaultWriterByteArraySize,
			ePrefix)
}

// SetPathFileNamesReadWrite
//
// Receives two input parameter strings identifying the
// path and file names for the 'read' and 'write' files.
// These two files will be used to reconfigure the
// internal io.Reader and io.Writer objects encapsulated
// by the current instance of FileIoReadWrite.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the current instance of
//	FileIoReadWrite.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	readerPathFileName				string
//
//		This string contains the path and file name of
//		the file which will be configured as an io.Reader
//		object encapsulated in the current instance of
//		FileIoReadWrite. As such, the file identified by
//		'readerPathFileName' will be used a data source
//		for all future 'read' operations.
//
//		If this file does not currently exist on an
//		attached storage drive, an error will be
//		returned.
//
//	openReadFileReadWrite			bool
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
//	defaultReaderByteArraySize		int
//
//		The default size of the byte array which will be
//	 	used to read data from the internal io.Reader
//	 	object encapsulated by current instance of
//	 	FileIoReadWrite.
//
//		If the value of 'defaultReaderByteArraySize' is
//		less than '16', it will be reset to a size of
//		'4096'.
//
//		Although the FileIoReadWrite type does not use
//		the 'buffered' read protocol, the size of the
//		byte array used to store bytes read from the
//		underlying io.Reader object is variable.
//
//	writerPathFileName				string
//
//		This string contains the path and file name of
//		the target 'write' file which will be used as
//		a data destination for 'write' operations by
//		the current instance of FileIoReadWrite.
//
//		If the target path and file do not currently
//		exist on an attached storage drive, this method
//		will attempt to create them.
//
//	openWriteFileReadWrite			bool
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
//	defaultWriterByteArraySize		int
//
//		The size of the byte array which will be used to
//		write data to the internal io.Writer object
//		encapsulated by the current instance of
//		FileIoReadWrite.
//
//		If the value of 'defaultByteArraySize' is
//		less than one ('1'), it will be reset to a size
//		of '4096'.
//
//		Although the FileIoReadWrite type does not use
//		the 'buffered' write protocol, the size of the
//		byte array used to store bytes written to the
//		underlying io.Writer object is variable.
//
//	truncateExistingWriteFile		bool
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
//	errorPrefix						interface{}
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
//	readerFileInfoPlus				FileInfoPlus
//
//		This returned instance of Type FileInfoPlus
//		contains data elements describing the file
//		identified by input parameter
//		'readerPathFileName'.
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
//	writerFileInfoPlus				FileInfoPlus
//
//		This returned instance of Type FileInfoPlus
//		contains data elements describing the file
//		identified by input parameter
//		'writerPathFileName'.
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
//	err								error
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
func (fIoReadWrite *FileIoReadWrite) SetPathFileNamesReadWrite(
	readerPathFileName string,
	openReadFileReadWrite bool,
	defaultReaderByteArraySize int,
	writerPathFileName string,
	openWriteFileReadWrite bool,
	defaultWriterByteArraySize int,
	truncateExistingWriteFile bool,
	errorPrefix interface{}) (
	readerFileInfoPlus FileInfoPlus,
	writerFileInfoPlus FileInfoPlus,
	err error) {

	if fIoReadWrite.lock == nil {
		fIoReadWrite.lock = new(sync.Mutex)
	}

	fIoReadWrite.lock.Lock()

	defer fIoReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileIoReadWrite."+
			"SetPathFileNamesReadWrite()",
		"")

	if err != nil {

		return readerFileInfoPlus,
			writerFileInfoPlus,
			err
	}

	readerFileInfoPlus,
		writerFileInfoPlus,
		err = new(fileIoReadWriteNanobot).
		setPathFileNamesReadWrite(
			fIoReadWrite,
			"fIoReadWrite",
			readerPathFileName,
			"readerPathFileName",
			openReadFileReadWrite,
			defaultReaderByteArraySize,
			writerPathFileName,
			"writerPathFileName",
			openWriteFileReadWrite,
			defaultWriterByteArraySize,
			truncateExistingWriteFile,
			ePrefix)

	return readerFileInfoPlus,
		writerFileInfoPlus,
		err
}

//	SetPathFileNameReader
//
// Receives an input parameter string specifying the path
// and file name identifying the file which will be
// configured as a data source for 'read' operations
// performed by the current instance of FileIoReadWrite.
//
// This 'read' file will be configured as an internal
// io.Reader object encapsulated in the current instance
// of FileIoReadWrite.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reconfigure
//	the member variable io.Reader object encapsulated in
//	the current instance of FileIoReadWrite.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	readerPathFileName				string
//
//		This string contains the path and file name of
//		the file which will be configured as an io.Reader
//		object encapsulated in the current instance of
//		FileIoReadWrite. As such, the file identified by
//		'readerPathFileName' will be used a data source
//		for all future 'read' operations.
//
//		If this file does not currently exist on an
//		attached storage drive, an error will be
//		returned.
//
//	openReadFileReadWrite			bool
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
//	defaultReaderByteArraySize		int
//
//		The default size of the byte array which will be
//	 	used to read data from the internal io.Reader
//	 	object encapsulated by current instance of
//	 	FileIoReadWrite ('readerPathFileName').
//
//		If the value of 'defaultReaderByteArraySize' is
//		less than '16', it will be reset to a size of
//		'4096'.
//
//		Although the FileIoReadWrite type does not use
//		the 'buffered' read protocol, the size of the
//		byte array used to store bytes read from the
//		underlying io.Reader object is variable.
//
//	errorPrefix						interface{}
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
//	readerFileInfoPlus				FileInfoPlus
//
//		This returned instance of Type FileInfoPlus
//		contains data elements describing the file
//		identified by input parameter
//		'readerPathFileName'.
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
//	err								error
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
func (fIoReadWrite *FileIoReadWrite) SetPathFileNameReader(
	readerPathFileName string,
	openReadFileReadWrite bool,
	defaultReaderByteArraySize int,
	errorPrefix interface{}) (
	readerFileInfoPlus FileInfoPlus,
	err error) {

	if fIoReadWrite.lock == nil {
		fIoReadWrite.lock = new(sync.Mutex)
	}

	fIoReadWrite.lock.Lock()

	defer fIoReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileIoReadWrite."+
			"SetPathFileNameReader()",
		"")

	if err != nil {

		return readerFileInfoPlus, err
	}

	readerFileInfoPlus,
		err = new(fileIoReadWriteAtom).
		setPathFileNameReader(
			fIoReadWrite,
			"fIoReadWrite",
			readerPathFileName,
			"readerPathFileName",
			openReadFileReadWrite,
			defaultReaderByteArraySize,
			ePrefix)

	return readerFileInfoPlus, err
}

// SetPathFileNameWriter
//
// Receives an input parameter string specifying the path
// and file name identifying the file which will be
// configured as an output destination for 'write'
// operations performed by the current instance of
// FileIoReadWrite.
//
// This 'write' file will be configured as an internal
// io.Writer object encapsulated in the current instance
// of FileIoReadWrite.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reconfigure
//	the member variable io.Writer object encapsulated in
//	the current instance of FileIoReadWrite.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	writerPathFileName				string
//
//		This string contains the path and file name of
//		the target 'write' file which will be used as
//		a data destination for 'write' operations by
//		the current instance of FileIoReadWrite.
//
//		If the target path and file do not currently
//		exist on an attached storage drive, this method
//		will attempt to create them.
//
//	openWriteFileReadWrite			bool
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
//	defaultWriterByteArraySize		int
//
//		The size of the byte array which will be used to
//		write data to the internal io.Writer object
//		encapsulated by the current instance of
//		FileIoReadWrite.
//
//		If the value of 'defaultByteArraySize' is
//		less than one ('1'), it will be reset to a size
//		of '4096'.
//
//		Although the FileIoReadWrite type does not use
//		the 'buffered' write protocol, the size of the
//		byte array used to store bytes written to the
//		underlying io.Writer object is variable.
//
//	truncateExistingWriteFile		bool
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
//	errorPrefix						interface{}
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
//	writerFileInfoPlus				FileInfoPlus
//
//		This returned instance of Type FileInfoPlus
//		contains data elements describing the file
//		identified by input parameter
//		'writerPathFileName'.
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
//	err								error
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
func (fIoReadWrite *FileIoReadWrite) SetPathFileNameWriter(
	writerPathFileName string,
	openWriteFileReadWrite bool,
	defaultWriterByteArraySize int,
	truncateExistingWriteFile bool,
	errorPrefix interface{}) (
	writerFileInfoPlus FileInfoPlus,
	err error) {

	if fIoReadWrite.lock == nil {
		fIoReadWrite.lock = new(sync.Mutex)
	}

	fIoReadWrite.lock.Lock()

	defer fIoReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileIoReadWrite."+
			"SetPathFileNameWriter()",
		"")

	if err != nil {

		return writerFileInfoPlus, err
	}

	writerFileInfoPlus,
		err = new(fileIoReadWriteAtom).
		setPathFileNameWriter(
			fIoReadWrite,
			"fIoReadWrite",
			writerPathFileName,
			"writerPathFileName",
			openWriteFileReadWrite,
			defaultWriterByteArraySize,
			truncateExistingWriteFile,
			ePrefix)

	return writerFileInfoPlus, err
}

type fileIoReadWriteMicrobot struct {
	lock *sync.Mutex
}

// readerWriterCloseRelease
//
// This method will perform all required Clean-Up
// operations on an instance of FileIoReadWrite
// passed as input parameter 'fIoReadWrite'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	After calling this method the FileIoReadWrite
//	instance passed as 'fIoReadWrite' will be invalid
//	and unavailable for any future read/write operations.
//
//	The specific Clean-Up procedures performed by this
//	method are listed as follows:
//
//	(1)	Closing the internal io.reader object.
//
//	(2)	Closing the internal io.writer object.
//
//	(3) Releasing all internal memory resources.
//		This action will synchronize internal flags and
//		prevent multiple calls to 'close' methods.
//		Performing a 'close' operation multiple times
//		on a single io.reader or io.writer object can
//		produce unexpected results.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fIoReadWrite						*FileIoReadWrite
//
//		A pointer to an instance of FileIoReadWrite.
//
//		This method will perform Clean-Up operations on
//		the internal io.reader and io.writer objects
//		encapsulated in	this FileIoReadWrite instance.
//
//		This method will effectively render the
//		FileIoReadWrite instance, 'fIoReadWrite',
//		invalid and unusable for any future 'read' and/or
//		'write' operations.
//
//	fIoReadWriteLabel					string
//
//		The name or label associated with input parameter
//		'fIoReadWrite' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fIoReadWrite" will
//		be automatically applied.
//
//	releaseReaderWriterMemResources		bool
//
//		If this parameter is set to 'true', this method
//		will release the memory resources for the
//		internal io.reader and io.writer objects
//		encapsulated by 'fIoReadWrite'.
//
//	releaseFIoReadWriteMemResources		bool
//
//		If this parameter is set to 'true', this method
//		will release the internal memory resources for
//		the FileIoReadWrite instance passed as
//		'fIoReadWrite'.
//
//	errPrefDto							*ePref.ErrPrefixDto
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
func (fIoReadWriteMicrobot *fileIoReadWriteMicrobot) readerWriterCloseRelease(
	fIoReadWrite *FileIoReadWrite,
	fIoReadWriteLabel string,
	releaseReaderWriterMemResources bool,
	releaseFIoReadWriteMemResources bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fIoReadWriteMicrobot.lock == nil {
		fIoReadWriteMicrobot.lock = new(sync.Mutex)
	}

	fIoReadWriteMicrobot.lock.Lock()

	defer fIoReadWriteMicrobot.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileIoReadWriteMicrobot." +
		"readerWriterCloseRelease()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return err
	}

	if len(fIoReadWriteLabel) == 0 {

		fIoReadWriteLabel = "fIoReadWrite"
	}

	if fIoReadWrite == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is a nil pointer!\n"+
			"%v is invalid.\n",
			ePrefix.String(),
			fIoReadWriteLabel,
			fIoReadWriteLabel)

		return err
	}

	var fIoReadWriteElectron = new(fileIoReadWriteElectron)

	var err2, err3 error

	err2 = fIoReadWriteElectron.
		readerCloseRelease(
			fIoReadWrite,
			fIoReadWriteLabel,
			releaseReaderWriterMemResources,
			releaseFIoReadWriteMemResources,
			ePrefix)

	if err2 != nil {

		err3 = fmt.Errorf("%v\n"+
			"Error occurred while closing the %v.\n"+
			"Error:\n%v\n",
			funcName,
			fIoReadWriteLabel+".reader",
			err2.Error())

		err = errors.Join(err3)
	}

	err2 = fIoReadWriteElectron.
		writerCloseRelease(
			fIoReadWrite,
			fIoReadWriteLabel,
			releaseReaderWriterMemResources,
			releaseFIoReadWriteMemResources,
			ePrefix)

	if err2 != nil {

		err3 = fmt.Errorf("%v\n"+
			"Error occurred while closing the %v.\n"+
			"Error:\n%v\n",
			funcName,
			fIoReadWriteLabel+".writer",
			err2.Error())

		err = errors.Join(err3)
	}

	return err
}

// setFileMgrsReadWrite
//
// Receives two instances of FileMgr as input parameters
// identifying the internal io.Reader and io.Writer
// objects which will be configured for the
// FileIoReadWrite instance passed as input parameter
// 'fIoReadWrite'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the instance of
//	FileIoReadWrite passed as input parameter
//	'fIoReadWrite'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fIoReadWrite					*FileIoReadWrite
//
//		A pointer to an instance of FileIoReadWrite.
//
//		The internal io.Reader and io.Writer objects
//		encapsulated in this instance will be deleted and
//		reinitialized using the 'read' and 'write' files
//		passed as input parameters 'readerFileMgr' and
//		'writerFileMgr'.
//
//	fIoReadWriteLabel				string
//
//		The name or label associated with input parameter
//		'fIoReadWrite' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fIoReadWrite" will
//		be automatically applied.
//
//	readerFileMgr					*FileMgr
//
//		A pointer to an instance of FileMgr. The file
//		identified by 'readerFileMgr' will be used as a
//		data source for 'read' operations performed by
//		the instance of FileIoReadWrite passed as
//		input parameter 'fIoReadWrite'.
//
//		If the path and file name encapsulated by
//		'readerFileMgr' do not currently exist on an
//		attached storage drive, an error will be
//		returned.
//
//	readerFileMgrLabel				string
//
//		The name or label associated with input parameter
//		'readerFileMgr' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "readerFileMgr" will
//		be automatically applied.
//
//	openReadFileReadWrite			bool
//
//		If this parameter is set to 'true', the target
//		'read' file identified by input parameter
//		'readerFileMgr' will be opened for 'read' and
//		'write' operations.
//
//		If 'openFileReadWrite' is set to 'false', the
//		target read file will be opened for 'read-only'
//		operations.
//
//	defaultReaderByteArraySize		int
//
//		The size of the default byte array which will be
//		used to read data from the internal io.Reader
//		object encapsulated by the FileIoReadWrite
//		instance passed as input parameter
//		'fIoReadWrite'.
//
//		If the value of 'defaultReaderByteArraySize' is
//		less than '16', it will be reset to a size of
//		'4096'.
//
//	writerFileMgr					*FileMgr
//
//		A pointer to an instance of FileMgr. The file
//		identified by 'writerFileMgr' will be used as an
//		output destination for 'write' operations
//		performed by the instance of FileIoReadWrite
//		passed as input parameter 'fIoReadWrite'.
//
//		If the path and file name encapsulated by
//		'writerFileMgr' do not currently exist on an
//		attached storage drive, this method will attempt
//		to create them.
//
//	writerFileMgrLabel				string
//
//		The name or label associated with input parameter
//		'writerFileMgr' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "writerFileMgr" will
//		be automatically applied.
//
//	openWriteFileReadWrite			bool
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
//	defaultWriterByteArraySize		int
//
//		The size of the default byte array which will be
//		used to write data to the internal io.Writer
//		object encapsulated by the FileIoReadWrite
//		instance passed as input parameter
//		'fIoReadWrite'.
//
//		If the value of 'defaultWriterByteArraySize' is
//		less than one ('1'), it will be reset to a size
//		of '4096'.
//
//	truncateExistingWriteFile		bool
//
//		If this parameter is set to 'true', the target
//		'write' file ('writerFileMgr') will be opened for
//		'write' operations. If the target 'write' file
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
//	errPrefDto						*ePref.ErrPrefixDto
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
//	readerFileInfoPlus			FileInfoPlus
//
//		This returned instance of Type FileInfoPlus
//		contains data elements describing the file
//		identified by input parameter 'readerFileMgr'.
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
//	writerFileInfoPlus			FileInfoPlus
//
//		This returned instance of Type FileInfoPlus
//		contains data elements describing the file
//		identified by input parameter 'writerFileMgr'.
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
func (fIoReadWriteMicrobot *fileIoReadWriteMicrobot) setFileMgrsReadWrite(
	fIoReadWrite *FileIoReadWrite,
	fIoReadWriteLabel string,
	readerFileMgr *FileMgr,
	readerFileMgrLabel string,
	openReadFileReadWrite bool,
	defaultReaderByteArraySize int,
	writerFileMgr *FileMgr,
	writerFileMgrLabel string,
	openWriteFileReadWrite bool,
	defaultWriterByteArraySize int,
	truncateExistingWriteFile bool,
	errPrefDto *ePref.ErrPrefixDto) (
	readerFileInfoPlus FileInfoPlus,
	writerFileInfoPlus FileInfoPlus,
	err error) {

	if fIoReadWriteMicrobot.lock == nil {
		fIoReadWriteMicrobot.lock = new(sync.Mutex)
	}

	fIoReadWriteMicrobot.lock.Lock()

	defer fIoReadWriteMicrobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileIoReadWriteMicrobot." +
		"setFileMgrsReadWrite()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		return readerFileInfoPlus,
			writerFileInfoPlus,
			err
	}

	if fIoReadWrite == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is a nil pointer!\n"+
			"FileIoReadWrite instance %v is invalid.\n",
			ePrefix.String(),
			fIoReadWriteLabel,
			fIoReadWriteLabel)

		return readerFileInfoPlus, writerFileInfoPlus, err
	}

	var fIoReadWriteNanobot = new(fileIoReadWriteNanobot)

	readerFileInfoPlus,
		err = fIoReadWriteNanobot.
		setFileMgrReader(
			fIoReadWrite,
			fIoReadWriteLabel,
			readerFileMgr,
			readerFileMgrLabel,
			openReadFileReadWrite,
			defaultReaderByteArraySize,
			ePrefix)

	if err != nil {

		return readerFileInfoPlus, writerFileInfoPlus, err
	}

	writerFileInfoPlus,
		err = fIoReadWriteNanobot.
		setFileMgrWriter(
			fIoReadWrite,
			fIoReadWriteLabel,
			writerFileMgr,
			writerFileMgrLabel,
			openWriteFileReadWrite,
			defaultWriterByteArraySize,
			truncateExistingWriteFile,
			ePrefix)

	return readerFileInfoPlus, writerFileInfoPlus, err
}

type fileIoReadWriteNanobot struct {
	lock *sync.Mutex
}

// setFileMgrReader
//
// Receives an instance of File Manager (FileMgr) which
// will be used to configure an io.Reader object
// encapsulated by the FileIoReadWrite instance
// passed as input parameter 'fIoReadWrite'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reconfigure
//	the member variable io.Reader object encapsulated in
//	the instance of FileIoReadWrite passed as input
//	parameter 'fIoReadWrite':
//
//		fIoReadWrite.reader
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fIoReadWrite					*FileIoReadWrite
//
//		A pointer to an instance of FileIoReadWrite.
//
//		The internal io.Reader object encapsulated in
//		this instance of FileIoReadWrite will be
//		deleted and reconfigured using the FileMgr
//		instance passed as input parameter
//		'readerFileMgr'.
//
//	fIoReadWriteLabel				string
//
//		The name or label associated with input parameter
//		'fIoReadWrite' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fIoReadWrite" will
//		be automatically applied.
//
//	readerFileMgr					*FileMgr
//
//		A pointer to an instance of FileMgr. The file
//		identified by 'readerFileMgr' will be used as a
//		data source for 'read' operations and will be
//		configured as an internal io.Reader for the
//		FileIoReadWrite instance passed as input
//		parameter 'fIoReadWrite'.
//
//		If the path and file name encapsulated by
//		'readerFileMgr' do not currently exist on an
//		attached storage drive, an error will be
//		returned.
//
//	readerLabel						string
//
//		The name or label associated with input parameter
//		'reader' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "readerFileMgr" will
//		be automatically applied.
//
//	openReadFileReadWrite			bool
//
//		If this parameter is set to 'true', the target
//		'read' file identified by input parameter
//		'readerFileMgr' will be opened for both 'read'
//		and 'write' operations.
//
//		If 'openReadFileReadWrite' is set to 'false', the
//		target 'read' file ('readerFileMgr') will be
//		opened for 'read-only' operations.
//
//	defaultReaderByteArraySize		int
//
//		The size of the default byte array which will be
//		used to read data from the internal io.Reader
//		object encapsulated by the FileIoReadWrite
//		instance passed as input parameter
//		'fIoReadWrite'.
//
//		If the value of 'defaultReaderByteArraySize' is
//		less than '16', it will be reset to a size of
//		'4096'.
//
//	errPrefDto						*ePref.ErrPrefixDto
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
//	readerFileInfoPlus				FileInfoPlus
//
//		This returned instance of Type FileInfoPlus
//		contains data elements describing the file
//		identified by input parameter 'readerFileMgr'.
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
//	err								error
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
func (fIoReadWriteNanobot *fileIoReadWriteNanobot) setFileMgrReader(
	fIoReadWrite *FileIoReadWrite,
	fIoReadWriteLabel string,
	readerFileMgr *FileMgr,
	readerFileMgrLabel string,
	openReadFileReadWrite bool,
	defaultReaderByteArraySize int,
	errPrefDto *ePref.ErrPrefixDto) (
	readerFileInfoPlus FileInfoPlus,
	err error) {

	if fIoReadWriteNanobot.lock == nil {
		fIoReadWriteNanobot.lock = new(sync.Mutex)
	}

	fIoReadWriteNanobot.lock.Lock()

	defer fIoReadWriteNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fIoReadWriteNanobot." +
		"setFileMgrReader()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		return readerFileInfoPlus, err
	}

	if len(fIoReadWriteLabel) == 0 {

		fIoReadWriteLabel = "fIoReadWrite"
	}

	if len(readerFileMgrLabel) == 0 {

		readerFileMgrLabel = "readerFileMgr"
	}

	if fIoReadWrite == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is a nil pointer!\n"+
			"%v is invalid.\n",
			ePrefix.String(),
			fIoReadWriteLabel,
			fIoReadWriteLabel)

		return readerFileInfoPlus, err
	}

	var err2 error

	err2 = new(fileMgrHelperAtom).isFileMgrValid(
		readerFileMgr,
		ePrefix.XCpy(readerFileMgrLabel))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter %v is invalid.\n"+
			"Reader FileMgr %v failed the validity test.\n"+
			"Error=\n%v\n",
			funcName,
			readerFileMgrLabel,
			readerFileMgrLabel,
			err2.Error())

		return readerFileInfoPlus, err
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

		return readerFileInfoPlus, err
	}

	readerFileMgrLabel += ".absolutePathFileName"

	readerFileInfoPlus,
		err = new(fileIoReadWriteAtom).
		setPathFileNameReader(
			fIoReadWrite,
			fIoReadWriteLabel,
			readerFileMgr.absolutePathFileName,
			readerFileMgrLabel,
			openReadFileReadWrite,
			defaultReaderByteArraySize,
			ePrefix)

	return readerFileInfoPlus, err
}

// setFileMgrWriter
//
// Receives an instance of File Manager (FileMgr) which
// will be used to configure an io.Writer object
// encapsulated by the FileIoReadWrite instance
// passed as input parameter 'fIoReadWrite'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reconfigure
//	the member variable io.Writer object encapsulated in
//	the instance of FileIoReadWrite passed as input
//	parameter 'fIoReadWrite':
//
//		fIoReadWrite.writer
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fIoReadWrite					*FileIoReadWrite
//
//		A pointer to an instance of FileIoReadWrite.
//
//		The internal io.Writer object encapsulated in
//		this instance of FileIoReadWrite will be
//		deleted and reconfigured using the FileMgr
//		instance passed as input parameter
//		'writerFileMgr'.
//
//	fIoReadWriteLabel				string
//
//		The name or label associated with input parameter
//		'fIoReadWrite' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fIoReadWrite" will
//		be automatically applied.
//
//	writerFileMgr					*FileMgr
//
//		A pointer to an instance of FileMgr. The file
//		identified by 'writerFileMgr' will be used as a
//		data source for 'write' operations and will be
//		configured as an internal io.Writer for the
//		FileIoReadWrite instance passed as input
//		parameter 'fIoReadWrite'.
//
//		If the path and file name encapsulated by
//		'writerFileMgr' do not currently exist on an
//		attached storage drive, an error will be
//		returned.
//
//	writerFileMgrLabel				string
//
//		The name or label associated with input parameter
//		'writer' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "writerFileMgr" will
//		be automatically applied.
//
//	openWriteFileReadWrite			bool
//
//		If this parameter is set to 'true', the target
//		'read' file identified by input parameter
//		'writerFileMgr' will be opened for both 'read'
//		and 'write' operations.
//
//		If 'openWriteFileReadWrite' is set to 'false', the
//		target 'read' file will be opened for 'read-only'
//		operations.
//
//	defaultWriterByteArraySize		int
//
//		The size of the default byte array which will be
//		used to write data to the internal io.Writer
//		object encapsulated by the FileIoReadWrite
//		instance passed as input parameter 'fIoReadWrite'.
//
//		If the value of 'defaultWriterByteArraySize' is
//		less than one ('1'), it will be reset to a size
//		of '4096'.
//
//	truncateExistingWriteFile		bool
//
//		If this parameter is set to 'true', the target
//		'write' file ('writerFileMgr') will be opened for
//		write operations. If the target write file
//		previously existed, it will be truncated. This
//		means that the file's previous contents will be
//		deleted.
//
//		If this parameter is set to 'false', the target
//		'write' file ('writerFileMgr') will be opened for
//		write operations. If the target 'write' file
//		previously existed, the new text written to this
//		file will be appended to the end of the previous
//		file contents.
//
//	errPrefDto						*ePref.ErrPrefixDto
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
//	writerFileInfoPlus				FileInfoPlus
//
//		This returned instance of Type FileInfoPlus
//		contains data elements describing the file
//		identified by input parameter 'writerFileMgr'.
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
//	err								error
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
func (fIoReadWriteNanobot *fileIoReadWriteNanobot) setFileMgrWriter(
	fIoReadWrite *FileIoReadWrite,
	fIoReadWriteLabel string,
	writerFileMgr *FileMgr,
	writerFileMgrLabel string,
	openWriteFileReadWrite bool,
	defaultWriterByteArraySize int,
	truncateExistingWriteFile bool,
	errPrefDto *ePref.ErrPrefixDto) (
	writerFileInfoPlus FileInfoPlus,
	err error) {

	if fIoReadWriteNanobot.lock == nil {
		fIoReadWriteNanobot.lock = new(sync.Mutex)
	}

	fIoReadWriteNanobot.lock.Lock()

	defer fIoReadWriteNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fIoReadWriteNanobot." +
		"setFileMgrWriter()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		return writerFileInfoPlus, err
	}

	if len(fIoReadWriteLabel) == 0 {

		fIoReadWriteLabel = "fIoReadWrite"
	}

	if len(writerFileMgrLabel) == 0 {

		writerFileMgrLabel = "writerFileMgr"
	}

	if fIoReadWrite == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is a nil pointer!\n"+
			"%v is invalid.\n",
			ePrefix.String(),
			fIoReadWriteLabel,
			fIoReadWriteLabel)

		return writerFileInfoPlus, err
	}

	var err2 error

	err2 = new(fileMgrHelperAtom).isFileMgrValid(
		writerFileMgr,
		ePrefix.XCpy(writerFileMgrLabel))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter %v is invalid.\n"+
			"Writer FileMgr %v failed the validity test.\n"+
			"Error=\n%v\n",
			funcName,
			writerFileMgrLabel,
			writerFileMgrLabel,
			err2.Error())

		return writerFileInfoPlus, err
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

		return writerFileInfoPlus, err
	}

	writerFileMgrLabel += ".absolutePathFileName"

	writerFileInfoPlus,
		err = new(fileIoReadWriteAtom).
		setPathFileNameWriter(
			fIoReadWrite,
			fIoReadWriteLabel,
			writerFileMgr.absolutePathFileName,
			writerFileMgrLabel,
			openWriteFileReadWrite,
			defaultWriterByteArraySize,
			truncateExistingWriteFile,
			ePrefix)

	return writerFileInfoPlus, err
}

// setPathFileNamesReadWrite
//
// Receives two strings as input parameters for the path
// and file names identifying the io.Reader and io.Writer
// objects which will be configured for the
// FileIoReadWrite instance passed as input parameter
// 'fIoReadWrite'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the instance of
//	FileIoReadWrite passed as input parameter
//	'fIoReadWrite'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fIoReadWrite					*FileIoReadWrite
//
//		A pointer to an instance of FileIoReadWrite.
//
//		The internal FileIoReader and FileIoWriter
//		objects encapsulated in this instance will be
//		deleted and reinitialized using the path and file
//		names passed as input parameters
//		'readerPathFileName' and 'writerPathFileName'.
//
//	fIoReadWriteLabel				string
//
//		The name or label associated with input parameter
//		'fIoReadWrite' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fIoReadWrite" will
//		be automatically applied.
//
//	readerPathFileName				string
//
//		This string contains the path and file name of
//		the file which will be configured as an io.Reader
//		object encapsulated in the FileIoReadWrite
//		instance passed as input parameter
//		'fIoReadWrite'. As such, the file identified by
//		'readerPathFileName' will be used a data source
//		for 'read' operations.
//
//		If this file does not currently exist on an
//		attached storage drive, an error will be
//		returned.
//
//	readerPathFileNameLabel			string
//
//		The name or label associated with input parameter
//		'readerPathFileName' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "readerPathFileName"
//		will be automatically applied.
//
//	openReadFileReadWrite			bool
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
//	defaultReaderByteArraySize		int
//
//		The default size of the byte array which will be
//		used to read data from the internal io.Reader
//		object encapsulated by the FileIoReadWrite
//		instance passed as input parameter 'fIoReadWrite'.
//
//		If the value of 'defaultReaderByteArraySize' is
//		less than '16', it will be reset to a size of
//		'4096'.
//
//	writerPathFileName				string
//
//		This string contains the path and file name of
//		the target 'write' file which will be used as
//		a data destination for 'write' operations.
//
//		If the target path and file do not currently
//		exist on an attached storage drive, this method
//		will attempt to create them.
//
//	writerPathFileNameLabel			string
//
//		The name or label associated with input parameter
//		'writerPathFileName' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "writerPathFileName"
//		will be automatically applied.
//
//	openWriteFileReadWrite			bool
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
//	defaultWriterByteArraySize		int
//
//		The default size of the byte array which will be
//		used to write data to the internal io.Writer
//		object encapsulated by FileIoWriter input
//		parameter 'fIoWriter'.
//
//		If the value of 'defaultByteArraySize' is
//		less than one ('1'), it will be reset to a size
//		of '4096'.
//
//	truncateExistingWriteFile		bool
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
//	errPrefDto						*ePref.ErrPrefixDto
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
//	readerFileInfoPlus				FileInfoPlus
//
//		This returned instance of Type FileInfoPlus
//		contains data elements describing the file
//		identified by input parameter
//		'readerPathFileName'.
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
//	writerFileInfoPlus				FileInfoPlus
//
//		This returned instance of Type FileInfoPlus
//		contains data elements describing the file
//		identified by input parameter
//		'writerPathFileName'.
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
//	err								error
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
func (fIoReadWriteNanobot *fileIoReadWriteNanobot) setPathFileNamesReadWrite(
	fIoReadWrite *FileIoReadWrite,
	fIoReadWriteLabel string,
	readerPathFileName string,
	readerPathFileNameLabel string,
	openReadFileReadWrite bool,
	defaultReaderByteArraySize int,
	writerPathFileName string,
	writerPathFileNameLabel string,
	openWriteFileReadWrite bool,
	defaultWriterByteArraySize int,
	truncateExistingWriteFile bool,
	errPrefDto *ePref.ErrPrefixDto) (
	readerFileInfoPlus FileInfoPlus,
	writerFileInfoPlus FileInfoPlus,
	err error) {

	if fIoReadWriteNanobot.lock == nil {
		fIoReadWriteNanobot.lock = new(sync.Mutex)
	}

	fIoReadWriteNanobot.lock.Lock()

	defer fIoReadWriteNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fIoReadWriteNanobot." +
		"setPathFileNamesReadWrite()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		return readerFileInfoPlus,
			writerFileInfoPlus,
			err
	}

	var fIoReadWriteAtom = new(fileIoReadWriteAtom)

	readerFileInfoPlus,
		err = fIoReadWriteAtom.
		setPathFileNameReader(
			fIoReadWrite,
			fIoReadWriteLabel,
			readerPathFileName,
			readerPathFileNameLabel,
			openReadFileReadWrite,
			defaultReaderByteArraySize,
			ePrefix.XCpy(
				fIoReadWriteLabel+".reader"))

	if err != nil {

		return readerFileInfoPlus,
			writerFileInfoPlus,
			err
	}

	writerFileInfoPlus,
		err = fIoReadWriteAtom.
		setPathFileNameWriter(
			fIoReadWrite,
			fIoReadWriteLabel,
			writerPathFileName,
			writerPathFileNameLabel,
			openWriteFileReadWrite,
			defaultWriterByteArraySize,
			truncateExistingWriteFile,
			ePrefix.XCpy(
				fIoReadWriteLabel+".writer"))

	return readerFileInfoPlus,
		writerFileInfoPlus,
		err
}

type fileIoReadWriteMolecule struct {
	lock *sync.Mutex
}

// setIoReaderIoWriter
//
// Receives two objects which implements io.Reader and
// io.Writer interfaces. These objects are then used to
// configure the internal io.Reader and io.Writer member
// variable encapsulated in the FileIoReadWrite instance
// passed as input parameter 'fIoReadWrite'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reconfigure
//	the member variable io.Reader and io.Writer objects
//	encapsulated in the instance of FileIoReadWrite
//	passed as input parameter 'fIoReadWrite':
//
//		fIoReadWrite.reader
//		fIoReadWrite.writer
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fIoReadWrite					*FileIoReadWrite
//
//		A pointer to an instance of FileIoReadWrite.
//
//		The internal io.Reader and io.Writer objects
//		encapsulated in this instance of FileIoReadWrite
//		will be deleted and reconfigured using the io.Reader
//		and io.Writer instances passed as input parameters
//		'ioReader' and 'ioWriter'.
//
//	fIoReadWriteLabel				string
//
//		The name or label associated with input parameter
//		'fIoReadWrite' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fIoReadWrite" will
//		be automatically applied.
//
//	ioReader						io.Reader
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
//		in addition to calling the local method:
//
//		FileIoReadWrite.CloseRelease()
//
//		While the 'read' services provided by
//		FileIoReadWrite are primarily designed to read
//		data from disk files, this type of 'reader' will
//		in fact read data from any object implementing
//		the io.Reader interface.
//
//	ioReaderLabel					string
//
//		The name or label associated with input parameter
//		'ioReader' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "ioReader" will be
//		automatically applied.
//
//	defaultReaderByteArraySize		int
//
//		The size of the byte array which will be used to
//		read data from the internal io.Reader object
//		encapsulated by the FileIoReadWrite instance
//		passed as input parameter 'fIoReadWrite'.
//
//		If the value of 'defaultReaderByteArraySize' is
//		less than '16', it will be reset to a size of
//		'4096'.
//
//	ioWriter						io.Writer
//
//		This parameter will accept any object
//		implementing the io.Writer interface.
//
//		This object may be a file pointer of type
//		*os.File. File pointers of this type implement
//		the io.Writer interface.
//
//		A file pointer (*os.File) will facilitate writing
//		output data to destination files residing on an
//		attached storage drive. However, with this
//		configuration, the user is responsible for
//		manually closing the file and performing any
//		other required clean-up operations in addition to
//		calling local method:
//
//		FileIoReadWrite.CloseRelease()
//
//		While the 'write' services provided by the
//		FileIoReadWrite are primarily designed for
//		writing data to disk files, this type of 'writer'
//		will in fact write data to any object
//		implementing the io.Writer interface.
//
//	ioWriterLabel					string
//
//		The name or label associated with input parameter
//		'ioWriter' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "ioWriter" will be
//		automatically applied.
//
//	defaultWriterByteArraySize		int
//
//		The size of the byte array which will be used to
//		write data to the internal io.Writer object
//		encapsulated by FileIoReadWrite input parameter
//		'fIoReadWrite'.
//
//		If the value of 'defaultWriterByteArraySize' is
//		less than one ('1'), it will be reset to a size
//		of '4096'.
//
//		Although the FileIoReadWrite type does not use
//		the 'buffered' write protocol, the size of the
//		byte array used to write bytes to the underlying
//		io.Writer object is variable.
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
func (fIoReadWriteMolecule *fileIoReadWriteMolecule) setIoReaderIoWriter(
	fIoReadWrite *FileIoReadWrite,
	fIoReadWriteLabel string,
	ioReader io.Reader,
	ioReaderLabel string,
	defaultReaderByteArraySize int,
	ioWriter io.Writer,
	ioWriterLabel string,
	defaultWriterByteArraySize int,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fIoReadWriteMolecule.lock == nil {
		fIoReadWriteMolecule.lock = new(sync.Mutex)
	}

	fIoReadWriteMolecule.lock.Lock()

	defer fIoReadWriteMolecule.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileIoReadWriteMolecule." +
		"setIoReaderIoWriter()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return err
	}

	if len(fIoReadWriteLabel) == 0 {

		fIoReadWriteLabel = "fIoReadWrite"
	}

	if len(ioReaderLabel) == 0 {

		ioReaderLabel = "ioReader"
	}

	if len(ioWriterLabel) == 0 {

		ioWriterLabel = "ioWriter"
	}

	if fIoReadWrite == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is a nil pointer!\n"+
			"%v is invalid.\n",
			ePrefix.String(),
			fIoReadWriteLabel,
			fIoReadWriteLabel)

		return err
	}

	if ioReader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: The io.Reader instance passed\n"+
			"as input parameter '%v' is invalid!\n"+
			"'%v' is a 'nil' pointer.\n",
			ePrefix.String(),
			ioReaderLabel,
			ioReaderLabel)

		return err
	}

	if ioWriter == nil {

		err = fmt.Errorf("%v\n"+
			"Error: The io.Writer instance passed\n"+
			"as input parameter '%v' is invalid!\n"+
			"'%v' is a 'nil' pointer.\n",
			ePrefix.String(),
			ioWriterLabel,
			ioWriterLabel)

		return err
	}

	var fIoReadWriteAtom = new(fileIoReadWriteAtom)

	err = fIoReadWriteAtom.
		setIoReader(
			fIoReadWrite,
			fIoReadWriteLabel,
			ioReader,
			ioReaderLabel,
			defaultReaderByteArraySize,
			ePrefix)

	if err != nil {

		return err
	}

	err = fIoReadWriteAtom.
		setIoWriter(
			fIoReadWrite,
			fIoReadWriteLabel,
			ioWriter,
			ioWriterLabel,
			defaultWriterByteArraySize,
			ePrefix)

	return err
}

type fileIoReadWriteAtom struct {
	lock *sync.Mutex
}

// setIoReader
//
// Receives an object which implements io.Reader
// interface. This object is then used to configure
// the internal io.Reader member variable encapsulated in
// the FileIoReadWrite instance passed as input parameter
// 'fIoReadWrite'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reconfigure
//	the member variable io.Reader object encapsulated in
//	the instance of FileIoReadWrite passed as input
//	parameter 'fIoReadWrite':
//
//		fIoReadWrite.reader
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fIoReadWrite					*FileIoReadWrite
//
//		A pointer to an instance of FileIoReadWrite.
//
//		The internal io.Reader object encapsulated in
//		this instance of FileIoReadWrite will be
//		deleted and reconfigured using the io.Reader
//		instance passed as input parameter 'ioReader'.
//
//	fIoReadWriteLabel				string
//
//		The name or label associated with input parameter
//		'fIoReadWrite' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fIoReadWrite" will
//		be automatically applied.
//
//	ioReader						io.Reader
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
//		in addition to calling the local method:
//
//		FileIoReadWrite.CloseRelease()
//
//		While the 'read' services provided by
//		FileIoReadWrite are primarily designed to read
//		data from disk files, this type of 'reader' will
//		in fact read data from any object implementing
//		the io.Reader interface.
//
//	ioReaderLabel					string
//
//		The name or label associated with input parameter
//		'ioReader' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "ioReader" will be
//		automatically applied.
//
//	defaultReaderByteArraySize		int
//
//		The size of the byte array which will be used to
//		read data from the internal io.Reader object
//		encapsulated by the FileIoReadWrite instance
//		passed as input parameter 'fIoReadWrite'.
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
func (fIoReadWriteAtom *fileIoReadWriteAtom) setIoReader(
	fIoReadWrite *FileIoReadWrite,
	fIoReadWriteLabel string,
	ioReader io.Reader,
	ioReaderLabel string,
	defaultReaderByteArraySize int,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fIoReadWriteAtom.lock == nil {
		fIoReadWriteAtom.lock = new(sync.Mutex)
	}

	fIoReadWriteAtom.lock.Lock()

	defer fIoReadWriteAtom.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileIoReadWriteAtom." +
		"setIoReader()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return err
	}

	if len(fIoReadWriteLabel) == 0 {

		fIoReadWriteLabel = "fIoReadWrite"
	}

	if len(ioReaderLabel) == 0 {

		ioReaderLabel = "ioReader"
	}

	if fIoReadWrite == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is a nil pointer!\n"+
			"%v is invalid.\n",
			ePrefix.String(),
			fIoReadWriteLabel,
			fIoReadWriteLabel)

		return err
	}

	if ioReader == nil {

		err = fmt.Errorf("%v\n"+
			"Error: The io.Reader instance passed\n"+
			"as input parameter '%v' is invalid!\n"+
			"'%v' is a 'nil' pointer.\n",
			ePrefix.String(),
			ioReaderLabel,
			ioReaderLabel)

		return err
	}

	// Close and release the old FileIoReadWrite.reader
	err = new(fileIoReadWriteElectron).
		readerCloseRelease(
			fIoReadWrite,
			fIoReadWriteLabel,
			true, // releaseReaderMemResources
			true, // releaseFBuffWriterLocalMemRes
			ePrefix)

	if err != nil {

		return err
	}

	fIoReadWrite.reader = nil
	fIoReadWrite.readerFilePathName = ""

	var newIoReader FileIoReader
	var err2 error

	err2 =
		new(fileIoReaderNanobot).
			setIoReader(
				&newIoReader,
				fIoReadWriteLabel+".reader",
				ioReader,
				ioReaderLabel,
				defaultReaderByteArraySize,
				ePrefix)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"An error occurred while creating the new %v.reader.\n"+
			"Error=\n%v\n",
			funcName,
			fIoReadWriteLabel,
			err2.Error())

		return err
	}

	fIoReadWrite.reader = &newIoReader

	fIoReadWrite.readerFilePathName =
		newIoReader.targetReadFileName

	return err
}

// setIoWriter
//
// Receives an object which implements io.Writer
// interface. This object is then used to configure
// the internal io.Writer member variable encapsulated in
// the FileIoReadWrite instance passed as input parameter
// 'fIoReadWrite'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reconfigure
//	the member variable io.Writer object encapsulated in
//	the instance of FileIoReadWrite passed as input
//	parameter 'fIoReadWrite':
//
//		fIoReadWrite.writer
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fIoReadWrite					*FileIoReadWrite
//
//		A pointer to an instance of FileIoReadWrite.
//
//		The internal io.Writer object encapsulated in
//		this instance of FileIoReadWrite will be
//		deleted and reconfigured using the io.Writer
//		object passed as input parameter 'ioWriter'.
//
//	fIoReadWriteLabel				string
//
//		The name or label associated with input parameter
//		'fIoReadWrite' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fIoReadWrite" will
//		be automatically applied.
//
//	ioWriter						io.Writer
//
//		This parameter will accept any object
//		implementing the io.Writer interface.
//
//		This object may be a file pointer of type
//		*os.File. File pointers of this type implement
//		the io.Writer interface.
//
//		A file pointer (*os.File) will facilitate writing
//		output data to destination files residing on an
//		attached storage drive. However, with this
//		configuration, the user is responsible for
//		manually closing the file and performing any
//		other required clean-up operations in addition to
//		calling local method:
//
//		FileIoReadWrite.CloseRelease()
//
//		While the 'write' services provided by the
//		FileIoReadWrite are primarily designed for
//		writing data to disk files, this type of 'writer'
//		will in fact write data to any object
//		implementing the io.Writer interface.
//
//	ioWriterLabel					string
//
//		The name or label associated with input parameter
//		'ioWriter' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "ioWriter" will be
//		automatically applied.
//
//	defaultWriterByteArraySize		int
//
//		The size of the byte array which will be used to
//		write data to the internal io.Writer object
//		encapsulated by FileIoReadWrite input parameter
//		'fIoReadWrite'.
//
//		If the value of 'defaultWriterByteArraySize' is
//		less than one ('1'), it will be reset to a size
//		of '4096'.
//
//		Although the FileIoReadWrite type does not use
//		the 'buffered' write protocol, the size of the
//		byte array used to write bytes to the underlying
//		io.Writer object is variable.
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
func (fIoReadWriteAtom *fileIoReadWriteAtom) setIoWriter(
	fIoReadWrite *FileIoReadWrite,
	fIoReadWriteLabel string,
	ioWriter io.Writer,
	ioWriterLabel string,
	defaultWriterByteArraySize int,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fIoReadWriteAtom.lock == nil {
		fIoReadWriteAtom.lock = new(sync.Mutex)
	}

	fIoReadWriteAtom.lock.Lock()

	defer fIoReadWriteAtom.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileIoReadWriteAtom." +
		"setIoWriter()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return err
	}

	if len(fIoReadWriteLabel) == 0 {

		fIoReadWriteLabel = "fIoReadWrite"
	}

	if len(ioWriterLabel) == 0 {

		ioWriterLabel = "ioWriter"
	}

	if fIoReadWrite == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is a nil pointer!\n"+
			"%v is invalid.\n",
			ePrefix.String(),
			fIoReadWriteLabel,
			fIoReadWriteLabel)

		return err
	}

	if ioWriter == nil {

		err = fmt.Errorf("%v\n"+
			"Error: The io.Writer instance passed\n"+
			"as input parameter '%v' is invalid!\n"+
			"'%v' is a 'nil' pointer.\n",
			ePrefix.String(),
			ioWriterLabel,
			ioWriterLabel)

		return err
	}

	// Close and release the old FileIoReadWrite.writer
	err = new(fileIoReadWriteElectron).
		writerCloseRelease(
			fIoReadWrite,
			fIoReadWriteLabel,
			true, // releaseWriterMemResources
			true, // releaseFBuffWriterLocalMemRes
			ePrefix)

	if err != nil {

		return err
	}

	var err2 error
	var newFileIoWriter FileIoWriter

	err2 = new(fileIoWriterNanobot).
		setIoWriter(
			&newFileIoWriter,
			ioWriterLabel+".writer",
			ioWriter,
			ioWriterLabel,
			defaultWriterByteArraySize,
			ePrefix)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"An error occurred while creating the new %v.writer.\n"+
			"Error=\n%v\n",
			funcName,
			fIoReadWriteLabel,
			err2.Error())

		return err
	}

	fIoReadWrite.writer = &newFileIoWriter

	fIoReadWrite.readerFilePathName =
		newFileIoWriter.targetWriteFileName

	return err
}

// setPathFileNameReader
//
// Receives an input parameter string specifying the path
// and file name identifying the file which will be
// configured as a data source for 'read' operations.
// This file will be configured as an internal io.Reader
// object for the FileIoReadWrite instance passed as
// input parameter 'fIoReadWrite'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reconfigure
//	the member variable io.Reader object encapsulated in
//	the instance of FileIoReadWrite passed as input
//	parameter 'fIoReadWrite':
//
//			fIoReadWrite.reader
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fIoReadWrite					*FileIoReadWrite
//
//		A pointer to an instance of FileIoReadWrite.
//
//		The internal io.Reader object encapsulated in
//		this instance of FileIoReadWrite will be deleted
//		and configured using the file identified by input
//		parameter 'readerPathFileName'.
//
//	fIoReadWriteLabel				string
//
//		The name or label associated with input parameter
//		'fIoReadWrite' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fIoReadWrite" will
//		be automatically applied.
//
//	readerPathFileName				string
//
//		This string contains the path and file name of
//		the file which will be configured as an io.Reader
//		object encapsulated in the FileIoReadWrite
//		instance passed as input parameter
//		'fIoReadWrite'. As such, the file identified by
//		'readerPathFileName' will be used a data source
//		for 'read' operations.
//
//		If this file does not currently exist on an
//		attached storage drive, an error will be
//		returned.
//
//	readerPathFileNameLabel			string
//
//		The name or label associated with input parameter
//		'readerPathFileName' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "readerPathFileName"
//		will be automatically applied.
//
//	openReadFileReadWrite			bool
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
//	defaultReaderByteArraySize		int
//
//		The size of the byte array which will be used to
//		read data from the internal io.Reader object
//		encapsulated by the FileIoReadWrite instance
//		passed as input parameter 'fIoReadWrite'.
//
//		If the value of 'defaultReaderByteArraySize' is
//		less than '16', it will be reset to a size of
//		'4096'.
//
//		Although the FileIoReadWrite type does not use
//		the 'buffered' read protocol, the size of the
//		byte array used to store bytes read from the
//		underlying io.Reader object is variable.
//
//	errPrefDto						*ePref.ErrPrefixDto
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
//	fileInfoPlus					FileInfoPlus
//
//		This returned instance of Type FileInfoPlus
//		contains data elements describing the file
//		identified by input parameter
//		'readerPathFileName'.
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
//	err								error
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
func (fIoReadWriteAtom *fileIoReadWriteAtom) setPathFileNameReader(
	fIoReadWrite *FileIoReadWrite,
	fIoReadWriteLabel string,
	readerPathFileName string,
	readerPathFileNameLabel string,
	openReadFileReadWrite bool,
	defaultReaderByteArraySize int,
	errPrefDto *ePref.ErrPrefixDto) (
	fInfoPlus FileInfoPlus,
	err error) {

	if fIoReadWriteAtom.lock == nil {
		fIoReadWriteAtom.lock = new(sync.Mutex)
	}

	fIoReadWriteAtom.lock.Lock()

	defer fIoReadWriteAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileIoReadWriteAtom." +
		"setPathFileNameReader()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		return fInfoPlus, err
	}

	if len(fIoReadWriteLabel) == 0 {

		fIoReadWriteLabel = "fIoReadWrite"
	}

	if len(readerPathFileNameLabel) == 0 {

		readerPathFileNameLabel = "readerPathFileName"
	}

	if fIoReadWrite == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is a nil pointer!\n"+
			"%v is invalid.\n",
			ePrefix.String(),
			fIoReadWriteLabel,
			fIoReadWriteLabel)

		return fInfoPlus, err
	}

	if len(readerPathFileName) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is invalid!\n"+
			"'%v' is an empty string with a length of zero (0).\n",
			ePrefix.String(),
			readerPathFileNameLabel,
			readerPathFileNameLabel)

		return fInfoPlus, err
	}

	err = new(fileIoReadWriteElectron).readerCloseRelease(
		fIoReadWrite,
		fIoReadWriteLabel,
		true, // releaseMemoryResources
		true, // releaseFBuffReaderLocalMemRes
		ePrefix.XCpy("Close-Reader"))

	if err != nil {

		return fInfoPlus, err
	}

	var newIoReader FileIoReader
	var err2 error

	fInfoPlus,
		err2 = new(fileIoReaderNanobot).
		setPathFileName(
			&newIoReader,
			fIoReadWriteLabel+".newIoReader",
			readerPathFileName,
			readerPathFileNameLabel,
			openReadFileReadWrite,
			defaultReaderByteArraySize,
			ePrefix)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"An error occurred while creating the new %v.reader.\n"+
			"Error=\n%v\n",
			funcName,
			fIoReadWriteLabel,
			err2.Error())

		return fInfoPlus, err
	}

	fIoReadWrite.reader = &newIoReader
	fIoReadWrite.readerFilePathName = readerPathFileName

	return fInfoPlus, err
}

// setPathFileNameWriter
//
// Receives an input parameter string specifying the path
// and file name identifying the file which will be
// configured as an output destination for 'write'
// operations. This file will be configured as an
// internal io.Writer object for the FileIoReadWrite
// instance passed as input parameter 'fIoReadWrite'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reconfigure
//	the member variable io.Writer object encapsulated in
//	the instance of FileIoReadWrite passed as input
//	parameter 'fIoReadWrite':
//
//			fIoReadWrite.writer
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fIoReadWrite					*FileIoReadWrite
//
//		A pointer to an instance of FileBufferWriter.
//
//		The internal io.Writer object encapsulated in
//		this instance of FileIoReadWrite will be
//		deleted and configured using the file identified
//		by input parameter 'writerPathFileName'.
//
//	fIoReadWriteLabel				string
//
//		The name or label associated with input parameter
//		'fIoReadWrite' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fIoReadWrite" will
//		be automatically applied.
//
//	writerPathFileName				string
//
//		This string contains the path and file name of
//		the target 'write' file which will be used as
//		a data destination for 'write' operations.
//
//		If the target path and file do not currently
//		exist on an attached storage drive, this method
//		will attempt to create them.
//
//	writerPathFileNameLabel			string
//
//		The name or label associated with input parameter
//		'writerPathFileName' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "writerPathFileName"
//		will be automatically applied.
//
//	openWriteFileReadWrite			bool
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
//	defaultWriterByteArraySize		int
//
//		The size of the default byte array which will be
//		used to write data to the internal io.Writer
//		object encapsulated by FileIoWriter input
//		parameter 'fIoWriter'.
//
//		If the value of 'defaultByteArraySize' is
//		less than one ('1'), it will be reset to a size
//		of '4096'.
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
//	fileInfoPlus				FileInfoPlus
//
//		This returned instance of Type FileInfoPlus
//		contains data elements describing the file
//		identified by input parameter
//		'writerPathFileName'.
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
func (fIoReadWriteAtom *fileIoReadWriteAtom) setPathFileNameWriter(
	fIoReadWrite *FileIoReadWrite,
	fIoReadWriteLabel string,
	writerPathFileName string,
	writerPathFileNameLabel string,
	openWriteFileReadWrite bool,
	defaultWriterByteArraySize int,
	truncateExistingWriteFile bool,
	errPrefDto *ePref.ErrPrefixDto) (
	fInfoPlus FileInfoPlus,
	err error) {

	if fIoReadWriteAtom.lock == nil {
		fIoReadWriteAtom.lock = new(sync.Mutex)
	}

	fIoReadWriteAtom.lock.Lock()

	defer fIoReadWriteAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileIoReadWriteAtom." +
		"setPathFileNameWriter()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		return fInfoPlus, err
	}

	if len(fIoReadWriteLabel) == 0 {

		fIoReadWriteLabel = "fIoReadWrite"
	}

	if len(writerPathFileNameLabel) == 0 {

		writerPathFileNameLabel = "writerPathFileName"
	}

	if fIoReadWrite == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is a nil pointer!\n"+
			"%v is invalid.\n",
			ePrefix.String(),
			fIoReadWriteLabel,
			fIoReadWriteLabel)

		return fInfoPlus, err
	}

	if len(writerPathFileName) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is invalid!\n"+
			"'%v' is an empty string with a length of zero (0).\n",
			ePrefix.String(),
			writerPathFileNameLabel,
			writerPathFileNameLabel)

		return fInfoPlus, err
	}

	err = new(fileIoReadWriteElectron).writerCloseRelease(
		fIoReadWrite,
		fIoReadWriteLabel,
		true, // releaseMemoryResources
		true, // releaseFBuffReaderLocalMemRes
		ePrefix.XCpy("Close-Writer"))

	if err != nil {

		return fInfoPlus, err
	}

	var newIoWriter FileIoWriter
	var err2 error

	fInfoPlus,
		err2 = new(fileIoWriterNanobot).
		setPathFileName(
			&newIoWriter,
			fIoReadWriteLabel+".newIoWriter",
			writerPathFileName,
			writerPathFileNameLabel,
			openWriteFileReadWrite,
			defaultWriterByteArraySize,
			truncateExistingWriteFile,
			ePrefix)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"An error occurred while creating the new %v.writer.\n"+
			"Error=\n%v\n",
			funcName,
			fIoReadWriteLabel,
			err2.Error())

		return fInfoPlus, err
	}

	fIoReadWrite.writer = &newIoWriter
	fIoReadWrite.readerFilePathName = writerPathFileName

	return fInfoPlus, err
}

type fileIoReadWriteElectron struct {
	lock *sync.Mutex
}

// empty
//
// This method deletes all internal member variables and
// release all the internal memory resources for an
// instance of FileIoReadWrite passed as input parameter
// 'fBufReadWrite'.
//
// Specifically the following internal member variables
// are set to 'nil' or their initial zero values:
//
//	FileIoReadWrite.reader = nil
//	FileIoReadWrite.writer = nil
//	FileIoReadWrite.readerFilePathName = ""
//	FileIoReadWrite.writerFilePathName = ""
//
// After calling this method, the instance of
// FileIoReadWrite, passed as input parameter
// 'fIoReadWrite', will become invalid and unavailable
// for future read/write operations.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	This method does NOT perform the 'close' procedure.
//	To perform the 'close' procedures while
//	simultaneously releasing all internal memory
//	resources, call local method:
//
//			FileBufferReadWrite.Close()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fIoReadWrite				*FileIoReadWrite
//
//		A pointer to an instance of FileIoReadWrite.
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
func (fIoReadWriteElectron *fileIoReadWriteElectron) empty(
	fIoReadWrite *FileIoReadWrite) {

	if fIoReadWriteElectron.lock == nil {
		fIoReadWriteElectron.lock = new(sync.Mutex)
	}

	fIoReadWriteElectron.lock.Lock()

	defer fIoReadWriteElectron.lock.Unlock()

	if fIoReadWrite == nil {
		return
	}

	fIoReadWrite.writer = nil
	fIoReadWrite.reader = nil
	fIoReadWrite.writerFilePathName = ""
	fIoReadWrite.readerFilePathName = ""

	return
}

// isFileIoReadWriteValid
//
// This method receives a pointer to an instance of
// FileIoReadWrite ('fIoReadWrite') which will be
// analyzed to determine if all the member variables
// contain valid data values.
//
// If input parameter 'fIoReadWrite' is determined to be
// invalid, this method returns an error.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fIoReadWrite				*FileIoReadWrite
//
//		A pointer to an instance of FileIoReadWrite.
//
//		If any of the internal member variable data
//		values encapsulated in 'fIoReadWrite' are
//		determined to be invalid, this method will return
//		an error.
//
//	fIoReadWriteLabel			string
//
//		The name or label associated with input parameter
//		'fIoReadWrite' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fIoReadWrite" will
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
//		contained in the instance of FileIoReadWrite
//		passed as 'fIoReadWrite' are found to be
//		invalid, this method will return an error
//		configured with an appropriate message
//		identifying the invalid	member data variable.
//
//		If all internal member data variables evaluate
//		as valid, this returned error value will be set
//		to 'nil'.
//
//		If errors are encountered during processing or if
//		any 'fIoReadWrite' internal member data values
//		are found to be invalid, the returned error Type
//		will encapsulate an appropriate error message.
//	 	This returned error message will incorporate the
//	 	method chain and text passed by input parameter,
//	 	'errorPrefix'. The 'errorPrefix' text will be
//	 	prefixed to the beginning of the error message.
func (fIoReadWriteElectron *fileIoReadWriteElectron) isFileIoReadWriteValid(
	fIoReadWrite *FileIoReadWrite,
	fIoReadWriteLabel string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fIoReadWriteElectron.lock == nil {
		fIoReadWriteElectron.lock = new(sync.Mutex)
	}

	fIoReadWriteElectron.lock.Lock()

	defer fIoReadWriteElectron.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileIoReadWriteElectron." +
		"isFileIoReadWriteValid()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return err
	}

	if len(fIoReadWriteLabel) == 0 {

		fIoReadWriteLabel = "fIoReadWrite"
	}

	if fIoReadWrite == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is a nil pointer!\n"+
			"%v is invalid.\n",
			ePrefix.String(),
			fIoReadWriteLabel,
			fIoReadWriteLabel)

		return err
	}

	if fIoReadWrite.reader == nil &&
		fIoReadWrite.writer == nil {

		err = fmt.Errorf("%v\n"+
			" -----------------------------------------------------------\n"+
			" ERROR: The %v instance of FileIoReadWrite\n"+
			" is invalid! The internal io.Reader and io.Writer objects\n"+
			" were never initialized. Call one of the 'New' methods or\n"+
			" 'Setter' methods to create a valid instance of\n"+
			" FileBufferReadWrite.\n",
			ePrefix.String(),
			fIoReadWriteLabel)

		return err
	}

	if fIoReadWrite.reader == nil {

		err = fmt.Errorf(" %v\n"+
			" -------------------------------------------------------------------\n"+
			" ERROR: The %v instance of FileIoReadWrite is invalid!\n"+
			" The internal io.Reader object was never initialized.\n"+
			" Call one of the 'New' methods or 'Setter' methods to create\n"+
			" a valid instance of FileIoReadWrite.\n",
			ePrefix.String(),
			fIoReadWriteLabel)

		return err
	}

	if fIoReadWrite.writer == nil {

		err = fmt.Errorf(" %v\n"+
			" -------------------------------------------------------------------\n"+
			" ERROR: The %v instance of FileIoReadWrite is invalid!\n"+
			" The internal io.Writer object was never initialized.\n"+
			" Call one of the 'New' methods or 'Setter' methods to create\n"+
			" a valid instance of FileIoReadWrite.\n",
			ePrefix.String(),
			fIoReadWriteLabel)

		return err
	}

	var err2 error

	err2 = new(fileIoReaderAtom).isFileIoReaderValid(
		fIoReadWrite.reader,
		"fIoReadWrite.reader",
		ePrefix)

	if err2 != nil {

		err = fmt.Errorf(" %v\n"+
			" -------------------------------------------------------------------\n"+
			" ERROR: The %v instance of FileIoReadWrite is invalid!\n"+
			" The internal io.reader returned a validation error.\n"+
			" Error:\n"+
			"%v\n",
			funcName,
			fIoReadWriteLabel,
			err2.Error())

		return err
	}

	err2 = new(fileIoWriterAtom).isFileIoWriterValid(
		fIoReadWrite.writer,
		"fIoReadWrite.writer",
		ePrefix)

	if err2 != nil {

		err = fmt.Errorf(" %v\n"+
			" -------------------------------------------------------------------\n"+
			" ERROR: The %v instance of FileIoReadWrite is invalid!\n"+
			" The internal io.Writer returned a validation error.\n"+
			" Error:\n"+
			"%v\n",
			funcName,
			fIoReadWriteLabel,
			err2.Error())

	}

	return err
}

// readerCloseRelease
//
// This method will perform Clean-Up operations on the
// internal io.Reader object encapsulated in the
// FileIoReadWrite instance passed as input parameter
// 'fIoReadWrite':
//
//	fIoReadWrite.reader
//
// Upon completion, this method will effectively render
// the 'fIoReadWrite' instance invalid and unusable for
// any future 'read' operations.
//
// This Clean-Up operation is accomplished by and closing
// the internal io.Reader object before finally releasing
// the memory resources contained in that object:
//
//	fIoReadWrite.reader
//
// The 'release memory resources' actions are implemented
// independently based on the values passed for input
// parameters 'releaseReaderMemResources', and
// 'releaseFBuffReaderLocalMemRes'.
//
// 'releaseReaderMemResources' controls the release of
// memory resources associated with the internal io.Reader
// object.
//
// 'releaseFBuffReaderLocalMemRes' controls the release of
// FileIoReadWrite memory resources associated with the
// io.Reader object.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fIoReadWrite						*FileIoReadWrite
//
//		A pointer to an instance of FileBufferReader.
//
//		This method will perform Clean-Up operations on
//		the internal io.Reader object encapsulated in
//		this FileIoReadWrite instance:
//
//			fIoReadWrite.reader
//
//		Upon completion, this method will effectively
//		render the FileIoReadWrite instance,
//		'fIoReadWrite', invalid and unusable for any
//		future 'read' operations.
//
//	fIoReadWriteLabel					string
//
//		The name or label associated with input parameter
//		'fIoReadWrite' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fIoReadWrite" will
//		be automatically applied.
//
//	releaseReaderMemResources			bool
//
//		If 'releaseReaderMemResources' is set to 'true',
//		this method	will release the memory resources
//		contained in the internal io.reader object
//		encapsulated by	'fIoReadWrite':
//
//			fIoReadWrite.reader
//
//		Releasing internal memory resources synchronizes
//		internal flags and prevents multiple calls to the
//		'close' method. Calling the 'close' method more
//		than once may produce unexpected results.
//
//	releaseFIoReaderLocalMemRes			bool
//
//		If 'releaseFIoReaderLocalMemRes' is set to
//		'true', this method will release the local memory
//		resources for the FileIoReadWrite reader object
//		(fIoReadWrite):
//
//			fIoReadWrite.reader = nil
//			fIoReadWrite.readerFilePathName = ""
//
//		Releasing internal memory resources synchronizes
//		internal flags and prevents multiple calls to the
//		'close' method. Calling the 'close' method more
//		than once may produce unexpected results.
//
//	errPrefDto							*ePref.ErrPrefixDto
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
func (fIoReadWriteElectron *fileIoReadWriteElectron) readerCloseRelease(
	fIoReadWrite *FileIoReadWrite,
	fBufReadWriteLabel string,
	releaseReaderMemResources bool,
	releaseFIoReaderLocalMemRes bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fIoReadWriteElectron.lock == nil {
		fIoReadWriteElectron.lock = new(sync.Mutex)
	}

	fIoReadWriteElectron.lock.Lock()

	defer fIoReadWriteElectron.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileIoReadWriteElectron." +
		"readerCloseRelease()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return err
	}

	if len(fBufReadWriteLabel) == 0 {

		fBufReadWriteLabel = "fIoReadWrite"
	}

	if fIoReadWrite == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is a nil pointer!\n"+
			"%v is invalid.\n",
			ePrefix.String(),
			fBufReadWriteLabel,
			fBufReadWriteLabel)

		return err
	}

	var err2 error

	if fIoReadWrite.reader != nil {

		err2 = new(fileIoReaderMolecule).
			closeAndRelease(
				fIoReadWrite.reader,
				fBufReadWriteLabel+".reader",
				releaseReaderMemResources,
				ePrefix)

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Error occurred while closing the %v.\n"+
				"Error:\n%v\n",
				funcName,
				fBufReadWriteLabel+".reader",
				err2.Error())

			return err
		}
	}

	if releaseFIoReaderLocalMemRes == true {

		fIoReadWrite.reader = nil
		fIoReadWrite.readerFilePathName = ""

	}

	return err
}

// writerCloseRelease
//
// This method will perform Clean-Up operations on the
// internal io.Writer object encapsulated in the
// FileIoReadWrite instance passed as input parameter
// 'fIoReadWrite':
//
//	fIoReadWrite.writer
//
// Upon completion, this method will effectively render
// the 'fIoReadWrite' instance invalid and unusable for
// any future 'write' operations.
//
// This Clean-Up operation is accomplished by and closing
// the internal io.Writer object before finally releasing
// the memory resources contained in that object:
//
//	fIoReadWrite.writer
//
// The 'release memory resources' actions are implemented
// independently based on the values passed for input
// parameters 'releaseWriterMemResources', and
// 'releaseFBuffWriterLocalMemRes'.
//
// 'releaseWriterMemResources' controls the release of
// memory resources associated with the internal io.Writer
// object.
//
// 'releaseFBuffWriterLocalMemRes' controls the release of
// FileIoReadWrite memory resources associated with the
// io.Writer object.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fIoReadWrite						*FileIoReadWrite
//
//		A pointer to an instance of FileBufferWriter.
//
//		This method will perform Clean-Up operations on
//		the internal io.Writer object encapsulated in
//		this FileIoReadWrite instance:
//
//			fIoReadWrite.writer
//
//		Upon completion, this method will effectively
//		render the FileIoReadWrite instance,
//		'fIoReadWrite', invalid and unusable for any
//		future 'write' operations.
//
//	fIoReadWriteLabel					string
//
//		The name or label associated with input parameter
//		'fIoReadWrite' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fIoReadWrite" will
//		be automatically applied.
//
//	releaseWriterMemResources			bool
//
//		If 'releaseWriterMemResources' is set to 'true',
//		this method	will release the memory resources
//		contained in the internal io.writer object
//		encapsulated by	'fIoReadWrite':
//
//			fIoReadWrite.writer
//
//		Releasing internal memory resources synchronizes
//		internal flags and prevents multiple calls to the
//		'close' method. Calling the 'close' method more
//		than once may produce unexpected results.
//
//	releaseFIoWriterLocalMemRes		bool
//
//		If 'releaseFIoWriterLocalMemRes' is set to
//		'true', this method will release the local memory
//		resources for the FileIoReadWrite writer object
//		(fIoReadWrite):
//
//			fIoReadWrite.writer = nil
//			fIoReadWrite.writerFilePathName = ""
//
//		Releasing internal memory resources synchronizes
//		internal flags and prevents multiple calls to the
//		'close' method. Calling the 'close' method more
//		than once may produce unexpected results.
//
//	errPrefDto							*ePref.ErrPrefixDto
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
func (fIoReadWriteElectron *fileIoReadWriteElectron) writerCloseRelease(
	fIoReadWrite *FileIoReadWrite,
	fBufReadWriteLabel string,
	releaseWriterMemResources bool,
	releaseFBuffWriterLocalMemRes bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fIoReadWriteElectron.lock == nil {
		fIoReadWriteElectron.lock = new(sync.Mutex)
	}

	fIoReadWriteElectron.lock.Lock()

	defer fIoReadWriteElectron.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileIoReadWriteElectron." +
		"writerCloseRelease()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return err
	}

	if len(fBufReadWriteLabel) == 0 {

		fBufReadWriteLabel = "fIoReadWrite"
	}

	if fIoReadWrite == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is a nil pointer!\n"+
			"%v is invalid.\n",
			ePrefix.String(),
			fBufReadWriteLabel,
			fBufReadWriteLabel)

		return err
	}

	var err2 error

	if fIoReadWrite.writer != nil {

		err2 = new(fileIoWriterMolecule).
			closeAndRelease(
				fIoReadWrite.writer,
				fBufReadWriteLabel+".writer",
				releaseWriterMemResources,
				ePrefix)

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Error occurred while closing the %v.\n"+
				"Error:\n%v\n",
				funcName,
				fBufReadWriteLabel+".writer",
				err2.Error())

			return err
		}

	}

	if releaseFBuffWriterLocalMemRes == true {

		fIoReadWrite.writer = nil
		fIoReadWrite.writerFilePathName = ""

	}

	return err
}
