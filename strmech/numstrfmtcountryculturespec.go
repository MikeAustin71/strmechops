package strmech

import "sync"

type NumStrFmtCountryCultureSpec struct {
	IdNo uint64
	//	Optional
	//	An identification number used to differentiate
	//	and track multiple Country Culture Specification
	//	objects.

	IdString string
	//	Optional
	//	An identification string of text characters
	//	used to differentiate and track multiple
	//	Country Culture Specification objects.

	Description string
	//	Optional
	//	This string contains descriptive text describing
	//	the Country Culture Specification instance.

	Tag string
	//	Optional
	//	This string contains descriptive text describing
	//	the Country Culture Specification instance.

	CountryIdNo uint64
	//	Optional
	//	A number identifying the specific country
	//	or culture specified by the current Country
	//	Culture Specification instance.

	CountryIdString string
	//	Optional
	//	A string of text characters identifying the
	//	specific country or culture specified by the
	//	current Country	Culture Specification instance.

	CountryDescription string
	//	Optional
	//	A string of characters providing a	narrative
	//	text description of the country or culture
	//	identified by the current Country Culture
	//	Specification instance.

	CountryTag string
	//	Optional
	//	A string containing a brief text description
	//	 of the country or culture identified by the
	//	 current Country Culture Specification instance.

	CountryCultureName string
	//	Required
	//	The full name of the country or culture
	//	identified by the current Country Culture
	//	Specification instance.

	CountryAbbreviatedName string
	//	Optional
	//	An abbreviated name for the country or culture
	//	identified by the current Country Culture
	//	Specification instance.

	CountryAlternateNames []string
	//	Optional
	//	An alternate or additional name for the country
	//	or culture identified by the current Country
	//	Culture Specification instance.

	CountryCodeTwoChar string
	//	Optional
	//	The unique Two Character code identifying the
	//	country or culture associated with the current
	//	Country Culture Specification instance.

	CountryCodeThreeChar string
	//	Optional
	//	The unique Three Character code identifying the
	//	country or culture associated with the current
	//	Country Culture Specification instance.

	CountryCodeNumber string
	//	Optional
	//	The official code identifier for the country or
	//	culture associated with the current Country
	//	Culture Specification instance.

	CurrencyCode string
	//	Optional
	//	The official currency code associated with
	//	the country or culture identified by the
	//	current Country Culture Specification
	//	instance.

	CurrencyCodeNo string
	//	Optional
	//	The official currency code number associated
	//	with the country or culture identified by the
	//	current Country Culture Specification
	//	instance.

	CurrencyName string
	//	Optional
	//	The official currency code number associated
	//	with the country or culture identified by the
	//	current Country Culture Specification
	//	instance.

	CurrencySymbols []rune
	//	Required
	//	The official currency symbol or symbols for
	//	the country or culture identified by the
	//	current Country Culture Specification
	//	instance.

	MinorCurrencyName string
	//	Optional
	//	The name of the minor currency associated
	//	with the country or culture identified by
	//	the current Country Culture Specification
	//	instance. In the United States, the minor
	//	currency name is "Cents"

	MinorCurrencySymbols []rune
	//	Optional
	//	The Minor Currency symbol or symbols. In
	//	the United States, the minor currency
	//	name is "Cents" and the minor currency
	//	symbol is "¢".

	CurrencyNumDecSep DecimalSeparatorSpec
	//	Required for Currency Number String
	//	Formatting.
	//
	//	The Decimal Separator specifies the
	//	character or characters which serve
	//	as the radix point. Decimal separators
	//	are used to separate integer and
	//	fractional numeric digits in a floating
	//	point numeric value. In the United
	//	States, the Decimal Separator is the
	//	period character ('.') or decimal point.

	CurrencyTurnOnIntegerDigitsSeparation bool
	//	Required for Currency Number String
	//	Formatting.
	//
	//	When set to 'true' Currency Number Strings
	//	will be formatted with integer separation.
	//	This usually means the integer portion
	//	of the numeric value is separated into
	//	thousands.
	//		United States Example: 1,000,000

	CurrencyIntGroupingSpec IntegerSeparatorSpec
	//	Required for Currency Number String
	//	Formatting.
	//
	//	Integer Separator Specification. This
	//	parameter specifies the type of integer
	//	grouping and integer separator characters
	//	which will be applied to the number
	//	string formatting operations.

	CurrencyNegativeValueFmt NumStrNumberSymbolSpec
	//	Required for Currency Number String
	//	Formatting.
	//
	//	The Number String Negative Number Sign
	//	Specification is used to configure negative
	//	number sign symbols for negative numeric
	//	values formatted and displayed in number
	//	stings.
	//
	//	For currency presentations, the currency
	//	symbol is combined with the negative number
	//	sign.

	CurrencyPositiveValueFmt NumStrNumberSymbolSpec
	//	Required for Currency Number String
	//	Formatting.
	//
	//	Positive number signs are commonly implied
	//	and not specified. However, the user has
	//	the option to specify a positive number sign
	//	character or characters for positive numeric
	//	values using a Number String Positive Number
	//	Sign Specification.
	//
	//	For currency presentations, the currency
	//	symbol is combined with the positive number
	//	sign.

	CurrencyZeroValueFmt NumStrNumberSymbolSpec
	//	Required for Currency Number String
	//	Formatting.
	//
	//	The Currency Zero Value Format Specification
	//	is used to configure number symbols for zero
	//	numeric values formatted and displayed in
	//	number stings.
	//
	//	For currency presentations, the currency
	//	symbol is combined with the zero number
	//	sign.

	SignedNumValDecSep DecimalSeparatorSpec
	//	Required for Signed Number String
	//	Formatting.
	//
	//	The Decimal Separator specifies the
	//	character or characters which serve
	//	as the radix point. Decimal separators
	//	are used to separate integer and
	//	fractional numeric digits in a floating
	//	point numeric value. In the United
	//	States, the Decimal Separator is the
	//	period character ('.') or decimal point.

	SignedNumValTurnOnIntegerDigitsSeparation bool
	//	Required for Signed Number String
	//	Formatting.
	//
	//	When set to 'true' Signed Number Strings
	//	will be formatted with integer separation.
	//	This usually means the integer portion
	//	of the numeric value is separated into
	//	thousands.
	//		United States Example: 1,000,000

	SignedNumValIntGroupingSpec IntegerSeparatorSpec
	//	Required for Signed Number String
	//	Formatting.
	//
	//	Integer Separator Specification. This
	//	parameter specifies the type of integer
	//	grouping and integer separator characters
	//	which will be applied to the number
	//	string formatting operations.

	SignedNumValNegativeValueFmt NumStrNumberSymbolSpec
	//	Required for Signed Number String
	//	Formatting.
	//
	//	The Number String Negative Number Sign
	//	Specification is used to configure negative
	//	number sign symbols for negative numeric
	//	values formatted and displayed in number
	//	stings.

	SignedNumValPositiveValueFmt NumStrNumberSymbolSpec
	//	Required for Currency Number String
	//	Formatting.
	//
	//	Positive number signs are commonly implied
	//	and not specified. However, the user has
	//	the option to specify a positive number sign
	//	character or characters for positive numeric
	//	values using a Number String Positive Number
	//	Sign Specification.

	SignedNumValZeroValueFmt NumStrNumberSymbolSpec
	//	Required for Signed Number String
	//	Formatting.
	//
	//	The Currency Zero Value Format Specification
	//	is used to configure number symbols for zero
	//	numeric values formatted and displayed in
	//	number stings.

	lock *sync.Mutex
}
