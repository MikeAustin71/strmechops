package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type textSpecificationMolecule struct {
	lock *sync.Mutex
}

// getDefaultTimeFormat - Returns a string containing the default
// format for text display and output of time values.
func (txtSpecMolecule *textSpecificationMolecule) getDefaultTimeFormat() string {

	if txtSpecMolecule.lock == nil {
		txtSpecMolecule.lock = new(sync.Mutex)
	}

	txtSpecMolecule.lock.Lock()

	defer txtSpecMolecule.lock.Unlock()

	return "2006-01-02 15:04:05.000000000 -0700 MST"
}

// getFormattedText - Formats text using text string, field
// length and text justification values.
//
// If the length of the text label string is zero and the field
// length is zero this method returns an empty string.
//
// If the length of the text label string is zero and the field
// length is greater than zero, this method returns a string with
// a length equal to field length and content equal to white space
// (the space character " " x field length).
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  textRunes                  []rune
//     - Contains the text context which will be formatted within a
//       text field specified by input parameter 'fieldLen'.
//
//  fieldLen                   int
//     - The length of the text field in which the 'textLabel' will
//       be displayed. If 'fieldLen' is less than the length of the
//       'textLabel' string, it will be automatically set equal to
//       the 'textLabel' string length.
//
//       To automatically set the value of 'fieldLen' to the length
//       of 'textLabel', set this parameter to a value of minus one
//       (-1).
//
//       If this parameter is submitted with a value less than
//       minus one (-1) or greater than 1-million (1,000,000), an
//       error will be returned.
//
//
//  textJustify                TextJustify
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
//  formattedText              string
//     - The formatted text is returned as a string. If an error
//       occurs, the error message is included in this string.
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
func (txtSpecMolecule *textSpecificationMolecule) getFormattedText(
	textRunes []rune,
	fieldLen int,
	textJustify TextJustify,
	errPrefDto *ePref.ErrPrefixDto) (
	formattedText string,
	err error) {

	if txtSpecMolecule.lock == nil {
		txtSpecMolecule.lock = new(sync.Mutex)
	}

	txtSpecMolecule.lock.Lock()

	defer txtSpecMolecule.lock.Unlock()

	formattedText = ""

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textSpecificationMolecule.getFormattedText()",
		"")

	if err != nil {
		err = fmt.Errorf("%v\n", err.Error())

		return formattedText, err
	}

	txtLabelElectron := textFieldSpecLabelElectron{}

	_,
		err = txtLabelElectron.isTextLabelValid(
		textRunes,
		ePrefix.XCtx("textRunes"))

	if err != nil {
		return formattedText, err
	}

	err = txtLabelElectron.isFieldLengthValid(
		fieldLen,
		ePrefix.XCtx("fieldLen"))

	if err != nil {
		return formattedText, err
	}

	err = txtLabelElectron.isTextJustificationValid(
		textRunes,
		fieldLen,
		textJustify,
		ePrefix.XCtx("textJustify"))

	if err != nil {
		return formattedText, err
	}

	formattedText,
		err = strMechNanobot{}.ptr().
		justifyTextInStrField(
			string(textRunes),
			fieldLen,
			textJustify,
			ePrefix)

	return formattedText, err
}

// ptr - Returns a pointer to a new instance of
// textSpecificationMolecule.
//
func (txtSpecMolecule textSpecificationMolecule) ptr() *textSpecificationMolecule {

	if txtSpecMolecule.lock == nil {
		txtSpecMolecule.lock = new(sync.Mutex)
	}

	txtSpecMolecule.lock.Lock()

	defer txtSpecMolecule.lock.Unlock()

	return &textSpecificationMolecule{
		lock: new(sync.Mutex),
	}
}
