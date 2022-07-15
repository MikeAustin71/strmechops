package strmech

import (
	"sync"
)

type TextLineColumnsDto struct {
	FormatType TextFieldType
	// Required. This enumeration value specifies the type of Text
	// Format Operation to be performed. For the Line1Column
	// Formatter Type, this value is always set to:
	//   TxtFieldType.Line1Column()

	Col1FieldText string
	Col2FieldText string
	Col3FieldText string
	Col4FieldText string
	Col5FieldText string
	Col6FieldText string
	Col7FieldText string
	Col8FieldText string

	FmtParameters TextFmtParamsLineColumns

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

	if fmtLine1Col.FormatType != TxtFieldType.Line1Column() {

		fmtLine1Col.FormatType = TxtFieldType.Line1Column()
	}

	return fmtLine1Col.FormatType
}
