package examples

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"github.com/MikeAustin71/strmechops/strmech"
	"strings"
	"time"
)

type MainTest02 struct {
	input string
}

func (mTest02 MainTest02) ExtractNumberRunes01() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainTest02.ExtractNumberRunes01()",
		"")

	var outputStr string

	strBuilder := strings.Builder{}

	strBuilder.Grow(1024)

	sMech := strmech.StrMech{}

	// Total available Length of Output Line
	const maxFieldLen = 70

	txtBuilder := strmech.TextStrBuilder{}

	// Leading Title Marquee
	var fmtrs []strmech.TextFormatterDto

	// Blank Line
	txtFmt := strmech.TextFormatterDto{}
	txtFmt.FormatType = strmech.TxtFieldType.BlankLine()
	txtFmt.BlankLine.NumOfBlankLines = 1
	fmtrs = append(fmtrs, txtFmt)

	// Filler =======
	// Marquee Top
	txtFmt = strmech.TextFormatterDto{}
	txtFmt.FormatType = strmech.TxtFieldType.Filler()
	txtFmt.Filler.LeftMarginStr = " "
	txtFmt.Filler.FillerCharacters = "="
	txtFmt.Filler.FillerCharsRepeatCount = maxFieldLen - 2
	txtFmt.Filler.RightMarginStr = " "
	txtFmt.Filler.LineTerminator = "\n"
	fmtrs = append(fmtrs, txtFmt)

	// Title Line 1
	txtFmt = strmech.TextFormatterDto{}
	txtFmt.FormatType = strmech.TxtFieldType.Label()
	txtFmt.Label.LeftMarginStr = ""
	txtFmt.Label.FieldText = ePrefix.String()
	txtFmt.Label.FieldLength = maxFieldLen
	txtFmt.Label.FieldJustify = strmech.TxtJustify.Center()
	txtFmt.Label.RightMarginStr = ""
	txtFmt.Label.LineTerminator = "\n"
	fmtrs = append(fmtrs, txtFmt)

	// Title Line 2
	txtFmt = strmech.TextFormatterDto{}
	txtFmt.FormatType = strmech.TxtFieldType.Label()
	txtFmt.Label.LeftMarginStr = ""
	txtFmt.Label.FieldText = "Testing StrMech.ExtractNumberRunes()"
	txtFmt.Label.FieldLength = maxFieldLen
	txtFmt.Label.FieldJustify = strmech.TxtJustify.Center()
	txtFmt.Label.RightMarginStr = ""
	txtFmt.Label.LineTerminator = "\n"
	fmtrs = append(fmtrs, txtFmt)

	// Title Line  3 Date/Time
	txtFmt = strmech.TextFormatterDto{}
	txtFmt.FormatType = strmech.TxtFieldType.DateTime()
	txtFmt.DateTime.LeftMarginStr = ""
	txtFmt.DateTime.FieldDateTime = time.Now()
	txtFmt.DateTime.FieldLength = maxFieldLen
	txtFmt.DateTime.FieldJustify = strmech.TxtJustify.Center()
	txtFmt.DateTime.FieldDateTimeFormat =
		"Monday 2006-01-02 15:04:05.000000000 -0700 MST"
	txtFmt.DateTime.RightMarginStr = ""
	txtFmt.DateTime.LineTerminator = "\n"
	fmtrs = append(fmtrs, txtFmt)

	// Filler Line '========='
	// Marquee Bottom
	txtFmt = strmech.TextFormatterDto{}
	txtFmt.FormatType = strmech.TxtFieldType.Filler()
	txtFmt.Filler.LeftMarginStr = " "
	txtFmt.Filler.FillerCharacters = "="
	txtFmt.Filler.FillerCharsRepeatCount = maxFieldLen - 2
	txtFmt.Filler.RightMarginStr = " "
	txtFmt.Filler.LineTerminator = "\n"
	fmtrs = append(fmtrs, txtFmt)

	// Trailing Blank Line
	txtFmt = strmech.TextFormatterDto{}
	txtFmt.FormatType = strmech.TxtFieldType.BlankLine()
	txtFmt.BlankLine.NumOfBlankLines = 1
	fmtrs = append(fmtrs, txtFmt)

	var err error
	var strBuilder2 strings.Builder

	strBuilder2,
		err = txtBuilder.BuildTextFormatters(
		fmtrs,
		ePrefix.XCpy(
			"strBuilder<-Marquee Top"))

	if err != nil {

		outputStr =
			fmt.Sprintf("%v\n"+
				"Error: txtBuilder.BuildTextFormatters()\n"+
				"Error=\n%v\n",
				ePrefix.String(),
				err.Error())

		fmt.Printf(outputStr + "\n")

		return
	}

	fmtrs = nil

	strBuilder.WriteString(strBuilder2.String())

	strBuilder2.Reset()

	// End Of Marquee

	fmt.Printf(strBuilder.String() + "\n")

	strBuilder.Reset()

	numberStr := " 123456.789 "

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

	strBuilder.Reset()

	strBuilder2,
		err = searchResults.GetParameterTextListing(
		false,
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

	strBuilder.WriteString(strBuilder2.String())
	strBuilder2.Reset()

	fmt.Printf(strBuilder.String() + "\n")
	strBuilder.Reset()

	strBuilder2,
		err = numStrKernel.GetParameterTextListing(
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

	strBuilder.WriteString(strBuilder2.String())
	strBuilder2.Reset()

	fmt.Printf(strBuilder.String() + "\n")
	strBuilder.Reset()

	// Trailing Title Marquee
	// Top Blank Line
	fmtrs = nil
	txtFmt = strmech.TextFormatterDto{}
	txtFmt.FormatType = strmech.TxtFieldType.BlankLine()
	txtFmt.BlankLine.NumOfBlankLines = 1
	fmtrs = append(fmtrs, txtFmt)

	// Filler =======
	// Marquee Top
	txtFmt = strmech.TextFormatterDto{}
	txtFmt.FormatType = strmech.TxtFieldType.Filler()
	txtFmt.Filler.LeftMarginStr = " "
	txtFmt.Filler.FillerCharacters = "="
	txtFmt.Filler.FillerCharsRepeatCount = maxFieldLen - 2
	txtFmt.Filler.RightMarginStr = " "
	txtFmt.Filler.LineTerminator = "\n"
	fmtrs = append(fmtrs, txtFmt)

	// Title # 1
	txtFmt = strmech.TextFormatterDto{}
	txtFmt.FormatType = strmech.TxtFieldType.Label()
	txtFmt.Label.LeftMarginStr = ""
	txtFmt.Label.FieldText = ePrefix.String()
	txtFmt.Label.FieldLength = maxFieldLen
	txtFmt.Label.FieldJustify = strmech.TxtJustify.Center()
	txtFmt.Label.RightMarginStr = ""
	txtFmt.Label.LineTerminator = "\n"
	fmtrs = append(fmtrs, txtFmt)

	// Title # 2
	txtFmt = strmech.TextFormatterDto{}
	txtFmt.FormatType = strmech.TxtFieldType.Label()
	txtFmt.Label.LeftMarginStr = ""
	txtFmt.Label.FieldText = "SUCCESSFUL COMPLETION"
	txtFmt.Label.FieldLength = maxFieldLen
	txtFmt.Label.FieldJustify = strmech.TxtJustify.Center()
	txtFmt.Label.RightMarginStr = ""
	txtFmt.Label.LineTerminator = "\n"
	fmtrs = append(fmtrs, txtFmt)

	// Filler =======
	// Marquee Bottom
	txtFmt = strmech.TextFormatterDto{}
	txtFmt.FormatType = strmech.TxtFieldType.Filler()
	txtFmt.Filler.LeftMarginStr = " "
	txtFmt.Filler.FillerCharacters = "="
	txtFmt.Filler.FillerCharsRepeatCount = maxFieldLen - 2
	txtFmt.Filler.RightMarginStr = " "
	txtFmt.Filler.LineTerminator = "\n"
	fmtrs = append(fmtrs, txtFmt)

	// Blank Line
	txtFmt = strmech.TextFormatterDto{}
	txtFmt.FormatType = strmech.TxtFieldType.BlankLine()
	txtFmt.BlankLine.NumOfBlankLines = 2
	fmtrs = append(fmtrs, txtFmt)

	strBuilder2,
		err = txtBuilder.BuildTextFormatters(
		fmtrs,
		ePrefix.XCpy(
			"Marquee-Bottom"))
	fmtrs = nil

	strBuilder.WriteString(strBuilder2.String())

	strBuilder2.Reset()

	fmt.Printf(strBuilder.String() + "\n")

	strBuilder.Reset()

	return
}
