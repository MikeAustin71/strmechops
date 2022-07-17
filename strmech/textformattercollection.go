package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// TextFormatterCollection - This type contains a collection of
// Text Formatter Data Transfer Objects (TextFormatterDto). These
// object contain the specifications used to format text fields and
// lines of text for  screen displays, file output and printing.
//
type TextFormatterCollection struct {
	fmtCollection []TextFormatterDto
	// Text

	stdTextLineParamCollection []TextFmtParamsLineColumns
	// Standard Format Parameter Collection.
	// Provides standard text formats

	lock *sync.Mutex
}

// AddLine1Col - Adds Field Text and Format Parameters for
// Format Type 'Line1Column'.
//
// The 'Line1Column' Text Line type is designed to produce a single
// line of text consisting of one text field with optional left and
// right margins. This single text field is referred to as
// 'Column1'.
//
// This method will assign previously configured (a.k.a. default)
// Format Parameters to this 'Line1Column' Text Line. The prior
// configuration of these 'Line1Column' Format Parameters is
// a requirement and errors will be generated if these standard
// Format Parameters have not yet been created.
//
// This method will extract those previously created Standard
// Format Parameters for 'Line1Column' Text Lines from the
// Standard Text Line Parameters collection maintained by this
// instance of TextFormatterCollection.
//
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// The standard parameters for 'Line1Column' Text Lines must be
// configured in the Standard Text Line Parameters Collection
// before calling this method.
//
// If the standard parameters for 'Line1Column' Text Lines were
// not previously configured in the Standard Text Line Parameters
// Collection, an error will be returned.
//
// ----------------------------------------------------------------
//
// To configure the standard parameters for 'Line1Column' Text
// Lines, call one of the following methods:
//   TextFormatterCollection.CfgLine1Col()
//   TextFormatterCollection.SetStdFormatParamsLine1Col()
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  column1Field               interface{}
//     - This parameter is an empty interface which must contain
//       of several specific types. This empty interface type will
//       be converted to a string and configured as the single text
//       field in this 'Line1Column' Text Line.
//
//       Supported types which may be submitted through this empty
//       interface parameter are listed as follows:
//          string
//          bool
//          uint, uint8, uint16, uint32, uint64,
//          int, int8, int16, int32, int64
//          float32, float64
//          *big.Int *big.Float
//          fmt.Stringer (types that support this interface)
//
//       If the 'column1Field' is not convertible to one of the
//       supported types, an error will be returned.
//
//       If the converted string value for 'column1Field' is empty,
//       it will be defaulted to a single white space character
//       (" ").
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings
//          containing error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of
//                          ErrPrefixDto. ErrorPrefixInfo from this
//                          object will be copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package,
//       "github.com/MikeAustin71/errpref".
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (txtFmtCollection *TextFormatterCollection) AddLine1Col(
	column1Field interface{},
	errorPrefix interface{}) error {

	if txtFmtCollection.lock == nil {
		txtFmtCollection.lock = new(sync.Mutex)
	}

	txtFmtCollection.lock.Lock()

	defer txtFmtCollection.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFormatterCollection."+
			"AddLine1Col()",
		"")

	if err != nil {
		return err
	}

	var foundStdParams bool
	var stdLineColsFmt TextFmtParamsLineColumns

	foundStdParams,
		stdLineColsFmt,
		err =
		textFormatterCollectionElectron{}.ptr().
			findStdTxtLineParameters(
				txtFmtCollection,
				TxtFieldType.Line1Column(),
				ePrefix.XCpy(
					"TxtFieldType.Line1Column()"))

	if err != nil {
		return err
	}

	if !foundStdParams {

		err = fmt.Errorf("%v - Error\n"+
			"Could NOT locate Standard Text Line Parameter Format\n"+
			"for Text Field Type: %v.\n"+
			"Use one of the 'SetStdFormatParams' to configure a\n"+
			"new Standard Text Line Column Parameters Format. ",
			ePrefix.String(),
			TxtFieldType.Line1Column().String())

		return err
	}

	var column1FieldText string

	column1FieldText,
		err = textSpecificationAtom{}.ptr().
		convertParamEmptyInterfaceToString(
			column1Field,
			"column1Field",
			ePrefix.XCpy(
				"column1Field"))

	if err != nil {
		return err
	}

	newLine1Col := TextFormatterDto{
		FormatType: TxtFieldType.Line1Column(),
		DateTime:   TextFieldDateTimeDto{},
		Filler:     TextFieldFillerDto{},
		Label:      TextFieldLabelDto{},
		Spacer:     TextFieldSpacerDto{},
		BlankLine:  TextLineBlankDto{},
		SolidLine:  TextLineSolidLineDto{},
		LineColumns: TextLineColumnsDto{
			FormatType:    TxtFieldType.Line1Column(),
			Col1FieldText: column1FieldText,
			Col2FieldText: "",
			Col3FieldText: "",
			Col4FieldText: "",
			Col5FieldText: "",
			Col6FieldText: "",
			Col7FieldText: "",
			Col8FieldText: "",
			FmtParameters: stdLineColsFmt,
			lock:          nil,
		},
	}

	txtFmtCollection.fmtCollection =
		append(
			txtFmtCollection.fmtCollection,
			newLine1Col)

	return err
}

