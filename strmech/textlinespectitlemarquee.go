package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"strings"
	"sync"
)

// TextLineSpecTitleMarquee
//
// This Text Line Specification is designed to format
// multiple text lines used in the presentation of title
// marquees for screen displays, file output and
// printing.
//
// Type TextLineSpecTitleMarquee encapsulates three types
// of text lines used in generating title marquees:
//
//  1. Leading Marquee Lines
//     Usually consists of leading blank lines
//     and solid lines.
//
//  2. Title Lines
//     Consists entirely of text strings functioning
//     as the main title lines.
//
//  3. Trailing Marquee Lines
//     Usually consists of trailing blank lines
//     and solid lines.
//
// These text line types are consolidated to produce
// the formatted text lines which comprise the Title
// Marquee.
//
// Type TextLineSpecTitleMarquee implements the
// ITextLineSpecification interface.
type TextLineSpecTitleMarquee struct {
	leadingMarqueeLines  TextLineSpecLinesCollection
	titleLines           TextLineSpecLinesCollection
	trailingMarqueeLines TextLineSpecLinesCollection
	textLineReader       *strings.Reader

	lock *sync.Mutex
}

// AddBlankLine
//
// Adds a blank line to the Text Specification Lines
// Collection maintained by the current instance of
// TextLineSpecTitleMarquee.
//
// Type TextLineSpecTitleMarquee encapsulates three types
// of text lines used in generating title marquees:
//
//  1. Leading Marquee Lines
//     Usually consists of leading blank lines
//     and solid lines.
//
//  2. Title Lines
//     Consists entirely of text strings functioning
//     as the main title lines.
//
//  3. Trailing Marquee Lines
//     Usually consists of trailing blank lines
//     and solid lines.
//
// This method will add the blank line to one of these
// three collections as specified by parameter,
// 'titleMarqueeLineType'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numOfBlankLines				int
//
//		The number of blank lines which will be generated
//		and added to one of the three Text Specification
//		Lines Collections maintained by the current
//		instance of TextLineSpecTitleMarquee
//
//		If input parameter 'numOfBlankLines' is less than
//		one (1), it is invalid and an error will be
//		returned.
//
//		If input parameter 'numOfBlankLines' is greater
//		than one-million (1,000,000), it is invalid and
//		an error will be returned.
//
//	titleMarqueeLineType		TextTileLineType
//
//		Type TextTileLineType is an enumeration of
//		Title Marquee Text Line Types. This parameter
//		determines which text line collection will
//		receive the newly created blank line.
//
//		If this parameter is not set to one of the
//		following valid values, an error will be
//		returned.
//
//		Formal TextTileLineType Syntax
//
//			TextTileLineType(0).LeadingMarqueeLine()
//			TextTileLineType(0).TitleLine()
//			TextTileLineType(0).TrailingMarqueeLine()
//
//		Abbreviated TextTileLineType Syntax
//
//			TitleLineType.LeadingMarqueeLine()
//			TitleLineType.TitleLine()
//			TitleLineTypeTrailingMarqueeLine()
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
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
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (txtLineSpecTitleMarquee *TextLineSpecTitleMarquee) AddBlankLine(
	numOfBlankLines int,
	titleMarqueeLineType TextTileLineType,
	errorPrefix interface{}) error {

	if txtLineSpecTitleMarquee.lock == nil {
		txtLineSpecTitleMarquee.lock = new(sync.Mutex)
	}

	txtLineSpecTitleMarquee.lock.Lock()

	defer txtLineSpecTitleMarquee.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecTitleMarquee."+
			"AddBlankLine()",
		"")

	if err != nil {
		return err
	}

	var txtLineCollection *TextLineSpecLinesCollection

	var txtLineCollectionName string

	switch titleMarqueeLineType {

	case TitleLineType.LeadingMarqueeLine():

		txtLineCollection = &txtLineSpecTitleMarquee.leadingMarqueeLines

		txtLineCollectionName = "leadingMarqueeLines"

	case TitleLineType.TitleLine():

		txtLineCollection = &txtLineSpecTitleMarquee.titleLines

		txtLineCollectionName = "titleLines"

	case TitleLineType.TrailingMarqueeLine():

		txtLineCollection = &txtLineSpecTitleMarquee.trailingMarqueeLines

		txtLineCollectionName = "trailingMarqueeLines"

	default:

		err := fmt.Errorf("%v\n"+
			"Error: Input parameter 'titleMarqueeLineType' is invalid!\n"+
			" titleMarqueeLineType string value = '%v'\n"+
			"titleMarqueeLineType integer value = '%v'\n",
			ePrefix.String(),
			titleMarqueeLineType.String(),
			titleMarqueeLineType.XValueInt())

		return err
	}

	err = txtLineCollection.
		AddBlankLine(
			numOfBlankLines,
			ePrefix.XCpy(
				fmt.Sprintf("txtLineSpecTitleMarquee.%v",
					txtLineCollectionName)))

	return err
}

// AddSolidLine
//
// Adds a solid line to the Text Specification Lines
// Collection maintained by the current instance of
// TextLineSpecTitleMarquee.
//
// Type TextLineSpecTitleMarquee encapsulates three types
// of text lines used in generating title marquees:
//
//  1. Leading Marquee Lines
//     Usually consists of leading blank lines
//     and solid lines.
//
//  2. Title Lines
//     Consists entirely of text strings functioning
//     as the main title lines.
//
//  3. Trailing Marquee Lines
//     Usually consists of trailing blank lines
//     and solid lines.
//
// This method will add the solid line to one of these
// three collections as specified by parameter,
// 'titleMarqueeLineType'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	leftMarginStr				string
//
//		A string containing the text characters to be
//		positioned on the left side of the Solid Line.
//
//		If no left margin is required, set this parameter
//		to an empty string.
//
//		Example:
//			leftMarginStr  = "   " // 3-spaces
//			solidLineChars = "*"
//			solidLineCharsRepeatCount = 5
//			rightMarginStr = "" // Empty string
//			Solid line = "   *****"
//
//		If the 'leftMarginStr' string length is greater
//		than one-million (1,000,000), an error will be
//		returned.
//
//	solidLineChars				string
//
//		This string specifies the character or characters
//		which will comprise the solid line.
//
//		Example:
//			solidLineChars = "*"
//			solidLineCharsRepeatCount = 5
//			Solid line = "*****"
//
//		If this parameter is submitted as a zero length
//		string, an error will be returned.
//
//	solidLineCharsRepeatCount	int
//
//		This integer value specifies the number of times
//		that parameter 'solidLineChars' will be repeated
//		in constructing the solid line.
//
//		If this parameter is submitted with a value less
//		than one (1), an error will be returned.
//
//		Example:
//			solidLineChars = "*"
//			solidLineCharsRepeatCount = 5
//			Solid line = "*****"
//
//	rightMarginStr				string
//
//		A string containing the text characters to
//		be positioned on the right side of the Solid
//		Line.
//
//		If no right margin is required, set this
//		parameter to an empty string.
//
//		Example:
//			solidLineChars = "*"
//			solidLineCharsRepeatCount = 5
//			leftMarginStr = "" // Empty string
//			rightMarginStr = "   " // 3-spaces
//			Solid line = "*****   "
//
//		If the 'rightMarginStr' string length is greater
//		than one-million (1,000,000), an error will be
//		returned.
//
//	lineTerminator				string
//
//		Also known as 'New Line characters'. This string
//		contains one or more characters which will be
//		used to terminate the solid text line added to
//		the Text Specification Lines Collection for the
//		current instance of TextLineSpecTitleMarquee.
//
//		Example:
//			solidLineChars = "*"
//			solidLineCharsRepeatCount = 5
//			newLineChars = "??\n\n"
//			Solid line = "*****??\n\n"
//
//		If this parameter is submitted as a zero length
//		string, 'newLineChars' will be set to the default
//		new line character ("\n").
//
//		If this parameter is submitted with a string
//		length greater than one-million (1,000,000),
//		'newLineChars' will be set to the default new
//		line character ('\n').
//
//	turnLineTerminatorOff		bool
//
//		The 'turnLineTerminatorOff' flag controls whether
//		a line termination character or characters will
//		be automatically appended to each solid line of
//		text added to the Text Specification Lines
//		Collection maintained by the current instance of
//		TextLineSpecTitleMarquee.
//
//		When the boolean flag 'turnLineTerminatorOff' is
//		set to 'false', line terminators as defined by
//		parameter 'lineTerminator' will be applied as a
//		line termination sequence for each line of text
//		added to the Text Specification	Lines Collection.
//
//		When this boolean value is set to 'true', it
//		turns off or cancels the automatic generation of
//		line terminators for each line of text added to
//		the Text Specification Lines Collection.
//
//	numOfSolidLines				int
//
//		The number of solid lines to be created. If this
//		value is less than one (1), an error will be
//		returned.
//
//		The parameter controls the number of solid lines
//		which will be generated and added to the Text
//		Specification Lines Collection for the current
//		instance of TextLineSpecLinesCollection.
//
//	titleMarqueeLineType		TextTileLineType
//
//		Type TextTileLineType is an enumeration of
//		Title Marquee Text Line Types. This parameter
//		determines which text line collection will
//		receive the newly created solid line.
//
//		If this parameter is not set to one of the
//		following valid values, an error will be
//		returned.
//
//		Formal TextTileLineType Syntax
//
//			TextTileLineType(0).LeadingMarqueeLine()
//			TextTileLineType(0).TitleLine()
//			TextTileLineType(0).TrailingMarqueeLine()
//
//		Abbreviated TextTileLineType Syntax
//
//			TitleLineType.LeadingMarqueeLine()
//			TitleLineType.TitleLine()
//			TitleLineTypeTrailingMarqueeLine()
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
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
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (txtLineSpecTitleMarquee *TextLineSpecTitleMarquee) AddSolidLine(
	leftMarginStr string,
	solidLineChars string,
	solidLineCharsRepeatCount int,
	rightMarginStr string,
	lineTerminator string,
	turnLineTerminatorOff bool,
	numOfSolidLines int,
	titleMarqueeLineType TextTileLineType,
	errorPrefix interface{}) error {

	if txtLineSpecTitleMarquee.lock == nil {
		txtLineSpecTitleMarquee.lock = new(sync.Mutex)
	}

	txtLineSpecTitleMarquee.lock.Lock()

	defer txtLineSpecTitleMarquee.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecTitleMarquee."+
			"AddSolidLine()",
		"")

	if err != nil {
		return err
	}

	var txtLineCollection *TextLineSpecLinesCollection

	var txtLineCollectionName string

	switch titleMarqueeLineType {

	case TitleLineType.LeadingMarqueeLine():

		txtLineCollection = &txtLineSpecTitleMarquee.leadingMarqueeLines

		txtLineCollectionName = "leadingMarqueeLines"

	case TitleLineType.TitleLine():

		txtLineCollection = &txtLineSpecTitleMarquee.titleLines

		txtLineCollectionName = "titleLines"

	case TitleLineType.TrailingMarqueeLine():

		txtLineCollection = &txtLineSpecTitleMarquee.trailingMarqueeLines

		txtLineCollectionName = "trailingMarqueeLines"

	default:

		err := fmt.Errorf("%v\n"+
			"Error: Input parameter 'titleMarqueeLineType' is invalid!\n"+
			" titleMarqueeLineType string value = '%v'\n"+
			"titleMarqueeLineType integer value = '%v'\n",
			ePrefix.String(),
			titleMarqueeLineType.String(),
			titleMarqueeLineType.XValueInt())

		return err
	}

	err = txtLineCollection.
		AddSolidLine(
			leftMarginStr,
			solidLineChars,
			solidLineCharsRepeatCount,
			rightMarginStr,
			lineTerminator,
			turnLineTerminatorOff,
			numOfSolidLines,
			ePrefix.XCpy(
				fmt.Sprintf("txtLineSpecTitleMarquee.%v",
					txtLineCollectionName)))

	return err
}

