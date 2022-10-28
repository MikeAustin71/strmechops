package strmech

import (
	"fmt"
	"strings"
	"sync"
)

//	Lock lockScientificNotationFormat before accessing
//	these maps!

var mSciNotationFmtCodeToString = map[ScientificNotationFormat]string{
	ScientificNotationFormat(0): "None",
	ScientificNotationFormat(1): "Exponential",
	ScientificNotationFormat(2): "ENotation",
}

var mSciNotationFmtStringToCode = map[string]ScientificNotationFormat{
	"None":        ScientificNotationFormat(0),
	"Exponential": ScientificNotationFormat(0),
	"ENotation":   ScientificNotationFormat(0),
}

var mSciNotationFmtLwrCaseStringToCode = map[string]ScientificNotationFormat{
	"none":        ScientificNotationFormat(0),
	"exponential": ScientificNotationFormat(0),
	"enotation":   ScientificNotationFormat(0),
}

//	ScientificNotationFormat
//
//
//	The 'Scientific Notation Format' is an
//	enumeration of type codes used for classifying
//	the output or display formats for Scientific
//	Notation.
//
// ----------------------------------------------------------------
//
// # Terminology
//
//	Scientific notation is a way of expressing numbers
//	that are too large or too small (usually would result
//	in a long string of digits) to be conveniently
//	written in decimal form. It may be referred to as
//	scientific form or standard index form, or standard
//	form in the United Kingdom. This base ten notation
//	is commonly used by scientists, mathematicians, and
//	engineers, in part because it can simplify certain
//	arithmetic operations. On scientific calculators
//	it is usually known as "SCI" display mode.
//
//	In scientific notation, nonzero numbers are written
//	in the form	"m × 10n" or m times ten raised to the
//	power of n, where n is an integer, and the coefficient
//	m is a nonzero real number (usually between 1 and 10
//	in absolute value, and nearly always written as a
//	terminating decimal).
//
//		Wikipedia
//
//	Reference:
//		https://en.wikipedia.org/wiki/Scientific_notation
//
//	Scientific notation values are either displayed in the
//	Exponential Format (2.652 x 10^8) or the E-Notation
//	Format (2.652e+8).
//
//	Reference:
//		https://en.wikipedia.org/wiki/Scientific_notation#E_notation
//
// ----------------------------------------------------------------
//
// # Enumeration Values
//
//	Since the Go Programming Language does not directly
//	support enumerations, the ScientificNotationFormat
//	has been adapted to function in a manner similar to
//	classic enumerations.
//
//	ScientificNotationFormat is declared as a type
//	'int'. The method names associated with this type
//	effectively represent an enumeration of Scientific
//	Notation Display Types. These methods are listed as
//	follows:
//
//	Method		Integer
//	 Name	 	 Value
//	------		-------
//
//	None	   	   0
//
//		Signals that 'ScientificNotationFormat' has
//		not been initialized and therefore has no value.
//		This is an error condition.
//
//	Exponential	   1
//
//		Signals that a Scientific Notation value will be
//		displayed using the Exponential notation format.
//
//		Exponential notation is displayed in the form
//		"m × 10n" or m times ten raised to the power of
//		n, where n is an integer, and the coefficient m
//		is a nonzero real number (usually between 1 and
//		10 in absolute value, and nearly always written
//		as a terminating decimal).
//
//		Reference:
//			https://en.wikipedia.org/wiki/Scientific_notation
//
//		Exponential Format Example:
//
//			Numeric Value: 265,200,000
//			Exponential Display Format: "2.652 x 10^8"
//
//	ENotation	   2
//
//		A type of scientific notation in which the phrase
//		“times 10 to the power of” is replaced by the
//		letter E; for example, 3.1 × 107 is written
//		3.1E+7 and 5.1 × 10-9 is written 5.1E-9.
//
//		Most calculators and many computer programs
//		present very large and very small results in
//		scientific notation, typically invoked by a
//		key labelled EXP (for exponent), EEX (for enter
//		exponent), EE, EX, E, or ×10x depending on vendor
//		and model.
//
//		Because superscripted exponents like 107 cannot
//		always be conveniently displayed, the letter E
//		(or e) is often used to represent "times ten
//		raised to the power of" (which would be written
//		as "× 10n") and is followed by the value of the
//		exponent; in other words, for any real number m
//		and integer n, the usage of "mEn" would indicate
//		a value of m × 10n. In this usage the character
//		e is not related to the mathematical constant e
//		or the exponential function ex (a confusion that
//		is unlikely if scientific notation is represented
//		by a capital E). Although the E stands for
//		exponent, the notation is usually referred to as
//		(scientific) E notation rather than (scientific)
//		exponential notation. The use of E notation
//		facilitates data entry and readability in textual
//		communication since it minimizes keystrokes,
//		avoids reduced font sizes and provides a simpler
//		and more concise display, but it is not encouraged
//		in some publications.
//
//	 	Example E-Notation Format:
//
//			Numeric Value: 265,200,000
//	 		E-Notation Display Format: "2.652e+8"
//
//		Reference:
//			https://encyclopedia2.thefreedictionary.com/E+notation
//			https://en.wikipedia.org/wiki/Scientific_notation#E_notation
//
// ----------------------------------------------------------------
//
// # Usage
//
//	For easy access to these enumeration values, use the
//	global constant SciNotFmt.
//
//		Example:
//			SciNotFmt.ENotation()
//
//	Otherwise you will need to use the formal syntax.
//
//		Example:
//			ScientificNotationFormat(0).ENotation()
//
//	Depending on your editor, intellisense (a.k.a.
//	intelligent code completion) may not list the
//	ScientificNotationFormat methods in alphabetical
//	order.
//
//	Be advised that all ScientificNotationFormat
//	methods beginning with 'X', as well as the method
//	'String()', are utility methods, and NOT part of the
//	enumeration values.
type ScientificNotationFormat int

