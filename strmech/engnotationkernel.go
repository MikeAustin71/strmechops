package strmech

import (
	"math/big"
	"sync"
)

//	EngNotationKernel
//
//	Type EngNotationKernel contains the Engineering
//	Notation value used to format number strings
//	displaying numeric values as Engineering Notation.
//
//	The type only contains the scientific or engineering
//	notation values, NOT the format specifications.
//
// ----------------------------------------------------------------
//
//	# Definition of Terms
//
// ----------------------------------------------------------------
//
//	in scientific notation.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	# Engineering Notation
//
//	https://en.wikipedia.org/wiki/Engineering_notation
//	https://sciencing.com/convert-between-base-number-systems-8442032.html
//	https://mathworld.wolfram.com/EngineeringNotation.html
//	https://www.youtube.com/watch?v=WfnTO_Pr3HE
type EngNotationKernel struct {
	coefficient *big.Int
	//	The coefficient consists of the leading integer and
	//	fractional digits of the engineering notation.
	//
	//	The coefficient is also referred to as the
	//	significand.
	//
	// In the example '2.652E+8', the coefficient is '2.652'.

	exponent *big.Int
	// The exponent portion of the Engineering Notation
	// number string presentation.
	//
	// In the example '2.652E+8', the exponent is '8'

	lock *sync.Mutex
}
