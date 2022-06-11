package numberstr

import (
	"sync"
)

// NumStrFormatCountry - Returns the number string formats used
// by specific countries.
//
// Sources:
//  https://gist.github.com/bzerangue/5484121
//  http://symbologic.info/currency.htm
//  http://www.xe.com/symbols.php
//  https://www.countrycode.org/
//  https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes
//  https://www.codeproject.com/articles/78175/international-number-formats
//  https://www.thefinancials.com/Default.aspx?SubSectionID=curformat
//  https://en.wikipedia.org/wiki/List_of_circulating_currencies - Symbols with decoding
//  https://docs.oracle.com/cd/E19455-01/806-0169/overview-9/index.html
//  https://fastspring.com/blog/how-to-format-30-currencies-from-countries-all-over-the-world/
//  https://en.wikipedia.org/wiki/Decimal_separator
//  https://en.wikipedia.org/wiki/ISO_4217   Currency Codes
//  https://english.stackexchange.com/questions/124797/how-to-write-negative-currency-in-text
//  https://freeformatter.com/i18n-standards-code-snippets.html
//  https://www.evertype.com/standards/euro/formats.html
//  https://www.unicode.org/charts/PDF/U20A0.pdf
//  https://www.rapidtables.com/code/text/unicode-characters.html
//  https://en.wikipedia.org/wiki/Currency_symbol
//  https://www.ip2currency.com/currency-symbol
//  https://www.xe.com/iso4217.php#U
//  https://unicode-table.com/en
//
//  https://en.wikipedia.org/wiki/Indian_numbering_system
//  https://en.wikipedia.org/wiki/Chinese_numerals
//
//
// Countries:
//
//  Argentina
//  Australia
//  Austria
//  Canada
//  CanadaFrench
//  Chile
//  China
//  Columbia
//  Czechia
//  France
//  Germany
//  Israel
//  Italy
//  United Kingdom
//  United States

type NumStrFormatCountry struct {
	lock *sync.Mutex
}

// Ptr - Returns a pointer to a new instance of NumStrFormatCountry.
//
func (nStrFmtCountry NumStrFormatCountry) Ptr() *NumStrFormatCountry {
	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	return &NumStrFormatCountry{
		lock: new(sync.Mutex),
	}
}

