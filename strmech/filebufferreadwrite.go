package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"sync"
)

// FileBufferReadWrite
//
// This type implements both the io.Reader and
// io.Writer interfaces. It is designed to provide a
// read/write capability for files and any objects
// supporting the io.Reader/io.Writer interfaces.
//
// User can utilize this type to perform read and
// write operations in a single method call.
//
// Read and Write operations are performed using
// private, internal FileBufferReader and
// FileBufferWriter objects.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	To create a valid instance of FileBufferReadWrite,
//	users MUST call one of the 'New' methods.
type FileBufferReadWrite struct {
	writer            *FileBufferWriter
	reader            *FileBufferReader
	writeFilePathName string
	readFilePathName  string

	lock *sync.Mutex
}

// New
//
// Creates and returns a new, fully configured instance
// of FileBufferReadWrite
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
//		in addition to calling FileBufferReadWrite.Close().
//
//		While the 'read' services provided by
//		FileBufferReadWrite are primarily designed to
//		read data from disk files, this type of 'reader'
//		will in fact read data from any object
//		implementing the io.Reader interface.
//
//	readerBuffSize				int
//
//		This integer value controls the size of the
//		buffer created for the io.Reader object passed as
//		input parameter 'reader'.
//
//		'readerBuffSize' should be configured to maximize
//		performance for 'read' operations subject to
//		prevailing memory limitations.
//
//		The minimum reader buffer size is 16-bytes. If
//		'readerBuffSize' is set to a value less than
//		"16", it will be automatically reset to the
//		default buffer size of 4096-bytes.
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
//		FileBufferReadWrite.Close().
//
//		While the 'write' services provided by the
//		FileBufferReadWrite are primarily designed for
//		writing data to disk files, this type of 'writer'
//		will in fact write data to any object
//		implementing the io.Writer interface.
//
//	writerBuffSize				int
//
//		This integer value controls the size of the
//		buffer created for the io.Writer object passed as
//		input parameter 'writer'.
//
//		'writerBuffSize' should be configured to maximize
//		performance for 'write' operations subject to
//		prevailing memory limitations.
//
//		The minimum write buffer size is 1-byte. If
//		'writerBuffSize' is set to a size less than or
//		equal to zero, it will be automatically reset to
//		the default buffer size of 4096-bytes.
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
//	FileBufferReadWrite
//
//		If this method completes successfully, it will
//		return a fully configured instance of
//		FileBufferReadWrite.
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
func (fBufReadWrite *FileBufferReadWrite) New(
	reader io.Reader,
	readerBuffSize int,
	writer io.Writer,
	writerBuffSize int,
	errorPrefix interface{}) (
	FileBufferReadWrite,
	error) {

	if fBufReadWrite.lock == nil {
		fBufReadWrite.lock = new(sync.Mutex)
	}

	fBufReadWrite.lock.Lock()

	defer fBufReadWrite.lock.Unlock()

	var newFBuffReadWrite = FileBufferReadWrite{}

	var ePrefix *ePref.ErrPrefixDto
	var err error
	funcName := "FileBufferReadWrite." +
		"New()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {
		return newFBuffReadWrite, err
	}

	err = new(fileBufferReadWriteNanobot).
		setIoReaderWriter(
			&newFBuffReadWrite,
			"newFBuffReadWrite",
			reader,
			"reader",
			readerBuffSize,
			writer,
			"writer",
			writerBuffSize,
			ePrefix)

	return newFBuffReadWrite, err
}

func (fBufReadWrite *FileBufferReadWrite) NewFileMgrs(
	readerFileMgr *FileMgr,
	openReadFileReadWrite bool,
	readerBuffSize int,
	writerFileMgr *FileMgr,
	openWriteFileReadWrite bool,
	writerBuffSize int,
	truncateExistingWriteFile bool,
	errorPrefix interface{}) (
	FileBufferReadWrite,
	error) {

	if fBufReadWrite.lock == nil {
		fBufReadWrite.lock = new(sync.Mutex)
	}

	fBufReadWrite.lock.Lock()

	defer fBufReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	var newFBuffReadWrite = FileBufferReadWrite{}

	funcName := "FileBufferReadWrite." +
		"NewFileMgrs()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {

		return newFBuffReadWrite, err
	}

	var newFileBufReader FileBufferReader
	var err2 error

	err2 = new(fileBufferReaderMicrobot).
		setFileMgr(
			&newFileBufReader,
			"newFileBufReader",
			readerFileMgr,
			"readerFileMgr",
			openReadFileReadWrite,
			readerBuffSize,
			ePrefix.XCpy(
				"readerFileMgr"))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"An error occurred while configuring the reader\n"+
			"from input parameter 'readerFileMgr'.\n"+
			"Error=\n%v\n",
			funcName,
			err2.Error())

		return newFBuffReadWrite, err
	}

	return newFBuffReadWrite, err
}

