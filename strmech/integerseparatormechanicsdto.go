package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type integerSeparatorDtoMechanics struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// integerSeparatorDtoMechanics.
func (intSeparatorMech integerSeparatorDtoMechanics) ptr() *integerSeparatorDtoMechanics {

	if intSeparatorMech.lock == nil {
		intSeparatorMech.lock = new(sync.Mutex)
	}

	intSeparatorMech.lock.Lock()

	defer intSeparatorMech.lock.Unlock()

	newIntSeparatorMech := new(integerSeparatorDtoMechanics)

	newIntSeparatorMech.lock = new(sync.Mutex)

	return newIntSeparatorMech
}

// setToUSADefaults - Receives a pointer to an instance of
// IntegerSeparatorDto and proceeds to overwrite and set the
// internal member variable data values to default values
// used in the United States. Integer separator values used
// in the United States consist of the comma character (','),
// an integer grouping of three ('3') and unlimited repetitions of
// this sequence.
//   United States Integer Separation Example:
//         '1,000,000,000,000'
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  nStrIntSep                 *IntegerSeparatorDto
//     - A pointer to an instance of IntegerSeparatorDto. All the
//       internal member variable data values contained in this
//       object will be overwritten and reset to default integer
//       separator values used in the United States.
//           United States Integer Separation Example:
//                  '1,000,000,000,000'
//
//
//  errPrefDto                 *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'errPrefDto'. The
//       'errPrefDto' text will be attached to the beginning of the
//       error message.
//
func (intSeparatorMech *integerSeparatorDtoMechanics) setToUSADefaults(
	nStrIntSep *IntegerSeparatorDto,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if intSeparatorMech.lock == nil {
		intSeparatorMech.lock = new(sync.Mutex)
	}

	intSeparatorMech.lock.Lock()

	defer intSeparatorMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"integerSeparatorDtoMechanics."+
			"setToUSADefaults()",
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

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.intSeparatorChars =
		make([]rune, 1)

	nStrIntSep.intSeparatorChars[0] = ','

	nStrIntSep.intSeparatorGrouping = 3

	nStrIntSep.intSeparatorRepetitions = 0

	nStrIntSep.restartIntGroupingSequence = false

	return err
}

// setWithComponents - Receives a pointer to an instance of
// IntegerSeparatorDto and proceeds to overwrite and set the
// internal member variable data values based on the other input
// parameters.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  nStrIntSep                 *IntegerSeparatorDto
//     - A pointer to an instance of IntegerSeparatorDto. All the
//       internal member variable data values contained in this
//       object will be overwritten and reset based on the
//       following input parameters.
//
//
//  intSeparatorChars          []rune
//     - A series of runes or characters used to separate integer
//       digits in a number string. These characters are commonly
//       known as the 'thousands separator'. A 'thousands
//       separator' is used to separate groups of integer digits to
//       the left of the decimal separator (a.k.a. decimal point).
//       In the United States, the standard integer digits
//       separator is the single comma character (','). Other
//       countries and cultures use periods, spaces, apostrophes or
//       multiple characters to separate integers.
//             United States Example:  1,000,000,000
//
//
//  intSeparatorGrouping       uint
//     - This unsigned integer values specifies the number of
//       integer digits within a group. This value is used to group
//       integers within a number string.
//
//       In most western countries integer digits to the left of
//       the decimal separator (a.k.a. decimal point) are separated
//       into groups of three digits representing a grouping of
//       'thousands' like this: '1,000,000,000'. In this case the
//       intSeparatorGrouping value would be set to three ('3').
//
//       In some countries and cultures other integer groupings are
//       used. In India, for example, a number might be formatted
//       like this: '6,78,90,00,00,00,00,000'.
//
//
//  intSeparatorRepetitions    uint
//     - This unsigned integer value specifies the number of times
//       this integer grouping is repeated. A value of zero signals
//       that this integer grouping will be repeated indefinitely.
//
//
//  restartIntGroupingSequence bool
//     - If the IntegerSeparatorDto is the last element in an array
//       of IntegerSeparatorDto objects, this boolean flag signals
//       whether the entire integer grouping sequence will be
//       restarted from array element zero.
//
//
//  errPrefDto                 *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  err                        error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'errPrefDto'. The
//       'errPrefDto' text will be attached to the beginning of the
//       error message.
//
func (intSeparatorMech *integerSeparatorDtoMechanics) setWithComponents(
	nStrIntSep *IntegerSeparatorDto,
	intSeparatorChars []rune,
	intSeparatorGrouping uint,
	intSeparatorRepetitions uint,
	restartIntGroupingSequence bool,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if intSeparatorMech.lock == nil {
		intSeparatorMech.lock = new(sync.Mutex)
	}

	intSeparatorMech.lock.Lock()

	defer intSeparatorMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"integerSeparatorDtoMechanics."+
			"setWithComponents()",
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

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	if intSeparatorChars == nil {
		intSeparatorChars =
			make([]rune, 0, 5)
	}

	lenIntSepChars := len(intSeparatorChars)

	if lenIntSepChars == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: 'intSeparatorChars' is invalid!\n"+
			"'intSeparatorChars' is a ZERO length rune array.\n",
			ePrefix.String())

		return err
	}

	nStrIntSep.intSeparatorChars =
		make([]rune, lenIntSepChars, lenIntSepChars+5)

	for i := 0; i < lenIntSepChars; i++ {
		nStrIntSep.intSeparatorChars[i] =
			intSeparatorChars[i]
	}

	nStrIntSep.intSeparatorGrouping =
		intSeparatorGrouping

	nStrIntSep.intSeparatorRepetitions =
		intSeparatorRepetitions

	nStrIntSep.restartIntGroupingSequence =
		restartIntGroupingSequence

	return err
}
