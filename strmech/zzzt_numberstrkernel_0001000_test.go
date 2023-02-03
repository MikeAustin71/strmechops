package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"strings"
	"testing"
)

func TestNumberStrKernel_Compare_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_Compare_000100()",
		"")

	origIntStr := "1234"
	origFracStr := "5678"

	var err error
	var baseValue NumberStrKernel
	var intDigitsDto, fracDigitsDto RuneArrayDto

	intDigitsDto,
		err = new(RuneArrayDto).NewString(
		origIntStr,
		CharSearchType.LinearTargetStartingIndex(),
		ePrefix.XCpy(
			"origIntStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	fracDigitsDto,
		err = new(RuneArrayDto).NewString(
		origFracStr,
		CharSearchType.LinearTargetStartingIndex(),
		ePrefix.XCpy(
			"origIntStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	baseValue,
		err = new(NumberStrKernel).NewFromRuneDto(
		&intDigitsDto,
		&fracDigitsDto,
		NumSignVal.Positive(),
		ePrefix.XCpy(
			"baseValue<-"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var testNStrValue01 NumberStrKernel

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
		t.Errorf("\n%v\n"+
			"Error return from testBigFloat.Parse(testValue,10)\n"+
			"testValue = '%v'\n",
			ePrefix.String(),
			testValue)

		return
	}

	minPrecision := testBigFloat.MinPrec()

	testBigFloat.SetPrec(minPrecision)

	testNStrValue01,
		err = new(NumberStrKernel).NewFromNumericValue(
		testBigFloat,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"testNStrValue01<-testBigFloat"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var actualNumSign NumericSignValueType

	actualNumSign,
		err = testNStrValue01.GetNumberSign(
		ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumSign != NumSignVal.Positive() {

		t.Errorf("\n%v\n"+
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
		t.Errorf("\n%v\n"+
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

		t.Errorf("\n%v\n"+
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
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if comparisonResult != -1 {
		t.Errorf("\n%v\n"+
			"Test#7\n"+
			"Error: Expected a comparisonResult of -1.\n"+
			"Instead, comparisonResult = '%v'\n",
			ePrefix.String(),
			comparisonResult)

		return
	}

	var testNStrValue02 NumberStrKernel

	testValueIntDigits = "234"
	testValueFracDigits = "5678"

	testValue = "-" +
		testValueIntDigits +
		"." +
		testValueFracDigits

	testNStrValue02,
		_,
		err = new(NumberStrKernel).NewParsePureNumberStr(
		testValue,
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"testNStrValue02<-testValue"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualIntStr = testNStrValue02.GetIntegerString()

	if actualIntStr != testValueIntDigits {
		t.Errorf("\n%v\n"+
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

		t.Errorf("\n%v\n"+
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
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumSign != NumSignVal.Negative() {

		t.Errorf("\n%v\n"+
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
			"<-testNStrValue02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if comparisonResult != 1 {
		t.Errorf("\n%v\n"+
			"Test#10\n"+
			"Error: Expected a comparisonResult of +1.\n"+
			"Instead, comparisonResult = '%v'\n",
			ePrefix.String(),
			comparisonResult)
	}

	var testNStrValue03 NumberStrKernel

	testValueIntDigits = "1234"
	testValueFracDigits = "5679"

	testValue = testValueIntDigits +
		"." +
		testValueFracDigits

	testNStrValue03,
		_,
		err = new(NumberStrKernel).NewParsePureNumberStr(
		testValue,
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"testNStrValue03<-testValue"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualIntStr = testNStrValue03.GetIntegerString()

	if actualIntStr != testValueIntDigits {
		t.Errorf("\n%v\n"+
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

		t.Errorf("\n%v\n"+
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
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumSign != NumSignVal.Positive() {

		t.Errorf("\n%v\n"+
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
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if comparisonResult != -1 {
		t.Errorf("\n%v\n"+
			"Test#14\n"+
			"Error: Expected a comparisonResult of -1.\n"+
			"Instead, comparisonResult = '%v'\n",
			ePrefix.String(),
			comparisonResult)
	}

	var testNStrValue04 NumberStrKernel

	origIntStr = "5678"
	origFracStr = "9012"
	origValueStr := "-" +
		origIntStr +
		"." +
		origFracStr

	baseValue,
		_,
		err = new(NumberStrKernel).NewParsePureNumberStr(
		origValueStr,
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"baseValue<-origValueStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumSign,
		err = baseValue.GetNumberSign(
		ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumSign != NumSignVal.Negative() {

		t.Errorf("\n%v\n"+
			"Test #17 baseValue\n"+
			"Error: Expected Number Sign Equal To 'Negative'.\n"+
			"Instead, Number Sign Equals Integer Value ='%v'"+
			"Actual Number Sign string value = '%v'\n",
			ePrefix.String(),
			actualNumSign.XValueInt(),
			actualNumSign.String())

		return
	}

	testValueIntDigits = "5678"
	testValueFracDigits = "9013"

	testValue = "-" +
		testValueIntDigits +
		"." +
		testValueFracDigits

	testNStrValue04,
		_,
		err = new(NumberStrKernel).NewParsePureNumberStr(
		testValue,
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"testNStrValue04<-testValue"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualIntStr = testNStrValue04.GetIntegerString()

	if actualIntStr != testValueIntDigits {
		t.Errorf("\n%v\n"+
			"Test#15\n"+
			"Error: actualIntStr != testValueIntDigits\n"+
			"actualIntStr         = '%v'\n"+
			"testValueIntDigits   = '%v'\n",
			ePrefix.String(),
			actualIntStr,
			testValueIntDigits)

		return
	}

	actualFracStr = testNStrValue04.GetFractionalString()

	if actualFracStr != testValueFracDigits {

		t.Errorf("\n%v\n"+
			"Test#16\n"+
			"Error: actualFracStr != testValueFracDigits\n"+
			"actualFracStr         = '%v'\n"+
			"testValueFracDigits   = '%v'\n",
			ePrefix.String(),
			actualFracStr,
			testValueFracDigits)

		return
	}

	actualNumSign,
		err = testNStrValue04.GetNumberSign(
		ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumSign != NumSignVal.Negative() {

		t.Errorf("\n%v\n"+
			"Test #17\n"+
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
		&testNStrValue04,
		ePrefix.XCpy(
			"<-testNStrValue04"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if comparisonResult != 1 {
		t.Errorf("\n%v\n"+
			"Test#18\n"+
			"Error: Expected a comparisonResult of +1.\n"+
			"Instead, comparisonResult = '%v'\n",
			ePrefix.String(),
			comparisonResult)
	}

	return
}

func TestNumberStrKernel_DeleteLeadingTrailingFractionalChars_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"\nTestNumberStrKernel_DeleteLeadingTrailingFractionalChars_000100()",
		"")

	origIntStr := "1234"
	origFracStr := "5678"

	var err error
	var numStrKernel01 NumberStrKernel

	numStrKernel01,
		err = new(NumberStrKernel).NewFromStringDigits(
		origIntStr,
		origFracStr,
		NumSignVal.Positive(),
		ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualIntStr := numStrKernel01.GetIntegerString()

	if actualIntStr != origIntStr {
		t.Errorf("%v\n"+
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

		t.Errorf("%v\n"+
			"Test#2\n"+
			"Error: actualFracStr != origFracStr\n"+
			"actualFracStr = '%v'\n"+
			"origFracStr   = '%v'\n",
			ePrefix.String(),
			actualFracStr,
			origFracStr)

		return
	}

	err = numStrKernel01.DeleteLeadingTrailingIntegerChars(
		2,
		false,
		ePrefix.XCpy(
			"Delete Leading Chars"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	//origIntStr := "1234"

	expectedIntStr := "34"

	actualIntStr = numStrKernel01.GetIntegerString()

	if actualIntStr != expectedIntStr {

		t.Errorf("%v\n"+
			"Test#3\n"+
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

	//origFracStr := "5678"

	err = numStrKernel01.DeleteLeadingTrailingFractionalChars(
		3,
		true,
		ePrefix.XCpy(
			"Delete Trailing Chars"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	expectedFracStr := "5"

	actualFracStr = numStrKernel01.GetFractionalString()

	if actualFracStr != expectedFracStr {

		t.Errorf("%v\n"+
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

	return
}

func TestNumberStrKernel_DeleteLeadingTrailingFractionalChars_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"\nTestNumberStrKernel_DeleteLeadingTrailingFractionalChars_000200()",
		"")

	origIntStr := "1234"
	origFracStr := ""

	var err error
	var numStrKernel01 NumberStrKernel

	numStrKernel01,
		err = new(NumberStrKernel).NewFromRuneDigits(
		[]rune(origIntStr),
		[]rune(origFracStr),
		NumSignVal.Positive(),
		&ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualIntStr := numStrKernel01.GetIntegerString()

	if actualIntStr != origIntStr {
		t.Errorf("%v\n"+
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

		t.Errorf("%v\n"+
			"Test#2\n"+
			"Error: actualFracStr != origFracStr\n"+
			"actualFracStr = '%v'\n"+
			"origFracStr   = '%v'\n",
			ePrefix.String(),
			actualFracStr,
			origFracStr)

		return
	}

	err = numStrKernel01.DeleteLeadingTrailingIntegerChars(
		uint64(len(origIntStr)),
		false,
		ePrefix.XCpy(
			"Delete Leading Chars"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	//origIntStr := "1234"

	expectedIntStr := ""

	actualIntStr = numStrKernel01.GetIntegerString()

	if actualIntStr != expectedIntStr {

		t.Errorf("%v\n"+
			"Test#3\n"+
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

	return
}

func TestNumberStrKernel_ExtendIntegerDigitsArray_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"\nTestNumberStrKernel_ExtendIntegerDigitsArray_000100()",
		"")

	origIntStr := "1234"
	origFracStr := "5678"
	origNumStr := origIntStr + "." + origFracStr

	var err error
	var numStrKernel01 NumberStrKernel
	var numberStrSearchResults CharSearchNumStrParseResultsDto

	numberStrSearchResults,
		numStrKernel01,
		err = new(NumberStrKernel).NewParseUSNumberStr(
		origNumStr,
		0,
		-1,
		nil,
		false,
		&ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !numberStrSearchResults.FoundIntegerDigits {

		t.Errorf("%v\n"+
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

		t.Errorf("%v\n"+
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
		t.Errorf("%v\n"+
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

		t.Errorf("%v\n"+
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
		t.Errorf("%v\n",
			err.Error())
		return
	}

	//origIntStr := "1234"

	expectedIntStr := "1234000"

	actualIntStr = numStrKernel01.GetIntegerString()

	if actualIntStr != expectedIntStr {

		t.Errorf("%v\n"+
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
		t.Errorf("%v\n",
			err.Error())
		return
	}

	//origFracStr := "5678"
	expectedFracStr := "0005678"

	actualFracStr = numStrKernel01.GetFractionalString()

	if actualFracStr != expectedFracStr {

		t.Errorf("%v\n"+
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

	return
}

func TestNumberStrKernel_ExtendIntegerDigitsArray_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"\nTestNumberStrKernel_ExtendIntegerDigitsArray_000200()",
		"")

	origIntStr := "1234"
	origFracStr := "5678"
	origNumStr := origIntStr + "." + origFracStr

	var err error
	var numStrKernel01 NumberStrKernel
	var numberStrSearchResults CharSearchNumStrParseResultsDto

	numberStrSearchResults,
		numStrKernel01,
		err = new(NumberStrKernel).NewParseUSNumberStr(
		origNumStr,
		0,
		-1,
		nil,
		false,
		&ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !numberStrSearchResults.FoundIntegerDigits {

		t.Errorf("%v\n"+
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

		t.Errorf("%v\n"+
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
		t.Errorf("%v\n"+
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

		t.Errorf("%v\n"+
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
		false,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	//origIntStr := "1234"

	expectedIntStr := "0001234"

	actualIntStr = numStrKernel01.GetIntegerString()

	if actualIntStr != expectedIntStr {

		t.Errorf("%v\n"+
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
		true,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	//origFracStr := "5678"
	expectedFracStr := "5678000"

	actualFracStr = numStrKernel01.GetFractionalString()

	if actualFracStr != expectedFracStr {

		t.Errorf("%v\n"+
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

	err = numStrKernel01.ExtendIntegerDigitsArray(
		'x',
		3,
		false,
		ePrefix)

	if err == nil {

		t.Errorf("%v\n"+
			"Error: Expected an error return from\n"+
			"numStrKernel01.ExtendIntegerDigitsArray()\n"+
			"because the rune digit to add is non-numeric.\n"+
			"'numCharToAdd' == 'x'\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!!\n",
			ePrefix)

		return
	}

	err = numStrKernel01.ExtendFractionalDigitsArray(
		'!',
		3,
		false,
		ePrefix)

	if err == nil {

		t.Errorf("%v\n"+
			"Error: Expected an error return from\n"+
			"numStrKernel01.ExtendIntegerDigitsArray()\n"+
			"because the rune digit to add is non-numeric.\n"+
			"'numCharToAdd' == '!'\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!!\n",
			ePrefix)

		return
	}

	return
}

func TestNumberStrKernel_FmtCharReplacementStr_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_FmtCharReplacementStr_000100()",
		"")

	var nStrKernel01 NumberStrKernel
	var err error

	intDigits := "0115550101"
	expectedStr := "(011) 555-0101"

	nStrKernel01,
		err = new(NumberStrKernel).
		NewFromStringDigits(
			intDigits,
			"",
			NumSignVal.Positive(),
			ePrefix.XCpy(
				"Test #1-A: nStrKernel01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	testName := fmt.Sprintf("Test #1-A FmtCharReplacementStr()\n"+
		"intDigits Str(%v)\n",
		intDigits)

	numFmtSpec := NumStrFmtCharReplacementSpec{
		NumberFormat:       "(NNN) NNN-NNNN",
		NumReplacementChar: 'N',
	}

	var formattedNumStr, remainingIntFracDigits string

	formattedNumStr,
		remainingIntFracDigits,
		err = nStrKernel01.FmtCharReplacementStr(
		numFmtSpec,
		ePrefix.XCpy(
			"Test#1-A nStrKernel01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if formattedNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: formattedNumStr != expectedStr\n"+
			"formattedNumStr  = '%v'\n"+
			"   expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			formattedNumStr,
			expectedStr)

		return

	}

	testName = fmt.Sprintf("Test #1-B remainingIntFracDigits()\n"+
		"remainingIntFracDigits = '%v'\n",
		remainingIntFracDigits)

	if len(remainingIntFracDigits) != 0 {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: len(remainingIntFracDigits) > 0\n"+
			"formattedNumStr  = '%v'\n"+
			"   expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			formattedNumStr,
			expectedStr)

		return

	}

	intDigits = "0115550101"
	fracDigits := "4128"

	var nStrKernel02 NumberStrKernel

	nStrKernel02,
		err = new(NumberStrKernel).
		NewFromStringDigits(
			intDigits,
			fracDigits,
			NumSignVal.Positive(),
			ePrefix.XCpy(
				"Test #2-A: nStrKernel02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	testName = fmt.Sprintf("Test #2-A FmtCharReplacementStr()\n"+
		" intDigits Str = %v\n"+
		"fracDigits Str = %v\n",
		intDigits,
		fracDigits)

	numFmtSpec.NumberFormat = "(ZZZ) ZZZ-ZZZZ"
	numFmtSpec.NumReplacementChar = 'Z'

	formattedNumStr = ""
	remainingIntFracDigits = ""

	formattedNumStr,
		remainingIntFracDigits,
		err = nStrKernel02.FmtCharReplacementStr(
		numFmtSpec,
		ePrefix.XCpy(
			"Test#2-A nStrKernel02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if formattedNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: formattedNumStr != expectedStr\n"+
			"formattedNumStr  = '%v'\n"+
			"   expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			formattedNumStr,
			expectedStr)

		return

	}

	testName = fmt.Sprintf("Test #2-B remainingIntFracDigits()\n"+
		"remainingIntFracDigits = '%v'\n",
		remainingIntFracDigits)

	expectedStr = "4128"

	if remainingIntFracDigits != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: remainingIntFracDigits != expectedStr\n"+
			"remainingIntFracDigits  = '%v'\n"+
			"          expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			remainingIntFracDigits,
			expectedStr)

		return

	}

	var nStrKernel03 NumberStrKernel

	intDigits = "011555010"
	fracDigits = ""

	nStrKernel03,
		err = new(NumberStrKernel).
		NewFromStringDigits(
			intDigits,
			"",
			NumSignVal.Positive(),
			ePrefix.XCpy(
				"Test #3-A: nStrKernel03"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	testName = fmt.Sprintf("Test #2-A FmtCharReplacementStr()\n"+
		"Insufficient Integer Digits\n"+
		" intDigits Str = %v\n"+
		"fracDigits Str = %v\n",
		intDigits,
		fracDigits)

	numFmtSpec.NumberFormat = "(ZZZ) ZZZ-ZZZZ"
	numFmtSpec.NumReplacementChar = 'Z'

	formattedNumStr,
		remainingIntFracDigits,
		err = nStrKernel03.FmtCharReplacementStr(
		numFmtSpec,
		ePrefix.XCpy(
			"Test#3-A nStrKernel03"))

	if err == nil {
		t.Errorf("\n%v\n"+
			"%v\n"+
			"Expected an error return from\n"+
			"nStrKernel02.FmtCharReplacementStr()\n"+
			"because there were insufficient numeric digits\n"+
			"to fill the target string.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String(),
			testName)

		return
	}

	var nStrKernel04 NumberStrKernel

	intDigits = "0115550101"
	fracDigits = "4128"

	nStrKernel04,
		err = new(NumberStrKernel).
		NewFromStringDigits(
			intDigits,
			fracDigits,
			NumSignVal.Positive(),
			ePrefix.XCpy(
				"Test #4-A: nStrKernel04"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	numFmtSpec.NumberFormat = "(ZZZ) ZZZ-ZZZZ"
	numFmtSpec.NumReplacementChar = 'X'

	testName = fmt.Sprintf("Test #4-A: Invalid NumStrFmtCharReplacementSpec\n"+
		" numFmtSpec.NumberFormat = %v\n"+
		"numFmtSpec.NumReplacementChar = %v",
		numFmtSpec.NumberFormat,
		string(numFmtSpec.NumReplacementChar))

	formattedNumStr,
		remainingIntFracDigits,
		err = nStrKernel04.FmtCharReplacementStr(
		numFmtSpec,
		ePrefix.XCpy(
			"Test#4-A nStrKernel04"))

	if err == nil {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Expected an error return from\n"+
			"nStrKernel04.FmtCharReplacementStr()\n"+
			"because NumStrFmtCharReplacementSpec is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String(),
			testName)

		return
	}

	return
}

func TestNumberStrKernel_FmtCharReplacementStr_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_FmtCharReplacementStr_000200()",
		"")

	var nStrKernel01 NumberStrKernel
	var err error

	intDigitsStr := "0115550101"
	fracDigitsStr := "4567"
	expectedStr := "(011) 555-0101"

	nStrKernel01,
		err = new(NumberStrKernel).
		NewFromRuneDigits(
			[]rune(intDigitsStr),
			[]rune(fracDigitsStr),
			NumSignVal.Positive(),
			ePrefix.XCpy(
				"Test #1-A: nStrKernel01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	testName := fmt.Sprintf("Test #1-A FmtCharReplacementStr()\n"+
		"intDigitsStr = (%v)\n",
		intDigitsStr)

	numFmtSpec := NumStrFmtCharReplacementSpec{
		NumberFormat:       "(NNN) NNN-NNNN",
		NumReplacementChar: 'N',
	}

	var formattedNumStr, remainingIntFracDigits string

	formattedNumStr,
		remainingIntFracDigits,
		err = nStrKernel01.FmtCharReplacementStr(
		numFmtSpec,
		ePrefix.XCpy(
			"Test#1-A nStrKernel01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if formattedNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: formattedNumStr != expectedStr\n"+
			"formattedNumStr  = '%v'\n"+
			"   expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			formattedNumStr,
			expectedStr)

		return

	}

	expectedStr = "4567"

	testName = fmt.Sprintf("Test #1-B Remaining IntFrac Digits\n"+
		"expectedStr = %v\n",
		expectedStr)

	if remainingIntFracDigits != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: remainingIntFracDigits != expectedStr\n"+
			"remainingIntFracDigits  = '%v'\n"+
			"          expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			remainingIntFracDigits,
			expectedStr)

		return

	}

	return
}

func TestNumberStrKernel_FmtCurrencyDefaultsFrance_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_FmtCurrencyDefaultsFrance_000100()",
		"")

	var intDigits, fracDigits string

	var err error

	intDigits = "1234"

	fracDigits = "56"

	numSign := NumSignVal.Negative()

	var numStrKernel NumberStrKernel

	err = numStrKernel.SetStringDigits(
		intDigits,
		fracDigits,
		numSign,
		ePrefix.XCpy(
			"numStrKernel"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	origNumStrValue := "-1234.56"

	expectedNumberStr := origNumStrValue

	var actualNumberStr string

	actualNumberStr,
		_,
		err = numStrKernel.FmtNumStrNative(
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"numStrKernel Test#1"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	testName := fmt.Sprintf("Test #1 Original Number String Setup Verification\n"+
		"Original Native Number String Value = %v\n",
		origNumStrValue)

	if expectedNumberStr != actualNumberStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumberStr,
			expectedNumberStr)

		return
	}

	testName = fmt.Sprintf("Test #2 French Default Currency\n"+
		"Original Native Number String Value = %v\n",
		origNumStrValue)

	expectedNumberStr = "    -1 234,56 €"

	roundingSpec := NumStrRoundingSpec{
		roundingType:            NumRoundType.NoRounding(),
		roundToFractionalDigits: 0,
		lock:                    nil,
	}

	numberFieldSpec := NumStrNumberFieldSpec{
		fieldLength:        15,
		fieldJustification: TxtJustify.Right(),
		lock:               nil,
	}

	actualNumberStr,
		err = numStrKernel.FmtCurrencyDefaultsFrance(
		roundingSpec,
		numberFieldSpec,
		ePrefix.XCpy(
			"numStrKernel Test#2"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != actualNumberStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumberStr,
			expectedNumberStr)

		return
	}
}

func TestNumberStrKernel_FmtCurrencyDefaultsGermany_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_FmtCurrencyDefaultsGermany_000100()",
		"")

	var intDigits, fracDigits RuneArrayDto

	var err error

	intDigits,
		err = new(RuneArrayDto).NewRunes(
		[]rune("1234"),
		CharSearchType.LinearTargetStartingIndex(),
		ePrefix.XCpy(
			"intDigits"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	fracDigits,
		err = new(RuneArrayDto).NewRunes(
		[]rune("56"),
		CharSearchType.LinearTargetStartingIndex(),
		ePrefix.XCpy(
			"intDigits"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	numSign := NumSignVal.Negative()

	var numStrKernel NumberStrKernel

	err = numStrKernel.SetRuneDto(
		&intDigits,
		&fracDigits,
		numSign,
		ePrefix.XCpy(
			"numStrKernel"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	origNumStrValue := "-1234.56"

	expectedNumberStr := origNumStrValue

	var actualNumberStr string

	actualNumberStr,
		_,
		err = numStrKernel.FmtNumStrNative(
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"numStrKernel Test#1"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	testName := fmt.Sprintf("Test #1 Original Number String Setup Verification\n"+
		"Original Native Number String Value = %v\n",
		origNumStrValue)

	if expectedNumberStr != actualNumberStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumberStr,
			expectedNumberStr)

		return
	}

	testName = fmt.Sprintf("Test #2 German Default Currency\n"+
		"Original Native Number String Value = %v\n",
		origNumStrValue)

	expectedNumberStr = "    1.234,56- €"

	roundingSpec := NumStrRoundingSpec{
		roundingType:            NumRoundType.NoRounding(),
		roundToFractionalDigits: 0,
		lock:                    nil,
	}

	numberFieldSpec := NumStrNumberFieldSpec{
		fieldLength:        15,
		fieldJustification: TxtJustify.Right(),
		lock:               nil,
	}

	actualNumberStr,
		err = numStrKernel.FmtCurrencyDefaultsGermany(
		roundingSpec,
		numberFieldSpec,
		ePrefix.XCpy(
			"numStrKernel Test#2"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != actualNumberStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumberStr,
			expectedNumberStr)

		return
	}
}

func TestNumberStrKernel_FmtCurrencyDefaultsUKMinusInside_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_FmtCurrencyDefaultsUKMinusInside_000100()",
		"")

	intDigits := []rune("1234")

	fracDigits := []rune("56")

	numSign := NumSignVal.Negative()

	var err error

	var numStrKernel NumberStrKernel

	err = numStrKernel.SetRuneDigits(
		intDigits,
		fracDigits,
		numSign,
		ePrefix.XCpy(
			"numStrKernel"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	origNumStrValue := "-1234.56"

	expectedNumberStr := origNumStrValue

	var actualNumberStr string

	actualNumberStr,
		_,
		err = numStrKernel.FmtNumStrNative(
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"numStrKernel Test#1"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	testName := fmt.Sprintf("Test #1 Original Number String Setup Verification\n"+
		"Original Native Number String Value = %v\n",
		origNumStrValue)

	if expectedNumberStr != actualNumberStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumberStr,
			expectedNumberStr)

		return
	}

	testName = fmt.Sprintf("Test #2 UK Minus Inside\n"+
		"Original Native Number String Value = %v\n",
		origNumStrValue)

	expectedNumberStr = "    £ -1,234.56"

	roundingSpec := NumStrRoundingSpec{
		roundingType:            NumRoundType.NoRounding(),
		roundToFractionalDigits: 0,
		lock:                    nil,
	}

	numberFieldSpec := NumStrNumberFieldSpec{
		fieldLength:        15,
		fieldJustification: TxtJustify.Right(),
		lock:               nil,
	}

	actualNumberStr,
		err = numStrKernel.FmtCurrencyDefaultsUKMinusInside(
		roundingSpec,
		numberFieldSpec,
		ePrefix.XCpy(
			"numStrKernel Test#2"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != actualNumberStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumberStr,
			expectedNumberStr)

		return
	}

	testName = fmt.Sprintf("Test #3 UK Minus Outside\n"+
		"Original Native Number String Value = %v\n",
		origNumStrValue)

	expectedNumberStr = "    -£ 1,234.56"

	actualNumberStr,
		err = numStrKernel.FmtCurrencyDefaultsUKMinusOutside(
		roundingSpec,
		numberFieldSpec,
		ePrefix.XCpy(
			"numStrKernel Test#3"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != actualNumberStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumberStr,
			expectedNumberStr)

		return
	}

	return
}

func TestNumberStrKernel_FmtNumStrPure_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_FmtNumStrPure_000100",
		"")

	inputNumberStr := "7.12345678"

	inputNumFracDigits := len(inputNumberStr) -
		strings.Index(inputNumberStr, ".") -
		1
	expectedNumberStr := "7.12346"

	var err error
	var nStr01 NumberStrKernel

	nStr01,
		_,
		err = new(NumberStrKernel).
		NewParsePureNumberStr(
			inputNumberStr,
			".",
			true,
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy(
				"nStr01<-inputNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr01.IsValidInstanceError(
		ePrefix.XCpy(
			"nStr01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var nativeNumStrStats NumberStrStatsDto
	var actualNativeNumStr1 string

	actualNativeNumStr1,
		nativeNumStrStats,
		err = nStr01.FmtNumStrNative(
		NumRoundType.NoRounding(),
		0,
		ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	testName := "Test# 1-A Number Of Digits Test"

	if int(nativeNumStrStats.NumOfFractionalDigits) !=
		inputNumFracDigits {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: Expected Number Of Fractional Digits\n"+
			"NOT EQUAL TO Actual Number Of Fractional Digits!\n"+
			"  Actual Number Of Fractional Digits = '%v'\n"+
			"Expected Number Of Fractional Digits = '%v'\n",
			ePrefix.String(),
			testName,
			nativeNumStrStats.NumOfFractionalDigits,
			inputNumFracDigits)

		return

	}

	testName = "Test# 1-B Number String Test"

	var actualNativeNumStr1B string

	actualNativeNumStr1B,
		_,
		err = nStr01.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr01"))

	if inputNumberStr != actualNativeNumStr1B {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: Expected Input Number String\n"+
			"DOES NOT EQUAL Actual Number String!\n"+
			"Actual Number String = '%v'\n"+
			" Input Number String = '%v'\n",
			ePrefix.String(),
			testName,
			actualNativeNumStr1B,
			inputNumberStr)

		return

	}

	testName = "Test# 2-A Native Number String Test"

	if actualNativeNumStr1 != actualNativeNumStr1B {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: Actual Native Number String 1 DOES NOT EQUAL\n"+
			"Actual Native Number String 1-B!\n"+
			"  Actual Native Number String 1 = '%v'\n"+
			"Actual Native Number String 1-B = '%v'\n",
			ePrefix.String(),
			testName,
			actualNativeNumStr1,
			actualNativeNumStr1B)

	}

	testName = "Test# 2-B Expected Number String Test"

	var actualNativeNumStr2A string

	actualNativeNumStr2A,
		_,
		err = nStr01.FmtNumStrNative(
		NumRoundType.HalfAwayFromZero(),
		5,
		ePrefix.XCpy(
			"2-A nStr01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != actualNativeNumStr2A {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: Expected Number String DOES NOT EQUAL\n"+
			"Actual Native Number String 2-A !\n"+
			"Actual Native Number String 2-A = '%v'\n"+
			"Expected Number String = '%v'\n",
			ePrefix.String(),
			testName,
			actualNativeNumStr2A,
			expectedNumberStr)

		return

	}

	var nStr02 NumberStrKernel

	nStr02,
		_,
		err = new(NumberStrKernel).
		NewParsePureNumberStr(
			inputNumberStr,
			".",
			true,
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy(
				"nStr02<-inputNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr02.IsValidInstanceError(
		ePrefix.XCpy(
			"nStr02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var actualNativeNumStr3A,
		actualPureNumStr3 string

	actualNativeNumStr3A,
		_,
		err = nStr02.FmtNumStrNative(
		NumRoundType.HalfAwayFromZero(),
		5,
		ePrefix.XCpy(
			"actualNativeNumStr3A<-nStr02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var pureNumStrComponents PureNumberStrComponents

	actualPureNumStr3,
		pureNumStrComponents,
		err = nStr02.FmtNumStrPure(
		".",
		true,
		NumRoundType.HalfAwayFromZero(),
		5,
		ePrefix.XCpy(
			"actualNativeNumStr3B<-nStr02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	testName = "Test# 3-A Pure Number String Components Native Number String Test"

	if actualNativeNumStr3A !=
		pureNumStrComponents.NativeNumberStr {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: Actual Native Number String 3-A DOES NOT EQUAL\n"+
			"Pure Number String Components Native Number String!\n"+
			"                   Actual Native Number String 3-A = '%v'\n"+
			"Pure Number String Components Native Number String = '%v'\n",
			ePrefix.String(),
			testName,
			actualNativeNumStr3A,
			pureNumStrComponents.NativeNumberStr)

		return

	}

	testName = "Test# 3-B Pure Number String Test"

	if actualNativeNumStr3A !=
		actualPureNumStr3 {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: Actual Native Number String 3-A DOES NOT EQUAL\n"+
			"Pure Number String !\n"+
			"Actual Native Number String 3-A = '%v'\n"+
			"             Pure Number String = '%v'\n",
			ePrefix.String(),
			testName,
			actualNativeNumStr3A,
			actualPureNumStr3)

		return

	}

	testName = "Test# 4 NumberStrKernel Equality Test"

	if !nStr01.Equal(&nStr02) {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: nStr01 IS NOT EQUAL TO nStr02\n",
			ePrefix.String(),
			testName)

		return
	}

	return
}

func TestNumberStrKernel_FmtNumStrPure_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_FmtNumStrPure_000200()",
		"")

	origIntStr := "1234"
	origFracStr := "5678"

	origNumberStr := origIntStr +
		"." +
		origFracStr

	expectedNumberStr := "1234.5678"

	var err error
	var baseValueNStr NumberStrKernel

	baseValueNStr,
		_,
		err = new(NumberStrKernel).
		NewParsePureNumberStr(
			origNumberStr,
			".",
			true,
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy(
				"baseValueNStr<-origNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var actualFmtNumberStr string

	actualFmtNumberStr,
		_,
		err = baseValueNStr.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"actualFmtNumberStr<-baseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != actualFmtNumberStr {

		t.Errorf("%v\n"+
			"Test#1\n"+
			"Error: actualfmtNumberStr NOT EQUAL TO expectedNumberStr\n"+
			" actualfmtNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualFmtNumberStr,
			expectedNumberStr)

		return
	}

	actualFmtNumberStr,
		_,
		err = baseValueNStr.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"baseValueNStr"))

	if expectedNumberStr != actualFmtNumberStr {

		t.Errorf("%v\n"+
			"Test#2\n"+
			"baseValueNStr.FmtPureNumberStr()\n"+
			"Error: actualfmtNumberStr NOT EQUAL TO expectedNumberStr\n"+
			" actualfmtNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualFmtNumberStr,
			expectedNumberStr)

		return
	}

	err = baseValueNStr.SetNumberSign(
		NumSignVal.Negative(),
		ePrefix.XCpy(
			"baseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var nStr03 NumberStrKernel

	nStr03,
		err = baseValueNStr.CopyOut(
		ePrefix.XCpy(
			"nStr03<-baseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr03.Round(
		NumRoundType.HalfAwayFromZero(),
		3,
		ePrefix.XCpy(
			"nStr03<-HalfAwayFromZero"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualFmtNumberStr,
		_,
		err = nStr03.FmtNumStrPure(
		".",
		false,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"actualFmtNumberStr<-baseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	expectedNumberStr = "1234.568-"

	if expectedNumberStr != actualFmtNumberStr {

		t.Errorf("%v\n"+
			"Test#2\n"+
			"baseValueNStr.FmtPureNumberStr()\n"+
			"Error: actualfmtNumberStr NOT EQUAL TO expectedNumberStr\n"+
			" actualfmtNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualFmtNumberStr,
			expectedNumberStr)

		return
	}

	return
}

func TestNumberStrKernel_FmtNumStrPure_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_FmtNumStrPure_000300()",
		"")

	origNumberStr := "1234.5678"

	expectedNumberStr := "1234.57"

	var err error
	var baseValueNStr NumberStrKernel

	baseValueNStr,
		_,
		err = new(NumberStrKernel).
		NewParseNativeNumberStr(
			origNumberStr,
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy(
				"baseValueNStr<-origNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var actualFmtNumberStr string

	actualFmtNumberStr,
		_,
		err = baseValueNStr.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"#1 actualFmtNumberStr<-baseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	testName := "Test #1"

	if origNumberStr != actualFmtNumberStr {

		t.Errorf("%v\n"+
			"%v\n"+
			"baseValueNStr.FmtNumStrPure()\n"+
			"Error: actualfmtNumberStr NOT EQUAL TO origNumberStr\n"+
			" actualFmtNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualFmtNumberStr,
			expectedNumberStr)

		return
	}

	actualFmtNumberStr,
		_,
		err = baseValueNStr.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"baseValueNStr #2"))

	testName = "Test #2"

	if origNumberStr != actualFmtNumberStr {

		t.Errorf("%v\n"+
			"%v\n"+
			"baseValueNStr.FmtNumStrPure()\n"+
			"Error: actualfmtNumberStr NOT EQUAL TO origNumberStr\n"+
			"actualfmtNumberStr = '%v'\n"+
			"     origNumberStr = '%v'\n",
			ePrefix.String(),
			testName,
			actualFmtNumberStr,
			expectedNumberStr)

		return
	}

	actualFmtNumberStr,
		_,
		err = baseValueNStr.FmtNumStrNative(
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"baseValueNStr #3"))

	testName = "Test #3 - Native Number String"

	if origNumberStr != actualFmtNumberStr {

		t.Errorf("%v\n"+
			"%v\n"+
			"baseValueNStr.FmtPureNumberStr()\n"+
			"Error: actualfmtNumberStr NOT EQUAL TO origNumberStr\n"+
			"actualfmtNumberStr = '%v'\n"+
			"     origNumberStr = '%v'\n",
			ePrefix.String(),
			testName,
			actualFmtNumberStr,
			expectedNumberStr)

		return
	}

	return
}

func TestNumberStrKernel_FmtNumStrNative_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_FmtSignedSimpleNumber_000100()",
		"")

	origIntStr := "1234"
	origFracStr := "5678"

	origNumberStr := origIntStr +
		"." +
		origFracStr

	expectedNumberStr := "1234.5678"

	var err error
	var baseValueNStr NumberStrKernel

	baseValueNStr,
		_,
		err = new(NumberStrKernel).
		NewParsePureNumberStr(
			origNumberStr,
			".",
			true,
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy(
				"baseValueNStr<-origNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = baseValueNStr.IsValidInstanceError(
		ePrefix.XCpy(
			"baseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var fmtNumberStr string

	fmtNumberStr,
		_,
		err = baseValueNStr.FmtNumStrNative(
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"fmtNumberStr<-baseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != fmtNumberStr {

		t.Errorf("%v\n"+
			"Test#1\n"+
			"Error: actualfmtNumberStr NOT EQUAL TO expectedNumberStr\n"+
			" actualfmtNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			fmtNumberStr,
			expectedNumberStr)

		return
	}

}

func TestNumberStrKernel_FmtNumStrNative_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_FmtNumStrNative_000200()",
		"")

	origIntStr := "1234"
	origFracStr := "5678"

	origNumberStr := origIntStr +
		"." +
		origFracStr

	expectedNumberStr := "1234.5678"

	var err error
	var baseValueNStr NumberStrKernel

	baseValueNStr,
		_,
		err = new(NumberStrKernel).
		NewParsePureNumberStr(
			origNumberStr,
			".",
			true,
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy(
				"baseValueNStr<-origNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	isValid := baseValueNStr.IsValidInstance()

	if !isValid {

		t.Errorf("\n%v\n"+
			"baseValueNStr.IsValidInstance()\n"+
			"returned 'false'.\n",
			ePrefix.String())

		return
	}

	var fmtNumberStr string

	fmtNumberStr,
		_,
		err = baseValueNStr.FmtNumStrNative(
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"fmtNumberStr<-baseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != fmtNumberStr {

		t.Errorf("%v\n"+
			"Test#1\n"+
			"Error: actualfmtNumberStr NOT EQUAL TO expectedNumberStr\n"+
			" actualfmtNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			fmtNumberStr,
			expectedNumberStr)

		return
	}

	return
}

func TestNumberStrKernel_FmtNumStrNative_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_FmtNumStrNative_000300()",
		"")

	origIntStr := "001234"
	origFracStr := "567800"

	origNumberStr := origIntStr +
		"." +
		origFracStr

	var expectedNumberStr, fmtNumberStr, testName string

	var err error
	var nStrKernel01 NumberStrKernel

	nStrKernel01,
		_,
		err = new(NumberStrKernel).
		NewParsePureNumberStr(
			origNumberStr,
			".",
			true,
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy(
				"Test #1-A nStrKernel01<-origNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	isValid := nStrKernel01.IsValidInstance()

	if !isValid {

		t.Errorf("\n%v\n"+
			"nStrKernel01.IsValidInstance()\n"+
			"returned 'false'.\n",
			ePrefix.String())

		return
	}

	intRuneArray := nStrKernel01.GetIntegerRuneArray()

	expectedNumberStr = origIntStr

	testName = fmt.Sprintf("Test #1-A NewParsePureNumberStr()\n"+
		"Test Integer Rune Array\n"+
		"origIntStr  = %v",
		origIntStr)

	if string(intRuneArray) != expectedNumberStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: string(intRuneArray) != expectedNumberStr\n"+
			"string(intRuneArray)  = '%v'\n"+
			"   expectedNumberStr  = '%v'\n",
			ePrefix.String(),
			testName,
			string(intRuneArray),
			expectedNumberStr)

		return

	}

	fracArray := nStrKernel01.GetFractionalRuneArray()

	testName = fmt.Sprintf("Test #1-B NewParsePureNumberStr()\n"+
		"Test Fractional Rune Array\n"+
		"origNumberStr = %v",
		origNumberStr)

	expectedNumberStr = origFracStr

	if string(fracArray) != expectedNumberStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: string(fracArray) != expectedNumberStr\n"+
			"string(fracArray)  = '%v'\n"+
			"expectedNumberStr  = '%v'\n",
			ePrefix.String(),
			testName,
			string(fracArray),
			expectedNumberStr)

		return

	}

	fmtNumberStr,
		_,
		err = nStrKernel01.FmtNumStrNative(
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"Test #2 fmtNumberStr<-nStrKernel01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	expectedNumberStr = "1234.5678"

	testName = fmt.Sprintf("Test #2 FmtNumStrNative()\n"+
		"Test Returned Native Number String\n"+
		"origNumberStr = %v",
		origNumberStr)

	if expectedNumberStr != fmtNumberStr {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: fmtNumberStr NOT EQUAL TO expectedNumberStr\n"+
			" fmtNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			testName,
			fmtNumberStr,
			expectedNumberStr)

		return
	}

	return
}

func TestNumberStrKernel_FmtSignedNumStrBasic_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_SetFromNativeNumberStr_000100()",
		"")

	originalNumberStr := "1 123 456,775"

	//expectedStr  "1,123,456.78"

	var numberStrSearchResults CharSearchNumStrParseResultsDto
	var nStrKernel01 NumberStrKernel
	var err error

	numberStrSearchResults,
		nStrKernel01,
		err = new(NumberStrKernel).NewParseFrenchNumberStr(
		originalNumberStr,
		0,
		-1,
		nil,
		false,
		ePrefix.XCpy(
			"nStrKernel01<-"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if numberStrSearchResults.FoundNumericDigits == false {

		t.Errorf("\n%v\n"+
			"Error: Test #1 - NewParseFrenchNumberStr() Failed!\n"+
			"No Numeric Digits were found in French Number String.\n",
			ePrefix.String())

		return
	}

	err = nStrKernel01.IsValidInstanceError(
		ePrefix.XCpy(
			"Test #1 - nStrKernel01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var actualNumStr, expectedStr string

	actualNumStr,
		_,
		err = nStrKernel01.FmtNumStrNative(
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"Test #1 actualNumStr<-"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	testName := "Test #1 Native Number String\n" +
		"NewParseFrenchNumberStr()"

	expectedStr = "1123456.775"

	if expectedStr != actualNumStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: expectedStr != actualFmtStr\n"+
			" expectedStr = %v\n"+
			"actualNumStr = %v\n",
			ePrefix.String(),
			testName,
			expectedStr,
			actualNumStr)

		return
	}

	//expectedStr  "    1,123,456.78"
	var roundingSpec NumStrRoundingSpec

	roundingSpec,
		err = new(NumStrRoundingSpec).NewRoundingSpec(
		NumRoundType.HalfAwayFromZero(),
		2,
		ePrefix.XCpy(
			"roundingSpec"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumStr,
		err = nStrKernel01.FmtSignedNumStrBasic(
		roundingSpec,
		".",
		",",
		IntGroupingType.Thousands(),
		"(",
		")",
		16,
		TxtJustify.Right(),
		ePrefix.XCpy(
			"Test # 2 actualNumStr<-nStrKernel01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	testName = "Test #2 Native Number String\n" +
		"FmtSignedNumStrBasic()"

	expectedStr = "    1,123,456.78"

	if expectedStr != actualNumStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: expectedStr != actualFmtStr\n"+
			" expectedStr = %v\n"+
			"actualNumStr = %v\n",
			ePrefix.String(),
			testName,
			expectedStr,
			actualNumStr)

		return
	}

	return
}

func TestNumberStrKernel_IsZeroValue_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"\nTestNumberStrKernel_IsZeroValue_000100()",
		"")

	var err error
	var numStrKernel01 NumberStrKernel

	origNum := uint64(12345)

	origIntStr := fmt.Sprintf("%v",
		origNum)

	origFracStr := ""

	numStrKernel01,
		err = new(NumberStrKernel).NewFromNumericValue(
		origNum,
		NumRoundType.NoRounding(),
		0,
		&ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualIntStr := numStrKernel01.GetIntegerString()

	if actualIntStr != origIntStr {
		t.Errorf("%v\n"+
			"Test#1\n"+
			"Error: actualIntStr != origIntStr\n"+
			"         actualIntStr = '%v'\n"+
			"         origIntStr   = '%v'\n"+
			"Original uint64 Value = '%v'\n",
			ePrefix.String(),
			actualIntStr,
			origIntStr,
			origNum)

		return
	}

	actualFracStr := numStrKernel01.GetFractionalString()

	if actualFracStr != origFracStr {

		t.Errorf("%v\n"+
			"Test#2\n"+
			"Error: actualFracStr != origFracStr\n"+
			"actualFracStr = '%v'\n"+
			"origFracStr   = '%v'\n",
			ePrefix.String(),
			actualFracStr,
			origFracStr)

		return
	}

	isZeroValue := numStrKernel01.IsZeroValue()

	if isZeroValue == true {

		t.Errorf("%v\n"+
			"Test#3\n"+
			"Error: Expected numStrKernel01.IsZeroValue()\n"+
			"would be equal to 'false'. Instead the return value\n"+
			"was 'true'!\n",
			ePrefix.String())

		return
	}

	var numStrKernel02 NumberStrKernel

	uintNum := uint(0)

	numStrKernel02,
		err = new(NumberStrKernel).NewFromNumericValue(
		uintNum,
		NumRoundType.NoRounding(),
		0,
		&ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	expectedIntStr := "0"

	actualIntStr = numStrKernel02.GetIntegerString()

	if actualIntStr != expectedIntStr {
		t.Errorf("%v\n"+
			"Test#4\n"+
			"Error: actualIntStr != origIntStr\n"+
			"         actualIntStr = '%v'\n"+
			"     expectedIntStr   = '%v'\n"+
			"  Original uint Value = '%v'\n",
			ePrefix.String(),
			actualIntStr,
			origIntStr,
			uintNum)

		return
	}

	actualFracStr = numStrKernel02.GetFractionalString()

	if actualFracStr != origFracStr {

		t.Errorf("%v\n"+
			"Test#5\n"+
			"Error: actualFracStr != origFracStr\n"+
			"actualFracStr = '%v'\n"+
			"origFracStr   = '%v'\n",
			ePrefix.String(),
			actualFracStr,
			origFracStr)

		return
	}

	isZeroValue = numStrKernel02.IsZeroValue()

	if isZeroValue == false {

		t.Errorf("%v\n"+
			"Test#6\n"+
			"Error: Expected numStrKernel02.IsZeroValue()\n"+
			"would be equal to 'true'. Instead the return value\n"+
			"was 'false'!\n",
			ePrefix.String())

		return
	}

	var numStrKernel03 NumberStrKernel

	uint32Num := uint32(9879)

	expectedIntStr = fmt.Sprintf("%v",
		uint32Num)

	numStrKernel03,
		err = new(NumberStrKernel).NewFromNumericValue(
		uint32Num,
		NumRoundType.NoRounding(),
		0,
		&ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var intRuneArrayDto RuneArrayDto

	intRuneArrayDto,
		err = numStrKernel03.GetIntegerDigits(
		ePrefix.XCpy(
			"numStrKernel03"))
	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualIntStr = intRuneArrayDto.GetCharacterString()

	if actualIntStr != expectedIntStr {
		t.Errorf("%v\n"+
			"Test#7\n"+
			"Error: actualIntStr != origIntStr\n"+
			"         actualIntStr = '%v'\n"+
			"     expectedIntStr   = '%v'\n"+
			"  Original uint Value = '%v'\n",
			ePrefix.String(),
			actualIntStr,
			expectedIntStr,
			uint32Num)

		return
	}

	actualFracStr = numStrKernel03.GetFractionalString()

	if actualFracStr != origFracStr {

		t.Errorf("%v\n"+
			"Test#8\n"+
			"Error: actualFracStr != origFracStr\n"+
			"actualFracStr = '%v'\n"+
			"origFracStr   = '%v'\n",
			ePrefix.String(),
			actualFracStr,
			origFracStr)

		return
	}

	isZeroValue = numStrKernel03.IsZeroValue()

	if isZeroValue == true {

		t.Errorf("%v\n"+
			"Test#9\n"+
			"Error: Expected numStrKernel03.IsZeroValue()\n"+
			"would be equal to 'false'. Instead the return value\n"+
			"was 'false'!\n",
			ePrefix.String())

		return
	}

}

func TestNumberStrKernel_GetIntFracDigitsStr_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_FmtCharReplacementStr_000100()",
		"")

	intDigits := "123456789"
	fracDigits := "1234567890"

	var err error
	var nStrKernel01 NumberStrKernel

	nStrKernel01,
		err = new(NumberStrKernel).
		NewFromStringDigits(
			intDigits,
			fracDigits,
			NumSignVal.Positive(),
			ePrefix.XCpy(
				"Test #1-A: nStrKernel01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	testName := fmt.Sprintf("Test #1-A Positive Values - GetIntFracDigitsStr()\n"+
		" intDigits = %v\n"+
		"fracDigits = %v",
		intDigits,
		fracDigits)

	var formattedNumStr string

	formattedNumStr,
		err = nStrKernel01.GetIntFracDigitsStr(
		true,
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"Test #1-A nStrKernel01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	expectedStr := intDigits + fracDigits

	if formattedNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: formattedNumStr != expectedStr\n"+
			"formattedNumStr  = '%v'\n"+
			"   expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			formattedNumStr,
			expectedStr)

		return

	}

	testName = fmt.Sprintf("Test #2-A Negative Values - GetIntFracDigitsStr()\n"+
		" intDigits = %v\n"+
		"fracDigits = %v",
		intDigits,
		fracDigits)

	var nStrKernel02 NumberStrKernel

	nStrKernel02,
		err = new(NumberStrKernel).
		NewFromStringDigits(
			intDigits,
			fracDigits,
			NumSignVal.Negative(),
			ePrefix.XCpy(
				"Test #2-A: nStrKernel02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	formattedNumStr,
		err = nStrKernel02.GetIntFracDigitsStr(
		true,
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"Test #2-A nStrKernel02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	expectedStr = "-" + intDigits + fracDigits

	if formattedNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: formattedNumStr != expectedStr\n"+
			"formattedNumStr  = '%v'\n"+
			"   expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			formattedNumStr,
			expectedStr)

		return

	}

	testName = fmt.Sprintf("Test #3-A Trailing Negative Values - GetIntFracDigitsStr()\n"+
		" intDigits = %v\n"+
		"fracDigits = %v",
		intDigits,
		fracDigits)

	var nStrKernel03 NumberStrKernel

	nStrKernel03,
		err = new(NumberStrKernel).
		NewFromStringDigits(
			intDigits,
			fracDigits,
			NumSignVal.Negative(),
			ePrefix.XCpy(
				"Test #3-A: nStrKernel03"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	formattedNumStr,
		err = nStrKernel03.GetIntFracDigitsStr(
		true,
		false,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"Test #3-A nStrKernel03"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	expectedStr = intDigits + fracDigits + "-"

	if formattedNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: formattedNumStr != expectedStr\n"+
			"formattedNumStr  = '%v'\n"+
			"   expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			formattedNumStr,
			expectedStr)

		return

	}

	intDigits = "0"
	fracDigits = ""

	testName = fmt.Sprintf("Test #4-A Zero Value - GetIntFracDigitsStr()\n"+
		" intDigits = %v\n"+
		"fracDigits = %v",
		intDigits,
		fracDigits)

	var nStrKernel04 NumberStrKernel

	nStrKernel04,
		err = new(NumberStrKernel).
		NewFromStringDigits(
			intDigits,
			fracDigits,
			NumSignVal.Positive(),
			ePrefix.XCpy(
				"Test #4-A: nStrKernel04"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	testName = fmt.Sprintf("Test #4-A NumericSignValueType with Zero Value\n"+
		" intDigits = %v\n"+
		"fracDigits = %v",
		intDigits,
		fracDigits)

	var numValue NumericSignValueType

	numValue,
		err = nStrKernel04.GetNumberSign(
		ePrefix.XCpy("Test #4A nStrKernel04"))

	if numValue != NumSignVal.Zero() {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: Number Sign IS NOT EQUAL to Zero!\n"+
			"Number Sign  = '%v'\n",
			ePrefix.String(),
			testName,
			numValue.String())

		return
	}

	testName = fmt.Sprintf("Test #4-B Zero Value - GetIntFracDigitsStr()\n"+
		" intDigits = %v\n"+
		"fracDigits = %v",
		intDigits,
		fracDigits)

	formattedNumStr,
		err = nStrKernel04.GetIntFracDigitsStr(
		true,
		false,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"Test #4-B nStrKernel04"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	expectedStr = intDigits

	if formattedNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: formattedNumStr != expectedStr\n"+
			"formattedNumStr  = '%v'\n"+
			"   expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			formattedNumStr,
			expectedStr)

		return

	}

	intDigits = "12345"
	fracDigits = "5678"

	testName = fmt.Sprintf("Test #5-A Rounded Value - GetIntFracDigitsStr()\n"+
		" intDigits = %v\n"+
		"fracDigits = %v",
		intDigits,
		fracDigits)

	var nStrKernel05 NumberStrKernel

	nStrKernel05,
		err = new(NumberStrKernel).
		NewFromStringDigits(
			intDigits,
			fracDigits,
			NumSignVal.Positive(),
			ePrefix.XCpy(
				"Test #5-A: nStrKernel05"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	formattedNumStr,
		err = nStrKernel05.GetIntFracDigitsStr(
		true,
		true,
		NumRoundType.HalfAwayFromZero(),
		2,
		ePrefix.XCpy(
			"Test #5-A nStrKernel05"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	/*
		intDigits = "12345"
			fracDigits = "5678"

	*/

	expectedStr = "1234557"

	if formattedNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: formattedNumStr != expectedStr\n"+
			"formattedNumStr  = '%v'\n"+
			"   expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			formattedNumStr,
			expectedStr)

		return

	}

	return
}

func TestNumberStrKernel_GetBigRatNum_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_GetBigRatNum_000100()",
		"")

	var origIntDigits, origFracDigits string

	origIntDigits = "123"
	origFracDigits = "456"

	bigRatToFracDigits := 5

	expectedBigRatNumStr := "123.45600"

	origNumberSign := NumSignVal.Positive()

	var origNumStr string

	origNumStr += origIntDigits

	if len(origFracDigits) > 0 {

		origNumStr += "."
		origNumStr += origFracDigits
	}

	numStrKernelBase,
		err := new(NumberStrKernel).
		NewFromStringDigits(
			origIntDigits,
			origFracDigits,
			origNumberSign,
			ePrefix.XCpy(
				origNumStr))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var pureNumStr string

	pureNumStr,
		_,
		err = numStrKernelBase.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"numStrKernelBase"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if pureNumStr != origNumStr {

		t.Errorf("%v\n"+
			"Test#1\n"+
			"Error: pureNumStr != origNumStr\n"+
			"pureNumStr = '%v'\n"+
			"origNumStr = '%v'\n",
			ePrefix.String(),
			pureNumStr,
			origNumStr)

		return

	}

	bigRatNum := new(big.Rat)

	_,
		err = numStrKernelBase.FmtNumericValue(
		bigRatNum,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"bigRatNum<-"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var actualBigRatNumStr string

	actualBigRatNumStr =
		bigRatNum.FloatString(bigRatToFracDigits)

	if actualBigRatNumStr != expectedBigRatNumStr {

		t.Errorf("%v\n"+
			"Test#2\n"+
			"Error: actualBigRatNumStr != expectedBigRatNumStr\n"+
			"actualBigRatNumStr   = '%v'\n"+
			"expectedBigRatNumStr = '%v'\n",
			ePrefix.String(),
			actualBigRatNumStr,
			expectedBigRatNumStr)

		return

	}

	return
}

func TestNumberStrKernel_GetScientificNotation_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_GetScientificNotation_000100()",
		"")

	origIntStr := "1234567"

	origFracStr := "890"

	expectedSciNotStr := "1.23456789 x 10^6"

	compositeNumStr := origIntStr + "." + origFracStr

	var err error
	var numStrKernel01 NumberStrKernel

	numStrKernel01,
		_,
		err = new(NumberStrKernel).NewParsePureNumberStr(
		compositeNumStr,
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualIntStr := numStrKernel01.GetIntegerString()

	if actualIntStr != origIntStr {

		t.Errorf("%v\n"+
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

		t.Errorf("%v\n"+
			"Test#2\n"+
			"Error: actualFracStr != origFracStr\n"+
			"actualFracStr = '%v'\n"+
			"origFracStr   = '%v'\n",
			ePrefix.String(),
			actualFracStr,
			origFracStr)

		return
	}

	var sciNot01 SciNotationKernel

	sciNot01,
		err = numStrKernel01.GetScientificNotation(
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"sciNot01<-"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var actualSciNotStr string

	actualSciNotStr = sciNot01.GetNumStrExponentFmt()

	if actualSciNotStr != expectedSciNotStr {

		t.Errorf("%v\n"+
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

	return
}

func TestNumberStrKernel_GetScientificNotation_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_GetScientificNotation_000100()",
		"")

	origIntStr := "1234567"

	origFracStr := "8901"

	expectedSciNotStr := "1.2345678901E+6"

	compositeNumStr := origIntStr + "." + origFracStr

	var err error
	var numStrKernel01 NumberStrKernel
	var numStrSearchResults CharSearchNumStrParseResultsDto

	numStrSearchResults,
		numStrKernel01,
		err = new(NumberStrKernel).NewParseUSNumberStr(
		compositeNumStr,
		0,
		-1,
		nil,
		false,
		ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if !numStrSearchResults.FoundNumericDigits {

		t.Errorf("%v\n" +
			"Test#1\n" +
			"Error: Failed to locate numeric\n" +
			"digits in number string!\n" +
			ePrefix.String())

		return
	}

	actualIntStr := numStrKernel01.GetIntegerString()

	if actualIntStr != origIntStr {

		t.Errorf("%v\n"+
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

		t.Errorf("%v\n"+
			"Test#3\n"+
			"Error: actualFracStr != origFracStr\n"+
			"actualFracStr = '%v'\n"+
			"origFracStr   = '%v'\n",
			ePrefix.String(),
			actualFracStr,
			origFracStr)

		return
	}

	var sciNot01 SciNotationKernel

	sciNot01,
		err = numStrKernel01.GetScientificNotation(
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"sciNot01<-"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var actualSciNotStr string

	actualSciNotStr = sciNot01.GetENotationFmt(
		"",
		"")

	if actualSciNotStr != expectedSciNotStr {

		t.Errorf("%v\n"+
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

	return
}

func TestNumberStrKernel_NewFromBigRat_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_NewFromBigRat_000100()",
		"")

	newRat := big.NewRat(1, 3)

	expectedNumberStr := "0.3333333333"

	var numStrKernel01 NumberStrKernel

	var err error

	numStrKernel01,
		err = new(NumberStrKernel).
		NewFromNumericValue(
			newRat,
			NumRoundType.HalfAwayFromZero(),
			10,
			ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var actualNumberStr string

	actualNumberStr,
		_,
		err = numStrKernel01.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy("numStrKernel01"))

	if actualNumberStr != expectedNumberStr {

		t.Errorf("\n%v\n"+
			"Test#1\n"+
			"Error: actualNumberStr != expectedNumberStr\n"+
			"actualNumberStr   = '%v'\n"+
			"expectedNumberStr = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return

	}

	return
}

func TestNumberStrKernel_NewFromBigRat_000125(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_NewFromBigRat_000125()",
		"")

	newRat := big.NewRat(1, 3)

	expectedNumberStr := "0.3333333333"

	var numStrKernel01 NumberStrKernel

	var err error

	numStrKernel01,
		err = new(NumberStrKernel).
		NewFromBigRat(
			newRat,
			10,
			ePrefix.XCpy(
				"numStrKernel01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var actualNumberStr string

	actualNumberStr,
		_,
		err = numStrKernel01.FmtNumStrNative(
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"actualNumberStr<-"+
				"numStrKernel01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumberStr != expectedNumberStr {

		t.Errorf("\n%v\n"+
			"Test#1\n"+
			"Error: actualNumberStr != expectedNumberStr\n"+
			"actualNumberStr   = '%v'\n"+
			"expectedNumberStr = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return

	}

	newRat = big.NewRat(25, 100)

	expectedNumberStr = "0.25"

	var numStrKernel02 NumberStrKernel

	err = numStrKernel02.SetFromBigRat(
		newRat,
		2,
		ePrefix.XCpy(
			"numStrKernel02"))

	actualNumberStr,
		_,
		err = numStrKernel02.FmtNumStrNative(
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"actualNumberStr<-numStrKernel02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumberStr != expectedNumberStr {

		t.Errorf("\n%v\n"+
			"Test#2\n"+
			"Error: actualNumberStr != expectedNumberStr\n"+
			"actualNumberStr   = '%v'\n"+
			"expectedNumberStr = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return

	}

	return
}

func TestNumberStrKernel_NewFromBigRat_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_NewFromBigRat_000200()",
		"")

	newRat := big.NewRat(1, 3)

	expectedNumberStr := "0.3333333333"

	var numStrKernel01 NumberStrKernel

	var err error

	numStrKernel01,
		err = new(NumberStrKernel).
		NewFromBigRat(
			newRat,
			10,
			ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var actualNumberStr string

	actualNumberStr,
		_,
		err = numStrKernel01.FmtNumStrNative(
		NumRoundType.NoRounding(),
		0,
		ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumberStr != expectedNumberStr {

		t.Errorf("\n%v\n"+
			"Test#1\n"+
			"Error: actualNumberStr != expectedNumberStr\n"+
			"actualNumberStr   = '%v'\n"+
			"expectedNumberStr = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return

	}

	return
}

func TestNumberStrKernel_NewParseGermanNumberStr_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_NewParseGermanNumberStr_000100()",
		"")

	var expectedNumberStr, fmtNumberStr, testName,
		origNumberStr string

	origNumberStr = "1.234.567,89 €"

	var err error
	var nStrKernel01 NumberStrKernel
	var numberStrSearchResults CharSearchNumStrParseResultsDto

	numberStrSearchResults,
		nStrKernel01,
		err = new(NumberStrKernel).NewParseGermanNumberStr(
		origNumberStr,
		0,
		-1,
		nil,
		false,
		ePrefix.XCpy(
			"Test 1-A nStrKernel01<-origNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	testName = fmt.Sprintf("Test #1-A NewParseGermanNumberStr()\n"+
		"Test Number String Search Results FoundNumericDigits\n"+
		"origNumberStr  = %v",
		origNumberStr)

	if numberStrSearchResults.FoundNumericDigits == false {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: numberStrSearchResults.FoundNumericDigits == false\n"+
			"Failed To Find Numeric Digits while parsing German Number String.\n",
			ePrefix.String(),
			testName)

		return
	}

	testName = fmt.Sprintf("Test #1-B NewParseGermanNumberStr()\n"+
		"Test Number String Search Results FoundIntegerDigits\n"+
		"origNumberStr  = %v",
		origNumberStr)

	if numberStrSearchResults.FoundIntegerDigits == false {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: numberStrSearchResults.FoundIntegerDigits == false\n"+
			"Failed To Find Integer Digits while parsing German Number String.\n",
			ePrefix.String(),
			testName)

		return
	}

	testName = fmt.Sprintf("Test #1-C NewParseGermanNumberStr()\n"+
		"Test Number String Search Results FoundDecimalDigits\n"+
		"origNumberStr  = %v",
		origNumberStr)

	if numberStrSearchResults.FoundDecimalDigits == false {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: numberStrSearchResults.FoundDecimalDigits == false\n"+
			"Failed To Find Decimal (Fractional) Digits while parsing German Number String.\n",
			ePrefix.String(),
			testName)

		return
	}

	testName = fmt.Sprintf("Test #2 NewParseGermanNumberStr()\n"+
		"nStrKernel01.FmtNumStrNative()\n"+
		"origNumberStr  = %v",
		origNumberStr)

	fmtNumberStr,
		_,
		err = nStrKernel01.FmtNumStrNative(
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"Test #2 fmtNumberStr<-nStrKernel01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	expectedNumberStr = "1234567.89"

	if expectedNumberStr != fmtNumberStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: fmtNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"       fmtNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			testName,
			fmtNumberStr,
			expectedNumberStr)

		return
	}

	var nStrKernel02 NumberStrKernel

	origNumberStr = "1.234.567,89- €"

	_,
		nStrKernel02,
		err = new(NumberStrKernel).NewParseGermanNumberStr(
		origNumberStr,
		0,
		-1,
		nil,
		false,
		ePrefix.XCpy(
			"Test #3 nStrKernel01<-origNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	testName = fmt.Sprintf("Test #3 NewParseGermanNumberStr()\n"+
		"nStrKernel02.FmtNumStrNative()\n"+
		"origNumberStr  = %v",
		origNumberStr)

	fmtNumberStr,
		_,
		err = nStrKernel02.FmtNumStrNative(
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"Test #3 fmtNumberStr<-nStrKernel02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	expectedNumberStr = "-1234567.89"

	if expectedNumberStr != fmtNumberStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: fmtNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"       fmtNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			testName,
			fmtNumberStr,
			expectedNumberStr)

		return
	}

	testName = fmt.Sprintf("Test #4 FmtSignedNumStrDefaultsUSMinus()\n"+
		"origNumberStr  = %v",
		origNumberStr)

	numberFieldSpec := NumStrNumberFieldSpec{
		fieldLength:        17,
		fieldJustification: TxtJustify.Center(),
	}

	roundingSpec := NumStrRoundingSpec{
		roundingType:            NumRoundType.NoRounding(),
		roundToFractionalDigits: 0,
	}

	fmtNumberStr,
		err = nStrKernel02.FmtSignedNumStrDefaultsUSMinus(
		roundingSpec,
		numberFieldSpec,
		ePrefix.XCpy(
			"Test #4 fmtNumberStr<-nStrKernel02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	expectedNumberStr = "  -1,234,567.89  "

	if expectedNumberStr != fmtNumberStr {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: fmtNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"       fmtNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			testName,
			fmtNumberStr,
			expectedNumberStr)

		return
	}

	testName = fmt.Sprintf("Test #5 FmtCurrencyDefaultsUSMinus()\n"+
		"origNumberStr  = %v",
		origNumberStr)

	err = numberFieldSpec.SetFieldJustification(
		TxtJustify.Right(),
		ePrefix.XCpy(
			"Test #5 numberFieldSpec"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	fmtNumberStr,
		err = nStrKernel02.FmtCurrencyDefaultsUSMinus(
		roundingSpec,
		numberFieldSpec,
		ePrefix.XCpy(
			"Test #5 fmtNumberStr<-nStrKernel02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	expectedNumberStr = "  $ -1,234,567.89"

	if expectedNumberStr != fmtNumberStr {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: fmtNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"       fmtNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			testName,
			fmtNumberStr,
			expectedNumberStr)

		return
	}

}

func TestNumberStrKernel_NewParsePureNumberStr_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_NewParsePureNumberStr_000100()",
		"")

	origIntStr := "001234"
	origFracStr := "567800"

	origNumberStr := origIntStr +
		"." +
		origFracStr

	var err error
	var baseValueNStr NumberStrKernel
	var expectedNumberStr, actualFmtNumberStr,
		testName string

	baseValueNStr,
		_,
		err = new(NumberStrKernel).
		NewParsePureNumberStr(
			origNumberStr,
			".",
			true,
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy(
				"baseValueNStr<-origNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualFmtNumberStr,
		_,
		err = baseValueNStr.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"actualFmtNumberStr<-baseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	testName = fmt.Sprintf("Test #1-A FmtNumStrPure()\n"+
		"Test Leading & Trailing Zeros\n"+
		"origNumberStr  = %v",
		origNumberStr)

	expectedNumberStr = "001234.567800"

	if expectedNumberStr != actualFmtNumberStr {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: actualfmtNumberStr NOT EQUAL TO expectedNumberStr\n"+
			" actualfmtNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualFmtNumberStr,
			expectedNumberStr)

		return
	}

}
