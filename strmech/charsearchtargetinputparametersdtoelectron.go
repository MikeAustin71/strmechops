package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
)

// charSearchTargetInputParametersDtoElectron - Provides helper methods for type
// CharSearchTargetInputParametersDto.
//
type charSearchTargetInputParametersDtoElectron struct {
	lock *sync.Mutex
}

// buildFormattedMarqueeTitle - Builds formatted output text for a
// marquee with title.
//
func (searchTargetInputParmsElectron *charSearchTargetInputParametersDtoElectron) buildFormattedMarqueeTitle(
	solidLineChar string,
	title1 string,
	title2 string,
	lineLen int,
	strBuilder *strings.Builder,
	errPrefDto *ePref.ErrPrefixDto) error {

	if searchTargetInputParmsElectron.lock == nil {
		searchTargetInputParmsElectron.lock = new(sync.Mutex)
	}

	searchTargetInputParmsElectron.lock.Lock()

	defer searchTargetInputParmsElectron.lock.Unlock()
	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchTargetInputParametersDtoElectron."+
			"buildFormattedMarqueeTitle()",
		"")

	if err != nil {

		return err
	}

	if strBuilder == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'strBuilder' is invalid!\n"+
			"'strBuilder' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if len(title1) == 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'title1' is invalid!\n"+
			"'title1' is empty and has a string length of zero.\n",
			ePrefix.String())

		return err
	}

	if len(solidLineChar) == 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'solidLineChar' is invalid!\n"+
			"'solidLineChar' is empty and has a string length of zero.\n",
			ePrefix.String())

		return err
	}

	if lineLen < -1 {
		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'labelFieldLen' is invalid!\n"+
			"'labelFieldLen' has a value less than minus one (-1).\n"+
			"labelFieldLen = '%v'\n",
			ePrefix.String(),
			lineLen)

		return err

	}

	var solidLine TextLineSpecSolidLine

	// Title Marquee
	solidLine,
		err = TextLineSpecSolidLine{}.NewFullSolidLineConfig(
		1,
		1,
		solidLineChar,
		lineLen-2,
		"\n",
		ePrefix.XCpy(
			"solidLine"))

	if err != nil {

		return err
	}

	var blankLine *TextLineSpecBlankLines

	blankLine,
		err = TextLineSpecBlankLines{}.NewPtrBlankLines(
		1,
		"\n",
		ePrefix.XCpy(
			"blankLine"))

	if err != nil {

		return err
	}

	var titleLabel *TextFieldSpecLabel

	titleLabel,
		err = TextFieldSpecLabel{}.NewPtrTextLabel(
		title1,
		lineLen,
		TxtJustify.Center(),
		ePrefix.XCpy(
			"title1->titleLabel"))

	if err != nil {

		return err
	}

	err = blankLine.TextBuilder(
		strBuilder,
		ePrefix.XCpy(
			"blankLine#1"))

	if err != nil {
		return err
	}

	err = solidLine.TextBuilder(
		strBuilder,
		ePrefix.XCpy(
			"solidLine#1"))

	if err != nil {

		return err
	}

	stdLine := TextLineSpecStandardLine{}.NewPtr()

	_,
		err = stdLine.AddTextField(
		titleLabel,
		ePrefix)

	if err != nil {

		return err
	}

	err = stdLine.TextBuilder(
		strBuilder,
		ePrefix.XCpy(
			"stdLine->strBuilder #1"))

	if err != nil {

		return err
	}

	if len(title2) > 0 {

		err = titleLabel.SetText(
			title2,
			ePrefix.XCpy(
				"titleLabel<-title2"))

		if err != nil {

			return err
		}

		stdLine.EmptyTextFields()

		_,
			err = stdLine.AddTextField(
			titleLabel,
			ePrefix.XCpy(""+
				"stdLine<-titleLabel"))

		if err != nil {

			return err
		}

		err = stdLine.TextBuilder(
			strBuilder,
			ePrefix.XCpy(
				"stdLine->strBuilder #2"))

		if err != nil {

			return err
		}

	}

	err = solidLine.TextBuilder(
		strBuilder,
		ePrefix.XCpy(
			"solidLine#1"))

	if err != nil {

		return err
	}

	err = blankLine.TextBuilder(
		strBuilder,
		ePrefix.XCpy(
			"blankLine#1"))

	return err
}

