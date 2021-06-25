package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type integerSeparatorDtoUtility struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// integerSeparatorDtoUtility.
func (intSeparatorUtil integerSeparatorDtoUtility) ptr() *integerSeparatorDtoUtility {

	if intSeparatorUtil.lock == nil {
		intSeparatorUtil.lock = new(sync.Mutex)
	}

	intSeparatorUtil.lock.Lock()

	defer intSeparatorUtil.lock.Unlock()

	newIntSepUtility := new(integerSeparatorDtoUtility)

	newIntSepUtility.lock = new(sync.Mutex)

	return newIntSepUtility
}

// setBasic - Overwrites all the member variable data values for
// the input parameter 'intSep', an instance of type
// IntegerSeparatorDto.
//
// This method is intended to configure a basic or simple integer
// separator object using default values and a minimum number of
// input parameters.
//
// The input parameter 'integerDigitsSeparators' is string
// containing the integer separator characters. The integer digit
// grouping is defaulted to a value of three (3). The 'separator
// repetitions' value is defaulted to zero (0) signaling unlimited
// repetitions.
//
// This means that integer digits will be separated into 'thousands'
// with each group containing three digits each (Example:
// 1,000,000,000). Users have the option of specifying integer
// separator characters through input parameter
// 'integerDigitsSeparators'.
//
// IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// input parameter 'intSep'.
//
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  intSep                     *IntegerSeparatorDto
//     - A pointer to an instance of IntegerSeparatorDto. All of
//       the member variable data values will be overwritten and
//       reset using the input parameters listed below and default
//       values.
//
//
//  integerDigitsSeparators    string
//     - One or more characters used to separate groups of
//       integers. This separator is also known as the 'thousands'
//       separator. It is used to separate groups of integer digits
//       to the left of the decimal separator
//       (a.k.a. decimal point). In the United States, the standard
//       integer digits separator is the comma (",").
//
//             Example:  1,000,000,000
//
//       If this input parameter contains a zero length string, an
//       error will be returned.
//
//
//  errPrefDto                 *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  err                        error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'errPrefDto'. The
//       'errPrefDto' text will be attached to the beginning of the
//       error message.
//
func (intSeparatorUtil *integerSeparatorDtoUtility) setBasic(
	intSep *IntegerSeparatorDto,
	integerDigitsSeparators string,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if intSeparatorUtil.lock == nil {
		intSeparatorUtil.lock = new(sync.Mutex)
	}

	intSeparatorUtil.lock.Lock()

	defer intSeparatorUtil.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"integerSeparatorDtoUtility."+
			"setBasic()",
		"")

	if err != nil {
		return err
	}

	if intSep == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'intSep' is invalid!\n"+
			"'intSep' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if len(integerDigitsSeparators) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'integerDigitsSeparators' is invalid!\n"+
			"'integerDigitsSeparators' is an empty string.\n",
			ePrefix.String())

		return err
	}

	err =
		integerSeparatorDtoMechanics{}.ptr().
			setWithComponents(
				intSep,
				[]rune(integerDigitsSeparators),
				3,
				0,
				false,
				ePrefix)

	return err
}

