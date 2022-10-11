package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// decimalSepSpecNanobot - Provides helper methods for type
// DecimalSeparatorSpec.
type decimalSepSpecNanobot struct {
	lock *sync.Mutex
}

// copyDecimalSeparator - Copies all data from input parameter
// 'incomingDecSepSpec' to input parameter
// 'targetDecSepSpec'. Both instances are of type
// DecimalSeparatorSpec.
//
// # IMPORTANT
// -----------------------------------------------------------------
// Be advised that the data fields in 'targetDecSepSpec' will
// be deleted and overwritten.
//
// -----------------------------------------------------------------
//
// # Input Parameters
//
//	destinationDecSepSpec      *DecimalSeparatorSpec
//	   - A pointer to a DecimalSeparatorSpec instance. All the
//	     member variable data fields in this object will be
//	     replaced by data values copied from input parameter
//	     'incomingDecSepSpec'.
//
//	     'targetDecSepSpec' is the target of this copy
//	     operation.
//
//
//	sourceDecSepSpec           *DecimalSeparatorSpec
//	   - A pointer to another DecimalSeparatorSpec instance. All
//	     the member variable data values from this object will
//	     be copied to corresponding member variables in
//	     'destinationDecSepSpec'.
//
//	     'sourceDecSepSpec' is the source for this copy
//	     operation.
//
//	     If 'sourceDecSepSpec' is determined to be invalid,
//	     an error will be returned.
//
//
//	errPrefDto                 *ePref.ErrPrefixDto
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
func (decSepSpecNanobot *decimalSepSpecNanobot) copyDecimalSeparator(
	destinationDecSepSpec *DecimalSeparatorSpec,
	sourceDecSepSpec *DecimalSeparatorSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if decSepSpecNanobot.lock == nil {
		decSepSpecNanobot.lock = new(sync.Mutex)
	}

	decSepSpecNanobot.lock.Lock()

	defer decSepSpecNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"decimalSepSpecNanobot."+
			"copyIn()",
		"")

	if err != nil {

		return err

	}

	if destinationDecSepSpec == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'destinationDecSepSpec' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if sourceDecSepSpec == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'sourceDecSepSpec' is a 'nil' pointer!\n",
			ePrefix.String())

		return err
	}

	decSepSpecAtom := decimalSeparatorSpecAtom{}

	decSepSpecAtom.empty(
		destinationDecSepSpec)

	var err2 error

	_,
		err2 = decSepSpecAtom.testValidityOfDecSepSearchSpec(
		sourceDecSepSpec,
		nil)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: Validation of input parameter 'sourceDecSepSpec' failed!\n"+
			"This instance of DecimalSeparatorSpec is invalid.\n"+
			"Validation error message reads as follows:\n"+
			"%v\n",
			ePrefix.String(),
			err2.Error())

		return err
	}

	err = destinationDecSepSpec.decimalSeparatorChars.CopyIn(
		&sourceDecSepSpec.decimalSeparatorChars,
		ePrefix.XCpy(
			"destinationDecSepSpec<-sourceDecSepSpec"))

	return err
}

//	setFrenchGermanDecSep
//
//	Deletes and resets the member variable data values
//	stored in the instance of DecimalSeparatorSpec passed
//	as input parameter 'decSeparatorSpec'.
//
//	Reconfigures the NumStrFormatSpec instance,
//	'decSeparatorSpec', using decimal separator
//	conventions typically applied in France and
//	Germany.
//
//	For French and German numeric values, the
//	radix point or decimal separator is set to
//	the comma character (',').
//
//	As such, the comma character is used to separate
//	integer and fractional digits within a floating
//	point numeric value configured according to French
//	and German standards.
//
// -----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields in 'decSeparatorSpec'
//	will be deleted and overwritten using decimal
//	separator conventions typically applied in France and
//	Germany.
//
// -----------------------------------------------------------------
//
// # Input Parameters
//
//	decSeparatorSpec			*DecimalSeparatorSpec
//
//		A pointer to a DecimalSeparatorSpec instance. All
//		the member variable data values in this object will
//		be replaced and configured for decimal separator
//		conventions typically applied in France and Germany.
//
//		As such, this instance of DecimalSeparatorSpec
//		will be configured with a single comma character
//		(',') for use as a radix point or decimal separator.
//
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string
//		which is included in all returned error
//		messages. Usually, it contains the name of the
//		calling method or methods listed as a function
//		chain.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref'
//		software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//		the error message.
//
//	err							error
//
//		If this method completes successfully, this
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (decSepSpecNanobot *decimalSepSpecNanobot) setFrenchGermanDecSep(
	decSeparatorSpec *DecimalSeparatorSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if decSepSpecNanobot.lock == nil {
		decSepSpecNanobot.lock = new(sync.Mutex)
	}

	decSepSpecNanobot.lock.Lock()

	defer decSepSpecNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"decimalSepSpecNanobot."+
			"setFrenchGermanDecSep()",
		"")

	if err != nil {

		return err

	}

	if decSeparatorSpec == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'decSeparatorSpec' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	new(decimalSeparatorSpecAtom).
		empty(
			decSeparatorSpec)

	decSeparatorSpec.decimalSeparatorChars.CharsArray = []rune{','}
	decSeparatorSpec.decimalSeparatorChars.charSearchType =
		CharSearchType.LinearTargetStartingIndex()

	return err
}
