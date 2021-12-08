package examples

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"github.com/MikeAustin71/strmechops/strmech"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type MainTest struct {
	input string
}

func (mt MainTest) TextLineSpecSolidLine01() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainTest.TextLineSpecSolidLine01()",
		"")

	txtSolidLine03 := strmech.TextLineSpecSolidLine{}

	_,
		err :=
		txtSolidLine03.CopyOutITextLine(
			ePrefix.XCtx(
				"txtSolidLine03->_"))

	if err == nil {

		errorStr := fmt.Sprintf("%v\n"+
			"Error: txtSolidLine03.CopyOutITextLine()\n"+
			"Expected an error return because 'txtSolidLine03' is empty!\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		fmt.Printf(errorStr + "\n")

	}

	return

}

func (mt MainTest) TextLineSpecPlainText01() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainTest.TextLineSpecBlkLines01()",
		"")

	leftMarginChars := "  "
	rightMarginChars := "  "
	textString := "The cow jumped over the moon!"
	newLineChars := "\n"
	outStr := ""

	_,
		err := strmech.TextLineSpecPlainText{}.NewPtrPlainTextStrings(
		leftMarginChars,
		rightMarginChars,
		textString,
		newLineChars,
		false,
		strmech.StrMech{})

	if err != nil {
		outStr = fmt.Sprintf("%v\n"+
			"Error returned from: NewPtrPlainTextStrings()\n"+
			"%v\n",
			ePrefix.String(),
			err.Error())
	} else {
		outStr = fmt.Sprintf("%v\n"+
			"TEST FAILED!! No Error was returned from\n"+
			"NewPtrPlainTextStrings()\n",
			ePrefix.String())
	}

	fmt.Printf(outStr)

	return
}

func (mt MainTest) TextLineSpecBlkLines01() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainTest.TextLineSpecBlkLines01()",
		"")

	var errText string

	blankLinesAlpha,
		err := strmech.TextLineSpecBlankLines{}.NewDefaultBlankLines(
		3,
		ePrefix.XCtx(
			"blankLinesAlpha"))

	if err != nil {
		errText = fmt.Sprintf("%v\n",
			err.Error())

		fmt.Println(errText)

		return
	}

	err = blankLinesAlpha.IsValidInstanceError(
		ePrefix.XCtx(
			"blankLinesAlpha"))

	if err != nil {
		errText = fmt.Sprintf("%v\n",
			err.Error())

		fmt.Println(errText)

		return
	}

	blankLinesBravo := strmech.TextLineSpecBlankLines{}

	err =
		blankLinesBravo.CopyIn(
			&blankLinesAlpha,
			ePrefix.XCtx(
				"blankLinesAlpha->blankLinesBravo"))

	if err != nil {
		errText = fmt.Sprintf("%v\n",
			err.Error())

		fmt.Println(errText)

		return
	}

	err = blankLinesBravo.IsValidInstanceError(
		ePrefix.XCtx(
			"blankLinesBravo"))

	if err != nil {
		errText = fmt.Sprintf("%v\n",
			err.Error())

		fmt.Println(errText)

		return
	}

	errText = fmt.Sprintf("%v\nSuccessful Completion\n",
		ePrefix.XCtxEmpty().String())

	fmt.Println(errText)

	return
}

func (mt MainTest) DateTimeTextField01() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainTest.DateTimeTextField01()",
		"")

	timeZoneName := "America/Chicago"

	tzLocPtr, err := time.LoadLocation(timeZoneName)

	var errTxt string

	if err != nil {
		errTxt = fmt.Sprintf("%v\n"+
			"Error returned by time.LoadLocation(timeZoneName)\n"+
			"timeZoneName='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			timeZoneName,
			err.Error())

		fmt.Println(errTxt)

		return
	}

	dateTime := time.Date(
		2021,
		time.Month(10),
		21,
		14,
		19,
		3,
		0,
		tzLocPtr)

	dateTimeFormat :=
		"2006-01-02 15:04:05.000000000 -0700 MST"

	fieldLen := 39 + 4

	textJustification := strmech.TxtJustify.Center()

	var txtFieldDateTimeOne strmech.TextFieldSpecDateTime

	txtFieldDateTimeOne,
		err = strmech.TextFieldSpecDateTime{}.NewDateTimeField(
		dateTime,
		fieldLen,
		dateTimeFormat,
		textJustification,
		ePrefix.XCtx("txtFieldDateTimeOne"))

	if err != nil {
		errTxt = fmt.Sprintf("%v\n"+
			"Error returned by TextFieldSpecDateTime{}.NewDateTimeField()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		fmt.Println(errTxt)

		return
	}

	var fmtText, convertedFmtText string

	fmtText,
		err = txtFieldDateTimeOne.GetFormattedText(
		ePrefix)

	if err != nil {
		errTxt = fmt.Sprintf("%v\n"+
			"Error returned by txtFieldDateTimeOne.GetFormattedText()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		fmt.Println(errTxt)

		return
	}

	fmt.Printf("Formatted string = \n"+
		"          \"%v\"\n",
		fmtText)
	fmt.Println()

	sMech := strmech.StrMech{}

	convertedFmtText = sMech.ConvertNonPrintableChars(
		[]rune("\""+fmtText+"\""),
		false)

	fmt.Printf("Printable string = \n"+
		"          %v\n",
		convertedFmtText)

	return
}

func (mt MainTest) DateTimeTextField02() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainTest.DateTimeTextField02()",
		"")

	timeZoneName := "America/Los_Angeles"

	tzLocPtr, err := time.LoadLocation(timeZoneName)

	var errTxt string

	if err != nil {
		errTxt = fmt.Sprintf("%v\n"+
			"Error returned by time.LoadLocation(timeZoneName)\n"+
			"timeZoneName='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			timeZoneName,
			err.Error())

		fmt.Println(errTxt)

		return
	}

	dateTime := time.Date(
		2021,
		time.Month(10),
		10,
		20,
		13,
		34,
		0,
		tzLocPtr)

	dateTimeFormat :=
		"2006-01-02 15:04:05.000000000 -0700 MST"

	fieldLen := 39 + 4

	textJustification := strmech.TxtJustify.Center()

	var txtFieldDateTimeOne strmech.TextFieldSpecDateTime

	txtFieldDateTimeOne,
		err = strmech.TextFieldSpecDateTime{}.NewDateTimeField(
		dateTime,
		fieldLen,
		dateTimeFormat,
		textJustification,
		ePrefix.XCtx("txtFieldDateTimeOne"))

	if err != nil {
		errTxt = fmt.Sprintf("%v\n"+
			"Error returned by TextFieldSpecDateTime{}.NewDateTimeField()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		fmt.Println(errTxt)

		return
	}

	var fmtText, convertedFmtText string

	fmtText,
		err = txtFieldDateTimeOne.GetFormattedText(
		ePrefix)

	if err != nil {
		errTxt = fmt.Sprintf("%v\n"+
			"Error returned by txtFieldDateTimeOne.GetFormattedText()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		fmt.Println(errTxt)

		return
	}

	fmt.Printf("Formatted string = \n"+
		"          \"%v\"\n",
		fmtText)
	fmt.Println()

	sMech := strmech.StrMech{}

	convertedFmtText = sMech.ConvertNonPrintableChars(
		[]rune("\""+fmtText+"\""),
		false)

	fmt.Printf("Printable string = \n"+
		"          %v\n",
		convertedFmtText)

	return
}

func (mt MainTest) DateTimeTextField03() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainTest.DateTimeTextField03()",
		"")

	timeZoneName := "America/Chicago"

	tzLocPtr, err := time.LoadLocation(timeZoneName)

	var errTxt string

	if err != nil {
		errTxt = fmt.Sprintf("%v\n"+
			"Error returned by time.LoadLocation(timeZoneName)\n"+
			"timeZoneName='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			timeZoneName,
			err.Error())

		fmt.Println(errTxt)

		return
	}

	dateTime := time.Date(
		2021,
		time.Month(10),
		21,
		14,
		19,
		3,
		0,
		tzLocPtr)

	dateTimeFormat :=
		"Monday January 2, 2006 15:04:05.000000000 -0700 MST"

	fieldLen := 52 + 6

	textJustification := strmech.TxtJustify.Center()

	var txtFieldDateTimeOne strmech.TextFieldSpecDateTime

	txtFieldDateTimeOne,
		err = strmech.TextFieldSpecDateTime{}.NewDateTimeField(
		dateTime,
		fieldLen,
		dateTimeFormat,
		textJustification,
		ePrefix.XCtx("txtFieldDateTimeOne"))

	if err != nil {
		errTxt = fmt.Sprintf("%v\n"+
			"Error returned by TextFieldSpecDateTime{}.NewDateTimeField()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		fmt.Println(errTxt)

		return
	}

	var fmtText, convertedFmtText string

	fmtText,
		err = txtFieldDateTimeOne.GetFormattedText(
		ePrefix)

	if err != nil {
		errTxt = fmt.Sprintf("%v\n"+
			"Error returned by txtFieldDateTimeOne.GetFormattedText()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		fmt.Println(errTxt)

		return
	}

	fmt.Printf("Formatted string = \n"+
		"          \"%v\"\n",
		fmtText)
	fmt.Println()

	sMech := strmech.StrMech{}

	convertedFmtText = sMech.ConvertNonPrintableChars(
		[]rune("\""+fmtText+"\""),
		false)

	fmt.Printf("Printable string = \n"+
		"          %v\n",
		convertedFmtText)

	return
}

func (mt MainTest) TimerEventText01() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainTest.TimerEventText01()",
		"")

	startTime := time.Now()

	elapsedTime := (time.Microsecond * 5) + 999
	endTime := startTime.Add(elapsedTime)

	fmt.Println()
	fmt.Println("---------------------------------")
	fmt.Println(ePrefix.String())
	fmt.Println("---------------------------------")
	fmt.Println()

	var errTxt string

	timerLines,
		err := strmech.TextLineSpecTimerLines{}.NewDefaultFullTimerEvent(
		startTime,
		endTime,
		ePrefix)

	if err != nil {
		errTxt = fmt.Sprintf(
			"%v\n", err.Error())
		fmt.Println(errTxt)
		return
	}

	err = timerLines.IsValidInstanceError(
		ePrefix)

	if err != nil {
		errTxt = fmt.Sprintf(
			"%v\n", err.Error())
		fmt.Println(errTxt)
		return
	}

	fmt.Println("Time Lines Output")
	fmt.Println()
	fmt.Printf(timerLines.String())

}

func (mt MainTest) IntSeparateNumRunes01() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainTest.IntSeparateNumRunes01()",
		"")

	fmt.Println()
	fmt.Println("---------------------------------")
	fmt.Println(ePrefix.String())
	fmt.Println("---------------------------------")
	fmt.Println()

	var outStr string

	intSepDto,
		err := strmech.IntegerSeparatorDto{}.NewDetail(
		",",
		[]uint{3},
		false,
		ePrefix)

	if err != nil {
		outStr = fmt.Sprintf(
			"%v\n\n\n", err.Error())
		fmt.Printf(outStr)
		return
	}

	pureNumStr := []rune("123456789012")
	expectedNStr := "123,456,789,012"
	sMech := strmech.StrMech{}

	var numStrWithIntSeps []rune
	var startTime, endTime time.Time

	startTime = time.Now()

	numStrWithIntSeps,
		err = sMech.IntSeparateNumStr(
		&intSepDto,
		pureNumStr,
		ePrefix)

	endTime = time.Now()

	if err != nil {
		outStr = fmt.Sprintf(
			"%v\n\n\n", err.Error())
		fmt.Printf(outStr)
		return
	}

	var timerLines *strmech.TextLineSpecTimerLines

	timerLines,
		err = strmech.TextLineSpecTimerLines{}.NewDefaultFullTimerEvent(
		startTime,
		endTime,
		ePrefix)

	if err != nil {
		outStr = fmt.Sprintf(
			"%v\n\n\n", err.Error())
		fmt.Printf(outStr)
		return
	}

	actualNumStr := string(numStrWithIntSeps)

	fmt.Printf("------------------------------\n")
	fmt.Printf("Expected Num Str: '%v'\n",
		expectedNStr)
	fmt.Printf("  Actual Num Str: '%v'\n",
		actualNumStr)
	fmt.Printf("------------------------------\n")
	fmt.Printf("Length Expected Num Str: %v\n",
		len(expectedNStr))
	fmt.Printf("  Length Actual Num Str: %v\n",
		len(actualNumStr))
	fmt.Printf("------------------------------\n")
	// Decimal 48 = 0
	fmt.Printf("Actual Rune Array: '%v'\n",
		numStrWithIntSeps)

	fmt.Printf("------------------------------\n\n")
	fmt.Printf(timerLines.String())
	fmt.Printf("------------------------------\n\n")

	if actualNumStr == expectedNStr {
		fmt.Printf("SUCCESSFUL COMPLETION!!\n\n")
	} else {
		fmt.Printf("ERROR Expected vs Actual\n" +
			"DO NOT MATCH!\n\n")
	}

	fmt.Printf("\nEND of %v\n\n",
		ePrefix.String())

}

