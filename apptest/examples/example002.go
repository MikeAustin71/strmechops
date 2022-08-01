package examples

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"github.com/MikeAustin71/strmechops/strmech"
	"io"
	"strings"
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

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainTest02.ExtractNumberRunes01()",
		"")

	var outputStr string

	strBuilder := strings.Builder{}

	sMech := strmech.StrMech{}

	// Total available Length of Output Line
	const maxLineLen = 78

	txtFormatCol := strmech.TextFormatterCollection{}
	var err error

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

		outputStr =
			fmt.Sprintf("%v\n"+
				"Error: txtFormatCol.SetStdFormatParamsLine1Col()\n"+
				"Error=\n%v\n",
				ePrefix.String(),
				err.Error())

		fmt.Printf(outputStr + "\n")

		return
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

	// Title Line 2
	err = txtFormatCol.AddLine1Col(
		"Testing StrMech.ExtractNumberRunes()",
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

	// Title Line  3 Date/Time

	err = txtFormatCol.AddLine1Col(
		strmech.TextInputParamFieldDateTimeDto{
			FieldDateTime:       time.Now(),
			FieldDateTimeFormat: "Monday 2006-01-02 15:04:05.000000000 -0700 MST",
		},
		ePrefix.XCpy(
			"Top-Title Line 4"))

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

	numberStr := " 1.234 "

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
