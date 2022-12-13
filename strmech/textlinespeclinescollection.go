package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
)

// TextLineSpecLinesCollection
//
// A collection of objects implementing the
// ITextLineSpecification interface.
type TextLineSpecLinesCollection struct {
	textLines []ITextLineSpecification
	lock      *sync.Mutex
}

//	AddBlankLine
//
//	Creates an instance of TextLineSpecBlankLines and
//	adds it to the Text Line collection maintained by
//	this instance of TextLineSpecLinesCollection.
//
//	The TextLineSpecBlankLines type is a specialized
//	form of text line specification which is used to
//	create one or more blank lines of text.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numOfBlankLines				int
//
//		The number of blank lines which will be generated
//		by the newly created instance of
//		TextLineSpecBlankLines.
//
//		If input parameter 'numOfBlankLines' is less than
//		one (1), it is invalid and an error will be
//		returned.
//
//		If input parameter 'numOfBlankLines' is greater
//		than one-million (1,000,000), it is invalid and
//		an error will be returned.
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
func (txtLinesSpecCol *TextLineSpecLinesCollection) AddBlankLine(
	numOfBlankLines int,
	errorPrefix interface{}) error {

	if txtLinesSpecCol.lock == nil {
		txtLinesSpecCol.lock = new(sync.Mutex)
	}

	txtLinesSpecCol.lock.Lock()

	defer txtLinesSpecCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecLinesCollection."+
			"AddBlankLine()",
		"")

	if err != nil {
		return err
	}

	var newBlankLine *TextLineSpecBlankLines

	newBlankLine,
		err = new(TextLineSpecBlankLines).
		NewPtrDefaultBlankLines(
			numOfBlankLines,
			ePrefix.XCpy(
				"newBlankLine<-"))

	if err != nil {
		return err
	}

	err = new(textLineSpecLinesCollectionNanobot).addTextLine(
		txtLinesSpecCol,
		newBlankLine,
		ePrefix.XCpy(
			"txtLinesSpecCol<-newBlankLine"))

	return err
}

// AddSolidLine
//
// Adds a solid line to the Text Specification Lines
// Collection maintained by the current instance of
// TextLineSpecLinesCollection.
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
//		used to terminate the solid text line.
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
//		TextLineSpecLinesCollection.
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
func (txtLinesSpecCol *TextLineSpecLinesCollection) AddSolidLine(
	leftMarginStr string,
	solidLineChars string,
	solidLineCharsRepeatCount int,
	rightMarginStr string,
	lineTerminator string,
	turnLineTerminatorOff bool,
	numOfSolidLines int,
	errorPrefix interface{}) error {

	if txtLinesSpecCol.lock == nil {
		txtLinesSpecCol.lock = new(sync.Mutex)
	}

	txtLinesSpecCol.lock.Lock()

	defer txtLinesSpecCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecLinesCollection."+
			"AddSolidLine()",
		"")

	if err != nil {
		return err
	}

	if numOfSolidLines < 1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numOfSolidLines' is invalid!\n"+
			"'numOfSolidLines' has a value less than one (1).\n"+
			"'numOfSolidLines' controls the number of solid lines\n"+
			"which will be generated.\n"+
			"numOfSolidLines = '%v'\n",
			ePrefix.String(),
			numOfSolidLines)

		return err
	}

	var txtSolidLine TextLineSpecSolidLine

	txtSolidLine,
		err = TextLineSpecSolidLine{}.
		NewSolidLineAllParms(
			leftMarginStr,
			rightMarginStr,
			solidLineChars,
			solidLineCharsRepeatCount,
			lineTerminator,
			turnLineTerminatorOff,
			ePrefix.XCpy(
				"txtSolidLine"))

	for i := 0; i < numOfSolidLines; i++ {

		var copyTxtSolidLine TextLineSpecSolidLine

		copyTxtSolidLine,
			err = txtSolidLine.CopyOut(
			ePrefix.XCpy(
				fmt.Sprintf("txtSolildLine[%v]",
					i)))

		err = new(textLineSpecLinesCollectionNanobot).
			addTextLine(
				txtLinesSpecCol,
				&copyTxtSolidLine,
				ePrefix.XCpy(
					"txtLinesSpecCol<-copyTxtSolidLine"))

	}

	return err
}

// AddStdLineColumns
//
// Adds a standard text line to the Text Specification
// Lines Collection maintained by the current instance of
// TextLineSpecLinesCollection.
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
//		instance of TextLineSpecLinesCollection.
//
//		By default, each line of text generated by
//		TextLineSpecStandardLine will be terminated with
//		a new line character ('\n'). However, this
//		parameter allows the user to specify the
//		character or characters to be used as a line
//		termination sequence for each line of text added
//		to the Text Specification Lines Collection.
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
//		TextLineSpecLinesCollection.
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
func (txtLinesSpecCol *TextLineSpecLinesCollection) AddStdLineColumns(
	lineTerminator string,
	turnLineTerminatorOff bool,
	errorPrefix interface{},
	textFieldColumns ...ITextFieldFormatDto) error {

	if txtLinesSpecCol.lock == nil {
		txtLinesSpecCol.lock = new(sync.Mutex)
	}

	txtLinesSpecCol.lock.Lock()

	defer txtLinesSpecCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecLinesCollection."+
			"AddBlankLine()",
		"")

	if err != nil {
		return err
	}

	var newStdLine TextLineSpecStandardLine

	newStdLine,
		err = TextLineSpecStandardLine{}.NewStdLineColumns(
		lineTerminator,
		turnLineTerminatorOff,
		ePrefix.XCpy(
			"newStdLine<-textFieldColumns..."),
		textFieldColumns...)

	if err != nil {
		return err
	}

	err = new(textLineSpecLinesCollectionNanobot).addTextLine(
		txtLinesSpecCol,
		&newStdLine,
		ePrefix.XCpy(
			"txtLinesSpecCol<-newStdLine"))

	return err
}

// AddTextLineSpec - Adds a ITextLineSpecification object to the
// end of the Text Line collection maintained by this instance of
// TextLineSpecLinesCollection.
//
// A deep copy of this ITextLineSpecification object ('textLine')
// is appended to the end of the array of ITextLineSpecification
// objects maintained by this instance of
// TextLineSpecLinesCollection.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	textLine                   ITextLineSpecification
//	   - A text line object which implements the
//	     ITextLineSpecification interface. A deep copy of this
//	     object will be added to the text lines collection
//	     maintained by this instance of
//	     TextLineSpecLinesCollection.
//
//	     If member variable data values contained in this
//	     'textLine' parameter are found to be invalid, an error
//	     will be returned.
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings containing
//	                    error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	                        ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	err                        error
//	   - If the method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (txtLinesSpecCol *TextLineSpecLinesCollection) AddTextLineSpec(
	textLine ITextLineSpecification,
	errorPrefix interface{}) (
	err error) {

	if txtLinesSpecCol.lock == nil {
		txtLinesSpecCol.lock = new(sync.Mutex)
	}

	txtLinesSpecCol.lock.Lock()

	defer txtLinesSpecCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecLinesCollection."+
			"AddTextLineSpec()",
		"")

	if err != nil {
		return err
	}

	err = new(textLineSpecLinesCollectionNanobot).
		addTextLine(
			txtLinesSpecCol,
			textLine,
			ePrefix.XCpy(
				"txtLinesSpecCol<-textLine"))

	return err
}

// CopyIn - Copies the text line collection from an incoming
// instance of TextLineSpecLinesCollection
// ('incomingTxtLinesCol') to the current
// TextLineSpecLinesCollection instance ('txtLinesCol').
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// All data values in current TextLineSpecLinesCollection instance
// ('txtLinesCol') will be deleted and overwritten.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	incomingTxtLinesCol        *TextLineSpecLinesCollection
//	   - A pointer to an instance of TextLineSpecLinesCollection.
//	     This method will NOT change the data values contained in
//	     this instance.
//
//	     All text line collection member elements in this
//	     TextLineSpecLinesCollection instance will be copied to the
//	     current TextLineSpecLinesCollection instance ('txtLinesCol').
//
//	     If 'incomingTextLineCol' contains invalid member data
//	     variables, this method will return an error.
//
//	     If 'incomingTextLineCol' contains an empty, or zero
//	     length, Text Lines Collection, an error will be returned.
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	                    containing error prefix and error context
//	                    information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	                        ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If the method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (txtLinesSpecCol *TextLineSpecLinesCollection) CopyIn(
	incomingTxtLinesCol *TextLineSpecLinesCollection,
	errorPrefix interface{}) error {

	if txtLinesSpecCol.lock == nil {
		txtLinesSpecCol.lock = new(sync.Mutex)
	}

	txtLinesSpecCol.lock.Lock()

	defer txtLinesSpecCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecLinesCollection.CopyIn()",
		"")

	if err != nil {
		return err
	}

	return new(textLineSpecLinesCollectionNanobot).
		copyIn(
			txtLinesSpecCol,
			incomingTxtLinesCol,
			ePrefix.XCpy(
				"incomingTxtLinesCol->"+
					"txtLinesSpecCol"))
}

