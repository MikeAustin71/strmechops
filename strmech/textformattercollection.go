package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
	"time"
)

// TextFormatterCollection - This type contains a collection of
// Text Formatter Data Transfer Objects (TextFormatterDto). These
// object contain the specifications used to format text fields and
// lines of text for  screen displays, file output and printing.
type TextFormatterCollection struct {
	fmtCollection []TextFormatterDto
	// Text

	stdTextLineParamCollection []TextFmtParamsLineColumnsDto
	// Standard Format Parameter Collection.
	// Provides standard text formats

	lock *sync.Mutex
}

// AddAdHocText adds raw, or ad hoc text to the Formatter
// Collection.
//
// The input parameters for this method are used to inject user
// generated text into the stream of text characters being
// formatted for screen display, file output or printing.
//
// Except for line breaks configured at the user's discretion,
// no additional formatting is performed on this text, and it is
// inserted raw, or "as is", in to the final output of formatted
// text.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	leftMarginStr                      string
//	   - A string containing the text characters to be positioned
//	     on the Left side of the Ad Hoc Text.
//
//	     If no Left margin is required, set this parameter to an
//	     empty string.
//
//
//	adHocText                          string
//	   - This strings holds the raw, ad hoc text which will be
//	     inserted "as is", into the final output of formatted
//	     text.
//
//	     If this parameter is submitted as an empty or zero
//	     length string, an error will be generated when
//	     attempting to create formatted text output.
//
//
//
//	rightMarginStr                     string
//	   - The contents of the string will be used as the right
//	     margin for the Text Ad Hoc string.
//
//	     If no right margin is required, set 'rightMarginStr' to
//	     a zero length or empty string, and no right margin will
//	     be created.
//
//
//	turnLineTerminationOff             bool
//	   - By default, a new line string terminator ('\n') will be
//	     appended to the Ad Hoc text ('AdHocText'). If this
//	     parameter is set to 'true', no line termination sequence
//	     will be applied.
//
//	     This parameter controls the operation of parameter
//	     'LineTerminator'. If 'TurnLineTerminationOff' is set to
//	     'true', 'LineTerminator' will be completely ignored and
//	      have no effect.
//
//
//	lineTerminator                     string
//	   - This string holds the character or characters which
//	     will be used to terminate the formatted line of text
//	     output, if parameter 'TurnLineTerminationOff' is set
//	     to 'false'.
//
//	     The most common usage sets this string to a new line
//	     character ("\n").
//
//	     If 'lineTerminator' is configured as an empty string
//	     (string length zero), a single new line character ('\n')
//	     will be automatically applied to produce line termination
//	     depending on the setting for parameter
//	     'turnLineTerminationOff'.
//
//	     LineTerminator works in conjunction with parameter
//	     'turnLineTerminationOff'. 'turnLineTerminationOff'
//	     controls the application of a line terminator. Setting
//	     'turnLineTerminationOff' to 'true' means that NO line
//	     terminator will be applied.
//
//	     Setting 'turnLineTerminationOff' to 'true' means that
//	     parameter 'lineTerminator' will be completely ignored
//	     and have no effect.
//
//
//	maxLineLength                      int
//	   - The maximum length of the line on which the solid line
//	     text characters will be presented.
//
//	     Set this parameter to minus one (-1) to specify an
//	     unlimited line length for this text line.
//
//	     If the value of 'maxLineLength' is less than one (1),
//	     it will be automatically converted to minus one (-1).
//
//	     'maxLineLength' is used in conjunction with parameter
//	     'turnAutoLineLengthBreaksOn' to automatically place text
//	     on separate text lines when that text exceeds the maximum
//	     text line length ('maxLineLength'). Therefore, parameter
//	     'turnAutoLineLengthBreaksOn' controls whether automatic
//	     line breaks using 'maxLineLength' will be applied.
//
//
//	turnAutoLineLengthBreaksOn         bool
//	   - This parameter controls whether text lines which exceed
//	     the maximum line length ('maxLineLength') are positioned
//	     on the following line as a separate line of text.
//
//	     To apply automatic line breaking at the maximum line
//	     length, set the value of this parameter to 'true'.
//
//	     When this parameter is set to 'true', text fields which
//	     extend beyond the maximum line length ('maxLineLength')
//	     will be formatted as a separate line of text on the
//	     following line.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	NONE
func (txtFmtCollection *TextFormatterCollection) AddAdHocText(
	leftMarginStr string,
	adHocText string,
	rightMarginStr string,
	turnLineTerminationOff bool,
	lineTerminator string,
	maxLineLength int,
	turnAutoLineLengthBreaksOn bool) {

	if txtFmtCollection.lock == nil {
		txtFmtCollection.lock = new(sync.Mutex)
	}

	txtFmtCollection.lock.Lock()

	defer txtFmtCollection.lock.Unlock()

	if maxLineLength < 1 {
		maxLineLength = -1
	}

	newTextFormatter := TextFormatterDto{
		FormatType:          TxtFieldType.TextAdHoc(),
		DateTime:            TextFieldDateTimeDto{},
		Filler:              TextFieldFillerDto{},
		Label:               TextFieldLabelDto{},
		Spacer:              TextFieldSpacerDto{},
		BlankLine:           TextLineBlankDto{},
		SolidLine:           TextLineSolidDto{},
		LineColumns:         TextLineColumnsDto{},
		LinesTimerStartStop: TextLineTimerStartStopDto{},
		TextAdHoc: TextAdHocDto{
			FormatType:                 TxtFieldType.TextAdHoc(),
			LeftMarginStr:              leftMarginStr,
			AdHocText:                  adHocText,
			RightMarginStr:             rightMarginStr,
			TurnLineTerminationOff:     turnLineTerminationOff,
			LineTerminator:             lineTerminator,
			MaxLineLength:              maxLineLength,
			TurnAutoLineLengthBreaksOn: turnAutoLineLengthBreaksOn,
			lock:                       nil,
		},

		lock: nil,
	}

	txtFmtCollection.fmtCollection =
		append(
			txtFmtCollection.fmtCollection,
			newTextFormatter)

	return
}

// AddAdHocTextDto - Adds raw, or ad hoc text to the Formatter
// Collection.
//
// The ad hoc text is configured from an instance of TextAdHocDto
// passed as an input parameter.
//
// The input parameters for this method are used to inject user
// generated into the stream of text characters being formatted
// for screen display, file output or printing.
//
// Except for line breaks configured at the user's discretion,
// no additional formatting is performed on this text, and it is
// inserted raw, or "as is", in to the final output of formatted
// text.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	txtAdHocDto                        TextAdHocDto
//	   - An instance of TextFieldLabelDto which contains all the
//	     necessary data parameters to produce a text label.
//
//	     The Text Field Label Data Transfer Object is defined as
//	     follows:
//
//	      type TextAdHocDto struct {
//
//	       FormatType                  TextFieldType
//	        Required. This enumeration value specifies the type of
//	        Text Format Operation to be performed.
//
//	        For TextAdHocDto 'FormatType' this parameter should
//	        be set to: TxtFieldType.TextAdHoc()
//
//	       LeftMarginStr               string
//	        A string containing the text characters to be
//	        positioned on the Left side of the Ad Hoc Text.
//
//	        If no Left margin is required, set this parameter to an
//	        empty string.
//
//	       AdHocText                   string
//	        This strings holds the raw ad hoc text.
//
//	       RightMarginStr              string
//	        The contents of the string will be used as the right
//	        margin for the Text Ad Hoc string.
//
//	        If no right margin is required, set 'RightMarginStr' to
//	        a zero length or empty string, and no right margin will
//	        be created.
//
//	       TurnLineTerminationOff      bool
//	        By default, a new line string terminator ('\n') will be
//	        appended to the Ad Hoc text ('AdHocText'). If this
//	        parameter is set to 'true', no line termination
//	        sequence will be applied.
//
//	        This parameter controls the operation of parameter
//	        'LineTerminator'. If 'TurnLineTerminationOff' is set
//	        to 'true', 'LineTerminator' will be completely
//	        ignored and have no effect.
//
//	       LineTerminator              string
//	        This string holds the character or characters which
//	        will be used to terminate the formatted line of text
//	        output, if parameter 'TurnLineTerminationOff' is set
//	        to 'false'.
//
//	        The most common usage sets this string to a new line
//	        character ("\n").
//
//	        If 'LineTerminator' is configured as an empty string
//	        (string length zero), a single new line character
//	        ('\n') will be automatically applied to produce line
//	        termination depending on the setting for parameter
//	        'TurnLineTerminationOff'.
//
//	        LineTerminator works in conjunction with member
//	        variable 'TurnLineTerminationOff'.
//
//	        'TurnLineTerminationOff' controls the application of a
//	        line terminator. Setting 'TurnLineTerminationOff' to
//	        'true' means that NO line terminator will be applied.
//
//	        Setting 'TurnLineTerminationOff' to 'true' means that
//	        parameter 'LineTerminator' will be completely ignored
//	        and have no effect.
//
//	       MaxLineLength               int
//	        The maximum length of the line on which the ad hoc
//	        text characters ('AdHocText') will be presented.
//
//	        Set this parameter to minus one (-1) to specify an
//	        unlimited line length for this text line.
//
//	        If the value of 'MaxLineLength' is less than one (1),
//	        it will be automatically converted to minus one (-1).
//
//	        'MaxLineLength' is used in conjunction with parameter
//	        'TurnAutoLineLengthBreaksOn' to automatically place
//	        text on separate text lines when that text exceeds the
//	        maximum text line length ('MaxLineLength'). Therefore,
//	        paramter 'TurnAutoLineLengthBreaksOn' controls whether
//	        automatic line breaks using 'MaxLineLength' will be
//	        applied.
//
//	       TurnAutoLineLengthBreaksOn  bool
//	        This parameter controls whether text lines which exceed
//	        the maximum line length ('MaxLineLength') are
//	        positioned on the following line as a separate line of
//	        text.
//
//	        To apply automatic line breaking at the maximum line
//	        length, set the value of this parameter to 'true'.
//
//	        When this parameter is set to 'true', text fields which
//	        extend beyond the maximum line length 'MaxLineLength'
//	        will be formatted as a separate line of text on the
//	        following line.
//	      }
//
// ----------------------------------------------------------------
//
// Return Values
//
//	NONE
func (txtFmtCollection *TextFormatterCollection) AddAdHocTextDto(
	txtAdHocDto TextAdHocDto) {

	if txtFmtCollection.lock == nil {
		txtFmtCollection.lock = new(sync.Mutex)
	}

	txtFmtCollection.lock.Lock()

	defer txtFmtCollection.lock.Unlock()

	if txtAdHocDto.MaxLineLength < 1 {
		txtAdHocDto.MaxLineLength = -1
	}

	if txtAdHocDto.FormatType != TxtFieldType.TextAdHoc() {

		txtAdHocDto.FormatType = TxtFieldType.TextAdHoc()
	}

	newTextFormatter := TextFormatterDto{
		FormatType:          TxtFieldType.TextAdHoc(),
		DateTime:            TextFieldDateTimeDto{},
		Filler:              TextFieldFillerDto{},
		Label:               TextFieldLabelDto{},
		Spacer:              TextFieldSpacerDto{},
		BlankLine:           TextLineBlankDto{},
		SolidLine:           TextLineSolidDto{},
		LineColumns:         TextLineColumnsDto{},
		LinesTimerStartStop: TextLineTimerStartStopDto{},
		TextAdHoc:           txtAdHocDto.CopyOut(),

		lock: nil,
	}

	newTextFormatter.TextAdHoc.CopyIn(
		txtAdHocDto)

	txtFmtCollection.fmtCollection =
		append(
			txtFmtCollection.fmtCollection,
			newTextFormatter)

	return
}

// AddFieldDateTime - Adds a date time value formatted as a text
// field to the Formatter Collection.
//
// Users have the option to format this Text Date Time Field with
// a line terminator (a.k.a. new line character '\n') thereby
// creating a separate stand-alone line of text.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	leftMarginStr              string
//	   - The contents of this string will be used as the left
//	     margin for the date time text field.
//
//	     If no left margin is required, set 'leftMarginStr' to a
//	     zero length or empty string, and no left margin will be
//	     created.
//
//
//	fieldDateTime              time.Time
//	   - The date/time value which will be formatted as a text
//	     string.
//
//	     If this parameter is set equal to zero, an error will be
//	     generated when attempting to create a text field with a
//	     Text Builder method.
//
//
//	fieldDateTimeFormat        string
//	  - This string will be used to format the date/time value
//	    'fieldDateTime' as a text string.
//
//	     If this 'fieldDateTimeFormat' string is empty (has a zero
//	     length), a default Date/Time format string will be applied
//	     as follows:
//	       "2006-01-02 15:04:05.000000000 -0700 MST"
//
//
//	fieldLength                int
//	   - Used to format date time text field. This is the length of
//	     the text field in which the formatted date time text
//	     string will be displayed. If 'fieldLength' is less than the
//	     length of the date time string, it will be automatically
//	     set equal to the date time string length.
//
//	     If 'fieldLength' is greater than the length of the date
//	     time text string, the 'fieldJustify' parameter will be
//	     used to configure or justify the text within the
//	     boundaries the text field defined by 'fieldLength'.
//
//	     To automatically set the value of 'fieldLength' to the
//	     length of the date time text field, set this parameter to
//	     a value of  minus one (-1).
//
//
//	fieldJustify               TextJustify
//	    An enumeration value specifying the justification of the
//	    date time text string within the text field specified by
//	    'fieldLength'.
//
//	    Text justification can only be evaluated in the context of
//	    a text label, field length and a Text Justification object
//	    of type TextJustify. This is because text labels with a
//	    field length equal to or less than the length of the text
//	    label never use text justification. In these cases, text
//	    justification is completely ignored.
//
//	    If the field length is greater than the length of the date
//	    time text string, text justification must be equal to one
//	    of these three valid values:
//	        TextJustify(0).Left()
//	        TextJustify(0).Right()
//	        TextJustify(0).Center()
//
//	    You can also use the abbreviated text justification
//	    enumeration syntax as follows:
//
//	        TxtJustify.Left()
//	        TxtJustify.Right()
//	        TxtJustify.Center()
//
//
//	lineTerminator             string
//	   - This string holds the character or characters which will
//	     be used to terminate the formatted date time text thereby
//	     converting this text element into a valid line of text.
//
//	     If a text line is required, setting this string to include
//	     a new line character ('\n') will ensure that the text line
//	     consists of the text label field and no other text
//	     elements. Any string of text characters will be accepted
//	     for this parameter.
//
//	     Again, the most common usage sets this string to a new
//	     line character ("\n").
//
//	     If Line Termination is NOT required, set 'lineTerminator'
//	     to a zero length or empty string and no line termination
//	     characters will be created.
//
//
//	maxLineLength              int
//	   - The maximum length of the line on which this date time
//	     text field will be presented.
//
//	     Set this parameter to minus one (-1) to specify an
//	     unlimited line length for this text line.
//
//	     'maxLineLength' is used in conjunction with parameter
//	     'turnAutoLineLengthBreaksOn' to automatically place text
//	     fields on separate text lines when that text exceeds the
//	     maximum text line length ('maxLineLength'). Therefore,
//	     paramter 'turnAutoLineLengthBreaksOn' controls whether
//	     automatic line breaks using 'maxLineLength' will be
//	     applied.
//
//	     If the value of 'maxLineLength' is less than zero (0), it
//	     will be automatically converted to minus one (-1).
//
//
//	turnAutoLineLengthBreaksOn bool
//	   - This parameter controls whether text lines which exceed
//	     the maximum line length ('maxLineLength') are broken up
//	     and presented on the following line.
//
//	     To apply automatic line breaking at the maximum line
//	     length, set the value of this parameter to 'true'.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	NONE
func (txtFmtCollection *TextFormatterCollection) AddFieldDateTime(
	leftMarginStr string,
	fieldDateTime time.Time,
	fieldDateTimeFormat string,
	fieldLength int,
	fieldJustify TextJustify,
	rightMarginStr string,
	lineTerminator string,
	maxLineLength int,
	turnAutoLineLengthBreaksOn bool) {

	if txtFmtCollection.lock == nil {
		txtFmtCollection.lock = new(sync.Mutex)
	}

	txtFmtCollection.lock.Lock()

	defer txtFmtCollection.lock.Unlock()

	if maxLineLength < 1 {
		maxLineLength = -1
	}

	if fieldLength < 1 {
		fieldLength = -1
	}

	if len(fieldDateTimeFormat) == 0 {
		fieldDateTimeFormat =
			new(textSpecificationMolecule).
				getDefaultDateTimeFormat()
	}

	newTextFormatter := TextFormatterDto{
		FormatType: TxtFieldType.DateTime(),
		DateTime: TextFieldDateTimeDto{
			FormatType:                 TxtFieldType.DateTime(),
			LeftMarginStr:              leftMarginStr,
			FieldDateTime:              fieldDateTime,
			FieldDateTimeFormat:        fieldDateTimeFormat,
			FieldLength:                fieldLength,
			FieldJustify:               fieldJustify,
			RightMarginStr:             rightMarginStr,
			LineTerminator:             lineTerminator,
			MaxLineLength:              maxLineLength,
			TurnAutoLineLengthBreaksOn: turnAutoLineLengthBreaksOn,
		},
		Filler:              TextFieldFillerDto{},
		Label:               TextFieldLabelDto{},
		Spacer:              TextFieldSpacerDto{},
		BlankLine:           TextLineBlankDto{},
		SolidLine:           TextLineSolidDto{},
		LineColumns:         TextLineColumnsDto{},
		LinesTimerStartStop: TextLineTimerStartStopDto{},
	}

	txtFmtCollection.fmtCollection =
		append(
			txtFmtCollection.fmtCollection,
			newTextFormatter)

	return
}

// AddFieldDateTimeDto - Adds a date time value formatted as a text
// field to the Formatter Collection. Users have the option to
// format this date time text field with a line terminator
// (a.k.a. new line character '\n').
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	textDateTimeDto            TextFieldDateTimeDto
//	   - An instance of TextFieldDateTimeDto which contains all the
//	     necessary data parameters to produce a date time text
//	     field.
//
//	     The Text Field Date Time Data Transfer Object is defined
//	     as follows:
//
//
//	     type TextFieldDateTimeDto struct {
//
//	       FormatType                  TextFieldType
//	        Required. This enumeration value specifies the type of
//	        Text Format Operation to be performed.
//
//	        For the TextFieldDateTimeDto Format Type, this value
//	        should always be set to:
//	          TxtFieldType.DateTime()
//
//	       LeftMarginStr               string
//	        The contents of the string will be used as the left
//	        margin for the Text Field.
//
//	        If no left margin is required, set 'LeftMarginStr' to a
//	        zero length or empty string, and no left margin will be
//	        created.
//
//	       FieldDateTime               time.Time
//	        This time value will be converted to a string and used
//	        to populate the Date Time Text Field.
//
//	       FieldDateTimeFormat         string
//	        This string will be used to format the Date/Time value
//	        contained in parameter 'FieldDateTime' as text.
//
//	        If 'FieldDateTime' is set to a value greater than zero
//	        and this 'FieldDateTimeFormat' string is empty (has a
//	        zero length), a default Date/Time format string will be
//	        applied as follows:
//	                "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	       FieldLength                 int
//	        Used to format Text Fields. This is the length of the
//	        text field in which the date time text field will be
//	        displayed. If 'FieldLength' is less than the length of
//	        the date time text string, it will be automatically set
//	        equal to the date time text string length.
//
//	        If 'FieldLength' is greater than the length of the date
//	        time text string, the 'FieldJustify' parameter will be
//	        used to configure or justify the text with the
//	        boundaries of the text field defined by 'FieldLength'.
//
//	        To automatically set the value of 'FieldLength' to the
//	        length of the date time text string, set this parameter
//	        to a value of minus one (-1).
//
//	        If this parameter is submitted with a value less than
//	        minus one (-1) or greater than 1-million (1,000,000),
//	        an error will be returned when attempting to build the
//	        final text output.
//
//
//	       FieldJustify                TextJustify
//	        An enumeration which specifies the justification of the
//	        date time text string within the text field specified
//	        by 'FieldLength'.
//
//	        Text justification can only be evaluated in the context
//	        of a text label, field length and a Text Justification
//	        object of type TextJustify. This is because text labels
//	        with a field length equal to or less than the length of
//	        the text label never use text justification. In these
//	        cases, text justification is completely ignored.
//
//	        If the field length is greater than the length of the
//	        date time text string, text justification must be equal
//	        to one of these three valid values:
//	            TextJustify(0).Left()
//	            TextJustify(0).Right()
//	            TextJustify(0).Center()
//
//	        You can also use the abbreviated text justification
//	        enumeration syntax as follows:
//
//	            TxtJustify.Left()
//	            TxtJustify.Right()
//	            TxtJustify.Center()
//
//	       RightMarginStr string
//	        The contents of the string will be used as the right margin
//	        for the Text Field.
//
//	        If no right margin is required, set 'RightMarginStr' to a
//	        zero length or empty string, and no right margin will be
//	        created.
//
//	       LineTerminator              string
//	        This string holds the character or characters which will be
//	        used to terminate the formatted line of text output.
//
//	        The most common usage sets this string to a new line
//	        character ("\n").
//
//	        If no Line Terminator is required, set 'lineTerminator' to
//	        a zero length or empty string and no line termination
//	        characters will be created.
//
//	       MaxLineLength               int
//	        Set this parameter to minus one -1 to specify an
//	        unlimited line length for this text line.
//
//	       TurnAutoLineLengthBreaksOn  bool
//	        When this parameter is set to 'true', text fields which
//	        extend beyond the maximum line length ('MaxLineLength')
//	        will be terminated with a new line character ('\n') and
//	        placed on the following line of text.
//
//	     }
//
// ----------------------------------------------------------------
//
// Return Values
//
//	NONE
func (txtFmtCollection *TextFormatterCollection) AddFieldDateTimeDto(
	textDateTimeDto TextFieldDateTimeDto) {

	if txtFmtCollection.lock == nil {
		txtFmtCollection.lock = new(sync.Mutex)
	}

	txtFmtCollection.lock.Lock()

	defer txtFmtCollection.lock.Unlock()

	if textDateTimeDto.MaxLineLength < 1 {
		textDateTimeDto.MaxLineLength = -1
	}

	if textDateTimeDto.FieldLength < 1 {
		textDateTimeDto.FieldLength = -1
	}

	if len(textDateTimeDto.FieldDateTimeFormat) == 0 {
		textDateTimeDto.FieldDateTimeFormat =
			new(textSpecificationMolecule).
				getDefaultDateTimeFormat()
	}

	newTextFormatter := TextFormatterDto{
		FormatType:          TxtFieldType.DateTime(),
		DateTime:            textDateTimeDto.CopyOut(),
		Filler:              TextFieldFillerDto{},
		Label:               TextFieldLabelDto{},
		Spacer:              TextFieldSpacerDto{},
		BlankLine:           TextLineBlankDto{},
		SolidLine:           TextLineSolidDto{},
		LineColumns:         TextLineColumnsDto{},
		LinesTimerStartStop: TextLineTimerStartStopDto{},
	}

	txtFmtCollection.fmtCollection =
		append(
			txtFmtCollection.fmtCollection,
			newTextFormatter)

	return
}

