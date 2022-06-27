package strmech

import (
	"sync"
	"time"
)

// TextLabelParameterValueFieldDto - The Text Label Parameter Value
// Field Data Transfer Object specifies the parameters necessary to
// construct a formatted line of text consisting of a left margin,
// text label field, text label right margin, parameter value field,
// parameter value right margin and a line termination sequence.
//
// Example Layout
//
//         " " + "Inflation Rate" + ": " + "8.3%"+"\n"
//
// Be advised that the actual Parameter Value Text Field can be
// populated with one of two values: a Date/Time Value or a Text
// String.
//
// If the 'ParamValueDateTime' field is set to a value greater than
// zero, the Parameter Value Text Field will be formatted as a
// Date/Time string. Otherwise, if the 'ParamValueDateTime' field
// is set to zero, the Parameter Value Text Field will be
// constructed using the string value of 'ParamValueStr'.
//
type TextLabelParameterValueFieldDto struct {
	LeftMarginStr string
	// The contents of the string will be used as the left margin
	// for the Label Field.
	//
	// If no left margin is required, set 'LeftMarginStr' to a zero
	// length or empty string, and no left margin will be created.

	ParamLabelStr string
	// This string represents the contents of the Parameter Label.
	// If this string is empty (has a zero (0) length), it will be
	// skipped an ignored.
	//
	// The 'ParamLabelStr' field should be used to provide narrative
	// text describing the Parameter Value displayed in the
	// 'ParamValueStrStr' field.

	ParamLabelLength int
	// Used to format 'ParamLabelStr' field. This is the length of the
	// text field in which the 'ParamLabelStr' will be displayed. If
	// 'ParamLabelLength' is less than the length of the
	// 'ParamLabelStr' string, it will be automatically set equal to
	// the 'ParamLabelStr' string length.
	//
	// To automatically set the value of 'ParamLabelLength' to the
	// length of 'ParamLabelStr', set this parameter to a value of
	// minus one (-1).
	//
	// If 'ParamLabelLength' is submitted with a value less than
	// minus one (-1) or greater than 1-million (1,000,000), an
	// error will be returned.
	//

	ParamLabelJustify TextJustify
	// An enumeration which specifies the justification of the
	// 'ParamLabelStr' string within the text field specified by
	// 'ParamLabelLength'.
	//
	// Text justification can only be evaluated in the context of
	// a text label, field length and a Text Justification object
	// of type TextJustify. This is because text labels with a
	// field length equal to or less than the length of the text
	// label never use text justification. In these cases, text
	// justification is completely ignored.
	//
	// If the field length is greater than the length of the text
	// label, text justification must be equal to one of these
	// three valid values:
	//     TextJustify(0).Left()
	//     TextJustify(0).Right()
	//     TextJustify(0).Center()
	//
	// You can also use the abbreviated text justification
	// enumeration syntax as follows:
	//
	//     TxtJustify.Left()
	//     TxtJustify.Right()
	//     TxtJustify.Center()

	ParamLabelRightMarginStr string
	// The contents of the string will be used as the right margin
	// for the 'ParamLabelStr' Field.
	//
	// If no Parameter Label right margin is required, set
	// 'ParamLabelRightMarginStr' to a zero length or empty string,
	// and no Parameter Label right margin will be created.

	ParamValueDateTime time.Time
	// If 'ParamValueDateTime' is populated with a value greater
	// than zero, the Parameter value will be formatted as at
	// Date/Time value using the 'ParamValueDateTimeFormat' string.
	//
	// If 'ParamValueDateTime' is set equal to zero, this field
	// will be skipped and ignored and the 'ParamValueStr' field
	// will be used to construct the Parameter value.
	//

	ParamValueDateTimeFormat string
	// If 'ParamValueDateTime' is set to a value greater than zero,
	// this field will be used to format 'ParamValueDateTime' as a
	// string for text output.
	//
	// If 'ParamValueDateTime' is set to a value greater than zero
	// and this 'ParamValueDateTimeFormat' string is empty (has a
	// zero length), a default Date/Time format string will be
	// applied as follows:
	//         "2006-01-02 15:04:05.000000000 -0700 MST"

	ParamValueStr string
	// The Parameter Value formatted as a string. If
	// 'ParamValueDateTime' is set equal to zero (0),
	// 'ParamValueStr' will be used to populate the Parameter Value
	// field. This string will be formatted as a TextFieldSpecLabel
	// and formatted for text output.

	ParamValueLength int
	// Used to format Parameter Value Text Field. This is the
	// string length of the text field in which the Parameter
	// Value will be displayed. If 'ParamValueLength' is less than
	// the length of the Parameter Value Text Field string, it will
	// be automatically set equal to the Parameter Value Text Field
	// string length.
	//
	// To automatically set the value of 'ParamValueLength' to the
	// length of the Parameter Value Text Field, set this parameter
	// to a value of minus one (-1).
	//
	// If 'ParamValueLength' is submitted with a value less than
	// minus one (-1) or greater than 1-million (1,000,000), an
	// error will be returned.

	ParamValueJustify TextJustify
	// An enumeration which specifies the justification of the
	// Parameter Value Text Field string within the text field
	// specified by 'ParamValueLength'.
	//
	// Text justification can only be evaluated in the context of
	// a text string, field length and a Text Justification object
	// of type TextJustify. This is because text strings with a
	// field length equal to or less than the length of the text
	// string never use text justification. In these cases, text
	// justification is completely ignored.
	//
	// If the field length is greater than the length of the text
	// string, text justification must be equal to one of these
	// three valid values:
	//     TextJustify(0).Left()
	//     TextJustify(0).Right()
	//     TextJustify(0).Center()
	//
	// You can also use the abbreviated text justification
	// enumeration syntax as follows:
	//
	//     TxtJustify.Left()
	//     TxtJustify.Right()
	//     TxtJustify.Center()

	ParamValueRightMarginStr string
	// The contents of the string will be used as the right margin
	// for the Parameter Value text field.
	//
	// If no right margin is required, set
	// 'ParamValueRightMarginStr' to a zero length or empty string,
	// and no right margin will be created.

	LineTerminator string
	// This string holds the character or characters which will be
	// used to terminate the formatted line of text output.
	//
	// The most common usage sets this string to a new line
	// character ("\n").
	//
	// If no Line Terminator is required, set 'lineTerminator' to
	// a zero length or empty string and no line termination
	// characters will be created.

	lock *sync.Mutex
}