// Albania - Returns the number string format used in
// The Republic of Albania.
//
//  https://www.xe.com/currency/all-albanian-lek
//  https://en.wikipedia.org/wiki/Decimal_separator
//
func (nStrFmtCountry *NumStrFormatCountry) Albania() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 8
	setupDto.IdString = "008"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 8
	setupDto.CountryIdString = "008"
	setupDto.CountryDescription = "Country Setup - Albania"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Albania"
	setupDto.CountryAbbreviatedName = "Albania"

	setupDto.CountryAlternateNames =
		[]string{
			"The Republic of Albania",
			"Republic of Albania"}

	setupDto.CountryCodeTwoChar = "AL"
	setupDto.CountryCodeThreeChar = "ALB"
	setupDto.CountryCodeNumber = "008"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$ 127.54"
	setupDto.CurrencyNegativeValueFmt = "$ -127.54"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "ALL"
	setupDto.CurrencyCodeNo = "008"
	setupDto.CurrencyName = "Lek"

	setupDto.CurrencySymbols = []rune{
		'\U0000004c',
		'\U00000065',
		'\U0000006b',
	}

	setupDto.MinorCurrencyName = "Qindarkë"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = ' ' // Space
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "-127.54"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Argentina - Returns the number string format used in the
// Argentina.
//
// https://freeformatter.com/argentina-standards-code-snippets.html
//
func (nStrFmtCountry *NumStrFormatCountry) Argentina() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 32
	setupDto.IdString = "032"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 32
	setupDto.CountryIdString = "032"
	setupDto.CountryDescription = "Country Setup - Argentina"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Argentina"
	setupDto.CountryAbbreviatedName = "Argentina"

	setupDto.CountryAlternateNames =
		[]string{
			"Argentine Republic",
			"The Argentine Republic"}

	setupDto.CountryCodeTwoChar = "AR"
	setupDto.CountryCodeThreeChar = "ARG"
	setupDto.CountryCodeNumber = "032"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$ 127.54"
	setupDto.CurrencyNegativeValueFmt = "-$ 127.54"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "ARS"
	setupDto.CurrencyCodeNo = "032"
	setupDto.CurrencyName = "Peso"
	setupDto.CurrencySymbols = []rune{'\U00000024'}

	setupDto.MinorCurrencyName = "Centavo"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = '.'
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "- 127.54"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Australia - Returns the number string format used in
// Australia.
//
func (nStrFmtCountry *NumStrFormatCountry) Australia() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 36
	setupDto.IdString = "36"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 36
	setupDto.CountryIdString = "036"
	setupDto.CountryDescription = "Country Setup - Australia"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Australia"
	setupDto.CountryAbbreviatedName = "Australia"

	setupDto.CountryAlternateNames =
		[]string{
			"Commonwealth of Australia",
			"The Commonwealth of Australia"}

	setupDto.CountryCodeTwoChar = "AU"
	setupDto.CountryCodeThreeChar = "AUS"
	setupDto.CountryCodeNumber = "036"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$ 127.54"
	setupDto.CurrencyNegativeValueFmt = "$ -127.54"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "AUD"
	setupDto.CurrencyCodeNo = "036"
	setupDto.CurrencyName = "Dollar"
	setupDto.CurrencySymbols = []rune{'\U00000024'}

	setupDto.MinorCurrencyName = "Cent"
	setupDto.MinorCurrencySymbols = []rune{'\U000000a2'}

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = ','
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "-127.54"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Austria - Returns the number string format used in
// Austria.
//
//  https://fastspring.com/blog/how-to-format-30-currencies-from-countries-all-over-the-world/
//  https://en.wikipedia.org/wiki/Decimal_separator
//
func (nStrFmtCountry *NumStrFormatCountry) Austria() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 40
	setupDto.IdString = "040"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 40
	setupDto.CountryIdString = "040"
	setupDto.CountryDescription = "Country Setup - Austria"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Austria"
	setupDto.CountryAbbreviatedName = "Austria"

	setupDto.CountryAlternateNames =
		[]string{
			"The Republic of Austria",
			"Republic of Austria"}

	setupDto.CountryCodeTwoChar = "AT"
	setupDto.CountryCodeThreeChar = "AUT"
	setupDto.CountryCodeNumber = "040"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "127.54- $"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "ATS"
	setupDto.CurrencyCodeNo = "978"
	setupDto.CurrencyName = "Euro"
	setupDto.CurrencySymbols = []rune{'\U000020ac'}

	setupDto.MinorCurrencyName = "Cent"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = '.'
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "-127.54"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Bahrain - Returns the number string format used in
// The Kingdom of Bahrain.
//
// https://en.wikipedia.org/wiki/ISO_4217
// https://en.wikipedia.org/wiki/Currency_symbol
// https://www.xe.com/currency/bhd-bahraini-dinar
//
func (nStrFmtCountry *NumStrFormatCountry) Bahrain() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 48
	setupDto.IdString = "048"
	setupDto.Description = "Country Setup - Bahrain"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 48
	setupDto.CountryIdString = "048"
	setupDto.CountryDescription = "Country Setup - Bahrain"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Bahrain"
	setupDto.CountryAbbreviatedName = "Bahrain"

	setupDto.CountryAlternateNames =
		[]string{
			"The Kingdom of Bahrain",
			"Kingdom of Bahrain"}

	setupDto.CountryCodeTwoChar = "BH"
	setupDto.CountryCodeThreeChar = "BHR"
	setupDto.CountryCodeNumber = "048"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "127.54- $"
	setupDto.CurrencyDecimalDigits = 3
	setupDto.CurrencyCode = "BHD"
	setupDto.CurrencyCodeNo = "048"
	setupDto.CurrencyName = "Dinar"
	setupDto.CurrencySymbols = []rune{
		'B',
		'D',
	}

	setupDto.MinorCurrencyName = "Fils"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = ','
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "127.54-"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Bangladesh - Returns the number string format used in
// The People's Republic of Bangladesh.
//
func (nStrFmtCountry *NumStrFormatCountry) Bangladesh() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 050
	setupDto.IdString = "050"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 050
	setupDto.CountryIdString = "050"
	setupDto.CountryDescription = "Country Setup - Bangladesh"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Bangladesh"
	setupDto.CountryAbbreviatedName = "Bangladesh"

	setupDto.CountryAlternateNames =
		[]string{
			"The People's Republic of Bangladesh",
			"People's Republic of Bangladesh"}

	setupDto.CountryCodeTwoChar = "BD"
	setupDto.CountryCodeThreeChar = "BGD"
	setupDto.CountryCodeNumber = "050"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$ 127.54"
	setupDto.CurrencyNegativeValueFmt = "$ -127.54"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "BDT"
	setupDto.CurrencyCodeNo = "050"
	setupDto.CurrencyName = "Taka"
	setupDto.CurrencySymbols = []rune{'\U000009f3'}

	setupDto.MinorCurrencyName = "Paisa"
	setupDto.MinorCurrencySymbols =
		make([]rune, 0, 5)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = ','
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3, 2}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "-127.54"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Belarus - Returns the number string format used in
// The Republic of Belarus.
//
//  https://en.wikipedia.org/wiki/ISO_4217
//  https://en.wikipedia.org/wiki/Currency_symbol
//
func (nStrFmtCountry *NumStrFormatCountry) Belarus() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 112
	setupDto.IdString = "112"
	setupDto.Description = "Country Setup - Belarus"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 112
	setupDto.CountryIdString = "112"
	setupDto.CountryDescription = "Country Setup - Belarus"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Belarus"
	setupDto.CountryAbbreviatedName = "Belarus"

	setupDto.CountryAlternateNames =
		[]string{
			"The Republic of Belarus",
			"Republic of Belarus"}

	setupDto.CountryCodeTwoChar = "BY"
	setupDto.CountryCodeThreeChar = "BLR"
	setupDto.CountryCodeNumber = "112"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "127.54- $"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "BYN"
	setupDto.CurrencyCodeNo = "974"
	setupDto.CurrencyName = "Ruble"
	// 42, 72
	setupDto.CurrencySymbols = []rune{
		'\U00000042',
		'\U00000072',
	}

	setupDto.MinorCurrencyName = "Kapeyka"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = ' ' // space
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "127.54-"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Belgium - Returns the number string format used in The Kingdom
// of Belgium.
//
// https://en.wikipedia.org/wiki/ISO_4217
// https://en.wikipedia.org/wiki/Currency_symbol
//
func (nStrFmtCountry *NumStrFormatCountry) Belgium() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 56
	setupDto.IdString = "056"
	setupDto.Description = "Country Setup - Belgium"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 56
	setupDto.CountryIdString = "056"
	setupDto.CountryDescription = "Country Setup - Belgium"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Belgium"
	setupDto.CountryAbbreviatedName = "Belgium"

	setupDto.CountryAlternateNames =
		[]string{
			"The Kingdom of Belgium",
			"Kingdom of Belgium"}

	setupDto.CountryCodeTwoChar = "BE"
	setupDto.CountryCodeThreeChar = "BEL"
	setupDto.CountryCodeNumber = "056"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "127.54- $"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "BEF"
	setupDto.CurrencyCodeNo = "978"
	setupDto.CurrencyName = "Euro"
	setupDto.CurrencySymbols = []rune{'\U000020ac'}

	setupDto.MinorCurrencyName = "Cent"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = '.'
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "127.54-"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Bitcoin - Returns the number string format used in
// the crypto currency, Bitcoin.
//
func (nStrFmtCountry *NumStrFormatCountry) Bitcoin() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 999999
	setupDto.IdString = "999999"
	setupDto.Description = "Crypto Currency Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 999999
	setupDto.CountryIdString = "999999"
	setupDto.CountryDescription = "Crypto Currency Setup - Bitcoin"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Bitcoin"
	setupDto.CountryAbbreviatedName = "Bitcoin"

	setupDto.CountryAlternateNames =
		[]string{
			"Bitcoin",
		}

	setupDto.CountryCodeTwoChar = "BC"
	setupDto.CountryCodeThreeChar = "BCH"
	setupDto.CountryCodeNumber = "999999"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$ 127.54"
	setupDto.CurrencyNegativeValueFmt = "$ -127.54"
	setupDto.CurrencyDecimalDigits = 8
	setupDto.CurrencyCode = "BCH"
	setupDto.CurrencyCodeNo = "999999"
	setupDto.CurrencyName = "Bitcoin"
	setupDto.CurrencySymbols =
		[]rune{'\U000020bf'}

	setupDto.MinorCurrencyName = ""
	setupDto.MinorCurrencySymbols =
		make([]rune, 0, 5)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = ','
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "-127.54"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// BosniaHerzegovina - Bosnia and Herzegovina. Returns the number
// string format used in the country of Bosnia and Herzegovina.
//
//  https://en.wikipedia.org/wiki/Decimal_separator
//  https://www.xe.com/currency/bam-bosnian-convertible-mark
//
func (nStrFmtCountry *NumStrFormatCountry) BosniaHerzegovina() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 70
	setupDto.IdString = "070"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 70
	setupDto.CountryIdString = "070"
	setupDto.CountryDescription = "Country Setup - Bosnia and Herzegovina"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Bosnia and Herzegovina"
	setupDto.CountryAbbreviatedName = "BosniaHerzegovina"

	setupDto.CountryAlternateNames =
		[]string{
			"BosniaHerzegovina",
		}

	setupDto.CountryCodeTwoChar = "BA"
	setupDto.CountryCodeThreeChar = "BIH"
	setupDto.CountryCodeNumber = "070"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$ 127.54"
	setupDto.CurrencyNegativeValueFmt = "$ -127.54"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "BAM"
	setupDto.CurrencyCodeNo = "977"
	setupDto.CurrencyName = "Marka"
	setupDto.CurrencySymbols = []rune{
		'\U0000004b',
		'\U0000004d',
	}

	setupDto.MinorCurrencyName = "Fenning"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = '.'
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "127.54-"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Brazil - Returns the number string format used in the
// Brazil.
//
func (nStrFmtCountry *NumStrFormatCountry) Brazil() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 76
	setupDto.IdString = "076"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 76
	setupDto.CountryIdString = "076"
	setupDto.CountryDescription = "Country Setup - Brazil"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Brazil"
	setupDto.CountryAbbreviatedName = "Brazil"

	setupDto.CountryAlternateNames =
		[]string{
			"The Federative Republic of Brazil"}

	setupDto.CountryCodeTwoChar = "BR"
	setupDto.CountryCodeThreeChar = "BRA"
	setupDto.CountryCodeNumber = "076"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$ 127.54"
	setupDto.CurrencyNegativeValueFmt = "$ -127.54"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "BRL"
	setupDto.CurrencyCodeNo = "986"
	setupDto.CurrencyName = "Real"
	setupDto.CurrencySymbols = []rune{'\U00000052', '\U00000024'}

	setupDto.MinorCurrencyName = "Centavo"
	setupDto.MinorCurrencySymbols = []rune{}

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = '.'
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "-127.54"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Bulgaria - Returns the number string format used in
// The Republic of Bulgaria.
//
//  https://fastspring.com/blog/how-to-format-30-currencies-from-countries-all-over-the-world/
//  https://en.wikipedia.org/wiki/Decimal_separator
//
func (nStrFmtCountry *NumStrFormatCountry) Bulgaria() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 100
	setupDto.IdString = "100"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 100
	setupDto.CountryIdString = "100"
	setupDto.CountryDescription = "Country Setup - Bulgaria"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Bulgaria"
	setupDto.CountryAbbreviatedName = "Bulgaria"

	setupDto.CountryAlternateNames =
		[]string{
			"The Republic of Bulgaria",
			"Republic of Bulgaria"}

	setupDto.CountryCodeTwoChar = "BG"
	setupDto.CountryCodeThreeChar = "BGR"
	setupDto.CountryCodeNumber = "100"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "127.54- $"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "BGN"
	setupDto.CurrencyCodeNo = "978"
	setupDto.CurrencyName = "Euro"
	setupDto.CurrencySymbols = []rune{'\U000020ac'}

	setupDto.MinorCurrencyName = "Cent"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = ' ' // Space
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "127.54-"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Canada - Returns the number string format used in
// Canada.
//
func (nStrFmtCountry *NumStrFormatCountry) Canada() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 124
	setupDto.IdString = "124"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 124
	setupDto.CountryIdString = "124"
	setupDto.CountryDescription = "Country Setup - Canada"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Canada"
	setupDto.CountryAbbreviatedName = "Canada"

	setupDto.CountryAlternateNames =
		[]string{
			"Canada"}

	setupDto.CountryCodeTwoChar = "CA"
	setupDto.CountryCodeThreeChar = "CAN"
	setupDto.CountryCodeNumber = "124"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$ 127.54"
	setupDto.CurrencyNegativeValueFmt = "-$ 127.54"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "CAD"
	setupDto.CurrencyCodeNo = "124"
	setupDto.CurrencyName = "Dollar"
	setupDto.CurrencySymbols = []rune{'\U00000024'}

	setupDto.MinorCurrencyName = "Cent"
	setupDto.MinorCurrencySymbols = []rune{'\U000000a2'}

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = ','
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "-127.54"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// CanadaFrench - Returns the number string format used in
// French Canada.
//
func (nStrFmtCountry *NumStrFormatCountry) CanadaFrench() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 124
	setupDto.IdString = "124"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 124
	setupDto.CountryIdString = "124"
	setupDto.CountryDescription = "Country Setup - Canada French"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Canada French"
	setupDto.CountryAbbreviatedName = "Canada French"

	setupDto.CountryAlternateNames =
		[]string{
			"Canada French",
			"French Canadian"}

	setupDto.CountryCodeTwoChar = "CA"
	setupDto.CountryCodeThreeChar = "CAN"
	setupDto.CountryCodeNumber = "124"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "127.54- $"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "CAD"
	setupDto.CurrencyCodeNo = "124"
	setupDto.CurrencyName = "Dollar"
	setupDto.CurrencySymbols = []rune{'\U00000024'}

	setupDto.MinorCurrencyName = "Cent"
	setupDto.MinorCurrencySymbols = []rune{'\U000000a2'}

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = '.'
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "127.54 -"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Chile - Returns the number string format used in the
// the The Republic of Chile.
//
func (nStrFmtCountry *NumStrFormatCountry) Chile() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 152
	setupDto.IdString = "152"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 152
	setupDto.CountryIdString = "152"
	setupDto.CountryDescription = "Country Setup - Chile"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Chile"
	setupDto.CountryAbbreviatedName = "Chile"

	setupDto.CountryAlternateNames =
		[]string{
			"The Republic of Chile"}

	setupDto.CountryCodeTwoChar = "CL"
	setupDto.CountryCodeThreeChar = "CHL"
	setupDto.CountryCodeNumber = "152"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$127.54"
	setupDto.CurrencyNegativeValueFmt = "-$127.54"
	setupDto.CurrencyDecimalDigits = 0
	setupDto.CurrencyCode = "CLP"
	setupDto.CurrencyCodeNo = "152"
	setupDto.CurrencyName = "Peso"
	setupDto.CurrencySymbols = []rune{'\U00000024'}

	setupDto.MinorCurrencyName = "Centavo"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = ','
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "-127.54"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// China - Returns the number string format used in the
// Peoples Republic of China.
//
func (nStrFmtCountry *NumStrFormatCountry) China() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 156
	setupDto.IdString = "156"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 156
	setupDto.CountryIdString = "156"
	setupDto.CountryDescription = "Country Setup - China"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "China"
	setupDto.CountryAbbreviatedName = "CHN"

	setupDto.CountryAlternateNames =
		[]string{
			"Peoples Republic of China",
			"The Peoples Republic of China"}

	setupDto.CountryCodeTwoChar = "CN"
	setupDto.CountryCodeThreeChar = "CHN"
	setupDto.CountryCodeNumber = "156"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$ 127.54"
	setupDto.CurrencyNegativeValueFmt = "-$ 127.54"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "CNY"
	setupDto.CurrencyCodeNo = "156"
	setupDto.CurrencyName = "Yuan"
	setupDto.CurrencySymbols = []rune{'\U000000a5'}

	setupDto.MinorCurrencyName = "Jiao"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = ','
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "-127.54"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Columbia - Returns the number string format used in the
// Columbia.
//
func (nStrFmtCountry *NumStrFormatCountry) Columbia() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 170
	setupDto.IdString = "170"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 170
	setupDto.CountryIdString = "170"
	setupDto.CountryDescription = "Country Setup - Columbia"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Columbia"
	setupDto.CountryAbbreviatedName = "Columbia"

	setupDto.CountryAlternateNames =
		[]string{
			"The Republic of Colombia",
			"Republic of Colombia"}

	setupDto.CountryCodeTwoChar = "CO"
	setupDto.CountryCodeThreeChar = "COL"
	setupDto.CountryCodeNumber = "170"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$ 127.54"
	setupDto.CurrencyNegativeValueFmt = "-$ 127.54"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "COP"
	setupDto.CurrencyCodeNo = "170"
	setupDto.CurrencyName = "Peso"
	setupDto.CurrencySymbols = []rune{'\U00000024'}

	setupDto.MinorCurrencyName = "Centavo"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = '.'
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "-127.54"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Congo - Returns the number string format used in
// The Democratic Republic of the Congo.
//
//  https://www.xe.com/currency/cdf-congolese-franc
//
func (nStrFmtCountry *NumStrFormatCountry) Congo() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 180
	setupDto.IdString = "180"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 180
	setupDto.CountryIdString = "180"
	setupDto.CountryDescription = "Country Setup - Congo"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Congo"
	setupDto.CountryAbbreviatedName = "Congo"

	setupDto.CountryAlternateNames =
		[]string{
			"The Democratic Republic of the Congo",
			"Democratic Republic of the Congo"}

	setupDto.CountryCodeTwoChar = "CD"
	setupDto.CountryCodeThreeChar = "COD"
	setupDto.CountryCodeNumber = "180"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "127.54- $"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "CDF"
	setupDto.CurrencyCodeNo = "180"
	setupDto.CurrencyName = "Franc"
	setupDto.CurrencySymbols = []rune{
		'C',
		'D',
		'F'}

	setupDto.MinorCurrencyName = ""
	setupDto.MinorCurrencySymbols =
		make([]rune, 0, 5)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = ','
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "127.54-"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// CostaRica - Returns the number string format used in
// The Republic of Costa Rica.
//
//  https://en.wikipedia.org/wiki/Decimal_separator
//
func (nStrFmtCountry *NumStrFormatCountry) CostaRica() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 188
	setupDto.IdString = "188"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 188
	setupDto.CountryIdString = "188"
	setupDto.CountryDescription = "Country Setup - Costa Rica"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Costa Rica"
	setupDto.CountryAbbreviatedName = "Costa Rica"

	setupDto.CountryAlternateNames =
		[]string{
			"The Republic of Costa Rica",
			"Republic of Costa Rica"}

	setupDto.CountryCodeTwoChar = "CR"
	setupDto.CountryCodeThreeChar = "CRI"
	setupDto.CountryCodeNumber = "188"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$ 127.54"
	setupDto.CurrencyNegativeValueFmt = "$ -127.54"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "CRC"
	setupDto.CurrencyCodeNo = "188"
	setupDto.CurrencyName = "Colon"
	setupDto.CurrencySymbols = []rune{'\U000020a1'}

	setupDto.MinorCurrencyName = "Centimo"
	setupDto.MinorCurrencySymbols =
		make([]rune, 0, 5)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = ' ' // Space
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "-127.54"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Croatia - Returns the number string format used in
// The Republic of Croatia.
//
//  https://fastspring.com/blog/how-to-format-30-currencies-from-countries-all-over-the-world/
//  https://en.wikipedia.org/wiki/Decimal_separator
//
func (nStrFmtCountry *NumStrFormatCountry) Croatia() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 191
	setupDto.IdString = "191"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 191
	setupDto.CountryIdString = "191"
	setupDto.CountryDescription = "Country Setup - Croatia"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Croatia"
	setupDto.CountryAbbreviatedName = "Croatia"

	setupDto.CountryAlternateNames =
		[]string{
			"The Republic of Croatia",
			"Republic of Croatia"}

	setupDto.CountryCodeTwoChar = "HR"
	setupDto.CountryCodeThreeChar = "HRV"
	setupDto.CountryCodeNumber = "191"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "127.54- $"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "HRK"
	setupDto.CurrencyCodeNo = "978"
	setupDto.CurrencyName = "Euro"
	setupDto.CurrencySymbols = []rune{'\U000020ac'}

	setupDto.MinorCurrencyName = "Cent"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = ' ' // Space
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "127.54-"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Cuba - Returns the number string format used in
// The Republic of Cuba.
//
func (nStrFmtCountry *NumStrFormatCountry) Cuba() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 192
	setupDto.IdString = "192"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 192
	setupDto.CountryIdString = "192"
	setupDto.CountryDescription = "Country Setup - Cuba"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Cuba"
	setupDto.CountryAbbreviatedName = "Cuba"

	setupDto.CountryAlternateNames =
		[]string{
			"The Republic of Cuba",
			"Republic of Cuba"}

	setupDto.CountryCodeTwoChar = "CU"
	setupDto.CountryCodeThreeChar = "CUB"
	setupDto.CountryCodeNumber = "192"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$ 127.54"
	setupDto.CurrencyNegativeValueFmt = "$ -127.54"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "CUP"
	setupDto.CurrencyCodeNo = "192"
	setupDto.CurrencyName = "Peso"
	setupDto.CurrencySymbols = []rune{'\U000020b1'}

	setupDto.MinorCurrencyName = "Centavo"
	setupDto.MinorCurrencySymbols =
		make([]rune, 0, 5)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = ','
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "-127.54"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Cyprus - Returns the number string format used in
// The Republic of Cyprus.
//
// https://fastspring.com/blog/how-to-format-30-currencies-from-countries-all-over-the-world/
// https://en.wikipedia.org/wiki/Decimal_separator
//
func (nStrFmtCountry *NumStrFormatCountry) Cyprus() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 196
	setupDto.IdString = "196"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 196
	setupDto.CountryIdString = "196"
	setupDto.CountryDescription = "Country Setup - Cyprus"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Cyprus"
	setupDto.CountryAbbreviatedName = "Cyprus"

	setupDto.CountryAlternateNames =
		[]string{
			"European Union",
			"Euro"}

	setupDto.CountryCodeTwoChar = "CY"
	setupDto.CountryCodeThreeChar = "CYP"
	setupDto.CountryCodeNumber = "196"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "127.54- $"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "CYP"
	setupDto.CurrencyCodeNo = "978"
	setupDto.CurrencyName = "Euro"
	setupDto.CurrencySymbols = []rune{'\U000020ac'}

	setupDto.MinorCurrencyName = "Cent"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = ','
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "127.54-"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// CzechiaEuro - Returns the number string format used in
// The Czech Republic.
//
// Czechia, or The Czech Republic, is a member of the European
// Union. As such, it is legally bound to adopt the 'Euro' as its
// official currency. As of 2020, it has been hesitant to do so
// and the Czech Koruna still remains in wide spread use.
//
// This format configures number strings with the Euro currency
// symbol. For the Czech Koruna currency, reference:
//   NumStrFormatCountry.CzechiaKoruna()
//
//  https://fastspring.com/blog/how-to-format-30-currencies-from-countries-all-over-the-world/
//  https://en.wikipedia.org/wiki/Decimal_separator
//
func (nStrFmtCountry *NumStrFormatCountry) CzechiaEuro() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 203
	setupDto.IdString = "203"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 203
	setupDto.CountryIdString = "203"
	setupDto.CountryDescription = "Country Setup -  Czechia Euro"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Czechia"
	setupDto.CountryAbbreviatedName = "Czechia"

	setupDto.CountryAlternateNames =
		[]string{
			"The Czech Republic",
			"Czech Republic"}

	setupDto.CountryCodeTwoChar = "CZ"
	setupDto.CountryCodeThreeChar = "CZE"
	setupDto.CountryCodeNumber = "203"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "127.54- $"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "CZK"
	setupDto.CurrencyCodeNo = "978"
	setupDto.CurrencyName = "Euro"
	setupDto.CurrencySymbols = []rune{'\U000020ac'}

	setupDto.MinorCurrencyName = "Cent"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = ' ' // Space
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "127.54-"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// CzechiaKoruna - Returns the number string format used in the
// The Czech Republic.
//
// Czechia, or The Czech Republic, is a member of the European
// Union. As such, it is legally bound to adopt the 'Euro' as its
// official currency. As of 2020, it has been hesitant to do so
// and the Czech Koruna still remains in wide spread use.
//
// This format configures number strings with the Koruna currency
// symbol. For the Czech Euro currency, reference:
//   NumStrFormatCountry.CzechiaEuro()
//
//  https://fastspring.com/blog/how-to-format-30-currencies-from-countries-all-over-the-world/
//  https://en.wikipedia.org/wiki/Decimal_separator
//
func (nStrFmtCountry *NumStrFormatCountry) CzechiaKoruna() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 9203
	setupDto.IdString = "9203"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 203
	setupDto.CountryIdString = "203"
	setupDto.CountryDescription = "Country Setup - Czechia"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Czechia"
	setupDto.CountryAbbreviatedName = "Czechia"

	setupDto.CountryAlternateNames =
		[]string{
			"The Czech Republic",
			"Czech Republic"}

	setupDto.CountryCodeTwoChar = "CZ"
	setupDto.CountryCodeThreeChar = "CZE"
	setupDto.CountryCodeNumber = "203"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$127.54"
	setupDto.CurrencyNegativeValueFmt = "-$127.54"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "CZK"
	setupDto.CurrencyCodeNo = "203"
	setupDto.CurrencyName = "Koruna"
	setupDto.CurrencySymbols = []rune{
		'\U0000004b', '\U0000010d'}

	setupDto.MinorCurrencyName = "Haler"
	setupDto.MinorCurrencySymbols =
		make([]rune, 0)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = ' ' // Space
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "-127.54"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Denmark - Returns the number string format used in the
// Kingdom of Denmark.
//
// https://en.wikipedia.org/wiki/ISO_4217
// https://en.wikipedia.org/wiki/Currency_symbol
// https://docs.microsoft.com/en-us/globalization/locale/currency-formatting
//
func (nStrFmtCountry *NumStrFormatCountry) Denmark() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 208
	setupDto.IdString = "208"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 208
	setupDto.CountryIdString = "208"
	setupDto.CountryDescription = "Country Setup - Denmark"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Denmark"
	setupDto.CountryAbbreviatedName = "Denmark"

	setupDto.CountryAlternateNames =
		[]string{
			"The Kingdom of Denmark",
			"Kingdom of Denmark"}

	setupDto.CountryCodeTwoChar = "DK"
	setupDto.CountryCodeThreeChar = "DNK"
	setupDto.CountryCodeNumber = "208"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$127.54"
	setupDto.CurrencyNegativeValueFmt = "$-127.54"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "DKK"
	setupDto.CurrencyCodeNo = "208"
	setupDto.CurrencyName = "Krone"
	setupDto.CurrencySymbols = []rune{'\U0000006b', '\U00000072'}

	setupDto.MinorCurrencyName = "øre"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = '.'
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "127.54-"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Egypt - Returns the number string format used in
// The Arab Republic of Egypt.
//
// https://www.thefinancials.com/Default.aspx?SubSectionID=curformat
//
func (nStrFmtCountry *NumStrFormatCountry) Egypt() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 818
	setupDto.IdString = "818"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 818
	setupDto.CountryIdString = "818"
	setupDto.CountryDescription = "Country Setup - Egypt"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Egypt"
	setupDto.CountryAbbreviatedName = "Egypt"

	setupDto.CountryAlternateNames =
		[]string{
			"The Arab Republic of Egypt",
			"Arab Republic of Egypt",
		}

	setupDto.CountryCodeTwoChar = "EG"
	setupDto.CountryCodeThreeChar = "EGY"
	setupDto.CountryCodeNumber = "818"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$127.54"
	setupDto.CurrencyNegativeValueFmt = "-$127.54"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "EGP"
	setupDto.CurrencyCodeNo = "818"
	setupDto.CurrencyName = "Pound"
	setupDto.CurrencySymbols = []rune{'\U000000a3'}

	setupDto.MinorCurrencyName = "Pence"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = ','
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "-127.54"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1
	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Estonia - Returns the number string format used in The Republic