// AddFieldDateTimeInputDto - Adds a date time value formatted as a
// text field to the Formatter Collection. Users have the option to
// format  this date time text field with a line terminator (a.k.a.
// new line character '\n').
//
// The date time value and date format are configured through input
// parameter 'dateTimeInputDto', an instance of
// TextInputParamFieldDateTimeDto.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	leftMarginStr              string
//	   - The contents of this string will be used as the left
//	     margin for the date time text field.
//
//	     If no left margin is required, set 'leftMarginStr' to a
//	     zero length or empty string, and no left margin will be
//	     created.
//
//
//	dateTimeDto                TextInputParamFieldDateTimeDto
//	   - A data structure containing both the date time value and
//	     the date time format string necessary to convert the date
//	     time value into formatted text.
//
//	     The TextInputParamFieldDateTimeDto data structure is
//	     defined as follows:
//
//	      type TextInputParamFieldDateTimeDto struct {
//
//	        FieldDateTime time.Time
//	         This time value will be used to populate a Text Field
//	         used for screen display, file output or printing.
//
//	        FieldDateTimeFormat string
//	         This string will be used to format the date time value
//	         contained in the 'FieldDateTime' data element.
//
//	         If 'FieldDateTime' is set to a value greater than zero
//	         and this 'FieldDateTimeFormat' string is empty (has a
//	         zero length), a default Date/Time format string will be
//	         applied as follows:
//	             "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	      }
//
//
//	fieldLength                int
//	   - Used to format date time text field. This is the length of
//	     the text field in which the formatted date time text
//	     string will be displayed. If 'fieldLength' is less than the
//	     length of the date time string, it will be automatically
//	     set equal to the date time string length.
//
//	     If 'fieldLength' is greater than the length of the date
//	     time text string, the 'fieldJustify' parameter will be
//	     used to configure or justify the text within the
//	     boundaries the text field defined by 'fieldLength'.
//
//	     To automatically set the value of 'fieldLength' to the
//	     length of the date time text field, set this parameter to
//	     a value of  minus one (-1).
//
//
//	fieldJustify               TextJustify
//	    An enumeration value specifying the justification of the
//	    date time text string within the text field specified by
//	    'fieldLength'.
//
//	    Text justification can only be evaluated in the context of
//	    a text label, field length and a Text Justification object
//	    of type TextJustify. This is because text labels with a
//	    field length equal to or less than the length of the text
//	    label never use text justification. In these cases, text
//	    justification is completely ignored.
//
//	    If the field length is greater than the length of the date
//	    time text string, text justification must be equal to one
//	    of these three valid values:
//	        TextJustify(0).Left()
//	        TextJustify(0).Right()
//	        TextJustify(0).Center()
//
//	    You can also use the abbreviated text justification
//	    enumeration syntax as follows:
//
//	        TxtJustify.Left()
//	        TxtJustify.Right()
//	        TxtJustify.Center()
//
//
//	lineTerminator             string
//	   - This string holds the character or characters which will
//	     be used to terminate the formatted date time text thereby
//	     converting this text element into a valid line of text.
//
//	     If a text line is required, setting this string to include
//	     a new line character ('\n') will ensure that the text line
//	     consists of the text label field and no other text
//	     elements. Any string of text characters will be accepted
//	     for this parameter.
//
//	     Again, the most common usage sets this string to a new
//	     line character ("\n").
//
//	     If Line Termination is NOT required, set 'lineTerminator'
//	     to a zero length or empty string and no line termination
//	     characters will be created.
//
//
//	maxLineLength              int
//	   - The maximum length of the line on which this date time
//	     text field will be presented.
//
//	     Set this parameter to minus one (-1) to specify an
//	     unlimited line length for this text line.
//
//	     'maxLineLength' is used in conjunction with parameter
//	     'turnAutoLineLengthBreaksOn' to automatically place text
//	     fields on separate text lines when that text exceeds the
//	     maximum text line length ('maxLineLength'). Therefore,
//	     paramter 'turnAutoLineLengthBreaksOn' controls whether
//	     automatic line breaks using 'maxLineLength' will be
//	     applied.
//
//	     If the value of 'maxLineLength' is less than zero (0), it
//	     will be automatically converted to minus one (-1).
//
//
//	turnAutoLineLengthBreaksOn bool
//	   - This parameter controls whether text lines which exceed
//	     the maximum line length ('maxLineLength') are broken up
//	     and presented on the following line.
//
//	     To apply automatic line breaking at the maximum line
//	     length, set the value of this parameter to 'true'.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	NONE
func (txtFmtCollection *TextFormatterCollection) AddFieldDateTimeInputDto(
	leftMarginStr string,
	dateTimeDto TextInputParamFieldDateTimeDto,
	fieldLength int,
	fieldJustify TextJustify,
	rightMarginStr string,
	lineTerminator string,
	maxLineLength int,
	turnAutoLineLengthBreaksOn bool) {

	if txtFmtCollection.lock == nil {
		txtFmtCollection.lock = new(sync.Mutex)
	}

	txtFmtCollection.lock.Lock()

	defer txtFmtCollection.lock.Unlock()

	if maxLineLength < 1 {
		maxLineLength = -1
	}

	if fieldLength < 1 {
		fieldLength = -1
	}

	if len(dateTimeDto.FieldDateTimeFormat) == 0 {
		dateTimeDto.FieldDateTimeFormat =
			new(textSpecificationMolecule).
				getDefaultDateTimeFormat()
	}

	newTextFormatter := TextFormatterDto{
		FormatType: TxtFieldType.DateTime(),
		DateTime: TextFieldDateTimeDto{
			FormatType:                 TxtFieldType.DateTime(),
			LeftMarginStr:              leftMarginStr,
			FieldDateTime:              dateTimeDto.FieldDateTime,
			FieldDateTimeFormat:        dateTimeDto.FieldDateTimeFormat,
			FieldLength:                fieldLength,
			FieldJustify:               fieldJustify,
			RightMarginStr:             rightMarginStr,
			LineTerminator:             lineTerminator,
			MaxLineLength:              maxLineLength,
			TurnAutoLineLengthBreaksOn: turnAutoLineLengthBreaksOn,
		},
		Filler:              TextFieldFillerDto{},
		Label:               TextFieldLabelDto{},
		Spacer:              TextFieldSpacerDto{},
		BlankLine:           TextLineBlankDto{},
		SolidLine:           TextLineSolidDto{},
		LineColumns:         TextLineColumnsDto{},
		LinesTimerStartStop: TextLineTimerStartStopDto{},
	}

	txtFmtCollection.fmtCollection =
		append(
			txtFmtCollection.fmtCollection,
			newTextFormatter)

	return
}

// AddFieldFiller - Adds a Text Filler Field to the Formatter
// Collection.
//
// Text Filler Fields are commonly used as margins containing
// multiple white space characters, or line separators containing
// multiple dashes, equal signs or underscore characters. Text
// Filler Fields consist of filler characters ('fillerCharacters')
// and the filler characters repeat count
// ('fillerCharsRepeatCount'). A filler field is made up of one or
// more filler characters. These filler characters are repeated one
// or more times in order to construct the complete filler field as
// shown in the following examples:
//
//	Example 1:
//	 Filler Characters = "-"
//	 Filler Characters Repeat Count = 3
//	 Formatted Text = "---"
//
//	Example 2:
//	 Filler Characters = "-*"
//	 Filler Characters Repeat Count = 3
//	 Formatted Text = "-*-*-*"
//
//	Example 3:
//	 Filler Characters = "-*"
//	 Filler Characters Repeat Count = 3
//	 Line Terminator = "\n"
//	 Formatted Text = "-*-*-*\n"
//
// Typically, Text Filler Fields are designed to be configured
// within a line of text. However, users have the option of
// configuring a Text Filler Field as a separate stand-alone
// line of text by means of the input parameter 'lineTerminator'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	leftMarginStr              string
//	   - The contents of this string will be used as the left
//	     margin for the Text Filler Field.
//
//	     If no left margin is required, set 'leftMarginStr' to a
//	     zero length or empty string, and no left margin will be
//	     created.
//
//
//	fillerCharacters           string
//	   - A string containing the text characters which will be
//	     included in the Text Filler Field. The final Text Filler
//	     Field will be constructed from the filler characters
//	     repeated one or more times as specified by the
//	     'fillerCharsRepeatCount' parameter.
//
//	      Text Field Filler Length =
//	        Length of fillerCharacters X fillerCharsRepeatCount
//
//	        Example #1: FillerCharacters = "-*"
//	                    FillerCharsRepeatCount = 3
//	                    Final Text Filler Field = "-*-*-*"
//
//	        Example #2: FillerCharacters = "-"
//	                    FillerCharsRepeatCount = 3
//	                    Final Text Filler Field = "---"
//
//
//	fillerCharsRepeatCount     int
//	     Controls the number of times 'fillerCharacters' is
//	     repeated when constructing the final Text Filler Field.
//	     The actual length of the string which will populate the
//	     completed Text Filler Field is equal to the length of
//	     'fillerCharacters' times the value of
//	     'fillerCharsRepeatCount'.
//
//	      Text Field Filler Length =
//	        Length of fillerCharacters X fillerCharsRepeatCount
//
//	        Example #1: fillerCharacters = "-*"
//	                    fillerRepeatCount = 3
//	                    Final Text Filler Field = "-*-*-*"
//
//	        Example #2: fillerCharacters = "-"
//	                    fillerRepeatCount = 3
//	                    Final Text Filler Field = "---"
//
//	     If 'fillerCharsRepeatCount' has a value less than one (1)
//	     or greater than one-million (1,000,000), an error will be
//	     generated when attempting to format text.
//
//
//	lineTerminator             string
//	   - This string holds the character or characters which will
//	     be used to terminate the formatted text thereby converting
//	     this text element into a valid line of text.
//
//	     If a text line is required, setting this string to include
//	     a new line character ('\n') will ensure that the text line
//	     consists of the text label field and no other text
//	     elements. Any string of text characters will be accepted
//	     for this parameter.
//
//	     Again, the most common usage sets this string to a new
//	     line character ("\n").
//
//	     If Line Termination is NOT required, set 'lineTerminator'
//	     to a zero length or empty string and no line termination
//	     characters will be created.
//
//
//	maxLineLength              int
//	   - The maximum length of the line on which this label text
//	     will be presented.
//
//	     Set this parameter to minus one (-1) to specify an
//	     unlimited line length for this text line.
//
//	     'maxLineLength' is used in conjunction with parameter
//	     'turnAutoLineLengthBreaksOn' to automatically place text
//	     fields on separate text lines when that text exceeds the
//	     maximum text line length ('maxLineLength'). Therefore,
//	     paramter 'turnAutoLineLengthBreaksOn' controls whether
//	     automatic line breaks using 'maxLineLength' will be
//	     applied.
//
//	     If the value of 'maxLineLength' is less than zero (0), it
//	     will be automatically converted to minus one (-1).
//
//
//	turnAutoLineLengthBreaksOn bool
//	   - This parameter controls whether text lines which exceed
//	     the maximum line length ('maxLineLength') are broken up
//	     and presented on the following line.
//
//	     To apply automatic line breaking at the maximum line
//	     length, set the value of this parameter to 'true'.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	NONE
func (txtFmtCollection *TextFormatterCollection) AddFieldFiller(
	leftMarginStr string,
	fillerCharacters string,
	fillerCharsRepeatCount int,
	rightMarginStr string,
	lineTerminator string,
	maxLineLength int,
	turnAutoLineLengthBreaksOn bool) {

	if txtFmtCollection.lock == nil {
		txtFmtCollection.lock = new(sync.Mutex)
	}

	txtFmtCollection.lock.Lock()

	defer txtFmtCollection.lock.Unlock()

	if maxLineLength < 1 {
		maxLineLength = -1
	}

	newTextFormatter := TextFormatterDto{
		FormatType: TxtFieldType.Filler(),
		DateTime:   TextFieldDateTimeDto{},
		Filler: TextFieldFillerDto{
			FormatType:                 TxtFieldType.Filler(),
			LeftMarginStr:              leftMarginStr,
			FillerCharacters:           fillerCharacters,
			FillerCharsRepeatCount:     fillerCharsRepeatCount,
			RightMarginStr:             rightMarginStr,
			LineTerminator:             lineTerminator,
			MaxLineLength:              maxLineLength,
			TurnAutoLineLengthBreaksOn: turnAutoLineLengthBreaksOn,
			lock:                       nil,
		},
		Label:               TextFieldLabelDto{},
		Spacer:              TextFieldSpacerDto{},
		BlankLine:           TextLineBlankDto{},
		SolidLine:           TextLineSolidDto{},
		LineColumns:         TextLineColumnsDto{},
		LinesTimerStartStop: TextLineTimerStartStopDto{},
	}

	txtFmtCollection.fmtCollection =
		append(
			txtFmtCollection.fmtCollection,
			newTextFormatter)

	return
}

// AddFieldFillerDto - Adds a Text Filler Field to the Formatter
// Collection generated from an input parameter of type
// TextFieldFillerDto.
//
// Text Filler Fields are commonly used as margins containing
// multiple white space characters, or line separators containing
// multiple dashes, equal signs or underscore characters. Text
// Filler Fields consist of filler characters ('fillerCharacters')
// and the filler characters repeat count
// ('fillerCharsRepeatCount'). A filler field is made up of one or
// more filler characters. These filler characters are repeated one
// or more times in order to construct the complete filler field as
// shown in the following examples:
//
//	Example 1:
//	 Filler Characters = "-"
//	 Filler Characters Repeat Count = 3
//	 Formatted Text = "---"
//
//	Example 2:
//	 Filler Characters = "-*"
//	 Filler Characters Repeat Count = 3
//	 Formatted Text = "-*-*-*"
//
//	Example 3:
//	 Filler Characters = "-*"
//	 Filler Characters Repeat Count = 3
//	 Line Terminator = "\n"
//	 Formatted Text = "-*-*-*\n"
//
// Typically, Text Filler Fields are designed to be configured
// within a line of text. However, users have the option of
// configuring a Text Filler Field as a separate stand-alone
// line of text.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	textFillerDto              TextFieldFillerDto
//	   - An instance of TextFieldFillerDto which contains all the
//	     necessary data parameters to produce a Text Filler Field.
//
//	     The Text Field Filler Data Transfer Object is defined as
//	     follows:
//
//	     type TextFieldFillerDto struct {
//
//	      FormatType                   TextFieldType
//	       Required. This enumeration value specifies the type of
//	       Text Format Operation to be performed.
//
//	       For the TextFieldFillerDto Format Type, this value
//	       should always be set to:
//	         TxtFieldType.Filler()
//
//	      LeftMarginStr                string
//	       The contents of the string will be used as the left
//	       margin for the Text Field.
//
//	       If no left margin is required, set 'LeftMarginStr' to a
//	       zero length or empty string, and no left margin will be
//	       created.
//
//	      FillerCharacters             string
//	       A string containing the text characters which will be
//	       included in the Text Filler Field. The final Text Filler
//	       Field will be constructed from the filler characters
//	       repeated one or more times as specified by the
//	       'FillerCharsRepeatCount' parameter.
//
//	        Text Field Filler Length =
//	          Length of fillerCharacters X fillerCharsRepeatCount
//
//	          Example #1: FillerCharacters = "-*"
//	                      FillerCharsRepeatCount = 3
//	                      Final Text Filler Field = "-*-*-*"
//
//	          Example #2: FillerCharacters = "-"
//	                      FillerCharsRepeatCount = 3
//	                      Final Text Filler Field = "---"
//
//
//	      FillerCharsRepeatCount       int
//	       Controls the number of times 'FillerCharacters' is
//	       repeated when constructing the final Text Filler Field.
//	       The actual length of the string which will populate the
//	       completed Text Filler Field is equal to the length of
//	       'FillerCharacters' times the value of
//	       'FillerCharsRepeatCount'.
//
//	        Text Field Filler Length =
//	          Length of FillerCharacters X FillerCharsRepeatCount
//
//	          Example #1: FillerCharacters = "-*"
//	                      FillerRepeatCount = 3
//	                      Final Text Filler Field = "-*-*-*"
//
//	          Example #2: FillerCharacters = "-"
//	                      FillerRepeatCount = 3
//	                      Final Text Filler Field = "---"
//
//	       If 'FillerCharsRepeatCount' has a value less than one
//	       (1) or greater than one-million (1,000,000), an error
//	       will be returned when attempting to format text.
//
//	      RightMarginStr               string
//	       The contents of the string will be used as the right
//	       margin for the Text Filler Field.
//
//	       If no right margin is required, set 'RightMarginStr' to
//	       a zero length or empty string, and no right margin will
//	       be created.
//
//	      LineTerminator               string
//	       This string holds the character or characters which will
//	       be used to terminate the formatted line of text output
//	       thereby converting this text element into a valid
//	       stand-alone line of text. Line Termination is optional.
//	       Populate this string only if this text output should be
//	       formatted as a separate line of text.
//
//	       The most common usage sets this string to a new line
//	       character ("\n").
//
//	       If no Line Terminator is required, set 'lineTerminator'
//	       to a zero length or empty string and no line termination
//	       characters will be created.
//
//	      MaxLineLength                int
//	       The maximum length of the line on which this label text
//	       will be presented.
//
//	       Set this parameter to minus one (-1) to specify an
//	       unlimited line length for this text line.
//
//	       'MaxLineLength' is used in conjunction with parameter
//	       'TurnAutoLineLengthBreaksOn' to automatically place text
//	       fields on separate text lines when that text exceeds the
//	       maximum text line length ('MaxLineLength'). Therefore,
//	       paramter 'turnAutoLineLengthBreaksOn' controls whether
//	       automatic line breaks using 'MaxLineLength' will be
//	       applied.
//
//	       If the value of 'maxLineLength' is less than one (1), it
//	       will be automatically converted to minus one (-1).
//
//	       Set this parameter to minus one (-1) to specify an
//	       unlimited line length for this text line.
//
//	      TurnAutoLineLengthBreaksOn   bool
//	       This parameter controls whether text lines which exceed
//	       the maximum line length ('MaxLineLength') are broken up
//	       and presented on the following line.
//
//	       To apply automatic line breaking at the maximum line
//	       length, set the value of this parameter to 'true'.
//	     }
//
// ----------------------------------------------------------------
//
// Return Values
//
//	NONE
func (txtFmtCollection *TextFormatterCollection) AddFieldFillerDto(
	textFillerDto TextFieldFillerDto) {

	if txtFmtCollection.lock == nil {
		txtFmtCollection.lock = new(sync.Mutex)
	}

	txtFmtCollection.lock.Lock()

	defer txtFmtCollection.lock.Unlock()

	if textFillerDto.MaxLineLength < 1 {
		textFillerDto.MaxLineLength = -1
	}

	textFillerDto.FormatType = TxtFieldType.Filler()

	newTextFormatter := TextFormatterDto{
		FormatType:          TxtFieldType.Filler(),
		DateTime:            TextFieldDateTimeDto{},
		Filler:              textFillerDto.CopyOut(),
		Label:               TextFieldLabelDto{},
		Spacer:              TextFieldSpacerDto{},
		BlankLine:           TextLineBlankDto{},
		SolidLine:           TextLineSolidDto{},
		LineColumns:         TextLineColumnsDto{},
		LinesTimerStartStop: TextLineTimerStartStopDto{},
	}

	txtFmtCollection.fmtCollection =
		append(
			txtFmtCollection.fmtCollection,
			newTextFormatter)

	return
}

// AddFieldLabel - Adds a single text label field to the Formatter
// Collection.
//
// Users have the option to format this text label with a line
// terminator (a.k.a. new line character '\n') thereby creating a
// separate stand-alone line of text.
//
// Text Label Example 1:
//
//	leftMarginStr = "" // Empty String
//	fieldText = "Hello"  // Length = 5 characters
//	fieldLength = 7
//	fieldJustify = TxtJustify.Center()
//	rightMarginStr = "" // Empty String
//	lineTerminator = "" // Empty String
//	maxLineLength = -1
//	turnAutoLineLengthBreaksOn = false
//	Final Text Label string = " Hello "
//
// Text Label Example 2:
//
//	leftMarginStr = "" // Empty String
//	fieldText = "Hello"  // Length = 5 characters
//	fieldLength = 7
//	fieldJustify = TxtJustify.Center()
//	rightMarginStr = "" // Empty String
//	lineTerminator = "\n" // Empty String
//	maxLineLength = -1
//	turnAutoLineLengthBreaksOn = false
//	Final Text Label string = " Hello \n"
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	leftMarginStr              string
//	   - The contents of this string will be used as the left
//	     margin for the text label field.
//
//	     If no left margin is required, set 'leftMarginStr' to a
//	     zero length or empty string, and no left margin will be
//	     created.
//
//
//	fieldText                  interface{}
//	   - This parameter is an empty interface which must contain
//	     one of several specific types. This empty interface type
//	     will be converted to a string and configured as the text
//	     field for this label.
//
//	     Supported types which may be submitted through this empty
//	     interface parameter are listed as follows:
//	        time.Time (Converted using default format)
//	        string
//	        bool
//	        uint, uint8, uint16, uint32, uint64,
//	        int, int8, int16, int32, int64
//	        float32, float64
//	        *big.Int *big.Float
//	        fmt.Stringer (types that support this interface)
//	        TextInputParamFieldDateTimeDto
//	              (Converts date time to string)
//
//	     If the 'fieldText' is not convertible to one of the
//	     supported types, the 'column1Field' string will be
//	     populated with an error message.
//
//	     If the converted string value for 'column1Field' is empty,
//	     it will be defaulted to a single white space character
//	     (" ").
//
//
//	fieldLength                int
//	   - Used to format Text Label Fields. This is the length of
//	     the text field in which the formatted 'fieldText' string
//	     will be displayed. If 'fieldLength' is less than the
//	     length of the 'fieldText' string, it will be automatically
//	     set equal to the 'labelText' string length.
//
//	     If 'fieldLength' is greater than the length of the
//	     'fieldText' text string, the 'fieldJustify' parameter will
//	     be used to configure or justify the text within the
//	     boundaries of the text field defined by 'fieldLength'.
//
//	     To automatically set the value of 'fieldLength' to the
//	     length of 'fieldText', set this parameter to a value
//	     of  minus one (-1).
//
//	     If this parameter is submitted with a value less than
//	     minus one (-1) or greater than 1-million (1,000,000),
//	     an error will be returned when attempting to build the
//	     final text output.
//
//
//	fieldJustify               TextJustify
//	    An enumeration value specifying the justification of the
//	    'fieldText' string within the text field specified by
//	    'fieldLength'.
//
//	    Text justification can only be evaluated in the context of
//	    a text label, field length and a Text Justification object
//	    of type TextJustify. This is because text labels with a
//	    field length equal to or less than the length of the text
//	    label never use text justification. In these cases, text
//	    justification is completely ignored.
//
//	    If the field length is greater than the length of the text
//	    label, text justification must be equal to one of these
//	    three valid values:
//	        TextJustify(0).Left()
//	        TextJustify(0).Right()
//	        TextJustify(0).Center()
//
//	    You can also use the abbreviated text justification
//	    enumeration syntax as follows:
//
//	        TxtJustify.Left()
//	        TxtJustify.Right()
//	        TxtJustify.Center()
//
//
//	rightMarginStr             string
//	   - The contents of this string will be used as the right
//	     margin for the text label field.
//
//	     If no right margin is required, set 'rightMarginStr' to a
//	     zero length or empty string, and no right margin will be
//	     created.
//
//
//	lineTerminator             string
//	   - This string holds the character or characters which will
//	     be used to terminate the formatted text thereby converting
//	     this text element into a valid line of text.
//
//	     If a text line is required, setting this string to include
//	     a new line character ('\n') will ensure that the text line
//	     consists of the text label field and no other text
//	     elements. Any string of text characters will be accepted
//	     for this parameter.
//
//	     Again, the most common usage sets this string to a new
//	     line character ("\n").
//
//	     If Line Termination is NOT required, set 'lineTerminator'
//	     to a zero length or empty string and no line termination
//	     characters will be created.
//
//
//	maxLineLength              int
//	   - The maximum length of the line on which this label text
//	     will be presented.
//
//	     Set this parameter to minus one (-1) to specify an
//	     unlimited line length for this text line.
//
//	     If the value of 'maxLineLength' is less than zero (0), it
//	     will be automatically converted to minus one (-1).
//
//	     'maxLineLength' is used in conjunction with parameter
//	     'turnAutoLineLengthBreaksOn' to automatically place text
//	     fields on separate text lines when that text exceeds the
//	     maximum text line length ('maxLineLength'). Therefore,
//	     paramter 'turnAutoLineLengthBreaksOn' controls whether
//	     automatic line breaks using 'maxLineLength' will be
//	     applied.
//
//
//	turnAutoLineLengthBreaksOn bool
//	   - This parameter controls whether text lines which exceed
//	     the maximum line length ('maxLineLength') are broken up
//	     and presented on the following line.
//
//	     To apply automatic line breaking at the maximum line
//	     length, set the value of this parameter to 'true'.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	NONE
func (txtFmtCollection *TextFormatterCollection) AddFieldLabel(
	leftMarginStr string,
	fieldText interface{},
	fieldLength int,
	fieldJustify TextJustify,
	rightMarginStr string,
	lineTerminator string,
	maxLineLength int,
	turnAutoLineLengthBreaksOn bool) {

	if txtFmtCollection.lock == nil {
		txtFmtCollection.lock = new(sync.Mutex)
	}

	txtFmtCollection.lock.Lock()

	defer txtFmtCollection.lock.Unlock()

	if maxLineLength < 1 {
		maxLineLength = -1
	}

	if fieldLength < 1 {
		fieldLength = -1
	}

	var fieldTextStr string
	var err error

	fieldTextStr,
		err = new(textSpecificationAtom).
		convertParamEmptyInterfaceToString(
			fieldText,
			"fieldText",
			nil)

	if err != nil {
		fieldTextStr = "ERROR interface{} to string conversion!"
	}

	newTextFormatter := TextFormatterDto{
		FormatType: TxtFieldType.Label(),
		DateTime:   TextFieldDateTimeDto{},
		Filler:     TextFieldFillerDto{},
		Label: TextFieldLabelDto{
			FormatType:                 TxtFieldType.Label(),
			LeftMarginStr:              leftMarginStr,
			FieldText:                  fieldTextStr,
			FieldLength:                fieldLength,
			FieldJustify:               fieldJustify,
			RightMarginStr:             rightMarginStr,
			LineTerminator:             lineTerminator,
			MaxLineLength:              maxLineLength,
			TurnAutoLineLengthBreaksOn: turnAutoLineLengthBreaksOn,
		},
		Spacer:              TextFieldSpacerDto{},
		BlankLine:           TextLineBlankDto{},
		SolidLine:           TextLineSolidDto{},
		LineColumns:         TextLineColumnsDto{},
		LinesTimerStartStop: TextLineTimerStartStopDto{},
	}

	txtFmtCollection.fmtCollection =
		append(
			txtFmtCollection.fmtCollection,
			newTextFormatter)

	return
}