// NewPathFileNames
//
// Opens a 'read' file and a 'write' file using path and
// file name strings passed as input parameters. The
// 'read' and 'write' files are then configured as
// io.Reader and io.Writer instances. Finally, this
// configuration is returned as a new and full populated
// instance of FileBufferReadWrite
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	readerPathFileName			string
//
//		This string contains the path and file name of
//		the file which will be used a data source for
//		'read' operations.
//
//		If this file does not currently exist on an
//		attached storage drive, an error will be
//		returned.
//
//	openReadFileReadWrite		bool
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
//	readerBuffSize					int
//
//		This integer value controls the size of the
//		'read' buffer created for reading data from
//		the file identified by 'readerPathFileName'.
//
//		'readerBuffSize' should be configured to maximize
//		performance for 'read' operations subject to
//		prevailing memory limitations.
//
//		The minimum reader buffer size is 16-bytes. If
//		'readerBuffSize' is set to a size less than "16",
//		it will be automatically set to the default buffer
//		size of 4096-bytes.
//
//	writerPathFileName			string
//
//		This string contains the path and file name of
//		the target 'write' file which will be used as
//		a data destination for 'write' operations.
//
//		If the target path and file do not currently
//		exist on an attached storage drive, this method
//		will attempt to create them.
//
//	openWriteFileReadWrite		bool
//
//		If this parameter is set to 'true', the target
//		'write' file identified by input parameter
//		'pathFileName' will be opened for 'read' and
//		'write' operations.
//
//		If 'openWriteFileReadWrite' is set to 'false',
//		the target write file will be opened for
//		'write-only' operations.
//
//	writerBuffSize				int
//
//		This integer value controls the size of the
//		'write' buffer created for writing data to the
//		file identified by 'writerPathFileName'.
//
//		'writerBuffSize' should be configured to maximize
//		performance for 'write' operations subject to
//		prevailing memory limitations.
//
//		If 'writerBuffSize' is set to a value less than
//		or equal to zero (0), it will be automatically
//		reset to the default value of 4096-bytes.
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
func (fBufReadWrite *FileBufferReadWrite) NewPathFileNames(
	readerPathFileName string,
	openReadFileReadWrite bool,
	readerBuffSize int,
	writerPathFileName string,
	openWriteFileReadWrite bool,
	writerBuffSize int,
	truncateExistingWriteFile bool,
	errorPrefix interface{}) (
	FileBufferReadWrite,
	error) {

	if fBufReadWrite.lock == nil {
		fBufReadWrite.lock = new(sync.Mutex)
	}

	fBufReadWrite.lock.Lock()

	defer fBufReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	var newFBuffReadWrite = FileBufferReadWrite{}

	funcName := "FileBufferReadWrite." +
		"NewPathFileNames()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {
		return newFBuffReadWrite, err
	}

	if len(readerPathFileName) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'readerPathFileName' is invalid!\n"+
			"'readerPathFileName' is an empty string with a length of zero (0).\n",
			ePrefix.String())

		return newFBuffReadWrite, err
	}

	if len(writerPathFileName) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'writerPathFileName' is invalid!\n"+
			"'readerPathFileName' is an empty string with a length of zero (0).\n",
			ePrefix.String())

		return newFBuffReadWrite, err
	}

	var buffReader FileBufferReader

	err = new(fileBufferReaderNanobot).
		setPathFileName(
			&buffReader,
			"buffReader",
			readerPathFileName,
			"readerPathFileName",
			openReadFileReadWrite, // openFileReadWrite
			readerBuffSize,
			ePrefix.XCpy("buffReader<-readerPathFileName"))

	if err != nil {

		return newFBuffReadWrite, err
	}

	newFBuffReadWrite.reader = &buffReader

	var buffWriter FileBufferWriter

	err = new(fileBufferWriterNanobot).
		setPathFileName(
			&buffWriter,
			"buffWriter",
			writerPathFileName,
			"writerPathFileName",
			openWriteFileReadWrite, // openFileReadWrite
			writerBuffSize,
			truncateExistingWriteFile, // truncateExistingFile
			ePrefix.XCpy("buffWriter<-writerPathFileName"))

	if err != nil {

		return newFBuffReadWrite, err
	}

	newFBuffReadWrite.writer = &buffWriter

	return newFBuffReadWrite, err
}

