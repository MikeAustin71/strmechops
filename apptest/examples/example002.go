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

func (MainTest02) BigRat002() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"BigRat002()",
		"")

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Starting Run!\n"+
		"Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

	var origIntDigits, origFracDigits string

	origIntDigits = "123456"
	origFracDigits = ""

	bigRatToFracDigits := 5

	origNumberSign := strmech.NumSignVal.Negative()

	var origNumStr string

	if origNumberSign == strmech.NumSignVal.Negative() {
		origNumStr += "-"
	}

	origNumStr += origIntDigits

	if len(origFracDigits) > 0 {

		origNumStr += "."
		origNumStr += origFracDigits
	}

	numStrKernelBase,
		err := new(strmech.NumberStrKernel).
		NewFromStringDigits(
			origIntDigits,
			origFracDigits,
			origNumberSign,
			ePrefix.XCpy(
				origNumStr))

	if err != nil {
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	var pureNumStr string

	pureNumStr = numStrKernelBase.GetPureNumberStr(
		".",
		true)

	fmt.Printf("%v\n"+
		"Original NumStr    = '%v'\n"+
		"Initial Pure NumStr= '%v'\n\n",
		ePrefix.String(),
		origNumStr,
		pureNumStr)

	var bigRatNum *big.Rat

	bigRatNum,
		err = numStrKernelBase.GetBigRatNum(
		strmech.NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"bigRatNum<-"))

	if err != nil {
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	var bigRatNumStr, bigRatTwoElementStr string

	bigRatNumStr =
		bigRatNum.FloatString(bigRatToFracDigits)

	bigRatTwoElementStr = bigRatNum.RatString()

	bigRatNumerator := bigRatNum.Num()
	bigRatDenominator := bigRatNum.Denom()

	fmt.Printf("%v\n"+
		"bigRatNum to %v-Digits\n"+
		"bigRatNum         = '%v'\n"+
		"bigRatString      = '%v'\n"+
		"bigRatNumerator   = '%v'\n"+
		"bigRatDenominator = '%v'\n\n",
		ePrefix.String(),
		bigRatToFracDigits,
		bigRatNumStr,
		bigRatTwoElementStr,
		bigRatNumerator.Text(10),
		bigRatDenominator.Text(10))

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		"Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

}

func (MainTest02) BigRatHelper01() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"BigRatHelper01()",
		"")

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Starting Run!\n"+
		"Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

	newRat := big.NewRat(1, 3)

	var bigFloatNum *big.Float

	var err error

	bigFloatNum,
		err = new(strmech.MathBigRatHelper).BigRatToBigFloat(
		newRat,
		10,
		ePrefix)

	if err != nil {
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	ratNum := newRat.Num()

	ratDenom := newRat.Denom()

	fmt.Printf("newRat = %v/%v\n",
		ratNum.Text(10),
		ratDenom.Text(10))

	fmt.Printf("bigFloatNum = %v\n",
		bigFloatNum.Text('f', -1))

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		"Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

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
		true,
		"=",
		"=",
		true,
		titles,
		txtFmtParams,
		true,
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
		true,
		"=",
		"=",
		true,
		titles,
		txtFmtParams,
		ePrefix.XCpy(
			"strBuilder<-txtFmtParams"))

	// End Of Marquee

	return err
}
