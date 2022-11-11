package examples

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"github.com/MikeAustin71/strmechops/strmech"
	"io"
	"math/big"
	"strings"
	"sync"
)

type MainTest02 struct {
	input string
}

func (MainTest02) NumberStrKernelCompare01() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"NumberStrKernelCompare01()",
		"")

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Starting Run!\n"+
		"Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

	origIntStr := "1234"
	origFracStr := "5678"

	var err error
	var baseValue strmech.NumberStrKernel
	var intDigitsDto, fracDigitsDto strmech.RuneArrayDto

	intDigitsDto,
		err = new(strmech.RuneArrayDto).NewString(
		origIntStr,
		strmech.CharSearchType.LinearTargetStartingIndex(),
		ePrefix.XCpy(
			"origIntStr"))

	if err != nil {
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	fracDigitsDto,
		err = new(strmech.RuneArrayDto).NewString(
		origFracStr,
		strmech.CharSearchType.LinearTargetStartingIndex(),
		ePrefix.XCpy(
			"origIntStr"))

	if err != nil {
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	baseValue,
		err = new(strmech.NumberStrKernel).NewFromRuneDto(
		&intDigitsDto,
		&fracDigitsDto,
		strmech.NumSignVal.Positive(),
		ePrefix.XCpy(
			"baseValue<-"))

	if err != nil {
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	var testNStrValue01 strmech.NumberStrKernel

	testValueIntDigits := "5234"
	testValueFracDigits := "5678"

	testValue := testValueIntDigits +
		"." +
		testValueFracDigits

	testBigFloat := big.NewFloat(0.0)

	_,
		_,
		err = testBigFloat.Parse(testValue, 10)

	if err != nil {
		fmt.Printf("\n%v\n"+
			"Error return from testBigFloat.Parse(testValue,10)\n"+
			"testValue = '%v'\n",
			ePrefix.String(),
			testValue)

		return
	}

	fmt.Printf("Verifying 'testBigFloat': %v\n",
		testBigFloat.Text('f', -1))

	minPrecision := testBigFloat.MinPrec()

	testBigFloat.SetPrec(minPrecision)

	testNStrValue01,
		err = new(strmech.NumberStrKernel).NewFromFloatValue(
		testBigFloat,
		ePrefix.XCpy(
			"testNStrValue01<-testBigFloat"))

	if err != nil {
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	var actualNumSign strmech.NumericSignValueType

	actualNumSign,
		err = testNStrValue01.GetNumberSign(
		ePrefix)

	if err != nil {
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	if actualNumSign != strmech.NumSignVal.Positive() {

		fmt.Printf("\n%v\n"+
			"Test # 0.01 testNStrValue01 Floating Point Value\n"+
			"Error: Expected Number Sign Equal To 'Positive'.\n"+
			"Instead, Number Sign Equals Integer Value ='%v'"+
			"Actual Number Sign string value = '%v'\n",
			ePrefix.String(),
			actualNumSign.XValueInt(),
			actualNumSign.String())

		return
	}

	actualIntStr := testNStrValue01.GetIntegerString()

	if actualIntStr != testValueIntDigits {
		fmt.Printf("\n%v\n"+
			"Test#5\n"+
			"Error: actualIntStr != testValueIntDigits\n"+
			"actualIntStr         = '%v'\n"+
			"testValueIntDigits   = '%v'\n",
			ePrefix.String(),
			actualIntStr,
			testValueIntDigits)

		return
	}

	actualFracStr := testNStrValue01.GetFractionalString()

	if actualFracStr != testValueFracDigits {

		fmt.Printf("\n%v\n"+
			"Test#6\n"+
			"Error: actualFracStr != testValueFracDigits\n"+
			"actualFracStr         = '%v'\n"+
			"testValueFracDigits   = '%v'\n",
			ePrefix.String(),
			actualFracStr,
			testValueFracDigits)

		return
	}

	var comparisonResult int

	comparisonResult,
		err = baseValue.Compare(
		&testNStrValue01,
		ePrefix.XCpy(
			"<-testNStrValue01"))

	if err != nil {
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	if comparisonResult != -1 {
		fmt.Printf("\n%v\n"+
			"Test#7\n"+
			"Error: Expected a comparisonResult of -1.\n"+
			"Instead, comparisonResult = '%v'\n",
			ePrefix.String(),
			comparisonResult)

		return
	}

	var testNStrValue02 strmech.NumberStrKernel

	testValueIntDigits = "234"
	testValueFracDigits = "5678"

	testValue = "-" + testValueIntDigits +
		"." +
		testValueFracDigits

	testNStrValue02,
		err = new(strmech.NumberStrKernel).NewParsePureNumberStr(
		testValue,
		".",
		true,
		ePrefix.XCpy(
			"testNStrValue02<-testValue"))

	if err != nil {
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	actualIntStr = testNStrValue02.GetIntegerString()

	if actualIntStr != testValueIntDigits {
		fmt.Printf("\n%v\n"+
			"Test#8\n"+
			"Error: actualIntStr != testValueIntDigits\n"+
			"actualIntStr         = '%v'\n"+
			"testValueIntDigits   = '%v'\n",
			ePrefix.String(),
			actualIntStr,
			testValueIntDigits)

		return
	}

	actualFracStr = testNStrValue02.GetFractionalString()

	if actualFracStr != testValueFracDigits {

		fmt.Printf("\n%v\n"+
			"Test#9\n"+
			"Error: actualFracStr != testValueFracDigits\n"+
			"actualFracStr         = '%v'\n"+
			"testValueFracDigits   = '%v'\n",
			ePrefix.String(),
			actualFracStr,
			testValueFracDigits)

		return
	}

	actualNumSign,
		err = testNStrValue02.GetNumberSign(
		ePrefix)

	if err != nil {
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	if actualNumSign != strmech.NumSignVal.Negative() {

		fmt.Printf("\n%v\n"+
			"Test # 9.5"+
			"Error: Expected Number Sign Equal To 'Negative'.\n"+
			"Instead, Number Sign Equals Integer Value ='%v'"+
			"Actual Number Sign string value = '%v'\n",
			ePrefix.String(),
			actualNumSign.XValueInt(),
			actualNumSign.String())

		return
	}

	comparisonResult,
		err = baseValue.Compare(
		&testNStrValue02,
		ePrefix.XCpy(
			"<-testNStrValue01"))

	if err != nil {
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	if comparisonResult != 1 {
		fmt.Printf("\n%v\n"+
			"Test#10\n"+
			"Error: Expected a comparisonResult of +1.\n"+
			"Instead, comparisonResult = '%v'\n",
			ePrefix.String(),
			comparisonResult)
	}

	var testNStrValue03 strmech.NumberStrKernel

	testValueIntDigits = "1234"
	testValueFracDigits = "5679"

	testValue = testValueIntDigits +
		"." +
		testValueFracDigits

	testNStrValue03,
		err = new(strmech.NumberStrKernel).NewParsePureNumberStr(
		testValue,
		".",
		true,
		ePrefix.XCpy(
			"testNStrValue03<-testValue"))

	if err != nil {
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	actualIntStr = testNStrValue03.GetIntegerString()

	if actualIntStr != testValueIntDigits {
		fmt.Printf("\n%v\n"+
			"Test#11\n"+
			"Error: actualIntStr != testValueIntDigits\n"+
			"actualIntStr         = '%v'\n"+
			"testValueIntDigits   = '%v'\n",
			ePrefix.String(),
			actualIntStr,
			testValueIntDigits)

		return
	}

	actualFracStr = testNStrValue03.GetFractionalString()

	if actualFracStr != testValueFracDigits {

		fmt.Printf("\n%v\n"+
			"Test#12\n"+
			"Error: actualFracStr != testValueFracDigits\n"+
			"actualFracStr         = '%v'\n"+
			"testValueFracDigits   = '%v'\n",
			ePrefix.String(),
			actualFracStr,
			testValueFracDigits)

		return
	}

	actualNumSign,
		err = testNStrValue03.GetNumberSign(
		ePrefix)

	if err != nil {
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	if actualNumSign != strmech.NumSignVal.Positive() {

		fmt.Printf("\n%v\n"+
			"Test #13\n"+
			"Error: Expected Number Sign Equal To 'Positive'.\n"+
			"Instead, Number Sign Equals Integer Value ='%v'"+
			"Actual Number Sign string value = '%v'\n",
			ePrefix.String(),
			actualNumSign.XValueInt(),
			actualNumSign.String())

		return
	}

	comparisonResult,
		err = baseValue.Compare(
		&testNStrValue03,
		ePrefix.XCpy(
			"<-testNStrValue03"))

	if err != nil {
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	if comparisonResult != -1 {
		fmt.Printf("\n%v\n"+
			"Test#14\n"+
			"Error: Expected a comparisonResult of -1.\n"+
			"Instead, comparisonResult = '%v'\n",
			ePrefix.String(),
			comparisonResult)
	}

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		"Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

	return
}

func (MainTest02) NumberStrKernelPureNumStr01() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"NumberStrKernelPureNumStr01()",
		"")

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Starting Run!\n"+
		"Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

	origIntStr := "1234567"

	origFracStr := "890"

	expectedSciNotStr := "1.23456789 x 10^6"

	compositeNumStr := origIntStr + "." + origFracStr

	var err error
	var numStrKernel01 strmech.NumberStrKernel

	numStrKernel01,
		err = new(strmech.NumberStrKernel).NewParsePureNumberStr(
		compositeNumStr,
		".",
		true,
		ePrefix)

	if err != nil {
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	actualIntStr := numStrKernel01.GetIntegerString()

	if actualIntStr != origIntStr {
		fmt.Printf("%v\n"+
			"Test#1\n"+
			"Error: actualIntStr != origIntStr\n"+
			"actualIntStr = '%v'\n"+
			"origIntStr   = '%v'\n",
			ePrefix.String(),
			actualIntStr,
			origIntStr)

		return
	}

	actualFracStr := numStrKernel01.GetFractionalString()

	if actualFracStr != origFracStr {

		fmt.Printf("%v\n"+
			"Test#2\n"+
			"Error: actualFracStr != origFracStr\n"+
			"actualFracStr = '%v'\n"+
			"origFracStr   = '%v'\n",
			ePrefix.String(),
			actualFracStr,
			origFracStr)

		return
	}

	var sciNot01 strmech.SciNotationKernel

	sciNot01,
		err = numStrKernel01.GetScientificNotation(
		strmech.NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"sciNot01<-"))

	if err != nil {
		fmt.Printf("%v\n",
			err.Error())
		return
	}

	var actualSciNotStr string

	actualSciNotStr = sciNot01.GetNumStrExponentFmt()

	if actualSciNotStr != expectedSciNotStr {

		fmt.Printf("%v\n"+
			"Test#3\n"+
			"Scientific Notaion String Error\n"+
			"actualSciNotStr  !=  expectedSciNotStr\n"+
			"actualSciNotStr   = '%v'\n"+
			"expectedSciNotStr = '%v'\n",
			ePrefix.String(),
			actualSciNotStr,
			expectedSciNotStr)

		return

	}

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		"Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

	return

}

func (MainTest02) NumberStrKernelExtendArrays01() {

	funcName := "\nMainTest02.NumberStrKernelExtendArrays01()"

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		funcName,
		"")

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Starting Run!\n"+
		"Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

	origIntStr := "1234"
	origFracStr := "5678"
	origNumStr := origIntStr + "." + origFracStr

	var err error
	var numStrKernel01 strmech.NumberStrKernel
	var numberStrSearchResults strmech.CharSearchNumStrParseResultsDto

	numberStrSearchResults,
		numStrKernel01,
		err = new(strmech.NumberStrKernel).NewParseUSNumberStr(
		origNumStr,
		0,
		-1,
		nil,
		false,
		&ePrefix)

	if err != nil {
		fmt.Printf("%v\n",
			err.Error())
		return
	}

	if !numberStrSearchResults.FoundIntegerDigits {

		fmt.Printf("%v\n"+
			"Test#1\n"+
			"Error: No Integer Digits Found!\n"+
			"The number string search failed to find integer digits.\n"+
			"origNumStr   = '%v'\n"+
			"origIntStr   = '%v'\n",
			ePrefix.String(),
			origNumStr,
			origIntStr)

		return

	}

	if !numberStrSearchResults.FoundDecimalDigits {

		fmt.Printf("%v\n"+
			"Test#1\n"+
			"Error: No Fractional Digits Found!\n"+
			"The number string search failed to find fractional digits.\n"+
			"origNumStr   = '%v'\n"+
			"origFracStr   = '%v'\n",
			ePrefix.String(),
			origNumStr,
			origFracStr)

		return

	}

	actualIntStr := numStrKernel01.GetIntegerString()

	if actualIntStr != origIntStr {
		fmt.Printf("%v\n"+
			"Test#2\n"+
			"Error: actualIntStr != origIntStr\n"+
			"actualIntStr = '%v'\n"+
			"origIntStr   = '%v'\n",
			ePrefix.String(),
			actualIntStr,
			origIntStr)

		return
	}

	actualFracStr := numStrKernel01.GetFractionalString()

	if actualFracStr != origFracStr {

		fmt.Printf("%v\n"+
			"Test#3\n"+
			"Error: actualFracStr != origFracStr\n"+
			"actualFracStr = '%v'\n"+
			"origFracStr   = '%v'\n",
			ePrefix.String(),
			actualFracStr,
			origFracStr)

		return
	}

	err = numStrKernel01.ExtendIntegerDigitsArray(
		'0',
		3,
		true,
		ePrefix)

	if err != nil {
		fmt.Printf("%v\n",
			err.Error())
		return
	}

	//origIntStr := "1234"

	expectedIntStr := "1234000"

	actualIntStr = numStrKernel01.GetIntegerString()

	if actualIntStr != expectedIntStr {

		fmt.Printf("%v\n"+
			"Test#4\n"+
			"Error: actualIntStr != expectedIntStr\n"+
			"origIntStr       = '%v'\n"+
			"actualIntStr     = '%v'\n"+
			"expectedIntStr   = '%v'\n",
			ePrefix.String(),
			origIntStr,
			actualIntStr,
			expectedIntStr)

		return
	}

	err = numStrKernel01.ExtendFractionalDigitsArray(
		'0',
		3,
		false,
		ePrefix)

	if err != nil {
		fmt.Printf("%v\n",
			err.Error())
		return
	}

	//origFracStr := "5678"
	expectedFracStr := "0005678"

	actualFracStr = numStrKernel01.GetFractionalString()

	if actualFracStr != expectedFracStr {

		fmt.Printf("%v\n"+
			"Test#4\n"+
			"Error: actualFracStr != expectedFracStr\n"+
			"origFracStr       = '%v'\n"+
			"actualFracStr     = '%v'\n"+
			"expectedFracStr   = '%v'\n",
			ePrefix.String(),
			origFracStr,
			actualFracStr,
			expectedFracStr)

		return
	}

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		"Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

	return
}

func (MainTest02) RuneArrayDtoDeleteChars01() {

	funcName := "MainTest02.RuneArrayDtoDeleteChars01()"

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		funcName,
		"")

	breakStr := strings.Repeat("=", 50)

	var runeArrayDto strmech.RuneArrayDto
	var err error

	originalStr := "How now brown cow."
	expectedStr := "How now brown"
	numOfCharsToDelete := uint64(5)

	runeArrayDto = new(strmech.RuneArrayDto).NewStringDefault(
		originalStr)

	preDeleteStr := runeArrayDto.GetCharacterString()

	fmt.Printf("\nPre Delete Str = '%v'\n",
		preDeleteStr)

	err = runeArrayDto.DeleteLeadingTrailingChars(
		numOfCharsToDelete,
		true,
		ePrefix.XCpy(
			"runeArrayDto"))

	if err != nil {
		fmt.Printf("%v\n",
			err.Error())
		return
	}

	actualStr := runeArrayDto.GetCharacterString()

	if actualStr != expectedStr {

		fmt.Printf("\n%v\n"+
			"Test # 2\n"+
			"Original String = '%v'\n"+
			"Error: Returned 'actualStr' is NOT EQUAL\n"+
			"to 'expectedStr'!\n"+
			"expectedStr = '%v'\n"+
			"actualStr   = '%v'\n",
			ePrefix.String(),
			originalStr,
			expectedStr,
			actualStr)

		return
	}

	fmt.Printf("\nPost Delet Str = '%v'\n",
		actualStr)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		"Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

	return
}

func (MainTest02) NumStrKernelParseUSNumStr() {

	funcName := "MainTest02.NumStrKernelParseUSNumStr()"

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		funcName,
		"")

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\nFunction: %v\n",
		ePrefix.String())

	fmt.Printf(breakStr + "\n\n")

	expectedNumStr := "-1234567.1234567"

	sMechNStrKernel := strmech.NumberStrKernel{}

	searchResults,
		nStrKernel,
		err := sMechNStrKernel.NewParseUSNumberStr(
		expectedNumStr,
		0,
		-1,
		[]string{},
		false,
		ePrefix)

	if err != nil {
		fmt.Println(
			fmt.Sprintf("%v",
				err.Error()))

		return
	}

	sBuilder := strings.Builder{}

	err = searchResults.GetParameterTextListing(
		&sBuilder,
		false,
		false,
		ePrefix)

	if err != nil {
		fmt.Println(
			fmt.Sprintf("%v",
				err.Error()))

		return
	}

	fmt.Printf("Search Results:\n"+
		"%v\n",
		sBuilder.String())

	actualIntegerDigits := nStrKernel.GetIntegerString()

	actualFractionalDigits := nStrKernel.GetFractionalString()

	var actualNumStr string

	var numSign strmech.NumericSignValueType

	numSign,
		err = nStrKernel.GetNumberSign(
		ePrefix)

	if err != nil {
		fmt.Println(
			fmt.Sprintf("%v",
				err.Error()))

		return
	}

	if numSign == strmech.NumSignVal.Negative() {
		actualNumStr += "-"
	}

	actualNumStr += actualIntegerDigits

	if len(actualFractionalDigits) > 0 {
		actualNumStr += "."
		actualNumStr += actualFractionalDigits
	}

	fmt.Printf("Expected Number:       %v\n",
		expectedNumStr)

	fmt.Printf("Actual Number:         %v\n",
		actualNumStr)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		"Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")
}

func (MainTest02) NumStrKernelSetSignedIntValue() {

	funcName := "MainTest02.NumStrKernelSetSignedIntValue()"

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		funcName,
		"")

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\nFunction: %v\n",
		ePrefix.String())

	fmt.Printf(breakStr + "\n\n")

	expectedNumStr := "1234567"

	sMechNStrKernel := strmech.NumberStrKernel{}

	bigIntNum := big.NewInt(1234567)

	nStrKernel,
		err := sMechNStrKernel.NewFromSignedIntValue(
		bigIntNum,
		ePrefix)

	if err != nil {
		fmt.Println(
			fmt.Sprintf("%v",
				err.Error()))

		return
	}

	actualIntegerDigits := nStrKernel.GetIntegerString()

	actualFractionalDigits := nStrKernel.GetFractionalString()

	actualNumStr := actualIntegerDigits

	if len(actualFractionalDigits) > 0 {
		actualNumStr += "."
		actualNumStr += actualFractionalDigits
	}

	fmt.Printf("Big Int Number:        %v\n",
		bigIntNum.Text(10))

	fmt.Printf("Expected Number:       %v\n",
		expectedNumStr)

	fmt.Printf("Actual Number:         %v\n",
		actualNumStr)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		"Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

}

func (MainTest02) NumStrKernelToFloatConversion() {

	funcName := "MainTest02.NumStrKernelToFloatConversion()"

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		funcName,
		"")

	intDigits := "1234567"
	fracDigits := "1234567"
	expectedNumStr :=
		intDigits +
			"." +
			fracDigits

	sMechNStrKernel := strmech.NumberStrKernel{}
	numberSign := strmech.NumSignVal.None()

	numStrKernel,
		err := sMechNStrKernel.NewFromStringDigits(
		intDigits,
		fracDigits,
		numberSign,
		ePrefix.XCpy(
			fmt.Sprintf(
				"intDigits=%v",
				intDigits)))

	if err != nil {
		fmt.Println(
			fmt.Sprintf("%v",
				err.Error()))

		return
	}

	var bigFloatNumber *big.Float

	var roundType = strmech.NumRoundType.NoRounding()
	var numOfFractionalDigits int
	var roundToFractionalDigits = 0

	bigFloatNumber,
		numOfFractionalDigits,
		err = numStrKernel.GetBigFloatNum(
		roundType,
		roundToFractionalDigits,
		ePrefix)

	if err != nil {
		fmt.Println(
			fmt.Sprintf("%v",
				err.Error()))

		return
	}

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\nFunction: %v\n",
		ePrefix.String())

	fmt.Printf(breakStr + "\n\n")

	fmt.Printf("Big Float String Value:         '%v'\n",
		expectedNumStr)

	fmt.Printf(breakStr + "\n\n")

	fmt.Printf("Big Float Intial Precision:     '%v'\n",
		bigFloatNumber.Prec())

	fmt.Printf("Big Float Intial Min Precision: '%v'\n",
		bigFloatNumber.MinPrec())

	fmt.Printf("Big Float Intial Mode: '%v'\n",
		bigFloatNumber.Mode().String())

	fmt.Printf("Big Float Intial Accuracy:      '%v'\n",
		bigFloatNumber.Acc().String())

	fmt.Printf("Big Float Numeric Value g -1:   '%v'\n",
		bigFloatNumber.Text('g', -1))

	fmt.Printf("Big Float Numeric Value f -1:   '%v'\n",
		bigFloatNumber.Text('f', -1))

	fmt.Printf("Big Float Val f numFracDigits:  '%v'\n",
		bigFloatNumber.Text('f', numOfFractionalDigits))

	fmt.Printf("Big Float Val f 6:              '%v'\n",
		bigFloatNumber.Text('f', 6))

	fmt.Printf("Big Float Val f 7:              '%v'\n",
		bigFloatNumber.Text('f', 7))

	fmt.Printf("Big Float Val .11g:             '%.11g'\n",
		bigFloatNumber)

	fmt.Printf("Big Float Val .11f:             '%.11f'\n",
		bigFloatNumber)

	fmt.Printf("Big Float Val .20f:             '%.20f'\n",
		bigFloatNumber)

	fmt.Printf("Big Float Val f  20 Prec :      '%v'\n",
		bigFloatNumber.Text('f', 20))

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		"Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

}

