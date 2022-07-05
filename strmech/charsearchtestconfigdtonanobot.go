package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// charSearchTestConfigDtoNanobot - Provides helper methods for type
// CharSearchTestConfigDto.
//
type charSearchTestConfigDtoNanobot struct {
	lock *sync.Mutex
}

// copyIn - Copies all data from input parameter
// 'sourceTestCfgDto' to input parameter
// 'destinationTestCfgDto'. Both instances are of type
// CharSearchTestConfigDto.
//
// IMPORTANT
//
// -----------------------------------------------------------------
//
// Be advised that the data fields in 'destinationTestCfgDto'
// will be deleted and overwritten.
//
// Also, NO validation is performed on 'sourceTestCfgDto'.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  destinationTestCfgDto     *CharSearchTestConfigDto
//     - A pointer to a CharSearchTestConfigDto instance.
//       All the member variable data fields in this object will be
//       replaced by data values copied from input parameter
//       'sourceTestCfgDto'.
//
//       'destinationTestCfgDto' is the destination for this
//       copy operation.
//
//
//  sourceTestCfgDto          *CharSearchTestConfigDto
//     - A pointer to another CharSearchTestConfigDto
//       instance. All the member variable data values from this
//       object will be copied to corresponding member variables in
//       'destinationTestCfgDto'.
//
//       'sourceTestCfgDto' is the source for this copy
//       operation.
//
//       No data validation is performed on
//       'sourceTestCfgDto'.
//
//
//  errPrefDto          *ePref.ErrPrefixDto
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
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (searchTestConfigNanobot *charSearchTestConfigDtoNanobot) copyIn(
	destinationTestCfgDto *CharSearchTestConfigDto,
	sourceTestCfgDto *CharSearchTestConfigDto,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if searchTestConfigNanobot.lock == nil {
		searchTestConfigNanobot.lock = new(sync.Mutex)
	}

	searchTestConfigNanobot.lock.Lock()

	defer searchTestConfigNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchTargetInputParametersDtoNanobot."+
			"copyIn()",
		"")

	if err != nil {

		return err

	}

	if destinationTestCfgDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'destinationTestCfgDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if sourceTestCfgDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'sourceTestCfgDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	charSearchTestConfigDtoAtom{}.ptr().empty(
		destinationTestCfgDto)

	destinationTestCfgDto.TestInputParametersName =
		sourceTestCfgDto.TestInputParametersName

	destinationTestCfgDto.TestStringName =
		sourceTestCfgDto.TestStringName

	destinationTestCfgDto.TestStringLengthName =
		sourceTestCfgDto.TestStringLengthName

	destinationTestCfgDto.TestStringStartingIndex =
		sourceTestCfgDto.TestStringStartingIndex

	destinationTestCfgDto.TestStringStartingIndexName =
		sourceTestCfgDto.TestStringStartingIndexName

	destinationTestCfgDto.TestStringDescription1 =
		sourceTestCfgDto.TestStringDescription1

	destinationTestCfgDto.TestStringDescription2 =
		sourceTestCfgDto.TestStringDescription2

	destinationTestCfgDto.CollectionTestObjIndex =
		sourceTestCfgDto.CollectionTestObjIndex

	destinationTestCfgDto.NumValueType =
		sourceTestCfgDto.NumValueType

	destinationTestCfgDto.NumStrFormatType =
		sourceTestCfgDto.NumStrFormatType

	destinationTestCfgDto.NumSymbolLocation =
		sourceTestCfgDto.NumSymbolLocation

	destinationTestCfgDto.NumSymbolClass =
		sourceTestCfgDto.NumSymbolClass

	destinationTestCfgDto.NumSignValue =
		sourceTestCfgDto.NumSignValue

	destinationTestCfgDto.PrimaryNumSignPosition =
		sourceTestCfgDto.PrimaryNumSignPosition

	destinationTestCfgDto.SecondaryNumSignPosition =
		sourceTestCfgDto.SecondaryNumSignPosition

	destinationTestCfgDto.TextCharSearchType =
		sourceTestCfgDto.TextCharSearchType

	return err
}

