package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"strings"
	"sync"
	"time"
)

type textSpecificationAtom struct {
	lock *sync.Mutex
}

// convertParamsToNativeElements
//
// Receives an empty interface and attempts to convert that
// interface to one of over fifty eligible before return the
// contents as one of two return types:
//
//	A string array
//		Or
//	A byte array
//
//				Eligible Data Types
//
//			   1.	[]byte
//			   2.	*[]byte
//			   3.	string
//			   4.	*string
//			   5.	[]string
//			   6.	*[]string
//			   7.	strings.Builder
//			   8.	*strings.Builder
//		 	   9.	StringArrayDto
//			  10.	*StringArrayDto
//			  11.	[]rune
//			  12.	*[]rune
//			  13.	RuneArrayDto
//			  14.	*RuneArrayDto
//			  15.	ITextFieldFormatDto
//			  16.	ITextFieldSpecification
//			  17.	ITextLineSpecification
//			  18.	float32
//			  19.	*float32
//			  20.	float64
//			  21.	*float64
//			  22.	BigFloatDto
//			  23.	*BigFloatDto
//			  24.	big.Float
//			  25.	*big.Float
//			  26.	big.Rat
//			  27.	*big.Rat
//			  28.	int8
//			  29.	*int8
//			  30.	int16
//			  31.	*int16
//			  32.	int
//			  33.	*int
//			  34.	int32
//			  35.	*int32
//			  36.	int64
//			  37.	*int64
//			  38.	uint8
//			  39.	*uint8
//			  40.	uint16
//			  41.	*uint16
//			  42.	uint
//			  43.	*uint
//			  44.	uint32
//			  45.	*uint32
//			  46.	uint64,
//			  47.	*uint64
//			  48.	big.Int
//			  49.	*big.Int
//			  50.	TextFieldFormatDtoFloat64
//			  51.	*TextFieldFormatDtoFloat64
//			  52.	TextFieldFormatDtoBigFloat
//			  53.	*TextFieldFormatDtoBigFloat
//			  54.	NumberStrKernel
//			  55.	*NumberStrKernel
//			  56.	[]NumberStrKernel
//			  57.	*[]NumberStrKernel
func (txtSpecAtom *textSpecificationAtom) convertParamsToNativeElements(
	charsToConvert interface{},
	charsToConvertLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	strResult string,
	lenOfStrResult int,
	strArrayDto StringArrayDto,
	numOfStringArrayElements int,
	byteArray []byte,
	numOfByteArrayElements int,
	err error) {

	if txtSpecAtom.lock == nil {
		txtSpecAtom.lock = new(sync.Mutex)
	}

	txtSpecAtom.lock.Lock()

	defer txtSpecAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textSpecificationAtom."+
			"convertToStringArrayOrByteArray()",
		"")

	if err != nil {

		return strResult,
			lenOfStrResult,
			strArrayDto,
			numOfStringArrayElements,
			byteArray,
			numOfByteArrayElements,
			err
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

		return strResult,
			lenOfStrResult,
			strArrayDto,
			numOfStringArrayElements,
			byteArray,
			numOfByteArrayElements,
			err
	}

	var convertedString string
	var ok bool
	var err2 error

	switch charsToConvert.(type) {

	case []byte:

		byteArray, ok = charsToConvert.([]byte)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"ERROR: Input parameter '%v' is invalid!\n"+
				"'%v' was identified as a byte array []byte.\n"+
				"However, the cast from '%v' to []byte Failed.\n",
				ePrefix.String(),
				charsToConvertLabel,
				charsToConvertLabel,
				charsToConvertLabel)

			return strResult,
				lenOfStrResult,
				strArrayDto,
				numOfStringArrayElements,
				byteArray,
				numOfByteArrayElements,
				err
		}

		numOfByteArrayElements = len(byteArray)

		if numOfByteArrayElements == 0 {

			err = fmt.Errorf("%v\n"+
				"Error: %v converted to a byte array.\n"+
				"However, the byte array is empty and\n"+
				"has a zero length!\n",
				ePrefix.String(),
				charsToConvertLabel)
		}

	case *[]byte:

		var byteArrayPtr *[]byte

		byteArrayPtr, ok = charsToConvert.(*[]byte)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"Error: Input parameter '%v' is invalid!\n"+
				"'%v' was identified as a byte array ponter (*[]byte).\n"+
				"However, the cast from '%v' to *[]byte Failed.\n",
				ePrefix.String(),
				charsToConvertLabel,
				charsToConvertLabel,
				charsToConvertLabel)

			return strResult,
				lenOfStrResult,
				strArrayDto,
				numOfStringArrayElements,
				byteArray,
				numOfByteArrayElements,
				err
		}

		byteArray = *byteArrayPtr

		numOfByteArrayElements = len(byteArray)

		if numOfByteArrayElements == 0 {
			err = fmt.Errorf("%v\n"+
				"Error: %v converted to a byte array pointer\n"+
				"(*[]byte). However, the byte array is empty\n"+
				"and has a zero length!\n",
				ePrefix.String(),
				charsToConvertLabel)
		}

	case string:

		convertedString, ok = charsToConvert.(string)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"Error: Input parameter '%v' is invalid!\n"+
				"'%v' was identified as a string.\n"+
				"However, the cast from '%v' to string Failed.\n",
				ePrefix.String(),
				charsToConvertLabel,
				charsToConvertLabel,
				charsToConvertLabel)

			return strResult,
				lenOfStrResult,
				strArrayDto,
				numOfStringArrayElements,
				byteArray,
				numOfByteArrayElements,
				err
		}

		if len(convertedString) == 0 {

			err = fmt.Errorf("%v\n"+
				"Error: %v converted to a string.\n"+
				"However, that string is empty\n"+
				"and has a zero length!\n",
				ePrefix.String(),
				charsToConvertLabel)

			return strResult,
				lenOfStrResult,
				strArrayDto,
				numOfStringArrayElements,
				byteArray,
				numOfByteArrayElements,
				err
		}

		strResult = convertedString

		lenOfStrResult = len(strResult)

	case *string:
		// string pointer

		var strPtr *string

		strPtr, ok = charsToConvert.(*string)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"Error: Input parameter '%v' is invalid!\n"+
				"'%v' was identified as a string pointer.\n"+
				"However, the cast from '%v' to string pointer Failed.\n",
				ePrefix.String(),
				charsToConvertLabel,
				charsToConvertLabel,
				charsToConvertLabel)

			return strResult,
				lenOfStrResult,
				strArrayDto,
				numOfStringArrayElements,
				byteArray,
				numOfByteArrayElements,
				err
		}

		strResult = *strPtr

		lenOfStrResult = len(strResult)

		if lenOfStrResult == 0 {

			err = fmt.Errorf("%v\n"+
				"Error: %v converted to a string pointer\n"+
				"(*string). However, that string is empty\n"+
				"and has a zero length!\n",
				ePrefix.String(),
				charsToConvertLabel)

			return strResult,
				lenOfStrResult,
				strArrayDto,
				numOfStringArrayElements,
				byteArray,
				numOfByteArrayElements,
				err
		}

	case []string:
		// string array

		strArrayDto.StrArray, ok = charsToConvert.([]string)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"ERROR: Input parameter '%v' is invalid!\n"+
				"'%v' was identified as a string array.\n"+
				"However, the cast from '%v' to string array Failed.\n",
				ePrefix.String(),
				charsToConvertLabel,
				charsToConvertLabel,
				charsToConvertLabel)

			return strResult,
				lenOfStrResult,
				strArrayDto,
				numOfStringArrayElements,
				byteArray,
				numOfByteArrayElements,
				err
		}

		numOfStringArrayElements =
			strArrayDto.GetStringArrayLength()

		if numOfStringArrayElements == 0 {

			err = fmt.Errorf("%v\n"+
				"Error: %v converted to a string array\n"+
				"([]string). However, that string array\n"+
				"is empty and has a zero length!\n",
				ePrefix.String(),
				charsToConvertLabel)

			return strResult,
				lenOfStrResult,
				strArrayDto,
				numOfStringArrayElements,
				byteArray,
				numOfByteArrayElements,
				err
		}

	case *[]string:
		// string array

		var strArrayPtr *[]string

		strArrayPtr, ok = charsToConvert.(*[]string)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"ERROR: Input parameter '%v' is invalid!\n"+
				"'%v' was identified as a string array pointer\n"+
				"(*[]string). However, the cast from '%v' to string\n"+
				"array Failed.\n",
				ePrefix.String(),
				charsToConvertLabel,
				charsToConvertLabel,
				charsToConvertLabel)

			return strResult,
				lenOfStrResult,
				strArrayDto,
				numOfStringArrayElements,
				byteArray,
				numOfByteArrayElements,
				err
		}

		if len(*strArrayPtr) == 0 {

			err = fmt.Errorf("%v\n"+
				"Error: %v converted to a string array pointer\n"+
				"(*[]string). However, that string array is\n"+
				"empty and has a zero length!\n",
				ePrefix.String(),
				charsToConvertLabel)

			return strResult,
				lenOfStrResult,
				strArrayDto,
				numOfStringArrayElements,
				byteArray,
				numOfByteArrayElements,
				err
		}

		strArrayDto = new(StringArrayDto).NewStringArray(
			*strArrayPtr,
			"",
			"")

		numOfStringArrayElements =
			strArrayDto.GetStringArrayLength()

	case strings.Builder:

		var strBuilder strings.Builder

		strBuilder, ok = charsToConvert.(strings.Builder)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"ERROR: Input parameter '%v' is invalid!\n"+
				"'%v' was identified as a string builder\n"+
				"(strings.Builder). However, the cast from '%v' to\n"+
				"strings.Builder Failed.\n",
				ePrefix.String(),
				charsToConvertLabel,
				charsToConvertLabel,
				charsToConvertLabel)

			return strResult,
				lenOfStrResult,
				strArrayDto,
				numOfStringArrayElements,
				byteArray,
				numOfByteArrayElements,
				err
		}

		strResult = strBuilder.String()

		lenOfStrResult = len(strResult)

	case *strings.Builder:

		var strBuilderPtr *strings.Builder

		strBuilderPtr, ok = charsToConvert.(*strings.Builder)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"ERROR: Input parameter '%v' is invalid!\n"+
				"'%v' was identified as a string builder pointer\n"+
				"(*strings.Builder). However, the cast from '%v' to\n"+
				"*strings.Builder Failed.\n",
				ePrefix.String(),
				charsToConvertLabel,
				charsToConvertLabel,
				charsToConvertLabel)

			return strResult,
				lenOfStrResult,
				strArrayDto,
				numOfStringArrayElements,
				byteArray,
				numOfByteArrayElements,
				err
		}

		strResult = strBuilderPtr.String()

		lenOfStrResult = len(strResult)

	case StringArrayDto:

		strArrayDto, ok = charsToConvert.(StringArrayDto)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"Error: Input parameter '%v' is invalid!\n"+
				"'%v' was identified as a StringArrayDto.\n"+
				"However, the cast from '%v' to StringArrayDto Failed.\n",
				ePrefix.String(),
				charsToConvertLabel,
				charsToConvertLabel,
				charsToConvertLabel)

			return strResult,
				lenOfStrResult,
				strArrayDto,
				numOfStringArrayElements,
				byteArray,
				numOfByteArrayElements,
				err
		}

		numOfStringArrayElements =
			strArrayDto.GetStringArrayLength()

	case *StringArrayDto:

		var strArrayDtoPtr *StringArrayDto

		strArrayDtoPtr, ok = charsToConvert.(*StringArrayDto)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"----------------------------------------------------------------\n"+
				"ERROR: Input parameter '%v' is invalid!\n"+
				"'%v' was identified as a StringArrayDto Pointer\n"+
				"(*StringArrayDto). However, the cast from '%v'\n"+
				"to *StringArrayDto Failed.\n",
				ePrefix.String(),
				charsToConvertLabel,
				charsToConvertLabel,
				charsToConvertLabel)

			return strResult,
				lenOfStrResult,
				strArrayDto,
				numOfStringArrayElements,
				byteArray,
				numOfByteArrayElements,
				err
		}

		err2 = strArrayDto.
			CopyIn(
				strArrayDtoPtr,
				ePrefix.XCpy("strArrayDto<-strArrayDtoPtr"))

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Error: strArrayDto.CopyIn(strArrayDtoPtr)\n"+
				"An error occurred while copying from 'strArrayDtoPtr'\n"+
				"to output parameter 'strArrayDto'\n",
				ePrefix.String())

			return strResult,
				lenOfStrResult,
				strArrayDto,
				numOfStringArrayElements,
				byteArray,
				numOfByteArrayElements,
				err

		}

		numOfStringArrayElements =
			strArrayDto.GetStringArrayLength()

	case []rune:

		var runesToWrite []rune

		runesToWrite, ok = charsToConvert.([]rune)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"Error: Input parameter '%v' is invalid!\n"+
				"'%v' was identified as a rune array ([]rune).\n"+
				"However, the cast from '%v' to []rune Failed.\n",
				ePrefix.String(),
				charsToConvertLabel,
				charsToConvertLabel,
				charsToConvertLabel)

			return strResult,
				lenOfStrResult,
				strArrayDto,
				numOfStringArrayElements,
				byteArray,
				numOfByteArrayElements,
				err
		}

		strResult = string(runesToWrite)

		lenOfStrResult = len(strResult)

	case *[]rune:

		var runeArrayPtr *[]rune

		runeArrayPtr, ok = charsToConvert.(*[]rune)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"Error: Input parameter '%v' is invalid!\n"+
				"'%v' was identified as a rune array pointer\n"+
				"(*[]rune). However, the cast from '%v' to\n"+
				"*[]rune Failed.\n",
				ePrefix.String(),
				charsToConvertLabel,
				charsToConvertLabel,
				charsToConvertLabel)

			return strResult,
				lenOfStrResult,
				strArrayDto,
				numOfStringArrayElements,
				byteArray,
				numOfByteArrayElements,
				err
		}

		strResult = string(*runeArrayPtr)

		lenOfStrResult = len(strResult)

	case RuneArrayDto:

		var runesArrayDto RuneArrayDto

		runesArrayDto, ok = charsToConvert.(RuneArrayDto)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"Error: Input parameter '%v' is invalid!\n"+
				"'%v' was identified as a type RuneArrayDto.\n"+
				"However, the cast from '%v' to\n"+
				"RuneArrayDto Failed.\n",
				ePrefix.String(),
				charsToConvertLabel,
				charsToConvertLabel,
				charsToConvertLabel)

			return strResult,
				lenOfStrResult,
				strArrayDto,
				numOfStringArrayElements,
				byteArray,
				numOfByteArrayElements,
				err
		}

		strResult = string(runesArrayDto.CharsArray)

		lenOfStrResult = len(strResult)

	case *RuneArrayDto:

		var runeArrayDtoPtr *RuneArrayDto

		runeArrayDtoPtr, ok = charsToConvert.(*RuneArrayDto)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"Error: Input parameter '%v' is invalid!\n"+
				"'%v' was identified as a type RuneArrayDto pointer\n"+
				"(*RuneArrayDto). However, the cast from '%v'\n"+
				"to *RuneArrayDto Failed.\n",
				ePrefix.String(),
				charsToConvertLabel,
				charsToConvertLabel,
				charsToConvertLabel)

			return strResult,
				lenOfStrResult,
				strArrayDto,
				numOfStringArrayElements,
				byteArray,
				numOfByteArrayElements,
				err
		}

		strResult = string(runeArrayDtoPtr.CharsArray)

		lenOfStrResult = len(strResult)

	}

	return strResult,
		lenOfStrResult,
		strArrayDto,
		numOfStringArrayElements,
		byteArray,
		numOfByteArrayElements,
		err

}

