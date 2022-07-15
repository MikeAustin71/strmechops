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
// Format Type 'Line1Column' Text. The 'Line1Column' Text is
// designed to produce a single line of text consisting of one
// text field with optional left and right margins.
//
// This method will assign previously configured (a.k.a. default)
// Format Parameters to this 'Line1Column' Text Line. The prior
// configuration of these 'Line1Column' Format Parameters is
// a requirement and errors will be generated if these standard
// Format Parameters have not yet been created.
//
// This method will extract the previously created Standard Format
// Parameters for 'Line1Column' Text Lines from the Standard Text
// Line Parameters collection maintained by this instance of
// TextFormatterCollection.
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
// Lines call one of the following methods:
//   TextFormatterCollection.CfgLine1Col()
//   TextFormatterCollection.SetStdFormatParamsLine1Col()
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  column1FieldText           string
//     - This string contains the text which will be configured
//       for the single text field created for the 'Line1Column'
//       Text Line created by this method.
//
//       If this string parameter is empty, it will be defaulted to
//       a single white space character (" ").
func (txtFmtCollection *TextFormatterCollection) AddLine1Col(
	column1FieldText string,
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

func (txtFmtCollection *TextFormatterCollection) CfgLine1Col(
	leftMarginStr string,
	col1FieldText string,
	col1FieldLength int,
	col1FieldJustify TextJustify,
	rightMarginStr string,
	lineTerminator string,
	maxLineLength int,
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

	if col1FieldLength < -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'Col1FieldLength' is invalid!\n"+
			"'Col1FieldLength' has a value less than minus one (-1)\n"+
			"Col1FieldLength = '%v'\n",
			ePrefix.String(),
			col1FieldLength)

		return err
	}

	if !col1FieldJustify.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'Col1FieldJustify' is invalid!\n"+
			"'Col1FieldJustify' must be set to 'Left', 'Right' or 'Center'.\n"+
			"Col1FieldJustify String Value  = '%v'\n"+
			"Col1FieldJustify Integer Value = '%v'\n",
			ePrefix.String(),
			col1FieldJustify.String(),
			col1FieldJustify.XValueInt())

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
		FormatType:         TxtFieldType.Line1Column(),
		Col1LeftMarginStr:  leftMarginStr,
		Col1FieldLength:    col1FieldLength,
		Col1FieldJustify:   col1FieldJustify,
		Col1RightMarginStr: rightMarginStr,
		Col2LeftMarginStr:  "",
		Col2FieldLength:    0,
		Col2FieldJustify:   0,
		Col2RightMarginStr: "",
		Col3LeftMarginStr:  "",
		Col3FieldLength:    0,
		Col3FieldJustify:   0,
		Col3RightMarginStr: "",
		Col4LeftMarginStr:  "",
		Col4FieldLength:    0,
		Col4FieldJustify:   0,
		Col4RightMarginStr: "",
		Col5LeftMarginStr:  "",
		Col5FieldLength:    0,
		Col5FieldJustify:   0,
		Col5RightMarginStr: "",
		Col6LeftMarginStr:  "",
		Col6FieldLength:    0,
		Col6FieldJustify:   0,
		Col6RightMarginStr: "",
		Col7LeftMarginStr:  "",
		Col7FieldLength:    0,
		Col7FieldJustify:   0,
		Col7RightMarginStr: "",
		Col8LeftMarginStr:  "",
		Col8FieldLength:    0,
		Col8FieldJustify:   0,
		Col8RightMarginStr: "",
		LineTerminator:     lineTerminator,
		MaxLineLength:      maxLineLength,
		isValid:            true,
		lock:               nil,
	}

	lenStdTxtLineCol :=
		len(txtFmtCollection.stdTextLineParamCollection)

	foundStdTxtLineColFmt := false

	if lenStdTxtLineCol > 0 {

		for i := 0; i < lenStdTxtLineCol; i++ {

			if txtFmtCollection.stdTextLineParamCollection[i].
				FormatType == TxtFieldType.Line1Column() {

				txtFmtCollection.stdTextLineParamCollection[i].
					CopyIn(&newStdFmtParams)

				foundStdTxtLineColFmt = true

			}

		}
	}

	if !foundStdTxtLineColFmt {

		txtFmtCollection.stdTextLineParamCollection =
			append(
				txtFmtCollection.stdTextLineParamCollection,
				newStdFmtParams)

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
			Col1FieldText: col1FieldText,
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

// SetStdFormatParamsLine1Col - Sets the standard format parameters
// for a text line consisting of one text column (Line1Column).
// This standard format will be applied as the default format of
// all 'Line1Column' Text Format Operations.
//
func (txtFmtCollection *TextFormatterCollection) SetStdFormatParamsLine1Col(
	leftMarginStr string,
	fieldLength int,
	fieldJustify TextJustify,
	rightMarginStr string,
	lineTerminator string,
	maxLineLength int,
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

	if fieldLength < -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fieldLength' is invalid!\n"+
			"'fieldLength' has a value less than minus one (-1)\n"+
			"fieldLength = '%v'\n",
			ePrefix.String(),
			fieldLength)

		return err
	}

	if !fieldJustify.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fieldJustify' is invalid!\n"+
			"'fieldJustify' must be set to 'Left', 'Right' or 'Center'.\n"+
			"fieldJustify String Value  = '%v'\n"+
			"fieldJustify Integer Value = '%v'\n",
			ePrefix.String(),
			fieldJustify.String(),
			fieldJustify.XValueInt())

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
		FormatType:         TxtFieldType.Line1Column(),
		Col1LeftMarginStr:  leftMarginStr,
		Col1FieldLength:    fieldLength,
		Col1FieldJustify:   fieldJustify,
		Col1RightMarginStr: rightMarginStr,
		Col2LeftMarginStr:  "",
		Col2FieldLength:    0,
		Col2FieldJustify:   0,
		Col2RightMarginStr: "",
		Col3LeftMarginStr:  "",
		Col3FieldLength:    0,
		Col3FieldJustify:   0,
		Col3RightMarginStr: "",
		Col4LeftMarginStr:  "",
		Col4FieldLength:    0,
		Col4FieldJustify:   0,
		Col4RightMarginStr: "",
		Col5LeftMarginStr:  "",
		Col5FieldLength:    0,
		Col5FieldJustify:   0,
		Col5RightMarginStr: "",
		Col6LeftMarginStr:  "",
		Col6FieldLength:    0,
		Col6FieldJustify:   0,
		Col6RightMarginStr: "",
		Col7LeftMarginStr:  "",
		Col7FieldLength:    0,
		Col7FieldJustify:   0,
		Col7RightMarginStr: "",
		Col8LeftMarginStr:  "",
		Col8FieldLength:    0,
		Col8FieldJustify:   0,
		Col8RightMarginStr: "",
		LineTerminator:     lineTerminator,
		MaxLineLength:      maxLineLength,
		isValid:            true,
		lock:               nil,
	}

	lenStdTextLineCol :=
		len(txtFmtCollection.stdTextLineParamCollection)

	for i := 0; i < lenStdTextLineCol; i++ {

		if txtFmtCollection.
			stdTextLineParamCollection[i].FormatType ==
			TxtFieldType.Line1Column() {

			txtFmtCollection.
				stdTextLineParamCollection[i].
				CopyIn(&newStdFmtParams)

			return err
		}
	}

	txtFmtCollection.stdTextLineParamCollection =
		append(txtFmtCollection.stdTextLineParamCollection,
			newStdFmtParams)

	return err
}