func (MainTest02) NumStrKernelToIntConversion() {

	funcName := "MainTest02.NumStrKernelToIntConversion()"

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		funcName,
		"")

	intDigits := "12345"
	fracDigits := "6"

	sMechNStrKernel := strmech.NumberStrKernel{}
	numberSign := strmech.NumSignVal.None()

	numStrKernel,
		err := sMechNStrKernel.NewFromStringDigits(
		intDigits,
		fracDigits,
		numberSign,
		ePrefix.XCpy(
			fmt.Sprintf(
				"intDigits=%v",
				intDigits)))

	if err != nil {
		fmt.Println(
			fmt.Sprintf("%v",
				err.Error()))

		return
	}

	var bigIntNumber *big.Int

	var roundType = strmech.NumRoundType.HalfAwayFromZero()

	bigIntNumber,
		err = numStrKernel.GetBigIntNum(
		roundType,
		ePrefix)

	if err != nil {
		fmt.Println(
			fmt.Sprintf("%v",
				err.Error()))

		return
	}

	fmt.Printf("Original Integer Digits: %v\n",
		intDigits)

	fmt.Printf("Original Fractional Digits: %v\n",
		fracDigits)

	fmt.Printf("NumStrKernel Integer Number: %v\n",
		bigIntNumber.Text(10))

}

