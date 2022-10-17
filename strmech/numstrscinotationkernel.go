package strmech

import (
	"math/big"
	"sync"
)

//	NumStrSciNotationKernel
//
//	This structure contains the data elements necessary
//	to store and format a numeric value using scientific
//	notation.
//
// ----------------------------------------------------------------
//
//	Scientific Notation Terminology
//
// ----------------------------------------------------------------
//
//	 	Example: 2.652e+8
//
//	Definition of Terms
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
//	https://sciencing.com/convert-between-base-number-systems-8442032.html
//	https://mathworld.wolfram.com/EngineeringNotation.html
//
// ----------------------------------------------------------------
//
// # Scientific Notation:
//
//	1.	Scientific Notation has only 1-digit to the
//		left of the decimal point.

type NumStrSciNotationKernel struct {
	significand *big.Int
	//	The significand consists of the leading integer and
	//	fractional digits of the scientific notation.
	//
	// In the example '2.652E+8', the significand is '2.652'
	// and will be stored as a *big.Int number, '2652'.

	exponent *big.Int
	// The exponent portion of the scientific notation string
	// In the example '2.652E+8', the exponent is '8'

	mantissaLength uint
	//	The length of the fractional digits in the
	//	significand.
	//
	//	In the example '2.652E+8', the mantissa length is
	//	'3'.

	lock *sync.Mutex
}

// NumStrSciNotationFormatSpec
//
// Contains all the specification parameters required
// to format a numeric value in scientific notation for
// presentation as a number string.
type NumStrSciNotationFormatSpec struct {
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
	// 	The default is lower case 'e'.

	lock *sync.Mutex
}