// CopyOut - Returns a deep copy of the current
// TextLineSpecLinesCollection instance.
//
// If the current TextLineSpecLinesCollection instance contains
// invalid member variables, this method will return an error.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings containing
//	                    error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	                        ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	TextLineSpecLinesCollection
//	   - If this method completes successfully and no errors are
//	     encountered, this parameter will return a deep copy of the
//	     current TextLineSpecLinesCollection instance.
//
//
//	error
//	   - If the method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (txtLinesSpecCol *TextLineSpecLinesCollection) CopyOut(
	errorPrefix interface{}) (
	TextLineSpecLinesCollection,
	error) {

	if txtLinesSpecCol.lock == nil {
		txtLinesSpecCol.lock = new(sync.Mutex)
	}

	txtLinesSpecCol.lock.Lock()

	defer txtLinesSpecCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecLinesCollection.CopyOut()",
		"")

	if err != nil {
		return TextLineSpecLinesCollection{}, err
	}

	return new(textLineSpecLinesCollectionNanobot).
		copyOut(
			txtLinesSpecCol,
			ePrefix.XCpy(
				"txtLinesSpecCol->"))
}

// DeleteTextLineMember - Deletes a member of the Text Lines
// collection encapsulated by the current instance of
// TextLineSpecLinesCollection.
//
// Input parameter 'zeroBasedIndex' is used to specify the
// collection member which will be deleted. If the operation
// is completed successfully, the total number of member elements
// in the collection will be reduced by one (1).
//
// If input parameter 'zeroBasedIndex' is less than zero or greater
// than the last member element in collection, an error will be
// returned. Also, if this method is called on an empty collection,
// an error will be returned.
//
// Remember that indexes in the Text Lines collection are zero
// based. This means the first element in the collection is index
// zero.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
// If you delete the last element in the collection, the current
// instance TextLineSpecLinesCollection will be rendered invalid,
// and cannot be used until more Text Line elements are added to
// the collection
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	zeroBasedIndex             int
//	   - Specifies the index of the member element in the Text
//	     Lines collection which will be deleted. If this input
//	     parameter is found to be invalid or if the Text Lines
//	     collection is empty, an error will be returned.
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	                    containing error prefix and error context
//	                    information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	                        ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If the method completes successfully and no errors are
//	     encountered, this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (txtLinesSpecCol *TextLineSpecLinesCollection) DeleteTextLineMember(
	zeroBasedIndex int,
	errorPrefix interface{}) error {

	if txtLinesSpecCol.lock == nil {
		txtLinesSpecCol.lock = new(sync.Mutex)
	}

	txtLinesSpecCol.lock.Lock()

	defer txtLinesSpecCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecLinesCollection."+
			"DeleteTextLineMember()",
		"")

	if err != nil {
		return err
	}

	err = new(textLineSpecLinesCollectionElectron).
		deleteTextLineElement(
			txtLinesSpecCol,
			zeroBasedIndex,
			ePrefix.XCpy(
				fmt.Sprintf(
					"Delete Element txtLinesSpecCol[%v]",
					zeroBasedIndex)))

	return err
}

// Empty - Empties the text line collection and resets all member
// variables to their initial or zero values.
//
// Call this method when you intend to delete the
// TextLineSpecLinesCollection permanently as it will not be
// available for immediate reuse.
//
// If you wish to delete the text line collection and immediately
// reuse this TextLineSpecLinesCollection instance, use method
// TextLineSpecLinesCollection.EmptyTextLines() instead.
func (txtLinesSpecCol *TextLineSpecLinesCollection) Empty() {

	if txtLinesSpecCol.lock == nil {
		txtLinesSpecCol.lock = new(sync.Mutex)
	}

	txtLinesSpecCol.lock.Lock()

	new(textLineSpecLinesCollectionAtom).
		emptyCollection(txtLinesSpecCol)

	txtLinesSpecCol.lock.Unlock()

	txtLinesSpecCol.lock = nil

	return
}

// EmptyTextLines - Empties the text line collection and resets all member
// variables to their initial or zero values.
//
// This method only deletes the current text line collection. This
// means that the TextLineSpecLinesCollection instance is
// immediately for reuse and new text lines may be added to the now
// empty collection.
func (txtLinesSpecCol *TextLineSpecLinesCollection) EmptyTextLines() {

	if txtLinesSpecCol.lock == nil {
		txtLinesSpecCol.lock = new(sync.Mutex)
	}

	txtLinesSpecCol.lock.Lock()

	defer txtLinesSpecCol.lock.Unlock()

	new(textLineSpecLinesCollectionAtom).
		emptyCollection(txtLinesSpecCol)

	return
}

// Equal - Receives a pointer to another instance of
// TextLineSpecLinesCollection and proceeds to compare the member
// variables to those of the current TextLineSpecLinesCollection
// instance in order to determine if they are equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables of both instances are equal in
// all respects, this flag is set to 'true'. Otherwise, this method
// returns 'false'.
func (txtLinesSpecCol *TextLineSpecLinesCollection) Equal(
	textLinesCol02 *TextLineSpecLinesCollection) bool {

	if txtLinesSpecCol.lock == nil {
		txtLinesSpecCol.lock = new(sync.Mutex)
	}

	txtLinesSpecCol.lock.Lock()

	defer txtLinesSpecCol.lock.Unlock()

	return new(textLineSpecLinesCollectionAtom).
		equalCollections(
			txtLinesSpecCol,
			textLinesCol02)
}

// GetNumberOfTextLines - Returns the number of text lines
// encapsulated by the current TextLineSpecLinesCollection
// instance.
//
// Analyzing the number of text lines in the collection provides
// verification that text lines exist and are ready for formatting.
// Once properly formatted text lines may be presented for text
// display, file output or printing.
func (txtLinesSpecCol *TextLineSpecLinesCollection) GetNumberOfTextLines() int {

	if txtLinesSpecCol.lock == nil {
		txtLinesSpecCol.lock = new(sync.Mutex)
	}

	txtLinesSpecCol.lock.Lock()

	defer txtLinesSpecCol.lock.Unlock()

	return len(txtLinesSpecCol.textLines)
}

//	GetFormattedText
//
//	Generates formatted text strings for all member
//	elements in the Text Line Specification Collection
//	maintained by the current instance of
//	TextLineSpecLinesCollection. These formatted text
//	strings are then written to a String Builder
//	(strBuilder) passed as an input parameter.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	strBuilder					*strings.Builder
//
//			A pointer to an instance of strings.Builder.
//			The formatted text strings generated from
//			the current instance of
//			TextLineSpecLinesCollection
//			('txtLinesSpecCol') will be written to this
//			instance of strings.Builder.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
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
//	maxLineLen					int
//
//		This parameter returns the length of the longest
//		text line generated from this collection of Text
//		Line Specifications ('textLinesCol').
//
//	err							error
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
func (txtLinesSpecCol *TextLineSpecLinesCollection) GetFormattedText(
	strBuilder *strings.Builder,
	errorPrefix interface{}) (
	maxLineLen int,
	err error) {

	if txtLinesSpecCol.lock == nil {
		txtLinesSpecCol.lock = new(sync.Mutex)
	}

	txtLinesSpecCol.lock.Lock()

	defer txtLinesSpecCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecLinesCollection."+
			"GetFormattedText()",
		"")

	if err != nil {
		return maxLineLen, err
	}

	return new(textLineSpecLinesCollectionNanobot).
		getFormattedText(
			strBuilder,
			txtLinesSpecCol,
			ePrefix.XCpy(
				"txtLinesSpecCol"))
}

