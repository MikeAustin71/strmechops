package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type textFieldSpecLabelElectron struct {
	lock *sync.Mutex
}

// isFieldLengthValid - Receives the 'fieldLen' member variable for
// a TextFieldSpecLabel object and subjects it to a diagnostic
// review in order to determine if the field length value is valid.
//
// If input parameter 'fieldLen' is judged invalid, this method
// will return an error along with an appropriate error message.
//
// If input parameter 'fieldLen' is valid, this method will return
// a 'nil' value.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  fieldLen                   int
//     - The length of the text field in which the Text Label will
//       be displayed. If 'fieldLen' is less than the length of
//       Text Label, it will be automatically set equal to the Text
//       Label length.
//
//       If this parameter is submitted with a value less than
//       minus one (-1) or greater than 1-million (1,000,000), an
//       error will be returned.
//
//
//  errPrefDto                 *ErrPrefixDto
//     - This error prefix data transfer object encapsulates an
//       error prefix string which is included in all returned
//       error messages. Usually, it contains the names of the
//       calling functions or methods included in this chain of
//       code execution.
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
//     - If 'fieldLen' is found to be valid, this return parameter
//       will be set to 'nil'.
//
//     - If 'fieldLen' is found to be invalid, this method will
//       return an error along with an appropriate error message.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'errPrefDto'. The
//       'errPrefDto' text will be attached to the beginning of the
//       error message.
//
func (txtFieldLabelElectron *textFieldSpecLabelElectron) isFieldLengthValid(
	fieldLen int,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if txtFieldLabelElectron.lock == nil {
		txtFieldLabelElectron.lock = new(sync.Mutex)
	}

	txtFieldLabelElectron.lock.Lock()

	defer txtFieldLabelElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textFieldSpecLabelElectron.isFieldLengthValid()",
		"")

	if err != nil {
		return err
	}

	if fieldLen < -1 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fieldLen' is less than minus one (-1)!\n"+
			"'fieldLen' controls the length of the Text Label Field.\n",
			ePrefix.String())

		return err
	}

	if fieldLen > 1000000 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fieldLen' is greater than one-million (1,000,000)!\n"+
			"'fieldLen' controls the length of the Text Label Field.\n",
			ePrefix.String())
		return err
	}

	return nil
}

// isTextJustificationValid -  Receives the text justification
// member variable for a TextFieldSpecLabel object and subjects it
// to a diagnostic review in order to determine if the text
// justification value is valid.
//
// Text justification can only be evaluated in the context of a
// text label, field length and 'textJustification' object of type
// TextJustify. This is because text labels with a field length
// equal to or less than the length of the text label never use
// text justification. In these cases, text justification is
// completely ignored.
//
// If the field length is greater than the length of the text
// label, text justification must be equal to one of these three
// valid values:
//           TextJustify(0).Left()
//           TextJustify(0).Right()
//           TextJustify(0).Center()
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  textLabel                  []rune
//     - The text characters to be displayed within the text label.
//
//
//  fieldLen                   int
//     - The length of the text field in which the Text Label will
//       be displayed. If 'fieldLen' is less than the length of
//       Text Label, it will be automatically set equal to the Text
//       Label length.
//
//       If this parameter is submitted with a value less than
//       minus one (-1) or greater than 1-million (1,000,000), an
//       error will be returned.
//
//
//  textJustification          TextJustify
//     - An enumeration which specifies the justification of the
//       'textLabel' within the field specified by 'fieldLen'.
//
//       Text justification can only be evaluated in the context of
//       a text label, field length and 'textJustification' object
//       of type TextJustify. This is because text labels with a
//       field length equal to or less than the length of the text
//       label never use text justification. In these cases, text
//       justification is completely ignored.
//
//       If the field length is greater than the length of the text
//       label, text justification must be equal to one of these
//       three valid values:
//           TextJustify(0).Left()
//           TextJustify(0).Right()
//           TextJustify(0).Center()
//
//
//  errPrefDto                 *ErrPrefixDto
//     - This error prefix data transfer object encapsulates an
//       error prefix string which is included in all returned
//       error messages. Usually, it contains the names of the
//       calling functions or methods included in this chain of
//       code execution.
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
//     - If 'textJustification' is found to be valid, this return
//       parameter will be set to 'nil'.
//
//     - If 'textJustification' is found to be invalid, this method
//       will return an error along with an appropriate error
//       message.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'errPrefDto'. The
//       'errPrefDto' text will be attached to the beginning of the
//       error message.
//
func (txtFieldLabelElectron *textFieldSpecLabelElectron) isTextJustificationValid(
	textLabel []rune,
	fieldLen int,
	textJustification TextJustify,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if txtFieldLabelElectron.lock == nil {
		txtFieldLabelElectron.lock = new(sync.Mutex)
	}

	txtFieldLabelElectron.lock.Lock()

	defer txtFieldLabelElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textFieldSpecLabelElectron.isTextJustificationValid()",
		"")

	if err != nil {
		return err
	}

	if fieldLen <= len(textLabel) {
		return nil
	}

	if !textJustification.XIsValid() {
		err = fmt.Errorf("%v\n"+
			"Error: Text Justification is INVALID!\n"+
			"Text Justification MUST be set to\n"+
			"Left, Right or Center.\n"+
			"Text Label= '%v'\n"+
			"Field Length= '%v'\n"+
			"Text Justification String Value= '%v'\n"+
			"Text Justification Integer Vale= '%v'\n",
			ePrefix.String(),
			string(textLabel),
			fieldLen,
			textJustification.String(),
			textJustification.XValueInt())

		return err
	}

	return nil
}

