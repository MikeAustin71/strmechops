package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// numStrFmtCountryCultureSpecMech
//
// Provides helper methods for type
// NumStrFmtCountryCultureSpec
type numStrFmtCountryCultureSpecMech struct {
	lock *sync.Mutex
}

//	copyCountryCulture
//
//	Copies all data from input parameter 'sourceSpec'
//	to input parameter 'destinationSpec'. Both instances
//	are of type NumStrFmtCountryCultureSpec.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields in 'destinationSpec'
//	will be deleted and overwritten.
//
//	Also, NO data validation is performed on 'sourceSpec'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	destinationSpec				*NumStrFmtCountryCultureSpec
//
//		A pointer to a NumStrFmtCountryCultureSpec instance.
//		All the member variable data fields in this object
//		will be replaced by data values copied from input
//		parameter 'sourceSignedNumFmtSpec'.
//
//		'destinationSpec' is the destination for this
//		copy operation.
//
//	sourceSpec			*NumStrFormatSpec
//
//		A pointer to another *NumStrFmtCountryCultureSpec
//		instance. All the member variable data values from this
//		object will be copied to corresponding member variables in
//		'destinationSpec'.
//
//		'sourceSpec' is the source for this copy operation.
//
//		No data validation is performed on 'sourceSpec'.
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string
//		which is included in all returned error
//		messages. Usually, it contains the name of the
//		calling method or methods listed as a function
//		chain.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref'
//		software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	err							error
//
//		If this method completes successfully, this
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (nStrFmtCountryCultureMech *numStrFmtCountryCultureSpecMech) copyCountryCulture(
	destinationSpec *NumStrFmtCountryCultureSpec,
	sourceSpec *NumStrFmtCountryCultureSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if nStrFmtCountryCultureMech.lock == nil {
		nStrFmtCountryCultureMech.lock = new(sync.Mutex)
	}

	nStrFmtCountryCultureMech.lock.Lock()

	defer nStrFmtCountryCultureMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"nStrFmtCountryCultureMech."+
			"copyCountryCulture()",
		"")

	if err != nil {
		return err
	}

	if destinationSpec == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'destinationSpec' is invalid!\n"+
			"'destinationSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if sourceSpec == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'sourceSpec' is invalid!\n"+
			"'sourceSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	new(numStrFmtCountryCultureSpecAtom).empty(
		destinationSpec)

	destinationSpec.IdNo = sourceSpec.IdNo

	destinationSpec.IdString = sourceSpec.IdString

	destinationSpec.Description = sourceSpec.Description

	destinationSpec.Tag = sourceSpec.Tag

	destinationSpec.CountryIdNo = sourceSpec.CountryIdNo

	destinationSpec.CountryIdString = sourceSpec.CountryIdString

	destinationSpec.CountryDescription =
		sourceSpec.CountryDescription

	destinationSpec.CountryTag = sourceSpec.CountryTag

	destinationSpec.CountryCultureName =
		sourceSpec.CountryCultureName

	destinationSpec.CountryCultureOfficialStateName =
		sourceSpec.CountryCultureOfficialStateName

	destinationSpec.CountryAbbreviatedName =
		sourceSpec.CountryAbbreviatedName

	arrayLen := len(sourceSpec.CountryAlternateNames)

	destinationSpec.CountryAlternateNames =
		make([]string, arrayLen)

	for i := 0; i < arrayLen; i++ {
		destinationSpec.CountryAlternateNames[i] =
			sourceSpec.CountryAlternateNames[i]
	}

	destinationSpec.CountryCodeTwoChar =
		sourceSpec.CountryCodeTwoChar

	destinationSpec.CountryCodeThreeChar =
		sourceSpec.CountryCodeThreeChar

	destinationSpec.CountryCodeNumber =
		sourceSpec.CountryCodeNumber

	destinationSpec.CurrencyCode =
		sourceSpec.CurrencyCode

	destinationSpec.CurrencyCodeNo =
		sourceSpec.CurrencyCodeNo

	destinationSpec.CurrencyName =
		sourceSpec.CurrencyName

	arrayLen = len(sourceSpec.CurrencySymbols)

	destinationSpec.CurrencySymbols =
		make([]rune, arrayLen)

	for i := 0; i < arrayLen; i++ {
		destinationSpec.CurrencySymbols[i] =
			sourceSpec.CurrencySymbols[i]
	}

	destinationSpec.MinorCurrencyName =
		sourceSpec.MinorCurrencyName

	arrayLen = len(sourceSpec.MinorCurrencySymbols)

	destinationSpec.MinorCurrencySymbols =
		make([]rune, arrayLen)

	for i := 0; i < arrayLen; i++ {
		destinationSpec.MinorCurrencySymbols[i] =
			sourceSpec.MinorCurrencySymbols[i]
	}

	err = destinationSpec.CurrencyNumStrFormat.CopyIn(
		&sourceSpec.CurrencyNumStrFormat,
		ePrefix.XCpy(
			"destinationSpec.CurrencyNumStrFormat"))

	if err != nil {
		return err
	}

	err = destinationSpec.SignedNumStrFormat.CopyIn(
		&sourceSpec.SignedNumStrFormat,
		ePrefix.XCpy(
			"destinationSpec.SignedNumStrFormat"))

	return err
}

