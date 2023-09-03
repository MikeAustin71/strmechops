package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"sync"
)

// fileWriterHelperAtom
//
// A series of utility methods used by types:
//
//	FileBufferWriter
//	FileIoWriter
type fileWriterHelperAtom struct {
	lock *sync.Mutex
}

// writeBytes
//
// Writes a byte array to the io.Writer object
// contained in the FileIoWriter instance passed as input
// parameter 'ioWriter'
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	If the byte array passed as input parameter
//	'byteArray' is empty or contains zero array elements,
//	this method will take no action, no error will be
//	returned and the returned number of bytes written
//	('numOfBytesWritten') will be set to zero.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	ioWriter 					*io.Writer
//
//		A pointer to an instance of io.Writer.
//
//		The contents of the byte array passed as input
//		parameter 'byteArray' will be written to this
//		io.Writer object.
//
//		If this parameter is submitted with a value of
//		'nil', an error will be returned.
//
//	ioWriterLabel				string
//
//		The name or label associated with input parameter
//		'ioWriter' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "ioWriter" will be
//		automatically applied.
//
//	byteArray					[]byte
//
//		An array of bytes which will be written to the
//		io.Writer object passed as input parameter
//		'ioWriter'.
//
//		If parameter 'endOfLineTerminator' has a length
//		greater than zero, 'endOfLineTerminator' characters
//		will be to the byte array written to the io.Writer
//		object.
//
//		If 'byteArray' is empty or passed as a zero
//		length byte array, the method will take no
//		action, no error will be returned and the
//		returned number of bytes written
//		('numOfBytesWritten') will be set to zero.
//
//	byteArrayLabel				string
//
//		The name or label associated with input parameter
//		'byteArray' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "byteArray" will be
//		automatically applied.
//
//	endOfLineTerminator string
//
//		If this parameter has a string length greater
//		than zero, the text characters contained therein
//		will be appended to the byte array ('byteArray')
//		written to the io.Writer object passed as input
//		parameter 'ioWriter'.
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
//	numOfBytesWritten			int64
//
//		The number of bytes written to the io.Writer
//		object passed as input parameter 'ioWriter'.
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
//	 	text passed by input parameter, 'errPrefDto'.
//	 	The 'errPrefDto' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fWriterHelperAtom *fileWriterHelperAtom) writeBytes(
	ioWriter *io.Writer,
	ioWriterLabel string,
	byteArray []byte,
	byteArrayLabel string,
	endOfLineTerminator string,
	errPrefDto *ePref.ErrPrefixDto) (
	numOfBytesWritten int64,
	err error) {

	if fWriterHelperAtom.lock == nil {
		fWriterHelperAtom.lock = new(sync.Mutex)
	}

	fWriterHelperAtom.lock.Lock()

	defer fWriterHelperAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileWriterHelperAtom." +
		"writeBytes()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		return numOfBytesWritten, err
	}

	if len(ioWriterLabel) == 0 {

		ioWriterLabel = "writer"
	}

	if len(byteArrayLabel) == 0 {

		byteArrayLabel = "byteArray"
	}

	if ioWriter == nil {

		err = fmt.Errorf("%v\n"+
			"-------------------------------------------\n"+
			"Error: The io.Writer instance passed\n"+
			"as input parameter '%v' is invalid!\n"+
			"'%v' is a 'nil' pointer.\n",
			ePrefix.String(),
			ioWriterLabel,
			ioWriterLabel)

		return numOfBytesWritten, err
	}

	lenByteArray := len(byteArray)

	if lenByteArray == 0 {

		return numOfBytesWritten, err
	}

	var writer = *ioWriter
	var err2 error
	var localNumBytesWritten int

	localNumBytesWritten,
		err2 = writer.Write(
		byteArray)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by writer.Write(byteArray)"+
			"while writing original byte array.\n"+
			"byteArray= '%v'\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			string(byteArray),
			err2.Error())

		return numOfBytesWritten, err

	} else {

		numOfBytesWritten += int64(localNumBytesWritten)
	}

	if len(endOfLineTerminator) > 0 {

		localNumBytesWritten,
			err2 = writer.Write(
			[]byte(endOfLineTerminator))

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Error returned by writer.Write(byteArray)"+
				"while writing original byte array.\n"+
				"byteArray= '%v'\n"+
				"Error=\n%v\n",
				ePrefix.String(),
				string(byteArray),
				err2.Error())

		} else {

			numOfBytesWritten += int64(localNumBytesWritten)
		}
	}

	return numOfBytesWritten, err
}

