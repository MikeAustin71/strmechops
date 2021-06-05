package strmech

import (
	"fmt"
	"strings"
	"sync"
)

var mDataFieldTrailingDelimiterStringToCode = map[string]DataFieldTrailingDelimiterType{
	"Unknown":     DataFieldTrailingDelimiterType(0).Unknown(),
	"EndOfField":  DataFieldTrailingDelimiterType(0).EndOfField(),
	"Comment":     DataFieldTrailingDelimiterType(0).Comment(),
	"EndOfLine":   DataFieldTrailingDelimiterType(0).EndOfLine(),
	"EndOfString": DataFieldTrailingDelimiterType(0).EndOfString(),
}

var mDataFieldTrailingDelimiterLwrCaseStringToCode = map[string]DataFieldTrailingDelimiterType{
	"unknown":     DataFieldTrailingDelimiterType(0).Unknown(),
	"endoffield":  DataFieldTrailingDelimiterType(0).EndOfField(),
	"comment":     DataFieldTrailingDelimiterType(0).Comment(),
	"endofline":   DataFieldTrailingDelimiterType(0).EndOfLine(),
	"endofstring": DataFieldTrailingDelimiterType(0).EndOfString(),
}

var mDataFieldTrailingDelimiterToString = map[DataFieldTrailingDelimiterType]string{
	DataFieldTrailingDelimiterType(0).Unknown():     "Unknown",
	DataFieldTrailingDelimiterType(0).EndOfField():  "EndOfField",
	DataFieldTrailingDelimiterType(0).Comment():     "Comment",
	DataFieldTrailingDelimiterType(0).EndOfLine():   "EndOfLine",
	DataFieldTrailingDelimiterType(0).EndOfString(): "EndOfString",
}

// DataFieldTrailingDelimiterType - Enumerates the type of delimiters used to mark the
// end of a a data field within a host string.
//
// DataFieldTrailingDelimiterType has been adapted to function as an enumeration of valid
// Data Field Delimiter Types. Since Go does not directly support enumerations, the
// 'DataFieldTrailingDelimiterType' has been configured to function in a manner similar
// to classic enumerations found in other languages like C#. For additional
// information, reference:
//
//      Jeffrey Richter Using Reflection to implement enumerated types
//             https://www.youtube.com/watch?v=DyXJy_0v0_U
//
// Valid Enumerations for DataFieldTrailingDelimiterType are invoked by calling the
// appropriate method on this type:
//
//      DataFieldTrailingDelimiterType(0).Unknown()
//      DataFieldTrailingDelimiterType(0).EndOfField()
//      DataFieldTrailingDelimiterType(0).Comment()
//      DataFieldTrailingDelimiterType(0).EndOfLine()
//      DataFieldTrailingDelimiterType(0).EndOfString()
//
// Alternatively the shorthand method of invoking this enumeration may be employed as
// follows:
//
//      DfTrailDelimiter.Unknown()
//      DfTrailDelimiter.EndOfField()
//      DfTrailDelimiter.Comment()
//      DfTrailDelimiter.EndOfLine()
//      DfTrailDelimiter.EndOfString()
//
//    Note: The variable DfTrailDelimiter is discussed below.
//
//
type DataFieldTrailingDelimiterType int

var lockDataFieldTrailingDelimiterType sync.Mutex

// Unknown - Data Field Trailing Delimiter Type is unknown or undetermined.
//
func (dfTrailDelimiter DataFieldTrailingDelimiterType) Unknown() DataFieldTrailingDelimiterType {

	lockDataFieldTrailingDelimiterType.Lock()

	defer lockDataFieldTrailingDelimiterType.Unlock()

	return DataFieldTrailingDelimiterType(0)
}

// EndOfField - The Data Field is terminated by a trailing end of field separator.
//
func (dfTrailDelimiter DataFieldTrailingDelimiterType) EndOfField() DataFieldTrailingDelimiterType {

	lockDataFieldTrailingDelimiterType.Lock()

	defer lockDataFieldTrailingDelimiterType.Unlock()

	return DataFieldTrailingDelimiterType(1)
}

// Comment - The Data Field is terminated by a trailing 'comment' separator.
//
func (dfTrailDelimiter DataFieldTrailingDelimiterType) Comment() DataFieldTrailingDelimiterType {

	lockDataFieldTrailingDelimiterType.Lock()

	defer lockDataFieldTrailingDelimiterType.Unlock()

	return DataFieldTrailingDelimiterType(2)
}

// EndOfLine - The Data Field is terminated by a designated end-of-line separator. Often this is
// a designated new line character such as '\n'.
//
func (dfTrailDelimiter DataFieldTrailingDelimiterType) EndOfLine() DataFieldTrailingDelimiterType {

	lockDataFieldTrailingDelimiterType.Lock()

	defer lockDataFieldTrailingDelimiterType.Unlock()

	return DataFieldTrailingDelimiterType(3)
}

// EndOfString - No specific character terminated the data field. The next character after the
// data field represents the end of the host string.
//
func (dfTrailDelimiter DataFieldTrailingDelimiterType) EndOfString() DataFieldTrailingDelimiterType {

	lockDataFieldTrailingDelimiterType.Lock()

	defer lockDataFieldTrailingDelimiterType.Unlock()

	return DataFieldTrailingDelimiterType(4)
}

// =============================================================================
// Utility Methods
// =============================================================================