// AddFieldLabelDto - Adds a single text label field to the Formatter
// Collection. Users have the option to format the text label with
// a line terminator (a.k.a. new line character '\n').
//
// This text label is created from an instance of TextFieldLabelDto
// passed as an input parameter.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	textLabelDto                       TextFieldLabelDto
//	   - An instance of TextFieldLabelDto which contains all the
//	     necessary data parameters to produce a text label.
//
//	     The Text Field Label Data Transfer Object is defined as
//	     follows:
//
//	     type TextFieldLabelDto struct {
//
//	       FormatType                  TextFieldType
//	         Required. This enumeration value specifies the type of
//	         Text Format Operation to be performed.
//
//	         For TextFieldLabelDto Format Type, this value should
//	         always be set to:
//	           TxtFieldType.Label()
//
//	       LeftMarginStr               string
//	         The contents of the string will be used as the left
//	         margin for the Text Field.
//
//	         If no left margin is required, set 'LeftMarginStr' to
//	         a zero length or empty string, and no left margin will
//	         be created.
//
//	       FieldText                   string
//	         The Text Field string or contents. If this string is
//	         empty (has a zero (0) length) it will be automatically
//	         converted to a single white space character (" ").
//
//	         This string represents the contents of the Text Field.
//
//	       FieldLength                 int
//	         Used to format FieldText string. This is the length of
//	         the text field in which the 'FieldText' will be
//	         displayed. If 'FieldLength' is less than the length of
//	         the 'FieldText' string, it will be automatically set
//	         equal to the 'FieldText' string length.
//
//	         If 'FieldLength' is greater than the length of the
//	         'FieldText' text string, the 'FieldJustify' parameter
//	         will be used to configure or justify the text within
//	         the boundaries of the text field defined by
//	         'FieldLength'.
//
//	         To automatically set the value of 'FieldLength' to the
//	         length of 'FieldText', set this parameter to a value
//	         of minus one (-1).
//
//	         If this parameter is submitted with a value less than
//	         minus one (-1) or greater than 1-million (1,000,000),
//	         an error will be generated when attempting to format
//	         text.
//
//	       FieldJustify                TextJustify
//	         An enumeration which specifies the justification of
//	         the 'FieldText' string within the text field specified
//	         by 'FieldLength'.
//
//	         Text justification can only be evaluated in the
//	         context of a text label, field length and a Text
//	         Justification object of type TextJustify. This is
//	         because text labels with a field length equal to or
//	         less than the length of the text label never use text
//	         justification. In these cases, text justification is
//	         completely ignored.
//
//	         If the field length is greater than the length of the
//	         text label, text justification must be equal to one of
//	         these three valid values:
//	             TextJustify(0).Left()
//	             TextJustify(0).Right()
//	             TextJustify(0).Center()
//
//	         You can also use the abbreviated text justification
//	         enumeration syntax as follows:
//
//	             TxtJustify.Left()
//	             TxtJustify.Right()
//	             TxtJustify.Center()
//
//	       RightMarginStr              string
//	         The contents of the string will be used as the right
//	         margin for the Text Field.
//
//	         If no right margin is required, set 'RightMarginStr'
//	         to a zero length or empty string, and no right margin
//	         will be created.
//
//	       LineTerminator              string
//	         This string holds the character or characters which
//	         will be used to terminate the formatted line of text
//	         output thereby converting this text element into a
//	         valid line of text. Line Termination is optional.
//	         Populate this string only if this text output should
//	         be formatted as a separate line of text.
//
//	         The most common usage sets this string to a new line
//	         character ("\n"); however, any string of text
//	         characters will be accepted.
//
//	         If no Line Terminator is required, set
//	         'LineTerminator' to a zero length or empty string and
//	         no line termination characters will be created.
//
//	       MaxLineLength               int
//	         The maximum length of the line on which this label
//	         text will be presented.
//
//	         Set this parameter to minus one (-1) to specify an
//	         unlimited line length for this text line.
//
//	         'MaxLineLength' is used in conjunction with parameter
//	         'TurnAutoLineLengthBreaksOn' to automatically place
//	         text fields on separate text lines when that text
//	         exceeds the maximum text line length('MaxLineLength').
//	         Therefore, parameter 'turnAutoLineLengthBreaksOn'
//	         controls whether automatic line breaks using
//	         'MaxLineLength' will be applied.
//
//	         If the value of 'MaxLineLength' is less than one (1),
//	         it will be automatically converted to minus one (-1).
//
//	         Set this parameter to minus one (-1) to specify an
//	         unlimited line length for this text line.
//
//	       TurnAutoLineLengthBreaksOn  bool
//	         This parameter controls whether text lines which
//	         exceed the maximum line length ('MaxLineLength') are
//	         broken up and presented on the following line.
//
//	         To apply automatic line breaking at the maximum line
//	         length, set the value of this parameter to 'true'.
//
//	     }
//
// ----------------------------------------------------------------
//
// Return Values
//
//	NONE
func (txtFmtCollection *TextFormatterCollection) AddFieldLabelDto(
	textLabelDto TextFieldLabelDto) {

	if txtFmtCollection.lock == nil {
		txtFmtCollection.lock = new(sync.Mutex)
	}

	txtFmtCollection.lock.Lock()

	defer txtFmtCollection.lock.Unlock()

	if textLabelDto.MaxLineLength < 1 {
		textLabelDto.MaxLineLength = -1
	}

	if textLabelDto.FieldLength < 1 {
		textLabelDto.FieldLength = -1
	}

	newTextFormatter := TextFormatterDto{
		FormatType:          TxtFieldType.Label(),
		DateTime:            TextFieldDateTimeDto{},
		Filler:              TextFieldFillerDto{},
		Label:               textLabelDto.CopyOut(),
		Spacer:              TextFieldSpacerDto{},
		BlankLine:           TextLineBlankDto{},
		SolidLine:           TextLineSolidDto{},
		LineColumns:         TextLineColumnsDto{},
		LinesTimerStartStop: TextLineTimerStartStopDto{},
	}

	txtFmtCollection.fmtCollection =
		append(
			txtFmtCollection.fmtCollection,
			newTextFormatter)

	return
}

// AddFieldSpacer - Adds a Text Spacer Field to the Formatter
// Collection.
//
// Text Spacer Fields consist of one or more white space characters
// (" ").
//
// Text Spacer Field Example-1:
//
//	LeftMarginStr = ""
//	FieldLength = 3
//	RightMarginStr = ""
//	LineTerminator = ""
//	Final Text Spacer string = "   " // 3-white spaces
//
// Text Spacer Field Example-2:
//
//	LeftMarginStr = ""
//	FieldLength = 3
//	RightMarginStr = ""
//	LineTerminator = "\n"
//	Final Text Spacer string = "   \n" // 3-white spaces and
//	                                   // 1-new line character
//	                                   // ('\n')
//
// Typically, Text Spacer Fields are designed to be configured
// within a line of text. However, users have the option of
// configuring a Text Filler Field as a separate stand-alone
// line of text by configuring the input parameter
// 'lineTerminator'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	leftMarginStr              string
//	   - The contents of this string will be used as the left
//	     margin for the Text Spacer Field.
//
//	     If no left margin is required, set 'leftMarginStr' to a
//	     zero length or empty string, and no left margin will be
//	     created.
//
//
//	fieldLength                int
//	   - An integer value used to specify the number of white space
//	     characters in the Text Spacer Field.
//
//	     If the value of this parameter is less than zero and greater
//	     than one-million (1,000,000), an error will be generated when
//	     attempting to create formatted text output.
//
//	     Examples:
//	      fieldLen = 1 produces text field " "  // 1-white space
//	      fieldLen = 2 produces text field "  " // 2-white spaces
//	      fieldLen = 5 produces text field "     " // 5-white spaces
//
//
//	lineTerminator             string
//	   - This string holds the character or characters which will
//	     be used to terminate the formatted text thereby converting
//	     this text element into a valid line of text.
//
//	     If a text line is required, setting this string to include
//	     a new line character ('\n') will ensure that the text line
//	     consists of the text spacer field and no other text
//	     elements. Any string of text characters will be accepted
//	     for this parameter.
//
//	     Again, the most common usage sets this string to a new
//	     line character ("\n").
//
//	     If Line Termination is NOT required, set 'lineTerminator'
//	     to a zero length or empty string and no line termination
//	     characters will be created.
//
//
//	maxLineLength              int
//	   - The maximum length of the line on which this Text Spacer
//	     Field will be presented.
//
//	     Set this parameter to minus one (-1) to specify an
//	     unlimited line length for this text line.
//
// /
//
//	     'maxLineLength' is used in conjunction with parameter
//	     'turnAutoLineLengthBreaksOn' to automatically place text
//	     fields on separate text lines when that text exceeds the
//	     maximum text line length ('maxLineLength'). Therefore,
//	     paramter 'turnAutoLineLengthBreaksOn' controls whether
//	     automatic line breaks using 'maxLineLength' will be
//	     applied.
//
//	     If the value of 'maxLineLength' is less than zero (0), it
//	     will be automatically converted to minus one (-1).
//
//
//	turnAutoLineLengthBreaksOn bool
//	   - This parameter controls whether text lines which exceed
//	     the maximum line length ('maxLineLength') are broken up
//	     and presented on the following line.
//
//	     To apply automatic line breaking at the maximum line
//	     length, set the value of this parameter to 'true'.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	NONE
func (txtFmtCollection *TextFormatterCollection) AddFieldSpacer(
	leftMarginStr string,
	fieldLength int,
	rightMarginStr string,
	lineTerminator string,
	maxLineLength int,
	turnAutoLineLengthBreaksOn bool) {

	if txtFmtCollection.lock == nil {
		txtFmtCollection.lock = new(sync.Mutex)
	}

	txtFmtCollection.lock.Lock()

	defer txtFmtCollection.lock.Unlock()

	if maxLineLength < 1 {
		maxLineLength = -1
	}

	newTextFormatter := TextFormatterDto{
		FormatType: TxtFieldType.Spacer(),
		DateTime:   TextFieldDateTimeDto{},
		Filler:     TextFieldFillerDto{},
		Label:      TextFieldLabelDto{},
		Spacer: TextFieldSpacerDto{
			FormatType:                 TxtFieldType.Spacer(),
			LeftMarginStr:              leftMarginStr,
			FieldLength:                fieldLength,
			RightMarginStr:             rightMarginStr,
			LineTerminator:             lineTerminator,
			MaxLineLength:              maxLineLength,
			TurnAutoLineLengthBreaksOn: turnAutoLineLengthBreaksOn,
			lock:                       nil,
		},
		BlankLine:           TextLineBlankDto{},
		SolidLine:           TextLineSolidDto{},
		LineColumns:         TextLineColumnsDto{},
		LinesTimerStartStop: TextLineTimerStartStopDto{},
	}

	txtFmtCollection.fmtCollection =
		append(
			txtFmtCollection.fmtCollection,
			newTextFormatter)

	return
}

// AddFieldSpacerDto - Adds a Text Spacer Field to the Formatter
// Collection generated from an input parameter of type
// TextFieldSpacerDto.
//
// Text Spacer Fields consist of one or more white space
// characters (" ").
//
// Text Spacer Field Example-1:
//
//	LeftMarginStr = ""
//	FieldLength = 3
//	RightMarginStr = ""
//	LineTerminator = ""
//	Final Text Spacer string = "   " // 3-white spaces
//
// Text Spacer Field Example-2:
//
//	LeftMarginStr = ""
//	FieldLength = 3
//	RightMarginStr = ""
//	LineTerminator = "\n"
//	Final Text Spacer string = "   \n" // 3-white spaces and
//	                                   // 1-new line character
//	                                   // ('\n')
//
// Typically, Text Spacer Fields are designed to be configured
// within a line of text. However, users have the option of
// configuring a Text Filler Field as a separate stand-alone
// line of text by means of the input parameter 'lineTerminator'.
// line of text by means of the input parameter 'lineTerminator'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	txtFieldSpacerDto          TextFieldSpacerDto
//	   - An instance of TextFieldSpacerDto which contains all the
//	     necessary data parameters to produce a Text Spacer Field.
//
//	     The Text Field Spacer Data Transfer Object is defined as
//	     follows:
//
//	     type TextFieldSpacerDto struct {
//
//	      FormatType                   TextFieldType
//	       Required. This enumeration value specifies the type of
//	       Text Format Operation to be performed.
//
//	       For the TextFieldSpacerDto Format Type, this value
//	       should always be set to: TxtFieldType.Spacer()
//
//	      LeftMarginStr                string
//	       The contents of the string will be used as the left
//	       margin for the Text Spacer Field.
//
//	       If no left margin is required, set 'leftMarginStr' to a
//	       zero length or empty string, and no left margin will be
//	       created.
//
//	      FieldLength                  int
//	       An integer value used to specify the number of white
//	       space characters in the Text Spacer Field.
//
//	       Values less than zero and greater than one million
//	       (1,000,001) will generate errors when attempting to
//	       format text output.
//
//	       Examples:
//	        fieldLen = 1 produces text field " "
//	            // 1-white space
//
//	        fieldLen = 2 produces text field "  "
//	           // 2-white spaces
//
//	        fieldLen = 5 produces text field "     "
//	           // 5-white spaces
//
//	      RightMarginStr               string
//	       The contents of the string will be used as the right
//	       margin for the Text Spacer Field.
//
//	       If no right margin is required, set 'rightMarginStr' to
//	       a zero length or empty string, and no right margin will
//	       be created.
//
//	      LineTerminator               string
//	       This string holds the character or characters which will
//	       be used to terminate the formatted line of text output.
//
//	       The most common usage sets this string to a new line
//	       character ("\n").
//
//	       If no Line Terminator is required, set 'lineTerminator'
//	       to a zero length or empty string and no line
//	       termination characters will be created.
//
//	      MaxLineLength                int
//	       The maximum length of the line on which this label text
//	       will be presented.
//
//	       Set this parameter to minus one (-1) to specify an
//	       unlimited line length for this text line.
//
//	       'MaxLineLength' is used in conjunction with parameter
//	       'TurnAutoLineLengthBreaksOn' to automatically place text
//	       fields on separate text lines when that text exceeds the
//	       maximum text line length ('MaxLineLength'). Therefore,
//	       paramter 'turnAutoLineLengthBreaksOn' controls whether
//	       automatic line breaks using 'MaxLineLength' will be
//	       applied.
//
//	       If the value of 'maxLineLength' is less than one (1), it
//	       will be automatically converted to minus one (-1).
//
//	       Set this parameter to minus one (-1) to specify an
//	       unlimited line length for this text line.
//
//	      TurnAutoLineLengthBreaksOn   bool
//	       This parameter controls whether text lines which exceed
//	       the maximum line length ('MaxLineLength') are broken up
//	       and presented on the following line.
//
//	       To apply automatic line breaking at the maximum line
//	       length, set the value of this parameter to 'true'.
//
//	     }
//
// ----------------------------------------------------------------
//
// Return Values
//
//	NONE
func (txtFmtCollection *TextFormatterCollection) AddFieldSpacerDto(
	txtFieldSpacerDto TextFieldSpacerDto) {

	if txtFmtCollection.lock == nil {
		txtFmtCollection.lock = new(sync.Mutex)
	}

	txtFmtCollection.lock.Lock()

	defer txtFmtCollection.lock.Unlock()

	if txtFieldSpacerDto.MaxLineLength < 1 {
		txtFieldSpacerDto.MaxLineLength = -1
	}

	newTextFormatter := TextFormatterDto{
		FormatType:          TxtFieldType.Spacer(),
		DateTime:            TextFieldDateTimeDto{},
		Filler:              TextFieldFillerDto{},
		Label:               TextFieldLabelDto{},
		Spacer:              txtFieldSpacerDto.CopyOut(),
		BlankLine:           TextLineBlankDto{},
		SolidLine:           TextLineSolidDto{},
		LineColumns:         TextLineColumnsDto{},
		LinesTimerStartStop: TextLineTimerStartStopDto{},
	}

	txtFmtCollection.fmtCollection =
		append(
			txtFmtCollection.fmtCollection,
			newTextFormatter)

	return
}

// AddLine1Col - Adds a single Text Field used to generate a
// 1-Column Text Line.
//
// The 1-Column Text Line is designed to produce a single line
// of text consisting of one text field or text column with
// optional left and right margins. This single text field is
// referred to as 'Column-1'.
//
// This method will assign previously configured (a.k.a. default)
// Format Parameters to this 1-Column Text Line. The prior
// configuration of these 1-Column Format Parameters is a
// requirement and errors will be generated if these Standard
// Format Parameters have not previously been created.
//
// This method will extract those previously created Standard
// Format Parameters for 1-Column Text Lines from the Standard Text
// Line Parameters collection maintained by this instance of
// TextFormatterCollection.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// The standard parameters for 1-Column Text Lines must be
// configured in the Standard Text Line Parameters Collection
// before calling this method.
//
// If the standard parameters for 1-Column Text Lines were not
// previously configured in the Standard Text Line Parameters
// Collection, an error will be returned.
//
// ----------------------------------------------------------------
//
// To configure the standard parameters for 1-Column Text Lines,
// call one of the following methods:
//
//	TextFormatterCollection.CfgLine1Col()
//	TextFormatterCollection.SetStdFormatParamsLine1Col()
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//		column1Field               interface{}
//		   - This parameter is an empty interface which must contain
//		     one of several specific types. This empty interface type
//		     will be converted to a string and configured as the single
//		     text field in this 'Line1Column' Text Line.
//
//		     Supported types which may be submitted through this empty
//		     interface parameter are listed as follows:
//		        time.Time (Converted using default format)
//		        string
//		        bool
//		        uint, uint8, uint16, uint32, uint64,
//		        int, int8, int16, int32, int64
//		        float32, float64
//		        *big.Int *big.Float
//		        fmt.Stringer (types that support this interface)
//		        TextInputParamFieldDateTimeDto
//		              (Converts date time to string)
//
//		     If the 'column1Field' is not convertible to one of the
//		     supported types, an error will be returned.
//
//		     If the converted string value for 'column1Field' is empty,
//		     it will be defaulted to a single white space character
//		     (" ").
//
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
//
//		     This empty interface must be convertible to one of the
//		     following types:
//
//		     1. nil - A nil value is valid and generates an empty
//		        collection of error prefix and error context
//		        information.
//
//		     2. string - A string containing error prefix information.
//
//		     3. []string A one-dimensional slice of strings containing
//		        error prefix information
//
//		     4. [][2]string A two-dimensional slice of strings
//		        containing error prefix and error context information.
//
//		     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		        from this object will be copied for use in error and
//		        informational messages.
//
//		     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		        Information from this object will be copied for use in
//		        error and informational messages.
//
//		     7. IBasicErrorPrefix - An interface to a method generating
//		        a two-dimensional slice of strings containing error
//		        prefix and error context information.
//
//		     If parameter 'errorPrefix' is NOT convertible to one of
//		     the valid types listed above, it will be considered
//		     invalid and trigger the return of an error.
//
//		     Types ErrPrefixDto and IBasicErrorPrefix are included in
//		     the 'errpref' software package,
//		     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
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
func (txtFmtCollection *TextFormatterCollection) AddLine1Col(
	column1Field interface{},
	errorPrefix interface{}) error {

	if txtFmtCollection.lock == nil {
		txtFmtCollection.lock = new(sync.Mutex)
	}

	txtFmtCollection.lock.Lock()

	defer txtFmtCollection.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFormatterCollection."+
			"AddLine1Col()",
		"")

	if err != nil {
		return err
	}

	var foundStdParams bool
	var stdLineColsFmt TextFmtParamsLineColumnsDto

	foundStdParams,
		stdLineColsFmt,
		err =
		new(textFormatterCollectionElectron).
			findStdTxtLineParameters(
				txtFmtCollection,
				1, // This is a one column Text Line
				ePrefix.XCpy(
					"Text Line 1-Column"))

	if err != nil {
		return err
	}

	if !foundStdParams {

		err = fmt.Errorf("%v - Error\n"+
			"Could NOT locate Standard Text Line Parameter Format\n"+
			"for a 1-Column Text Line.\n"+
			"Use the 'SetStdFormatParams()' method to configure a\n"+
			"new Standard Text Line Column Parameters Format for this\n"+
			"1-Column Text Line Type. ",
			ePrefix.String())

		return err
	}

	var column1FieldText string

	column1FieldText,
		err = new(textSpecificationAtom).
		convertParamEmptyInterfaceToString(
			column1Field,
			"column1Field",
			ePrefix.XCpy(
				"column1Field"))

	if err != nil {
		return err
	}

	if len(column1FieldText) == 0 {
		column1FieldText = " "
	}

	newTextLine1Cols := TextLineColumnsDto{
		FormatType: TxtFieldType.LineColumns(),
		TextFieldsContent: []TextFieldsContentDto{
			{
				TextFieldString:   column1FieldText,
				TextFieldDateTime: time.Time{},
				lock:              nil,
			},
		},
		FmtParameters: stdLineColsFmt,
		lock:          nil,
	}

	newTextFormatter := TextFormatterDto{
		FormatType:          TxtFieldType.LineColumns(),
		DateTime:            TextFieldDateTimeDto{},
		Filler:              TextFieldFillerDto{},
		Label:               TextFieldLabelDto{},
		Spacer:              TextFieldSpacerDto{},
		BlankLine:           TextLineBlankDto{},
		SolidLine:           TextLineSolidDto{},
		LineColumns:         newTextLine1Cols,
		LinesTimerStartStop: TextLineTimerStartStopDto{},
	}

	txtFmtCollection.fmtCollection =
		append(
			txtFmtCollection.fmtCollection,
			newTextFormatter)

	return err
}

// AddLine2Col - Adds two Text Fields used to generate a 2-Column
// Text Lne.
//
// The 2-Column Text Line is designed to produce a single line of
// text consisting of two text fields or text columns with
// optional left and right margins.
//
// This method will assign previously configured (a.k.a. default)
// Format Parameters to this 2-Column Text Line. The prior
// configuration of these 2-Column Format Parameters is a
// requirement and errors will be generated if these Standard
// Format Parameters have not previously been created.
//
// This method will extract those previously created Standard
// Format Parameters for 2-Column Text Lines from the Standard Text
// Line Parameters collection maintained by this instance of
// TextFormatterCollection.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// The standard parameters for 2-Column Text Lines must be
// configured in the Standard Text Line Parameters Collection
// before calling this method.
//
// If the standard parameters for 2-Column Text Lines were not
// previously configured in the Standard Text Line Parameters
// Collection, an error will be returned.
//
// ----------------------------------------------------------------
//
// To configure the standard parameters for 2-Column Text Lines,
// call one of the following methods:
//
//	TextFormatterCollection.CfgLine2Col()
//	TextFormatterCollection.SetStdFormatParamsLine2Col()
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//		column1Field               interface{}
//		   - This is the first of two text fields used to generate a
//		     2-Column Text Line. This parameter is an empty interface
//		     which must contain one of several specific data types.
//		     This empty interface type will be converted to a string
//		     and configured as the first text field in this 2-Column
//		     Text Line.
//
//		     Supported types which may be submitted through this empty
//		     interface parameter are listed as follows:
//		        time.Time (Converted using default format)
//		        string
//		        bool
//		        uint, uint8, uint16, uint32, uint64,
//		        int, int8, int16, int32, int64
//		        float32, float64
//		        *big.Int *big.Float
//		        fmt.Stringer (types that support this interface)
//		        TextInputParamFieldDateTimeDto
//		              (Converts date time to string)
//
//		     If the 'column1Field' is not convertible to one of the
//		     supported types, an error will be returned.
//
//		     If the converted string value for 'column1Field' is empty,
//		     it will be defaulted to a single white space character
//		     (" ").
//
//
//		column2Field               interface{}
//		   - This is the second of two text fields used to generate a
//		     2-Column Text Line. This parameter is an empty interface
//		     which must contain one of several specific data types.
//		     This empty interface type will be converted to a string
//		     and configured as the second text field in this 2-Column
//		     Text Line.
//
//		     Supported types which may be submitted through this empty
//		     interface parameter are listed as follows:
//		        time.Time (Converted using default format)
//		        string
//		        bool
//		        uint, uint8, uint16, uint32, uint64,
//		        int, int8, int16, int32, int64
//		        float32, float64
//		        *big.Int *big.Float
//		        fmt.Stringer (types that support this interface)
//		        TextInputParamFieldDateTimeDto
//		              (Converts date time to string)
//
//		     If the 'column2Field' is not convertible to one of the
//		     supported types, an error will be returned.
//
//		     If the converted string value for 'column2Field' is empty,
//		     it will be defaulted to a single white space character
//		     (" ").
//
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
//
//		     This empty interface must be convertible to one of the
//		     following types:
//
//		     1. nil - A nil value is valid and generates an empty
//		        collection of error prefix and error context
//		        information.
//
//		     2. string - A string containing error prefix information.
//
//		     3. []string A one-dimensional slice of strings containing
//		        error prefix information
//
//		     4. [][2]string A two-dimensional slice of strings
//		        containing error prefix and error context information.
//
//		     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		        from this object will be copied for use in error and
//		        informational messages.
//
//		     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		        Information from this object will be copied for use in
//		        error and informational messages.
//
//		     7. IBasicErrorPrefix - An interface to a method generating
//		        a two-dimensional slice of strings containing error
//		        prefix and error context information.
//
//		     If parameter 'errorPrefix' is NOT convertible to one of
//		     the valid types listed above, it will be considered
//		     invalid and trigger the return of an error.
//
//		     Types ErrPrefixDto and IBasicErrorPrefix are included in
//		     the 'errpref' software package,
//		     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
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
func (txtFmtCollection *TextFormatterCollection) AddLine2Col(
	column1Field interface{},
	column2Field interface{},
	errorPrefix interface{}) error {

	if txtFmtCollection.lock == nil {
		txtFmtCollection.lock = new(sync.Mutex)
	}

	txtFmtCollection.lock.Lock()

	defer txtFmtCollection.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFormatterCollection."+
			"AddLine2Col()",
		"")

	if err != nil {
		return err
	}

	var foundStdParams bool
	var stdLineColsFmt TextFmtParamsLineColumnsDto

	foundStdParams,
		stdLineColsFmt,
		err =
		new(textFormatterCollectionElectron).
			findStdTxtLineParameters(
				txtFmtCollection,
				2, // This is a two column Text Line
				ePrefix.XCpy(
					"Text Line 2-Column"))

	if err != nil {
		return err
	}

	if !foundStdParams {

		err = fmt.Errorf("%v - Error\n"+
			"Could NOT locate Standard Text Line Parameter Format\n"+
			"for a 2-Column Text Line.\n"+
			"Use the 'SetStdFormatParams()' method to configure a\n"+
			"new Standard Text Line Column Parameters Format for this"+
			"2-Column Text Line Type. ",
			ePrefix.String())

		return err
	}

	var column1FieldText string

	column1FieldText,
		err = new(textSpecificationAtom).
		convertParamEmptyInterfaceToString(
			column1Field,
			"column1Field",
			ePrefix.XCpy(
				"column1Field"))

	if err != nil {
		return err
	}

	if len(column1FieldText) == 0 {
		column1FieldText = " "
	}

	var column2FieldText string

	column2FieldText,
		err = new(textSpecificationAtom).
		convertParamEmptyInterfaceToString(
			column2Field,
			"column2Field",
			ePrefix.XCpy(
				"column2Field"))

	if err != nil {
		return err
	}

	if len(column2FieldText) == 0 {
		column2FieldText = " "
	}

	newTextLine1Cols := TextLineColumnsDto{
		FormatType: TxtFieldType.LineColumns(),
		TextFieldsContent: []TextFieldsContentDto{
			{
				TextFieldString:   column1FieldText,
				TextFieldDateTime: time.Time{},
				lock:              nil,
			},
			{
				TextFieldString:   column2FieldText,
				TextFieldDateTime: time.Time{},
				lock:              nil,
			},
		},
		FmtParameters: stdLineColsFmt,
		lock:          nil,
	}

	newTextFormatter := TextFormatterDto{
		FormatType:          TxtFieldType.LineColumns(),
		DateTime:            TextFieldDateTimeDto{},
		Filler:              TextFieldFillerDto{},
		Label:               TextFieldLabelDto{},
		Spacer:              TextFieldSpacerDto{},
		BlankLine:           TextLineBlankDto{},
		SolidLine:           TextLineSolidDto{},
		LineColumns:         newTextLine1Cols,
		LinesTimerStartStop: TextLineTimerStartStopDto{},
	}

	txtFmtCollection.fmtCollection =
		append(
			txtFmtCollection.fmtCollection,
			newTextFormatter)

	return err
}

