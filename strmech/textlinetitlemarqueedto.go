package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
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
	StandardSolidLineLeftMargin string
	//	The standard left margin characters applied
	//	to all Solid Lines created for this Title
	//	Marquee

	StandardSolidLineRightMargin string
	//	The standard right margin characters applied
	//	to all Solid Lines created for this Title
	//	Marquee

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
	//
	//	'StandardTextFieldLen' defines the length of the
	//	text field in which the Title Line string will be
	//	displayed. If 'StandardTextFieldLen' is less
	//	than the length of the Title Line string, it will
	//	be automatically set equal to the Title Line
	//	string length.
	//
	//	To automatically set the value of
	//	'StandardTextFieldLen' to the length of the Title
	//	Line string, set this parameter to a value of
	//	minus one (-1).
	//
	//	Field Length Examples
	//
	//		Example-1
	//          Title Line String = "Hello World!"
	//			Title Line String Length = 12
	//			StandardTextFieldLen = 18
	//			StandardTextJustification = TxtJustify.Center()
	//			Formatted Title Line String =
	//				"   Hello World!   "
	//
	//		Example-2
	//          Title Line String = "Hello World!"
	//			Title Line String Length = 12
	//			StandardTextFieldLen = 18
	//			StandardTextJustification = TxtJustify.Left()
	//			Formatted Title Line String =
	//				"Hello World!      "
	//
	//		Example-3
	//          Title Line String = "Hello World!"
	//			Title Line String Length = 12
	//			StandardTextFieldLen = -1
	//			StandardTextJustification = TxtJustify.Center()
	//				// Text Justification Ignored. Field
	//				// Length Equals Title Line String Length
	//			Formatted Title Line String =
	//				"Hello World!"
	//
	//		Example-4
	//          Title Line String = "Hello World!"
	//			Title Line String Length = 12
	//			StandardTextFieldLen = 2
	//			StandardTextJustification = TxtJustify.Center()
	//				// Justification Ignored because Field
	//				// Length Less Than Title Line String Length.
	//			Formatted Title Line String =
	//				"Hello World!"

	StandardTextJustification TextJustify
	//	The standard text field justification applied to
	//	all Text Title Lines in the 'TitleLines' array.
	//
	//	Type 'TextJustify' is an enumeration which
	//	specifies the justification of the text field
	//	contents string within the text	field length
	//	specified by 'StandardTextFieldLen'.
	//
	//	Text justification can only be evaluated in the
	//	context of a text label, field length and a Text
	//	Justification object of type TextJustify. This is
	//	because text labels with a field length equal to
	//	or less than the length of the text label string
	//	will never use text justification. In these cases,
	//	text justification is completely ignored.
	//
	//	If the field length is greater than the length of
	//	the text label string, text justification must be
	//	equal to one of these three valid values:
	//
	//	    TextJustify(0).Left()
	//	    TextJustify(0).Right()
	//	    TextJustify(0).Center()
	//
	//	Users can also specify the abbreviated text
	//	justification enumeration syntax as follows:
	//
	//	    TxtJustify.Left()
	//	    TxtJustify.Right()
	//	    TxtJustify.Center()
	//
	//	Text Justification Examples
	//
	//		Example-1
	//          Title Line String = "Hello World!"
	//			Title Line String Length = 12
	//			StandardTextFieldLen = 18
	//			StandardTextJustification = TxtJustify.Center()
	//			Formatted Title Line String =
	//				"   Hello World!   "
	//
	//		Example-2
	//          Title Line String = "Hello World!"
	//			Title Line String Length = 12
	//			StandardTextFieldLen = 18
	//			StandardTextJustification = TxtJustify.Left()
	//			Formatted Title Line String =
	//				"Hello World!      "
	//
	//		Example-3
	//          Title Line String = "Hello World!"
	//			Title Line String Length = 12
	//			StandardTextFieldLen = -1
	//			StandardTextJustification = TxtJustify.Center()
	//				// Text Justification Ignored. Field
	//				// Length Equals Title Line String Length
	//			Formatted Title Line String =
	//				"Hello World!"
	//
	//		Example-4
	//          Title Line String = "Hello World!"
	//			Title Line String Length = 12
	//			StandardTextFieldLen = 2
	//			StandardTextJustification = TxtJustify.Center()
	//				// Justification Ignored because Field
	//				// Length Less Than Title Line String Length.
	//			Formatted Title Line String =
	//				"Hello World!"

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
	//
	//	If this value is less than one (1), no
	//	solid line will be created.

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
	//
	//	If this value is less than one (1), no
	//	solid line will be created.

	NumTrailingBlankLines int
	//	The number of blank lines or 'new lines'
	//	inserted after the Trailing Solid Line.

	lock *sync.Mutex
}

//	AddTitleLineDateTimeFmtDto
//
//	Adds a Date Time text title line to the text
//	title lines array contained in the current instance
//	of TextLineTitleMarqueeDto.
//
//	Be advised that the left and right margins for this
//	Date Time Title field will automatically use the
//	TextLineTitleMarqueeDto StandardTitleLeftMargin and
//	StandardTitleRightMargin.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	If any data field contained in the current instance
//	of TextLineTitleMarqueeDto is judged to be invalid,
//	an error will be returned.
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
// # Input Parameters
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
//		Field Length Examples
//
//			Example-1
//		  	     Title Line String = "Hello World!"
//				Title Line String Length = 12
//				StandardTextFieldLen = 18
//				StandardTextJustification = TxtJustify.Center()
//				Formatted Title Line String =
//					"   Hello World!   "
//
//			Example-2
//		  	     Title Line String = "Hello World!"
//				Title Line String Length = 12
//				StandardTextFieldLen = 18
//				StandardTextJustification = TxtJustify.Left()
//				Formatted Title Line String =
//					"Hello World!      "
//
//			Example-3
//		  	     Title Line String = "Hello World!"
//				Title Line String Length = 12
//				StandardTextFieldLen = -1
//				StandardTextJustification = TxtJustify.Center()
//					// Text Justification Ignored. Field
//					// Length Equals Title Line String Length
//				Formatted Title Line String =
//					"Hello World!"
//
//			Example-4
//		  	     Title Line String = "Hello World!"
//				Title Line String Length = 12
//				StandardTextFieldLen = 2
//				StandardTextJustification = TxtJustify.Center()
//					// Justification Ignored because Field
//					// Length Less Than Title Line String Length.
//				Formatted Title Line String =
//					"Hello World!"
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
//		The format operations are also documented at:
//			https://pkg.go.dev/time#Time.Format
//			https://www.golanglearn.com/golang-tutorials/go-date-and-time-formatting/
//			https://gosamples.dev/date-time-format-cheatsheet/
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
//		Text Justification Examples
//
//			Example-1
//		  	     Title Line String = "Hello World!"
//				Title Line String Length = 12
//				StandardTextFieldLen = 18
//				StandardTextJustification = TxtJustify.Center()
//				Formatted Title Line String =
//					"   Hello World!   "
//
//			Example-2
//		  	     Title Line String = "Hello World!"
//				Title Line String Length = 12
//				StandardTextFieldLen = 18
//				StandardTextJustification = TxtJustify.Left()
//				Formatted Title Line String =
//					"Hello World!      "
//
//			Example-3
//		  	     Title Line String = "Hello World!"
//				Title Line String Length = 12
//				StandardTextFieldLen = -1
//				StandardTextJustification = TxtJustify.Center()
//					// Text Justification Ignored. Field
//					// Length Equals Title Line String Length
//				Formatted Title Line String =
//					"Hello World!"
//
//			Example-4
//		  	     Title Line String = "Hello World!"
//				Title Line String Length = 12
//				StandardTextFieldLen = 2
//				StandardTextJustification = TxtJustify.Center()
//					// Justification Ignored because Field
//					// Length Less Than Title Line String Length.
//				Formatted Title Line String =
//					"Hello World!"
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
func (txtLineTitleMarqueeDto *TextLineTitleMarqueeDto) AddTitleLineDateTimeFmtDto(
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
		"TextLineTitleMarqueeDto.AddTitleLineLabel()",
		"")

	if err != nil {
		return err
	}

	_,
		err = new(textLineTitleMarqueeDtoAtom).
		testValidityOfTitleMarqueeDto(
			txtLineTitleMarqueeDto,
			ePrefix.XCpy(
				"txtLineTitleMarqueeDto"))

	if err != nil {
		return err
	}

	if len(dateTimeFormat) == 0 {

		dateTimeFormat = new(textSpecificationMolecule).
			getDefaultDateTimeFormat()
	}

	fieldLen,
		err = new(textLineTitleMarqueeDtoAtom).
		calcTextFieldLen(
			txtLineTitleMarqueeDto,
			fieldLen,
			ePrefix)

	if err != nil {
		return err
	}

	txtFieldFmtDtoDate := TextFieldFormatDtoDate{
		LeftMarginStr:       txtLineTitleMarqueeDto.StandardTitleLeftMargin,
		FieldDateTime:       dateTime,
		FieldDateTimeFormat: dateTimeFormat,
		FieldLength:         fieldLen,
		FieldJustify:        textJustification,
		RightMarginStr:      txtLineTitleMarqueeDto.StandardTitleRightMargin,
	}

	err = txtFieldFmtDtoDate.IsValidInstanceError(
		ePrefix.XCpy(
			"txtFieldFmtDtoDate"))

	if err != nil {
		return err
	}

	return new(textLineTitleMarqueeDtoNanobot).
		addTitleLineFmtDtos(
			txtLineTitleMarqueeDto,
			ePrefix.XCpy(
				"txtFieldFmtDtoDate"),
			&txtFieldFmtDtoDate)
}

