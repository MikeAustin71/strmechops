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
		err = numStrKernelBase.GetNumericValue(
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

func TestNumberStrKernel_FmtSignedSimpleNumber_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_FmtSignedSimpleNumber_000100()",
		"")

	origIntStr := "1234"
	origFracStr := "5678"

	origNumberStr := origIntStr +
		"." +
		origFracStr

	expectedNumberStr := "1,234.5678"

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

	roundingSpec := NumStrRoundingSpec{}

	roundingSpec,
		err = new(NumStrRoundingSpec).NewRoundingSpec(
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"roundingSpec<-"+
				"NumRoundType.NoRounding()"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var fmtNumberStr string

	fmtNumberStr,
		err = baseValueNStr.FmtSignedNumStrSimple(
		".",
		",",
		true,
		-1,
		TxtJustify.Right(),
		roundingSpec,
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

func TestNumberStrKernel_FmtSignedSimpleNumber_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_FmtSignedSimpleNumber_000200()",
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

	roundingSpec := NumStrRoundingSpec{}

	roundingSpec,
		err = new(NumStrRoundingSpec).NewRoundingSpec(
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"roundingSpec<-"+
				"NumRoundType.NoRounding()"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var fmtNumberStr string

	fmtNumberStr,
		err = baseValueNStr.FmtSignedNumStrSimple(
		".",
		"",
		true,
		-1,
		TxtJustify.Right(),
		roundingSpec,
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

func TestNumberStrKernel_FmtSignedPureNumberStr_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_FmtSignedPureNumberStr_000100()",
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

func TestNumberStrKernel_FmtSignedPureNumberStr_000120(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_FmtSignedPureNumberStr_000120()",
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

func TestNumberStrKernel_RoundNoRounding_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_RoundNoRounding_000100",
		"")

	inputNumberStr := "7.5"
	expectedNumberStr := "7.5"

	var err error
	var baseVal, nStr01, nStr02 NumberStrKernel

	baseVal,
		_,
		err = new(NumberStrKernel).
		NewParsePureNumberStr(
			inputNumberStr,
			".",
			true,
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy(
				"baseVal<-inputNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = baseVal.IsValidInstanceError(
		ePrefix.XCpy(
			"baseVal"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if baseVal.IsZeroValue() {

		t.Errorf("%v\n"+
			"Setup Test#1 - baseVal.IsZeroValue()\n"+
			"Error: IsZeroValue returned 'true'.\n"+
			"However baseVal = '%v'\n",
			ePrefix.String(),
			expectedNumberStr)

		return
	}

	if !baseVal.IsFloatingPointValue() {

		t.Errorf("%v\n"+
			"Setup Test#2 - baseVal.IsFloatingPointValue()\n"+
			"Error: IsZeroValue returned 'false'.\n"+
			"However baseVal = '%v'\n",
			ePrefix.String(),
			expectedNumberStr)

		return
	}

	var intNumSign = 0

	intNumSign,
		err = baseVal.GetNumberSignAsInt(
		ePrefix.XCpy(
			"baseVal"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if intNumSign != 1 {

		t.Errorf("%v\n"+
			"Setup Test#3 - baseVal.GetNumberSignAsInt()\n"+
			"Error: Integer Number Sign should be '1'.\n"+
			"However intNumSign = '%v'\n",
			ePrefix.String(),
			intNumSign)

		return
	}

	strBuilder := &strings.Builder{}

	err = baseVal.GetParameterTextListing(
		strBuilder,
		false,
		ePrefix.XCpy(
			""))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if strBuilder.Len() < 50 {

		t.Errorf("%v\n"+
			"Setup Test#4 - baseVal.GetParameterTextListing()\n"+
			"Error: Returned string length should be greater than 50.\n"+
			"However string length = '%v'\n",
			ePrefix.String(),
			strBuilder.Len())

		return
	}

	var numOfFracDigits, numOfIntDigits, numOfDigits int

	numOfFracDigits = baseVal.GetNumberOfFractionalDigits()

	if numOfFracDigits != 1 {

		t.Errorf("%v\n"+
			"Setup Test#5 - baseVal.GetNumberOfFractionalDigits()\n"+
			"Error: Returned number of fractional digits should be '1'.\n"+
			"However actual number of fractional digits = '%v'\n",
			ePrefix.String(),
			numOfFracDigits)

		return

	}

	numOfIntDigits = baseVal.GetNumberOfIntegerDigits()

	if numOfIntDigits != 1 {

		t.Errorf("%v\n"+
			"Setup Test#6 - baseVal.GetNumberOfIntegerDigits()\n"+
			"Error: Returned number of integer digits should be '1'.\n"+
			"However actual number of integer digits = '%v'\n",
			ePrefix.String(),
			numOfIntDigits)

		return

	}

	numOfDigits = baseVal.GetNumberOfNumericDigits()

	if numOfDigits != 2 {

		t.Errorf("%v\n"+
			"Setup Test#6 - baseVal.GetNumberOfIntegerDigits()\n"+
			"Error: Returned total number of numeric digits should be '2'.\n"+
			"However actual total number of numeric digits = '%v'\n",
			ePrefix.String(),
			numOfDigits)

		return

	}

	err = nStr01.CopyIn(
		&baseVal,
		ePrefix.XCpy(
			"nStr01<-baseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr02.CopyIn(
		&baseVal,
		ePrefix.XCpy(
			"nStr02<-baseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr02.SetNumberSign(
		NumSignVal.Negative(),
		ePrefix.XCpy(
			"nStr02-NumSignVal.Negative()"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr01.Round(
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr01-NoRounding"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var actualNumberStr string
	var pureNumStrComponents PureNumberStrComponents

	actualNumberStr,
		pureNumStrComponents,
		err = nStr01.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#1A - Positive Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	if pureNumStrComponents.NativeNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#1B - Positive Test\n"+
			"pureNumStrComponents\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"pureNumStrComponents.NativeNumStr = '%v'\n"+
			"              expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			pureNumStrComponents.NativeNumberStr,
			expectedNumberStr)

		return
	}

	err = nStr02.Round(
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr02-NoRounding"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumberStr,
		_,
		err = nStr02.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	expectedNumberStr = "-7.5"

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#2 - Negative Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	return
}

func TestNumberStrKernel_RoundHalfUpWithNegNums_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_RoundHalfUpWithNegNums_000200",
		"")

	inputNumberStr := "7.5"
	expectedNumberStr := "8"

	var err error
	var nStr01, nStr02 NumberStrKernel

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

	err = nStr02.CopyIn(
		&nStr01,
		ePrefix.XCpy(
			"nStr02<-baseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr02.SetNumberSign(
		NumSignVal.Negative(),
		ePrefix.XCpy(
			"nStr02-NumSignVal.Negative()"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr01.Round(
		NumRoundType.HalfUpWithNegNums(),
		0,
		ePrefix.XCpy(
			"nStr01-HalfUpWithNegNums"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var actualNumberStr string

	actualNumberStr,
		_,
		err = nStr01.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#1 - Positive Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	err = nStr02.Round(
		NumRoundType.HalfUpWithNegNums(),
		0,
		ePrefix.XCpy(
			"nStr02-HalfUpWithNegNums"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumberStr,
		_,
		err = nStr02.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	expectedNumberStr = "-7"

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#2 - Negative Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	return
}

func TestNumberStrKernel_RoundHalfDownWithNegNums_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_RoundHalfDownWithNegNums_000300()",
		"")

	inputNumberStr := "7.5"
	expectedNumberStr := "7"

	var err error
	var baseVal, nStr01, nStr02 NumberStrKernel

	baseVal,
		_,
		err = new(NumberStrKernel).
		NewParsePureNumberStr(
			inputNumberStr,
			".",
			true,
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy(
				"baseVal<-inputNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = baseVal.IsValidInstanceError(
		ePrefix.XCpy(
			"baseVal"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr01.CopyIn(
		&baseVal,
		ePrefix.XCpy(
			"nStr01<-baseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr02.CopyIn(
		&baseVal,
		ePrefix.XCpy(
			"nStr02<-baseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr02.SetNumberSign(
		NumSignVal.Negative(),
		ePrefix.XCpy(
			"nStr02-NumSignVal.Negative()"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr01.Round(
		NumRoundType.HalfDownWithNegNums(),
		0,
		ePrefix.XCpy(
			"nStr01-HalfDownWithNegNums"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var actualNumberStr string

	actualNumberStr,
		_,
		err = nStr01.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#1 - Positive Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	err = nStr02.Round(
		NumRoundType.HalfDownWithNegNums(),
		0,
		ePrefix.XCpy(
			"nStr02-HalfDownWithNegNums"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumberStr,
		_,
		err = nStr02.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	expectedNumberStr = "-8"

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#2 - Negative Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	return
}

func TestNumberStrKernel_RoundHalfAwayFromZero_000400(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_RoundHalfAwayFromZero_000400",
		"")

	inputNumberStr := "7.5"
	expectedNumberStr := "8"

	var err error
	var baseVal, nStr01, nStr02 NumberStrKernel

	baseVal,
		_,
		err = new(NumberStrKernel).
		NewParsePureNumberStr(
			inputNumberStr,
			".",
			true,
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy(
				"baseVal<-inputNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = baseVal.IsValidInstanceError(
		ePrefix.XCpy(
			"baseVal"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr01.CopyIn(
		&baseVal,
		ePrefix.XCpy(
			"nStr01<-baseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr02.CopyIn(
		&baseVal,
		ePrefix.XCpy(
			"nStr02<-baseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr02.SetNumberSign(
		NumSignVal.Negative(),
		ePrefix.XCpy(
			"nStr02-NumSignVal.Negative()"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr01.Round(
		NumRoundType.HalfAwayFromZero(),
		0,
		ePrefix.XCpy(
			"nStr01-HalfAwayFromZero"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var actualNumberStr string

	actualNumberStr,
		_,
		err = nStr01.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#1 - Positive Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	err = nStr02.Round(
		NumRoundType.HalfAwayFromZero(),
		0,
		ePrefix.XCpy(
			"nStr02-HalfAwayFromZero"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumberStr,
		_,
		err = nStr02.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	expectedNumberStr = "-8"

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#2 - Negative Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	return
}

func TestNumberStrKernel_RoundHalfTowardsZero_000500(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_RoundHalfTowardsZero_000500",
		"")

	inputNumberStr := "7.5"
	expectedNumberStr := "7"

	var err error
	var baseVal, nStr01, nStr02,
		nStr03 NumberStrKernel

	baseVal,
		_,
		err = new(NumberStrKernel).
		NewParsePureNumberStr(
			inputNumberStr,
			".",
			true,
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy(
				"baseVal<-inputNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = baseVal.IsValidInstanceError(
		ePrefix.XCpy(
			"baseVal"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr01.CopyIn(
		&baseVal,
		ePrefix.XCpy(
			"nStr01<-baseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr02.CopyIn(
		&baseVal,
		ePrefix.XCpy(
			"nStr02<-baseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr02.SetNumberSign(
		NumSignVal.Negative(),
		ePrefix.XCpy(
			"nStr02-NumSignVal.Negative()"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr01.Round(
		NumRoundType.HalfTowardsZero(),
		0,
		ePrefix.XCpy(
			"nStr01-HalfTowardsZero"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var actualNumberStr string

	actualNumberStr,
		_,
		err = nStr01.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#1 - Positive Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	err = nStr02.Round(
		NumRoundType.HalfTowardsZero(),
		0,
		ePrefix.XCpy(
			"nStr02-HalfTowardsZero"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumberStr,
		_,
		err = nStr02.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr02"))

	expectedNumberStr = "-7"

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#2 - Negative Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	expectedNumberStr = "-8"
	inputNumberStr = "-7.6"

	nStr03,
		_,
		err = new(NumberStrKernel).
		NewParsePureNumberStr(
			inputNumberStr,
			".",
			true,
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy(
				"nStr03<-inputNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr03.Round(
		NumRoundType.HalfTowardsZero(),
		0,
		ePrefix.XCpy(
			"nStr03-HalfTowardsZero"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumberStr,
		_,
		err = nStr03.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr03"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#3 - Negative Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	return
}

func TestNumberStrKernel_RoundHalfToEven_000600(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_RoundHalfToEven_000600",
		"")

	inputNumberStr := "7.5"
	expectedNumberStr := "8"

	var err error
	var baseVal, nStr01, nStr02,
		nStr03, nStr04 NumberStrKernel

	baseVal,
		_,
		err = new(NumberStrKernel).
		NewParsePureNumberStr(
			inputNumberStr,
			".",
			true,
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy(
				"baseVal<-inputNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = baseVal.IsValidInstanceError(
		ePrefix.XCpy(
			"baseVal"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr01.CopyIn(
		&baseVal,
		ePrefix.XCpy(
			"nStr01<-baseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr02.CopyIn(
		&baseVal,
		ePrefix.XCpy(
			"nStr02<-baseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr02.SetNumberSign(
		NumSignVal.Negative(),
		ePrefix.XCpy(
			"nStr02-NumSignVal.Negative()"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr01.Round(
		NumRoundType.HalfToEven(),
		0,
		ePrefix.XCpy(
			"nStr01-HalfToEven"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var actualNumberStr string

	actualNumberStr,
		_,
		err = nStr01.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#1 - Positive Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	err = nStr02.Round(
		NumRoundType.HalfToEven(),
		0,
		ePrefix.XCpy(
			"nStr02-HalfToEven"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumberStr,
		_,
		err = nStr02.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	expectedNumberStr = "-8"

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#2 - Negative Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	expectedNumberStr = "6"
	inputNumberStr = "6.5"

	nStr03,
		_,
		err = new(NumberStrKernel).
		NewParsePureNumberStr(
			inputNumberStr,
			".",
			true,
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy(
				"nStr03<-inputNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr03.Round(
		NumRoundType.HalfToEven(),
		0,
		ePrefix.XCpy(
			"nStr03-HalfToEven"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumberStr,
		_,
		err = nStr03.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr03"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#3 - Postive Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	expectedNumberStr = "-6"
	inputNumberStr = "-6.5"

	nStr04,
		_,
		err = new(NumberStrKernel).
		NewParsePureNumberStr(
			inputNumberStr,
			".",
			true,
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy(
				"nStr04<-inputNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr04.Round(
		NumRoundType.HalfToEven(),
		0,
		ePrefix.XCpy(
			"nStr04-HalfToEven"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumberStr,
		_,
		err = nStr04.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr04"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#4 - Positive Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	return
}

func TestNumberStrKernel_RoundHalfToOdd_000700(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_RoundHalfToOdd_000700",
		"")

	inputNumberStr := "7.5"
	expectedNumberStr := "7"

	var err error
	var baseVal, nStr01, nStr02,
		nStr03, nStr04 NumberStrKernel

	baseVal,
		_,
		err = new(NumberStrKernel).
		NewParsePureNumberStr(
			inputNumberStr,
			".",
			true,
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy(
				"baseVal<-inputNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = baseVal.IsValidInstanceError(
		ePrefix.XCpy(
			"baseVal"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr01.CopyIn(
		&baseVal,
		ePrefix.XCpy(
			"nStr01<-baseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr02.CopyIn(
		&baseVal,
		ePrefix.XCpy(
			"nStr02<-baseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr02.SetNumberSign(
		NumSignVal.Negative(),
		ePrefix.XCpy(
			"nStr02-NumSignVal.Negative()"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr01.Round(
		NumRoundType.HalfToOdd(),
		0,
		ePrefix.XCpy(
			"nStr01-HalfToOdd"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var actualNumberStr string

	actualNumberStr,
		_,
		err = nStr01.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#1 - Positive Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	err = nStr02.Round(
		NumRoundType.HalfToOdd(),
		0,
		ePrefix.XCpy(
			"nStr02-HalfToOdd"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumberStr,
		_,
		err = nStr02.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	expectedNumberStr = "-7"

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#2 - Negative Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	expectedNumberStr = "7"
	inputNumberStr = "6.5"

	nStr03,
		_,
		err = new(NumberStrKernel).
		NewParsePureNumberStr(
			inputNumberStr,
			".",
			true,
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy(
				"nStr03<-inputNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr03.Round(
		NumRoundType.HalfToOdd(),
		0,
		ePrefix.XCpy(
			"nStr03-HalfToOdd"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumberStr,
		_,
		err = nStr03.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr03"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#3 - Postive Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	expectedNumberStr = "-6"
	inputNumberStr = "-6.4"

	nStr04,
		_,
		err = new(NumberStrKernel).
		NewParsePureNumberStr(
			inputNumberStr,
			".",
			true,
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy(
				"nStr04<-inputNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr04.Round(
		NumRoundType.HalfToOdd(),
		0,
		ePrefix.XCpy(
			"nStr04-HalfToOdd"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumberStr,
		_,
		err = nStr04.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr04"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#4 - Negative Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	return
}

func TestNumberStrKernel_RoundRandomly_000800(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_RoundRandomly_000800()",
		"")

	inputNumberIntStr := "7"
	inputNumberFracStr := "5"
	var numberSign = NumSignVal.Positive()
	var loopCount = 50

	upperVal := "8"
	lowerVal := "7"
	var actualNumberStr string

	var upperValCount = 0
	var lowerValCount = 0

	var err error

	for i := 0; i < loopCount; i++ {

		var nStr01 NumberStrKernel

		nStr01,
			err = new(NumberStrKernel).NewFromStringDigits(
			inputNumberIntStr,
			inputNumberFracStr,
			numberSign,
			ePrefix.XCpy(
				"nStr01<-"))

		if err != nil {
			t.Errorf("\n%v\n",
				err.Error())
			return
		}

		err = nStr01.Round(
			NumRoundType.Randomly(),
			0,
			ePrefix.XCpy(
				"nStr01-Randomly"))

		if err != nil {
			t.Errorf("\n%v\n",
				err.Error())
			return
		}

		actualNumberStr,
			_,
			err = nStr01.FmtNumStrPure(
			".",
			true,
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy(
				"nStr01"))

		if err != nil {
			t.Errorf("\n%v\n",
				err.Error())
			return
		}

		if actualNumberStr == upperVal {
			upperValCount++
		}

		if actualNumberStr == lowerVal {
			lowerValCount++
		}

	}

	if upperValCount == loopCount {

		t.Errorf("%v\n"+
			"Test#1 - Random Rounding Test\n"+
			"Error: upperValCount == loopCount\n"+
			"Upper Value Rounding Result = '%v'"+
			"Number of roundings = '%v'\n"+
			"Number of results equal to Upper Value = '%v'\n",
			ePrefix.String(),
			upperVal,
			loopCount,
			upperValCount)

		return

	}

	if lowerValCount == loopCount {

		t.Errorf("%v\n"+
			"Test#1 - Random Rounding Test\n"+
			"Error: lowerValCount == loopCount\n"+
			"Lower Value Rounding Result = '%v'"+
			"Number of roundings = '%v'\n"+
			"Number of results equal to Lower Value = '%v'\n",
			ePrefix.String(),
			lowerVal,
			loopCount,
			lowerValCount)

		return

	}

	return
}

func TestNumberStrKernel_RoundFloor_000900(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_RoundFloor_000900",
		"")

	inputNumberStr := "2.9"
	expectedNumberStr := "2"

	var err error
	var baseVal, nStr01, nStr02,
		nStr03, nStr04 NumberStrKernel

	baseVal,
		_,
		err = new(NumberStrKernel).
		NewParsePureNumberStr(
			inputNumberStr,
			".",
			true,
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy(
				"baseVal<-inputNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = baseVal.IsValidInstanceError(
		ePrefix.XCpy(
			"baseVal"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr01.CopyIn(
		&baseVal,
		ePrefix.XCpy(
			"nStr01<-baseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr02.CopyIn(
		&baseVal,
		ePrefix.XCpy(
			"nStr02<-baseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr02.SetNumberSign(
		NumSignVal.Negative(),
		ePrefix.XCpy(
			"nStr02-NumSignVal.Negative()"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr01.Round(
		NumRoundType.Floor(),
		0,
		ePrefix.XCpy(
			"nStr01-Floor"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var actualNumberStr string

	actualNumberStr,
		_,
		err = nStr01.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#1 - Positive Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	err = nStr02.Round(
		NumRoundType.Floor(),
		0,
		ePrefix.XCpy(
			"nStr02-Floor"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumberStr,
		_,
		err = nStr02.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	expectedNumberStr = "-3"

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#2 - Negative Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	expectedNumberStr = "2"
	inputNumberStr = "2.4"

	nStr03,
		_,
		err = new(NumberStrKernel).
		NewParsePureNumberStr(
			inputNumberStr,
			".",
			true,
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy(
				"nStr03<-inputNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr03.Round(
		NumRoundType.Floor(),
		0,
		ePrefix.XCpy(
			"nStr03-Floor"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumberStr,
		_,
		err = nStr03.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr03"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#3 - Postive Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	expectedNumberStr = "-2"
	inputNumberStr = "-2"

	nStr04,
		_,
		err = new(NumberStrKernel).
		NewParsePureNumberStr(
			inputNumberStr,
			".",
			true,
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy(
				"nStr04<-inputNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr04.Round(
		NumRoundType.Floor(),
		0,
		ePrefix.XCpy(
			"nStr04-Floor"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumberStr,
		_,
		err = nStr04.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr04"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#4 - Negative Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	return
}

func TestNumberStrKernel_RoundCeiling_001000(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_RoundCeiling_001000",
		"")

	inputNumberStr := "2.9"
	expectedNumberStr := "3"

	var err error
	var baseVal, nStr01, nStr02,
		nStr03, nStr04 NumberStrKernel

	baseVal,
		_,
		err = new(NumberStrKernel).
		NewParsePureNumberStr(
			inputNumberStr,
			".",
			true,
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy(
				"baseVal<-inputNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = baseVal.IsValidInstanceError(
		ePrefix.XCpy(
			"baseVal"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr01.CopyIn(
		&baseVal,
		ePrefix.XCpy(
			"nStr01<-baseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr02.CopyIn(
		&baseVal,
		ePrefix.XCpy(
			"nStr02<-baseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr02.SetNumberSign(
		NumSignVal.Negative(),
		ePrefix.XCpy(
			"nStr02-NumSignVal.Negative()"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr01.Round(
		NumRoundType.Ceiling(),
		0,
		ePrefix.XCpy(
			"nStr01-Ceiling"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var actualNumberStr string

	actualNumberStr,
		_,
		err = nStr01.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#1 - Positive Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	err = nStr02.Round(
		NumRoundType.Ceiling(),
		0,
		ePrefix.XCpy(
			"nStr02-Ceiling"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumberStr,
		_,
		err = nStr02.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	expectedNumberStr = "-2"

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#2 - Negative Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	expectedNumberStr = "3"
	inputNumberStr = "2.4"

	nStr03,
		_,
		err = new(NumberStrKernel).
		NewParsePureNumberStr(
			inputNumberStr,
			".",
			true,
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy(
				"nStr03<-inputNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr03.Round(
		NumRoundType.Ceiling(),
		0,
		ePrefix.XCpy(
			"nStr03-Ceiling"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumberStr,
		_,
		err = nStr03.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#3 - Postive Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	expectedNumberStr = "-2"
	inputNumberStr = "-2"

	nStr04,
		_,
		err = new(NumberStrKernel).
		NewParsePureNumberStr(
			inputNumberStr,
			".",
			true,
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy(
				"nStr04<-inputNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr04.Round(
		NumRoundType.Ceiling(),
		0,
		ePrefix.XCpy(
			"nStr04-Ceiling"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumberStr,
		_,
		err = nStr04.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#4 - Negative Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	return
}

func TestNumberStrKernel_String_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_FmtSignedPureNumberStr_000100()",
		"")

	origIntStr := "1234"
	origFracStr := "5678"

	origNumberStr := origIntStr +
		"." +
		origFracStr

	expectedNumberStr := "1234.5678"

	var err error
	var baseValueNStr, nStr02 NumberStrKernel

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

	err = nStr02.CopyIn(
		&baseValueNStr,
		ePrefix.XCpy(
			"nStr02<-baseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if !nStr02.Equal(&baseValueNStr) {

		t.Errorf("%v\n"+
			"Test#1\n"+
			"Error: nStr02 NOT EQUAL TO baseValueNStr\n"+
			"After CopyIn() operation they were expected\n"+
			"to be Equal. HOWEVER, THEY ARE NOT EQUAL!!!\n",
			ePrefix.String())

		return
	}

	err = nStr02.SetDefaultPureNumStrFormatSpec(
		".",
		true,
		-1,
		TxtJustify.Right(),
		ePrefix.XCpy(
			"nStr02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var actualFmtNumberStr string

	actualFmtNumberStr,
		err = nStr02.FmtNumStrDefault(
		ePrefix.XCpy(
			"actualFmtNumberStr<-nStr02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != actualFmtNumberStr {

		t.Errorf("%v\n"+
			"Test#2\n"+
			"Error: actualfmtNumberStr NOT EQUAL TO expectedNumberStr\n"+
			" actualFmtNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualFmtNumberStr,
			expectedNumberStr)

		return
	}

	expectedNumberStr = "1,234.5678"

	err = nStr02.SetDefaultSimpleNumStrFormatSpec(
		".",
		",",
		true,
		-1,
		TxtJustify.Right(),
		ePrefix.XCpy(
			"nStr02"))

	actualFmtNumberStr,
		err = nStr02.FmtNumStrDefault(
		ePrefix.XCpy("actualFmtNumberStr" +
			"<-nStr02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != actualFmtNumberStr {

		t.Errorf("%v\n"+
			"Test#3\n"+
			"Error: actualfmtNumberStr NOT EQUAL TO expectedNumberStr\n"+
			" actualFmtNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualFmtNumberStr,
			expectedNumberStr)

		return
	}

	return
}
