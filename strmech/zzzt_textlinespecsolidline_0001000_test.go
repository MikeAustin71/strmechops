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
		ePrefix.XCtx(
			"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine01.IsValidInstanceError(
		ePrefix.XCtx("txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var txtSolidLine01FmtText, txtSolidLine02FmtText string

	txtSolidLine01FmtText,
		err = txtSolidLine01.GetFormattedText(
		ePrefix.XCtx(
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
			ePrefix.XCtx(
				"txtSolidLine02<-txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtSolidLine02FmtText,
		err = txtSolidLine02.GetFormattedText(
		ePrefix.XCtx(
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
			ePrefix.XCtxEmpty().String(),
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
			ePrefix.XCtxEmpty().String(),
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
			ePrefix.XCtxEmpty().String())

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
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecSolidLine_CopyOut_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_CopyOut_000100()",
		"")

	leftMargin := 2
	rightMargin := 2
	solidLineChars := "-"
	solidLineCharsRepeatCount := 35
	newLineChars := "\n"

	expectedSolidLineStr :=
		strings.Repeat(" ", leftMargin) +
			strings.Repeat(
				solidLineChars,
				solidLineCharsRepeatCount) +
			strings.Repeat(" ", rightMargin) +
			newLineChars

	txtSolidLine01,
		err := TextLineSpecSolidLine{}.NewFullSolidLineConfig(
		leftMargin,
		rightMargin,
		solidLineChars,
		solidLineCharsRepeatCount,
		newLineChars,
		ePrefix.XCtx(
			"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine01.IsValidInstanceError(
		ePrefix.XCtx("txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var txtSolidLine01FmtText, txtSolidLine02FmtText string

	txtSolidLine01FmtText,
		err = txtSolidLine01.GetFormattedText(
		ePrefix.XCtx(
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
			ePrefix.XCtx(
				"txtSolidLine02<-txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtSolidLine02FmtText,
		err = txtSolidLine02.GetFormattedText(
		ePrefix.XCtx(
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
			ePrefix.XCtxEmpty().String(),
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
			ePrefix.XCtxEmpty().String(),
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
			ePrefix.XCtxEmpty().String())

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
			ePrefix.XCtxEmpty().String())

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
			ePrefix.XCtxEmpty().String())

		return
	}

	txtSolidLine04 := TextLineSpecSolidLine{}

	_,
		err =
		txtSolidLine04.CopyOut(
			ePrefix.XCtx(
				"txtSolidLine04"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLine04."+
			"CopyOut()\n"+
			"because 'txtSolidLine04' is empty.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecSolidLine_CopyOutITextLine_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_CopyOutITextLine_000100()",
		"")

	leftMargin := 2
	rightMargin := 2
	solidLineChars := "-"
	solidLineCharsRepeatCount := 35
	newLineChars := "\n\n"

	expectedSolidLineStr :=
		strings.Repeat(" ", leftMargin) +
			strings.Repeat(
				solidLineChars,
				solidLineCharsRepeatCount) +
			strings.Repeat(" ", rightMargin) +
			newLineChars

	txtSolidLine01,
		err := TextLineSpecSolidLine{}.
		NewFullSolidLineRunesConfig(
			leftMargin,
			rightMargin,
			[]rune(solidLineChars),
			solidLineCharsRepeatCount,
			[]rune(newLineChars),
			ePrefix.XCtx(
				"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine01.IsValidInstanceError(
		ePrefix.XCtx("txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var iTextLine ITextLineSpecification

	iTextLine, err =
		txtSolidLine01.CopyOutITextLine(
			ePrefix.XCtx(
				"txtSolidLine01->iTextLine"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtSolidLine02, ok := iTextLine.(*TextLineSpecSolidLine)

	if !ok {
		t.Errorf("%v\n"+
			"Error: iTextLine.(*TextLineSpecPlainText)\n"+
			"Could not convert 'iTextLine' to TextLineSpecPlainText\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	if !txtSolidLine02.Equal(&txtSolidLine01) {
		t.Errorf("%v\n"+
			"Error: Expected txtSolidLine02 == txtSolidLine01\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	var txtSolidLine02FmtText string

	txtSolidLine02FmtText,
		err = txtSolidLine02.GetFormattedText(
		ePrefix.XCtx(
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
			ePrefix.XCtxEmpty().String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	if !txtSolidLine01.EqualITextLine(iTextLine) {
		t.Errorf("%v\n"+
			"Error: txtSolidLine01.EqualITextLine(iTextLine)\n"+
			"Expected txtSolidLine01 == iTextLine\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	txtSolidLine03 := TextLineSpecSolidLine{}

	_,
		err =
		txtSolidLine03.CopyOutITextLine(
			ePrefix.XCtx(
				"txtSolidLine03->_"))

	if err == nil {
		t.Errorf("%v\n"+
			"Error: txtSolidLine03.CopyOutITextLine()\n"+
			"Expected an error return because 'txtSolidLine03' is empty!\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	var txtSolidLine04 *TextLineSpecSolidLine

	txtSolidLine04,
		err = TextLineSpecSolidLine{}.
		NewPtrFullSolidLineRunesConfig(
			leftMargin,
			rightMargin,
			[]rune(solidLineChars),
			solidLineCharsRepeatCount,
			[]rune(newLineChars),
			ePrefix.XCtx(
				"txtSolidLine04"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine04.IsValidInstanceError(
		ePrefix.XCtx("txtSolidLine04"))

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
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecSolidLine_CopyOutPtr_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_CopyOutPtr_000100()",
		"")

	leftMargin := 2
	rightMargin := 2
	solidLineChars := "-"
	solidLineCharsRepeatCount := 35
	newLineChars := "\n"

	expectedSolidLineStr :=
		strings.Repeat(" ", leftMargin) +
			strings.Repeat(
				solidLineChars,
				solidLineCharsRepeatCount) +
			strings.Repeat(" ", rightMargin) +
			newLineChars

	txtSolidLine01,
		err := TextLineSpecSolidLine{}.NewSolidLine(
		leftMargin,
		rightMargin,
		solidLineChars,
		solidLineCharsRepeatCount,
		ePrefix.XCtx(
			"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine01.IsValidInstanceError(
		ePrefix.XCtx("txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var txtSolidLine01FmtText, txtSolidLine02FmtText string

	txtSolidLine01FmtText,
		err = txtSolidLine01.GetFormattedText(
		ePrefix.XCtx(
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
			ePrefix.XCtx(
				"txtSolidLine02<-txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtSolidLine02FmtText,
		err = txtSolidLine02.GetFormattedText(
		ePrefix.XCtx(
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
			ePrefix.XCtxEmpty().String(),
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
			ePrefix.XCtxEmpty().String(),
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
			ePrefix.XCtxEmpty().String())

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
			ePrefix.XCtxEmpty().String())

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
			ePrefix.XCtxEmpty().String())

		return
	}

	txtSolidLine04 := TextLineSpecSolidLine{}

	_,
		err =
		txtSolidLine04.CopyOutPtr(
			ePrefix.XCtx(
				"txtSolidLine04"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLine04."+
			"CopyOutPtr()\n"+
			"because 'txtSolidLine04' is empty.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecSolidLine_Empty_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_CopyOut_000100()",
		"")

	leftMargin := 2
	rightMargin := 2
	solidLineChars := "-"
	solidLineCharsRepeatCount := 35

	txtSolidLine01,
		err := TextLineSpecSolidLine{}.NewPtrSolidLine(
		leftMargin,
		rightMargin,
		solidLineChars,
		solidLineCharsRepeatCount,
		ePrefix.XCtx(
			"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine01.IsValidInstanceError(
		ePrefix.XCtx("txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtSolidLine01.Empty()

	if txtSolidLine01.leftMargin != 0 {

		t.Errorf("%v - ERROR\n"+
			"txtSolidLine01.Empty() should have set all\n"+
			"member variables to their zero values.\n"+
			"However 'txtSolidLine01.leftMargin' is NOT equal to zero!\n"+
			"txtSolidLine01.leftMargin= '%v'\n",
			ePrefix.String(),
			txtSolidLine01.leftMargin)

		return
	}

	if txtSolidLine01.rightMargin != 0 {

		t.Errorf("%v - ERROR\n"+
			"txtSolidLine01.Empty() should have set all\n"+
			"member variables to their zero values.\n"+
			"However 'txtSolidLine01.rightMargin' is NOT equal to zero!\n"+
			"txtSolidLine01.rightMargin= '%v'\n",
			ePrefix.String(),
			txtSolidLine01.rightMargin)

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

	txtSolidLine02 := TextLineSpecSolidLine{}

	txtSolidLine02.Empty()

	return
}
