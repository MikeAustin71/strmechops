package strmech

import "sync"

//	Lock lockScientificNotationCalcType before accessing
//	these maps!

var mSciNotationCalcTypeCodeToString = map[ScientificNotationCalcType]string{
	ScientificNotationCalcType(0): "None",
	ScientificNotationCalcType(1): "Standard",
	ScientificNotationCalcType(2): "Engineering",
}

var mSciNotationCalcTypeStringToCode = map[string]ScientificNotationCalcType{
	"None":        ScientificNotationCalcType(0),
	"Standard":    ScientificNotationCalcType(1),
	"Engineering": ScientificNotationCalcType(2),
}

var mSciNotationCalcTypeLwrCaseStringToCode = map[string]ScientificNotationCalcType{
	"none":        ScientificNotationCalcType(0),
	"standard":    ScientificNotationCalcType(1),
	"engineering": ScientificNotationCalcType(2),
}

// ScientificNotationCalcType
//
//	The 'Scientific Notation Calculation Type' is an
//	enumeration of type codes used top specify the
//	type of calculation used to generate a numeric
//	value expressed in Scientific Notation.
//
// ----------------------------------------------------------------
//
// # Terminology
//
//	Scientific Notation is a way of expressing numbers
//	that are too large or too small (usually would result
//	in a long string of digits) to be conveniently
//	written in decimal form.
//
//		E-Notation Format Example
//
//			Base Numeric Value in Decimal Format:
//				265,200,000
//
//		Scientific Notation Value in E-Notation Format:
//				2.652e+8
//
//	Scientific Notation can be understood as having two
//	dimensions.
//
//	Display Format
//
//	First, there is manner in which a Scientific Notation
//	value is formatted for output or display. There are
//	two common formats for display of Scientific Notation
//	values:
//
//		1. Exponential Format
//
//			Exponential Example:
//				Base Numeric Value: 265,200,000
//				Expnential Display Format: "2.652 x 10^8"
//
//		2. E-Notation Format
//
//		 	E-Notation Example:
//				Base Numeric Value: 265,200,000
//		 		E-Notation Display Format: "2.652e+8"
//
//		For an enmeration of Scientific Notation Display
//		Formats see type 'ScientificNotationFormat'.
//
//	Standard Calculation Type
//
//	Second, there is the manner in which the Scientific
//	Notation value is calulated. Generally, there are
//	two different methods for calculating the Scientific
//	Notation value.
//
//		1. Standard Scientific Notation Calculation Type
//
//		Using the Normalized or Standard Calculation Type
//		for Scientific Notation, nonzero numbers are
//		written in the form:
//
//				m × 10^n
//
//		This means that m times ten raised to the power
//		of n, where n is an integer, and the coefficient
//		m is a nonzero real number (usually between 1 and
//		10 in absolute value, and nearly always written
//		as a terminating decimal).
//
//		The integer n is called the exponent and the real
//		number m is called the significand or mantissa.
//
//		If the base numeric value is negative then a minus
//		sign precedes m, as in ordinary decimal notation.
//
//			Base Numeric Value: -53,000
//			Standard Calculation Value: −5.3×10^4
//
//		In the normalized or standard calculation
//		methodology, the exponent is chosen so that the
//		absolute value (modulus) of the significand m is
//		at least 1 but less than 10.
//
//								Scientific Notation
//			Decimal Notation	Standard Calculation
//			----------------	--------------------
//
//				    2				2×10^0
//
//				  300				3×10^2
//
//			4,321.768				4.321768×10^3
//
//			  −53,000			   −5.3×10^4
//
//		6,720,000,000				6.72×10^9
//
//				  0.2				2×10^−1
//
//				  987				9.87×10^2
//
//		0.00000000751				7.51×10^−9
//
//		2. Engineering Calculation Type
//
//		Engineering notation or engineering form is a
//		version of scientific notation in which the
//		exponent of ten must be divisible by three
//		(i.e., they are powers of a thousand, but written
//		as, for example, 10^6 instead of 1000^2). As an
//		alternative to writing powers of 10, SI prefixes
//		(Metric prefixes) can be used, which also usually
//		provide steps of a factor of a thousand.
//
//		Engineering notation is a version of scientific
//		notation in which the exponent n in expressions
//		of the form m × 10^n is chosen to always be
//		divisible by 3.
//
//		Numbers of forms such as 12×10^-6, 230×10^-3,
//		340, and 4.5×10^3 therefore correspond to an
//		Engineering	Calcuation Type.
//
//		Numbers such as 12×10^-2, 2×10^2, and 123×10^5 DO
//		NOT correspond to the Engineering Calcuation Type.
//
//		Care must be taken when attempting to infer
//		significant digits from numbers expressed in
//		engineering notation. For example, while 3.0×10^5
//		unambiguously indicates two significant digits in
//		scientific notation, the number of significant
//		digits in 300×10^3 is not clear, with the only
//		obvious way to indicate significance being to
//		write the number in the awkward form 0.30×10^6.
//
//
//								Engineering Notation
//			Decimal Notation		 Calculation
//			----------------	--------------------
//
//				  2,700 			 2.7 × 10^3
//				 27,000 			27 × 10^3
//				270,000			   270 × 10^3
//			  2,700,000 			 2.7 × 10^6
//
// ----------------------------------------------------------------
//
// # Reference
//
//	https://en.wikipedia.org/wiki/Scientific_notation
//	https://sciencing.com/e-5491586.html
//	https://en.wikipedia.org/wiki/Engineering_notation
//	https://mathworld.wolfram.com/EngineeringNotation.html
//	https://www.mathsisfun.com/numbers/scientific-notation.html
//	https://www.mathsisfun.com/definitions/engineering-notation.html
//	https://sciencing.com/e-5491586.html
//	https://www.engineeringtoolbox.com/standard-form-scientific-engineering-notation-d_1801.html
//	https://openoregon.pressbooks.pub/techmath/chapter/module-11-scientific-notation/
//	https://www.rapidtables.com/convert/number/scientific-notation-converter.html
//
// ----------------------------------------------------------------
//
// # Enumeration Values
//
//	Since the Go Programming Language does not directly
//	support enumerations, the ScientificNotationCalcType
//	has been adapted to function in a manner similar to
//	classic enumerations.
//
//	ScientificNotationCalcType is declared as a type
//	'int'. The method names associated with this type
//	effectively represent an enumeration of Scientific
//	Notation Calculation Types. These methods are listed
//	as follows:
//
//	Method		Integer
//	 Name	 	 Value
//	------		-------
//
//	None	   	   0
//
//		Signals that 'ScientificNotationCalcType' has
//		not been initialized and therefore has no value.
//		This is an error condition.
//
//	Standard		1
//
//		Specifies a Standard Scientific Notaion
//		Calculation. A numeric value calulated in
//		Standard Scientific Notation is written in the
//		form:
//
//			m × 10^n
//
//			Where:
//
//				1.	m is a number greater than or equal
//					to one and less than 10.
//
//				2. n is an integer.
//
//	Engineering		2
//
//		Specifies an Engineering Notation Calculation.
//		A numeric value calculated in Engineering
//		Notation is written in the form:
//
//			m × 10^n
//
//			Where:
//
//				1.	m is a number greater than zero and less
//					than 999.
//
//				2.	n is an integer multiple of 3
//
// ----------------------------------------------------------------
//
// # Usage
//
//	For easy access to these enumeration values, use the
//	global constant SciNotCalcType.
//
//		Example:
//			SciNotCalcType.Standard()
//
//	Otherwise you will need to use the formal syntax.
//
//		Example:
//			ScientificNotationCalcType(0).Standard()
//
//	Depending on your editor, intellisense (a.k.a.
//	intelligent code completion) may not list the
//	ScientificNotationCalcType methods in alphabetical
//	order.
//
//	Be advised that all ScientificNotationCalcType
//	methods beginning with 'X', as well as the method
//	'String()', are utility methods, and NOT part of the
//	enumeration values.
type ScientificNotationCalcType int

