package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// NegativeNumberSignSpec - Negative Number Sign Specification
//
type NegativeNumberSignSpec struct {
	negNumSignPosition           NumSignSymbolPosition // Before(), After(), BeforeAndAfter()
	leadingNegNumSignSymbols     []rune
	foundLeadingNegNumSign       bool
	foundLeadingNegNumSignIndex  int
	trailingNegNumSignSymbols    []rune
	foundTrailingNegNumSign      bool
	foundTrailingNegNumSignIndex int
	lock                         *sync.Mutex
}

func (negNumSignSpec *NegativeNumberSignSpec) IsValidInstanceError(
	errorPrefix interface{}) (
	err error) {

	if negNumSignSpec.lock == nil {
		negNumSignSpec.lock = new(sync.Mutex)
	}

	negNumSignSpec.lock.Lock()

	defer negNumSignSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSignSpec.IsValidInstanceError()",
		"")

	if err != nil {
		return err
	}

	if !negNumSignSpec.negNumSignPosition.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of 'NegativeNumberSignSpec' is invalid!\n"+
			"The internal member variable 'negNumSignPosition' is NOT configured.\n",
			ePrefix.String())

		return err
	}

	sMechPreon := strMechPreon{}

	var err2 error

	if negNumSignSpec.negNumSignPosition == NSignSymPos.Before() {

		if len(negNumSignSpec.trailingNegNumSignSymbols) > 0 {

			err = fmt.Errorf("%v\n"+
				"Error: This instance of 'NegativeNumberSignSpec' is invalid!\n"+
				"It is configured as a Leading Negative Number Sign.\n"+
				"However, it contains Trailing Negative Number Sign characters.\n",
				ePrefix.String())

			return err

		}

		if len(negNumSignSpec.leadingNegNumSignSymbols) == 0 {

			err = fmt.Errorf("%v\n"+
				"Error: This instance of 'NegativeNumberSignSpec' is invalid!\n"+
				"It is configured as a Leading Negative Number Sign.\n"+
				"However, no Leading Negative Number Sign characters are configured.\n",
				ePrefix.String())

			return err
		}

		_,
			err2 = sMechPreon.testValidityOfRuneCharArray(
			negNumSignSpec.leadingNegNumSignSymbols,
			nil)

		if err2 != nil {
			err = fmt.Errorf("%v\n"+
				"Error: This instance of 'NegativeNumberSignSpec' is invalid!\n"+
				"It is configured as a Leading Negative Number Sign Symbol.\n"+
				"Internal member variable 'leadingNegNumSignSymbols' returned\n"+
				"the following validation error:\n"+
				"%v\n",
				ePrefix.String(),
				err2.Error())

			return err
		}

		return err
	}

	if negNumSignSpec.negNumSignPosition == NSignSymPos.After() {

		if len(negNumSignSpec.leadingNegNumSignSymbols) > 0 {

			err = fmt.Errorf("%v\n"+
				"Error: This instance of 'NegativeNumberSignSpec' is invalid!\n"+
				"It is configured as a Trailing Negative Number Sign.\n"+
				"However, it contains Leading Negative Number Sign characters.\n",
				ePrefix.String())

			return err

		}

		if len(negNumSignSpec.trailingNegNumSignSymbols) == 0 {

			err = fmt.Errorf("%v\n"+
				"Error: This instance of 'NegativeNumberSignSpec' is invalid!\n"+
				"It is configured as a Trailing Negative Number Sign.\n"+
				"However, no Trailing Negative Number Sign characters are configured.\n",
				ePrefix.String())

			return err

		}

		_,
			err2 = sMechPreon.testValidityOfRuneCharArray(
			negNumSignSpec.trailingNegNumSignSymbols,
			nil)

		if err2 != nil {
			err = fmt.Errorf("%v\n"+
				"Error: This instance of 'NegativeNumberSignSpec' is invalid!\n"+
				"It is configured as a Trailing Negative Number Sign Symbol.\n"+
				"Internal member variable 'trailingNegNumSignSymbols' returned\n"+
				"the following validation error:\n"+
				"%v\n",
				ePrefix.String(),
				err2.Error())

			return err
		}

		return err
	}

	if negNumSignSpec.negNumSignPosition == NSignSymPos.BeforeAndAfter() {

		if len(negNumSignSpec.leadingNegNumSignSymbols) == 0 {

			err = fmt.Errorf("%v\n"+
				"Error: This instance of 'NegativeNumberSignSpec' is invalid!\n"+
				"It is configured as a Leading and Trailing Negative Number Sign.\n"+
				"However, it contains NO Leading Negative Number Sign characters.\n",
				ePrefix.String())

			return err

		}

		_,
			err2 = sMechPreon.testValidityOfRuneCharArray(
			negNumSignSpec.leadingNegNumSignSymbols,
			nil)

		if err2 != nil {
			err = fmt.Errorf("%v\n"+
				"Error: This instance of 'NegativeNumberSignSpec' is invalid!\n"+
				"It is configured as a Leading and Trailing Negative Number Sign Symbol.\n"+
				"Internal member variable 'leadingNegNumSignSymbols' returned\n"+
				"the following validation error:\n"+
				"%v\n",
				ePrefix.String(),
				err2.Error())

			return err
		}

		if len(negNumSignSpec.trailingNegNumSignSymbols) == 0 {

			err = fmt.Errorf("%v\n"+
				"Error: This instance of 'NegativeNumberSignSpec' is invalid!\n"+
				"It is configured as a Trailing Negative Number Sign.\n"+
				"However, it contains NO Trailing Negative Number Sign characters.\n",
				ePrefix.String())

			return err

		}

		_,
			err2 = sMechPreon.testValidityOfRuneCharArray(
			negNumSignSpec.trailingNegNumSignSymbols,
			nil)

		if err2 != nil {
			err = fmt.Errorf("%v\n"+
				"Error: This instance of 'NegativeNumberSignSpec' is invalid!\n"+
				"It is configured as a Leading and Trailing Negative Number Sign Symbol.\n"+
				"Internal member variable 'trailingNegNumSignSymbols' returned\n"+
				"the following validation error:\n"+
				"%v\n",
				ePrefix.String(),
				err2.Error())

			return err
		}

		return err
	}

	return err
}