// setBasicRunes - Overwrites all the member variable data values
// for the input parameter 'intSep', an instance of type
// IntegerSeparatorDto.
//
// This method is intended to configure a basic or simple integer
// separator object using default values and a minimum number of
// input parameters.
//
// The input parameter 'integerDigitsSeparators' is an array of
// runes containing the integer separator characters. The integer
// digit grouping is defaulted to a value of three (3). The
// 'separator repetitions' value is defaulted to zero (0) signaling
// unlimited repetitions.
//
// This means that integer digits will be separated into 'thousands'
// with each group containing three digits each (Example:
// 1,000,000,000). Users have the option of specifying integer
// separator characters through input parameter
// 'integerDigitsSeparators'.
//
// This method is an alternative to method
// integerSeparatorDtoUtility.setBasic() in that this method accepts
// integer separator characters as an array of runes instead
// of a string.
//
// IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// input parameter 'intSep'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  intSep                     *IntegerSeparatorDto
//     - A pointer to an instance of IntegerSeparatorDto. All of
//       the member variable data values will be overwritten and
//       reset using the input parameters listed below and default
//       values.
//
//
//  integerDigitsSeparators    []rune
//     - One or more characters used to separate groups of
//       integers. This separator is also known as the 'thousands'
//       separator. It is used to separate groups of integer digits
//       to the left of the decimal separator
//       (a.k.a. decimal point). In the United States, the standard
//       integer digits separator is the comma (',').
//
//             Example:  1,000,000,000
//
//       If this input parameter contains a zero length string, an
//       error will be returned.
//
//
//  errPrefDto                 *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  err                        error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'errPrefDto'. The
//       'errPrefDto' text will be attached to the beginning of the
//       error message.
//
func (intSeparatorUtil *integerSeparatorDtoUtility) setBasicRunes(
	intSep *IntegerSeparatorDto,
	integerDigitsSeparators []rune,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if intSeparatorUtil.lock == nil {
		intSeparatorUtil.lock = new(sync.Mutex)
	}

	intSeparatorUtil.lock.Lock()

	defer intSeparatorUtil.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"integerSeparatorDtoUtility."+
			"setBasicRunes()",
		"")

	if err != nil {
		return err
	}

	if intSep == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'intSep' is invalid!\n"+
			"'intSep' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if len(integerDigitsSeparators) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'integerDigitsSeparators' is invalid!\n"+
			"'integerDigitsSeparators' is an empty string.\n",
			ePrefix.String())

		return err
	}

	err =
		integerSeparatorDtoMechanics{}.ptr().
			setWithComponents(
				intSep,
				integerDigitsSeparators,
				3,
				0,
				false,
				ePrefix)

	return err
}

// setToUSADefaultsIfEmpty - If any of the IntegerSeparatorDto data
// values are zero or invalid, this method will reset ALL data
// elements to United States default values.
//
// If the current IntegerSeparatorDto instance is valid and
// populated with data, this method will take no action and exit.
//
// United States default numeric separators are listed as follows:
//
//  Decimal Separator = '.'
//  Thousands Separator (a.k.a Integer Digits Separator) = ','
//  Integer Digits Grouping Sequence = 3
//  Example Floating Point Number String: 1,000,000,000.456
//
// IMPORTANT
//
// This method MAY overwrite all pre-existing data values in the
// input parameter, 'intSep'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  intSep                     *IntegerSeparatorDto
//     - A pointer to an instance of IntegerSeparatorDto. If this
//       object is invalid or contains zero data values, all
//       member variable data values will be overwritten and reset
//       to United States default integer separator values.
//
//
//  errPrefDto                 *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  err                        error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'errPrefDto'. The
//       'errPrefDto' text will be attached to the beginning of the
//       error message.
//
func (intSeparatorUtil *integerSeparatorDtoUtility) setToUSADefaultsIfEmpty(
	intSep *IntegerSeparatorDto,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if intSeparatorUtil.lock == nil {
		intSeparatorUtil.lock = new(sync.Mutex)
	}

	intSeparatorUtil.lock.Lock()

	defer intSeparatorUtil.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"integerSeparatorDtoUtility."+
			"setToUSADefaultsIfEmpty()",
		"")

	if err != nil {
		return err
	}

	if intSep == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'intSep' is invalid!\n"+
			"'intSep' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	ePrefix.SetEPrefCtx(
		"IntegerSeparatorDto.IsValidInstanceError()",
		"Testing Validity of 'intSep'")
	_,
		err =
		integerSeparatorDtoQuark{}.ptr().
			testValidityOfNumStrIntSeparator(
				intSep,
				ePrefix)

	if err == nil {
		return err
	}

	err = integerSeparatorDtoMechanics{}.ptr().
		setToUSADefaults(
			intSep,
			ePrefix.XCtx(
				"intSep"))

	return err
}
