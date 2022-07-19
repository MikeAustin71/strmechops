package strmech

import "time"

type TextInputParamFieldDateTimeDto struct {
	FieldDateTime time.Time
	// This time value will be used to populate a Text Field used
	// for screen display, file output or printing.

	FieldDateTimeFormat string
	// This string will be used to format the date time value
	// contained in the 'FieldDateTime' data element.
	//
	// If 'FieldDateTime' is set to a value greater than zero and
	// this 'FieldDateTimeFormat' string is empty (has a zero
	// length), a default Date/Time format string will be applied
	// as follows:
	//     "2006-01-02 15:04:05.000000000 -0700 MST"

}
