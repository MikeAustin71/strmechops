package strmech

import "sync"

type NumberBuilder struct {
	numericValueType NumericValueType     // Integer or Floating Point
	numberFormatType NumStrFormatTypeCode // Absolute Value, Binary, Currency,
	//                                     // Hexadecimal, Octal, SignedNumber,
	//                                     // Scientific Notation
	numberStrFormat string // A string containing format specifications for
	//                       //  displaying the numeric value encapsulated by
	//                       //  this NumberBuilder instance.
	integerElement    []rune // The integer component of the number
	fractionalElement []rune // The fractional component of the number
	exponentElement   []rune // The exponent component of scientific notation
	decimalSeparator  []rune // The decimal separator separating integer and
	//                       //  fractional elements.
	currencySymbol   []rune              // The currency symbol
	numberSignSymbol []rune              // The number sign symbol
	integerSeparator IntegerSeparatorDto // Integer separator format information.
	//                                    //  This is usually the thousands separator.
	numberFieldLen int // The length of the text number field in which the
	//                       //  numeric value will be displayed
	textJustification TextJustify // The justification specification which will
	//                            //  position the numeric value in the text
	//                            //  number field: Left, Right or Center
	lock *sync.Mutex
}