// AddTitleLineDateTimeStr
//
// Receives a time.Time instance and converts it to a
// string using the date/time format input parameter
// 'dateTimeFormat'. The string generated from this
// format parameter is used to construct a title line
// using the left margin, right margin, field length and
// text justification specifications contained in the
// current instance of TextLineTitleMarqueeDto.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	If any data field contained in the current instance
//	of TextLineTitleMarqueeDto is judged to be invalid,
//	an error will be returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//		The format operations are also documented at:
//			https://pkg.go.dev/time#Time.Format
//			https://www.golanglearn.com/golang-tutorials/go-date-and-time-formatting/
//			https://gosamples.dev/date-time-format-cheatsheet/
//
//		If this parameter is submitted as an empty string,
//		parameter 'dateTimeFormat' will be assigned a
//		default value of:
//			"2006-01-02 15:04:05.000000000 -0700 MST".
//
//		Example Formats:
//
//		 Example 1:
//		 dateTimeFormat =
//		   "2006-01-02 15:04:05.000000000 -0700 MST"
//
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
//		 Example 3:
//		  dateTimeFormat =
//		   "Monday 2006-01-02 15:04:05.000000000 -0700 MST"
//
//		 Result =
//		   "Monday 2021-10-21 14:19:03.000000000 -0500 CDT"
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
func (txtLineTitleMarqueeDto *TextLineTitleMarqueeDto) AddTitleLineDateTimeStr(
	dateTime time.Time,
	dateTimeFormat string,
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
		"TextLineTitleMarqueeDto.AddTitleLineLabel()",
		"")

	if err != nil {
		return err
	}

	_,
		err = new(textLineTitleMarqueeDtoAtom).
		testValidityOfTitleMarqueeDto(
			txtLineTitleMarqueeDto,
			ePrefix.XCpy("txtLineTitleMarqueeDto"))

	if err != nil {
		return err
	}

	if len(dateTimeFormat) == 0 {

		dateTimeFormat = new(textSpecificationMolecule).
			getDefaultDateTimeFormat()
	}

	txtFieldFmtDtoDate := TextFieldFormatDtoDate{
		LeftMarginStr:       txtLineTitleMarqueeDto.StandardTitleLeftMargin,
		FieldDateTime:       dateTime,
		FieldDateTimeFormat: dateTimeFormat,
		FieldLength:         txtLineTitleMarqueeDto.StandardTextFieldLen,
		FieldJustify:        txtLineTitleMarqueeDto.StandardTextJustification,
		RightMarginStr:      txtLineTitleMarqueeDto.StandardTitleRightMargin,
	}

	err = txtFieldFmtDtoDate.IsValidInstanceError(
		ePrefix.XCpy(
			"txtFieldFmtDtoDate"))

	if err != nil {
		return err
	}

	return new(textLineTitleMarqueeDtoNanobot).
		addTitleLineFmtDtos(
			txtLineTitleMarqueeDto,
			ePrefix.XCpy(
				"txtFieldFmtDtoDate"),
			&txtFieldFmtDtoDate)

}

//	AddTitleLineLabel
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
// # IMPORTANT
//
//	If any data field contained in the current instance
//	of TextLineTitleMarqueeDto is judged to be invalid,
//	an error will be returned.
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
func (txtLineTitleMarqueeDto *TextLineTitleMarqueeDto) AddTitleLineLabel(
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
			"AddTitleLineLabel()",
		"")

	if err != nil {
		return err
	}

	if len(textLabel) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'textLabel' is INVALID!\n"+
			"'textLabel' is an empty string with zero text characters!\n",
			ePrefix.String())

		return err

	}

	_,
		err = new(textLineTitleMarqueeDtoAtom).
		testValidityOfTitleMarqueeDto(
			txtLineTitleMarqueeDto,
			ePrefix.XCpy(
				"txtLineTitleMarqueeDto"))

	if err != nil {
		return err
	}

	fieldLen,
		err = new(textLineTitleMarqueeDtoAtom).
		calcTextFieldLen(
			txtLineTitleMarqueeDto,
			fieldLen,
			ePrefix)

	if err != nil {
		return err
	}

	txtLabel := TextFieldFormatDtoLabel{
		LeftMarginStr:  txtLineTitleMarqueeDto.StandardTitleLeftMargin,
		FieldContents:  textLabel,
		FieldLength:    fieldLen,
		FieldJustify:   textJustification,
		RightMarginStr: txtLineTitleMarqueeDto.StandardTitleRightMargin,
	}

	err = txtLabel.IsValidInstanceError(
		ePrefix.XCpy(
			"txtLineTitleMarqueeDto"))

	if err != nil {
		return err
	}

	return new(textLineTitleMarqueeDtoNanobot).
		addTitleLineFmtDtos(
			txtLineTitleMarqueeDto,
			ePrefix.XCpy("txtLineTitleMarqueeDto"),
			&txtLabel)
}

// AddTitleLineFmtDtos
//
// Receives one or more ITextFieldFormatDto objects
// passed as a variadic parameter.
//
// Each ITextFieldFormatDto passed to this method will
// constitute a separate text field used to create a
// single text title line which will then be added to
// the Text Title Line Collection for the current
// instance of TextLineTitleMarqueeDto.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	Only one (1) title line will be created and added
//	to the Text Title Line Collections. The
//	ITextFieldFormatDto objects passed to this method
//	represent one or more text fields which will be
//	consolidated to form one line of text.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	If any data field contained in the current instance
//	of TextLineTitleMarqueeDto is judged to be invalid,
//	an error will be returned.
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
//	textFieldColumns			...ITextFieldFormatDto
//
//		This variadic parameter is used to pass one or
//		more instances of objects implementing the
//		ITextFieldFormatDto interface.
//
//		These ITextFieldFormatDto object contains all the
//		text field content and formatting specifications
//		necessary to format one or more text fields in
//		a standard line of text. This method will use
//		these formatting specifications to create a
//		text title line and add it to the text title
//		line collection maintained by the current
//		instance of TextLineTitleMarqueeDto.
//
//		Examples of concrete types implementing the
//		ITextFieldFormatDto interface are:
//
//			TextFieldFormatDtoBigFloat
//			TextFieldFormatDtoDate
//			TextFieldFormatDtoLabel
//			TextFieldFormatDtoFiller
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
func (txtLineTitleMarqueeDto *TextLineTitleMarqueeDto) AddTitleLineFmtDtos(
	errorPrefix interface{},
	textFieldColumns ...ITextFieldFormatDto) error {

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
			"AddTitleLineFmtDtos()",
		"")

	if err != nil {
		return err
	}

	_,
		err = new(textLineTitleMarqueeDtoAtom).
		testValidityOfTitleMarqueeDto(
			txtLineTitleMarqueeDto,
			ePrefix.XCpy(
				"txtLineTitleMarqueeDto"))

	if err != nil {
		return err
	}

	return new(textLineTitleMarqueeDtoNanobot).
		addTitleLineFmtDtos(
			txtLineTitleMarqueeDto,
			ePrefix.XCpy(
				"txtLineTitleMarqueeDto"),
			textFieldColumns...)

}