// AddStdLineColumns
//
// Adds a standard text line to the Text Specification
// Lines Collection maintained by the current instance of
// TextLineSpecTitleMarquee.
//
// Type TextLineSpecTitleMarquee encapsulates three types
// of text lines used in generating title marquees:
//
//  1. Leading Marquee Lines
//     Usually consists of leading blank lines
//     and solid lines.
//
//  2. Title Lines
//     Consists entirely of text strings functioning
//     as the main title lines.
//
//  3. Trailing Marquee Lines
//     Usually consists of trailing blank lines
//     and solid lines.
//
// This method will add the standard line to one of these
// three collections as specified by parameter,
// 'titleMarqueeLineType'.
//
// The standard text line created and added to the
// collection will consist of one or more columns of text
// as specified by input parameter 'textFieldColumns'.
//
// 'textFieldColumns' is a variadic parameter which
// accepts a variable number of arguments of type
// ITextFieldFormatDto. These ITextFieldFormatDto
// objects contain all the specifications necessary to
// construct a text field. Each ITextFieldFormatDto
// object will be instantiated in the standard text line
// as a separate text field.
//
// ----------------------------------------------------------------
//
// # ITextFieldFormatDto Interface
//
//		This method processes objects implementing the
//		ITextFieldFormatDto interface to define text field
//		specifications used to generate multi-column lines of
//		text.
//
//		These text fields are then bundled to configure a
//		line of text returned as an instance of
//		TextLineSpecStandardLine.
//
//		Examples of concrete types implementing the
//		ITextFieldFormatDto interface are:
//
//				TextFieldFormatDtoBigFloat
//				TextFieldFormatDtoDate
//				TextFieldFormatDtoLabel
//				TextFieldFormatDtoFiller
//
//		The most frequently used type is the
//		TextFieldFormatDtoLabel structure which is defined
//		as follows:
//
//			type TextFieldFormatDtoLabel struct {
//
//				LeftMarginStr string
//					One or more characters used to create a left
//					margin for this Text Field.
//
//					If this parameter is set to an empty string, no
//					left margin will be configured for this Text
//					Field.
//
//				FieldContents interface{}
//					This parameter may contain one of several
//					specific data types. This empty interface type
//					will be converted to a string and configured as
//					the text column content within a text line.
//
//					Supported types which may be submitted through
//					this empty interface parameter are listed as
//					follows:
//
//					   time.Time (Converted using default format)
//					   string
//					   bool
//					   uint, uint8, uint16, uint32, uint64,
//					   int, int8, int16, int32, int64
//					   float32, float64
//					   *big.Int *big.Float
//					   fmt.Stringer (types that support this interface)
//					   TextInputParamFieldDateTimeDto
//					         (Converts date time to string. The best way
//					          to transmit and configure date time values.)
//
//				 FieldLength int
//					The length of the text field in which the
//					'FieldContents' will be displayed. If
//					'FieldLength' is less than the length of the
//					'FieldContents' string, it will be automatically
//					set equal to the 'FieldContents' string length.
//
//					To automatically set the value of 'FieldLength'
//					to the length of 'FieldContents', set this
//					parameter to a value of minus one (-1).
//
//					If this parameter is submitted with a value less
//					than minus one (-1) or greater than 1-million
//					(1,000,000), an error will be returned.
//
//					Field Length Examples
//
//						Example-1
//	 			        FieldContents String = "Hello World!"
//							FieldContents String Length = 12
//							FieldLength = 18
//							FieldJustify = TxtJustify.Center()
//							Text Field String =
//								"   Hello World!   "
//
//						Example-2
//	 			        FieldContents = "Hello World!"
//							FieldContents String Length = 12
//							FieldLength = 18
//							FieldJustify = TxtJustify.Left()
//							Text Field String =
//								"Hello World!      "
//
//						Example-3
//	 			        FieldContents = "Hello World!"
//							FieldContents String Length = 12
//							FieldLength = -1
//							FieldJustify = TxtJustify.Center() // Ignored
//							Text Field String =
//								"Hello World!"
//
//						Example-4
//	 			        FieldContents = "Hello World!"
//							FieldContents String Length = 12
//							FieldLength = 2
//							FieldJustify = TxtJustify.Center()
//								Ignored, because FieldLength Less
//								Than FieldContents String Length.
//							Text Field String =
//								"Hello World!"
//
//				 FieldJustify TextJustify
//					An enumeration which specifies the justification
//					of the 'FieldContents' string within the text
//					field length specified by 'FieldLength'.
//
//					Text justification can only be evaluated in the
//					context of a text label ('FieldContents'), field
//					length ('FieldLength') and a Text Justification
//					object of type TextJustify. This is because text
//					labels with a field length equal to or less than
//					the length of the text label string will never
//					use text justification. In these cases, text
//					justification is completely ignored.
//
//					If the field length is greater than the length of
//					the text label string, text justification must be
//					equal to one of these three valid values:
//
//					    TextJustify(0).Left()
//					    TextJustify(0).Right()
//					    TextJustify(0).Center()
//
//					Users can also specify the abbreviated text
//					justification enumeration syntax as follows:
//
//					    TxtJustify.Left()
//					    TxtJustify.Right()
//					    TxtJustify.Center()
//
//					Text Justification Examples
//
//						Example-1
//	 			        FieldContents String = "Hello World!"
//							FieldContents String Length = 12
//							FieldLength = 18
//							FieldJustify = TxtJustify.Center()
//							Text Field String =
//								"   Hello World!   "
//
//						Example-2
//	 			        FieldContents = "Hello World!"
//							FieldContents String Length = 12
//							FieldLength = 18
//							FieldJustify = TxtJustify.Left()
//							Text Field String =
//								"Hello World!      "
//
//						Example-3
//	 			        FieldContents = "Hello World!"
//							FieldContents String Length = 12
//							FieldLength = -1
//							FieldJustify = TxtJustify.Center() // Ignored
//							Text Field String =
//								"Hello World!"
//
//						Example-4
//	 			        FieldContents = "Hello World!"
//							FieldContents String Length = 12
//							FieldLength = 2
//							FieldJustify = TxtJustify.Center()
//								Ignored, because FieldLength Less
//								Than FieldContents String Length.
//							Text Field String =
//								"Hello World!"
//
//				RightMarginStr string
//					One or more characters used to create a right
//					margin for this Text Field.
//
//					If this parameter is set to an empty string, no
//					right margin will be configured for this Text
//					Field.
//			}
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	lineTerminator				string
//
//		Also known as 'New Line characters'. This string
//		contains the text characters which will be
//		applied as line termination characters for each
//		line of text added to the Text Specification
//		Lines Collection maintained by the current
//		instance of TextLineSpecTitleMarquee.
//
//		By default, each line of text generated will be
//		terminated with	a new line character ('\n').
//		However, this parameter allows the user to
//		specify the character or characters to be used
//		as a line termination sequence for each line of
//		text added to the Text Specification Lines
//		Collection.
//
//		If this parameter is submitted as an empty string
//		with zero string length, this method will set
//		'lineTerminator' to the default new line
//		termination character ("\n").
//
//	turnLineTerminatorOff		bool
//
//		The 'turnLineTerminatorOff' flag controls whether
//		a line termination character or characters will
//		be automatically appended to each line of text
//		added to the Text Specification Lines Collection
//		maintained by the current instance of
//		TextLineSpecTitleMarquee.
//
//		When the boolean flag 'turnLineTerminatorOff' is
//		set to 'false', line terminators as defined by
//		parameter 'lineTerminator' will be applied as a
//		line termination sequence for each line of text
//		added to the Text Specification Lines Collection.
//
//		When this boolean value is set to 'true', it
//		turns off or cancels the automatic generation of
//		line terminators for each line of text produced
//		by TextLineSpecStandardLine.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
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
//	textFieldColumns			...ITextFieldFormatDto
//
//		This variadic parameter is used to pass one or
//		more instances of objects implementing the
//		ITextFieldFormatDto interface.
//
//		In the Go Programming language variadic
//		parameters will accept a variable number of
//		arguments.
//
//		These ITextFieldFormatDto object contains all the
//		text field content and formatting specifications
//		necessary to format one or more text fields in
//		a standard line of text.
//
//		Examples of concrete types implementing the
//		ITextFieldFormatDto interface are:
//
//			TextFieldFormatDtoBigFloat
//			TextFieldFormatDtoDate
//			TextFieldFormatDtoLabel
//			TextFieldFormatDtoFiller
//
//		For additional information on the
//		ITextFieldFormatDto interface, see above.
//
//		Note: In the Go Programming language the
//		variadic arguments must be positioned last
//		in the parameter list.
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
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (txtLineSpecTitleMarquee *TextLineSpecTitleMarquee) AddStdLineColumns(
	lineTerminator string,
	turnLineTerminatorOff bool,
	titleMarqueeLineType TextTileLineType,
	errorPrefix interface{},
	textFieldColumns ...ITextFieldFormatDto) error {

	if txtLineSpecTitleMarquee.lock == nil {
		txtLineSpecTitleMarquee.lock = new(sync.Mutex)
	}

	txtLineSpecTitleMarquee.lock.Lock()

	defer txtLineSpecTitleMarquee.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecTitleMarquee."+
			"AddStdLineColumns()",
		"")

	if err != nil {
		return err
	}
	//
	//var newStdLine TextLineSpecStandardLine
	//
	//newStdLine,
	//	err = TextLineSpecStandardLine{}.NewStdLineColumns(
	//	lineTerminator,
	//	turnLineTerminatorOff,
	//	ePrefix.XCpy(
	//		"newStdLine<-textFieldColumns..."),
	//	textFieldColumns...)

	if err != nil {
		return err
	}

	var txtLineCollection *TextLineSpecLinesCollection

	var txtLineCollectionName string

	switch titleMarqueeLineType {

	case TitleLineType.LeadingMarqueeLine():

		txtLineCollection = &txtLineSpecTitleMarquee.leadingMarqueeLines

		txtLineCollectionName = "leadingMarqueeLines"

	case TitleLineType.TitleLine():

		txtLineCollection = &txtLineSpecTitleMarquee.titleLines

		txtLineCollectionName = "titleLines"

	case TitleLineType.TrailingMarqueeLine():

		txtLineCollection = &txtLineSpecTitleMarquee.trailingMarqueeLines

		txtLineCollectionName = "trailingMarqueeLines"

	default:

		err := fmt.Errorf("%v\n"+
			"Error: Input parameter 'titleMarqueeLineType' is invalid!\n"+
			" titleMarqueeLineType string value = '%v'\n"+
			"titleMarqueeLineType integer value = '%v'\n",
			ePrefix.String(),
			titleMarqueeLineType.String(),
			titleMarqueeLineType.XValueInt())

		return err
	}

	err = txtLineCollection.
		AddStdLineColumns(
			lineTerminator,
			turnLineTerminatorOff,
			ePrefix.XCpy(
				fmt.Sprintf("txtLineSpecTitleMarquee.%v",
					txtLineCollectionName)),
			textFieldColumns...)

	return err
}

