package strmech

import (
	"bufio"
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"os"
	"sync"
)

// FileBufferWriter
//
// This structure and the associated methods are designed
// to facilitate data 'write' operations. The most common
// destination for these 'write' operations is assumed to
// be a data file residing on an attached storage drive.
// However, any object implementing the io.Writer
// interface may be used as a 'write' destination.
//
// The FileBufferWriter type is a wrapper for
// 'bufio.Writer'. As such, FileBufferWriter supports
// incremental or buffered 'write' operations to the
// target output destination.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://pkg.go.dev/bufio
//	https://pkg.go.dev/bufio#Writer
//	https://pkg.go.dev/io#Writer
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	Use methods 'New' and 'NewPathFileName' to create
//		new instances of FileBufferWriter.
//
//	(2)	FileBufferWriter implements the io.Writer
//		interface.
//
// ----------------------------------------------------------------
//
// # Best Practice
//
//	(1)	Create a new instance of FileBufferWriter using
//		either the New() method or the NewPathFileName()
//		method.
//
//		(a)	The New() method is used when an instance of
//			io.Writer is created externally by the user
//			and passed to the FileBufferWriter.New()
//			method.
//
//			Under this scenario, the user is independently
//			responsible for clean-up of the io.Writer
//			object after FileBufferWriter 'write'
//			operations have been completed.
//
//			Once all FileBufferWriter 'write' operations
//			have been completed, call methods Flush() and
//			Close() to perform local FileBufferWriter
//			clean-up tasks.
//
//		(b)	The NewPathFileName() method allows for the
//			creation of an internal file pointer to a
//			file passed as a path and file name by the
//			user. This file serves at the target
//			io.Writer object to which data will be
//			written.
//
//			Under this scenario, the user simply calls
//			methods Flush() and Close() in sequence to
//			perform all required clean-up tasks after
//			'write' operations have been completed.
//
//			Once method Close() is called, the current
//			FileBufferWriter instance becomes unusable
//			and should be discarded.
//
//	(2)	After creating an instance of FileBufferWriter,
//		the user calls the Write() method to write bytes
//		of data to target io.Writer object. This 'write'
//		target may be a file or any other object which
//		implements the io.Writer interface.
//
//		The Write() method should be called repeatedly
//		until all data has been written to the underlying
//		io.Writer object.
//
//		Upon completion of the 'write' operation, call
//		methods Flush() and Close() in sequence to
//		perform required clean-up tasks.
//
//	(3)	After all data bytes have been written to the
//		target io.Writer object, the user must call
//		methods Flush() and Close() to perform necessary
//		clean-up operations.
//
//		Once method Close() is called, the current
//		FileBufferWriter instance becomes unusable
//		and should be discarded.
type FileBufferWriter struct {
	fileWriter          *bufio.Writer
	filePtr             *os.File
	targetWriteFileName string

	lock *sync.Mutex
}

// Close
//
// This method is used to close any open file pointers
// and perform necessary clean-up operations after all
// data has been written to the destination io.Writer
// object.
//
// These clean-up operations include flushing the
// write buffer to ensure that all data is written to
// the destination io.Writer object.
//
// After calling this method, the current instance of
// FileBufferWriter will be unusable and should be
// discarded.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	Call this method after completing all write
//		operations. Calling this method is essential to
//		performance of necessary clean-up tasks after
//		completion of all 'write' operations.
//
//	(2)	Once this method completes the 'Close' operation,
//		this instance of FileBufferWriter becomes
//		invalid, unusable and unavailable for further
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
func (fBufWriter *FileBufferWriter) Close(
	errorPrefix interface{}) error {

	if fBufWriter.lock == nil {
		fBufWriter.lock = new(sync.Mutex)
	}

	fBufWriter.lock.Lock()

	defer fBufWriter.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileBufferWriter."+
			"Close()",
		"")

	if err != nil {
		return err
	}

	var err2 error
	var errs []error

	if fBufWriter.fileWriter != nil {

		err2 = fBufWriter.fileWriter.Flush()

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Error returned while flushing the 'write' buffer!\n"+
				"fBufWriter.fileWriter.Flush()\n"+
				"Error = \n%v\n",
				ePrefix.String(),
				err2.Error())

			errs = append(errs, err)

		}

		fBufWriter.fileWriter = nil

	}

	if fBufWriter.filePtr != nil {

		err2 = fBufWriter.filePtr.Close()

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Error returned while closing the target 'write' file!\n"+
				"fBufWriter.filePtr.Close()\n"+
				"Error = \n%v\n",
				ePrefix.String(),
				err2.Error())

			errs = append(errs, err)

		}

		fBufWriter.filePtr = nil
	}

	if len(errs) > 0 {
		err = new(StrMech).ConsolidateErrors(errs)

	} else {

		err = nil
	}

	fBufWriter.fileWriter = nil
	fBufWriter.targetWriteFileName = ""
	fBufWriter.filePtr = nil

	return err
}

