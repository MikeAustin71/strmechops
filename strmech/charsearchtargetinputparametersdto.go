package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// CharSearchTargetInputParametersDto - Target Input Parameters are
// more easily understood in the context of text character search
// operations.
//
// ----------------------------------------------------------------
//
// TERMINOLOGY
//
// Text Character Search algorithms typically perform comparisons
// between two strings or groups of text characters to determine
// the search outcome. A successful search outcome usually involves
// finding one or more text characters from one string inside a
// second string. A successful search outcome is often referred to
// as a 'Match' condition because characters in one string were
// compared and matched with characters in another string.
//
// Character Search algorithms using the Character Search Type
// ('CharacterSearchType') rely on a framework consisting of a
// 'Target Search String' and a 'Test String'.
//
//    Target String        - A string character or characters which
//                           will be searched for the occurrence of
//                           another predefined character or
//                           characters referred to as a Test
//                           String.
//
//
//    Test String          - A string character or characters which
//                           will be used to search for matching
//                           characters in a Target Search String.
//
// A comparison of text characters contained in the Target Search
// String and the Test String serves as the basis for determining
// a 'Match' condition or successful outcome from a text character
// search algorithm. The specific criterion for determining a
// 'Match' condition vary between the different Character Search
// Types.
//
// ----------------------------------------------------------------
//
// The Character Search Target Input Parameters Data Transfer
// Object type (CharSearchTargetInputParametersDto) is used to
// transmit Target String input parameters to methods performing
// search operations.
//
type CharSearchTargetInputParametersDto struct {
	TargetString *RuneArrayDto
	// A pointer to the RuneArrayDto containing the Target
	// Search String text characters used in the search
	// algorithm. Target Characters are compared against
	// Test Characters to determine if a 'Match' condition
	// exists.

	TargetStringName string
	// The label or name of the 'TargetString' parameter.
	// Used in error and informational messages.

	TargetStringLength int
	// Actual number of text characters in the entire
	// Target Search String ('TargetString').

	TargetStringLengthName string
	// The label or name of the 'TargetStringLength' parameter.
	// Used in error and informational messages.

	TargetStringStartingSearchIndex int
	// The index in 'TargetString' at which the search
	// operation begins.

	TargetStringStartingSearchIndexName string
	// The label or name of the
	// TargetStringStartingSearchIndex parameter.
	// Used in error and informational messages.

	TargetStringSearchLength int
	// The actual number of characters within the Target
	// Search String that are included in the search
	// operation. This value may be less than the actual
	// length of the Target Search String.

	TargetStringSearchLengthName string
	// The label or name of the TargetStringSearchLength
	// parameter. Used in error and informational
	// messages.

	TargetStringAdjustedSearchLength int
	// The adjusted or corrected Target String Search
	// Length. This value is guaranteed to be equal to or
	// less than the actual Target String Length.

	TargetStringDescription1 string
	// First of two optional description strings
	// describing the Target Search String in the context
	// of the current search operation.

	TargetStringDescription2 string
	// Second of two optional description strings
	// describing the Target Search String in the context
	// of the current search operation.

	FoundFirstNumericDigitInNumStr bool
	// When set to 'true' this signals that the first
	// numeric digit has been identified in the text
	// characters specified by 'TargetString'

	TextCharSearchType CharacterSearchType
	// Optional. An enumeration value signaling the type
	// of text character search algorithm used to conduct
	// this search operation. When set to a valid value,
	// this specification will override the search
	// specification contained in the Test Input
	// Parameters Data Transfer Object.
	//
	// Valid CharSearch Type values are listed as follows:
	//  TextCharSearchType.None() - Invalid Value
	//  TextCharSearchType.LinearTargetStartingIndex() - Default
	//  TextCharSearchType.SingleTargetChar()
	//  TextCharSearchType.LinearEndOfString()

	lock *sync.Mutex
}

func (targetSearchParms *CharSearchTargetInputParametersDto) Empty() {

	if targetSearchParms.lock == nil {
		targetSearchParms.lock = new(sync.Mutex)
	}

	targetSearchParms.lock.Lock()

	targetSearchParms.TargetString = nil

	targetSearchParms.TargetStringName = ""

	targetSearchParms.TargetStringLength = -1

	targetSearchParms.TargetStringLengthName = ""

	targetSearchParms.TargetStringStartingSearchIndex = -1

	targetSearchParms.TargetStringStartingSearchIndexName = ""

	targetSearchParms.TargetStringSearchLength = -2

	targetSearchParms.TargetStringSearchLengthName = ""

	targetSearchParms.TargetStringAdjustedSearchLength = -2

	targetSearchParms.TargetStringDescription1 = ""

	targetSearchParms.TargetStringDescription2 = ""

	targetSearchParms.FoundFirstNumericDigitInNumStr = false

	targetSearchParms.TextCharSearchType = CharSearchType.None()

	targetSearchParms.lock.Unlock()

	targetSearchParms.lock = nil
}

// New - Returns a new, uninitialized instance of
// CharSearchTargetInputParametersDto.
//
func (targetSearchParms CharSearchTargetInputParametersDto) New() CharSearchTargetInputParametersDto {

	if targetSearchParms.lock == nil {
		targetSearchParms.lock = new(sync.Mutex)
	}

	targetSearchParms.lock.Lock()

	defer targetSearchParms.lock.Unlock()

	newEmptyTargetInputParms := CharSearchTargetInputParametersDto{}

	newEmptyTargetInputParms.Empty()

	return newEmptyTargetInputParms
}

