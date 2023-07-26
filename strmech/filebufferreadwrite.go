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
type FileBufferReadWrite struct {
	writer            FileBufferWriter
	reader            FileBufferReader
	writeFilePathName string
	readFilePathName  string

	lock *sync.Mutex
}

// New
//
// Creates and returns a new, fully configured instance
// of FileBufferReadWrite
func (fBufReadWrite *FileBufferReadWrite) New(
	reader io.Reader,
	readerBuffSize int,
	writer io.Writer,
	writerBuffSize int) FileBufferReadWrite {

	if fBufReadWrite.lock == nil {
		fBufReadWrite.lock = new(sync.Mutex)
	}

	fBufReadWrite.lock.Lock()

	defer fBufReadWrite.lock.Unlock()

	var newFBuffReadWrite = FileBufferReadWrite{}

	newFBuffReadWrite.reader =
		new(FileBufferReader).New(
			reader,
			readerBuffSize)

	newFBuffReadWrite.writer =
		new(FileBufferWriter).New(
			writer,
			writerBuffSize)

	return newFBuffReadWrite
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

	newFBuffReadWrite.reader,
		err =
		new(FileBufferReader).NewPathFileName(
			readerPathFileName,
			openReadFileReadWrite, // openFileReadWrite
			readerBuffSize,
			ePrefix.XCpy("<-readerPathFileName"))

	if err != nil {

		return newFBuffReadWrite, err
	}

	newFBuffReadWrite.writer,
		err = new(FileBufferWriter).NewPathFileName(
		writerPathFileName,
		openWriteFileReadWrite, // openFileReadWrite
		writerBuffSize,
		truncateExistingWriteFile, // truncateExistingFile
		ePrefix.XCpy("<-"))

	return newFBuffReadWrite, err
}