// Flush
//
// Calling this method ensures that all remaining data in
// 'write' buffer will be written to the destination
// io.Writer object.
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
func (fBufWriter *FileBufferWriter) Flush(
	errorPrefix interface{}) error {

	if fBufWriter.lock == nil {
		fBufWriter.lock = new(sync.Mutex)
	}

	fBufWriter.lock.Lock()

	defer fBufWriter.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileBufferWriter."+
			"Close()",
		"")

	if err != nil {
		return err
	}

	if fBufWriter.fileWriter == nil {
		return err
	}

	var err2 error

	err2 = fBufWriter.fileWriter.Flush()

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned while flushing the 'write' buffer!\n"+
			"fBufWriter.fileWriter.Flush()\n"+
			"Error = \n%v\n",
			ePrefix.String(),
			err2.Error())

	}

	return err
}

// New
//
// This method returns a fully initialized instance of
// FileBufferWriter.
//
// This returned instance of FileBufferWriter is created
// using an object implementing the io.Writer interface
// and passed as input parameter 'writer'.
//
// The size of the internal read buffer is controlled by
// input parameter 'bufSize'. The minimum buffer size is
// 16-bytes.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	Input parameter 'writer' will accept a pointer to an
//	instance of os.File because os.File implements the
//	io.Writer interface.
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
//		operations.
//
//		While the returned instance of FileBufferWriter
//		is primarily designed for writing data to disk
//		files, this 'writer' will in fact write data to
//		any object implementing the io.Writer interface.
//
//	bufSize						int
//
//		This integer value controls the size of the
//		buffer created for the returned instance of
//		FileBufferWriter.
//
//		'bufSize' should be configured to maximize
//		performance for 'write' operations subject to
//		prevailing memory limitations.
//
//		If 'bufSize' is set to a value less than or equal
//		to zero (0), it will automatically be reset to
//		the default buffer size of 4096-bytes.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	FileBufferWriter
//
//		This method will return a fully configured
//		instance of FileBufferWriter.
func (fBufWriter *FileBufferWriter) New(
	writer io.Writer,
	bufSize int) FileBufferWriter {

	if fBufWriter.lock == nil {
		fBufWriter.lock = new(sync.Mutex)
	}

	fBufWriter.lock.Lock()

	defer fBufWriter.lock.Unlock()

	if bufSize <= 0 {

		bufSize = 4096
	}

	var newFileBufWriter FileBufferWriter

	newFileBufWriter.fileWriter = bufio.NewWriterSize(
		writer,
		bufSize)

	return newFileBufWriter
}

