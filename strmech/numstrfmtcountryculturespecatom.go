package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// numStrFmtCountryCultureSpecAtom
//
// Provides helper methods for type
// NumStrFmtCountryCultureSpec
type numStrFmtCountryCultureSpecAtom struct {
	lock *sync.Mutex
}

func (nStrFmtCountryCultureAtom *numStrFmtCountryCultureSpecAtom) empty(
	countryCultureSpec *NumStrFmtCountryCultureSpec) {

	if nStrFmtCountryCultureAtom.lock == nil {
		nStrFmtCountryCultureAtom.lock = new(sync.Mutex)
	}

	nStrFmtCountryCultureAtom.lock.Lock()

	defer nStrFmtCountryCultureAtom.lock.Unlock()

	if countryCultureSpec == nil {
		return
	}

	countryCultureSpec.IdNo = 0

	countryCultureSpec.IdString = ""

	countryCultureSpec.Description = ""

	countryCultureSpec.Tag = ""

	countryCultureSpec.CountryIdNo = 0

	countryCultureSpec.CountryIdString = ""

	countryCultureSpec.CountryDescription = ""

	countryCultureSpec.CountryTag = ""

	countryCultureSpec.CountryCultureName = ""

	countryCultureSpec.CountryCultureOfficialStateName = ""

	countryCultureSpec.CountryAbbreviatedName = ""

	countryCultureSpec.CountryAlternateNames = nil

	countryCultureSpec.CountryCodeTwoChar = ""

	countryCultureSpec.CountryCodeThreeChar = ""

	countryCultureSpec.CountryCodeNumber = ""

	countryCultureSpec.CurrencyCode = ""

	countryCultureSpec.CurrencyCodeNo = ""

	countryCultureSpec.CurrencyName = ""

	countryCultureSpec.CurrencySymbols = nil

	countryCultureSpec.MinorCurrencyName = ""

	countryCultureSpec.MinorCurrencySymbols = nil

	countryCultureSpec.CurrencyNumStrFormat.Empty()

	countryCultureSpec.SignedNumStrFormat.Empty()
}

//	copyNumStrFormatSpec
//
//	Copies a source instance of NumStrFormatSpec to a
//	destination instance of NumStrFormatSpec.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and reset the data values
//	for the destination instance of NumStrFormatSpec,
//	input parameter 'destinationNStrFormatSpec'.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	destinationNStrFormatSpec	*NumStrFormatSpec
//
//		A pointer to an instance of NumStrFormatSpec. The
//		data values contained in this input parameter
//		('destinationNStrFormatSpec') will be deleted and
//		reset to new values extracted from input parameter,
//		'sourceNStrFormatSpec'.
//
//		'destinationNStrFormatSpec' is the destination for
//		this copy operation.
//
//	sourceNStrFormatSpec		*NumStrFormatSpec
//
//		A pointer to an instance of NumStrFormatSpec. The
//		data values contained in this input parameter
//		('sourceNStrFormatSpec') will be copied to input
//		parameter, 'destinationNStrFormatSpec'.
//
//		'sourceNStrFormatSpec' is the data source for
//		this copy operation.
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
//	error
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
func (nStrFmtCountryCultureAtom *numStrFmtCountryCultureSpecAtom) copyNumStrFormatSpec(
	destinationNStrFormatSpec *NumStrFormatSpec,
	sourceNStrFormatSpec *NumStrFormatSpec,
	errPrefDto *ePref.ErrPrefixDto) error {

	if nStrFmtCountryCultureAtom.lock == nil {
		nStrFmtCountryCultureAtom.lock = new(sync.Mutex)
	}

	nStrFmtCountryCultureAtom.lock.Lock()

	defer nStrFmtCountryCultureAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtCountryCultureSpecAtom."+
			"copyNumStrFormatSpec()",
		"")

	if err != nil {
		return err
	}

	if destinationNStrFormatSpec == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'destinationNStrFormatSpec' is invalid!\n"+
			"'destinationNStrFormatSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if sourceNStrFormatSpec == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'sourceNStrFormatSpec' is invalid!\n"+
			"'sourceNStrFormatSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	destinationNStrFormatSpec.Empty()

	return destinationNStrFormatSpec.CopyIn(
		sourceNStrFormatSpec,
		ePrefix.XCpy(
			"destinationNStrFormatSpec<-"+
				"sourceNStrFormatSpec"))
}

