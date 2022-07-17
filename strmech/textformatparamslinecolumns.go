package strmech

import "sync"

// TextFmtParamsLineColumns - Contains Text Format Parameters for a
// Line containing one or more columns of text.
type TextFmtParamsLineColumns struct {
	FormatType TextFieldType
	// Required. This enumeration value specifies the type of Text
	// Format/Parameter Operation to be performed. For the
	// Line1Column Parameter Type, this value is always set to:
	//   TxtFieldType.Line1Column()

	Col1LeftMarginStr  string
	Col1FieldLength    int
	Col1FieldJustify   TextJustify
	Col1RightMarginStr string
	Col2LeftMarginStr  string
	Col2FieldLength    int
	Col2FieldJustify   TextJustify
	Col2RightMarginStr string
	Col3LeftMarginStr  string
	Col3FieldLength    int
	Col3FieldJustify   TextJustify
	Col3RightMarginStr string
	Col4LeftMarginStr  string
	Col4FieldLength    int
	Col4FieldJustify   TextJustify
	Col4RightMarginStr string
	Col5LeftMarginStr  string
	Col5FieldLength    int
	Col5FieldJustify   TextJustify
	Col5RightMarginStr string
	Col6LeftMarginStr  string
	Col6FieldLength    int
	Col6FieldJustify   TextJustify
	Col6RightMarginStr string
	Col7LeftMarginStr  string
	Col7FieldLength    int
	Col7FieldJustify   TextJustify
	Col7RightMarginStr string
	Col8LeftMarginStr  string
	Col8FieldLength    int
	Col8FieldJustify   TextJustify
	Col8RightMarginStr string

	TurnLineTerminationOff bool
	// If this parameter is set to 'true' no Line Termination
	// Sequence will be applied for this text line.

	LineTerminator string
	// If this parameter is submitted as an empty string,
	// the default new line terminator ('\n') will be applied.
	// If this parameter is populated, this character sequence
	// will be used as the Line Terminator for this text line.
	// Remember that the line termination operation will only
	// be performed if paramter 'TurnLineTerminationOff' is set
	// to 'false'.

	MaxLineLength int
	// Set this parameter to minus one -1 to specify an unlimited
	// line length for this text line.

	TurnAutoLineLengthBreaksOn bool
	// When this parameter is set to 'true', text fields which extend
	// beyond the maximum line length 'MaxLineLength' will be placed
	// on the following line of text.

	isValid bool

	lock *sync.Mutex
}

func (paramsLineCol *TextFmtParamsLineColumns) CopyIn(
	incomingParams *TextFmtParamsLineColumns) {

	if paramsLineCol.lock == nil {
		paramsLineCol.lock = new(sync.Mutex)
	}

	paramsLineCol.lock.Lock()

	defer paramsLineCol.lock.Unlock()

	paramsLineCol.FormatType = incomingParams.FormatType
	paramsLineCol.Col1LeftMarginStr = incomingParams.Col1LeftMarginStr
	paramsLineCol.Col1FieldLength = incomingParams.Col1FieldLength
	paramsLineCol.Col1FieldJustify = incomingParams.Col1FieldJustify
	paramsLineCol.Col1RightMarginStr = incomingParams.Col1RightMarginStr
	paramsLineCol.Col2LeftMarginStr = incomingParams.Col2LeftMarginStr
	paramsLineCol.Col2FieldLength = incomingParams.Col2FieldLength
	paramsLineCol.Col2FieldJustify = incomingParams.Col2FieldJustify
	paramsLineCol.Col2RightMarginStr = incomingParams.Col2RightMarginStr
	paramsLineCol.Col3FieldLength = incomingParams.Col3FieldLength
	paramsLineCol.Col3FieldJustify = incomingParams.Col3FieldJustify
	paramsLineCol.Col3RightMarginStr = incomingParams.Col3RightMarginStr
	paramsLineCol.Col4LeftMarginStr = incomingParams.Col4LeftMarginStr
	paramsLineCol.Col4FieldLength = incomingParams.Col4FieldLength
	paramsLineCol.Col4FieldJustify = incomingParams.Col4FieldJustify
	paramsLineCol.Col4RightMarginStr = incomingParams.Col4RightMarginStr
	paramsLineCol.Col5LeftMarginStr = incomingParams.Col5LeftMarginStr
	paramsLineCol.Col5FieldLength = incomingParams.Col5FieldLength
	paramsLineCol.Col5FieldJustify = incomingParams.Col5FieldJustify
	paramsLineCol.Col5RightMarginStr = incomingParams.Col5RightMarginStr
	paramsLineCol.Col6LeftMarginStr = incomingParams.Col6LeftMarginStr
	paramsLineCol.Col6FieldLength = incomingParams.Col6FieldLength
	paramsLineCol.Col6FieldJustify = incomingParams.Col6FieldJustify
	paramsLineCol.Col6RightMarginStr = incomingParams.Col6RightMarginStr
	paramsLineCol.Col7LeftMarginStr = incomingParams.Col7LeftMarginStr
	paramsLineCol.Col7FieldLength = incomingParams.Col7FieldLength
	paramsLineCol.Col7FieldJustify = incomingParams.Col7FieldJustify
	paramsLineCol.Col7RightMarginStr = incomingParams.Col7RightMarginStr
	paramsLineCol.Col8LeftMarginStr = incomingParams.Col8LeftMarginStr
	paramsLineCol.Col8FieldLength = incomingParams.Col8FieldLength
	paramsLineCol.Col8FieldJustify = incomingParams.Col8FieldJustify
	paramsLineCol.Col8RightMarginStr = incomingParams.Col8RightMarginStr
	paramsLineCol.TurnLineTerminationOff = incomingParams.TurnLineTerminationOff
	paramsLineCol.LineTerminator = incomingParams.LineTerminator
	paramsLineCol.MaxLineLength = incomingParams.MaxLineLength
	paramsLineCol.TurnAutoLineLengthBreaksOn =
		incomingParams.TurnAutoLineLengthBreaksOn
	paramsLineCol.isValid = incomingParams.isValid

	return
}