func (mTest02 MainTest02) CharSearchTermination01() {

	funcName := "MainTest02.CharSearchTermination01()"

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		funcName,
		"")

	var outputStr string
	const maxLineLen = 78

	strBuilder := strings.Builder{}

	mt02Nanobot := mainTest02Nanobot{}

	err := mt02Nanobot.SetLeadingMarquee(
		&strBuilder,
		funcName,
		"Testing StrMech.CharSearchTerminationType()",
		maxLineLen,
		false,
		&ePrefix)

	if err != nil {
		fmt.Println(
			fmt.Sprintf("%v",
				err.Error()))

		return
	}

	charSearchTermType := strmech.CharSearchTerminationType(-972)

	strName := charSearchTermType.XReturnNoneIfInvalid()

	if strName.String() != "None" {
		outputStr = fmt.Sprintf(
			"%v\n"+
				"Error: Expected strName == 'None'\n"+
				"Actual strName integer value = '%v'\n",
			ePrefix.String(),
			strName.XValueInt())

		fmt.Println(outputStr)

		return

	}

	// Trailing Title Marquee
	err = mt02Nanobot.SetTrailingMarquee(
		&strBuilder,
		funcName,
		"SUCCESSFUL COMPLETION!",
		maxLineLen,
		false,
		ePrefix.XCpy(
			"Trailing Title Marquee"))

	if err != nil {
		fmt.Println(
			fmt.Sprintf("%v",
				err.Error()))

		return
	}

	fmt.Println(strBuilder.String())
}

