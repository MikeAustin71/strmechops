package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"testing"
)

// createTestTextLineSpecStandardLine01 - Generates a
// TextLineSpecStandardLine object used for test operations.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  errPrefDto          *ePref.ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods listed
//       as a function chain.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       Type ErrPrefixDto is included in the 'errpref' software
//       package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  TextLineSpecStandardLine
//     - If this method completes successfully, it will create and
//       return a new, populated concrete instance of
//       TextLineSpecStandardLine which is fully configured with all
//       the parameters necessary to format one or more standard lines
//       of text for text display, file output or printing.
//
//
//  error
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func createTestTextLineSpecStandardLine01(
	errorPrefix interface{}) (
	TextLineSpecStandardLine,
	error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error
	stdLine01 := TextLineSpecStandardLine{}.New()

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TestDataGeneration - "+
			"createTestTextLineSpecStandardLine01()",
		"")

	if err != nil {
		return stdLine01, err
	}

	rightMarginLen := 5
	var rightMarginSpec TextFieldSpecSpacer

	rightMarginSpec,
		err = TextFieldSpecSpacer{}.NewSpacer(
		rightMarginLen,
		ePrefix.XCtx(
			"rightMarginSpec"))

	if err != nil {
		return stdLine01, err
	}

	var leftMarginSpec TextFieldSpecSpacer

	leftMarginLen := 6

	leftMarginSpec,
		err = TextFieldSpecSpacer{}.NewSpacer(
		leftMarginLen,
		ePrefix.XCtx(
			"leftMarginSpec"))

	if err != nil {
		return stdLine01, err
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
		return stdLine01, err
	}

	_,
		err = stdLine01.AddTextField(
		&leftMarginSpec,
		ePrefix.XCtx(
			"stdLine01<-leftMarginSpec"))

	if err != nil {
		return stdLine01, err
	}

	_,
		err = stdLine01.AddTextField(
		&labelSpec,
		ePrefix.XCtx(
			"stdLine01<-labelSpec"))

	if err != nil {
		return stdLine01, err
	}

	_,
		err = stdLine01.AddTextField(
		&rightMarginSpec,
		ePrefix.XCtx(
			"stdLine01<-rightMarginSpec"))

	if err != nil {
		return stdLine01, err
	}

	err = stdLine01.IsValidInstanceError(
		ePrefix.XCtx(
			"stdLine01"))

	if err != nil {
		return stdLine01, err
	}

	return stdLine01, err
}

// createTestTextLineSpecStandardLine02 - Generates a
// TextLineSpecStandardLine object used for test operations.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  errPrefDto          *ePref.ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods listed
//       as a function chain.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       Type ErrPrefixDto is included in the 'errpref' software
//       package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  TextLineSpecStandardLine
//     - If this method completes successfully, it will create and
//       return a new, populated concrete instance of
//       TextLineSpecStandardLine which is fully configured with all
//       the parameters necessary to format one or more standard lines
//       of text for text display, file output or printing.
//
//
//  error
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func createTestTextLineSpecStandardLine02(
	errorPrefix interface{}) (
	TextLineSpecStandardLine,
	error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error
	stdLine02 := TextLineSpecStandardLine{}.New()

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TestDataGeneration - "+
			"createTestTextLineSpecStandardLine02()",
		"")

	if err != nil {
		return stdLine02, err
	}

	rightMarginLen := 2
	var rightMarginSpec TextFieldSpecSpacer

	rightMarginSpec,
		err = TextFieldSpecSpacer{}.NewSpacer(
		rightMarginLen,
		ePrefix.XCtx(
			"rightMarginSpec"))

	if err != nil {
		return stdLine02, err
	}

	var leftMarginSpec TextFieldSpecSpacer

	leftMarginLen := 3

	leftMarginSpec,
		err = TextFieldSpecSpacer{}.NewSpacer(
		leftMarginLen,
		ePrefix.XCtx(
			"leftMarginSpec"))

	if err != nil {
		return stdLine02, err
	}

	label := "Good Morning Sunshine!"
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
		return stdLine02, err
	}

	_,
		err = stdLine02.AddTextField(
		&leftMarginSpec,
		ePrefix.XCtx(
			"stdLine02<-leftMarginSpec"))

	if err != nil {
		return stdLine02, err
	}

	_,
		err = stdLine02.AddTextField(
		&labelSpec,
		ePrefix.XCtx(
			"stdLine02<-labelSpec"))

	if err != nil {
		return stdLine02, err
	}

	_,
		err = stdLine02.AddTextField(
		&rightMarginSpec,
		ePrefix.XCtx(
			"stdLine02<-rightMarginSpec"))

	if err != nil {
		return stdLine02, err
	}

	err = stdLine02.IsValidInstanceError(
		ePrefix.XCtx(
			"stdLine02"))

	return stdLine02, err
}

