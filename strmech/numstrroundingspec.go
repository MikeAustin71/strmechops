package strmech

type NumStrRoundingSpec struct {
	RoundFractionalDigits bool
	// When set to 'true' this signals that rounding should
	// be performed on fractional numbers. When set to 'false'
	// this parameter signals that no rounding should be
	// performed.

	RoundingType NumberRoundingType
	// This enumeration parameter is used to specify the type
	// of rounding algorithm that will be applied in the
	// rounding operation.
	//
	// Possible values are listed as follows:
	//	NumRoundType.None(),
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

	RoundToFractionalDigits int
	// When set to a positive integer value, this
	// parameter controls the number of digits to
	// the right of the decimal separator (a.k.a.
	// decimal point) which will remain after
	// completion of the number rounding operation.

}
