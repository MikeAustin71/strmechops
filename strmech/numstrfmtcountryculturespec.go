package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

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
	//	The ISO 3166 name of the country or culture
	//	identified by the current Country Culture
	//	Specification instance.

	CountryCultureOfficialStateName string
	//	Optional
	//	The ISO 3166 official state name of the country
	//	or culture identified by the current Country
	//	Culture Specification instance.

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
	//	The unique ISO 3166-1 alpha-2 Two Character code
	//	identifying the country or culture associated
	//	with the current Country Culture Specification
	//	instance.

	CountryCodeThreeChar string
	//	Optional
	//	The unique ISO 3166-1 alpha-3 Three Character code
	//	identifying the country or culture associated with
	//	the current Country Culture Specification instance.

	CountryCodeNumber string
	//	Optional
	//	The official ISO 3166-1 numeric code identifier
	//	for the country or culture associated with the
	//	current Country Culture Specification instance.

	CurrencyCode string
	//	Optional
	//	The official ISO 4217 currency code associated
	//	with the country or culture identified by the
	//	current Country Culture Specification instance.

	CurrencyCodeNo string
	//	Optional
	//	The official ISO 4217 currency code number
	//	associated with the country or culture identified
	//	by the current Country Culture Specification
	//	instance.

	CurrencyName string
	//	Optional
	//	The official ISO 4217 currency name associated
	//	with the country or culture identified by the
	//	current Country Culture Specification
	//	instance.

	CurrencySymbols []rune
	//	Optional. Not required for Currency Number
	//	String Formatting.
	//
	//	The official ISO 4217 currency symbol or symbols
	//	for the country or culture identified by the
	//	current Country Culture Specification instance.

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

	CurrencyNumStrFormat NumStrFormatSpec
	//	Required for Currency Number String
	//	Formatting.
	//
	//	This NumStrFormatSpec instance
	//	contains all the parameters necessary
	//	to produce a formatted Currency
	//	Number String.

	SignedNumStrFormat NumStrFormatSpec
	//	Required for Signed Number String
	//	Formatting.
	//
	//	This NumStrFormatSpec instance
	//	contains all the parameters necessary
	//	to produce a formatted Signed
	//	Number String.

	lock *sync.Mutex
}

//	CopyIn
//
//	Copies the data fields from an incoming instance of
//	NumStrFmtCountryCultureSpec ('sourceCountryCultureSpec')
//	to the data fields of the current
//	NumStrFmtCountryCultureSpec instance
//	('nStrFmtCountryCultureSpec').
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	All the member variable data values in the current
//	NumStrFmtCountryCultureSpec instance
//	('nStrFmtCountryCultureSpec') will	be deleted and
//	replaced.
//
//	No data validation is performed on input parameter,
//	'incomingSignedNumFmt'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	sourceCountryCultureSpec		*NumStrFmtCountryCultureSpec
//
//		A pointer to an instance of NumStrFmtCountryCultureSpec.
//		This method will NOT change the values of internal
//		member variables contained in this instance.
//
//		All data values in this NumStrFmtCountryCultureSpec
//		instance will be copied to current
//		NumStrFmtCountryCultureSpec instance
//		('nStrFmtCountryCultureSpec').
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (nStrFmtCountryCultureSpec *NumStrFmtCountryCultureSpec) CopyIn(
	sourceCountryCultureSpec *NumStrFmtCountryCultureSpec,
	errorPrefix interface{}) error {

	if nStrFmtCountryCultureSpec.lock == nil {
		nStrFmtCountryCultureSpec.lock = new(sync.Mutex)
	}

	nStrFmtCountryCultureSpec.lock.Lock()

	defer nStrFmtCountryCultureSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFmtCountryCultureSpec."+
			"CopyIn()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtCountryCultureSpecMech).
		copyCountryCulture(
			nStrFmtCountryCultureSpec,
			sourceCountryCultureSpec,
			ePrefix.XCpy(
				"nStrFmtCountryCultureSpec<-"))
}

