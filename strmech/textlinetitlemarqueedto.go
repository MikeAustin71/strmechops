package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
	"time"
)

//	TextLineTitleMarqueeDto
//
//	This data transfer object (DTO) is designed to store
//	and transfer all specifications necessary to produce
//	a Title Marquee for file output or text display.
//
// This type is compatible with and used by type
// TextLineSpecTitleMarquee
type TextLineTitleMarqueeDto struct {
	StandardTitleLeftMargin string
	//	The standard left margin characters applied
	//	to all Text Title Lines in the 'TitleLines'
	//	array.

	StandardTitleRightMargin string
	//	The standard left margin characters applied
	//	to all Text Title Lines in the 'TitleLines'
	//	array.

	StandardMaxLineLen int
	//	The maximum number of characters allowed on
	//	a text title line.

	StandardTextFieldLen int
	//	The standard field length applied to all
	//	Text Title Lines in the 'TitleLines' array.

	NumLeadingBlankLines int
	//	The number of blank lines or 'new lines'
	//	inserted above the Leading Solid Line.

	LeadingSolidLineChar string
	//	The character used to create the Leading
	//	Solid Line displayed above the Title
	//	Lines.

	NumLeadingSolidLines int
	//	The Number of Leading Solid Lines to
	//	display above the Title Lines.

	NumTopTitleBlankLines int
	//	The number of blank lines or 'new lines' to
	//	insert immediately above the Title Lines
	//	Display.

	TitleLines TextLineSpecLinesCollection
	//	A collection of text line objects containing
	//	all specifications necessary to display the
	//	Text Title Lines.

	NumBottomTitleBlankLines int
	//	The number of blank lines or 'new lines' to
	//	insert immediately below the Title Lines
	//	Display.

	TrailingSolidLineChar string
	//	The character used to create the Trailing
	//	Solid Line displayed below the Title
	//	Lines.

	NumTrailingSolidLines int
	//	The Number of Trailing Solid Lines to
	//	display below the Title Lines.

	NumTrailingBlankLines int
	//	The number of blank lines or 'new lines'
	//	inserted after the Trailing Solid Line.

	lock *sync.Mutex
}

