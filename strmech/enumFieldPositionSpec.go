package strmech

import (
	"errors"
	"fmt"
	"strings"
)

var mFieldPositionSpecStringToCode = map[string]FieldPositionSpec{
	"None":         FieldPositionSpec(0).None(),
	"LeftJustify":  FieldPositionSpec(0).LeftJustify(),
	"RightJustify": FieldPositionSpec(0).RightJustify(),
	"Center":       FieldPositionSpec(0).Center(),
}

var mFieldPositionSpecLwrCaseStringToCode = map[string]FieldPositionSpec{
	"none":         FieldPositionSpec(0).None(),
	"leftjustify":  FieldPositionSpec(0).LeftJustify(),
	"rightjustify": FieldPositionSpec(0).RightJustify(),
	"center":       FieldPositionSpec(0).Center(),
}

var mFieldPositionSpecToString = map[FieldPositionSpec]string{
	FieldPositionSpec(0).None():         "None",
	FieldPositionSpec(0).LeftJustify():  "LeftJustify",
	FieldPositionSpec(0).RightJustify(): "RightJustify",
	FieldPositionSpec(0).Center():       "Center",
}

type FieldPositionSpec int

// None - Signifies that the FieldPositionSpec value has not been
// configured.
//
// This method is part of the standard FieldPositionSpec enumeration.
//
func (fPos FieldPositionSpec) None() FieldPositionSpec { return FieldPositionSpec(0) }

// LeftJustify - Value will begin at the left most index of the field
// length.
func (fPos FieldPositionSpec) LeftJustify() FieldPositionSpec { return FieldPositionSpec(1) }

// RightJustify - Value will terminate at the right most index of the field
// length.
func (fPos FieldPositionSpec) RightJustify() FieldPositionSpec { return FieldPositionSpec(2) }

// Center - Value will be centered in the field length. Margins on the left and
// right will be equal.
func (fPos FieldPositionSpec) Center() FieldPositionSpec { return FieldPositionSpec(3) }

// =============================================================================
// Utility Methods
// =============================================================================

// String - Returns a string with the name of the enumeration associated
// with this instance of 'FieldPositionSpec'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Return Value:
//
//  string - The string label or description for the current enumeration
//           value. If, the FieldPositionSpec value is invalid, this
//           method will return an empty string.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	t:= FieldPositionSpec(0).LeftJustify()
//	str := t.String()
//	    str is now equal to "LeftJustify"
//
func (fPos FieldPositionSpec) String() string {

	label, ok := mFieldPositionSpecToString[fPos]

	if !ok {
		return ""
	}

	return label
}

// UtilityIsValid - If the value of the current FieldPositionSpec instance
// is 'invalid', this method will return an error.
//
// If the FieldPositionSpec is instance is 'valid', this method will
// return a value of 'nil'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
func (fPos FieldPositionSpec) UtilityIsValid() error {

	_, ok := mFieldPositionSpecToString[fPos]

	if !ok {
		ePrefix := "FieldPositionSpec.UtilityIsValid()\n"
		return fmt.Errorf(ePrefix+
			"Error: The current FieldPositionSpec is INVALID! "+
			"FieldPositionSpec Value='%v'", int(fPos))
	}

	return nil
}

// UtilityParseString - Receives a string and attempts to match it with
// the string value of the supported enumeration. If successful,
// a new instance of FieldPositionSpec is returned set to
// the value of the associated enumeration.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ------------------------------------------------------------------------
//
// Input Parameters:
//
//	valueString   string - A string which will be matched against the
//	                       enumeration string values. If 'valueString'
//	                       is equal to one of the enumeration names, this
//	                       method will proceed to successful completion
//
//	caseSensitive   bool - If 'true' the search for enumeration names
//	                       will be case sensitive and will require an
//	                       exact match. Therefore, 'valid' will NOT
//	                       match the enumeration name, 'Valid'.
//
//	                       If 'false' a case insensitive search is
//	                       conducted for the enumeration name. In
//	                       this case, 'valid' will match the
//	                       enumeration name 'Valid'.
//
// ------------------------------------------------------------------------
//
// Return Values:
//
//	FieldPositionSpec        - Upon successful completion, this method will return a new
//	                          instance of FieldPositionSpec set to the value of the
//	                          enumeration matched by the string search performed on
//	                          input parameter,'valueString'.
//
//	error                   - If this method completes successfully, the returned error
//	                          Type is set equal to 'nil'. If an error condition is encountered,
//	                          this method will return an error Type which encapsulates an
//	                          appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage:
//
//  t, err := FieldPositionSpec(0).UtilityParseString("LeftJustify", true)
//                            OR
//  t, err := FieldPositionSpec(0).UtilityParseString("LeftJustify()", true)
//                            OR
//  t, err := FieldPositionSpec(0).UtilityParseString("leftjustify", false)
//
//  For all of the cases shown above,
//  t is now equal to FieldPositionSpec(0).LeftJustify()
//
func (fPos FieldPositionSpec) UtilityParseString(
	valueString string,
	caseSensitive bool) (FieldPositionSpec, error) {

	ePrefix := "FieldPositionSpec.UtilityParseString() "

	lenValueStr := len(valueString)

	if strings.HasSuffix(valueString, "()") {
		valueString = valueString[0 : lenValueStr-2]
		lenValueStr -= 2
	}

	if lenValueStr < 4 {
		return FieldPositionSpec(0).None(),
			fmt.Errorf(ePrefix+
				"Input parameter 'valueString' is INVALID! Length Less than 4-characters\n"+
				"valueString='%v'\n", valueString)
	}

	var ok bool

	var tzClassCode FieldPositionSpec

	if caseSensitive {

		tzClassCode, ok = mFieldPositionSpecStringToCode[valueString]

		if !ok {
			return FieldPositionSpec(0).None(),
				errors.New(ePrefix + "Invalid Field Position Specification!\n")
		}

	} else {

		valueString = strings.ToLower(valueString)

		tzClassCode, ok = mFieldPositionSpecLwrCaseStringToCode[valueString]

		if !ok {
			return FieldPositionSpec(0).None(),
				errors.New(ePrefix + "Invalid Field Position Specification Code!\n")
		}
	}

	return tzClassCode, nil
}

// UtilityValue - Returns the value of the FieldPositionSpec instance
// as type FieldPositionSpec.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
func (fPos FieldPositionSpec) UtilityValue() FieldPositionSpec {

	return fPos
}

// FieldPos - public global variable of type FieldPositionSpec.
//
// This variable serves as an easier, short hand technique for
//accessing FieldPositionSpec values.
//
// Usage:
//  FieldPos.None()
//  FieldPos.LeftJustify()
//  FieldPos.RightJustify()
//  FieldPos.Center()
//
var FieldPos FieldPositionSpec
