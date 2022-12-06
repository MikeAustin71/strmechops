package strmech

import "sync"

// TextFieldFormatDto
//
// Used to specify content and format specifications for
// a Text Field formatted for screen display, file
// output and printing.
//
// This Data Transfer Object (Dto) contains all the format
// parameters necessary format a single text field.
type TextFieldFormatDto struct {
	LeftMarginStr string
	//	One or more characters used to create a left
	//	margin for this Text Field.
	//
	//	If this parameter is set to an empty string, no
	//	left margin will be configured for this Text
	//	Field.

	FieldContents interface{}
	//	This parameter may contain one of several
	//	specific data types. This empty interface type
	//	will be converted to a string and configured as
	//	the text column content within a text line.
	//
	//	Supported types which may be submitted through
	//	this empty interface parameter are listed as
	//	follows:
	//
	//	   time.Time (Converted using default format)
	//	   string
	//	   bool
	//	   uint, uint8, uint16, uint32, uint64,
	//	   int, int8, int16, int32, int64
	//	   float32, float64
	//	   *big.Int *big.Float
	//	   fmt.Stringer (types that support this interface)
	//	   TextInputParamFieldDateTimeDto
	//	         (Converts date time to string. The best way
	//	          to transmit and configure date time values.)

	FieldLength int
	//	The length of the text field in which the
	//	'FieldContents' will be displayed. If
	//	'FieldLength' is less than the length of the
	//	'FieldContents' string, it will be automatically
	//	set equal to the 'FieldContents' string length.
	//
	//	To automatically set the value of 'FieldLength'
	//	to the length of 'FieldContents', set this
	//	parameter to a value of minus one (-1).
	//
	//	If this parameter is submitted with a value less
	//	than minus one (-1) or greater than 1-million
	//	(1,000,000), an error will be returned.

	FieldJustify TextJustify
	//	An enumeration which specifies the justification
	//	of the 'FieldContents' string within the text
	//	field length specified by 'FieldLength'.
	//
	//	Text justification can only be evaluated in the
	//	context of a text label ('FieldContents'), field
	//	length ('FieldLength') and a Text Justification
	//	object of type TextJustify. This is because text
	//	labels with a field length equal to or less than
	//	the length of the text label string will never
	//	use text justification. In these cases, text
	//	justification is completely ignored.
	//
	//	If the field length is greater than the length of
	//	the text label string, text justification must be
	//	equal to one of these three valid values:
	//
	//	    TextJustify(0).Left()
	//	    TextJustify(0).Right()
	//	    TextJustify(0).Center()
	//
	//	Users can also specify the abbreviated text
	//	justification enumeration syntax as follows:
	//
	//	    TxtJustify.Left()
	//	    TxtJustify.Right()
	//	    TxtJustify.Center()

	RightMarginStr string
	//	One or more characters used to create a right
	//	margin for this Text Field.
	//
	//	If this parameter is set to an empty string, no
	//	right margin will be configured for this Text
	//	Field.

	lock *sync.Mutex
}