// isTextLabelValid - Receives the 'textLabel' member variable for
// a TextFieldSpecLabel object and subjects it to a diagnostic
// review in order to determine if the text label is valid.
//
// If input parameter 'textLabel' is judged invalid, this method
// will return an error along with an appropriate error message.
//
// If input parameter 'textLabel' is valid, this method will return
// a 'nil' value.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  textLabel                  []rune
//     - The text characters to be displayed within the text label.
//
//
//  errPrefDto                 *ErrPrefixDto
//     - This error prefix data transfer object encapsulates an
//       error prefix string which is included in all returned
//       error messages. Usually, it contains the names of the
//       calling functions or methods included in this chain of
//       code execution.
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
//  lenTextLabel               int
//     - The length of input parameter 'textLabel'.
//
//
//  err                        error
//     - If 'textLabel' is found to be valid, this return parameter
//       will be set to 'nil'.
//
//     - If 'textLabel' is found to be invalid, this method will
//       return an error along with an appropriate error message.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'errPrefDto'. The
//       'errPrefDto' text will be attached to the beginning of the
//       error message.
//
func (txtFieldLabelElectron *textFieldSpecLabelElectron) isTextLabelValid(
	textLabel []rune,
	errPrefDto *ePref.ErrPrefixDto) (
	lenTextLabel int,
	err error) {

	if txtFieldLabelElectron.lock == nil {
		txtFieldLabelElectron.lock = new(sync.Mutex)
	}

	txtFieldLabelElectron.lock.Lock()

	defer txtFieldLabelElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	lenTextLabel = 0

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"txtFieldFillerElectron.isFillerCharsValid()",
		"")

	if err != nil {
		return lenTextLabel, err
	}

	if textLabel == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'textLabel' is 'nil' and "+
			"has a zero length!\n",
			ePrefix.String())

		return lenTextLabel, err
	}

	lenTextLabel = len(textLabel)

	if lenTextLabel == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'textLabel' is has a "+
			"zero length!\n",
			ePrefix.String())

		return lenTextLabel, err
	}

	var zeroCnt int

	for i := 0; i < lenTextLabel; i++ {

		if textLabel[i] == 0 {
			zeroCnt++
		}
	}

	if zeroCnt == lenTextLabel {
		err = fmt.Errorf("%v\n"+
			"Error: All of the characters in Input parameter\n"+
			"'textLabel' have a zero value!\n",
			ePrefix.String())
	}

	return lenTextLabel, err
}

// ptr - Returns a pointer to a new instance of
// textFieldSpecLabelElectron.
//
func (txtFieldLabelElectron textFieldSpecLabelElectron) ptr() *textFieldSpecLabelElectron {

	if txtFieldLabelElectron.lock == nil {
		txtFieldLabelElectron.lock = new(sync.Mutex)
	}

	txtFieldLabelElectron.lock.Lock()

	defer txtFieldLabelElectron.lock.Unlock()

	return &textFieldSpecLabelElectron{
		lock: new(sync.Mutex),
	}
}
