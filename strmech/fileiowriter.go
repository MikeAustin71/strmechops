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
//	(1)	Input parameter 'writer' will accept a pointer to an
//		instance of os.File because os.File implements the
//		io.Writer interface.
//
//	(2) The returned instance of FileIoWriter will retain a
//		pointer reference to input parameter 'writer'.
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
