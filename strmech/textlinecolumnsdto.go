package strmech

import (
	"sync"
)

type TextLineColumnsDto struct {
	formatType TextFieldType
	// Required. This enumeration value specifies the type of Text
	// Format Operation to be performed. For the Line1Column
	// Formatter Type, this value is always set to:
	//   TxtFieldType.Line1Column()

	col1FieldText string
	col2FieldText string
	col3FieldText string
	col4FieldText string
	col5FieldText string
	col6FieldText string
	col7FieldText string
	col8FieldText string

	fmtParameters TextFmtParamsLineColumns

	lock *sync.Mutex
}

// GetTextFieldType - Returns the internal member variable which
// stores the value of TextFieldType. For type
// TextLineColumnsDto, TextFieldType should be set to
// TxtFieldType.Line1Column().
//
func (fmtLine1Col *TextLineColumnsDto) GetTextFieldType() TextFieldType {

	if fmtLine1Col.lock == nil {
		fmtLine1Col.lock = new(sync.Mutex)
	}

	fmtLine1Col.lock.Lock()

	defer fmtLine1Col.lock.Unlock()

	if fmtLine1Col.formatType != TxtFieldType.Line1Column() {

		fmtLine1Col.formatType = TxtFieldType.Line1Column()
	}

	return fmtLine1Col.formatType
}