// convertParamEmptyInterfaceToString
//
// Receives an object styled as an empty interface and
// attempts to convert that object to a string.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	emptyIFace                 interface{}
//		This object will be converted to a type of string and
//		returned to the calling function.
//
//		This parameter is an empty interface which must contain
//		of several specific types. This empty interface type will
//		be converted to a string and configured as the single text
//		field in this 'Line1Column' Text Line.
//
//		Supported types which may be submitted through this empty
//		interface parameter are listed as follows:
//
//			time.Time (Converted using default format)
//			NumberStrKernel, *NumberStrKernel
//			string
//			bool
//			uint, uint8, uint16, uint32, uint64,
//			int, int8, int16, int32, int64
//			float32, float64
//			*big.Int, big.Int
//			*big.Float, big.Float
//			*big.Rat, big.Rat
//			fmt.Stringer (types that support this interface)
//			TextInputParamFieldDateTimeDto
//	               (Converts date time to string)
//			ITextLineSpecification
//			ITextFieldSpecification
//				TextFieldFormatDtoBigFloat - Formats big.Float numbers
//				TextFieldFormatDtoDate
//				TextFieldFormatDtoLabel
//				TextFieldFormatDtoFiller
//
//		If the 'emptyIFace' object is not convertible to
//		one of the supported types, an error will be
//		returned.
//
//	 emptyIFaceParamName        string
//	    - This is the name or text label used to describe input
//	      parameter 'emptyIFace' when formatting informational or
//	      error messages. If this parameter is submitted as an empty
//	      string its value will be defaulted to 'emptyIFace'.
//
//
//	 errPrefDto                 *ePref.ErrPrefixDto
//	    - This object encapsulates an error prefix string which is
//	      included in all returned error messages. Usually, it
//	      contains the name of the calling method or methods listed
//	      as a function chain.
//
//	      If no error prefix information is needed, set this parameter
//	      to 'nil'.
//
//	      Type ErrPrefixDto is included in the 'errpref' software
//	      package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	convertedString            string
//	   - If this method completes successfully, this parameter will
//	     be populated with a string value extracted and converted
//	     from the empty interface input parameter, 'emptyIFace'.
//
//
//	err                        error
//	   - If this method completes successfully and no errors are
//	     encountered, this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (txtSpecAtom *textSpecificationAtom) convertParamEmptyInterfaceToString(
	emptyIFace interface{},
	emptyIFaceParamName string,
	errPrefDto *ePref.ErrPrefixDto) (
	convertedString string,
	err error) {

	if txtSpecAtom.lock == nil {
		txtSpecAtom.lock = new(sync.Mutex)
	}

	txtSpecAtom.lock.Lock()

	defer txtSpecAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var ok bool
	var plainString, defaultDateTimeFormat string
	var iStringer fmt.Stringer
	var dateTimeInputDto TextInputParamFieldDateTimeDto
	var dateTimeValue time.Time
	var bFloat big.Float
	var bFloatPtr *big.Float
	var bInt big.Int
	var bIntPtr *big.Int

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textSpecificationAtom."+
			"convertEmptyInterfaceToString()",
		"")

	if err != nil {
		return convertedString, err
	}

	if len(emptyIFaceParamName) == 0 {
		emptyIFaceParamName = "emptyIFace"
	}

	if emptyIFace == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'emptyIFace' is INVALID!\n"+
			"'emptyIFace' is a nil pointer.\n",
			ePrefix.String())

		return convertedString, err
	}

	defaultDateTimeFormat =
		new(textSpecificationMolecule).getDefaultDateTimeFormat()

	switch emptyIFace.(type) { // the switch uses the type of the interface

	case time.Time:

		dateTimeValue,
			ok = emptyIFace.(time.Time)

		if !ok {
			err = fmt.Errorf("%v\n"+
				"Error: Failed to convert empty interface\n"+
				"(%v) to Date Time Value!\n"+
				"String Conversion Error.\n",
				ePrefix.String(),
				emptyIFaceParamName)

			return convertedString, err
		}

		convertedString =
			dateTimeValue.Format(defaultDateTimeFormat)

		return convertedString, err

	case TextInputParamFieldDateTimeDto:

		goto dateTimeConversion

	case bool:

		goto standardConversion

	case uint, uint8, uint16, uint32, uint64:

		goto standardConversion

	case int, int8, int16, int32, int64:

		goto standardConversion

	case float32, float64:

		goto standardConversion

	case big.Int:

		bInt,
			ok = emptyIFace.(big.Int)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"Error: Failed to convert empty interface\n"+
				"to big.Int value!\n"+
				"String Conversion Error.\n",
				ePrefix.String())

			return convertedString, err
		}

		convertedString = bInt.Text(10)

		return convertedString, err

	case *big.Int:

		bIntPtr,
			ok = emptyIFace.(*big.Int)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"Error: Failed to convert empty interface\n"+
				"to *big.Int value!\n"+
				"String Conversion Error.\n",
				ePrefix.String())

			return convertedString, err
		}

		convertedString = bIntPtr.Text(10)

		return convertedString, err

	case big.Float:

		bFloat,
			ok = emptyIFace.(big.Float)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"Error: Failed to convert empty interface\n"+
				"to big.Float value!\n"+
				"String Conversion Error.\n",
				ePrefix.String())

			return convertedString, err
		}

		convertedString = bFloat.Text('f', -1)

		return convertedString, err

	case *big.Float:

		bFloatPtr,
			ok = emptyIFace.(*big.Float)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"Error: Failed to convert empty interface\n"+
				"to *big.Float value!\n"+
				"String Conversion Error.\n",
				ePrefix.String())

			return convertedString, err
		}

		convertedString = bFloatPtr.Text('f', -1)

		return convertedString, err

	case big.Rat:

		var bRat big.Rat

		bRat,
			ok = emptyIFace.(big.Rat)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"Error: Failed to convert empty interface\n"+
				"to big.Rat value!\n"+
				"String Conversion Error.\n",
				ePrefix.String())

			return convertedString, err
		}

		convertedString = bRat.String()

		return convertedString, err

	case *big.Rat:

		var bRatPtr *big.Rat

		bRatPtr,
			ok = emptyIFace.(*big.Rat)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"Error: Failed to convert empty interface\n"+
				"to *big.Rat value!\n"+
				"String Conversion Error.\n",
				ePrefix.String())

			return convertedString, err
		}

		convertedString = bRatPtr.String()

		return convertedString, err

	case NumberStrKernel:

		var nStrKernel NumberStrKernel

		nStrKernel,
			ok = emptyIFace.(NumberStrKernel)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"Error: Failed to convert empty interface\n"+
				"to NumberStrKernel value!\n"+
				"String Conversion Error.\n",
				ePrefix.String())

			return convertedString, err
		}

		convertedString,
			err = nStrKernel.FmtNumStrDefault(
			ePrefix.XCpy(
				"convertedString<-nStrKernel"))

		return convertedString, err

	case *NumberStrKernel:

		var nStrKernelPtr *NumberStrKernel

		nStrKernelPtr,
			ok = emptyIFace.(*NumberStrKernel)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"Error: Failed to convert empty interface\n"+
				"to *NumberStrKernel value!\n"+
				"String Conversion Error.\n",
				ePrefix.String())

			return convertedString, err
		}

		convertedString,
			err = nStrKernelPtr.FmtNumStrDefault(
			ePrefix.XCpy(
				"convertedString<-nStrKernel"))

		return convertedString, err

	case fmt.Stringer:

		goto stringerConversion

	case string:
		goto straightStringConversion

	case ITextFieldSpecification:

		iTextFieldSpec,
			ok := emptyIFace.(ITextFieldSpecification)

		if !ok {
			err = fmt.Errorf("%v\n"+
				"Error: Failed to convert empty interface\n"+
				"(%v) ITextFieldSpecification!\n"+
				"String Conversion Error.\n",
				ePrefix.String(),
				emptyIFaceParamName)

			return convertedString, err
		}

		convertedString,
			err = iTextFieldSpec.GetFormattedText(
			ePrefix.XCpy(
				"convertedString<-iTextFieldSpec"))

		return convertedString, err

	case ITextLineSpecification:

		iTextLineSpec,
			ok := emptyIFace.(ITextLineSpecification)

		if !ok {
			err = fmt.Errorf("%v\n"+
				"Error: Failed to convert empty interface\n"+
				"(%v) ITextLineSpecification!\n"+
				"String Conversion Error.\n",
				ePrefix.String(),
				emptyIFaceParamName)

			return convertedString, err
		}

		convertedString,
			err = iTextLineSpec.GetFormattedText(
			ePrefix.XCpy(
				"convertedString<-iTextLineSpec"))

		return convertedString, err

	case ITextFieldFormatDto:

		iTextFieldFormatDto,
			ok := emptyIFace.(ITextFieldFormatDto)

		if !ok {
			err = fmt.Errorf("%v\n"+
				"Error: Failed to convert empty interface\n"+
				"(%v) ITextFieldFormatDto!\n"+
				"String Conversion Error.\n",
				ePrefix.String(),
				emptyIFaceParamName)

			return convertedString, err
		}

		convertedString,
			err = iTextFieldFormatDto.GetFormattedTextFieldStr(
			ePrefix.XCpy(
				"convertedString"))

		return convertedString, err

	case nil:
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is invalid!\n"+
			"'%v' is a 'nil' value.\n",
			ePrefix.String(),
			emptyIFaceParamName,
			emptyIFaceParamName)

		return convertedString, err

	default:
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is invalid!\n"+
			"'%v' is an unsupported type.\n"+
			"'%v' is Type: %T",
			ePrefix.String(),
			emptyIFaceParamName,
			emptyIFaceParamName,
			emptyIFaceParamName,
			emptyIFace)

		return convertedString, err
	}