// of Estonia.
//
// https://en.wikipedia.org/wiki/ISO_4217
// https://en.wikipedia.org/wiki/Currency_symbol
//
func (nStrFmtCountry *NumStrFormatCountry) Estonia() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 233
	setupDto.IdString = "233"
	setupDto.Description = "Country Setup - Estonia"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 233
	setupDto.CountryIdString = "233"
	setupDto.CountryDescription = "Country Setup - Estonia"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Estonia"
	setupDto.CountryAbbreviatedName = "Estonia"

	setupDto.CountryAlternateNames =
		[]string{
			"The Republic of Estonia",
			"Republic of Estonia"}

	setupDto.CountryCodeTwoChar = "EE"
	setupDto.CountryCodeThreeChar = "EST"
	setupDto.CountryCodeNumber = "233"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "127.54- $"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "EEK"
	setupDto.CurrencyCodeNo = "978"
	setupDto.CurrencyName = "Euro"
	setupDto.CurrencySymbols = []rune{'\U000020ac'}

	setupDto.MinorCurrencyName = "Cent"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = ' ' // Space
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "127.54-"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Euro - Returns the number string format used in the
// European Union.
//
// https://fastspring.com/blog/how-to-format-30-currencies-from-countries-all-over-the-world/
//
func (nStrFmtCountry *NumStrFormatCountry) Euro() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 77777
	setupDto.IdString = "77777"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 77777
	setupDto.CountryIdString = "77777"
	setupDto.CountryDescription = "Country Setup - European Union"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "European Union"
	setupDto.CountryAbbreviatedName = "EU"

	setupDto.CountryAlternateNames =
		[]string{
			"European Union",
			"Euro"}

	setupDto.CountryCodeTwoChar = "EU"
	setupDto.CountryCodeThreeChar = "EUR"
	setupDto.CountryCodeNumber = "77777"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "127.54- $"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "EUR"
	setupDto.CurrencyCodeNo = "77777"
	setupDto.CurrencyName = "Euro"
	setupDto.CurrencySymbols = []rune{'\U000020ac'}

	setupDto.MinorCurrencyName = "Cent"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = '.'
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "127.54-"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Finland - Returns the number string format used in The Republic
// of Finland.
//
//  https://en.wikipedia.org/wiki/ISO_4217
//  https://en.wikipedia.org/wiki/Currency_symbol
//  https://en.wikipedia.org/wiki/Decimal_separator
//
func (nStrFmtCountry *NumStrFormatCountry) Finland() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 246
	setupDto.IdString = "246"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 246
	setupDto.CountryIdString = "246"
	setupDto.CountryDescription = "Country Setup - Finland"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Finland"
	setupDto.CountryAbbreviatedName = "Finland"

	setupDto.CountryAlternateNames =
		[]string{
			"The Republic of Finland",
			"Republic of Finland"}

	setupDto.CountryCodeTwoChar = "FI"
	setupDto.CountryCodeThreeChar = "FIN"
	setupDto.CountryCodeNumber = "246"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "127.54- $"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "FIM"
	setupDto.CurrencyCodeNo = "978"
	setupDto.CurrencyName = "Euro"
	setupDto.CurrencySymbols = []rune{'\U000020ac'}

	setupDto.MinorCurrencyName = "Cent"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = ' ' // Space
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "127.54-"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// France - Returns the number string format used in the
// French Republic.
//
// https://www.ibm.com/support/pages/english-and-french-currency-formats
// https://freeformatter.com/france-standards-code-snippets.html
// https://docs.microsoft.com/en-us/globalization/locale/currency-formatting
//
func (nStrFmtCountry *NumStrFormatCountry) France() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 250
	setupDto.IdString = "250"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 250
	setupDto.CountryIdString = "250"
	setupDto.CountryDescription = "Country Setup - France"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "France"
	setupDto.CountryAbbreviatedName = "France"

	setupDto.CountryAlternateNames =
		[]string{
			"French Republic",
			"The French Republic"}

	setupDto.CountryCodeTwoChar = "FR"
	setupDto.CountryCodeThreeChar = "FRA"
	setupDto.CountryCodeNumber = "250"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "-127.54 $"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "FRF"
	setupDto.CurrencyCodeNo = "978"
	setupDto.CurrencyName = "Euro"
	setupDto.CurrencySymbols = []rune{'\U000020ac'}

	setupDto.MinorCurrencyName = "Cent"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = ' ' // Space
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "-127.54"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Germany - Returns the number string format used in the
// Federal Republic of Germany.
//
// https://freeformatter.com/germany-standards-code-snippets.html
// https://www.evertype.com/standards/euro/formats.html
//
func (nStrFmtCountry *NumStrFormatCountry) Germany() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 276
	setupDto.IdString = "276"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 276
	setupDto.CountryIdString = "276"
	setupDto.CountryDescription = "Country Setup - Germany"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Germany"
	setupDto.CountryAbbreviatedName = "Germany"

	setupDto.CountryAlternateNames =
		[]string{
			"Federal Republic of Germany",
			"The Federal Republic of Germany"}

	setupDto.CountryCodeTwoChar = "DE"
	setupDto.CountryCodeThreeChar = "DEU"
	setupDto.CountryCodeNumber = "276"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "127.54- $"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "DEM"
	setupDto.CurrencyCodeNo = "978"
	setupDto.CurrencyName = "Euro"
	setupDto.CurrencySymbols = []rune{'\U000020ac'}

	setupDto.MinorCurrencyName = "Cent"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = '.'
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "127.54-"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Greece - Returns the number string format used in The Hellenic
// Republic.
//
// https://en.wikipedia.org/wiki/ISO_4217
// https://en.wikipedia.org/wiki/Currency_symbol
//
func (nStrFmtCountry *NumStrFormatCountry) Greece() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 300
	setupDto.IdString = "300"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 300
	setupDto.CountryIdString = "300"
	setupDto.CountryDescription = "Country Setup - Greece"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Greece"
	setupDto.CountryAbbreviatedName = "Greece"

	setupDto.CountryAlternateNames =
		[]string{
			"The Hellenic Republic",
			"Hellenic Republic"}

	setupDto.CountryCodeTwoChar = "GR"
	setupDto.CountryCodeThreeChar = "GRC"
	setupDto.CountryCodeNumber = "300"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "127.54- $"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "GRD"
	setupDto.CurrencyCodeNo = "978"
	setupDto.CurrencyName = "Euro"
	setupDto.CurrencySymbols = []rune{'\U000020ac'}

	setupDto.MinorCurrencyName = "Cent"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = '.'
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "127.54-"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// HongKong - Returns the number string format used in
// The Hong Kong Special Administrative Region of China.
//
func (nStrFmtCountry *NumStrFormatCountry) HongKong() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 344
	setupDto.IdString = "344"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 344
	setupDto.CountryIdString = "344"
	setupDto.CountryDescription = "Country Setup - Hong Kong"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Hong Kong"
	setupDto.CountryAbbreviatedName = "Hong Kong"

	setupDto.CountryAlternateNames =
		[]string{
			"The Hong Kong Special Administrative Region of China",
			"Hong Kong Special Administrative Region of China",
		}

	setupDto.CountryCodeTwoChar = "HK"
	setupDto.CountryCodeThreeChar = "HKG"
	setupDto.CountryCodeNumber = "344"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$ 127.54"
	setupDto.CurrencyNegativeValueFmt = "$ -127.54"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "HKD"
	setupDto.CurrencyCodeNo = "344"
	setupDto.CurrencyName = "Dollar"
	setupDto.CurrencySymbols = []rune{'\U00000024'}

	setupDto.MinorCurrencyName = "Cent"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 5)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = ','
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "-127.54"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Hungary - Returns the number string format used in the
// country of Hungary.
//
//  https://en.wikipedia.org/wiki/ISO_4217
//  https://en.wikipedia.org/wiki/Currency_symbol
//  https://en.wikipedia.org/wiki/Decimal_separator
//
func (nStrFmtCountry *NumStrFormatCountry) Hungary() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 348
	setupDto.IdString = "348"
	setupDto.Description = "Country Setup - Hungary"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 348
	setupDto.CountryIdString = "348"
	setupDto.CountryDescription = "Country Setup - Hungary"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Hungary"
	setupDto.CountryAbbreviatedName = "Hungary"

	setupDto.CountryAlternateNames =
		[]string{
			"Hungary"}

	setupDto.CountryCodeTwoChar = "HU"
	setupDto.CountryCodeThreeChar = "HUN"
	setupDto.CountryCodeNumber = "348"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "127.54- $"
	setupDto.CurrencyDecimalDigits = 0
	setupDto.CurrencyCode = "RON"
	setupDto.CurrencyCodeNo = "946"
	setupDto.CurrencyName = "Forint"
	setupDto.CurrencySymbols = []rune{
		'\U00000046',
		'\U00000074'}

	setupDto.MinorCurrencyName = "NONE"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = ' ' // Space
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "127.54-"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Iceland - Returns the number string format used in the
// country of Iceland.
//
//  https://www.lonelyplanet.com/iceland/a/nar-gr/money-and-costs
//  https://www.xe.com/currency/isk-icelandic-krona
//  https://www.xe.com/symbols.php
//  https://www.thefinancials.com/Default.aspx?SubSectionID=curformat
//  https://www.swedishnomad.com/currency-iceland/
//
func (nStrFmtCountry *NumStrFormatCountry) Iceland() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 352
	setupDto.IdString = "352"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 352
	setupDto.CountryIdString = "352"
	setupDto.CountryDescription = "Country Setup - Iceland"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Iceland"
	setupDto.CountryAbbreviatedName = "Iceland"

	setupDto.CountryAlternateNames =
		[]string{
			"Iceland",
		}

	setupDto.CountryCodeTwoChar = "IS"
	setupDto.CountryCodeThreeChar = "ISL"
	setupDto.CountryCodeNumber = "352"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$ 127.54"
	setupDto.CurrencyNegativeValueFmt = "$ -127.54"
	setupDto.CurrencyDecimalDigits = 0
	setupDto.CurrencyCode = "ISK"
	setupDto.CurrencyCodeNo = "352"
	setupDto.CurrencyName = "Krona"
	setupDto.CurrencySymbols = []rune{
		'\U0000006b',
		'\U00000072',
	}

	setupDto.MinorCurrencyName = ""
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = ','
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "-127.54"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// India - Returns the number string format used in
// The Republic of India.
//
//  https://freeformatter.com/israel-standards-code-snippets.html
//
func (nStrFmtCountry *NumStrFormatCountry) India() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 356
	setupDto.IdString = "356"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 356
	setupDto.CountryIdString = "356"
	setupDto.CountryDescription = "Country Setup - India"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "India"
	setupDto.CountryAbbreviatedName = "India"

	setupDto.CountryAlternateNames =
		[]string{
			"The Republic of India",
			"Republic of India"}

	setupDto.CountryCodeTwoChar = "IN"
	setupDto.CountryCodeThreeChar = "IND"
	setupDto.CountryCodeNumber = "356"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$ 127.54"
	setupDto.CurrencyNegativeValueFmt = "$ -127.54"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "INR"
	setupDto.CurrencyCodeNo = "356"
	setupDto.CurrencyName = "Rupee"
	setupDto.CurrencySymbols = []rune{'\U000020b9'}

	setupDto.MinorCurrencyName = "Paise"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = ','
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3, 2}

	setupDto.SignedNumValPositiveValueFmt = ""
	setupDto.SignedNumValNegativeValueFmt = ""
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1
	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Indonesia - Returns the number string format used in
