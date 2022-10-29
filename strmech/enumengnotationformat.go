package strmech

import (
	"fmt"
	"strings"
	"sync"
)

//	Lock lockEngineeringNotationFormat before accessing
//	these maps!

var mEngNotationFmtCodeToString = map[EngineeringNotationFormat]string{
	EngineeringNotationFormat(0): "None",
	EngineeringNotationFormat(1): "Exponential",
	EngineeringNotationFormat(2): "ENotUprCaseELeadPlus",
	EngineeringNotationFormat(3): "ENotUprCaseENoLeadPlus",
	EngineeringNotationFormat(4): "ENotLwrCaseELeadPlus",
	EngineeringNotationFormat(5): "ENotLwrCaseENoLeadPlus",
	EngineeringNotationFormat(6): "SIPrefixSymbol",
	EngineeringNotationFormat(7): "SIPrefixName",
}

var mEngNotationFmtStringToCode = map[string]EngineeringNotationFormat{
	"None":                   EngineeringNotationFormat(0),
	"Exponential":            EngineeringNotationFormat(1),
	"ENotUprCaseELeadPlus":   EngineeringNotationFormat(2),
	"ENotUprCaseENoLeadPlus": EngineeringNotationFormat(3),
	"ENotLwrCaseELeadPlus":   EngineeringNotationFormat(4),
	"ENotLwrCaseENoLeadPlus": EngineeringNotationFormat(5),
	"SIPrefixSymbol":         EngineeringNotationFormat(6),
	"SIPrefixName":           EngineeringNotationFormat(7),
}

var mEngNotationFmtLwrCaseStringToCode = map[string]EngineeringNotationFormat{
	"none":                   EngineeringNotationFormat(0),
	"exponential":            EngineeringNotationFormat(1),
	"enotuprcaseeleadplus":   EngineeringNotationFormat(2),
	"enotuprcaseenoleadplus": EngineeringNotationFormat(3),
	"enotlwrcaseeleadplus":   EngineeringNotationFormat(4),
	"enotlwrcaseenoleadplus": EngineeringNotationFormat(5),
	"siprefixsymbol":         EngineeringNotationFormat(6),
	"siprefixname":           EngineeringNotationFormat(7),
}

