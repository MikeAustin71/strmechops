package strmech

import "sync"

type NumStrRoundingSpec struct {
	roundFractionalDigits bool
	// When set to 'true' this signals that rounding should
	// be performed on fractional numbers. When set to 'false'
	// this parameter signals that no rounding should be
	// performed.

	roundingType NumberRoundingType
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

	roundToFractionalDigits int
	// When set to a positive integer value, this
	// parameter controls the number of digits to
	// the right of the decimal separator (a.k.a.
	// decimal point) which will remain after
	// completion of the number rounding operation.

	lock *sync.Mutex
}

func (nStrRoundingSpec *NumStrRoundingSpec) Empty() {

	if nStrRoundingSpec.lock == nil {
		nStrRoundingSpec.lock = new(sync.Mutex)
	}

	nStrRoundingSpec.lock.Lock()

	defer nStrRoundingSpec.lock.Unlock()

	nStrRoundingSpec.roundFractionalDigits = false

	nStrRoundingSpec.roundingType =
		NumRoundType.None()

	nStrRoundingSpec.roundToFractionalDigits = -1

	return
}

func (nStrRoundingSpec *NumStrRoundingSpec) Equal(
	incomingNStrRoundingSpec *NumStrRoundingSpec) bool {

	if nStrRoundingSpec.lock == nil {
		nStrRoundingSpec.lock = new(sync.Mutex)
	}

	nStrRoundingSpec.lock.Lock()

	defer nStrRoundingSpec.lock.Unlock()

	if nStrRoundingSpec.roundFractionalDigits !=
		incomingNStrRoundingSpec.roundFractionalDigits {

		return false
	}

	if nStrRoundingSpec.roundingType !=
		incomingNStrRoundingSpec.roundingType {

		return false
	}

	if nStrRoundingSpec.roundToFractionalDigits !=
		incomingNStrRoundingSpec.roundToFractionalDigits {

		return false
	}

	return true
}
