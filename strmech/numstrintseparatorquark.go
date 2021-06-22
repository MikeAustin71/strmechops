package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type numStrIntSeparatorQuark struct {
	lock *sync.Mutex
}

// empty - Deletes and resets the data values for all member
// variables within a NumStrIntSeparator instance to their
// 'zero' values.
//
func (nStrIntSepQuark *numStrIntSeparatorQuark) empty(
	nStrIntSep *NumStrIntSeparator,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if nStrIntSepQuark.lock == nil {
		nStrIntSepQuark.lock = new(sync.Mutex)
	}

	nStrIntSepQuark.lock.Lock()

	defer nStrIntSepQuark.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrIntSeparatorQuark.empty()",
		"")

	if err != nil {
		return err
	}

	if nStrIntSep == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrIntSep' is invalid!\n"+
			"'nStrIntSep' is a nil pointer.",
			ePrefix.String())

		return err
	}

	nStrIntSep.intSeparatorChars = nil

	nStrIntSep.intSeparatorGrouping = 0

	nStrIntSep.intSeparatorRepetitions = 0

	nStrIntSep.restartIntGroupingSequence = false

	return err
}

// testValidityOfNumStrIntSeparator - Tests the validity of a
// NumStrIntSeparator instance.
//
// If the NumStrIntSeparator instance is judged invalid, this
// method will return an error and set 'isValid' to false.
//
func (nStrIntSepQuark *numStrIntSeparatorQuark) testValidityOfNumStrIntSeparator(
	nStrIntSep *NumStrIntSeparator,
	errPrefDto *ePref.ErrPrefixDto) (
	isValid bool,
	err error) {

	if nStrIntSepQuark.lock == nil {
		nStrIntSepQuark.lock = new(sync.Mutex)
	}

	nStrIntSepQuark.lock.Lock()

	defer nStrIntSepQuark.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrIntSeparatorQuark."+
			"testValidityOfNumStrIntSeparator()",
		"")

	if err != nil {
		return isValid, err
	}

	if nStrIntSep == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrIntSep' is invalid!\n"+
			"'nStrIntSep' is a nil pointer.",
			ePrefix.String())

		return isValid, err
	}

	lIntSepChars := len(nStrIntSep.intSeparatorChars)

	if lIntSepChars == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: 'intSeparatorChars' is invalid!\n"+
			"'intSeparatorChars' is a ZERO length rune array.\n",
			ePrefix.String())

		return isValid, err
	}

	if nStrIntSep.intSeparatorGrouping > 10000000 {

		err = fmt.Errorf("%v\n"+
			"Error: 'nStrIntSep.intSeparatorGrouping' is invalid!\n"+
			"nStrIntSep.intSeparatorGrouping='%v'\n",
			ePrefix,
			nStrIntSep.intSeparatorGrouping)

		return isValid, err
	}

	if nStrIntSep.intSeparatorRepetitions > 10000000 {

		err = fmt.Errorf("%v\n"+
			"Error: 'nStrIntSep.intSeparatorRepetitions' is invalid!\n"+
			"nStrIntSep.intSeparatorRepetitions='%v'\n",
			ePrefix,
			nStrIntSep.intSeparatorRepetitions)

		return isValid, err
	}

	isValid = true

	return isValid, err
}

// ptr - Returns a pointer to a new instance of
// numStrIntSeparatorQuark
func (nStrIntSepQuark numStrIntSeparatorQuark) ptr() *numStrIntSeparatorQuark {

	if nStrIntSepQuark.lock == nil {
		nStrIntSepQuark.lock = new(sync.Mutex)
	}

	nStrIntSepQuark.lock.Lock()

	defer nStrIntSepQuark.lock.Unlock()

	newIntSepQuark := new(numStrIntSeparatorQuark)

	newIntSepQuark.lock = new(sync.Mutex)

	return newIntSepQuark
}
