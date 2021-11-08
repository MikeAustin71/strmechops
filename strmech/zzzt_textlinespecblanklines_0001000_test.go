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
			"Expected an error return from txtFieldSpacerOne."+
			"CopyOut()\n"+
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
