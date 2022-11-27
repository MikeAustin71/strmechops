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

//	pureNumStrToRunes
//
//	Receives a pure number string and proceeds to return
//	the extracted numeric value as integer and fractional
//	rune arrays which taken together constitute the
//	extracted numeric value.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	pureNumberStr				string
//
//		A pure number string consists entirely of numeric
//		digit text characters.
//
//		In the case of negative numeric values, a leading
//		or trailing minus sign ('-') must be present. A
//		valid minus sign will therefore classify the
//		numeric value as a negative number.
//
//		Leading minus signs ('-') are prefixed at the
//		beginning of the pure number string.
//
//		Trailing minus signs ('-') are located at the end
//		of the pure number string.
//
//		The input parameter, 'leadingMinusSign' will
//		determine whether a leading or a tailing minus
//		sign is expected.
//
//		Location of a valid trailing minus sign will
//		immediately terminate the search for numeric
//		characters.
//
//	intDigits					*RuneArrayDto
//
//		A pointer to an instance of RuneArrayDto. The
//		integer numeric digits extracted from
//		'pureNumberStr' will be stored as text
//		characters in the rune array encapsulated by
//		this RuneArrayDto object.
//
//		The positive or negative number sign for the
//		extracted integer digits, can be determined by
//		examining the statistics returned by parameter
//		'numberStats' (numberStats.NumberSign).
//
//	fracDigits					*RuneArrayDto
//
//		A pointer to an instance of RuneArrayDto. The
//		fractional numeric digits extracted from
//		'pureNumberStr' will be stored as text
//		characters in the rune array encapsulated by
//		this RuneArrayDto object.
//
//		The positive or negative number sign for the
//		extracted integer digits, can be determined by
//		examining the statistics returned by parameter
//		'numberStats' (numberStats.NumberSign).
//
//	decSeparatorChars		*RuneArrayDto
//
//		A pointer to an instance of RuneArrayDto.
//
//		This RuneArrayDto encapsulates the rune array
//		containing the radix point or decimal separator
//		character or characters which will be used to
//		separate integer and fractional digits within
//		the pure number string input parameter,
//		'pureNumberStr'.
//
//		In the US, UK, Australia and most of Canada, the
//		decimal separator is the period character ('.')
//		known as the decimal point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//	leadingMinusSign			bool
//
//		When set to 'true', the pure number string
//		parsing algorithm will search for a leading minus
//		sign ('-') at the beginning of the number string.
//
//		Leading minus signs represent the standard means
//		for designating negative numeric values in the
//		US, UK, Australia, most of Canada and many other
//		countries.
//
//		Example Leading Minus Sign:
//			"-123.456" or "- 123.456"
//
//		When set to 'false', the pure number string
//		parsing algorithm will search for trailing minus
//		signs ('-') located at the end of the pure number
//		string.
//
//		Trailing minus signs represent the standard for
//		France, Germany and many countries in the
//		European Union.
//
//		NOTE: Identification of a trailing minus sign in
//		the pure number string ('pureNumberStr') will
//		immediately terminate the search for numeric
//		characters.
//
//		Example Trailing Number Symbols:
//			"123.456-" or "123.456 -"
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
//	numberStats					NumberStrStatsDto
//
//		This data transfer object will return critical
//		statistics on the numeric value represented
//		by the integer and fractional digits extracted
//		from 'pureNumberStr' and stored in the
//		'intDigits' and 'fracDigits' RuneArrayDto
//		objects.
//
//		type NumberStrStatsDto struct {
//
//		NumOfIntegerDigits					uint64
//
//			The total number of integer digits to the
//			left of the radix point or, decimal point, in
//			the subject numeric value.
//
//		NumOfSignificantIntegerDigits		uint64
//
//			The number of nonzero integer digits to the
//			left of the radix point or, decimal point, in
//			the subject numeric value.
//
//		NumOfFractionalDigits				uint64
//
//			The total number of fractional digits to the
//			right of the radix point or, decimal point,
//			in the subject numeric value.
//
//		NumOfSignificantFractionalDigits	uint64
//
//			The number of nonzero fractional digits to
//			the right of the radix point or, decimal
//			point, in the subject numeric value.
//
//		NumberValueType 					NumericValueType
//
//			This enumeration value specifies whether the
//			subject numeric value is classified either as
//			an integer or a floating point number.
//
//			Possible enumeration values are listed as
//			follows:
//				NumValType.None()
//				NumValType.FloatingPoint()
//				NumValType.Integer()
//
//		NumberSign							NumericSignValueType
//
//			An enumeration specifying the number sign
//			associated with the numeric value. Possible
//			values are listed as follows:
//				NumSignVal.None()		= Invalid Value
//				NumSignVal.Negative()	= -1
//				NumSignVal.Zero()		=  0
//				NumSignVal.Positive()	=  1
//
//		IsZeroValue							bool
//
//			If 'true', the subject numeric value is equal
//			to zero ('0').
//
//			If 'false', the subject numeric value is
//			greater than or less than zero ('0').
//		}
//
//	err							error
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
func (nStrMathQuark *numStrMathQuark) pureNumStrToRunes(
	pureNumberStr string,
	intDigits *RuneArrayDto,
	fracDigits *RuneArrayDto,
	decSeparatorChars *RuneArrayDto,
	leadingMinusSign bool,
	errPrefDto *ePref.ErrPrefixDto) (
	numberStats NumberStrStatsDto,
	err error) {

	if nStrMathQuark.lock == nil {
		nStrMathQuark.lock = new(sync.Mutex)
	}

	nStrMathQuark.lock.Lock()

	defer nStrMathQuark.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	numberStats.NumberSign = NumSignVal.Zero()

	numberStats.IsZeroValue = true

	numberStats.NumberValueType = NumValType.Integer()

	numberStats.NumOfIntegerDigits = 0

	numberStats.NumOfSignificantIntegerDigits = 0

	numberStats.NumOfFractionalDigits = 0

	numberStats.NumOfSignificantFractionalDigits = 0

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"mathIntHelperAtom."+
			"pureNumStrToRunes()",
		"")

	if err != nil {

		return numberStats, err
	}

	if intDigits == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'intDigits' is a nil pointer!\n",
			ePrefix.String())

		return numberStats, err
	}

	if fracDigits == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'fracDigits' is a nil pointer!\n",
			ePrefix.String())

		return numberStats, err
	}

	lenDecSepChars := len(decSeparatorChars.CharsArray)

	if lenDecSepChars == 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'decSeparatorSpec' is invalid!\n"+
			"The Rune Array for Decimal Separators is empty with a zero length!\n",
			ePrefix.String())

		return numberStats, err

	}

	new(runeArrayDtoAtom).empty(
		intDigits)

	intDigits.charSearchType =
		CharSearchType.LinearTargetStartingIndex()

	new(runeArrayDtoAtom).empty(
		fracDigits)

	fracDigits.charSearchType =
		CharSearchType.LinearTargetStartingIndex()

	if len(pureNumberStr) == 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'pureNumberStr' is invalid!\n"+
			"'pureNumberStr' is an empty string and has a\n"+
			"length of zero characters.\n",
			ePrefix.String())

		return numberStats, err

	}

	numberRunes := []rune(pureNumberStr)

	lenNumberRunes := len(numberRunes)

	foundMinusSign := false

	foundRadixPoint := false

	foundFirstNumericDigit := false

	isZero := true

	lastNumberRunesIdx := lenNumberRunes - 1

	var newLastIndex int

	for i := 0; i < lenNumberRunes; i++ {

		if numberRunes[i] == '-' {

			if leadingMinusSign == true &&
				foundFirstNumericDigit == false {

				foundMinusSign = true

				continue
			}

			if leadingMinusSign == false &&
				foundFirstNumericDigit == true {

				foundMinusSign = true

				break

			}

			continue
		}

		// Test For Decimal Separator
		if numberRunes[i] == decSeparatorChars.CharsArray[0] {

			newLastIndex = lenDecSepChars + i - 1

			if newLastIndex > lastNumberRunesIdx {
				continue
			}

			j := i

			for k := 0; k < lenDecSepChars; k++ {

				if numberRunes[j] != decSeparatorChars.CharsArray[k] {
					continue
				}

				j++
			}

			foundRadixPoint = true

			i = newLastIndex

			continue
		}

		// intDigits *RuneArrayDto,
		// fracDigits *RuneArrayDto,

		if numberRunes[i] >= '0' &&
			numberRunes[i] <= '9' {

			foundFirstNumericDigit = true

			if numberRunes[i] > '0' {
				isZero = false
			}

			if !foundRadixPoint {

				intDigits.CharsArray = append(
					intDigits.CharsArray, numberRunes[i])

			} else {
				// MUST BE -
				//	foundRadixPoint == true

				fracDigits.CharsArray = append(
					fracDigits.CharsArray, numberRunes[i])

			}
		}

	}

	numberStats.NumOfIntegerDigits =
		uint64(len(intDigits.CharsArray))

	numberStats.NumOfSignificantIntegerDigits =
		numberStats.NumOfIntegerDigits -
			intDigits.GetCountLeadingZeros()

	numberStats.NumOfFractionalDigits =
		uint64(len(fracDigits.CharsArray))

	numberStats.NumOfSignificantFractionalDigits =
		numberStats.NumOfFractionalDigits -
			fracDigits.GetCountTrailingZeros()

	if numberStats.NumOfFractionalDigits > 0 {

		numberStats.NumberValueType =
			NumValType.FloatingPoint()

	} else if numberStats.NumOfIntegerDigits > 0 {

		numberStats.NumberValueType =
			NumValType.Integer()

	} else {

		numberStats.NumberValueType =
			NumValType.None()
	}

	numberStats.IsZeroValue = isZero

	if !numberStats.IsZeroValue {

		if foundMinusSign {

			numberStats.NumberSign =
				NumSignVal.Negative()

		} else {

			numberStats.NumberSign =
				NumSignVal.Positive()

		}

	} else {
		// MUST BE -
		// numberStats.IsZeroValue == true

		numberStats.NumberSign =
			NumSignVal.Zero()

	}

	return numberStats, err
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
