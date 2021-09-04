package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type textLineSpecPlainTextNanobot struct {
	lock *sync.Mutex
}

// copyIn - Copies all data from input parameter
// 'incomingPlainTextLine' to input parameter
// 'targetPlainTextLine'.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// The pre-existing data fields for input parameter 'targetStdLine'
// will be overwritten and deleted.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  targetPlainTextLine        *TextLineSpecPlainText
//     - A pointer to an instance of TextLineSpecPlainText. Data
//       extracted from input parameter 'incomingPlainTextLine'
//       will be copied to this input parameter,
//       'targetPlainTextLine'. If this method completes
//       successfully, all member data variables encapsulated in
//       'targetPlainTextLine' will be identical to those contained
//       in input parameter, 'incomingPlainTextLine'.
//
//       Be advised that the pre-existing data fields in input
//       parameter 'targetPlainTextLine' will be overwritten and
//       deleted.
//
//
//  incomingStdLine            *TextLineSpecPlainText
//     - A pointer to an instance of TextLineSpecPlainText.
//
//       All data values in this TextLineSpecPlainText instance
//       will be copied to input parameter 'targetPlainTextLine'.
//
//       The original member variable data values encapsulated in
//       'incomingPlainTextLine' will remain unchanged and will NOT
//       be overwritten or deleted.
//
//       If 'incomingPlainTextLine' contains invalid member data
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
func (txtLinePlainTextNanobot *textLineSpecPlainTextNanobot) copyIn(
	targetPlainTextLine *TextLineSpecPlainText,
	incomingPlainTextLine *TextLineSpecPlainText,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtLinePlainTextNanobot.lock == nil {
		txtLinePlainTextNanobot.lock = new(sync.Mutex)
	}

	txtLinePlainTextNanobot.lock.Lock()

	defer txtLinePlainTextNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecPlainTextNanobot."+
			"copyIn()",
		"")

	if err != nil {
		return err
	}

	if targetPlainTextLine == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetPlainTextLine' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if incomingPlainTextLine == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'incomingPlainTextLine' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return err
	}

	_,
		err = textLineSpecPlainTextAtom{}.ptr().
		testValidityOfTextLineSpecPlainText(
			incomingPlainTextLine,
			ePrefix.XCtx(
				"incomingPlainTextLine"))

	if err != nil {
		return err
	}

	sMechPreon := strMechPreon{}

	err = sMechPreon.copyRuneArrays(
		&targetPlainTextLine.leftMarginChars,
		&incomingPlainTextLine.leftMarginChars,
		true,
		ePrefix.XCtx(
			"incomingPlainTextLine.leftMarginChars->"+
				"targetPlainTextLine.leftMarginChars"))

	if err != nil {
		return err
	}

	err = sMechPreon.copyRuneArrays(
		&targetPlainTextLine.rightMarginChars,
		&incomingPlainTextLine.rightMarginChars,
		true,
		ePrefix.XCtx(
			"incomingPlainTextLine.rightMarginChars->"+
				"targetPlainTextLine.rightMarginChars"))

	if err != nil {
		return err
	}

	targetPlainTextLine.textString =
		incomingPlainTextLine.textString

	targetPlainTextLine.turnLineTerminatorOff =
		incomingPlainTextLine.turnLineTerminatorOff

	err = sMechPreon.copyRuneArrays(
		&targetPlainTextLine.newLineChars,
		&incomingPlainTextLine.newLineChars,
		true,
		ePrefix.XCtx(
			"incomingPlainTextLine.newLineChars->"+
				"targetPlainTextLine.newLineChars"))

	return err
}

// ptr - Returns a pointer to a new instance of
// textLineSpecPlainTextNanobot.
//
func (txtLinePlainTextNanobot textLineSpecPlainTextNanobot) ptr() *textLineSpecPlainTextNanobot {

	if txtLinePlainTextNanobot.lock == nil {
		txtLinePlainTextNanobot.lock = new(sync.Mutex)
	}

	txtLinePlainTextNanobot.lock.Lock()

	defer txtLinePlainTextNanobot.lock.Unlock()

	return &textLineSpecPlainTextNanobot{
		lock: new(sync.Mutex),
	}
}