// AddLineBlank - Adds Blank Text Lines to the Formatter
// Collection.
//
// Blank Lines typically consist of one or more new line
// characters ('\n') and nothing more. However, users have
// the option to provide alternate or custom line termination
// characters which will be applied instead.
//
// The number of blank lines created is controlled by input
// parameter 'numOfBlankLines'.
//
//	Example-1 :
//	 NumOfBlankLines = 3
//	 LineTerminator = ""
//	 Final Blank Line Text = "\n\n\n" // 3-new line characters
//
//	Example-2 :
//	 NumOfBlankLines = 2
//	 LineTerminator = "\n x \n"
//	 Final Blank Line Text = "\n x \n\n x \n"
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	numOfBlankLines            int
//	   - Specifies the number of blank lines which will be created.
//	     Essentially, this parameter controls the number of new
//	     line characters configured in the resulting text string.
//
//
//	lineTerminator             string
//	   - If this parameter is submitted as an empty string,
//	     the default new line terminator ('\n') will be applied as
//	     a line termination sequence.
//
//	     If this parameter is populated, this character sequence
//	     will be used as the Line Terminator for this text blank
//	     line.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	NONE
func (txtFmtCollection *TextFormatterCollection) AddLineBlank(
	numOfBlankLines int,
	lineTerminator string) {

	if txtFmtCollection.lock == nil {
		txtFmtCollection.lock = new(sync.Mutex)
	}

	txtFmtCollection.lock.Lock()

	defer txtFmtCollection.lock.Unlock()

	newTextFormatter := TextFormatterDto{
		FormatType: TxtFieldType.BlankLine(),
		DateTime:   TextFieldDateTimeDto{},
		Filler:     TextFieldFillerDto{},
		Label:      TextFieldLabelDto{},
		Spacer:     TextFieldSpacerDto{},
		BlankLine: TextLineBlankDto{
			FormatType:      TxtFieldType.BlankLine(),
			NumOfBlankLines: numOfBlankLines,
			LineTerminator:  lineTerminator,
			lock:            nil,
		},
		SolidLine:           TextLineSolidDto{},
		LineColumns:         TextLineColumnsDto{},
		LinesTimerStartStop: TextLineTimerStartStopDto{},
	}

	txtFmtCollection.fmtCollection =
		append(
			txtFmtCollection.fmtCollection,
			newTextFormatter)

	return
}

// AddLineBlankDto - Adds Blank Text Lines to the Formatter
// Collection generated from an input parameter of type
// TextLineBlankDto.
//
// Blank Lines typically consist of one or more new line
// characters ('\n') and nothing more. However, users have
// the option to provide alternate or custom line termination
// characters which will be applied instead.
//
// The number of blank lines created is controlled by input
// parameter 'numOfBlankLines'.
//
//	Example-1 :
//	 NumOfBlankLines = 3
//	 LineTerminator = ""
//	 Final Blank Line Text = "\n\n\n" // 3-new line characters
//
//	Example-2 :
//	 NumOfBlankLines = 2
//	 LineTerminator = "\n x \n"
//	 Final Blank Line Text = "\n x \n\n x \n"
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	txtLineBlankDto            TextLineBlankDto
//	   - An instance of TextLineBlankDto which contains all the
//	     necessary data parameters to produce one or more Blank
//	     Text Lines.
//
//	     The Text Line Blank Data Transfer Object is defined as
//	     follows:
//
//	     type TextLineBlankDto struct {
//
//	      FormatType                 TextFieldType
//	       Required. This enumeration value specifies the type of
//	       Text Format Operation to be performed.
//
//	       For the TextLineBlankDto Format Type, this value
//	       should always be set to: TxtFieldType.BlankLine()
//
//	      NumOfBlankLines            int
//	       The number of blank lines which will be generated by this
//	       method.
//
//	       If this value is less than one (+1) or greater than
//	       one-million, an error will be returned when attempting
//	       to format the Text Blank Line for text output.
//
//	      LineTerminator             string
//	      If this parameter is submitted as an empty string,
//	      the default new line terminator ('\n') will be applied.
//	      If this parameter is populated, this character sequence
//	      will be used as the Line Terminator for this Text Blank
//	      Line.
//	     }
//
// ----------------------------------------------------------------
//
// Return Values
//
//	NONE
func (txtFmtCollection *TextFormatterCollection) AddLineBlankDto(
	txtLineBlankDto TextLineBlankDto) {

	if txtFmtCollection.lock == nil {
		txtFmtCollection.lock = new(sync.Mutex)
	}

	txtFmtCollection.lock.Lock()

	defer txtFmtCollection.lock.Unlock()

	newTextFormatter := TextFormatterDto{
		FormatType:          TxtFieldType.BlankLine(),
		DateTime:            TextFieldDateTimeDto{},
		Filler:              TextFieldFillerDto{},
		Label:               TextFieldLabelDto{},
		Spacer:              TextFieldSpacerDto{},
		BlankLine:           txtLineBlankDto.CopyOut(),
		SolidLine:           TextLineSolidDto{},
		LineColumns:         TextLineColumnsDto{},
		LinesTimerStartStop: TextLineTimerStartStopDto{},
	}

	txtFmtCollection.fmtCollection =
		append(
			txtFmtCollection.fmtCollection,
			newTextFormatter)

	return
}

// AddLineMultiCol - Adds a single Text Line consisting of one or
// more columns to the Text Formatter Collection.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will assign previously configured (a.k.a. default)
// Format Parameters to this multi-column Text Line. The prior
// configuration of these Format Parameters is a requirement and
// errors will be generated if these Standard Format Parameters
// have not previously been created. The number of columns in the
// Standard Format Parameters instance MUST MATCH the number of
// elements in the input array parameter, 'textFields'.
//
// Example: If input parameter 'textFields' contains 3-items, then
// a 3-column Standard Format Parameters instance must be
// previously configured.
//
// To configure the standard parameters for Text Lines, call one of
// the following methods:
//
//	TextFormatterCollection.CfgLineMultiCol()
//	TextFormatterCollection.SetStdFormatParamsMultiCol()
//	TextFormatterCollection.SetStdFormatParamsManyCol()
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//		textFields                 []interface{}
//		   - An array of type empty interface ([]interface{}). Each
//		     element in this array will be converted to a string and
//		     configured as a text column in this Multi-Column Text
//		     Line.
//
//		     Supported types which may be submitted through this empty
//		     interface array are listed as follows:
//		        time.Time (Converted using default format)
//		        string
//		        bool
//		        uint, uint8, uint16, uint32, uint64,
//		        int, int8, int16, int32, int64
//		        float32, float64
//		        *big.Int *big.Float
//		        fmt.Stringer (types that support this interface)
//		        TextInputParamFieldDateTimeDto
//		              (Converts date time to string)
//
//		     If an element in the 'textFields' array is not convertible
//		     to one of the supported types, an error will be returned.
//
//		     If the converted string value for a 'textFields' element
//		     is empty (has a zero length), it will be defaulted to a
//		     single white space character (" ").
//
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
//
//		     This empty interface must be convertible to one of the
//		     following types:
//
//		     1. nil - A nil value is valid and generates an empty
//		        collection of error prefix and error context
//		        information.
//
//		     2. string - A string containing error prefix information.
//
//		     3. []string A one-dimensional slice of strings containing
//		        error prefix information
//
//		     4. [][2]string A two-dimensional slice of strings
//		        containing error prefix and error context information.
//
//		     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		        from this object will be copied for use in error and
//		        informational messages.
//
//		     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		        Information from this object will be copied for use in
//		        error and informational messages.
//
//		     7. IBasicErrorPrefix - An interface to a method generating
//		        a two-dimensional slice of strings containing error
//		        prefix and error context information.
//
//		     If parameter 'errorPrefix' is NOT convertible to one of
//		     the valid types listed above, it will be considered
//		     invalid and trigger the return of an error.
//
//		     Types ErrPrefixDto and IBasicErrorPrefix are included in
//		     the 'errpref' software package,
//		     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
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
func (txtFmtCollection *TextFormatterCollection) AddLineMultiCol(
	textFields []interface{},
	errorPrefix interface{}) error {

	txtFmtCollection.lock.Lock()

	defer txtFmtCollection.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFormatterCollection."+
			"AddLineMultiCol()",
		"")

	if err != nil {
		return err
	}

	lenTextFields := len(textFields)

	var foundStdParams bool
	var stdLineColsFmt TextFmtParamsLineColumnsDto

	foundStdParams,
		stdLineColsFmt,
		err =
		new(textFormatterCollectionElectron).
			findStdTxtLineParameters(
				txtFmtCollection,
				lenTextFields, // Must Match number of text fields
				ePrefix.XCpy(
					fmt.Sprintf("Text Line %v-Column",
						lenTextFields)))

	if err != nil {
		return err
	}

	if !foundStdParams {

		err = fmt.Errorf("%v - Error\n"+
			"Could NOT locate Standard Text Line Parameter Format\n"+
			"for a %v-Column Text Line.\n"+
			"Use the 'SetStdFormatParams()' method to configure a\n"+
			"new Standard Text Line Column Parameters Format for this"+
			"%v-Column Text Line. ",
			ePrefix.String(),
			lenTextFields,
			lenTextFields)

		return err
	}

	textFieldsContent := make([]TextFieldsContentDto, lenTextFields)

	for i := 0; i < lenTextFields; i++ {

		textFieldsContent[i].TextFieldString,
			err =
			new(textSpecificationAtom).
				convertParamEmptyInterfaceToString(
					textFields[i],
					fmt.Sprintf("textFields[%v]", i),
					ePrefix.XCpy(
						fmt.Sprintf("textFields[%v]", i)))

		if err != nil {
			return err
		}

		textFieldsContent[i].TextFieldDateTime = time.Time{}
		textFieldsContent[i].lock = nil
	}

	newTextLineCols := TextLineColumnsDto{
		FormatType:        TxtFieldType.LineColumns(),
		TextFieldsContent: textFieldsContent,
		FmtParameters:     stdLineColsFmt,
		lock:              nil,
	}

	newTextFormatter := TextFormatterDto{
		FormatType:          TxtFieldType.LineColumns(),
		DateTime:            TextFieldDateTimeDto{},
		Filler:              TextFieldFillerDto{},
		Label:               TextFieldLabelDto{},
		Spacer:              TextFieldSpacerDto{},
		BlankLine:           TextLineBlankDto{},
		SolidLine:           TextLineSolidDto{},
		LineColumns:         newTextLineCols,
		LinesTimerStartStop: TextLineTimerStartStopDto{},
	}

	txtFmtCollection.fmtCollection =
		append(
			txtFmtCollection.fmtCollection,
			newTextFormatter)

	return err
}

// AddLineManyCol - Adds a single Text Line consisting of one or
// more columns to the Text Formatter Collection. This method is
// configured as a variadic function with an input parameter that
// accepts a variable number of arguments ('textFields').
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will assign previously configured (a.k.a. default)
// Format Parameters to this Multiple Column Text Line. The prior
// configuration of these Format Parameters is a requirement and
// errors will be generated if these Standard Format Parameters
// have not previously been created. The number of columns in the
// Standard Format Parameters instance MUST MATCH the number of
// elements passed through input parameter, 'textFields'.
//
// Example: If input parameter 'textFields' contains 3-items, then
// a 3-column Standard Format Parameters instance must be
// previously configured.
//
// To configure the standard parameters for Text Lines, call one of
// the following methods:
//
//	TextFormatterCollection.CfgLineMultiCol()
//	TextFormatterCollection.SetStdFormatParamsMultiCol()
//	TextFormatterCollection.SetStdFormatParamsManyCol()
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
//
//		     This empty interface must be convertible to one of the
//		     following types:
//
//		     1. nil - A nil value is valid and generates an empty
//		        collection of error prefix and error context
//		        information.
//
//		     2. string - A string containing error prefix information.
//
//		     3. []string A one-dimensional slice of strings containing
//		        error prefix information
//
//		     4. [][2]string A two-dimensional slice of strings
//		        containing error prefix and error context information.
//
//		     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		        from this object will be copied for use in error and
//		        informational messages.
//
//		     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		        Information from this object will be copied for use in
//		        error and informational messages.
//
//		     7. IBasicErrorPrefix - An interface to a method generating
//		        a two-dimensional slice of strings containing error
//		        prefix and error context information.
//
//		     If parameter 'errorPrefix' is NOT convertible to one of
//		     the valid types listed above, it will be considered
//		     invalid and trigger the return of an error.
//
//		     Types ErrPrefixDto and IBasicErrorPrefix are included in
//		     the 'errpref' software package,
//		     "github.com/MikeAustin71/errpref".
//
//
//		textFields                 ...interface{}
//		   - This parameter accepts a variable number of arguments in
//		     the form of a series of empty interfaces (interface{}).
//		     Each argument passed through parameter will be converted
//		     to a string and configured as a text column in this
//		     Multiple Column Text Line.
//
//		     Supported types which may be submitted through the
//		     'textFields' parameter are listed as follows:
//		        time.Time (Converted using default format)
//		        string
//		        bool
//		        uint, uint8, uint16, uint32, uint64,
//		        int, int8, int16, int32, int64
//		        float32, float64
//		        *big.Int *big.Float
//		        fmt.Stringer (types that support this interface)
//		        TextInputParamFieldDateTimeDto
//		              (Converts date time to string)
//
//		     If any of the arguments passed through parameter
//		     'textFields' is not convertible to one of the supported
//		     types, an error will be returned.
//
//		     If the converted string value for a 'textFields' element
//		     is empty (has a zero length), it will be defaulted to a
//		     single white space character (" ").
//
// ----------------------------------------------------------------
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
func (txtFmtCollection *TextFormatterCollection) AddLineManyCol(
	errorPrefix interface{},
	textFields ...interface{}) error {

	txtFmtCollection.lock.Lock()

	defer txtFmtCollection.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFormatterCollection."+
			"AddLineMultiCol()",
		"")

	if err != nil {
		return err
	}

	textFieldsContent := make([]TextFieldsContentDto, 0)

	for idx, txtFieldVal := range textFields {

		txtFieldContent := TextFieldsContentDto{}

		txtFieldContent.TextFieldString,
			err =
			new(textSpecificationAtom).
				convertParamEmptyInterfaceToString(
					txtFieldVal,
					fmt.Sprintf("txtFieldVal[%v]", idx),
					ePrefix.XCpy(
						fmt.Sprintf("txtFieldVal[%v]", idx)))

		if err != nil {
			return err
		}

		txtFieldContent.TextFieldDateTime = time.Time{}
		txtFieldContent.lock = nil

		textFieldsContent = append(
			textFieldsContent,
			txtFieldContent)

	}

	lenTextFields := len(textFieldsContent)

	var foundStdParams bool
	var stdLineColsFmt TextFmtParamsLineColumnsDto

	foundStdParams,
		stdLineColsFmt,
		err =
		new(textFormatterCollectionElectron).
			findStdTxtLineParameters(
				txtFmtCollection,
				lenTextFields, // Must Match number of text fields
				ePrefix.XCpy(
					fmt.Sprintf("Text Line %v-Column",
						lenTextFields)))

	if err != nil {
		return err
	}

	if !foundStdParams {

		err = fmt.Errorf("%v - Error\n"+
			"Could NOT locate Standard Text Line Parameter Format\n"+
			"for a %v-Column Text Line.\n"+
			"Use the 'SetStdFormatParams()' method to configure a\n"+
			"new Standard Text Line Column Parameters Format for this"+
			"%v-Column Text Line. ",
			ePrefix.String(),
			lenTextFields,
			lenTextFields)

		return err
	}

	newTextLineCols := TextLineColumnsDto{
		FormatType:        TxtFieldType.LineColumns(),
		TextFieldsContent: textFieldsContent,
		FmtParameters:     stdLineColsFmt,
		lock:              nil,
	}

	newTextFormatter := TextFormatterDto{
		FormatType:          TxtFieldType.LineColumns(),
		DateTime:            TextFieldDateTimeDto{},
		Filler:              TextFieldFillerDto{},
		Label:               TextFieldLabelDto{},
		Spacer:              TextFieldSpacerDto{},
		BlankLine:           TextLineBlankDto{},
		SolidLine:           TextLineSolidDto{},
		LineColumns:         newTextLineCols,
		LinesTimerStartStop: TextLineTimerStartStopDto{},
	}

	txtFmtCollection.fmtCollection =
		append(
			txtFmtCollection.fmtCollection,
			newTextFormatter)

	return err
}

// AddLineSolid - Adds a Solid Text Line to the Formatter
// Collection.
//
// A solid line, as defined here, consists of a single character or
// multiple characters used in a repeating sequence to construct
// a solid line. Typically, solid lines consist of dashes ("---"),
// underscore characters ("____"), equal signs ("====="), asterisks
// ("*****") and other similar line break presentations. Multiple
// characters may be used to produce more complex line sequences
// ("--*--*--*"). The length of a solid line is specified by
// input parameter, 'SolidLineCharRepeatCount'.
//
// Therefore, the number of solid line characters produced is
// equal to:
//
//	(Number of  solid line characters 'solidLineChars')
//	     x 'solidLineCharRepeatCount'
//
// By default, all solid lines are terminated with a new line
// character ('\n'). This means that each instance of a solid line
// will be formatted as a single line of text. Users have the
// option to alter this behavior and control the content and
// application of line termination characters through two input
// parameters, 'turnLineTerminationOff' and 'lineTerminator'.
//
//	Example-1:
//	  SolidLineChars = "*"
//	  SolidLineCharRepeatCount = 5
//	  LeftMarginStr = ""
//	  RightMarginStr = ""
//	  TurnLineTerminationOff = false
//	  LineTerminator = ""
//	  Final Solid Line String = "*****\n"
//
//	Example-2:
//	  SolidLineChars = "*"
//	  SolidLineCharRepeatCount = 5
//	  LeftMarginStr = ""
//	  RightMarginStr = ""
//	  TurnLineTerminationOff = false
//	  LineTerminator = "\n\n"
//	  Final Solid Line String = "*****\n\n"
//
//	Example-3:
//	  SolidLineChars = "*"
//	  SolidLineCharRepeatCount = 5
//	  LeftMarginStr = ""
//	  RightMarginStr = ""
//	  TurnLineTerminationOff = true
//	  LineTerminator = "\n\n"
//	  Final Solid Line String = "*****"
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	leftMarginStr              string
//	   - A string containing the text characters to be positioned
//	     on the Left side of the Solid Line.
//
//	     If no Left margin is required, set this parameter to an
//	     empty string ("").
//
//	     Example A-1:
//	       LeftMarginStr = "xxx"
//	       SolidLineChars = "*"
//	       SolidLineCharRepeatCount = 5
//	       RightMarginStr = "" // Empty string
//	       Solid Line = "xxx*****"
//
//	     Example A-2:
//	       LeftMarginStr = ""
//	       SolidLineChars = "*"
//	       SolidLineCharRepeatCount = 5
//	       RightMarginStr = "" // Empty string
//	       Solid Line = "*****"
//
//	     If the 'LeftMarginStr' string length is greater than
//	     one-million (1,000,000), an error will be returned when
//	     attempting to create formatted text output.
//
//
//	solidLineChars             string
//	   - This string specifies the character or characters which
//	     will comprise the solid line string for screen display,
//	     file output or printing.
//
//	     Example B-1:
//	       solidLineChars = "*"
//	       solidLineCharsRepeatCount = 5
//	       Solid line = "*****"
//
//	     If this parameter is submitted as a zero length string, an
//	     error will be returned when attempting to create formatted
//	     text output.
//
//
//	solidLineCharRepeatCount   int
//	   - This integer value specifies the number of times that
//	     parameter 'solidLineChars' will be repeated in
//	     constructing the solid line.
//
//	     If this parameter is submitted with a value less than one
//	     (1), it will be automatically converted to one (1).
//
//	     Example C-1:
//	       solidLineChars = "*"
//	       solidLineCharsRepeatCount = 5
//	       Solid line = "*****"
//
//
//	rightMarginStr             string
//	   - A string containing the text characters to positioned on
//	     the Right side of the Solid Line.
//
//	     If no Right margin is required, set this parameter to an
//	     empty string.
//
//	     Example D-1:
//	       solidLineChars = "*"
//	       solidLineCharsRepeatCount = 5
//	       leftMarginStr = "" // Empty string
//	       RightMarginStr = "xxx"
//	       Solid line = "*****xxx"
//
//	     If the 'RightMarginStr' string length is greater than
//	     one-million (1,000,000), an error will be returned when
//	     attempting to created formatted text output.
//
//
//	turnLineTerminationOff     bool
//	   - If this parameter is set to 'true' no Line Termination
//	     Sequence will be applied for this text line (Example E-3).
//
//	           Example E-1:
//	             SolidLineChars = "*"
//	             SolidLineCharRepeatCount = 5
//	             LeftMarginStr = ""
//	             RightMarginStr = ""
//	             TurnLineTerminationOff = false
//	             LineTerminator = ""
//	             Final Solid Line String = "*****\n"
//
//	           Example E-2:
//	             SolidLineChars = "*"
//	             SolidLineCharRepeatCount = 5
//	             LeftMarginStr = ""
//	             RightMarginStr = ""
//	             TurnLineTerminationOff = false
//	             LineTerminator = "\n\n"
//	             Final Solid Line String = "*****\n\n"
//
//	           Example E-3:
//	             SolidLineChars = "*"
//	             SolidLineCharRepeatCount = 5
//	             LeftMarginStr = ""
//	             RightMarginStr = ""
//	             TurnLineTerminationOff = true
//	             LineTerminator = "\n\n"
//	             Final Solid Line String = "*****"
//
//
//	lineTerminator             string
//	   - This string holds the character or characters which
//	     will be used to terminate the formatted line of text
//	     output.
//
//	     The most common usage sets this string to a new line
//	     character ("\n").
//
//	     If 'LineTerminator' is configured as an empty string
//	     (string length zero), a single new line character ('\n')
//	     will be automatically applied to produce line termination.
//
//	     LineTerminator works in conjunction with member variable
//	     'TurnLineTerminationOff'. 'TurnLineTerminationOff'
//	     controls the application of a line terminator. Setting
//	     'TurnLineTerminationOff' to 'true' means that NO line
//	     terminator will be applied to this instance of Text Solid
//	     Line (Example F-3).
//
//	           Example F-1:
//	             SolidLineChars = "*"
//	             SolidLineCharRepeatCount = 5
//	             LeftMarginStr = ""
//	             RightMarginStr = ""
//	             TurnLineTerminationOff = false
//	             LineTerminator = ""
//	             Final Solid Line String = "*****\n"
//
//	           Example F-2:
//	             SolidLineChars = "*"
//	             SolidLineCharRepeatCount = 5
//	             LeftMarginStr = ""
//	             RightMarginStr = ""
//	             TurnLineTerminationOff = false
//	             LineTerminator = "\n\n"
//	             Final Solid Line String = "*****\n\n"
//
//	           Example F-3:
//	             SolidLineChars = "*"
//	             SolidLineCharRepeatCount = 5
//	             LeftMarginStr = ""
//	             RightMarginStr = ""
//	             TurnLineTerminationOff = true
//	             LineTerminator = "\n\n"
//	             Final Solid Line String = "*****"
//
//
//	maxLineLength              int
//	   - The maximum length of the line on which the solid line
//	     text characters will be presented.
//
//	     Set this parameter to minus one (-1) to specify an
//	     unlimited line length for this text line.
//
//	     If the value of 'maxLineLength' is less than one (1),
//	     it will be automatically converted to minus one (-1).
//
//	     'maxLineLength' is used in conjunction with parameter
//	     'turnAutoLineLengthBreaksOn' to automatically place text
//	     on separate text lines when that text exceeds the maximum
//	     text line length ('maxLineLength'). Therefore, paramter
//	     'turnAutoLineLengthBreaksOn' controls whether automatic
//	     line breaks using 'maxLineLength' will be applied.
//
//
//	turnAutoLineLengthBreaksOn bool
//	   - This parameter controls whether text lines which exceed
//	     the maximum line length ('maxLineLength') are positioned
//	     on the following line as a separate line of text.
//
//	     To apply automatic line breaking at the maximum line
//	     length ('maxLineLength'), set the value of this parameter
//	     to 'true'.
//
//	     When this parameter is set to 'true', text fields which
//	     extend beyond the maximum line length ('maxLineLength')
//	     will be formatted as a separate line of text on the
//	     following line.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	NONE
func (txtFmtCollection *TextFormatterCollection) AddLineSolid(
	leftMarginStr string,
	solidLineChars string,
	solidLineCharRepeatCount int,
	rightMarginStr string,
	turnLineTerminationOff bool,
	lineTerminator string,
	maxLineLength int,
	turnAutoLineLengthBreaksOn bool) {

	if txtFmtCollection.lock == nil {
		txtFmtCollection.lock = new(sync.Mutex)
	}

	txtFmtCollection.lock.Lock()

	defer txtFmtCollection.lock.Unlock()

	if maxLineLength < 1 {
		maxLineLength = -1
	}

	if solidLineCharRepeatCount < 1 {
		solidLineCharRepeatCount = 1
	}

	newTextFormatter := TextFormatterDto{
		FormatType: TxtFieldType.SolidLine(),
		DateTime:   TextFieldDateTimeDto{},
		Filler:     TextFieldFillerDto{},
		Label:      TextFieldLabelDto{},
		Spacer:     TextFieldSpacerDto{},
		BlankLine:  TextLineBlankDto{},
		SolidLine: TextLineSolidDto{
			FormatType:                 TxtFieldType.SolidLine(),
			LeftMarginStr:              leftMarginStr,
			SolidLineChars:             solidLineChars,
			SolidLineCharRepeatCount:   solidLineCharRepeatCount,
			RightMarginStr:             rightMarginStr,
			TurnLineTerminationOff:     turnLineTerminationOff,
			LineTerminator:             lineTerminator,
			MaxLineLength:              maxLineLength,
			TurnAutoLineLengthBreaksOn: turnAutoLineLengthBreaksOn,
			lock:                       nil,
		},

		LineColumns:         TextLineColumnsDto{},
		LinesTimerStartStop: TextLineTimerStartStopDto{},
	}

	txtFmtCollection.fmtCollection =
		append(
			txtFmtCollection.fmtCollection,
			newTextFormatter)

	return
}

