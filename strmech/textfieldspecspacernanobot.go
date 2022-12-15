package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
)

type textFieldSpecSpacerNanobot struct {
	lock *sync.Mutex
}

// copyIn - Copies all data from input parameter
// 'incomingTxtFieldSpacer' to input parameter
// 'targetTxtFieldSpacer'.
//
// IMPORTANT
// -----------------------------------------------------------------
// Be advised that the data fields in 'targetTxtFieldSpacer' will be
// overwritten and replaced.
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//	targetTxtFieldSpacer       *TextFieldSpecSpacer
//	   - A pointer to a TextFieldSpecSpacer instance. All the
//	     member variable data fields in this object will be
//	     replaced by data values extracted from input parameter
//	     'incomingTxtFieldSpacer'.
//
//	     'targetTxtFieldSpacer' is the target of this copy
//	     operation.
//
//
//	incomingTxtFieldSpacer     *TextFieldSpecSpacer
//	   - A pointer to another TextFieldSpecSpacer instance. All
//	     the member variable data values from this object will
//	     be copied to corresponding member variables in
//	     'targetTxtFieldSpacer'.
//
//	     'incomingTxtFieldSpacer' is the source for this copy
//	     operation.
//
//	     If 'incomingTxtFieldSpacer' is determined to be invalid,
//	     an error will be returned.
//
//
//	errPrefDto          *ePref.ErrPrefixDto
//	   - This object encapsulates an error prefix string which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods listed
//	     as a function chain.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     Type ErrPrefixDto is included in the 'errpref' software
//	     package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If this method completes successfully, this returned error
//	     Type is set equal to 'nil'. If errors are encountered during
//	     processing, the returned error Type will encapsulate an error
//	     message.
//
//	     If an error message is returned, the text value for input
//	     parameter 'errPrefDto' (error prefix) will be prefixed or
//	     attached at the beginning of the error message.
func (txtFieldSpacerNanobot *textFieldSpecSpacerNanobot) copyIn(
	targetTxtFieldSpacer *TextFieldSpecSpacer,
	incomingTxtFieldSpacer *TextFieldSpecSpacer,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtFieldSpacerNanobot.lock == nil {
		txtFieldSpacerNanobot.lock = new(sync.Mutex)
	}

	txtFieldSpacerNanobot.lock.Lock()

	defer txtFieldSpacerNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textFieldSpecSpacerNanobot.copyIn()",
		"")

	if err != nil {
		return err
	}

	if targetTxtFieldSpacer == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetTxtFieldSpacer' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if incomingTxtFieldSpacer == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'incomingTxtFieldSpacer' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return err
	}

	targetTxtFieldSpacer.textLineReader = nil

	_,
		err = textFieldSpecSpacerElectron{}.ptr().
		isFieldLenValidError(
			incomingTxtFieldSpacer.fieldLen,
			ePrefix.XCpy("incomingTxtFieldSpacer.fieldLen is invalid!"))

	if err != nil {

		return err
	}

	targetTxtFieldSpacer.fieldLen =
		incomingTxtFieldSpacer.fieldLen

	return err
}