// CopyIn
//
//	Copies the data fields from an incoming instance of
//	NumberStrKernel ('incomingNumStrKernel') to the	data
//	fields of the current NumberStrKernel instance
//	('numStrKernel').
//
// # IMPORTANT
//
// ----------------------------------------------------------------
//
// All the data fields in current NumberStrKernel instance
// ('numStrKernel') will be modified and overwritten.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	incomingTitleMarquee		*TextLineSpecTitleMarquee
//
//		A pointer to an instance of
//		TextLineSpecTitleMarquee. This method will NOT
//		change the values of internal member data
//		variables contained in this instance.
//
//		All data values in this TextLineSpecTitleMarquee
//		instance will be copied to the current
//		TextLineSpecTitleMarquee instance
//		('txtLineSpecTitleMarquee').
//
//		If parameter 'incomingTitleMarquee' is determined
//		to be invalid, an error will be returned.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
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
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (txtLineSpecTitleMarquee *TextLineSpecTitleMarquee) CopyIn(
	incomingTitleMarquee *TextLineSpecTitleMarquee,
	errorPrefix interface{}) error {

	if txtLineSpecTitleMarquee.lock == nil {
		txtLineSpecTitleMarquee.lock = new(sync.Mutex)
	}

	txtLineSpecTitleMarquee.lock.Lock()

	defer txtLineSpecTitleMarquee.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecTitleMarquee."+
			"CopyIn()",
		"")

	if err != nil {
		return err
	}

	return new(textLineSpecTitleMarqueeNanobot).
		copy(txtLineSpecTitleMarquee,
			incomingTitleMarquee,
			ePrefix.XCpy(
				"txtLineSpecTitleMarquee<-"+
					"incomingTitleMarquee"))
}

//	CopyOut
//
//	Returns a deep copy of the current
//	TextLineSpecTitleMarquee instance.
//
//	If the current TextLineSpecTitleMarquee instance
//	contains invalid member variables, this method will
//	return an error.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
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
//	TextLineSpecTitleMarquee
//
//		If this method completes successfully, a deep
//		copy of the current TextLineSpecTitleMarquee
//		instance will be returned through this parameter.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (txtLineSpecTitleMarquee *TextLineSpecTitleMarquee) CopyOut(
	errorPrefix interface{}) (
	TextLineSpecTitleMarquee,
	error) {

	if txtLineSpecTitleMarquee.lock == nil {
		txtLineSpecTitleMarquee.lock = new(sync.Mutex)
	}

	txtLineSpecTitleMarquee.lock.Lock()

	defer txtLineSpecTitleMarquee.lock.Unlock()

	deepCopyTitleMarquee := TextLineSpecTitleMarquee{}

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecTitleMarquee."+
			"CopyOut()",
		"")

	if err != nil {
		return deepCopyTitleMarquee, err
	}

	err = new(textLineSpecTitleMarqueeNanobot).
		copy(&deepCopyTitleMarquee,
			txtLineSpecTitleMarquee,
			ePrefix.XCpy(
				"deepCopyTitleMarquee<-"+
					"txtLineSpecTitleMarquee"))

	return deepCopyTitleMarquee, err
}