//	CopyOut
//
//	Returns a deep copy of the current
//	NumStrFmtCountryCultureSpec instance.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	NumStrFmtCountryCultureSpec
//
//		If this method completes successfully and no errors are
//		encountered, this parameter will return a deep copy of
//		the current NumStrFmtCountryCultureSpec instance.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (nStrFmtCountryCultureSpec *NumStrFmtCountryCultureSpec) CopyOut(
	errorPrefix interface{}) (
	NumStrFmtCountryCultureSpec,
	error) {

	if nStrFmtCountryCultureSpec.lock == nil {
		nStrFmtCountryCultureSpec.lock = new(sync.Mutex)
	}

	nStrFmtCountryCultureSpec.lock.Lock()

	defer nStrFmtCountryCultureSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	var newNStrFmtCountryCulture NumStrFmtCountryCultureSpec

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFmtCountryCultureSpec."+
			"CopyOut()",
		"")

	if err != nil {
		return newNStrFmtCountryCulture, err
	}

	err = new(numStrFmtCountryCultureSpecMech).
		copyCountryCulture(
			&newNStrFmtCountryCulture,
			nStrFmtCountryCultureSpec,
			ePrefix.XCpy(
				"newNStrFmtCountryCulture<-"))

	return newNStrFmtCountryCulture, err
}

//	Empty
//
//	Resets all internal member variables for the current
//	instance of NumStrFmtCountryCultureSpec to their
//	initial or zero values.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete all pre-existing internal
//	member variable data values in the current instance
//	of NumStrFmtCountryCultureSpec.
//
// ------------------------------------------------------------------------
//
// # Input Parameters
//
//	NONE
//
// ------------------------------------------------------------------------
//
// # Return Values
//
//	NONE
func (nStrFmtCountryCultureSpec *NumStrFmtCountryCultureSpec) Empty() {

	if nStrFmtCountryCultureSpec.lock == nil {
		nStrFmtCountryCultureSpec.lock = new(sync.Mutex)
	}

	nStrFmtCountryCultureSpec.lock.Lock()

	new(numStrFmtCountryCultureSpecAtom).empty(
		nStrFmtCountryCultureSpec)

	nStrFmtCountryCultureSpec.lock.Unlock()

	nStrFmtCountryCultureSpec.lock = nil
}

//	Equal
//
//	Receives a pointer to another instance of
//	NumStrFormatSpec and proceeds to compare its internal
//	member variables to those of the current
//	NumStrFormatSpec instance in order to determine if
//	they are equivalent.
//
//	A boolean flag showing the result of this comparison
//	is returned. If the member variables for both
//	instances are equal in all respects, this flag is set
//	to 'true'. Otherwise, this method returns 'false'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	incomingCountryCulture		*NumStrFmtCountryCultureSpec
//
//		A pointer to an external instance of
//		NumStrFormatSpec. The internal member variable
//		data values in this instance will be compared to
//		those in the current instance of
//		NumStrFmtCountryCultureSpec. The results of this
//		comparison will be returned to the calling
//		function as a boolean value.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	bool
//
//		If the internal member variable data values
//		contained in input parameter 'incomingCountryCulture'
//		are equivalent in all respects to those contained
//		in the current instance of
//		'NumStrFmtCountryCultureSpec', this return value
//		will be set to 'true'.
//
//		Otherwise, this method will return 'false'.
func (nStrFmtCountryCultureSpec *NumStrFmtCountryCultureSpec) Equal(
	incomingCountryCulture *NumStrFmtCountryCultureSpec) bool {

	if nStrFmtCountryCultureSpec.lock == nil {
		nStrFmtCountryCultureSpec.lock = new(sync.Mutex)
	}

	nStrFmtCountryCultureSpec.lock.Lock()

	defer nStrFmtCountryCultureSpec.lock.Unlock()

	return new(numStrFmtCountryCultureSpecAtom).equal(
		nStrFmtCountryCultureSpec,
		incomingCountryCulture)
}

