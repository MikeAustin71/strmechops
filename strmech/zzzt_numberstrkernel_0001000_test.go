package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
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
		err = new(NumberStrKernel).NewFromFloatValue(
		testBigFloat,
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
		err = new(NumberStrKernel).NewParsePureNumberStr(
		testValue,
		".",
		true,
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
		err = new(NumberStrKernel).NewParsePureNumberStr(
		testValue,
		".",
		true,
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
		err = new(NumberStrKernel).NewParsePureNumberStr(
		origValueStr,
		".",
		true,
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
		err = new(NumberStrKernel).NewParsePureNumberStr(
		testValue,
		".",
		true,
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
		err = new(NumberStrKernel).NewFromUnsignedIntValue(
		origNum,
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
		err = new(NumberStrKernel).NewFromUnsignedIntValue(
		uintNum,
		NumSignVal.Positive(),
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
		err = new(NumberStrKernel).NewFromUnsignedIntValue(
		uint32Num,
		NumSignVal.Positive(),
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

	pureNumStr = numStrKernelBase.GetPureNumberStr(
		".",
		true)

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

	var bigRatNum *big.Rat

	bigRatNum,
		err = numStrKernelBase.GetBigRatNum(
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
		err = new(NumberStrKernel).NewParsePureNumberStr(
		compositeNumStr,
		".",
		true,
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
		NewFromBigRat(
			newRat,
			10,
			ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumberStr := numStrKernel01.GetPureNumberStr(
		".",
		true)

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
		err = new(NumberStrKernel).
		NewParsePureNumberStr(
			origNumberStr,
			".",
			true,
			ePrefix.XCpy(
				"baseValueNStr<-origNumberStr"))

	if err != nil {
		fmt.Printf("\n%v\n",
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
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	var fmtNumberStr string

	fmtNumberStr,
		err = baseValueNStr.FmtSignedSimpleNumberStr(
		".",
		",",
		true,
		-1,
		TxtJustify.Right(),
		roundingSpec,
		ePrefix.XCpy(
			"fmtNumberStr<-baseValueNStr"))

	if err != nil {
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != fmtNumberStr {

		t.Errorf("%v\n"+
			"Test#1\n"+
			"Error: actualfmtNumberStr NOT EUQAL TO expectedNumberStr\n"+
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
		err = new(NumberStrKernel).
		NewParsePureNumberStr(
			origNumberStr,
			".",
			true,
			ePrefix.XCpy(
				"baseValueNStr<-origNumberStr"))

	if err != nil {
		fmt.Printf("\n%v\n",
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
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	var fmtNumberStr string

	fmtNumberStr,
		err = baseValueNStr.FmtSignedSimpleNumberStr(
		".",
		"",
		true,
		-1,
		TxtJustify.Right(),
		roundingSpec,
		ePrefix.XCpy(
			"fmtNumberStr<-baseValueNStr"))

	if err != nil {
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != fmtNumberStr {

		t.Errorf("%v\n"+
			"Test#1\n"+
			"Error: actualfmtNumberStr NOT EUQAL TO expectedNumberStr\n"+
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
		err = new(NumberStrKernel).
		NewParsePureNumberStr(
			origNumberStr,
			".",
			true,
			ePrefix.XCpy(
				"baseValueNStr<-origNumberStr"))

	if err != nil {
		fmt.Printf("\n%v\n",
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
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	var fmtNumberStr string

	fmtNumberStr,
		err = baseValueNStr.FmtSignedPureNumberStr(
		".",
		true,
		-1,
		TxtJustify.Right(),
		roundingSpec,
		ePrefix.XCpy(
			"fmtNumberStr<-baseValueNStr"))

	if err != nil {
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != fmtNumberStr {

		t.Errorf("%v\n"+
			"Test#1\n"+
			"Error: actualfmtNumberStr NOT EUQAL TO expectedNumberStr\n"+
			" actualfmtNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			fmtNumberStr,
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
		err = new(NumberStrKernel).
		NewParsePureNumberStr(
			origNumberStr,
			".",
			true,
			ePrefix.XCpy(
				"baseValueNStr<-origNumberStr"))

	if err != nil {
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	err = nStr02.CopyIn(
		&baseValueNStr,
		ePrefix.XCpy(
			"nStr02<-baseValueNStr"))

	if err != nil {
		fmt.Printf("\n%v\n",
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
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	actualFmtNumberStr := nStr02.String()

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

	actualFmtNumberStr = nStr02.String()

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
}