//	AddDateTimeTitleLine
//
//	Adds a Date Time text title line to the text
//	title lines array contained in the current instance
//	of TextLineTitleMarqueeDto.
//
//	Be advised that the left and right margins for this
//	Date Time Title field will automatically use the
//	Text Line Title standard left and right margins.
//
// ----------------------------------------------------------------
//
//	# Terminology
//
//	Maximum Available Text Field Length =
//		TextLineTitleMarqueeDto.StandardMaxLineLen -
//		1 -
//		len(TextLineTitleMarqueeDto.StandardTitleLeftMargin) -
//		len(TextLineTitleMarqueeDto.StandardTitleRightMargin)
//
//	Standard Text Field Length for all text lines is
//		defined at initialization by internal member
//		variable:
//
//			TextLineTitleMarqueeDto.StandardTextFieldLen
//
// ----------------------------------------------------------------
//
//	# Field Length Values And Defaults
//
//		fieldLen > Maximum Available Text Field Length
//			fieldLen = Maximum Available Text Field Length
//
//		fieldLen >= 1 &&
//			fieldLen < len(txtLabel)
//				fieldLen = textLabel string length
//
//		fieldLen = 0
//			fieldLen = Maximum Available Text Field Length
//
//		fieldLen = -1
//			fieldLen = textLabel string length
//
//		fieldLen = -2
//			fieldLen =
//			TextLineTitleMarqueeDto.StandardTextFieldLen
//
//		fieldLen < -2
//			Return value = Error
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	dateTime					time.Time
//
//		A valid date time value which is used to generate
//		a formatted Date/Time text string. Type time.Time
//		is part of the Golang time package:
//	             https://pkg.go.dev/time.
//
//		If this parameter is submitted as a zero value, an
//		error will be returned.
//
//	fieldLen					int
//
//		The length of the text field in which the
//		formatted 'dateTime' value will be displayed.
//
//		If 'fieldLen' is less than the length of the
//		formatted 'dateTime' string, it will be
//		automatically set equal to the formatted
//		'dateTime' string length.
//
//		To automatically set the value of 'fieldLen' to
//		the Maximum Available Text Field Length for this
//		Text Line Title Marquee instance, set this
//		parameter to zero (0).
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
//		Field Length Values And Defaults
//		--------------------------------
//
//		fieldLen > Maximum Available Text Field Length
//			fieldLen = Maximum Available Text Field Length
//
//		fieldLen >= 1 &&
//			fieldLen < len(txtLabel)
//				fieldLen = textLabel string length
//
//		fieldLen = 0
//			fieldLen = Maximum Available Text Field Length
//
//		fieldLen = -1
//			fieldLen = textLabel string length
//
//		fieldLen = -2
//			fieldLen =
//			TextLineTitleMarqueeDto.StandardTextFieldLen
//
//		fieldLen < -2
//			Return value = Error
//
//		If 'fieldLen' is greater than the length of the
//		formatted 'dateTime' string, 'dateTime' will be
//		positioned within a text field with a length equal
//		to 'fieldLen'. In this case, the position of the
//		'dateTime' string within the text field will be
//		controlled by the text justification value
//		contained in parameter, 'textJustification'.
//
//	dateTimeFormat				string
//
//		This string holds the Date/Time format parameters
//		used to format the 'dateTime' value when
//		generating a 'dateTime' text string. The formatted
//		'dateTime' text string is used to display time
//		stamps as a title line in title marquee
//		presentations.
//
//		The Date/Time format is documented in the Golang
//		time.Time package:
//			https://pkg.go.dev/time
//
//		The format operations are documented at:
//			https://pkg.go.dev/time#Time.Format
//
//		If this parameter is submitted as an empty string,
//		parameter 'dateTimeFormat' will be assigned a
//		default value of:
//			"2006-01-02 15:04:05.000000000 -0700 MST".
//
//		Example Formats:
//		 Example 1:
//		  dateTimeFormat =
//		   "2006-01-02 15:04:05.000000000 -0700 MST"
//		 Result =
//		   "2021-10-21 14:19:03.000000000 -0500 CDT"
//
//		 Example 2:
//		  dateTimeFormat =
//		   "Monday January 2, 2006 15:04:05.000000000 -0700 MST"
//
//		   Result =
//		    "Thursday October 21, 2021 14:19:03.000000000 -0500 CDT"
//
//	textJustification			TextJustify
//
//		An enumeration which specifies the justification
//		of the 'dateTime' string within a text field.
//
//		The text field length is taken from input
//		parameter 'fieldLen'.
//
//		Text justification can only be evaluated in the
//		context of a 'dateTime' text string, field length
//		and a 'textJustification' object of type
//		TextJustify. This is because a field length
//		('fieldLen') value equal to or less than the
//		length of the 'dateTime' text string will never
//		use text justification. In these cases, text
//		justification is completely	ignored because the
//		length of the text field is automatically set
//		equal to the length of the 'dateTime' text
//		string.
//
//		If the field length ('fieldLen') is greater than
//		the length of the 'dateTime' text string, text
//		justification must be equal to one of these three
//		valid values:
//
//		    TextJustify(0).Left()
//		    TextJustify(0).Right()
//		    TextJustify(0).Center()
//
//		Users can also apply the abbreviated text
//		justification enumeration syntax as follows:
//
//		    TxtJustify.Left()
//		    TxtJustify.Right()
//		    TxtJustify.Center()
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
func (txtLineTitleMarqueeDto *TextLineTitleMarqueeDto) AddDateTimeTitleLine(
	dateTime time.Time,
	fieldLen int,
	dateTimeFormat string,
	textJustification TextJustify,
	errorPrefix interface{}) error {

	if txtLineTitleMarqueeDto.lock == nil {
		txtLineTitleMarqueeDto.lock = new(sync.Mutex)
	}

	txtLineTitleMarqueeDto.lock.Lock()

	defer txtLineTitleMarqueeDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineTitleMarqueeDto.AddTextLabelTitleLine()",
		"")

	if err != nil {
		return err
	}

	lenStr := len(dateTimeFormat)

	if lenStr == 0 {
		dateTimeFormat = new(textSpecificationMolecule).
			getDefaultDateTimeFormat()
	}

	stdLine := TextLineSpecStandardLine{}.New()

	fieldLen,
		err = new(textLineTitleMarqueeDtoMechanics).
		calcTextFieldLen(
			txtLineTitleMarqueeDto,
			fieldLen,
			ePrefix)

	if err != nil {
		return err
	}

	lenStr = len(txtLineTitleMarqueeDto.StandardTitleLeftMargin)

	// Left Margin Label
	if lenStr > 0 {

		_,
			err = stdLine.AddTextFieldLabel(
			txtLineTitleMarqueeDto.StandardTitleLeftMargin,
			lenStr,
			TxtJustify.Left(),
			ePrefix.XCpy(
				"StandardTitleLeftMargin"))

		if err != nil {
			return err
		}

	}

	_,
		err = stdLine.AddTextFieldDateTime(
		dateTime,
		fieldLen,
		dateTimeFormat,
		textJustification,
		ePrefix.XCpy(
			"dateTime Label"))

	lenStr =
		len(txtLineTitleMarqueeDto.StandardTitleRightMargin)

	if lenStr > 0 {

		_,
			err = stdLine.AddTextFieldLabel(
			txtLineTitleMarqueeDto.StandardTitleRightMargin,
			lenStr,
			TxtJustify.Left(),
			ePrefix.XCpy(
				"StandardTitleRightMargin"))

		if err != nil {
			return err
		}

	}

	err = stdLine.IsValidInstanceError(
		ePrefix.XCpy(
			"stdLine<-dateTime"))

	if err != nil {

		return err
	}

	err = txtLineTitleMarqueeDto.TitleLines.AddTextLineSpec(
		&stdLine,
		ePrefix.XCpy(
			"stdLine<-"))

	return err
}

