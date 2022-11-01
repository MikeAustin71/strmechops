package strmech

import (
	"sync"
)

// BigDecimal
//
// Contains a representing of integer or floating
// point numeric values.
//
// This type incorporates integer math routines to
// ensure a high degree of accuracy when managing
// floating point calculations. The type is also
// capable of processing very large or very small
// numeric values.
//
// The encapsulated numeric value is maintained
// internally as form of scientific notation.
type BigDecimal struct {
	significand NumberStrKernel
	//	The significand is also known as the mantissa or
	//	coefficient in scientific notation. It represents
	//	the significant digits in the base number shown
	//	in the following example:
	//
	//		Scientific Notation Calculation:
	//			265,200,000 = 2.652 x 10^8
	//		Base Numeric Value: 265,200,000
	//		Scientific Notation Format: "2.652 x 10^8"
	//		significand: '2.652'
	//		exponent: 8

	exponent NumberStrKernel
	//	The component of scientific notation known as the
	//	'exponent' is illustrated in the following example:
	//
	//		Scientific Notation Calculation:
	//			265,200,000 = 2.652 x 10^8
	//		Base Numeric Value: 265,200,000
	//		Scientific Notation Format: "2.652 x 10^8"
	//		significand: '2.652'
	//		exponent: 8

	lock *sync.Mutex
}