// AddLineSolidDto - Adds a Solid Text Line to the Formatter
// Collection generated from an input parameter of type
// TextFormatterCollection.
//
// A solid line, as defined here, consists of a single character or
// multiple characters used in a repeating sequence to construct
// a solid line. Typically, solid lines consist of dashes ("---"),
// underscore characters ("____"), equal signs ("====="), asterisks
// ("*****") and other similar line break presentations. Multiple
// characters may be used to produce more complex line sequences
// ("--*--*--*"). The length of a solid line is specified by the
// Text Line Solid Data Transfer Object member variable,
// 'SolidLineCharRepeatCount'.
//
// Therefore, the number of solid line characters produced is
// equal to:
//
//	(Number of  solid line characters 'solidLineChars')
//	     x 'solidLineCharRepeatCount'
//
// By default, all solid lines are terminated with a new line
// character ('\n'). This means that each instance of a solid line
// will be formatted as a single line of text. Users have the
// option to alter this behavior and control the content and
// application of line termination characters through two input
// parameters, 'turnLineTerminationOff' and 'lineTerminator'.
//
//	Example-1:
//	  SolidLineChars = "*"
//	  SolidLineCharRepeatCount = 5
//	  LeftMarginStr = ""
//	  RightMarginStr = ""
//	  TurnLineTerminationOff = false
//	  LineTerminator = ""
//	  Final Solid Line String = "*****\n"
//
//	Example-2:
//	  SolidLineChars = "*"
//	  SolidLineCharRepeatCount = 5
//	  LeftMarginStr = ""
//	  RightMarginStr = ""
//	  TurnLineTerminationOff = false
//	  LineTerminator = "\n\n"
//	  Final Solid Line String = "*****\n\n"
//
//	Example-3:
//	  SolidLineChars = "*"
//	  SolidLineCharRepeatCount = 5
//	  LeftMarginStr = ""
//	  RightMarginStr = ""
//	  TurnLineTerminationOff = true
//	  LineTerminator = "\n\n"
//	  Final Solid Line String = "*****"
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	txtLineSolidDto            TextLineSolidDto
//	   - An instance of TextLineSolidDto which contains all the
//	     data parameters required to produce a Text Solid Line.
//
//	     The Text Line Solid Data Transfer Object is defined as
//	     follows:
//	     type TextLineSolidDto struct {
//
//	      FormatType                   TextFieldType
//	       Required. This enumeration value specifies the type of
//	       Text Format Operation to be performed.
//
//	       For TextLineSolidDto 'FormatType' this parameter should
//	       be set to: TxtFieldType.SolidLine()
//
//	      LeftMarginStr                string
//	       A string containing the text characters to be positioned
//	       on the Left side of the Solid Line.
//
//	       If no Left margin is required, set this parameter to an
//	       empty string.
//
//	       Example:
//	         LeftMarginStr = "xxx"
//	         SolidLineChars = "*"
//	         SolidLineCharRepeatCount = 5
//	         RightMarginStr = "" // Empty string
//	         Solid Line = "xxx*****"
//
//	       If the 'LeftMarginStr' string length is greater than
//	       one-million (1,000,000), an error will be returned when
//	       attempting to create formatted text output.
//
//	      SolidLineChars               string
//	       This string specifies the character or characters which
//	       will comprise the solid line string for screen display,
//	       file output or printing.
//
//	       Example:
//	         solidLineChars = "*"
//	         solidLineCharsRepeatCount = 5
//	         Solid line = "*****"
//
//	       If this parameter is submitted as a zero length string,
//	       an error will be returned when attempting to create
//	       formatted.
//
//	      SolidLineCharRepeatCount     int
//	       This integer value specifies the number of times that
//	       parameter 'solidLineChars' will be repeated in
//	       constructing the solid line.
//
//	       If this parameter is submitted with a value less than
//	       one (1), it will be automatically converted to one (1).
//
//	       Example:
//	         solidLineChars = "*"
//	         solidLineCharsRepeatCount = 5
//	         Solid line = "*****"
//
//	      RightMarginStr               string
//	       A string containing the text characters to positioned on
//	       the Right side of the Solid Line.
//
//	       If no Right margin is required, set this parameter to an
//	       empty string.
//
//	       Example:
//	         solidLineChars = "*"
//	         solidLineCharsRepeatCount = 5
//	         leftMarginStr = "" // Empty string
//	         RightMarginStr = "xxx"
//	         Solid line = "*****xxx"
//
//	       If the 'RightMarginStr' string length is greater than
//	       one-million (1,000,000), an error will be returned.
//
//	      TurnLineTerminationOff     bool
//	       If this parameter is set to 'true' no Line Termination
//	       Sequence will be applied for this text line (Example-3).
//
//	            Example-1:
//	              SolidLineChars = "*"
//	              SolidLineCharRepeatCount = 5
//	              LeftMarginStr = ""
//	              RightMarginStr = ""
//	              TurnLineTerminationOff = false
//	              LineTerminator = ""
//	              Final Solid Line String = "*****\n"
//
//	            Example-2:
//	              SolidLineChars = "*"
//	              SolidLineCharRepeatCount = 5
//	              LeftMarginStr = ""
//	              RightMarginStr = ""
//	              TurnLineTerminationOff = false
//	              LineTerminator = "\n\n"
//	              Final Solid Line String = "*****\n\n"
//
//	            Example-3:
//	              SolidLineChars = "*"
//	              SolidLineCharRepeatCount = 5
//	              LeftMarginStr = ""
//	              RightMarginStr = ""
//	              TurnLineTerminationOff = true
//	              LineTerminator = "\n\n"
//	              Final Solid Line String = "*****"
//
//	      LineTerminator               string
//	       This string holds the character or characters which
//	       will be used to terminate the formatted line of text
//	       output.
//
//	       The most common usage sets this string to a new line
//	       character ("\n").
//
//	       If 'LineTerminator' is configured as an empty string
//	       (string length zero), a single new line character ('\n')
//	       will be automatically applied to produce line termination.
//
//	       LineTerminator works in conjunction with member variable
//	       'TurnLineTerminationOff'. 'TurnLineTerminationOff'
//	       controls the application of a line terminator. Setting
//	       'TurnLineTerminationOff' to 'true' means that NO line
//	       terminator will be applied to this instance of Text Solid
//	       Line.
//
//	             Example-1:
//	               SolidLineChars = "*"
//	               SolidLineCharRepeatCount = 5
//	               LeftMarginStr = ""
//	               RightMarginStr = ""
//	               TurnLineTerminationOff = false
//	               LineTerminator = ""
//	               Final Solid Line String = "*****\n"
//
//	             Example-2:
//	               SolidLineChars = "*"
//	               SolidLineCharRepeatCount = 5
//	               LeftMarginStr = ""
//	               RightMarginStr = ""
//	               TurnLineTerminationOff = false
//	               LineTerminator = "\n\n"
//	               Final Solid Line String = "*****\n\n"
//
//	             Example-3:
//	               SolidLineChars = "*"
//	               SolidLineCharRepeatCount = 5
//	               LeftMarginStr = ""
//	               RightMarginStr = ""
//	               TurnLineTerminationOff = true
//	               LineTerminator = "\n\n"
//	               Final Solid Line String = "*****"
//
//	      MaxLineLength               int
//	       The maximum length of the line on which the solid line
//	       text characters will be presented.
//
//	       Set this parameter to minus one (-1) to specify an
//	       unlimited line length for this text line.
//
//	       If the value of 'MaxLineLength' is less than one (1),
//	       it will be automatically converted to minus one (-1).
//
//	       'MaxLineLength' is used in conjunction with parameter
//	       'TurnAutoLineLengthBreaksOn' to automatically place text
//	       fields on separate text lines when that text exceeds the
//	       maximum text line length ('MaxLineLength'). Therefore,
//	       paramter 'TurnAutoLineLengthBreaksOn' controls whether
//	       automatic line breaks using 'MaxLineLength' will be
//	       applied.
//
//	       TurnAutoLineLengthBreaksOn  bool
//	       This parameter controls whether text lines which exceed
//	       the maximum line length ('MaxLineLength') are
//	       positioned on the following line as a separate line of
//	       text.
//
//	       To apply automatic line breaking at the maximum line
//	       length, set the value of this parameter to 'true'.
//
//	       When this parameter is set to 'true', text fields which
//	       extend beyond the maximum line length 'MaxLineLength'
//	       will be formatted as a separate line of text on the
//	       following line.
//
//	     }
//
// ----------------------------------------------------------------
//
// Return Values
//
//	NONE
func (txtFmtCollection *TextFormatterCollection) AddLineSolidDto(
	txtLineSolidDto TextLineSolidDto) {

	if txtFmtCollection.lock == nil {
		txtFmtCollection.lock = new(sync.Mutex)
	}

	txtFmtCollection.lock.Lock()

	defer txtFmtCollection.lock.Unlock()

	if txtLineSolidDto.MaxLineLength < 1 {
		txtLineSolidDto.MaxLineLength = -1
	}

	if txtLineSolidDto.SolidLineCharRepeatCount < 1 {
		txtLineSolidDto.SolidLineCharRepeatCount = 1
	}

	newTextFormatter := TextFormatterDto{
		FormatType:          TxtFieldType.SolidLine(),
		DateTime:            TextFieldDateTimeDto{},
		Filler:              TextFieldFillerDto{},
		Label:               TextFieldLabelDto{},
		Spacer:              TextFieldSpacerDto{},
		BlankLine:           TextLineBlankDto{},
		SolidLine:           txtLineSolidDto.CopyOut(),
		LineColumns:         TextLineColumnsDto{},
		LinesTimerStartStop: TextLineTimerStartStopDto{},
	}

	txtFmtCollection.fmtCollection =
		append(
			txtFmtCollection.fmtCollection,
			newTextFormatter)

	return
}

// AddLineTimerStartStop - Adds Timer Start Stop Text Lines to the
// Formatter Collection.
//
// Text Line Timer Start Stop type records, computes and formats an
// elapsed time. This format process requires user input specifying
// a start time and ending time.
//
// The final formatted output string is composed of four lines of
// text for output to screen display, file output or printing.
//
// The first line of text shows the Starting Time. The second line
// shows the Ending Time. The third line displays the time duration
// or the difference between starting time and ending time. The
// fourth line displays the total elapsed time in nanoseconds.
//
// The third line contains Time duration, or elapsed time, and is
// broken down by days, hours, minutes, seconds, microseconds,
// milliseconds and nanoseconds. The display has a variable line
// length and will begin with the first category containing valid
// time duration data.
//
// Sample Output
//
//	  Start Time: 2021-08-13 03:19:32.462108100 -0500 CDT
//	    End Time: 2021-08-13 03:19:32.462163100 -0500 CDT
//	Elapsed Time: 55 Microseconds 0 Nanoseconds
//	              Total Elapsed Nanoseconds: 55,000
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	leftMarginStr              string
//	   - The character or characters which will comprise the left
//	     margin of the text label.
//
//	     If the string length of 'leftMarginStr' plus
//	     'rightMarginStr' plus the text label field length
//	     ('testLabelFieldLength') exceeds the maximum length of
//	     55-characters, an error will be generated when
//	     attempting to create formatted text output.
//
//	     If leftMarginStr is submitted as an empty or zero length
//	     string, no error will be triggered and the left margin
//	     will simply be omitted.
//
//
//	startTimeLabel             string
//	   - A string containing the text characters constituting the
//	     starting time text label.
//
//	     If this string is submitted as a zero length or empty
//	     string, 'startTimeLabel' will be assigned a default value
//	     of "Start Time".
//
//	     If the string length of 'leftMarginStr' plus
//	     'rightMarginStr' plus the text label field length
//	     ('testLabelFieldLength') exceeds the maximum length of
//	     55-characters, an error will be generated when
//	     attempting to create formatted text output.
//
//
//	startTime                  time.Time
//	   - A time value which will be used in conjunction with
//	     'endTime' parameter to compute the time duration or
//	     elapsed for the timer event.
//
//	     If this parameter is submitted as a zero time value,
//	     'startTime' will be defaulted to value of July 4, 1776
//	     9:30AM UTC.
//
//
//	endTimeLabel               string
//	   - A string containing the text characters constituting the
//	     ending time text label.
//
//	     If this string is submitted as a zero length or empty
//	     string, 'endTimeLabel' will be assigned a default value of
//	     "End Time".
//
//	     If the string length of 'leftMarginStr' plus
//	     'rightMarginStr' plus the text label field length
//	     ('testLabelFieldLength') exceeds the maximum length of
//	     55-characters, an error will be generated when
//	     attempting to create formatted text output.
//
//
//	endTime                    time.Time
//	   - A time value which will be used in conjunction with
//	     'startTime' parameter to compute the time duration or
//	     elapsed for the timer event.
//
//	     If this parameter is submitted as a zero time value,
//	     'endTime' will be defaulted to value of 'startTime'.
//
//
//	timeFormat                 string
//	   - This string holds the time format parameters used to
//	     format starting time and ending time values for text
//	     display and output.
//
//	     If this parameter is submitted as an empty string,
//	     parameter 'timeFormat' will be assigned a default value
//	     of "2006-01-02 15:04:05.000000000 -0700 MST"
//
//
//	timeDurationLabel          string
//	   - The text label used to describe the time duration or
//	     elapsed time computed from the 'startTime' and 'endTime'
//	     parameters.
//
//	     If this string is submitted as a zero length or empty
//	     string, 'timeDurationLabel' will be assigned a default
//	     value of "Elapsed Time".
//
//	     If the string length of 'leftMarginStr' plus
//	     'rightMarginStr' plus the text label field length
//	     ('testLabelFieldLength') exceeds the maximum length of
//	     55-characters, an error will be generated when
//	     attempting to create formatted text output.
//
//
//	testLabelFieldLength          int
//	   - A user entered value which defines the length of the text
//	     field used by all three text labels, 'startTimeLabel',
//	     'endTimeLabel' and 'timeDurationLabel'.
//
//	     The length of the text fields which will be used to
//	     position and display the three text labels provided by
//	     input parameters 'startTimeLabel', 'endTimeLabel' and
//	     'timeDurationLabel'.
//
//	     If 'testLabelFieldLength' is less than the length of the
//	     longest text label it will be defaulted to the length
//	     of the longest text label ('startTimeLabel',
//	     'endTimeLabel' or 'timeDurationLabel').
//
//	     If the string length of 'leftMarginStr' plus
//	     'rightMarginStr' plus the text label field length
//	     ('testLabelFieldLength') exceeds the maximum length of
//	     55-characters, an error will be generated when
//	     attempting to create formatted text output.
//
//
//	textLabelJustification         TextJustify
//	   - An enumeration which specifies the justification of the
//	     three text labels 'startTimeLabel', 'endTimeLabel' and
//	     'timeDurationLabel' within the field length specified by
//	     'testLabelFieldLength'.
//
//	     Label justification must be equal to one of these three
//	     valid values:
//	         TextJustify(0).Left()
//	         TextJustify(0).Right()
//	         TextJustify(0).Center()
//
//	     The abbreviated text justification enumeration syntax can
//	     also be used:
//
//	         TxtJustify.Left()
//	         TxtJustify.Right()
//	         TxtJustify.Center()
//
//
//	rightMarginStr      string
//	   - This string contains the character or characters which
//	     will be used to separate the text labels
//	     ('startTimeLabel', 'endTimeLabel' and 'timeDurationLabel')
//	     from the output or data values displayed on the same line.
//	     Example:
//	      Start Time[right margin chars]2010-01-02 15:04:05.000000000 -0700 MST
//
//	     Often this parameter is set to a single white space
//	     character (" ") or a colon plus white space character,
//	     (": ").
//
//	     If this string is submitted as a zero length or empty
//	     string, 'rightMarginStr' will be assigned a
//	     default value of  ": ". Example Output:
//	      Start Time: 2010-01-02 15:04:05.000000000 -0700 MST
//
//	     If the string length of 'leftMarginStr' plus
//	     'rightMarginStr' plus the text label field length
//	     ('testLabelFieldLength') exceeds the maximum length of
//	     55-characters, an error will be generated when
//	     attempting to create formatted text output.
//
//	     If leftMarginStr is submitted as a zero length
//	     string, no error will be triggered. In this case, the left
//	     margin will simply be omitted.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	NONE
func (txtFmtCollection *TextFormatterCollection) AddLineTimerStartStop(
	leftMarginStr string,
	startTimeLabel string,
	startTime time.Time,
	endTimeLabel string,
	endTime time.Time,
	timeFormat string,
	timeDurationLabel string,
	textLabelFieldLength int,
	textLabelJustification TextJustify,
	rightMarginStr string) {

	if txtFmtCollection.lock == nil {
		txtFmtCollection.lock = new(sync.Mutex)
	}

	txtFmtCollection.lock.Lock()

	defer txtFmtCollection.lock.Unlock()

	timerStartStopDto := TextLineTimerStartStopDto{
		FormatType:             TxtFieldType.TimerStartStop(),
		LeftMarginStr:          leftMarginStr,
		StartTimeLabel:         startTimeLabel,
		StartTime:              startTime,
		EndTimeLabel:           endTimeLabel,
		EndTime:                endTime,
		TimeFormat:             timeFormat,
		TimeDurationLabel:      timeDurationLabel,
		TextLabelFieldLength:   textLabelFieldLength,
		TextLabelJustification: textLabelJustification,
		RightMarginStr:         rightMarginStr,
		lock:                   nil,
	}

	newTextFormatter := TextFormatterDto{
		FormatType:          TxtFieldType.TimerStartStop(),
		DateTime:            TextFieldDateTimeDto{},
		Filler:              TextFieldFillerDto{},
		Label:               TextFieldLabelDto{},
		Spacer:              TextFieldSpacerDto{},
		BlankLine:           TextLineBlankDto{},
		SolidLine:           TextLineSolidDto{},
		LineColumns:         TextLineColumnsDto{},
		LinesTimerStartStop: timerStartStopDto,
		lock:                nil,
	}

	txtFmtCollection.fmtCollection =
		append(
			txtFmtCollection.fmtCollection,
			newTextFormatter)

	return
}

// AddLineTimerStartStopDto - Adds Timer Start Stop Text Lines to
// the Formatter Collection generated from an input parameter of
// type TextLineTimerStartStopDto.
//
// Text Line Timer Start Stop type records, computes and formats an
// elapsed time. This format process requires user input specifying
// a start time and ending time.
//
// The final formatted output string is composed of four lines of
// text for output to screen display, file output or printing.
//
// The first line of text shows the Starting Time. The second line
// shows the Ending Time. The third line displays the time duration
// or the difference between starting time and ending time. The
// fourth line displays the total elapsed time in nanoseconds.
//
// The third line contains Time duration, or elapsed time, and is
// broken down by days, hours, minutes, seconds, microseconds,
// milliseconds and nanoseconds. The display has a variable line
// length and will begin with the first category containing valid
// time duration data.
//
// Sample Output
//
//	  Start Time: 2021-08-13 03:19:32.462108100 -0500 CDT
//	    End Time: 2021-08-13 03:19:32.462163100 -0500 CDT
//	Elapsed Time: 55 Microseconds 0 Nanoseconds
//	              Total Elapsed Nanoseconds: 55,000
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	timerStartStopDto          TextLineTimerStartStopDto
//	   - An instance of TextLineTimerStartStopDto which contains
//	     all the data parameters required to produce Timer Start
//	     Stop Text Lines describing a timer event.
//
//	     The Text Line Timer Start Stop Data Transfer Object is
//	     defined as follows:
//
//	     type TextLineTimerStartStopDto struct {
//
//	      FormatType                   TextFieldType
//	       Required. This enumeration value specifies the type of
//	       Text Format Operation to be performed.
//
//	       For TextLineTimerStartStopDto Format Type, this value
//	       should always be set to:
//	       TxtFieldType.TimerStartStop()
//
//	      LeftMarginStr                string
//	       The contents of the string will be used as the left
//	       margin for the timer event lines.
//
//	       If no left margin is required, set 'LeftMarginStr' to a
//	       zero length or empty string, and no left margin will be
//	       created.
//
//	      StartTimeLabel               string
//	       A string containing the text characters constituting the
//	       starting time text label.
//
//	       If this parameter is submitted as a zero length or empty
//	       rune array, or if the rune array contains invalid zero
//	       rune values, 'StartTimeLabel' will be assigned a default
//	       value of "Start Time".
//
//	       If the length in characters of 'LeftMarginStr' plus
//	       'RightMarginStr' plus the text label field length
//	       ('TextLabelFieldLength') exceeds the maximum length
//	       of 55-characters, an error will be generated when
//	       attempting to create formatted text output.
//
//	      StartTime                    time.Time
//	       A time value which will be used in conjunction with
//	       the 'EndTime' parameter to compute the time duration
//	       or elapsed time for the timer event.
//
//	       If this parameter is submitted as a zero time value,
//	       'startTime' will be defaulted to value of July 4, 1776
//	       9:30AM UTC.
//
//	      EndTimeLabel                 string
//	       A string containing the text characters constituting the
//	       ending time text label.
//
//	       If this parameter is submitted as a zero length or empty
//	       string, 'EndTimeLabel' will be assigned a default
//	       value of "End Time".
//
//	       If the length in characters of 'LeftMarginStr' plus
//	       'RightMarginStr' plus the text label field length
//	       ('TextLabelFieldLength') exceeds the maximum length
//	       of 55-characters, an error will be generated when
//	       attempting to create formatted text output.
//
//	      EndTime                      time.Time
//	       A time value which will be used in conjunction with
//	       the 'StartTime' parameter to compute the time duration
//	       or elapsed time for the timer event.
//
//	       If this parameter is submitted as a zero time value,
//	       'EndTime' will be defaulted to value of 'StartTime'.
//
//	      TimeFormat                   string
//	       This string holds the time format parameters used to
//	       format starting time and ending time values for text
//	       output.
//
//	       If this parameter is submitted as an empty string,
//	       parameter 'TimeFormat' will be assigned a default
//	       value of:
//	        "2006-01-02 15:04:05.000000000 -0700 MST"
//
//	      TimeDurationLabel            string
//	       The text label used to describe the time duration or
//	       elapsed time computed from the 'startTime' and 'endTime'
//	       parameters.
//
//	       If this string is submitted as a zero length or empty
//	       string, 'TimeDurationLabel' will be assigned a default
//	       value of "Elapsed Time".
//
//	       If the string length of 'LeftMarginStr' plus
//	       'RightMarginStr' plus the text label field length
//	       ('TextLabelFieldLength') exceeds the maximum length
//	       of 55-characters, an error will be generated when
//	       attempting to create formatted text output.
//
//	      TextLabelFieldLength         int
//	       A user entered value which defines the length of the
//	       text field used by all three text labels,
//	       'StartTimeLabel', 'EndTimeLabel' and
//	       'TimeDurationLabel'.
//
//	       This text length value will be used to position and
//	       display the three text labels provided by
//	       input parameters 'StartTimeLabel', 'EndTimeLabel' and
//	       'TimeDurationLabel'.
//
//	       If 'TextLabelFieldLength' is less than the length of
//	       the longest text label it will be defaulted to the
//	       length of the longest text label ('StartTimeLabel',
//	       'EndTimeLabel' or 'TimeDurationLabel').
//
//	       If the string length of 'LeftMarginStr' plus
//	       'RightMarginStr' plus the text label field length
//	       ('TextLabelFieldLength') exceeds the maximum length
//	       of 55-characters, an error will be generated when
//	       attempting to create formatted text output.
//
//	      TextLabelJustification       TextJustify
//	       An enumeration which specifies the text justification of
//	       the three text labels 'StartTimeLabel', 'EndTimeLabel'
//	       and 'TimeDurationLabel' within the field length
//	       specified by 'TextLabelFieldLength'.
//
//	       Label justification must be equal to one of these three
//	       valid values:
//	         TextJustify(0).Left()
//	         TextJustify(0).Right()
//	         TextJustify(0).Center()
//
//	       The abbreviated text justification enumeration syntax
//	       can also be used:
//
//	         TxtJustify.Left()
//	         TxtJustify.Right()
//	         TxtJustify.Center()
//
//	      RightMarginStr               string
//	       This string contains the character or characters which
//	       will be used to separate the text labels
//	       ('startTimeLabel', 'endTimeLabel' and
//	       'timeDurationLabel') from the output or  data values
//	       displayed on the same line.
//	        Example:
//	        Start Time[RightMarginStr]2010-01-02 15:04:05.000000000 -0700 MST
//
//	       Often this parameter is set to a single white space
//	       character (" ") or a colon plus white space character,
//	       (": ").
//
//	       If this string is submitted as a zero length or empty
//	       string, 'RightMarginStr' will be assigned a default
//	       value of  ": ".
//	        Example Output:
//	        Start Time: 2010-01-02 15:04:05.000000000 -0700 MST
//
//	       If the string length of 'LeftMarginStr' plus
//	       'RightMarginStr' plus the text label field length
//	       ('TextLabelFieldLength') exceeds the maximum length
//	       of 55-characters, an error will be generated when
//	       attempting to create formatted text output.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	NONE
func (txtFmtCollection *TextFormatterCollection) AddLineTimerStartStopDto(
	timerStartStopDto TextLineTimerStartStopDto) {

	if txtFmtCollection.lock == nil {
		txtFmtCollection.lock = new(sync.Mutex)
	}

	txtFmtCollection.lock.Lock()

	defer txtFmtCollection.lock.Unlock()

	newTextFormatter := TextFormatterDto{
		FormatType:          TxtFieldType.TimerStartStop(),
		DateTime:            TextFieldDateTimeDto{},
		Filler:              TextFieldFillerDto{},
		Label:               TextFieldLabelDto{},
		Spacer:              TextFieldSpacerDto{},
		BlankLine:           TextLineBlankDto{},
		SolidLine:           TextLineSolidDto{},
		LineColumns:         TextLineColumnsDto{},
		LinesTimerStartStop: timerStartStopDto,
		lock:                nil,
	}

	txtFmtCollection.fmtCollection =
		append(
			txtFmtCollection.fmtCollection,
			newTextFormatter)

	return
}

// BuildText - Generates the formatted text using the contents
// of the Text Formatter Collection maintained by the current
// instance of TextFormatterCollection.
//
// If the Text Formatter Collection is empty (contains zero
// elements), an error will be returned.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//		strBuilder                 *strings.Builder
//		   - A pointer to an instance of *strings.Builder. The
//		     formatted text characters produced by this method will be
//		     written to this instance of strings.Builder.
//
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
//
//		     This empty interface must be convertible to one of the
//		     following types:
//
//		     1. nil - A nil value is valid and generates an empty
//		        collection of error prefix and error context
//		        information.
//
//		     2. string - A string containing error prefix information.
//
//		     3. []string A one-dimensional slice of strings containing
//		        error prefix information
//
//		     4. [][2]string A two-dimensional slice of strings
//		        containing error prefix and error context information.
//
//		     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		        from this object will be copied for use in error and
//		        informational messages.
//
//		     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		        Information from this object will be copied for use in
//		        error and informational messages.
//
//		     7. IBasicErrorPrefix - An interface to a method generating
//		        a two-dimensional slice of strings containing error
//		        prefix and error context information.
//
//		     If parameter 'errorPrefix' is NOT convertible to one of
//		     the valid types listed above, it will be considered
//		     invalid and trigger the return of an error.
//
//		     Types ErrPrefixDto and IBasicErrorPrefix are included in
//		     the 'errpref' software package,
//		     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
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
func (txtFmtCollection *TextFormatterCollection) BuildText(
	strBuilder *strings.Builder,
	errorPrefix interface{}) error {

	if txtFmtCollection.lock == nil {
		txtFmtCollection.lock = new(sync.Mutex)
	}

	txtFmtCollection.lock.Lock()

	defer txtFmtCollection.lock.Unlock()
	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFormatterCollection."+
			"BuildText()",
		"")

	if err != nil {
		return err
	}

	if strBuilder == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'strBuilder' is invalid!\n"+
			"'strBuilder' is a 'nil' pointer.\n",
			ePrefix.String())

		return err

	}

	txtBuilder := TextStrBuilder{}

	err = txtBuilder.BuildText(
		strBuilder,
		txtFmtCollection,
		ePrefix.XCpy(
			"txtFmtCollection"))

	return err
}

