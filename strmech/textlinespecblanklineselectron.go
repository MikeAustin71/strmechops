package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type textLineSpecBlankLinesElectron struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// textLineSpecBlankLinesElectron.
//
func (txtBlankLinesElectron textLineSpecBlankLinesElectron) ptr() *textLineSpecBlankLinesElectron {

	if txtBlankLinesElectron.lock == nil {
		txtBlankLinesElectron.lock = new(sync.Mutex)
	}

	txtBlankLinesElectron.lock.Lock()

	defer txtBlankLinesElectron.lock.Unlock()

	return &textLineSpecBlankLinesElectron{
		lock: new(sync.Mutex),
	}
}

// testValidityNumOfBlankLines - Tests the validity of the "number of
// blank lines" parameter used in configuring instances of
// TextLineSpecBlankLines.
//
// If the "number of blank lines" parameter is invalid, this method
// will return an error.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  numOfBlankLines            int
//     - The number of blank lines which will be generated by an
//       instance of TextLineSpecBlankLines. This value will be
//       tested for validity.
//
//       If this value is found to be invalid, an error will be
//       returned.
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
//     - If input parameter 'numOfBlankLines' is evaluated as a
//       valid value, this returned error Type will be set equal to
//       'nil'. If input parameter 'numOfBlankLines' is found to be
//       invalid, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (txtBlankLinesElectron *textLineSpecBlankLinesElectron) testValidityNumOfBlankLines(
	numOfBlankLines int,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtBlankLinesElectron.lock == nil {
		txtBlankLinesElectron.lock = new(sync.Mutex)
	}

	txtBlankLinesElectron.lock.Lock()

	defer txtBlankLinesElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecBlankLinesElectron."+
			"testValidityNumOfBlankLines()",
		"")

	if err != nil {
		return err
	}

	if numOfBlankLines < 1 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numOfBlankLines' is invalid!\n"+
			"'numOfBlankLines' is less than one (1).\n",
			ePrefix.String())

		return err
	}

	if numOfBlankLines > 1000000 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numOfBlankLines' is invalid!\n"+
			"'numOfBlankLines' is greater than one-million (1,000,000).\n",
			ePrefix.String())

		return err
	}

	return err
}

// testValidityNewLinesChar - Tests the validity of the
// "new-characters" parameter used in configuring instances of
// TextLineSpecNewLinesChar.
//
// If the "new-characters" parameter is invalid, this method will
// return an error.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  newLineChars               []rune
//     - This rune array holds the text character or characters
//       which will be used as the "new-line" character(s) for an
//       instance of TextLineSpecBlankLines.
//
//       If this array has a 'nil' value or a length of zero, an
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
//  error
//     - If input parameter 'newLineChars' is evaluated as a valid
//       value, this returned error Type will be set equal to
//       'nil'. If input parameter 'newLineChars' is found to be
//       invalid, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (txtBlankLinesElectron *textLineSpecBlankLinesElectron) testValidityNewLinesChars(
	newLineChars []rune,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtBlankLinesElectron.lock == nil {
		txtBlankLinesElectron.lock = new(sync.Mutex)
	}

	txtBlankLinesElectron.lock.Lock()

	defer txtBlankLinesElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecNewLinesCharElectron."+
			"testValidityNewLinesChars()",
		"")

	if err != nil {
		return err
	}

	if newLineChars == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'newLineChars' is invalid!\n"+
			"'newLineChars' is 'nil'.\n",
			ePrefix.String())

		return err
	}

	if len(newLineChars) == 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'newLineChars' is invalid!\n"+
			"The length of 'newLineChars' is Zero (0).\n",
			ePrefix.String())

		return err
	}

	return err
}