// The Republic of Indonesia.
//
func (nStrFmtCountry *NumStrFormatCountry) Indonesia() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 360
	setupDto.IdString = "360"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 360
	setupDto.CountryIdString = "360"
	setupDto.CountryDescription = "Country Setup - Indonesia"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Indonesia"
	setupDto.CountryAbbreviatedName = "Indonesia"

	setupDto.CountryAlternateNames =
		[]string{
			"The Republic of Indonesia",
			"Republic of Indonesia"}

	setupDto.CountryCodeTwoChar = "ID"
	setupDto.CountryCodeThreeChar = "IDN"
	setupDto.CountryCodeNumber = "360"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$ 127.54"
	setupDto.CurrencyNegativeValueFmt = "$ -127.54"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "IDR"
	setupDto.CurrencyCodeNo = "360"
	setupDto.CurrencyName = "Rupiah"

	setupDto.CurrencySymbols = []rune{
		'\U00000052',
		'\U00000070',
	}

	setupDto.MinorCurrencyName = "Sen"
	setupDto.MinorCurrencySymbols =
		make([]rune, 0, 5)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = '.'
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "-127.54"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Iran - Returns the number string format used in
// The Islamic Republic of Iran.
//
//  https://en.wikipedia.org/wiki/Decimal_separator
//  https://www.xe.com/currency/irr-iranian-rial
//
func (nStrFmtCountry *NumStrFormatCountry) Iran() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 364
	setupDto.IdString = "364"
	setupDto.Description = "Country Setup - Iran"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 364
	setupDto.CountryIdString = "364"
	setupDto.CountryDescription = "Country Setup - Iran"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Iran"
	setupDto.CountryAbbreviatedName = "Iran"

	setupDto.CountryAlternateNames =
		[]string{
			"The Islamic Republic of Iran",
			"Islamic Republic of Iran",
		}

	setupDto.CountryCodeTwoChar = "IR"
	setupDto.CountryCodeThreeChar = "IRN"
	setupDto.CountryCodeNumber = "364"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "127.54- $"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "IRR"
	setupDto.CurrencyCodeNo = "364"
	setupDto.CurrencyName = "Rial"

	setupDto.CurrencySymbols = []rune{
		'\U000000fd',
		'\U000000fc',
	}

	setupDto.MinorCurrencyName = "Dinar"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = ','
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "127.54-"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Ireland - Returns the number string format used in Ireland.
//
// https://en.wikipedia.org/wiki/ISO_4217
// https://en.wikipedia.org/wiki/Currency_symbol
//
func (nStrFmtCountry *NumStrFormatCountry) Ireland() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 372
	setupDto.IdString = "372"
	setupDto.Description = "Country Setup - Ireland"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 372
	setupDto.CountryIdString = "372"
	setupDto.CountryDescription = "Country Setup - Ireland"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Ireland"
	setupDto.CountryAbbreviatedName = "Ireland"

	setupDto.CountryAlternateNames =
		[]string{
			"Ireland"}

	setupDto.CountryCodeTwoChar = "IE"
	setupDto.CountryCodeThreeChar = "IRL"
	setupDto.CountryCodeNumber = "372"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$127.54"
	setupDto.CurrencyNegativeValueFmt = "$-127.54"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "IEP"
	setupDto.CurrencyCodeNo = "978"
	setupDto.CurrencyName = "Euro"
	setupDto.CurrencySymbols = []rune{'\U000020ac'}

	setupDto.MinorCurrencyName = "Cent"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = ','
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "-127.54"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Israel - Returns the number string format used in the
// State of Israel.
//
//  https://freeformatter.com/israel-standards-code-snippets.html
//
func (nStrFmtCountry *NumStrFormatCountry) Israel() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 376
	setupDto.IdString = "376"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 376
	setupDto.CountryIdString = "376"
	setupDto.CountryDescription = "Country Setup - Israel"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Israel"
	setupDto.CountryAbbreviatedName = "Israel"

	setupDto.CountryAlternateNames =
		[]string{
			"State of Israel",
			"The State of Israel"}

	setupDto.CountryCodeTwoChar = "IL"
	setupDto.CountryCodeThreeChar = "ISR"
	setupDto.CountryCodeNumber = "376"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$ 127.54"
	setupDto.CurrencyNegativeValueFmt = "-$ 127.54"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "ILS"
	setupDto.CurrencyCodeNo = "376"
	setupDto.CurrencyName = "Shekel"
	setupDto.CurrencySymbols = []rune{'\U000020aa'}

	setupDto.MinorCurrencyName = "Agorot"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = ','
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = ""
	setupDto.SignedNumValNegativeValueFmt = ""
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1
	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Italy - Returns the number string format used in the
