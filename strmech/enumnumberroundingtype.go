package strmech

import (
	"fmt"
	"strings"
	"sync"
)

// Lock lockNumberRoundingType before accessing these maps!

var mNumberRoundingTypeCodeToString = map[NumberRoundingType]string{
	NumberRoundingType(0):  "None",
	NumberRoundingType(1):  "NoRounding",
	NumberRoundingType(2):  "HalfUpWithNegNums",
	NumberRoundingType(3):  "HalfDownWithNegNums",
	NumberRoundingType(4):  "HalfAwayFromZero",
	NumberRoundingType(5):  "HalfTowardsZero",
	NumberRoundingType(6):  "HalfToEven",
	NumberRoundingType(7):  "HalfToOdd",
	NumberRoundingType(8):  "Randomly",
	NumberRoundingType(9):  "Floor",
	NumberRoundingType(10): "Ceiling",
	NumberRoundingType(11): "Truncate",
}

var mNumberRoundingTypeStringToCode = map[string]NumberRoundingType{
	"None":                NumberRoundingType(0),
	"NoRounding":          NumberRoundingType(1),
	"HalfUpWithNegNums":   NumberRoundingType(2),
	"HalfDownWithNegNums": NumberRoundingType(3),
	"HalfAwayFromZero":    NumberRoundingType(4),
	"HalfTowardsZero":     NumberRoundingType(5),
	"HalfToEven":          NumberRoundingType(6),
	"HalfToOdd":           NumberRoundingType(7),
	"Randomly":            NumberRoundingType(8),
	"Floor":               NumberRoundingType(9),
	"Ceiling":             NumberRoundingType(10),
	"Truncate":            NumberRoundingType(11),
}

var mNumberRoundingTypeLwrCaseStringToCode = map[string]NumberRoundingType{
	"none":                NumberRoundingType(0),
	"norounding":          NumberRoundingType(1),
	"halfupwithnegnums":   NumberRoundingType(2),
	"halfdownwithnegnums": NumberRoundingType(3),
	"halfawayfromzero":    NumberRoundingType(4),
	"halftowardszero":     NumberRoundingType(5),
	"halftoeven":          NumberRoundingType(6),
	"halftoodd":           NumberRoundingType(7),
	"randomly":            NumberRoundingType(8),
	"floor":               NumberRoundingType(9),
	"ceiling":             NumberRoundingType(10),
	"truncate":            NumberRoundingType(11),
}