func (txtLabelParamValueDto *TextLabelParameterValueFieldDto) CopyIn() {

	if txtLabelParamValueDto.lock == nil {
		txtLabelParamValueDto.lock = new(sync.Mutex)
	}

	txtLabelParamValueDto.lock.Lock()

	defer txtLabelParamValueDto.lock.Unlock()

}

// CopyOut - Returns a deep copy of the current
// TextLabelParameterValueFieldDto instance.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// This method does NOT perform validation checks on the current
// instance of TextLabelParameterValueFieldDto before creating and
// returning the deep copy of this instance.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  NONE
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  TextLabelParameterValueFieldDto
//     - If this method completes successfully and no errors are
//       encountered, this parameter will return a deep copy of the
//       current TextLabelParameterValueFieldDto instance.
//
func (txtLabelParamValueDto *TextLabelParameterValueFieldDto) CopyOut() TextLabelParameterValueFieldDto {

	if txtLabelParamValueDto.lock == nil {
		txtLabelParamValueDto.lock = new(sync.Mutex)
	}

	txtLabelParamValueDto.lock.Lock()

	defer txtLabelParamValueDto.lock.Unlock()

	var newCopyTxtLabelParamDto TextLabelParameterValueFieldDto

	newCopyTxtLabelParamDto.LeftMarginStr =
		txtLabelParamValueDto.LeftMarginStr

	newCopyTxtLabelParamDto.ParamLabelStr =
		txtLabelParamValueDto.ParamLabelStr

	newCopyTxtLabelParamDto.ParamLabelLength =
		txtLabelParamValueDto.ParamLabelLength

	newCopyTxtLabelParamDto.ParamLabelJustify =
		txtLabelParamValueDto.ParamLabelJustify

	newCopyTxtLabelParamDto.ParamLabelRightMarginStr =
		txtLabelParamValueDto.ParamLabelRightMarginStr

	newCopyTxtLabelParamDto.ParamValueDateTime =
		txtLabelParamValueDto.ParamValueDateTime

	newCopyTxtLabelParamDto.ParamValueDateTimeFormat =
		txtLabelParamValueDto.ParamValueDateTimeFormat

	newCopyTxtLabelParamDto.ParamValueStr =
		txtLabelParamValueDto.ParamValueStr

	newCopyTxtLabelParamDto.ParamValueLength =
		txtLabelParamValueDto.ParamValueLength

	newCopyTxtLabelParamDto.ParamValueJustify =
		txtLabelParamValueDto.ParamValueJustify

	newCopyTxtLabelParamDto.ParamValueRightMarginStr =
		txtLabelParamValueDto.ParamValueRightMarginStr

	newCopyTxtLabelParamDto.LineTerminator =
		txtLabelParamValueDto.LineTerminator

	return newCopyTxtLabelParamDto
}

// Empty - Resets all internal member variables for the current
// instance of TextLabelParameterValueFieldDto to their zero or
// uninitialized states.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// This method will delete all member variable data values in this
// current instance of TextLabelParameterValueFieldDto. All member
// variable data values will be reset to their zero or
// uninitialized states.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  NONE
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  NONE
//
func (txtLabelParamValueDto *TextLabelParameterValueFieldDto) Empty() {

	if txtLabelParamValueDto.lock == nil {
		txtLabelParamValueDto.lock = new(sync.Mutex)
	}

	txtLabelParamValueDto.lock.Lock()

	txtLabelParamValueDto.LeftMarginStr = ""

	txtLabelParamValueDto.ParamLabelStr = ""

	txtLabelParamValueDto.ParamLabelLength = 0

	txtLabelParamValueDto.ParamLabelJustify = TxtJustify.None()

	txtLabelParamValueDto.ParamLabelRightMarginStr = ""

	txtLabelParamValueDto.ParamValueDateTime = time.Time{}

	txtLabelParamValueDto.ParamValueDateTimeFormat = ""

	txtLabelParamValueDto.ParamValueStr = ""

	txtLabelParamValueDto.ParamValueLength = 0

	txtLabelParamValueDto.ParamValueJustify = TxtJustify.None()

	txtLabelParamValueDto.ParamValueRightMarginStr = ""

	txtLabelParamValueDto.LineTerminator = ""

	txtLabelParamValueDto.ParamLabelStr = ""

	txtLabelParamValueDto.ParamLabelLength = 0

	txtLabelParamValueDto.ParamLabelJustify = TxtJustify.None()

	txtLabelParamValueDto.ParamLabelRightMarginStr = ""

	txtLabelParamValueDto.ParamValueDateTime = time.Time{}

	txtLabelParamValueDto.ParamValueDateTimeFormat = ""

	txtLabelParamValueDto.ParamValueStr = ""

	txtLabelParamValueDto.ParamValueLength = 0

	txtLabelParamValueDto.ParamValueJustify = TxtJustify.None()

	txtLabelParamValueDto.ParamValueRightMarginStr = ""

	txtLabelParamValueDto.LineTerminator = ""

	txtLabelParamValueDto.lock.Unlock()

	txtLabelParamValueDto.lock = nil

}
