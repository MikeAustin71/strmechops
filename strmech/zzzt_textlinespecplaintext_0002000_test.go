package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
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

func TestTextLineSpecPlainText_setPlainTextSpec_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_setPlainTextSpec_000100()",
		"")

	leftMarginSpaces := 2
	rightMarginSpaces := 3
	textString := "How now brown cow"

	leftMargin := strings.Repeat(" ", leftMarginSpaces)

	rightMargin := strings.Repeat(" ", rightMarginSpaces)

	plainTextLine01 := TextLineSpecPlainText{}

	txtLinePlainTextAtom := textLineSpecPlainTextAtom{}

	err :=
		txtLinePlainTextAtom.setPlainTextSpec(
			&plainTextLine01,
			[]rune(leftMargin),
			[]rune(rightMargin),
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

	err =
		txtLinePlainTextAtom.setPlainTextSpec(
			nil,
			[]rune(leftMargin),
			[]rune(rightMargin),
			textString,
			[]rune{'\n'},
			false,
			ePrefix.XCtx(
				"plainTextLine01"))

	if err == nil {
		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from txtLinePlainTextAtom."+
			"setPlainTextSpec()\n"+
			"because 'plainTxtLine' is nil.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	err =
		txtLinePlainTextAtom.setPlainTextSpec(
			&plainTextLine01,
			[]rune(leftMargin),
			[]rune(rightMargin),
			"",
			[]rune{'\n'},
			false,
			ePrefix.XCtx(
				"plainTextLine01"))

	if err == nil {
		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from txtLinePlainTextAtom."+
			"setPlainTextSpec()\n"+
			"because 'textString' is an empty string.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	textString =
		strings.Repeat("X", 1000001)

	err =
		txtLinePlainTextAtom.setPlainTextSpec(
			&plainTextLine01,
			[]rune(leftMargin),
			[]rune(rightMargin),
			textString,
			[]rune{'\n'},
			false,
			ePrefix.XCtx(
				"plainTextLine01"))

	if err == nil {
		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from txtLinePlainTextAtom."+
			"setPlainTextSpec()\n"+
			"because 'textString' has a charcter length of 1,000,001.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	textString = "How now brown cow!"

	var marginRunes []rune

	marginRunes,
		err =
		strMechPreon{}.ptr().getRepeatRuneChar(
			1000001,
			'X',
			ePrefix.XCtx(
				"->marginRunes"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err =
		txtLinePlainTextAtom.setPlainTextSpec(
			&plainTextLine01,
			marginRunes,
			[]rune(rightMargin),
			textString,
			[]rune{'\n'},
			false,
			ePrefix.XCtx(
				"plainTextLine01"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from txtLinePlainTextAtom."+
			"setPlainTextSpec()\n"+
			"because 'leftMarginRunes' has an array length of 1,000,001.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	err =
		txtLinePlainTextAtom.setPlainTextSpec(
			&plainTextLine01,
			[]rune(leftMargin),
			marginRunes,
			textString,
			[]rune{'\n'},
			false,
			ePrefix.XCtx(
				"plainTextLine01"))

	if err == nil {
		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from txtLinePlainTextAtom."+
			"setPlainTextSpec()\n"+
			"because 'rightMarginRunes' has an array length of 1,000,001.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	err =
		txtLinePlainTextAtom.setPlainTextSpec(
			&plainTextLine01,
			[]rune(leftMargin),
			[]rune(rightMargin),
			textString,
			marginRunes,
			false,
			ePrefix.XCtx(
				"plainTextLine01"))

	if err == nil {
		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from txtLinePlainTextAtom."+
			"setPlainTextSpec()\n"+
			"because 'newLineChars' has an array length of 1,000,001.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	marginRunes = nil

	err =
		txtLinePlainTextAtom.setPlainTextSpec(
			&plainTextLine01,
			nil,
			[]rune(rightMargin),
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

	err =
		txtLinePlainTextAtom.setPlainTextSpec(
			&plainTextLine01,
			[]rune(leftMargin),
			nil,
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

	err =
		txtLinePlainTextAtom.setPlainTextSpec(
			&plainTextLine01,
			[]rune(leftMargin),
			[]rune(rightMargin),
			textString,
			nil,
			false,
			ePrefix.XCtx(
				"plainTextLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	return
}

func TestTextLineSpecPlainText_testValidityOfTextLineSpecPlainText_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_testValidityOfTextLineSpecPlainText_000100()",
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

	txtLinePlainTextAtom := textLineSpecPlainTextAtom{}

	var isValid bool

	isValid,
		err =
		txtLinePlainTextAtom.testValidityOfTextLineSpecPlainText(
			plainTextLine01,
			ePrefix.XCtx(
				"plainTextLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !isValid {
		t.Errorf("%v - ERROR\n"+
			"txtLinePlainTextAtom.testValidityOfTextLineSpecPlainText()\n"+
			"Expected a return 'isValid' value of 'true'.\n"+
			"HOWEVER, 'isValid' == 'false'!\n",
			ePrefix.XCtxEmpty().String())
	}

	isValid,
		err =
		txtLinePlainTextAtom.testValidityOfTextLineSpecPlainText(
			nil,
			ePrefix.XCtx(
				"plainTextLine01"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinePlainTextAtom{}."+
			"testValidityOfTextLineSpecPlainText()\n"+
			"because 'plainTextLine' is 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	if isValid {
		t.Errorf("%v - ERROR\n"+
			"'plainTextLine' is 'nil' test\n"+
			"txtLinePlainTextAtom.testValidityOfTextLineSpecPlainText()\n"+
			"Expected a return 'isValid' value of 'false'.\n"+
			"HOWEVER, 'isValid' == 'true'!\n",
			ePrefix.XCtxEmpty().String())
	}

	plainTextLine01.newLineChars,
		err =
		strMechPreon{}.ptr().getRepeatRuneChar(
			1000001,
			'X',
			ePrefix.XCtx(
				"leftMarginRunes"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	isValid,
		err =
		txtLinePlainTextAtom.testValidityOfTextLineSpecPlainText(
			plainTextLine01,
			ePrefix.XCtx(
				"plainTextLine01"))

	plainTextLine01.newLineChars = nil

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinePlainTextAtom{}."+
			"testValidityOfTextLineSpecPlainText()\n"+
			"because 'plainTextLine01.newLineChars' has an array length\n"+
			"of 1,000,001.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	if isValid {
		t.Errorf("%v - ERROR\n"+
			"Invalid 'plainTextLine01.newLineChars' test\n"+
			"txtLinePlainTextAtom.testValidityOfTextLineSpecPlainText()\n"+
			"Expected a return 'isValid' value of 'false'.\n"+
			"HOWEVER, 'isValid' == 'true'!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecPlainText_getFormattedText_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_testValidityOfTextLineSpecPlainText_000100()",
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

	txtLinePlainTextNanobot := textLineSpecPlainTextNanobot{}

	var formattedStr string

	formattedStr,
		err =
		txtLinePlainTextNanobot.getFormattedText(
			&plainTextLine01,
			ePrefix.XCtx(
				"plainTextLine01-formattedStr"))

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
			[]rune(formattedStr),
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

	formattedStr,
		err =
		txtLinePlainTextNanobot.getFormattedText(
			nil,
			ePrefix.XCtx(
				"plainTxtLine is 'nil'"))

	if err == nil {

		t.Errorf("%v\n"+
			"Error:\n"+
			"Expected error return from plainTextLine01.TextBuilder()\n"+
			"because 'plainTextLine01' is empty.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

}

func TestTextLineSpecPlainText_setPlainTextSpecRunes_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_setPlainTextSpecRunes_000100()",
		"")

	leftMarginSpaces := 2
	rightMarginSpaces := 3

	textString := "How now brown cow"

	leftMargin := strings.Repeat(" ", leftMarginSpaces)

	rightMargin := strings.Repeat(" ", rightMarginSpaces)

	plainTextLine01 := TextLineSpecPlainText{}

	txtLinePlainTextNanobot := textLineSpecPlainTextNanobot{}

	err :=
		txtLinePlainTextNanobot.setPlainTextSpecRunes(
			&plainTextLine01,
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

	var marginRunes []rune

	marginRunes,
		err =
		strMechPreon{}.ptr().getRepeatRuneChar(
			1000001,
			'X',
			ePrefix.XCtx(
				"->marginRunes"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err =
		txtLinePlainTextNanobot.setPlainTextSpecRunes(
			&plainTextLine01,
			marginRunes,
			[]rune(rightMargin),
			[]rune(textString),
			[]rune{'\n'},
			false,
			ePrefix.XCtx(
				"plainTextLine01-leftMargin Error"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from txtLinePlainTextNanobot."+
			"setPlainTextSpecRunes()\n"+
			"because 'leftMarginRunes' has an array length of 1,000,001.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	err =
		txtLinePlainTextNanobot.setPlainTextSpecRunes(
			&plainTextLine01,
			[]rune(leftMargin),
			marginRunes,
			[]rune(textString),
			[]rune{'\n'},
			false,
			ePrefix.XCtx(
				"plainTextLine01-right margin Error"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from txtLinePlainTextNanobot."+
			"setPlainTextSpecRunes()\n"+
			"because 'rightMarginRunes' has an array length of 1,000,001.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	err =
		txtLinePlainTextNanobot.setPlainTextSpecRunes(
			&plainTextLine01,
			[]rune(leftMargin),
			[]rune(rightMargin),
			[]rune(textString),
			marginRunes,
			false,
			ePrefix.XCtx(
				"plainTextLine01 New Line Chars Error"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from txtLinePlainTextNanobot."+
			"setPlainTextSpecRunes()\n"+
			"because 'newLineChars' has an array length of 1,000,001.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	err =
		txtLinePlainTextNanobot.setPlainTextSpecRunes(
			&plainTextLine01,
			[]rune(leftMargin),
			[]rune(rightMargin),
			marginRunes,
			[]rune{'\n'},
			false,
			ePrefix.XCtx(
				"plainTextLine01-Text Runes Error"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from txtLinePlainTextNanobot."+
			"setPlainTextSpecRunes()\n"+
			"because 'textRunes' has an array length of 1,000,001.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	marginRunes = nil

	err =
		txtLinePlainTextNanobot.setPlainTextSpecRunes(
			&plainTextLine01,
			[]rune(leftMargin),
			[]rune(rightMargin),
			nil,
			[]rune{'\n'},
			false,
			ePrefix.XCtx(
				"plainTextLine01-Text Runes Error"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from txtLinePlainTextNanobot."+
			"setPlainTextSpecRunes()\n"+
			"because 'textRunes' has a value of 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

}

func TestTextLineSpecPlainText_setPlainTextSpecStrings_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecPlainText_setPlainTextSpecStrings_000100()",
		"")

	expectedLeftMarginChars := "   "
	expectedRightMarginChars := "   "
	expectedNewLineChars := "\n\n"
	expectedTextString := "How now brown cow!"

	plainTextLine01 := TextLineSpecPlainText{}

	txtLinePlainTextNanobot := textLineSpecPlainTextNanobot{}

	err :=
		txtLinePlainTextNanobot.setPlainTextSpecStrings(
			&plainTextLine01,
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

	err =
		txtLinePlainTextNanobot.setPlainTextSpecStrings(
			nil,
			expectedLeftMarginChars,
			expectedRightMarginChars,
			expectedTextString,
			expectedNewLineChars,
			false,
			ePrefix.XCtx(
				"plainTextLine01"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from txtLinePlainTextNanobot."+
			"setPlainTextSpecStrings()\n"+
			"because 'plainTextLine01' has a valueof 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	invalidCharLen :=
		strings.Repeat("X", 1000001)

	err =
		txtLinePlainTextNanobot.setPlainTextSpecStrings(
			&plainTextLine01,
			invalidCharLen,
			expectedRightMarginChars,
			expectedTextString,
			expectedNewLineChars,
			false,
			ePrefix.XCtx(
				"plainTextLine01->Left Margin Error"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from txtLinePlainTextNanobot."+
			"setPlainTextSpecStrings()\n"+
			"because 'leftMarginChars' has a length of 1,000,001 characters.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	err =
		txtLinePlainTextNanobot.setPlainTextSpecStrings(
			&plainTextLine01,
			expectedLeftMarginChars,
			invalidCharLen,
			expectedTextString,
			expectedNewLineChars,
			false,
			ePrefix.XCtx(
				"plainTextLine01->Left Margin Error"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from txtLinePlainTextNanobot."+
			"setPlainTextSpecStrings()\n"+
			"because 'rightMarginChars' has a length of 1,000,001 characters.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	err =
		txtLinePlainTextNanobot.setPlainTextSpecStrings(
			&plainTextLine01,
			expectedLeftMarginChars,
			expectedRightMarginChars,
			expectedTextString,
			invalidCharLen,
			false,
			ePrefix.XCtx(
				"plainTextLine01->Left Margin Error"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from txtLinePlainTextNanobot."+
			"setPlainTextSpecStrings()\n"+
			"because 'newLineChars' has a length of 1,000,001 characters.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	err =
		txtLinePlainTextNanobot.setPlainTextSpecStrings(
			&plainTextLine01,
			expectedLeftMarginChars,
			expectedRightMarginChars,
			invalidCharLen,
			expectedNewLineChars,
			false,
			ePrefix.XCtx(
				"plainTextLine01->Left Margin Error"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from txtLinePlainTextNanobot."+
			"setPlainTextSpecStrings()\n"+
			"because 'textString' has a length of 1,000,001 characters.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	invalidCharLen = ""

	err =
		txtLinePlainTextNanobot.setPlainTextSpecStrings(
			&plainTextLine01,
			expectedLeftMarginChars,
			expectedRightMarginChars,
			"",
			expectedNewLineChars,
			false,
			ePrefix.XCtx(
				"plainTextLine01->Left Margin Error"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from txtLinePlainTextNanobot."+
			"setPlainTextSpecStrings()\n"+
			"because 'textString' an empty string.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	plainTextLine02 := TextLineSpecPlainText{}

	err =
		txtLinePlainTextNanobot.setPlainTextSpecStrings(
			&plainTextLine02,
			"",
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

	err =
		txtLinePlainTextNanobot.setPlainTextSpecStrings(
			&plainTextLine02,
			expectedLeftMarginChars,
			"",
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

	err =
		txtLinePlainTextNanobot.setPlainTextSpecStrings(
			&plainTextLine02,
			expectedLeftMarginChars,
			expectedRightMarginChars,
			expectedTextString,
			"",
			false,
			ePrefix.XCtx(
				"plainTextLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

}
