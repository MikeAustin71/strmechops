package strmech

import "time"

type TextInputParamFieldDateTimeDto struct {
	FieldDateTime time.Time
	// If this Text Field is designated as a 'DateTime' Field, this
	// time value will be used to populate the Text Field.

	FieldDateTimeFormat string
	// If this Text Field is designated as a 'DateTime' Field, this
	// string will be used to format the Date/Time.
	//
	// If 'FieldDateTime' is set to a value greater than zero and
	// this 'ParamValueDateTimeFormat' string is empty (has a zero
	// length), a default Date/Time format string will be applied
	// as follows:
	//         "2006-01-02 15:04:05.000000000 -0700 MST"

}
