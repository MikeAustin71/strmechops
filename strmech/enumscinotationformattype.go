package strmech

import (
	"fmt"
	"strings"
	"sync"
)

// Lock lockSciNotationFormatType before accessing these
// 'maps'.

var mapSciNotationFormatTypeCodeToString = map[SciNotationFormatType]string{
	SciNotationFormatType(0): "None",
	SciNotationFormatType(1): "Exponent",
	SciNotationFormatType(2): "ENotation",
}

var mapSciNotationFormatTypeStringToCode = map[string]SciNotationFormatType{
	"None":      SciNotationFormatType(0),
	"Exponent":  SciNotationFormatType(1),
	"ENotation": SciNotationFormatType(2),
}

var mapSciNotationFormatTypeLwrCaseStringToCode = map[string]SciNotationFormatType{
	"none":      SciNotationFormatType(0),
	"exponent":  SciNotationFormatType(1),
	"eNotation": SciNotationFormatType(2),
}

//	SciNotationFormatType
//
//	An enumeration of format types used to display
//	Scientific Notation values in Number Strings.
//
// ----------------------------------------------------------------
//
//	# Definition of Terms
//
// ----------------------------------------------------------------
//
//	Scientific notation is a way of expressing numbers
//	that are too large or too small (usually would
//	result in a long string of digits) to be
//	conveniently written in decimal form.
//
//	It may be referred to as scientific form or standard
//	index form, or standard form in the United Kingdom.
//
//	This base ten notation is commonly used by scientists,
//	mathematicians, and engineers, in part because it can
//	simplify certain arithmetic operations. On scientific
//	calculators it is usually known as "SCI" display mode.
//
//	In scientific notation, nonzero numbers are written in
//	the form  m × 10^n or m times ten raised to the power
//	of n, where n is an integer, and the coefficient m is
//	a nonzero real number (usually between 1 and 10 in
//	absolute value, and nearly always written as a
//	terminating decimal). The integer n is called the
//	exponent and the real number m is called the
//	significand or mantissa.
//
//	The term "mantissa" can be ambiguous where logarithms
//	are involved, because it is also the traditional name
//	of the fractional part of the common logarithm. If the
//	number is negative then a minus sign precedes m, as in
//	ordinary decimal notation. In normalized notation, the
//	exponent is chosen so that the absolute value (modulus)
//	of the significand m is at least 1 but less than 10.
//
//		https://en.wikipedia.org/wiki/Scientific_notation
//
// ----------------------------------------------------------------
//
//	# Enumerations In Go
//
// ----------------------------------------------------------------
//
//	Since the Go Programming Language does not directly
//	support enumerations, the SciNotationFormatType has
//	been adapted to function in a manner similar to classic
//	enumerations.
//
//	SciNotationFormatType is declared as a type 'int'. The
//	method names for this type effectively represent an
//	enumeration of numeric sign value types.
//
//	These methods are listed as follows:
//
// ----------------------------------------------------------------
//
//	Method			Integer
//
//	 Name			 Value
//
//	------			-------
//
//	None			(0)
//
//		Signals that 'SciNotationFormatType' has not been
//		initialized and therefore has no value. This is
//		an error condition.
//
//	Exponent		(1)
//		The Exponent Format signals that the Scientific
//		Notation value will be expressed in the form,
//		m × 10^n.
//
//			Example: 2.53 x 10^8
//
//	ENotation		(2)
//
//		The ENotation Format mimics the format used on
//		numerical calculators. Under this format, the
//		Scientific Notation value will be expressed in
//		the form 'mEn'.
//
//			Example: 2.53E8
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://en.wikipedia.org/wiki/Scientific_notation
//	https://www.medcalc.org/manual/scientific-notation.php
//	https://researchtweet.com/scientific-notation-definition-calculation/
//	https://www.wikihow.com/Multiply-Scientific-Notation
//
//	For easy access to these enumeration values, use the
//	global constant SciNotFmtType.
//
//		Example: SciNotFmtType.Exponent()
//
//	Otherwise you will need to use the formal syntax.
//
//		Example: SciNotationFormatType(0).Exponent()
//
//	Depending on your editor, intellisense (a.k.a.
//	intelligent code completion) may not list the
//	SciNotationFormatType methods in alphabetical order.
//
//	Be advised that all SciNotationFormatType methods
//	beginning with 'X', as well as the method "String()",
//	are utility methods, and NOT part of the enumeration
//	values.
type SciNotationFormatType int

var lockSciNotationFormatType sync.Mutex

