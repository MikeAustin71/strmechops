package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"strings"
	"testing"
)

func TestTextLineSpecBlankLines_CopyIn_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_CopyIn_000100()",
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
			StrMech{})

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

func TestTextLineSpecBlankLines_CopyIn_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_CopyIn_000200()",
		"")

	numOfBlankLines := 3

	newLineRunes := []rune{'!', '\n'}

	blankLinesOne,
		err := TextLineSpecBlankLines{}.NewBlankLineRunes(
		numOfBlankLines,
		newLineRunes,
		ePrefix.XCtx(
			"blankLinesOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = blankLinesOne.IsValidInstanceError(
		ePrefix.XCtx(
			"blankLinesOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	blankLinesTwo := TextLineSpecBlankLines{}

	err = blankLinesTwo.CopyIn(
		&blankLinesOne,
		ePrefix.XCtx(
			"blankLinesOne->blankLinesTwo"))

	err = blankLinesTwo.IsValidInstanceError(
		ePrefix.XCtx(
			"blankLinesTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !blankLinesTwo.Equal(&blankLinesOne) {

		t.Errorf("%v - ERROR\n"+
			"blankLinesTwo.Equal(&blankLinesOne)\n"+
			"Expected blankLinesOne==blankLinesTwo\n"+
			"HOWEVER, THESE INSTANCES ARE NOT EQUAL!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	expectedFmtStr :=
		strings.Repeat(string(newLineRunes), numOfBlankLines)

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedFmtStr),
			true)

	var actualFmtStr string

	actualFmtStr,
		err =
		blankLinesTwo.GetFormattedText(
			ePrefix.XCtx(
				"blankLinesTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	printableActualStr :=
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
		StrMech{})

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
		StrMech{})

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

func TestTextLineSpecBlankLines_CopyOut_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_CopyOut_000200()",
		"")

	numOfBlankLines := 4

	newLineRunes := []rune{'!', '\n'}

	blankLinesOne,
		err := TextLineSpecBlankLines{}.NewPtrBlankLineRunes(
		numOfBlankLines,
		newLineRunes,
		ePrefix.XCtx(
			"blankLinesOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = blankLinesOne.IsValidInstanceError(
		ePrefix.XCtx(
			"blankLinesOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var blankLinesTwo TextLineSpecBlankLines

	blankLinesTwo,
		err = blankLinesOne.CopyOut(
		ePrefix.XCtx(
			"blankLinesOne->blankLinesTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = blankLinesTwo.IsValidInstanceError(
		ePrefix.XCtx(
			"blankLinesTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !blankLinesTwo.Equal(blankLinesOne) {

		t.Errorf("%v - ERROR\n"+
			"blankLinesTwo.Equal(blankLinesOne)\n"+
			"Expected blankLinesOne==blankLinesTwo\n"+
			"HOWEVER, THESE INSTANCES ARE NOT EQUAL!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	expectedFmtStr :=
		strings.Repeat(string(newLineRunes), numOfBlankLines)

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedFmtStr),
			true)

	var actualFmtStr string

	actualFmtStr,
		err =
		blankLinesTwo.GetFormattedText(
			ePrefix.XCtx(
				"blankLinesTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	printableActualStr :=
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
			StrMech{})

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
		StrMech{})

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
		StrMech{})

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
		"TestTextLineSpecBlankLines_GetFormattedText_000100()",
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
		StrMech{})

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
		"TestTextLineSpecBlankLines_GetNewLineChars_000100()",
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

func TestTextLineSpecBlankLines_GetNewLineRunes_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_GetNewLineRunes_000100()",
		"")

	numOfBlankLines := 3

	txtBlankLinesZero := TextLineSpecBlankLines{}

	expectedFmtStr := ""

	actualRuneArray := txtBlankLinesZero.GetNewLineRunes()

	if actualRuneArray != nil {

		t.Errorf("%v - ERROR\n"+
			"txtBlankLinesZero.GetNewLineRunes()\n"+
			"Expected newLineRunes = 'nil'\n"+
			"Instead, newLineRunes = '%v'\n"+
			"Runes Value = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			string(actualRuneArray),
			actualRuneArray)

		return
	}

	actualRuneArray = []rune{'!', '\n'}

	txtBlankLinesOne,
		err := TextLineSpecBlankLines{}.NewBlankLineRunes(
		numOfBlankLines,
		actualRuneArray,
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
		strings.Repeat(string(actualRuneArray), numOfBlankLines)

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			actualRuneArray,
			true)

	actualRuneArray = txtBlankLinesOne.GetNewLineRunes()

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			actualRuneArray,
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #2\n"+
			"Expected newLineRunes = '%v'\n"+
			"Instead, newLineRunes = '%v'\n",
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

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

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

func TestTextLineSpecBlankLines_GetNumberOfBlankLines_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_GetNumberOfBlankLines_000100()",
		"")

	expectedNumOfBlankLines := 2

	newLineChars := "!\n"

	txtBlankLinesZero := TextLineSpecBlankLines{}

	actualNumOfBlankLines :=
		txtBlankLinesZero.GetNumOfBlankLines()

	if actualNumOfBlankLines != 0 {

		t.Errorf("%v - ERROR\n"+
			"txtBlankLinesZero.GetNumOfBlankLines()\n"+
			"Expected Number of Blank Lines = '0'\n"+
			"Instead, Number of Blank Lines = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			actualNumOfBlankLines)

		return
	}

	txtBlankLinesOne,
		err := TextLineSpecBlankLines{}.NewPtrBlankLines(
		expectedNumOfBlankLines,
		newLineChars,
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

	actualNumOfBlankLines =
		txtBlankLinesOne.GetNumOfBlankLines()

	if expectedNumOfBlankLines != actualNumOfBlankLines {

		t.Errorf("%v - ERROR\n"+
			"txtBlankLinesOne.GetNumOfBlankLines()\n"+
			"Expected Number of Blank Lines = '%v'\n"+
			"Instead, Number of Blank Lines = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedNumOfBlankLines,
			actualNumOfBlankLines)

		return
	}

	return
}

func TestTextLineSpecBlankLines_IsValidInstance_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_IsValidInstanceError_000100()",
		"")

	txtBlankLinesZero := TextLineSpecBlankLines{}

	isValid :=
		txtBlankLinesZero.IsValidInstance()

	if isValid {

		t.Errorf("%v - ERROR\n"+
			"isValid := txtBlankLinesZero.IsValidInstance()\n"+
			"Expected isValid == 'false' because \n"+
			"'txtBlankLinesZero' is empty.\n"+
			"HOWEVER, isValid=='true'!!!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	expectedNumOfBlankLines := 3

	newLineRunes := []rune{'!', '\n'}

	txtBlankLinesOne,
		err := TextLineSpecBlankLines{}.NewPtrBlankLineRunes(
		expectedNumOfBlankLines,
		newLineRunes,
		ePrefix.XCtx(
			"txtBlankLinesOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	isValid = txtBlankLinesOne.IsValidInstance()

	if !isValid {

		t.Errorf("%v - ERROR\n"+
			"isValid := txtBlankLinesOne.IsValidInstance()\n"+
			"Expected isValid == 'true' because \n"+
			"'txtBlankLinesOne' is a valid instance.\n"+
			"HOWEVER, isValid=='false'!!!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecBlankLines_IsValidInstanceError_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_IsValidInstanceError_000100()",
		"")

	txtBlankLinesZero := TextLineSpecBlankLines{}

	err :=
		txtBlankLinesZero.IsValidInstanceError(
			ePrefix.XCtx(
				"txtBlankLinesZero"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"txtBlankLinesZero.IsValidInstanceError()\n"+
			"Expected an error return from IsValidInstanceError()\n"+
			"because 'txtBlankLinesZero' is empty.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return

	}

	var txtBlankLinesOne *TextLineSpecBlankLines

	expectedNumOfBlankLines := 3

	newLineRunes := []rune{'!', '\n'}

	txtBlankLinesOne,
		err = TextLineSpecBlankLines{}.NewPtrBlankLineRunes(
		expectedNumOfBlankLines,
		newLineRunes,
		ePrefix.XCtx(
			"txtBlankLinesOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err =
		txtBlankLinesOne.IsValidInstanceError(
			ePrefix.XCtx(
				"txtBlankLinesOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err =
		txtBlankLinesOne.IsValidInstanceError(
			StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtBlankLinesOne."+
			"IsValidInstanceError()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecBlankLines_NewBlankLines_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_NewBlankLines_000100()",
		"")

	expectedNumOfBlankLines := 3

	newLineChars := "!\n"

	expectedBlankLineText :=
		strings.Repeat(
			newLineChars,
			expectedNumOfBlankLines)

	txtBlankLinesOne,
		err := TextLineSpecBlankLines{}.NewBlankLines(
		expectedNumOfBlankLines,
		newLineChars,
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

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedBlankLineText),
			true)

	var actualStr string

	actualStr,
		err = txtBlankLinesOne.GetFormattedText(
		ePrefix.XCtx(
			"txtBlankLinesOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

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

	actualNumOfBlankLines :=
		txtBlankLinesOne.GetNumOfBlankLines()

	if expectedNumOfBlankLines != actualNumOfBlankLines {

		t.Errorf("%v - ERROR\n"+
			"txtBlankLinesOne.GetNumOfBlankLines()\n"+
			"Expected Number of Blank Lines = '%v'\n"+
			"Instead, Number of Blank Lines = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedNumOfBlankLines,
			actualNumOfBlankLines)

		return
	}

	_,
		err = TextLineSpecBlankLines{}.NewBlankLines(
		expectedNumOfBlankLines,
		newLineChars,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from TextLineSpecBlankLines{}."+
			"NewBlankLines()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecBlankLines_NewBlankLineRunes_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_NewBlankLineRunes_000100()",
		"")

	numOfBlankLines := 3

	newLineRunes := []rune{'!', '\n'}

	txtBlankLinesOne,
		err := TextLineSpecBlankLines{}.NewBlankLineRunes(
		numOfBlankLines,
		newLineRunes,
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

	_,
		err = TextLineSpecBlankLines{}.NewBlankLineRunes(
		numOfBlankLines,
		newLineRunes,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from TextLineSpecBlankLines{}."+
			"NewBlankLineRunes()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	var txtBlankLinesTwo TextLineSpecBlankLines

	txtBlankLinesTwo,
		err = TextLineSpecBlankLines{}.NewBlankLineRunes(
		numOfBlankLines,
		newLineRunes,
		ePrefix.XCtx(
			"txtBlankLinesTwo"))

	expectedFmtStr :=
		strings.Repeat(string(newLineRunes), numOfBlankLines)

	var actualStr string

	actualStr,
		err = txtBlankLinesTwo.GetFormattedText(
		ePrefix.XCtx(
			"txtBlankLinesTwo"))

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
			"txtBlankLinesTwo.GetFormattedText()\n"+
			"Error: Expected Text String DOES NOT match\n"+
			"Actual Text String.\n"+
			"Expected Text String = '%v'\n"+
			"Instead, Text String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	return
}

func TestTextLineSpecBlankLines_NewBlankLineRunes_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_NewBlankLines_000100()",
		"")

	expectedNumOfBlankLines := 3

	newLineChars := "!\n"

	expectedBlankLineText :=
		strings.Repeat(
			newLineChars,
			expectedNumOfBlankLines)

	txtBlankLinesOne,
		err := TextLineSpecBlankLines{}.NewBlankLineRunes(
		expectedNumOfBlankLines,
		[]rune(newLineChars),
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

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedBlankLineText),
			true)

	actualStr := txtBlankLinesOne.String()

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

	actualNumOfBlankLines :=
		txtBlankLinesOne.GetNumOfBlankLines()

	if expectedNumOfBlankLines != actualNumOfBlankLines {

		t.Errorf("%v - ERROR\n"+
			"txtBlankLinesOne.GetNumOfBlankLines()\n"+
			"Expected Number of Blank Lines = '%v'\n"+
			"Instead, Number of Blank Lines = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedNumOfBlankLines,
			actualNumOfBlankLines)

		return
	}

	_,
		err = TextLineSpecBlankLines{}.NewBlankLineRunes(
		expectedNumOfBlankLines,
		[]rune(newLineChars),
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from TextLineSpecBlankLines{}."+
			"NewBlankLineRunes()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return

}

func TestTextLineSpecBlankLines_NewDefaultBlankLines_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_NewDefaultBlankLines_000100()",
		"")

	expectedNumOfBlankLines := 3

	newLineChars := "\n"

	expectedBlankLineText :=
		strings.Repeat(
			newLineChars,
			expectedNumOfBlankLines)

	txtBlankLinesOne,
		err := TextLineSpecBlankLines{}.NewDefaultBlankLines(
		expectedNumOfBlankLines,
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

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedBlankLineText),
			true)

	var actualStr string

	actualStr,
		err = txtBlankLinesOne.GetFormattedText(
		ePrefix.XCtx(
			"txtBlankLinesOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

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

	actualNumOfBlankLines :=
		txtBlankLinesOne.GetNumOfBlankLines()

	if expectedNumOfBlankLines != actualNumOfBlankLines {

		t.Errorf("%v - ERROR\n"+
			"txtBlankLinesOne.GetNumOfBlankLines()\n"+
			"Expected Number of Blank Lines = '%v'\n"+
			"Instead, Number of Blank Lines = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedNumOfBlankLines,
			actualNumOfBlankLines)

		return
	}

	_,
		err = TextLineSpecBlankLines{}.NewDefaultBlankLines(
		expectedNumOfBlankLines,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from TextLineSpecBlankLines{}."+
			"NewDefaultBlankLines()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecBlankLines_NewPtrBlankLines_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_NewPtrBlankLines_000100()",
		"")

	expectedNumOfBlankLines := 3

	newLineChars := "!\n"

	expectedBlankLineText :=
		strings.Repeat(
			newLineChars,
			expectedNumOfBlankLines)

	txtBlankLinesOne,
		err := TextLineSpecBlankLines{}.NewPtrBlankLines(
		expectedNumOfBlankLines,
		newLineChars,
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

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedBlankLineText),
			true)

	var actualStr string

	actualStr,
		err = txtBlankLinesOne.GetFormattedText(
		ePrefix.XCtx(
			"txtBlankLinesOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

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

	actualNumOfBlankLines :=
		txtBlankLinesOne.GetNumOfBlankLines()

	if expectedNumOfBlankLines != actualNumOfBlankLines {

		t.Errorf("%v - ERROR\n"+
			"txtBlankLinesOne.GetNumOfBlankLines()\n"+
			"Expected Number of Blank Lines = '%v'\n"+
			"Instead, Number of Blank Lines = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedNumOfBlankLines,
			actualNumOfBlankLines)

		return
	}

	_,
		err = TextLineSpecBlankLines{}.NewPtrBlankLines(
		expectedNumOfBlankLines,
		newLineChars,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from TextLineSpecBlankLines{}."+
			"NewPtrBlankLines()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	var txtBlankLinesTwo *TextLineSpecBlankLines

	txtBlankLinesTwo,
		err = TextLineSpecBlankLines{}.NewPtrBlankLines(
		expectedNumOfBlankLines,
		newLineChars,
		ePrefix.XCtx(
			"txtBlankLinesTwo"))

	expectedFmtStr :=
		strings.Repeat(
			newLineChars,
			expectedNumOfBlankLines)

	actualStr,
		err = txtBlankLinesTwo.GetFormattedText(
		ePrefix.XCtx(
			"txtBlankLinesTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	printableExpectedStr =
		sMech.ConvertNonPrintableChars(
			[]rune(expectedFmtStr),
			true)

	printableActualStr =
		sMech.ConvertNonPrintableChars(
			[]rune(actualStr),
			true)

	if printableExpectedStr != printableActualStr {
		t.Errorf("%v\n"+
			"txtBlankLinesTwo.GetFormattedText()\n"+
			"Error: Expected Text String DOES NOT match\n"+
			"Actual Text String.\n"+
			"Expected Text String = '%v'\n"+
			"Instead, Text String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	return
}

func TestTextLineSpecBlankLines_NewPtrDefaultBlankLines_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_NewPtrDefaultBlankLines_000100()",
		"")

	numOfBlankLines := 3

	txtBlankLinesOne,
		err := TextLineSpecBlankLines{}.NewPtrDefaultBlankLines(
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

	_,
		err = TextLineSpecBlankLines{}.NewPtrDefaultBlankLines(
		numOfBlankLines,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from TextLineSpecBlankLines{}."+
			"NewPtrDefaultBlankLines()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	var txtBlankLinesTwo *TextLineSpecBlankLines

	txtBlankLinesTwo,
		err = TextLineSpecBlankLines{}.NewPtrDefaultBlankLines(
		numOfBlankLines,
		ePrefix.XCtx(
			"txtBlankLinesTwo"))

	expectedFmtStr :=
		strings.Repeat("\n", numOfBlankLines)

	var actualStr string

	actualStr,
		err = txtBlankLinesTwo.GetFormattedText(
		ePrefix.XCtx(
			"txtBlankLinesTwo"))

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
			"txtBlankLinesTwo.GetFormattedText()\n"+
			"Error: Expected Text String DOES NOT match\n"+
			"Actual Text String.\n"+
			"Expected Text String = '%v'\n"+
			"Instead, Text String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	return
}

func TestTextLineSpecBlankLines_NewPtrRunesBlankLines_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_NewPtrRunesBlankLines_000100()",
		"")

	numOfBlankLines := 3

	newLineRunes := []rune{'!', '\n'}

	txtBlankLinesOne,
		err := TextLineSpecBlankLines{}.NewPtrBlankLineRunes(
		numOfBlankLines,
		newLineRunes,
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

	_,
		err = TextLineSpecBlankLines{}.NewPtrBlankLineRunes(
		numOfBlankLines,
		newLineRunes,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from TextLineSpecBlankLines{}."+
			"NewPtrBlankLineRunes()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	var txtBlankLinesTwo *TextLineSpecBlankLines

	txtBlankLinesTwo,
		err = TextLineSpecBlankLines{}.NewPtrBlankLineRunes(
		numOfBlankLines,
		newLineRunes,
		ePrefix.XCtx(
			"txtBlankLinesTwo"))

	expectedFmtStr :=
		strings.Repeat(string(newLineRunes), numOfBlankLines)

	var actualStr string

	actualStr,
		err = txtBlankLinesTwo.GetFormattedText(
		ePrefix.XCtx(
			"txtBlankLinesTwo"))

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
			"txtBlankLinesTwo.GetFormattedText()\n"+
			"Error: Expected Text String DOES NOT match\n"+
			"Actual Text String.\n"+
			"Expected Text String = '%v'\n"+
			"Instead, Text String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	return
}

func TestTextLineSpecBlankLines_Read_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_Read_000100()",
		"")

	numOfBlankLines := 3

	newLineRunes := []rune{'!', '\n'}

	txtBlankLinesOne,
		err := TextLineSpecBlankLines{}.NewPtrBlankLineRunes(
		numOfBlankLines,
		newLineRunes,
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

	expectedTextStr :=
		strings.Repeat(string(newLineRunes), numOfBlankLines)

	lenExpectedStr := len(expectedTextStr)

	p := make([]byte, lenExpectedStr+1)

	var n, readBytesCnt int
	var actualStr string

	for {

		n,
			err = txtBlankLinesOne.Read(p)

		if n == 0 {
			break
		}

		actualStr += string(p[:n])
		readBytesCnt += n
	}

	if err != nil &&
		err != io.EOF {
		t.Errorf("%v\n"+
			"Error Returned From txtBlankLinesOne.Read(p)\n"+
			"Error = \n%v\n",
			ePrefix.XCtxEmpty().String(),
			err.Error())

		return
	}

	if err == nil {
		t.Errorf("%v\n"+
			"Error: After completing Read Operation\n"+
			"the returned error should equal io.EOF.\n"+
			"HOWEVER, returned error == nil!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	if err != io.EOF {
		t.Errorf("%v\n"+
			"Error: After completing Read Operation\n"+
			"the returned error should equal io.EOF.\n"+
			"HOWEVER, returned error is NOT equal io.EOF!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	if txtBlankLinesOne.textLineReader != nil {
		t.Errorf("%v\n"+
			"Error: After completing Read Operation\n"+
			"txtBlankLinesOne.textLineReader != 'nil'\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	if readBytesCnt != lenExpectedStr {
		t.Errorf("%v\n"+
			"Byte Length Error: txtBlankLinesOne.Read(p)\n"+
			"The actual length of bytes read\n"+
			"does NOT match the expected length.\n"+
			"Expected Bytes Read = '%v'\n"+
			"       Actual Bytes = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			lenExpectedStr,
			readBytesCnt)

		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedTextStr),
			true)

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(actualStr),
			true)

	if printableExpectedStr != printableActualStr {
		t.Errorf("%v\n"+
			"Error: Expected Text String DOES NOT match\n"+
			"Actual Text String.\n"+
			"Expected Text String = '%v'\n"+
			"Instead, Text String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedStr,
			printableActualStr)
	}

	return
}

func TestTextLineSpecBlankLines_Read_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_Read_000200()",
		"")

	numOfBlankLines := 3

	newLineRunes := []rune{'!', '\n'}

	txtBlankLinesOne,
		err := TextLineSpecBlankLines{}.NewPtrBlankLineRunes(
		numOfBlankLines,
		newLineRunes,
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

	expectedTextStr :=
		strings.Repeat(string(newLineRunes), numOfBlankLines)

	lenExpectedStr := len(expectedTextStr)

	p := make([]byte, 1)

	var actualStr string

	var n, readBytesCnt int

	for {

		n,
			err = txtBlankLinesOne.Read(p)

		if n == 0 {
			break
		}

		actualStr += string(p[:n])
		readBytesCnt += n

	}

	if err != nil &&
		err != io.EOF {
		t.Errorf("%v\n"+
			"Error Returned From txtBlankLinesOne.Read(p)\n"+
			"Error = \n%v\n",
			ePrefix.XCtxEmpty().String(),
			err.Error())

		return
	}

	if err != nil &&
		err != io.EOF {
		t.Errorf("%v\n"+
			"Error Returned From txtBlankLinesOne.Read(p)\n"+
			"Error = \n%v\n",
			ePrefix.XCtxEmpty().String(),
			err.Error())

		return
	}

	if err == nil {
		t.Errorf("%v\n"+
			"Error: After completing Read Operation\n"+
			"the returned error should equal io.EOF.\n"+
			"HOWEVER, returned error == nil!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	if err != io.EOF {
		t.Errorf("%v\n"+
			"Error: After completing Read Operation\n"+
			"the returned error should equal io.EOF.\n"+
			"HOWEVER, returned error is NOT equal io.EOF!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	if readBytesCnt != lenExpectedStr {
		t.Errorf("%v\n"+
			"Byte Length Error: txtBlankLinesOne.Read(p)\n"+
			"The actual length of bytes read\n"+
			"does NOT match the expected length.\n"+
			"Expected Bytes Read = '%v'\n"+
			"       Actual Bytes = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			lenExpectedStr,
			readBytesCnt)

		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedTextStr),
			true)

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(actualStr),
			true)

	if printableExpectedStr != printableActualStr {
		t.Errorf("%v\n"+
			"Error: Expected Text String DOES NOT match\n"+
			"Actual Text String.\n"+
			"Expected Text String = '%v'\n"+
			"Instead, Text String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedStr,
			printableActualStr)
		return
	}

	if txtBlankLinesOne.textLineReader != nil {
		t.Errorf("%v\n"+
			"Error: After a successful series of byte reads,\n"+
			"txtBlankLinesOne.textLineReader pointer has NOT\n"+
			"BEEN RESET TO 'nil'!\n",
			ePrefix.XCtxEmpty().String())
		return
	}

	p = make([]byte, 100)
	readBytesCnt = 0
	actualStr = ""

	for {

		n,
			err = txtBlankLinesOne.Read(p)

		if n == 0 {
			break
		}

		actualStr += string(p[:n])
		readBytesCnt += n
	}

	if err != nil &&
		err != io.EOF {
		t.Errorf("%v\n"+
			"Error: Test # 2\n"+
			"Error Returned From txtBlankLinesOne.Read(p)\n"+
			"Error = \n%v\n",
			ePrefix.XCtxEmpty().String(),
			err.Error())

		return
	}

	printableActualStr =
		sMech.ConvertNonPrintableChars(
			[]rune(actualStr),
			true)

	if printableExpectedStr != printableActualStr {
		t.Errorf("%v\n"+
			"Error: Test # 2\n"+
			"Expected Text String DOES NOT match\n"+
			"Actual Text String.\n"+
			"Expected Text String = '%v'\n"+
			"Instead, Text String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedStr,
			printableActualStr)
		return
	}

	return
}

func TestTextLineSpecBlankLines_Read_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_Read_000300()",
		"")

	numOfBlankLines := 3

	newLineRunes := []rune{'!', '\n'}

	txtBlankLinesOne,
		err := TextLineSpecBlankLines{}.NewPtrBlankLineRunes(
		numOfBlankLines,
		newLineRunes,
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

	expectedTextStr :=
		strings.Repeat(string(newLineRunes), numOfBlankLines)

	lenExpectedStr := len(expectedTextStr)

	txtSpecAtom := textSpecificationAtom{}

	var n int
	p := make([]byte, 100)

	n,
		err = txtSpecAtom.readBytes(
		nil,
		p,
		ePrefix.XCtx("plainTextLine == 'nil'"))

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected error return from txtSpecAtom.readBytes()"+
			"because input parameter 'plainTextLine' == 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	var formattedTxtStr string
	blankLinesMolecule := textLineSpecBlankLinesMolecule{}

	formattedTxtStr,
		err =
		blankLinesMolecule.getFormattedText(
			txtBlankLinesOne,
			ePrefix.XCtx("txtBlankLinesOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	p = make([]byte, 0)

	txtBlankLinesOne.textLineReader =
		strings.NewReader(formattedTxtStr)

	n,
		err = txtSpecAtom.readBytes(
		txtBlankLinesOne.textLineReader,
		p,
		ePrefix.XCtx("p == zero length"))

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected error return from txtSpecAtom.readBytes()"+
			"because input parameter 'p' is a zero length byte array.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	p = make([]byte, 100)

	var readBytesCnt int
	var actualStr string

	for {

		n,
			err = txtSpecAtom.readBytes(
			txtBlankLinesOne.textLineReader,
			p,
			ePrefix.XCtx("txtBlankLinesOne is valid"))

		if n == 0 {
			break
		}

		actualStr += string(p[:n])
		readBytesCnt += n
	}

	if err != nil &&
		err != io.EOF {
		t.Errorf("%v\n"+
			"Error Returned From txtSpecAtom.readBytes(p)\n"+
			"Error = \n%v\n",
			ePrefix.XCtxEmpty().String(),
			err.Error())

		return
	}

	if readBytesCnt != lenExpectedStr {
		t.Errorf("%v\n"+
			"Byte Length Error: txtSpecAtom.readBytes(p)\n"+
			"The actual length of bytes read\n"+
			"does NOT match the expected length.\n"+
			"Expected Bytes Read = '%v'\n"+
			"       Actual Bytes = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			lenExpectedStr,
			readBytesCnt)

		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedTextStr),
			true)

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(actualStr),
			true)

	if printableExpectedStr != printableActualStr {
		t.Errorf("%v\n"+
			"Error: Expected Text String DOES NOT match\n"+
			"Actual Text String.\n"+
			"Expected Text String = '%v'\n"+
			"Instead, Text String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	txtBlankLinesTwo := TextLineSpecBlankLines{}

	p = make([]byte, lenExpectedStr+1)

	_,
		err = txtBlankLinesTwo.Read(p)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from "+
			"txtBlankLinesTwo.Read(p)\n"+
			"because parameter 'txtBlankLinesTwo' is "+
			"empty and invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtx("Missing Error Return"))

		return
	}

	return
}

func TestTextLineSpecBlankLines_ReaderInitialize_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_ReaderInitialize_000100()",
		"")

	numOfBlankLines := 3

	newLineRunes := []rune{'!', '\n'}

	txtBlankLinesOne,
		err := TextLineSpecBlankLines{}.NewPtrBlankLineRunes(
		numOfBlankLines,
		newLineRunes,
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

	expectedTextStr :=
		strings.Repeat(string(newLineRunes), numOfBlankLines)

	lenExpectedStr := len(expectedTextStr)

	p := make([]byte, 5)

	var n int

	n,
		err = txtBlankLinesOne.Read(p)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by txtBlankLinesOne.Read(p)\n"+
			"Error:\n%v\n",
			ePrefix.XCtxEmpty().String(),
			err.Error())

		return
	}

	if n != 5 {
		t.Errorf("%v\n"+
			"Error: txtBlankLinesOne.Read(p)\n"+
			"Expected n == 5\n"+
			"Instead, n == %v\n",
			ePrefix.XCtxEmpty().String(),
			n)

		return
	}

	p = make([]byte, 100)

	txtBlankLinesOne.ReaderInitialize()

	var readBytesCnt int
	var actualStr string

	for {

		n,
			err = txtBlankLinesOne.Read(p)

		if n == 0 {
			break
		}

		actualStr += string(p[:n])
		readBytesCnt += n
	}

	if err != nil &&
		err != io.EOF {
		t.Errorf("%v\n"+
			"Error Returned From txtBlankLinesOne.Read(p)\n"+
			"Error = \n%v\n",
			ePrefix.XCtxEmpty().String(),
			err.Error())

		return
	}

	if err == nil {
		t.Errorf("%v\n"+
			"Error: After completing Read Operation\n"+
			"the returned error should equal io.EOF.\n"+
			"HOWEVER, returned error == nil!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	if err != io.EOF {
		t.Errorf("%v\n"+
			"Error: After completing Read Operation\n"+
			"the returned error should equal io.EOF.\n"+
			"HOWEVER, returned error is NOT equal io.EOF!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	if txtBlankLinesOne.textLineReader != nil {
		t.Errorf("%v\n"+
			"Error: After completing Read Operation\n"+
			"txtBlankLinesOne.textLineReader != 'nil'\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	if readBytesCnt != lenExpectedStr {
		t.Errorf("%v\n"+
			"Byte Length Error: txtBlankLinesOne.Read(p)\n"+
			"The actual length of bytes read\n"+
			"does NOT match the expected length.\n"+
			"Expected Bytes Read = '%v'\n"+
			"       Actual Bytes = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			lenExpectedStr,
			readBytesCnt)

		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedTextStr),
			true)

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(actualStr),
			true)

	if printableExpectedStr != printableActualStr {
		t.Errorf("%v\n"+
			"Error: Expected Text String DOES NOT match\n"+
			"Actual Text String.\n"+
			"Expected Text String = '%v'\n"+
			"Instead, Text String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	if txtBlankLinesOne.textLineReader != nil {
		t.Errorf("%v Test #1\n"+
			"Completed Read Operation but txtBlankLinesOne.textLineReader\n"+
			"is NOT equal to 'nil'!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	p = make([]byte, 100)
	actualStr = ""
	readBytesCnt = 0

	for {

		n,
			err = txtBlankLinesOne.Read(p)

		if n == 0 {
			break
		}

		actualStr += string(p[:n])
		readBytesCnt += n
	}

	if err != nil &&
		err != io.EOF {
		t.Errorf("%v\n"+
			"Test # 2"+
			"Error Returned From txtBlankLinesOne.Read(p)\n"+
			"Error = \n%v\n",
			ePrefix.XCtxEmpty().String(),
			err.Error())

		return
	}

	if err == nil {
		t.Errorf("%v\n"+
			"Test # 2"+
			"Error: After completing Read Operation\n"+
			"the returned error should equal io.EOF.\n"+
			"HOWEVER, returned error == nil!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	if err != io.EOF {
		t.Errorf("%v\n"+
			"Test # 2"+
			"Error: After completing Read Operation\n"+
			"the returned error should equal io.EOF.\n"+
			"HOWEVER, returned error is NOT equal io.EOF!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	printableActualStr =
		sMech.ConvertNonPrintableChars(
			[]rune(actualStr),
			true)

	if printableExpectedStr != printableActualStr {
		t.Errorf("%v Test #2\n"+
			"Error: Expected Text String DOES NOT match\n"+
			"Actual Text String.\n"+
			"Expected Text String = '%v'\n"+
			"Instead, Text String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	if txtBlankLinesOne.textLineReader != nil {
		t.Errorf("%v Test #2\n"+
			"Completed Read Operation but plainTextLine01.textLineReader\n"+
			"is NOT equal to 'nil'!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	txtBlankLinesTwo := TextLineSpecBlankLines{}

	txtBlankLinesTwo.ReaderInitialize()

	return
}

func TestTextLineSpecBlankLines_SetNewLineChars_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_SetNewLineChars_000100()",
		"")

	numOfBlankLines := 3

	expectedNewLineChars := "!\n"

	expectedNewLineRunes := []rune(expectedNewLineChars)

	txtBlankLinesZero := TextLineSpecBlankLines{}

	err :=
		txtBlankLinesZero.SetNewLineChars(
			expectedNewLineChars,
			ePrefix.XCtx(
				"txtBlankLinesZero"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	_,
		err = txtBlankLinesZero.GetFormattedText(
		ePrefix.XCtx(
			"Empty txtBlankLinesZero"))

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from "+
			"txtBlankLinesZero.GetFormattedText()\n"+
			"because 'txtBlankLinesZero' is "+
			"zero and invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtx("Missing Error Return"))

		return
	}

	actualNewLineChars :=
		txtBlankLinesZero.GetNewLineChars()

	if actualNewLineChars != expectedNewLineChars {

		t.Errorf("%v - ERROR\n"+
			"txtBlankLinesZero.GetNewLineChars()\n"+
			"Expected expectedNewLineChars == actualNewLineChars\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"expectedNewLineChars = '%v'\n"+
			"  actualNewLineChars = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedNewLineChars,
			actualNewLineChars)

		return
	}

	actualNewLineRunes :=
		txtBlankLinesZero.GetNewLineRunes()

	areEqual := strMechPreon{}.ptr().equalRuneArrays(
		expectedNewLineRunes,
		actualNewLineRunes)

	if !areEqual {

		t.Errorf("%v - ERROR\n"+
			"txtBlankLinesZero.GetNewLineRunes()\n"+
			"Expected expectedNewLineRunes == actualNewLineRunes\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"expectedNewLineRunes = '%v'\n"+
			"  actualNewLineRunes = '%v'\n"+
			"  actualNewLineRunes = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			string(expectedNewLineRunes),
			string(actualNewLineRunes),
			actualNewLineRunes)

		return
	}

	var txtBlankLinesOne TextLineSpecBlankLines

	txtBlankLinesOne,
		err = TextLineSpecBlankLines{}.NewBlankLines(
		numOfBlankLines,
		expectedNewLineChars,
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

	actualNewLineChars =
		txtBlankLinesOne.GetNewLineChars()

	if actualNewLineChars != expectedNewLineChars {

		t.Errorf("%v - ERROR\n"+
			"txtBlankLinesOne.GetNewLineChars()\n"+
			"Expected expectedNewLineChars == actualNewLineChars\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"expectedNewLineChars = '%v'\n"+
			"  actualNewLineChars = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedNewLineChars,
			actualNewLineChars)

		return
	}

	expectedNewLineChars = "**!\n"

	err =
		txtBlankLinesOne.SetNewLineChars(
			expectedNewLineChars,
			ePrefix.XCtx(
				"txtBlankLinesOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualNewLineChars =
		txtBlankLinesOne.GetNewLineChars()

	if actualNewLineChars != expectedNewLineChars {

		t.Errorf("%v - ERROR\n"+
			"Test # 2\n"+
			"txtBlankLinesOne.GetNewLineChars()\n"+
			"Expected expectedNewLineChars == actualNewLineChars\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"expectedNewLineChars = '%v'\n"+
			"  actualNewLineChars = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedNewLineChars,
			actualNewLineChars)

		return
	}

	expectedFmtStr :=
		strings.Repeat(expectedNewLineChars, numOfBlankLines)

	var actualFmtStr string

	actualFmtStr,
		err =
		txtBlankLinesOne.GetFormattedText(
			ePrefix.XCtx(
				"txtBlankLinesOne-Revised"))

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedFmtStr),
			true)

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(actualFmtStr),
			true)

	if printableExpectedStr != printableActualStr {
		t.Errorf("%v - ERROR\n"+
			"txtBlankLinesOne.GetFormattedText()\n"+
			"Expected Text String DOES NOT match\n"+
			"Actual Text String.\n"+
			"Expected Text String = '%v'\n"+
			"Instead, Text String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	err =
		txtBlankLinesOne.SetNewLineChars(
			"xxx\n",
			StrMech{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtBlankLinesOne."+
			"SetNewLineChars()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	actualNewLineChars =
		txtBlankLinesOne.GetNewLineChars()

	if actualNewLineChars != expectedNewLineChars {

		t.Errorf("%v - ERROR\n"+
			"Test # 3\n"+
			"txtBlankLinesOne.GetNewLineChars()\n"+
			"Expected expectedNewLineChars == actualNewLineChars\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"expectedNewLineChars = '%v'\n"+
			"  actualNewLineChars = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedNewLineChars,
			actualNewLineChars)

		return
	}

	var txtBlankLinesTwo TextLineSpecBlankLines

	txtBlankLinesTwo,
		err = TextLineSpecBlankLines{}.NewBlankLines(
		numOfBlankLines,
		expectedNewLineChars,
		ePrefix.XCtx(
			"txtBlankLinesTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualNewLineChars = ""

	err = txtBlankLinesTwo.SetNewLineChars(
		actualNewLineChars,
		ePrefix.XCtx(
			"txtBlankLinesTwo"))

	if err == nil {

		t.Errorf("%v\n"+
			"Expected an error return from "+
			"txtBlankLinesTwo.SetNewLineChars()\n"+
			"because input parameter 'actualNewLineChars' is "+
			"empty and invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtx("Missing Error Return"))

		return
	}

	return
}

func TestTextLineSpecBlankLines_SetNewLineRunes_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_SetNewLineRunes_000100()",
		"")

	numOfBlankLines := 3

	expectedNewLineChars := "!\n"

	expectedNewLineRunes := []rune(expectedNewLineChars)

	txtBlankLinesZero := TextLineSpecBlankLines{}

	err :=
		txtBlankLinesZero.SetNewLineRunes(
			expectedNewLineRunes,
			ePrefix.XCtx(
				"txtBlankLinesZero"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	_,
		err = txtBlankLinesZero.GetFormattedText(
		ePrefix.XCtx(
			"Empty txtBlankLinesZero"))

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from "+
			"txtBlankLinesZero.GetFormattedText()\n"+
			"because 'txtBlankLinesZero' is "+
			"zero and invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtx("Missing Error Return"))

		return
	}

	actualNewLineChars :=
		txtBlankLinesZero.GetNewLineChars()

	if actualNewLineChars != expectedNewLineChars {

		t.Errorf("%v - ERROR\n"+
			"txtBlankLinesZero.GetNewLineChars()\n"+
			"Expected expectedNewLineChars == actualNewLineChars\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"expectedNewLineChars = '%v'\n"+
			"  actualNewLineChars = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedNewLineChars,
			actualNewLineChars)

		return
	}

	actualNewLineRunes :=
		txtBlankLinesZero.GetNewLineRunes()

	sMechPreon := strMechPreon{}

	areEqual := sMechPreon.equalRuneArrays(
		expectedNewLineRunes,
		actualNewLineRunes)

	if !areEqual {

		t.Errorf("%v - ERROR\n"+
			"txtBlankLinesZero.GetNewLineRunes()\n"+
			"Expected expectedNewLineRunes == actualNewLineRunes\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"expectedNewLineRunes = '%v'\n"+
			"  actualNewLineRunes = '%v'\n"+
			"  actualNewLineRunes = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			string(expectedNewLineRunes),
			string(actualNewLineRunes),
			actualNewLineRunes)

		return
	}

	var txtBlankLinesOne TextLineSpecBlankLines

	txtBlankLinesOne,
		err = TextLineSpecBlankLines{}.NewBlankLineRunes(
		numOfBlankLines,
		expectedNewLineRunes,
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

	actualNewLineChars =
		txtBlankLinesOne.GetNewLineChars()

	if actualNewLineChars != expectedNewLineChars {

		t.Errorf("%v - ERROR\n"+
			"txtBlankLinesOne.GetNewLineChars()\n"+
			"Expected expectedNewLineChars == actualNewLineChars\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"expectedNewLineChars = '%v'\n"+
			"  actualNewLineChars = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedNewLineChars,
			actualNewLineChars)

		return
	}

	expectedNewLineChars = "**!\n"
	expectedNewLineRunes = []rune(expectedNewLineChars)

	err =
		txtBlankLinesOne.SetNewLineRunes(
			expectedNewLineRunes,
			ePrefix.XCtx(
				"txtBlankLinesOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualNewLineChars =
		txtBlankLinesOne.GetNewLineChars()

	if actualNewLineChars != expectedNewLineChars {

		t.Errorf("%v - ERROR\n"+
			"Test # 2\n"+
			"txtBlankLinesOne.GetNewLineChars()\n"+
			"Expected expectedNewLineChars == actualNewLineChars\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"expectedNewLineChars = '%v'\n"+
			"  actualNewLineChars = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedNewLineChars,
			actualNewLineChars)

		return
	}

	expectedFmtStr :=
		strings.Repeat(expectedNewLineChars, numOfBlankLines)

	var actualFmtStr string

	actualFmtStr,
		err =
		txtBlankLinesOne.GetFormattedText(
			ePrefix.XCtx(
				"txtBlankLinesOne-Revised"))

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedFmtStr),
			true)

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(actualFmtStr),
			true)

	if printableExpectedStr != printableActualStr {
		t.Errorf("%v - ERROR\n"+
			"txtBlankLinesOne.GetFormattedText()\n"+
			"Expected Text String DOES NOT match\n"+
			"Actual Text String.\n"+
			"Expected Text String = '%v'\n"+
			"Instead, Text String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	err =
		txtBlankLinesOne.SetNewLineRunes(
			[]rune("xxx\n"),
			StrMech{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtBlankLinesOne."+
			"SetNewLineRunes()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	actualNewLineChars =
		txtBlankLinesOne.GetNewLineChars()

	if actualNewLineChars != expectedNewLineChars {

		t.Errorf("%v - ERROR\n"+
			"Test # 3\n"+
			"txtBlankLinesOne.GetNewLineChars()\n"+
			"Expected expectedNewLineChars == actualNewLineChars\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"expectedNewLineChars = '%v'\n"+
			"  actualNewLineChars = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedNewLineChars,
			actualNewLineChars)

		return
	}

	actualNewLineRunes =
		txtBlankLinesOne.GetNewLineRunes()

	areEqual = sMechPreon.equalRuneArrays(
		expectedNewLineRunes,
		actualNewLineRunes)

	if !areEqual {

		t.Errorf("%v - ERROR - Test #3\n"+
			"txtBlankLinesZero.GetNewLineRunes()\n"+
			"Expected expectedNewLineRunes == actualNewLineRunes\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"expectedNewLineRunes = '%v'\n"+
			"  actualNewLineRunes = '%v'\n"+
			"  actualNewLineRunes = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			string(expectedNewLineRunes),
			string(actualNewLineRunes),
			actualNewLineRunes)

		return
	}

	txtBlankLinesTwo := TextLineSpecBlankLines{}

	err = txtBlankLinesTwo.SetNewLineRunes(
		nil,
		ePrefix.XCtx(
			"txtBlankLinesTwo"))

	if err == nil {

		t.Errorf("%v\n"+
			"Expected an error return from "+
			"txtBlankLinesTwo.SetNewLineRunes()\n"+
			"because input parameter 'newLineRunes' is "+
			"'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtx("Missing Error Return"))

		return
	}

	return
}

func TestTextLineSpecBlankLines_SetNumberOfBlankLines_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_SetNumberOfBlankLines_000100()",
		"")

	numOfBlankLines := 3

	txtBlankLinesOne,
		err := TextLineSpecBlankLines{}.NewPtrDefaultBlankLines(
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

	err = txtBlankLinesOne.SetNumberOfBlankLines(
		-99,
		ePrefix.XCtx(
			"txtBlankLinesOne"))

	if err == nil {

		t.Errorf("%v\n"+
			"Expected an error return from "+
			"txtBlankLinesTwo.SetNumberOfBlankLines()\n"+
			"because input parameter 'newOfBlankLines' is "+
			"invalid (-99).\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtx("Missing Error Return"))

		return
	}

	err = txtBlankLinesOne.SetNumberOfBlankLines(
		5,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtBlankLinesOne."+
			"SetNumberOfBlankLines()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecBlankLines_SetSpecBlankLines_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_SetSpecBlankLines_000100()",
		"")

	numOfBlankLines := 3

	newLineChars := "!?\n"

	expectedFmtStr :=
		strings.Repeat(newLineChars, numOfBlankLines)

	txtBlankLinesOne := TextLineSpecBlankLines{}

	err :=
		txtBlankLinesOne.SetSpecBlankLines(
			numOfBlankLines,
			newLineChars,
			ePrefix.XCtx("txtBlankLinesOne"))

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

	var txtBlankLinesTwo TextLineSpecBlankLines

	txtBlankLinesTwo,
		err = TextLineSpecBlankLines{}.NewBlankLines(
		numOfBlankLines,
		newLineChars,
		ePrefix.XCtx(
			"txtBlankLinesTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtBlankLinesTwo.IsValidInstanceError(
		ePrefix.XCtx(
			"txtBlankLinesTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !txtBlankLinesOne.Equal(&txtBlankLinesTwo) {
		t.Errorf("%v - ERROR\n"+
			"Expected txtBlankLinesOne==txtBlankLinesTwo\n"+
			"HOWEVER, THESE INSTANCES ARE NOT EQUAL!\n",
			ePrefix.XCtxEmpty().String())
		return
	}

	var actualStr string

	actualStr,
		err = txtBlankLinesOne.GetFormattedText(
		ePrefix.XCtx(
			"txtBlankLinesOne"))

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
			"txtBlankLinesOne.GetFormattedText()\n"+
			"Error: Expected Text String DOES NOT match\n"+
			"Actual Text String.\n"+
			"Expected Text String = '%v'\n"+
			"Instead, Text String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	txtBlankLinesThree := TextLineSpecBlankLines{}

	numOfBlankLines = 1000001

	newLineChars = "!?\n"

	err =
		txtBlankLinesThree.SetSpecBlankLines(
			numOfBlankLines,
			newLineChars,
			ePrefix.XCtx("txtBlankLinesThree"))

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Test #1 - numOfBlankLines = 1000001\n"+
			"Expected an error return from txtBlankLinesThree."+
			"SetSpecBlankLines()\n"+
			"because 'numOfBlankLines' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	txtBlankLinesThreeB := TextLineSpecBlankLines{}

	numOfBlankLines = -99

	newLineChars = "!?\n"

	err =
		txtBlankLinesThreeB.SetSpecBlankLines(
			numOfBlankLines,
			newLineChars,
			ePrefix.XCtx("txtBlankLinesThree"))

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Test #2 - numOfBlankLines = -99\n"+
			"Expected an error return from txtBlankLinesThreeB."+
			"SetSpecBlankLines()\n"+
			"because 'numOfBlankLines' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	txtBlankLinesFour := TextLineSpecBlankLines{}

	numOfBlankLines = 3

	newLineChars = ""

	err =
		txtBlankLinesFour.SetSpecBlankLines(
			numOfBlankLines,
			newLineChars,
			ePrefix.XCtx("txtBlankLinesFour"))

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtBlankLinesThree."+
			"SetSpecBlankLines()\n"+
			"because 'newLineChars' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	txtBlankLinesFive := TextLineSpecBlankLines{}

	numOfBlankLines = 3

	newLineChars = "   \n"

	err =
		txtBlankLinesFive.SetSpecBlankLines(
			numOfBlankLines,
			newLineChars,
			StrMech{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtBlankLinesThree."+
			"SetSpecBlankLines()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecBlankLines_SetSpecDefaultBlankLines_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_SetSpecDefaultBlankLines_000100()",
		"")

	numOfBlankLines := 3

	newLineChars := "\n"

	expectedFmtStr :=
		strings.Repeat(newLineChars, numOfBlankLines)

	txtBlankLinesOne := TextLineSpecBlankLines{}

	err :=
		txtBlankLinesOne.SetSpecDefaultBlankLines(
			numOfBlankLines,
			ePrefix.XCtx("txtBlankLinesOne"))

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

	var txtBlankLinesTwo TextLineSpecBlankLines

	txtBlankLinesTwo,
		err = TextLineSpecBlankLines{}.NewBlankLines(
		numOfBlankLines,
		newLineChars,
		ePrefix.XCtx(
			"txtBlankLinesTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtBlankLinesTwo.IsValidInstanceError(
		ePrefix.XCtx(
			"txtBlankLinesTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !txtBlankLinesOne.Equal(&txtBlankLinesTwo) {
		t.Errorf("%v - ERROR\n"+
			"Expected txtBlankLinesOne==txtBlankLinesTwo\n"+
			"HOWEVER, THESE INSTANCES ARE NOT EQUAL!\n",
			ePrefix.XCtxEmpty().String())
		return
	}

	var actualStr string

	actualStr,
		err = txtBlankLinesOne.GetFormattedText(
		ePrefix.XCtx(
			"txtBlankLinesOne"))

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
			"txtBlankLinesOne.GetFormattedText()\n"+
			"Error: Expected Text String DOES NOT match\n"+
			"Actual Text String.\n"+
			"Expected Text String = '%v'\n"+
			"Instead, Text String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	txtBlankLinesThree := TextLineSpecBlankLines{}

	numOfBlankLines = -50

	err =
		txtBlankLinesThree.SetSpecDefaultBlankLines(
			numOfBlankLines,
			ePrefix.XCtx("txtBlankLinesThree"))

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Test #1 'numOfBlankLines' = -50\n"+
			"Expected an error return from txtBlankLinesThree."+
			"SetSpecDefaultBlankLines()\n"+
			"because 'numOfBlankLines' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	txtBlankLinesFour := TextLineSpecBlankLines{}

	numOfBlankLines = 1000005

	err =
		txtBlankLinesFour.SetSpecDefaultBlankLines(
			numOfBlankLines,
			ePrefix.XCtx("txtBlankLinesFour"))

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Test #2 'numOfBlankLines' = 1000005\n"+
			"Expected an error return from txtBlankLinesThree."+
			"SetSpecDefaultBlankLines()\n"+
			"because 'numOfBlankLines' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	txtBlankLinesFive := TextLineSpecBlankLines{}

	numOfBlankLines = 3

	err =
		txtBlankLinesFive.SetSpecDefaultBlankLines(
			numOfBlankLines,
			StrMech{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtBlankLinesThree."+
			"SetSpecDefaultBlankLines()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecBlankLines_SetSpecRunesBlankLines_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_SetSpecRunesBlankLines_000100()",
		"")

	numOfBlankLines := 3

	newLineChars := "!?\n"

	newLineRunes := []rune(newLineChars)

	expectedFmtStr :=
		strings.Repeat(newLineChars, numOfBlankLines)

	txtBlankLinesOne := TextLineSpecBlankLines{}

	err :=
		txtBlankLinesOne.SetSpecBlankLineRunes(
			numOfBlankLines,
			newLineRunes,
			ePrefix.XCtx("txtBlankLinesOne"))

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

	var txtBlankLinesTwo TextLineSpecBlankLines

	txtBlankLinesTwo,
		err = TextLineSpecBlankLines{}.NewBlankLineRunes(
		numOfBlankLines,
		newLineRunes,
		ePrefix.XCtx(
			"txtBlankLinesTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtBlankLinesTwo.IsValidInstanceError(
		ePrefix.XCtx(
			"txtBlankLinesTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !txtBlankLinesOne.Equal(&txtBlankLinesTwo) {
		t.Errorf("%v - ERROR\n"+
			"Expected txtBlankLinesOne==txtBlankLinesTwo\n"+
			"HOWEVER, THESE INSTANCES ARE NOT EQUAL!\n",
			ePrefix.XCtxEmpty().String())
		return
	}

	var actualStr string

	actualStr,
		err = txtBlankLinesOne.GetFormattedText(
		ePrefix.XCtx(
			"txtBlankLinesOne"))

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
			"txtBlankLinesOne.GetFormattedText()\n"+
			"Error: Expected Text String DOES NOT match\n"+
			"Actual Text String.\n"+
			"Expected Text String = '%v'\n"+
			"Instead, Text String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	txtBlankLinesThree := TextLineSpecBlankLines{}

	numOfBlankLines = 1000001

	newLineChars = "!?\n"

	newLineRunes = []rune(newLineChars)

	err =
		txtBlankLinesThree.SetSpecBlankLineRunes(
			numOfBlankLines,
			newLineRunes,
			ePrefix.XCtx("txtBlankLinesThree"))

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtBlankLinesThree."+
			"SetSpecBlankLineRunes()\n"+
			"because 'numOfBlankLines' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	txtBlankLinesFour := TextLineSpecBlankLines{}

	numOfBlankLines = 3

	newLineChars = ""

	newLineRunes = nil

	err =
		txtBlankLinesFour.SetSpecBlankLineRunes(
			numOfBlankLines,
			newLineRunes,
			ePrefix.XCtx("txtBlankLinesFour"))

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Test # 1 - newLineRunes = nil\n"+
			"Expected an error return from txtBlankLinesFour."+
			"SetSpecBlankLineRunes()\n"+
			"because 'newLineRunes' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	txtBlankLinesFive := TextLineSpecBlankLines{}

	newLineRunes = make([]rune, 0)

	err =
		txtBlankLinesFive.SetSpecBlankLineRunes(
			numOfBlankLines,
			newLineRunes,
			ePrefix.XCtx("txtBlankLinesFour"))

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Test # 2 - newLineRunes = make([]rune, 0)\n"+
			"Expected an error return from txtBlankLinesFive."+
			"SetSpecBlankLines()\n"+
			"because 'newLineRunes' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	txtBlankLinesSix := TextLineSpecBlankLines{}

	numOfBlankLines = 3

	newLineChars = "   \n"

	newLineRunes = []rune(newLineChars)

	err =
		txtBlankLinesSix.SetSpecBlankLineRunes(
			numOfBlankLines,
			newLineRunes,
			StrMech{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtBlankLinesThree."+
			"SetSpecBlankLineRunes()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecBlankLines_String_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_String_000100()",
		"")

	txtBlankLinesBase := TextLineSpecBlankLines{}

	outPutStr := txtBlankLinesBase.String()

	if !strings.Contains(outPutStr, "Error") {

		t.Errorf("%v\n"+
			"Test #1 \n"+
			"Expected an error return from "+
			"txtBlankLinesBase.String()\n"+
			"because 'txtBlankLinesBase' is "+
			"empty and invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtx("Missing Error Return"))

		return
	}

	numOfBlankLines := 3
	newLineChars := "\n"

	var err error

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
		strings.Repeat(newLineChars, 3)

	outPutStr = txtBlankLinesBase.String()

	if strings.Contains(outPutStr, "Error") {

		t.Errorf("%v\n"+
			"Test #2 \n"+
			"Did NOT expect an error return from "+
			"txtBlankLinesBase.String()\n"+
			"because 'txtBlankLinesBase' is "+
			"valid.\n"+
			"HOWEVER, AN ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedFmtStr),
			true)

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(outPutStr),
			true)

	if printableExpectedStr != printableActualStr {
		t.Errorf("%v\n"+
			"Test #3\n"+
			"txtBlankLinesBase.String()\n"+
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

	numOfBlankLines = 4

	newLineChars = "!!\n"

	err = txtBlankLinesBaseAlpha.SetSpecBlankLines(
		3,
		newLineChars,
		ePrefix.XCtx(
			"txtBlankLinesBaseAlpha"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	outPutStr =
		txtBlankLinesBaseAlpha.String()

	expectedFmtStr =
		strings.Repeat(newLineChars, 3)

	printableExpectedStr =
		sMech.ConvertNonPrintableChars(
			[]rune(expectedFmtStr),
			true)

	printableActualStr =
		sMech.ConvertNonPrintableChars(
			[]rune(outPutStr),
			true)

	if printableExpectedStr != printableActualStr {
		t.Errorf("%v\n"+
			"Test #4\n"+
			"txtBlankLinesBaseAlpha.String()\n"+
			"Error: Expected Text String DOES NOT match\n"+
			"Actual Text String.\n"+
			"Expected Text String = '%v'\n"+
			"Instead, Text String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	return
}

func TestTextLineSpecBlankLines_TextLineBuilder_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_TextLineBuilder_000100()",
		"")

	numOfBlankLines := 3

	newLineChars := "  \n"

	newLineRunes := []rune(newLineChars)

	expectedTextStr :=
		strings.Repeat(newLineChars, numOfBlankLines)

	txtBlankLinesOne,
		err := TextLineSpecBlankLines{}.NewBlankLineRunes(
		numOfBlankLines,
		newLineRunes,
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

	sb := strings.Builder{}

	err = txtBlankLinesOne.TextBuilder(
		&sb,
		ePrefix.XCtx("txtBlankLinesOne->sb"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedTextStr),
			true)

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(sb.String()),
			true)

	if printableExpectedStr != printableActualStr {
		t.Errorf("%v\n"+
			"Test #1 - txtBlankLinesOne.TextBuilder()\n"+
			"Error: Expected Text String DOES NOT match\n"+
			"Actual Text String.\n"+
			"Expected Text String = '%v'\n"+
			"Instead, Text String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	return
}

func TestTextLineSpecBlankLines_TextLineBuilder_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_TextLineBuilder_000200()",
		"")

	txtBlankLinesOne := TextLineSpecBlankLines{}

	sb := strings.Builder{}

	err := txtBlankLinesOne.TextBuilder(
		&sb,
		ePrefix.XCtx("empty txtBlankLinesOne->sb"))

	if err == nil {
		t.Errorf("%v\n"+
			"Error:\n"+
			"Expected error return from txtBlankLinesOne.TextBuilder()\n"+
			"because 'txtBlankLinesOne' is empty.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}
	numOfBlankLines := 3

	newLineChars := "!!\n"

	newLineRunes := []rune(newLineChars)

	expectedTextStr :=
		strings.Repeat(newLineChars, numOfBlankLines)

	var txtBlankLinesTwo TextLineSpecBlankLines

	txtBlankLinesTwo,
		err = TextLineSpecBlankLines{}.NewBlankLineRunes(
		numOfBlankLines,
		newLineRunes,
		ePrefix.XCtx(
			"txtBlankLinesTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtBlankLinesTwo.IsValidInstanceError(
		ePrefix.XCtx(
			"txtBlankLinesTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtBlankLinesTwo.TextBuilder(
		nil,
		ePrefix.XCtx("txtBlankLinesTwo->nil sb"))

	if err == nil {
		t.Errorf("%v\n"+
			"Error:\n"+
			"Expected error return from txtBlankLinesTwo.TextBuilder()\n"+
			"because strings.Builder pointer is 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	txtBlankLinesThree := TextLineSpecBlankLines{}

	err = txtBlankLinesThree.CopyIn(
		&txtBlankLinesTwo,
		ePrefix.XCtx(
			"txtBlankLinesThree<-txtBlankLinesTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}
	err = txtBlankLinesThree.IsValidInstanceError(
		ePrefix.XCtx(
			"txtBlankLinesTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtBlankLinesThree.TextBuilder(
		&sb,
		ePrefix.XCtx("valid txtBlankLinesThree->sb"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedTextStr),
			true)

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(sb.String()),
			true)

	if printableExpectedStr != printableActualStr {
		t.Errorf("%v\n"+
			"txtBlankLinesThree.TextBuilder()\n"+
			"Error: Expected Text String DOES NOT match\n"+
			"Actual Text String.\n"+
			"Expected Text String = '%v'\n"+
			"Instead, Text String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	err = txtBlankLinesThree.TextBuilder(
		&sb,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtBlankLinesThree."+
			"TextBuilder()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecBlankLines_TextLineSpecName_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_TextLineSpecName_000100()",
		"")

	txtBlankLinesOne := TextLineSpecBlankLines{}

	textLineSpecName :=
		txtBlankLinesOne.TextLineSpecName()

	expectedTextLineSpecName := "TextLineSpecBlankLines"

	if textLineSpecName != expectedTextLineSpecName {

		t.Errorf("%v\n"+
			"txtBlankLinesOne.TextLineSpecName()\n"+
			"Error: Expected Text String DOES NOT match\n"+
			"Actual Text String.\n"+
			"Expected Text Line Spec Name = '%v'\n"+
			"Instead, Text Line Spec Name = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedTextLineSpecName,
			textLineSpecName)

		return
	}

	return
}

func TestTextLineSpecBlankLines_TextTypeName_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecBlankLines_TextTypeName_000100()",
		"")

	txtBlankLinesOne := TextLineSpecBlankLines{}

	textTypeName :=
		txtBlankLinesOne.TextTypeName()

	expectedTextTypeName := "TextLineSpecBlankLines"

	if textTypeName != expectedTextTypeName {

		t.Errorf("%v\n"+
			"txtBlankLinesOne.TextLineSpecName()\n"+
			"Error: Expected Text String DOES NOT match\n"+
			"Actual Text String.\n"+
			"Expected Text Type Name = '%v'\n"+
			"Instead, Text Type Name = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedTextTypeName,
			textTypeName)

		return
	}

	return
}
