package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type integerSeparatorDtoUtility struct {
	lock *sync.Mutex
}

// setChineseNumberingRunes - Overwrites all the member variable
// data values for the input parameter 'intSep', an instance of
// type IntegerSeparatorDto.
//
// This method is intended to configure Chinese Numbering System
// integer grouping using one or more integer separator
// characters passed through input parameter 'intSeparatorChars'.
//
// The input parameter 'intSeparatorChars' is an array of
// runes containing the integer separator character or characters.
//
// The integer digit grouping is defaulted to a value of four (4).
// Example: '6,7890,000,0000,0000'
//
// The 'integer separator repetitions' value is defaulted to zero
// (0) signaling unlimited repetitions.
//
// This means that integer digits will be separated using the
// Chinese Numbering System (Example: '6,7890,000,0000,0000').
// The Chinese Numbering System separates integer digits into
// groups of four (4) digits. With this method, users have the
// option of specifying one or more integer separator characters
// through input parameter 'intSeparatorChars'. However,
// the// commonly used integer separator for the Chinese Numbering
// system is the comma (',').
//
// # Reference:
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// input parameter 'intSep'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	intSep                     *IntegerSeparatorDto
//	   - A pointer to an instance of IntegerSeparatorDto. All the
//	     member variable data values will be overwritten and reset
//	     using the input parameters listed below and default
//	     values.
//
//
//	intSeparatorChars          []rune
//	   - One or more characters used to separate groups of
//	     integers. It is used to separate groups of integer digits
//	     to the left of the decimal separator (a.k.a. decimal
//	     point). The most common integer separator used in
//	     conjunction with the Chinese Numbering System is the
//	     comma (",").
//	       Example:  6,7890,000,0000,0000
//
//	     If this input parameter contains a zero length array, an
//	     error will be returned.
//
//
//	errPrefDto                 *ErrPrefixDto
//	   - This object encapsulates an error prefix string which is
//	     included in all returned error messages. Usually, it
//	     contains the names of the calling method or methods.
//
//	     Type ErrPrefixDto is included in the 'errpref' software
//	     package, "github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// # Return Values
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
//
// [Wikipedia Chinese Numbering System]: https://en.wikipedia.org/wiki/Chinese_numerals
func (intSeparatorUtil *integerSeparatorDtoUtility) setChineseNumberingRunes(
	intSep *IntegerSeparatorDto,
	intSeparatorChars []rune,
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
			"setIndiaNumberingRunes()",
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

	if len(intSeparatorChars) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'intSeparatorChars' is invalid!\n"+
			"'intSeparatorChars' is an empty string.\n",
			ePrefix.String())

		return err
	}

	_,
		err = new(strMechPreon).
		testValidityOfRuneCharArray(
			intSeparatorChars,
			ePrefix.XCpy(
				"intSeparatorChars"))

	if err != nil {
		return err
	}

	err = new(integerSeparatorDtoQuark).
		empty(
			intSep,
			ePrefix.XCpy(
				"intSep"))

	if err != nil {
		return err
	}

	err =
		new(integerSeparatorDtoMechanics).
			setWithComponents(
				intSep,
				intSeparatorChars,
				[]uint{4},
				false,
				ePrefix.XCpy(
					"intSep []uint{4}"))

	return err
}

// setIndiaNumberingRunes - Overwrites all the member variable
// data values for the input parameter 'intSep', an instance of
// type IntegerSeparatorDto.
//
// This method is intended to configure India Numbering System
// integer grouping using one or more integer separator
// characters passed through input parameter 'intSeparatorChars'.
//
// The input parameter 'intSeparatorChars' is an array of
// runes containing the integer separator character or characters.
//
// The integer digit grouping is defaulted to a value of three (3)
// and two (2). Example: '6,78,90,00,00,00,00,000'
//
// The 'integer separator repetitions' value is defaulted to zero
// (0) signaling unlimited repetitions.
//
// This means that integer digits will be separated using the
// India Numbering System (Example: '6,78,90,00,00,00,00,000')
// with the first group consisting of three digits and the
// remaining groups consisting of two digits. Users have the
// option of specifying integer separator characters through
// input parameter 'intSeparatorChars'. However, the commonly
// used integer separator for the India Numbering system is
// the comma (',').
//
// # Reference:
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// input parameter 'intSep'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	intSep                     *IntegerSeparatorDto
//	   - A pointer to an instance of IntegerSeparatorDto. All the
//	     member variable data values will be overwritten and reset
//	     using the input parameters listed below and default
//	     values.
//
//
//	intSeparatorChars          []rune
//	   - One or more characters used to separate groups of
//	     integers. These characters are used to separate groups of
//	     integer digits to the left of the decimal separator
//	     (a.k.a. decimal point). The most common integer separator
//	     used in conjunction with the India Numbering System is the
//	     comma ([]rune{,}).
//	       Example:  6,78,90,00,00,00,00,000
//
//	     If this input parameter contains a zero length array, an
//	     error will be returned.
//
//
//	errPrefDto                 *ErrPrefixDto
//	   - This object encapsulates an error prefix string which is
//	     included in all returned error messages. Usually, it
//	     contains the names of the calling method or methods.
//
//	     Type ErrPrefixDto is included in the 'errpref' software
//	     package, "github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// # Return Values
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
//
// [Wikipedia India Numbering System]: https://en.wikipedia.org/wiki/Indian_numbering_system
func (intSeparatorUtil *integerSeparatorDtoUtility) setIndiaNumberingRunes(
	intSep *IntegerSeparatorDto,
	intSeparatorChars []rune,
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
			"setIndiaNumberingRunes()",
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

	if len(intSeparatorChars) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'intSeparatorChars' is invalid!\n"+
			"'intSeparatorChars' is an empty string.\n",
			ePrefix.String())

		return err
	}

	_,
		err = new(strMechPreon).
		testValidityOfRuneCharArray(
			intSeparatorChars,
			ePrefix.XCpy(
				"intSeparatorChars"))

	if err != nil {
		return err
	}

	err = new(integerSeparatorDtoQuark).
		empty(
			intSep,
			ePrefix.XCpy(
				"intSep"))

	if err != nil {
		return err
	}

	err =
		new(integerSeparatorDtoMechanics).
			setWithComponents(
				intSep,
				intSeparatorChars,
				[]uint{3, 2},
				false,
				ePrefix.XCpy(
					"intSep []uint{3, 2}"))

	return err
}