// createTestTextLineSpecStandardLine02 - Generates a
// TextLineSpecStandardLine object used for test operations.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  errPrefDto          *ePref.ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods listed
//       as a function chain.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       Type ErrPrefixDto is included in the 'errpref' software
//       package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  TextLineSpecStandardLine
//     - If this method completes successfully, it will create and
//       return a new, populated concrete instance of
//       TextLineSpecStandardLine which is fully configured with all
//       the parameters necessary to format one or more standard lines
//       of text for text display, file output or printing.
//
//
//  error
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func createTestTextLineSpecStandardLine03(
	leftMarginSpaces int,
	labelTxt string,
	txtFieldLen int,
	txtJustify TextJustify,
	rightMarginSpaces int,
	errorPrefix interface{}) (
	TextLineSpecStandardLine,
	error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error
	stdLine03 := TextLineSpecStandardLine{}.New()

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TestDataGeneration - "+
			"createTestTextLineSpecStandardLine03()",
		"")

	if err != nil {
		return stdLine03, err
	}

	var rightMarginSpec TextFieldSpecSpacer

	rightMarginSpec,
		err = TextFieldSpecSpacer{}.NewSpacer(
		rightMarginSpaces,
		ePrefix.XCtx(
			"rightMarginSpec"))

	if err != nil {
		return stdLine03, err
	}

	var leftMarginSpec TextFieldSpecSpacer

	leftMarginSpec,
		err = TextFieldSpecSpacer{}.NewSpacer(
		leftMarginSpaces,
		ePrefix.XCtx(
			"leftMarginSpec"))

	if err != nil {
		return stdLine03, err
	}

	var labelSpec TextFieldSpecLabel

	labelSpec,
		err = TextFieldSpecLabel{}.NewTextLabel(
		labelTxt,
		txtFieldLen,
		txtJustify,
		ePrefix.XCtx(
			"labelSpec"))

	if err != nil {
		return stdLine03, err
	}

	_,
		err = stdLine03.AddTextField(
		&leftMarginSpec,
		ePrefix.XCtx(
			"stdLine03<-leftMarginSpec"))

	if err != nil {
		return stdLine03, err
	}

	_,
		err = stdLine03.AddTextField(
		&labelSpec,
		ePrefix.XCtx(
			"stdLine03<-labelSpec"))

	if err != nil {
		return stdLine03, err
	}

	_,
		err = stdLine03.AddTextField(
		&rightMarginSpec,
		ePrefix.XCtx(
			"stdLine03<-rightMarginSpec"))

	if err != nil {
		return stdLine03, err
	}

	err = stdLine03.IsValidInstanceError(
		ePrefix.XCtx(
			"stdLine03"))

	return stdLine03, err
}

