package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type textFieldSpecDateTimeNanobot struct {
	lock *sync.Mutex
}

// copyIn - Copies all data from input parameter
//'incomingDateTimeTxtField' to input parameter
// 'targetDateTimeTxtField'.
//
// IMPORTANT
// -----------------------------------------------------------------
// Be advised that the data fields in 'targetDateTimeTxtField' will
// be overwritten.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  targetDateTimeTxtField     *TextFieldSpecDateTime
//     - A pointer to a TextFieldSpecDateTime instance. All the
//       member variable data fields in this object will be
//       replaced by data values extracted from input parameter
//       'incomingDateTimeTxtField'.
//
//       'targetDateTimeTxtField' is the target of this copy
//       operation.
//
//
//  incomingDateTimeTxtField   *TextFieldSpecDateTime
//     - A pointer to another TextFieldSpecDateTime instance. All
//       the member variable data values from this object will
//       be copied to corresponding member variables in
//       'targetDateTimeTxtField'.
//
//       'incomingDateTimeTxtField' is the source for this copy
//       operation.
//
//       If 'incomingDateTimeTxtField' is determined to be invalid,
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
func (txtFieldDateTimeNanobot *textFieldSpecDateTimeNanobot) copyIn(
	targetDateTimeTxtField *TextFieldSpecDateTime,
	incomingDateTimeTxtField *TextFieldSpecDateTime,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtFieldDateTimeNanobot.lock == nil {
		txtFieldDateTimeNanobot.lock = new(sync.Mutex)
	}

	txtFieldDateTimeNanobot.lock.Lock()

	defer txtFieldDateTimeNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textFieldSpecDateTimeNanobot.copyIn()",
		"")

	if err != nil {
		return err
	}

	if targetDateTimeTxtField == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetDateTimeTxtField' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if incomingDateTimeTxtField == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'incomingDateTimeTxtField' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return err
	}

	_,
		err = textFieldSpecDateTimeAtom{}.ptr().isValidTextFieldDateTime(
		incomingDateTimeTxtField,
		ePrefix.XCpy("testing incomingDateTimeTxtField"))

	if err != nil {
		return err
	}

	targetDateTimeTxtField.textLineReader = nil

	targetDateTimeTxtField.dateTime =
		incomingDateTimeTxtField.dateTime

	targetDateTimeTxtField.fieldLen =
		incomingDateTimeTxtField.fieldLen

	targetDateTimeTxtField.dateTimeFormat =
		incomingDateTimeTxtField.dateTimeFormat

	targetDateTimeTxtField.textJustification =
		incomingDateTimeTxtField.textJustification

	return err
}

// copyOut - Returns a deep copy of the input parameter
// 'dateTimeTxtField'
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  dateTimeTxtField           *TextFieldSpecDateTime
//     - A pointer to an instance of TextFieldSpecDateTime. A deep
//       copy of the internal member variables will be created
//       and returned in a new instance of TextFieldSpecDateTime.
//
//       If the member variable data values encapsulated by this
//       'dateTimeTxtField' are found to be invalid, this method
//       will return an error
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
//  TextFieldSpecDateTime
//     - If this method completes successfully, a deep copy of
//       input parameter 'dateTimeTxtField' will be created and
//       returned in a new instance of TextFieldSpecDateTime.
//
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
func (txtFieldDateTimeNanobot *textFieldSpecDateTimeNanobot) copyOut(
	dateTimeTxtField *TextFieldSpecDateTime,
	errPrefDto *ePref.ErrPrefixDto) (
	TextFieldSpecDateTime, error) {

	if txtFieldDateTimeNanobot.lock == nil {
		txtFieldDateTimeNanobot.lock = new(sync.Mutex)
	}

	txtFieldDateTimeNanobot.lock.Lock()

	defer txtFieldDateTimeNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newDateTimeTxtField := TextFieldSpecDateTime{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"txtFieldDateTimeNanobot.copyOut()",
		"")

	if err != nil {
		return newDateTimeTxtField, err
	}

	if dateTimeTxtField == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'dateTimeTxtField' is a nil pointer!\n",
			ePrefix.String())

		return newDateTimeTxtField, err
	}

	_,
		err = textFieldSpecDateTimeAtom{}.ptr().isValidTextFieldDateTime(
		dateTimeTxtField,
		ePrefix.XCpy("testing validity of dateTimeTxtField"))

	if err != nil {
		return newDateTimeTxtField, err
	}

	newDateTimeTxtField.textLineReader = nil

	newDateTimeTxtField.dateTime =
		dateTimeTxtField.dateTime

	newDateTimeTxtField.fieldLen =
		dateTimeTxtField.fieldLen

	newDateTimeTxtField.dateTimeFormat =
		dateTimeTxtField.dateTimeFormat

	newDateTimeTxtField.textJustification =
		dateTimeTxtField.textJustification

	return newDateTimeTxtField, err
}

// getFormattedText - Returns the formatted text generated by the
// instance of TextFieldSpecDateTime passed as an input parameter.
//
// If input parameter 'dateTimeTxtField' is invalid, an error will
// be returned.
//
func (txtFieldDateTimeNanobot *textFieldSpecDateTimeNanobot) getFormattedText(
	dateTimeTxtField *TextFieldSpecDateTime,
	errPrefDto *ePref.ErrPrefixDto) (
	string,
	error) {

	if txtFieldDateTimeNanobot.lock == nil {
		txtFieldDateTimeNanobot.lock = new(sync.Mutex)
	}

	txtFieldDateTimeNanobot.lock.Lock()

	defer txtFieldDateTimeNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"txtFieldDateTimeNanobot.getFormattedText()",
		"")

	if err != nil {
		return "", err
	}

	if dateTimeTxtField == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'dateTimeTxtField' is a nil pointer!\n",
			ePrefix.String())

		return "", err
	}

	_,
		err = textFieldSpecDateTimeAtom{}.ptr().
		isValidTextFieldDateTime(
			dateTimeTxtField,
			ePrefix.XCpy("testing validity of dateTimeTxtField"))

	if err != nil {
		return "", err
	}

	textLabel := dateTimeTxtField.dateTime.Format(
		dateTimeTxtField.dateTimeFormat)

	return textSpecificationMolecule{}.ptr().
		getFormattedText(
			[]rune(textLabel),
			dateTimeTxtField.fieldLen,
			dateTimeTxtField.textJustification,
			ePrefix.XCpy(
				"dateTimeTxtField"))
}

// ptr - Returns a pointer to a new instance of
// textFieldSpecDateTimeNanobot.
//
func (txtFieldDateTimeNanobot textFieldSpecDateTimeNanobot) ptr() *textFieldSpecDateTimeNanobot {

	if txtFieldDateTimeNanobot.lock == nil {
		txtFieldDateTimeNanobot.lock = new(sync.Mutex)
	}

	txtFieldDateTimeNanobot.lock.Lock()

	defer txtFieldDateTimeNanobot.lock.Unlock()

	return &textFieldSpecDateTimeNanobot{
		lock: new(sync.Mutex),
	}
}
