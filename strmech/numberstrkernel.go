package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// NumberStrKernel - Contains all the base specifications necessary
// to construct an integer or floating pointer number string.
//
type NumberStrKernel struct {
	integerDigits RuneArrayDto // An array of numeric digits
	//                             constituting the integer portion
	//                             of the numeric value represented
	//                             by this instance of NumberStrKernel

	hasIntegerDigits bool

	fractionalDigits RuneArrayDto // An array of numeric digits
	//                             constituting the integer portion
	//                             of the numeric value represented
	//                             by this instance of NumberStrKernel

	hasFractionalDigits bool // If set to 'true' this signals that
	//                           the numeric value represented by
	//                           this instance of NumberStrKernel
	//                           contains fractional numeric digits.

	numberSign int // Set to one of three values:
	//                           0  = Number has Zero Numeric Value
	//                           -1 = Number Sign is Negative
	//                           1  = (+1) Number Sign is Positive

	isNonZeroValue bool // If 'true' the Numeric Value is greater than
	//                     or less than zero ('0').
	//                     If 'false' the Numeric Value is equal to zero.

	lock *sync.Mutex
}

func (numStrKernel *NumberStrKernel) AddIntegerDigit(
	integerDigit rune,
	errorPrefix interface{}) (
	err error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"AddIntegerDigit()",
		"")

	if err != nil {
		return err
	}

	if integerDigit < '0' ||
		integerDigit > '9' {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'integerDigit' is invalid!\n"+
			"Integer Runes must represent a numberic character between\n"+
			"'0' and '9', inclusive. 'integerDigit' fails to meet this criterion.\n"+
			"The rune value of 'integerDigit' is %v\n"+
			"The string value of 'integerDigit' is %v\n",
			ePrefix.XCpy("integerDigit Error"),
			integerDigit,
			string(integerDigit))

		return err
	}

	numStrKernel.integerDigits.CharsArray =
		append(numStrKernel.integerDigits.CharsArray,
			integerDigit)

	numStrKernel.hasIntegerDigits = true

	if integerDigit != '0' {
		numStrKernel.isNonZeroValue = true
	}

	return err
}

func (numStrKernel *NumberStrKernel) AddFractionalDigit(
	fractionalDigit rune,
	errorPrefix interface{}) (
	err error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"AddFractionalDigit()",
		"")

	if err != nil {
		return err
	}

	if fractionalDigit < '0' ||
		fractionalDigit > '9' {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fractionalDigit' is invalid!\n"+
			"Fractional Runes must represent a numberic character between\n"+
			"'0' and '9', inclusive. 'fractionalDigit' fails to meet this criterion.\n"+
			"The rune value of 'fractionalDigit' is %v\n"+
			"The string value of 'fractionalDigit' is %v\n",
			ePrefix.XCpy("fractionalDigit Error"),
			fractionalDigit,
			string(fractionalDigit))

		return err
	}

	numStrKernel.fractionalDigits.CharsArray =
		append(numStrKernel.fractionalDigits.CharsArray,
			fractionalDigit)

	numStrKernel.hasFractionalDigits = true

	if fractionalDigit != '0' {
		numStrKernel.isNonZeroValue = true
	}

	return err
}

func (numStrKernel *NumberStrKernel) GetIntegerDigits(
	errorPrefix interface{}) (
	copyOfIntegerDigits RuneArrayDto,
	err error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"GetIntegerDigits()",
		"")

	if len(numStrKernel.integerDigits.CharsArray) == 0 {
		return copyOfIntegerDigits, err
	}

	copyOfIntegerDigits,
		err = numStrKernel.integerDigits.CopyOut(
		ePrefix.XCpy("copyOfIntegerDigits<-numStrKernel.integerDigits"))

	return copyOfIntegerDigits, err
}

func (numStrKernel *NumberStrKernel) GetIntegerRuneArray() []rune {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	return numStrKernel.integerDigits.GetRuneArray()
}

func (numStrKernel *NumberStrKernel) GetFractionalDigits(
	errorPrefix interface{}) (
	copyOfFractionalDigits RuneArrayDto,
	err error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"GetFractionalDigits()",
		"")

	if len(numStrKernel.fractionalDigits.CharsArray) == 0 {
		return copyOfFractionalDigits, err
	}

	copyOfFractionalDigits,
		err = numStrKernel.fractionalDigits.CopyOut(
		ePrefix.XCpy("copyOfFractionalDigits<-" +
			"numStrKernel.fractionalDigits"))

	return copyOfFractionalDigits, err
}

func (numStrKernel *NumberStrKernel) GetFractionalRuneArray() []rune {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	return numStrKernel.fractionalDigits.GetRuneArray()
}

func (numStrKernel *NumberStrKernel) SetNumberSign(
	numberSign int,
	errorPrefix interface{}) (
	err error) {

	if numStrKernel.lock == nil {
		numStrKernel.lock = new(sync.Mutex)
	}

	numStrKernel.lock.Lock()

	defer numStrKernel.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumberStrKernel."+
			"SetNumberSign()",
		"")

	if err != nil {
		return err
	}

	if numberSign > 1 ||
		numberSign < -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numberSign' is invalid!\n"+
			"'numberSign' must be set to one of three values:"+
			"     0 = Numeric Value of Zero ('0').\n"+
			"     1 = + or positive number sign\n"+
			"    -1 = - or negative number sign\n"+
			"Input parameter 'numberSign' = '%v'\n",
			ePrefix.String(),
			numberSign)

		return err
	}

	if numberSign == 0 {
		numStrKernel.isNonZeroValue = false
	} else {
		numStrKernel.isNonZeroValue = true
	}

	numStrKernel.numberSign = numberSign

	return err
}
