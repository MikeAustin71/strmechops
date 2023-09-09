package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"strconv"
	"strings"
	"sync"
	"time"
)

type typeConversionsAtom struct {
	lock *sync.Mutex
}

// convertParamsToBaseTypes
//
// Receives an empty interface and attempts to convert that
// interface to one of over sixty eligible before return the
// contents:
//
//		Eligible Data Types
//
//	   1.	[]byte
//	   2.	*[]byte
//	   3.	string
//	   4.	*string
//	   5.	[]string
//	   6.	*[]string
//	   7.	Stringer (fmt.Stringer) Interface
//	   8.	strings.Builder
//	   9.	*strings.Builder
//	  10.	StringArrayDto
//	  11.	*StringArrayDto
//	  12.	[]rune
//	  13.	*[]rune
//	  14.	RuneArrayDto
//	  15.	*RuneArrayDto
//	  16.	RuneArrayCollection
//	  17.	*RuneArrayCollection
//	  18.	ITextFieldFormatDto
//	  19.	ITextFieldSpecification
//	  20.	ITextLineSpecification
//	  21.	TextLineSpecLinesCollection
//	  22.	bool
//	  23.	TextLineTitleMarqueeDto
//	  24.	time.Time
//	  25.	TextInputParamFieldDateTimeDto
//	  26.	float32
//	  27.	*float32
//	  28.	float64
//	  29.	*float64
//	  30.	BigFloatDto
//	  31.	*BigFloatDto
//	  32.	big.Float
//	  33.	*big.Float
//	  34.	big.Rat
//	  35.	*big.Rat
//	  36.	int8
//	  37.	*int8
//	  38.	int16
//	  39.	*int16
//	  40.	int
//	  41.	*int
//	  42.	int32
//	  43.	*int32
//	  44.	int64
//	  45.	*int64
//	  46.	uint8
//	  47.	*uint8
//	  48.	uint16
//	  49.	*uint16
//	  50.	uint
//	  51.	*uint
//	  52.	uint32
//	  53.	*uint32
//	  54.	uint64,
//	  55.	*uint64
//	  56.	big.Int
//	  57.	*big.Int
//	  58.	TextFieldFormatDtoFloat64
//	  59.	*TextFieldFormatDtoFloat64
//	  60.	TextFieldFormatDtoBigFloat
//	  61.	*TextFieldFormatDtoBigFloat
//	  62.	NumberStrKernel
//	  63.	*NumberStrKernel
//	  64.	[]NumberStrKernel
//	  65.	*[]NumberStrKernel
//
// The results of the interface conversion are returned
// and reported through an instance of BaseTypeDto.
// Depending on the data type passed through input
// parameter 'charsToConvert', this empty interface
// input parameter will be converted to one of the
// following four data elements encapsulated in the
// returned instance of BaseTypeDto:
//
//	BaseTypeDto.AByteArrayDto
//
//	BaseTypeDto.ARuneArrayDto
//
//	BaseTypeDto.AString
//
//	BaseTypeDto.AStringArrayDto
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	charsToConvert				interface{}
//
//		An empty interface containing an object matching
//		one of the more than 60-eligible types described
//		above.
//
//		If this object is NOT convertible to one of the
//		60+ eligible types supported by this method, an
//		error will be returned.
//
//		Eligible types passed through this parameter will be
//		one of the four base type specified by the returned
//		instance of BaseTypeDto.
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
//	baseTypeConversion			BaseTypeDto
//
//		This method will return an instance of
//		BaseTypeDto. Input parameter 'charsToConvert'
//		will be converted to one of four base types
//		encapsulated by this returned instance of
//		'BaseTypeDto'.
//
//		Type BaseTypeDto consists of the following data
//		elements:
//
//		type BaseTypeDto struct {
//
//			IsAByteArrayDto     bool
//			AByteArrayDto       ByteArrayDto
//			AByteArrayDtoLength int
//			AByteArrayDtoDesc1  string
//			AByteArrayDtoDesc2  string
//
//			IsARuneArrayDto     bool
//			ARuneArrayDto       RuneArrayDto
//			ARuneArrayDtoLength int
//			ARuneArrayDtoDesc1  string
//			ARuneArrayDtoDesc2  string
//
//			IsAString     		bool
//			AString       		string
//			AStringLength 		int
//			AStringDesc1  		string
//			AStringDesc2  		string
//
//			IsAStringArrayDto     bool
//			AStringArrayDto       StringArrayDto
//			AStringArrayDtoLength int
//			AStringArrayDtoDesc1  string
//			AStringArrayDtoDesc2  string
//		}
//
//		----------------------------------------------------
//		BE ADVISED
//
//			The description #1 (...Desc1) data element
//			will always be populated with the name of the
//			original source type passed through the empty
//			interface input parameter 'charsToConvert'.
//		----------------------------------------------------
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
func (typeConvAtom *typeConversionsAtom) convertParamsToBaseTypes(
	charsToConvert interface{},
	charsToConvertLabel string,
	errPrefDto *ePref.ErrPrefixDto) (
	baseTypeConversion BaseTypeDto,
	err error) {

	if typeConvAtom.lock == nil {
		typeConvAtom.lock = new(sync.Mutex)
	}

	typeConvAtom.lock.Lock()

	defer typeConvAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	funcName := "typeConversionsAtom." +
		"convertToStringArrayOrByteArray()"

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		funcName,
		"")

	if err != nil {

		return baseTypeConversion, err
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

		return baseTypeConversion, err
	}

	var ok bool
	var err2 error
	var numStrKernelArray []NumberStrKernel

	switch charsToConvert.(type) {

	case []byte:

		baseTypeConversion.AByteArrayDto.ByteArray,
			ok = charsToConvert.([]byte)

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

			return baseTypeConversion, err
		}

		baseTypeConversion.AByteArrayDtoDesc1 = "[]byte"

		baseTypeConversion.AByteArrayDtoLength =
			len(baseTypeConversion.AByteArrayDto.ByteArray)

		baseTypeConversion.IsAByteArrayDto = true

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

			return baseTypeConversion, err
		}

		baseTypeConversion.AByteArrayDto.ByteArray =
			*byteArrayPtr

		baseTypeConversion.AByteArrayDtoDesc1 = "*[]byte"

		baseTypeConversion.AByteArrayDtoLength =
			len(baseTypeConversion.AByteArrayDto.ByteArray)

		baseTypeConversion.IsAByteArrayDto = true

	case string:

		baseTypeConversion.AString,
			ok = charsToConvert.(string)

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

			return baseTypeConversion, err
		}

		baseTypeConversion.AStringDesc1 = "string"

		baseTypeConversion.AStringLength =
			len(baseTypeConversion.AString)

		baseTypeConversion.IsAString = true

	case *string:
		// string pointer

		var strPtr *string

		strPtr, ok = charsToConvert.(*string)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"Error: Input parameter '%v' is invalid!\n"+
				"'%v' was identified as a string pointer\n"+
				"(*string) However, the cast from '%v' to\n"+
				"string pointer Failed.\n",
				ePrefix.String(),
				charsToConvertLabel,
				charsToConvertLabel,
				charsToConvertLabel)

			return baseTypeConversion, err
		}

		baseTypeConversion.AStringDesc1 = "*string"

		baseTypeConversion.AString = *strPtr

		baseTypeConversion.AStringLength =
			len(baseTypeConversion.AString)

		baseTypeConversion.IsAString = true

	case []string:
		// string array

		var strArray []string

		strArray,
			ok = charsToConvert.([]string)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"ERROR: Input parameter '%v' is invalid!\n"+
				"'%v' was identified as a string array\n"+
				"([]string). However, the cast from '%v' to\n"+
				"string array Failed.\n",
				ePrefix.String(),
				charsToConvertLabel,
				charsToConvertLabel,
				charsToConvertLabel)

			return baseTypeConversion, err
		}

		baseTypeConversion.AStringArrayDtoDesc1 =
			"[]string"

		baseTypeConversion.AStringArrayDto =
			new(StringArrayDto).NewStringArray(
				strArray,
				"",
				"")

		baseTypeConversion.AStringArrayDtoLength =
			baseTypeConversion.AStringArrayDto.
				GetStringArrayLength()

		baseTypeConversion.IsAStringArrayDto = true

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

			return baseTypeConversion, err
		}

		baseTypeConversion.AStringArrayDto =
			new(StringArrayDto).NewStringArray(
				*strArrayPtr,
				"",
				"")

		baseTypeConversion.AStringArrayDtoDesc1 =
			"*[]string"

		baseTypeConversion.AStringArrayDtoLength =
			baseTypeConversion.AStringArrayDto.
				GetStringArrayLength()

		baseTypeConversion.IsAStringArrayDto = true

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

			return baseTypeConversion, err
		}

		baseTypeConversion.AString = strBuilder.String()

		baseTypeConversion.AStringDesc1 =
			"strings.Builder"

		baseTypeConversion.AStringLength =
			len(baseTypeConversion.AString)

		baseTypeConversion.IsAString = true

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

			return baseTypeConversion, err
		}

		baseTypeConversion.AString = strBuilderPtr.String()

		baseTypeConversion.AStringDesc1 =
			"*strings.Builder"

		baseTypeConversion.AStringLength =
			len(baseTypeConversion.AString)

		baseTypeConversion.IsAString = true

	case StringArrayDto:

		var strArrayDto StringArrayDto

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

			return baseTypeConversion, err
		}

		err2 = baseTypeConversion.AStringArrayDto.
			CopyIn(
				&strArrayDto,
				ePrefix.XCpy("strArrayDto"))

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"Error copying String Array Dto ('strArrayDto') to\n"+
				"baseTypeConversion.AStringArrayDto.\n"+
				"Error=\n%v\n",
				ePrefix.String(),
				err2.Error())

			return baseTypeConversion, err
		}

		baseTypeConversion.AStringArrayDtoDesc1 =
			"StringArrayDto"

		baseTypeConversion.AStringArrayDtoLength =
			baseTypeConversion.AStringArrayDto.GetStringArrayLength()

		baseTypeConversion.IsAStringArrayDto = true

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

			return baseTypeConversion, err
		}

		err2 = baseTypeConversion.AStringArrayDto.
			CopyIn(
				strArrayDtoPtr,
				ePrefix.XCpy("strArrayDtoPtr"))

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------------\n"+
				"Error copying String Array Dto Pointer '%v'"+
				"(*StringArrayDto) to baseTypeConversion.AStringArrayDto.\n"+
				"Error=\n%v\n",
				ePrefix.String(),
				charsToConvertLabel,
				err2.Error())

			return baseTypeConversion, err
		}

		baseTypeConversion.AStringArrayDtoDesc1 =
			"*StringArrayDto"

		baseTypeConversion.AStringArrayDtoLength =
			baseTypeConversion.AStringArrayDto.GetStringArrayLength()

		baseTypeConversion.IsAStringArrayDto = true

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

			return baseTypeConversion, err
		}

		baseTypeConversion.ARuneArrayDtoDesc1 =
			"[]rune"

		baseTypeConversion.ARuneArrayDto =
			new(RuneArrayDto).NewRunesDefault(
				runesToWrite)

		baseTypeConversion.ARuneArrayDtoLength =
			baseTypeConversion.ARuneArrayDto.
				GetRuneArrayLength()

		baseTypeConversion.IsARuneArrayDto = true

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

			return baseTypeConversion, err
		}

		baseTypeConversion.ARuneArrayDto =
			new(RuneArrayDto).NewRunesDefault(
				*runeArrayPtr)

		baseTypeConversion.ARuneArrayDtoDesc1 =
			"*[]rune"

		baseTypeConversion.ARuneArrayDtoLength =
			baseTypeConversion.ARuneArrayDto.
				GetRuneArrayLength()

		baseTypeConversion.IsARuneArrayDto = true

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

			return baseTypeConversion, err
		}

		err2 = baseTypeConversion.ARuneArrayDto.
			CopyIn(
				&runesArrayDto,
				ePrefix.XCpy("runesArrayDto"))

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"Error occurred while copying 'RuneArrayDto' to\n"+
				"baseTypeConversion.ARuneArrayDto\n"+
				"Error=\n%v\n",
				ePrefix.String(),
				err2.Error())

			return baseTypeConversion, err
		}

		baseTypeConversion.ARuneArrayDtoDesc1 =
			"RuneArrayDto"

		baseTypeConversion.ARuneArrayDtoLength =
			baseTypeConversion.ARuneArrayDto.
				GetRuneArrayLength()

		baseTypeConversion.IsARuneArrayDto = true

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

			return baseTypeConversion, err
		}

		err2 = baseTypeConversion.ARuneArrayDto.
			CopyIn(
				runeArrayDtoPtr,
				ePrefix.XCpy("runeArrayDtoPtr"))

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"Error occurred while copying '*RuneArrayDto' to\n"+
				"baseTypeConversion.ARuneArrayDto\n"+
				"Error=\n%v\n",
				ePrefix.String(),
				err2.Error())

			return baseTypeConversion, err
		}

		baseTypeConversion.ARuneArrayDtoDesc1 =
			"*RuneArrayDto"

		baseTypeConversion.ARuneArrayDtoLength =
			baseTypeConversion.ARuneArrayDto.
				GetRuneArrayLength()

		baseTypeConversion.IsARuneArrayDto = true

	case RuneArrayCollection:

		var runeArrayCol RuneArrayCollection

		runeArrayCol, ok = charsToConvert.(RuneArrayCollection)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"Error: Input parameter '%v' is invalid!\n"+
				"'%v' was identified as a type RuneArrayCollection\n"+
				"However, the cast from '%v' to RuneArrayCollection\n"+
				"Failed.\n",
				ePrefix.String(),
				charsToConvertLabel,
				charsToConvertLabel,
				charsToConvertLabel)

			return baseTypeConversion, err
		}

		baseTypeConversion.AStringArrayDtoDesc1 =
			"RuneArrayCollection"

		baseTypeConversion.AStringArrayDto =
			runeArrayCol.GetStringArrayDto()

		baseTypeConversion.AStringArrayDtoLength =
			baseTypeConversion.AStringArrayDto.GetStringArrayLength()

		baseTypeConversion.IsAStringArrayDto = true

	case *RuneArrayCollection:

		var runeArrayColPtr *RuneArrayCollection

		runeArrayColPtr, ok = charsToConvert.(*RuneArrayCollection)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"Error: Input parameter '%v' is invalid!\n"+
				"'%v' was identified as a type *RuneArrayCollection\n"+
				"However, the cast from '%v' to *RuneArrayCollection\n"+
				"Failed.\n",
				ePrefix.String(),
				charsToConvertLabel,
				charsToConvertLabel,
				charsToConvertLabel)

			return baseTypeConversion, err
		}

		baseTypeConversion.AStringArrayDtoDesc1 =
			"*RuneArrayCollection"

		baseTypeConversion.AStringArrayDto =
			runeArrayColPtr.GetStringArrayDto()

		baseTypeConversion.AStringArrayDtoLength =
			baseTypeConversion.AStringArrayDto.GetStringArrayLength()

		baseTypeConversion.IsAStringArrayDto = true

	case ITextFieldFormatDto:

		var textFileFormatDto ITextFieldFormatDto

		textFileFormatDto, ok = charsToConvert.(ITextFieldFormatDto)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"Error: Input parameter '%v' is invalid!\n"+
				"'%v' was identified as a type ITextFieldFormatDto.\n"+
				"The cast from '%v' to ITextFieldFormatDto Failed.\n",
				ePrefix.String(),
				charsToConvertLabel,
				charsToConvertLabel,
				charsToConvertLabel)

			return baseTypeConversion, err
		}

		baseTypeConversion.AString,
			err = textFileFormatDto.
			GetFormattedTextFieldStr(
				ePrefix.XCpy("strResult<-ITextFieldFormatDto"))

		if err != nil {

			return baseTypeConversion, err
		}

		baseTypeConversion.AStringDesc1 =
			"ITextFieldFormatDto"

		baseTypeConversion.AStringLength =
			len(baseTypeConversion.AString)

		baseTypeConversion.IsAString = true

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

			return baseTypeConversion, err
		}

		var fieldSpecStrBuilder strings.Builder

		err = textFieldSpecification.
			TextBuilder(
				&fieldSpecStrBuilder,
				ePrefix.XCpy("ITextFieldSpecification"))

		if err != nil {

			return baseTypeConversion, err
		}

		baseTypeConversion.AString =
			fieldSpecStrBuilder.String()

		baseTypeConversion.AStringDesc1 =
			"ITextFieldSpecification"

		baseTypeConversion.AStringLength =
			len(baseTypeConversion.AString)

		baseTypeConversion.IsAString = true

	case ITextLineSpecification:

		var texLineSpecification ITextLineSpecification

		texLineSpecification, ok =
			charsToConvert.(ITextLineSpecification)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"Error: Input parameter '%v' is invalid!\n"+
				"'%v' was identified as a type ITextLineSpecification.\n"+
				"However, the cast from '%v' to ITextLineSpecification\n"+
				"Failed.\n",
				ePrefix.String(),
				charsToConvertLabel,
				charsToConvertLabel,
				charsToConvertLabel)

			return baseTypeConversion, err
		}

		var fieldSpecStrBuilder strings.Builder

		err = texLineSpecification.
			TextBuilder(
				&fieldSpecStrBuilder,
				ePrefix.XCpy("fieldSpecStrBuilder<-ITextLineSpecification"))

		if err != nil {

			return baseTypeConversion, err
		}

		baseTypeConversion.AString =
			fieldSpecStrBuilder.String()

		baseTypeConversion.AStringDesc1 =
			"ITextLineSpecification"

		baseTypeConversion.AStringLength =
			len(baseTypeConversion.AString)

		baseTypeConversion.IsAString = true

	case TextLineSpecLinesCollection:

		var txtLineSpecCol TextLineSpecLinesCollection

		txtLineSpecCol, ok =
			charsToConvert.(TextLineSpecLinesCollection)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-----------------------------------------------------------\n"+
				"Error: Input parameter '%v' is invalid!\n"+
				"'%v' was identified as a type TextLineSpecLinesCollection.\n"+
				"However, the cast from '%v' to TextLineSpecLinesCollection\n"+
				"Failed.\n",
				ePrefix.String(),
				charsToConvertLabel,
				charsToConvertLabel,
				charsToConvertLabel)

			return baseTypeConversion, err
		}

		baseTypeConversion.AStringArrayDto,
			err2 = txtLineSpecCol.
			GetFmtTextStrArray(
				ePrefix.XCpy("AStringArrayDto<-TextLineSpecLinesCollection"))

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"Error extracting String Array Dto ('strArrayDto') from\n"+
				"TextLineSpecLinesCollection.GetFmtTextStrArray().\n"+
				"Error=\n%v\n",
				ePrefix.String(),
				err2.Error())

			return baseTypeConversion, err
		}

		baseTypeConversion.AStringArrayDtoDesc1 =
			"TextLineSpecLinesCollection"

		baseTypeConversion.AStringArrayDtoLength =
			baseTypeConversion.AStringArrayDto.GetStringArrayLength()

		baseTypeConversion.IsAStringArrayDto = true

	case TextLineTitleMarqueeDto:

		var txtLineTitleMarquee TextLineTitleMarqueeDto

		txtLineTitleMarquee, ok =
			charsToConvert.(TextLineTitleMarqueeDto)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------------\n"+
				"Error: Input parameter '%v' is invalid!\n"+
				"'%v' was identified as a type TextLineTitleMarqueeDto.\n"+
				"However, the cast from '%v' to TextLineTitleMarqueeDto\n"+
				"Failed.\n",
				ePrefix.String(),
				charsToConvertLabel,
				charsToConvertLabel,
				charsToConvertLabel)

			return baseTypeConversion, err
		}

		var strBuilder strings.Builder

		err2 = txtLineTitleMarquee.
			TextBuilder(
				&strBuilder,
				ePrefix.XCpy(
					"strBuilder<-txtLineTitleMarquee"))

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Error: String creation from converted\n"+
				"TextLineTitleMarqueeDto Failed!\n"+
				"Error=\n%v\n",
				funcName,
				err2.Error())

			return baseTypeConversion, err
		}

		baseTypeConversion.AString =
			strBuilder.String()

		baseTypeConversion.AStringDesc1 =
			"TextLineTitleMarqueeDto"

		baseTypeConversion.AStringLength =
			len(baseTypeConversion.AString)

		baseTypeConversion.IsAString = true

	case bool:

		var booleanValue bool

		booleanValue, ok =
			charsToConvert.(bool)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"Error: Input parameter '%v' is invalid!\n"+
				"'%v' was identified as a type bool.\n"+
				"However, the cast from '%v' to bool Failed.\n",
				ePrefix.String(),
				charsToConvertLabel,
				charsToConvertLabel,
				charsToConvertLabel)

			return baseTypeConversion, err
		}

		baseTypeConversion.AString =
			strconv.FormatBool(booleanValue)

		baseTypeConversion.AStringDesc1 =
			"bool"

		baseTypeConversion.AStringLength =
			len(baseTypeConversion.AString)

		baseTypeConversion.IsAString = true

	case time.Time:

		var dateTimeValue time.Time

		dateTimeValue, ok =
			charsToConvert.(time.Time)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"Error: Input parameter '%v' is invalid!\n"+
				"'%v' was identified as a type time.Time.\n"+
				"However, the cast from '%v' to time.Time\n"+
				"Failed.\n",
				ePrefix.String(),
				charsToConvertLabel,
				charsToConvertLabel,
				charsToConvertLabel)

			return baseTypeConversion, err
		}

		defaultDateTimeFormat :=
			new(textSpecificationMolecule).getDefaultDateTimeFormat()

		baseTypeConversion.AString =
			dateTimeValue.Format(defaultDateTimeFormat)

		baseTypeConversion.AStringDesc1 =
			"time.Time"

		baseTypeConversion.AStringLength =
			len(baseTypeConversion.AString)

		baseTypeConversion.IsAString = true

	case TextInputParamFieldDateTimeDto:

		var dateTimeInputDto TextInputParamFieldDateTimeDto

		dateTimeInputDto,
			ok = charsToConvert.(TextInputParamFieldDateTimeDto)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"------------------------------------------------------------------\n"+
				"Error: Input parameter '%v' is invalid!\n"+
				"'%v' was identified as a type TextInputParamFieldDateTimeDto.\n"+
				"However, the cast from '%v' to TextInputParamFieldDateTimeDto\n"+
				"Failed.\n",
				ePrefix.String(),
				charsToConvertLabel,
				charsToConvertLabel,
				charsToConvertLabel)

			return baseTypeConversion, err
		}

		if len(dateTimeInputDto.FieldDateTimeFormat) == 0 {
			dateTimeInputDto.FieldDateTimeFormat =
				new(textSpecificationMolecule).
					getDefaultDateTimeFormat()

		}

		baseTypeConversion.AString =
			dateTimeInputDto.FieldDateTime.
				Format(
					dateTimeInputDto.FieldDateTimeFormat)

		baseTypeConversion.AStringDesc1 =
			"TextInputParamFieldDateTimeDto"

		baseTypeConversion.AStringLength =
			len(baseTypeConversion.AString)

		baseTypeConversion.IsAString = true

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

		baseTypeConversion.AString,
			baseTypeConversion.AStringDesc1,
			err2 = new(mathHelperNanobot).
			numericValueToNativeNumStr(
				charsToConvert,
				ePrefix.XCpy("<-charsToConvert"))

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Error converting numeric value to a number string!\n"+
				"Error=\n%v\n",
				funcName,
				err2.Error())

			return baseTypeConversion, err

		}

		baseTypeConversion.AStringLength =
			len(baseTypeConversion.AString)

		baseTypeConversion.IsAString = true

	case []NumberStrKernel:

		numStrKernelArray, ok =
			charsToConvert.([]NumberStrKernel)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"ERROR: Input parameter '%v' is invalid!\n"+
				"'%v' was identified as a type []NumberStrKernel.\n"+
				"However, the cast from '%v' to []NumberStrKernel\n"+
				"Failed.\n",
				ePrefix.String(),
				charsToConvertLabel,
				charsToConvertLabel,
				charsToConvertLabel)

			return baseTypeConversion, err
		}

		baseTypeConversion.AStringArrayDtoDesc1 =
			"[]NumberStrKernel"

		goto convertNumStrKernelArray

	case *[]NumberStrKernel:

		var numStrKernelArrayPtr *[]NumberStrKernel

		numStrKernelArrayPtr, ok =
			charsToConvert.(*[]NumberStrKernel)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"--------------------------------------------------------------\n"+
				"ERROR: Input parameter '%v' is invalid!\n"+
				"'%v' was identified as a type *[]NumberStrKernel \n"+
				"or a pointer to a Number String Kernel array.\n"+
				"However, the cast from '%v' to *[]NumberStrKernel Failed.\n",
				ePrefix.String(),
				charsToConvertLabel,
				charsToConvertLabel,
				charsToConvertLabel)

			return baseTypeConversion, err
		}

		numStrKernelArray = *numStrKernelArrayPtr

		baseTypeConversion.AStringArrayDtoDesc1 =
			"*[]NumberStrKernel"

		goto convertNumStrKernelArray

	case fmt.Stringer:

		var iStringer fmt.Stringer

		iStringer,
			ok = charsToConvert.(fmt.Stringer)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"-------------------------------------------------------\n"+
				"Input parameter '%v' is ERROR!\n"+
				"'%v' was identified as a type Stringer\n"+
				"interface (fmt.Stringer). However, the cast from\n"+
				"'%v' to fmt.Stringer Failed.\n",
				ePrefix.String(),
				charsToConvertLabel,
				charsToConvertLabel,
				charsToConvertLabel)

			return baseTypeConversion, err
		}

		baseTypeConversion.AString =
			iStringer.String()

		baseTypeConversion.AStringDesc1 =
			"fmt.Stringer"

		baseTypeConversion.AStringLength =
			len(baseTypeConversion.AString)

		baseTypeConversion.IsAString = true

	}

	return baseTypeConversion, err

