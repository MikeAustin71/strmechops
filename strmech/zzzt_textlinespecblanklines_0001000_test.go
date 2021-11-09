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
		err := TextLineSpecBlankLines{}.NewBlankLines(
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

func TestTextLineSpecBlankLines_CopyIn_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_copyIn_000100()",
		"")

	blankLinesAlpha,
		err := TextLineSpecBlankLines{}.NewBlankLines(
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

	blankLinesBravo := TextLineSpecBlankLines{}

	err =
		blankLinesBravo.CopyIn(
			&blankLinesAlpha,
			ePrefix.XCtx(
				"blankLinesAlpha->blankLinesBravo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = blankLinesBravo.IsValidInstanceError(
		ePrefix.XCtx(
			"blankLinesBravo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !blankLinesBravo.Equal(&blankLinesAlpha) {
		t.Errorf("%v - ERROR\n"+
			"Expected blankLinesAlpha==blankLinesBravo\n"+
			"HOWEVER, THESE INSTANCES ARE NOT EQUAL!\n",
			ePrefix.XCtxEmpty().String())
		return
	}

	blankLinesCharlie := TextLineSpecBlankLines{}

	err =
		blankLinesCharlie.CopyIn(
			&blankLinesAlpha,
			TextFieldSpecDateTime{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from blankLinesCharlie."+
			"CopyIn()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

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
		err := TextLineSpecBlankLines{}.NewBlankLines(
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

func TestTextLineSpecBlankLines_CopyOut_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_CopyOut_000100()",
		"")

	numOfBlankLines := 3

	txtBlankLinesBase,
		err := TextLineSpecBlankLines{}.NewPtrBlankLines(
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

	txtBlankLinesOut := TextLineSpecBlankLines{}

	txtBlankLinesOut,
		err =
		txtBlankLinesBase.CopyOut(
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

	if !txtBlankLinesOut.Equal(txtBlankLinesBase) {
		t.Errorf("%v - ERROR\n"+
			"Expected txtBlankLinesOut==txtBlankLinesBase\n"+
			"HOWEVER, THESE INSTANCES ARE NOT EQUAL!\n",
			ePrefix.XCtxEmpty().String())
		return
	}

	txtBlankLinesBase2 := TextLineSpecBlankLines{}

	_,
		err = txtBlankLinesBase2.CopyOut(
		ePrefix.XCtx(
			"txtBlankLinesBase2 is empty!"))

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from "+
			"txtBlankLinesBase2.CopyOut()\n"+
			"because input parameter 'txtBlankLinesBase2' is "+
			"empty and invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtx("Missing Error Return"))

		return
	}

	var txtBlankLinesBase3 TextLineSpecBlankLines

	txtBlankLinesBase3,
		err = TextLineSpecBlankLines{}.NewBlankLines(
		numOfBlankLines,
		ePrefix.XCtx(
			"->txtBlankLinesBase3"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtBlankLinesBase3.IsValidInstanceError(
		ePrefix.XCtx(
			"txtBlankLinesBase3"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	_,
		err = txtBlankLinesBase2.CopyOut(
		ePrefix.XCtx(
			"txtBlankLinesBase2 is empty!"))

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtBlankLinesBase2."+
			"CopyOut()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecBlankLines_empty_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_empty_000100()",
		"")

	blankLinesAlpha,
		err := TextLineSpecBlankLines{}.NewBlankLines(
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
			"Instead, blankLinesAlpha.numBlankLines = '%v'\n"+
			" blankLinesAlpha.numBlankLines array = '%v'",
			ePrefix.XCtxEmpty().String(),
			string(blankLinesAlpha.newLineChars),
			blankLinesAlpha.numBlankLines)

		return
	}

	return
}

func TestTextLineSpecBlankLines_equal_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_equal_000100()",
		"")

	blankLinesAlpha,
		err := TextLineSpecBlankLines{}.NewBlankLines(
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

func TestTextLineSpecBlankLines_Equal_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_Equal_000100()",
		"")

	blankLinesAlpha,
		err := TextLineSpecBlankLines{}.NewBlankLines(
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

	blankLinesBravo := TextLineSpecBlankLines{}

	if blankLinesBravo.Equal(&blankLinesAlpha) {
		t.Errorf("%v - ERROR\n"+
			"Expected blankLinesBravo.Equal(&blankLinesAlpha)== 'false'\n"+
			"because 'blankLinesBravo' is empty.\n"+
			"HOWEVER, THE RETURNED VALUE WAS 'true'!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	if blankLinesAlpha.Equal(nil) {
		t.Errorf("%v - ERROR\n"+
			"Expected blankLinesAlpha.Equal(nil) == 'false'\n"+
			"because input parameter is 'nil'.\n"+
			"HOWEVER, THE METHOD RETURNED 'true'!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	blankLinesBravo,
		err = blankLinesAlpha.CopyOut(
		ePrefix.XCtx(
			"blankLinesAlpha->blankLinesBravo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !blankLinesBravo.Equal(&blankLinesAlpha) {
		t.Errorf("%v - ERROR \n"+
			"Expected blankLinesAlpha==blankLinesBravo\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}