//	CopyOutITextLine
//
//	Returns a deep copy of the current
//	TextLineSpecTitleMarquee instance cast as a type
//	ITextLineSpecification.
//
//	This method fulfills requirements of the
//	ITextLineSpecification interface.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
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
//	ITextLineSpecification
//
//		If this method completes successfully and no
//		errors are encountered, this parameter will
//		return a deep copy of the current
//		TextLineSpecTitleMarquee instance cast as an
//		ITextLineSpecification object.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (txtLineSpecTitleMarquee *TextLineSpecTitleMarquee) CopyOutITextLine(
	errorPrefix interface{}) (
	ITextLineSpecification,
	error) {

	if txtLineSpecTitleMarquee.lock == nil {
		txtLineSpecTitleMarquee.lock = new(sync.Mutex)
	}

	txtLineSpecTitleMarquee.lock.Lock()

	defer txtLineSpecTitleMarquee.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	deepCopyTitleMarquee := TextLineSpecTitleMarquee{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecTitleMarquee."+
			"CopyOutITextLine()",
		"")

	if err != nil {
		return ITextLineSpecification(&deepCopyTitleMarquee), err
	}

	err = new(textLineSpecTitleMarqueeNanobot).
		copy(&deepCopyTitleMarquee,
			txtLineSpecTitleMarquee,
			ePrefix.XCpy(
				"deepCopyTitleMarquee<-"+
					"txtLineSpecTitleMarquee"))

	return ITextLineSpecification(&deepCopyTitleMarquee), err
}

//	CopyOutPtr
//
//	Returns a pointer to a deep copy of the current
//	TextLineSpecTitleMarquee instance.
//
//	If the current TextLineSpecTitleMarquee instance
//	contains invalid member variables, this method will
//	return an error.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
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
//	*TextLineSpecTitleMarquee
//
//		If this method completes successfully, a pointer
//		to a deep copy of the current
//		TextLineSpecTitleMarquee instance will be
//		returned through this parameter.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (txtLineSpecTitleMarquee *TextLineSpecTitleMarquee) CopyOutPtr(
	errorPrefix interface{}) (
	*TextLineSpecTitleMarquee,
	error) {

	if txtLineSpecTitleMarquee.lock == nil {
		txtLineSpecTitleMarquee.lock = new(sync.Mutex)
	}

	txtLineSpecTitleMarquee.lock.Lock()

	defer txtLineSpecTitleMarquee.lock.Unlock()

	deepCopyTitleMarquee := TextLineSpecTitleMarquee{}

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecTitleMarquee."+
			"CopyOutPtr()",
		"")

	if err != nil {
		return &deepCopyTitleMarquee, err
	}

	err = new(textLineSpecTitleMarqueeNanobot).
		copy(&deepCopyTitleMarquee,
			txtLineSpecTitleMarquee,
			ePrefix.XCpy(
				"deepCopyTitleMarquee<-"+
					"txtLineSpecTitleMarquee"))

	return &deepCopyTitleMarquee, err
}

//	Empty
//
//	Resets all internal member variables for the current
//	instance of TextLineSpecTitleMarquee to their initial
//	or zero values.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete all pre-existing internal
//	member variable data values in the current instance
//	of TextLineSpecTitleMarquee.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	NONE
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	NONE
func (txtLineSpecTitleMarquee *TextLineSpecTitleMarquee) Empty() {

	if txtLineSpecTitleMarquee.lock == nil {
		txtLineSpecTitleMarquee.lock = new(sync.Mutex)
	}

	txtLineSpecTitleMarquee.lock.Lock()

	new(textLineSpecTitleMarqueeElectron).
		empty(txtLineSpecTitleMarquee)

	txtLineSpecTitleMarquee.lock.Unlock()

	txtLineSpecTitleMarquee.lock = nil
}

// EmptyLeadingMarqueeLines
//
// Type TextLineSpecTitleMarquee encapsulates three types
// of text lines used in generating title marquees:
//
//  1. Leading Marquee Lines
//     Usually consists of leading blank lines
//     and solid lines.
//
//  2. Title Lines
//     Consists entirely of text strings functioning
//     as the main title lines.
//
//  3. Trailing Marquee Lines
//     Usually consists of trailing blank lines
//     and solid lines.
//
// This method will delete all Leading Marquee Lines in
// the current instance of TextLineSpecTitleMarquee.
//
// The internal member variable to be deleted is:
//
//	TextLineSpecTitleMarquee.leadingMarqueeLines
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	NONE
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	NONE
func (txtLineSpecTitleMarquee *TextLineSpecTitleMarquee) EmptyLeadingMarqueeLines() {

	if txtLineSpecTitleMarquee.lock == nil {
		txtLineSpecTitleMarquee.lock = new(sync.Mutex)
	}

	txtLineSpecTitleMarquee.lock.Lock()

	defer txtLineSpecTitleMarquee.lock.Unlock()

	new(textLineSpecTitleMarqueeElectron).
		emptyLeadingMarqueeLines(txtLineSpecTitleMarquee)

	return
}

// EmptyTitleLines
//
// Type TextLineSpecTitleMarquee encapsulates three types
// of text lines used in generating title marquees:
//
//  1. Leading Marquee Lines
//     Usually consists of leading blank lines
//     and solid lines.
//
//  2. Title Lines
//     Consists entirely of text strings functioning
//     as the main title lines.
//
//  3. Trailing Marquee Lines
//     Usually consists of trailing blank lines
//     and solid lines.
//
// This method will delete all Title Lines. The internal
// member variable to be deleted is:
//
//	TextLineSpecTitleMarquee.titleLines
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	NONE
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	NONE
func (txtLineSpecTitleMarquee *TextLineSpecTitleMarquee) EmptyTitleLines() {

	if txtLineSpecTitleMarquee.lock == nil {
		txtLineSpecTitleMarquee.lock = new(sync.Mutex)
	}

	txtLineSpecTitleMarquee.lock.Lock()

	defer txtLineSpecTitleMarquee.lock.Unlock()

	new(textLineSpecTitleMarqueeElectron).
		emptyTitleLines(txtLineSpecTitleMarquee)

	return
}

// EmptyTrailingMarqueeLines
//
// Type TextLineSpecTitleMarquee encapsulates three types
// of text lines used in generating title marquees:
//
//  1. Leading Marquee Lines
//     Usually consists of leading blank lines
//     and solid lines.
//
//  2. Title Lines
//     Consists entirely of text strings functioning
//     as the main title lines.
//
//  3. Trailing Marquee Lines
//     Usually consists of trailing blank lines
//     and solid lines.
//
// This method will delete all Trailing Marquee Lines in
// the current instance of TextLineSpecTitleMarquee.
//
// The internal member variable to be deleted is:
//
//	TextLineSpecTitleMarquee.trailingMarqueeLines
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	NONE
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	NONE
func (txtLineSpecTitleMarquee *TextLineSpecTitleMarquee) EmptyTrailingMarqueeLines() {

	if txtLineSpecTitleMarquee.lock == nil {
		txtLineSpecTitleMarquee.lock = new(sync.Mutex)
	}

	txtLineSpecTitleMarquee.lock.Lock()

	defer txtLineSpecTitleMarquee.lock.Unlock()

	new(textLineSpecTitleMarqueeElectron).
		emptyTrailingMarqueeLines(txtLineSpecTitleMarquee)

	return
}

//	Equal
//
//	Receives a pointer to another instance of
//	TextLineSpecTitleMarquee and proceeds to compare the
//	member variables to those of the current
//	TextLineSpecTitleMarquee instance in order to
//	determine if they are equivalent.
//
//	A boolean flag showing the result of this comparison
//	is returned. If the member variables of both
//	instances are equal in all respects, this flag is set
//	to 'true'. Otherwise, this method returns 'false'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	incomingTitleMarquee	*TextLineSpecTitleMarquee
//
//		A pointer to an incoming instance of
//		TextLineSpecTitleMarquee. This method will
//		compare all member variable data values in this
//		instance against those contained in the current
//		instance of TextLineSpecTitleMarquee. If the data
//		values in both instances are found to be equal in
//		all respects, this method will return a boolean
//		value of 'true'.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		If the member variable data values contained in
//		input parameter 'incomingTitleMarquee' are equal
//		in all respects to those contained in the current
//		instance of TextLineSpecTitleMarquee, this method
//		will return a boolean value of 'true'. Otherwise,
//		a value of 'false' will be returned.
func (txtLineSpecTitleMarquee *TextLineSpecTitleMarquee) Equal(
	incomingTitleMarquee *TextLineSpecTitleMarquee) bool {

	if txtLineSpecTitleMarquee.lock == nil {
		txtLineSpecTitleMarquee.lock = new(sync.Mutex)
	}

	txtLineSpecTitleMarquee.lock.Lock()

	defer txtLineSpecTitleMarquee.lock.Unlock()

	return new(textLineSpecTitleMarqueeElectron).
		equal(
			txtLineSpecTitleMarquee,
			incomingTitleMarquee)
}

