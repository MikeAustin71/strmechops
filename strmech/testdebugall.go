package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/rand"
	"sync"
	"time"
)

type TestDebugAll struct {
	lock *sync.Mutex
}

func (testDebug *TestDebugAll) MathRounding(
	errorPrefix interface{}) {

	if testDebug.lock == nil {
		testDebug.lock = new(sync.Mutex)
	}

	testDebug.lock.Lock()

	defer testDebug.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TestDebugAll."+
			"MathRounding()",
		"")

	if err != nil {

		fmt.Printf("%v\n",
			err.Error())

		return
	}

	expectedIntDigits := "7"

	expectedFracDigits := "5"

	roundToFractionalDigits := 0

	expectedStr := "7"

	numberSign := NumSignVal.Negative()

	var intRunes RuneArrayDto

	intRunes,
		err = RuneArrayDto{}.NewRunes(
		[]rune(expectedIntDigits),
		CharSearchType.LinearTargetStartingIndex(),
		ePrefix.XCpy(
			"intRunes"))

	if err != nil {

		fmt.Printf("%v\n",
			err.Error())

		return
	}

	var fracRunes RuneArrayDto

	fracRunes,
		err = RuneArrayDto{}.NewRunes(
		[]rune(expectedFracDigits),
		CharSearchType.LinearTargetStartingIndex(),
		ePrefix.XCpy(
			"fracRunes"))

	if err != nil {

		fmt.Printf("%v\n",
			err.Error())

		return
	}

	nStrMathRoundAtom := numStrMathRoundingAtom{}

	err = nStrMathRoundAtom.roundHalfTowardsZero(
		&intRunes,
		&fracRunes,
		roundToFractionalDigits,
		numberSign,
		ePrefix.XCpy(
			fmt.Sprintf("Test #5 roundToFracDigits=%v",
				roundToFractionalDigits)))

	if err != nil {

		fmt.Printf("%v\n",
			err.Error())

		return
	}

	actualStr := intRunes.GetCharacterString()

	if fracRunes.GetRuneArrayLength() > 0 {

		actualStr += "." +
			fracRunes.GetCharacterString()

	}

	if expectedStr != actualStr {

		fmt.Printf("%v\n"+
			"Error: nStrMathRoundAtom.roundHalfAwayFromZero() #5\n"+
			"Expected String = '%v'\n"+
			"  Actual String = '%v'\n",
			ePrefix.String(),
			expectedStr,
			actualStr)

		return
	}

	fmt.Printf("%v\n"+
		"No Errors!\n"+
		"Expected String = '%v'\n"+
		"  Actual String = '%v'\n",
		ePrefix.String(),
		expectedStr,
		actualStr)

	return
}

func (testDebug *TestDebugAll) MathRounding2(
	errorPrefix interface{}) {

	if testDebug.lock == nil {
		testDebug.lock = new(sync.Mutex)
	}

	testDebug.lock.Lock()

	defer testDebug.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TestDebugAll."+
			"MathRounding2()",
		"")

	if err != nil {

		fmt.Printf("%v\n",
			err.Error())

		return
	}

	var actualStr string

	var intRunes, fracRunes RuneArrayDto

	intRunes,
		err = RuneArrayDto{}.NewRunes(
		[]rune("71"),
		CharSearchType.LinearTargetStartingIndex(),
		ePrefix)

	if err != nil {

		fmt.Printf("%v\n",
			err.Error())

		return
	}

	fracRunes,
		err = RuneArrayDto{}.NewRunes(
		[]rune("5"),
		CharSearchType.LinearTargetStartingIndex(),
		ePrefix)

	if err != nil {

		fmt.Printf("%v\n",
			err.Error())

		return
	}

	baseNum := "71.5"

	var numberSign NumericSignValueType
	rand.Seed(3200053)

	for i := 0; i < 30; i++ {

		if i%2 == 0 {
			numberSign = NumSignVal.Positive()
		} else {
			numberSign = NumSignVal.Negative()
		}

		waitVal := rand.Intn(400000000) + 100000000
		time.Sleep(time.Duration(waitVal))

		err = new(numStrMathRoundingAtom).roundRandomly(
			&intRunes,
			&fracRunes,
			0,
			numberSign,
			ePrefix.XCpy(
				fmt.Sprintf(
					"roundRandomly[%v]",
					i)))

		if err != nil {

			fmt.Printf("%v\n",
				err.Error())

			return
		}

		actualStr = intRunes.GetCharacterString()

		if fracRunes.GetRuneArrayLength() > 0 {

			actualStr += "." +
				fracRunes.GetCharacterString()

		}

		fmt.Printf("%v. %v\n"+
			"nStrMathRoundAtom.roundRandomly()\n"+
			"baseNum         = '%v'\n"+
			"Round To Digit  = '0'\n"+
			"Number Sign     = '%v'\n"+
			"  Actual String = '%v'\n\n",
			i,
			ePrefix.String(),
			baseNum,
			numberSign.String(),
			actualStr)

	}

	return
}

func (testDebug *TestDebugAll) RandomNum() {

}