var lockScientificNotationFormat sync.Mutex

// None
//
// Signals that 'ScientificNotationFormat' has not been
// initialized and therefore has no value.
//
// This is an error condition.
func (sciNotationFmt ScientificNotationFormat) None() ScientificNotationFormat {

	lockScientificNotationFormat.Lock()

	defer lockScientificNotationFormat.Unlock()

	return ScientificNotationFormat(0)
}

// Exponential
//
// Signals that a Scientific Notation value will be
// displayed using the Exponential notation format.
//
// Exponential notation is displayed in the form
// "m × 10n" or m times ten raised to the power of n,
// where n is an integer, and the coefficient m is a
// nonzero real number (usually between 1 and 10 in
// absolute value, and nearly always written as a
// terminating decimal).
//
// Reference:
//
//	https://en.wikipedia.org/wiki/Scientific_notation
//
// Exponential Format Example:
//
//	Numeric Value: 265,200,000
//
//	Exponential Display Format: "2.652 x 10^8"
func (sciNotationFmt ScientificNotationFormat) Exponential() ScientificNotationFormat {

	lockScientificNotationFormat.Lock()

	defer lockScientificNotationFormat.Unlock()

	return ScientificNotationFormat(1)
}

// ENotation
//
// A type of scientific notation in which the phrase
// “times 10 to the power of” is replaced by the letter
// E; for example, 3.1 × 107 is written 3.1E+7 and
// 5.1 × 10-9 is written 5.1E-9.
//
// Most calculators and many computer programs present
// very large and very small results in scientific
// notation, typically invoked by a key labelled EXP
// (for exponent), EEX (for enter exponent), EE, EX,
// E, or ×10x depending on vendor and model.
//
// Because superscripted exponents like 107 cannot
// always be conveniently displayed, the letter E (or e)
// is often used to represent "times ten raised to the
// power of" (which would be written as "× 10n") and is
// followed by the value of the exponent; in other
// words, for any real number m and integer n, the usage
// of "mEn" would indicate a value of m × 10n. In this
// usage the character e is not related to the
// mathematical constant e or the exponential function
// ex (a confusion that is unlikely if scientific
// notation is represented by a capital E). Although the
// E stands for exponent, the notation is usually
// referred to as (scientific) E notation rather than
// (scientific) exponential notation. The use of E
// notation facilitates data entry and readability in
// textual communication since it minimizes keystrokes,
// avoids reduced font sizes and provides a simpler and
// more concise display, but it is not encouraged in
// some publications.
//
// Example E-Notation Format:
//
//	Numeric Value: 265,200,000
//
//	E-Notation Display Format: "2.652e+8"
//
// Reference:
//
//	https://encyclopedia2.thefreedictionary.com/E+notation
//	https://en.wikipedia.org/wiki/Scientific_notation#E_notation
func (sciNotationFmt ScientificNotationFormat) ENotation() ScientificNotationFormat {

	lockScientificNotationFormat.Lock()

	defer lockScientificNotationFormat.Unlock()

	return ScientificNotationFormat(2)
}

