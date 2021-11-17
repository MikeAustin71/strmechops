package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"testing"
)

func TestTextLineSpecBlankLines_copyIn_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_copyIn_000100()",
		"")

	numOfBlankLines := 3

	incomingBlankLines,
		err := TextLineSpecBlankLines{}.NewDefaultBlankLines(
		numOfBlankLines,
		ePrefix.XCtx(
			"incomingBlankLines"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = incomingBlankLines.IsValidInstanceError(
		ePrefix.XCtx(
			"incomingBlankLines"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	targetBlankLines := TextLineSpecBlankLines{}

	txtBlankLinesMolecule := textLineSpecBlankLinesMolecule{}

	err =
		txtBlankLinesMolecule.copyIn(
			&targetBlankLines,
			&incomingBlankLines,
			ePrefix.XCtx(
				"incomingBlankLines->targetBlankLines"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = targetBlankLines.IsValidInstanceError(
		ePrefix.XCtx(
			"incomingBlankLines"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !targetBlankLines.Equal(&incomingBlankLines) {
		t.Errorf("%v - ERROR\n"+
			"Expected targetBlankLines==incomingBlankLines\n"+
			"HOWEVER, THESE INSTANCES ARE NOT EQUAL!\n",
			ePrefix.XCtxEmpty().String())
		return
	}

	targetBlankLinesTwo := TextLineSpecBlankLines{}

	err =
		txtBlankLinesMolecule.copyIn(
			&targetBlankLinesTwo,
			nil,
			ePrefix.XCtx(
				"'nil'->targetBlankLinesTwo"))

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from "+
			"txtBlankLinesMolecule.copyIn()\n"+
			"because input parameter 'incomingBlkLines' is 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtx("Missing Error Return"))
		return
	}

	txtBlankLinesMolecule2 := textLineSpecBlankLinesMolecule{}

	err =
		txtBlankLinesMolecule2.copyIn(
			nil,
			&incomingBlankLines,
			ePrefix.XCtx(
				"'nil'->targetBlankLinesTwo"))

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from "+
			"txtBlankLinesMolecule.copyIn()\n"+
			"because input parameter 'targetBlkLines' is 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtx("Missing Error Return"))
		return
	}

	incomingBlankLines.numBlankLines = -9999

	err = textLineSpecBlankLinesMolecule{}.ptr().
		copyIn(
			&targetBlankLinesTwo,
			&incomingBlankLines,
			ePrefix.XCtx(
				"incomingBlankLines is invalid!"))

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from "+
			"txtBlankLinesMolecule.copyIn()\n"+
			"because input parameter 'incomingBlankLines' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtx("Missing Error Return"))
		return
	}

	return
}

func TestTextLineSpecBlankLines_copyOut_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_copyOut_000100()",
		"")

	numOfBlankLines := 3

	txtBlankLinesBase,
		err := TextLineSpecBlankLines{}.NewDefaultBlankLines(
		numOfBlankLines,
		ePrefix.XCtx(
			"txtBlankLinesBase"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtBlankLinesBase.IsValidInstanceError(
		ePrefix.XCtx(
			"txtBlankLinesBase"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtBlankLinesMolecule := textLineSpecBlankLinesMolecule{}

	var txtBlankLinesOut TextLineSpecBlankLines

	txtBlankLinesOut,
		err =
		txtBlankLinesMolecule.copyOut(
			&txtBlankLinesBase,
			ePrefix.XCtx(
				"txtBlankLinesBase->txtBlankLinesOut"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtBlankLinesOut.IsValidInstanceError(
		ePrefix.XCtx(
			"txtBlankLinesOut"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !txtBlankLinesOut.Equal(&txtBlankLinesBase) {
		t.Errorf("%v - ERROR\n"+
			"Expected txtBlankLinesOut==txtBlankLinesBase\n"+
			"HOWEVER, THESE INSTANCES ARE NOT EQUAL!\n",
			ePrefix.XCtxEmpty().String())
		return
	}

	_,
		err =
		txtBlankLinesMolecule.copyOut(
			nil,
			ePrefix.XCtx(
				"'nil'->txtBlankLinesOutTwo"))

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from "+
			"txtBlankLinesMolecule.copyOut()\n"+
			"because input parameter 'txtBlankLines' is 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtx("Missing Error Return"))
		return
	}

	txtBlankLinesBase.numBlankLines = -9999

	_,
		err = textLineSpecBlankLinesMolecule{}.ptr().
		copyOut(
			&txtBlankLinesBase,
			ePrefix.XCtx(
				"txtBlankLinesBase is invalid!"))

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from "+
			"txtBlankLinesMolecule.copyOut()\n"+
			"because input parameter 'txtBlankLinesBase' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtx("Missing Error Return"))
		return
	}

	return
}

func TestTextLineSpecBlankLines_empty_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_empty_000100()",
		"")

	blankLinesAlpha,
		err := TextLineSpecBlankLines{}.NewDefaultBlankLines(
		3,
		ePrefix.XCtx(
			"blankLinesAlpha"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = blankLinesAlpha.IsValidInstanceError(
		ePrefix.XCtx(
			"blankLinesAlpha"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if blankLinesAlpha.numBlankLines != 3 {
		t.Errorf("%v - ERROR\n"+
			"Expected blankLinesAlpha.numBlankLines == 3\n"+
			"Instead, blankLinesAlpha.numBlankLines == %v\n",
			ePrefix.XCtxEmpty().String(),
			blankLinesAlpha.numBlankLines)

		return
	}

	txtBlankLinesMolecule := textLineSpecBlankLinesMolecule{}

	txtBlankLinesMolecule.empty(
		nil)

	txtBlankLinesMolecule.empty(
		&blankLinesAlpha)

	if blankLinesAlpha.numBlankLines != 0 {
		t.Errorf("%v - ERROR\n"+
			"Expected blankLinesAlpha.numBlankLines == 0\n"+
			"Instead, blankLinesAlpha.numBlankLines == %v\n",
			ePrefix.XCtxEmpty().String(),
			blankLinesAlpha.numBlankLines)

		return
	}

	if blankLinesAlpha.newLineChars != nil {
		t.Errorf("%v - ERROR\n"+
			"Expected blankLinesAlpha.newLineChars == nil\n"+
			"Instead, blankLinesAlpha.newLineChars = '%v'\n"+
			" blankLinesAlpha.newLineChars array = '%v'",
			ePrefix.XCtxEmpty().String(),
			string(blankLinesAlpha.newLineChars),
			blankLinesAlpha.newLineChars)

		return
	}

	return
}

func TestTextLineSpecBlankLines_equal_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_equal_000100()",
		"")

	blankLinesAlpha,
		err := TextLineSpecBlankLines{}.NewDefaultBlankLines(
		3,
		ePrefix.XCtx(
			"blankLinesAlpha"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = blankLinesAlpha.IsValidInstanceError(
		ePrefix.XCtx(
			"blankLinesAlpha"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var blankLinesBravo TextLineSpecBlankLines

	blankLinesBravo,
		err = blankLinesAlpha.CopyOut(
		ePrefix.XCtx(
			"blankLinesAlpha->blankLinesBravo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtBlankLinesMolecule := textLineSpecBlankLinesMolecule{}

	areEqual :=
		txtBlankLinesMolecule.equal(
			&blankLinesAlpha,
			nil)

	if areEqual {
		t.Errorf("%v - ERROR\n"+
			"Expected txtBlankLinesMolecule.equal() == 'false'\n"+
			"because input parameter 'txtBlankLinesTwo' is 'nil'.\n"+
			"HOWEVER, THE METHOD RETURNED 'true'!!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	areEqual =
		txtBlankLinesMolecule.equal(
			nil,
			&blankLinesBravo)

	if areEqual {
		t.Errorf("%v - ERROR\n"+
			"Expected txtBlankLinesMolecule.equal() == 'false'\n"+
			"because input parameter 'txtBlankLinesOne' is 'nil'.\n"+
			"HOWEVER, THE METHOD RETURNED 'true'!!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	var blankLinesCharlie TextLineSpecBlankLines

	blankLinesCharlie,
		err = blankLinesAlpha.CopyOut(
		ePrefix.XCtx(
			"blankLinesAlpha->blankLinesCharlie"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	blankLinesCharlie.numBlankLines = 95

	areEqual =
		txtBlankLinesMolecule.equal(
			&blankLinesAlpha,
			&blankLinesCharlie)

	if areEqual {
		t.Errorf("%v - ERROR\n"+
			"Expected txtBlankLinesMolecule.equal() == 'false' because\n"+
			"input parameter 'blankLinesCharlie' has a different"+
			"number of blank lines.\n"+
			"HOWEVER, THE METHOD RETURNED 'true'!!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	var blankLinesDelta TextLineSpecBlankLines

	blankLinesDelta,
		err = blankLinesAlpha.CopyOut(
		ePrefix.XCtx(
			"blankLinesAlpha->blankLinesDelta"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	blankLinesDelta.newLineChars =
		[]rune{'H', 'e', 'l', 'l', 'o'}

	areEqual =
		txtBlankLinesMolecule.equal(
			&blankLinesAlpha,
			&blankLinesDelta)

	if areEqual {
		t.Errorf("%v - ERROR\n"+
			"Expected txtBlankLinesMolecule.equal() == 'false' because\n"+
			"input parameter 'blankLinesDelta' has different"+
			"new line characters.\n"+
			"HOWEVER, THE METHOD RETURNED 'true'!!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecBlankLines_testValidityOfTextLineSpecBlankLines_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_testValidityOfTextLineSpecBlankLines_000100()",
		"")

	txtBlankLinesAtom := textLineSpecBlankLinesAtom{}

	_,
		err := txtBlankLinesAtom.testValidityOfTextLineSpecBlankLines(
		nil,
		ePrefix.XCtx(
			"txtBlankLines== 'nil'"))

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtBlankLinesAtom."+
			"testValidityOfTextLineSpecBlankLines()\n"+
			"because 'txtBlankLines' is 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	numOfBlankLines := 3

	var txtBlankLinesOne TextLineSpecBlankLines

	txtBlankLinesOne,
		err = TextLineSpecBlankLines{}.NewDefaultBlankLines(
		numOfBlankLines,
		ePrefix.XCtx(
			"txtBlankLinesOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtBlankLinesOne.IsValidInstanceError(
		ePrefix.XCtx(
			"txtBlankLinesOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtBlankLinesOne.numBlankLines = 1000001

	_,
		err = txtBlankLinesAtom.testValidityOfTextLineSpecBlankLines(
		&txtBlankLinesOne,
		ePrefix.XCtx(
			"txtBlankLinesOne.numBlankLines = 1000001"))

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtBlankLinesAtom."+
			"testValidityOfTextLineSpecBlankLines()\n"+
			"because 'numBlankLines' is '1000001' and invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecBlankLines_testValidityNewLinesChars_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_testValidityNewLinesChars_000100()",
		"")

	txtBlankLinesElectron := textLineSpecBlankLinesElectron{}

	newLineRunes := []rune{'\n'}

	err := txtBlankLinesElectron.testValidityNewLinesChars(
		newLineRunes,
		ePrefix.XCtx(
			"newLineRunes := []rune{'\\n'}"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	return
}

func TestTextLineSpecBlankLines_getFormattedText_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_getFormattedText_000100()",
		"")

	txtBlankLinesMolecule := textLineSpecBlankLinesMolecule{}

	_,
		err := txtBlankLinesMolecule.getFormattedText(
		nil,
		ePrefix.XCtx(
			"txtBlankLines ='nil'"))

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from textLineSpecBlankLinesMolecule."+
			"getFormattedText()\n"+
			"because 'txtBlankLines' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecBlankLines_setTextLinesSpecBlankLines_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_setTextLinesSpecBlankLines_000100()",
		"")

	txtBlankLinesMolecule := textLineSpecBlankLinesMolecule{}

	numOfBlankLines := 3

	newLineRunes := []rune{'\n'}

	err := txtBlankLinesMolecule.setTextLinesSpecBlankLines(
		nil,
		numOfBlankLines,
		newLineRunes,
		ePrefix.XCtx(
			"txtBlankLines ='nil'"))

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from textLineSpecBlankLinesMolecule."+
			"setTextLinesSpecBlankLines()\n"+
			"because 'txtBlankLines' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}
