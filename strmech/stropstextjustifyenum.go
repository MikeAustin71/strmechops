package strmech

import (
	"fmt"
	"strings"
	"sync"
)

var mStrOpsTextJustifyCodeToString = map[TextJustify]string{
	TextJustify(0): "None",
	TextJustify(1): "Left",
	TextJustify(2): "Right",
	TextJustify(3): "Center",
}

var mStrOpsTextJustifyStringToCode = map[string]TextJustify{
	"None":     TextJustify(0),
	"Left":     TextJustify(1),
	"Right":    TextJustify(2),
	"Center":   TextJustify(3),
	"Centered": TextJustify(3),
}

var mStrOpsTextJustifyLwrCaseStringToCode = map[string]TextJustify{
	"none":     TextJustify(0),
	"left":     TextJustify(1),
	"right":    TextJustify(2),
	"center":   TextJustify(3),
	"centered": TextJustify(3),
}

// TextJustify - An enumeration of text justification designations.
// TextJustify is used to specify 'Right-Justified',
// 'Left-Justified' and 'Centered' string positioning within text
// fields.
//
// Since Go does not directly support enumerations, the 'TextJustify'
// type has been adapted to function in a manner similar to classic
// enumerations. 'TextJustify' is declared as a type 'int'. The
// method names effectively represent an enumeration of text
// justification formats. These methods are listed as follows:
//
// None            (0) - Signals that 'TextJustify' value has NOT
//                       been initialized. This is an error condition.
//
//
// Left            (1) - Signals that the text justification format is
//                       set to 'Left-Justify'. Strings within text
//                       fields will be flush with the left margin.
//
//                           Example: "TextString      "
//
//
// Right           (2) - Signals that the text justification format is
//                       set to 'Right-Justify'. Strings within text
//                       fields will terminate at the right margin.
//
//                           Example: "      TextString"
//
//
// Center          (3) - Signals that the text justification format is
//                       is set to 'Centered'. Strings will be positioned
//                       in the center of the text field equidistant
//                       from the left and right margins.
//
//                           Example: "   TextString   "
//
//
// For easy access to these enumeration values, use the global variable
// 'TxtJustify'. Example: TxtJustify.Right()
//
// Otherwise you will need to use the formal syntax.
// Example: TextJustify(0).Right()
//
// Depending on your editor, intellisense (a.k.a. intelligent code completion) may not
// list the TextJustify methods in alphabetical order. Be advised that all 'TextJustify'
// methods beginning with 'X', as well as the method 'String()', are utility methods and
// not part of the enumeration values.
//
type TextJustify int

var lockStrOpsTextJustify sync.Mutex

// None - Signals that 'SOpsTextJustify' value has NOT been initialized.
// This is an error condition.
//
// The 'None' TextJustify integer value is zero (0).
//
// This method is part of the standard enumeration.
//
func (sopsTxtJustify TextJustify) None() TextJustify {

	lockStrOpsTextJustify.Lock()

	defer lockStrOpsTextJustify.Unlock()

	return TextJustify(0)
}

// Left - Signals that the text justification format is set to
// 'Left-Justify'. Strings within text fields will be flush with
// the left margin.
//
//        Example: "TextString      "
//
// The 'Left' text justification has a TextJustify
// integer value of one (+1).
//
// This method is part of the standard enumeration.
//
func (sopsTxtJustify TextJustify) Left() TextJustify {

	lockStrOpsTextJustify.Lock()

	defer lockStrOpsTextJustify.Unlock()

	return TextJustify(1)
}

// Right - Signals that the text justification format is
// set to 'Right-Justify'. Strings within text fields will
// terminate at the right margin.
//
//        Example: "      TextString"
//
// The 'Right' text justification has a TextJustify
// integer value of two (+2).
//
// This method is part of the standard enumeration.
//
func (sopsTxtJustify TextJustify) Right() TextJustify {

	lockStrOpsTextJustify.Lock()

	defer lockStrOpsTextJustify.Unlock()

	return TextJustify(2)
}

// Center - Signals that the text justification format is
// is set to 'Center'. Strings will be positioned in the
// center of the text field equidistant from the left and
// right margins.
//
//        Example: "   TextString   "
//
// The 'Center' text justification has a TextJustify
// integer value of three (+3).
//
// This method is part of the standard enumeration.
//
func (sopsTxtJustify TextJustify) Center() TextJustify {

	lockStrOpsTextJustify.Lock()

	defer lockStrOpsTextJustify.Unlock()

	return TextJustify(3)
}