//	GetFmtTextStrArray
//
//	Generates formatted text strings from the Text Line
//	Specifications collection maintained by the current
//	instance of TextLineSpecLinesCollection. These text
//	strings are then returned as individual elements in
//	a string array.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
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
//	StringArrayDto
//
//		If this method completes successfully, an
//		instance of StringArrayDto will be returned
//		encapsulating a string array containing the
//		formatted text strings generated from the
//		Text Line Specification objects contained in
//		the current instance of
//		TextLineSpecLinesCollection.
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
func (txtLinesSpecCol *TextLineSpecLinesCollection) GetFmtTextStrArray(
	errorPrefix interface{}) (
	StringArrayDto,
	error) {

	if txtLinesSpecCol.lock == nil {
		txtLinesSpecCol.lock = new(sync.Mutex)
	}

	txtLinesSpecCol.lock.Lock()

	defer txtLinesSpecCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	txtStrArray := new(StringArrayDto).New()

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecLinesCollection."+
			"GetFmtTextStrArray()",
		"")

	if err != nil {
		return txtStrArray, err
	}

	var strArrayDto StringArrayDto

	err = new(textLineSpecLinesCollectionNanobot).
		getFormattedTextStrArray(
			&strArrayDto,
			txtLinesSpecCol,
			ePrefix.XCpy(
				"txtLinesSpecCol"))

	return strArrayDto, err
}

// GetTextLine - Returns a deep copy of the Text Line Collection
// member element specified by input parameter, 'zeroBasedIndex'.
//
// If the Text Line collection maintained by the current
// TextLineSpecLinesCollection instance is empty (contains zero
// elements), an error will be returned.
//
// Remember that indexes in the Text Lines collection are zero
// based. This means the first element in the collection is index
// zero.
//
// If input parameter 'zeroBasedIndex' is less than zero or greater
// than the last member element in the collection, an error will be
// returned.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
// This method ( GetTextField() ) is functionally equivalent to
// method:
//
//	TextLineSpecLinesCollection.PeekAtTextField()
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	zeroBasedIndex             int
//	   - Specifies the index of the member element in the Text
//	     Lines collection which will be returned as a deep copy
//	     of the original. If this input parameter is found to be
//	     invalid or if the Text Lines collection is empty, an
//	     error will be returned.
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	                    containing error prefix and error context
//	                    information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	                        ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	ITextLineSpecification
//	   - If this method completes successfully, a deep copy of the
//	     Text Line member element specified by input parameter
//	     'zeroBasedIndex' will be returned.
//
//
//	error
//	   - If the method completes successfully and no errors are
//	     encountered, this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//	When casting ITextLineSpecification returned from this method,
//	use the following syntax to cast the interface object to a
//	concrete type.
//
//	It is necessary to cast the interface object ('iTxtLineSpec')
//	as a pointer to the concrete type ('stdLine'). This is because
//	the concrete type uses methods with pointer receivers.
//
//	------------------------------------------------------------
//	   var iTxtLineSpec ITextLineSpecification
//
//	   iTxtLineSpec,
//	   err = txtLinesCol01.GetTextLine(
//	           2,  // Return Text Line at index '2'
//	           ePrefix.XCpy(
//	           "txtLinesCol01"))
//
//	   if err != nil {
//	     return err
//	   }
//
//	   var stdLine *TextLineSpecStandardLine
//
//	   var ok bool
//
//	   stdLine, ok = iTxtLineSpec.(*TextLineSpecStandardLine)
//
//	   if !ok {
//
//	     err = fmt.Errorf("%v - Error\n"+
//	     "stdLine, ok := iTxtLineSpec.(*TextLineSpecStandardLine)\n"+
//	     "Expected return of type 'TextLineSpecStandardLine'.\n"+
//	     "HOWEVER, THAT TYPE WAS NOT RETURNED!\n",
//	     ePrefix.String())
//
//	     return err
//	   }
//
//	   // 'stdLine' is now available for use
//	   // as a concrete object.
//	   stdLineLen := stdLine.GetSingleLineLength()
func (txtLinesSpecCol *TextLineSpecLinesCollection) GetTextLine(
	zeroBasedIndex int,
	errorPrefix interface{}) (
	ITextLineSpecification,
	error) {

	if txtLinesSpecCol.lock == nil {
		txtLinesSpecCol.lock = new(sync.Mutex)
	}

	txtLinesSpecCol.lock.Lock()

	defer txtLinesSpecCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	var iTextLineSpec ITextLineSpecification

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecLinesCollection."+
			"GetTextLine()",
		"")

	if err != nil {
		return iTextLineSpec, err
	}

	iTextLineSpec,
		err = new(textLineSpecLinesCollectionAtom).
		peekPopTextLine(
			txtLinesSpecCol,
			zeroBasedIndex,
			false,
			ePrefix.XCpy(
				fmt.Sprintf(
					"txtLinesSpecCol[%v]",
					zeroBasedIndex)))

	return iTextLineSpec, err
}

// GetTextLineCollection - Returns a deep copy of the text lines
// contained in the current TextLineSpecLinesCollection instance.
//
// These text lines are returned in an array of
// ITextLineSpecification objects.
//
// If the text line collection maintained by the current
// TextLineSpecLinesCollection instance is empty (contains zero
// elements), an error will be returned.
//
// If any of the text lines within the collection maintained by
// the current TextLineSpecLinesCollection instance are invalid,
// an error will be returned.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings containing
//	                    error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	                        ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	[]ITextLineSpecification
//	   - If this method completes successfully, a deep copy of the
//	     text line collection maintained by the current
//	     TextLineSpecLinesCollection instance will be returned.
//	     These text lines are returned as an array of objects
//	     implementing the ITextLineSpecification interface.
//
//
//	error
//	   - If the method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (txtLinesSpecCol *TextLineSpecLinesCollection) GetTextLineCollection(
	errorPrefix interface{}) (
	[]ITextLineSpecification,
	error) {

	if txtLinesSpecCol.lock == nil {
		txtLinesSpecCol.lock = new(sync.Mutex)
	}

	txtLinesSpecCol.lock.Lock()

	defer txtLinesSpecCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecLinesCollection."+
			"GetTextLineCollection()",
		"")

	if err != nil {
		return nil, err
	}

	lenTxtLines := len(txtLinesSpecCol.textLines)

	if lenTxtLines == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: The text lines collection is empty!\n"+
			"TextLineSpecLinesCollection.textLines contains zero text line objects!\n",
			ePrefix.String())

		return nil, err
	}

	newTextLines :=
		make([]ITextLineSpecification, lenTxtLines)

	for i := 0; i < lenTxtLines; i++ {

		if txtLinesSpecCol.textLines[i] == nil {
			err = fmt.Errorf("%v\n"+
				"Error: Text Line element txtLinesSpecCol.textLines[%v]\n"+
				"has a 'nil' value!\n",
				ePrefix.String(),
				i)

			return nil, err
		}

		err = txtLinesSpecCol.textLines[i].IsValidInstanceError(
			ePrefix.XCpy(
				fmt.Sprintf(
					"txtLinesSpecCol.textLines[%v] invalid",
					i)))

		if err != nil {
			return nil, err
		}

		var newTextLine ITextLineSpecification

		newTextLine,
			err = txtLinesSpecCol.textLines[i].CopyOutITextLine(
			ePrefix.XCpy(
				fmt.Sprintf(
					"txtLinesSpecCol.textLines[%v] copy error",
					i)))

		if err != nil {
			return nil, err
		}

		newTextLines[i] = newTextLine
	}

	return newTextLines, err
}