// ReadWriteAll
//
// This method reads all data from the 'reader' data
// source and writes all that data to the 'writer'
// destination.
//
// If the total number of bytes read does NOT equal the
// total number of bytes written, an error will be
// returned.
//
// The 'read' and 'write' operations use the io.Reader
// and io.Writer objects created when the current
// instance of FileBufferReadWrite was initialized with
// one of the 'New' methods.
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
func (fBufReadWrite *FileBufferReadWrite) ReadWriteAll(
	errorPrefix interface{}) (
	totalBytesRead int,
	totalBytesWritten int,
	err error) {

	if fBufReadWrite.lock == nil {
		fBufReadWrite.lock = new(sync.Mutex)
	}

	fBufReadWrite.lock.Lock()

	defer fBufReadWrite.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "FileBufferReadWrite." +
		"ReadWriteAll()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {

		return totalBytesRead, totalBytesWritten, err
	}

	if fBufReadWrite.reader == nil ||
		fBufReadWrite.writer == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: The current instance of\n"+
			"FileBufferReadWrite is invalid!\n"+
			"Call one of the 'New' methods to\n"+
			"create a valid instance of FileBufferReadWrite.\n",
			ePrefix.String())

		return totalBytesRead, totalBytesWritten, err
	}

	var readErr, writeErr error
	var numOfBytesRead, numOfBytesWritten, cycleCount int

	byteArray := make([]byte,
		fBufReadWrite.reader.fileReader.Size())

	var errs []error

	for {

		cycleCount++

		numOfBytesRead,
			readErr =
			fBufReadWrite.reader.Read(byteArray)

		if readErr != nil &&
			readErr != io.EOF {

			var err2 error

			err2 = fmt.Errorf("%v\n"+
				"Error Reading Target Read File!\n"+
				"Cycle Count= %v\n"+
				"Error= \n%v\n",
				ePrefix.String(),
				cycleCount,
				readErr.Error())

			errs = append(
				errs, err2)

			var err3 error

			err3 = fBufReadWrite.reader.Close(
				ePrefix.XCpy("reader-Close"))

			if err3 != nil {

				errs = append(
					errs, err3)

			}

			var err4 error

			err4 = fBufReadWrite.writer.Close(
				ePrefix.XCpy("writer-Close"))

			if err4 != nil {

				errs = append(
					errs, err4)

			}

			err = new(StrMech).ConsolidateErrors(errs)

			return totalBytesRead, totalBytesWritten, err
		}

		if numOfBytesRead > 0 {

			totalBytesRead += numOfBytesRead

			numOfBytesWritten,
				writeErr = fBufReadWrite.writer.Write(
				byteArray[0:numOfBytesRead])

			if writeErr != nil {

				var err2 error

				err2 = fmt.Errorf("%v\n"+
					"Error Writing Bytes To File!\n"+
					"Write Error=\n%v\n",
					ePrefix.String(),
					writeErr.Error())

				errs = append(
					errs, err2)

				var err3 error

				err3 = fBufReadWrite.reader.Close(
					ePrefix.XCpy("reader-Close"))

				if err3 != nil {

					errs = append(
						errs, err3)

				}

				var err4 error

				err4 = fBufReadWrite.writer.Close(
					ePrefix.XCpy("writer-Close"))

				if err4 != nil {

					errs = append(
						errs, err4)

				}

				err = new(StrMech).ConsolidateErrors(errs)

				return totalBytesRead, totalBytesWritten, err
			}

			totalBytesWritten += numOfBytesWritten
		}

		if numOfBytesRead != numOfBytesWritten {

			var err2 error

			err2 = fmt.Errorf("%v\n"+
				"Error Writing Bytes To File!\n"+
				"numOfBytesRead != numOfBytesWritten\n"+
				"numOfBytesRead = %v\n"+
				"numOfBytesWritten = %v\n",
				ePrefix.String(),
				numOfBytesRead,
				numOfBytesWritten)

			errs = append(
				errs, err2)

			var err3 error

			err3 = fBufReadWrite.reader.Close(
				ePrefix.XCpy("reader-Close"))

			if err3 != nil {

				errs = append(
					errs, err3)

			}

			var err4 error

			err4 = fBufReadWrite.writer.Close(
				ePrefix.XCpy("writer-Close"))

			if err4 != nil {

				errs = append(
					errs, err4)

			}

			err = new(StrMech).ConsolidateErrors(errs)

			return totalBytesRead, totalBytesWritten, err
		}

		if readErr == io.EOF {

			break
		}

	}

	return totalBytesRead, totalBytesWritten, err
}

