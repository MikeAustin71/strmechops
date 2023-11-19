package strmech

import (
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
	FileIoReadWrite,
	error) {

	if fIoReadWrite.lock == nil {
		fIoReadWrite.lock = new(sync.Mutex)
	}

	fIoReadWrite.lock.Lock()

	defer fIoReadWrite.lock.Unlock()

	var newFIoReadWrite = FileIoReadWrite{}

	var ePrefix *ePref.ErrPrefixDto
	var err error
	funcName := "FileIoReadWrite." +
		"NewIoReadWrite()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {
		return newFIoReadWrite, err
	}

	var newReader FileIoReader

	newReader,
		err = new(FileIoReader).
		NewIoReader(
			reader,
			defaultReaderByteArraySize,
			ePrefix.XCpy("newReader"))

	if err != nil {
		return newFIoReadWrite, err
	}

	var newWriter FileIoWriter

	newWriter,
		err = new(FileIoWriter).
		NewIoWriter(
			writer,
			defaultWriterByteArraySize,
			ePrefix.XCpy("newWriter"))

	if err != nil {
		return newFIoReadWrite, err
	}

	newFIoReadWrite.reader = &newReader
	newFIoReadWrite.readerFilePathName =
		newReader.targetReadFileName

	newFIoReadWrite.writer = &newWriter
	newFIoReadWrite.writerFilePathName =
		newWriter.targetWriteFileName

	return newFIoReadWrite, err
}

type fileIoReadWriteElectron struct {
	lock *sync.Mutex
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
//	releaseFIoReaderLocalMemRes		bool
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

	if releaseFBuffWriterLocalMemRes == true {

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