//	AddTitleLineITextFields
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
// # IMPORTANT
//
//	If any data field contained in the current instance
//	of TextLineTitleMarqueeDto is judged to be invalid,
//	an error will be returned.
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
func (txtLineTitleMarqueeDto *TextLineTitleMarqueeDto) AddTitleLineITextFields(
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
			"AddTitleLineITextFields()",
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

	_,
		err = new(textLineTitleMarqueeDtoAtom).
		testValidityOfTitleMarqueeDto(
			txtLineTitleMarqueeDto,
			ePrefix.XCpy(
				"txtLineTitleMarqueeDto"))

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
// # IMPORTANT
//
//	If any data field contained in the current instance
//	of TextLineTitleMarqueeDto is judged to be invalid,
//	an error will be returned.
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
		err = new(textLineTitleMarqueeDtoAtom).
		testValidityOfTitleMarqueeDto(
			txtLineTitleMarqueeDto,
			ePrefix.XCpy(
				"txtLineTitleMarqueeDto"))

	if err != nil {
		return err
	}

	_,
		err = new(textLineSpecStandardLineAtom).
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

// AddTitleLineStrArrayDto
//
// Receives one or more strings and incorporates each one
// as a separate title line the title Marque. Each string
// is therefore formatted and saved as a separate title
// line in the internal Title Lines collection maintained
// by the current instance of TextLineTitleMarqueeDto:
//
//	TextLineTitleMarqueeDto.TitleLines
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	If any data field contained in the current instance
//	of TextLineTitleMarqueeDto is judged to be invalid,
//	an error will be returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	titleLineStrArrayDto		StringArrayDto
//
//		An instance of StringArrayDto which encapsulates
//		an array strings which will be converted to
//		Marquee title lines and stored in the Title Line
//		Collection maintained by the current instance of
//		TextLineTitleMarqueeDto. This internal title lines
//		collection is designated as:
//			TextLineTitleMarqueeDto.TitleLines
//
//		The field length and text justification for each
//		string converted to a title line will be taken
//		from internal member variables:
//			TextLineTitleMarqueeDto.StandardTextFieldLen
//			TextLineTitleMarqueeDto.StandardTextJustification
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
func (txtLineTitleMarqueeDto *TextLineTitleMarqueeDto) AddTitleLineStrArrayDto(
	titleLineStrArrayDto StringArrayDto,
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
			"AddTitleLineStrArrayDto()",
		"")

	if err != nil {
		return err
	}

	_,
		err = new(textLineTitleMarqueeDtoAtom).
		testValidityOfTitleMarqueeDto(
			txtLineTitleMarqueeDto,
			ePrefix.XCpy(
				"txtLineTitleMarqueeDto"))

	if err != nil {
		return err
	}

	fieldFmtDto := TextFieldFormatDtoLabel{
		LeftMarginStr:  txtLineTitleMarqueeDto.StandardTitleLeftMargin,
		FieldContents:  nil,
		FieldLength:    txtLineTitleMarqueeDto.StandardTextFieldLen,
		FieldJustify:   txtLineTitleMarqueeDto.StandardTextJustification,
		RightMarginStr: txtLineTitleMarqueeDto.StandardTitleRightMargin,
	}
	txtTitleMarqueeDtoAtom := textLineTitleMarqueeDtoNanobot{}

	for idx, titleStr := range titleLineStrArrayDto.StrArray {

		fieldFmtDto.FieldContents = titleStr

		err = txtTitleMarqueeDtoAtom.addTitleLineFmtDtos(
			txtLineTitleMarqueeDto,
			ePrefix.XCpy(
				fmt.Sprintf("titleLineStrings[%v]",
					idx)),
			&fieldFmtDto)

		if err != nil {
			return err
		}

	}

	return err
}

// AddTitleLineStrings
//
// Adds one or more strings to the Title Line Collection
// maintained by the current instance of
// TextLineTitleMarqueeDto.
//
// A separate title line will be created for each string
// passed through input parameter 'titleLineStrings'.
//
// This method applies the Standard Left Margin, Right
// Margin, Field Length and Field Justification values
// previously configured for this instance of
// TextLineTitleMarqueeDto.
//
// Multiple strings may be passed through the variadic
// input paramter, 'titleLineStrings'. Variadic
// parameters accept a variable number of arguments.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	If any data field contained in the current instance
//	of TextLineTitleMarqueeDto is judged to be invalid,
//	an error will be returned.
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
//	titleLineStrings			...string
//
//		One or more strings containing a title line to be
//		displayed in the title marquee generated by the
//		current instance of TextLineTitleMarqueeDto.
//
//		This variadic parameter accepts a variable number
//		of arguments.
//
//		This method will apply the Standard Left Margin,
//		Right Margin, Field Length and Field
//		Justification previously configured for this
//		instance of TextLineTitleMarqueeDto. Using
//		these standard formatting parameters, the
//		'titleLineStr' will be converted into a
//		formatted text line.
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
func (txtLineTitleMarqueeDto *TextLineTitleMarqueeDto) AddTitleLineStrings(
	errorPrefix interface{},
	titleLineStrings ...string) error {

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
			"AddTitleLineStr()",
		"")

	if err != nil {
		return err
	}

	_,
		err = new(textLineTitleMarqueeDtoAtom).
		testValidityOfTitleMarqueeDto(
			txtLineTitleMarqueeDto,
			ePrefix.XCpy(
				"txtLineTitleMarqueeDto"))

	if err != nil {

		return err
	}

	itemCnt := 0

	txtLabelFmtDto := TextFieldFormatDtoLabel{
		LeftMarginStr:  txtLineTitleMarqueeDto.StandardTitleLeftMargin,
		FieldContents:  nil,
		FieldLength:    txtLineTitleMarqueeDto.StandardTextFieldLen,
		FieldJustify:   txtLineTitleMarqueeDto.StandardTextJustification,
		RightMarginStr: txtLineTitleMarqueeDto.StandardTitleRightMargin,
	}

	txtLineMarqueeNanobot := textLineTitleMarqueeDtoNanobot{}

	for idx, titleStr := range titleLineStrings {

		txtLabelFmtDto.FieldContents = titleStr

		err = txtLabelFmtDto.IsValidInstanceError(
			ePrefix.XCpy(
				fmt.Sprintf("titleStr[%v]",
					idx)))

		if err != nil {

			return err
		}

		err = txtLineMarqueeNanobot.
			addTitleLineFmtDtos(
				txtLineTitleMarqueeDto,
				ePrefix.XCpy(
					fmt.Sprintf("titleStr[%v]",
						idx)),
				&txtLabelFmtDto)

		if err != nil {

			return err
		}

		itemCnt++
	}

	if itemCnt == 0 {
		err = fmt.Errorf("\n%v\n"+
			"Error: Input parameter 'titleLineStrings' is INVALID!\n"+
			"'titleLineStrings' is empty and contains zero strings.\n",
			ePrefix.String())
	}

	return err
}