//	AddTextLabelTitleLine
//
//	Adds a text label title line to the text lines
//	array contained in the current instance of
//	TextLineTitleMarqueeDto.
//
//	Be advised that the left and right margins for this
//	Date Time Title field will automatically use the
//	Text Line Title standard left and right margins.
//
// ----------------------------------------------------------------
//
//	# Terminology
//
//	Maximum Available Text Field Length =
//		TextLineTitleMarqueeDto.StandardMaxLineLen -
//		1 -
//		len(TextLineTitleMarqueeDto.StandardTitleLeftMargin) -
//		len(TextLineTitleMarqueeDto.StandardTitleRightMargin)
//
//	Standard Text Field Length for all text lines is
//		defined at initialization by internal member
//		variable:
//
//			TextLineTitleMarqueeDto.StandardTextFieldLen
//
// ----------------------------------------------------------------
//
//	# Field Length Values And Defaults
//
//		fieldLen > Maximum Available Text Field Length
//			fieldLen = Maximum Available Text Field Length
//
//		fieldLen >= 1 &&
//			fieldLen < len(txtLabel)
//				fieldLen = textLabel string length
//
//		fieldLen = 0
//			fieldLen = Maximum Available Text Field Length
//
//		fieldLen = -1
//			fieldLen = textLabel string length
//
//		fieldLen = -2
//			fieldLen =
//			TextLineTitleMarqueeDto.StandardTextFieldLen
//
//		fieldLen < -2
//			Return value = Error
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	textLabel					string
//
//		The string content to be displayed within the
//		text label.
//
//		If this parameter is submitted as a zero length
//		string, an error will be returned.
//
//	fieldLen					int
//
//		The length of the text field in which the
//		'textLabel' value will be displayed.
//
//		If 'fieldLen' is less than the length of the
//		formatted 'textLabel' string, it will be
//		automatically set equal to the 'textLabel'
//		string length.
//
//		To automatically set the value of 'fieldLen' to
//		the Maximum Available Text Field Length for this
//		Text Line Title Marquee instance, set this
//		parameter to zero (0).
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
//		Standard Text Field Length for all text lines is
//		defined at initialization by internal member
//		variable:
//
//			TextLineTitleMarqueeDto.StandardTextFieldLen
//
//		Field Length Values And Defaults
//		--------------------------------
//
//		fieldLen > Maximum Available Text Field Length
//			fieldLen = Maximum Available Text Field Length
//
//		fieldLen >= 1 &&
//			fieldLen < len(txtLabel)
//				fieldLen = textLabel string length
//
//		fieldLen = 0
//			fieldLen = Maximum Available Text Field Length
//
//		fieldLen = -1
//			fieldLen = textLabel string length
//
//		fieldLen = -2
//			fieldLen =
//			TextLineTitleMarqueeDto.StandardTextFieldLen
//
//		fieldLen < -2
//			Return value = Error
//
//		If 'fieldLen' is greater than the length of the
//		'textLabel' string, 'textLabel' will be
//		positioned within a text field with a length equal
//		to 'fieldLen'. In this case, the position of the
//		'textLabel' string within the text field will be
//		controlled by the text justification value
//		contained in parameter, 'textJustification'.
//
//	textJustification			TextJustify
//
//		An enumeration which specifies the justification
//		of the 'textLabel' string within a text field.
//
//		The text field length is taken from input
//		parameter 'fieldLen'.
//
//		Text justification can only be evaluated in the
//		context of a text label, field length and a Text
//		Justification object of type TextJustify. This is
//		because text labels with a field length equal to
//		or less than the length of the text label will
//		never use text justification. In these cases,
//		text justification is completely ignored.
//
//
//		If the field length ('fieldLen') is greater than
//		the length of the 'textLabel' string, text
//		justification must be equal to one of these three
//		valid values:
//
//	         TextJustify(0).Left()
//	         TextJustify(0).Right()
//	         TextJustify(0).Center()
//
//		Users can also use the abbreviated text
//		justification enumeration syntax as follows:
//
//	         TxtJustify.Left()
//	         TxtJustify.Right()
//	         TxtJustify.Center()
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
func (txtLineTitleMarqueeDto *TextLineTitleMarqueeDto) AddTextLabelTitleLine(
	textLabel string,
	fieldLen int,
	textJustification TextJustify,
	errorPrefix interface{}) error {

	if txtLineTitleMarqueeDto.lock == nil {
		txtLineTitleMarqueeDto.lock = new(sync.Mutex)
	}

	txtLineTitleMarqueeDto.lock.Lock()

	defer txtLineTitleMarqueeDto.lock.Unlock()
	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineTitleMarqueeDto."+
			"AddTextLabelTitleLine()",
		"")

	if err != nil {
		return err
	}

	stdLine := TextLineSpecStandardLine{}.New()

	if len(textLabel) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'textLabel' is INVALID!\n"+
			"'textLabel' is an empty string with zero text characters!\n",
			ePrefix.String())

		return err

	}

	fieldLen,
		err = new(textLineTitleMarqueeDtoMechanics).
		calcTextFieldLen(
			txtLineTitleMarqueeDto,
			fieldLen,
			ePrefix)

	if err != nil {
		return err
	}

	lenStr := len(txtLineTitleMarqueeDto.StandardTitleLeftMargin)

	// Left Margin Label
	if lenStr > 0 {

		_,
			err = stdLine.AddTextFieldLabel(
			txtLineTitleMarqueeDto.StandardTitleLeftMargin,
			lenStr,
			TxtJustify.Left(),
			ePrefix.XCpy(
				"StandardTitleLeftMargin"))

		if err != nil {
			return err
		}

	}

	_,
		err = stdLine.AddTextFieldLabel(
		textLabel,
		fieldLen,
		textJustification,
		ePrefix.XCpy(
			"textLabel"))

	lenStr =
		len(txtLineTitleMarqueeDto.StandardTitleRightMargin)

	if lenStr > 0 {

		_,
			err = stdLine.AddTextFieldLabel(
			txtLineTitleMarqueeDto.StandardTitleRightMargin,
			lenStr,
			TxtJustify.Left(),
			ePrefix.XCpy(
				"StandardTitleRightMargin"))

		if err != nil {
			return err
		}

	}

	err = stdLine.IsValidInstanceError(
		ePrefix.XCpy(
			"stdLine<-textLabel"))

	if err != nil {

		return err
	}

	err = txtLineTitleMarqueeDto.TitleLines.AddTextLineSpec(
		&stdLine,
		ePrefix.XCpy(
			"stdLine<-"))

	return err
}

