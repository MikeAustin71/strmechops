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

//	addStrings
//
//	Receives a series of strings as a variadic argument
//	and proceeds to add them to the rune array contained
//	in an instance of RuneArrayDto passed as an input
//	parameter.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//
//	runeArrayDto				*RuneArrayDto
//
//		A pointer to an instance of RuneArrayDto. The
//		rune array contained in this RuneArrayDto
//		instance will be modified. A single rune text
//		character will either be added to the beginning,
//		or the end, of this rune array:
//
//			runeArrayDto.CharsArray
//
//	addTrailingChar				bool
//
//		This parameter determines whether 'charToAdd'
//		will be added to the existing rune array,
//		'runeArrayDto.CharsArray', as a trailing
//		character or as a leading character.
//
//		If 'addTrailingChars' is set to 'true',
//		'charsToAdd' will be added to the end of the
//		existing rune array as a trailing character.
//
//		If 'addTrailingChars' is set to 'false',
//		'charToAdd' will be added to the beginning of
//		the existing rune array as leading character.
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
func (runeDtoNanobot *runeArrayDtoNanobot) addStrings(
	runeArrayDto *RuneArrayDto,
	addTrailingChars bool,
	errPrefDto *ePref.ErrPrefixDto,
	strings ...string) error {

	if runeDtoNanobot.lock == nil {
		runeDtoNanobot.lock = new(sync.Mutex)
	}

	runeDtoNanobot.lock.Lock()

	defer runeDtoNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"runeArrayDtoNanobot."+
			"addStrings()",
		"")

	if err != nil {

		return err

	}

	if runeArrayDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'runeArrayDto' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return err
	}

	for _, charStr := range strings {

		err = new(runeArrayDtoAtom).addRunes(
			runeArrayDto,
			[]rune(charStr),
			addTrailingChars,
			ePrefix.XCpy(
				"runeArrayDto<-"))

		if err != nil {
			return err
		}
	}

	return err
}

//	copyRuneArrayDto
//
//	Copies all data from input parameter 'sourceRunesDto'
//	to input parameter 'destinationRunesDto'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	The pre-existing data fields for input parameter
//	'destinationRunesDto' will be overwritten and
//	deleted.
//
//	NO DATA VALIDATION is performed on input parameter,
//	'sourceRunesDto'.
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

// characterSearchExecutor
//
// This high level function analyzes a search request
// and assigns the search operation to the appropriate
// low-level search method.
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
