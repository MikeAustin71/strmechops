package strmech

import ePref "github.com/MikeAustin71/errpref"

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
	stdLine01 := TextLineSpecStandardLine{}

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
	stdLine02 := TextLineSpecStandardLine{}

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