// InsertTextLine - Receives a Text Line instance in the form of a
// type ITextLineSpecification. This Text Line object is then
// inserted into the Text Lines Collection maintained by the
// current instance of TextLineSpecLinesCollection.
//
// The Text Line input parameter, 'textLine', is inserted into
// the internal Text Lines collection at the array element index
// position indicated by input parameter, 'zeroBasedIndex'.
//
// After this method completes, the number of elements in the Text
// Lines Collection will be increased by one.
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//	textLine                   ITextLineSpecification
//	   - A Text Line object which implements the
//	     ITextLineSpecification interface. A deep copy of this
//	     text field will be inserted into the Text Lines Collection
//	     maintained by the current instance of
//	     TextLineSpecLinesCollection.
//
//	     After the insertion operation is completed, the
//	     'textLine' object will be located at array element
//	     'zeroBasedIndex' immediately BEFORE the original array
//	      element previously located at that array index.
//
//	     NOTE: You will need to pass the concrete instance of
//	     'textLine' as a pointer to the Text Line (&textLine).
//
//	     If the 'textLine' parameter is found to be invalid, an
//	     error will be returned.
//
//
//	zeroBasedIndex             int
//	   - This index number designates the array element index in
//	     the Text Lines Collection of the current
//	     TextLineSpecLinesCollection instance where the Text Line
//	     parameter, 'textLine' will be inserted. After insertion,
//	     the 'textLine' object will be positioned immediately
//	     BEFORE the original array element previously located at
//	     that array index.
//
//	     If 'zeroBasedIndex' is set to '4', the original Text Line
//	     object at index '4' will be moved to index position '5'
//	     after the insertion operation is completed.
//
//	     If the value of 'zeroBasedIndex' is less than zero, it
//	     will be reset to zero. This means that the 'textLine'
//	     object will be inserted in the first array element
//	     position of the Text Fields Collection maintained by the
//	     current TextLineSpecLinesCollection instance.
//
//	     If the value of 'zeroBasedIndex' is greater the last array
//	     element index in the Text Fields Collection, the
//	     'textLine' object will be appended to the end of the Text
//	     Lines Collection.
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings containing
//	                    error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	                        ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// Return Values
//
//	lastIndexId                int
//	   - If this method completes successfully, the internal array
//	     index of the last text line object for the Text Lines
//	     Collection maintained by the current
//	     TextLineSpecLinesCollection instance will be returned as
//	     an integer value. Remember, this is a zero based index
//	     value which is always one less than the length of the Text
//	     Line Collection.
//
//	     In the event of an error, 'lastIndexId' will be set to a
//	     value of minus one (-1).
//
//
//	err                        error
//	   - If the method completes successfully and no errors are
//	     encountered, this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (txtLinesSpecCol *TextLineSpecLinesCollection) InsertTextLine(
	textLine ITextLineSpecification,
	zeroBasedIndex int,
	errorPrefix interface{}) (
	lastIndexId int,
	err error) {

	if txtLinesSpecCol.lock == nil {
		txtLinesSpecCol.lock = new(sync.Mutex)
	}

	txtLinesSpecCol.lock.Lock()

	defer txtLinesSpecCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	lastIndexId = -1

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecLinesCollection."+
			"InsertTextLine()",
		"")

	if err != nil {
		return lastIndexId, err
	}

	lastIndexId,
		err = new(textLineSpecLinesCollectionAtom).
		insertTextLine(
			txtLinesSpecCol,
			textLine,
			zeroBasedIndex,
			ePrefix.XCpy(
				fmt.Sprintf(
					"txtLinesSpecCol[%v]<-textLine",
					zeroBasedIndex)))

	return lastIndexId, err
}

// IsValidInstance - Performs a diagnostic review of the data
// values encapsulated in the current TextLineSpecLinesCollection
// instance to determine if they are valid.
//
// If all data element evaluate as valid, this method returns
// 'true'. If any data element is invalid, this method returns
// 'false'.
//
// ------------------------------------------------------------------------
//
// # BE ADVISED
//
// If the current instance of TextLineSpecLinesCollection contains
// zero Text Line members in the collection, this method will
// return 'false'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	--- NONE ---
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	isValid             bool
//	   - If all data elements encapsulated by the current instance
//	     of TextLineSpecLinesCollection are valid, this returned
//	     boolean value is set to 'true'. If any data values are
//	     invalid, this return parameter is set to 'false'.
func (txtLinesSpecCol *TextLineSpecLinesCollection) IsValidInstance() (
	isValid bool) {

	if txtLinesSpecCol.lock == nil {
		txtLinesSpecCol.lock = new(sync.Mutex)
	}

	txtLinesSpecCol.lock.Lock()

	defer txtLinesSpecCol.lock.Unlock()

	isValid,
		_ = new(textLineSpecLinesCollectionAtom).
		testValidityOfTextLinesCollection(
			txtLinesSpecCol,
			nil)

	return isValid
}

// IsValidInstanceError - Performs a diagnostic review of the data
// values encapsulated in the current TextLineSpecLinesCollection
// instance to determine if they are valid.
//
// If any data element evaluates as invalid, this method will
// return an error.
//
// ------------------------------------------------------------------------
//
// # BE ADVISED
//
// If the current instance of TextLineSpecLinesCollection contains
// zero Text Line members in the collection, this method will
// return an error.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	errorPrefix         interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings containing
//	                    error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	                        ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If any of the internal member data variables contained in
//	     the current instance of TextLineSpecLinesCollection are found
//	     to be invalid, this method will return an error.
//
//	     Also, if the current instance of
//	     TextLineSpecLinesCollection contains zero Text Line members
//	     in the collection, this method will return an error.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' (error prefix) will be inserted or
//	     prefixed at the beginning of the error message.
func (txtLinesSpecCol *TextLineSpecLinesCollection) IsValidInstanceError(
	errorPrefix interface{}) error {

	if txtLinesSpecCol.lock == nil {
		txtLinesSpecCol.lock = new(sync.Mutex)
	}

	txtLinesSpecCol.lock.Lock()

	defer txtLinesSpecCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecLinesCollection."+
			"IsValidInstanceError()",
		"")

	if err != nil {
		return err
	}

	_,
		err = new(textLineSpecLinesCollectionAtom).
		testValidityOfTextLinesCollection(
			txtLinesSpecCol,
			ePrefix.XCpy(
				"txtLinesSpecCol"))

	return err
}

// New - Returns a new, empty instance of
// TextLineSpecLinesCollection.
//
// The Text Line Collection for this returned instance is empty and
// contains zero member elements.
//
// To add Text Lines to the collection encapsulated by this
// instance of TextLineSpecLinesCollection, call the method
//
//	TextLineSpecLinesCollection.AddTextLineSpec()
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	--- NONE ---
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	TextLineSpecLinesCollection
//	   - This method will return an empty or uninitialized instance
//	     of TextLineSpecLinesCollection. The Text Line Collection
//	     encapsulated by this instance contains zero member
//	     elements.
func (txtLinesSpecCol TextLineSpecLinesCollection) New() TextLineSpecLinesCollection {

	if txtLinesSpecCol.lock == nil {
		txtLinesSpecCol.lock = new(sync.Mutex)
	}

	txtLinesSpecCol.lock.Lock()

	defer txtLinesSpecCol.lock.Unlock()

	newTxtLineSpec := TextLineSpecLinesCollection{}

	return newTxtLineSpec
}

// NewTextLine - Returns a new instance of
// TextLineSpecLinesCollection. The Text Lines Collection of this
// new TextLineSpecLinesCollection instance will be populated with
// a single Text Line object passed as input parameter, 'textLine'
//
// To add more Text Lines to the Text Lines Collection maintained
// by the returned instance of TextLineSpecLinesCollection, call
// the method
//
//	TextLineSpecLinesCollection.AddTextLineSpec()
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	textLine                   ITextLineSpecification
//	   - A text line object which implements the
//	     ITextLineSpecification interface. A deep copy of this
//	     object will be added to the text lines collection
//	     maintained by the returned instance of
//	     TextLineSpecLinesCollection.
//
//	     If member variable data values contained in this
//	     'textLine' parameter are found to be invalid, an error
//	     will be returned.
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings containing
//	                    error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	                        ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	newTextLineCol             TextLineSpecLinesCollection
//	   - This method will return a new instance of
//	     TextLineSpecLinesCollection. The Text Line Collection
//	     encapsulated by this returned instance will contain a
//	     single Text Line object passed as input parameter,
//	     'textLine'.
//
//
//	err                        error
//	   - If the method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (txtLinesSpecCol TextLineSpecLinesCollection) NewTextLine(
	textLine ITextLineSpecification,
	errorPrefix interface{}) (
	newTextLineCol TextLineSpecLinesCollection,
	err error) {

	if txtLinesSpecCol.lock == nil {
		txtLinesSpecCol.lock = new(sync.Mutex)
	}

	txtLinesSpecCol.lock.Lock()

	defer txtLinesSpecCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	newTextLineCol = TextLineSpecLinesCollection{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecLinesCollection.NewTextLine()",
		"")

	if err != nil {
		return newTextLineCol, err
	}

	err = new(textLineSpecLinesCollectionNanobot).
		addTextLine(
			&newTextLineCol,
			textLine,
			ePrefix.XCpy(
				"newTextLineCol<-textLine"))

	return newTextLineCol, err
}