// CfgLine1Col - Allows the user to configure both the field value
// and the Format Parameters for text line Format Type
// 'Line1Column'.
//
// The '1-Column' Text Line is designed to produce a single line of
// text consisting of one text field with optional left and right
// margins.
//
// Unlike method TextFormatterCollection.AddLine1Col(), this method
// has no requirement for previously configured Standard Format
// Parameters because those parameters are created in a single call
// to this method provided that input parameter
// 'saveFmtParamsAsDefault' is set to 'true'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// When input parameter 'saveFmtParamsAsDefault' is set to 'true',
// this method will create the Standard Format Parameters for this
// and all future 'Line1Column' Text types created by this instance
// of TextFormatterCollection.
//
// If input parameter  'saveFmtParamsAsDefault' is set to 'true',
// then in the future, users should call
// TextFormatterCollection.AddLine1Col() to reduce the number of
// input parameters required to produce other '1-Column' Text Line.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//		leftMarginStr              string
//		   - The contents of this string will be used as the left
//		     margin for the 'Column1' field.
//
//		     If no left margin is required, set 'leftMarginStr' to a
//		     zero length or empty string, and no left margin will be
//		     created.
//
//
//		column1Field               interface{}
//		   - This parameter is an empty interface which must contain
//		     one of several specific types. This empty interface type
//		     will be converted to a string and configured as the single
//		     text field in this 'Line1Column' Text Line.
//
//		     Supported types which may be submitted through this empty
//		     interface parameter are listed as follows:
//		        time.Time (Converted using default format)
//		        string
//		        bool
//		        uint, uint8, uint16, uint32, uint64,
//		        int, int8, int16, int32, int64
//		        float32, float64
//		        *big.Int *big.Float
//		        fmt.Stringer (types that support this interface)
//		        TextInputParamFieldDateTimeDto
//		              (Converts date time to string)
//
//		     If the 'column1Field' is not convertible to one of the
//		     supported types, an error will be returned.
//
//		     If the converted string value for 'column1Field' is empty,
//		     it will be defaulted to a single white space character
//		     (" ").
//
//
//		column1FieldLength         int
//		   - This is the length of the text field in which the
//		     formatted 'column1Field' string will be displayed. If
//		     'column1FieldLength' is less than the length of the
//		     'column1Field' string, it will be automatically set equal
//		     to the 'column1Field' string length.
//
//		     If 'column1FieldLength' is greater than the length of the
//		     'column1Field' text string, the 'column1FieldJustify'
//		     parameter will be used to configure or justify the text
//		     within the boundaries of the text field defined by
//		     'column1FieldLength'.
//
//		     To automatically set the value of 'column1FieldLength' to
//		     the length of 'column1Field', set this parameter to a
//		     value of minus one (-1).
//
//		     If this parameter is submitted with a value less than
//		     minus one (-1) or greater than 1-million (1,000,000), an
//		     error will be returned.
//
//
//		column1FieldJustify        TextJustify
//		    An enumeration value specifying the justification of the
//		    'column1Field' string within the text field specified by
//		    'column1FieldLength'.
//
//		    Text justification can only be evaluated in the context of
//		    a text label, field length and a Text Justification object
//		    of type TextJustify. This is because text labels with a
//		    field length equal to or less than the length of the text
//		    label never use text justification. In these cases, text
//		    justification is completely ignored.
//
//		    If the field length is greater than the length of the text
//		    label, text justification must be equal to one of these
//		    three valid values:
//		        TextJustify(0).Left()
//		        TextJustify(0).Right()
//		        TextJustify(0).Center()
//
//		    You can also use the abbreviated text justification
//		    enumeration syntax as follows:
//
//		        TxtJustify.Left()
//		        TxtJustify.Right()
//		        TxtJustify.Center()
//
//
//		rightMarginStr             string
//		   - The contents of the string will be used as the right
//		     margin for the 'Column1' field.
//
//		     If no right margin is required, set 'rightMarginStr' to a
//		     zero length or empty string, and no right margin will be
//		     created.
//
//
//		 turnLineTerminationOff    bool
//		   - When this parameter is set to 'true', no line termination
//		     sequence will be configured for this 'Line1Column' Text
//		     Line.
//
//		     Text Lines operate on the assumption that a line
//		     termination is standard operating procedure. The default
//		     line terminator for text lines is the new line character,
//		     '\n'.
//
//		     Users have the option of turning off the entire line
//		     termination sequence if this parameter is set to 'true'.
//
//
//		lineTerminator             string
//		   - If this parameter is set to an empty string, the default
//		     line terminator, a new line character '\n', will be
//		     applied to this 'Line1Column' Text Line.
//
//		     If this string is populated and the string length is
//		     greater than zero (0), an alternate line termination
//		     sequence will be configured using the characters provided
//		     in the 'lineTerminator' string.
//
//		     Remember that the application of a line termination
//		     sequence is controlled by parameter
//		     'turnLineTerminationOff'. If 'turnLineTerminationOff' is
//		     set to 'true', no line termination characters will be
//		     configured for this 'Line1Column' Text Line.
//
//
//		maxLineLength              int
//		   - The maximum length of the line on which this 'Line1Column'
//		     Text will be presented.
//
//		     Set this parameter to minus one (-1), and no maximum line
//		     limits will be applied.
//
//		     'maxLineLength' is used in conjunction with parameter
//		     'turnAutoLineLengthBreaksOn' to automatically place text
//		     fields on separate text lines when that text exceeds the
//		     maximum text line length ('maxLineLength').
//
//		     If the value of 'maxLineLength' is less than minus one
//		     (-1), an error will be returned.
//
//		     If the value of 'maxLineLength' is zero (0), an error will
//		     be returned.
//
//
//		turnAutoLineLengthBreaksOn bool
//		   - This parameter controls whether text lines which exceed
//		     the maximum line length ('maxLineLength') are broken up
//		     and presented on the following line.
//
//		     To apply automatic line breaking at the maximum line
//		     length, set the value of this parameter to 'true'.
//
//
//		saveFmtParamsAsDefault     bool
//		   - When this parameter is set to 'true', the Text Field
//		     Format Parameters will be saved as the default Format
//		     Parameters for this specific number of Text Columns.
//
//		     If this parameter is set to 'false', the Text Field
//		     Format Parameters associated with this number of Text
//		     Columns will NOT be saved as a default and therefore
//		     these format parameters will not be used in subsequent
//		     calls to the 'Add' methods.
//
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
//
//		     This empty interface must be convertible to one of the
//		     following types:
//
//		     1. nil - A nil value is valid and generates an empty
//		        collection of error prefix and error context
//		        information.
//
//		     2. string - A string containing error prefix information.
//
//		     3. []string A one-dimensional slice of strings containing
//		        error prefix information
//
//		     4. [][2]string A two-dimensional slice of strings
//		        containing error prefix and error context information.
//
//		     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		        from this object will be copied for use in error and
//		        informational messages.
//
//		     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		        Information from this object will be copied for use in
//		        error and informational messages.
//
//		     7. IBasicErrorPrefix - An interface to a method generating
//		        a two-dimensional slice of strings containing error
//		        prefix and error context information.
//
//		     If parameter 'errorPrefix' is NOT convertible to one of
//		     the valid types listed above, it will be considered
//		     invalid and trigger the return of an error.
//
//		     Types ErrPrefixDto and IBasicErrorPrefix are included in
//		     the 'errpref' software package,
//		     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If this method completes successfully and no errors are
//	     encountered, this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (txtFmtCollection *TextFormatterCollection) CfgLine1Col(
	leftMarginStr string,
	column1Field interface{},
	column1FieldLength int,
	column1FieldJustify TextJustify,
	rightMarginStr string,
	turnLineTerminationOff bool,
	lineTerminator string,
	maxLineLength int,
	turnAutoLineLengthBreaksOn bool,
	saveFmtParamsAsDefault bool,
	errorPrefix interface{}) error {

	if txtFmtCollection.lock == nil {
		txtFmtCollection.lock = new(sync.Mutex)
	}

	txtFmtCollection.lock.Lock()

	defer txtFmtCollection.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFormatterCollection."+
			"CfgLine1Col()",
		"")

	if err != nil {
		return err
	}

	var column1FieldText string

	column1FieldText,
		err =
		new(textSpecificationAtom).
			convertParamEmptyInterfaceToString(
				column1Field,
				"column1Field",
				ePrefix.XCpy(
					"column1Field"))

	if err != nil {
		return err
	}

	if column1FieldLength < -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'Col1FieldLength' is invalid!\n"+
			"'column1FieldLength' has a value less than minus one (-1)\n"+
			"column1FieldLength = '%v'\n",
			ePrefix.String(),
			column1FieldLength)

		return err
	}

	if !column1FieldJustify.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'Col1FieldJustify' is invalid!\n"+
			"'Col1FieldJustify' must be set to 'Left', 'Right' or 'Center'.\n"+
			"Col1FieldJustify String Value  = '%v'\n"+
			"Col1FieldJustify Integer Value = '%v'\n",
			ePrefix.String(),
			column1FieldJustify.String(),
			column1FieldJustify.XValueInt())

		return err

	}

	if maxLineLength < -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'maxLineLength' is invalid!\n"+
			"'maxLineLength' has a value less than minus one (-1)\n"+
			"maxLineLength = '%v'\n",
			ePrefix.String(),
			maxLineLength)

		return err

	}

	if maxLineLength == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'maxLineLength' is invalid!\n"+
			"'maxLineLength' has a value of zero (0).\n",
			ePrefix.String())

		return err

	}

	newStdFmtParams := TextFmtParamsLineColumnsDto{
		FormatType: TxtFieldType.LineColumns(),
		FieldFormatParams: []TextFieldFmtParamsDto{
			{
				LeftMarginStr:  leftMarginStr,
				FieldLength:    column1FieldLength,
				FieldJustify:   column1FieldJustify,
				DateTimeFormat: "",
				RightMarginStr: rightMarginStr,
				lock:           nil,
			},
		},
		TurnLineTerminationOff:     turnLineTerminationOff,
		LineTerminator:             lineTerminator,
		MaxLineLength:              maxLineLength,
		TurnAutoLineLengthBreaksOn: turnAutoLineLengthBreaksOn,
		lock:                       nil,
	}

	newTextLine1Cols := TextLineColumnsDto{
		FormatType: TxtFieldType.LineColumns(),
		TextFieldsContent: []TextFieldsContentDto{
			{
				TextFieldString:   column1FieldText,
				TextFieldDateTime: time.Time{},
				lock:              nil,
			},
		},
		FmtParameters: newStdFmtParams,
		lock:          nil,
	}

	newTextFormatter := TextFormatterDto{
		FormatType:          TxtFieldType.LineColumns(),
		DateTime:            TextFieldDateTimeDto{},
		Filler:              TextFieldFillerDto{},
		Label:               TextFieldLabelDto{},
		Spacer:              TextFieldSpacerDto{},
		BlankLine:           TextLineBlankDto{},
		SolidLine:           TextLineSolidDto{},
		LineColumns:         newTextLine1Cols,
		LinesTimerStartStop: TextLineTimerStartStopDto{},
	}

	if saveFmtParamsAsDefault == true {

		err =
			new(textFormatterCollectionElectron).
				cfgNewStdTxtLineParameters(
					txtFmtCollection,
					newStdFmtParams,
					ePrefix.XCpy(
						"newStdFmtParams"))

		if err != nil {
			return err
		}
	}

	txtFmtCollection.fmtCollection =
		append(
			txtFmtCollection.fmtCollection,
			newTextFormatter)

	return err
}

// CfgLine2Col - Allows the user to configure both the field value
// and the Format Parameters for 2-Column Text Lines.
//
// The '2-Column' Text Line is designed to produce a single line of
// text consisting of two text fields each with optional left and
// right margins.
//
// Unlike method TextFormatterCollection.AddLine2Col(), this method
// has no requirement for previously configured Standard Format
// Parameters because those parameters are created in a single call
// to this method provided that input parameter
// 'saveFmtParamsAsDefault' is set to 'true'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// When input parameter 'saveFmtParamsAsDefault' is set to 'true',
// this method will create the Standard Format Parameters for this
// and all future '2-Column' Text Lines created by this instance of
// TextFormatterCollection.
//
// If input parameter 'saveFmtParamsAsDefault' is set to 'true',
// then in the future, users should call
// TextFormatterCollection.AddLine2Col() to reduce the number of
// input parameters required to produce other '2-Column' Text Line.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//		leftMarginStr              string
//		   - The contents of this string will be used as the left
//		     margin for the 'Column1' field.
//
//		     If no left margin is required, set 'leftMarginStr' to a
//		     zero length or empty string, and no left margin will be
//		     created.
//
//
//		column1Field               interface{}
//		   - This parameter is an empty interface which must contain
//		     one of several specific types. This empty interface type
//		     will be converted to a string and configured as the single
//		     text field in this 'Line1Column' Text Line.
//
//		     Supported types which may be submitted through this empty
//		     interface parameter are listed as follows:
//		        time.Time (Converted using default format)
//		        string
//		        bool
//		        uint, uint8, uint16, uint32, uint64,
//		        int, int8, int16, int32, int64
//		        float32, float64
//		        *big.Int *big.Float
//		        fmt.Stringer (types that support this interface)
//		        TextInputParamFieldDateTimeDto
//		              (Converts date time to string)
//
//		     If the 'column1Field' is not convertible to one of the
//		     supported types, an error will be returned.
//
//		     If the converted string value for 'column1Field' is empty,
//		     it will be defaulted to a single white space character
//		     (" ").
//
//
//		column1FieldLength         int
//		   - This is the length of the text field in which the
//		     formatted 'column1Field' string will be displayed. If
//		     'column1FieldLength' is less than the length of the
//		     'column1Field' string, it will be automatically set equal
//		     to the 'column1Field' string length.
//
//		     If 'column1FieldLength' is greater than the length of the
//		     'column1Field' text string, the 'column1FieldJustify'
//		     parameter will be used to configure or justify the text
//		     within the boundaries of the text field defined by
//		     'column1FieldLength'.
//
//		     To automatically set the value of 'column1FieldLength' to
//		     the length of 'column1Field', set this parameter to a
//		     value of minus one (-1).
//
//		     If this parameter is submitted with a value less than
//		     minus one (-1) or greater than 1-million (1,000,000), an
//		     error will be returned.
//
//
//		column1FieldJustify        TextJustify
//		    An enumeration value specifying the justification of the
//		    'column1Field' string within the text field specified by
//		    'column1FieldLength'.
//
//		    Text justification can only be evaluated in the context of
//		    a text label, field length and a Text Justification object
//		    of type TextJustify. This is because text labels with a
//		    field length equal to or less than the length of the text
//		    label never use text justification. In these cases, text
//		    justification is completely ignored.
//
//		    If the field length is greater than the length of the text
//		    label, text justification must be equal to one of these
//		    three valid values:
//		        TextJustify(0).Left()
//		        TextJustify(0).Right()
//		        TextJustify(0).Center()
//
//		    You can also use the abbreviated text justification
//		    enumeration syntax as follows:
//
//		        TxtJustify.Left()
//		        TxtJustify.Right()
//		        TxtJustify.Center()
//
//
//		column1RightMarginStr      string
//		   - The contents of the string will be used as the right
//		     margin for the 'Column1' field.
//
//		     If no right margin is required, set
//		     'column1RightMarginStr' to a zero length or empty string,
//		     and no right margin will be created.
//
//
//		column2Field               interface{}
//		   - This parameter is an empty interface which must contain
//		     one of several specific types. This empty interface type
//		     will be converted to a string and configured as the second
//		     text field in this '2-Column' Text Line.
//
//		     Supported types which may be submitted through this empty
//		     interface parameter are listed as follows:
//		        time.Time (Converted using default format)
//		        string
//		        bool
//		        uint, uint8, uint16, uint32, uint64,
//		        int, int8, int16, int32, int64
//		        float32, float64
//		        *big.Int *big.Float
//		        fmt.Stringer (types that support this interface)
//		        TextInputParamFieldDateTimeDto
//		              (Converts date time to string)
//
//		     If the 'column1Field' is not convertible to one of the
//		     supported types, an error will be returned.
//
//		     If the converted string value for 'column2Field' is empty,
//		     it will be defaulted to a single white space character
//		     (" ").
//
//
//		column2FieldLength         int
//		   - This is the length of the text field in which the
//		     formatted 'column2Field' string will be displayed. If
//		     'column2FieldLength' is less than the length of the
//		     'column2Field' string, it will be automatically set equal
//		     to the 'column2Field' string length.
//
//		     If 'column2FieldLength' is greater than the length of the
//		     'column2Field' text string, the 'column2FieldJustify'
//		     parameter will be used to configure or justify the text
//		     within the boundaries of the text field defined by
//		     'column2FieldLength'.
//
//		     To automatically set the value of 'column2FieldLength' to
//		     the length of the 'column2Field' string, set this
//		     parameter to a value of minus one (-1).
//
//		     If this parameter is submitted with a value less than
//		     minus one (-1) or greater than 1-million (1,000,000), an
//		     error will be returned.
//
//
//		column2FieldJustify        TextJustify
//		    An enumeration value specifying the justification of the
//		    'column2Field' string within the text field specified by
//		    'column2FieldLength'.
//
//		    Text justification can only be evaluated in the context of
//		    a text label, field length and a Text Justification object
//		    of type TextJustify. This is because text labels with a
//		    field length equal to or less than the length of the text
//		    label never use text justification. In these cases, text
//		    justification is completely ignored.
//
//		    If the field length is greater than the length of the text
//		    label, text justification must be equal to one of these
//		    three valid values:
//		        TextJustify(0).Left()
//		        TextJustify(0).Right()
//		        TextJustify(0).Center()
//
//		    You can also use the abbreviated text justification
//		    enumeration syntax as follows:
//
//		        TxtJustify.Left()
//		        TxtJustify.Right()
//		        TxtJustify.Center()
//
//
//		column2RightMarginStr      string
//		   - The contents of the string will be used as the right
//		     margin for the 'Column2' field.
//
//		     If no right margin is required, set
//		     'column2RightMarginStr' to a zero length or empty string,
//		     and no right margin will be created.
//
//
//		 turnLineTerminationOff    bool
//		   - When this parameter is set to 'true', no line termination
//		     sequence will be configured for this 'Line1Column' Text
//		     Line.
//
//		     Text Lines operate on the assumption that a line
//		     termination is standard operating procedure. The default
//		     line terminator for text lines is the new line character,
//		     '\n'.
//
//		     Users have the option of turning off the entire line
//		     termination sequence if this parameter is set to 'true'.
//
//
//		lineTerminator             string
//		   - If this parameter is set to an empty string, the default
//		     line terminator, a new line character '\n', will be
//		     applied to this '2-Column' Text Line.
//
//		     If this string is populated and the string length is
//		     greater than zero (0), an alternate line termination
//		     sequence will be configured using the characters provided
//		     in the 'lineTerminator' string.
//
//		     Remember that the application of a line termination
//		     sequence is controlled by parameter
//		     'turnLineTerminationOff'. If 'turnLineTerminationOff' is
//		     set to 'true', no line termination characters will be
//		     configured for this 'Line1Column' Text Line.
//
//
//		maxLineLength              int
//		   - The maximum length of the line on which this 'Line1Column'
//		     Text will be presented.
//
//		     Set this parameter to minus one (-1), and no maximum line
//		     limits will be applied.
//
//		     'maxLineLength' is used in conjunction with parameter
//		     'turnAutoLineLengthBreaksOn' to automatically place text
//		     fields on separate text lines when that text exceeds the
//		     maximum text line length ('maxLineLength').
//
//		     If the value of 'maxLineLength' is less than minus one
//		     (-1), an error will be returned.
//
//		     If the value of 'maxLineLength' is zero (0), an error will
//		     be returned.
//
//
//		turnAutoLineLengthBreaksOn bool
//		   - This parameter controls whether text lines which exceed
//		     the maximum line length ('maxLineLength') are broken up
//		     and presented on the following line.
//
//		     To apply automatic line breaking at the maximum line
//		     length, set the value of this parameter to 'true'.
//
//
//		saveFmtParamsAsDefault     bool
//		   - When this parameter is set to 'true', the Text Field
//		     Format Parameters will be saved as the default Format
//		     Parameters for this specific number of Text Columns.
//
//		     If this parameter is set to 'false', the Text Field
//		     Format Parameters associated with this number of Text
//		     Columns will NOT be saved as a default and therefore
//		     these format parameters will not be used in subsequent
//		     calls to the 'Add' methods.
//
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
//
//		     This empty interface must be convertible to one of the
//		     following types:
//
//		     1. nil - A nil value is valid and generates an empty
//		        collection of error prefix and error context
//		        information.
//
//		     2. string - A string containing error prefix information.
//
//		     3. []string A one-dimensional slice of strings containing
//		        error prefix information
//
//		     4. [][2]string A two-dimensional slice of strings
//		        containing error prefix and error context information.
//
//		     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		        from this object will be copied for use in error and
//		        informational messages.
//
//		     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		        Information from this object will be copied for use in
//		        error and informational messages.
//
//		     7. IBasicErrorPrefix - An interface to a method generating
//		        a two-dimensional slice of strings containing error
//		        prefix and error context information.
//
//		     If parameter 'errorPrefix' is NOT convertible to one of
//		     the valid types listed above, it will be considered
//		     invalid and trigger the return of an error.
//
//		     Types ErrPrefixDto and IBasicErrorPrefix are included in
//		     the 'errpref' software package,
//		     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If this method completes successfully and no errors are
//	     encountered, this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (txtFmtCollection *TextFormatterCollection) CfgLine2Col(
	leftMarginStr string,
	column1Field interface{},
	column1FieldLength int,
	column1FieldJustify TextJustify,
	column1RightMarginStr string,
	column2Field interface{},
	column2FieldLength int,
	column2FieldJustify TextJustify,
	column2RightMarginStr string,
	turnLineTerminationOff bool,
	lineTerminator string,
	maxLineLength int,
	turnAutoLineLengthBreaksOn bool,
	saveFmtParamsAsDefault bool,
	errorPrefix interface{}) error {

	if txtFmtCollection.lock == nil {
		txtFmtCollection.lock = new(sync.Mutex)
	}

	txtFmtCollection.lock.Lock()

	defer txtFmtCollection.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFormatterCollection."+
			"CfgLine1Col()",
		"")

	if err != nil {
		return err
	}

	var column1FieldText string

	column1FieldText,
		err =
		new(textSpecificationAtom).
			convertParamEmptyInterfaceToString(
				column1Field,
				"column1Field",
				ePrefix.XCpy(
					"column1Field"))

	if err != nil {
		return err
	}

	var column2FieldText string

	column2FieldText,
		err = textSpecificationAtom{}.ptr().
		convertParamEmptyInterfaceToString(
			column2Field,
			"column2Field",
			ePrefix.XCpy(
				"column2Field"))

	if err != nil {
		return err
	}

	if column1FieldLength < -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'Col1FieldLength' is invalid!\n"+
			"'column1FieldLength' has a value less than minus one (-1)\n"+
			"column1FieldLength = '%v'\n",
			ePrefix.String(),
			column1FieldLength)

		return err
	}

	if !column1FieldJustify.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'Col1FieldJustify' is invalid!\n"+
			"'Col1FieldJustify' must be set to 'Left', 'Right' or 'Center'.\n"+
			"Col1FieldJustify String Value  = '%v'\n"+
			"Col1FieldJustify Integer Value = '%v'\n",
			ePrefix.String(),
			column1FieldJustify.String(),
			column1FieldJustify.XValueInt())

		return err

	}

	if column2FieldLength < -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'Col2FieldLength' is invalid!\n"+
			"'column2FieldLength' has a value less than minus one (-1)\n"+
			"column2FieldLength = '%v'\n",
			ePrefix.String(),
			column2FieldLength)

		return err
	}

	if !column2FieldJustify.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'Col2FieldJustify' is invalid!\n"+
			"'Col2FieldJustify' must be set to 'Left', 'Right' or 'Center'.\n"+
			"Col2FieldJustify String Value  = '%v'\n"+
			"Col2FieldJustify Integer Value = '%v'\n",
			ePrefix.String(),
			column2FieldJustify.String(),
			column2FieldJustify.XValueInt())

		return err

	}

	if maxLineLength < -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'maxLineLength' is invalid!\n"+
			"'maxLineLength' has a value less than minus one (-1)\n"+
			"maxLineLength = '%v'\n",
			ePrefix.String(),
			maxLineLength)

		return err

	}

	if maxLineLength == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'maxLineLength' is invalid!\n"+
			"'maxLineLength' has a value of zero (0).\n",
			ePrefix.String())

		return err

	}

	newStdFmtParams := TextFmtParamsLineColumnsDto{
		FormatType: TxtFieldType.LineColumns(),
		FieldFormatParams: []TextFieldFmtParamsDto{
			{
				LeftMarginStr:  leftMarginStr,
				FieldLength:    column1FieldLength,
				FieldJustify:   column1FieldJustify,
				DateTimeFormat: "",
				RightMarginStr: column1RightMarginStr,
				lock:           nil,
			},
			{
				LeftMarginStr:  "",
				FieldLength:    column2FieldLength,
				FieldJustify:   column2FieldJustify,
				DateTimeFormat: "",
				RightMarginStr: column2RightMarginStr,
				lock:           nil,
			},
		},
		TurnLineTerminationOff:     turnLineTerminationOff,
		LineTerminator:             lineTerminator,
		MaxLineLength:              maxLineLength,
		TurnAutoLineLengthBreaksOn: turnAutoLineLengthBreaksOn,
		lock:                       nil,
	}

	newTextLine1Cols := TextLineColumnsDto{
		FormatType: TxtFieldType.LineColumns(),
		TextFieldsContent: []TextFieldsContentDto{
			{
				TextFieldString:   column1FieldText,
				TextFieldDateTime: time.Time{},
				lock:              nil,
			},
			{
				TextFieldString:   column2FieldText,
				TextFieldDateTime: time.Time{},
				lock:              nil,
			},
		},
		FmtParameters: newStdFmtParams,
		lock:          nil,
	}

	newTextFormatter := TextFormatterDto{
		FormatType:          TxtFieldType.LineColumns(),
		DateTime:            TextFieldDateTimeDto{},
		Filler:              TextFieldFillerDto{},
		Label:               TextFieldLabelDto{},
		Spacer:              TextFieldSpacerDto{},
		BlankLine:           TextLineBlankDto{},
		SolidLine:           TextLineSolidDto{},
		LineColumns:         newTextLine1Cols,
		LinesTimerStartStop: TextLineTimerStartStopDto{},
	}

	if saveFmtParamsAsDefault == true {

		err =
			new(textFormatterCollectionElectron).
				cfgNewStdTxtLineParameters(
					txtFmtCollection,
					newStdFmtParams,
					ePrefix.XCpy(
						"newStdFmtParams"))

		if err != nil {
			return err
		}
	}

	txtFmtCollection.fmtCollection =
		append(
			txtFmtCollection.fmtCollection,
			newTextFormatter)

	return err
}

