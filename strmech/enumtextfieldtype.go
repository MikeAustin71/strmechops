package strmech

import (
	"fmt"
	"strings"
	"sync"
)

// Do NOT access these maps without first getting
// the lock on 'lockTextFieldType'.

var mTextFieldTypeCodeToString = map[TextFieldType]string{
	TextFieldType(0): "None",
	TextFieldType(1): "Label",
	TextFieldType(2): "DateTime",
	TextFieldType(3): "Filler",
	TextFieldType(4): "Spacer",
	TextFieldType(5): "BlankLine",
	TextFieldType(6): "SolidLine",
	TextFieldType(7): "LineColumns",
	TextFieldType(8): "TimerStartStop",
	TextFieldType(9): "TextAdHoc",
}

var mTextFieldTypeStringToCode = map[string]TextFieldType{
	"None":           TextFieldType(0),
	"Label":          TextFieldType(1),
	"DateTime":       TextFieldType(2),
	"Date Time":      TextFieldType(2),
	"Date":           TextFieldType(2),
	"Filler":         TextFieldType(3),
	"Spacer":         TextFieldType(4),
	"BlankLine":      TextFieldType(5),
	"SolidLine":      TextFieldType(6),
	"LineColumns":    TextFieldType(7),
	"TimerStartStop": TextFieldType(8),
	"TextAdHoc":      TextFieldType(9),
}

var mTextFieldTypeLwrCaseStringToCode = map[string]TextFieldType{
	"none":           TextFieldType(0),
	"label":          TextFieldType(1),
	"datetime":       TextFieldType(2),
	"date time":      TextFieldType(2),
	"date":           TextFieldType(2),
	"filler":         TextFieldType(3),
	"spacer":         TextFieldType(4),
	"blankline":      TextFieldType(5),
	"solidline":      TextFieldType(6),
	"linecolumns":    TextFieldType(7),
	"timerstartstop": TextFieldType(8),
	"textadhoc":      TextFieldType(9),
}

