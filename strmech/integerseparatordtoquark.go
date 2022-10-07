package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type integerSeparatorDtoQuark struct {
	lock *sync.Mutex
}

// empty - Deletes and resets the data values for all member
// variables within a IntegerSeparatorDto instance to their
// 'zero' values.
func (nStrIntSepQuark *integerSeparatorDtoQuark) empty(
	nStrIntSep *IntegerSeparatorDto,
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
		"integerSeparatorDtoQuark.empty()",
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

	nStrIntSep.intGroupingSequence = nil

	nStrIntSep.restartIntGroupingSequence = false

	return err
}

// testValidityOfNumStrIntSeparator - Tests the validity of a
// IntegerSeparatorDto instance.
//
// If the IntegerSeparatorDto instance is judged invalid, this
// method will return an error and set 'isValid' to false.
func (nStrIntSepQuark *integerSeparatorDtoQuark) testValidityOfNumStrIntSeparator(
	nStrIntSep *IntegerSeparatorDto,
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
		"integerSeparatorDtoQuark."+
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

	isValid,
		err = strMechPreon{}.ptr().
		testValidityOfRuneCharArray(
			nStrIntSep.intSeparatorChars,
			ePrefix)

	if err != nil {
		return isValid, err
	}

	arrayLen := len(nStrIntSep.intGroupingSequence)

	if arrayLen == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: 'intGroupingSequence' is invalid!\n"+
			"'intGroupingSequence' is a ZERO length rune array.\n",
			ePrefix.String())

		return isValid, err
	}

	for i := 0; i < arrayLen; i++ {

		if nStrIntSep.intGroupingSequence[i] == 0 {
			err = fmt.Errorf("%v\n"+
				"Error: 'intGroupingSequence' is invalid!\n"+
				"'intGroupingSequence' element number '%v' has zero value.\n",
				ePrefix.String(),
				i)

			return isValid, err
		}

		if nStrIntSep.intGroupingSequence[i] > 1000000 {
			err = fmt.Errorf("%v\n"+
				"Error: 'intGroupingSequence' is invalid!\n"+
				"'intGroupingSequence' element number '%v' has value\n"+
				"greater than 1,000,000!\n",
				ePrefix.String(),
				i)

			return isValid, err
		}

	}

	isValid = true

	return isValid, err
}
