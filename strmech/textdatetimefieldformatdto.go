package strmech

import (
	"sync"
	"time"
)

// TextDateFieldFormatDto
//
// The TextDateFieldFormatDto type encapsulates input
// specifications for a text field populated with a
// formatted Date/Time string created from a type
// time.Time.
//
// Field Format Data Transfer Objects (Dto) are used
// to facilitate easy data entry which creating and
// configuring text lines strings for screen display,
// file output or printing.
type TextDateFieldFormatDto struct {
	LeftMarginStr string
	//	One or more characters used to create a left
	//	margin for this 'FieldDateTime' Text Field.
	//
	//	If this parameter is set to an empty string, no
	//	left margin will be configured for this
	//	'FieldDateTime' Text Field.

	FieldDateTime time.Time
	// This time value will be used to populate a Text
	// Field used for screen display, file output or
	// printing.

	FieldDateTimeFormat string
	// This string will be used to format the date time
	// value contained in the 'FieldDateTime' data
	// element.
	//
	// If 'FieldDateTime' is set to a value greater than
	// zero and this 'FieldDateTimeFormat' string is
	// empty (has a zero string length), a default
	// Date/Time format string will be applied as
	// follows:
	//     "2006-01-02 15:04:05.000000000 -0700 MST"

	FieldLength int
	//	The length of the text field in which the
	//	'FieldDateTime' string will be displayed. If
	//	'FieldLength' is less than the length of the
	//	'FieldDateTime' string, it will be automatically
	//	set equal to the 'FieldDateTime' string length.
	//
	//	To automatically set the value of 'FieldLength'
	//	to the length of 'FieldDateTime', set this
	//	parameter to a value of minus one (-1).
	//
	//	If this parameter is submitted with a value less
	//	than minus one (-1) or greater than 1-million
	//	(1,000,000), an error will be returned.
	//
	//	Field Length Examples
	//
	//		Example-1
	//          Date/Time String = "2006-01-02 15:04:05.000000000 -0700 MST"
	//			Date/Time String Length = 39
	//			FieldLength = 45
	//			FieldJustify = TxtJustify.Center()
	//			Text Field String =
	//				"   2006-01-02 15:04:05.000000000 -0700 MST   "
	//
	//		Example-2
	//          Date/Time String = "2006-01-02 15:04:05.000000000 -0700 MST"
	//			Date/Time String Length = 39
	//			FieldLength = 45
	//			FieldJustify = TxtJustify.Left()
	//			Text Field String =
	//				"2006-01-02 15:04:05.000000000 -0700 MST      "
	//
	//		Example-3
	//          Date/Time String = "2006-01-02 15:04:05.000000000 -0700 MST"
	//			Date/Time String Length = 39
	//			FieldLength = -1
	//			FieldJustify = TxtJustify.Center() // Ignored
	//			Text Field String =
	//				"2006-01-02 15:04:05.000000000 -0700 MST"

	FieldJustify TextJustify
	//	An enumeration which specifies the justification
	//	of the 'FieldDateTime' string within the text
	//	field length specified by 'FieldLength'.
	//
	//	Text justification can only be evaluated in the
	//	context of a text label ('FieldDateTime'), field
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
	//
	//	Text Justification Examples
	//
	//		Example-1
	//          Date/Time String = "2006-01-02 15:04:05.000000000 -0700 MST"
	//			Date/Time String Length = 39
	//			FieldLength = 45
	//			FieldJustify = TxtJustify.Center()
	//			Text Field String =
	//				"   2006-01-02 15:04:05.000000000 -0700 MST   "
	//
	//		Example-2
	//          Date/Time String = "2006-01-02 15:04:05.000000000 -0700 MST"
	//			Date/Time String Length = 39
	//			FieldLength = 45
	//			FieldJustify = TxtJustify.Left()
	//			Text Field String =
	//				"2006-01-02 15:04:05.000000000 -0700 MST      "
	//
	//		Example-3
	//          Date/Time String = "2006-01-02 15:04:05.000000000 -0700 MST"
	//			Date/Time String Length = 39
	//			FieldLength = -1
	//			FieldJustify = TxtJustify.Center() // Ignored
	//			Text Field String =
	//				"2006-01-02 15:04:05.000000000 -0700 MST"

	RightMarginStr string
	//	One or more characters used to create a right
	//	margin for this 'FieldDateTime' Text Field.
	//
	//	If this parameter is set to an empty string, no
	//	right margin will be configured for this
	//	'FieldDateTime' Text Field.

	lock *sync.Mutex
}

// textDateFieldFormatDtoAtom - Provides helper
// methods for TextDateFieldFormatDto.
type textDateFieldFormatDtoAtom struct {
	lock *sync.Mutex
}

// empty
//
// Receives a pointer to an instance of
// TextDateFieldFormatDto and proceeds to set all the
// member variables to their zero or uninitialized
// states.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and reset all pre-existing
//	data values contained within the
//	TextDateFieldFormatDto instance passed as input
//	parameter 'txtDateFieldDto' to their zero or
//	uninitialized states.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	txtDateFieldDto				*TextDateFieldFormatDto
//
//		A pointer to an instance of TextDateFieldFormatDto.
//		All data values contained within this instance
//		will be deleted and reset to their zero or
//		uninitialized states.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	NONE
func (txtDateFieldDtoAtom *textDateFieldFormatDtoAtom) empty(
	txtDateFieldDto *TextDateFieldFormatDto) {

	if txtDateFieldDtoAtom.lock == nil {
		txtDateFieldDtoAtom.lock = new(sync.Mutex)
	}

	txtDateFieldDtoAtom.lock.Lock()

	defer txtDateFieldDtoAtom.lock.Unlock()

	if txtDateFieldDto == nil {

		return
	}

	txtDateFieldDto.LeftMarginStr = ""

	txtDateFieldDto.FieldDateTime = time.Time{}

	txtDateFieldDto.FieldDateTimeFormat = ""

	txtDateFieldDto.FieldLength = 0

	txtDateFieldDto.FieldJustify = TxtJustify.None()

	txtDateFieldDto.RightMarginStr = ""

	return
}