// Italian Republic.
//
// https://freeformatter.com/italy-standards-code-snippets.html
// https://italian.stackexchange.com/questions/5674/what-is-the-correct-way-to-format-currency-in-italian
//
func (nStrFmtCountry *NumStrFormatCountry) Italy() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 380
	setupDto.IdString = "380"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 380
	setupDto.CountryIdString = "380"
	setupDto.CountryDescription = "Country Setup - Italy"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Italy"
	setupDto.CountryAbbreviatedName = "Italy"

	setupDto.CountryAlternateNames =
		[]string{
			"Italian Republic",
			"The Italian Republic"}

	setupDto.CountryCodeTwoChar = "IT"
	setupDto.CountryCodeThreeChar = "ITA"
	setupDto.CountryCodeNumber = "380"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "127.54- $"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "EUR"
	setupDto.CurrencyName = "ITL"
	setupDto.CurrencyCodeNo = "978"
	setupDto.CurrencySymbols = []rune{'\U000020ac'}

	setupDto.MinorCurrencyName = "Cent"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = '.'
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "127.54-"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Japan - Returns the number string format used in
// the country of Japan.
//
// https://fastspring.com/blog/how-to-format-30-currencies-from-countries-all-over-the-world/
//
func (nStrFmtCountry *NumStrFormatCountry) Japan() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 392
	setupDto.IdString = "392"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 392
	setupDto.CountryIdString = "392"
	setupDto.CountryDescription = "Country Setup - Japan"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Japan"
	setupDto.CountryAbbreviatedName = "Japan"

	setupDto.CountryAlternateNames =
		[]string{
			"Japan",
		}

	setupDto.CountryCodeTwoChar = "JP"
	setupDto.CountryCodeThreeChar = "JPN"
	setupDto.CountryCodeNumber = "392"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$ 127.54"
	setupDto.CurrencyNegativeValueFmt = "$ -127.54"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "JPY"
	setupDto.CurrencyCodeNo = "392"
	setupDto.CurrencyName = "Yen"

	setupDto.CurrencySymbols = []rune{
		'\U000000a5',
	}

	setupDto.MinorCurrencyName = "Sen"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 5)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = ','
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "-127.54"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Kenya - Returns the number string format used in
// The Republic of Kenya.
//
//  https://www.xe.com/currency/kes-kenyan-shilling
//  https://www.tuko.co.ke/307213-new-kenyan-currency-everything-know.html
//  https://www.thefinancials.com/Default.aspx?SubSectionID=curformat
//
func (nStrFmtCountry *NumStrFormatCountry) Kenya() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 404
	setupDto.IdString = "404"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 404
	setupDto.CountryIdString = "404"
	setupDto.CountryDescription = "Country Setup - Kenya"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Kenya"
	setupDto.CountryAbbreviatedName = "Kenya"

	setupDto.CountryAlternateNames =
		[]string{
			"The Republic of Kenya",
			"Republic of Kenya"}

	setupDto.CountryCodeTwoChar = "KE"
	setupDto.CountryCodeThreeChar = "KEN"
	setupDto.CountryCodeNumber = "404"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$ 127.54"
	setupDto.CurrencyNegativeValueFmt = "$ -127.54"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "KES"
	setupDto.CurrencyCodeNo = "404"
	setupDto.CurrencyName = "Shilling"
	setupDto.CurrencySymbols = []rune{
		'K',
		'S',
		'h',
	}

	setupDto.MinorCurrencyName = "Cent"
	setupDto.MinorCurrencySymbols = []rune{'c'}

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = ','
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "-127.54"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// KoreaSouth - Returns the number string format used in
// The Republic of Korea.
//
// https://fastspring.com/blog/how-to-format-30-currencies-from-countries-all-over-the-world/
//
func (nStrFmtCountry *NumStrFormatCountry) KoreaSouth() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 410
	setupDto.IdString = "410"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 410
	setupDto.CountryIdString = "410"
	setupDto.CountryDescription = "Country Setup - Korea"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Korea"
	setupDto.CountryAbbreviatedName = "Korea"

	setupDto.CountryAlternateNames =
		[]string{
			"The Republic of Korea",
			"South Korea",
		}

	setupDto.CountryCodeTwoChar = "KR"
	setupDto.CountryCodeThreeChar = "KOR"
	setupDto.CountryCodeNumber = "410"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$ 127.54"
	setupDto.CurrencyNegativeValueFmt = "$ -127.54"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "KRW"
	setupDto.CurrencyCodeNo = "410"
	setupDto.CurrencyName = "Won"

	setupDto.CurrencySymbols = []rune{
		'\U000020a9',
	}

	setupDto.MinorCurrencyName = "Sen"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 5)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = ','
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "-127.54"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Kuwait - Returns the number string format used in
// The State of Kuwait.
//
// https://en.wikipedia.org/wiki/ISO_4217
// https://en.wikipedia.org/wiki/Currency_symbol
// https://fastspring.com/blog/how-to-format-30-currencies-from-countries-all-over-the-world/
//
func (nStrFmtCountry *NumStrFormatCountry) Kuwait() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 414
	setupDto.IdString = "414"
	setupDto.Description = "Country Setup - Kuwait"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 414
	setupDto.CountryIdString = "414"
	setupDto.CountryDescription = "Country Setup - Kuwait"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Kuwait"
	setupDto.CountryAbbreviatedName = "Kuwait"

	setupDto.CountryAlternateNames =
		[]string{
			"The State of Kuwait",
			"State of Kuwait"}

	setupDto.CountryCodeTwoChar = "KW"
	setupDto.CountryCodeThreeChar = "KWT"
	setupDto.CountryCodeNumber = "414"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "127.54- $"
	setupDto.CurrencyDecimalDigits = 3
	setupDto.CurrencyCode = "KWD"
	setupDto.CurrencyCodeNo = "414"
	setupDto.CurrencyName = "Dinar"
	setupDto.CurrencySymbols = []rune{'\U00000643'}

	setupDto.MinorCurrencyName = "Fil"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = ','
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "127.54-"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Latvia - Returns the number string format used in
// The Republic of Latvia.
//
//  https://fastspring.com/blog/how-to-format-30-currencies-from-countries-all-over-the-world/
//  https://en.wikipedia.org/wiki/Decimal_separator
//
func (nStrFmtCountry *NumStrFormatCountry) Latvia() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 428
	setupDto.IdString = "428"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 428
	setupDto.CountryIdString = "428"
	setupDto.CountryDescription = "Country Setup - Latvia"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Latvia"
	setupDto.CountryAbbreviatedName = "Latvia"

	setupDto.CountryAlternateNames =
		[]string{
			"The Republic of Latvia",
			"Republic of Latvia"}

	setupDto.CountryCodeTwoChar = "LV"
	setupDto.CountryCodeThreeChar = "LVA"
	setupDto.CountryCodeNumber = "428"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "127.54- $"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "LVL"
	setupDto.CurrencyCodeNo = "978"
	setupDto.CurrencyName = "Euro"
	setupDto.CurrencySymbols = []rune{'\U000020ac'}

	setupDto.MinorCurrencyName = "Cent"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = ' ' // Space
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "127.54-"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Liechtenstein - Returns the number string format used in
// The Principality of Liechtenstein.
//
// https://en.wikipedia.org/wiki/Decimal_separator
// https://en.wikipedia.org/wiki/Currency_symbol
//
func (nStrFmtCountry *NumStrFormatCountry) Liechtenstein() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 438
	setupDto.IdString = "438"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 438
	setupDto.CountryIdString = "438"
	setupDto.CountryDescription = "Country Setup - Liechtenstein"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Liechtenstein"
	setupDto.CountryAbbreviatedName = "Liechtenstein"

	setupDto.CountryAlternateNames =
		[]string{
			"The Principality of Liechtenstein",
			"Principality of Liechtenstein"}

	setupDto.CountryCodeTwoChar = "LI"
	setupDto.CountryCodeThreeChar = "LIE"
	setupDto.CountryCodeNumber = "438"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "127.54- $"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "CHF"
	setupDto.CurrencyCodeNo = "756"
	setupDto.CurrencyName = "Franc"
	setupDto.CurrencySymbols = []rune{
		'\U00000043',
		'\U00000048',
		'\U00000046'}

	setupDto.MinorCurrencyName = "Rappen"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = '\U00000027' // Apostrophe
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "127.54-"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Lithuania - Returns the number string format used in
// The Republic of Lithuania.
//
//  https://fastspring.com/blog/how-to-format-30-currencies-from-countries-all-over-the-world/
//  https://en.wikipedia.org/wiki/Decimal_separator
//
func (nStrFmtCountry *NumStrFormatCountry) Lithuania() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 440
	setupDto.IdString = "440"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 440
	setupDto.CountryIdString = "440"
	setupDto.CountryDescription = "Country Setup - Lithuania"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Lithuania"
	setupDto.CountryAbbreviatedName = "Lithuania"

	setupDto.CountryAlternateNames =
		[]string{
			"The Republic of Lithuania",
			"Republic of Lithuania"}

	setupDto.CountryCodeTwoChar = "LT"
	setupDto.CountryCodeThreeChar = "LTU"
	setupDto.CountryCodeNumber = "440"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "127.54- $"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "LTL"
	setupDto.CurrencyCodeNo = "978"
	setupDto.CurrencyName = "Euro"
	setupDto.CurrencySymbols = []rune{'\U000020ac'}

	setupDto.MinorCurrencyName = "Cent"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = ' ' // Space
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "127.54-"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Luxembourg - Returns the number string format used in The
// Grand Duchy of Luxembourg.
//
// https://freeformatter.com/germany-standards-code-snippets.html
// https://www.evertype.com/standards/euro/formats.html
//
func (nStrFmtCountry *NumStrFormatCountry) Luxembourg() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 442
	setupDto.IdString = "442"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 442
	setupDto.CountryIdString = "442"
	setupDto.CountryDescription = "Country Setup - Luxembourg"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Luxembourg"
	setupDto.CountryAbbreviatedName = "Luxembourg"

	setupDto.CountryAlternateNames =
		[]string{
			"The Grand Duchy of Luxembourg",
			"Grand Duchy of Luxembourg"}

	setupDto.CountryCodeTwoChar = "LU"
	setupDto.CountryCodeThreeChar = "LUX"
	setupDto.CountryCodeNumber = "442"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "127.54- $"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "LUF"
	setupDto.CurrencyCodeNo = "978"
	setupDto.CurrencyName = "Euro"
	setupDto.CurrencySymbols = []rune{'\U000020ac'}

	setupDto.MinorCurrencyName = "Cent"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = '.'
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "127.54-"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Malaysia - Returns the number string format used in
// the country of Malaysia.
//
// https://fastspring.com/blog/how-to-format-30-currencies-from-countries-all-over-the-world/
//
func (nStrFmtCountry *NumStrFormatCountry) Malaysia() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 458
	setupDto.IdString = "458"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 458
	setupDto.CountryIdString = "458"
	setupDto.CountryDescription = "Country Setup - Malaysia"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Malaysia"
	setupDto.CountryAbbreviatedName = "Malaysia"

	setupDto.CountryAlternateNames =
		[]string{
			"Malaysia",
		}

	setupDto.CountryCodeTwoChar = "MY"
	setupDto.CountryCodeThreeChar = "MYS"
	setupDto.CountryCodeNumber = "458"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$ 127.54"
	setupDto.CurrencyNegativeValueFmt = "$ -127.54"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "MYR"
	setupDto.CurrencyCodeNo = "458"
	setupDto.CurrencyName = "Ringgit"

	setupDto.CurrencySymbols = []rune{
		'\U00000052',
		'\U0000004d',
	}

	setupDto.MinorCurrencyName = "Sen"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 5)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = ','
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "-127.54"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Malta - Returns the number string format used in
// The Republic of Malta.
//
//  https://fastspring.com/blog/how-to-format-30-currencies-from-countries-all-over-the-world/
//  https://en.wikipedia.org/wiki/Decimal_separator
//
func (nStrFmtCountry *NumStrFormatCountry) Malta() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 470
	setupDto.IdString = "470"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 470
	setupDto.CountryIdString = "470"
	setupDto.CountryDescription = "Country Setup - Malta"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Malta"
	setupDto.CountryAbbreviatedName = "Malta"

	setupDto.CountryAlternateNames =
		[]string{
			"The Republic of Malta",
			"Republic of Malta"}

	setupDto.CountryCodeTwoChar = "MT"
	setupDto.CountryCodeThreeChar = "MLT"
	setupDto.CountryCodeNumber = "470"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "127.54- $"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "LUF"
	setupDto.CurrencyCodeNo = "978"
	setupDto.CurrencyName = "Euro"
	setupDto.CurrencySymbols = []rune{'\U000020ac'}

	setupDto.MinorCurrencyName = "Cent"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = ','
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "127.54-"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Mexico - Returns the number string format used in
// The United Mexican States.
//
func (nStrFmtCountry *NumStrFormatCountry) Mexico() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 484
	setupDto.IdString = "484"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 484
	setupDto.CountryIdString = "484"
	setupDto.CountryDescription = "Country Setup - Mexico"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Mexico"
	setupDto.CountryAbbreviatedName = "Mexico"

	setupDto.CountryAlternateNames =
		[]string{
			"The United Mexican States",
			"United Mexican States"}

	setupDto.CountryCodeTwoChar = "MX"
	setupDto.CountryCodeThreeChar = "MEX"
	setupDto.CountryCodeNumber = "484"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$127.54"
	setupDto.CurrencyNegativeValueFmt = "$-127.54"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "MXN"
	setupDto.CurrencyCodeNo = "484"
	setupDto.CurrencyName = "Peso"
	setupDto.CurrencySymbols = []rune{'\U00000024'}

	setupDto.MinorCurrencyName = "Centavo"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 5)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = ','
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "-127.54"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Morocco - Returns the number string format used in
// The Kingdom of Morocco.
//
// https://en.wikipedia.org/wiki/ISO_4217
// https://en.wikipedia.org/wiki/Currency_symbol
// https://fastspring.com/blog/how-to-format-30-currencies-from-countries-all-over-the-world/
//
func (nStrFmtCountry *NumStrFormatCountry) Morocco() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 504
	setupDto.IdString = "504"
	setupDto.Description = "Country Setup - Morocco"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 504
	setupDto.CountryIdString = "504"
	setupDto.CountryDescription = "Country Setup - Morocco"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Morocco"
	setupDto.CountryAbbreviatedName = "Morocco"

	setupDto.CountryAlternateNames =
		[]string{
			"The Kingdom of Morocco",
			"Kingdom of Morocco"}

	setupDto.CountryCodeTwoChar = "MA"
	setupDto.CountryCodeThreeChar = "MAR"
	setupDto.CountryCodeNumber = "504"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "127.54- $"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "MAD"
	setupDto.CurrencyCodeNo = "504"
	setupDto.CurrencyName = "Dirham"

	setupDto.CurrencySymbols = []rune{
		'\U0000002e',
		'\U0000062f',
		'\U0000002e',
		'\U00000645',
	}

	setupDto.MinorCurrencyName = "Centime"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = ','
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "127.54-"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Namibia - Returns the number string format used in
// The Republic of Namibia.
//
//  https://www.xe.com/currency/nad-namibian-dollar
//  https://en.wikipedia.org/wiki/Decimal_separator
//
func (nStrFmtCountry *NumStrFormatCountry) Namibia() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 516
	setupDto.IdString = "516"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 516
	setupDto.CountryIdString = "516"
	setupDto.CountryDescription = "Country Setup - Namibia"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Namibia"
	setupDto.CountryAbbreviatedName = "Namibia"

	setupDto.CountryAlternateNames =
		[]string{
			"The Republic of Namibia",
			"Republic of Namibia"}

	setupDto.CountryCodeTwoChar = "NA"
	setupDto.CountryCodeThreeChar = "NAM"
	setupDto.CountryCodeNumber = "516"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$ 127.54"
	setupDto.CurrencyNegativeValueFmt = "$ -127.54"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "NAD"
	setupDto.CurrencyCodeNo = "516"
	setupDto.CurrencyName = "Dollar"
	setupDto.CurrencySymbols = []rune{'\U00000024'}

	setupDto.MinorCurrencyName = "Cent"
	setupDto.MinorCurrencySymbols = []rune{'c'}

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = ','
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "-127.54"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Nepal - Returns the number string format used in
// The Federal Democratic Republic of Nepal.
//
//  https://en.wikipedia.org/wiki/Decimal_separator
//  https://www.xe.com/currency/npr-nepalese-rupee
//
func (nStrFmtCountry *NumStrFormatCountry) Nepal() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 524
	setupDto.IdString = "524"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 524
	setupDto.CountryIdString = "524"
	setupDto.CountryDescription = "Country Setup - Nepal"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Nepal"
	setupDto.CountryAbbreviatedName = "Nepal"

	setupDto.CountryAlternateNames =
		[]string{
			"The Federal Democratic Republic of Nepal",
			"Federal Democratic Republic of Nepal"}

	setupDto.CountryCodeTwoChar = "NP"
	setupDto.CountryCodeThreeChar = "NPL"
	setupDto.CountryCodeNumber = "524"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$ 127.54"
	setupDto.CurrencyNegativeValueFmt = "$ -127.54"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "NPR"
	setupDto.CurrencyCodeNo = "524"
	setupDto.CurrencyName = "Rupee"
	setupDto.CurrencySymbols = []rune{'\U000020a8'}

	setupDto.MinorCurrencyName = "Paisa"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = ','
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3, 2}

	setupDto.SignedNumValPositiveValueFmt = ""
	setupDto.SignedNumValNegativeValueFmt = ""
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1
	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Netherlands - Returns the number string format used in the
