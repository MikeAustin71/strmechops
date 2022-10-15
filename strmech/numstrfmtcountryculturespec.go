package strmech

import "sync"

type NumStrFmtCountryCultureSpec struct {
	IdNo                                      uint64
	IdString                                  string
	Description                               string
	Tag                                       string
	CountryIdNo                               uint64
	CountryIdString                           string
	CountryDescription                        string
	CountryTag                                string
	CountryCultureName                        string
	CountryAbbreviatedName                    string
	CountryAlternateNames                     []string
	CountryCodeTwoChar                        string
	CountryCodeThreeChar                      string
	CountryCodeNumber                         string
	CurrencyCode                              string
	CurrencyCodeNo                            string
	CurrencyName                              string
	CurrencySymbols                           []rune
	MinorCurrencyName                         string
	MinorCurrencySymbols                      []rune
	CurrencyTurnOnIntegerDigitsSeparation     bool
	CurrencyNumDecSep                         DecimalSeparatorSpec
	CurrencyIntGroupingSpec                   IntegerSeparatorSpec
	CurrencyPositiveValueFmt                  NumStrNumberSymbolSpec
	CurrencyNegativeValueFmt                  NumStrNumberSymbolSpec
	CurrencyZeroValueFmt                      NumStrNumberSymbolSpec
	SignedNumValTurnOnIntegerDigitsSeparation bool
	SignedNumValDecSep                        DecimalSeparatorSpec
	SignedNumValIntGroupingSpec               IntegerSeparatorSpec
	SignedNumValPositiveValueFmt              NumStrNumberSymbolSpec
	SignedNumValNegativeValueFmt              NumStrNumberSymbolSpec
	SignedNumValZeroValueFmt                  NumStrNumberSymbolSpec
	lock                                      *sync.Mutex
}
