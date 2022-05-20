package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"testing"
	"time"
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
		ePrefix.XCpy(
			"rightMarginSpec"))

	if err != nil {
		return stdLine01, err
	}

	var leftMarginSpec TextFieldSpecSpacer

	leftMarginLen := 6

	leftMarginSpec,
		err = TextFieldSpecSpacer{}.NewSpacer(
		leftMarginLen,
		ePrefix.XCpy(
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
		ePrefix.XCpy(
			"labelSpec"))

	if err != nil {
		return stdLine01, err
	}

	//Index 0
	_,
		err = stdLine01.AddTextField(
		&leftMarginSpec,
		ePrefix.XCpy(
			"stdLine01<-leftMarginSpec"))

	if err != nil {
		return stdLine01, err
	}

	//Index 1

	_,
		err = stdLine01.AddTextField(
		&labelSpec,
		ePrefix.XCpy(
			"stdLine01<-labelSpec"))

	if err != nil {
		return stdLine01, err
	}

	//Index 2
	_,
		err = stdLine01.AddTextField(
		&rightMarginSpec,
		ePrefix.XCpy(
			"stdLine01<-rightMarginSpec"))

	if err != nil {
		return stdLine01, err
	}

	err = stdLine01.IsValidInstanceError(
		ePrefix.XCpy(
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
		ePrefix.XCpy(
			"rightMarginSpec"))

	if err != nil {
		return stdLine02, err
	}

	var leftMarginSpec TextFieldSpecSpacer

	leftMarginLen := 3

	leftMarginSpec,
		err = TextFieldSpecSpacer{}.NewSpacer(
		leftMarginLen,
		ePrefix.XCpy(
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
		ePrefix.XCpy(
			"labelSpec"))

	if err != nil {
		return stdLine02, err
	}

	_,
		err = stdLine02.AddTextField(
		&leftMarginSpec,
		ePrefix.XCpy(
			"stdLine02<-leftMarginSpec"))

	if err != nil {
		return stdLine02, err
	}

	_,
		err = stdLine02.AddTextField(
		&labelSpec,
		ePrefix.XCpy(
			"stdLine02<-labelSpec"))

	if err != nil {
		return stdLine02, err
	}

	_,
		err = stdLine02.AddTextField(
		&rightMarginSpec,
		ePrefix.XCpy(
			"stdLine02<-rightMarginSpec"))

	if err != nil {
		return stdLine02, err
	}

	err = stdLine02.IsValidInstanceError(
		ePrefix.XCpy(
			"stdLine02"))

	return stdLine02, err
}

// createTestTextLineSpecStandardLine03 - Generates a
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
		ePrefix.XCpy(
			"rightMarginSpec"))

	if err != nil {
		return stdLine03, err
	}

	var leftMarginSpec TextFieldSpecSpacer

	leftMarginSpec,
		err = TextFieldSpecSpacer{}.NewSpacer(
		leftMarginSpaces,
		ePrefix.XCpy(
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
		ePrefix.XCpy(
			"labelSpec"))

	if err != nil {
		return stdLine03, err
	}

	_,
		err = stdLine03.AddTextField(
		&leftMarginSpec,
		ePrefix.XCpy(
			"stdLine03<-leftMarginSpec"))

	if err != nil {
		return stdLine03, err
	}

	_,
		err = stdLine03.AddTextField(
		&labelSpec,
		ePrefix.XCpy(
			"stdLine03<-labelSpec"))

	if err != nil {
		return stdLine03, err
	}

	_,
		err = stdLine03.AddTextField(
		&rightMarginSpec,
		ePrefix.XCpy(
			"stdLine03<-rightMarginSpec"))

	if err != nil {
		return stdLine03, err
	}

	err = stdLine03.IsValidInstanceError(
		ePrefix.XCpy(
			"stdLine03"))

	return stdLine03, err
}

// createTestTextLineSpecStandardLine04 - Generates a
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
func createTestTextLineSpecStandardLine04(
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
			"createTestTextLineSpecStandardLine04()",
		"")

	if err != nil {
		return stdLine02, err
	}

	rightMarginLen := 2
	var rightMarginSpec TextFieldSpecSpacer

	rightMarginSpec,
		err = TextFieldSpecSpacer{}.NewSpacer(
		rightMarginLen,
		ePrefix.XCpy(
			"rightMarginSpec"))

	if err != nil {
		return stdLine02, err
	}

	var leftMarginSpec TextFieldSpecSpacer

	leftMarginLen := 3

	leftMarginSpec,
		err = TextFieldSpecSpacer{}.NewSpacer(
		leftMarginLen,
		ePrefix.XCpy(
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
		ePrefix.XCpy(
			"labelSpec"))

	if err != nil {
		return stdLine02, err
	}

	_,
		err = stdLine02.AddTextField(
		&leftMarginSpec,
		ePrefix.XCpy(
			"stdLine02<-leftMarginSpec"))

	if err != nil {
		return stdLine02, err
	}

	_,
		err = stdLine02.AddTextField(
		&labelSpec,
		ePrefix.XCpy(
			"stdLine02<-labelSpec"))

	if err != nil {
		return stdLine02, err
	}

	_,
		err = stdLine02.AddTextField(
		&rightMarginSpec,
		ePrefix.XCpy(
			"stdLine02<-rightMarginSpec"))

	if err != nil {
		return stdLine02, err
	}

	// Starting Series-2 Fields
	rightMarginLen = 3

	rightMarginSpec,
		err = TextFieldSpecSpacer{}.NewSpacer(
		rightMarginLen,
		ePrefix.XCpy(
			"Series-2 rightMarginSpec"))

	if err != nil {
		return stdLine02, err
	}

	leftMarginLen = 4

	leftMarginSpec,
		err = TextFieldSpecSpacer{}.NewSpacer(
		leftMarginLen,
		ePrefix.XCpy(
			"Series-2 leftMarginSpec"))

	if err != nil {
		return stdLine02, err
	}

	label = "Ain't gonna make war no more!"
	fieldLen = len(label) + 6
	txtJustify = TxtJustify.Left()

	labelSpec,
		err = TextFieldSpecLabel{}.NewTextLabel(
		label,
		fieldLen,
		txtJustify,
		ePrefix.XCpy(
			"Series-2 labelSpec"))

	if err != nil {
		return stdLine02, err
	}

	_,
		err = stdLine02.AddTextField(
		&leftMarginSpec,
		ePrefix.XCpy(
			"stdLine02<-leftMarginSpec"))

	if err != nil {
		return stdLine02, err
	}

	_,
		err = stdLine02.AddTextField(
		&labelSpec,
		ePrefix.XCpy(
			"stdLine02<-labelSpec"))

	if err != nil {
		return stdLine02, err
	}

	_,
		err = stdLine02.AddTextField(
		&rightMarginSpec,
		ePrefix.XCpy(
			"stdLine02<-rightMarginSpec"))

	if err != nil {
		return stdLine02, err
	}

	err = stdLine02.IsValidInstanceError(
		ePrefix.XCpy(
			"Series-2 stdLine02"))

	if err != nil {
		return stdLine02, err
	}

	return stdLine02, err
}

// createTestTextLineSpecStandardLine05 - Generates a
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
func createTestTextLineSpecStandardLine05(
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
			"createTestTextLineSpecStandardLine05()",
		"")

	if err != nil {
		return stdLine01, err
	}

	var leftMarginSpec TextFieldSpecSpacer

	leftMarginLen := 6

	leftMarginSpec,
		err = TextFieldSpecSpacer{}.NewSpacer(
		leftMarginLen,
		ePrefix.XCpy(
			"leftMarginSpec"))

	if err != nil {
		return stdLine01, err
	}

	rightMarginLen := 5
	var rightMarginSpec TextFieldSpecSpacer

	rightMarginSpec,
		err = TextFieldSpecSpacer{}.NewSpacer(
		rightMarginLen,
		ePrefix.XCpy(
			"rightMarginSpec"))

	if err != nil {
		return stdLine01, err
	}

	label := "This Is A Date Time:"
	fieldLen := len(label) + 1
	txtJustify := TxtJustify.Right()

	var labelSpec TextFieldSpecLabel

	labelSpec,
		err = TextFieldSpecLabel{}.NewTextLabel(
		label,
		fieldLen,
		txtJustify,
		ePrefix.XCpy(
			"labelSpec"))

	if err != nil {
		return stdLine01, err
	}

	//Index 0
	_,
		err = stdLine01.AddTextField(
		&leftMarginSpec,
		ePrefix.XCpy(
			"stdLine01<-leftMarginSpec"))

	if err != nil {
		return stdLine01, err
	}

	//Index 1

	_,
		err = stdLine01.AddTextField(
		&labelSpec,
		ePrefix.XCpy(
			"stdLine01<-labelSpec"))

	if err != nil {
		return stdLine01, err
	}

	//Index 2
	_,
		err = stdLine01.AddTextField(
		&rightMarginSpec,
		ePrefix.XCpy(
			"stdLine01<-rightMarginSpec"))

	if err != nil {
		return stdLine01, err
	}

	timeZoneName := "America/Chicago"

	tzLocPtr, err2 := time.LoadLocation(timeZoneName)

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"Error returned by time.LoadLocation(timeZoneName)\n"+
			"timeZoneName='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			timeZoneName,
			err2.Error())

		return stdLine01, err
	}

	dateTime := time.Date(
		2021,
		time.Month(10),
		14,
		15,
		28,
		0,
		0,
		tzLocPtr)

	dateTimeFormat :=
		"Monday January 2, 2006 15:04:05.000000000 -0700 MST"

	fieldLen = len(dateTimeFormat)

	textJustification := TxtJustify.Left()

	var incomingTxtFieldDateTime TextFieldSpecDateTime

	incomingTxtFieldDateTime,
		err = TextFieldSpecDateTime{}.NewDateTimeField(
		dateTime,
		fieldLen,
		dateTimeFormat,
		textJustification,
		ePrefix.XCpy("incomingTxtFieldDateTime"))

	if err != nil {

		return stdLine01, err

	}

	//Index 3
	_,
		err = stdLine01.AddTextField(
		&incomingTxtFieldDateTime,
		ePrefix.XCpy(
			"stdLine01<-incomingTxtFieldDateTime"))

	if err != nil {
		return stdLine01, err
	}

	rightMarginLen = 2

	rightMarginSpec,
		err = TextFieldSpecSpacer{}.NewSpacer(
		rightMarginLen,
		ePrefix.XCpy(
			"rightMarginSpec#2"))

	if err != nil {
		return stdLine01, err
	}

	//Index 4
	_,
		err = stdLine01.AddTextField(
		&rightMarginSpec,
		ePrefix.XCpy(
			"stdLine01<-rightMarginSpec#2"))

	if err != nil {
		return stdLine01, err
	}

	fillerChars := "-"
	fillerRepeatCnt := 5

	var fillerTxtFieldOne TextFieldSpecFiller

	fillerTxtFieldOne,
		err = TextFieldSpecFiller{}.NewTextFiller(
		fillerChars,
		fillerRepeatCnt,
		ePrefix)

	//Index 5
	_,
		err = stdLine01.AddTextField(
		&fillerTxtFieldOne,
		ePrefix.XCpy(
			"stdLine01<-fillerTxtFieldOne"))

	if err != nil {
		return stdLine01, err
	}

	rightMarginLen = 1

	rightMarginSpec,
		err = TextFieldSpecSpacer{}.NewSpacer(
		rightMarginLen,
		ePrefix.XCpy(
			"rightMarginSpec#3"))

	if err != nil {
		return stdLine01, err
	}

	//Index 6
	_,
		err = stdLine01.AddTextField(
		&rightMarginSpec,
		ePrefix.XCpy(
			"stdLine01<-rightMarginSpec#3"))

	if err != nil {
		return stdLine01, err
	}

	err = stdLine01.IsValidInstanceError(
		ePrefix.XCpy(
			"stdLine01-final"))

	return stdLine01, err
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
		ePrefix.XCpy(
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
			ePrefix.String(),
			expectedLastIndexId,
			lastIndexId)

		return
	}

	lastIndexId,
		err = stdLine03.AddStandardLine(
		&stdLine02,
		ePrefix.XCpy(
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
			ePrefix.String(),
			expectedLastIndexId,
			lastIndexId)

		return
	}

	stdLine04 := TextLineSpecStandardLine{}.New()

	for idx, val := range stdLine01.textFields {

		var newTextField ITextFieldSpecification

		newTextField,
			err = val.CopyOutITextField(
			ePrefix.XCpy(fmt.Sprintf(
				"stdLine01.textFields[%v]->newTextField",
				idx)))

		if err != nil {
			t.Errorf("%v - ERROR\n"+
				"stdLine01.CopyOutITextField() Failed!\n",
				ePrefix.String())

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
			ePrefix.XCpy(fmt.Sprintf(
				"stdLine01.textFields[%v]->newTextField",
				idx)))

		if err != nil {
			t.Errorf("%v - ERROR\n"+
				"stdLine02.CopyOutITextField() Failed!\n",
				ePrefix.String())

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
			ePrefix.String())

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
		ePrefix.XCpy(
			"stdLine03<-stdLine01"))

	if lastIndexId != expectedLastIndexId {

		t.Errorf("%v - ERROR\n"+
			"stdLine03<-stdLine01\n"+
			"Expected Last Index Id = '%v'\n"+
			"  Actual Last Index Id = '%v'\n",
			ePrefix.String(),
			expectedLastIndexId,
			lastIndexId)

		return
	}

	lastIndexId,
		err = stdLine03.AddStandardLine(
		&stdLine02,
		ePrefix.XCpy(
			"stdLine03<-stdLine02"))

	expectedLastIndexId =
		len(stdLine01.textFields) +
			len(stdLine02.textFields) - 1

	if lastIndexId != expectedLastIndexId {

		t.Errorf("%v - ERROR\n"+
			"stdLine03<-stdLine02\n"+
			"Expected Last Index Id = '%v'\n"+
			"  Actual Last Index Id = '%v'\n",
			ePrefix.String(),
			expectedLastIndexId,
			lastIndexId)

		return
	}

	stdLine04 := TextLineSpecStandardLine{}.New()

	for idx, val := range stdLine01.textFields {

		var newTextField ITextFieldSpecification

		newTextField,
			err = val.CopyOutITextField(
			ePrefix.XCpy(fmt.Sprintf(
				"stdLine01.textFields[%v]->newTextField",
				idx)))

		if err != nil {
			t.Errorf("%v - ERROR\n"+
				"stdLine01.CopyOutITextField() Failed!\n",
				ePrefix.String())

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
			ePrefix.XCpy(fmt.Sprintf(
				"stdLine01.textFields[%v]->newTextField",
				idx)))

		if err != nil {
			t.Errorf("%v - ERROR\n"+
				"stdLine02.CopyOutITextField() Failed!\n",
				ePrefix.String())

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
			ePrefix.String())

		return

	}
}
