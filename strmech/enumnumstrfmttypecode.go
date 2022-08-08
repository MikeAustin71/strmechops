package strmech

import (
	"fmt"
	"strings"
	"sync"
)

// Do NOT access these maps without first getting
// the lock on 'lockNumStrFormatTypeCode'.

var mNumStrFmtTypeCodeToString = map[NumStrFormatTypeCode]string{
	NumStrFormatTypeCode(0): "None",
	NumStrFormatTypeCode(1): "AbsoluteValue",
	NumStrFormatTypeCode(2): "Binary",
	NumStrFormatTypeCode(3): "Currency",
	NumStrFormatTypeCode(4): "Hexadecimal",
	NumStrFormatTypeCode(5): "Octal",
	NumStrFormatTypeCode(6): "SignedNumber",
	NumStrFormatTypeCode(7): "ScientificNotation",
}

var mNumStrFmtTypeCodeStringToCode = map[string]NumStrFormatTypeCode{
	"None":                 NumStrFormatTypeCode(0),
	"AbsoluteValue":        NumStrFormatTypeCode(1),
	"Absolute Value":       NumStrFormatTypeCode(1),
	"Absolute":             NumStrFormatTypeCode(1),
	"Binary":               NumStrFormatTypeCode(2),
	"Currency":             NumStrFormatTypeCode(3),
	"Currency Value":       NumStrFormatTypeCode(3),
	"CurrencyValue":        NumStrFormatTypeCode(3),
	"Hexadecimal":          NumStrFormatTypeCode(4),
	"Octal":                NumStrFormatTypeCode(5),
	"SignedNumber":         NumStrFormatTypeCode(6),
	"Signed Number Value":  NumStrFormatTypeCode(6),
	"SignedNumericValue":   NumStrFormatTypeCode(6),
	"Signed Numeric Value": NumStrFormatTypeCode(6),
	"Signed Number":        NumStrFormatTypeCode(6),
	"Signed":               NumStrFormatTypeCode(6),
	"ScientificNotation":   NumStrFormatTypeCode(7),
	"Scientific Notation":  NumStrFormatTypeCode(7),
	"SCI":                  NumStrFormatTypeCode(7),
	"Scientific Form":      NumStrFormatTypeCode(7),
	"Standard Index Form":  NumStrFormatTypeCode(7),
	"Standard Form":        NumStrFormatTypeCode(7),
}

var mNumStrFmtTypeCodeLwrCaseStringToCode = map[string]NumStrFormatTypeCode{
	"none":                 NumStrFormatTypeCode(0),
	"absolutevalue":        NumStrFormatTypeCode(1),
	"absolute value":       NumStrFormatTypeCode(1),
	"absolute":             NumStrFormatTypeCode(1),
	"binary":               NumStrFormatTypeCode(2),
	"currency":             NumStrFormatTypeCode(3),
	"currency value":       NumStrFormatTypeCode(3),
	"currencyvalue":        NumStrFormatTypeCode(3),
	"hexadecimal":          NumStrFormatTypeCode(4),
	"octal":                NumStrFormatTypeCode(5),
	"signednumber":         NumStrFormatTypeCode(6),
	"signed number value":  NumStrFormatTypeCode(6),
	"signednumericvalue":   NumStrFormatTypeCode(6),
	"signed numeric value": NumStrFormatTypeCode(6),
	"signed number":        NumStrFormatTypeCode(6),
	"signed":               NumStrFormatTypeCode(6),
	"scientificnotation":   NumStrFormatTypeCode(7),
	"scientific notation":  NumStrFormatTypeCode(7),
	"sci":                  NumStrFormatTypeCode(7),
	"scientific form":      NumStrFormatTypeCode(7),
	"standard index form":  NumStrFormatTypeCode(7),
	"standard form":        NumStrFormatTypeCode(7),
}