func (mTest02 MainTest02) StrMechRead01() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainTest02.StrMechRead01()",
		"")

	var outputStr string

	expected := "Original sops1 base string"
	lenExpected := len(expected)

	p := make([]byte, 5, 15)

	s1 := new(strmech.StrMech)

	s1.SetStringData(expected)

	n := 0

	var err error

	err = nil

	b := strings.Builder{}

	b.Grow(len(expected) + 150)

	for err != io.EOF {

		n, err = s1.Read(p)

		if err != nil && err != io.EOF {
			outputStr = fmt.Sprintf(
				"%v\n"+
					"Error returned by s1.Read(p). "+
					"Error='%v' \n",
				ePrefix.XCpy(
					"s1.Read(p)"),
				err.Error())

			fmt.Println(outputStr)

			return
		}

		b.Write(p[:n])

		for i := 0; i < len(p); i++ {
			p[i] = byte(0)
		}

	}

	actualStr := b.String()

	if expected != actualStr {

		outputStr = fmt.Sprintf("%v\n"+
			"Error: Expected StrOut='%v'.\n"+
			"Instead, StrOut='%v'\n",
			ePrefix.String(),
			expected,
			actualStr)

		fmt.Println(outputStr)

		return
	}

	lenActual := len(actualStr)

	if lenExpected != lenActual {
		outputStr =
			fmt.Sprintf("%v\n"+
				"Error: Expected bytes read ='%v'.\n"+
				"Instead, bytes read='%v'\n",
				ePrefix.String(),
				lenExpected,
				lenActual)

		fmt.Println(outputStr)

		return
	}

	// Total available Length of Output Line
	const maxLineLen = 78

	txtFormatCol := strmech.TextFormatterCollection{}

	strBuilder := strings.Builder{}

	err = txtFormatCol.SetStdFormatParamsLine1Col(
		" ",
		maxLineLen,
		strmech.TxtJustify.Center(),
		"",
		false,
		"",
		-1,
		false,
		ePrefix.XCpy(
			"1-Column Setup"))

	if err != nil {

		outputStr =
			fmt.Sprintf("%v\n"+
				"Error: txtFormatCol.SetStdFormatParamsLine1Col()\n"+
				"Error=\n%v\n",
				ePrefix.String(),
				err.Error())

		fmt.Printf(outputStr + "\n")

		return
	}

	// Trailing Title Marquee
	// Top Blank Line

	txtFormatCol.AddLineBlank(
		1,
		"")

	// Filler =======
	// Marquee Top
	txtFormatCol.AddLineSolid(
		" ",
		"=",
		maxLineLen-2,
		" ",
		false,
		"",
		-1,
		false)

	// Title # 1
	err = txtFormatCol.AddLine1Col(
		ePrefix.String(),
		ePrefix.XCpy(
			"Title Line 1"))

	if err != nil {
		outputStr =
			fmt.Sprintf("%v\n"+
				"Error: txtFormatCol.AddLine1Col()\n"+
				"Error=\n%v\n",
				ePrefix.String(),
				err.Error())

		fmt.Printf(outputStr + "\n")

		return
	}

	// Title # 2
	err = txtFormatCol.AddLine1Col(
		fmt.Sprintf("Bytes Read: %v",
			lenActual),
		ePrefix.XCpy(
			"Title Line 1"))

	if err != nil {
		outputStr =
			fmt.Sprintf("%v\n"+
				"Error: txtFormatCol.AddLine1Col()\n"+
				"Error=\n%v\n",
				ePrefix.String(),
				err.Error())

		fmt.Printf(outputStr + "\n")

		return
	}

	// Title # 3
	err = txtFormatCol.AddLine1Col(
		"SUCCESSFUL COMPLETION",
		ePrefix.XCpy(
			"Title Line 1"))

	if err != nil {
		outputStr =
			fmt.Sprintf("%v\n"+
				"Error: txtFormatCol.AddLine1Col()\n"+
				"Error=\n%v\n",
				ePrefix.String(),
				err.Error())

		fmt.Printf(outputStr + "\n")

		return
	}

	// Filler =======
	// Marquee Bottom
	txtFormatCol.AddLineSolid(
		" ",
		"=",
		maxLineLen-2,
		" ",
		false,
		"",
		-1,
		false)

	// Blank Line
	txtStrBuilder := strmech.TextStrBuilder{}

	err =
		txtStrBuilder.BuildText(
			&strBuilder,
			&txtFormatCol,
			ePrefix.XCpy(
				"Final Output String"))

	if err != nil {
		outputStr =
			fmt.Sprintf("%v\n"+
				"Error: txtStrBuilder.BuildText()\n"+
				"Error=\n%v\n",
				ePrefix.String(),
				err.Error())

		fmt.Printf(outputStr + "\n")

		return
	}

	fmt.Printf(strBuilder.String() + "\n")

	strBuilder.Reset()

	return

}