func (mt MainTest) IntSeparateNumRunes02() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainTest.IntSeparateNumRunes02()",
		"")

	fmt.Println()
	fmt.Println("---------------------------------")
	fmt.Println(ePrefix.String())
	fmt.Println("---------------------------------")
	fmt.Println()

	var outStr string

	intSepDto,
		err := strmech.IntegerSeparatorDto{}.NewDetail(
		",",
		[]uint{3},
		false,
		ePrefix)

	if err != nil {
		outStr = fmt.Sprintf(
			"%v\n\n\n", err.Error())
		fmt.Printf(outStr)
		return
	}

	pureNumStr := []rune("123456789012345")
	expectedNStr := "123,456,789,012,345"
	sMech := strmech.StrMech{}

	var numStrWithIntSeps []rune

	numStrWithIntSeps,
		err = sMech.IntSeparateNumStr(
		&intSepDto,
		pureNumStr,
		ePrefix)

	if err != nil {
		outStr = fmt.Sprintf(
			"%v\n\n\n", err.Error())
		fmt.Printf(outStr)
		return
	}

	actualNumStr := string(numStrWithIntSeps)

	fmt.Printf("------------------------------\n")
	fmt.Printf("Expected Num Str: '%v'\n",
		expectedNStr)
	fmt.Printf("  Actual Num Str: '%v'\n",
		actualNumStr)
	fmt.Printf("------------------------------\n")
	fmt.Printf("Length Expected Num Str: %v\n",
		len(expectedNStr))
	fmt.Printf("  Length Actual Num Str: %v\n",
		len(actualNumStr))
	fmt.Printf("------------------------------\n")
	// Decimal 48 = 0
	fmt.Printf("Actual Rune Array: '%v'\n",
		numStrWithIntSeps)
	fmt.Printf("------------------------------\n\n")

	if actualNumStr == expectedNStr {
		fmt.Printf("SUCCESSFUL COMPLETION!!\n\n")
	} else {
		fmt.Printf("ERROR Expected vs Actual\n" +
			"DO NOT MATCH!\n\n")
	}

	fmt.Printf("\nEND of %v\n\n",
		ePrefix.String())

}

func (mt MainTest) IntSeparateNumRunes03() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainTest.IntSeparateNumRunes03()",
		"")

	fmt.Println()
	fmt.Println("---------------------------------")
	fmt.Println(ePrefix.String())
	fmt.Println("---------------------------------")
	fmt.Println()

	var outStr string

	intSepDto,
		err := strmech.IntegerSeparatorDto{}.NewDetail(
		",",
		[]uint{3, 2},
		false,
		ePrefix)

	if err != nil {
		outStr = fmt.Sprintf(
			"%v\n\n\n", err.Error())
		fmt.Printf(outStr)
		return
	}

	pureNumStr := []rune("6789000000000000")
	expectedNStr := "6,78,90,00,00,00,00,000"
	sMech := strmech.StrMech{}

	var numStrWithIntSeps []rune

	numStrWithIntSeps,
		err = sMech.IntSeparateNumStr(
		&intSepDto,
		pureNumStr,
		ePrefix)

	if err != nil {
		outStr = fmt.Sprintf(
			"%v\n\n\n", err.Error())
		fmt.Printf(outStr)
		return
	}

	actualNumStr := string(numStrWithIntSeps)

	fmt.Printf("------------------------------\n")
	fmt.Printf("Expected Num Str: '%v'\n",
		expectedNStr)
	fmt.Printf("  Actual Num Str: '%v'\n",
		actualNumStr)
	fmt.Printf("------------------------------\n")
	fmt.Printf("Length Expected Num Str: %v\n",
		len(expectedNStr))
	fmt.Printf("  Length Actual Num Str: %v\n",
		len(actualNumStr))
	fmt.Printf("------------------------------\n")
	// Decimal 48 = 0
	fmt.Printf("Actual Rune Array: '%v'\n",
		numStrWithIntSeps)
	fmt.Printf("------------------------------\n\n")

	if actualNumStr == expectedNStr {
		fmt.Printf("SUCCESSFUL COMPLETION!!\n\n")
	} else {
		fmt.Printf("ERROR Expected vs Actual\n" +
			"DO NOT MATCH!\n\n")
	}

	fmt.Printf("\nEND of %v\n\n",
		ePrefix.String())

}

func (mt MainTest) IntSeparateNumRunes04() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainTest.IntSeparateNumRunes04()",
		"")

	var outStr string
	var lineLen = 60

	txtLineBreak,
		err := strmech.TextFieldSpecFiller{}.
		NewTextFiller(
			"-",
			lineLen,
			ePrefix)

	if err != nil {
		outStr = fmt.Sprintf("%v\n",
			err.Error())
		fmt.Println(outStr)
		return
	}

	var txtFuncName, txtLabel *strmech.TextFieldSpecLabel

	txtFuncName,
		err = strmech.TextFieldSpecLabel{}.
		NewPtrTextLabel(
			ePrefix.String(),
			lineLen,
			strmech.TxtJustify.Center(),
			ePrefix.XCtx("ePrefix.String()"))

	if err != nil {
		outStr = fmt.Sprintf("%v\n",
			err.Error())
		fmt.Println(outStr)
		return
	}

	txtLabel,
		err = strmech.TextFieldSpecLabel{}.
		NewPtrTextLabel(
			"Example Of restartIntGroupingSequence",
			lineLen,
			strmech.TxtJustify.Center(),
			ePrefix.XCtx("txtLabel restartIntGroupingSequence"))

	if err != nil {
		outStr = fmt.Sprintf("%v\n",
			err.Error())
		fmt.Println(outStr)
		return
	}

	var str string

	fmt.Println()

	str,
		_ = txtLineBreak.GetFormattedText(ePrefix.XCtxEmpty())

	fmt.Println(str)
	fmt.Println(txtFuncName.String())
	fmt.Println(txtLabel.String())
	fmt.Println(txtLineBreak.GetFormattedText(nil))
	fmt.Println()

	var intSepDto strmech.IntegerSeparatorDto

	intSepDto,
		err = strmech.IntegerSeparatorDto{}.NewDetail(
		",",
		[]uint{3, 2},
		true,
		ePrefix)

	if err != nil {
		outStr = fmt.Sprintf(
			"%v\n\n\n", err.Error())
		fmt.Printf(outStr)
		return
	}

	pureNumStr := []rune("6789000000000000")
	expectedNStr := "6,78,900,00,000,00,000"
	sMech := strmech.StrMech{}

	var numStrWithIntSeps []rune

	numStrWithIntSeps,
		err = sMech.IntSeparateNumStr(
		&intSepDto,
		pureNumStr,
		ePrefix)

	if err != nil {
		outStr = fmt.Sprintf(
			"%v\n\n\n", err.Error())
		fmt.Printf(outStr)
		return
	}

	actualNumStr := string(numStrWithIntSeps)

	var lineBreak string

	lineBreak,
		err = txtLineBreak.GetFormattedText(
		ePrefix.XCtx("LineBreak"))

	if err != nil {
		outStr = fmt.Sprintf(
			"%v\n\n\n", err.Error())
		fmt.Printf(outStr)
		return
	}

	ePrefix.XCtxEmpty()

	fmt.Println(lineBreak)
	fmt.Printf("Expected Num Str: '%v'\n",
		expectedNStr)
	fmt.Printf("  Actual Num Str: '%v'\n",
		actualNumStr)
	fmt.Println(lineBreak)
	fmt.Printf("Length Expected Num Str: %v\n",
		len(expectedNStr))
	fmt.Printf("  Length Actual Num Str: %v\n",
		len(actualNumStr))
	fmt.Println(lineBreak)
	// Decimal 48 = 0
	fmt.Printf("Actual Rune Array: '%v'\n",
		numStrWithIntSeps)
	fmt.Println(lineBreak)
	fmt.Println()

	if actualNumStr == expectedNStr {
		_ = txtLabel.SetText(
			"SUCCESSFUL COMPLETION!!",
			nil)

		fmt.Println(txtLabel.String())
		fmt.Println()

	} else {

		_ = txtLabel.SetText(
			"ERROR Expected vs Actual",
			nil)

		fmt.Println(txtLabel.String())

		_ = txtLabel.SetText(
			"DO NOT MATCH!",
			nil)

		fmt.Println(txtLabel.String())

	}

	fmt.Printf("%v\n", txtLineBreak)

	_ = txtLabel.SetText("End Of ", nil)

	fmt.Printf("%s\n", txtLabel)
	fmt.Printf("%s\n", txtFuncName)
	fmt.Printf("%s\n", txtLineBreak)

}

//goland:noinspection GrazieInspection
func (mt MainTest) IntSeparateNumRunes05() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainTest.IntSeparateNumRunes05()",
		"")

	fmt.Println()
	fmt.Println("---------------------------------")
	fmt.Println(ePrefix.String())
	fmt.Println("Example Of Chinese Numerals")
	fmt.Println("---------------------------------")
	fmt.Println()

	var outStr string
	// Chinese Numerals have an integer grouping value of four
	// and are formatted like this: "12,3456,7890,2345". In this
	// case the 'intGroupingSequence' value would be set to four
	// uint[]{4}.

	intSepDto,
		err := strmech.IntegerSeparatorDto{}.NewDetail(
		",",
		[]uint{4},
		false,
		ePrefix)

	if err != nil {
		outStr = fmt.Sprintf(
			"%v\n\n\n", err.Error())
		fmt.Printf(outStr)
		return
	}

	pureNumStr := []rune("12345678902345")
	expectedNStr := "12,3456,7890,2345"
	sMech := strmech.StrMech{}

	var numStrWithIntSeps []rune

	numStrWithIntSeps,
		err = sMech.IntSeparateNumStr(
		&intSepDto,
		pureNumStr,
		ePrefix)

	if err != nil {
		outStr = fmt.Sprintf(
			"%v\n\n\n", err.Error())
		fmt.Printf(outStr)
		return
	}

	actualNumStr := string(numStrWithIntSeps)

	fmt.Printf("------------------------------\n")
	fmt.Printf("Expected Num Str: '%v'\n",
		expectedNStr)
	fmt.Printf("  Actual Num Str: '%v'\n",
		actualNumStr)
	fmt.Printf("------------------------------\n")
	fmt.Printf("Length Expected Num Str: %v\n",
		len(expectedNStr))
	fmt.Printf("  Length Actual Num Str: %v\n",
		len(actualNumStr))
	fmt.Printf("------------------------------\n")
	// Decimal 48 = 0
	fmt.Printf("Actual Rune Array: '%v'\n",
		numStrWithIntSeps)
	fmt.Printf("------------------------------\n\n")

	if actualNumStr == expectedNStr {
		fmt.Printf("SUCCESSFUL COMPLETION!!\n\n")
	} else {
		fmt.Printf("ERROR Expected vs Actual\n" +
			"DO NOT MATCH!\n\n")
	}

	fmt.Printf("\nEND of %v\n\n",
		ePrefix.String())

}