// NewPtr - Returns a pointer to a new, empty instance of
// TextLineSpecLinesCollection.
//
// The Text Line Collection for this returned instance is empty and
// contains zero member elements.
//
// To add Text Lines to the collection encapsulated by this
// instance of TextLineSpecLinesCollection, call the method
//
//	TextLineSpecLinesCollection.AddTextLineSpec()
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	--- NONE ---
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	*TextLineSpecLinesCollection
//	   - This method will return a pointer to an empty or
//	     uninitialized instance of TextLineSpecLinesCollection. The
//	     Text Line Collection encapsulated by this instance
//	     contains zero member elements.
func (txtLinesSpecCol TextLineSpecLinesCollection) NewPtr() *TextLineSpecLinesCollection {

	if txtLinesSpecCol.lock == nil {
		txtLinesSpecCol.lock = new(sync.Mutex)
	}

	txtLinesSpecCol.lock.Lock()

	defer txtLinesSpecCol.lock.Unlock()

	newTxtLineSpec := TextLineSpecLinesCollection{}

	return &newTxtLineSpec

}

// NewPtrTextLine - Returns a pointer to a new instance of
// TextLineSpecLinesCollection. The Text Lines Collection of this
// new TextLineSpecLinesCollection instance will be populated with
// a single Text Line object passed as input parameter, 'textLine'
//
// To add more Text Lines to the Text Lines Collection maintained
// by the returned instance of TextLineSpecLinesCollection, call
// the method
//
//	TextLineSpecLinesCollection.AddTextLineSpec()
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	textLine                   ITextLineSpecification
//	   - A text line object which implements the
//	     ITextLineSpecification interface. A deep copy of this
//	     object will be added to the text lines collection
//	     maintained by the returned instance of
//	     TextLineSpecLinesCollection.
//
//	     If member variable data values contained in this
//	     'textLine' parameter are found to be invalid, an error
//	     will be returned.
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings containing
//	                    error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	                        ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	newTextLineCol             *TextLineSpecLinesCollection
//	   - This method will return a pointer to a new instance of
//	     TextLineSpecLinesCollection. The Text Line Collection
//	     encapsulated by this returned instance will contain a
//	     single Text Line object passed as input parameter,
//	     'textLine'.
//
//
//	err                        error
//	   - If the method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (txtLinesSpecCol TextLineSpecLinesCollection) NewPtrTextLine(
	textLine ITextLineSpecification,
	errorPrefix interface{}) (
	newTextLineCol *TextLineSpecLinesCollection,
	err error) {

	if txtLinesSpecCol.lock == nil {
		txtLinesSpecCol.lock = new(sync.Mutex)
	}

	txtLinesSpecCol.lock.Lock()

	defer txtLinesSpecCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	newTextLineCol = &TextLineSpecLinesCollection{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecLinesCollection.NewPtrTextLine()",
		"")

	if err != nil {
		return newTextLineCol, err
	}

	err = new(textLineSpecLinesCollectionNanobot).
		addTextLine(
			newTextLineCol,
			textLine,
			ePrefix.XCpy(
				"newTextLineCol<-textLine"))

	return newTextLineCol, err
}

// PeekAtFirstTextLine - Returns a deep copy of the first Text Line
// ('ITextLineSpecification') object in the Text Lines Collection
// ('txtLinesCol.textLines[0]').
//
// As a 'Peek' method, the original Text Line object
// ('txtLinesCol.textLines[0]') WILL NOT be deleted from the Text
// Lines Collection encapsulated by the current instance of
// TextLineSpecLinesCollection.
//
// After completion of this method, the Text Line Collection array
// will remain unchanged.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of
//	                        ErrPrefixDto. ErrorPrefixInfo from this
//	                        object will be copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	iTxtLineSpec              ITextLineSpecification
//	   - If this method completes successfully, a deep copy of
//	     the designated member of the Text Lines Collection
//	     will be returned to the calling function. The returned
//	     object will implement the ITextLineSpecification
//	     interface.
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
//
// ----------------------------------------------------------------
//
// Example Usage
//
//	When casting ITextLineSpecification returned from this method,
//	use the following syntax to cast the interface object to a
//	concrete type.
//
//	It is necessary to cast the interface object ('iTxtLineSpec')
//	as a pointer to the concrete type ('stdLine'). This is because
//	the concrete type uses methods with pointer receivers.
//
//	---------------------------------------------------------------
//	   var iTxtLineSpec ITextLineSpecification
//
//	   iTxtLineSpec,
//	   err = txtLinesCol01.PeekAtFirstTextLine(
//	           ePrefix.XCpy(
//	           "txtLinesCol01"))
//
//	   if err != nil {
//	     return err
//	   }
//
//	   var stdLine *TextLineSpecStandardLine
//
//	   var ok bool
//
//	   stdLine, ok = iTxtLineSpec.(*TextLineSpecStandardLine)
//
//	   if !ok {
//
//	     err = fmt.Errorf("%v - Error\n"+
//	     "stdLine, ok := iTxtLineSpec.(*TextLineSpecStandardLine)\n"+
//	     "Expected return of type 'TextLineSpecStandardLine'.\n"+
//	     "HOWEVER, THAT TYPE WAS NOT RETURNED!\n",
//	     ePrefix.String())
//
//	     return err
//	   }
//
//	   // 'stdLine' is now available for use
//	   // as a concrete object.
//	   stdLineLen := stdLine.GetSingleLineLength()
func (txtLinesSpecCol *TextLineSpecLinesCollection) PeekAtFirstTextLine(
	errorPrefix interface{}) (
	ITextLineSpecification,
	error) {

	if txtLinesSpecCol.lock == nil {
		txtLinesSpecCol.lock = new(sync.Mutex)
	}

	txtLinesSpecCol.lock.Lock()

	defer txtLinesSpecCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	var iTextLineSpec ITextLineSpecification

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecLinesCollection."+
			"PeekAtFirstTextLine()",
		"")

	if err != nil {
		return iTextLineSpec, err
	}

	iTextLineSpec,
		err = new(textLineSpecLinesCollectionAtom).
		peekPopTextLine(
			txtLinesSpecCol,
			0,
			false,
			ePrefix.XCpy(
				"txtLinesSpecCol[0]"))

	return iTextLineSpec, err
}

// PeekAtLastTextLine - Returns a deep copy of the last Text Line
// ('ITextLineSpecification') object in the Text Lines Collection
// ('txtLinesCol.textLines[lastIdx]').
//
// As a 'Peek' method, the original Text Line object
// ('txtLinesCol.textLines[lastIdx]') WILL NOT be deleted from the Text
// Line Collection encapsulated by this instance of
// TextLineSpecStandardLine.
//
// After completion of this method, the Text Line Collection array
// will remain unchanged.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings containing
//	                    error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	                        ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	iTxtLineSpec              ITextLineSpecification
//	   - If this method completes successfully, a deep copy of
//	     the last member element of the Text Lines Collection
//	     will be returned to the calling function. The returned
//	     object will implement the ITextLineSpecification
//	     interface.
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
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//	When casting ITextLineSpecification returned from this method,
//	use the following syntax to cast the interface object to a
//	concrete type.
//
//	It is necessary to cast the interface object ('iTxtLineSpec')
//	as a pointer to the concrete type ('stdLine'). This is because
//	the concrete type uses methods with pointer receivers.
//
// /
//
//	------------------------------------------------------------
//	   var iTxtLineSpec ITextLineSpecification
//
//	   iTxtLineSpec,
//	   err = txtLinesCol01.PeekAtLastTextLine(
//	         ePrefix.XCpy(
//	         "txtLinesCol01"))
//
//	   if err != nil {
//	     return err
//	   }
//
//	   var stdLine *TextLineSpecStandardLine
//
//	   var ok bool
//
//	   stdLine, ok = iTxtLineSpec.(*TextLineSpecStandardLine)
//
//	   if !ok {
//
//	     err = fmt.Errorf("%v - Error\n"+
//	     "stdLine, ok := iTxtLineSpec.(*TextLineSpecStandardLine)\n"+
//	     "Expected return of type 'TextLineSpecStandardLine'.\n"+
//	     "HOWEVER, THAT TYPE WAS NOT RETURNED!\n",
//	     ePrefix.String())
//
//	     return err
//	   }
//
//	   // 'stdLine' is now available for use
//	   // as a concrete object.
//	   stdLineLen := stdLine.GetSingleLineLength()
func (txtLinesSpecCol *TextLineSpecLinesCollection) PeekAtLastTextLine(
	errorPrefix interface{}) (
	ITextLineSpecification,
	error) {

	if txtLinesSpecCol.lock == nil {
		txtLinesSpecCol.lock = new(sync.Mutex)
	}

	txtLinesSpecCol.lock.Lock()

	defer txtLinesSpecCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	var iTextLineSpec ITextLineSpecification

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecLinesCollection."+
			"PeekAtLastTextLine()",
		"")

	if err != nil {
		return iTextLineSpec, err
	}

	lastIdx := len(txtLinesSpecCol.textLines) - 1

	if lastIdx < 0 {
		err = fmt.Errorf("%v - ERROR\n"+
			"The Text Lines Collection is empty!\n",
			ePrefix.String())

		return iTextLineSpec, err
	}

	iTextLineSpec,
		err = new(textLineSpecLinesCollectionAtom).
		peekPopTextLine(
			txtLinesSpecCol,
			lastIdx,
			false,
			ePrefix.XCpy(
				fmt.Sprintf("txtLinesSpecCol[%v]",
					lastIdx)))

	return iTextLineSpec, err
}