//	EqualITextLine
//
//	Receives an object implementing the
//	ITextLineSpecification interface and proceeds to
//	compare the member variables to those of the current
//	TextLineSpecTitleMarquee instance in order to
//	determine if they are equivalent.
//
//	A boolean flag showing the result of this comparison
//	is returned. If the member variables from both
//	instances are equal in all respects, this flag is set
//	to 'true'. Otherwise, this method returns 'false'.
//
//	This method is required by interface
//	ITextLineSpecification.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	iTextLine	ITextLineSpecification
//
//		An object implementing the ITextLineSpecification
//		interface. If this object proves to be equal in
//		all respects to the current instance of
//		TextLineSpecTitleMarquee, this method will return
//		'true'.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		If input parameter 'iTextLine' is judged to be
//		equal in all respects to the current instance of
//		TextLineSpecTitleMarquee, this return parameter
//		will be set to 'true'.
//
//		If 'iTextLine' is NOT equal to the current
//		instance of TextLineSpecTitleMarquee, this return
//		parameter will be set to 'false'.
func (txtLineSpecTitleMarquee *TextLineSpecTitleMarquee) EqualITextLine(
	iTextLine ITextLineSpecification) bool {

	if txtLineSpecTitleMarquee.lock == nil {
		txtLineSpecTitleMarquee.lock = new(sync.Mutex)
	}

	txtLineSpecTitleMarquee.lock.Lock()

	defer txtLineSpecTitleMarquee.lock.Unlock()

	txtTitleMarqueeTwo, ok := iTextLine.(*TextLineSpecTitleMarquee)

	if !ok {
		return false
	}

	return new(textLineSpecTitleMarqueeElectron).
		equal(
			txtLineSpecTitleMarquee,
			txtTitleMarqueeTwo)
}

//	GetFormattedText
//
//	Returns the formatted text generated by this Text
//	Line Specification Title Marquee
//	(TextLineSpecTitleMarquee) for text display, file
//	output, screen display and printing.
//
//	This method is similar to
//	TextLineSpecTitleMarquee.String() with the sole
//	difference being that this method returns an error.
//
//	This method fulfills requirements of the
//	ITextLineSpecification interface.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
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
//	string
//
//		The formatted text lines generated by the current
//		instance of TextLineSpecTitleMarquee.
//
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (txtLineSpecTitleMarquee *TextLineSpecTitleMarquee) GetFormattedText(
	errorPrefix interface{}) (
	string,
	error) {

	if txtLineSpecTitleMarquee.lock == nil {
		txtLineSpecTitleMarquee.lock = new(sync.Mutex)
	}

	txtLineSpecTitleMarquee.lock.Lock()

	defer txtLineSpecTitleMarquee.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecTitleMarquee."+
			"GetFormattedText()",
		"")

	if err != nil {
		return "", err
	}

	strBuilder := strings.Builder{}

	_,
		_,
		err = new(textLineSpecTitleMarqueeMolecule).
		getFormattedText(
			&strBuilder,
			txtLineSpecTitleMarquee,
			ePrefix.XCpy(
				"txtLineSpecTitleMarquee"))

	return strBuilder.String(), err
}

//	IsValidInstance
//
//	Performs a diagnostic review of the data values
//	encapsulated in the current TextLineSpecTitleMarquee
//	instance to determine if they are valid.
//
//	If all data element evaluate as valid, this method
//	returns 'true'. If any data element is invalid, this
//	method returns 'false'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	--- NONE ---
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	isValid						bool
//
//		If all data elements encapsulated by the current
//		instance of TextLineSpecTitleMarquee are valid,
//		this returned boolean value is set to 'true'. If
//		any data values are invalid, this return
//		parameter is set to 'false'.
func (txtLineSpecTitleMarquee *TextLineSpecTitleMarquee) IsValidInstance() (
	isValid bool) {

	if txtLineSpecTitleMarquee.lock == nil {
		txtLineSpecTitleMarquee.lock = new(sync.Mutex)
	}

	txtLineSpecTitleMarquee.lock.Lock()

	defer txtLineSpecTitleMarquee.lock.Unlock()

	isValid,
		_ = new(textLineSpecTitleMarqueeElectron).
		testValidityTitleMarquee(
			txtLineSpecTitleMarquee,
			nil)

	return isValid
}

//	IsValidInstanceError
//
//	Performs a diagnostic review of the data values
//	encapsulated in the current TextLineSpecTitleMarquee
//	instance to determine if they are valid.
//
//	If any data element evaluates as invalid, this method
//	will return an error.
//
//	This method fulfills requirements of
//	ITextLineSpecification interface.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
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
//	error
//
//
//		If the current instance of
//		TextLineSpecTitleMarquee is judged to be valid in
//		all respects, the returned error Type is set
//		equal to 'nil'.
//
//		If input parameter 'TextLineSpecTitleMarquee' is
//		found to be invalid, the returned error Type will
//		encapsulate an appropriate error message.
//
//		If an error message is returned, the text value
//		for input parameter 'errorPrefix' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (txtLineSpecTitleMarquee *TextLineSpecTitleMarquee) IsValidInstanceError(
	errorPrefix interface{}) error {

	if txtLineSpecTitleMarquee.lock == nil {
		txtLineSpecTitleMarquee.lock = new(sync.Mutex)
	}

	txtLineSpecTitleMarquee.lock.Lock()

	defer txtLineSpecTitleMarquee.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecTitleMarquee."+
			"IsValidInstanceError()",
		"")

	if err != nil {
		return err
	}

	_,
		err = new(textLineSpecTitleMarqueeElectron).
		testValidityTitleMarquee(
			txtLineSpecTitleMarquee,
			ePrefix.XCpy(
				"txtLineSpecTitleMarquee"))

	return err
}