dateTimeConversion:

	dateTimeInputDto,
		ok = emptyIFace.(TextInputParamFieldDateTimeDto)

	if !ok {
		err = fmt.Errorf("%v\n"+
			"Error: Failed to convert empty interface\n"+
			"(%v) to Date Time!\n"+
			"String Conversion Error.\n",
			ePrefix.String(),
			emptyIFaceParamName)

		return convertedString, err
	}

	if len(dateTimeInputDto.FieldDateTimeFormat) == 0 {
		dateTimeInputDto.FieldDateTimeFormat =
			textSpecificationMolecule{}.ptr().
				getDefaultDateTimeFormat()

	}

	convertedString = dateTimeInputDto.FieldDateTime.
		Format(dateTimeInputDto.FieldDateTimeFormat)

	return convertedString, err

standardConversion:

	convertedString = fmt.Sprintf("%v",
		emptyIFace)

	return convertedString, err

straightStringConversion:

	plainString,
		ok = emptyIFace.(string)

	if !ok {
		err = fmt.Errorf("%v\n"+
			"Error: Failed to convert empty interface\n"+
			"(%v) to string!\n"+
			"String Conversion Error.\n",
			ePrefix.String(),
			emptyIFaceParamName)

		return convertedString, err
	}

	convertedString = plainString

	return convertedString, err