//	equal
//
//	Receives a pointer to two instances of
//	NumStrFmtCountryCultureSpec and proceeds to compare
//	their member variables in order to determine if they
//	are equivalent.
//
//	A boolean flag showing the result of this comparison
//	is returned. If the member variables for both instances
//	are equal in all respects, this flag is set to 'true'.
//	Otherwise, this method returns 'false'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	countryCultureOne			*NumStrFmtCountryCultureSpec
//
//		An instance of NumStrFmtCountryCultureSpec.
//		Internal member variables from 'countryCultureOne'
//		will be compared to those of 'signedNumFmtSpec2' to
//		determine if both instances are equivalent.
//
//
//	countryCultureTwo			*NumStrFmtCountryCultureSpec
//
//		An instance of NumStrFmtCountryCultureSpec
//		Internal member variables from 'countryCultureTwo'
//		will be compared to those of 'countryCultureOne' to
//		determine if both instances are equivalent.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		If the comparison of 'countryCultureOne' and
//		'countryCultureTwo' shows that all internal
//		member variables are equivalent, this method
//		will return a boolean value of 'true'.
//
//		If the two instances are NOT equal, this method
//		will return a boolean value of 'false' to the
//		calling function.
func (nStrFmtCountryCultureAtom *numStrFmtCountryCultureSpecAtom) equal(
	countryCultureOne *NumStrFmtCountryCultureSpec,
	countryCultureTwo *NumStrFmtCountryCultureSpec) bool {

	if nStrFmtCountryCultureAtom.lock == nil {
		nStrFmtCountryCultureAtom.lock = new(sync.Mutex)
	}

	nStrFmtCountryCultureAtom.lock.Lock()

	defer nStrFmtCountryCultureAtom.lock.Unlock()

	if countryCultureOne == nil {

		return false
	}

	if countryCultureTwo == nil {

		return false
	}

	if countryCultureOne.IdNo != countryCultureTwo.IdNo {

		return false
	}

	if countryCultureOne.IdString != countryCultureTwo.IdString {

		return false
	}

	if countryCultureOne.Description != countryCultureTwo.Description {

		return false
	}

	if countryCultureOne.Tag != countryCultureTwo.Tag {

		return false
	}

	if countryCultureOne.CountryIdNo != countryCultureTwo.CountryIdNo {

		return false
	}

	if countryCultureOne.CountryIdString != countryCultureTwo.CountryIdString {

		return false
	}

	if countryCultureOne.CountryDescription !=
		countryCultureTwo.CountryDescription {

		return false
	}

	if countryCultureOne.CountryTag != countryCultureTwo.CountryTag {

		return false
	}

	if countryCultureOne.CountryCultureName !=
		countryCultureTwo.CountryCultureName {

		return false
	}

	if countryCultureOne.CountryCultureOfficialStateName !=
		countryCultureTwo.CountryCultureOfficialStateName {

		return false
	}

	if countryCultureOne.CountryAbbreviatedName !=
		countryCultureTwo.CountryAbbreviatedName {

		return false
	}

	arrayLen := len(countryCultureTwo.CountryAlternateNames)

	if len(countryCultureOne.CountryAbbreviatedName) !=
		arrayLen {

		return false
	}

	for i := 0; i < arrayLen; i++ {

		if countryCultureOne.CountryAlternateNames[i] !=
			countryCultureTwo.CountryAlternateNames[i] {

			return false
		}
	}

	countryCultureOne.CountryCodeTwoChar =
		countryCultureTwo.CountryCodeTwoChar

	countryCultureOne.CountryCodeThreeChar =
		countryCultureTwo.CountryCodeThreeChar

	countryCultureOne.CountryCodeNumber =
		countryCultureTwo.CountryCodeNumber

	countryCultureOne.CurrencyCode =
		countryCultureTwo.CurrencyCode

	countryCultureOne.CurrencyCodeNo =
		countryCultureTwo.CurrencyCodeNo

	countryCultureOne.CurrencyName =
		countryCultureTwo.CurrencyName

	arrayLen = len(countryCultureTwo.CurrencySymbols)

	if len(countryCultureOne.CurrencySymbols) !=
		arrayLen {

		return false
	}

	for i := 0; i < arrayLen; i++ {

		if countryCultureOne.CurrencySymbols[i] !=
			countryCultureTwo.CurrencySymbols[i] {

			return false
		}
	}

	countryCultureOne.MinorCurrencyName =
		countryCultureTwo.MinorCurrencyName

	arrayLen = len(countryCultureTwo.MinorCurrencySymbols)

	if len(countryCultureOne.MinorCurrencySymbols) !=
		arrayLen {

		return false
	}

	for i := 0; i < arrayLen; i++ {

		if countryCultureOne.MinorCurrencySymbols[i] !=
			countryCultureTwo.MinorCurrencySymbols[i] {

			return false
		}
	}

	areEqual := countryCultureOne.CurrencyNumStrFormat.Equal(
		&countryCultureTwo.CurrencyNumStrFormat)

	if !areEqual {

		return false
	}

	areEqual = countryCultureOne.SignedNumStrFormat.Equal(
		&countryCultureTwo.SignedNumStrFormat)

	return areEqual
}