convertNumStrKernelArray:

	var stringToWrite string

	lastNumStrIdx := len(numStrKernelArray) - 1

	if lastNumStrIdx < 0 {
		err = fmt.Errorf("%v\n"+
			"Error: '%v' conversion to Number String\n"+
			"Kernel array failed.  '%v' converted to\n"+
			"an empty array with a zero length.\n",
			ePrefix.String(),
			charsToConvertLabel,
			charsToConvertLabel)

		return baseTypeConversion, err
	}

	for i := 0; i <= lastNumStrIdx; i++ {

		stringToWrite,
			_,
			err2 = numStrKernelArray[i].FmtNumStrNative(
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy(
				fmt.Sprintf(
					"numStrKernelArray[%v]", i)))

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Error: Number String Kernel Array Conversion\n"+
				"([]NumberStrKernel) numStrKernelArray[%v].FmtNumStrNative()\n"+
				"Error=\n%v\n",
				funcName,
				i,
				err2.Error())

			return baseTypeConversion, err

		}

		baseTypeConversion.AStringArrayDto.PushStr(stringToWrite)

		stringToWrite = ""

	}

	baseTypeConversion.AStringArrayDtoLength =
		baseTypeConversion.AStringArrayDto.GetStringArrayLength()

	baseTypeConversion.IsAStringArrayDto = true

	return baseTypeConversion, err
}
