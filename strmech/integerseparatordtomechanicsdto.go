package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type integerSeparatorDtoMechanics struct {
	lock *sync.Mutex
}

// setToFrenchDefaults
//
//	Receives a pointer to an instance of IntegerSeparatorDto
//	and proceeds to overwrite and set the internal member
//	variable data values to default values used in France.
//
//	Integer separator values used in France consist of the
//	space or empty character (' '), an integer grouping of
//	three ('3') and unlimited repetitions of this sequence.
//
//	French Integer Separation Example:
//		'1 000 000 000 000'
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	All the internal member variables in input parameter
//	'nStrIntSep' will be deleted and replaced.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	nStrIntSep					*IntegerSeparatorDto
//
//		A pointer to an instance of IntegerSeparatorDto. All
//		the internal member variable data values contained in
//		this object will be deleted and reset to default
//		integer separator values used in France.
//
//			French Integer Separation Example:
//					'1 000 000 000 000'
//
//	errPrefDto                 *ErrPrefixDto
//
//		This object encapsulates an error prefix string which
//		is included in all returned error messages. Usually,
//		it contains the names of the calling method or methods.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref'
//		software package:
//			"github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//		This returned error message will incorporate the method
//		chain and text passed by input parameter, 'errPrefDto'.
//		The 'errPrefDto' text will be attached to the beginning
//		of the error message.
func (intSeparatorMech *integerSeparatorDtoMechanics) setToFrenchDefaults(
	nStrIntSep *IntegerSeparatorDto,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if intSeparatorMech.lock == nil {
		intSeparatorMech.lock = new(sync.Mutex)
	}

	intSeparatorMech.lock.Lock()

	defer intSeparatorMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"integerSeparatorDtoMechanics."+
			"setToFrenchDefaults()",
		"")

	if err != nil {
		return err
	}

	if nStrIntSep == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrIntSep' is invalid!\n"+
			"'nStrIntSep' is a nil pointer.",
			ePrefix.String())

		return err
	}

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.intSeparatorChars = []rune{' '}

	nStrIntSep.intGroupingSequence = []uint{3}

	nStrIntSep.restartIntGroupingSequence = false

	return err
}

// setToGermanDefaults
//
//	Receives a pointer to an instance of IntegerSeparatorDto
//	and proceeds to overwrite and set the internal member
//	variable data values to default values used in Germany.
//
//	Integer separator values used in Germany consist of the
//	period character ('.'), an integer grouping of
//	three ('3') and unlimited repetitions of this sequence.
//
//	German Integer Separation Example:
//		'1.000.000.000.000'
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	All the internal member variables in input parameter
//	'nStrIntSep' will be deleted and replaced.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	nStrIntSep					*IntegerSeparatorDto
//
//		A pointer to an instance of IntegerSeparatorDto. All
//		the internal member variable data values contained in
//		this object will be deleted and reset to default
//		integer separator values used in Germany.
//
//			German Integer Separation Example:
//					'1.000.000.000.000'
//
//	errPrefDto                 *ErrPrefixDto
//
//		This object encapsulates an error prefix string which
//		is included in all returned error messages. Usually,
//		it contains the names of the calling method or methods.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref'
//		software package:
//			"github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error message.
//		This returned error message will incorporate the method
//		chain and text passed by input parameter, 'errPrefDto'.
//		The 'errPrefDto' text will be attached to the beginning
//		of the error message.
func (intSeparatorMech *integerSeparatorDtoMechanics) setToGermanDefaults(
	nStrIntSep *IntegerSeparatorDto,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if intSeparatorMech.lock == nil {
		intSeparatorMech.lock = new(sync.Mutex)
	}

	intSeparatorMech.lock.Lock()

	defer intSeparatorMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"integerSeparatorDtoMechanics."+
			"setToGermanDefaults()",
		"")

	if err != nil {
		return err
	}

	if nStrIntSep == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrIntSep' is invalid!\n"+
			"'nStrIntSep' is a nil pointer.",
			ePrefix.String())

		return err
	}

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.intSeparatorChars = []rune{'.'}

	nStrIntSep.intGroupingSequence = []uint{3}

	nStrIntSep.restartIntGroupingSequence = false

	return err
}

