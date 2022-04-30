package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"strings"
	"testing"
	"time"
)

func TestTextLineSpecTimerLines_CopyIn_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_CopyIn_000100()",
		"")

	outputStr,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var timerLines02 *TextLineSpecTimerLines

	_,
		timerLines02,
		err = createTestTextLineSpecTimerLines02(
		ePrefix.XCpy(
			"timerLines02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = timerLines02.CopyIn(
		timerLines01,
		ePrefix.XCpy(
			"timerLines02<-timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var actualStr string

	actualStr,
		err = timerLines02.GetFormattedText(
		ePrefix.XCpy(
			"timerLines02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sMech := StrMech{}

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(actualStr),
			true)

	if outputStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"timerLines02.GetFormattedText()"+
			"Expected string DOES NOT match Actual string\n"+
			"Expected string = '%v'\n"+
			"  Actual string = '%v'\n",
			ePrefix.String(),
			outputStr,
			printableActualStr)

		return

	}

	var timerLines03 *TextLineSpecTimerLines

	_,
		timerLines03,
		err = createTestTextLineSpecTimerLines02(
		ePrefix.XCpy(
			"timerLines03"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = timerLines03.CopyIn(
		nil,
		ePrefix.XCpy(
			"timerLines03<-nil"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from timerLines03."+
			"CopyIn()\n"+
			"because 'incomingTimerLines' is 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	err = timerLines03.CopyIn(
		timerLines02,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from timerLines03."+
			"CopyIn()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	timerLines04 := TextLineSpecTimerLines{}

	err = timerLines04.CopyIn(
		timerLines01,
		ePrefix.XCpy(
			"timerLines04"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	return
}

func TestTextLineSpecTimerLines_CopyOut_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_CopyOut_000100()",
		"")

	outputStr,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var timerLines02 TextLineSpecTimerLines

	timerLines02,
		err = timerLines01.CopyOut(
		ePrefix.XCpy(
			"timerLines02<-timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var actualStr string

	actualStr,
		err = timerLines02.GetFormattedText(
		ePrefix.XCpy(
			"timerLines02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sMech := StrMech{}

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(actualStr),
			true)

	if outputStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"timerLines02.GetFormattedText()\n"+
			"Expected string DOES NOT match Actual string\n"+
			"Expected string = '%v'\n"+
			"  Actual string = '%v'\n",
			ePrefix.String(),
			outputStr,
			printableActualStr)

		return
	}

	var timerLines03 *TextLineSpecTimerLines

	_,
		timerLines03,
		err = createTestTextLineSpecTimerLines02(
		ePrefix.XCpy(
			"timerLines03"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	*timerLines03,
		err = timerLines02.CopyOut(
		ePrefix.XCpy(
			"timerLines03<-timerLines02"))

	actualStr,
		err = timerLines03.GetFormattedText(
		ePrefix.XCpy(
			"timerLines03"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	printableActualStr =
		sMech.ConvertNonPrintableChars(
			[]rune(actualStr),
			true)

	if outputStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"timerLines03.GetFormattedText()"+
			"Expected string DOES NOT match Actual string\n"+
			"Expected string = '%v'\n"+
			"  Actual string = '%v'\n",
			ePrefix.String(),
			outputStr,
			printableActualStr)

		return

	}

	timerLines04 := TextLineSpecTimerLines{}

	_,
		err = timerLines04.CopyOut(
		ePrefix.XCpy(
			"timerLines04"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from timerLines04."+
			"CopyOut()\n"+
			"because 'timerLines04' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	_,
		err = timerLines04.CopyOut(
		ePrefix.XCpy(
			"timerLines04"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from timerLines04."+
			"CopyOut()\n"+
			"because 'timerLines04' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	_,
		err = timerLines01.CopyOut(
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from timerLines01."+
			"CopyOut()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecTimerLines_CopyOutITextLine_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_CopyOutITextLine_000100()",
		"")

	outputStr,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = timerLines01.IsValidInstanceError(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var iTextLine ITextLineSpecification

	iTextLine,
		err =
		timerLines01.CopyOutITextLine(
			ePrefix.XCpy(
				"timerLines01->iTextLine"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	timerLines02,
		ok := iTextLine.(*TextLineSpecTimerLines)

	if !ok {
		t.Errorf("%v\n"+
			"Error: iTextLine.(*TextLineSpecTimerLines)\n"+
			"Could not convert 'iTextLine' to TextLineSpecTimerLines\n",
			ePrefix.String())

		return
	}

	if !timerLines02.Equal(timerLines01) {
		t.Errorf("%v\n"+
			"Error: Expected timerLines02 == timerLines01\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return
	}

	var actualStr string

	actualStr,
		err = timerLines02.GetFormattedText(
		ePrefix.XCpy(
			"timerLines02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sMech := StrMech{}

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(actualStr),
			true)

	if outputStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"timerLines02.GetFormattedText()"+
			"Expected string DOES NOT match Actual string\n"+
			"Expected string = '%v'\n"+
			"  Actual string = '%v'\n",
			ePrefix.String(),
			outputStr,
			printableActualStr)

		return
	}

	if !timerLines01.EqualITextLine(iTextLine) {
		t.Errorf("%v\n"+
			"Error: plainTextLine01.EqualITextLine(iTextLine)\n"+
			"Expected timerLines01 == iTextLine\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return
	}

	timerLines03 := TextLineSpecTimerLines{}

	_,
		err =
		timerLines03.CopyOutITextLine(
			ePrefix.XCpy(
				"timerLines03->_"))

	if err == nil {
		t.Errorf("%v\n"+
			"Error: timerLines03.CopyOutITextLine()\n"+
			"Expected an error return because 'timerLines03' is empty!\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	var timerLines04 *TextLineSpecTimerLines

	outputStr,
		timerLines04,
		err = createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"->timerLines04"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !timerLines04.IsValidInstance() {
		t.Errorf("%v\n"+
			"Error: timerLines04.IsValidInstance()\n"+
			"Expected a return value of 'true' because "+
			"'timerLines04' should be valid.\n"+
			"HOWEVER, A VALUD OF 'false' WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	_,
		err =
		timerLines04.CopyOutITextLine(
			StrMech{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from timerLines04."+
			"CopyOutITextLine()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecTimerLines_CopyOutPtr_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_CopyOutPtr_000100()",
		"")

	outputStr,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = timerLines01.IsValidInstanceError(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var timerLines02 *TextLineSpecTimerLines

	timerLines02,
		err = timerLines01.CopyOutPtr(
		ePrefix.XCpy(
			"timerLines02<-timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var actualStr string

	actualStr,
		err = timerLines02.GetFormattedText(
		ePrefix.XCpy(
			"timerLines02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sMech := StrMech{}

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(actualStr),
			true)

	if outputStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"timerLines02.GetFormattedText()"+
			"Expected string DOES NOT match Actual string\n"+
			"Expected string = '%v'\n"+
			"  Actual string = '%v'\n",
			ePrefix.String(),
			outputStr,
			printableActualStr)

		return
	}

	var timerLines03 *TextLineSpecTimerLines

	_,
		timerLines03,
		err = createTestTextLineSpecTimerLines02(
		ePrefix.XCpy(
			"timerLines03"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	timerLines03,
		err = timerLines02.CopyOutPtr(
		ePrefix.XCpy(
			"timerLines03<-timerLines02"))

	actualStr,
		err = timerLines03.GetFormattedText(
		ePrefix.XCpy(
			"timerLines03"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	printableActualStr =
		sMech.ConvertNonPrintableChars(
			[]rune(actualStr),
			true)

	if outputStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"timerLines03.GetFormattedText()"+
			"Expected string DOES NOT match Actual string\n"+
			"Expected string = '%v'\n"+
			"  Actual string = '%v'\n",
			ePrefix.String(),
			outputStr,
			printableActualStr)

		return

	}

	timerLines04 := TextLineSpecTimerLines{}

	_,
		err = timerLines04.CopyOutPtr(
		ePrefix.XCpy(
			"timerLines04"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from timerLines04."+
			"CopyOutPtr()\n"+
			"because 'timerLines04' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	_,
		err = timerLines04.CopyOutPtr(
		ePrefix.XCpy(
			"timerLines04"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from timerLines04."+
			"CopyOut()\n"+
			"because 'timerLines04' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	_,
		err = timerLines01.CopyOutPtr(
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from timerLines01."+
			"CopyOutPtr()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecTimerLines_Empty_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_Empty_000100()",
		"")

	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = timerLines01.IsValidInstanceError(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	timerLines01.Empty()

	if timerLines01.labelLeftMarginChars != nil {

		t.Errorf("%v - ERROR\n"+
			"timerLines01.Empty()"+
			"Expected timerLines01.labelLeftMarginChars = nil\n"+
			"Instead  timerLines01.labelLeftMarginChars = '%v'\n",
			ePrefix.String(),
			string(timerLines01.labelLeftMarginChars))

		return
	}

	if timerLines01.startTimeLabel != nil {

		t.Errorf("%v - ERROR\n"+
			"timerLines01.Empty()"+
			"Expected timerLines01.startTimeLabel = nil\n"+
			"Instead  timerLines01.startTimeLabel = '%v'\n",
			ePrefix.String(),
			string(timerLines01.startTimeLabel))

		return
	}

	if !timerLines01.startTime.IsZero() {

		t.Errorf("%v - ERROR\n"+
			"timerLines01.Empty()"+
			"Expected timerLines01.startTime = Zero\n"+
			"Instead  timerLines01.startTime = '%v'\n",
			ePrefix.String(),
			timerLines01.startTime.String())

		return
	}

	if timerLines01.endTimeLabel != nil {

		t.Errorf("%v - ERROR\n"+
			"timerLines01.Empty()"+
			"Expected timerLines01.endTimeLabel = nil\n"+
			"Instead  timerLines01.endTimeLabel = '%v'\n",
			ePrefix.String(),
			string(timerLines01.endTimeLabel))

		return
	}

	if !timerLines01.endTime.IsZero() {

		t.Errorf("%v - ERROR\n"+
			"timerLines01.Empty()"+
			"Expected timerLines01.endTime = Zero\n"+
			"Instead  timerLines01.endTime = '%v'\n",
			ePrefix.String(),
			timerLines01.endTime.String())

		return
	}

	if len(timerLines01.timeFormat) != 0 {

		t.Errorf("%v - ERROR\n"+
			"timerLines01.Empty()"+
			"Expected timerLines01.timeFormat = \"\"\n"+
			"Instead  timerLines01.timeFormat = '%v'\n",
			ePrefix.String(),
			timerLines01.timeFormat)

		return
	}

	if timerLines01.timeDurationLabel != nil {

		t.Errorf("%v - ERROR\n"+
			"timerLines01.Empty()"+
			"Expected timerLines01.timeDurationLabel = nil\n"+
			"Instead  timerLines01.timeDurationLabel = '%v'\n",
			ePrefix.String(),
			string(timerLines01.timeDurationLabel))

		return
	}

	if timerLines01.textLabelFieldLen != 0 {

		t.Errorf("%v - ERROR\n"+
			"timerLines01.Empty()"+
			"Expected timerLines01.textLabelFieldLen = ZERO\n"+
			"Instead  timerLines01.textLabelFieldLen = '%v'\n",
			ePrefix.String(),
			timerLines01.textLabelFieldLen)

		return
	}

	if timerLines01.textLabelJustification != TxtJustify.None() {

		t.Errorf("%v - ERROR\n"+
			"timerLines01.Empty()"+
			"Expected timerLines01.textLabelJustification = None\n"+
			"Instead  timerLines01.textLabelFieldLen = '%v'\n",
			ePrefix.String(),
			timerLines01.textLabelJustification.String())

		return
	}

	if timerLines01.labelRightMarginChars != nil {

		t.Errorf("%v - ERROR\n"+
			"timerLines01.Empty()"+
			"Expected timerLines01.labelRightMarginChars = nil\n"+
			"Instead  timerLines01.labelRightMarginChars = '%v'\n",
			ePrefix.String(),
			string(timerLines01.labelRightMarginChars))

		return
	}

	timerLines02 := TextLineSpecTimerLines{}

	timerLines02.Empty()

	return
}

func TestTextLineSpecTimerLines_Equal_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_Equal_000100()",
		"")

	outputStr01,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var timerLines02 *TextLineSpecTimerLines
	var outputStr02 string

	outputStr02,
		timerLines02,
		err = createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if outputStr01 != outputStr02 {

		t.Errorf("%v - ERROR\n"+
			"Expected outputStr01 == outputStr02\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!\n"+
			"outputStr01 = '%v'\n"+
			"outputStr02 = '%v'\n",
			ePrefix.String(),
			outputStr01,
			outputStr02)

		return

	}

	if !timerLines01.Equal(timerLines02) {

		t.Errorf("%v - ERROR\n"+
			"timerLines01.Equal(&timerLines02)\n"+
			"Expected timerLines01 == timerLines02\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!\n",
			ePrefix.String())

		return
	}

	timerLines02.labelRightMarginChars = nil

	if timerLines01.Equal(timerLines02) {

		t.Errorf("%v - ERROR\n"+
			"timerLines01.Equal(&timerLines02)\n"+
			"Expected timerLines01 != timerLines02\n"+
			"because timerLines02.labelRightMarginChars = nil"+
			"HOWEVER, THEY ARE EQUAL!!\n",
			ePrefix.String())

		return
	}

	timeLines03 := TextLineSpecTimerLines{}

	if timeLines03.Equal(timerLines01) {

		t.Errorf("%v - ERROR\n"+
			"timeLines03.Equal(timerLines01)\n"+
			"Expected timeLines03 != timerLines01\n"+
			"because timeLines03 is empty."+
			"HOWEVER, THEY ARE EQUAL!!\n",
			ePrefix.String())

		return

	}

	return
}

func TestTextLineSpecTimerLines_Equal_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_Equal_000200()",
		"")

	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var timerLines02 *TextLineSpecTimerLines

	timerLines02,
		err = timerLines01.CopyOutPtr(
		ePrefix.XCpy(
			"timerLines02<-timerLines01 #0"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var loc *time.Location

	loc,
		err = time.LoadLocation(
		"America/Chicago")

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	badStartTime := time.Date(
		1966,
		4,
		5,
		10,
		0,
		0,
		0,
		loc)

	timerLines01.startTime = badStartTime

	if timerLines01.Equal(timerLines02) {

		t.Errorf("%v - ERROR\n"+
			"timerLines01.Equal(&timerLines02)\n"+
			"Expected 'timerLines01' NOT EQUAL to 'timerLines02'\n"+
			"because timerLines01.startTime and timerLines02.startTime\n"+
			"are NOT Equal. \n"+
			"HOWEVER, THE RETURNED ANALYSIS SHOWS THEY ARE EQUAL!!\n",
			ePrefix.String())

		return
	}

	timerLines01,
		err = timerLines02.CopyOutPtr(
		ePrefix.XCpy(
			"timerLines01<-timerLines02 #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	timerLines01.endTime = badStartTime

	if timerLines01.Equal(timerLines02) {

		t.Errorf("%v - ERROR\n"+
			"timerLines01.Equal(&timerLines02)\n"+
			"Expected 'timerLines01' NOT EQUAL to 'timerLines02'\n"+
			"because timerLines01.endTime and\n"+
			"timerLines02.endTime are NOT Equal.\n"+
			"HOWEVER, THE RETURNED ANALYSIS SHOWS THEY ARE EQUAL!!\n",
			ePrefix.String())

		return
	}

	timerLines01,
		err = timerLines02.CopyOutPtr(
		ePrefix.XCpy(
			"timerLines01<-timerLines02 #2"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	timerLines01.endTimeLabel = []rune("Hello World")

	if timerLines01.Equal(timerLines02) {

		t.Errorf("%v - ERROR\n"+
			"timerLines01.Equal(&timerLines02)\n"+
			"Expected 'timerLines01' NOT EQUAL to 'timerLines02'\n"+
			"because timerLines01.endTimeLabel and\n"+
			"timerLines02.endTimeLabel are NOT Equal.\n"+
			"HOWEVER, THE RETURNED ANALYSIS SHOWS THEY ARE EQUAL!!\n",
			ePrefix.String())

		return
	}

	timerLines01,
		err = timerLines02.CopyOutPtr(
		ePrefix.XCpy(
			"timerLines01<-timerLines02 #3"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	timerLines01.endTimeLabel = []rune("Hello World")

	if timerLines01.Equal(timerLines02) {

		t.Errorf("%v - ERROR\n"+
			"timerLines01.Equal(&timerLines02)\n"+
			"Expected 'timerLines01' NOT EQUAL to 'timerLines02'\n"+
			"because timerLines01.endTimeLabel and\n"+
			"timerLines02.endTimeLabel are NOT Equal.\n"+
			"HOWEVER, THE RETURNED ANALYSIS SHOWS THEY ARE EQUAL!!\n",
			ePrefix.String())

		return
	}

	timerLines01,
		err = timerLines02.CopyOutPtr(
		ePrefix.XCpy(
			"timerLines01<-timerLines02 #3"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	timerLines01.timeFormat = "Hello World"

	if timerLines01.Equal(timerLines02) {

		t.Errorf("%v - ERROR\n"+
			"timerLines01.Equal(&timerLines02)\n"+
			"Expected 'timerLines01' NOT EQUAL to 'timerLines02'\n"+
			"because timerLines01.timeFormat and\n"+
			"timerLines02.timeFormat are NOT Equal.\n"+
			"HOWEVER, THE RETURNED ANALYSIS SHOWS THEY ARE EQUAL!!\n",
			ePrefix.String())

		return
	}

	timerLines01,
		err = timerLines02.CopyOutPtr(
		ePrefix.XCpy(
			"timerLines01<-timerLines02 #4"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	timerLines01.timeDurationLabel = []rune("Hello World")

	if timerLines01.Equal(timerLines02) {

		t.Errorf("%v - ERROR\n"+
			"timerLines01.Equal(&timerLines02)\n"+
			"Expected 'timerLines01' NOT EQUAL to 'timerLines02'\n"+
			"because timerLines01.timeDurationLabel and\n"+
			"timerLines02.timeDurationLabel are NOT Equal.\n"+
			"HOWEVER, THE RETURNED ANALYSIS SHOWS THEY ARE EQUAL!!\n",
			ePrefix.String())

		return
	}

	timerLines01,
		err = timerLines02.CopyOutPtr(
		ePrefix.XCpy(
			"timerLines01<-timerLines02 #5"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	timerLines01.textLabelFieldLen = 7

	if timerLines01.Equal(timerLines02) {

		t.Errorf("%v - ERROR\n"+
			"timerLines01.Equal(&timerLines02)\n"+
			"Expected 'timerLines01' NOT EQUAL to 'timerLines02'\n"+
			"because timerLines01.textLabelFieldLen and\n"+
			"timerLines02.textLabelFieldLen are NOT Equal.\n"+
			"HOWEVER, THE RETURNED ANALYSIS SHOWS THEY ARE EQUAL!!\n",
			ePrefix.String())

		return
	}

	timerLines01,
		err = timerLines02.CopyOutPtr(
		ePrefix.XCpy(
			"timerLines01<-timerLines02 #6"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	timerLines01.textLabelJustification = TxtJustify.None()

	if timerLines01.Equal(timerLines02) {

		t.Errorf("%v - ERROR\n"+
			"timerLines01.Equal(&timerLines02)\n"+
			"Expected 'timerLines01' NOT EQUAL to 'timerLines02'\n"+
			"because timerLines01.textLabelJustification and\n"+
			"timerLines02.textLabelJustification are NOT Equal.\n"+
			"HOWEVER, THE RETURNED ANALYSIS SHOWS THEY ARE EQUAL!!\n",
			ePrefix.String())

		return
	}

}

func TestTextLineSpecTimerLines_EqualITextLine_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_EqualITextLine_000100()",
		"")

	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	areEqual := timerLines01.EqualITextLine(
		nil)

	if areEqual == true {

		t.Errorf("%v\n"+
			"Error: timerLines01.EqualITextLine()\n"+
			"Expected areEqual == true because input\n"+
			"parameter 'nil' is invalid.\n"+
			"HOWEVER, areEqual == true\n",
			ePrefix.String())

		return
	}

	var timerLines02 TextLineSpecTimerLines

	timerLines02,
		err = timerLines01.CopyOut(
		ePrefix.XCpy(
			"timerLines02<-timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	iTextLineSpec := ITextLineSpecification(timerLines01)

	areEqual = timerLines02.EqualITextLine(
		iTextLineSpec)

	if areEqual == false {

		t.Errorf("%v\n"+
			"Error: timerLines02.EqualITextLine()\n"+
			"Expected areEqual == true because input\n"+
			"parameter 'timerLines02' is valid and equal.\n"+
			"HOWEVER, areEqual == false\n",
			ePrefix.String())

		return
	}

	timerLines03 := TextLineSpecTimerLines{}

	areEqual = timerLines03.EqualITextLine(
		iTextLineSpec)

	if areEqual == true {

		t.Errorf("%v\n"+
			"Error: timerLines03.EqualITextLine()\n"+
			"Expected areEqual == false because input\n"+
			"parameter 'timerLines03' is empty.\n"+
			"HOWEVER, areEqual == true\n",
			ePrefix.String())

		return
	}

	var stdLine01 TextLineSpecStandardLine

	stdLine01,
		err = createTestTextLineSpecStandardLine04(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	iTextLineSpec = ITextLineSpecification(&stdLine01)

	areEqual = timerLines02.EqualITextLine(
		iTextLineSpec)

	if areEqual == true {

		t.Errorf("%v\n"+
			"Error: timerLines02.EqualITextLine()\n"+
			"Expected areEqual == false because input\n"+
			"parameter 'iTextLineSpec' is of type TextLineSpecStandardLine.\n"+
			"HOWEVER, areEqual == true\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecTimerLines_GetFormattedText_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_GetFormattedText_000100()",
		"")

	outputStr,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var actualStr string

	actualStr,
		err = timerLines01.GetFormattedText(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sMech := StrMech{}

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(actualStr),
			true)

	if outputStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"timerLines01.GetFormattedText()"+
			"Expected string DOES NOT match Actual string\n"+
			"Expected string = '%v'\n"+
			"  Actual string = '%v'\n",
			ePrefix.String(),
			outputStr,
			printableActualStr)

		return

	}

	return
}

func TestTextLineSpecTimerLines_GetFormattedText_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTESTSERIES_TESTMETHOD_000200()",
		"")

	outputStr,
		timerLines01,
		err := createTestTextLineSpecTimerLines02(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var actualStr string

	actualStr,
		err = timerLines01.GetFormattedText(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sMech := StrMech{}

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(actualStr),
			true)

	if outputStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"timerLines01.GetFormattedText()"+
			"Expected string DOES NOT match Actual string\n"+
			"Expected string = '%v'\n"+
			"  Actual string = '%v'\n",
			ePrefix.String(),
			outputStr,
			printableActualStr)

		return

	}

	return
}

func TestTextLineSpecTimerLines_GetFormattedText_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_GetFormattedText_000300()",
		"")

	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines02(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	_,
		err = timerLines01.GetFormattedText(
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from timerLines01."+
			"GetFormattedText()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecTimerLines_GetFormattedText_000400(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_GetFormattedText_000400()",
		"")

	timerLines01 := TextLineSpecTimerLines{}

	_,
		err := timerLines01.GetFormattedText(
		ePrefix.XCpy(
			"timerLines01"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from timerLines01.GetFormattedText()\n"+
			"because 'timerLines01' is empty.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!\n",
			ePrefix.String())

		return
	}

}

func TestTextLineSpecTimerLines_GetEndTimeLabel_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_GetEndTimeLabel_000100()",
		"")

	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	expectedStr := "End Time"

	actualStr := timerLines01.GetEndTimeLabel()

	if expectedStr != actualStr {

		t.Errorf("\n%v\n"+
			"Error: timerLines01.GetEndTimeLabel()\n"+
			"Expected string is not equal to atual string.\n"+
			"Expected String = '%v'\n"+
			"  Actual String = '%v'\n",
			ePrefix.String(),
			expectedStr,
			actualStr)
		return
	}

	timerLines02 := TextLineSpecTimerLines{}

	actualStr = timerLines02.GetEndTimeLabel()

	if actualStr != "" {

		t.Errorf("\n%v\n"+
			"Error: timerLines02.GetEndTimeLabel()\n"+
			"Expected an empty string because\n"+
			"timerLines02 is empty!\n"+
			"However, Actual String = '%v'\n",
			ePrefix.String(),
			actualStr)
		return

	}

	return
}

func TestTextLineSpecTimerLines_GetEndTime_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_GetEndTime_000100()",
		"")

	timerLines01 := TextLineSpecTimerLines{}

	expectedTime := timerLines01.GetEndTime()

	timeFormat := "2006-01-02 15:04:05.000000000 -0700 MST"

	if !expectedTime.IsZero() {

		t.Errorf("\n%v\n"+
			"Error: timerLines01.GetEndTime()\n"+
			"Expected the returned time to be Zero.\n"+
			"However, the Actual Returned Time = '%v'\n",
			ePrefix.String(),
			expectedTime.Format(timeFormat))
		return

	}

	_,
		timerLines02,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	expectedTime = timerLines02.endTime

	actualTime := timerLines02.GetEndTime()

	if !expectedTime.Equal(actualTime) {

		t.Errorf("\n%v\n"+
			"Error: timerLines02.GetEndTime()\n"+
			"Expected time is NOT EQUAL to Actual Time.\n"+
			"Expected Time = '%v'\n"+
			"  Actual Time = '%v'\n",
			ePrefix.String(),
			expectedTime.Format(timeFormat),
			actualTime.Format(timeFormat))
		return
	}

	expectedTimeStr := expectedTime.Format(timeFormat)

	actualTimeStr := actualTime.Format(timeFormat)

	if expectedTimeStr != actualTimeStr {

		t.Errorf("\n%v\n"+
			"Error: timerLines02.GetEndTime() Time Strings\n"+
			"expectedTimeStr != actualTimeStr\n"+
			"Expected Time = '%v'\n"+
			"  Actual Time = '%v'\n",
			ePrefix.String(),
			expectedTimeStr,
			actualTimeStr)
		return

	}

	return
}

func TestTextLineSpecTimerLines_GetStartTimeLabel_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_GetStartTimeLabel_000100()",
		"")

	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	expectedStr := "Start Time"

	actualStr := timerLines01.GetStartTimeLabel()

	if expectedStr != actualStr {

		t.Errorf("\n%v\n"+
			"Error: timerLines01.GetStartTimeLabel()\n"+
			"Expected string is not equal to atual string.\n"+
			"Expected String = '%v'\n"+
			"  Actual String = '%v'\n",
			ePrefix.String(),
			expectedStr,
			actualStr)
		return
	}

	timerLines02 := TextLineSpecTimerLines{}

	actualStr = timerLines02.GetStartTimeLabel()

	if actualStr != "" {

		t.Errorf("\n%v\n"+
			"Error: timerLines02.GetStartTimeLabel()\n"+
			"Expected an empty string because\n"+
			"timerLines02 is empty!\n"+
			"However, Actual String = '%v'\n",
			ePrefix.String(),
			actualStr)
		return

	}

	return
}

func TestTextLineSpecTimerLines_GetStartTime_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_GetStartTime_000100()",
		"")

	timerLines01 := TextLineSpecTimerLines{}

	expectedTime := timerLines01.GetStartTime()

	timeFormat := "2006-01-02 15:04:05.000000000 -0700 MST"

	if !expectedTime.IsZero() {

		t.Errorf("\n%v\n"+
			"Error: timerLines01.GetEndTime()\n"+
			"Expected the returned time to be Zero.\n"+
			"However, the Actual Returned Time = '%v'\n",
			ePrefix.String(),
			expectedTime.Format(timeFormat))
		return
	}

	_,
		timerLines02,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	expectedTime = timerLines02.startTime

	actualTime := timerLines02.GetStartTime()

	if !expectedTime.Equal(actualTime) {

		t.Errorf("\n%v\n"+
			"Error: timerLines02.GetStartTime()\n"+
			"Expected time is NOT EQUAL to Actual Time.\n"+
			"Expected Time = '%v'\n"+
			"  Actual Time = '%v'\n",
			ePrefix.String(),
			expectedTime.Format(timeFormat),
			actualTime.Format(timeFormat))
		return
	}

	expectedTimeStr := expectedTime.Format(timeFormat)

	actualTimeStr := actualTime.Format(timeFormat)

	if expectedTimeStr != actualTimeStr {

		t.Errorf("\n%v\n"+
			"Error: timerLines02.GetStartTime() Time Strings\n"+
			"expectedTimeStr != actualTimeStr\n"+
			"Expected Time = '%v'\n"+
			"  Actual Time = '%v'\n",
			ePrefix.String(),
			expectedTimeStr,
			actualTimeStr)
		return

	}

	return
}

func TestTextLineSpecTimerLines_GetLabelFieldLength_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_GetLabelFieldLength_000100()",
		"")

	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	expectedLabelFieldLength := timerLines01.textLabelFieldLen

	actualLabelFieldLength := timerLines01.GetLabelFieldLength()

	if expectedLabelFieldLength != actualLabelFieldLength {

		t.Errorf("\n%v\n"+
			"Error: timerLines01.GetLabelFieldLength()\n"+
			"Expected Label Field Length is NOT Equal to\n"+
			"Actual Label Field Length.\n"+
			"Expected Label Field Length = '%v'\n"+
			"  Actual Label Field Length = '%v'\n",
			ePrefix.String(),
			expectedLabelFieldLength,
			actualLabelFieldLength)

		return
	}

	timerLines02 := TextLineSpecTimerLines{}

	_ = timerLines02.GetLabelFieldLength()

	return
}

func TestTextLineSpecTimerLines_GetLabelJustification_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_GetLabelJustification_000100()",
		"")

	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	expectedLabelJustification := timerLines01.textLabelJustification

	actualLabelJustification := timerLines01.GetLabelJustification()

	if expectedLabelJustification != actualLabelJustification {

		t.Errorf("\n%v\n"+
			"Error: timerLines01.GetLabelJustification()\n"+
			"Expected Label Justification is NOT Equal to\n"+
			"the Actual Label Justification value.\n"+
			"Expected Label Justification = '%v'\n"+
			"  Actual Label Justification = '%v'\n",
			ePrefix.String(),
			expectedLabelJustification.String(),
			actualLabelJustification.String())

		return
	}

	timerLines02 := TextLineSpecTimerLines{}

	_ = timerLines02.GetLabelJustification()

	return
}

func TestTextLineSpecTimerLines_GetLabelOutputSeparationChars_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_GetLabelOutputSeparationChars_000100()",
		"")

	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	expectedLabelOutputChars := string(timerLines01.labelRightMarginChars)

	actualLabelOutputChars := timerLines01.GetLabelOutputSeparationChars()

	if expectedLabelOutputChars != actualLabelOutputChars {

		t.Errorf("\n%v\n"+
			"Error: timerLines01.GetLabelOutputSeparationChars()\n"+
			"Expected Label Separation Characters is NOT Equal to\n"+
			"the Actual Label Separation Characters.\n"+
			"Expected Label Separation Characters = '%v'\n"+
			"  Actual Label Separation Characters = '%v'\n",
			ePrefix.String(),
			expectedLabelOutputChars,
			actualLabelOutputChars)

		return
	}

	timerLines02 := TextLineSpecTimerLines{}

	_ = timerLines02.GetLabelOutputSeparationChars()

	return
}

func TestTextLineSpecTimerLines_GetTimeDurationLabel_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_GetTimeDurationLabel_000100()",
		"")

	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	expectedStr := string(timerLines01.timeDurationLabel)

	actualStr := timerLines01.GetTimeDurationLabel()

	if expectedStr != actualStr {

		t.Errorf("\n%v\n"+
			"Error: timerLines01.GetTimeDurationLabel()\n"+
			"Expected string is not equal to atual string.\n"+
			"Expected String = '%v'\n"+
			"  Actual String = '%v'\n",
			ePrefix.String(),
			expectedStr,
			actualStr)
		return
	}

	timerLines02 := TextLineSpecTimerLines{}

	actualStr = timerLines02.GetTimeDurationLabel()

	if actualStr != "" {

		t.Errorf("\n%v\n"+
			"Error: timerLines02.GetTimeDurationLabel()\n"+
			"Expected an empty string because\n"+
			"timerLines02 is empty!\n"+
			"However, Actual String = '%v'\n",
			ePrefix.String(),
			actualStr)
		return

	}

	return
}

func TestTextLineSpecTimerLines_GetTimeFormat_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_GetTimeFormat_000100()",
		"")

	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	expectedStr := timerLines01.timeFormat

	actualStr := timerLines01.GetTimeFormat()

	if expectedStr != actualStr {

		t.Errorf("\n%v\n"+
			"Error: timerLines01.GetTimeFormat()\n"+
			"Expected string is not equal to atual string.\n"+
			"Expected String = '%v'\n"+
			"  Actual String = '%v'\n",
			ePrefix.String(),
			expectedStr,
			actualStr)
		return
	}

	timerLines02 := TextLineSpecTimerLines{}

	actualStr = timerLines02.GetTimeFormat()

	if actualStr != "" {

		t.Errorf("\n%v\n"+
			"Error: timerLines02.GetTimeFormat()\n"+
			"Expected an empty string because\n"+
			"timerLines02 is empty!\n"+
			"However, Actual String = '%v'\n",
			ePrefix.String(),
			actualStr)
		return

	}

	return
}

func TestTextLineSpecTimerLines_IsValidInstance_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_IsValidInstance_000100()",
		"")

	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	isValid := timerLines01.IsValidInstance()

	if !isValid {

		t.Errorf("\n%v\n"+
			"Error: isValid := timerLines01.IsValidInstance()\n"+
			"Expected 'isValid' == 'true' because 'timerLines01' is valid.\n"+
			"HOWEVER 'isValid' == 'false' !!\n",
			ePrefix.String())
		return
	}

	timerLines02 := TextLineSpecTimerLines{}

	isValid = timerLines02.IsValidInstance()

	if isValid {

		t.Errorf("\n%v\n"+
			"Error: isValid := timerLines02.IsValidInstance()\n"+
			"Expected 'isValid' == 'false' because 'timerLines02' is empty.\n"+
			"HOWEVER 'isValid' == 'true' !!\n",
			ePrefix.String())
		return
	}

	return
}

func TestTextLineSpecTimerLines_IsValidInstanceError_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_IsValidInstanceError_000100()",
		"")

	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = timerLines01.IsValidInstanceError(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	timerLines02 := TextLineSpecTimerLines{}

	err = timerLines02.IsValidInstanceError(
		ePrefix.XCpy(
			"timerLines02"))

	if err == nil {

		t.Errorf("\n%v\n"+
			"Error: err = timerLines02.IsValidInstanceError()\n"+
			"Expected an error return because 'timerLines02' is empty.\n"+
			"HOWEVER NO ERROR WAS RETURNED!!\n",
			ePrefix.String())
		return
	}

	return
}

func TestTextLineSpecTimerLines_IsValidInstanceError_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_IsValidInstanceError_000200()",
		"")

	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	timerLines01.endTime = time.Time{}

	err = timerLines01.IsValidInstanceError(
		ePrefix.XCpy(
			"timerLines01"))

	if err == nil {

		t.Errorf("\n%v\n"+
			"Error: err = timerLines02.IsValidInstanceError()\n"+
			"Expected an error return because \n"+
			"'timerLines01.endTime' is zero.\n"+
			"HOWEVER NO ERROR WAS RETURNED!!\n",
			ePrefix.String())

		return
	}

	var timerLines02 *TextLineSpecTimerLines
	_,
		timerLines02,
		err = createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var loc *time.Location

	loc,
		err = time.LoadLocation(
		"America/Chicago")

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	badEndTime := time.Date(
		1966,
		4,
		5,
		10,
		0,
		0,
		0,
		loc)

	timerLines02.endTime = badEndTime

	err = timerLines02.IsValidInstanceError(
		ePrefix.XCpy(
			"timerLines02"))

	if err == nil {

		t.Errorf("\n%v\n"+
			"Error: err = timerLines02.IsValidInstanceError()\n"+
			"Expected an error return because \n"+
			"'timerLines02.endTime' is before start time.\n"+
			"HOWEVER NO ERROR WAS RETURNED!!\n",
			ePrefix.String())
		return
	}

	var timerLines03 *TextLineSpecTimerLines
	_,
		timerLines03,
		err = createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines03"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	timerLines03.textLabelFieldLen = 1000001

	err = timerLines03.IsValidInstanceError(
		ePrefix.XCpy(
			"timerLines03"))

	if err == nil {

		t.Errorf("\n%v\n"+
			"Error: err = timerLines03.IsValidInstanceError()\n"+
			"Expected an error return because \n"+
			"'timerLines03.textLabelFieldLen' is greater than '1000000'.\n"+
			"HOWEVER NO ERROR WAS RETURNED!!\n",
			ePrefix.String())

		return
	}

	timerLines03.textLabelFieldLen = -70

	err = timerLines03.IsValidInstanceError(
		ePrefix.XCpy(
			"timerLines03"))

	if err != nil {

		t.Errorf("\n%v\n"+
			"Error: err = timerLines03.IsValidInstanceError()\n"+
			"Test # 2\n"+
			"DID NOT Expect an error return because \n"+
			"'timerLines03.textLabelFieldLen' is '-70'.\n"+
			"HOWEVER AN ERROR WAS RETURNED!!\n"+
			"Error= \n'%v\n",
			ePrefix.String(),
			err)

		return
	}

	var timerLines04 *TextLineSpecTimerLines
	_,
		timerLines04,
		err = createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines04"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	timerLines04.timeFormat = ""

	err = timerLines04.IsValidInstanceError(
		ePrefix.XCpy(
			"timerLines04"))

	if err != nil {

		t.Errorf("\n%v\n"+
			"Error: err = timerLines04.IsValidInstanceError()\n"+
			"DID NOT Expect an error return because \n"+
			"'timerLines04.timeFormat' is empty.\n"+
			"HOWEVER AN ERROR WAS RETURNED!!\n"+
			"Error= \n'%v\n",
			ePrefix.String(),
			err)

		return
	}

	var timerLines05 *TextLineSpecTimerLines

	_,
		timerLines05,
		err = createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines05"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	timerLines05.startTimeLabel = []rune{}

	err = timerLines05.IsValidInstanceError(
		ePrefix.XCpy(
			"timerLines05"))

	if err != nil {

		t.Errorf("\n%v\n"+
			"Error: err = timerLines05.IsValidInstanceError()\n"+
			"DID NOT Expect an error return because \n"+
			"'timerLines05.startTimeLabel' is empty.\n"+
			"HOWEVER AN ERROR WAS RETURNED!!\n"+
			"Error= \n'%v\n",
			ePrefix.String(),
			err)

		return
	}

	var timerLines06 *TextLineSpecTimerLines
	_,
		timerLines06,
		err = createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines06"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	timerLines06.endTimeLabel = []rune{}

	err = timerLines06.IsValidInstanceError(
		ePrefix.XCpy(
			"timerLines06"))

	if err != nil {

		t.Errorf("\n%v\n"+
			"Error: err = timerLines06.IsValidInstanceError()\n"+
			"DID NOT Expect an error return because \n"+
			"'timerLines06.endTimeLabel' is empty.\n"+
			"HOWEVER AN ERROR WAS RETURNED!!\n"+
			"Error= \n'%v\n",
			ePrefix.String(),
			err)

		return
	}

	var timerLines07 *TextLineSpecTimerLines
	_,
		timerLines07,
		err = createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines07"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	timerLines07.timeDurationLabel = []rune{}

	err = timerLines07.IsValidInstanceError(
		ePrefix.XCpy(
			"timerLines07"))

	if err != nil {

		t.Errorf("\n%v\n"+
			"Error: err = timerLines07.IsValidInstanceError()\n"+
			"DID NOT Expect an error return because \n"+
			"'timerLines07.timeDurationLabel' is empty.\n"+
			"HOWEVER AN ERROR WAS RETURNED!!\n"+
			"Error= \n'%v\n",
			ePrefix.String(),
			err)

		return
	}

	var timerLines08 *TextLineSpecTimerLines
	_,
		timerLines08,
		err = createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines08"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	timerLines08.labelRightMarginChars = []rune{}

	err = timerLines08.IsValidInstanceError(
		ePrefix.XCpy(
			"timerLines08"))

	if err != nil {

		t.Errorf("\n%v\n"+
			"Error: err = timerLines08.IsValidInstanceError()\n"+
			"DID NOT Expect an error return because \n"+
			"'timerLines08.labelRightMarginChars' is empty.\n"+
			"HOWEVER AN ERROR WAS RETURNED!!\n"+
			"Error= \n'%v\n",
			ePrefix.String(),
			err)

		return
	}

	var timerLines09 *TextLineSpecTimerLines
	_,
		timerLines09,
		err = createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines09"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	timerLines09.textLabelFieldLen = 60

	err = timerLines09.IsValidInstanceError(
		ePrefix.XCpy(
			"timerLines09"))

	if err == nil {

		t.Errorf("\n%v\n"+
			"Error: err = timerLines09.IsValidInstanceError()\n"+
			"Expected an error return because \n"+
			"'timerLines09.textLabelFieldLen' is exceeds maximum field length.\n"+
			"HOWEVER NO ERROR WAS RETURNED!!\n",
			ePrefix.String())

		return
	}

	var timerLines10 *TextLineSpecTimerLines
	_,
		timerLines10,
		err = createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines10"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	timerLines10.textLabelJustification = -9

	err = timerLines10.IsValidInstanceError(
		ePrefix.XCpy(
			"timerLines10"))

	if err == nil {

		t.Errorf("\n%v\n"+
			"Error: err = timerLines10.IsValidInstanceError()\n"+
			"Expected an error return because \n"+
			"'timerLines10.textLabelJustification' is invalid.\n"+
			"HOWEVER NO ERROR WAS RETURNED!!\n",
			ePrefix.String())

		return
	}

	var timerLines11 *TextLineSpecTimerLines
	_,
		timerLines11,
		err = createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines11"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = timerLines11.IsValidInstanceError(
		StrMech{})

	if err == nil {

		t.Errorf("\n%v\n"+
			"Error: err = timerLines11.IsValidInstanceError()\n"+
			"Expected an error return because the error prefix\n"+
			"is invalid.\n"+
			"HOWEVER NO ERROR WAS RETURNED!!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecTimerLines_NewDefaultFullTimerEvent_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_NewDefaultFullTimerEvent_000100()",
		"")

	var loc *time.Location
	var err error

	loc,
		err = time.LoadLocation(
		"America/Chicago")

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	startTime := time.Date(
		2022,
		4,
		5,
		10,
		0,
		0,
		0,
		loc)

	endTime := startTime.Add((time.Microsecond * 5) + 999)

	_,
		err = TextLineSpecTimerLines{}.NewDefaultFullTimerEvent(
		startTime,
		endTime,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from TextLineSpecTimerLines{}"+
			"NewDefaultFullTimerEvent()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecTimerLines_NewDefaultShellTimerEvent_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_NewDefaultShellTimerEvent_000100()",
		"")
	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var loc *time.Location

	loc,
		err = time.LoadLocation(
		"America/Chicago")

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	startTime := time.Date(
		2022,
		4,
		5,
		10,
		0,
		0,
		0,
		loc)

	endTime := startTime.Add((time.Microsecond * 5) + 999)

	timerLines02 :=
		TextLineSpecTimerLines{}.NewDefaultShellTimerEvent()

	err = timerLines02.SetStartAndEndTime(
		startTime,
		endTime,
		ePrefix.XCpy(
			"timerLines02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	areEqual := timerLines01.Equal(timerLines02)

	if !areEqual {

		t.Errorf("\n%v\n"+
			"Error:\n"+
			"Expected 'timerLines01' EQUAL 'timerLines02'.\n"+
			"HOWEVER THEY ARE NOT EQUAL!!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecTimerLines_NewEmptyTimerEvent_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_NewEmptyTimerEvent_000100()",
		"")

	timerLines01 := TextLineSpecTimerLines{}.NewEmptyTimerEvent()

	isValid := timerLines01.IsValidInstance()

	if isValid {

		t.Errorf("\n%v\n"+
			"Error: isValid = timerLines01.IsValidInstance()\n"+
			"Expected 'isValid' equals 'false' because\n"+
			"'timerLines01' is empty and invalid.\n"+
			"HOWEVER 'isValid' IS 'true' !!\n",
			ePrefix.String())

		return
	}

}

func TestTextLineSpecTimerLines_NewFullTimerEvent_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_NewFullTimerEvent_000100()",
		"")
	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var timerLines02 *TextLineSpecTimerLines

	timerLines02,
		err = TextLineSpecTimerLines{}.NewFullTimerEvent(
		string(timerLines01.labelLeftMarginChars),
		string(timerLines01.startTimeLabel),
		timerLines01.startTime,
		string(timerLines01.endTimeLabel),
		timerLines01.endTime,
		timerLines01.timeFormat,
		string(timerLines01.timeDurationLabel),
		timerLines01.textLabelFieldLen,
		timerLines01.textLabelJustification,
		string(timerLines01.labelRightMarginChars),
		ePrefix.XCpy(
			"timerLines02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = timerLines02.IsValidInstanceError(
		ePrefix.XCpy(
			"timerLines02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !timerLines01.Equal(timerLines02) {

		t.Errorf("\n%v\n"+
			"Error timerLines02\n"+
			"Expected 'timerLines01' would be equal to 'timerLines02'.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!\n",
			ePrefix.String())

		return
	}

	badEndTime := timerLines01.startTime

	badEndTime = badEndTime.Add(-5000000000)

	_,
		err = TextLineSpecTimerLines{}.NewFullTimerEvent(
		string(timerLines01.labelLeftMarginChars),
		string(timerLines01.startTimeLabel),
		timerLines01.startTime,
		string(timerLines01.endTimeLabel),
		badEndTime,
		timerLines01.timeFormat,
		string(timerLines01.timeDurationLabel),
		timerLines01.textLabelFieldLen,
		timerLines01.textLabelJustification,
		string(timerLines01.labelRightMarginChars),
		ePrefix.XCpy(
			"timerLines03"))

	if err == nil {

		t.Errorf("\n%v\n"+
			"Error TextLineSpecTimerLines{}.NewFullTimerEvent()\n"+
			"Expected an error return because\n"+
			"input parameter 'endTime' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED !!\n",
			ePrefix.String())

		return
	}

	_,
		err = TextLineSpecTimerLines{}.NewFullTimerEvent(
		string(timerLines01.labelLeftMarginChars),
		string(timerLines01.startTimeLabel),
		timerLines01.startTime,
		string(timerLines01.endTimeLabel),
		timerLines01.endTime,
		timerLines01.timeFormat,
		string(timerLines01.timeDurationLabel),
		timerLines01.textLabelFieldLen,
		timerLines01.textLabelJustification,
		string(timerLines01.labelRightMarginChars),
		StrMech{})

	if err == nil {

		t.Errorf("\n%v\n"+
			"Error TextLineSpecTimerLines{}.NewFullTimerEvent()\n"+
			"Expected an error return because\n"+
			"input parameter 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED !!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecTimerLines_NewFullTimerEventRunes_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_NewFullTimerEventRunes_000100()",
		"")
	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var timerLines02 *TextLineSpecTimerLines

	timerLines02,
		err = TextLineSpecTimerLines{}.NewFullTimerEventRunes(
		timerLines01.labelLeftMarginChars,
		timerLines01.startTimeLabel,
		timerLines01.startTime,
		timerLines01.endTimeLabel,
		timerLines01.endTime,
		timerLines01.timeFormat,
		timerLines01.timeDurationLabel,
		timerLines01.textLabelFieldLen,
		timerLines01.textLabelJustification,
		timerLines01.labelRightMarginChars,
		ePrefix.XCpy(
			"timerLines02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = timerLines02.IsValidInstanceError(
		ePrefix.XCpy(
			"timerLines02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !timerLines01.Equal(timerLines02) {

		t.Errorf("\n%v\n"+
			"Error timerLines02\n"+
			"Expected 'timerLines01' would be equal to 'timerLines02'.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!\n",
			ePrefix.String())

		return
	}

	badEndTime := timerLines01.startTime

	badEndTime = badEndTime.Add(-5000000000)

	_,
		err = TextLineSpecTimerLines{}.NewFullTimerEventRunes(
		timerLines01.labelLeftMarginChars,
		timerLines01.startTimeLabel,
		timerLines01.startTime,
		timerLines01.endTimeLabel,
		badEndTime,
		timerLines01.timeFormat,
		timerLines01.timeDurationLabel,
		timerLines01.textLabelFieldLen,
		timerLines01.textLabelJustification,
		timerLines01.labelRightMarginChars,
		ePrefix.XCpy(
			"timerLines03"))

	if err == nil {

		t.Errorf("\n%v\n"+
			"Error TextLineSpecTimerLines{}.NewFullTimerEventRunes()\n"+
			"Expected an error return because\n"+
			"input parameter 'endTime' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED !!\n",
			ePrefix.String())

		return
	}

	_,
		err = TextLineSpecTimerLines{}.NewFullTimerEventRunes(
		timerLines01.labelLeftMarginChars,
		timerLines01.startTimeLabel,
		timerLines01.startTime,
		timerLines01.endTimeLabel,
		time.Time{},
		timerLines01.timeFormat,
		timerLines01.timeDurationLabel,
		timerLines01.textLabelFieldLen,
		timerLines01.textLabelJustification,
		timerLines01.labelRightMarginChars,
		ePrefix.XCpy(
			"timerLines04"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	_,
		err = TextLineSpecTimerLines{}.NewFullTimerEventRunes(
		timerLines01.labelLeftMarginChars,
		timerLines01.startTimeLabel,
		time.Time{},
		timerLines01.endTimeLabel,
		timerLines01.endTime,
		timerLines01.timeFormat,
		timerLines01.timeDurationLabel,
		timerLines01.textLabelFieldLen,
		timerLines01.textLabelJustification,
		timerLines01.labelRightMarginChars,
		ePrefix.XCpy(
			"timerLines05"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	_,
		err = TextLineSpecTimerLines{}.NewFullTimerEventRunes(
		timerLines01.labelLeftMarginChars,
		timerLines01.startTimeLabel,
		timerLines01.startTime,
		timerLines01.endTimeLabel,
		timerLines01.endTime,
		timerLines01.timeFormat,
		timerLines01.timeDurationLabel,
		timerLines01.textLabelFieldLen,
		timerLines01.textLabelJustification,
		timerLines01.labelRightMarginChars,
		StrMech{})

	if err == nil {

		t.Errorf("\n%v\n"+
			"Error TextLineSpecTimerLines{}.NewFullTimerEventRunes()\n"+
			"Expected an error return because\n"+
			"input parameter 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED !!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecTimerLines_NewShellTimerEvent_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_NewShellTimerEvent_000100()",
		"")
	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var timerLines02 *TextLineSpecTimerLines

	timerLines02,
		err = TextLineSpecTimerLines{}.NewShellTimerEvent(
		string(timerLines01.labelLeftMarginChars),
		string(timerLines01.startTimeLabel),
		string(timerLines01.endTimeLabel),
		timerLines01.timeFormat,
		string(timerLines01.timeDurationLabel),
		timerLines01.textLabelFieldLen,
		timerLines01.textLabelJustification,
		string(timerLines01.labelRightMarginChars),
		ePrefix.XCpy(
			"timerLines02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = timerLines02.SetStartAndEndTime(
		timerLines01.startTime,
		timerLines01.endTime,
		ePrefix.XCpy(
			"timerLines02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = timerLines02.IsValidInstanceError(
		ePrefix.XCpy(
			"timerLines02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !timerLines01.Equal(timerLines02) {

		t.Errorf("\n%v\n"+
			"Error timerLines02\n"+
			"Expected 'timerLines01' would be equal to 'timerLines02'.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!\n",
			ePrefix.String())

		return
	}

	_,
		err = TextLineSpecTimerLines{}.NewShellTimerEvent(
		string(timerLines01.labelLeftMarginChars),
		string(timerLines01.startTimeLabel),
		string(timerLines01.endTimeLabel),
		timerLines01.timeFormat,
		string(timerLines01.timeDurationLabel),
		timerLines01.textLabelFieldLen,
		timerLines01.textLabelJustification,
		string(timerLines01.labelRightMarginChars),
		StrMech{})

	if err == nil {

		t.Errorf("\n%v\n"+
			"Error TextLineSpecTimerLines{}.NewShellTimerEvent()\n"+
			"Expected an error return because\n"+
			"input parameter 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED !!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecTimerLines_Read_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_Read_000100()",
		"")

	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err =
		timerLines01.IsValidInstanceError(
			ePrefix.XCpy(
				"timerLines01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var expectedTextStr string

	expectedTextStr,
		err =
		timerLines01.GetFormattedText(
			ePrefix.XCpy(
				"timerLines01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	lenExpectedStr := len(expectedTextStr)

	p := make([]byte, lenExpectedStr+1)

	var n, readBytesCnt int
	var actualStr string

	for {

		n,
			err = timerLines01.Read(p)

		if n == 0 {
			break
		}

		actualStr += string(p[:n])
		readBytesCnt += n
	}

	if err != nil &&
		err != io.EOF {
		t.Errorf("%v\n"+
			"Error Returned From timerLines01.Read(p)\n"+
			"Error = \n%v\n",
			ePrefix.String(),
			err.Error())

		return
	}

	if err == nil {
		t.Errorf("%v\n"+
			"Error: timerLines01.Read(p)\n"+
			"After completing Read Operation\n"+
			"the returned error should equal io.EOF.\n"+
			"HOWEVER, returned error == nil!\n",
			ePrefix.String())

		return
	}

	if err != io.EOF {
		t.Errorf("%v\n"+
			"Error: timerLines01.Read(p)\n"+
			"After completing Read Operation\n"+
			"the returned error should equal io.EOF.\n"+
			"HOWEVER, returned error is NOT equal io.EOF!\n",
			ePrefix.String())

		return
	}

	if timerLines01.textLineReader != nil {
		t.Errorf("%v\n"+
			"Error: After completing Read Operation\n"+
			"timerLines01.textLineReader != 'nil'\n",
			ePrefix.String())

		return
	}

	if readBytesCnt != lenExpectedStr {
		t.Errorf("%v\n"+
			"Byte Length Error: timerLines01.Read(p)\n"+
			"The actual length of bytes read\n"+
			"does NOT match the expected length.\n"+
			"Expected Bytes Read = '%v'\n"+
			"       Actual Bytes = '%v'\n",
			ePrefix.String(),
			lenExpectedStr,
			readBytesCnt)

		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedTextStr),
			true)

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(actualStr),
			true)

	if printableExpectedStr != printableActualStr {
		t.Errorf("%v\n"+
			"Error: Expected Text String DOES NOT match\n"+
			"Actual Text String.\n"+
			"Expected Text String = '%v'\n"+
			"Instead, Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	return
}

func TestTextLineSpecTimerLines_Read_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_Read_000200()",
		"")

	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err =
		timerLines01.IsValidInstanceError(
			ePrefix.XCpy(
				"timerLines01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var expectedTextStr string

	expectedTextStr,
		err =
		timerLines01.GetFormattedText(
			ePrefix.XCpy(
				"timerLines01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	lenExpectedStr := len(expectedTextStr)

	p := make([]byte, 1)

	var actualStr string

	var n, readBytesCnt int

	for {

		n,
			err = timerLines01.Read(p)

		if n == 0 {
			break
		}

		actualStr += string(p[:n])
		readBytesCnt += n

	}

	if err != nil &&
		err != io.EOF {
		t.Errorf("%v\n"+
			"Error Returned From timerLines01.Read(p)\n"+
			"Error = \n%v\n",
			ePrefix.String(),
			err.Error())

		return
	}

	if err != nil &&
		err != io.EOF {
		t.Errorf("%v\n"+
			"Error Returned From timerLines01.Read(p)\n"+
			"Error = \n%v\n",
			ePrefix.String(),
			err.Error())

		return
	}

	if err == nil {
		t.Errorf("%v\n"+
			"Error: After completing Read Operation\n"+
			"the returned error should equal io.EOF.\n"+
			"HOWEVER, returned error == nil!\n",
			ePrefix.String())

		return
	}

	if err != io.EOF {
		t.Errorf("%v\n"+
			"Error: After completing Read Operation\n"+
			"the returned error should equal io.EOF.\n"+
			"HOWEVER, returned error is NOT equal io.EOF!\n",
			ePrefix.String())

		return
	}

	if readBytesCnt != lenExpectedStr {
		t.Errorf("%v\n"+
			"Byte Length Error: timerLines01.Read(p)\n"+
			"The actual length of bytes read\n"+
			"does NOT match the expected length.\n"+
			"Expected Bytes Read = '%v'\n"+
			"       Actual Bytes = '%v'\n",
			ePrefix.String(),
			lenExpectedStr,
			readBytesCnt)

		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedTextStr),
			true)

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(actualStr),
			true)

	if printableExpectedStr != printableActualStr {
		t.Errorf("%v\n"+
			"Error: Expected Text String DOES NOT match\n"+
			"Actual Text String.\n"+
			"Expected Text String = '%v'\n"+
			"Instead, Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)
		return
	}

	if timerLines01.textLineReader != nil {
		t.Errorf("%v\n"+
			"Error: After a successful series of byte reads,\n"+
			"timerLines01.textLineReader pointer has NOT\n"+
			"BEEN RESET TO 'nil'!\n",
			ePrefix.String())
		return
	}

	p = make([]byte, 200)
	readBytesCnt = 0
	actualStr = ""

	for {

		n,
			err = timerLines01.Read(p)

		if n == 0 {
			break
		}

		actualStr += string(p[:n])
		readBytesCnt += n
	}

	if err != nil &&
		err != io.EOF {
		t.Errorf("%v\n"+
			"Error: Test # 2\n"+
			"Error Returned From timerLines01.Read(p)\n"+
			"Error = \n%v\n",
			ePrefix.String(),
			err.Error())

		return
	}

	printableActualStr =
		sMech.ConvertNonPrintableChars(
			[]rune(actualStr),
			true)

	if printableExpectedStr != printableActualStr {
		t.Errorf("%v\n"+
			"Error: Test # 2\n"+
			"Expected Text String DOES NOT match\n"+
			"Actual Text String.\n"+
			"Expected Text String = '%v'\n"+
			"Instead, Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)
		return
	}

	return
}

func TestTextLineSpecTimerLines_Read_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_Read_000300()",
		"")

	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err =
		timerLines01.IsValidInstanceError(
			ePrefix.XCpy(
				"timerLines01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var expectedTextStr string

	expectedTextStr,
		err =
		timerLines01.GetFormattedText(
			ePrefix.XCpy(
				"timerLines01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	lenExpectedStr := len(expectedTextStr)

	txtSpecAtom := textSpecificationAtom{}

	var n int
	p := make([]byte, 100)

	n,
		err = txtSpecAtom.readBytes(
		nil,
		p,
		ePrefix.XCpy("textReader == 'nil'"))

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected error return from txtSpecAtom.readBytes()"+
			"because input parameter 'textReader' == 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	var formattedTxtStr string
	timerLinesMolecule := textLineSpecTimerLinesMolecule{}

	formattedTxtStr,
		err =
		timerLinesMolecule.getFormattedText(
			timerLines01,
			ePrefix.XCpy("timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	p = make([]byte, 0)

	timerLines01.textLineReader =
		strings.NewReader(formattedTxtStr)

	n,
		err = txtSpecAtom.readBytes(
		timerLines01.textLineReader,
		p,
		ePrefix.XCpy("p == zero length"))

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected error return from txtSpecAtom.readBytes()"+
			"because input parameter 'p' is a zero length byte array.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	p = make([]byte, 100)

	var readBytesCnt int
	var actualStr string

	for {

		n,
			err = txtSpecAtom.readBytes(
			timerLines01.textLineReader,
			p,
			ePrefix.XCpy("timerLines01 is valid"))

		if n == 0 {
			break
		}

		actualStr += string(p[:n])
		readBytesCnt += n
	}

	if err != nil &&
		err != io.EOF {
		t.Errorf("%v\n"+
			"Error Returned From txtSpecAtom.readBytes(p)\n"+
			"Error = \n%v\n",
			ePrefix.String(),
			err.Error())

		return
	}

	if readBytesCnt != lenExpectedStr {
		t.Errorf("%v\n"+
			"Byte Length Error: txtSpecAtom.readBytes(p)\n"+
			"The actual length of bytes read\n"+
			"does NOT match the expected length.\n"+
			"Expected Bytes Read = '%v'\n"+
			"       Actual Bytes = '%v'\n",
			ePrefix.String(),
			lenExpectedStr,
			readBytesCnt)

		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedTextStr),
			true)

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(actualStr),
			true)

	if printableExpectedStr != printableActualStr {
		t.Errorf("%v\n"+
			"Error: Expected Text String DOES NOT match\n"+
			"Actual Text String.\n"+
			"Expected Text String = '%v'\n"+
			"Instead, Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	return
}

func TestTextLineSpecTimerLines_Read_000400(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_Read_000400()",
		"")
	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err =
		timerLines01.IsValidInstanceError(
			ePrefix.XCpy(
				"timerLines01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var expectedTextStr string

	expectedTextStr,
		err =
		timerLines01.GetFormattedText(
			ePrefix.XCpy(
				"timerLines01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	lenExpectedStr := len(expectedTextStr)

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedTextStr),
			true)

	p := make([]byte, 15)

	var n, readBytesCnt int
	sb := strings.Builder{}
	sb.Grow(512)

	for {

		n,
			err = timerLines01.Read(p)

		if n == 0 {
			break
		}

		sb.Write(p[:n])
		readBytesCnt += n
	}

	if err != nil &&
		err != io.EOF {
		t.Errorf("%v\n"+
			"Error Returned From timerLines01.Read(p)\n"+
			"Error = \n%v\n",
			ePrefix.String(),
			err.Error())

		return
	}

	if err == nil {
		t.Errorf("%v\n"+
			"Error: After completing Read Operation\n"+
			"the returned error should equal io.EOF.\n"+
			"HOWEVER, returned error == nil!\n",
			ePrefix.String())

		return
	}

	if err != io.EOF {
		t.Errorf("%v\n"+
			"Error: After completing Read Operation\n"+
			"the returned error should equal io.EOF.\n"+
			"HOWEVER, returned error is NOT equal io.EOF!\n",
			ePrefix.String())

		return
	}

	if timerLines01.textLineReader != nil {
		t.Errorf("%v\n"+
			"Error: After completing Read Operation\n"+
			"plainTextLine01.textLineReader != 'nil'\n",
			ePrefix.String())

		return
	}

	if readBytesCnt != lenExpectedStr {
		t.Errorf("%v\n"+
			"Byte Length Error: plainTextLine01.Read(p)\n"+
			"The actual length of bytes read\n"+
			"does NOT match the expected length.\n"+
			"Expected Bytes Read = '%v'\n"+
			"       Actual Bytes = '%v'\n",
			ePrefix.String(),
			lenExpectedStr,
			readBytesCnt)

		return
	}

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(sb.String()),
			true)

	if printableExpectedStr != printableActualStr {
		t.Errorf("%v\n"+
			"Error: Expected Text String DOES NOT match\n"+
			"Actual Text String.\n"+
			"Expected Text String = '%v'\n"+
			"Instead, Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)
	}

	return
}

func TestTextLineSpecTimerLines_Read_000500(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_Read_000500()",
		"")
	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err =
		timerLines01.IsValidInstanceError(
			ePrefix.XCpy(
				"timerLines01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	badEndTime := timerLines01.startTime
	timerLines01.startTime = timerLines01.endTime
	timerLines01.endTime = badEndTime
	timerLines01.textLabelFieldLen = -99
	timerLines01.textLabelJustification = -99

	p := make([]byte, 15)

	_,
		err = timerLines01.Read(p)

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from timerLines01.Read(p)\n"+
			"because 'timerLines01' contains invalid data.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecTimerLines_Read_000600(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_Read_000600()",
		"")

	p := make([]byte, 5)

	timerLines01 := TextLineSpecTimerLines{}

	_,
		err := timerLines01.Read(p)

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from timerLines01.Read(p)\n"+
			"because 'timerLines01' contains invalid data.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return

	}

	return
}

func TestTextLineSpecTimerLines_ReaderInitialize_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_ReaderInitialize_000100()",
		"")

	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err =
		timerLines01.IsValidInstanceError(
			ePrefix.XCpy(
				"timerLines01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var expectedTextStr string

	expectedTextStr,
		err =
		timerLines01.GetFormattedText(
			ePrefix.XCpy(
				"timerLines01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	lenExpectedStr := len(expectedTextStr)

	p := make([]byte, 5)

	var n int

	n,
		err = timerLines01.Read(p)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by timerLines01.Read(p)\n"+
			"Error:\n%v\n",
			ePrefix.String(),
			err.Error())

		return
	}

	if n != 5 {
		t.Errorf("%v\n"+
			"Error: timerLines01.Read(p)\n"+
			"Expected n == 5\n"+
			"Instead, n == %v\n",
			ePrefix.String(),
			n)

		return
	}

	p = make([]byte, 200)

	timerLines01.ReaderInitialize()

	var readBytesCnt int
	var actualStr string

	for {

		n,
			err = timerLines01.Read(p)

		if n == 0 {
			break
		}

		actualStr += string(p[:n])
		readBytesCnt += n
	}

	if err != nil &&
		err != io.EOF {
		t.Errorf("%v\n"+
			"Error Returned From timerLines01.Read(p)\n"+
			"Error = \n%v\n",
			ePrefix.String(),
			err.Error())

		return
	}

	if err == nil {
		t.Errorf("%v\n"+
			"Error: After completing Read Operation\n"+
			"the returned error should equal io.EOF.\n"+
			"HOWEVER, returned error == nil!\n",
			ePrefix.String())

		return
	}

	if err != io.EOF {
		t.Errorf("%v\n"+
			"Error: After completing Read Operation\n"+
			"the returned error should equal io.EOF.\n"+
			"HOWEVER, returned error is NOT equal io.EOF!\n",
			ePrefix.String())

		return
	}

	if timerLines01.textLineReader != nil {
		t.Errorf("%v\n"+
			"Error: After completing Read Operation\n"+
			"plainTextLine01.textLineReader != 'nil'\n",
			ePrefix.String())

		return
	}

	if readBytesCnt != lenExpectedStr {
		t.Errorf("%v\n"+
			"Byte Length Error: plainTextLine01.Read(p)\n"+
			"The actual length of bytes read\n"+
			"does NOT match the expected length.\n"+
			"Expected Bytes Read = '%v'\n"+
			"       Actual Bytes = '%v'\n",
			ePrefix.String(),
			lenExpectedStr,
			readBytesCnt)

		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedTextStr),
			true)

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(actualStr),
			true)

	if printableExpectedStr != printableActualStr {
		t.Errorf("%v\n"+
			"Error: Expected Text String DOES NOT match\n"+
			"Actual Text String.\n"+
			"Expected Text String = '%v'\n"+
			"Instead, Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	if timerLines01.textLineReader != nil {
		t.Errorf("%v Test #1\n"+
			"Completed Read Operation but timerLines01.textLineReader\n"+
			"is NOT equal to 'nil'!\n",
			ePrefix.String())

		return
	}

	p = make([]byte, 200)
	actualStr = ""
	readBytesCnt = 0

	for {

		n,
			err = timerLines01.Read(p)

		if n == 0 {
			break
		}

		actualStr += string(p[:n])
		readBytesCnt += n
	}

	if err != nil &&
		err != io.EOF {
		t.Errorf("%v\n"+
			"Test # 2"+
			"Error Returned From timerLines01.Read(p)\n"+
			"Error = \n%v\n",
			ePrefix.String(),
			err.Error())

		return
	}

	if err == nil {
		t.Errorf("%v\n"+
			"Test # 2"+
			"Error: After completing Read Operation\n"+
			"the returned error should equal io.EOF.\n"+
			"HOWEVER, returned error == nil!\n",
			ePrefix.String())

		return
	}

	if err != io.EOF {
		t.Errorf("%v\n"+
			"Test # 2"+
			"Error: After completing Read Operation\n"+
			"the returned error should equal io.EOF.\n"+
			"HOWEVER, returned error is NOT equal io.EOF!\n",
			ePrefix.String())

		return
	}

	printableActualStr =
		sMech.ConvertNonPrintableChars(
			[]rune(actualStr),
			true)

	if printableExpectedStr != printableActualStr {
		t.Errorf("%v Test #2\n"+
			"Error: Expected Text String DOES NOT match\n"+
			"Actual Text String.\n"+
			"Expected Text String = '%v'\n"+
			"Instead, Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	if timerLines01.textLineReader != nil {
		t.Errorf("%v Test #2\n"+
			"Completed Read Operation but timerLines01.textLineReader\n"+
			"is NOT equal to 'nil'!\n",
			ePrefix.String())

		return
	}

	timerLines02 := TextLineSpecTimerLines{}

	timerLines02.ReaderInitialize()

	return
}

func TestTextLineSpecTimerLines_SetDefaultFullTimerEvent_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_ReaderInitialize_000100()",
		"")

	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	timerLines02 := TextLineSpecTimerLines{}

	err = timerLines02.SetDefaultFullTimerEvent(
		timerLines01.startTime,
		timerLines01.endTime,
		ePrefix.XCpy(
			"timerLines02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !timerLines01.Equal(&timerLines02) {

		t.Errorf("\n%v - ERROR\n"+
			"Expected that timerLines01 would\n"+
			"be Equal To timerLines02.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!\n",
			ePrefix.String())

		return
	}

	timerLines03 := TextLineSpecTimerLines{}

	err = timerLines03.SetDefaultFullTimerEvent(
		timerLines01.startTime,
		timerLines01.endTime,
		StrMech{})

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from timerLines01."+
			"SetDefaultFullTimerEvent()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecTimerLines_SetStartAndEndTime_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_SetStartAndEndTime_000100()",
		"")

	var loc *time.Location
	var err error

	loc,
		err = time.LoadLocation(
		"America/Chicago")

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	startTime := time.Date(
		2022,
		4,
		5,
		10,
		0,
		0,
		0,
		loc)

	endTime := startTime.Add((time.Microsecond * 5) + 999)

	timerLines01 :=
		TextLineSpecTimerLines{}.NewDefaultShellTimerEvent()

	err = timerLines01.SetStartAndEndTime(
		endTime,
		startTime,
		ePrefix.XCpy(
			"timerLines01 start and end times reversed"))

	if err == nil {

		t.Errorf("\n%v\n"+
			"Error: err = timerLines01.SetStartAndEndTime()\n"+
			"Expected an error return because \n"+
			"start and end times are invalid.\n"+
			"HOWEVER NO ERROR WAS RETURNED!!\n",
			ePrefix.String())

		return
	}

	var timerLines02 TextLineSpecTimerLines

	err = timerLines02.SetStartAndEndTime(
		startTime,
		endTime,
		ePrefix.XCpy(
			"timerLines02 is empty"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	timerLines03 := TextLineSpecTimerLines{}

	err = timerLines03.SetStartAndEndTime(
		time.Time{},
		endTime,
		ePrefix.XCpy(
			"timerLines03"))

	if err == nil {

		t.Errorf("\n%v\n"+
			"Error: timerLines03.SetStartAndEndTime()\n"+
			"Expected an error return because \n"+
			"'startTime' is invalid.\n"+
			"HOWEVER NO ERROR WAS RETURNED!!\n",
			ePrefix.String())

		return
	}

	timerLines04 := TextLineSpecTimerLines{}

	err = timerLines04.SetStartAndEndTime(
		startTime,
		time.Time{},
		ePrefix.XCpy(
			"timerLines04"))

	if err == nil {

		t.Errorf("\n%v\n"+
			"Error: timerLines04.SetStartAndEndTime()\n"+
			"Expected an error return because \n"+
			"'endTime' is invalid.\n"+
			"HOWEVER NO ERROR WAS RETURNED!!\n",
			ePrefix.String())

		return
	}

	return
}
