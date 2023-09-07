package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"strings"
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
//				Eligible Data Types
//
//			   1.	[]byte
//			   2.	*[]byte
//			   3.	string
//			   4.	*string
//			   5.	[]string
//		 	   6.	StringArrayDto
//			   7.	*StringArrayDto
//			   8.	[]rune
//			   9.	*[]rune
//			  10.	RuneArrayDto
//			  11.	*RuneArrayDto
//			  12.	ITextFieldFormatDto
//			  13.	ITextFieldSpecification
//			  14.	ITextLineSpecification
//			  15.	float32
//			  16.	*float32
//			  17.	float64
//			  18.	*float64
//			  19.	BigFloatDto
//			  20.	*BigFloatDto
//			  21.	big.Float
//			  22.	*big.Float
//			  23.	big.Rat
//			  24.	*big.Rat
//			  25.	int8
//			  26.	*int8
//			  27.	int16
//			  28.	*int16
//			  29.	int
//			  30.	*int
//			  31.	int32
//			  32.	*int32
//			  33.	int64
//			  34.	*int64
//			  35.	uint8
//			  36.	*uint8
//			  37.	uint16
//			  38.	*uint16
//			  39.	uint
//			  40.	*uint
//			  41.	uint32
//			  42.	*uint32
//			  43.	uint64,
//			  44.	*uint64
//			  45.	big.Int
//			  46.	*big.Int
//			  47.	TextFieldFormatDtoFloat64
//			  48.	*TextFieldFormatDtoFloat64
//			  49.	TextFieldFormatDtoBigFloat
//			  50.	*TextFieldFormatDtoBigFloat
//			  51.	NumberStrKernel
//			  52.	*NumberStrKernel
//			  53.	[]NumberStrKernel
//			  54.	*[]NumberStrKernel
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

	var stringToWrite string
	var ok bool
	var byteArray []byte
	var strArray []string
	var strArrayDto StringArrayDto
	var numStrKernelArray []NumberStrKernel

	switch charsToConvert.(type) {

	case []byte:

		byteArray, ok = charsToConvert.([]byte)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"ERROR: Input parameter '%v' is invalid!\n"+
				"'%v' was identified as a byte array []byte.\n"+
				"The cast from '%v' to []byte Failed.\n",
				ePrefix.String(),
				charsToConvertLabel,
				charsToConvertLabel,
				charsToConvertLabel)

			return numOfBytesWritten, err
		}

		goto writeToBytes

	case *[]byte:

		var byteArrayPtr *[]byte

		byteArrayPtr, ok = charsToConvert.(*[]byte)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"Input parameter '%v' is ERROR!\n"+
				"'%v' was identified as a byte array ponter (*[]byte).\n"+
				"The cast from '%v' to *[]byte Failed.\n",
				ePrefix.String(),
				charsToConvertLabel,
				charsToConvertLabel,
				charsToConvertLabel)

			return numOfBytesWritten, err
		}

		byteArray = *byteArrayPtr

		goto writeToBytes

	case string:
		// string

		stringToWrite, ok = charsToConvert.(string)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"Input parameter '%v' is ERROR!\n"+
				"'%v' was identified as a string.\n"+
				"The cast from '%v' to string Failed.\n",
				ePrefix.String(),
				charsToConvertLabel,
				charsToConvertLabel,
				charsToConvertLabel)

			return numOfBytesWritten, err
		}

		byteArray = []byte(stringToWrite)

		goto writeToBytes

	case *string:
		// string pointer

		var strPtr *string

		strPtr, ok = charsToConvert.(*string)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"Input parameter '%v' is ERROR!\n"+
				"'%v' was identified as a string pointer.\n"+
				"string cast from '%v' to string pointer Failed.\n",
				ePrefix.String(),
				charsToConvertLabel,
				charsToConvertLabel,
				charsToConvertLabel)

			return numOfBytesWritten, err
		}

		byteArray = []byte(*strPtr)

		goto writeToBytes

	case []string:
		// string array

		strArray, ok = charsToConvert.([]string)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"ERROR: Input parameter '%v' is invalid!\n"+
				"'%v' was identified as a string array.\n"+
				"string array cast from '%v' to string array Failed.\n",
				ePrefix.String(),
				charsToConvertLabel,
				charsToConvertLabel,
				charsToConvertLabel)

			return numOfBytesWritten, err
		}

		goto writeToStrArray

	case StringArrayDto:

		strArrayDto, ok = charsToConvert.(StringArrayDto)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"Input parameter '%v' is ERROR!\n"+
				"'%v' was identified as a StringArrayDto.\n"+
				"The cast from '%v' to StringArrayDto Failed.\n",
				ePrefix.String(),
				charsToConvertLabel,
				charsToConvertLabel,
				charsToConvertLabel)

			return numOfBytesWritten, err
		}

		strArray = strArrayDto.StrArray

		goto writeToStrArrayDto

	case *StringArrayDto:

		var strArrayDtoPtr *StringArrayDto

		strArrayDtoPtr, ok = charsToConvert.(*StringArrayDto)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"----------------------------------------------------------------\n"+
				"ERROR: Input parameter '%v' is invalid!\n"+
				"'%v' was identified as a StringArrayDto Pointer.\n"+
				"The cast from '%v' to *StringArrayDto Failed.\n",
				ePrefix.String(),
				charsToConvertLabel,
				charsToConvertLabel,
				charsToConvertLabel)

			return numOfBytesWritten, err
		}

		strArrayDto = *strArrayDtoPtr

		goto writeToStrArrayDto

	case strings.Builder:

		var strBuilder strings.Builder

		strBuilder, ok = charsToConvert.(strings.Builder)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"ERROR: Input parameter '%v' is invalid!\n"+
				"'%v' was identified as a string builder (strings.Builder).\n"+
				"The cast from '%v' to strings.Builder Failed.\n",
				ePrefix.String(),
				charsToConvertLabel,
				charsToConvertLabel,
				charsToConvertLabel)

			return numOfBytesWritten, err
		}

		byteArray = []byte(strBuilder.String())

		goto writeToBytes

	case *strings.Builder:

		var strBuilderPtr *strings.Builder

		strBuilderPtr, ok = charsToConvert.(*strings.Builder)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------------------------\n"+
				"ERROR: Input parameter '%v' is invalid!\n"+
				"'%v' was identified as a string builder pointer (*strings.Builder).\n"+
				"The cast from '%v' to *strings.Builder Failed.\n",
				ePrefix.String(),
				charsToConvertLabel,
				charsToConvertLabel,
				charsToConvertLabel)

			return numOfBytesWritten, err
		}

		byteArray = []byte(strBuilderPtr.String())

		goto writeToBytes

	case []rune:

		var runesToWrite []rune

		runesToWrite, ok = charsToConvert.([]rune)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"Input parameter '%v' is ERROR!\n"+
				"'%v' was identified as a rune array ([]rune).\n"+
				"The cast from '%v' to []rune Failed.\n",
				ePrefix.String(),
				charsToConvertLabel,
				charsToConvertLabel,
				charsToConvertLabel)

			return numOfBytesWritten, err
		}

		byteArray = []byte(string(runesToWrite))

		goto writeToBytes

	case *[]rune:

		var runeArrayPtr *[]rune

		runeArrayPtr, ok = charsToConvert.(*[]rune)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"Input parameter '%v' is ERROR!\n"+
				"'%v' was identified as a rune array pointer (*[]rune).\n"+
				"The cast from '%v' to *[]rune Failed.\n",
				ePrefix.String(),
				charsToConvertLabel,
				charsToConvertLabel,
				charsToConvertLabel)

			return numOfBytesWritten, err
		}

		byteArray = []byte(string(*runeArrayPtr))

		goto writeToBytes

	case RuneArrayDto:

		var runesToWriteDto RuneArrayDto

		runesToWriteDto, ok = charsToConvert.(RuneArrayDto)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"Input parameter '%v' is ERROR!\n"+
				"'%v' was identified as a type RuneArrayDto.\n"+
				"The cast from '%v' to RuneArrayDto Failed.\n",
				ePrefix.String(),
				charsToConvertLabel,
				charsToConvertLabel,
				charsToConvertLabel)

			return numOfBytesWritten, err
		}

		byteArray = []byte(string(runesToWriteDto.CharsArray))

		goto writeToBytes

	case *RuneArrayDto:

		var runesToWriteDtoPtr *RuneArrayDto

		runesToWriteDtoPtr, ok = charsToConvert.(*RuneArrayDto)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"Input parameter '%v' is ERROR!\n"+
				"'%v' was identified as a type *RuneArrayDto.\n"+
				"The cast from '%v' to *RuneArrayDto Failed.\n",
				ePrefix.String(),
				charsToConvertLabel,
				charsToConvertLabel,
				charsToConvertLabel)

			return numOfBytesWritten, err
		}

		byteArray = []byte(string(runesToWriteDtoPtr.CharsArray))

		goto writeToBytes

	case ITextFieldFormatDto:

		var textFileFormatDto ITextFieldFormatDto

		textFileFormatDto, ok = charsToConvert.(ITextFieldFormatDto)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"Input parameter '%v' is ERROR!\n"+
				"'%v' was identified as a type ITextFieldFormatDto.\n"+
				"The cast from '%v' to ITextFieldFormatDto Failed.\n",
				ePrefix.String(),
				charsToConvertLabel,
				charsToConvertLabel,
				charsToConvertLabel)

			return numOfBytesWritten, err
		}

		stringToWrite,
			err = textFileFormatDto.
			GetFormattedTextFieldStr(
				ePrefix.XCpy("textFileFormatDto"))

		if err != nil {

			return numOfBytesWritten, err

		}

		byteArray = []byte(stringToWrite)

		goto writeToBytes

	case ITextFieldSpecification:

		var textFieldSpecification ITextFieldSpecification

		textFieldSpecification, ok =
			charsToConvert.(ITextFieldSpecification)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"Input parameter '%v' is ERROR!\n"+
				"'%v' was identified as a type ITextFieldSpecification.\n"+
				"The cast from '%v' to ITextFieldSpecification Failed.\n",
				ePrefix.String(),
				charsToConvertLabel,
				charsToConvertLabel,
				charsToConvertLabel)

			return numOfBytesWritten, err
		}

		var fieldSpecStrBuilder strings.Builder

		err = textFieldSpecification.
			TextBuilder(
				&fieldSpecStrBuilder,
				ePrefix.XCpy("ITextFieldSpecification"))

		if err != nil {

			return numOfBytesWritten, err

		}

		byteArray = []byte(fieldSpecStrBuilder.String())

		goto writeToBytes

	case ITextLineSpecification:

		var texLineSpecification ITextLineSpecification

		texLineSpecification, ok =
			charsToConvert.(ITextLineSpecification)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"Input parameter '%v' is ERROR!\n"+
				"'%v' was identified as a type ITextLineSpecification.\n"+
				"The cast from '%v' to ITextLineSpecification Failed.\n",
				ePrefix.String(),
				charsToConvertLabel,
				charsToConvertLabel,
				charsToConvertLabel)

			return numOfBytesWritten, err
		}

		var fieldSpecStrBuilder strings.Builder

		err = texLineSpecification.
			TextBuilder(
				&fieldSpecStrBuilder,
				ePrefix.XCpy("ITextLineSpecification"))

		if err != nil {

			return numOfBytesWritten, err

		}

		byteArray = []byte(fieldSpecStrBuilder.String())

		goto writeToBytes

	case float32, *float32, float64, *float64, *BigFloatDto,
		BigFloatDto, *big.Float, big.Float, big.Rat, *big.Rat,
		int8, *int8, int16, *int16, int, *int, int32,
		*int32, int64, *int64, uint8, *uint8, uint16,
		*uint16, uint, *uint, uint32, *uint32, uint64,
		*uint64, big.Int, *big.Int, TextFieldFormatDtoFloat64,
		*TextFieldFormatDtoFloat64, TextFieldFormatDtoBigFloat,
		*TextFieldFormatDtoBigFloat, NumberStrKernel,
		*NumberStrKernel:

		// Writes numerical data to io.Writer

		stringToWrite,
			err = new(mathHelperNanobot).
			numericValueToNativeNumStr(
				charsToConvert,
				ePrefix.XCpy("<-charsToConvert"))

		if err != nil {

			return numOfBytesWritten,
				fmt.Errorf("%v\n"+
					"Error converting numeric value to a number string!\n"+
					"Error=\n%v\n",
					funcName,
					err.Error())

		}

		byteArray = []byte(stringToWrite)

		goto writeToBytes

	case []NumberStrKernel:

		numStrKernelArray, ok =
			charsToConvert.([]NumberStrKernel)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"ERROR: Input parameter '%v' is invalid!\n"+
				"'%v' was identified as a type []NumberStrKernel.\n"+
				"The cast from '%v' to []NumberStrKernel Failed.\n",
				ePrefix.String(),
				charsToConvertLabel,
				charsToConvertLabel,
				charsToConvertLabel)

			return numOfBytesWritten, err
		}

		goto writeNumStrKernelArray

	case *[]NumberStrKernel:

		var numStrKernelArrayPtr *[]NumberStrKernel

		numStrKernelArrayPtr, ok =
			charsToConvert.(*[]NumberStrKernel)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"ERROR: Input parameter '%v' is invalid!\n"+
				"'%v' was identified as a type *[]NumberStrKernel.\n"+
				"The cast from '%v' to *[]NumberStrKernel Failed.\n",
				ePrefix.String(),
				charsToConvertLabel,
				charsToConvertLabel,
				charsToConvertLabel)

			return numOfBytesWritten, err
		}

		numStrKernelArray = *numStrKernelArrayPtr

		goto writeNumStrKernelArray

	default:

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'charsToConvert' is an invalid type!\n"+
			"'charsToConvert' is unsupported type '%T'\n",
			ePrefix.String(),
			charsToConvert)

		return numOfBytesWritten, err
	}

