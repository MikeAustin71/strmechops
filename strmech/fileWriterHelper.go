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
// parameter 'charsToWrite' must match one of over fifty
// eligible data types.
//
// If 'charsToWrite' is set to an ineligible data type,
// an error will be returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	ioWriter					*io.Writer
//
//		A pointer to an instance of io.Writer. The
//		data type 'charsToWrite' will be converted to
//		a string or byte array and written to this
//		io.Writer object.
//
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
//	charsToWrite				interface{}
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
//	charsToWriteLabel			string
//
//		The name or label associated with input parameter
//		'charsToWrite' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "charsToWrite" will be
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
//		io.Writer object after all other text from
//		'charsToWrite' and 'writeEndOfLineChars' have
//		been processed and written.
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
	ioWriter *IFileWriterEngine,
	ioWriterLabel string,
	charsToWrite interface{},
	charsToWriteLabel string,
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

	if len(ioWriterLabel) == 0 {

		ioWriterLabel = "ioWriter"
	}

	if len(charsToWriteLabel) == 0 {

		charsToWriteLabel = "charsToWrite"
	}

	if ioWriter == nil {

		err = fmt.Errorf("%v\n"+
			"Error: The io.Writer instance passed\n"+
			"as input parameter '%v' is invalid!\n"+
			"'%v' is a 'nil' pointer.\n",
			ePrefix.String(),
			ioWriterLabel,
			ioWriterLabel)

		return numOfBytesWritten, err
	}

	if ioWriter == nil {

		err = fmt.Errorf("%v\n"+
			"------------------------------------------------------------------------------\n"+
			"Error: The io.Writer instance passed\n"+
			"as input parameter '%v' is invalid!\n"+
			"'%v' is a 'nil' pointer.\n",
			ePrefix.String(),
			ioWriterLabel,
			ioWriterLabel)

		return numOfBytesWritten, err
	}

	if charsToWrite == nil {

		err = fmt.Errorf("%v\n"+
			"-------------------------------------------------------\n"+
			"Error: Input parameter '%v' is invalid!\n"+
			"'%v' is a 'nil' value.\n",
			ePrefix.String(),
			charsToWriteLabel,
			charsToWriteLabel)

		return numOfBytesWritten, err
	}

	var stringToWrite string
	var ok bool
	var fWriterHlprAtom = new(fileWriterHelperAtom)

	switch charsToWrite.(type) {

	case []byte:

		var byteArray []byte

		byteArray, ok = charsToWrite.([]byte)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"ERROR: Input parameter '%v' is invalid!\n"+
				"'%v' was identified as a byte array []byte.\n"+
				"The cast from '%v' to []byte Failed.\n",
				ePrefix.String(),
				charsToWriteLabel,
				charsToWriteLabel,
				charsToWriteLabel)

			return numOfBytesWritten, err
		}

		numOfBytesWritten,
			err = fWriterHlprAtom.writeBytes(
			ioWriter,
			"ioWriter",
			byteArray,
			"byteArray",
			writeEndOfTextChars,
			ePrefix.XCpy("byteArray<-[]byte"))

	case *[]byte:

		var byteArrayPtr *[]byte

		byteArrayPtr, ok = charsToWrite.(*[]byte)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"Input parameter '%v' is ERROR!\n"+
				"'%v' was identified as a byte array ponter (*[]byte).\n"+
				"The cast from '%v' to *[]byte Failed.\n",
				ePrefix.String(),
				charsToWriteLabel,
				charsToWriteLabel,
				charsToWriteLabel)

			return numOfBytesWritten, err
		}

		numOfBytesWritten,
			err = fWriterHlprAtom.writeBytes(
			ioWriter,
			"ioWriter",
			*byteArrayPtr,
			"byteArrayPtr",
			writeEndOfLineChars+writeEndOfTextChars,
			ePrefix.XCpy("byteArray<-*[]byte"))

	case string:
		// string

		stringToWrite, ok = charsToWrite.(string)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"Input parameter '%v' is ERROR!\n"+
				"'%v' was identified as a string.\n"+
				"The cast from '%v' to string Failed.\n",
				ePrefix.String(),
				charsToWriteLabel,
				charsToWriteLabel,
				charsToWriteLabel)

			return numOfBytesWritten, err
		}

		numOfBytesWritten,
			err = fWriterHlprAtom.writeBytes(
			ioWriter,
			"ioWriter",
			[]byte(stringToWrite),
			"byteArray",
			writeEndOfLineChars+writeEndOfTextChars,
			ePrefix.XCpy("byteArray<-stringToWrite"))

	case *string:
		// string pointer

		var strPtr *string

		strPtr, ok = charsToWrite.(*string)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"Input parameter '%v' is ERROR!\n"+
				"'%v' was identified as a string pointer.\n"+
				"string cast from '%v' to string pointer Failed.\n",
				ePrefix.String(),
				charsToWriteLabel,
				charsToWriteLabel,
				charsToWriteLabel)

			return numOfBytesWritten, err
		}

		numOfBytesWritten,
			err = fWriterHlprAtom.writeBytes(
			ioWriter,
			"ioWriter",
			[]byte(*strPtr),
			"byteArray",
			writeEndOfLineChars+writeEndOfTextChars,
			ePrefix.XCpy("byteArray<-*string"))

	case []string:
		// string array

		var strArray []string

		strArray, ok = charsToWrite.([]string)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"ERROR: Input parameter '%v' is invalid!\n"+
				"'%v' was identified as a string array.\n"+
				"string array cast from '%v' to string array Failed.\n",
				ePrefix.String(),
				charsToWriteLabel,
				charsToWriteLabel,
				charsToWriteLabel)

			return numOfBytesWritten, err
		}

		numOfBytesWritten,
			err = fWriterHlprAtom.
			writeStringArray(
				ioWriter,
				"ioWriter",
				strArray,
				"strArray",
				writeEndOfLineChars,
				writeEndOfTextChars,
				ePrefix.XCpy("strArray<-[]string"))

	case StringArrayDto:

		var strArrayDto StringArrayDto

		strArrayDto, ok = charsToWrite.(StringArrayDto)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"Input parameter '%v' is ERROR!\n"+
				"'%v' was identified as a StringArrayDto.\n"+
				"The cast from '%v' to StringArrayDto Failed.\n",
				ePrefix.String(),
				charsToWriteLabel,
				charsToWriteLabel,
				charsToWriteLabel)

			return numOfBytesWritten, err
		}

		numOfBytesWritten,
			err = fWriterHlprAtom.
			writeStringArray(
				ioWriter,
				"ioWriter",
				strArrayDto.StrArray,
				"strArrayDto.StrArray",
				writeEndOfLineChars,
				writeEndOfTextChars,
				ePrefix.XCpy("strArray<-StringArrayDto"))

	case *StringArrayDto:

		var strArrayDtoPtr *StringArrayDto

		strArrayDtoPtr, ok = charsToWrite.(*StringArrayDto)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"----------------------------------------------------------------\n"+
				"ERROR: Input parameter '%v' is invalid!\n"+
				"'%v' was identified as a StringArrayDto Pointer.\n"+
				"The cast from '%v' to *StringArrayDto Failed.\n",
				ePrefix.String(),
				charsToWriteLabel,
				charsToWriteLabel,
				charsToWriteLabel)

			return numOfBytesWritten, err
		}

		numOfBytesWritten,
			err = fWriterHlprAtom.
			writeStringArray(
				ioWriter,
				"ioWriter",
				strArrayDtoPtr.StrArray,
				"strArrayDtoPtr.StrArray",
				writeEndOfLineChars,
				writeEndOfTextChars,
				ePrefix.XCpy("strArray<-*StringArrayDto"))

	case strings.Builder:

		var strBuilder strings.Builder

		strBuilder, ok = charsToWrite.(strings.Builder)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"ERROR: Input parameter '%v' is invalid!\n"+
				"'%v' was identified as a string builder (strings.Builder).\n"+
				"The cast from '%v' to strings.Builder Failed.\n",
				ePrefix.String(),
				charsToWriteLabel,
				charsToWriteLabel,
				charsToWriteLabel)

			return numOfBytesWritten, err
		}

		numOfBytesWritten,
			err = fWriterHlprAtom.writeBytes(
			ioWriter,
			"ioWriter",
			[]byte(strBuilder.String()),
			"strBuilder",
			writeEndOfLineChars+writeEndOfTextChars,
			ePrefix.XCpy("byteArray<-strBuilder"))

	case *strings.Builder:

		var strBuilderPtr *strings.Builder

		strBuilderPtr, ok = charsToWrite.(*strings.Builder)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------------------------\n"+
				"ERROR: Input parameter '%v' is invalid!\n"+
				"'%v' was identified as a string builder pointer (*strings.Builder).\n"+
				"The cast from '%v' to *strings.Builder Failed.\n",
				ePrefix.String(),
				charsToWriteLabel,
				charsToWriteLabel,
				charsToWriteLabel)

			return numOfBytesWritten, err
		}

		numOfBytesWritten,
			err = fWriterHlprAtom.writeBytes(
			ioWriter,
			"ioWriter",
			[]byte(strBuilderPtr.String()),
			"strBuilder",
			writeEndOfLineChars+writeEndOfTextChars,
			ePrefix.XCpy("byteArray<-strBuilderPtr"))

	case []rune:

		var runesToWrite []rune

		runesToWrite, ok = charsToWrite.([]rune)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"Input parameter '%v' is ERROR!\n"+
				"'%v' was identified as a rune array ([]rune).\n"+
				"The cast from '%v' to []rune Failed.\n",
				ePrefix.String(),
				charsToWriteLabel,
				charsToWriteLabel,
				charsToWriteLabel)

			return numOfBytesWritten, err
		}

		numOfBytesWritten,
			err = fWriterHlprAtom.writeBytes(
			ioWriter,
			"ioWriter",
			[]byte(string(runesToWrite)),
			"runesToWrite",
			writeEndOfLineChars+writeEndOfTextChars,
			ePrefix.XCpy("byteArray<-runeToWrite"))

	case *[]rune:

		var runeArrayPtr *[]rune

		runeArrayPtr, ok = charsToWrite.(*[]rune)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"Input parameter '%v' is ERROR!\n"+
				"'%v' was identified as a rune array pointer (*[]rune).\n"+
				"The cast from '%v' to *[]rune Failed.\n",
				ePrefix.String(),
				charsToWriteLabel,
				charsToWriteLabel,
				charsToWriteLabel)

			return numOfBytesWritten, err
		}

		numOfBytesWritten,
			err = fWriterHlprAtom.writeBytes(
			ioWriter,
			"ioWriter",
			[]byte(string(*runeArrayPtr)),
			"runeArray",
			writeEndOfLineChars+writeEndOfTextChars,
			ePrefix.XCpy("byteArray<-runeArray"))

	case RuneArrayDto:

		var runesToWriteDto RuneArrayDto

		runesToWriteDto, ok = charsToWrite.(RuneArrayDto)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"Input parameter '%v' is ERROR!\n"+
				"'%v' was identified as a type RuneArrayDto.\n"+
				"The cast from '%v' to RuneArrayDto Failed.\n",
				ePrefix.String(),
				charsToWriteLabel,
				charsToWriteLabel,
				charsToWriteLabel)

			return numOfBytesWritten, err
		}

		numOfBytesWritten,
			err = fWriterHlprAtom.writeBytes(
			ioWriter,
			"ioWriter",
			[]byte(string(runesToWriteDto.CharsArray)),
			"runesToWrite",
			writeEndOfLineChars+writeEndOfTextChars,
			ePrefix.XCpy("byteArray<-runeToWrite"))

	case *RuneArrayDto:

		var runesToWriteDtoPtr *RuneArrayDto

		runesToWriteDtoPtr, ok = charsToWrite.(*RuneArrayDto)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"Input parameter '%v' is ERROR!\n"+
				"'%v' was identified as a type *RuneArrayDto.\n"+
				"The cast from '%v' to *RuneArrayDto Failed.\n",
				ePrefix.String(),
				charsToWriteLabel,
				charsToWriteLabel,
				charsToWriteLabel)

			return numOfBytesWritten, err
		}

		numOfBytesWritten,
			err = fWriterHlprAtom.writeBytes(
			ioWriter,
			"ioWriter",
			[]byte(string(runesToWriteDtoPtr.CharsArray)),
			"runesToWrite",
			writeEndOfLineChars+writeEndOfTextChars,
			ePrefix.XCpy("byteArray<-runeToWrite"))

	case ITextFieldFormatDto:

		var textFileFormatDto ITextFieldFormatDto

		textFileFormatDto, ok = charsToWrite.(ITextFieldFormatDto)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"Input parameter '%v' is ERROR!\n"+
				"'%v' was identified as a type ITextFieldFormatDto.\n"+
				"The cast from '%v' to ITextFieldFormatDto Failed.\n",
				ePrefix.String(),
				charsToWriteLabel,
				charsToWriteLabel,
				charsToWriteLabel)

			return numOfBytesWritten, err
		}

		stringToWrite,
			err = textFileFormatDto.
			GetFormattedTextFieldStr(
				ePrefix.XCpy("textFileFormatDto"))

		if err != nil {

			return numOfBytesWritten, err

		}

		numOfBytesWritten,
			err = fWriterHlprAtom.writeBytes(
			ioWriter,
			"ioWriter",
			[]byte(stringToWrite),
			"runesToWrite",
			writeEndOfLineChars+writeEndOfTextChars,
			ePrefix.XCpy("byteArray<-ITextFieldFormatDto"))

	case ITextFieldSpecification:

		var textFieldSpecification ITextFieldSpecification

		textFieldSpecification, ok =
			charsToWrite.(ITextFieldSpecification)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"Input parameter '%v' is ERROR!\n"+
				"'%v' was identified as a type ITextFieldSpecification.\n"+
				"The cast from '%v' to ITextFieldSpecification Failed.\n",
				ePrefix.String(),
				charsToWriteLabel,
				charsToWriteLabel,
				charsToWriteLabel)

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

		numOfBytesWritten,
			err = fWriterHlprAtom.writeBytes(
			ioWriter,
			"ioWriter",
			[]byte(fieldSpecStrBuilder.String()),
			"ITextFieldSpecification",
			writeEndOfLineChars+writeEndOfTextChars,
			ePrefix.XCpy("byteArray<-ITextFieldSpecification"))

	case ITextLineSpecification:

		var texLineSpecification ITextLineSpecification

		texLineSpecification, ok =
			charsToWrite.(ITextLineSpecification)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"Input parameter '%v' is ERROR!\n"+
				"'%v' was identified as a type ITextLineSpecification.\n"+
				"The cast from '%v' to ITextLineSpecification Failed.\n",
				ePrefix.String(),
				charsToWriteLabel,
				charsToWriteLabel,
				charsToWriteLabel)

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

		numOfBytesWritten,
			err = fWriterHlprAtom.writeBytes(
			ioWriter,
			"ioWriter",
			[]byte(fieldSpecStrBuilder.String()),
			"ITextLineSpecification",
			writeEndOfLineChars+writeEndOfTextChars,
			ePrefix.XCpy("byteArray<-ITextLineSpecification"))

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
				charsToWrite,
				ePrefix.XCpy("<-charsToWrite"))

		if err != nil {

			return numOfBytesWritten,
				fmt.Errorf("%v\n"+
					"Error converting numeric value to a number string!\n"+
					"Error=\n%v\n",
					funcName,
					err.Error())

		}

		numOfBytesWritten,
			err = fWriterHlprAtom.writeBytes(
			ioWriter,
			"ioWriter",
			[]byte(stringToWrite),
			"Number String",
			writeEndOfLineChars+writeEndOfTextChars,
			ePrefix.XCpy("byteArray<-number string"))

	case []NumberStrKernel:

		var numStrKernelArray []NumberStrKernel

		numStrKernelArray, ok =
			charsToWrite.([]NumberStrKernel)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"ERROR: Input parameter '%v' is invalid!\n"+
				"'%v' was identified as a type []NumberStrKernel.\n"+
				"The cast from '%v' to []NumberStrKernel Failed.\n",
				ePrefix.String(),
				charsToWriteLabel,
				charsToWriteLabel,
				charsToWriteLabel)

			return numOfBytesWritten, err
		}

		lastNumStrIdx := len(numStrKernelArray) - 1

		if lastNumStrIdx < 0 {

			err = fmt.Errorf("%v\n"+
				"--------------------------------------------------------------\n"+
				"ERROR: Input parameter '%v' is invalid!\n"+
				"'%v' was identified as a type []NumberStrKernel.\n"+
				"However, '%v' is empty and contains zero array elements.\n",
				ePrefix.String(),
				charsToWriteLabel,
				charsToWriteLabel,
				charsToWriteLabel)

			return numOfBytesWritten, err
		}

		var localNumBytesWritten int64

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

			stringToWrite += writeEndOfLineChars

			if i == lastNumStrIdx {

				stringToWrite += writeEndOfTextChars
			}

			localNumBytesWritten,
				err = fWriterHlprAtom.writeBytes(
				ioWriter,
				"ioWriter",
				[]byte(stringToWrite),
				"byteArray",
				writeEndOfTextChars,
				ePrefix.XCpy(
					fmt.Sprintf(
						"byteArray<-numStrKernelArray[%v]", i)))

			if err != nil {

				return numOfBytesWritten, err
			}

			numOfBytesWritten += localNumBytesWritten
		}

	case *[]NumberStrKernel:

		var numStrKernelArrayPtr *[]NumberStrKernel

		numStrKernelArrayPtr, ok =
			charsToWrite.(*[]NumberStrKernel)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"ERROR: Input parameter '%v' is invalid!\n"+
				"'%v' was identified as a type *[]NumberStrKernel.\n"+
				"The cast from '%v' to *[]NumberStrKernel Failed.\n",
				ePrefix.String(),
				charsToWriteLabel,
				charsToWriteLabel,
				charsToWriteLabel)

			return numOfBytesWritten, err
		}

		var numStrKernelArray []NumberStrKernel

		numStrKernelArray = *numStrKernelArrayPtr

		lastNumStrIdx := len(numStrKernelArray) - 1

		if lastNumStrIdx < 0 {

			err = fmt.Errorf("%v\n"+
				"--------------------------------------------------------------\n"+
				"ERROR: Input parameter '%v' is invalid!\n"+
				"'%v' was identified as a type *[]NumberStrKernel.\n"+
				"However, '%v' is empty and contains zero array elements.\n",
				ePrefix.String(),
				charsToWriteLabel,
				charsToWriteLabel,
				charsToWriteLabel)

			return numOfBytesWritten, err
		}

		var localNumBytesWritten int64

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

			stringToWrite += writeEndOfLineChars

			if i == lastNumStrIdx {

				stringToWrite += writeEndOfTextChars
			}

			localNumBytesWritten,
				err = fWriterHlprAtom.writeBytes(
				ioWriter,
				"ioWriter",
				[]byte(stringToWrite),
				"byteArray",
				writeEndOfTextChars,
				ePrefix.XCpy(
					fmt.Sprintf(
						"byteArray<-*numStrKernelArray[%v]", i)))

			if err != nil {

				return numOfBytesWritten, err
			}

			numOfBytesWritten += localNumBytesWritten
		}

	default:

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'charsToWrite' is an invalid type!\n"+
			"'charsToWrite' is unsupported type '%T'\n",
			ePrefix.String(),
			charsToWrite)

		return numOfBytesWritten, err
	}

	return numOfBytesWritten, err
}

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
//	(1)	If the byte array passed as input parameter
//		'byteArray' is empty or contains zero array
//		elements, this method will take no action, no
//		error will be returned and the returned number of
//		bytes written ('numOfBytesWritten') will be set
//		to zero.
//
//	(2)	If the planned number of bytes to be written to
//		the io.Writer object does NOT match the actual
//		number of bytes written to the io.Writer object,
//		an error will be returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	ioWriter 					*IFileWriterEngine
//
//
//		A pointer to an object which implements the
//		IFileWriterEngine interface.
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
//	writeEndOfTextChars 		string
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
//		If the planned number of bytes to be written to
//		the io.Writer object does NOT match the actual
//		number of bytes written to the io.Writer object,
//		an error will be returned.
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
	ioWriter *IFileWriterEngine,
	ioWriterLabel string,
	byteArray []byte,
	byteArrayLabel string,
	writeEndOfTextChars string,
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
	var expectedNumBytesWritten int64

	expectedNumBytesWritten = int64(lenByteArray)

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

	if len(writeEndOfTextChars) > 0 {

		expectedNumBytesWritten +=
			int64(len(writeEndOfTextChars))

		localNumBytesWritten,
			err2 = writer.Write(
			[]byte(writeEndOfTextChars))

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

	if err == nil &&
		expectedNumBytesWritten != numOfBytesWritten {

		err = fmt.Errorf("%v\n"+
			"Error condition detected!\n"+
			"The expected number of bytes to be written does\n"+
			"NOT match the actual number of bytes written.\n"+
			"Expected Number of Bytes to be Written: %v\n"+
			"        Actual Number of Bytes Written: %v\n",
			ePrefix.String(),
			expectedNumBytesWritten,
			numOfBytesWritten)
	}

	return numOfBytesWritten, err
}