// NumberRoundingType - The 'Number Rounding Type' is an
// enumeration of type codes used for classification of
// numeric rounding methodologies and algorithms.
//
// The examples provided here relate to floating point rounding
// which is the most common application for this go package.
//
// ----------------------------------------------------------------
//
// # TERMINOLOGY
//
// Rounding means replacing a number with an approximate value that
// has a shorter, simpler, or more explicit representation. For
// example, replacing $23.4476 with $23.45, the fraction 312/937
// with 1/3, or the expression √2 with 1.414.
//
//	Wikipedia
//	https://en.wikipedia.org/wiki/Rounding
//
// Most rounding methodologies are fairly straight forward and
// understandable. The differences and questions surrounding which
// methodology to apply usually relate to handling a value of
// one-half or '.5' in the 'Round From Digit' position.
//
//	Round To Digit   - The digit we are rounding to.
//	                   Example: 23.14567
//
//	                   The 'Round To' Digit in this example is: '4'
//
//	Round From Digit - The digit to the immediate right of the
//	                   'Round To' Digit. The 'Round From' Digit
//	                   will determine what rounding procedure will
//	                   be applied to the 'Round To' Digit.
//	                   Example: 23.14567
//
//	                   The 'Round From' Digit in this example is: '5'
//
//	Example: 23.14567
//	                   Objective: Round to two decimal places to
//	                              the right of the decimal point.
//	                   Rounding Method: HalfAwayFromZero
//	                   Round To Digit:   4
//	                   Round From Digit: 5
//	                   Rounded Number:   23.15
//
//	Reference:
//	       https://www.mathsisfun.com/numbers/rounding-methods.html
//
//
//	Sources describe the 'HalfUpWithNegNums' and
//	'HalfDownWithNegNums' as the most common or intuitive types of
//	rounding methods. However, both of the methods have some issues
//	or unexpected outcomes when applied to negative numbers.
//
//	      HalfUpWithNegNums           HalfDownWithNegNums
//
//	      7.6 rounds up to 8          7.6 rounds up to 8
//	      7.5 rounds up to 8          7.5 rounds down to 7
//	      7.4 rounds down to 7        7.4 rounds down to 7
//	      -7.4 rounds up to -7        -7.4 rounds up to -7
//	      -7.5 rounds up to -7        -7.5 rounds down to -8
//	      -7.6 rounds down to -8      -7.6 rounds down to -8
//
//	Rounding methods 'HalfAwayFromZero' and 'HalfTowardsZero' may
//	provide more clarity and fewer surprises when dealing with
//	negative numbers.
//
//	    HalfAwayFromZero              HalfTowardsZero
//
//	    7.6 rounds away to 8          7.6 rounds away to 8
//	    7.5 rounds away to 8          7.5 rounds to 7
//	    7.4 rounds to 7               7.4 rounds to 7
//	    -7.4 rounds to -7             -7.4 rounds to -7
//	    -7.5 rounds away to -8        -7.5 rounds to -7
//	    -7.6 rounds away to -8        -7.6 rounds away to -8
//
//	 This utility method will return a default rounding method
//	 of 'HalfAwayFromZero':
//	   NumberRoundingType(0).XGetDefaultRoundingType()
//
// ----------------------------------------------------------------
//
// # BACKGROUND
//
// Type NumberRoundingType is styled as an enumeration. Since the
// Go Programming Language does not directly support enumerations,
// type NumberRoundingType has been adapted to function in a manner
// similar to classic enumerations.
//
// NumberRoundingType is declared as a type 'int' and includes two
// types of methods:
//
//	Enumeration Methods
//	      and
//	Utility Methods
//
// Enumeration methods have names which collectively represent an
// enumeration of different rounding methodologies and procedures
// which may be applied to numeric rounding operations.
//
//	  Examples Of Enumeration Method Names:
//	      HalfUpWithNegNums()
//	      HalfDownWithNegNums()
//	      HalfAwayFromZero()
//	      HalfTowardsZero()
//	      HalfToEven()
//	      HalfToOdd()
//
//	Enumeration methods return an integer value used to designate
//	a specific rounding methodology.
//
//	Utility methods make up the second type of method included in
//	NumberRoundingType. These methods are NOT part of the
//	enumeration but instead provide needed supporting services. All
//	utility methods, with the sole exception of method String(),
//	have names beginning with 'X' to separate them from standard
//	enumeration methods.
//	  Examples:
//	    XIsValid()
//	    XParseString()
//	    XValue()
//	    XValueInt()
//
//	The utility method 'String()' supports the Stringer Interface
//	and is not part of the standard enumeration.
//
// ----------------------------------------------------------------
//
// # Enumeration Methods
//
// The NumberRoundingType enumeration methods are described
// below:
//
// Method                   Integer
//
//	Name                     Value
//
// ------                   -------
//
// None                     Zero (0)
//
//	Signals that the Number Rounding Type
//	(NumberRoundingType) Type is empty and
//	uninitialized. This is an error condition.
//
// NoRounding					1
//
//	Signals that no rounding operation will be performed
//	on fractional digits contained in a number string.
//	The fractional digits will therefore remain unchanged.
//
// HalfUpWithNegNums			2
//
//	Half Round Up Including Negative Numbers. This method
//	is intuitive but may produce unexpected results when
//	applied to negative numbers.
//
//	'HalfUpWithNegNums' rounds .5 up.
//
//		Examples of 'HalfUpWithNegNums'
//		7.6 rounds up to 8
//		7.5 rounds up to 8
//		7.4 rounds down to 7
//		-7.4 rounds up to -7
//		-7.5 rounds up to -7
//		-7.6 rounds down to -8
//
// HalfDownWithNegNums          3
//
//	Half Round Down Including Negative Numbers. This method
//	is also considered intuitive but may produce unexpected
//	results when applied to negative numbers.
//
//	'HalfDownWithNegNums' rounds .5 down.
//
//		Examples of HalfDownWithNegNums
//
//		7.6 rounds up to 8
//		7.5 rounds down to 7
//		7.4 rounds down to 7
//		-7.4 rounds up to -7
//		-7.5 rounds down to -8
//		-7.6 rounds down to -8
//
// HalfAwayFromZero				4
//
//	Round Half Away From Zero. This rounding method is treated
//	as the default and this value is returned by method:
//	NumberRoundingType(0).XGetDefaultRoundingType()
//
//	The 'HalfAwayFromZero' method rounds .5 further away from zero.
//	It provides clear and consistent behavior when dealing with
//	negative numbers.
//
//		Examples of HalfAwayFromZero
//
//		7.6 rounds away to 8
//		7.5 rounds away to 8
//		7.4 rounds to 7
//		-7.4 rounds to -7
//		-7.5 rounds away to -8
//		-7.6 rounds away to -8
//
// HalfTowardsZero				5
//
//	Round Half Towards Zero. 'HalfTowardsZero' rounds 0.5
//	closer to zero. It provides clear and consistent behavior
//	when dealing with negative numbers.
//
//		Examples of HalfTowardsZero
//
//		7.6 rounds away to 8
//		7.5 rounds to 7
//		7.4 rounds to 7
//		-7.4 rounds to -7
//		-7.5 rounds to -7
//		-7.6 rounds away to -8
//
// HalfToEven					6
//
//	Round Half To Even Numbers. 'HalfToEven' is also called
//	Banker's Rounding. This method rounds 0.5 to the nearest
//	even digit.
//
//		Examples of HalfToEven
//
//		7.5 rounds up to 8 (because 8 is an even number)
//		but 6.5 rounds down to 6 (because 6 is an even number)
//
//		HalfToEven only applies to 0.5. Other numbers (not ending
//		in 0.5) round to nearest as usual, so:
//
//		7.6 rounds up to 8
//		7.5 rounds up to 8 (because 8 is an even number)
//		7.4 rounds down to 7
//		6.6 rounds up to 7
//		6.5 rounds down to 6 (because 6 is an even number)
//		6.4 rounds down to 6
//
// HalfToOdd					7
//
//	Round Half to Odd Numbers. Similar to 'HalfToEven', but
//	in this case 'HalfToOdd' rounds 0.5 towards odd numbers.
//
//		Examples of HalfToOdd
//
//		HalfToOdd only applies to 0.5. Other numbers (not ending
//		in 0.5) round to nearest as usual.
//
//		7.5 rounds down to 7 (because 7 is an odd number)
//
//		6.5 rounds up to 7 (because 7 is an odd number)
//
//		7.6 rounds up to 8
//		7.5 rounds down to 7 (because 7 is an odd number)
//		7.4 rounds down to 7
//		6.6 rounds up to 7
//		6.5 rounds up to 7 (because 7 is an odd number)
//		6.4 rounds down to 6
//
// Randomly						8
//
//	Round Half Randomly. Uses a Random Number Generator to choose
//	between rounding 0.5 up or down.
//
//	All numbers other than 0.5 round to the nearest as usual.
//
// Floor						9
//
//	Yields the nearest integer down. Floor does not apply any
//	special treatment to 0.5.
//
//	Floor Function: The greatest integer that is less than or
//	equal to x
//	Source: https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//	In mathematics and computer science, the floor function is
//	the function that takes as input a real number x, and gives
//	as output the greatest integer less than or equal to x,
//	denoted floor(x) or ⌊x⌋.
//	Source: https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//		Examples of Floor
//
//		Number     Floor
//		2           2
//		2.4         2
//		2.9         2
//		-2.5        -3
//		-2.7        -3
//		-2          -2
//
// Ceiling						10
//
//	Yields the nearest integer up. Ceiling does not apply any
//	special treatment to 0.5.
//
//	Ceiling Function: The least integer that is greater than or
//	equal to x.
//	Source: https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//	The ceiling function maps x to the least integer greater than
//	or equal to x, denoted ceil(x) or ⌈x⌉.[1]
//	Source: https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//		Examples of Ceiling
//
//		Number    Ceiling
//		2           2
//		2.4         3
//		2.9         3
//		-2.5        -2
//		-2.7        -2
//		-2          -2
//
// Truncate						11
//
//	Apply NO Rounding whatsoever. The Round From Digit is dropped
//	or deleted. The Round To Digit is NEVER changed.
//
//		Examples of Truncate
//
//		Example-1
//		Number: 23.14567
//		Objective: Round to two decimal places to
//		the right of the decimal point.
//		Rounding Method: Truncate
//		Round To Digit:   4
//		Round From Digit: 5
//		Rounded Number:   23.14 - The Round From Digit is dropped.
//
//		Example-2
//		Number: -23.14567
//		Objective: Round to two decimal places to
//		the right of the decimal point.
//		Rounding Method: Truncate
//		Round To Digit:   4
//		Round From Digit: 5
//		Rounded Number:  -23.14 - The Round From Digit is dropped.
//
// ----------------------------------------------------------------
//
// # REFERENCE
//
//	https://www.mathsisfun.com/numbers/rounding-methods.html
//	https://en.wikipedia.org/wiki/Rounding
//	https://www.mathsisfun.com/rounding-numbers.html
//	https://www.vedantu.com/maths/rounding-methods
//	https://rounding.to/the-most-common-rounding-methods/
//	https://www.wikihow.com/Round-Numbers
//
// ----------------------------------------------------------------
//
// # USAGE
//
// For easy access to these enumeration values, use the global
// constant 'NumRoundType'.
//
//	Example: NumRoundType.HalfAwayFromZero()
//
// Otherwise you will need to use the formal syntax.
//
//	Example: NumberRoundingType(0).HalfAwayFromZero()
//
// Depending on your editor, intellisense (a.k.a. intelligent code
// completion) may not list the NumberRoundingType methods in
// alphabetical order.
//
// Be advised that all 'NumberRoundingType' methods beginning with
// 'X', as well as the method 'String()', are utility methods and
// not part of the enumeration.
type NumberRoundingType int