// NewPathFileName
//
// Receives a path and file name as an input parameter.
// This target 'write' file is opened for either
// 'write-only' or 'read/write' operations depending on
// input parameter 'openFileReadWrite'.
//
// Upon completion, this method returns a fully
// configured instance of FileBufferWriter.
//
// If the target path and file do not currently exist on
// an attached storage drive, this method will attempt to
// create them.
//
// The size of the internal 'write' buffer is controlled
// by input parameter 'bufSize'. If 'bufSize' is set to a
// value less than or equal to zero (0), it will be
// automatically reset to the default value of
// 4096-bytes.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	The returned type, 'FileBufferWriter', implements the
//	io.Writer interface.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pathFileName				string
//
//		This string contains the path and file name of
//		the target 'write' file which will be used a
//		destination data source for 'write' operations
//		performed by method:
//
//			FileBufferWriter.Write()
//
//		If the target path and file do not currently
//		exist on an attached storage drive, this method
//		will attempt to create them.
//
//	openFileReadWrite			bool
//
//		If this parameter is set to 'true', the target
//		'write' file created from input parameter
//		'pathFileName' will be opened for 'read' and
//		'write' operations.
//
//		If 'openFileReadWrite' is set to 'false', the
//		target write file will be opened for 'write-only'
//		operations.
//
//	bufSize						int
//
//		This integer value controls the size of the
//		'write' buffer created for the returned instance
//		of FileBufferWriter.
//
//		'bufSize' should be configured to maximize
//		performance for 'write' operations subject to
//		prevailing memory limitations.
//
//		If 'bufSize' is set to a value less than or equal
//		to zero (0), it will be automatically reset to
//		the default value of 4096-bytes.
//
//	truncateExistingFile			bool
//
//		If this parameter is set to 'true', the target
//		'write' file will be opened for write operations.
//		If the target file previously existed, it will be
//		truncated. This means that the file's previous
//		contents will be deleted.
//
//		If this parameter is set to 'false', the target
//		file will be opened for write operations. If the
//		target file previously existed, the new text
//		written to the file will be appended to the
//		end of the previous file contents.
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
//	FileBufferWriter
//
//		If this method completes successfully, a fully
//		configured instance of FileBufferWriter will
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
func (fBufWriter *FileBufferWriter) NewPathFileName(
	pathFileName string,
	openFileReadWrite bool,
	bufSize int,
	truncateExistingFile bool,
	errorPrefix interface{}) (
	FileBufferWriter,
	error) {

	if fBufWriter.lock == nil {
		fBufWriter.lock = new(sync.Mutex)
	}

	fBufWriter.lock.Lock()

	defer fBufWriter.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error
	var newFileBufWriter FileBufferWriter

	funcName := "FileBufferWriter." +
		"NewPathFileName()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		"")

	if err != nil {
		return newFileBufWriter, err
	}

	if len(pathFileName) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'pathFileName' is invalid!\n"+
			"'pathFileName' is an empty string with a length of zero (0).\n",
			ePrefix.String())

		return newFileBufWriter, err
	}

	if bufSize <= 0 {

		bufSize = 4096

	}

	var err2 error

	pathFileName,
		_,
		_,
		err2 =
		new(fileHelperMolecule).
			doesPathFileExist(
				pathFileName,
				PreProcPathCode.AbsolutePath(), // Convert to Absolute Path
				ePrefix,
				"pathFileName")

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"An error occurred while testing for the existance\n"+
			"of 'pathFileName' on an attached storage drive.\n"+
			"pathFileName = '%v'\n"+
			"Error= \n%v\n",
			funcName,
			pathFileName,
			err2.Error())

		return newFileBufWriter, err
	}

	var filePermissionCfg FilePermissionConfig

	var filePermissionStr = "--w--w--w-"

	if openFileReadWrite == true {

		filePermissionStr = "-rw-rw-rw-"
	}

	filePermissionCfg,
		err = new(FilePermissionConfig).New(
		filePermissionStr,
		ePrefix.XCpy("filePermissionCfg<-"))

	if err != nil {

		return newFileBufWriter, err
	}

	var fileOpenCfg FileOpenConfig

	if truncateExistingFile {

		fileOpenCfg,
			err = new(FileOpenConfig).New(
			ePrefix.XCpy("fileOpenCfg<-"),
			FOpenType.TypeWriteOnly(),
			FOpenMode.ModeCreate(),
			FOpenMode.ModeTruncate())

		if err != nil {

			return newFileBufWriter, err
		}

	} else {
		// truncateExistingFile = 'false'
		// This signals Append to existing file.

		fileOpenCfg,
			err = new(FileOpenConfig).New(
			ePrefix.XCpy("fileOpenCfg<-"),
			FOpenType.TypeWriteOnly(),
			FOpenMode.ModeCreate(),
			FOpenMode.ModeAppend())

		if err != nil {

			return newFileBufWriter, err
		}

	}

	newFileBufWriter.filePtr,
		err = new(fileHelperBoson).
		openFile(
			pathFileName,
			false,
			fileOpenCfg,
			filePermissionCfg,
			"pathFileName",
			ePrefix)

	if err != nil {

		if newFileBufWriter.filePtr != nil {
			_ = newFileBufWriter.filePtr.Close()
		}

		return newFileBufWriter, err
	}

	newFileBufWriter.targetWriteFileName = pathFileName

	newFileBufWriter.fileWriter = bufio.NewWriterSize(
		newFileBufWriter.filePtr,
		bufSize)

	return newFileBufWriter, err
}

