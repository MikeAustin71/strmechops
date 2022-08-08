package strmech

import (
	"fmt"
	"strings"
	"sync"
)

// Do NOT access these maps without first getting
// the lock on 'lockDataFieldTrailingDelimiterType'.

var mDataFieldTrailingDelimiterStringToCode = map[string]DataFieldTrailingDelimiterType{
	"None":        DataFieldTrailingDelimiterType(0),
	"EndOfField":  DataFieldTrailingDelimiterType(1),
	"Comment":     DataFieldTrailingDelimiterType(2),
	"EndOfLine":   DataFieldTrailingDelimiterType(3),
	"EndOfString": DataFieldTrailingDelimiterType(4),
}

var (
	mDataFieldTrailingDelimiterLwrCaseStringToCode = map[string]DataFieldTrailingDelimiterType{
		"none":        DataFieldTrailingDelimiterType(0),
		"endoffield":  DataFieldTrailingDelimiterType(1),
		"comment":     DataFieldTrailingDelimiterType(2),
		"endofline":   DataFieldTrailingDelimiterType(3),
		"endofstring": DataFieldTrailingDelimiterType(4),
	}
)

var mDataFieldTrailingDelimiterToString = map[DataFieldTrailingDelimiterType]string{
	DataFieldTrailingDelimiterType(0): "None",
	DataFieldTrailingDelimiterType(1): "EndOfField",
	DataFieldTrailingDelimiterType(2): "Comment",
	DataFieldTrailingDelimiterType(3): "EndOfLine",
	DataFieldTrailingDelimiterType(4): "EndOfString",
}

// DataFieldTrailingDelimiterType - Enumerates the type of delimiters used to mark the
// end of a data field within a host string.
//
// DataFieldTrailingDelimiterType has been adapted to function as an enumeration of valid
// Data Field Delimiter Types. Since Go does not directly support enumerations, the
// 'DataFieldTrailingDelimiterType' has been configured to function in a manner similar
// to classic enumerations found in other languages like C#. For additional
// information, reference:
//
//	Jeffrey Richter Using Reflection to implement enumerated types
//	       https://www.youtube.com/watch?v=DyXJy_0v0_U
//
// Valid Enumerations for DataFieldTrailingDelimiterType are invoked by calling the
// appropriate method on this type:
//
//	DataFieldTrailingDelimiterType(0).None()
//	DataFieldTrailingDelimiterType(0).EndOfField()
//	DataFieldTrailingDelimiterType(0).Comment()
//	DataFieldTrailingDelimiterType(0).EndOfLine()
//	DataFieldTrailingDelimiterType(0).EndOfString()
//
// Alternatively the shorthand method of invoking this enumeration may be employed as
// follows:
//
//	  DfTrailDelimiter.None()
//	  DfTrailDelimiter.EndOfField()
//	  DfTrailDelimiter.Comment()
//	  DfTrailDelimiter.EndOfLine()
//	  DfTrailDelimiter.EndOfString()
//
//	Note: The variable DfTrailDelimiter is discussed below.
type DataFieldTrailingDelimiterType int

var lockDataFieldTrailingDelimiterType sync.Mutex

// None - Signals that the Data Field Trailing Delimiter Type is
// empty or uninitialized.
//
// 'None' is an invalid or error condition.
//
// This method is part of the standard enumeration.
func (dfTrailDelimiter DataFieldTrailingDelimiterType) None() DataFieldTrailingDelimiterType {

	lockDataFieldTrailingDelimiterType.Lock()

	defer lockDataFieldTrailingDelimiterType.Unlock()

	return DataFieldTrailingDelimiterType(0)
}

// EndOfField - The Data Field is terminated by a trailing end of
// field separator.
//
// This method is part of the standard enumeration.
func (dfTrailDelimiter DataFieldTrailingDelimiterType) EndOfField() DataFieldTrailingDelimiterType {

	lockDataFieldTrailingDelimiterType.Lock()

	defer lockDataFieldTrailingDelimiterType.Unlock()

	return DataFieldTrailingDelimiterType(1)
}

// Comment - The Data Field is terminated by a trailing 'comment'
// separator.
//
// This method is part of the standard enumeration.
func (dfTrailDelimiter DataFieldTrailingDelimiterType) Comment() DataFieldTrailingDelimiterType {

	lockDataFieldTrailingDelimiterType.Lock()

	defer lockDataFieldTrailingDelimiterType.Unlock()

	return DataFieldTrailingDelimiterType(2)
}

// EndOfLine - The Data Field is terminated by a designated
// end-of-line separator. Often this is a designated new line
// character such as '\n'.
//
// This method is part of the standard enumeration.
func (dfTrailDelimiter DataFieldTrailingDelimiterType) EndOfLine() DataFieldTrailingDelimiterType {

	lockDataFieldTrailingDelimiterType.Lock()

	defer lockDataFieldTrailingDelimiterType.Unlock()

	return DataFieldTrailingDelimiterType(3)
}

// EndOfString - No specific character terminated the data field.
// The next character after the data field represents the end of
// the host string.
//
// This method is part of the standard enumeration.
func (dfTrailDelimiter DataFieldTrailingDelimiterType) EndOfString() DataFieldTrailingDelimiterType {

	lockDataFieldTrailingDelimiterType.Lock()

	defer lockDataFieldTrailingDelimiterType.Unlock()

	return DataFieldTrailingDelimiterType(4)
}