//	EngineeringNotationFormat
//
//
//	The 'Engineering Notation Format' is an enumeration
//	of type codes used for classifying the output or
//	display formats for Engineering Notation.
//
// ----------------------------------------------------------------
//
// # Terminology
//
//	Engineering notation or engineering form is a version
//	of scientific notation in which the exponent of ten
//	must be divisible by three (i.e., they are powers of
//	a thousand, but written as, for example, 106 instead
//	of 10002).
//
//	As an alternative to writing powers of 10, SI or
//	Metric prefixes can be used, which also usually
//	provide steps of a factor of a thousand.
//
//	On most calculators, engineering notation is called
//	"ENG" mode.
//
//		Wikipedia
//		https://en.wikipedia.org/wiki/Engineering_notation
//
//	In Engineering Notation numeric values like
//	2,652,000,000 are either displayed in the Exponential
//	Format "2.652 x 10^9", E-Notation Format "2.652e+9" or
//	SI (Metric) Format "2.652 G".
//
//	Reference:
//		https://en.wikipedia.org/wiki/Engineering_notation#E_notation
//
// ----------------------------------------------------------------
//
// # Enumeration Values
//
//	Since the Go Programming Language does not directly
//	support enumerations, the EngineeringNotationFormat
//	has been adapted to function in a manner similar to
//	classic enumerations.
//
//	EngineeringNotationFormat is declared as a type
//	'int'. The method names associated with this type
//	effectively represent an enumeration of Engineering
//	Notation Format Types. These methods are listed as
//	follows:
//
//	Method				 Integer
//	 Name				  Value
//	------			 	 -------
//
//	None	   	   			0
//
//		Signals that 'EngineeringNotationFormat' has
//		not been initialized and therefore has no value.
//		This is an error condition.
//
//	Exponential	   			1
//
//		Signals that an Engineering Notation value will be
//		displayed using the Exponential format.
//
//		Exponential Format Example
//
//					Base Numeric Value: 2,652,000,000
//
//			Exponential Display Format: "2.652 x 10^9"
//
//		Exponential notation is displayed in the form
//		"m × 10n" or m times ten raised to the power of
//		n, where n is an integer, and the coefficient m
//		is a nonzero real number (usually between 1 and
//		10 in absolute value, and nearly always written
//		as a terminating decimal).
//
//		Reference:
//			https://en.wikipedia.org/wiki/Engineering_notation
//
//	ENotUprCaseELeadPlus	2
//
//		E-Notation Upper Case 'E' and Positive Exponent
//		Leading	Plus Sign.
//
//		ENotUprCaseELeadPlus Format Example
//
//						  Base Numeric Value: 2,652,000,000
//
//			E-NotationUprCase Display Format: "2.652E+9"
//
//		ENotUprCaseELeadPlus is a type of scientific
//		notation in which the phrase “times 10 to the power
//		of” is replaced by the upper case letter, 'E'.
//
//		Positive exponents have a leading plus sign (+).
//
//		For example, 5.1 × 10^+6 is written 5.1E+6.
//
//		"UprCase" in "ENotUprCaseELeadPlus" refers to the
//		upper case 'E' used in the E-Notation format.
//
//		"LeadPlus" in "ENotUprCaseELeadPlus" means that
//		positive exponents are formatted with a leading
//		plus sign ('+').
//
//	ENotUprCaseENoLeadPlus	3
//
//		E-Notation Upper Case 'E' and Positive Exponents
//		HAVE NO Leading Plus Sign.
//
//		ENotUprCaseENoLeadPlus Format Example
//
//						  Base Numeric Value: 2,652,000,000
//
//			E-NotationUprCase Display Format: "2.652E9"
//
//		ENotUprCaseENoLeadPlus is a type of scientific
//		notation in which the phrase “times 10 to the power
//		of” is replaced by the upper case letter, 'E'.
//
//		Positive exponents DO NOT HAVE a leading plus sign
//		(+).
//
//		For example, 5.1 × 10^9 is written 5.1E9.
//
//		"UprCase" in "ENotUprCaseENoLeadPlus" refers to the
//		upper case 'E' used in the E-Notation format.
//
//		"NoLeadPlus" in "ENotUprCaseENoLeadPlus" means that
//		positive exponents are NOT formatted with a leading
//		plus sign ('+').
//
//	ENotLwrCaseELeadPlus	4
//
//		E-Notation Lower Case 'e' and Positive Exponents
//		HAVE Leading Plus Signs.
//
//		E-Notation denotes the use of the letter 'e' in
//		the formatted Engineering Notation output display.
//
//		ENotLwrCaseELeadPlus Format Example
//
//					 Base Numeric Value: 2,652,000,000
//
//			ENotLwrCaseELeadPlus Format: "2.652e+9"
//
//		ENotLwrCaseELeadPlus is a type of scientific
//		notation in which the phrase “times 10 to the power
//		of” is replaced by the lower case letter, 'e'.
//
//		Positive exponents have a leading plus sign (+).
//
//		For example, 5.1 × 10^9 is written 5.1e+9.
//
//		"LwrCase" in "ENotLwrCaseELeadPlus" refers to the
//		lower case 'e' used in the E-Notation format.
//
//		"LeadPlus" in "ENotLwrCaseELeadPlus" means that
//		positive exponents HAVE leading plus signs ('+').
//
//	ENotLwrCaseENoLeadPlus	5
//
//		E-Notation Lower Case 'e' and Positive Exponent
//		HAVE NO Leading Plus Sign.
//
//		E-Notation denotes the use of the letter 'e' in
//		the formatted Engineering Notation output display.
//
//		ENotLwrCaseENoLeadPlus Example
//
//					 Base Numeric Value: 2,652,000,000
//
//			ENotLwrCaseELeadPlus Format: "2.652e9"
//
//		ENotLwrCaseENoLeadPlus is a type of scientific
//		notation in which the phrase “times 10 to the power
//		of” is replaced by the lower case letter, 'e'.
//
//		Positive exponents DO NOT have a leading plus sign
//		(+).
//
//		For example, 5.1 × 10^9 is written 5.1e9.
//
//		"LwrCase" in "ENotLwrCaseENoLeadPlus" refers to the
//		lower case 'e' used in the E-Notation format.
//
//		"NoLeadPlus" in "ENotLwrCaseENoLeadPlus" means that
//		positive exponents ARE NOT formatted with a leading
//		plus sign ('+').
//
//		ENotLwrCaseENoLeadPlus Example
//
//					 Base Numeric Value: 2,652,000,000
//
//			ENotLwrCaseELeadPlus Format: "2.652e9"
//
//	SIPrefixSymbol			6
//
//		Specifies the use of 'SI' Prefix Symbol in the
//		formatted display output for Engineering Notation
//		numeric values.
//
//		SI stands for the International System of Units.
//		SI is also referred to as a Metric prefix.
//
//		In the arithmetic of measurements having units, the
//		units are treated as multiplicative factors to
//		values. If they have prefixes, all but one of the
//		prefixes must be expanded to their numeric
//		multiplier, except when combining values with
//		identical units. Hence:
//
//			In SI Symbology 'm' = 10^−3
//
//			5 m = 0.005 = 5 x 10^-3 = 5 m
//
//		SIPrefixSymbol Example
//
//					   Base Numeric Value: 2,650
//
//			ENotLwrCaseENoLeadPlus Format: "2.65 k"
//
//		In SI Symbology 'k' (kilo) Equals: "10^3"
//
//
//	SIPrefixName			7
//
//		Specifies the use of 'SI' Prefix Name in the
//		formatted display output for Engineering Notation
//		numeric values.
//
//		SI stands for the International System of Units.
//		SI is also referred to as a Metric prefix.
//
//		SIPrefixName Example
//
//					   Base Numeric Value: 2,650
//
//			ENotLwrCaseENoLeadPlus Format: "2.65 kilo"
//
//			In SI Symbology 'kilo' Equals: "10^3"
//
//		In the arithmetic of measurements having units, the
//		units are treated as multiplicative factors to
//		values. If they have prefixes, all but one of the
//		prefixes must be expanded to their numeric
//		multiplier, except when combining values with
//		identical units. Hence:
//
//			In SI Symbology 'milli' = 10^−3
//
//			5 m = 0.005 = 5 x 10^-3 = 5 milli
//
//		SIPrefixName Example
//
//					   Base Numeric Value: 2,650
//
//			ENotLwrCaseENoLeadPlus Format: "2.65 kilo"
//
//			In SI Symbology 'kilo' Equals: "10^3"
//
// ----------------------------------------------------------------
//
// # Usage
//
//	For easy access to these enumeration values, use the
//	global constant EngNotFmt.
//
//		Example:
//			EngNotFmt.ENotUprCaseELeadPlus()
//
//	Otherwise you will need to use the formal syntax.
//
//		Example:
//			EngineeringNotationFormat(0).ENotUprCaseELeadPlus()
//
//	Depending on your editor, intellisense (a.k.a.
//	intelligent code completion) may not list the
//	EngineeringNotationFormat methods in alphabetical
//	order.
//
//	Be advised that all EngineeringNotationFormat
//	methods beginning with 'X', as well as the method
//	'String()', are utility methods, and NOT part of the
//	enumeration values.
type EngineeringNotationFormat int