// NewLeadingNegNumSign - Returns a fully populated specification
// for a Leading Negative Number Sign.
func (negNumSignSpec NegativeNumberSignSpec) NewLeadingNegNumSign(
	leadingNegNumSignSymbols []rune,
	errorPrefix interface{}) (
	leadingNegNumSignSpec NegativeNumberSignSpec,
	err error) {

	if negNumSignSpec.lock == nil {
		negNumSignSpec.lock = new(sync.Mutex)
	}

	negNumSignSpec.lock.Lock()

	defer negNumSignSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSignSpec.NewLeadingNegNumSign()",
		"")

	if err != nil {
		return leadingNegNumSignSpec, err
	}

	if len(leadingNegNumSignSymbols) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'leadingNegNumSignSymbols' is invalid!\n"+
			"'leadingNegNumSignSymbols' is an empty array. The array length\n"+
			"is zero (0)!\n",
			ePrefix.String())

		return leadingNegNumSignSpec, err
	}

	sMechPreon := strMechPreon{}

	var err2 error

	_,
		err2 = sMechPreon.testValidityOfRuneCharArray(
		leadingNegNumSignSymbols,
		nil)

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'leadingNegNumSignSymbols' is invalid!\n"+
			"'leadingNegNumSignSymbols' returned the following validation error:\n"+
			"%v\n",
			ePrefix.String(),
			err2.Error())

		return leadingNegNumSignSpec, err
	}

	err = sMechPreon.copyRuneArrays(
		&leadingNegNumSignSpec.leadingNegNumSignSymbols,
		&leadingNegNumSignSymbols,
		true,
		ePrefix.XCpy(
			"leadingNegNumSignSpec<-leadingNegNumSignSymbols"))

	if err != nil {
		return leadingNegNumSignSpec, err
	}

	leadingNegNumSignSpec.negNumSignPosition = NSignSymPos.Before()

	return leadingNegNumSignSpec, err
}