func (mt MainTest) CopyRuneArrays01() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainTest.CopyRuneArrays01()",
		"")

	var outputStr string

	fmt.Printf("%v\n\n",
		ePrefix.String())

	targetRuneArray := []rune("Hello World")

	outputStr = fmt.Sprintf("targetRuneArray= '%v'\n"+
		"targetRuneArray length= '%v'\n\n",
		string(targetRuneArray),
		len(targetRuneArray))

	fmt.Printf(outputStr)

	var sourceRuneArray = make([]rune, 0)

	outputStr = fmt.Sprintf("sourceRuneArray= '%v'\n"+
		"sourceRuneArray length= '%v'\n",
		string(sourceRuneArray),
		len(sourceRuneArray))

	fmt.Printf(outputStr)

	sMech := strmech.StrMech{}

	err := sMech.CopyRuneArrays(
		&targetRuneArray,
		&sourceRuneArray,
		true,
		&ePrefix)

	if err != nil {
		outputStr = fmt.Sprintf("%v\n",
			err.Error())

		fmt.Printf(outputStr)

		return
	}

	if targetRuneArray != nil {
		outputStr = fmt.Sprintf("%v - Error\n"+
			"After copy operation, expected that\n"+
			"'targetRuneArray' == nil. Instead,\n"+
			"len(targetRuneArray) = '%v'\n",
			ePrefix.String(),
			len(targetRuneArray))

		fmt.Printf(outputStr)

		return
	}

	areEqual := sMech.EqualRuneArraysNil(
		targetRuneArray,
		sourceRuneArray)

	if areEqual {
		outputStr = fmt.Sprintf("%v - Error\n"+
			"'targetRuneArray' and 'sourceRuneArray' are EQUAL!"+
			"This is wrong. 'targetRuneArray' should be 'nil'"+
			"and sourceRuneArray should be a zero length array!\n",
			ePrefix.String())

		fmt.Printf(outputStr)

		return
	}

	fmt.Printf(fmt.Sprintf("\nSUCCESS!\n"))

	var isNil bool

	if targetRuneArray == nil {
		isNil = true
	}

	outputStr = fmt.Sprintf("targetRuneArray= '%v'\n"+
		"targetRuneArray length= '%v'\n"+
		"targetRuneArray is 'nil'= '%v'\n",
		string(targetRuneArray),
		len(targetRuneArray),
		isNil)

	fmt.Printf(outputStr)

	isNil = false

	if sourceRuneArray == nil {
		isNil = true
	}

	outputStr = fmt.Sprintf("sourceRuneArray= '%v'\n"+
		"sourceRuneArray length= '%v'\n"+
		"sourceRuneArray is 'nil'= '%v'\n",
		string(sourceRuneArray),
		len(sourceRuneArray),
		isNil)

	fmt.Printf(outputStr)

}

func (mt MainTest) NumSignCollection01() {

	ePrefix := "NumSignCollection01"

	nSignCollection := strmech.NumberSignSymbolCollection{}
	var err error

	err =
		nSignCollection.AddNewSymbol(
			"+",
			"",
			false,
			ePrefix)

	if err != nil {
		str := fmt.Sprintf("%v", err.Error())
		fmt.Println(str)
		return
	}

	err =
		nSignCollection.AddNewSymbol(
			"-",
			"",
			true,
			ePrefix)

	if err != nil {
		str := fmt.Sprintf("%v", err.Error())
		fmt.Println(str)
		return
	}

	err =
		nSignCollection.AddNewSymbol(
			"",
			"+",
			false,
			ePrefix)

	if err != nil {
		str := fmt.Sprintf("%v", err.Error())
		fmt.Println(str)
		return
	}

	err =
		nSignCollection.AddNewSymbol(
			"",
			"-",
			true,
			ePrefix)

	if err != nil {
		str := fmt.Sprintf("%v", err.Error())
		fmt.Println(str)
		return
	}

	//                   0123456789
	hostRunes := []rune(" -+1234.56")
	startIndex := 1

	foundLeadingNumSign :=
		nSignCollection.IsLeadingNumSignAtHostIndex(
			hostRunes,
			startIndex)

	if !foundLeadingNumSign {
		str := fmt.Sprintf("%v - Round # 1 Error\n"+
			"Expected foundLeadingNumSign == 'true'\n"+
			"Instead foundLeadingNumSign == 'false'\n",
			ePrefix)

		fmt.Println(str)
		return
	}

	//                     0123456789
	//hostRunes := []rune(" -+1234.56")
	startIndex = 2

	foundLeadingNumSign =
		nSignCollection.IsLeadingNumSignAtHostIndex(
			hostRunes,
			startIndex)

	if !foundLeadingNumSign {
		str := fmt.Sprintf("%v - Round #2 Error\n"+
			"Expected foundLeadingNumSign == 'true'\n"+
			"Instead foundLeadingNumSign == 'false'\n",
			ePrefix)
		fmt.Println(str)

		return
	}

	var collection []strmech.NumberSignSymbolDto

	collection,
		err = nSignCollection.GetCollection(ePrefix)

	if collection[0].GetLeadingNumSignFoundIndex() != 2 {
		str := fmt.Sprintf("%v - Error\n"+
			"Expected nSignCollection.numSignSymbols[0].leadingNumSignFoundIndex == 2\n"+
			"Instead, nSignCollection.numSignSymbols[0].leadingNumSignFoundIndex == '%v'\n",
			ePrefix,
			collection[0].GetLeadingNumSignFoundIndex())

		fmt.Println(str)
		return
	}

	if collection[0].GetLeadingNumSignFoundInNumber() != true {
		str := fmt.Sprintf("%v - Error\n"+
			"Expected nSignCollection.numSignSymbols[0].leadingNumSignFoundInNumber == 'true'\n"+
			"Instead, nSignCollection.numSignSymbols[0].leadingNumSignFoundInNumber == 'false'\n",
			ePrefix)

		fmt.Println(str)
		return
	}

	if collection[0].GetNumSignSymPosition() != strmech.NSignSymPos.Before() {
		str := fmt.Sprintf("%v - Error\n"+
			"Expected nSignCollection.numSignSymbols[0].numSignPosition == NSignSymPos.Before()\n"+
			"Instead, nSignCollection.numSignSymbols[0].numSignPosition == '%v'\n",
			ePrefix,
			collection[0].GetNumSignSymPosition().String())

		fmt.Println(str)
		return
	}

	foundNumberSign,
		nSignSymDto,
		_ :=
		nSignCollection.GetFoundNumberSignSymbol()

	if foundNumberSign != true {
		str := fmt.Sprintf("%v - Error\n"+
			"Result: nSignCollection.GetFoundNumberSignSymbol()\n"+
			"Expected foundNumberSign == 'true'\n"+
			"Instead, foundNumberSign == 'false'\n",
			ePrefix)
		fmt.Println(str)
		return
	}

	actualNumSignPos := nSignSymDto.GetNumSignSymPosition()

	if actualNumSignPos != strmech.NSignSymPos.Before() {
		str := fmt.Sprintf("%v - Error\n"+
			"Expected actualNumSignPos == NSignSymPos.Before()\n"+
			"Instead, actualNumSignPos == '%v'\n",
			ePrefix,
			actualNumSignPos.String())
		fmt.Println(str)
		return
	}

	if nSignSymDto.GetLeadingNumSignFoundInNumber() != true {
		str := fmt.Sprintf("%v - Error\n"+
			"Expected nSignSymDto.GetLeadingNumSignFoundInNumber() == 'true'\n"+
			"Instead, nSignSymDto.GetLeadingNumSignFoundInNumber() == 'false'\n",
			ePrefix)
		fmt.Println(str)
		return
	}

	if nSignSymDto.GetLeadingNumSignFoundIndex() != 2 {
		str := fmt.Sprintf("%v - Error\n"+
			"Expected nSignSymDto.GetLeadingNumSignFoundIndex() == '2'\n"+
			"Instead, nSignSymDto.GetLeadingNumSignFoundIndex() == '%v'\n",
			ePrefix,
			nSignSymDto.GetLeadingNumSignFoundIndex())

		fmt.Println(str)
		return
	}

	fmt.Printf("%v\n"+
		"Successful Completion!\n",
		ePrefix)
}