// TextFieldType - The 'Text Field Type' is an enumeration of type
// codes used to identify a Text Field Specification like
// TextFieldSpecLabel, TextFieldSpecDateTime, TextFieldSpecFiller
// or TextFieldSpecSpacer. Using this enumeration to identify Text
// Field Specifications when passing parameters to methods enhances
// flexibility and efficiency in text formatting operations.
//
// ----------------------------------------------------------------
//
// TERMINOLOGY
//
// Text Field Specifications are used to format lines to text. They
// are designed to be configured as elements within a line of text.
// Those text lines can then be formatted for text displays,
// file output or printing.
//
// Type TextLineSpecStandardLine can be used to compose a line of
// text consisting of multiple Text Field Specifications like
// TextFieldSpecLabel, TextFieldSpecDateTime, TextFieldSpecFiller
// or TextFieldSpecSpacer. Text Field Specifications are therefore
// used as the components or building blocks for constructing
// single lines of formatted text.
//
// The TextFieldType enumeration is used to identify Text Field
// Specifications is in text formatting operations.
//
//
// ----------------------------------------------------------------
//
// Type TextFieldType is styled as an enumeration. Since the Go
// Programming Language does not directly support enumerations,
// type TextFieldType has been adapted to function in a manner
// similar to classic enumerations.
//
// TextFieldType is declared as a type 'int' and includes two
// types of methods:
//    Enumeration Methods
//          and
//    Utility Methods
//
// Enumeration methods have names which collectively represent an
// enumeration of different Text Field Specifications necessary
// for text formatting operations.
//    Examples Of Enumeration Method Names:
//        Label()
//        DateTime()
//        Filler()
//        Spacer()
//
//  Enumeration methods return an integer value used to designate
//  a specific Text Field Specification.
//
//  Utility methods make up the second type of method included in
//  TextFieldType. These methods are NOT part of the enumeration
//  but instead provide needed supporting services. All
//  utility methods, with the sole exception of method String(),
//  have names beginning with 'X' to separate them from standard
//  enumeration methods.
//    Examples:
//      XIsValid()
//      XParseString()
//      XValue()
//      XValueInt()
//
//  The utility method 'String()' supports the Stringer Interface
//  and is not part of the standard enumeration.
//
// ----------------------------------------------------------------
//
// Enumeration Methods
//
// The TextFieldType enumeration methods are described below:
//
// Method                   Integer
//  Name                     Value
// ------                   -------
//
// None                     Zero (0)
//  - Signals that the Text Field Type is empty and not
//    initialized. This is an invalid or error condition.
//
//
// Label                        1
//  - Identifies the Text Field Specification as a type
//    TextFieldSpecLabel.
//
//    The Label Text Field Specification is used to generate a text
//    label string. The text label is positioned inside a text
//    field with a given field length. Text Justification within
//    this text field is controlled by the Text Justification
//    specification value which may be set to 'Left', 'Right' or
//    'Center'. A text label contains a string of text characters.
//
//
// DateTime                     2
//  - Identifies the Text Field Specification as a type
//    TextFieldSpecDateTime.
//
//    The Date Time Text Field specification is used to produce a
//    formatted text string from a date/time value.
//
//
// Filler                       3
//  - Identifies the Text Field Specification as a type
//    TextFieldSpecFiller.
//
//    A Filler Text Field Specification is a single character or
//    character sequence which is replicated multiple times to
//    create the entire length of the Filler Text Field.
//
//
// Spacer                       4
//  - Identifies the Text Field Specification as a type
//    TextFieldSpecSpacer.
//
//    The Spacer Text Field Specification is used to create a Text
//    Field consisting of one or more white space characters (" ").
//
//
// BlankLine                    5
//  - Identifies a type TextLineSpecBlankLines which is used to
//    generate Blank Lines of text.
//
//
// SolidLine                    6
//  - Identifies a Solid Line Specification consisting of a left
//    margin, a 'Filler' field made up of a single or repeating
//    character sequence, a right margin and a line termination
//    character sequence. This type of line is implemented using
//    the TextLineSpecSolidLine Specification.
//
//
//  LineColumns                 7
//  - Identifies a Text Line consisting of one or more text fields
//    or text columns.
//
//    The line/column architecture differs from single text fields
//    in that text lines includes margins on both sides of the
//    column in addition to providing input parameters for
//    line-termination characters such as new line characters
//    ('\n').
//
//  TimerStartStop              8
//  - Identifies a display of four text lines presenting the
//    results of a timer event. This timer event is described with
//    a start time, stop time and time duration or elapsed time.
//
//    The first line of text shows the Starting Time. The second
//    line shows the Ending Time. The third line displays the time
//    duration or the difference between starting time and ending
//    time. The fourth line displays the total elapsed time in
//    nanoseconds.
//
//    For more information reference type TextLineSpecTimerLines.
//
//  TextAdHoc                   9
//  - Identifies a string of ad hoc text which will be inserted
//    into the stream of formatted text as is, without any
//    additional formatting being applied.
//
//
// ----------------------------------------------------------------
//
// USAGE
//
// For easy access to these enumeration values, use the global
// constant 'TxtFieldType'.
//     Example: TxtFieldType.Label()
//
// Otherwise you will need to use the formal syntax.
//     Example: TextFieldType(0).Label()
//
// Depending on your editor, intellisense (a.k.a. intelligent code
// completion) may not list the TextFieldType methods in
// alphabetical order.
//
// Be advised that all 'TextFieldType' methods beginning with
// 'X', as well as the method 'String()', are utility methods and
// not part of the enumeration.
//
type TextFieldType int

var lockTextFieldType sync.Mutex

// None - Signals that the TextFieldType specification is
// empty or uninitialized. This is an invalid or error condition.
//
// This method is part of the standard enumeration.
//
func (txtFieldType TextFieldType) None() TextFieldType {

	lockTextFieldType.Lock()

	defer lockTextFieldType.Unlock()

	return TextFieldType(0)
}

// Label - Identifies the Text Field Specification as a type
// TextFieldSpecLabel.
//
// The Label Text Field Specification is used to generate a text
// label string. The text label is positioned inside a text
// field with a given field length. Text Justification within
// this text field is controlled by the Text Justification
// specification value which may be set to 'Left', 'Right' or
// 'Center'. A text label contains a string of text characters.
//
// This method is part of the standard enumeration.
//
func (txtFieldType TextFieldType) Label() TextFieldType {

	lockTextFieldType.Lock()

	defer lockTextFieldType.Unlock()

	return TextFieldType(1)
}

