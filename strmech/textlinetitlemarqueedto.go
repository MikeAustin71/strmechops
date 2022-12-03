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
	//	Display above the Title Lines.

	NumTopTitleBlankLines int
	//	The number of blank lines or 'new lines' to
	//	insert immediately above the Title Lines
	//	Display.

	TitleLines []TextLineSpecStandardLine
	//	An array of TextLineSpecStandardLine objects
	//	containing all specifications necessary to
	//	display the Text Title Lines.

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
	//	Display below the Title Lines.

	NumTrailingBlankLines int
	//	The number of blank lines or 'new lines'
	//	inserted after the Trailing Solid Line.

	lock *sync.Mutex
}

//	AddDateTimeTitleLine
//
//	Adds a Date Time text line to the text lines
//	array contained in the current instance of
//	TextLineTitleMarqueeDto.
//
//	The generated Date Time Text Field will use the
//	standard field length configured Text Line Title
//	Marquee:
//		TextLineTitleMarqueeDto.StandardTextFieldLen
//
//	Be advised that the left and right margins for this
//	Date Time Title field will automatically use the
//	Text Line Title standard left and right margins.
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
//		The text field length is taken from the Title
//		Marquee standard field length.
//
//		Text justification can only be evaluated in the
//		context of a 'dateTime' text string, field length
//		and a 'textJustification' object of type
//		TextJustify. This is because a field length
//		('TextLineTitleMarqueeDto.StandardTextFieldLen') value
//		equal to or less than the length of the 'dateTime'
//		text string will never use text justification. In
//		these cases, text justification is completely
//		ignored because the length of the text field
//		is automatically set equal to the length of
//		the 'dateTime' text string.
//
//		If the field length is greater than the length of
//		the text label, text justification must be equal
//		to one of these	three valid values:
//		    TextJustify(0).Left()
//		    TextJustify(0).Right()
//		    TextJustify(0).Center()
//
//		You can also use the abbreviated text
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
		"TextLineTitleMarqueeDto.AddLabelTitleLine()",
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
		txtLineTitleMarqueeDto.StandardTextFieldLen,
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

	txtLineTitleMarqueeDto.TitleLines =
		append(
			txtLineTitleMarqueeDto.TitleLines,
			stdLine)

	return err
}

//	AddLabelTitleLine
//
//	Adds a text label title line to the text lines
//	array contained in the current instance of
//	TextLineTitleMarqueeDto.
//
//	This text label will use the standard field length
//	configured Text Line Title Marquee:
//		TextLineTitleMarqueeDto.StandardTextFieldLen
//
//	Be advised that the left and right margins for this
//	Date Time Title field will automatically use the
//	Text Line Title standard left and right margins.
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
//	textJustification			TextJustify
//
//		An enumeration which specifies the justification
//		of the 'textLabel' string within the text field
//		specified by the standard Text Line Title field
//		length:
//			TextLineTitleMarqueeDto.StandardTextFieldLen
//
//		Text justification can only be evaluated in the
//		context of a text label, field length and a Text
//		Justification object of type TextJustify. This is
//		because text labels with a field length equal to
//		or less than the length of the text label will
//		never use text justification. In these cases,
//		text justification is completely ignored.
//
//		If the standard Text Line Title field length
//	    (TextLineTitleMarqueeDto.StandardTextFieldLen) is greater
//	    than the length of the text label, text
//	    justification must be equal to one of these three
//	    valid values:
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
func (txtLineTitleMarqueeDto *TextLineTitleMarqueeDto) AddLabelTitleLine(
	textLabel string,
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
		"TextLineTitleMarqueeDto.AddLabelTitleLine()",
		"")

	if err != nil {
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
		err = stdLine.AddTextFieldLabel(
		textLabel,
		txtLineTitleMarqueeDto.StandardTextFieldLen,
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

	txtLineTitleMarqueeDto.TitleLines =
		append(
			txtLineTitleMarqueeDto.TitleLines,
			stdLine)

	return err
}

//	AddTitleLineTextFields
//
//	Receives an array of text fields which are used to
//	construct a single Title Text Line which will be
//	added to the Title Lines array maintained by the
//	current instance of TextLineTitleMarqueeDto.
//
//	Be advised that the left and right margins for this
//	newly constructed title line will automatically use
//	the Text Line Title standard left and right margins.
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

	txtLineTitleMarqueeDto.TitleLines =
		append(
			txtLineTitleMarqueeDto.TitleLines,
			stdLine)

	return err
}
