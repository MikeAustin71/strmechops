package strmech

import "sync"

// NumStrFmtCountryTelephoneNumSpec
//
// Stores and transports data specifications necessary
// for the creation and display of telephone numbers.
//
//	Resources:
//		http://www.wtng.info/
//		https://support.microsoft.com/en-us/office/phone-number-format-afbe9a4f-6f1c-4370-8a93-4e41d27460f5
//		https://countrycode.org/
type NumStrFmtCountryTelephoneNumSpec struct {
	CountryName string
	//	The name of the country associated with this
	//	Telephone data.

	CountryCodeTwoChar string
	//	The unique ISO 3166-1 alpha-2 Two Character code
	//	identifying the country or culture associated
	//	with the current Country Culture Specification
	//	instance.
	//	ISO 3166-1 alpha-2 Wikipedia
	//	https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2

	CountryCodeThreeChar string
	//	Optional
	//	The unique ISO 3166-1 alpha-3 Three Character code
	//	identifying the country or culture associated with
	//	the current Country Culture Specification instance.
	//	ISO 3166-1 alpha-3 Wikipedia
	//	https://en.wikipedia.org/wiki/ISO_3166-1_alpha-3

	InternationalDirectDialingNo string
	//	Usually the same as the InternationalPrefix

	InternationalPrefix string
	//	Lists the prefix used in the country when
	//	dialling an international number.

	TrunkPrefix string
	//	Lists the prefix used for calls within a
	//	country, usually for long distance calls
	//	outside a local area. Increasingly,
	//	some countries are requiring use of trunk
	//	prefix and area code for local calls.

	CountryTelephoneCode string
	//	The Telephone Country Code used to prefix
	//	Telephone Numbers for this country.

	AreaCodeMaxNumOfDigits string
	//	The Maximum Number of Area Code numeric digits
	//	in this country's Telephone Number.
	//
	//	Refers to the number of digits in the area code.
	//	(This does not include Trunk Prefix, if this is
	//	used in the country).
	//	Some nations only use the subscriber number, and
	//	do not use an area code.

	AreaCodeMinNumOfDigits string
	//	The Minimum Number of Area Code numeric digits
	//	in this country's Telephone Number.
	//
	//	Refers to the number of digits in the area code.
	//	(This does not include Trunk Prefix, if this is
	//	used in the country).
	//	Some nations only use the subscriber number, and
	//	do not use an area code.

	SubscriberNumMaxNumOfDigitsExternal string
	//	The Maximum Number of numeric digits in the
	//	Subscriber Number when calling from outside
	//	the host country.

	SubscriberNumMinNumOfDigitsExternal string
	//	The Minimum Number of numeric digits in the
	//	Subscriber Number when calling from outside
	//	the host country.

	SubscriberNumMaxNumOfDigitsInternal string
	//	The Maximum Number of numeric digits in the
	//	Subscriber Number when calling from inside
	//	the host country.

	SubscriberNumMinNumOfDigitsInternal string
	//	The Minimum Number of numeric digits in the
	//	Subscriber Number when calling from inside
	//	the host country.

	MobileNumMaxNumOfDigitsExternal string
	//	The Maximum Number of numeric digits in the
	//	Mobile Number when calling from outside
	//	the host country.

	MobileNumMinNumOfDigitsExternal string
	//	The Minimum Number of numeric digits in the
	//	Mobile Number when calling from outside
	//	the host country.

	MobileNumMaxNumOfDigitsInternal string
	//	The Maximum Number of numeric digits in the
	//	Mobile Number when calling from inside
	//	the host country.

	MobileNumMinNumOfDigitsInternal string
	//	The Minimum Number of numeric digits in the
	//	Mobile Number when calling from inside
	//	the host country.

	PhoneExtNumMaxNumOfDigitsExternal string
	//	The Maximum Number of numeric digits in the
	//	Phone Extension Number when calling from outside
	//	the host country.

	PhoneExtNumMinNumOfDigitsExternal string
	//	The Minimum Number of numeric digits in the
	//	Phone Extension Number when calling from outside
	//	the host country.

	PhoneExtNumMaxNumOfDigitsInternal string
	//	The Maximum Number of numeric digits in the
	//	Phone Extension Number when calling from inside
	//	the host country.

	PhoneExtNumMinNumOfDigitsInternal string
	//	The Minimum Number of numeric digits in the
	//	Phone Extension Number when calling from inside
	//	the host country.

	// ****** Subscriber Numbers ******

	SubscriberFmtFullExternal NumStrFmtTelephoneNumSpec
	//	This format is used with a full Subscriber phone
	//	number.
	//
	//	This is an 'External' format used when dialing in
	//	from outside the country.
	//
	//	NumStrFmtTelephoneNumSpec contains two formats,
	//	one for dialing and one for text display.
	//
	//	The formats are designed to use character
	//	replacement for number display. The character 'N'
	//	is used as a placeholder for numeric digits.
	//
	//	Example US: "1 (NNN) NNN-NNNN"

	SubscriberFmtAbbrExternal NumStrFmtTelephoneNumSpec
	//	This format is used with an abbreviated
	//	Subscriber phone number.
	//
	//	This is an 'External' format used when dialing in
	//	from outside the country.
	//
	//	NumStrFmtTelephoneNumSpec contains two formats,
	//	one for dialing and one for text display.
	//
	//	The formats are designed to use character
	//	replacement for number display. The character 'N'
	//	is used as a placeholder for numeric digits.
	//
	//	Example US: "1 (NNN) NNN-NNNN"

	SubscriberFmtFullInternal NumStrFmtTelephoneNumSpec
	//	This format is used with a full Subscriber phone
	//	number.
	//
	//	This is an 'Internal' format used when dialing in
	//	from outside the country.
	//
	//	NumStrFmtTelephoneNumSpec contains two formats,
	//	one for dialing and one for text display.
	//
	//	The formats are designed to use character
	//	replacement for number display. The character 'N'
	//	is used as a placeholder for numeric digits.
	//
	//	Example US: "1 (NNN) NNN-NNNN"

	SubscriberFmtAbbrInternal NumStrFmtTelephoneNumSpec
	//	This format is used with an abbreviated
	//	Subscriber phone number.
	//
	//	This is an 'Internal' format used when dialing in
	//	from outside the country.
	//
	//	NumStrFmtTelephoneNumSpec contains two formats,
	//	one for dialing and one for text display.
	//
	//	The formats are designed to use character
	//	replacement for number display. The character 'N'
	//	is used as a placeholder for numeric digits.
	//
	//	Example US: "1 (NNN) NNN-NNNN"

	// ****** END Of Subscriber Numbers ******

	// ****** Mobile Numbers ******

	MobileFmtFullExternal NumStrFmtTelephoneNumSpec
	//	This format is used with a full Mobile phone
	//	number.
	//
	//	This is an 'External' format used when dialing in
	//	from outside the country.
	//
	//	NumStrFmtTelephoneNumSpec contains two formats,
	//	one for dialing and one for text display.
	//
	//	The formats are designed to use character
	//	replacement for number display. The character 'N'
	//	is used as a placeholder for numeric digits.
	//
	//	Example US: "1 (NNN) NNN-NNNN"

	MobileFmtAbbrExternal NumStrFmtTelephoneNumSpec
	//	This format is used with an abbreviated
	//	Mobile phone number.
	//
	//	This is an 'External' format used when dialing in
	//	from outside the country.
	//
	//	NumStrFmtTelephoneNumSpec contains two formats,
	//	one for dialing and one for text display.
	//
	//	The formats are designed to use character
	//	replacement for number display. The character 'N'
	//	is used as a placeholder for numeric digits.
	//
	//	Example US: "1 (NNN) NNN-NNNN"

	MobileFmtFullInternal NumStrFmtTelephoneNumSpec
	//	This format is used with a full Mobile phone
	//	number.
	//
	//	This is an 'Internal' format used when dialing in
	//	from outside the country.
	//
	//	NumStrFmtTelephoneNumSpec contains two formats,
	//	one for dialing and one for text display.
	//
	//	The formats are designed to use character
	//	replacement for number display. The character 'N'
	//	is used as a placeholder for numeric digits.
	//
	//	Example US: "1 (NNN) NNN-NNNN"

	MobileFmtAbbrInternal NumStrFmtTelephoneNumSpec
	//	This format is used with an abbreviated
	//	Mobile phone number.
	//
	//	This is an 'Internal' format used when dialing in
	//	from outside the country.
	//
	//	NumStrFmtTelephoneNumSpec contains two formats,
	//	one for dialing and one for text display.
	//
	//	The formats are designed to use character
	//	replacement for number display. The character 'N'
	//	is used as a placeholder for numeric digits.
	//
	//	Example US: "1 (NNN) NNN-NNNN"

	// ****** END Of Mobile Numbers ******

	// ****** Extension Numbers ******

	PhoneExtFmtFullExternal NumStrFmtTelephoneNumSpec
	//	This format is used with a full Phone Extension
	//	phone number.
	//
	//	This is an 'External' format used when dialing in
	//	from outside the country.
	//
	//	NumStrFmtTelephoneNumSpec contains two formats,
	//	one for dialing and one for text display.
	//
	//	The formats are designed to use character
	//	replacement for number display. The character 'N'
	//	is used as a placeholder for numeric digits.
	//
	//	Example US: "1 (NNN) NNN-NNNN"

	PhoneExtFmtAbbrExternal NumStrFmtTelephoneNumSpec
	//	This format is used with an abbreviated
	//	Phone Extension phone number.
	//
	//	This is an 'External' format used when dialing in
	//	from outside the country.
	//
	//	NumStrFmtTelephoneNumSpec contains two formats,
	//	one for dialing and one for text display.
	//
	//	The formats are designed to use character
	//	replacement for number display. The character 'N'
	//	is used as a placeholder for numeric digits.
	//
	//	Example US: "1 (NNN) NNN-NNNN"

	PhoneExtFmtFullInternal NumStrFmtTelephoneNumSpec
	//	This format is used with a full Phone Extension
	//	phone number.
	//
	//	This is an 'Internal' format used when dialing in
	//	from outside the country.
	//
	//	NumStrFmtTelephoneNumSpec contains two formats,
	//	one for dialing and one for text display.
	//
	//	The formats are designed to use character
	//	replacement for number display. The character 'N'
	//	is used as a placeholder for numeric digits.
	//
	//	Example US: "1 (NNN) NNN-NNNN"

	PhoneExtFmtAbbrInternal NumStrFmtTelephoneNumSpec
	//	This format is used with an abbreviated
	//	Phone Extension phone number.
	//
	//	This is an 'Internal' format used when dialing in
	//	from outside the country.
	//
	//	NumStrFmtTelephoneNumSpec contains two formats,
	//	one for dialing and one for text display.
	//
	//	The formats are designed to use character
	//	replacement for number display. The character 'N'
	//	is used as a placeholder for numeric digits.
	//
	//	Example US: "1 (NNN) NNN-NNNN"

	// ****** END Of Extension Numbers ******

	lock *sync.Mutex
}