// CopyIn
//
// Copies all the data fields from an incoming instance
// of TextLineTitleMarqueeDto ('incomingTitleMarqueeDto')
// to the corresponding data fields of the current
// TextLineTitleMarqueeDto instance
// ('txtLineTitleMarqueeDto').
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and overwrite all
//	pre-existing data values contained within the
//	current instance of TextLineTitleMarqueeDto
//	('txtLineTitleMarqueeDto').
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	incomingTitleMarqueeDto		*TextLineTitleMarqueeDto
//
//		A pointer to an instance of
//		TextLineTitleMarqueeDto.
//
//		All the internal data field values in this
//		instance will be copied to corresponding data
//		fields of the current TextLineTitleMarqueeDto
//		instance.
//
//		The data fields contained in
//		'incomingTitleMarqueeDto' will NOT be
//		changed or modified.
//
//		If 'incomingTitleMarqueeDto' contains
//		invalid data values, an error will be returned.
//
//	errorPrefix						interface{}
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
func (txtLineTitleMarqueeDto *TextLineTitleMarqueeDto) CopyIn(
	incomingTitleMarqueeDto *TextLineTitleMarqueeDto,
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
			"CopyIn()",
		"")

	if err != nil {
		return err
	}

	return new(textLineTitleMarqueeDtoNanobot).copy(
		txtLineTitleMarqueeDto,
		incomingTitleMarqueeDto,
		ePrefix.XCpy("txtLineTitleMarqueeDto"))
}

// CopyOut
//
// Returns a deep copy of the current
// TextLineTitleMarqueeDto instance.
//
// If the current TextLineTitleMarqueeDto instance
// contains invalid member variable data values, this
// method will return an error.
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
//	TextLineTitleMarqueeDto
//
//		If this method completes successfully and no
//		errors are encountered, this parameter will
//		return a deep copy of the current
//		TextLineTitleMarqueeDto instance.
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
func (txtLineTitleMarqueeDto *TextLineTitleMarqueeDto) CopyOut(
	errorPrefix interface{}) (
	TextLineTitleMarqueeDto,
	error) {

	if txtLineTitleMarqueeDto.lock == nil {
		txtLineTitleMarqueeDto.lock = new(sync.Mutex)
	}

	txtLineTitleMarqueeDto.lock.Lock()

	defer txtLineTitleMarqueeDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	var newTxtLineMarqueeDto TextLineTitleMarqueeDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineTitleMarqueeDto."+
			"CopyIn()",
		"")

	if err != nil {
		return newTxtLineMarqueeDto, err
	}

	err = new(textLineTitleMarqueeDtoNanobot).copy(
		&newTxtLineMarqueeDto,
		txtLineTitleMarqueeDto,
		ePrefix.XCpy("txtLineTitleMarqueeDto"))

	return newTxtLineMarqueeDto, err
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

	new(textLineTitleMarqueeDtoAtom).empty(
		txtLineTitleMarqueeDto)

	txtLineTitleMarqueeDto.lock.Unlock()

	txtLineTitleMarqueeDto.lock = nil

	return
}

// GetFormattedText
//
//	Returns lines of text comprising the entire Title
//	Marquee. This text is generated from the
//	specifications contained in the current instance of
//	TextLineTitleMarqueeDto.
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
//		This parameter returns the formatted text lines
//		generated from the current instance of
//		TextLineSpecTitleMarquee.
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
func (txtLineTitleMarqueeDto *TextLineTitleMarqueeDto) GetFormattedText(
	errorPrefix interface{}) (
	string,
	error) {

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
			"GetFormattedText()",
		"")

	if err != nil {
		return "", err
	}

	_,
		err = new(textLineTitleMarqueeDtoAtom).
		testValidityOfTitleMarqueeDto(
			txtLineTitleMarqueeDto,
			ePrefix.XCpy(
				"txtLineTitleMarqueeDto"))

	if err != nil {
		return "", err
	}

	var strBuilder strings.Builder

	err = new(textLineTitleMarqueeDtoNanobot).
		getFormattedTitleMarquee(
			txtLineTitleMarqueeDto,
			&strBuilder,
			ePrefix.XCpy(
				"txtLineTitleMarqueeDto"))

	if err != nil {
		return "", err
	}

	return strBuilder.String(), err
}

// GetTitleLines
//
// Returns a copy of the internal Title Lines Collection
// from the current instance of TextLineTitleMarqueeDto.
//
//	TextLineTitleMarqueeDto.TitleLines
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
func (txtLineTitleMarqueeDto *TextLineTitleMarqueeDto) GetTitleLines(
	errorPrefix interface{}) (
	numberOfTitleLines int,
	titleLinesCollection TextLineSpecLinesCollection,
	err error) {

	if txtLineTitleMarqueeDto.lock == nil {
		txtLineTitleMarqueeDto.lock = new(sync.Mutex)
	}

	txtLineTitleMarqueeDto.lock.Lock()

	defer txtLineTitleMarqueeDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	numberOfTitleLines = 0

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineTitleMarqueeDto."+
			"GetTitleLines()",
		"")

	if err != nil {
		return numberOfTitleLines, titleLinesCollection, err
	}

	numberOfTitleLines =
		txtLineTitleMarqueeDto.TitleLines.GetNumberOfTextLines()

	if numberOfTitleLines == 0 {

		return numberOfTitleLines, titleLinesCollection, err

	}

	titleLinesCollection,
		err = txtLineTitleMarqueeDto.TitleLines.CopyOut(
		ePrefix.XCpy(
			"titleLinesCollection<-" +
				"txtLineTitleMarqueeDto.TitleLines"))

	return numberOfTitleLines, titleLinesCollection, err
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
		_ = new(textLineTitleMarqueeDtoAtom).
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
		err = new(textLineTitleMarqueeDtoAtom).
		testValidityOfTitleMarqueeDto(
			txtLineTitleMarqueeDto,
			ePrefix.XCpy(
				"txtLineTitleMarqueeDto"))

	return err
}

