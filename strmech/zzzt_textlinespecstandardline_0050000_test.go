package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"testing"
)

func TestTestTextLineSpecStandardLineNanobot_addTextFields_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTestTextLineSpecStandardLineNanobot_addTextFields_000100()",
		"")

	stdLine02,
		err := createTestTextLineSpecStandardLine01(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var textFields []ITextFieldSpecification

	textFields,
		err = stdLine02.GetTextFields(
		ePrefix.XCtx(
			"textFields<-stdLine02"))

	stdLineNanobot := textLineSpecStandardLineNanobot{}

	_,
		err =
		stdLineNanobot.addTextFields(
			nil,
			&textFields,
			ePrefix.XCtx(
				"txtStdLine is nil pointer"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from stdLineNanobot.addTextFields()\n"+
			"because 'txtStdLine' input parameter is 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!\n",
			ePrefix.XCtxEmpty().String())

		return

	}

	stdLine01 := TextLineSpecStandardLine{}.New()

	_,
		err =
		stdLineNanobot.addTextFields(
			&stdLine01,
			nil,
			ePrefix.XCtx(
				"textFields is nil pointer"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from stdLineNanobot.addTextFields()\n"+
			"because 'textFields' input parameter is 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!\n",
			ePrefix.XCtxEmpty().String())

		return

	}

}

func TestTestTextLineSpecStandardLineNanobot_addTextFields_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTestTextLineSpecStandardLineNanobot_addTextFields_000200()",
		"")

	stdLine02,
		err := createTestTextLineSpecStandardLine01(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var textFields []ITextFieldSpecification

	textFields,
		err = stdLine02.GetTextFields(
		ePrefix.XCtx(
			"textFields<-stdLine02"))

	stdLineNanobot := textLineSpecStandardLineNanobot{}

	textFields[1].Empty()

	stdLine01 := TextLineSpecStandardLine{}.New()

	_,
		err =
		stdLineNanobot.addTextFields(
			&stdLine01,
			&textFields,
			ePrefix.XCtx(
				"textFields[1] is invalid"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from stdLineNanobot.addTextFields()\n"+
			"because 'textFields[1]' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!\n",
			ePrefix.XCtxEmpty().String())

		return

	}

	return
}