writeToBytes:

	numOfBytesWritten,
		err = writeBytes(
		byteArray,
		writeEndOfLineChars+writeEndOfTextChars,
		ePrefix.XCpy("byteArray<-[]byte"))

	return numOfBytesWritten, err

writeToStrArray:

	numOfBytesWritten,
		err = new(fileWriterHelperMolecule).
		writeStrArray(
			writeBytes,
			strArray,
			charsToConvertLabel,
			writeEndOfLineChars,
			writeEndOfTextChars,
			ePrefix.XCpy(
				"byteArray<-strArray"))

	return numOfBytesWritten, err

writeToStrArrayDto:

	numOfBytesWritten,
		err = new(fileWriterHelperMolecule).
		writeStrArray(
			writeBytes,
			strArray,
			charsToConvertLabel,
			writeEndOfLineChars,
			writeEndOfTextChars,
			ePrefix.XCpy(
				"byteArray<-strArray"))

	return numOfBytesWritten, err

writeNumStrKernelArray:

	//lastNumStrIdx := len(numStrKernelArray) - 1

	numOfBytesWritten,
		err = new(fileWriterHelperNanobot).
		writeNumStrKernelArray(
			writeBytes,
			numStrKernelArray,
			charsToConvertLabel,
			writeEndOfLineChars,
			writeEndOfTextChars,
			ePrefix.XCpy(
				"byteArray<-numStrKernelArray"))

	return numOfBytesWritten, err
}