// CfgLine1Col - Allows the user to configure both the field value
// and the Format Parameters for text line Format Type
// 'Line1Column'.
//
// The 'Line1Column' Text Line type is designed to produce a single
// line of text consisting of one text field with optional left and
// right margins. This single text field is referred to as
// 'Column1'.
//
// Unlike method TextFormatterCollection.AddLine1Col(), this method
// has no requirement for previously configured Standard Format
// Parameters because those parameters are created in a single call
// to this method.
//
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// This method will create the Standard Format Parameters for this
// and all future 'Line1Column' Text types created by this instance
// of TextFormatterCollection. After calling this method once,
// users should call TextFormatterCollection.AddLine1Col() to
// reduce the number of input parameters required to produce other
// 'Line1Column' Text type
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  leftMarginStr              string
//     - The contents of the string will be used as the left margin
//       for the 'Column1' field.
//
//       If no left margin is required, set 'leftMarginStr' to a
//       zero length or empty string, and no left margin will be
//       created.
//
//
//  column1Field               interface{}
//     - This parameter is an empty interface which must contain
//       one of several specific types. This empty interface type
//       will be converted to a string and configured as the single
//       text field in this 'Line1Column' Text Line.
//
//       Supported types which may be submitted through this empty
//       interface parameter are listed as follows:
//          string
//          bool
//          uint, uint8, uint16, uint32, uint64,
//          int, int8, int16, int32, int64
//          float32, float64
//          *big.Int *big.Float
//          fmt.Stringer (types that support this interface)
//
//       If the 'column1Field' is not convertible to one of the
//       supported types, an error will be returned.
//
//       If the converted string value for 'column1Field' is empty,
//       it will be defaulted to a single white space character
//       (" ").
//
//
//  column1FieldLength         int
//     - This is the length of the text field in which the
//       formatted 'column1Field' string will be displayed. If
//       'column1FieldLength' is less than the length of the
//       'column1Field' string, it will be automatically set equal
//       to the 'column1Field' string length.
//
//       To automatically set the value of 'column1FieldLength' to
//       the length of 'column1Field', set this parameter to a
//       value of minus one (-1).
//
//       If this parameter is submitted with a value less than
//       minus one (-1) or greater than 1-million (1,000,000), an
//       error will be returned.
//
//
//  column1FieldJustify        TextJustify
//      An enumeration value specifying the justification of the
//      'column1Field' string within the text field specified by
//      'column1FieldLength'.
//
//      Text justification can only be evaluated in the context of
//      a text label, field length and a Text Justification object
//      of type TextJustify. This is because text labels with a
//      field length equal to or less than the length of the text
//      label never use text justification. In these cases, text
//      justification is completely ignored.
//
//      If the field length is greater than the length of the text
//      label, text justification must be equal to one of these
//      three valid values:
//          TextJustify(0).Left()
//          TextJustify(0).Right()
//          TextJustify(0).Center()
//
//      You can also use the abbreviated text justification
//      enumeration syntax as follows:
//
//          TxtJustify.Left()
//          TxtJustify.Right()
//          TxtJustify.Center()
//
//
//  rightMarginStr             string
//     - The contents of the string will be used as the right
//       margin for the 'Column1' field.
//
//       If no right margin is required, set 'rightMarginStr' to a
//       zero length or empty string, and no right margin will be
//       created.
//
//
//   turnLineTerminationOff    bool
//     - When this parameter is set to 'true', no line termination
//       sequence will be configured for this 'Line1Column' Text
//       Line.
//
//       Text Lines operate on the assumption that a line
//       termination is standard operating procedure. The default
//       line terminator for text lines is the new line character,
//       '\n'.
//
//       Users have the option of turning off the entire line
//       termination sequence if this parameter is set to 'true'.
//
//
//  lineTerminator             string
//     - If this parameter is set to an empty string, the default
//       line terminator, a new line character '\n', will be
//       applied to this 'Line1Column' Text Line.
//
//       If this string is populated and the string length is
//       greater than zero (0), an alternate line termination
//       sequence will be configured using the characters provided
//       in the 'lineTerminator' string.
//
//       Remember that the application of a line termination
//       sequence is controlled by parameter
//       'turnLineTerminationOff'. If 'turnLineTerminationOff' is
//       set to 'true', no line termination characters will be
//       configured for this 'Line1Column' Text Line.
//
//
//  maxLineLength              int
//     - The maximum length of the line on which this 'Line1Column'
//       Text will be presented.
//
//       Set this parameter to minus one (-1), and no maximum line
//       limits will be applied.
//
//       'maxLineLength' is used in conjunction with parameter
//       'turnAutoLineLengthBreaksOn' to automatically place text
//       fields on separate text lines when that text exceeds the
//       maximum text line length ('maxLineLength').
//
//       If the value of 'maxLineLength' is less than minus one
//       (-1), an error will be returned.
//
//       If the value of 'maxLineLength' is zero (0), an error will
//       be returned.
//
//
//  turnAutoLineLengthBreaksOn bool
//     - This parameter controls whether text lines which exceed
//       the maximum line length ('maxLineLength') are broken up
//       and presented on the following line.
//
//       To apply automatic line breaking at the maximum line
//       length, set the value of this parameter to 'true'.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings
//          containing error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of
//                          ErrPrefixDto. ErrorPrefixInfo from this
//                          object will be copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package,
//       "github.com/MikeAustin71/errpref".
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method completes successfully and no errors are
//       encountered, this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (txtFmtCollection *TextFormatterCollection) CfgLine1Col(
	leftMarginStr string,
	column1Field interface{},
	column1FieldLength int,
	column1FieldJustify TextJustify,
	rightMarginStr string,
	turnLineTerminationOff bool,
	lineTerminator string,
	maxLineLength int,
	turnAutoLineLengthBreaksOn bool,
	errorPrefix interface{}) error {

	if txtFmtCollection.lock == nil {
		txtFmtCollection.lock = new(sync.Mutex)
	}

	txtFmtCollection.lock.Lock()

	defer txtFmtCollection.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFormatterCollection."+
			"CfgLine1Col()",
		"")

	if err != nil {
		return err
	}

	var column1FieldText string

	column1FieldText,
		err = textSpecificationAtom{}.ptr().
		convertParamEmptyInterfaceToString(
			column1Field,
			"column1Field",
			ePrefix.XCpy(
				"column1Field"))

	if err != nil {
		return err
	}

	if column1FieldLength < -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'Col1FieldLength' is invalid!\n"+
			"'column1FieldLength' has a value less than minus one (-1)\n"+
			"column1FieldLength = '%v'\n",
			ePrefix.String(),
			column1FieldLength)

		return err
	}

	if !column1FieldJustify.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'Col1FieldJustify' is invalid!\n"+
			"'Col1FieldJustify' must be set to 'Left', 'Right' or 'Center'.\n"+
			"Col1FieldJustify String Value  = '%v'\n"+
			"Col1FieldJustify Integer Value = '%v'\n",
			ePrefix.String(),
			column1FieldJustify.String(),
			column1FieldJustify.XValueInt())

		return err

	}

	if maxLineLength < -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'maxLineLength' is invalid!\n"+
			"'maxLineLength' has a value less than minus one (-1)\n"+
			"maxLineLength = '%v'\n",
			ePrefix.String(),
			maxLineLength)

		return err

	}

	if maxLineLength == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'maxLineLength' is invalid!\n"+
			"'maxLineLength' has a value of zero (0).\n",
			ePrefix.String())

		return err

	}

	newStdFmtParams := TextFmtParamsLineColumns{
		FormatType:                 TxtFieldType.Line1Column(),
		Col1LeftMarginStr:          leftMarginStr,
		Col1FieldLength:            column1FieldLength,
		Col1FieldJustify:           column1FieldJustify,
		Col1RightMarginStr:         rightMarginStr,
		Col2LeftMarginStr:          "",
		Col2FieldLength:            0,
		Col2FieldJustify:           0,
		Col2RightMarginStr:         "",
		Col3LeftMarginStr:          "",
		Col3FieldLength:            0,
		Col3FieldJustify:           0,
		Col3RightMarginStr:         "",
		Col4LeftMarginStr:          "",
		Col4FieldLength:            0,
		Col4FieldJustify:           0,
		Col4RightMarginStr:         "",
		Col5LeftMarginStr:          "",
		Col5FieldLength:            0,
		Col5FieldJustify:           0,
		Col5RightMarginStr:         "",
		Col6LeftMarginStr:          "",
		Col6FieldLength:            0,
		Col6FieldJustify:           0,
		Col6RightMarginStr:         "",
		Col7LeftMarginStr:          "",
		Col7FieldLength:            0,
		Col7FieldJustify:           0,
		Col7RightMarginStr:         "",
		Col8LeftMarginStr:          "",
		Col8FieldLength:            0,
		Col8FieldJustify:           0,
		Col8RightMarginStr:         "",
		TurnLineTerminationOff:     turnLineTerminationOff,
		LineTerminator:             lineTerminator,
		MaxLineLength:              maxLineLength,
		TurnAutoLineLengthBreaksOn: turnAutoLineLengthBreaksOn,
		isValid:                    true,
		lock:                       nil,
	}

	newLine1Col := TextFormatterDto{
		FormatType: 0,
		DateTime:   TextFieldDateTimeDto{},
		Filler:     TextFieldFillerDto{},
		Label:      TextFieldLabelDto{},
		Spacer:     TextFieldSpacerDto{},
		BlankLine:  TextLineBlankDto{},
		SolidLine:  TextLineSolidLineDto{},
		LineColumns: TextLineColumnsDto{
			FormatType:    TxtFieldType.Line1Column(),
			Col1FieldText: column1FieldText,
			Col2FieldText: "",
			Col3FieldText: "",
			Col4FieldText: "",
			Col5FieldText: "",
			Col6FieldText: "",
			Col7FieldText: "",
			Col8FieldText: "",
			FmtParameters: TextFmtParamsLineColumns{},
			lock:          nil,
		},
	}

	newLine1Col.LineColumns.FmtParameters.CopyIn(
		&newStdFmtParams)

	txtFmtCollection.fmtCollection =
		append(
			txtFmtCollection.fmtCollection,
			newLine1Col)

	err = textFormatterCollectionElectron{}.ptr().
		cfgNewStdTxtLineParameters(
			txtFmtCollection,
			newStdFmtParams,
			ePrefix.XCpy(
				"newStdFmtParams"))

	return err
}

