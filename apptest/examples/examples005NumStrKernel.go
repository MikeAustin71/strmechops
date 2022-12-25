package examples

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"github.com/MikeAustin71/strmechops/strmech"
	"math/big"
	"strings"
)

type MainNumStrTest005 struct {
	input string
}

func (mainNumStrTest005 MainNumStrTest005) Float64Dto01() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainNumStrTest005.Float64Dto01()",
		"")

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Starting Run!\n"+
		"Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

	var txtFmtFloat64 strmech.TextFieldFormatDtoFloat64

	txtFmtFloat64 = strmech.TextFieldFormatDtoFloat64{
		LeftMarginStr:         "",
		Float64Num:            12345.123456789,
		LeadingMinusSign:      true,
		RoundingType:          strmech.NumRoundType.HalfAwayFromZero(),
		NumOfFractionalDigits: 6,
		DefaultNumStrFmt:      strmech.NumStrFormatSpec{},
		FieldLength:           -1,
		FieldJustify:          strmech.TxtJustify.Right(),
		RightMarginStr:        "",
	}

	numStrFmtSpec,
		err := new(strmech.NumStrFormatSpec).NewNumFmtParams(
		".",
		",",
		strmech.IntGroupingType.Thousands(),
		"",
		"",
		strmech.NumFieldSymPos.InsideNumField(),
		"-",
		"",
		strmech.NumFieldSymPos.InsideNumField(),
		"",
		"",
		strmech.NumFieldSymPos.InsideNumField(),
		23,
		strmech.TxtJustify.Center(),
		ePrefix.XCpy(
			"numStrFmtSpec<-"))

	if err != nil {
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	expectedStr := "     12,345.123457     "

	sMech := strmech.StrMech{}

	convertedExpectedStr := sMech.ConvertNonPrintableString(
		expectedStr,
		true)

	err = txtFmtFloat64.SetNumStrFmtDefault(
		numStrFmtSpec,
		ePrefix.XCpy(
			"txtFmtFloat64<-numStrFmtSpec"))

	if err != nil {
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	var testNumStrFmtSpec strmech.NumStrFormatSpec

	testNumStrFmtSpec,
		err = txtFmtFloat64.GetDefaultNumStrFmtSpec(
		ePrefix.XCpy(
			"testNumStrFmtSpec<-"))

	var fieldSpec strmech.NumStrNumberFieldSpec

	fieldSpec,
		err = testNumStrFmtSpec.GetNumberFieldSpec(
		ePrefix.XCpy(
			"testNumStrFmtSpec"))

	if err != nil {
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	fmt.Printf("Field Length: %v\n"+
		"Field Justification: %v\n",
		fieldSpec.GetNumFieldLength(),
		fieldSpec.GetNumFieldJustification().String())

	var actualStr string

	actualStr,
		err = txtFmtFloat64.FmtNumStrDefault(
		ePrefix.XCpy(
			""))

	convertedActualStr := sMech.ConvertNonPrintableString(
		actualStr,
		true)

	if convertedActualStr != convertedExpectedStr {

		fmt.Printf("%v\n"+
			"Test#1 - txtFmtFloat64.SetNumStrFmtDefault()\n"+
			"Error: convertedActualStr NOT EQUAL TO convertedExpectedStr\n"+
			"    convertedActualStr = '%v'\n"+
			"convertedExpectedStr   = '%v'\n",
			ePrefix.String(),
			convertedActualStr,
			convertedExpectedStr)

		return

	}

	fmt.Printf("%v\n"+
		"Success: convertedActualStr EQUAL TO convertedExpectedStr\n"+
		"    convertedActualStr = '%v'\n"+
		"convertedExpectedStr   = '%v'\n",
		ePrefix.String(),
		convertedActualStr,
		convertedExpectedStr)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		"Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

	return

}

func (mainNumStrTest005 MainNumStrTest005) NumStrRound01() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainNumStrTest005.NumStrFmtSignedSimple()",
		"")

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Starting Run!\n"+
		"Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

	origIntStr := "1234"
	origFracStr := "5678"

	origNumberStr := origIntStr +
		"." +
		origFracStr

	expectedNumberStr := "1234.568"

	var err error
	var baseValueNStr strmech.NumberStrKernel

	baseValueNStr,
		err = new(strmech.NumberStrKernel).
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

	err = baseValueNStr.Round(
		strmech.NumRoundType.HalfAwayFromZero(),
		3,
		ePrefix.XCpy(
			"baseValueNStr,HalfAwayFromZero,3-digits"))

	if err != nil {
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	var intDigitsDto, fracDigitsDto strmech.RuneArrayDto

	intDigitsDto,
		err = baseValueNStr.GetIntegerDigits(
		ePrefix.XCpy(
			"baseValueNStr"))

	if err != nil {
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	fracDigitsDto,
		err = baseValueNStr.GetFractionalDigits(
		ePrefix.XCpy(
			"baseValueNStr"))

	if err != nil {
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	var numberSign strmech.NumericSignValueType

	numberSign,
		err = baseValueNStr.GetNumberSign(
		ePrefix.XCpy(
			"baseValueNStr"))

	if err != nil {
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	fmt.Printf("\n\n%v\n"+
		"NumberStrKernel Digits After Rounding\n"+
		"   Integer Digits: '%v'\n"+
		"Fractional Digits: '%v'\n"+
		"Number Sign: '%v'\n",
		ePrefix.String(),
		intDigitsDto.GetCharacterString(),
		fracDigitsDto.GetCharacterString(),
		numberSign.String())

	roundingSpec := strmech.NumStrRoundingSpec{}

	roundingSpec,
		err = new(strmech.NumStrRoundingSpec).NewRoundingSpec(
		strmech.NumRoundType.NoRounding(),
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
		err = baseValueNStr.FmtSignedNumStrSimple(
		".",
		"",
		true,
		-1,
		strmech.TxtJustify.Right(),
		roundingSpec,
		ePrefix.XCpy(
			"fmtNumberStr<-baseValueNStr"))

	if err != nil {
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("%v\n"+
		"Expected NumStr = '%v'\n"+
		"  Actual NumStr = '%v'\n",
		ePrefix.String(),
		expectedNumberStr,
		fmtNumberStr)

	if expectedNumberStr != fmtNumberStr {

		fmt.Printf("\n\n%v\n"+
			"** Error **\n"+
			"Expected Number String NOT EQUAL\n"+
			"to Actual Number String!\n"+
			"Expected NumStr = '%v'\n"+
			"  Actual NumStr = '%v'\n",
			ePrefix.String(),
			expectedNumberStr,
			fmtNumberStr)

		return

	}

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		"Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

}

func (mainNumStrTest005 MainNumStrTest005) NumStrFmtSignedSimple() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainNumStrTest005.NumStrFmtSignedSimple()",
		"")

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Starting Run!\n"+
		"Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

	origIntStr := "1234"
	origFracStr := "5678"

	origNumberStr := origIntStr +
		"." +
		origFracStr

	expectedNumberStr := "1,234.5678"

	var err error
	var baseValueNStr strmech.NumberStrKernel

	baseValueNStr,
		err = new(strmech.NumberStrKernel).
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

	roundingSpec := strmech.NumStrRoundingSpec{}

	roundingSpec,
		err = new(strmech.NumStrRoundingSpec).NewRoundingSpec(
		strmech.NumRoundType.NoRounding(),
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
		err = baseValueNStr.FmtSignedNumStrSimple(
		".",
		",",
		true,
		-1,
		strmech.TxtJustify.Right(),
		roundingSpec,
		ePrefix.XCpy(
			"fmtNumberStr<-baseValueNStr"))

	if err != nil {
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("%v\n"+
		"Expected NumStr = '%v'\n"+
		"  Actual NumStr = '%v'\n",
		ePrefix.String(),
		expectedNumberStr,
		fmtNumberStr)

	if expectedNumberStr != fmtNumberStr {

		fmt.Printf("\n\n%v\n"+
			"** Error **\n"+
			"Expected Number String NOT EQUAL\n"+
			"to Actual Number String!\n"+
			"Expected NumStr = '%v'\n"+
			"  Actual NumStr = '%v'\n",
			ePrefix.String(),
			expectedNumberStr,
			fmtNumberStr)

		return

	}

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		"Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

}

func (mainNumStrTest005 MainNumStrTest005) NumberStrKernelCompare01() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainNumStrTest005.NumberStrKernelCompare01()",
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
		err = new(strmech.NumberStrKernel).NewFromNumericValue(
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

func (mainNumStrTest005 MainNumStrTest005) NumberStrKernelPureNumStr01() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainNumStrTest005.NumberStrKernelPureNumStr01()",
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

func (mainNumStrTest005 MainNumStrTest005) NumberStrKernelExtendArrays01() {

	funcName := "\nMainNumStrTest005.NumberStrKernelExtendArrays01()"

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

func (mainNumStrTest005 MainNumStrTest005) NumStrKernelParseUSNumStr() {

	funcName := "MainNumStrTest005.NumStrKernelParseUSNumStr()"

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

func (mainNumStrTest005 MainNumStrTest005) NumStrKernelSetSignedIntValue() {

	funcName := "MainNumStrTest005.NumStrKernelSetSignedIntValue()"

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
		err := sMechNStrKernel.NewFromNumericValue(
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

func (mainNumStrTest005 MainNumStrTest005) NumStrKernelToFloatConversion() {

	funcName := "MainNumStrTest005.NumStrKernelToFloatConversion()"

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

func (mainNumStrTest005 MainNumStrTest005) NumStrKernelToIntConversion() {

	funcName := "MainNumStrTest005.NumStrKernelToIntConversion()"

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
