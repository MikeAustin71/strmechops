package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type textLineSpecPlainTextAtom struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// textLineSpecPlainTextAtom.
func (txtLinePlainTextAtom textLineSpecPlainTextAtom) ptr() *textLineSpecPlainTextAtom {

	if txtLinePlainTextAtom.lock == nil {
		txtLinePlainTextAtom.lock = new(sync.Mutex)
	}

	txtLinePlainTextAtom.lock.Lock()

	defer txtLinePlainTextAtom.lock.Unlock()

	return &textLineSpecPlainTextAtom{
		lock: new(sync.Mutex),
	}
}

// setPlainTextSpec - Receives a pointer to an instance of
// TextLineSpecPlainText and proceeds to reset all the member
// variable data values using the values passed through input
// parameters.
//
// The TextLineSpecPlainText type provides formatting
// specifications for a simple line of text characters for text
// display, file output or printing.
//
// The plain text line consists of a left margin, the text string,
// the right margin and a line termination character or characters.
// The line termination character is usually a new line character
// ('\n').
//
// Left and right margins consist of zero or more characters
// customized and provided by the calling function.
//
// This method is similar to
// textLineSpecPlainTextNanobot.setPlainTextSpecRunes() with the
// sole exception being that this method receives input parameter
// 'textString' as a string instead of an array of runes.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// The pre-existing member variable data fields for the
// 'plainTxtLine' instance of TextLineSpecPlainText will be
// overwritten and deleted.
//
// Member variable plainTxtLine.textLineReader will be set to
// 'nil'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	 plainTxtLine               *TextLineSpecPlainText
//
//	    - A pointer to an instance of TextLineSpecPlainText.
//	      If this method completes successfully, all member data
//	      variables encapsulated in 'targetPlainTextLine' will be
//	      deleted and overwritten with new values extracted from the
//	      following input parameters.
//
//
//	 leftMarginRunes            []rune
//
//	    - An array of runes containing the text characters which
//	      will be used to construct the left margin of the plain
//	      text line generated by the current instance of
//	      TextLineSpecPlainText.
//
//	      If this parameter is submitted as a zero length or empty
//	      array, NO error will be generated and the existing left
//	      margin will be effectively eliminated from the plain text
//	      line generated by this instance of TextLineSpecPlainText.
//
//	      If the 'leftMarginRunes' array exceeds a length of
//	      one-million array elements, an error will be returned. If
//	      any of the array elements has a rune value of zero ('0'),
//	      an error will be returned.
//
//
//	 rightMarginRunes           []rune
//
//	    - An array of runes containing the text characters which
//	      will be used to construct the right margin of the plain
//	      text line generated by the current instance of
//	      TextLineSpecPlainText.
//
//	      If this parameter is submitted as a zero length or empty
//	      array, NO error will be generated and the existing left
//	      margin will be effectively eliminated from the plain text
//	      line generated by this instance of TextLineSpecPlainText.
//
//	      If the 'rightMarginRunes' array exceeds a length of
//	      one-million array elements, an error will be returned. If
//	      any of the array elements has a rune value of zero ('0'),
//	      an error will be returned.
//
//		textString					string
//
//			A string of text which will be used to configure
//			the text characters generated by the current
//			instance of TextLineSpecPlainText.
//
//			If this parameter is submitted as a zero length
//			or empty string, an error will be returned.
//
//			If this string contains more than one-million
//			characters, an error will be returned.
//
//		textFieldLength				int
//
//			This parameter defines the length of the text
//			field in which the numeric value will be
//			displayed within a number string.
//
//			If 'textFieldLength' is less than the length of
//			the text string ('textString'), it will be
//			automatically set equal to the length of that
//			numeric value string.
//
//			To automatically set the value of 'textFieldLength'
//			to the string length of text string
//			('textString'), set this parameter to a value of
//			minus one (-1).
//
//			If this parameter is submitted with a value less
//			than minus one (-1) or greater than 1-million
//			(1,000,000), an error will be returned.
//
//			Text Field Length Examples
//
//				Example-1
//		         FieldContents String = "1234.5678"
//					FieldContents String Length = 9
//					textFieldLength = 15
//					textFieldJustification = TxtJustify.Center()
//					Text Field String =
//						"   1234.5678   "
//
//				Example-2
//		         FieldContents = "1234.5678"
//					FieldContents String Length = 9
//					textFieldLength = 15
//					textFieldJustification = TxtJustify.Right()
//					Text Field String =
//						"      1234.5678"
//
//				Example-3
//		         FieldContents = "1234.5678"
//					FieldContents String Length = 9
//					textFieldLength = -1
//					textFieldJustification = TxtJustify.Center()
//						// Justification Ignored. Field Length
//						// Equals -1
//					Text Field String =
//						"1234.5678"
//
//				Example-4
//		         FieldContents = "1234.5678"
//					FieldContents String Length = 9
//					textFieldLength = 2
//					textFieldJustification = TxtJustify.Center()
//						// Ignored, because FieldLength Less
//						// Than FieldContents String Length.
//					Text Field String =
//						"1234.5678"
//
//		textFieldJustification		TextJustify
//
//			An enumeration which specifies the justification
//			of the text string ('textString') within the text
//			field length specified by input parameter
//			'textFieldLength'.
//
//			Text justification can only be evaluated in the
//			context of a number string, field length and a
//			'textJustification' object of type TextJustify.
//			This is because text strings with a field length
//			equal to or less than the length of the text string
//			('textString') never use text justification. In
//			these cases, text justification is completely
//			ignored.
//
//			If the field length parameter ('textFieldLength')
//			is greater than the length of the numeric value
//			string, text justification must be equal to one
//			of these three valid values:
//
//				TextJustify(0).Left()
//				TextJustify(0).Right()
//				TextJustify(0).Center()
//
//			You can also use the abbreviated text justification
//			enumeration syntax as follows:
//
//				TxtJustify.Left()
//				TxtJustify.Right()
//				TxtJustify.Center()
//
//			Text Justification Examples
//
//				Example-1
//		         FieldContents String = "1234.5678"
//					FieldContents String Length = 9
//					textFieldLength = 15
//					textFieldJustification = TxtJustify.Center()
//					Text Field String =
//						"   1234.5678   "
//
//				Example-2
//		         FieldContents = "1234.5678"
//					FieldContents String Length = 9
//					textFieldLength = 15
//					textFieldJustification = TxtJustify.Right()
//					Text Field String =
//						"      1234.5678"
//
//				Example-3
//		         FieldContents = "1234.5678"
//					FieldContents String Length = 9
//					textFieldLength = -1
//					textFieldJustification = TxtJustify.Center()
//						// Justification Ignored. Field Length
//						// Equals -1
//					Text Field String =
//						"1234.5678"
//
//				Example-4
//		         FieldContents = "1234.5678"
//					FieldContents String Length = 9
//					textFieldLength = 2
//					textFieldJustification = TxtJustify.Center()
//						// Ignored, because FieldLength Less
//						// Than FieldContents String Length.
//					Text Field String =
//						"1234.5678"
//
//	 newLineChars               []rune
//
//	    - An array of runes containing the character or characters
//	      used to terminate each line of text generated by the
//	      current instance of TextLineSpecPlainText.
//
//	      If this parameter is submitted as a zero length or empty
//	      rune array, it will be reset to the default new line value
//	      ('\n').
//
//	      If the rune array contains invalid zero rune values, an
//	      error will be returned.
//
//	 turnLineTerminatorOff      bool
//
//	    - The 'turnLineTerminatorOff' flag controls whether a line
//	      termination character or characters will be automatically
//	      appended to each line of text produced by
//	      TextLineSpecPlainText.
//
//	      When the boolean flag 'turnLineTerminatorOff' is set to
//	      'false', line terminators as defined by parameter
//	      'newLineChars' WILL BE applied as a line termination
//	      sequence for each line of text produced by
//	      TextLineSpecPlainText.
//
//	      When this boolean value is set to 'true', it turns off or
//	      cancels the automatic generation of line terminators for
//	      each line of text produced by TextLineSpecStandardLine.
//
//
//	 errPrefDto                 *ePref.ErrPrefixDto
//	    - This object encapsulates an error prefix string which is
//	      included in all returned error messages. Usually, it
//	      contains the name of the calling method or methods listed
//	      as a function chain.
//
//	      If no error prefix information is needed, set this parameter
//	      to 'nil'.
//
//	      Type ErrPrefixDto is included in the 'errpref' software
//	      package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	err                        error
//	   - If input parameter 'plainTextLine' is judged to be valid
//	     in all respects, this return parameter will be set to
//	     'nil'.
//
//	     If an error message is returned, the text value for input
//	     parameter 'errPrefDto' (error prefix) will be prefixed or
//	     attached at the beginning of the error message.
func (txtLinePlainTextAtom *textLineSpecPlainTextAtom) setPlainTextSpec(
	plainTxtLine *TextLineSpecPlainText,
	leftMarginChars []rune,
	rightMarginChars []rune,
	textString string,
	textFieldLength int,
	textFieldJustification TextJustify,
	newLineChars []rune,
	turnLineTerminatorOff bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtLinePlainTextAtom.lock == nil {
		txtLinePlainTextAtom.lock = new(sync.Mutex)
	}

	txtLinePlainTextAtom.lock.Lock()

	defer txtLinePlainTextAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecPlainTextAtom."+
			"setPlainTextSpec()",
		"")

	if err != nil {
		return err
	}

	if plainTxtLine == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'plainTxtLine' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	lenTextStr := len(textString)

	if lenTextStr == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'textString' is an empty or zero length string!\n",
			ePrefix.String())

		return err
	}

	if lenTextStr > 1000000 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'textString' is invalid!\n"+
			"'textString' contains over 1-million (1,000,000) characters.\n"+
			"Length of 'textString' = '%v'\n",
			ePrefix.String(),
			lenTextStr)

		return err
	}

	if textFieldLength <= lenTextStr {
		textFieldLength = -1
		textFieldJustification = TxtJustify.None()

	} else {
		// MUST BE -
		// textFieldLength > lenTextStr

		textString,
			err = new(strMechNanobot).
			justifyTextInStrField(
				textString,
				textFieldLength,
				textFieldJustification,
				ePrefix.XCpy(
					"textString<-"))

		if err != nil {
			return err
		}
	}

	lenLeftMargin := len(leftMarginChars)

	if lenLeftMargin > 1000000 {
		err = fmt.Errorf("%v\n"+
			"Error: The 'leftMarginsChars' rune array exceeds\n"+
			"one-million (1,000,000) characters in length.\n"+
			"Length of 'leftMarginsChars' = '%v'\n",
			ePrefix.String(),
			lenLeftMargin)

		return err
	}

	lenRightMargin := len(rightMarginChars)

	if lenRightMargin > 1000000 {
		err = fmt.Errorf("%v\n"+
			"Error: The 'rightMarginChars' rune array exceeds\n"+
			"one-million (1,000,000) characters in length.\n"+
			"Length of 'rightMarginChars' = '%v'\n",
			ePrefix.String(),
			lenRightMargin)

		return err
	}

	lenNewLineChars := len(newLineChars)

	if lenNewLineChars > 1000000 {
		err = fmt.Errorf("%v\n"+
			"Error: The 'newLineChars' rune array exceeds\n"+
			"one-million (1,000,000) characters in length.\n"+
			"Length of 'rightMarginChars' = '%v'\n",
			ePrefix.String(),
			lenNewLineChars)

		return err
	}

	sMechPreon := strMechPreon{}

	if lenLeftMargin > 0 {
		_,
			err =
			sMechPreon.testValidityOfRuneCharArray(
				leftMarginChars,
				ePrefix.XCpy(
					"input parameter leftMarginChars invalid!"))

		if err != nil {
			return err
		}

		err = sMechPreon.copyRuneArrays(
			&plainTxtLine.leftMarginChars,
			&leftMarginChars,
			true,
			ePrefix.XCpy(
				"leftMarginChars->"+
					"plainTxtLine.leftMarginChars"))

		if err != nil {
			return err
		}

	} else {
		// lenLeftMargin == 0
		plainTxtLine.leftMarginChars = nil
	}

	if lenRightMargin > 0 {

		_,
			err =
			sMechPreon.testValidityOfRuneCharArray(
				rightMarginChars,
				ePrefix.XCpy(
					"input parameter rightMarginChars invalid!"))

		if err != nil {
			return err
		}

		err = sMechPreon.copyRuneArrays(
			&plainTxtLine.rightMarginChars,
			&rightMarginChars,
			true,
			ePrefix.XCpy(
				"rightMarginChars->"+
					"plainTxtLine.rightMarginChars"))

		if err != nil {
			return err
		}

	} else {
		// lenRightMargin == 0

		plainTxtLine.rightMarginChars = nil

	}

	if len(newLineChars) > 0 {

		_,
			err =
			sMechPreon.testValidityOfRuneCharArray(
				newLineChars,
				ePrefix.XCpy(
					"input parameter newLineChars invalid!"))

		if err != nil {
			return err
		}

		err = sMechPreon.copyRuneArrays(
			&plainTxtLine.newLineChars,
			&newLineChars,
			true,
			ePrefix.XCpy(
				"newLineChars->"+
					"plainTxtLine.newLineChars"))

	} else {
		// len(newLineChars) == 0
		newLineChars = []rune{'\n'}

	}

	plainTxtLine.textString =
		textString

	plainTxtLine.turnLineTerminatorOff =
		turnLineTerminatorOff

	plainTxtLine.textLineReader = nil

	return err
}