// NumStrFormatTypeCode - The 'Number String Format Type Code' is
// an enumeration of specification or formatting codes for the
// display of numeric values in text number strings.
//
// Since the Go Programming Language does not directly support
// enumerations, the 'NumStrFormatTypeCode' type has been adapted
// to function in a manner similar to classic enumerations.
//
// 'NumStrFormatTypeCode' is declared as a type 'int'. The method
// names effectively represent an enumeration of numeric format
// display specifications. These methods are listed as
// follows:
//
// Method             Integer
// Name                Value
// ------             -------
//
// None                 (0)
//   - Signals that the Number String Value Specification (NumStrFormatTypeCode)
//     Type is not initialized. This is an error condition.
//
// AbsoluteValue        (1)
//
//   - This format specification signals that a numeric value will
//     be displayed in text as a positive number regardless of
//     whether the native value is positive or negative.
//     Effectively, this means that both negative values and
//     positive values will be displayed as positive numbers.
//
//     Examples:
//     Positive Values          Negative Values
//     +132 = +132              -123 = +123
//
// Binary               (2)
//   - The 'Binary' format specification provides support for the
//     display of text number strings in base-16 or binary format.
//
// Currency             (3)
//
//   - The 'Currency' format specification signals that all numeric
//     values will be displayed in number strings as currency
//     formatted with appropriate currency characters.
//
//     Currency number strings are always displayed as signed
//     numeric values with currency symbols included in the text
//     string. This means that positive values are displayed in text
//     as positive numbers with currency symbols (like the dollar
//     sign) included in the text string. Likewise, negative values
//     are displayed in text as negative numbers with currency
//     symbols (like the dollar sign) included in the text string.
//
//     Examples:
//     Positive Values          Negative Values
//     +132 = $132               -123 = ($123)
//
// Hexadecimal          (4)
//   - The 'Hexadecimal' format specification provides support for
//     the display of text number strings in base-16 or hexadecimal
//     format.
//
// Octal                (5)
//   - The 'Octal' format specification provides support for the
//     display of text number strings in base-8 or octal format.
//
// SignedNumber         (6)
//
//   - Signals that the numeric value will be displayed in text as a
//     standard positive or negative value contingent upon the
//     number sign associated with the numeric value. NO CURRENCY
//     Symbols will be display in the resulting text number strings.
//
//     This is the default handling for numeric values.
//
//     'SignedNumber' means that positive values will be displayed
//     as positive numbers and negative values will be displayed as
//     negative numbers.
//
//     Examples:
//     Positive Values          Negative Values
//     +132 = 132               -123 = -123
//
// ScientificNotation   (7)
//
//   - Signals that the numeric value will be displayed in text as
//     Scientific Notation.
//
//     Examples: '2.652e+8'     '2.652e-8'
//
// For easy access to these enumeration values, use the global
// constant 'NumStrFmtType'.
//
//	Example: NumStrFmtType.SignedNumber()
//
// Otherwise you will need to use the formal syntax.
//
//	Example: NumStrFormatTypeCode(0).SignedNumber()
//
// Depending on your editor, intellisense (a.k.a. intelligent code
// completion) may not list the NumStrFormatTypeCode methods in
// alphabetical order.
//
// Be advised that all 'NumStrFormatTypeCode' methods beginning
// with 'X', as well as the method 'String()', are utility methods
// and not part of the enumeration.
type NumStrFormatTypeCode int

var lockNumStrFormatTypeCode sync.Mutex

// None - Signals that the NumStrFormatTypeCode Type is uninitialized.
// This is an error condition.
//
// This method is part of the standard enumeration.
func (nStrValSpec NumStrFormatTypeCode) None() NumStrFormatTypeCode {

	lockNumStrFormatTypeCode.Lock()

	defer lockNumStrFormatTypeCode.Unlock()

	return NumStrFormatTypeCode(0)
}

// AbsoluteValue - The 'Absolute Integer Value' specification
// signals that both positive and negative integer numeric values
// will be displayed in number strings as positive numbers.
//
//	Examples:
//	Positive Values          Negative Values
//	 +132 = +132              -123 = +123
//
// This method is part of the standard enumeration.
func (nStrValSpec NumStrFormatTypeCode) AbsoluteValue() NumStrFormatTypeCode {

	lockNumStrFormatTypeCode.Lock()

	defer lockNumStrFormatTypeCode.Unlock()

	return NumStrFormatTypeCode(1)
}

// Binary - The 'Binary' format type provides support for number
// strings formatted as binary or base-2 numeric values.
//
//	Examples: 00001101
//
// This method is part of the standard enumeration.
func (nStrValSpec NumStrFormatTypeCode) Binary() NumStrFormatTypeCode {

	lockNumStrFormatTypeCode.Lock()

	defer lockNumStrFormatTypeCode.Unlock()

	return NumStrFormatTypeCode(2)
}