type fileBufferReadWriteNanobot struct {
	lock *sync.Mutex
}

// setIoReaderWriter
//
// This 'setter' method is used to initialize new values
// for internal member variables contained in the
// instance of FileBufferReadWrite passed as input
// parameter 'fBufReadWrite'. The new configuration will
// be based on the io.Reader and io.Writer object passed
// as input parameters.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the instance of
//	FileBufferReadWrite passed as input parameter
//	'fBufReadWrite'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fBufReadWrite				*FileBufferReadWrite
//
//		A pointer to an instance of FileBufferWriter.
//
//		The internal FileBufferReader and
//		FileBufferWriter objects encapsulated in this
//		instance be deleted and reinitialized using the
//		io.Reader and io.Writer objects passed as input
//		parameters 'reader' and 'writer'.
//
//	fBufReadWriteLabel			string
//
//		The name or label associated with input parameter
//		'fBufReadWrite' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fBufReadWrite" will
//		be automatically applied.
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
//		in addition to calling FileBufferReadWrite.Close().
//
//		While the 'read' services provided by
//		FileBufferReadWrite are primarily designed to
//		read data from disk files, this type of 'reader'
//		will in fact read data from any object
//		implementing the io.Reader interface.
//
//	readerLabel					string
//
//		The name or label associated with input parameter
//		'reader' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "reader" will
//		be automatically applied.
//
//	readerBuffSize				int
//
//		This integer value controls the size of the
//		buffer created for the io.Reader object passed as
//		input parameter 'reader'.
//
//		'readerBuffSize' should be configured to maximize
//		performance for 'read' operations subject to
//		prevailing memory limitations.
//
//		The minimum reader buffer size is 16-bytes. If
//		'readerBuffSize' is set to a value less than
//		"16", it will be automatically reset to the
//		default buffer size of 4096-bytes.
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
//		FileBufferReadWrite.Close().
//
//		While the 'write' services provided by the
//		FileBufferReadWrite are primarily designed for
//		writing data to disk files, this type of 'writer'
//		will in fact write data to any object
//		implementing the io.Writer interface.
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
//	writerBuffSize				int
//
//		This integer value controls the size of the
//		buffer created for the io.Writer object passed as
//		input parameter 'writer'.
//
//		'writerBuffSize' should be configured to maximize
//		performance for 'write' operations subject to
//		prevailing memory limitations.
//
//		The minimum write buffer size is 1-byte. If
//		'writerBuffSize' is set to a size less than or
//		equal to zero, it will be automatically reset to
//		the default buffer size of 4096-bytes.
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
func (fBufReadWriteNanobot *fileBufferReadWriteNanobot) setIoReaderWriter(
	fBufReadWrite *FileBufferReadWrite,
	fBufReadWriteLabel string,
	reader io.Reader,
	readerLabel string,
	readerBuffSize int,
	writer io.Writer,
	writerLabel string,
	writerBuffSize int,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fBufReadWriteNanobot.lock == nil {
		fBufReadWriteNanobot.lock = new(sync.Mutex)
	}

	fBufReadWriteNanobot.lock.Lock()

	defer fBufReadWriteNanobot.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileBufferReadWriteNanobot." +
		"setIoReaderWriter()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return err
	}

	if len(fBufReadWriteLabel) == 0 {

		fBufReadWriteLabel = "fBufReadWrite"
	}

	if len(readerLabel) == 0 {

		readerLabel = "reader"
	}

	if len(writerLabel) == 0 {

		writerLabel = "writer"
	}

	if fBufReadWrite == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is a nil pointer!\n"+
			"%v is invalid.\n",
			ePrefix.String(),
			fBufReadWriteLabel,
			fBufReadWriteLabel)

		return err
	}

	var newBuffReader FileBufferReader
	var err2 error

	err2 =
		new(fileBufferReaderNanobot).setIoReader(
			&newBuffReader,
			"newBuffReader",
			reader,
			"reader",
			readerBuffSize,
			ePrefix.XCpy("newBuffReader<-reader"))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"An error occurred while creating the new %v.reader.\n"+
			"Error=\n%v\n",
			funcName,
			fBufReadWriteLabel,
			err2.Error())

		return err

	}

	var newBuffWriter FileBufferWriter

	err2 =
		new(fileBufferWriterNanobot).setIoWriter(
			&newBuffWriter,
			"newBuffWriter",
			writer,
			"writer",
			writerBuffSize,
			ePrefix.XCpy("newBuffWriter<-writer"))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"An error occurred while creating the new %v.writer.\n"+
			"Error=\n%v\n",
			funcName,
			fBufReadWriteLabel,
			err2.Error())

		return err

	}

	fBufReadWrite.reader = &newBuffReader

	fBufReadWrite.writer = &newBuffWriter

	return err
}

