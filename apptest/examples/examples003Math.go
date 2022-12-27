package examples

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"github.com/MikeAustin71/strmechops/strmech"
	"math/big"
	"strings"
)

type MainTest03 struct {
	input string
}

func (MainTest03) PureNumStrToFloat64() {
	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainTest03.PureNumStrToFloat64_01()",
		"")

	expectedStr := "1234.5678"

	testName := fmt.Sprintf("Test #1 TextFieldFormatDtoFloat64 - txtFieldFmtDtoBFloat(%v)\n",
		expectedStr)

	txtFieldFmtDtoF64 := strmech.TextFieldFormatDtoFloat64{
		LeftMarginStr:         "",
		Float64Num:            0,
		LeadingMinusSign:      true,
		RoundingType:          strmech.NumRoundType.HalfAwayFromZero(),
		NumOfFractionalDigits: -1,
		DefaultNumStrFmt:      strmech.NumStrFormatSpec{},
		FieldLength:           -1,
		FieldJustify:          strmech.TxtJustify.Right(),
		RightMarginStr:        "",
	}

	var err error

	err = txtFieldFmtDtoF64.SetFromNativeNumStr(
		expectedStr,
		ePrefix.XCpy(
			"txtFieldFmtDtoF64<-expectedStr"))

	if err != nil {
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	var actualNumStr string

	actualNumStr,
		err = new(strmech.MathHelper).NumericValueToNativeNumStr(
		txtFieldFmtDtoF64,
		ePrefix.XCpy(
			"actualNumStr<-txtFieldFmtDtoF64"))

	if err != nil {
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		fmt.Printf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	expectedStr = "-8761.123456"

	testName = fmt.Sprintf("Test #2 *TextFieldFormatDtoFloat64 - ptrTxtFieldFmtDtoF64(%v)\n",
		expectedStr)

	txtFieldFmtDtoF642 := strmech.TextFieldFormatDtoFloat64{
		LeftMarginStr:         "",
		Float64Num:            0,
		LeadingMinusSign:      true,
		RoundingType:          strmech.NumRoundType.HalfAwayFromZero(),
		NumOfFractionalDigits: -1,
		DefaultNumStrFmt:      strmech.NumStrFormatSpec{},
		FieldLength:           -1,
		FieldJustify:          strmech.TxtJustify.Right(),
		RightMarginStr:        "",
	}

	var ptrTxtFieldFmtDtoF64 *strmech.TextFieldFormatDtoFloat64

	err = txtFieldFmtDtoF642.SetFromNativeNumStr(
		expectedStr,
		ePrefix.XCpy(
			"txtFieldFmtDtoF642<-expectedStr"))

	if err != nil {
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	ptrTxtFieldFmtDtoF64 = &txtFieldFmtDtoF642

	actualNumStr,
		err = new(strmech.MathHelper).NumericValueToNativeNumStr(
		ptrTxtFieldFmtDtoF64,
		ePrefix.XCpy(
			"actualNumStr<-txtFieldFmtDtoF64"))

	if err != nil {
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		fmt.Printf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	return

}

func (MainTest03) RaiseToExponent03() {
	// Tests RaiseToIntExponent()
	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainTest03.RaiseToExponent03()",
		"")

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Starting Run!\n"+
		"Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

	baseStr := "-16.7653333333218"

	exponent := int64(0)

	expectedResultStr := "-16.7653333333218"

	var pureNumStrStats strmech.NumberStrStatsDto
	var err error

	pureNumStrStats,
		err = new(strmech.NumStrMath).PureNumStrStats(
		expectedResultStr,
		".",
		true,
		ePrefix)

	if err != nil {

		fmt.Printf("\n\n%v\n"+
			"%v\n",
			ePrefix.String(),
			err.Error())

		return

	}

	expectedFracDigits :=
		int(pureNumStrStats.NumOfFractionalDigits)

	var numOfExtraDigitsBuffer int64
	var precisionBitsOverride uint
	var roundingMode big.RoundingMode

	numOfExtraDigitsBuffer = 10
	precisionBitsOverride = 0
	roundingMode = big.AwayFromZero

	floatHelper := strmech.MathFloatHelper{}

	var bFloatDto strmech.BigFloatDto

	bFloatDto,
		err = floatHelper.PureNumStrToBigFloatDto(
		baseStr,
		".",
		true,
		numOfExtraDigitsBuffer,
		precisionBitsOverride,
		big.AwayFromZero,
		ePrefix)

	if err != nil {
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	fmt.Printf("\nInitialization\n"+
		"----- RaiseToIntPositiveExponent() -----\n"+
		"               baseStr = %v\n"+
		"          baseFloat -1 = %v\n"+
		"        base Precision = %v\n"+
		"numOfExtraDigitsBuffer = %v\n"+
		"             base Mode = %v\n"+
		"         base Accuracy = %v\n"+
		"              exponent = %v\n\n",
		baseStr,
		bFloatDto.Value.Text('f', -1),
		bFloatDto.Value.Prec(),
		numOfExtraDigitsBuffer,
		bFloatDto.Value.Mode(),
		bFloatDto.Value.Acc(),
		exponent)

	var raisedToIntExponent *big.Float

	raisedToIntExponent,
		err = floatHelper.RaiseToIntPositiveExponent(
		&bFloatDto.Value,
		exponent,
		numOfExtraDigitsBuffer,
		precisionBitsOverride,
		roundingMode,
		ePrefix.XCpy(
			"raisedToIntExponent"))

	if err != nil {
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	raisedToIntExponentStr :=
		raisedToIntExponent.Text('f', -1)

	fmt.Printf("After Calculation Of Raise To Power\n"+
		"----- RaiseToIntPositiveExponent() -----\n"+
		"Actual raisedToIntExponent -1 = %v\n"+
		"   Actual raisedToIntExponent = %v\n"+
		" Expected raisedToIntExponent = %v\n"+
		"raisedToIntExponent Precision = %v\n"+
		"       numOfExtraDigitsBuffer = %v\n"+
		"     raisedToIntExponent Mode = %v\n"+
		" raisedToIntExponent Accuracy = %v\n"+
		"                     exponent = %v\n\n",
		raisedToIntExponentStr,
		raisedToIntExponent.Text('f', expectedFracDigits),
		expectedResultStr,
		raisedToIntExponent.Prec(),
		numOfExtraDigitsBuffer,
		raisedToIntExponent.Mode(),
		raisedToIntExponent.Acc(),
		exponent)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		"Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

}

func (MainTest03) RaiseToExponent02() {
	// Compares results of RaisedToIntExponent
	// and RaisedToFloatExponent

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainTest03.RaiseToExponent02()",
		"")

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Starting Run!\n"+
		"Function: %v\n"+
		"Compares results of RaisedToIntExponent()\n"+
		"and RaisedToFloatExponent()",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

	// int digits 4
	// frac digits 12
	baseStr := "5084.987654321000"

	exponent := int64(4)

	expectedResultStr := "668589591687777.75101222860206783"

	var pureNumStrStats strmech.NumberStrStatsDto
	var err error

	pureNumStrStats,
		err = new(strmech.NumStrMath).PureNumStrStats(
		expectedResultStr,
		".",
		true,
		ePrefix)

	if err != nil {

		fmt.Printf("\n\n%v\n"+
			"%v\n",
			ePrefix.String(),
			err.Error())

		return

	}

	expectedFracDigits :=
		int(pureNumStrStats.NumOfFractionalDigits)

	var numOfExtraDigitsBuffer int64
	var precisionBitsOverride uint
	var roundingMode big.RoundingMode

	numOfExtraDigitsBuffer = 10
	precisionBitsOverride = 0
	roundingMode = big.AwayFromZero

	floatHelper := strmech.MathFloatHelper{}

	var bFloatDto strmech.BigFloatDto

	bFloatDto,
		err = floatHelper.PureNumStrToBigFloatDto(
		baseStr,
		".",
		true,
		numOfExtraDigitsBuffer,
		precisionBitsOverride,
		big.AwayFromZero,
		ePrefix)

	if err != nil {
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	fmt.Printf("\nInitialization\n"+
		"----- RaiseToIntPositiveExponent() -----\n"+
		"               baseStr = %v\n"+
		"             baseFloat = %v\n"+
		"        base Precision = %v\n"+
		"numOfExtraDigitsBuffer = %v\n"+
		"             base Mode = %v\n"+
		"         base Accuracy = %v\n"+
		"              exponent = %v\n\n",
		baseStr,
		bFloatDto.Value.Text('f', -1),
		bFloatDto.Value.Prec(),
		numOfExtraDigitsBuffer,
		bFloatDto.Value.Mode(),
		bFloatDto.Value.Acc(),
		exponent)

	var raisedToIntExponent *big.Float

	raisedToIntExponent,
		err = floatHelper.RaiseToIntPositiveExponent(
		&bFloatDto.Value,
		exponent,
		numOfExtraDigitsBuffer,
		precisionBitsOverride,
		roundingMode,
		ePrefix.XCpy(
			"raisedToIntExponent"))

	if err != nil {
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	raisedToIntExponentStr :=
		raisedToIntExponent.Text('f', expectedFracDigits)

	fmt.Printf("After Calculation Of Raise To Power\n"+
		"----- RaiseToIntPositiveExponent() -----\n"+
		"Actual raisedToIntExponent -1 = %v\n"+
		"   Actual raisedToIntExponent = %v\n"+
		" Expected raisedToIntExponent = %v\n"+
		"raisedToIntExponent Precision = %v\n"+
		"       numOfExtraDigitsBuffer = %v\n"+
		"     raisedToIntExponent Mode = %v\n"+
		" raisedToIntExponent Accuracy = %v\n"+
		"                     exponent = %v\n\n",
		raisedToIntExponent.Text('f', -1),
		raisedToIntExponentStr,
		expectedResultStr,
		raisedToIntExponent.Prec(),
		numOfExtraDigitsBuffer,
		raisedToIntExponent.Mode(),
		raisedToIntExponent.Acc(),
		exponent)

	fmt.Printf("\nInitialization\n"+
		"----- RaiseToFloatPositiveExponent() -----\n"+
		"               baseStr = %v\n"+
		"             baseFloat = %v\n"+
		"        base Precision = %v\n"+
		"numOfExtraDigitsBuffer = %v\n"+
		"             base Mode = %v\n"+
		"         base Accuracy = %v\n"+
		"              exponent = %v\n\n",
		baseStr,
		bFloatDto.Value.Text('f', -1),
		bFloatDto.Value.Prec(),
		numOfExtraDigitsBuffer,
		bFloatDto.Value.Mode(),
		bFloatDto.Value.Acc(),
		exponent)

	var raisedToFloatExponent *big.Float

	numOfExtraDigitsBuffer = 10
	precisionBitsOverride = 0

	raisedToFloatExponent,
		err = floatHelper.RaiseToFloatPositiveExponent(
		&bFloatDto.Value,
		exponent,
		numOfExtraDigitsBuffer,
		precisionBitsOverride,
		roundingMode,
		ePrefix)

	if err != nil {
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	raisedToFloatExpStr :=
		raisedToFloatExponent.Text('f', expectedFracDigits)

	fmt.Printf("After Calculation Of Raise To Power\n"+
		"----- RaiseToFloatPositiveExponent() -----\n"+
		"Actual raisedToFloatExponent -1 = %v\n\n"+
		"Actual raisedToFloatExponent    = %v\n"+
		"Expected raisedToFloatExponent  = %v\n"+
		"raisedToFloatExponent Precision = %v\n"+
		"         numOfExtraDigitsBuffer = %v\n"+
		"raisedToFloatExponent Mode      = %v\n"+
		"raisedToFloatExponent Accuracy  = %v\n"+
		"exponent                        = %v\n\n",
		raisedToFloatExponent.Text('f', -1),
		raisedToFloatExpStr,
		expectedResultStr,
		raisedToFloatExponent.Prec(),
		numOfExtraDigitsBuffer,
		raisedToFloatExponent.Mode(),
		raisedToFloatExponent.Acc(),
		exponent)

	fmt.Printf("\nResults Comparison\n"+
		"RaiseToIntPositiveExponent vs. RaiseToFloatPositiveExponent\n"+
		"Actual raisedToFloatExponent -1 = %v\n"+
		"Actual raisedToIntExponent -1   = %v\n"+
		"Actual raisedToFloatExponent    = %v\n"+
		"Actual raisedToIntExponent      = %v\n"+
		"Expected raisedToIntExponent    = %v\n"+
		"raisedToIntExponent Precision   = %v\n"+
		"raisedToIntExponent Mode        = %v\n"+
		"raisedToIntExponent Accuracy    = %v\n"+
		"raisedToFloatExponent Precision = %v\n"+
		"raisedToFloatExponent Mode      = %v\n"+
		"raisedToFloatExponent Accuracy  = %v\n"+
		"exponent                        = %v\n\n",
		raisedToFloatExponent.Text('f', -1),
		raisedToIntExponent.Text('f', -1),
		raisedToFloatExpStr,
		raisedToIntExponentStr,
		expectedResultStr,
		raisedToIntExponent.Prec(),
		raisedToIntExponent.Mode(),
		raisedToIntExponent.Acc(),
		raisedToFloatExponent.Prec(),
		raisedToFloatExponent.Mode(),
		raisedToFloatExponent.Acc(),
		exponent)

	if raisedToIntExponentStr !=
		raisedToFloatExpStr {

		fmt.Printf("\nFunction: %v\n"+
			"&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&\n"+
			"&&&&&&&      Comparision FAILED!      &&&&&&&\n"+
			"&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&\n"+
			"Result from RaiseToIntExponent NOT EQUAL TO\n"+
			"Result from RaiseToFracExponent!\n"+
			" RaiseToIntExponent = %v\n"+
			"RaiseToFracExponent = %v\n"+
			"&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&\n",
			ePrefix.String(),
			raisedToIntExponentStr,
			raisedToFloatExpStr)

		return
	}

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\nFunction: %v\n"+
		"!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!\n"+
		"!!!!!!!    Comparision SUCCESSFUL     !!!!!!!\n"+
		"!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!\n"+
		"Result from RaiseToIntExponent EQUAL TO\n"+
		"Result from RaiseToFracExponent!\n"+
		" RaiseToIntExponent = %v\n"+
		"RaiseToFracExponent = %v\n"+
		"!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!\n",
		ePrefix.String(),
		raisedToIntExponentStr,
		raisedToFloatExpStr)

	fmt.Printf("\n" + breakStr + "\n")

}

func (MainTest03) RaiseToExponent01() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainTest03.RaiseToExponent01()",
		"")

	breakStr := strings.Repeat("=", 60)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Starting Run!\n"+
		"Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

	baseStr := "5"

	exponent := int64(2)

	expectedResultStr := "25"

	var pureNumStrStats strmech.NumberStrStatsDto
	var err error

	pureNumStrStats,
		err = new(strmech.NumStrMath).PureNumStrStats(
		expectedResultStr,
		".",
		true,
		ePrefix)

	if err != nil {

		fmt.Printf("\n\n%v\n"+
			"%v\n",
			ePrefix.String(),
			err.Error())

		return

	}

	expectedFracDigits :=
		int(pureNumStrStats.NumOfFractionalDigits)

	floatHelper := strmech.MathFloatHelper{}

	var bFloatDto strmech.BigFloatDto

	bFloatDto,
		err = floatHelper.PureNumStrToBigFloatDto(
		baseStr,
		".",
		true,
		50,
		0,
		big.AwayFromZero,
		ePrefix)

	if err != nil {

		fmt.Printf("\n\n%v\n"+
			"%v\n",
			ePrefix.String(),
			err.Error())

		return

	}

	bFloatDtoValueStr := bFloatDto.Value.Text('f', -1)

	fmt.Printf("\nInitialization\n"+
		"baseStr         = %v\n"+
		"baseFloat       = %v\n"+
		"base Precision  = %v\n"+
		"base Accuracy   = %v\n"+
		"base Round Mode = %v\n"+
		"base is Integer = %v\n"+
		"exponent        = %v\n\n",
		baseStr,
		bFloatDtoValueStr,
		bFloatDto.Value.Prec(),
		bFloatDto.Value.Acc(),
		bFloatDto.Value.Mode(),
		bFloatDto.Value.IsInt(),
		exponent)

	var raisedToExponent *big.Float

	raisedToExponent,
		err = floatHelper.RaiseToFloatPositiveExponent(
		&bFloatDto.Value,
		exponent,
		100,
		0,
		big.AwayFromZero,
		ePrefix.XCpy(
			"raisedToExponent"))

	if err != nil {
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	raisedToExponentStr :=
		raisedToExponent.Text('f', expectedFracDigits)

	pureNumStrStats,
		err = new(strmech.NumStrMath).PureNumStrStats(
		raisedToExponentStr,
		".",
		true,
		ePrefix)

	if err != nil {

		fmt.Printf("\n\n%v\n"+
			"%v\n",
			ePrefix.String(),
			err.Error())

		return

	}

	actualRaisedToExponentStr :=
		raisedToExponent.Text('f', -1)

	var actualRaisedToExpStats strmech.NumberStrStatsDto

	actualRaisedToExpStats,
		err = new(strmech.NumStrMath).PureNumStrStats(
		actualRaisedToExponentStr,
		".",
		true,
		ePrefix)

	if err != nil {

		fmt.Printf("\n\n%v\n"+
			"%v\n",
			ePrefix.String(),
			err.Error())

		return

	}

	fmt.Printf("After Calculation Of Raise To Power\n"+
		"    ---- Raise To Float Exponent ----\n"+
		"Actual raisedToExponent      = %v\n"+
		"Expected raisedToExponent    = %v\n"+
		"raisedToExponent Precision   = %v\n"+
		"raisedToExponent Accuracy    = %v\n"+
		"raisedToExponent Mode        = %v\n"+
		"raisedToExponent is Integer  = %v\n"+
		"raisedToExponent Int Digits  = %v\n"+
		"raisedToExponent Frac Digits = %v\n"+
		"Actual raised Int Digits     = %v\n"+
		"Actual raised Frac Digits    = %v\n\n",
		raisedToExponentStr,
		expectedResultStr,
		raisedToExponent.Prec(),
		raisedToExponent.Acc(),
		raisedToExponent.Mode(),
		raisedToExponent.IsInt(),
		pureNumStrStats.NumOfIntegerDigits,
		pureNumStrStats.NumOfFractionalDigits,
		actualRaisedToExpStats.NumOfIntegerDigits,
		actualRaisedToExpStats.NumOfFractionalDigits)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		"Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

}

func (MainTest03) RoundBigFloat01() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainTest03.RoundBigFloat01()",
		"")

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Starting Run!\n"+
		"Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

	requiredDigits := int64(150)

	floatHelper := strmech.MathFloatHelper{}

	requiredPrecision,
		err := floatHelper.
		DigitsToPrecisionEstimate(
			requiredDigits,
			ePrefix.XCpy(
				""))

	if err != nil {
		fmt.Printf("%v\n",
			err.Error())
		return
	}

	if requiredPrecision != 504 {

		fmt.Printf("\n%v\n"+
			"Error: Expected 504 bits of precision from "+
			"150 digits.\n"+
			"Instead, requiredPrecision = '%v'\n",
			ePrefix.String(),
			requiredPrecision)

		return
	}

	test01Float :=
		new(big.Float).
			SetPrec(requiredPrecision).
			SetMode(big.AwayFromZero).
			SetInt64(0)

	test02Float :=
		new(big.Float).
			SetPrec(requiredPrecision).
			SetMode(big.AwayFromZero).
			SetInt64(0)

	originalValueStr := "123456.78125"

	var ok bool

	_,
		ok = test01Float.SetString(originalValueStr)

	if !ok {

		fmt.Printf("\n%v\n"+
			"Error:\n"+
			"test01Float.SetString(originalValueStr) FAILED!\n"+
			"originalValueStr = '%v'\n",
			ePrefix.String(),
			originalValueStr)

		return
	}

	err = floatHelper.RoundBigFloat(
		test01Float,
		test02Float,
		strmech.NumRoundType.HalfAwayFromZero(),
		4,
		ePrefix.XCpy(
			"test02Float<- to 4-Frac digits"))

	actualRoundedValueStr :=
		test02Float.Text('f', 5)

	expectedRoundedValueStr := "123456.78130"

	if actualRoundedValueStr != expectedRoundedValueStr {

		fmt.Printf("\n%v\n"+
			"Error: originalValueStr = '%v'\n"+
			"Expected a rounded value      = '%v'\n"+
			"Instead, actual rounded value = '%v'\n\n",
			ePrefix.String(),
			originalValueStr,
			expectedRoundedValueStr,
			actualRoundedValueStr)

		return
	}

	fmt.Printf("\n\n%v\n"+
		"Successful Rounding Operation!\n"+
		"originalValueStr = '%v'\n"+
		"Expected a rounded value of   = '%v'\n"+
		"Matching Actual rounded value = '%v'\n\n",
		ePrefix.String(),
		originalValueStr,
		expectedRoundedValueStr,
		actualRoundedValueStr)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		"Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

	return
}

func (MainTest03) EstimatePrecisionBits01() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainTest03.EstimatePrecisionBits01()",
		"")

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Starting Run!\n"+
		"Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

	requiredDigits := int64(150)

	floatHelper := strmech.MathFloatHelper{}

	requiredPrecision,
		err := floatHelper.
		DigitsToPrecisionEstimate(
			requiredDigits,
			ePrefix.XCpy(
				""))

	if err != nil {
		fmt.Printf("%v\n",
			err.Error())
		return
	}

	if requiredPrecision != 504 {

		fmt.Printf("\n%v\n"+
			"Error: Expected 504 bits of precision from "+
			"150 digits.\n"+
			"Instead, requiredPrecision = '%v'\n",
			ePrefix.String(),
			requiredPrecision)

		return
	}

	fmt.Printf("requiredPrecision = %v\n"+
		"expected required precision = %v\n"+
		"SUCCESS!\n\n",
		requiredPrecision,
		504)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		"Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

}