func (mt MainTest) ExampleExtractDataField01() {

	funcName := "ExampleExtractDataField01()"

	sMech := strmech.StrMech{}

	// line feed (ascii character 12) and carriage return (ascii character 13)
	endOfLineDelimiters := []string{"\n"}
	commentDelimiters := []string{"#"}
	leadingFieldSeparators := []string{"\t", " ", "\r", "\f", "\v"}
	trailingFieldSeparators := []string{"\t", " ", "\r", "\f", "\v"}
	//            0          1         2          3         4          5         6         7
	//            0123456 78901234567890123 456789012345678901 23 456789012345678901234567890
	targetStr := " Zone:\t America/Chicago\t Link: US/Central\t!\n"
	expectedLastGoodIdx := strings.LastIndex(targetStr, "\n")
	expectedLastGoodIdx--
	lenTargetStr := len(targetStr)
	startIdx := 0
	leadingKeyWordDelimiters := []string{"Zone:", "Link:"}
	expectedLeadingKeyWordDelimiter := leadingKeyWordDelimiters[0]
	expectedDataFieldStr := "America/Chicago"
	expectedDataFieldIdx := strings.Index(targetStr, expectedDataFieldStr)
	expectedDataFieldLength := len(expectedDataFieldStr)
	expectedDataFieldTrailingDelimiter := "\t"
	expectedDataFieldTrailingDelimiterType := strmech.DfTrailDelimiter.EndOfField()
	expectedLeadingKeyWordDelimiterIndex := strings.Index(targetStr, expectedLeadingKeyWordDelimiter)
	expectedNextTargetIdx := expectedDataFieldIdx + expectedDataFieldLength

	if expectedNextTargetIdx > expectedLastGoodIdx {
		expectedNextTargetIdx = -1
	}

	var startTime, endTime time.Time

	startTime = time.Now()

	datDto,
		err := sMech.ExtractDataField(
		targetStr,
		leadingKeyWordDelimiters,
		startIdx,
		leadingFieldSeparators,
		trailingFieldSeparators,
		commentDelimiters,
		endOfLineDelimiters,
		funcName)

	endTime = time.Now()

	if err != nil {
		fmt.Printf("Error returned by strops.StrOps{}.ExtractDataField()\n"+
			"targetStr='%v'\tstartIdx='%v'\n"+
			"Error='%v'\n", targetStr, startIdx, err.Error())
		return
	}

	isError := false

	if targetStr != datDto.TargetStr {
		fmt.Printf("ERROR: Expected datDto.TargetStr='%v'.\n"+
			"Instead, datDto.TargetStr='%v'.\n",
			targetStr, datDto.TargetStr)
		isError = true
	}

	if lenTargetStr != datDto.TargetStrLength {
		fmt.Printf("ERROR: Expected datDto.TargetStrLength='%v'.\n"+
			"Instead, datDto.TargetStrLength='%v'.\n",
			lenTargetStr, datDto.TargetStrLength)
		isError = true
	}

	if startIdx != datDto.TargetStrStartIndex {
		fmt.Printf("ERROR: Expected datDto.TargetStrStartIndex='%v'.\n"+
			"Instead, datDto.TargetStrStartIndex='%v'.\n",
			startIdx, datDto.TargetStrStartIndex)
		isError = true
	}

	if expectedLeadingKeyWordDelimiter != datDto.LeadingKeyWordDelimiter {
		fmt.Printf("ERROR: Expected datDto.LeadingKeyWordDelimiter='%v'.\n"+
			"Instead, datDto.LeadingKeyWordDelimiter='%v'.\n",
			expectedLeadingKeyWordDelimiter, datDto.LeadingKeyWordDelimiter)
		isError = true
	}

	if expectedLeadingKeyWordDelimiterIndex != datDto.LeadingKeyWordDelimiterIndex {
		fmt.Printf("ERROR: Expected datDto.LeadingKeyWordDelimiterIndex='%v'.\n"+
			"Instead, datDto.LeadingKeyWordDelimiterIndex='%v'.\n",
			expectedLeadingKeyWordDelimiterIndex, datDto.LeadingKeyWordDelimiterIndex)
		isError = true
	}

	if expectedDataFieldStr != datDto.DataFieldStr {
		fmt.Printf("ERROR: Expected datDto.DataFieldStr='%v'.\n"+
			"Instead, datDto.DataFieldStr='%v'.\n",
			expectedDataFieldStr, datDto.DataFieldStr)
		isError = true
	}

	if expectedDataFieldLength != datDto.DataFieldLength {
		fmt.Printf("ERROR: Expected datDto.DataFieldLength='%v'.\n"+
			"Instead, datDto.DataFieldLength='%v'.\n",
			expectedDataFieldLength, datDto.DataFieldLength)
		isError = true
	}

	if expectedDataFieldIdx != datDto.DataFieldIndex {
		fmt.Printf("ERROR: Expected datDto.DataFieldIndex='%v'.\n"+
			"Instead, datDto.DataFieldIndex='%v'.\n",
			expectedDataFieldIdx, datDto.DataFieldIndex)
		isError = true
	}

	if expectedDataFieldTrailingDelimiter != datDto.DataFieldTrailingDelimiter {
		fmt.Printf("ERROR: Expected datDto.DataFieldTrailingDelimiter='%v'.\n"+
			"Instead, datDto.DataFieldTrailingDelimiter='%v'.\n",
			expectedDataFieldTrailingDelimiter, datDto.DataFieldTrailingDelimiter)
		isError = true
	}

	if expectedDataFieldTrailingDelimiterType != datDto.DataFieldTrailingDelimiterType {
		fmt.Printf("ERROR: Expected datDto.DataFieldTrailingDelimiterType='%v'.\n"+
			"Instead, datDto.DataFieldTrailingDelimiterType='%v'.\n",
			expectedDataFieldTrailingDelimiterType.String(), datDto.DataFieldTrailingDelimiterType.String())
		isError = true
	}

	if expectedLastGoodIdx != datDto.TargetStrLastGoodIndex {
		fmt.Printf("ERROR: Expected datDto.TargetStrLastGoodIndex='%v'.\n"+
			"Instead, datDto.TargetStrLastGoodIndex='%v'.\n",
			expectedLastGoodIdx, datDto.TargetStrLastGoodIndex)
		isError = true
	}

	if expectedNextTargetIdx != datDto.NextTargetStrIndex {
		fmt.Printf("ERROR: Expected datDto.NextTargetStrIndex='%v'.\n"+
			"Instead, datDto.NextTargetStrIndex='%v'.\n",
			expectedNextTargetIdx, datDto.NextTargetStrIndex)
		isError = true
	}

	elapsedTotalNanoSecs, elapsedTime := mt.Timer(startTime, endTime)

	fmt.Println("================================================")
	fmt.Println("           ExampleExtractDataField01            ")
	fmt.Println("================================================")
	if isError {
		fmt.Println("              @@@@ FAILURE @@@@                 ")
	} else {
		fmt.Println("                   SUCCESS!                     ")
	}
	fmt.Println("------------------------------------------------")
	fmt.Println("                    Base Data                   ")
	fmt.Println("------------------------------------------------")
	fmt.Printf("             TargetStr: %v", targetStr)
	fmt.Println("      TargetStr Length: ", lenTargetStr)
	fmt.Println("           Start Index: ", startIdx)
	fmt.Println("   Key Word Delimiters: ", leadingKeyWordDelimiters)
	fmt.Println("Key Word Delimiter Idx: ", expectedLeadingKeyWordDelimiterIndex)
	fmt.Println("------------------------------------------------")
	fmt.Println("                 Expected Results               ")
	fmt.Println("------------------------------------------------")
	fmt.Println("                    Field String: ", expectedDataFieldStr)
	fmt.Println("                Field Str Length: ", expectedDataFieldLength)
	fmt.Println("                     Field Index: ", expectedDataFieldIdx)
	fmt.Println("      Leading Key Word Delimiter: ", expectedLeadingKeyWordDelimiter)
	fmt.Println("Leading Key Word Delimiter Index: ", expectedLeadingKeyWordDelimiterIndex)
	fmt.Println("        Field Trailing Delimiter: ", expectedDataFieldTrailingDelimiter)
	fmt.Println("   Field Trailing Delimiter Type: ", expectedDataFieldTrailingDelimiterType.String())
	fmt.Println("      Target Str Last Good Index: ", expectedLastGoodIdx)
	fmt.Println("               Next Target Index: ", expectedNextTargetIdx)
	fmt.Println("------------------------------------------------")
	fmt.Println("                  Actual Results                ")
	fmt.Println("------------------------------------------------")
	fmt.Println("                  Field String: ", datDto.DataFieldStr)
	fmt.Println("              Field Str Length: ", datDto.DataFieldLength)
	fmt.Println("                   Field Index: ", datDto.DataFieldIndex)
	fmt.Println("      Field Trailing Delimiter: ", datDto.DataFieldTrailingDelimiter)
	fmt.Println(" Field Trailing Delimiter Type: ", datDto.DataFieldTrailingDelimiterType.String())
	fmt.Println("             Comment Delimiter: ", datDto.CommentDelimiter)
	fmt.Println("       Comment Delimiter Index: ", datDto.CommentDelimiterIndex)
	fmt.Println("         End Of Line Delimiter: ", datDto.EndOfLineDelimiter)
	fmt.Println("   End Of Line Delimiter Index: ", datDto.EndOfLineDelimiterIndex)
	fmt.Println("             Next Target Index: ", datDto.NextTargetStrIndex)
	fmt.Println("                 Target String: ", datDto.TargetStr)
	fmt.Println("             Target Str Length: ", datDto.TargetStrLength)
	fmt.Println("           Target Str StartIdx: ", datDto.TargetStrStartIndex)
	fmt.Println("    Target Str Last Good Index: ", datDto.TargetStrLastGoodIndex)
	fmt.Println("    Leading Key Delimiter Word: ", datDto.LeadingKeyWordDelimiter)
	fmt.Println("Leading Key Word Delimiter Idx: ", datDto.LeadingKeyWordDelimiterIndex)
	fmt.Println("------------------------------------------------")
	fmt.Println("                 Elapsed Time                   ")
	fmt.Println("------------------------------------------------")
	fmt.Println("     Elapsed Time: ", elapsedTime)
	fmt.Println("Total Nanoseconds: ", elapsedTotalNanoSecs)

}

func (mt MainTest) ExampleExtractDataField02() {

	funcName := "ExampleExtractDataField02()"
	endOfLineDelimiters := []string{"\n"}
	commentDelimiters := []string{"#"}
	leadingFieldSeparators := []string{"\t", " ", "\r", "\f", "\v"}
	trailingFieldSeparators := []string{"\t", " ", "\r", "\f", "\v"}
	targetStr := "\tZone:\tAmerica/Chicago\t\tZone:\tAmerica/New_York\t\tZone:\tAmerica/Los_Angeles\n"
	expectedLastGoodIdx := strings.LastIndex(targetStr, "\n")
	expectedLastGoodIdx--
	lenTargetStr := len(targetStr)
	startIdx := 0
	expectedStartIdx := 46
	leadingKeyWordDelimiters := []string{"Zone:"}
	expectedLeadingKeyWordDelimiter := leadingKeyWordDelimiters[0]
	expectedDataFieldStr := "America/Los_Angeles"
	expectedDataFieldIdx := strings.Index(targetStr, expectedDataFieldStr)
	expectedDataFieldLength := len(expectedDataFieldStr)
	expectedDataFieldTrailingDelimiter := "\n"
	expectedDataFieldTrailingDelimiterType := strmech.DfTrailDelimiter.EndOfLine()
	expectedLeadingKeyWordDelimiterIndex := strings.LastIndex(targetStr, expectedLeadingKeyWordDelimiter)
	expectedNextTargetIdx := expectedDataFieldIdx + expectedDataFieldLength

	if expectedNextTargetIdx > expectedLastGoodIdx {
		expectedNextTargetIdx = -1
	}

	var datDto strmech.DataFieldProfileDto
	var err error
	var startTime, endTime time.Time

	startTime = time.Now()

	sMech := strmech.StrMech{}

	for i := 0; i < 3; i++ {

		datDto,
			err = sMech.ExtractDataField(
			targetStr,
			leadingKeyWordDelimiters,
			startIdx,
			leadingFieldSeparators,
			trailingFieldSeparators,
			commentDelimiters,
			endOfLineDelimiters,
			funcName)

		if err != nil {
			fmt.Printf("Error returned by StrOps{}.ExtractDataField()\n"+
				"Cycle No='%v'\n"+
				"targetStr='%v'\tstartIdx='%v'\n"+
				"Error='%v'\n", i, targetStr, startIdx, err.Error())
			return
		}

		startIdx = datDto.NextTargetStrIndex

	}

	endTime = time.Now()

	isError := false

	if targetStr != datDto.TargetStr {
		fmt.Printf("ERROR: Expected datDto.TargetStr='%v'.\n"+
			"Instead, datDto.TargetStr='%v'.\n",
			targetStr, datDto.TargetStr)
		isError = true
	}

	if lenTargetStr != datDto.TargetStrLength {
		fmt.Printf("ERROR: Expected datDto.TargetStrLength='%v'.\n"+
			"Instead, datDto.TargetStrLength='%v'.\n",
			lenTargetStr, datDto.TargetStrLength)
		isError = true
	}

	if expectedStartIdx != datDto.TargetStrStartIndex {
		fmt.Printf("ERROR: Expected datDto.TargetStrStartIndex='%v'.\n"+
			"Instead, datDto.TargetStrStartIndex='%v'.\n",
			expectedStartIdx, datDto.TargetStrStartIndex)
		isError = true
	}

	if expectedLeadingKeyWordDelimiter != datDto.LeadingKeyWordDelimiter {
		fmt.Printf("ERROR: Expected datDto.LeadingKeyWordDelimiter='%v'.\n"+
			"Instead, datDto.LeadingKeyWordDelimiter='%v'.\n",
			expectedLeadingKeyWordDelimiter, datDto.LeadingKeyWordDelimiter)
		isError = true
	}

	if expectedLeadingKeyWordDelimiterIndex != datDto.LeadingKeyWordDelimiterIndex {
		fmt.Printf("ERROR: Expected datDto.LeadingKeyWordDelimiterIndex='%v'.\n"+
			"Instead, datDto.LeadingKeyWordDelimiterIndex='%v'.\n",
			expectedLeadingKeyWordDelimiterIndex, datDto.LeadingKeyWordDelimiterIndex)
		isError = true
	}

	if expectedDataFieldStr != datDto.DataFieldStr {
		fmt.Printf("ERROR: Expected datDto.DataFieldStr='%v'.\n"+
			"Instead, datDto.DataFieldStr='%v'.\n",
			expectedDataFieldStr, datDto.DataFieldStr)
		isError = true
	}

	if expectedDataFieldLength != datDto.DataFieldLength {
		fmt.Printf("ERROR: Expected datDto.DataFieldLength='%v'.\n"+
			"Instead, datDto.DataFieldLength='%v'.\n",
			expectedDataFieldLength, datDto.DataFieldLength)
		isError = true
	}

	if expectedDataFieldIdx != datDto.DataFieldIndex {
		fmt.Printf("ERROR: Expected datDto.DataFieldIndex='%v'.\n"+
			"Instead, datDto.DataFieldIndex='%v'.\n",
			expectedDataFieldIdx, datDto.DataFieldIndex)
		isError = true
	}

	if expectedDataFieldTrailingDelimiter != datDto.DataFieldTrailingDelimiter {
		fmt.Printf("ERROR: Expected datDto.DataFieldTrailingDelimiter='%v'.\n"+
			"Instead, datDto.DataFieldTrailingDelimiter='%v'.\n",
			expectedDataFieldTrailingDelimiter, datDto.DataFieldTrailingDelimiter)
		isError = true
	}

	if expectedDataFieldTrailingDelimiterType != datDto.DataFieldTrailingDelimiterType {
		fmt.Printf("ERROR: Expected datDto.DataFieldTrailingDelimiterType='%v'.\n"+
			"Instead, datDto.DataFieldTrailingDelimiterType='%v'.\n",
			expectedDataFieldTrailingDelimiterType.String(), datDto.DataFieldTrailingDelimiterType.String())
		isError = true
	}

	if expectedLastGoodIdx != datDto.TargetStrLastGoodIndex {
		fmt.Printf("ERROR: Expected datDto.TargetStrLastGoodIndex='%v'.\n"+
			"Instead, datDto.TargetStrLastGoodIndex='%v'.\n",
			expectedLastGoodIdx, datDto.TargetStrLastGoodIndex)
		isError = true
	}

	if expectedNextTargetIdx != datDto.NextTargetStrIndex {
		fmt.Printf("ERROR: Expected datDto.NextTargetStrIndex='%v'.\n"+
			"Instead, datDto.NextTargetStrIndex='%v'.\n",
			expectedNextTargetIdx, datDto.NextTargetStrIndex)
		isError = true
	}

	elapsedTotalNanoSecs, elapsedTime := mt.Timer(startTime, endTime)

	fmt.Println("================================================")
	fmt.Println("           ExampleExtractDataField02            ")
	fmt.Println("================================================")
	if isError {
		fmt.Println("              @@@@ FAILURE @@@@                 ")
	} else {
		fmt.Println("                   SUCCESS!                     ")
	}
	fmt.Println("------------------------------------------------")
	fmt.Println("                    Base Data                   ")
	fmt.Println("------------------------------------------------")
	fmt.Printf("             TargetStr: %v", targetStr)
	fmt.Println("      TargetStr Length: ", lenTargetStr)
	fmt.Println("           Start Index: ", startIdx)
	fmt.Println("    Key Word Delimiter: ", leadingKeyWordDelimiters)
	fmt.Println("------------------------------------------------")
	fmt.Println("                 Expected Results               ")
	fmt.Println("------------------------------------------------")
	fmt.Println("                   Field String: ", expectedDataFieldStr)
	fmt.Println("               Field Str Length: ", expectedDataFieldLength)
	fmt.Println("                    Field Index: ", expectedDataFieldIdx)
	fmt.Println("       Field Trailing Delimiter: ",
		sMech.ConvertNonPrintableChars([]rune(expectedDataFieldTrailingDelimiter), false))
	fmt.Println("  Field Trailing Delimiter Type: ", expectedDataFieldTrailingDelimiterType.String())
	fmt.Println("      Leading Keyword Delimiter: ", expectedLeadingKeyWordDelimiter)
	fmt.Println("Leading Keyword Delimiter Index: ", expectedLeadingKeyWordDelimiterIndex)
	fmt.Println("     Target Str Last Good Index: ", expectedLastGoodIdx)
	fmt.Println("              Next Target Index: ", expectedNextTargetIdx)
	fmt.Println("------------------------------------------------")
	fmt.Println("                  Actual Results                ")
	fmt.Println("------------------------------------------------")
	fmt.Println("                  Field String: ", datDto.DataFieldStr)
	fmt.Println("              Field Str Length: ", datDto.DataFieldLength)
	fmt.Println("                   Field Index: ", datDto.DataFieldIndex)
	fmt.Println("      Field Trailing Delimiter: ",
		sMech.ConvertNonPrintableChars([]rune(datDto.DataFieldTrailingDelimiter), false))
	fmt.Println(" Field Trailing Delimiter Type: ", datDto.DataFieldTrailingDelimiterType.String())
	fmt.Println("             Comment Delimiter: ",
		sMech.ConvertNonPrintableChars([]rune(datDto.CommentDelimiter), false))
	fmt.Println("       Comment Delimiter Index: ", datDto.CommentDelimiterIndex)
	fmt.Println("         End Of Line Delimiter: ",
		sMech.ConvertNonPrintableChars([]rune(datDto.EndOfLineDelimiter), false))
	fmt.Println("   End Of Line Delimiter Index: ", datDto.EndOfLineDelimiterIndex)
	fmt.Println("             Next Target Index: ", datDto.NextTargetStrIndex)
	fmt.Println("                 Target String: ", datDto.TargetStr)
	fmt.Println("             Target Str Length: ", datDto.TargetStrLength)
	fmt.Println("           Target Str StartIdx: ", datDto.TargetStrStartIndex)
	fmt.Println("    Target Str Last Good Index: ", datDto.TargetStrLastGoodIndex)
	fmt.Println("    Leading Key Delimiter Word: ", datDto.LeadingKeyWordDelimiter)
	fmt.Println("Leading Key Word Delimiter Idx: ", datDto.LeadingKeyWordDelimiterIndex)
	fmt.Println("------------------------------------------------")
	fmt.Println("                 Elapsed Time                   ")
	fmt.Println("------------------------------------------------")
	fmt.Println("     Elapsed Time: ", elapsedTime)
	fmt.Println("Total Nanoseconds: ", elapsedTotalNanoSecs)

}