// copyOut - Returns a deep copy of the input parameter
// 'txtFieldSpacer'
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	txtFieldSpacer       *TextFieldSpecSpacer
//	   - A pointer to an instance of TextFieldSpecSpacer. A deep
//	     copy of the internal member variables will be created
//	     and returned in a new instance of TextFieldSpecSpacer.
//
//	     If the member variable data values encapsulated by this
//	     'txtFieldSpacer' are found to be invalid, this method will
//	     return an error
//
//
//	errPrefDto          *ePref.ErrPrefixDto
//	   - This object encapsulates an error prefix string which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods listed
//	     as a function chain.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     Type ErrPrefixDto is included in the 'errpref' software
//	     package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	TextFieldSpecSpacer
//	   - If this method completes successfully, a deep copy of
//	     input parameter 'txtFieldSpacer' will be created and
//	     returned in a new instance of TextFieldSpecSpacer.
//
//
//	error
//	   - If this method completes successfully, this returned error
//	     Type is set equal to 'nil'. If errors are encountered during
//	     processing, the returned error Type will encapsulate an error
//	     message.
//
//	     If an error message is returned, the text value for input
//	     parameter 'errPrefDto' (error prefix) will be prefixed or
//	     attached at the beginning of the error message.
func (txtFieldSpacerNanobot *textFieldSpecSpacerNanobot) copyOut(
	txtFieldSpacer *TextFieldSpecSpacer,
	errPrefDto *ePref.ErrPrefixDto) (
	TextFieldSpecSpacer,
	error) {

	if txtFieldSpacerNanobot.lock == nil {
		txtFieldSpacerNanobot.lock = new(sync.Mutex)
	}

	txtFieldSpacerNanobot.lock.Lock()

	defer txtFieldSpacerNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newTxtFieldSpacer := TextFieldSpecSpacer{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textFieldSpecSpacerNanobot.copyOut()",
		"")

	if err != nil {
		return newTxtFieldSpacer, err
	}

	if txtFieldSpacer == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtFieldSpacer' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return newTxtFieldSpacer, err
	}

	_,
		err = textFieldSpecSpacerElectron{}.ptr().
		isFieldLenValidError(
			txtFieldSpacer.fieldLen,
			ePrefix.XCpy("txtFieldSpacer.fieldLen is invalid!"))

	if err != nil {

		return newTxtFieldSpacer, err
	}

	newTxtFieldSpacer.fieldLen =
		txtFieldSpacer.fieldLen

	newTxtFieldSpacer.textLineReader = nil

	return newTxtFieldSpacer, err
}

// empty - Receives a pointer to an instance of TextFieldSpecSpacer
// and proceeds to set all the internal member variables to
// their uninitialized or zero states.
//
// IMPORTANT
// ----------------------------------------------------------------
// The values of all member variables contained in input parameter
// 'txtFieldSpacer' will be overwritten and replaced.
func (txtFieldSpacerNanobot *textFieldSpecSpacerNanobot) empty(
	txtFieldSpacer *TextFieldSpecSpacer) {

	if txtFieldSpacerNanobot.lock == nil {
		txtFieldSpacerNanobot.lock = new(sync.Mutex)
	}

	txtFieldSpacerNanobot.lock.Lock()

	defer txtFieldSpacerNanobot.lock.Unlock()

	if txtFieldSpacer == nil {
		return
	}

	txtFieldSpacer.fieldLen = 0

	txtFieldSpacer.textLineReader = nil

	return
}

// equal - Receives a pointer to two instances of
// TextFieldSpecSpacer and proceeds to compare their member
// variables in order to determine if they are equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables for both instances are equal
// in all respects, this flag is set to 'true'. Otherwise, this
// method returns 'false'.
func (txtFieldSpacerNanobot *textFieldSpecSpacerNanobot) equal(
	txtFieldSpacer *TextFieldSpecSpacer,
	incomingTxtFieldSpacer *TextFieldSpecSpacer) bool {

	if txtFieldSpacerNanobot.lock == nil {
		txtFieldSpacerNanobot.lock = new(sync.Mutex)
	}

	txtFieldSpacerNanobot.lock.Lock()

	defer txtFieldSpacerNanobot.lock.Unlock()

	if txtFieldSpacer == nil ||
		incomingTxtFieldSpacer == nil {
		return false
	}

	if txtFieldSpacer.fieldLen !=
		incomingTxtFieldSpacer.fieldLen {
		return false
	}

	return true
}