var lockEngineeringNotationFormat sync.Mutex

// None
//
// Signals that 'EngineeringNotationFormat' has not been
// initialized and therefore has no value.
//
// This is an error condition.
func (engNotationFmt EngineeringNotationFormat) None() EngineeringNotationFormat {

	lockEngineeringNotationFormat.Lock()

	defer lockEngineeringNotationFormat.Unlock()

	return EngineeringNotationFormat(0)
}

//	Exponential
//
//	Signals that an Engineering Notation value will be
//	displayed using the Exponential format.
//
//	Exponential Format Example
//
//				Base Numeric Value: 2,652,000,000
//
//		Exponential Display Format: "2.652 x 10^9"
//
//	Exponential notation is displayed in the form
//	"m × 10n" or m times ten raised to the power of n,
//	where n is an integer, and the coefficient m is a
//	nonzero real number (usually between 1 and 10 in
//	absolute value, and nearly always written as a
//	terminating decimal).
//
//	This method is part of the EngineeringNotationFormat
//	enumeration.
//
// ----------------------------------------------------------------
//
//	Reference:
//
//	https://en.wikipedia.org/wiki/Engineering_notation
func (engNotationFmt EngineeringNotationFormat) Exponential() EngineeringNotationFormat {

	lockEngineeringNotationFormat.Lock()

	defer lockEngineeringNotationFormat.Unlock()

	return EngineeringNotationFormat(1)
}