// copyOut - Returns a deep copy of the input parameter
// 'searchTestCfgDto', a pointer to an instance of
// CharSearchTestConfigDto.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
//
// NO validation is performed on 'searchTestCfgDto'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  searchTestCfgDto                *CharSearchTestConfigDto
//     - A pointer to an instance of CharSearchTestConfigDto. A
//       deep copy of the internal member variables will be created
//       and returned in a new instance of CharSearchTestConfigDto.
//
//       No data validation is performed on 'searchTestCfgDto'.
//
//
//  errPrefDto                      *ePref.ErrPrefixDto
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
//  deepCopySearchTestCfgDto        CharSearchTestConfigDto
//     - If this method completes successfully, a deep copy of
//       input parameter 'searchTestCfgDto' will be created and
//       returned in a new instance of CharSearchTestConfigDto.
//
//
//  err                             error
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (searchTestConfigNanobot *charSearchTestConfigDtoNanobot) copyOut(
	searchTestCfgDto *CharSearchTestConfigDto,
	errPrefDto *ePref.ErrPrefixDto) (
	deepCopySearchTestCfgDto CharSearchTestConfigDto,
	err error) {

	if searchTestConfigNanobot.lock == nil {
		searchTestConfigNanobot.lock = new(sync.Mutex)
	}

	searchTestConfigNanobot.lock.Lock()

	defer searchTestConfigNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchTargetInputParametersDtoNanobot."+
			"copyOut()",
		"")

	if err != nil {

		return deepCopySearchTestCfgDto, err

	}

	if searchTestCfgDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'searchTestCfgDto' is a nil pointer!\n",
			ePrefix.String())

		return deepCopySearchTestCfgDto, err
	}

	charSearchTestConfigDtoAtom{}.ptr().empty(
		&deepCopySearchTestCfgDto)

	deepCopySearchTestCfgDto.TestInputParametersName =
		searchTestCfgDto.TestInputParametersName

	deepCopySearchTestCfgDto.TestStringName =
		searchTestCfgDto.TestStringName

	deepCopySearchTestCfgDto.TestStringLengthName =
		searchTestCfgDto.TestStringLengthName

	deepCopySearchTestCfgDto.TestStringStartingIndex =
		searchTestCfgDto.TestStringStartingIndex

	deepCopySearchTestCfgDto.TestStringStartingIndexName =
		searchTestCfgDto.TestStringStartingIndexName

	deepCopySearchTestCfgDto.TestStringDescription1 =
		searchTestCfgDto.TestStringDescription1

	deepCopySearchTestCfgDto.TestStringDescription2 =
		searchTestCfgDto.TestStringDescription2

	deepCopySearchTestCfgDto.CollectionTestObjIndex =
		searchTestCfgDto.CollectionTestObjIndex

	deepCopySearchTestCfgDto.NumValueType =
		searchTestCfgDto.NumValueType

	deepCopySearchTestCfgDto.NumStrFormatType =
		searchTestCfgDto.NumStrFormatType

	deepCopySearchTestCfgDto.NumSymbolLocation =
		searchTestCfgDto.NumSymbolLocation

	deepCopySearchTestCfgDto.NumSymbolClass =
		searchTestCfgDto.NumSymbolClass

	deepCopySearchTestCfgDto.NumSignValue =
		searchTestCfgDto.NumSignValue

	deepCopySearchTestCfgDto.PrimaryNumSignPosition =
		searchTestCfgDto.PrimaryNumSignPosition

	deepCopySearchTestCfgDto.SecondaryNumSignPosition =
		searchTestCfgDto.SecondaryNumSignPosition

	deepCopySearchTestCfgDto.TextCharSearchType =
		searchTestCfgDto.TextCharSearchType

	return deepCopySearchTestCfgDto, err
}

// ptr - Returns a pointer to a new instance of
// charSearchTestConfigDtoNanobot.
//
func (searchTestConfigNanobot charSearchTestConfigDtoNanobot) ptr() *charSearchTestConfigDtoNanobot {

	if searchTestConfigNanobot.lock == nil {
		searchTestConfigNanobot.lock = new(sync.Mutex)
	}

	searchTestConfigNanobot.lock.Lock()

	defer searchTestConfigNanobot.lock.Unlock()

	return &charSearchTestConfigDtoNanobot{
		lock: new(sync.Mutex),
	}
}