func (mt MainTest) ExampleExtractDataField03() {

	funcName := "ExampleExtractDataField03()"
	endOfLineDelimiters := []string{"\n"}
	commentDelimiters := []string{"#"}
	leadingFieldSeparators := []string{"\t", " ", "\r", "\f", "\v"}
	trailingFieldSeparators := []string{"\t", " ", "\r", "\f", "\v"}
	targetStr := " America/Chicago\t Good morning America!\n"
	lenTargetStr := len(targetStr)
	expectedLastGoodIdx := strings.LastIndex(targetStr, "\n")
	expectedLastGoodIdx--
	startIdx := 0
	leadingKeyWordDelimiter := []string{"Zone:"}
	expectedLeadingKeyWordDelimiter := ""
	expectedDataFieldStr := ""
	expectedDataFieldIdx := -1
	expectedDataFieldLength := len(expectedDataFieldStr)
	expectedDataFieldTrailingDelimiter := "\n"
	expectedDataFieldTrailingDelimiterType := strmech.DfTrailDelimiter.EndOfLine()
	expectedLeadingKeyWordDelimiterIndex := -1
	expectedNextTargetIdx := expectedDataFieldIdx + expectedDataFieldLength
	expectedNextTargetIdx = -1

	var startTime, endTime time.Time

	startTime = time.Now()

	sMech := strmech.StrMech{}

	datDto,
		err := sMech.ExtractDataField(
		targetStr,
		leadingKeyWordDelimiter,
		startIdx,
		leadingFieldSeparators,
		trailingFieldSeparators,
		commentDelimiters,
		endOfLineDelimiters,
		funcName)

	if err != nil {
		fmt.Printf("Error returned by sMech.ExtractDataField()\n"+
			"targetStr='%v'\tstartIdx='%v'\n"+
			"Error='%v'\n", targetStr, startIdx, err.Error())
		return
	}

	endTime = time.Now()

	isError := false

	if targetStr != datDto.TargetStr {
		fmt.Printf("ERROR: Expected datDto.TargetStr='%v'.\n"+
			"Instead, datDto.TargetStr='%v'.\n",
			targetStr, datDto.TargetStr)
		isError = true
	}

	if lenTargetStr != datDto.TargetStrLength {
		fmt.Printf("ERROR: Expected datDto.TargetStrLength='%v'.\n"+
			"Instead, datDto.TargetStrLength='%v'.\n",
			lenTargetStr, datDto.TargetStrLength)
		isError = true
	}

	if startIdx != datDto.TargetStrStartIndex {
		fmt.Printf("ERROR: Expected datDto.TargetStrStartIndex='%v'.\n"+
			"Instead, datDto.TargetStrStartIndex='%v'.\n",
			startIdx, datDto.TargetStrStartIndex)
		isError = true
	}

	if expectedLeadingKeyWordDelimiter != datDto.LeadingKeyWordDelimiter {
		fmt.Printf("ERROR: Expected datDto.LeadingKeyWordDelimiter='%v'.\n"+
			"Instead, datDto.LeadingKeyWordDelimiter='%v'.\n",
			leadingKeyWordDelimiter, datDto.LeadingKeyWordDelimiter)
		isError = true
	}

	if expectedLeadingKeyWordDelimiterIndex != datDto.LeadingKeyWordDelimiterIndex {
		fmt.Printf("ERROR: Expected datDto.LeadingKeyWordDelimiterIndex='%v'.\n"+
			"Instead, datDto.LeadingKeyWordDelimiterIndex='%v'.\n",
			expectedLeadingKeyWordDelimiterIndex, datDto.LeadingKeyWordDelimiterIndex)
		isError = true
	}

	if expectedDataFieldStr != datDto.DataFieldStr {
		fmt.Printf("ERROR: Expected datDto.DataFieldStr='%v'.\n"+
			"Instead, datDto.DataFieldStr='%v'.\n",
			expectedDataFieldStr, datDto.DataFieldStr)
		isError = true
	}

	if expectedDataFieldLength != datDto.DataFieldLength {
		fmt.Printf("ERROR: Expected datDto.DataFieldLength='%v'.\n"+
			"Instead, datDto.DataFieldLength='%v'.\n",
			expectedDataFieldLength, datDto.DataFieldLength)
		isError = true
	}

	if expectedDataFieldIdx != datDto.DataFieldIndex {
		fmt.Printf("ERROR: Expected datDto.DataFieldIndex='%v'.\n"+
			"Instead, datDto.DataFieldIndex='%v'.\n",
			expectedDataFieldIdx, datDto.DataFieldIndex)
		isError = true
	}

	if expectedDataFieldTrailingDelimiter != datDto.DataFieldTrailingDelimiter {
		fmt.Printf("ERROR: Expected datDto.DataFieldTrailingDelimiter='%v'.\n"+
			"Instead, datDto.DataFieldTrailingDelimiter='%v'.\n",
			expectedDataFieldTrailingDelimiter, datDto.DataFieldTrailingDelimiter)
		isError = true
	}

	if expectedDataFieldTrailingDelimiterType != datDto.DataFieldTrailingDelimiterType {
		fmt.Printf("ERROR: Expected datDto.DataFieldTrailingDelimiterType='%v'.\n"+
			"Instead, datDto.DataFieldTrailingDelimiterType='%v'.\n",
			expectedDataFieldTrailingDelimiterType.String(), datDto.DataFieldTrailingDelimiterType.String())
		isError = true
	}

	if expectedLastGoodIdx != datDto.TargetStrLastGoodIndex {
		fmt.Printf("ERROR: Expected datDto.TargetStrLastGoodIndex='%v'.\n"+
			"Instead, datDto.TargetStrLastGoodIndex='%v'.\n",
			expectedLastGoodIdx, datDto.TargetStrLastGoodIndex)
		isError = true
	}

	if expectedNextTargetIdx != datDto.NextTargetStrIndex {
		fmt.Printf("ERROR: Expected datDto.NextTargetStrIndex='%v'.\n"+
			"Instead, datDto.NextTargetStrIndex='%v'.\n",
			expectedNextTargetIdx, datDto.NextTargetStrIndex)
		isError = true
	}

	elapsedTotalNanoSecs, elapsedTime := mt.Timer(startTime, endTime)

	fmt.Println("================================================")
	fmt.Println("           ExampleExtractDataField03            ")
	fmt.Println("================================================")
	if isError {
		fmt.Println("              @@@@ FAILURE @@@@                 ")
	} else {
		fmt.Println("                   SUCCESS!                     ")
	}
	fmt.Println("------------------------------------------------")
	fmt.Println("                    Base Data                   ")
	fmt.Println("------------------------------------------------")
	fmt.Printf("             TargetStr: %v", targetStr)
	fmt.Println("      TargetStr Length: ", lenTargetStr)
	fmt.Println("           Start Index: ", startIdx)
	fmt.Println("    Key Word Delimiter: ", leadingKeyWordDelimiter)
	fmt.Println("Key Word Delimiter Idx: ", expectedLeadingKeyWordDelimiterIndex)
	fmt.Println("------------------------------------------------")
	fmt.Println("                 Expected Results               ")
	fmt.Println("------------------------------------------------")
	fmt.Println("                   Field String: ", expectedDataFieldStr)
	fmt.Println("               Field Str Length: ", expectedDataFieldLength)
	fmt.Println("                    Field Index: ", expectedDataFieldIdx)
	fmt.Printf("       Field Trailing Delimiter: %v", expectedDataFieldTrailingDelimiter)
	fmt.Println("  Field Trailing Delimiter Type: ", expectedDataFieldTrailingDelimiterType.String())
	fmt.Println("      Leading Keyword Delimiter: ", expectedLeadingKeyWordDelimiter)
	fmt.Println("Leading Keyword Delimiter Index: ", expectedLeadingKeyWordDelimiterIndex)
	fmt.Println("     Target Str Last Good Index: ", expectedLastGoodIdx)
	fmt.Println("              Next Target Index: ", expectedNextTargetIdx)
	fmt.Println("------------------------------------------------")
	fmt.Println("                  Actual Results                ")
	fmt.Println("------------------------------------------------")
	fmt.Println("                  Field String: ", datDto.DataFieldStr)
	fmt.Println("              Field Str Length: ", datDto.DataFieldLength)
	fmt.Println("                   Field Index: ", datDto.DataFieldIndex)
	fmt.Println("      Field Trailing Delimiter: ", datDto.DataFieldTrailingDelimiter)
	fmt.Println(" Field Trailing Delimiter Type: ", datDto.DataFieldTrailingDelimiterType.String())
	fmt.Println("             Next Target Index: ", datDto.NextTargetStrIndex)
	fmt.Println("                 Target String: ", datDto.TargetStr)
	fmt.Println("             Target Str Length: ", datDto.TargetStrLength)
	fmt.Println("           Target Str StartIdx: ", datDto.TargetStrStartIndex)
	fmt.Println("    Target Str Last Good Index: ", datDto.TargetStrLastGoodIndex)
	fmt.Println("    Leading Key Delimiter Word: ", datDto.LeadingKeyWordDelimiter)
	fmt.Println("Leading Key Word Delimiter Idx: ", datDto.LeadingKeyWordDelimiterIndex)
	fmt.Println("------------------------------------------------")
	fmt.Println("                 Elapsed Time                   ")
	fmt.Println("------------------------------------------------")
	fmt.Println("     Elapsed Time: ", elapsedTime)
	fmt.Println("Total Nanoseconds: ", elapsedTotalNanoSecs)

}

