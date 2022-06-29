package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// charSearchResultsDtoNanobot - Provides helper methods for type
// CharSearchResultsDto.
//
type charSearchResultsDtoNanobot struct {
	lock *sync.Mutex
}

// copyIn - Copies all data from input parameter
// 'incomingSearchResultsDto' to input parameter
// 'targetSearchResultsDto'. Both instances are of type
// CharSearchResultsDto.
//
// IMPORTANT
//
// ----------------------------------------------------------------
//
// Be advised that the data fields in 'targetSearchResultsDto'
// will be overwritten.
//
// Also, NO data validation is performed on
// 'incomingSearchResultsDto'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  targetSearchResultsDto     *CharSearchResultsDto
//     - A pointer to a CharSearchResultsDto instance. All the
//       member variable data fields in this object will be
//       replaced by data values copied from input parameter
//       'incomingSearchResultsDto'.
//
//       'targetSearchResultsDto' is the destination for this
//       copy operation.
//
//
//  incomingSearchResultsDto          *CharSearchResultsDto
//     - A pointer to another CharSearchResultsDto instance. All
//       the member variable data values from this object will be
//       copied to corresponding member variables in
//       'targetSearchResultsDto'.
//
//       'incomingSearchResultsDto' is the source for this copy
//       operation.
//
//       If 'incomingSearchResultsDto' is determined to be invalid,
//       an error will be returned.
//
//
//  errPrefDto          *ePref.ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods listed
//       as a function chain.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
//
//       Type ErrPrefixDto is included in the 'errpref' software
//       package, "github.com/MikeAustin71/errpref".
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (searchResultsDtoNanobot *charSearchResultsDtoNanobot) copyIn(
	targetSearchResultsDto *CharSearchResultsDto,
	incomingSearchResultsDto *CharSearchResultsDto,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if searchResultsDtoNanobot.lock == nil {
		searchResultsDtoNanobot.lock = new(sync.Mutex)
	}

	searchResultsDtoNanobot.lock.Lock()

	defer searchResultsDtoNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchTestInputParametersDtoNanobot."+
			"copyIn()",
		"")

	if err != nil {

		return err

	}

	if targetSearchResultsDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'targetSearchResultsDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if incomingSearchResultsDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'incomingSearchResultsDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	charSearchResultsDtoAtom{}.ptr().empty(
		targetSearchResultsDto)

	targetSearchResultsDto.FoundSearchTarget =
		incomingSearchResultsDto.FoundSearchTarget

	targetSearchResultsDto.FoundSearchTargetOnPreviousSearch =
		incomingSearchResultsDto.FoundSearchTargetOnPreviousSearch

	targetSearchResultsDto.FoundFirstNumericDigitInNumStr =
		incomingSearchResultsDto.FoundFirstNumericDigitInNumStr

	targetSearchResultsDto.TargetInputParametersName =
		incomingSearchResultsDto.TargetInputParametersName

	targetSearchResultsDto.TargetStringLength =
		incomingSearchResultsDto.TargetStringLength

	targetSearchResultsDto.TargetStringSearchLength =
		incomingSearchResultsDto.TargetStringSearchLength

	targetSearchResultsDto.TargetStringStartingSearchIndex =
		incomingSearchResultsDto.TargetStringStartingSearchIndex

	targetSearchResultsDto.TargetStringFirstFoundIndex =
		incomingSearchResultsDto.TargetStringFirstFoundIndex

	targetSearchResultsDto.TargetStringLastFoundIndex =
		incomingSearchResultsDto.TargetStringLastFoundIndex

	targetSearchResultsDto.TargetStringLastSearchIndex =
		incomingSearchResultsDto.TargetStringLastSearchIndex

	targetSearchResultsDto.TargetStringDescription1 =
		incomingSearchResultsDto.TargetStringDescription1

	targetSearchResultsDto.TargetStringDescription2 =
		incomingSearchResultsDto.TargetStringDescription2

	targetSearchResultsDto.TestInputParametersName =
		incomingSearchResultsDto.TestInputParametersName

	targetSearchResultsDto.TestStringLength =
		incomingSearchResultsDto.TestStringLength

	targetSearchResultsDto.TestStringStartingIndex =
		incomingSearchResultsDto.TestStringStartingIndex

	targetSearchResultsDto.TestStringFirstFoundIndex =
		incomingSearchResultsDto.TestStringFirstFoundIndex

	targetSearchResultsDto.TestStringLastFoundIndex =
		incomingSearchResultsDto.TestStringLastFoundIndex

	targetSearchResultsDto.TestStringDescription1 =
		incomingSearchResultsDto.TestStringDescription1

	targetSearchResultsDto.TestStringDescription2 =
		incomingSearchResultsDto.TestStringDescription2

	targetSearchResultsDto.CollectionTestObjIndex =
		incomingSearchResultsDto.CollectionTestObjIndex

	if incomingSearchResultsDto.ReplacementString != nil {

		targetSearchResultsDto.ReplacementString =
			&RuneArrayDto{}

		err = targetSearchResultsDto.ReplacementString.CopyIn(
			incomingSearchResultsDto.ReplacementString,
			ePrefix.XCpy(
				"targetSearchResultsDto.ReplacementString"+
					"<-incomingSearchResultsDto.ReplacementString"))
	}

	targetSearchResultsDto.NumValueType =
		incomingSearchResultsDto.NumValueType

	targetSearchResultsDto.NumStrFormatType =
		incomingSearchResultsDto.NumStrFormatType

	targetSearchResultsDto.NumSymLocation =
		incomingSearchResultsDto.NumSymLocation

	targetSearchResultsDto.NumSymbolClass =
		incomingSearchResultsDto.NumSymbolClass

	targetSearchResultsDto.NumSignValue =
		incomingSearchResultsDto.NumSignValue

	targetSearchResultsDto.PrimaryNumSignPosition =
		incomingSearchResultsDto.PrimaryNumSignPosition

	targetSearchResultsDto.SecondaryNumSignPosition =
		incomingSearchResultsDto.SecondaryNumSignPosition

	targetSearchResultsDto.TextCharSearchType =
		incomingSearchResultsDto.TextCharSearchType

	return err
}

