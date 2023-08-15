package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"os"
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
// buffered read techniques. Direct data reads are
// performed by means of the internal io.Reader object.
//
// For more information on 'buffered' data reads, see
// type 'FileBufferReader'
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://pkg.go.dev/io#Reader
type FileIoReader struct {
	fileIoReader       *io.Reader
	filePtr            *os.File
	targetReadFileName string

	lock *sync.Mutex
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
//	buffered read techniques.
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

	fIoReader.fileIoReader = &reader

	return err
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

	fIoReader.fileIoReader = nil

	return err
}