func (mTest02 MainTest02) ExtractNumberRunes01() {

	funcName := "MainTest02.ExtractNumberRunes01()"

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		funcName,
		"")

	var outputStr string

	strBuilder := strings.Builder{}

	sMech := strmech.StrMech{}

	numberStr := "123456.78  0123 "
	characterSearchLen := 9

	// Total available Length of Output Line
	const maxLineLen = 78

	var err error

	mt02Nanobot := mainTest02Nanobot{}

	err = mt02Nanobot.SetLeadingMarquee(
		&strBuilder,
		funcName,
		"Testing StrMech.ExtractNumberRunes()",
		maxLineLen,
		false,
		&ePrefix)

	if err != nil {
		fmt.Println(
			fmt.Sprintf("%v",
				err.Error()))

		return
	}

	var searchResults strmech.CharSearchNumStrParseResultsDto
	var numStrKernel strmech.NumberStrKernel

	decimalSeparatorSpec := strmech.DecimalSeparatorSpec{}
	numParsingTerminators := strmech.RuneArrayCollection{}
	negativeNumSearchSpecs := strmech.NegNumSearchSpecCollection{}

	err = decimalSeparatorSpec.SetDecimalSeparatorStr(
		".",
		ePrefix.XCpy("Decimal Point '.'"))

	if err != nil {
		outputStr =
			fmt.Sprintf("%v\n"+
				"Error: decimalSeparatorSpec.SetDecimalSeparatorStr()\n"+
				"Error=\n%v\n",
				ePrefix.String(),
				err.Error())

		fmt.Printf(outputStr + "\n")

		return
	}

	err = negativeNumSearchSpecs.AddLeadingNegNumSearchStr(
		"-",
		ePrefix.XCpy("Leading Minus Sign '-'"))

	if err != nil {
		outputStr =
			fmt.Sprintf("%v\n"+
				"Error: negativeNumSearchSpecs.AddLeadingNegNumSearchStr()\n"+
				"Error=\n%v\n",
				ePrefix.String(),
				err.Error())

		fmt.Printf(outputStr + "\n")

		return
	}

	runeArrayDto := strmech.RuneArrayDto{
		CharsArray:   []rune(numberStr),
		Description1: "",
		Description2: "",
	}

	searchResults,
		numStrKernel,
		err = sMech.ExtractNumberRunes(
		runeArrayDto,
		0,
		characterSearchLen,
		negativeNumSearchSpecs,
		decimalSeparatorSpec,
		numParsingTerminators,
		false,
		ePrefix.XCpy(
			numberStr))

	if err != nil {
		outputStr =
			fmt.Sprintf("%v\n"+
				"Error: sMech.ExtractNumberRunes()\n"+
				"Error=\n\n%v\n",
				ePrefix.String(),
				err.Error())

		fmt.Printf(outputStr + "\n")

		return
	}

	err = searchResults.GetParameterTextListing(
		&strBuilder,
		false,
		true,
		ePrefix.XCpy("searchResults"))

	if err != nil {
		outputStr =
			fmt.Sprintf("%v\n"+
				"Error: searchResults.GetParameterTextListing()\n"+
				"Error=\n\n%v\n",
				ePrefix.String(),
				err.Error())

		fmt.Printf(outputStr + "\n")

		return
	}

	fmt.Printf(strBuilder.String() + "\n")
	strBuilder.Reset()

	err = numStrKernel.GetParameterTextListing(
		&strBuilder,
		true,
		ePrefix.XCpy(
			"numStrKernel"))

	if err != nil {
		outputStr =
			fmt.Sprintf("%v\n"+
				"Error: searchResults.GetParameterTextListing()\n"+
				"Error=\n\n%v\n",
				ePrefix.String(),
				err.Error())

		fmt.Printf(outputStr + "\n")

		return
	}

	// Trailing Title Marquee
	err = mt02Nanobot.SetTrailingMarquee(
		&strBuilder,
		funcName,
		"SUCCESSFUL COMPLETION!",
		maxLineLen,
		false,
		ePrefix.XCpy(
			"Trailing Title Marquee"))

	if err != nil {
		fmt.Println(
			fmt.Sprintf("%v",
				err.Error()))

		return
	}

	fmt.Printf(strBuilder.String() + "\n")

	fmt.Printf("\n\n String Builder Stats\n")
	fmt.Println("=========================================")
	fmt.Printf("      String Builder Capacity: %v\n",
		strBuilder.Cap())
	fmt.Printf("        String Builder Length: %v\n",
		strBuilder.Len())
	fmt.Println("-----------------------------------------")
	fmt.Printf("  String Builder Net Capacity: %v\n\n",
		strBuilder.Cap()-strBuilder.Len())

	strBuilder.Reset()

	return
}

