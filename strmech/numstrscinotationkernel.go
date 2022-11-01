package strmech

import (
	"math/big"
	"sync"
)

//	NumStrSciNotationKernel
//
//	Type NumStrSciNotationKernel contains the scientific
//	notation value used to format number strings
//	displaying numeric values as scientific notation.
//
// ----------------------------------------------------------------
//
//	# Definition of Terms
//
//	In scientific notation, nonzero numbers are written
//	in the form	'm × 10n' or m times ten raised to the
//	power of n, where n is an integer, and the coefficient
//	m is a nonzero real number (usually between 1 and 10
//	in absolute value, and nearly always written as a
//	terminating decimal).
//
//										Wikipedia
//
//
//	This type only contains the scientific or engineering
//	notation values. It DOES NOT contain format specifications.
//
// ----------------------------------------------------------------
//
//	Exponential Format
//
//		Example:
//			Base Numeric Value: 265,200,000
//			Exponential Display Format: "2.652 x 10^8"
//
//	E-Notation Format
//
//	 	Example:
//			Base Numeric Value: 265,200,000
//	 		E-Notation Display Format: "2.652e+8"
//
//	significand 				=	'2.652'
//	significand integer digits 	= 	'2'
//	mantissa					= 	significand factional digits = '.652'
//	exponent    				= 	'8'  (10^8)
//	mantissaLength				=	length of fractional digits displayed
//									in scientific notation.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://en.wikipedia.org/wiki/Scientific_notation
//	https://encyclopedia2.thefreedictionary.com/E+notation
//	https://www.medcalc.org/manual/scientific-notation.php
//	https://researchtweet.com/scientific-notation-definition-calculation/
//	https://www.wikihow.com/Multiply-Scientific-Notation
//	https://en.wikipedia.org/wiki/Engineering_notation
//	https://www.engineeringtoolbox.com/standard-form-scientific-engineering-notation-d_1801.html
//
//	# Exponential Notation
//
//	https://calculatort.com/convert-scientific-notation-into-standard-form/
//
//	# E-Notation
//
//	https://www.gigacalculator.com/calculators/scientific-notation-calculator.php
//	https://www.medcalc.org/manual/scientific-notation.php
//	https://calculatort.com/convert-scientific-notation-into-standard-form/
//
//	# Engineering Notation
//
//	https://en.wikipedia.org/wiki/Engineering_notation
//	https://sciencing.com/convert-between-base-number-systems-8442032.html
//	https://mathworld.wolfram.com/EngineeringNotation.html
//	https://www.youtube.com/watch?v=WfnTO_Pr3HE
//
// ----------------------------------------------------------------
//
// # Scientific Notation:
//
//  1. Scientific Notation has only 1-digit to the
//     left of the decimal point.
//
// ----------------------------------------------------------------
//
// # Engineering Notation:
//
//  1. Numbers multiplied by 1 get no modifier. This
//     is a base number.
//
//  2. Multipliers
//
//     k kilo	1 x 10^3	1,000
//     M mega	1 x 10^6	1,000,000
//     G giga	1 x 10^9	1,000,000,000
//     T tera	1 x 10^12	1,000,000,000,000
//     P peta	1 x 10^15	1,000,000,000,000,000
//     E exa	1 x 10^18	1,000,000,000,000,000,000
//     Z zetta	1 x 10^21	1,000,000,000,000,000,000,000
//     Y yotta	1 x 10^24	1,000,000,000,000,000,000,000,000
//     Base	1 x 10^0	1
//     m milli	1 x 10^-3	0.001
//     μ micro	1 x 10^-6	0.000 001
//     n nano	1 x 10^-9	0.000 000 001
//     p pico	1 x 10^-12	0.000 000 000 001
//     f femto	1 x 10^-15	0.000 000 000 000 001
//     a atto	1 x 10^-18  0.000 000 000 000 000 001
//     y yocto	1 x 10^-24	0.000 000 000 000 000 000 001
type NumStrSciNotationKernel struct {
	significand NumberStrKernel
	//	The significand consists of the leading integer and
	//	fractional digits of the scientific notation.
	//
	// In the example '2.652E+8', the significand is '2.652'.

	exponent *big.Int
	// The exponent portion of the scientific notation string
	// In the example '2.652E+8', the exponent is '8'

	lock *sync.Mutex
}

// NumStrENotationFormatSpec
//
// Contains all the specification parameters required
// to format a numeric value in scientific notation using
// the E-Notation Format.
//
// ----------------------------------------------------------------
//
//	# Definition of Terms
//
// ----------------------------------------------------------------
//
//	There are two types of formats used to display numeric
//	values in scientific notation: The Exponential Format
//	and the E-Notation Format.
//
//		Exponential Format Example
//
//			"2.652 x 10^8"
//
//		E-Notation Format Example
//
//	 		"2.652e+8"
//
//	Type NumStrENotationFormatSpec contains the
//	specification parameter necessary to format a
//	scientific notation or engineering notation value in
//	E-Notation Format.
type NumStrENotationFormatSpec struct {
	decimalSeparator DecimalSeparatorSpec
	//	The radix point or decimal separator used to separate
	//	integer and	fractional digits in the significand. The
	//	default is the standard USA decimal separator, the
	//	decimal point ('.').
	//
	//	In the example '2.652E+8', the decimal separator is
	//	the period character or decimal point ('.').

	significandUsesLeadingPlus bool
	//	"Significand uses leading plus sign". This refers to the
	//	integer digit in a significand.
	//
	//	Positive significand integer digits may have a leading
	//	plus sign, '+2.652E+8'. The default is no leading plus
	//	sign, '2.652E+8'.
	//
	//	If this value is set to true, positive significand integer
	//	digit values will be prefixed with a leading plus sign
	//	('+').
	//	        Example: '+2.652E+8'
	//
	//	If this value is set to true, positive significand integer
	//	digit values will NOT be prefixed with a leading plus sign
	//	('+').
	//	        Example: '2.652E+8'

	exponentUsesLeadingPlus bool
	//	If true, positive exponent values are
	//	prefixed with a leading plus (+) sign.
	//	Example: '2.652e+8'.
	//
	//	If this boolean value is set to 'false',
	//	it signals that no leading plus will be
	//	prefixed to positive exponent values. In
	//	this case the positive value is implied.
	//	Example: '2.652e8'.

	exponentCharIsUpperCase bool
	//	When set to 'true', the Exponent label 'E'
	//	is set to upper case. Example: '2.652E8'
	//
	//	When set to 'false', the Exponent label 'e'
	//	is set to lower case. Example: '2.652e8'
	//
	// 	The default is upper case 'E' in order
	//	to avoid confusion with the abbreviation
	//	of Euler's Number (e) or the exponential
	//	function  f(x) = e^x.

	lock *sync.Mutex
}