// DateTime - Identifies the Text Field Specification as a type
// TextFieldSpecDateTime.
//
// The Date Time Text Field specification is used to produce a
// formatted text string from a date/time value.
//
//
// This method is part of the standard enumeration.
//
func (txtFieldType TextFieldType) DateTime() TextFieldType {

	lockTextFieldType.Lock()

	defer lockTextFieldType.Unlock()

	return TextFieldType(2)
}

// Filler - Identifies the Text Field Specification as a type
// TextFieldSpecFiller.
//
// A Filler Text Field Specification is a single character or
// character sequence which is replicated multiple times to create
// the entire length of the Filler Text Field.
//
// This method is part of the standard enumeration.
//
func (txtFieldType TextFieldType) Filler() TextFieldType {

	lockTextFieldType.Lock()

	defer lockTextFieldType.Unlock()

	return TextFieldType(3)
}

// Spacer - Identifies the Text Field Specification as a type
// TextFieldSpecSpacer.
//
// A Spacer Text Field Specification is used to create a Text
// Field consisting of one or more white space characters (" ").
//
// This method is part of the standard enumeration.
//
func (txtFieldType TextFieldType) Spacer() TextFieldType {

	lockTextFieldType.Lock()

	defer lockTextFieldType.Unlock()

	return TextFieldType(4)
}

// BlankLine - Identifies a Blank Line Specification as a type
// TextLineSpecBlankLines.
//
// A Blank Line Specification is used to create one or more blank
// lines of text. This type commonly employs the new line
// character "\n" to generate blank or empty text lines.
//
func (txtFieldType TextFieldType) BlankLine() TextFieldType {

	lockTextFieldType.Lock()

	defer lockTextFieldType.Unlock()

	return TextFieldType(5)
}

// SolidLine - Identifies a Solid Line Specification consisting
// of a left margin, a 'Filler' field made up of a single
// or repeating character sequence, a right margin and a line
// termination character sequence. This type of line is implemented
// using the TextLineSpecSolidLine Specification.
//
func (txtFieldType TextFieldType) SolidLine() TextFieldType {

	lockTextFieldType.Lock()

	defer lockTextFieldType.Unlock()

	return TextFieldType(6)
}

// LineColumns - Identifies a Text Line consisting of one or more
// text fields or text columns.
//
// The line/column architecture differs from single text fields
// in that text lines includes margins on both sides of the column
// in addition to providing input parameters for line-termination
// characters such as new line characters ('\n').
//
func (txtFieldType TextFieldType) LineColumns() TextFieldType {

	lockTextFieldType.Lock()

	defer lockTextFieldType.Unlock()

	return TextFieldType(7)
}

// TimerStartStop - Identifies a display of four text lines
// presenting the results of a timer event. This timer event is
// described with a start time, stop time and time duration or
// elapsed time.
//
// The first line of text shows the Starting Time. The second line
// shows the Ending Time. The third line displays the time duration
// or the difference between starting time and ending time. The
// fourth line displays the total elapsed time in nanoseconds.
//
// For more information reference type TextLineSpecTimerLines.
//
func (txtFieldType TextFieldType) TimerStartStop() TextFieldType {

	lockTextFieldType.Lock()

	defer lockTextFieldType.Unlock()

	return TextFieldType(8)
}

// TextAdHoc - Identifies a string of ad hoc text which will be
// inserted into the stream of formatted text as is, without any
// additional formatting being applied.
//
func (txtFieldType TextFieldType) TextAdHoc() TextFieldType {

	lockTextFieldType.Lock()

	defer lockTextFieldType.Unlock()

	return TextFieldType(9)
}

