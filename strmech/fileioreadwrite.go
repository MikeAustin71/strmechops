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

// writerFlushCloseRelease
//
// This method will perform Clean-Up operations on the
// internal io.Writer object encapsulated in the
// FileIoReadWrite instance passed as input parameter
// 'fIoReadWrite':
//
//	fIoReadWrite.writer
//
// Upon completion, this method will effectively render the
// 'fIoReadWrite' instance invalid and unusable for any
// future 'write' operations.
//
// This operation is accomplished by flushing and closing
// the internal bufio.Writer object before finally
// releasing the memory resources associated with that
// object:
//
//	fIoReadWrite.writer
//
// The 'flush' and 'release memory resources' actions are
// implemented independently based on the values passed
// for input parameters 'flushWriteBuffer',
// 'releaseWriterMemResources', and
// 'releaseFBuffWriterLocalMemRes'.
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
//		the internal bufio.Writer object encapsulated in
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
//		this method	will release the internal memory
//		resources for the internal io.writer object
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
		"writerFlushCloseRelease()"

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