//	AddTitleLineTextFields
//
//	Receives an array of text fields which are used to
//	construct a single Title Text Line (type
//	TextLineSpecStandardLine) which will be added to the
//	Title Lines array maintained by the current instance
//	of TextLineTitleMarqueeDto.
//
//	Be advised that the left and right margins for this
//	newly constructed title line will automatically apply
//	the Text Line Title standard left and right margins.
//
//	This method gives maximum flexibility in customizing
//	individual lines to text added to the internal title
//	lines array:
//
//		TextLineTitleMarqueeDto.TitleLines
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	textFields					*[]ITextFieldSpecification
//
//		A pointer to a text field collection whose
//		objects implement the ITextFieldSpecification
//		interface. A deep copy of each object in this
//		collection will be added to the text field
//		collection of a TextLineSpecStandardLine
//		instance which will in turn be added to the
//		current TextLineTitleMarqueeDto Title Lines
//		array.
//
//		If member variable data values contained in this
//		'textFields' parameter are found to be invalid,
//		an error will be returned.
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
func (txtLineTitleMarqueeDto *TextLineTitleMarqueeDto) AddTitleLineTextFields(
	textFields *[]ITextFieldSpecification,
	errorPrefix interface{}) error {

	if txtLineTitleMarqueeDto.lock == nil {
		txtLineTitleMarqueeDto.lock = new(sync.Mutex)
	}

	txtLineTitleMarqueeDto.lock.Lock()

	defer txtLineTitleMarqueeDto.lock.Unlock()
	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineTitleMarqueeDto."+
			"AddTitleLineTextFields()",
		"")

	if err != nil {
		return err
	}

	if textFields == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'textFields' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if len(*textFields) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'textFields' is INVALID!\n"+
			"'textFields' is an empty array with zero text fields.\n",
			ePrefix.String())

		return err
	}

	stdLine := TextLineSpecStandardLine{}.New()

	lenStr := len(txtLineTitleMarqueeDto.StandardTitleLeftMargin)

	// Left Margin Label
	if lenStr > 0 {

		_,
			err = stdLine.AddTextFieldLabel(
			txtLineTitleMarqueeDto.StandardTitleLeftMargin,
			lenStr,
			TxtJustify.Left(),
			ePrefix.XCpy(
				"StandardTitleLeftMargin"))

		if err != nil {
			return err
		}

	}

	_,
		err = stdLine.AddTextFields(
		textFields,
		ePrefix)

	if err != nil {

		return err
	}

	lenStr =
		len(txtLineTitleMarqueeDto.StandardTitleRightMargin)

	if lenStr > 0 {

		_,
			err = stdLine.AddTextFieldLabel(
			txtLineTitleMarqueeDto.StandardTitleRightMargin,
			lenStr,
			TxtJustify.Left(),
			ePrefix.XCpy(
				"StandardTitleRightMargin"))

		if err != nil {
			return err
		}

	}

	err = stdLine.IsValidInstanceError(
		ePrefix.XCpy(
			"stdLine<-textFields"))

	if err != nil {

		return err
	}

	err = txtLineTitleMarqueeDto.TitleLines.AddTextLineSpec(
		&stdLine,
		ePrefix.XCpy(
			"stdLine<-"))

	return err
}

