package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type negNumSignSpecNanobot struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// negNumSignSpecNanobot.
//
func (negNumSignNanobot negNumSignSpecNanobot) ptr() *negNumSignSpecNanobot {

	if negNumSignNanobot.lock == nil {
		negNumSignNanobot.lock = new(sync.Mutex)
	}

	negNumSignNanobot.lock.Lock()

	defer negNumSignNanobot.lock.Unlock()

	return &negNumSignSpecNanobot{
		lock: new(sync.Mutex),
	}
}

// setLeadingNegNumSignSpec - Receives an instance of
// NegativeNumberSignSpec and proceeds to configure that instance
// as a Leading Negative Number Sign Specification. All internal
// member variables are then configured using the input parameter
// 'leadingNegNumSignSymbols'.
//
// Any previous configuration data associated with this instance of
// NegativeNumberSignSpec will be deleted before applying the
// new configuration specifications.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  negNumSignSpec             *NegativeNumberSignSpec
//     - A pointer to an instance of NegativeNumberSignSpec. This
//       instance will be configured as a Leading Negative Number
//       Sign Specification. All previous configuration data will be
//       deleted and replaced with a new Leading Negative Number
//       Sign configuration.
//
//
//  leadingNegNumSignSymbols   []rune
//     - An array of runes identifying the character or characters
//       which comprise the Leading Negative Number Symbol used in
//       configuring the NegativeNumberSignSpec instance,
//       'negNumSignSpec'.
//
//       If this array is empty (zero length) or includes array
//       elements containing a zero value, an error will be
//       returned.
//
//
//  errPrefDto                 *ePref.ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods listed
//       as a function chain.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       Type ErrPrefixDto is included in the 'errpref' software
//       package, "github.com/MikeAustin71/errpref".
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered, this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (negNumSignNanobot *negNumSignSpecNanobot) setLeadingNegNumSignSpec(
	negNumSignSpec *NegativeNumberSignSpec,
	leadingNegNumSignSymbols []rune,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if negNumSignNanobot.lock == nil {
		negNumSignNanobot.lock = new(sync.Mutex)
	}

	negNumSignNanobot.lock.Lock()

	defer negNumSignNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"negNumSignSpecNanobot."+
			"setLeadingNegNumSignSpec()",
		"")

	if err != nil {
		return err
	}

	if negNumSignSpec == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'negNumSignSpec' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if len(leadingNegNumSignSymbols) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'leadingNegNumSignSymbols' is invalid!\n"+
			"'leadingNegNumSignSymbols' is an empty array. The array length\n"+
			"is zero (0)!\n",
			ePrefix.String())

		return err
	}

	negNumSignAtom := negNumSignSpecAtom{}

	negNumSignAtom.empty(
		negNumSignSpec)

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

		return err
	}

	err = sMechPreon.copyRuneArrays(
		&negNumSignSpec.leadingNegNumSignSymbols,
		&leadingNegNumSignSymbols,
		true,
		ePrefix.XCpy(
			"negNumSignSpec<-leadingNegNumSignSymbols"))

	if err != nil {
		return err
	}

	negNumSignSpec.negNumSignPosition = NSignSymPos.Before()

	return err
}

