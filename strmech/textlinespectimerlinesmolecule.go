package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type textLineSpecTimerLinesMolecule struct {
	lock *sync.Mutex
}

// copyIn - Copies all data from input parameter
// 'incomingTimerLines' to input parameter 'targetTimerLines'.
//
// Be advised that the pre-existing data fields in input parameter
// 'targetTimerLines' will be overwritten and deleted.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  targetTimerLines           *TextLineSpecTimerLines
//     - A pointer to an instance of TextLineSpecTimerLines. Data
//       extracted from input parameter 'incomingTimerLines' will
//       be copied to this input parameter, 'targetTimerLines'. If
//       this method completes successfully, all member data
//       variables encapsulated in 'targetTimerLines' will be
//       identical to those contained in input parameter,
//       'incomingTimerLines'.
//
//       Be advised that the pre-existing data fields in input
//       parameter 'targetTimerLines' will be overwritten and
//       deleted.
//
//
//  incomingTimerLines         *TextLineSpecTimerLines
//     - A pointer to an instance of TextLineSpecTimerLines.
//
//       All data values in this TextLineSpecTimerLines instance
//       will be copied to input parameter 'targetTimerLines'.
//
//       The original member variable data values encapsulated in
//       'incomingTimerLines' will remain unchanged and will NOT be
//       overwritten or deleted.
//
//       If 'incomingTimerLines' contains invalid member data
//       variables, this method will return an error.
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
func (txtTimerLinesMolecule textLineSpecTimerLinesMolecule) copyIn(
	targetTimerLines *TextLineSpecTimerLines,
	incomingTimerLines *TextLineSpecTimerLines,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtTimerLinesMolecule.lock == nil {
		txtTimerLinesMolecule.lock = new(sync.Mutex)
	}

	txtTimerLinesMolecule.lock.Lock()

	defer txtTimerLinesMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecTimerLinesMolecule.copyIn()",
		"")

	if err != nil {
		return err
	}

	if targetTimerLines == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetTimerLines' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if incomingTimerLines == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'incomingTimerLines' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	_,
		err = textLineSpecTimerLinesAtom{}.ptr().
		testValidityOfTxtSpecTimerLines(
			incomingTimerLines,
			ePrefix.XCtx("incomingTimerLines"))

	if err != nil {
		return err
	}

	sMechPreon := strMechPreon{}

	err = sMechPreon.copyRuneArrays(
		&targetTimerLines.startTimeLabel,
		&incomingTimerLines.startTimeLabel,
		true,
		ePrefix.XCtx("incomingTimerLines.startTimeLabel->"+
			"targetTimerLines.startTimeLabel"))

	if err != nil {
		return err
	}

	targetTimerLines.startTime =
		incomingTimerLines.startTime

	err = sMechPreon.copyRuneArrays(
		&targetTimerLines.endTimeLabel,
		&incomingTimerLines.endTimeLabel,
		true,
		ePrefix.XCtx("incomingTimerLines.endTimeLabel->"+
			"targetTimerLines.endTimeLabel"))

	if err != nil {
		return err
	}

	targetTimerLines.endTime =
		incomingTimerLines.endTime

	targetTimerLines.timeFormat =
		incomingTimerLines.timeFormat

	err = sMechPreon.copyRuneArrays(
		&targetTimerLines.timeDurationLabel,
		&incomingTimerLines.timeDurationLabel,
		true,
		ePrefix.XCtx("incomingTimerLines.timeDurationLabel->"+
			"targetTimerLines.timeDurationLabel"))

	if err != nil {
		return err
	}

	targetTimerLines.labelFieldLen =
		incomingTimerLines.labelFieldLen

	targetTimerLines.labelJustification =
		incomingTimerLines.labelJustification

	err = sMechPreon.copyRuneArrays(
		&targetTimerLines.labelOutputSeparationChars,
		&incomingTimerLines.labelOutputSeparationChars,
		true,
		ePrefix.XCtx("incomingTimerLines.labelOutputSeparationChars->"+
			"targetTimerLines.labelOutputSeparationChars"))

	return err
}

// ptr - Returns a pointer to a new instance of
// textLineSpecTimerLinesMolecule.
//
func (txtTimerLinesMolecule textLineSpecTimerLinesMolecule) ptr() *textLineSpecTimerLinesMolecule {

	if txtTimerLinesMolecule.lock == nil {
		txtTimerLinesMolecule.lock = new(sync.Mutex)
	}

	txtTimerLinesMolecule.lock.Lock()

	defer txtTimerLinesMolecule.lock.Unlock()

	return &textLineSpecTimerLinesMolecule{
		lock: new(sync.Mutex),
	}
}