// Currency - The 'Currency' format specification
// signals that all numeric values will be displayed in
// number strings as currency formatted with appropriate
// currency characters.
//
// Currency number strings are always displayed as signed
// numeric values with currency symbols included in the
// text string.
//
// This means that positive values are displayed in text
// as positive numbers with currency symbols (like the
// dollar sign) included in the text string.
//
// Likewise, negative values are displayed in text as
// negative numbers with currency symbols (like the dollar
// sign) included in the text string.
//
// Examples:
//
//	Positive Values          Negative Values
//	 +132 = $132              -123 = ($123)
//
// Note: Placement and formatting characters such as
// plus ('+'), minus ('-') parentheses ('()') and
// currency ('$') is controlled by type,
// 'FormatterCurrency'.
//
// This method is part of the standard enumeration.
func (nStrValSpec NumStrFormatTypeCode) Currency() NumStrFormatTypeCode {

	lockNumStrFormatTypeCode.Lock()

	defer lockNumStrFormatTypeCode.Unlock()

	return NumStrFormatTypeCode(3)
}

// Hexadecimal - The 'Hexadecimal' format type provides support for
// number strings formatted as base-16, hexadecimal numeric values.
//
// This is the default for processing and converting numeric
// values into text strings.
//
// This method is part of the standard enumeration.
func (nStrValSpec NumStrFormatTypeCode) Hexadecimal() NumStrFormatTypeCode {

	lockNumStrFormatTypeCode.Lock()

	defer lockNumStrFormatTypeCode.Unlock()

	return NumStrFormatTypeCode(4)
}

// Octal - The 'Octal' format type provides support for
// number strings formatted as base-8, octal numeric values.
//
// This is the default for processing and converting numeric
// values into text strings.
//
// This method is part of the standard enumeration.
func (nStrValSpec NumStrFormatTypeCode) Octal() NumStrFormatTypeCode {

	lockNumStrFormatTypeCode.Lock()

	defer lockNumStrFormatTypeCode.Unlock()

	return NumStrFormatTypeCode(5)
}

// SignedNumber - The 'Signed Number' specification
// signals that positive numeric values will be displayed in
// number strings as positive numbers and negative numeric
// values will be displayed as negative numbers.
//
// This is the default for processing and converting numeric
// values into text strings.
//
// This method is part of the standard enumeration.
func (nStrValSpec NumStrFormatTypeCode) SignedNumber() NumStrFormatTypeCode {

	lockNumStrFormatTypeCode.Lock()

	defer lockNumStrFormatTypeCode.Unlock()

	return NumStrFormatTypeCode(6)
}

// ScientificNotation - The 'Scientific Notation' specification
// signals that numeric values, both positive and negative, will
// be displayed in text number strings as Scientific Notation.
//
//	Example Text Display:
//	    "2.652e+8"
//	    "2.652e-8"
//
// Reference:
//
//	https://en.wikipedia.org/wiki/Scientific_notation
//
// This is the default for processing and converting numeric
// values into text strings.
//
// This method is part of the standard enumeration.
func (nStrValSpec NumStrFormatTypeCode) ScientificNotation() NumStrFormatTypeCode {

	lockNumStrFormatTypeCode.Lock()

	defer lockNumStrFormatTypeCode.Unlock()

	return NumStrFormatTypeCode(7)
}

// String - Returns a string with the name of the enumeration associated
// with this current instance of 'NumStrFormatTypeCode'.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// # Usage
//
// t:= NumStrFormatTypeCode(0).SignedNumber()
// str := t.String()
//
//	str is now equal to 'SignedNumber'
func (nStrValSpec NumStrFormatTypeCode) String() string {

	lockNumStrFormatTypeCode.Lock()

	defer lockNumStrFormatTypeCode.Unlock()

	result, ok := mNumStrFmtTypeCodeToString[nStrValSpec]

	if !ok {
		return "Error: Number String Format Type Code Specification UNKNOWN!"
	}

	return result
}

