package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type integerSeparatorSpecQuark struct {
	lock *sync.Mutex
}

// empty - Deletes and resets the data values for all member
// variables within a IntegerSeparatorSpec instance to their
// 'zero' values.
func (nStrIntSepQuark *integerSeparatorSpecQuark) empty(
	nStrIntSep *IntegerSeparatorSpec,
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
		"integerSeparatorSpecQuark.empty()",
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

	nStrIntSep.intSeparatorGrouping = nil

	nStrIntSep.restartIntGroupingSequence = false

	nStrIntSep.turnOffIntegerSeparation = false

	return err
}

// testValidityOfNumStrIntSeparator - Tests the validity of a
// IntegerSeparatorSpec instance.
//
// If the IntegerSeparatorSpec instance is judged invalid, this
// method will return an error and set 'isValid' to false.
func (nStrIntSepQuark *integerSeparatorSpecQuark) testValidityOfNumStrIntSeparator(
	nStrIntSep *IntegerSeparatorSpec,
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
		"integerSeparatorSpecQuark."+
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

	if len(nStrIntSep.intSeparatorChars) == 0 {
		// This is a NOP Integer Separator
		// Integer separation is turned off.

		isValid = true

		return isValid, err
	}

	isValid,
		err = strMechPreon{}.ptr().
		testValidityOfRuneCharArray(
			nStrIntSep.intSeparatorChars,
			ePrefix)

	if err != nil {
		return isValid, err
	}

	arrayLen := len(nStrIntSep.intSeparatorGrouping)

	if arrayLen == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: 'intSeparatorGrouping' is invalid!\n"+
			"'intSeparatorGrouping' is a ZERO length rune array.\n",
			ePrefix.String())

		return isValid, err
	}

	for i := 0; i < arrayLen; i++ {

		if nStrIntSep.intSeparatorGrouping[i] == 0 {
			err = fmt.Errorf("%v\n"+
				"Error: 'intSeparatorGrouping' is invalid!\n"+
				"'intSeparatorGrouping' element number '%v' has zero value.\n",
				ePrefix.String(),
				i)

			return isValid, err
		}

		if nStrIntSep.intSeparatorGrouping[i] > 1000000 {
			err = fmt.Errorf("%v\n"+
				"Error: 'intSeparatorGrouping' is invalid!\n"+
				"'intSeparatorGrouping' element number '%v' has value\n"+
				"greater than 1,000,000!\n",
				ePrefix.String(),
				i)

			return isValid, err
		}

	}

	isValid = true

	return isValid, err
}