// PeekAtTextLine - Returns a deep copy of the Text Line Collection
// member element specified by input parameter, 'zeroBasedIndex'.
//
// If the Text Line collection maintained by the current
// TextLineSpecLinesCollection instance is empty (contains zero
// elements), an error will be returned.
//
// Remember that indexes in the Text Lines collection are zero
// based. This means the first element in the collection is index
// zero.
//
// If input parameter 'zeroBasedIndex' is less than zero or greater
// than the index of the last member element in the collection, an
// error will be returned.
//
// -----------------------------------------------------------------
//
// # BE ADVISED
//
// This method ( PeekAtTextField() ) is functionally equivalent to
// method:
//
//	TextLineSpecLinesCollection.GetTextField()
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	zeroBasedIndex             int
//	   - Specifies the index of the member element in the Text
//	     Lines collection which will be returned as a deep copy
//	     of the original. If this input parameter is found to be
//	     invalid or if the Text Lines collection is empty, an
//	     error will be returned.
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	                    containing error prefix and error context
//	                    information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	                        ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	ITextLineSpecification
//	   - If this method completes successfully, a deep copy of the
//	     Text Line member element specified by input parameter
//	     'zeroBasedIndex' will be returned.
//
//
//	error
//	   - If the method completes successfully and no errors are
//	     encountered, this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//	When casting ITextLineSpecification returned from this method,
//	use the following syntax to cast the interface object to a
//	concrete type.
//
//	It is necessary to cast the interface object ('iTxtLineSpec')
//	as a pointer to the concrete type ('stdLine'). This is because
//	the concrete type uses methods with pointer receivers.
//
//	------------------------------------------------------------
//	   var iTxtLineSpec ITextLineSpecification
//
//	   iTxtLineSpec,
//	   err = txtLinesCol01.PeekAtTextLine(
//	           2,  // Return Text Line at index '2'
//	           ePrefix.XCpy(
//	           "txtLinesCol01"))
//
//	   if err != nil {
//	     return err
//	   }
//
//	   var stdLine *TextLineSpecStandardLine
//
//	   var ok bool
//
//	   stdLine, ok = iTxtLineSpec.(*TextLineSpecStandardLine)
//
//	   if !ok {
//
//	     err = fmt.Errorf("%v - Error\n"+
//	     "stdLine, ok := iTxtLineSpec.(*TextLineSpecStandardLine)\n"+
//	     "Expected return of type 'TextLineSpecStandardLine'.\n"+
//	     "HOWEVER, THAT TYPE WAS NOT RETURNED!\n",
//	     ePrefix.String())
//
//	     return err
//	   }
//
//	   // 'stdLine' is now available for use
//	   // as a concrete object.
//	   stdLineLen := stdLine.GetSingleLineLength()
func (txtLinesSpecCol *TextLineSpecLinesCollection) PeekAtTextLine(
	zeroBasedIndex int,
	errorPrefix interface{}) (
	ITextLineSpecification,
	error) {

	if txtLinesSpecCol.lock == nil {
		txtLinesSpecCol.lock = new(sync.Mutex)
	}

	txtLinesSpecCol.lock.Lock()

	defer txtLinesSpecCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	var iTextLineSpec ITextLineSpecification

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecLinesCollection."+
			"PeekAtTextLine()",
		"")

	if err != nil {
		return iTextLineSpec, err
	}

	iTextLineSpec,
		err = new(textLineSpecLinesCollectionAtom).
		peekPopTextLine(
			txtLinesSpecCol,
			zeroBasedIndex,
			false,
			ePrefix.XCpy(
				fmt.Sprintf(
					"txtLinesSpecCol[%v]",
					zeroBasedIndex)))

	return iTextLineSpec, err
}

// PopFirstTextLine - Returns a deep copy of the first Text Line
// ('ITextLineSpecification') object in the Text Line Collection
// ('txtLinesCol.textLines[0]').
//
// As a 'Pop' method, the first Text Line object will be deleted
// from the Text Line Collection encapsulated by this instance of
// TextLineSpecStandardLine. Parameter 'remainingNumOfTextLines'
// will be returned to the calling function containing the number
// of array elements still remaining in the Text Line Collection
// after deletion of the first array element.
//
// The number of array elements remaining in the Text Line
// Collection after this operation will always be one less than
// the original number of array elements.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// After successful completion of this method, the first member of
// the Text Line Collection will be DELETED and the Text Line
// Collection array will have a length which is one less than the
// original array length.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings containing
//	                    error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	                        ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	iTxtLineSpec               ITextLineSpecification
//	   - If this method completes successfully, a deep copy of the
//	     first member of the Text Lines Collection will be returned
//	     to the calling function. The returned object will
//	     implement the ITextLineSpecification interface.
//
//	     After completion, the first element of Text Lines
//	     Collection will be deleted.
//
//
//	remainingNumOfTextLines    int
//	   - If this method completes successfully, the first array
//	     element in the Text Lines Collection will be deleted.
//	     After deleting that element, this parameter will return
//	     the number of array elements still remaining in the
//	     Text Lines Collection.
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
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//	When casting ITextLineSpecification returned from this method,
//	use the following syntax to cast the interface object to a
//	concrete type.
//
//	It is necessary to cast the interface object ('iTxtLineSpec')
//	as a pointer to the concrete type ('stdLine'). This is because
//	the concrete type uses methods with pointer receivers.
//
//	------------------------------------------------------------
//	   var iTxtLineSpec ITextLineSpecification
//
//	   iTxtLineSpec,
//	   err = txtLinesCol01.PopFirstTextLine(
//	           ePrefix.XCpy(
//	           "txtLinesCol01"))
//
//	   if err != nil {
//	     return err
//	   }
//
//	   // BE ADVISED
//	   // This 'Pop' METHOD WILL DELETE THE FIRST
//	   // MEMBER OF THE TEXT LINES COLLECTION!!!
//
//	   var stdLine *TextLineSpecStandardLine
//
//	   var ok bool
//
//	   stdLine, ok = iTxtLineSpec.(*TextLineSpecStandardLine)
//
//	   if !ok {
//
//	     err = fmt.Errorf("%v - Error\n"+
//	     "stdLine, ok := iTxtLineSpec.(*TextLineSpecStandardLine)\n"+
//	     "Expected return of type 'TextLineSpecStandardLine'.\n"+
//	     "HOWEVER, THAT TYPE WAS NOT RETURNED!\n",
//	     ePrefix.String())
//
//	     return err
//	   }
//
//	   // 'stdLine' is now available for use
//	   // as a concrete object.
//	   stdLineLen := stdLine.GetSingleLineLength()
func (txtLinesSpecCol *TextLineSpecLinesCollection) PopFirstTextLine(
	errorPrefix interface{}) (
	iTxtLineSpec ITextLineSpecification,
	remainingNumOfTextLines int,
	err error) {

	if txtLinesSpecCol.lock == nil {
		txtLinesSpecCol.lock = new(sync.Mutex)
	}

	txtLinesSpecCol.lock.Lock()

	defer txtLinesSpecCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var iTextLineSpec ITextLineSpecification

	remainingNumOfTextLines = -1

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecLinesCollection."+
			"PopFirstTextLine()",
		"")

	if err != nil {
		return iTextLineSpec, remainingNumOfTextLines, err
	}

	iTextLineSpec,
		err = new(textLineSpecLinesCollectionAtom).
		peekPopTextLine(
			txtLinesSpecCol,
			0,
			true,
			ePrefix.XCpy(
				"txtLinesSpecCol[0]"))

	remainingNumOfTextLines = len(txtLinesSpecCol.textLines)

	return iTextLineSpec, remainingNumOfTextLines, err
}

