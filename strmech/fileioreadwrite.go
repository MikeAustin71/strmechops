package strmech

import (
	"errors"
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"io"
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
// This method returns an empty or 'blank' instance of
// FileIoReadWrite. All the member variables in this
// returned instance are initialized to their zero or
// initial values. This means the returned instance is
// invalid and unusable for standard 'read' and 'write'
// operations.
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

/*
func (fIoReadWrite *FileIoReadWrite)  NewFileMgrs(
	readerFileMgr *FileMgr,
	openReadFileReadWrite bool,
	readerBuffSize int,
	writerFileMgr *FileMgr,
	openWriteFileReadWrite bool,
	writerBuffSize int,
	truncateExistingWriteFile bool,
	errorPrefix interface{}) (
	readerFileInfoPlus FileInfoPlus,
	writerFileInfoPlus FileInfoPlus,
	newFBuffReadWrite *FileIoReadWrite,
	err error) {

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

}

*/

// NewIoReadWrite
//
// Creates and returns a new, fully configured instance
// of FileIoReadWrite using io.Reader and io.Writer
// input parameters.
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
//		FileIoReadWrite.CloseAndRelease()
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
//		encapsulated by the returned FileIoReadWrite
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
//	FileIoReadWrite
//
//		If this method completes successfully, it will
//		return a fully configured instance of
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
func (fIoReadWrite *FileIoReadWrite) NewIoReadWrite(
	reader io.Reader,
	defaultReaderByteArraySize int,
	writer io.Writer,
	defaultWriterByteArraySize int,
	errorPrefix interface{}) (
	*FileIoReadWrite,
	error) {

	if fIoReadWrite.lock == nil {
		fIoReadWrite.lock = new(sync.Mutex)
	}

	fIoReadWrite.lock.Lock()

	defer fIoReadWrite.lock.Unlock()

	var newFIoReadWrite = new(FileIoReadWrite)

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileIoReadWrite."+
			"NewIoReadWrite()",
		"")

	if err != nil {
		return newFIoReadWrite, err
	}

	err = new(fileIoReadWriteMolecule).
		setIoReaderIoWriter(
			newFIoReadWrite,
			"newFIoReadWrite",
			reader,
			"reader",
			defaultReaderByteArraySize,
			writer,
			"writer",
			defaultWriterByteArraySize,
			ePrefix)

	return newFIoReadWrite, err
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
//		The size of the byte array which will be used to
//		write data to the internal io.Writer object
//		encapsulated by FileIoWriter input parameter
//		'fIoWriter'.
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