// ValidateTargetParameters - Validates the Target Search String
// and related member variables contained in the current instance
// of CharSearchTargetInputParametersDto.
//
func (targetSearchParms *CharSearchTargetInputParametersDto) ValidateTargetParameters(
	errorPrefix interface{}) error {

	if targetSearchParms.lock == nil {
		targetSearchParms.lock = new(sync.Mutex)
	}

	targetSearchParms.lock.Lock()

	defer targetSearchParms.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"CharSearchTargetInputParametersDto."+
			"ValidateTargetSearchString()",
		"")

	if err != nil {

		return err

	}

	if len(targetSearchParms.TargetStringName) == 0 {
		targetSearchParms.TargetStringName =
			"TargetString"
	}

	if targetSearchParms.TargetString == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter '%v' is a nil pointer!\n",
			ePrefix.String(),
			targetSearchParms.TargetStringName)

		return err
	}

	if len(targetSearchParms.TargetStringLengthName) == 0 {
		targetSearchParms.TargetStringLengthName =
			"TargetStringLength"
	}

	targetSearchParms.TargetStringLength =
		len(targetSearchParms.TargetString.CharsArray)

	if targetSearchParms.TargetStringLength == 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter '%v' is invalid!\n"+
			"The rune array encapsulated by '%v' is empty\n"+
			"Length of %v.CharsArray is Zero (0).\n",
			ePrefix.String(),
			targetSearchParms.TargetStringLengthName,
			targetSearchParms.TargetStringName,
			targetSearchParms.TargetStringName)

		return err
	}

	if len(targetSearchParms.TargetStringStartingSearchIndexName) == 0 {
		targetSearchParms.TargetStringStartingSearchIndexName =
			"TargetStringStartingSearchIndex"
	}

	if targetSearchParms.TargetStringStartingSearchIndex < 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter %v is invalid!\n"+
			"%v is less than zero (0)\n"+
			"%v = '%v'\n",
			ePrefix.String(),
			targetSearchParms.TargetStringStartingSearchIndexName,
			targetSearchParms.TargetStringStartingSearchIndexName,
			targetSearchParms.TargetStringStartingSearchIndexName,
			targetSearchParms.TargetStringStartingSearchIndex)

		return err
	}

	if targetSearchParms.TargetStringStartingSearchIndex >=
		targetSearchParms.TargetStringLength {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter %v is invalid!\n"+
			"%v has a value greater than the last\n"+
			"index in '%v.CharsArray'.\n"+
			"Last Index in %v.CharsArray = '%v'\n"+
			"%v = '%v'\n",
			ePrefix.String(),
			targetSearchParms.TargetStringStartingSearchIndexName,
			targetSearchParms.TargetStringStartingSearchIndexName,
			targetSearchParms.TargetStringName,
			targetSearchParms.TargetStringName,
			targetSearchParms.TargetStringLength-1,
			targetSearchParms.TargetStringStartingSearchIndexName,
			targetSearchParms.TargetStringStartingSearchIndex)

		return err
	}

	if len(targetSearchParms.TargetStringSearchLengthName) == 0 {
		targetSearchParms.TargetStringSearchLengthName =
			"TargetStringSearchLength"
	}

	if targetSearchParms.TargetStringSearchLength < -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter %v is invalid!\n"+
			"%v has a value less than minus one (-1)\n"+
			"%v = '%v'\n",
			ePrefix.String(),
			targetSearchParms.TargetStringSearchLengthName,
			targetSearchParms.TargetStringSearchLengthName,
			targetSearchParms.TargetStringName,
			targetSearchParms.TargetStringSearchLength)

		return err
	}

	if targetSearchParms.TargetStringSearchLength == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter %v is invalid!\n"+
			"%v has a value of Zero (0)\n",
			ePrefix.String(),
			targetSearchParms.TargetStringSearchLengthName,
			targetSearchParms.TargetStringSearchLengthName)

	}

	if targetSearchParms.TargetStringSearchLength == -1 {

		targetSearchParms.TargetStringAdjustedSearchLength =
			targetSearchParms.TargetStringLength -
				targetSearchParms.TargetStringStartingSearchIndex
	} else {

		targetSearchParms.TargetStringAdjustedSearchLength =
			targetSearchParms.TargetStringSearchLength

	}

	targetSearchParms.TargetStringAdjustedSearchLength =
		targetSearchParms.TargetStringStartingSearchIndex +
			targetSearchParms.TargetStringAdjustedSearchLength

	if targetSearchParms.TargetStringAdjustedSearchLength >
		targetSearchParms.TargetStringLength {

		targetSearchParms.TargetStringAdjustedSearchLength =
			targetSearchParms.TargetStringLength

	}

	return err
}
func (targetSearchParms *CharSearchTargetInputParametersDto) ValidateCharSearchType(
	errorPrefix interface{}) error {

	if targetSearchParms.lock == nil {
		targetSearchParms.lock = new(sync.Mutex)
	}

	targetSearchParms.lock.Lock()

	defer targetSearchParms.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"CharSearchTargetInputParametersDto."+
			"ValidateCharSearchType()",
		"")

	if err != nil {

		return err

	}

	if !targetSearchParms.TextCharSearchType.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"ERROR: The Character Search Type is invalid!\n"+
			"Character Search Type must be set to one of these\n"+
			"enumeration values:\n"+
			"  CharacterSearchType(0).LinearTargetStartingIndex()\n"+
			"  CharacterSearchType(0).SingleTargetChar()\n"+
			"  CharacterSearchType(0).LinearEndOfString()\n"+
			"The invalid Input Character Search Type is currently\n"+
			"configured as:\n"+
			" Character Search Type   String Name: %v\n"+
			" Character Search Type Integer Value: %v\n",
			ePrefix.String(),
			targetSearchParms.TextCharSearchType.String(),
			targetSearchParms.TextCharSearchType.XValueInt())

	}

	return err
}
