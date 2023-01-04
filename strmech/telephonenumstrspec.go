package strmech

import "sync"

// TelephoneNumStrSpec
//
// Stores and transports data specifications necessary
// for the creation and display of telephone numbers.
//
//	Resources:
//		http://www.wtng.info/
//		https://support.microsoft.com/en-us/office/phone-number-format-afbe9a4f-6f1c-4370-8a93-4e41d27460f5
//		https://countrycode.org/
type TelephoneNumStrSpec struct {
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

	DialingSubNumFmtFullExternal string
	//	The format used when dialing a Subscriber number.
	//	The character 'N' is used as a placeholder for
	//	numeric digits. This is the format used when
	//	dialing in from outside the country.
	//	Example US: "1 NNNNNNNNNN"

	DisplaySubNumFmtFullExternal string
	//	The format used to display a Subscriber phone
	//	number in text.
	//
	//	The character 'N' is used as a placeholder for
	//	numeric digits. This is the format used when
	//	dialing in from outside the country.
	//	Example US: 1 (NNN) NNN-NNNN

	DialingMobileNumFmtFullExternal string
	//	The format used when dialing a Mobile number.
	//	The character 'N' is used as a placeholder for
	//	numeric digits. This is the format used when
	//	dialing in from outside the country.
	//	Example US: "1 NNNNNNNNNN"

	DisplayMobileNumFmtFullExternal string
	//	The format used to display a Mobile phone
	//	number in text.
	//
	//	The character 'N' is used as a placeholder for
	//	numeric digits. This is the format used when
	//	dialing in from outside the country.
	//	Example US: 1 (NNN) NNN-NNNN

	DialingSubNumFmtFullInternal string
	//	The format used when dialing a Subscriber number.
	//	The character 'N' is used as a placeholder for
	//	numeric digits. This is the format used when
	//	dialing a number from inside the country.
	//	Example US: "1 NNNNNNNNNN"

	DisplaySubNumFmtFullInternal string
	//	The format used to display a Subscriber phone
	//	number in text.
	//
	//	The character 'N' is used as a placeholder for
	//	numeric digits. This is the format used when
	//	dialing a number from inside the country.
	//	Example US: 1 (NNN) NNN-NNNN

	DialingMobileNumFmtFullInternal string
	//	The format used when dialing a Mobile number.
	//	The character 'N' is used as a placeholder for
	//	numeric digits. This is the format used when
	//	dialing a number from inside the country.
	//	Example US: "1 NNNNNNNNNN"

	DisplayMobileNumFmtFullInternal string
	//	The format used to display a Mobile phone
	//	number in text.
	//
	//	The character 'N' is used as a placeholder for
	//	numeric digits. This is the format used when
	//	dialing a number from inside the country.
	//	Example US: 1 (NNN) NNN-NNNN

	DisplaySubNumFmtAbbrExternal string
	//	The format used to display an abbreviated
	//	Subscriber phone number in text.
	//
	//	The character 'N' is used as a placeholder for
	//	numeric digits. This is the format used when
	//	dialing in from outside the country.
	//	Example US: (NNN) NNN-NNNN

	DialingMobileNumFmtAbbrExternal string
	//	The format used when dialing an abbreviated
	//	Mobile number.
	//
	//	The character 'N' is used as a placeholder for
	//	numeric digits. This is the format used when
	//	dialing in from outside the country.
	//	Example US: "NNNNNNNNNN"

	DisplayMobileNumFmtAbbrExternal string
	//	The format used to display an abbreviated
	//	Mobile phone number in text.
	//
	//	The character 'N' is used as a placeholder for
	//	numeric digits. This is the format used when
	//	dialing in from outside the country.
	//	Example US: 1 (NNN) NNN-NNNN

	DialingSubNumFmtAbbrInternal string
	//	The format used when dialing an abbreviated
	//	Subscriber number.
	//
	//	The character 'N' is used as a placeholder for
	//	numeric digits. This is the format used when
	//	dialing a number from inside the country.
	//	Example US: "1 NNNNNNNNNN"

	DisplaySubNumFmtAbbrInternal string
	//	The format used to display an abbreviated
	//	Subscriber phone number in text.
	//
	//	The character 'N' is used as a placeholder for
	//	numeric digits. This is the format used when
	//	dialing a number from inside the country.
	//	Example US: 1 (NNN) NNN-NNNN

	DialingMobileNumFmtAbbrInternal string
	//	The format used when dialing an abbreviated
	//	Mobile number.
	//
	//	The character 'N' is used as a placeholder for
	//	numeric digits. This is the format used when
	//	dialing a number from inside the country.
	//	Example US: "1 NNNNNNNNNN"

	DisplayMobileNumFmtAbbrInternal string
	//	The format used to display an abbreviated
	//	Mobile phone number in text.
	//
	//	The character 'N' is used as a placeholder for
	//	numeric digits. This is the format used when
	//	dialing a number from inside the country.
	//	Example US: 1 (NNN) NNN-NNNN

	lock *sync.Mutex
}
