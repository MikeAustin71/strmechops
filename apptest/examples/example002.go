package examples

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"github.com/MikeAustin71/strmechops/strmech"
	"io"
	"strings"
	"sync"
	"time"
)

type MainTest02 struct {
	input string
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

	numberStr := "1 2 3 4 5678"
	characterSearchLen := 7

	// Total available Length of Output Line
	const maxLineLen = 78

	txtFormatCol := strmech.TextFormatterCollection{}
	var err error

	mt02Nanobot := mainTest02Nanobot{}

	err = mt02Nanobot.SetLeadingMarquee(
		funcName,
		"Testing StrMech.ExtractNumberRunes()",
		maxLineLen,
		&txtFormatCol,
		&ePrefix)

	if err != nil {
		fmt.Println(
			fmt.Sprintf("%v",
				err.Error()))

		return
	}

	txtStrBuilder := strmech.TextStrBuilder{}

	err = txtStrBuilder.BuildText(
		&strBuilder,
		&txtFormatCol,
		ePrefix.XCpy(
			"strBuilder<-txtFormatCol"))

	if err != nil {
		fmt.Println(
			fmt.Sprintf("%v",
				err.Error()))

		return
	}

	txtFormatCol.EmptyFormatterCollection()

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

	strBuilder.Grow(256)

	err = searchResults.GetParameterTextListing(
		&strBuilder,
		true,
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
		funcName,
		"SUCCESSFUL COMPLETION!",
		maxLineLen,
		&txtFormatCol,
		ePrefix.XCpy(
			"Trailing Title Marquee"))

	if err != nil {
		fmt.Println(
			fmt.Sprintf("%v",
				err.Error()))

		return
	}

	err = txtStrBuilder.BuildText(
		&strBuilder,
		&txtFormatCol,
		ePrefix.XCpy(
			"strBuilder<-txtFormatCol"))

	if err != nil {
		fmt.Println(
			fmt.Sprintf("%v",
				err.Error()))

		return
	}

	fmt.Printf(strBuilder.String() + "\n")

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

	numberStr := " -1.2 34 "

	// Total available Length of Output Line
	const maxLineLen = 78

	txtFormatCol := strmech.TextFormatterCollection{}
	var err error

	// Leading Title Marquee

	mt02Nanobot := mainTest02Nanobot{}

	err = mt02Nanobot.SetLeadingMarquee(
		funcName,
		"Testing StrMech.ExtractNumberRunes()",
		maxLineLen,
		&txtFormatCol,
		&ePrefix)

	if err != nil {
		fmt.Println(
			fmt.Sprintf("%v",
				err.Error()))

		return
	}

	// End Of Marquee

	txtStrBuilder := strmech.TextStrBuilder{}

	err = txtStrBuilder.BuildText(
		&strBuilder,
		&txtFormatCol,
		ePrefix.XCpy(
			"strBuilder<-txtFormatCol"))

	if err != nil {
		fmt.Println(
			fmt.Sprintf("%v",
				err.Error()))

		return
	}

	txtFormatCol.EmptyFormatterCollection()

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

	strBuilder.Grow(256)

	err = searchResults.GetParameterTextListing(
		&strBuilder,
		true,
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
		funcName,
		"SUCCESSFUL COMPLETION!",
		maxLineLen,
		&txtFormatCol,
		ePrefix.XCpy(
			"Trailing Title Marquee"))

	if err != nil {
		fmt.Println(
			fmt.Sprintf("%v",
				err.Error()))

		return
	}

	err = txtStrBuilder.BuildText(
		&strBuilder,
		&txtFormatCol,
		ePrefix.XCpy(
			"strBuilder<-txtFormatCol"))

	if err != nil {
		fmt.Println(
			fmt.Sprintf("%v",
				err.Error()))

		return
	}

	fmt.Printf(strBuilder.String() + "\n")

	strBuilder.Reset()

	return
}

type mainTest02Nanobot struct {
	lock *sync.Mutex
}

func (mt02Nanobot *mainTest02Nanobot) SetLeadingMarquee(
	funcName string,
	testTitle string,
	maxLineLen int,
	txtFormatCol *strmech.TextFormatterCollection,
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

	if txtFormatCol == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtFormatCol' is a nil pointer.\n",
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

	// Leading Title Marquee

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
		return err
	}

	// Blank Line
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

	// Title Line 1
	err = txtFormatCol.AddLine1Col(
		funcName,
		ePrefix.XCpy(
			"Title Line 1"))

	if err != nil {
		return err
	}

	// Title Line 2
	err = txtFormatCol.AddLine1Col(
		testTitle,
		ePrefix.XCpy(
			"Title Line 1"))

	if err != nil {
		return err
	}

	// Title Line  3 Date/Time

	err = txtFormatCol.AddLine1Col(
		strmech.TextInputParamFieldDateTimeDto{
			FieldDateTime:       time.Now(),
			FieldDateTimeFormat: "Monday 2006-01-02 15:04:05.000000000 -0700 MST",
		},
		ePrefix.XCpy(
			"Top-Title Line 3"))

	if err != nil {
		return
	}

	// Filler Line '========='
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

	// Trailing Blank Line
	txtFormatCol.AddLineBlank(
		1,
		"")

	// End Of Marquee

	return err
}

func (mt02Nanobot *mainTest02Nanobot) SetTrailingMarquee(
	funcName string,
	completionDesc string,
	maxLineLen int,
	txtFormatCol *strmech.TextFormatterCollection,
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

	if txtFormatCol == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtFormatCol' is a nil pointer.\n",
			ePrefix.String())

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

	// Trailing Title Marquee

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
		return err
	}

	// Blank Line
	txtFormatCol.AddLineBlank(
		1,
		"")

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

	// Title Line 1
	err = txtFormatCol.AddLine1Col(
		funcName,
		ePrefix.XCpy(
			"Bottom-Title Line 1"))

	if err != nil {
		return err
	}

	// Title Line 2
	err = txtFormatCol.AddLine1Col(
		completionDesc,
		ePrefix.XCpy(
			"Bottom-Title Line 2"))

	if err != nil {
		return err
	}

	// Filler Line '========='
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

	// Trailing Blank Line
	txtFormatCol.AddLineBlank(
		1,
		"")

	// End Of Marquee

	return err
}