func (mTest02 MainTest02) ExtractNumberRunes02() {

	funcName := "MainTest02.ExtractNumberRunes02()"

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		funcName,
		"")

	var outputStr string

	strBuilder := strings.Builder{}

	sMech := strmech.StrMech{}

	numberStr := " - 1.234 "

	// Total available Length of Output Line
	const maxLineLen = 78

	var err error

	// Leading Title Marquee

	mt02Nanobot := mainTest02Nanobot{}

	err = mt02Nanobot.SetLeadingMarquee(
		&strBuilder,
		funcName,
		"Testing StrMech.ExtractNumberRunes()",
		maxLineLen,
		false,
		&ePrefix)

	if err != nil {
		fmt.Println(
			fmt.Sprintf("%v",
				err.Error()))

		return
	}

	// End Of Marquee

	var searchResults strmech.CharSearchNumStrParseResultsDto
	var numStrKernel strmech.NumberStrKernel

	decimalSeparatorSpec := strmech.DecimalSeparatorSpec{}
	numParsingTerminators := strmech.RuneArrayCollection{}
	negativeNumSearchSpecs := strmech.NegNumSearchSpecCollection{}

	err = decimalSeparatorSpec.SetDecimalSeparatorStr(
		".",
		ePrefix.XCpy("Decimal Point '.'"))

	if err != nil {
		outputStr =
			fmt.Sprintf("%v\n"+
				"Error: decimalSeparatorSpec.SetDecimalSeparatorStr()\n"+
				"Error=\n%v\n",
				ePrefix.String(),
				err.Error())

		fmt.Printf(outputStr + "\n")

		return
	}

	err = negativeNumSearchSpecs.AddLeadingNegNumSearchStr(
		"-",
		ePrefix.XCpy("Leading Minus Sign '-'"))

	if err != nil {
		outputStr =
			fmt.Sprintf("%v\n"+
				"Error: negativeNumSearchSpecs.AddLeadingNegNumSearchStr()\n"+
				"Error=\n%v\n",
				ePrefix.String(),
				err.Error())

		fmt.Printf(outputStr + "\n")

		return
	}

	runeArrayDto := strmech.RuneArrayDto{
		CharsArray:   []rune(numberStr),
		Description1: "",
		Description2: "",
	}

	searchResults,
		numStrKernel,
		err = sMech.ExtractNumberRunes(
		runeArrayDto,
		0,
		-1,
		negativeNumSearchSpecs,
		decimalSeparatorSpec,
		numParsingTerminators,
		false,
		ePrefix.XCpy(
			numberStr))

	if err != nil {
		outputStr =
			fmt.Sprintf("%v\n"+
				"Error: sMech.ExtractNumberRunes()\n"+
				"Error=\n\n%v\n",
				ePrefix.String(),
				err.Error())

		fmt.Printf(outputStr + "\n")

		return
	}

	err = searchResults.GetParameterTextListing(
		&strBuilder,
		false,
		true,
		ePrefix.XCpy("searchResults"))

	if err != nil {
		outputStr =
			fmt.Sprintf("%v\n"+
				"Error: searchResults.GetParameterTextListing()\n"+
				"Error=\n\n%v\n",
				ePrefix.String(),
				err.Error())

		fmt.Printf(outputStr + "\n")

		return
	}

	fmt.Printf(strBuilder.String() + "\n")
	strBuilder.Reset()

	err = numStrKernel.GetParameterTextListing(
		&strBuilder,
		true,
		ePrefix.XCpy(
			"numStrKernel"))

	if err != nil {
		outputStr =
			fmt.Sprintf("%v\n"+
				"Error: searchResults.GetParameterTextListing()\n"+
				"Error=\n\n%v\n",
				ePrefix.String(),
				err.Error())

		fmt.Printf(outputStr + "\n")

		return
	}

	// Trailing Title Marquee
	// Top Blank Line

	err = mt02Nanobot.SetTrailingMarquee(
		&strBuilder,
		funcName,
		"SUCCESSFUL COMPLETION!",
		maxLineLen,
		false,
		ePrefix.XCpy(
			"Trailing Title Marquee"))

	if err != nil {
		fmt.Println(
			fmt.Sprintf("%v",
				err.Error()))

		return
	}

	fmt.Printf(strBuilder.String() + "\n")

	fmt.Printf(strBuilder.String() + "\n")

	fmt.Printf("\n\n String Builder Stats\n")
	fmt.Println("=========================================")
	fmt.Printf("      String Builder Capacity: %v\n",
		strBuilder.Cap())
	fmt.Printf("        String Builder Length: %v\n",
		strBuilder.Len())
	fmt.Println("-----------------------------------------")
	fmt.Printf("  String Builder Net Capacity: %v\n\n",
		strBuilder.Cap()-strBuilder.Len())

	strBuilder.Reset()

	return
}