var lockNumberRoundingType sync.Mutex

// None
//
// Signals that the Number Round Type (NumberRoundingType)
// Type is empty and uninitialized.
//
// This is an error condition.
//
// This method is part of the standard enumeration.
func (numRoundingType NumberRoundingType) None() NumberRoundingType {

	lockNumberRoundingType.Lock()

	defer lockNumberRoundingType.Unlock()

	return NumberRoundingType(0)
}

// NoRounding
//
// Signals that no rounding operation will be performed
// on fractional digits contained in a number string.
// The fractional digits will therefore remain unchanged.
//
// This method is part of the standard enumeration.
func (numRoundingType NumberRoundingType) NoRounding() NumberRoundingType {

	lockNumberRoundingType.Lock()

	defer lockNumberRoundingType.Unlock()

	return NumberRoundingType(1)
}

// HalfUpWithNegNums
//
// Half Round Up Including Negative Numbers.
// This method is intuitive but may produce unexpected results
// when applied to negative numbers.
//
// 'HalfUpWithNegNums' rounds 0.5 up.
//
//	Examples of 'HalfUpWithNegNums'
//
//	7.6 rounds up to 8
//	7.5 rounds up to 8
//	7.4 rounds down to 7
//	-7.4 rounds up to -7
//	-7.5 rounds up to -7
//	-7.6 rounds down to -8
//
// This method is part of the standard enumeration.
func (numRoundingType NumberRoundingType) HalfUpWithNegNums() NumberRoundingType {

	lockNumberRoundingType.Lock()

	defer lockNumberRoundingType.Unlock()

	return NumberRoundingType(2)
}

