package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
	"time"
)

type TextUtility struct {
	lock *sync.Mutex
}

//	BuildOneColLeadingMarquee
//
//	Builds a one-column top title display or 'Marquee'
//	and writes the formatted text to a strings.Builder
//	instance passed as an input parameter.
//
//	Be advised. This is the top half of a title Marquee.
//	The bottom half of the title marquee is configured
//	with a call to method:
//
//		TextUtility.BuildOneColTrailingMarquee()
//
// Other alternatives for creating Tile Marquees are
// provided by Type, 'TextLineTitleMarqueeDto'
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	strBuilder					*strings.Builder
//
//		A pointer to an instance of strings.Builder. The
//		formatted text created by this method will be
//		written to this instance of string builder.
//
//	addLeadingBlankLine			bool
//		If set to 'true' a blank line or new line will be
//		positioned immediately before the Top Marquee
//		Solid Line.
//
//		Example: addLeadingBlankLine = 'true'
//			\n (Empty Blank Line)
//			--------------------------
//			        Hello World!
//
//		Example: addLeadingBlankLine = 'false'
//			(No Preceding Blank Line)
//			--------------------------
//			        Hello World!
//
//	marqueeTopSolidLineChar		string
//
//		A single character used to create a solid line
//		above the tiles displayed in the title marquee
//		text output.
//
//		Example: marqueeTopSolidLineChar = "-"
//			The dotted line is the "Top Solid Line Char"
//
//			--------------------------
//			        Hello World!
//
//		If 'marqueeTopSolidLineChar' is submitted as an
//		empty string, no solid line will be configured
//		above the title display.
//
//		Example: marqueeTopSolidLineChar = ""
//			No solid line displayed above first title.
//
//			        Hello World!
//
//	marqueeBottomSolidLineChar	string
//
//		A single character used to create a solid line
//		above the tiles displayed in the title marquee
//		text output.
//
//		Example: marqueeBottomSolidLineChar = "-"
//			The dotted line is the "Bottom Solid Line Char"
//
//			      Goodbye World!
//			--------------------------
//
//		If 'marqueeBottomSolidLineChar' is submitted as
//		empty string, no solid line will be configured
//		below the titles display.
//
//		Example: marqueeBottomSolidLineChar = ""
//			No solid line displayed below last title.
//
//			        Goodbye World!
//
//	addTrailingBlankLine		bool
//		If set to 'true' a blank line or new line will be
//		positioned immediately after the Bottom Marquee
//		Solid Line.
//
//		Example addLeadingBlankLine = 'true'
//
//			        Hello World!
//			--------------------------
//			\n (Empty Blank Line)
//
//		Example addLeadingBlankLine = 'false'
//			(No Preceding Blank Line)
//			==========================
//			        Hello World!
//
//	titles						StringArrayDto
//
//		This String Array Data Transfer Object
//		encapsulates a string array which will supply the
//		Marquee titles. Each title will be treated as a
//		1-Column display and placed on a separate text
//		line.
//
//		If the first element in this array is an empty
//		string, an error will be returned. Empty strings
//		in the remaining array elements will be skipped
//		and ignored.
//
//	txtFmtParams				TextFmtParamsLineColumnsDto
//
//		This instance of TextFmtParamsLineColumnsDto
//		should be configured for a 1-Column Display. The
//		format parameters contained in this instance will be
//		used to format the text for all title lines formatted
//		for text output.
//
//		If this instance of TextFmtParamsLineColumnsDto is
//		invalid, default values will be used as shown below:
//
//		txtFmtParams = TextFmtParamsLineColumnsDto{
//		  FormatType: TxtFieldType.LineColumns(),
//		  FieldFormatParams: []TextFieldFmtParamsDto{
//		     {
//		      LeftMarginStr:  " ",
//		      FieldLength:    79,
//		      FieldJustify:   TxtJustify.Center(),
//		      DateTimeFormat: "2006-01-02 15:04:05.000000000 -0700 MST",
//		      RightMarginStr: "",
//		     },
//		  },
//		  TurnLineTerminationOff:     false,
//		  LineTerminator:             "\n",
//		  MaxLineLength:              79,
//		  TurnAutoLineLengthBreaksOn: true,
//		}
//
//	addTrailingDateTimeStamp	bool
//
//		If this parameter is set to 'true', a date
//		time stamp will be added using the Date/Time
//		format string supplied in the first element of
//		the txtFmtParams.FieldFormatParams array:
//			txtFmtParams.FieldFormatParams[0]
//
//		This date time stamp will be positioned below
//		the last title in the 'titles' array.
//
//		If 'addTrailingDateTimeStamp' is set to false, no
//		date time stamp will be displayed.
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
func (txtUtility *TextUtility) BuildOneColLeadingMarquee(
	strBuilder *strings.Builder,
	addLeadingBlankLine bool,
	marqueeTopSolidLineChar string,
	marqueeBottomSolidLineChar string,
	addTrailingBlankLine bool,
	titles StringArrayDto,
	txtFmtParams TextFmtParamsLineColumnsDto,
	addTrailingDateTimeStamp bool,
	errorPrefix interface{}) (
	err error) {

	if txtUtility.lock == nil {
		txtUtility.lock = new(sync.Mutex)
	}

	txtUtility.lock.Lock()

	defer txtUtility.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextUtility."+
			"BuildOneColLeadingMarquee()",
		"")

	if err != nil {
		return err
	}

	lenTitles := titles.GetStringArrayLength()

	if lenTitles == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'titles' is a zero length\n"+
			"string array!\n",
			ePrefix.String())

		return err
	}

	if txtFmtParams.MaxLineLength < 5 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtFmtParams.MaxLineLength' is invalid!\n"+
			"The minimum value for 'txtFmtParams.MaxLineLength' is five (5).\n"+
			"standardMaxLineLen = '%v'\n",
			ePrefix.String(),
			txtFmtParams.MaxLineLength)

		return err
	}

	cntEmptyStrs := 0

	for i := 0; i < lenTitles; i++ {
		if len(titles.StrArray[i]) == 0 {
			cntEmptyStrs++
		}
	}

	if cntEmptyStrs == lenTitles {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'titles' is invalid!\n"+
			"All the strings in string array 'titles' are\n"+
			"zero length or empty strings.\n",
			ePrefix.String())

		return err

	}

	txtFormatCol := TextFormatterCollection{}

	if txtFmtParams.GetNumOfFieldFmtParams() != 1 {
		// txtFmtParams is invalid,
		// apply default standard format parameters
		// for 1-Column text displays.

		txtFmtParams.FieldFormatParams = nil

		txtFmtParams = TextFmtParamsLineColumnsDto{
			FormatType: TxtFieldType.LineColumns(),
			FieldFormatParams: []TextFieldFmtParamsDto{
				{
					LeftMarginStr: " ",
					FieldLength:   79,
					FieldJustify:  TxtJustify.Center(),
					DateTimeFormat: new(textSpecificationMolecule).
						getDefaultDateTimeFormat(),
					RightMarginStr: "",
					lock:           nil,
				},
			},
			TurnLineTerminationOff:     false,
			LineTerminator:             "\n",
			MaxLineLength:              79,
			TurnAutoLineLengthBreaksOn: true,
			lock:                       nil,
		}

	}

	err = txtFormatCol.
		SetStdFormatParams(
			txtFmtParams,
			ePrefix.XCpy(
				"Default txtFmtParams 1-Column"))

	if err != nil {
		return err
	}

	if addLeadingBlankLine == true {

		// Blank Line
		txtFormatCol.AddLineBlank(
			1,
			"")

	}

	if len(marqueeTopSolidLineChar) > 0 {

		// Filler =======
		// Marquee Top
		txtFormatCol.AddLineSolid(
			" ",
			marqueeTopSolidLineChar,
			txtFmtParams.MaxLineLength-2,
			" ",
			false,
			"",
			-1,
			false,
			" ")
	}

	// Title Line 1
	for i := 0; i < lenTitles; i++ {

		title := titles.StrArray[i]

		if len(title) > 0 {

			err = txtFormatCol.AddLine1Col(
				title,
				ePrefix.XCpy(
					fmt.Sprintf(
						"titles.StrArray[i]")))

			if err != nil {
				return err
			}

		}

	}
	// Last Line Date/Time

	if addTrailingDateTimeStamp == true {

		err = txtFormatCol.AddLine1Col(
			TextInputParamFieldDateTimeDto{
				FieldDateTime:       time.Now(),
				FieldDateTimeFormat: txtFmtParams.FieldFormatParams[0].DateTimeFormat,
			},
			ePrefix.XCpy(
				"Last Line Date/Time"))

		if err != nil {
			return err
		}
	}

	if len(marqueeBottomSolidLineChar) > 0 {
		// Filler Line '========='
		// Marquee Bottom
		txtFormatCol.AddLineSolid(
			" ",
			"=",
			txtFmtParams.MaxLineLength-2,
			" ",
			false,
			"",
			-1,
			false,
			" ")
	}

	if addTrailingBlankLine == true {

		// Trailing Blank Line
		txtFormatCol.AddLineBlank(
			1,
			"")

	}

	// End Of Marquee

	err = txtFormatCol.BuildText(
		strBuilder,
		ePrefix.XCpy(
			"strBuilder<-txtFormatCol"))

	return err
}