//	AddStandardTitleLine
//
//	Adds a deep copy of a TextLineSpecStandardLine object
//	to the member variable title lines array maintained
//	by the current instance of TextLineTitleMarqueeDto.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	stdTitleLine				*TextLineSpecStandardLine
//
//		A pointer to an instance of
//		TextLineSpecStandardLine. A deep copy of this
//		object will be added to the member variable title
//		lines array maintained by the current instance of
//		TextLineTitleMarqueeDto:
//			TextLineTitleMarqueeDto.TitleLines
//
//		If this instance of TextLineSpecStandardLine is
//		judged to be invalid, an error will be returned.
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
func (txtLineTitleMarqueeDto *TextLineTitleMarqueeDto) AddStandardTitleLine(
	stdTitleLine *TextLineSpecStandardLine,
	errorPrefix interface{}) error {

	if txtLineTitleMarqueeDto.lock == nil {
		txtLineTitleMarqueeDto.lock = new(sync.Mutex)
	}

	txtLineTitleMarqueeDto.lock.Lock()

	defer txtLineTitleMarqueeDto.lock.Unlock()
	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineTitleMarqueeDto."+
			"AddStandardTitleLine()",
		"")

	if err != nil {
		return err
	}

	_,
		err = textLineSpecStandardLineAtom{}.ptr().
		testValidityOfTextLineSpecStdLine(
			stdTitleLine,
			false, // allowZeroLengthTextFieldsArray
			ePrefix.XCpy("stdTitleLine"))

	if err != nil {
		return err
	}

	var deepCopyStdTitleLine TextLineSpecStandardLine

	deepCopyStdTitleLine,
		err = stdTitleLine.CopyOut(
		ePrefix.XCpy("deepCopyStdTitleLine<-"))

	if err != nil {
		return err
	}

	err = deepCopyStdTitleLine.IsValidInstanceError(
		ePrefix.XCpy(
			"deepCopyStdTitleLine<-textLabel"))

	if err != nil {

		return err
	}

	err = txtLineTitleMarqueeDto.TitleLines.AddTextLineSpec(
		&deepCopyStdTitleLine,
		ePrefix.XCpy(
			"deepCopyStdTitleLine<-"))

	return err
}

//	Empty
//
//	Resets all internal member variables for the current
//	instance of TextLineTitleMarqueeDto to their zero or
//	uninitialized states.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete all member variable data
//	values in the current instance of
//	TextLineTitleMarqueeDto.
//
//	All member variable data values in this instance will
//	be reset to their zero or uninitialized states.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	NONE
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	NONE
func (txtLineTitleMarqueeDto *TextLineTitleMarqueeDto) Empty() {

	if txtLineTitleMarqueeDto.lock == nil {
		txtLineTitleMarqueeDto.lock = new(sync.Mutex)
	}

	txtLineTitleMarqueeDto.lock.Lock()

	new(textLineTitleMarqueeDtoMechanics).empty(
		txtLineTitleMarqueeDto)

	txtLineTitleMarqueeDto.lock.Unlock()

	txtLineTitleMarqueeDto.lock = nil

	return
}

// IsValidInstance - Performs a diagnostic review of the data
// values encapsulated in the current TextLineSpecLinesCollection
// instance to determine if they are valid.
//
// If all data element evaluate as valid, this method returns
// 'true'. If any data element is invalid, this method returns
// 'false'.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
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
//		instance of TextLineTitleMarqueeDto are valid,
//		this returned boolean value is set to 'true'. If
//		any data values are invalid, this return
//		parameter is set to 'false'.
func (txtLineTitleMarqueeDto *TextLineTitleMarqueeDto) IsValidInstance() (
	isValid bool) {

	if txtLineTitleMarqueeDto.lock == nil {
		txtLineTitleMarqueeDto.lock = new(sync.Mutex)
	}

	txtLineTitleMarqueeDto.lock.Lock()

	defer txtLineTitleMarqueeDto.lock.Unlock()

	isValid,
		_ = new(textLineTitleMarqueeDtoMechanics).
		testValidityOfTitleMarqueeDto(
			txtLineTitleMarqueeDto,
			nil)

	return isValid
}

//	IsValidInstanceError
//
//	Performs a diagnostic review of the data values
//	encapsulated in the current TextLineTitleMarqueeDto
//	instance to determine if they are valid.
//
//	If any data element evaluates as invalid, this method
//	will return an error.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
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
func (txtLineTitleMarqueeDto *TextLineTitleMarqueeDto) IsValidInstanceError(
	errorPrefix interface{}) error {

	if txtLineTitleMarqueeDto.lock == nil {
		txtLineTitleMarqueeDto.lock = new(sync.Mutex)
	}

	txtLineTitleMarqueeDto.lock.Lock()

	defer txtLineTitleMarqueeDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineTitleMarqueeDto."+
			"IsValidInstanceError()",
		"")

	if err != nil {
		return err
	}

	_,
		err = new(textLineTitleMarqueeDtoMechanics).
		testValidityOfTitleMarqueeDto(
			txtLineTitleMarqueeDto,
			ePrefix.XCpy(
				"txtLineTitleMarqueeDto"))

	return err
}