// HalfDownWithNegNums
//
// Half Round Down Including Negative Numbers. This method
// is considered intuitive but may produce unexpected
// results when applied to negative numbers.
//
// 'HalfDownWithNegNums' rounds 0.5 down.
//
//	Examples of HalfDownWithNegNums
//
//	7.6 rounds up to 8
//	7.5 rounds down to 7
//	7.4 rounds down to 7
//	-7.4 rounds up to -7
//	-7.5 rounds down to -8
//	-7.6 rounds down to -8
//
// This method is part of the standard enumeration.
func (numRoundingType NumberRoundingType) HalfDownWithNegNums() NumberRoundingType {

	lockNumberRoundingType.Lock()

	defer lockNumberRoundingType.Unlock()

	return NumberRoundingType(3)
}

// HalfAwayFromZero
//
// Round Half Away From Zero. This rounding method is
// treated as the default and this value is returned by
// method:
//
//	NumberRoundingType(0).XGetDefaultRoundingType()
//
// 'HalfAwayFromZero' rounds 0.5 further away from zero. It
// provides clear and consistent behavior when dealing with
// negative numbers.
//
//	Examples of HalfAwayFromZero
//
//	7.6 rounds away to 8
//	7.5 rounds away to 8
//	7.4 rounds to 7
//	-7.4 rounds to -7
//	-7.5 rounds away to -8
//	-7.6 rounds away to -8
//
// This method is part of the standard enumeration.
func (numRoundingType NumberRoundingType) HalfAwayFromZero() NumberRoundingType {

	lockNumberRoundingType.Lock()

	defer lockNumberRoundingType.Unlock()

	return NumberRoundingType(4)
}

