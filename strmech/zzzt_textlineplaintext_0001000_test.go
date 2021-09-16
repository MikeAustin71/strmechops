package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"testing"
)

func TestTextLineSpecPlainText_copyIn_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_copyIn_000100()",
		"")

	leftMargin := 3
	rightMargin := 3
	textString := "How now brown cow!"

	plainTextLine01,
		err := TextLineSpecPlainText{}.NewPtrDefault(
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

	txtLinePlainTxtNanobot := textLineSpecPlainTextNanobot{}

	err = txtLinePlainTxtNanobot.copyIn(
		nil,
		plainTextLine01,
		ePrefix.XCtx(
			"targetPlainTextLine=nil"))

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from "+
			"textLineSpecPlainTextNanobot{}.copyIn()\n"+
			"because input parameter 'targetPlainTextLine' is 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtx("Missing Error Return"))
		return
	}

	txtLinePlainTxtNanobot2 := textLineSpecPlainTextNanobot{}

	err = txtLinePlainTxtNanobot2.copyIn(
		plainTextLine01,
		nil,
		ePrefix.XCtx(
			"incomingPlainTextLine=nil"))

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from "+
			"txtLinePlainTxtNanobot2.copyIn()\n"+
			"because input parameter 'incomingPlainTextLine' is 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtx("Missing Error Return"))
		return
	}

	badTextLinePlainTxt99 := TextLineSpecPlainText{}

	badTextLinePlainTxt99.leftMarginChars =
		[]rune{' ', ' ', ' '}

	badTextLinePlainTxt99.rightMarginChars =
		[]rune{' ', ' ', ' '}

	badTextLinePlainTxt99.textString = ""

	badTextLinePlainTxt99.turnLineTerminatorOff = false

	badTextLinePlainTxt99.newLineChars = []rune{'\n'}

	txtLinePlainTxtNanobot3 := textLineSpecPlainTextNanobot{}

	plainTextLine02 := TextLineSpecPlainText{}

	err = txtLinePlainTxtNanobot3.copyIn(
		&plainTextLine02,
		&badTextLinePlainTxt99,
		ePrefix.XCtx(
			"badTextLinePlainTxt99 invalid"))

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from "+
			"txtLinePlainTxtNanobot3.copyIn()\n"+
			"because input parameter 'incomingPlainTextLine' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtx("Missing Error Return"))
		return
	}

	txtLinePlainTxtNanobot4 := textLineSpecPlainTextNanobot{}

	var plainTextLine03 TextLineSpecPlainText

	plainTextLine03,
		err = TextLineSpecPlainText{}.NewDefault(
		leftMargin,
		rightMargin,
		textString,
		ePrefix.XCtx("plainTextLine03"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	plainTextLine04 := TextLineSpecPlainText{}

	err = txtLinePlainTxtNanobot4.copyIn(
		&plainTextLine04,
		&plainTextLine03,
		ePrefix.XCtx(
			"plainTextLine03->plainTextLine04"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = plainTextLine04.IsValidInstanceError(
		ePrefix.XCtx("plainTextLine04"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !plainTextLine04.Equal(&plainTextLine03) {
		t.Errorf("%v\n"+
			"Error: Expected plainTextLine04 == plainTextLine03\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.XCtxEmpty().String())
	}

	return
}

func TestTextLineSpecPlainText_CopyIn_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_CopyIn_000100()",
		"")

	leftMargin := 3
	rightMargin := 3
	textString := "How now brown cow!"

	expectedLeftMarginChars := []rune{' ', ' ', ' '}
	expectedRightMarginChars := []rune{' ', ' ', ' '}
	expectedNewLineChars := []rune{'\n'}

	plainTextLine01 := TextLineSpecPlainText{}

	err := plainTextLine01.SetPlainTextDefault(
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

	sMechPreon := strMechPreon{}

	if !sMechPreon.equalRuneArrays(
		plainTextLine01.leftMarginChars,
		expectedLeftMarginChars) {

		t.Errorf("%v\n"+
			"Error: plainTextLine01\n"+
			"Expected Left Margin Chars = '%v'\n"+
			"Instead, Left Margin Chars = '%v'\n",
			ePrefix.String(),
			string(expectedLeftMarginChars),
			string(plainTextLine01.leftMarginChars))

		return
	}

	if !sMechPreon.equalRuneArrays(
		plainTextLine01.rightMarginChars,
		expectedRightMarginChars) {

		t.Errorf("%v\n"+
			"Error: plainTextLine01\n"+
			"Expected Right Margin Chars = '%v'\n"+
			"Instead, Right Margin Chars = '%v'\n",
			ePrefix.String(),
			string(expectedLeftMarginChars),
			string(plainTextLine01.leftMarginChars))

		return
	}

	if !sMechPreon.equalRuneArrays(
		plainTextLine01.newLineChars,
		expectedNewLineChars) {

		t.Errorf("%v\n"+
			"Error: plainTextLine01\n"+
			"Expected New Line Chars = '%v'\n"+
			"Instead, New Line Chars = '%v'\n",
			ePrefix.String(),
			string(expectedLeftMarginChars),
			string(plainTextLine01.leftMarginChars))

		return
	}

	if plainTextLine01.turnLineTerminatorOff != false {

		t.Errorf("%v\n"+
			"Error: plainTextLine01\n"+
			"Expected turnLineTerminatorOff = 'false'\n"+
			"Instead, turnLineTerminatorOff = 'true'\n",
			ePrefix.String())

		return
	}

	if plainTextLine01.textString !=
		textString {

		t.Errorf("%v\n"+
			"Error: plainTextLine01\n"+
			"Expected textString = '%v'\n"+
			"Instead, textSTring = '%v'\n",
			ePrefix.String(),
			textString,
			plainTextLine01.textString)

		return
	}

	plainTextLine02 := TextLineSpecPlainText{}

	err = plainTextLine02.CopyIn(
		nil,
		ePrefix.XCtx("incomingPlainTxtLine='nil'"))

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from "+
			"textLineSpecPlainTextNanobot{}.copyIn()\n"+
			"because input parameter 'incomingPlainTxtLine' is 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtx("incomingPlainTxtLine - Missing Error Return"))
		return
	}

	plainTextLine03 := TextLineSpecPlainText{}

	err = plainTextLine03.CopyIn(
		&plainTextLine01,
		ePrefix.XCtx("incomingPlainTxtLine='plainTextLine01'"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = plainTextLine03.IsValidInstanceError(
		ePrefix.XCtx("plainTextLine03"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !plainTextLine03.Equal(&plainTextLine01) {
		t.Errorf("%v\n"+
			"Error: Expected plainTextLine03 == plainTextLine01\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	plainTextLine04 := TextLineSpecPlainText{}

	badTextLinePlainTxt99 := TextLineSpecPlainText{}

	badTextLinePlainTxt99.leftMarginChars =
		[]rune{' ', ' ', ' '}

	badTextLinePlainTxt99.rightMarginChars =
		[]rune{' ', ' ', ' '}

	badTextLinePlainTxt99.textString = ""

	badTextLinePlainTxt99.turnLineTerminatorOff = false

	badTextLinePlainTxt99.newLineChars = []rune{'\n'}

	err = plainTextLine04.CopyIn(
		&badTextLinePlainTxt99,
		ePrefix.XCtx("incomingPlainTxtLine='badTextLinePlainTxt99'"))

	if err == nil {

		t.Errorf("%v\n"+
			"Expected an error return from "+
			"plainTextLine04.CopyIn()\n"+
			"because input parameter 'incomingPlainTextLine' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtx("Missing Error Return"))

		return
	}

	return
}
