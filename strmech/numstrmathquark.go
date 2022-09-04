package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/rand"
	"sync"
	"time"
)

type numStrMathQuark struct {
	lock *sync.Mutex
}

// extendRunes - Receives an array of runes and proceeds
// to extend the length of that array using one or more
// 'fill' characters passed as input parameters.
//
// The 'fill' characters will be appended to the beginning
// or end of the original rune array depending on the
// value of input parameter, 'extendRunesRight'.
//
// If 'extendRunesRight' is 'true', the fill characters
// will be appended to the end of the original rune
// array.
//
// If 'extendRunesRight' is 'false', the fill characters
// will be appended to the beginning of the original rune
// array.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	incomingCharsToExtend			*RuneArrayDto
//
//		A pointer to an instance of RuneArrayDto containing the
//		original rune array which be extended with fill
//		characters.
//
//		If 'extendedOutputChars' points to an instance of
//		RuneArrayDto other than the 'incomingCharsToExtend'
//		RuneArrayDto, 'incomingCharsToExtend' will NOT be
//		modified.
//
//	extendedOutputChars 			*RuneArrayDto
//
//		A pointer to an instance of RuneArrayDto. This instance
//		will contain the final extended version of the
//		'incomingCharsToExtend' rune array.
//
//		If 'extendedOutputChars' points to an instance of
//		RuneArrayDto other than the 'incomingCharsToExtend'
//		RuneArrayDto, 'incomingCharsToExtend' will NOT be
//		modified.
//
//	fillChar						rune
//
//		This rune contains the fill character which will be used
//		to extend the original rune array
//		('incomingCharsToExtend').
//
//	newCharArrayLength				int
//
//		The new length of the extended rune array.
//
//		If this value is equal to or less than the length of
//		the original rune array ('incomingCharsToExtend'),
//		this method will immediately exit and no error will
//		be returned.
//
//	extendRunesRight				bool
//
//		If this parameter is set to 'true', the fill characters
//		will be added to the end of the original rune array,
//		'incomingCharsToExtend'.
//
//		If this parameter is set to 'false', the fill characters
//		will be added to the beginning of the original rune
//		array, 'incomingCharsToExtend'.
//
//	errPrefDto						*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string which
//		is included in all returned error messages. Usually,
//		it contains the name of the calling method or methods
//		listed as a function chain.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref' software
//		package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, this returned
//		error Type is set equal to 'nil'. If errors are
//		encountered during	processing, the returned error
//		Type will encapsulate an error	message.
//
//		If an error message is returned, the text value for
//		input parameter 'errPrefDto' (error prefix) will be
//		prefixed or attached at the beginning of the error
//		message.
func (nStrMathQuark *numStrMathQuark) extendRunes(
	incomingCharsToExtend *RuneArrayDto,
	extendedOutputChars *RuneArrayDto,
	fillChar rune,
	newCharArrayLength int,
	extendRunesRight bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	if nStrMathQuark.lock == nil {
		nStrMathQuark.lock = new(sync.Mutex)
	}

	nStrMathQuark.lock.Lock()

	defer nStrMathQuark.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrMathQuark."+
			"extendRunesRight()",
		"")

	if err != nil {
		return err
	}

	if incomingCharsToExtend == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'incomingCharsToExtend' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if extendedOutputChars == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'extendedOutputChars' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return err
	}

	charArrayLen :=
		incomingCharsToExtend.GetRuneArrayLength()

	if charArrayLen >= newCharArrayLength {
		// Nothing to do.

		if charArrayLen !=
			extendedOutputChars.GetRuneArrayLength() {
			// These are not equal rune arrays

			err = extendedOutputChars.CopyIn(
				incomingCharsToExtend,
				ePrefix.XCpy(
					"extendedOutputChars<-"+
						"incomingCharsToExtend"))

			return err
		}

		for i := 0; i < charArrayLen; i++ {
			extendedOutputChars.CharsArray[i] =
				incomingCharsToExtend.CharsArray[i]
		}

		return err
	}

	lengthDelta := newCharArrayLength - charArrayLen

	newFillArray := make([]rune, lengthDelta)

	for i := 0; i < lengthDelta; i++ {
		newFillArray[i] = fillChar
	}

	if extendRunesRight {

		extendedOutputChars.CharsArray =
			append(
				incomingCharsToExtend.CharsArray,
				newFillArray...)

	} else {

		extendedOutputChars.CharsArray =
			append(
				newFillArray,
				incomingCharsToExtend.CharsArray...)

	}

	return err
}

// randomInt
//
// Returns a random integer number created with the psuedo
// random number generator provided by the math/rand
// package
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	NONE
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	NONE
func (nStrMathQuark *numStrMathQuark) randomInt() int {

	if nStrMathQuark.lock == nil {
		nStrMathQuark.lock = new(sync.Mutex)
	}

	nStrMathQuark.lock.Lock()

	defer nStrMathQuark.lock.Unlock()

	rand.Seed(time.Now().UnixNano())

	intNums := rand.Perm(30)

	selection := rand.Intn(29)

	intNums = rand.Perm(30)

	selection = rand.Intn(29)

	return intNums[selection]
}