// GetLengthFormatterCollection - Returns the length of the Text
// Formatter Collection contained in the current instance of
// TextFormatterCollection.
//
// The Text Formatter Collection is an array of Text Formatter Data
// Transfer Objects (TextFormatterDto).
//
// The collection length is returned as an integer value.
//
func (txtFmtCollection *TextFormatterCollection) GetLengthFormatterCollection() int {

	if txtFmtCollection.lock == nil {
		txtFmtCollection.lock = new(sync.Mutex)
	}

	txtFmtCollection.lock.Lock()

	defer txtFmtCollection.lock.Unlock()

	return len(txtFmtCollection.fmtCollection)
}

// GetLengthStdTextLineParamCollection - Returns the length of the
// Standard Text Line Parameter Collection contained in the current
// instance of TextFormatterCollection.
//
// The Standard Text Line Parameter Collection is an array of Text
// Format Parameters for Lines and Columns
// (TextFmtParamsLineColumns).
//
// The Standard Text Line Parameter Collection is used to produced
// standardized lines of texts containing between one and eight
// label fields.
//
// The collection length is returned as an integer value.
//
func (txtFmtCollection *TextFormatterCollection) GetLengthStdTextLineParamCollection() int {

	if txtFmtCollection.lock == nil {
		txtFmtCollection.lock = new(sync.Mutex)
	}

	txtFmtCollection.lock.Lock()

	defer txtFmtCollection.lock.Unlock()

	return len(txtFmtCollection.stdTextLineParamCollection)
}