// CopyOut - Returns a deep copy of the TextFmtParamsLineColumns
// current instance.
func (paramsLineCol *TextFmtParamsLineColumns) CopyOut() TextFmtParamsLineColumns {

	if paramsLineCol.lock == nil {
		paramsLineCol.lock = new(sync.Mutex)
	}

	paramsLineCol.lock.Lock()

	defer paramsLineCol.lock.Unlock()

	newParamsLineCols := TextFmtParamsLineColumns{
		FormatType:                 paramsLineCol.FormatType,
		Col1LeftMarginStr:          paramsLineCol.Col1LeftMarginStr,
		Col1FieldLength:            paramsLineCol.Col1FieldLength,
		Col1FieldJustify:           paramsLineCol.Col1FieldJustify,
		Col1RightMarginStr:         paramsLineCol.Col1RightMarginStr,
		Col2LeftMarginStr:          paramsLineCol.Col2LeftMarginStr,
		Col2FieldLength:            paramsLineCol.Col2FieldLength,
		Col2FieldJustify:           paramsLineCol.Col2FieldJustify,
		Col2RightMarginStr:         paramsLineCol.Col2RightMarginStr,
		Col3FieldLength:            paramsLineCol.Col3FieldLength,
		Col3FieldJustify:           paramsLineCol.Col3FieldJustify,
		Col3RightMarginStr:         paramsLineCol.Col3RightMarginStr,
		Col4LeftMarginStr:          paramsLineCol.Col4LeftMarginStr,
		Col4FieldLength:            paramsLineCol.Col4FieldLength,
		Col4FieldJustify:           paramsLineCol.Col4FieldJustify,
		Col4RightMarginStr:         paramsLineCol.Col4RightMarginStr,
		Col5LeftMarginStr:          paramsLineCol.Col5LeftMarginStr,
		Col5FieldLength:            paramsLineCol.Col5FieldLength,
		Col5FieldJustify:           paramsLineCol.Col5FieldJustify,
		Col5RightMarginStr:         paramsLineCol.Col5RightMarginStr,
		Col6LeftMarginStr:          paramsLineCol.Col6LeftMarginStr,
		Col6FieldLength:            paramsLineCol.Col6FieldLength,
		Col6FieldJustify:           paramsLineCol.Col6FieldJustify,
		Col6RightMarginStr:         paramsLineCol.Col6RightMarginStr,
		Col7LeftMarginStr:          paramsLineCol.Col7LeftMarginStr,
		Col7FieldLength:            paramsLineCol.Col7FieldLength,
		Col7FieldJustify:           paramsLineCol.Col7FieldJustify,
		Col7RightMarginStr:         paramsLineCol.Col7RightMarginStr,
		Col8LeftMarginStr:          paramsLineCol.Col8LeftMarginStr,
		Col8FieldLength:            paramsLineCol.Col8FieldLength,
		Col8FieldJustify:           paramsLineCol.Col8FieldJustify,
		Col8RightMarginStr:         paramsLineCol.Col8RightMarginStr,
		TurnLineTerminationOff:     paramsLineCol.TurnLineTerminationOff,
		LineTerminator:             paramsLineCol.LineTerminator,
		MaxLineLength:              paramsLineCol.MaxLineLength,
		TurnAutoLineLengthBreaksOn: paramsLineCol.TurnAutoLineLengthBreaksOn,
		isValid:                    paramsLineCol.isValid,
		lock:                       nil,
	}

	return newParamsLineCols
}

// GetTextParamsType - Returns the Text Field Type associated with
// Line 1-Column Format Parameters.
func (paramsLineCol *TextFmtParamsLineColumns) GetTextParamsType() TextFieldType {

	if paramsLineCol.lock == nil {
		paramsLineCol.lock = new(sync.Mutex)
	}

	paramsLineCol.lock.Lock()

	defer paramsLineCol.lock.Unlock()

	if paramsLineCol.FormatType != TxtFieldType.Line1Column() {

		paramsLineCol.FormatType = TxtFieldType.Line1Column()
	}

	return paramsLineCol.FormatType

}