// setPathFileNames
func (fBufReadWriteNanobot *fileBufferReadWriteNanobot) setPathFileNames(
	fBufReadWrite *FileBufferReadWrite,
	fBufReadWriteLabel string,
	readerPathFileName string,
	readerPathFileNameLabel string,
	openReadFileReadWrite bool,
	readerBuffSize int,
	writerPathFileName string,
	writerPathFileNameLabel string,
	openWriteFileReadWrite bool,
	writerBuffSize int,
	truncateExistingWriteFile bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fBufReadWriteNanobot.lock == nil {
		fBufReadWriteNanobot.lock = new(sync.Mutex)
	}

	fBufReadWriteNanobot.lock.Lock()

	defer fBufReadWriteNanobot.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileBufferReadWriteNanobot." +
		"setIoReaderWriter()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return err
	}

	if len(fBufReadWriteLabel) == 0 {

		fBufReadWriteLabel = "fBufReadWrite"
	}

	if len(readerPathFileNameLabel) == 0 {

		readerPathFileNameLabel = "readerPathFileName"
	}

	if len(writerPathFileNameLabel) == 0 {

		writerPathFileNameLabel = "writerPathFileName"
	}

	if fBufReadWrite == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is a nil pointer!\n"+
			"%v is invalid.\n",
			ePrefix.String(),
			fBufReadWriteLabel,
			fBufReadWriteLabel)

		return err
	}

	if len(readerPathFileName) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is invalid!\n"+
			"'%v' is an empty string with a length of zero (0).\n",
			ePrefix.String(),
			readerPathFileNameLabel,
			readerPathFileNameLabel)

		return err
	}

}

type fileBufferReadWriteMolecule struct {
	lock *sync.Mutex
}

// close
//
// This method is designed to perform clean up tasks
// after completion of all 'read' and 'write' operations.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will effectively render the instance of
//	FileBufferReadWrite, passed as input parameter
//	'fBufReadWrite', invalid and unusable for any
//	future 'read' or 'write' operations.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fBufReadWrite				*FileBufferReadWrite
//
//		A pointer to an instance of FileBufferWriter.
//
//		The internal FileBufferReader and
//		FileBufferWriter objects encapsulated in this
//		instance of FileBufferReadWrite will be deleted
//		as part of this 'close' operation.
//
//	fBufReadWriteLabel			string
//
//		The name or label associated with input parameter
//		'fBufReadWrite' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fBufReadWrite" will
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
func (fBuffReadWriteMolecule *fileBufferReadWriteMolecule) close(
	fBufReadWrite *FileBufferReadWrite,
	fBufReadWriteLabel string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fBuffReadWriteMolecule.lock == nil {
		fBuffReadWriteMolecule.lock = new(sync.Mutex)
	}

	fBuffReadWriteMolecule.lock.Lock()

	defer fBuffReadWriteMolecule.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileBufferReadWriteMolecule." +
		"close()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return err
	}

	if len(fBufReadWriteLabel) == 0 {

		fBufReadWriteLabel = "fBufReadWrite"
	}

	if fBufReadWrite == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is a nil pointer!\n"+
			"%v is invalid.\n",
			ePrefix.String(),
			fBufReadWriteLabel,
			fBufReadWriteLabel)

		return err
	}

	var err2 error

	if fBufReadWrite.reader != nil {

		err2 = new(fileBufferReaderMolecule).close(
			fBufReadWrite.reader,
			fBufReadWriteLabel+".reader",
			ePrefix.XCpy(
				fBufReadWriteLabel+".reader"))

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"An error occurred while closing %v.reader.\n"+
				"Error=\n%v\n",
				funcName,
				fBufReadWriteLabel,
				err2.Error())

			return err
		}

		fBufReadWrite.reader = nil
		fBufReadWrite.readFilePathName = ""

	}

	if fBufReadWrite.writer != nil {

		err2 = new(fileBufferWriterMolecule).close(
			fBufReadWrite.writer,
			fBufReadWriteLabel+".writer",
			ePrefix.XCpy(
				fBufReadWriteLabel+".writer"))

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"An error occurred while closing %v.writer.\n"+
				"Error=\n%v\n",
				funcName,
				fBufReadWriteLabel,
				err2.Error())

			return err
		}

		fBufReadWrite.writer = nil
		fBufReadWrite.writeFilePathName = ""

	}

	return err
}