func TestTextLineSpecStandardLine_AddTextFields_000010(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_AddTextFields_000010()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine01(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var stdLine02 TextLineSpecStandardLine

	stdLine02,
		err = createTestTextLineSpecStandardLine02(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	lastIndexId := -1

	expectedLastIndexId :=
		stdLine01.GetNumOfTextFields() - 1

	stdLine03 := TextLineSpecStandardLine{}.New()

	lastIndexId,
		err = stdLine03.AddStandardLine(
		&stdLine01,
		ePrefix.XCtx(
			"stdLine03<-stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if lastIndexId != expectedLastIndexId {

		t.Errorf("%v - ERROR\n"+
			"stdLine03<-stdLine01\n"+
			"Expected Last Index Id = '%v'\n"+
			"  Actual Last Index Id = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedLastIndexId,
			lastIndexId)

		return
	}

	lastIndexId,
		err = stdLine03.AddStandardLine(
		&stdLine02,
		ePrefix.XCtx(
			"stdLine03<-stdLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	expectedLastIndexId =
		len(stdLine01.textFields) +
			len(stdLine02.textFields) - 1

	if lastIndexId != expectedLastIndexId {

		t.Errorf("%v - ERROR\n"+
			"stdLine03<-stdLine02\n"+
			"Expected Last Index Id = '%v'\n"+
			"  Actual Last Index Id = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedLastIndexId,
			lastIndexId)

		return
	}

	stdLine04 := TextLineSpecStandardLine{}.New()

	for idx, val := range stdLine01.textFields {

		var newTextField ITextFieldSpecification

		newTextField,
			err = val.CopyOutITextField(
			ePrefix.XCtx(fmt.Sprintf(
				"stdLine01.textFields[%v]->newTextField",
				idx)))

		if err != nil {
			t.Errorf("%v - ERROR\n"+
				"stdLine01.CopyOutITextField() Failed!\n",
				ePrefix.XCtxEmpty().String())

			return
		}

		stdLine04.textFields =
			append(stdLine04.textFields,
				newTextField)

	}

	for idx, val := range stdLine02.textFields {

		var newTextField ITextFieldSpecification

		newTextField,
			err = val.CopyOutITextField(
			ePrefix.XCtx(fmt.Sprintf(
				"stdLine01.textFields[%v]->newTextField",
				idx)))

		if err != nil {
			t.Errorf("%v - ERROR\n"+
				"stdLine02.CopyOutITextField() Failed!\n",
				ePrefix.XCtxEmpty().String())

			return
		}

		stdLine04.textFields =
			append(stdLine04.textFields,
				newTextField)

	}

	areEqual := stdLine03.Equal(&stdLine04)

	if !areEqual {

		t.Errorf("%v - ERROR\n"+
			"Expected stdLine03==stdLine04\n"+
			"HOWEVER THEY ARE NOT EQUAL!\n",
			ePrefix.XCtxEmpty().String())

		return

	}

}

func TestTextLineSpecStandardLine_AddStandardLine_000020(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_AddStandardLine_000020()",
		"")

	labelTxt := "I deny everything!"

	stdLine01,
		err := createTestTextLineSpecStandardLine03(
		3,
		labelTxt,
		len(labelTxt)+4,
		TxtJustify.Center(),
		5,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var stdLine02 TextLineSpecStandardLine

	labelTxt = "They went that way!"

	stdLine02,
		err = createTestTextLineSpecStandardLine03(
		6,
		labelTxt,
		len(labelTxt)+4,
		TxtJustify.Center(),
		5,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	lastIndexId := -1

	expectedLastIndexId :=
		stdLine01.GetNumOfTextFields() - 1

	stdLine03 := TextLineSpecStandardLine{}.New()

	lastIndexId,
		err = stdLine03.AddStandardLine(
		&stdLine01,
		ePrefix.XCtx(
			"stdLine03<-stdLine01"))

	if lastIndexId != expectedLastIndexId {

		t.Errorf("%v - ERROR\n"+
			"stdLine03<-stdLine01\n"+
			"Expected Last Index Id = '%v'\n"+
			"  Actual Last Index Id = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedLastIndexId,
			lastIndexId)

		return
	}

	lastIndexId,
		err = stdLine03.AddStandardLine(
		&stdLine02,
		ePrefix.XCtx(
			"stdLine03<-stdLine02"))

	expectedLastIndexId =
		len(stdLine01.textFields) +
			len(stdLine02.textFields) - 1

	if lastIndexId != expectedLastIndexId {

		t.Errorf("%v - ERROR\n"+
			"stdLine03<-stdLine02\n"+
			"Expected Last Index Id = '%v'\n"+
			"  Actual Last Index Id = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedLastIndexId,
			lastIndexId)

		return
	}

	stdLine04 := TextLineSpecStandardLine{}.New()

	for idx, val := range stdLine01.textFields {

		var newTextField ITextFieldSpecification

		newTextField,
			err = val.CopyOutITextField(
			ePrefix.XCtx(fmt.Sprintf(
				"stdLine01.textFields[%v]->newTextField",
				idx)))

		if err != nil {
			t.Errorf("%v - ERROR\n"+
				"stdLine01.CopyOutITextField() Failed!\n",
				ePrefix.XCtxEmpty().String())

			return
		}

		stdLine04.textFields =
			append(stdLine04.textFields,
				newTextField)

	}

	for idx, val := range stdLine02.textFields {

		var newTextField ITextFieldSpecification

		newTextField,
			err = val.CopyOutITextField(
			ePrefix.XCtx(fmt.Sprintf(
				"stdLine01.textFields[%v]->newTextField",
				idx)))

		if err != nil {
			t.Errorf("%v - ERROR\n"+
				"stdLine02.CopyOutITextField() Failed!\n",
				ePrefix.XCtxEmpty().String())

			return
		}

		stdLine04.textFields =
			append(stdLine04.textFields,
				newTextField)

	}

	areEqual := stdLine03.Equal(&stdLine04)

	if !areEqual {

		t.Errorf("%v - ERROR\n"+
			"Expected stdLine03==stdLine04\n"+
			"HOWEVER THEY ARE NOT EQUAL!\n",
			ePrefix.XCtxEmpty().String())

		return

	}
}