var lockScientificNotationCalcType sync.Mutex

// None
//
// Signals that 'ScientificNotationCalcType' has not
// been initialized and therefore has no value.
//
// This is an error condition.
func (sciNotationCalcType ScientificNotationCalcType) None() ScientificNotationCalcType {

	lockScientificNotationCalcType.Lock()

	defer lockScientificNotationCalcType.Unlock()

	return ScientificNotationCalcType(0)
}

//	Standard
//
//	Specifies the Standard Scientific Notation
//	Calculation Type.
//
//	A numeric value calulated in Standard Scientific
//	Notation is written in the form:
//
//			m × 10^n
//
//			Where:
//
//				1.	m is a number greater than zero and
//					less than 10.
//
//				2.	n is an integer.
//
//	This method is part of the ScientificNotationCalcType
//	enumeration.
//
// ----------------------------------------------------------------
//
// # Reference
//
//	https://en.wikipedia.org/wiki/Scientific_notation
//	https://sciencing.com/e-5491586.html
//	https://en.wikipedia.org/wiki/Scientific_notation
//	https://www.mathsisfun.com/numbers/scientific-notation.html
//	https://www.rapidtables.com/convert/number/scientific-notation-converter.html
func (sciNotationCalcType ScientificNotationCalcType) Standard() ScientificNotationCalcType {

	lockScientificNotationCalcType.Lock()

	defer lockScientificNotationCalcType.Unlock()

	return ScientificNotationCalcType(1)
}

//	Engineering
//
//	Specifies an Engineering Notation Calculation. A
//	numeric value calculated in Engineering Notation is
//	written in the form:
//
//			m × 10^n
//
//			Where:
//
//				1.	m is a number greater than zero and less
//					than 999.
//
//				2.	n is an integer multiple of 3
//
//	This method is part of the ScientificNotationCalcType
//	enumeration.
//
// ----------------------------------------------------------------
//
// # Reference
//
// /	https://en.wikipedia.org/wiki/Engineering_notation
//
//	https://mathworld.wolfram.com/EngineeringNotation.html
//	https://www.mathsisfun.com/definitions/engineering-notation.html
//	https://sciencing.com/e-5491586.html
//	https://www.engineeringtoolbox.com/standard-form-scientific-engineering-notation-d_1801.html
//	https://openoregon.pressbooks.pub/techmath/chapter/module-11-scientific-notation/
//	https://www.rapidtables.com/convert/number/scientific-notation-converter.html
func (sciNotationCalcType ScientificNotationCalcType) Engineering() ScientificNotationCalcType {

	lockScientificNotationCalcType.Lock()

	defer lockScientificNotationCalcType.Unlock()

	return ScientificNotationCalcType(2)
}