// HalfTowardsZero
//
// Round Half Towards Zero.  'HalfTowardsZero' rounds 0.5 closer
// to zero. It provides clear and consistent behavior when dealing
// with negative numbers.
//
//	Examples of HalfTowardsZero
//
//	 7.6 rounds away to 8
//	 7.5 rounds to 7
//	 7.4 rounds to 7
//	 -7.4 rounds to -7
//	 -7.5 rounds to -7
//	 -7.6 rounds away to -8
//
// This method is part of the standard enumeration.
func (numRoundingType NumberRoundingType) HalfTowardsZero() NumberRoundingType {

	lockNumberRoundingType.Lock()

	defer lockNumberRoundingType.Unlock()

	return NumberRoundingType(5)
}

// HalfToEven
//
// Round to Even (Banker's Rounding). Round 0.5 to the
// nearest even digit.
//
//		Examples of HalfToEven
//
//		7.5 rounds up to 8 (because 8 is an even number)
//		but 6.5 rounds down to 6 (because 6 is an even number)
//
//		HalfToEven only applies to 0.5. Other numbers (not ending
//		in 0.5) round to nearest as usual, so:
//
//		7.6 rounds up to 8
//		7.5 rounds up to 8 (because 8 is an even number)
//		7.4 rounds down to 7
//		6.6 rounds up to 7
//		6.5 rounds down to 6 (because 7 is an odd number)
//		6.4 rounds down to 6
//		-23.5 rounds to -24 (because 4 is an even number)
//	 -24.5 rounds to -24 (because 5 is an odd number)
//
// Reference:
//
//	https://en.wikipedia.org/wiki/Rounding#Round_half_to_even
//	https://rounding.to/understanding-the-bankers-rounding/
//	https://www.mathsisfun.com/numbers/rounding-methods.html
//	https://en.wikipedia.org/wiki/Rounding
//	https://www.mathsisfun.com/rounding-numbers.html
//	https://www.vedantu.com/maths/rounding-methods
//	https://rounding.to/the-most-common-rounding-methods/
//	https://www.wikihow.com/Round-Numbers
//
// This method is part of the standard enumeration.
func (numRoundingType NumberRoundingType) HalfToEven() NumberRoundingType {

	lockNumberRoundingType.Lock()

	defer lockNumberRoundingType.Unlock()

	return NumberRoundingType(6)
}

// HalfToOdd
//
// Round Half to Odd Numbers. Similar to 'HalfToEven',
// but in this case 'HalfToOdd' rounds 0.5 towards odd numbers.
//
//	Examples of HalfToOdd
//
//	HalfToOdd only applies to 0.5. Other numbers
//	(not ending in 0.5) round to nearest as usual.
//
//	7.5 rounds down to 7 (because 8 is an even number)
//
//	6.5 rounds up to 7 (because 7 is an odd number)
//
//	7.6 rounds up to 8
//	7.5 rounds down to 7 (because 8 is an even number)
//	7.4 rounds down to 7
//	6.6 rounds up to 7
//	6.5 rounds up to 7 (because 7 is an odd number)
//	6.4 rounds down to 6
//
// This method is part of the standard enumeration.
func (numRoundingType NumberRoundingType) HalfToOdd() NumberRoundingType {

	lockNumberRoundingType.Lock()

	defer lockNumberRoundingType.Unlock()

	return NumberRoundingType(7)
}

