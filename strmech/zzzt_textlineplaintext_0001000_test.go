package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"strings"
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

		return
	}

	return
}

func TestTextLineSpecPlainText_copyOut_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_copyOut_000100()",
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

func TestTextLineSpecPlainText_CopyOutITextLine_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_CopyOutITextLine_000100()",
		"")

	expectedLeftMarginChars := "   "
	expectedRightMarginChars := "   "
	expectedTextString := "The cow jumped over the moon!"
	expectedNewLineChars := "\n"

	plainTextLine01,
		err := TextLineSpecPlainText{}.NewPlainTextStrings(
		expectedLeftMarginChars,
		expectedRightMarginChars,
		expectedTextString,
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

	actualTurnLineTerminatorOff :=
		plainTextLine01.GetTurnLineTerminatorOff()

	if false != actualTurnLineTerminatorOff {
		t.Errorf("%v\n"+
			"Error: plainTextLine01.GetTurnLineTerminatorOff()"+
			"Expected TurnLineTerminatorOff == 'false'\n"+
			"Instead, TurnLineTerminatorOff == 'true'\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	var iTextLine ITextLineSpecification

	iTextLine, err =
		plainTextLine01.CopyOutITextLine(
			ePrefix.XCtx(
				"plainTextLine01->iTextLine"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	plainTextLine02, ok := iTextLine.(*TextLineSpecPlainText)

	if !ok {
		t.Errorf("%v\n"+
			"Error: iTextLine.(*TextLineSpecPlainText)\n"+
			"Could not convert 'iTextLine' to TextLineSpecPlainText\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	if !plainTextLine02.Equal(&plainTextLine01) {
		t.Errorf("%v\n"+
			"Error: Expected plainTextLine02 == plainTextLine01\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	if !plainTextLine01.EqualITextLine(iTextLine) {
		t.Errorf("%v\n"+
			"Error: plainTextLine01.EqualITextLine(iTextLine)\n"+
			"Expected plainTextLine01 == iTextLine\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	plainTextLine03 := TextLineSpecPlainText{}

	_,
		err =
		plainTextLine03.CopyOutITextLine(
			ePrefix.XCtx(
				"plainTextLine03->iTextLine"))

	if err == nil {
		t.Errorf("%v\n"+
			"Error: plainTextLine03.CopyOutITextLine()\n"+
			"Expected an error return because 'plainTextLine03' is empty!\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecPlainText_CopyOutPtr_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_CopyOutPtr_000100()",
		"")

	leftMargin := 3
	rightMargin := 3
	textString := "How now brown cow!"

	plainTextLine01,
		err := TextLineSpecPlainText{}.NewDefault(
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

	var plainTextLine02 *TextLineSpecPlainText

	plainTextLine02,
		err = plainTextLine01.CopyOutPtr(
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

	var plainTextLine03 *TextLineSpecPlainText

	plainTextLine02.textString = ""

	plainTextLine03,
		err = plainTextLine02.CopyOutPtr(
		ePrefix.XCtx(
			"plainTextLine02->plainTextLine03"))

	if err == nil {

		t.Errorf("%v\n"+
			"Expected an error return from "+
			"plainTextLine02.CopyOutPtr()\n"+
			"because 'plainTextLine02' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtx("Missing Error Return"))

		return
	}

	err = plainTextLine03.IsValidInstanceError(
		ePrefix.XCtx("plainTextLine03"))

	if err == nil {

		t.Errorf("%v\n"+
			"Expected an error return from "+
			"plainTextLine03.IsValidInstanceError()\n"+
			"because 'plainTextLine03' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtx("Missing Error Return"))

		return
	}

	return
}

func TestTextLineSpecPlainText_empty_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_empty_000100()",
		"")

	txtLinePlainTextElectron01 := textLineSpecPlainTextElectron{}

	txtLinePlainTextElectron01.empty(
		nil)

	expectedLeftMarginChars := "   "
	expectedRightMarginChars := "   "
	expectedTextString := "The cow jumped over the moon!"
	expectedNewLineChars := "\n"

	plainTextLine01,
		err := TextLineSpecPlainText{}.NewPtrPlainTextStrings(
		expectedLeftMarginChars,
		expectedRightMarginChars,
		expectedTextString,
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

	plainTextLine01.TurnAutoLineTerminationOff()

	txtLinePlainTextElectron02 := textLineSpecPlainTextElectron{}

	txtLinePlainTextElectron02.empty(
		plainTextLine01)

	if len(plainTextLine01.textString) != 0 {
		t.Errorf("%v\n"+
			"Error: plainTextLine01.textString\n"+
			"Expected length of textString == 0\n"+
			"Instead, length of textString == %v\n",
			ePrefix.XCtxEmpty().String(),
			len(plainTextLine01.textString))

		return
	}

	if plainTextLine01.leftMarginChars != nil {
		t.Errorf("%v\n"+
			"Error: plainTextLine01.leftMarginChars\n"+
			"Expected leftMarginChars == 'nil'\n"+
			"Instead, leftMarginChars == '%v'\n",
			ePrefix.XCtxEmpty().String(),
			string(plainTextLine01.leftMarginChars))

		return
	}

	if plainTextLine01.rightMarginChars != nil {
		t.Errorf("%v\n"+
			"Error: plainTextLine01.rightMarginChars\n"+
			"Expected rightMarginChars == 'nil'\n"+
			"Instead, rightMarginChars == '%v'\n",
			ePrefix.XCtxEmpty().String(),
			string(plainTextLine01.rightMarginChars))

		return
	}

	if plainTextLine01.newLineChars != nil {
		t.Errorf("%v\n"+
			"Error: plainTextLine01.newLineChars\n"+
			"Expected newLineChars == 'nil'\n"+
			"Instead, newLineChars == '%v'\n",
			ePrefix.XCtxEmpty().String(),
			string(plainTextLine01.newLineChars))

		return
	}

	if plainTextLine01.turnLineTerminatorOff == true {
		t.Errorf("%v\n"+
			"Error: plainTextLine01.turnLineTerminatorOff\n"+
			"Expected turnLineTerminatorOff == 'false'\n"+
			"Instead, turnLineTerminatorOff == 'true'\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecPlainText_Empty_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_Empty_000100()",
		"")

	expectedLeftMarginChars := "   "
	expectedRightMarginChars := "   "
	expectedTextString := "The cow jumped over the moon!"
	expectedNewLineChars := "\n"

	plainTextLine01,
		err := TextLineSpecPlainText{}.NewPtrPlainTextStrings(
		expectedLeftMarginChars,
		expectedRightMarginChars,
		expectedTextString,
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

	plainTextLine01.TurnAutoLineTerminationOff()

	plainTextLine01.Empty()

	if len(plainTextLine01.textString) != 0 {
		t.Errorf("%v\n"+
			"Error: plainTextLine01.textString\n"+
			"Expected length of textString == 0\n"+
			"Instead, length of textString == %v\n",
			ePrefix.XCtxEmpty().String(),
			len(plainTextLine01.textString))

		return
	}

	if plainTextLine01.leftMarginChars != nil {
		t.Errorf("%v\n"+
			"Error: plainTextLine01.leftMarginChars\n"+
			"Expected leftMarginChars == 'nil'\n"+
			"Instead, leftMarginChars == '%v'\n",
			ePrefix.XCtxEmpty().String(),
			string(plainTextLine01.leftMarginChars))

		return
	}

	if plainTextLine01.rightMarginChars != nil {
		t.Errorf("%v\n"+
			"Error: plainTextLine01.rightMarginChars\n"+
			"Expected rightMarginChars == 'nil'\n"+
			"Instead, rightMarginChars == '%v'\n",
			ePrefix.XCtxEmpty().String(),
			string(plainTextLine01.rightMarginChars))

		return
	}

	if plainTextLine01.newLineChars != nil {
		t.Errorf("%v\n"+
			"Error: plainTextLine01.newLineChars\n"+
			"Expected newLineChars == 'nil'\n"+
			"Instead, newLineChars == '%v'\n",
			ePrefix.XCtxEmpty().String(),
			string(plainTextLine01.newLineChars))

		return
	}

	if plainTextLine01.turnLineTerminatorOff == true {
		t.Errorf("%v\n"+
			"Error: plainTextLine01.turnLineTerminatorOff\n"+
			"Expected turnLineTerminatorOff == 'false'\n"+
			"Instead, turnLineTerminatorOff == 'true'\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecPlainText_Empty_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_Empty_000200()",
		"")

	plainTextLine00 := TextLineSpecPlainText{}

	plainTextLine00.Empty()

	expectedLeftMarginChars := "   "
	expectedRightMarginChars := "   "
	expectedTextString := "The cow jumped over the moon!"
	expectedNewLineChars := "\n"

	var plainTextLine01 *TextLineSpecPlainText
	var err error

	plainTextLine01,
		err = TextLineSpecPlainText{}.NewPtrPlainTextStrings(
		expectedLeftMarginChars,
		expectedRightMarginChars,
		expectedTextString,
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

	plainTextLine01.TurnAutoLineTerminationOff()

	plainTextLine01.Empty()

	if len(plainTextLine01.textString) != 0 {
		t.Errorf("%v\n"+
			"Error: plainTextLine01.textString\n"+
			"Expected length of textString == 0\n"+
			"Instead, length of textString == %v\n",
			ePrefix.XCtxEmpty().String(),
			len(plainTextLine01.textString))

		return
	}

	if plainTextLine01.leftMarginChars != nil {
		t.Errorf("%v\n"+
			"Error: plainTextLine01.leftMarginChars\n"+
			"Expected leftMarginChars == 'nil'\n"+
			"Instead, leftMarginChars == '%v'\n",
			ePrefix.XCtxEmpty().String(),
			string(plainTextLine01.leftMarginChars))

		return
	}

	if plainTextLine01.rightMarginChars != nil {
		t.Errorf("%v\n"+
			"Error: plainTextLine01.rightMarginChars\n"+
			"Expected rightMarginChars == 'nil'\n"+
			"Instead, rightMarginChars == '%v'\n",
			ePrefix.XCtxEmpty().String(),
			string(plainTextLine01.rightMarginChars))

		return
	}

	if plainTextLine01.newLineChars != nil {
		t.Errorf("%v\n"+
			"Error: plainTextLine01.newLineChars\n"+
			"Expected newLineChars == 'nil'\n"+
			"Instead, newLineChars == '%v'\n",
			ePrefix.XCtxEmpty().String(),
			string(plainTextLine01.newLineChars))

		return
	}

	if plainTextLine01.turnLineTerminatorOff == true {
		t.Errorf("%v\n"+
			"Error: plainTextLine01.turnLineTerminatorOff\n"+
			"Expected turnLineTerminatorOff == 'false'\n"+
			"Instead, turnLineTerminatorOff == 'true'\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecPlainText_equal_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_equal_000100()",
		"")

	expectedLeftMarginChars := []rune{' ', ' ', ' '}
	expectedRightMarginChars := []rune{' ', ' ', ' '}
	expectedNewLineChars := []rune{'\n', '\n'}

	expectedTextString := "How now brown cow!"
	expectedTextChars := []rune(expectedTextString)

	plainTextLine01,
		err := TextLineSpecPlainText{}.NewPlainTextRunes(
		expectedLeftMarginChars,
		expectedRightMarginChars,
		expectedTextChars,
		expectedNewLineChars,
		true,
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

	txtLinePlainTextElectron := textLineSpecPlainTextElectron{}

	areEqual :=
		txtLinePlainTextElectron.equal(
			nil,
			&plainTextLine01)

	if areEqual == true {
		t.Errorf("%v\n"+
			"Error: txtLinePlainTextElectron.equal()\n"+
			"Expected areEqual == false because input\n"+
			"parameter 'plainTxtLineOne is 'nil'.\n"+
			"HOWEVER, areEqual == true\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	areEqual =
		txtLinePlainTextElectron.equal(
			&plainTextLine01,
			nil)

	if areEqual == true {
		t.Errorf("%v\n"+
			"Error: txtLinePlainTextElectron.equal()\n"+
			"Expected areEqual == false because input\n"+
			"parameter 'plainTxtLineTwo is 'nil'.\n"+
			"HOWEVER, areEqual == true\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	var plainTextLine02 *TextLineSpecPlainText

	plainTextLine02,
		err = TextLineSpecPlainText{}.NewPtrPlainTextRunes(
		expectedLeftMarginChars,
		expectedRightMarginChars,
		expectedTextChars,
		expectedNewLineChars,
		true,
		ePrefix.XCtx(
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

	txtLinePlainTextElectron2 := textLineSpecPlainTextElectron{}

	areEqual =
		txtLinePlainTextElectron2.equal(
			&plainTextLine01,
			plainTextLine02)

	if areEqual == false {
		t.Errorf("%v\n"+
			"Error: txtLinePlainTextElectron2.equal()\n"+
			"Expected areEqual == true because input\n"+
			"parameters 'plainTextLine01' and 'plainTextLine02'"+
			"are euqal.\n"+
			"HOWEVER, areEqual == false\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	plainTextLine03 := TextLineSpecPlainText{}

	_,
		err = plainTextLine03.CopyOut(
		ePrefix.XCtxEmpty())

	if err == nil {
		t.Errorf("%v\n"+
			"Error: plainTextLine03.CopyOut()\n"+
			"Expected an error return because \n"+
			"'plainTextLine03' is empty an invalid\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	plainTextLine04 := TextLineSpecPlainText{}

	_,
		err = plainTextLine04.CopyOutPtr(
		ePrefix.XCtxEmpty())

	if err == nil {
		t.Errorf("%v\n"+
			"Error: plainTextLine04.CopyOutPtr()\n"+
			"Expected an error return because \n"+
			"'plainTextLine04' is empty an invalid\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	plainTextLine03,
		err = plainTextLine01.CopyOut(
		ePrefix.XCtx(
			"plainTextLine01->plainTextLine03"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	areEqual =
		txtLinePlainTextElectron2.equal(
			&plainTextLine01,
			&plainTextLine03)

	if areEqual == false {
		t.Errorf("%v\n"+
			"Error: txtLinePlainTextElectron2.equal()\n"+
			"Expected areEqual == true because input\n"+
			"parameters 'plainTextLine01' and 'plainTextLine03'"+
			"are equal.\n"+
			"HOWEVER, areEqual == false\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	plainTextLine03.leftMarginChars = []rune{'x'}

	areEqual =
		txtLinePlainTextElectron2.equal(
			&plainTextLine01,
			&plainTextLine03)

	if areEqual == true {
		t.Errorf("%v\n"+
			"Error: txtLinePlainTextElectron2.equal()\n"+
			"Expected areEqual == false because input\n"+
			"'plainTextLine01.leftMarginChars' and\n"+
			"'plainTextLine03.leftMarginChars'"+
			"are ARE NOT equal.\n"+
			"HOWEVER, areEqual == true\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	plainTextLine05 := TextLineSpecPlainText{}

	err =
		plainTextLine05.CopyIn(
			&plainTextLine01,
			ePrefix.XCtx(
				"plainTextLine01->plainTextLine05"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	plainTextLine05.rightMarginChars = []rune{'B', 'B'}

	areEqual =
		txtLinePlainTextElectron2.equal(
			&plainTextLine01,
			&plainTextLine05)

	if areEqual == true {
		t.Errorf("%v\n"+
			"Error: txtLinePlainTextElectron2.equal()\n"+
			"Expected areEqual == false because input\n"+
			"'plainTextLine01.rightMarginChars' and\n"+
			"'plainTextLine05.rightMarginChars'"+
			"are ARE NOT equal.\n"+
			"HOWEVER, areEqual == true\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	err =
		plainTextLine04.CopyIn(
			&plainTextLine01,
			ePrefix.XCtx(
				"plainTextLine01->plainTextLine04"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	plainTextLine04.textString = "X X X X X X X!!!!&&&&"

	areEqual =
		txtLinePlainTextElectron2.equal(
			&plainTextLine01,
			&plainTextLine04)

	if areEqual == true {
		t.Errorf("%v\n"+
			"Error: txtLinePlainTextElectron2.equal()\n"+
			"Expected areEqual == false because input\n"+
			"'plainTextLine01.textString' and\n"+
			"'plainTextLine04.textString'"+
			"are ARE NOT equal.\n"+
			"HOWEVER, areEqual == true\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	err =
		plainTextLine03.CopyIn(
			&plainTextLine01,
			ePrefix.XCtx(
				"plainTextLine01->plainTextLine03 #2"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	plainTextLine03.TurnAutoLineTerminationOn()

	areEqual =
		txtLinePlainTextElectron2.equal(
			&plainTextLine01,
			&plainTextLine03)

	if areEqual == true {
		t.Errorf("%v\n"+
			"Error: txtLinePlainTextElectron2.equal()\n"+
			"Expected areEqual == false because input\n"+
			"'plainTextLine01.turnLineTerminatorOff' and\n"+
			"'plainTextLine03.turnLineTerminatorOff'"+
			"are ARE NOT equal.\n"+
			"HOWEVER, areEqual == true\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	err =
		plainTextLine02.CopyIn(
			&plainTextLine01,
			ePrefix.XCtx(
				"plainTextLine01->plainTextLine02 #2"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	plainTextLine02.newLineChars = []rune{'y', 'v', '5'}

	areEqual =
		txtLinePlainTextElectron2.equal(
			&plainTextLine01,
			plainTextLine02)

	if areEqual == true {
		t.Errorf("%v\n"+
			"Error: txtLinePlainTextElectron2.equal()\n"+
			"Expected areEqual == false because input\n"+
			"'plainTextLine01.newLineChars' and\n"+
			"'plainTextLine02.newLineChars'"+
			"are ARE NOT equal.\n"+
			"HOWEVER, areEqual == true\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	plainTextLine06 := TextLineSpecPlainText{}

	areEqual =
		plainTextLine06.Equal(
			&plainTextLine01)

	if areEqual == true {
		t.Errorf("%v\n"+
			"Error: txtLinePlainTextElectron2.equal()\n"+
			"Expected areEqual == false because input\n"+
			"'plainTextLine01' and 'plainTextLine06'"+
			"are ARE NOT equal.\n"+
			"HOWEVER, areEqual == true\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	plainTextLine07 := TextLineSpecPlainText{}

	areEqual =
		plainTextLine07.EqualITextLine(
			&plainTextLine01)

	if areEqual == true {
		t.Errorf("%v\n"+
			"Error: plainTextLine07.EqualITextLine()\n"+
			"Expected areEqual == false because 'plainTextLine07'\n"+
			"and input parameter 'plainTextLine01'"+
			"are ARE NOT equal.\n"+
			"HOWEVER, areEqual == true\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecPlainText_EqualITextLine_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_EqualITextLine_000100()",
		"")

	expectedLeftMarginChars := []rune{' ', ' ', ' '}
	expectedRightMarginChars := []rune{' ', ' ', ' '}
	expectedNewLineChars := []rune{'\n', '\n'}

	expectedTextString := "How now brown cow!"

	plainTextLine01,
		err := TextLineSpecPlainText{}.NewPtrPlainText(
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

	var standardLineSpec *TextLineSpecStandardLine

	standardLineSpec = TextLineSpecStandardLine{}.NewPtr()

	plainTextLine02 := TextLineSpecPlainText{}

	areEqual :=
		plainTextLine02.EqualITextLine(
			standardLineSpec)

	if areEqual == true {
		t.Errorf("%v\n"+
			"Error: plainTextLine02.EqualITextLine()\n"+
			"Expected areEqual == false because input\n"+
			"parameter 'standardLineSpec is the wrong type.\n"+
			"HOWEVER, areEqual == true\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	areEqual =
		plainTextLine01.EqualITextLine(
			standardLineSpec)

	if areEqual == true {
		t.Errorf("%v\n"+
			"Error: plainTextLine01.EqualITextLine()\n"+
			"Expected areEqual == false because input\n"+
			"parameter 'standardLineSpec is the wrong type.\n"+
			"HOWEVER, areEqual == true\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	err =
		plainTextLine02.CopyIn(
			plainTextLine01,
			ePrefix.XCtx(
				"plainTextLine01->plainTextLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	areEqual =
		plainTextLine01.EqualITextLine(
			&plainTextLine02)

	if areEqual == false {
		t.Errorf("%v\n"+
			"Error: plainTextLine01.EqualITextLine(plainTextLine02)\n"+
			"Expected areEqual == true because \n"+
			"'plainTextLine01' == 'plainTextLine02'\n"+
			"HOWEVER, areEqual == false\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecPlainText_setDefaultPlainTextSpec_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_setDefaultPlainTextSpec_000100()",
		"")

	leftMarginSpaces := 2
	rightMarginSpaces := 3
	textString := "How now brown cow"

	txtLinePlainTextNanobot := textLineSpecPlainTextNanobot{}

	err :=
		txtLinePlainTextNanobot.setDefaultPlainTextSpec(
			nil,
			leftMarginSpaces,
			rightMarginSpaces,
			textString,
			ePrefix.XCtx(
				"'plainTxtLine' == 'nil'"))

	if err == nil {
		t.Errorf("%v\n"+
			"Error: txtLinePlainTextNanobot.setDefaultPlainTextSpec()\n"+
			"Expected an error return because input parameter\n"+
			"'plainTxtLine' == 'nil'\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	plainTextLine01 := TextLineSpecPlainText{}

	txtLinePlainTextNanobot2 := textLineSpecPlainTextNanobot{}

	err =
		txtLinePlainTextNanobot2.setDefaultPlainTextSpec(
			&plainTextLine01,
			leftMarginSpaces,
			rightMarginSpaces,
			textString,
			ePrefix.XCtx(
				"plainTextLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err =
		txtLinePlainTextNanobot2.setDefaultPlainTextSpec(
			&plainTextLine01,
			-1,
			rightMarginSpaces,
			textString,
			ePrefix.XCtx(
				"plainTextLine01"))

	if err == nil {
		t.Errorf("%v\n"+
			"Error: txtLinePlainTextNanobot2.setDefaultPlainTextSpec()\n"+
			"Expected an error return because input parameter\n"+
			"'leftMarginSpaces' == -1\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	err =
		txtLinePlainTextNanobot2.setDefaultPlainTextSpec(
			&plainTextLine01,
			leftMarginSpaces,
			-1,
			textString,
			ePrefix.XCtx(
				"plainTextLine01"))

	if err == nil {
		t.Errorf("%v\n"+
			"Error: txtLinePlainTextNanobot2.setDefaultPlainTextSpec()\n"+
			"Expected an error return because input parameter\n"+
			"'rightMarginSpaces' == -1\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	err =
		txtLinePlainTextNanobot2.setDefaultPlainTextSpec(
			&plainTextLine01,
			1000001,
			rightMarginSpaces,
			textString,
			ePrefix.XCtx(
				"plainTextLine01"))

	if err == nil {
		t.Errorf("%v\n"+
			"Error: txtLinePlainTextNanobot2.setDefaultPlainTextSpec()\n"+
			"Expected an error return because input parameter\n"+
			"'leftMarginSpaces' == 1,000,001\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	err =
		txtLinePlainTextNanobot2.setDefaultPlainTextSpec(
			&plainTextLine01,
			leftMarginSpaces,
			1000001,
			textString,
			ePrefix.XCtx(
				"plainTextLine01"))

	if err == nil {
		t.Errorf("%v\n"+
			"Error: txtLinePlainTextNanobot2.setDefaultPlainTextSpec()\n"+
			"Expected an error return because input parameter\n"+
			"'rightMarginSpaces' == 1,000,001\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	textString = ""

	err =
		txtLinePlainTextNanobot2.setDefaultPlainTextSpec(
			&plainTextLine01,
			leftMarginSpaces,
			rightMarginSpaces,
			textString,
			ePrefix.XCtx(
				"plainTextLine01"))

	if err == nil {
		t.Errorf("%v\n"+
			"Error: txtLinePlainTextNanobot2.setDefaultPlainTextSpec()\n"+
			"Expected an error return because input parameter\n"+
			"'textString' is an empty string.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	textString = strings.Repeat("x", 1000001)

	err =
		txtLinePlainTextNanobot2.setDefaultPlainTextSpec(
			&plainTextLine01,
			leftMarginSpaces,
			rightMarginSpaces,
			textString,
			ePrefix.XCtx(
				"plainTextLine01"))

	if err == nil {
		t.Errorf("%v\n"+
			"Error: txtLinePlainTextNanobot2.setDefaultPlainTextSpec()\n"+
			"Expected an error return because input parameter\n"+
			"'textString' has a length of 1,000,001 characters.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	textString = ""

}

func TestTextLineSpecPlainText_Read_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_Read_000100()",
		"")

	leftMarginSpaces := 2
	rightMarginSpaces := 2
	textString := "How now brown cow"

	expectedTextStr :=
		strings.Repeat(" ", leftMarginSpaces) +
			textString +
			strings.Repeat(" ", rightMarginSpaces) +
			"\n"

	lenExpectedStr := len(expectedTextStr)

	plainTextLine01 := TextLineSpecPlainText{}

	err :=
		plainTextLine01.SetPlainTextDefault(
			leftMarginSpaces,
			rightMarginSpaces,
			textString,
			ePrefix.XCtx(
				"plainTextLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	p := make([]byte, lenExpectedStr+1)

	var n, readBytesCnt int
	var actualStr string

	for {

		n,
			err = plainTextLine01.Read(p)

		if n == 0 {
			break
		}

		actualStr += string(p[:n])
		readBytesCnt += n
	}

	if err != nil &&
		err != io.EOF {
		t.Errorf("%v\n"+
			"Error Returned From plainTextLine01.Read(p)\n"+
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

	if plainTextLine01.textLineReader != nil {
		t.Errorf("%v\n"+
			"Error: After completing Read Operation\n"+
			"plainTextLine01.textLineReader != 'nil'\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	if readBytesCnt != lenExpectedStr {
		t.Errorf("%v\n"+
			"Byte Length Error: plainTextLine01.Read(p)\n"+
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

func TestTextLineSpecPlainText_Read_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_Read_000200()",
		"")

	leftMarginSpaces := 2
	rightMarginSpaces := 2
	textString := "How now brown cow"

	expectedTextStr :=
		strings.Repeat(" ", leftMarginSpaces) +
			textString +
			strings.Repeat(" ", rightMarginSpaces) +
			"\n"

	lenExpectedStr := len(expectedTextStr)

	plainTextLine01 := TextLineSpecPlainText{}

	err :=
		plainTextLine01.SetPlainTextDefault(
			leftMarginSpaces,
			rightMarginSpaces,
			textString,
			ePrefix.XCtx(
				"plainTextLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	p := make([]byte, 1)

	var actualStr string

	var n, readBytesCnt int

	for {

		n,
			err = plainTextLine01.Read(p)

		if n == 0 {
			break
		}

		actualStr += string(p[:n])
		readBytesCnt += n

	}

	if err != nil &&
		err != io.EOF {
		t.Errorf("%v\n"+
			"Error Returned From plainTextLine01.Read(p)\n"+
			"Error = \n%v\n",
			ePrefix.XCtxEmpty().String(),
			err.Error())

		return
	}

	if err != nil &&
		err != io.EOF {
		t.Errorf("%v\n"+
			"Error Returned From plainTextLine01.Read(p)\n"+
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
			"Byte Length Error: plainTextLine01.Read(p)\n"+
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

	if plainTextLine01.textLineReader != nil {
		t.Errorf("%v\n"+
			"Error: After a successful series of byte reads,\n"+
			"plainTextLine01.textLineReader pointer has NOT\n"+
			"BEEN RESET TO 'nil'!\n",
			ePrefix.XCtxEmpty().String())
		return
	}

	p = make([]byte, 100)
	readBytesCnt = 0
	actualStr = ""

	for {

		n,
			err = plainTextLine01.Read(p)

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
			"Error Returned From plainTextLine01.Read(p)\n"+
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

func TestTextLineSpecPlainText_Read_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_Read_000300()",
		"")

	leftMarginSpaces := 2
	rightMarginSpaces := 2
	textString := "How now brown cow"

	expectedTextStr :=
		strings.Repeat(" ", leftMarginSpaces) +
			textString +
			strings.Repeat(" ", rightMarginSpaces) +
			"\n"

	lenExpectedStr := len(expectedTextStr)

	plainTextLine01 := TextLineSpecPlainText{}

	err :=
		plainTextLine01.SetPlainTextSpec(
			[]rune(strings.Repeat(" ", leftMarginSpaces)),
			[]rune(strings.Repeat(" ", rightMarginSpaces)),
			textString,
			[]rune{'\n'},
			false,
			ePrefix.XCtx(
				"plainTextLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

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
	plainTxtNanobot := textLineSpecPlainTextNanobot{}

	formattedTxtStr,
		err =
		plainTxtNanobot.getFormattedText(
			&plainTextLine01,
			ePrefix.XCtx("plainTextLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	p = make([]byte, 0)

	plainTextLine01.textLineReader =
		strings.NewReader(formattedTxtStr)

	n,
		err = txtSpecAtom.readBytes(
		plainTextLine01.textLineReader,
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
			plainTextLine01.textLineReader,
			p,
			ePrefix.XCtx("plainTextLine is valid"))

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

	return
}

func TestTextLineSpecPlainText_ReaderInitialize_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_Read_000200()",
		"")

	leftMarginSpaces := 2
	rightMarginSpaces := 2
	textString := "How now brown cow"

	expectedTextStr :=
		strings.Repeat(" ", leftMarginSpaces) +
			textString +
			strings.Repeat(" ", rightMarginSpaces) +
			"\n"

	lenExpectedStr := len(expectedTextStr)

	plainTextLine01 := TextLineSpecPlainText{}

	err :=
		plainTextLine01.SetPlainTextDefault(
			leftMarginSpaces,
			rightMarginSpaces,
			textString,
			ePrefix.XCtx(
				"plainTextLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	p := make([]byte, 5)

	var n int

	n,
		err = plainTextLine01.Read(p)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by plainTextLine01.Read(p)\n"+
			"Error:\n%v\n",
			ePrefix.XCtxEmpty().String(),
			err.Error())

		return
	}

	if n != 5 {
		t.Errorf("%v\n"+
			"Error: plainTextLine01.Read(p)\n"+
			"Expected n == 5\n"+
			"Instead, n == %v\n",
			ePrefix.XCtxEmpty().String(),
			n)

		return
	}

	p = make([]byte, 100)

	plainTextLine01.ReaderInitialize()

	var readBytesCnt int
	var actualStr string

	for {

		n,
			err = plainTextLine01.Read(p)

		if n == 0 {
			break
		}

		actualStr += string(p[:n])
		readBytesCnt += n
	}

	if err != nil &&
		err != io.EOF {
		t.Errorf("%v\n"+
			"Error Returned From plainTextLine01.Read(p)\n"+
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

	if plainTextLine01.textLineReader != nil {
		t.Errorf("%v\n"+
			"Error: After completing Read Operation\n"+
			"plainTextLine01.textLineReader != 'nil'\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	if readBytesCnt != lenExpectedStr {
		t.Errorf("%v\n"+
			"Byte Length Error: plainTextLine01.Read(p)\n"+
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

	if plainTextLine01.textLineReader != nil {
		t.Errorf("%v Test #1\n"+
			"Completed Read Operation but plainTextLine01.textLineReader\n"+
			"is NOT equal to 'nil'!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	p = make([]byte, 100)
	actualStr = ""
	readBytesCnt = 0

	for {

		n,
			err = plainTextLine01.Read(p)

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
			"Error Returned From plainTextLine01.Read(p)\n"+
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

	if plainTextLine01.textLineReader != nil {
		t.Errorf("%v Test #2\n"+
			"Completed Read Operation but plainTextLine01.textLineReader\n"+
			"is NOT equal to 'nil'!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecPlainText_TextLineBuilder_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_TextLineBuilder_000100()",
		"")

	leftMarginSpaces := 2
	rightMarginSpaces := 3
	textString := "How now brown cow"

	leftMargin := strings.Repeat(" ", leftMarginSpaces)

	rightMargin := strings.Repeat(" ", rightMarginSpaces)

	expectedTextStr :=
		leftMargin +
			textString +
			rightMargin +
			"\n"

	plainTextLine01 := TextLineSpecPlainText{}

	err :=
		plainTextLine01.SetPlainTextSpecRunes(
			[]rune(leftMargin),
			[]rune(rightMargin),
			[]rune(textString),
			[]rune{'\n'},
			false,
			ePrefix.XCtx(
				"plainTextLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sb := strings.Builder{}

	err = plainTextLine01.TextLineBuilder(
		&sb,
		ePrefix.XCtx("plainTextLine01->sb"))

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

func TestTextLineSpecPlainText_TextLineBuilder_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_TextLineBuilder_000100()",
		"")

	plainTextLine01 := TextLineSpecPlainText{}
	sb := strings.Builder{}

	err := plainTextLine01.TextLineBuilder(
		&sb,
		ePrefix.XCtx("empty plainTextLine01->sb"))

	if err == nil {
		t.Errorf("%v\n"+
			"Error:\n"+
			"Expected error return from plainTextLine01.TextLineBuilder()\n"+
			"because 'plainTextLine01' is empty.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	leftMarginSpaces := 2
	rightMarginSpaces := 3
	textString := "How now brown cow"

	leftMargin := strings.Repeat(" ", leftMarginSpaces)

	rightMargin := strings.Repeat(" ", rightMarginSpaces)

	expectedTextStr :=
		leftMargin +
			textString +
			rightMargin +
			"\n"

	err =
		plainTextLine01.SetPlainTextSpecRunes(
			[]rune(leftMargin),
			[]rune(rightMargin),
			[]rune(textString),
			[]rune{'\n'},
			false,
			ePrefix.XCtx(
				"plainTextLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = plainTextLine01.TextLineBuilder(
		nil,
		ePrefix.XCtx("plainTextLine01->nil sb"))

	if err == nil {
		t.Errorf("%v\n"+
			"Error:\n"+
			"Expected error return from plainTextLine01.TextLineBuilder()\n"+
			"because strings.Builder pointer is 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	err = plainTextLine01.TextLineBuilder(
		&sb,
		ePrefix.XCtx("valid plainTextLine01->sb"))

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

func TestTextLineSpecPlainText_TextTypeName_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_TextTypeName_000100()",
		"")

	plainTextLine01 := TextLineSpecPlainText{}

	actualTxtTypeName :=
		plainTextLine01.TextTypeName()

	expectedTxtTypeName := "TextLineSpecPlainText"

	if expectedTxtTypeName !=
		actualTxtTypeName {
		t.Errorf("%v\n"+
			"Error: plainTextLine01.TextTypeName()\n"+
			"Expected Text Type Name = '%v'\n"+
			"Instead, Text Type Name = '%v'\n",
			ePrefix.String(),
			expectedTxtTypeName,
			actualTxtTypeName)
	}

	return
}

func TestTextLineSpecPlainText_TextLineSpecName_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_TextLineSpecName_000100()",
		"")

	plainTextLine01 := TextLineSpecPlainText{}

	actualTxtLineSpecName :=
		plainTextLine01.TextLineSpecName()

	expectedTxtLineSpecName := "TextLineSpecPlainText"

	if expectedTxtLineSpecName !=
		actualTxtLineSpecName {
		t.Errorf("%v\n"+
			"Error: plainTextLine01.TextLineSpecName()\n"+
			"Expected Text Line Spec Name = '%v'\n"+
			"Instead, Text Line Spec Name = '%v'\n",
			ePrefix.String(),
			expectedTxtLineSpecName,
			actualTxtLineSpecName)
	}

	return
}