// None - Signals that 'SciNotationFormatType' has not been
// initialized and therefore has no value. This is an error
// condition.
func (sciNotFmtType SciNotationFormatType) None() SciNotationFormatType {

	lockSciNotationFormatType.Lock()

	defer lockSciNotationFormatType.Unlock()

	return SciNotationFormatType(0)
}

// Exponent
//
//	The Exponent Format signals that the Scientific
//	Notation value will be expressed in the form,
//	m × 10^n.
//
//			Example: 2.53 x 10^8
func (sciNotFmtType SciNotationFormatType) Exponent() SciNotationFormatType {

	lockSciNotationFormatType.Lock()

	defer lockSciNotationFormatType.Unlock()

	return SciNotationFormatType(1)
}

// ENotation
//
// The ENotation Format mimics the format used on
// numerical calculators. Under this format, the
// Scientific Notation value will be expressed in the
// form 'mEn'.
//
//	Example: 2.53E8
func (sciNotFmtType SciNotationFormatType) ENotation() SciNotationFormatType {

	lockSciNotationFormatType.Lock()

	defer lockSciNotationFormatType.Unlock()

	return SciNotationFormatType(2)
}

//	String
//
//	Returns a string with the name of the enumeration
//	associated with this instance of SciNotationFormatType.
//
//	This is a standard utility method and is not part of
//	the valid enumerations for this type.
//
// ----------------------------------------------------------------
//
// Usage
//
//	t:= SciNotationFormatType(0).ENotation()
//	str := t.String()
//	   str is now equal to 'ENotation'
func (sciNotFmtType SciNotationFormatType) String() string {

	lockSciNotationFormatType.Lock()

	defer lockSciNotationFormatType.Unlock()

	result, ok :=
		mapSciNotationFormatTypeCodeToString[sciNotFmtType]

	if !ok {
		return "Error: SciNotationFormatType code UNKNOWN!"
	}

	return result
}

//	XIsValid
//	Returns a boolean value signaling whether the current
//	SciNotationFormatType value is valid.
//
//	Be advised, the enumeration value "None" is considered
//	NOT VALID.
//
//	This is a standard utility method and is not part of
//	the valid enumerations for this type.
//
// ------------------------------------------------------------------------
//
//	Usage
//
//	sciNotFmt := SciNotationFormatType(0).Exponent()
//
//	isValid := sciNotFmt.XIsValid() // isValid == true
//
//	sciNotFmt = SciNotationFormatType(0).None()
//
//	isValid = sciNotFmt.XIsValid() // isValid == false
func (sciNotFmtType SciNotationFormatType) XIsValid() bool {

	lockSciNotationFormatType.Lock()

	defer lockSciNotationFormatType.Unlock()

	return new(sciNotationFormatTypeNanobot).
		isValidSciNotFmtType(
			sciNotFmtType)
}

//	XParseString
//
//	Receives a string and attempts to match it with the
//	string value of a supported enumeration. If successful,
//	a new instance of SciNotationFormatType is returned set
//	to the value of the associated enumeration.
//
//	This is a standard utility method and is not part of the
//	valid enumerations for this type.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	valueString			string
//
//		A string which will be matched against the
//		enumeration string values. If 'valueString' is
//		equal to one of the enumeration names, this
//		method will proceed to successful completion and
//		return the correct enumeration value.
//
//	caseSensitive		bool
//
//		If 'true', the search for enumeration names will
//		be case-sensitive and will require an exact match.
//		Therefore, 'exponent' will NOT match the
//		enumeration name, 'Exponent'.
//
//		If 'false', a case-insensitive search is conducted
//		for the enumeration name. In this case, 'exponent'
//		will match the enumeration name 'Exponent'.
//
// ------------------------------------------------------------------------
//
// # Return Values
//
//	SciNotationFormatType
//
//		Upon successful completion, this method will
//		return a new instance of SciNotationFormatType
//		set to the value of the enumeration matched by
//		the string search performed on input parameter,
//		'valueString'.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'. If an
//		error condition is encountered, this method will
//		return an error type which encapsulates an
//		appropriate error message.
//
// ------------------------------------------------------------------------
//
// # Usage
//
//	t, err := SciNotationFormatType(0).XParseString("Exponent", true)
//
//	t is now equal to SciNotationFormatType(0).Exponent()
func (sciNotFmtType SciNotationFormatType) XParseString(
	valueString string,
	caseSensitive bool) (SciNotationFormatType, error) {

	lockSciNotationFormatType.Lock()

	defer lockSciNotationFormatType.Unlock()

	ePrefix := "SciNotationFormatType.XParseString() "

	if len(valueString) < 4 {
		return SciNotationFormatType(0),
			fmt.Errorf(ePrefix+"\n"+
				"Input parameter 'valueString' is INVALID!\n"+
				"String length is less than '4'.\n"+
				"valueString='%v'\n", valueString)
	}

	var ok bool
	var strNumSignValType SciNotationFormatType

	if caseSensitive {

		strNumSignValType, ok =
			mapSciNotationFormatTypeStringToCode[valueString]

		if !ok {
			return SciNotationFormatType(-2),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid SciNotationFormatType Value.\n"+
					"valueString='%v'\n", valueString)
		}

	} else {

		strNumSignValType, ok =
			mapSciNotationFormatTypeLwrCaseStringToCode[strings.ToLower(valueString)]

		if !ok {
			return SciNotationFormatType(-2),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid SciNotationFormatType Value.\n"+
					"valueString='%v'\n", valueString)
		}
	}

	return strNumSignValType, nil
}

