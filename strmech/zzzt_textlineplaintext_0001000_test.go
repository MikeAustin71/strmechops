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
			string(expectedRightMarginChars),
			string(plainTextLine01.rightMarginChars))

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
			string(expectedNewLineChars),
			string(plainTextLine01.newLineChars))

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

	}

	return
}

func TestTextLineSpecPlainText_copyOut_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_CopyOut_000100()",
		"")

	expectedLeftMarginChars := []rune{' ', ' ', ' '}
	expectedRightMarginChars := []rune{' ', ' ', ' '}
	expectedNewLineChars := []rune{'\n', '\n'}

	expectedTextStringRunes := []rune("How now brown cow!")
	expectedTextString := string(expectedTextStringRunes)

	txtLinePlainTextNanobot := textLineSpecPlainTextNanobot{}

	err :=
		txtLinePlainTextNanobot.setPlainTextSpecRunes(
			nil,
			expectedLeftMarginChars,
			expectedRightMarginChars,
			expectedTextStringRunes,
			expectedNewLineChars,
			false,
			&ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Error: txtLinePlainTextNanobot.setPlainTextSpecRunes()\n"+
			"Expected an error return because input parameter\n"+
			"'plainTxtLine' is nil.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	plainTextLine01 := TextLineSpecPlainText{}

	err =
		txtLinePlainTextNanobot.setPlainTextSpecRunes(
			&plainTextLine01,
			expectedLeftMarginChars,
			expectedRightMarginChars,
			nil,
			expectedNewLineChars,
			false,
			&ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Error: txtLinePlainTextNanobot.setPlainTextSpecRunes()\n"+
			"Expected an error return because input parameter\n"+
			"'textRune' is 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	badRunes := make([]rune, 1000001)

	for i := 0; i < 1000001; i++ {
		badRunes[i] = 'x'
	}

	err =
		txtLinePlainTextNanobot.setPlainTextSpecRunes(
			&plainTextLine01,
			expectedLeftMarginChars,
			expectedRightMarginChars,
			badRunes,
			expectedNewLineChars,
			false,
			&ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Error: txtLinePlainTextNanobot.setPlainTextSpecRunes()\n"+
			"Expected an error return because input parameter\n"+
			"'textRunes' is has a length of 1,000,001.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	badRunes = make([]rune, 0)

	err =
		txtLinePlainTextNanobot.setPlainTextSpecRunes(
			&plainTextLine01,
			expectedLeftMarginChars,
			expectedRightMarginChars,
			expectedTextStringRunes,
			expectedNewLineChars,
			false,
			ePrefix.XCtx(
				"plainTextLine01"))

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

	txtLinePlainTextNanobot2 := textLineSpecPlainTextNanobot{}

	plainTextLine02 := TextLineSpecPlainText{}

	plainTextLine02,
		err = txtLinePlainTextNanobot2.copyOut(
		nil,
		&ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Error: txtLinePlainTextNanobot2.copyOut()\n"+
			"Expected an error return because input parameter\n"+
			"'plainTxtLine' is 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	plainTextLine02,
		err = txtLinePlainTextNanobot2.copyOut(
		&plainTextLine01,
		ePrefix.XCtx(
			"plainTextLine01->plainTextLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = plainTextLine02.IsValidInstanceError(
		ePrefix.XCtx("plainTextLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !plainTextLine02.Equal(&plainTextLine01) {
		t.Errorf("%v\n"+
			"Error: Expected plainTextLine02 == plainTextLine01\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	actualTextString := plainTextLine02.GetTextString()

	if expectedTextString != actualTextString {
		t.Errorf("%v\n"+
			"Error:\n"+
			"Expected Text String = '%v'\n"+
			"Instead, Text String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedTextString,
			actualTextString)

		return
	}

	plainTextLine02.textString = ""

	_,
		err = txtLinePlainTextNanobot2.copyOut(
		&plainTextLine02,
		&ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Error: txtLinePlainTextNanobot2.copyOut()\n"+
			"Expected an error return because input parameter\n"+
			"'plainTxtLine' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecPlainText_CopyOut_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_CopyOut_000100()",
		"")

	expectedLeftMarginChars := []rune{' ', ' ', ' '}
	expectedRightMarginChars := []rune{' ', ' ', ' '}
	expectedNewLineChars := []rune{'\n', '\n'}

	expectedTextString := "How now brown cow!"

	plainTextLine01,
		err := TextLineSpecPlainText{}.NewPlainText(
		expectedLeftMarginChars,
		expectedRightMarginChars,
		expectedTextString,
		expectedNewLineChars,
		false,
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

	plainTextLine02 := TextLineSpecPlainText{}

	plainTextLine02,
		err = plainTextLine01.CopyOut(
		ePrefix.XCtx("plainTextLine01->" +
			"plainTextLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = plainTextLine02.IsValidInstanceError(
		ePrefix.XCtx("plainTextLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !plainTextLine02.Equal(&plainTextLine01) {
		t.Errorf("%v\n"+
			"Error: Expected plainTextLine02 == plainTextLine01\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	actualTextStr :=
		plainTextLine02.GetTextString()

	if expectedTextString != actualTextStr {
		t.Errorf("%v\n"+
			"Error:\n"+
			"Expected Text String = '%v'\n"+
			"Instead, Text String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedTextString,
			actualTextStr)

		return
	}

	sMechPreon := strMechPreon{}

	if !sMechPreon.equalRuneArrays(
		plainTextLine02.leftMarginChars,
		expectedLeftMarginChars) {

		t.Errorf("%v\n"+
			"Error: plainTextLine02\n"+
			"Expected Left Margin Chars = '%v'\n"+
			"Instead, Left Margin Chars = '%v'\n",
			ePrefix.String(),
			string(expectedLeftMarginChars),
			string(plainTextLine02.leftMarginChars))

		return
	}

	actualLeftMarginChars :=
		plainTextLine02.GetLeftMarginRunes()

	if !sMechPreon.equalRuneArrays(
		actualLeftMarginChars,
		expectedLeftMarginChars) {

		t.Errorf("%v\n"+
			"Error: plainTextLine02.GetLeftMarginRunes()\n"+
			"Expected Left Margin Chars = '%v'\n"+
			"Instead, Left Margin Chars = '%v'\n",
			ePrefix.String(),
			string(expectedLeftMarginChars),
			string(actualLeftMarginChars))

		return
	}

	if !sMechPreon.equalRuneArrays(
		plainTextLine02.rightMarginChars,
		expectedRightMarginChars) {

		t.Errorf("%v\n"+
			"Error: plainTextLine02\n"+
			"Expected Right Margin Chars = '%v'\n"+
			"Instead, Right Margin Chars = '%v'\n",
			ePrefix.String(),
			string(expectedRightMarginChars),
			string(plainTextLine02.rightMarginChars))

		return
	}

	actualRightMarginChars :=
		plainTextLine02.GetRightMarginRunes()

	if !sMechPreon.equalRuneArrays(
		actualRightMarginChars,
		expectedRightMarginChars) {

		t.Errorf("%v\n"+
			"Error: plainTextLine02.GetRightMarginRunes()\n"+
			"Expected Right Margin Chars = '%v'\n"+
			"Instead, Right Margin Chars = '%v'\n",
			ePrefix.String(),
			string(actualRightMarginChars),
			string(actualRightMarginChars))

		return
	}

	if !sMechPreon.equalRuneArrays(
		plainTextLine02.newLineChars,
		expectedNewLineChars) {

		t.Errorf("%v\n"+
			"Error: plainTextLine02\n"+
			"Expected New Line Chars = '%v'\n"+
			"Instead, New Line Chars = '%v'\n",
			ePrefix.String(),
			string(expectedNewLineChars),
			string(plainTextLine02.newLineChars))

		return
	}

	actualNewLineChars :=
		plainTextLine02.GetLineTerminationRunes()

	if !sMechPreon.equalRuneArrays(
		actualNewLineChars,
		expectedNewLineChars) {

		t.Errorf("%v\n"+
			"Error: plainTextLine02.GetLineTerminationRunes()\n"+
			"Expected New Line Chars = '%v'\n"+
			"Instead, New Line Chars = '%v'\n",
			ePrefix.String(),
			string(expectedNewLineChars),
			string(actualNewLineChars))

		return
	}

	if plainTextLine02.turnLineTerminatorOff != false {

		t.Errorf("%v\n"+
			"Error: plainTextLine02\n"+
			"Expected turnLineTerminatorOff = 'false'\n"+
			"Instead, turnLineTerminatorOff = 'true'\n",
			ePrefix.String())

		return
	}

	if plainTextLine02.GetTurnLineTerminatorOff() != false {

		t.Errorf("%v\n"+
			"Error: plainTextLine02.GetTurnLineTerminatorOff()\n"+
			"Expected turnLineTerminatorOff = 'false'\n"+
			"Instead, turnLineTerminatorOff = 'true'\n",
			ePrefix.String())

		return
	}

	return
}
