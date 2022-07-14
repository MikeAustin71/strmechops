package strmech

import "sync"

type TextFormatterLine2Column struct {
	FormatType TextFieldType
	// Required. This enumeration value specifies the type of Text
	// Format Operation to be performed. For the Line1Column
	// Formatter Type, this value is always set to:
	//   TxtFieldType.Line2Column()

	Line1ColumnFieldText string
	Line2ColumnFieldText string

	lock *sync.Mutex
}

type TextFormatParamsLine2Column struct {
	Line1ColumnLeftMarginStr  string
	Line1ColumnFieldLength    int
	Line1ColumnFieldJustify   TextJustify
	Line1ColumnRightMarginStr string
	Line2ColumnRightMarginStr string
	LineTerminator            string

	lock *sync.Mutex
}