// buildFormattedParameterText - Configures and returns low level
// elements of the formatted text generated for an instance of
// CharSearchTargetInputParametersDto.
//
func (searchTargetInputParmsElectron *charSearchTargetInputParametersDtoElectron) buildFormattedParameterText(
	parameterLabel string,
	labelFieldLen int,
	labelRightSpacer *TextFieldSpecLabel,
	parameterValueStr string,
	paramValueFieldLen int,
	strBuilder *strings.Builder,
	errPrefDto *ePref.ErrPrefixDto) error {

	if searchTargetInputParmsElectron.lock == nil {
		searchTargetInputParmsElectron.lock = new(sync.Mutex)
	}

	searchTargetInputParmsElectron.lock.Lock()

	defer searchTargetInputParmsElectron.lock.Unlock()
	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchTargetInputParametersDtoElectron."+
			"buildFormattedParameterText()",
		"")

	if err != nil {

		return err
	}

	if len(parameterLabel) == 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'parameterLabel' is invalid!\n"+
			"'parameterLabel' is empty and has a string length of zero.\n",
			ePrefix.String())

		return err
	}

	if len(parameterValueStr) == 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'parameterValueStr' is invalid!\n"+
			"'parameterValueStr' is empty and has a string length of zero.\n",
			ePrefix.String())

		return err
	}

	if labelRightSpacer == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'labelRightSpacer' is invalid!\n"+
			"'labelRightSpacer' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if strBuilder == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'strBuilder' is invalid!\n"+
			"'strBuilder' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if labelFieldLen < -1 {
		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'labelFieldLen' is invalid!\n"+
			"'labelFieldLen' has a value less than minus one (-1).\n"+
			"labelFieldLen = '%v'\n",
			ePrefix.String(),
			labelFieldLen)

		return err

	}

	if paramValueFieldLen < -1 {
		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'paramValueFieldLen' is invalid!\n"+
			"'paramValueFieldLen' has a value less than minus one (-1).\n"+
			"paramValueFieldLen = '%v'\n",
			ePrefix.String(),
			paramValueFieldLen)

		return err

	}

	stdLine := TextLineSpecStandardLine{}.NewPtr()

	_,
		err = stdLine.AddTextFieldLabel(
		parameterLabel,
		labelFieldLen,
		TxtJustify.Right(),
		ePrefix.XCpy(
			fmt.Sprintf("%v Label",
				parameterLabel)))

	if err != nil {

		return err
	}

	_,
		err = stdLine.AddTextField(
		labelRightSpacer,
		ePrefix.XCpy(
			"labelRightSpacer"))

	if err != nil {

		return err
	}
	_,
		err = stdLine.AddTextFieldLabel(
		parameterValueStr,
		paramValueFieldLen,
		TxtJustify.Left(),
		ePrefix.XCpy(
			fmt.Sprintf("%v Parameter",
				parameterLabel)))

	if err != nil {

		return err
	}

	err = stdLine.TextBuilder(
		strBuilder,
		ePrefix.XCpy(
			fmt.Sprintf("Output %v Text Line",
				parameterLabel)))

	if err != nil {

		return err
	}

	stdLine.Empty()

	return err
}

// buildFormattedTargetString - Creates and formats the Target
// String rune array for text output.
//
func (searchTargetInputParmsElectron *charSearchTargetInputParametersDtoElectron) buildFormattedTargetString(
	labelFieldLen int,
	labelRightSpacer *TextFieldSpecLabel,
	targetInputParms *CharSearchTargetInputParametersDto,
	parameterFieldLen int,
	strBuilder *strings.Builder,
	errPrefDto *ePref.ErrPrefixDto) error {

	if searchTargetInputParmsElectron.lock == nil {
		searchTargetInputParmsElectron.lock = new(sync.Mutex)
	}

	searchTargetInputParmsElectron.lock.Lock()

	defer searchTargetInputParmsElectron.lock.Unlock()
	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchTargetInputParametersDtoElectron."+
			"buildFormattedParameterText()",
		"")

	if err != nil {

		return err
	}

	if labelFieldLen < -1 {
		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'labelFieldLen' is invalid!\n"+
			"'labelFieldLen' has a value less than minus one (-1).\n"+
			"labelFieldLen = '%v'\n",
			ePrefix.String(),
			labelFieldLen)

		return err

	}

	if parameterFieldLen < -1 {
		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'parameterFieldLen' is invalid!\n"+
			"'parameterFieldLen' has a value less than minus one (-1).\n"+
			"labelFieldLen = '%v'\n",
			ePrefix.String(),
			parameterFieldLen)

		return err

	}

	if targetInputParms == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'targetInputParms' is invalid!\n"+
			"'targetInputParms' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if labelRightSpacer == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'labelRightSpacer' is invalid!\n"+
			"'labelRightSpacer' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if strBuilder == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'strBuilder' is invalid!\n"+
			"'strBuilder' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	txtParam := targetInputParms.TargetStringName

	if len(txtParam) == 0 {
		txtParam = "Target String"
	}

	stdLine := TextLineSpecStandardLine{}.NewPtr()

	_,
		err = stdLine.AddTextFieldLabel(
		txtParam,
		labelFieldLen,
		TxtJustify.Right(),
		ePrefix.XCpy(
			"TargetString Label"))

	if err != nil {

		return err
	}
	_,
		err = stdLine.AddTextField(
		labelRightSpacer,
		ePrefix.XCpy(
			"labelRightSpacer"))

	if err != nil {

		return err
	}

	err = stdLine.TextBuilder(
		strBuilder,
		ePrefix.XCpy(
			"Output TargetString Label"))

	if err != nil {

		return err
	}

	stdLine.EmptyTextFields()

	_,
		err = stdLine.AddTextFieldSpacer(
		2,
		ePrefix.XCpy("spacer=2"))

	if err != nil {

		return err
	}

	if targetInputParms.TargetString == nil {

		_,
			err = stdLine.AddTextFieldLabel(
			"Target String is EMPTY!",
			-1,
			TxtJustify.Left(),
			ePrefix.XCpy(
				"TargetString Parameter"))

	} else {

		_,
			err = stdLine.AddTextFieldLabel(
			" "+
				string(targetInputParms.TargetString.CharsArray),
			parameterFieldLen,
			TxtJustify.Left(),
			ePrefix.XCpy(
				"TargetString Parameter"))

	}

	if err != nil {

		return err
	}

	err = stdLine.TextBuilder(
		strBuilder,
		ePrefix.XCpy(
			"Output TargetString Parameter"))

	return err
}

