package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"testing"
)

func TestTextLineSpecSolidLine_CopyIn_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_CopyIn_000100()",
		"")

	solidLineChars := "-"
	solidLineCharsRepeatCount := 35

	expectedSolidLineStr :=
		strings.Repeat(
			solidLineChars,
			solidLineCharsRepeatCount) +
			"\n"

	txtSolidLine01,
		err := TextLineSpecSolidLine{}.NewDefaultSolidLine(
		solidLineChars,
		solidLineCharsRepeatCount,
		ePrefix.XCpy(
			"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine01.IsValidInstanceError(
		ePrefix.XCpy("txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var txtSolidLine01FmtText, txtSolidLine02FmtText string

	txtSolidLine01FmtText,
		err = txtSolidLine01.GetFormattedText(
		ePrefix.XCpy(
			"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtSolidLine02 := TextLineSpecSolidLine{}

	err =
		txtSolidLine02.CopyIn(
			&txtSolidLine01,
			ePrefix.XCpy(
				"txtSolidLine02<-txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtSolidLine02FmtText,
		err = txtSolidLine02.GetFormattedText(
		ePrefix.XCpy(
			"txtSolidLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(txtSolidLine01FmtText),
			true)

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(txtSolidLine02FmtText),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #1"+
			"Expected Formatted Text String 01 DOES NOT match\n"+
			"Actual Formatted Text String 02.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	printableExpectedStr =
		sMech.ConvertNonPrintableChars(
			[]rune(expectedSolidLineStr),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #2"+
			"Expected Formatted Text String DOES NOT match\n"+
			"Actual Formatted Text String 02.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	if !txtSolidLine02.Equal(&txtSolidLine01) {
		t.Errorf("%v - ERROR\n"+
			"Test #3"+
			"Expected Text Solid Line Object #1 would be\n"+
			"Equal to Text Solid Line Object #2.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!!\n",
			ePrefix.String())

		return
	}

	txtSolidLine03 := TextLineSpecSolidLine{}

	err =
		txtSolidLine03.CopyIn(
			&txtSolidLine02,
			StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLine03."+
			"CopyIn()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecSolidLine_CopyOut_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_CopyOut_000100()",
		"")

	leftMargin := "  "
	rightMargin := "  "
	solidLineChars := "-"
	solidLineCharsRepeatCount := 35
	newLineChars := "\n"

	expectedSolidLineStr :=
		leftMargin +
			strings.Repeat(
				solidLineChars,
				solidLineCharsRepeatCount) +
			rightMargin +
			newLineChars

	txtSolidLine01,
		err := TextLineSpecSolidLine{}.NewFullSolidLineConfig(
		leftMargin,
		rightMargin,
		solidLineChars,
		solidLineCharsRepeatCount,
		newLineChars,
		ePrefix.XCpy(
			"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine01.IsValidInstanceError(
		ePrefix.XCpy("txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var txtSolidLine01FmtText, txtSolidLine02FmtText string

	txtSolidLine01FmtText,
		err = txtSolidLine01.GetFormattedText(
		ePrefix.XCpy(
			"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var txtSolidLine02 TextLineSpecSolidLine

	txtSolidLine02,
		err =
		txtSolidLine01.CopyOut(
			ePrefix.XCpy(
				"txtSolidLine02<-txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtSolidLine02FmtText,
		err = txtSolidLine02.GetFormattedText(
		ePrefix.XCpy(
			"txtSolidLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(txtSolidLine01FmtText),
			true)

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(txtSolidLine02FmtText),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #1"+
			"Expected Formatted Text String 01 DOES NOT match\n"+
			"Actual Formatted Text String 02.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	printableExpectedStr =
		sMech.ConvertNonPrintableChars(
			[]rune(expectedSolidLineStr),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #2"+
			"Expected Formatted Text String DOES NOT match\n"+
			"Actual Formatted Text String 02.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	if !txtSolidLine02.Equal(&txtSolidLine01) {
		t.Errorf("%v - ERROR\n"+
			"Test #3"+
			"Expected Text Solid Line Object #1 would be\n"+
			"Equal to Text Solid Line Object #2.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!!\n",
			ePrefix.String())

		return
	}

	txtSolidLine03 := TextLineSpecSolidLine{}

	_,
		err =
		txtSolidLine03.CopyOut(
			StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLine03."+
			"CopyOut()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	_,
		err =
		txtSolidLine01.CopyOut(
			StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLine01."+
			"CopyOut()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	txtSolidLine04 := TextLineSpecSolidLine{}

	_,
		err =
		txtSolidLine04.CopyOut(
			ePrefix.XCpy(
				"txtSolidLine04"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLine04."+
			"CopyOut()\n"+
			"because 'txtSolidLine04' is empty.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecSolidLine_CopyOutITextLine_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_CopyOutITextLine_000100()",
		"")

	leftMargin := "  "
	rightMargin := "  "
	solidLineChars := "-"
	solidLineCharsRepeatCount := 35
	newLineChars := "\n\n"

	expectedSolidLineStr :=
		leftMargin +
			strings.Repeat(
				solidLineChars,
				solidLineCharsRepeatCount) +
			rightMargin +
			newLineChars

	txtSolidLine01,
		err := TextLineSpecSolidLine{}.
		NewFullSolidLineRunesConfig(
			[]rune(leftMargin),
			[]rune(rightMargin),
			[]rune(solidLineChars),
			solidLineCharsRepeatCount,
			[]rune(newLineChars),
			ePrefix.XCpy(
				"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine01.IsValidInstanceError(
		ePrefix.XCpy("txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var iTextLine ITextLineSpecification

	iTextLine, err =
		txtSolidLine01.CopyOutITextLine(
			ePrefix.XCpy(
				"txtSolidLine01->iTextLine"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtSolidLine02, ok := iTextLine.(*TextLineSpecSolidLine)

	if !ok {
		t.Errorf("%v\n"+
			"Error: iTextLine.(*TextLineSpecSolidLine)\n"+
			"Could not convert 'iTextLine' to TextLineSpecSolidLine\n",
			ePrefix.String())

		return
	}

	if !txtSolidLine02.Equal(&txtSolidLine01) {
		t.Errorf("%v\n"+
			"Error: Expected txtSolidLine02 == txtSolidLine01\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return
	}

	var txtSolidLine02FmtText string

	txtSolidLine02FmtText,
		err = txtSolidLine02.GetFormattedText(
		ePrefix.XCpy(
			"txtSolidLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedSolidLineStr),
			true)

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(txtSolidLine02FmtText),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #1"+
			"Expected Formatted Text String 01 DOES NOT match\n"+
			"Actual Formatted Text String 02.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	if !txtSolidLine01.EqualITextLine(iTextLine) {
		t.Errorf("%v\n"+
			"Error: txtSolidLine01.EqualITextLine(iTextLine)\n"+
			"Expected txtSolidLine01 == iTextLine\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return
	}

	txtSolidLine03 := TextLineSpecSolidLine{}

	_,
		err =
		txtSolidLine03.CopyOutITextLine(
			ePrefix.XCpy(
				"txtSolidLine03->_"))

	if err == nil {
		t.Errorf("%v\n"+
			"Error: txtSolidLine03.CopyOutITextLine()\n"+
			"Expected an error return because 'txtSolidLine03' is empty!\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	var txtSolidLine04 *TextLineSpecSolidLine

	txtSolidLine04,
		err = TextLineSpecSolidLine{}.
		NewPtrFullSolidLineRunesConfig(
			[]rune(leftMargin),
			[]rune(rightMargin),
			[]rune(solidLineChars),
			solidLineCharsRepeatCount,
			[]rune(newLineChars),
			ePrefix.XCpy(
				"txtSolidLine04"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine04.IsValidInstanceError(
		ePrefix.XCpy("txtSolidLine04"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	_,
		err =
		txtSolidLine04.CopyOutITextLine(
			StrMech{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLine04."+
			"CopyOutITextLine()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecSolidLine_CopyOutPtr_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_CopyOutPtr_000100()",
		"")

	leftMargin := "  "
	rightMargin := "  "
	solidLineChars := "-"
	solidLineCharsRepeatCount := 35
	newLineChars := "\n"

	expectedSolidLineStr :=
		leftMargin +
			strings.Repeat(
				solidLineChars,
				solidLineCharsRepeatCount) +
			rightMargin +
			newLineChars

	txtSolidLine01,
		err := TextLineSpecSolidLine{}.NewSolidLine(
		leftMargin,
		rightMargin,
		solidLineChars,
		solidLineCharsRepeatCount,
		ePrefix.XCpy(
			"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine01.IsValidInstanceError(
		ePrefix.XCpy("txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var txtSolidLine01FmtText, txtSolidLine02FmtText string

	txtSolidLine01FmtText,
		err = txtSolidLine01.GetFormattedText(
		ePrefix.XCpy(
			"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var txtSolidLine02 *TextLineSpecSolidLine

	txtSolidLine02,
		err =
		txtSolidLine01.CopyOutPtr(
			ePrefix.XCpy(
				"txtSolidLine02<-txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtSolidLine02FmtText,
		err = txtSolidLine02.GetFormattedText(
		ePrefix.XCpy(
			"txtSolidLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(txtSolidLine01FmtText),
			true)

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(txtSolidLine02FmtText),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #1\n"+
			"Expected Formatted Text String 01 DOES NOT match\n"+
			"Actual Formatted Text String 02.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	printableExpectedStr =
		sMech.ConvertNonPrintableChars(
			[]rune(expectedSolidLineStr),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #2\n"+
			"Expected Formatted Text String DOES NOT match\n"+
			"Actual Formatted Text String 02.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	if !txtSolidLine02.Equal(&txtSolidLine01) {
		t.Errorf("%v - ERROR\n"+
			"Test #3\n"+
			"Expected Text Solid Line Object #1 would be\n"+
			"Equal to Text Solid Line Object #2.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!!\n",
			ePrefix.String())

		return
	}

	txtSolidLine03 := TextLineSpecSolidLine{}

	_,
		err =
		txtSolidLine03.CopyOutPtr(
			StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLine03."+
			"CopyOutPtr()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	_,
		err =
		txtSolidLine01.CopyOutPtr(
			StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLine01."+
			"CopyOutPtr()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	txtSolidLine04 := TextLineSpecSolidLine{}

	_,
		err =
		txtSolidLine04.CopyOutPtr(
			ePrefix.XCpy(
				"txtSolidLine04"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLine04."+
			"CopyOutPtr()\n"+
			"because 'txtSolidLine04' is empty.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecSolidLine_Empty_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_CopyOut_000100()",
		"")

	leftMargin := "  "
	rightMargin := "  "
	solidLineChars := "-"
	solidLineCharsRepeatCount := 35

	txtSolidLine01,
		err := TextLineSpecSolidLine{}.NewPtrSolidLine(
		leftMargin,
		rightMargin,
		solidLineChars,
		solidLineCharsRepeatCount,
		ePrefix.XCpy(
			"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine01.IsValidInstanceError(
		ePrefix.XCpy("txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtSolidLine01.Empty()

	if len(txtSolidLine01.leftMarginChars) != 0 {

		t.Errorf("%v - ERROR\n"+
			"txtSolidLine01.Empty() should have set all\n"+
			"member variables to their zero values.\n"+
			"However 'txtSolidLine01.leftMarginChars' is NOT equal to zero!\n"+
			"txtSolidLine01.leftMargin= '%v'\n",
			ePrefix.String(),
			len(txtSolidLine01.leftMarginChars))

		return
	}

	if len(txtSolidLine01.rightMarginChars) != 0 {

		t.Errorf("%v - ERROR\n"+
			"txtSolidLine01.Empty() should have set all\n"+
			"member variables to their zero values.\n"+
			"However 'txtSolidLine01.rightMarginChars' is NOT equal to zero!\n"+
			"txtSolidLine01.rightMarginChars length = '%v'\n",
			ePrefix.String(),
			len(txtSolidLine01.rightMarginChars))

		return
	}

	if txtSolidLine01.solidLineChars != nil {

		t.Errorf("%v - ERROR\n"+
			"txtSolidLine01.Empty() should have set all\n"+
			"member variables to their zero values.\n"+
			"However 'txtSolidLine01.solidLineChars' is NOT equal to 'nil'!\n"+
			"txtSolidLine01.solidLineChars= '%v'\n",
			ePrefix.String(),
			txtSolidLine01.solidLineChars)

		return
	}

	if txtSolidLine01.solidLineCharsRepeatCount != 0 {

		t.Errorf("%v - ERROR\n"+
			"txtSolidLine01.Empty() should have set all\n"+
			"member variables to their zero values.\n"+
			"However 'txtSolidLine01.solidLineCharsRepeatCount' is NOT equal to zero!\n"+
			"txtSolidLine01.solidLineCharsRepeatCount= '%v'\n",
			ePrefix.String(),
			txtSolidLine01.solidLineCharsRepeatCount)

		return
	}

	if txtSolidLine01.newLineChars != nil {

		t.Errorf("%v - ERROR\n"+
			"txtSolidLine01.Empty() should have set all\n"+
			"member variables to their zero values.\n"+
			"However 'txtSolidLine01.newLineChars' is NOT equal to 'nil'!\n"+
			"txtSolidLine01.newLineChars= '%v'\n",
			ePrefix.String(),
			txtSolidLine01.newLineChars)

		return
	}

	if txtSolidLine01.turnLineTerminatorOff != false {

		t.Errorf("%v - ERROR\n"+
			"txtSolidLine01.Empty() should have set all\n"+
			"member variables to their zero values.\n"+
			"However 'txtSolidLine01.turnLineTerminatorOff' is NOT equal to 'false'!\n"+
			"txtSolidLine01.turnLineTerminatorOff = '%v'\n",
			ePrefix.String(),
			txtSolidLine01.turnLineTerminatorOff)

	}

	txtSolidLine02 := TextLineSpecSolidLine{}

	txtSolidLine02.Empty()

	return
}

func TestTextLineSpecSolidLine_Equal_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_CopyOut_000100()",
		"")

	leftMargin := "  "
	rightMargin := "  "
	solidLineChars := "-"
	solidLineCharsRepeatCount := 35
	newLineChars := "\n-\n"

	expectedSolidLineStr :=
		leftMargin +
			strings.Repeat(
				solidLineChars,
				solidLineCharsRepeatCount) +
			rightMargin +
			newLineChars

	txtSolidLine01 := TextLineSpecSolidLine{}

	err :=
		txtSolidLine01.SetFullSolidLineConfig(
			leftMargin,
			rightMargin,
			solidLineChars,
			solidLineCharsRepeatCount,
			newLineChars,
			ePrefix.XCpy(
				"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !txtSolidLine01.IsValidInstance() {

		t.Errorf("%v - ERROR\n" +
			"Initial creation of instance 'txtSolidLine01' is invalid!!!\n" +
			ePrefix.String())

		return
	}

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var txtSolidLine02 TextLineSpecSolidLine

	txtSolidLine02,
		err = TextLineSpecSolidLine{}.NewFullSolidLineConfig(
		leftMargin,
		rightMargin,
		solidLineChars,
		solidLineCharsRepeatCount,
		newLineChars,
		ePrefix.XCpy(
			"txtSolidLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine02.IsValidInstanceError(
		ePrefix.XCpy("txtSolidLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !txtSolidLine01.Equal(&txtSolidLine02) {

		t.Errorf("%v - ERROR\n"+
			"Expected txtSolidLine01==txtSolidLine02\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!\n",
			ePrefix.String())

		return
	}

	var txtSolidLine02FmtText string

	txtSolidLine02FmtText,
		err = txtSolidLine01.GetFormattedText(
		ePrefix.XCpy(
			"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedSolidLineStr),
			true)

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(txtSolidLine02FmtText),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #1\n"+
			"Expected Formatted Text String 01 DOES NOT match\n"+
			"Actual Formatted Text String 02.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	txtSolidLine01.solidLineChars[0] =
		'X'

	if txtSolidLine01.Equal(&txtSolidLine02) {

		t.Errorf("%v - ERROR\n"+
			"Test #1\n"+
			"Expected txtSolidLine01 WOULD NOT EQUAL txtSolidLine02\n"+
			"because solidLineChars are different.\n"+
			"HOWEVER, THEY ARE EQUAL!!\n",
			ePrefix.String())

		return
	}

	err = txtSolidLine01.CopyIn(
		&txtSolidLine02,
		ePrefix.XCpy(
			"txtSolidLine01<-txtSolidLine02 #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtSolidLine01.leftMarginChars = nil

	if txtSolidLine01.Equal(&txtSolidLine02) {

		t.Errorf("%v - ERROR\n"+
			"Test #2\n"+
			"Expected txtSolidLine01 WOULD NOT EQUAL txtSolidLine02\n"+
			"because leftMarginChars is different.\n"+
			"HOWEVER, THEY ARE EQUAL!!\n",
			ePrefix.String())

		return
	}

	err = txtSolidLine01.CopyIn(
		&txtSolidLine02,
		ePrefix.XCpy(
			"txtSolidLine01<-txtSolidLine02 #2"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtSolidLine01.rightMarginChars = nil

	if txtSolidLine01.Equal(&txtSolidLine02) {

		t.Errorf("%v - ERROR\n"+
			"Test #3\n"+
			"Expected txtSolidLine01 WOULD NOT EQUAL txtSolidLine02\n"+
			"because rightMarginChars is different.\n"+
			"HOWEVER, THEY ARE EQUAL!!\n",
			ePrefix.String())

		return
	}

	err = txtSolidLine01.CopyIn(
		&txtSolidLine02,
		ePrefix.XCpy(
			"txtSolidLine01<-txtSolidLine02 #3"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtSolidLine01.solidLineCharsRepeatCount = 0

	if txtSolidLine01.Equal(&txtSolidLine02) {

		t.Errorf("%v - ERROR\n"+
			"Test #4\n"+
			"Expected txtSolidLine01 WOULD NOT EQUAL txtSolidLine02\n"+
			"because solidLineCharsRepeatCount is different.\n"+
			"HOWEVER, THEY ARE EQUAL!!\n",
			ePrefix.String())

		return
	}

	err = txtSolidLine01.CopyIn(
		&txtSolidLine02,
		ePrefix.XCpy(
			"txtSolidLine01<-txtSolidLine02 #3"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtSolidLine01.newLineChars = []rune{'X', '-', '\n'}

	if txtSolidLine01.Equal(&txtSolidLine02) {

		t.Errorf("%v - ERROR\n"+
			"Test #5\n"+
			"Expected txtSolidLine01 WOULD NOT EQUAL txtSolidLine02\n"+
			"because newLineChars are different.\n"+
			"HOWEVER, THEY ARE EQUAL!!\n",
			ePrefix.String())

		return
	}

	txtSolidLine03 := TextLineSpecSolidLine{}

	txtSolidLine04 := TextLineSpecSolidLine{}

	if !txtSolidLine03.Equal(&txtSolidLine04) {

		t.Errorf("%v - ERROR\n"+
			"Expected txtSolidLine03==txtSolidLine04"+
			"because they are both empty.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecSolidLine_EqualITextLine_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_EqualITextLine_000100()",
		"")

	leftMargin := "  "
	rightMargin := "  "
	solidLineChars := "-"
	solidLineCharsRepeatCount := 35
	newLineChars := "\n\n"

	expectedSolidLineStr :=
		leftMargin +
			strings.Repeat(
				solidLineChars,
				solidLineCharsRepeatCount) +
			rightMargin +
			newLineChars

	txtSolidLine01 := TextLineSpecSolidLine{}

	err :=
		txtSolidLine01.SetFullSolidLineRunesConfig(
			[]rune(leftMargin),
			[]rune(rightMargin),
			[]rune(solidLineChars),
			solidLineCharsRepeatCount,
			[]rune(newLineChars),
			ePrefix.XCpy(
				"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine01.IsValidInstanceError(
		ePrefix.XCpy("txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	standardLineSpec := TextLineSpecStandardLine{}.NewPtr()

	areEqual :=
		txtSolidLine01.EqualITextLine(
			standardLineSpec)

	if areEqual == true {

		t.Errorf("%v\n"+
			"Error: txtSolidLine01.EqualITextLine()\n"+
			"Expected areEqual == 'false' because input\n"+
			"parameter 'standardLineSpec' is the wrong type.\n"+
			"HOWEVER, areEqual == 'true'!!\n",
			ePrefix.String())

		return
	}

	var txtSolidLine02 TextLineSpecSolidLine

	txtSolidLine02,
		err = txtSolidLine01.CopyOut(
		ePrefix.XCpy(
			"txtSolidLine01->txtSolidLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !txtSolidLine01.Equal(&txtSolidLine02) {

		t.Errorf("%v - ERROR\n"+
			"Expected txtSolidLine01 == txtSolidLine02.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedSolidLineStr),
			true)

	var actualFmtTxtStr02 string

	actualFmtTxtStr02,
		err =
		txtSolidLine02.GetFormattedText(
			ePrefix.XCpy(
				"txtSolidLine02->actualFmtTxtStr"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(actualFmtTxtStr02),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Expected Formatted Text String DOES NOT match\n"+
			"Actual Formatted Text String.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	areEqual =
		txtSolidLine01.EqualITextLine(
			&txtSolidLine02)

	if !areEqual {

		t.Errorf("%v - ERROR\n"+
			"txtSolidLine01.EqualITextLine(&txtSolidLine02)"+
			"Expected areEqual == 'true' because the two"+
			"instances are identical.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return
	}

	txtSolidLine03 := TextLineSpecSolidLine{}

	areEqual =
		txtSolidLine03.EqualITextLine(
			&txtSolidLine02)

	if areEqual == true {

		t.Errorf("%v\n"+
			"Error: txtSolidLine03.EqualITextLine(&txtSolidLine02)\n"+
			"Expected areEqual == 'false' because instance\n"+
			"'txtSolidLine03' is empty.\n"+
			"HOWEVER, areEqual == 'true'!!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecSolidLine_GetLeftMarginLength_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_GetLeftMarginLength_000100()",
		"")

	expectedLeftMarginLength := 2
	leftMargin := "  "
	rightMargin := "  "
	solidLineChars := "-"
	solidLineCharsRepeatCount := 35

	txtSolidLine01 := TextLineSpecSolidLine{}

	err :=
		txtSolidLine01.SetSolidLine(
			leftMargin,
			rightMargin,
			solidLineChars,
			solidLineCharsRepeatCount,
			ePrefix.XCpy(
				"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualLeftMarginLength :=
		txtSolidLine01.GetLeftMarginLength()

	if expectedLeftMarginLength != actualLeftMarginLength {

		t.Errorf("%v - ERROR\n"+
			"txtSolidLine01.GetLeftMarginLength()\n"+
			"Expected Left Margin Length DOES NOT match\n"+
			"Actual Left Margin Length.\n"+
			"Expected Left Margin Length = '%v'\n"+
			"Instead, Left Margin Length = '%v'\n",
			ePrefix.String(),
			expectedLeftMarginLength,
			actualLeftMarginLength)

		return
	}

	expectedLeftMarginLength = 0

	txtSolidLine02 := TextLineSpecSolidLine{}

	actualLeftMarginLength =
		txtSolidLine02.GetLeftMarginLength()

	if expectedLeftMarginLength != actualLeftMarginLength {

		t.Errorf("%v - ERROR\n"+
			"txtSolidLine02.GetLeftMarginLength()\n"+
			"Note: 'txtSolidLine02' is empty!\n"+
			"Expected Left Margin Length DOES NOT match\n"+
			"Actual Left Margin Length.\n"+
			"Expected Left Margin Length = '%v'\n"+
			"Instead, Left Margin Length = '%v'\n",
			ePrefix.String(),
			expectedLeftMarginLength,
			actualLeftMarginLength)

		return
	}

	return
}

func TestTextLineSpecSolidLine_GetRightMarginLength_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_GetRightMarginLength_000100()",
		"")

	leftMargin := "  "
	rightMargin := "  "
	expectedRightMarginLength := 2
	solidLineChars := "-"
	solidLineCharsRepeatCount := 35

	txtSolidLine01 := TextLineSpecSolidLine{}

	err :=
		txtSolidLine01.SetSolidLine(
			leftMargin,
			rightMargin,
			solidLineChars,
			solidLineCharsRepeatCount,
			ePrefix.XCpy(
				"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualRightMarginLength :=
		txtSolidLine01.GetRightMarginLength()

	if expectedRightMarginLength != actualRightMarginLength {

		t.Errorf("%v - ERROR\n"+
			"txtSolidLine01.GetLeftMarginLength()\n"+
			"Expected Right Margin Length DOES NOT match\n"+
			"Actual Right Margin Length.\n"+
			"Expected Right Margin Length = '%v'\n"+
			"Instead, Right Margin Length = '%v'\n",
			ePrefix.String(),
			expectedRightMarginLength,
			actualRightMarginLength)

		return
	}

	expectedRightMarginLength = 0

	txtSolidLine02 := TextLineSpecSolidLine{}

	actualRightMarginLength =
		txtSolidLine02.GetRightMarginLength()

	if expectedRightMarginLength != actualRightMarginLength {

		t.Errorf("%v - ERROR\n"+
			"txtSolidLine02.GetLeftMarginLength()\n"+
			"Note: 'txtSolidLine02' is empty!\n"+
			"Expected Right Margin Length DOES NOT match\n"+
			"Actual Right Margin Length.\n"+
			"Expected Right Margin Length = '%v'\n"+
			"Instead, Right Margin Length = '%v'\n",
			ePrefix.String(),
			expectedRightMarginLength,
			actualRightMarginLength)

		return
	}

	return
}

func TestTextLineSpecSolidLine_GetFormattedText_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_GetFormattedText_000100()",
		"")

	leftMargin := "  "
	rightMargin := "  "
	solidLineChars := "-"
	solidLineCharsRepeatCount := 35
	newLineChars := "\n"

	expectedSolidLineStr :=
		leftMargin +
			strings.Repeat(
				solidLineChars,
				solidLineCharsRepeatCount) +
			rightMargin +
			newLineChars

	txtSolidLine01 := TextLineSpecSolidLine{}

	err :=
		txtSolidLine01.SetSolidLine(
			leftMargin,
			rightMargin,
			solidLineChars,
			solidLineCharsRepeatCount,
			ePrefix.XCpy(
				"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine01.IsValidInstanceError(
		ePrefix.XCpy("txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedSolidLineStr),
			true)

	var actualFmtTxtStr01 string

	actualFmtTxtStr01,
		err =
		txtSolidLine01.GetFormattedText(
			ePrefix.XCpy(
				"txtSolidLine02->actualFmtTxtStr"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(actualFmtTxtStr01),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"txtSolidLine01.GetFormattedText()\n"+
			"Expected Formatted Text String DOES NOT match\n"+
			"Actual Formatted Text String.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	_,
		err =
		txtSolidLine01.GetFormattedText(
			StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLine01."+
			"GetFormattedText()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	txtSolidLine02 := TextLineSpecSolidLine{}

	_,
		err = txtSolidLine02.GetFormattedText(
		ePrefix.XCpy(
			"txtSolidLine02==empty"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLine02."+
			"CopyOut()\n"+
			"because 'txtSolidLine02' is empty.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecSolidLine_GetNewLineChars_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_GetNewLineChars_000100()",
		"")

	txtSolidLine01 := TextLineSpecSolidLine{}

	var newLineChars string

	newLineChars = txtSolidLine01.GetNewLineChars()

	if len(newLineChars) != 0 {

		t.Errorf("%v - ERROR\n"+
			"txtSolidLine01.GetNewLineChars()\n"+
			"Expected New Line Characters length == Zero"+
			"because 'txtSolidLine01' is empty.\n"+
			"HOWEVER, New Line Characters length = '%v'\n"+
			"New Line Characters = '%v'\n",
			ePrefix.String(),
			len(newLineChars),
			newLineChars)

		return
	}

	leftMargin := "  "
	rightMargin := "  "
	solidLineChars := "-"
	solidLineCharsRepeatCount := 35
	expectedNewLineChars := "\n-\n"

	txtSolidLine02,
		err := TextLineSpecSolidLine{}.
		NewFullSolidLineRunesConfig(
			[]rune(leftMargin),
			[]rune(rightMargin),
			[]rune(solidLineChars),
			solidLineCharsRepeatCount,
			[]rune(expectedNewLineChars),
			ePrefix.XCpy(
				"txtSolidLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualNewLineChars :=
		txtSolidLine02.GetNewLineChars()

	if expectedNewLineChars != actualNewLineChars {

		t.Errorf("%v - ERROR\n"+
			"txtSolidLine02.GetNewLineChars()\n"+
			"Expected New Line Characters = '%v'\n"+
			"Instead, New Line Characters = '%v'\n"+
			ePrefix.String(),
			expectedNewLineChars,
			actualNewLineChars)

		return
	}

	return
}

func TestTextLineSpecSolidLine_GetNewLineRunes_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_GetNewLineRunes_000100()",
		"")

	txtSolidLine01 := TextLineSpecSolidLine{}

	var newLineRunes []rune

	newLineRunes = txtSolidLine01.GetNewLineRunes()

	if len(newLineRunes) != 0 {

		t.Errorf("%v - ERROR\n"+
			"txtSolidLine01.GetNewLineRunes()\n"+
			"Expected New Line Runes length == Zero"+
			"because 'txtSolidLine01' is empty.\n"+
			"HOWEVER, New Line Runes length = '%v'\n"+
			"New Line Runes = '%v'\n",
			ePrefix.String(),
			len(newLineRunes),
			newLineRunes)

		return
	}

	leftMargin := "  "
	rightMargin := "  "
	solidLineChars := "-"
	solidLineCharsRepeatCount := 35
	expectedNewLineRunes := []rune{'\n', '-', '\n'}

	txtSolidLine02,
		err := TextLineSpecSolidLine{}.
		NewFullSolidLineRunesConfig(
			[]rune(leftMargin),
			[]rune(rightMargin),
			[]rune(solidLineChars),
			solidLineCharsRepeatCount,
			expectedNewLineRunes,
			ePrefix.XCpy(
				"txtSolidLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualNewLineRunes :=
		txtSolidLine02.GetNewLineRunes()

	sMechPreon := strMechPreon{}

	areEqual :=
		sMechPreon.equalRuneArrays(
			expectedNewLineRunes,
			actualNewLineRunes)

	if !areEqual {

		t.Errorf("%v - ERROR\n"+
			"txtSolidLine02.GetNewLineRunes()\n"+
			"Expected New Line Runes = '%v'\n"+
			"Instead, New Line Runes = '%v'\n"+
			"Expected New Line Runs Length = %v\n"+
			" Actual New Line Runes Length = %v\n",
			ePrefix.String(),
			string(expectedNewLineRunes),
			string(actualNewLineRunes),
			len(expectedNewLineRunes),
			len(actualNewLineRunes))

		return
	}

	return
}

func TestTextLineSpecSolidLine_GetSolidLineChars_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_GetSolidLineChars_000100()",
		"")

	txtSolidLine01 := TextLineSpecSolidLine{}

	var solidLineChars string

	solidLineChars = txtSolidLine01.GetSolidLineChars()

	if len(solidLineChars) != 0 {

		t.Errorf("%v - ERROR\n"+
			"txtSolidLine01.GetNewLineChars()\n"+
			"Expected Solid Line Characters length == Zero"+
			"because 'txtSolidLine01' is empty.\n"+
			"HOWEVER, Actual Solid Line Characters length = '%v'\n"+
			"Actual Solid Line Characters = '%v'\n",
			ePrefix.String(),
			len(solidLineChars),
			solidLineChars)

		return
	}

	leftMargin := "  "
	rightMargin := "  "
	solidLineChars = "-"
	solidLineCharsRepeatCount := 35
	newLineChars := "\n-\n"

	txtSolidLine02,
		err := TextLineSpecSolidLine{}.
		NewFullSolidLineRunesConfig(
			[]rune(leftMargin),
			[]rune(rightMargin),
			[]rune(solidLineChars),
			solidLineCharsRepeatCount,
			[]rune(newLineChars),
			ePrefix.XCpy(
				"txtSolidLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualSolidLineChars :=
		txtSolidLine02.GetSolidLineChars()

	if solidLineChars != actualSolidLineChars {

		t.Errorf("%v - ERROR\n"+
			"txtSolidLine02.GetSolidLineChars()\n"+
			"Expected Solid Line Characters = '%v'\n"+
			"Instead, Solid Line Characters = '%v'\n"+
			ePrefix.String(),
			solidLineChars,
			actualSolidLineChars)

		return
	}

	return
}

func TestTextLineSpecSolidLine_GetSolidLineCharsRepeatCount_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_GetSolidLineRunes_000100()",
		"")

	txtSolidLine01 := TextLineSpecSolidLine{}

	var expectedSolidLineCharsRepeatCount int

	expectedSolidLineCharsRepeatCount =
		txtSolidLine01.GetSolidLineCharsRepeatCount()

	if expectedSolidLineCharsRepeatCount != 0 {

		t.Errorf("%v - ERROR\n"+
			"txtSolidLine01.GetSolidLineCharsRepeatCount()\n"+
			"Expected Solid Line Characters Repeat Count == Zero"+
			"because 'txtSolidLine01' is empty.\n"+
			"HOWEVER, Actual Solid Line Characters Repeat Count = '%v'\n",
			ePrefix.String(),
			expectedSolidLineCharsRepeatCount)

		return
	}

	leftMargin := "  "
	rightMargin := "  "
	solidLineChars := "-"
	expectedSolidLineCharsRepeatCount = 35
	newLineChars := "\n-\n"

	txtSolidLine02,
		err := TextLineSpecSolidLine{}.
		NewFullSolidLineRunesConfig(
			[]rune(leftMargin),
			[]rune(rightMargin),
			[]rune(solidLineChars),
			expectedSolidLineCharsRepeatCount,
			[]rune(newLineChars),
			ePrefix.XCpy(
				"txtSolidLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualSolidLineCharsRepeatCount :=
		txtSolidLine02.GetSolidLineCharsRepeatCount()

	if expectedSolidLineCharsRepeatCount !=
		actualSolidLineCharsRepeatCount {

		t.Errorf("%v - ERROR\n"+
			"txtSolidLine02.GetSolidLineChars()\n"+
			"Expected Solid Line Characters Repeat Count = '%v'\n"+
			"  Actual Solid Line Characters Repeat Count = '%v'\n"+
			ePrefix.String(),
			expectedSolidLineCharsRepeatCount,
			actualSolidLineCharsRepeatCount)

		return
	}

	return
}

func TestTextLineSpecSolidLine_GetSolidLineRunes_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_GetSolidLineRunes_000100()",
		"")

	txtSolidLine01 := TextLineSpecSolidLine{}

	var expectedSolidLineRunes []rune

	expectedSolidLineRunes = txtSolidLine01.GetSolidLineRunes()

	if len(expectedSolidLineRunes) != 0 {

		t.Errorf("%v - ERROR\n"+
			"txtSolidLine01.GetSolidLineRunes()\n"+
			"Expected Solid Line Runes length == Zero"+
			"because 'txtSolidLine01' is empty.\n"+
			"HOWEVER, Solid Line Runes length = '%v'\n"+
			"Solid Line Runes = '%v'\n",
			ePrefix.String(),
			len(expectedSolidLineRunes),
			expectedSolidLineRunes)

		return
	}

	leftMargin := []rune{' ', ' '}
	rightMargin := []rune{' ', ' '}
	expectedSolidLineRunes = []rune{'-', '*'}
	solidLineCharsRepeatCount := 35
	expectedNewLineRunes := []rune{'\n', '-', '\n'}

	txtSolidLine02,
		err := TextLineSpecSolidLine{}.
		NewFullSolidLineRunesConfig(
			leftMargin,
			rightMargin,
			expectedSolidLineRunes,
			solidLineCharsRepeatCount,
			expectedNewLineRunes,
			ePrefix.XCpy(
				"txtSolidLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualSolidLineRunes :=
		txtSolidLine02.GetSolidLineRunes()

	sMechPreon := strMechPreon{}

	areEqual :=
		sMechPreon.equalRuneArrays(
			expectedSolidLineRunes,
			actualSolidLineRunes)

	if !areEqual {

		t.Errorf("%v - ERROR\n"+
			"txtSolidLine02.GetNewLineRunes()\n"+
			"Expected Solid Line Runes = '%v'\n"+
			"Instead, Solid Line Runes = '%v'\n"+
			"Expected Solid Line Runs Length = %v\n"+
			" Actual Solid Line Runes Length = %v\n",
			ePrefix.String(),
			string(expectedSolidLineRunes),
			string(actualSolidLineRunes),
			len(expectedSolidLineRunes),
			len(actualSolidLineRunes))

		return
	}

	return
}

func TestTextLineSpecSolidLine_GetTurnLineTerminatorOff_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_GetTurnLineTerminatorOff_000100()",
		"")

	txtSolidLine01 := TextLineSpecSolidLine{}

	if txtSolidLine01.GetTurnLineTerminatorOff() == true {

		t.Errorf("%v - ERROR\n"+
			"txtSolidLine01.GetTurnLineTerminatorOff() == 'true'\n"+
			"Expected turnLineTerminatorOff == 'false'\n"+
			"because 'txtSolidLine01' is empty.\n"+
			"Instead, turnLineTerminatorOff == 'true'\n",
			ePrefix.String())

		return
	}

	leftMargin := "  "
	rightMargin := "  "
	solidLineChars := "-"
	solidLineCharsRepeatCount := 35
	newLineChars := "\n\n\n"
	var turnAutoLineTerminationOff bool

	expectedSolidLineStr :=
		leftMargin +
			strings.Repeat(
				solidLineChars,
				solidLineCharsRepeatCount) +
			rightMargin +
			newLineChars

	turnAutoLineTerminationOff = false

	txtSolidLine02,
		err := TextLineSpecSolidLine{}.NewSolidLineAllParms(
		leftMargin,
		rightMargin,
		solidLineChars,
		solidLineCharsRepeatCount,
		newLineChars,
		turnAutoLineTerminationOff,
		ePrefix.XCpy(
			"txtSolidLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine02.IsValidInstanceError(
		ePrefix.XCpy("txtSolidLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedSolidLineStr),
			true)

	txtSolidLine02FmtText :=
		txtSolidLine02.String()

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(txtSolidLine02FmtText),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #1 - txtSolidLine02.String()\n"+
			"Expected Formatted Text String 01 DOES NOT match\n"+
			"Actual Formatted Text String 01.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	if txtSolidLine02.GetTurnLineTerminatorOff() == true {

		t.Errorf("%v - ERROR\n"+
			"Test #1\n"+
			"txtSolidLine02.GetTurnLineTerminatorOff() == 'true'\n"+
			"Expected turnLineTerminatorOff == 'false'\n"+
			"Instead, turnLineTerminatorOff == 'true'\n",
			ePrefix.String())

		return
	}

	txtSolidLine02.TurnAutoLineTerminationOff()

	if txtSolidLine02.GetTurnLineTerminatorOff() == false {

		t.Errorf("%v - ERROR\n"+
			"Test #2 \n"+
			"txtSolidLine02.TurnAutoLineTerminationOff()\n"+
			"Expected turnLineTerminatorOff == 'true'\n"+
			"Instead, turnLineTerminatorOff == 'false'\n",
			ePrefix.String())

		return
	}

	txtSolidLine02.TurnAutoLineTerminationOn()

	if txtSolidLine02.GetTurnLineTerminatorOff() == true {

		t.Errorf("%v - ERROR\n"+
			"Test #3 \n"+
			"txtSolidLine02.TurnAutoLineTerminationOn()\n"+
			"Expected turnLineTerminatorOff == 'false'\n"+
			"Instead, turnLineTerminatorOff == 'true'\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecSolidLine_IsValidInstance_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_IsValidInstance_000100()",
		"")

	txtSolidLine01 := TextLineSpecSolidLine{}

	if txtSolidLine01.IsValidInstance() {

		t.Errorf("%v - ERROR\n"+
			"txtSolidLine01.IsValidInstance() returned a \n"+
			"value of 'true'. It should have returned a\n"+
			"value of 'false' because 'txtSolidLine01' is invalid!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecSolidLine_IsValidInstanceError_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_IsValidInstanceError_000100()",
		"")

	txtSolidLine01 := TextLineSpecSolidLine{}

	err :=
		txtSolidLine01.IsValidInstanceError(
			ePrefix.XCpy(
				"txtSolidLine01"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLine01."+
			"IsValidInstanceError()\n"+
			"because 'txtSolidLine01' is empty.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	solidLineChars := "-"
	solidLineCharsRepeatCount := 35

	var txtSolidLine02 TextLineSpecSolidLine

	txtSolidLine02,
		err =
		txtSolidLine02.NewDefaultSolidLine(
			solidLineChars,
			solidLineCharsRepeatCount,
			ePrefix.XCpy(
				"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine02.IsValidInstanceError(
		ePrefix.XCpy("txtSolidLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine01.IsValidInstanceError(
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLine01."+
			"IsValidInstanceError()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecSolidLine_NewDefaultSolidLine_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_NewDefaultSolidLine_000100()",
		"")

	solidLineChars := "-"
	solidLineCharsRepeatCount := 35

	_,
		err := TextLineSpecSolidLine{}.NewDefaultSolidLine(
		solidLineChars,
		solidLineCharsRepeatCount,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from TextLineSpecSolidLine{}."+
			"NewDefaultSolidLine()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecSolidLine_NewFullSolidLineConfig_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_NewFullSolidLineConfig_000100()",
		"")

	leftMargin := "  "
	rightMargin := "  "
	solidLineChars := "-"
	solidLineCharsRepeatCount := 35
	newLineChars := "\n\n\n"

	expectedSolidLineStr :=
		leftMargin +
			strings.Repeat(
				solidLineChars,
				solidLineCharsRepeatCount) +
			rightMargin +
			newLineChars

	txtSolidLine01,
		err := TextLineSpecSolidLine{}.NewFullSolidLineConfig(
		leftMargin,
		rightMargin,
		solidLineChars,
		solidLineCharsRepeatCount,
		newLineChars,
		ePrefix.XCpy(
			"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine01.IsValidInstanceError(
		ePrefix.XCpy("txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedSolidLineStr),
			true)

	txtSolidLine01FmtText :=
		txtSolidLine01.String()

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(txtSolidLine01FmtText),
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
		err = TextLineSpecSolidLine{}.NewFullSolidLineConfig(
		leftMargin,
		rightMargin,
		solidLineChars,
		solidLineCharsRepeatCount,
		newLineChars,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from TextLineSpecSolidLine{}."+
			"NewFullSolidLineConfig()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecSolidLine_NewFullSolidLineRunesConfig_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_NewFullSolidLineRunesConfig_000100()",
		"")

	leftMargin := "  "
	rightMargin := "  "
	solidLineChars := "-!"
	solidLineCharsRepeatCount := 35
	newLineChars := "\n\n\n"

	expectedSolidLineStr :=
		leftMargin +
			strings.Repeat(
				solidLineChars,
				solidLineCharsRepeatCount) +
			rightMargin +
			newLineChars

	txtSolidLine01,
		err := TextLineSpecSolidLine{}.
		NewFullSolidLineRunesConfig(
			[]rune(leftMargin),
			[]rune(rightMargin),
			[]rune(solidLineChars),
			solidLineCharsRepeatCount,
			[]rune(newLineChars),
			ePrefix.XCpy(
				"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine01.IsValidInstanceError(
		ePrefix.XCpy("txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedSolidLineStr),
			true)

	txtSolidLine01FmtText :=
		txtSolidLine01.String()

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(txtSolidLine01FmtText),
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
		err = TextLineSpecSolidLine{}.
		NewFullSolidLineRunesConfig(
			[]rune(leftMargin),
			[]rune(rightMargin),
			[]rune(solidLineChars),
			solidLineCharsRepeatCount,
			[]rune(newLineChars),
			StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from TextLineSpecSolidLine{}."+
			"NewFullSolidLineRunesConfig()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecSolidLine_NewSolidLine_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_NewSolidLine_000100()",
		"")

	leftMargin := "  "
	rightMargin := "  "
	solidLineChars := "-"
	solidLineCharsRepeatCount := 35
	newLineChars := "\n"

	expectedSolidLineStr :=
		leftMargin +
			strings.Repeat(
				solidLineChars,
				solidLineCharsRepeatCount) +
			rightMargin +
			newLineChars

	txtSolidLine01,
		err := TextLineSpecSolidLine{}.NewSolidLine(
		leftMargin,
		rightMargin,
		solidLineChars,
		solidLineCharsRepeatCount,
		ePrefix.XCpy(
			"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine01.IsValidInstanceError(
		ePrefix.XCpy("txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedSolidLineStr),
			true)

	txtSolidLine01FmtText :=
		txtSolidLine01.String()

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(txtSolidLine01FmtText),
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
		err = TextLineSpecSolidLine{}.NewSolidLine(
		leftMargin,
		rightMargin,
		solidLineChars,
		solidLineCharsRepeatCount,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from TextLineSpecSolidLine{}."+
			"NewSolidLine()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecSolidLine_NewSolidLineAllParms_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_NewSolidLineAllParms_000100()",
		"")

	leftMargin := "  "
	rightMargin := "  "
	solidLineChars := "-"
	solidLineCharsRepeatCount := 35
	newLineChars := "\n\n\n"
	var turnAutoLineTerminationOff bool

	expectedSolidLineStr :=
		leftMargin +
			strings.Repeat(
				solidLineChars,
				solidLineCharsRepeatCount) +
			rightMargin +
			newLineChars

	turnAutoLineTerminationOff = false

	txtSolidLine01,
		err := TextLineSpecSolidLine{}.NewSolidLineAllParms(
		leftMargin,
		rightMargin,
		solidLineChars,
		solidLineCharsRepeatCount,
		newLineChars,
		turnAutoLineTerminationOff,
		ePrefix.XCpy(
			"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine01.IsValidInstanceError(
		ePrefix.XCpy("txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if txtSolidLine01.GetTurnLineTerminatorOff() == true {

		t.Errorf("%v - ERROR\n"+
			"txtSolidLine01.GetTurnLineTerminatorOff() == 'true'\n"+
			"Expected turnLineTerminatorOff == 'false'\n"+
			"Instead, turnLineTerminatorOff == 'true'\n",
			ePrefix.String())

		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedSolidLineStr),
			true)

	txtSolidLine01FmtText :=
		txtSolidLine01.String()

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(txtSolidLine01FmtText),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #1\n"+
			"Expected Formatted Text String 01 DOES NOT match\n"+
			"Actual Formatted Text String 01.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	var txtSolidLine02 TextLineSpecSolidLine

	turnAutoLineTerminationOff = true

	expectedSolidLineStr =
		leftMargin +
			strings.Repeat(
				solidLineChars,
				solidLineCharsRepeatCount) +
			rightMargin

	txtSolidLine02,
		err = TextLineSpecSolidLine{}.NewSolidLineAllParms(
		leftMargin,
		rightMargin,
		solidLineChars,
		solidLineCharsRepeatCount,
		newLineChars,
		turnAutoLineTerminationOff,
		ePrefix.XCpy(
			"txtSolidLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine02.IsValidInstanceError(
		ePrefix.XCpy("txtSolidLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if txtSolidLine02.GetTurnLineTerminatorOff() == false {

		t.Errorf("%v - ERROR\n"+
			"txtSolidLine02.GetTurnLineTerminatorOff() == 'false'\n"+
			"Expected turnLineTerminatorOff == 'true'\n"+
			"Instead, turnLineTerminatorOff == 'false'\n",
			ePrefix.String())

		return
	}

	printableExpectedStr =
		sMech.ConvertNonPrintableChars(
			[]rune(expectedSolidLineStr),
			true)

	txtSolidLine02FmtText :=
		txtSolidLine02.String()

	printableActualStr =
		sMech.ConvertNonPrintableChars(
			[]rune(txtSolidLine02FmtText),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #2 - txtSolidLine02.String()\n"+
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
		err = TextLineSpecSolidLine{}.NewSolidLineAllParms(
		leftMargin,
		rightMargin,
		solidLineChars,
		solidLineCharsRepeatCount,
		newLineChars,
		turnAutoLineTerminationOff,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from TextLineSpecSolidLine{}."+
			"NewSolidLineAllParms()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecSolidLine_NewPtrDefaultSolidLine_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_NewPtrDefaultSolidLine_000100()",
		"")

	solidLineChars := "-"
	solidLineCharsRepeatCount := 35

	expectedSolidLineStr :=
		strings.Repeat(
			solidLineChars,
			solidLineCharsRepeatCount) +
			"\n"

	txtSolidLine01,
		err := TextLineSpecSolidLine{}.NewPtrDefaultSolidLine(
		solidLineChars,
		solidLineCharsRepeatCount,
		ePrefix.XCpy(
			"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine01.IsValidInstanceError(
		ePrefix.XCpy("txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedSolidLineStr),
			true)

	txtSolidLine01FmtText :=
		txtSolidLine01.String()

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(txtSolidLine01FmtText),
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
		err = TextLineSpecSolidLine{}.NewPtrDefaultSolidLine(
		solidLineChars,
		solidLineCharsRepeatCount,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from TextLineSpecSolidLine{}."+
			"NewPtrDefaultSolidLine()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecSolidLine_NewPtrFullSolidLineConfig_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_NewPtrFullSolidLineConfig_000100()",
		"")

	leftMargin := "  "
	rightMargin := "  "
	solidLineChars := "-"
	solidLineCharsRepeatCount := 35
	newLineChars := "\n"

	expectedSolidLineStr :=
		leftMargin +
			strings.Repeat(
				solidLineChars,
				solidLineCharsRepeatCount) +
			rightMargin +
			newLineChars

	txtSolidLine01,
		err := TextLineSpecSolidLine{}.NewPtrFullSolidLineConfig(
		leftMargin,
		rightMargin,
		solidLineChars,
		solidLineCharsRepeatCount,
		newLineChars,
		ePrefix.XCpy(
			"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine01.IsValidInstanceError(
		ePrefix.XCpy("txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedSolidLineStr),
			true)

	txtSolidLine01FmtText :=
		txtSolidLine01.String()

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(txtSolidLine01FmtText),
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
		err = TextLineSpecSolidLine{}.NewPtrFullSolidLineConfig(
		leftMargin,
		rightMargin,
		solidLineChars,
		solidLineCharsRepeatCount,
		newLineChars,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from TextLineSpecSolidLine{}."+
			"NewPtrFullSolidLineConfig()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecSolidLine_NewPtrFullSolidLineRunesConfig_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_NewPtrFullSolidLineRunesConfig_000100()",
		"")

	leftMargin := "  "
	rightMargin := "  "
	solidLineChars := "-"
	solidLineCharsRepeatCount := 35
	newLineChars := "\n\n"

	expectedSolidLineStr :=
		leftMargin +
			strings.Repeat(
				solidLineChars,
				solidLineCharsRepeatCount) +
			rightMargin +
			newLineChars

	txtSolidLine01,
		err := TextLineSpecSolidLine{}.
		NewPtrFullSolidLineRunesConfig(
			[]rune(leftMargin),
			[]rune(rightMargin),
			[]rune(solidLineChars),
			solidLineCharsRepeatCount,
			[]rune(newLineChars),
			ePrefix.XCpy(
				"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine01.IsValidInstanceError(
		ePrefix.XCpy("txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedSolidLineStr),
			true)

	txtSolidLine01FmtText :=
		txtSolidLine01.String()

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(txtSolidLine01FmtText),
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
		err = TextLineSpecSolidLine{}.
		NewPtrFullSolidLineRunesConfig(
			[]rune(leftMargin),
			[]rune(rightMargin),
			[]rune(solidLineChars),
			solidLineCharsRepeatCount,
			[]rune(newLineChars),
			StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from TextLineSpecSolidLine{}."+
			"NewPtrFullSolidLineRunesConfig()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecSolidLine_NewPtrSolidLine_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_NewPtrSolidLine_000100()",
		"")

	leftMargin := "  "
	rightMargin := "  "
	solidLineChars := "-"
	solidLineCharsRepeatCount := 35
	newLineChars := "\n"

	expectedSolidLineStr :=
		leftMargin +
			strings.Repeat(
				solidLineChars,
				solidLineCharsRepeatCount) +
			rightMargin +
			newLineChars

	txtSolidLine01,
		err := TextLineSpecSolidLine{}.NewPtrSolidLine(
		leftMargin,
		rightMargin,
		solidLineChars,
		solidLineCharsRepeatCount,
		ePrefix.XCpy(
			"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine01.IsValidInstanceError(
		ePrefix.XCpy("txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedSolidLineStr),
			true)

	txtSolidLine01FmtText :=
		txtSolidLine01.String()

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(txtSolidLine01FmtText),
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
		err = TextLineSpecSolidLine{}.NewPtrSolidLine(
		leftMargin,
		rightMargin,
		solidLineChars,
		solidLineCharsRepeatCount,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from TextLineSpecSolidLine{}."+
			"NewPtrSolidLine()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecSolidLine_NewPtrSolidLineAllParms_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_NewPtrSolidLineAllParms_000100()",
		"")

	leftMargin := "  "
	rightMargin := "  "
	solidLineChars := "-"
	solidLineCharsRepeatCount := 35
	newLineChars := "\n\n\n"
	var turnAutoLineTerminationOff bool

	expectedSolidLineStr :=
		leftMargin +
			strings.Repeat(
				solidLineChars,
				solidLineCharsRepeatCount) +
			rightMargin +
			newLineChars

	turnAutoLineTerminationOff = false

	txtSolidLine01,
		err := TextLineSpecSolidLine{}.NewPtrSolidLineAllParms(
		leftMargin,
		rightMargin,
		solidLineChars,
		solidLineCharsRepeatCount,
		newLineChars,
		turnAutoLineTerminationOff,
		ePrefix.XCpy(
			"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine01.IsValidInstanceError(
		ePrefix.XCpy("txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if txtSolidLine01.GetTurnLineTerminatorOff() == true {

		t.Errorf("%v - ERROR\n"+
			"txtSolidLine01.GetTurnLineTerminatorOff() == 'true'\n"+
			"Expected turnLineTerminatorOff == 'false'\n"+
			"Instead, turnLineTerminatorOff == 'true'\n",
			ePrefix.String())

		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedSolidLineStr),
			true)

	txtSolidLine01FmtText :=
		txtSolidLine01.String()

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(txtSolidLine01FmtText),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #1\n"+
			"Expected Formatted Text String 01 DOES NOT match\n"+
			"Actual Formatted Text String 01.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	var txtSolidLine02 *TextLineSpecSolidLine

	turnAutoLineTerminationOff = true

	expectedSolidLineStr =
		leftMargin +
			strings.Repeat(
				solidLineChars,
				solidLineCharsRepeatCount) +
			rightMargin

	txtSolidLine02,
		err = TextLineSpecSolidLine{}.NewPtrSolidLineAllParms(
		leftMargin,
		rightMargin,
		solidLineChars,
		solidLineCharsRepeatCount,
		newLineChars,
		turnAutoLineTerminationOff,
		ePrefix.XCpy(
			"txtSolidLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine02.IsValidInstanceError(
		ePrefix.XCpy("txtSolidLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if txtSolidLine02.turnLineTerminatorOff == false {

		t.Errorf("%v - ERROR\n"+
			"txtSolidLine02.turnLineTerminatorOff == 'false'\n"+
			"Expected turnLineTerminatorOff == 'true'\n"+
			"Instead, turnLineTerminatorOff == 'false'\n",
			ePrefix.String())

		return
	}

	printableExpectedStr =
		sMech.ConvertNonPrintableChars(
			[]rune(expectedSolidLineStr),
			true)

	txtSolidLine02FmtText :=
		txtSolidLine02.String()

	printableActualStr =
		sMech.ConvertNonPrintableChars(
			[]rune(txtSolidLine02FmtText),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #2 - txtSolidLine02.String()\n"+
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
		err = TextLineSpecSolidLine{}.NewPtrSolidLineAllParms(
		leftMargin,
		rightMargin,
		solidLineChars,
		solidLineCharsRepeatCount,
		newLineChars,
		turnAutoLineTerminationOff,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from TextLineSpecSolidLine{}."+
			"NewPtrSolidLineAllParms()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}