func (mt MainTest) ExampleExtractNumStr01() {
	funcName := "ExampleExtractNumStr01()"
	// Etc/GMT-4
	// "Etc/GMT+11"
	// "November 12, 2016 1:6:3pm -(+0000) UTC"
	// "Hello World! Your bank account =$(1,250,364.33).44 What do you think?"
	targetStr := "Hello World! Your bank account =$(1,250,364.33).44 What do you think?"

	expectedNumStr := "$(1,250,364.33)"
	expectedLeadingSignChar := ""
	startIndex := 0
	keepLeadingChars := "$(+-"
	keepInteriorChars := ",."
	keepTrailingChars := ")"

	expectedLeadingSignIndex := -1

	expectedNumStrLen := len(expectedNumStr)
	expectedNumIdx := strings.Index(targetStr, expectedNumStr)
	expectedNextTargetStrIdx := expectedNumIdx + expectedNumStrLen

	if expectedNextTargetStrIdx >= len(targetStr) {
		expectedNextTargetStrIdx = -1
	}
	var startTime, endTime time.Time
	startTime = time.Now()

	nStrDto,
		err := strmech.StrMech{}.Ptr().ExtractNumericDigits(
		targetStr,
		startIndex,
		keepLeadingChars,
		keepInteriorChars,
		keepTrailingChars,
		funcName)

	endTime = time.Now()

	if err != nil {
		fmt.Printf("Error returned by StrOps{}.ExtractNumericDigits(targetStr, 0)\n"+
			"targetStr='%v'\nError='%v'\n", targetStr, err.Error())
		return
	}

	isError := false

	if expectedNumIdx != nStrDto.FirstNumCharIndex {
		fmt.Printf("Expected starting numeric index='%v'\n"+
			"Instead, staring numeric index='%v'\n",
			expectedNumIdx, nStrDto.FirstNumCharIndex)
		isError = true
	}

	if expectedNumStr != nStrDto.NumStr {
		fmt.Printf("Expected number string ='%v'\n"+
			"Instead, number string ='%v'\n",
			expectedNumStr, nStrDto.NumStr)
		isError = true
	}

	if expectedNumStrLen != nStrDto.NumStrLen {
		fmt.Printf("Expected number string length ='%v'\n"+
			"Instead, number string length ='%v'\n",
			expectedNumStrLen, nStrDto.NumStrLen)
		isError = true
	}

	if expectedLeadingSignChar != nStrDto.LeadingSignChar {
		fmt.Printf("Expected leading sign char ='%v'\n"+
			"Instead, leading sign char ='%v'\n",
			expectedLeadingSignChar, nStrDto.LeadingSignChar)
		isError = true
	}

	if expectedLeadingSignIndex != nStrDto.LeadingSignIndex {
		fmt.Printf("Expected leading sign index ='%v'\n"+
			"Instead, leading sign index ='%v'\n",
			expectedLeadingSignIndex, nStrDto.LeadingSignIndex)
		isError = true
	}

	if expectedNextTargetStrIdx != nStrDto.NextTargetStrIndex {
		fmt.Printf("Expected Next TargetStr Char Index ='%v'\n"+
			"Instead, Next TargetStr Char Index ='%v'\n",
			expectedNextTargetStrIdx, nStrDto.NextTargetStrIndex)
		isError = true
	}

	totalNanoSecs,
		elapsedTime := mt.Timer(startTime, endTime)

	fmt.Println("  mainTest.ExampleExtractNumStr01()  ")
	fmt.Println("-------------------------------------")
	if isError {
		fmt.Println("     @@@@@  FAILURE @@@@@@           ")
	} else {
		fmt.Println("          SUCCESS!!!")
	}

	fmt.Println("-------------------------------------")
	fmt.Println("          TargetStr: ", targetStr)
	fmt.Println("           startIdx: ", startIndex)
	fmt.Println("-------------------------------------")
	fmt.Println("           Expected                  ")
	fmt.Println("-------------------------------------")
	fmt.Println("       Number Index: ", expectedNumIdx)
	fmt.Println("         Num Length: ", expectedNumStrLen)
	fmt.Println("  Leading Sign Char: ", expectedLeadingSignChar)
	fmt.Println(" Leading Sign Index: ", expectedLeadingSignIndex)
	fmt.Println("      Number String: ", expectedNumStr)
	fmt.Println(" Next TargetStr Idx: ", expectedNextTargetStrIdx)
	fmt.Println("-------------------------------------")
	fmt.Println("            Results                  ")
	fmt.Println("-------------------------------------")
	fmt.Println("        NumberIndex: ", nStrDto.FirstNumCharIndex)
	fmt.Println("         Num Length: ", nStrDto.NumStrLen)
	fmt.Println("  Leading Sign Char: ", nStrDto.LeadingSignChar)
	fmt.Println(" Leading Sign Index: ", nStrDto.LeadingSignIndex)
	fmt.Println("      Number String: ", nStrDto.NumStr)
	fmt.Println("Target Str Next Idx: ", nStrDto.NextTargetStrIndex)
	fmt.Println("-------------------------------------")
	fmt.Println("       Elapsed Time: ", elapsedTime)
	fmt.Println("  Total NanoSeconds: ", totalNanoSecs)
	fmt.Println("-------------------------------------")

}

func (mt MainTest) ExampleExtractNumStr02() {
	funcName := "ExampleExtractNumStr02()"
	// Etc/GMT-4
	// "Etc/GMT+11"
	// "November 12, 2016 1:6:3pm -(+0000) UTC"
	// "Hello World! Your bank account =$(1,250,364.33).44 What do you think?"
	targetStr := "Hello World! 1234 ?"

	expectedNumStr := "1234"
	expectedLeadingSignChar := ""
	startIndex := 0
	keepLeadingChars := ""
	keepInteriorChars := ""
	keepTrailingChars := ""

	expectedLeadingSignIndex := -1

	expectedNumStrLen := len(expectedNumStr)
	expectedNumIdx := strings.Index(targetStr, expectedNumStr)
	expectedNextTargetStrIdx := expectedNumIdx + expectedNumStrLen

	if expectedNextTargetStrIdx >= len(targetStr) {
		expectedNextTargetStrIdx = -1
	}

	var startTime, endTime time.Time
	startTime = time.Now()

	nStrDto,
		err := strmech.StrMech{}.Ptr().ExtractNumericDigits(
		targetStr,
		startIndex,
		keepLeadingChars,
		keepInteriorChars,
		keepTrailingChars,
		funcName)

	endTime = time.Now()

	if err != nil {
		fmt.Printf("Error returned by StrOps{}.ExtractNumericDigits(targetStr, 0)\n"+
			"targetStr='%v'\nError='%v'\n", targetStr, err.Error())
		return
	}

	totalNanoSeconds,
		elapsedTime := mt.Timer(startTime, endTime)

	isError := false

	if expectedNumIdx != nStrDto.FirstNumCharIndex {
		fmt.Printf("Expected starting numeric index='%v'\n"+
			"Instead, staring numeric index='%v'\n",
			expectedNumIdx, nStrDto.FirstNumCharIndex)
		isError = true
	}

	if expectedNumStr != nStrDto.NumStr {
		fmt.Printf("Expected number string ='%v'\n"+
			"Instead, number string ='%v'\n",
			expectedNumStr, nStrDto.NumStr)
		isError = true
	}

	if expectedNumStrLen != nStrDto.NumStrLen {
		fmt.Printf("Expected number string length ='%v'\n"+
			"Instead, number string length ='%v'\n",
			expectedNumStrLen, nStrDto.NumStrLen)
		isError = true
	}

	if expectedLeadingSignChar != nStrDto.LeadingSignChar {
		fmt.Printf("Expected leading sign char ='%v'\n"+
			"Instead, leading sign char ='%v'\n",
			expectedLeadingSignChar, nStrDto.LeadingSignChar)
		isError = true
	}

	if expectedLeadingSignIndex != nStrDto.LeadingSignIndex {
		fmt.Printf("Expected leading sign index ='%v'\n"+
			"Instead, leading sign index ='%v'\n",
			expectedLeadingSignIndex, nStrDto.LeadingSignIndex)
		isError = true
	}

	if expectedNextTargetStrIdx != nStrDto.NextTargetStrIndex {
		fmt.Printf("Expected Next TargetStr Char Index ='%v'\n"+
			"Instead, Next TargetStr Char Index ='%v'\n",
			expectedNextTargetStrIdx, nStrDto.NextTargetStrIndex)
		isError = true
	}

	fmt.Println("  mainTest.ExampleExtractNumStr02()  ")
	fmt.Println("-------------------------------------")
	if isError {
		fmt.Println("     @@@@@  FAILURE @@@@@@           ")
	} else {
		fmt.Println("          SUCCESS!!!")
	}

	fmt.Println("-------------------------------------")
	fmt.Println("          TargetStr: ", targetStr)
	fmt.Println("           startIdx: ", startIndex)
	fmt.Println("-------------------------------------")
	fmt.Println("           Expected                  ")
	fmt.Println("-------------------------------------")
	fmt.Println("       Number Index: ", expectedNumIdx)
	fmt.Println("         Num Length: ", expectedNumStrLen)
	fmt.Println("  Leading Sign Char: ", expectedLeadingSignChar)
	fmt.Println(" Leading Sign Index: ", expectedLeadingSignIndex)
	fmt.Println("      Number String: ", expectedNumStr)
	fmt.Println(" Next TargetStr Idx: ", expectedNextTargetStrIdx)
	fmt.Println("-------------------------------------")
	fmt.Println("            Results                  ")
	fmt.Println("-------------------------------------")
	fmt.Println("        NumberIndex: ", nStrDto.FirstNumCharIndex)
	fmt.Println("         Num Length: ", nStrDto.NumStrLen)
	fmt.Println("  Leading Sign Char: ", nStrDto.LeadingSignChar)
	fmt.Println(" Leading Sign Index: ", nStrDto.LeadingSignIndex)
	fmt.Println("      Number String: ", nStrDto.NumStr)
	fmt.Println("Target Str Next Idx: ", nStrDto.NextTargetStrIndex)
	fmt.Println("-------------------------------------")
	fmt.Println("       Elapsed Time: ", elapsedTime)
	fmt.Println("  Total NanoSeconds: ", totalNanoSeconds)
	fmt.Println("-------------------------------------")
}

func (mt MainTest) ExampleStripLeadingChars01() {

	badChars := []string{
		" ",
		"/",
		"//",
		"../",
		".",
		"..\\",
		"\\\\\\",
		"..",
		"./",
		"///",
		"..."}
	expectedStr := "SomeString"
	expectedStrLen := len(expectedStr)
	testString := "..........      ./../.\\.\\..\\////   " + expectedStr

	actualString, actualStrLen :=
		strmech.StrMech{}.Ptr().StripLeadingChars(testString, badChars)

	if expectedStr != actualString {
		fmt.Printf("ERROR: Expected result string='%v'\n"+
			"Instead, result string='%v'\n",
			expectedStr, actualString)
		return
	}

	if expectedStrLen != actualStrLen {
		fmt.Printf("ERROR: Expected result string length='%v'\n"+
			"Instead, result string length='%v'\n",
			expectedStrLen, actualStrLen)
	}

}

func (mt MainTest) ExampleSortStrLenHighestToLowest01() {
	badChars := []string{
		"aaaaa",
		"bbbbb",
		"cccccccccc",
		"z",
		"fffffffffff",
		"xx",
		"ddddddddd",
		"eeeeeeeeeee"}

	fmt.Println("Sort by Length Highest To Lowest")
	fmt.Println("          Unordered List")
	fmt.Println("================================")
	fmt.Println()

	for i := 0; i < len(badChars); i++ {
		fmt.Printf("%3d. %v\n", i+1, badChars[i])
	}

	sort.Sort(strmech.SortStrLengthHighestToLowest(badChars))

	fmt.Println()
	fmt.Println("================================")
	fmt.Println("Sort by Length Highest To Lowest")
	fmt.Println("          Ordered List")
	fmt.Println("================================")
	fmt.Println()

	for i := 0; i < len(badChars); i++ {
		fmt.Printf("%3d. %v\n", i+1, badChars[i])
	}

}

