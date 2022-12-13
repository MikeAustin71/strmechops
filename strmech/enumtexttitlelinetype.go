package strmech

import (
	"fmt"
	"strings"
	"sync"
)

// Lock lockTextTitleLineType before accessing these
// 'maps'.

var mapTextTitleLineTypeCodeToString = map[TextTileLineType]string{
	TextTileLineType(0): "None",
	TextTileLineType(1): "LeadingMarqueeLine",
	TextTileLineType(2): "TitleLine",
	TextTileLineType(3): "TrailingMarqueeLine",
}

var mapTextTileLineTypeStringToCode = map[string]TextTileLineType{
	"None":                TextTileLineType(0),
	"LeadingMarqueeLine":  TextTileLineType(1),
	"TitleLine":           TextTileLineType(2),
	"TrailingMarqueeLine": TextTileLineType(3),
}

var mapTextTileLineTypeLwrCaseStringToCode = map[string]TextTileLineType{
	"none":                TextTileLineType(0),
	"leadingmarqueeline":  TextTileLineType(1),
	"titleline":           TextTileLineType(2),
	"trailingmarqueeline": TextTileLineType(3),
}

// TextTileLineType
//
// An enumeration of Text Title Line Types. Title Lines
// are typically used in generating Title Marquees.
//
// Tittle Marquee display are divided into three distinct
// components. The first are leading marquee text lines
// which usually consists of blank lines or solid lines.
//
// After leading marquee text lines, the actual text
// title lines are displayed. These strings contain the
// actual title text.
//
// Finally, the third component of trailing marquee lines
// consists of trailing blank lines and/or solid lines.
//
// The Text Title Line Type enumeration is used to
// classify Text Title Marquee lines of text.
//
// Since the Go Programming Language does not directly
// support enumerations, the TextTileLineType has been
// adapted to function in a manner similar to classic
// enumerations.
//
// TextTileLineType is declared as a type 'int'. The method
// names effectively represent an enumeration of numeric sign
// value types. These methods are listed as follows:
//
// ----------------------------------------------------------------
//
// Method         		Integer
//
//	Name          		 Value
//
// ------        		-------
//
//	None           		  (0)	Signals that 'TextTileLineType' has
//	                     		not been initialized and therefore
//	                     		has no value. This is an error
//	                     		condition.
//
//	LeadingMarqueeLine	  (1)	Signals that the text line is
//								classified as a leading title
//								marquee text line. This type of
//								line usually consists of blank
//								lines and solid lines.
//
//	TitleLine			  (2)	Signals that the text line is
//								classified as Title Line typically
//								displayed in the center of the Title
//								Marquee.
//
//	TrailingMarqueeLine	  (3)	Classifies the line of text as a
//								trailing title marquee text line.
//								This type of line usually consists
//								of blank lines and solid lines.
//
// For easy access to these enumeration values, use the
// global constant TitleLineType.
//
//	Example: TitleLineType.TitleLine()
//
// Otherwise you will need to use the formal syntax.
// Example: TextTileLineType(0).TitleLine()
//
// Depending on your editor, intellisense (a.k.a.
// intelligent code completion) may not list the
// TitleLineType methods in alphabetical order.
//
// Be advised that all TitleLineType methods beginning
// with 'X', as well as the method 'String()', are
// utility methods, and NOT part of the enumeration
// values.
type TextTileLineType int

var lockTextTitleLineType sync.Mutex

// None
//
// Signals that 'TextTileLineType' has not been
// initialized and therefore has no value. This is an
// error condition.
func (txtTitleLineType TextTileLineType) None() TextTileLineType {

	lockTextTitleLineType.Lock()

	defer lockTextTitleLineType.Unlock()

	return TextTileLineType(0)
}

// LeadingMarqueeLine
//
// Classifies the line of text as a leading title marquee
// text line.
func (txtTitleLineType TextTileLineType) LeadingMarqueeLine() TextTileLineType {

	lockTextTitleLineType.Lock()

	defer lockTextTitleLineType.Unlock()

	return TextTileLineType(1)

}

