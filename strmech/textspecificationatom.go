package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
)

type textSpecificationAtom struct {
	lock *sync.Mutex
}

// convertParamEmptyInterfaceToString
//
// Receives an object styled as an empty interface and
// attempts to convert that object to a string.
//
// If the empty interface object does NOT match one of
// over sixty supported data types, an error will be
// returned.
//
// Note that the processing of string arrays will result
// in a concatenation of all string array elements to
// generate a single string result.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	emptyIFace                 interface{}
//
//		This object will be converted to a type of string
//		and returned to the calling function.
//
//		This empty interface parameter must contain one
//		of over sixty supported data types. Eligible
//		objects will then be converted and returned as a
//		string.
//
//		Supported eligible types which may be submitted
//		through this empty interface parameter are listed
//		as follows:
//
//		Eligible Data Types
//
//	  		 1.	[]byte
//	  		 2.	*[]byte
//	  		 3.	string
//	  		 4.	*string
//	  		 5.	[]string
//	  		 6.	*[]string
//	  		 7.	Stringer (fmt.Stringer) Interface
//	  		 8.	strings.Builder
//	  		 9.	*strings.Builder
//	  		10.	StringArrayDto
//	  		11.	*StringArrayDto
//	  		12.	[]rune
//	  		13.	*[]rune
//	  		14.	RuneArrayDto
//	  		15.	*RuneArrayDto
//	  		16.	RuneArrayCollection
//	  		17.	*RuneArrayCollection
//	  		18.	ITextFieldFormatDto
//	  		19.	ITextFieldSpecification
//	  		20.	ITextLineSpecification
//	  		21.	TextLineSpecLinesCollection
//	  		22.	bool
//	  		23.	TextLineTitleMarqueeDto
//	  		24.	time.Time (Converted using default format)
//	  		25.	TextInputParamFieldDateTimeDto
//	  		26.	float32
//	  		27.	*float32
//	  		28.	float64
//	  		29.	*float64
//	  		30.	BigFloatDto
//	  		31.	*BigFloatDto
//	  		32.	big.Float
//	  		33.	*big.Float
//	  		34.	big.Rat
//	  		35.	*big.Rat
//	  		36.	int8
//	  		37.	*int8
//	  		38.	int16
//	  		39.	*int16
//	  		40.	int
//	  		41.	*int
//	  		42.	int32
//	  		43.	*int32
//	  		44.	int64
//	  		45.	*int64
//	  		46.	uint8
//	  		47.	*uint8
//	  		48.	uint16
//	  		49.	*uint16
//	  		50.	uint
//	  		51.	*uint
//	  		52.	uint32
//	  		53.	*uint32
//	  		54.	uint64,
//	  		55.	*uint64
//	  		56.	big.Int
//	  		57.	*big.Int
//	  		58.	TextFieldFormatDtoFloat64
//	  		59.	*TextFieldFormatDtoFloat64
//	  		60.	TextFieldFormatDtoBigFloat
//	  		61.	*TextFieldFormatDtoBigFloat
//	  		62.	NumberStrKernel
//	  		63.	*NumberStrKernel
//	  		64.	[]NumberStrKernel
//	  		65.	*[]NumberStrKernel
//
//		If the 'emptyIFace' object is not convertible to
//		one of the supported types, an error will be
//		returned.
//
//		----------------------------------------------------
//		BE ADVISED
//
//			If an 'emptyIFace' object which resolves as a
//			string array, this method will concatenate the
//			member string array elements to produce a single
//			string result.
//		----------------------------------------------------
//
//	 emptyIFaceParamName        string
//
//		This is the name or text label used to describe
//		input parameter 'emptyIFace' when formatting
//		informational or error messages. If this
//		parameter is submitted as an empty string its
//		value will be defaulted to 'emptyIFace'.
//
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
// ------------------------------------------------------------------------
//
// Return Values
//
//	convertedString				string
//
//		If this method completes successfully, this
//		parameter will be populated with a string value
//		extracted and converted from the empty interface
//		input parameter, 'emptyIFace'.
//
//		Be advised, the returned string may be empty or
//		have a zero length.
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

	var baseTypeConversion BaseTypeDto

	baseTypeConversion,
		err = new(typeConversionsAtom).
		convertParamsToBaseTypes(
			emptyIFace,
			emptyIFaceParamName,
			ePrefix.XCpy("emptyIFace"))

	if err != nil {

		return convertedString, err
	}

	if baseTypeConversion.IsAByteArrayDto == true {

		convertedString =
			string(baseTypeConversion.AByteArrayDto.ByteArray)

	} else if baseTypeConversion.IsARuneArrayDto == true {

		convertedString =
			string(baseTypeConversion.ARuneArrayDto.CharsArray)

	} else if baseTypeConversion.IsAString == true {

		convertedString = baseTypeConversion.AString

	} else if baseTypeConversion.IsAStringArrayDto == true {

		convertedString =
			baseTypeConversion.AStringArrayDto.
				ConcatenateStrings("")

	} else {

		err = fmt.Errorf("%v\n"+
			"-------------------------------------------------------\n"+
			"Error: typeConversionsAtom.convertParamsToBaseTypes()\n"+
			"returned an invalid 'BaseTypeDto' object. The return"+
			"value did not resolve as a byte array, rune array,"+
			"string array or string.\n",
			ePrefix.String())

	}

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