// String - Returns a string with the name of the enumeration
// associated with this current instance of 'TextFieldType'.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ----------------------------------------------------------------
//
// Usage
//
// t:= TextFieldType(0).Label()
// str := t.String()
//     str is now equal to 'Label'
//
func (txtFieldType TextFieldType) String() string {

	lockTextFieldType.Lock()

	defer lockTextFieldType.Unlock()

	result, ok := mTextFieldTypeCodeToString[txtFieldType]

	if !ok {

		return "Error: Text Field Type Specification UNKNOWN!"

	}

	return result
}

// XReturnNoneIfInvalid - Provides a standardized value for invalid
// instances of enumeration TextFieldType.
//
// If the current instance of TextFieldType is invalid, this
// method will always return a value of TextFieldType(0).None().
//
// Background
//
// Enumeration TextFieldType has an underlying type of integer
// (int). This means the type could conceivably be set to any
// integer value. This method ensures that all invalid
// TextFieldType instances are consistently classified as 'None'
// (TextFieldType(0).None()). Remember that 'None' is considered
// an invalid value.
//
// For example, assume that TextFieldType was set to an integer
// value of -848972. Calling this method on a TextFieldType with
// this invalid integer value will return an integer value of zero
// or the equivalent of TextFieldType(0).None(). This conversion is
// useful in generating text strings for meaningful informational
// and error messages.
//
func (txtFieldType TextFieldType) XReturnNoneIfInvalid() TextFieldType {

	lockTextFieldType.Lock()

	defer lockTextFieldType.Unlock()

	isValid := textFieldTypeNanobot{}.ptr().
		isValidTextField(txtFieldType)

	if !isValid {
		return TextFieldType(0).None()
	}

	return txtFieldType
}

// XIsValid - Returns a boolean value signaling whether the current
// TextFieldType value is valid.
//
// Be advised, the enumeration value "None" is considered NOT
// VALID. "None" represents an error condition.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
//  textType :=
// 			TextFieldType(0).Label()
//
//  isValid := textType.XIsValid() // isValid == true
//
//  textType = TextFieldType(0).None()
//
//  isValid = textType.XIsValid() // isValid == false
//
func (txtFieldType TextFieldType) XIsValid() bool {

	lockTextFieldType.Lock()

	defer lockTextFieldType.Unlock()

	txtFieldNanobot := textFieldTypeNanobot{}

	return txtFieldNanobot.isValidTextField(
		txtFieldType)
}

// XParseString - Receives a string and attempts to match it with
// the string value of a supported enumeration. If successful, a
// new instance of TextFieldType is returned set to the value
// of the associated enumeration.
//
// This is a standard utility method and is NOT part of the valid
// enumerations for this type.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
// valueString   string
//     - A string which will be matched against the enumeration string
//       values. If 'valueString' is equal to one of the enumeration
//       names, this method will proceed to successful completion and
//       return the correct enumeration value.
//
// caseSensitive   bool
//     - If 'true' the search for enumeration names will be
//       case-sensitive and will require an exact match. Therefore,
//       'label' WILL NOT match the enumeration name, 'Label'.
//
//       A case-sensitive search will match any of the following
//       strings:
//           "None"
//           "Label"
//           "DateTime"
//           "Date Time"
//           "Date"
//           "Filler"
//           "Spacer"
//           "BlankLine"
//           "SolidLine"
//           "LineColumns"
//           "TimerStartStop"
//           "TextAdHoc"
//
//       If 'false', a case-insensitive search is conducted for the
//       enumeration name. In this example, 'label'
//       WILL MATCH the enumeration name, 'Label'.
//
//       A case-insensitive search will match any of the following
//       lower case names:
//           "none"
//           "label"
//           "datetime"
//           "date time"
//           "date"
//           "filler"
//           "spacer"
//           "blankline"
//           "solidline"
//           "linecolumns"
//           "timerstartstop"
//           "textadhoc"
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  TextFieldType
//     - Upon successful completion, this method will return a new
//       instance of TextFieldType set to the value of the
//       enumeration matched by the string search performed on
//       input parameter, 'valueString'.
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If an error condition is
//       encountered, this method will return an error type which
//       encapsulates an appropriate error message.
//
//
// ----------------------------------------------------------------
//
// Usage
//
//  t, err := TextFieldType(0).
//               XParseString("Label", true)
//
//  t is now equal to TextFieldType(0).Label()
//
func (txtFieldType TextFieldType) XParseString(
	valueString string,
	caseSensitive bool) (
	TextFieldType,
	error) {

	lockTextFieldType.Lock()

	defer lockTextFieldType.Unlock()

	ePrefix := "TextFieldType.XParseString() "

	if len(valueString) < 4 {
		return TextFieldType(0),
			fmt.Errorf(ePrefix+
				"\nInput parameter 'valueString' is INVALID!\n"+
				"String length is less than '4'.\n"+
				"valueString='%v'\n", valueString)
	}

	var ok bool
	var enumTextFieldType TextFieldType

	if caseSensitive {

		enumTextFieldType, ok =
			mTextFieldTypeStringToCode[valueString]

		if !ok {
			return TextFieldType(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid TextFieldType Specification.\n"+
					"valueString='%v'\n", valueString)
		}

	} else {

		enumTextFieldType, ok =
			mTextFieldTypeLwrCaseStringToCode[strings.ToLower(valueString)]

		if !ok {
			return TextFieldType(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid TextFieldType Specification.\n"+
					"valueString='%v'\n", valueString)
		}
	}

	return enumTextFieldType, nil
}