// =============================================================================
// Utility Methods
// =============================================================================

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
//	string - The string label or description for the current enumeration
//	         value. If, the DataFieldTrailingDelimiterType value is invalid,
//	         this method will return an empty string.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	t:= DataFieldTrailingDelimiterType(0).EndOfLine()
//	str := t.String()
//	    str is now equal to "EndOfLine"
func (dfTrailDelimiter DataFieldTrailingDelimiterType) String() string {

	lockDataFieldTrailingDelimiterType.Lock()

	defer lockDataFieldTrailingDelimiterType.Unlock()

	label, ok := mDataFieldTrailingDelimiterToString[dfTrailDelimiter]

	if !ok {
		return "Error: Data Field Trailing Delimiter Type Specification UNKNOWN!"
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
//	trailDelimType := DataFieldTrailingDelimiterType(0).EndOfField()
//
//	isValid := trailDelimType.XIsValid()
func (dfTrailDelimiter DataFieldTrailingDelimiterType) XIsValid() bool {

	lockDataFieldTrailingDelimiterType.Lock()

	defer lockDataFieldTrailingDelimiterType.Unlock()

	return new(dataFieldTrailDelimiterNanobot).
		isValidDfTrailDelimiter(dfTrailDelimiter)
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
//	                       will be case-sensitive and will require an
//	                       exact match. Therefore, 'valid' will NOT
//	                       match the enumeration name, 'Valid'.
//
//	                       If 'false' a case-insensitive search is
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
//	t, err := DataFieldTrailingDelimiterType(0).ParseString("EndOfLine", true)
//	                          OR
//	t, err := DataFieldTrailingDelimiterType(0).ParseString("EndOfLine()", true)
//	                          OR
//	t, err := DataFieldTrailingDelimiterType(0).ParseString("endofline", false)
//
//	For all the cases shown above, t is now equal to
//	DataFieldTrailingDelimiterType(0).EndOfLine()
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

// XReturnNoneIfInvalid - Provides a standardized value for invalid
// instances of enumeration DataFieldTrailingDelimiterType.
//
// If the current instance of DataFieldTrailingDelimiterType is
// invalid, this method will always return a value of
// DataFieldTrailingDelimiterType(0).None().
//
// # Background
//
// Enumeration DataFieldTrailingDelimiterType has an underlying
// type of integer (int). This means the type could conceivably be
// set to any integer value. This method ensures that all invalid
// DataFieldTrailingDelimiterType instances are consistently
// classified as 'None' (DataFieldTrailingDelimiterType(0).None()).
// Remember that 'None' is considered an invalid value.
//
// For example, assume that DataFieldTrailingDelimiterType was set
// to an invalid integer value of -848972. Calling this method on a
// DataFieldTrailingDelimiterType with this invalid integer value
// will return an integer value of zero or the equivalent of
// DataFieldTrailingDelimiterType(0).None(). This conversion is
// useful in generating text strings for meaningful informational
// and error messages.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
func (dfTrailDelimiter DataFieldTrailingDelimiterType) XReturnNoneIfInvalid() DataFieldTrailingDelimiterType {

	lockDataFieldTrailingDelimiterType.Lock()

	defer lockDataFieldTrailingDelimiterType.Unlock()

	isValid := new(dataFieldTrailDelimiterNanobot).
		isValidDfTrailDelimiter(dfTrailDelimiter)

	if !isValid {
		return DataFieldTrailingDelimiterType(0)
	}

	return dfTrailDelimiter
}

// XValue - This method returns the enumeration value of the
// current DataFieldTrailingDelimiterType instance.
//
// This is a standard utility method and is NOT part of the valid
// enumerations for this type.
func (dfTrailDelimiter DataFieldTrailingDelimiterType) XValue() DataFieldTrailingDelimiterType {

	lockDataFieldTrailingDelimiterType.Lock()

	defer lockDataFieldTrailingDelimiterType.Unlock()

	return dfTrailDelimiter
}

// XValueInt Value - Returns the value of the DataFieldTrailingDelimiterType instance
// as type int.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
func (dfTrailDelimiter DataFieldTrailingDelimiterType) XValueInt() int {

	lockDataFieldTrailingDelimiterType.Lock()

	defer lockDataFieldTrailingDelimiterType.Unlock()

	return int(dfTrailDelimiter)
}

// DfTrailDelimiter - public global variable of
// type DataFieldTrailingDelimiterType.
//
// This variable serves as an easier, shorthand
// technique for accessing DataFieldTrailingDelimiterType
// values.
//
// Usage:
//
//	DfTrailDelimiter.None()
//	DfTrailDelimiter.EndOfField()
//	DfTrailDelimiter.Comment()
//	DfTrailDelimiter.EndOfLine()
//	DfTrailDelimiter.EndOfString()
const DfTrailDelimiter = DataFieldTrailingDelimiterType(0)

// dataFieldTrailDelimiterNanobot - Provides helper methods for
// enumeration DataFieldTrailingDelimiterType.
type dataFieldTrailDelimiterNanobot struct {
	lock *sync.Mutex
}

// isValidDfTrailDelimiter - Receives an instance of DataFieldTrailingDelimiterType and
// returns a boolean value signaling whether that DataFieldTrailingDelimiterType
// instance is valid.
//
// If the passed instance of DataFieldTrailingDelimiterType is valid, this method
// returns 'true'.
//
// Be advised, the enumeration value "None" is considered NOT
// VALID. "None" represents an error condition.
//
// This is a standard utility method and is not part of the valid
// DataFieldTrailingDelimiterType enumeration.
func (textFieldNanobot *dataFieldTrailDelimiterNanobot) isValidDfTrailDelimiter(
	DataFieldTrailDelimiter DataFieldTrailingDelimiterType) bool {

	if textFieldNanobot.lock == nil {
		textFieldNanobot.lock = new(sync.Mutex)
	}

	textFieldNanobot.lock.Lock()

	defer textFieldNanobot.lock.Unlock()

	if DataFieldTrailDelimiter < 1 ||
		DataFieldTrailDelimiter > 4 {

		return false
	}

	return true
}