//	New
//
//	Creates and returns a new instance of
//	NumStrFmtCountryCultureSpec.
//
//	The returned instance will be configured with the
//	currency and signed number string formatters. In
//	addition, the Country/Culture name will also be
//	populated.
//
//	While the returned instance is sufficient for
//	generating formatted number strings for currency
//	and signed numbers, all the description member
//	variables will remain empty. These description
//	member variables are public variables and may be
//	populated manually by the user as needed.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	countryCultureName			string
//
//		It is recommended that this parameter be set
//		to the ISO 3166 name of the country or culture.
//
//	currencyNumStrFormat		NumStrFormatSpec
//
//		An instance of NumStrFormatSpec which will be
//		used to generate Number Strings using the
//		currency format associated with the
//		designated Country or Culture.
//
//	signedNumStrFormat			NumStrFormatSpec
//
//		An instance of NumStrFormatSpec which will be
//		used to generate Number Strings using the
//		signed number format associated with the
//		designated Country or Culture.
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	NumStrFmtCountryCultureSpec
//
//		If this method completes successfully, a partially
//		completed instance of NumStrFmtCountryCultureSpec
//		will be returned.
//
//		While the returned instance is sufficient for
//		generating formatted number strings for currency
//		and signed numbers, all the description member
//		variables will remain empty. These description
//		member variables are public variables and may be
//		populated manually by the user as needed.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (nStrFmtCountryCultureSpec *NumStrFmtCountryCultureSpec) New(
	countryCultureName string,
	currencyNumStrFormat NumStrFormatSpec,
	signedNumStrFormat NumStrFormatSpec,
	errorPrefix interface{}) (
	NumStrFmtCountryCultureSpec,
	error) {

	if nStrFmtCountryCultureSpec.lock == nil {
		nStrFmtCountryCultureSpec.lock = new(sync.Mutex)
	}

	nStrFmtCountryCultureSpec.lock.Lock()

	defer nStrFmtCountryCultureSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	var newCountryCultureSpec NumStrFmtCountryCultureSpec

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFmtCountryCultureSpec."+
			"CopyIn()",
		"")

	if err != nil {
		return newCountryCultureSpec, err
	}

	fmtCountryCultureAtom := numStrFmtCountryCultureSpecAtom{}

	err = fmtCountryCultureAtom.copyNumStrFormatSpec(
		&newCountryCultureSpec.CurrencyNumStrFormat,
		&currencyNumStrFormat,
		ePrefix.XCpy(
			"newCountryCultureSpec<-"+
				"currencyNumStrFormat"))

	if err != nil {
		return newCountryCultureSpec, err
	}

	err = fmtCountryCultureAtom.copyNumStrFormatSpec(
		&newCountryCultureSpec.SignedNumStrFormat,
		&signedNumStrFormat,
		ePrefix.XCpy(
			"newCountryCultureSpec<-"+
				"signedNumStrFormat"))

	if err != nil {
		return newCountryCultureSpec, err
	}

	newCountryCultureSpec.CountryCultureName =
		countryCultureName

	return newCountryCultureSpec, err
}

//	SetCurrencyNumberFieldSpec
//
//	Deletes and resets the Currency Number Field
//	Specification for the current instance of
//	NumStrFmtCountryCultureSpec.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	No data validation is performed on input parameter,
//	'currencyNumFieldSpec'.
//
//	For the current instance of NumStrFmtCountryCultureSpec,
//	this method will modify the internal member variable:
//
//	NumStrFmtCountryCultureSpec.CurrencyNumStrFormat.numberFieldSpec
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	currencyNumFieldSpec		NumStrNumberFieldSpec
//
//		Contains the Number String Number Field Specification
//		which will be copied to the Currency Number String
//		Number Field Specification encapsulated within
//		the current instance of NumStrFmtCountryCultureSpec.
func (nStrFmtCountryCultureSpec *NumStrFmtCountryCultureSpec) SetCurrencyNumberFieldSpec(
	currencyNumFieldSpec NumStrNumberFieldSpec,
	errorPrefix interface{}) error {

	if nStrFmtCountryCultureSpec.lock == nil {
		nStrFmtCountryCultureSpec.lock = new(sync.Mutex)
	}

	nStrFmtCountryCultureSpec.lock.Lock()

	defer nStrFmtCountryCultureSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFmtCountryCultureSpec."+
			"SetCurrencyNumberFieldSpec()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtCountryCultureSpecElectron).
		copyNumberFieldSpec(
			&nStrFmtCountryCultureSpec.CurrencyNumStrFormat.numberFieldSpec,
			&currencyNumFieldSpec,
			ePrefix.XCpy(
				"nStrFmtCountryCultureSpec<-currencyNumFieldSpec"))
}

//	SetCurrencyNumStrFormatSpec
//
//	Deletes and resets the Currency Number String Format
//	Specification for the current instance of
//	NumStrFmtCountryCultureSpec.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	All data values in the Currency Number String Format
//	Specification for the current instance of
//	NumStrFmtCountryCultureSpec will be deleted and reset
//	to new values.
//
//	No data validation is performed on input parameter,
//	'currencyNumStrFormat'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	currencyNumStrFormat		NumStrFormatSpec
//
//		An instance of NumStrFormatSpec which will be
//		used to generate Number Strings using the
//		currency format associated with the
//		designated Country or Culture.
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (nStrFmtCountryCultureSpec *NumStrFmtCountryCultureSpec) SetCurrencyNumStrFormatSpec(
	currencyNumStrFormat NumStrFormatSpec,
	errorPrefix interface{}) error {

	if nStrFmtCountryCultureSpec.lock == nil {
		nStrFmtCountryCultureSpec.lock = new(sync.Mutex)
	}

	nStrFmtCountryCultureSpec.lock.Lock()

	defer nStrFmtCountryCultureSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFmtCountryCultureSpec."+
			"SetCurrencyNumStrFormatSpec()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtCountryCultureSpecAtom).
		copyNumStrFormatSpec(
			&nStrFmtCountryCultureSpec.CurrencyNumStrFormat,
			&currencyNumStrFormat,
			ePrefix.XCpy(
				"nStrFmtCountryCultureSpec<-"+
					"currencyNumStrFormat"))
}