// CfgLineMultiCol - Adds a single Text Line consisting of one or
// more columns to the Text Formatter Collection.
//
// The number of 'textFields' MUST MATCH the number of Field
// Format Parameters
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// When input parameter 'saveFmtParamsAsDefault' is set to 'true',
// this method will create the Standard Format Parameters for this
// and all future Text Lines with the same number of columns.
//
// If input parameter  'saveFmtParamsAsDefault' is set to 'true',
// then in the future, users should call
// TextFormatterCollection.AddLineMultiCol() to reduce the number
// of input parameters required to produce other Text Lines with
// the same number of text columns.
//
// If the number of elements in the textFields array does NOT MATCH
// the number of elements in the 'fieldFormatParams' array, an
// error will be returned.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//		textFields               []interface{}
//		   - This parameter is an array of empty interfaces which must
//		     contain one of several specific data types. This empty
//		     interface type will be converted to a string and configured
//		     as the text columns within this Text Lines.
//
//		     Supported types which may be submitted through this empty
//		     interface parameter are listed as follows:
//		        time.Time (Converted using default format)
//		        string
//		        bool
//		        uint, uint8, uint16, uint32, uint64,
//		        int, int8, int16, int32, int64
//		        float32, float64
//		        *big.Int *big.Float
//		        fmt.Stringer (types that support this interface)
//		        TextInputParamFieldDateTimeDto
//		              (Converts date time to string)
//
//		     If the 'column1Field' is not convertible to one of the
//		     supported types, an error will be returned.
//
//		     If the converted string value for 'column1Field' is empty,
//		     it will be defaulted to a single white space character
//		     (" ").
//
//		     If the number of elements in this array does NOT MATCH the
//		     number of elements in the 'fieldFormatParams' array, an
//		     error will be returned.
//
//
//		fieldFormatParams          []TextFieldFmtParamsDto
//		   - An array of Text Field Format Parameters
//		    (TextFieldFmtParamsDto).
//
//		    If the number of elements in this array does NOT MATCH the
//		    number of elements in the 'textFields' array, an error will
//		    be returned.
//
//		    The TextFieldFmtParamsDto type is defined as follows:
//
//		      type TextFieldFmtParamsDto struct {
//		        LeftMarginStr  string
//		        FieldLength    int
//		        FieldJustify   TextJustify
//		        DateTimeFormat string
//		        RightMarginStr string
//		      }
//
//
//		 turnLineTerminationOff    bool
//		   - When this parameter is set to 'true', no line termination
//		     sequence will be configured for this 'Line1Column' Text
//		     Line.
//
//		     Text Lines operate on the assumption that a line
//		     termination is standard operating procedure. The default
//		     line terminator for text lines is the new line character,
//		     '\n'.
//
//		     Users have the option of turning off the entire line
//		     termination sequence if this parameter is set to 'true'.
//
//
//		lineTerminator             string
//		   - If this parameter is set to an empty string, the default
//		     line terminator, a new line character '\n', will be
//		     applied to this 'Line1Column' Text Line.
//
//		     If this string is populated and the string length is
//		     greater than zero (0), an alternate line termination
//		     sequence will be configured using the characters provided
//		     in the 'lineTerminator' string.
//
//		     Remember that the application of a line termination
//		     sequence is controlled by parameter
//		     'turnLineTerminationOff'. If 'turnLineTerminationOff' is
//		     set to 'true', no line termination characters will be
//		     configured for this 'Line1Column' Text Line.
//
//
//		maxLineLength              int
//		   - The maximum length of the line on which this 'Line1Column'
//		     Text will be presented.
//
//		     Set this parameter to minus one (-1), and no maximum line
//		     limits will be applied.
//
//		     'maxLineLength' is used in conjunction with parameter
//		     'turnAutoLineLengthBreaksOn' to automatically place text
//		     fields on separate text lines when that text exceeds the
//		     maximum text line length ('maxLineLength').
//
//		     If the value of 'maxLineLength' is less than minus one
//		     (-1), an error will be returned.
//
//		     If the value of 'maxLineLength' is zero (0), an error will
//		     be returned.
//
//
//		turnAutoLineLengthBreaksOn bool
//		   - This parameter controls whether text lines which exceed
//		     the maximum line length ('maxLineLength') are broken up
//		     and presented on the following line.
//
//		     To apply automatic line breaking at the maximum line
//		     length, set the value of this parameter to 'true'.
//
//
//		saveFmtParamsAsDefault     bool
//		   - When this parameter is set to 'true', the Text Field
//		     Format Parameters will be saved as the default Format
//		     Parameters for this specific number of Text Columns.
//
//		     If this parameter is set to 'false', the Text Field
//		     Format Parameters associated with this number of Text
//		     Columns will NOT be saved as a default and therefore
//		     these format parameters will not be used in subsequent
//		     calls to the 'Add' methods.
//
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
//
//		     This empty interface must be convertible to one of the
//		     following types:
//
//		     1. nil - A nil value is valid and generates an empty
//		        collection of error prefix and error context
//		        information.
//
//		     2. string - A string containing error prefix information.
//
//		     3. []string A one-dimensional slice of strings containing
//		        error prefix information
//
//		     4. [][2]string A two-dimensional slice of strings
//		        containing error prefix and error context information.
//
//		     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		        from this object will be copied for use in error and
//		        informational messages.
//
//		     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		        Information from this object will be copied for use in
//		        error and informational messages.
//
//		     7. IBasicErrorPrefix - An interface to a method generating
//		        a two-dimensional slice of strings containing error
//		        prefix and error context information.
//
//		     If parameter 'errorPrefix' is NOT convertible to one of
//		     the valid types listed above, it will be considered
//		     invalid and trigger the return of an error.
//
//		     Types ErrPrefixDto and IBasicErrorPrefix are included in
//		     the 'errpref' software package,
//		     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If this method completes successfully and no errors are
//	     encountered, this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (txtFmtCollection *TextFormatterCollection) CfgLineMultiCol(
	textFields []interface{},
	fieldFormatParams []TextFieldFmtParamsDto,
	turnLineTerminationOff bool,
	lineTerminator string,
	maxLineLength int,
	turnAutoLineLengthBreaksOn bool,
	saveFmtParamsAsDefault bool,
	errorPrefix interface{}) error {

	if txtFmtCollection.lock == nil {
		txtFmtCollection.lock = new(sync.Mutex)
	}

	txtFmtCollection.lock.Lock()

	defer txtFmtCollection.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFormatterCollection."+
			"CfgLineMultiCol()",
		"")

	if err != nil {
		return err
	}

	lenTextFields := len(textFields)

	lenFieldFmtParams := len(fieldFormatParams)

	if lenTextFields != lenFieldFmtParams {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameters 'textFields' and\n"+
			"'fieldFormatParams' are invalid!\n"+
			"Both arrays must have the same length\n"+
			"or same number of elements!\n"+
			"'textFields' Length       = '%v'\n"+
			"'fieldFormatParams Length = '%v'\n",
			ePrefix.String(),
			lenTextFields,
			lenFieldFmtParams)

		return err
	}

	textFieldsContent := make([]TextFieldsContentDto, lenTextFields)

	for i := 0; i < lenTextFields; i++ {

		textFieldsContent[i].TextFieldString,
			err = textSpecificationAtom{}.ptr().
			convertParamEmptyInterfaceToString(
				textFields[i],
				fmt.Sprintf("textFields[%v]", i),
				ePrefix.XCpy(
					fmt.Sprintf("textFields[%v]", i)))

		if err != nil {
			return err
		}

		textFieldsContent[i].TextFieldDateTime = time.Time{}
		textFieldsContent[i].lock = nil
	}

	newStdFmtParams := TextFmtParamsLineColumnsDto{
		FormatType:                 TxtFieldType.LineColumns(),
		FieldFormatParams:          fieldFormatParams,
		TurnLineTerminationOff:     turnLineTerminationOff,
		LineTerminator:             lineTerminator,
		MaxLineLength:              maxLineLength,
		TurnAutoLineLengthBreaksOn: turnAutoLineLengthBreaksOn,
		lock:                       nil,
	}

	newTextLineCols := TextLineColumnsDto{
		FormatType:        TxtFieldType.LineColumns(),
		TextFieldsContent: textFieldsContent,
		FmtParameters:     newStdFmtParams,
		lock:              nil,
	}

	newTextFormatter := TextFormatterDto{
		FormatType:          TxtFieldType.LineColumns(),
		DateTime:            TextFieldDateTimeDto{},
		Filler:              TextFieldFillerDto{},
		Label:               TextFieldLabelDto{},
		Spacer:              TextFieldSpacerDto{},
		BlankLine:           TextLineBlankDto{},
		SolidLine:           TextLineSolidDto{},
		LineColumns:         newTextLineCols,
		LinesTimerStartStop: TextLineTimerStartStopDto{},
	}

	if saveFmtParamsAsDefault == true {

		err =
			new(textFormatterCollectionElectron).
				cfgNewStdTxtLineParameters(
					txtFmtCollection,
					newStdFmtParams,
					ePrefix.XCpy(
						"newStdFmtParams"))

		if err != nil {
			return err
		}

	}

	txtFmtCollection.fmtCollection =
		append(
			txtFmtCollection.fmtCollection,
			newTextFormatter)

	return err
}

// CopyIn - Copies the data fields from an incoming instance of
// TextFormatterCollection ('incomingTxtFmtCol') to the data fields
// of the current TextFormatterCollection instance
// ('txtFmtCollection').
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// All the data fields in current TextFormatterCollection instance
// ('txtFmtCollection') will be deleted and overwritten.
//
// NO DATA VALIDATION is performed on input parameter
// 'incomingTxtFmtCol'
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//		incomingTxtFmtCol     *TextFormatterCollection
//		   - A pointer to an instance of TextFormatterCollection. This
//		     method will NOT change the data values of member variables
//		     contained in this instance.
//
//		     All data values in this TextFormatterCollection instance
//		     ('incomingTxtFmtCol') will be copied to the current
//		     TextFormatterCollection instance ('txtFmtCollection').
//
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
//
//		     This empty interface must be convertible to one of the
//		     following types:
//
//		     1. nil - A nil value is valid and generates an empty
//		        collection of error prefix and error context
//		        information.
//
//		     2. string - A string containing error prefix information.
//
//		     3. []string A one-dimensional slice of strings containing
//		        error prefix information
//
//		     4. [][2]string A two-dimensional slice of strings
//		        containing error prefix and error context information.
//
//		     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		        from this object will be copied for use in error and
//		        informational messages.
//
//		     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		        Information from this object will be copied for use in
//		        error and informational messages.
//
//		     7. IBasicErrorPrefix - An interface to a method generating
//		        a two-dimensional slice of strings containing error
//		        prefix and error context information.
//
//		     If parameter 'errorPrefix' is NOT convertible to one of
//		     the valid types listed above, it will be considered
//		     invalid and trigger the return of an error.
//
//		     Types ErrPrefixDto and IBasicErrorPrefix are included in
//		     the 'errpref' software package,
//		     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
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
func (txtFmtCollection *TextFormatterCollection) CopyIn(
	incomingTxtFmtCol *TextFormatterCollection,
	errorPrefix interface{}) error {

	if txtFmtCollection.lock == nil {
		txtFmtCollection.lock = new(sync.Mutex)
	}

	txtFmtCollection.lock.Lock()

	defer txtFmtCollection.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFormatterCollection."+
			"CopyIn()",
		"")

	if err != nil {
		return err
	}

	return new(textFormatterCollectionNanobot).
		copy(
			txtFmtCollection,
			incomingTxtFmtCol,
			ePrefix.XCpy(
				"txtFmtCollection<-incomingTxtFmtCol"))

}

// CopyOut - Returns a deep copy of the current
// TextFormatterCollection instance.
//
// No data validation is performed on the current
// TextFormatterCollection instance prior to returning the deep
// copy of this instance.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
//
//		     This empty interface must be convertible to one of the
//		     following types:
//
//		     1. nil - A nil value is valid and generates an empty
//		        collection of error prefix and error context
//		        information.
//
//		     2. string - A string containing error prefix information.
//
//		     3. []string A one-dimensional slice of strings containing
//		        error prefix information
//
//		     4. [][2]string A two-dimensional slice of strings
//		        containing error prefix and error context information.
//
//		     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		        from this object will be copied for use in error and
//		        informational messages.
//
//		     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		        Information from this object will be copied for use in
//		        error and informational messages.
//
//		     7. IBasicErrorPrefix - An interface to a method generating
//		        a two-dimensional slice of strings containing error
//		        prefix and error context information.
//
//		     If parameter 'errorPrefix' is NOT convertible to one of
//		     the valid types listed above, it will be considered
//		     invalid and trigger the return of an error.
//
//		     Types ErrPrefixDto and IBasicErrorPrefix are included in
//		     the 'errpref' software package,
//		     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	deepCopyTxtFmtCol          TextFormatterCollection
//	   - If this method completes successfully and no errors are
//	     encountered, this parameter will return a deep copy of the
//	     current TextFormatterCollection instance.
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
func (txtFmtCollection *TextFormatterCollection) CopyOut(
	errorPrefix interface{}) (
	deepCopyTxtFmtCol TextFormatterCollection,
	err error) {

	if txtFmtCollection.lock == nil {
		txtFmtCollection.lock = new(sync.Mutex)
	}

	txtFmtCollection.lock.Lock()

	defer txtFmtCollection.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFormatterCollection."+
			"CopyOut()",
		"")

	if err != nil {
		return deepCopyTxtFmtCol, err
	}

	err = new(textFormatterCollectionNanobot).
		copy(
			&deepCopyTxtFmtCol,
			txtFmtCollection,
			ePrefix.XCpy(
				"deepCopyTxtFmtCol<-txtFmtCollection"))

	return deepCopyTxtFmtCol, err
}

// Empty - Resets all internal member variables for the current
// instance of TextFormatterCollection to their zero or
// uninitialized states. This method will leave the current
// instance of TextFormatterCollection in an invalid state and
// unavailable for immediate reuse.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete all member variable data values in this
// current instance of TextFormatterCollection. All member
// variable data values will be reset to their zero or
// uninitialized states.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	NONE
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	NONE
func (txtFmtCollection *TextFormatterCollection) Empty() {

	if txtFmtCollection.lock == nil {
		txtFmtCollection.lock = new(sync.Mutex)
	}

	txtFmtCollection.lock.Lock()

	new(textFormatterCollectionMolecule).
		empty(
			txtFmtCollection)

	txtFmtCollection.lock.Unlock()

	txtFmtCollection.lock = nil

}

// EmptyFormatterCollection - Deletes all member elements of the
// Text Formatter Collection.
//
// Internal member variable
//
//	TextFormatterCollection.fmtCollection
//
// will be set to 'nil'.
//
// The internal member variable that holds the Text Formatter
// Collection is defined as follows:
//
//	fmtCollection []TextFormatterDto
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	NONE
//
// ----------------------------------------------------------------
//
// Return Values
//
//	NONE
func (txtFmtCollection *TextFormatterCollection) EmptyFormatterCollection() {

	if txtFmtCollection.lock == nil {
		txtFmtCollection.lock = new(sync.Mutex)
	}

	txtFmtCollection.lock.Lock()

	defer txtFmtCollection.lock.Unlock()

	new(textFormatterCollectionAtom).
		emptyFormatterCollection(txtFmtCollection)

	return
}

// EmptyLineParamCollection - Deletes all member elements in the
// Standard Text Line Parameters Collection.
//
// Internal member variable
//
//	TextFormatterCollection.stdTextLineParamCollection
//
// will be set to 'nil'.
//
// The internal member variable that holds the Standard Text Line
// Parameters Collection is defined as follows:
//
//	stdTextLineParamCollection []TextFmtParamsLineColumnsDto
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	NONE
//
// ----------------------------------------------------------------
//
// Return Values
//
//	NONE
func (txtFmtCollection *TextFormatterCollection) EmptyLineParamCollection() {

	if txtFmtCollection.lock == nil {
		txtFmtCollection.lock = new(sync.Mutex)
	}

	txtFmtCollection.lock.Lock()

	defer txtFmtCollection.lock.Unlock()

	new(textFormatterCollectionAtom).
		emptyLineParamCollection(
			txtFmtCollection)

	return

}

// Equal - Receives a pointer to another instance of
// TextFormatterCollection and proceeds to compare the member
// variables to those of the current TextFormatterCollection
// instance in order to determine if they are equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables of both instances are equal in
// all respects, this flag is set to 'true'. Otherwise, this method
// returns 'false'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	incomingTxtFmtCol    *TextFormatterCollection
//	   - A pointer to an incoming instance of
//	     TextFormatterCollection. This method will compare all
//	     member variable data values in this instance against those
//	     contained in the current instance of
//	     TextFormatterCollection. If the data values in both
//	     instances are found to be equal in all respects, this
//	     method will return a boolean value of 'true'.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	bool
//	   - If the member variable data values contained in input
//	     parameter 'incomingTxtFmtCol' are equal in all respects to
//	     those contained in the current instance of
//	     TextFormatterCollection, this method will return a
//	     boolean value of 'true'. Otherwise a value of 'false' will
//	     be returned to the calling function.
func (txtFmtCollection *TextFormatterCollection) Equal(
	incomingTxtFmtCol *TextFormatterCollection) bool {

	if txtFmtCollection.lock == nil {
		txtFmtCollection.lock = new(sync.Mutex)
	}

	txtFmtCollection.lock.Lock()

	defer txtFmtCollection.lock.Unlock()

	return new(textFormatterCollectionMolecule).
		equal(
			txtFmtCollection,
			incomingTxtFmtCol)
}

// GetLengthFormatterCollection - Returns the length of the Text
// Formatter Collection contained in the current instance of
// TextFormatterCollection.
//
// The Text Formatter Collection is an array of Text Formatter Data
// Transfer Objects (TextFormatterDto).
//
// The collection length is returned as an integer value.
func (txtFmtCollection *TextFormatterCollection) GetLengthFormatterCollection() int {

	if txtFmtCollection.lock == nil {
		txtFmtCollection.lock = new(sync.Mutex)
	}

	txtFmtCollection.lock.Lock()

	defer txtFmtCollection.lock.Unlock()

	return len(txtFmtCollection.fmtCollection)
}

// GetLengthStdTextLineParamCollection - Returns the length of the
// Standard Text Line Parameter Collection contained in the current
// instance of TextFormatterCollection.
//
// The Standard Text Line Parameter Collection is an array of Text
// Format Parameters for Lines and Columns
// (TextFmtParamsLineColumnsDto).
//
// The Standard Text Line Parameter Collection is used to produced
// standardized lines of texts containing between one and eight
// label fields.
//
// The collection length is returned as an integer value.
func (txtFmtCollection *TextFormatterCollection) GetLengthStdTextLineParamCollection() int {

	if txtFmtCollection.lock == nil {
		txtFmtCollection.lock = new(sync.Mutex)
	}

	txtFmtCollection.lock.Lock()

	defer txtFmtCollection.lock.Unlock()

	return len(txtFmtCollection.stdTextLineParamCollection)
}

// SetStdFormatParamsLine1Col - Sets the Standard Format Parameters
// for a Text Line consisting of one text column. This standard
// format will be applied as the default format for all Text Lines
// consisting of one text column.
//
// After configuring Standard Format Parameters for 1-Column
// Text Lines, users should configure additional 1-Column Text
// Lines using method TextFormatterCollection.AddLine1Col() in
// order to reduce the number of input parameters required to
// produce a 1-Column Text Line.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//		leftMarginStr              string
//		   - The contents of this string will be used as the left
//		     margin for the 'Column1' field.
//
//		     If no left margin is required, set 'leftMarginStr' to a
//		     zero length or empty string, and no left margin will be
//		     created.
//
//
//		column1FieldLength         int
//		   - This is the length of the text field in which the
//		     formatted 'column1Field' string will be displayed. If
//		     'column1FieldLength' is less than the length of the
//		     'column1Field' string, it will be automatically set equal
//		     to the 'column1Field' string length.
//
//		     To automatically set the value of 'column1FieldLength' to
//		     the length of 'column1Field', set this parameter to a
//		     value of minus one (-1).
//
//		     If this parameter is submitted with a value less than
//		     minus one (-1) or greater than 1-million (1,000,000), an
//		     error will be returned.
//
//
//		column1FieldJustify        TextJustify
//		    An enumeration value specifying the justification of the
//		    'column1Field' string within the text field specified by
//		    'column1FieldLength'.
//
//		    Text justification can only be evaluated in the context of
//		    a text label, field length and a Text Justification object
//		    of type TextJustify. This is because text labels with a
//		    field length equal to or less than the length of the text
//		    label never use text justification. In these cases, text
//		    justification is completely ignored.
//
//		    If the field length is greater than the length of the text
//		    label, text justification must be equal to one of these
//		    three valid values:
//		        TextJustify(0).Left()
//		        TextJustify(0).Right()
//		        TextJustify(0).Center()
//
//		    You can also use the abbreviated text justification
//		    enumeration syntax as follows:
//
//		        TxtJustify.Left()
//		        TxtJustify.Right()
//		        TxtJustify.Center()
//
//
//		rightMarginStr             string
//		   - The contents of the string will be used as the right
//		     margin for the 'Column1' field.
//
//		     If no right margin is required, set 'rightMarginStr' to a
//		     zero length or empty string, and no right margin will be
//		     created.
//
//
//		 turnLineTerminationOff    bool
//		   - When this parameter is set to 'true', no line termination
//		     sequence will be configured for this 'Line1Column' Text
//		     Line.
//
//		     Text Lines operate on the assumption that a line
//		     termination is standard operating procedure. The default
//		     line terminator for text lines is the new line character,
//		     '\n'.
//
//		     Users have the option of turning off the entire line
//		     termination sequence if this parameter is set to 'true'.
//
//
//		lineTerminator             string
//		   - If this parameter is set to an empty string, the default
//		     line terminator, a new line character '\n', will be
//		     applied to 'Line1Column' Text Lines.
//
//		     If this string is populated and the string length is
//		     greater than zero (0), an alternate line termination
//		     sequence will be configured using the characters provided
//		     in the 'lineTerminator' string.
//
//		     Remember that the application of a line termination
//		     sequence is controlled by parameter
//		     'turnLineTerminationOff'. If 'turnLineTerminationOff' is
//		     set to 'true', no line termination characters will be
//		     configured for this 'Line1Column' Text Lines.
//
//
//		maxLineLength              int
//		   - The maximum length of the line on which this 'Line1Column'
//		     Text will be presented.
//
//		     Set this parameter to minus one (-1), and no maximum line
//		     limits will be applied.
//
//		     'maxLineLength' is used in conjunction with parameter
//		     'turnAutoLineLengthBreaksOn' to automatically place text
//		     fields on separate text lines when that text exceeds the
//		     maximum text line length ('maxLineLength').
//
//		     If the value of 'maxLineLength' is less than minus one
//		     (-1), an error will be returned.
//
//		     If the value of 'maxLineLength' is zero (0), an error will
//		     be returned.
//
//
//		turnAutoLineLengthBreaksOn bool
//		   - This parameter controls whether text lines which exceed
//		     the maximum line length ('maxLineLength') are broken up
//		     and presented on the following line.
//
//		     To apply automatic line breaking at the maximum line
//		     length, set the value of this parameter to 'true'.
//
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
//
//		     This empty interface must be convertible to one of the
//		     following types:
//
//		     1. nil - A nil value is valid and generates an empty
//		        collection of error prefix and error context
//		        information.
//
//		     2. string - A string containing error prefix information.
//
//		     3. []string A one-dimensional slice of strings containing
//		        error prefix information
//
//		     4. [][2]string A two-dimensional slice of strings
//		        containing error prefix and error context information.
//
//		     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		        from this object will be copied for use in error and
//		        informational messages.
//
//		     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		        Information from this object will be copied for use in
//		        error and informational messages.
//
//		     7. IBasicErrorPrefix - An interface to a method generating
//		        a two-dimensional slice of strings containing error
//		        prefix and error context information.
//
//		     If parameter 'errorPrefix' is NOT convertible to one of
//		     the valid types listed above, it will be considered
//		     invalid and trigger the return of an error.
//
//		     Types ErrPrefixDto and IBasicErrorPrefix are included in
//		     the 'errpref' software package,
//		     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If this method completes successfully and no errors are
//	     encountered, this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (txtFmtCollection *TextFormatterCollection) SetStdFormatParamsLine1Col(
	leftMarginStr string,
	column1FieldLength int,
	column1FieldJustify TextJustify,
	rightMarginStr string,
	turnLineTerminationOff bool,
	lineTerminator string,
	maxLineLength int,
	turnAutoLineLengthBreaksOn bool,
	errorPrefix interface{}) error {

	if txtFmtCollection.lock == nil {
		txtFmtCollection.lock = new(sync.Mutex)
	}

	txtFmtCollection.lock.Lock()

	defer txtFmtCollection.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFormatterCollection."+
			"SetStdFormatParamsLine1Col()",
		"")

	if err != nil {
		return err
	}

	if column1FieldLength < -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'column1FieldLength' is invalid!\n"+
			"'column1FieldLength' has a value less than minus one (-1)\n"+
			"column1FieldLength = '%v'\n",
			ePrefix.String(),
			column1FieldLength)

		return err
	}

	if !column1FieldJustify.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'column1FieldJustify' is invalid!\n"+
			"'column1FieldJustify' must be set to 'Left', 'Right' or 'Center'.\n"+
			"column1FieldJustify String Value  = '%v'\n"+
			"column1FieldJustify Integer Value = '%v'\n",
			ePrefix.String(),
			column1FieldJustify.String(),
			column1FieldJustify.XValueInt())

		return err

	}

	if maxLineLength < -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'maxLineLength' is invalid!\n"+
			"'maxLineLength' has a value less than minus one (-1)\n"+
			"maxLineLength = '%v'\n",
			ePrefix.String(),
			maxLineLength)

		return err

	}

	newStdFmtParams := TextFmtParamsLineColumnsDto{
		FormatType: TxtFieldType.LineColumns(),
		FieldFormatParams: []TextFieldFmtParamsDto{
			{
				LeftMarginStr:  leftMarginStr,
				FieldLength:    column1FieldLength,
				FieldJustify:   column1FieldJustify,
				DateTimeFormat: "",
				RightMarginStr: rightMarginStr,
				lock:           nil,
			},
		},
		TurnLineTerminationOff:     turnLineTerminationOff,
		LineTerminator:             lineTerminator,
		MaxLineLength:              maxLineLength,
		TurnAutoLineLengthBreaksOn: turnAutoLineLengthBreaksOn,
		lock:                       nil,
	}

	err =
		new(textFormatterCollectionElectron).
			cfgNewStdTxtLineParameters(
				txtFmtCollection,
				newStdFmtParams,
				ePrefix.XCpy(
					"newStdFmtParams"))

	return err
}

