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

func TestTextLineSpecSolidLine_equal_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_equal_000100()",
		"")

	leftMargin := 2
	rightMargin := 2
	solidLineChars := "-"
	solidLineCharsRepeatCount := 35
	newLineChars := "\n-\n"

	txtSolidLine01 := TextLineSpecSolidLine{}

	err :=
		txtSolidLine01.SetFullSolidLineConfig(
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

	var txtSolidLine02 TextLineSpecSolidLine

	txtSolidLine02,
		err = txtSolidLine01.CopyOut(
		ePrefix.XCtx(
			"txtSolidLine01->txtSolidLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtSolidLineAtom := textLineSpecSolidLineAtom{}

	areEqual :=
		txtSolidLineAtom.equal(
			&txtSolidLine01,
			&txtSolidLine02)

	if !areEqual {

		t.Errorf("%v - ERROR\n"+
			"Test #1\n"+
			"Expected txtSolidLine01==txtSolidLine02.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	areEqual =
		txtSolidLineAtom.equal(
			nil,
			&txtSolidLine02)

	if areEqual {

		t.Errorf("%v - ERROR\n"+
			"Test #2\n"+
			"Expected areEqual WOULD BE 'false'\n"+
			"because 'txtSolidLineOne' has a 'nil' value.\n"+
			"HOWEVER, 'areEqual' is 'true'!!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	areEqual =
		txtSolidLineAtom.equal(
			&txtSolidLine01,
			nil)

	if areEqual {

		t.Errorf("%v - ERROR\n"+
			"Test #3\n"+
			"Expected areEqual WOULD BE 'false'\n"+
			"because 'txtSolidLineTwo' has a 'nil' value.\n"+
			"HOWEVER, 'areEqual' is 'true'!!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	leftMargin = 2
	rightMargin = 2
	solidLineChars = "-"
	solidLineCharsRepeatCount = 35
	newLineChars = "\n-\n"

	txtSolidLine03 := TextLineSpecSolidLine{}

	err =
		txtSolidLine03.SetFullSolidLineConfig(
			leftMargin,
			rightMargin,
			solidLineChars,
			solidLineCharsRepeatCount,
			newLineChars,
			ePrefix.XCtx(
				"txtSolidLine03"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !txtSolidLine03.IsValidInstance() {
		t.Errorf("%v - ERROR\n"+
			"txtSolidLine03.IsValidInstance()=='false'\n",
			ePrefix.XCtxEmpty().String())
	}

	var txtSolidLine04 TextLineSpecSolidLine

	txtSolidLine04,
		err = txtSolidLine03.CopyOut(
		ePrefix.XCtx(
			"txtSolidLine03->txtSolidLine04"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	areEqual =
		txtSolidLineAtom.equal(
			&txtSolidLine03,
			&txtSolidLine04)

	if !areEqual {

		t.Errorf("%v - ERROR\n"+
			"Test #4\n"+
			"Expected areEqual WOULD BE 'true'\n"+
			"because 'txtSolidLine03' and 'txtSolidLine04' are equivalent.\n"+
			"HOWEVER, 'areEqual' is 'false'!!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	txtSolidLine04.turnLineTerminatorOff = true

	areEqual =
		txtSolidLineAtom.equal(
			&txtSolidLine03,
			&txtSolidLine04)

	if areEqual {

		t.Errorf("%v - ERROR\n"+
			"Test #4\n"+
			"Expected areEqual WOULD BE 'false'\n"+
			"because 'txtSolidLine03' and 'txtSolidLine04' are NOT equal.\n"+
			"txtSolidLine04.turnLineTerminatorOff = true\n"+
			"HOWEVER, 'areEqual' is 'false'!!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecSolidLine_getFormattedText_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_getFormattedText_000100()",
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

func TestTextLineSpecSolidLine_testValidityOfTextSpecSolidLine_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_testValidityOfTextSpecSolidLine_000100()",
		"")

	leftMargin := 2
	rightMargin := 2
	solidLineChars := "-"
	solidLineCharsRepeatCount := 35
	newLineChars := "\n\n"

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

	var txtSolidLine02 TextLineSpecSolidLine

	txtSolidLine02,
		err = txtSolidLine01.CopyOut(
		ePrefix.XCtx(
			"txtSolidLine01->txtSolidLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtSolidLine02.solidLineCharsRepeatCount = -97

	var isValid bool
	txtSolidLineAtom := textLineSpecSolidLineAtom{}

	isValid,
		err = txtSolidLineAtom.testValidityOfTextSpecSolidLine(
		&txtSolidLine02,
		ePrefix.XCtx(
			"Testing 'txtSolidLine02'"))

	if isValid == true {
		t.Errorf("%v - ERROR\n"+
			"Expected 'isValid' equal 'false' because\n"+
			"txtSolidLine02.solidLineCharsRepeatCount = -97.\n"+
			"HOWEVER, 'isValid' IS 'true'!!!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from \n"+
			"txtSolidLineAtom.testValidityOfTextSpecSolidLine() \n"+
			"because txtSolidLine02.solidLineCharsRepeatCount = -97\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	var txtSolidLine03 TextLineSpecSolidLine

	txtSolidLine03,
		err = txtSolidLine01.CopyOut(
		ePrefix.XCtx(
			"txtSolidLine01->txtSolidLine03"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtSolidLine03.leftMargin = -97

	isValid,
		err = txtSolidLineAtom.testValidityOfTextSpecSolidLine(
		&txtSolidLine03,
		ePrefix.XCtx(
			"Testing 'txtSolidLine03'"))

	if isValid == false {
		t.Errorf("\n%v - ERROR\n"+
			"Expected 'isValid' equal 'true' because\n"+
			"txtSolidLine03.leftMargin = -97 defaulted to zero.\n"+
			"HOWEVER, 'isValid' IS 'false'!!!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	if err != nil {
		t.Errorf("\n%v - ERROR\n"+
			"Expected no errors from \n"+
			"txtSolidLineAtom.testValidityOfTextSpecSolidLine() \n"+
			"because txtSolidLine03.leftMargin = -97 defaulted to zero.\n"+
			"HOWEVER, AN ERROR WAS RETURNED!!!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	var txtSolidLine04 TextLineSpecSolidLine

	txtSolidLine04,
		err = txtSolidLine01.CopyOut(
		ePrefix.XCtx(
			"txtSolidLine01->txtSolidLine04"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtSolidLine04.rightMargin = -97

	isValid,
		err = txtSolidLineAtom.testValidityOfTextSpecSolidLine(
		&txtSolidLine04,
		ePrefix.XCtx(
			"Testing 'txtSolidLine04'"))

	if isValid == false {
		t.Errorf("%v - ERROR\n"+
			"Expected 'isValid' equal 'true' because\n"+
			"txtSolidLine04.rightMargin = -97 defaulted to zero.\n"+
			"HOWEVER, 'isValid' IS 'false'!!!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	if err != nil {
		t.Errorf("%v - ERROR\n"+
			"Expected no errors returned from \n"+
			"txtSolidLineAtom.testValidityOfTextSpecSolidLine() \n"+
			"because txtSolidLine04.rightMargin = -97 defaulted to zero.\n"+
			"HOWEVER, AN ERROR WAS RETURNED!!!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	var txtSolidLine05 TextLineSpecSolidLine

	txtSolidLine05,
		err = txtSolidLine01.CopyOut(
		ePrefix.XCtx(
			"txtSolidLine01->txtSolidLine05"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtSolidLine05.leftMargin = 1000001

	isValid,
		err = txtSolidLineAtom.testValidityOfTextSpecSolidLine(
		&txtSolidLine05,
		ePrefix.XCtx(
			"Testing 'txtSolidLine05'"))

	if isValid == true {
		t.Errorf("%v - ERROR\n"+
			"Expected 'isValid' equal 'false' because\n"+
			"txtSolidLine05.leftMargin = 1000001.\n"+
			"HOWEVER, 'isValid' IS 'true'!!!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from \n"+
			"txtSolidLineAtom.testValidityOfTextSpecSolidLine() \n"+
			"because txtSolidLine05.leftMargin = 1000001\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	var txtSolidLine06 TextLineSpecSolidLine

	txtSolidLine06,
		err = txtSolidLine01.CopyOut(
		ePrefix.XCtx(
			"txtSolidLine01->txtSolidLine06"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtSolidLine06.rightMargin = 1000001

	isValid,
		err = txtSolidLineAtom.testValidityOfTextSpecSolidLine(
		&txtSolidLine06,
		ePrefix.XCtx(
			"Testing 'txtSolidLine06'"))

	if isValid == true {
		t.Errorf("%v - ERROR\n"+
			"Expected 'isValid' equal 'false' because\n"+
			"txtSolidLine06.rightMargin = 1000001.\n"+
			"HOWEVER, 'isValid' IS 'true'!!!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from \n"+
			"txtSolidLineAtom.testValidityOfTextSpecSolidLine() \n"+
			"because txtSolidLine06.rightMargin = 1000001\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	var txtSolidLine07 TextLineSpecSolidLine

	txtSolidLine07,
		err = txtSolidLine01.CopyOut(
		ePrefix.XCtx(
			"txtSolidLine01->txtSolidLine06"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtSolidLine07.newLineChars = nil

	isValid,
		err = txtSolidLineAtom.testValidityOfTextSpecSolidLine(
		&txtSolidLine07,
		ePrefix.XCtx(
			"Testing 'txtSolidLine07'"))

	if isValid == false {
		t.Errorf("%v - ERROR\n"+
			"Expected 'isValid' equal 'true' because\n"+
			"txtSolidLine07 is a valid instance\n"+
			"HOWEVER, 'isValid' IS 'false'!!!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	if err != nil {
		t.Errorf("%v - ERROR\n"+
			"Expected no errors from \n"+
			"txtSolidLineAtom.testValidityOfTextSpecSolidLine() \n"+
			"because txtSolidLine07 is a valid instance.\n"+
			"HOWEVER, AN ERROR WAS RETURNED!!!\n"+
			"Error='%v'\n",
			ePrefix.XCtxEmpty().String(),
			err.Error())

		return
	}

	isValid,
		err = txtSolidLineAtom.testValidityOfTextSpecSolidLine(
		nil,
		ePrefix.XCtx(
			"'txtSolidLine'=='nil'"))

	if isValid == true {
		t.Errorf("%v - ERROR\n"+
			"Expected 'isValid' equal 'false' because\n"+
			"txtSolidLine =='nil'.\n"+
			"HOWEVER, 'isValid' IS 'true'!!!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from \n"+
			"txtSolidLineAtom.testValidityOfTextSpecSolidLine() \n"+
			"because txtSolidLine =='nil'\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	var txtSolidLine08 TextLineSpecSolidLine

	txtSolidLine08,
		err = txtSolidLine01.CopyOut(
		ePrefix.XCtx(
			"txtSolidLine01->txtSolidLine08"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtSolidLine08.newLineChars = []rune{'\n', 0, 0, '\n'}

	isValid,
		err = txtSolidLineAtom.testValidityOfTextSpecSolidLine(
		&txtSolidLine08,
		ePrefix.XCtx(
			"Testing 'txtSolidLine08'"))

	if isValid == true {
		t.Errorf("%v - ERROR\n"+
			"Expected 'isValid' equal 'false' because\n"+
			"txtSolidLine08.newLineChars = []rune{'\\n',0,0,'\\n'}\n"+
			"HOWEVER, 'isValid' IS 'true'!!!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from \n"+
			"txtSolidLineAtom.testValidityOfTextSpecSolidLine() \n"+
			"because txtSolidLine08.newLineChars = []rune{'\\n',0,0,'\\n'}\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	var txtSolidLine09 TextLineSpecSolidLine

	txtSolidLine09,
		err = txtSolidLine01.CopyOut(
		ePrefix.XCtx(
			"txtSolidLine01->txtSolidLine09"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtSolidLine09.solidLineChars = []rune{'-', 0, 0, '-'}

	isValid,
		err = txtSolidLineAtom.testValidityOfTextSpecSolidLine(
		&txtSolidLine09,
		ePrefix.XCtx(
			"Testing 'txtSolidLine09'"))

	if isValid == true {
		t.Errorf("%v - ERROR\n"+
			"Expected 'isValid' equal 'false' because\n"+
			"txtSolidLine09.solidLineChars = []rune{'-', 0, 0, '-'}\n"+
			"HOWEVER, 'isValid' IS 'true'!!!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from \n"+
			"txtSolidLineAtom.testValidityOfTextSpecSolidLine() \n"+
			"because txtSolidLine09.solidLineChars = []rune{'-', 0, 0, '-'}\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

}