func (mt MainTest) ExampleSortStrLenLowestToHighest01() {

	badChars := []string{
		"aaaaa",
		"bbbbb",
		"cccccccccc",
		"z",
		"fffffffffff",
		"xx",
		"ddddddddd",
		"eeeeeeeeeee"}

	fmt.Println("Sort by Length Lowest To Highest")
	fmt.Println("          Unordered List")
	fmt.Println("================================")
	fmt.Println()

	for i := 0; i < len(badChars); i++ {
		fmt.Printf("%3d. %v\n", i+1, badChars[i])
	}

	sort.Sort(strmech.SortStrLengthLowestToHighest(badChars))

	fmt.Println()
	fmt.Println("================================")
	fmt.Println("Sort by Length Lowest To Highest")
	fmt.Println("          Ordered List")
	fmt.Println("================================")
	fmt.Println()

	for i := 0; i < len(badChars); i++ {
		fmt.Printf("%3d. %v\n", i+1, badChars[i])
	}

}

func (mt MainTest) ExampleExpressions01() {
	//AMpm Match \d{1}\s?(?i)[pa][.\s]*(?i)m[.]*
	//PM Match V1 "\\d{1}\\s?(?i)p[.\\s]*(?i)m[.]*"
	//PM Match V2 "\\d{1}\\s{0,4}(?i)p[.]*\\s{0,4}(?i)m[.]*"
	//AM Match V1 "\\d{1}\\s?(?i)a[.\\s]*(?i)m[.]*"
	//AM Match V2 "\\d{1}\\s{0,4}(?i)a[.]*\\s{0,4}(?i)m[.]*"
	regexAMpm := "\\d{1}\\s{0,4}(?i)a[.]*\\s{0,4}(?i)m[.]*"

	samples := []string{
		"12:15 AM",
		"03:25AM",
		"11:19 A M",
		"11:19 A. M.",
		"11:19 A.M.",
		"11:19A.M.",
		"11:19  A.M.",
		"11:19  AM",
		"11:19  AM",
		"11:19  A  M",
		"12:15 am",
		"03:25am",
		"11:19 a m",
		"11:19a m",
		"11:19 a. m.",
		"11:19 a.m.",
		"11:19a.m.",
		"11:19  A  M",
		"11:19  A. M.",
		"11:19  a  m",
		"11:19  a. m.",
		"11:19 m",
		"11:19 a",
		"10:25 PM",
		"02:15PM",
		"10:18 P M",
		"01:19 P. M.",
		"12:19 P.M.",
		"10:19P.M.",
		"10:15 pm",
		"04:25pm",
		"10:19 p m",
		"10:19p m",
		"10:19 p. m.",
		"10:19p.m.",
		"15:35:03",
		"10:19:16 p.m.",
		"10:15 pm -0600 MST",
		"10:15 pm-0600 MST",
		"10:15 pm PST",
		"10:15  pm -0600 MST",
		"10:15 p.m -0600 MST",
		"10:15 pm. -0600 MST",
		"10:15 m -0600 MST",
		"10:15 p -0600 MST",
		"11:19  P.M.",
		"11:19  PM",
		"11:19  PM",
		"11:19  P  M",
		"11:19  P. M.",
		"11:19  p  m",
		"11:19  p. m.",
	}

	lArray := len(samples)
	for i := 0; i < lArray; i++ {
		match, err := mt.FindExpressionExample01(samples[i], regexAMpm)

		if err != nil {
			if err.Error() == "No Match" {
				fmt.Printf("No Match - testStr == %v  regex == %v\n", samples[i], regexAMpm)
				continue
			} else {
				panic(err)
			}
		}

		fmt.Printf("Match! - testStr == %v  regex == %v  match string: %v \n", samples[i], regexAMpm, match)
	}

}

// FindExpressionExample01 - FindExpression_Example_01 - Example function.
func (mt MainTest) FindExpressionExample01(targetStr string, regex string) (string, error) {

	funcName := "FindExpressionExample01()"

	if len(targetStr) < 1 {
		return "", fmt.Errorf("%v\n"+
			"Invalid Target String: %v",
			funcName,
			targetStr)
	}

	// \d{1}\s?(?i)[pa][.\s]*(?i)m[.]*
	r, err := regexp.Compile(regex)

	if err != nil {
		return "", fmt.Errorf(
			"%v\n"+
				"Regex failed to Compile!\n"+
				"regex== %v.\n"+
				"Error: %v",
			funcName, regex, err.Error())
	}

	bTargetStr := []byte(targetStr)

	loc := r.FindIndex(bTargetStr)

	if loc == nil {
		return "", fmt.Errorf("%v\n"+
			"No Match!",
			funcName)
	}

	return string(bTargetStr[loc[0]:loc[1]]), nil

}

func (mt MainTest) TrimMultipleStringsExample01(tStr string, trimChar rune) {

	funcName := "TrimMultipleStringsExample01()"

	su := strmech.StrMech{}

	r, err := su.TrimMultipleChars(
		tStr,
		trimChar,
		funcName)

	if err != nil {
		fmt.Println("Error Return from TrimMultipleChars: ", err.Error())
		return
	}

	fmt.Println("Original String: ", tStr)
	fmt.Println(" Trimmed String: ", r)
	fmt.Println("Original String Length: ", len(tStr))
	fmt.Println(" Trimmed String Length: ", len(r))
	tStr2 := strings.Replace(tStr, " ", "!", -1)
	fmt.Println("Original String TrimChar Locations: ", tStr2)
	r2 := strings.Replace(r, " ", "!", -1)
	fmt.Println(" Trimmed String TrimChar Locations: ", r2)

}

// RegExFindSingleTimeDigitsExample01 - Regex test
func (mt MainTest) RegExFindSingleTimeDigitsExample01() {

	funcName := "RegExFindSingleTimeDigitsExample01()"

	regex := "\\d:\\d:\\d"
	targetStr := "November 12, 2016 1:6:3pm +0000 UTC"

	fmt.Println("targetStr = ", targetStr)
	su := strmech.StrMech{}

	idx := su.FindRegExIndex(targetStr, regex)

	if idx == nil {
		panic(fmt.Errorf(
			"%v\n"+
				"Did not locate Regular Expression,'%v', in 'targetStr', '%v'.\n",
			funcName,
			regex,
			targetStr))
	}

	fmt.Println("Success - Found Regular Expression in targetStr!")
	fmt.Println("idx = ", idx)

	s := []byte(targetStr)

	extract := s[idx[0]:idx[1]]

	sExtract := string(extract)

	fmt.Println("Extracted String: ", sExtract)

	result := strings.Split(sExtract, ":")

	if len(result) == 0 {
		panic(fmt.Errorf("%v\n"+
			"Split returned array of zero length\n",
			funcName))
	}

	fmt.Println("Printing result array:")
	for j := 0; j < len(result); j++ {
		fmt.Println(result[j])
	}

	hrs, _ := strconv.Atoi(result[0])
	min, _ := strconv.Atoi(result[1])
	sec, _ := strconv.Atoi(result[2])

	fmt.Println("Printing Formatted Time String")
	fmt.Printf("%02d:%02d:%02d\n", hrs, min, sec)

	fmt.Println("Reprint with 2-digit seconds")
	fmt.Printf("%02d:%02d:%02d\n", hrs, min, 14)

}

func (mt MainTest) PrintFmtExample01() {

	s1 := fmt.Sprintf("No1: %d  No2: %d", 1, 2)

	fmt.Println(s1)
}

func (mt MainTest) ExampleWrite03() {

	fmt.Println("ExampleWrite03() - Version 2")

	originalStr := "Original base string written to sops1"

	sops1 := strmech.StrMech{}.NewPtr()

	sops1.SetStringData(originalStr)

	sops2 := strmech.StrMech{}.Ptr()

	n, err := io.Copy(sops2, sops1)

	if err != nil {
		fmt.Printf("Error returned by io.Copy(sops2, sops1). Error='%v' \n", err.Error())
		return
	}

	actualStr := sops2.GetStringData()

	if originalStr != actualStr {
		fmt.Printf("Error: Expected string='%v'. Instead, string='%v'. \n",
			originalStr, actualStr)
	}

	fmt.Println()
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("                       ExampleWrite03()")
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("            Original String: ", originalStr)
	fmt.Println("               sops2 String: ", actualStr)
	fmt.Println("  Length of Original String: ", len(originalStr))
	fmt.Println("     Length of sops2 String: ", len(actualStr))
	fmt.Println("                    n Value: ", n)

}

func (mt MainTest) ExampleWrite01() {

	fmt.Println("ExampleWrite01() - Version 2")

	originalStr := "Hello World"

	sops1 := strmech.StrMech{}.Ptr()

	lenOriginalStr := len(originalStr)

	nArray := [4]int{}

	p := make([]byte, 3)

	for i := 0; i < 4; i++ {

		if i == 0 {
			p[0] = 'H'
			p[1] = 'e'
			p[2] = 'l'
		} else if i == 1 {
			p[0] = 'l'
			p[1] = 'o'
			p[2] = ' '
		} else if i == 2 {
			p[0] = 'W'
			p[1] = 'o'
			p[2] = 'r'

		} else if i == 3 {
			p[0] = 'l'
			p[1] = 'd'
			p[2] = byte(0)

		}

		n, err := sops1.Write(p)

		if err != nil {
			fmt.Printf("Error returned by sops1.Write(p). Error='%v' ", err.Error())
			return
		}

		nArray[i] = n

	}

	actualStr := sops1.GetStringData()

	fmt.Println("=========================================================")
	fmt.Println("                  ExampleWrite01()")
	fmt.Println("=========================================================")
	fmt.Println("       Original String: ", originalStr)
	fmt.Println("Original String Length: ", lenOriginalStr)
	fmt.Println("         Actual String: ", actualStr)
	fmt.Println("  Actual String Length: ", len(actualStr))
	fmt.Println("               N Array: ", nArray)

}

func (mt MainTest) ExampleWrite02() {

	fmt.Println("ExampleWrite02() - Version 2")

	originalStr := "Original base string written to sops1"

	sops1 := strmech.StrMech{}.NewPtr()

	lenOriginalStr := len(originalStr)

	p := []byte(originalStr)

	n, err := sops1.Write(p)

	if err != nil {
		fmt.Printf("Error returned by sops1.Write(p). Error='%v' \n", err.Error())
		return
	}

	actualStr := sops1.GetStringData()

	if originalStr != actualStr {
		fmt.Printf("Error: Expected string='%v'. Instead, string='%v'. \n",
			originalStr, actualStr)
	}

	if lenOriginalStr != n {
		fmt.Printf("Error: Expected Length='%v'. Instead, Bytes Written='%v'. \n",
			lenOriginalStr, n)
	}

	fmt.Println()
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("                       ExampleWrite02()")
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("         Original String: ", originalStr)
	fmt.Println("    Actual Output String: ", actualStr)
	fmt.Println("           Length of 'p': ", len(p))
	fmt.Println("                 n Value: ", n)

}

func (mt MainTest) ExampleRead01() {

	fmt.Println("ExampleRead01() - Version 2")

	originalStr := "Original sops1 base string"

	sops1 := strmech.StrMech{}.NewPtr()
	sops1.SetStringData(originalStr)

	p := make([]byte, 5, 15)

	n := 0
	var err error
	err = nil
	cntr := uint64(0)

	b := strings.Builder{}
	b.Grow(len(originalStr) + 150)
	counterArray := make([]uint64, 0, 50)

	for err != io.EOF {

		n, err = sops1.Read(p)

		if err != nil && err != io.EOF {
			fmt.Printf("Error returned by sops1.Read(p). "+
				"Error='%v' \n", err.Error())
			return
		}

		b.Write(p[:n])

		for i := 0; i < len(p); i++ {
			p[i] = byte(0)
		}

		cntr++

	}

	strBuilderStr := b.String()

	fmt.Println("         Original Str: ", originalStr)
	fmt.Println("  Original Str Length: ", len(originalStr))
	fmt.Println("         sops1.StrOut: ", sops1.GetStringData())
	fmt.Println("  sops1.StrOut Length: ", len(sops1.GetStringData()))
	fmt.Println("              Counter: ", cntr)
	fmt.Println("Counter History Array: ", counterArray)
	fmt.Println("       String Builder: ", strBuilderStr)
	fmt.Println("String Builder Length: ", len(strBuilderStr))
	fmt.Println("                    n: ", n)
	fmt.Println("                    p: ", p)
	fmt.Println("**********************************************")

}