//	ENotUprCaseELeadPlus
//
//	E-Notation Upper Case 'E' and Positive Exponent
//	Leading	Plus Sign.
//
//	E-Notation denotes the use of the letter 'E' in
//	the formatted Engineering Notation output display.
//
//	ENotUprCaseELeadPlus Format Example
//
//					  Base Numeric Value: 2,652,000,000
//
//		E-NotationUprCase Display Format: "2.652E+9"
//
//	ENotUprCaseELeadPlus is a type of scientific
//	notation in which the phrase “times 10 to the power
//	of” is replaced by the upper case letter, 'E'.
//
//	Positive exponents have a leading plus sign (+).
//
//	For example, 5.1 × 10^+6 is written 5.1E+6.
//
//	"UprCase" in "ENotUprCaseELeadPlus" refers to the
//	upper case 'E' used in the E-Notation format.
//
//	"LeadPlus" in "ENotUprCaseELeadPlus" means that
//	positive exponents are formatted with a leading
//	plus sign ('+').
//
//	ENotUprCaseELeadPlus Format Example
//
//					  Base Numeric Value: 2,652,000,000
//
//		E-NotationUprCase Display Format: "2.652E+9"
//
//	This method is part of the EngineeringNotationFormat
//	enumeration.
//
// ----------------------------------------------------------------
//
//	Reference:
//
//	https://encyclopedia2.thefreedictionary.com/E+notation
//	https://en.wikipedia.org/wiki/Engineering_notation#E_notation
//	https://en.wikipedia.org/wiki/Engineering_notation#E_notation
func (engNotationFmt EngineeringNotationFormat) ENotUprCaseELeadPlus() EngineeringNotationFormat {

	lockEngineeringNotationFormat.Lock()

	defer lockEngineeringNotationFormat.Unlock()

	return EngineeringNotationFormat(2)
}

//	ENotUprCaseENoLeadPlus
//
//	E-Notation Upper Case 'E' and Positive Exponents
//	HAVE NO Leading Plus Sign.
//
//	ENotUprCaseENoLeadPlus Format Example
//
//					  Base Numeric Value: 2,652,000,000
//
//		E-NotationUprCase Display Format: "2.652E9"
//
//	ENotUprCaseENoLeadPlus is a type of scientific
//	notation in which the phrase “times 10 to the power
//	of” is replaced by the upper case letter, 'E'.
//
//	Positive exponents DO NOT HAVE a leading plus sign
//	(+).
//
//	For example, 5.1 × 10^9 is written 5.1E9.
//
//	"UprCase" in "ENotUprCaseENoLeadPlus" refers to the
//	upper case 'E' used in the E-Notation format.
//
//	"NoLeadPlus" in "ENotUprCaseENoLeadPlus" means that
//	positive exponents are NOT formatted with a leading
//	plus sign ('+').
//
//	ENotUprCaseENoLeadPlus Format Example
//
//					  Base Numeric Value: 2,652,000,000
//
//		E-NotationUprCase Display Format: "2.652E9"
//
//	This method is part of the EngineeringNotationFormat
//	enumeration.
//
// ----------------------------------------------------------------
//
//	Reference:
//
//	https://encyclopedia2.thefreedictionary.com/E+notation
//	https://en.wikipedia.org/wiki/Engineering_notation#E_notation
//	https://en.wikipedia.org/wiki/Engineering_notation#E_notation
func (engNotationFmt EngineeringNotationFormat) ENotUprCaseENoLeadPlus() EngineeringNotationFormat {

	lockEngineeringNotationFormat.Lock()

	defer lockEngineeringNotationFormat.Unlock()

	return EngineeringNotationFormat(3)
}

//	ENotLwrCaseELeadPlus
//
//	E-Notation Lower Case 'e' and Positive Exponents
//	HAVE Leading Plus Signs.
//
//	E-Notation denotes the use of the letter 'e' in
//	the formatted Engineering Notation output display.
//
//	ENotLwrCaseELeadPlus Format Example
//
//				 Base Numeric Value: 2,652,000,000
//
//		ENotLwrCaseELeadPlus Format: "2.652e+9"
//
//	ENotLwrCaseELeadPlus is a type of scientific
//	notation in which the phrase “times 10 to the power
//	of” is replaced by the lower case letter, 'e'.
//
//	Positive exponents have a leading plus sign (+).
//
//	For example, 5.1 × 10^9 is written 5.1e+9.
//
//	"LwrCase" in "ENotLwrCaseELeadPlus" refers to the
//	lower case 'e' used in the E-Notation format.
//
//	"LeadPlus" in "ENotLwrCaseELeadPlus" means that
//	positive exponents HAVE leading plus signs ('+').
//
//	ENotLwrCaseELeadPlusFormat Example
//
//				 Base Numeric Value: 2,652,000,000
//
//		ENotLwrCaseELeadPlus Format: "2.652e+9"
//
//	This method is part of the EngineeringNotationFormat
//	enumeration.
//
// ----------------------------------------------------------------
//
//	Reference:
//
//	https://encyclopedia2.thefreedictionary.com/E+notation
//	https://en.wikipedia.org/wiki/Engineering_notation#E_notation
func (engNotationFmt EngineeringNotationFormat) ENotLwrCaseELeadPlus() EngineeringNotationFormat {

	lockEngineeringNotationFormat.Lock()

	defer lockEngineeringNotationFormat.Unlock()

	return EngineeringNotationFormat(4)
}

