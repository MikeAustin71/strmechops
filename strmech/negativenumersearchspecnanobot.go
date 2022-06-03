package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type negNumSignSearchNanobot struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// negNumSignSearchNanobot.
//
func (negNumSearchNanobot negNumSignSearchNanobot) ptr() *negNumSignSearchNanobot {

	if negNumSearchNanobot.lock == nil {
		negNumSearchNanobot.lock = new(sync.Mutex)
	}

	negNumSearchNanobot.lock.Lock()

	defer negNumSearchNanobot.lock.Unlock()

	return &negNumSignSearchNanobot{
		lock: new(sync.Mutex),
	}
}

// setLeadingNegNumSearchSpec - Receives an instance of
// NegativeNumberSearchSpec and proceeds to configure that instance
// as a Leading Negative Number Sign Specification. All internal
// member variables are then configured using the input parameter
// 'leadingNegNumSignSymbols'.
//
// Any previous configuration data associated with this instance of
// NegativeNumberSearchSpec will be deleted before applying the
// new configuration specifications.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  negNumSignSpec             *NegativeNumberSearchSpec
//     - A pointer to an instance of NegativeNumberSearchSpec. This
//       instance will be configured as a Leading Negative Number
//       Sign Specification. All previous configuration data will be
//       deleted and replaced with a new Leading Negative Number
//       Sign configuration.
//
//
//  leadingNegNumSignSymbols   []rune
//     - An array of runes identifying the character or characters
//       which comprise the Leading Negative Number Symbol used in
//       configuring the NegativeNumberSearchSpec instance,
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
func (negNumSearchNanobot *negNumSignSearchNanobot) setLeadingNegNumSearchSpec(
	negNumSignSpec *NegativeNumberSearchSpec,
	leadingNegNumSignSymbols []rune,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if negNumSearchNanobot.lock == nil {
		negNumSearchNanobot.lock = new(sync.Mutex)
	}

	negNumSearchNanobot.lock.Lock()

	defer negNumSearchNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"negNumSignSearchNanobot."+
			"setLeadingNegNumSearchSpec()",
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

	negNumSignAtom := negNumSearchSpecAtom{}

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

// setLeadingAndTrailingNegNumSearchSpec - Receives an instance of
// NegativeNumberSearchSpec and proceeds to configure that instance
// as a Leading and Trailing Negative Number Sign Specification.
// All internal member variables are then configured using the
// input parameter 'leadingNegNumSignSymbols' and
// 'trailingNegNumSignSymbols'.
//
// Any previous configuration data associated with this instance of
// NegativeNumberSearchSpec will be deleted and replaced with the new
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
//  negNumSignSpec             *NegativeNumberSearchSpec
//     - A pointer to an instance of NegativeNumberSearchSpec. This
//       instance will be configured as a Leading and Trailing
//       Negative Number Sign Specification. All previous
//       configuration data will be deleted and replaced with a new
//       Leading and Trailing Negative Number Sign configuration.
//
//
//  leadingNegNumSignSymbols   []rune
//     - An array of runes identifying the character or characters
//       which comprise the Leading Negative Number Symbol used in
//       configuring the NegativeNumberSearchSpec instance,
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
//       configuring the NegativeNumberSearchSpec instance,
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
func (negNumSearchNanobot *negNumSignSearchNanobot) setLeadingAndTrailingNegNumSearchSpec(
	negNumSignSpec *NegativeNumberSearchSpec,
	leadingNegNumSignSymbols []rune,
	trailingNegNumSignSymbols []rune,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if negNumSearchNanobot.lock == nil {
		negNumSearchNanobot.lock = new(sync.Mutex)
	}

	negNumSearchNanobot.lock.Lock()

	defer negNumSearchNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"negNumSignSearchNanobot."+
			"setLeadingAndTrailingNegNumSearchSpec()",
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

// setTrailingNegNumSearchSpec - Receives an instance of
// NegativeNumberSearchSpec and proceeds to configure that instance
// as a Trailing Negative Number Sign Specification. All internal
// member variables are then configured using the input parameter
// 'trailingNegNumSignSymbols'.
//
// Any previous configuration data associated with this instance of
// NegativeNumberSearchSpec will be deleted before applying the
// new configuration specifications.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  negNumSignSpec             *NegativeNumberSearchSpec
//     - A pointer to an instance of NegativeNumberSearchSpec. This
//       instance will be configured as a Trailing Negative Number
//       Sign Specification. All previous configuration data will
//       be deleted and replaced with a new Trailing Negative Number
//       Sign configuration.
//
//
//  trailingNegNumSignSymbols  []rune
//     - An array of runes identifying the character or characters
//       which comprise the Trailing Negative Number Symbol used in
//       configuring the NegativeNumberSearchSpec instance,
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
func (negNumSearchNanobot *negNumSignSearchNanobot) setTrailingNegNumSearchSpec(
	negNumSignSpec *NegativeNumberSearchSpec,
	trailingNegNumSignSymbols []rune,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if negNumSearchNanobot.lock == nil {
		negNumSearchNanobot.lock = new(sync.Mutex)
	}

	negNumSearchNanobot.lock.Lock()

	defer negNumSearchNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"negNumSignSearchNanobot."+
			"setTrailingNegNumSearchSpec()",
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

	negNumSearchSpecAtom{}.ptr().empty(
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