// setToUSADefaults - Receives a pointer to an instance of
// IntegerSeparatorDto and proceeds to overwrite and set the
// internal member variable data values to default values
// used in the United States. Integer separator values used
// in the United States consist of the comma character (','),
// an integer grouping of three ('3') and unlimited repetitions of
// this sequence.
//
//	United States Integer Separation Example:
//	      '1,000,000,000,000'
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// All the internal member variables in input parameter
// 'nStrIntSep' will be deleted and replaced.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	nStrIntSep                 *IntegerSeparatorDto
//	   - A pointer to an instance of IntegerSeparatorDto. All the
//	     internal member variable data values contained in this
//	     object will be deleted and reset to default integer
//	     separator values used in the United States.
//	         United States Integer Separation Example:
//	                '1,000,000,000,000'
//
//
//	errPrefDto                 *ErrPrefixDto
//	   - This object encapsulates an error prefix string which is
//	     included in all returned error messages. Usually, it
//	     contains the names of the calling method or methods.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     Type ErrPrefixDto is included in the 'errpref' software
//	     package, "github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If this method completes successfully, the returned error
//	     Type is set equal to 'nil'.
//
//	     If errors are encountered during processing, the returned
//	     error Type will encapsulate an error message. This
//	     returned error message will incorporate the method chain
//	     and text passed by input parameter, 'errPrefDto'. The
//	     'errPrefDto' text will be attached to the beginning of the
//	     error message.
func (intSeparatorMech *integerSeparatorDtoMechanics) setToUSADefaults(
	nStrIntSep *IntegerSeparatorDto,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if intSeparatorMech.lock == nil {
		intSeparatorMech.lock = new(sync.Mutex)
	}

	intSeparatorMech.lock.Lock()

	defer intSeparatorMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"integerSeparatorDtoMechanics."+
			"setToUSADefaults()",
		"")

	if err != nil {
		return err
	}

	if nStrIntSep == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrIntSep' is invalid!\n"+
			"'nStrIntSep' is a nil pointer.",
			ePrefix.String())

		return err
	}

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	nStrIntSep.intSeparatorChars = []rune{','}

	nStrIntSep.intGroupingSequence = []uint{3}

	nStrIntSep.restartIntGroupingSequence = false

	return err
}