// equalTargetStrings - Compares the Target Strings from two
// instances of CharSearchTargetInputParametersDto to determine if
// they are equivalent.
//
// If both TargetString member variables are 'nil' pointers, this
// method classifies them as equivalent.
//
// If both TargetString member variables are equal in all respects,
// this method returns a boolean value of 'true'. Otherwise, a
// value of 'false' is returned to the calling routine.
//
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  targetInputParms1          *CharSearchTargetInputParametersDto
//     - An instance of CharSearchTargetInputParametersDto. If the
//       internal member variable, 'TargetString', will be compared
//       to the same internal member variable ('TargetString') in
//       parameter 'targetInputParms2' to determine if the two
//       target strings are equivalent.
//
//
//  targetInputParms2          *CharSearchTargetInputParametersDto
//     - An instance of CharSearchTargetInputParametersDto. If the
//       internal member variable, 'TargetString', will be compared
//       to the same internal member variable ('TargetString') in
//       parameter 'targetInputParms1' to determine if the two
//       target strings are equivalent.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  bool
//     - If the comparison of target strings in parameters
//       'targetInputParms1' and 'targetInputParms2' show that the
//       internal member varaiables 'TargetString' are equivalent,
//       this method will return a boolean value of 'true'.
//
//       If the two target strings are NOT equal, this method will
//       return a boolean value of 'false' to the calling function.
//
func (searchTargetInputParmsElectron *charSearchTargetInputParametersDtoElectron) equalTargetStrings(
	targetInputParms1 *CharSearchTargetInputParametersDto,
	targetInputParms2 *CharSearchTargetInputParametersDto) bool {

	if searchTargetInputParmsElectron.lock == nil {
		searchTargetInputParmsElectron.lock = new(sync.Mutex)
	}

	searchTargetInputParmsElectron.lock.Lock()

	defer searchTargetInputParmsElectron.lock.Unlock()

	if targetInputParms1 == nil ||
		targetInputParms2 == nil {
		return false
	}

	if targetInputParms1.TargetString == nil &&
		targetInputParms2.TargetString == nil {

		return true
	}

	if targetInputParms1.TargetString == nil &&
		targetInputParms2.TargetString != nil {

		return false
	}

	if targetInputParms1.TargetString != nil &&
		targetInputParms2.TargetString == nil {

		return false
	}

	if targetInputParms1.TargetString != nil &&
		targetInputParms2.TargetString != nil {

		if !targetInputParms1.TargetString.Equal(
			targetInputParms2.TargetString) {
			return false
		}
	}

	return true
}

// ptr - Returns a pointer to a new instance of
// numberStrKernelElectron.
//
func (searchTargetInputParmsElectron charSearchTargetInputParametersDtoElectron) ptr() *charSearchTargetInputParametersDtoElectron {

	if searchTargetInputParmsElectron.lock == nil {
		searchTargetInputParmsElectron.lock = new(sync.Mutex)
	}

	searchTargetInputParmsElectron.lock.Lock()

	defer searchTargetInputParmsElectron.lock.Unlock()

	return &charSearchTargetInputParametersDtoElectron{
		lock: new(sync.Mutex),
	}
}
