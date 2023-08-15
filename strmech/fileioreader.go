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
