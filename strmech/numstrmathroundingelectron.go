package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type numStrMathRoundingElectron struct {
	lock *sync.Mutex
}

// roundDown
//
// Rounds a value down with truncation.
func (nStrMathRoundElectron *numStrMathRoundingElectron) roundDown(
	integerDigits *RuneArrayDto,
	fractionalDigits *RuneArrayDto,
	roundToFractionalDigits int,
	numberSign NumericSignValueType,
	errPrefDto *ePref.ErrPrefixDto) error {

	if nStrMathRoundElectron.lock == nil {
		nStrMathRoundElectron.lock = new(sync.Mutex)
	}

	nStrMathRoundElectron.lock.Lock()

	defer nStrMathRoundElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"nStrMathRoundAtom."+
			"roundDown()",
		"")

	if err != nil {
		return err
	}

	err = new(numStrMathRoundingQuark).preRoundingValidation(
		integerDigits,
		fractionalDigits,
		roundToFractionalDigits,
		numberSign,
		ePrefix)

	if err != nil {
		return err
	}

	// Truncate Fractional Digits array!
	// Last Index = roundToFractionalDigits - 1

	fractionalDigits.CharsArray =
		fractionalDigits.CharsArray[:roundToFractionalDigits]

	return err
}

// roundUp
//
// Rounds a value up by one.
func (nStrMathRoundElectron *numStrMathRoundingElectron) roundUp(
	integerDigits *RuneArrayDto,
	fractionalDigits *RuneArrayDto,
	roundToFractionalDigits int,
	numberSign NumericSignValueType,
	errPrefDto *ePref.ErrPrefixDto) error {

	if nStrMathRoundElectron.lock == nil {
		nStrMathRoundElectron.lock = new(sync.Mutex)
	}

	nStrMathRoundElectron.lock.Lock()

	defer nStrMathRoundElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"nStrMathRoundAtom."+
			"roundUp()",
		"")

	if err != nil {
		return err
	}

	err = new(numStrMathRoundingQuark).preRoundingValidation(
		integerDigits,
		fractionalDigits,
		roundToFractionalDigits,
		numberSign,
		ePrefix)

	if err != nil {
		return err
	}

	// Truncate Fractional Digits array!
	// Last Index = roundToFractionalDigits - 1

	fractionalDigits.CharsArray =
		fractionalDigits.CharsArray[:roundToFractionalDigits]

	lenFracDigits := fractionalDigits.GetRuneArrayLength()

	var isCarry bool

	if lenFracDigits > 0 {

		isCarry,
			err = new(numStrMathAtom).addOneToRunes(
			fractionalDigits,
			fractionalDigits,
			false,
			ePrefix.XCpy(
				"fractionalDigits<-total"))

		if err != nil {
			return err
		}

	} else {

		isCarry = true
	}

	if isCarry {

		_,
			err = new(numStrMathAtom).addOneToRunes(
			integerDigits,
			integerDigits,
			true,
			ePrefix.XCpy(
				"fractionalDigits<-total"))

	}

	return err
}
