package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type textLineSpecSolidLineMolecule struct {
	lock *sync.Mutex
}

// setTxtSolidLine - Sets the member variable data values for an
// instance of TextLineSpecSolidLine passed as input parameter
// 'txtSolidLine'.
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
// By default, all solid lines of text are terminated with a new
// line character ('\n). This means that each instance of a solid
// will be formatted as a single line of text. The new line
// character or characters may be customized by the calling
// function.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  txtSolidLine               *TextLineSpecSolidLine
//     - A pointer to an instance of TextLineSpecSolidLine. The
//       member variables data values encapsulated in this object
//       will be set to new values extracted from the following
//       input parameters.
//
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
//  solidLineChars             []rune
//     - The array of runes specifies the character or characters
//       which will comprise the solid line output for text display
//       or printing.
//
//       Example:
//         solidLineChars = "*"
//         solidLineCharsRepeatCount = 5
//         Solid line = "*****"
///
//       If this parameter is submitted as a zero length rune
//       array, an error will be returned.
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
//  newLineChars               []rune
//     - This rune array contains one or more characters which will
//       be used to terminate the solid text line.
//
//       Example:
//         solidLineChars = '*'
//         solidLineCharsRepeatCount = 5
//         newLineChars = []rune{'??\n')
//         Solid line = "*****??\n"
//
//       If this parameter is submitted as a zero length rune array,
//       it will by default be set to a new line character ('\n').
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
func (txtSpecSolidLine *textLineSpecSolidLineMolecule) setTxtSolidLine(
	txtSolidLine *TextLineSpecSolidLine,
	leftMargin int,
	solidLineChars []rune,
	solidLineCharsRepeatCount int,
	newLineChars []rune,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecSolidLineMolecule."+
			"setTxtSolidLine()",
		"")

	if err != nil {
		return err
	}

	if txtSolidLine == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtSolidLine' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if len(solidLineChars) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'solidLineChars' is invalid!\n"+
			"The length of 'solidLineChars' is Zero.\n",
			ePrefix.String())

		return err
	}

	if solidLineCharsRepeatCount < 1 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'solidLineCharsRepeatCount' is invalid!\n"+
			"The value of 'solidLineCharsRepeatCount' is less than one ('1').\n"+
			"solidLineCharsRepeatCount = '%v'.\n",
			ePrefix.String(),
			solidLineCharsRepeatCount)

		return err
	}

	if leftMargin < 0 {
		leftMargin = 0
	}

	if len(newLineChars) == 0 {
		newLineChars = []rune{'\n'}
	}

	txtSolidLine.leftMargin = leftMargin

	sMechPreon := strMechPreon{}

	err = sMechPreon.copyRuneArrays(
		&txtSolidLine.solidLineChars,
		&solidLineChars,
		true,
		ePrefix.XCtx(
			"solidLineChars->"+
				"txtSolidLine.solidLineChars"))

	if err != nil {
		return err
	}

	txtSolidLine.solidLineCharsRepeatCount =
		solidLineCharsRepeatCount

	err = sMechPreon.copyRuneArrays(
		&txtSolidLine.newLineChars,
		&newLineChars,
		true,
		ePrefix.XCtx(
			"newLineChars->"+
				"txtSolidLine.newLineChars"))

	return err
}

// ptr - Returns a pointer to a new instance of
// textLineSpecSolidLineMolecule.
//
func (txtSpecSolidLine textLineSpecSolidLineMolecule) ptr() *textLineSpecSolidLineMolecule {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	return &textLineSpecSolidLineMolecule{
		lock: new(sync.Mutex),
	}
}
