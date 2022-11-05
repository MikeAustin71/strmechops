package strmech

import (
	"math/big"
	"sync"
)

//	SciNotationKernel
//
//	Type SciNotationKernel contains the scientific
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
//			E-Notation Display Format has two
//			variants:
//	 			E-Notation Display Format: "2.652E+8"
//							or
//				E-Notation Display Format: "2.652e+8"
//
//	significand 				=	'2.652'
//	significand integer digits 	= 	'2'
//	exponent    				= 	'8'  (10^8)
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
//	https://calculatort.com/convert-scientific-notation-into-standard-form/
//
//	# E-Notation
//
//	https://www.gigacalculator.com/calculators/scientific-notation-calculator.php
//	https://www.medcalc.org/manual/scientific-notation.php
//	https://calculatort.com/convert-scientific-notation-into-standard-form/
//
// ----------------------------------------------------------------
//
// # Scientific Notation:
//
//  1. Scientific Notation has only 1-integer digit to the
//     left of the decimal point.
//
//  2. The value of the single integer digit positioned to
//     the left of the decimal point must be greater than
//     or equal the one (1) and less than or equal to nine
//     (9).
type SciNotationKernel struct {
	significand *big.Int
	//	The significand consists of the leading integer and
	//	fractional digits of the scientific notation.
	//
	// In the example '2.652E+8', the significand is '2.652'.

	exponent *big.Int
	// The exponent portion of the scientific notation string
	// In the example '2.652E+8', the exponent is '8'

	lock *sync.Mutex
}