// setThousandsRunes - Overwrites all the member variable data values
// for the input parameter 'intSep', an instance of type
// IntegerSeparatorDto.
//
// This method is intended to configure a basic or simple integer
// separator object using default values and a minimum number of
// input parameters.
//
// The input parameter 'intSeparatorChars' is an array of
// runes containing the integer separator characters. The integer
// digit grouping is defaulted to a value of three (3). The
// 'separator repetitions' value is defaulted to zero (0)
// signaling unlimited repetitions.
//
// This means that integer digits will be separated into
// 'thousands' with each group containing three digits each
// (Example: 1,000,000,000). Users have the option of specifying
// integer separator characters through input parameter
// 'intSeparatorChars'.
//
// This method is an alternative to method
// integerSeparatorDtoUtility.setThousands() in that this method
// accepts integer separator characters as an array of runes
// instead of a string.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will overwrite all pre-existing data values in the
// input parameter 'intSep'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	intSep                     *IntegerSeparatorDto
//	   - A pointer to an instance of IntegerSeparatorDto. All the
//	     member variable data values will be overwritten and reset
//	     using the input parameters listed below and default
//	     values.
//
//
//	intSeparatorChars          []rune
//	   - A character, or series of characters, used to separate
//	     integer digits in a number string. These characters are
//	     commonly known as the 'thousands separator'. A 'thousands
//	     separator' is used to separate groups of integer digits to
//	     the left of the decimal separator (a.k.a. decimal point).
//	     In the United States, the standard integer digits
//	     separator is the single comma character (',').
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
//	errPrefDto                 *ErrPrefixDto
//	   - This object encapsulates an error prefix string which is
//	     included in all returned error messages. Usually, it
//	     contains the names of the calling method or methods.
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
func (intSeparatorUtil *integerSeparatorDtoUtility) setThousandsRunes(
	intSep *IntegerSeparatorDto,
	intSeparatorChars []rune,
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
			"setThousandsRunes()",
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

	if len(intSeparatorChars) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'intSeparatorChars' is invalid!\n"+
			"'intSeparatorChars' is an empty string.\n",
			ePrefix.String())

		return err
	}

	_,
		err = new(strMechPreon).
		testValidityOfRuneCharArray(
			intSeparatorChars,
			ePrefix.XCpy(
				"intSeparatorChars"))

	if err != nil {
		return err
	}

	err = new(integerSeparatorDtoQuark).
		empty(
			intSep,
			ePrefix.XCpy(
				"intSep"))

	if err != nil {
		return err
	}

	err =
		new(integerSeparatorDtoMechanics).
			setWithComponents(
				intSep,
				intSeparatorChars,
				[]uint{3},
				false,
				ePrefix.XCpy(
					"intSep []uint{3}"))

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
//	Decimal Separator = '.'
//	Thousands Separator (a.k.a. Integer Digits Separator) = ','
//	Integer Digits Grouping Sequence = 3
//	Example Floating Point Number String: 1,000,000,000.456
//
// # IMPORTANT
//
// This method MAY overwrite all pre-existing data values in the
// input parameter, 'intSep'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	intSep                     *IntegerSeparatorDto
//	   - A pointer to an instance of IntegerSeparatorDto. If this
//	     object is invalid or contains zero data values, all
//	     member variable data values will be overwritten and reset
//	     to United States default integer separator values.
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
		new(integerSeparatorDtoQuark).
			testValidityOfNumStrIntSeparator(
				intSep,
				ePrefix)

	if err == nil {
		return err
	}

	err = new(integerSeparatorDtoQuark).
		empty(
			intSep,
			ePrefix.XCpy(
				"intSep"))

	if err != nil {
		return err
	}

	err = new(integerSeparatorDtoMechanics).
		setToUSADefaults(
			intSep,
			ePrefix.XCpy(
				"intSep"))

	return err
}