// getFormattedText - Returns the formatted text generated by the
// instance of TextFieldSpecSpacer passed as an input parameter.
//
// TextFieldSpecSpacer is a Text Field Specification which produces
// a text string equal to one or more white space characters. The
// number of white space characters in the string is determined by
// the field length parameter. (See examples below)
//
// This method fulfills the requirements of the
// ITextFieldSpecification interface.
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//	fieldLen = 1 produces text field " "
//	fieldLen = 2 produces text field "  "
//	fieldLen = 5 produces text field "     "
func (txtFieldSpacerNanobot *textFieldSpecSpacerNanobot) getFormattedText(
	txtFieldSpacer *TextFieldSpecSpacer,
	errPrefDto *ePref.ErrPrefixDto) (
	formattedText string,
	err error) {

	if txtFieldSpacerNanobot.lock == nil {
		txtFieldSpacerNanobot.lock = new(sync.Mutex)
	}

	txtFieldSpacerNanobot.lock.Lock()

	defer txtFieldSpacerNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textFieldSpecSpacerNanobot."+
			"getFormattedText()",
		"")

	if err != nil {
		return "", err
	}

	if txtFieldSpacer == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtFieldSpacer' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return "", err
	}

	_,
		err = textFieldSpecSpacerElectron{}.ptr().
		isFieldLenValidError(
			txtFieldSpacer.fieldLen,
			ePrefix.XCpy(
				"txtFieldSpacer.fieldLen invalid!"))

	if err != nil {
		return "", err
	}

	formattedText = strings.Repeat(" ",
		txtFieldSpacer.fieldLen)

	return formattedText, err
}

// ptr - Returns a pointer to a new instance of
// textFieldSpecSpacerNanobot.
func (txtFieldSpacerNanobot textFieldSpecSpacerNanobot) ptr() *textFieldSpecSpacerNanobot {

	if txtFieldSpacerNanobot.lock == nil {
		txtFieldSpacerNanobot.lock = new(sync.Mutex)
	}

	txtFieldSpacerNanobot.lock.Lock()

	defer txtFieldSpacerNanobot.lock.Unlock()

	return &textFieldSpecSpacerNanobot{
		lock: new(sync.Mutex),
	}
}

// setTextFieldSpacer - Receives a pointer to an instance of
// TextFieldSpecSpacer and resets the member variable data values
// based on input parameters passed by the calling function.
//
// IMPORTANT
// -----------------------------------------------------------------
// Be advised that the data fields in 'txtFieldSpacer' will be
// overwritten and replaced.
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//	txtFieldSpacer             *TextFieldSpecSpacer
//	   - A pointer to an instance of TextFieldSpecSpacer. The
//	     member variable data values in this instance will be
//	     overwritten and replaced based on the following input
//	     parameters.
//
//
//	fieldLen                   int
//	   - An integer value which specifies the number of white space
//	     characters to be included in the spacer text field.
//
//	     Examples:
//	        fieldLen = 1 produces text field " "
//	        fieldLen = 2 produces text field "  "
//	        fieldLen = 5 produces text field "     "
//
//	     If 'fieldLen' is less than one (+1), an error will be
//	     returned.
//
//	     If 'fieldLen' is greater than one-million (+1,000,000), an
//	     error will be returned.
//
//
//	errPrefDto          *ePref.ErrPrefixDto
//	   - This object encapsulates an error prefix string which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods listed
//	     as a function chain.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     Type ErrPrefixDto is included in the 'errpref' software
//	     package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If this method completes successfully, this returned error
//	     Type is set equal to 'nil'. If errors are encountered during
//	     processing, the returned error Type will encapsulate an error
//	     message.
//
//	     If an error message is returned, the text value for input
//	     parameter 'errPrefDto' (error prefix) will be prefixed or
//	     attached at the beginning of the error message.
func (txtFieldSpacerNanobot *textFieldSpecSpacerNanobot) setTextFieldSpacer(
	txtFieldSpacer *TextFieldSpecSpacer,
	fieldLen int,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtFieldSpacerNanobot.lock == nil {
		txtFieldSpacerNanobot.lock = new(sync.Mutex)
	}

	txtFieldSpacerNanobot.lock.Lock()

	defer txtFieldSpacerNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textFieldSpecSpacerNanobot.setTextFieldSpacer()",
		"")

	if err != nil {
		return err
	}

	if txtFieldSpacer == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtFieldSpacer' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return err
	}

	_,
		err = textFieldSpecSpacerElectron{}.ptr().
		isFieldLenValidError(
			fieldLen,
			ePrefix.XCpy("fieldLen is invalid!"))

	if err != nil {

		return err
	}

	txtFieldSpacer.fieldLen = fieldLen

	return err
}
