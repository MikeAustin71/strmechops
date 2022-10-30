package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// runeArrayDtoNanobot - Provides helper methods for type
// RuneArrayDto.
type runeArrayDtoNanobot struct {
	lock *sync.Mutex
}

// copyRuneArrayDto - Copies all data from input parameter
// 'sourceRunesDto' to input parameter 'destinationRunesDto'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// The pre-existing data fields for input parameter
// 'destinationRunesDto' will be overwritten and deleted.
//
// NO DATA VALIDATION is performed on input parameter,
// 'sourceRunesDto'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	destinationRunesDto        *RuneArrayDto
//	   - A pointer to an instance of RuneArrayDto. All data
//	     contained in the internal member variables of input
//	     parameter 'sourceRunesDto' will be copied to the member
//	     variables of this input parameter, 'destinationRunesDto'.
//	     If this method completes successfully, all member data
//	     variables encapsulated in 'destinationRunesDto' will be
//	     identical to those contained in input parameter,
//	     'sourceRunesDto'.
//
//	     Be advised that the pre-existing data fields in input
//	     parameter 'destinationRunesDto' will be overwritten and
//	     deleted.
//
//
//	sourceRunesDto             *RuneArrayDto
//	   - A pointer to an instance of RuneArrayDto.
//
//	     All data contained in the member variables of this
//	     RuneArrayDto instance will be copied to corresponding
//	     member variables contained within input parameter
//	     'destinationRunesDto'.
//
//	     The original member variable data values encapsulated in
//	     'sourceRunesDto' will remain unchanged and will NOT be
//	     overwritten or deleted.
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
func (runeDtoNanobot *runeArrayDtoNanobot) copyRuneArrayDto(
	destinationRunesDto *RuneArrayDto,
	sourceRunesDto *RuneArrayDto,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if runeDtoNanobot.lock == nil {
		runeDtoNanobot.lock = new(sync.Mutex)
	}

	runeDtoNanobot.lock.Lock()

	defer runeDtoNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"runeArrayDtoNanobot."+
			"copyRuneArrayDto()",
		"")

	if err != nil {

		return err

	}

	if sourceRunesDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'sourceRunesDto' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if destinationRunesDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'destinationRunesDto' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return err
	}

	new(runeArrayDtoElectron).emptyCharsArray(
		destinationRunesDto)

	lenSourceCharsArray := len(sourceRunesDto.CharsArray)

	if lenSourceCharsArray > 0 {

		destinationRunesDto.CharsArray =
			make([]rune, lenSourceCharsArray)

		for i := 0; i < lenSourceCharsArray; i++ {
			destinationRunesDto.CharsArray[i] =
				sourceRunesDto.CharsArray[i]
		}

	} else {
		// MUST BE
		// lenSourceCharsArray == 0
		destinationRunesDto.CharsArray = nil

	}

	destinationRunesDto.Description1 =
		sourceRunesDto.Description1

	destinationRunesDto.Description2 =
		sourceRunesDto.Description2

	destinationRunesDto.charSearchType =
		sourceRunesDto.charSearchType

	return err
}

func (runeDtoNanobot *runeArrayDtoNanobot) characterSearchExecutor(
	targetInputParms CharSearchTargetInputParametersDto,
	testInputParms CharSearchTestInputParametersDto,
	errPrefDto *ePref.ErrPrefixDto) (
	CharSearchRuneArrayResultsDto,
	error) {

	if runeDtoNanobot.lock == nil {
		runeDtoNanobot.lock = new(sync.Mutex)
	}

	runeDtoNanobot.lock.Lock()

	defer runeDtoNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	searchResults := CharSearchRuneArrayResultsDto{}

	searchResults.Empty()

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"runeArrayDtoNanobot."+
			"characterSearchExecutor()",
		"")

	if err != nil {

		return searchResults, err

	}

	err = testInputParms.ValidateTestParameters(
		ePrefix)

	if err != nil {

		return searchResults, err

	}

	err = targetInputParms.ValidateTargetParameters(
		ePrefix)

	if err != nil {

		return searchResults, err

	}

	runeDtoElectron := runeArrayDtoElectron{}

	if testInputParms.TextCharSearchType == CharSearchType.LinearEndOfString() {

		return runeDtoElectron.linearEndOfStringSearch(
			targetInputParms,
			testInputParms,
			ePrefix)

	}

	if testInputParms.TextCharSearchType == CharSearchType.LinearTargetStartingIndex() {

		return runeDtoElectron.linearTargetStartingIndexSearch(
			targetInputParms,
			testInputParms,
			ePrefix)
	}

	if testInputParms.TextCharSearchType == CharSearchType.SingleTargetChar() {

		return runeDtoElectron.singleCharacterSearch(
			targetInputParms,
			testInputParms,
			ePrefix)

	}

	return searchResults, err
}
