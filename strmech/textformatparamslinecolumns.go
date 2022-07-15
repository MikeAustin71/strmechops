package strmech

import "sync"

// TextFmtParamsLineColumns - Contains Text Format Parameters for a
// Line containing one or more columns of text.
type TextFmtParamsLineColumns struct {
	formatType TextFieldType
	// Required. This enumeration value specifies the type of Text
	// Format/Parameter Operation to be performed. For the
	// Line1Column Parameter Type, this value is always set to:
	//   TxtFieldType.Line1Column()

	col1LeftMarginStr  string
	col1FieldLength    int
	col1FieldJustify   TextJustify
	col1RightMarginStr string
	col2LeftMarginStr  string
	col2FieldLength    int
	col2FieldJustify   TextJustify
	col2RightMarginStr string
	col3LeftMarginStr  string
	col3FieldLength    int
	col3FieldJustify   TextJustify
	col3RightMarginStr string
	col4LeftMarginStr  string
	col4FieldLength    int
	col4FieldJustify   TextJustify
	col4RightMarginStr string
	col5LeftMarginStr  string
	col5FieldLength    int
	col5FieldJustify   TextJustify
	col5RightMarginStr string
	col6LeftMarginStr  string
	col6FieldLength    int
	col6FieldJustify   TextJustify
	col6RightMarginStr string
	col7LeftMarginStr  string
	col7FieldLength    int
	col7FieldJustify   TextJustify
	col7RightMarginStr string
	col8LeftMarginStr  string
	col8FieldLength    int
	col8FieldJustify   TextJustify
	col8RightMarginStr string
	lineTerminator     string
	maxLineLength      int
	isValid            bool

	lock *sync.Mutex
}

func (paramsLineCol *TextFmtParamsLineColumns) CopyIn(
	incomingParams *TextFmtParamsLineColumns) {

	if paramsLineCol.lock == nil {
		paramsLineCol.lock = new(sync.Mutex)
	}

	paramsLineCol.lock.Lock()

	defer paramsLineCol.lock.Unlock()

	paramsLineCol.formatType = incomingParams.formatType
	paramsLineCol.col1LeftMarginStr = incomingParams.col1LeftMarginStr
	paramsLineCol.col1FieldLength = incomingParams.col1FieldLength
	paramsLineCol.col1FieldJustify = incomingParams.col1FieldJustify
	paramsLineCol.col1RightMarginStr = incomingParams.col1RightMarginStr
	paramsLineCol.col2LeftMarginStr = incomingParams.col2LeftMarginStr
	paramsLineCol.col2FieldLength = incomingParams.col2FieldLength
	paramsLineCol.col2FieldJustify = incomingParams.col2FieldJustify
	paramsLineCol.col2RightMarginStr = incomingParams.col2RightMarginStr
	paramsLineCol.col3FieldLength = incomingParams.col3FieldLength
	paramsLineCol.col3FieldJustify = incomingParams.col3FieldJustify
	paramsLineCol.col3RightMarginStr = incomingParams.col3RightMarginStr
	paramsLineCol.col4LeftMarginStr = incomingParams.col4LeftMarginStr
	paramsLineCol.col4FieldLength = incomingParams.col4FieldLength
	paramsLineCol.col4FieldJustify = incomingParams.col4FieldJustify
	paramsLineCol.col4RightMarginStr = incomingParams.col4RightMarginStr
	paramsLineCol.col5LeftMarginStr = incomingParams.col5LeftMarginStr
	paramsLineCol.col5FieldLength = incomingParams.col5FieldLength
	paramsLineCol.col5FieldJustify = incomingParams.col5FieldJustify
	paramsLineCol.col5RightMarginStr = incomingParams.col5RightMarginStr
	paramsLineCol.col6LeftMarginStr = incomingParams.col6LeftMarginStr
	paramsLineCol.col6FieldLength = incomingParams.col6FieldLength
	paramsLineCol.col6FieldJustify = incomingParams.col6FieldJustify
	paramsLineCol.col6RightMarginStr = incomingParams.col6RightMarginStr
	paramsLineCol.col7LeftMarginStr = incomingParams.col7LeftMarginStr
	paramsLineCol.col7FieldLength = incomingParams.col7FieldLength
	paramsLineCol.col7FieldJustify = incomingParams.col7FieldJustify
	paramsLineCol.col7RightMarginStr = incomingParams.col7RightMarginStr
	paramsLineCol.col8LeftMarginStr = incomingParams.col8LeftMarginStr
	paramsLineCol.col8FieldLength = incomingParams.col8FieldLength
	paramsLineCol.col8FieldJustify = incomingParams.col8FieldJustify
	paramsLineCol.col8RightMarginStr = incomingParams.col8RightMarginStr
	paramsLineCol.lineTerminator = incomingParams.lineTerminator
	paramsLineCol.maxLineLength = incomingParams.maxLineLength
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
		formatType:         paramsLineCol.formatType,
		col1LeftMarginStr:  paramsLineCol.col1LeftMarginStr,
		col1FieldLength:    paramsLineCol.col1FieldLength,
		col1FieldJustify:   paramsLineCol.col1FieldJustify,
		col1RightMarginStr: paramsLineCol.col1RightMarginStr,
		col2LeftMarginStr:  paramsLineCol.col2LeftMarginStr,
		col2FieldLength:    paramsLineCol.col2FieldLength,
		col2FieldJustify:   paramsLineCol.col2FieldJustify,
		col2RightMarginStr: paramsLineCol.col2RightMarginStr,
		col3FieldLength:    paramsLineCol.col3FieldLength,
		col3FieldJustify:   paramsLineCol.col3FieldJustify,
		col3RightMarginStr: paramsLineCol.col3RightMarginStr,
		col4LeftMarginStr:  paramsLineCol.col4LeftMarginStr,
		col4FieldLength:    paramsLineCol.col4FieldLength,
		col4FieldJustify:   paramsLineCol.col4FieldJustify,
		col4RightMarginStr: paramsLineCol.col4RightMarginStr,
		col5LeftMarginStr:  paramsLineCol.col5LeftMarginStr,
		col5FieldLength:    paramsLineCol.col5FieldLength,
		col5FieldJustify:   paramsLineCol.col5FieldJustify,
		col5RightMarginStr: paramsLineCol.col5RightMarginStr,
		col6LeftMarginStr:  paramsLineCol.col6LeftMarginStr,
		col6FieldLength:    paramsLineCol.col6FieldLength,
		col6FieldJustify:   paramsLineCol.col6FieldJustify,
		col6RightMarginStr: paramsLineCol.col6RightMarginStr,
		col7LeftMarginStr:  paramsLineCol.col7LeftMarginStr,
		col7FieldLength:    paramsLineCol.col7FieldLength,
		col7FieldJustify:   paramsLineCol.col7FieldJustify,
		col7RightMarginStr: paramsLineCol.col7RightMarginStr,
		col8LeftMarginStr:  paramsLineCol.col8LeftMarginStr,
		col8FieldLength:    paramsLineCol.col8FieldLength,
		col8FieldJustify:   paramsLineCol.col8FieldJustify,
		col8RightMarginStr: paramsLineCol.col8RightMarginStr,
		lineTerminator:     paramsLineCol.lineTerminator,
		maxLineLength:      paramsLineCol.maxLineLength,
		isValid:            paramsLineCol.isValid,
		lock:               nil,
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

	if paramsLineCol.formatType != TxtFieldType.Line1Column() {

		paramsLineCol.formatType = TxtFieldType.Line1Column()
	}

	return paramsLineCol.formatType

}