// XIsValid - Returns a boolean value signaling whether the current
// Number String Format Type Code Specification
// (NumStrFormatTypeCode) is valid.
//
// This is a standard utility method and is not part of the valid enumerations
// for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	nStrValueSpec := NumStrFormatTypeCode(0).AbsoluteValue()
//
//	isValid := nStrValueSpec.XIsValid()
//
//	In this case the boolean value of 'isValid' is 'true'.
//
//	Be advised, the value NumStrFormatTypeCode(0).None() is
//	classified as an 'invalid' value.
func (nStrValSpec NumStrFormatTypeCode) XIsValid() bool {

	lockNumStrFormatTypeCode.Lock()

	defer lockNumStrFormatTypeCode.Unlock()

	return new(numStrFmtTypeCodeNanobot).isValidNumStrFmtTypeCode(
		nStrValSpec)
}

// XParseString - Receives a string and attempts to match it with the
// string value of a supported enumeration. If successful, a new
// instance of NumStrFormatTypeCode is returned set to the value of the
// associated enumeration.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ------------------------------------------------------------------------
//
// # Input Parameters
//
// valueString   string
//   - A string which will be matched against the enumeration string
//     values. If 'valueString' is equal to one of the enumeration
//     names, this method will proceed to successful completion and
//     return the correct enumeration value.
//
// caseSensitive   bool
//
//   - If 'true' the search for enumeration names will be
//     case-sensitive and will require an exact match. Therefore,
//     'scientificnotation' will NOT match the enumeration name,
//     'ScientificNotation'.
//
//     A case-sensitive search will match any of the following strings:
//     "None"
//     "AbsoluteValue"
//     "Absolute Value"
//     "Absolute"
//     "Binary"
//     "Currency"
//     "Currency Value"
//     "CurrencyValue"
//     "Hexadecimal"
//     "Octal"
//     "SignedNumber"
//     "Signed Number Value"
//     "SignedNumericValue"
//     "Signed Numeric Value"
//     "Signed Number"
//     "Signed"
//     "ScientificNotation"
//     "Scientific Notation"
//     "SCI"
//
//     If 'false', a case-insensitive search is conducted for the
//     enumeration name. In this example, 'scientificnotation'
//     will match enumeration name 'ScientificNotation'.
//
//     A case-insensitive search will match any of the following
//     lower case names:
//     "none"
//     "absolutevalue"
//     "absolute value"
//     "absolute"
//     "binary"
//     "currency"
//     "currency value"
//     "currencyvalue"
//     "hexadecimal"
//     "octal"
//     "signednumber"
//     "signed number value"
//     "signednumericvalue"
//     "signed numeric value"
//     "signed number"
//     "signed"
//     "scientificnotation"
//     "scientific notation"
//     "sci"
//     "scientific form"
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	NumStrFormatTypeCode
//	   - Upon successful completion, this method will return a new
//	     instance of NumStrFormatTypeCode set to the value of the enumeration
//	     matched by the string search performed on input parameter,
//	     'valueString'.
//
//	error
//	   - If this method completes successfully, the returned error
//	     Type is set equal to 'nil'. If an error condition is encountered,
//	     this method will return an error type which encapsulates an
//	     appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	t, err := NumStrFormatTypeCode(0).XParseString("SignedNumber", true)
//
//	   t is now equal to NumStrFormatTypeCode(0).SignedNumber()
func (nStrValSpec NumStrFormatTypeCode) XParseString(
	valueString string,
	caseSensitive bool) (NumStrFormatTypeCode, error) {

	lockNumStrFormatTypeCode.Lock()

	defer lockNumStrFormatTypeCode.Unlock()

	ePrefix := "NumStrFormatTypeCode.XParseString() "

	if len(valueString) < 3 {
		return NumStrFormatTypeCode(0),
			fmt.Errorf(ePrefix+
				"\nInput parameter 'valueString' is INVALID!\n"+
				"String length is less than '3'.\n"+
				"valueString='%v'\n", valueString)
	}

	var ok bool
	var nStrValueSpec NumStrFormatTypeCode

	if caseSensitive {

		nStrValueSpec, ok = mNumStrFmtTypeCodeStringToCode[valueString]

		if !ok {
			return NumStrFormatTypeCode(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid NumStrFormatTypeCode Value.\n"+
					"valueString='%v'\n", valueString)
		}

	} else {

		nStrValueSpec, ok = mNumStrFmtTypeCodeLwrCaseStringToCode[strings.ToLower(valueString)]

		if !ok {
			return NumStrFormatTypeCode(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid NumStrFormatTypeCode Value.\n"+
					"valueString='%v'\n", valueString)
		}
	}

	return nStrValueSpec, nil
}