stringerConversion:

	iStringer,
		ok = emptyIFace.(fmt.Stringer)

	if !ok {
		err = fmt.Errorf("%v\n"+
			"Error: Failed to convert empty interface\n"+
			"(%v) to string!\n"+
			"fmt.Stringer Conversion Error.\n",
			ePrefix.String(),
			emptyIFaceParamName)

		return convertedString, err
	}

	convertedString = iStringer.String()
	return convertedString, err
}

// getDefaultNewLineChars - Returns the default new line characters
// used by instances of TextLineSpecStandardLine and
// TextLineSpecPlainText.
//
// By default, each line of text generated by an instance of
// TextLineSpecPlainText and TextLineSpecPlainText will be
// terminated with a new line character. That new line character is
// currently defined as the character '\n'.
//
// This method returns that default new line character ('\n').
func (txtSpecAtom *textSpecificationAtom) getDefaultNewLineChars() []rune {

	if txtSpecAtom.lock == nil {
		txtSpecAtom.lock = new(sync.Mutex)
	}

	txtSpecAtom.lock.Lock()

	defer txtSpecAtom.lock.Unlock()

	newLineChars := []rune{'\n'}

	return newLineChars
}

// ptr - Returns a pointer to a new instance of
// textSpecificationAtom.
func (txtSpecAtom textSpecificationAtom) ptr() *textSpecificationAtom {

	if txtSpecAtom.lock == nil {
		txtSpecAtom.lock = new(sync.Mutex)
	}

	txtSpecAtom.lock.Lock()

	defer txtSpecAtom.lock.Unlock()

	return &textSpecificationAtom{
		lock: new(sync.Mutex),
	}
}

