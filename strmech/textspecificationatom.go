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
