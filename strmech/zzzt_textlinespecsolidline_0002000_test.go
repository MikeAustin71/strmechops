package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"testing"
)

func TestTextLineSpecSolidLine_copyIn_000100(t *testing.T) {

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
		err := TextLineSpecSolidLine{}.NewPtrDefaultSolidLine(
		solidLineChars,
		solidLineCharsRepeatCount,
		ePrefix.XCtx(
			"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtSolidLine02 := TextLineSpecSolidLine{}

	txtSolidLineMolecule := textLineSpecSolidLineMolecule{}

	err =
		txtSolidLineMolecule.copyIn(
			&txtSolidLine02,
			txtSolidLine01,
			ePrefix.XCtx(
				"txtSolidLine02<-txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var txtSolidLine01FmtText, txtSolidLine02FmtText string

	txtSolidLine01FmtText,
		err = txtSolidLine01.GetFormattedText(
		ePrefix.XCtx(
			"txtSolidLine01->txtSolidLine01FmtText"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtSolidLine02FmtText,
		err = txtSolidLine02.GetFormattedText(
		ePrefix.XCtx(
			"txtSolidLine02->txtSolidLine02FmtText"))

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

	if !txtSolidLine02.Equal(txtSolidLine01) {
		t.Errorf("%v - ERROR\n"+
			"Test #3"+
			"Expected Text Solid Line Object #1 would be\n"+
			"Equal to Text Solid Line Object #2.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	txtSolidLine01.solidLineChars = nil

	txtSolidLine03 := TextLineSpecSolidLine{}

	err =
		txtSolidLineMolecule.copyIn(
			&txtSolidLine03,
			txtSolidLine01,
			ePrefix.XCtx(
				"txtSolidLine02<-txtSolidLine01"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLineMolecule."+
			"copyIn()\n"+
			"because 'txtSolidLine01.solidLineChars' has a 'nil' value.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	err =
		txtSolidLineMolecule.copyIn(
			nil,
			&txtSolidLine02,
			ePrefix.XCtx(
				"txtSolidLine02<-txtSolidLine01"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLineMolecule."+
			"copyIn()\n"+
			"because 'targetTxtSolidLine' has a 'nil' value.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	err =
		txtSolidLineMolecule.copyIn(
			&txtSolidLine03,
			nil,
			ePrefix.XCtx(
				"txtSolidLine02<-txtSolidLine01"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLineMolecule."+
			"copyIn()\n"+
			"because 'incomingTxtSolidLine' has a 'nil' value.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecSolidLine_copyOut_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_copyOut_000100()",
		"")

	leftMargin := 2
	rightMargin := 2
	solidLineChars := "-"
	solidLineCharsRepeatCount := 35
	newLineChars := "\n"

	expectedSolidLineStr :=
		strings.Repeat(" ", 2) +
			strings.Repeat(
				solidLineChars,
				solidLineCharsRepeatCount) +
			strings.Repeat(" ", rightMargin) +
			newLineChars

	txtSolidLine01,
		err := TextLineSpecSolidLine{}.NewPtrFullSolidLineConfig(
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

	var txtSolidLine02 TextLineSpecSolidLine

	txtSolidLineMolecule := textLineSpecSolidLineMolecule{}

	txtSolidLine02,
		err =
		txtSolidLineMolecule.copyOut(
			txtSolidLine01,
			ePrefix.XCtx(
				"txtSolidLine02<-txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var txtSolidLine01FmtText, txtSolidLine02FmtText string

	txtSolidLine01FmtText,
		err = txtSolidLine01.GetFormattedText(
		ePrefix.XCtx(
			"txtSolidLine01->txtSolidLine01FmtText"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtSolidLine02FmtText,
		err = txtSolidLine02.GetFormattedText(
		ePrefix.XCtx(
			"txtSolidLine02->txtSolidLine02FmtText"))

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

	if !txtSolidLine02.Equal(txtSolidLine01) {
		t.Errorf("%v - ERROR\n"+
			"Test #3"+
			"Expected Text Solid Line Object #1 would be\n"+
			"Equal to Text Solid Line Object #2.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	txtSolidLine01.solidLineChars = nil

	_,
		err =
		txtSolidLineMolecule.copyOut(
			txtSolidLine01,
			ePrefix.XCtx(
				"txtSolidLine01 invalid test #1"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLineMolecule."+
			"copyOut()\n"+
			"because 'txtSolidLine01.solidLineChars' has a 'nil' value.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	_,
		err =
		txtSolidLineMolecule.copyOut(
			nil,
			ePrefix.XCtx(
				"txtSolidLine == nil"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLineMolecule."+
			"copyIn()\n"+
			"because 'targetTxtSolidLine' has a 'nil' value.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecSolidLine_empty_000100(t *testing.T) {

	txtSolidLineAtom := textLineSpecSolidLineAtom{}

	txtSolidLineAtom.empty(
		nil)

}

func TestTextLineSpecSolidLine_getFormattedText_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_copyOut_000100()",
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
		NewPtrFullSolidLineRunesConfig(
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

	var txtSolidLine01FmtText string

	txtSolidLineMolecule := textLineSpecSolidLineMolecule{}

	txtSolidLine01FmtText,
		err = txtSolidLineMolecule.getFormattedText(
		txtSolidLine01,
		ePrefix.XCtx(
			"txtSolidLine01->txtSolidLine01FmtText"))

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
			[]rune(txtSolidLine01FmtText),
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

	txtSolidLineMolecule2 := textLineSpecSolidLineMolecule{}

	_,
		err =
		txtSolidLineMolecule2.getFormattedText(
			nil,
			ePrefix.XCtx(
				"txtSolidLine=='nil'"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLineMolecule2."+
			"getFormattedText()\n"+
			"because 'txtSolidLine' has a 'nil' value.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	txtSolidLine01.rightMargin = 1000001

	_,
		err =
		txtSolidLineMolecule2.getFormattedText(
			txtSolidLine01,
			ePrefix.XCtx(
				"txtSolidLine01 is invalid"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLineMolecule2."+
			"getFormattedText()\n"+
			"because 'txtSolidLine01.rightMargin = 1000001'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}