//	readBytes
//
//	This helper method is designed to support the
//	io.Reader interface.
//
//	This method reads up to len(p) bytes into p. It
//	returns the number of bytes read (0 <= n <= len(p))
//	and any error encountered. Even if read returns
//	n < len(p), it may use all of p as scratch space
//	during the call.
//
//	If some data is available but not len(p) bytes,
//	readBytes() conventionally returns what is available
//	instead of waiting for more.
//
//	When this method encounters an error or end-of-file
//	condition after successfully reading n > 0 bytes, it
//	returns the number of bytes read. It may return the
//	(non-nil) error from the same call or return the error
//	(and n == 0) from a subsequent call. An instance of
//	this general case is that a Reader returning a
//	non-zero number of bytes at the end of the input
//	stream may return either 'err == EOF' or
//	'err == nil'. The next Read operation should return
//	'0, EOF'.
//
//	Callers should always process the n > 0 bytes
//	returned before considering the error 'err'. Doing so
//	correctly handles I/O errors that happen after
//	reading some bytes and also both of the allowed EOF
//	behaviors.
//
//	This method supports buffered 'read' operations.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	textReader					*strings.Reader
//
//		A pointer to an instance of strings.Reader. The
//		text string encapsulated by 'textReader' will be
//		written to the byte buffer 'p'.
//
//
//	p							[]byte
//
//		The byte buffer into which the text string
//		encapsulated by parameter 'textReader' will be
//		written.
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
//	n                          int
//
//		The number of bytes written to byte buffer 'p'.
//
//
//	err                        error
//
//		If this method completes successfully, this
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		readBytes() reads up to len(p) bytes into p. It
//		returns the number of bytes read
//		(0 <= n <= len(p)) and any error encountered.
//		Even if readBytes() returns n < len(p), it may
//		use all of p as scratch space during the call. If
//		some data is available but not len(p) bytes,
//		readBytes() conventionally returns what is
//		available instead of waiting for more.
//
//		When readBytes() encounters an error or
//		'end-of-file' condition after successfully
//		reading n > 0 bytes, it returns the number of
//		bytes read. It may return the (non-nil) error
//		from the same call or return the error
//		(and n == 0) from a subsequent call. An instance
//		of this general case is that a Reader returning a
//		non-zero number of bytes at the end of the input
//		stream may return either 'err == EOF' or
//		'err == nil'. The next Read operation should
//		return '0, EOF'.
//
//		If an error message is returned and the error is
//		NOT equal to io.EOF, the text value for input
//		parameter 'errPrefDto' (error prefix) will be
//		prefixed or attached at the beginning of the
//		error message.
func (txtSpecAtom *textSpecificationAtom) readBytes(
	textReader *strings.Reader,
	p []byte,
	errPrefDto *ePref.ErrPrefixDto) (
	n int,
	err error) {

	if txtSpecAtom.lock == nil {
		txtSpecAtom.lock = new(sync.Mutex)
	}

	txtSpecAtom.lock.Lock()

	defer txtSpecAtom.lock.Unlock()

	n = 0

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textSpecificationAtom."+
			"readBytes()",
		"")

	if err != nil {
		return n, err
	}

	if textReader == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'textReader' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return n, err
	}

	pLen := len(p)

	if pLen == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input byte array 'p' has zero length!\n",
			ePrefix)

		return n, err
	}

	n,
		err = textReader.Read(p)

	return n, err
}
