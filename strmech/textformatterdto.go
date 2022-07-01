package strmech

type TextFormatterDto struct {
	FormatType TextFieldType
	// Required. This enumeration value specifies the type of Text
	// Format Operation to be performed. Depending on this value,
	// other member variables in the type will be accessed and used
	//
	// generated.
	//
	// Possible values are listed as follows:
	//   TxtFieldType.None()      - Invalid
	//   TxtFieldType.Label()     - Valid
	//   TxtFieldType.DateTime()  - Valid
	//   TxtFieldType.Filler()    - Valid
	//   TxtFieldType.Spacer()    - Valid
	//   TxtFieldType.BlankLine() - Valid

	TxtDateTime TextFieldDateTimeDto
	// A structure containing data elements necessary for the
	// creation of Text Date/Time Fields.

	TxtFiller TextFieldFillerDto
	// A structure containing data elements necessary for the
	// creation of Text Filler Fields.

	TxtLabel TextFieldLabelDto
	// A structure containing data elements necessary for the
	// creation of Text Label Fields.

	TxtSpacer TextFieldSpacerDto
	// A structure containing data elements necessary for the
	// creation of Text Spacer Fields.

	TxtBlankLine TextLineBlankDto
	// A structure containing data elements necessary for the
	// creation of Blank Lines or New Lines.
}