// Randomly
//
// Round Half Randomly. Uses a Random Number Generator
// to choose between rounding 0.5 up or down.
//
// All numbers other than 0.5 round to the nearest as usual.
//
// This method is part of the standard enumeration.
func (numRoundingType NumberRoundingType) Randomly() NumberRoundingType {

	lockNumberRoundingType.Lock()

	defer lockNumberRoundingType.Unlock()

	return NumberRoundingType(8)
}

// Floor
//
// Yields the nearest integer down. Floor does not apply
// any special treatment to 0.5.
//
//	Floor Function: The greatest integer that is less than or
//	                equal to x
//
//	Source: https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
// In mathematics and computer science, the floor function is the
// function that takes as input a real number x, and gives as
// output the greatest integer less than or equal to x, denoted
// floor(x) or ⌊x⌋.
//
//	Source: https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//		Examples of Floor
//
//		  Number     Floor
//		   2           2
//		   2.4         2
//		   2.9         2
//		  -2.5        -3
//		  -2.7        -3
//		  -2          -2
//
// This method is part of the standard enumeration.
func (numRoundingType NumberRoundingType) Floor() NumberRoundingType {

	lockNumberRoundingType.Lock()

	defer lockNumberRoundingType.Unlock()

	return NumberRoundingType(9)
}

// Ceiling
//
// Yields the nearest integer up. Ceiling does not apply
// any special treatment to 0.5.
//
//	Ceiling Function: The least integer that is greater than or
//	                  equal to x.
//
//	Source: https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//	The ceiling function maps x to the least integer greater than
//	or equal to x, denoted ceil(x) or ⌈x⌉.[1]
//
//	Source: https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//		Examples of Ceiling
//
//		  Number    Ceiling
//		   2           2
//		   2.4         3
//		   2.9         3
//		  -2.5        -2
//		  -2.7        -2
//		  -2          -2
//
// This method is part of the standard enumeration.
func (numRoundingType NumberRoundingType) Ceiling() NumberRoundingType {

	lockNumberRoundingType.Lock()

	defer lockNumberRoundingType.Unlock()

	return NumberRoundingType(10)
}

// Truncate
//
// Apply NO Rounding whatsoever. The Round From Digit is
// dropped or deleted. The Round To Digit is NEVER changed.
//
//	Examples of Truncate
//
//	 Example-1
//	   Number: 23.14567
//	   Objective: Round to two decimal places to
//	              the right of the decimal point.
//	   Rounding Method: Truncate
//	   Round To Digit:   4
//	   Round From Digit: 5
//	   Rounded Number:   23.14 - The Round From Digit
//	   	is dropped.
//
//	 Example-2
//	   Number: -23.14567
//	   Objective: Round to two decimal places to
//	              the right of the decimal point.
//	   Rounding Method: Truncate
//	   Round To Digit:   4
//	   Round From Digit: 5
//	   Rounded Number:  -23.14 - The Round From Digit
//	   	is dropped.
//
// This method is part of the standard enumeration.
func (numRoundingType NumberRoundingType) Truncate() NumberRoundingType {

	lockNumberRoundingType.Lock()

	defer lockNumberRoundingType.Unlock()

	return NumberRoundingType(11)
}

// String - Returns a string with the name of the enumeration
// associated with this current instance of 'NumberRoundingType'.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ----------------------------------------------------------------
//
// # Usage
//
// t:= NumberRoundingType(0).Floor()
// str := t.String()
//
//	str is now equal to 'Floor'
func (numRoundingType NumberRoundingType) String() string {

	lockNumberRoundingType.Lock()

	defer lockNumberRoundingType.Unlock()

	result, ok := mNumberRoundingTypeCodeToString[numRoundingType]

	if !ok {

		return "Error: Number Rounding Type Specification UNKNOWN!"

	}

	return result
}