// NewBasicTitleMarqueeDto
//
// Creates and returns a fully populated instance of
// TextLineTitleMarqueeDto using a basic Title Marquee
// profile incorporating default text line
// specifications.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	If the 'titleLines' paramter is submitted as an
//	empty or nil value, no error will be returned.
//
//	However, the user will then be responsible for
//	populating the title lines array using the 'Add'
//	methods documented above.
//
// ----------------------------------------------------------------
//
// # Default Values
//
//	The returned instance of TextLineTitleMarqueeDto will
//	be constructed as a basic Text Title Marquee
//	incorporating the following default values.
//
//	StandardTitleLeftMargin = standardTitleLeftMargin
//
//	StandardTitleRightMargin = standardTitleRightMargin
//
//	StandardMaxLineLen = standardMaxLineLen
//	StandardTextFieldLen =
//		standardMaxLineLen -
//		1 -
//		len(standardTitleLeftMargin) -
//		len(standardTitleRightMargin)
//
//	StandardTextJustification = TxtJustify.Center()
//
//	NumLeadingBlankLines = 1
//
//	LeadingSolidLineChar = solidLineChar
//
//	NumTopTitleBlankLines = 1
//
//	NumBottomTitleBlankLines = 1
//
//	TrailingSolidLineChar = solidLineChar
//
//	NumTrailingBlankLines = 1
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	standardTitleLeftMargin		string
//
//		The standard left margin characters applied
//		to all solid lines and Text Title Lines in the
//		Title Lines array
//		(TextLineTitleMarqueeDto.TitleLines).
//
//	standardTitleRightMargin	string
//
//		The standard right margin characters applied
//		to all solid lines and Text Title Lines in the
//		Title Lines array
//		(TextLineSpecTitleMarquee.titleLines).
//
//	standardMaxLineLen			int
//
//		The maximum number of characters allowed on
//		a text title line. This maximum limit will be
//		applied to the length of all text lines generated
//		by the returned instance of
//		TextLineSpecTitleMarquee.
//
//		In addition, this value will serve as the
//		standard field length for all solid lines and
//		text lines formatted for the Title Marquee.
//
//	solidLineChar				string
//
//		The character used to create the Leading
//		Solid Line displayed above the Title
//		Lines.
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
//	titleLines					...interface{}
//
//		This is a variadic parameter which can accept a
//		variable number of arguments.
//
//		Zero, one or more objects may be passed through
//		this input parameter.
//
//		If an empty string is passed through this
//		parameter, an error will be returned.
//
//		If one or more objects are passed through this
//		parameter, they must be convertable to one of the
//		following types:
//
//			time.Time (Converted using default format)
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
//			TextFieldFormatDtoBigFloat - Formats big.Float numbers
//
//		If the 'emptyIFace' object is not convertible to
//		one of the supported types, an error will be
//		returned.
//
//		Be advised. If this paramter is submitted as an
//		empty or nil value, no error will be returned.
//		However, the user will then be responsible for
//		populating the title lines array using the 'Add'
//		methods documented above.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	TextLineSpecTitleMarquee
//
//		If this method completes successfully, a new
//		fully configured instance of
//		TextLineTitleMarqueeDto will be returned.
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
func (txtLineTitleMarqueeDto *TextLineTitleMarqueeDto) NewBasicTitleMarqueeDto(
	standardSolidLineLeftMargin string,
	standardSolidLineRightMargin string,
	standardTitleLeftMargin string,
	standardTitleRightMargin string,
	standardMaxLineLen int,
	solidLineChar string,
	errorPrefix interface{},
	titleLines ...interface{}) (
	TextLineTitleMarqueeDto,
	error) {

	if txtLineTitleMarqueeDto.lock == nil {
		txtLineTitleMarqueeDto.lock = new(sync.Mutex)
	}

	txtLineTitleMarqueeDto.lock.Lock()

	defer txtLineTitleMarqueeDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	var newTextLineTitleMarqueeDto TextLineTitleMarqueeDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineTitleMarqueeDto."+
			"NewBasicTitleMarqueeDto()",
		"")

	if err != nil {
		return newTextLineTitleMarqueeDto, err
	}

	err = new(textLineTitleMarqueeDtoNanobot).
		configureBasicTitleMarqueeDto(
			&newTextLineTitleMarqueeDto,
			standardSolidLineLeftMargin,
			standardSolidLineRightMargin,
			standardTitleLeftMargin,
			standardTitleRightMargin,
			standardMaxLineLen,
			solidLineChar,
			ePrefix.XCpy(
				"newTextLineTitleMarqueeDto"),
			titleLines...)

	if err != nil {
		return newTextLineTitleMarqueeDto, err
	}

	_,
		err = new(textLineTitleMarqueeDtoAtom).
		testValidityOfTitleMarqueeDto(
			&newTextLineTitleMarqueeDto,
			ePrefix.XCpy(
				"configured newTextLineTitleMarqueeDto"))

	return newTextLineTitleMarqueeDto, err
}

// TextBuilder
//
//	Returns lines of text comprising the entire Title
//	Marquee. This text is generated from the
//	specifications contained in the current instance of
//	TextLineTitleMarqueeDto.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	strBuilder                 *strings.Builder
//
//		An instance of strings.Builder. The line of text
//		produced by the current instance of
//		TextLineTitleMarqueeDto will be written to
//		'sBuilder'.
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
func (txtLineTitleMarqueeDto *TextLineTitleMarqueeDto) TextBuilder(
	strBuilder *strings.Builder,
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
			"TextBuilder()",
		"")

	if err != nil {
		return err
	}

	_,
		err = new(textLineTitleMarqueeDtoAtom).
		testValidityOfTitleMarqueeDto(
			txtLineTitleMarqueeDto,
			ePrefix.XCpy(
				"txtLineTitleMarqueeDto"))

	if err != nil {
		return err
	}

	return new(textLineTitleMarqueeDtoNanobot).
		getFormattedTitleMarquee(
			txtLineTitleMarqueeDto,
			strBuilder,
			ePrefix.XCpy(
				"txtLineTitleMarqueeDto"))

}

//	textLineTitleMarqueeDtoNanobot
//
//	Provides helper methods for type
//	TextLineTitleMarqueeDto

type textLineTitleMarqueeDtoNanobot struct {
	lock *sync.Mutex
}

// copy
//
// Copies all data from a source instance of
// TextLineTitleMarqueeDto to a destination instance of
// TextLineTitleMarqueeDto.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and overwrite all
//	pre-existing data values contained within the
//	TextLineTitleMarqueeDto instance passed as input
//	parameter 'destinationTitleMarqueeDto'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	destinationTitleMarqueeDto		*TextLineTitleMarqueeDto
//
//		A pointer to an instance of
//		TextLineTitleMarqueeDto.
//
//		Data extracted from input parameter
//		'sourceTitleMarqueeDto' will be copied to this
//		input parameter, 'destinationTitleMarqueeDto'.
//
//		'destinationTitleMarqueeDto' is the destination
//		for this copy operation.
//
//		If this method completes successfully, all member
//		data variables encapsulated in
//		'destinationTitleMarqueeDto' will be identical to
//		those contained in input parameter,
//		'sourceTitleMarqueeDto'.
//
//		Be advised that the pre-existing data fields
//		contained within input parameter
//		'destinationTitleMarqueeDto' will be deleted and
//		overwritten.
//
//	sourceTitleMarqueeDto			*TextLineTitleMarqueeDto
//
//		A pointer to an instance of
//		TextLineTitleMarqueeDto.
//
//		All data values in this TextLineTitleMarqueeDto
//		instance will be copied to input parameter
//		'destinationTitleMarqueeDto'.
//
//		'sourceTitleMarqueeDto' is the source of the
//		copy operation.
//
//		The original member variable data values
//		encapsulated within 'sourceTitleMarqueeDto'
//		will remain unchanged and unmodified.
//
//		If 'sourceTitleMarqueeDto' contains invalid
//		member data variables, this method will return
//		an error.
//
//	errPrefDto						*ePref.ErrPrefixDto
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
func (txtTitleMarqueeDtoNanobot *textLineTitleMarqueeDtoNanobot) copy(
	destinationTitleMarqueeDto *TextLineTitleMarqueeDto,
	sourceTitleMarqueeDto *TextLineTitleMarqueeDto,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtTitleMarqueeDtoNanobot.lock == nil {
		txtTitleMarqueeDtoNanobot.lock = new(sync.Mutex)
	}

	txtTitleMarqueeDtoNanobot.lock.Lock()

	defer txtTitleMarqueeDtoNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineTitleMarqueeDtoNanobot."+
			"addTitleLineFmtDtos()",
		"")

	if err != nil {
		return err
	}

	if destinationTitleMarqueeDto == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'destinationTitleMarqueeDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if sourceTitleMarqueeDto == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'sourceTitleMarqueeDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	txtLineMarqueeDtoAtom := textLineTitleMarqueeDtoAtom{}

	_,
		err = txtLineMarqueeDtoAtom.
		testValidityOfTitleMarqueeDto(
			sourceTitleMarqueeDto,
			ePrefix.XCpy(
				"sourceTitleMarqueeDto"))

	if err != nil {
		return err
	}

	txtLineMarqueeDtoAtom.empty(
		destinationTitleMarqueeDto)

	destinationTitleMarqueeDto.StandardSolidLineLeftMargin =
		sourceTitleMarqueeDto.StandardSolidLineLeftMargin

	destinationTitleMarqueeDto.StandardSolidLineRightMargin =
		sourceTitleMarqueeDto.StandardSolidLineRightMargin

	destinationTitleMarqueeDto.StandardTitleLeftMargin =
		sourceTitleMarqueeDto.StandardTitleLeftMargin

	destinationTitleMarqueeDto.StandardTitleRightMargin =
		sourceTitleMarqueeDto.StandardTitleRightMargin

	destinationTitleMarqueeDto.StandardMaxLineLen =
		sourceTitleMarqueeDto.StandardMaxLineLen

	destinationTitleMarqueeDto.StandardTextFieldLen =
		sourceTitleMarqueeDto.StandardTextFieldLen

	destinationTitleMarqueeDto.StandardTextJustification =
		sourceTitleMarqueeDto.StandardTextJustification

	destinationTitleMarqueeDto.NumLeadingBlankLines =
		sourceTitleMarqueeDto.NumLeadingBlankLines

	destinationTitleMarqueeDto.LeadingSolidLineChar =
		sourceTitleMarqueeDto.LeadingSolidLineChar

	destinationTitleMarqueeDto.NumLeadingSolidLines =
		sourceTitleMarqueeDto.NumLeadingSolidLines

	destinationTitleMarqueeDto.NumTopTitleBlankLines =
		sourceTitleMarqueeDto.NumTopTitleBlankLines

	if sourceTitleMarqueeDto.TitleLines.
		GetNumberOfTextLines() > 0 {

		err = destinationTitleMarqueeDto.TitleLines.
			CopyIn(
				&sourceTitleMarqueeDto.TitleLines,
				ePrefix.XCpy(
					"sourceTitleMarqueeDto.TitleLines"))

		if err != nil {
			return err
		}

	}

	destinationTitleMarqueeDto.NumBottomTitleBlankLines =
		sourceTitleMarqueeDto.NumBottomTitleBlankLines

	destinationTitleMarqueeDto.TrailingSolidLineChar =
		sourceTitleMarqueeDto.TrailingSolidLineChar

	destinationTitleMarqueeDto.NumTrailingSolidLines =
		sourceTitleMarqueeDto.NumTrailingSolidLines

	destinationTitleMarqueeDto.NumTrailingBlankLines =
		sourceTitleMarqueeDto.NumTrailingBlankLines

	return err
}