//	NewAllParams
//
//	Creates and returns a new instance of
//	TextLineSpecTitleMarquee.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	standardTitleLeftMargin		string
//
//		The standard left margin characters applied
//		to all Text Title Lines in the Title Lines
//		array (TextLineSpecTitleMarquee.titleLines).
//
//	standardTitleRightMargin		string
//
//		The standard right margin characters applied
//		to all Text Title Lines in the Title Lines
//		array (TextLineSpecTitleMarquee.titleLines).
//
//	standardMaxLineLen					int
//
//		The maximum number of characters allowed on
//		a text title line. This maximum limit will be
//		applied to the length of all text lines generated
//		by the returned instance of
//		TextLineSpecTitleMarquee.
//
//	standardTextFieldLen		int
//
//		The standard field length applied to Text
//		Title Lines in the 'TitleLines' array unless
//		overridden by user customizations.
//
//		If the standardTextFieldLen exceeds the value of
//		the Maximum Available Text Field Length, it will
//		be reset and defaulted to the Maximum Available
//		Text Field Length.
//
//		The Maximum Available Text Field Length is
//		calculated as follows:
//
//		Maximum Available Text Field Length =
//			TextLineTitleMarqueeDto.StandardMaxLineLen -
//			1 -
//			len(TextLineTitleMarqueeDto.StandardTitleLeftMargin) -
//			len(TextLineTitleMarqueeDto.StandardTitleRightMargin)
//
//	standardTextJustification	TextJustify
//
//		The standard Justification specification applied to
//		the standard text field.
//
//		Type 'TextJustify' is an enumeration which
//		specifies the justification of the text field
//		contents string within the text	field length
//		specified by 'standardTextFieldLen'.
//
//		Text justification can only be evaluated in the
//		context of a text label, field length and a Text
//		Justification object of type TextJustify. This is
//		because text labels with a field length equal to
//		or less than the length of the text label string
//		will never use text justification. In these cases,
//		text justification is completely ignored.
//
//		If the field length is greater than the length of
//		the text label string, text justification must be
//		equal to one of these three valid values:
//
//		    TextJustify(0).Left()
//		    TextJustify(0).Right()
//		    TextJustify(0).Center()
//
//		Users can also specify the abbreviated text
//		justification enumeration syntax as follows:
//
//		    TxtJustify.Left()
//		    TxtJustify.Right()
//		    TxtJustify.Center()
//
//	numLeadingBlankLines		int
//
//		The number of blank lines or 'new lines'
//		inserted above the Leading Solid Line.
//
//	leadingSolidLineChar		string
//
//		The character used to create the Leading
//		Solid Line displayed above the Title
//		Lines.
//
//	numLeadingSolidLines		int
//
//		The Number of Leading Solid Lines to display
//		above the Title Lines.
//
//	numTopTitleBlankLines		int
//
//		The number of blank lines or 'new lines' to
//		insert immediately above the Title Lines
//		Display.
//
//	titleLines					TextLineSpecLinesCollection
//
//		A collection of text line objects containing all
//		specifications necessary to display the Text
//		Title Lines.
//
//		If this parameter is empty with zero Text Line
//		member elements, no error will be returned.
//		However, the user will be responsible for
//		populating the title lines using the 'Add'
//		methods on the returned instance of
//		TextLineSpecTitleMarquee.
//
//	numBottomTitleBlankLines	int
//
//		The number of blank lines or 'new lines' to
//		insert immediately below the Title Lines Display.
//
//	trailingSolidLineChar		string
//
//		The character used to create the Trailing Solid
//		Line displayed below the Title Lines.
//
//	numTrailingSolidLines		int
//
//		The Number of Trailing Solid Lines to display
//		below the Title Lines.
//
//	numTrailingBlankLines		int
//
//		The number of blank lines or 'new lines' inserted
//		after the Trailing Solid Line.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
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
//	TextLineSpecTitleMarquee
//
//		If this method completes successfully, a new
//		fully configured instance of
//		TextLineSpecTitleMarquee will be returned.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (txtLineSpecTitleMarquee *TextLineSpecTitleMarquee) NewAllParams(
	standardTitleLeftMargin string,
	standardTitleRightMargin string,
	standardMaxLineLen int,
	standardTextFieldLen int,
	standardTextJustification TextJustify,
	numLeadingBlankLines int,
	leadingSolidLineChar string,
	numLeadingSolidLines int,
	numTopTitleBlankLines int,
	titleLines TextLineSpecLinesCollection,
	numBottomTitleBlankLines int,
	trailingSolidLineChar string,
	numTrailingSolidLines int,
	numTrailingBlankLines int,
	errorPrefix interface{}) (
	TextLineSpecTitleMarquee,
	error) {

	if txtLineSpecTitleMarquee.lock == nil {
		txtLineSpecTitleMarquee.lock = new(sync.Mutex)
	}

	txtLineSpecTitleMarquee.lock.Lock()

	defer txtLineSpecTitleMarquee.lock.Unlock()

	var newTxtLineTitleMarquee TextLineSpecTitleMarquee

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecTitleMarquee."+
			"NewAllParams()",
		"")

	if err != nil {
		return newTxtLineTitleMarquee, err
	}

	titleMarqueeDto := TextLineTitleMarqueeDto{
		StandardTitleLeftMargin:   standardTitleLeftMargin,
		StandardTitleRightMargin:  standardTitleRightMargin,
		StandardMaxLineLen:        standardMaxLineLen,
		StandardTextFieldLen:      standardTextFieldLen,
		StandardTextJustification: standardTextJustification,
		NumLeadingBlankLines:      numLeadingBlankLines,
		LeadingSolidLineChar:      leadingSolidLineChar,
		NumLeadingSolidLines:      numLeadingSolidLines,
		NumTopTitleBlankLines:     numTopTitleBlankLines,
		NumBottomTitleBlankLines:  numBottomTitleBlankLines,
		TrailingSolidLineChar:     trailingSolidLineChar,
		NumTrailingSolidLines:     numTrailingSolidLines,
		NumTrailingBlankLines:     numTrailingBlankLines,
	}

	if titleLines.GetNumberOfTextLines() == 0 {

		titleMarqueeDto.TitleLines.Empty()

	} else {

		err = titleMarqueeDto.TitleLines.CopyIn(
			&titleLines,
			ePrefix.XCpy("<-titleLines"))

		if err != nil {
			return newTxtLineTitleMarquee, err
		}

	}

	err = new(textLineSpecTitleMarqueeMechanics).
		setTxtLineTitleMarqueeDto(
			&newTxtLineTitleMarquee,
			&titleMarqueeDto,
			ePrefix.XCpy(
				"newTxtLineTitleMarquee"))

	return newTxtLineTitleMarquee, err
}

// NewMarqueeDto
//
// Receives an instance of TextLineTitleMarqueeDto and
// extracts the encapsulated Text Line Marquee
// specifications to construct and return a fully
// configured instance of TextLineTitleMarqueeDto.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	configSpecs					*TextLineTitleMarqueeDto
//
//		A pointer to an instance of
//		TextLineTitleMarqueeDto. The Text Line Title
//		Marquee Data Transfer Object (Dto) is designed to
//		ease and facilitate the data entry operation for
//		text specification parameters required to
//		configure an instance of
//		TextLineSpecTitleMarquee.
//
//		The TextLineTitleMarqueeDto data structure is
//		defined as follows:
//
//		type TextLineTitleMarqueeDto struct {
//
//			StandardTitleLeftMargin string
//				The standard left margin characters applied
//				to all Text Title Lines in the 'TitleLines'
//				array.
//
//			StandardTitleRightMargin string
//				The standard left margin characters applied
//				to all Text Title Lines in the 'TitleLines'
//				array.
//
//			StandardMaxLineLen int
//				The maximum number of characters allowed on
//				a text title line.
//
//			StandardTextFieldLen int
//				The standard field length applied to all
//				Text Title Lines in the 'TitleLines' array.
//
//				'StandardTextFieldLen' defines the length of the
//				text field in which the Title Line string will be
//				displayed. If  'StandardTextFieldLen' is less
//				than the length of the Title Line string, it will
//				be automatically set equal to the Title Line
//				string length.
//
//				To automatically set the value of
//				'StandardTextFieldLen' to the length of the Title
//				Line string, set this parameter to a value of
//				minus one (-1).
//
//				Field Length Examples
//
//					Example-1
//			          Title Line String = "Hello World!"
//						Title Line String Length = 12
//						StandardTextFieldLen = 18
//						StandardTextJustification = TxtJustify.Center()
//						Formatted Title Line String =
//							"   Hello World!   "
//
//					Example-2
//			          Title Line String = "Hello World!"
//						Title Line String Length = 12
//						StandardTextFieldLen = 18
//						StandardTextJustification = TxtJustify.Left()
//						Formatted Title Line String =
//							"Hello World!      "
//
//					Example-3
//			          Title Line String = "Hello World!"
//						Title Line String Length = 12
//						StandardTextFieldLen = -1
//						StandardTextJustification = TxtJustify.Center()
//							// Text Justification Ignored. Field
//							// Length Equals Title Line String Length
//						Formatted Title Line String =
//							"Hello World!"
//
//					Example-4
//			          Title Line String = "Hello World!"
//						Title Line String Length = 12
//						StandardTextFieldLen = 2
//						StandardTextJustification = TxtJustify.Center()
//							// Justification Ignored because Field
//							// Length Less Than Title Line String Length.
//						Formatted Title Line String =
//							"Hello World!"
//
//			StandardTextJustification TextJustify
//				The standard field length applied to all
//				Text Title Lines in the 'TitleLines' array.
//
//				Type 'TextJustify' is an enumeration which
//				specifies the justification of the text field
//				contents string within the text	field length
//				specified by 'StandardTextFieldLen'.
//
//				Text justification can only be evaluated in the
//				context of a text label, field length and a Text
//				Justification object of type TextJustify. This is
//				because text labels with a field length equal to
//				or less than the length of the text label string
//				will never use text justification. In these cases,
//				text justification is completely ignored.
//
//				If the field length is greater than the length of
//				the text label string, text justification must be
//				equal to one of these three valid values:
//
//				    TextJustify(0).Left()
//				    TextJustify(0).Right()
//				    TextJustify(0).Center()
//
//				Users can also specify the abbreviated text
//				justification enumeration syntax as follows:
//
//				    TxtJustify.Left()
//				    TxtJustify.Right()
//				    TxtJustify.Center()
//
//				Text Justification Examples
//
//					Example-1
//			          Title Line String = "Hello World!"
//						Title Line String Length = 12
//						StandardTextFieldLen = 18
//						StandardTextJustification = TxtJustify.Center()
//						Formatted Title Line String =
//							"   Hello World!   "
//
//					Example-2
//			          Title Line String = "Hello World!"
//						Title Line String Length = 12
//						StandardTextFieldLen = 18
//						StandardTextJustification = TxtJustify.Left()
//						Formatted Title Line String =
//							"Hello World!      "
//
//					Example-3
//			          Title Line String = "Hello World!"
//						Title Line String Length = 12
//						StandardTextFieldLen = -1
//						StandardTextJustification = TxtJustify.Center()
//							// Text Justification Ignored. Field
//							// Length Equals Title Line String Length
//						Formatted Title Line String =
//							"Hello World!"
//
//					Example-4
//			          Title Line String = "Hello World!"
//						Title Line String Length = 12
//						StandardTextFieldLen = 2
//						StandardTextJustification = TxtJustify.Center()
//							// Justification Ignored because Field
//							// Length Less Than Title Line String Length.
//						Formatted Title Line String =
//							"Hello World!"
//
//			NumLeadingBlankLines int
//				The number of blank lines or 'new lines'
//				inserted above the Leading Solid Line.
//
//			LeadingSolidLineChar string
//				The character used to create the Leading
//				Solid Line displayed above the Title
//				Lines.
//
//			NumLeadingSolidLines int
//				The Number of Leading Solid Lines to
//				display above the Title Lines.
//
//			NumTopTitleBlankLines int
//				The number of blank lines or 'new lines' to
//				insert immediately above the Title Lines
//				Display.
//
//			TitleLines TextLineSpecLinesCollection
//				A collection of text line objects containing
//				all specifications necessary to display the
//				Text Title Lines.
//
//			NumBottomTitleBlankLines int
//				The number of blank lines or 'new lines' to
//				insert immediately below the Title Lines
//				Display.
//
//			TrailingSolidLineChar string
//				The character used to create the Trailing
//				Solid Line displayed below the Title
//				Lines.
//
//			NumTrailingSolidLines int
//				The Number of Trailing Solid Lines to
//				display below the Title Lines.
//
//			NumTrailingBlankLines int
//				The number of blank lines or 'new lines'
//				inserted after the Trailing Solid Line.
//		}
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
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
//	TextLineSpecTitleMarquee
//
//		If this method completes successfully, a new
//		fully configured instance of
//		TextLineSpecTitleMarquee will be returned.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (txtLineSpecTitleMarquee *TextLineSpecTitleMarquee) NewMarqueeDto(
	configSpecs *TextLineTitleMarqueeDto,
	errorPrefix interface{}) (
	TextLineSpecTitleMarquee,
	error) {

	if txtLineSpecTitleMarquee.lock == nil {
		txtLineSpecTitleMarquee.lock = new(sync.Mutex)
	}

	txtLineSpecTitleMarquee.lock.Lock()

	defer txtLineSpecTitleMarquee.lock.Unlock()

	var newTxtLineTitleMarquee TextLineSpecTitleMarquee

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecTitleMarquee."+
			"NewMarqueeDto()",
		"")

	if err != nil {
		return newTxtLineTitleMarquee, err
	}

	err = new(textLineSpecTitleMarqueeMechanics).
		setTxtLineTitleMarqueeDto(
			&newTxtLineTitleMarquee,
			configSpecs,
			ePrefix.XCpy(
				"newTxtLineTitleMarquee<-"+
					"configSpecs"))

	return newTxtLineTitleMarquee, err
}