// Write
//
// Writes the contents of the byte array input paramter
// ('bytesToWrite') to the destination io.Writer object
// previously configured for this instance of
// FileBufferWriter.
//
// If for any reason, the returned number of bytes
// written ('numBytesWritten') to the destination
// io.Writer object is less than the length of the byte
// array passed as input parameter 'bytesToWrite', an
// error containing an explanation for this event will
// be returned.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	After all write operations have been completed, the
//	user MUST call FileBufferWriter.Close() to perform
//	necessary clean-up operations.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	bytesToWrite				[]byte
//
//		The contents of this byte array will be written
//		to the destination io.Writer object previously
//		configured for the current instance of
//		FileBufferWriter.
//
//		Typically, the destination io.Writer object will
//		be a data file existing on an attached storage
//		drive. However, the destination io.Writer object
//		may be any object implementing the io.Writer
//		interface.
//
//		If for any reason, the returned number of bytes
//		written ('numBytesWritten') to the destination
//		io.Writer object is less than the length of this
//		byte array ('bytesToWrite'), an error containing
//		an explanation for this event will be returned.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	numBytesWritten				int
//
//		This parameter returns the number of bytes
//		written to the destination io.Writer object
//		configured for the current instance of
//		FileBufferWriter.
//
//		If 'numBytesWritten' is less than the length
//		of the byte array input parameter 'bytesToWrite',
//		an error will also be returned.
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
func (fBufWriter *FileBufferWriter) Write(
	bytesToWrite []byte) (
	numBytesWritten int,
	err error) {

	if fBufWriter.lock == nil {
		fBufWriter.lock = new(sync.Mutex)
	}

	fBufWriter.lock.Lock()

	defer fBufWriter.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"FileBufferWriter."+
			"Write()",
		"")

	if err != nil {

		return numBytesWritten, err
	}

	if fBufWriter.fileWriter == nil {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of 'FileBufferWriter' is invalid!\n"+
			"The internal bufio.Writer has NOT been initialized.\n"+
			"Call one of the 'New' methods when creating an instance\n"+
			"of 'FileBufferWriter'\n",
			ePrefix.String())

		return numBytesWritten, err
	}

	if len(bytesToWrite) <= 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'bytesToWrite' is invalid!\n"+
			"The 'bytesToWrite' byte array is empty. It has zero bytes.\n",
			ePrefix.String())

		return numBytesWritten, err
	}

	var err2 error

	numBytesWritten,
		err2 = fBufWriter.fileWriter.Write(bytesToWrite)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by fBufWriter.fileWriter.Write(bytesToWrite).\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			err2.Error())
	}

	return numBytesWritten, err
}

type fileBufferWriterNanobot struct {
	lock *sync.Mutex
}

type fileBufferWriterMolecule struct {
	lock *sync.Mutex
}