//	SetSignedNumberFieldSpec
//
//	Deletes and resets the Signed Number Field
//	Specification for the current instance of
//	NumStrFmtCountryCultureSpec.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	No data validation is performed on input parameter,
//	'signedNumFieldSpec'.
//
//	For the current instance of NumStrFmtCountryCultureSpec,
//	this method will modify the internal member variable:
//
//	NumStrFmtCountryCultureSpec.SignedNumStrFormat.numberFieldSpec
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	signedNumFieldSpec		NumStrNumberFieldSpec
//
//		Contains the Number String Number Field Specification
//		which will be copied to the Signed Number String
//		Number Field Specification encapsulated within
//		the current instance of NumStrFmtCountryCultureSpec.
func (nStrFmtCountryCultureSpec *NumStrFmtCountryCultureSpec) SetSignedNumberFieldSpec(
	signedNumFieldSpec NumStrNumberFieldSpec,
	errorPrefix interface{}) error {

	if nStrFmtCountryCultureSpec.lock == nil {
		nStrFmtCountryCultureSpec.lock = new(sync.Mutex)
	}

	nStrFmtCountryCultureSpec.lock.Lock()

	defer nStrFmtCountryCultureSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFmtCountryCultureSpec."+
			"SetSignedNumberFieldSpec()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtCountryCultureSpecElectron).
		copyNumberFieldSpec(
			&nStrFmtCountryCultureSpec.SignedNumStrFormat.numberFieldSpec,
			&signedNumFieldSpec,
			ePrefix.XCpy(
				"nStrFmtCountryCultureSpec<-signedNumFieldSpec"))
}

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

// numStrFmtCountryCultureSpecAtom
//
// Provides helper methods for type
// NumStrFmtCountryCultureSpec
type numStrFmtCountryCultureSpecElectron struct {
	lock *sync.Mutex
}

//	copyNumberFieldSpec
//
//	Copies a source Number String Number Field
//	Specification (NumStrNumberFieldSpec) to a destination
//	Number String Number Field Specification.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	No data validation is performed on input parameter,
//	'sourceNumFieldSpec'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	destinationNumFieldSpec			*NumStrNumberFieldSpec
//
//		A pointer to a NumStrNumberFieldSpec
//		instance. The Number String Number Field
//		Specification from input parameter
//		'sourceNumFieldSpec' will be copied to this
//		parameter 'destinationNumFieldSpec'.
//
//		All data values contained in
//		'destinationNumFieldSpec' will be deleted and
//		reset to new values extracted from
//		'sourceNumFieldSpec'.
//
//	sourceNumFieldSpec				*NumStrNumberFieldSpec
//
//		A pointer to a NumStrNumberFieldSpec
//		instance. The Number String Number Field
//		Specification from this parameter
//		'sourceNumFieldSpec' will be copied to the
//		destination input parameter,
//		'destinationNumFieldSpec'.
//
//		The data values contained in 'sourceNumFieldSpec'
//		WILL NOT be modified by this method.
//
//	errPrefDto						*ePref.ErrPrefixDto
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
func (nStrFmtCountryCultureElectron *numStrFmtCountryCultureSpecElectron) copyNumberFieldSpec(
	destinationNumFieldSpec *NumStrNumberFieldSpec,
	sourceNumFieldSpec *NumStrNumberFieldSpec,
	errPrefDto *ePref.ErrPrefixDto) error {

	if nStrFmtCountryCultureElectron.lock == nil {
		nStrFmtCountryCultureElectron.lock = new(sync.Mutex)
	}

	nStrFmtCountryCultureElectron.lock.Lock()

	defer nStrFmtCountryCultureElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtCountryCultureSpecElectron."+
			"copyNumberFieldSpec()",
		"")

	if err != nil {
		return err
	}

	if destinationNumFieldSpec == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'destinationNumFieldSpec' is invalid!\n"+
			"'countryCultureSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if sourceNumFieldSpec == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'sourceNumFieldSpec' is invalid!\n"+
			"'sourceNumFieldSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	destinationNumFieldSpec.Empty()

	return destinationNumFieldSpec.CopyIn(
		sourceNumFieldSpec,
		ePrefix.XCpy(
			"<-signedNumFieldSpec"))
}