// XValue - This method returns the enumeration value of the
// current TextFieldType instance.
//
// This is a standard utility method and is NOT part of the valid
// enumerations for this type.
//
func (txtFieldType TextFieldType) XValue() TextFieldType {

	lockTextFieldType.Lock()

	defer lockTextFieldType.Unlock()

	return txtFieldType
}

// XValueInt - This method returns the integer value of the current
// TextFieldType instance.
//
// This is a standard utility method and is NOT part of the valid
// enumerations for this type.
//
func (txtFieldType TextFieldType) XValueInt() int {

	lockTextFieldType.Lock()

	defer lockTextFieldType.Unlock()

	return int(txtFieldType)
}

// TxtFieldType - public global constant of type TextFieldType.
//
// This variable serves as an easier, shorthand technique for
// accessing TextFieldType values.
//
// For easy access to these enumeration values, use the
// global variable TxtFieldType.
//  Example: TxtFieldType.Label()
//
// Otherwise you will need to use the formal syntax.
//  Example: TextFieldType(0).Label()
//
// Usage:
//  TxtFieldType.None()
//  TxtFieldType.Label()
//  TxtFieldType.DateTime()
//  TxtFieldType.Filler()
//  TxtFieldType.Spacer()
//  TxtFieldType.BlankLine()
//  TxtFieldType.SolidLine()
//  TxtFieldType.LineColumns()
//  TxtFieldType.TimerStartStop()
//
const TxtFieldType = TextFieldType(0)

// textFieldTypeNanobot - Provides helper methods for
// enumeration TextFieldType.
//
type textFieldTypeNanobot struct {
	lock *sync.Mutex
}

// isValidTextField - Receives an instance of TextFieldType and
// returns a boolean value signaling whether that TextFieldType
// instance is valid.
//
// If the passed instance of TextFieldType is valid, this method
// returns 'true'.
//
// Be advised, the enumeration value "None" is considered NOT
// VALID. "None" represents an error condition.
//
// This is a standard utility method and is not part of the valid
// TextFieldType enumeration.
//
func (textFieldNanobot *textFieldTypeNanobot) isValidTextField(
	textFieldType TextFieldType) bool {

	if textFieldNanobot.lock == nil {
		textFieldNanobot.lock = new(sync.Mutex)
	}

	textFieldNanobot.lock.Lock()

	defer textFieldNanobot.lock.Unlock()

	if textFieldType < 1 ||
		textFieldType > 9 {

		return false
	}

	return true
}

// ptr - Returns a pointer to a new instance of
// textFieldTypeNanobot.
//
func (textFieldNanobot textFieldTypeNanobot) ptr() *textFieldTypeNanobot {

	if textFieldNanobot.lock == nil {
		textFieldNanobot.lock = new(sync.Mutex)
	}

	textFieldNanobot.lock.Lock()

	defer textFieldNanobot.lock.Unlock()

	return &textFieldTypeNanobot{
		lock: new(sync.Mutex),
	}
}