// XGetDefaultRoundingType - Returns a default rounding type of
// HalfAwayFromZero.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
func (numRoundingType NumberRoundingType) XGetDefaultRoundingType() NumberRoundingType {

	lockNumberRoundingType.Lock()

	defer lockNumberRoundingType.Unlock()

	// Default Rounding Type is HalfAwayFromZero
	return NumberRoundingType(3)
}

// XIsValid - Returns a boolean value signaling whether the current
// NumberRoundingType value is valid.
//
// Be advised, the enumeration value "None" is considered a VALID
// selection for 'NumberRoundingType'.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	 roundingType :=
//				NumberRoundingType(0).HalfAwayFromZero()
//
//	 isValid := roundingType.XIsValid() // isValid == true
//
//	 roundingType = NumberRoundingType(-999)
//
//	 isValid = roundingType.XIsValid() // isValid == false
func (numRoundingType NumberRoundingType) XIsValid() bool {

	lockNumberRoundingType.Lock()

	defer lockNumberRoundingType.Unlock()

	return new(numberRoundingTypeNanobot).
		isValidNumRoundType(
			numRoundingType)
}

// XParseString - Receives a string and attempts to match it with
// the string value of a supported enumeration. If successful, a
// new instance of NumberRoundingType is returned set to the value
// of the associated enumeration.
//
// This is a standard utility method and is NOT part of the valid
// enumerations for this type.
//
// ----------------------------------------------------------------
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
//     'halfawayfromzero' will NOT match the enumeration name,
//     'HalfAwayFromZero'.
//
//     A case-sensitive search will match any of the following
//     strings:
//     "None"
//     "HalfUpWithNegNums"
//     "HalfDownWithNegNums"
//     "HalfAwayFromZero"
//     "HalfTowardsZero"
//     "HalfToEven"
//     "HalfToOdd"
//     "Randomly"
//     "Floor"
//     "Ceiling"
//     "Truncate"
//
//     If 'false', a case-insensitive search is conducted for the
//     enumeration name. In this example, 'halfawayfromzero'
//     WILL MATCH the enumeration name, 'HalfAwayFromZero'.
//
//     A case-insensitive search will match any of the following
//     lower case names:
//     "none"
//     "halfupwithnegnums"
//     "halfdownwithnegnums"
//     "halfawayfromzero"
//     "halftowardszero"
//     "halftoeven"
//     "halftoodd"
//     "randomly"
//     "floor"
//     "ceiling"
//     "truncate"
//
// ----------------------------------------------------------------
//
// Return Values
//
//	NumberRoundingType
//	   - Upon successful completion, this method will return a new
//	     instance of NumberRoundingType set to the value of the
//	     enumeration matched by the string search performed on
//	     input parameter, 'valueString'.
//
//	error
//	   - If this method completes successfully, the returned error
//	     Type is set equal to 'nil'. If an error condition is
//	     encountered, this method will return an error type which
//	     encapsulates an appropriate error message.
//
// ----------------------------------------------------------------
//
// Usage
//
//	t, err := NumberRoundingType(0).
//	             XParseString("HalfAwayFromZero", true)
//
//	t is now equal to NumberRoundingType(0).HalfAwayFromZero()
func (numRoundingType NumberRoundingType) XParseString(
	valueString string,
	caseSensitive bool) (NumberRoundingType, error) {

	lockNumberRoundingType.Lock()

	defer lockNumberRoundingType.Unlock()

	ePrefix := "NumberRoundingType.XParseString() "

	var ok bool
	var numberRoundingType NumberRoundingType

	if caseSensitive {

		numberRoundingType, ok =
			mNumberRoundingTypeStringToCode[valueString]

		if !ok {
			return NumberRoundingType(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid NumberRoundingType Specification.\n"+
					"valueString='%v'\n", valueString)
		}

	} else {

		numberRoundingType, ok =
			mNumberRoundingTypeLwrCaseStringToCode[strings.ToLower(valueString)]

		if !ok {
			return NumberRoundingType(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid NumberRoundingType Specification.\n"+
					"valueString='%v'\n", valueString)
		}
	}

	return numberRoundingType, nil
}

