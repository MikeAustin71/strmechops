package strmech

import (
	"sync"
)

//	SciNotationKernel
//
//	Type SciNotationKernel contains a numeric value in
//	scientific notation. This type can therefore be used
//	in arithmetic calculations and the creation of number
//	strings for the display numeric values as scientific
//	notation.
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
//	This type only contains the scientific notation values.
//	It DOES NOT contain engineering notation values.
//
//	For Engineering Notation reference type,
//	EngNotationKernel.
//
//	Example
//
//		Numeric Value				=	265,200,000
//		Scientific Notation Value	=	'2.652 x 10^8'
//		significand 				=	'2.652'
//		significand integer digits 	= 	'2'
//	    significand fractional digits =	'652'
//		exponent    				= 	'8'  (10^8)
//
// ----------------------------------------------------------------
//
// # Scientific Notation Display Formats
//
//	Scientific Notation may be formatted using one of
//	of two display formats.
//
//	1.	Exponential Format
//
//		Example:
//			Numeric Value: 265,200,000
//
//			Exponential Display Format: "2.652 x 10^8"
//
//	2.	E-Notation Format
//
//	 	Example:
//			Numeric Value: 265,200,000
//
//			E-Notation Display Format has two
//			variants:
//	 			E-Notation Display Format: "2.652E+8"
//							or
//				E-Notation Display Format: "2.652e+8"
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
	significand NumberStrKernel
	//	The significand consists of the leading integer and
	//	fractional digits of the scientific notation.
	//
	// In the example '2.652E+8', the significand is '2.652'.

	exponent NumberStrKernel
	// The exponent portion of the scientific notation string
	// In the example '2.652E+8', the exponent is '8'

	lock *sync.Mutex
}

//	GetENotationFmt
//
//	Converts the numeric value configured for the current
//	instance of SciNotationKernel to a Number String
//	configured with the Scientific Notation E-Notation
//	Format.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://en.wikipedia.org/wiki/Scientific_notation
//	https://encyclopedia2.thefreedictionary.com/E+notation
//
// ----------------------------------------------------------------
//
// # Usage
//
//	Example-1
//		 Numeric Value: 265,200,000
//		Display Format: "2.652E+8"
//
//	Example-2
//		 Numeric Value: 0.0002652
//		Display Format: "2.652E-4"
//
//	Example-3
//		 Numeric Value: 265,200,000
//		Display Format: "2.652E8"
//
//	Example-4
//		 Numeric Value: 265,200,000
//		Display Format: "2.652e+8"
//
//	Example-5
//		 Numeric Value: -265,200,000
//		Display Format: "-2.652e8"
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	positiveExponentFmt			string
//
//		This string contains the exponent format to
//		be applied when the exponent is a positive
//		numeric value. In the example "2.652E+8",
//		the positive exponent format is "E+"
//
//		If this parameter is submitted as an empty
//		string, it will be defaulted to, "E+"
//
//		Recommended options for this string value
//		are listed as follows:
//			"E+"
//			"E"
//			"e+"
//			"e"
//
//	negativeExponentFmt			string
//
//		This string contains the exponent format to
//		be applied when the exponent is a negative
//		numeric value. In the example "2.652E-4",
//		the negative exponent format is "E-"
//
//		If this parameter is submitted as an empty
//		string, it will be defaulted to, "E-"
//
//		Recommended options for this string value
//		are listed as follows:
//			"E-"
//			"e-"
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	string
//
//		This returned string will contain the Scientific
//		Notation numeric value configured with the
//		E-Notation display format.
func (sciNotKernel *SciNotationKernel) GetENotationFmt(
	positiveExponentFmt string,
	negativeExponentFmt string) string {

	if sciNotKernel.lock == nil {
		sciNotKernel.lock = new(sync.Mutex)
	}

	sciNotKernel.lock.Lock()

	defer sciNotKernel.lock.Unlock()

	if positiveExponentFmt == "" {
		positiveExponentFmt = "E+"
	}

	if negativeExponentFmt == "" {
		negativeExponentFmt = "E-"
	}

	sigNifIntStr := sciNotKernel.significand.GetIntegerString()

	sigNifFracStr := sciNotKernel.significand.GetFractionalString()

	sigNifNumSign,
		_ := sciNotKernel.significand.GetNumberSign(
		nil)

	var significandStr string

	if sigNifNumSign == NumSignVal.Negative() {

		significandStr += "-"

	}

	significandStr += sigNifIntStr

	if len(sigNifFracStr) == 0 {

		significandStr += ".0"

	} else {

		significandStr += "." + sigNifFracStr
	}

	exponentIntStr := sciNotKernel.exponent.GetIntegerString()

	exponentFracStr := sciNotKernel.exponent.GetFractionalString()

	exponentNumSign,
		_ := sciNotKernel.exponent.GetNumberSign(
		nil)

	var exponentStr = ""

	if exponentNumSign == NumSignVal.Negative() {

		exponentStr += negativeExponentFmt

	} else {

		exponentStr += positiveExponentFmt
	}

	exponentStr += exponentIntStr

	if len(exponentFracStr) > 0 {
		exponentStr += "." + exponentFracStr
	}

	return significandStr + exponentStr
}

//	GetNumStrExponentFmt
//
//	Converts the numeric value configured for the current
//	instance of SciNotationKernel to a Number String
//	configured with the Scientific Notation Exponential
//	Format.
//
//		Example:
//			Numeric Value: 265,200,000
//
//			Exponential Display Format: "2.652 x 10^8"
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://en.wikipedia.org/wiki/Scientific_notation
//
// ----------------------------------------------------------------
//
// # Usage
//
//	Example-1
//
//		Numeric Value: 265,200,000
//
//		Exponential Display Format: "2.652 x 10^8"
//
//	Example-2
//
//		Numeric Value: -265,200,000
//
//		Exponential Display Format: "-2.652 x 10^8"
//
//	Example-3
//
//		Numeric Value: 0.0002652
//
//		Exponential Display Format: "2.652 x 10^-4"
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	string
//
//		This returned string will contain the Scientific
//		Notation numeric value configured with the
//		Exponential display format.
func (sciNotKernel *SciNotationKernel) GetNumStrExponentFmt() string {

	if sciNotKernel.lock == nil {
		sciNotKernel.lock = new(sync.Mutex)
	}

	sciNotKernel.lock.Lock()

	defer sciNotKernel.lock.Unlock()

	sigNifIntStr := sciNotKernel.significand.GetIntegerString()

	sigNifFracStr := sciNotKernel.significand.GetFractionalString()

	sigNifNumSign,
		_ := sciNotKernel.significand.GetNumberSign(
		nil)

	var significandStr string

	if sigNifNumSign == NumSignVal.Negative() {
		significandStr += "-"
	}

	significandStr += sigNifIntStr

	if len(sigNifFracStr) == 0 {

		significandStr += ".0"

	} else {

		significandStr += "." + sigNifFracStr
	}

	exponentIntStr := sciNotKernel.exponent.GetIntegerString()

	exponentFracStr := sciNotKernel.exponent.GetFractionalString()

	exponentNumSign,
		_ := sciNotKernel.exponent.GetNumberSign(
		nil)

	exponentStr := " x 10^"

	if exponentNumSign == NumSignVal.Negative() {
		exponentStr += "-"
	}

	exponentStr += exponentIntStr

	if len(exponentFracStr) > 0 {
		exponentStr += "." + exponentFracStr
	}

	return significandStr + exponentStr
}