// addTitleLineFmtDtos
//
// Receives one or more ITextFieldFormatDto objects
// passed as a variadic parameter.
//
// Each ITextFieldFormatDto passed to this method will
// constitute a separate text field used to create a
// text title line which will then be added to the
// text title line collection for the current instance
// of TextLineTitleMarqueeDto.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	txtTitleMarqueeDto 			*TextLineTitleMarqueeDto
//
//		A pointer to an instance of
//		TextLineTitleMarqueeDto.
//
//		The Text Line Titles Collection contained in this
//		TextLineTitleMarqueeDto instance will be updated
//		by adding a new title line created from input
//		parameter 'titleLineFmtDtos'.
//
//		If 'txtTitleMarqueeDto' contains invalid data
//		elements, an error will be returned.
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
//	titleLineFmtDtos			...ITextFieldFormatDto
//
//		This variadic parameter is used to pass one or
//		more instances of objects implementing the
//		ITextFieldFormatDto interface.
//
//		These ITextFieldFormatDto object contains all the
//		text field content and formatting specifications
//		necessary to format one or more text fields in
//		a standard line of text. This method will use
//		these formatting specifications to create a
//		text title line and add it to the text title
//		line collection maintained by the current
//		instance of TextLineTitleMarqueeDto.
//
//		Examples of concrete types implementing the
//		ITextFieldFormatDto interface are:
//
//			TextFieldFormatDtoBigFloat
//			TextFieldFormatDtoDate
//			TextFieldFormatDtoLabel
//			TextFieldFormatDtoFiller
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
func (txtTitleMarqueeDtoNanobot *textLineTitleMarqueeDtoNanobot) addTitleLineFmtDtos(
	txtTitleMarqueeDto *TextLineTitleMarqueeDto,
	errPrefDto *ePref.ErrPrefixDto,
	titleLineFmtDtos ...ITextFieldFormatDto) error {

	if txtTitleMarqueeDtoNanobot.lock == nil {
		txtTitleMarqueeDtoNanobot.lock = new(sync.Mutex)
	}

	txtTitleMarqueeDtoNanobot.lock.Lock()

	defer txtTitleMarqueeDtoNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineTitleMarqueeDtoNanobot."+
			"addTitleLineFmtDtos()",
		"")

	if err != nil {
		return err
	}

	if txtTitleMarqueeDto == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtTitleMarqueeDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	_,
		err = new(textLineTitleMarqueeDtoAtom).
		testValidityOfTitleMarqueeDto(
			txtTitleMarqueeDto,
			ePrefix.XCpy(
				"txtTitleMarqueeDto"))

	if err != nil {
		return err
	}

	var stdLine TextLineSpecStandardLine

	stdLine,
		err = TextLineSpecStandardLine{}.NewStdLineColumns(
		"\n",
		false,
		ePrefix.XCpy("stdLine"),
		titleLineFmtDtos...)

	if err != nil {

		return err
	}

	err = stdLine.IsValidInstanceError(
		ePrefix.XCpy(
			"stdLine<-titleLineFmtDtos..."))

	if err != nil {

		return err
	}

	return txtTitleMarqueeDto.
		TitleLines.AddTextLineSpec(
		&stdLine,
		ePrefix.XCpy(
			"stdLine"))
}