//	setCountryUS
//
//	Receives a pointer to an instance of
//	NumStrFmtCountryCultureSpec and proceeds to configure
//	that instance with country and Number String
//	formatting specifications commonly used in the US
//	(United States).
//
//	These same formatting specifications are also used in
//	Australia and most of Canada.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields in 'countryNStrFmtSpec'
//	instance NumStrFmtCountryCultureSpec will be deleted
//	and overwritten with US country parameters.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	countryNStrFmtSpec	*NumStrFmtCountryCultureSpec
//
//		A pointer to a NumStrFmtCountryCultureSpec instance.
//		All the member variable data fields in this object
//		will be replaced with country and Number String
//		Formatting specifications commonly used in the
//		US (United States).
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string
//		which is included in all returned error
//		messages. Usually, it contains the name of the
//		calling method or methods listed as a function
//		chain.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref'
//		software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	err							error
//
//		If this method completes successfully, this
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (nStrFmtCountryCultureMech *numStrFmtCountryCultureSpecMech) setCountryUS(
	countryNStrFmtSpec *NumStrFmtCountryCultureSpec,
	errPrefDto *ePref.ErrPrefixDto) (err error) {

	if nStrFmtCountryCultureMech.lock == nil {
		nStrFmtCountryCultureMech.lock = new(sync.Mutex)
	}

	nStrFmtCountryCultureMech.lock.Lock()

	defer nStrFmtCountryCultureMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"nStrFmtCountryCultureMech."+
			"setCountryUS()",
		"")

	if err != nil {
		return err
	}

	if countryNStrFmtSpec == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'countryNStrFmtSpec' is invalid!\n"+
			"'countryNStrFmtSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	new(numStrFmtCountryCultureSpecAtom).empty(
		countryNStrFmtSpec)

	numStrFmtSpec := NumStrFormatSpec{}

	nopNumFieldSpec := new(NumStrNumberFieldSpec).NewNOP()

	countryNStrFmtSpec.CurrencyNumStrFormat,
		err = numStrFmtSpec.NewCurrencyNumFmtUS(
		nopNumFieldSpec,
		ePrefix.XCpy(
			"countryNStrFmtSpec.CurrencyNumStrFormat"))

	if err != nil {
		return err
	}

	countryNStrFmtSpec.SignedNumStrFormat,
		err = numStrFmtSpec.NewSignedNumFmtUS(
		nopNumFieldSpec,
		ePrefix.XCpy(
			"countryNStrFmtSpec.SignedNumStrFormat"))

	countryNStrFmtSpec.IdNo = 840
	countryNStrFmtSpec.IdString = "840"
	countryNStrFmtSpec.Description = "Country Setup"
	countryNStrFmtSpec.Tag = ""
	countryNStrFmtSpec.CountryIdNo = 840
	countryNStrFmtSpec.CountryIdString = "840"
	countryNStrFmtSpec.CountryDescription = "Country Setup - United States"
	countryNStrFmtSpec.CountryTag = ""
	countryNStrFmtSpec.CountryCultureName = "United States"
	countryNStrFmtSpec.CountryAbbreviatedName = "USA"

	countryNStrFmtSpec.CountryAlternateNames =
		[]string{
			"The United States of America",
			"United States of America",
			"America"}

	countryNStrFmtSpec.CountryCodeTwoChar = "US"
	countryNStrFmtSpec.CountryCodeThreeChar = "USA"
	countryNStrFmtSpec.CountryCodeNumber = "840"
	countryNStrFmtSpec.CurrencyDecimalDigits = 2
	countryNStrFmtSpec.CurrencyCode = "USD"
	countryNStrFmtSpec.CurrencyCodeNo = "840"
	countryNStrFmtSpec.CurrencyName = "Dollar"
	countryNStrFmtSpec.CurrencySymbols = []rune{'\U00000024'}
	countryNStrFmtSpec.MinorCurrencyName = "Cent"
	countryNStrFmtSpec.MinorCurrencySymbols = []rune{'\U000000a2'}

	return err
}