// String - Returns a string with the name of the enumeration associated
// with this instance of 'TextJustify'.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
// t:= TextJustify(0).Center()
// str := t.String()
//     str is now equal to 'Center'
//
func (sopsTxtJustify TextJustify) String() string {

	lockStrOpsTextJustify.Lock()

	defer lockStrOpsTextJustify.Unlock()

	result, ok :=
		mStrOpsTextJustifyCodeToString[sopsTxtJustify]

	if !ok {
		return "Error: TextJustify code UNKNOWN!"
	}

	return result
}

// XIsValid - Returns a boolean value signaling whether the current
// TextJustify value is valid.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
//  textJustification := TextJustify(0).Right()
//
//  isValid := textJustification.XIsValid()
//
func (sopsTxtJustify TextJustify) XIsValid() bool {

	lockStrOpsTextJustify.Lock()

	defer lockStrOpsTextJustify.Unlock()

	if sopsTxtJustify > 3 ||
		sopsTxtJustify < 1 {
		return false
	}

	return true
}

// XParseString - Receives a string and attempts to match it with
// the string value of a supported enumeration. If successful, a
// new instance of TextJustify is returned set to the value
// of the associated enumeration.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
// valueString   string - A string which will be matched against the
//                        enumeration string values. If 'valueString'
//                        is equal to one of the enumeration names, this
//                        method will proceed to successful completion
//                        and return the correct enumeration value.
//
// caseSensitive   bool - If 'true' the search for enumeration names
//                        will be case sensitive and will require an
//                        exact match. Therefore, 'gregorian' will NOT
//                        match the enumeration name, 'Gregorian'.
//
//                        If 'false' a case insensitive search is conducted
//                        for the enumeration name. In this case, 'gregorian'
//                        will match match enumeration name 'Gregorian'.
//
// ------------------------------------------------------------------------
//
// Return Values
//
// TextJustify
//     - Upon successful completion, this method will return a new
//       instance of TextJustify set to the value of the enumeration
//       matched by the string search performed on input parameter,
//       'valueString'.
//
// error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If an error condition is encountered,
//       this method will return an error type which encapsulates an
//       appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
// t, err := TextJustify(0).XParseString("Right", true)
//
//     t is now equal to TextJustify(0).Right()
//
func (sopsTxtJustify TextJustify) XParseString(
	valueString string,
	caseSensitive bool) (TextJustify, error) {

	lockStrOpsTextJustify.Lock()

	defer lockStrOpsTextJustify.Unlock()

	ePrefix := "TextJustify.XParseString() "

	if len(valueString) < 4 {
		return TextJustify(0),
			fmt.Errorf(ePrefix+
				"\nInput parameter 'valueString' is INVALID!\n"+
				"String length is less than '4'.\n"+
				"valueString='%v'\n", valueString)
	}

	var ok bool
	var strOpsTxtJustification TextJustify

	if caseSensitive {

		strOpsTxtJustification, ok = mStrOpsTextJustifyStringToCode[valueString]

		if !ok {
			return TextJustify(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid TextJustify Value.\n"+
					"valueString='%v'\n", valueString)
		}

	} else {

		strOpsTxtJustification, ok = mStrOpsTextJustifyLwrCaseStringToCode[strings.ToLower(valueString)]

		if !ok {
			return TextJustify(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid TextJustify Value.\n"+
					"valueString='%v'\n", valueString)
		}
	}

	return strOpsTxtJustification, nil
}

// XValue - This method returns the enumeration value of the current
// TextJustify instance.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
//
func (sopsTxtJustify TextJustify) XValue() TextJustify {

	lockStrOpsTextJustify.Lock()

	defer lockStrOpsTextJustify.Unlock()

	return sopsTxtJustify
}

// XValueInt - This method returns the integer value of the current
// TextJustify instance.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
//
func (sopsTxtJustify TextJustify) XValueInt() int {

	lockStrOpsTextJustify.Lock()

	defer lockStrOpsTextJustify.Unlock()

	return int(sopsTxtJustify)
}

// TxtJustify - public global variable of
// type TextJustify.
//
// This variable serves as an easier, short hand
// technique for accessing TextJustify values.
//
// Usage:
// TxtJustify.None(),
// TxtJustify.Left(),
// TxtJustify.Right(),
// TxtJustify.Center(),
//
var TxtJustify TextJustify
