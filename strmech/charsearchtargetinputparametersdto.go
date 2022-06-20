package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type CharSearchTargetInputParametersDto struct {
	TargetString *RuneArrayDto // A pointer to the RuneArrayDto containing
	//                            the Target Search String text characters.

	TargetStringName string // The label or name of the TargetString
	//                               parameter.

	TargetStringLength int // Actual Full Length of the Target Search
	//                              String.

	TargetStringLengthName string // The label or name of the
	//                                     TargetStringLength parameter.

	TargetStringStartingSearchIndex int // The index at which the search
	//                               	   operation commenced.

	TargetStringStartingSearchIndexName string // The label or name of the
	//                                            TargetStringStartingSearchIndex
	//                                            parameter.

	TargetStringSearchLength int // The actual number of characters within
	//                              Target Search String that are included
	//                              in the search specification.

	TargetStringSearchLengthName string // The label or name of the
	//                                     TargetStringSearchLength parameter.

	TargetStringAdjustedSearchLength int // The adjusted Target String Search Length
	//                                   guaranteed to be equal to or less than
	//                                   the actual Target String Length.

	TargetStringDescription1 string // First optional description string
	//                                  describing the Target Search String
	//                                  used in this search

	TargetStringDescription2 string // Second Optional description string
	//                                  describing the Target Search String
	//                                  used in this search

	FoundFirstNumericDigitInNumStr bool // When set to 'true' this signals
	//                                     that the first numeric digit has
	//                                     been identified in a string of text
	//                                     characters.

	TextCharSearchType CharacterSearchType // An enumeration value signaling
	//                                the type of search algorithm which
	//                                was used to conduct this search.
	//                                 TextCharSearchType.None()
	//                                 TextCharSearchType.LinearTargetStartingIndex() - Default
	//                                 TextCharSearchType.SingleTargetChar()
	//                                 TextCharSearchType.LinearEndOfString()

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