//	textLineTitleMarqueeDtoMechanics
//
//	Provides helper methods for type
//	TextLineTitleMarqueeDto

type textLineTitleMarqueeDtoMechanics struct {
	lock *sync.Mutex
}

//	calcTextFieldLen
//
//	Receives a field length value and proceeds to validate
//	that value setting automatic default values where
//	required.
//
//	The validated or default field length value is then
//	returned to the calling function.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	txtTitleMarqueeDto 			*TextLineTitleMarqueeDto
//
//		A pointer to an instance of TextLineTitleMarqueeDto.
//		No data elements in this instance will be modified.
//
//		The internal member data elements contained in this
//		instance will be used to compute a valid value for
//		text field length passed as input paramter 'fieldLen'.
//
//	fieldLen					int
//
//		A text field length value which will be validated.
//		In certain cases, default values will be applied
//		and returned to the calling function.
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
//		Field Length Values And Defaults
//		--------------------------------
//
//		fieldLen > Maximum Available Text Field Length
//			fieldLen = Maximum Available Text Field Length
//
//		fieldLen >= 1 &&
//			fieldLen < len(txtLabel)
//				fieldLen = textLabel string length
//
//		fieldLen = 0
//			fieldLen = Maximum Available Text Field Length
//
//		fieldLen = -1
//			fieldLen = textLabel string length
//
//		fieldLen = -2
//			fieldLen =
//			TextLineTitleMarqueeDto.StandardTextFieldLen
//
//		fieldLen < -2
//			Return value = Error
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
//	validFieldLen				int
//
//		If this method completes successfully, valid
//		text field length value will be returned.
//
//		Be advised that the returned value may not
//		equal the original 'fieldLen' input parameter,
//		as default values may be applied in certain
//		cases.
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
//		Field Length Values And Defaults
//		--------------------------------
//
//		fieldLen > Maximum Available Text Field Length
//			fieldLen = Maximum Available Text Field Length
//
//		fieldLen >= 1 &&
//			fieldLen < len(txtLabel)
//				fieldLen = textLabel string length
//
//		fieldLen = 0
//			fieldLen = Maximum Available Text Field Length
//
//		fieldLen = -1
//			fieldLen = textLabel string length
//
//		fieldLen = -2
//			fieldLen =
//			TextLineTitleMarqueeDto.StandardTextFieldLen
//
//		fieldLen < -2
//			Return value = Error
//
//	err							error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (txtTitleDtoMech *textLineTitleMarqueeDtoMechanics) calcTextFieldLen(
	txtTitleMarqueeDto *TextLineTitleMarqueeDto,
	fieldLen int,
	errPrefDto *ePref.ErrPrefixDto) (
	validFieldLen int,
	err error) {

	if txtTitleDtoMech.lock == nil {
		txtTitleDtoMech.lock = new(sync.Mutex)
	}

	txtTitleDtoMech.lock.Lock()

	defer txtTitleDtoMech.lock.Unlock()

	validFieldLen = 0

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineTitleMarqueeDtoMechanics."+
			"calcTextFieldLen()",
		"")

	if err != nil {
		return validFieldLen, err
	}

	if txtTitleMarqueeDto == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtTitleMarqueeDto' is a nil pointer!\n",
			ePrefix.String())

		return validFieldLen, err
	}

	if txtTitleMarqueeDto.StandardMaxLineLen < 1 {

		err = fmt.Errorf("%v\n"+
			"Error: The TextLineTitleMarqueeDto object ('txtTitleMarqueeDto') is INVALID!\n"+
			"'txtTitleMarqueeDto.StandardMaxLineLen' has NOT been properly configured.\n"+
			"'txtTitleMarqueeDto.StandardMaxLineLen' has a value less than one (1).\n"+
			"txtTitleMarqueeDto.StandardMaxLineLen = %v\n",
			ePrefix.String(),
			txtTitleMarqueeDto.StandardMaxLineLen)

		return validFieldLen, err

	}

	if fieldLen < -2 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fieldLen' is INVALID!\n"+
			"'fieldLen' has a value less than -2\n"+
			"fieldLen = %v\n",
			ePrefix.String(),
			fieldLen)

		return validFieldLen, err

	}

	maxAvailableTextFieldLen :=
		txtTitleMarqueeDto.StandardMaxLineLen -
			1 -
			len(txtTitleMarqueeDto.StandardTitleLeftMargin) -
			len(txtTitleMarqueeDto.StandardTitleRightMargin)

	if maxAvailableTextFieldLen < 1 {

		err = fmt.Errorf("%v\n"+
			"Error: 'maxAvailableTextFieldLen' produces a value less than one (1).\n"+
			"maxAvailableTextFieldLen = \n"+
			"	StandardMaxLineLen - Left Margin Length - Right Margin Length -1\n"+
			"StandardMaxLineLen is probably invalid."+
			"StandardMaxLineLen  = %v\n"+
			"Left Margin Length  = %v\n"+
			"Right Margin Length = %v\n",
			ePrefix.String(),
			txtTitleMarqueeDto.StandardMaxLineLen,
			len(txtTitleMarqueeDto.StandardTitleLeftMargin),
			len(txtTitleMarqueeDto.StandardTitleRightMargin))

		return validFieldLen, err

	}

	if fieldLen > maxAvailableTextFieldLen ||
		fieldLen == 0 {

		validFieldLen = maxAvailableTextFieldLen
	}

	if fieldLen == -2 {
		fieldLen = txtTitleMarqueeDto.StandardTextFieldLen
	}

	return validFieldLen, err
}

