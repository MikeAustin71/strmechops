package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// decimalSepSearchNanobot - Provides helper methods for type
// DecimalSeparatorSpec.
//
type decimalSepSearchNanobot struct {
	lock *sync.Mutex
}

// copyIn - Copies all data from input parameter
// 'incomingDecSepSearchSpec' to input parameter
// 'targetDecSepSearchSpec'. Both instances are of type
// DecimalSeparatorSpec.
//
// IMPORTANT
// -----------------------------------------------------------------
// Be advised that the data fields in 'targetDecSepSearchSpec' will
// be deleted and overwritten.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  targetDecSepSearchSpec        *DecimalSeparatorSpec
//     - A pointer to a DecimalSeparatorSpec instance. All the
//       member variable data fields in this object will be
//       replaced by data values copied from input parameter
//       'incomingDecSepSearchSpec'.
//
//       'targetDecSepSearchSpec' is the target of this copy
//       operation.
//
//
//  incomingDecSepSearchSpec      *DecimalSeparatorSpec
//     - A pointer to another DecimalSeparatorSpec instance. All
//       the member variable data values from this object will
//       be copied to corresponding member variables in
//       'targetDecSepSearchSpec'.
//
//       'incomingDecSepSearchSpec' is the source for this copy
//       operation.
//
//       If 'incomingDecSepSearchSpec' is determined to be invalid,
//       an error will be returned.
//
//
//  errPrefDto          *ePref.ErrPrefixDto
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
// ------------------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (decSepSearchNanobot *decimalSepSearchNanobot) copyIn(
	targetDecSepSearchSpec *DecimalSeparatorSpec,
	incomingDecSepSearchSpec *DecimalSeparatorSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if decSepSearchNanobot.lock == nil {
		decSepSearchNanobot.lock = new(sync.Mutex)
	}

	decSepSearchNanobot.lock.Lock()

	defer decSepSearchNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"negNumSignSearchNanobot."+
			"copyIn()",
		"")

	if err != nil {

		return err

	}

	if targetDecSepSearchSpec == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'targetNegNumSearchSpec' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if incomingDecSepSearchSpec == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'incomingNegNumSearchSpec' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	var err2 error

	decSepSpecAtom := decimalSeparatorSpecAtom{}

	_,
		err2 = decSepSpecAtom.testValidityOfDecSepSearchSpec(
		incomingDecSepSearchSpec,
		nil)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: Validation of input parameter 'incomingDecSepSearchSpec' failed!\n"+
			"This instance of DecimalSeparatorSpec is invalid.\n"+
			"Validation error message reads as follows:\n"+
			"%v\n",
			ePrefix.String(),
			err2.Error())

		return err
	}

	decSepSpecAtom.empty(
		targetDecSepSearchSpec)

	lenDecSepChars :=
		len(incomingDecSepSearchSpec.decimalSeparatorChars)

	if lenDecSepChars == 0 {
		incomingDecSepSearchSpec.decimalSeparatorChars = nil
		targetDecSepSearchSpec.decimalSeparatorChars = nil
	} else {
		targetDecSepSearchSpec.decimalSeparatorChars =
			make([]rune, lenDecSepChars)

		for i := 0; i < lenDecSepChars; i++ {

			targetDecSepSearchSpec.decimalSeparatorChars[i] =
				incomingDecSepSearchSpec.decimalSeparatorChars[i]

		}
	}

	targetDecSepSearchSpec.foundFirstNumericDigitInNumStr =
		incomingDecSepSearchSpec.foundFirstNumericDigitInNumStr

	targetDecSepSearchSpec.foundDecimalSeparatorSymbols =
		incomingDecSepSearchSpec.foundDecimalSeparatorSymbols

	targetDecSepSearchSpec.foundDecimalSeparatorIndex =
		incomingDecSepSearchSpec.foundDecimalSeparatorIndex

	return err
}

// ptr - Returns a pointer to a new instance of
// decimalSepSearchNanobot.
//
func (decSepSearchNanobot decimalSepSearchNanobot) ptr() *decimalSepSearchNanobot {

	if decSepSearchNanobot.lock == nil {
		decSepSearchNanobot.lock = new(sync.Mutex)
	}

	decSepSearchNanobot.lock.Lock()

	defer decSepSearchNanobot.lock.Unlock()

	return &decimalSepSearchNanobot{
		lock: new(sync.Mutex),
	}
}
