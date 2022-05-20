package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"testing"
	"time"
)

func TestTextLineSpecStandardLine_AddStandardLine_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_AddStandardLine_000100()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine01(
		ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	stdLine03 := TextLineSpecStandardLine{}
	lastIndex := 0

	lastIndex,
		err = stdLine03.AddStandardLine(
		&stdLine01,
		ePrefix.XCpy(
			"stdLine03<-stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	stdLine03.SetNumOfStdLines(1)

	expectedLastIndex := stdLine01.GetNumOfTextFields() - 1

	if lastIndex != expectedLastIndex {

		t.Errorf("\n%v - ERROR\n"+
			"Expected Last Index is NOT EQUAL to Actual Last Index!\n"+
			"Expected Last Index = '%v'\n"+
			"  Actual Last Index = '%v'\n",
			ePrefix.String(),
			expectedLastIndex,
			lastIndex)

		return

	}

	var stdLine02 TextLineSpecStandardLine

	stdLine02,
		err = createTestTextLineSpecStandardLine02(
		ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	_,
		err = stdLine03.AddStandardLine(
		&stdLine02,
		textLineSpecStandardLineElectron{})

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from stdLine02.DeleteTextField()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

}

func TestTextLineSpecStandardLine_AddStandardLine_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_AddStandardLine_000200()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine01(
		ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	_,
		err = stdLine01.AddStandardLine(
		nil,
		ePrefix.XCpy(
			"incomingStdLine==nil"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from stdLine01.AddStandardLine()\n"+
			"because 'incomingStdLine' is 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

}

func TestTextLineSpecStandardLine_AddTextField_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_AddTextField_000100()",
		"")

	stdLine01 := TextLineSpecStandardLine{}.New()

	labelTxt,
		err := TextFieldSpecLabel{}.NewTextLabel(
		"Hello World",
		-1,
		TxtJustify.Left(),
		ePrefix.XCpy(
			"labelTxt"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var indexId int

	indexId,
		err = stdLine01.AddTextField(
		&labelTxt,
		ePrefix.XCpy(
			"labelTxt"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if indexId != 0 {
		t.Errorf("%v - ERROR\n"+
			"stdLine01.AddTextField() should have\n"+
			"returned 'indexId' = 0\n"+
			"HOWEVER, indexId = %v\n",
			ePrefix.String(),
			indexId)

		return
	}

	stdLine02 := TextLineSpecStandardLine{}.New()

	_,
		err = stdLine02.AddTextField(
		&labelTxt,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from stdLine02{}."+
			"AddTextField()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	stdLine03 := TextLineSpecStandardLine{}.New()

	labelTxt.textLabel = nil
	labelTxt.fieldLen = -95

	_,
		err = stdLine03.AddTextField(
		&labelTxt,
		ePrefix.XCpy(
			"stdLine03<-labelTxt"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from stdLine03{}."+
			"AddTextField()\n"+
			"because 'labelTxt' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	stdLine04 := TextLineSpecStandardLine{}

	labelTxt.textLabel = []rune("Hello World")
	labelTxt.fieldLen = -1

	_,
		err = stdLine04.AddTextField(
		&labelTxt,
		ePrefix.XCpy(
			"stdLine04<-labelTxt"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	stdLine05 := TextLineSpecStandardLine{}

	_,
		err = stdLine05.AddTextField(
		nil,
		ePrefix.XCpy(
			"stdLine05<-labelTxt"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from stdLine05{}."+
			"AddTextField()\n"+
			"because 'textField' is nil.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecStandardLine_AddTextFields_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TextLineSpecStandardLine_AddTextFields_000100()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine01(
		ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var textFields []ITextFieldSpecification

	textFields,
		err = stdLine01.GetTextFieldCollection(
		ePrefix.XCpy(
			"textFields<-stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	stdLine03 := TextLineSpecStandardLine{}.New()
	lastIndex := 0

	lastIndex,
		err = stdLine03.AddTextFields(
		&textFields,
		ePrefix.XCpy(
			"stdLine03<-stdLine01.textFields"))

	expectedLastIndex := stdLine01.GetNumOfTextFields() - 1

	if lastIndex != expectedLastIndex {

		t.Errorf("\n%v - ERROR\n"+
			"stdLine01 vs stdLine03\n"+
			"Expected Last Index is NOT EQUAL to Actual Last Index!\n"+
			"Expected Last Index = '%v'\n"+
			"  Actual Last Index = '%v'\n",
			ePrefix.String(),
			expectedLastIndex,
			lastIndex)

		return

	}

	if !stdLine01.Equal(&stdLine03) {

		t.Errorf("\n%v - ERROR\n"+
			"Expected stdLine01==stdLine03\n"+
			"HOWEVER THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return

	}

	var stdLine02 TextLineSpecStandardLine

	stdLine02,
		err = createTestTextLineSpecStandardLine02(
		ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	textFields,
		err = stdLine02.GetTextFieldCollection(
		ePrefix.XCpy(
			"textFields<-stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	_,
		err = stdLine03.AddTextFields(
		&textFields,
		textLineSpecStandardLineElectron{})

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from stdLine03.AddTextFields()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

}

func TestTextLineSpecStandardLine_AddTextFields_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_AddTextFields_000200()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine01(
		ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var textFields []ITextFieldSpecification

	textFields,
		err = stdLine01.GetTextFieldCollection(
		ePrefix.XCpy(
			"textFields<-stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	stdLine02 := TextLineSpecStandardLine{}

	_,
		err = stdLine02.AddTextFields(
		&textFields,
		ePrefix.XCpy(
			"stdLine02 <- textFields"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	stdLine02.SetNumOfStdLines(1)

	textFields = nil

	_,
		err = stdLine02.AddTextFields(
		&textFields,
		ePrefix.XCpy(
			"Series-2: stdLine02 <- textFields"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from stdLine02.AddTextFields()\n"+
			"because 'textFields' is empty.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}
}

func TestTextLineSpecStandardLine_AddTextFieldDateTime_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_AddTextFieldDateTime_000100()",
		"")

	timeZoneName := "America/Chicago"

	tzLocPtr, err := time.LoadLocation(timeZoneName)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by time.LoadLocation(timeZoneName)\n"+
			"timeZoneName='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			timeZoneName,
			err.Error())

		return

	}

	dateTime := time.Date(
		2021,
		time.Month(12),
		27,
		15,
		28,
		0,
		0,
		tzLocPtr)

	dateTimeFormat :=
		"Monday January 2, 2006 15:04:05.000000000 -0700 MST"

	fieldLen := len(dateTimeFormat) + 8

	textJustification := TxtJustify.Center()

	stdLine01 := TextLineSpecStandardLine{}.New()

	var indexId int

	indexId,
		err = stdLine01.AddTextFieldDateTime(
		dateTime,
		fieldLen,
		dateTimeFormat,
		textJustification,
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = stdLine01.IsValidInstanceError(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if indexId != 0 {
		t.Errorf("%v - ERROR\n"+
			"stdLine01.AddTextFieldDateTime() should have\n"+
			"returned 'indexId' = 0\n"+
			"HOWEVER, indexId = %v\n",
			ePrefix.String(),
			indexId)

		return
	}

	// Add a second Date Time Field
	indexId,
		err = stdLine01.AddTextFieldDateTime(
		dateTime,
		fieldLen,
		dateTimeFormat,
		textJustification,
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if indexId != 1 {
		t.Errorf("%v - ERROR\n"+
			"Test # 2\n"+
			"stdLine01.AddTextFieldDateTime() should have\n"+
			"returned 'indexId' = 1\n"+
			"HOWEVER, indexId = %v\n",
			ePrefix.String(),
			indexId)

		return
	}

	stdLine02 := TextLineSpecStandardLine{}

	indexId,
		err = stdLine02.AddTextFieldDateTime(
		dateTime,
		fieldLen,
		dateTimeFormat,
		textJustification,
		ePrefix.XCpy(
			"stdLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if indexId != 0 {
		t.Errorf("%v - ERROR\n"+
			"stdLine02.AddTextFieldDateTime() should have\n"+
			"returned 'indexId' = 0\n"+
			"HOWEVER, indexId = %v\n",
			ePrefix.String(),
			indexId)

		return
	}

	stdLine03 := TextLineSpecStandardLine{}

	dateTimeBad := time.Time{}
	fieldLenBad := -97

	_,
		err = stdLine03.AddTextFieldDateTime(
		dateTimeBad,
		fieldLenBad,
		dateTimeFormat,
		textJustification,
		ePrefix.XCpy(
			"stdLine03"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from stdLine03{}."+
			"AddTextFieldDateTime()\n"+
			"because 'dateTimeBad' and 'fieldLenBad' are invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	stdLine04 := TextLineSpecStandardLine{}

	_,
		err = stdLine04.AddTextFieldDateTime(
		dateTime,
		fieldLen,
		dateTimeFormat,
		textJustification,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from stdLine04{}."+
			"AddTextFieldDateTime()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecStandardLine_AddTextFieldFiller_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_AddTextFieldFiller_000100()",
		"")

	stdLine01 := TextLineSpecStandardLine{}.New()

	fillerCharacters := " "
	fillerCharsRepeatCount := 5

	indexId,
		err := stdLine01.AddTextFieldFiller(
		fillerCharacters,
		fillerCharsRepeatCount,
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if indexId != 0 {

		t.Errorf("%v - ERROR\n"+
			"stdLine01.AddTextFieldFiller() should have\n"+
			"returned 'indexId' = 0\n"+
			"HOWEVER, indexId = %v\n",
			ePrefix.String(),
			indexId)

		return
	}

	// Add a second Text Filler Field
	indexId,
		err = stdLine01.AddTextFieldFiller(
		fillerCharacters,
		fillerCharsRepeatCount,
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = stdLine01.IsValidInstanceError(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if indexId != 1 {

		t.Errorf("%v - ERROR\n"+
			"stdLine01.AddTextFieldFiller() should have\n"+
			"returned 'indexId' = 1\n"+
			"HOWEVER, indexId = %v\n",
			ePrefix.String(),
			indexId)

		return
	}

	stdLine02 := TextLineSpecStandardLine{}

	_,
		err = stdLine02.AddTextFieldFiller(
		fillerCharacters,
		fillerCharsRepeatCount,
		ePrefix.XCpy(
			"stdLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	badFillerChars := ""

	_,
		err = stdLine02.AddTextFieldFiller(
		badFillerChars,
		fillerCharsRepeatCount,
		ePrefix.XCpy(
			"stdLine02 - badFillerChars"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from stdLine03{}."+
			"AddTextFieldDateTime()\n"+
			"because 'badFillerChars' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}
	_,
		err = stdLine02.AddTextFieldFiller(
		badFillerChars,
		fillerCharsRepeatCount,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from stdLine03{}."+
			"AddTextFieldDateTime()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecStandardLine_AddTextFieldLabel_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_AddTextFieldLabel_000100()",
		"")

	label := "12345"
	fieldLen := 13
	txtJustify := TxtJustify.Center()

	expectedStdLineText :=
		strings.Repeat(" ", 4) +
			label +
			strings.Repeat(" ", 4) +
			"\n"

	stdLine01 := TextLineSpecStandardLine{}.New()

	indexId,
		err := stdLine01.AddTextFieldLabel(
		label,
		fieldLen,
		txtJustify,
		ePrefix.XCpy(
			"stdLine01 - valid label"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if indexId != 0 {
		t.Errorf("%v - ERROR\n"+
			"stdLine01.AddTextFieldLabel() should have\n"+
			"returned 'indexId' = 0\n"+
			"HOWEVER, indexId = %v\n",
			ePrefix.String(),
			indexId)

		return
	}

	err = stdLine01.IsValidInstanceError(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedStdLineText),
			true)

	var actualStdLineText string

	actualStdLineText,
		err = stdLine01.GetFormattedText(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(actualStdLineText),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #1"+
			"Expected Formatted Text String 01 DOES NOT match\n"+
			"Actual Formatted Text String 01.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	stdLine02 := TextLineSpecStandardLine{}

	_,
		err = stdLine02.AddTextFieldLabel(
		label,
		fieldLen,
		txtJustify,
		ePrefix.XCpy(
			"stdLine02 is empty."))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	stdLine03 := TextLineSpecStandardLine{}.NewPtr()

	_,
		err = stdLine03.AddTextFieldLabel(
		label,
		fieldLen,
		txtJustify,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from stdLine03{}."+
			"AddTextFieldLabel()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	badFieldLen := -97

	_,
		err = stdLine03.AddTextFieldLabel(
		label,
		badFieldLen,
		txtJustify,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from stdLine03{}."+
			"AddTextFieldLabel()\n"+
			"because 'badFieldLen' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecStandardLine_AddTextFieldLabel_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_AddTextFieldLabel_000200()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine01(
		ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	label := ""
	fieldLen := -99
	txtJustify := TxtJustify.None()

	_,
		err = stdLine01.AddTextFieldLabel(
		label,
		fieldLen,
		txtJustify,
		ePrefix.XCpy(
			"stdLine01 - invalid label input parms"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from  stdLine01.AddTextFieldLabel()\n"+
			"because all input parameters are invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecStandardLine_AddTextFieldSpacer_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_AddTextFieldSpacer_000100()",
		"")

	expectedFieldLen := 4

	expectedStdLineText :=
		strings.Repeat(" ", expectedFieldLen) +
			"\n"

	stdLine01 := TextLineSpecStandardLine{}.New()

	indexId,
		err := stdLine01.AddTextFieldSpacer(
		expectedFieldLen,
		ePrefix.XCpy(
			"stdLine01 - valid Spacer"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if indexId != 0 {
		t.Errorf("%v - ERROR\n"+
			"stdLine01.AddTextFieldSpacer() should have\n"+
			"returned 'indexId' = 0\n"+
			"HOWEVER, indexId = %v\n",
			ePrefix.String(),
			indexId)

		return
	}

	err = stdLine01.IsValidInstanceError(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedStdLineText),
			true)

	var actualStdLineText string

	actualStdLineText,
		err = stdLine01.GetFormattedText(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(actualStdLineText),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #1"+
			"Expected Formatted Text String 01 DOES NOT match\n"+
			"Actual Formatted Text String 01.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	stdLine02 := TextLineSpecStandardLine{}

	_,
		err = stdLine02.AddTextFieldSpacer(
		expectedFieldLen,
		ePrefix.XCpy(
			"stdLine02 is empty."))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	stdLine03 := TextLineSpecStandardLine{}.NewPtr()

	_,
		err = stdLine03.AddTextFieldSpacer(
		expectedFieldLen,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from stdLine03{}."+
			"AddTextFieldSpacer()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	badFieldLen := -97

	_,
		err = stdLine03.AddTextFieldSpacer(
		badFieldLen,
		ePrefix.XCpy(
			"stdLine03 - badFieldLen invalid"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from stdLine03{}."+
			"AddTextFieldSpacer()\n"+
			"because 'badFieldLen' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecStandardLine_CopyIn_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_CopyIn_000100()",
		"")

	stdLine := TextLineSpecStandardLine{}.NewPtr()

	numOfStdLines := stdLine.GetNumOfStdLines()

	if numOfStdLines != 1 {
		t.Errorf("%v\n"+
			"Error: Expected Number of Standard Lines = '1'.\n"+
			"Instead, Number of Standard Lines = '%v'\n",
			ePrefix.String(),
			numOfStdLines)

		return
	}

	numOfTextFields := stdLine.GetNumOfTextFields()

	if numOfTextFields != 0 {
		t.Errorf("%v\n"+
			"Error: Expected Number of Standard Lines = '0'.\n"+
			"Instead, Number of Standard Lines = '%v'\n",
			ePrefix.String(),
			numOfTextFields)

		return
	}

	newLineChars := stdLine.GetNewLineChars()

	if newLineChars != "\n" {
		t.Errorf("%v\n"+
			"Error: Expected newLineChars = \"\\n\".\n"+
			"Instead, newLineChars = '%v'\n",
			ePrefix.String(),
			[]rune(newLineChars))

		return
	}

	newLineRunes := stdLine.GetNewLineRunes()
	expectedNewLineRunes := "\n"

	if string(newLineRunes) != expectedNewLineRunes {
		t.Errorf("%v\n"+
			"Error: Expected newLineRunes = \"\\n\".\n"+
			"Instead, newLineRunes = '%v'\n",
			ePrefix.String(),
			[]rune(newLineChars))

		return

	}

	turnLineTerminatorOff :=
		stdLine.GetTurnLineTerminatorOff()

	if turnLineTerminatorOff != false {
		t.Errorf("%v\n"+
			"Error: Expected turnLineTerminatorOff = 'false'.\n"+
			"Instead, turnLineTerminatorOff = 'true'\n",
			ePrefix.String())

		return
	}

	leftMargin,
		err := TextFieldSpecFiller{}.NewPtrTextFiller(
		" ",
		2,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by leftMargin := TextFieldSpecFiller{}.NewTextFiller().\n"+
			"Error =\n%v\n",
			ePrefix.String(),
			err.Error())

		return
	}

	var rightMargin *TextFieldSpecFiller

	rightMargin,
		err = TextFieldSpecFiller{}.NewPtrTextFiller(
		" ",
		2,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by rightMargin = TextFieldSpecFiller{}.NewTextFiller().\n"+
			"Error =\n%v\n",
			ePrefix.String(),
			err.Error())

		return
	}

	labelTxt := "Hello World!"

	expectedFinalTxt := "  " + labelTxt + "  \\n"

	var centerLabel *TextFieldSpecLabel

	centerLabel,
		err = TextFieldSpecLabel{}.NewPtrTextLabel(
		labelTxt,
		-1,
		TxtJustify.Left(),
		ePrefix)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by centerLabel = TextFieldSpecLabel{}.NewTextLabel().\n"+
			"Error =\n%v\n",
			ePrefix.String(),
			err.Error())

		return
	}

	_,
		err = stdLine.AddTextField(
		leftMargin,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by stdLine.AddTextField(leftMargin).\n"+
			"Error =\n%v\n",
			ePrefix.String(),
			err.Error())

		return
	}

	_,
		err = stdLine.AddTextField(
		centerLabel,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by stdLine.AddTextField(centerLabel).\n"+
			"Error =\n%v\n",
			ePrefix.String(),
			err.Error())

		return
	}

	_,
		err = stdLine.AddTextField(
		rightMargin,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by stdLine.AddTextField(rightMargin).\n"+
			"Error =\n%v\n",
			ePrefix.String(),
			err.Error())

		return
	}

	err = stdLine.IsValidInstanceError(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by stdLine.IsValidInstanceError().\n"+
			"Error =\n%v\n",
			ePrefix.String(),
			err.Error())

		return
	}

	stdLineTwo := TextLineSpecStandardLine{}

	err = stdLineTwo.CopyIn(
		stdLine,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by stdLineTwo.CopyIn(stdLine).\n"+
			"Error =\n%v\n",
			ePrefix.String(),
			err.Error())

		return
	}

	lenTxtFields := stdLineTwo.GetNumOfTextFields()

	if lenTxtFields == 0 {
		t.Errorf("%v - ERROR\n"+
			"stdLineTwo.CopyIn(stdLine) Failed To Copy Text Fields!\n"+
			"Length Of Text Fields is Zero!\n",
			ePrefix.String())

		return

	}

	err = stdLineTwo.IsValidInstanceError(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by stdLineTwo.IsValidInstanceError().\n"+
			"Error =\n%v\n",
			ePrefix.String(),
			err.Error())

		return
	}

	if !stdLineTwo.Equal(stdLine) {
		t.Errorf("%v\n"+
			"Error: Expected stdLineTwo.Equal(stdLine) == 'true'.\n"+
			"Instead, stdLineTwo.Equal(stdLine) == 'false'!\n",
			ePrefix.String())

		return
	}

	var rawOutput string

	rawOutput,
		err = stdLine.GetFormattedText(
		ePrefix.XCpy("stdLine->rawOutput"))

	if err != nil {
		t.Errorf("%v",
			err.Error())

		return
	}

	sMech := StrMech{}

	actualStr := sMech.ConvertNonPrintableChars(
		[]rune(rawOutput),
		false)

	if expectedFinalTxt != actualStr {
		t.Errorf("%v\n"+
			"Error: Expected stdLine final text output = '%v'.\n"+
			"Instead, stdLine final text output        = '%v'!\n",
			ePrefix.String(),
			expectedFinalTxt,
			actualStr)

		return
	}

	rawOutput,
		err = stdLineTwo.GetFormattedText(
		"stdLineTwo->rawOutput")

	if err != nil {
		t.Errorf("%v",
			err.Error())

		return
	}

	actualStr = sMech.ConvertNonPrintableChars([]rune(rawOutput), false)

	if expectedFinalTxt != actualStr {
		t.Errorf("%v\n"+
			"Error: Expected stdLineTwo final text output = '%v'.\n"+
			"Instead, stdLineTwo final text output        = '%v'!\n",
			ePrefix.String(),
			expectedFinalTxt,
			actualStr)

		return
	}

	stdLineThree := TextLineSpecStandardLine{}.New()

	err = stdLineThree.CopyIn(
		&stdLineTwo,
		textLineSpecStandardLineElectron{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from TextLineSpecStandardLine{}."+
			"CopyIn()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecStandardLine_CopyOut_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_CopyIn_000100()",
		"")

	stdLine01 := TextLineSpecStandardLine{}.NewPtr()

	rightMarginLen := 5

	rightMarginSpec,
		err := TextFieldSpecSpacer{}.NewSpacer(
		rightMarginLen,
		ePrefix.XCpy(
			"rightMarginSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var leftMarginSpec TextFieldSpecSpacer

	leftMarginLen := 6

	leftMarginSpec,
		err = TextFieldSpecSpacer{}.NewSpacer(
		leftMarginLen,
		ePrefix.XCpy(
			"leftMarginSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	label := "How Now Brown Cow!"
	fieldLen := len(label) + 4
	txtJustify := TxtJustify.Center()

	var labelSpec TextFieldSpecLabel

	labelSpec,
		err = TextFieldSpecLabel{}.NewTextLabel(
		label,
		fieldLen,
		txtJustify,
		ePrefix.XCpy(
			"labelSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	expectedStdLineText :=
		strings.Repeat(" ", leftMarginLen) +
			"  " + label + "  " +
			strings.Repeat(" ", rightMarginLen) +
			"\n"

	var indexId int

	indexId,
		err = stdLine01.AddTextField(
		&leftMarginSpec,
		ePrefix.XCpy(
			"stdLine01<-leftMarginSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if indexId != 0 {
		t.Errorf("%v - ERROR\n"+
			"stdLine01.AddTextField() should have\n"+
			"returned 'indexId' = 0\n"+
			"HOWEVER, indexId = %v\n",
			ePrefix.String(),
			indexId)

		return
	}

	collectionCount :=
		stdLine01.GetNumOfTextFields()

	if collectionCount != 1 {
		t.Errorf("%v - ERROR\n"+
			"Test #1\n"+
			"stdLine01.GetNumOfTextFields() should have\n"+
			"returned 'collectionCount' = 1\n"+
			"HOWEVER, 'collectionCount' = %v\n",
			ePrefix.String(),
			collectionCount)

		return
	}

	indexId,
		err = stdLine01.AddTextField(
		&labelSpec,
		ePrefix.XCpy(
			"stdLine01<-labelSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if indexId != 1 {
		t.Errorf("%v - ERROR\n"+
			"Test #2\n"+
			"stdLine01.AddTextField() should have\n"+
			"returned 'indexId' = 1\n"+
			"HOWEVER, indexId = %v\n",
			ePrefix.String(),
			indexId)

		return
	}

	collectionCount =
		stdLine01.GetNumOfTextFields()

	if collectionCount != 2 {
		t.Errorf("%v - ERROR\n"+
			"Test #2\n"+
			"stdLine01.GetNumOfTextFields() should have\n"+
			"returned 'collectionCount' = 2\n"+
			"HOWEVER, 'collectionCount' = %v\n",
			ePrefix.String(),
			collectionCount)

		return
	}

	indexId,
		err = stdLine01.AddTextField(
		&rightMarginSpec,
		ePrefix.XCpy(
			"stdLine01<-rightMarginSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if indexId != 2 {
		t.Errorf("%v - ERROR\n"+
			"Test #3\n"+
			"stdLine01.AddTextField() should have\n"+
			"returned 'indexId' = 2\n"+
			"HOWEVER, indexId = %v\n",
			ePrefix.String(),
			indexId)

		return
	}

	collectionCount =
		stdLine01.GetNumOfTextFields()

	if collectionCount != 3 {
		t.Errorf("%v - ERROR\n"+
			"Test #3\n"+
			"stdLine01.GetNumOfTextFields() should have\n"+
			"returned 'collectionCount' = 3\n"+
			"HOWEVER, 'collectionCount' = %v\n",
			ePrefix.String(),
			collectionCount)

		return
	}

	err = stdLine01.IsValidInstanceError(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedStdLineText),
			true)

	var actualStdLineText string

	actualStdLineText,
		err = stdLine01.GetFormattedText(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(actualStdLineText),
			true)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #1"+
			"Expected Formatted Text String 01 DOES NOT match\n"+
			"Actual Formatted Text String 01.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	var stdLine02 TextLineSpecStandardLine

	stdLine02,
		err = stdLine01.CopyOut(
		ePrefix.XCpy(
			"stdLine01->stdLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = stdLine02.IsValidInstanceError(
		ePrefix.XCpy(
			"stdLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !stdLine02.Equal(stdLine01) {
		t.Errorf("%v - ERROR\n"+
			"stdLine02.Equal(stdLine01)\n"+
			"Expected that stdLine02 == stdLine01\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!\n",
			ePrefix.String())

		return
	}

	actualStdLineText,
		err = stdLine02.GetFormattedText(
		ePrefix.XCpy(
			"stdLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	printableActualStr =
		sMech.ConvertNonPrintableChars(
			[]rune(actualStdLineText),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #1"+
			"Expected Formatted Text String 01 DOES NOT match\n"+
			"Actual Formatted Text String 01.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	_,
		err = stdLine01.CopyOut(
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from stdLine01{}."+
			"CopyOut()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	stdLine03 := TextLineSpecStandardLine{}

	_,
		err = stdLine03.CopyOut(
		ePrefix.XCpy(
			"stdLine03 is empty"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from stdLine03{}."+
			"CopyOut()\n"+
			"because 'stdLine03' is empty.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecStandardLine_CopyOutITextLine_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_CopyOutITextLine_000100()",
		"")

	stdLine01 := TextLineSpecStandardLine{}.NewPtr()

	rightMarginLen := 5

	rightMarginSpec,
		err := TextFieldSpecSpacer{}.NewSpacer(
		rightMarginLen,
		ePrefix.XCpy(
			"rightMarginSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var leftMarginSpec TextFieldSpecSpacer

	leftMarginLen := 6

	leftMarginSpec,
		err = TextFieldSpecSpacer{}.NewSpacer(
		leftMarginLen,
		ePrefix.XCpy(
			"leftMarginSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	label := "How Now Brown Cow!"
	fieldLen := len(label) + 4
	txtJustify := TxtJustify.Center()

	var labelSpec TextFieldSpecLabel

	labelSpec,
		err = TextFieldSpecLabel{}.NewTextLabel(
		label,
		fieldLen,
		txtJustify,
		ePrefix.XCpy(
			"labelSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	expectedStdLineText :=
		strings.Repeat(" ", leftMarginLen) +
			"  " + label + "  " +
			strings.Repeat(" ", rightMarginLen) +
			"\n"

	_,
		err = stdLine01.AddTextField(
		&leftMarginSpec,
		ePrefix.XCpy(
			"stdLine01<-leftMarginSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	_,
		err = stdLine01.AddTextField(
		&labelSpec,
		ePrefix.XCpy(
			"stdLine01<-labelSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	_,
		err = stdLine01.AddTextField(
		&rightMarginSpec,
		ePrefix.XCpy(
			"stdLine01<-rightMarginSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = stdLine01.IsValidInstanceError(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var iTextLine ITextLineSpecification

	iTextLine, err =
		stdLine01.CopyOutITextLine(
			ePrefix.XCpy(
				"stdLine01->iTextLine"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	stdLine02, ok := iTextLine.(*TextLineSpecStandardLine)

	if !ok {
		t.Errorf("%v\n"+
			"Error: iTextLine.(*TextLineSpecStandardLine)\n"+
			"Could not convert 'iTextLine' to TextLineSpecStandardLine\n",
			ePrefix.String())

		return
	}

	err = stdLine02.IsValidInstanceError(
		ePrefix.XCpy(
			"stdLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !stdLine02.Equal(stdLine01) {
		t.Errorf("%v\n"+
			"Error: Expected stdLine02 == stdLine01\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedStdLineText),
			true)

	var actualStdLineText string

	actualStdLineText,
		err = stdLine01.GetFormattedText(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(actualStdLineText),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #1 - stdLine01"+
			"Expected Formatted Text String 01 DOES NOT match\n"+
			"Actual Formatted Text String 01.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	actualStdLineText,
		err = stdLine02.GetFormattedText(
		ePrefix.XCpy(
			"stdLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	printableActualStr =
		sMech.ConvertNonPrintableChars(
			[]rune(actualStdLineText),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #2 - stdLine02"+
			"Expected Formatted Text String 02 DOES NOT match\n"+
			"Actual Formatted Text String 02.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	_,
		err =
		stdLine02.CopyOutITextLine(
			StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from stdLine02{}."+
			"CopyOutITextLine()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	stdLine03 := TextLineSpecStandardLine{}

	_,
		err = stdLine03.CopyOutITextLine(
		ePrefix.XCpy(
			"stdLine03 is empty"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from stdLine03{}."+
			"CopyOutITextLine()\n"+
			"because 'stdLine03' is empty.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecStandardLine_CopyOutPtr_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_CopyOutPtr_000100()",
		"")

	stdLine01 := TextLineSpecStandardLine{}.NewPtr()

	rightMarginLen := 5

	rightMarginSpec,
		err := TextFieldSpecSpacer{}.NewSpacer(
		rightMarginLen,
		ePrefix.XCpy(
			"rightMarginSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var leftMarginSpec TextFieldSpecSpacer

	leftMarginLen := 6

	leftMarginSpec,
		err = TextFieldSpecSpacer{}.NewSpacer(
		leftMarginLen,
		ePrefix.XCpy(
			"leftMarginSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	label := "How Now Brown Cow!"
	fieldLen := len(label) + 4
	txtJustify := TxtJustify.Center()

	var labelSpec TextFieldSpecLabel

	labelSpec,
		err = TextFieldSpecLabel{}.NewTextLabel(
		label,
		fieldLen,
		txtJustify,
		ePrefix.XCpy(
			"labelSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	expectedStdLineText :=
		strings.Repeat(" ", leftMarginLen) +
			"  " + label + "  " +
			strings.Repeat(" ", rightMarginLen) +
			"\n"

	var indexId int

	indexId,
		err = stdLine01.AddTextField(
		&leftMarginSpec,
		ePrefix.XCpy(
			"stdLine01<-leftMarginSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if indexId != 0 {
		t.Errorf("%v - ERROR\n"+
			"stdLine01.AddTextField() should have\n"+
			"returned 'indexId' = 0\n"+
			"HOWEVER, indexId = %v\n",
			ePrefix.String(),
			indexId)

		return
	}

	collectionCount :=
		stdLine01.GetNumOfTextFields()

	if collectionCount != 1 {
		t.Errorf("%v - ERROR\n"+
			"Test #1\n"+
			"stdLine01.GetNumOfTextFields() should have\n"+
			"returned 'collectionCount' = 1\n"+
			"HOWEVER, 'collectionCount' = %v\n",
			ePrefix.String(),
			collectionCount)

		return
	}

	indexId,
		err = stdLine01.AddTextField(
		&labelSpec,
		ePrefix.XCpy(
			"stdLine01<-labelSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if indexId != 1 {
		t.Errorf("%v - ERROR\n"+
			"Test #2\n"+
			"stdLine01.AddTextField() should have\n"+
			"returned 'indexId' = 1\n"+
			"HOWEVER, indexId = %v\n",
			ePrefix.String(),
			indexId)

		return
	}

	collectionCount =
		stdLine01.GetNumOfTextFields()

	if collectionCount != 2 {
		t.Errorf("%v - ERROR\n"+
			"Test #2\n"+
			"stdLine01.GetNumOfTextFields() should have\n"+
			"returned 'collectionCount' = 2\n"+
			"HOWEVER, 'collectionCount' = %v\n",
			ePrefix.String(),
			collectionCount)

		return
	}

	indexId,
		err = stdLine01.AddTextField(
		&rightMarginSpec,
		ePrefix.XCpy(
			"stdLine01<-rightMarginSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if indexId != 2 {
		t.Errorf("%v - ERROR\n"+
			"Test #3\n"+
			"stdLine01.AddTextField() should have\n"+
			"returned 'indexId' = 2\n"+
			"HOWEVER, indexId = %v\n",
			ePrefix.String(),
			indexId)

		return
	}

	collectionCount =
		stdLine01.GetNumOfTextFields()

	if collectionCount != 3 {
		t.Errorf("%v - ERROR\n"+
			"Test #3\n"+
			"stdLine01.GetNumOfTextFields() should have\n"+
			"returned 'collectionCount' = 3\n"+
			"HOWEVER, 'collectionCount' = %v\n",
			ePrefix.String(),
			collectionCount)

		return
	}

	err = stdLine01.IsValidInstanceError(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedStdLineText),
			true)

	var actualStdLineText string

	actualStdLineText,
		err = stdLine01.GetFormattedText(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(actualStdLineText),
			true)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #1"+
			"Expected Formatted Text String 01 DOES NOT match\n"+
			"Actual Formatted Text String 01.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	var stdLine02 *TextLineSpecStandardLine

	stdLine02,
		err = stdLine01.CopyOutPtr(
		ePrefix.XCpy(
			"stdLine01->stdLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = stdLine02.IsValidInstanceError(
		ePrefix.XCpy(
			"stdLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !stdLine02.Equal(stdLine01) {
		t.Errorf("%v - ERROR\n"+
			"stdLine02.Equal(stdLine01)\n"+
			"Expected that stdLine02 == stdLine01\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!\n",
			ePrefix.String())

		return
	}

	actualStdLineText,
		err = stdLine02.GetFormattedText(
		ePrefix.XCpy(
			"stdLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	printableActualStr =
		sMech.ConvertNonPrintableChars(
			[]rune(actualStdLineText),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #1"+
			"Expected Formatted Text String 01 DOES NOT match\n"+
			"Actual Formatted Text String 01.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	_,
		err = stdLine01.CopyOutPtr(
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from stdLine01{}."+
			"CopyOutPtr()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	stdLine03 := TextLineSpecStandardLine{}

	_,
		err = stdLine03.CopyOutPtr(
		ePrefix.XCpy(
			"stdLine03 is empty"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from stdLine03{}."+
			"CopyOutPtr()\n"+
			"because 'stdLine03' is empty.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return

}

func TestTextLineSpecStandardLine_DeleteTextField_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_DeleteTextField_000100()",
		"")

	stdLine01 := TextLineSpecStandardLine{}.NewPtr()

	rightMarginLen := 5

	rightMarginSpec,
		err := TextFieldSpecSpacer{}.NewSpacer(
		rightMarginLen,
		ePrefix.XCpy(
			"rightMarginSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var leftMarginSpec TextFieldSpecSpacer

	leftMarginLen := 6

	leftMarginSpec,
		err = TextFieldSpecSpacer{}.NewSpacer(
		leftMarginLen,
		ePrefix.XCpy(
			"leftMarginSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	label := "How Now Brown Cow!"
	fieldLen := len(label) + 4
	txtJustify := TxtJustify.Center()

	var labelSpec TextFieldSpecLabel

	labelSpec,
		err = TextFieldSpecLabel{}.NewTextLabel(
		label,
		fieldLen,
		txtJustify,
		ePrefix.XCpy(
			"labelSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	expectedStdLineText :=
		strings.Repeat(" ", leftMarginLen) +
			"  " + label + "  " +
			strings.Repeat(" ", rightMarginLen) +
			"\n"

	var indexId int

	indexId,
		err = stdLine01.AddTextField(
		&leftMarginSpec,
		ePrefix.XCpy(
			"stdLine01<-leftMarginSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if indexId != 0 {
		t.Errorf("%v - ERROR\n"+
			"stdLine01.AddTextField() should have\n"+
			"returned 'indexId' = 0\n"+
			"HOWEVER, indexId = %v\n",
			ePrefix.String(),
			indexId)

		return
	}

	collectionCount :=
		stdLine01.GetNumOfTextFields()

	if collectionCount != 1 {
		t.Errorf("%v - ERROR\n"+
			"Test #1\n"+
			"stdLine01.GetNumOfTextFields() should have\n"+
			"returned 'collectionCount' = 1\n"+
			"HOWEVER, 'collectionCount' = %v\n",
			ePrefix.String(),
			collectionCount)

		return
	}

	indexId,
		err = stdLine01.AddTextField(
		&labelSpec,
		ePrefix.XCpy(
			"stdLine01<-labelSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if indexId != 1 {
		t.Errorf("%v - ERROR\n"+
			"Test #2\n"+
			"stdLine01.AddTextField() should have\n"+
			"returned 'indexId' = 1\n"+
			"HOWEVER, indexId = %v\n",
			ePrefix.String(),
			indexId)

		return
	}

	collectionCount =
		stdLine01.GetNumOfTextFields()

	if collectionCount != 2 {
		t.Errorf("%v - ERROR\n"+
			"Test #2\n"+
			"stdLine01.GetNumOfTextFields() should have\n"+
			"returned 'collectionCount' = 2\n"+
			"HOWEVER, 'collectionCount' = %v\n",
			ePrefix.String(),
			collectionCount)

		return
	}

	indexId,
		err = stdLine01.AddTextField(
		&rightMarginSpec,
		ePrefix.XCpy(
			"stdLine01<-rightMarginSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if indexId != 2 {
		t.Errorf("%v - ERROR\n"+
			"Test #3\n"+
			"stdLine01.AddTextField() should have\n"+
			"returned 'indexId' = 2\n"+
			"HOWEVER, indexId = %v\n",
			ePrefix.String(),
			indexId)

		return
	}

	collectionCount =
		stdLine01.GetNumOfTextFields()

	if collectionCount != 3 {
		t.Errorf("%v - ERROR\n"+
			"Test #3\n"+
			"stdLine01.GetNumOfTextFields() should have\n"+
			"returned 'collectionCount' = 3\n"+
			"HOWEVER, 'collectionCount' = %v\n",
			ePrefix.String(),
			collectionCount)

		return
	}

	err = stdLine01.IsValidInstanceError(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedStdLineText),
			true)

	var actualStdLineText string

	actualStdLineText,
		err = stdLine01.GetFormattedText(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(actualStdLineText),
			true)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #1"+
			"Expected Formatted Text String 01 DOES NOT match\n"+
			"Actual Formatted Text String 01.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	err =
		stdLine01.DeleteTextField(
			2,
			ePrefix.XCpy(
				"Delete Index 2"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err =
		stdLine01.DeleteTextField(
			0,
			ePrefix.XCpy(
				"Delete Index 0"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	expectedStdLineText =
		"  " + label + "  " +
			"\n"

	printableExpectedStr =
		sMech.ConvertNonPrintableChars(
			[]rune(expectedStdLineText),
			true)

	actualStdLineText,
		err = stdLine01.GetFormattedText(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	printableActualStr =
		sMech.ConvertNonPrintableChars(
			[]rune(actualStdLineText),
			true)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #2"+
			"Expected Formatted Text String 01 DOES NOT match\n"+
			"Actual Formatted Text String 01.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

}

func TestTextLineSpecStandardLine_DeleteTextField_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_DeleteTextField_000200()",
		"")

	stdLine01 := TextLineSpecStandardLine{}

	err := stdLine01.DeleteTextField(
		5,
		ePrefix.XCpy(
			"Empty Object, Invalid Index"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected error return from stdLine01.DeleteTextField()\n"+
			"because stdLine01 is empty and index is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	var stdLine02 TextLineSpecStandardLine

	stdLine02,
		err = createTestTextLineSpecStandardLine02(
		ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = stdLine02.DeleteTextField(
		2,
		textLineSpecStandardLineElectron{})

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from stdLine02.DeleteTextField()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

}

func TestTextLineSpecStandardLine_Empty_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_Empty_000100()",
		"")

	stdLine01 := TextLineSpecStandardLine{}

	stdLine01.Empty()

	stdLine02,
		err := createTestTextLineSpecStandardLine01(
		ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	stdLine02.Empty()

	numOfTxtFields :=
		stdLine02.GetNumOfTextFields()

	if numOfTxtFields != 0 {

		t.Errorf("\n%v - ERROR\n"+
			"Expected stdLine02.GetNumOfTextFields() to return zero.\n"+
			"However, Number of Text Fields = '%v'\n"+
			"This value persisted after a call to stdLine02.Empty()\n",
			ePrefix.String(),
			numOfTxtFields)

		return
	}

	stdLineRepetitions :=
		stdLine02.GetNumOfStdLines()

	if stdLineRepetitions != 0 {

		t.Errorf("\n%v - ERROR\n"+
			"Expected stdLine02.GetNumOfStdLines() to return zero.\n"+
			"However, Number of Standard Line Repetions is '%v'\n"+
			"This value persisted after a call to stdLine02.Empty()\n",
			ePrefix.String(),
			stdLineRepetitions)

		return
	}

	newLineChars :=
		stdLine02.GetNewLineRunes()

	if newLineChars != nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected stdLine02.GetNewLineChars() to return 'nil'.\n"+
			"However, the length of new line characters is '%v'\n"+
			"This value persisted after a call to stdLine02.Empty()\n",
			ePrefix.String(),
			len(newLineChars))

		return
	}

	return
}

func TestTextLineSpecStandardLine_EmptyTextFields_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_EmptyTextFields_000100()",
		"")

	stdLine01 := TextLineSpecStandardLine{}

	stdLine01.EmptyTextFields()

	stdLine02,
		err := createTestTextLineSpecStandardLine01(
		ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	stdLine02.EmptyTextFields()

	numOfTxtFields :=
		stdLine02.GetNumOfTextFields()

	if numOfTxtFields != 0 {

		t.Errorf("\n%v - ERROR\n"+
			"Expected stdLine02.GetNumOfTextFields() to return zero.\n"+
			"However, Number of Text Fields = '%v'\n"+
			"This value persisted after a call to stdLine02.EmptyTextFields()\n",
			ePrefix.String(),
			numOfTxtFields)

		return
	}

	stdLineRepetitions :=
		stdLine02.GetNumOfStdLines()

	if stdLineRepetitions != 1 {

		t.Errorf("\n%v - ERROR\n"+
			"Expected stdLine02.GetNumOfStdLines() to return '1'.\n"+
			"However, Number of Standard Line Repetions is '%v'.\n"+
			"This value persisted after a call to stdLine02.EmptyTextFields()\n",
			ePrefix.String(),
			stdLineRepetitions)

		return
	}

	lenNewLineRunes := len(stdLine02.GetNewLineRunes())

	if lenNewLineRunes != 1 {

		t.Errorf("\n%v - ERROR\n"+
			"Expected stdLine02.GetNewLineRunes() to return an array length of '1'.\n"+
			"However, the length of new line characters is '%v'.\n"+
			"This value persisted after a call to stdLine02.EmptyTextFields()\n",
			ePrefix.String(),
			lenNewLineRunes)

		return
	}

	return
}
