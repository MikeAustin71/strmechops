package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type textLineSpecTimerLinesNanobot struct {
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
func (txtTimerLinesNanobot *textLineSpecTimerLinesNanobot) copyIn(
	targetTimerLines *TextLineSpecTimerLines,
	incomingTimerLines *TextLineSpecTimerLines,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtTimerLinesNanobot.lock == nil {
		txtTimerLinesNanobot.lock = new(sync.Mutex)
	}

	txtTimerLinesNanobot.lock.Lock()

	defer txtTimerLinesNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecTimerLinesNanobot.copyIn()",
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

	err = textLineSpecTimerLinesMolecule{}.ptr().
		setTxtLineSpecTimerLines(
			targetTimerLines,
			incomingTimerLines.startTimeLabel,
			incomingTimerLines.startTime,
			incomingTimerLines.endTimeLabel,
			incomingTimerLines.endTime,
			incomingTimerLines.timeFormat,
			incomingTimerLines.timeDurationLabel,
			incomingTimerLines.labelFieldLen,
			incomingTimerLines.labelJustification,
			incomingTimerLines.labelOutputSeparationChars,
			ePrefix.XCtx("incomingTimerLines->targetTimerLines"))

	return err
}

// copyOut - Returns a deep copy of the TextLineSpecTimerLines
// input parameter 'txtTimerLines'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  txtTimerLines              *TextLineSpecTimerLines
//     - A pointer to an instance of TextLineSpecTimerLines. A deep
//       copy of the internal member variables will be created
//       and returned in a new instance of TextLineSpecTimerLines.
//
//       If the member variable data values encapsulated by this
//       'txtTimerLines' parameter are found to be invalid, an
//       error will be returned.
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
//  TextLineSpecTimerLines
//     - If this method completes successfully, a deep copy of
//       input parameter 'txtTimerLines' will be created and
//       returned in a new instance of TextLineSpecTimerLines.
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
func (txtTimerLinesNanobot *textLineSpecTimerLinesNanobot) copyOut(
	txtTimerLines *TextLineSpecTimerLines,
	errPrefDto *ePref.ErrPrefixDto) (
	TextLineSpecTimerLines,
	error) {

	if txtTimerLinesNanobot.lock == nil {
		txtTimerLinesNanobot.lock = new(sync.Mutex)
	}

	txtTimerLinesNanobot.lock.Lock()

	defer txtTimerLinesNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecTimerLinesNanobot.copyOut()",
		"")

	if err != nil {
		return TextLineSpecTimerLines{}, err
	}

	if txtTimerLines == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtTimerLines' is a nil pointer!\n",
			ePrefix.String())

		return TextLineSpecTimerLines{}, err
	}

	_,
		err = textLineSpecTimerLinesAtom{}.ptr().
		testValidityOfTxtSpecTimerLines(
			txtTimerLines,
			ePrefix.XCtx("txtTimerLines"))

	if err != nil {
		return TextLineSpecTimerLines{}, err
	}

	newTxtTimerLines := TextLineSpecTimerLines{}

	err = textLineSpecTimerLinesMolecule{}.ptr().
		setTxtLineSpecTimerLines(
			&newTxtTimerLines,
			txtTimerLines.startTimeLabel,
			txtTimerLines.startTime,
			txtTimerLines.endTimeLabel,
			txtTimerLines.endTime,
			txtTimerLines.timeFormat,
			txtTimerLines.timeDurationLabel,
			txtTimerLines.labelFieldLen,
			txtTimerLines.labelJustification,
			txtTimerLines.labelOutputSeparationChars,
			ePrefix.XCtx(
				"txtTimerLines->newTxtTimerLines"))

	if err != nil {
		return TextLineSpecTimerLines{}, err
	}

	newTxtTimerLines.lock = new(sync.Mutex)

	return newTxtTimerLines, nil
}

// ptr - Returns a pointer to a new instance of
// textLineSpecTimerLinesNanobot.
//
func (txtTimerLinesNanobot textLineSpecTimerLinesNanobot) ptr() *textLineSpecTimerLinesNanobot {

	if txtTimerLinesNanobot.lock == nil {
		txtTimerLinesNanobot.lock = new(sync.Mutex)
	}

	txtTimerLinesNanobot.lock.Lock()

	defer txtTimerLinesNanobot.lock.Unlock()

	return &textLineSpecTimerLinesNanobot{
		lock: new(sync.Mutex),
	}
}