// Kingdom of Netherlands.
//
// https://en.wikipedia.org/wiki/ISO_4217
// https://en.wikipedia.org/wiki/Currency_symbol
// https://docs.microsoft.com/en-us/globalization/locale/currency-formatting
//
func (nStrFmtCountry *NumStrFormatCountry) Netherlands() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 528
	setupDto.IdString = "528"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 528
	setupDto.CountryIdString = "528"
	setupDto.CountryDescription = "Country Setup - Netherlands"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Netherlands"
	setupDto.CountryAbbreviatedName = "Netherlands"

	setupDto.CountryAlternateNames =
		[]string{
			"The Kingdom of the Netherlands",
			"Kingdom of the Netherlands"}

	setupDto.CountryCodeTwoChar = "NL"
	setupDto.CountryCodeThreeChar = "NLD"
	setupDto.CountryCodeNumber = "528"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$ 127.54"
	setupDto.CurrencyNegativeValueFmt = " $ 127.54-"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "NLG"
	setupDto.CurrencyCodeNo = "978"
	setupDto.CurrencyName = "Euro"
	setupDto.CurrencySymbols = []rune{'\U000020ac'}

	setupDto.MinorCurrencyName = "Cent"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = ','
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "-127.54"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// NewZealand - Returns the number string format used in
// New Zealand.
//
func (nStrFmtCountry *NumStrFormatCountry) NewZealand() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 554
	setupDto.IdString = "554"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 554
	setupDto.CountryIdString = "554"
	setupDto.CountryDescription = "Country Setup - New Zealand"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "New Zealand"
	setupDto.CountryAbbreviatedName = "New Zealand"

	setupDto.CountryAlternateNames =
		[]string{
			"New Zealand",
		}

	setupDto.CountryCodeTwoChar = "NZ"
	setupDto.CountryCodeThreeChar = "NZL"
	setupDto.CountryCodeNumber = "554"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$ 127.54"
	setupDto.CurrencyNegativeValueFmt = "$ -127.54"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "NZD"
	setupDto.CurrencyCodeNo = "554"
	setupDto.CurrencyName = "Dollar"
	setupDto.CurrencySymbols = []rune{'\U00000024'}

	setupDto.MinorCurrencyName = "Cent"
	setupDto.MinorCurrencySymbols = []rune{'\U000000a2'}

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = ','
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "-127.54"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Norway - Returns the number string format used in the
// Kingdom of Norway.
//
//  https://en.wikipedia.org/wiki/ISO_4217
//  https://en.wikipedia.org/wiki/Currency_symbol
//  https://en.wikipedia.org/wiki/Decimal_separator
//
func (nStrFmtCountry *NumStrFormatCountry) Norway() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 578
	setupDto.IdString = "578"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 578
	setupDto.CountryIdString = "578"
	setupDto.CountryDescription = "Country Setup - Norway"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Norway"
	setupDto.CountryAbbreviatedName = "Norway"

	setupDto.CountryAlternateNames =
		[]string{
			"The Kingdom of Norway",
			"Kingdom of Norway"}

	setupDto.CountryCodeTwoChar = "NO"
	setupDto.CountryCodeThreeChar = "NOR"
	setupDto.CountryCodeNumber = "578"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "127.54- $"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "NOK"
	setupDto.CurrencyCodeNo = "578"
	setupDto.CurrencyName = "Krone"
	setupDto.CurrencySymbols = []rune{'\U0000006b', '\U00000072'}

	setupDto.MinorCurrencyName = "øre"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = ' ' // Space
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "127.54-"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Oman - Returns the number string format used in
// The Sultanate of Oman.
//
// https://en.wikipedia.org/wiki/ISO_4217
// https://en.wikipedia.org/wiki/Currency_symbol
// https://www.xe.com/currency/omr-omani-rial
//
func (nStrFmtCountry *NumStrFormatCountry) Oman() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 512
	setupDto.IdString = "512"
	setupDto.Description = "Country Setup - Oman"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 512
	setupDto.CountryIdString = "512"
	setupDto.CountryDescription = "Country Setup - Oman"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Oman"
	setupDto.CountryAbbreviatedName = "Oman"

	setupDto.CountryAlternateNames =
		[]string{
			"The Sultanate of Oman",
			"Sultanate of Oman"}

	setupDto.CountryCodeTwoChar = "OM"
	setupDto.CountryCodeThreeChar = "OMN"
	setupDto.CountryCodeNumber = "512"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "127.54- $"
	setupDto.CurrencyDecimalDigits = 3
	setupDto.CurrencyCode = "OMR"
	setupDto.CurrencyCodeNo = "512"
	setupDto.CurrencyName = "Rial"
	setupDto.CurrencySymbols = []rune{'\U0000fdfc'}

	setupDto.MinorCurrencyName = "Baiza"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = ','
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "127.54-"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Pakistan - Returns the number string format used in
// The Islamic Republic of Pakistan.
//
// https://en.wikipedia.org/wiki/ISO_4217
// https://en.wikipedia.org/wiki/Currency_symbol
// https://freeformatter.com/pakistan-standards-code-snippets.html
//
//
func (nStrFmtCountry *NumStrFormatCountry) Pakistan() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 586
	setupDto.IdString = "586"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 586
	setupDto.CountryIdString = "586"
	setupDto.CountryDescription = "Country Setup - Pakistan"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Pakistan"
	setupDto.CountryAbbreviatedName = "Pakistan"

	setupDto.CountryAlternateNames =
		[]string{
			"The Islamic Republic of Pakistan",
			"Islamic Republic of Pakistan"}

	setupDto.CountryCodeTwoChar = "PK"
	setupDto.CountryCodeThreeChar = "PAK"
	setupDto.CountryCodeNumber = "586"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$127.54"
	setupDto.CurrencyNegativeValueFmt = "$-127.54"
	setupDto.CurrencyDecimalDigits = 0
	setupDto.CurrencyCode = "PKR"
	setupDto.CurrencyCodeNo = "586"
	setupDto.CurrencyName = "Rupee"
	setupDto.CurrencySymbols = []rune{'\U000020a8'}

	setupDto.MinorCurrencyName = "Paisa"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = ','
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "-127.54"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Peru - Returns the number string format used in
// The Republic of Perú.
//
// https://en.wikipedia.org/wiki/Decimal_separator
//
func (nStrFmtCountry *NumStrFormatCountry) Peru() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 604
	setupDto.IdString = "604"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 604
	setupDto.CountryIdString = "604"
	setupDto.CountryDescription = "Country Setup - Peru"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Peru"
	setupDto.CountryAbbreviatedName = "Peru"

	setupDto.CountryAlternateNames =
		[]string{
			"The Republic of Perú",
			"Republic of Perú"}

	setupDto.CountryCodeTwoChar = "PE"
	setupDto.CountryCodeThreeChar = "PER"
	setupDto.CountryCodeNumber = "604"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$ 127.54"
	setupDto.CurrencyNegativeValueFmt = "$ -127.54"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "PEN"
	setupDto.CurrencyCodeNo = "604"
	setupDto.CurrencyName = "Sol"

	setupDto.CurrencySymbols = []rune{
		'\U00000053',
		'\U0000002f',
		'\U0000002e',
	}

	setupDto.MinorCurrencyName = "Centimo"
	setupDto.MinorCurrencySymbols =
		make([]rune, 0, 5)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = ' ' // space
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "-127.54"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Philippines - Returns the number string format used in
// The Republic of the Philippines.
//
func (nStrFmtCountry *NumStrFormatCountry) Philippines() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 608
	setupDto.IdString = "608"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 608
	setupDto.CountryIdString = "608"
	setupDto.CountryDescription = "Country Setup - Philippines"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Philippines"
	setupDto.CountryAbbreviatedName = "Philippines"

	setupDto.CountryAlternateNames =
		[]string{
			"The Republic of the Philippines",
			"The Philippines",
			"Republic of the Philippines"}

	setupDto.CountryCodeTwoChar = "PH"
	setupDto.CountryCodeThreeChar = "PHL"
	setupDto.CountryCodeNumber = "608"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$127.54"
	setupDto.CurrencyNegativeValueFmt = "$-127.54"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "PHP"
	setupDto.CurrencyCodeNo = "608"
	setupDto.CurrencyName = "Peso"
	setupDto.CurrencySymbols = []rune{'\U000020b1'}

	setupDto.MinorCurrencyName = "Centavo"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 5)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = ','
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "-127.54"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Poland - Returns the number string format used in
// The Republic of Poland.
//
// https://en.wikipedia.org/wiki/ISO_4217
// https://en.wikipedia.org/wiki/Currency_symbol
//
func (nStrFmtCountry *NumStrFormatCountry) Poland() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 616
	setupDto.IdString = "616"
	setupDto.Description = "Country Setup - Poland"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 616
	setupDto.CountryIdString = "616"
	setupDto.CountryDescription = "Country Setup - Poland"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Poland"
	setupDto.CountryAbbreviatedName = "Poland"

	setupDto.CountryAlternateNames =
		[]string{
			"The Republic of Poland",
			"Republic of Poland"}

	setupDto.CountryCodeTwoChar = "PL"
	setupDto.CountryCodeThreeChar = "POL"
	setupDto.CountryCodeNumber = "616"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "127.54- $"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "PLN"
	setupDto.CurrencyCodeNo = "985"
	setupDto.CurrencyName = "Zloty"
	setupDto.CurrencySymbols = []rune{'\U0000007a', '\U00000142'}

	setupDto.MinorCurrencyName = "Grosz"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = ' ' // Space
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "127.54-"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Portugal - Returns the number string format used in The
// Portuguese Republic.
//
//  https://en.wikipedia.org/wiki/ISO_4217
//  https://en.wikipedia.org/wiki/Currency_symbol
//  https://en.wikipedia.org/wiki/Decimal_separator
//
func (nStrFmtCountry *NumStrFormatCountry) Portugal() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 620
	setupDto.IdString = "620"
	setupDto.Description = "Country Setup - Portugal"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 620
	setupDto.CountryIdString = "620"
	setupDto.CountryDescription = "Country Setup - Portugal"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Portugal"
	setupDto.CountryAbbreviatedName = "Portugal"

	setupDto.CountryAlternateNames =
		[]string{
			"The Portuguese Republic",
			"Portuguese Republic"}

	setupDto.CountryCodeTwoChar = "PT"
	setupDto.CountryCodeThreeChar = "PRT"
	setupDto.CountryCodeNumber = "620"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "127.54- $"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "PTE"
	setupDto.CurrencyCodeNo = "978"
	setupDto.CurrencyName = "Euro"
	setupDto.CurrencySymbols = []rune{'\U000020ac'}

	setupDto.MinorCurrencyName = "Cent"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = ' ' // Space
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "127.54-"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Qatar - Returns the number string format used in
// The State of Qatar.
//
// https://en.wikipedia.org/wiki/ISO_4217
// https://en.wikipedia.org/wiki/Currency_symbol
// https://en.wikipedia.org/wiki/Decimal_separator
// https://www.xe.com/currency/qar-qatari-riyal
//
func (nStrFmtCountry *NumStrFormatCountry) Qatar() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 634
	setupDto.IdString = "634"
	setupDto.Description = "Country Setup - Qatar"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 634
	setupDto.CountryIdString = "634"
	setupDto.CountryDescription = "Country Setup - Qatar"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Qatar"
	setupDto.CountryAbbreviatedName = "Qatar"

	setupDto.CountryAlternateNames =
		[]string{
			"The State of Qatar",
			"State of Qatar"}

	setupDto.CountryCodeTwoChar = "QA"
	setupDto.CountryCodeThreeChar = "QAT"
	setupDto.CountryCodeNumber = "634"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "127.54- $"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "QAR"
	setupDto.CurrencyCodeNo = "634"
	setupDto.CurrencyName = "Riyal"
	setupDto.CurrencySymbols = []rune{'\U0000fdfc'}

	setupDto.MinorCurrencyName = "Dirham"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = ','
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "127.54-"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Romania - Returns the number string format used in the
// country of Romania.
//
// https://en.wikipedia.org/wiki/ISO_4217
// https://en.wikipedia.org/wiki/Currency_symbol
//
func (nStrFmtCountry *NumStrFormatCountry) Romania() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 642
	setupDto.IdString = "642"
	setupDto.Description = "Country Setup - Romania"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 642
	setupDto.CountryIdString = "642"
	setupDto.CountryDescription = "Country Setup - Romania"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Romania"
	setupDto.CountryAbbreviatedName = "Romania"

	setupDto.CountryAlternateNames =
		[]string{
			"Romania"}

	setupDto.CountryCodeTwoChar = "RO"
	setupDto.CountryCodeThreeChar = "ROU"
	setupDto.CountryCodeNumber = "642"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "127.54- $"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "RON"
	setupDto.CurrencyCodeNo = "946"
	setupDto.CurrencyName = "Lei"
	setupDto.CurrencySymbols = []rune{
		'\U0000006c',
		'\U00000065',
		'\U00000069'}

	setupDto.MinorCurrencyName = "Bani"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = '.'
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "127.54-"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Russia - Returns the number string format used in
// The Russian Federation.
//
//  https://en.wikipedia.org/wiki/ISO_4217
//  https://en.wikipedia.org/wiki/Currency_symbol
//
func (nStrFmtCountry *NumStrFormatCountry) Russia() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 643
	setupDto.IdString = "643"
	setupDto.Description = "Country Setup - Russia"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 643
	setupDto.CountryIdString = "643"
	setupDto.CountryDescription = "Country Setup - Russia"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Russia"
	setupDto.CountryAbbreviatedName = "Russia"

	setupDto.CountryAlternateNames =
		[]string{
			"The Russian Federation",
			"Russian Federation"}

	setupDto.CountryCodeTwoChar = "RU"
	setupDto.CountryCodeThreeChar = "RUS"
	setupDto.CountryCodeNumber = "643"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "127.54- $"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "RUB"
	setupDto.CurrencyCodeNo = "643"
	setupDto.CurrencyName = "Ruble"
	setupDto.CurrencySymbols = []rune{'\U000020bd'}

	setupDto.MinorCurrencyName = "Kopeck"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = ' ' // space
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "127.54-"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// SaudiArabia - Returns the number string format used in
// The Kingdom of Saudi Arabia.
//
// https://en.wikipedia.org/wiki/ISO_4217
// https://en.wikipedia.org/wiki/Currency_symbol
// https://fastspring.com/blog/how-to-format-30-currencies-from-countries-all-over-the-world/
//
func (nStrFmtCountry *NumStrFormatCountry) SaudiArabia() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 682
	setupDto.IdString = "682"
	setupDto.Description = "Country Setup - Saudi Arabia"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 682
	setupDto.CountryIdString = "682"
	setupDto.CountryDescription = "Country Setup - Saudi Arabia"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Saudi Arabia"
	setupDto.CountryAbbreviatedName = "Saudi Arabia"

	setupDto.CountryAlternateNames =
		[]string{
			"The Kingdom of Saudi Arabia",
			"Kingdom of Saudi Arabia"}

	setupDto.CountryCodeTwoChar = "SA"
	setupDto.CountryCodeThreeChar = "SAU"
	setupDto.CountryCodeNumber = "682"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "127.54- $"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "SAR"
	setupDto.CurrencyCodeNo = "682"
	setupDto.CurrencyName = "Riyal"
	setupDto.CurrencySymbols = []rune{'\U0000fdfc'}

	setupDto.MinorCurrencyName = "Halalat"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = ','
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "127.54-"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Serbia - Returns the number string format used in
// The Republic of Serbia.
//
//  https://en.wikipedia.org/wiki/Decimal_separator
//  https://www.xe.com/currency/rsd-serbian-dinar
//
func (nStrFmtCountry *NumStrFormatCountry) Serbia() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 688
	setupDto.IdString = "688"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 688
	setupDto.CountryIdString = "688"
	setupDto.CountryDescription = "Country Setup - Serbia"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Serbia"
	setupDto.CountryAbbreviatedName = "Serbia"

	setupDto.CountryAlternateNames =
		[]string{
			"The Republic of Serbia",
			"Republic of Serbia",
		}

	setupDto.CountryCodeTwoChar = "RS"
	setupDto.CountryCodeThreeChar = "SRB"
	setupDto.CountryCodeNumber = "688"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$ 127.54"
	setupDto.CurrencyNegativeValueFmt = "$ -127.54"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "RSD"
	setupDto.CurrencyCodeNo = "941"
	setupDto.CurrencyName = "Dinar"

	setupDto.CurrencySymbols = []rune{
		'\U00000414',
		'\U00000438',
		'\U0000043d',
		'\U0000002e',
	}

	setupDto.MinorCurrencyName = "Para"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = '.'
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "127.54-"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Singapore - Returns the number string format used in
// The Republic of Singapore.
//
// https://fastspring.com/blog/how-to-format-30-currencies-from-countries-all-over-the-world/
//
func (nStrFmtCountry *NumStrFormatCountry) Singapore() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 702
	setupDto.IdString = "702"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 702
	setupDto.CountryIdString = "702"
	setupDto.CountryDescription = "Country Setup - Singapore"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Singapore"
	setupDto.CountryAbbreviatedName = "Singapore"

	setupDto.CountryAlternateNames =
		[]string{
			"The Republic of Singapore",
			"Republic of Singapore",
		}

	setupDto.CountryCodeTwoChar = "SG"
	setupDto.CountryCodeThreeChar = "SGP"
	setupDto.CountryCodeNumber = "702"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$127.54"
	setupDto.CurrencyNegativeValueFmt = "($127.54)"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "SGD"
	setupDto.CurrencyCodeNo = "702"
	setupDto.CurrencyName = "Dollar"
	setupDto.CurrencySymbols = []rune{'\U00000024'}

	setupDto.MinorCurrencyName = "Cent"
	setupDto.MinorCurrencySymbols = []rune{'\U000000a2'}

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = ','
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "-127.54"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Slovakia - Returns the number string format used in
// The Slovak Republic.
//
//  https://fastspring.com/blog/how-to-format-30-currencies-from-countries-all-over-the-world/
//  https://en.wikipedia.org/wiki/Decimal_separator
//
func (nStrFmtCountry *NumStrFormatCountry) Slovakia() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 703
	setupDto.IdString = "703"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 703
	setupDto.CountryIdString = "703"
	setupDto.CountryDescription = "Country Setup - Slovakia"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Slovakia"
	setupDto.CountryAbbreviatedName = "Slovakia"

	setupDto.CountryAlternateNames =
		[]string{
			"The Slovak Republic",
			"Slovak Republic"}

	setupDto.CountryCodeTwoChar = "SK"
	setupDto.CountryCodeThreeChar = "SVK"
	setupDto.CountryCodeNumber = "703"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "-127.54 $"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "SKK"
	setupDto.CurrencyCodeNo = "978"
	setupDto.CurrencyName = "Euro"
	setupDto.CurrencySymbols = []rune{'\U000020ac'}

	setupDto.MinorCurrencyName = "Cent"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = ' ' // Space
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "127.54 -"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Slovenia - Returns the number string format used in
// The Republic of Slovenia.
//
//  https://fastspring.com/blog/how-to-format-30-currencies-from-countries-all-over-the-world/
//  https://en.wikipedia.org/wiki/Decimal_separator
//
func (nStrFmtCountry *NumStrFormatCountry) Slovenia() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 705
	setupDto.IdString = "705"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 705
	setupDto.CountryIdString = "705"
	setupDto.CountryDescription = "Country Setup - Slovenia"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Slovenia"
	setupDto.CountryAbbreviatedName = "Slovenia"

	setupDto.CountryAlternateNames =
		[]string{
			"The Republic of Slovenia",
			"Republic of Slovenia"}

	setupDto.CountryCodeTwoChar = "SI"
	setupDto.CountryCodeThreeChar = "SVN"
	setupDto.CountryCodeNumber = "705"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "127.54- $"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "SIT"
	setupDto.CurrencyCodeNo = "978"
	setupDto.CurrencyName = "Euro"
	setupDto.CurrencySymbols = []rune{'\U000020ac'}

	setupDto.MinorCurrencyName = "Cent"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = '.'
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "127.54-"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// SouthAfrica - Returns the number string format used in
// The Republic of South Africa.
//
//  https://en.wikipedia.org/wiki/Decimal_separator
//
func (nStrFmtCountry *NumStrFormatCountry) SouthAfrica() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 710
	setupDto.IdString = "710"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 710
	setupDto.CountryIdString = "710"
	setupDto.CountryDescription = "Country Setup - South Africa"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "South Africa"
	setupDto.CountryAbbreviatedName = "South Africa"

	setupDto.CountryAlternateNames =
		[]string{
			"The Republic of South Africa",
			"Republic of South Africa",
		}

	setupDto.CountryCodeTwoChar = "ZA"
	setupDto.CountryCodeThreeChar = "ZAF"
	setupDto.CountryCodeNumber = "710"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$127.54"
	setupDto.CurrencyNegativeValueFmt = "($127.54)"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "ZAR"
	setupDto.CurrencyCodeNo = "710"
	setupDto.CurrencyName = "Rand"
	setupDto.CurrencySymbols = []rune{'\U00000052'}

	setupDto.MinorCurrencyName = "Cent"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 5)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = ' ' // Space
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "-127.54"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Spain - Returns the number string format used in The Kingdom
// of Spain.
//
//  https://en.wikipedia.org/wiki/ISO_4217
//  https://en.wikipedia.org/wiki/Currency_symbol
//  https://en.wikipedia.org/wiki/Decimal_separator
//
func (nStrFmtCountry *NumStrFormatCountry) Spain() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 724
	setupDto.IdString = "724"
	setupDto.Description = "Country Setup - Spain"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 724
	setupDto.CountryIdString = "724"
	setupDto.CountryDescription = "Country Setup - Spain"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Spain"
	setupDto.CountryAbbreviatedName = "Spain"

	setupDto.CountryAlternateNames =
		[]string{
			"The Kingdom of Spain",
			"Kingdom of Spain"}

	setupDto.CountryCodeTwoChar = "ES"
	setupDto.CountryCodeThreeChar = "ESP"
	setupDto.CountryCodeNumber = "724"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "127.54- $"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "ESP"
	setupDto.CurrencyCodeNo = "978"
	setupDto.CurrencyName = "Euro"
	setupDto.CurrencySymbols = []rune{'\U000020ac'}

	setupDto.MinorCurrencyName = "Cent"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = ' ' // Space
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "127.54-"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// SriLanka - Sri Lanka. Returns the number string format used in
// The Democratic Socialist Republic of Sri Lanka.
//
//  https://www.srilankalocaltours.com/sri-lanka-currency/
//  https://www.xe.com/currency/lkr-sri-lankan-rupee
//
func (nStrFmtCountry *NumStrFormatCountry) SriLanka() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 144
	setupDto.IdString = "144"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 144
	setupDto.CountryIdString = "144"
	setupDto.CountryDescription = "Country Setup - Sri Lanka"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Sri Lanka"
	setupDto.CountryAbbreviatedName = "Sri Lanka"

	setupDto.CountryAlternateNames =
		[]string{
			"The Democratic Socialist Republic of Sri Lanka",
			"Democratic Socialist Republic of Sri Lanka"}

	setupDto.CountryCodeTwoChar = "LK"
	setupDto.CountryCodeThreeChar = "LKA"
	setupDto.CountryCodeNumber = "144"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$ 127.54"
	setupDto.CurrencyNegativeValueFmt = "$ -127.54"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "LKR"
	setupDto.CurrencyCodeNo = "144"
	setupDto.CurrencyName = "Rupee"
	setupDto.CurrencySymbols = []rune{'\U000020a8'}

	setupDto.MinorCurrencyName = "Cent"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = '.'
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "-127.54"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Sweden - Returns the number string format used in the
// Kingdom of Sweden.
//
//  https://en.wikipedia.org/wiki/ISO_4217
//  https://en.wikipedia.org/wiki/Currency_symbol
//  https://en.wikipedia.org/wiki/Decimal_separator
//
func (nStrFmtCountry *NumStrFormatCountry) Sweden() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 752
	setupDto.IdString = "752"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 752
	setupDto.CountryIdString = "752"
	setupDto.CountryDescription = "Country Setup - Sweden"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Sweden"
	setupDto.CountryAbbreviatedName = "Sweden"

	setupDto.CountryAlternateNames =
		[]string{
			"The Kingdom of Sweden",
			"Kingdom of Sweden"}

	setupDto.CountryCodeTwoChar = "SE"
	setupDto.CountryCodeThreeChar = "SWE"
	setupDto.CountryCodeNumber = "752"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "127.54- $"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "SEK"
	setupDto.CurrencyCodeNo = "752"
	setupDto.CurrencyName = "Krona"
	setupDto.CurrencySymbols = []rune{'\U0000006b', '\U00000072'}

	setupDto.MinorCurrencyName = "øre"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = ' ' // Space
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "127.54-"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Switzerland - Returns the number string format used in
// The Swiss Confederation.
//
// https://en.wikipedia.org/wiki/ISO_4217
// https://en.wikipedia.org/wiki/Currency_symbol
//
func (nStrFmtCountry *NumStrFormatCountry) Switzerland() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 756
	setupDto.IdString = "756"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 756
	setupDto.CountryIdString = "756"
	setupDto.CountryDescription = "Country Setup - Switzerland"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Switzerland"
	setupDto.CountryAbbreviatedName = "Switzerland"

	setupDto.CountryAlternateNames =
		[]string{
			"The Swiss Confederation",
			"Swiss Confederation"}

	setupDto.CountryCodeTwoChar = "CH"
	setupDto.CountryCodeThreeChar = "CHE"
	setupDto.CountryCodeNumber = "756"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "127.54- $"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "CHF"
	setupDto.CurrencyCodeNo = "756"
	setupDto.CurrencyName = "Franc"
	setupDto.CurrencySymbols = []rune{'\U00000043', '\U00000048', '\U00000046'}

	setupDto.MinorCurrencyName = "Rappen"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = '\U00000027' // Apostrophe
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "127.54-"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Taiwan - Returns the number string format used in
// The Republic of China.
//
// https://fastspring.com/blog/how-to-format-30-currencies-from-countries-all-over-the-world/
//
func (nStrFmtCountry *NumStrFormatCountry) Taiwan() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 158
	setupDto.IdString = "158"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 158
	setupDto.CountryIdString = "158"
	setupDto.CountryDescription = "Country Setup - Taiwan"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Taiwan"
	setupDto.CountryAbbreviatedName = "Taiwan"

	setupDto.CountryAlternateNames =
		[]string{
			"The Republic of China",
			"Republic of China",
		}

	setupDto.CountryCodeTwoChar = "TW"
	setupDto.CountryCodeThreeChar = "TWN"
	setupDto.CountryCodeNumber = "158"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$127.54"
	setupDto.CurrencyNegativeValueFmt = "($127.54)"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "TWD"
	setupDto.CurrencyCodeNo = "901"
	setupDto.CurrencyName = "New Dollar"
	setupDto.CurrencySymbols = []rune{'\U00005143'}

	setupDto.MinorCurrencyName = ""
	setupDto.MinorCurrencySymbols = make([]rune, 0, 5)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = ','
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "-127.54"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Thailand - Returns the number string format used in
// The Kingdom of Thailand.
//
func (nStrFmtCountry *NumStrFormatCountry) Thailand() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 764
	setupDto.IdString = "764"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 764
	setupDto.CountryIdString = "764"
	setupDto.CountryDescription = "Country Setup - Thailand"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Thailand"
	setupDto.CountryAbbreviatedName = "Thailand"

	setupDto.CountryAlternateNames =
		[]string{
			"The Kingdom of Thailand",
			"Kingdom of Thailand",
		}

	setupDto.CountryCodeTwoChar = "TH"
	setupDto.CountryCodeThreeChar = "THA"
	setupDto.CountryCodeNumber = "764"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "127.54- $"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "THB"
	setupDto.CurrencyCodeNo = "764"
	setupDto.CurrencyName = "Baht"
	setupDto.CurrencySymbols = []rune{'\U00000e3f'}

	setupDto.MinorCurrencyName = "Satang"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 5)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = ','
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "127.54-"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Turkey - Returns the number string format used in
// The Republic of Turkey.
//
// https://en.wikipedia.org/wiki/ISO_4217
// https://en.wikipedia.org/wiki/Currency_symbol
//
func (nStrFmtCountry *NumStrFormatCountry) Turkey() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 792
	setupDto.IdString = "792"
	setupDto.Description = "Country Setup - Turkey"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 792
	setupDto.CountryIdString = "792"
	setupDto.CountryDescription = "Country Setup - Turkey"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Turkey"
	setupDto.CountryAbbreviatedName = "Turkey"

	setupDto.CountryAlternateNames =
		[]string{
			"The Republic of Turkey",
			"Republic of Turkey"}

	setupDto.CountryCodeTwoChar = "TR"
	setupDto.CountryCodeThreeChar = "TUR"
	setupDto.CountryCodeNumber = "792"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "127.54- $"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "TRY"
	setupDto.CurrencyCodeNo = "949"
	setupDto.CurrencyName = "Lira"
	setupDto.CurrencySymbols = []rune{'\U000020ba'}

	setupDto.MinorCurrencyName = "Kurus"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = '.'
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "127.54-"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Ukraine - Returns the number string format used in
// the country of Ukraine.
//
// https://en.wikipedia.org/wiki/ISO_4217
// https://en.wikipedia.org/wiki/Currency_symbol
//
func (nStrFmtCountry *NumStrFormatCountry) Ukraine() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 804
	setupDto.IdString = "804"
	setupDto.Description = "Country Setup - Ukraine"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 804
	setupDto.CountryIdString = "804"
	setupDto.CountryDescription = "Country Setup - Ukraine"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Ukraine"
	setupDto.CountryAbbreviatedName = "Ukraine"

	setupDto.CountryAlternateNames =
		[]string{
			"Ukraine"}

	setupDto.CountryCodeTwoChar = "UA"
	setupDto.CountryCodeThreeChar = "UKR"
	setupDto.CountryCodeNumber = "804"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "127.54- $"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "UAH"
	setupDto.CurrencyCodeNo = "980"
	setupDto.CurrencyName = "Hryvnia"
	setupDto.CurrencySymbols = []rune{'\U000020b4'}

	setupDto.MinorCurrencyName = "Kopiyka"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = ' ' // space
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "127.54-"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// UnitedArabEmirates - Returns the number string format used in
// The United Arab Emirates.
//
// https://en.wikipedia.org/wiki/ISO_4217
// https://en.wikipedia.org/wiki/Currency_symbol
// https://www.xe.com/currency/aed-emirati-dirham
//
func (nStrFmtCountry *NumStrFormatCountry) UnitedArabEmirates() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 784
	setupDto.IdString = "784"
	setupDto.Description = "Country Setup - UnitedArabEmirates"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 784
	setupDto.CountryIdString = "784"
	setupDto.CountryDescription = "Country Setup - United Arab Emirates"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "United Arab Emirates"
	setupDto.CountryAbbreviatedName = "United Arab Emirates"

	setupDto.CountryAlternateNames =
		[]string{
			"The United Arab Emirates",
		}

	setupDto.CountryCodeTwoChar = "AE"
	setupDto.CountryCodeThreeChar = "ARE"
	setupDto.CountryCodeNumber = "784"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "127.54 $"
	setupDto.CurrencyNegativeValueFmt = "127.54- $"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "AED"
	setupDto.CurrencyCodeNo = "784"
	setupDto.CurrencyName = "Dirham"
	setupDto.CurrencySymbols = []rune{
		'\U00000625',
		'\U0000002e',
		'\U0000062f',
	}

	setupDto.MinorCurrencyName = "Fil"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = ','
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "127.54-"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// UnitedKingdom - Returns the number string format used in the
// United Kingdom of Great Britain and Northern Ireland.
//
// https://docs.microsoft.com/en-us/globalization/locale/currency-formatting
//
func (nStrFmtCountry *NumStrFormatCountry) UnitedKingdom() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 826
	setupDto.IdString = "826"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 826
	setupDto.CountryIdString = "826"
	setupDto.CountryDescription = "Country Setup - United Kingdom"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "United Kingdom"
	setupDto.CountryAbbreviatedName = "UK"

	setupDto.CountryAlternateNames =
		[]string{
			"United Kingdom of Great Britain and Northern Ireland",
			"England",
			"Great Britain"}

	setupDto.CountryCodeTwoChar = "GB"
	setupDto.CountryCodeThreeChar = "GBR"
	setupDto.CountryCodeNumber = "826"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$127.54"
	setupDto.CurrencyNegativeValueFmt = "-$127.54"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "GBP"
	setupDto.CurrencyCodeNo = "826"
	setupDto.CurrencyName = "Pound"
	setupDto.CurrencySymbols = []rune{'\U000000a3'}

	setupDto.MinorCurrencyName = "Pence"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = ','
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "-127.54"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1
	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// UnitedStates - Returns the number string format used in the
