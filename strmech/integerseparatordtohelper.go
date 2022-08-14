package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type integerSeparatorDtoHelper struct {
	lock *sync.Mutex
}

// setFromIntGroupEnum - Configures an instance of
// IntegerSeparatorDto based on an enumeration value
// ('intGroupingType') passed as an input parameter.
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
// intGroupingType             IntegerGroupingType
//
//   - This instance of IntegerGroupingType enumeration defines
//     how 'intSep' will be configured for integer grouping.
//     The enumeration value 'intGroupingType' must be set to one
//     of three values:
//     IntGroupingType.Thousands()
//     IntGroupingType.IndiaNumbering()
//     IntGroupingType.ChineseNumbering()
//
//     Any value other than the three listed above will generate
//     an error.
//
//     intSeparatorChars          []rune
//
//   - A character, or series of characters, used to separate
//     integer digits in a number string. These characters are
//     commonly known as the 'thousands separator'. A 'thousands
//     separator' is used to separate groups of integer digits to
//     the left of the decimal separator (a.k.a. decimal point).
//     In the United States, the standard integer digits
//     separator is the single comma character (',').
//     United States Example:  1,000,000,000
//
//     In many European countries, a single period ('.') is used
//     as the integer separator character.
//     European Example: 1.000.000.000
//
//     Other countries and cultures use spaces, apostrophes or
//     multiple characters to separate integers.
//
//     If this parameter is submitted as a zero length array, an
//     error will be returned.
//
//     errPrefDto                 *ErrPrefixDto
//
//   - This object encapsulates an error prefix string which is
//     included in all returned error messages. Usually, it
//     contains the names of the calling method or methods.
//
//     Type ErrPrefixDto is included in the 'errpref' software
//     package, "github.com/MikeAustin71/errpref".
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
func (intSeparatorHelper *integerSeparatorDtoHelper) setFromIntGroupEnum(
	intSep *IntegerSeparatorDto,
	intGroupingType IntegerGroupingType,
	intSeparatorChars []rune,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if intSeparatorHelper.lock == nil {
		intSeparatorHelper.lock = new(sync.Mutex)
	}

	intSeparatorHelper.lock.Lock()

	defer intSeparatorHelper.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"integerSeparatorDtoUtility."+
			"SetFromIntGroupEnum()",
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

	intSeparatorUtil := integerSeparatorDtoUtility{}

	switch intGroupingType {

	case IntGroupingType.Thousands():

		err = intSeparatorUtil.setThousandsRunes(
			intSep,
			intSeparatorChars,
			ePrefix.XCpy(
				"intSep<-"))

	case intGroupingType.ChineseNumbering():

		err = intSeparatorUtil.setChineseNumberingRunes(
			intSep,
			intSeparatorChars,
			ePrefix.XCpy(
				"intSep<-"))

	case intGroupingType.IndiaNumbering():

		err = intSeparatorUtil.setIndiaNumberingRunes(
			intSep,
			intSeparatorChars,
			ePrefix.XCpy(
				"intSep<-"))

	default:

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'intGroupingType' is invalid!\n"+
			"'intGroupingType' string  value = '%v'\n"+
			"'intGroupingType' integer value = '%v'\n",
			ePrefix.String(),
			intGroupingType.String(),
			intGroupingType.XValueInt())
	}

	return err
}
