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
// Format Type 'Line1Column' Text. This method will assign
// Format Parameter to this 'Line1Column' Text Line based
// on the Standard Parameters for this Text Type.
//
// This method will extract format parameters from the Standard
// Text Line Parameters collection maintained by this instance of
// TextFormatterCollection.
//
// If the standard parameters for 'Line1Column' Text Lines do
// not exist in the Standard Text Line Parameters Collection,
// an error will be returned.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// The standard parameters for 'Line1Column' Text Lines must be
// configured in the Standard Text Line Parameters Collection
// before calling this method.
//
// ----------------------------------------------------------------
//
// To configure the standard parameters for 'Line1Column' Text
// Lines call one of the following methods:
//   TextFormatterCollection.CfgLine1Col()
//   TextFormatterCollection.SetStdFormatParamsLine1Col()
//
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
			formatType:    TxtFieldType.Line1Column(),
			col1FieldText: column1FieldText,
			col2FieldText: "",
			col3FieldText: "",
			col4FieldText: "",
			col5FieldText: "",
			col6FieldText: "",
			col7FieldText: "",
			col8FieldText: "",
			fmtParameters: stdLineColsFmt,
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
			"Error: Input parameter 'col1FieldLength' is invalid!\n"+
			"'col1FieldLength' has a value less than minus one (-1)\n"+
			"col1FieldLength = '%v'\n",
			ePrefix.String(),
			col1FieldLength)

		return err
	}

	if !col1FieldJustify.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'col1FieldJustify' is invalid!\n"+
			"'col1FieldJustify' must be set to 'Left', 'Right' or 'Center'.\n"+
			"col1FieldJustify String Value  = '%v'\n"+
			"col1FieldJustify Integer Value = '%v'\n",
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
		formatType:         TxtFieldType.Line1Column(),
		col1LeftMarginStr:  leftMarginStr,
		col1FieldLength:    col1FieldLength,
		col1FieldJustify:   col1FieldJustify,
		col1RightMarginStr: rightMarginStr,
		col2LeftMarginStr:  "",
		col2FieldLength:    0,
		col2FieldJustify:   0,
		col2RightMarginStr: "",
		col3LeftMarginStr:  "",
		col3FieldLength:    0,
		col3FieldJustify:   0,
		col3RightMarginStr: "",
		col4LeftMarginStr:  "",
		col4FieldLength:    0,
		col4FieldJustify:   0,
		col4RightMarginStr: "",
		col5LeftMarginStr:  "",
		col5FieldLength:    0,
		col5FieldJustify:   0,
		col5RightMarginStr: "",
		col6LeftMarginStr:  "",
		col6FieldLength:    0,
		col6FieldJustify:   0,
		col6RightMarginStr: "",
		col7LeftMarginStr:  "",
		col7FieldLength:    0,
		col7FieldJustify:   0,
		col7RightMarginStr: "",
		col8LeftMarginStr:  "",
		col8FieldLength:    0,
		col8FieldJustify:   0,
		col8RightMarginStr: "",
		lineTerminator:     lineTerminator,
		maxLineLength:      maxLineLength,
		isValid:            true,
		lock:               nil,
	}

	lenStdTxtLineCol :=
		len(txtFmtCollection.stdTextLineParamCollection)

	foundStdTxtLineColFmt := false

	if lenStdTxtLineCol > 0 {

		for i := 0; i < lenStdTxtLineCol; i++ {

			if txtFmtCollection.stdTextLineParamCollection[i].
				formatType == TxtFieldType.Line1Column() {

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
			formatType:    TxtFieldType.Line1Column(),
			col1FieldText: col1FieldText,
			col2FieldText: "",
			col3FieldText: "",
			col4FieldText: "",
			col5FieldText: "",
			col6FieldText: "",
			col7FieldText: "",
			col8FieldText: "",
			fmtParameters: TextFmtParamsLineColumns{},
			lock:          nil,
		},
	}

	newLine1Col.LineColumns.fmtParameters.CopyIn(
		&newStdFmtParams)

	txtFmtCollection.fmtCollection =
		append(
			txtFmtCollection.fmtCollection,
			newLine1Col)

	return err
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
		formatType:         TxtFieldType.Line1Column(),
		col1LeftMarginStr:  leftMarginStr,
		col1FieldLength:    fieldLength,
		col1FieldJustify:   fieldJustify,
		col1RightMarginStr: rightMarginStr,
		col2LeftMarginStr:  "",
		col2FieldLength:    0,
		col2FieldJustify:   0,
		col2RightMarginStr: "",
		col3LeftMarginStr:  "",
		col3FieldLength:    0,
		col3FieldJustify:   0,
		col3RightMarginStr: "",
		col4LeftMarginStr:  "",
		col4FieldLength:    0,
		col4FieldJustify:   0,
		col4RightMarginStr: "",
		col5LeftMarginStr:  "",
		col5FieldLength:    0,
		col5FieldJustify:   0,
		col5RightMarginStr: "",
		col6LeftMarginStr:  "",
		col6FieldLength:    0,
		col6FieldJustify:   0,
		col6RightMarginStr: "",
		col7LeftMarginStr:  "",
		col7FieldLength:    0,
		col7FieldJustify:   0,
		col7RightMarginStr: "",
		col8LeftMarginStr:  "",
		col8FieldLength:    0,
		col8FieldJustify:   0,
		col8RightMarginStr: "",
		lineTerminator:     lineTerminator,
		maxLineLength:      maxLineLength,
		isValid:            true,
		lock:               nil,
	}

	lenStdTextLineCol :=
		len(txtFmtCollection.stdTextLineParamCollection)

	for i := 0; i < lenStdTextLineCol; i++ {

		if txtFmtCollection.
			stdTextLineParamCollection[i].formatType ==
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