//	empty
//
//	Resets all internal member variables for the
//	TextLineTitleMarqueeDto input parameter
//	'txtTitleMarqueeDto' to their zero or uninitialized
//	states.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete all member variable data
//	values in instance of TextLineTitleMarqueeDto passed
//	as input parameter 'txtTitleMarqueeDto'.
//
//	All member variable data values in this instance will
//	be reset to their zero or uninitialized states.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//
//	txtTitleMarqueeDto 			*TextLineTitleMarqueeDto
//
//		A pointer to an instance of TextLineTitleMarqueeDto.
//
//		All the data elements in this instance will be
//		deleted and reset to their zero values or
//		uninitialized states.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	NONE
func (txtTitleDtoMech *textLineTitleMarqueeDtoMechanics) empty(
	txtTitleMarqueeDto *TextLineTitleMarqueeDto) {

	if txtTitleDtoMech.lock == nil {
		txtTitleDtoMech.lock = new(sync.Mutex)
	}

	txtTitleDtoMech.lock.Lock()

	defer txtTitleDtoMech.lock.Unlock()

	if txtTitleMarqueeDto == nil {
		return
	}

	txtTitleMarqueeDto.StandardTitleLeftMargin = ""

	txtTitleMarqueeDto.StandardTitleRightMargin = ""

	txtTitleMarqueeDto.StandardMaxLineLen = 0

	txtTitleMarqueeDto.StandardTextFieldLen = 0

	txtTitleMarqueeDto.NumLeadingBlankLines = 0

	txtTitleMarqueeDto.LeadingSolidLineChar = ""

	txtTitleMarqueeDto.NumLeadingSolidLines = 0

	txtTitleMarqueeDto.NumTopTitleBlankLines = 0

	txtTitleMarqueeDto.TitleLines.Empty()

	txtTitleMarqueeDto.NumBottomTitleBlankLines = 0

	txtTitleMarqueeDto.TrailingSolidLineChar = ""

	txtTitleMarqueeDto.NumTrailingSolidLines = 0

	txtTitleMarqueeDto.NumTrailingBlankLines = 0
}