type mainTest02Nanobot struct {
	lock *sync.Mutex
}

func (mt02Nanobot *mainTest02Nanobot) SetLeadingMarquee(
	strBuilder *strings.Builder,
	funcName string,
	testTitle string,
	maxLineLen int,
	turnAutoLineLengthBreaksOn bool,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if mt02Nanobot.lock == nil {
		mt02Nanobot.lock = new(sync.Mutex)
	}

	mt02Nanobot.lock.Lock()

	defer mt02Nanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"mainTest02Nanobot."+
			"SetLeadingMarquee()",
		"")

	if err != nil {

		return err

	}

	if strBuilder == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'strBuilder' is a 'nil' pointer.\n",
			ePrefix.String())

		return err

	}

	if len(funcName) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'funcName' is an empty string.\n",
			ePrefix.String())

		return err
	}

	if len(testTitle) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'testTitle' is an empty string.\n",
			ePrefix.String())

		return err
	}

	if maxLineLen < 10 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'maxLineLen' has a value less than ten (10).\n"+
			"maxLineLen = '%v'\n",
			ePrefix.String(),
			maxLineLen)

		return err
	}

	titles := strmech.StringArrayDto{}

	titles.AddString(funcName)

	titles.AddString(testTitle)

	// Standard Format Parameter

	txtFmtParams := strmech.TextFmtParamsLineColumnsDto{
		FormatType: strmech.TxtFieldType.LineColumns(),
		FieldFormatParams: []strmech.TextFieldFmtParamsDto{
			{
				LeftMarginStr:  " ",
				FieldLength:    maxLineLen,
				FieldJustify:   strmech.TxtJustify.Center(),
				DateTimeFormat: "Monday 2006-01-02 15:04:05.000000000 -0700 MST",
				RightMarginStr: "",
			},
		},
		TurnLineTerminationOff:     false,
		LineTerminator:             "\n",
		MaxLineLength:              maxLineLen,
		TurnAutoLineLengthBreaksOn: turnAutoLineLengthBreaksOn,
	}

	// Leading Title Marquee

	err = new(strmech.TextUtility).BuildOneColLeadingMarquee(
		strBuilder,
		titles,
		txtFmtParams,
		ePrefix.XCpy(
			"strBuilder<-txtFmtParams"))

	return err
}