// copyOut - Returns a deep copy of the input parameter
// 'searchResultsDto', a pointer to an instance of
// CharSearchResultsDto.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// NO validation is performed on 'searchResultsDto'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  searchResultsDto           *CharSearchResultsDto
//     - A pointer to an instance of CharSearchResultsDto. A
//       deep copy of the internal member variables will be created
//       and returned in a new instance of CharSearchResultsDto.
//
//       No data validation is performed on 'searchResultsDto'.
//
//
//  errPrefDto                 *ePref.ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods listed
//       as a function chain.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
//
//       Type ErrPrefixDto is included in the 'errpref' software
//       package, "github.com/MikeAustin71/errpref".
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  copySearchResultsDto       CharSearchResultsDto
//     - If this method completes successfully, a deep copy of
//       input parameter 'searchResultsDto' will be created and
//       returned in a new instance of
//       CharSearchResultsDto.
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
func (searchResultsDtoNanobot *charSearchResultsDtoNanobot) copyOut(
	searchResultsDto *CharSearchResultsDto,
	errPrefDto *ePref.ErrPrefixDto) (
	copySearchResultsDto CharSearchResultsDto,
	err error) {

	if searchResultsDtoNanobot.lock == nil {
		searchResultsDtoNanobot.lock = new(sync.Mutex)
	}

	searchResultsDtoNanobot.lock.Lock()

	defer searchResultsDtoNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchTestInputParametersDtoNanobot."+
			"copyOut()",
		"")

	if err != nil {

		return copySearchResultsDto, err

	}

	if searchResultsDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'targetSearchResultsDto' is a nil pointer!\n",
			ePrefix.String())

		return copySearchResultsDto, err
	}

	charSearchResultsDtoAtom{}.ptr().empty(
		&copySearchResultsDto)

	copySearchResultsDto.FoundSearchTarget =
		searchResultsDto.FoundSearchTarget

	copySearchResultsDto.FoundSearchTargetOnPreviousSearch =
		searchResultsDto.FoundSearchTargetOnPreviousSearch

	copySearchResultsDto.FoundFirstNumericDigitInNumStr =
		searchResultsDto.FoundFirstNumericDigitInNumStr

	copySearchResultsDto.TargetInputParametersName =
		searchResultsDto.TargetInputParametersName

	copySearchResultsDto.TargetStringLength =
		searchResultsDto.TargetStringLength

	copySearchResultsDto.TargetStringSearchLength =
		searchResultsDto.TargetStringSearchLength

	copySearchResultsDto.TargetStringStartingSearchIndex =
		searchResultsDto.TargetStringStartingSearchIndex

	copySearchResultsDto.TargetStringFirstFoundIndex =
		searchResultsDto.TargetStringFirstFoundIndex

	copySearchResultsDto.TargetStringLastFoundIndex =
		searchResultsDto.TargetStringLastFoundIndex

	copySearchResultsDto.TargetStringLastSearchIndex =
		searchResultsDto.TargetStringLastSearchIndex

	copySearchResultsDto.TargetStringDescription1 =
		searchResultsDto.TargetStringDescription1

	copySearchResultsDto.TargetStringDescription2 =
		searchResultsDto.TargetStringDescription2

	copySearchResultsDto.TestInputParametersName =
		searchResultsDto.TestInputParametersName

	copySearchResultsDto.TestStringLength =
		searchResultsDto.TestStringLength

	copySearchResultsDto.TestStringStartingIndex =
		searchResultsDto.TestStringStartingIndex

	copySearchResultsDto.TestStringFirstFoundIndex =
		searchResultsDto.TestStringFirstFoundIndex

	copySearchResultsDto.TestStringLastFoundIndex =
		searchResultsDto.TestStringLastFoundIndex

	copySearchResultsDto.TestStringDescription1 =
		searchResultsDto.TestStringDescription1

	copySearchResultsDto.TestStringDescription2 =
		searchResultsDto.TestStringDescription2

	copySearchResultsDto.CollectionTestObjIndex =
		searchResultsDto.CollectionTestObjIndex

	copySearchResultsDto.NumValueType =
		searchResultsDto.NumValueType

	copySearchResultsDto.NumStrFormatType =
		searchResultsDto.NumStrFormatType

	copySearchResultsDto.NumSymLocation =
		searchResultsDto.NumSymLocation

	copySearchResultsDto.NumSymbolClass =
		searchResultsDto.NumSymbolClass

	copySearchResultsDto.NumSignValue =
		searchResultsDto.NumSignValue

	copySearchResultsDto.PrimaryNumSignPosition =
		searchResultsDto.PrimaryNumSignPosition

	copySearchResultsDto.SecondaryNumSignPosition =
		searchResultsDto.SecondaryNumSignPosition

	copySearchResultsDto.TextCharSearchType =
		searchResultsDto.TextCharSearchType

	return copySearchResultsDto, err
}

// ptr - Returns a pointer to a new instance of
// charSearchResultsDtoNanobot.
//
func (searchResultsDtoNanobot charSearchResultsDtoNanobot) ptr() *charSearchResultsDtoNanobot {

	if searchResultsDtoNanobot.lock == nil {
		searchResultsDtoNanobot.lock = new(sync.Mutex)
	}

	searchResultsDtoNanobot.lock.Lock()

	defer searchResultsDtoNanobot.lock.Unlock()

	return &charSearchResultsDtoNanobot{
		lock: new(sync.Mutex),
	}
}