//	Read
//
//	Implements the io.Reader interface for type
//	TextLineSpecTitleMarquee.
//
//	The formatted text line string generated by the
//	current instance of TextLineSpecTitleMarquee will be
//	written to the byte buffer 'p'. If the length of 'p'
//	is less than the length of the formatted text line
//	string, multiple calls to this method will write the
//	remaining unread characters to the byte buffer 'p'.
//
//	Read() supports buffered 'read' operations.
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
//	(non-nil) error from the same call or return the
//	error (and n == 0) from a subsequent call. An
//	instance of this general case is that a Reader
//	returning a non-zero number of bytes at the end of
//	the input stream may return either err == EOF or
//	err == nil. The next read operation should return 0,
//	EOF.
//
//	Callers should always process the n > 0 bytes returned
//	before considering the error err. Doing so correctly
//	handles I/O errors that happen after reading some bytes
//	and also both of the allowed EOF behaviors.
//
//	The last read operation performed on the formatted text
//	string will always return n==0 and err==io.EOF.
//
//	This method fulfills requirements of the
//	ITextLineSpecification interface.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	p							[]byte
//
//		The byte buffer into which the formatted text
//		line string generated by the current
//		TextLineSpecTitleMarquee instance will be
//		written.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	n							int
//
//		The number of bytes written to byte buffer 'p'.
//
//		Read() reads up to len(p) bytes into p. It
//		returns the number of bytes read
//		(0 <= n <= len(p)) and any error encountered.
//		Even if Read() returns n < len(p), it may use all
//		of 'p' as scratch space during the call. If some
//		data is available but not len(p) bytes, Read()
//		conventionally returns what is available instead
//		of waiting for more.
//
//	err							error
//
//		If this method completes successfully, this
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		When Read() encounters an error or end-of-file
//		condition after successfully reading n > 0 bytes,
//		it returns the number of bytes read. It may
//		return the (non-nil) error from the same call or
//		return the error (and n == 0) from a subsequent
//		call. An instance of this general case is that a
//		Reader returning a non-zero number of bytes at
//		the end of the input stream may return either
//		err == EOF or err == nil. The next read operation
//		should return 0, EOF.
//
// ----------------------------------------------------------------
//
// # Usage
//
//	Example # 1
//
//	p := make([]byte, 50)
//
//	var n, readBytesCnt int
//	sb := strings.Builder{}
//
//	for {
//
//	  n,
//	  err = txtTitleMarquee.Read(p)
//
//	  if n == 0 {
//	    break
//	  }
//
//	  sb.Write(p[:n])
//	  readBytesCnt += n
//	}
//
//	if err != nil &&
//	  err != io.EOF {
//	   return fmt.Error(
//	    "Error Returned From txtTitleMarquee.Read(p)\n"+
//	    "Error = \n%v\n",
//	     err.Error())
//	}
//
//	fmt.Printf("Text Line String: %s\n",
//	              sb.String())
//
//	fmt.Printf("Number of bytes Read: %v\n",
//	              readBytesCnt)
//
//	Example # 2
//
//	p := make([]byte, 50)
//
//	var n, readBytesCnt int
//	var actualStr string
//
//	for {
//
//	  n,
//	  err = txtTitleMarquee.Read(p)
//
//	  if n == 0 {
//	    break
//	  }
//
//	  actualStr += string(p[:n])
//	  readBytesCnt += n
//	}
//
//	if err != nil &&
//	  err != io.EOF {
//	   return fmt.Error(
//	    "Error Returned From txtTitleMarquee.Read(p)\n"+
//	    "Error = \n%v\n",
//	     err.Error())
//	}
//
//	fmt.Printf("Text Line String: %v\n",
//	              actualStr)
//
//	fmt.Printf("Number of bytes Read: %v\n",
//	              readBytesCnt)
func (txtLineSpecTitleMarquee *TextLineSpecTitleMarquee) Read(
	p []byte) (
	n int,
	err error) {

	if txtLineSpecTitleMarquee.lock == nil {
		txtLineSpecTitleMarquee.lock = new(sync.Mutex)
	}

	txtLineSpecTitleMarquee.lock.Lock()

	defer txtLineSpecTitleMarquee.lock.Unlock()

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TextLineSpecTitleMarquee.Read()",
		"")

	if txtLineSpecTitleMarquee.textLineReader == nil {

		strBuilder := strings.Builder{}

		_,
			_,
			err = new(textLineSpecTitleMarqueeMolecule).
			getFormattedText(
				&strBuilder,
				txtLineSpecTitleMarquee,
				ePrefix.XCpy("txtLineSpecTitleMarquee"))

		if err != nil {
			return n, err
		}

		txtLineSpecTitleMarquee.textLineReader =
			strings.NewReader(strBuilder.String())

		if txtLineSpecTitleMarquee.textLineReader == nil {
			err = fmt.Errorf("%v\n"+
				"Error: strings.NewReader(formattedText)\n"+
				"returned a nil pointer.\n"+
				"txtLineSpecTitleMarquee.textLineReader == nil\n",
				ePrefix.String())

			return n, err
		}

		strBuilder.Reset()
	}

	n,
		err = new(textLineSpecTitleMarqueeMolecule).
		readBytes(
			txtLineSpecTitleMarquee.textLineReader,
			p,
			ePrefix.XCpy(
				"p -> "+
					"txtLineSpecTitleMarquee.textLineReader"))

	if err == io.EOF {

		txtLineSpecTitleMarquee.textLineReader = nil

	}

	return n, err
}

//	ReaderInitialize
//
//	This method will reset the internal member variable
//	'TextLineSpecTitleMarquee.textLineReader' to its
//	initial zero state of 'nil'. Effectively, this resets
//	the internal strings.Reader object for use in future
//	read operations.
//
//	This method is rarely used or needed. It provides a
//	means of reinitializing the internal strings.Reader
//	object in case an error occurs during a read
//	operation initiated by method
//	TextLineSpecTitleMarquee.Read().
//
//	Calling this method cleans up the residue from an
//	aborted read operation and prepares the
//	strings.Reader object for future read operations.
//
//	If any errors are returned by method
//	TextLineSpecTitleMarquee.Read() which are NOT equal
//	to io.EOF, call this method,
//	TextLineSpecTitleMarquee.ReaderInitialize(), to reset
//	and prepare the internal reader for future read
//	operations.
//
//	This method fulfills requirements of the
//	ITextLineSpecification interface.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	NONE
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	NONE
func (txtLineSpecTitleMarquee *TextLineSpecTitleMarquee) ReaderInitialize() {

	if txtLineSpecTitleMarquee.lock == nil {
		txtLineSpecTitleMarquee.lock = new(sync.Mutex)
	}

	txtLineSpecTitleMarquee.lock.Lock()

	defer txtLineSpecTitleMarquee.lock.Unlock()

	txtLineSpecTitleMarquee.textLineReader = nil

	return
}

