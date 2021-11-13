package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"strings"
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

func TestTextLineSpecBlankLines_CopyIn_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_copyIn_000100()",
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

func TestTextLineSpecBlankLines_CopyOut_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_CopyOut_000100()",
		"")

	numOfBlankLines := 3

	txtBlankLinesBase,
		err := TextLineSpecBlankLines{}.NewPtrDefaultBlankLines(
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
		err = TextLineSpecBlankLines{}.NewDefaultBlankLines(
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
		TextFieldSpecDateTime{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtBlankLinesBase2."+
			"CopyOut()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	_,
		err = txtBlankLinesBase3.CopyOut(
		TextFieldSpecDateTime{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtBlankLinesBase3."+
			"CopyOut()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecBlankLines_CopyOutITextLine_000100(t *testing.T) {

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

	var iTextLine ITextLineSpecification

	iTextLine,
		err =
		txtBlankLinesBase.CopyOutITextLine(
			ePrefix.XCtx(
				"txtBlankLinesBase->iTextLine"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtBlankLinesBravo, ok := iTextLine.(*TextLineSpecBlankLines)

	if !ok {
		t.Errorf("%v\n"+
			"Error: iTextLine.(*TextLineSpecBlankLines)\n"+
			"Could not convert 'iTextLine' to TextLineSpecBlankLines\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	if !txtBlankLinesBase.Equal(txtBlankLinesBravo) {
		t.Errorf("%v\n"+
			"Error: Expected txtBlankLinesBase == txtBlankLinesBravo\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	if !txtBlankLinesBase.EqualITextLine(iTextLine) {
		t.Errorf("%v\n"+
			"Error: txtBlankLinesBase.EqualITextLine(iTextLine)\n"+
			"Expected txtBlankLinesBase == iTextLine\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	if !txtBlankLinesBravo.EqualITextLine(iTextLine) {
		t.Errorf("%v\n"+
			"Error: txtBlankLinesBravo.EqualITextLine(iTextLine)\n"+
			"Expected txtBlankLinesBravo == iTextLine\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	txtBlankLinesCharlie := TextLineSpecBlankLines{}

	_,
		err =
		txtBlankLinesCharlie.CopyOutITextLine(
			ePrefix.XCtx(
				"txtBlankLinesCharlie->_"))

	if err == nil {
		t.Errorf("%v\n"+
			"Error: txtBlankLinesCharlie.CopyOutITextLine()\n"+
			"Expected an error return because 'txtBlankLinesCharlie'"+
			" is empty!\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	var txtBlankLinesDelta TextLineSpecBlankLines

	txtBlankLinesDelta,
		err = TextLineSpecBlankLines{}.NewDefaultBlankLines(
		numOfBlankLines,
		ePrefix.XCtx(
			"txtBlankLinesBase"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtBlankLinesDelta.IsValidInstanceError(
		ePrefix.XCtx(
			"txtBlankLinesBase"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	_,
		err =
		txtBlankLinesCharlie.CopyOutITextLine(
			TextFieldSpecDateTime{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtBlankLinesCharlie."+
			"CopyOutITextLine()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecBlankLines_CopyOutPtr_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_CopyOutPtr_000100()",
		"")

	numOfBlankLines := 3

	txtBlankLinesBase,
		err := TextLineSpecBlankLines{}.NewPtrDefaultBlankLines(
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

	var txtBlankLinesOut *TextLineSpecBlankLines

	txtBlankLinesOut,
		err =
		txtBlankLinesBase.CopyOutPtr(
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
		err = txtBlankLinesBase2.CopyOutPtr(
		ePrefix.XCtx(
			"txtBlankLinesBase2 is empty!"))

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from "+
			"txtBlankLinesBase2.CopyOutPtr()\n"+
			"because input parameter 'txtBlankLinesBase2' is "+
			"empty and invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtx("Missing Error Return"))

		return
	}

	var txtBlankLinesBase3 TextLineSpecBlankLines

	txtBlankLinesBase3,
		err = TextLineSpecBlankLines{}.NewDefaultBlankLines(
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
		err = txtBlankLinesBase2.CopyOutPtr(
		TextFieldSpecDateTime{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtBlankLinesBase2."+
			"CopyOutPtr()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	_,
		err = txtBlankLinesBase3.CopyOutPtr(
		TextFieldSpecDateTime{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtBlankLinesBase3."+
			"CopyOutPtr()\n"+
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

func TestTextLineSpecBlankLines_Empty_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_Empty_000100()",
		"")

	blankLinesAlpha := TextLineSpecBlankLines{}

	blankLinesAlpha.Empty()

	var err error

	blankLinesAlpha,
		err = TextLineSpecBlankLines{}.NewDefaultBlankLines(
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

	blankLinesAlpha.Empty()

	if blankLinesAlpha.numBlankLines != 0 {
		t.Errorf("%v - ERROR\n"+
			"blankLinesAlpha.Empty()\n"+
			"Expected blankLinesAlpha.numBlankLines == 0\n"+
			"Instead, blankLinesAlpha.numBlankLines == %v\n",
			ePrefix.XCtxEmpty().String(),
			blankLinesAlpha.numBlankLines)

		return
	}

	if blankLinesAlpha.newLineChars != nil {
		t.Errorf("%v - ERROR\n"+
			"blankLinesAlpha.Empty()\n"+
			"Expected blankLinesAlpha.newLineChars == nil\n"+
			"Instead, blankLinesAlpha.newLineChars = '%v'\n"+
			" blankLinesAlpha.newLineChars array = '%v'",
			ePrefix.XCtxEmpty().String(),
			string(blankLinesAlpha.newLineChars),
			blankLinesAlpha.newLineChars)

		return
	}

	if blankLinesAlpha.lock != nil {
		t.Errorf("%v - ERROR\n" +
			"blankLinesAlpha.Empty()\n" +
			"Expected blankLinesAlpha.lock == nil\n" +
			"HOWEVER, 'blankLinesAlpha.lock' IS NOT 'nil'!\n" +
			ePrefix.XCtxEmpty().String())

		return
	}

	var blankLinesBravo TextLineSpecBlankLines

	blankLinesBravo,
		err = TextLineSpecBlankLines{}.NewDefaultBlankLines(
		3,
		ePrefix.XCtx(
			"blankLinesBravo"))

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

	blankLinesBravo.textLineReader =
		strings.NewReader("Hello World!")

	blankLinesBravo.Empty()

	if blankLinesBravo.textLineReader != nil {
		t.Errorf("%v - ERROR\n"+
			"Expected blankLinesBravo.textLineReader == nil\n"+
			"HOWEVER 'blankLinesBravo.textLineReader' IS NOT 'nil'!\n",
			ePrefix.XCtxEmpty().String())

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

func TestTextLineSpecBlankLines_Equal_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_Equal_000100()",
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

func TestTextLineSpecBlankLines_EqualITextLine_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_copyOut_000100()",
		"")

	var err error

	txtBlankLinesBase := TextLineSpecBlankLines{}

	areEqual := txtBlankLinesBase.EqualITextLine(nil)

	if areEqual {

		t.Errorf("%v - ERROR\n"+
			"areEqual = txtBlankLinesBase.EqualITextLine(nil)\n"+
			"Expected areEqual == 'false' because\n"+
			"'txtBlankLinesBase' is empty\n"+
			"HOWEVER, 'areEqual' == 'true'!!!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	numOfBlankLines := 3

	txtBlankLinesBase,
		err = TextLineSpecBlankLines{}.NewDefaultBlankLines(
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

	leftMargin := 3
	rightMargin := 3
	textString := "How now brown cow!"

	var plainTextLine01 TextLineSpecPlainText

	plainTextLine01,
		err = TextLineSpecPlainText{}.NewDefault(
		leftMargin,
		rightMargin,
		textString,
		ePrefix.XCtx("plainTextLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = plainTextLine01.IsValidInstanceError(
		ePrefix.XCtx("plainTextLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	areEqual =
		txtBlankLinesBase.EqualITextLine(&plainTextLine01)

	if areEqual {

		t.Errorf("%v - ERROR\n"+
			"areEqual = txtBlankLinesBase.EqualITextLine(&plainTextLine01)\n"+
			"Expected areEqual == 'false' because\n"+
			"'plainTextLine01' is of type 'TextLineSpecPlainText'.\n"+
			"HOWEVER, 'areEqual' == 'true'!!!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	var txtBlankLinesBravo TextLineSpecBlankLines

	txtBlankLinesBravo,
		err = txtBlankLinesBase.CopyOut(
		ePrefix.XCtx(
			"txtBlankLinesBase->txtBlankLinesBravo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtBlankLinesBravo.IsValidInstanceError(
		ePrefix.XCtx(
			"txtBlankLinesBravo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	areEqual = txtBlankLinesBase.Equal(
		&txtBlankLinesBravo)

	if !areEqual {

		t.Errorf("%v - ERROR\n"+
			"areEqual = txtBlankLinesBase.Equal(&txtBlankLinesBravo)\n"+
			"Expected areEqual == 'true' because\n"+
			"'txtBlankLinesBravo' is a copy of 'txtBlankLinesBase'.\n"+
			"HOWEVER, 'areEqual' == 'false'!!!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	areEqual = txtBlankLinesBase.EqualITextLine(
		&txtBlankLinesBravo)

	if !areEqual {

		t.Errorf("%v - ERROR\n"+
			"areEqual = txtBlankLinesBase.EqualITextLine(&txtBlankLinesBravo)\n"+
			"Expected areEqual == 'true' because\n"+
			"'txtBlankLinesBravo' is a copy of 'txtBlankLinesBase'.\n"+
			"HOWEVER, 'areEqual' == 'false'!!!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecBlankLines_GetFormattedText_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_CopyOut_000100()",
		"")

	txtBlankLinesBase := TextLineSpecBlankLines{}

	_,
		err := txtBlankLinesBase.GetFormattedText(
		ePrefix.XCtx(
			"Empty txtBlankLinesBase"))

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from "+
			"txtBlankLinesBase.GetFormattedText()\n"+
			"because 'txtBlankLinesBase' is "+
			"empty and invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtx("Missing Error Return"))

		return
	}

	numOfBlankLines := 3

	txtBlankLinesBase,
		err = TextLineSpecBlankLines{}.NewDefaultBlankLines(
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

	expectedFmtStr :=
		strings.Repeat("\n", 3)

	var actualStr string

	actualStr,
		err = txtBlankLinesBase.GetFormattedText(
		ePrefix.XCtx(
			"valid txtBlankLinesBase"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedFmtStr),
			true)

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(actualStr),
			true)

	if printableExpectedStr != printableActualStr {
		t.Errorf("%v\n"+
			"txtBlankLinesBase.GetFormattedText()\n"+
			"Error: Expected Text String DOES NOT match\n"+
			"Actual Text String.\n"+
			"Expected Text String = '%v'\n"+
			"Instead, Text String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	txtBlankLinesBaseAlpha := TextLineSpecBlankLines{}

	err = txtBlankLinesBaseAlpha.SetNumberOfBlankLines(
		3,
		ePrefix.XCtx(
			"txtBlankLinesBaseAlpha"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	_,
		err = txtBlankLinesBaseAlpha.GetFormattedText(
		TextFieldSpecDateTime{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtBlankLinesBaseAlpha."+
			"GetFormattedText()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecBlankLines_GetNewLineChars_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_CopyOut_000100()",
		"")

	numOfBlankLines := 3

	txtBlankLinesZero := TextLineSpecBlankLines{}

	expectedFmtStr := ""

	actualStr := txtBlankLinesZero.GetNewLineChars()

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedFmtStr),
			true)

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(actualStr),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Expected newLineChars = '%v'\n"+
			"Instead, newLineChars = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	actualStr = "!\n"

	txtBlankLinesOne,
		err := TextLineSpecBlankLines{}.NewBlankLines(
		numOfBlankLines,
		actualStr,
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

	expectedFmtStr =
		strings.Repeat(actualStr, numOfBlankLines)

	printableExpectedStr =
		sMech.ConvertNonPrintableChars(
			[]rune(actualStr),
			true)

	actualStr = txtBlankLinesOne.GetNewLineChars()

	printableActualStr =
		sMech.ConvertNonPrintableChars(
			[]rune(actualStr),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #2\n"+
			"Expected newLineChars = '%v'\n"+
			"Instead, newLineChars = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	printableExpectedStr =
		sMech.ConvertNonPrintableChars(
			[]rune(expectedFmtStr),
			true)

	var actualFmtStr string

	actualFmtStr,
		err =
		txtBlankLinesOne.GetFormattedText(
			ePrefix.XCtx(
				"txtBlankLinesOne"))

	printableActualStr =
		sMech.ConvertNonPrintableChars(
			[]rune(actualFmtStr),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #3\n"+
			"Expected Formatted String = '%v'\n"+
			"Instead, Formatted String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	return
}