type fileWriterHelperNanobot struct {
	lock *sync.Mutex
}

func (fWriterHlprNanobot *fileWriterHelperNanobot) writeNumStrKernelArray(
	writeBytes func(
		bytesToWrite []byte,
		xWriteEndOfTextChars string,
		xErrPref *ePref.ErrPrefixDto) (
		xNumOfBytesWritten int64,
		xErr error),
	numStrKernelArray []NumberStrKernel,
	numStrKernelArrayLabel string,
	writeEndOfLineChars string,
	writeEndOfTextChars string,
	errPrefDto *ePref.ErrPrefixDto) (
	numOfBytesWritten int64,
	err error) {

	if fWriterHlprNanobot.lock == nil {
		fWriterHlprNanobot.lock = new(sync.Mutex)
	}

	fWriterHlprNanobot.lock.Lock()

	defer fWriterHlprNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "fileWriterHelperNanobot." +
		"writeNumStrKernelArray()"

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

	if len(numStrKernelArrayLabel) == 0 {

		numStrKernelArrayLabel = "numStrKernelArray"
	}

	lastNumStrIdx := len(numStrKernelArray) - 1

	if lastNumStrIdx < 0 {

		err = fmt.Errorf("%v\n"+
			"--------------------------------------------------------------\n"+
			"ERROR: Input parameter '%v' is invalid!\n"+
			"'%v' was identified as a type []NumberStrKernel.\n"+
			"However, '%v' is empty and contains zero array elements.\n",
			ePrefix.String(),
			numStrKernelArrayLabel,
			numStrKernelArrayLabel,
			numStrKernelArrayLabel)

		return numOfBytesWritten, err
	}

	var stringToWrite string
	var strArray []string

	for i := 0; i <= lastNumStrIdx; i++ {

		stringToWrite,
			_,
			err = numStrKernelArray[i].FmtNumStrNative(
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy(
				fmt.Sprintf(
					"numStrKernelArray[%v]", i)))

		if err != nil {

			return numOfBytesWritten,
				fmt.Errorf("%v\n"+
					"Error returned by numStrKernelArray[%v].FmtNumStrNative()\n"+
					"Error=\n%v\n",
					funcName,
					i,
					err.Error())

		}

		strArray = append(strArray, stringToWrite)

		stringToWrite = ""

	}

	numOfBytesWritten,
		err = new(fileWriterHelperMolecule).
		writeStrArray(
			writeBytes,
			strArray,
			numStrKernelArrayLabel,
			writeEndOfLineChars,
			writeEndOfTextChars,
			ePrefix.XCpy(
				"byteArray<-strArray"))

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