//	testValidityOfTitleMarqueeDto
//
//	Receives a pointer to an instance of
//	TextLineTitleMarqueeDto and performs a diagnostic
//	analysis to determine if that instance is valid in
//	all respects.
//
//	If the input parameter 'txtTitleMarqueeDto' is
//	determined to be invalid, this method will return a
//	boolean flag ('isValid') of 'false'. In addition, an
//	instance of type error ('err') will be returned
//	configured with an appropriate error message.
//
//	If the input parameter 'txtTitleMarqueeDto' is valid,
//	this method will return a boolean flag ('isValid') of
//	'true' and the returned error type ('err') will be
//	set to 'nil'.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	txtTitleMarqueeDto 			*TextLineTitleMarqueeDto
//
//		A pointer to an instance of TextLineTitleMarqueeDto.
//		No data elements in this instance will be modified.
//
//		The internal member data elements contained in this
//		instance will be analyzed to determine if they are
//		valid in all respects.
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
// ----------------------------------------------------------------
//
// # Return Values
//
//	isValid                    bool
//
//		If input parameter 'txtTitleMarqueeDto' is judged
//		to be valid in all respects, this return parameter
//		will be set to 'true'.
//
//		If input parameter 'txtTitleMarqueeDto' is found to
//		be invalid, this return parameter will be set to
//		'false'.
//
//	err							error
//
//		If input parameter 'txtTitleMarqueeDto' is judged
//		to be valid in all respects, the returned error
//		Type is set equal to 'nil'.
//
//		If input parameter 'txtTitleMarqueeDto' is found
//		to be invalid, the returned error Type will
//		encapsulate an appropriate error message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (txtTitleDtoMech *textLineTitleMarqueeDtoMechanics) testValidityOfTitleMarqueeDto(
	txtTitleMarqueeDto *TextLineTitleMarqueeDto,
	errPrefDto *ePref.ErrPrefixDto) (
	isValid bool,
	err error) {

	if txtTitleDtoMech.lock == nil {
		txtTitleDtoMech.lock = new(sync.Mutex)
	}

	txtTitleDtoMech.lock.Lock()

	defer txtTitleDtoMech.lock.Unlock()

	isValid = false

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineTitleMarqueeDtoMechanics."+
			"testValidityOfTitleMarqueeDto()",
		"")

	if err != nil {
		return isValid, err
	}

	if txtTitleMarqueeDto == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtTitleMarqueeDto' is a nil pointer!\n",
			ePrefix.String())

		return isValid, err
	}

	if txtTitleMarqueeDto.StandardMaxLineLen < 1 {

		err = fmt.Errorf("%v\n"+
			"Error: The TextLineTitleMarqueeDto contains invalid data values!\n"+
			"'StandardMaxLineLen' has NOT been properly configured.\n"+
			"'StandardMaxLineLen' has a value less than one (1).\n"+
			"StandardMaxLineLen = %v\n",
			ePrefix.String(),
			txtTitleMarqueeDto.StandardMaxLineLen)

		return isValid, err
	}

	if txtTitleMarqueeDto.StandardTextFieldLen < 1 {

		err = fmt.Errorf("%v\n"+
			"Error: The TextLineTitleMarqueeDto contains invalid data values!\n"+
			"'StandardTextFieldLen' has NOT been properly configured.\n"+
			"'StandardTextFieldLen' has a value less than one (1).\n"+
			"StandardTextFieldLen = %v\n",
			ePrefix.String(),
			txtTitleMarqueeDto.StandardTextFieldLen)

		return isValid, err
	}

	maxAvailableTextFieldLen :=
		txtTitleMarqueeDto.StandardMaxLineLen -
			1 -
			len(txtTitleMarqueeDto.StandardTitleLeftMargin) -
			len(txtTitleMarqueeDto.StandardTitleRightMargin)

	if maxAvailableTextFieldLen < 1 {

		err = fmt.Errorf("%v\n"+
			"Error: The TextLineTitleMarqueeDto contains invalid data values!\n"+
			"'maxAvailableTextFieldLen' produces a value less than one (1).\n"+
			"maxAvailableTextFieldLen = \n"+
			"	StandardMaxLineLen - Left Margin Length - Right Margin Length -1\n"+
			"StandardMaxLineLen is probably invalid."+
			"StandardMaxLineLen  = %v\n"+
			"Left Margin Length  = %v\n"+
			"Right Margin Length = %v\n",
			ePrefix.String(),
			txtTitleMarqueeDto.StandardMaxLineLen,
			len(txtTitleMarqueeDto.StandardTitleLeftMargin),
			len(txtTitleMarqueeDto.StandardTitleRightMargin))

		return isValid, err
	}

	if txtTitleMarqueeDto.StandardTextFieldLen > maxAvailableTextFieldLen {

		txtTitleMarqueeDto.StandardTextFieldLen =
			maxAvailableTextFieldLen
	}

	if txtTitleMarqueeDto.NumLeadingBlankLines < 0 {

		txtTitleMarqueeDto.NumLeadingBlankLines = 0
	}

	if txtTitleMarqueeDto.NumLeadingSolidLines < 0 {

		txtTitleMarqueeDto.NumLeadingSolidLines = 0
	}

	if txtTitleMarqueeDto.NumTopTitleBlankLines < 0 {

		txtTitleMarqueeDto.NumTopTitleBlankLines = 0
	}

	if txtTitleMarqueeDto.NumBottomTitleBlankLines < 0 {

		txtTitleMarqueeDto.NumBottomTitleBlankLines = 0
	}

	if txtTitleMarqueeDto.NumTrailingSolidLines < 0 {

		txtTitleMarqueeDto.NumTrailingSolidLines = 0
	}

	if txtTitleMarqueeDto.NumTrailingBlankLines < 0 {

		txtTitleMarqueeDto.NumTrailingBlankLines = 0
	}

	if txtTitleMarqueeDto.NumLeadingSolidLines > 0 &&
		len(txtTitleMarqueeDto.LeadingSolidLineChar) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: The TextLineTitleMarqueeDto contains invalid data values!\n"+
			"'LeadingSolidLineChar' has NOT been properly configured.\n"+
			"'NumLeadingSolidLines' has a value greater than zero, but"+
			"'LeadingSolidLineChar' is an empty string.\n"+
			"NumLeadingSolidLines = %v\n",
			ePrefix.String(),
			txtTitleMarqueeDto.NumLeadingSolidLines)

		return isValid, err

	}

	if txtTitleMarqueeDto.NumTrailingSolidLines > 0 &&
		len(txtTitleMarqueeDto.TrailingSolidLineChar) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: The TextLineTitleMarqueeDto contains invalid data values!\n"+
			"'TrailingSolidLineChar' has NOT been properly configured.\n"+
			"'NumTrailingSolidLines' has a value greater than zero, but"+
			"'TrailingSolidLineChar' is an empty string.\n"+
			"NumTrailingSolidLines = %v\n",
			ePrefix.String(),
			txtTitleMarqueeDto.NumTrailingSolidLines)

		return isValid, err

	}

	isValid = true

	return isValid, err
}
