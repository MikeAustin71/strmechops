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

	baseNum := "71.5"

	var numberSign NumericSignValueType

	nStrMathRoundAtom := numStrMathRoundingAtom{}

	for i := 0; i < 30; i++ {

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

		if i%2 == 0 {
			numberSign = NumSignVal.Positive()
		} else {
			numberSign = NumSignVal.Negative()
		}

		err = nStrMathRoundAtom.roundRandomly(
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

		time.Sleep(time.Duration(500000000))

	}

	return
}

func (testDebug *TestDebugAll) RandomNum1(
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
			"RandomNum1()",
		"")

	if err != nil {

		fmt.Printf("%v\n",
			err.Error())

		return
	}

	var intNums []int
	var selection int

	newRandGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 10; i++ {

		// Old Code Pre Go 1.20
		// rand.Seed(time.Now().UnixNano())
		//intNums = rand.Perm(10)

		intNums = newRandGenerator.Perm(10)

		for j := 0; j < 3; j++ {
			selection = rand.Intn(9)
		}

		fmt.Printf("%v. "+
			"%v\n"+
			"Selection: %v\n",
			i+1,
			intNums,
			intNums[selection])

		time.Sleep(500000000)
	}

	fmt.Printf("\n\nSuccessful Completion\n"+
		"%v\n",
		ePrefix.String())

	return
}

func (testDebug *TestDebugAll) RandomNum2(
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
			"RandomNum2()",
		"")

	if err != nil {

		fmt.Printf("%v\n",
			err.Error())

		return
	}

	for i := 0; i < 30; i++ {

		x := time.Now().UnixNano()
		y := x
		z := y % 2

		fmt.Printf("%v. "+
			"%v - Remainder\n"+
			"%v - Nanoseconds\n",
			i+1,
			z,
			y)

		time.Sleep(500000000)
	}

	fmt.Printf("\n\nSuccessful Completion\n"+
		"%v\n",
		ePrefix.String())

	return
}

func (testDebug *TestDebugAll) RandomNum3Crypto(
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
			"RandomNum3Crypto()",
		"")

	if err != nil {

		fmt.Printf("\n%v\n"+
			"Error return from ePref.ErrPrefixDto{}.NewIEmpty()\n"+
			"Error:\n%v\n\n",
			"RandomNum3Crypto()",
			err.Error())

		return
	}

	fmt.Printf("\n%v\n"+
		"Crypto Random Number Test\n\n",
		ePrefix.String())

	nStrMathQuark := numStrMathQuark{}

	var intervalUpperBound, randomNum int64

	intervalUpperBound = 30

	for i := 0; i < 30; i++ {

		randomNum,
			err = nStrMathQuark.randomInt64(
			intervalUpperBound,
			ePrefix.XCpy(
				fmt.Sprintf("Cycle No: %v",
					i)))

		if err != nil {
			fmt.Printf("%v\n"+
				"Error return from nStrMathQuark.randomInt64()\n"+
				"Cycle = %v\n"+
				"intervalUpperBound = %v\n",
				ePrefix.String(),
				i,
				intervalUpperBound)

			return
		}

		fmt.Printf("%v. Remainder: %v Random Num: %v\n",
			i+1,
			randomNum%2,
			randomNum)

		time.Sleep(500000000)
	}

	fmt.Printf("\n\nSuccessful Completion\n"+
		"%v\n",
		ePrefix.String())

	return
}