// setWithComponents - Receives a pointer to an instance of
// IntegerSeparatorDto and proceeds to overwrite and set the
// internal member variable data values based on the other input
// parameters.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// All the internal member variables in input parameter
// 'nStrIntSep' will be deleted and replaced.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	nStrIntSep                 *IntegerSeparatorDto
//	   - A pointer to an instance of IntegerSeparatorDto. All the
//	     internal member variable data values contained in this
//	     object will be deleted and reset based on the
//	     following input parameters.
//
//
//	intSeparatorChars          []rune
//	   - A series of runes or characters used to separate integer
//	     digits in a number string. These characters are commonly
//	     known as the 'thousands separator'. A 'thousands
//	     separator' is used to separate groups of integer digits to
//	     the left of the decimal separator (a.k.a. decimal point).
//	     In the United States, the standard integer digits
//	     separator is the single comma character (','). Other
//	     countries and cultures use periods, spaces, apostrophes or
//	     multiple characters to separate integers.
//	           United States Example:  1,000,000,000
//
//	     In many European countries, a single period ('.') is used
//	     as the integer separator character.
//	           European Example: 1.000.000.000
//
//	     Other countries and cultures use spaces, apostrophes or
//	     multiple characters to separate integers.
//
//	     If this parameter is submitted as a zero length array, an
//	     error will be returned.
//
//
//	intGroupingSequence        []uint
//	   - These unsigned integer values specify the number of
//	     integer digits within each group. This value is used to
//	     group integers within a number string.
//
//	     In most western countries, integer digits to the left of
//	     the decimal separator (a.k.a. decimal point) are separated
//	     into groups of three digits representing a grouping of
//	     'thousands' like this: '1,000,000,000'. In this case the
//	     intGroupingSequence value would be set to three
//	     ([]uint{3}).
//
//	     In some countries and cultures other integer groupings are
//	     used. In India, for example, a number might be formatted
//	     like this: '6,78,90,00,00,00,00,000'. The grouping sequence
//	     for the Indian numbering system is therefore []uint{3,2} .
//
//
//	restartIntGroupingSequence bool
//	   - If this flag is set to 'true', the grouping sequence
//	     will be restarted at the beginning of the
//	     'intGroupingSequence' array after completion of the last
//	      group in the 'intGroupingSequence' array.
//	      Example:
//	        restartIntGroupingSequence = 'true'
//	        intGroupingSequence = uint{3,2}
//	        integer = 1234567890123456
//	        result  = 1,23,456,78,901,23,456
//
//	     If this flag is set to 'false', the last element or
//	     grouping in the 'intGroupingSequence' array will simply be
//	     repeated for all the remaining integer digits.
//	      Example:
//	        restartIntGroupingSequence = 'false'
//	        intGroupingSequence = uint{3,2}
//	        integer = 1234567890123456
//	        result  = 1,23,45,67,89,01,23,456
//
//	     The need to set this value to 'true' is exceedingly rare.
//	     For the vast majority of integer separation scenarios, set
//	     this parameter should be set to 'false'.
//
//
//	errPrefDto                 *ErrPrefixDto
//	   - This object encapsulates an error prefix string which is
//	     included in all returned error messages. Usually, it
//	     contains the names of the calling method or methods.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     Type ErrPrefixDto is included in the 'errpref' software
//	     package, "github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// Return Values
//
//	err                        error
//	   - If this method completes successfully, the returned error
//	     Type is set equal to 'nil'.
//
//	     If errors are encountered during processing, the returned
//	     error Type will encapsulate an error message. This
//	     returned error message will incorporate the method chain
//	     and text passed by input parameter, 'errPrefDto'. The
//	     'errPrefDto' text will be attached to the beginning of the
//	     error message.
func (intSeparatorMech *integerSeparatorDtoMechanics) setWithComponents(
	nStrIntSep *IntegerSeparatorDto,
	intSeparatorChars []rune,
	intGroupingSequence []uint,
	restartIntGroupingSequence bool,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if intSeparatorMech.lock == nil {
		intSeparatorMech.lock = new(sync.Mutex)
	}

	intSeparatorMech.lock.Lock()

	defer intSeparatorMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"integerSeparatorDtoMechanics."+
			"setWithComponents()",
		"")

	if err != nil {
		return err
	}

	if nStrIntSep == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrIntSep' is invalid!\n"+
			"'nStrIntSep' is a nil pointer.",
			ePrefix.String())

		return err
	}

	if nStrIntSep.lock == nil {
		nStrIntSep.lock = new(sync.Mutex)
	}

	if intSeparatorChars == nil {
		intSeparatorChars =
			make([]rune, 0, 5)
	}

	lenIntSepChars := len(intSeparatorChars)

	if lenIntSepChars == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: 'intSeparatorChars' is invalid!\n"+
			"'intSeparatorChars' is a ZERO length rune array.\n",
			ePrefix.String())

		return err
	}

	sMechPreon := strMechPreon{}

	err = sMechPreon.copyRuneArrays(
		&nStrIntSep.intSeparatorChars,
		&intSeparatorChars,
		true,
		ePrefix.XCpy(
			"intSeparatorChars->nStrIntSep.intSeparatorChars"))

	if err != nil {
		return err
	}

	err = sMechPreon.copyUnsignedIntArrays(
		&nStrIntSep.intGroupingSequence,
		&intGroupingSequence,
		true,
		ePrefix.XCpy(
			"intGroupingSequence->nStrIntSep.intGroupingSequence"))

	if err != nil {
		return err
	}

	nStrIntSep.restartIntGroupingSequence =
		restartIntGroupingSequence

	return err
}
