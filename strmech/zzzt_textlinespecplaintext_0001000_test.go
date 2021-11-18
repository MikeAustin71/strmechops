package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"strings"
	"testing"
)

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

	plainTextLine05 := TextLineSpecPlainText{}

	err = plainTextLine05.SetPlainTextDefault(
		leftMargin,
		rightMargin,
		textString,
		ePrefix.XCtx("plainTextLine05"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = plainTextLine05.IsValidInstanceError(
		ePrefix.XCtx("plainTextLine05"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	plainTextLine06 := TextLineSpecPlainText{}

	err =
		plainTextLine06.CopyIn(
			&plainTextLine05,
			TextFieldSpecDateTime{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from plainTextLine06."+
			"CopyIn()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

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

	var plainTextLine04 TextLineSpecPlainText

	plainTextLine04,
		err = TextLineSpecPlainText{}.NewPlainText(
		expectedLeftMarginChars,
		expectedRightMarginChars,
		expectedTextString,
		expectedNewLineChars,
		false,
		ePrefix.XCtx("plainTextLine04"))

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

	_,
		err =
		plainTextLine04.CopyOut(
			TextFieldSpecDateTime{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from plainTextLine04."+
			"CopyOut()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

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
				"plainTextLine03->_"))

	if err == nil {
		t.Errorf("%v\n"+
			"Error: plainTextLine03.CopyOutITextLine()\n"+
			"Expected an error return because 'plainTextLine03' is empty!\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	var plainTextLine04 TextLineSpecPlainText

	plainTextLine04,
		err = TextLineSpecPlainText{}.NewPlainTextStrings(
		expectedLeftMarginChars,
		expectedRightMarginChars,
		expectedTextString,
		expectedNewLineChars,
		false,
		ePrefix.XCtx(
			"plainTextLine04"))

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

	_,
		err =
		plainTextLine04.CopyOutITextLine(
			TextFieldSpecDateTime{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from plainTextLine04."+
			"CopyOutITextLine()\n"+
			"because 'errorPrefix' is invalid.\n"+
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

	var plainTextLine04 TextLineSpecPlainText

	plainTextLine04,
		err = TextLineSpecPlainText{}.NewDefault(
		leftMargin,
		rightMargin,
		textString,
		ePrefix.XCtx("plainTextLine04"))

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

	_,
		err = plainTextLine04.CopyOutPtr(
		TextFieldSpecDateTime{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from plainTextLine04."+
			"CopyOutPtr()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
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

func TestTextLineSpecPlainText_GetFormattedText_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_GetFormattedText_000100()",
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

	expectedFmtTextStr :=
		string(expectedLeftMarginChars) +
			expectedTextString +
			string(expectedRightMarginChars) +
			string(expectedNewLineChars)

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedFmtTextStr),
			true)

	var actualFmtTxtStr string

	actualFmtTxtStr,
		err =
		plainTextLine01.GetFormattedText(
			ePrefix.XCtx(
				"plainTextLine01->actualFmtTxtStr"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(actualFmtTxtStr),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Expected Formatted Text String DOES NOT match\n"+
			"Actual Formatted Text String.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	_,
		err =
		plainTextLine01.GetFormattedText(
			TextFieldSpecDateTime{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from plainTextLine01."+
			"GetFormattedText()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	plainTextLine02 := TextLineSpecPlainText{}

	_,
		err = plainTextLine02.GetFormattedText(
		ePrefix.XCtx(
			"plainTextLine02"))

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from plainTextLine02."+
			"GetFormattedText()\n"+
			"because 'plainTextLine02' is empty and invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecPlainText_GetLeftMarginStr_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_GetLeftMarginStr_000100()",
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

	expectedLeftMarginStr :=
		string(expectedLeftMarginChars)

	actualLeftMarginStr :=
		plainTextLine01.GetLeftMarginStr()

	if expectedLeftMarginStr != actualLeftMarginStr {

		t.Errorf("%v - ERROR\n"+
			"Test #1 - plainTextLine01.GetLeftMarginStr()\n"+
			"Expected Left Margin String DOES NOT match\n"+
			"Actual Left Margin String.\n"+
			"Expected Left Margin String = '%v'\n"+
			"Instead, Left Margin String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedLeftMarginStr,
			actualLeftMarginStr)

		return
	}

	plainTextLine02 := TextLineSpecPlainText{}

	expectedLeftMarginStr = ""

	actualLeftMarginStr =
		plainTextLine02.GetLeftMarginStr()

	if expectedLeftMarginStr != actualLeftMarginStr {

		t.Errorf("%v - ERROR\n"+
			"Test #2 - plainTextLine02.GetLeftMarginStr()\n"+
			"Expected Left Margin String DOES NOT match\n"+
			"Actual Left Margin String.\n"+
			"Expected Left Margin String = '%v'\n"+
			"Instead, Left Margin String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedLeftMarginStr,
			actualLeftMarginStr)

		return
	}

	return
}

func TestTextLineSpecPlainText_GetLeftMarginRunes_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_GetLeftMarginRunes_000100()",
		"")

	expectedLeftMarginRunes := []rune{' ', ' ', ' '}
	expectedRightMarginRunes := []rune{' ', ' ', ' '}
	expectedNewLineRunes := []rune{'\n', '\n'}
	expectedTextString := "How now brown cow!"

	plainTextLine01,
		err := TextLineSpecPlainText{}.NewPtrPlainTextRunes(
		expectedLeftMarginRunes,
		expectedRightMarginRunes,
		[]rune(expectedTextString),
		expectedNewLineRunes,
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

	actualLeftMarginRunes :=
		plainTextLine01.GetLeftMarginRunes()

	areEqual :=
		strMechPreon{}.ptr().equalRuneArrays(
			expectedLeftMarginRunes,
			actualLeftMarginRunes)

	if !areEqual {

		t.Errorf("%v - ERROR\n"+
			"Test #1 - plainTextLine01.GetLeftMarginRunes()\n"+
			"Expected Left Margin Runes DO NOT match\n"+
			"Actual Left Margin Runes.\n"+
			"Expected Left Margin Runes = '%v'\n"+
			"Instead, Left Margin Runes = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			string(expectedLeftMarginRunes),
			string(actualLeftMarginRunes))

		return
	}

	plainTextLine02 := TextLineSpecPlainText{}

	actualLeftMarginRunes =
		plainTextLine02.GetLeftMarginRunes()

	if actualLeftMarginRunes != nil {

		t.Errorf("%v - ERROR\n"+
			"Test #2 - plainTextLine02.GetLeftMarginRunes()\n"+
			"Expected Left Margin Runes to equal 'nil'\n"+
			"because 'plainTextLine02' is empty.\n"+
			"HOWEVER, actualLeftMarginRunes != 'nil'!!\n"+
			"Instead, Actual Left Margin Runes = '%v'\n"+
			"Actual Left Margin Rune Array = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			string(actualLeftMarginRunes),
			actualLeftMarginRunes)

		return
	}

	return
}

func TestTextLineSpecPlainText_GetLineTerminationChars_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_GetLineTerminationChars_000100()",
		"")

	expectedLeftMarginRunes := []rune{' ', ' ', ' '}
	expectedRightMarginRune := []rune{' ', ' ', ' '}
	expectedNewLineRunes := []rune{'\n', '\n'}
	expectedTextString := "How now brown cow!"

	plainTextLine01,
		err := TextLineSpecPlainText{}.NewPlainText(
		expectedLeftMarginRunes,
		expectedRightMarginRune,
		expectedTextString,
		expectedNewLineRunes,
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

	actualLineTerminationChars :=
		plainTextLine01.GetLineTerminationChars()

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			expectedNewLineRunes,
			true)

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(actualLineTerminationChars),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("\n%v - ERROR\n"+
			"Test #1 - plainTextLine01.GetLineTerminationChars()\n"+
			"Expected Line Termination String DOES NOT match\n"+
			"Actual Line Termination String.\n"+
			"Expected Line Termination String = '%v'\n"+
			"Instead, Line Termination String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	plainTextLine02 := TextLineSpecPlainText{}

	actualLineTerminationChars =
		plainTextLine02.GetLineTerminationChars()

	printableActualStr =
		sMech.ConvertNonPrintableChars(
			[]rune(actualLineTerminationChars),
			true)

	printableExpectedStr =
		sMech.ConvertNonPrintableChars(
			[]rune(""),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("\n%v - ERROR\n"+
			"Test #2 - plainTextLine02.GetLineTerminationChars()\n"+
			"Expected Line Termination String DOES NOT match\n"+
			"Actual Line Termination String.\n"+
			"Expected Line Termination String = '%v'\n"+
			"Instead, Line Termination String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	return
}

func TestTextLineSpecPlainText_GetLineTerminationRunes_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_GetLineTerminationRunes_000100()",
		"")

	expectedLeftMarginRunes := []rune{' ', ' ', ' '}
	expectedRightMarginRune := []rune{' ', ' ', ' '}
	expectedNewLineRunes := []rune{'\n', '\n'}
	expectedTextString := "How now brown cow!"

	plainTextLine01,
		err := TextLineSpecPlainText{}.NewPlainText(
		expectedLeftMarginRunes,
		expectedRightMarginRune,
		expectedTextString,
		expectedNewLineRunes,
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

	actualLineTerminationRunes :=
		plainTextLine01.GetLineTerminationRunes()

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			expectedNewLineRunes,
			true)

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			actualLineTerminationRunes,
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("\n%v - ERROR\n"+
			"Test #1 - plainTextLine01.GetLineTerminationRunes()\n"+
			"Expected Line Termination String DOES NOT match\n"+
			"Actual Line Termination String.\n"+
			"Expected Line Termination String = '%v'\n"+
			"Instead, Line Termination String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	areEqual :=
		strMechPreon{}.ptr().equalRuneArrays(
			expectedNewLineRunes,
			actualLineTerminationRunes)

	if !areEqual {

		t.Errorf("\n%v - ERROR\n"+
			"Test #2 - plainTextLine01.GetLineTerminationRunes()\n"+
			"Rune Array Comparison\n"+
			"Expected Line Termination String DOES NOT match\n"+
			"Actual Line Termination String.\n"+
			"Expected Line Termination String = '%v'\n"+
			"Instead, Line Termination String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			string(expectedNewLineRunes),
			string(actualLineTerminationRunes))

		return
	}

	plainTextLine02 := TextLineSpecPlainText{}

	actualLineTerminationRunes =
		plainTextLine02.GetLineTerminationRunes()

	printableActualStr =
		sMech.ConvertNonPrintableChars(
			actualLineTerminationRunes,
			true)

	expectedNewLineRunes = nil

	printableExpectedStr =
		sMech.ConvertNonPrintableChars(
			expectedNewLineRunes,
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("\n%v - ERROR\n"+
			"Test #3 - plainTextLine02.GetLineTerminationRunes()\n"+
			"Expected Line Termination String DOES NOT match\n"+
			"Actual Line Termination String.\n"+
			"Expected Line Termination String = '%v'\n"+
			"Instead, Line Termination String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	areEqual =
		strMechPreon{}.ptr().equalRuneArrays(
			expectedNewLineRunes,
			actualLineTerminationRunes)

	if !areEqual {

		t.Errorf("\n%v - ERROR\n"+
			"Test #4 - plainTextLine02.GetLineTerminationRunes()\n"+
			"Rune Array Comparison\n"+
			"Expected Line Termination String DOES NOT match\n"+
			"Actual Line Termination String.\n"+
			"Expected Line Termination String = '%v'\n"+
			"Instead, Line Termination String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			string(expectedNewLineRunes),
			string(actualLineTerminationRunes))

		return
	}

	return
}

func TestTextLineSpecPlainText_GetRightMarginStr_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_GetRightMarginStr_000100()",
		"")

	expectedLeftMarginRunes := []rune{' ', ' ', ' '}
	expectedRightMarginRunes := []rune{' ', ' ', ' '}
	expectedNewLineRunes := []rune{'\n', '\n'}
	expectedTextString := "How now brown cow!"

	plainTextLine01,
		err := TextLineSpecPlainText{}.NewPlainText(
		expectedLeftMarginRunes,
		expectedRightMarginRunes,
		expectedTextString,
		expectedNewLineRunes,
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

	expectedRightMarginStr :=
		string(expectedRightMarginRunes)

	actualRightMarginStr :=
		plainTextLine01.GetRightMarginStr()

	if expectedRightMarginStr != actualRightMarginStr {

		t.Errorf("%v - ERROR\n"+
			"Test #1 - plainTextLine01.GetRightMarginStr()\n"+
			"Expected Right Margin String DOES NOT match\n"+
			"Actual Right Margin String.\n"+
			"Expected Right Margin String = '%v'\n"+
			"Instead, Right Margin String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedRightMarginStr,
			actualRightMarginStr)

		return
	}

	plainTextLine02 := TextLineSpecPlainText{}

	expectedRightMarginStr = ""

	actualRightMarginStr =
		plainTextLine02.GetRightMarginStr()

	if expectedRightMarginStr != actualRightMarginStr {

		t.Errorf("%v - ERROR\n"+
			"Test #2 - plainTextLine02.GetRightMarginStr()\n"+
			"Expected Right Margin String DOES NOT match\n"+
			"Actual Right Margin String.\n"+
			"Expected Right Margin String = '%v'\n"+
			"Instead, Right Margin String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedRightMarginStr,
			actualRightMarginStr)

		return
	}

	return
}

func TestTextLineSpecPlainText_GetRightMarginRunes_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_GetRightMarginRunes_000100()",
		"")

	expectedLeftMarginRunes := []rune{' ', ' ', ' '}
	expectedRightMarginRunes := []rune{' ', ' ', ' '}
	expectedNewLineRunes := []rune{'\n', '\n'}
	expectedTextString := "How now brown cow!"

	plainTextLine01,
		err := TextLineSpecPlainText{}.NewPtrPlainTextRunes(
		expectedLeftMarginRunes,
		expectedRightMarginRunes,
		[]rune(expectedTextString),
		expectedNewLineRunes,
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

	actualRightMarginRunes :=
		plainTextLine01.GetRightMarginRunes()

	areEqual :=
		strMechPreon{}.ptr().equalRuneArrays(
			expectedRightMarginRunes,
			actualRightMarginRunes)

	if !areEqual {

		t.Errorf("%v - ERROR\n"+
			"Test #1 - plainTextLine01.GetRightMarginRunes()\n"+
			"Expected Right Margin Runes DO NOT match\n"+
			"Actual Right Margin Runes.\n"+
			"Expected Right Margin Runes = '%v'\n"+
			"Instead, Right Margin Runes = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			string(expectedRightMarginRunes),
			string(actualRightMarginRunes))

		return
	}

	plainTextLine02 := TextLineSpecPlainText{}

	actualRightMarginRunes =
		plainTextLine02.GetRightMarginRunes()

	if actualRightMarginRunes != nil {

		t.Errorf("%v - ERROR\n"+
			"Test #2 - plainTextLine02.GetRightMarginRunes()\n"+
			"Expected Right Margin Runes to equal 'nil'\n"+
			"because 'plainTextLine02' is empty.\n"+
			"HOWEVER, actualRightMarginRunes != 'nil'!!\n"+
			"Instead, Actual Right Margin Runes = '%v'\n"+
			"Actual Right Margin Rune Array = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			string(actualRightMarginRunes),
			actualRightMarginRunes)

		return
	}

	return
}

func TestTextLineSpecPlainText_GetTextString_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"\nTestTextLineSpecPlainText_GetTextString_000100()",
		"")

	expectedLeftMarginRunes := []rune{' ', ' ', ' '}
	expectedRightMarginRunes := []rune{' ', ' ', ' '}
	expectedNewLineRunes := []rune{'\n', '\n'}
	expectedTextString := "How now brown cow!"

	plainTextLine01,
		err := TextLineSpecPlainText{}.NewPlainText(
		expectedLeftMarginRunes,
		expectedRightMarginRunes,
		expectedTextString,
		expectedNewLineRunes,
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

	actualTextStr :=
		plainTextLine01.GetTextString()

	if expectedTextString != actualTextStr {

		t.Errorf("%v - ERROR\n"+
			"Test #1 - plainTextLine01.GetTextString()\n"+
			"Expected Text String DOES NOT match\n"+
			"Actual Text String.\n"+
			"Expected Text String = '%v'\n"+
			"Instead, Text String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedTextString,
			actualTextStr)

		return
	}

	plainTextLine02 := TextLineSpecPlainText{}

	expectedTextString = ""

	actualTextStr =
		plainTextLine02.GetTextString()

	if expectedTextString != actualTextStr {

		t.Errorf("%v - ERROR\n"+
			"Test #2 - plainTextLine02.GetTextString()\n"+
			"Expected Text String DOES NOT match\n"+
			"Actual Text String.\n"+
			"Expected Text String = '%v'\n"+
			"Instead, Text String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedTextString,
			actualTextStr)

		return
	}

	return
}

func TestTextLineSpecPlainText_GetTurnLineTerminatorOff_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"\nTestTextLineSpecPlainText_GetTurnLineTerminatorOff_000100()",
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

	actualTurnLineTerminatorOff :=
		plainTextLine01.GetTurnLineTerminatorOff()

	if actualTurnLineTerminatorOff == true {

		t.Errorf("%v - ERROR\n"+
			"Test #1\n"+
			"TurnLineTerminatorOff value is INVALID!\n"+
			"Expected TurnLineTerminatorOff = 'true'\n"+
			"Instead, TurnLineTerminatorOff = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			actualTurnLineTerminatorOff)

		return
	}

	expectedFmtTextStr :=
		string(expectedLeftMarginChars) +
			expectedTextString +
			string(expectedRightMarginChars) +
			string(expectedNewLineChars)

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedFmtTextStr),
			true)

	var actualFmtTextStr string

	actualFmtTextStr,
		err = plainTextLine01.GetFormattedText(
		ePrefix.XCtx(
			"Test#1 plainTextLine01->actualFmtTextStr"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(actualFmtTextStr),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #1\n"+
			"Expected Formatted Text String DOES NOT match\n"+
			"Actual Formatted Text String.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	expectedFmtTextStr =
		string(expectedLeftMarginChars) +
			expectedTextString +
			string(expectedRightMarginChars)

	printableExpectedStr =
		sMech.ConvertNonPrintableChars(
			[]rune(expectedFmtTextStr),
			true)

	plainTextLine01.TurnAutoLineTerminationOff()

	actualFmtTextStr,
		err = plainTextLine01.GetFormattedText(
		ePrefix.XCtx(
			"Test#2 plainTextLine01->actualFmtTextStr"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	printableActualStr =
		sMech.ConvertNonPrintableChars(
			[]rune(actualFmtTextStr),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #2\n"+
			"Expected Formatted Text String DOES NOT match\n"+
			"Actual Formatted Text String.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	actualTurnLineTerminatorOff =
		plainTextLine01.GetTurnLineTerminatorOff()

	if actualTurnLineTerminatorOff == false {

		t.Errorf("%v - ERROR\n"+
			"Test #2\n"+
			"TurnLineTerminatorOff value is INVALID!\n"+
			"Expected TurnLineTerminatorOff = 'false'\n"+
			"Instead, TurnLineTerminatorOff = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			actualTurnLineTerminatorOff)

		return
	}

	plainTextLine01.TurnAutoLineTerminationOn()

	actualTurnLineTerminatorOff =
		plainTextLine01.GetTurnLineTerminatorOff()

	if actualTurnLineTerminatorOff == true {

		t.Errorf("%v - ERROR\n"+
			"Test #3\n"+
			"TurnLineTerminatorOff value is INVALID!\n"+
			"Expected TurnLineTerminatorOff = 'true'\n"+
			"Instead, TurnLineTerminatorOff = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			actualTurnLineTerminatorOff)

		return
	}

	expectedFmtTextStr =
		string(expectedLeftMarginChars) +
			expectedTextString +
			string(expectedRightMarginChars) +
			string(expectedNewLineChars)

	printableExpectedStr =
		sMech.ConvertNonPrintableChars(
			[]rune(expectedFmtTextStr),
			true)

	actualFmtTextStr,
		err = plainTextLine01.GetFormattedText(
		ePrefix.XCtx(
			"Test#2 plainTextLine01->actualFmtTextStr"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	printableActualStr =
		sMech.ConvertNonPrintableChars(
			[]rune(actualFmtTextStr),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #3\n"+
			"Expected Formatted Text String DOES NOT match\n"+
			"Actual Formatted Text String.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	plainTextLine02 := TextLineSpecPlainText{}

	actualTurnLineTerminatorOff =
		plainTextLine02.GetTurnLineTerminatorOff()

	if actualTurnLineTerminatorOff == true {

		t.Errorf("%v - ERROR\n"+
			"Test #4 - plainTextLine02.GetTurnLineTerminatorOff()\n"+
			"TurnLineTerminatorOff value is INVALID!\n"+
			"Expected TurnLineTerminatorOff = 'false'\n"+
			"Instead, TurnLineTerminatorOff = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			actualTurnLineTerminatorOff)

		return
	}

}

func TestTextLineSpecPlainText_IsValidInstanceError_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_IsValidInstanceError_000100()",
		"")

	plainTextLine01 := TextLineSpecPlainText{}

	err :=
		plainTextLine01.IsValidInstanceError(
			ePrefix.XCtx(
				"plainTextLine01"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from plainTextLine01."+
			"IsValidInstanceError()\n"+
			"because 'plainTextLine01' is empty and invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	expectedLeftMarginRunes := []rune{' ', ' ', ' '}
	expectedRightMarginRunes := []rune{' ', ' ', ' '}
	expectedNewLineRunes := []rune{'\n', '\n'}
	expectedTextString := "How now brown cow!"

	var plainTextLine02 TextLineSpecPlainText

	plainTextLine02,
		err = TextLineSpecPlainText{}.NewPlainText(
		expectedLeftMarginRunes,
		expectedRightMarginRunes,
		expectedTextString,
		expectedNewLineRunes,
		false,
		ePrefix.XCtx("plainTextLine02"))

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

	plainTextLine02.textString = ""

	err = plainTextLine02.IsValidInstanceError(
		ePrefix.XCtx("plainTextLine02"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from plainTextLine01."+
			"IsValidInstanceError()\n"+
			"because plainTextLine02.textString = \"\".\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	plainTextLine02.textString = expectedTextString

	plainTextLine02.newLineChars = nil

	err = plainTextLine02.IsValidInstanceError(
		ePrefix.XCtx("plainTextLine02"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from plainTextLine01."+
			"IsValidInstanceError()\n"+
			"because plainTextLine02.newLineChars = nil.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	var plainTextLine03 TextLineSpecPlainText

	plainTextLine03,
		err = TextLineSpecPlainText{}.NewPlainText(
		expectedLeftMarginRunes,
		expectedRightMarginRunes,
		expectedTextString,
		expectedNewLineRunes,
		false,
		ePrefix.XCtx("plainTextLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = plainTextLine03.IsValidInstanceError(
		TextFieldSpecDateTime{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from plainTextLine04."+
			"CopyOut()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
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

		return
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

func TestTextLineSpecPlainText_Read_000400(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_Read_000100()",
		"")

	leftMarginSpaces := 2
	rightMarginSpaces := 2
	textString := "How now brown cow"

	leftMargin := strings.Repeat(" ", leftMarginSpaces)
	rightMargin := strings.Repeat(" ", rightMarginSpaces)
	newLineTerminator := "\n"

	expectedTextStr :=
		leftMargin +
			textString +
			rightMargin +
			newLineTerminator

	lenExpectedStr := len(expectedTextStr)

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedTextStr),
			true)

	plainTextLine01,
		err := TextLineSpecPlainText{}.NewPlainTextStrings(
		leftMargin,
		rightMargin,
		textString,
		newLineTerminator,
		false,
		ePrefix.XCtx("plainTextLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	p := make([]byte, 15)

	var n, readBytesCnt int
	sb := strings.Builder{}
	sb.Grow(128)

	for {

		n,
			err = plainTextLine01.Read(p)

		if n == 0 {
			break
		}

		sb.Write(p[:n])
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

	err = plainTextLine01.TextBuilder(
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

	err := plainTextLine01.TextBuilder(
		&sb,
		ePrefix.XCtx("empty plainTextLine01->sb"))

	if err == nil {
		t.Errorf("%v\n"+
			"Error:\n"+
			"Expected error return from plainTextLine01.TextBuilder()\n"+
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

	err = plainTextLine01.TextBuilder(
		nil,
		ePrefix.XCtx("plainTextLine01->nil sb"))

	if err == nil {
		t.Errorf("%v\n"+
			"Error:\n"+
			"Expected error return from plainTextLine01.TextBuilder()\n"+
			"because strings.Builder pointer is 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	err = plainTextLine01.TextBuilder(
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