// writeStringArray
//
// Writes a string array to the io.Writer object
// passed as input parameter 'ioWriter'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	If the string array passed as input parameter
//		'strArray' is empty or contains zero array
//		elements, this method will take no action, no
//		error will be returned and the returned number of
//		bytes written ('numOfBytesWritten') will be set
//		to zero.
//
//	(2)	If the planned number of bytes to be written to
//		the io.Writer object does NOT match the actual
//		number of bytes written to the io.Writer object,
//		an error will be returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	ioWriter 					*IFileWriterEngine
//
//		A pointer to an object which implements the
//		IFileWriterEngine interface.
//
//		All the strings contained in the string array
//		passed as input parameter 'strArray' will be
//		written to this io.Writer object.
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
//	writeEndOfLineChars 		string
//
//		If this parameter has a string length greater
//		than zero, this string will be appended to
//		each string array element ('strArray') written
//		to the io.Writer object ('ioWriter').
//
//	writeEndOfTextChars			string
//
//		If this parameter has a string length greater
//		than zero, this string will be the last item
//		written to the io.Writer object ('ioWriter').
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
//		If the planned number of bytes to be written to
//		the io.Writer object does NOT match the actual
//		number of bytes written to the io.Writer object,
//		an error will be returned.
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
func (fWriterHelperAtom *fileWriterHelperAtom) writeStringArray(
	ioWriter *IFileWriterEngine,
	ioWriterLabel string,
	strArray []string,
	strArrayLabel string,
	writeEndOfLineChars string,
	writeEndOfTextChars string,
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
		"writeStringArray()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		return numOfBytesWritten, err
	}

	if len(ioWriterLabel) == 0 {

		ioWriterLabel = "ioWriter"
	}

	if len(strArrayLabel) == 0 {

		strArrayLabel = "strArray"
	}

	if ioWriter == nil {

		err = fmt.Errorf("%v\n"+
			"-------------------------------------------\n"+
			"Error: The IFileWriterEngine instance passed\n"+
			"as input parameter '%v' is invalid!\n"+
			"'%v' is a 'nil' pointer.\n",
			ePrefix.String(),
			ioWriterLabel,
			ioWriterLabel)

		return numOfBytesWritten, err
	}

	if ioWriter == nil {

		err = fmt.Errorf("%v\n"+
			"-------------------------------------------------------\n"+
			"Error: The IFileWriterEngine instance passed\n"+
			"as input parameter '%v' is invalid!\n"+
			"'%v' is a 'nil' pointer.\n\n",
			ePrefix.String(),
			ioWriterLabel,
			ioWriterLabel)

		return numOfBytesWritten, err
	}

	lastStrArrayIdx := len(strArray) - 1

	if lastStrArrayIdx < 0 {

		// String array is empty
		return numOfBytesWritten, err
	}

	lenWriteEndOfLineChars := len(writeEndOfLineChars)

	var err2 error
	var localNumBytesWritten int
	var expectedNumBytesWritten int64
	var writer = *ioWriter
	var strToWrite string

	for i := 0; i <= lastStrArrayIdx; i++ {

		strToWrite = strArray[i]

		if lenWriteEndOfLineChars > 0 {
			strToWrite += writeEndOfLineChars
		}

		if i == lastStrArrayIdx {
			strToWrite += writeEndOfTextChars
		}

		expectedNumBytesWritten += int64(len(strToWrite))

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

			break

		} else {

			numOfBytesWritten += int64(localNumBytesWritten)
		}

		strToWrite = ""
	}

	if err == nil &&
		expectedNumBytesWritten != numOfBytesWritten {

		err = fmt.Errorf("%v\n"+
			"Error condition detected!\n"+
			"The expected number of bytes to be written does\n"+
			"NOT match the actual number of bytes written.\n"+
			"Expected Number of Bytes to be Written: %v\n"+
			"        Actual Number of Bytes Written: %v\n",
			ePrefix.String(),
			expectedNumBytesWritten,
			numOfBytesWritten)
	}

	return numOfBytesWritten, err
}
