package examples

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"github.com/MikeAustin71/strmechops/strmech"
	"math/big"
	"strings"
)

type MainNumStrTest006 struct {
	input string
}

func (mainNumStrTest006 MainNumStrTest006) BigFloatDto03() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"mainNumStrTest006.BigFloatDto03()",
		"")

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Starting Run!\n"+
		"Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

	mtTest6Mech := MainNumStrTest006Mechanics{}

	var isNegativeValue bool

	isNegativeValue = false

	var numberSignDescription string

	if isNegativeValue == true {

		numberSignDescription = "Negative Number"
	} else {
		numberSignDescription = "Positive Number"
	}

	expectedStr := mtTest6Mech.GetTestNativeNumStr(
		1,
		0,
		isNegativeValue)

	testName := fmt.Sprintf("Test #1 %v: TextFieldFormatDtoBigFloat - txtFieldFmtDtoBFloat(%v)\n",
		numberSignDescription,
		expectedStr)

	txtFieldFmtDtoBFloat := strmech.TextFieldFormatDtoBigFloat{
		LeftMarginStr:         "",
		BigFloatNum:           big.Float{},
		LeadingMinusSign:      false,
		NativeRoundingMode:    big.AwayFromZero,
		RoundingType:          strmech.NumRoundType.HalfAwayFromZero(),
		NumOfFractionalDigits: -1,
		DefaultNumStrFmt:      strmech.NumStrFormatSpec{},
		FieldLength:           -1,
		FieldJustify:          strmech.TxtJustify.Right(),
		RightMarginStr:        "",
	}

	var err error

	err = txtFieldFmtDtoBFloat.SetFromNativeNumStr(
		expectedStr,
		ePrefix.XCpy(
			"txtFieldFmtDtoBFloat<-expectedStr"))

	if err != nil {
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	var actualNumStr string

	actualNumStr,
		err = new(strmech.MathHelper).NumericValueToNativeNumStr(
		txtFieldFmtDtoBFloat,
		ePrefix.XCpy(
			"actualNumStr<-txtFieldFmtDtoBFloat"))

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

	isNegativeValue = true

	expectedStr = mtTest6Mech.GetTestNativeNumStr(
		1,
		0,
		isNegativeValue)

	if isNegativeValue == true {

		numberSignDescription = "Negative Number"
	} else {
		numberSignDescription = "Positive Number"
	}

	testName = fmt.Sprintf("Test #2: %v *TextFieldFormatDtoBigFloat - ptrTxtFieldFmtDtoBFloat2(%v)\n",
		numberSignDescription,
		expectedStr)

	txtFieldFmtDtoBFloat2 := strmech.TextFieldFormatDtoBigFloat{
		LeftMarginStr:         "",
		BigFloatNum:           big.Float{},
		LeadingMinusSign:      false,
		NativeRoundingMode:    big.AwayFromZero,
		RoundingType:          strmech.NumRoundType.HalfAwayFromZero(),
		NumOfFractionalDigits: -1,
		DefaultNumStrFmt:      strmech.NumStrFormatSpec{},
		FieldLength:           -1,
		FieldJustify:          strmech.TxtJustify.Right(),
		RightMarginStr:        "",
	}

	err = txtFieldFmtDtoBFloat2.SetFromNativeNumStr(
		expectedStr,
		ePrefix.XCpy(
			"txtFieldFmtDtoBFloat2<-expectedStr"))

	if err != nil {
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	/*	fmt.Printf("\n\n%v\n"+
		"Initial Value from txtFieldFmtDtoBFloat2.SetFromNativeNumStr()\n"+
		"txtFieldFmtDtoBFloat2.BigFloatNum = %v\n\n",
		ePrefix.String(),
		txtFieldFmtDtoBFloat2.BigFloatNum.Text('f', -1))
	*/

	actualNumStr,
		err = new(strmech.MathHelper).NumericValueToNativeNumStr(
		&txtFieldFmtDtoBFloat2,
		ePrefix.XCpy(
			"actualNumStr<-ptrTxtFieldFmtDtoBFloat2"))

	if err != nil {
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		fmt.Printf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = \n'%v'\n\n"+
			"expectedStr   = \n'%v'\n\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	fmt.Printf("\n\n" + breakStr + "\n\n")

	fmt.Printf("\n%v\n"+
		"SUCCESS: actualNumStr == expectedStr\n"+
		"actualNumStr  = \n'%v'\n\n"+
		"expectedStr   = \n'%v'\n\n",
		ePrefix.String(),
		actualNumStr,
		expectedStr)

	fmt.Printf("\n\n" + breakStr + "\n\n")

	fmt.Printf("\n Successful Completion!\n"+
		"Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

	return

}

func (mainNumStrTest006 MainNumStrTest006) BigFloatDto02() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"mainNumStrTest006.BigFloatDto01()",
		"")

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Starting Run!\n"+
		"Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

	expectedStr := "1234.5678"

	testName := fmt.Sprintf("Test #1 TextFieldFormatDtoBigFloat - txtFieldFmtDtoBFloat(%v)\n",
		expectedStr)

	txtFieldFmtDtoBFloat := strmech.TextFieldFormatDtoBigFloat{
		LeftMarginStr:         "",
		BigFloatNum:           big.Float{},
		LeadingMinusSign:      false,
		NativeRoundingMode:    big.AwayFromZero,
		RoundingType:          strmech.NumRoundType.HalfAwayFromZero(),
		NumOfFractionalDigits: -1,
		DefaultNumStrFmt:      strmech.NumStrFormatSpec{},
		FieldLength:           -1,
		FieldJustify:          strmech.TxtJustify.Right(),
		RightMarginStr:        "",
	}

	var err error

	err = txtFieldFmtDtoBFloat.SetFromNativeNumStr(
		expectedStr,
		ePrefix.XCpy(
			"txtFieldFmtDtoBFloat<-expectedStr"))

	if err != nil {
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	var actualNumStr string

	actualNumStr,
		err = new(strmech.MathHelper).NumericValueToNativeNumStr(
		txtFieldFmtDtoBFloat,
		ePrefix.XCpy(
			"actualNumStr<-txtFieldFmtDtoBFloat"))

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

	expectedStr = "-87654321.12345678901234567892012345678901"

	testName = fmt.Sprintf("Test #2 Negative Number *TextFieldFormatDtoBigFloat - ptrTxtFieldFmtDtoBFloat2(%v)\n",
		expectedStr)

	txtFieldFmtDtoBFloat2 := strmech.TextFieldFormatDtoBigFloat{
		LeftMarginStr:         "",
		BigFloatNum:           big.Float{},
		LeadingMinusSign:      false,
		NativeRoundingMode:    big.AwayFromZero,
		RoundingType:          strmech.NumRoundType.HalfAwayFromZero(),
		NumOfFractionalDigits: -1,
		DefaultNumStrFmt:      strmech.NumStrFormatSpec{},
		FieldLength:           -1,
		FieldJustify:          strmech.TxtJustify.Right(),
		RightMarginStr:        "",
	}

	err = txtFieldFmtDtoBFloat2.SetFromNativeNumStr(
		expectedStr,
		ePrefix.XCpy(
			"txtFieldFmtDtoBFloat2<-expectedStr"))

	if err != nil {
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	fmt.Printf("\n\n%v\n"+
		"Initial Value from txtFieldFmtDtoBFloat2.SetFromNativeNumStr()\n"+
		"txtFieldFmtDtoBFloat2.BigFloatNum = %v\n\n",
		ePrefix.String(),
		txtFieldFmtDtoBFloat2.BigFloatNum.Text('f', -1))

	actualNumStr,
		err = new(strmech.MathHelper).NumericValueToNativeNumStr(
		&txtFieldFmtDtoBFloat2,
		ePrefix.XCpy(
			"actualNumStr<-ptrTxtFieldFmtDtoBFloat2"))

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

	fmt.Printf("\n\n" + breakStr + "\n\n")

	fmt.Printf("\n%v\n"+
		"SUCCESS: actualNumStr == expectedStr\n"+
		"actualNumStr  = '%v'\n"+
		"expectedStr   = '%v'\n",
		ePrefix.String(),
		actualNumStr,
		expectedStr)

	fmt.Printf("\n\n" + breakStr + "\n\n")

	fmt.Printf("\n Successful Completion!\n"+
		"Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

	return

}

func (mainNumStrTest006 MainNumStrTest006) BigFloatDto01() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"mainNumStrTest006.BigFloatDto01()",
		"")

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Starting Run!\n"+
		"Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

	expectedStr := "1234567.890123456789992345678"

	testName := fmt.Sprintf("Test #1 TextFieldFormatDtoBigFloat - txtFieldFmtDtoBFloat = (%v)\n",
		expectedStr)

	txtFieldFmtDtoBFloat := strmech.TextFieldFormatDtoBigFloat{
		LeftMarginStr:         "",
		BigFloatNum:           big.Float{},
		LeadingMinusSign:      true,
		NativeRoundingMode:    big.AwayFromZero,
		RoundingType:          strmech.NumRoundType.HalfAwayFromZero(),
		NumOfFractionalDigits: -1,
		DefaultNumStrFmt:      strmech.NumStrFormatSpec{},
		FieldLength:           -1,
		FieldJustify:          strmech.TxtJustify.Right(),
		RightMarginStr:        "",
	}

	var err error

	err = new(strmech.MathHelper).NativeNumStrToNumericValue(
		expectedStr,
		&txtFieldFmtDtoBFloat,
		ePrefix.XCpy(
			"#1 txtFieldFmtDtoBFloat<-expectedStr"))

	if err != nil {
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	auxActualStr := txtFieldFmtDtoBFloat.BigFloatNum.Text(
		'f',
		-1)

	fmt.Printf("\n%v\n"+
		"Initial value of txtFieldFmtDtoBFloat.BigFloatNum\n"+
		"BigFloatNum = '%v'\n\n",
		ePrefix.String(),
		auxActualStr)

	var actualNumStr string

	actualNumStr,
		err = txtFieldFmtDtoBFloat.
		FmtNumStrNative(
			ePrefix.XCpy(
				"#1 actualNumStr<-txtFieldFmtDtoBFloat"))

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

	fmt.Printf("\n\n" + breakStr + "\n\n")

	fmt.Printf("\n%v\n"+
		"SUCCESS: actualNumStr == expectedStr\n"+
		"actualNumStr  = '%v'\n"+
		"expectedStr   = '%v'\n",
		ePrefix.String(),
		actualNumStr,
		expectedStr)

	fmt.Printf("\n\n" + breakStr + "\n\n")

	fmt.Printf("\n Successful Completion!\n"+
		"Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

	return

}

func (mainNumStrTest006 MainNumStrTest006) Float64Dto01() {

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

type MainNumStrTest006Mechanics struct {
	input string
}

func (mainNumStrTest006Mech MainNumStrTest006Mechanics) GetTestNativeNumStr(
	integerDigits int,
	fractionalDigits int,
	isNegativeValue bool) string {

	var nativeNumStr []rune

	if isNegativeValue {

		nativeNumStr = append(nativeNumStr, '-')

	}

	nextDigit := '1'

	if integerDigits < 1 {
		nativeNumStr = append(nativeNumStr, '0')
	} else {

		for i := 0; i < integerDigits; i++ {

			nativeNumStr = append(nativeNumStr, nextDigit)

			nextDigit += 1

			if nextDigit > '9' {
				nextDigit = '1'
			}

		}

	}

	if fractionalDigits < 1 {

		return string(nativeNumStr)
	}

	nativeNumStr = append(nativeNumStr, '.')

	baseFracDigits := fractionalDigits - 1

	nextDigit = '1'

	for i := 0; i < baseFracDigits; i++ {

		nativeNumStr = append(nativeNumStr, nextDigit)

		nextDigit += 1

		if nextDigit > '9' {
			nextDigit = '1'
		}

	}

	// Last Fractional digit is always 5
	nativeNumStr = append(nativeNumStr, '5')

	return string(nativeNumStr)
}