//	String
//
//	Returns a string with the name of the enumeration
//	associated with this current instance of
//	'ScientificNotationFormat'.
//
//	This is a standard utility method and is not part of
//	the valid enumerations for this type.
//
// ----------------------------------------------------------------
//
// # Usage
//
// t:= ScientificNotationFormat(0).Exponential()
// str := t.String()
//
//	str is now equal to 'Exponential'
func (sciNotationFmt ScientificNotationFormat) String() string {

	lockScientificNotationFormat.Lock()

	defer lockScientificNotationFormat.Unlock()

	result, ok := mSciNotationFmtCodeToString[sciNotationFmt]

	if !ok {

		return "Error: Scientific Notation Format code is UNKNOWN!"

	}

	return result
}

//	XIsValid
//
//	Returns a boolean value signaling whether the current
//	ScientificNotationFormat value is valid.
//
//	Be advised, the enumeration value "None" is
//	considered an INVALID selection for
//	'ScientificNotationFormat'.
//
//	This is a standard utility method and is not part of
//	the valid enumerations for this type.
//
// ------------------------------------------------------------------------
//
// # Usage
//
//	 sciNotFmt :=
//				NumberRoundingType(0).HalfAwayFromZero()
//
//	 isValid := roundingType.XIsValid() // isValid == true
//
//	 roundingType = NumberRoundingType(-999)
//
//	 isValid = roundingType.XIsValid() // isValid == false
func (sciNotationFmt ScientificNotationFormat) XIsValid() bool {

	lockScientificNotationFormat.Lock()

	defer lockScientificNotationFormat.Unlock()

	return new(scientificNotationFormatNanobot).
		isValidNumSciNotFmt(
			sciNotationFmt)
}

//	XParseString
//
//	Receives a string and attempts to match it with the
//	string value of a supported enumeration. If
//	successful, a new instance of ScientificNotationFormat
//	is returned set to the value of the associated
//	enumeration.
//
//	This is a standard utility method and is NOT part of
//	the valid enumerations for this type.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//		valueString			string
//
//			A string which will be matched against the
//			enumeration string values. If 'valueString' is
//			equal to one of the enumeration names, this
//			method will proceed to successful completion and
//			return the correct enumeration value.
//
//		caseSensitive		bool
//
//			If 'true' the search for enumeration names will
//			be case-sensitive and will require an exact
//			match. Therefore, 'exponential' will NOT match
//			the enumeration name, 'Exponential'.
//
//			A case-sensitive search will match any of the
//			following strings:
//
//	    "None"
//	    "Exponential"
//	    "ENotation"
//
//			If 'false', a case-insensitive search is conducted
//			for the enumeration name. In this example,
//			'exponential' WILL MATCH the enumeration name,
//			'Exponential'.
//
//			A case-insensitive search will match any of the
//			following lower case names:
//
//	    "none"
//	    "exponential"
//	    "enotation"
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	ScientificNotationFormat
//
//		Upon successful completion, this method will
//		return a new instance of ScientificNotationFormat
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
// ----------------------------------------------------------------
//
// # Usage
//
//	t, err := ScientificNotationFormat(0).
//	             XParseString("ENotation", true)
//
//	t is now equal to ScientificNotationFormat(0).ENotation()
func (sciNotationFmt ScientificNotationFormat) XParseString(
	valueString string,
	caseSensitive bool) (ScientificNotationFormat, error) {

	lockScientificNotationFormat.Lock()

	defer lockScientificNotationFormat.Unlock()

	ePrefix := "ScientificNotationFormat.XParseString() "

	var ok bool
	var sciNotFmt ScientificNotationFormat

	if caseSensitive {

		sciNotFmt, ok =
			mSciNotationFmtStringToCode[valueString]

		if !ok {
			return ScientificNotationFormat(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid Scientific Notation Format value.\n"+
					"valueString='%v'\n", valueString)
		}

	} else {

		sciNotFmt, ok =
			mSciNotationFmtLwrCaseStringToCode[strings.ToLower(valueString)]

		if !ok {
			return ScientificNotationFormat(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid Scientific Notation Format value.\n"+
					"valueString='%v'\n", valueString)
		}
	}

	return sciNotFmt, nil
}

