package strmech

import (
	"bufio"
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
type FileBufferWriter struct {
	fileWriter          *bufio.Writer
	filePtr             *os.File
	targetWriteFileName string

	lock *sync.Mutex
}

// New
//
// This method returns a fully initialized instance of
// FileBufferWriter.
//
// The size of the internal read buffer is controlled by
// input parameter 'bufSize'. The minimum buffer size is
// 16-bytes.
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