//	AddTitleLineLabel
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
//		# Input Parameters
//
//		txtTitleMarqueeDto 			*TextLineTitleMarqueeDto
//
//			A pointer to an instance of
//			TextLineTitleMarqueeDto.
//
//			The Text Line Titles Collection contained in this
//			TextLineTitleMarqueeDto instance will be updated
//			by adding a new title line created from input
//			parameter 'textLabel'.
//
//			If 'txtTitleMarqueeDto' contains invalid data
//			elements, an error will be returned.
//
//		textLabel					string
//
//			The string content to be displayed within the
//			text label.
//
//			If this parameter is submitted as a zero length
//			string, an error will be returned.
//
//		fieldLen					int
//
//			The length of the text field in which the
//			'textLabel' value will be displayed.
//
//			If 'fieldLen' is less than the length of the
//			formatted 'textLabel' string, it will be
//			automatically set equal to the 'textLabel'
//			string length.
//
//			To automatically set the value of 'fieldLen' to
//			the Maximum Available Text Field Length for this
//			Text Line Title Marquee instance, set this
//			parameter to zero (0).
//
//			The Maximum Available Text Field Length is
//			calculated as follows:
//
//			Maximum Available Text Field Length =
//				TextLineTitleMarqueeDto.StandardMaxLineLen -
//				1 -
//				len(TextLineTitleMarqueeDto.StandardTitleLeftMargin) -
//				len(TextLineTitleMarqueeDto.StandardTitleRightMargin)
//
//			Standard Text Field Length for all text lines is
//			defined at initialization by internal member
//			variable:
//
//				TextLineTitleMarqueeDto.StandardTextFieldLen
//
//			Field Length Values And Defaults
//			--------------------------------
//
//			fieldLen > Maximum Available Text Field Length
//				fieldLen = Maximum Available Text Field Length
//
//			fieldLen >= 1 &&
//				fieldLen < len(txtLabel)
//					fieldLen = textLabel string length
//
//			fieldLen = 0
//				fieldLen = Maximum Available Text Field Length
//
//			fieldLen = -1
//				fieldLen = textLabel string length
//
//			fieldLen = -2
//				fieldLen =
//				TextLineTitleMarqueeDto.StandardTextFieldLen
//
//			fieldLen < -2
//				Return value = Error
//
//			If 'fieldLen' is greater than the length of the
//			'textLabel' string, 'textLabel' will be
//			positioned within a text field with a length equal
//			to 'fieldLen'. In this case, the position of the
//			'textLabel' string within the text field will be
//			controlled by the text justification value
//			contained in parameter, 'textJustification'.
//
//		textJustification			TextJustify
//
//			An enumeration which specifies the justification
//			of the 'textLabel' string within a text field.
//
//			The text field length is taken from input
//			parameter 'fieldLen'.
//
//			Text justification can only be evaluated in the
//			context of a text label, field length and a Text
//			Justification object of type TextJustify. This is
//			because text labels with a field length equal to
//			or less than the length of the text label will
//			never use text justification. In these cases,
//			text justification is completely ignored.
//
//
//			If the field length ('fieldLen') is greater than
//			the length of the 'textLabel' string, text
//			justification must be equal to one of these three
//			valid values:
//
//		         TextJustify(0).Left()
//		         TextJustify(0).Right()
//		         TextJustify(0).Center()
//
//			Users can also use the abbreviated text
//			justification enumeration syntax as follows:
//
//		         TxtJustify.Left()
//		         TxtJustify.Right()
//		         TxtJustify.Center()
//
//			Text Justification Examples
//
//				Example-1
//	 	        Title Line String = "Hello World!"
//					Title Line String Length = 12
//					StandardTextFieldLen = 18
//					StandardTextJustification = TxtJustify.Center()
//					Formatted Title Line String =
//						"   Hello World!   "
//
//				Example-2
//	 	        Title Line String = "Hello World!"
//					Title Line String Length = 12
//					StandardTextFieldLen = 18
//					StandardTextJustification = TxtJustify.Left()
//					Formatted Title Line String =
//						"Hello World!      "
//
//				Example-3
//	 	        Title Line String = "Hello World!"
//					Title Line String Length = 12
//					StandardTextFieldLen = -1
//					StandardTextJustification = TxtJustify.Center()
//						// Text Justification Ignored. Field
//						// Length Equals Title Line String Length
//					Formatted Title Line String =
//						"Hello World!"
//
//				Example-4
//	 	        Title Line String = "Hello World!"
//					Title Line String Length = 12
//					StandardTextFieldLen = 2
//					StandardTextJustification = TxtJustify.Center()
//						// Justification Ignored because Field
//						// Length Less Than Title Line String Length.
//					Formatted Title Line String =
//						"Hello World!"
//
//		errPrefDto					*ePref.ErrPrefixDto
//
//			This object encapsulates an error prefix string
//			which is included in all returned error
//			messages. Usually, it contains the name of the
//			calling method or methods listed as a function
//			chain.
//
//			If no error prefix information is needed, set
//			this parameter to 'nil'.
//
//			Type ErrPrefixDto is included in the 'errpref'
//			software package:
//				"github.com/MikeAustin71/errpref".
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
func (txtTitleMarqueeDtoNanobot *textLineTitleMarqueeDtoNanobot) addTitleLineLabel(
	txtTitleMarqueeDto *TextLineTitleMarqueeDto,
	textLabel string,
	fieldLen int,
	textJustification TextJustify,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtTitleMarqueeDtoNanobot.lock == nil {
		txtTitleMarqueeDtoNanobot.lock = new(sync.Mutex)
	}

	txtTitleMarqueeDtoNanobot.lock.Lock()

	defer txtTitleMarqueeDtoNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineTitleMarqueeDtoNanobot."+
			"addTitleLineLabel()",
		"")

	if err != nil {
		return err
	}

	if txtTitleMarqueeDto == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtTitleMarqueeDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	_,
		err = new(textLineTitleMarqueeDtoAtom).
		testValidityOfTitleMarqueeDto(
			txtTitleMarqueeDto,
			ePrefix.XCpy(
				"txtTitleMarqueeDto"))

	if err != nil {
		return err
	}

	if len(textLabel) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'textLabel' is INVALID!\n"+
			"'textLabel' is an empty string with zero text characters!\n",
			ePrefix.String())

		return err

	}

	stdLine := TextLineSpecStandardLine{}.New()

	fieldLen,
		err = new(textLineTitleMarqueeDtoAtom).
		calcTextFieldLen(
			txtTitleMarqueeDto,
			fieldLen,
			ePrefix)

	if err != nil {
		return err
	}

	lenStr := len(txtTitleMarqueeDto.StandardTitleLeftMargin)

	// Left Margin Label
	if lenStr > 0 {

		_,
			err = stdLine.AddTextFieldLabel(
			txtTitleMarqueeDto.StandardTitleLeftMargin,
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
		len(txtTitleMarqueeDto.StandardTitleRightMargin)

	if lenStr > 0 {

		_,
			err = stdLine.AddTextFieldLabel(
			txtTitleMarqueeDto.StandardTitleRightMargin,
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

	err = txtTitleMarqueeDto.TitleLines.AddTextLineSpec(
		&stdLine,
		ePrefix.XCpy(
			"stdLine<-"))

	return err
}

// configureBasicTitleMarqueeDto
//
// Reconfigures an instance of TextLineTitleMarqueeDto as
// a basic Title Marquee using a standard profile
// incorporating default text line specifications.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	txtTitleMarqueeDto 			*TextLineTitleMarqueeDto
//
//		A pointer to an instance of TextLineTitleMarqueeDto.
//
//		All data elements in this instance will be configured
//		with new values based on the following input parameters.
//
//	standardSolidLineLeftMargin		string
//
//		The standard left margin characters applied to
//		all Solid Lines created for this Title Marquee
//
//		If no solid line left margin is required, set
//		this parameter to an empty string.
//
//	standardSolidLineRightMargin	string
//
//		The standard right margin characters applied to
//		all Solid Lines created for this Title Marquee
//
//		If no solid line left margin is required, set
//		this parameter to an empty string.
//
//	standardTitleLeftMargin		string
//
//		The standard left margin characters applied
//		to all Text Title Lines in the Title Lines array,
//		TextLineSpecTitleMarquee.titleLines.
//
//		If no Title Line left margin is required, set
//		this parameter to an empty string.
//
//	standardTitleRightMargin	string
//
//		The standard right margin characters applied
//		to all Text Title Lines in the Title Lines array,
//		TextLineSpecTitleMarquee.titleLines.
//
//		If no Title Line right margin is required, set
//		this parameter to an empty string.
//
//	standardMaxLineLen			int
//
//		The maximum number of characters allowed on
//		a text title line. This maximum limit will be
//		applied to the length of all text lines generated
//		by the returned instance of
//		TextLineSpecTitleMarquee.
//
//		In addition, this value will serve as the
//		standard field length for all solid lines and
//		text lines formatted for the Title Marquee.
//
//	solidLineChar				string
//
//		The character used to create the Leading
//		Solid Line displayed above the Title
//		Lines.
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
//	titleLines					...interface{}
//
//		This is a variadic parameter which can accept a
//		variable number of arguments.
//
//		Zero, one or more objects may be passed through
//		this input parameter.
//
//		If one or more objects are passed through this
//		parameter, they must be convertable to one of the
//		following types:
//
//			time.Time (Converted using default format)
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
//			TextFieldFormatDtoBigFloat - Formats big.Float numbers
//
//		If the 'emptyIFace' object is not convertible to
//		one of the supported types, an error will be
//		returned.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
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
func (txtTitleMarqueeDtoNanobot *textLineTitleMarqueeDtoNanobot) configureBasicTitleMarqueeDto(
	txtTitleMarqueeDto *TextLineTitleMarqueeDto,
	standardSolidLineLeftMargin string,
	standardSolidLineRightMargin string,
	standardTitleLeftMargin string,
	standardTitleRightMargin string,
	standardMaxLineLen int,
	solidLineChar string,
	errPrefDto *ePref.ErrPrefixDto,
	titleLines ...interface{}) error {

	if txtTitleMarqueeDtoNanobot.lock == nil {
		txtTitleMarqueeDtoNanobot.lock = new(sync.Mutex)
	}

	txtTitleMarqueeDtoNanobot.lock.Lock()

	defer txtTitleMarqueeDtoNanobot.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineTitleMarqueeDtoAtom."+
			"configureBasicTitleMarqueeDto()",
		"")

	if err != nil {
		return err
	}

	if txtTitleMarqueeDto == nil {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtTitleMarqueeDto' is invalid!\n"+
			"'txtTitleMarqueeDto' is a nil pointer.\n",
			ePrefix.String())

	}

	standardFieldLen := standardMaxLineLen -
		1 -
		len(standardTitleLeftMargin) -
		len(standardTitleRightMargin)

	txtTitleMarqueeDto.StandardSolidLineLeftMargin =
		standardSolidLineLeftMargin

	txtTitleMarqueeDto.StandardSolidLineRightMargin =
		standardSolidLineRightMargin

	txtTitleMarqueeDto.StandardTitleLeftMargin =
		standardTitleLeftMargin

	txtTitleMarqueeDto.StandardTitleRightMargin =
		standardTitleRightMargin

	txtTitleMarqueeDto.StandardMaxLineLen =
		standardMaxLineLen

	txtTitleMarqueeDto.StandardTextFieldLen =
		standardFieldLen

	txtTitleMarqueeDto.StandardTextJustification =
		TxtJustify.Center()

	txtTitleMarqueeDto.NumLeadingBlankLines = 1

	txtTitleMarqueeDto.LeadingSolidLineChar =
		solidLineChar

	txtTitleMarqueeDto.NumLeadingSolidLines = 1

	txtTitleMarqueeDto.NumTopTitleBlankLines = 1

	txtTitleMarqueeDto.TitleLines =
		TextLineSpecLinesCollection{}

	txtTitleMarqueeDto.NumBottomTitleBlankLines = 1

	txtTitleMarqueeDto.TrailingSolidLineChar = solidLineChar

	txtTitleMarqueeDto.NumTrailingSolidLines = 1

	txtTitleMarqueeDto.NumTrailingBlankLines = 1

	txtLabelDto := TextFieldFormatDtoLabel{
		LeftMarginStr:  standardTitleLeftMargin,
		FieldContents:  nil,
		FieldLength:    standardFieldLen,
		FieldJustify:   TxtJustify.Center(),
		RightMarginStr: standardTitleRightMargin,
	}

	for idx, titleLineObj := range titleLines {

		txtLabelDto.FieldContents = titleLineObj

		err = txtLabelDto.IsValidInstanceError(
			ePrefix.XCpy(
				fmt.Sprintf("txtLabelDto[%v]",
					idx)))

		if err != nil {
			return err
		}

		err = txtTitleMarqueeDto.TitleLines.
			AddStdLineColumns(
				"\n",
				false,
				ePrefix.XCpy(
					fmt.Sprintf(
						"txtTitleMarqueeDto.TitleLines<-"+
							"txtLabelDto[%v]",
						idx)),
				&txtLabelDto)

		if err != nil {
			return err
		}
	}

	return err
}

// getFormattedTitleMarquee
//
// Processes an instance of TextLineTitleMarqueeDto
// passed as an input parameter, and generates the
// formatted lines of text which comprise the Title
// Marquee.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	txtTitleMarqueeDto 			*TextLineTitleMarqueeDto
//
//		A pointer to an instance of
//		TextLineTitleMarqueeDto.
//
//		The Title Marquee specifications contained in
//		this instance of TextLineTitleMarqueeDto will be
//		used to generate the lines of text comprising the
//		Title Marquee.
//
//		If 'txtTitleMarqueeDto' contains invalid data
//		elements, an error will be returned.
//
//	strBuilder                 *strings.Builder
//
//		An instance of strings.Builder. The line of text
//		produced by the current instance of
//		TextLineTitleMarqueeDto will be written to
//		'sBuilder'.
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
func (txtTitleMarqueeDtoNanobot *textLineTitleMarqueeDtoNanobot) getFormattedTitleMarquee(
	titleMarqueeDto *TextLineTitleMarqueeDto,
	strBuilder *strings.Builder,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtTitleMarqueeDtoNanobot.lock == nil {
		txtTitleMarqueeDtoNanobot.lock = new(sync.Mutex)
	}

	txtTitleMarqueeDtoNanobot.lock.Lock()

	defer txtTitleMarqueeDtoNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineTitleMarqueeDtoNanobot."+
			"getFormattedTitleMarquee()",
		"")

	if err != nil {
		return err
	}

	if titleMarqueeDto == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'titleMarqueeDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if strBuilder == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'strBuilder' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	_,
		err = new(textLineTitleMarqueeDtoAtom).
		testValidityOfTitleMarqueeDto(
			titleMarqueeDto,
			ePrefix.XCpy(
				"titleMarqueeDto"))

	if err != nil {
		return err
	}

	txtLineSpecTitleMarquee := TextLineSpecTitleMarquee{}

	err = new(textLineSpecTitleMarqueeMechanics).
		setTxtLineTitleMarqueeDto(
			&txtLineSpecTitleMarquee,
			titleMarqueeDto,
			ePrefix.XCpy(
				"txtLineSpecTitleMarquee<-"+
					"titleMarqueeDto"))

	if err != nil {
		return err
	}

	_,
		_,
		err = new(textLineSpecTitleMarqueeMolecule).
		getFormattedText(
			strBuilder,
			&txtLineSpecTitleMarquee,
			ePrefix.XCpy(
				"txtLineSpecTitleMarquee"))

	return err
}