// NewTrailingNegNumSign - Returns a fully populated specification
// for a Trailing Negative Number Sign.
func (negNumSignSpec NegativeNumberSignSpec) NewTrailingNegNumSign(
	trailingNegNumSignSymbols []rune,
	errorPrefix interface{}) (
	trailingNegNumSignSpec NegativeNumberSignSpec,
	err error) {

	if negNumSignSpec.lock == nil {
		negNumSignSpec.lock = new(sync.Mutex)
	}

	negNumSignSpec.lock.Lock()

	defer negNumSignSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSignSpec."+
			"NewTrailingNegNumSign()",
		"")

	if err != nil {
		return trailingNegNumSignSpec, err
	}

	if len(trailingNegNumSignSymbols) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'trailingNegNumSignSymbols' is invalid!\n"+
			"'trailingNegNumSignSymbols' is an empty array. The array length\n"+
			"is zero (0)!\n",
			ePrefix.String())

		return trailingNegNumSignSpec, err
	}

	sMechPreon := strMechPreon{}

	var err2 error

	_,
		err2 = sMechPreon.testValidityOfRuneCharArray(
		trailingNegNumSignSymbols,
		nil)

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'trailingNegNumSignSymbols' is invalid!\n"+
			"'trailingNegNumSignSymbols' returned the following validation error:\n"+
			"%v\n",
			ePrefix.String(),
			err2.Error())

		return trailingNegNumSignSpec, err
	}

	err = sMechPreon.copyRuneArrays(
		&trailingNegNumSignSpec.trailingNegNumSignSymbols,
		&trailingNegNumSignSymbols,
		true,
		ePrefix.XCpy(
			"trailingNegNumSignSpec<-trailingNegNumSignSymbols"))

	if err != nil {
		return trailingNegNumSignSpec, err
	}

	trailingNegNumSignSpec.negNumSignPosition = NSignSymPos.After()

	return trailingNegNumSignSpec, err
}

// NewLeadingAndTrailingNegNumSign - Returns a fully populated
// specification for a Leading and Trailing Negative Number Sign.
// Example: (123) - Parenthesis designates a negative numeric value.
//
func (negNumSignSpec NegativeNumberSignSpec) NewLeadingAndTrailingNegNumSign(
	leadingNegNumSignSymbols []rune,
	trailingNegNumSignSymbols []rune,
	errorPrefix interface{}) (
	leadingAndTrailingNegNumSignSpec NegativeNumberSignSpec,
	err error) {

	if negNumSignSpec.lock == nil {
		negNumSignSpec.lock = new(sync.Mutex)
	}

	negNumSignSpec.lock.Lock()

	defer negNumSignSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSignSpec."+
			"NewLeadingAndTrailingNegNumSign()",
		"")

	if err != nil {
		return leadingAndTrailingNegNumSignSpec, err
	}

	if len(leadingNegNumSignSymbols) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'leadingNegNumSignSymbols' is invalid!\n"+
			"'leadingNegNumSignSymbols' is an empty array. The array length\n"+
			"is zero (0)!\n",
			ePrefix.String())

		return leadingAndTrailingNegNumSignSpec, err
	}

	sMechPreon := strMechPreon{}

	var err2 error

	_,
		err2 = sMechPreon.testValidityOfRuneCharArray(
		leadingNegNumSignSymbols,
		nil)

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'leadingNegNumSignSymbols' is invalid!\n"+
			"'leadingNegNumSignSymbols' returned the following validation error:\n"+
			"%v\n",
			ePrefix.String(),
			err2.Error())

		return leadingAndTrailingNegNumSignSpec, err
	}

	if len(trailingNegNumSignSymbols) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'trailingNegNumSignSymbols' is invalid!\n"+
			"'trailingNegNumSignSymbols' is an empty array. The array length\n"+
			"is zero (0)!\n",
			ePrefix.String())

		return leadingAndTrailingNegNumSignSpec, err
	}

	_,
		err2 = sMechPreon.testValidityOfRuneCharArray(
		trailingNegNumSignSymbols,
		nil)

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'trailingNegNumSignSymbols' is invalid!\n"+
			"'trailingNegNumSignSymbols' returned the following validation error:\n"+
			"%v\n",
			ePrefix.String(),
			err2.Error())

		return leadingAndTrailingNegNumSignSpec, err
	}

	err = sMechPreon.copyRuneArrays(
		&leadingAndTrailingNegNumSignSpec.leadingNegNumSignSymbols,
		&leadingNegNumSignSymbols,
		true,
		ePrefix.XCpy(
			"leadingAndTrailingNegNumSignSpec<-leadingNegNumSignSymbols"))

	if err != nil {
		return leadingAndTrailingNegNumSignSpec, err
	}

	err = sMechPreon.copyRuneArrays(
		&leadingAndTrailingNegNumSignSpec.trailingNegNumSignSymbols,
		&trailingNegNumSignSymbols,
		true,
		ePrefix.XCpy(
			"leadingAndTrailingNegNumSignSpec<-trailingNegNumSignSymbols"))

	if err != nil {
		return leadingAndTrailingNegNumSignSpec, err
	}

	leadingAndTrailingNegNumSignSpec.negNumSignPosition = NSignSymPos.BeforeAndAfter()

	return leadingAndTrailingNegNumSignSpec, err
}