// close
//
// This method is designed to perform clean up
// operations after completion of all 'write'
// operations.
//
// All internal member variable data values for the
// instance of FileBufferWriter passed as input parameter
// ”.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	This method will delete all pre-existing data
//		values in the instance of FileBufferWriter passed
//		as input parameter 'fBufWriter'.
//
//		After completion of this method this
//		FileBufferWriter instance will be unusable,
//		invalid and unavailable for future 'write'
//		operations.
//
//	(2)	This 'close' method will NOT flush the 'write'
//		buffer. To flush the 'write' buffer call:
//			fileBufferWriterMolecule.flush()
//
//		Be sure to call the 'flush()' method before you
//		call this method ('close()').
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fBufWriter					*FileBufferWriter
//
//		A pointer to an instance of FileBufferWriter.
//
//		All internal member variable data values in
//		this instance will be deleted.
//
//		If a file pointer (*os.File) was previously
//		configured for 'fBufWriter', it will be closed
//		and set to 'nil' by this method.
//
//	fBufWriterLabel				string
//
//		The name or label associated with input parameter
//		'fBufWriter' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fBufWriter" will be
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
func (fBufWriterMolecule *fileBufferWriterMolecule) close(
	fBufWriter *FileBufferWriter,
	fBufWriterLabel string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fBufWriterMolecule.lock == nil {
		fBufWriterMolecule.lock = new(sync.Mutex)
	}

	fBufWriterMolecule.lock.Lock()

	defer fBufWriterMolecule.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileBufferWriterMolecule." +
		"close()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return err
	}

	if len(fBufWriterLabel) == 0 {

		fBufWriterLabel = "fBufWriter"
	}

	if fBufWriter == nil {

		err = fmt.Errorf("%v\n"+
			"Error: The FileBufferReader instance passed\n"+
			"as input parameter '%v' is invalid!\n"+
			"'%v' is a 'nil' pointer.\n",
			ePrefix.String(),
			fBufWriterLabel,
			fBufWriterLabel)

		return err
	}

	if fBufWriter.filePtr != nil {

		var err2 error

		err2 = fBufWriter.filePtr.Close()

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Error returned while closing the target 'target' file!\n"+
				"fBufWriter.filePtr.Close()\n"+
				"Target Read File = '%v'\n"+
				"Error = \n%v\n",
				ePrefix.String(),
				fBufWriter.targetWriteFileName,
				err2.Error())

		}
	}

	fBufWriter.targetWriteFileName = ""

	fBufWriter.filePtr = nil

	fBufWriter.fileWriter = nil

	return err
}

// flush
//
// This method performs one function. Namely, it flushes
// all data from the write file effectively ensuring that
// all data in the buffer is written to the file or
// underlying device defined as an io.Writer.
//
// Specifically, this method does NOT 'close' the
// FileBufferWriter instance passed as input parameter
// 'fBufWriter'. As such, this method does not delete
// member variable data contained in 'fBufWriter'. To
// fully close the 'fBufWriter' instance, make a separate
// call to method 'fileBufferWriterMolecule.close()'.
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	This method will NOT modify the internal member
//	variable data values for the instance of
//	FileBufferWriter passed as input parameter
//	'fBufWriter'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fBufWriter					*FileBufferWriter
//
//		A pointer to an instance of FileBufferWriter.
//		Any data remaining in the 'write' buffer will
//		be written to the underlying data file by the
//		flush 'operation' performed by this method.
//
//	fBufWriterLabel				string
//
//		The name or label associated with input parameter
//		'fBufWriter' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "fBufWriter" will be
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
func (fBufWriterMolecule *fileBufferWriterMolecule) flush(
	fBufWriter *FileBufferWriter,
	fBufWriterLabel string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if fBufWriterMolecule.lock == nil {
		fBufWriterMolecule.lock = new(sync.Mutex)
	}

	fBufWriterMolecule.lock.Lock()

	defer fBufWriterMolecule.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileBufferWriterMolecule." +
		"flush()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {
		return err
	}

	if len(fBufWriterLabel) == 0 {

		fBufWriterLabel = "fBufWriter"
	}

	if fBufWriter == nil {

		err = fmt.Errorf("%v\n"+
			"Error: The FileBufferReader instance passed\n"+
			"as input parameter '%v' is invalid!\n"+
			"'%v' is a 'nil' pointer.\n",
			ePrefix.String(),
			fBufWriterLabel,
			fBufWriterLabel)

		return err
	}

	if fBufWriter.fileWriter == nil {
		return err
	}

	var err2 error

	err2 = fBufWriter.fileWriter.Flush()

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned while flushing the 'write' buffer!\n"+
			"%v.fileWriter.Flush()\n"+
			"Error = \n%v\n",
			ePrefix.String(),
			fBufWriterLabel,
			err2.Error())

	}

	return err
}