//	textLineTitleMarqueeDtoAtom
//
//	Provides helper methods for type
//	TextLineTitleMarqueeDto

type textLineTitleMarqueeDtoAtom struct {
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
func (txtTitleMarqueeDtoAtom *textLineTitleMarqueeDtoAtom) calcTextFieldLen(
	txtTitleMarqueeDto *TextLineTitleMarqueeDto,
	fieldLen int,
	errPrefDto *ePref.ErrPrefixDto) (
	validFieldLen int,
	err error) {

	if txtTitleMarqueeDtoAtom.lock == nil {
		txtTitleMarqueeDtoAtom.lock = new(sync.Mutex)
	}

	txtTitleMarqueeDtoAtom.lock.Lock()

	defer txtTitleMarqueeDtoAtom.lock.Unlock()

	validFieldLen = 0

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineTitleMarqueeDtoAtom."+
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
func (txtTitleMarqueeDtoAtom *textLineTitleMarqueeDtoAtom) empty(
	txtTitleMarqueeDto *TextLineTitleMarqueeDto) {

	if txtTitleMarqueeDtoAtom.lock == nil {
		txtTitleMarqueeDtoAtom.lock = new(sync.Mutex)
	}

	txtTitleMarqueeDtoAtom.lock.Lock()

	defer txtTitleMarqueeDtoAtom.lock.Unlock()

	if txtTitleMarqueeDto == nil {
		return
	}

	txtTitleMarqueeDto.StandardSolidLineLeftMargin = ""

	txtTitleMarqueeDto.StandardSolidLineRightMargin = ""

	txtTitleMarqueeDto.StandardTitleLeftMargin = ""

	txtTitleMarqueeDto.StandardTitleRightMargin = ""

	txtTitleMarqueeDto.StandardMaxLineLen = 0

	txtTitleMarqueeDto.StandardTextFieldLen = 0

	txtTitleMarqueeDto.StandardTextJustification =
		TxtJustify.None()

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
//		If input parameter 'txtTitleMarqueeDto' is invalid,
//		the returned error Type will encapsulate an
//		appropriate error message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (txtTitleMarqueeDtoAtom *textLineTitleMarqueeDtoAtom) testValidityOfTitleMarqueeDto(
	txtTitleMarqueeDto *TextLineTitleMarqueeDto,
	errPrefDto *ePref.ErrPrefixDto) (
	isValid bool,
	err error) {

	if txtTitleMarqueeDtoAtom.lock == nil {
		txtTitleMarqueeDtoAtom.lock = new(sync.Mutex)
	}

	txtTitleMarqueeDtoAtom.lock.Lock()

	defer txtTitleMarqueeDtoAtom.lock.Unlock()

	isValid = false

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineTitleMarqueeDtoAtom."+
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