//	ENotLwrCaseENoLeadPlus
//
//	E-Notation Lower Case 'e' and Positive Exponent
//	HAVE NO Leading Plus Sign.
//
//	E-Notation denotes the use of the letter 'e' in
//	the formatted Engineering Notation output display.
//
//	ENotLwrCaseENoLeadPlus Example
//
//				 Base Numeric Value: 2,652,000,000
//
//		ENotLwrCaseELeadPlus Format: "2.652e9"
//
//	ENotLwrCaseENoLeadPlus is a type of scientific
//	notation in which the phrase “times 10 to the power
//	of” is replaced by the lower case letter, 'e'.
//
//	Positive exponents DO NOT have a leading plus sign
//	(+).
//
//	For example, 5.1 × 10^9 is written 5.1e9.
//
//	"LwrCase" in "ENotLwrCaseENoLeadPlus" refers to the
//	lower case 'e' used in the E-Notation format.
//
//	"NoLeadPlus" in "ENotLwrCaseENoLeadPlus" means that
//	positive exponents ARE NOT formatted with a leading
//	plus sign ('+').
//
//	ENotLwrCaseENoLeadPlus Example
//
//				 Base Numeric Value: 2,652,000,000
//
//		ENotLwrCaseELeadPlus Format: "2.652e9"
//
//	This method is part of the EngineeringNotationFormat
//	enumeration.
//
// ----------------------------------------------------------------
//
//	Reference:
//
//	https://en.wikipedia.org/wiki/Engineering_notation
//	https://encyclopedia2.thefreedictionary.com/E+notation
//	https://en.wikipedia.org/wiki/Engineering_notation#E_notation
func (engNotationFmt EngineeringNotationFormat) ENotLwrCaseENoLeadPlus() EngineeringNotationFormat {

	lockEngineeringNotationFormat.Lock()

	defer lockEngineeringNotationFormat.Unlock()

	return EngineeringNotationFormat(5)
}

//	SIPrefixSymbol
//
//	Specifies the use of 'SI' Prefix Symbol in the
//	formatted display output for Engineering Notation
//	numeric values.
//
//		SI stands for the International System of Units.
//		SI is also referred to as a Metric prefix.
//
//	SIPrefixSymbol Example
//
//				   Base Numeric Value: 2,650
//
//		ENotLwrCaseENoLeadPlus Format: "2.65 k"
//
//	In SI Symbology 'k' (kilo) Equals: "10^3"
//
//	In the arithmetic of measurements having units, the
//	units are treated as multiplicative factors to
//	values. If they have prefixes, all but one of the
//	prefixes must be expanded to their numeric
//	multiplier, except when combining values with
//	identical units. Hence:
//
//		In SI Symbology 'm' = 10^−3
//
//		5 m = 0.005 = 5 x 10^-3 = 5 m
//
//	SIPrefixSymbol Example
//
//				   Base Numeric Value: 2,650
//
//		ENotLwrCaseENoLeadPlus Format: "2.65 k"
//
//	In SI Symbology 'k' (kilo) Equals: "10^3"
//
//	SI is also referred to as a Metric prefix. A metric
//	prefix is a unit prefix that precedes a basic unit
//	of measure to indicate a multiple or sub-multiple
//	of the unit.
//
//	All metric prefixes used today are decadic (of or
//	relating to the decimal system of counting). Each
//	prefix has a unique symbol that is prepended to
//	any unit symbol.
//
//						SI prefixes
//					Prefix	Representations
//
//						Base 	 Base
//		Name	Symbol	1000	  10		Value
//		----	------  -----	-----	-------------
//		yotta	  Y		1000^8	 10^24	1 000 000 000 000 000 000 000 000
//		zetta	  Z		1000^7	 10^21	1 000 000 000 000 000 000 000
//		exa		  E		1000^6	 10^18	1 000 000 000 000 000 000
//		peta 	  P		1000^5	 10^15	1 000 000 000 000 000
//		tera	  T		1000^4	 10^12	1 000 000 000 000
//		giga	  G		1000^3	 10^9	1 000 000 000
//		mega	  M		1000^2	 10^6	1 000 000
//		kilo	  k		1000^1	 10^3	1 000
//						1000^0	 10^0	1
//		milli	  m		1000^−1	 10^−3	0.001
//		micro	  μ		1000^−2	 10^−6	0.000 001
//		nano	  n		1000^−3	 10^−9	0.000 000 001
//		pico	  p		1000^−4	 10^−12	0.000 000 000 001
//		femto	  f		1000^−5	 10^−15	0.000 000 000 000 001
//		atto	  a		1000^−6	 10^−18	0.000 000 000 000 000 001
//		zepto	  z		1000^−7	 10^−21	0.000 000 000 000 000 000 001
//		yocto	  y		1000^−8	 10^−24	0.000 000 000 000 000 000 000 001
//
//	This method is part of the EngineeringNotationFormat
//	enumeration.
//
// ----------------------------------------------------------------
//
//	Reference:
//
//	https://www.nist.gov/pml/owm/metric-si-prefixes
//	https://en.wikipedia.org/wiki/Metric_prefix
//	https://en.wikipedia.org/wiki/International_System_of_Units
func (engNotationFmt EngineeringNotationFormat) SIPrefixSymbol() EngineeringNotationFormat {

	lockEngineeringNotationFormat.Lock()

	defer lockEngineeringNotationFormat.Unlock()

	return EngineeringNotationFormat(6)
}