// SetMarqueeLinesCollections
//
// Deletes and replaces one of the Marquee Line
// Collections contained in the current instance of
// TextLineSpecTitleMarquee.
//
// Type TextLineSpecTitleMarquee encapsulates three types
// of text lines used in generating title marquees:
//
//  1. Leading Marquee Lines
//     Usually consists of leading blank lines
//     and solid lines.
//
//  2. Title Lines
//     Consists entirely of text strings functioning
//     as the main title lines.
//
//  3. Trailing Marquee Lines
//     Usually consists of trailing blank lines
//     and solid lines.
//
// This method replaces the pre-existing Leading Marquee
// Lines, Title Lines or Trailing Marquee Lines
// Collection depending on setting for input parameter
// 'titleMarqueeLineType'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and replace one of the
//	Marquee Line Collections contained in the current
//	instance of TextLineSpecTitleMarquee. The precise
//	collection to be deleted and replaced is determined
//	by input parameter 'titleMarqueeLineType'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	leadingMarqueeLines			TextLineSpecLinesCollection
//
//		An instance of TextLineSpecLinesCollection
//		containing specifications for creating text
//		lines. The text lines will replace the
//		pre-existing Leading Marquee Text Lines
//		Collection encapsulated by the current instance
//		of TextLineSpecTitleMarquee.
//
//		If 'leadingMarqueeLines' contains invalid data
//		values, an error will be returned.
//
//	titleMarqueeLineType		TextTileLineType
//
//		Type TextTileLineType is an enumeration of
//		Title Marquee Text Line Types. This parameter
//		determines which text line collection will
//		be deleted and reset with the new Marquee
//		Text Collection passed by input parameter,
//		'leadingMarqueeLines'.
//
//		If this parameter is not set to one of the
//		following valid values, an error will be
//		returned.
//
//		Formal TextTileLineType Syntax
//
//			TextTileLineType(0).LeadingMarqueeLine()
//			TextTileLineType(0).TitleLine()
//			TextTileLineType(0).TrailingMarqueeLine()
//
//		Abbreviated TextTileLineType Syntax
//
//			TitleLineType.LeadingMarqueeLine()
//			TitleLineType.TitleLine()
//			TitleLineTypeTrailingMarqueeLine()
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
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
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (txtLineSpecTitleMarquee *TextLineSpecTitleMarquee) SetMarqueeLinesCollections(
	marqueeTextLines *TextLineSpecLinesCollection,
	titleMarqueeLineType TextTileLineType,
	errorPrefix interface{}) error {

	if txtLineSpecTitleMarquee.lock == nil {
		txtLineSpecTitleMarquee.lock = new(sync.Mutex)
	}

	txtLineSpecTitleMarquee.lock.Lock()

	defer txtLineSpecTitleMarquee.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecTitleMarquee."+
			"SetMarqueeLinesCollections()",
		"")

	if err != nil {
		return err
	}

	if marqueeTextLines == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'marqueeTextLines' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	_,
		err = new(textLineSpecLinesCollectionAtom).
		testValidityOfTextLinesCollection(
			marqueeTextLines,
			ePrefix.XCpy(
				"leadingMarqueeLines invalid!"))

	if err != nil {
		return err
	}

	var txtLineCollection *TextLineSpecLinesCollection

	var txtLineCollectionName string

	switch titleMarqueeLineType {

	case TitleLineType.LeadingMarqueeLine():

		txtLineCollection = &txtLineSpecTitleMarquee.leadingMarqueeLines

		txtLineCollectionName = "leadingMarqueeLines"

	case TitleLineType.TitleLine():

		txtLineCollection = &txtLineSpecTitleMarquee.titleLines

		txtLineCollectionName = "titleLines"

	case TitleLineType.TrailingMarqueeLine():

		txtLineCollection = &txtLineSpecTitleMarquee.trailingMarqueeLines

		txtLineCollectionName = "trailingMarqueeLines"

	default:

		err := fmt.Errorf("%v\n"+
			"Error: Input parameter 'titleMarqueeLineType' is invalid!\n"+
			" titleMarqueeLineType string value = '%v'\n"+
			"titleMarqueeLineType integer value = '%v'\n",
			ePrefix.String(),
			titleMarqueeLineType.String(),
			titleMarqueeLineType.XValueInt())

		return err
	}

	new(textLineSpecLinesCollectionAtom).
		emptyCollection(txtLineCollection)

	err = txtLineCollection.CopyIn(
		marqueeTextLines,
		ePrefix.XCpy(
			fmt.Sprintf("txtLineSpecTitleMarquee.%v",
				txtLineCollectionName)))

	return err
}

//	String
//
//	Returns the formatted text generated by this Text
//	Line Specification for Title Marquee for use in
//	file output, screen displays and printing.
//
//	If an error occurs, the error message will be
//	included in the returned string.
//
//	This method is similar to
//	TextLineSpecTitleMarquee.GetFormattedText() with the
//	sole difference being that this method does not
//	return an error.
//
//	This method fulfills requirements of the
//	ITextLineSpecification interface.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	NONE
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	string
//
//		The formatted text lines generated by the current
//		instance of TextLineSpecTitleMarquee.
func (txtLineSpecTitleMarquee TextLineSpecTitleMarquee) String() string {

	if txtLineSpecTitleMarquee.lock == nil {
		txtLineSpecTitleMarquee.lock = new(sync.Mutex)
	}

	txtLineSpecTitleMarquee.lock.Lock()

	defer txtLineSpecTitleMarquee.lock.Unlock()

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TextLineSpecStandardLine.GetFormattedText()",
		"")

	strBuilder := strings.Builder{}

	var err error
	var formattedText string

	_,
		_,
		err = new(textLineSpecTitleMarqueeMolecule).
		getFormattedText(
			&strBuilder,
			&txtLineSpecTitleMarquee,
			ePrefix.XCpy("txtLineSpecTitleMarquee"))

	if err != nil {
		formattedText = fmt.Sprintf("%v\n",
			err.Error())
	} else {
		formattedText = strBuilder.String()
	}

	return formattedText
}

//	TextBuilder
//
//	Configures the line of text produced by this instance
//	of TextLineSpecTitleMarquee, and writes it to an
//	instance of strings.Builder.
//
//	This method fulfills requirements of the
//	ITextLineSpecification interface.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	strBuilder					*strings.Builder
//
//		A pointer to an instance of *strings.Builder. The
//		formatted text characters produced by this method
//		will be written to this instance of
//		strings.Builder.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
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
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (txtLineSpecTitleMarquee *TextLineSpecTitleMarquee) TextBuilder(
	strBuilder *strings.Builder,
	errorPrefix interface{}) error {

	if txtLineSpecTitleMarquee.lock == nil {
		txtLineSpecTitleMarquee.lock = new(sync.Mutex)
	}

	txtLineSpecTitleMarquee.lock.Lock()

	defer txtLineSpecTitleMarquee.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecTitleMarquee."+
			"TextBuilder()",
		"")

	if err != nil {
		return err
	}

	_,
		_,
		err = new(textLineSpecTitleMarqueeMolecule).
		getFormattedText(
			strBuilder,
			txtLineSpecTitleMarquee,
			ePrefix.XCpy(
				"strBuilder<-"+
					"txtLineSpecTitleMarquee"))

	return err
}

//	TextLineSpecName
//
//	Returns a string specifying the name of this Text
//	Line Specification (TextLineSpecTitleMarquee).
//
//	This method fulfills requirements of the
//	ITextLineSpecification interface.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	NONE
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	NONE
func (txtLineSpecTitleMarquee *TextLineSpecTitleMarquee) TextLineSpecName() string {

	if txtLineSpecTitleMarquee.lock == nil {
		txtLineSpecTitleMarquee.lock = new(sync.Mutex)
	}

	txtLineSpecTitleMarquee.lock.Lock()

	defer txtLineSpecTitleMarquee.lock.Unlock()

	return "TextLineSpecTitleMarquee"
}

// TextTypeName
//
// Returns a string specifying the type of Text Line
// Specification (TextLineSpecTitleMarquee).
//
// This method fulfills requirements of the
// ITextSpecification interface.
func (txtLineSpecTitleMarquee *TextLineSpecTitleMarquee) TextTypeName() string {

	if txtLineSpecTitleMarquee.lock == nil {
		txtLineSpecTitleMarquee.lock = new(sync.Mutex)
	}

	txtLineSpecTitleMarquee.lock.Lock()

	defer txtLineSpecTitleMarquee.lock.Unlock()

	return "TextLineSpecTitleMarquee"
}