// SetStdFormatParamsLine2Col - Sets the Standard Format Parameters
// for a Text Line consisting of one text column. This standard
// format will be applied as the default format for all Text Lines
// consisting of one text column.
//
// After configuring Standard Format Parameters for 1-Column
// Text Lines, users should configure additional 1-Column Text
// Lines using method TextFormatterCollection.AddLine1Col() in
// order to reduce the number of input parameters required to
// produce a 1-Column Text Line.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//		leftMarginStr              string
//		   - The contents of this string will be used as the left
//		     margin for the 'Column1' field.
//
//		     If no left margin is required, set 'leftMarginStr' to a
//		     zero length or empty string, and no left margin will be
//		     created.
//
//
//		column1FieldLength         int
//		   - This is the length of the text field in which the
//		     formatted 'column1Field' string will be displayed. If
//		     'column1FieldLength' is less than the length of the
//		     'column1Field' string, it will be automatically set equal
//		     to the 'column1Field' string length.
//
//		     To automatically set the value of 'column1FieldLength' to
//		     the length of 'column1Field', set this parameter to a
//		     value of minus one (-1).
//
//		     If this parameter is submitted with a value less than
//		     minus one (-1) or greater than 1-million (1,000,000), an
//		     error will be returned.
//
//
//		column1FieldJustify        TextJustify
//		    An enumeration value specifying the justification of the
//		    'column1Field' string within the text field specified by
//		    'column1FieldLength'.
//
//		    Text justification can only be evaluated in the context of
//		    a text label, field length and a Text Justification object
//		    of type TextJustify. This is because text labels with a
//		    field length equal to or less than the length of the text
//		    label never use text justification. In these cases, text
//		    justification is completely ignored.
//
//		    If the field length is greater than the length of the text
//		    label, text justification must be equal to one of these
//		    three valid values:
//		        TextJustify(0).Left()
//		        TextJustify(0).Right()
//		        TextJustify(0).Center()
//
//		    You can also use the abbreviated text justification
//		    enumeration syntax as follows:
//
//		        TxtJustify.Left()
//		        TxtJustify.Right()
//		        TxtJustify.Center()
//
//
//		column1RightMarginStr      string
//		   - The contents of the string will be used as the right
//		     margin for the 'Column1' field.
//
//		     If no right margin is required, set
//		     'column1RightMarginStr' to a zero length or empty string,
//		     and no right margin will be created.
//
//
//		column2FieldLength         int
//		   - This is the length of the text field in which the
//		     formatted 'column2Field' string will be displayed. If
//		     'column2FieldLength' is less than the length of the
//		     'column2Field' string, it will be automatically set equal
//		     to the 'column2Field' string length.
//
//		     To automatically set the value of 'column2FieldLength' to
//		     the length of 'column2Field', set this parameter to a
//		     value of minus one (-1).
//
//		     If this parameter is submitted with a value less than
//		     minus one (-1) or greater than 1-million (1,000,000), an
//		     error will be returned.
//
//
//		column2FieldJustify        TextJustify
//		    An enumeration value specifying the justification of the
//		    'column2Field' string within the text field specified by
//		    'column2FieldLength'.
//
//		    Text justification can only be evaluated in the context of
//		    a text label, field length and a Text Justification object
//		    of type TextJustify. This is because text labels with a
//		    field length equal to or less than the length of the text
//		    label never use text justification. In these cases, text
//		    justification is completely ignored.
//
//		    If the field length is greater than the length of the text
//		    label, text justification must be equal to one of these
//		    three valid values:
//		        TextJustify(0).Left()
//		        TextJustify(0).Right()
//		        TextJustify(0).Center()
//
//		    You can also use the abbreviated text justification
//		    enumeration syntax as follows:
//
//		        TxtJustify.Left()
//		        TxtJustify.Right()
//		        TxtJustify.Center()
//
//
//		column2RightMarginStr      string
//		   - The contents of the string will be used as the right
//		     margin for the 'Column2' field.
//
//		     If no right margin is required, set
//		     'column1RightMarginStr' to a zero length or empty string,
//		     and no right margin will be created.
//
//
//		 turnLineTerminationOff    bool
//		   - When this parameter is set to 'true', no line termination
//		     sequence will be configured for this 'Line1Column' Text
//		     Line.
//
//		     Text Lines operate on the assumption that a line
//		     termination is standard operating procedure. The default
//		     line terminator for text lines is the new line character,
//		     '\n'.
//
//		     Users have the option of turning off the entire line
//		     termination sequence if this parameter is set to 'true'.
//
//
//		lineTerminator             string
//		   - If this parameter is set to an empty string, the default
//		     line terminator, a new line character '\n', will be
//		     applied to 'Line1Column' Text Lines.
//
//		     If this string is populated and the string length is
//		     greater than zero (0), an alternate line termination
//		     sequence will be configured using the characters provided
//		     in the 'lineTerminator' string.
//
//		     Remember that the application of a line termination
//		     sequence is controlled by parameter
//		     'turnLineTerminationOff'. If 'turnLineTerminationOff' is
//		     set to 'true', no line termination characters will be
//		     configured for this 'Line1Column' Text Lines.
//
//
//		maxLineLength              int
//		   - The maximum length of the line on which this 'Line1Column'
//		     Text will be presented.
//
//		     Set this parameter to minus one (-1), and no maximum line
//		     limits will be applied.
//
//		     'maxLineLength' is used in conjunction with parameter
//		     'turnAutoLineLengthBreaksOn' to automatically place text
//		     fields on separate text lines when that text exceeds the
//		     maximum text line length ('maxLineLength').
//
//		     If the value of 'maxLineLength' is less than minus one
//		     (-1), an error will be returned.
//
//		     If the value of 'maxLineLength' is zero (0), an error will
//		     be returned.
//
//
//		turnAutoLineLengthBreaksOn bool
//		   - This parameter controls whether text lines which exceed
//		     the maximum line length ('maxLineLength') are broken up
//		     and presented on the following line.
//
//		     To apply automatic line breaking at the maximum line
//		     length, set the value of this parameter to 'true'.
//
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
//
//		     This empty interface must be convertible to one of the
//		     following types:
//
//		     1. nil - A nil value is valid and generates an empty
//		        collection of error prefix and error context
//		        information.
//
//		     2. string - A string containing error prefix information.
//
//		     3. []string A one-dimensional slice of strings containing
//		        error prefix information
//
//		     4. [][2]string A two-dimensional slice of strings
//		        containing error prefix and error context information.
//
//		     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		        from this object will be copied for use in error and
//		        informational messages.
//
//		     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		        Information from this object will be copied for use in
//		        error and informational messages.
//
//		     7. IBasicErrorPrefix - An interface to a method generating
//		        a two-dimensional slice of strings containing error
//		        prefix and error context information.
//
//		     If parameter 'errorPrefix' is NOT convertible to one of
//		     the valid types listed above, it will be considered
//		     invalid and trigger the return of an error.
//
//		     Types ErrPrefixDto and IBasicErrorPrefix are included in
//		     the 'errpref' software package,
//		     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If this method completes successfully and no errors are
//	     encountered, this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (txtFmtCollection *TextFormatterCollection) SetStdFormatParamsLine2Col(
	leftMarginStr string,
	column1FieldLength int,
	column1FieldJustify TextJustify,
	column1RightMarginStr string,
	column2FieldLength int,
	column2FieldJustify TextJustify,
	column2RightMarginStr string,
	turnLineTerminationOff bool,
	lineTerminator string,
	maxLineLength int,
	turnAutoLineLengthBreaksOn bool,
	errorPrefix interface{}) error {

	if txtFmtCollection.lock == nil {
		txtFmtCollection.lock = new(sync.Mutex)
	}

	txtFmtCollection.lock.Lock()

	defer txtFmtCollection.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFormatterCollection."+
			"SetStdFormatParamsLine2Col()",
		"")

	if err != nil {
		return err
	}

	if column1FieldLength < -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'column1FieldLength' is invalid!\n"+
			"'column1FieldLength' has a value less than minus one (-1)\n"+
			"column1FieldLength = '%v'\n",
			ePrefix.String(),
			column1FieldLength)

		return err
	}

	if !column1FieldJustify.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'column1FieldJustify' is invalid!\n"+
			"'column1FieldJustify' must be set to 'Left', 'Right' or 'Center'.\n"+
			"column1FieldJustify String Value  = '%v'\n"+
			"column1FieldJustify Integer Value = '%v'\n",
			ePrefix.String(),
			column1FieldJustify.String(),
			column1FieldJustify.XValueInt())

		return err

	}

	if column2FieldLength < -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'column2FieldLength' is invalid!\n"+
			"'column2FieldLength' has a value less than minus one (-1)\n"+
			"column2FieldLength = '%v'\n",
			ePrefix.String(),
			column2FieldLength)

		return err
	}

	if !column2FieldJustify.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'column2FieldJustify' is invalid!\n"+
			"'column2FieldJustify' must be set to 'Left', 'Right' or 'Center'.\n"+
			"column1FieldJustify String Value  = '%v'\n"+
			"column1FieldJustify Integer Value = '%v'\n",
			ePrefix.String(),
			column2FieldJustify.String(),
			column2FieldJustify.XValueInt())

		return err

	}

	if maxLineLength < -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'maxLineLength' is invalid!\n"+
			"'maxLineLength' has a value less than minus one (-1)\n"+
			"maxLineLength = '%v'\n",
			ePrefix.String(),
			maxLineLength)

		return err

	}

	newStdFmtParams := TextFmtParamsLineColumnsDto{
		FormatType: TxtFieldType.LineColumns(),
		FieldFormatParams: []TextFieldFmtParamsDto{
			{
				LeftMarginStr:  leftMarginStr,
				FieldLength:    column1FieldLength,
				FieldJustify:   column1FieldJustify,
				DateTimeFormat: "",
				RightMarginStr: column1RightMarginStr,
				lock:           nil,
			},
			{
				LeftMarginStr:  leftMarginStr,
				FieldLength:    column2FieldLength,
				FieldJustify:   column2FieldJustify,
				DateTimeFormat: "",
				RightMarginStr: column2RightMarginStr,
				lock:           nil,
			},
		},
		TurnLineTerminationOff:     turnLineTerminationOff,
		LineTerminator:             lineTerminator,
		MaxLineLength:              maxLineLength,
		TurnAutoLineLengthBreaksOn: turnAutoLineLengthBreaksOn,
		lock:                       nil,
	}

	err =
		new(textFormatterCollectionElectron).
			cfgNewStdTxtLineParameters(
				txtFmtCollection,
				newStdFmtParams,
				ePrefix.XCpy(
					"newStdFmtParams"))

	return err
}

// SetStdFormatParamsMultiCol - Sets the Standard Format Parameters
// for a Text Line consisting of one or more columns. The number of
// columns associated with this format is determined by the number
// of elements in the TextFieldFmtParamsDto array submitted as
// input parameter, 'fieldFormatParams'.
//
// These Standard Format Parameters will be applied as the default
// format for all Text Lines with same number of columns.
//
// After configuring Standard Format Parameters for this specific
// number columns, users should configure additional Text Lines
// with the same number of columns using method AddLineMultiCol()
// or AddLineManyCol() in order to reduce the number of input
// parameters required to produce a Text Line with the same number
// of columns.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//		fieldFormatParams          []TextFieldFmtParamsDto
//		   - An array of TextFieldFmtParamsDto objects containing the
//		     Text Field Format Parameters (TextFieldFmtParamsDto).
//
//		    The TextFieldFmtParamsDto type is defined as follows:
//
//		      type TextFieldFmtParamsDto struct {
//		        LeftMarginStr  string
//		        FieldLength    int
//		        FieldJustify   TextJustify
//		        DateTimeFormat string
//		        RightMarginStr string
//		      }
//
//
//		 turnLineTerminationOff    bool
//		   - When this parameter is set to 'true', no line termination
//		     sequence will be configured for this 'Line1Column' Text
//		     Line.
//
//		     Text Lines operate on the assumption that a line
//		     termination is standard operating procedure. The default
//		     line terminator for text lines is the new line character,
//		     '\n'.
//
//		     Users have the option of turning off the entire line
//		     termination sequence if this parameter is set to 'true'.
//
//
//		lineTerminator             string
//		   - If this parameter is set to an empty string, the default
//		     line terminator, a new line character '\n', will be
//		     applied to 'Line1Column' Text Lines.
//
//		     If this string is populated and the string length is
//		     greater than zero (0), an alternate line termination
//		     sequence will be configured using the characters provided
//		     in the 'lineTerminator' string.
//
//		     Remember that the application of a line termination
//		     sequence is controlled by parameter
//		     'turnLineTerminationOff'. If 'turnLineTerminationOff' is
//		     set to 'true', no line termination characters will be
//		     configured for this 'Line1Column' Text Lines.
//
//
//		maxLineLength              int
//		   - The maximum length of the line on which this 'Line1Column'
//		     Text will be presented.
//
//		     Set this parameter to minus one (-1), and no maximum line
//		     limits will be applied.
//
//		     'maxLineLength' is used in conjunction with parameter
//		     'turnAutoLineLengthBreaksOn' to automatically place text
//		     fields on separate text lines when that text exceeds the
//		     maximum text line length ('maxLineLength').
//
//		     If the value of 'maxLineLength' is less than minus one
//		     (-1), an error will be returned.
//
//		     If the value of 'maxLineLength' is zero (0), an error will
//		     be returned.
//
//
//		turnAutoLineLengthBreaksOn bool
//		   - This parameter controls whether text lines which exceed
//		     the maximum line length ('maxLineLength') are broken up
//		     and presented on the following line.
//
//		     To apply automatic line breaking at the maximum line
//		     length, set the value of this parameter to 'true'.
//
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
//
//		     This empty interface must be convertible to one of the
//		     following types:
//
//		     1. nil - A nil value is valid and generates an empty
//		        collection of error prefix and error context
//		        information.
//
//		     2. string - A string containing error prefix information.
//
//		     3. []string A one-dimensional slice of strings containing
//		        error prefix information
//
//		     4. [][2]string A two-dimensional slice of strings
//		        containing error prefix and error context information.
//
//		     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		        from this object will be copied for use in error and
//		        informational messages.
//
//		     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		        Information from this object will be copied for use in
//		        error and informational messages.
//
//		     7. IBasicErrorPrefix - An interface to a method generating
//		        a two-dimensional slice of strings containing error
//		        prefix and error context information.
//
//		     If parameter 'errorPrefix' is NOT convertible to one of
//		     the valid types listed above, it will be considered
//		     invalid and trigger the return of an error.
//
//		     Types ErrPrefixDto and IBasicErrorPrefix are included in
//		     the 'errpref' software package,
//		     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If this method completes successfully and no errors are
//	     encountered, this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (txtFmtCollection *TextFormatterCollection) SetStdFormatParamsMultiCol(
	fieldFormatParams []TextFieldFmtParamsDto,
	turnLineTerminationOff bool,
	lineTerminator string,
	maxLineLength int,
	turnAutoLineLengthBreaksOn bool,
	errorPrefix interface{}) error {

	if txtFmtCollection.lock == nil {
		txtFmtCollection.lock = new(sync.Mutex)
	}

	txtFmtCollection.lock.Lock()

	defer txtFmtCollection.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFormatterCollection."+
			"SetStdFormatParamsMultiCol()",
		"")

	if err != nil {
		return err
	}

	newStdFmtParams := TextFmtParamsLineColumnsDto{
		FormatType:                 TxtFieldType.LineColumns(),
		FieldFormatParams:          fieldFormatParams,
		TurnLineTerminationOff:     turnLineTerminationOff,
		LineTerminator:             lineTerminator,
		MaxLineLength:              maxLineLength,
		TurnAutoLineLengthBreaksOn: turnAutoLineLengthBreaksOn,
		lock:                       nil,
	}

	err =
		new(textFormatterCollectionElectron).
			cfgNewStdTxtLineParameters(
				txtFmtCollection,
				newStdFmtParams,
				ePrefix.XCpy(
					"newStdFmtParams"))

	return err

}

// SetStdFormatParamsManyCol - Sets the Standard Format Parameters
// for a Text Line consisting of one or more columns. The number of
// columns associated with this format is determined by the number
// of elements passed in  input parameter, 'fieldFormatParams'.
//
// This is a Variadic method and input parameter
// 'fieldFormatParams' accepts a variable number of
// TextFieldFmtParamsDto objects.
//
// These Standard Format Parameters will be applied as the default
// format for all Text Lines with same number of columns. The
// number of columns is determined the number of
// TextFieldFmtParamsDto objects contained in input parameter
// 'fieldFormatParams'.
//
// After configuring Standard Format Parameters for this specific
// number columns, users should configure additional Text Lines
// with the same number of columns using method AddLineMultiCol()
// or AddLineManyCol() in order to reduce the number of input
// parameters required to produce a Text Line with the same number
// of columns.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
//
//		     This empty interface must be convertible to one of the
//		     following types:
//
//		     1. nil - A nil value is valid and generates an empty
//		        collection of error prefix and error context
//		        information.
//
//		     2. string - A string containing error prefix information.
//
//		     3. []string A one-dimensional slice of strings containing
//		        error prefix information
//
//		     4. [][2]string A two-dimensional slice of strings
//		        containing error prefix and error context information.
//
//		     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		        from this object will be copied for use in error and
//		        informational messages.
//
//		     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		        Information from this object will be copied for use in
//		        error and informational messages.
//
//		     7. IBasicErrorPrefix - An interface to a method generating
//		        a two-dimensional slice of strings containing error
//		        prefix and error context information.
//
//		     If parameter 'errorPrefix' is NOT convertible to one of
//		     the valid types listed above, it will be considered
//		     invalid and trigger the return of an error.
//
//		     Types ErrPrefixDto and IBasicErrorPrefix are included in
//		     the 'errpref' software package,
//		     "github.com/MikeAustin71/errpref".
//
//
//		 turnLineTerminationOff    bool
//		   - When this parameter is set to 'true', no line termination
//		     sequence will be configured for this 'Line1Column' Text
//		     Line.
//
//		     Text Lines operate on the assumption that a line
//		     termination is standard operating procedure. The default
//		     line terminator for text lines is the new line character,
//		     '\n'.
//
//		     Users have the option of turning off the entire line
//		     termination sequence if this parameter is set to 'true'.
//
//
//		lineTerminator             string
//		   - If this parameter is set to an empty string, the default
//		     line terminator, a new line character '\n', will be
//		     applied to 'Line1Column' Text Lines.
//
//		     If this string is populated and the string length is
//		     greater than zero (0), an alternate line termination
//		     sequence will be configured using the characters provided
//		     in the 'lineTerminator' string.
//
//		     Remember that the application of a line termination
//		     sequence is controlled by parameter
//		     'turnLineTerminationOff'. If 'turnLineTerminationOff' is
//		     set to 'true', no line termination characters will be
//		     configured for this 'Line1Column' Text Lines.
//
//
//		maxLineLength              int
//		   - The maximum length of the line on which this 'Line1Column'
//		     Text will be presented.
//
//		     Set this parameter to minus one (-1), and no maximum line
//		     limits will be applied.
//
//		     'maxLineLength' is used in conjunction with parameter
//		     'turnAutoLineLengthBreaksOn' to automatically place text
//		     fields on separate text lines when that text exceeds the
//		     maximum text line length ('maxLineLength').
//
//		     If the value of 'maxLineLength' is less than minus one
//		     (-1), an error will be returned.
//
//		     If the value of 'maxLineLength' is zero (0), an error will
//		     be returned.
//
//
//		turnAutoLineLengthBreaksOn bool
//		   - This parameter controls whether text lines which exceed
//		     the maximum line length ('maxLineLength') are broken up
//		     and presented on the following line.
//
//		     To apply automatic line breaking at the maximum line
//		     length, set the value of this parameter to 'true'.
//
//
//		fieldFormatParams          []TextFieldFmtParamsDto
//		   - An array of TextFieldFmtParamsDto objects containing the
//		     Text Field Format Parameters (TextFieldFmtParamsDto).
//
//		    The TextFieldFmtParamsDto type is defined as follows:
//
//		      type TextFieldFmtParamsDto struct {
//		        LeftMarginStr  string
//		        FieldLength    int
//		        FieldJustify   TextJustify
//		        DateTimeFormat string
//		        RightMarginStr string
//		      }
//
// ----------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If this method completes successfully and no errors are
//	     encountered, this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (txtFmtCollection *TextFormatterCollection) SetStdFormatParamsManyCol(
	errorPrefix interface{},
	turnLineTerminationOff bool,
	lineTerminator string,
	maxLineLength int,
	turnAutoLineLengthBreaksOn bool,
	fieldFormatParams ...TextFieldFmtParamsDto) error {

	if txtFmtCollection.lock == nil {
		txtFmtCollection.lock = new(sync.Mutex)
	}

	txtFmtCollection.lock.Lock()

	defer txtFmtCollection.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFormatterCollection."+
			"SetStdFormatParamsManyCol()",
		"")

	if err != nil {
		return err
	}

	fieldFmtParameters := make([]TextFieldFmtParamsDto, 0)

	for _, textFieldFmtParameter := range fieldFormatParams {

		fieldFmtParameters = append(
			fieldFmtParameters,
			textFieldFmtParameter)
	}

	newStdFmtParams := TextFmtParamsLineColumnsDto{
		FormatType:                 TxtFieldType.LineColumns(),
		FieldFormatParams:          fieldFmtParameters,
		TurnLineTerminationOff:     turnLineTerminationOff,
		LineTerminator:             lineTerminator,
		MaxLineLength:              maxLineLength,
		TurnAutoLineLengthBreaksOn: turnAutoLineLengthBreaksOn,
		lock:                       nil,
	}

	lenFmtParams := len(fieldFmtParameters)

	err = new(textFormatterCollectionElectron).
		cfgNewStdTxtLineParameters(
			txtFmtCollection,
			newStdFmtParams,
			ePrefix.XCpy(
				fmt.Sprintf("newStdFmtParams %v-Columns",
					lenFmtParams)))

	return err
}

// SetStdFormatParams - Sets the Standard Format Parameters
// for a Text Line consisting of one or more columns. The number of
// columns associated with this format is determined by the number
// of elements passed in  input parameter, 'fieldFormatParams'.
//
// The Text Field Format Parameters are passed as an input
// parameter of type TextFmtParamsLineColumnsDto.
//
// These Standard Format Parameters will be applied as the default
// format for all Text Lines with same number of columns. The
// number of columns is determined the number of
// TextFieldFmtParamsDto objects contained in input parameter
// 'fieldFormatParams'.
//
// After configuring Standard Format Parameters for this specific
// number columns, users should configure additional Text Lines
// with the same number of columns using method AddLineMultiCol()
// or AddLineManyCol() in order to reduce the number of input
// parameters required to produce a Text Line with the same number
// of columns.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
//
//		     This empty interface must be convertible to one of the
//		     following types:
//
//		     1. nil - A nil value is valid and generates an empty
//		        collection of error prefix and error context
//		        information.
//
//		     2. string - A string containing error prefix information.
//
//		     3. []string A one-dimensional slice of strings containing
//		        error prefix information
//
//		     4. [][2]string A two-dimensional slice of strings
//		        containing error prefix and error context information.
//
//		     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		        from this object will be copied for use in error and
//		        informational messages.
//
//		     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		        Information from this object will be copied for use in
//		        error and informational messages.
//
//		     7. IBasicErrorPrefix - An interface to a method generating
//		        a two-dimensional slice of strings containing error
//		        prefix and error context information.
//
//		     If parameter 'errorPrefix' is NOT convertible to one of
//		     the valid types listed above, it will be considered
//		     invalid and trigger the return of an error.
//
//		     Types ErrPrefixDto and IBasicErrorPrefix are included in
//		     the 'errpref' software package,
//		     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If this method completes successfully and no errors are
//	     encountered, this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (txtFmtCollection *TextFormatterCollection) SetStdFormatParams(
	txtFmtParams TextFmtParamsLineColumnsDto,
	errorPrefix interface{}) error {

	if txtFmtCollection.lock == nil {
		txtFmtCollection.lock = new(sync.Mutex)
	}

	txtFmtCollection.lock.Lock()

	defer txtFmtCollection.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFormatterCollection."+
			"SetStdFormatParamsManyCol()",
		"")

	if err != nil {
		return err
	}

	if txtFmtParams.GetNumOfFieldFmtParams() == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtFmtParams' is invalid!\n"+
			"The Text Field Format Parameters array is empty.\n",
			ePrefix.String())

		return err
	}

	err = new(textFormatterCollectionElectron).
		cfgNewStdTxtLineParameters(
			txtFmtCollection,
			txtFmtParams,
			ePrefix.XCpy(
				"txtFmtCollection<-txtFmtParams"))

	return err
}