// setLeadingAndTrailingNegNumSignSpec - Receives an instance of
// NegativeNumberSignSpec and proceeds to configure that instance
// as a Leading and Trailing Negative Number Sign Specification.
// All internal member variables are then configured using the
// input parameter 'leadingNegNumSignSymbols' and
// 'trailingNegNumSignSymbols'.
//
// Any previous configuration data associated with this instance of
// NegativeNumberSignSpec will be deleted and replaced with the new
// configuration specifications.
//
// In certain nations and cultures, a pair of symbols is used to
// designate a numeric value as negative. These pairs of symbols
// are described here as a Leading and Trailing Negative Number
// Sign Specification. As an example, in the US and Canada
// parentheses "()" are used to indicate negative numeric
// values. Examples: (127.45) = -127.45  (4,654.00) = -4,654.00
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  negNumSignSpec             *NegativeNumberSignSpec
//     - A pointer to an instance of NegativeNumberSignSpec. This
//       instance will be configured as a Leading and Trailing
//       Negative Number Sign Specification. All previous
//       configuration data will be deleted and replaced with a new
//       Leading and Trailing Negative Number Sign configuration.
//
//
//  leadingNegNumSignSymbols   []rune
//     - An array of runes identifying the character or characters
//       which comprise the Leading Negative Number Symbol used in
//       configuring the NegativeNumberSignSpec instance,
//       'negNumSignSpec'.
//
//       If this array is empty (zero length) or includes array
//       elements containing a zero value, an error will be
//       returned.
//
//
//  trailingNegNumSignSymbols  []rune
//     - An array of runes identifying the character or characters
//       which comprise the Trailing Negative Number Symbol used in
//       configuring the NegativeNumberSignSpec instance,
//       'negNumSignSpec'.
//
//       If this array is empty (zero length) or includes array
//       elements containing a zero value, an error will be
//       returned.
//
//
//  errPrefDto                 *ePref.ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods listed
//       as a function chain.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       Type ErrPrefixDto is included in the 'errpref' software
//       package, "github.com/MikeAustin71/errpref".
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered, this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (negNumSignNanobot *negNumSignSpecNanobot) setLeadingAndTrailingNegNumSignSpec(
	negNumSignSpec *NegativeNumberSignSpec,
	leadingNegNumSignSymbols []rune,
	trailingNegNumSignSymbols []rune,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if negNumSignNanobot.lock == nil {
		negNumSignNanobot.lock = new(sync.Mutex)
	}

	negNumSignNanobot.lock.Lock()

	defer negNumSignNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"negNumSignSpecNanobot."+
			"setLeadingAndTrailingNegNumSignSpec()",
		"")

	if err != nil {
		return err
	}

	if negNumSignSpec == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'negNumSignSpec' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if len(leadingNegNumSignSymbols) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'leadingNegNumSignSymbols' is invalid!\n"+
			"'leadingNegNumSignSymbols' is an empty array. The array length\n"+
			"is zero (0)!\n",
			ePrefix.String())

		return err
	}

	if len(trailingNegNumSignSymbols) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'trailingNegNumSignSymbols' is invalid!\n"+
			"'trailingNegNumSignSymbols' is an empty array. The array length\n"+
			"is zero (0)!\n",
			ePrefix.String())

		return err
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

		return err
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

		return err
	}

	err = sMechPreon.copyRuneArrays(
		&negNumSignSpec.leadingNegNumSignSymbols,
		&leadingNegNumSignSymbols,
		true,
		ePrefix.XCpy(
			"negNumSignSpec<-leadingNegNumSignSymbols"))

	if err != nil {
		return err
	}

	err = sMechPreon.copyRuneArrays(
		&negNumSignSpec.trailingNegNumSignSymbols,
		&trailingNegNumSignSymbols,
		true,
		ePrefix.XCpy(
			"negNumSignSpec<-trailingNegNumSignSymbols"))

	if err != nil {
		return err
	}

	negNumSignSpec.negNumSignPosition = NSignSymPos.BeforeAndAfter()

	return err
}

// setTrailingNegNumSignSpec - Receives an instance of
// NegativeNumberSignSpec and proceeds to configure that instance
// as a Trailing Negative Number Sign Specification. All internal
// member variables are then configured using the input parameter
// 'trailingNegNumSignSymbols'.
//
// Any previous configuration data associated with this instance of
// NegativeNumberSignSpec will be deleted before applying the
// new configuration specifications.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  negNumSignSpec             *NegativeNumberSignSpec
//     - A pointer to an instance of NegativeNumberSignSpec. This
//       instance will be configured as a Trailing Negative Number
//       Sign Specification. All previous configuration data will
//       be deleted and replaced with a new Trailing Negative Number
//       Sign configuration.
//
//
//  trailingNegNumSignSymbols  []rune
//     - An array of runes identifying the character or characters
//       which comprise the Trailing Negative Number Symbol used in
//       configuring the NegativeNumberSignSpec instance,
//       'negNumSignSpec'.
//
//       If this array is empty (zero length) or includes array
//       elements containing a zero value, an error will be
//       returned.
//
//
//  errPrefDto                 *ePref.ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods listed
//       as a function chain.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       Type ErrPrefixDto is included in the 'errpref' software
//       package, "github.com/MikeAustin71/errpref".
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered, this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (negNumSignNanobot *negNumSignSpecNanobot) setTrailingNegNumSignSpec(
	negNumSignSpec *NegativeNumberSignSpec,
	trailingNegNumSignSymbols []rune,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if negNumSignNanobot.lock == nil {
		negNumSignNanobot.lock = new(sync.Mutex)
	}

	negNumSignNanobot.lock.Lock()

	defer negNumSignNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"negNumSignSpecNanobot."+
			"setTrailingNegNumSignSpec()",
		"")

	if err != nil {
		return err
	}

	if negNumSignSpec == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'negNumSignSpec' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if len(trailingNegNumSignSymbols) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'trailingNegNumSignSymbols' is invalid!\n"+
			"'trailingNegNumSignSymbols' is an empty array. The array length\n"+
			"is zero (0)!\n",
			ePrefix.String())

		return err
	}

	negNumSignSpecAtom{}.ptr().empty(
		negNumSignSpec)

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

		return err
	}

	err = sMechPreon.copyRuneArrays(
		&negNumSignSpec.trailingNegNumSignSymbols,
		&trailingNegNumSignSymbols,
		true,
		ePrefix.XCpy(
			"negNumSignSpec<-trailingNegNumSignSymbols"))

	if err != nil {
		return err
	}

	negNumSignSpec.negNumSignPosition = NSignSymPos.After()

	return err
}
