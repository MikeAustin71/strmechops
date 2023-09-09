package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type fileWriterHelperMicrobot struct {
	lock *sync.Mutex
}

// writeCharacters
//
// This method will accept many different text or numeric
// data types which are then converted to a byte or
// string array and written to the io.Writer object
// passed as input parameter 'ioWriter'.
//
// The text or numeric data type passed as input
// parameter 'charsToConvert' must match one of over fifty
// eligible data types.
//
// If 'charsToConvert' is set to an ineligible data type,
// an error will be returned.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://golangdocs.com/anonymous-functions-in-golang
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	writeBytes 	func(
//					bytesToWrite []byte,
//					xWriteEndOfTextChars string,
//					xEPrefix *ePref.ErrPrefixDto) (
//					xNumOfBytesWritten int64,
//					xErr error)
//
//		This parameter is an anonymous function with the
//		indicated function signature. See the Reference
//		section on "anonymous functions" above.
//
//			Input Parameters
//
//				bytesToWrite 			[]byte
//					A byte array containing the bytes to
//					be written to the internal io.Writer
//					object wrapped by this anonymous
//					function.
//
//				xWriteEndOfTextChars	string
//					A string containing the text characters
//					which will be written to the internal
//					io.Writer object after all the bytes in
//					'bytesToWrite' have been processed and
//					written.
//
//				xErrPref 				*ePref.ErrPrefixDto
//
//					This object encapsulates an error prefix
//					string which is included in all returned
//					error messages. Usually, it contains the
//					name of the calling method or methods
//					listed as a function chain.
//
//					If no error prefix information is needed,
//					set this parameter to 'nil'.
//
//					Type ErrPrefixDto is included in the
//					'errpref' software package:
//						"github.com/MikeAustin71/errpref".
//
//			Return Values
//
//				xNumOfBytesWritten 		int64
//					The number of bytes written to the internal
//					io.Writer.
//
//				xErr 					error
//					If this anonymous method completes
//					successfully, the returned error Type is set
//					equal to 'nil'.
//
//					If errors are encountered during processing,
//					the returned error Type will encapsulate an
//					appropriate error message. This returned
//					error message will incorporate the method
//					chain and text passed by input parameter,
//					'xErrPref'. The 'xErrPref' text will be
//					prefixed or attached to the	beginning of the
//					error message.
//
//	charsToConvert				interface{}
//
//		This empty interface is used to transmit an
//		eligible text or numeric data type which will be
//		to a string or byte array and written to the
//		io.Writer object passed as input parameter
//		'ioWriter'.
//
//		If the type transmitted through this parameter
//		does not one of the following data types, an
//		error will be returned.
//
//			Eligible Data Types
//
//			 1.	[]byte
//			 2.	*[]byte
//			 3.	string
//			 4.	*string
//			 5.	[]string
//			 6.	*[]string
//			 7.	Stringer (fmt.Stringer) Interface
//			 8.	strings.Builder
//			 9.	*strings.Builder
//			10.	StringArrayDto
//			11.	*StringArrayDto
//			12.	[]rune
//			13.	*[]rune
//			14.	RuneArrayDto
//			15.	*RuneArrayDto
//			16.	RuneArrayCollection
//			17.	*RuneArrayCollection
//			18.	ITextFieldFormatDto
//			19.	ITextFieldSpecification
//			20.	ITextLineSpecification
//			21.	TextLineSpecLinesCollection
//			22.	bool
//			23.	TextLineTitleMarqueeDto
//			24.	time.Time
//			25.	TextInputParamFieldDateTimeDto
//			26.	float32
//			27.	*float32
//			28.	float64
//			29.	*float64
//			30.	BigFloatDto
//			31.	*BigFloatDto
//			32.	big.Float
//			33.	*big.Float
//			34.	big.Rat
//			35.	*big.Rat
//			36.	int8
//			37.	*int8
//			38.	int16
//			39.	*int16
//			40.	int
//			41.	*int
//			42.	int32
//			43.	*int32
//			44.	int64
//			45.	*int64
//			46.	uint8
//			47.	*uint8
//			48.	uint16
//			49.	*uint16
//			50.	uint
//			51.	*uint
//			52.	uint32
//			53.	*uint32
//			54.	uint64,
//			55.	*uint64
//			56.	big.Int
//			57.	*big.Int
//			58.	TextFieldFormatDtoFloat64
//			59.	*TextFieldFormatDtoFloat64
//			60.	TextFieldFormatDtoBigFloat
//			61.	*TextFieldFormatDtoBigFloat
//			62.	NumberStrKernel
//			63.	*NumberStrKernel
//			64.	[]NumberStrKernel
//			65.	*[]NumberStrKernel
//
//	charsToConvertLabel			string
//
//		The name or label associated with input parameter
//		'charsToConvert' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "charsToConvert" will be
//		automatically applied.
//
//	writeEndOfLineChars string
//
//		This character string is appended to each line of
//		text written to the io.Writer object. This
//		capability is more useful when processing string
//		arrays when each element of the array is written
//		separately to the io.Writer object.
//
//		Remember that on Windows, line-endings are
//		terminated with a combination of a carriage
//		return (ASCII 0x0d or \r) and a newline(\n), also
//		referred to as CR/LF (\r\n).
//
//		On UNIX or Linux, text file line-endings are
//		terminated with a newline character (ASCII 0x0a,
//		represented by the \n escape sequence in most
//		languages), also referred to as a linefeed (LF).
//
//		If 'writeEndOfLineChars' is set to an empty
//		string, it will be ignored and no additional
//		characters will be appended to each line written
//		to the io.Writer object.
//
//	writeEndOfTextChars			string
//
//		A character string which will be written to the
//		internal io.Writer object after all other text
//		from 'charsToConvert' and 'writeEndOfLineChars'
//		has been processed and written.
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
func (fWriterHlprMicrobot *fileWriterHelperMicrobot) writeCharacters(
	writeBytes func(
		bytesToWrite []byte,
		xWriteEndOfTextChars string,
		xErrPref *ePref.ErrPrefixDto) (
		xNumOfBytesWritten int64,
		xErr error),
	charsToConvert interface{},
	charsToConvertLabel string,
	writeEndOfLineChars string,
	writeEndOfTextChars string,
	errPrefDto *ePref.ErrPrefixDto) (
	numOfBytesWritten int64,
	err error) {

	if fWriterHlprMicrobot.lock == nil {
		fWriterHlprMicrobot.lock = new(sync.Mutex)
	}

	fWriterHlprMicrobot.lock.Lock()

	defer fWriterHlprMicrobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileWriterHelperMicrobot." +
		"writeCharacters()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		return numOfBytesWritten, err
	}

	if len(charsToConvertLabel) == 0 {

		charsToConvertLabel = "charsToConvert"
	}

	if charsToConvert == nil {

		err = fmt.Errorf("%v\n"+
			"-------------------------------------------------------\n"+
			"Error: Input parameter '%v' is invalid!\n"+
			"'%v' is a 'nil' value.\n",
			ePrefix.String(),
			charsToConvertLabel,
			charsToConvertLabel)

		return numOfBytesWritten, err
	}

	var strArray []string
	var baseTypeConversion BaseTypeDto

	baseTypeConversion,
		err = new(typeConversionsAtom).
		convertParamsToBaseTypes(
			charsToConvert,
			"charsToConvert",
			ePrefix)

	if err != nil {

		return numOfBytesWritten, err

	}

	if baseTypeConversion.IsAByteArrayDto == true {

		if baseTypeConversion.AByteArrayDto.
			GetArrayLength() == 0 {

			err = fmt.Errorf("%v\n"+
				"Error: The converted Byte Array is empty\n"+
				"and has a zero length.\n"+
				"The byte array was converted from original\n"+
				"input type %v.\n",
				ePrefix.String(),
				baseTypeConversion.AByteArrayDtoSourceDataType)

			return numOfBytesWritten, err
		}

		numOfBytesWritten,
			err = writeBytes(
			baseTypeConversion.AByteArrayDto.ByteArray,
			writeEndOfLineChars+writeEndOfTextChars,
			ePrefix.XCpy("byteArray<-AByteArrayDto"))

	} else if baseTypeConversion.IsARuneArrayDto == true {

		if baseTypeConversion.ARuneArrayDto.
			GetRuneArrayLength() == 0 {

			err = fmt.Errorf("%v\n"+
				"Error: The converted Rune Array in\n"+
				"RuneArrayDto is empty and has a zero length.\n"+
				"This rune array was converted from original\n"+
				"input type %v.\n",
				ePrefix.String(),
				baseTypeConversion.ARuneArrayDtoSourceDataType)

			return numOfBytesWritten, err
		}

		numOfBytesWritten,
			err = writeBytes(
			baseTypeConversion.ARuneArrayDto.
				GetByteArrayDto().ByteArray,
			writeEndOfLineChars+writeEndOfTextChars,
			ePrefix.XCpy("byteArray<-ARuneArrayDto"))

	} else if baseTypeConversion.IsAString == true {

		if baseTypeConversion.AStringLength == 0 {

			err = fmt.Errorf("%v\n"+
				"Error: The converted string is empty\n"+
				"and has a zero length.\n"+
				"This string was converted from original\n"+
				"input type %v.\n",
				ePrefix.String(),
				baseTypeConversion.AStringSourceDataType)

			return numOfBytesWritten, err
		}

		numOfBytesWritten,
			err = writeBytes(
			[]byte(baseTypeConversion.AString),
			writeEndOfLineChars+writeEndOfTextChars,
			ePrefix.XCpy("byteArray<-AString"))

	} else if baseTypeConversion.IsAStringArrayDto == true {

		if baseTypeConversion.AStringArrayDtoLength == 0 {

			err = fmt.Errorf("%v\n"+
				"Error: The converted string array in\n"+
				"StringArrayDto is empty and has a zero length.\n"+
				"This string array was converted from original\n"+
				"input type %v.\n",
				ePrefix.String(),
				baseTypeConversion.AStringArrayDtoSourceDataType)

			return numOfBytesWritten, err
		}

		numOfBytesWritten,
			err = new(fileWriterHelperMolecule).
			writeStrArray(
				writeBytes,
				strArray,
				charsToConvertLabel,
				writeEndOfLineChars,
				writeEndOfTextChars,
				ePrefix.XCpy(
					"byteArray<-string array"))

	} else {

		err = fmt.Errorf("%v\n"+
			"-------------------------------------------------------\n"+
			"Error: typeConversionsAtom.convertParamsToBaseTypes()\n"+
			"returned an invalid 'BaseTypeDto' object. The return"+
			"value did not resolve as a byte array, rune array,"+
			"string array or string.\n",
			ePrefix.String())

	}

	return numOfBytesWritten, err
}