func TestTextLineSpecStandardLineNanobot_copyIn_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLineNanobot_copyIn_000100()",
		"")

	stdLine01 := TextLineSpecStandardLine{}.NewPtr()

	rightMarginLen := 5

	rightMarginSpec,
		err := TextFieldSpecSpacer{}.NewSpacer(
		rightMarginLen,
		ePrefix.XCtx(
			"rightMarginSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var leftMarginSpec TextFieldSpecSpacer

	leftMarginLen := 6

	leftMarginSpec,
		err = TextFieldSpecSpacer{}.NewSpacer(
		leftMarginLen,
		ePrefix.XCtx(
			"leftMarginSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	label := "How Now Brown Cow!"
	fieldLen := len(label) + 4
	txtJustify := TxtJustify.Center()

	var labelSpec TextFieldSpecLabel

	labelSpec,
		err = TextFieldSpecLabel{}.NewTextLabel(
		label,
		fieldLen,
		txtJustify,
		ePrefix.XCtx(
			"labelSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	_,
		err = stdLine01.AddTextField(
		&leftMarginSpec,
		ePrefix.XCtx(
			"stdLine01<-leftMarginSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	_,
		err = stdLine01.AddTextField(
		&labelSpec,
		ePrefix.XCtx(
			"stdLine01<-labelSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	_,
		err = stdLine01.AddTextField(
		&rightMarginSpec,
		ePrefix.XCtx(
			"stdLine01<-rightMarginSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = stdLine01.IsValidInstanceError(
		ePrefix.XCtx(
			"stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	stdLine02 := TextLineSpecStandardLine{}

	txtStdLineNanobot := textLineSpecStandardLineNanobot{}

	err =
		txtStdLineNanobot.copyIn(
			&stdLine02,
			stdLine01,
			ePrefix.XCtx(
				"stdLine02<-stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = stdLine02.IsValidInstanceError(
		ePrefix.XCtx(
			"stdLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !stdLine02.Equal(stdLine01) {
		t.Errorf("%v\n"+
			"Error: Expected stdLine02 == stdLine01\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	err =
		txtStdLineNanobot.copyIn(
			&stdLine02,
			nil,
			ePrefix.XCtx(
				"incomingStdLine==nil"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtStdLineNanobot{}."+
			"copyIn()\n"+
			"because 'incomingStdLine' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	err =
		txtStdLineNanobot.copyIn(
			nil,
			stdLine01,
			ePrefix.XCtx(
				"nil<-stdLine01"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtStdLineNanobot{}."+
			"copyIn()\n"+
			"because 'targetStdLine' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	stdLine01.numOfStdLines = -47

	err =
		txtStdLineNanobot.copyIn(
			&stdLine02,
			stdLine01,
			ePrefix.XCtx(
				"stdLine02<-invalid stdLine01"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtStdLineNanobot{}."+
			"copyIn()\n"+
			"because 'targetStdLine' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	stdLine01.numOfStdLines = 1

	stdLine01.newLineChars = nil

	err =
		txtStdLineNanobot.copyIn(
			&stdLine02,
			stdLine01,
			ePrefix.XCtx(
				"stdLine02<-stdLine01 newLineChars==nil"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	stdLine01.newLineChars = []rune{'\n', 0, '\n', 0}

	err =
		txtStdLineNanobot.copyIn(
			&stdLine02,
			stdLine01,
			ePrefix.XCtx(
				"stdLine02<-stdLine01 newLineChars invalid"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtStdLineNanobot{}."+
			"copyIn()\n"+
			"because 'stdLine01.newLineChars' contains invalid characters.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecStandardLineNanobot_copyIn_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLineNanobot_copyIn_000200()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine01(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	stdLine02 := TextLineSpecStandardLine{}.New()

	stdLine01.newLineChars = []rune{'\n', 0, '\n', 0}

	txtStdLineNanobot := textLineSpecStandardLineNanobot{}

	err =
		txtStdLineNanobot.copyIn(
			&stdLine02,
			&stdLine01,
			ePrefix.XCtx(
				"stdLine02<-stdLine01"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtStdLineNanobot{}."+
			"copyIn()\n"+
			"because 'stdLine01.newLineChars' are invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTestTextLineSpecStandardLineNanobot_copyOut_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTestTextLineSpecStandardLineNanobot_copyOut_000100()",
		"")

	stdLine01 := TextLineSpecStandardLine{}.NewPtr()

	rightMarginLen := 5

	rightMarginSpec,
		err := TextFieldSpecSpacer{}.NewSpacer(
		rightMarginLen,
		ePrefix.XCtx(
			"rightMarginSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var leftMarginSpec TextFieldSpecSpacer

	leftMarginLen := 6

	leftMarginSpec,
		err = TextFieldSpecSpacer{}.NewSpacer(
		leftMarginLen,
		ePrefix.XCtx(
			"leftMarginSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	label := "How Now Brown Cow!"
	fieldLen := len(label) + 4
	txtJustify := TxtJustify.Center()

	var labelSpec TextFieldSpecLabel

	labelSpec,
		err = TextFieldSpecLabel{}.NewTextLabel(
		label,
		fieldLen,
		txtJustify,
		ePrefix.XCtx(
			"labelSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	_,
		err = stdLine01.AddTextField(
		&leftMarginSpec,
		ePrefix.XCtx(
			"stdLine01<-leftMarginSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	_,
		err = stdLine01.AddTextField(
		&labelSpec,
		ePrefix.XCtx(
			"stdLine01<-labelSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	_,
		err = stdLine01.AddTextField(
		&rightMarginSpec,
		ePrefix.XCtx(
			"stdLine01<-rightMarginSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = stdLine01.IsValidInstanceError(
		ePrefix.XCtx(
			"stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	stdLine02 := TextLineSpecStandardLine{}

	txtStdLineNanobot := textLineSpecStandardLineNanobot{}

	stdLine02,
		err =
		txtStdLineNanobot.copyOut(
			stdLine01,
			ePrefix.XCtx(
				"stdLine02<-stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = stdLine02.IsValidInstanceError(
		ePrefix.XCtx(
			"stdLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !stdLine02.Equal(stdLine01) {
		t.Errorf("%v\n"+
			"Error: Expected stdLine02 == stdLine01\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	_,
		err =
		txtStdLineNanobot.copyOut(
			nil,
			ePrefix.XCtx(
				"txtStdLine==nil"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtStdLineNanobot{}."+
			"copyOut()\n"+
			"because 'txtStdLine' is 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	stdLine01.numOfStdLines = -47
	_,
		err =
		txtStdLineNanobot.copyOut(
			stdLine01,
			ePrefix.XCtx(
				"stdLine01 is invalid"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtStdLineNanobot{}."+
			"copyOut()\n"+
			"because 'stdLine01.numOfStdLines' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	stdLine01.numOfStdLines = 1
	stdLine01.newLineChars = nil
	// nil newLinChars defaults to '\n'

	_,
		err =
		txtStdLineNanobot.copyOut(
			stdLine01,
			ePrefix.XCtx(
				"stdLine01.newLineChars==nil"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	stdLine01.newLineChars = []rune{'\n', 0, '\n', 0}
	_,
		err =
		txtStdLineNanobot.copyOut(
			stdLine01,
			ePrefix.XCtx(
				"stdLine01 is invalid"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtStdLineNanobot{}."+
			"copyOut()\n"+
			"because 'stdLine01.newLineChars' contains invalid characters.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

}

func TestTextLineSpecStandardLineNanobot_setTxtSpecStandardLine_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTESTSERIES_FUNCNAME_000100()",
		"")

	numOfStdLines := 1

	newLineChars := []rune{'\n'}

	stdLine02,
		err := createTestTextLineSpecStandardLine01(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var textFields []ITextFieldSpecification

	textFields,
		err = stdLine02.GetTextFields(
		ePrefix.XCtx(
			"textFields<-stdLine02"))

	stdLineNanobot := textLineSpecStandardLineNanobot{}

	err = stdLineNanobot.setTxtSpecStandardLine(
		nil,
		numOfStdLines,
		textFields,
		newLineChars,
		false,
		ePrefix.XCtx(
			"txtStdLine==nil"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from stdLineNanobot.setTxtSpecStandardLine()\n"+
			"because 'txtStdLine' is  'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	stdLine01 := TextLineSpecStandardLine{}.New()

	err = stdLineNanobot.setTxtSpecStandardLine(
		&stdLine01,
		numOfStdLines,
		nil,
		newLineChars,
		false,
		ePrefix.XCtx(
			"textFields[]==nil"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from stdLineNanobot.setTxtSpecStandardLine()\n"+
			"because 'textFields' is  'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	numOfStdLines = -1

	err = stdLineNanobot.setTxtSpecStandardLine(
		&stdLine01,
		numOfStdLines,
		textFields,
		newLineChars,
		false,
		ePrefix.XCtx(
			"textFields[]==nil"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from stdLineNanobot.setTxtSpecStandardLine()\n"+
			"because 'numOfStdLines' is  '-1'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	numOfStdLines = 1000001

	err = stdLineNanobot.setTxtSpecStandardLine(
		&stdLine01,
		numOfStdLines,
		textFields,
		newLineChars,
		false,
		ePrefix.XCtx(
			"textFields[]==nil"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from stdLineNanobot.setTxtSpecStandardLine()\n"+
			"because 'numOfStdLines' is  '1000001'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	numOfStdLines = 1
	newLineChars = nil

	err = stdLineNanobot.setTxtSpecStandardLine(
		&stdLine01,
		numOfStdLines,
		textFields,
		newLineChars,
		false,
		ePrefix.XCtx(
			"textFields[]==nil"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from stdLineNanobot.setTxtSpecStandardLine()\n"+
			"because 'newLineChars' is  'nill'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	numOfStdLines = 1
	newLineChars = []rune{'\n', 0, '\n', 0}

	err = stdLineNanobot.setTxtSpecStandardLine(
		&stdLine01,
		numOfStdLines,
		textFields,
		newLineChars,
		false,
		ePrefix.XCtx(
			"textFields[]==nil"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from stdLineNanobot.setTxtSpecStandardLine()\n"+
			"because 'newLineChars' is  invalid'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

}

func TestTextLineSpecStandardLineElectron_deleteTextField_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLineElectron_deleteTextField_000100()",
		"")

	indexId := 2

	txtStdLineElectron := textLineSpecStandardLineElectron{}

	err := txtStdLineElectron.deleteTextField(
		nil,
		indexId,
		&ePrefix)

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error from txtStdLineElectron.deleteTextField()\n"+
			"because txtStdLine is a 'nil' pointer!\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	stdLine01 := TextLineSpecStandardLine{}.NewPtr()

	rightMarginLen := 5

	rightMarginSpec,
		err := TextFieldSpecSpacer{}.NewSpacer(
		rightMarginLen,
		ePrefix.XCtx(
			"rightMarginSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var leftMarginSpec TextFieldSpecSpacer

	leftMarginLen := 6

	leftMarginSpec,
		err = TextFieldSpecSpacer{}.NewSpacer(
		leftMarginLen,
		ePrefix.XCtx(
			"leftMarginSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	label := "How Now Brown Cow!"
	fieldLen := len(label) + 4
	txtJustify := TxtJustify.Center()

	var labelSpec TextFieldSpecLabel

	labelSpec,
		err = TextFieldSpecLabel{}.NewTextLabel(
		label,
		fieldLen,
		txtJustify,
		ePrefix.XCtx(
			"labelSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	_,
		err = stdLine01.AddTextField(
		&leftMarginSpec,
		ePrefix.XCtx(
			"stdLine01<-leftMarginSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	_,
		err = stdLine01.AddTextField(
		&labelSpec,
		ePrefix.XCtx(
			"stdLine01<-labelSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	_,
		err = stdLine01.AddTextField(
		&rightMarginSpec,
		ePrefix.XCtx(
			"stdLine01<-rightMarginSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = stdLine01.IsValidInstanceError(
		ePrefix.XCtx(
			"stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtStdLineElectron02 := textLineSpecStandardLineElectron{}

	indexId = 18

	err = txtStdLineElectron02.deleteTextField(
		stdLine01,
		indexId,
		&ePrefix)

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error from txtStdLineElectron02.deleteTextField()\n"+
			"because indexId is invalid!\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	indexId = -2

	err = txtStdLineElectron02.deleteTextField(
		stdLine01,
		indexId,
		&ePrefix)

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error from txtStdLineElectron02.deleteTextField()\n"+
			"because indexId is less than zero!\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	indexId = 0

	stdLine01.textFields = nil

	err = txtStdLineElectron02.deleteTextField(
		stdLine01,
		indexId,
		&ePrefix)

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error from txtStdLineElectron02.deleteTextField()\n"+
			"because stdLine01.textFields = nil!\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

}

func TestTextLineSpecStandardLineElectron_deleteTextField_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLineElectron_deleteTextField_000200()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine02(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err =
		textLineSpecStandardLineElectron{}.ptr().deleteTextField(
			&stdLine01,
			1,
			ePrefix.XCtx(
				"stdLine01 delete index 1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	return
}

func TestTextLineSpecStandardLineElectron_emptyStandardLine_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLineElectron_emptyStandardLine_000100()",
		"")

	txtStdLineElectron := textLineSpecStandardLineElectron{}

	err := txtStdLineElectron.emptyStandardLine(
		nil,
		ePrefix.XCtx(
			"txtStdLine is 'nil'"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from txtStdLineElectron.emptyStandardLine()\n"+
			"because 'txtStdLine' is a nil pointer.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

}

func TestTextLineSpecStandardLineElectron_testValidityOfTextFields_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLineElectron_testValidityOfTextFields_000100()",
		"")

	txtStdLineElectron := textLineSpecStandardLineElectron{}

	_,
		err := txtStdLineElectron.testValidityOfTextFields(
		nil,
		ePrefix.XCtx(
			"txtFields is empty"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from txtStdLineElectron.emptyStandardLine()\n"+
			"because 'txtFields' is  'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

}

func TestTextLineSpecStandardLineElectron_testValidityOfTextFields_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLineElectron_testValidityOfTextFields_000200()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine01(
		ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var textFields []ITextFieldSpecification

	textFields,
		err =
		stdLine01.GetTextFields(
			ePrefix.XCtx(
				"textFields<-stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	txtStdLineElectron := textLineSpecStandardLineElectron{}

	textFields[1] = nil

	_,
		err = txtStdLineElectron.testValidityOfTextFields(
		textFields,
		ePrefix.XCtx(
			"textFields[1] = nil"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from txtStdLineElectron.emptyStandardLine()\n"+
			"because 'txtFields[1]' is  'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	textFields,
		err =
		stdLine01.GetTextFields(
			ePrefix.XCtx(
				"#2 textFields<-stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	textFields[1].Empty()

	_,
		err = txtStdLineElectron.testValidityOfTextFields(
		textFields,
		ePrefix.XCtx(
			"txtFields[1] is invalid"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from txtStdLineElectron.emptyStandardLine()\n"+
			"because 'txtFields[1]' is empty and invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

}
