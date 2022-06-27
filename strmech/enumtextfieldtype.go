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
}

var mTextFieldTypeStringToCode = map[string]TextFieldType{
	"None":      TextFieldType(0),
	"Label":     TextFieldType(1),
	"DateTime":  TextFieldType(2),
	"Date Time": TextFieldType(2),
	"Date":      TextFieldType(2),
	"Filler":    TextFieldType(3),
	"Spacer":    TextFieldType(4),
	"BlankLine": TextFieldType(5),
}
var mTextFieldTypeLwrCaseStringToCode = map[string]TextFieldType{
	"none":      TextFieldType(0),
	"label":     TextFieldType(1),
	"datetime":  TextFieldType(2),
	"date time": TextFieldType(2),
	"date":      TextFieldType(2),
	"filler":    TextFieldType(3),
	"spacer":    TextFieldType(4),
	"blankline": TextFieldType(5),
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
// Spacer                       4
//  - Identifies the Text Field Specification as a type
//    TextFieldSpecSpacer.
//
//    The Spacer Text Field Specification is used to create a Text
//    Field consisting of one or more white space characters (" ").
//
// BlankLine                    5
//  - Identifies a type TextLineSpecBlankLines which is used to
//    generate Blank Lines of text.
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
//  roundingType :=
// 			TextFieldType(0).Label()
//
//  isValid := roundingType.XIsValid() // isValid == true
//
//  roundingType = TextFieldType(0).None()
//
//  isValid = roundingType.XIsValid() // isValid == false
//
func (txtFieldType TextFieldType) XIsValid() bool {

	lockTextFieldType.Lock()

	defer lockTextFieldType.Unlock()

	if txtFieldType < 1 ||
		txtFieldType > 5 {

		return false
	}

	return true
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
//  Example: TextFieldType(0).FloatingPoint()
//
// Usage:
//  TxtFieldType.None(),
//  TxtFieldType.Label(),
//  TxtFieldType.DateTime(),
//  TxtFieldType.Filler(),
//  TxtFieldType.Spacer(),
//
const TxtFieldType = TextFieldType(0)