// TitleLine
//
// Classifies the line of text as a main text title line.
// Text titles lines are typically displayed in the center
// of the Title Marquee.
func (txtTitleLineType TextTileLineType) TitleLine() TextTileLineType {

	lockTextTitleLineType.Lock()

	defer lockTextTitleLineType.Unlock()

	return TextTileLineType(2)

}

// TrailingMarqueeLine
//
// Classifies the line of text as a trailing title marquee
// text line.
func (txtTitleLineType TextTileLineType) TrailingMarqueeLine() TextTileLineType {

	lockTextTitleLineType.Lock()

	defer lockTextTitleLineType.Unlock()

	return TextTileLineType(3)

}

// String
//
// Returns a string with the name of the enumeration
// associated with this instance of TextTileLineType.
//
// This is a standard utility method and is not part
// of the valid enumerations for this type.
//
// ----------------------------------------------------------------
//
// Usage
//
//	t:= TextTileLineType(0).TrailingMarqueeLine()
//	str := t.String()
//	   str is now equal to 'TrailingMarqueeLine'
func (txtTitleLineType TextTileLineType) String() string {

	lockTextTitleLineType.Lock()

	defer lockTextTitleLineType.Unlock()

	result, ok :=
		mapTextTitleLineTypeCodeToString[txtTitleLineType]

	if !ok {
		return "Error: TextTileLineType code UNKNOWN!"
	}

	return result
}

// XIsValid
//
// Returns a boolean value signaling whether the current
// TextTileLineType value is valid.
//
// Be advised, the enumeration value "None" is considered
// NOT VALID.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	titleLineVal := TextTileLineType(0).TrailingMarqueeLine()
//
//	isValid := titleLineVal.XIsValid() // isValid == true
//
//	titleLineVal = TextTileLineType(0).None()
//
//	isValid = titleLineVal.XIsValid() // isValid == false
func (txtTitleLineType TextTileLineType) XIsValid() bool {

	lockTextTitleLineType.Lock()

	defer lockTextTitleLineType.Unlock()

	return new(textTileLineTypeNanobot).
		isValidTextTitleLineType(
			txtTitleLineType)
}

// XParseString
//
// Receives a string and attempts to match it with the
// string value of a supported enumeration. If
// successful, a new instance of TextTileLineType is
// returned set to the value of the associated
// enumeration.
//
// This is a standard utility method and is not part of
// the valid enumerations for this type.
//
// ------------------------------------------------------------------------
//
// # Input Parameters
//
//	valueString				string
//
//		A string which will be matched against the
//		enumeration string values. If 'valueString'
//		is equal to one of the enumeration names, this
//		method will proceed to successful completion
//		and return the correct enumeration value.
//
//	caseSensitive			bool
//
//		If 'true' the search for enumeration names will
//		be case-sensitive and will require an exact
//		match. Therefore, 'titleline' will NOT match the
//		enumeration name, 'TitleLine'.
//
//		If 'false' a case-insensitive search is conducted
//		for the enumeration name. In this case,
//		'titleline' will match the enumeration name
//		'TitleLine'.
//
// ------------------------------------------------------------------------
//
// # Return Values
//
//	TextTileLineType
//
//		Upon successful completion, this method will
//		return a new instance of TextTileLineType set to
//		the value of the enumeration matched by the
//		string search performed on input parameter,
//		'valueString'.
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
//
// ------------------------------------------------------------------------
//
// # Usage
//
// t, err :=
//
//		TextTileLineType(0).XParseString("TitleLine", true)
//
//	t is now equal to TextTileLineType(0).TitleLine()
func (txtTitleLineType TextTileLineType) XParseString(
	valueString string,
	caseSensitive bool) (
	TextTileLineType,
	error) {

	lockTextTitleLineType.Lock()

	defer lockTextTitleLineType.Unlock()

	ePrefix := "TextTileLineType.XParseString() "

	if len(valueString) < 4 {
		return TextTileLineType(0),
			fmt.Errorf(ePrefix+"\n"+
				"Input parameter 'valueString' is INVALID!\n"+
				"String length is less than '4'.\n"+
				"valueString='%v'\n", valueString)
	}

	var ok bool
	var strTextTitleLineType TextTileLineType

	if caseSensitive {

		strTextTitleLineType, ok =
			mapTextTileLineTypeStringToCode[valueString]

		if !ok {
			return TextTileLineType(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid TextTileLineType Value.\n"+
					"valueString='%v'\n", valueString)
		}

	} else {

		strTextTitleLineType, ok =
			mapTextTileLineTypeLwrCaseStringToCode[strings.ToLower(valueString)]

		if !ok {
			return TextTileLineType(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid NumericSignValueType Value.\n"+
					"valueString='%v'\n", valueString)
		}
	}

	return strTextTitleLineType, nil
}