// PopLastTextLine - Returns a deep copy of the last Text Line
// ('ITextLineSpecification') element in the Text Line Collection
// ('txtLinesCol.textLines[lastIndex]').
//
// As a 'Pop' method, the last Text Line object will be deleted
// from the Text Line Collection encapsulated by this instance of
// TextLineSpecStandardLine. Parameter 'remainingNumOfTextLines'
// will be returned to the calling function containing the number
// of array elements still remaining in the Text Line Collection
// after deletion of the last array element.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// After successful completion of this method, the last member of
// the Text Line Collection will be DELETED and the Text Line
// Collection array will have a length which is one less than the
// original array length.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings containing
//	                    error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	                        ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	iTxtLineSpec               ITextLineSpecification
//	   - If this method completes successfully, a deep copy of
//	     if the last member of the Text Lines Collection will be
//	     returned to the calling function. The returned object
//	     will implement the ITextLineSpecification interface.
//
//	     After completion, the last element of Text Lines
//	     Collection will be deleted.
//
//
//	remainingNumOfTextLines    int
//	   - If this method completes successfully, the last array
//	     element in the Text Lines Collection will be deleted.
//	     After deleting that element, this parameter will return
//	     the number of array elements still remaining in the
//	     Text Lines Collection.
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
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//	When casting ITextLineSpecification returned from this method,
//	use the following syntax to cast the interface object to a
//	concrete type.
//
//	It is necessary to cast the interface object ('iTxtLineSpec')
//	as a pointer to the concrete type ('stdLine'). This is because
//	the concrete type uses methods with pointer receivers.
//
//	------------------------------------------------------------
//	   var iTxtLineSpec ITextLineSpecification
//
//	   iTxtLineSpec,
//	   err = txtLinesCol01.PopLastTextLine(
//	           ePrefix.XCpy(
//	           "txtLinesCol01"))
//
//	   if err != nil {
//	     return err
//	   }
//
//	   // BE ADVISED
//	   // This 'Pop' METHOD WILL DELETE THE LAST
//	   // MEMBER OF THE TEXT LINES COLLECTION!!!
//
//	   var stdLine *TextLineSpecStandardLine
//
//	   var ok bool
//
//	   stdLine, ok = iTxtLineSpec.(*TextLineSpecStandardLine)
//
//	   if !ok {
//
//	     err = fmt.Errorf("%v - Error\n"+
//	     "stdLine, ok := iTxtLineSpec.(*TextLineSpecStandardLine)\n"+
//	     "Expected return of type 'TextLineSpecStandardLine'.\n"+
//	     "HOWEVER, THAT TYPE WAS NOT RETURNED!\n",
//	     ePrefix.String())
//
//	     return err
//	   }
//
//	   // 'stdLine' is now available for use
//	   // as a concrete object.
//	   stdLineLen := stdLine.GetSingleLineLength()
func (txtLinesSpecCol *TextLineSpecLinesCollection) PopLastTextLine(
	errorPrefix interface{}) (
	iTxtLineSpec ITextLineSpecification,
	remainingNumOfTextLines int,
	err error) {

	if txtLinesSpecCol.lock == nil {
		txtLinesSpecCol.lock = new(sync.Mutex)
	}

	txtLinesSpecCol.lock.Lock()

	defer txtLinesSpecCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var iTextLineSpec ITextLineSpecification

	remainingNumOfTextLines = -1

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecLinesCollection."+
			"PopLastTextLine()",
		"")

	if err != nil {
		return iTextLineSpec, remainingNumOfTextLines, err
	}

	lastIdx := len(txtLinesSpecCol.textLines) - 1

	if lastIdx < 0 {
		err = fmt.Errorf("%v - ERROR\n"+
			"The Text Lines Collection is empty!\n",
			ePrefix.String())

		return iTextLineSpec, remainingNumOfTextLines, err
	}

	iTextLineSpec,
		err = new(textLineSpecLinesCollectionAtom).
		peekPopTextLine(
			txtLinesSpecCol,
			lastIdx,
			true,
			ePrefix.XCpy(
				"txtLinesSpecCol[0]"))

	remainingNumOfTextLines = len(txtLinesSpecCol.textLines)

	return iTextLineSpec, remainingNumOfTextLines, err
}

// PopTextLine - Returns a deep copy of the Text Line Collection
// member element specified by input parameter, 'zeroBasedIndex'.
//
// If the Text Line collection maintained by the current
// TextLineSpecLinesCollection instance is empty (contains zero
// elements), an error will be returned.
//
// Remember that indexes in the Text Lines collection are zero
// based. This means the first element in the collection is always
// index zero.
//
// If input parameter 'zeroBasedIndex' is less than zero or greater
// than the last member element in the collection, an error will be
// returned.
//
// -----------------------------------------------------------------
//
// # BE ADVISED
//
// This method ( PopTextField() ) is similar to the following
// methods:
//
//	TextLineSpecLinesCollection.GetTextField()
//	TextLineSpecLinesCollection.PeekAtTextField()
//
// The sole difference between this method and the two methods
// cited above is that this method deletes the target collection
// member element.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// After successful completion of this method, the target member of
// the Text Line Collection specified by input parameter
// 'zeroBasedIndex' will be DELETED.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	zeroBasedIndex             int
//	   - Specifies the index of the member element in the Text
//	     Lines collection which will be returned as a deep copy
//	     of the original. After returning a deep copy of the
//	     collection element specified by 'zeroBasedIndex',
//	     that element will be deleted from the collection.
//
//	     If this input parameter is found to be invalid or if the
//	     Text Lines collection is empty, an error will be returned.
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	                    containing error prefix and error context
//	                    information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	                        ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	ITextLineSpecification
//	   - If this method completes successfully, a deep copy of the
//	     Text Line member element specified by input parameter
//	     'zeroBasedIndex' will be returned. In addition, the
//	     original collection element specified by 'zeroBasedIndex'
//	     will be deleted.
//
//
//	error
//	   - If the method completes successfully and no errors are
//	     encountered, this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//	When casting ITextLineSpecification returned from this method,
//	use the following syntax to cast the interface object to a
//	concrete type.
//
//	It is necessary to cast the interface object ('iTxtLineSpec')
//	as a pointer to the concrete type ('stdLine'). This is because
//	the concrete type uses methods with pointer receivers.
//
//	------------------------------------------------------------
//	   var iTxtLineSpec ITextLineSpecification
//
//	   iTxtLineSpec,
//	   err = txtLinesCol01.PopTextLine(
//	           2,  // Return a copy of and delete Text Line index '2'
//	           ePrefix.XCpy(
//	           "txtLinesCol01"))
//
//	   if err != nil {
//	     return err
//	   }
//
//	   // BE ADVISED
//	   // This 'Pop' METHOD WILL DELETE THE TARGET
//	   // MEMBER OF THE TEXT LINES COLLECTION DESIGNATED
//	   // BY 'zeroBasedIndex'
//
//	   var stdLine *TextLineSpecStandardLine
//
//	   var ok bool
//
//	   stdLine, ok = iTxtLineSpec.(*TextLineSpecStandardLine)
//
//	   if !ok {
//
//	     err = fmt.Errorf("%v - Error\n"+
//	     "stdLine, ok := iTxtLineSpec.(*TextLineSpecStandardLine)\n"+
//	     "Expected return of type 'TextLineSpecStandardLine'.\n"+
//	     "HOWEVER, THAT TYPE WAS NOT RETURNED!\n",
//	     ePrefix.String())
//
//	     return err
//	   }
//
//	   // 'stdLine' is now available for use
//	   // as a concrete object.
//	   stdLineLen := stdLine.GetSingleLineLength()
func (txtLinesSpecCol *TextLineSpecLinesCollection) PopTextLine(
	zeroBasedIndex int,
	errorPrefix interface{}) (
	iTxtLineSpec ITextLineSpecification,
	remainingNumOfTextLines int,
	err error) {

	if txtLinesSpecCol.lock == nil {
		txtLinesSpecCol.lock = new(sync.Mutex)
	}

	txtLinesSpecCol.lock.Lock()

	defer txtLinesSpecCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var iTextLineSpec ITextLineSpecification

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecLinesCollection."+
			"PopTextLine()",
		"")

	if err != nil {
		return iTextLineSpec, remainingNumOfTextLines, err
	}

	iTextLineSpec,
		err = new(textLineSpecLinesCollectionAtom).
		peekPopTextLine(
			txtLinesSpecCol,
			zeroBasedIndex,
			true,
			ePrefix.XCpy(
				fmt.Sprintf(
					"txtLinesSpecCol[%v]",
					zeroBasedIndex)))

	remainingNumOfTextLines = len(txtLinesSpecCol.textLines)

	return iTextLineSpec, remainingNumOfTextLines, err
}

