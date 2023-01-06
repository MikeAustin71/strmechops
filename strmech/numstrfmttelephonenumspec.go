package strmech

// NumStrFmtTelephoneNumSpec
//
// Specifies the parameters for presentation of two types
// of telephone numbers, a telephone number used in the
// actual 'dialing' process and a telephone number
// formatted for a text presentation.
type NumStrFmtTelephoneNumSpec struct {
	PhoneNoDialFmt NumStrFmtCharReplacementSpec
	//	The format used when dialing a telephone
	//	number.
	//
	//	The character 'N' is used as a placeholder for
	//	numeric digits.
	//
	//	This is the format used when dialing in from
	//	outside the country.
	//
	//	Example US: "1 NNNNNNNNNN"
	//
	//	Type NumStrFmtCharReplacementSpec contains a
	//	pattern and a designated character used for
	//	replacing said character with a numeric digit
	//	(0-9).

	PhoneNoDisplayFmt NumStrFmtCharReplacementSpec
	//	The format used to display a telephone
	//	number in text.
	//
	//	The character 'N' is used as a placeholder for
	//	numeric digits.
	//
	//	This is the format used when dialing in from
	//	outside the country.
	//
	//	Example US: 1 (NNN) NNN-NNNN
	//
	//	Type NumStrFmtCharReplacementSpec contains a
	//	pattern and a designated character used for
	//	replacing said character with a numeric digit
	//	(0-9).

}