//	SIPrefixName
//
//	Specifies the use of 'SI' Prefix Name in the
//	formatted display output for Engineering Notation
//	numeric values.
//
//	SI stands for the International System of Units.
//	SI is also referred to as a Metric prefix.
//
//	SIPrefixName Example
//
//				   Base Numeric Value: 2,650
//
//		ENotLwrCaseENoLeadPlus Format: "2.65 kilo"
//
//		In SI Symbology 'kilo' Equals: "10^3"
//
//	In the arithmetic of measurements having units, the
//	units are treated as multiplicative factors to
//	values. If they have prefixes, all but one of the
//	prefixes must be expanded to their numeric
//	multiplier, except when combining values with
//	identical units. Hence:
//
//		In SI Symbology 'milli' = 10^−3
//
//		5 m = 0.005 = 5 x 10^-3 = 5 milli
//
//	SIPrefixName Example
//
//				   Base Numeric Value: 2,650
//
//		ENotLwrCaseENoLeadPlus Format: "2.65 kilo"
//
//		In SI Symbology 'kilo' Equals: "10^3"
//
//	SI is also referred to as a Metric prefix. A metric
//	prefix is a unit prefix that precedes a basic unit
//	of measure to indicate a multiple or sub-multiple
//	of the unit.
//
//	All metric prefixes used today are decadic (of or
//	relating to the decimal system of counting). Each
//	prefix has a unique symbol that is prepended to
//	any unit symbol.
//
//						SI prefixes
//					Prefix	Representations
//
//						Base 	 Base
//		Name	Symbol	1000	  10		Value
//		----	------  -----	-----	-------------
//		yotta	  Y		1000^8	 10^24	1 000 000 000 000 000 000 000 000
//		zetta	  Z		1000^7	 10^21	1 000 000 000 000 000 000 000
//		exa		  E		1000^6	 10^18	1 000 000 000 000 000 000
//		peta 	  P		1000^5	 10^15	1 000 000 000 000 000
//		tera	  T		1000^4	 10^12	1 000 000 000 000
//		giga	  G		1000^3	 10^9	1 000 000 000
//		mega	  M		1000^2	 10^6	1 000 000
//		kilo	  k		1000^1	 10^3	1 000
//						1000^0	 10^0	1
//		milli	  m		1000^−1	 10^−3	0.001
//		micro	  μ		1000^−2	 10^−6	0.000 001
//		nano	  n		1000^−3	 10^−9	0.000 000 001
//		pico	  p		1000^−4	 10^−12	0.000 000 000 001
//		femto	  f		1000^−5	 10^−15	0.000 000 000 000 001
//		atto	  a		1000^−6	 10^−18	0.000 000 000 000 000 001
//		zepto	  z		1000^−7	 10^−21	0.000 000 000 000 000 000 001
//		yocto	  y		1000^−8	 10^−24	0.000 000 000 000 000 000 000 001
//
//	This method is part of the EngineeringNotationFormat
//	enumeration.
//
// ----------------------------------------------------------------
//
//	Reference:
//
//	https://www.nist.gov/pml/owm/metric-si-prefixes
//	https://en.wikipedia.org/wiki/Metric_prefix
//	https://en.wikipedia.org/wiki/International_System_of_Units
func (engNotationFmt EngineeringNotationFormat) SIPrefixName() EngineeringNotationFormat {

	lockEngineeringNotationFormat.Lock()

	defer lockEngineeringNotationFormat.Unlock()

	return EngineeringNotationFormat(7)
}