// ReplaceTextLine - Receives an object which implements the
// ITextLineSpecification interface. This object will replace an
// existing text line object within the text line collection
// maintained by this TextLineSpecLinesCollection instance.
//
// The text line object to be replaced must exist at the index
// specified by input parameter, 'replaceAtIndex'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	textLine                   ITextLineSpecification
//	   - A text line object which implements the
//	     ITextLineSpecification interface. A deep copy of this
//	     object will replace an existing element within the
//	     text lines collection maintained by this instance of
//	     TextLineSpecLinesCollection. The text line object to
//	     be replaced is identified by the collection element index
//	     supplied by input parameter 'replaceAtIndex'.
//
//	     If member variable data values contained in this
//	     'textLine' parameter are found to be invalid, an error
//	     will be returned.
//
//	     NOTE: You will need to pass the concrete instance of
//	     'textLine' as a pointer to the Text Line (&textLine).
//
//
//	replaceAtIndex             int
//	   - The index of an element within the text lines collection
//	     maintained by the current TextLineSpecLinesCollection
//	     instance which will be replaced by input parameter
//	     'textLine'.
//
//	     Remember that the text fields collection maintained by
//	     the current TextLineSpecLinesCollection instance is a zero
//	     based array. Therefore, the first index in the collection
//	     is zero (0).
//
//	     If 'replaceAtIndex' proves to be an invalid index, an error
//	     will be returned.
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of
//	                        ErrPrefixDto. ErrorPrefixInfo from this
//	                        object will be copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	err                        error
//	   - If the method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (txtLinesSpecCol *TextLineSpecLinesCollection) ReplaceTextLine(
	textLine ITextLineSpecification,
	replaceAtIndex int,
	errorPrefix interface{}) (
	err error) {

	if txtLinesSpecCol.lock == nil {
		txtLinesSpecCol.lock = new(sync.Mutex)
	}

	txtLinesSpecCol.lock.Lock()

	defer txtLinesSpecCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecLinesCollection."+
			"ReplaceTextLine()",
		"")

	if err != nil {
		return err
	}

	lenOfTextLinesCol := len(txtLinesSpecCol.textLines)

	if lenOfTextLinesCol == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: The text lines collection is empty and contains no text lines!\n"+
			"First add some text lines before trying to replace a text line.\n",
			ePrefix.String())

		return err
	}

	if textLine == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'textLine' is 'nil' and invalid!\n",
			ePrefix.String())

		return err

	}

	err = textLine.IsValidInstanceError(
		ePrefix.XCpy("Input Parameter: textLine"))

	if err != nil {
		return err
	}

	if replaceAtIndex < 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'replaceAtIndex' is invalid!\n"+
			"'replaceAtIndex' is less than zero (0).\n"+
			"replaceAtIndex = '%v'\n",
			ePrefix.String(),
			replaceAtIndex)

		return err
	}

	lenOfTextLinesCol--

	if replaceAtIndex > lenOfTextLinesCol {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'replaceAtIndex' is out of range and invalid!\n"+
			"'replaceAtIndex' is greater than the maximum collection index.\n"+
			"The last element in the text lines collection is index '%v'.\n"+
			"Input parameter 'replaceAtIndex' = '%v'\n",
			ePrefix.String(),
			lenOfTextLinesCol,
			replaceAtIndex)

		return err
	}

	var newTextLine ITextLineSpecification

	newTextLine,
		err = textLine.CopyOutITextLine(
		ePrefix.XCpy(
			"newTextLine"))

	if err != nil {
		return err
	}

	if txtLinesSpecCol.textLines[replaceAtIndex] != nil {

		txtLinesSpecCol.textLines[replaceAtIndex].Empty()

		txtLinesSpecCol.textLines[replaceAtIndex] = nil

	}

	txtLinesSpecCol.textLines[replaceAtIndex] = newTextLine

	return err
}

// SetTextLineCollection - Deletes the Text Line Collection for the
// current instance of TextLineSpecLinesCollection. This collection
// is then replaced with the new Text Line Collection passed as
// input parameter, 'newTextLineCol'.
//
// Input parameter 'newTextLineCol' is an array of
// ITextLineSpecification objects.
//
// Only deep copies of the member elements of 'newTextLineCol' will
// be transferred to the new Text Line Collection for the current
// instance of TextLineSpecLinesCollection.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will DELETE the current Text Line Collection and
// replace it with a new collection passed as an input parameter.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	newTextLineCol             []ITextLineSpecification
//	   - An array of ITextLineSpecification objects which will
//	     replace the Text Line Collection for the current instance
//	     of TextLineSpecLinesCollection.
//
//	     The replacement operation will first create deep copies
//	     of 'newTextLineCol' member elements before adding those
//	     deep copies to the TextLineSpecLinesCollection Text Line
//	     Collection.
//
//	     If any of the member elements of this array are found to
//	     be invalid, or if this parameter is passed as a zero
//	     length array, an error will be returned.
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings containing
//	                    error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	                        ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If the method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (txtLinesSpecCol *TextLineSpecLinesCollection) SetTextLineCollection(
	newTextLineCol []ITextLineSpecification,
	errorPrefix interface{}) error {

	if txtLinesSpecCol.lock == nil {
		txtLinesSpecCol.lock = new(sync.Mutex)
	}

	txtLinesSpecCol.lock.Lock()

	defer txtLinesSpecCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecLinesCollection."+
			"SetTextLineCollection()",
		"")

	if err != nil {
		return err
	}

	lenOfNewTxtLines := len(newTextLineCol)

	if lenOfNewTxtLines == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'newTextLineCol' is invalid!\n"+
			"'newTextLineCol' contains zero text line objects!\n",
			ePrefix.String())

		return err
	}

	var err2 error

	for i := 0; i < lenOfNewTxtLines; i++ {

		if newTextLineCol[i] == nil {
			err = fmt.Errorf("%v\n"+
				"Error: Input parameter element newTextLineCol[%v] is invalid!\n"+
				"newTextLineCol[%v] has a 'nil' value.\n",
				ePrefix.String(),
				i,
				i)

			return err
		}

		err2 = newTextLineCol[i].IsValidInstanceError(
			nil)

		if err2 != nil {
			err = fmt.Errorf("%v\n"+
				"Error: Input parameter element newTextLineCol[%v] is invalid!\n"+
				"newTextLineCol[%v] produced the following validation error:\n"+
				"%v\n",
				ePrefix.String(),
				i,
				i,
				err2.Error())

			return err
		}
	}

	textLineColAtom := textLineSpecLinesCollectionAtom{}

	textLineColAtom.emptyCollection(
		txtLinesSpecCol)

	txtLinesSpecCol.textLines = make([]ITextLineSpecification, lenOfNewTxtLines)

	for j := 0; j < lenOfNewTxtLines; j++ {

		txtLinesSpecCol.textLines[j],
			err2 = newTextLineCol[j].CopyOutITextLine(
			nil)

		if err2 != nil {

			textLineColAtom.emptyCollection(
				txtLinesSpecCol)

			err = fmt.Errorf("%v\n"+
				"Error: newTextLineCol[%v] Deep Copy Failed!\n"+
				"The Text Line Collection for the Current Instance\n"+
				"of TextLineSpecLinesCollection is now deleted and empty.\n"+
				"The copy operation for newTextLineCol[%v] produced the\n"+
				"following error:\n"+
				"%v\n",
				ePrefix.String(),
				j,
				j,
				err2.Error())

			return err
		}

	}

	return err
}