// XReturnNoneIfInvalid
//
// Provides a standardized value for invalid instances of
// enumeration TextTileLineType.
//
// If the current instance of TextTileLineType is invalid,
// this method will always return a value of
// TextTileLineType(0).None().
//
// # Background
//
// Enumeration TextTileLineType has an underlying type of
// integer (int). This means the type could conceivably
// be set to any integer value. This method ensures that
// all invalid TextTileLineType instances are
// consistently classified as 'None'
// (TextTileLineType(0).None()). Remember that 'None' is
// considered an invalid value.
//
// For example, assume that TextTileLineType was set to an
// integer value of -848972. Calling this method on a
// TextTileLineType with this invalid integer value will
// return an integer value of zero or the equivalent of
// TextTileLineType(0).None(). This conversion is useful
// in generating text strings for meaningful
// informational and error messages.
//
// This is a standard utility method and is not part of
// the valid enumerations for this type.
func (txtTitleLineType TextTileLineType) XReturnNoneIfInvalid() TextTileLineType {

	lockTextTitleLineType.Lock()

	defer lockTextTitleLineType.Unlock()

	isValid := new(textTileLineTypeNanobot).
		isValidTextTitleLineType(
			txtTitleLineType)

	if !isValid {
		return TextTileLineType(0)
	}

	return txtTitleLineType
}

// XValue
//
// This method returns the enumeration value of the
// current TextTileLineType instance.
//
// This is a standard utility method and is not part
// of the valid enumerations for this type.
func (txtTitleLineType TextTileLineType) XValue() TextTileLineType {

	lockTextTitleLineType.Lock()

	defer lockTextTitleLineType.Unlock()

	return txtTitleLineType
}

// XValueInt
//
// This method returns the integer value of the current
// TextTileLineType instance.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
func (txtTitleLineType TextTileLineType) XValueInt() int {

	lockTextTitleLineType.Lock()

	defer lockTextTitleLineType.Unlock()

	return int(txtTitleLineType)
}

// TitleLineType
//
// A public global constant of type TextTileLineType.
//
// This variable serves as an easier, shorthand technique
// for accessing TextTileLineType values.
//
// Usage:
// TitleLineType.None(),
// TitleLineType.LeadingMarqueeLine(),
// TitleLineType.TitleLine(),
// TitleLineType.TrailingMarqueeLine(),
const TitleLineType = TextTileLineType(0)

// textTileLineTypeNanobot
//
// Provides helper methods for enumeration
// TextTileLineType.
type textTileLineTypeNanobot struct {
	lock *sync.Mutex
}

// isValidTextTitleLineType
//
// Receives an instance of TextTileLineType and returns a
// boolean value signaling whether that TextTileLineType
// instance is valid.
//
// If the passed instance of TextTileLineType is valid,
// this method returns 'true'.
//
// Be advised, the enumeration value "None" is considered
// NOT VALID. "None" represents an error condition.
//
// This is a standard utility method and is not part of
// the valid TextTileLineType enumeration.
func (txtTileLineTypeNanobot *textTileLineTypeNanobot) isValidTextTitleLineType(
	txtTileLineType TextTileLineType) bool {

	if txtTileLineTypeNanobot.lock == nil {
		txtTileLineTypeNanobot.lock = new(sync.Mutex)
	}

	txtTileLineTypeNanobot.lock.Lock()

	defer txtTileLineTypeNanobot.lock.Unlock()

	if txtTileLineType < 1 ||
		txtTileLineType > 3 {

		return false
	}

	return true
}
