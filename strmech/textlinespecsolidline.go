package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// TextLineSpecSolidLine - Used to create a solid line of text
// characters for text display or printing.
//
// The TextLineSpecSolidLine type provides formatting
// specifications for solid lines output for text display and
// printing. A solid line consists of a single character or
// multiple characters which constitute a solid line and are often
// used for line breaks. Typically, solid lines consist of dashes
// ("---"), underscore characters ("____"), equal signs ("=====")
// or asterisks ("*****"). Multiple characters may be used to
// produce different line sequences ("--*--*--*"). The length of a
// solid is specified by the calling function using input
// parameter 'solidLineCharsRepeatCount'.
//
// By default, all solid lines are terminated by a new line
// character. This means that each instance of a solid will be
// formatted as a single line of text. If another new line
// character or characters is required, reference method:
//    TextLineSpecSolidLine.SetNewLineCharsOverride()
//
type TextLineSpecSolidLine struct {
	leftMargin                int
	solidLineChars            []rune
	solidLineCharsRepeatCount int
	newLineChars              []rune
	lock                      *sync.Mutex
}

// NewSolidLine - Creates and returns a new instance of
// TextLineSpecSolidLine.
//
// The TextLineSpecSolidLine type provides formatting
// specifications for solid lines output for text display and
// printing. A solid line consists of a single character or
// multiple characters which constitute a solid line and are often
// used for line breaks. Typically, solid lines consist of dashes
// ("---"), underscore characters ("____"), equal signs ("=====")
// or asterisks ("*****"). Multiple characters may be used to
// produce different line sequences ("--*--*--*"). The length of a
// solid is specified by the calling function using input
// parameter 'solidLineCharsRepeatCount'. The use of these types of
// solid lines with text display are often described as
// "line breaks" or "breaking lines".
//
// By default, all solid lines of text are terminated with a new
// line character ('\n'). This means that each instance of a solid
// will be formatted as a single line of text. If another new line
// character or characters are required, reference method:
//    TextLineSpecSolidLine.SetNewLineCharsOverride()
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  leftMargin                 int
//     - The number of white space characters which will be
//       inserted on the left side of the solid line.
//
//       Example:
//         solidLineChars = "*"
//         solidLineCharsRepeatCount = 5
//         leftMargin = 3
//         Solid line = "   *****"
//
//       If this value is less than zero (0), it will be set to a
//       default value of zero (0).
//
//
//  solidLineChars             string
//     - This string specifies the character or characters which
//       will comprise the solid line output for text display or
//       printing.
//
//       Example:
//         solidLineChars = "*"
//         solidLineCharsRepeatCount = 5
//         Solid line = "*****"
//
//       If this parameter is submitted as a zero length rune array,
//       or if 'newLineChars' contains invalid characters, it will
//       be set to the default new line character ('\n').
//
//
//  solidLineCharsRepeatCount  int
//     - This integer value specifies the number of times that
//       parameter 'solidLineChars' will be repeated in
//       constructing the solid line.
//
//       If this parameter is submitted with a value less than one
//       (1), an error will be returned.
//
//       Example:
//         solidLineChars = '*'
//         solidLineCharsRepeatCount = 5
//         Solid line = "*****"
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
//  TextLineSpecSolidLine
//     - If this method completes successfully, it will create and
//       return a new instance of TextLineSpecSolidLine which is
//       fully configured with all the parameters necessary to
//       format a solid line of characters for text display output
//       or printing.
//
//
//  err                        error
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (solidLine TextLineSpecSolidLine) NewSolidLine(
	leftMargin int,
	solidLineChars string,
	solidLineCharsRepeatCount int,
	errorPrefix interface{}) (
	txtSolidLine TextLineSpecSolidLine,
	err error) {

	if solidLine.lock == nil {
		solidLine.lock = new(sync.Mutex)
	}

	solidLine.lock.Lock()

	defer solidLine.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecSolidLine.NewSolidLine()",
		"")

	if err != nil {
		return txtSolidLine, err
	}

	err = textLineSpecSolidLineMolecule{}.ptr().
		setTxtSolidLine(
			&txtSolidLine,
			leftMargin,
			[]rune(solidLineChars),
			solidLineCharsRepeatCount,
			[]rune{'\n'},
			ePrefix.XCtx("txtSolidLine"))

	return txtSolidLine, err
}

// SetNewLineCharsOverride - Sets the new line character or
// characters which will be used to terminate the solid line
// generated by this instance of TextLineSpecSolidLine.
//
// By default, all solid lines of text are terminated by a new line
// character ('\n). This means that each instance of a solid will
// be formatted as a single line of text.
//
// The new line character or characters may be customized by
// calling this method.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  newLineChars               string
//     - This rune array contains one or more characters which will
//       be used to terminate the solid text line.
//
//       This example creates a line termination followed by a
//       second blank line:
//         solidLineChars = '*'
//         solidLineCharsRepeatCount = 5
//         newLineChars = "??\n\n"
//         Solid line = "*****??\n"
//                      "\n"
//
//       If this parameter is submitted as a zero length rune array,
//       it will by default be set to a new line character ('\n').
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  --- NONE ---
//
func (solidLine *TextLineSpecSolidLine) SetNewLineCharsOverride(
	newLineChars string) {

	if solidLine.lock == nil {
		solidLine.lock = new(sync.Mutex)
	}

	solidLine.lock.Lock()

	defer solidLine.lock.Unlock()

	if len(newLineChars) == 0 {
		newLineChars = "\n"
	}

	sMechPreon := strMechPreon{}

	newLineRunes := []rune(newLineChars)

	_ = sMechPreon.copyRuneArrays(
		&solidLine.newLineChars,
		&newLineRunes,
		true,
		nil)
}

// TextTypeName - returns a string specifying the type
// of Text Line Specification.
//
// This method fulfills requirements of ITextSpecification
// interface.
//
func (solidLine TextLineSpecSolidLine) TextTypeName() string {

	if solidLine.lock == nil {
		solidLine.lock = new(sync.Mutex)
	}

	solidLine.lock.Lock()

	defer solidLine.lock.Unlock()

	return "TextLineSpecSolidLine"
}

// TextLineSpecName - returns a string specifying the name
// of this Text Line Specification.
//
// This method fulfills requirements of ITextLineSpecification
// interface.
//
func (solidLine TextLineSpecSolidLine) TextLineSpecName() string {

	if solidLine.lock == nil {
		solidLine.lock = new(sync.Mutex)
	}

	solidLine.lock.Lock()

	defer solidLine.lock.Unlock()

	return "TextLineSpecStandardLine"
}