// writeStringArray
//
// Writes a string array to the io.Writer object
// contained in the FileIoWriter instance passed as input
// parameter 'fIoWriter'
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	If the string array passed as input parameter
//	'strArray' is empty or contains zero array elements,
//	this method will take no action, no error will be
//	returned and the returned number of bytes written
//	'numOfBytesWritten' will be set to zero.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	fIoWriter					*FileIoWriter
//
//		A pointer to an instance of FileIoWriter.
//
//		The contents of the string array passed as input
//		parameter '' will be written to the internal
//		io.Writer object encapsulated by this
//		FileIoWriter instance.
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
//	strArray					[]string
//
//		An array of strings which will be written to
//		the internal io.Writer object encapsulated
//		within the FileIoWriter instance passed as input
//		parameter 'fIoWriter'.
//
//		If parameter 'endOfLineTerminator' has a length
//		greater than zero, 'endOfLineTerminator' will be
//		appended to each string written to the io.Writer
//		object.
//
//		If 'strArray' is empty or passed as a zero length
//		byte array, the method will take no action, no
//		error will be returned and the returned number of
//		bytes written ('numOfBytesWritten') will be set
//		to zero.
//
//	strArrayLabel				string
//
//		The name or label associated with input parameter
//		'strArray' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "strArray" will be
//		automatically applied.
//
//	endOfLineTerminator string
//
//		If this parameter has a string length greater
//		than zero, this string will be appended to
//		each string array element ('strArray') written
//		to the io.Writer object contained in 'fIoWriter'.
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
//	numOfBytesWritten			int64
//
//		The number of bytes written to the io.Writer
//		object encapsulated in the FileIoWriter instance
//		passed as input parameter 'fIoWriter'.
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
//	 	text passed by input parameter, 'errPrefDto'.
//	 	The 'errPrefDto' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (fWriterAtom *fileWriterHelperAtom) writeStringArray(
	fIoWriter *FileIoWriter,
	fIoWriterLabel string,
	strArray []string,
	strArrayLabel string,
	endOfLineTerminator string,
	errPrefDto *ePref.ErrPrefixDto) (
	numOfBytesWritten int64,
	err error) {

	if fIoWriterAtom.lock == nil {
		fIoWriterAtom.lock = new(sync.Mutex)
	}

	fIoWriterAtom.lock.Lock()

	defer fIoWriterAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileIoWriterAtom." +
		"writeStringArray()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		return numOfBytesWritten, err
	}

	if len(fIoWriterLabel) == 0 {

		fIoWriterLabel = "fIoWriter"
	}

	if len(strArrayLabel) == 0 {

		strArrayLabel = "strArray"
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

		return numOfBytesWritten, err
	}

	if fIoWriter.ioWriter == nil {

		err = fmt.Errorf("%v\n"+
			"-------------------------------------------------------\n"+
			"Error: This instance of 'FileIoWriter' passed"+
			"as input parameter '%v' is invalid!\n"+
			"The internal io.Writer has NOT been initialized.\n"+
			"Call one of the 'New' or 'Setter' methods when creating\n"+
			"a new valid instance of 'FileIoWriter'\n",
			fIoWriterLabel,
			ePrefix.String())

		return numOfBytesWritten, err
	}

	lenStrArray := len(strArray)
	lenEOLTerminator := len(endOfLineTerminator)

	if lenStrArray == 0 {

		return numOfBytesWritten, err
	}

	var err2 error
	var localNumBytesWritten int
	var writer = *fIoWriter.ioWriter
	var strToWrite string

	for i := 0; i <= lenStrArray; i++ {

		strToWrite = strArray[i]

		if lenEOLTerminator > 0 {
			strToWrite += endOfLineTerminator
		}

		localNumBytesWritten,
			err2 = writer.Write(
			[]byte(strToWrite))

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Error returned by writer.Write([]byte(strToWrite)).\n"+
				"strToWrite= '%v'\n"+
				"Error=\n%v\n",
				ePrefix.String(),
				strToWrite,
				err2.Error())

			return numOfBytesWritten, err

		} else {

			numOfBytesWritten += int64(localNumBytesWritten)
		}

		strToWrite = ""
	}

	return numOfBytesWritten, err
}