// XReturnNoneIfInvalid - Provides a standardized value for invalid
// instances of enumeration NumStrFormatTypeCode.
//
// If the current instance of NumStrFormatTypeCode is invalid, this
// method will always return a value of
// NumStrFormatTypeCode(0).None().
//
// # Background
//
// Enumeration NumStrFormatTypeCode has an underlying type of
// integer (int). This means the type could conceivably be set to
// any integer value. This method ensures that all invalid
// NumStrFormatTypeCode instances are consistently classified as
// 'None' (NumStrFormatTypeCode(0).None()). Remember that 'None' is
// considered an invalid value.
//
// For example, assume that NumStrFormatTypeCode was set to an
// invalid integer value of -848972. Calling this method on a
// NumStrFormatTypeCode with this invalid integer value will return
// an integer value of zero or the equivalent of
// NumStrFormatTypeCode(0).None(). This conversion is useful in
// generating text strings for meaningful informational and error
// messages.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
func (nStrValSpec NumStrFormatTypeCode) XReturnNoneIfInvalid() NumStrFormatTypeCode {

	lockNumStrFormatTypeCode.Lock()

	defer lockNumStrFormatTypeCode.Unlock()

	isValid := new(numStrFmtTypeCodeNanobot).
		isValidNumStrFmtTypeCode(nStrValSpec)

	if !isValid {
		return NumStrFormatTypeCode(0)
	}

	return nStrValSpec
}

// XValue - This method returns the enumeration value of the
// current NumStrFormatTypeCode instance.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
func (nStrValSpec NumStrFormatTypeCode) XValue() NumStrFormatTypeCode {

	lockNumStrFormatTypeCode.Lock()

	defer lockNumStrFormatTypeCode.Unlock()

	return nStrValSpec
}

// XValueInt - This method returns the integer value of the current
// NumStrFormatTypeCode.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
func (nStrValSpec NumStrFormatTypeCode) XValueInt() int {

	lockNumStrFormatTypeCode.Lock()

	defer lockNumStrFormatTypeCode.Unlock()

	return int(nStrValSpec)
}

// NumStrFmtType - public global variable of type
// NumStrFormatTypeCode.
//
// This variable serves as an easier, shorthand
// technique for accessing NumStrFormatTypeCode values.
//
// For easy access to these enumeration values, use the
// global variable 'NumStrFmtType'.
//
//	Example: NumStrFmtType.SignedNumber()
//
// Otherwise you will need to use the formal syntax.
//
//	Example: NumStrFormatTypeCode(0).SignedNumber()
//
// Usage:
//
//	NumStrFmtType.None()
//	NumStrFmtType.AbsoluteValue()
//	NumStrFmtType.Binary()
//	NumStrFmtType.Currency()
//	NumStrFmtType.Hexadecimal()
//	NumStrFmtType.Octal()
//	NumStrFmtType.SignedNumber()
//	NumStrFmtType.ScientificNotation()
const NumStrFmtType = NumStrFormatTypeCode(0)

// numStrFmtTypeCodeNanobot - Provides helper methods for
// enumeration NumStrFormatTypeCode.
type numStrFmtTypeCodeNanobot struct {
	lock *sync.Mutex
}

// isValidNumStrFmtTypeCode - Receives an instance of
// NumStrFormatTypeCode and returns a boolean value signaling
// whether that NumStrFormatTypeCode instance is valid.
//
// If the passed instance of NumStrFormatTypeCode is valid, this
// method returns 'true'.
//
// Be advised, the enumeration value "None" is considered NOT
// VALID. "None" represents an error condition.
//
// This is a standard utility method and is not part of the valid
// NumStrFormatTypeCode enumeration.
func (numStrFmtTypeNanobot *numStrFmtTypeCodeNanobot) isValidNumStrFmtTypeCode(
	numStrFmtTypeCode NumStrFormatTypeCode) bool {

	if numStrFmtTypeNanobot.lock == nil {
		numStrFmtTypeNanobot.lock = new(sync.Mutex)
	}

	numStrFmtTypeNanobot.lock.Lock()

	defer numStrFmtTypeNanobot.lock.Unlock()

	if numStrFmtTypeCode < 1 ||
		numStrFmtTypeCode > 9 {

		return false
	}

	return true
}
