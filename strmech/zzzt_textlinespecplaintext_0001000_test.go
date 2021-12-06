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
			StrMech{})

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
			StrMech{})

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
			StrMech{})

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
		StrMech{})

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

	expectedFmtTextStr :=
		string(expectedLeftMarginRunes) +
			expectedTextString +
			string(expectedRightMarginRunes) +
			string(expectedNewLineRunes)

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
			StrMech{})

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

	plainTextLine02.newLineChars = expectedNewLineRunes

	plainTextLine02.textString =
		strings.Repeat("x", 1000001)

	err = plainTextLine02.IsValidInstanceError(
		ePrefix.XCtx("plainTextLine02"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from plainTextLine01."+
			"IsValidInstanceError()\n"+
			"because plainTextLine02.textString has\n"+
			"over 1-million characters (1,000,001).\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	plainTextLine02.textString = expectedTextString

	plainTextLine02.leftMarginChars =
		[]rune(strings.Repeat("x", 1000001))

	err = plainTextLine02.IsValidInstanceError(
		ePrefix.XCtx("plainTextLine02"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from plainTextLine01."+
			"IsValidInstanceError()\n"+
			"because plainTextLine02.leftMarginChars has\n"+
			"over 1-million characters (1,000,001)\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	plainTextLine02.leftMarginChars =
		make([]rune, 0)

	plainTextLine02.leftMarginChars = expectedLeftMarginRunes

	plainTextLine02.rightMarginChars =
		[]rune(strings.Repeat("x", 1000001))

	err = plainTextLine02.IsValidInstanceError(
		ePrefix.XCtx("plainTextLine02"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from plainTextLine01."+
			"IsValidInstanceError()\n"+
			"because plainTextLine02.rightMarginChars has\n"+
			"over 1-million characters (1,000,001)\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	plainTextLine02.rightMarginChars =
		make([]rune, 0)

	plainTextLine02.rightMarginChars = expectedRightMarginRunes

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
		StrMech{})

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

func TestTextLineSpecPlainText_NewDefault_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_NewDefault_000100()",
		"")

	leftMarginSpaces := 3
	rightMarginSpaces := 3
	textString := "How now brown cow!"

	plainTextLine01,
		err := TextLineSpecPlainText{}.NewDefault(
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

	err = plainTextLine01.IsValidInstanceError(
		ePrefix.XCtx("plainTextLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	_,
		err = TextLineSpecPlainText{}.NewDefault(
		leftMarginSpaces,
		rightMarginSpaces,
		textString,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from TextLineSpecPlainText{}."+
			"NewDefault()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	_,
		err = TextLineSpecPlainText{}.NewDefault(
		1000001,
		rightMarginSpaces,
		textString,
		ePrefix.XCtx(
			"Invalid Left Margin Spaces"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from TextLineSpecPlainText{}."+
			"NewDefault()\n"+
			"because 'leftMarginSpaces' is '1,000,001'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	_,
		err = TextLineSpecPlainText{}.NewDefault(
		3,
		1000001,
		textString,
		ePrefix.XCtx(
			"Invalid Right Margin Spaces"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from TextLineSpecPlainText{}."+
			"NewDefault()\n"+
			"because 'rightMarginSpaces' is '1,000,001'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	textString = strings.Repeat(
		"x", 1000001)
	_,
		err = TextLineSpecPlainText{}.NewDefault(
		3,
		3,
		textString,
		ePrefix.XCtx(
			"Invalid Right Margin Spaces"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from TextLineSpecPlainText{}."+
			"NewDefault()\n"+
			"because 'textString' has '1,000,001' characters.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	textString = ""

	return
}

func TestTextLineSpecPlainText_NewPlainText_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_GetFormattedText_000100()",
		"")

	expectedLeftMarginRunes := []rune{' ', ' ', ' '}
	expectedRightMarginRunes := []rune{' ', ' ', ' '}
	expectedNewLineRunes := []rune{'\n'}
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

	_,
		err = TextLineSpecPlainText{}.NewPlainText(
		expectedLeftMarginRunes,
		expectedRightMarginRunes,
		expectedTextString,
		expectedNewLineRunes,
		false,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from TextLineSpecPlainText{}."+
			"NewPlainText()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecPlainText_NewPlainTextRunes_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_NewPlainTextRunes_000100()",
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

	turnLineTerminatorOff :=
		plainTextLine01.GetTurnLineTerminatorOff()

	if turnLineTerminatorOff == false {

		t.Errorf("\n%v - ERROR\n"+
			"plainTextLine01.GetTurnLineTerminatorOff()\n"+
			"Expected turnLineTerminatorOff == 'true'\n"+
			"Instead, turnLineTerminatorOff == 'false'\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	_,
		err = TextLineSpecPlainText{}.NewPlainTextRunes(
		expectedLeftMarginChars,
		expectedRightMarginChars,
		expectedTextChars,
		expectedNewLineChars,
		true,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from TextLineSpecPlainText{}."+
			"NewPlainTextRunes()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecPlainText_NewPlainTextStrings_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_NewPlainTextStrings_000100()",
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

	_,
		err = TextLineSpecPlainText{}.NewPlainTextStrings(
		expectedLeftMarginChars,
		expectedRightMarginChars,
		expectedTextString,
		expectedNewLineChars,
		false,
		StrMech{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from TextLineSpecPlainText{}."+
			"NewPlainTextStrings()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecPlainText_NewPtrDefault_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_NewPtrDefault_000100()",
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

	expectedLeftMarginChars := []rune{' ', ' ', ' '}
	expectedRightMarginChars := []rune{' ', ' ', ' '}
	expectedNewLineChars := []rune{'\n'}

	expectedTextString := "How now brown cow!"
	expectedTextChars := []rune(expectedTextString)

	var plainTextLine02 TextLineSpecPlainText

	plainTextLine02,
		err = TextLineSpecPlainText{}.NewPlainTextRunes(
		expectedLeftMarginChars,
		expectedRightMarginChars,
		expectedTextChars,
		expectedNewLineChars,
		false,
		ePrefix.XCtx(
			"plainTextLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !plainTextLine02.Equal(
		plainTextLine01) {

		t.Errorf("%v - ERROR\n"+
			"Expected plainTextLine01 == plainTextLine02\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	_,
		err = TextLineSpecPlainText{}.NewPtrDefault(
		leftMargin,
		rightMargin,
		textString,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from TextLineSpecPlainText{}."+
			"NewPtrDefault()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecPlainText_NewPtrPlainText_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_NewPtrPlainText_000100()",
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

	_,
		err = TextLineSpecPlainText{}.NewPtrPlainText(
		expectedLeftMarginChars,
		expectedRightMarginChars,
		expectedTextString,
		expectedNewLineChars,
		false,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from TextLineSpecPlainText{}."+
			"NewPtrPlainText()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecPlainText_NewPtrPlainTextRunes_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_NewPtrPlainTextRunes_000100()",
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

	_,
		err = TextLineSpecPlainText{}.NewPtrPlainTextRunes(
		expectedLeftMarginRunes,
		expectedRightMarginRunes,
		[]rune(expectedTextString),
		expectedNewLineRunes,
		false,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from TextLineSpecPlainText{}."+
			"NewPtrPlainTextRunes()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecPlainText_NewPtrPlainTextStrings_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_NewPtrPlainTextStrings_000100()",
		"")

	expectedLeftMarginRunes := "   "
	expectedRightMarginRunes := "   "
	expectedNewLineRunes := "\n\n"
	expectedTextString := "How now brown cow!"

	plainTextLine01,
		err := TextLineSpecPlainText{}.NewPtrPlainTextStrings(
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

	_,
		err = TextLineSpecPlainText{}.NewPtrPlainTextStrings(
		expectedLeftMarginRunes,
		expectedRightMarginRunes,
		expectedTextString,
		expectedNewLineRunes,
		false,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from TextLineSpecPlainText{}."+
			"NewPtrPlainTextStrings()\n"+
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
		"TestTextLineSpecPlainText_Read_000400()",
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

func TestTextLineSpecPlainText_Read_000500(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_Read_000500()",
		"")

	leftMarginSpaces := 2
	rightMarginSpaces := 2
	textString := "How now brown cow"

	leftMargin := strings.Repeat(" ", leftMarginSpaces)
	rightMargin := strings.Repeat(" ", rightMarginSpaces)
	newLineTerminator := "\n"

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

	err = plainTextLine01.IsValidInstanceError(
		ePrefix.XCtx("plainTextLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	plainTextLine01.textString = ""
	plainTextLine01.rightMarginChars = nil
	plainTextLine01.leftMarginChars = nil
	plainTextLine01.newLineChars = nil

	p := make([]byte, 15)

	_,
		err = plainTextLine01.Read(p)

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from plainTextLine01.Read(p)\n"+
			"because 'plainTextLine01' contains invalid data.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecPlainText_ReaderInitialize_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_ReaderInitialize_000100()",
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

	plainTextLine02 := TextLineSpecPlainText{}

	plainTextLine02.ReaderInitialize()

	return
}

func TestTextLineSpecPlainText_SetLeftMarginChars_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_SetLeftMarginChars_000100()",
		"")

	expectedLeftMarginStr := "   "
	expectedRightMarginStr := "   "
	expectedNewLineStr := "\n\n"
	expectedTextString := "How now brown cow!"

	plainTextLineZero := TextLineSpecPlainText{}

	err :=
		plainTextLineZero.SetLeftMarginChars(
			expectedLeftMarginStr,
			ePrefix.XCtx(
				"plainTextLineZero"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualLeftMarginStr :=
		plainTextLineZero.GetLeftMarginStr()

	if expectedLeftMarginStr !=
		actualLeftMarginStr {

		t.Errorf("%v Test #1\n"+
			"plainTextLineZero.GetLeftMarginStr()\n"+
			"Error: Expected 'expectedLeftMarginStr' == 'actualLeftMarginStr'\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!!\n"+
			"Expected Text String = '%v'\n"+
			"Instead, Text String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedLeftMarginStr,
			actualLeftMarginStr)

		return
	}

	var plainTextLine01 TextLineSpecPlainText

	plainTextLine01,
		err = TextLineSpecPlainText{}.NewPlainTextStrings(
		expectedLeftMarginStr,
		expectedRightMarginStr,
		expectedTextString,
		expectedNewLineStr,
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

	expectedLeftMarginStr = "!!!!"

	err =
		plainTextLine01.SetLeftMarginChars(
			expectedLeftMarginStr,
			ePrefix.XCtx(
				"plainTextLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualLeftMarginStr =
		plainTextLine01.GetLeftMarginStr()

	if expectedLeftMarginStr !=
		actualLeftMarginStr {

		t.Errorf("%v Test #2\n"+
			"plainTextLine01.GetLeftMarginStr()\n"+
			"Error: Expected 'expectedLeftMarginStr' == 'actualLeftMarginStr'\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!!\n"+
			"Expected Text String = '%v'\n"+
			"Instead, Text String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedLeftMarginStr,
			actualLeftMarginStr)

		return
	}

	err =
		plainTextLine01.SetLeftMarginChars(
			expectedLeftMarginStr,
			StrMech{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from plainTextLine01."+
			"SetLeftMarginChars()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	var plainTextLine02 TextLineSpecPlainText

	plainTextLine02,
		err = TextLineSpecPlainText{}.NewPlainTextStrings(
		expectedLeftMarginStr,
		expectedRightMarginStr,
		expectedTextString,
		expectedNewLineStr,
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

	expectedLeftMarginStr = ""

	err =
		plainTextLine02.SetLeftMarginChars(
			expectedLeftMarginStr,
			ePrefix.XCtx(
				"plainTextLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualLeftMarginStr =
		plainTextLine02.GetLeftMarginStr()

	if expectedLeftMarginStr !=
		actualLeftMarginStr {

		t.Errorf("%v Test #3\n"+
			"plainTextLine02.GetLeftMarginStr()\n"+
			"Error: Expected 'expectedLeftMarginStr' == 'actualLeftMarginStr'\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!!\n"+
			"Expected Text String = '%v'\n"+
			"Instead, Text String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedLeftMarginStr,
			actualLeftMarginStr)

		return
	}

	expectedLeftMarginStr = "   "

	var plainTextLine03 TextLineSpecPlainText

	plainTextLine03,
		err = TextLineSpecPlainText{}.NewPlainTextStrings(
		expectedLeftMarginStr,
		expectedRightMarginStr,
		expectedTextString,
		expectedNewLineStr,
		false,
		ePrefix.XCtx("plainTextLine03"))

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

	expectedLeftMarginStr =
		strings.Repeat("X", 1000001)

	err =
		plainTextLine03.SetLeftMarginChars(
			expectedLeftMarginStr,
			ePrefix.XCtx(
				"plainTextLine03"))

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from plainTextLine03."+
			"SetLeftMarginChars()\n"+
			"because 'expectedLeftMarginStr' has 1,000,001 characters.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())
	}

	expectedLeftMarginStr = ""

	return
}

func TestTextLineSpecPlainText_SetLeftMarginRunes_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_SetLeftMarginRunes_000100()",
		"")

	expectedLeftMarginRunes := []rune{' ', ' ', ' '}
	expectedRightMarginRunes := []rune{' ', ' ', ' '}
	expectedNewLineRunes := []rune{'\n', '\n'}
	expectedTextString := "How now brown cow!"

	expectedLeftMarginStr :=
		string(expectedLeftMarginRunes)

	plainTextLineZero := TextLineSpecPlainText{}

	err := plainTextLineZero.SetLeftMarginRunes(
		expectedLeftMarginRunes,
		ePrefix.XCtx(
			"plainTextLineZero"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualLeftMarginStr :=
		plainTextLineZero.GetLeftMarginStr()

	if expectedLeftMarginStr !=
		actualLeftMarginStr {

		t.Errorf("%v Test #1\n"+
			"plainTextLineZero.GetLeftMarginStr()\n"+
			"Error: Expected 'expectedLeftMarginStr' == 'actualLeftMarginStr'\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!!\n"+
			"Expected Text String = '%v'\n"+
			"Instead, Text String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedLeftMarginStr,
			actualLeftMarginStr)

		return
	}

	var plainTextLine01 TextLineSpecPlainText

	plainTextLine01,
		err = TextLineSpecPlainText{}.NewPlainText(
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

	expectedLeftMarginStr = "!!!!"

	expectedLeftMarginRunes =
		[]rune(expectedLeftMarginStr)

	err = plainTextLine01.SetLeftMarginRunes(
		expectedLeftMarginRunes,
		ePrefix.XCtx(
			"plainTextLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualLeftMarginStr =
		plainTextLine01.GetLeftMarginStr()

	if expectedLeftMarginStr !=
		actualLeftMarginStr {

		t.Errorf("%v Test #2\n"+
			"plainTextLine01.GetLeftMarginStr()\n"+
			"Error: Expected 'expectedLeftMarginStr' == 'actualLeftMarginStr'\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!!\n"+
			"Expected Text String = '%v'\n"+
			"Instead, Text String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedLeftMarginStr,
			actualLeftMarginStr)

		return
	}

	err = plainTextLine01.SetLeftMarginRunes(
		expectedLeftMarginRunes,
		StrMech{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from plainTextLine01."+
			"SetLeftMarginChars()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

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

	expectedLeftMarginStr = ""
	expectedLeftMarginRunes = nil

	err = plainTextLine02.SetLeftMarginRunes(
		expectedLeftMarginRunes,
		ePrefix.XCtx(
			"plainTextLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualLeftMarginStr =
		plainTextLine02.GetLeftMarginStr()

	if expectedLeftMarginStr !=
		actualLeftMarginStr {

		t.Errorf("%v Test #3\n"+
			"plainTextLine02.GetLeftMarginStr()\n"+
			"Error: Expected 'expectedLeftMarginStr' == 'actualLeftMarginStr'\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!!\n"+
			"Expected Text String = '%v'\n"+
			"Instead, Text String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedLeftMarginStr,
			actualLeftMarginStr)

		return
	}

	expectedLeftMarginStr = "   "
	expectedLeftMarginRunes =
		[]rune(expectedLeftMarginStr)

	var plainTextLine03 TextLineSpecPlainText

	plainTextLine03,
		err = TextLineSpecPlainText{}.NewPlainText(
		expectedLeftMarginRunes,
		expectedRightMarginRunes,
		expectedTextString,
		expectedNewLineRunes,
		false,
		ePrefix.XCtx("plainTextLine03"))

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

	expectedLeftMarginRunes = nil

	expectedLeftMarginRunes,
		err = strMechPreon{}.ptr().getRepeatRuneChar(
		1000001,
		'X',
		ePrefix.XCtx(
			"Repeat Count = 1000001"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = plainTextLine03.SetLeftMarginRunes(
		expectedLeftMarginRunes,
		ePrefix.XCtx(
			"plainTextLine03"))

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from plainTextLine03."+
			"SetLeftMarginChars()\n"+
			"because 'expectedLeftMarginStr'has 1,000,001 characters.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())
	}

	expectedLeftMarginRunes = nil

	return
}

func TestTextLineSpecPlainText_SetLineTerminationChars_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_SetLineTerminationChars_000100()",
		"")

	expectedLeftMarginStr := "   "
	expectedRightMarginStr := "   "
	expectedNewLineStr := "\n\n"
	expectedTextString := "How now brown cow!"

	plainTextLineZero := TextLineSpecPlainText{}

	err :=
		plainTextLineZero.SetLineTerminationChars(
			expectedNewLineStr,
			ePrefix.XCtx(
				"plainTextLineZero"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualNewLineStr :=
		plainTextLineZero.GetLineTerminationChars()

	sMech := StrMech{}

	printableExpectedNewLineStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedNewLineStr),
			true)

	printableActualNewLineStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(actualNewLineStr),
			true)

	if printableExpectedNewLineStr !=
		printableActualNewLineStr {

		t.Errorf("%v Test #1\n"+
			"plainTextLineZero.GetLineTerminationChars()\n"+
			"Error: Expected 'expectedNewLineStr' == 'actualNewLineStr'\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!!\n"+
			"Expected Line Termination String = '%v'\n"+
			"Instead, Line Termination String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedNewLineStr,
			printableActualNewLineStr)

		return
	}

	var plainTextLine01 TextLineSpecPlainText

	plainTextLine01,
		err = TextLineSpecPlainText{}.NewPlainTextStrings(
		expectedLeftMarginStr,
		expectedRightMarginStr,
		expectedTextString,
		expectedNewLineStr,
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

	expectedNewLineStr = "\n\n\n\n"

	err =
		plainTextLine01.SetLineTerminationChars(
			expectedNewLineStr,
			ePrefix.XCtx(
				"plainTextLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualNewLineStr =
		plainTextLine01.GetLineTerminationChars()

	printableExpectedNewLineStr =
		sMech.ConvertNonPrintableChars(
			[]rune(expectedNewLineStr),
			true)

	printableActualNewLineStr =
		sMech.ConvertNonPrintableChars(
			[]rune(actualNewLineStr),
			true)

	if printableExpectedNewLineStr !=
		printableActualNewLineStr {

		t.Errorf("%v Test #2\n"+
			"plainTextLine01.GetLineTerminationChars()\n"+
			"Error: Expected 'expectedNewLineStr' == 'actualNewLineStr'\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!!\n"+
			"Expected Line Termination String = '%v'\n"+
			"Instead, Line Termination String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedNewLineStr,
			printableActualNewLineStr)

		return
	}

	err =
		plainTextLine01.SetLineTerminationChars(
			expectedNewLineStr,
			StrMech{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from plainTextLine01."+
			"SetLeftMarginChars()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	var plainTextLine02 TextLineSpecPlainText

	plainTextLine02,
		err = TextLineSpecPlainText{}.NewPlainTextStrings(
		expectedLeftMarginStr,
		expectedRightMarginStr,
		expectedTextString,
		expectedNewLineStr,
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

	expectedNewLineStr = ""

	err =
		plainTextLine02.SetLineTerminationChars(
			expectedNewLineStr,
			ePrefix.XCtx(
				"plainTextLine02"))

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from plainTextLine02."+
			"SetLineTerminationChars()\n"+
			"because 'expectedNewLineStr' is an empty string.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	expectedNewLineStr =
		strings.Repeat("X", 1000001)

	err =
		plainTextLine02.SetLineTerminationChars(
			expectedNewLineStr,
			ePrefix.XCtx(
				"plainTextLine02"))

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from plainTextLine03."+
			"SetLineTerminationChars()\n"+
			"because 'expectedNewLineStr' has 1,000,001 characters.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())
	}

	expectedNewLineStr = ""

	return
}

func TestTextLineSpecPlainText_SetLineTerminationRunes_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_SetLineTerminationRunes_000100()",
		"")

	expectedLeftMarginRunes := []rune{' ', ' ', ' '}
	expectedRightMarginRunes := []rune{' ', ' ', ' '}
	expectedNewLineRunes := []rune{'\n', '\n'}
	expectedTextString := "How now brown cow!"

	plainTextLineZero := TextLineSpecPlainText{}

	err := plainTextLineZero.SetLineTerminationRunes(
		expectedNewLineRunes,
		ePrefix.XCtx(
			"plainTextLineZero"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualNewLineRunes :=
		plainTextLineZero.GetLineTerminationRunes()

	sMech := StrMech{}

	printableExpectedNewLineStr :=
		sMech.ConvertNonPrintableChars(
			expectedNewLineRunes,
			true)

	printableActualNewLineStr :=
		sMech.ConvertNonPrintableChars(
			actualNewLineRunes,
			true)

	if printableExpectedNewLineStr !=
		printableActualNewLineStr {

		t.Errorf("%v Test #1\n"+
			"plainTextLineZero.GetLineTerminationRunes()\n"+
			"Error: Expected 'expectedNewLineRunes' == 'actualNewLineRunes'\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!!\n"+
			"Expected Line Termination String = '%v'\n"+
			"Instead, Line Termination String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedNewLineStr,
			printableActualNewLineStr)

		return
	}

	var plainTextLine01 TextLineSpecPlainText

	plainTextLine01,
		err = TextLineSpecPlainText{}.NewPlainText(
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

	expectedNewLineRunes = []rune("!!!!")

	err = plainTextLine01.SetLineTerminationRunes(
		expectedNewLineRunes,
		ePrefix.XCtx(
			"plainTextLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualNewLineRunes =
		plainTextLine01.GetLineTerminationRunes()

	printableExpectedNewLineStr =
		sMech.ConvertNonPrintableChars(
			expectedNewLineRunes,
			true)

	printableActualNewLineStr =
		sMech.ConvertNonPrintableChars(
			actualNewLineRunes,
			true)

	if printableExpectedNewLineStr !=
		printableActualNewLineStr {

		t.Errorf("%v Test #2\n"+
			"plainTextLine01.GetLineTerminationRunes()\n"+
			"Error: Expected 'expectedNewLineRunes' == 'actualNewLineRunes'\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!!\n"+
			"Expected Line Termination String = '%v'\n"+
			"Instead, Line Termination String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedNewLineStr,
			printableActualNewLineStr)

		return
	}

	err = plainTextLine01.SetLineTerminationRunes(
		expectedNewLineRunes,
		StrMech{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Test #3\n"+
			"Expected an error return from plainTextLine01."+
			"SetLineTerminationRunes()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

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

	expectedNewLineRunes = nil

	err = plainTextLine02.SetLineTerminationRunes(
		expectedNewLineRunes,
		ePrefix.XCtx(
			"plainTextLine02"))

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Test #4\n"+
			"Expected an error return from plainTextLine01."+
			"SetLineTerminationRunes()\n"+
			"because 'expectedNewLineRunes' is invalid.\n"+
			"'expectedNewLineRunes' is a zero length array.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	expectedNewLineRunes = []rune("   ")

	var plainTextLine03 TextLineSpecPlainText

	plainTextLine03,
		err = TextLineSpecPlainText{}.NewPlainText(
		expectedLeftMarginRunes,
		expectedRightMarginRunes,
		expectedTextString,
		expectedNewLineRunes,
		false,
		ePrefix.XCtx("plainTextLine03"))

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

	expectedNewLineRunes = nil

	expectedNewLineRunes,
		err = strMechPreon{}.ptr().getRepeatRuneChar(
		1000001,
		'X',
		ePrefix.XCtx(
			"Repeat Count = 1000001"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = plainTextLine03.SetLineTerminationRunes(
		expectedNewLineRunes,
		ePrefix.XCtx(
			"plainTextLine03"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Test #5\n"+
			"Expected an error return from plainTextLine03."+
			"SetLineTerminationRunes()\n"+
			"because 'expectedNewLineRunes' is invalid.\n"+
			"'expectedNewLineRunes' has an array length greater\n"+
			"than 1-million (1,000,000) characters.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	expectedNewLineRunes = nil

	return
}

func TestTextLineSpecPlainText_SetPlainTextSpec_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_SetPlainTextSpec_000100()",
		"")

	plainTextLine01 := TextLineSpecPlainText{}

	leftMarginSpaces := 2
	rightMarginSpaces := 2
	textString := "How now brown cow"

	err :=
		plainTextLine01.SetPlainTextSpec(
			[]rune(strings.Repeat(" ", leftMarginSpaces)),
			[]rune(strings.Repeat(" ", rightMarginSpaces)),
			textString,
			[]rune{'\n'},
			false,
			StrMech{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from plainTextLine01."+
			"SetPlainTextSpec()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

	}

	return
}

func TestTextLineSpecPlainText_SetPlainTextSpecRunes_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_SetPlainTextSpecRunes_000100()",
		"")

	leftMarginSpaces := 2
	rightMarginSpaces := 3
	textString := "How now brown cow"

	leftMargin := strings.Repeat(" ", leftMarginSpaces)

	rightMargin := strings.Repeat(" ", rightMarginSpaces)

	plainTextLine01 := TextLineSpecPlainText{}

	err :=
		plainTextLine01.SetPlainTextSpecRunes(
			[]rune(leftMargin),
			[]rune(rightMargin),
			[]rune(textString),
			[]rune{'\n'},
			false,
			StrMech{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from plainTextLine01."+
			"SetPlainTextSpecRunes()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

	}

	return
}

func TestTextLineSpecPlainText_SetRightMarginChars_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_SetRightMarginChars_000100()",
		"")

	plainTextLine01 := TextLineSpecPlainText{}

	rightMarginStr := strings.Repeat(" ", 2)

	err :=
		plainTextLine01.SetRightMarginChars(
			rightMarginStr,
			ePrefix.XCtx(
				"plainTextLine01<-rightMarginStr"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	rightMarginStr = strings.Repeat(" ", 1000001)

	err =
		plainTextLine01.SetRightMarginChars(
			rightMarginStr,
			ePrefix.XCtx(
				"plainTextLine01<-rightMarginStr-invalid"))

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from plainTextLine01."+
			"SetRightMarginChars()\n"+
			"because 'rightMarginStr' has a string length > 1,000,000.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	rightMarginStr = ""

	err =
		plainTextLine01.SetRightMarginChars(
			rightMarginStr,
			ePrefix.XCtx(
				"plainTextLine01<-rightMarginStr-zero-length-string"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	expectedLeftMarginChars := "   "
	expectedRightMarginChars := "   "
	expectedTextString := "The cow jumped over the moon!"
	expectedNewLineChars := "\n"

	var plainTextLine02 TextLineSpecPlainText

	plainTextLine02,
		err = TextLineSpecPlainText{}.NewPlainTextStrings(
		expectedLeftMarginChars,
		expectedRightMarginChars,
		expectedTextString,
		expectedNewLineChars,
		false,
		ePrefix.XCtx(
			"plainTextLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	rightMarginStr = "XXXXXX"

	err =
		plainTextLine02.SetRightMarginChars(
			rightMarginStr,
			ePrefix.XCtx(
				"plainTextLine02<-rightMarginStr"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualRightMarginStr :=
		plainTextLine02.GetRightMarginStr()

	if actualRightMarginStr !=
		rightMarginStr {

		t.Errorf("%v - ERROR\n"+
			"Expected Right Margin String does NOT match\n"+
			"Actual Right Margin String.\n"+
			"Expected Right Margin String = '%v'\n"+
			"Actual Right Margin String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			rightMarginStr,
			actualRightMarginStr)

		return
	}

	err =
		plainTextLine01.SetRightMarginChars(
			rightMarginStr,
			StrMech{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from plainTextLine01."+
			"SetRightMarginChars()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecPlainText_SetRightMarginRunes_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_SetRightMarginRunes_000100()",
		"")

	plainTextLine01 := TextLineSpecPlainText{}

	rightMarginRunes :=
		[]rune(strings.Repeat(" ", 2))

	err :=
		plainTextLine01.SetRightMarginRunes(
			rightMarginRunes,
			ePrefix.XCtx(
				"plainTextLine01<-rightMarginRunes"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	rightMarginRunes =
		[]rune(strings.Repeat(" ", 1000001))

	err =
		plainTextLine01.SetRightMarginRunes(
			rightMarginRunes,
			ePrefix.XCtx(
				"plainTextLine01<-rightMarginRunes-invalid"))

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from plainTextLine01."+
			"SetRightMarginRunes()\n"+
			"because 'rightMarginRunes' has an array length > 1,000,000.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	rightMarginRunes = nil

	err =
		plainTextLine01.SetRightMarginRunes(
			rightMarginRunes,
			ePrefix.XCtx(
				"plainTextLine01<-rightMarginRunes-zero-length-string"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	expectedLeftMarginChars := "   "
	expectedRightMarginChars := "   "
	expectedTextString := "The cow jumped over the moon!"
	expectedNewLineChars := "\n"

	var plainTextLine02 TextLineSpecPlainText

	plainTextLine02,
		err = TextLineSpecPlainText{}.NewPlainTextStrings(
		expectedLeftMarginChars,
		expectedRightMarginChars,
		expectedTextString,
		expectedNewLineChars,
		false,
		ePrefix.XCtx(
			"plainTextLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	rightMarginRunes = []rune("XXXXXX")

	err =
		plainTextLine02.SetRightMarginRunes(
			rightMarginRunes,
			ePrefix.XCtx(
				"plainTextLine02<-rightMarginRunes"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualRightMarginStr :=
		plainTextLine02.GetRightMarginStr()

	expectedRightMarginStr :=
		string(rightMarginRunes)

	if actualRightMarginStr !=
		expectedRightMarginStr {

		t.Errorf("%v - ERROR\n"+
			"Expected Right Margin String does NOT match\n"+
			"Actual Right Margin String.\n"+
			"Expected Right Margin String = '%v'\n"+
			"Actual Right Margin String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedRightMarginStr,
			actualRightMarginStr)

		return
	}

	err =
		plainTextLine01.SetRightMarginRunes(
			rightMarginRunes,
			StrMech{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from plainTextLine01."+
			"SetRightMarginRunes()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecPlainText_SetTextRunes_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_SetTextRunes_000100()",
		"")

	plainTextLine01 := TextLineSpecPlainText{}

	textRunes :=
		[]rune("How now brown cow!")

	err :=
		plainTextLine01.SetTextRunes(
			textRunes,
			ePrefix.XCtx(
				"plainTextLine01<-textRunes"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	textRunes =
		[]rune(strings.Repeat(" ", 1000001))

	err =
		plainTextLine01.SetTextRunes(
			textRunes,
			ePrefix.XCtx(
				"plainTextLine01<-textRunes-invalid"))

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from plainTextLine01."+
			"SetTextRunes()\n"+
			"because 'textRunes' has an array length > 1,000,000.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	textRunes = nil

	err =
		plainTextLine01.SetTextRunes(
			textRunes,
			ePrefix.XCtx(
				"plainTextLine01<-textRunes-zero-length-string"))

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from plainTextLine01."+
			"SetTextRunes()\n"+
			"because 'textRunes' is a zero length array.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	expectedLeftMarginChars := "   "
	expectedRightMarginChars := "   "
	expectedTextString := "The cow jumped over the moon!"
	expectedNewLineChars := "\n"

	var plainTextLine02 TextLineSpecPlainText

	plainTextLine02,
		err = TextLineSpecPlainText{}.NewPlainTextStrings(
		expectedLeftMarginChars,
		expectedRightMarginChars,
		expectedTextString,
		expectedNewLineChars,
		false,
		ePrefix.XCtx(
			"plainTextLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	textRunes = []rune("XXXXXX")

	err =
		plainTextLine02.SetTextRunes(
			textRunes,
			ePrefix.XCtx(
				"plainTextLine02<-textRunes"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualTextStr :=
		plainTextLine02.GetTextString()

	expectedTextStr :=
		string(textRunes)

	if actualTextStr !=
		expectedTextStr {

		t.Errorf("%v - ERROR\n"+
			"Expected Text String does NOT match\n"+
			"Actual Text String.\n"+
			"Expected Text String = '%v'\n"+
			"Actual Text String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedTextStr,
			actualTextStr)

		return
	}

	err =
		plainTextLine01.SetTextRunes(
			textRunes,
			StrMech{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from plainTextLine01."+
			"SetTextRunes()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecPlainText_SetTextString_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_SetTextString_000100()",
		"")

	plainTextLine01 := TextLineSpecPlainText{}

	textStr := "How now brown cow!"

	err :=
		plainTextLine01.SetTextString(
			textStr,
			ePrefix.XCtx(
				"plainTextLine01<-textStr"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	textStr =
		strings.Repeat(" ", 1000001)

	err =
		plainTextLine01.SetTextString(
			textStr,
			ePrefix.XCtx(
				"plainTextLine01<-textStr-invalid"))

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from plainTextLine01."+
			"SetTextString()\n"+
			"because 'textStr' has an array length > 1,000,000.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	textStr = ""

	err =
		plainTextLine01.SetTextString(
			textStr,
			ePrefix.XCtx(
				"plainTextLine01<-textStr-zero-length-string"))

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from plainTextLine01."+
			"SetTextRunes()\n"+
			"because 'textStr' is an empty string.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	expectedLeftMarginChars := "   "
	expectedRightMarginChars := "   "
	expectedTextString := "The cow jumped over the moon!"
	expectedNewLineChars := "\n"

	var plainTextLine02 TextLineSpecPlainText

	plainTextLine02,
		err = TextLineSpecPlainText{}.NewPlainTextStrings(
		expectedLeftMarginChars,
		expectedRightMarginChars,
		expectedTextString,
		expectedNewLineChars,
		false,
		ePrefix.XCtx(
			"plainTextLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	textStr = "XXXXXX"

	err =
		plainTextLine02.SetTextString(
			textStr,
			ePrefix.XCtx(
				"plainTextLine02<-textStr"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualTextStr :=
		plainTextLine02.GetTextString()

	expectedTextStr := textStr

	if actualTextStr !=
		expectedTextStr {

		t.Errorf("%v - ERROR\n"+
			"Expected Text String does NOT match\n"+
			"Actual Text String.\n"+
			"Expected Text String = '%v'\n"+
			"Actual Text String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedTextStr,
			actualTextStr)

		return
	}

	err =
		plainTextLine01.SetTextString(
			textStr,
			StrMech{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from plainTextLine01."+
			"SetTextString()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
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
