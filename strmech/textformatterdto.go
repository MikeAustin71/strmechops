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
	//  TxtFieldType.Label()
	//  TxtFieldType.DateTime()
	//  TxtFieldType.Filler()
	//  TxtFieldType.Spacer()
	//  TxtFieldType.BlankLine()
	//  TxtFieldType.SolidLine()
	//  TxtFieldType.Line1Column()
	//  TxtFieldType.Line2Column()
	//  TxtFieldType.Line3Column()
	//  TxtFieldType.Line4Column()
	//  TxtFieldType.Line5Column()
	//  TxtFieldType.Line6Column()
	//  TxtFieldType.Line7Column()
	//  TxtFieldType.Line8Column()

	DateTime TextFieldDateTimeDto
	// A structure containing data elements necessary for the
	// creation of Text Date/Time Fields.

	Filler TextFieldFillerDto
	// A structure containing data elements necessary for the
	// creation of Text Filler Fields.

	Label TextFieldLabelDto
	// A structure containing data elements necessary for the
	// creation of Text Label Fields.

	Spacer TextFieldSpacerDto
	// A structure containing data elements necessary for the
	// creation of Text Spacer Fields.

	BlankLine TextLineBlankDto
	// A structure containing data elements necessary for the
	// creation of Blank Lines or New Lines.

	SolidLine TextLineSolidLineDto
	// A structure containing data elements necessary for the
	// creation of a solid line of text characters.
	//
	// A solid line, as used here, consists of a single character
	// or multiple characters used in a repeating sequence to
	// construct a solid line. Typically, solid lines consist of
	// dashes ("---"), underscore characters ("____"), equal signs
	// ("====="), asterisks ("*****") and other similar line break
	// presentations. The length of a solid line is specified by the
	// specification parameter, 'SolidLineCharsRepeatCount'.

	LineColumns TextLineColumnsDto
	// A structure containing data elements necessary for the
	// creation of a text line containing at least one, and up
	// to 8-columns. TextFieldType's covered by 'LineColumns'
	// are listed as follows:
	//  TxtFieldType.Line1Column()
	//  TxtFieldType.Line2Column()
	//  TxtFieldType.Line3Column()
	//  TxtFieldType.Line4Column()
	//  TxtFieldType.Line5Column()
	//  TxtFieldType.Line6Column()
	//  TxtFieldType.Line7Column()
	//  TxtFieldType.Line8Column()

}