//	XReturnNoneIfInvalid
//
//	Provides a standardized value for invalid instances
//	of enumeration ScientificNotationFormat.
//
//	If the current instance of ScientificNotationFormat
//	is invalid, this method will always return a value
//	of ScientificNotationFormat(0).None().
//
// ----------------------------------------------------------------
//
// # Background
//
//	Enumeration ScientificNotationFormat has an
//	underlying type of integer (int). This means the type
//	could conceivably be set to any integer value. This
//	method ensures that all invalid ScientificNotationFormat
//	instances are consistently classified as 'None'
//	(ScientificNotationFormat(0).None()). Remember that
//	'None' is considered an INVALID selection for
//	'ScientificNotationFormat'.
//
//	For example, assume that NumberRoundingType was set
//	to an integer value of -848972. Calling this method on
//	a ScientificNotationFormat with this invalid integer
//	value will return an integer value of zero or the
//	equivalent of ScientificNotationFormat(0).None(). This
//	conversion is useful in generating text strings for
//	meaningful informational and error messages.
//
//	This is a standard utility method and is not part of the
//	valid enumerations for this type.
func (sciNotationFmt ScientificNotationFormat) XReturnNoneIfInvalid() ScientificNotationFormat {

	lockScientificNotationFormat.Lock()

	defer lockScientificNotationFormat.Unlock()

	isValid := new(scientificNotationFormatNanobot).
		isValidNumSciNotFmt(sciNotationFmt)

	if !isValid {
		return ScientificNotationFormat(0)
	}

	return sciNotationFmt
}

// XValue
//
// This method returns the enumeration value of the
// current ScientificNotationFormat instance.
//
// This is a standard utility method and is NOT part of
// the valid enumerations for this type.
func (sciNotationFmt ScientificNotationFormat) XValue() ScientificNotationFormat {

	lockScientificNotationFormat.Lock()

	defer lockScientificNotationFormat.Unlock()

	return sciNotationFmt
}

// XValueInt
//
// This method returns the integer value of the current
// ScientificNotationFormat instance.
//
// This is a standard utility method and is NOT part of
// the valid enumerations for this type.
func (sciNotationFmt ScientificNotationFormat) XValueInt() int {

	lockScientificNotationFormat.Lock()

	defer lockScientificNotationFormat.Unlock()

	return int(sciNotationFmt)
}

//	SciNotFmt
//
//	Public global constant of type ScientificNotationFormat.
//
//	This variable serves as an easier, shorthand technique
//	for accessing ScientificNotationFormat values.
//
//	For easy access to these enumeration values, use the
//	global	variable SciNotFmt.
//
//		Example:
//
//			SciNotFmt.Exponential()
//
//	Otherwise you will need to use the formal syntax.
//
//	Example:
//
//		ScientificNotationFormat(0).Exponential()
//
// ----------------------------------------------------------------
//
// # Usage
//
//	SciNotFmt.None()
//	SciNotFmt.Exponential()
//	SciNotFmt.ENotation()
const SciNotFmt = ScientificNotationFormat(0)

// scientificNotationFormatNanobot
//
// Provides helper methods for enumeration
// ScientificNotationFormat.
type scientificNotationFormatNanobot struct {
	lock *sync.Mutex
}

// isValidNumSciNotFmt
//
// Receives an instance of ScientificNotationFormat and
// returns a boolean value signaling whether that
// ScientificNotationFormat instance is valid.
//
// If the passed instance of ScientificNotationFormat is
// valid, this method returns 'true'.
//
// Be advised, the enumeration value "None" is
// considered an INVALID selection for
// 'ScientificNotationFormat'.
//
// This is a standard utility method and is not part of
// the valid ScientificNotationFormat enumeration.
func (sciNotFmtNanobot *scientificNotationFormatNanobot) isValidNumSciNotFmt(
	sciNotFmt ScientificNotationFormat) bool {

	if sciNotFmtNanobot.lock == nil {
		sciNotFmtNanobot.lock = new(sync.Mutex)
	}

	sciNotFmtNanobot.lock.Lock()

	defer sciNotFmtNanobot.lock.Unlock()

	if sciNotFmt < 1 ||
		sciNotFmt > 2 {

		return false
	}

	return true
}