// SetStdFormatParamsLine1Col - Sets the Standard Format Parameters
// for a text line consisting of one text column (Line1Column).
// This standard format will be applied as the default format of
// all 'Line1Column' Text Format Operations.
//
// After configuring Standard Format Parameters for 'Line1Column'
// Text Lines, users should configure additional 'Line1Column' Text
// Lines using method TextFormatterCollection.AddLine1Col() in
// order to reduce the number of input parameters required to
// produce a 'Line1Column' Text Line.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  leftMarginStr              string
//     - The contents of the string will be used as the left margin
//       for the 'Column1' field.
//
//       If no left margin is required, set 'leftMarginStr' to a
//       zero length or empty string, and no left margin will be
//       created.
//
//
//  column1FieldLength         int
//     - This is the length of the text field in which the
//       formatted 'column1Field' string will be displayed. If
//       'column1FieldLength' is less than the length of the
//       'column1Field' string, it will be automatically set equal
//       to the 'column1Field' string length.
//
//       To automatically set the value of 'column1FieldLength' to
//       the length of 'column1Field', set this parameter to a
//       value of minus one (-1).
//
//       If this parameter is submitted with a value less than
//       minus one (-1) or greater than 1-million (1,000,000), an
//       error will be returned.
//
//
//  column1FieldJustify        TextJustify
//      An enumeration value specifying the justification of the
//      'column1Field' string within the text field specified by
//      'column1FieldLength'.
//
//      Text justification can only be evaluated in the context of
//      a text label, field length and a Text Justification object
//      of type TextJustify. This is because text labels with a
//      field length equal to or less than the length of the text
//      label never use text justification. In these cases, text
//      justification is completely ignored.
//
//      If the field length is greater than the length of the text
//      label, text justification must be equal to one of these
//      three valid values:
//          TextJustify(0).Left()
//          TextJustify(0).Right()
//          TextJustify(0).Center()
//
//      You can also use the abbreviated text justification
//      enumeration syntax as follows:
//
//          TxtJustify.Left()
//          TxtJustify.Right()
//          TxtJustify.Center()
//
//
//  rightMarginStr             string
//     - The contents of the string will be used as the right
//       margin for the 'Column1' field.
//
//       If no right margin is required, set 'rightMarginStr' to a
//       zero length or empty string, and no right margin will be
//       created.
//
//
//   turnLineTerminationOff    bool
//     - When this parameter is set to 'true', no line termination
//       sequence will be configured for this 'Line1Column' Text
//       Line.
//
//       Text Lines operate on the assumption that a line
//       termination is standard operating procedure. The default
//       line terminator for text lines is the new line character,
//       '\n'.
//
//       Users have the option of turning off the entire line
//       termination sequence if this parameter is set to 'true'.
//
//
//  lineTerminator             string
//     - If this parameter is set to an empty string, the default
//       line terminator, a new line character '\n', will be
//       applied to 'Line1Column' Text Lines.
//
//       If this string is populated and the string length is
//       greater than zero (0), an alternate line termination
//       sequence will be configured using the characters provided
//       in the 'lineTerminator' string.
//
//       Remember that the application of a line termination
//       sequence is controlled by parameter
//       'turnLineTerminationOff'. If 'turnLineTerminationOff' is
//       set to 'true', no line termination characters will be
//       configured for this 'Line1Column' Text Lines.
//
//
//  maxLineLength              int
//     - The maximum length of the line on which this 'Line1Column'
//       Text will be presented.
//
//       Set this parameter to minus one (-1), and no maximum line
//       limits will be applied.
//
//       'maxLineLength' is used in conjunction with parameter
//       'turnAutoLineLengthBreaksOn' to automatically place text
//       fields on separate text lines when that text exceeds the
//       maximum text line length ('maxLineLength').
//
//       If the value of 'maxLineLength' is less than minus one
//       (-1), an error will be returned.
//
//       If the value of 'maxLineLength' is zero (0), an error will
//       be returned.
//
//
//  turnAutoLineLengthBreaksOn bool
//     - This parameter controls whether text lines which exceed
//       the maximum line length ('maxLineLength') are broken up
//       and presented on on the following line.
//
//       To apply automatic line breaking at the maximum line
//       length, set the value of this parameter to 'true'.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings
//          containing error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of
//                          ErrPrefixDto. ErrorPrefixInfo from this
//                          object will be copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package,
//       "github.com/MikeAustin71/errpref".
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method completes successfully and no errors are
//       encountered, this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (txtFmtCollection *TextFormatterCollection) SetStdFormatParamsLine1Col(
	leftMarginStr string,
	column1FieldLength int,
	column1FieldJustify TextJustify,
	rightMarginStr string,
	lineTerminator string,
	turnLineTerminationOff bool,
	maxLineLength int,
	turnAutoLineLengthBreaksOn bool,
	errorPrefix interface{}) error {

	if txtFmtCollection.lock == nil {
		txtFmtCollection.lock = new(sync.Mutex)
	}

	txtFmtCollection.lock.Lock()

	defer txtFmtCollection.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFormatterCollection."+
			"SetStdFormatParamsLine1Col()",
		"")

	if err != nil {
		return err
	}

	if column1FieldLength < -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'column1FieldLength' is invalid!\n"+
			"'column1FieldLength' has a value less than minus one (-1)\n"+
			"column1FieldLength = '%v'\n",
			ePrefix.String(),
			column1FieldLength)

		return err
	}

	if !column1FieldJustify.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'column1FieldJustify' is invalid!\n"+
			"'column1FieldJustify' must be set to 'Left', 'Right' or 'Center'.\n"+
			"column1FieldJustify String Value  = '%v'\n"+
			"column1FieldJustify Integer Value = '%v'\n",
			ePrefix.String(),
			column1FieldJustify.String(),
			column1FieldJustify.XValueInt())

		return err

	}

	if maxLineLength < -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'maxLineLength' is invalid!\n"+
			"'maxLineLength' has a value less than minus one (-1)\n"+
			"maxLineLength = '%v'\n",
			ePrefix.String(),
			maxLineLength)

		return err

	}

	newStdFmtParams := TextFmtParamsLineColumns{
		FormatType:                 TxtFieldType.Line1Column(),
		Col1LeftMarginStr:          leftMarginStr,
		Col1FieldLength:            column1FieldLength,
		Col1FieldJustify:           column1FieldJustify,
		Col1RightMarginStr:         rightMarginStr,
		Col2LeftMarginStr:          "",
		Col2FieldLength:            0,
		Col2FieldJustify:           0,
		Col2RightMarginStr:         "",
		Col3LeftMarginStr:          "",
		Col3FieldLength:            0,
		Col3FieldJustify:           0,
		Col3RightMarginStr:         "",
		Col4LeftMarginStr:          "",
		Col4FieldLength:            0,
		Col4FieldJustify:           0,
		Col4RightMarginStr:         "",
		Col5LeftMarginStr:          "",
		Col5FieldLength:            0,
		Col5FieldJustify:           0,
		Col5RightMarginStr:         "",
		Col6LeftMarginStr:          "",
		Col6FieldLength:            0,
		Col6FieldJustify:           0,
		Col6RightMarginStr:         "",
		Col7LeftMarginStr:          "",
		Col7FieldLength:            0,
		Col7FieldJustify:           0,
		Col7RightMarginStr:         "",
		Col8LeftMarginStr:          "",
		Col8FieldLength:            0,
		Col8FieldJustify:           0,
		Col8RightMarginStr:         "",
		TurnLineTerminationOff:     turnLineTerminationOff,
		LineTerminator:             lineTerminator,
		MaxLineLength:              maxLineLength,
		TurnAutoLineLengthBreaksOn: turnAutoLineLengthBreaksOn,
		isValid:                    true,
		lock:                       nil,
	}

	err = textFormatterCollectionElectron{}.ptr().
		cfgNewStdTxtLineParameters(
			txtFmtCollection,
			newStdFmtParams,
			ePrefix.XCpy(
				"newStdFmtParams"))

	return err
}