// United States.
//
func (nStrFmtCountry *NumStrFormatCountry) UnitedStates() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 840
	setupDto.IdString = "840"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 840
	setupDto.CountryIdString = "840"
	setupDto.CountryDescription = "Country Setup - United States"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "United States"
	setupDto.CountryAbbreviatedName = "USA"

	setupDto.CountryAlternateNames =
		[]string{
			"The United States of America",
			"United States of America",
			"America"}

	setupDto.CountryCodeTwoChar = "US"
	setupDto.CountryCodeThreeChar = "USA"
	setupDto.CountryCodeNumber = "840"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$127.54"
	setupDto.CurrencyNegativeValueFmt = "($127.54)"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "USD"
	setupDto.CurrencyCodeNo = "840"
	setupDto.CurrencyName = "Dollar"
	setupDto.CurrencySymbols = []rune{'\U00000024'}

	setupDto.MinorCurrencyName = "Cent"
	setupDto.MinorCurrencySymbols = []rune{'\U000000a2'}

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = '.'
	setupDto.IntegerDigitsSeparator = ','
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "-127.54"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SciNotNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// Venezuela - Returns the number string format used in
// The Bolivarian Republic of Venezuela.
//
// https://support.sas.com/documentation/cdl/en/nlsref/61893/HTML/default/viewer.htm#a003090801.htm
//
func (nStrFmtCountry *NumStrFormatCountry) Venezuela() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 862
	setupDto.IdString = "862"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 862
	setupDto.CountryIdString = "862"
	setupDto.CountryDescription = "Country Setup - Venezuela"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Venezuela"
	setupDto.CountryAbbreviatedName = "Venezuela"

	setupDto.CountryAlternateNames =
		[]string{
			"The Bolivarian Republic of Venezuela",
			"Bolivarian Republic of Venezuela"}

	setupDto.CountryCodeTwoChar = "VE"
	setupDto.CountryCodeThreeChar = "VEN"
	setupDto.CountryCodeNumber = "862"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$ 127.54"
	setupDto.CurrencyNegativeValueFmt = "$ -127.54"
	setupDto.CurrencyDecimalDigits = 2
	setupDto.CurrencyCode = "VES"
	setupDto.CurrencyCodeNo = "862"
	setupDto.CurrencyName = "Bolivar"

	setupDto.CurrencySymbols = []rune{
		'\U00000042',
		'\U00000073',
	}

	setupDto.MinorCurrencyName = "Centimo"
	setupDto.MinorCurrencySymbols =
		make([]rune, 0, 5)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = '.'
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "-127.54"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}