//	String
//
//	Returns a string with the name of the enumeration
//	associated with this current instance of
//	'EngineeringNotationFormat'.
//
//	This is a standard utility method and is not part of
//	the valid enumerations for this type.
//
// ----------------------------------------------------------------
//
// # Usage
//
// t:= EngineeringNotationFormat(0).Exponential()
// str := t.String()
//
//	str is now equal to 'Exponential'
func (engNotationFmt EngineeringNotationFormat) String() string {

	lockEngineeringNotationFormat.Lock()

	defer lockEngineeringNotationFormat.Unlock()

	result, ok := mEngNotationFmtCodeToString[engNotationFmt]

	if !ok {

		return "Error: Engineering Notation Format code is UNKNOWN!"

	}

	return result
}

//	XIsValid
//
//	Returns a boolean value signaling whether the current
//	EngineeringNotationFormat value is valid.
//
//	Be advised, the enumeration value "None" is
//	considered an INVALID selection for
//	'EngineeringNotationFormat'.
//
//	This is a standard utility method and is not part of
//	the valid enumerations for this type.
//
// ------------------------------------------------------------------------
//
// # Usage
//
//	 sciNotFmt :=
//				EngineeringNotationFormat(0).ENotation()
//
//	 isValid := sciNotFmt.XIsValid() // isValid == true
//
//	 sciNotFmt = EngineeringNotationFormat(-999)
//
//	 isValid = sciNotFmt.XIsValid() // isValid == false
func (engNotationFmt EngineeringNotationFormat) XIsValid() bool {

	lockEngineeringNotationFormat.Lock()

	defer lockEngineeringNotationFormat.Unlock()

	return new(engineeringNotationFormatNanobot).
		isValidNumEngNotFmt(
			engNotationFmt)
}

//	XParseString
//
//	Receives a string and attempts to match it with the
//	string value of a supported enumeration. If
//	successful, a new instance of EngineeringNotationFormat
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
//		If 'true' the search for enumeration names will
//		be case-sensitive and will require an exact
//		match. Therefore, 'exponential' will NOT match
//		the enumeration name, 'Exponential'.
//
//		A case-sensitive search will match any of the
//		following strings:
//
//			"None"
//			"Exponential"
//			"ENotUprCaseELeadPlus"
//			"ENotUprCaseENoLeadPlus"
//			"ENotLwrCaseELeadPlus"
//			"ENotLwrCaseENoLeadPlus"
//			"SIPrefixSymbol"
//			"SIPrefixName"
//
//		If 'false', a case-insensitive search is conducted
//		for the enumeration name. In this example,
//		'exponential' WILL MATCH the enumeration name,
//		'Exponential'.
//
//		A case-insensitive search will match any of the
//		following lower case names:
//
//			"none"
//			"exponential"
//			"enotuprcaseeleadplus"
//			"enotuprcaseenoleadplus"
//			"enotlwrcaseeleadplus"
//			"enotlwrcaseenoleadplus"
//			"siprefixsymbol"
//			"siprefixname"
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	EngineeringNotationFormat
//
//		Upon successful completion, this method will
//		return a new instance of EngineeringNotationFormat
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
//	t, err := EngineeringNotationFormat(0).
//	             XParseString("ENotation", true)
//
//	t is now equal to EngineeringNotationFormat(0).ENotation()
func (engNotationFmt EngineeringNotationFormat) XParseString(
	valueString string,
	caseSensitive bool) (EngineeringNotationFormat, error) {

	lockEngineeringNotationFormat.Lock()

	defer lockEngineeringNotationFormat.Unlock()

	ePrefix := "EngineeringNotationFormat.XParseString() "

	var ok bool
	var sciNotFmt EngineeringNotationFormat

	if caseSensitive {

		sciNotFmt, ok =
			mEngNotationFmtStringToCode[valueString]

		if !ok {
			return EngineeringNotationFormat(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid Engineering Notation Format value.\n"+
					"valueString='%v'\n", valueString)
		}

	} else {

		sciNotFmt, ok =
			mEngNotationFmtLwrCaseStringToCode[strings.ToLower(valueString)]

		if !ok {
			return EngineeringNotationFormat(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid Engineering Notation Format value.\n"+
					"valueString='%v'\n", valueString)
		}
	}

	return sciNotFmt, nil
}