// testValidityOfTextLineSpecPlainText - Receives a pointer to an
// instance of TextLineSpecPlainText and performs a diagnostic
// analysis to determine if that instance is valid in all respects.
//
// If the input parameter 'plainTextLine' is determined to be
// invalid, this method will return a boolean flag ('isValid') of
// 'false'. In addition, an instance of type error ('err') will be
// returned configured with an appropriate error message.
//
// If the input parameter 'plainTextLine' is valid, this method
// will return a boolean flag ('isValid') of 'true' and the
// returned error type ('err') will be set to 'nil'.
//
// If plainTextLine.newLineChars is a zero length array, this method
// will automatically set this value to the default new line
// character ('\n').
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	plainTextLine              *TextLineSpecPlainText
//	   - A pointer to an instance of TextLineSpecPlainText. This
//	     object will be subjected to diagnostic analysis in order
//	     to determine if all the member variables contain valid
//	     values.
//
//
//	errPrefDto                 *ePref.ErrPrefixDto
//	   - This object encapsulates an error prefix string which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods listed
//	     as a function chain.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     Type ErrPrefixDto is included in the 'errpref' software
//	     package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	isValid                    bool
//	   - If input parameter 'plainTextLine' is judged to be valid
//	     in all respects, this return parameter will be set to
//	     'true'.
//
//	   - If input parameter 'plainTextLine' is found to be invalid,
//	     this return parameter will be set to 'false'.
//
//
//	err                        error
//	   - If input parameter 'plainTextLine' is judged to be valid
//	     in all respects, this return parameter will be set to
//	     'nil'.
//
//	     If input parameter, 'plainTextLine' is found to be
//	     invalid, this return parameter will be configured with an
//	     appropriate error message.
//
//	     If an error message is returned, the text value for input
//	     parameter 'errPrefDto' (error prefix) will be prefixed or
//	     attached at the beginning of the error message.
func (txtLinePlainTextAtom *textLineSpecPlainTextAtom) testValidityOfTextLineSpecPlainText(
	plainTextLine *TextLineSpecPlainText,
	errPrefDto *ePref.ErrPrefixDto) (
	isValid bool,
	err error) {

	if txtLinePlainTextAtom.lock == nil {
		txtLinePlainTextAtom.lock = new(sync.Mutex)
	}

	txtLinePlainTextAtom.lock.Lock()

	defer txtLinePlainTextAtom.lock.Unlock()

	isValid = false

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecPlainTextAtom."+
			"testValidityOfTextLineSpecPlainText()",
		"")

	if err != nil {
		return isValid, err
	}

	if plainTextLine == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'plainTextLine' is a nil pointer!\n",
			ePrefix.String())

		return isValid, err
	}

	lenTextStr := len(plainTextLine.textString)

	if lenTextStr == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: plainTextLine.textString is an empty string!\n"+
			"No Text String have been configured for\n"+
			"this Plain Text Line Specification!\n",
			ePrefix.String())

		return isValid, err
	}

	if lenTextStr > 1000000 {

		err = fmt.Errorf("%v\n"+
			"Error: plainTextLine.textString is invalid!\n"+
			"The Length Of 'plainTextLine.textString' is greater than "+
			"1-million (1,000,000) characters\n",
			ePrefix.String())

		return isValid, err
	}

	lenNewLineChars := len(plainTextLine.newLineChars)

	if lenNewLineChars == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: plainTextLine.newLineChars is empty!\n"+
			"New Line Characters have NOT been configured for\n"+
			"this Plain Text Line Specification!\n",
			ePrefix.String())

		return isValid, err
	}

	if lenNewLineChars > 1000000 {

		err = fmt.Errorf("%v\n"+
			"Error: plainTextLine.newLineChars is invalid!\n"+
			"The number of new line characters is greater than\n"+
			"1-million (1,000,000) characters!\n",
			ePrefix.String())

		return isValid, err
	}

	sMechPreon := strMechPreon{}

	lenLeftMarginChars := len(plainTextLine.leftMarginChars)

	if lenLeftMarginChars > 1000000 {

		err = fmt.Errorf("%v\n"+
			"Error: plainTextLine.leftMarginChars is invalid!\n"+
			"The number of left margin characters is greater than\n"+
			"1-million (1,000,000) characters!\n",
			ePrefix.String())

		return isValid, err
	}

	if lenLeftMarginChars > 0 {
		_,
			err =
			sMechPreon.testValidityOfRuneCharArray(
				plainTextLine.leftMarginChars,
				ePrefix.XCpy(
					"plainTextLine.leftMarginChars invalid!"))

		if err != nil {
			return isValid, err
		}

	}

	lenRightMarginChars := len(plainTextLine.rightMarginChars)

	if lenRightMarginChars > 1000000 {

		err = fmt.Errorf("%v\n"+
			"Error: plainTextLine.rightMarginChars is invalid!\n"+
			"The number of right margin characters is greater than\n"+
			"1-million (1,000,000) characters!\n",
			ePrefix.String())

		return isValid, err
	}

	if len(plainTextLine.rightMarginChars) > 0 {
		_,
			err =
			sMechPreon.testValidityOfRuneCharArray(
				plainTextLine.rightMarginChars,
				ePrefix.XCpy(
					"plainTextLine.rightMarginChars invalid!"))

		if err != nil {
			return isValid, err
		}

	}

	isValid = true

	return isValid, err
}