func (mt MainTest) ExampleRead02() {
	fmt.Println("ExampleRead02 - Version 2")

	originalStr := "Original sops1 base string"

	sops1 := strmech.StrMech{}.NewPtr()
	sops1.SetStringData(originalStr)

	p := make([]byte, 3, 100)

	_, err := sops1.Read(p)

	if err != nil && err != io.EOF {
		fmt.Printf("Error returned by sops1.Read(p). "+
			"Error='%v' \n", err.Error())
		return
	}

	sops2 := strmech.StrMech{}.NewPtr()
	n, err := sops2.Write(p)

	fmt.Println("        Original Str: ", originalStr)
	fmt.Println(" Original Str Length: ", len(originalStr))
	fmt.Println("        sops1.StrOut: ", sops1.GetStringData())
	fmt.Println(" sops1.StrOut Length: ", len(sops1.GetStringData()))
	fmt.Println("    sops1 Bytes Read: ", n)
	fmt.Println("**********************************************")

}

func (mt MainTest) ExampleIoCopy02() {
	fmt.Println("ExampleIOCopy_02() - Version 2")

	originalStr := "Original sops1 base string"

	sops1 := strmech.StrMech{}.NewPtr()
	sops1.SetStringData(originalStr)
	sops2 := strmech.StrMech{}.Ptr()

	n, err := io.Copy(sops2, sops1)

	if err != nil {
		fmt.Printf("Error returned by io.Copy(sops2, sops1). "+
			"Error='%v' \n", err.Error())
		return
	}

	fmt.Println("           Original Str: ", originalStr)
	fmt.Println("    Original Str Length: ", len(originalStr))
	fmt.Println("       sops1.stringData: ", sops1.GetStringData())
	fmt.Println("       sops2.stringData: ", sops2.GetStringData())
	fmt.Println("sops2.stringData Length: ", len(sops2.GetStringData()))
	fmt.Println("          Bytes Written: ", n)
	fmt.Println("**********************************************")
	fmt.Println("      Copying sops2 To StringData")
	fmt.Println("**********************************************")

	n, err = io.Copy(os.Stdout, sops2)

	if err != nil {
		fmt.Printf("Error returned by io.Copy(os.Stdout, sops2). "+
			"Error='%v' \n", err.Error())
		return
	}

	fmt.Println()
	fmt.Println("New value of n: ", n)
}

func (mt MainTest) ExampleExtractNumRunes01() {
	funcName := "ExampleExtractNumRunes01()"
	rawNumStrRunes := []rune("1234.5678")
	leadingNegativeSignChars := []rune{'-'}
	trailingNegativeSignChars := []rune{0}
	decimalSeparatorChars := []rune{'.'}

	sMech := strmech.StrMech{}

	var startTime, endTime time.Time
	startTime = time.Now()

	intRunes,
		fracRunes,
		numberSign,
		digitsFound,
		err := sMech.ExtractNumberRunes(
		rawNumStrRunes,
		leadingNegativeSignChars,
		trailingNegativeSignChars,
		decimalSeparatorChars,
		funcName)

	endTime = time.Now()

	isError := false
	errStr := ""
	if err != nil {
		isError = true
		errStr = err.Error()
	}

	totalNanoSeconds,
		elapsedTime := mt.Timer(startTime, endTime)

	lineBrk := strings.Repeat("-", 70)

	fmt.Println("  " + funcName)
	fmt.Println(lineBrk)
	if isError {
		fmt.Println("     @@@@@  FAILURE @@@@@@           ")
		fmt.Printf("%v\n\n",
			errStr)
		return
	} else {
		fmt.Println("          SUCCESS!!!")
	}

	fmt.Println(lineBrk)
	fmt.Printf("   NumberStr: %v\n", string(rawNumStrRunes))
	fmt.Printf("    intRunes: %v\n", intRunes)
	fmt.Printf("   intString: %v\n", string(intRunes))
	fmt.Printf("   fracRunes: %v\n", fracRunes)
	fmt.Printf("  fracString: %v\n", string(fracRunes))
	fmt.Printf(" Number Sign: %v\n", numberSign)
	fmt.Printf("Digits Found: %v\n", digitsFound)
	fmt.Println(lineBrk)
	fmt.Println("       Elapsed Time: ", elapsedTime)
	fmt.Println("  Total NanoSeconds: ", totalNanoSeconds)
	fmt.Println(lineBrk)
	fmt.Println()
}

func (mt MainTest) ExampleExtractNumRunes02() {
	funcName := "ExampleExtractNumRunes02()"
	rawNumStrRunes := []rune("-1234.5678")
	leadingNegativeSignChars := []rune{'-'}
	trailingNegativeSignChars := []rune{0}
	decimalSeparatorChars := []rune{'.'}

	sMech := strmech.StrMech{}

	var startTime, endTime time.Time
	startTime = time.Now()

	intRunes,
		fracRunes,
		numberSign,
		digitsFound,
		err := sMech.ExtractNumberRunes(
		rawNumStrRunes,
		leadingNegativeSignChars,
		trailingNegativeSignChars,
		decimalSeparatorChars,
		funcName)

	endTime = time.Now()

	isError := false
	errStr := ""
	if err != nil {
		isError = true
		errStr = err.Error()
	}

	totalNanoSeconds,
		elapsedTime := mt.Timer(startTime, endTime)

	lineBrk := strings.Repeat("-", 70)

	fmt.Println("  " + funcName)
	fmt.Println(lineBrk)
	if isError {
		fmt.Println("     @@@@@  FAILURE @@@@@@           ")
		fmt.Printf("%v\n\n",
			errStr)
		return
	} else {
		fmt.Println("          SUCCESS!!!")
	}

	fmt.Println(lineBrk)
	fmt.Printf("   NumberStr: %v\n", string(rawNumStrRunes))
	fmt.Printf("    intRunes: %v\n", intRunes)
	fmt.Printf("   intString: %v\n", string(intRunes))
	fmt.Printf("   fracRunes: %v\n", fracRunes)
	fmt.Printf("  fracString: %v\n", string(fracRunes))
	fmt.Printf(" Number Sign: %v\n", numberSign)
	fmt.Printf("Digits Found: %v\n", digitsFound)
	fmt.Println(lineBrk)
	fmt.Println("       Elapsed Time: ", elapsedTime)
	fmt.Println("  Total NanoSeconds: ", totalNanoSeconds)
	fmt.Println(lineBrk)
	fmt.Println()
}

func (mt MainTest) ExampleExtractNumRunes03(
	numStrRunes []rune,
	leadingNegativeSignChars []rune,
	trailingNegativeSignChars []rune,
	decimalSeparatorChars []rune,
	errorPrefix interface{}) (err error) {

	funcName := "MainTest.ExampleExtractNumRunes03()"
	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		funcName,
		fmt.Sprintf(
			"       numStrRunes='%s'\n"+
				"  leadingNegSignChars='%v'\n"+
				" trailingNegSignChars='%v'\n"+
				"decimalSeparatorChars='%v'\n",
			string(numStrRunes),
			string(leadingNegativeSignChars),
			trailingNegativeSignChars,
			decimalSeparatorChars))

	if err != nil {
		return err
	}

	sMech := strmech.StrMech{}

	var startTime, endTime time.Time
	startTime = time.Now()

	intRunes,
		fracRunes,
		numberSign,
		digitsFound,
		err := sMech.ExtractNumberRunes(
		numStrRunes,
		leadingNegativeSignChars,
		trailingNegativeSignChars,
		decimalSeparatorChars,
		ePrefix)

	endTime = time.Now()

	isError := false
	errStr := ""
	if err != nil {
		isError = true
		errStr = err.Error()
	}

	totalNanoSeconds,
		elapsedTime := mt.Timer(startTime, endTime)

	lineBrk := strings.Repeat("-", 70)

	fmt.Println("  " + funcName)
	fmt.Println(lineBrk)
	if isError {
		fmt.Println("     @@@@@  FAILURE @@@@@@           ")
		fmt.Printf("%v\n\n",
			errStr)
		return
	} else {
		fmt.Println("          SUCCESS!!!")
	}

	fmt.Println(lineBrk)
	fmt.Printf("   NumberStr: %v\n", string(numStrRunes))
	fmt.Printf("    intRunes: %v\n", intRunes)
	fmt.Printf("   intString: %v\n", string(intRunes))
	fmt.Printf("   fracRunes: %v\n", fracRunes)
	fmt.Printf("  fracString: %v\n", string(fracRunes))
	fmt.Printf(" Number Sign: %v\n", numberSign)
	fmt.Printf("Digits Found: %v\n", digitsFound)
	fmt.Println(lineBrk)
	fmt.Println("       Elapsed Time: ", elapsedTime)
	fmt.Println("  Total NanoSeconds: ", totalNanoSeconds)
	fmt.Println(lineBrk)
	fmt.Println()

	return nil
}

func (mt MainTest) Timer(
	startTime, endTime time.Time) (totalNanoSecs int64, elapsedTime string) {

	// MicroSecondNanoseconds - Number of Nanoseconds in a Microsecond.
	// A MicroSecond is 1/1,000,000 or 1 one-millionth of a second
	MicroSecondNanoseconds := int64(time.Microsecond)

	// MilliSecondNanoseconds - Number of Nanoseconds in a MilliSecond.
	//	 A millisecond is 1/1,000 or 1 one-thousandth of a second
	MilliSecondNanoseconds := int64(time.Millisecond)

	// SecondNanoseconds - Number of Nanoseconds in a Second
	SecondNanoseconds := int64(time.Second)

	// MinuteNanoseconds - Number of Nanoseconds in a minute
	MinuteNanoseconds := int64(time.Minute)

	// HourNanoseconds - Number of Nanoseconds in an hour
	HourNanoseconds := int64(time.Hour)

	t2Dur := endTime.Sub(startTime)

	str := ""

	totalNanoseconds := t2Dur.Nanoseconds()
	var numOfHours, numOfMinutes, numOfSeconds, numOfMilliseconds,
		numOfMicroseconds, numOfNanoseconds int64

	if totalNanoseconds >= HourNanoseconds {
		numOfHours = totalNanoseconds / HourNanoseconds
		totalNanoseconds = totalNanoseconds - (numOfHours * HourNanoseconds)
	}

	if totalNanoseconds >= MinuteNanoseconds {
		numOfMinutes = totalNanoseconds / MinuteNanoseconds
		totalNanoseconds = totalNanoseconds - (numOfMinutes * MinuteNanoseconds)
	}

	if totalNanoseconds >= SecondNanoseconds {
		numOfSeconds = totalNanoseconds / SecondNanoseconds
		totalNanoseconds = totalNanoseconds - (numOfSeconds * SecondNanoseconds)
	}

	if totalNanoseconds >= SecondNanoseconds {
		numOfSeconds = totalNanoseconds / SecondNanoseconds
		totalNanoseconds = totalNanoseconds - (numOfSeconds * SecondNanoseconds)
	}

	if totalNanoseconds >= MilliSecondNanoseconds {
		numOfMilliseconds = totalNanoseconds / MilliSecondNanoseconds
		totalNanoseconds = totalNanoseconds - (numOfMilliseconds * MilliSecondNanoseconds)
	}

	if totalNanoseconds >= MicroSecondNanoseconds {
		numOfMicroseconds = totalNanoseconds / MicroSecondNanoseconds
		totalNanoseconds = totalNanoseconds - (numOfMicroseconds * MicroSecondNanoseconds)
	}

	numOfNanoseconds = totalNanoseconds

	if numOfHours > 0 {

		str += fmt.Sprintf("%v-Hours ", numOfHours)

	}

	if numOfMinutes > 0 {

		str += fmt.Sprintf("%v-Minutes ", numOfMinutes)

	}

	if numOfSeconds > 0 || str != "" {

		str += fmt.Sprintf("%v-Seconds ", numOfSeconds)

	}

	if numOfMilliseconds > 0 || str != "" {

		str += fmt.Sprintf("%v-Milliseconds ", numOfMilliseconds)

	}

	if numOfMicroseconds > 0 || str != "" {

		str += fmt.Sprintf("%v-Microseconds ", numOfMicroseconds)

	}

	str += fmt.Sprintf("%v-Nanoseconds", numOfNanoseconds)

	return t2Dur.Nanoseconds(), str
}