//	XReturnNoneIfInvalid
//
//	Provides a standardized value for invalid instances
//	of enumeration EngineeringNotationFormat.
//
//	If the current instance of EngineeringNotationFormat
//	is invalid, this method will always return a value
//	of EngineeringNotationFormat(0).None().
//
// ----------------------------------------------------------------
//
// # Background
//
//	Enumeration EngineeringNotationFormat has an
//	underlying type of integer (int). This means the type
//	could conceivably be set to any integer value. This
//	method ensures that all invalid EngineeringNotationFormat
//	instances are consistently classified as 'None'
//	(EngineeringNotationFormat(0).None()). Remember that
//	'None' is considered an INVALID selection for
//	'EngineeringNotationFormat'.
//
//	For example, assume that EngineeringNotationFormat was
//	set to an integer value of -848972. Calling this
//	method on a EngineeringNotationFormat with this
//	invalid integer value will return an integer value of
//	zero or the equivalent of EngineeringNotationFormat(0).None().
//
//	This conversion is useful in generating text strings for
//	meaningful informational and error messages.
//
//	This is a standard utility method and is not part of the
//	valid enumerations for this type.
func (engNotationFmt EngineeringNotationFormat) XReturnNoneIfInvalid() EngineeringNotationFormat {

	lockEngineeringNotationFormat.Lock()

	defer lockEngineeringNotationFormat.Unlock()

	isValid := new(engineeringNotationFormatNanobot).
		isValidNumEngNotFmt(engNotationFmt)

	if !isValid {
		return EngineeringNotationFormat(0)
	}

	return engNotationFmt
}

// XValue
//
// This method returns the enumeration value of the
// current EngineeringNotationFormat instance.
//
// This is a standard utility method and is NOT part of
// the valid enumerations for this type.
func (engNotationFmt EngineeringNotationFormat) XValue() EngineeringNotationFormat {

	lockEngineeringNotationFormat.Lock()

	defer lockEngineeringNotationFormat.Unlock()

	return engNotationFmt
}

// XValueInt
//
// This method returns the integer value of the current
// EngineeringNotationFormat instance.
//
// This is a standard utility method and is NOT part of
// the valid enumerations for this type.
func (engNotationFmt EngineeringNotationFormat) XValueInt() int {

	lockEngineeringNotationFormat.Lock()

	defer lockEngineeringNotationFormat.Unlock()

	return int(engNotationFmt)
}

//	EngNotFmt
//
//	Public global constant of type EngineeringNotationFormat.
//
//	This variable serves as an easier, shorthand technique
//	for accessing EngineeringNotationFormat values.
//
//	For easy access to these enumeration values, use the
//	global	variable EngNotFmt.
//
//		Example:
//
//			EngNotFmt.Exponential()
//
//	Otherwise you will need to use the formal syntax.
//
//	Example:
//
//		EngineeringNotationFormat(0).Exponential()
//
// ----------------------------------------------------------------
//
// # Usage
//
//	EngNotFmt.None()
//	EngNotFmt.Exponential()
//	EngNotFmt.ENotUprCaseELeadPlus()
//	EngNotFmt.ENotUprCaseELeadPlus()
//	EngNotFmt.ENotUprCaseENoLeadPlus()
//	EngNotFmt.ENotLwrCaseELeadPlus()
//	EngNotFmt.ENotLwrCaseENoLeadPlus()
const EngNotFmt = EngineeringNotationFormat(0)

// engineeringNotationFormatNanobot
//
// Provides helper methods for enumeration
// EngineeringNotationFormat.
type engineeringNotationFormatNanobot struct {
	lock *sync.Mutex
}

// isValidNumEngNotFmt
//
// Receives an instance of EngineeringNotationFormat and
// returns a boolean value signaling whether that
// EngineeringNotationFormat instance is valid.
//
// If the passed instance of EngineeringNotationFormat is
// valid, this method returns 'true'.
//
// Be advised, the enumeration value "None" is considered
// an INVALID selection for 'EngineeringNotationFormat'.
//
// This is a standard utility method and is not part of
// the valid EngineeringNotationFormat enumeration.
func (engNotFmtNanobot *engineeringNotationFormatNanobot) isValidNumEngNotFmt(
	sciNotFmt EngineeringNotationFormat) bool {

	if engNotFmtNanobot.lock == nil {
		engNotFmtNanobot.lock = new(sync.Mutex)
	}

	engNotFmtNanobot.lock.Lock()

	defer engNotFmtNanobot.lock.Unlock()

	if sciNotFmt < 1 ||
		sciNotFmt > 7 {

		return false
	}

	return true
}