//	XReturnNoneIfInvalid
//
//	Provides a standardized value for invalid instances
//	of enumeration SciNotationFormatType.
//
//	If the current instance of SciNotationFormatType is
//	invalid, this method will always return a value of
//	SciNotationFormatType(0).None().
//
// # Background
//
//	Enumeration SciNotationFormatType has an underlying
//	type of integer (int). This means the type could
//	conceivably be set to any integer value. This method
//	ensures that all invalid SciNotationFormatType instances
//	are consistently classified as 'None'
//	(SciNotationFormatType(0).None()). Remember that 'None'
//	is considered an invalid value.
//
//	For example, assume that SciNotationFormatType was set
//	to an integer value of -848972. Calling this method on
//	a SciNotationFormatType with this invalid integer value
//	will return an integer value of zero or the equivalent
//	of SciNotationFormatType(0).None(). This conversion is
//	useful in generating text strings for meaningful
//	informational and error messages.
//
//	This is a standard utility method and is not part of the
//	valid enumerations for this type.
func (sciNotFmtType SciNotationFormatType) XReturnNoneIfInvalid() SciNotationFormatType {

	lockSciNotationFormatType.Lock()

	defer lockSciNotationFormatType.Unlock()

	isValid := new(sciNotationFormatTypeNanobot).
		isValidSciNotFmtType(
			sciNotFmtType)

	if !isValid {
		return SciNotationFormatType(0)
	}

	return sciNotFmtType
}

// XValue
//
// This method returns the enumeration value of the
// current SciNotationFormatType instance.
//
// This is a standard utility method and is not part
// of the valid enumerations for this type.
func (sciNotFmtType SciNotationFormatType) XValue() SciNotationFormatType {

	lockSciNotationFormatType.Lock()

	defer lockSciNotationFormatType.Unlock()

	return sciNotFmtType
}

// XValueInt
//
// This method returns the integer value of the current
// SciNotationFormatType instance.
//
// This is a standard utility method and is not part of
// the valid enumerations for this type.
func (sciNotFmtType SciNotationFormatType) XValueInt() int {

	lockSciNotationFormatType.Lock()

	defer lockSciNotationFormatType.Unlock()

	return int(sciNotFmtType)
}

//	SciNotFmtType
//
//	Public global constant of type NumericSignValueType.
//
//	This variable serves as an easier, shorthand
//	technique for accessing SciNotationFormatType values.
//
//		Example: SciNotFmtType.Exponent()
//
//	Otherwise you will need to use the formal syntax.
//
//		Example: SciNotationFormatType(0).Exponent()
//
// Usage:
// SciNotFmtType.None(),
// SciNotFmtType.Exponent(),
// SciNotFmtType.ENotation(),
const SciNotFmtType = SciNotationFormatType(0)

type sciNotationFormatTypeNanobot struct {
	lock *sync.Mutex
}

// isValidSciNotFmtType
//
// Receives an instance of SciNotationFormatType and
// returns a boolean value signaling whether that
// SciNotationFormatType instance is valid.
//
// If the passed instance of SciNotationFormatType is
// valid, this method returns 'true'.
//
// Be advised, the enumeration value "None" is
// considered NOT VALID. "None" represents an error
// condition.
//
// This is a standard utility method and is not part of
// the valid SciNotationFormatType enumeration.
func (sciNotFmtTypeNanobot *sciNotationFormatTypeNanobot) isValidSciNotFmtType(
	sciNotFmtType SciNotationFormatType) bool {

	if sciNotFmtTypeNanobot.lock == nil {
		sciNotFmtTypeNanobot.lock = new(sync.Mutex)
	}

	sciNotFmtTypeNanobot.lock.Lock()

	defer sciNotFmtTypeNanobot.lock.Unlock()

	if sciNotFmtType < 1 ||
		sciNotFmtType > 2 {

		return false
	}

	return true
}
