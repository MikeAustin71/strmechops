package strmech

import (
	"crypto/rand"
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"sync"
)

type numStrMathQuark struct {
	lock *sync.Mutex
}

//	extendRunes
//
//	Receives an array of runes and proceeds to extend the
//	length of that array using one or more 'fill'
//	characters passed as input parameters.
//
//	The 'fill' characters will be appended to the
//	beginning or end of the original rune array depending
//	on the value of input parameter, 'extendRunesRight'.
//
//	If 'extendRunesRight' is 'true', the fill characters
//	will be appended to the end of the original rune
//	array.
//
//	If 'extendRunesRight' is 'false', the fill characters
//	will be appended to the beginning of the original
//	rune array.
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
		len(incomingCharsToExtend.CharsArray)

	if charArrayLen >= newCharArrayLength {
		// Nothing to do.

		if charArrayLen !=
			len(extendedOutputChars.CharsArray) {
			// These are not equal rune arrays

			err = new(runeArrayDtoNanobot).copyRuneArrayDto(
				extendedOutputChars,
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

//	randomInt64
//
//	Returns a random integer number created with the
//	pseudo random number generator provided by the
//	"crypto/rand" package.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://golangdocs.com/the-crypto-rand-package-in-golang
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	intervalUpperBound			int64
//
//		This method utilizes a non-negative pseudo-random
//		number in the range [0,intervalUpperBound].
//
//		If this value is less than one (+1), an error will
//		be returned.
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
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	int64
//
//		This parameter returns a pseudo-random number
//		between zero (0) and 'intervalUpperBound'.
//
//	error
//
//		If this method completes successfully, this
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (nStrMathQuark *numStrMathQuark) randomInt64(
	intervalUpperBound int64,
	errPrefDto *ePref.ErrPrefixDto) (
	int64,
	error) {

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
			"randomInt64()",
		"")

	if err != nil {

		return -1, err
	}

	if intervalUpperBound < 1 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'intervalUpperBound' is invalid!\n"+
			"'intervalUpperBound' has a value less than one (+1).\n"+
			"intervalUpperBound = '%v'\n",
			ePrefix.String(),
			intervalUpperBound)

		return -1, err
	}

	maxNumLimit := big.NewInt(intervalUpperBound)

	var bigRandomNum *big.Int

	bigRandomNum,
		err = rand.Int(rand.Reader, maxNumLimit)

	if err != nil {
		err = fmt.Errorf("%v\n"+
			"Error returned from big.NewInt(intervalUpperBound)\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			err.Error())

		return -1, err

	}

	return bigRandomNum.Int64(), err
}