// XReturnNoneIfInvalid - Provides a standardized value for invalid
// instances of enumeration NumberRoundingType.
//
// If the current instance of NumberRoundingType is invalid, this
// method will always return a value of
// NumberRoundingType(0).None().
//
// # Background
//
// Enumeration NumberRoundingType has an underlying type of
// integer (int). This means the type could conceivably be set
// to any integer value. This method ensures that all invalid
// NumberRoundingType instances are consistently classified as
// 'None' (NumberRoundingType(0).None()). Remember that 'None'
// is considered an INVALID selection for 'NumberRoundingType'.
//
// For example, assume that NumberRoundingType was set to an
// integer value of -848972. Calling this method on a
// NumberRoundingType with this invalid integer value will
// return an integer value of zero or the equivalent of
// NumberRoundingType(0).None(). This conversion is useful in
// generating text strings for meaningful informational and
// error messages.
//
// This is a standard utility method and is not part of the
// valid enumerations for this type.
func (numRoundingType NumberRoundingType) XReturnNoneIfInvalid() NumberRoundingType {

	lockNumberRoundingType.Lock()

	defer lockNumberRoundingType.Unlock()

	isValid := new(numberRoundingTypeNanobot).
		isValidNumRoundType(numRoundingType)

	if !isValid {
		return NumberRoundingType(0)
	}

	return numRoundingType
}

// XValue - This method returns the enumeration value of the
// current NumberRoundingType instance.
//
// This is a standard utility method and is NOT part of the valid
// enumerations for this type.
func (numRoundingType NumberRoundingType) XValue() NumberRoundingType {

	lockNumberRoundingType.Lock()

	defer lockNumberRoundingType.Unlock()

	return numRoundingType
}

// XValueInt - This method returns the integer value of the current
// NumberRoundingType instance.
//
// This is a standard utility method and is NOT part of the valid
// enumerations for this type.
func (numRoundingType NumberRoundingType) XValueInt() int {

	lockNumberRoundingType.Lock()

	defer lockNumberRoundingType.Unlock()

	return int(numRoundingType)
}

// NumRoundType - Public global constant of type
// NumberRoundingType.
//
// This variable serves as an easier, shorthand technique for
// accessing NumberRoundingType values.
//
// For easy access to these enumeration values, use the global
// variable NumRoundType.
//
//	Example: NumRoundType.HalfAwayFromZero()
//
// Otherwise you will need to use the formal syntax.
//
//	Example: NumberRoundingType(0).HalfAwayFromZero()
//
// Usage:
//
//	NumRoundType.None(),
//	NumRoundType.NoRounding(),
//	NumRoundType.HalfUpWithNegNums(),
//	NumRoundType.HalfDownWithNegNums(),
//	NumRoundType.HalfAwayFromZero(),
//	NumRoundType.HalfTowardsZero(),
//	NumRoundType.HalfToEven(),
//	NumRoundType.HalfToOdd(),
//	NumRoundType.Randomly(),
//	NumRoundType.Floor(),
//	NumRoundType.Ceiling(),
//	NumRoundType.Truncate(),
const NumRoundType = NumberRoundingType(0)

// numberRoundingTypeNanobot - Provides helper methods for
// enumeration NumberRoundingType.
type numberRoundingTypeNanobot struct {
	lock *sync.Mutex
}

// isValidNumRoundType - Receives an instance of
// NumberRoundingType and returns a boolean value signaling whether
// that NumberRoundingType instance is valid.
//
// If the passed instance of NumberRoundingType is valid, this
// method returns 'true'.
//
// Be advised, the enumeration value "None" is considered a
// VALID selection for 'NumberRoundingType'.
//
// This is a standard utility method and is not part of the valid
// NumberRoundingType enumeration.
func (numRoundTypeNanobot *numberRoundingTypeNanobot) isValidNumRoundType(
	numberRoundingType NumberRoundingType) bool {

	if numRoundTypeNanobot.lock == nil {
		numRoundTypeNanobot.lock = new(sync.Mutex)
	}

	numRoundTypeNanobot.lock.Lock()

	defer numRoundTypeNanobot.lock.Unlock()

	if numberRoundingType < 1 ||
		numberRoundingType > 11 {

		return false
	}

	return true
}