func (mt02Nanobot *mainTest02Nanobot) SetTrailingMarquee(
	strBuilder *strings.Builder,
	funcName string,
	completionDesc string,
	maxLineLen int,
	turnAutoLineLengthBreaksOn bool,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if mt02Nanobot.lock == nil {
		mt02Nanobot.lock = new(sync.Mutex)
	}

	mt02Nanobot.lock.Lock()

	defer mt02Nanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"mainTest02Nanobot."+
			"SetLeadingMarquee()",
		"")

	if err != nil {

		return err

	}

	if len(funcName) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'funcName' is an empty string.\n",
			ePrefix.String())

		return err
	}

	if len(completionDesc) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'completionDesc' is an empty string.\n",
			ePrefix.String())

		return err
	}

	if maxLineLen < 10 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'maxLineLen' has a value less than ten (10).\n"+
			"maxLineLen = '%v'\n",
			ePrefix.String(),
			maxLineLen)

		return err
	}

	titles := strmech.StringArrayDto{}

	titles.AddString(funcName)

	titles.AddString(completionDesc)

	// Standard Format Parameter
	txtFmtParams := strmech.TextFmtParamsLineColumnsDto{
		FormatType: strmech.TxtFieldType.LineColumns(),
		FieldFormatParams: []strmech.TextFieldFmtParamsDto{
			{
				LeftMarginStr:  " ",
				FieldLength:    maxLineLen,
				FieldJustify:   strmech.TxtJustify.Center(),
				DateTimeFormat: "Monday 2006-01-02 15:04:05.000000000 -0700 MST",
				RightMarginStr: "",
			},
		},
		TurnLineTerminationOff:     false,
		LineTerminator:             "\n",
		MaxLineLength:              maxLineLen,
		TurnAutoLineLengthBreaksOn: turnAutoLineLengthBreaksOn,
	}

	// Trailing Title Marquee
	err = new(strmech.TextUtility).BuildOneColTrailingMarquee(
		strBuilder,
		titles,
		txtFmtParams,
		ePrefix.XCpy(
			"strBuilder<-txtFmtParams"))

	// End Of Marquee

	return err
}