// VietNam - Returns the number string format used in
// The Socialist Republic of Viet Nam.
//
//  https://en.wikipedia.org/wiki/ISO_4217
//  https://en.wikipedia.org/wiki/Currency_symbol
//  https://en.wikipedia.org/wiki/Decimal_separator
//
func (nStrFmtCountry *NumStrFormatCountry) VietNam() (setupDto NumStrFmtSpecSetupDto) {

	if nStrFmtCountry.lock == nil {
		nStrFmtCountry.lock = new(sync.Mutex)
	}

	nStrFmtCountry.lock.Lock()

	defer nStrFmtCountry.lock.Unlock()

	setupDto.Lock = new(sync.Mutex)

	setupDto.IdNo = 704
	setupDto.IdString = "704"
	setupDto.Description = "Country Setup"
	setupDto.Tag = ""
	setupDto.CountryIdNo = 704
	setupDto.CountryIdString = "704"
	setupDto.CountryDescription = "Country Setup - Viet Nam"
	setupDto.CountryTag = ""
	setupDto.CountryCultureName = "Viet Nam"
	setupDto.CountryAbbreviatedName = "Viet Nam"

	setupDto.CountryAlternateNames =
		[]string{
			"The Socialist Republic of Viet Nam",
			"Socialist Republic of Viet Nam"}

	setupDto.CountryCodeTwoChar = "VN"
	setupDto.CountryCodeThreeChar = "VNM"
	setupDto.CountryCodeNumber = "704"

	setupDto.AbsoluteValFmt = "127.54"
	setupDto.AbsoluteValTurnOnIntegerDigitsSeparation = true
	setupDto.AbsoluteValNumFieldLen = -1

	setupDto.AbsoluteValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.CurrencyPositiveValueFmt = "$127.54"
	setupDto.CurrencyNegativeValueFmt = "127.54- $"
	setupDto.CurrencyDecimalDigits = 0
	setupDto.CurrencyCode = "VND"
	setupDto.CurrencyCodeNo = "704"
	setupDto.CurrencyName = "Dong"
	setupDto.CurrencySymbols = []rune{'\U000020ab'}

	setupDto.MinorCurrencyName = "Hao,Xu"
	setupDto.MinorCurrencySymbols = make([]rune, 0, 10)

	setupDto.CurrencyTurnOnIntegerDigitsSeparation = true
	setupDto.CurrencyNumFieldLen = -1

	setupDto.CurrencyNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.DecimalSeparator = ','
	setupDto.IntegerDigitsSeparator = '.'
	setupDto.IntegerDigitsGroupingSequence =
		[]uint{3}

	setupDto.SignedNumValPositiveValueFmt = "127.54"
	setupDto.SignedNumValNegativeValueFmt = "127.54-"
	setupDto.SignedNumValTurnOnIntegerDigitsSeparation = true
	setupDto.SignedNumValNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	setupDto.SciNotSignificandUsesLeadingPlus = false
	setupDto.SciNotMantissaLength = 6
	setupDto.SciNotExponentChar = 'E'
	setupDto.SciNotExponentUsesLeadingPlus = true
	setupDto.SciNotNumFieldLen = -1

	setupDto.SignedNumValNumFieldTextJustify =
		TextJustify(0).Right()

	return setupDto
}