type fileWriterHelperMolecule struct {
	lock *sync.Mutex
}

func (fWriterHlprMolecule *fileWriterHelperMolecule) writeStrArray(
	writeBytes func(
		bytesToWrite []byte,
		xWriteEndOfTextChars string,
		xErrPref *ePref.ErrPrefixDto) (
		xNumOfBytesWritten int64,
		xErr error),
	strArray []string,
	strArrayLabel string,
	writeEndOfLineChars string,
	writeEndOfTextChars string,
	errPrefDto *ePref.ErrPrefixDto) (
	numOfBytesWritten int64,
	err error) {

	if fWriterHlprMolecule.lock == nil {
		fWriterHlprMolecule.lock = new(sync.Mutex)
	}

	fWriterHlprMolecule.lock.Lock()

	defer fWriterHlprMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileWriterHelperMicrobot." +
		"writeCharacters()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		return numOfBytesWritten, err
	}

	if writeBytes == nil {

		err = fmt.Errorf("%v\n"+
			"Error: The anonymous function input parameter 'writeBytes'\n"+
			"is invalid. 'writeBytes' has a value of 'nil'.\n",
			ePrefix.String())

		return numOfBytesWritten, err
	}

	if len(strArrayLabel) == 0 {

		strArrayLabel = "strArray"
	}

	lastStrArrayIdx := len(strArray) - 1

	if lastStrArrayIdx < 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is invalid!\n"+
			"This string array is empty. It has zero bytes.\n",
			ePrefix.String(),
			strArrayLabel)

		return numOfBytesWritten, err
	}
	var localNumOfBytesWritten int64
	var err2 error
	var strToWrite string

	for i := 0; i <= lastStrArrayIdx; i++ {

		strToWrite = strArray[i]

		strToWrite += writeEndOfLineChars

		if i == lastStrArrayIdx {

			strToWrite += writeEndOfTextChars
		}

		if len(strToWrite) == 0 {
			continue
		}

		localNumOfBytesWritten,
			err2 = writeBytes(
			[]byte(strToWrite),
			"",
			ePrefix.XCpy("[]byte(strToWrite)<-strArray"))

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Error returned by  writeBytes() anonymous function!\n"+
				"strToWrite= '%v'\n"+
				"String Array Index= %v\n"+
				"Error=\n%v\n",
				ePrefix.String(),
				strToWrite,
				i,
				err2.Error())

			return numOfBytesWritten, err
		}

		numOfBytesWritten += localNumOfBytesWritten

	}

	return numOfBytesWritten, err
}