//	BuildOneColTrailingMarquee
//
//	Builds a one-column bottom title display or 'Marquee'
//	and writes the formatted text to a strings.Builder
//	instance passed as an input parameter.
//
//	Be advised. This is the bottom half of a title Marquee.
//	The top half of the title marquee is configured with a
//	call to method:
//
//		TextUtility.BuildOneColLeadingMarquee()
//
// Other alternatives for creating Tile Marquees are
// provided by Type, 'TextLineTitleMarqueeDto'
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	strBuilder					*strings.Builder
//
//		A pointer to an instance of strings.Builder. The
//		formatted text created by this method will be
//		written to this instance of string builder.
//
//	strBuilder					*strings.Builder
//
//		A pointer to an instance of strings.Builder. The
//		formatted text created by this method will be
//		written to this instance of string builder.
//
//	addLeadingBlankLine			bool
//		If set to 'true' a blank line or new line will be
//		positioned immediately before the Top Marquee
//		Solid Line.
//
//		Example: addLeadingBlankLine = 'true'
//			\n (Empty Blank Line)
//			--------------------------
//			        Hello World!
//
//		Example: addLeadingBlankLine = 'false'
//			(No Preceding Blank Line)
//			--------------------------
//			        Hello World!
//
//	marqueeTopSolidLineChar		string
//
//		A single character used to create a solid line
//		above the tiles displayed in the title marquee
//		text output.
//
//		Example: marqueeTopSolidLineChar = "-"
//			The dotted line is the "Top Solid Line Char"
//
//			--------------------------
//			        Hello World!
//
//		If 'marqueeTopSolidLineChar' is submitted as an
//		empty string, no solid line will be configured
//		above the title display.
//
//		Example: marqueeTopSolidLineChar = ""
//			No solid line displayed above first title.
//
//			        Hello World!
//
//	marqueeBottomSolidLineChar	string
//
//		A single character used to create a solid line
//		above the tiles displayed in the title marquee
//		text output.
//
//		Example: marqueeBottomSolidLineChar = "-"
//			The dotted line is the "Bottom Solid Line Char"
//
//			      Goodbye World!
//			--------------------------
//
//		If 'marqueeBottomSolidLineChar' is submitted as
//		empty string, no solid line will be configured
//		below the titles display.
//
//		Example: marqueeBottomSolidLineChar = ""
//			No solid line displayed below last title.
//
//			        Goodbye World!
//
//	addTrailingBlankLine		bool
//		If set to 'true' a blank line or new line will be
//		positioned immediately after the Bottom Marquee
//		Solid Line.
//
//		Example addLeadingBlankLine = 'true'
//
//			        Hello World!
//			--------------------------
//			\n (Empty Blank Line)
//
//		Example addLeadingBlankLine = 'false'
//			(No Preceding Blank Line)
//			==========================
//			        Hello World!
//
//	titles						StringArrayDto
//
//		This String Array Data Transfer Object
//		encapsulates a string array which will supply the
//		Marquee titles. Each title will be treated as a
//		1-Column display and placed on a separate text
//		line.
//
//		If the first element in this array is an empty
//		string, an error will be returned. Empty strings
//		in the remaining array elements will be skipped
//		and ignored.
//
//	txtFmtParams				TextFmtParamsLineColumnsDto
//
//		This instance of TextFmtParamsLineColumnsDto
//		should be configured for a 1-Column Display. The
//		format parameters contained in this instance will be
//		used to format the text for all title lines formatted
//		for text output.
//
//		If this instance of TextFmtParamsLineColumnsDto is
//		invalid, default values will be used as shown below:
//
//		txtFmtParams = TextFmtParamsLineColumnsDto{
//		  FormatType: TxtFieldType.LineColumns(),
//		  FieldFormatParams: []TextFieldFmtParamsDto{
//		     {
//		      LeftMarginStr:  " ",
//		      FieldLength:    79,
//		      FieldJustify:   TxtJustify.Center(),
//		      DateTimeFormat: "2006-01-02 15:04:05.000000000 -0700 MST",
//		      RightMarginStr: "",
//		     },
//		  },
//		  TurnLineTerminationOff:     false,
//		  LineTerminator:             "\n",
//		  MaxLineLength:              79,
//		  TurnAutoLineLengthBreaksOn: true,
//		}
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
func (txtUtility *TextUtility) BuildOneColTrailingMarquee(
	strBuilder *strings.Builder,
	addLeadingBlankLine bool,
	marqueeTopSolidLineChar string,
	marqueeBottomSolidLineChar string,
	addTrailingBlankLine bool,
	titles StringArrayDto,
	txtFmtParams TextFmtParamsLineColumnsDto,
	errorPrefix interface{}) (
	err error) {

	if txtUtility.lock == nil {
		txtUtility.lock = new(sync.Mutex)
	}

	txtUtility.lock.Lock()

	defer txtUtility.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextUtility."+
			"BuildOneColTrailingMarquee()",
		"")

	if err != nil {
		return err
	}

	lenTitles := titles.GetStringArrayLength()

	if lenTitles == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'titles' is a zero length\n"+
			"string array!\n",
			ePrefix.String())

		return err
	}

	if txtFmtParams.MaxLineLength < 5 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtFmtParams.MaxLineLength' is invalid!\n"+
			"The minimum value for 'txtFmtParams.MaxLineLength' is five (5).\n"+
			"standardMaxLineLen = '%v'\n",
			ePrefix.String(),
			txtFmtParams.MaxLineLength)

		return err
	}

	cntEmptyStrs := 0

	for i := 0; i < lenTitles; i++ {
		if len(titles.StrArray[i]) == 0 {
			cntEmptyStrs++
		}
	}

	if cntEmptyStrs == lenTitles {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'titles' is invalid!\n"+
			"All the strings in string array 'titles' are\n"+
			"zero length or empty strings.\n",
			ePrefix.String())

		return err

	}

	txtFormatCol := TextFormatterCollection{}

	if txtFmtParams.GetNumOfFieldFmtParams() != 1 {
		// txtFmtParams is invalid,
		// apply default standard format parameters
		// for 1-Column text displays.

		txtFmtParams.FieldFormatParams = nil

		txtFmtParams = TextFmtParamsLineColumnsDto{
			FormatType: TxtFieldType.LineColumns(),
			FieldFormatParams: []TextFieldFmtParamsDto{
				{
					LeftMarginStr: " ",
					FieldLength:   79,
					FieldJustify:  TxtJustify.Center(),
					DateTimeFormat: new(textSpecificationMolecule).
						getDefaultDateTimeFormat(),
					RightMarginStr: "",
					lock:           nil,
				},
			},
			TurnLineTerminationOff:     false,
			LineTerminator:             "\n",
			MaxLineLength:              79,
			TurnAutoLineLengthBreaksOn: true,
			lock:                       nil,
		}
	}

	err = txtFormatCol.
		SetStdFormatParams(
			txtFmtParams,
			ePrefix.XCpy(
				"Default txtFmtParams 1-Column"))

	if err != nil {
		return err
	}

	if addLeadingBlankLine == true {

		// Blank Line
		txtFormatCol.AddLineBlank(
			1,
			"")

	}

	if len(marqueeTopSolidLineChar) > 0 {

		// Filler =======
		// Marquee Top
		txtFormatCol.AddLineSolid(
			" ",
			marqueeTopSolidLineChar,
			txtFmtParams.MaxLineLength-2,
			" ",
			false,
			"",
			-1,
			false,
			" ")

	}

	// Format Title Lines
	for i := 0; i < lenTitles; i++ {

		title := titles.StrArray[i]

		if len(title) > 0 {

			err = txtFormatCol.AddLine1Col(
				title,
				ePrefix.XCpy(
					fmt.Sprintf(
						"titles.StrArray[i]")))

			if err != nil {
				return err
			}

		}

	}

	if len(marqueeBottomSolidLineChar) > 0 {

		// Filler Line '========='
		// Marquee Bottom
		txtFormatCol.AddLineSolid(
			" ",
			marqueeBottomSolidLineChar,
			txtFmtParams.MaxLineLength-2,
			" ",
			false,
			"",
			-1,
			false,
			" ")

	}

	if addTrailingBlankLine == true {

		// Trailing Blank Line
		txtFormatCol.AddLineBlank(
			1,
			"")
	}

	// End Of Marquee

	err = txtFormatCol.BuildText(
		strBuilder,
		ePrefix.XCpy(
			"strBuilder<-txtFormatCol"))

	return err
}