// StatusIsValid - If the value of the current DataFieldTrailingDelimiterType instance
// is 'invalid', this method will return an error.
//
// If the DataFieldTrailingDelimiterType is instance is 'valid', this method will
// return a value of 'nil'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
func (dfTrailDelimiter DataFieldTrailingDelimiterType) StatusIsValid() error {

	lockDataFieldTrailingDelimiterType.Lock()

	defer lockDataFieldTrailingDelimiterType.Unlock()

	_, ok := mDataFieldTrailingDelimiterToString[dfTrailDelimiter]

	if !ok {
		ePrefix := "DataFieldTrailingDelimiterType.StatusIsValid()\n"
		return fmt.Errorf(ePrefix+
			"Error: The current DataFieldTrailingDelimiterType is INVALID! "+
			"DataFieldTrailingDelimiterType Value='%v'", int(dfTrailDelimiter))
	}

	return nil
}

// String - Returns a string with the name of the enumeration associated
// with this instance of 'DataFieldTrailingDelimiterType'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Return Value:
//
//  string - The string label or description for the current enumeration
//           value. If, the DataFieldTrailingDelimiterType value is invalid,
//           this method will return an empty string.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	t:= DataFieldTrailingDelimiterType(0).EndOfLine()
//	str := t.String()
//	    str is now equal to "EndOfLine"
//
func (dfTrailDelimiter DataFieldTrailingDelimiterType) String() string {

	lockDataFieldTrailingDelimiterType.Lock()

	defer lockDataFieldTrailingDelimiterType.Unlock()

	label, ok := mDataFieldTrailingDelimiterToString[dfTrailDelimiter]

	if !ok {
		return ""
	}

	return label
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
//  trailDelimType := DataFieldTrailingDelimiterType(0).EndOfField()
//
//  isValid := trailDelimType.XIsValid()
//
func (dfTrailDelimiter DataFieldTrailingDelimiterType) XIsValid() bool {

	lockDataFieldTrailingDelimiterType.Lock()

	defer lockDataFieldTrailingDelimiterType.Unlock()

	if dfTrailDelimiter > 4 ||
		dfTrailDelimiter < 0 {
		return false
	}

	return true
}

// XParseString - Receives a string and attempts to match it with
// the string value of the supported enumeration. If successful,
// a new instance of TimeZoneClass is returned set to
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
//	DataFieldTrailingDelimiterType  - Upon successful completion, this method will return a new
//	                                  instance of DataFieldTrailingDelimiterType set to the value
//	                                  of the enumeration matched by the string search performed on
//	                                  input parameter,'valueString'.
//
//	error                           - If this method completes successfully, the returned error
//	                                  Type is set equal to 'nil'. If an error condition is encountered,
//	                                  this method will return an error Type which encapsulates an
//	                                  appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage:
//
//  t, err := DataFieldTrailingDelimiterType(0).ParseString("EndOfLine", true)
//                            OR
//  t, err := DataFieldTrailingDelimiterType(0).ParseString("EndOfLine()", true)
//                            OR
//  t, err := DataFieldTrailingDelimiterType(0).ParseString("endofline", false)
//
//  For all of the cases shown above,
//  t is now equal to DataFieldTrailingDelimiterType(0).EndOfLine()
//
func (dfTrailDelimiter DataFieldTrailingDelimiterType) XParseString(
	valueString string,
	caseSensitive bool) (DataFieldTrailingDelimiterType, error) {

	lockDataFieldTrailingDelimiterType.Lock()

	defer lockDataFieldTrailingDelimiterType.Unlock()

	ePrefix := "DataFieldTrailingDelimiterType.XParseString() "

	lenValueStr := len(valueString)

	if strings.HasSuffix(valueString, "()") {
		valueString = valueString[0 : lenValueStr-2]
		lenValueStr -= 2
	}

	if lenValueStr < 3 {
		return DataFieldTrailingDelimiterType(0),
			fmt.Errorf(ePrefix+
				"Input parameter 'valueString' is INVALID! Length Less than 3-characters\n"+
				"valueString='%v'\n", valueString)
	}

	var ok bool

	var dfTrailType DataFieldTrailingDelimiterType

	if caseSensitive {

		dfTrailType, ok = mDataFieldTrailingDelimiterStringToCode[valueString]

		if !ok {
			return DataFieldTrailingDelimiterType(0),
				fmt.Errorf("%v\n"+
					"Invalid DataFieldTrailingDelimiterType Code!\n",
					ePrefix)
		}

	} else {

		valueString = strings.ToLower(valueString)

		dfTrailType, ok = mDataFieldTrailingDelimiterLwrCaseStringToCode[valueString]

		if !ok {
			return DataFieldTrailingDelimiterType(0),
				fmt.Errorf("%v\n"+
					"Invalid DataFieldTrailingDelimiterType Code!\n",
					ePrefix)
		}

	}

	return dfTrailType, nil
}

// XValueInt Value - Returns the value of the DataFieldTrailingDelimiterType instance
// as type int.
//
func (dfTrailDelimiter DataFieldTrailingDelimiterType) XValueInt() int {

	lockDataFieldTrailingDelimiterType.Lock()

	defer lockDataFieldTrailingDelimiterType.Unlock()

	return int(dfTrailDelimiter)
}

// DfTrailDelimiter - public global variable of
// type DataFieldTrailingDelimiterType.
//
// This variable serves as an easier, short hand
// technique for accessing DataFieldTrailingDelimiterType
// values.
//
// Usage:
//  DfTrailDelimiter.Unknown()
//  DfTrailDelimiter.EndOfField()
//  DfTrailDelimiter.Comment()
//  DfTrailDelimiter.EndOfLine()
//  DfTrailDelimiter.EndOfString()
//
const DfTrailDelimiter = DataFieldTrailingDelimiterType(0)
