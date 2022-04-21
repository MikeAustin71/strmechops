package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"testing"
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
