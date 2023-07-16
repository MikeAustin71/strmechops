package strmech

import (
	"bufio"
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"sync"
)

// FileBufferReader
//
// This structure and the associated methods are designed
// to facilitate data read operations. The most common
// data source for these read operations is assumed to be
// a data file residing on an attached storage drive.
// However, any object implementing the io.Reader
// interface may be used as a data source.
//
// The FileBufferReader type is a wrapper for
// 'bufio.Reader'. As such, FileBufferReader supports
// incremental read operations from the data source.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://pkg.go.dev/bufio
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Use the 'New' method to create new instances of
//	FileBufferReader.
type FileBufferReader struct {
	fileReader *bufio.Reader

	lock *sync.Mutex
}

// New
//
// This method returns a fully initialized instance of
// FileBufferReader.
//
// The size of the internal buffer is controlled by input
// parameter 'bufSize'. The minimum buffer size is
// 16-bytes.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	reader						io.Reader
//
//		An object which implements io.Reader interface.
//		Typically, this is a file pointer of type
//		*os.File.
//
//		A file pointer (*os.File) will facilitate reading
//		file data from files residing on an attached
//		storage drive.
//
//		While the returned instance of FileBufferReader
//		is primarily designed for reading file data, the
//		reader will in fact read data from any object
//		implementing the io.Reader interface.
//
//	bufSize						int
//
//		This integer value controls the size of the
//		buffer created for the returned instance of
//		FileBufferReader.
//
//		'bufSize' should be configured to maximize
//		performance for 'read' operations subject to
//		prevailing memory limitations.
//
//		The minimum reader buffer size is 16-bytes. If
//		'bufSize' is set to a size less than '16', an
//		error will be returned.
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
//	FileBufferReader
//
//		If this method completes successfully, a fully
//		configured instance of FileBufferReader will
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
func (fBufReader *FileBufferReader) New(
	reader io.Reader,
	bufSize int,
	errorPrefix interface{}) (
	FileBufferReader,
	error) {

	if fBufReader.lock == nil {
		fBufReader.lock = new(sync.Mutex)
	}

	fBufReader.lock.Lock()

	defer fBufReader.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error
	var newFileBufReader FileBufferReader

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"FileBufferReader."+
			"NewInitializeReader()",
		"")

	if err != nil {
		return newFileBufReader, err
	}

	if bufSize < 16 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'bufSize' is invalid!\n"+
			"The minimum buffer size is '16'.\n"+
			"Input parameter 'bufSize' = '%v'\n",
			ePrefix.String(),
			bufSize)

	}

	newFileBufReader.fileReader = bufio.NewReaderSize(
		reader,
		bufSize)

	return newFileBufReader, err
}
