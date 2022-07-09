package strmech

type CharSearchNumStrParseResultsDto struct {
	SearchResultsName string
	// Optional. The Name, Label or descriptive Tag
	// associated with the current instance of
	// CharSearchNumStrParseResultsDto.

	SearchResultsFunctionChain string
	// Optional. A listing of the functions which led to the
	// creation of this instance of
	// CharSearchNumStrParseResultsDto.

	FoundNumericDigits bool
	// Signals a successful Number String Parsing operation.
	// When set to 'true', this means one or more numeric
	// digit characters ('0' through '9' inclusive) were
	// located in the Target String.

	FoundNonZeroValue bool
	// If numeric digit characters were located in the
	// Target String (see 'FoundNumericDigits' above), this
	// signals whether the numeric value is zero or nonzero.
	// When set to 'true', the found numeric value is some
	// value other than zero. If set to 'false', this means
	// that the found numeric value is zero.

	FoundDecimalSeparatorSymbols bool
	// When set to 'true', this signals that one or more
	// Decimal Separator Symbol characters have been
	// identified in the 'Target String'.

	FoundDecimalDigits bool
	// When set to 'true', this signals that one or more
	// numeric digit characters ('0' through '9' inclusive)
	// were located in the Target String to the right of
	// the Decimal Separator Symbol (a.k.a Decimal Point).

	NumSignValue NumericSignValueType
	// If a numeric value was extracted from the number string
	// (see 'FoundNumericDigits' above), this parameter
	// specifies the number sign associated with that value.
	//
	// Possible values are listed as follows:
	//  NumSignVal.None() - Signals no numeric value was found
	//  NumSignVal.Negative()
	//  NumSignVal.Zero()
	//  NumSignVal.Positive()

	NumValueType NumericValueType
	// If a numeric value was extracted from the number string
	// (see 'FoundNumericDigits' above), this parameter signals
	// whether the numeric value is an integer or a floating
	// point (digits to the right of the decimal) value.
	//
	// Possible values are listed as follows:
	//  NumValType.None() - Signals no numeric value was found
	//  NumValType.FloatingPoint()
	//  NumValType.Integer()

	RemainderString RuneArrayDto
	// A Rune Array Data Transfer Object containing the
	// remaining characters in a Target String which were
	// NOT included in the search operation and which remain
	// to be searched in future search operations.
	//
	// The 'RemainderString' is only returned if requested.

	DecimalSeparatorSearchResults CharSearchDecimalSeparatorResultsDto
	// An instance of CharSearchDecimalSeparatorResultsDto detailing
	// the results of a search for a decimal separator in the Target
	// String.

	NegativeNumberSymbolSearchResults CharSearchNegativeNumberResultsDto
	// An instance of CharSearchNegativeNumberResultsDto detailing the
	// results of a search for negative number symbols in the Target
	// String.

	ParsingTerminatorSearchResults CharSearchRuneArrayResultsDto
	// An instance of CharSearchRuneArrayResultsDto detailing the
	// results of a search for a Number String Parsing Delimiter
	// character or characters. When a Parsing Terminator Delimiter
	// is located in the Target String, the Number Parsing operation
	// is immediately terminated.

}
